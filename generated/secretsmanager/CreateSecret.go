package secretsmanager

// CreateSecret is generated as a reference stub.
// Executable command wiring lives under cmd/secretsmanager.go.
//
// Creates a new secret. A secret can be a password, a set of credentials such as
// a user name and password, an OAuth token, or other secret information that you
// store in an encrypted form in Secrets Manager. The secret also includes the
// connection information to access a database or other service, which Secrets
// Manager doesn't encrypt. A secret in Secrets Manager consists of both the
// protected secret data and the important information needed to manage the secret.
//
// For secrets that use managed rotation, you need to create the secret through
// the managing service. For more information, see [Secrets Manager secrets managed by other Amazon Web Services services].
//
// For information about creating a secret in the console, see [Create a secret].
//
// To create a secret, you can provide the secret value to be encrypted in either
// the SecretString parameter or the SecretBinary parameter, but not both. If you
// include SecretString or SecretBinary then Secrets Manager creates an initial
// secret version and automatically attaches the staging label AWSCURRENT to it.
//
// For database credentials you want to rotate, for Secrets Manager to be able to
// rotate the secret, you must make sure the JSON you store in the SecretString
// matches the [JSON structure of a database secret].
//
// If you don't specify an KMS encryption key, Secrets Manager uses the Amazon Web
// Services managed key aws/secretsmanager . If this key doesn't already exist in
// your account, then Secrets Manager creates it for you automatically. All users
// and roles in the Amazon Web Services account automatically have access to use
// aws/secretsmanager . Creating aws/secretsmanager can result in a one-time
// significant delay in returning the result.
//
// If the secret is in a different Amazon Web Services account from the
// credentials calling the API, then you can't use aws/secretsmanager to encrypt
// the secret, and you must create and use a customer managed KMS key.
//
// Secrets Manager generates a CloudTrail log entry when you call this action. Do
// not include sensitive information in request parameters except SecretBinary or
// SecretString because it might be logged. For more information, see [Logging Secrets Manager events with CloudTrail].
//
// Required permissions: secretsmanager:CreateSecret . If you include tags in the
// secret, you also need secretsmanager:TagResource . To add replica Regions, you
// must also have secretsmanager:ReplicateSecretToRegions . For more information,
// see [IAM policy actions for Secrets Manager]and [Authentication and access control in Secrets Manager].
//
// To encrypt the secret with a KMS key other than aws/secretsmanager , you need
// kms:GenerateDataKey and kms:Decrypt permission to the key.
//
// When you enter commands in a command shell, there is a risk of the command
// history being accessed or utilities having access to your command parameters.
// This is a concern if the command includes the value of a secret. Learn how to [Mitigate the risks of using command-line tools to store Secrets Manager secrets].
//
// [Authentication and access control in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
// [Logging Secrets Manager events with CloudTrail]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html
// [Secrets Manager secrets managed by other Amazon Web Services services]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/service-linked-secrets.html
// [Mitigate the risks of using command-line tools to store Secrets Manager secrets]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/security_cli-exposure-risks.html
// [Create a secret]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/manage_create-basic-secret.html
// [IAM policy actions for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions
// [JSON structure of a database secret]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_secret_json_structure.html
