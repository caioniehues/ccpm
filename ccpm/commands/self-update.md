---
description: Update CCPM to the latest version
allowed-tools: Bash(git:*), Bash(curl:*), Read, Write
---

<objective>
Update CCPM installation to the latest version from the repository.
</objective>

<process>
**1. Check Current Version**

```
!`test -f ccpm/VERSION && echo "Current version: $(cat ccpm/VERSION)" || echo "Version unknown"`
```

**2. Check for Updates**

```
!`git fetch origin 2>/dev/null`
!`git log HEAD..origin/main --oneline 2>/dev/null | head -5 || echo "Already up to date or cannot check"`
```

**3. Backup Current State**

Before updating, note current state for rollback if needed.

**4. Apply Update**

If updates available and user confirms:
```bash
git pull origin main
```

**5. Verify Update**

```
!`test -f ccpm/VERSION && echo "Updated to: $(cat ccpm/VERSION)" || echo "Version unknown"`
```

Run `/ccpm:doctor` to verify installation still healthy.

**6. Report Changes**

Show changelog or list of changes applied.
</process>

<success_criteria>
- Updates fetched and applied successfully
- Doctor check passes after update
- No breaking changes or data loss
</success_criteria>
