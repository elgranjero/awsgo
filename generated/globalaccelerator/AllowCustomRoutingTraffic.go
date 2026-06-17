package globalaccelerator

// AllowCustomRoutingTraffic is generated as a reference stub.
// Executable command wiring lives under cmd/globalaccelerator.go.
//
// Specify the Amazon EC2 instance (destination) IP addresses and ports for a VPC
// subnet endpoint that can receive traffic for a custom routing accelerator. You
// can allow traffic to all destinations in the subnet endpoint, or allow traffic
// to a specified list of destination IP addresses and ports in the subnet. Note
// that you cannot specify IP addresses or ports outside of the range that you
// configured for the endpoint group.
//
// After you make changes, you can verify that the updates are complete by
// checking the status of your accelerator: the status changes from IN_PROGRESS to
// DEPLOYED.
