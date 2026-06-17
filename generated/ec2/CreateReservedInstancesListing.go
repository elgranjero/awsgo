package ec2

// CreateReservedInstancesListing is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Creates a listing for Amazon EC2 Standard Reserved Instances to be sold in the
// Reserved Instance Marketplace. You can submit one Standard Reserved Instance
// listing at a time. To get a list of your Standard Reserved Instances, you can
// use the DescribeReservedInstancesoperation.
//
// Only Standard Reserved Instances can be sold in the Reserved Instance
// Marketplace. Convertible Reserved Instances cannot be sold.
//
// The Reserved Instance Marketplace matches sellers who want to resell Standard
// Reserved Instance capacity that they no longer need with buyers who want to
// purchase additional capacity. Reserved Instances bought and sold through the
// Reserved Instance Marketplace work like any other Reserved Instances.
//
// To sell your Standard Reserved Instances, you must first register as a seller
// in the Reserved Instance Marketplace. After completing the registration process,
// you can create a Reserved Instance Marketplace listing of some or all of your
// Standard Reserved Instances, and specify the upfront price to receive for them.
// Your Standard Reserved Instance listings then become available for purchase. To
// view the details of your Standard Reserved Instance listing, you can use the DescribeReservedInstancesListings
// operation.
//
// For more information, see [Sell in the Reserved Instance Marketplace] in the Amazon EC2 User Guide.
//
// [Sell in the Reserved Instance Marketplace]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ri-market-general.html
