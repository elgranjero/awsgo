package ec2

// CopySnapshot is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Creates an exact copy of an Amazon EBS snapshot.
//
// The location of the source snapshot determines whether you can copy it or not,
// and the allowed destinations for the snapshot copy.
//
// - If the source snapshot is in a Region, you can copy it within that Region,
// to another Region, to an Outpost associated with that Region, or to a Local Zone
// in that Region.
//
// - If the source snapshot is in a Local Zone, you can copy it within that
// Local Zone, to another Local Zone in the same zone group, or to the parent
// Region of the Local Zone.
//
// - If the source snapshot is on an Outpost, you can't copy it.
//
// When copying snapshots to a Region, the encryption outcome for the snapshot
// copy depends on the Amazon EBS encryption by default setting for the destination
// Region, the encryption status of the source snapshot, and the encryption
// parameters you specify in the request. For more information, see [Encryption and snapshot copying].
//
// Snapshots copied to an Outpost must be encrypted. Unencrypted snapshots are not
// supported on Outposts. For more information, [Amazon EBS local snapshots on Outposts].
//
// Snapshots copies have an arbitrary source volume ID. Do not use this volume ID
// for any purpose.
//
// For more information, see [Copy an Amazon EBS snapshot] in the Amazon EBS User Guide.
//
// [Encryption and snapshot copying]: https://docs.aws.amazon.com/ebs/latest/userguide/ebs-copy-snapshot.html#creating-encrypted-snapshots
// [Copy an Amazon EBS snapshot]: https://docs.aws.amazon.com/ebs/latest/userguide/ebs-copy-snapshot.html
// [Amazon EBS local snapshots on Outposts]: https://docs.aws.amazon.com/ebs/latest/userguide/snapshots-outposts.html#considerations
