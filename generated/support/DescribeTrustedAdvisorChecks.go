package support

// DescribeTrustedAdvisorChecks is generated as a reference stub.
// Executable command wiring lives under cmd/support.go.
//
// Returns information about all available Trusted Advisor checks, including the
// name, ID, category, description, and metadata. You must specify a language code.
//
// The response contains a TrustedAdvisorCheckDescription object for each check. You must set the Amazon Web
// Services Region to us-east-1.
//
// - You must have a Business, Enterprise On-Ramp, or Enterprise Support plan to
// use the Amazon Web Services Support API.
//
// - If you call the Amazon Web Services Support API from an account that
// doesn't have a Business, Enterprise On-Ramp, or Enterprise Support plan, the
// SubscriptionRequiredException error message appears. For information about
// changing your support plan, see [Amazon Web Services Support].
//
// - The names and descriptions for Trusted Advisor checks are subject to
// change. We recommend that you specify the check ID in your code to uniquely
// identify a check.
//
// To call the Trusted Advisor operations in the Amazon Web Services Support API,
// you must use the US East (N. Virginia) endpoint. Currently, the US West (Oregon)
// and Europe (Ireland) endpoints don't support the Trusted Advisor operations. For
// more information, see [About the Amazon Web Services Support API]in the Amazon Web Services Support User Guide.
//
// [Amazon Web Services Support]: http://aws.amazon.com/premiumsupport/
// [About the Amazon Web Services Support API]: https://docs.aws.amazon.com/awssupport/latest/user/about-support-api.html#endpoint
