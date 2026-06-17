package secretsmanager

// DeleteSecret is generated as a reference stub.
// Executable command wiring lives under cmd/secretsmanager.go.
//
// Deletes a secret and all of its versions. You can specify a recovery window
// during which you can restore the secret. The minimum recovery window is 7 days.
// The default recovery window is 30 days. Secrets Manager attaches a DeletionDate
// stamp to the secret that specifies the end of the recovery window. At the end of
// the recovery window, Secrets Manager deletes the secret permanently.
//
// You can't delete a primary secret that is replicated to other Regions. You must
// first delete the replicas using RemoveRegionsFromReplication, and then delete the primary secret. When you
// delete a replica, it is deleted immediately.
//
// You can't directly delete a version of a secret. Instead, you remove all
// staging labels from the version using UpdateSecretVersionStage. This marks the version as deprecated,
// and then Secrets Manager can automatically delete the version in the background.
//
// To determine whether an application still uses a secret, you can create an
// Amazon CloudWatch alarm to alert you to any attempts to access a secret during
// the recovery window. For more information, see [Monitor secrets scheduled for deletion].
//
// Secrets Manager performs the permanent secret deletion at the end of the
// waiting period as a background task with low priority. There is no guarantee of
// a specific time after the recovery window for the permanent delete to occur.
//
// At any time before recovery window ends, you can use RestoreSecret to remove the DeletionDate
// and cancel the deletion of the secret.
//
// When a secret is scheduled for deletion, you cannot retrieve the secret value.
// You must first cancel the deletion with RestoreSecretand then you can retrieve the secret.
//
// Secrets Manager generates a CloudTrail log entry when you call this action. Do
// not include sensitive information in request parameters because it might be
// logged. For more information, see [Logging Secrets Manager events with CloudTrail].
//
// Required permissions: secretsmanager:DeleteSecret . For more information, see [IAM policy actions for Secrets Manager]
// and [Authentication and access control in Secrets Manager].
//
// [Authentication and access control in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
// [Logging Secrets Manager events with CloudTrail]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html
// [Monitor secrets scheduled for deletion]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/monitoring_cloudwatch_deleted-secrets.html
// [IAM policy actions for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions
