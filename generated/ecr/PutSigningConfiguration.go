package ecr

// PutSigningConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/ecr.go.
//
// Creates or updates the registry's signing configuration, which defines rules
// for automatically signing images with Amazon Web Services Signer.
//
// For more information, see [Managed signing] in the Amazon Elastic Container Registry User Guide.
//
// To successfully generate a signature, the IAM principal pushing images must
// have permission to sign payloads with the Amazon Web Services Signer signing
// profile referenced in the signing configuration.
//
// [Managed signing]: https://docs.aws.amazon.com/AmazonECR/latest/userguide/managed-signing.html
