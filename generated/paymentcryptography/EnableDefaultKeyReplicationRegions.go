package paymentcryptography

// EnableDefaultKeyReplicationRegions is generated as a reference stub.
// Executable command wiring lives under cmd/paymentcryptography.go.
//
// Enables [Multi-Region key replication] settings for your Amazon Web Services account, causing new keys to be
// automatically replicated to the specified Amazon Web Services Regions when
// created.
//
// When Multi-Region key replication are enabled, any new keys created in your
// account will automatically be replicated to these regions unless you explicitly
// override this behavior during key creation. This simplifies key management for
// applications that operate across multiple regions.
//
// Existing keys are not affected by this operation - only keys created after
// enabling default replication will be automatically replicated.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// [DisableDefaultKeyReplicationRegions]
//
// [GetDefaultKeyReplicationRegions]
//
// [Multi-Region key replication]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-multi-region-replication.html
// [DisableDefaultKeyReplicationRegions]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_DisableDefaultKeyReplicationRegions.html
// [GetDefaultKeyReplicationRegions]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_GetDefaultKeyReplicationRegions.html
