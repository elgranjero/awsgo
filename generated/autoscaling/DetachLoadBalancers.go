package autoscaling

// DetachLoadBalancers is generated as a reference stub.
// Executable command wiring lives under cmd/autoscaling.go.
//
// This API operation is superseded by [DetachTrafficSources], which can detach multiple traffic sources
// types. We recommend using DetachTrafficSources to simplify how you manage
// traffic sources. However, we continue to support DetachLoadBalancers . You can
// use both the original DetachLoadBalancers API operation and DetachTrafficSources
// on the same Auto Scaling group.
//
// Detaches one or more Classic Load Balancers from the specified Auto Scaling
// group.
//
// This operation detaches only Classic Load Balancers. If you have Application
// Load Balancers, Network Load Balancers, or Gateway Load Balancers, use the [DetachLoadBalancerTargetGroups]API
// instead.
//
// When you detach a load balancer, it enters the Removing state while
// deregistering the instances in the group. When all instances are deregistered,
// then you can no longer describe the load balancer using the [DescribeLoadBalancers]API call. The
// instances remain running.
//
// [DetachLoadBalancerTargetGroups]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DetachLoadBalancerTargetGroups.html
// [DescribeLoadBalancers]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribeLoadBalancers.html
// [DetachTrafficSources]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DetachTrafficSources.html
