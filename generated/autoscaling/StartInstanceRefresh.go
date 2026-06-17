package autoscaling

// StartInstanceRefresh is generated as a reference stub.
// Executable command wiring lives under cmd/autoscaling.go.
//
// Starts an instance refresh.
//
// This operation is part of the [instance refresh feature] in Amazon EC2 Auto Scaling, which helps you
// update instances in your Auto Scaling group. This feature is helpful, for
// example, when you have a new AMI or a new user data script. You just need to
// create a new launch template that specifies the new AMI or user data script.
// Then start an instance refresh to immediately begin the process of updating
// instances in the group.
//
// If successful, the request's response contains a unique ID that you can use to
// track the progress of the instance refresh. To query its status, call the [DescribeInstanceRefreshes]API.
// To describe the instance refreshes that have already run, call the [DescribeInstanceRefreshes]API. To
// cancel an instance refresh that is in progress, use the [CancelInstanceRefresh]API.
//
// An instance refresh might fail for several reasons, such as EC2 launch
// failures, misconfigured health checks, or not ignoring or allowing the
// termination of instances that are in Standby state or protected from scale in.
// You can monitor for failed EC2 launches using the scaling activities. To find
// the scaling activities, call the [DescribeScalingActivities]API.
//
// If you enable auto rollback, your Auto Scaling group will be rolled back
// automatically when the instance refresh fails. You can enable this feature
// before starting an instance refresh by specifying the AutoRollback property in
// the instance refresh preferences. Otherwise, to roll back an instance refresh
// before it finishes, use the [RollbackInstanceRefresh]API.
//
// [DescribeScalingActivities]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribeScalingActivities.html
// [instance refresh feature]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-instance-refresh.html
// [DescribeInstanceRefreshes]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribeInstanceRefreshes.html
// [CancelInstanceRefresh]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_CancelInstanceRefresh.html
// [RollbackInstanceRefresh]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_RollbackInstanceRefresh.html
