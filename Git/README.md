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

## Git delete all commit history
```bash
# Checkkout main branch and then create orphan branch
git checkout --orphan temp_branch

#Add all files and commit
git add -A ; git commit -m "Initial commit"

# Delete main branch and rename current branch
git branch -D main ; git branch -m main

# Push changes
git push --force origin main 

```