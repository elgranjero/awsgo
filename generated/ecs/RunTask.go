package ecs

// RunTask is generated as a reference stub.
// Executable command wiring lives under cmd/ecs.go.
//
// Starts a new task using the specified task definition.
//
// On March 21, 2024, a change was made to resolve the task definition revision
// before authorization. When a task definition revision is not specified,
// authorization will occur using the latest revision of a task definition.
//
// Amazon Elastic Inference (EI) is no longer available to customers.
//
// You can allow Amazon ECS to place tasks for you, or you can customize how
// Amazon ECS places tasks using placement constraints and placement strategies.
// For more information, see [Scheduling Tasks]in the Amazon Elastic Container Service Developer
// Guide.
//
// Alternatively, you can use StartTask to use your own scheduler or place tasks
// manually on specific container instances.
//
// You can attach Amazon EBS volumes to Amazon ECS tasks by configuring the volume
// when creating or updating a service. For more information, see [Amazon EBS volumes]in the Amazon
// Elastic Container Service Developer Guide.
//
// The Amazon ECS API follows an eventual consistency model. This is because of
// the distributed nature of the system supporting the API. This means that the
// result of an API command you run that affects your Amazon ECS resources might
// not be immediately visible to all subsequent commands you run. Keep this in mind
// when you carry out an API command that immediately follows a previous API
// command.
//
// To manage eventual consistency, you can do the following:
//
// - Confirm the state of the resource before you run a command to modify it.
// Run the DescribeTasks command using an exponential backoff algorithm to ensure
// that you allow enough time for the previous command to propagate through the
// system. To do this, run the DescribeTasks command repeatedly, starting with a
// couple of seconds of wait time and increasing gradually up to five minutes of
// wait time.
//
// - Add wait time between subsequent commands, even if the DescribeTasks
// command returns an accurate response. Apply an exponential backoff algorithm
// starting with a couple of seconds of wait time, and increase gradually up to
// about five minutes of wait time.
//
// If you get a ConflictException error, the RunTask request could not be
// processed due to conflicts. The provided clientToken is already in use with a
// different RunTask request. The resourceIds are the existing task ARNs which are
// already associated with the clientToken .
//
// To fix this issue:
//
// - Run RunTask with a unique clientToken .
//
// - Run RunTask with the clientToken and the original set of parameters
//
// If you get a ClientException error, the RunTask could not be processed because
// you use managed scaling and there is a capacity error because the quota of tasks
// in the PROVISIONING per cluster has been reached. For information about the
// service quotas, see [Amazon ECS service quotas].
//
// [Scheduling Tasks]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/scheduling_tasks.html
// [Amazon EBS volumes]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ebs-volumes.html#ebs-volume-types
// [Amazon ECS service quotas]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-quotas.html
