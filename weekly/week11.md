# Weekly Note: 17 July - 23 July

## Delivered this week

+ Thought about adding a readOnly field but decided against it since it will go against kube's way of handling resources.
    - this is present in `hyperschema` but we don't use it now.
+ Got rid of the `enqueue` for informers. Fixed the cache invalidation.
+ Added custom deepcopy functions because we use interfaces now. Or else, it gives this error: https://gist.github.com/nikhita/8cc04bfc68357695625da2e1ac92ba9a#file-deecopy-txt-L10.
    - maybe we can have a special tag for deepcopy gen to avoid writing big deepcopy functions.
    - used the one for `Table` in meta as a reference.
+ Added custom fuzzer functions for interface fields. Why we need this: interface can have invalid JSON values too. This breaks round trip tests. Also, https://github.com/kubernetes/kubernetes/pull/49307.
+ Proto doesn't work too because of the interface fiels. A similar issue was faced in `Table`. Decided to not support proto for now.
    - there is a file which contains all packages where proto is supported. removed apiextensions/v1beta1 from there.
+ Changed the JSON tags on the `$ref` fields.
+ Had a bit of trouble with `update-staging-godeps` that was also reported with someone else (https://github.com/kubernetes/kubernetes/issues/49210) Fixed with a clean checkout.
+ Refractored some more code!