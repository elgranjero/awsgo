package secretsmanager

// CancelRotateSecret is generated as a reference stub.
// Executable command wiring lives under cmd/secretsmanager.go.
//
// Turns off automatic rotation, and if a rotation is currently in progress,
// cancels the rotation.
//
// If you cancel a rotation in progress, it can leave the VersionStage labels in
// an unexpected state. You might need to remove the staging label AWSPENDING from
// the partially created version. You also need to determine whether to roll back
// to the previous version of the secret by moving the staging label AWSCURRENT to
// the version that has AWSPENDING . To determine which version has a specific
// staging label, call ListSecretVersionIds. Then use UpdateSecretVersionStage to change staging labels. For more information,
// see [How rotation works].
//
// To turn on automatic rotation again, call RotateSecret.
//
// Secrets Manager generates a CloudTrail log entry when you call this action. Do
// not include sensitive information in request parameters because it might be
// logged. For more information, see [Logging Secrets Manager events with CloudTrail].
//
// Required permissions: secretsmanager:CancelRotateSecret . For more information,
// see [IAM policy actions for Secrets Manager]and [Authentication and access control in Secrets Manager].
//
// [Authentication and access control in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
// [Logging Secrets Manager events with CloudTrail]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html
// [How rotation works]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotate-secrets_how.html
// [IAM policy actions for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions
