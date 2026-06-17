package ssm

// UpdateResourceDataSync is generated as a reference stub.
// Executable command wiring lives under cmd/ssm.go.
//
// Update a resource data sync. After you create a resource data sync for a
// Region, you can't change the account options for that sync. For example, if you
// create a sync in the us-east-2 (Ohio) Region and you choose the Include only
// the current account option, you can't edit that sync later and choose the
// Include all accounts from my Organizations configuration option. Instead, you
// must delete the first resource data sync, and create a new one.
//
// This API operation only supports a resource data sync that was created with a
// SyncFromSource SyncType .
