package sns

// Unsubscribe is generated as a reference stub.
// Executable command wiring lives under cmd/sns.go.
//
// Deletes a subscription. If the subscription requires authentication for
// deletion, only the owner of the subscription or the topic's owner can
// unsubscribe, and an Amazon Web Services signature is required. If the
// Unsubscribe call does not require authentication and the requester is not the
// subscription owner, a final cancellation message is delivered to the endpoint,
// so that the endpoint owner can easily resubscribe to the topic if the
// Unsubscribe request was unintended.
//
// This action is throttled at 100 transactions per second (TPS).
