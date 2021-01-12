package cmd

import (
	"fmt"
	"os"

	"github.com/dlsrb6342/git-branchcut/git"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "branchcut",
		Short: "A Git branch cleanup addon, branchcut",
		Long:  "branchcut is a git addon for deleting git branches with various options.",
	}
	executableGitPath string
	displayVersion    string
)

// Execute executes the root command.
func Execute(version string) {
	git.SetGitExecutable(executableGitPath)
	displayVersion = version
	rootCmd.SetHelpTemplate(fmt.Sprintf("Version: github.com/dlsr6342/%s\n%s\n",
		version, rootCmd.HelpTemplate()))

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&executableGitPath, "executable-git-path", "g", "git",
		"The git executable to use")
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}
