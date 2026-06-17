package ec2

// CreateSnapshot is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Creates a snapshot of an EBS volume and stores it in Amazon S3. You can use
// snapshots for backups, to make copies of EBS volumes, and to save data before
// shutting down an instance.
//
// The location of the source EBS volume determines where you can create the
// snapshot.
//
// - If the source volume is in a Region, you must create the snapshot in the
// same Region as the volume.
//
// - If the source volume is in a Local Zone, you can create the snapshot in the
// same Local Zone or in its parent Amazon Web Services Region.
//
// - If the source volume is on an Outpost, you can create the snapshot on the
// same Outpost or in its parent Amazon Web Services Region.
//
// When a snapshot is created, any Amazon Web Services Marketplace product codes
// that are associated with the source volume are propagated to the snapshot.
//
// You can take a snapshot of an attached volume that is in use. However,
// snapshots only capture data that has been written to your Amazon EBS volume at
// the time the snapshot command is issued; this might exclude any data that has
// been cached by any applications or the operating system. If you can pause any
// file systems on the volume long enough to take a snapshot, your snapshot should
// be complete. However, if you cannot pause all file writes to the volume, you
// should unmount the volume from within the instance, issue the snapshot command,
// and then remount the volume to ensure a consistent and complete snapshot. You
// may remount and use your volume while the snapshot status is pending .
//
// When you create a snapshot for an EBS volume that serves as a root device, we
// recommend that you stop the instance before taking the snapshot.
//
// Snapshots that are taken from encrypted volumes are automatically encrypted.
// Volumes that are created from encrypted snapshots are also automatically
// encrypted. Your encrypted volumes and any associated snapshots always remain
// protected. For more information, see [Amazon EBS encryption]in the Amazon EBS User Guide.
//
// [Amazon EBS encryption]: https://docs.aws.amazon.com/ebs/latest/userguide/ebs-encryption.html
