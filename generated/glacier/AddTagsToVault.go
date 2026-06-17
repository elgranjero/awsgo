package glacier

// AddTagsToVault is generated as a reference stub.
// Executable command wiring lives under cmd/glacier.go.
//
// This operation adds the specified tags to a vault. Each tag is composed of a
// key and a value. Each vault can have up to 10 tags. If your request would cause
// the tag limit for the vault to be exceeded, the operation throws the
// LimitExceededException error. If a tag already exists on the vault under a
// specified key, the existing key value will be overwritten. For more information
// about tags, see [Tagging Amazon Glacier Resources].
//
// [Tagging Amazon Glacier Resources]: https://docs.aws.amazon.com/amazonglacier/latest/dev/tagging.html
