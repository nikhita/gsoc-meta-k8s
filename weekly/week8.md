# Weekly Note: 26 June - 2 July

## Delivered this week

+ Fixed an error in `local-up-cluster.sh`. https://github.com/kubernetes/kubernetes/pull/48076
+ The example in TPR client-go was outdated for release 1.7. Fixed it: https://github.com/kubernetes/kubernetes/pull/47748.
+ Created a cherry pick to port this to release-1.7 branch.
    - came across a few minor errors in the cherry pic doc. Fixed them: https://github.com/kubernetes/community/pull/766.
    - the cherry pick PR: https://github.com/kubernetes/kubernetes/pull/48114.
    - note: for cherry-picking specific PRs, use the indivisual cherry-pick script.
+ Decided to use `GetURL` to go through a JSON string for the conversion to avoid accessing the internal fields.
+ The defaulting for go-openapi is on track.
+ The godoc for apiextensions was not generated because:
    - a typo: https://github.com/kubernetes/kubernetes/commit/9ee991837c2ba22becee1fb324deb81de3327903
    - the bot had not synced since the last 5 days.
+ Addressed reviews on the proposal. The proposal now looks good to get merged! :tada:




