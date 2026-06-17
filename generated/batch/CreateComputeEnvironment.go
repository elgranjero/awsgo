package batch

// CreateComputeEnvironment is generated as a reference stub.
// Executable command wiring lives under cmd/batch.go.
//
// Creates an Batch compute environment. You can create MANAGED or UNMANAGED
// compute environments. MANAGED compute environments can use Amazon EC2 or
// Fargate resources. UNMANAGED compute environments can only use EC2 resources.
//
// In a managed compute environment, Batch manages the capacity and instance types
// of the compute resources within the environment. This is based on the compute
// resource specification that you define or the [launch template]that you specify when you create
// the compute environment. Either, you can choose to use EC2 On-Demand Instances
// and EC2 Spot Instances. Or, you can use Fargate and Fargate Spot capacity in
// your managed compute environment. You can optionally set a maximum price so that
// Spot Instances only launch when the Spot Instance price is less than a specified
// percentage of the On-Demand price.
//
// In an unmanaged compute environment, you can manage your own EC2 compute
// resources and have flexibility with how you configure your compute resources.
// For example, you can use custom AMIs. However, you must verify that each of your
// AMIs meet the Amazon ECS container instance AMI specification. For more
// information, see [container instance AMIs]in the Amazon Elastic Container Service Developer Guide. After
// you created your unmanaged compute environment, you can use the DescribeComputeEnvironmentsoperation to
// find the Amazon ECS cluster that's associated with it. Then, launch your
// container instances into that Amazon ECS cluster. For more information, see [Launching an Amazon ECS container instance]in
// the Amazon Elastic Container Service Developer Guide.
//
// Batch doesn't automatically upgrade the AMIs in a compute environment after
// it's created. For more information on how to update a compute environment's AMI,
// see [Updating compute environments]in the Batch User Guide.
//
// [Updating compute environments]: https://docs.aws.amazon.com/batch/latest/userguide/updating-compute-environments.html
// [launch template]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-launch-templates.html
// [Launching an Amazon ECS container instance]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_container_instance.html
// [container instance AMIs]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/container_instance_AMIs.html
