package waf

// GetSampledRequests is generated as a reference stub.
// Executable command wiring lives under cmd/waf.go.
//
// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Gets detailed information about a specified number of requests--a sample--that
// AWS WAF randomly selects from among the first 5,000 requests that your AWS
// resource received during a time range that you choose. You can specify a sample
// size of up to 500 requests, and you can specify any time range in the previous
// three hours.
//
// GetSampledRequests returns a time range, which is usually the time range that
// you specified. However, if your resource (such as a CloudFront distribution)
// received 5,000 requests before the specified time range elapsed,
// GetSampledRequests returns an updated time range. This new time range indicates
// the actual period during which AWS WAF selected the requests in the sample.
//
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
