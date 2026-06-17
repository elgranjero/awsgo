package secretsmanager

// RestoreSecret is generated as a reference stub.
// Executable command wiring lives under cmd/secretsmanager.go.
//
// Cancels the scheduled deletion of a secret by removing the DeletedDate time
// stamp. You can access a secret again after it has been restored.
//
// Secrets Manager generates a CloudTrail log entry when you call this action. Do
// not include sensitive information in request parameters because it might be
// logged. For more information, see [Logging Secrets Manager events with CloudTrail].
//
// Required permissions: secretsmanager:RestoreSecret . For more information, see [IAM policy actions for Secrets Manager]
// and [Authentication and access control in Secrets Manager].
//
// [Authentication and access control in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
// [Logging Secrets Manager events with CloudTrail]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html
// [IAM policy actions for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions
