package storagegateway

// UpdateBandwidthRateLimitSchedule is generated as a reference stub.
// Executable command wiring lives under cmd/storagegateway.go.
//
// Updates the bandwidth rate limit schedule for a specified gateway. By default,
//
// gateways do not have bandwidth rate limit schedules, which means no bandwidth
// rate limiting is in effect. Use this to initiate or update a gateway's bandwidth
// rate limit schedule. This operation is supported for volume, tape, and S3 file
// gateways. S3 file gateways support bandwidth rate limits for upload only. FSx
// file gateways do not support bandwidth rate limits.
