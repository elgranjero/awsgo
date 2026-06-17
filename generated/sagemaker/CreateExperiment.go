package sagemaker

// CreateExperiment is generated as a reference stub.
// Executable command wiring lives under cmd/sagemaker.go.
//
// Creates a SageMaker experiment. An experiment is a collection of trials that
// are observed, compared and evaluated as a group. A trial is a set of steps,
// called trial components, that produce a machine learning model.
//
// In the Studio UI, trials are referred to as run groups and trial components are
// referred to as runs.
//
// The goal of an experiment is to determine the components that produce the best
// model. Multiple trials are performed, each one isolating and measuring the
// impact of a change to one or more inputs, while keeping the remaining inputs
// constant.
//
// When you use SageMaker Studio or the SageMaker Python SDK, all experiments,
// trials, and trial components are automatically tracked, logged, and indexed.
// When you use the Amazon Web Services SDK for Python (Boto), you must use the
// logging APIs provided by the SDK.
//
// You can add tags to experiments, trials, trial components and then use the [Search] API
// to search for the tags.
//
// To add a description to an experiment, specify the optional Description
// parameter. To add a description later, or to change the description, call the [UpdateExperiment]
// API.
//
// To get a list of all your experiments, call the [ListExperiments] API. To view an experiment's
// properties, call the [DescribeExperiment]API. To get a list of all the trials associated with an
// experiment, call the [ListTrials]API. To create a trial call the [CreateTrial] API.
//
// [DescribeExperiment]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeExperiment.html
// [Search]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_Search.html
// [CreateTrial]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateTrial.html
// [ListExperiments]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_ListExperiments.html
// [UpdateExperiment]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_UpdateExperiment.html
// [ListTrials]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_ListTrials.html
