package vembed

var gitCommit, gitBranch, gitState, gitSummary, buildDate, version string

// Version acts as a way to access the injected version information.
var Version Vembed

// Vembed represents the version of the package.
type Vembed struct {
	gitCommit  string
	gitBranch  string
	gitState   string
	gitSummary string
	buildDate  string
	version    string
}

// GetGitCommit return the git commit.
func (v *Vembed) GetGitCommit() string {
	return v.gitCommit
}

// GetGitBranch return the git branch
func (v *Vembed) GetGitBranch() string {
	return v.gitBranch
}

// GetGitState returns the git state.
func (v *Vembed) GetGitState() string {
	return v.gitState
}

// GetGitSummary returns the git summary.
func (v *Vembed) GetGitSummary() string {
	return v.gitSummary
}

// GetBuildDate returns the build date.
func (v *Vembed) GetBuildDate() string {
	return v.buildDate
}

// GetVersion returns the version.
func (v *Vembed) GetVersion() string {
	return v.version
}

// String returns the version.
func (v *Vembed) String() string {
	if v.version == "" {
		return "unknown"
	}
	return v.GetVersion()
}

func init() {
	Version.gitCommit = gitCommit
	Version.gitBranch = gitBranch
	Version.gitState = gitState
	Version.gitSummary = gitSummary
	Version.buildDate = buildDate
	Version.version = version
}
