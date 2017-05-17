# Weekly Note: 4 May - 14 May

## Delivered this week

- Created the meta repo with waffle integration. The issues will serve as tasks. Also added milestones to reflect "community bonding period", "first evaluation", etc.
- Came up with the following two go json schema libraries for implementing validation. Wrote benchmark tests for the same.
    - `go-jsonschema`
    - `go-openapi`
    - benchmarks can be found [here](https://github.com/nikhita/gsoc-meta-k8s/issues/1).
- Chose a library comparing their strength's and weaknesses. Decided to go with `go-openapi`.
- Came up with a few ideas and discussed general implementation outline for validation. Will submit a draft proposal next week.
- Went through the new TPR code and got comfortable with the codebase.
- Setup the new kube-apiextensions-server and wrote it in the form of a [script](https://github.com/nikhita/gsoc-meta-k8s/blob/master/notes/set-kube-apiextensions-server.sh) so that it can be easily exposed.
- Add the following integration tests for kube-apiextensions-server in the [PR](https://github.com/kubernetes/kubernetes/pull/45721):
    - [x] test namespace scoped resources
    - [x] test cluster scoped resources
    - [x] test discovery

## Pitfalls

I believe I can improve even further by writing more code. The first couple of days were spent in setting up the repository and getting started. This inertia could have been avoided. I will keep this in mind for the coming weeks!
