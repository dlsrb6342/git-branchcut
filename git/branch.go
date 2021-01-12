package git

import (
	"fmt"
	"strings"

	"github.com/codeskyblue/go-sh"
)

// SwitchToDefault Switch to default branch
func SwitchToDefault(mainBranch string) {
	_, err := sh.Command(gitExecutable, "switch", mainBranch).Output()
	checkErr(err, "Default branch is invalid")
}

// Branches get all branches
func Branches() []string {
	output, err := sh.Command(gitExecutable, "branch").Output()
	checkErr(err, "Get Branch list failed")

	branches := strings.Split(string(output), "\n")
	for i := range branches {
		branches[i] = strings.TrimSpace(branches[i])
	}
	return branches
}

// BranchesByPattern get branches by pattern
func BranchesByPattern(pattern string) []string {
	output, err := sh.Command(gitExecutable, "branch", "--list", pattern).Output()
	checkErr(err, "Get Branch by pattern failed")

	branches := strings.Split(string(output), "\n")
	for i := range branches {
		branches[i] = strings.TrimSpace(branches[i])
	}
	return branches
}

// DeleteBranches Delete brances
func DeleteBranches(branches []interface{}, force bool, dryRun bool) {
	var args []interface{}
	args = append(args, "branch", "--delete")
	if force {
		args = append(args, "--force")
	}

	var branchNames []string
	for i := 0; i < len(branches); i++ {
		branch := strings.TrimSpace(fmt.Sprintf("%v", branches[i]))
		if len(branch) > 0 {
			branchNames = append(branchNames, branch)
		}
	}

	args = append(args, strings.Join(branchNames[:], " "))
	if dryRun {
		fmt.Println(gitExecutable, joinInterfaceArray(args, " "))
	} else {
		sh.Command(gitExecutable, args).Run()
	}

}

func joinInterfaceArray(a []interface{}, sep string) string {
	str := ""

	for index := 0; index < len(a); index++ {
		str1 := fmt.Sprintf("%v", a[index])
		str += sep + str1
	}
	return str
}
