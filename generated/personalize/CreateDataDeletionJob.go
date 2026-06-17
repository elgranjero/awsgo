package personalize

// CreateDataDeletionJob is generated as a reference stub.
// Executable command wiring lives under cmd/personalize.go.
//
// Creates a batch job that deletes all references to specific users from an
// Amazon Personalize dataset group in batches. You specify the users to delete in
// a CSV file of userIds in an Amazon S3 bucket. After a job completes, Amazon
// Personalize no longer trains on the users’ data and no longer considers the
// users when generating user segments. For more information about creating a data
// deletion job, see [Deleting users].
//
// - Your input file must be a CSV file with a single USER_ID column that lists
// the users IDs. For more information about preparing the CSV file, see [Preparing your data deletion file and uploading it to Amazon S3].
//
// - To give Amazon Personalize permission to access your input CSV file of
// userIds, you must specify an IAM service role that has permission to read from
// the data source. This role needs GetObject and ListBucket permissions for the
// bucket and its content. These permissions are the same as importing data. For
// information on granting access to your Amazon S3 bucket, see [Giving Amazon Personalize Access to Amazon S3 Resources].
//
// After you create a job, it can take up to a day to delete all references to the
// users from datasets and models. Until the job completes, Amazon Personalize
// continues to use the data when training. And if you use a User Segmentation
// recipe, the users might appear in user segments.
//
// # Status
//
// A data deletion job can have one of the following statuses:
//
// - PENDING > IN_PROGRESS > COMPLETED -or- FAILED
//
// To get the status of the data deletion job, call [DescribeDataDeletionJob] API operation and specify the
// Amazon Resource Name (ARN) of the job. If the status is FAILED, the response
// includes a failureReason key, which describes why the job failed.
//
// # Related APIs
//
// [ListDataDeletionJobs]
//
// [DescribeDataDeletionJob]
//
// [ListDataDeletionJobs]: https://docs.aws.amazon.com/personalize/latest/dg/API_ListDataDeletionJobs.html
// [Giving Amazon Personalize Access to Amazon S3 Resources]: https://docs.aws.amazon.com/personalize/latest/dg/granting-personalize-s3-access.html
// [Deleting users]: https://docs.aws.amazon.com/personalize/latest/dg/delete-records.html
// [Preparing your data deletion file and uploading it to Amazon S3]: https://docs.aws.amazon.com/personalize/latest/dg/prepare-deletion-input-file.html
// [DescribeDataDeletionJob]: https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeDataDeletionJob.html
