package redshift

// RestoreTableFromClusterSnapshot is generated as a reference stub.
// Executable command wiring lives under cmd/redshift.go.
//
// Creates a new table from a table in an Amazon Redshift cluster snapshot. You
// must create the new table within the Amazon Redshift cluster that the snapshot
// was taken from.
//
// You cannot use RestoreTableFromClusterSnapshot to restore a table with the same
// name as an existing table in an Amazon Redshift cluster. That is, you cannot
// overwrite an existing table in a cluster with a restored table. If you want to
// replace your original table with a new, restored table, then rename or drop your
// original table before you call RestoreTableFromClusterSnapshot . When you have
// renamed your original table, then you can pass the original name of the table as
// the NewTableName parameter value in the call to RestoreTableFromClusterSnapshot
// . This way, you can replace the original table with the table created from the
// snapshot.
//
// You can't use this operation to restore tables with [interleaved sort keys].
//
// [interleaved sort keys]: https://docs.aws.amazon.com/redshift/latest/dg/t_Sorting_data.html#t_Sorting_data-interleaved
