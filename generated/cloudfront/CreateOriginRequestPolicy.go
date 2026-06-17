package cloudfront

// CreateOriginRequestPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/cloudfront.go.
//
// Creates an origin request policy.
//
// After you create an origin request policy, you can attach it to one or more
// cache behaviors. When it's attached to a cache behavior, the origin request
// policy determines the values that CloudFront includes in requests that it sends
// to the origin. Each request that CloudFront sends to the origin includes the
// following:
//
// - The request body and the URL path (without the domain name) from the viewer
// request.
//
// - The headers that CloudFront automatically includes in every origin request,
// including Host , User-Agent , and X-Amz-Cf-Id .
//
// - All HTTP headers, cookies, and URL query strings that are specified in the
// cache policy or the origin request policy. These can include items from the
// viewer request and, in the case of headers, additional ones that are added by
// CloudFront.
//
// CloudFront sends a request when it can't find a valid object in its cache that
// matches the request. If you want to send values to the origin and also include
// them in the cache key, use CachePolicy .
//
// For more information about origin request policies, see [Controlling origin requests] in the Amazon
// CloudFront Developer Guide.
//
// [Controlling origin requests]: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-origin-requests.html
