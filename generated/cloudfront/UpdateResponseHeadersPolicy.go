package cloudfront

// UpdateResponseHeadersPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/cloudfront.go.
//
// Updates a response headers policy.
//
// When you update a response headers policy, the entire policy is replaced. You
// cannot update some policy fields independent of others. To update a response
// headers policy configuration:
//
// - Use GetResponseHeadersPolicyConfig to get the current policy's configuration.
//
// - Modify the fields in the response headers policy configuration that you
// want to update.
//
// - Call UpdateResponseHeadersPolicy , providing the entire response headers
// policy configuration, including the fields that you modified and those that you
// didn't.
