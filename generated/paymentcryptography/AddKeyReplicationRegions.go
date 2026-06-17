package paymentcryptography

// AddKeyReplicationRegions is generated as a reference stub.
// Executable command wiring lives under cmd/paymentcryptography.go.
//
// Adds replication Amazon Web Services Regions to an existing Amazon Web Services
// Payment Cryptography key, enabling the key to be used for cryptographic
// operations in additional Amazon Web Services Regions.
//
// [Multi-Region key replication]allow you to use the same key material across multiple Amazon Web Services
// Regions, providing lower latency for applications distributed across regions.
// When you add Replication Regions, Amazon Web Services Payment Cryptography
// securely replicates the key material to the specified Amazon Web Services
// Regions.
//
// The key must be in an active state to add Replication Regions. You can add
// multiple regions in a single operation, and the key will be available for use in
// those regions once replication is complete.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// [RemoveKeyReplicationRegions]
//
// [EnableDefaultKeyReplicationRegions]
//
// [GetDefaultKeyReplicationRegions]
//
// [RemoveKeyReplicationRegions]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_RemoveKeyReplicationRegions.html
// [Multi-Region key replication]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-multi-region-replication.html
// [EnableDefaultKeyReplicationRegions]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_EnableDefaultKeyReplicationRegions.html
// [GetDefaultKeyReplicationRegions]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_GetDefaultKeyReplicationRegions.html
