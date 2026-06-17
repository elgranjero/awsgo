package autoscaling

// CancelInstanceRefresh is generated as a reference stub.
// Executable command wiring lives under cmd/autoscaling.go.
//
// Cancels an instance refresh or rollback that is in progress. If an instance
// refresh or rollback is not in progress, an ActiveInstanceRefreshNotFound error
// occurs.
//
// This operation is part of the [instance refresh feature] in Amazon EC2 Auto Scaling, which helps you
// update instances in your Auto Scaling group after you make configuration
// changes.
//
// When you cancel an instance refresh, this does not roll back any changes that
// it made. Use the [RollbackInstanceRefresh]API to roll back instead.
//
// [instance refresh feature]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-instance-refresh.html
// [RollbackInstanceRefresh]: https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_RollbackInstanceRefresh.html
