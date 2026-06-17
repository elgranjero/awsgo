package organizations

// ListRoots is generated as a reference stub.
// Executable command wiring lives under cmd/organizations.go.
//
// Lists the roots that are defined in the current organization.
//
// When calling List* operations, always check the NextToken response parameter
// value, even if you receive an empty result set. These operations can
// occasionally return an empty set of results even when more results are
// available. Continue making requests until NextToken returns null. A null
// NextToken value indicates that you have retrieved all available results.
//
// You can only call this operation from the management account or a member
// account that is a delegated administrator.
//
// Policy types can be enabled and disabled in roots. This is distinct from
// whether they're available in the organization. When you enable all features, you
// make policy types available for use in that organization. Individual policy
// types can then be enabled and disabled in a root. To see the availability of a
// policy type in an organization, use DescribeOrganization.
