# rally-puller

`rally-puller` is a simple CLI utility written in Go that pulls defect details from Rally (CA Agile Central) and saves them locally as a formatted Markdown file (`task.md` by default).

## Main Use Case

The primary use case for `rally-puller` is to streamline development workflows by automatically fetching context about a task or defect when starting work. 

Specifically, it is designed to be added as a hook to [git-wt](https://github.com/k1LoW/git-wt) (a tool for managing git worktrees). By integrating `rally-puller` into a `git-wt` hook, a Markdown file containing the defect's description will be generated automatically whenever you check out or create a new worktree project.

If your git branches are named with the Rally defect ID (e.g., `feature/DE12345-fix-login`), `rally-puller` can automatically detect the ID from the branch name without any manual input!

## Installation

You can build the binary from the source:

```bash
go build -o rally-puller main.go
```

Then, move it to a location in your system's `$PATH` (e.g., `/usr/local/bin` or `~/bin`).

## Usage

```bash
rally-puller pull [flags]
```

### Flags

* `-d, --defect`: The Rally defect FormattedID (e.g., `DE12345`). If omitted, `rally-puller` attempts to resolve the defect ID automatically from the current git branch name.
* `-f, --filename`: The name of the markdown file to be written. Defaults to `task.md`.
* `-c, --cwd`: The working directory to execute the tool in.

### Configuration

Ensure you have your Rally credentials and API keys configured according to the tool's internal configuration loading (`internal/conf`).

## `git-wt` Hook Configuration

You can easily configure this utility to run as a `git-wt` hook. When `git-wt` sets up a new worktree, it will automatically execute `rally-puller` to pull the defect data into that directory.

Run the following command to add the hook to your git configuration:

```bash
git config wt.hook "rally-puller pull --cwd \$PWD"
```

*(Note: You can also use `git config --global wt.hook ...` if you want this to apply to all repositories globally).*

With this in place, every new worktree associated with a defect branch will have a handy `task.md` waiting for you, complete with the problem description directly from Rally.
