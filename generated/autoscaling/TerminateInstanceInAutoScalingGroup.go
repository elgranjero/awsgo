package autoscaling

// TerminateInstanceInAutoScalingGroup is generated as a reference stub.
// Executable command wiring lives under cmd/autoscaling.go.
//
// Terminates the specified instance and optionally adjusts the desired group
// size. This operation cannot be called on instances in a warm pool.
//
// This call simply makes a termination request. The instance is not terminated
// immediately. When an instance is terminated, the instance status changes to
// terminated . You can't connect to or start an instance after you've terminated
// it.
//
// If you do not specify the option to decrement the desired capacity, Amazon EC2
// Auto Scaling launches instances to replace the ones that are terminated.
//
// By default, Amazon EC2 Auto Scaling balances instances across all Availability
// Zones. If you decrement the desired capacity, your Auto Scaling group can become
// unbalanced between Availability Zones. Amazon EC2 Auto Scaling tries to
// rebalance the group, and rebalancing might terminate instances in other zones.
// For more information, see [Manual scaling]in the Amazon EC2 Auto Scaling User Guide.
//
// [Manual scaling]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-scaling-manually.html
