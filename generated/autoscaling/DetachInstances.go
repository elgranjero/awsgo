package autoscaling

// DetachInstances is generated as a reference stub.
// Executable command wiring lives under cmd/autoscaling.go.
//
// Removes one or more instances from the specified Auto Scaling group.
//
// After the instances are detached, you can manage them independent of the Auto
// Scaling group.
//
// If you do not specify the option to decrement the desired capacity, Amazon EC2
// Auto Scaling launches instances to replace the ones that are detached.
//
// If there is a Classic Load Balancer attached to the Auto Scaling group, the
// instances are deregistered from the load balancer. If there are target groups
// attached to the Auto Scaling group, the instances are deregistered from the
// target groups.
//
// For more information, see [Detach or attach instances] in the Amazon EC2 Auto Scaling User Guide.
//
// [Detach or attach instances]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-detach-attach-instances.html
