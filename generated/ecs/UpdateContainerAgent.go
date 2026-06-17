package ecs

// UpdateContainerAgent is generated as a reference stub.
// Executable command wiring lives under cmd/ecs.go.
//
// Updates the Amazon ECS container agent on a specified container instance.
// Updating the Amazon ECS container agent doesn't interrupt running tasks or
// services on the container instance. The process for updating the agent differs
// depending on whether your container instance was launched with the Amazon
// ECS-optimized AMI or another operating system.
//
// The UpdateContainerAgent API isn't supported for container instances using the
// Amazon ECS-optimized Amazon Linux 2 (arm64) AMI. To update the container agent,
// you can update the ecs-init package. This updates the agent. For more
// information, see [Updating the Amazon ECS container agent]in the Amazon Elastic Container Service Developer Guide.
//
// Agent updates with the UpdateContainerAgent API operation do not apply to
// Windows container instances. We recommend that you launch new container
// instances to update the agent version in your Windows clusters.
//
// The UpdateContainerAgent API requires an Amazon ECS-optimized AMI or Amazon
// Linux AMI with the ecs-init service installed and running. For help updating
// the Amazon ECS container agent on other operating systems, see [Manually updating the Amazon ECS container agent]in the Amazon
// Elastic Container Service Developer Guide.
//
// [Updating the Amazon ECS container agent]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/agent-update-ecs-ami.html
// [Manually updating the Amazon ECS container agent]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-update.html#manually_update_agent
