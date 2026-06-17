package lakeformation

// DeleteObjectsOnCancel is generated as a reference stub.
// Executable command wiring lives under cmd/lakeformation.go.
//
// For a specific governed table, provides a list of Amazon S3 objects that will
// be written during the current transaction and that can be automatically deleted
// if the transaction is canceled. Without this call, no Amazon S3 objects are
// automatically deleted when a transaction cancels.
//
// The Glue ETL library function write_dynamic_frame.from_catalog() includes an
// option to automatically call DeleteObjectsOnCancel before writes. For more
// information, see [Rolling Back Amazon S3 Writes].
//
// [Rolling Back Amazon S3 Writes]: https://docs.aws.amazon.com/lake-formation/latest/dg/transactions-data-operations.html#rolling-back-writes
