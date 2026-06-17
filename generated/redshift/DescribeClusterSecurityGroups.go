package redshift

// DescribeClusterSecurityGroups is generated as a reference stub.
// Executable command wiring lives under cmd/redshift.go.
//
// Returns information about Amazon Redshift security groups. If the name of a
// security group is specified, the response will contain only information about
// only that security group.
//
// For information about managing security groups, go to [Amazon Redshift Cluster Security Groups] in the Amazon Redshift
// Cluster Management Guide.
//
// If you specify both tag keys and tag values in the same request, Amazon
// Redshift returns all security groups that match any combination of the specified
// keys and values. For example, if you have owner and environment for tag keys,
// and admin and test for tag values, all security groups that have any
// combination of those values are returned.
//
// If both tag keys and values are omitted from the request, security groups are
// returned regardless of whether they have tag keys or values associated with
// them.
//
// [Amazon Redshift Cluster Security Groups]: https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-security-groups.html
