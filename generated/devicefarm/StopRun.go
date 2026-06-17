package devicefarm

// StopRun is generated as a reference stub.
// Executable command wiring lives under cmd/devicefarm.go.
//
// Initiates a stop request for the current test run. AWS Device Farm immediately
// stops the run on devices where tests have not started. You are not billed for
// these devices. On devices where tests have started executing, setup suite and
// teardown suite tests run to completion on those devices. You are billed for
// setup, teardown, and any tests that were in progress or already completed.
