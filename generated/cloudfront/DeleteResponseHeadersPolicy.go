package cloudfront

// DeleteResponseHeadersPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/cloudfront.go.
//
// Deletes a response headers policy.
//
// You cannot delete a response headers policy if it's attached to a cache
// behavior. First update your distributions to remove the response headers policy
// from all cache behaviors, then delete the response headers policy.
//
// To delete a response headers policy, you must provide the policy's identifier
// and version. To get these values, you can use ListResponseHeadersPolicies or
// GetResponseHeadersPolicy .
