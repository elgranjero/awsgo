package gamelift

// DescribeScalingPolicies is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2
//
// Retrieves all scaling policies applied to a fleet.
//
// To get a fleet's scaling policies, specify the fleet ID. You can filter this
// request by policy status, such as to retrieve only active scaling policies. Use
// the pagination parameters to retrieve results as a set of sequential pages. If
// successful, set of ScalingPolicy objects is returned for the fleet.
//
// A fleet may have all of its scaling policies suspended. This operation does not
// affect the status of the scaling policies, which remains ACTIVE.
