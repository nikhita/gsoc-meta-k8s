# Weekly Note: 24 July - 30 July

## Delivered this week

+ `TestCustomResourceValidationErrors` was now constantly failing.
    - after a LOT of debugging, turns out there was a problem upstream.
    - after even more debugging, turns out we could fix it on our own. :smile:
+ Added validations for validating that the _schema_ is correct (not only the custom resource).
+ Fixed some nits.
+ Updated the PR to get all tests *green*! :tada:
+ Discussed about defaulting the `$schema` field. Changing the defaults would break the api. So, defaulting seemed like a gray area.
+ Started with the sub-resources proposal + a LOT of research on existing resources.
+ Added a new PR for fixing conversion. Improved some renaming from the gengo PR: https://github.com/kubernetes/kubernetes/pull/49747.
+ The api-reference doc was not updated for the apiextensions. Turned to be a deeper issue than that. Created https://github.com/kubernetes/kubernetes/issues/49811.
+ Since the GC fix was in place with tests, closed down the earlier GC intergration test.