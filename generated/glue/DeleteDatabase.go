package glue

// DeleteDatabase is generated as a reference stub.
// Executable command wiring lives under cmd/glue.go.
//
// Removes a specified database from a Data Catalog.
//
// After completing this operation, you no longer have access to the tables (and
// all table versions and partitions that might belong to the tables) and the
// user-defined functions in the deleted database. Glue deletes these "orphaned"
// resources asynchronously in a timely manner, at the discretion of the service.
//
// To ensure the immediate deletion of all related resources, before calling
// DeleteDatabase , use DeleteTableVersion or BatchDeleteTableVersion ,
// DeletePartition or BatchDeletePartition , DeleteUserDefinedFunction , and
// DeleteTable or BatchDeleteTable , to delete any resources that belong to the
// database.
