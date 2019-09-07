package vembed

var gitCommit, gitBranch, gitState, gitSummary, buildDate, version string

var Version Vembed

type Vembed struct {
	gitCommit  string
	gitBranch  string
	gitState   string
	gitSummary string
	buildDate  string
	version    string
}

func (v *Vembed) GetGitCommit() string {
	return v.gitCommit
}
func (v *Vembed) GetGitBranch() string {
	return v.gitBranch
}
func (v *Vembed) GetGitState() string {
	return v.gitState
}
func (v *Vembed) GetGitSummary() string {
	return v.gitSummary
}
func (v *Vembed) GetBuildDate() string {
	return v.buildDate
}
func (v *Vembed) GetVersion() string {
	return v.version
}

func init() {
	Version.gitCommit = gitCommit
	Version.gitBranch = gitBranch
	Version.gitState = gitState
	Version.gitSummary = gitSummary
	Version.buildDate = buildDate
	Version.version = version
}
