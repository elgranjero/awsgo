package autoscaling

// AttachLoadBalancers is generated as a reference stub.
// Executable command wiring lives under cmd/autoscaling.go.
//
// This API operation is superseded by [AttachTrafficSources], which can attach multiple traffic sources
// types. We recommend using AttachTrafficSources to simplify how you manage
// traffic sources. However, we continue to support AttachLoadBalancers . You can
// use both the original AttachLoadBalancers API operation and AttachTrafficSources
// on the same Auto Scaling group.
//
// Attaches one or more Classic Load Balancers to the specified Auto Scaling
// group. Amazon EC2 Auto Scaling registers the running instances with these
// Classic Load Balancers.
//
// To describe the load balancers for an Auto Scaling group, call the [DescribeLoadBalancers] API. To
// detach a load balancer from the Auto Scaling group, call the [DetachLoadBalancers]API.
//
// This operation is additive and does not detach existing Classic Load Balancers
// or target groups from the Auto Scaling group.
//
// For more information, see [Use Elastic Load Balancing to distribute traffic across the instances in your Auto Scaling group] in the Amazon EC2 Auto Scaling User Guide.
//
// [DetachLoadBalancers]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DetachLoadBalancers.html
// [DescribeLoadBalancers]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribeLoadBalancers.html
// [AttachTrafficSources]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_AttachTrafficSources.html
// [Use Elastic Load Balancing to distribute traffic across the instances in your Auto Scaling group]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/autoscaling-load-balancer.html
