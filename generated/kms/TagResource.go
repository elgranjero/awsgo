package kms

// TagResource is generated as a reference stub.
// Executable command wiring lives under cmd/kms.go.
//
// Adds or edits tags on a [customer managed key].
//
// Tagging or untagging a KMS key can allow or deny permission to the KMS key. For
// details, see [ABAC for KMS]in the Key Management Service Developer Guide.
//
// Each tag consists of a tag key and a tag value, both of which are
// case-sensitive strings. The tag value can be an empty (null) string. To add a
// tag, specify a new tag key and a tag value. To edit a tag, specify an existing
// tag key and a new tag value.
//
// You can use this operation to tag a [customer managed key], but you cannot tag an [Amazon Web Services managed key], an [Amazon Web Services owned key], a [custom key store], or an [alias].
//
// You can also add tags to a KMS key while creating it (CreateKey ) or replicating it (ReplicateKey ).
//
// For information about using tags in KMS, see [Tagging keys]. For general information about
// tags, including the format and syntax, see [Tagging Amazon Web Services resources]in the Amazon Web Services General
// Reference.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: No. You cannot perform this operation on a KMS key in a
// different Amazon Web Services account.
//
// Required permissions: [kms:TagResource] (key policy)
//
// # Related operations
//
// # CreateKey
//
// # ListResourceTags
//
// # ReplicateKey
//
// # UntagResource
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Amazon Web Services owned key]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-owned-key
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [kms:TagResource]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [customer managed key]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#customer-mgn-key
// [Tagging keys]: https://docs.aws.amazon.com/kms/latest/developerguide/tagging-keys.html
// [alias]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-alias.html
// [ABAC for KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/abac.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [Amazon Web Services managed key]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-key
// [custom key store]: https://docs.aws.amazon.com/kms/latest/developerguide/key-store-overview.html
// [Tagging Amazon Web Services resources]: https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html
