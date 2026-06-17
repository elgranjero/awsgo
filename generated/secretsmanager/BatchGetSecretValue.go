package secretsmanager

// BatchGetSecretValue is generated as a reference stub.
// Executable command wiring lives under cmd/secretsmanager.go.
//
// Retrieves the contents of the encrypted fields SecretString or SecretBinary for
// up to 20 secrets. To retrieve a single secret, call GetSecretValue.
//
// To choose which secrets to retrieve, you can specify a list of secrets by name
// or ARN, or you can use filters. If Secrets Manager encounters errors such as
// AccessDeniedException while attempting to retrieve any of the secrets, you can
// see the errors in Errors in the response.
//
// Secrets Manager generates CloudTrail GetSecretValue log entries for each secret
// you request when you call this action. Do not include sensitive information in
// request parameters because it might be logged. For more information, see [Logging Secrets Manager events with CloudTrail].
//
// Required permissions: secretsmanager:BatchGetSecretValue , and you must have
// secretsmanager:GetSecretValue for each secret. If you use filters, you must also
// have secretsmanager:ListSecrets . If the secrets are encrypted using
// customer-managed keys instead of the Amazon Web Services managed key
// aws/secretsmanager , then you also need kms:Decrypt permissions for the keys.
// For more information, see [IAM policy actions for Secrets Manager]and [Authentication and access control in Secrets Manager].
//
// [Authentication and access control in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
// [Logging Secrets Manager events with CloudTrail]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html
// [IAM policy actions for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions
