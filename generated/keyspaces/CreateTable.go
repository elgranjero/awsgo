package keyspaces

// CreateTable is generated as a reference stub.
// Executable command wiring lives under cmd/keyspaces.go.
//
// The CreateTable operation adds a new table to the specified keyspace. Within a
// keyspace, table names must be unique.
//
// CreateTable is an asynchronous operation. When the request is received, the
// status of the table is set to CREATING . You can monitor the creation status of
// the new table by using the GetTable operation, which returns the current status
// of the table. You can start using a table when the status is ACTIVE .
//
// For more information, see [Create a table] in the Amazon Keyspaces Developer Guide.
//
// [Create a table]: https://docs.aws.amazon.com/keyspaces/latest/devguide/getting-started.tables.html
