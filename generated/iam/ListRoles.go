package iam

// ListRoles is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Lists the IAM roles that have the specified path prefix. If there are none, the
// operation returns an empty list. For more information about roles, see [IAM roles]in the
// IAM User Guide.
//
// IAM resource-listing operations return a subset of the available attributes for
// the resource. This operation does not return the following attributes, even
// though they are an attribute of the returned object:
//
// - PermissionsBoundary
//
// - RoleLastUsed
//
// - Tags
//
// To view all of the information for a role, see [GetRole].
//
// You can paginate the results using the MaxItems and Marker parameters.
//
// [GetRole]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetRole.html
// [IAM roles]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html
