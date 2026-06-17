package partnercentralselling

// StartEngagementFromOpportunityTask is generated as a reference stub.
// Executable command wiring lives under cmd/partnercentralselling.go.
//
// Similar to StartEngagementByAcceptingInvitationTask , this action is
// asynchronous and performs multiple steps before completion. This action
// orchestrates a comprehensive workflow that combines multiple API operations into
// a single task to create and initiate an engagement from an existing opportunity.
// It automatically executes a sequence of operations including GetOpportunity ,
// CreateEngagement (if it doesn't exist), CreateResourceSnapshot ,
// CreateResourceSnapshotJob , CreateEngagementInvitation (if not already
// invited/accepted), and SubmitOpportunity .
