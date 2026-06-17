package configservice

// PutRemediationConfigurations is generated as a reference stub.
// Executable command wiring lives under cmd/configservice.go.
//
// Adds or updates the remediation configuration with a specific Config rule with
// the selected target or action. The API creates the RemediationConfiguration
// object for the Config rule. The Config rule must already exist for you to add a
// remediation configuration. The target (SSM document) must exist and have
// permissions to use the target.
//
// # Be aware of backward incompatible changes
//
// If you make backward incompatible changes to the SSM document, you must call
// this again to ensure the remediations can run.
//
// This API does not support adding remediation configurations for service-linked
// Config Rules such as Organization Config rules, the rules deployed by
// conformance packs, and rules deployed by Amazon Web Services Security Hub.
//
// # Required fields
//
// For manual remediation configuration, you need to provide a value for
// automationAssumeRole or use a value in the assumeRole field to remediate your
// resources. The SSM automation document can use either as long as it maps to a
// valid parameter.
//
// However, for automatic remediation configuration, the only valid assumeRole
// field value is AutomationAssumeRole and you need to provide a value for
// AutomationAssumeRole to remediate your resources.
//
// # Auto remediation can be initiated even for compliant resources
//
// If you enable auto remediation for a specific Config rule using the [PutRemediationConfigurations] API or the
// Config console, it initiates the remediation process for all non-compliant
// resources for that specific rule. The auto remediation process relies on the
// compliance data snapshot which is captured on a periodic basis. Any
// non-compliant resource that is updated between the snapshot schedule will
// continue to be remediated based on the last known compliance data snapshot.
//
// This means that in some cases auto remediation can be initiated even for
// compliant resources, since the bootstrap processor uses a database that can have
// stale evaluation results based on the last known compliance data snapshot.
//
// [PutRemediationConfigurations]: https://docs.aws.amazon.com/config/latest/APIReference/emAPI_PutRemediationConfigurations.html
