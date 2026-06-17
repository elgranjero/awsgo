package timestreamwrite

// DeleteTable is generated as a reference stub.
// Executable command wiring lives under cmd/timestreamwrite.go.
//
// Deletes a given Timestream table. This is an irreversible operation. After a
// Timestream database table is deleted, the time-series data stored in the table
// cannot be recovered.
//
// Due to the nature of distributed retries, the operation can return either
// success or a ResourceNotFoundException. Clients should consider them equivalent.
//
// See [code sample] for details.
//
// [code sample]: https://docs.aws.amazon.com/timestream/latest/developerguide/code-samples.delete-table.html
