package sagemaker

// DeleteTags is generated as a reference stub.
// Executable command wiring lives under cmd/sagemaker.go.
//
// Deletes the specified tags from an SageMaker resource.
//
// To list a resource's tags, use the ListTags API.
//
// When you call this API to delete tags from a hyperparameter tuning job, the
// deleted tags are not removed from training jobs that the hyperparameter tuning
// job launched before you called this API.
//
// When you call this API to delete tags from a SageMaker Domain or User Profile,
// the deleted tags are not removed from Apps that the SageMaker Domain or User
// Profile launched before you called this API.
