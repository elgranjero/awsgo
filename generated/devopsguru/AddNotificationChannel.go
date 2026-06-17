package devopsguru

// AddNotificationChannel is generated as a reference stub.
// Executable command wiring lives under cmd/devopsguru.go.
//
// Adds a notification channel to DevOps Guru. A notification channel is used to
//
// notify you about important DevOps Guru events, such as when an insight is
// generated.
//
// If you use an Amazon SNS topic in another account, you must attach a policy to
// it that grants DevOps Guru permission to send it notifications. DevOps Guru adds
// the required policy on your behalf to send notifications using Amazon SNS in
// your account. DevOps Guru only supports standard SNS topics. For more
// information, see [Permissions for Amazon SNS topics].
//
// If you use an Amazon SNS topic that is encrypted by an Amazon Web Services Key
// Management Service customer-managed key (CMK), then you must add permissions to
// the CMK. For more information, see [Permissions for Amazon Web Services KMS–encrypted Amazon SNS topics].
//
// [Permissions for Amazon SNS topics]: https://docs.aws.amazon.com/devops-guru/latest/userguide/sns-required-permissions.html
// [Permissions for Amazon Web Services KMS–encrypted Amazon SNS topics]: https://docs.aws.amazon.com/devops-guru/latest/userguide/sns-kms-permissions.html
