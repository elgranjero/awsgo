package cloudformation

// CreateChangeSet is generated as a reference stub.
// Executable command wiring lives under cmd/cloudformation.go.
//
// Creates a list of changes that will be applied to a stack so that you can
// review the changes before executing them. You can create a change set for a
// stack that doesn't exist or an existing stack. If you create a change set for a
// stack that doesn't exist, the change set shows all of the resources that
// CloudFormation will create. If you create a change set for an existing stack,
// CloudFormation compares the stack's information with the information that you
// submit in the change set and lists the differences. Use change sets to
// understand which resources CloudFormation will create or change, and how it will
// change resources in an existing stack, before you create or update a stack.
//
// To create a change set for a stack that doesn't exist, for the ChangeSetType
// parameter, specify CREATE . To create a change set for an existing stack,
// specify UPDATE for the ChangeSetType parameter. To create a change set for an
// import operation, specify IMPORT for the ChangeSetType parameter. After the
// CreateChangeSet call successfully completes, CloudFormation starts creating the
// change set. To check the status of the change set or to review it, use the DescribeChangeSet
// action.
//
// When you are satisfied with the changes the change set will make, execute the
// change set by using the ExecuteChangeSetaction. CloudFormation doesn't make changes until you
// execute the change set.
//
// To create a change set for the entire stack hierarchy, set IncludeNestedStacks
// to True .
