package shield

// AssociateProactiveEngagementDetails is generated as a reference stub.
// Executable command wiring lives under cmd/shield.go.
//
// Initializes proactive engagement and sets the list of contacts for the Shield
// Response Team (SRT) to use. You must provide at least one phone number in the
// emergency contact list.
//
// After you have initialized proactive engagement using this call, to disable or
// enable proactive engagement, use the calls DisableProactiveEngagement and
// EnableProactiveEngagement .
//
// This call defines the list of email addresses and phone numbers that the SRT
// can use to contact you for escalations to the SRT and to initiate proactive
// customer support.
//
// The contacts that you provide in the request replace any contacts that were
// already defined. If you already have contacts defined and want to use them,
// retrieve the list using DescribeEmergencyContactSettings and then provide it to
// this call.
