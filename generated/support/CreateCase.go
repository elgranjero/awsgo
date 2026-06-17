package support

// CreateCase is generated as a reference stub.
// Executable command wiring lives under cmd/support.go.
//
// Creates a case in the Amazon Web Services Support Center. This operation is
// similar to how you create a case in the Amazon Web Services Support Center [Create Case]page.
//
// The Amazon Web Services Support API doesn't support requesting service limit
// increases. You can submit a service limit increase in the following ways:
//
// - Submit a request from the Amazon Web Services Support Center [Create Case]page.
//
// - Use the Service Quotas [RequestServiceQuotaIncrease]operation.
//
// A successful CreateCase request returns an Amazon Web Services Support case
// number. You can use the DescribeCasesoperation and specify the case number to get existing
// Amazon Web Services Support cases. After you create a case, use the AddCommunicationToCaseoperation
// to add additional communication or attachments to an existing case.
//
// The caseId is separate from the displayId that appears in the [Amazon Web Services Support Center]. Use the DescribeCases
// operation to get the displayId .
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
// [RequestServiceQuotaIncrease]: https://docs.aws.amazon.com/servicequotas/2019-06-24/apireference/API_RequestServiceQuotaIncrease.html
// [Amazon Web Services Support Center]: https://console.aws.amazon.com/support
