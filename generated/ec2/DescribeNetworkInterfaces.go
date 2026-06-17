package ec2

// DescribeNetworkInterfaces is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Describes the specified network interfaces or all your network interfaces.
//
// If you have a large number of network interfaces, the operation fails unless
// you use pagination or one of the following filters: group-id , mac-address ,
// private-dns-name , private-ip-address , subnet-id , or vpc-id .
//
// We strongly recommend using only paginated requests. Unpaginated requests are
// susceptible to throttling and timeouts.
