package iam

// DeleteUser is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Deletes the specified IAM user. Unlike the Amazon Web Services Management
// Console, when you delete a user programmatically, you must delete the items
// attached to the user manually, or the deletion fails. For more information, see [Deleting an IAM user]
// . Before attempting to delete a user, remove the following items:
//
// - Password ([DeleteLoginProfile] )
//
// - Access keys ([DeleteAccessKey] )
//
// - Signing certificate ([DeleteSigningCertificate] )
//
// - SSH public key ([DeleteSSHPublicKey] )
//
// - Git credentials ([DeleteServiceSpecificCredential] )
//
// - Multi-factor authentication (MFA) device ([DeactivateMFADevice] , [DeleteVirtualMFADevice])
//
// - Inline policies ([DeleteUserPolicy] )
//
// - Attached managed policies ([DetachUserPolicy] )
//
// - Group memberships ([RemoveUserFromGroup] )
//
// [DetachUserPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DetachUserPolicy.html
// [DeleteAccessKey]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DeleteAccessKey.html
// [DeleteVirtualMFADevice]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DeleteVirtualMFADevice.html
// [Deleting an IAM user]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_users_manage.html#id_users_deleting_cli
// [DeleteUserPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DeleteUserPolicy.html
// [RemoveUserFromGroup]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_RemoveUserFromGroup.html
// [DeleteLoginProfile]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DeleteLoginProfile.html
// [DeleteServiceSpecificCredential]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DeleteServiceSpecificCredential.html
// [DeleteSigningCertificate]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DeleteSigningCertificate.html
// [DeleteSSHPublicKey]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DeleteSSHPublicKey.html
// [DeactivateMFADevice]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DeactivateMFADevice.html
