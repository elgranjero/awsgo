package kinesisvideo

// DeleteStream is generated as a reference stub.
// Executable command wiring lives under cmd/kinesisvideo.go.
//
// Deletes a Kinesis video stream and the data contained in the stream.
//
// This method marks the stream for deletion, and makes the data in the stream
// inaccessible immediately.
//
// To ensure that you have the latest version of the stream before deleting it,
// you can specify the stream version. Kinesis Video Streams assigns a version to
// each stream. When you update a stream, Kinesis Video Streams assigns a new
// version number. To get the latest stream version, use the DescribeStream API.
//
// This operation requires permission for the KinesisVideo:DeleteStream action.
