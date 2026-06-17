package ec2

// EnableSnapshotBlockPublicAccess is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Enables or modifies the block public access for snapshots setting at the
// account level for the specified Amazon Web Services Region. After you enable
// block public access for snapshots in a Region, users can no longer request
// public sharing for snapshots in that Region. Snapshots that are already publicly
// shared are either treated as private or they remain publicly shared, depending
// on the State that you specify.
//
// Enabling block public access for snapshots in block all sharing mode does not
// change the permissions for snapshots that are already publicly shared. Instead,
// it prevents these snapshots from be publicly visible and publicly accessible.
// Therefore, the attributes for these snapshots still indicate that they are
// publicly shared, even though they are not publicly available.
//
// If you later disable block public access or change the mode to block new
// sharing, these snapshots will become publicly available again.
//
// For more information, see [Block public access for snapshots] in the Amazon EBS User Guide.
//
// [Block public access for snapshots]: https://docs.aws.amazon.com/ebs/latest/userguide/block-public-access-snapshots.html
