package version

import "fmt"

var (
	gitTag       string
	version      string
	buildDate    string
	gitCommit    string
	gitTreeState string
)

func Version() {
	fmt.Println(fmt.Sprintf("GitTag=%s \t\n"+
		"Version=%s \t\n"+
		"BuildDate=%s \t\n"+
		"GitCommit=%s \t\n"+
		"GitTreeState=%s \t\n", gitTag, version, buildDate, gitCommit, gitTreeState))
}
