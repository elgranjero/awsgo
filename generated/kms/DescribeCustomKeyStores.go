package kms

// DescribeCustomKeyStores is generated as a reference stub.
// Executable command wiring lives under cmd/kms.go.
//
// Gets information about [custom key stores] in the account and Region.
//
// This operation is part of the custom key stores feature in KMS, which combines
// the convenience and extensive integration of KMS with the isolation and control
// of a key store that you own and manage.
//
// By default, this operation returns information about all custom key stores in
// the account and Region. To get only information about a particular custom key
// store, use either the CustomKeyStoreName or CustomKeyStoreId parameter (but not
// both).
//
// To determine whether the custom key store is connected to its CloudHSM cluster
// or external key store proxy, use the ConnectionState element in the response.
// If an attempt to connect the custom key store failed, the ConnectionState value
// is FAILED and the ConnectionErrorCode element in the response indicates the
// cause of the failure. For help interpreting the ConnectionErrorCode , see CustomKeyStoresListEntry.
//
// Custom key stores have a DISCONNECTED connection state if the key store has
// never been connected or you used the DisconnectCustomKeyStoreoperation to disconnect it. Otherwise, the
// connection state is CONNECTED. If your custom key store connection state is
// CONNECTED but you are having trouble using it, verify that the backing store is
// active and available. For an CloudHSM key store, verify that the associated
// CloudHSM cluster is active and contains the minimum number of HSMs required for
// the operation, if any. For an external key store, verify that the external key
// store proxy and its associated external key manager are reachable and enabled.
//
// For help repairing your CloudHSM key store, see the [Troubleshooting CloudHSM key stores]. For help repairing your
// external key store, see the [Troubleshooting external key stores]. Both topics are in the Key Management Service
// Developer Guide.
//
// Cross-account use: No. You cannot perform this operation on a custom key store
// in a different Amazon Web Services account.
//
// Required permissions: [kms:DescribeCustomKeyStores] (IAM policy)
//
// Related operations:
//
// # ConnectCustomKeyStore
//
// # CreateCustomKeyStore
//
// # DeleteCustomKeyStore
//
// # DisconnectCustomKeyStore
//
// # UpdateCustomKeyStore
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [kms:DescribeCustomKeyStores]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [custom key stores]: https://docs.aws.amazon.com/kms/latest/developerguide/key-store-overview.html
// [Troubleshooting CloudHSM key stores]: https://docs.aws.amazon.com/kms/latest/developerguide/fix-keystore.html
// [Troubleshooting external key stores]: https://docs.aws.amazon.com/kms/latest/developerguide/xks-troubleshooting.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
