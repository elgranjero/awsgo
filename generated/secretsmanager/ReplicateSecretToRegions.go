package secretsmanager

// ReplicateSecretToRegions is generated as a reference stub.
// Executable command wiring lives under cmd/secretsmanager.go.
//
// Replicates the secret to a new Regions. See [Multi-Region secrets].
//
// Secrets Manager generates a CloudTrail log entry when you call this action. Do
// not include sensitive information in request parameters because it might be
// logged. For more information, see [Logging Secrets Manager events with CloudTrail].
//
// Required permissions: secretsmanager:ReplicateSecretToRegions . If the primary
// secret is encrypted with a KMS key other than aws/secretsmanager , you also need
// kms:Decrypt permission to the key. To encrypt the replicated secret with a KMS
// key other than aws/secretsmanager , you need kms:GenerateDataKey and kms:Encrypt
// to the key. For more information, see [IAM policy actions for Secrets Manager]and [Authentication and access control in Secrets Manager].
//
// [Authentication and access control in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
// [Logging Secrets Manager events with CloudTrail]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html
// [Multi-Region secrets]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/create-manage-multi-region-secrets.html
// [IAM policy actions for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions
