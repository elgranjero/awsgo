package iam

// ListSSHPublicKeys is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Returns information about the SSH public keys associated with the specified IAM
// user. If none exists, the operation returns an empty list.
//
// The SSH public keys returned by this operation are used only for authenticating
// the IAM user to an CodeCommit repository. For more information about using SSH
// keys to authenticate to an CodeCommit repository, see [Set up CodeCommit for SSH connections]in the CodeCommit User
// Guide.
//
// Although each user is limited to a small number of keys, you can still paginate
// the results using the MaxItems and Marker parameters.
//
// [Set up CodeCommit for SSH connections]: https://docs.aws.amazon.com/codecommit/latest/userguide/setting-up-credentials-ssh.html
