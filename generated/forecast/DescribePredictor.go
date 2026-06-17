package forecast

// DescribePredictor is generated as a reference stub.
// Executable command wiring lives under cmd/forecast.go.
//
// This operation is only valid for legacy predictors created with
//
// CreatePredictor. If you are not using a legacy predictor, use DescribeAutoPredictor.
//
// Describes a predictor created using the CreatePredictor operation.
//
// In addition to listing the properties provided in the CreatePredictor request,
// this operation lists the following properties:
//
// - DatasetImportJobArns - The dataset import jobs used to import training data.
//
// - AutoMLAlgorithmArns - If AutoML is performed, the algorithms that were
// evaluated.
//
// - CreationTime
//
// - LastModificationTime
//
// - Status
//
// - Message - If an error occurred, information about the error.
