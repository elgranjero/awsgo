package ram

// ReplacePermissionAssociations is generated as a reference stub.
// Executable command wiring lives under cmd/ram.go.
//
// Updates all resource shares that use a managed permission to a different
// managed permission. This operation always applies the default version of the
// target managed permission. You can optionally specify that the update applies to
// only resource shares that currently use a specified version. This enables you to
// update to the latest version, without changing the which managed permission is
// used.
//
// You can use this operation to update all of your resource shares to use the
// current default version of the permission by specifying the same value for the
// fromPermissionArn and toPermissionArn parameters.
//
// You can use the optional fromPermissionVersion parameter to update only those
// resources that use a specified version of the managed permission to the new
// managed permission.
//
// To successfully perform this operation, you must have permission to update the
// resource-based policy on all affected resource types.
