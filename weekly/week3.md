# Weekly Note: 22 May - 28 May

## Delivered this week

+ Created a PR [#46200](https://github.com/kubernetes/kubernetes/pull/46200) for NameConflicts. Came across many problems while doing this and fixed a few bugs!
    - This initially caused a panic ([stack trace](https://gist.github.com/nikhita/345c3629fc4a942b7c9ee5275dfeb74b)). This occurred because now we deleted from storage instead of the client (after this [PR](https://github.com/kubernetes/kubernetes/pull/46296)). This was racing the storage and causing a nil dereference. PR with fix: [#46430](https://github.com/kubernetes/kubernetes/pull/46430).
    - Did a rebase again but still did not work. Now I got:  `"reason": "InstanceDeletionFailed"` and `"message": "could not list instances: unable to find a custom resource client for noxus.mygroup.example.com"`. The problem turned out to be that the storage map was only filled if a CR (not the CRD!) was accessed via some CRUD request. So this worked only if you created CRs before the CRD deletion. PR with fix: [#46501](https://github.com/kubernetes/kubernetes/pull/46501). 
    - Fix naming controller so that when a resource is deleted, other resources in the same group are added to the queue. This is necessary since the deleted CRD can free names.
    - Added a `requeueAllOtherGroupCRDs` for the above. Had a little trouble making the synchronization was right. (Finally got it working!)
+ Wrote the schema for validation of Pod for the proposal.
+ Did more community bonding:
    - [Talked](https://kubernetes.slack.com/archives/C0EG7JC6T/p1495715788294534) to David about the project and CRDs.
    - Helped a few folks on Slack and answered questions about getting started with contributing: [1](https://kubernetes.slack.com/archives/C0EG7JC6T/p1495741622231635), [2](https://kubernetes.slack.com/archives/C0EG7JC6T/p1495963475724039), [3](https://kubernetes.slack.com/archives/C09R23FHP/p1495975931968264).
+ Did my first [code review](https://github.com/kubernetes/kubernetes/pull/46439)!
+ Some twitter fun:

<blockquote class="twitter-tweet" data-lang="en"><p lang="en" dir="ltr"><a href="https://twitter.com/TheNikhita">@TheNikhita</a> Thanks for your work so far <a href="https://twitter.com/TheNikhita">@TheNikhita</a>! You fixed a bug that we hit in Kubernetes and is getting backported to the next 1.6.x branch 👍👍👍</p>&mdash; Brandon Philips (@BrandonPhilips) <a href="https://twitter.com/BrandonPhilips/status/867820077063602176">May 25, 2017</a></blockquote>
   
## Pitfalls

Really need to ramp up on the proposal. Want to get in the tests before the code freeze on June 1 and will then continue to work on the proposal completely.