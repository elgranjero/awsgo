package iotsecuretunneling

// CloseTunnel is generated as a reference stub.
// Executable command wiring lives under cmd/iotsecuretunneling.go.
//
// Closes a tunnel identified by the unique tunnel id. When a CloseTunnel request
// is received, we close the WebSocket connections between the client and proxy
// server so no data can be transmitted.
//
// Requires permission to access the [CloseTunnel] action.
//
// [CloseTunnel]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
