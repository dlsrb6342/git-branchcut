package git

import (
	"strings"
	"time"

	"github.com/codeskyblue/go-sh"
)

var (
	dateLayout = "2006-01-02 15:04:05 -0700"
)

// CommitTimestamp Get Timestamp of commit
func CommitTimestamp(hash string) time.Time {
	output, err := sh.Command(gitExecutable, "show", "-s", "--format=%ci", hash).Output()
	checkErr(err, "Get commit timestamp is failed")

	time, _ := time.Parse(dateLayout, strings.TrimSpace(string(output)))
	return time
}

// CommitHash Get Commit hash of branch
func CommitHash(branch string) string {
	output, err := sh.Command(gitExecutable, "rev-parse", branch).Output()
	checkErr(err, "Get Commit Hash of Branch failed")

	return strings.TrimSpace(string(output))
}
