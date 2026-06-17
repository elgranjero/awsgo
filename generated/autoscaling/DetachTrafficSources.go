package autoscaling

// DetachTrafficSources is generated as a reference stub.
// Executable command wiring lives under cmd/autoscaling.go.
//
// Detaches one or more traffic sources from the specified Auto Scaling group.
//
// When you detach a traffic source, it enters the Removing state while
// deregistering the instances in the group. When all instances are deregistered,
// then you can no longer describe the traffic source using the [DescribeTrafficSources]API call. The
// instances continue to run.
//
// [DescribeTrafficSources]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribeTrafficSources.html
