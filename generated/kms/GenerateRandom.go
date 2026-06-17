package kms

// GenerateRandom is generated as a reference stub.
// Executable command wiring lives under cmd/kms.go.
//
// Returns a random byte string that is cryptographically secure.
//
// You must use the NumberOfBytes parameter to specify the length of the random
// byte string. There is no default value for string length.
//
// By default, the random byte string is generated in KMS. To generate the byte
// string in the CloudHSM cluster associated with an CloudHSM key store, use the
// CustomKeyStoreId parameter.
//
// GenerateRandom also supports [Amazon Web Services Nitro Enclaves], which provide an isolated compute environment in
// Amazon EC2. To call GenerateRandom for a Nitro enclave or NitroTPM, use the [Amazon Web Services Nitro Enclaves SDK] or
// any Amazon Web Services SDK. Use the Recipient parameter to provide the
// attestation document for the attested environment. Instead of plaintext bytes,
// the response includes the plaintext bytes encrypted under the public key from
// the attestation document ( CiphertextForRecipient ). For information about the
// interaction between KMS and Amazon Web Services Nitro Enclaves or Amazon Web
// Services NitroTPM, see [Cryptographic attestation support in KMS]in the Key Management Service Developer Guide.
//
// For more information about entropy and random number generation, see [Entropy and random number generation] in the
// Key Management Service Developer Guide.
//
// Cross-account use: Not applicable. GenerateRandom does not use any
// account-specific resources, such as KMS keys.
//
// Required permissions: [kms:GenerateRandom] (IAM policy)
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Cryptographic attestation support in KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/cryptographic-attestation.html
// [kms:GenerateRandom]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [Amazon Web Services Nitro Enclaves]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/nitro-enclave.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [Amazon Web Services Nitro Enclaves SDK]: https://docs.aws.amazon.com/enclaves/latest/user/developing-applications.html#sdk
// [Entropy and random number generation]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-cryptography.html#entropy-and-random-numbers
