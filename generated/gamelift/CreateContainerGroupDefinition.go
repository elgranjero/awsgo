package gamelift

// CreateContainerGroupDefinition is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: Container
//
// Creates a ContainerGroupDefinition that describes a set of containers for
// hosting your game server with Amazon GameLift Servers managed containers
// hosting. An Amazon GameLift Servers container group is similar to a container
// task or pod. Use container group definitions when you create a container fleet
// with [CreateContainerFleet].
//
// A container group definition determines how Amazon GameLift Servers deploys
// your containers to each instance in a container fleet. You can maintain multiple
// versions of a container group definition.
//
// There are two types of container groups:
//
// - A game server container group has the containers that run your game server
// application and supporting software. A game server container group can have
// these container types:
//
// - Game server container. This container runs your game server. You can define
// one game server container in a game server container group.
//
// - Support container. This container runs software in parallel with your game
// server. You can define up to 8 support containers in a game server group.
//
// When building a game server container group definition, you can choose to
//
// bundle your game server executable and all dependent software into a single game
// server container. Alternatively, you can separate the software into one game
// server container and one or more support containers.
//
// On a container fleet instance, a game server container group can be deployed
//
// multiple times (depending on the compute resources of the instance). This means
// that all containers in the container group are replicated together.
//
// - A per-instance container group has containers for processes that aren't
// replicated on a container fleet instance. This might include background
// services, logging, test processes, or processes that need to persist
// independently of the game server container group. When building a per-instance
// container group, you can define up to 10 support containers.
//
// This operation requires Identity and Access Management (IAM) permissions to
// access container images in Amazon ECR repositories. See [IAM permissions for Amazon GameLift Servers]for help setting the
// appropriate permissions.
//
// # Request options
//
// Use this operation to make the following types of requests. You can specify
// values for the minimum required parameters and customize optional values later.
//
// - Create a game server container group definition. Provide the following
// required parameter values:
//
// - Name
//
// - ContainerGroupType ( GAME_SERVER )
//
// - OperatingSystem (omit to use default value)
//
// - TotalMemoryLimitMebibytes (omit to use default value)
//
// - TotalVcpuLimit (omit to use default value)
//
// - At least one GameServerContainerDefinition
//
// - ContainerName
//
// - ImageUrl
//
// - PortConfiguration
//
// - ServerSdkVersion (omit to use default value)
//
// - Create a per-instance container group definition. Provide the following
// required parameter values:
//
// - Name
//
// - ContainerGroupType ( PER_INSTANCE )
//
// - OperatingSystem (omit to use default value)
//
// - TotalMemoryLimitMebibytes (omit to use default value)
//
// - TotalVcpuLimit (omit to use default value)
//
// - At least one SupportContainerDefinition
//
// - ContainerName
//
// - ImageUrl
//
// # Results
//
// If successful, this request creates a ContainerGroupDefinition resource and
// assigns a unique ARN value. You can update most properties of a container group
// definition by calling [UpdateContainerGroupDefinition], and optionally save the update as a new version.
//
// [UpdateContainerGroupDefinition]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_UpdateContainerGroupDefinition.html
// [CreateContainerFleet]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_CreateContainerFleet.html
// [IAM permissions for Amazon GameLift Servers]: https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-iam-policy-examples.html
