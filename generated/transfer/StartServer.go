package transfer

// StartServer is generated as a reference stub.
// Executable command wiring lives under cmd/transfer.go.
//
// Changes the state of a file transfer protocol-enabled server from OFFLINE to
// ONLINE . It has no impact on a server that is already ONLINE . An ONLINE server
// can accept and process file transfer jobs.
//
// The state of STARTING indicates that the server is in an intermediate state,
// either not fully able to respond, or not fully online. The values of
// START_FAILED can indicate an error condition.
//
// No response is returned from this call.
