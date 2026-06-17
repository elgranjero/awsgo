package eventbridge

// PutTargets is generated as a reference stub.
// Executable command wiring lives under cmd/eventbridge.go.
//
// Adds the specified targets to the specified rule, or updates the targets if
// they are already associated with the rule.
//
// Targets are the resources that are invoked when a rule is triggered.
//
// The maximum number of entries per request is 10.
//
// Each rule can have up to five (5) targets associated with it at one time.
//
// For a list of services you can configure as targets for events, see [EventBridge targets] in the
// Amazon EventBridge User Guide .
//
// Creating rules with built-in targets is supported only in the Amazon Web
// Services Management Console. The built-in targets are:
//
// - Amazon EBS CreateSnapshot API call
//
// - Amazon EC2 RebootInstances API call
//
// - Amazon EC2 StopInstances API call
//
// - Amazon EC2 TerminateInstances API call
//
// For some target types, PutTargets provides target-specific parameters. If the
// target is a Kinesis data stream, you can optionally specify which shard the
// event goes to by using the KinesisParameters argument. To invoke a command on
// multiple EC2 instances with one rule, you can use the RunCommandParameters
// field.
//
// To be able to make API calls against the resources that you own, Amazon
// EventBridge needs the appropriate permissions:
//
// - For Lambda and Amazon SNS resources, EventBridge relies on resource-based
// policies.
//
// - For EC2 instances, Kinesis Data Streams, Step Functions state machines and
// API Gateway APIs, EventBridge relies on IAM roles that you specify in the
// RoleARN argument in PutTargets .
//
// For more information, see [Authentication and Access Control] in the Amazon EventBridge User Guide .
//
// If another Amazon Web Services account is in the same region and has granted
// you permission (using PutPermission ), you can send events to that account. Set
// that account's event bus as a target of the rules in your account. To send the
// matched events to the other account, specify that account's event bus as the Arn
// value when you run PutTargets . If your account sends events to another account,
// your account is charged for each sent event. Each event sent to another account
// is charged as a custom event. The account receiving the event is not charged.
// For more information, see [Amazon EventBridge Pricing].
//
// Input , InputPath , and InputTransformer are not available with PutTarget if
// the target is an event bus of a different Amazon Web Services account.
//
// If you are setting the event bus of another account as the target, and that
// account granted permission to your account through an organization instead of
// directly by the account ID, then you must specify a RoleArn with proper
// permissions in the Target structure. For more information, see [Sending and Receiving Events Between Amazon Web Services Accounts] in the Amazon
// EventBridge User Guide.
//
// If you have an IAM role on a cross-account event bus target, a PutTargets call
// without a role on the same target (same Id and Arn ) will not remove the role.
//
// For more information about enabling cross-account events, see [PutPermission].
//
// Input, InputPath, and InputTransformer are mutually exclusive and optional
// parameters of a target. When a rule is triggered due to a matched event:
//
// - If none of the following arguments are specified for a target, then the
// entire event is passed to the target in JSON format (unless the target is Amazon
// EC2 Run Command or Amazon ECS task, in which case nothing from the event is
// passed to the target).
//
// - If Input is specified in the form of valid JSON, then the matched event is
// overridden with this constant.
//
// - If InputPath is specified in the form of JSONPath (for example, $.detail ),
// then only the part of the event specified in the path is passed to the target
// (for example, only the detail part of the event is passed).
//
// - If InputTransformer is specified, then one or more specified JSONPaths are
// extracted from the event and used as values in a template that you specify as
// the input to the target.
//
// When you specify InputPath or InputTransformer , you must use JSON dot notation,
// not bracket notation.
//
// When you add targets to a rule and the associated rule triggers soon after, new
// or updated targets might not be immediately invoked. Allow a short period of
// time for changes to take effect.
//
// This action can partially fail if too many requests are made at the same time.
// If that happens, FailedEntryCount is non-zero in the response and each entry in
// FailedEntries provides the ID of the failed target and the error code.
//
// [Authentication and Access Control]: https://docs.aws.amazon.com/eventbridge/latest/userguide/auth-and-access-control-eventbridge.html
// [Amazon EventBridge Pricing]: http://aws.amazon.com/eventbridge/pricing/
// [Sending and Receiving Events Between Amazon Web Services Accounts]: https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-cross-account-event-delivery.html
// [PutPermission]: https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_PutPermission.html
// [EventBridge targets]: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-targets.html
