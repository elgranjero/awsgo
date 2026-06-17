package support

// AddCommunicationToCase is generated as a reference stub.
// Executable command wiring lives under cmd/support.go.
//
// Adds additional customer communication to an Amazon Web Services Support case.
// Use the caseId parameter to identify the case to which to add communication.
// You can list a set of email addresses to copy on the communication by using the
// ccEmailAddresses parameter. The communicationBody value contains the text of
// the communication.
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
