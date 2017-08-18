## Benchmarks

Bechmarks for validation using:
1. go-openapi
2. go-jsonschema

Note: The files in this directory were *not* used for testing. Standard test data for JSON Schema validation was used. However, the JSON files in this directory are provided as examples.

```
BenchmarkOpenAPI-4        	    5000	    240116 ns/op	   71790 B/op	     649 allocs/op
BenchmarkGoJSONSchema-4   	    5000	    314191 ns/op	   66629 B/op	     965 allocs/op
```