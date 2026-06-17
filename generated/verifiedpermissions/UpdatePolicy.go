package verifiedpermissions

// UpdatePolicy is generated as a reference stub.
// Executable command wiring lives under cmd/verifiedpermissions.go.
//
// Modifies a Cedar static policy in the specified policy store. You can change
// only certain elements of the [UpdatePolicyDefinition]parameter. You can directly update only static
// policies. To change a template-linked policy, you must update the template
// instead, using [UpdatePolicyTemplate].
//
// - If policy validation is enabled in the policy store, then updating a static
// policy causes Verified Permissions to validate the policy against the schema in
// the policy store. If the updated static policy doesn't pass validation, the
// operation fails and the update isn't stored.
//
// - When you edit a static policy, you can change only certain elements of a
// static policy:
//
// - The action referenced by the policy.
//
// - A condition clause, such as when and unless.
//
// You can't change these elements of a static policy:
//
// - Changing a policy from a static policy to a template-linked policy.
//
// - Changing the effect of a static policy from permit or forbid.
//
// - The principal referenced by a static policy.
//
// - The resource referenced by a static policy.
//
// - To update a template-linked policy, you must update the template instead.
//
// Verified Permissions is [eventually consistent] . It can take a few seconds for a new or changed
// element to propagate through the service and be visible in the results of other
// Verified Permissions operations.
//
// [eventually consistent]: https://wikipedia.org/wiki/Eventual_consistency
// [UpdatePolicyTemplate]: https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_UpdatePolicyTemplate.html
// [UpdatePolicyDefinition]: https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_UpdatePolicyInput.html#amazonverifiedpermissions-UpdatePolicy-request-UpdatePolicyDefinition
