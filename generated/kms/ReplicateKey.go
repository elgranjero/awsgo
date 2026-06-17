package kms

// ReplicateKey is generated as a reference stub.
// Executable command wiring lives under cmd/kms.go.
//
// Replicates a multi-Region key into the specified Region. This operation creates
// a multi-Region replica key based on a multi-Region primary key in a different
// Region of the same Amazon Web Services partition. You can create multiple
// replicas of a primary key, but each must be in a different Region. To create a
// multi-Region primary key, use the CreateKeyoperation.
//
// This operation supports multi-Region keys, an KMS feature that lets you create
// multiple interoperable KMS keys in different Amazon Web Services Regions.
// Because these KMS keys have the same key ID, key material, and other metadata,
// you can use them interchangeably to encrypt data in one Amazon Web Services
// Region and decrypt it in a different Amazon Web Services Region without
// re-encrypting the data or making a cross-Region call. For more information about
// multi-Region keys, see [Multi-Region keys in KMS]in the Key Management Service Developer Guide.
//
// A replica key is a fully-functional KMS key that can be used independently of
// its primary and peer replica keys. A primary key and its replica keys share
// properties that make them interoperable. They have the same [key ID]and key material.
// They also have the same key spec, key usage, key material origin, and automatic
// key rotation status. KMS automatically synchronizes these shared properties
// among related multi-Region keys. All other properties of a replica key can
// differ, including its [key policy], [tags], [aliases], and [key state]. KMS pricing and quotas for KMS keys apply to
// each primary key and replica key.
//
// When this operation completes, the new replica key has a transient key state of
// Creating . This key state changes to Enabled (or PendingImport ) after a few
// seconds when the process of creating the new replica key is complete. While the
// key state is Creating , you can manage key, but you cannot yet use it in
// cryptographic operations. If you are creating and using the replica key
// programmatically, retry on KMSInvalidStateException or call DescribeKey to
// check its KeyState value before using it. For details about the Creating key
// state, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// You cannot create more than one replica of a primary key in any Region. If the
// Region already includes a replica of the key you're trying to replicate,
// ReplicateKey returns an AlreadyExistsException error. If the key state of the
// existing replica is PendingDeletion , you can cancel the scheduled key deletion (CancelKeyDeletion
// ) or wait for the key to be deleted. The new replica key you create will have
// the same [shared properties]as the original replica key.
//
// The CloudTrail log of a ReplicateKey operation records a ReplicateKey operation
// in the primary key's Region and a CreateKeyoperation in the replica key's Region.
//
// If you replicate a multi-Region primary key with imported key material, the
// replica key is created with no key material. You must import the same key
// material that you imported into the primary key.
//
// To convert a replica key to a primary key, use the UpdatePrimaryRegion operation.
//
// ReplicateKey uses different default values for the KeyPolicy and Tags
// parameters than those used in the KMS console. For details, see the parameter
// descriptions.
//
// Cross-account use: No. You cannot use this operation to create a replica key in
// a different Amazon Web Services account.
//
// Required permissions:
//
// - kms:ReplicateKey on the primary key (in the primary key's Region). Include
// this permission in the primary key's key policy.
//
// - kms:CreateKey in an IAM policy in the replica Region.
//
// - To use the Tags parameter, kms:TagResource in an IAM policy in the replica
// Region.
//
// # Related operations
//
// # CreateKey
//
// # UpdatePrimaryRegion
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [key ID]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id-key-id
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [aliases]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-alias.html
// [Multi-Region keys in KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html
// [key policy]: https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html
// [key state]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [tags]: https://docs.aws.amazon.com/kms/latest/developerguide/tagging-keys.html
// [shared properties]: https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html#mrk-sync-properties
