package cloudwatchlogs

// StartLiveTail is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatchlogs.go.
//
// Starts a Live Tail streaming session for one or more log groups. A Live Tail
// session returns a stream of log events that have been recently ingested in the
// log groups. For more information, see [Use Live Tail to view logs in near real time].
//
// The response to this operation is a response stream, over which the server
// sends live log events and the client receives them.
//
// The following objects are sent over the stream:
//
// - A single [LiveTailSessionStart]object is sent at the start of the session.
//
// - Every second, a [LiveTailSessionUpdate]object is sent. Each of these objects contains an array of
// the actual log events.
//
// If no new log events were ingested in the past second, the LiveTailSessionUpdate
//
// object will contain an empty array.
//
// The array of log events contained in a LiveTailSessionUpdate can include as many
//
// as 500 log events. If the number of log events matching the request exceeds 500
// per second, the log events are sampled down to 500 log events to be included in
// each LiveTailSessionUpdate object.
//
// If your client consumes the log events slower than the server produces them,
//
// CloudWatch Logs buffers up to 10 LiveTailSessionUpdate events or 5000 log
// events, after which it starts dropping the oldest events.
//
// - A [SessionStreamingException]object is returned if an unknown error occurs on the server side.
//
// - A [SessionTimeoutException]object is returned when the session times out, after it has been kept
// open for three hours.
//
// The StartLiveTail API routes requests using SDK host prefix injection. SDK
// versions released before April 1, 2026 route to
// streaming-logs.Region.amazonaws.com , which does not support VPC endpoints. SDK
// versions released on or after April 1, 2026 route to
// stream-logs.Region.amazonaws.com , which supports VPC endpoints. To set up a VPC
// endpoint for this API, see [Creating a VPC endpoint for CloudWatch Logs].
//
// You can end a session before it times out by closing the session stream or by
// closing the client that is receiving the stream. The session also ends if the
// established connection between the client and the server breaks.
//
// For examples of using an SDK to start a Live Tail session, see [Start a Live Tail session using an Amazon Web Services SDK].
//
// [LiveTailSessionStart]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_LiveTailSessionStart.html
// [LiveTailSessionUpdate]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_LiveTailSessionUpdate.html
// [Use Live Tail to view logs in near real time]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatchLogs_LiveTail.html
// [Creating a VPC endpoint for CloudWatch Logs]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/cloudwatch-logs-and-interface-VPC.html#create-VPC-endpoint-for-CloudWatchLogs
// [Start a Live Tail session using an Amazon Web Services SDK]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/example_cloudwatch-logs_StartLiveTail_section.html
//
// [SessionTimeoutException]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_StartLiveTailResponseStream.html#CWL-Type-StartLiveTailResponseStream-SessionTimeoutException
// [SessionStreamingException]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_StartLiveTailResponseStream.html#CWL-Type-StartLiveTailResponseStream-SessionStreamingException
