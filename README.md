# GSoC Improve TPRs Project Meta Tracker

[![In Progress](https://badge.waffle.io/nikhita/gsoc-meta-k8s.svg?label=In%20Progress&title=In%20Progress)](http://waffle.io/nikhita/gsoc-meta-k8s)

Issues created in this repository correspond to the tasks I'm working on. Please see the waffle board linked below for a better perspective of how the project is progressing.

## Project Abstract

Over time, the project goal has somewhat changed but this is the original abstract:

> ThirdPartyResources are already available but the implementation has languished with multiple outstanding capabilities that are missing. They did not complete the list of requirements for graduating to beta.

> Hence, there are multiple problems present in the current implementation of ThirdPartyResources. This project aims at working towards a number of known shortcomings to drive the ongoing effort toward a stable TPR release forward.

## General Info

+ Name: Nikhita Raghunath
+ Github: [nikhita](https://github.com/nikhita)
+ Email: hello@nikhita.dev
+ Slack nick: nikhita
+ Blog: https://www.nikhita.dev
+ Mentor: [Dr. Stefan Schimanski](https://github.com/sttts)

## Links

+ [GSoC Proposal](notes/gsoc-proposal.pdf)
+ [Project on GSoC website](https://summerofcode.withgoogle.com/organizations/6540924424290304/#5982049109278720)
+ [TPR/CRD features issue](https://github.com/kubernetes/features/issues/95)
+ [Waffle board](https://waffle.io/nikhita/gsoc-meta-k8s)
+ [Daily update log](https://docs.google.com/document/d/1iCH03_jdyUsBfXKwgZtut7tXVbsl6lgusDt_lXLNNGo/edit?usp=sharing)
+ [cncf-soc github repo](https://github.com/cncf/soc)
+ [Initial CNCF blogpost on GSoC](https://www.cncf.io/blog/2017/05/04/cncf-brings-kubernetes-coredns-opentracing-prometheus-google-summer-code-2017/)
+ [Final blogpost on the CNCF website](https://www.cncf.io/blog/2017/09/06/gsoc-17-making-customresources-kubernetes-awesome/)
+ [Mirror of above blogpost on my website](https://nikhita.github.io/gsoc-kubernetes-making-customresources-more-awesome)

## Design Proposals

+ [Design Proposal: Validation for CustomResources](https://github.com/kubernetes/community/pull/708)
+ [Design Proposal: SubResources for CustomResources](https://github.com/kubernetes/community/pull/913)

## Talks

[How to Contribute to Kubernetes](https://kccncna17.sched.com/event/CU74/how-to-contribute-to-kubernetes-b-nikhita-raghunath-student) at KubeCon+CNCFCon

## Weekly Summary

### Community Bonding Period - May

+ [Week 1](weekly/week1.md)
+ [Week 2](weekly/week2.md)
+ [Week 3](weekly/week3.md)
+ [Week 4](weekly/week4.md)

### Coding Period - June

+ [Week 5](weekly/week5.md)
+ [Week 6](weekly/week6.md)
+ [Week 7](weekly/week7.md)
+ [Week 8](weekly/week8.md)

### Coding Period - July

+ [Week 9](weekly/week9.md)
+ [Week 10](weekly/week10.md)
+ [Week 11](weekly/week11.md)
+ [Week 12](weekly/week12.md)

### Coding Period - August

+ [Week 13](weekly/week13.md)
+ [Week 14](weekly/week14.md)
+ [Week 15](weekly/week15.md)
+ [Week 16](weekly/week16.md)

## Pull Requests and Issues

The following list is automatically generated using https://github.com/nikhita/github-contrib.

**TOTAL : 69**

**Repository: kubernetes**

Total Pull Requests Created: 35
1. [kubernetes/kubernetes#53558](https://github.com/kubernetes/kubernetes/pull/53558) - Fix error for strategic merge patch of custom resources
2. [kubernetes/kubernetes#53426](https://github.com/kubernetes/kubernetes/pull/53426) - Automated cherry pick of #48036
3. [kubernetes/kubernetes#53312](https://github.com/kubernetes/kubernetes/pull/53312) - Rename TPR to CRD
4. [kubernetes/kubernetes#52793](https://github.com/kubernetes/kubernetes/pull/52793) - apiextensions: add round trip tests for CRD schema conversion
5. [kubernetes/kubernetes#52281](https://github.com/kubernetes/kubernetes/pull/52281) - apiextensions: fix conversion of CRD schema
6. [kubernetes/kubernetes#52034](https://github.com/kubernetes/kubernetes/pull/52034) - apiextensions: add defaulting for customresource validation
7. [kubernetes/kubernetes#51712](https://github.com/kubernetes/kubernetes/pull/51712) - apiextensions: add maximum for validation
8. [kubernetes/kubernetes#51204](https://github.com/kubernetes/kubernetes/pull/51204) - roundtrip: fix error messages
9. [kubernetes/kubernetes#50964](https://github.com/kubernetes/kubernetes/pull/50964) - Update Registry interface for deployment and endpoints
10. [kubernetes/kubernetes#50764](https://github.com/kubernetes/kubernetes/pull/50764) - apiextensions: update CRD strategy
11. [kubernetes/kubernetes#50638](https://github.com/kubernetes/kubernetes/pull/50638) - FeatureGate: update comments
12. [kubernetes/kubernetes#50444](https://github.com/kubernetes/kubernetes/pull/50444) - jsonpath: fix comments
13. [kubernetes/kubernetes#50250](https://github.com/kubernetes/kubernetes/pull/50250) - Automated cherry pick of #50098
14. [kubernetes/kubernetes#50098](https://github.com/kubernetes/kubernetes/pull/50098) - apiextensions: fix data race in storage
15. [kubernetes/kubernetes#50085](https://github.com/kubernetes/kubernetes/pull/50085) - fix kube-openapi imports
16. [kubernetes/kubernetes#49747](https://github.com/kubernetes/kubernetes/pull/49747) - conversion-gen: support recursive types
17. [kubernetes/kubernetes#49307](https://github.com/kubernetes/kubernetes/pull/49307) - fuzzer: remove unreachable code
18. [kubernetes/kubernetes#48630](https://github.com/kubernetes/kubernetes/pull/48630) - update vendored gengo
19. [kubernetes/kubernetes#48389](https://github.com/kubernetes/kubernetes/pull/48389) - apiextensions: add cleanup section to client-go
20. [kubernetes/kubernetes#48114](https://github.com/kubernetes/kubernetes/pull/48114) - Automated cherry pick of #47748
21. [kubernetes/kubernetes#48076](https://github.com/kubernetes/kubernetes/pull/48076) - Fix error in local-cluster-up
22. [kubernetes/kubernetes#47748](https://github.com/kubernetes/kubernetes/pull/47748) - Update custom-resources example in client-go
23. [kubernetes/kubernetes#47684](https://github.com/kubernetes/kubernetes/pull/47684) - Fix link to apiextensions client-go example
24. [kubernetes/kubernetes#47263](https://github.com/kubernetes/kubernetes/pull/47263) - apiextensions: validation for customresources
25. [kubernetes/kubernetes#46624](https://github.com/kubernetes/kubernetes/pull/46624) - Add test for advanced CRUD for apiextensions
26. [kubernetes/kubernetes#46585](https://github.com/kubernetes/kubernetes/pull/46585) - [WIP] apiextensions: add integration test for GC
27. [kubernetes/kubernetes#46200](https://github.com/kubernetes/kubernetes/pull/46200) - apiextensions: add integration test for name conflicts
28. [kubernetes/kubernetes#46059](https://github.com/kubernetes/kubernetes/pull/46059) - Integration test for kube-apiextensions-server: integers
29. [kubernetes/kubernetes#45793](https://github.com/kubernetes/kubernetes/pull/45793) - Add plural name for CustomResources example
30. [kubernetes/kubernetes#45721](https://github.com/kubernetes/kubernetes/pull/45721) - Add integration tests for kube-apiextensions-server
31. [kubernetes/kubernetes#44612](https://github.com/kubernetes/kubernetes/pull/44612) - Fix kube-apiserver crash when patching TPR data
32. [kubernetes/kubernetes#44026](https://github.com/kubernetes/kubernetes/pull/44026) - Preserve int data when unmarshalling for TPR
33. [kubernetes/kubernetes#43606](https://github.com/kubernetes/kubernetes/pull/43606) - Improve timeout error message for kubectl delete
34. [kubernetes/kubernetes#43591](https://github.com/kubernetes/kubernetes/pull/43591) - Update `kubectl help` descriptions and examples
35. [kubernetes/kubernetes#43573](https://github.com/kubernetes/kubernetes/pull/43573) - fixed formatting for examples README.md

Total Issues Opened: 4
1. [kubernetes/kubernetes#49811](https://github.com/kubernetes/kubernetes/issues/49811) - api-reference docs don't contain types from staging repos
2. [kubernetes/kubernetes#48675](https://github.com/kubernetes/kubernetes/issues/48675) - Auto-generation by deepcopy is broken
3. [kubernetes/kubernetes#48593](https://github.com/kubernetes/kubernetes/issues/48593) - hack/godep-save.sh pulls down docker/docker/project dir
4. [kubernetes/kubernetes#47743](https://github.com/kubernetes/kubernetes/issues/47743) - No examples for CR in client-go

Total Pull Requests Reviewed: 5
1. [kubernetes/kubernetes#53308](https://github.com/kubernetes/kubernetes/pull/53308) - apiextensions/examples: remove unnecessary function
2. [kubernetes/kubernetes#53051](https://github.com/kubernetes/kubernetes/pull/53051) - fix todo
3. [kubernetes/kubernetes#52753](https://github.com/kubernetes/kubernetes/pull/52753) - sample-controller: add example CRD controller
4. [kubernetes/kubernetes#47223](https://github.com/kubernetes/kubernetes/pull/47223) - Fix cross-repo link
5. [kubernetes/kubernetes#46439](https://github.com/kubernetes/kubernetes/pull/46439) - Test finalization for CRs


**Repository: helm**

Total Pull Requests Created: 1
1. [kubernetes/helm#2523](https://github.com/kubernetes/helm/pull/2523) - fix(helm): Fix log import


**Repository: kubernetes.github.io**

Total Pull Requests Created: 5
1. [kubernetes/kubernetes.github.io#5708](https://github.com/kubernetes/kubernetes.github.io/pull/5708) - Fix link to Design Overview
2. [kubernetes/kubernetes.github.io#5619](https://github.com/kubernetes/kubernetes.github.io/pull/5619) - Fix link after design proposals move
3. [kubernetes/kubernetes.github.io#5290](https://github.com/kubernetes/kubernetes.github.io/pull/5290) - Add docs for CustomResource validation
4. [kubernetes/kubernetes.github.io#4986](https://github.com/kubernetes/kubernetes.github.io/pull/4986) - JSONPath: rename title
5. [kubernetes/kubernetes.github.io#2791](https://github.com/kubernetes/kubernetes.github.io/pull/2791) - Fixed typo


**Repository: test-infra**

Total Issues Opened: 1
1. [kubernetes/test-infra#2787](https://github.com/kubernetes/test-infra/issues/2787) - Prow issue: Keep hitting a flake


**Repository: features**

Total Pull Requests Created: 1
1. [kubernetes/features#455](https://github.com/kubernetes/features/pull/455) - Fix links in 1.8 release notes

Total Pull Requests Reviewed: 1
1. [kubernetes/features#420](https://github.com/kubernetes/features/pull/420) - 1.8: Add relnotes for TPR and other deprecation removals.


**Repository: community**

Total Pull Requests Created: 8
1. [kubernetes/community#1128](https://github.com/kubernetes/community/pull/1128) - Fix links after design proposals move
2. [kubernetes/community#913](https://github.com/kubernetes/community/pull/913) - Proposal: SubResources for CustomResources
3. [kubernetes/community#882](https://github.com/kubernetes/community/pull/882) - api-conventions: status subresource updates metadata
4. [kubernetes/community#779](https://github.com/kubernetes/community/pull/779) - Fix link to point to English docs instead of Chinese
5. [kubernetes/community#766](https://github.com/kubernetes/community/pull/766) - Fix minor typos in cherry-picks doc
6. [kubernetes/community#708](https://github.com/kubernetes/community/pull/708) - Proposal: Validation for CustomResources.
7. [kubernetes/community#621](https://github.com/kubernetes/community/pull/621) - Fix link for factory methods of informers
8. [kubernetes/community#448](https://github.com/kubernetes/community/pull/448) - Fix link to contributing instructions


**Repository: frakti**

Total Pull Requests Created: 1
1. [kubernetes/frakti#238](https://github.com/kubernetes/frakti/pull/238) - Fix link after design proposal move


**Repository: client-go**

Total Pull Requests Created: 2
1. [kubernetes/client-go#274](https://github.com/kubernetes/client-go/pull/274) - README: remove non-existent update-staging-client-go.sh
2. [kubernetes/client-go#212](https://github.com/kubernetes/client-go/pull/212) - README: fix script to update staging area


**Repository: gengo**

Total Pull Requests Created: 1
1. [kubernetes/gengo#61](https://github.com/kubernetes/gengo/pull/61) - Fix defaulter-gen for recursive types


**Repository: ingress-nginx**

Total Pull Requests Created: 1
1. [kubernetes/ingress-nginx#1418](https://github.com/kubernetes/ingress-nginx/pull/1418) - Fix links after design proposals move


**Repository: bootkube**

Total Pull Requests Created: 1
1. [kubernetes-incubator/bootkube#727](https://github.com/kubernetes-incubator/bootkube/pull/727) - README: fix link to design proposal


**Repository: cri-o**

Total Pull Requests Created: 1
1. [kubernetes-incubator/cri-o#970](https://github.com/kubernetes-incubator/cri-o/pull/970) - tutorial: fix link after design proposal move


**Repository: cri-containerd**

Total Pull Requests Created: 1
1. [kubernetes-incubator/cri-containerd#318](https://github.com/kubernetes-incubator/cri-containerd/pull/318) - Fix link after design proposals move
