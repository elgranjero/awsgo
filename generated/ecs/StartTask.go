package ecs

// StartTask is generated as a reference stub.
// Executable command wiring lives under cmd/ecs.go.
//
// Starts a new task from the specified task definition on the specified container
// instance or instances.
//
// On March 21, 2024, a change was made to resolve the task definition revision
// before authorization. When a task definition revision is not specified,
// authorization will occur using the latest revision of a task definition.
//
// Amazon Elastic Inference (EI) is no longer available to customers.
//
// Alternatively, you can use RunTask to place tasks for you. For more
// information, see [Scheduling Tasks]in the Amazon Elastic Container Service Developer Guide.
//
// You can attach Amazon EBS volumes to Amazon ECS tasks by configuring the volume
// when creating or updating a service. For more information, see [Amazon EBS volumes]in the Amazon
// Elastic Container Service Developer Guide.
//
// [Scheduling Tasks]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/scheduling_tasks.html
// [Amazon EBS volumes]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ebs-volumes.html#ebs-volume-types
