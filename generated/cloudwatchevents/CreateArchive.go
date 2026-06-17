package cloudwatchevents

// CreateArchive is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatchevents.go.
//
// Creates an archive of events with the specified settings. When you create an
// archive, incoming events might not immediately start being sent to the archive.
// Allow a short period of time for changes to take effect. If you do not specify a
// pattern to filter events sent to the archive, all events are sent to the archive
// except replayed events. Replayed events are not sent to an archive.
