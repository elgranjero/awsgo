package timestreamwrite

// CreateBatchLoadTask is generated as a reference stub.
// Executable command wiring lives under cmd/timestreamwrite.go.
//
// Creates a new Timestream batch load task. A batch load task processes data from
// a CSV source in an S3 location and writes to a Timestream table. A mapping from
// source to target is defined in a batch load task. Errors and events are written
// to a report at an S3 location. For the report, if the KMS key is not specified,
// the report will be encrypted with an S3 managed key when SSE_S3 is the option.
// Otherwise an error is thrown. For more information, see [Amazon Web Services managed keys]. [Service quotas apply]. For details, see [code sample].
//
// [Service quotas apply]: https://docs.aws.amazon.com/timestream/latest/developerguide/ts-limits.html
// [code sample]: https://docs.aws.amazon.com/timestream/latest/developerguide/code-samples.create-batch-load.html
// [Amazon Web Services managed keys]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-cmk
