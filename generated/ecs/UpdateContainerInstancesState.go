package ecs

// UpdateContainerInstancesState is generated as a reference stub.
// Executable command wiring lives under cmd/ecs.go.
//
// Modifies the status of an Amazon ECS container instance.
//
// Once a container instance has reached an ACTIVE state, you can change the
// status of a container instance to DRAINING to manually remove an instance from
// a cluster, for example to perform system updates, update the Docker daemon, or
// scale down the cluster size.
//
// A container instance can't be changed to DRAINING until it has reached an ACTIVE
// status. If the instance is in any other status, an error will be received.
//
// When you set a container instance to DRAINING , Amazon ECS prevents new tasks
// from being scheduled for placement on the container instance and replacement
// service tasks are started on other container instances in the cluster if the
// resources are available. Service tasks on the container instance that are in the
// PENDING state are stopped immediately.
//
// Service tasks on the container instance that are in the RUNNING state are
// stopped and replaced according to the service's deployment configuration
// parameters, minimumHealthyPercent and maximumPercent . You can change the
// deployment configuration of your service using [UpdateService].
//
// - If minimumHealthyPercent is below 100%, the scheduler can ignore
// desiredCount temporarily during task replacement. For example, desiredCount is
// four tasks, a minimum of 50% allows the scheduler to stop two existing tasks
// before starting two new tasks. If the minimum is 100%, the service scheduler
// can't remove existing tasks until the replacement tasks are considered healthy.
// Tasks for services that do not use a load balancer are considered healthy if
// they're in the RUNNING state. Tasks for services that use a load balancer are
// considered healthy if they're in the RUNNING state and are reported as healthy
// by the load balancer.
//
// - The maximumPercent parameter represents an upper limit on the number of
// running tasks during task replacement. You can use this to define the
// replacement batch size. For example, if desiredCount is four tasks, a maximum
// of 200% starts four new tasks before stopping the four tasks to be drained,
// provided that the cluster resources required to do this are available. If the
// maximum is 100%, then replacement tasks can't start until the draining tasks
// have stopped.
//
// Any PENDING or RUNNING tasks that do not belong to a service aren't affected.
// You must wait for them to finish or stop them manually.
//
// A container instance has completed draining when it has no more RUNNING tasks.
// You can verify this using [ListTasks].
//
// When a container instance has been drained, you can set a container instance to
// ACTIVE status and once it has reached that status the Amazon ECS scheduler can
// begin scheduling tasks on the instance again.
//
// [UpdateService]: https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_UpdateService.html
// [ListTasks]: https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ListTasks.html
