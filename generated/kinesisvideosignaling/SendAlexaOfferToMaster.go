package kinesisvideosignaling

// SendAlexaOfferToMaster is generated as a reference stub.
// Executable command wiring lives under cmd/kinesisvideosignaling.go.
//
// This API allows you to connect WebRTC-enabled devices with Alexa display
// devices. When invoked, it sends the Alexa Session Description Protocol (SDP)
// offer to the master peer. The offer is delivered as soon as the master is
// connected to the specified signaling channel. This API returns the SDP answer
// from the connected master. If the master is not connected to the signaling
// channel, redelivery requests are made until the message expires.
