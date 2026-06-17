package kinesisvideoarchivedmedia

// GetMediaForFragmentList is generated as a reference stub.
// Executable command wiring lives under cmd/kinesisvideoarchivedmedia.go.
//
// Gets media for a list of fragments (specified by fragment number) from the
// archived data in an Amazon Kinesis video stream.
//
// You must first call the GetDataEndpoint API to get an endpoint. Then send the
// GetMediaForFragmentList requests to this endpoint using the [--endpoint-url parameter].
//
// For limits, see [Kinesis Video Streams Limits].
//
// If an error is thrown after invoking a Kinesis Video Streams archived media
// API, in addition to the HTTP status code and the response body, it includes the
// following pieces of information:
//
// - x-amz-ErrorType HTTP header – contains a more specific error type in
// addition to what the HTTP status code provides.
//
// - x-amz-RequestId HTTP header – if you want to report an issue to Amazon Web
// Services, the support team can better diagnose the problem if given the Request
// Id.
//
// Both the HTTP status code and the ErrorType header can be utilized to make
// programmatic decisions about whether errors are retry-able and under what
// conditions, as well as provide information on what actions the client programmer
// might need to take in order to successfully try again.
//
// For more information, see the Errors section at the bottom of this topic, as
// well as [Common Errors].
//
// [--endpoint-url parameter]: https://docs.aws.amazon.com/cli/latest/reference/
// [Common Errors]: https://docs.aws.amazon.com/kinesisvideostreams/latest/dg/CommonErrors.html
// [Kinesis Video Streams Limits]: http://docs.aws.amazon.com/kinesisvideostreams/latest/dg/limits.html
