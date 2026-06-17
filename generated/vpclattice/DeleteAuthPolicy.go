package vpclattice

// DeleteAuthPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/vpclattice.go.
//
// Deletes the specified auth policy. If an auth is set to AWS_IAM and the auth
// policy is deleted, all requests are denied. If you are trying to remove the auth
// policy completely, you must set the auth type to NONE . If auth is enabled on
// the resource, but no auth policy is set, all requests are denied.
