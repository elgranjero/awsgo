package ec2

// AssociateDhcpOptions is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Associates a set of DHCP options (that you've previously created) with the
// specified VPC, or associates no DHCP options with the VPC.
//
// After you associate the options with the VPC, any existing instances and all
// new instances that you launch in that VPC use the options. You don't need to
// restart or relaunch the instances. They automatically pick up the changes within
// a few hours, depending on how frequently the instance renews its DHCP lease. You
// can explicitly renew the lease using the operating system on the instance.
//
// For more information, see [DHCP option sets] in the Amazon VPC User Guide.
//
// [DHCP option sets]: https://docs.aws.amazon.com/vpc/latest/userguide/VPC_DHCP_Options.html
