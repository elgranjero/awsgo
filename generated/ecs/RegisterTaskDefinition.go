package ecs

// RegisterTaskDefinition is generated as a reference stub.
// Executable command wiring lives under cmd/ecs.go.
//
// Registers a new task definition from the supplied family and
// containerDefinitions . Optionally, you can add data volumes to your containers
// with the volumes parameter. For more information about task definition
// parameters and defaults, see [Amazon ECS Task Definitions]in the Amazon Elastic Container Service Developer
// Guide.
//
// You can specify a role for your task with the taskRoleArn parameter. When you
// specify a role for a task, its containers can then use the latest versions of
// the CLI or SDKs to make API requests to the Amazon Web Services services that
// are specified in the policy that's associated with the role. For more
// information, see [IAM Roles for Tasks]in the Amazon Elastic Container Service Developer Guide.
//
// You can specify a Docker networking mode for the containers in your task
// definition with the networkMode parameter. If you specify the awsvpc network
// mode, the task is allocated an elastic network interface, and you must specify a
// [NetworkConfiguration]when you create a service or run a task with the task definition. For more
// information, see [Task Networking]in the Amazon Elastic Container Service Developer Guide.
//
// [Amazon ECS Task Definitions]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_defintions.html
// [Task Networking]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-networking.html
// [IAM Roles for Tasks]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-iam-roles.html
// [NetworkConfiguration]: https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_NetworkConfiguration.html
