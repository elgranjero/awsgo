package storagegateway

// UpdateSnapshotSchedule is generated as a reference stub.
// Executable command wiring lives under cmd/storagegateway.go.
//
// Updates a snapshot schedule configured for a gateway volume. This operation is
// only supported in the cached volume and stored volume gateway types.
//
// The default snapshot schedule for volume is once every 24 hours, starting at
// the creation time of the volume. You can use this API to change the snapshot
// schedule configured for the volume.
//
// In the request you must identify the gateway volume whose snapshot schedule you
// want to update, and the schedule information, including when you want the
// snapshot to begin on a day and the frequency (in hours) of snapshots.
