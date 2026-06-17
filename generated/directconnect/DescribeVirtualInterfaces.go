package directconnect

// DescribeVirtualInterfaces is generated as a reference stub.
// Executable command wiring lives under cmd/directconnect.go.
//
// Displays all virtual interfaces for an Amazon Web Services account. Virtual
// interfaces deleted fewer than 15 minutes before you make the request are also
// returned. If you specify a connection ID, only the virtual interfaces associated
// with the connection are returned. If you specify a virtual interface ID, then
// only a single virtual interface is returned.
//
// A virtual interface (VLAN) transmits the traffic between the Direct Connect
// location and the customer network.
//
// - If you're using an asn , the response includes ASN value in both the asn and
// asnLong fields.
//
// - If you're using asnLong , the response returns a value of 0 (zero) for the
// asn attribute because it exceeds the highest ASN value of 2,147,483,647 that
// it can support
