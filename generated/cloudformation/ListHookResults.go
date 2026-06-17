package cloudformation

// ListHookResults is generated as a reference stub.
// Executable command wiring lives under cmd/cloudformation.go.
//
// Returns summaries of invoked Hooks. For more information, see [View invocation summaries for CloudFormation Hooks] in the
// CloudFormation Hooks User Guide.
//
// This operation supports the following parameter combinations:
//
// - No parameters: Returns all Hook invocation summaries.
//
// - TypeArn only: Returns summaries for a specific Hook.
//
// - TypeArn and Status : Returns summaries for a specific Hook filtered by
// status.
//
// - TargetId and TargetType : Returns summaries for a specific Hook invocation
// target.
//
// [View invocation summaries for CloudFormation Hooks]: https://docs.aws.amazon.com/cloudformation-cli/latest/hooks-userguide/hooks-view-invocations.html
