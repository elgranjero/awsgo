package kinesis

// RemoveTagsFromStream is generated as a reference stub.
// Executable command wiring lives under cmd/kinesis.go.
//
// Removes tags from the specified Kinesis data stream. Removed tags are deleted
// and cannot be recovered after this operation successfully completes.
//
// When invoking this API, you must use either the StreamARN or the StreamName
// parameter, or both. It is recommended that you use the StreamARN input
// parameter when you invoke this API.
//
// If you specify a tag that does not exist, it is ignored.
//
// RemoveTagsFromStreamhas a limit of five transactions per second per account.
