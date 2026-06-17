package globalaccelerator

// AddCustomRoutingEndpoints is generated as a reference stub.
// Executable command wiring lives under cmd/globalaccelerator.go.
//
// Associate a virtual private cloud (VPC) subnet endpoint with your custom
// routing accelerator.
//
// The listener port range must be large enough to support the number of IP
// addresses that can be specified in your subnet. The number of ports required is:
// subnet size times the number of ports per destination EC2 instances. For
// example, a subnet defined as /24 requires a listener port range of at least 255
// ports.
//
// Note: You must have enough remaining listener ports available to map to the
// subnet ports, or the call will fail with a LimitExceededException.
//
// By default, all destinations in a subnet in a custom routing accelerator cannot
// receive traffic. To enable all destinations to receive traffic, or to specify
// individual port mappings that can receive traffic, see the [AllowCustomRoutingTraffic]operation.
//
// [AllowCustomRoutingTraffic]: https://docs.aws.amazon.com/global-accelerator/latest/api/API_AllowCustomRoutingTraffic.html
