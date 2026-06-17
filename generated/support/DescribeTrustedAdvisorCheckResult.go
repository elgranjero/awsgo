package support

// DescribeTrustedAdvisorCheckResult is generated as a reference stub.
// Executable command wiring lives under cmd/support.go.
//
// Returns the results of the Trusted Advisor check that has the specified check
// ID. You can get the check IDs by calling the DescribeTrustedAdvisorChecksoperation.
//
// The response contains a TrustedAdvisorCheckResult object, which contains these three objects:
//
// # TrustedAdvisorCategorySpecificSummary
//
// # TrustedAdvisorResourceDetail
//
// # TrustedAdvisorResourcesSummary
//
// In addition, the response contains these fields:
//
// - status - The alert status of the check can be ok (green), warning (yellow),
// error (red), or not_available .
//
// - timestamp - The time of the last refresh of the check.
//
// - checkId - The unique identifier for the check.
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
