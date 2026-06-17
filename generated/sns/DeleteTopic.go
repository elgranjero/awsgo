package sns

// DeleteTopic is generated as a reference stub.
// Executable command wiring lives under cmd/sns.go.
//
// Deletes a topic and all its subscriptions. Deleting a topic might prevent some
// messages previously sent to the topic from being delivered to subscribers. This
// action is idempotent, so deleting a topic that does not exist does not result in
// an error.
