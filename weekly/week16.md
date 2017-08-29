# Weekly Note: 21 August - 29 August

## Delivered this week

+ Continued work on fuzzers.
    - this took a lot of time + a lot of debugging.
+ Submitted a talk for KubeCon.
+ Create a separate PR for round trip test error messages: https://github.com/kubernetes/kubernetes/pull/51204.
+ clayton pointed out that disabling protobuf is a breaking API change.
+ Discussed an idea/solution to implement proto for interface fields and got it working.
    - essentially, proto is only for external types so we want to avoid interface fields in the external types.
    - to do so, we can have a struct with a `Raw` field - basically, a stripped `runtime.Unknown`.
    - Use this JSON blob and convert to the internal types (which is still an interface)
    - added custom conversion functions + marshal functions + tests
+ Now the validation PR is ready! :rocket: