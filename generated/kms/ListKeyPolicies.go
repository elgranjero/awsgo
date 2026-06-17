package kms

// ListKeyPolicies is generated as a reference stub.
// Executable command wiring lives under cmd/kms.go.
//
// Gets the names of the key policies that are attached to a KMS key. This
// operation is designed to get policy names that you can use in a GetKeyPolicyoperation.
// However, the only valid policy name is default .
//
// Cross-account use: No. You cannot perform this operation on a KMS key in a
// different Amazon Web Services account.
//
// Required permissions: [kms:ListKeyPolicies] (key policy)
//
// Related operations:
//
// # GetKeyPolicy
//
// [PutKeyPolicy]
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [kms:ListKeyPolicies]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [PutKeyPolicy]: https://docs.aws.amazon.com/kms/latest/APIReference/API_PutKeyPolicy.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
