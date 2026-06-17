package iam

// ListUsers is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Lists the IAM users that have the specified path prefix. If no path prefix is
// specified, the operation returns all users in the Amazon Web Services account.
// If there are none, the operation returns an empty list.
//
// IAM resource-listing operations return a subset of the available attributes for
// the resource. This operation does not return the following attributes, even
// though they are an attribute of the returned object:
//
// - PermissionsBoundary
//
// - Tags
//
// To view all of the information for a user, see [GetUser].
//
// You can paginate the results using the MaxItems and Marker parameters.
//
// [GetUser]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetUser.html
