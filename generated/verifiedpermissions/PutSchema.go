package verifiedpermissions

// PutSchema is generated as a reference stub.
// Executable command wiring lives under cmd/verifiedpermissions.go.
//
// Creates or updates the policy schema in the specified policy store. The schema
// is used to validate any Cedar policies and policy templates submitted to the
// policy store. Any changes to the schema validate only policies and templates
// submitted after the schema change. Existing policies and templates are not
// re-evaluated against the changed schema. If you later update a policy, then it
// is evaluated against the new schema at that time.
//
// Verified Permissions is [eventually consistent] . It can take a few seconds for a new or changed
// element to propagate through the service and be visible in the results of other
// Verified Permissions operations.
//
// [eventually consistent]: https://wikipedia.org/wiki/Eventual_consistency
