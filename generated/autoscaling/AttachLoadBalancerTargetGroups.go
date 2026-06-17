package autoscaling

// AttachLoadBalancerTargetGroups is generated as a reference stub.
// Executable command wiring lives under cmd/autoscaling.go.
//
// This API operation is superseded by [AttachTrafficSources], which can attach multiple traffic sources
// types. We recommend using AttachTrafficSources to simplify how you manage
// traffic sources. However, we continue to support AttachLoadBalancerTargetGroups
// . You can use both the original AttachLoadBalancerTargetGroups API operation
// and AttachTrafficSources on the same Auto Scaling group.
//
// Attaches one or more target groups to the specified Auto Scaling group.
//
// This operation is used with the following load balancer types:
//
// - Application Load Balancer - Operates at the application layer (layer 7) and
// supports HTTP and HTTPS.
//
// - Network Load Balancer - Operates at the transport layer (layer 4) and
// supports TCP, TLS, and UDP.
//
// - Gateway Load Balancer - Operates at the network layer (layer 3).
//
// To describe the target groups for an Auto Scaling group, call the [DescribeLoadBalancerTargetGroups] API. To
// detach the target group from the Auto Scaling group, call the [DetachLoadBalancerTargetGroups]API.
//
// This operation is additive and does not detach existing target groups or
// Classic Load Balancers from the Auto Scaling group.
//
// For more information, see [Use Elastic Load Balancing to distribute traffic across the instances in your Auto Scaling group] in the Amazon EC2 Auto Scaling User Guide.
//
// [DescribeLoadBalancerTargetGroups]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribeLoadBalancerTargetGroups.html
// [DetachLoadBalancerTargetGroups]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DetachLoadBalancerTargetGroups.html
// [AttachTrafficSources]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_AttachTrafficSources.html
// [Use Elastic Load Balancing to distribute traffic across the instances in your Auto Scaling group]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/autoscaling-load-balancer.html
