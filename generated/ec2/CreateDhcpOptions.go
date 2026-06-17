package ec2

// CreateDhcpOptions is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Creates a custom set of DHCP options. After you create a DHCP option set, you
// associate it with a VPC. After you associate a DHCP option set with a VPC, all
// existing and newly launched instances in the VPC use this set of DHCP options.
//
// The following are the individual DHCP options you can specify. For more
// information, see [DHCP option sets]in the Amazon VPC User Guide.
//
// - domain-name - If you're using AmazonProvidedDNS in us-east-1 , specify
// ec2.internal . If you're using AmazonProvidedDNS in any other Region, specify
// region.compute.internal . Otherwise, specify a custom domain name. This value
// is used to complete unqualified DNS hostnames.
//
// Some Linux operating systems accept multiple domain names separated by spaces.
//
// However, Windows and other Linux operating systems treat the value as a single
// domain, which results in unexpected behavior. If your DHCP option set is
// associated with a VPC that has instances running operating systems that treat
// the value as a single domain, specify only one domain name.
//
// - domain-name-servers - The IP addresses of up to four DNS servers, or
// AmazonProvidedDNS. To specify multiple domain name servers in a single
// parameter, separate the IP addresses using commas. To have your instances
// receive custom DNS hostnames as specified in domain-name , you must specify a
// custom DNS server.
//
// - ntp-servers - The IP addresses of up to eight Network Time Protocol (NTP)
// servers (four IPv4 addresses and four IPv6 addresses).
//
// - netbios-name-servers - The IP addresses of up to four NetBIOS name servers.
//
// - netbios-node-type - The NetBIOS node type (1, 2, 4, or 8). We recommend that
// you specify 2. Broadcast and multicast are not supported. For more information
// about NetBIOS node types, see [RFC 2132].
//
// - ipv6-address-preferred-lease-time - A value (in seconds, minutes, hours, or
// years) for how frequently a running instance with an IPv6 assigned to it goes
// through DHCPv6 lease renewal. Acceptable values are between 140 and 2147483647
// seconds (approximately 68 years). If no value is entered, the default lease time
// is 140 seconds. If you use long-term addressing for EC2 instances, you can
// increase the lease time and avoid frequent lease renewal requests. Lease renewal
// typically occurs when half of the lease time has elapsed.
//
// [DHCP option sets]: https://docs.aws.amazon.com/vpc/latest/userguide/VPC_DHCP_Options.html
//
// [RFC 2132]: https://www.ietf.org/rfc/rfc2132.txt
