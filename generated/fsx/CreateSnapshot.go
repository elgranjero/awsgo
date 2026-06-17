package fsx

// CreateSnapshot is generated as a reference stub.
// Executable command wiring lives under cmd/fsx.go.
//
// Creates a snapshot of an existing Amazon FSx for OpenZFS volume. With
// snapshots, you can easily undo file changes and compare file versions by
// restoring the volume to a previous version.
//
// If a snapshot with the specified client request token exists, and the
// parameters match, this operation returns the description of the existing
// snapshot. If a snapshot with the specified client request token exists, and the
// parameters don't match, this operation returns IncompatibleParameterError . If a
// snapshot with the specified client request token doesn't exist, CreateSnapshot
// does the following:
//
// - Creates a new OpenZFS snapshot with an assigned ID, and an initial
// lifecycle state of CREATING .
//
// - Returns the description of the snapshot.
//
// By using the idempotent operation, you can retry a CreateSnapshot operation
// without the risk of creating an extra snapshot. This approach can be useful when
// an initial call fails in a way that makes it unclear whether a snapshot was
// created. If you use the same client request token and the initial call created a
// snapshot, the operation returns a successful result because all the parameters
// are the same.
//
// The CreateSnapshot operation returns while the snapshot's lifecycle state is
// still CREATING . You can check the snapshot creation status by calling the [DescribeSnapshots]
// operation, which returns the snapshot state along with other information.
//
// [DescribeSnapshots]: https://docs.aws.amazon.com/fsx/latest/APIReference/API_DescribeSnapshots.html
