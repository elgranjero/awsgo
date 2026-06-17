package workspaces

// TerminateWorkspaces is generated as a reference stub.
// Executable command wiring lives under cmd/workspaces.go.
//
// Terminates the specified WorkSpaces.
//
// Terminating a WorkSpace is a permanent action and cannot be undone. The user's
// data is destroyed. If you need to archive any user data, contact Amazon Web
// Services Support before terminating the WorkSpace.
//
// You can terminate a WorkSpace that is in any state except SUSPENDED .
//
// This operation is asynchronous and returns before the WorkSpaces have been
// completely terminated. After a WorkSpace is terminated, the TERMINATED state is
// returned only briefly before the WorkSpace directory metadata is cleaned up, so
// this state is rarely returned. To confirm that a WorkSpace is terminated, check
// for the WorkSpace ID by using [DescribeWorkSpaces]. If the WorkSpace ID isn't returned, then the
// WorkSpace has been successfully terminated.
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
// [DescribeWorkSpaces]: https://docs.aws.amazon.com/workspaces/latest/api/API_DescribeWorkspaces.html
// [Delete the Directory for Your WorkSpaces]: https://docs.aws.amazon.com/workspaces/latest/adminguide/delete-workspaces-directory.html
// [Directory Service pricing terms]: http://aws.amazon.com/directoryservice/pricing/
