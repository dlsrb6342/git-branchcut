package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dlsrb6342/git-branchcut/git"
	"github.com/juliangruber/go-intersect"
	"github.com/spf13/cobra"
)

var (
	executeCmd = &cobra.Command{
		Use:   "execute",
		Short: "Cut branches",
		Run:   execute,
	}
	pattern    string
	dayOffset  int
	dryRun     bool
	force      bool
	mainBranch string
)

func init() {
	executeCmd.Flags().StringVarP(&pattern, "pattern", "p", "",
		"Pattern for branch name")
	executeCmd.Flags().IntVarP(&dayOffset, "day-offset", "o", 0,
		"Day offset. Branches that is committed before day offset would be deleted")
	executeCmd.Flags().StringVarP(&mainBranch, "main-branch", "m", "main",
		"Default branch name")
	executeCmd.Flags().BoolVar(&dryRun, "dry-run", false,
		"Only print branch list that would be delete, without deleting it.")
	executeCmd.Flags().BoolVar(&force, "force", false,
		"Allow deleting the branch irrespective of its merged status.")

	rootCmd.AddCommand(executeCmd)
}

func execute(cmd *cobra.Command, args []string) {
	if false == validateFlags() {
		fmt.Println("You must enter either day-offset or pattern.")
		rootCmd.Usage()
		os.Exit(1)
	}

	git.SwitchToDefault(mainBranch)

	filteredByDate := filterByDate()
	filteredByPattern := filterByPattern()

	intersected := intersect.Simple(filteredByDate, filteredByPattern)
	git.DeleteBranches(intersected.([]interface{}), force, dryRun)
}

func validateFlags() bool {
	return len(pattern) > 0 || dayOffset > 0
}

func filterByDate() []string {
	branches := git.Branches()
	if dayOffset == 0 {
		return branches
	}

	var filteredBranch []string
	pivotDate := time.Now().AddDate(0, 0, dayOffset*-1)

	for _, branch := range git.Branches() {
		if strings.Contains(branch, mainBranch) || len(branch) == 0 {
			continue
		}
		hash := git.CommitHash(branch)
		commitTime := git.CommitTimestamp(hash)

		if commitTime.Before(pivotDate) {
			filteredBranch = append(filteredBranch, branch)
		}
	}

	return filteredBranch
}

func filterByPattern() []string {
	if len(pattern) == 0 {
		return git.Branches()
	}
	return git.BranchesByPattern(pattern)
}
