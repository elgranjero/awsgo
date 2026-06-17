package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/paymentcryptographydata/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"decrypt-data", "encrypt-data", "generate-as2805-kek-validation", "generate-card-validation-data", "generate-mac", "generate-mac-emv-pin-change", "generate-pin-data", "re-encrypt-data", "translate-key-material", "translate-pin-data", "verify-auth-request-cryptogram", "verify-card-validation-data", "verify-mac", "verify-pin-data"},
		OperationSet: map[string]bool{"decrypt-data": true, "encrypt-data": true, "generate-as2805-kek-validation": true, "generate-card-validation-data": true, "generate-mac": true, "generate-mac-emv-pin-change": true, "generate-pin-data": true, "re-encrypt-data": true, "translate-key-material": true, "translate-pin-data": true, "verify-auth-request-cryptogram": true, "verify-card-validation-data": true, "verify-mac": true, "verify-pin-data": true},
		OperationInputs: map[string][]string{
			"decrypt-data":                   {"CipherText", "DecryptionAttributes", "KeyIdentifier", "WrappedKey"},
			"encrypt-data":                   {"EncryptionAttributes", "KeyIdentifier", "PlainText", "WrappedKey"},
			"generate-as2805-kek-validation": {"KekValidationType", "KeyIdentifier", "RandomKeySendVariantMask"},
			"generate-card-validation-data":  {"GenerationAttributes", "KeyIdentifier", "PrimaryAccountNumber", "ValidationDataLength"},
			"generate-mac":                   {"GenerationAttributes", "KeyIdentifier", "MacLength", "MessageData"},
			"generate-mac-emv-pin-change":    {"DerivationMethodAttributes", "MessageData", "NewEncryptedPinBlock", "NewPinPekIdentifier", "PinBlockFormat", "SecureMessagingConfidentialityKeyIdentifier", "SecureMessagingIntegrityKeyIdentifier"},
			"generate-pin-data":              {"EncryptionKeyIdentifier", "EncryptionWrappedKey", "GenerationAttributes", "GenerationKeyIdentifier", "PinBlockFormat", "PinDataLength", "PrimaryAccountNumber"},
			"re-encrypt-data":                {"CipherText", "IncomingEncryptionAttributes", "IncomingKeyIdentifier", "IncomingWrappedKey", "OutgoingEncryptionAttributes", "OutgoingKeyIdentifier", "OutgoingWrappedKey"},
			"translate-key-material":         {"IncomingKeyMaterial", "KeyCheckValueAlgorithm", "OutgoingKeyMaterial"},
			"translate-pin-data":             {"EncryptedPinBlock", "IncomingAs2805Attributes", "IncomingDukptAttributes", "IncomingKeyIdentifier", "IncomingTranslationAttributes", "IncomingWrappedKey", "OutgoingDukptAttributes", "OutgoingKeyIdentifier", "OutgoingTranslationAttributes", "OutgoingWrappedKey"},
			"verify-auth-request-cryptogram": {"AuthRequestCryptogram", "AuthResponseAttributes", "KeyIdentifier", "MajorKeyDerivationMode", "SessionKeyDerivationAttributes", "TransactionData"},
			"verify-card-validation-data":    {"KeyIdentifier", "PrimaryAccountNumber", "ValidationData", "VerificationAttributes"},
			"verify-mac":                     {"KeyIdentifier", "Mac", "MacLength", "MessageData", "VerificationAttributes"},
			"verify-pin-data":                {"DukptAttributes", "EncryptedPinBlock", "EncryptionKeyIdentifier", "EncryptionWrappedKey", "PinBlockFormat", "PinDataLength", "PrimaryAccountNumber", "VerificationAttributes", "VerificationKeyIdentifier"},
		},
		OperationInputTypes: map[string]map[string]string{
			"decrypt-data":                   {"CipherText": "*string", "DecryptionAttributes": "types.EncryptionDecryptionAttributes", "KeyIdentifier": "*string", "WrappedKey": "*types.WrappedKey"},
			"encrypt-data":                   {"EncryptionAttributes": "types.EncryptionDecryptionAttributes", "KeyIdentifier": "*string", "PlainText": "*string", "WrappedKey": "*types.WrappedKey"},
			"generate-as2805-kek-validation": {"KekValidationType": "types.As2805KekValidationType", "KeyIdentifier": "*string", "RandomKeySendVariantMask": "types.RandomKeySendVariantMask"},
			"generate-card-validation-data":  {"GenerationAttributes": "types.CardGenerationAttributes", "KeyIdentifier": "*string", "PrimaryAccountNumber": "*string", "ValidationDataLength": "*int32"},
			"generate-mac":                   {"GenerationAttributes": "types.MacAttributes", "KeyIdentifier": "*string", "MacLength": "*int32", "MessageData": "*string"},
			"generate-mac-emv-pin-change":    {"DerivationMethodAttributes": "types.DerivationMethodAttributes", "MessageData": "*string", "NewEncryptedPinBlock": "*string", "NewPinPekIdentifier": "*string", "PinBlockFormat": "types.PinBlockFormatForEmvPinChange", "SecureMessagingConfidentialityKeyIdentifier": "*string", "SecureMessagingIntegrityKeyIdentifier": "*string"},
			"generate-pin-data":              {"EncryptionKeyIdentifier": "*string", "EncryptionWrappedKey": "*types.WrappedKey", "GenerationAttributes": "types.PinGenerationAttributes", "GenerationKeyIdentifier": "*string", "PinBlockFormat": "types.PinBlockFormatForPinData", "PinDataLength": "*int32", "PrimaryAccountNumber": "*string"},
			"re-encrypt-data":                {"CipherText": "*string", "IncomingEncryptionAttributes": "types.ReEncryptionAttributes", "IncomingKeyIdentifier": "*string", "IncomingWrappedKey": "*types.WrappedKey", "OutgoingEncryptionAttributes": "types.ReEncryptionAttributes", "OutgoingKeyIdentifier": "*string", "OutgoingWrappedKey": "*types.WrappedKey"},
			"translate-key-material":         {"IncomingKeyMaterial": "types.IncomingKeyMaterial", "KeyCheckValueAlgorithm": "types.KeyCheckValueAlgorithm", "OutgoingKeyMaterial": "types.OutgoingKeyMaterial"},
			"translate-pin-data":             {"EncryptedPinBlock": "*string", "IncomingAs2805Attributes": "*types.As2805PekDerivationAttributes", "IncomingDukptAttributes": "*types.DukptDerivationAttributes", "IncomingKeyIdentifier": "*string", "IncomingTranslationAttributes": "types.TranslationIsoFormats", "IncomingWrappedKey": "*types.WrappedKey", "OutgoingDukptAttributes": "*types.DukptDerivationAttributes", "OutgoingKeyIdentifier": "*string", "OutgoingTranslationAttributes": "types.TranslationIsoFormats", "OutgoingWrappedKey": "*types.WrappedKey"},
			"verify-auth-request-cryptogram": {"AuthRequestCryptogram": "*string", "AuthResponseAttributes": "types.CryptogramAuthResponse", "KeyIdentifier": "*string", "MajorKeyDerivationMode": "types.MajorKeyDerivationMode", "SessionKeyDerivationAttributes": "types.SessionKeyDerivation", "TransactionData": "*string"},
			"verify-card-validation-data":    {"KeyIdentifier": "*string", "PrimaryAccountNumber": "*string", "ValidationData": "*string", "VerificationAttributes": "types.CardVerificationAttributes"},
			"verify-mac":                     {"KeyIdentifier": "*string", "Mac": "*string", "MacLength": "*int32", "MessageData": "*string", "VerificationAttributes": "types.MacAttributes"},
			"verify-pin-data":                {"DukptAttributes": "*types.DukptAttributes", "EncryptedPinBlock": "*string", "EncryptionKeyIdentifier": "*string", "EncryptionWrappedKey": "*types.WrappedKey", "PinBlockFormat": "types.PinBlockFormatForPinData", "PinDataLength": "*int32", "PrimaryAccountNumber": "*string", "VerificationAttributes": "types.PinVerificationAttributes", "VerificationKeyIdentifier": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"decrypt-data":                   {"CipherText", "DecryptionAttributes", "KeyIdentifier"},
			"encrypt-data":                   {"EncryptionAttributes", "KeyIdentifier", "PlainText"},
			"generate-as2805-kek-validation": {"KekValidationType", "KeyIdentifier", "RandomKeySendVariantMask"},
			"generate-card-validation-data":  {"GenerationAttributes", "KeyIdentifier", "PrimaryAccountNumber"},
			"generate-mac":                   {"GenerationAttributes", "KeyIdentifier", "MessageData"},
			"generate-mac-emv-pin-change":    {"DerivationMethodAttributes", "MessageData", "NewEncryptedPinBlock", "NewPinPekIdentifier", "PinBlockFormat", "SecureMessagingConfidentialityKeyIdentifier", "SecureMessagingIntegrityKeyIdentifier"},
			"generate-pin-data":              {"EncryptionKeyIdentifier", "GenerationAttributes", "GenerationKeyIdentifier", "PinBlockFormat"},
			"re-encrypt-data":                {"CipherText", "IncomingEncryptionAttributes", "IncomingKeyIdentifier", "OutgoingEncryptionAttributes", "OutgoingKeyIdentifier"},
			"translate-key-material":         {"IncomingKeyMaterial", "OutgoingKeyMaterial"},
			"translate-pin-data":             {"EncryptedPinBlock", "IncomingKeyIdentifier", "IncomingTranslationAttributes", "OutgoingKeyIdentifier", "OutgoingTranslationAttributes"},
			"verify-auth-request-cryptogram": {"AuthRequestCryptogram", "KeyIdentifier", "MajorKeyDerivationMode", "SessionKeyDerivationAttributes", "TransactionData"},
			"verify-card-validation-data":    {"KeyIdentifier", "PrimaryAccountNumber", "ValidationData", "VerificationAttributes"},
			"verify-mac":                     {"KeyIdentifier", "Mac", "MessageData", "VerificationAttributes"},
			"verify-pin-data":                {"EncryptedPinBlock", "EncryptionKeyIdentifier", "PinBlockFormat", "VerificationAttributes", "VerificationKeyIdentifier"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("paymentcryptographydata", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
