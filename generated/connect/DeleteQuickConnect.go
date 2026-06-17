package connect

// DeleteQuickConnect is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// Deletes a quick connect.
//
// After calling [DeleteUser], it's important to call DeleteQuickConnect to delete any records
// related to the deleted users. This will help you:
//
// - Avoid dangling resources that impact your service quotas.
//
// - Remove deleted users so they don't appear to agents as transfer options.
//
// - Avoid the disruption of other Amazon Connect processes, such as instance
// replication and syncing if you're using [Amazon Connect Global Resiliency].
//
// [Amazon Connect Global Resiliency]: https://docs.aws.amazon.com/connect/latest/adminguide/setup-connect-global-resiliency.html
// [DeleteUser]: https://docs.aws.amazon.com/connect/latest/APIReference/API_DeleteUser.html
