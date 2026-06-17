package autoscaling

// DescribeAutoScalingGroups is generated as a reference stub.
// Executable command wiring lives under cmd/autoscaling.go.
//
// Gets information about the Auto Scaling groups in the account and Region.
//
// If you specify Auto Scaling group names, the output includes information for
// only the specified Auto Scaling groups. If you specify filters, the output
// includes information for only those Auto Scaling groups that meet the filter
// criteria. If you do not specify group names or filters, the output includes
// information for all Auto Scaling groups.
//
// This operation also returns information about instances in Auto Scaling groups.
// To retrieve information about the instances in a warm pool, you must call the [DescribeWarmPool]
// API.
//
// [DescribeWarmPool]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribeWarmPool.html
