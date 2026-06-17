package autoscaling

// PutScheduledUpdateGroupAction is generated as a reference stub.
// Executable command wiring lives under cmd/autoscaling.go.
//
// Creates or updates a scheduled scaling action for an Auto Scaling group.
//
// For more information, see [Scheduled scaling] in the Amazon EC2 Auto Scaling User Guide.
//
// You can view the scheduled actions for an Auto Scaling group using the [DescribeScheduledActions] API
// call. If you are no longer using a scheduled action, you can delete it by
// calling the [DeleteScheduledAction]API.
//
// If you try to schedule your action in the past, Amazon EC2 Auto Scaling returns
// an error message.
//
// [DeleteScheduledAction]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DeleteScheduledAction.html
// [DescribeScheduledActions]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribeScheduledActions.html
// [Scheduled scaling]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-scheduled-scaling.html
