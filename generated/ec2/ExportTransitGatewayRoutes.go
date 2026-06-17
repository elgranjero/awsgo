package ec2

// ExportTransitGatewayRoutes is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Exports routes from the specified transit gateway route table to the specified
// S3 bucket. By default, all routes are exported. Alternatively, you can filter by
// CIDR range.
//
// The routes are saved to the specified bucket in a JSON file. For more
// information, see [Export route tables to Amazon S3]in the Amazon Web Services Transit Gateways Guide.
//
// [Export route tables to Amazon S3]: https://docs.aws.amazon.com/vpc/latest/tgw/tgw-route-tables.html#tgw-export-route-tables
