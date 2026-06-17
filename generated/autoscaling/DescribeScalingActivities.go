package autoscaling

// DescribeScalingActivities is generated as a reference stub.
// Executable command wiring lives under cmd/autoscaling.go.
//
// Gets information about the scaling activities in the account and Region.
//
// When scaling events occur, you see a record of the scaling activity in the
// scaling activities. For more information, see [Verify a scaling activity for an Auto Scaling group]in the Amazon EC2 Auto Scaling
// User Guide.
//
// If the scaling event succeeds, the value of the StatusCode element in the
// response is Successful . If an attempt to launch instances failed, the
// StatusCode value is Failed or Cancelled and the StatusMessage element in the
// response indicates the cause of the failure. For help interpreting the
// StatusMessage , see [Troubleshooting Amazon EC2 Auto Scaling] in the Amazon EC2 Auto Scaling User Guide.
//
// [Troubleshooting Amazon EC2 Auto Scaling]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/CHAP_Troubleshooting.html
// [Verify a scaling activity for an Auto Scaling group]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-verify-scaling-activity.html
