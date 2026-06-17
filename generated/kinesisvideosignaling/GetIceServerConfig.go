package kinesisvideosignaling

// GetIceServerConfig is generated as a reference stub.
// Executable command wiring lives under cmd/kinesisvideosignaling.go.
//
// Gets the Interactive Connectivity Establishment (ICE) server configuration
// information, including URIs, username, and password which can be used to
// configure the WebRTC connection. The ICE component uses this configuration
// information to setup the WebRTC connection, including authenticating with the
// Traversal Using Relays around NAT (TURN) relay server.
//
// TURN is a protocol that is used to improve the connectivity of peer-to-peer
// applications. By providing a cloud-based relay service, TURN ensures that a
// connection can be established even when one or more peers are incapable of a
// direct peer-to-peer connection. For more information, see [A REST API For Access To TURN Services].
//
// You can invoke this API to establish a fallback mechanism in case either of the
// peers is unable to establish a direct peer-to-peer connection over a signaling
// channel. You must specify either a signaling channel ARN or the client ID in
// order to invoke this API.
//
// [A REST API For Access To TURN Services]: https://tools.ietf.org/html/draft-uberti-rtcweb-turn-rest-00
