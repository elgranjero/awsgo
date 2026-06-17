package dynamodb

// DescribeTable is generated as a reference stub.
// Executable command wiring lives under cmd/dynamodb.go.
//
// Returns information about the table, including the current status of the table,
// when it was created, the primary key schema, and any indexes on the table.
//
// If you issue a DescribeTable request immediately after a CreateTable request,
// DynamoDB might return a ResourceNotFoundException . This is because
// DescribeTable uses an eventually consistent query, and the metadata for your
// table might not be available at that moment. Wait for a few seconds, and then
// try the DescribeTable request again.
