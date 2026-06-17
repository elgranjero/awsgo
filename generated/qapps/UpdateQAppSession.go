package qapps

// UpdateQAppSession is generated as a reference stub.
// Executable command wiring lives under cmd/qapps.go.
//
// Updates the session for a given Q App sessionId . This is only valid when at
// least one card of the session is in the WAITING state. Data for each WAITING
// card can be provided as input. If inputs are not provided, the call will be
// accepted but session will not move forward. Inputs for cards that are not in the
// WAITING status will be ignored.
