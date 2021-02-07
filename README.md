# git-branchcut
![go](https://img.shields.io/github/go-mod/go-version/dlsrb6342/git-branchcut)
[![Software License](https://img.shields.io/badge/License-MIT-orange.svg)](https://github.com/dlsrb6342/git-branchcut/blob/master/LICENSE)
[![version](https://img.shields.io/github/v/release/dlsrb6342/git-branchcut)](https://github.com/dlsrb6342/git-branchcut/releases/latest)


git-branchcut is a git add-on for cleaning up git outdated branches with various options.

### Why? 
I created this because there are too many branches stacked in my macbook. But there is no easy way to cleanup many outdated branches at once. Instead of deleting branches manually, this can be a solution.

## Install
### Homebrew
You can install git-branchcut via homebrew.

```
brew tap dlsrb6342/git-branchcut https://github.com/dlsrb6342/git-branchcut
brew install dlsrb6342/git-branchcut
```

### Download
Alternately, you can download executable file for your platform from [github release](https://github.com/dlsrb6342/git-branchcut/releases/latest).


## Usage
```
git branchcut help
git branchcut version
git branchcut execute [opts]
```

```
Cut branches

Usage:
  branchcut execute [flags]

Flags:
  -o, --day-offset int       Day offset. Branches that is committed before day offset would be deleted
      --dry-run              Only print branch list that would be delete, without deleting it.
      --force                Allow deleting the branch irrespective of its merged status.
  -h, --help                 help for execute
  -m, --main-branch string   Default branch name (default "main")
  -p, --pattern string       Pattern for branch name

Global Flags:
  -g, --executable-git-path string   The git executable to use (default "git")
```
You can cleanup branches with `day-offset` or `pattern`. If you set both two options, only branches that meet both conditions will be deleted.

* `day-offset` means that branches which is committed before the days of offset would be deleted.
  * For example, `foo` branch is commited 3 days ago and `bar` branch is commited 6 days ago. In this situation, if you set `day-offset 5`, only `bar` branch will be deleted.
* `pattern` means Pattern for deleted branch name.
  * For example, `feature/foo` branch is commited 3 days ago and `bugfix/bar` branch is commited 6 days ago. In this situation, if you set `pattern "feature/*"`, only `feature/foo` branch will be deleted.
