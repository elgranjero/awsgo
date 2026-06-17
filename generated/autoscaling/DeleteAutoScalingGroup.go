package autoscaling

// DeleteAutoScalingGroup is generated as a reference stub.
// Executable command wiring lives under cmd/autoscaling.go.
//
// Deletes the specified Auto Scaling group.
//
// If the group has instances or scaling activities in progress, you must specify
// the option to force the deletion in order for it to succeed. The force delete
// operation will also terminate the EC2 instances. If the group has a warm pool,
// the force delete option also deletes the warm pool.
//
// To remove instances from the Auto Scaling group before deleting it, call the [DetachInstances]
// API with the list of instances and the option to decrement the desired capacity.
// This ensures that Amazon EC2 Auto Scaling does not launch replacement instances.
//
// To terminate all instances before deleting the Auto Scaling group, call the [UpdateAutoScalingGroup]
// API and set the minimum size and desired capacity of the Auto Scaling group to
// zero.
//
// If the group has scaling policies, deleting the group deletes the policies, the
// underlying alarm actions, and any alarm that no longer has an associated action.
//
// For more information, see [Delete your Auto Scaling infrastructure] in the Amazon EC2 Auto Scaling User Guide.
//
// [Delete your Auto Scaling infrastructure]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-process-shutdown.html
// [DetachInstances]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DetachInstances.html
// [UpdateAutoScalingGroup]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_UpdateAutoScalingGroup.html
