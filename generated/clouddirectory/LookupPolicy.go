package clouddirectory

// LookupPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/clouddirectory.go.
//
// Lists all policies from the root of the Directory to the object specified. If there are
// no policies present, an empty list is returned. If policies are present, and if
// some objects don't have the policies attached, it returns the ObjectIdentifier
// for such objects. If policies are present, it returns ObjectIdentifier ,
// policyId , and policyType . Paths that don't lead to the root from the target
// object are ignored. For more information, see [Policies].
//
// [Policies]: https://docs.aws.amazon.com/clouddirectory/latest/developerguide/key_concepts_directory.html#key_concepts_policies
