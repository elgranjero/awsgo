package fsx

// CreateBackup is generated as a reference stub.
// Executable command wiring lives under cmd/fsx.go.
//
// Creates a backup of an existing Amazon FSx for Windows File Server file system,
// Amazon FSx for Lustre file system, Amazon FSx for NetApp ONTAP volume, or Amazon
// FSx for OpenZFS file system. We recommend creating regular backups so that you
// can restore a file system or volume from a backup if an issue arises with the
// original file system or volume.
//
// For Amazon FSx for Lustre file systems, you can create a backup only for file
// systems that have the following configuration:
//
// - A Persistent deployment type
//
// - Are not linked to a data repository
//
// For more information about backups, see the following:
//
// - For Amazon FSx for Lustre, see [Working with FSx for Lustre backups].
//
// - For Amazon FSx for Windows, see [Working with FSx for Windows backups].
//
// - For Amazon FSx for NetApp ONTAP, see [Working with FSx for NetApp ONTAP backups].
//
// - For Amazon FSx for OpenZFS, see [Working with FSx for OpenZFS backups].
//
// If a backup with the specified client request token exists and the parameters
// match, this operation returns the description of the existing backup. If a
// backup with the specified client request token exists and the parameters don't
// match, this operation returns IncompatibleParameterError . If a backup with the
// specified client request token doesn't exist, CreateBackup does the following:
//
// - Creates a new Amazon FSx backup with an assigned ID, and an initial
// lifecycle state of CREATING .
//
// - Returns the description of the backup.
//
// By using the idempotent operation, you can retry a CreateBackup operation
// without the risk of creating an extra backup. This approach can be useful when
// an initial call fails in a way that makes it unclear whether a backup was
// created. If you use the same client request token and the initial call created a
// backup, the operation returns a successful result because all the parameters are
// the same.
//
// The CreateBackup operation returns while the backup's lifecycle state is still
// CREATING . You can check the backup creation status by calling the [DescribeBackups] operation,
// which returns the backup state along with other information.
//
// [Working with FSx for OpenZFS backups]: https://docs.aws.amazon.com/fsx/latest/OpenZFSGuide/using-backups.html
// [Working with FSx for NetApp ONTAP backups]: https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/using-backups.html
// [Working with FSx for Lustre backups]: https://docs.aws.amazon.com/fsx/latest/LustreGuide/using-backups-fsx.html
// [Working with FSx for Windows backups]: https://docs.aws.amazon.com/fsx/latest/WindowsGuide/using-backups.html
// [DescribeBackups]: https://docs.aws.amazon.com/fsx/latest/APIReference/API_DescribeBackups.html
