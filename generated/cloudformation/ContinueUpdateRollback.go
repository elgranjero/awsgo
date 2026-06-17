package cloudformation

// ContinueUpdateRollback is generated as a reference stub.
// Executable command wiring lives under cmd/cloudformation.go.
//
// Continues rolling back a stack from UPDATE_ROLLBACK_FAILED to
// UPDATE_ROLLBACK_COMPLETE state. Depending on the cause of the failure, you can
// manually fix the error and continue the rollback. By continuing the rollback,
// you can return your stack to a working state (the UPDATE_ROLLBACK_COMPLETE
// state) and then try to update the stack again.
//
// A stack enters the UPDATE_ROLLBACK_FAILED state when CloudFormation can't roll
// back all changes after a failed stack update. For example, this might occur when
// a stack attempts to roll back to an old database that was deleted outside of
// CloudFormation. Because CloudFormation doesn't know the instance was deleted, it
// assumes the instance still exists and attempts to roll back to it, causing the
// update rollback to fail.
//
// For more information, see [Continue rolling back an update] in the CloudFormation User Guide. For information
// for troubleshooting a failed update rollback, see [Update rollback failed].
//
// [Continue rolling back an update]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-continueupdaterollback.html
// [Update rollback failed]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/troubleshooting.html#troubleshooting-errors-update-rollback-failed
