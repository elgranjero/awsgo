package paymentcryptographydata

// VerifyCardValidationData is generated as a reference stub.
// Executable command wiring lives under cmd/paymentcryptographydata.go.
//
// Verifies card-related validation data using algorithms such as Card
// Verification Values (CVV/CVV2), Dynamic Card Verification Values (dCVV/dCVV2)
// and Card Security Codes (CSC). For more information, see [Verify card data]in the Amazon Web
// Services Payment Cryptography User Guide.
//
// This operation validates the CVV or CSC codes that is printed on a payment
// credit or debit card during card payment transaction. The input values are
// typically provided as part of an inbound transaction to an issuer or supporting
// platform partner. Amazon Web Services Payment Cryptography uses CVV or CSC, PAN
// (Primary Account Number) and expiration date of the card to check its validity
// during transaction processing. In this operation, the CVK (Card Verification
// Key) encryption key for use with card data verification is same as the one in
// used for GenerateCardValidationData.
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
// # VerifyAuthRequestCryptogram
//
// # VerifyPinData
//
// [Verify card data]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/verify-card-data.html
// [Key types for specific data operations]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/crypto-ops-validkeys-ops.html
// [Understanding key attributes]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-validattributes.html
