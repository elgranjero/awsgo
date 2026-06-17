package cloudformation

// RollbackStack is generated as a reference stub.
// Executable command wiring lives under cmd/cloudformation.go.
//
// When specifying RollbackStack , you preserve the state of previously provisioned
// resources when an operation fails. You can check the status of the stack through
// the DescribeStacksoperation.
//
// Rolls back the specified stack to the last known stable state from CREATE_FAILED
// or UPDATE_FAILED stack statuses.
//
// This operation will delete a stack if it doesn't contain a last known stable
// state. A last known stable state includes any status in a *_COMPLETE . This
// includes the following stack statuses.
//
// - CREATE_COMPLETE
//
// - UPDATE_COMPLETE
//
// - UPDATE_ROLLBACK_COMPLETE
//
// - IMPORT_COMPLETE
//
// - IMPORT_ROLLBACK_COMPLETE
