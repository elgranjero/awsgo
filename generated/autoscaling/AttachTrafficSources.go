package autoscaling

// AttachTrafficSources is generated as a reference stub.
// Executable command wiring lives under cmd/autoscaling.go.
//
// Attaches one or more traffic sources to the specified Auto Scaling group.
//
// You can use any of the following as traffic sources for an Auto Scaling group:
//
// - Application Load Balancer
//
// - Classic Load Balancer
//
// - Gateway Load Balancer
//
// - Network Load Balancer
//
// - VPC Lattice
//
// This operation is additive and does not detach existing traffic sources from
// the Auto Scaling group.
//
// After the operation completes, use the [DescribeTrafficSources] API to return details about the state
// of the attachments between traffic sources and your Auto Scaling group. To
// detach a traffic source from the Auto Scaling group, call the [DetachTrafficSources]API.
//
// [DescribeTrafficSources]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribeTrafficSources.html
// [DetachTrafficSources]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DetachTrafficSources.html
