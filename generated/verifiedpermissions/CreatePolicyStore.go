package verifiedpermissions

// CreatePolicyStore is generated as a reference stub.
// Executable command wiring lives under cmd/verifiedpermissions.go.
//
// Creates a policy store. A policy store is a container for policy resources.
//
// Although [Cedar supports multiple namespaces], Verified Permissions currently supports only one namespace per
// policy store.
//
// Verified Permissions is [eventually consistent] . It can take a few seconds for a new or changed
// element to propagate through the service and be visible in the results of other
// Verified Permissions operations.
//
// [eventually consistent]: https://wikipedia.org/wiki/Eventual_consistency
// [Cedar supports multiple namespaces]: https://docs.cedarpolicy.com/schema/schema.html#namespace
