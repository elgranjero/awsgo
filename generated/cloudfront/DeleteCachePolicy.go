package cloudfront

// DeleteCachePolicy is generated as a reference stub.
// Executable command wiring lives under cmd/cloudfront.go.
//
// Deletes a cache policy.
//
// You cannot delete a cache policy if it's attached to a cache behavior. First
// update your distributions to remove the cache policy from all cache behaviors,
// then delete the cache policy.
//
// To delete a cache policy, you must provide the policy's identifier and version.
// To get these values, you can use ListCachePolicies or GetCachePolicy .
