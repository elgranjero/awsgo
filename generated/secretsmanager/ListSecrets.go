package secretsmanager

// ListSecrets is generated as a reference stub.
// Executable command wiring lives under cmd/secretsmanager.go.
//
// Lists the secrets that are stored by Secrets Manager in the Amazon Web Services
// account, not including secrets that are marked for deletion. To see secrets
// marked for deletion, use the Secrets Manager console.
//
// All Secrets Manager operations are eventually consistent. ListSecrets might not
// reflect changes from the last five minutes. You can get more recent information
// for a specific secret by calling DescribeSecret.
//
// To list the versions of a secret, use ListSecretVersionIds.
//
// To retrieve the values for the secrets, call BatchGetSecretValue or GetSecretValue.
//
// For information about finding secrets in the console, see [Find secrets in Secrets Manager].
//
// Secrets Manager generates a CloudTrail log entry when you call this action. Do
// not include sensitive information in request parameters because it might be
// logged. For more information, see [Logging Secrets Manager events with CloudTrail].
//
// Required permissions: secretsmanager:ListSecrets . For more information, see [IAM policy actions for Secrets Manager]
// and [Authentication and access control in Secrets Manager].
//
// [Authentication and access control in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
// [Logging Secrets Manager events with CloudTrail]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html
// [IAM policy actions for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions
// [Find secrets in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/manage_search-secret.html
