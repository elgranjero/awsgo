package support

// RefreshTrustedAdvisorCheck is generated as a reference stub.
// Executable command wiring lives under cmd/support.go.
//
// Refreshes the Trusted Advisor check that you specify using the check ID. You
// can get the check IDs by calling the DescribeTrustedAdvisorChecksoperation.
//
// Some checks are refreshed automatically. If you call the
// RefreshTrustedAdvisorCheck operation to refresh them, you might see the
// InvalidParameterValue error.
//
// The response contains a TrustedAdvisorCheckRefreshStatus object.
//
// - You must have a Business, Enterprise On-Ramp, or Enterprise Support plan to
// use the Amazon Web Services Support API.
//
// - If you call the Amazon Web Services Support API from an account that
// doesn't have a Business, Enterprise On-Ramp, or Enterprise Support plan, the
// SubscriptionRequiredException error message appears. For information about
// changing your support plan, see [Amazon Web Services Support].
//
// To call the Trusted Advisor operations in the Amazon Web Services Support API,
// you must use the US East (N. Virginia) endpoint. Currently, the US West (Oregon)
// and Europe (Ireland) endpoints don't support the Trusted Advisor operations. For
// more information, see [About the Amazon Web Services Support API]in the Amazon Web Services Support User Guide.
//
// [Amazon Web Services Support]: http://aws.amazon.com/premiumsupport/
// [About the Amazon Web Services Support API]: https://docs.aws.amazon.com/awssupport/latest/user/about-support-api.html#endpoint
