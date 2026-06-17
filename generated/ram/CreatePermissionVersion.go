package ram

// CreatePermissionVersion is generated as a reference stub.
// Executable command wiring lives under cmd/ram.go.
//
// Creates a new version of the specified customer managed permission. The new
// version is automatically set as the default version of the customer managed
// permission. New resource shares automatically use the default permission.
// Existing resource shares continue to use their original permission versions, but
// you can use ReplacePermissionAssociationsto update them.
//
// If the specified customer managed permission already has the maximum of 5
// versions, then you must delete one of the existing versions before you can
// create a new one.
