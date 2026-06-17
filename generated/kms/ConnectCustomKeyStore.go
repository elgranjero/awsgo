package kms

// ConnectCustomKeyStore is generated as a reference stub.
// Executable command wiring lives under cmd/kms.go.
//
// Connects or reconnects a [custom key store] to its backing key store. For an CloudHSM key store,
// ConnectCustomKeyStore connects the key store to its associated CloudHSM cluster.
// For an external key store, ConnectCustomKeyStore connects the key store to the
// external key store proxy that communicates with your external key manager.
//
// The custom key store must be connected before you can create KMS keys in the
// key store or use the KMS keys it contains. You can disconnect and reconnect a
// custom key store at any time.
//
// The connection process for a custom key store can take an extended amount of
// time to complete. This operation starts the connection process, but it does not
// wait for it to complete. When it succeeds, this operation quickly returns an
// HTTP 200 response and a JSON object with no properties. However, this response
// does not indicate that the custom key store is connected. To get the connection
// state of the custom key store, use the DescribeCustomKeyStoresoperation.
//
// This operation is part of the custom key stores feature in KMS, which combines
// the convenience and extensive integration of KMS with the isolation and control
// of a key store that you own and manage.
//
// The ConnectCustomKeyStore operation might fail for various reasons. To find the
// reason, use the DescribeCustomKeyStoresoperation and see the ConnectionErrorCode in the response. For
// help interpreting the ConnectionErrorCode , see CustomKeyStoresListEntry.
//
// To fix the failure, use the DisconnectCustomKeyStore operation to disconnect the custom key store,
// correct the error, use the UpdateCustomKeyStoreoperation if necessary, and then use
// ConnectCustomKeyStore again.
//
// # CloudHSM key store
//
// During the connection process for an CloudHSM key store, KMS finds the CloudHSM
// cluster that is associated with the custom key store, creates the connection
// infrastructure, connects to the cluster, logs into the CloudHSM client as the
// kmsuser CU, and rotates its password.
//
// To connect an CloudHSM key store, its associated CloudHSM cluster must have at
// least one active HSM. To get the number of active HSMs in a cluster, use the [DescribeClusters]
// operation. To add HSMs to the cluster, use the [CreateHsm]operation. Also, the [kmsuser crypto user]kmsuser
// (CU) must not be logged into the cluster. This prevents KMS from using this
// account to log in.
//
// If you are having trouble connecting or disconnecting a CloudHSM key store, see [Troubleshooting an CloudHSM key store]
// in the Key Management Service Developer Guide.
//
// # External key store
//
// When you connect an external key store that uses public endpoint connectivity,
// KMS tests its ability to communicate with your external key manager by sending a
// request via the external key store proxy.
//
// When you connect to an external key store that uses VPC endpoint service
// connectivity, KMS establishes the networking elements that it needs to
// communicate with your external key manager via the external key store proxy.
// This includes creating an interface endpoint to the VPC endpoint service and a
// private hosted zone for traffic between KMS and the VPC endpoint service.
//
// To connect an external key store, KMS must be able to connect to the external
// key store proxy, the external key store proxy must be able to communicate with
// your external key manager, and the external key manager must be available for
// cryptographic operations.
//
// If you are having trouble connecting or disconnecting an external key store,
// see [Troubleshooting an external key store]in the Key Management Service Developer Guide.
//
// Cross-account use: No. You cannot perform this operation on a custom key store
// in a different Amazon Web Services account.
//
// Required permissions: [kms:ConnectCustomKeyStore] (IAM policy)
//
// # Related operations
//
// # CreateCustomKeyStore
//
// # DeleteCustomKeyStore
//
// # DescribeCustomKeyStores
//
// # DisconnectCustomKeyStore
//
// # UpdateCustomKeyStore
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [DescribeClusters]: https://docs.aws.amazon.com/cloudhsm/latest/APIReference/API_DescribeClusters.html
// [kmsuser crypto user]: https://docs.aws.amazon.com/kms/latest/developerguide/keystore-cloudhsm.html#concept-kmsuser
// [Troubleshooting an CloudHSM key store]: https://docs.aws.amazon.com/kms/latest/developerguide/fix-keystore.html
// [CreateHsm]: https://docs.aws.amazon.com/cloudhsm/latest/APIReference/API_CreateHsm.html
// [kms:ConnectCustomKeyStore]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [Troubleshooting an external key store]: https://docs.aws.amazon.com/kms/latest/developerguide/xks-troubleshooting.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [custom key store]: https://docs.aws.amazon.com/kms/latest/developerguide/key-store-overview.html
