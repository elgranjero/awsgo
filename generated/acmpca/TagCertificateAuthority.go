package acmpca

// TagCertificateAuthority is generated as a reference stub.
// Executable command wiring lives under cmd/acmpca.go.
//
// Adds one or more tags to your private CA. Tags are labels that you can use to
// identify and organize your Amazon Web Services resources. Each tag consists of a
// key and an optional value. You specify the private CA on input by its Amazon
// Resource Name (ARN). You specify the tag by using a key-value pair. You can
// apply a tag to just one private CA if you want to identify a specific
// characteristic of that CA, or you can apply the same tag to multiple private CAs
// if you want to filter for a common relationship among those CAs. To remove one
// or more tags, use the [UntagCertificateAuthority]action. Call the [ListTags] action to see what tags are associated
// with your CA.
//
// To attach tags to a private CA during the creation procedure, a CA
// administrator must first associate an inline IAM policy with the
// CreateCertificateAuthority action and explicitly allow tagging. For more
// information, see [Attaching tags to a CA at the time of creation].
//
// [Attaching tags to a CA at the time of creation]: https://docs.aws.amazon.com/privateca/latest/userguide/auth-InlinePolicies.html#policy-tag-ca
// [UntagCertificateAuthority]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_UntagCertificateAuthority.html
// [ListTags]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_ListTags.html
