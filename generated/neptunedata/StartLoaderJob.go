package neptunedata

// StartLoaderJob is generated as a reference stub.
// Executable command wiring lives under cmd/neptunedata.go.
//
// Starts a Neptune bulk loader job to load data from an Amazon S3 bucket into a
// Neptune DB instance. See [Using the Amazon Neptune Bulk Loader to Ingest Data].
//
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:StartLoaderJob]IAM action in that cluster.
//
// [neptune-db:StartLoaderJob]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#startloaderjob
// [Using the Amazon Neptune Bulk Loader to Ingest Data]: https://docs.aws.amazon.com/neptune/latest/userguide/bulk-load.html
