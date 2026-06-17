package cloudformation

// DeleteChangeSet is generated as a reference stub.
// Executable command wiring lives under cmd/cloudformation.go.
//
// Deletes the specified change set. Deleting change sets ensures that no one
// executes the wrong change set.
//
// If the call successfully completes, CloudFormation successfully deleted the
// change set.
//
// If IncludeNestedStacks specifies True during the creation of the nested change
// set, then DeleteChangeSet will delete all change sets that belong to the stacks
// hierarchy and will also delete all change sets for nested stacks with the status
// of REVIEW_IN_PROGRESS .
