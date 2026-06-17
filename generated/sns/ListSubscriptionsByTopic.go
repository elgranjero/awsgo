package sns

// ListSubscriptionsByTopic is generated as a reference stub.
// Executable command wiring lives under cmd/sns.go.
//
// Returns a list of the subscriptions to a specific topic. Each call returns a
// limited list of subscriptions, up to 100. If there are more subscriptions, a
// NextToken is also returned. Use the NextToken parameter in a new
// ListSubscriptionsByTopic call to get further results.
//
// This action is throttled at 30 transactions per second (TPS).
