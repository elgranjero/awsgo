package mturk

// ApproveAssignment is generated as a reference stub.
// Executable command wiring lives under cmd/mturk.go.
//
// The ApproveAssignment operation approves the results of a completed
//
// assignment.
//
// Approving an assignment initiates two payments from the Requester's Amazon.com
// account
//
// - The Worker who submitted the results is paid the reward specified in the
// HIT.
//
// - Amazon Mechanical Turk fees are debited.
//
// If the Requester's account does not have adequate funds for these payments, the
// call to ApproveAssignment returns an exception, and the approval is not
// processed. You can include an optional feedback message with the approval, which
// the Worker can see in the Status section of the web site.
//
// You can also call this operation for assignments that were previous rejected
// and approve them by explicitly overriding the previous rejection. This only
// works on rejected assignments that were submitted within the previous 30 days
// and only if the assignment's related HIT has not been deleted.
