package kinesisvideo

// GetSignalingChannelEndpoint is generated as a reference stub.
// Executable command wiring lives under cmd/kinesisvideo.go.
//
// Provides an endpoint for the specified signaling channel to send and receive
// messages. This API uses the SingleMasterChannelEndpointConfiguration input
// parameter, which consists of the Protocols and Role properties.
//
// Protocols is used to determine the communication mechanism. For example, if you
// specify WSS as the protocol, this API produces a secure websocket endpoint. If
// you specify HTTPS as the protocol, this API generates an HTTPS endpoint. If you
// specify WEBRTC as the protocol, but the signaling channel isn't configured for
// ingestion, you will receive the error InvalidArgumentException .
//
// Role determines the messaging permissions. A MASTER role results in this API
// generating an endpoint that a client can use to communicate with any of the
// viewers on the channel. A VIEWER role results in this API generating an
// endpoint that a client can use to communicate only with a MASTER .
