package dynamodb

// RestoreTableToPointInTime is generated as a reference stub.
// Executable command wiring lives under cmd/dynamodb.go.
//
// Restores the specified table to the specified point in time within
// EarliestRestorableDateTime and LatestRestorableDateTime . You can restore your
// table to any point in time in the last 35 days. You can set the recovery period
// to any value between 1 and 35 days. Any number of users can execute up to 50
// concurrent restores (any type of restore) in a given account.
//
// When you restore using point in time recovery, DynamoDB restores your table
// data to the state based on the selected date and time (day:hour:minute:second)
// to a new table.
//
// Along with data, the following are also included on the new restored table
// using point in time recovery:
//
// - Global secondary indexes (GSIs)
//
// - Local secondary indexes (LSIs)
//
// - Provisioned read and write capacity
//
// - Encryption settings
//
// All these settings come from the current settings of the source table at the
//
// time of restore.
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
//
// - Point in time recovery settings
