# Git command quick ref

## Working with tags Tags

```bash
# Create new tag
git tag -am "Fresh copy" v1.0.0
# Push tag to remote
git push origin v1.0.0
# Delete tag
git tag --delete v1.0.0
# Delete tag remote
git push --delete origin v1.0.0
# Create branch from tag
git checkout -b newbranch v1.0.0

```
