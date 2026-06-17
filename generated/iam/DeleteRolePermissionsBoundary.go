package iam

// DeleteRolePermissionsBoundary is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Deletes the permissions boundary for the specified IAM role.
//
// You cannot set the boundary for a service-linked role.
//
// Deleting the permissions boundary for a role might increase its permissions.
// For example, it might allow anyone who assumes the role to perform all the
// actions granted in its permissions policies.
