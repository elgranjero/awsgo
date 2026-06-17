package cloudfront

// UpdateDistributionWithStagingConfig is generated as a reference stub.
// Executable command wiring lives under cmd/cloudfront.go.
//
// Copies the staging distribution's configuration to its corresponding primary
// distribution. The primary distribution retains its Aliases (also known as
// alternate domain names or CNAMEs) and ContinuousDeploymentPolicyId value, but
// otherwise its configuration is overwritten to match the staging distribution.
//
// You can use this operation in a continuous deployment workflow after you have
// tested configuration changes on the staging distribution. After using a
// continuous deployment policy to move a portion of your domain name's traffic to
// the staging distribution and verifying that it works as intended, you can use
// this operation to copy the staging distribution's configuration to the primary
// distribution. This action will disable the continuous deployment policy and move
// your domain's traffic back to the primary distribution.
//
// This API operation requires the following IAM permissions:
//
// [GetDistribution]
//
// [UpdateDistribution]
//
// [GetDistribution]: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_GetDistribution.html
// [UpdateDistribution]: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_UpdateDistribution.html
