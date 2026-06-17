package mturk

// DeleteHIT is generated as a reference stub.
// Executable command wiring lives under cmd/mturk.go.
//
// The DeleteHIT operation is used to delete HIT that is no longer needed. Only
//
// the Requester who created the HIT can delete it.
//
// You can only dispose of HITs that are in the Reviewable state, with all of
// their submitted assignments already either approved or rejected. If you call the
// DeleteHIT operation on a HIT that is not in the Reviewable state (for example,
// that has not expired, or still has active assignments), or on a HIT that is
// Reviewable but without all of its submitted assignments already approved or
// rejected, the service will return an error.
//
// - HITs are automatically disposed of after 120 days.
//
// - After you dispose of a HIT, you can no longer approve the HIT's rejected
// assignments.
//
// - Disposed HITs are not returned in results for the ListHITs operation.
//
// - Disposing HITs can improve the performance of operations such as
// ListReviewableHITs and ListHITs.
