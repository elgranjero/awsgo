package autoscaling

// DetachLoadBalancerTargetGroups is generated as a reference stub.
// Executable command wiring lives under cmd/autoscaling.go.
//
// This API operation is superseded by [DetachTrafficSources], which can detach multiple traffic sources
// types. We recommend using DetachTrafficSources to simplify how you manage
// traffic sources. However, we continue to support DetachLoadBalancerTargetGroups
// . You can use both the original DetachLoadBalancerTargetGroups API operation
// and DetachTrafficSources on the same Auto Scaling group.
//
// Detaches one or more target groups from the specified Auto Scaling group.
//
// When you detach a target group, it enters the Removing state while
// deregistering the instances in the group. When all instances are deregistered,
// then you can no longer describe the target group using the [DescribeLoadBalancerTargetGroups]API call. The
// instances remain running.
//
// You can use this operation to detach target groups that were attached by using [AttachLoadBalancerTargetGroups]
// , but not for target groups that were attached by using [AttachTrafficSources].
//
// [AttachLoadBalancerTargetGroups]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_AttachLoadBalancerTargetGroups.html
// [DescribeLoadBalancerTargetGroups]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribeLoadBalancerTargetGroups.html
// [DetachTrafficSources]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DetachTrafficSources.html
// [AttachTrafficSources]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_AttachTrafficSources.html
