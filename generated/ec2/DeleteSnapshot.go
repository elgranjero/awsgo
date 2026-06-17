package ec2

// DeleteSnapshot is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Deletes the specified snapshot.
//
// When you make periodic snapshots of a volume, the snapshots are incremental,
// and only the blocks on the device that have changed since your last snapshot are
// saved in the new snapshot. When you delete a snapshot, only the data not needed
// for any other snapshot is removed. So regardless of which prior snapshots have
// been deleted, all active snapshots will have access to all the information
// needed to restore the volume.
//
// You cannot delete a snapshot of the root device of an EBS volume used by a
// registered AMI. You must first deregister the AMI before you can delete the
// snapshot.
//
// For more information, see [Delete an Amazon EBS snapshot] in the Amazon EBS User Guide.
//
// [Delete an Amazon EBS snapshot]: https://docs.aws.amazon.com/ebs/latest/userguide/ebs-deleting-snapshot.html
