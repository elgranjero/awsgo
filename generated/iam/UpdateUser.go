package iam

// UpdateUser is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Updates the name and/or the path of the specified IAM user.
//
// You should understand the implications of changing an IAM user's path or name.
// For more information, see [Renaming an IAM user]and [Renaming an IAM group] in the IAM User Guide.
//
// To change a user name, the requester must have appropriate permissions on both
// the source object and the target object. For example, to change Bob to Robert,
// the entity making the request must have permission on Bob and Robert, or must
// have permission on all (*). For more information about permissions, see [Permissions and policies].
//
// [Renaming an IAM user]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_users_manage.html#id_users_renaming
// [Renaming an IAM group]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_groups_manage_rename.html
// [Permissions and policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/PermissionsAndPolicies.html
