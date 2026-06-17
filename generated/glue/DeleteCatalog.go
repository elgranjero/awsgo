package glue

// DeleteCatalog is generated as a reference stub.
// Executable command wiring lives under cmd/glue.go.
//
// Removes the specified catalog from the Glue Data Catalog.
//
// After completing this operation, you no longer have access to the databases,
// tables (and all table versions and partitions that might belong to the tables)
// and the user-defined functions in the deleted catalog. Glue deletes these
// "orphaned" resources asynchronously in a timely manner, at the discretion of the
// service.
//
// To ensure the immediate deletion of all related resources before calling the
// DeleteCatalog operation, use DeleteTableVersion (or BatchDeleteTableVersion ),
// DeletePartition (or BatchDeletePartition ), DeleteTable (or BatchDeleteTable ),
// DeleteUserDefinedFunction and DeleteDatabase to delete any resources that
// belong to the catalog.
