package configservice

// PutRemediationExceptions is generated as a reference stub.
// Executable command wiring lives under cmd/configservice.go.
//
// A remediation exception is when a specified resource is no longer considered
// for auto-remediation. This API adds a new exception or updates an existing
// exception for a specified resource with a specified Config rule.
//
// # Exceptions block auto remediation
//
// Config generates a remediation exception when a problem occurs running a
// remediation action for a specified resource. Remediation exceptions blocks
// auto-remediation until the exception is cleared.
//
// # Manual remediation is recommended when placing an exception
//
// When placing an exception on an Amazon Web Services resource, it is recommended
// that remediation is set as manual remediation until the given Config rule for
// the specified resource evaluates the resource as NON_COMPLIANT . Once the
// resource has been evaluated as NON_COMPLIANT , you can add remediation
// exceptions and change the remediation type back from Manual to Auto if you want
// to use auto-remediation. Otherwise, using auto-remediation before a
// NON_COMPLIANT evaluation result can delete resources before the exception is
// applied.
//
// # Exceptions can only be performed on non-compliant resources
//
// Placing an exception can only be performed on resources that are NON_COMPLIANT .
// If you use this API for COMPLIANT resources or resources that are NOT_APPLICABLE
// , a remediation exception will not be generated. For more information on the
// conditions that initiate the possible Config evaluation results, see [Concepts | Config Rules]in the
// Config Developer Guide.
//
// # Exceptions cannot be placed on service-linked remediation actions
//
// You cannot place an exception on service-linked remediation actions, such as
// remediation actions put by an organizational conformance pack.
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
// [Concepts | Config Rules]: https://docs.aws.amazon.com/config/latest/developerguide/config-concepts.html#aws-config-rules
// [PutRemediationConfigurations]: https://docs.aws.amazon.com/config/latest/APIReference/emAPI_PutRemediationConfigurations.html
