# GSoC Improve TPRs Project Meta Tracker


[![In Progress](https://badge.waffle.io/nikhita/gsoc-meta-k8s.svg?label=In%20Progress&title=In%20Progress)](http://waffle.io/nikhita/gsoc-meta-k8s)


Issues created in this repository correspond to the tasks I'm working on. Please see the waffle board linked below for a better perspective of how the project is progressing.

## Student Info

+ Name: Nikhita Raghunath
+ Github: [nikhita](https://github.com/nikhita)
+ Email: nikitaraghunath@gmail.com
+ Slack nick: nikhita
+ Blog: https://nikhita.github.io/

## Links

+ [GSoC Proposal](notes/gsoc-proposal.pdf)
+ [Project on GSoC website](https://summerofcode.withgoogle.com/organizations/6540924424290304/#5982049109278720)
+ [TPR features issue](https://github.com/kubernetes/features/issues/95)
+ [Design Proposal: Validation for CustomResources](https://github.com/kubernetes/community/pull/708)
+ [Waffle board](https://waffle.io/nikhita/gsoc-meta-k8s)
+ [Daily update log](https://docs.google.com/document/d/1iCH03_jdyUsBfXKwgZtut7tXVbsl6lgusDt_lXLNNGo/edit?usp=sharing)
+ [cncf-soc github repo](https://github.com/cncf/soc)
+ [CNCF Blog post on GSoC](https://www.cncf.io/blog/2017/05/04/cncf-brings-kubernetes-coredns-opentracing-prometheus-google-summer-code-2017/)

## Weekly Summary

### Community Bonding Period - May

+ [Week 1](weekly/week1.md)
+ [Week 2](weekly/week2.md)
+ [Week 3](weekly/week3.md)

### Coding Period - June

+ [Week 4](weekly/week4.md)
+ [Week 5](weekly/week5.md)
+ [Week 6](weekly/week6.md)
+ [Week 7](weekly/week7.md)


## Blog posts

A RSS feed for blog posts only related to GSoC is available [here](https://nikhita.github.io/feed-gsoc.xml).

Note: I have decided not to write weekly updates as blog posts anymore but to include it as a quick weekly summary only. The blog posts will focus more on technical content.

+ [GSoC with Kubernetes - Week 1](https://nikhita.github.io/google-summer-of-code-kubernetes-week-1)
+ [GSoC with Kubernetes - Week 2 and 3](https://nikhita.github.io/google-summer-of-code-kubernetes-week2-3-community-bonding)


## Related Pull Requests/Issues

There are a few PRs to other Kubernetes projects as well like `kubernetes/community`, `kubernetes/examples`, `kubernetes/kubernetes.github.io`, etc. However, the following lists the PRs/Issues pertaining to my GSoC project only.

+ [#44026](https://github.com/kubernetes/kubernetes/pull/44026) - Preserve int data when unmarshalling for TPR
+ [#44612](https://github.com/kubernetes/kubernetes/pull/44612) - Fix kube-apiserver crash when patching TPR data
+ [#45721](https://github.com/kubernetes/kubernetes/pull/45721) - Add integration tests for kube-apiextensions-server
+ [#45793](https://github.com/kubernetes/kubernetes/pull/45793) - Add plural name for CustomResources example
+ [#46059](https://github.com/kubernetes/kubernetes/pull/46059) - Integration test for kube-apiextensions-server: integers
+ [#46200](https://github.com/kubernetes/kubernetes/pull/46200) - Add integration test for name conflicts
+ [#46585](https://github.com/kubernetes/kubernetes/pull/46585) - [WIP] Add integration test for GC (blocked on [#46000](https://github.com/kubernetes/kubernetes/pull/46000))
+ [#46624](https://github.com/kubernetes/kubernetes/pull/46624) - Add test for advanced CRUD for apiextensions
+ [#47263](https://github.com/kubernetes/kubernetes/pull/47263) - [WIP] apiextensions: validation of CustomResources
+ [#61](https://github.com/kubernetes/gengo/pull/61) - Fix defaulter codegen for recursive types
+ [#3](https://github.com/nikhita/kubernetes/pull/3) - Fix for code generation in kube
+ [#1](https://github.com/caesarxuchao/gengo/pull/1) - Add tests for recursive types in defaulter-gen
+ [#708](https://github.com/kubernetes/community/pull/708) - Proposal: Validation for CustomResources
+ [#47684](https://github.com/kubernetes/kubernetes/pull/47684) - Fix link to apiextensions client-go example
+ [#47748](https://github.com/kubernetes/kubernetes/pull/47748) - Update client-go to point to CR example
+ [#47743](https://github.com/kubernetes/kubernetes/issues/47743) - No examples for CR in client-go
+ [#48076](https://github.com/kubernetes/kubernetes/pull/48076) - Fix error in local-up-cluster
+ [#48114](https://github.com/kubernetes/kubernetes/pull/48114) - Automated cherry pick of #47748

