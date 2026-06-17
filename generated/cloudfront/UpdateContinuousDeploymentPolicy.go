package cloudfront

// UpdateContinuousDeploymentPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/cloudfront.go.
//
// Updates a continuous deployment policy. You can update a continuous deployment
// policy to enable or disable it, to change the percentage of traffic that it
// sends to the staging distribution, or to change the staging distribution that it
// sends traffic to.
//
// When you update a continuous deployment policy configuration, all the fields
// are updated with the values that are provided in the request. You cannot update
// some fields independent of others. To update a continuous deployment policy
// configuration:
//
// - Use GetContinuousDeploymentPolicyConfig to get the current configuration.
//
// - Locally modify the fields in the continuous deployment policy configuration
// that you want to update.
//
// - Use UpdateContinuousDeploymentPolicy , providing the entire continuous
// deployment policy configuration, including the fields that you modified and
// those that you didn't.
