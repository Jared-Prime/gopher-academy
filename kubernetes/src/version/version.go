package version

// we'll supply these at build time via the Makefile
var (
	BuildTime = "unset"
	Commit    = "unset"
	Release   = "unset"
)
