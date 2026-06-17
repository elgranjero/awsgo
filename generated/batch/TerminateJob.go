package batch

// TerminateJob is generated as a reference stub.
// Executable command wiring lives under cmd/batch.go.
//
// Terminates a job in a job queue. Jobs that are in the STARTING or RUNNING state
// are terminated, which causes them to transition to FAILED . Jobs that have not
// progressed to the STARTING state are cancelled.
