package autoscaling

// DescribeInstanceRefreshes is generated as a reference stub.
// Executable command wiring lives under cmd/autoscaling.go.
//
// Gets information about the instance refreshes for the specified Auto Scaling
// group from the previous six weeks.
//
// This operation is part of the [instance refresh feature] in Amazon EC2 Auto Scaling, which helps you
// update instances in your Auto Scaling group after you make configuration
// changes.
//
// To help you determine the status of an instance refresh, Amazon EC2 Auto
// Scaling returns information about the instance refreshes you previously
// initiated, including their status, start time, end time, the percentage of the
// instance refresh that is complete, and the number of instances remaining to
// update before the instance refresh is complete. If a rollback is initiated while
// an instance refresh is in progress, Amazon EC2 Auto Scaling also returns
// information about the rollback of the instance refresh.
//
// [instance refresh feature]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-instance-refresh.html
