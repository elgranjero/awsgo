package autoscaling

// RollbackInstanceRefresh is generated as a reference stub.
// Executable command wiring lives under cmd/autoscaling.go.
//
// Cancels an instance refresh that is in progress and rolls back any changes that
// it made. Amazon EC2 Auto Scaling replaces any instances that were replaced
// during the instance refresh. This restores your Auto Scaling group to the
// configuration that it was using before the start of the instance refresh.
//
// This operation is part of the [instance refresh feature] in Amazon EC2 Auto Scaling, which helps you
// update instances in your Auto Scaling group after you make configuration
// changes.
//
// A rollback is not supported in the following situations:
//
// - There is no desired configuration specified for the instance refresh.
//
// - The Auto Scaling group has a launch template that uses an Amazon Web
// Services Systems Manager parameter instead of an AMI ID for the ImageId
// property.
//
// - The Auto Scaling group uses the launch template's $Latest or $Default
// version.
//
// When you receive a successful response from this operation, Amazon EC2 Auto
// Scaling immediately begins replacing instances. You can check the status of this
// operation through the [DescribeInstanceRefreshes]API operation.
//
// [instance refresh feature]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-instance-refresh.html
// [DescribeInstanceRefreshes]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribeInstanceRefreshes.html
