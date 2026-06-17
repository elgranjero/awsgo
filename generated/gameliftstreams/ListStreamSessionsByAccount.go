package gameliftstreams

// ListStreamSessionsByAccount is generated as a reference stub.
// Executable command wiring lives under cmd/gameliftstreams.go.
//
// Retrieves a list of Amazon GameLift Streams stream sessions that this user
// account has access to.
//
// In the returned list of stream sessions, the ExportFilesMetadata property only
// shows the Status value. To get the OutpurUri and StatusReason values, use [GetStreamSession].
//
// We don't recommend using this operation to regularly check stream session
// statuses because it's costly. Instead, to check status updates for a specific
// stream session, use [GetStreamSession].
//
// [GetStreamSession]: https://docs.aws.amazon.com/gameliftstreams/latest/apireference/API_GetStreamSession.html
