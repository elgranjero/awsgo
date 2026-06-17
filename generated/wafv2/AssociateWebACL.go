package wafv2

// AssociateWebACL is generated as a reference stub.
// Executable command wiring lives under cmd/wafv2.go.
//
// Associates a web ACL with a resource, to protect the resource.
//
// Use this for all resource types except for Amazon CloudFront distributions. For
// Amazon CloudFront, call UpdateDistribution for the distribution and provide the
// Amazon Resource Name (ARN) of the web ACL in the web ACL ID. For information,
// see [UpdateDistribution]in the Amazon CloudFront Developer Guide.
//
// # Required permissions for customer-managed IAM policies
//
// This call requires permissions that are specific to the protected resource
// type. For details, see [Permissions for AssociateWebACL]in the WAF Developer Guide.
//
// # Temporary inconsistencies during updates
//
// When you create or change a web ACL or other WAF resources, the changes take a
// small amount of time to propagate to all areas where the resources are stored.
// The propagation time can be from a few seconds to a number of minutes.
//
// The following are examples of the temporary inconsistencies that you might
// notice during change propagation:
//
// - After you create a web ACL, if you try to associate it with a resource, you
// might get an exception indicating that the web ACL is unavailable.
//
// - After you add a rule group to a web ACL, the new rule group rules might be
// in effect in one area where the web ACL is used and not in another.
//
// - After you change a rule action setting, you might see the old action in
// some places and the new action in others.
//
// - After you add an IP address to an IP set that is in use in a blocking rule,
// the new address might be blocked in one area while still allowed in another.
//
// [UpdateDistribution]: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_UpdateDistribution.html
// [Permissions for AssociateWebACL]: https://docs.aws.amazon.com/waf/latest/developerguide/security_iam_service-with-iam.html#security_iam_action-AssociateWebACL
