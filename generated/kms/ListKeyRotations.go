package kms

// ListKeyRotations is generated as a reference stub.
// Executable command wiring lives under cmd/kms.go.
//
// Returns information about the key materials associated with the specified KMS
// key. You can use the optional IncludeKeyMaterial parameter to control which key
// materials are included in the response.
//
// You must specify the KMS key in all requests. You can refine the key rotations
// list by limiting the number of rotations returned.
//
// For detailed information about automatic and on-demand key rotations, see [Rotate KMS keys] in
// the Key Management Service Developer Guide.
//
// Cross-account use: No. You cannot perform this operation on a KMS key in a
// different Amazon Web Services account.
//
// Required permissions: [kms:ListKeyRotations] (key policy)
//
// Related operations:
//
// # EnableKeyRotation
//
// # DeleteImportedKeyMaterial
//
// # DisableKeyRotation
//
// # GetKeyRotationStatus
//
// # ImportKeyMaterial
//
// # RotateKeyOnDemand
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Rotate KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html
// [kms:ListKeyRotations]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
