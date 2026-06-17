package mgn

// RetryDataReplication is generated as a reference stub.
// Executable command wiring lives under cmd/mgn.go.
//
// Causes the data replication initiation sequence to begin immediately upon next
// Handshake for specified SourceServer IDs, regardless of when the previous
// initiation started. This command will not work if the SourceServer is not
// stalled or is in a DISCONNECTED or STOPPED state.
