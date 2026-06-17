package forecast

// CreateDatasetGroup is generated as a reference stub.
// Executable command wiring lives under cmd/forecast.go.
//
// Creates a dataset group, which holds a collection of related datasets. You can
// add datasets to the dataset group when you create the dataset group, or later by
// using the [UpdateDatasetGroup]operation.
//
// After creating a dataset group and adding datasets, you use the dataset group
// when you create a predictor. For more information, see [Dataset groups].
//
// To get a list of all your datasets groups, use the [ListDatasetGroups] operation.
//
// The Status of a dataset group must be ACTIVE before you can use the dataset
// group to create a predictor. To get the status, use the [DescribeDatasetGroup]operation.
//
// [ListDatasetGroups]: https://docs.aws.amazon.com/forecast/latest/dg/API_ListDatasetGroups.html
// [DescribeDatasetGroup]: https://docs.aws.amazon.com/forecast/latest/dg/API_DescribeDatasetGroup.html
// [UpdateDatasetGroup]: https://docs.aws.amazon.com/forecast/latest/dg/API_UpdateDatasetGroup.html
// [Dataset groups]: https://docs.aws.amazon.com/forecast/latest/dg/howitworks-datasets-groups.html
