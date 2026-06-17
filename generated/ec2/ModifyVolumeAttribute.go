package ec2

// ModifyVolumeAttribute is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Modifies a volume attribute.
//
// By default, all I/O operations for the volume are suspended when the data on
// the volume is determined to be potentially inconsistent, to prevent
// undetectable, latent data corruption. The I/O access to the volume can be
// resumed by first enabling I/O access and then checking the data consistency on
// your volume.
//
// You can change the default behavior to resume I/O operations. We recommend that
// you change this only for boot volumes or for volumes that are stateless or
// disposable.
