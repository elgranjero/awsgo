package sagemaker

// CreateModelPackage is generated as a reference stub.
// Executable command wiring lives under cmd/sagemaker.go.
//
// Creates a model package that you can use to create SageMaker models or list on
// Amazon Web Services Marketplace, or a versioned model that is part of a model
// group. Buyers can subscribe to model packages listed on Amazon Web Services
// Marketplace to create models in SageMaker.
//
// To create a model package by specifying a Docker container that contains your
// inference code and the Amazon S3 location of your model artifacts, provide
// values for InferenceSpecification . To create a model from an algorithm resource
// that you created or subscribed to in Amazon Web Services Marketplace, provide a
// value for SourceAlgorithmSpecification .
//
// There are two types of model packages:
//
// - Versioned - a model that is part of a model group in the model registry.
//
// - Unversioned - a model package that is not part of a model group.
