package autoscaling

// CreateLaunchConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/autoscaling.go.
//
// Creates a launch configuration.
//
// If you exceed your maximum limit of launch configurations, the call fails. To
// query this limit, call the [DescribeAccountLimits]API. For information about updating this limit, see [Quotas for Amazon EC2 Auto Scaling]
// in the Amazon EC2 Auto Scaling User Guide.
//
// For more information, see [Launch configurations] in the Amazon EC2 Auto Scaling User Guide.
//
// Amazon EC2 Auto Scaling configures instances launched as part of an Auto
// Scaling group using either a launch template or a launch configuration. We
// strongly recommend that you do not use launch configurations. They do not
// provide full functionality for Amazon EC2 Auto Scaling or Amazon EC2. For
// information about using launch templates, see [Launch templates]in the Amazon EC2 Auto Scaling
// User Guide.
//
// [DescribeAccountLimits]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribeAccountLimits.html
// [Quotas for Amazon EC2 Auto Scaling]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-quotas.html
// [Launch configurations]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/launch-configurations.html
// [Launch templates]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/launch-templates.html
