package sagemaker

// CreateTrainingPlan is generated as a reference stub.
// Executable command wiring lives under cmd/sagemaker.go.
//
// Creates a new training plan in SageMaker to reserve compute capacity.
//
// Amazon SageMaker Training Plan is a capability within SageMaker that allows
// customers to reserve and manage GPU capacity for large-scale AI model training.
// It provides a way to secure predictable access to computational resources within
// specific timelines and budgets, without the need to manage underlying
// infrastructure.
//
// # How it works
//
// Plans can be created for specific resources such as SageMaker Training Jobs or
// SageMaker HyperPod clusters, automatically provisioning resources, setting up
// infrastructure, executing workloads, and handling infrastructure failures.
//
// Plan creation workflow
//
// - Users search for available plan offerings based on their requirements
// (e.g., instance type, count, start time, duration) using the [SearchTrainingPlanOfferings]API operation.
//
// - They create a plan that best matches their needs using the ID of the plan
// offering they want to use.
//
// - After successful upfront payment, the plan's status becomes Scheduled .
//
// - The plan can be used to:
//
// - Queue training jobs.
//
// - Allocate to an instance group of a SageMaker HyperPod cluster.
//
// - When the plan start date arrives, it becomes Active . Based on available
// reserved capacity:
//
// - Training jobs are launched.
//
// - Instance groups are provisioned.
//
// # Plan composition
//
// A plan can consist of one or more Reserved Capacities, each defined by a
// specific instance type, quantity, Availability Zone, duration, and start and end
// times. For more information about Reserved Capacity, see [ReservedCapacitySummary].
//
// [SearchTrainingPlanOfferings]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_SearchTrainingPlanOfferings.html
// [ReservedCapacitySummary]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_ReservedCapacitySummary.html
