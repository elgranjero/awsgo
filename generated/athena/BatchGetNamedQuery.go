package athena

// BatchGetNamedQuery is generated as a reference stub.
// Executable command wiring lives under cmd/athena.go.
//
// Returns the details of a single named query or a list of up to 50 queries,
// which you provide as an array of query ID strings. Requires you to have access
// to the workgroup in which the queries were saved. Use ListNamedQueriesInputto get the list of named
// query IDs in the specified workgroup. If information could not be retrieved for
// a submitted query ID, information about the query ID submitted is listed under UnprocessedNamedQueryId
// . Named queries differ from executed queries. Use BatchGetQueryExecutionInputto get details about each
// unique query execution, and ListQueryExecutionsInputto get a list of query execution IDs.
