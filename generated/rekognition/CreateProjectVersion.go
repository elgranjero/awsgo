package rekognition

// CreateProjectVersion is generated as a reference stub.
// Executable command wiring lives under cmd/rekognition.go.
//
// Creates a new version of Amazon Rekognition project (like a Custom Labels model
// or a custom adapter) and begins training. Models and adapters are managed as
// part of a Rekognition project. The response from CreateProjectVersion is an
// Amazon Resource Name (ARN) for the project version.
//
// The FeatureConfig operation argument allows you to configure specific model or
// adapter settings. You can provide a description to the project version by using
// the VersionDescription argment. Training can take a while to complete. You can
// get the current status by calling DescribeProjectVersions. Training completed successfully if the
// value of the Status field is TRAINING_COMPLETED . Once training has successfully
// completed, call DescribeProjectVersionsto get the training results and evaluate the model.
//
// This operation requires permissions to perform the
// rekognition:CreateProjectVersion action.
//
// The following applies only to projects with Amazon Rekognition Custom Labels as
// the chosen feature:
//
// You can train a model in a project that doesn't have associated datasets by
// specifying manifest files in the TrainingData and TestingData fields.
//
// If you open the console after training a model with manifest files, Amazon
// Rekognition Custom Labels creates the datasets for you using the most recent
// manifest files. You can no longer train a model version for the project by
// specifying manifest files.
//
// Instead of training with a project without associated datasets, we recommend
// that you use the manifest files to create training and test datasets for the
// project.
