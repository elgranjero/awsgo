package connect

// UpdateViewContent is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// Updates the view content of the given view identifier in the specified Amazon
// Connect instance.
//
// It performs content validation if Status is set to SAVED and performs full
// content validation if Status is PUBLISHED . Note that the $SAVED alias' content
// will always be updated, but the $LATEST alias' content will only be updated if
// Status is PUBLISHED .
