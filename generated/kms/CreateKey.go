package kms

// CreateKey is generated as a reference stub.
// Executable command wiring lives under cmd/kms.go.
//
// Creates a unique customer managed [KMS key] in your Amazon Web Services account and
// Region. You can use a KMS key in cryptographic operations, such as encryption
// and signing. Some Amazon Web Services services let you use KMS keys that you
// create and manage to protect your service resources.
//
// A KMS key is a logical representation of a cryptographic key. In addition to
// the key material used in cryptographic operations, a KMS key includes metadata,
// such as the key ID, key policy, creation date, description, and key state.
//
// Use the parameters of CreateKey to specify the type of KMS key, the source of
// its key material, its key policy, description, tags, and other properties.
//
// KMS has replaced the term customer master key (CMK) with Key Management Service
// key and KMS key. The concept has not changed. To prevent breaking changes, KMS
// is keeping some variations of this term.
//
// To create different types of KMS keys, use the following guidance:
//
// Symmetric encryption KMS key By default, CreateKey creates a symmetric
// encryption KMS key with key material that KMS generates. This is the basic and
// most widely used type of KMS key, and provides the best performance.
//
// To create a symmetric encryption KMS key, you don't need to specify any
// parameters. The default value for KeySpec , SYMMETRIC_DEFAULT , the default
// value for KeyUsage , ENCRYPT_DECRYPT , and the default value for Origin ,
// AWS_KMS , create a symmetric encryption KMS key with KMS key material.
//
// If you need a key for basic encryption and decryption or you are creating a KMS
// key to protect your resources in an Amazon Web Services service, create a
// symmetric encryption KMS key. The key material in a symmetric encryption key
// never leaves KMS unencrypted. You can use a symmetric encryption KMS key to
// encrypt and decrypt data up to 4,096 bytes, but they are typically used to
// generate data keys and data keys pairs. For details, see GenerateDataKeyand GenerateDataKeyPair.
//
// Asymmetric KMS keys To create an asymmetric KMS key, use the KeySpec parameter
// to specify the type of key material in the KMS key. Then, use the KeyUsage
// parameter to determine whether the KMS key will be used to encrypt and decrypt
// or sign and verify. You can't change these properties after the KMS key is
// created.
//
// Asymmetric KMS keys contain an RSA key pair, Elliptic Curve (ECC) key pair,
// ML-DSA key pair or an SM2 key pair (China Regions only). The private key in an
// asymmetric KMS key never leaves KMS unencrypted. However, you can use the GetPublicKey
// operation to download the public key so it can be used outside of KMS. Each KMS
// key can have only one key usage. KMS keys with RSA key pairs can be used to
// encrypt and decrypt data or sign and verify messages (but not both). KMS keys
// with NIST-standard ECC key pairs can be used to sign and verify messages or
// derive shared secrets (but not both). KMS keys with ECC_SECG_P256K1 can be used
// only to sign and verify messages. KMS keys with ML-DSA key pairs can be used to
// sign and verify messages. KMS keys with SM2 key pairs (China Regions only) can
// be used to either encrypt and decrypt data, sign and verify messages, or derive
// shared secrets (you must choose one key usage type). For information about
// asymmetric KMS keys, see [Asymmetric KMS keys]in the Key Management Service Developer Guide.
//
// HMAC KMS key To create an HMAC KMS key, set the KeySpec parameter to a key spec
// value for HMAC KMS keys. Then set the KeyUsage parameter to GENERATE_VERIFY_MAC
// . You must set the key usage even though GENERATE_VERIFY_MAC is the only valid
// key usage value for HMAC KMS keys. You can't change these properties after the
// KMS key is created.
//
// HMAC KMS keys are symmetric keys that never leave KMS unencrypted. You can use
// HMAC keys to generate (GenerateMac ) and verify (VerifyMac ) HMAC codes for messages up to 4096
// bytes.
//
// Multi-Region primary keys To create a multi-Region primary key in the local
// Amazon Web Services Region, use the MultiRegion parameter with a value of True .
// To create a multi-Region replica key, that is, a KMS key with the same key ID
// and key material as a primary key, but in a different Amazon Web Services
// Region, use the ReplicateKeyoperation. To change a replica key to a primary key, and its
// primary key to a replica key, use the UpdatePrimaryRegionoperation.
//
// You can create multi-Region KMS keys for all supported KMS key types: symmetric
// encryption KMS keys, HMAC KMS keys, asymmetric encryption KMS keys, and
// asymmetric signing KMS keys. You can also create multi-Region keys with imported
// key material. However, you can't create multi-Region keys in a custom key store.
//
// This operation supports multi-Region keys, an KMS feature that lets you create
// multiple interoperable KMS keys in different Amazon Web Services Regions.
// Because these KMS keys have the same key ID, key material, and other metadata,
// you can use them interchangeably to encrypt data in one Amazon Web Services
// Region and decrypt it in a different Amazon Web Services Region without
// re-encrypting the data or making a cross-Region call. For more information about
// multi-Region keys, see [Multi-Region keys in KMS]in the Key Management Service Developer Guide.
//
// Imported key material To import your own key material into a KMS key, begin by
// creating a KMS key with no key material. To do this, use the Origin parameter
// of CreateKey with a value of EXTERNAL . Next, use GetParametersForImport operation to get a public
// key and import token. Use the wrapping public key to encrypt your key material.
// Then, use ImportKeyMaterialwith your import token to import the key material. For step-by-step
// instructions, see [Importing Key Material]in the Key Management Service Developer Guide .
//
// You can import key material into KMS keys of all supported KMS key types:
// symmetric encryption KMS keys, HMAC KMS keys, asymmetric encryption KMS keys,
// and asymmetric signing KMS keys. You can also create multi-Region keys with
// imported key material. However, you can't import key material into a KMS key in
// a custom key store.
//
// To create a multi-Region primary key with imported key material, use the Origin
// parameter of CreateKey with a value of EXTERNAL and the MultiRegion parameter
// with a value of True . To create replicas of the multi-Region primary key, use
// the ReplicateKeyoperation. For instructions, see [Importing key material step 1]. For more information about multi-Region
// keys, see [Multi-Region keys in KMS]in the Key Management Service Developer Guide.
//
// Custom key store A [custom key store] lets you protect your Amazon Web Services resources using
// keys in a backing key store that you own and manage. When you request a
// cryptographic operation with a KMS key in a custom key store, the operation is
// performed in the backing key store using its cryptographic keys.
//
// KMS supports [CloudHSM key stores] backed by an CloudHSM cluster and [external key stores] backed by an external key
// manager outside of Amazon Web Services. When you create a KMS key in an CloudHSM
// key store, KMS generates an encryption key in the CloudHSM cluster and
// associates it with the KMS key. When you create a KMS key in an external key
// store, you specify an existing encryption key in the external key manager.
//
// Some external key managers provide a simpler method for creating a KMS key in
// an external key store. For details, see your external key manager documentation.
//
// Before you create a KMS key in a custom key store, the ConnectionState of the
// key store must be CONNECTED . To connect the custom key store, use the ConnectCustomKeyStore
// operation. To find the ConnectionState , use the DescribeCustomKeyStores operation.
//
// To create a KMS key in a custom key store, use the CustomKeyStoreId . Use the
// default KeySpec value, SYMMETRIC_DEFAULT , and the default KeyUsage value,
// ENCRYPT_DECRYPT to create a symmetric encryption key. No other key type is
// supported in a custom key store.
//
// To create a KMS key in an [CloudHSM key store], use the Origin parameter with a value of
// AWS_CLOUDHSM . The CloudHSM cluster that is associated with the custom key store
// must have at least two active HSMs in different Availability Zones in the Amazon
// Web Services Region.
//
// To create a KMS key in an [external key store], use the Origin parameter with a value of
// EXTERNAL_KEY_STORE and an XksKeyId parameter that identifies an existing
// external key.
//
// Some external key managers provide a simpler method for creating a KMS key in
// an external key store. For details, see your external key manager documentation.
//
// Cross-account use: No. You cannot use this operation to create a KMS key in a
// different Amazon Web Services account.
//
// Required permissions: [kms:CreateKey] (IAM policy). To use the Tags parameter, [kms:TagResource] (IAM policy).
// For examples and information about related permissions, see [Allow a user to create KMS keys]in the Key
// Management Service Developer Guide.
//
// Related operations:
//
// # DescribeKey
//
// # ListKeys
//
// # ScheduleKeyDeletion
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [CloudHSM key stores]: https://docs.aws.amazon.com/kms/latest/developerguide/keystore-cloudhsm.html
// [external key store]: https://docs.aws.amazon.com/kms/latest/developerguide/create-xks-keys.html
// [external key stores]: https://docs.aws.amazon.com/kms/latest/developerguide/keystore-external.html
// [Asymmetric KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html
// [Multi-Region keys in KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html
// [Importing key material step 1]: https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys-create-cmk.html
// [KMS key]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#kms-keys
// [Allow a user to create KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/customer-managed-policies.html#iam-policy-example-create-key
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [kms:TagResource]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [CloudHSM key store]: https://docs.aws.amazon.com/kms/latest/developerguide/create-cmk-keystore.html
// [kms:CreateKey]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [Importing Key Material]: https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html
// [custom key store]: https://docs.aws.amazon.com/kms/latest/developerguide/key-store-overview.html
