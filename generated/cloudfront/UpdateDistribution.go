package cloudfront

// UpdateDistribution is generated as a reference stub.
// Executable command wiring lives under cmd/cloudfront.go.
//
// Updates the configuration for a CloudFront distribution.
//
// The update process includes getting the current distribution configuration,
// updating it to make your changes, and then submitting an UpdateDistribution
// request to make the updates.
//
// To update a web distribution using the CloudFront API
//
// - Use GetDistributionConfig to get the current configuration, including the
// version identifier ( ETag ).
//
// - Update the distribution configuration that was returned in the response.
// Note the following important requirements and restrictions:
//
// - You must copy the ETag field value from the response. (You'll use it for the
// IfMatch parameter in your request.) Then, remove the ETag field from the
// distribution configuration.
//
// - You can't change the value of CallerReference .
//
// - Submit an UpdateDistribution request, providing the updated distribution
// configuration. The new configuration replaces the existing configuration. The
// values that you specify in an UpdateDistribution request are not merged into
// your existing configuration. Make sure to include all fields: the ones that you
// modified and also the ones that you didn't.
