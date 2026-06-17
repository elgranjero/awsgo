package secretsmanager

// UpdateSecretVersionStage is generated as a reference stub.
// Executable command wiring lives under cmd/secretsmanager.go.
//
// Modifies the staging labels attached to a version of a secret. Secrets Manager
// uses staging labels to track a version as it progresses through the secret
// rotation process. Each staging label can be attached to only one version at a
// time. To add a staging label to a version when it is already attached to another
// version, Secrets Manager first removes it from the other version first and then
// attaches it to this one. For more information about versions and staging labels,
// see [Concepts: Version].
//
// The staging labels that you specify in the VersionStage parameter are added to
// the existing list of staging labels for the version.
//
// You can move the AWSCURRENT staging label to this version by including it in
// this call.
//
// Whenever you move AWSCURRENT , Secrets Manager automatically moves the label
// AWSPREVIOUS to the version that AWSCURRENT was removed from.
//
// If this action results in the last label being removed from a version, then the
// version is considered to be 'deprecated' and can be deleted by Secrets Manager.
//
// Secrets Manager generates a CloudTrail log entry when you call this action. Do
// not include sensitive information in request parameters because it might be
// logged. For more information, see [Logging Secrets Manager events with CloudTrail].
//
// Required permissions: secretsmanager:UpdateSecretVersionStage . For more
// information, see [IAM policy actions for Secrets Manager]and [Authentication and access control in Secrets Manager].
//
// [Authentication and access control in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
// [Logging Secrets Manager events with CloudTrail]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html
// [Concepts: Version]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/getting-started.html#term_version
// [IAM policy actions for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions
