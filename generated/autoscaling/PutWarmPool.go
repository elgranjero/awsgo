package autoscaling

// PutWarmPool is generated as a reference stub.
// Executable command wiring lives under cmd/autoscaling.go.
//
// Creates or updates a warm pool for the specified Auto Scaling group. A warm
// pool is a pool of pre-initialized EC2 instances that sits alongside the Auto
// Scaling group. Whenever your application needs to scale out, the Auto Scaling
// group can draw on the warm pool to meet its new desired capacity.
//
// This operation must be called from the Region in which the Auto Scaling group
// was created.
//
// You can view the instances in the warm pool using the [DescribeWarmPool] API call. If you are no
// longer using a warm pool, you can delete it by calling the [DeleteWarmPool]API.
//
// For more information, see [Warm pools for Amazon EC2 Auto Scaling] in the Amazon EC2 Auto Scaling User Guide.
//
// [DeleteWarmPool]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DeleteWarmPool.html
// [DescribeWarmPool]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribeWarmPool.html
// [Warm pools for Amazon EC2 Auto Scaling]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-warm-pools.html
