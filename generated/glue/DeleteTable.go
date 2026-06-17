package glue

// DeleteTable is generated as a reference stub.
// Executable command wiring lives under cmd/glue.go.
//
// Removes a table definition from the Data Catalog.
//
// After completing this operation, you no longer have access to the table
// versions and partitions that belong to the deleted table. Glue deletes these
// "orphaned" resources asynchronously in a timely manner, at the discretion of the
// service.
//
// To ensure the immediate deletion of all related resources, before calling
// DeleteTable , use DeleteTableVersion or BatchDeleteTableVersion , and
// DeletePartition or BatchDeletePartition , to delete any resources that belong to
// the table.
