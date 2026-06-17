package paymentcryptographydata

// GenerateMacEmvPinChange is generated as a reference stub.
// Executable command wiring lives under cmd/paymentcryptographydata.go.
//
// Generates an issuer script mac for EMV payment cards that use offline PINs as
// the cardholder verification method (CVM).
//
// This operation generates an authenticated issuer script response by appending
// the incoming message data (APDU command) with the target encrypted PIN block in
// ISO2 format. The command structure and method to send the issuer script update
// to the card is not defined by this operation and is typically determined by the
// applicable payment card scheme.
//
// The primary inputs to this operation include the incoming new encrypted
// pinblock, PIN encryption key (PEK), issuer master key (IMK), primary account
// number (PAN), and the payment card derivation method.
//
// The operation uses two issuer master keys - secure messaging for
// confidentiality (IMK-SMC) and secure messaging for integrity (IMK-SMI). The SMC
// key is used to internally derive a key to secure the pin, while SMI key is used
// to internally derive a key to authenticate the script reponse as per the [EMV 4.4 - Book 2 - Security and Key Management]
// specification.
//
// This operation supports Amex, EMV2000, EMVCommon, Mastercard and Visa
// derivation methods, each requiring specific input parameters. Users must follow
// the specific derivation method and input parameters defined by the respective
// payment card scheme.
//
// Use GenerateMac operation when sending a script update to an EMV card that does not
// involve PIN change. When assigning IAM permissions, it is important to
// understand that EncryptDatausing EMV keys and GenerateMac perform similar functions to this command.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// # EncryptData
//
// # GenerateMac
//
// [EMV 4.4 - Book 2 - Security and Key Management]: https://www.emvco.com/specifications/
