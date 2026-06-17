package secretsmanager

// PutSecretValue is generated as a reference stub.
// Executable command wiring lives under cmd/secretsmanager.go.
//
// Creates a new version of your secret by creating a new encrypted value and
// attaching it to the secret. version can contain a new SecretString value or a
// new SecretBinary value.
//
// Do not call PutSecretValue at a sustained rate of more than once every 10
// minutes. When you update the secret value, Secrets Manager creates a new version
// of the secret. Secrets Manager keeps 100 of the most recent versions, but it
// keeps all secret versions created in the last 24 hours. If you call
// PutSecretValue more than once every 10 minutes, you will create more versions
// than Secrets Manager removes, and you will reach the quota for secret versions.
//
// You can specify the staging labels to attach to the new version in VersionStages
// . If you don't include VersionStages , then Secrets Manager automatically moves
// the staging label AWSCURRENT to this version. If this operation creates the
// first version for the secret, then Secrets Manager automatically attaches the
// staging label AWSCURRENT to it. If this operation moves the staging label
// AWSCURRENT from another version to this version, then Secrets Manager also
// automatically moves the staging label AWSPREVIOUS to the version that AWSCURRENT
// was removed from.
//
// This operation is idempotent. If you call this operation with a
// ClientRequestToken that matches an existing version's VersionId, and you specify
// the same secret data, the operation succeeds but does nothing. However, if the
// secret data is different, then the operation fails because you can't modify an
// existing version; you can only create new ones.
//
// Secrets Manager generates a CloudTrail log entry when you call this action. Do
// not include sensitive information in request parameters except SecretBinary ,
// SecretString , or RotationToken because it might be logged. For more
// information, see [Logging Secrets Manager events with CloudTrail].
//
// Required permissions: secretsmanager:PutSecretValue . For more information, see [IAM policy actions for Secrets Manager]
// and [Authentication and access control in Secrets Manager].
//
// When you enter commands in a command shell, there is a risk of the command
// history being accessed or utilities having access to your command parameters.
// This is a concern if the command includes the value of a secret. Learn how to [Mitigate the risks of using command-line tools to store Secrets Manager secrets].
//
// [Authentication and access control in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
// [Logging Secrets Manager events with CloudTrail]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html
// [Mitigate the risks of using command-line tools to store Secrets Manager secrets]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/security_cli-exposure-risks.html
// [IAM policy actions for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions
