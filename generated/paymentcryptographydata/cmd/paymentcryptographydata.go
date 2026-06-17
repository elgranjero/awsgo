package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/paymentcryptographydata"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// paymentcryptographydataCmd represents the paymentcryptographydata command
var _paymentcryptographydataCmd = &cobra.Command{
	Use:   "paymentcryptographydata",
	Short: "AWS paymentcryptographydata CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := paymentcryptographydata.NewFromConfig(cfg)
		if _paymentcryptographydataDecryptData {
			paymentcryptographydata_DecryptData(cfg, client)
			return
		}
		if _paymentcryptographydataEncryptData {
			paymentcryptographydata_EncryptData(cfg, client)
			return
		}
		if _paymentcryptographydataGenerateAs2805KekValidation {
			paymentcryptographydata_GenerateAs2805KekValidation(cfg, client)
			return
		}
		if _paymentcryptographydataGenerateCardValidationData {
			paymentcryptographydata_GenerateCardValidationData(cfg, client)
			return
		}
		if _paymentcryptographydataGenerateMac {
			paymentcryptographydata_GenerateMac(cfg, client)
			return
		}
		if _paymentcryptographydataGenerateMacEmvPinChange {
			paymentcryptographydata_GenerateMacEmvPinChange(cfg, client)
			return
		}
		if _paymentcryptographydataGeneratePinData {
			paymentcryptographydata_GeneratePinData(cfg, client)
			return
		}
		if _paymentcryptographydataReEncryptData {
			paymentcryptographydata_ReEncryptData(cfg, client)
			return
		}
		if _paymentcryptographydataTranslateKeyMaterial {
			paymentcryptographydata_TranslateKeyMaterial(cfg, client)
			return
		}
		if _paymentcryptographydataTranslatePinData {
			paymentcryptographydata_TranslatePinData(cfg, client)
			return
		}
		if _paymentcryptographydataVerifyAuthRequestCryptogram {
			paymentcryptographydata_VerifyAuthRequestCryptogram(cfg, client)
			return
		}
		if _paymentcryptographydataVerifyCardValidationData {
			paymentcryptographydata_VerifyCardValidationData(cfg, client)
			return
		}
		if _paymentcryptographydataVerifyMac {
			paymentcryptographydata_VerifyMac(cfg, client)
			return
		}
		if _paymentcryptographydataVerifyPinData {
			paymentcryptographydata_VerifyPinData(cfg, client)
			return
		}

	},
}

var (
	_paymentcryptographydataDecryptData                 bool
	_paymentcryptographydataEncryptData                 bool
	_paymentcryptographydataGenerateAs2805KekValidation bool
	_paymentcryptographydataGenerateCardValidationData  bool
	_paymentcryptographydataGenerateMac                 bool
	_paymentcryptographydataGenerateMacEmvPinChange     bool
	_paymentcryptographydataGeneratePinData             bool
	_paymentcryptographydataReEncryptData               bool
	_paymentcryptographydataTranslateKeyMaterial        bool
	_paymentcryptographydataTranslatePinData            bool
	_paymentcryptographydataVerifyAuthRequestCryptogram bool
	_paymentcryptographydataVerifyCardValidationData    bool
	_paymentcryptographydataVerifyMac                   bool
	_paymentcryptographydataVerifyPinData               bool

	_paymentcryptographydataAuthRequestCryptogram                       string
	_paymentcryptographydataAuthResponseAttributes                      string
	_paymentcryptographydataCipherText                                  string
	_paymentcryptographydataDecryptionAttributes                        string
	_paymentcryptographydataDerivationMethodAttributes                  string
	_paymentcryptographydataDukptAttributes                             string
	_paymentcryptographydataEncryptedPinBlock                           string
	_paymentcryptographydataEncryptionAttributes                        string
	_paymentcryptographydataEncryptionKeyIdentifier                     string
	_paymentcryptographydataEncryptionWrappedKey                        string
	_paymentcryptographydataGenerationAttributes                        string
	_paymentcryptographydataGenerationKeyIdentifier                     string
	_paymentcryptographydataIncomingAs2805Attributes                    string
	_paymentcryptographydataIncomingDukptAttributes                     string
	_paymentcryptographydataIncomingEncryptionAttributes                string
	_paymentcryptographydataIncomingKeyIdentifier                       string
	_paymentcryptographydataIncomingKeyMaterial                         string
	_paymentcryptographydataIncomingTranslationAttributes               string
	_paymentcryptographydataIncomingWrappedKey                          string
	_paymentcryptographydataKekValidationType                           string
	_paymentcryptographydataKeyCheckValueAlgorithm                      string
	_paymentcryptographydataKeyIdentifier                               string
	_paymentcryptographydataMac                                         string
	_paymentcryptographydataMacLength                                   string
	_paymentcryptographydataMajorKeyDerivationMode                      string
	_paymentcryptographydataMessageData                                 string
	_paymentcryptographydataNewEncryptedPinBlock                        string
	_paymentcryptographydataNewPinPekIdentifier                         string
	_paymentcryptographydataOutgoingDukptAttributes                     string
	_paymentcryptographydataOutgoingEncryptionAttributes                string
	_paymentcryptographydataOutgoingKeyIdentifier                       string
	_paymentcryptographydataOutgoingKeyMaterial                         string
	_paymentcryptographydataOutgoingTranslationAttributes               string
	_paymentcryptographydataOutgoingWrappedKey                          string
	_paymentcryptographydataPinBlockFormat                              string
	_paymentcryptographydataPinDataLength                               string
	_paymentcryptographydataPlainText                                   string
	_paymentcryptographydataPrimaryAccountNumber                        string
	_paymentcryptographydataRandomKeySendVariantMask                    string
	_paymentcryptographydataSecureMessagingConfidentialityKeyIdentifier string
	_paymentcryptographydataSecureMessagingIntegrityKeyIdentifier       string
	_paymentcryptographydataSessionKeyDerivationAttributes              string
	_paymentcryptographydataTransactionData                             string
	_paymentcryptographydataValidationData                              string
	_paymentcryptographydataValidationDataLength                        string
	_paymentcryptographydataVerificationAttributes                      string
	_paymentcryptographydataVerificationKeyIdentifier                   string
	_paymentcryptographydataWrappedKey                                  string
)

// Decrypts ciphertext data to plaintext using a symmetric (TDES, AES), asymmetric
// (RSA), or derived (DUKPT or EMV) encryption key scheme. For more information,
// see [Decrypt data]in the Amazon Web Services Payment Cryptography User Guide.
//
// You can use an decryption key generated within Amazon Web Services Payment
// Cryptography, or you can import your own decryption key by calling [ImportKey]. For this
// operation, the key must have KeyModesOfUse set to Decrypt . In asymmetric
// decryption, Amazon Web Services Payment Cryptography decrypts the ciphertext
// using the private component of the asymmetric encryption key pair. For data
// encryption outside of Amazon Web Services Payment Cryptography, you can export
// the public component of the asymmetric key pair by calling [GetPublicCertificate].
//
// This operation also supports dynamic keys, allowing you to pass a dynamic
// decryption key as a TR-31 WrappedKeyBlock. This can be used when key material is
// frequently rotated, such as during every card transaction, and there is need to
// avoid importing short-lived keys into Amazon Web Services Payment Cryptography.
// To decrypt using dynamic keys, the keyARN is the Key Encryption Key (KEK) of
// the TR-31 wrapped decryption key material. The incoming wrapped key shall have a
// key purpose of D0 with a mode of use of B or D. For more information, see [Using Dynamic Keys]in
// the Amazon Web Services Payment Cryptography User Guide.
//
// For symmetric and DUKPT decryption, Amazon Web Services Payment Cryptography
// supports TDES and AES algorithms. For EMV decryption, Amazon Web Services
// Payment Cryptography supports TDES algorithms. For asymmetric decryption,
// Amazon Web Services Payment Cryptography supports RSA .
//
// When you use TDES or TDES DUKPT, the ciphertext data length must be a multiple
// of 8 bytes. For AES or AES DUKPT, the ciphertext data length must be a multiple
// of 16 bytes. For RSA, it sould be equal to the key size unless padding is
// enabled.
//
// For information about valid keys for this operation, see [Understanding key attributes] and [Key types for specific data operations] in the Amazon
// Web Services Payment Cryptography User Guide.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// # EncryptData
//
// [GetPublicCertificate]
//
// [ImportKey]
//
// [Using Dynamic Keys]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/use-cases-acquirers-dynamickeys.html
// [GetPublicCertificate]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_GetPublicKeyCertificate.html
// [ImportKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_ImportKey.html
// [Key types for specific data operations]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/crypto-ops-validkeys-ops.html
// [Decrypt data]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/decrypt-data.html
// [Understanding key attributes]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-validattributes.html
func paymentcryptographydata_DecryptData(cfg aws.Config, client *paymentcryptographydata.Client) {
	input := &paymentcryptographydata.DecryptDataInput{
		// CipherText: *string, // Required
		// DecryptionAttributes: types.EncryptionDecryptionAttributes, // Required
		// KeyIdentifier: *string, // Required
	}

	if len(_paymentcryptographydataCipherText) > 0 {
		input.CipherText = aws.String(_paymentcryptographydataCipherText)
	}
	if len(_paymentcryptographydataDecryptionAttributes) > 0 {
		if err := assignInputField(input, "DecryptionAttributes", _paymentcryptographydataDecryptionAttributes); err != nil {
			log.Errorf("invalid --decryption-attributes: %s", err.Error())
			return
		}
	}
	if len(_paymentcryptographydataKeyIdentifier) > 0 {
		input.KeyIdentifier = aws.String(_paymentcryptographydataKeyIdentifier)
	}
	if len(_paymentcryptographydataWrappedKey) > 0 {
		if err := assignInputField(input, "WrappedKey", _paymentcryptographydataWrappedKey); err != nil {
			log.Errorf("invalid --wrapped-key: %s", err.Error())
			return
		}
	}

	if resp, err := client.DecryptData(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Encrypts plaintext data to ciphertext using a symmetric (TDES, AES), asymmetric
// (RSA), or derived (DUKPT or EMV) encryption key scheme. For more information,
// see [Encrypt data]in the Amazon Web Services Payment Cryptography User Guide.
//
// You can generate an encryption key within Amazon Web Services Payment
// Cryptography by calling [CreateKey]. You can import your own encryption key by calling [ImportKey].
//
// For this operation, the key must have KeyModesOfUse set to Encrypt . In
// asymmetric encryption, plaintext is encrypted using public component. You can
// import the public component of an asymmetric key pair created outside Amazon Web
// Services Payment Cryptography by calling [ImportKey].
//
// This operation also supports dynamic keys, allowing you to pass a dynamic
// encryption key as a TR-31 WrappedKeyBlock. This can be used when key material is
// frequently rotated, such as during every card transaction, and there is need to
// avoid importing short-lived keys into Amazon Web Services Payment Cryptography.
// To encrypt using dynamic keys, the keyARN is the Key Encryption Key (KEK) of
// the TR-31 wrapped encryption key material. The incoming wrapped key shall have a
// key purpose of D0 with a mode of use of B or D. For more information, see [Using Dynamic Keys]in
// the Amazon Web Services Payment Cryptography User Guide.
//
// For symmetric and DUKPT encryption, Amazon Web Services Payment Cryptography
// supports TDES and AES algorithms. For EMV encryption, Amazon Web Services
// Payment Cryptography supports TDES algorithms.For asymmetric encryption, Amazon
// Web Services Payment Cryptography supports RSA .
//
// When you use TDES or TDES DUKPT, the plaintext data length must be a multiple
// of 8 bytes. For AES or AES DUKPT, the plaintext data length must be a multiple
// of 16 bytes. For RSA, it sould be equal to the key size unless padding is
// enabled.
//
// To encrypt using DUKPT, you must already have a BDK (Base Derivation Key) key
// in your account with KeyModesOfUse set to DeriveKey , or you can generate a new
// DUKPT key by calling [CreateKey]. To encrypt using EMV, you must already have an IMK
// (Issuer Master Key) key in your account with KeyModesOfUse set to DeriveKey .
//
// For information about valid keys for this operation, see [Understanding key attributes] and [Key types for specific data operations] in the Amazon
// Web Services Payment Cryptography User Guide.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// # DecryptData
//
// [GetPublicCertificate]
//
// [ImportKey]
//
// # ReEncryptData
//
// [Using Dynamic Keys]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/use-cases-acquirers-dynamickeys.html
// [GetPublicCertificate]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_GetPublicKeyCertificate.html
// [Encrypt data]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/encrypt-data.html
// [ImportKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_ImportKey.html
// [Key types for specific data operations]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/crypto-ops-validkeys-ops.html
// [Understanding key attributes]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-validattributes.html
// [CreateKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_CreateKey.html
func paymentcryptographydata_EncryptData(cfg aws.Config, client *paymentcryptographydata.Client) {
	input := &paymentcryptographydata.EncryptDataInput{
		// EncryptionAttributes: types.EncryptionDecryptionAttributes, // Required
		// KeyIdentifier: *string, // Required
		// PlainText: *string, // Required
	}

	if len(_paymentcryptographydataEncryptionAttributes) > 0 {
		if err := assignInputField(input, "EncryptionAttributes", _paymentcryptographydataEncryptionAttributes); err != nil {
			log.Errorf("invalid --encryption-attributes: %s", err.Error())
			return
		}
	}
	if len(_paymentcryptographydataKeyIdentifier) > 0 {
		input.KeyIdentifier = aws.String(_paymentcryptographydataKeyIdentifier)
	}
	if len(_paymentcryptographydataPlainText) > 0 {
		input.PlainText = aws.String(_paymentcryptographydataPlainText)
	}
	if len(_paymentcryptographydataWrappedKey) > 0 {
		if err := assignInputField(input, "WrappedKey", _paymentcryptographydataWrappedKey); err != nil {
			log.Errorf("invalid --wrapped-key: %s", err.Error())
			return
		}
	}

	if resp, err := client.EncryptData(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

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
func paymentcryptographydata_GenerateAs2805KekValidation(cfg aws.Config, client *paymentcryptographydata.Client) {
	input := &paymentcryptographydata.GenerateAs2805KekValidationInput{
		// KekValidationType: types.As2805KekValidationType, // Required
		// KeyIdentifier: *string, // Required
		// RandomKeySendVariantMask: types.RandomKeySendVariantMask, // Required
	}

	if len(_paymentcryptographydataKekValidationType) > 0 {
		if err := assignInputField(input, "KekValidationType", _paymentcryptographydataKekValidationType); err != nil {
			log.Errorf("invalid --kek-validation-type: %s", err.Error())
			return
		}
	}
	if len(_paymentcryptographydataKeyIdentifier) > 0 {
		input.KeyIdentifier = aws.String(_paymentcryptographydataKeyIdentifier)
	}
	if len(_paymentcryptographydataRandomKeySendVariantMask) > 0 {
		if err := assignInputField(input, "RandomKeySendVariantMask", _paymentcryptographydataRandomKeySendVariantMask); err != nil {
			log.Errorf("invalid --random-key-send-variant-mask: %s", err.Error())
			return
		}
	}

	if resp, err := client.GenerateAs2805KekValidation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

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
func paymentcryptographydata_GenerateCardValidationData(cfg aws.Config, client *paymentcryptographydata.Client) {
	input := &paymentcryptographydata.GenerateCardValidationDataInput{
		// GenerationAttributes: types.CardGenerationAttributes, // Required
		// KeyIdentifier: *string, // Required
		// PrimaryAccountNumber: *string, // Required
	}

	if len(_paymentcryptographydataGenerationAttributes) > 0 {
		if err := assignInputField(input, "GenerationAttributes", _paymentcryptographydataGenerationAttributes); err != nil {
			log.Errorf("invalid --generation-attributes: %s", err.Error())
			return
		}
	}
	if len(_paymentcryptographydataKeyIdentifier) > 0 {
		input.KeyIdentifier = aws.String(_paymentcryptographydataKeyIdentifier)
	}
	if len(_paymentcryptographydataPrimaryAccountNumber) > 0 {
		input.PrimaryAccountNumber = aws.String(_paymentcryptographydataPrimaryAccountNumber)
	}
	if len(_paymentcryptographydataValidationDataLength) > 0 {
		if err := assignInputField(input, "ValidationDataLength", _paymentcryptographydataValidationDataLength); err != nil {
			log.Errorf("invalid --validation-data-length: %s", err.Error())
			return
		}
	}

	if resp, err := client.GenerateCardValidationData(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

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
func paymentcryptographydata_GenerateMac(cfg aws.Config, client *paymentcryptographydata.Client) {
	input := &paymentcryptographydata.GenerateMacInput{
		// GenerationAttributes: types.MacAttributes, // Required
		// KeyIdentifier: *string, // Required
		// MessageData: *string, // Required
	}

	if len(_paymentcryptographydataGenerationAttributes) > 0 {
		if err := assignInputField(input, "GenerationAttributes", _paymentcryptographydataGenerationAttributes); err != nil {
			log.Errorf("invalid --generation-attributes: %s", err.Error())
			return
		}
	}
	if len(_paymentcryptographydataKeyIdentifier) > 0 {
		input.KeyIdentifier = aws.String(_paymentcryptographydataKeyIdentifier)
	}
	if len(_paymentcryptographydataMessageData) > 0 {
		input.MessageData = aws.String(_paymentcryptographydataMessageData)
	}
	if len(_paymentcryptographydataMacLength) > 0 {
		if err := assignInputField(input, "MacLength", _paymentcryptographydataMacLength); err != nil {
			log.Errorf("invalid --mac-length: %s", err.Error())
			return
		}
	}

	if resp, err := client.GenerateMac(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

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
func paymentcryptographydata_GenerateMacEmvPinChange(cfg aws.Config, client *paymentcryptographydata.Client) {
	input := &paymentcryptographydata.GenerateMacEmvPinChangeInput{
		// DerivationMethodAttributes: types.DerivationMethodAttributes, // Required
		// MessageData: *string, // Required
		// NewEncryptedPinBlock: *string, // Required
		// NewPinPekIdentifier: *string, // Required
		// PinBlockFormat: types.PinBlockFormatForEmvPinChange, // Required
		// SecureMessagingConfidentialityKeyIdentifier: *string, // Required
		// SecureMessagingIntegrityKeyIdentifier: *string, // Required
	}

	if len(_paymentcryptographydataDerivationMethodAttributes) > 0 {
		if err := assignInputField(input, "DerivationMethodAttributes", _paymentcryptographydataDerivationMethodAttributes); err != nil {
			log.Errorf("invalid --derivation-method-attributes: %s", err.Error())
			return
		}
	}
	if len(_paymentcryptographydataMessageData) > 0 {
		input.MessageData = aws.String(_paymentcryptographydataMessageData)
	}
	if len(_paymentcryptographydataNewEncryptedPinBlock) > 0 {
		input.NewEncryptedPinBlock = aws.String(_paymentcryptographydataNewEncryptedPinBlock)
	}
	if len(_paymentcryptographydataNewPinPekIdentifier) > 0 {
		input.NewPinPekIdentifier = aws.String(_paymentcryptographydataNewPinPekIdentifier)
	}
	if len(_paymentcryptographydataPinBlockFormat) > 0 {
		if err := assignInputField(input, "PinBlockFormat", _paymentcryptographydataPinBlockFormat); err != nil {
			log.Errorf("invalid --pin-block-format: %s", err.Error())
			return
		}
	}
	if len(_paymentcryptographydataSecureMessagingConfidentialityKeyIdentifier) > 0 {
		input.SecureMessagingConfidentialityKeyIdentifier = aws.String(_paymentcryptographydataSecureMessagingConfidentialityKeyIdentifier)
	}
	if len(_paymentcryptographydataSecureMessagingIntegrityKeyIdentifier) > 0 {
		input.SecureMessagingIntegrityKeyIdentifier = aws.String(_paymentcryptographydataSecureMessagingIntegrityKeyIdentifier)
	}

	if resp, err := client.GenerateMacEmvPinChange(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

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
func paymentcryptographydata_GeneratePinData(cfg aws.Config, client *paymentcryptographydata.Client) {
	input := &paymentcryptographydata.GeneratePinDataInput{
		// EncryptionKeyIdentifier: *string, // Required
		// GenerationAttributes: types.PinGenerationAttributes, // Required
		// GenerationKeyIdentifier: *string, // Required
		// PinBlockFormat: types.PinBlockFormatForPinData, // Required
	}

	if len(_paymentcryptographydataEncryptionKeyIdentifier) > 0 {
		input.EncryptionKeyIdentifier = aws.String(_paymentcryptographydataEncryptionKeyIdentifier)
	}
	if len(_paymentcryptographydataGenerationAttributes) > 0 {
		if err := assignInputField(input, "GenerationAttributes", _paymentcryptographydataGenerationAttributes); err != nil {
			log.Errorf("invalid --generation-attributes: %s", err.Error())
			return
		}
	}
	if len(_paymentcryptographydataGenerationKeyIdentifier) > 0 {
		input.GenerationKeyIdentifier = aws.String(_paymentcryptographydataGenerationKeyIdentifier)
	}
	if len(_paymentcryptographydataPinBlockFormat) > 0 {
		if err := assignInputField(input, "PinBlockFormat", _paymentcryptographydataPinBlockFormat); err != nil {
			log.Errorf("invalid --pin-block-format: %s", err.Error())
			return
		}
	}
	if len(_paymentcryptographydataEncryptionWrappedKey) > 0 {
		if err := assignInputField(input, "EncryptionWrappedKey", _paymentcryptographydataEncryptionWrappedKey); err != nil {
			log.Errorf("invalid --encryption-wrapped-key: %s", err.Error())
			return
		}
	}
	if len(_paymentcryptographydataPinDataLength) > 0 {
		if err := assignInputField(input, "PinDataLength", _paymentcryptographydataPinDataLength); err != nil {
			log.Errorf("invalid --pin-data-length: %s", err.Error())
			return
		}
	}
	if len(_paymentcryptographydataPrimaryAccountNumber) > 0 {
		input.PrimaryAccountNumber = aws.String(_paymentcryptographydataPrimaryAccountNumber)
	}

	if resp, err := client.GeneratePinData(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Re-encrypt ciphertext using DUKPT or Symmetric data encryption keys.
// You can either generate an encryption key within Amazon Web Services Payment
// Cryptography by calling [CreateKey]or import your own encryption key by calling [ImportKey]. The
// KeyArn for use with this operation must be in a compatible key state with
// KeyModesOfUse set to Encrypt .
//
// This operation also supports dynamic keys, allowing you to pass a dynamic
// encryption key as a TR-31 WrappedKeyBlock. This can be used when key material is
// frequently rotated, such as during every card transaction, and there is need to
// avoid importing short-lived keys into Amazon Web Services Payment Cryptography.
// To re-encrypt using dynamic keys, the keyARN is the Key Encryption Key (KEK) of
// the TR-31 wrapped encryption key material. The incoming wrapped key shall have a
// key purpose of D0 with a mode of use of B or D. For more information, see [Using Dynamic Keys]in
// the Amazon Web Services Payment Cryptography User Guide.
//
// For symmetric and DUKPT encryption, Amazon Web Services Payment Cryptography
// supports TDES and AES algorithms. To encrypt using DUKPT, a DUKPT key must
// already exist within your account with KeyModesOfUse set to DeriveKey or a new
// DUKPT can be generated by calling [CreateKey].
//
// For information about valid keys for this operation, see [Understanding key attributes] and [Key types for specific data operations] in the Amazon
// Web Services Payment Cryptography User Guide.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// # DecryptData
//
// # EncryptData
//
// [GetPublicCertificate]
//
// [ImportKey]
//
// [Using Dynamic Keys]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/use-cases-acquirers-dynamickeys.html
// [GetPublicCertificate]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_GetPublicKeyCertificate.html
// [ImportKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_ImportKey.html
// [Key types for specific data operations]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/crypto-ops-validkeys-ops.html
// [Understanding key attributes]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-validattributes.html
// [CreateKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_CreateKey.html
func paymentcryptographydata_ReEncryptData(cfg aws.Config, client *paymentcryptographydata.Client) {
	input := &paymentcryptographydata.ReEncryptDataInput{
		// CipherText: *string, // Required
		// IncomingEncryptionAttributes: types.ReEncryptionAttributes, // Required
		// IncomingKeyIdentifier: *string, // Required
		// OutgoingEncryptionAttributes: types.ReEncryptionAttributes, // Required
		// OutgoingKeyIdentifier: *string, // Required
	}

	if len(_paymentcryptographydataCipherText) > 0 {
		input.CipherText = aws.String(_paymentcryptographydataCipherText)
	}
	if len(_paymentcryptographydataIncomingEncryptionAttributes) > 0 {
		if err := assignInputField(input, "IncomingEncryptionAttributes", _paymentcryptographydataIncomingEncryptionAttributes); err != nil {
			log.Errorf("invalid --incoming-encryption-attributes: %s", err.Error())
			return
		}
	}
	if len(_paymentcryptographydataIncomingKeyIdentifier) > 0 {
		input.IncomingKeyIdentifier = aws.String(_paymentcryptographydataIncomingKeyIdentifier)
	}
	if len(_paymentcryptographydataOutgoingEncryptionAttributes) > 0 {
		if err := assignInputField(input, "OutgoingEncryptionAttributes", _paymentcryptographydataOutgoingEncryptionAttributes); err != nil {
			log.Errorf("invalid --outgoing-encryption-attributes: %s", err.Error())
			return
		}
	}
	if len(_paymentcryptographydataOutgoingKeyIdentifier) > 0 {
		input.OutgoingKeyIdentifier = aws.String(_paymentcryptographydataOutgoingKeyIdentifier)
	}
	if len(_paymentcryptographydataIncomingWrappedKey) > 0 {
		if err := assignInputField(input, "IncomingWrappedKey", _paymentcryptographydataIncomingWrappedKey); err != nil {
			log.Errorf("invalid --incoming-wrapped-key: %s", err.Error())
			return
		}
	}
	if len(_paymentcryptographydataOutgoingWrappedKey) > 0 {
		if err := assignInputField(input, "OutgoingWrappedKey", _paymentcryptographydataOutgoingWrappedKey); err != nil {
			log.Errorf("invalid --outgoing-wrapped-key: %s", err.Error())
			return
		}
	}

	if resp, err := client.ReEncryptData(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Translates an cryptographic key between different wrapping keys without
// importing the key into Amazon Web Services Payment Cryptography.
//
// This operation can be used when key material is frequently rotated, such as
// during every card transaction, and there is a need to avoid importing
// short-lived keys into Amazon Web Services Payment Cryptography. It translates
// short-lived transaction keys such as [PEK]generated for each transaction and wrapped
// with an [ECDH]derived wrapping key to another [KEK] wrapping key.
//
// Before using this operation, you must first request the public key certificate
// of the ECC key pair generated within Amazon Web Services Payment Cryptography to
// establish an ECDH key agreement. In TranslateKeyData , the service uses its own
// ECC key pair, public certificate of receiving ECC key pair, and the key
// derivation parameters to generate a derived key. The service uses this derived
// key to unwrap the incoming transaction key received as a TR31WrappedKeyBlock and
// re-wrap using a user provided KEK to generate an outgoing Tr31WrappedKeyBlock.
//
// For information about valid keys for this operation, see [Understanding key attributes] and [Key types for specific data operations] in the Amazon
// Web Services Payment Cryptography User Guide.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// [CreateKey]
//
// [GetPublicCertificate]
//
// [ImportKey]
//
// [KEK]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/terminology.html#terms.kek
// [ECDH]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/terminology.html#terms.ecdh
// [GetPublicCertificate]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_GetPublicKeyCertificate.html
// [PEK]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/terminology.html#terms.pek
// [Key types for specific data operations]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/crypto-ops-validkeys-ops.html
// [ImportKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_ImportKey.html
// [Understanding key attributes]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-validattributes.html
// [CreateKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_CreateKey.html
func paymentcryptographydata_TranslateKeyMaterial(cfg aws.Config, client *paymentcryptographydata.Client) {
	input := &paymentcryptographydata.TranslateKeyMaterialInput{
		// IncomingKeyMaterial: types.IncomingKeyMaterial, // Required
		// OutgoingKeyMaterial: types.OutgoingKeyMaterial, // Required
	}

	if len(_paymentcryptographydataIncomingKeyMaterial) > 0 {
		if err := assignInputField(input, "IncomingKeyMaterial", _paymentcryptographydataIncomingKeyMaterial); err != nil {
			log.Errorf("invalid --incoming-key-material: %s", err.Error())
			return
		}
	}
	if len(_paymentcryptographydataOutgoingKeyMaterial) > 0 {
		if err := assignInputField(input, "OutgoingKeyMaterial", _paymentcryptographydataOutgoingKeyMaterial); err != nil {
			log.Errorf("invalid --outgoing-key-material: %s", err.Error())
			return
		}
	}
	if len(_paymentcryptographydataKeyCheckValueAlgorithm) > 0 {
		if err := assignInputField(input, "KeyCheckValueAlgorithm", _paymentcryptographydataKeyCheckValueAlgorithm); err != nil {
			log.Errorf("invalid --key-check-value-algorithm: %s", err.Error())
			return
		}
	}

	if resp, err := client.TranslateKeyMaterial(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

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
func paymentcryptographydata_TranslatePinData(cfg aws.Config, client *paymentcryptographydata.Client) {
	input := &paymentcryptographydata.TranslatePinDataInput{
		// EncryptedPinBlock: *string, // Required
		// IncomingKeyIdentifier: *string, // Required
		// IncomingTranslationAttributes: types.TranslationIsoFormats, // Required
		// OutgoingKeyIdentifier: *string, // Required
		// OutgoingTranslationAttributes: types.TranslationIsoFormats, // Required
	}

	if len(_paymentcryptographydataEncryptedPinBlock) > 0 {
		input.EncryptedPinBlock = aws.String(_paymentcryptographydataEncryptedPinBlock)
	}
	if len(_paymentcryptographydataIncomingKeyIdentifier) > 0 {
		input.IncomingKeyIdentifier = aws.String(_paymentcryptographydataIncomingKeyIdentifier)
	}
	if len(_paymentcryptographydataIncomingTranslationAttributes) > 0 {
		if err := assignInputField(input, "IncomingTranslationAttributes", _paymentcryptographydataIncomingTranslationAttributes); err != nil {
			log.Errorf("invalid --incoming-translation-attributes: %s", err.Error())
			return
		}
	}
	if len(_paymentcryptographydataOutgoingKeyIdentifier) > 0 {
		input.OutgoingKeyIdentifier = aws.String(_paymentcryptographydataOutgoingKeyIdentifier)
	}
	if len(_paymentcryptographydataOutgoingTranslationAttributes) > 0 {
		if err := assignInputField(input, "OutgoingTranslationAttributes", _paymentcryptographydataOutgoingTranslationAttributes); err != nil {
			log.Errorf("invalid --outgoing-translation-attributes: %s", err.Error())
			return
		}
	}
	if len(_paymentcryptographydataIncomingAs2805Attributes) > 0 {
		if err := assignInputField(input, "IncomingAs2805Attributes", _paymentcryptographydataIncomingAs2805Attributes); err != nil {
			log.Errorf("invalid --incoming-as2805-attributes: %s", err.Error())
			return
		}
	}
	if len(_paymentcryptographydataIncomingDukptAttributes) > 0 {
		if err := assignInputField(input, "IncomingDukptAttributes", _paymentcryptographydataIncomingDukptAttributes); err != nil {
			log.Errorf("invalid --incoming-dukpt-attributes: %s", err.Error())
			return
		}
	}
	if len(_paymentcryptographydataIncomingWrappedKey) > 0 {
		if err := assignInputField(input, "IncomingWrappedKey", _paymentcryptographydataIncomingWrappedKey); err != nil {
			log.Errorf("invalid --incoming-wrapped-key: %s", err.Error())
			return
		}
	}
	if len(_paymentcryptographydataOutgoingDukptAttributes) > 0 {
		if err := assignInputField(input, "OutgoingDukptAttributes", _paymentcryptographydataOutgoingDukptAttributes); err != nil {
			log.Errorf("invalid --outgoing-dukpt-attributes: %s", err.Error())
			return
		}
	}
	if len(_paymentcryptographydataOutgoingWrappedKey) > 0 {
		if err := assignInputField(input, "OutgoingWrappedKey", _paymentcryptographydataOutgoingWrappedKey); err != nil {
			log.Errorf("invalid --outgoing-wrapped-key: %s", err.Error())
			return
		}
	}

	if resp, err := client.TranslatePinData(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

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
func paymentcryptographydata_VerifyAuthRequestCryptogram(cfg aws.Config, client *paymentcryptographydata.Client) {
	input := &paymentcryptographydata.VerifyAuthRequestCryptogramInput{
		// AuthRequestCryptogram: *string, // Required
		// KeyIdentifier: *string, // Required
		// MajorKeyDerivationMode: types.MajorKeyDerivationMode, // Required
		// SessionKeyDerivationAttributes: types.SessionKeyDerivation, // Required
		// TransactionData: *string, // Required
	}

	if len(_paymentcryptographydataAuthRequestCryptogram) > 0 {
		input.AuthRequestCryptogram = aws.String(_paymentcryptographydataAuthRequestCryptogram)
	}
	if len(_paymentcryptographydataKeyIdentifier) > 0 {
		input.KeyIdentifier = aws.String(_paymentcryptographydataKeyIdentifier)
	}
	if len(_paymentcryptographydataMajorKeyDerivationMode) > 0 {
		if err := assignInputField(input, "MajorKeyDerivationMode", _paymentcryptographydataMajorKeyDerivationMode); err != nil {
			log.Errorf("invalid --major-key-derivation-mode: %s", err.Error())
			return
		}
	}
	if len(_paymentcryptographydataSessionKeyDerivationAttributes) > 0 {
		if err := assignInputField(input, "SessionKeyDerivationAttributes", _paymentcryptographydataSessionKeyDerivationAttributes); err != nil {
			log.Errorf("invalid --session-key-derivation-attributes: %s", err.Error())
			return
		}
	}
	if len(_paymentcryptographydataTransactionData) > 0 {
		input.TransactionData = aws.String(_paymentcryptographydataTransactionData)
	}
	if len(_paymentcryptographydataAuthResponseAttributes) > 0 {
		if err := assignInputField(input, "AuthResponseAttributes", _paymentcryptographydataAuthResponseAttributes); err != nil {
			log.Errorf("invalid --auth-response-attributes: %s", err.Error())
			return
		}
	}

	if resp, err := client.VerifyAuthRequestCryptogram(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

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
func paymentcryptographydata_VerifyCardValidationData(cfg aws.Config, client *paymentcryptographydata.Client) {
	input := &paymentcryptographydata.VerifyCardValidationDataInput{
		// KeyIdentifier: *string, // Required
		// PrimaryAccountNumber: *string, // Required
		// ValidationData: *string, // Required
		// VerificationAttributes: types.CardVerificationAttributes, // Required
	}

	if len(_paymentcryptographydataKeyIdentifier) > 0 {
		input.KeyIdentifier = aws.String(_paymentcryptographydataKeyIdentifier)
	}
	if len(_paymentcryptographydataPrimaryAccountNumber) > 0 {
		input.PrimaryAccountNumber = aws.String(_paymentcryptographydataPrimaryAccountNumber)
	}
	if len(_paymentcryptographydataValidationData) > 0 {
		input.ValidationData = aws.String(_paymentcryptographydataValidationData)
	}
	if len(_paymentcryptographydataVerificationAttributes) > 0 {
		if err := assignInputField(input, "VerificationAttributes", _paymentcryptographydataVerificationAttributes); err != nil {
			log.Errorf("invalid --verification-attributes: %s", err.Error())
			return
		}
	}

	if resp, err := client.VerifyCardValidationData(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Verifies a Message Authentication Code (MAC).
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
func paymentcryptographydata_VerifyMac(cfg aws.Config, client *paymentcryptographydata.Client) {
	input := &paymentcryptographydata.VerifyMacInput{
		// KeyIdentifier: *string, // Required
		// Mac: *string, // Required
		// MessageData: *string, // Required
		// VerificationAttributes: types.MacAttributes, // Required
	}

	if len(_paymentcryptographydataKeyIdentifier) > 0 {
		input.KeyIdentifier = aws.String(_paymentcryptographydataKeyIdentifier)
	}
	if len(_paymentcryptographydataMac) > 0 {
		input.Mac = aws.String(_paymentcryptographydataMac)
	}
	if len(_paymentcryptographydataMessageData) > 0 {
		input.MessageData = aws.String(_paymentcryptographydataMessageData)
	}
	if len(_paymentcryptographydataVerificationAttributes) > 0 {
		if err := assignInputField(input, "VerificationAttributes", _paymentcryptographydataVerificationAttributes); err != nil {
			log.Errorf("invalid --verification-attributes: %s", err.Error())
			return
		}
	}
	if len(_paymentcryptographydataMacLength) > 0 {
		if err := assignInputField(input, "MacLength", _paymentcryptographydataMacLength); err != nil {
			log.Errorf("invalid --mac-length: %s", err.Error())
			return
		}
	}

	if resp, err := client.VerifyMac(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

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
func paymentcryptographydata_VerifyPinData(cfg aws.Config, client *paymentcryptographydata.Client) {
	input := &paymentcryptographydata.VerifyPinDataInput{
		// EncryptedPinBlock: *string, // Required
		// EncryptionKeyIdentifier: *string, // Required
		// PinBlockFormat: types.PinBlockFormatForPinData, // Required
		// VerificationAttributes: types.PinVerificationAttributes, // Required
		// VerificationKeyIdentifier: *string, // Required
	}

	if len(_paymentcryptographydataEncryptedPinBlock) > 0 {
		input.EncryptedPinBlock = aws.String(_paymentcryptographydataEncryptedPinBlock)
	}
	if len(_paymentcryptographydataEncryptionKeyIdentifier) > 0 {
		input.EncryptionKeyIdentifier = aws.String(_paymentcryptographydataEncryptionKeyIdentifier)
	}
	if len(_paymentcryptographydataPinBlockFormat) > 0 {
		if err := assignInputField(input, "PinBlockFormat", _paymentcryptographydataPinBlockFormat); err != nil {
			log.Errorf("invalid --pin-block-format: %s", err.Error())
			return
		}
	}
	if len(_paymentcryptographydataVerificationAttributes) > 0 {
		if err := assignInputField(input, "VerificationAttributes", _paymentcryptographydataVerificationAttributes); err != nil {
			log.Errorf("invalid --verification-attributes: %s", err.Error())
			return
		}
	}
	if len(_paymentcryptographydataVerificationKeyIdentifier) > 0 {
		input.VerificationKeyIdentifier = aws.String(_paymentcryptographydataVerificationKeyIdentifier)
	}
	if len(_paymentcryptographydataDukptAttributes) > 0 {
		if err := assignInputField(input, "DukptAttributes", _paymentcryptographydataDukptAttributes); err != nil {
			log.Errorf("invalid --dukpt-attributes: %s", err.Error())
			return
		}
	}
	if len(_paymentcryptographydataEncryptionWrappedKey) > 0 {
		if err := assignInputField(input, "EncryptionWrappedKey", _paymentcryptographydataEncryptionWrappedKey); err != nil {
			log.Errorf("invalid --encryption-wrapped-key: %s", err.Error())
			return
		}
	}
	if len(_paymentcryptographydataPinDataLength) > 0 {
		if err := assignInputField(input, "PinDataLength", _paymentcryptographydataPinDataLength); err != nil {
			log.Errorf("invalid --pin-data-length: %s", err.Error())
			return
		}
	}
	if len(_paymentcryptographydataPrimaryAccountNumber) > 0 {
		input.PrimaryAccountNumber = aws.String(_paymentcryptographydataPrimaryAccountNumber)
	}

	if resp, err := client.VerifyPinData(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_paymentcryptographydataCmd)
	_paymentcryptographydataCmd.Flags().SortFlags = false

	_paymentcryptographydataCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_paymentcryptographydataCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_paymentcryptographydataCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataAuthRequestCryptogram, "auth-request-cryptogram", "", "", "Auth Request Cryptogram")
	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataAuthResponseAttributes, "auth-response-attributes", "", "", "Auth Response Attributes")
	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataCipherText, "cipher-text", "", "", "Cipher Text")
	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataDecryptionAttributes, "decryption-attributes", "", "", "Decryption Attributes")
	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataDerivationMethodAttributes, "derivation-method-attributes", "", "", "Derivation Method Attributes")
	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataDukptAttributes, "dukpt-attributes", "", "", "Dukpt Attributes")
	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataEncryptedPinBlock, "encrypted-pin-block", "", "", "Encrypted Pin Block")
	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataEncryptionAttributes, "encryption-attributes", "", "", "Encryption Attributes")
	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataEncryptionKeyIdentifier, "encryption-key-identifier", "", "", "Encryption Key Identifier")
	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataEncryptionWrappedKey, "encryption-wrapped-key", "", "", "Encryption Wrapped Key")
	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataGenerationAttributes, "generation-attributes", "", "", "Generation Attributes")
	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataGenerationKeyIdentifier, "generation-key-identifier", "", "", "Generation Key Identifier")
	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataIncomingAs2805Attributes, "incoming-as2805-attributes", "", "", "Incoming As2805 Attributes")
	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataIncomingDukptAttributes, "incoming-dukpt-attributes", "", "", "Incoming Dukpt Attributes")
	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataIncomingEncryptionAttributes, "incoming-encryption-attributes", "", "", "Incoming Encryption Attributes")
	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataIncomingKeyIdentifier, "incoming-key-identifier", "", "", "Incoming Key Identifier")
	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataIncomingKeyMaterial, "incoming-key-material", "", "", "Incoming Key Material")
	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataIncomingTranslationAttributes, "incoming-translation-attributes", "", "", "Incoming Translation Attributes")
	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataIncomingWrappedKey, "incoming-wrapped-key", "", "", "Incoming Wrapped Key")
	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataKekValidationType, "kek-validation-type", "", "", "Kek Validation Type")
	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataKeyCheckValueAlgorithm, "key-check-value-algorithm", "", "", "Key Check Value Algorithm")
	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataKeyIdentifier, "key-identifier", "", "", "Key Identifier")
	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataMac, "mac", "", "", "Mac")
	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataMacLength, "mac-length", "", "", "Mac Length")
	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataMajorKeyDerivationMode, "major-key-derivation-mode", "", "", "Major Key Derivation Mode")
	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataMessageData, "message-data", "", "", "Message Data")
	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataNewEncryptedPinBlock, "new-encrypted-pin-block", "", "", "New Encrypted Pin Block")
	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataNewPinPekIdentifier, "new-pin-pek-identifier", "", "", "New Pin Pek Identifier")
	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataOutgoingDukptAttributes, "outgoing-dukpt-attributes", "", "", "Outgoing Dukpt Attributes")
	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataOutgoingEncryptionAttributes, "outgoing-encryption-attributes", "", "", "Outgoing Encryption Attributes")
	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataOutgoingKeyIdentifier, "outgoing-key-identifier", "", "", "Outgoing Key Identifier")
	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataOutgoingKeyMaterial, "outgoing-key-material", "", "", "Outgoing Key Material")
	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataOutgoingTranslationAttributes, "outgoing-translation-attributes", "", "", "Outgoing Translation Attributes")
	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataOutgoingWrappedKey, "outgoing-wrapped-key", "", "", "Outgoing Wrapped Key")
	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataPinBlockFormat, "pin-block-format", "", "", "Pin Block Format")
	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataPinDataLength, "pin-data-length", "", "", "Pin Data Length")
	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataPlainText, "plain-text", "", "", "Plain Text")
	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataPrimaryAccountNumber, "primary-account-number", "", "", "Primary Account Number")
	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataRandomKeySendVariantMask, "random-key-send-variant-mask", "", "", "Random Key Send Variant Mask")
	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataSecureMessagingConfidentialityKeyIdentifier, "secure-messaging-confidentiality-key-identifier", "", "", "Secure Messaging Confidentiality Key Identifier")
	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataSecureMessagingIntegrityKeyIdentifier, "secure-messaging-integrity-key-identifier", "", "", "Secure Messaging Integrity Key Identifier")
	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataSessionKeyDerivationAttributes, "session-key-derivation-attributes", "", "", "Session Key Derivation Attributes")
	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataTransactionData, "transaction-data", "", "", "Transaction Data")
	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataValidationData, "validation-data", "", "", "Validation Data")
	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataValidationDataLength, "validation-data-length", "", "", "Validation Data Length")
	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataVerificationAttributes, "verification-attributes", "", "", "Verification Attributes")
	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataVerificationKeyIdentifier, "verification-key-identifier", "", "", "Verification Key Identifier")
	_paymentcryptographydataCmd.Flags().StringVarP(&_paymentcryptographydataWrappedKey, "wrapped-key", "", "", "Wrapped Key")

	_paymentcryptographydataCmd.Flags().BoolVarP(&_paymentcryptographydataDecryptData, "decrypt-data", "", false, "Decrypt Data")
	_paymentcryptographydataCmd.Flags().BoolVarP(&_paymentcryptographydataEncryptData, "encrypt-data", "", false, "Encrypt Data")
	_paymentcryptographydataCmd.Flags().BoolVarP(&_paymentcryptographydataGenerateAs2805KekValidation, "generate-as2805-kek-validation", "", false, "Generate As2805 Kek Validation")
	_paymentcryptographydataCmd.Flags().BoolVarP(&_paymentcryptographydataGenerateCardValidationData, "generate-card-validation-data", "", false, "Generate Card Validation Data")
	_paymentcryptographydataCmd.Flags().BoolVarP(&_paymentcryptographydataGenerateMac, "generate-mac", "", false, "Generate Mac")
	_paymentcryptographydataCmd.Flags().BoolVarP(&_paymentcryptographydataGenerateMacEmvPinChange, "generate-mac-emv-pin-change", "", false, "Generate Mac Emv Pin Change")
	_paymentcryptographydataCmd.Flags().BoolVarP(&_paymentcryptographydataGeneratePinData, "generate-pin-data", "", false, "Generate Pin Data")
	_paymentcryptographydataCmd.Flags().BoolVarP(&_paymentcryptographydataReEncryptData, "re-encrypt-data", "", false, "Re Encrypt Data")
	_paymentcryptographydataCmd.Flags().BoolVarP(&_paymentcryptographydataTranslateKeyMaterial, "translate-key-material", "", false, "Translate Key Material")
	_paymentcryptographydataCmd.Flags().BoolVarP(&_paymentcryptographydataTranslatePinData, "translate-pin-data", "", false, "Translate Pin Data")
	_paymentcryptographydataCmd.Flags().BoolVarP(&_paymentcryptographydataVerifyAuthRequestCryptogram, "verify-auth-request-cryptogram", "", false, "Verify Auth Request Cryptogram")
	_paymentcryptographydataCmd.Flags().BoolVarP(&_paymentcryptographydataVerifyCardValidationData, "verify-card-validation-data", "", false, "Verify Card Validation Data")
	_paymentcryptographydataCmd.Flags().BoolVarP(&_paymentcryptographydataVerifyMac, "verify-mac", "", false, "Verify Mac")
	_paymentcryptographydataCmd.Flags().BoolVarP(&_paymentcryptographydataVerifyPinData, "verify-pin-data", "", false, "Verify Pin Data")

}
