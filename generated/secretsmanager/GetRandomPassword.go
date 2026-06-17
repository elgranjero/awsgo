package secretsmanager

// GetRandomPassword is generated as a reference stub.
// Executable command wiring lives under cmd/secretsmanager.go.
//
// Generates a random password. We recommend that you specify the maximum length
// and include every character type that the system you are generating a password
// for can support. By default, Secrets Manager uses uppercase and lowercase
// letters, numbers, and the following characters in passwords:
// !\"#$%&'()*+,-./:;<=>?(at)[\\]^_`{|}~
//
// Secrets Manager generates a CloudTrail log entry when you call this action.
//
// Required permissions: secretsmanager:GetRandomPassword . For more information,
// see [IAM policy actions for Secrets Manager]and [Authentication and access control in Secrets Manager].
//
// [Authentication and access control in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
// [IAM policy actions for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions
