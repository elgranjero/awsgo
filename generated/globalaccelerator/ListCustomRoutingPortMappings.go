package globalaccelerator

// ListCustomRoutingPortMappings is generated as a reference stub.
// Executable command wiring lives under cmd/globalaccelerator.go.
//
// Provides a complete mapping from the public accelerator IP address and port to
// destination EC2 instance IP addresses and ports in the virtual public cloud
// (VPC) subnet endpoint for a custom routing accelerator. For each subnet endpoint
// that you add, Global Accelerator creates a new static port mapping for the
// accelerator. The port mappings don't change after Global Accelerator generates
// them, so you can retrieve and cache the full mapping on your servers.
//
// If you remove a subnet from your accelerator, Global Accelerator removes
// (reclaims) the port mappings. If you add a subnet to your accelerator, Global
// Accelerator creates new port mappings (the existing ones don't change). If you
// add or remove EC2 instances in your subnet, the port mappings don't change,
// because the mappings are created when you add the subnet to Global Accelerator.
//
// The mappings also include a flag for each destination denoting which
// destination IP addresses and ports are allowed or denied traffic.
