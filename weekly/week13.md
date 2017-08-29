# Weekly Note: 31 July - 6 August

## Delivered this week

+ Created the spreadsheet after lots of discussion about changing the JSONSchema struct: https://docs.google.com/spreadsheets/d/1Mkm9L7CXGvRorV0Cr4Vwfu0DH7XRi24YHPiDK1NZWo4/edit?usp=sharing
    - went through the openAPI spec
    - Additional discussion: https://github.com/kubernetes/kubernetes/issues/49879.
+ Discussed more about subresources (integration in kubectl, etc).
+ Created the WIP subresources proposal gdoc and iterated a LOT over it. Asked anthony for a review too.
+ Added feature gate for `CustomResourceValidation`.
    - fixed feature gate godoc: https://github.com/kubernetes/kubernetes/pull/50638.
+ There was a data race for CRD `storageMap`.
    - Fixed that: https://github.com/kubernetes/kubernetes/pull/50098
    - Cherry-picked it to 1.7: https://github.com/kubernetes/kubernetes/pull/50250
+ Fixed the fuzzer for JSONSchemaProps.