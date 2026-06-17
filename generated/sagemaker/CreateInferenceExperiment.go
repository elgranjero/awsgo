package sagemaker

// CreateInferenceExperiment is generated as a reference stub.
// Executable command wiring lives under cmd/sagemaker.go.
//
// Creates an inference experiment using the configurations specified in the
//
// request.
//
// Use this API to setup and schedule an experiment to compare model variants on a
// Amazon SageMaker inference endpoint. For more information about inference
// experiments, see [Shadow tests].
//
// Amazon SageMaker begins your experiment at the scheduled time and routes
// traffic to your endpoint's model variants based on your specified configuration.
//
// While the experiment is in progress or after it has concluded, you can view
// metrics that compare your model variants. For more information, see [View, monitor, and edit shadow tests].
//
// [Shadow tests]: https://docs.aws.amazon.com/sagemaker/latest/dg/shadow-tests.html
// [View, monitor, and edit shadow tests]: https://docs.aws.amazon.com/sagemaker/latest/dg/shadow-tests-view-monitor-edit.html
