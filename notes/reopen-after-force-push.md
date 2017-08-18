# How to re-open a PR if it's force pushed after it is closed?

You need the rights to reopen pull requests on the repository.

1. Write down the current commit hash of your PR-branch `git log --oneline -1 <PR-BRANCH>`
2. Write down the latest commit hash on github before the PR has been closed.
3. `git push -f origin <GITHUB-HASH-FROM-STEP-2>:<PR-BRANCH>`
4. Reopen the PR.
5. `git push -f origin <HASH-FROM-STEP-1>:<PR-BRANCH>`

#### Example
You've a PR branch `my-feature` currently at `1234567`. Looking at the the PRs page, we see that the PR was closed when `my-feature` pointed at `0abcdef`.
* `git push -f origin 0abcdef:my-feature` `#pushing the old commit the PR has been closed with`
* Reopen the PR.
* `git push -f origin 1234567:my-feature` `#pushing the latest commit`