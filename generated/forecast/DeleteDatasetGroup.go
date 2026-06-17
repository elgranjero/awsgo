package forecast

// DeleteDatasetGroup is generated as a reference stub.
// Executable command wiring lives under cmd/forecast.go.
//
// Deletes a dataset group created using the [CreateDatasetGroup] operation. You can only delete
// dataset groups that have a status of ACTIVE , CREATE_FAILED , or UPDATE_FAILED .
// To get the status, use the [DescribeDatasetGroup]operation.
//
// This operation deletes only the dataset group, not the datasets in the group.
//
// [DescribeDatasetGroup]: https://docs.aws.amazon.com/forecast/latest/dg/API_DescribeDatasetGroup.html
// [CreateDatasetGroup]: https://docs.aws.amazon.com/forecast/latest/dg/API_CreateDatasetGroup.html
