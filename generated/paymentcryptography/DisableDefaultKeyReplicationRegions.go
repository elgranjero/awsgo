package paymentcryptography

// DisableDefaultKeyReplicationRegions is generated as a reference stub.
// Executable command wiring lives under cmd/paymentcryptography.go.
//
// Disables [Multi-Region key replication] settings for the specified Amazon Web Services Regions in your Amazon
// Web Services account, preventing new keys from being automatically replicated to
// those regions.
//
// After disabling Multi-Region key replication for specific regions, new keys
// created in your account will not be automatically replicated to those regions.
// You can still manually add replication to those regions for individual keys
// using the [AddKeyReplicationRegions]operation.
//
// This operation does not affect existing keys or their current replication
// configuration.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// [EnableDefaultKeyReplicationRegions]
//
// [GetDefaultKeyReplicationRegions]
//
// [Multi-Region key replication]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-multi-region-replication.html
// [AddKeyReplicationRegions]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_AddKeyReplicationRegions.html
// [EnableDefaultKeyReplicationRegions]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_EnableDefaultKeyReplicationRegions.html
// [GetDefaultKeyReplicationRegions]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_GetDefaultKeyReplicationRegions.html
