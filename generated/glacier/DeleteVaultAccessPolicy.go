package glacier

// DeleteVaultAccessPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/glacier.go.
//
// This operation deletes the access policy associated with the specified vault.
// The operation is eventually consistent; that is, it might take some time for
// Amazon Glacier to completely remove the access policy, and you might still see
// the effect of the policy for a short time after you send the delete request.
//
// This operation is idempotent. You can invoke delete multiple times, even if
// there is no policy associated with the vault. For more information about vault
// access policies, see [Amazon Glacier Access Control with Vault Access Policies].
//
// [Amazon Glacier Access Control with Vault Access Policies]: https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-access-policy.html
