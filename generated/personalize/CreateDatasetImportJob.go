package personalize

// CreateDatasetImportJob is generated as a reference stub.
// Executable command wiring lives under cmd/personalize.go.
//
// Creates a job that imports training data from your data source (an Amazon S3
// bucket) to an Amazon Personalize dataset. To allow Amazon Personalize to import
// the training data, you must specify an IAM service role that has permission to
// read from the data source, as Amazon Personalize makes a copy of your data and
// processes it internally. For information on granting access to your Amazon S3
// bucket, see [Giving Amazon Personalize Access to Amazon S3 Resources].
//
// If you already created a recommender or deployed a custom solution version with
// a campaign, how new bulk records influence recommendations depends on the domain
// use case or recipe that you use. For more information, see [How new data influences real-time recommendations].
//
// By default, a dataset import job replaces any existing data in the dataset that
// you imported in bulk. To add new records without replacing existing data,
// specify INCREMENTAL for the import mode in the CreateDatasetImportJob operation.
//
// # Status
//
// A dataset import job can be in one of the following states:
//
// - CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE FAILED
//
// To get the status of the import job, call [DescribeDatasetImportJob], providing the Amazon Resource Name
// (ARN) of the dataset import job. The dataset import is complete when the status
// shows as ACTIVE. If the status shows as CREATE FAILED, the response includes a
// failureReason key, which describes why the job failed.
//
// Importing takes time. You must wait until the status shows as ACTIVE before
// training a model using the dataset.
//
// # Related APIs
//
// [ListDatasetImportJobs]
//
// [DescribeDatasetImportJob]
//
// [ListDatasetImportJobs]: https://docs.aws.amazon.com/personalize/latest/dg/API_ListDatasetImportJobs.html
// [Giving Amazon Personalize Access to Amazon S3 Resources]: https://docs.aws.amazon.com/personalize/latest/dg/granting-personalize-s3-access.html
// [DescribeDatasetImportJob]: https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeDatasetImportJob.html
// [How new data influences real-time recommendations]: https://docs.aws.amazon.com/personalize/latest/dg/how-new-data-influences-recommendations.html
