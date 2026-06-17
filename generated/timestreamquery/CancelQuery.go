package timestreamquery

// CancelQuery is generated as a reference stub.
// Executable command wiring lives under cmd/timestreamquery.go.
//
// Cancels a query that has been issued. Cancellation is provided only if the
//
// query has not completed running before the cancellation request was issued.
// Because cancellation is an idempotent operation, subsequent cancellation
// requests will return a CancellationMessage , indicating that the query has
// already been canceled. See [code sample]for details.
//
// [code sample]: https://docs.aws.amazon.com/timestream/latest/developerguide/code-samples.cancel-query.html
