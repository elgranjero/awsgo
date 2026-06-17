package workspaces

// RebuildWorkspaces is generated as a reference stub.
// Executable command wiring lives under cmd/workspaces.go.
//
// Rebuilds the specified WorkSpace.
//
// You cannot rebuild a WorkSpace unless its state is AVAILABLE , ERROR , UNHEALTHY
// , STOPPED , or REBOOTING .
//
// Rebuilding a WorkSpace is a potentially destructive action that can result in
// the loss of data. For more information, see [Rebuild a WorkSpace].
//
// This operation is asynchronous and returns before the WorkSpaces have been
// completely rebuilt.
//
// [Rebuild a WorkSpace]: https://docs.aws.amazon.com/workspaces/latest/adminguide/reset-workspace.html
