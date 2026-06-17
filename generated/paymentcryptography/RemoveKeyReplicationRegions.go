package paymentcryptography

// RemoveKeyReplicationRegions is generated as a reference stub.
// Executable command wiring lives under cmd/paymentcryptography.go.
//
// Removes Replication Regions from an existing Amazon Web Services Payment
// Cryptography key, disabling the key's availability for cryptographic operations
// in the specified Amazon Web Services Regions.
//
// When you remove Replication Regions, the key material is securely deleted from
// those regions and can no longer be used for cryptographic operations there. This
// operation is irreversible for the specified Amazon Web Services Regions. For
// more information, see [Multi-Region key replication].
//
// Ensure that no active cryptographic operations or applications depend on the
// key in the regions you're removing before performing this operation.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// [AddKeyReplicationRegions]
//
// [DisableDefaultKeyReplicationRegions]
//
// [Multi-Region key replication]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-multi-region-replication.html
// [DisableDefaultKeyReplicationRegions]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_DisableDefaultKeyReplicationRegions.html
// [AddKeyReplicationRegions]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_AddKeyReplicationRegions.html
