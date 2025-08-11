package version

import "fmt"

var (
	// These are set during build time via ldflags
	Version   = "v0.9.3-final"
	GitCommit = "4179c2745b4fb22fcb37f36de27b3daa39f114f0"
	BuiltAt   = "2024-07-14T09:48:32"
)

// GetVersion returns version information
func GetVersion() string {
	return fmt.Sprintf("Version: %s\nGit Commit: %s\nBuilt At: %s", Version, GitCommit, BuiltAt)
}