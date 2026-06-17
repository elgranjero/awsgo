package kinesis

// AddTagsToStream is generated as a reference stub.
// Executable command wiring lives under cmd/kinesis.go.
//
// Adds or updates tags for the specified Kinesis data stream. You can assign up
// to 50 tags to a data stream.
//
// When invoking this API, you must use either the StreamARN or the StreamName
// parameter, or both. It is recommended that you use the StreamARN input
// parameter when you invoke this API.
//
// If tags have already been assigned to the stream, AddTagsToStream overwrites
// any existing tags that correspond to the specified tag keys.
//
// AddTagsToStreamhas a limit of five transactions per second per account.
