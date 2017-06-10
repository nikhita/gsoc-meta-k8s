# Weekly Note: 29 May - 4 June

## Delivered this week

+ Updated the [PR](https://github.com/kubernetes/kubernetes/pull/46200) for naming conflicts and fixed minor nits.
+ Added a test for checking if self link for cluster scoped resources work right.
+ Created a new [PR](https://github.com/kubernetes/kubernetes/pull/46585) to check for GC. However it is blocked on [#46000](https://github.com/kubernetes/kubernetes/pull/46000) which makes GC occasionally refresh its discovery information. Moved the test for 1.8 due to the blocker.
+ Created a [PR](https://github.com/kubernetes/kubernetes/pull/46624) for a test for advanced CRUD for CRDs. However, later it turned out that the PR was redundant because CRs don't have a status/spec split. Nevertheless, I learnt a lot of bash while writing this so it was awesome!
+ Started out with checking that Helm works right with CRDs. However the master branch was broken and the installation faced lots of problems. Created a [PR](https://github.com/kubernetes/helm/pull/2523) on Helm for that. However, Helm is just a use case and proved difficult to vendor and test so moved it to 1.8.
+ Discussed about an accessible client for CRDs - [issue](https://github.com/kubernetes/kubernetes/issues/46702).
+ Suggested fixes for the docs [PR](https://github.com/kubernetes/kubernetes.github.io/pull/3940) - we probably want to mention that GC won't work for CRDs.

## Pitfalls

Wasted a lot of time on a very silly mistake while extracting the group name from the CRD name. The PR for advanced CRUD was not really needed - however it turned out to be a good learning experience!

