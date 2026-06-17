package partnercentralselling

// CreateOpportunity is generated as a reference stub.
// Executable command wiring lives under cmd/partnercentralselling.go.
//
// Creates an Opportunity record in Partner Central. Use this operation to create
// a potential business opportunity for submission to Amazon Web Services. Creating
// an opportunity sets Lifecycle.ReviewStatus to Pending Submission .
//
// To submit an opportunity, follow these steps:
//
// - To create the opportunity, use CreateOpportunity .
//
// - To associate a solution with the opportunity, use AssociateOpportunity .
//
// - To start the engagement with AWS, use StartEngagementFromOpportunity .
//
// After submission, you can't edit the opportunity until the review is complete.
// But opportunities in the Pending Submission state must have complete details.
// You can update the opportunity while it's in the Pending Submission state.
//
// There's a set of mandatory fields to create opportunities, but consider
// providing optional fields to enrich the opportunity record.
