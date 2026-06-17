package storagegateway

// DescribeBandwidthRateLimit is generated as a reference stub.
// Executable command wiring lives under cmd/storagegateway.go.
//
// Returns the bandwidth rate limits of a gateway. By default, these limits are
// not set, which means no bandwidth rate limiting is in effect. This operation is
// supported only for the stored volume, cached volume, and tape gateway types. To
// describe bandwidth rate limits for S3 file gateways, use DescribeBandwidthRateLimitSchedule.
//
// This operation returns a value for a bandwidth rate limit only if the limit is
// set. If no limits are set for the gateway, then this operation returns only the
// gateway ARN in the response body. To specify which gateway to describe, use the
// Amazon Resource Name (ARN) of the gateway in your request.
