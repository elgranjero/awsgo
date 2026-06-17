package personalize

// CreateDataset is generated as a reference stub.
// Executable command wiring lives under cmd/personalize.go.
//
// Creates an empty dataset and adds it to the specified dataset group. Use [CreateDatasetImportJob] to
// import your training data to a dataset.
//
// There are 5 types of datasets:
//
// - Item interactions
//
// - Items
//
// - Users
//
// - Action interactions
//
// - Actions
//
// Each dataset type has an associated schema with required field types. Only the
// Item interactions dataset is required in order to train a model (also referred
// to as creating a solution).
//
// A dataset can be in one of the following states:
//
// - CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE FAILED
//
// - DELETE PENDING > DELETE IN_PROGRESS
//
// To get the status of the dataset, call [DescribeDataset].
//
// # Related APIs
//
// [CreateDatasetGroup]
//
// [ListDatasets]
//
// [DescribeDataset]
//
// [DeleteDataset]
//
// [CreateDatasetImportJob]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateDatasetImportJob.html
// [DescribeDataset]: https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeDataset.html
// [ListDatasets]: https://docs.aws.amazon.com/personalize/latest/dg/API_ListDatasets.html
// [DeleteDataset]: https://docs.aws.amazon.com/personalize/latest/dg/API_DeleteDataset.html
// [CreateDatasetGroup]: https://docs.aws.amazon.com/personalize/latest/dg/API_CreateDatasetGroup.html
