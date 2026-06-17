package fsx

// CreateFileSystemFromBackup is generated as a reference stub.
// Executable command wiring lives under cmd/fsx.go.
//
// Creates a new Amazon FSx for Lustre, Amazon FSx for Windows File Server, or
// Amazon FSx for OpenZFS file system from an existing Amazon FSx backup.
//
// If a file system with the specified client request token exists and the
// parameters match, this operation returns the description of the file system. If
// a file system with the specified client request token exists but the parameters
// don't match, this call returns IncompatibleParameterError . If a file system
// with the specified client request token doesn't exist, this operation does the
// following:
//
// - Creates a new Amazon FSx file system from backup with an assigned ID, and
// an initial lifecycle state of CREATING .
//
// - Returns the description of the file system.
//
// Parameters like the Active Directory, default share name, automatic backup, and
// backup settings default to the parameters of the file system that was backed up,
// unless overridden. You can explicitly supply other settings.
//
// By using the idempotent operation, you can retry a CreateFileSystemFromBackup
// call without the risk of creating an extra file system. This approach can be
// useful when an initial call fails in a way that makes it unclear whether a file
// system was created. Examples are if a transport level timeout occurred, or your
// connection was reset. If you use the same client request token and the initial
// call created a file system, the client receives a success message as long as the
// parameters are the same.
//
// The CreateFileSystemFromBackup call returns while the file system's lifecycle
// state is still CREATING . You can check the file-system creation status by
// calling the [DescribeFileSystems]operation, which returns the file system state along with other
// information.
//
// [DescribeFileSystems]: https://docs.aws.amazon.com/fsx/latest/APIReference/API_DescribeFileSystems.html
