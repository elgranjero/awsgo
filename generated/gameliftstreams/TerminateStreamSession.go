package gameliftstreams

// TerminateStreamSession is generated as a reference stub.
// Executable command wiring lives under cmd/gameliftstreams.go.
//
// Permanently terminates an active stream session. When called, the stream
// session status changes to TERMINATING . You can terminate a stream session in
// any status except ACTIVATING . If the stream session is in ACTIVATING status,
// an exception is thrown.
