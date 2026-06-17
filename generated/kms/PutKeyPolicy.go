package kms

// PutKeyPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/kms.go.
//
// Attaches a key policy to the specified KMS key.
//
// For more information about key policies, see [Key Policies] in the Key Management Service
// Developer Guide. For help writing and formatting a JSON policy document, see the
// [IAM JSON Policy Reference]in the Identity and Access Management User Guide . For examples of adding a key
// policy in multiple programming languages, see [Use PutKeyPolicy with an Amazon Web Services SDK or CLI]in the Key Management Service
// Developer Guide.
//
// Cross-account use: No. You cannot perform this operation on a KMS key in a
// different Amazon Web Services account.
//
// Required permissions: [kms:PutKeyPolicy] (key policy)
//
// Related operations: GetKeyPolicy
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [IAM JSON Policy Reference]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies.html
// [kms:PutKeyPolicy]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [Use PutKeyPolicy with an Amazon Web Services SDK or CLI]: https://docs.aws.amazon.com/kms/latest/developerguide/example_kms_PutKeyPolicy_section.html
// [Key Policies]: https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
