package ec2

// DeleteVpcEndpoints is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Deletes the specified VPC endpoints.
//
// When you delete a gateway endpoint, we delete the endpoint routes in the route
// tables for the endpoint.
//
// When you delete a Gateway Load Balancer endpoint, we delete its endpoint
// network interfaces. You can only delete Gateway Load Balancer endpoints when the
// routes that are associated with the endpoint are deleted.
//
// When you delete an interface endpoint, we delete its endpoint network
// interfaces.
