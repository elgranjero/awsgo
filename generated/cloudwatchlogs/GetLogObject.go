package cloudwatchlogs

// GetLogObject is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatchlogs.go.
//
// Retrieves a large logging object (LLO) and streams it back. This API is used to
// fetch the content of large portions of log events that have been ingested
// through the PutOpenTelemetryLogs API. When log events contain fields that would
// cause the total event size to exceed 1MB, CloudWatch Logs automatically
// processes up to 10 fields, starting with the largest fields. Each field is
// truncated as needed to keep the total event size as close to 1MB as possible.
// The excess portions are stored as Large Log Objects (LLOs) and these fields are
// processed separately and LLO reference system fields (in the format
// (at)ptr.$[path.to.field] ) are added. The path in the reference field reflects the
// original JSON structure where the large field was located. For example, this
// could be (at)ptr.$['input']['message'] , (at)ptr.$['AAA']['BBB']['CCC']['DDD'] ,
// (at)ptr.$['AAA'] , or any other path matching your log structure.
//
// The GetLogObject API routes requests using SDK host prefix injection. SDK
// versions released before April 1, 2026 route to
// streaming-logs.Region.amazonaws.com , which does not support VPC endpoints. SDK
// versions released on or after April 1, 2026 route to
// stream-logs.Region.amazonaws.com , which supports VPC endpoints. To set up a VPC
// endpoint for this API, see [Creating a VPC endpoint for CloudWatch Logs].
//
// [Creating a VPC endpoint for CloudWatch Logs]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/cloudwatch-logs-and-interface-VPC.html#create-VPC-endpoint-for-CloudWatchLogs
