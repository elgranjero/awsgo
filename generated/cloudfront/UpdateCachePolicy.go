package cloudfront

// UpdateCachePolicy is generated as a reference stub.
// Executable command wiring lives under cmd/cloudfront.go.
//
// Updates a cache policy configuration.
//
// When you update a cache policy configuration, all the fields are updated with
// the values provided in the request. You cannot update some fields independent of
// others. To update a cache policy configuration:
//
// - Use GetCachePolicyConfig to get the current configuration.
//
// - Locally modify the fields in the cache policy configuration that you want
// to update.
//
// - Call UpdateCachePolicy by providing the entire cache policy configuration,
// including the fields that you modified and those that you didn't.
//
// If your minimum TTL is greater than 0, CloudFront will cache content for at
// least the duration specified in the cache policy's minimum TTL, even if the
// Cache-Control: no-cache , no-store , or private directives are present in the
// origin headers.
