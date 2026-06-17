package gameliftstreams

// CreateStreamSessionConnection is generated as a reference stub.
// Executable command wiring lives under cmd/gameliftstreams.go.
//
// Enables clients to reconnect to a stream session while preserving all session
// state and data in the disconnected session. This reconnection process can be
// initiated when a stream session is in either PENDING_CLIENT_RECONNECTION or
// ACTIVE status. The process works as follows:
//
// - Initial disconnect:
//
// - When a client disconnects or loses connection, the stream session
// transitions from CONNECTED to PENDING_CLIENT_RECONNECTION
//
// - Reconnection time window:
//
// - Clients have ConnectionTimeoutSeconds (defined in [StartStreamSession]) to reconnect before
// session termination
//
// - Your backend server must call CreateStreamSessionConnection to initiate
// reconnection
//
// - Session transitions to RECONNECTING status
//
// - Reconnection completion:
//
// - On successful CreateStreamSessionConnection, session status changes to
// ACTIVE
//
// - Provide the new connection information to the requesting client
//
// - Client must establish connection within ConnectionTimeoutSeconds
//
// - Session terminates automatically if client fails to connect in time
//
// For more information about the stream session lifecycle, see [Stream sessions] in the Amazon
// GameLift Streams Developer Guide.
//
// To begin re-connecting to an existing stream session, specify the stream group
// ID and stream session ID that you want to reconnect to, and the signal request
// to use with the stream.
//
// [Stream sessions]: https://docs.aws.amazon.com/gameliftstreams/latest/developerguide/stream-sessions.html
// [StartStreamSession]: https://docs.aws.amazon.com/gameliftstreams/latest/apireference/API_StartStreamSession.html
