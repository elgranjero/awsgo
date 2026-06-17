package paymentcryptography

// CreateKey is generated as a reference stub.
// Executable command wiring lives under cmd/paymentcryptography.go.
//
// Creates an Amazon Web Services Payment Cryptography key, a logical
// representation of a cryptographic key, that is unique in your account and Amazon
// Web Services Region. You use keys for cryptographic functions such as encryption
// and decryption.
//
// In addition to the key material used in cryptographic operations, an Amazon Web
// Services Payment Cryptography key includes metadata such as the key ARN, key
// usage, key origin, creation date, description, and key state.
//
// When you create a key, you specify both immutable and mutable data about the
// key. The immutable data contains key attributes that define the scope and
// cryptographic operations that you can perform using the key, for example key
// class (example: SYMMETRIC_KEY ), key algorithm (example: TDES_2KEY ), key usage
// (example: TR31_P0_PIN_ENCRYPTION_KEY ) and key modes of use (example: Encrypt ).
// Amazon Web Services Payment Cryptography binds key attributes to keys using key
// blocks when you store or export them. Amazon Web Services Payment Cryptography
// stores the key contents wrapped and never stores or transmits them in the clear.
//
// For information about valid combinations of key attributes, see [Understanding key attributes] in the Amazon
// Web Services Payment Cryptography User Guide. The mutable data contained within
// a key includes usage timestamp and key deletion timestamp and can be modified
// after creation.
//
// You can use the CreateKey operation to generate an ECC (Elliptic Curve
// Cryptography) key pair used for establishing an ECDH (Elliptic Curve
// Diffie-Hellman) key agreement between two parties. In the ECDH key agreement
// process, both parties generate their own ECC key pair with key usage K3 and
// exchange the public keys. Each party then use their private key, the received
// public key from the other party, and the key derivation parameters including key
// derivation function, hash algorithm, derivation data, and key algorithm to
// derive a shared key.
//
// To maintain the single-use principle of cryptographic keys in payments, ECDH
// derived keys should not be used for multiple purposes, such as a
// TR31_P0_PIN_ENCRYPTION_KEY and TR31_K1_KEY_BLOCK_PROTECTION_KEY . When creating
// ECC key pairs in Amazon Web Services Payment Cryptography you can optionally set
// the DeriveKeyUsage parameter, which defines the key usage bound to the
// symmetric key that will be derived using the ECC key pair.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// [DeleteKey]
//
// [GetKey]
//
// [ListKeys]
//
// [DeleteKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_DeleteKey.html
// [GetKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_GetKey.html
// [ListKeys]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_ListKeys.html
// [Understanding key attributes]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-validattributes.html
