package secretsmanager

// UpdateSecret is generated as a reference stub.
// Executable command wiring lives under cmd/secretsmanager.go.
//
// Modifies the details of a secret, including metadata and the secret value. To
// change the secret value, you can also use PutSecretValue.
//
// To change the rotation configuration of a secret, use RotateSecret instead.
//
// To change a secret so that it is managed by another service, you need to
// recreate the secret in that service. See [Secrets Manager secrets managed by other Amazon Web Services services].
//
// We recommend you avoid calling UpdateSecret at a sustained rate of more than
// once every 10 minutes. When you call UpdateSecret to update the secret value,
// Secrets Manager creates a new version of the secret. Secrets Manager removes
// outdated versions when there are more than 100, but it does not remove versions
// created less than 24 hours ago. If you update the secret value more than once
// every 10 minutes, you create more versions than Secrets Manager removes, and you
// will reach the quota for secret versions.
//
// If you include SecretString or SecretBinary to create a new secret version,
// Secrets Manager automatically moves the staging label AWSCURRENT to the new
// version. Then it attaches the label AWSPREVIOUS to the version that AWSCURRENT
// was removed from.
//
// If you call this operation with a ClientRequestToken that matches an existing
// version's VersionId , the operation results in an error. You can't modify an
// existing version, you can only create a new version. To remove a version, remove
// all staging labels from it. See UpdateSecretVersionStage.
//
// Secrets Manager generates a CloudTrail log entry when you call this action. Do
// not include sensitive information in request parameters except SecretBinary or
// SecretString because it might be logged. For more information, see [Logging Secrets Manager events with CloudTrail].
//
// Required permissions: secretsmanager:UpdateSecret . For more information, see [IAM policy actions for Secrets Manager]
// and [Authentication and access control in Secrets Manager]. If you use a customer managed key, you must also have kms:GenerateDataKey
// , kms:Encrypt , and kms:Decrypt permissions on the key. If you change the KMS
// key and you don't have kms:Encrypt permission to the new key, Secrets Manager
// does not re-encrypt existing secret versions with the new key. For more
// information, see [Secret encryption and decryption].
//
// When you enter commands in a command shell, there is a risk of the command
// history being accessed or utilities having access to your command parameters.
// This is a concern if the command includes the value of a secret. Learn how to [Mitigate the risks of using command-line tools to store Secrets Manager secrets].
//
// [Authentication and access control in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
// [Logging Secrets Manager events with CloudTrail]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html
// [Secret encryption and decryption]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/security-encryption.html
// [Secrets Manager secrets managed by other Amazon Web Services services]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/service-linked-secrets.html
// [Mitigate the risks of using command-line tools to store Secrets Manager secrets]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/security_cli-exposure-risks.html
// [IAM policy actions for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions
