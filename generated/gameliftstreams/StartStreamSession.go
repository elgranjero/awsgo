package gameliftstreams

// StartStreamSession is generated as a reference stub.
// Executable command wiring lives under cmd/gameliftstreams.go.
//
// This action initiates a new stream session and outputs connection information
//
// that clients can use to access the stream. A stream session refers to an
// instance of a stream that Amazon GameLift Streams transmits from the server to
// the end-user. A stream session runs on a compute resource that a stream group
// has allocated. The start stream session process works as follows:
//
// - Prerequisites:
//
// - You must have a stream group in ACTIVE status
//
// - You must have idle or on-demand capacity in a stream group in the location
// you want to stream from
//
// - You must have at least one application associated to the stream group (use [AssociateApplications]
// if needed)
//
// - Start stream request:
//
// - Your backend server calls StartStreamSession to initiate connection
//
// - Amazon GameLift Streams creates the stream session resource, assigns an
// Amazon Resource Name (ARN) value, and begins searching for available stream
// capacity to run the stream
//
// - Session transitions to ACTIVATING status
//
// - Placement completion:
//
// - If Amazon GameLift Streams is successful in finding capacity for the
// stream, the stream session status changes to ACTIVE status and
// StartStreamSession returns stream connection information
//
// - If Amazon GameLift Streams was not successful in finding capacity within
// the placement timeout period (defined according to the capacity type and
// platform type), the stream session status changes to ERROR status and
// StartStreamSession returns a StatusReason of placementTimeout
//
// - Connection completion:
//
// - Provide the new connection information to the requesting client
//
// - Client must establish connection within ConnectionTimeoutSeconds (specified
// in StartStreamSession parameters)
//
// - Session terminates automatically if client fails to connect in time
//
// For more information about the stream session lifecycle, see [Stream sessions] in the Amazon
// GameLift Streams Developer Guide.
//
// Timeouts to be aware of that affect a stream session:
//
// - Placement timeout: The amount of time that Amazon GameLift Streams has to
// find capacity for a stream request. Placement timeout varies based on the
// capacity type used to fulfill your stream request:
//
// - Always-on capacity: 75 seconds
//
// - On-demand capacity:
//
// - Linux/Proton runtimes: 90 seconds
//
// - Windows runtime: 10 minutes
//
// - Connection timeout: The amount of time that Amazon GameLift Streams waits
// for a client to connect to a stream session in ACTIVE status, or reconnect to
// a stream session in PENDING_CLIENT_RECONNECTION status, the latter of which
// occurs when a client disconnects or loses connection from a stream session. If
// no client connects before the timeout, Amazon GameLift Streams terminates the
// stream session. This value is specified by ConnectionTimeoutSeconds in the
// StartStreamSession parameters.
//
// - Idle timeout: A stream session will be terminated if no user input has been
// received for 60 minutes.
//
// - Maximum session length: A stream session will be terminated after this
// amount of time has elapsed since it started, regardless of any existing client
// connections. This value is specified by SessionLengthSeconds in the
// StartStreamSession parameters.
//
// To start a new stream session, specify a stream group ID and application ID,
// along with the transport protocol and signal request to use with the stream
// session.
//
// For stream groups that have multiple locations, provide a set of locations
// ordered by priority using a Locations parameter. Amazon GameLift Streams will
// start a single stream session in the next available location. An application
// must be finished replicating to a remote location before the remote location can
// host a stream.
//
// To reconnect to a stream session after a client disconnects or loses
// connection, use [CreateStreamSessionConnection].
//
// [Stream sessions]: https://docs.aws.amazon.com/gameliftstreams/latest/developerguide/stream-sessions.html
// [AssociateApplications]: https://docs.aws.amazon.com/gameliftstreams/latest/apireference/API_AssociateApplications.html
// [CreateStreamSessionConnection]: https://docs.aws.amazon.com/gameliftstreams/latest/apireference/API_CreateStreamSessionConnection.html
