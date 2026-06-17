package cloudwatchevents

// DeleteRule is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatchevents.go.
//
// Deletes the specified rule.
//
// Before you can delete the rule, you must remove all targets, using [RemoveTargets].
//
// When you delete a rule, incoming events might continue to match to the deleted
// rule. Allow a short period of time for changes to take effect.
//
// If you call delete rule multiple times for the same rule, all calls will
// succeed. When you call delete rule for a non-existent custom eventbus,
// ResourceNotFoundException is returned.
//
// Managed rules are rules created and managed by another Amazon Web Services
// service on your behalf. These rules are created by those other Amazon Web
// Services services to support functionality in those services. You can delete
// these rules using the Force option, but you should do so only if you are sure
// the other service is not still using that rule.
//
// [RemoveTargets]: https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_RemoveTargets.html
