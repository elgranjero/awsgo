package iam

// UpdateAccessKey is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Changes the status of the specified access key from Active to Inactive, or vice
// versa. This operation can be used to disable a user's key as part of a key
// rotation workflow.
//
// If the UserName is not specified, the user name is determined implicitly based
// on the Amazon Web Services access key ID used to sign the request. If a
// temporary access key is used, then UserName is required. If a long-term key is
// assigned to the user, then UserName is not required. This operation works for
// access keys under the Amazon Web Services account. Consequently, you can use
// this operation to manage Amazon Web Services account root user credentials even
// if the Amazon Web Services account has no associated users.
//
// For information about rotating keys, see [Managing keys and certificates] in the IAM User Guide.
//
// [Managing keys and certificates]: https://docs.aws.amazon.com/IAM/latest/UserGuide/ManagingCredentials.html
