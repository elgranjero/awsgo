package comprehend

// CreateFlywheel is generated as a reference stub.
// Executable command wiring lives under cmd/comprehend.go.
//
// A flywheel is an Amazon Web Services resource that orchestrates the ongoing
// training of a model for custom classification or custom entity recognition. You
// can create a flywheel to start with an existing trained model, or Comprehend can
// create and train a new model.
//
// When you create the flywheel, Comprehend creates a data lake in your account.
// The data lake holds the training data and test data for all versions of the
// model.
//
// To use a flywheel with an existing trained model, you specify the active model
// version. Comprehend copies the model's training data and test data into the
// flywheel's data lake.
//
// To use the flywheel with a new model, you need to provide a dataset for
// training data (and optional test data) when you create the flywheel.
//
// For more information about flywheels, see [Flywheel overview] in the Amazon Comprehend Developer
// Guide.
//
// [Flywheel overview]: https://docs.aws.amazon.com/comprehend/latest/dg/flywheels-about.html
