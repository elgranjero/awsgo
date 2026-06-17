package verifiedpermissions

// UpdatePolicyTemplate is generated as a reference stub.
// Executable command wiring lives under cmd/verifiedpermissions.go.
//
// Updates the specified policy template. You can update only the description and
// the some elements of the [policyBody].
//
// Changes you make to the policy template content are immediately (within the
// constraints of eventual consistency) reflected in authorization decisions that
// involve all template-linked policies instantiated from this template.
//
// Verified Permissions is [eventually consistent] . It can take a few seconds for a new or changed
// element to propagate through the service and be visible in the results of other
// Verified Permissions operations.
//
// [eventually consistent]: https://wikipedia.org/wiki/Eventual_consistency
// [policyBody]: https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_UpdatePolicyTemplate.html#amazonverifiedpermissions-UpdatePolicyTemplate-request-policyBody
