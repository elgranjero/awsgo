package cloudformation

// DescribeEvents is generated as a reference stub.
// Executable command wiring lives under cmd/cloudformation.go.
//
// Returns CloudFormation events based on flexible query criteria. Groups events
// by operation ID, enabling you to focus on individual stack operations during
// deployment.
//
// An operation is any action performed on a stack, including stack lifecycle
// actions (Create, Update, Delete, Rollback), change set creation, nested stack
// creation, and automatic rollbacks triggered by failures. Each operation has a
// unique identifier (Operation ID) and represents a discrete change attempt on the
// stack.
//
// Returns different types of events including:
//
// - Progress events - Status updates during stack operation execution.
//
// - Validation errors - Failures from CloudFormation Early Validations.
//
// - Provisioning errors - Resource creation and update failures.
//
// - Hook invocation errors - Failures from CloudFormation Hook during stack
// operations.
//
// One of ChangeSetName , OperationId or StackName must be specified as input.
