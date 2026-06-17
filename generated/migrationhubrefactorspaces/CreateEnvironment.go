package migrationhubrefactorspaces

// CreateEnvironment is generated as a reference stub.
// Executable command wiring lives under cmd/migrationhubrefactorspaces.go.
//
// Creates an Amazon Web Services Migration Hub Refactor Spaces environment. The
// caller owns the environment resource, and all Refactor Spaces applications,
// services, and routes created within the environment. They are referred to as the
// environment owner. The environment owner has cross-account visibility and
// control of Refactor Spaces resources that are added to the environment by other
// accounts that the environment is shared with.
//
// When creating an environment with a [CreateEnvironment:NetworkFabricType] of TRANSIT_GATEWAY , Refactor Spaces
// provisions a transit gateway to enable services in VPCs to communicate directly
// across accounts. If [CreateEnvironment:NetworkFabricType]is NONE , Refactor Spaces does not create a transit gateway
// and you must use your network infrastructure to route traffic to services with
// private URL endpoints.
//
// [CreateEnvironment:NetworkFabricType]: https://docs.aws.amazon.com/migrationhub-refactor-spaces/latest/APIReference/API_CreateEnvironment.html#migrationhubrefactorspaces-CreateEnvironment-request-NetworkFabricType
