package paymentcryptographydata

// EncryptData is generated as a reference stub.
// Executable command wiring lives under cmd/paymentcryptographydata.go.
//
// Encrypts plaintext data to ciphertext using a symmetric (TDES, AES), asymmetric
// (RSA), or derived (DUKPT or EMV) encryption key scheme. For more information,
// see [Encrypt data]in the Amazon Web Services Payment Cryptography User Guide.
//
// You can generate an encryption key within Amazon Web Services Payment
// Cryptography by calling [CreateKey]. You can import your own encryption key by calling [ImportKey].
//
// For this operation, the key must have KeyModesOfUse set to Encrypt . In
// asymmetric encryption, plaintext is encrypted using public component. You can
// import the public component of an asymmetric key pair created outside Amazon Web
// Services Payment Cryptography by calling [ImportKey].
//
// This operation also supports dynamic keys, allowing you to pass a dynamic
// encryption key as a TR-31 WrappedKeyBlock. This can be used when key material is
// frequently rotated, such as during every card transaction, and there is need to
// avoid importing short-lived keys into Amazon Web Services Payment Cryptography.
// To encrypt using dynamic keys, the keyARN is the Key Encryption Key (KEK) of
// the TR-31 wrapped encryption key material. The incoming wrapped key shall have a
// key purpose of D0 with a mode of use of B or D. For more information, see [Using Dynamic Keys]in
// the Amazon Web Services Payment Cryptography User Guide.
//
// For symmetric and DUKPT encryption, Amazon Web Services Payment Cryptography
// supports TDES and AES algorithms. For EMV encryption, Amazon Web Services
// Payment Cryptography supports TDES algorithms.For asymmetric encryption, Amazon
// Web Services Payment Cryptography supports RSA .
//
// When you use TDES or TDES DUKPT, the plaintext data length must be a multiple
// of 8 bytes. For AES or AES DUKPT, the plaintext data length must be a multiple
// of 16 bytes. For RSA, it sould be equal to the key size unless padding is
// enabled.
//
// To encrypt using DUKPT, you must already have a BDK (Base Derivation Key) key
// in your account with KeyModesOfUse set to DeriveKey , or you can generate a new
// DUKPT key by calling [CreateKey]. To encrypt using EMV, you must already have an IMK
// (Issuer Master Key) key in your account with KeyModesOfUse set to DeriveKey .
//
// For information about valid keys for this operation, see [Understanding key attributes] and [Key types for specific data operations] in the Amazon
// Web Services Payment Cryptography User Guide.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// # DecryptData
//
// [GetPublicCertificate]
//
// [ImportKey]
//
// # ReEncryptData
//
// [Using Dynamic Keys]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/use-cases-acquirers-dynamickeys.html
// [GetPublicCertificate]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_GetPublicKeyCertificate.html
// [Encrypt data]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/encrypt-data.html
// [ImportKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_ImportKey.html
// [Key types for specific data operations]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/crypto-ops-validkeys-ops.html
// [Understanding key attributes]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-validattributes.html
// [CreateKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_CreateKey.html
