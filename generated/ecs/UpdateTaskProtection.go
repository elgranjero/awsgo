package ecs

// UpdateTaskProtection is generated as a reference stub.
// Executable command wiring lives under cmd/ecs.go.
//
// Updates the protection status of a task. You can set protectionEnabled to true
// to protect your task from termination during scale-in events from [Service Autoscaling]or [deployments].
//
// Task-protection, by default, expires after 2 hours at which point Amazon ECS
// clears the protectionEnabled property making the task eligible for termination
// by a subsequent scale-in event.
//
// You can specify a custom expiration period for task protection from 1 minute to
// up to 2,880 minutes (48 hours). To specify the custom expiration period, set the
// expiresInMinutes property. The expiresInMinutes property is always reset when
// you invoke this operation for a task that already has protectionEnabled set to
// true . You can keep extending the protection expiration period of a task by
// invoking this operation repeatedly.
//
// To learn more about Amazon ECS task protection, see [Task scale-in protection] in the Amazon Elastic
// Container Service Developer Guide .
//
// This operation is only supported for tasks belonging to an Amazon ECS service.
// Invoking this operation for a standalone task will result in an TASK_NOT_VALID
// failure. For more information, see [API failure reasons].
//
// If you prefer to set task protection from within the container, we recommend
// using the [Task scale-in protection endpoint].
//
// [deployments]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html
// [API failure reasons]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/api_failures_messages.html
// [Task scale-in protection endpoint]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-scale-in-protection-endpoint.html
// [Task scale-in protection]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-scale-in-protection.html
// [Service Autoscaling]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-auto-scaling.html
