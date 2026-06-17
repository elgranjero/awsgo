package paymentcryptographydata

// VerifyPinData is generated as a reference stub.
// Executable command wiring lives under cmd/paymentcryptographydata.go.
//
// Verifies pin-related data such as PIN and PIN Offset using algorithms including
// VISA PVV and IBM3624. For more information, see [Verify PIN data]in the Amazon Web Services
// Payment Cryptography User Guide.
//
// This operation verifies PIN data for user payment card. A card holder PIN data
// is never transmitted in clear to or from Amazon Web Services Payment
// Cryptography. This operation uses PIN Verification Key (PVK) for PIN or PIN
// Offset generation and then encrypts it using PIN Encryption Key (PEK) to create
// an EncryptedPinBlock for transmission from Amazon Web Services Payment
// Cryptography.
//
// For information about valid keys for this operation, see [Understanding key attributes] and [Key types for specific data operations] in the Amazon
// Web Services Payment Cryptography User Guide.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// # GeneratePinData
//
// # TranslatePinData
//
// [Key types for specific data operations]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/crypto-ops-validkeys-ops.html
// [Understanding key attributes]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-validattributes.html
// [Verify PIN data]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/verify-pin-data.html
