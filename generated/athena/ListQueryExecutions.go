package athena

// ListQueryExecutions is generated as a reference stub.
// Executable command wiring lives under cmd/athena.go.
//
// Provides a list of available query execution IDs for the queries in the
// specified workgroup. Athena keeps a query history for 45 days. If a workgroup is
// not specified, returns a list of query execution IDs for the primary workgroup.
// Requires you to have access to the workgroup in which the queries ran.
