package paymentcryptographydata

// VerifyAuthRequestCryptogram is generated as a reference stub.
// Executable command wiring lives under cmd/paymentcryptographydata.go.
//
// Verifies Authorization Request Cryptogram (ARQC) for a EMV chip payment card
// authorization. For more information, see [Verify auth request cryptogram]in the Amazon Web Services Payment
// Cryptography User Guide.
//
// ARQC generation is done outside of Amazon Web Services Payment Cryptography and
// is typically generated on a point of sale terminal for an EMV chip card to
// obtain payment authorization during transaction time. For ARQC verification, you
// must first import the ARQC generated outside of Amazon Web Services Payment
// Cryptography by calling [ImportKey]. This operation uses the imported ARQC and an major
// encryption key (DUKPT) created by calling [CreateKey]to either provide a boolean ARQC
// verification result or provide an APRC (Authorization Response Cryptogram)
// response using Method 1 or Method 2. The ARPC_METHOD_1 uses AuthResponseCode to
// generate ARPC and ARPC_METHOD_2 uses CardStatusUpdate to generate ARPC.
//
// For information about valid keys for this operation, see [Understanding key attributes] and [Key types for specific data operations] in the Amazon
// Web Services Payment Cryptography User Guide.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// # VerifyCardValidationData
//
// # VerifyPinData
//
// [Verify auth request cryptogram]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/data-operations.verifyauthrequestcryptogram.html
// [ImportKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_ImportKey.html
// [Key types for specific data operations]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/crypto-ops-validkeys-ops.html
// [Understanding key attributes]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-validattributes.html
// [CreateKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_CreateKey.html
