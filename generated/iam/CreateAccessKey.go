package iam

// CreateAccessKey is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Creates a new Amazon Web Services secret access key and corresponding Amazon
//
// Web Services access key ID for the specified user. The default status for new
// keys is Active .
//
// If you do not specify a user name, IAM determines the user name implicitly
// based on the Amazon Web Services access key ID signing the request. This
// operation works for access keys under the Amazon Web Services account.
// Consequently, you can use this operation to manage Amazon Web Services account
// root user credentials. This is true even if the Amazon Web Services account has
// no associated users.
//
// For information about quotas on the number of keys you can create, see [IAM and STS quotas] in the
// IAM User Guide.
//
// To ensure the security of your Amazon Web Services account, the secret access
// key is accessible only during key and user creation. You must save the key (for
// example, in a text file) if you want to be able to access it again. If a secret
// key is lost, you can delete the access keys for the associated user and then
// create new keys.
//
// [IAM and STS quotas]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html
