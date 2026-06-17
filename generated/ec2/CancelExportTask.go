package ec2

// CancelExportTask is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Cancels an active export task. The request removes all artifacts of the export,
// including any partially-created Amazon S3 objects. If the export task is
// complete or is in the process of transferring the final disk image, the command
// fails and returns an error.
