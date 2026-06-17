package directconnect

// CreateBGPPeer is generated as a reference stub.
// Executable command wiring lives under cmd/directconnect.go.
//
// Creates a BGP peer on the specified virtual interface.
//
// You must create a BGP peer for the corresponding address family (IPv4/IPv6) in
// order to access Amazon Web Services resources that also use that address family.
//
// If logical redundancy is not supported by the connection, interconnect, or LAG,
// the BGP peer cannot be in the same address family as an existing BGP peer on the
// virtual interface.
//
// When creating a IPv6 BGP peer, omit the Amazon address and customer address.
// IPv6 addresses are automatically assigned from the Amazon pool of IPv6
// addresses; you cannot specify custom IPv6 addresses.
//
// If you let Amazon Web Services auto-assign IPv4 addresses, a /30 CIDR will be
// allocated from 169.254.0.0/16. Amazon Web Services does not recommend this
// option if you intend to use the customer router peer IP address as the source
// and destination for traffic. Instead you should use RFC 1918 or other
// addressing, and specify the address yourself. For more information about RFC
// 1918 see [Address Allocation for Private Internets].
//
// For a public virtual interface, the Autonomous System Number (ASN) must be
// private or already on the allow list for the virtual interface.
//
// [Address Allocation for Private Internets]: https://datatracker.ietf.org/doc/html/rfc1918
