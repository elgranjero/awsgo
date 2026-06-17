package athena

// TerminateSession is generated as a reference stub.
// Executable command wiring lives under cmd/athena.go.
//
// Terminates an active session. A TerminateSession call on a session that is
// already inactive (for example, in a FAILED , TERMINATED or TERMINATING state)
// succeeds but has no effect. Calculations running in the session when
// TerminateSession is called are forcefully stopped, but may display as FAILED
// instead of STOPPED .
