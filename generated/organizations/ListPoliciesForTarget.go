package organizations

// ListPoliciesForTarget is generated as a reference stub.
// Executable command wiring lives under cmd/organizations.go.
//
// Lists the policies that are directly attached to the specified target root,
// organizational unit (OU), or account. You must specify the policy type that you
// want included in the returned list.
//
// When calling List* operations, always check the NextToken response parameter
// value, even if you receive an empty result set. These operations can
// occasionally return an empty set of results even when more results are
// available. Continue making requests until NextToken returns null. A null
// NextToken value indicates that you have retrieved all available results.
//
// You can only call this operation from the management account or a member
// account that is a delegated administrator.
