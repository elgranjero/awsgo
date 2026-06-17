package globalaccelerator

// CreateCustomRoutingAccelerator is generated as a reference stub.
// Executable command wiring lives under cmd/globalaccelerator.go.
//
// Create a custom routing accelerator. A custom routing accelerator directs
// traffic to one of possibly thousands of Amazon EC2 instance destinations running
// in a single or multiple virtual private clouds (VPC) subnet endpoints.
//
// Be aware that, by default, all destination EC2 instances in a VPC subnet
// endpoint cannot receive traffic. To enable all destinations to receive traffic,
// or to specify individual port mappings that can receive traffic, see the [AllowCustomRoutingTraffic]
// operation.
//
// Global Accelerator is a global service that supports endpoints in multiple
// Amazon Web Services Regions but you must specify the US West (Oregon) Region to
// create, update, or otherwise work with accelerators. That is, for example,
// specify --region us-west-2 on Amazon Web Services CLI commands.
//
// [AllowCustomRoutingTraffic]: https://docs.aws.amazon.com/global-accelerator/latest/api/API_AllowCustomRoutingTraffic.html
