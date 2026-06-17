package fsx

// DescribeSnapshots is generated as a reference stub.
// Executable command wiring lives under cmd/fsx.go.
//
// Returns the description of specific Amazon FSx for OpenZFS snapshots, if a
// SnapshotIds value is provided. Otherwise, this operation returns all snapshots
// owned by your Amazon Web Services account in the Amazon Web Services Region of
// the endpoint that you're calling.
//
// When retrieving all snapshots, you can optionally specify the MaxResults
// parameter to limit the number of snapshots in a response. If more backups
// remain, Amazon FSx returns a NextToken value in the response. In this case,
// send a later request with the NextToken request parameter set to the value of
// NextToken from the last response.
//
// Use this operation in an iterative process to retrieve a list of your
// snapshots. DescribeSnapshots is called first without a NextToken value. Then
// the operation continues to be called with the NextToken parameter set to the
// value of the last NextToken value until a response has no NextToken value.
//
// When using this operation, keep the following in mind:
//
// - The operation might return fewer than the MaxResults value of snapshot
// descriptions while still including a NextToken value.
//
// - The order of snapshots returned in the response of one DescribeSnapshots
// call and the order of backups returned across the responses of a multi-call
// iteration is unspecified.
