package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/paymentcryptographydata"
)

var fields_decrypt_data = []leanruntime.Field{
	{Name: "CipherText", Flag: "cipher-text", Type: "*string", Required: true},
	{Name: "DecryptionAttributes", Flag: "decryption-attributes", Type: "types.EncryptionDecryptionAttributes", Required: true},
	{Name: "KeyIdentifier", Flag: "key-identifier", Type: "*string", Required: true},
	{Name: "WrappedKey", Flag: "wrapped-key", Type: "*types.WrappedKey", Required: false},
}

var fields_encrypt_data = []leanruntime.Field{
	{Name: "EncryptionAttributes", Flag: "encryption-attributes", Type: "types.EncryptionDecryptionAttributes", Required: true},
	{Name: "KeyIdentifier", Flag: "key-identifier", Type: "*string", Required: true},
	{Name: "PlainText", Flag: "plain-text", Type: "*string", Required: true},
	{Name: "WrappedKey", Flag: "wrapped-key", Type: "*types.WrappedKey", Required: false},
}

var fields_generate_as2805_kek_validation = []leanruntime.Field{
	{Name: "KekValidationType", Flag: "kek-validation-type", Type: "types.As2805KekValidationType", Required: true},
	{Name: "KeyIdentifier", Flag: "key-identifier", Type: "*string", Required: true},
	{Name: "RandomKeySendVariantMask", Flag: "random-key-send-variant-mask", Type: "types.RandomKeySendVariantMask", Required: true},
}

var fields_generate_card_validation_data = []leanruntime.Field{
	{Name: "GenerationAttributes", Flag: "generation-attributes", Type: "types.CardGenerationAttributes", Required: true},
	{Name: "KeyIdentifier", Flag: "key-identifier", Type: "*string", Required: true},
	{Name: "PrimaryAccountNumber", Flag: "primary-account-number", Type: "*string", Required: true},
	{Name: "ValidationDataLength", Flag: "validation-data-length", Type: "*int32", Required: false},
}

var fields_generate_mac = []leanruntime.Field{
	{Name: "GenerationAttributes", Flag: "generation-attributes", Type: "types.MacAttributes", Required: true},
	{Name: "KeyIdentifier", Flag: "key-identifier", Type: "*string", Required: true},
	{Name: "MacLength", Flag: "mac-length", Type: "*int32", Required: false},
	{Name: "MessageData", Flag: "message-data", Type: "*string", Required: true},
}

var fields_generate_mac_emv_pin_change = []leanruntime.Field{
	{Name: "DerivationMethodAttributes", Flag: "derivation-method-attributes", Type: "types.DerivationMethodAttributes", Required: true},
	{Name: "MessageData", Flag: "message-data", Type: "*string", Required: true},
	{Name: "NewEncryptedPinBlock", Flag: "new-encrypted-pin-block", Type: "*string", Required: true},
	{Name: "NewPinPekIdentifier", Flag: "new-pin-pek-identifier", Type: "*string", Required: true},
	{Name: "PinBlockFormat", Flag: "pin-block-format", Type: "types.PinBlockFormatForEmvPinChange", Required: true},
	{Name: "SecureMessagingConfidentialityKeyIdentifier", Flag: "secure-messaging-confidentiality-key-identifier", Type: "*string", Required: true},
	{Name: "SecureMessagingIntegrityKeyIdentifier", Flag: "secure-messaging-integrity-key-identifier", Type: "*string", Required: true},
}

var fields_generate_pin_data = []leanruntime.Field{
	{Name: "EncryptionKeyIdentifier", Flag: "encryption-key-identifier", Type: "*string", Required: true},
	{Name: "EncryptionWrappedKey", Flag: "encryption-wrapped-key", Type: "*types.WrappedKey", Required: false},
	{Name: "GenerationAttributes", Flag: "generation-attributes", Type: "types.PinGenerationAttributes", Required: true},
	{Name: "GenerationKeyIdentifier", Flag: "generation-key-identifier", Type: "*string", Required: true},
	{Name: "PinBlockFormat", Flag: "pin-block-format", Type: "types.PinBlockFormatForPinData", Required: true},
	{Name: "PinDataLength", Flag: "pin-data-length", Type: "*int32", Required: false},
	{Name: "PrimaryAccountNumber", Flag: "primary-account-number", Type: "*string", Required: false},
}

var fields_re_encrypt_data = []leanruntime.Field{
	{Name: "CipherText", Flag: "cipher-text", Type: "*string", Required: true},
	{Name: "IncomingEncryptionAttributes", Flag: "incoming-encryption-attributes", Type: "types.ReEncryptionAttributes", Required: true},
	{Name: "IncomingKeyIdentifier", Flag: "incoming-key-identifier", Type: "*string", Required: true},
	{Name: "IncomingWrappedKey", Flag: "incoming-wrapped-key", Type: "*types.WrappedKey", Required: false},
	{Name: "OutgoingEncryptionAttributes", Flag: "outgoing-encryption-attributes", Type: "types.ReEncryptionAttributes", Required: true},
	{Name: "OutgoingKeyIdentifier", Flag: "outgoing-key-identifier", Type: "*string", Required: true},
	{Name: "OutgoingWrappedKey", Flag: "outgoing-wrapped-key", Type: "*types.WrappedKey", Required: false},
}

var fields_translate_key_material = []leanruntime.Field{
	{Name: "IncomingKeyMaterial", Flag: "incoming-key-material", Type: "types.IncomingKeyMaterial", Required: true},
	{Name: "KeyCheckValueAlgorithm", Flag: "key-check-value-algorithm", Type: "types.KeyCheckValueAlgorithm", Required: false},
	{Name: "OutgoingKeyMaterial", Flag: "outgoing-key-material", Type: "types.OutgoingKeyMaterial", Required: true},
}

var fields_translate_pin_data = []leanruntime.Field{
	{Name: "EncryptedPinBlock", Flag: "encrypted-pin-block", Type: "*string", Required: true},
	{Name: "IncomingAs2805Attributes", Flag: "incoming-as2805-attributes", Type: "*types.As2805PekDerivationAttributes", Required: false},
	{Name: "IncomingDukptAttributes", Flag: "incoming-dukpt-attributes", Type: "*types.DukptDerivationAttributes", Required: false},
	{Name: "IncomingKeyIdentifier", Flag: "incoming-key-identifier", Type: "*string", Required: true},
	{Name: "IncomingTranslationAttributes", Flag: "incoming-translation-attributes", Type: "types.TranslationIsoFormats", Required: true},
	{Name: "IncomingWrappedKey", Flag: "incoming-wrapped-key", Type: "*types.WrappedKey", Required: false},
	{Name: "OutgoingDukptAttributes", Flag: "outgoing-dukpt-attributes", Type: "*types.DukptDerivationAttributes", Required: false},
	{Name: "OutgoingKeyIdentifier", Flag: "outgoing-key-identifier", Type: "*string", Required: true},
	{Name: "OutgoingTranslationAttributes", Flag: "outgoing-translation-attributes", Type: "types.TranslationIsoFormats", Required: true},
	{Name: "OutgoingWrappedKey", Flag: "outgoing-wrapped-key", Type: "*types.WrappedKey", Required: false},
}

var fields_verify_auth_request_cryptogram = []leanruntime.Field{
	{Name: "AuthRequestCryptogram", Flag: "auth-request-cryptogram", Type: "*string", Required: true},
	{Name: "AuthResponseAttributes", Flag: "auth-response-attributes", Type: "types.CryptogramAuthResponse", Required: false},
	{Name: "KeyIdentifier", Flag: "key-identifier", Type: "*string", Required: true},
	{Name: "MajorKeyDerivationMode", Flag: "major-key-derivation-mode", Type: "types.MajorKeyDerivationMode", Required: true},
	{Name: "SessionKeyDerivationAttributes", Flag: "session-key-derivation-attributes", Type: "types.SessionKeyDerivation", Required: true},
	{Name: "TransactionData", Flag: "transaction-data", Type: "*string", Required: true},
}

var fields_verify_card_validation_data = []leanruntime.Field{
	{Name: "KeyIdentifier", Flag: "key-identifier", Type: "*string", Required: true},
	{Name: "PrimaryAccountNumber", Flag: "primary-account-number", Type: "*string", Required: true},
	{Name: "ValidationData", Flag: "validation-data", Type: "*string", Required: true},
	{Name: "VerificationAttributes", Flag: "verification-attributes", Type: "types.CardVerificationAttributes", Required: true},
}

var fields_verify_mac = []leanruntime.Field{
	{Name: "KeyIdentifier", Flag: "key-identifier", Type: "*string", Required: true},
	{Name: "Mac", Flag: "mac", Type: "*string", Required: true},
	{Name: "MacLength", Flag: "mac-length", Type: "*int32", Required: false},
	{Name: "MessageData", Flag: "message-data", Type: "*string", Required: true},
	{Name: "VerificationAttributes", Flag: "verification-attributes", Type: "types.MacAttributes", Required: true},
}

var fields_verify_pin_data = []leanruntime.Field{
	{Name: "DukptAttributes", Flag: "dukpt-attributes", Type: "*types.DukptAttributes", Required: false},
	{Name: "EncryptedPinBlock", Flag: "encrypted-pin-block", Type: "*string", Required: true},
	{Name: "EncryptionKeyIdentifier", Flag: "encryption-key-identifier", Type: "*string", Required: true},
	{Name: "EncryptionWrappedKey", Flag: "encryption-wrapped-key", Type: "*types.WrappedKey", Required: false},
	{Name: "PinBlockFormat", Flag: "pin-block-format", Type: "types.PinBlockFormatForPinData", Required: true},
	{Name: "PinDataLength", Flag: "pin-data-length", Type: "*int32", Required: false},
	{Name: "PrimaryAccountNumber", Flag: "primary-account-number", Type: "*string", Required: false},
	{Name: "VerificationAttributes", Flag: "verification-attributes", Type: "types.PinVerificationAttributes", Required: true},
	{Name: "VerificationKeyIdentifier", Flag: "verification-key-identifier", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"decrypt-data": {
			Name:   "decrypt-data",
			Fields: fields_decrypt_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DecryptDataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_decrypt_data, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DecryptData(ctx, input)
			},
		},
		"encrypt-data": {
			Name:   "encrypt-data",
			Fields: fields_encrypt_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EncryptDataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_encrypt_data, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EncryptData(ctx, input)
			},
		},
		"generate-as2805-kek-validation": {
			Name:   "generate-as2805-kek-validation",
			Fields: fields_generate_as2805_kek_validation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GenerateAs2805KekValidationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_generate_as2805_kek_validation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GenerateAs2805KekValidation(ctx, input)
			},
		},
		"generate-card-validation-data": {
			Name:   "generate-card-validation-data",
			Fields: fields_generate_card_validation_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GenerateCardValidationDataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_generate_card_validation_data, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GenerateCardValidationData(ctx, input)
			},
		},
		"generate-mac": {
			Name:   "generate-mac",
			Fields: fields_generate_mac,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GenerateMacInput{}
				if _, err := leanruntime.ApplyInput(input, fields_generate_mac, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GenerateMac(ctx, input)
			},
		},
		"generate-mac-emv-pin-change": {
			Name:   "generate-mac-emv-pin-change",
			Fields: fields_generate_mac_emv_pin_change,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GenerateMacEmvPinChangeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_generate_mac_emv_pin_change, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GenerateMacEmvPinChange(ctx, input)
			},
		},
		"generate-pin-data": {
			Name:   "generate-pin-data",
			Fields: fields_generate_pin_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GeneratePinDataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_generate_pin_data, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GeneratePinData(ctx, input)
			},
		},
		"re-encrypt-data": {
			Name:   "re-encrypt-data",
			Fields: fields_re_encrypt_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ReEncryptDataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_re_encrypt_data, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ReEncryptData(ctx, input)
			},
		},
		"translate-key-material": {
			Name:   "translate-key-material",
			Fields: fields_translate_key_material,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TranslateKeyMaterialInput{}
				if _, err := leanruntime.ApplyInput(input, fields_translate_key_material, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TranslateKeyMaterial(ctx, input)
			},
		},
		"translate-pin-data": {
			Name:   "translate-pin-data",
			Fields: fields_translate_pin_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TranslatePinDataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_translate_pin_data, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TranslatePinData(ctx, input)
			},
		},
		"verify-auth-request-cryptogram": {
			Name:   "verify-auth-request-cryptogram",
			Fields: fields_verify_auth_request_cryptogram,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.VerifyAuthRequestCryptogramInput{}
				if _, err := leanruntime.ApplyInput(input, fields_verify_auth_request_cryptogram, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.VerifyAuthRequestCryptogram(ctx, input)
			},
		},
		"verify-card-validation-data": {
			Name:   "verify-card-validation-data",
			Fields: fields_verify_card_validation_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.VerifyCardValidationDataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_verify_card_validation_data, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.VerifyCardValidationData(ctx, input)
			},
		},
		"verify-mac": {
			Name:   "verify-mac",
			Fields: fields_verify_mac,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.VerifyMacInput{}
				if _, err := leanruntime.ApplyInput(input, fields_verify_mac, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.VerifyMac(ctx, input)
			},
		},
		"verify-pin-data": {
			Name:   "verify-pin-data",
			Fields: fields_verify_pin_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.VerifyPinDataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_verify_pin_data, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.VerifyPinData(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("paymentcryptographydata", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
