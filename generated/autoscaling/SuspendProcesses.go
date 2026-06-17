package autoscaling

// SuspendProcesses is generated as a reference stub.
// Executable command wiring lives under cmd/autoscaling.go.
//
// Suspends the specified auto scaling processes, or all processes, for the
// specified Auto Scaling group.
//
// If you suspend either the Launch or Terminate process types, it can prevent
// other process types from functioning properly. For more information, see [Suspend and resume Amazon EC2 Auto Scaling processes]in the
// Amazon EC2 Auto Scaling User Guide.
//
// To resume processes that have been suspended, call the [ResumeProcesses] API.
//
// [ResumeProcesses]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_ResumeProcesses.html
// [Suspend and resume Amazon EC2 Auto Scaling processes]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-suspend-resume-processes.html
