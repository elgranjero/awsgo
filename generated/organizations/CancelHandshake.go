package organizations

// CancelHandshake is generated as a reference stub.
// Executable command wiring lives under cmd/organizations.go.
//
// Cancels a Handshake.
//
// Only the account that sent a handshake can call this operation. The recipient
// of the handshake can't cancel it, but can use DeclineHandshaketo decline. After a handshake is
// canceled, the recipient can no longer respond to the handshake.
//
// You can view canceled handshakes in API responses for 30 days before they are
// deleted.
