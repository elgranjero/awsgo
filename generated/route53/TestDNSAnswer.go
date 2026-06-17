package route53

// TestDNSAnswer is generated as a reference stub.
// Executable command wiring lives under cmd/route53.go.
//
// Gets the value that Amazon Route 53 returns in response to a DNS request for a
// specified record name and type. You can optionally specify the IP address of a
// DNS resolver, an EDNS0 client subnet IP address, and a subnet mask.
//
// This call only supports querying public hosted zones.
//
// The TestDnsAnswer  returns information similar to what you would expect from
// the answer section of the dig command. Therefore, if you query for the name
// servers of a subdomain that point to the parent name servers, those will not be
// returned.
