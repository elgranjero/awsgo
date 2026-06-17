package paymentcryptographydata

// TranslateKeyMaterial is generated as a reference stub.
// Executable command wiring lives under cmd/paymentcryptographydata.go.
//
// Translates an cryptographic key between different wrapping keys without
// importing the key into Amazon Web Services Payment Cryptography.
//
// This operation can be used when key material is frequently rotated, such as
// during every card transaction, and there is a need to avoid importing
// short-lived keys into Amazon Web Services Payment Cryptography. It translates
// short-lived transaction keys such as [PEK]generated for each transaction and wrapped
// with an [ECDH]derived wrapping key to another [KEK] wrapping key.
//
// Before using this operation, you must first request the public key certificate
// of the ECC key pair generated within Amazon Web Services Payment Cryptography to
// establish an ECDH key agreement. In TranslateKeyData , the service uses its own
// ECC key pair, public certificate of receiving ECC key pair, and the key
// derivation parameters to generate a derived key. The service uses this derived
// key to unwrap the incoming transaction key received as a TR31WrappedKeyBlock and
// re-wrap using a user provided KEK to generate an outgoing Tr31WrappedKeyBlock.
//
// For information about valid keys for this operation, see [Understanding key attributes] and [Key types for specific data operations] in the Amazon
// Web Services Payment Cryptography User Guide.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// [CreateKey]
//
// [GetPublicCertificate]
//
// [ImportKey]
//
// [KEK]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/terminology.html#terms.kek
// [ECDH]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/terminology.html#terms.ecdh
// [GetPublicCertificate]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_GetPublicKeyCertificate.html
// [PEK]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/terminology.html#terms.pek
// [Key types for specific data operations]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/crypto-ops-validkeys-ops.html
// [ImportKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_ImportKey.html
// [Understanding key attributes]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-validattributes.html
// [CreateKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_CreateKey.html
