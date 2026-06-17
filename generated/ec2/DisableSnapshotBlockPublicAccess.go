package ec2

// DisableSnapshotBlockPublicAccess is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Disables the block public access for snapshots setting at the account level for
// the specified Amazon Web Services Region. After you disable block public access
// for snapshots in a Region, users can publicly share snapshots in that Region.
//
// Enabling block public access for snapshots in block-all-sharing mode does not
// change the permissions for snapshots that are already publicly shared. Instead,
// it prevents these snapshots from be publicly visible and publicly accessible.
// Therefore, the attributes for these snapshots still indicate that they are
// publicly shared, even though they are not publicly available.
//
// If you disable block public access , these snapshots will become publicly
// available again.
//
// For more information, see [Block public access for snapshots] in the Amazon EBS User Guide .
//
// [Block public access for snapshots]: https://docs.aws.amazon.com/ebs/latest/userguide/block-public-access-snapshots.html
