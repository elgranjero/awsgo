package kms

// UpdateCustomKeyStore is generated as a reference stub.
// Executable command wiring lives under cmd/kms.go.
//
// Changes the properties of a custom key store. You can use this operation to
// change the properties of an CloudHSM key store or an external key store.
//
// Use the required CustomKeyStoreId parameter to identify the custom key store.
// Use the remaining optional parameters to change its properties. This operation
// does not return any property values. To verify the updated property values, use
// the DescribeCustomKeyStoresoperation.
//
// This operation is part of the custom key stores feature in KMS, which combines
// the convenience and extensive integration of KMS with the isolation and control
// of a key store that you own and manage.
//
// When updating the properties of an external key store, verify that the updated
// settings connect your key store, via the external key store proxy, to the same
// external key manager as the previous settings, or to a backup or snapshot of the
// external key manager with the same cryptographic keys. If the updated connection
// settings fail, you can fix them and retry, although an extended delay might
// disrupt Amazon Web Services services. However, if KMS permanently loses its
// access to cryptographic keys, ciphertext encrypted under those keys is
// unrecoverable.
//
// For external key stores:
//
// Some external key managers provide a simpler method for updating an external
// key store. For details, see your external key manager documentation.
//
// When updating an external key store in the KMS console, you can upload a
// JSON-based proxy configuration file with the desired values. You cannot upload
// the proxy configuration file to the UpdateCustomKeyStore operation. However,
// you can use the file to help you determine the correct values for the
// UpdateCustomKeyStore parameters.
//
// For an CloudHSM key store, you can use this operation to change the custom key
// store friendly name ( NewCustomKeyStoreName ), to tell KMS about a change to the
// kmsuser crypto user password ( KeyStorePassword ), or to associate the custom
// key store with a different, but related, CloudHSM cluster ( CloudHsmClusterId ).
// To update any property of an CloudHSM key store, the ConnectionState of the
// CloudHSM key store must be DISCONNECTED .
//
// For an external key store, you can use this operation to change the custom key
// store friendly name ( NewCustomKeyStoreName ), or to tell KMS about a change to
// the external key store proxy authentication credentials (
// XksProxyAuthenticationCredential ), connection method ( XksProxyConnectivity ),
// external proxy endpoint ( XksProxyUriEndpoint ) and path ( XksProxyUriPath ).
// For external key stores with an XksProxyConnectivity of VPC_ENDPOINT_SERVICE ,
// you can also update the Amazon VPC endpoint service name (
// XksProxyVpcEndpointServiceName ). To update most properties of an external key
// store, the ConnectionState of the external key store must be DISCONNECTED .
// However, you can update the CustomKeyStoreName ,
// XksProxyAuthenticationCredential , and XksProxyUriPath of an external key store
// when it is in the CONNECTED or DISCONNECTED state.
//
// If your update requires a DISCONNECTED state, before using UpdateCustomKeyStore
// , use the DisconnectCustomKeyStoreoperation to disconnect the custom key store. After the
// UpdateCustomKeyStore operation completes, use the ConnectCustomKeyStore to reconnect the custom key
// store. To find the ConnectionState of the custom key store, use the DescribeCustomKeyStores operation.
//
// Before updating the custom key store, verify that the new values allow KMS to
// connect the custom key store to its backing key store. For example, before you
// change the XksProxyUriPath value, verify that the external key store proxy is
// reachable at the new path.
//
// If the operation succeeds, it returns a JSON object with no properties.
//
// Cross-account use: No. You cannot perform this operation on a custom key store
// in a different Amazon Web Services account.
//
// Required permissions: [kms:UpdateCustomKeyStore] (IAM policy)
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
// # DisconnectCustomKeyStore
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [kms:UpdateCustomKeyStore]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
