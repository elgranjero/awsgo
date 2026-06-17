package ecs

// CreateService is generated as a reference stub.
// Executable command wiring lives under cmd/ecs.go.
//
// Runs and maintains your desired number of tasks from a specified task
// definition. If the number of tasks running in a service drops below the
// desiredCount , Amazon ECS runs another copy of the task in the specified
// cluster. To update an existing service, use [UpdateService].
//
// On March 21, 2024, a change was made to resolve the task definition revision
// before authorization. When a task definition revision is not specified,
// authorization will occur using the latest revision of a task definition.
//
// Amazon Elastic Inference (EI) is no longer available to customers.
//
// In addition to maintaining the desired count of tasks in your service, you can
// optionally run your service behind one or more load balancers. The load
// balancers distribute traffic across the tasks that are associated with the
// service. For more information, see [Service load balancing]in the Amazon Elastic Container Service
// Developer Guide.
//
// You can attach Amazon EBS volumes to Amazon ECS tasks by configuring the volume
// when creating or updating a service. volumeConfigurations is only supported for
// REPLICA service and not DAEMON service. For more information, see [Amazon EBS volumes]in the Amazon
// Elastic Container Service Developer Guide.
//
// Tasks for services that don't use a load balancer are considered healthy if
// they're in the RUNNING state. Tasks for services that use a load balancer are
// considered healthy if they're in the RUNNING state and are reported as healthy
// by the load balancer.
//
// There are two service scheduler strategies available:
//
// - REPLICA - The replica scheduling strategy places and maintains your desired
// number of tasks across your cluster. By default, the service scheduler spreads
// tasks across Availability Zones. You can use task placement strategies and
// constraints to customize task placement decisions. For more information, see [Service scheduler concepts]
// in the Amazon Elastic Container Service Developer Guide.
//
// - DAEMON - The daemon scheduling strategy deploys exactly one task on each
// active container instance that meets all of the task placement constraints that
// you specify in your cluster. The service scheduler also evaluates the task
// placement constraints for running tasks. It also stops tasks that don't meet the
// placement constraints. When using this strategy, you don't need to specify a
// desired number of tasks, a task placement strategy, or use Service Auto Scaling
// policies. For more information, see [Amazon ECS services]in the Amazon Elastic Container Service
// Developer Guide.
//
// The deployment controller is the mechanism that determines how tasks are
// deployed for your service. The valid options are:
//
// - ECS
//
// When you create a service which uses the ECS deployment controller, you can
//
// choose between the following deployment strategies (which you can set in the “
// strategy ” field in “ deploymentConfiguration ”): :
//
// - ROLLING : When you create a service which uses the rolling update ( ROLLING
// ) deployment strategy, the Amazon ECS service scheduler replaces the currently
// running tasks with new tasks. The number of tasks that Amazon ECS adds or
// removes from the service during a rolling update is controlled by the service
// deployment configuration. For more information, see [Deploy Amazon ECS services by replacing tasks]in the Amazon Elastic
// Container Service Developer Guide.
//
// Rolling update deployments are best suited for the following scenarios:
//
// - Gradual service updates: You need to update your service incrementally
// without taking the entire service offline at once.
//
// - Limited resource requirements: You want to avoid the additional resource
// costs of running two complete environments simultaneously (as required by
// blue/green deployments).
//
// - Acceptable deployment time: Your application can tolerate a longer
// deployment process, as rolling updates replace tasks one by one.
//
// - No need for instant roll back: Your service can tolerate a rollback process
// that takes minutes rather than seconds.
//
// - Simple deployment process: You prefer a straightforward deployment approach
// without the complexity of managing multiple environments, target groups, and
// listeners.
//
// - No load balancer requirement: Your service doesn't use or require a load
// balancer, Application Load Balancer, Network Load Balancer, or Service Connect
// (which are required for blue/green deployments).
//
// - Stateful applications: Your application maintains state that makes it
// difficult to run two parallel environments.
//
// - Cost sensitivity: You want to minimize deployment costs by not running
// duplicate environments during deployment.
//
// Rolling updates are the default deployment strategy for services and provide a
//
// balance between deployment safety and resource efficiency for many common
// application scenarios.
//
// - BLUE_GREEN : A blue/green deployment strategy ( BLUE_GREEN ) is a release
// methodology that reduces downtime and risk by running two identical production
// environments called blue and green. With Amazon ECS blue/green deployments, you
// can validate new service revisions before directing production traffic to them.
// This approach provides a safer way to deploy changes with the ability to quickly
// roll back if needed. For more information, see [Amazon ECS blue/green deployments]in the Amazon Elastic
// Container Service Developer Guide.
//
// Amazon ECS blue/green deployments are best suited for the following scenarios:
//
// - Service validation: When you need to validate new service revisions before
// directing production traffic to them
//
// - Zero downtime: When your service requires zero-downtime deployments
//
// - Instant roll back: When you need the ability to quickly roll back if issues
// are detected
//
// - Load balancer requirement: When your service uses Application Load
// Balancer, Network Load Balancer, or Service Connect
//
// - LINEAR : A linear deployment strategy ( LINEAR ) gradually shifts traffic
// from the current production environment to a new environment in equal percentage
// increments. With Amazon ECS linear deployments, you can control the pace of
// traffic shifting and validate new service revisions with increasing amounts of
// production traffic.
//
// Linear deployments are best suited for the following scenarios:
//
// - Gradual validation: When you want to gradually validate your new service
// version with increasing traffic
//
// - Performance monitoring: When you need time to monitor metrics and
// performance during the deployment
//
// - Risk minimization: When you want to minimize risk by exposing the new
// version to production traffic incrementally
//
// - Load balancer requirement: When your service uses Application Load Balancer
// or Service Connect
//
// - CANARY : A canary deployment strategy ( CANARY ) shifts a small percentage
// of traffic to the new service revision first, then shifts the remaining traffic
// all at once after a specified time period. This allows you to test the new
// version with a subset of users before full deployment.
//
// Canary deployments are best suited for the following scenarios:
//
// - Feature testing: When you want to test new features with a small subset of
// users before full rollout
//
// - Production validation: When you need to validate performance and
// functionality with real production traffic
//
// - Blast radius control: When you want to minimize blast radius if issues are
// discovered in the new version
//
// - Load balancer requirement: When your service uses Application Load Balancer
// or Service Connect
//
// - External
//
// Use a third-party deployment controller.
//
// - Blue/green deployment (powered by CodeDeploy)
//
// CodeDeploy installs an updated version of the application as a new replacement
//
// task set and reroutes production traffic from the original application task set
// to the replacement task set. The original task set is terminated after a
// successful deployment. Use this deployment controller to verify a new deployment
// of a service before sending production traffic to it.
//
// When creating a service that uses the EXTERNAL deployment controller, you can
// specify only parameters that aren't controlled at the task set level. The only
// required parameter is the service name. You control your services using the [CreateTaskSet].
// For more information, see [Amazon ECS deployment types]in the Amazon Elastic Container Service Developer
// Guide.
//
// When the service scheduler launches new tasks, it determines task placement.
// For information about task placement and task placement strategies, see [Amazon ECS task placement]in the
// Amazon Elastic Container Service Developer Guide
//
// [Amazon ECS task placement]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-placement.html
// [Service scheduler concepts]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs_services.html
// [Amazon ECS deployment types]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html
// [UpdateService]: https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_UpdateService.html
// [CreateTaskSet]: https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_CreateTaskSet.html
// [Amazon ECS services]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs_services.html
// [Service load balancing]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-load-balancing.html
// [Amazon EBS volumes]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ebs-volumes.html#ebs-volume-types
//
// [Amazon ECS blue/green deployments]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-type-blue-green.html
// [Deploy Amazon ECS services by replacing tasks]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-type-ecs.html
