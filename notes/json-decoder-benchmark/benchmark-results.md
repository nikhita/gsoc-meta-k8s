# JSON Decoder Comparison

This is in relation to:
- PR: https://github.com/kubernetes/kubernetes/pull/44026
- Issue: https://github.com/kubernetes/kubernetes/issues/30213

The comparison is between Kubernetes' custom JSON package and Go's JSON package.

## Benchmark functions

I ran benchmark tests for the json package:

- `json.go` - https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apimachinery/pkg/util/json/json.go
- `json_test.go` - https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apimachinery/pkg/util/json/json_test.go

I used the test cases that were present in the `json_test.go` file. Note: This is really a worst-case, i.e. a real-world object will have a much smaller int-to-non-int ratio. So, we should be fine.

**Custom JSON package:**

```go
func BenchmarkEvaluateTypes(b *testing.B) {
    for i := 0; i < b.N; i++ {
        for _, tc := range testCases {
            inputJSON := fmt.Sprintf(`{"data":%s}`, tc.In)
            m := map[string]interface{}{}
            if tc.Err != true {
                err := Unmarshal([]byte(inputJSON), &m)
                if err != nil {
                    b.Errorf("%s: error decoding: %v", tc.In, err)
                }
            }

            _, err := Marshal(m)
            if err != nil {
                b.Errorf("%s: error encoding: %v", tc.In, err)
            }
          }
    }
}
```
**Default JSON package:**

```go
func BenchmarkEvaluateTypes(b *testing.B) {
    for i := 0; i < b.N; i++ {
        for _, tc := range testCases {
            inputJSON := fmt.Sprintf(`{"data":%s}`, tc.In)
            m := map[string]interface{}{}
            if tc.Err != true {
                err := gojson.Unmarshal([]byte(inputJSON), &m)
                if err != nil {
                    b.Errorf("%s: error decoding: %v", tc.In, err)
                }
            }

            _, err := gojson.Marshal(m)
            if err != nil {
                b.Errorf("%s: error encoding: %v", tc.In, err)
            }
         }
    }
}
```

## Results

### Benchmark tests for both cases:

CASE 1: Custom JSON package

```bash
$ go test ./staging/src/k8s.io/apimachinery/pkg/util/json/ -benchmem -run=XXX -bench=BenchmarkEvaluateTypes -cpuprofile=cpu.out | tee bench0
```

Output:

![output](images/1.png?raw=true)

CASE 2: Default JSON package

```bash
$ go test ./staging/src/k8s.io/apimachinery/pkg/util/json/ -benchmem -run=XXX -bench=BenchmarkEvaluateTypes -cpuprofile=cpu.out | tee bench1
```

Output:

![output](images/2.png?raw=true)

### Comparison of benchmarks

```bash
$ benchcmp bench1 bench0
benchmark                    old ns/op      new ns/op    delta
BenchmarkEvaluateTypes-4     144682         172127       +18.97%

benchmark                    old allocs  new allocs   delta
BenchmarkEvaluateTypes-4     865         1022         +18.15%

benchmark                    old bytes      new bytes     delta
BenchmarkEvaluateTypes-4     43715          71183         +62.83%
```

### CPU Profiling

Analyzing how much CPU usage `convertMapNumbers` takes:

![cpu-profiling](images/3.png?raw=true)


- Unmarshalling takes a total of 910 ms.
- `convertMapNumbers` takes 90ms overall.

### Hotspots in `convertMapNumbers`

![hotspots-1](images/4.png?raw=true)

![hotspots-2](images/5.png?raw=true)

### Memory profiling

Doing a memory profiling for `Unmarshal` and `convertMapNumbers`:

![unmarshal-memory-profiling](images/6.png?raw=true)

![convertMapNumbers-memory-profiling](images/7.png?raw=true)


