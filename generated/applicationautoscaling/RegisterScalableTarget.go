package applicationautoscaling

// RegisterScalableTarget is generated as a reference stub.
// Executable command wiring lives under cmd/applicationautoscaling.go.
//
// Registers or updates a scalable target, which is the resource that you want to
// scale.
//
// Scalable targets are uniquely identified by the combination of resource ID,
// scalable dimension, and namespace, which represents some capacity dimension of
// the underlying service.
//
// When you register a new scalable target, you must specify values for the
// minimum and maximum capacity. If the specified resource is not active in the
// target service, this operation does not change the resource's current capacity.
// Otherwise, it changes the resource's current capacity to a value that is inside
// of this range.
//
// If you add a scaling policy, current capacity is adjustable within the
// specified range when scaling starts. Application Auto Scaling scaling policies
// will not scale capacity to values that are outside of the minimum and maximum
// range.
//
// After you register a scalable target, you do not need to register it again to
// use other Application Auto Scaling operations. To see which resources have been
// registered, use [DescribeScalableTargets]. You can also view the scaling policies for a service
// namespace by using [DescribeScalableTargets]. If you no longer need a scalable target, you can
// deregister it by using [DeregisterScalableTarget].
//
// To update a scalable target, specify the parameters that you want to change.
// Include the parameters that identify the scalable target: resource ID, scalable
// dimension, and namespace. Any parameters that you don't specify are not changed
// by this update request.
//
// If you call the RegisterScalableTarget API operation to create a scalable
// target, there might be a brief delay until the operation achieves [eventual consistency]. You might
// become aware of this brief delay if you get unexpected errors when performing
// sequential operations. The typical strategy is to retry the request, and some
// Amazon Web Services SDKs include automatic backoff and retry logic.
//
// If you call the RegisterScalableTarget API operation to update an existing
// scalable target, Application Auto Scaling retrieves the current capacity of the
// resource. If it's below the minimum capacity or above the maximum capacity,
// Application Auto Scaling adjusts the capacity of the scalable target to place it
// within these bounds, even if you don't include the MinCapacity or MaxCapacity
// request parameters.
//
// [DescribeScalableTargets]: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_DescribeScalableTargets.html
// [eventual consistency]: https://en.wikipedia.org/wiki/Eventual_consistency
// [DeregisterScalableTarget]: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_DeregisterScalableTarget.html
