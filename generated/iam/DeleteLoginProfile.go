package iam

// DeleteLoginProfile is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Deletes the password for the specified IAM user or root user, For more
// information, see [Managing passwords for IAM users].
//
// You can use the CLI, the Amazon Web Services API, or the Users page in the IAM
// console to delete a password for any IAM user. You can use [ChangePassword]to update, but not
// delete, your own password in the My Security Credentials page in the Amazon Web
// Services Management Console.
//
// Deleting a user's password does not prevent a user from accessing Amazon Web
// Services through the command line interface or the API. To prevent all user
// access, you must also either make any access keys inactive or delete them. For
// more information about making keys inactive or deleting them, see [UpdateAccessKey]and [DeleteAccessKey].
//
// [ChangePassword]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_ChangePassword.html
// [DeleteAccessKey]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DeleteAccessKey.html
// [Managing passwords for IAM users]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_passwords_admin-change-user.html
// [UpdateAccessKey]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_UpdateAccessKey.html
