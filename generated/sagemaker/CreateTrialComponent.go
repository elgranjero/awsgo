package sagemaker

// CreateTrialComponent is generated as a reference stub.
// Executable command wiring lives under cmd/sagemaker.go.
//
// Creates a trial component, which is a stage of a machine learning trial. A
// trial is composed of one or more trial components. A trial component can be used
// in multiple trials.
//
// Trial components include pre-processing jobs, training jobs, and batch
// transform jobs.
//
// When you use SageMaker Studio or the SageMaker Python SDK, all experiments,
// trials, and trial components are automatically tracked, logged, and indexed.
// When you use the Amazon Web Services SDK for Python (Boto), you must use the
// logging APIs provided by the SDK.
//
// You can add tags to a trial component and then use the [Search] API to search for the
// tags.
//
// [Search]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_Search.html
