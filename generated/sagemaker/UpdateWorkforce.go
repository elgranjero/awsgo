package sagemaker

// UpdateWorkforce is generated as a reference stub.
// Executable command wiring lives under cmd/sagemaker.go.
//
// Use this operation to update your workforce. You can use this operation to
// require that workers use specific IP addresses to work on tasks and to update
// your OpenID Connect (OIDC) Identity Provider (IdP) workforce configuration.
//
// The worker portal is now supported in VPC and public internet.
//
// Use SourceIpConfig to restrict worker access to tasks to a specific range of IP
// addresses. You specify allowed IP addresses by creating a list of up to ten [CIDRs].
// By default, a workforce isn't restricted to specific IP addresses. If you
// specify a range of IP addresses, workers who attempt to access tasks using any
// IP address outside the specified range are denied and get a Not Found error
// message on the worker portal.
//
// To restrict public internet access for all workers, configure the SourceIpConfig
// CIDR value. For example, when using SourceIpConfig with an IpAddressType of IPv4
// , you can restrict access to the IPv4 CIDR block "10.0.0.0/16". When using an
// IpAddressType of dualstack , you can specify both the IPv4 and IPv6 CIDR blocks,
// such as "10.0.0.0/16" for IPv4 only, "2001:db8:1234:1a00::/56" for IPv6 only, or
// "10.0.0.0/16" and "2001:db8:1234:1a00::/56" for dual stack.
//
// Amazon SageMaker does not support Source Ip restriction for worker portals in
// VPC.
//
// Use OidcConfig to update the configuration of a workforce created using your
// own OIDC IdP.
//
// You can only update your OIDC IdP configuration when there are no work teams
// associated with your workforce. You can delete work teams using the [DeleteWorkteam]operation.
//
// After restricting access to a range of IP addresses or updating your OIDC IdP
// configuration with this operation, you can view details about your update
// workforce using the [DescribeWorkforce]operation.
//
// This operation only applies to private workforces.
//
// [DescribeWorkforce]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeWorkforce.html
// [DeleteWorkteam]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DeleteWorkteam.html
// [CIDRs]: https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html
