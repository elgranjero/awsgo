package autoscaling

// DeletePolicy is generated as a reference stub.
// Executable command wiring lives under cmd/autoscaling.go.
//
// Deletes the specified scaling policy.
//
// Deleting either a step scaling policy or a simple scaling policy deletes the
// underlying alarm action, but does not delete the alarm, even if it no longer has
// an associated action.
//
// For more information, see [Delete a scaling policy] in the Amazon EC2 Auto Scaling User Guide.
//
// [Delete a scaling policy]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/deleting-scaling-policy.html
