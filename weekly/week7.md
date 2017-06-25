# Weekly Note: 19 May - 25 June

## Delivered this week

+ This week mainly involved trying to get a prototype for validation working. Please see: https://github.com/kubernetes/kubernetes/pull/47263.
+ Decided not to use conversion-gen for converting our types since it would have been a huge time sink to get the generator working - the code is in a vendor directory and would cause problems with makefiles.
+ Wrote the [conversion funcs](https://github.com/kubernetes/kubernetes/pull/47263/files#diff-54fdd67617b0f2ef1b5a09148f77303e) manually (drawing inspiration from the generator).
    - had a little trouble adhering to the same style for converting maps - the generator does not work on maps as of now. The problem occurred because we can't write to nil maps.
    - `JSONSchemaRef` conversion does not work because of non-exported fields in the original go-openapi code. Maybe a PR to go-openapi would be a fix. Need to discuss this further.
+ Performed the validation using the go-openapi library.
+ Integrated this validation in the strategy of customresources and in the `customresource_handler`.
+ Added integration tests for checking validation.
    - general validation
    - validation on update of CRs
    - checking error messages
+ Things still TODO:
    - fix conversion of `JSONSchemaRef`. Note: the conversion to internal works but conversion from internal to external types does not work yet.
    - Add test to check for drift between our types and OpenAPI types once the ref field works.
    - When CRD is updated, the CR should validate again i.e the handler is installed again. Have ideas for this but need to bounce off someone for this.
+ Also changed a few more things:
    - `Validation` is now a `*CustomResourceValidation` since it is optional and could be nil.
    - added the `ReferenceURL` field because that is necessary for the `$ref` field.
+ Addressed more reviews on the proposal.

## Pitfalls

The prototype works but still some more things are remaining to get the validation working _completely_.


