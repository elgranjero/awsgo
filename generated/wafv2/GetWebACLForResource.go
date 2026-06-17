package wafv2

// GetWebACLForResource is generated as a reference stub.
// Executable command wiring lives under cmd/wafv2.go.
//
// Retrieves the WebACL for the specified resource.
//
// This call uses GetWebACL , to verify that your account has permission to access
// the retrieved web ACL. If you get an error that indicates that your account
// isn't authorized to perform wafv2:GetWebACL on the resource, that error won't
// be included in your CloudTrail event history.
//
// For Amazon CloudFront, don't use this call. Instead, call the CloudFront action
// GetDistributionConfig . For information, see [GetDistributionConfig] in the Amazon CloudFront API
// Reference.
//
// # Required permissions for customer-managed IAM policies
//
// This call requires permissions that are specific to the protected resource
// type. For details, see [Permissions for GetWebACLForResource]in the WAF Developer Guide.
//
// [GetDistributionConfig]: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_GetDistributionConfig.html
// [Permissions for GetWebACLForResource]: https://docs.aws.amazon.com/waf/latest/developerguide/security_iam_service-with-iam.html#security_iam_action-GetWebACLForResource
