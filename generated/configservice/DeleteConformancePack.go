package configservice

// DeleteConformancePack is generated as a reference stub.
// Executable command wiring lives under cmd/configservice.go.
//
// Deletes the specified conformance pack and all the Config rules, remediation
// actions, and all evaluation results within that conformance pack.
//
// Config sets the conformance pack to DELETE_IN_PROGRESS until the deletion is
// complete. You cannot update a conformance pack while it is in this state.
//
// Recommendation: Consider excluding the AWS::Config::ResourceCompliance resource
// type from recording before deleting rules
//
// Deleting rules creates configuration items (CIs) for
// AWS::Config::ResourceCompliance that can affect your costs for the configuration
// recorder. If you are deleting rules which evaluate a large number of resource
// types, this can lead to a spike in the number of CIs recorded.
//
// To avoid the associated costs, you can opt to disable recording for the
// AWS::Config::ResourceCompliance resource type before deleting rules, and
// re-enable recording after the rules have been deleted.
//
// However, since deleting rules is an asynchronous process, it might take an hour
// or more to complete. During the time when recording is disabled for
// AWS::Config::ResourceCompliance , rule evaluations will not be recorded in the
// associated resource’s history.
