package verifiedpermissions

// CreatePolicyTemplate is generated as a reference stub.
// Executable command wiring lives under cmd/verifiedpermissions.go.
//
// Creates a policy template. A template can use placeholders for the principal
// and resource. A template must be instantiated into a policy by associating it
// with specific principals and resources to use for the placeholders. That
// instantiated policy can then be considered in authorization decisions. The
// instantiated policy works identically to any other policy, except that it is
// dynamically linked to the template. If the template changes, then any policies
// that are linked to that template are immediately updated as well.
//
// Verified Permissions is [eventually consistent] . It can take a few seconds for a new or changed
// element to propagate through the service and be visible in the results of other
// Verified Permissions operations.
//
// [eventually consistent]: https://wikipedia.org/wiki/Eventual_consistency
