package ec2

// DescribeRouteTables is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Describes your route tables. The default is to describe all your route tables.
// Alternatively, you can specify specific route table IDs or filter the results to
// include only the route tables that match specific criteria.
//
// Each subnet in your VPC must be associated with a route table. If a subnet is
// not explicitly associated with any route table, it is implicitly associated with
// the main route table. This command does not return the subnet ID for implicit
// associations.
//
// For more information, see [Route tables] in the Amazon VPC User Guide.
//
// [Route tables]: https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Route_Tables.html
