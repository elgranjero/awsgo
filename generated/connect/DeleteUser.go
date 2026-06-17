package connect

// DeleteUser is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// Deletes a user account from the specified Amazon Connect instance.
//
// For information about what happens to a user's data when their account is
// deleted, see [Delete Users from Your Amazon Connect Instance]in the Amazon Connect Administrator Guide.
//
// After calling DeleteUser, call [DeleteQuickConnect] to delete any records related to the deleted
// users. This will help you:
//
// - Avoid dangling resources that impact your service quotas.
//
// - Remove deleted users so they don't appear to agents as transfer options.
//
// - Avoid the disruption of other Amazon Connect processes, such as instance
// replication and syncing if you're using [Amazon Connect Global Resiliency].
//
// [Delete Users from Your Amazon Connect Instance]: https://docs.aws.amazon.com/connect/latest/adminguide/delete-users.html
// [Amazon Connect Global Resiliency]: https://docs.aws.amazon.com/connect/latest/adminguide/setup-connect-global-resiliency.html
// [DeleteQuickConnect]: https://docs.aws.amazon.com/connect/latest/APIReference/API_DeleteQuickConnect.html
