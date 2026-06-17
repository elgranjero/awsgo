package cloudwatch

// SetAlarmState is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatch.go.
//
// Temporarily sets the state of an alarm for testing purposes. When the updated
// state differs from the previous value, the action configured for the appropriate
// state is invoked. For example, if your alarm is configured to send an Amazon SNS
// message when an alarm is triggered, temporarily changing the alarm state to
// ALARM sends an SNS message.
//
// Metric alarms returns to their actual state quickly, often within seconds.
// Because the metric alarm state change happens quickly, it is typically only
// visible in the alarm's History tab in the Amazon CloudWatch console or through [DescribeAlarmHistory].
//
// If you use SetAlarmState on a composite alarm, the composite alarm is not
// guaranteed to return to its actual state. It returns to its actual state only
// once any of its children alarms change state. It is also reevaluated if you
// update its configuration.
//
// If an alarm triggers EC2 Auto Scaling policies or application Auto Scaling
// policies, you must include information in the StateReasonData parameter to
// enable the policy to take the correct action.
//
// [DescribeAlarmHistory]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_DescribeAlarmHistory.html
