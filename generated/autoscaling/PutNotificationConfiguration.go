package autoscaling

// PutNotificationConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/autoscaling.go.
//
// Configures an Auto Scaling group to send notifications when specified events
// take place. Subscribers to the specified topic can have messages delivered to an
// endpoint such as a web server or an email address.
//
// This configuration overwrites any existing configuration.
//
// For more information, see [Amazon SNS notification options for Amazon EC2 Auto Scaling] in the Amazon EC2 Auto Scaling User Guide.
//
// If you exceed your maximum limit of SNS topics, which is 10 per Auto Scaling
// group, the call fails.
//
// [Amazon SNS notification options for Amazon EC2 Auto Scaling]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-sns-notifications.html
