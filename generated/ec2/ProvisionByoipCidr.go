package ec2

// ProvisionByoipCidr is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Provisions an IPv4 or IPv6 address range for use with your Amazon Web Services
// resources through bring your own IP addresses (BYOIP) and creates a
// corresponding address pool. After the address range is provisioned, it is ready
// to be advertised.
//
// Amazon Web Services verifies that you own the address range and are authorized
// to advertise it. You must ensure that the address range is registered to you and
// that you created an RPKI ROA to authorize Amazon ASNs 16509 and 14618 to
// advertise the address range. For more information, see [Bring your own IP addresses (BYOIP)]in the Amazon EC2 User
// Guide.
//
// Provisioning an address range is an asynchronous operation, so the call returns
// immediately, but the address range is not ready to use until its status changes
// from pending-provision to provisioned . For more information, see [Onboard your address range].
//
// [Bring your own IP addresses (BYOIP)]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-byoip.html
// [Onboard your address range]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/byoip-onboard.html
