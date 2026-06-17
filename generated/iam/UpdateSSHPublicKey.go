package iam

// UpdateSSHPublicKey is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Sets the status of an IAM user's SSH public key to active or inactive. SSH
// public keys that are inactive cannot be used for authentication. This operation
// can be used to disable a user's SSH public key as part of a key rotation work
// flow.
//
// The SSH public key affected by this operation is used only for authenticating
// the associated IAM user to an CodeCommit repository. For more information about
// using SSH keys to authenticate to an CodeCommit repository, see [Set up CodeCommit for SSH connections]in the
// CodeCommit User Guide.
//
// [Set up CodeCommit for SSH connections]: https://docs.aws.amazon.com/codecommit/latest/userguide/setting-up-credentials-ssh.html
