package ec2

// ModifySubnetAttribute is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Modifies a subnet attribute. You can only modify one attribute at a time.
//
// Use this action to modify subnets on Amazon Web Services Outposts.
//
// - To modify a subnet on an Outpost rack, set both MapCustomerOwnedIpOnLaunch
// and CustomerOwnedIpv4Pool . These two parameters act as a single attribute.
//
// - To modify a subnet on an Outpost server, set either EnableLniAtDeviceIndex
// or DisableLniAtDeviceIndex .
//
// For more information about Amazon Web Services Outposts, see the following:
//
// [Outpost servers]
//
// [Outpost racks]
//
// [Outpost servers]: https://docs.aws.amazon.com/outposts/latest/userguide/how-servers-work.html
// [Outpost racks]: https://docs.aws.amazon.com/outposts/latest/userguide/how-racks-work.html
