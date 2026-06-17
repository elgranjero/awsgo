package support

// DescribeServices is generated as a reference stub.
// Executable command wiring lives under cmd/support.go.
//
// Returns the current list of Amazon Web Services services and a list of service
// categories for each service. You then use service names and categories in your CreateCase
// requests. Each Amazon Web Services service has its own set of categories.
//
// The service codes and category codes correspond to the values that appear in
// the Service and Category lists on the Amazon Web Services Support Center [Create Case]page.
// The values in those fields don't necessarily match the service codes and
// categories returned by the DescribeServices operation. Always use the service
// codes and categories that the DescribeServices operation returns, so that you
// have the most recent set of service and category codes.
//
// - You must have a Business, Enterprise On-Ramp, or Enterprise Support plan to
// use the Amazon Web Services Support API.
//
// - If you call the Amazon Web Services Support API from an account that
// doesn't have a Business, Enterprise On-Ramp, or Enterprise Support plan, the
// SubscriptionRequiredException error message appears. For information about
// changing your support plan, see [Amazon Web Services Support].
//
// [Amazon Web Services Support]: http://aws.amazon.com/premiumsupport/
// [Create Case]: https://console.aws.amazon.com/support/home#/case/create
