package kms

// CreateCustomKeyStore is generated as a reference stub.
// Executable command wiring lives under cmd/kms.go.
//
// Creates a [custom key store] backed by a key store that you own and manage. When you use a KMS
// key in a custom key store for a cryptographic operation, the cryptographic
// operation is actually performed in your key store using your keys. KMS supports [CloudHSM key stores]
// backed by an [CloudHSM cluster]and [external key stores] backed by an external key store proxy and external key
// manager outside of Amazon Web Services.
//
// This operation is part of the custom key stores feature in KMS, which combines
// the convenience and extensive integration of KMS with the isolation and control
// of a key store that you own and manage.
//
// Before you create the custom key store, the required elements must be in place
// and operational. We recommend that you use the test tools that KMS provides to
// verify the configuration your external key store proxy. For details about the
// required elements and verification tests, see [Assemble the prerequisites (for CloudHSM key stores)]or [Assemble the prerequisites (for external key stores)] in the Key Management Service
// Developer Guide.
//
// To create a custom key store, use the following parameters.
//
// - To create an CloudHSM key store, specify the CustomKeyStoreName ,
// CloudHsmClusterId , KeyStorePassword , and TrustAnchorCertificate . The
// CustomKeyStoreType parameter is optional for CloudHSM key stores. If you
// include it, set it to the default value, AWS_CLOUDHSM . For help with
// failures, see [Troubleshooting an CloudHSM key store]in the Key Management Service Developer Guide.
//
// - To create an external key store, specify the CustomKeyStoreName and a
// CustomKeyStoreType of EXTERNAL_KEY_STORE . Also, specify values for
// XksProxyConnectivity , XksProxyAuthenticationCredential , XksProxyUriEndpoint
// , and XksProxyUriPath . If your XksProxyConnectivity value is
// VPC_ENDPOINT_SERVICE , specify the XksProxyVpcEndpointServiceName parameter.
// For help with failures, see [Troubleshooting an external key store]in the Key Management Service Developer Guide.
//
// For external key stores:
//
// Some external key managers provide a simpler method for creating an external
// key store. For details, see your external key manager documentation.
//
// When creating an external key store in the KMS console, you can upload a
// JSON-based proxy configuration file with the desired values. You cannot use a
// proxy configuration with the CreateCustomKeyStore operation. However, you can
// use the values in the file to help you determine the correct values for the
// CreateCustomKeyStore parameters.
//
// When the operation completes successfully, it returns the ID of the new custom
// key store. Before you can use your new custom key store, you need to use the ConnectCustomKeyStore
// operation to connect a new CloudHSM key store to its CloudHSM cluster, or to
// connect a new external key store to the external key store proxy for your
// external key manager. Even if you are not going to use your custom key store
// immediately, you might want to connect it to verify that all settings are
// correct and then disconnect it until you are ready to use it.
//
// Cross-account use: No. You cannot perform this operation on a custom key store
// in a different Amazon Web Services account.
//
// Required permissions: [kms:CreateCustomKeyStore] (IAM policy).
//
// Related operations:
//
// # ConnectCustomKeyStore
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
// [CloudHSM key stores]: https://docs.aws.amazon.com/kms/latest/developerguide/keystore-cloudhsm.html
// [CloudHSM cluster]: https://docs.aws.amazon.com/cloudhsm/latest/userguide/clusters.html
// [external key stores]: https://docs.aws.amazon.com/kms/latest/developerguide/keystore-external.html
// [Troubleshooting an CloudHSM key store]: https://docs.aws.amazon.com/kms/latest/developerguide/fix-keystore.html
// [Assemble the prerequisites (for CloudHSM key stores)]: https://docs.aws.amazon.com/kms/latest/developerguide/create-keystore.html#before-keystore
// [Assemble the prerequisites (for external key stores)]: https://docs.aws.amazon.com/kms/latest/developerguide/create-xks-keystore.html#xks-requirements
// [Troubleshooting an external key store]: https://docs.aws.amazon.com/kms/latest/developerguide/xks-troubleshooting.html
// [kms:CreateCustomKeyStore]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [custom key store]: https://docs.aws.amazon.com/kms/latest/developerguide/key-store-overview.html
