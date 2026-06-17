package dynamodb

// TransactGetItems is generated as a reference stub.
// Executable command wiring lives under cmd/dynamodb.go.
//
// TransactGetItems is a synchronous operation that atomically retrieves multiple
// items from one or more tables (but not from indexes) in a single account and
// Region. A TransactGetItems call can contain up to 100 TransactGetItem objects,
// each of which contains a Get structure that specifies an item to retrieve from
// a table in the account and Region. A call to TransactGetItems cannot retrieve
// items from tables in more than one Amazon Web Services account or Region. The
// aggregate size of the items in the transaction cannot exceed 4 MB.
//
// DynamoDB rejects the entire TransactGetItems request if any of the following is
// true:
//
// - A conflicting operation is in the process of updating an item to be read.
//
// - There is insufficient provisioned capacity for the transaction to be
// completed.
//
// - There is a user error, such as an invalid data format.
//
// - The aggregate size of the items in the transaction exceeded 4 MB.
