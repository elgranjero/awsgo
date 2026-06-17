package sagemaker

// CreateHyperParameterTuningJob is generated as a reference stub.
// Executable command wiring lives under cmd/sagemaker.go.
//
// Starts a hyperparameter tuning job. A hyperparameter tuning job finds the best
// version of a model by running many training jobs on your dataset using the
// algorithm you choose and values for hyperparameters within ranges that you
// specify. It then chooses the hyperparameter values that result in a model that
// performs the best, as measured by an objective metric that you choose.
//
// A hyperparameter tuning job automatically creates Amazon SageMaker experiments,
// trials, and trial components for each training job that it runs. You can view
// these entities in Amazon SageMaker Studio. For more information, see [View Experiments, Trials, and Trial Components].
//
// Do not include any security-sensitive information including account access IDs,
// secrets, or tokens in any hyperparameter fields. As part of the shared
// responsibility model, you are responsible for any potential exposure,
// unauthorized access, or compromise of your sensitive data if caused by any
// security-sensitive information included in the request hyperparameter variable
// or plain text fields..
//
// [View Experiments, Trials, and Trial Components]: https://docs.aws.amazon.com/sagemaker/latest/dg/experiments-view-compare.html#experiments-view
