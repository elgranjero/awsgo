package secretsmanager

// RotateSecret is generated as a reference stub.
// Executable command wiring lives under cmd/secretsmanager.go.
//
// Configures and starts the asynchronous process of rotating the secret. For
// information about rotation, see [Rotate secrets]in the Secrets Manager User Guide. If you
// include the configuration parameters, the operation sets the values for the
// secret and then immediately starts a rotation. If you don't include the
// configuration parameters, the operation starts a rotation with the values
// already stored in the secret.
//
// When rotation is successful, the AWSPENDING staging label might be attached to
// the same version as the AWSCURRENT version, or it might not be attached to any
// version. If the AWSPENDING staging label is present but not attached to the
// same version as AWSCURRENT , then any later invocation of RotateSecret assumes
// that a previous rotation request is still in progress and returns an error. When
// rotation is unsuccessful, the AWSPENDING staging label might be attached to an
// empty secret version. For more information, see [Troubleshoot rotation]in the Secrets Manager User
// Guide.
//
// Secrets Manager generates a CloudTrail log entry when you call this action. Do
// not include sensitive information in request parameters because it might be
// logged. For more information, see [Logging Secrets Manager events with CloudTrail].
//
// Required permissions: secretsmanager:RotateSecret . For more information, see [IAM policy actions for Secrets Manager]
// and [Authentication and access control in Secrets Manager]. You also need lambda:InvokeFunction permissions on the rotation function.
// For more information, see [Permissions for rotation].
//
// [Authentication and access control in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
// [Permissions for rotation]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotating-secrets-required-permissions-function.html
// [Rotate secrets]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotating-secrets.html
// [Logging Secrets Manager events with CloudTrail]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html
// [Troubleshoot rotation]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/troubleshoot_rotation.html
// [IAM policy actions for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions
