package iotwireless

// DeleteWirelessGateway is generated as a reference stub.
// Executable command wiring lives under cmd/iotwireless.go.
//
// Deletes a wireless gateway.
//
// When deleting a wireless gateway, you might run into duplication errors for the
// following reasons.
//
// - If you specify a GatewayEui value that already exists.
//
// - If you used a ClientRequestToken with the same parameters within the last 10
// minutes.
//
// To avoid this error, make sure that you use unique identifiers and parameters
// for each request within the specified time period.
