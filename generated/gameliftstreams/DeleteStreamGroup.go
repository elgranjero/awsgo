package gameliftstreams

// DeleteStreamGroup is generated as a reference stub.
// Executable command wiring lives under cmd/gameliftstreams.go.
//
// Permanently deletes all compute resources and information related to a stream
// group. To delete a stream group, specify the unique stream group identifier.
// During the deletion process, the stream group's status is DELETING . This
// operation stops streams in progress and prevents new streams from starting. As a
// best practice, before deleting the stream group, call [ListStreamSessions]to check for streams in
// progress and take action to stop them. When you delete a stream group, any
// application associations referring to that stream group are automatically
// removed.
//
// [ListStreamSessions]: https://docs.aws.amazon.com/gameliftstreams/latest/apireference/API_ListStreamSessions.html
