package partnercentralselling

// StartEngagementByAcceptingInvitationTask is generated as a reference stub.
// Executable command wiring lives under cmd/partnercentralselling.go.
//
// This action starts the engagement by accepting an EngagementInvitation . The
// task is asynchronous and involves the following steps: accepting the invitation,
// creating an opportunity in the partner’s account from the AWS opportunity, and
// copying details for tracking. When completed, an Opportunity Created event is
// generated, indicating that the opportunity has been successfully created in the
// partner's account.
