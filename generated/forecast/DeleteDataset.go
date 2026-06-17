package forecast

// DeleteDataset is generated as a reference stub.
// Executable command wiring lives under cmd/forecast.go.
//
// Deletes an Amazon Forecast dataset that was created using the [CreateDataset] operation. You
// can only delete datasets that have a status of ACTIVE or CREATE_FAILED . To get
// the status use the [DescribeDataset]operation.
//
// Forecast does not automatically update any dataset groups that contain the
// deleted dataset. In order to update the dataset group, use the [UpdateDatasetGroup]operation,
// omitting the deleted dataset's ARN.
//
// [CreateDataset]: https://docs.aws.amazon.com/forecast/latest/dg/API_CreateDataset.html
// [UpdateDatasetGroup]: https://docs.aws.amazon.com/forecast/latest/dg/API_UpdateDatasetGroup.html
// [DescribeDataset]: https://docs.aws.amazon.com/forecast/latest/dg/API_DescribeDataset.html
