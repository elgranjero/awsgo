package paymentcryptographydata

// DecryptData is generated as a reference stub.
// Executable command wiring lives under cmd/paymentcryptographydata.go.
//
// Decrypts ciphertext data to plaintext using a symmetric (TDES, AES), asymmetric
// (RSA), or derived (DUKPT or EMV) encryption key scheme. For more information,
// see [Decrypt data]in the Amazon Web Services Payment Cryptography User Guide.
//
// You can use an decryption key generated within Amazon Web Services Payment
// Cryptography, or you can import your own decryption key by calling [ImportKey]. For this
// operation, the key must have KeyModesOfUse set to Decrypt . In asymmetric
// decryption, Amazon Web Services Payment Cryptography decrypts the ciphertext
// using the private component of the asymmetric encryption key pair. For data
// encryption outside of Amazon Web Services Payment Cryptography, you can export
// the public component of the asymmetric key pair by calling [GetPublicCertificate].
//
// This operation also supports dynamic keys, allowing you to pass a dynamic
// decryption key as a TR-31 WrappedKeyBlock. This can be used when key material is
// frequently rotated, such as during every card transaction, and there is need to
// avoid importing short-lived keys into Amazon Web Services Payment Cryptography.
// To decrypt using dynamic keys, the keyARN is the Key Encryption Key (KEK) of
// the TR-31 wrapped decryption key material. The incoming wrapped key shall have a
// key purpose of D0 with a mode of use of B or D. For more information, see [Using Dynamic Keys]in
// the Amazon Web Services Payment Cryptography User Guide.
//
// For symmetric and DUKPT decryption, Amazon Web Services Payment Cryptography
// supports TDES and AES algorithms. For EMV decryption, Amazon Web Services
// Payment Cryptography supports TDES algorithms. For asymmetric decryption,
// Amazon Web Services Payment Cryptography supports RSA .
//
// When you use TDES or TDES DUKPT, the ciphertext data length must be a multiple
// of 8 bytes. For AES or AES DUKPT, the ciphertext data length must be a multiple
// of 16 bytes. For RSA, it sould be equal to the key size unless padding is
// enabled.
//
// For information about valid keys for this operation, see [Understanding key attributes] and [Key types for specific data operations] in the Amazon
// Web Services Payment Cryptography User Guide.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// # EncryptData
//
// [GetPublicCertificate]
//
// [ImportKey]
//
// [Using Dynamic Keys]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/use-cases-acquirers-dynamickeys.html
// [GetPublicCertificate]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_GetPublicKeyCertificate.html
// [ImportKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_ImportKey.html
// [Key types for specific data operations]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/crypto-ops-validkeys-ops.html
// [Decrypt data]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/decrypt-data.html
// [Understanding key attributes]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-validattributes.html
