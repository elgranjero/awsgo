package ec2

// DescribeAddressTransfers is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Describes an Elastic IP address transfer. For more information, see [Transfer Elastic IP addresses] in the
// Amazon VPC User Guide.
//
// When you transfer an Elastic IP address, there is a two-step handshake between
// the source and transfer Amazon Web Services accounts. When the source account
// starts the transfer, the transfer account has seven days to accept the Elastic
// IP address transfer. During those seven days, the source account can view the
// pending transfer by using this action. After seven days, the transfer expires
// and ownership of the Elastic IP address returns to the source account. Accepted
// transfers are visible to the source account for 14 days after the transfers have
// been accepted.
//
// [Transfer Elastic IP addresses]: https://docs.aws.amazon.com/vpc/latest/userguide/vpc-eips.html#transfer-EIPs-intro
