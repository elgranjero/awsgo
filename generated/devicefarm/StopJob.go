package devicefarm

// StopJob is generated as a reference stub.
// Executable command wiring lives under cmd/devicefarm.go.
//
// Initiates a stop request for the current job. AWS Device Farm immediately stops
// the job on the device where tests have not started. You are not billed for this
// device. On the device where tests have started, setup suite and teardown suite
// tests run to completion on the device. You are billed for setup, teardown, and
// any tests that were in progress or already completed.
