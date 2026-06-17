package cloudfront

// UpdateOriginRequestPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/cloudfront.go.
//
// Updates an origin request policy configuration.
//
// When you update an origin request policy configuration, all the fields are
// updated with the values provided in the request. You cannot update some fields
// independent of others. To update an origin request policy configuration:
//
// - Use GetOriginRequestPolicyConfig to get the current configuration.
//
// - Locally modify the fields in the origin request policy configuration that
// you want to update.
//
// - Call UpdateOriginRequestPolicy by providing the entire origin request policy
// configuration, including the fields that you modified and those that you didn't.
