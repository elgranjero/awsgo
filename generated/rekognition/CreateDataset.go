package rekognition

// CreateDataset is generated as a reference stub.
// Executable command wiring lives under cmd/rekognition.go.
//
// This operation applies only to Amazon Rekognition Custom Labels.
//
// Creates a new Amazon Rekognition Custom Labels dataset. You can create a
// dataset by using an Amazon Sagemaker format manifest file or by copying an
// existing Amazon Rekognition Custom Labels dataset.
//
// To create a training dataset for a project, specify TRAIN for the value of
// DatasetType . To create the test dataset for a project, specify TEST for the
// value of DatasetType .
//
// The response from CreateDataset is the Amazon Resource Name (ARN) for the
// dataset. Creating a dataset takes a while to complete. Use DescribeDatasetto check the current
// status. The dataset created successfully if the value of Status is
// CREATE_COMPLETE .
//
// To check if any non-terminal errors occurred, call ListDatasetEntries and check for the presence
// of errors lists in the JSON Lines.
//
// Dataset creation fails if a terminal error occurs ( Status = CREATE_FAILED ).
// Currently, you can't access the terminal error information.
//
// For more information, see Creating dataset in the Amazon Rekognition Custom
// Labels Developer Guide.
//
// This operation requires permissions to perform the rekognition:CreateDataset
// action. If you want to copy an existing dataset, you also require permission to
// perform the rekognition:ListDatasetEntries action.
