package organizations

// ListChildren is generated as a reference stub.
// Executable command wiring lives under cmd/organizations.go.
//
// Lists all of the organizational units (OUs) or accounts that are contained in
// the specified parent OU or root. This operation, along with ListParentsenables you to
// traverse the tree structure that makes up this root.
//
// When calling List* operations, always check the NextToken response parameter
// value, even if you receive an empty result set. These operations can
// occasionally return an empty set of results even when more results are
// available. Continue making requests until NextToken returns null. A null
// NextToken value indicates that you have retrieved all available results.
//
// You can only call this operation from the management account or a member
// account that is a delegated administrator.
