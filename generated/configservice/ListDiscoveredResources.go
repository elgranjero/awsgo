package configservice

// ListDiscoveredResources is generated as a reference stub.
// Executable command wiring lives under cmd/configservice.go.
//
// Returns a list of resource resource identifiers for the specified resource
// types for the resources of that type. A resource identifier includes the
// resource type, ID, and (if available) the custom resource name.
//
// The results consist of resources that Config has discovered, including those
// that Config is not currently recording. You can narrow the results to include
// only resources that have specific resource IDs or a resource name.
//
// You can specify either resource IDs or a resource name, but not both, in the
// same request.
//
// # CloudFormation stack recording behavior in Config
//
// When a CloudFormation stack fails to create (for example, it enters the
// ROLLBACK_FAILED state), Config does not record a configuration item (CI) for
// that stack. Configuration items are only recorded for stacks that reach the
// following states:
//
// - CREATE_COMPLETE
//
// - UPDATE_COMPLETE
//
// - UPDATE_ROLLBACK_COMPLETE
//
// - UPDATE_ROLLBACK_FAILED
//
// - DELETE_FAILED
//
// - DELETE_COMPLETE
//
// Because no CI is created for a failed stack creation, you won't see
// configuration history for that stack in Config, even after the stack is deleted.
// This helps make sure that Config only tracks resources that were successfully
// provisioned.
