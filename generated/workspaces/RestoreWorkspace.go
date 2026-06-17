package workspaces

// RestoreWorkspace is generated as a reference stub.
// Executable command wiring lives under cmd/workspaces.go.
//
// Restores the specified WorkSpace to its last known healthy state.
//
// You cannot restore a WorkSpace unless its state is  AVAILABLE , ERROR ,
// UNHEALTHY , or STOPPED .
//
// Restoring a WorkSpace is a potentially destructive action that can result in
// the loss of data. For more information, see [Restore a WorkSpace].
//
// This operation is asynchronous and returns before the WorkSpace is completely
// restored.
//
// [Restore a WorkSpace]: https://docs.aws.amazon.com/workspaces/latest/adminguide/restore-workspace.html
