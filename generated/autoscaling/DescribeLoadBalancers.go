package autoscaling

// DescribeLoadBalancers is generated as a reference stub.
// Executable command wiring lives under cmd/autoscaling.go.
//
// This API operation is superseded by [DescribeTrafficSources], which can describe multiple traffic
// sources types. We recommend using DescribeTrafficSources to simplify how you
// manage traffic sources. However, we continue to support DescribeLoadBalancers .
// You can use both the original DescribeLoadBalancers API operation and
// DescribeTrafficSources on the same Auto Scaling group.
//
// Gets information about the load balancers for the specified Auto Scaling group.
//
// This operation describes only Classic Load Balancers. If you have Application
// Load Balancers, Network Load Balancers, or Gateway Load Balancers, use the [DescribeLoadBalancerTargetGroups]API
// instead.
//
// To determine the attachment status of the load balancer, use the State element
// in the response. When you attach a load balancer to an Auto Scaling group, the
// initial State value is Adding . The state transitions to Added after all Auto
// Scaling instances are registered with the load balancer. If Elastic Load
// Balancing health checks are enabled for the Auto Scaling group, the state
// transitions to InService after at least one Auto Scaling instance passes the
// health check. When the load balancer is in the InService state, Amazon EC2 Auto
// Scaling can terminate and replace any instances that are reported as unhealthy.
// If no registered instances pass the health checks, the load balancer doesn't
// enter the InService state.
//
// Load balancers also have an InService state if you attach them in the [CreateAutoScalingGroup] API
// call. If your load balancer state is InService , but it is not working properly,
// check the scaling activities by calling [DescribeScalingActivities]and take any corrective actions
// necessary.
//
// For help with failed health checks, see [Troubleshooting Amazon EC2 Auto Scaling: Health checks] in the Amazon EC2 Auto Scaling User
// Guide. For more information, see [Use Elastic Load Balancing to distribute traffic across the instances in your Auto Scaling group]in the Amazon EC2 Auto Scaling User Guide.
//
// [Troubleshooting Amazon EC2 Auto Scaling: Health checks]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/ts-as-healthchecks.html
// [DescribeScalingActivities]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribeScalingActivities.html
// [DescribeLoadBalancerTargetGroups]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribeLoadBalancerTargetGroups.html
// [CreateAutoScalingGroup]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_CreateAutoScalingGroup.html
// [DescribeTrafficSources]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribeTrafficSources.html
// [Use Elastic Load Balancing to distribute traffic across the instances in your Auto Scaling group]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/autoscaling-load-balancer.html
