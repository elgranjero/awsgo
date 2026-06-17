package glue

// SearchTables is generated as a reference stub.
// Executable command wiring lives under cmd/glue.go.
//
// Searches a set of tables based on properties in the table metadata as well as
// on the parent database. You can search against text or filter conditions.
//
// You can only get tables that you have access to based on the security policies
// defined in Lake Formation. You need at least a read-only access to the table for
// it to be returned. If you do not have access to all the columns in the table,
// these columns will not be searched against when returning the list of tables
// back to you. If you have access to the columns but not the data in the columns,
// those columns and the associated metadata for those columns will be included in
// the search.
