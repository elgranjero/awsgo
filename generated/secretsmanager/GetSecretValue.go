package secretsmanager

// GetSecretValue is generated as a reference stub.
// Executable command wiring lives under cmd/secretsmanager.go.
//
// Retrieves the contents of the encrypted fields SecretString or SecretBinary
// from the specified version of a secret, whichever contains content.
//
// To retrieve the values for a group of secrets, call BatchGetSecretValue.
//
// We recommend that you cache your secret values by using client-side caching.
// Caching secrets improves speed and reduces your costs. For more information, see
// [Cache secrets for your applications].
//
// To retrieve the previous version of a secret, use VersionStage and specify
// AWSPREVIOUS. To revert to the previous version of a secret, call [UpdateSecretVersionStage].
//
// Secrets Manager generates a CloudTrail log entry when you call this action. Do
// not include sensitive information in request parameters because it might be
// logged. For more information, see [Logging Secrets Manager events with CloudTrail].
//
// Required permissions: secretsmanager:GetSecretValue . If the secret is encrypted
// using a customer-managed key instead of the Amazon Web Services managed key
// aws/secretsmanager , then you also need kms:Decrypt permissions for that key.
// For more information, see [IAM policy actions for Secrets Manager]and [Authentication and access control in Secrets Manager].
//
// [Authentication and access control in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
// [Logging Secrets Manager events with CloudTrail]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html
// [UpdateSecretVersionStage]: https://docs.aws.amazon.com/cli/latest/reference/secretsmanager/update-secret-version-stage.html
// [IAM policy actions for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions
// [Cache secrets for your applications]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieving-secrets.html
