# Git Finish - Complete Feature Branch

Finalize the feature branch and prepare for merging or creating a pull request.

## Instructions

1. Run `git status` to ensure all changes are committed
2. If there are uncommitted changes, prompt user to commit them first (suggest using `/git-step`)
3. Run `git log main..HEAD` to show all commits on this branch
4. Ask the user what they want to do:
   - Create a pull request
   - Merge directly to main (if allowed by workflow)
   - Push branch without merging
   - Cancel
5. Based on user choice:

### Option A: Create Pull Request
1. Ensure branch is pushed to remote: `git push -u origin <branch-name>`
2. Analyze all commits to generate PR title and description
3. Use `gh pr create` to create the pull request with:
   - Title: Summarize the feature/fix
   - Body: Include conventional commit format with:
     - Summary section (bullet points of main changes)
     - Detailed description
     - Test plan (checklist of testing steps)
     - Related issues (if any)
4. Return the PR URL

### Option B: Merge to Main
1. Switch to main branch
2. Pull latest changes: `git pull origin main`
3. Merge feature branch: `git merge --no-ff <branch-name>`
4. Push to remote: `git push origin main`
5. Optionally delete feature branch locally and remotely
6. Confirm merge successful

### Option C: Push Only
1. Push branch to remote: `git push -u origin <branch-name>`
2. Provide instructions for creating PR manually
3. Confirm push successful

## Pull Request Template

```markdown
## Summary
- <bullet point of main change>
- <bullet point of another change>

## Description
<Detailed description of what was done and why>

## Test Plan
- [ ] Unit tests pass
- [ ] Integration tests pass
- [ ] Manual testing completed
- [ ] Documentation updated

## Related Issues
Closes #<issue-number>
```

## Pre-Merge Checklist

Before finishing, verify:
- [ ] All tests pass
- [ ] Code follows project conventions
- [ ] Documentation is updated
- [ ] No merge conflicts with main
- [ ] Commit messages follow conventional commits
- [ ] No sensitive data in commits

## Branch Cleanup

After successful merge, optionally:
- Delete local branch: `git branch -d <branch-name>`
- Delete remote branch: `git push origin --delete <branch-name>`
