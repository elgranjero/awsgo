package storagegateway

// UpdateBandwidthRateLimit is generated as a reference stub.
// Executable command wiring lives under cmd/storagegateway.go.
//
// Updates the bandwidth rate limits of a gateway. You can update both the upload
// and download bandwidth rate limit or specify only one of the two. If you don't
// set a bandwidth rate limit, the existing rate limit remains. This operation is
// supported only for the stored volume, cached volume, and tape gateway types. To
// update bandwidth rate limits for S3 file gateways, use UpdateBandwidthRateLimitSchedule.
//
// By default, a gateway's bandwidth rate limits are not set. If you don't set any
// limit, the gateway does not have any limitations on its bandwidth usage and
// could potentially use the maximum available bandwidth.
//
// To specify which gateway to update, use the Amazon Resource Name (ARN) of the
// gateway in your request.
