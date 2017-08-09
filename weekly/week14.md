# Weekly Note: 7 August - 13 August

## Delivered this week

+ Updated the proposal. Some noteworthy points:
    - when we post the whole object, the status is ignored completely.
    - objectGeneration is incremented only when the spec status and NOT when the status changes.
    - discussed more about the structure of CustomResourceSubResourceType.
    - spec is a really special beast. :/
+ We'll probably have to implement our own Scale struct?
    - can this go to meta? this way all groups can access it. discussion here: https://github.com/kubernetes/kubernetes/issues/49504#issuecomment-317629927. Current discussion prefers autoscaling. But autoscaling is not the only way to "scale" resources. So waiting for the discussion to reach a conclusion...
    - but there are also PRs like this: https://github.com/kubernetes/kubernetes/pull/50204 which prefer the move to meta. :)
+ Talked to Jeff about the bazel bug. Now there is a fix for it! \o/
    - discussion: https://kubernetes.slack.com/messages/C5NV3MT97/p1502126389358226
    - fix: https://github.com/kubernetes/repo-infra/pull/32
+ the api-conventions doc says that the status sub-resource only updates the status path. However, as can be seen from the existing resources, it also updates metadata. Opened a PR: https://github.com/kubernetes/community/pull/882 to start the discussions.
    - we don't want to update metadata from the status subresource anymore (at least for the new resources).
    - we want to try to make the existing resources follow this pattern as well but that's a pretty breaking change.
    - one question remains: metadata.annotations is very helpful for controllers. With this change, this field will no longer be accessible. Is that ok or do we move with our old method (which is kind of ugly)?
+ Discussed at length about adding support for openAPI spec (v2 and v3) for the CRD Schema with Mehdy. Finally reached a conclusion: https://github.com/kubernetes/kubernetes/issues/49879. The Schema looks good now!
+ Refractored the validation of the schema with interfaces! :sparkles: