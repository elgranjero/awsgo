package kendra

// DeleteDataSource is generated as a reference stub.
// Executable command wiring lives under cmd/kendra.go.
//
// Deletes an Amazon Kendra data source connector. An exception is not thrown if
// the data source is already being deleted. While the data source is being
// deleted, the Status field returned by a call to the DescribeDataSource API is
// set to DELETING . For more information, see [Deleting Data Sources].
//
// Deleting an entire data source or re-syncing your index after deleting specific
// documents from a data source could take up to an hour or more, depending on the
// number of documents you want to delete.
//
// [Deleting Data Sources]: https://docs.aws.amazon.com/kendra/latest/dg/delete-data-source.html
