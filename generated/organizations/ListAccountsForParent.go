package organizations

// ListAccountsForParent is generated as a reference stub.
// Executable command wiring lives under cmd/organizations.go.
//
// Lists the accounts in an organization that are contained by the specified
// target root or organizational unit (OU). If you specify the root, you get a list
// of all the accounts that aren't in any OU. If you specify an OU, you get a list
// of all the accounts in only that OU and not in any child OUs. To get a list of
// all accounts in the organization, use the ListAccountsoperation.
//
// When calling List* operations, always check the NextToken response parameter
// value, even if you receive an empty result set. These operations can
// occasionally return an empty set of results even when more results are
// available. Continue making requests until NextToken returns null. A null
// NextToken value indicates that you have retrieved all available results.
//
// You can only call this operation from the management account or a member
// account that is a delegated administrator.
