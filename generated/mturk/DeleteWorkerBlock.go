package mturk

// DeleteWorkerBlock is generated as a reference stub.
// Executable command wiring lives under cmd/mturk.go.
//
// The DeleteWorkerBlock operation allows you to reinstate a blocked Worker to
// work on your HITs. This operation reverses the effects of the CreateWorkerBlock
// operation. You need the Worker ID to use this operation. If the Worker ID is
// missing or invalid, this operation fails and returns the message “WorkerId is
// invalid.” If the specified Worker is not blocked, this operation returns
// successfully.
