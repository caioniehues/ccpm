# Installation Guide

This guide covers the installation, configuration, and verification of Claude Code PM (CCPM) in your project.

## Quick Install

Use the automated installer to set up CCPM in your project directory.

### Unix/Linux/macOS

```bash
curl -sSL https://automaze.io/ccpm/install | bash
```

### Windows (PowerShell)

```powershell
iwr -useb https://automaze.io/ccpm/install | iex
```

## What Gets Installed

The installation script initializes the following structure in your project root:

1.  **`ccpm/` Directory**: Contains the core system files.
    *   `agents/`: Specialized AI agent definitions.
    *   `commands/`: Slash command prompt files.
    *   `context/`: Project context templates and system prompts.
    *   `hooks/`: Git hooks for workflow automation.
    *   `scripts/`: Utility scripts for project management.

2.  **Command Symlinks**: Makes commands discoverable by your AI assistant.
    *   Creates `.claude/commands/` (for Claude Code) and links to CCPM commands.
    *   Creates `.opencode/command/` (for OpenCode) and links to CCPM commands.

3.  **Project Directories**:
    *   `.claude/prds/`: Storage for Product Requirements Documents.
    *   `.claude/epics/`: Storage for Epic planning and task tracking files.

4.  **Git Configuration**:
    *   Updates `.gitignore` to ensure local planning files are not committed accidentally.

## Manual Installation

If you prefer to install manually or are contributing to CCPM:

1.  **Clone the Repository**:
    ```bash
    git clone https://github.com/automazeio/ccpm.git ccpm
    ```

2.  **Create Directory Structure**:
    ```bash
    mkdir -p .claude/prds .claude/epics .claude/commands
    mkdir -p .opencode/command
    ```

3.  **Link Commands**:
    Create symbolic links from `ccpm/commands/*` to `.claude/commands/` and `.opencode/command/`.

    *For Unix/Linux/macOS:*
    ```bash
    # Link Claude Code commands
    ln -sf ../../ccpm/commands/*.md .claude/commands/
    mkdir -p .claude/commands/pm
    ln -sf ../../../ccpm/commands/pm/*.md .claude/commands/pm/

    # Link OpenCode commands
    ln -sf ../../ccpm/commands/*.md .opencode/command/
    mkdir -p .opencode/command/pm
    ln -sf ../../../ccpm/commands/pm/*.md .opencode/command/pm/
    ```

4.  **Configure Git Ignore**:
    Add the following to your `.gitignore`:
    ```
    # CCPM
    .claude/prds/
    .claude/epics/
    .claude/settings.local.json
    ```

## GitHub CLI Setup

CCPM relies on the GitHub CLI (`gh`) for issue management and synchronization.

1.  **Install GitHub CLI**:
    *   **macOS**: `brew install gh`
    *   **Windows**: `winget install GitHub.cli`
    *   **Linux**: See [official instructions](https://github.com/cli/cli/blob/trunk/docs/install_linux.md)

2.  **Authenticate**:
    ```bash
    gh auth login
    ```
    Select `GitHub.com` and `HTTPS` (or `SSH` if configured). Ensure you authorize with `repo` and `read:org` scopes.

3.  **Install Extensions**:
    CCPM uses `gh-sub-issue` for managing parent-child relationships between epics and tasks.
    ```bash
    gh extension install yahsan2/gh-sub-issue
    ```

4.  **Create Labels**:
    The system uses specific labels to track work items.
    ```bash
    gh label create "epic" --color "0E8A16" --description "Epic issue containing multiple related tasks" --force
    gh label create "task" --color "1D76DB" --description "Individual task within an epic" --force
    ```

## Verification

After installation, verify everything is working correctly:

1.  **Restart your session** (Claude Code or OpenCode) to load the new commands.

2.  **Run the doctor command**:
    ```
    /doctor
    ```
    This will check for:
    *   Command availability
    *   Directory structure existence
    *   GitHub CLI authentication
    *   Required extensions

## Updating CCPM

To update to the latest version of CCPM:

### Automatic Update
Run the self-update command inside your session:
```
/self-update
```

### Manual Update
Re-run the installation script:
```bash
curl -sSL https://automaze.io/ccpm/install | bash
```
The script detects existing installations and asks if you want to update.

## Uninstalling

To remove CCPM from your project:

### Using Command
```
/uninstall
```

### Manual Removal
1.  Remove the `ccpm` directory:
    ```bash
    rm -rf ccpm
    ```
2.  Remove symlinks in `.claude/commands` and `.opencode/command`.
3.  (Optional) Remove project data (PRDs and Epics):
    ```bash
    rm -rf .claude/prds .claude/epics
    ```

## Troubleshooting

### Commands not appearing
If `/pm:...` commands don't appear in the slash command menu:
1.  **Restart your session**. Commands are loaded only at startup.
2.  Check that symlinks exist in `.claude/commands/` or `.opencode/command/`.

### "gh: command not found"
Ensure GitHub CLI is installed and in your system PATH. Run `gh --version` to verify.

### Permission Denied errors
If you see permission errors during installation or execution:
*   Ensure you have write permissions to the project directory.
*   **Do not** run the installer with `sudo` unless strictly necessary (it may mess up file ownership).

### GitHub Authentication issues
If sync fails:
1.  Run `gh auth status` to check your login state.
2.  If token is expired, run `gh auth refresh`.
3.  Ensure your token has `repo` scope access.
