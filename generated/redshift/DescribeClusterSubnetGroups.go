package redshift

// DescribeClusterSubnetGroups is generated as a reference stub.
// Executable command wiring lives under cmd/redshift.go.
//
// Returns one or more cluster subnet group objects, which contain metadata about
// your cluster subnet groups. By default, this operation returns information about
// all cluster subnet groups that are defined in your Amazon Web Services account.
//
// If you specify both tag keys and tag values in the same request, Amazon
// Redshift returns all subnet groups that match any combination of the specified
// keys and values. For example, if you have owner and environment for tag keys,
// and admin and test for tag values, all subnet groups that have any combination
// of those values are returned.
//
// If both tag keys and values are omitted from the request, subnet groups are
// returned regardless of whether they have tag keys or values associated with
// them.
