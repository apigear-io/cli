# Git Start - Create Feature Branch

Create a new feature branch from the main branch following best practices.

## Instructions

1. Ask the user for a feature name/description if not provided as an argument
2. Generate a branch name using the format: `feature/<descriptive-name>` or `fix/<descriptive-name>`
   - Use kebab-case for the branch name
   - Keep it concise but descriptive
   - Suggest the branch name to the user for approval
3. Check the current git status to ensure working directory is clean
4. If there are uncommitted changes, ask the user what to do (commit, stash, or abort)
5. Switch to main branch and pull latest changes
6. Create and checkout the new feature branch
7. Confirm the new branch has been created and is active

## Example Usage

```
/git-start user-authentication
/git-start fix login bug
/git-start
```

## Branch Naming Convention

- `feature/<name>` - For new features
- `fix/<name>` - For bug fixes
- `refactor/<name>` - For refactoring
- `docs/<name>` - For documentation changes
- `test/<name>` - For test improvements

## Best Practices

- Always start from an updated main branch
- Use descriptive branch names that reflect the work
- Ensure working directory is clean before branching
