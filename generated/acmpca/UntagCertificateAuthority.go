package acmpca

// UntagCertificateAuthority is generated as a reference stub.
// Executable command wiring lives under cmd/acmpca.go.
//
// Remove one or more tags from your private CA. A tag consists of a key-value
// pair. If you do not specify the value portion of the tag when calling this
// action, the tag will be removed regardless of value. If you specify a value, the
// tag is removed only if it is associated with the specified value. To add tags to
// a private CA, use the [TagCertificateAuthority]. Call the [ListTags] action to see what tags are associated with
// your CA.
//
// [TagCertificateAuthority]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_TagCertificateAuthority.html
// [ListTags]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_ListTags.html
