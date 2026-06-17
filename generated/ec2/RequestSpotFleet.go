package ec2

// RequestSpotFleet is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Creates a Spot Fleet request.
//
// The Spot Fleet request specifies the total target capacity and the On-Demand
// target capacity. Amazon EC2 calculates the difference between the total capacity
// and On-Demand capacity, and launches the difference as Spot capacity.
//
// You can submit a single request that includes multiple launch specifications
// that vary by instance type, AMI, Availability Zone, or subnet.
//
// By default, the Spot Fleet requests Spot Instances in the Spot Instance pool
// where the price per unit is the lowest. Each launch specification can include
// its own instance weighting that reflects the value of the instance type to your
// application workload.
//
// Alternatively, you can specify that the Spot Fleet distribute the target
// capacity across the Spot pools included in its launch specifications. By
// ensuring that the Spot Instances in your Spot Fleet are in different Spot pools,
// you can improve the availability of your fleet.
//
// You can specify tags for the Spot Fleet request and instances launched by the
// fleet. You cannot tag other resource types in a Spot Fleet request because only
// the spot-fleet-request and instance resource types are supported.
//
// For more information, see [Spot Fleet requests] in the Amazon EC2 User Guide.
//
// We strongly discourage using the RequestSpotFleet API because it is a legacy
// API with no planned investment. For options for requesting Spot Instances, see [Which is the best Spot request method to use?]
// in the Amazon EC2 User Guide.
//
// [Spot Fleet requests]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-fleet-requests.html
// [Which is the best Spot request method to use?]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-best-practices.html#which-spot-request-method-to-use
