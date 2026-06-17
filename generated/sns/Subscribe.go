package sns

// Subscribe is generated as a reference stub.
// Executable command wiring lives under cmd/sns.go.
//
// Subscribes an endpoint to an Amazon SNS topic. If the endpoint type is HTTP/S
// or email, or if the endpoint and the topic are not in the same Amazon Web
// Services account, the endpoint owner must run the ConfirmSubscription action to
// confirm the subscription.
//
// You call the ConfirmSubscription action with the token from the subscription
// response. Confirmation tokens are valid for two days.
//
// This action is throttled at 100 transactions per second (TPS).
