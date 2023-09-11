# Github Notes

## Workflow
- Protect main branch
- Define pull request requirement
  - Run Code quality checks
  - Unit tests
  - Discussions and ammendments
  - Ensure origin pulled
- Protect Pull requst with requiring approval
- Define who can merge after approval
- Create templated for issues and pull requests

# Branch Protection
- Require Pull before merge to master
- Require approval
- Review from Code owners
- Require status checks to pass
- Require conversation resolution
- Bypass protection

# Merge strategies
- Fastforward : Commit history is single linear comment. So no merge commit. Git rebase is fast forward.
- Recursive : merges all commits.

# Event triggers
- Github events
    Each event can have types.  For eg: pull_request to trigger when type assign, unassign
  ```yaml
  on:
    pull_request:
      type: [assigned unassigned]
  ```
- Scheduled events
- Webhook event


https://github-actions-hero.vercel.app/
https://developers.giphy.com/
