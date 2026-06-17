package ecs

// DeleteTaskDefinitions is generated as a reference stub.
// Executable command wiring lives under cmd/ecs.go.
//
// Deletes one or more task definitions.
//
// You must deregister a task definition revision before you delete it. For more
// information, see [DeregisterTaskDefinition].
//
// When you delete a task definition revision, it is immediately transitions from
// the INACTIVE to DELETE_IN_PROGRESS . Existing tasks and services that reference
// a DELETE_IN_PROGRESS task definition revision continue to run without
// disruption. Existing services that reference a DELETE_IN_PROGRESS task
// definition revision can still scale up or down by modifying the service's
// desired count.
//
// You can't use a DELETE_IN_PROGRESS task definition revision to run new tasks or
// create new services. You also can't update an existing service to reference a
// DELETE_IN_PROGRESS task definition revision.
//
// A task definition revision will stay in DELETE_IN_PROGRESS status until all the
// associated tasks and services have been terminated.
//
// When you delete all INACTIVE task definition revisions, the task definition
// name is not displayed in the console and not returned in the API. If a task
// definition revisions are in the DELETE_IN_PROGRESS state, the task definition
// name is displayed in the console and returned in the API. The task definition
// name is retained by Amazon ECS and the revision is incremented the next time you
// create a task definition with that name.
//
// [DeregisterTaskDefinition]: https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_DeregisterTaskDefinition.html
