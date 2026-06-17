package ecs

// UpdateService is generated as a reference stub.
// Executable command wiring lives under cmd/ecs.go.
//
// Modifies the parameters of a service.
//
// On March 21, 2024, a change was made to resolve the task definition revision
// before authorization. When a task definition revision is not specified,
// authorization will occur using the latest revision of a task definition.
//
// For services using the rolling update ( ECS ) you can update the desired count,
// deployment configuration, network configuration, load balancers, service
// registries, enable ECS managed tags option, propagate tags option, task
// placement constraints and strategies, and task definition. When you update any
// of these parameters, Amazon ECS starts new tasks with the new configuration.
//
// You can attach Amazon EBS volumes to Amazon ECS tasks by configuring the volume
// when starting or running a task, or when creating or updating a service. For
// more information, see [Amazon EBS volumes]in the Amazon Elastic Container Service Developer Guide.
// You can update your volume configurations and trigger a new deployment.
// volumeConfigurations is only supported for REPLICA service and not DAEMON
// service. If you leave volumeConfigurations null , it doesn't trigger a new
// deployment. For more information on volumes, see [Amazon EBS volumes]in the Amazon Elastic
// Container Service Developer Guide.
//
// For services using the blue/green ( CODE_DEPLOY ) deployment controller, only
// the desired count, deployment configuration, health check grace period, task
// placement constraints and strategies, enable ECS managed tags option, and
// propagate tags can be updated using this API. If the network configuration,
// platform version, task definition, or load balancer need to be updated, create a
// new CodeDeploy deployment. For more information, see [CreateDeployment]in the CodeDeploy API
// Reference.
//
// For services using an external deployment controller, you can update only the
// desired count, task placement constraints and strategies, health check grace
// period, enable ECS managed tags option, and propagate tags option, using this
// API. If the launch type, load balancer, network configuration, platform version,
// or task definition need to be updated, create a new task set For more
// information, see [CreateTaskSet].
//
// You can add to or subtract from the number of instantiations of a task
// definition in a service by specifying the cluster that the service is running in
// and a new desiredCount parameter.
//
// You can attach Amazon EBS volumes to Amazon ECS tasks by configuring the volume
// when starting or running a task, or when creating or updating a service. For
// more information, see [Amazon EBS volumes]in the Amazon Elastic Container Service Developer Guide.
//
// If you have updated the container image of your application, you can create a
// new task definition with that image and deploy it to your service. The service
// scheduler uses the minimum healthy percent and maximum percent parameters (in
// the service's deployment configuration) to determine the deployment strategy.
//
// If your updated Docker image uses the same tag as what is in the existing task
// definition for your service (for example, my_image:latest ), you don't need to
// create a new revision of your task definition. You can update the service using
// the forceNewDeployment option. The new tasks launched by the deployment pull
// the current image/tag combination from your repository when they start.
//
// You can also update the deployment configuration of a service. When a
// deployment is triggered by updating the task definition of a service, the
// service scheduler uses the deployment configuration parameters,
// minimumHealthyPercent and maximumPercent , to determine the deployment strategy.
//
// - If minimumHealthyPercent is below 100%, the scheduler can ignore
// desiredCount temporarily during a deployment. For example, if desiredCount is
// four tasks, a minimum of 50% allows the scheduler to stop two existing tasks
// before starting two new tasks. Tasks for services that don't use a load balancer
// are considered healthy if they're in the RUNNING state. Tasks for services
// that use a load balancer are considered healthy if they're in the RUNNING
// state and are reported as healthy by the load balancer.
//
// - The maximumPercent parameter represents an upper limit on the number of
// running tasks during a deployment. You can use it to define the deployment batch
// size. For example, if desiredCount is four tasks, a maximum of 200% starts
// four new tasks before stopping the four older tasks (provided that the cluster
// resources required to do this are available).
//
// When [UpdateService] stops a task during a deployment, the equivalent of docker stop is issued
// to the containers running in the task. This results in a SIGTERM and a
// 30-second timeout. After this, SIGKILL is sent and the containers are forcibly
// stopped. If the container handles the SIGTERM gracefully and exits within 30
// seconds from receiving it, no SIGKILL is sent.
//
// When the service scheduler launches new tasks, it determines task placement in
// your cluster with the following logic.
//
// - Determine which of the container instances in your cluster can support your
// service's task definition. For example, they have the required CPU, memory,
// ports, and container instance attributes.
//
// - By default, the service scheduler attempts to balance tasks across
// Availability Zones in this manner even though you can choose a different
// placement strategy.
//
// - Sort the valid container instances by the fewest number of running tasks
// for this service in the same Availability Zone as the instance. For example, if
// zone A has one running service task and zones B and C each have zero, valid
// container instances in either zone B or C are considered optimal for placement.
//
// - Place the new service task on a valid container instance in an optimal
// Availability Zone (based on the previous steps), favoring container instances
// with the fewest number of running tasks for this service.
//
// When the service scheduler stops running tasks, it attempts to maintain balance
// across the Availability Zones in your cluster using the following logic:
//
// - Sort the container instances by the largest number of running tasks for
// this service in the same Availability Zone as the instance. For example, if zone
// A has one running service task and zones B and C each have two, container
// instances in either zone B or C are considered optimal for termination.
//
// - Stop the task on a container instance in an optimal Availability Zone
// (based on the previous steps), favoring container instances with the largest
// number of running tasks for this service.
//
// [CreateTaskSet]: https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_CreateTaskSet.html
// [UpdateService]: https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_UpdateService.html
// [Amazon EBS volumes]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ebs-volumes.html#ebs-volume-types
// [CreateDeployment]: https://docs.aws.amazon.com/codedeploy/latest/APIReference/API_CreateDeployment.html
