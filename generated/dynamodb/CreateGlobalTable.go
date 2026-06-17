package dynamodb

// CreateGlobalTable is generated as a reference stub.
// Executable command wiring lives under cmd/dynamodb.go.
//
// Creates a global table from an existing table. A global table creates a
// replication relationship between two or more DynamoDB tables with the same table
// name in the provided Regions.
//
// This documentation is for version 2017.11.29 (Legacy) of global tables, which
// should be avoided for new global tables. Customers should use [Global Tables version 2019.11.21 (Current)]when possible,
// because it provides greater flexibility, higher efficiency, and consumes less
// write capacity than 2017.11.29 (Legacy).
//
// To determine which version you're using, see [Determining the global table version you are using]. To update existing global tables
// from version 2017.11.29 (Legacy) to version 2019.11.21 (Current), see [Upgrading global tables].
//
// If you want to add a new replica table to a global table, each of the following
// conditions must be true:
//
// - The table must have the same primary key as all of the other replicas.
//
// - The table must have the same name as all of the other replicas.
//
// - The table must have DynamoDB Streams enabled, with the stream containing
// both the new and the old images of the item.
//
// - None of the replica tables in the global table can contain any data.
//
// If global secondary indexes are specified, then the following conditions must
// also be met:
//
// - The global secondary indexes must have the same name.
//
// - The global secondary indexes must have the same hash key and sort key (if
// present).
//
// If local secondary indexes are specified, then the following conditions must
// also be met:
//
// - The local secondary indexes must have the same name.
//
// - The local secondary indexes must have the same hash key and sort key (if
// present).
//
// Write capacity settings should be set consistently across your replica tables
// and secondary indexes. DynamoDB strongly recommends enabling auto scaling to
// manage the write capacity settings for all of your global tables replicas and
// indexes.
//
// If you prefer to manage write capacity settings manually, you should provision
// equal replicated write capacity units to your replica tables. You should also
// provision equal replicated write capacity units to matching secondary indexes
// across your global table.
//
// [Global Tables version 2019.11.21 (Current)]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GlobalTables.html
// [Upgrading global tables]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/V2globaltables_upgrade.html
// [Determining the global table version you are using]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.DetermineVersion.html
