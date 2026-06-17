package secretsmanager

// TagResource is generated as a reference stub.
// Executable command wiring lives under cmd/secretsmanager.go.
//
// Attaches tags to a secret. Tags consist of a key name and a value. Tags are
// part of the secret's metadata. They are not associated with specific versions of
// the secret. This operation appends tags to the existing list of tags.
//
// For tag quotas and naming restrictions, see [Service quotas for Tagging] in the Amazon Web Services General
// Reference guide.
//
// If you use tags as part of your security strategy, then adding or removing a
// tag can change permissions. If successfully completing this operation would
// result in you losing your permissions for this secret, then the operation is
// blocked and returns an Access Denied error.
//
// Secrets Manager generates a CloudTrail log entry when you call this action. Do
// not include sensitive information in request parameters because it might be
// logged. For more information, see [Logging Secrets Manager events with CloudTrail].
//
// Required permissions: secretsmanager:TagResource . For more information, see [IAM policy actions for Secrets Manager]
// and [Authentication and access control in Secrets Manager].
//
// [Authentication and access control in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
// [Logging Secrets Manager events with CloudTrail]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html
// [IAM policy actions for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions
// [Service quotas for Tagging]: https://docs.aws.amazon.com/general/latest/gr/arg.html#taged-reference-quotas
