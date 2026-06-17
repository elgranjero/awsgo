package kms

// DisconnectCustomKeyStore is generated as a reference stub.
// Executable command wiring lives under cmd/kms.go.
//
// Disconnects the [custom key store] from its backing key store. This operation disconnects an
// CloudHSM key store from its associated CloudHSM cluster or disconnects an
// external key store from the external key store proxy that communicates with your
// external key manager.
//
// This operation is part of the custom key stores feature in KMS, which combines
// the convenience and extensive integration of KMS with the isolation and control
// of a key store that you own and manage.
//
// While a custom key store is disconnected, you can manage the custom key store
// and its KMS keys, but you cannot create or use its KMS keys. You can reconnect
// the custom key store at any time.
//
// While a custom key store is disconnected, all attempts to create KMS keys in
// the custom key store or to use existing KMS keys in [cryptographic operations]will fail. This action can
// prevent users from storing and accessing sensitive data.
//
// When you disconnect a custom key store, its ConnectionState changes to
// Disconnected . To find the connection state of a custom key store, use the DescribeCustomKeyStores
// operation. To reconnect a custom key store, use the ConnectCustomKeyStoreoperation.
//
// If the operation succeeds, it returns a JSON object with no properties.
//
// Cross-account use: No. You cannot perform this operation on a custom key store
// in a different Amazon Web Services account.
//
// Required permissions: [kms:DisconnectCustomKeyStore] (IAM policy)
//
// Related operations:
//
// # ConnectCustomKeyStore
//
// # CreateCustomKeyStore
//
// # DeleteCustomKeyStore
//
// # DescribeCustomKeyStores
//
// # UpdateCustomKeyStore
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [cryptographic operations]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-cryptography.html#cryptographic-operations
// [kms:DisconnectCustomKeyStore]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [custom key store]: https://docs.aws.amazon.com/kms/latest/developerguide/key-store-overview.html
