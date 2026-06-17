package lightsail

// DeleteDiskSnapshot is generated as a reference stub.
// Executable command wiring lives under cmd/lightsail.go.
//
// Deletes the specified disk snapshot.
//
// When you make periodic snapshots of a disk, the snapshots are incremental, and
// only the blocks on the device that have changed since your last snapshot are
// saved in the new snapshot. When you delete a snapshot, only the data not needed
// for any other snapshot is removed. So regardless of which prior snapshots have
// been deleted, all active snapshots will have access to all the information
// needed to restore the disk.
//
// The delete disk snapshot operation supports tag-based access control via
// resource tags applied to the resource identified by disk snapshot name . For
// more information, see the [Amazon Lightsail Developer Guide].
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
