package wafv2

// DisassociateWebACL is generated as a reference stub.
// Executable command wiring lives under cmd/wafv2.go.
//
// Disassociates the specified resource from its web ACL association, if it has
// one.
//
// Use this for all resource types except for Amazon CloudFront distributions. For
// Amazon CloudFront, call UpdateDistribution for the distribution and provide an
// empty web ACL ID. For information, see [UpdateDistribution]in the Amazon CloudFront API Reference.
//
// # Required permissions for customer-managed IAM policies
//
// This call requires permissions that are specific to the protected resource
// type. For details, see [Permissions for DisassociateWebACL]in the WAF Developer Guide.
//
// [Permissions for DisassociateWebACL]: https://docs.aws.amazon.com/waf/latest/developerguide/security_iam_service-with-iam.html#security_iam_action-DisassociateWebACL
// [UpdateDistribution]: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_UpdateDistribution.html
