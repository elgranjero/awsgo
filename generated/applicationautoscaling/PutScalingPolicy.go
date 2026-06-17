package applicationautoscaling

// PutScalingPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/applicationautoscaling.go.
//
// Creates or updates a scaling policy for an Application Auto Scaling scalable
// target.
//
// Each scalable target is identified by a service namespace, resource ID, and
// scalable dimension. A scaling policy applies to the scalable target identified
// by those three attributes. You cannot create a scaling policy until you have
// registered the resource as a scalable target.
//
// Multiple scaling policies can be in force at the same time for the same
// scalable target. You can have one or more target tracking scaling policies, one
// or more step scaling policies, or both. However, there is a chance that multiple
// policies could conflict, instructing the scalable target to scale out or in at
// the same time. Application Auto Scaling gives precedence to the policy that
// provides the largest capacity for both scale out and scale in. For example, if
// one policy increases capacity by 3, another policy increases capacity by 200
// percent, and the current capacity is 10, Application Auto Scaling uses the
// policy with the highest calculated capacity (200% of 10 = 20) and scales out to
// 30.
//
// We recommend caution, however, when using target tracking scaling policies with
// step scaling policies because conflicts between these policies can cause
// undesirable behavior. For example, if the step scaling policy initiates a
// scale-in activity before the target tracking policy is ready to scale in, the
// scale-in activity will not be blocked. After the scale-in activity completes,
// the target tracking policy could instruct the scalable target to scale out
// again.
//
// For more information, see [Target tracking scaling policies], [Step scaling policies], and [Predictive scaling policies] in the Application Auto Scaling User Guide.
//
// If a scalable target is deregistered, the scalable target is no longer
// available to use scaling policies. Any scaling policies that were specified for
// the scalable target are deleted.
//
// [Step scaling policies]: https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-step-scaling-policies.html
// [Predictive scaling policies]: https://docs.aws.amazon.com/autoscaling/application/userguide/aas-create-predictive-scaling-policy.html
// [Target tracking scaling policies]: https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-target-tracking.html
