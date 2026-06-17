package applicationautoscaling

// DeleteScalingPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/applicationautoscaling.go.
//
// Deletes the specified scaling policy for an Application Auto Scaling scalable
// target.
//
// Deleting a step scaling policy deletes the underlying alarm action, but does
// not delete the CloudWatch alarm associated with the scaling policy, even if it
// no longer has an associated action.
//
// For more information, see [Delete a step scaling policy] and [Delete a target tracking scaling policy] in the Application Auto Scaling User Guide.
//
// [Delete a target tracking scaling policy]: https://docs.aws.amazon.com/autoscaling/application/userguide/create-target-tracking-policy-cli.html#delete-target-tracking-policy
// [Delete a step scaling policy]: https://docs.aws.amazon.com/autoscaling/application/userguide/create-step-scaling-policy-cli.html#delete-step-scaling-policy
