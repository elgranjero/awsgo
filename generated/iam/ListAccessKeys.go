package iam

// ListAccessKeys is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Returns information about the access key IDs associated with the specified IAM
// user. If there is none, the operation returns an empty list.
//
// Although each user is limited to a small number of keys, you can still paginate
// the results using the MaxItems and Marker parameters.
//
// If the UserName is not specified, the user name is determined implicitly based
// on the Amazon Web Services access key ID used to sign the request. If a
// temporary access key is used, then UserName is required. If a long-term key is
// assigned to the user, then UserName is not required.
//
// This operation works for access keys under the Amazon Web Services account. If
// the Amazon Web Services account has no associated users, the root user returns
// it's own access key IDs by running this command.
//
// To ensure the security of your Amazon Web Services account, the secret access
// key is accessible only during key and user creation.
