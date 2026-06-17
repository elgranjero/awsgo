package paymentcryptographydata

// GenerateAs2805KekValidation is generated as a reference stub.
// Executable command wiring lives under cmd/paymentcryptographydata.go.
//
// Establishes node-to-node initialization between payment processing nodes such
// as an acquirer, issuer or payment network using Australian Standard 2805
// (AS2805).
//
// During node-to-node initialization, both communicating nodes must validate that
// they possess the correct Key Encrypting Keys (KEKs) before proceeding with
// session key exchange. In AS2805, the sending KEK (KEKs) of one node corresponds
// to the receiving KEK (KEKr) of its partner node. Each node uses its KEK to
// encrypt and decrypt session keys exchanged between the nodes. A KEK can be
// created or imported into Amazon Web Services Payment Cryptography using either
// the [CreateKey]or [ImportKey] operations.
//
// The node initiating communication can use GenerateAS2805KekValidation to
// generate a combined KEK validation request and KEK validation response to send
// to the partnering node for validation. When invoked, the API internally
// generates a random sending key encrypted under KEKs and provides a receiving key
// encrypted under KEKr as response. The initiating node sends the response
// returned by this API to its partner for validation.
//
// For information about valid keys for this operation, see [Understanding key attributes] and [Key types for specific data operations] in the Amazon
// Web Services Payment Cryptography User Guide.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// [ImportKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_ImportKey.html
// [Key types for specific data operations]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/crypto-ops-validkeys-ops.html
// [Understanding key attributes]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-validattributes.html
// [CreateKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_CreateKey.html
