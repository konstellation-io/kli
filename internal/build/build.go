package build

//nolint:gochecknoglobals
var (
	// Version is dynamically set by the toolchain or overridden by the Makefile.
	Version = "DEV"

	// Date is dynamically set at build time in the Makefile.
	Date = "" // YYYY-MM-DD
)
