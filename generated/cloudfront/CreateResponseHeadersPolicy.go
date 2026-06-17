package cloudfront

// CreateResponseHeadersPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/cloudfront.go.
//
// Creates a response headers policy.
//
// A response headers policy contains information about a set of HTTP headers. To
// create a response headers policy, you provide some metadata about the policy and
// a set of configurations that specify the headers.
//
// After you create a response headers policy, you can use its ID to attach it to
// one or more cache behaviors in a CloudFront distribution. When it's attached to
// a cache behavior, the response headers policy affects the HTTP headers that
// CloudFront includes in HTTP responses to requests that match the cache behavior.
// CloudFront adds or removes response headers according to the configuration of
// the response headers policy.
//
// For more information, see [Adding or removing HTTP headers in CloudFront responses] in the Amazon CloudFront Developer Guide.
//
// [Adding or removing HTTP headers in CloudFront responses]: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/modifying-response-headers.html
