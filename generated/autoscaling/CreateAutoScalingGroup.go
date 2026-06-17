package autoscaling

// CreateAutoScalingGroup is generated as a reference stub.
// Executable command wiring lives under cmd/autoscaling.go.
//
// We strongly recommend using a launch template when calling this operation to
//
// ensure full functionality for Amazon EC2 Auto Scaling and Amazon EC2.
//
// Creates an Auto Scaling group with the specified name and attributes.
//
// If you exceed your maximum limit of Auto Scaling groups, the call fails. To
// query this limit, call the [DescribeAccountLimits]API. For information about updating this limit, see [Quotas for Amazon EC2 Auto Scaling]
// in the Amazon EC2 Auto Scaling User Guide.
//
// If you're new to Amazon EC2 Auto Scaling, see the introductory tutorials in [Get started with Amazon EC2 Auto Scaling] in
// the Amazon EC2 Auto Scaling User Guide.
//
// Every Auto Scaling group has three size properties ( DesiredCapacity , MaxSize ,
// and MinSize ). Usually, you set these sizes based on a specific number of
// instances. However, if you configure a mixed instances policy that defines
// weights for the instance types, you must specify these sizes with the same units
// that you use for weighting instances.
//
// [DescribeAccountLimits]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribeAccountLimits.html
// [Get started with Amazon EC2 Auto Scaling]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/get-started-with-ec2-auto-scaling.html
// [Quotas for Amazon EC2 Auto Scaling]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-quotas.html
