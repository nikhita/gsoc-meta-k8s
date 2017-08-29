# Weekly Note: 3 July - 10 July

## Delivered this week

+ For the proposal: only open question is about the order of validations. We would probably overengineer this if we define the order via a slice. This could be part of the webhook validation type. Though even multiple webhooks is probably too much.
+ Created a temp PR for conversion of JSONSchemaRef: https://github.com/nikhita/jsonreference/pull/1. Edit: decided not to use this, but the constructor which is already available upstream.
+ We probably don't need the 5 bools in `JSONSchemaRef`  The parsing code only checks that it is a URL. everything after that is unconditionally, i.e. without added value for validation of the `JSONSchemaRef` value.
+ The [update-on-CRD case](https://github.com/kubernetes/kubernetes/pull/47263#discussion_r125323976) did not work because of no informers! :/ Invalidated the storage cache.
+ Also, added a check that the url created by `NewRef` is a valid url.
+ Discussed the non-trivial cases of `OneOf`
    - one can construct json objects where the default values in a schema do not make sense.
    - this `{allOf: [{minItems: 1, items:{default:7}}, {items:[{default:8}]}]}` will give either `[7]` or `[8]`, both are valid.
+ For default:
    - default has no direct impact on validation. so it does not need to satisfy the schema.
    - most validators dont take defaulting into account for this very reason.
    - main insight: if defaulting does not influence validation, the defaulting of slices (in the sense of appending missing elements) cannot happen. :sparkle:
+ Updated the Ref field to a string. When I do that, it validates _every time_. Still figuring out why changing one field is affecting the whole validation to pass.
+ Updated the gengo PR!
    - discussed with lavalamp and got the PR merged! :boom:
    - simplified and refractored the existing code as well. the `buildTreeFromTypes` func is now very simple!
+ Plan to add Custom resources spec dynamically to main OpenAPI spec using the aggregator. (https://github.com/kubernetes/kubernetes/pull/46734)
+ Worked on updating the vendored gengo in k/k. And had trouble with godep...again!
    - kept hitting this error: https://github.com/kubernetes/kubernetes/issues/35356#issuecomment-258054506, https://github.com/kubernetes/kubernetes/pull/42967#issuecomment-288593136.
    - kept breaking on the bazel update.
        + after godep-save, this lists as an untracked file: `vendor/github.com/docker/docker/project/`.
        + then bazel panics in the next step.
    - turns out `CONTRIBUTORS.md` from `docker/docker` has a symlink which keeps breaking bazel. https://github.com/moby/moby/issues/31825.
    - adding a simple `rm` in godep-save with a good comment that it's a workaround would be good. (edit: done)
    - The main problem was why godep was pulling down the `CONTRIBUTORS.md` file, there should no dependecies on this!
    - Solution:
        + godep pulls it down because it a "legal" file. https://github.com/tools/godep/pull/522.
        + Related: https://github.com/moby/moby/pull/31847.
    - continuing on the vendoring, get this huge change now: https://github.com/nikhita/kubernetes/pull/5. apimachinery is changed! :scream:
    - Note: why `godep-save.sh` takes such a long time:
        + The problem is that godep accesses k8s.io for every kube package (package, not repo!)
        + k8s.io returns the meta info for "go get", i.e. where to find k8s.io/kubernetes on github there is no caching for that.
        + in CI everything runs very near to the http://k8s.io server. So roundtrip is much much faster in CI => godep restore takes only 2 minutes or so in CI .
    - The apimachinery problem is solved by making sure it does not exist in GOPATH before. Now there is a PR to make sure this does not happen anymore! :smile:
    - After some more trouble with gengo (starting with a clean gopath + `make clean`), had the gengo PR out! https://github.com/kubernetes/kubernetes/pull/48630.
+ To top of everything, deepcopy generation broke. https://github.com/kubernetes/kubernetes/issues/48675.
