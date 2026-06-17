package cloudfront

// CopyDistribution is generated as a reference stub.
// Executable command wiring lives under cmd/cloudfront.go.
//
// Creates a staging distribution using the configuration of the provided primary
// distribution. A staging distribution is a copy of an existing distribution
// (called the primary distribution) that you can use in a continuous deployment
// workflow.
//
// After you create a staging distribution, you can use UpdateDistribution to
// modify the staging distribution's configuration. Then you can use
// CreateContinuousDeploymentPolicy to incrementally move traffic to the staging
// distribution.
//
// This API operation requires the following IAM permissions:
//
// [GetDistribution]
//
// [CreateDistribution]
//
// [CopyDistribution]
//
// [CopyDistribution]: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_CopyDistribution.html
// [GetDistribution]: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_GetDistribution.html
// [CreateDistribution]: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_CreateDistribution.html
