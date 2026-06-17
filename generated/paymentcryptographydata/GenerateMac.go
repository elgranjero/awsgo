package paymentcryptographydata

// GenerateMac is generated as a reference stub.
// Executable command wiring lives under cmd/paymentcryptographydata.go.
//
// Generates a Message Authentication Code (MAC) cryptogram within Amazon Web
// Services Payment Cryptography.
//
// You can use this operation to authenticate card-related data by using known
// data values to generate MAC for data validation between the sending and
// receiving parties. This operation uses message data, a secret encryption key and
// MAC algorithm to generate a unique MAC value for transmission. The receiving
// party of the MAC must use the same message data, secret encryption key and MAC
// algorithm to reproduce another MAC value for comparision.
//
// You can use this operation to generate a DUPKT, CMAC, HMAC or EMV MAC by
// setting generation attributes and algorithm to the associated values. The MAC
// generation encryption key must have valid values for KeyUsage such as
// TR31_M7_HMAC_KEY for HMAC generation, and the key must have KeyModesOfUse set
// to Generate .
//
// For information about valid keys for this operation, see [Understanding key attributes] and [Key types for specific data operations] in the Amazon
// Web Services Payment Cryptography User Guide.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// # VerifyMac
//
// [Key types for specific data operations]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/crypto-ops-validkeys-ops.html
// [Understanding key attributes]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-validattributes.html
