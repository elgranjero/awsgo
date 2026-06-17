package secretsmanager

// StopReplicationToReplica is generated as a reference stub.
// Executable command wiring lives under cmd/secretsmanager.go.
//
// Removes the link between the replica secret and the primary secret and promotes
// the replica to a primary secret in the replica Region.
//
// You must call this operation from the Region in which you want to promote the
// replica to a primary secret.
//
// Secrets Manager generates a CloudTrail log entry when you call this action. Do
// not include sensitive information in request parameters because it might be
// logged. For more information, see [Logging Secrets Manager events with CloudTrail].
//
// Required permissions: secretsmanager:StopReplicationToReplica . For more
// information, see [IAM policy actions for Secrets Manager]and [Authentication and access control in Secrets Manager].
//
// [Authentication and access control in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
// [Logging Secrets Manager events with CloudTrail]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html
// [IAM policy actions for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions
