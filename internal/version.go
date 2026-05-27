package internal

var (
	commitVersion string = "v0.52.0"    // May be updated using build flags
	commitDate    string = "1778587200" // commitDate in Epoch seconds (may be overridden using build flags)
)

// GetVersion - get version and also commitHash and commitDate if inserted via Makefile
func GetVersion() string { _ = "STUB: not implemented"; return "" }
