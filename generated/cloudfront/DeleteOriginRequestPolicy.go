package cloudfront

// DeleteOriginRequestPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/cloudfront.go.
//
// Deletes an origin request policy.
//
// You cannot delete an origin request policy if it's attached to any cache
// behaviors. First update your distributions to remove the origin request policy
// from all cache behaviors, then delete the origin request policy.
//
// To delete an origin request policy, you must provide the policy's identifier
// and version. To get the identifier, you can use ListOriginRequestPolicies or
// GetOriginRequestPolicy .
