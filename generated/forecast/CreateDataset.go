package forecast

// CreateDataset is generated as a reference stub.
// Executable command wiring lives under cmd/forecast.go.
//
// Creates an Amazon Forecast dataset. The information about the dataset that you
// provide helps Forecast understand how to consume the data for model training.
// This includes the following:
//
// - DataFrequency - How frequently your historical time-series data is collected.
//
// - Domain and DatasetType - Each dataset has an associated dataset domain and a
// type within the domain. Amazon Forecast provides a list of predefined domains
// and types within each domain. For each unique dataset domain and type within the
// domain, Amazon Forecast requires your data to include a minimum set of
// predefined fields.
//
// - Schema - A schema specifies the fields in the dataset, including the field
// name and data type.
//
// After creating a dataset, you import your training data into it and add the
// dataset to a dataset group. You use the dataset group to create a predictor. For
// more information, see [Importing datasets].
//
// To get a list of all your datasets, use the [ListDatasets] operation.
//
// For example Forecast datasets, see the [Amazon Forecast Sample GitHub repository].
//
// The Status of a dataset must be ACTIVE before you can import training data. Use
// the [DescribeDataset]operation to get the status.
//
// [Amazon Forecast Sample GitHub repository]: https://github.com/aws-samples/amazon-forecast-samples
// [DescribeDataset]: https://docs.aws.amazon.com/forecast/latest/dg/API_DescribeDataset.html
// [ListDatasets]: https://docs.aws.amazon.com/forecast/latest/dg/API_ListDatasets.html
// [Importing datasets]: https://docs.aws.amazon.com/forecast/latest/dg/howitworks-datasets-groups.html
