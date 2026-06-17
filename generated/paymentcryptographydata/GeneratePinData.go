package paymentcryptographydata

// GeneratePinData is generated as a reference stub.
// Executable command wiring lives under cmd/paymentcryptographydata.go.
//
// Generates pin-related data such as PIN, PIN Verification Value (PVV), PIN
// Block, and PIN Offset during new card issuance or reissuance. For more
// information, see [Generate PIN data]in the Amazon Web Services Payment Cryptography User Guide.
//
// PIN data is never transmitted in clear to or from Amazon Web Services Payment
// Cryptography. This operation generates PIN, PVV, or PIN Offset and then encrypts
// it using Pin Encryption Key (PEK) to create an EncryptedPinBlock for
// transmission from Amazon Web Services Payment Cryptography. This operation uses
// a separate Pin Verification Key (PVK) for VISA PVV generation.
//
// Using ECDH key exchange, you can receive cardholder selectable PINs into Amazon
// Web Services Payment Cryptography. The ECDH derived key protects the incoming
// PIN block. You can also use it for reveal PIN, wherein the generated PIN block
// is protected by the ECDH derived key before transmission from Amazon Web
// Services Payment Cryptography. For more information on establishing ECDH derived
// keys, see the [Generating keys]in the Amazon Web Services Payment Cryptography User Guide.
//
// For information about valid keys for this operation, see [Understanding key attributes] and [Key types for specific data operations] in the Amazon
// Web Services Payment Cryptography User Guide.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// # GenerateCardValidationData
//
// # TranslatePinData
//
// # VerifyPinData
//
// [Generate PIN data]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/generate-pin-data.html
// [Key types for specific data operations]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/crypto-ops-validkeys-ops.html
// [Understanding key attributes]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-validattributes.html
// [Generating keys]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/create-keys.html
