package ec2

// ReplaceRouteTableAssociation is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Changes the route table associated with a given subnet, internet gateway, or
// virtual private gateway in a VPC. After the operation completes, the subnet or
// gateway uses the routes in the new route table. For more information about route
// tables, see [Route tables]in the Amazon VPC User Guide.
//
// You can also use this operation to change which table is the main route table
// in the VPC. Specify the main route table's association ID and the route table ID
// of the new main route table.
//
// [Route tables]: https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Route_Tables.html
