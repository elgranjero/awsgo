package rekognition

// DistributeDatasetEntries is generated as a reference stub.
// Executable command wiring lives under cmd/rekognition.go.
//
// This operation applies only to Amazon Rekognition Custom Labels.
//
// Distributes the entries (images) in a training dataset across the training
// dataset and the test dataset for a project. DistributeDatasetEntries moves 20%
// of the training dataset images to the test dataset. An entry is a JSON Line that
// describes an image.
//
// You supply the Amazon Resource Names (ARN) of a project's training dataset and
// test dataset. The training dataset must contain the images that you want to
// split. The test dataset must be empty. The datasets must belong to the same
// project. To create training and test datasets for a project, call CreateDataset.
//
// Distributing a dataset takes a while to complete. To check the status call
// DescribeDataset . The operation is complete when the Status field for the
// training dataset and the test dataset is UPDATE_COMPLETE . If the dataset split
// fails, the value of Status is UPDATE_FAILED .
//
// This operation requires permissions to perform the
// rekognition:DistributeDatasetEntries action.
