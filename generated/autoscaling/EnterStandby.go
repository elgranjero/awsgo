package autoscaling

// EnterStandby is generated as a reference stub.
// Executable command wiring lives under cmd/autoscaling.go.
//
// Moves the specified instances into the standby state.
//
// If you choose to decrement the desired capacity of the Auto Scaling group, the
// instances can enter standby as long as the desired capacity of the Auto Scaling
// group after the instances are placed into standby is equal to or greater than
// the minimum capacity of the group.
//
// If you choose not to decrement the desired capacity of the Auto Scaling group,
// the Auto Scaling group launches new instances to replace the instances on
// standby.
//
// For more information, see [Temporarily removing instances from your Auto Scaling group] in the Amazon EC2 Auto Scaling User Guide.
//
// [Temporarily removing instances from your Auto Scaling group]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-enter-exit-standby.html
