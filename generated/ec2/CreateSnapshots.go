package ec2

// CreateSnapshots is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Creates crash-consistent snapshots of multiple EBS volumes attached to an
// Amazon EC2 instance. Volumes are chosen by specifying an instance. Each volume
// attached to the specified instance will produce one snapshot that is
// crash-consistent across the instance. You can include all of the volumes
// currently attached to the instance, or you can exclude the root volume or
// specific data (non-root) volumes from the multi-volume snapshot set.
//
// The location of the source instance determines where you can create the
// snapshots.
//
// - If the source instance is in a Region, you must create the snapshots in the
// same Region as the instance.
//
// - If the source instance is in a Local Zone, you can create the snapshots in
// the same Local Zone or in its parent Amazon Web Services Region.
//
// - If the source instance is on an Outpost, you can create the snapshots on
// the same Outpost or in its parent Amazon Web Services Region.
