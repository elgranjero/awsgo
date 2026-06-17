package transfer

// StopServer is generated as a reference stub.
// Executable command wiring lives under cmd/transfer.go.
//
// Changes the state of a file transfer protocol-enabled server from ONLINE to
// OFFLINE . An OFFLINE server cannot accept and process file transfer jobs.
// Information tied to your server, such as server and user properties, are not
// affected by stopping your server.
//
// Stopping the server does not reduce or impact your file transfer protocol
// endpoint billing; you must delete the server to stop being billed.
//
// The state of STOPPING indicates that the server is in an intermediate state,
// either not fully able to respond, or not fully offline. The values of
// STOP_FAILED can indicate an error condition.
//
// No response is returned from this call.
