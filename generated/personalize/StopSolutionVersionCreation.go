package personalize

// StopSolutionVersionCreation is generated as a reference stub.
// Executable command wiring lives under cmd/personalize.go.
//
// Stops creating a solution version that is in a state of CREATE_PENDING or
// CREATE IN_PROGRESS.
//
// Depending on the current state of the solution version, the solution version
// state changes as follows:
//
// - CREATE_PENDING > CREATE_STOPPED
//
// or
//
// - CREATE_IN_PROGRESS > CREATE_STOPPING > CREATE_STOPPED
//
// You are billed for all of the training completed up until you stop the solution
// version creation. You cannot resume creating a solution version once it has been
// stopped.
