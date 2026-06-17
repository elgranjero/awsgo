package cloudformation

// SignalResource is generated as a reference stub.
// Executable command wiring lives under cmd/cloudformation.go.
//
// Sends a signal to the specified resource with a success or failure status. You
// can use the SignalResource operation in conjunction with a creation policy or
// update policy. CloudFormation doesn't proceed with a stack creation or update
// until resources receive the required number of signals or the timeout period is
// exceeded. The SignalResource operation is useful in cases where you want to
// send signals from anywhere other than an Amazon EC2 instance.
