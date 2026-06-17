package workspaces

// DeregisterWorkspaceDirectory is generated as a reference stub.
// Executable command wiring lives under cmd/workspaces.go.
//
// Deregisters the specified directory. This operation is asynchronous and returns
// before the WorkSpace directory is deregistered. If any WorkSpaces are registered
// to this directory, you must remove them before you can deregister the directory.
//
// Simple AD and AD Connector are made available to you free of charge to use with
// WorkSpaces. If there are no WorkSpaces being used with your Simple AD or AD
// Connector directory for 30 consecutive days, this directory will be
// automatically deregistered for use with Amazon WorkSpaces, and you will be
// charged for this directory as per the [Directory Service pricing terms].
//
// To delete empty directories, see [Delete the Directory for Your WorkSpaces]. If you delete your Simple AD or AD Connector
// directory, you can always create a new one when you want to start using
// WorkSpaces again.
//
// [Delete the Directory for Your WorkSpaces]: https://docs.aws.amazon.com/workspaces/latest/adminguide/delete-workspaces-directory.html
// [Directory Service pricing terms]: http://aws.amazon.com/directoryservice/pricing/
