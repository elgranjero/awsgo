package kms

// ListResourceTags is generated as a reference stub.
// Executable command wiring lives under cmd/kms.go.
//
// Returns all tags on the specified KMS key.
//
// For general information about tags, including the format and syntax, see [Tagging Amazon Web Services resources] in
// the Amazon Web Services General Reference. For information about using tags in
// KMS, see [Tags in KMS].
//
// Cross-account use: No. You cannot perform this operation on a KMS key in a
// different Amazon Web Services account.
//
// Required permissions: [kms:ListResourceTags] (key policy)
//
// Related operations:
//
// # CreateKey
//
// # ReplicateKey
//
// # TagResource
//
// # UntagResource
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [kms:ListResourceTags]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [Tags in KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/tagging-keys.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [Tagging Amazon Web Services resources]: https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html
