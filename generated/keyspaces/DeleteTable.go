package keyspaces

// DeleteTable is generated as a reference stub.
// Executable command wiring lives under cmd/keyspaces.go.
//
// The DeleteTable operation deletes a table and all of its data. After a
// DeleteTable request is received, the specified table is in the DELETING state
// until Amazon Keyspaces completes the deletion. If the table is in the ACTIVE
// state, you can delete it. If a table is either in the CREATING or UPDATING
// states, then Amazon Keyspaces returns a ResourceInUseException . If the
// specified table does not exist, Amazon Keyspaces returns a
// ResourceNotFoundException . If the table is already in the DELETING state, no
// error is returned.
