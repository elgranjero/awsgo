package networkfirewall

// DeleteVpcEndpointAssociation is generated as a reference stub.
// Executable command wiring lives under cmd/networkfirewall.go.
//
// Deletes the specified VpcEndpointAssociation.
//
// You can check whether an endpoint association is in use by reviewing the route
// tables for the Availability Zones where you have the endpoint subnet mapping.
// You can retrieve the subnet mapping by calling DescribeVpcEndpointAssociation. You define and update the
// route tables through Amazon VPC. As needed, update the route tables for the
// Availability Zone to remove the firewall endpoint for the association. When the
// route tables no longer use the firewall endpoint, you can remove the endpoint
// association safely.
