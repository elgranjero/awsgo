package storagegateway

// DescribeBandwidthRateLimitSchedule is generated as a reference stub.
// Executable command wiring lives under cmd/storagegateway.go.
//
// Returns information about the bandwidth rate limit schedule of a gateway. By
//
// default, gateways do not have bandwidth rate limit schedules, which means no
// bandwidth rate limiting is in effect. This operation is supported only for
// volume, tape and S3 file gateways. FSx file gateways do not support bandwidth
// rate limits.
//
// This operation returns information about a gateway's bandwidth rate limit
// schedule. A bandwidth rate limit schedule consists of one or more bandwidth rate
// limit intervals. A bandwidth rate limit interval defines a period of time on one
// or more days of the week, during which bandwidth rate limits are specified for
// uploading, downloading, or both.
//
// A bandwidth rate limit interval consists of one or more days of the week, a
// start hour and minute, an ending hour and minute, and bandwidth rate limits for
// uploading and downloading
//
// If no bandwidth rate limit schedule intervals are set for the gateway, this
// operation returns an empty response. To specify which gateway to describe, use
// the Amazon Resource Name (ARN) of the gateway in your request.
