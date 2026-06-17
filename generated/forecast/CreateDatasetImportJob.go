package forecast

// CreateDatasetImportJob is generated as a reference stub.
// Executable command wiring lives under cmd/forecast.go.
//
// Imports your training data to an Amazon Forecast dataset. You provide the
// location of your training data in an Amazon Simple Storage Service (Amazon S3)
// bucket and the Amazon Resource Name (ARN) of the dataset that you want to import
// the data to.
//
// You must specify a [DataSource] object that includes an Identity and Access Management
// (IAM) role that Amazon Forecast can assume to access the data, as Amazon
// Forecast makes a copy of your data and processes it in an internal Amazon Web
// Services system. For more information, see [Set up permissions].
//
// The training data must be in CSV or Parquet format. The delimiter must be a
// comma (,).
//
// You can specify the path to a specific file, the S3 bucket, or to a folder in
// the S3 bucket. For the latter two cases, Amazon Forecast imports all files up to
// the limit of 10,000 files.
//
// Because dataset imports are not aggregated, your most recent dataset import is
// the one that is used when training a predictor or generating a forecast. Make
// sure that your most recent dataset import contains all of the data you want to
// model off of, and not just the new data collected since the previous import.
//
// To get a list of all your dataset import jobs, filtered by specified criteria,
// use the [ListDatasetImportJobs]operation.
//
// [ListDatasetImportJobs]: https://docs.aws.amazon.com/forecast/latest/dg/API_ListDatasetImportJobs.html
// [Set up permissions]: https://docs.aws.amazon.com/forecast/latest/dg/aws-forecast-iam-roles.html
// [DataSource]: https://docs.aws.amazon.com/forecast/latest/dg/API_DataSource.html
