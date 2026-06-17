package connect

// CreateView is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// Creates a new view with the possible status of SAVED or PUBLISHED .
//
// The views will have a unique name for each connect instance.
//
// It performs basic content validation if the status is SAVED or full content
// validation if the status is set to PUBLISHED . An error is returned if
// validation fails. It associates either the $SAVED qualifier or both of the
// $SAVED and $LATEST qualifiers with the provided view content based on the
// status. The view is idempotent if ClientToken is provided.
