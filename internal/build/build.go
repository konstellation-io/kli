package build

//nolint:gochecknoglobals
var (
	// Version is dynamically set by the toolchain or overridden at build time.
	Version = "DEV"

	// Date is dynamically set at build time.
	Date = "" // YYYY-MM-DD
)
