package databasemigrationservice

// ModifyEndpoint is generated as a reference stub.
// Executable command wiring lives under cmd/databasemigrationservice.go.
//
// Modifies the specified endpoint.
//
// For a MySQL source or target endpoint, don't explicitly specify the database
// using the DatabaseName request parameter on the ModifyEndpoint API call.
// Specifying DatabaseName when you modify a MySQL endpoint replicates all the
// task tables to this single database. For MySQL endpoints, you specify the
// database only when you specify the schema in the table-mapping rules of the DMS
// task.
