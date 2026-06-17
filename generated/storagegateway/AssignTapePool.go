package storagegateway

// AssignTapePool is generated as a reference stub.
// Executable command wiring lives under cmd/storagegateway.go.
//
// Assigns a tape to a tape pool for archiving. The tape assigned to a pool is
// archived in the S3 storage class that is associated with the pool. When you use
// your backup application to eject the tape, the tape is archived directly into
// the S3 storage class (S3 Glacier or S3 Glacier Deep Archive) that corresponds to
// the pool.
