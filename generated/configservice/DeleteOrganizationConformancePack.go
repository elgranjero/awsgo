package configservice

// DeleteOrganizationConformancePack is generated as a reference stub.
// Executable command wiring lives under cmd/configservice.go.
//
// Deletes the specified organization conformance pack and all of the Config rules
// and remediation actions from all member accounts in that organization.
//
// Only a management account or a delegated administrator account can delete an
// organization conformance pack. When calling this API with a delegated
// administrator, you must ensure Organizations ListDelegatedAdministrator
// permissions are added.
//
// Config sets the state of a conformance pack to DELETE_IN_PROGRESS until the
// deletion is complete. You cannot update a conformance pack while it is in this
// state.
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
