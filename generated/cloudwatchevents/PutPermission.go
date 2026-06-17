package cloudwatchevents

// PutPermission is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatchevents.go.
//
// Running PutPermission permits the specified Amazon Web Services account or
// Amazon Web Services organization to put events to the specified event bus.
// Amazon EventBridge (CloudWatch Events) rules in your account are triggered by
// these events arriving to an event bus in your account.
//
// For another account to send events to your account, that external account must
// have an EventBridge rule with your account's event bus as a target.
//
// To enable multiple Amazon Web Services accounts to put events to your event
// bus, run PutPermission once for each of these accounts. Or, if all the accounts
// are members of the same Amazon Web Services organization, you can run
// PutPermission once specifying Principal as "*" and specifying the Amazon Web
// Services organization ID in Condition , to grant permissions to all accounts in
// that organization.
//
// If you grant permissions using an organization, then accounts in that
// organization must specify a RoleArn with proper permissions when they use
// PutTarget to add your account's event bus as a target. For more information, see [Sending and Receiving Events Between Amazon Web Services Accounts]
// in the Amazon EventBridge User Guide.
//
// The permission policy on the event bus cannot exceed 10 KB in size.
//
// [Sending and Receiving Events Between Amazon Web Services Accounts]: https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-cross-account-event-delivery.html
