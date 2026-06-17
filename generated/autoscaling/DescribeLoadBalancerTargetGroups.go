package autoscaling

// DescribeLoadBalancerTargetGroups is generated as a reference stub.
// Executable command wiring lives under cmd/autoscaling.go.
//
// This API operation is superseded by [DescribeTrafficSources], which can describe multiple traffic
// sources types. We recommend using DetachTrafficSources to simplify how you
// manage traffic sources. However, we continue to support
// DescribeLoadBalancerTargetGroups . You can use both the original
// DescribeLoadBalancerTargetGroups API operation and DescribeTrafficSources on
// the same Auto Scaling group.
//
// Gets information about the Elastic Load Balancing target groups for the
// specified Auto Scaling group.
//
// To determine the attachment status of the target group, use the State element
// in the response. When you attach a target group to an Auto Scaling group, the
// initial State value is Adding . The state transitions to Added after all Auto
// Scaling instances are registered with the target group. If Elastic Load
// Balancing health checks are enabled for the Auto Scaling group, the state
// transitions to InService after at least one Auto Scaling instance passes the
// health check. When the target group is in the InService state, Amazon EC2 Auto
// Scaling can terminate and replace any instances that are reported as unhealthy.
// If no registered instances pass the health checks, the target group doesn't
// enter the InService state.
//
// Target groups also have an InService state if you attach them in the [CreateAutoScalingGroup] API call.
// If your target group state is InService , but it is not working properly, check
// the scaling activities by calling [DescribeScalingActivities]and take any corrective actions necessary.
//
// For help with failed health checks, see [Troubleshooting Amazon EC2 Auto Scaling: Health checks] in the Amazon EC2 Auto Scaling User
// Guide. For more information, see [Use Elastic Load Balancing to distribute traffic across the instances in your Auto Scaling group]in the Amazon EC2 Auto Scaling User Guide.
//
// You can use this operation to describe target groups that were attached by
// using [AttachLoadBalancerTargetGroups], but not for target groups that were attached by using [AttachTrafficSources].
//
// [Troubleshooting Amazon EC2 Auto Scaling: Health checks]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/ts-as-healthchecks.html
// [AttachLoadBalancerTargetGroups]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_AttachLoadBalancerTargetGroups.html
// [DescribeScalingActivities]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribeScalingActivities.html
// [CreateAutoScalingGroup]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_CreateAutoScalingGroup.html
// [DescribeTrafficSources]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribeTrafficSources.html
// [AttachTrafficSources]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_AttachTrafficSources.html
// [Use Elastic Load Balancing to distribute traffic across the instances in your Auto Scaling group]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/autoscaling-load-balancer.html
