package secretsmanager

// ListSecretVersionIds is generated as a reference stub.
// Executable command wiring lives under cmd/secretsmanager.go.
//
// Lists the versions of a secret. Secrets Manager uses staging labels to indicate
// the different versions of a secret. For more information, see [Secrets Manager concepts: Versions].
//
// To list the secrets in the account, use ListSecrets.
//
// Secrets Manager generates a CloudTrail log entry when you call this action. Do
// not include sensitive information in request parameters because it might be
// logged. For more information, see [Logging Secrets Manager events with CloudTrail].
//
// Required permissions: secretsmanager:ListSecretVersionIds . For more
// information, see [IAM policy actions for Secrets Manager]and [Authentication and access control in Secrets Manager].
//
// [Authentication and access control in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
// [Logging Secrets Manager events with CloudTrail]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html
// [Secrets Manager concepts: Versions]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/getting-started.html#term_version
// [IAM policy actions for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions
