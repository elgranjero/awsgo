package bcmdataexports

// CreateExport is generated as a reference stub.
// Executable command wiring lives under cmd/bcmdataexports.go.
//
// Creates a data export and specifies the data query, the delivery preference,
// and any optional resource tags.
//
// A DataQuery consists of both a QueryStatement and TableConfigurations .
//
// The QueryStatement is an SQL statement. Data Exports only supports a limited
// subset of the SQL syntax. For more information on the SQL syntax that is
// supported, see [Data query]. To view the available tables and columns, see the [Data Exports table dictionary].
//
// The TableConfigurations is a collection of specified TableProperties for the
// table being queried in the QueryStatement . TableProperties are additional
// configurations you can provide to change the data and schema of a table. Each
// table can have different TableProperties. However, tables are not required to
// have any TableProperties. Each table property has a default value that it
// assumes if not specified. For more information on table configurations, see [Data query].
// To view the table properties available for each table, see the [Data Exports table dictionary]or use the
// ListTables API to get a response of all tables and their available properties.
//
// [Data Exports table dictionary]: https://docs.aws.amazon.com/cur/latest/userguide/de-table-dictionary.html
// [Data query]: https://docs.aws.amazon.com/cur/latest/userguide/de-data-query.html
