package iam

// ListSigningCertificates is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Returns information about the signing certificates associated with the
// specified IAM user. If none exists, the operation returns an empty list.
//
// Although each user is limited to a small number of signing certificates, you
// can still paginate the results using the MaxItems and Marker parameters.
//
// If the UserName field is not specified, the user name is determined implicitly
// based on the Amazon Web Services access key ID used to sign the request for this
// operation. This operation works for access keys under the Amazon Web Services
// account. Consequently, you can use this operation to manage Amazon Web Services
// account root user credentials even if the Amazon Web Services account has no
// associated users.
