package eventbridge

// CreateArchive is generated as a reference stub.
// Executable command wiring lives under cmd/eventbridge.go.
//
// Creates an archive of events with the specified settings. When you create an
// archive, incoming events might not immediately start being sent to the archive.
// Allow a short period of time for changes to take effect. If you do not specify a
// pattern to filter events sent to the archive, all events are sent to the archive
// except replayed events. Replayed events are not sent to an archive.
//
// If you have specified that EventBridge use a customer managed key for
// encrypting the source event bus, we strongly recommend you also specify a
// customer managed key for any archives for the event bus as well.
//
// For more information, see [Encrypting archives] in the Amazon EventBridge User Guide.
//
// [Encrypting archives]: https://docs.aws.amazon.com/eventbridge/latest/userguide/encryption-archives.html
