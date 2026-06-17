package dynamodb

// UpdateTable is generated as a reference stub.
// Executable command wiring lives under cmd/dynamodb.go.
//
// Modifies the provisioned throughput settings, global secondary indexes, or
// DynamoDB Streams settings for a given table.
//
// You can only perform one of the following operations at once:
//
// - Modify the provisioned throughput settings of the table.
//
// - Remove a global secondary index from the table.
//
// - Create a new global secondary index on the table. After the index begins
// backfilling, you can use UpdateTable to perform other operations.
//
// UpdateTable is an asynchronous operation; while it's executing, the table
// status changes from ACTIVE to UPDATING . While it's UPDATING , you can't issue
// another UpdateTable request. When the table returns to the ACTIVE state, the
// UpdateTable operation is complete.
