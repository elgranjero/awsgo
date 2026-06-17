package dynamodb

// UpdateGlobalTable is generated as a reference stub.
// Executable command wiring lives under cmd/dynamodb.go.
//
// Adds or removes replicas in the specified global table. The global table must
// already exist to be able to use this operation. Any replica to be added must be
// empty, have the same name as the global table, have the same key schema, have
// DynamoDB Streams enabled, and have the same provisioned and maximum write
// capacity units.
//
// This documentation is for version 2017.11.29 (Legacy) of global tables, which
// should be avoided for new global tables. Customers should use [Global Tables version 2019.11.21 (Current)]when possible,
// because it provides greater flexibility, higher efficiency, and consumes less
// write capacity than 2017.11.29 (Legacy).
//
// To determine which version you're using, see [Determining the global table version you are using]. To update existing global tables
// from version 2017.11.29 (Legacy) to version 2019.11.21 (Current), see [Upgrading global tables].
//
// If you are using global tables [Version 2019.11.21] (Current) you can use [UpdateTable] instead.
//
// Although you can use UpdateGlobalTable to add replicas and remove replicas in a
// single request, for simplicity we recommend that you issue separate requests for
// adding or removing replicas.
//
// If global secondary indexes are specified, then the following conditions must
// also be met:
//
// - The global secondary indexes must have the same name.
//
// - The global secondary indexes must have the same hash key and sort key (if
// present).
//
// - The global secondary indexes must have the same provisioned and maximum
// write capacity units.
//
// [UpdateTable]: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_UpdateTable.html
// [Global Tables version 2019.11.21 (Current)]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GlobalTables.html
// [Upgrading global tables]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/V2globaltables_upgrade.html
// [Version 2019.11.21]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GlobalTables.html
// [Determining the global table version you are using]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.DetermineVersion.html
