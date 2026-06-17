package sagemaker

// CreateTrial is generated as a reference stub.
// Executable command wiring lives under cmd/sagemaker.go.
//
// Creates an SageMaker trial. A trial is a set of steps called trial components
// that produce a machine learning model. A trial is part of a single SageMaker
// experiment.
//
// When you use SageMaker Studio or the SageMaker Python SDK, all experiments,
// trials, and trial components are automatically tracked, logged, and indexed.
// When you use the Amazon Web Services SDK for Python (Boto), you must use the
// logging APIs provided by the SDK.
//
// You can add tags to a trial and then use the [Search] API to search for the tags.
//
// To get a list of all your trials, call the [ListTrials] API. To view a trial's properties,
// call the [DescribeTrial]API. To create a trial component, call the [CreateTrialComponent] API.
//
// [DescribeTrial]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeTrial.html
// [Search]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_Search.html
// [ListTrials]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_ListTrials.html
// [CreateTrialComponent]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateTrialComponent.html
