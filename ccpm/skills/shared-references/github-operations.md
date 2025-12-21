<overview>
Standard patterns for GitHub CLI operations. Includes critical repository protection to prevent accidental modifications to the CCPM template repository.
</overview>

<repository_protection>
**CRITICAL: Before ANY GitHub write operation, check if targeting the template repository:**

```bash
remote_url=$(git remote get-url origin 2>/dev/null || echo "")
if [[ "$remote_url" == *"automazeio/ccpm"* ]] || [[ "$remote_url" == *"automazeio/ccpm.git"* ]]; then
  echo "ERROR: You're trying to sync with the CCPM template repository!"
  echo ""
  echo "This repository (automazeio/ccpm) is a template for others to use."
  echo "You should NOT create issues or PRs here."
  echo ""
  echo "To fix: Update your remote origin:"
  echo "  git remote set-url origin https://github.com/YOUR_USERNAME/YOUR_REPO.git"
  exit 1
fi
```

This check MUST be performed before:
- `gh issue create` / `gh issue edit` / `gh issue comment`
- `gh pr create` / `gh pr edit`
- Any operation that modifies the GitHub repository
</repository_protection>

<authentication>
Don't pre-check authentication. Just run the command and handle failure:

```bash
gh {command} || echo "GitHub CLI failed. Run: gh auth login"
```
</authentication>

<common_operations>
**Get issue details:**
```bash
gh issue view {number} --json state,title,labels,body
```

**Create issue:**
```bash
remote_url=$(git remote get-url origin 2>/dev/null || echo "")
REPO=$(echo "$remote_url" | sed 's|.*github.com[:/]||' | sed 's|\.git$||')
[ -z "$REPO" ] && REPO="user/repo"
gh issue create --repo "$REPO" --title "{title}" --body-file {file} --label "{labels}"
```

**Update issue:**
```bash
# ALWAYS check remote origin first!
gh issue edit {number} --add-label "{label}" --add-assignee @me
```

**Add comment:**
```bash
# ALWAYS check remote origin first!
gh issue comment {number} --body-file {file}
```
</common_operations>

<error_handling>
If any gh command fails:

1. Show clear error: `"GitHub operation failed: {command}"`
2. Suggest fix: `"Run: gh auth login"` or check issue number
3. Don't retry automatically
</error_handling>

<best_practices>
- **ALWAYS** check remote origin before ANY write operation
- Trust that gh CLI is installed and authenticated
- Use `--json` for structured output when parsing
- Keep operations atomic - one gh command per action
- Don't check rate limits preemptively
</best_practices>
