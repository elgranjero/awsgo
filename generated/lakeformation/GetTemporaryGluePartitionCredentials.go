package lakeformation

// GetTemporaryGluePartitionCredentials is generated as a reference stub.
// Executable command wiring lives under cmd/lakeformation.go.
//
// This API is identical to GetTemporaryTableCredentials except that this is used
// when the target Data Catalog resource is of type Partition. Lake Formation
// restricts the permission of the vended credentials with the same scope down
// policy which restricts access to a single Amazon S3 prefix.
