package iotsecuretunneling

// RotateTunnelAccessToken is generated as a reference stub.
// Executable command wiring lives under cmd/iotsecuretunneling.go.
//
// Revokes the current client access token (CAT) and returns new CAT for clients
// to use when reconnecting to secure tunneling to access the same tunnel.
//
// Requires permission to access the [RotateTunnelAccessToken] action.
//
// Rotating the CAT doesn't extend the tunnel duration. For example, say the
// tunnel duration is 12 hours and the tunnel has already been open for 4 hours.
// When you rotate the access tokens, the new tokens that are generated can only be
// used for the remaining 8 hours.
//
// [RotateTunnelAccessToken]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
