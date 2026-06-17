package secretsmanager

// ValidateResourcePolicy is generated as a reference stub.
// Executable command wiring lives under cmd/secretsmanager.go.
//
// Validates that a resource policy does not grant a wide range of principals
// access to your secret. A resource-based policy is optional for secrets.
//
// The API performs three checks when validating the policy:
//
// - Sends a call to [Zelkova], an automated reasoning engine, to ensure your resource
// policy does not allow broad access to your secret, for example policies that use
// a wildcard for the principal.
//
// - Checks for correct syntax in a policy.
//
// - Verifies the policy does not lock out a caller.
//
// Secrets Manager generates a CloudTrail log entry when you call this action. Do
// not include sensitive information in request parameters because it might be
// logged. For more information, see [Logging Secrets Manager events with CloudTrail].
//
// Required permissions: secretsmanager:ValidateResourcePolicy and
// secretsmanager:PutResourcePolicy . For more information, see [IAM policy actions for Secrets Manager] and [Authentication and access control in Secrets Manager].
//
// [Authentication and access control in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
// [Logging Secrets Manager events with CloudTrail]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html
// [IAM policy actions for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions
// [Zelkova]: https://aws.amazon.com/blogs/security/protect-sensitive-data-in-the-cloud-with-automated-reasoning-zelkova/
