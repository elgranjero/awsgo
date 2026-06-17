package health

// EnableHealthServiceAccessForOrganization is generated as a reference stub.
// Executable command wiring lives under cmd/health.go.
//
// Enables Health to work with Organizations. You can use the organizational view
// feature to aggregate events from all Amazon Web Services accounts in your
// organization in a centralized location.
//
// This operation also creates a service-linked role for the management account in
// the organization.
//
// To call this operation, you must meet the following requirements:
//
// - You must have a Business, Enterprise On-Ramp, or Enterprise Support plan
// from [Amazon Web Services Support]to use the Health API. If you call the Health API from an Amazon Web
// Services account that doesn't have a Business, Enterprise On-Ramp, or Enterprise
// Support plan, you receive a SubscriptionRequiredException error.
//
// - You must have permission to call this operation from the organization's
// management account. For example IAM policies, see [Health identity-based policy examples].
//
// If you don't have the required support plan, you can instead use the Health
// console to enable the organizational view feature. For more information, see [Aggregating Health events]in
// the Health User Guide.
//
// [Amazon Web Services Support]: http://aws.amazon.com/premiumsupport/
// [Aggregating Health events]: https://docs.aws.amazon.com/health/latest/ug/aggregate-events.html
// [Health identity-based policy examples]: https://docs.aws.amazon.com/health/latest/ug/security_iam_id-based-policy-examples.html
