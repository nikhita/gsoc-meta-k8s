# Weekly Note: 5 May - 11 June

## Delivered this week

+ This week mainly involved making the proposal publish-worthy. Discussed ideas for server side and client side validation of CustomResources.
+ Discussed implementation plan for validation.
+ Added JSON Schema examples, client side/server side validation sections and updated the proposal on the whole while collaborating with sttts.
+ Shared the proposal with deads2k, enisoc and xiao.zhou (and integrated ideas from her proposal).
+ Addressed David's comments and updated the proposal.
+ Looked for whether webhook admission will work for CRDs. This [PR](https://github.com/kubernetes/kubernetes/pull/46316) should make webhook admission work with CRs out of the box since admission plugins for apiextensions-apiserver are inherited from kube-apiserver.
 + Added types for validation (prototype) in this [PR](https://github.com/kubernetes/kubernetes/pull/47263).
    - However the types involved recursive types and the script for code generation builds a tree for automatically generating the types.
    - This tree was going into an infinite recursion  and was giving out a stackoverflow due to the recursive types. Basically, we needed a mechanism to make sure it does not undergo recursion while coming across a type it has already seen.
    - Decided to fix the default-gen and conversion-gen instead of modifying our types. 
    - Added [PR](https://github.com/kubernetes/gengo/pull/61) for defaulter fix - this goes in `kubernetes/gengo`. Currently this works for CRD types but breaks other stuff. Need to fix it and add tests for defaulters. thockin suggested to add it in the output tests [here](https://github.com/kubernetes/gengo/tree/master/examples/deepcopy-gen/output_tests).
    - Correction for conversion goes into `kubernetes/kubernetes` so it is included in the types PR.


## Pitfalls

Could improve spending less time on the proposal. The default-gen and conversion-gen still don't work well even though they seem to work for the CRD types. Need to fix them so that they work for native resources as well.

