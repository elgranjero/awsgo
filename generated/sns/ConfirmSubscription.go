package sns

// ConfirmSubscription is generated as a reference stub.
// Executable command wiring lives under cmd/sns.go.
//
// Verifies an endpoint owner's intent to receive messages by validating the token
// sent to the endpoint by an earlier Subscribe action. If the token is valid, the
// action creates a new subscription and returns its Amazon Resource Name (ARN).
// This call requires an AWS signature only when the AuthenticateOnUnsubscribe
// flag is set to "true".
