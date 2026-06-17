package athena

// StopCalculationExecution is generated as a reference stub.
// Executable command wiring lives under cmd/athena.go.
//
// Requests the cancellation of a calculation. A StopCalculationExecution call on
// a calculation that is already in a terminal state (for example, STOPPED , FAILED
// , or COMPLETED ) succeeds but has no effect.
//
// Cancelling a calculation is done on a best effort basis. If a calculation
// cannot be cancelled, you can be charged for its completion. If you are concerned
// about being charged for a calculation that cannot be cancelled, consider
// terminating the session in which the calculation is running.
