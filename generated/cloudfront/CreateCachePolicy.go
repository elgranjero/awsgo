package cloudfront

// CreateCachePolicy is generated as a reference stub.
// Executable command wiring lives under cmd/cloudfront.go.
//
// Creates a cache policy.
//
// After you create a cache policy, you can attach it to one or more cache
// behaviors. When it's attached to a cache behavior, the cache policy determines
// the following:
//
// - The values that CloudFront includes in the cache key. These values can
// include HTTP headers, cookies, and URL query strings. CloudFront uses the cache
// key to find an object in its cache that it can return to the viewer.
//
// - The default, minimum, and maximum time to live (TTL) values that you want
// objects to stay in the CloudFront cache.
//
// If your minimum TTL is greater than 0, CloudFront will cache content for at
//
// least the duration specified in the cache policy's minimum TTL, even if the
// Cache-Control: no-cache , no-store , or private directives are present in the
// origin headers.
//
// The headers, cookies, and query strings that are included in the cache key are
// also included in requests that CloudFront sends to the origin. CloudFront sends
// a request when it can't find an object in its cache that matches the request's
// cache key. If you want to send values to the origin but not include them in the
// cache key, use OriginRequestPolicy .
//
// For more information about cache policies, see [Controlling the cache key] in the Amazon CloudFront
// Developer Guide.
//
// [Controlling the cache key]: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html
