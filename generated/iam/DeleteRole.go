package iam

// DeleteRole is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Deletes the specified role. Unlike the Amazon Web Services Management Console,
// when you delete a role programmatically, you must delete the items attached to
// the role manually, or the deletion fails. For more information, see [Deleting an IAM role]. Before
// attempting to delete a role, remove the following attached items:
//
// - Inline policies ([DeleteRolePolicy] )
//
// - Attached managed policies ([DetachRolePolicy] )
//
// - Instance profile ([RemoveRoleFromInstanceProfile] )
//
// - Optional – Delete instance profile after detaching from role for resource
// clean up ([DeleteInstanceProfile] )
//
// Make sure that you do not have any Amazon EC2 instances running with the role
// you are about to delete. Deleting a role or instance profile that is associated
// with a running instance will break any applications running on the instance.
//
// [DetachRolePolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DetachRolePolicy.html
// [RemoveRoleFromInstanceProfile]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_RemoveRoleFromInstanceProfile.html
// [DeleteRolePolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DeleteRolePolicy.html
// [DeleteInstanceProfile]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DeleteInstanceProfile.html
// [Deleting an IAM role]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_manage_delete.html#roles-managingrole-deleting-cli
