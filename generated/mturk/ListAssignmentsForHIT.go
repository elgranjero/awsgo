package mturk

// ListAssignmentsForHIT is generated as a reference stub.
// Executable command wiring lives under cmd/mturk.go.
//
// The ListAssignmentsForHIT operation retrieves completed assignments for a HIT.
//
// You can use this operation to retrieve the results for a HIT.
//
// You can get assignments for a HIT at any time, even if the HIT is not yet
// Reviewable. If a HIT requested multiple assignments, and has received some
// results but has not yet become Reviewable, you can still retrieve the partial
// results with this operation.
//
// Use the AssignmentStatus parameter to control which set of assignments for a
// HIT are returned. The ListAssignmentsForHIT operation can return submitted
// assignments awaiting approval, or it can return assignments that have already
// been approved or rejected. You can set AssignmentStatus=Approved,Rejected to get
// assignments that have already been approved and rejected together in one result
// set.
//
// Only the Requester who created the HIT can retrieve the assignments for that
// HIT.
//
// Results are sorted and divided into numbered pages and the operation returns a
// single page of results. You can use the parameters of the operation to control
// sorting and pagination.
