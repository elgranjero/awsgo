package cloudfront

// CreateContinuousDeploymentPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/cloudfront.go.
//
// Creates a continuous deployment policy that distributes traffic for a custom
// domain name to two different CloudFront distributions.
//
// To use a continuous deployment policy, first use CopyDistribution to create a
// staging distribution, then use UpdateDistribution to modify the staging
// distribution's configuration.
//
// After you create and update a staging distribution, you can use a continuous
// deployment policy to incrementally move traffic to the staging distribution.
// This workflow enables you to test changes to a distribution's configuration
// before moving all of your domain's production traffic to the new configuration.
