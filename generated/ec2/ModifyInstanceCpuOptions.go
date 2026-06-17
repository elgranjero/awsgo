package ec2

// ModifyInstanceCpuOptions is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// By default, all vCPUs for the instance type are active when you launch an
// instance. When you configure the number of active vCPUs for the instance, it can
// help you save on licensing costs and optimize performance. The base cost of the
// instance remains unchanged.
//
// The number of active vCPUs equals the number of threads per CPU core multiplied
// by the number of cores. The instance must be in a Stopped state before you make
// changes.
//
// Some instance type options do not support this capability. For more
// information, see [Supported CPU options]in the Amazon EC2 User Guide.
//
// [Supported CPU options]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/cpu-options-supported-instances-values.html
