package secretsmanager

// UntagResource is generated as a reference stub.
// Executable command wiring lives under cmd/secretsmanager.go.
//
// Removes specific tags from a secret.
//
// This operation is idempotent. If a requested tag is not attached to the secret,
// no error is returned and the secret metadata is unchanged.
//
// If you use tags as part of your security strategy, then removing a tag can
// change permissions. If successfully completing this operation would result in
// you losing your permissions for this secret, then the operation is blocked and
// returns an Access Denied error.
//
// Secrets Manager generates a CloudTrail log entry when you call this action. Do
// not include sensitive information in request parameters because it might be
// logged. For more information, see [Logging Secrets Manager events with CloudTrail].
//
// Required permissions: secretsmanager:UntagResource . For more information, see [IAM policy actions for Secrets Manager]
// and [Authentication and access control in Secrets Manager].
//
// [Authentication and access control in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
// [Logging Secrets Manager events with CloudTrail]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html
// [IAM policy actions for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions
