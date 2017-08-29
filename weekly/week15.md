# Weekly Note: 14 August - 20 August

## Delivered this week

+ Refractored the existing feature gates even more.
+ Enabled feature gate in integration tests.
+ Changed `CustomResourceValidation` to `*CustomResourceValidation`.
+ Published the proposal: https://github.com/kubernetes/community/pull/913.
+ Started work on the subresources prototype.
    - added the status and scale subresources.
    - updated docs for json path: https://github.com/kubernetes/kubernetes/pull/50444, https://github.com/kubernetes/kubernetes.github.io/pull/4986.
    - todo: add tests.
+ Noticed that the spec and status for CRDs did not support incrementing `metadata.Generation`. Created: https://github.com/kubernetes/kubernetes/pull/50764.
+ Addressed reviews in the proposal + made some changes to it.
+ Added custom marshal functions. But round-tripping broke so needed to update fuzzers.
