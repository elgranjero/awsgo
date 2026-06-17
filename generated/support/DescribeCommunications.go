package support

// DescribeCommunications is generated as a reference stub.
// Executable command wiring lives under cmd/support.go.
//
// Returns communications and attachments for one or more support cases. Use the
// afterTime and beforeTime parameters to filter by date. You can use the caseId
// parameter to restrict the results to a specific case.
//
// Case data is available for 12 months after creation. If a case was created more
// than 12 months ago, a request for data might cause an error.
//
// You can use the maxResults and nextToken parameters to control the pagination
// of the results. Set maxResults to the number of cases that you want to display
// on each page, and use nextToken to specify the resumption of pagination.
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
