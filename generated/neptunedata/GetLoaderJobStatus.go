package neptunedata

// GetLoaderJobStatus is generated as a reference stub.
// Executable command wiring lives under cmd/neptunedata.go.
//
// Gets status information about a specified load job. Neptune keeps track of the
// most recent 1,024 bulk load jobs, and stores the last 10,000 error details per
// job.
//
// See [Neptune Loader Get-Status API] for more information.
//
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:GetLoaderJobStatus]IAM action in that cluster..
//
// [neptune-db:GetLoaderJobStatus]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#getloaderjobstatus
// [Neptune Loader Get-Status API]: https://docs.aws.amazon.com/neptune/latest/userguide/load-api-reference-status.htm
