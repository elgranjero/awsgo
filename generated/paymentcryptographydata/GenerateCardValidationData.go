package paymentcryptographydata

// GenerateCardValidationData is generated as a reference stub.
// Executable command wiring lives under cmd/paymentcryptographydata.go.
//
// Generates card-related validation data using algorithms such as Card
// Verification Values (CVV/CVV2), Dynamic Card Verification Values (dCVV/dCVV2),
// or Card Security Codes (CSC). For more information, see [Generate card data]in the Amazon Web
// Services Payment Cryptography User Guide.
//
// This operation generates a CVV or CSC value that is printed on a payment credit
// or debit card during card production. The CVV or CSC, PAN (Primary Account
// Number) and expiration date of the card are required to check its validity
// during transaction processing. To begin this operation, a CVK (Card Verification
// Key) encryption key is required. You can use [CreateKey]or [ImportKey] to establish a CVK within
// Amazon Web Services Payment Cryptography. The KeyModesOfUse should be set to
// Generate and Verify for a CVK encryption key.
//
// For information about valid keys for this operation, see [Understanding key attributes] and [Key types for specific data operations] in the Amazon
// Web Services Payment Cryptography User Guide.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// [ImportKey]
//
// # VerifyCardValidationData
//
// [Generate card data]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/generate-card-data.html
// [ImportKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_ImportKey.html
// [Key types for specific data operations]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/crypto-ops-validkeys-ops.html
// [Understanding key attributes]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-validattributes.html
// [CreateKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_CreateKey.html
