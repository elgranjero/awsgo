package wafv2

// GetSampledRequests is generated as a reference stub.
// Executable command wiring lives under cmd/wafv2.go.
//
// Gets detailed information about a specified number of requests--a sample--that
// WAF randomly selects from among the first 5,000 requests that your Amazon Web
// Services resource received during a time range that you choose. You can specify
// a sample size of up to 500 requests, and you can specify any time range in the
// previous three hours.
//
// GetSampledRequests returns a time range, which is usually the time range that
// you specified. However, if your resource (such as a CloudFront distribution)
// received 5,000 requests before the specified time range elapsed,
// GetSampledRequests returns an updated time range. This new time range indicates
// the actual period during which WAF selected the requests in the sample.
