package paymentcryptographydata

// VerifyMac is generated as a reference stub.
// Executable command wiring lives under cmd/paymentcryptographydata.go.
//
// Verifies a Message Authentication Code (MAC).
//
// You can use this operation to verify MAC for message data authentication such
// as . In this operation, you must use the same message data, secret encryption
// key and MAC algorithm that was used to generate MAC. You can use this operation
// to verify a DUPKT, CMAC, HMAC or EMV MAC by setting generation attributes and
// algorithm to the associated values.
//
// For information about valid keys for this operation, see [Understanding key attributes] and [Key types for specific data operations] in the Amazon
// Web Services Payment Cryptography User Guide.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// # GenerateMac
//
// [Key types for specific data operations]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/crypto-ops-validkeys-ops.html
// [Understanding key attributes]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-validattributes.html
