package main

import (
	"fmt"

	"github.com/dlsrb6342/git-branchcut/cmd"
)

var (
	version = "v0.0.2"
	dirty   = ""
)

func main() {
	displayVersion := fmt.Sprintf("git-branchcut %s%s",
		version,
		dirty)
	cmd.Execute(displayVersion)
}
