package version

import (
	"runtime"
)

var (
	version   string
	goVersion string
	commit    string
	buildDate string
)

// VersionInfo holds versioning information
type versionInfo struct {
	Version         string
	GoVersion       string
	Commit          string
	BuildDate       string
	OperatingSystem string
	Architechture   string
}

// GetVersionInfo returns the version information of the application.
func GetVersionInfo() versionInfo {
	return versionInfo{
		Version:         version,
		GoVersion:       goVersion,
		Commit:          commit,
		BuildDate:       buildDate,
		OperatingSystem: runtime.GOOS,
		Architechture:   runtime.GOARCH,
	}
}
