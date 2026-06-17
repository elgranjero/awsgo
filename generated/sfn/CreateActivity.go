package sfn

// CreateActivity is generated as a reference stub.
// Executable command wiring lives under cmd/sfn.go.
//
// Creates an activity. An activity is a task that you write in any programming
// language and host on any machine that has access to Step Functions. Activities
// must poll Step Functions using the GetActivityTask API action and respond using
// SendTask* API actions. This function lets Step Functions know the existence of
// your activity and returns an identifier for use in a state machine and when
// polling from the activity.
//
// This operation is eventually consistent. The results are best effort and may
// not reflect very recent updates and changes.
//
// CreateActivity is an idempotent API. Subsequent requests won’t create a
// duplicate resource if it was already created. CreateActivity 's idempotency
// check is based on the activity name . If a following request has different tags
// values, Step Functions will ignore these differences and treat it as an
// idempotent request of the previous. In this case, tags will not be updated,
// even if they are different.
