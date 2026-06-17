package kms

// DeriveSharedSecret is generated as a reference stub.
// Executable command wiring lives under cmd/kms.go.
//
// Derives a shared secret using a key agreement algorithm.
//
// You must use an asymmetric NIST-standard elliptic curve (ECC) or SM2 (China
// Regions only) KMS key pair with a KeyUsage value of KEY_AGREEMENT to call
// DeriveSharedSecret.
//
// DeriveSharedSecret uses the [Elliptic Curve Cryptography Cofactor Diffie-Hellman Primitive] (ECDH) to establish a key agreement between two
// peers by deriving a shared secret from their elliptic curve public-private key
// pairs. You can use the raw shared secret that DeriveSharedSecret returns to
// derive a symmetric key that can encrypt and decrypt data that is sent between
// the two peers, or that can generate and verify HMACs. KMS recommends that you
// follow [NIST recommendations for key derivation]when using the raw shared secret to derive a symmetric key.
//
// The following workflow demonstrates how to establish key agreement over an
// insecure communication channel using DeriveSharedSecret.
//
// - Alice calls CreateKeyto create an asymmetric KMS key pair with a KeyUsage value of
// KEY_AGREEMENT .
//
// The asymmetric KMS key must use a NIST-standard elliptic curve (ECC) or SM2
//
// (China Regions only) key spec.
//
// - Bob creates an elliptic curve key pair.
//
// Bob can call CreateKeyto create an asymmetric KMS key pair or generate a key pair
//
// outside of KMS. Bob's key pair must use the same NIST-standard elliptic curve
// (ECC) or SM2 (China Regions ony) curve as Alice.
//
// - Alice and Bob exchange their public keys through an insecure communication
// channel (like the internet).
//
// Use GetPublicKeyto download the public key of your asymmetric KMS key pair.
//
// KMS strongly recommends verifying that the public key you receive came from the
//
// expected party before using it to derive a shared secret.
//
// - Alice calls DeriveSharedSecret.
//
// KMS uses the private key from the KMS key pair generated in Step 1, Bob's
//
// public key, and the Elliptic Curve Cryptography Cofactor Diffie-Hellman
// Primitive to derive the shared secret. The private key in your KMS key pair
// never leaves KMS unencrypted. DeriveSharedSecret returns the raw shared secret.
//
// - Bob uses the Elliptic Curve Cryptography Cofactor Diffie-Hellman Primitive
// to calculate the same raw secret using his private key and Alice's public key.
//
// To derive a shared secret you must provide a key agreement algorithm, the
// private key of the caller's asymmetric NIST-standard elliptic curve or SM2
// (China Regions only) KMS key pair, and the public key from your peer's
// NIST-standard elliptic curve or SM2 (China Regions only) key pair. The public
// key can be from another asymmetric KMS key pair or from a key pair generated
// outside of KMS, but both key pairs must be on the same elliptic curve.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: Yes. To perform this operation with a KMS key in a different
// Amazon Web Services account, specify the key ARN or alias ARN in the value of
// the KeyId parameter.
//
// Required permissions: [kms:DeriveSharedSecret] (key policy)
//
// Related operations:
//
// # CreateKey
//
// # GetPublicKey
//
// # DescribeKey
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [kms:DeriveSharedSecret]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [Elliptic Curve Cryptography Cofactor Diffie-Hellman Primitive]: https://nvlpubs.nist.gov/nistpubs/SpecialPublications/NIST.SP.800-56Ar3.pdf#page=60
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [NIST recommendations for key derivation]: https://nvlpubs.nist.gov/nistpubs/SpecialPublications/NIST.SP.800-56Cr2.pdf
