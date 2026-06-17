package personalize

// CreateDatasetExportJob is generated as a reference stub.
// Executable command wiring lives under cmd/personalize.go.
//
// Creates a job that exports data from your dataset to an Amazon S3 bucket. To
//
// allow Amazon Personalize to export the training data, you must specify an
// service-linked IAM role that gives Amazon Personalize PutObject permissions for
// your Amazon S3 bucket. For information, see [Exporting a dataset]in the Amazon Personalize developer
// guide.
//
// # Status
//
// A dataset export job can be in one of the following states:
//
// - CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE FAILED
//
// To get the status of the export job, call [DescribeDatasetExportJob], and specify the Amazon Resource
// Name (ARN) of the dataset export job. The dataset export is complete when the
// status shows as ACTIVE. If the status shows as CREATE FAILED, the response
// includes a failureReason key, which describes why the job failed.
//
// [Exporting a dataset]: https://docs.aws.amazon.com/personalize/latest/dg/export-data.html
// [DescribeDatasetExportJob]: https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeDatasetExportJob.html
