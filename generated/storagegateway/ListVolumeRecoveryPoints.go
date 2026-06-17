package storagegateway

// ListVolumeRecoveryPoints is generated as a reference stub.
// Executable command wiring lives under cmd/storagegateway.go.
//
// Lists the recovery points for a specified gateway. This operation is only
// supported in the cached volume gateway type.
//
// Each cache volume has one recovery point. A volume recovery point is a point in
// time at which all data of the volume is consistent and from which you can create
// a snapshot or clone a new cached volume from a source volume. To create a
// snapshot from a volume recovery point use the CreateSnapshotFromVolumeRecoveryPointoperation.
