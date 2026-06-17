package paymentcryptographydata

// TranslatePinData is generated as a reference stub.
// Executable command wiring lives under cmd/paymentcryptographydata.go.
//
// Translates encrypted PIN block from and to ISO 9564 formats 0,1,3,4. For more
// information, see [Translate PIN data]in the Amazon Web Services Payment Cryptography User Guide.
//
// PIN block translation involves changing a PIN block from one encryption key to
// another and optionally change its format. PIN block translation occurs entirely
// within the HSM boundary and PIN data never enters or leaves Amazon Web Services
// Payment Cryptography in clear text. The encryption key transformation can be
// from PEK (Pin Encryption Key) to BDK (Base Derivation Key) for DUKPT or from BDK
// for DUKPT to PEK.
//
// Amazon Web Services Payment Cryptography also supports use of dynamic keys and
// ECDH (Elliptic Curve Diffie-Hellman) based key exchange for this operation.
//
// Dynamic keys allow you to pass a PEK as a TR-31 WrappedKeyBlock. They can be
// used when key material is frequently rotated, such as during every card
// transaction, and there is need to avoid importing short-lived keys into Amazon
// Web Services Payment Cryptography. To translate PIN block using dynamic keys,
// the keyARN is the Key Encryption Key (KEK) of the TR-31 wrapped PEK. The
// incoming wrapped key shall have a key purpose of P0 with a mode of use of B or
// D. For more information, see [Using Dynamic Keys]in the Amazon Web Services Payment Cryptography
// User Guide.
//
// Using ECDH key exchange, you can receive cardholder selectable PINs into Amazon
// Web Services Payment Cryptography. The ECDH derived key protects the incoming
// PIN block, which is translated to a PEK encrypted PIN block for use within the
// service. You can also use ECDH for reveal PIN, wherein the service translates
// the PIN block from PEK to a ECDH derived encryption key. For more information on
// establishing ECDH derived keys, see the [Creating keys]in the Amazon Web Services Payment
// Cryptography User Guide.
//
// The allowed combinations of PIN block format translations are guided by PCI. It
// is important to note that not all encrypted PIN block formats (example, format
// 1) require PAN (Primary Account Number) as input. And as such, PIN block format
// that requires PAN (example, formats 0,3,4) cannot be translated to a format
// (format 1) that does not require a PAN for generation.
//
// For information about valid keys for this operation, see [Understanding key attributes] and [Key types for specific data operations] in the Amazon
// Web Services Payment Cryptography User Guide.
//
// Amazon Web Services Payment Cryptography currently supports ISO PIN block 4
// translation for PIN block built using legacy PAN length. That is, PAN is the
// right most 12 digits excluding the check digits.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// # GeneratePinData
//
// # VerifyPinData
//
// [Using Dynamic Keys]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/use-cases-acquirers-dynamickeys.html
// [Creating keys]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/create-keys.html
// [Translate PIN data]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/translate-pin-data.html
// [Key types for specific data operations]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/crypto-ops-validkeys-ops.html
// [Understanding key attributes]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-validattributes.html
