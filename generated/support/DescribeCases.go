package support

// DescribeCases is generated as a reference stub.
// Executable command wiring lives under cmd/support.go.
//
// Returns a list of cases that you specify by passing one or more case IDs. You
// can use the afterTime and beforeTime parameters to filter the cases by date.
// You can set values for the includeResolvedCases and includeCommunications
// parameters to specify how much information to return.
//
// The response returns the following in JSON format:
//
// - One or more [CaseDetails]data types.
//
// - One or more nextToken values, which specify where to paginate the returned
// records represented by the CaseDetails objects.
//
// Case data is available for 12 months after creation. If a case was created more
// than 12 months ago, a request might return an error.
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
// [CaseDetails]: https://docs.aws.amazon.com/awssupport/latest/APIReference/API_CaseDetails.html
