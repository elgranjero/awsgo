package storagegateway

// UpdateChapCredentials is generated as a reference stub.
// Executable command wiring lives under cmd/storagegateway.go.
//
// Updates the Challenge-Handshake Authentication Protocol (CHAP) credentials for
// a specified iSCSI target. By default, a gateway does not have CHAP enabled;
// however, for added security, you might use it. This operation is supported in
// the volume and tape gateway types.
//
// When you update CHAP credentials, all existing connections on the target are
// closed and initiators must reconnect with the new credentials.
