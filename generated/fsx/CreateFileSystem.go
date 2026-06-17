package fsx

// CreateFileSystem is generated as a reference stub.
// Executable command wiring lives under cmd/fsx.go.
//
// Creates a new, empty Amazon FSx file system. You can create the following
// supported Amazon FSx file systems using the CreateFileSystem API operation:
//
// - Amazon FSx for Lustre
//
// - Amazon FSx for NetApp ONTAP
//
// - Amazon FSx for OpenZFS
//
// - Amazon FSx for Windows File Server
//
// This operation requires a client request token in the request that Amazon FSx
// uses to ensure idempotent creation. This means that calling the operation
// multiple times with the same client request token has no effect. By using the
// idempotent operation, you can retry a CreateFileSystem operation without the
// risk of creating an extra file system. This approach can be useful when an
// initial call fails in a way that makes it unclear whether a file system was
// created. Examples are if a transport level timeout occurred, or your connection
// was reset. If you use the same client request token and the initial call created
// a file system, the client receives success as long as the parameters are the
// same.
//
// If a file system with the specified client request token exists and the
// parameters match, CreateFileSystem returns the description of the existing file
// system. If a file system with the specified client request token exists and the
// parameters don't match, this call returns IncompatibleParameterError . If a file
// system with the specified client request token doesn't exist, CreateFileSystem
// does the following:
//
// - Creates a new, empty Amazon FSx file system with an assigned ID, and an
// initial lifecycle state of CREATING .
//
// - Returns the description of the file system in JSON format.
//
// The CreateFileSystem call returns while the file system's lifecycle state is
// still CREATING . You can check the file-system creation status by calling the [DescribeFileSystems]
// operation, which returns the file system state along with other information.
//
// [DescribeFileSystems]: https://docs.aws.amazon.com/fsx/latest/APIReference/API_DescribeFileSystems.html
