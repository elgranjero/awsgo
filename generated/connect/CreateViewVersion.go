package connect

// CreateViewVersion is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// Publishes a new version of the view identifier.
//
// Versions are immutable and monotonically increasing.
//
// It returns the highest version if there is no change in content compared to
// that version. An error is displayed if the supplied ViewContentSha256 is
// different from the ViewContentSha256 of the $LATEST alias.
