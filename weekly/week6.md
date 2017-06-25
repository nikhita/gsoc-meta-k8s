# Weekly Note: 12 May - 18 June

## Delivered this week

+ Fixed the defaulter-gen to support recursive types in this [PR](https://github.com/kubernetes/gengo/pull/61).
    - Earlier it was working for CRD types but breaking for other types.
    - The trick was to mark the type as "unseen" in the tree if it was no longer a nested struct.
    - [This](https://github.com/nikhita/kubernetes/pull/3) shows that there are no changes in kube after the  cumulative change in defaulter-gen and conversion-gen.
+ Added [tests](https://github.com/caesarxuchao/gengo/pull/1/files) for defaulter-gen on top of Chao's [PR](https://github.com/kubernetes/gengo/pull/55) as suggested by thockin.
+ Discussed and decided that we will need to implement our own custom marshaller for the `$ref` property.
+ SUBMITTED THE PROPOSAL!!! https://github.com/kubernetes/community/pull/708
    - Mentioned the proposal in the [#95](https://github.com/kubernetes/features/issues/95) issue and in [#38117](https://github.com/kubernetes/kubernetes/issues/38117).
    - Sent an [email](https://groups.google.com/forum/#!topic/kubernetes-sig-api-machinery/1vkhOSRKKtc) to SIG API Machinery about it.
+ Addressed review comments in the PR (against PRs in my fork first with sttts).
+ Had a problem with defaulting (we don't support it currently) but this [PR](https://github.com/go-openapi/validate/pull/27), when merged, should solve it.
+ Discussed code conversion for the types to OpenAPI types. We'll probably need recursive conversion funcs here.

## Pitfalls

The defaulter-gen fix that was intially creating problems could have been fixed sooner if I had figured out that one tiny little problem! Had originally said that defaulting would be supported but later after reading [[1](http://python-jsonschema.readthedocs.io/en/latest/faq/)][[2](https://groups.google.com/forum/#!topic/json-schema/xiTmft0Oatg)]found that go-openapi doesn't have defaulting support.
