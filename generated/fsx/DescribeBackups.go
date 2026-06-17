package fsx

// DescribeBackups is generated as a reference stub.
// Executable command wiring lives under cmd/fsx.go.
//
// Returns the description of a specific Amazon FSx backup, if a BackupIds value
// is provided for that backup. Otherwise, it returns all backups owned by your
// Amazon Web Services account in the Amazon Web Services Region of the endpoint
// that you're calling.
//
// When retrieving all backups, you can optionally specify the MaxResults
// parameter to limit the number of backups in a response. If more backups remain,
// Amazon FSx returns a NextToken value in the response. In this case, send a
// later request with the NextToken request parameter set to the value of the
// NextToken value from the last response.
//
// This operation is used in an iterative process to retrieve a list of your
// backups. DescribeBackups is called first without a NextToken value. Then the
// operation continues to be called with the NextToken parameter set to the value
// of the last NextToken value until a response has no NextToken value.
//
// When using this operation, keep the following in mind:
//
// - The operation might return fewer than the MaxResults value of backup
// descriptions while still including a NextToken value.
//
// - The order of the backups returned in the response of one DescribeBackups
// call and the order of the backups returned across the responses of a multi-call
// iteration is unspecified.
