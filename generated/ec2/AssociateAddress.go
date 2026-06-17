package ec2

// AssociateAddress is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Associates an Elastic IP address, or carrier IP address (for instances that are
// in subnets in Wavelength Zones) with an instance or a network interface. Before
// you can use an Elastic IP address, you must allocate it to your account.
//
// If the Elastic IP address is already associated with a different instance, it
// is disassociated from that instance and associated with the specified instance.
// If you associate an Elastic IP address with an instance that has an existing
// Elastic IP address, the existing address is disassociated from the instance, but
// remains allocated to your account.
//
// [Subnets in Wavelength Zones] You can associate an IP address from the
// telecommunication carrier to the instance or network interface.
//
// You cannot associate an Elastic IP address with an interface in a different
// network border group.
//
// This is an idempotent operation. If you perform the operation more than once,
// Amazon EC2 doesn't return an error, and you may be charged for each time the
// Elastic IP address is remapped to the same instance. For more information, see
// the Elastic IP Addresses section of [Amazon EC2 Pricing].
//
// [Amazon EC2 Pricing]: http://aws.amazon.com/ec2/pricing/
