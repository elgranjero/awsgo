package efs

// DeleteAccessPoint is generated as a reference stub.
// Executable command wiring lives under cmd/efs.go.
//
// Deletes the specified access point. After deletion is complete, new clients can
// no longer connect to the access points. Clients connected to the access point at
// the time of deletion will continue to function until they terminate their
// connection.
//
// This operation requires permissions for the elasticfilesystem:DeleteAccessPoint
// action.
