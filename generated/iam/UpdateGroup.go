package iam

// UpdateGroup is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Updates the name and/or the path of the specified IAM group.
//
// You should understand the implications of changing a group's path or name. For
// more information, see [Renaming users and groups]in the IAM User Guide.
//
// The person making the request (the principal), must have permission to change
// the role group with the old name and the new name. For example, to change the
// group named Managers to MGRs , the principal must have a policy that allows them
// to update both groups. If the principal has permission to update the Managers
// group, but not the MGRs group, then the update fails. For more information
// about permissions, see [Access management].
//
// [Access management]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access.html
// [Renaming users and groups]: https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_WorkingWithGroupsAndUsers.html
