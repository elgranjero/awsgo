package rekognition

// StartProjectVersion is generated as a reference stub.
// Executable command wiring lives under cmd/rekognition.go.
//
// This operation applies only to Amazon Rekognition Custom Labels.
//
// Starts the running of the version of a model. Starting a model takes a while to
// complete. To check the current state of the model, use DescribeProjectVersions.
//
// Once the model is running, you can detect custom labels in new images by
// calling DetectCustomLabels.
//
// You are charged for the amount of time that the model is running. To stop a
// running model, call StopProjectVersion.
//
// This operation requires permissions to perform the
// rekognition:StartProjectVersion action.
