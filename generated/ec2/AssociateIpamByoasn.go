package ec2

// AssociateIpamByoasn is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Associates your Autonomous System Number (ASN) with a BYOIP CIDR that you own
// in the same Amazon Web Services Region. For more information, see [Tutorial: Bring your ASN to IPAM]in the Amazon
// VPC IPAM guide.
//
// After the association succeeds, the ASN is eligible for advertisement. You can
// view the association with [DescribeByoipCidrs]. You can advertise the CIDR with [AdvertiseByoipCidr].
//
// [DescribeByoipCidrs]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeByoipCidrs.html
// [AdvertiseByoipCidr]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AdvertiseByoipCidr.html
// [Tutorial: Bring your ASN to IPAM]: https://docs.aws.amazon.com/vpc/latest/ipam/tutorials-byoasn.html
