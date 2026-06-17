package ec2

// AdvertiseByoipCidr is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Advertises an IPv4 or IPv6 address range that is provisioned for use with your
// Amazon Web Services resources through bring your own IP addresses (BYOIP).
//
// You can perform this operation at most once every 10 seconds, even if you
// specify different address ranges each time.
//
// We recommend that you stop advertising the BYOIP CIDR from other locations when
// you advertise it from Amazon Web Services. To minimize down time, you can
// configure your Amazon Web Services resources to use an address from a BYOIP CIDR
// before it is advertised, and then simultaneously stop advertising it from the
// current location and start advertising it through Amazon Web Services.
//
// It can take a few minutes before traffic to the specified addresses starts
// routing to Amazon Web Services because of BGP propagation delays.
