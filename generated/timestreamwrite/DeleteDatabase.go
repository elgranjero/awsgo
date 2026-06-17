package timestreamwrite

// DeleteDatabase is generated as a reference stub.
// Executable command wiring lives under cmd/timestreamwrite.go.
//
// Deletes a given Timestream database. This is an irreversible operation. After a
// database is deleted, the time-series data from its tables cannot be recovered.
//
// All tables in the database must be deleted first, or a ValidationException
// error will be thrown.
//
// Due to the nature of distributed retries, the operation can return either
// success or a ResourceNotFoundException. Clients should consider them equivalent.
//
// See [code sample] for details.
//
// [code sample]: https://docs.aws.amazon.com/timestream/latest/developerguide/code-samples.delete-db.html
