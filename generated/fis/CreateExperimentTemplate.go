package fis

// CreateExperimentTemplate is generated as a reference stub.
// Executable command wiring lives under cmd/fis.go.
//
// Creates an experiment template.
//
// An experiment template includes the following components:
//
// - Targets: A target can be a specific resource in your Amazon Web Services
// environment, or one or more resources that match criteria that you specify, for
// example, resources that have specific tags.
//
// - Actions: The actions to carry out on the target. You can specify multiple
// actions, the duration of each action, and when to start each action during an
// experiment.
//
// - Stop conditions: If a stop condition is triggered while an experiment is
// running, the experiment is automatically stopped. You can define a stop
// condition as a CloudWatch alarm.
//
// For more information, see [experiment templates] in the Fault Injection Service User Guide.
//
// [experiment templates]: https://docs.aws.amazon.com/fis/latest/userguide/experiment-templates.html
