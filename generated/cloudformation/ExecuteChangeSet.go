package cloudformation

// ExecuteChangeSet is generated as a reference stub.
// Executable command wiring lives under cmd/cloudformation.go.
//
// Updates a stack using the input information that was provided when the
// specified change set was created. After the call successfully completes,
// CloudFormation starts updating the stack. Use the DescribeStacksaction to view the status of
// the update.
//
// When you execute a change set, CloudFormation deletes all other change sets
// associated with the stack because they aren't valid for the updated stack.
//
// If a stack policy is associated with the stack, CloudFormation enforces the
// policy during the update. You can't specify a temporary stack policy that
// overrides the current policy.
//
// To create a change set for the entire stack hierarchy, IncludeNestedStacks must
// have been set to True .
