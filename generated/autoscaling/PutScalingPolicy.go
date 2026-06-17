package autoscaling

// PutScalingPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/autoscaling.go.
//
// Creates or updates a scaling policy for an Auto Scaling group. Scaling policies
// are used to scale an Auto Scaling group based on configurable metrics. If no
// policies are defined, the dynamic scaling and predictive scaling features are
// not used.
//
// For more information about using dynamic scaling, see [Target tracking scaling policies] and [Step and simple scaling policies] in the Amazon EC2
// Auto Scaling User Guide.
//
// For more information about using predictive scaling, see [Predictive scaling for Amazon EC2 Auto Scaling] in the Amazon EC2
// Auto Scaling User Guide.
//
// You can view the scaling policies for an Auto Scaling group using the [DescribePolicies] API
// call. If you are no longer using a scaling policy, you can delete it by calling
// the [DeletePolicy]API.
//
// [Step and simple scaling policies]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-scaling-simple-step.html
// [DeletePolicy]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DeletePolicy.html
// [Target tracking scaling policies]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-scaling-target-tracking.html
// [DescribePolicies]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribePolicies.html
// [Predictive scaling for Amazon EC2 Auto Scaling]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-predictive-scaling.html
