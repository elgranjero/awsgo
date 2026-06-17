package snowdevicemanagement

// CancelTask is generated as a reference stub.
// Executable command wiring lives under cmd/snowdevicemanagement.go.
//
// Sends a cancel request for a specified task. You can cancel a task only if it's
// still in a QUEUED state. Tasks that are already running can't be cancelled.
//
// A task might still run if it's processed from the queue before the CancelTask
// operation changes the task's state.
