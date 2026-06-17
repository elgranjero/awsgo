package verifiedpermissions

// CreatePolicy is generated as a reference stub.
// Executable command wiring lives under cmd/verifiedpermissions.go.
//
// Creates a Cedar policy and saves it in the specified policy store. You can
// create either a static policy or a policy linked to a policy template.
//
// - To create a static policy, provide the Cedar policy text in the StaticPolicy
// section of the PolicyDefinition .
//
// - To create a policy that is dynamically linked to a policy template, specify
// the policy template ID and the principal and resource to associate with this
// policy in the templateLinked section of the PolicyDefinition . If the policy
// template is ever updated, any policies linked to the policy template
// automatically use the updated template.
//
// Creating a policy causes it to be validated against the schema in the policy
// store. If the policy doesn't pass validation, the operation fails and the policy
// isn't stored.
//
// Verified Permissions is [eventually consistent] . It can take a few seconds for a new or changed
// element to propagate through the service and be visible in the results of other
// Verified Permissions operations.
//
// [eventually consistent]: https://wikipedia.org/wiki/Eventual_consistency
