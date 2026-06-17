package ec2

// ReleaseAddress is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Releases the specified Elastic IP address.
//
// [Default VPC] Releasing an Elastic IP address automatically disassociates it
// from any instance that it's associated with. Alternatively, you can disassociate
// an Elastic IP address without releasing it.
//
// [Nondefault VPC] You must disassociate the Elastic IP address before you can
// release it. Otherwise, Amazon EC2 returns an error ( InvalidIPAddress.InUse ).
//
// After releasing an Elastic IP address, it is released to the IP address pool.
// Be sure to update your DNS records and any servers or devices that communicate
// with the address. If you attempt to release an Elastic IP address that you
// already released, you'll get an AuthFailure error if the address is already
// allocated to another Amazon Web Services account.
//
// After you release an Elastic IP address, you might be able to recover it. For
// more information, see [Release an Elastic IP address].
//
// [Release an Elastic IP address]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-instance-addressing-eips-releasing.html
