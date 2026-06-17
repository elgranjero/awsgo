package connect

// DescribeView is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// Retrieves the view for the specified Amazon Connect instance and view
// identifier.
//
// The view identifier can be supplied as a ViewId or ARN.
//
// $SAVED needs to be supplied if a view is unpublished.
//
// The view identifier can contain an optional qualifier, for example, :$SAVED ,
// which is either an actual version number or an Amazon Connect managed qualifier
// $SAVED | $LATEST . If it is not supplied, then $LATEST is assumed for customer
// managed views and an error is returned if there is no published content
// available. Version 1 is assumed for Amazon Web Services managed views.
