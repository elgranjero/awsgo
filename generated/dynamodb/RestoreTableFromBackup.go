package dynamodb

// RestoreTableFromBackup is generated as a reference stub.
// Executable command wiring lives under cmd/dynamodb.go.
//
// Creates a new table from an existing backup. Any number of users can execute up
// to 50 concurrent restores (any type of restore) in a given account.
//
// You can call RestoreTableFromBackup at a maximum rate of 10 times per second.
//
// You must manually set up the following on the restored table:
//
// - Auto scaling policies
//
// - IAM policies
//
// - Amazon CloudWatch metrics and alarms
//
// - Tags
//
// - Stream settings
//
// - Time to Live (TTL) settings
