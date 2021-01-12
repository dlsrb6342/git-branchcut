package git

import (
	"fmt"
	"os"
)

var (
	gitExecutable = "git"
)

// SetGitExecutable Set executable git path.
func SetGitExecutable(gitExecutablePath string) {
	gitExecutable = gitExecutablePath
}

func checkErr(err error, msg string) {
	if err != nil {
		fmt.Printf("Branch Command Failed: %s\n", msg)
		os.Exit(1)
	}

}
