package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/paymentcryptography"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// paymentcryptographyCmd represents the paymentcryptography command
var _paymentcryptographyCmd = &cobra.Command{
	Use:   "paymentcryptography",
	Short: "AWS paymentcryptography CLI",
	Run: func(cmd *cobra.Command, args []string) {
		_awsOutput = resolveAWSOutput(_awsProfile, cmd.Flags().Changed("output"))
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := paymentcryptography.NewFromConfig(cfg)
		if _paymentcryptographyAddKeyReplicationRegions {
			paymentcryptography_AddKeyReplicationRegions(cfg, client)
			return
		}
		if _paymentcryptographyCreateAlias {
			paymentcryptography_CreateAlias(cfg, client)
			return
		}
		if _paymentcryptographyCreateKey {
			paymentcryptography_CreateKey(cfg, client)
			return
		}
		if _paymentcryptographyDeleteAlias {
			paymentcryptography_DeleteAlias(cfg, client)
			return
		}
		if _paymentcryptographyDeleteKey {
			paymentcryptography_DeleteKey(cfg, client)
			return
		}
		if _paymentcryptographyDisableDefaultKeyReplicationRegions {
			paymentcryptography_DisableDefaultKeyReplicationRegions(cfg, client)
			return
		}
		if _paymentcryptographyEnableDefaultKeyReplicationRegions {
			paymentcryptography_EnableDefaultKeyReplicationRegions(cfg, client)
			return
		}
		if _paymentcryptographyExportKey {
			paymentcryptography_ExportKey(cfg, client)
			return
		}
		if _paymentcryptographyGetAlias {
			paymentcryptography_GetAlias(cfg, client)
			return
		}
		if _paymentcryptographyGetCertificateSigningRequest {
			paymentcryptography_GetCertificateSigningRequest(cfg, client)
			return
		}
		if _paymentcryptographyGetDefaultKeyReplicationRegions {
			paymentcryptography_GetDefaultKeyReplicationRegions(cfg, client)
			return
		}
		if _paymentcryptographyGetKey {
			paymentcryptography_GetKey(cfg, client)
			return
		}
		if _paymentcryptographyGetParametersForExport {
			paymentcryptography_GetParametersForExport(cfg, client)
			return
		}
		if _paymentcryptographyGetParametersForImport {
			paymentcryptography_GetParametersForImport(cfg, client)
			return
		}
		if _paymentcryptographyGetPublicKeyCertificate {
			paymentcryptography_GetPublicKeyCertificate(cfg, client)
			return
		}
		if _paymentcryptographyImportKey {
			paymentcryptography_ImportKey(cfg, client)
			return
		}
		if _paymentcryptographyListAliases {
			paymentcryptography_ListAliases(cfg, client)
			return
		}
		if _paymentcryptographyListKeys {
			paymentcryptography_ListKeys(cfg, client)
			return
		}
		if _paymentcryptographyListTagsForResource {
			paymentcryptography_ListTagsForResource(cfg, client)
			return
		}
		if _paymentcryptographyRemoveKeyReplicationRegions {
			paymentcryptography_RemoveKeyReplicationRegions(cfg, client)
			return
		}
		if _paymentcryptographyRestoreKey {
			paymentcryptography_RestoreKey(cfg, client)
			return
		}
		if _paymentcryptographyStartKeyUsage {
			paymentcryptography_StartKeyUsage(cfg, client)
			return
		}
		if _paymentcryptographyStopKeyUsage {
			paymentcryptography_StopKeyUsage(cfg, client)
			return
		}
		if _paymentcryptographyTagResource {
			paymentcryptography_TagResource(cfg, client)
			return
		}
		if _paymentcryptographyUntagResource {
			paymentcryptography_UntagResource(cfg, client)
			return
		}
		if _paymentcryptographyUpdateAlias {
			paymentcryptography_UpdateAlias(cfg, client)
			return
		}

	},
}

var (
	_paymentcryptographyAddKeyReplicationRegions            bool
	_paymentcryptographyCreateAlias                         bool
	_paymentcryptographyCreateKey                           bool
	_paymentcryptographyDeleteAlias                         bool
	_paymentcryptographyDeleteKey                           bool
	_paymentcryptographyDisableDefaultKeyReplicationRegions bool
	_paymentcryptographyEnableDefaultKeyReplicationRegions  bool
	_paymentcryptographyExportKey                           bool
	_paymentcryptographyGetAlias                            bool
	_paymentcryptographyGetCertificateSigningRequest        bool
	_paymentcryptographyGetDefaultKeyReplicationRegions     bool
	_paymentcryptographyGetKey                              bool
	_paymentcryptographyGetParametersForExport              bool
	_paymentcryptographyGetParametersForImport              bool
	_paymentcryptographyGetPublicKeyCertificate             bool
	_paymentcryptographyImportKey                           bool
	_paymentcryptographyListAliases                         bool
	_paymentcryptographyListKeys                            bool
	_paymentcryptographyListTagsForResource                 bool
	_paymentcryptographyRemoveKeyReplicationRegions         bool
	_paymentcryptographyRestoreKey                          bool
	_paymentcryptographyStartKeyUsage                       bool
	_paymentcryptographyStopKeyUsage                        bool
	_paymentcryptographyTagResource                         bool
	_paymentcryptographyUntagResource                       bool
	_paymentcryptographyUpdateAlias                         bool

	_paymentcryptographyAliasName              string
	_paymentcryptographyCertificateSubject     string
	_paymentcryptographyDeleteKeyInDays        string
	_paymentcryptographyDeriveKeyUsage         string
	_paymentcryptographyEnabled                string
	_paymentcryptographyExportAttributes       string
	_paymentcryptographyExportKeyIdentifier    string
	_paymentcryptographyExportable             string
	_paymentcryptographyKeyArn                 string
	_paymentcryptographyKeyAttributes          string
	_paymentcryptographyKeyCheckValueAlgorithm string
	_paymentcryptographyKeyIdentifier          string
	_paymentcryptographyKeyMaterial            string
	_paymentcryptographyKeyMaterialType        string
	_paymentcryptographyKeyState               string
	_paymentcryptographyMaxResults             string
	_paymentcryptographyNextToken              string
	_paymentcryptographyReplicationRegions     []string
	_paymentcryptographyResourceArn            string
	_paymentcryptographySigningAlgorithm       string
	_paymentcryptographySigningKeyAlgorithm    string
	_paymentcryptographyTagKeys                []string
	_paymentcryptographyTags                   string
	_paymentcryptographyWrappingKeyAlgorithm   string
)

// Adds replication Amazon Web Services Regions to an existing Amazon Web Services
// Payment Cryptography key, enabling the key to be used for cryptographic
// operations in additional Amazon Web Services Regions.
//
// [Multi-Region key replication]allow you to use the same key material across multiple Amazon Web Services
// Regions, providing lower latency for applications distributed across regions.
// When you add Replication Regions, Amazon Web Services Payment Cryptography
// securely replicates the key material to the specified Amazon Web Services
// Regions.
//
// The key must be in an active state to add Replication Regions. You can add
// multiple regions in a single operation, and the key will be available for use in
// those regions once replication is complete.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// [RemoveKeyReplicationRegions]
//
// [EnableDefaultKeyReplicationRegions]
//
// [GetDefaultKeyReplicationRegions]
//
// [RemoveKeyReplicationRegions]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_RemoveKeyReplicationRegions.html
// [Multi-Region key replication]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-multi-region-replication.html
// [EnableDefaultKeyReplicationRegions]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_EnableDefaultKeyReplicationRegions.html
// [GetDefaultKeyReplicationRegions]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_GetDefaultKeyReplicationRegions.html
func paymentcryptography_AddKeyReplicationRegions(cfg aws.Config, client *paymentcryptography.Client) {
	input := &paymentcryptography.AddKeyReplicationRegionsInput{
		// KeyIdentifier: *string, // Required
		// ReplicationRegions: []string, // Required
	}

	if len(_paymentcryptographyKeyIdentifier) > 0 {
		input.KeyIdentifier = aws.String(_paymentcryptographyKeyIdentifier)
	}
	if len(_paymentcryptographyReplicationRegions) > 0 {
		input.ReplicationRegions = append([]string(nil), _paymentcryptographyReplicationRegions...)
	}

	if resp, err := client.AddKeyReplicationRegions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an alias, or a friendly name, for an Amazon Web Services Payment
// Cryptography key. You can use an alias to identify a key in the console and when
// you call cryptographic operations such as [EncryptData]or [DecryptData].
//
// You can associate the alias with any key in the same Amazon Web Services
// Region. Each alias is associated with only one key at a time, but a key can have
// multiple aliases. You can't create an alias without a key. The alias must be
// unique in the account and Amazon Web Services Region, but you can create another
// alias with the same name in a different Amazon Web Services Region.
//
// To change the key that's associated with the alias, call [UpdateAlias]. To delete the alias,
// call [DeleteAlias]. These operations don't affect the underlying key. To get the alias that
// you created, call [ListAliases].
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// [DeleteAlias]
//
// [GetAlias]
//
// [ListAliases]
//
// [UpdateAlias]
//
// [ListAliases]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_ListAliases.html
// [DeleteAlias]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_DeleteAlias.html
// [UpdateAlias]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_UpdateAlias.html
// [GetAlias]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_GetAlias.html
// [EncryptData]: https://docs.aws.amazon.com/payment-cryptography/latest/DataAPIReference/API_EncryptData.html
// [DecryptData]: https://docs.aws.amazon.com/payment-cryptography/latest/DataAPIReference/API_DecryptData.html
func paymentcryptography_CreateAlias(cfg aws.Config, client *paymentcryptography.Client) {
	input := &paymentcryptography.CreateAliasInput{
		// AliasName: *string, // Required
	}

	if len(_paymentcryptographyAliasName) > 0 {
		input.AliasName = aws.String(_paymentcryptographyAliasName)
	}
	if len(_paymentcryptographyKeyArn) > 0 {
		input.KeyArn = aws.String(_paymentcryptographyKeyArn)
	}

	if resp, err := client.CreateAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Web Services Payment Cryptography key, a logical
// representation of a cryptographic key, that is unique in your account and Amazon
// Web Services Region. You use keys for cryptographic functions such as encryption
// and decryption.
//
// In addition to the key material used in cryptographic operations, an Amazon Web
// Services Payment Cryptography key includes metadata such as the key ARN, key
// usage, key origin, creation date, description, and key state.
//
// When you create a key, you specify both immutable and mutable data about the
// key. The immutable data contains key attributes that define the scope and
// cryptographic operations that you can perform using the key, for example key
// class (example: SYMMETRIC_KEY ), key algorithm (example: TDES_2KEY ), key usage
// (example: TR31_P0_PIN_ENCRYPTION_KEY ) and key modes of use (example: Encrypt ).
// Amazon Web Services Payment Cryptography binds key attributes to keys using key
// blocks when you store or export them. Amazon Web Services Payment Cryptography
// stores the key contents wrapped and never stores or transmits them in the clear.
//
// For information about valid combinations of key attributes, see [Understanding key attributes] in the Amazon
// Web Services Payment Cryptography User Guide. The mutable data contained within
// a key includes usage timestamp and key deletion timestamp and can be modified
// after creation.
//
// You can use the CreateKey operation to generate an ECC (Elliptic Curve
// Cryptography) key pair used for establishing an ECDH (Elliptic Curve
// Diffie-Hellman) key agreement between two parties. In the ECDH key agreement
// process, both parties generate their own ECC key pair with key usage K3 and
// exchange the public keys. Each party then use their private key, the received
// public key from the other party, and the key derivation parameters including key
// derivation function, hash algorithm, derivation data, and key algorithm to
// derive a shared key.
//
// To maintain the single-use principle of cryptographic keys in payments, ECDH
// derived keys should not be used for multiple purposes, such as a
// TR31_P0_PIN_ENCRYPTION_KEY and TR31_K1_KEY_BLOCK_PROTECTION_KEY . When creating
// ECC key pairs in Amazon Web Services Payment Cryptography you can optionally set
// the DeriveKeyUsage parameter, which defines the key usage bound to the
// symmetric key that will be derived using the ECC key pair.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// [DeleteKey]
//
// [GetKey]
//
// [ListKeys]
//
// [DeleteKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_DeleteKey.html
// [GetKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_GetKey.html
// [ListKeys]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_ListKeys.html
// [Understanding key attributes]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-validattributes.html
func paymentcryptography_CreateKey(cfg aws.Config, client *paymentcryptography.Client) {
	input := &paymentcryptography.CreateKeyInput{
		// Exportable: *bool, // Required
		// KeyAttributes: *types.KeyAttributes, // Required
	}

	if len(_paymentcryptographyExportable) > 0 {
		if err := assignInputField(input, "Exportable", _paymentcryptographyExportable); err != nil {
			log.Errorf("invalid --exportable: %s", err.Error())
			return
		}
	}
	if len(_paymentcryptographyKeyAttributes) > 0 {
		if err := assignInputField(input, "KeyAttributes", _paymentcryptographyKeyAttributes); err != nil {
			log.Errorf("invalid --key-attributes: %s", err.Error())
			return
		}
	}
	if len(_paymentcryptographyDeriveKeyUsage) > 0 {
		if err := assignInputField(input, "DeriveKeyUsage", _paymentcryptographyDeriveKeyUsage); err != nil {
			log.Errorf("invalid --derive-key-usage: %s", err.Error())
			return
		}
	}
	if len(_paymentcryptographyEnabled) > 0 {
		if err := assignInputField(input, "Enabled", _paymentcryptographyEnabled); err != nil {
			log.Errorf("invalid --enabled: %s", err.Error())
			return
		}
	}
	if len(_paymentcryptographyKeyCheckValueAlgorithm) > 0 {
		if err := assignInputField(input, "KeyCheckValueAlgorithm", _paymentcryptographyKeyCheckValueAlgorithm); err != nil {
			log.Errorf("invalid --key-check-value-algorithm: %s", err.Error())
			return
		}
	}
	if len(_paymentcryptographyReplicationRegions) > 0 {
		input.ReplicationRegions = append([]string(nil), _paymentcryptographyReplicationRegions...)
	}
	if len(_paymentcryptographyTags) > 0 {
		if err := assignInputField(input, "Tags", _paymentcryptographyTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the alias, but doesn't affect the underlying key.
// Each key can have multiple aliases. To get the aliases of all keys, use the [UpdateAlias]
// operation. To change the alias of a key, first use [DeleteAlias]to delete the current alias
// and then use [CreateAlias]to create a new alias. To associate an existing alias with a
// different key, call [UpdateAlias].
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// [CreateAlias]
//
// [GetAlias]
//
// [ListAliases]
//
// [UpdateAlias]
//
// [ListAliases]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_ListAliases.html
// [DeleteAlias]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_DeleteAlias.html
// [UpdateAlias]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_UpdateAlias.html
// [CreateAlias]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_CreateAlias.html
// [GetAlias]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_GetAlias.html
func paymentcryptography_DeleteAlias(cfg aws.Config, client *paymentcryptography.Client) {
	input := &paymentcryptography.DeleteAliasInput{
		// AliasName: *string, // Required
	}

	if len(_paymentcryptographyAliasName) > 0 {
		input.AliasName = aws.String(_paymentcryptographyAliasName)
	}

	if resp, err := client.DeleteAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the key material and metadata associated with Amazon Web Services
// Payment Cryptography key.
//
// Key deletion is irreversible. After a key is deleted, you can't perform
// cryptographic operations using the key. For example, you can't decrypt data that
// was encrypted by a deleted Amazon Web Services Payment Cryptography key, and the
// data may become unrecoverable. Because key deletion is destructive, Amazon Web
// Services Payment Cryptography has a safety mechanism to prevent accidental
// deletion of a key. When you call this operation, Amazon Web Services Payment
// Cryptography disables the specified key but doesn't delete it until after a
// waiting period set using DeleteKeyInDays . The default waiting period is 7 days.
// During the waiting period, the KeyState is DELETE_PENDING . After the key is
// deleted, the KeyState is DELETE_COMPLETE .
//
// You should delete a key only when you are sure that you don't need to use it
// anymore and no other parties are utilizing this key. If you aren't sure,
// consider deactivating it instead by calling [StopKeyUsage].
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// [RestoreKey]
//
// [StartKeyUsage]
//
// [StopKeyUsage]
//
// [StartKeyUsage]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_StartKeyUsage.html
// [StopKeyUsage]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_StopKeyUsage.html
// [RestoreKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_RestoreKey.html
func paymentcryptography_DeleteKey(cfg aws.Config, client *paymentcryptography.Client) {
	input := &paymentcryptography.DeleteKeyInput{
		// KeyIdentifier: *string, // Required
	}

	if len(_paymentcryptographyKeyIdentifier) > 0 {
		input.KeyIdentifier = aws.String(_paymentcryptographyKeyIdentifier)
	}
	if len(_paymentcryptographyDeleteKeyInDays) > 0 {
		if err := assignInputField(input, "DeleteKeyInDays", _paymentcryptographyDeleteKeyInDays); err != nil {
			log.Errorf("invalid --delete-key-in-days: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables [Multi-Region key replication] settings for the specified Amazon Web Services Regions in your Amazon
// Web Services account, preventing new keys from being automatically replicated to
// those regions.
//
// After disabling Multi-Region key replication for specific regions, new keys
// created in your account will not be automatically replicated to those regions.
// You can still manually add replication to those regions for individual keys
// using the [AddKeyReplicationRegions]operation.
//
// This operation does not affect existing keys or their current replication
// configuration.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// [EnableDefaultKeyReplicationRegions]
//
// [GetDefaultKeyReplicationRegions]
//
// [AddKeyReplicationRegions]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_AddKeyReplicationRegions.html
// [EnableDefaultKeyReplicationRegions]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_EnableDefaultKeyReplicationRegions.html
// [GetDefaultKeyReplicationRegions]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_GetDefaultKeyReplicationRegions.html
// [Multi-Region key replication]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-multi-region-replication.html
func paymentcryptography_DisableDefaultKeyReplicationRegions(cfg aws.Config, client *paymentcryptography.Client) {
	input := &paymentcryptography.DisableDefaultKeyReplicationRegionsInput{
		// ReplicationRegions: []string, // Required
	}

	if len(_paymentcryptographyReplicationRegions) > 0 {
		input.ReplicationRegions = append([]string(nil), _paymentcryptographyReplicationRegions...)
	}

	if resp, err := client.DisableDefaultKeyReplicationRegions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables [Multi-Region key replication] settings for your Amazon Web Services account, causing new keys to be
// automatically replicated to the specified Amazon Web Services Regions when
// created.
//
// When Multi-Region key replication are enabled, any new keys created in your
// account will automatically be replicated to these regions unless you explicitly
// override this behavior during key creation. This simplifies key management for
// applications that operate across multiple regions.
//
// Existing keys are not affected by this operation - only keys created after
// enabling default replication will be automatically replicated.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// [DisableDefaultKeyReplicationRegions]
//
// [GetDefaultKeyReplicationRegions]
//
// [DisableDefaultKeyReplicationRegions]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_DisableDefaultKeyReplicationRegions.html
// [GetDefaultKeyReplicationRegions]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_GetDefaultKeyReplicationRegions.html
// [Multi-Region key replication]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-multi-region-replication.html
func paymentcryptography_EnableDefaultKeyReplicationRegions(cfg aws.Config, client *paymentcryptography.Client) {
	input := &paymentcryptography.EnableDefaultKeyReplicationRegionsInput{
		// ReplicationRegions: []string, // Required
	}

	if len(_paymentcryptographyReplicationRegions) > 0 {
		input.ReplicationRegions = append([]string(nil), _paymentcryptographyReplicationRegions...)
	}

	if resp, err := client.EnableDefaultKeyReplicationRegions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Exports a key from Amazon Web Services Payment Cryptography.
// Amazon Web Services Payment Cryptography simplifies key exchange by replacing
// the existing paper-based approach with a modern electronic approach. With
// ExportKey you can export symmetric keys using either symmetric and asymmetric
// key exchange mechanisms. Using this operation, you can share your Amazon Web
// Services Payment Cryptography generated keys with other service partners to
// perform cryptographic operations outside of Amazon Web Services Payment
// Cryptography
//
// For symmetric key exchange, Amazon Web Services Payment Cryptography uses the
// ANSI X9 TR-31 norm in accordance with PCI PIN guidelines. And for asymmetric key
// exchange, Amazon Web Services Payment Cryptography supports ANSI X9 TR-34 norm,
// RSA unwrap, and ECDH (Elliptic Curve Diffie-Hellman) key exchange mechanisms.
// Asymmetric key exchange methods are typically used to establish bi-directional
// trust between the two parties exhanging keys and are used for initial key
// exchange such as Key Encryption Key (KEK). After which you can export working
// keys using symmetric method to perform various cryptographic operations within
// Amazon Web Services Payment Cryptography.
//
// PCI requires specific minimum key strength of wrapping keys used to protect the
// keys being exchanged electronically. These requirements can change when PCI
// standards are revised. The rules specify that wrapping keys used for transport
// must be at least as strong as the key being protected. For more information on
// recommended key strength of wrapping keys and key exchange mechanism, see [Importing and exporting keys]in
// the Amazon Web Services Payment Cryptography User Guide.
//
// You can also use ExportKey functionality to generate and export an IPEK
// (Initial Pin Encryption Key) from Amazon Web Services Payment Cryptography using
// either TR-31 or TR-34 export key exchange. IPEK is generated from BDK (Base
// Derivation Key) and ExportDukptInitialKey attribute KSN ( KeySerialNumber ). The
// generated IPEK does not persist within Amazon Web Services Payment Cryptography
// and has to be re-generated each time during export.
//
// For key exchange using TR-31 or TR-34 key blocks, you can also export optional
// blocks within the key block header which contain additional attribute
// information about the key. The KeyVersion within KeyBlockHeaders indicates the
// version of the key within the key block. Furthermore, KeyExportability within
// KeyBlockHeaders can be used to further restrict exportability of the key after
// export from Amazon Web Services Payment Cryptography.
//
// The OptionalBlocks contain the additional data related to the key. For
// information on data type that can be included within optional blocks, refer to [ASC X9.143-2022].
//
// Data included in key block headers is signed but transmitted in clear text.
// Sensitive or confidential information should not be included in optional blocks.
// Refer to ASC X9.143-2022 standard for information on allowed data type.
//
// # To export initial keys (KEK) or IPEK using TR-34
//
// Using this operation, you can export initial key using TR-34 asymmetric key
// exchange. You can only export KEK generated within Amazon Web Services Payment
// Cryptography. In TR-34 terminology, the sending party of the key is called Key
// Distribution Host (KDH) and the receiving party of the key is called Key
// Receiving Device (KRD). During key export process, KDH is Amazon Web Services
// Payment Cryptography which initiates key export and KRD is the user receiving
// the key.
//
// To initiate TR-34 key export, the KRD must obtain an export token by calling [GetParametersForExport].
// This operation also generates a key pair for the purpose of key export, signs
// the key and returns back the signing public key certificate (also known as KDH
// signing certificate) and root certificate chain. The KDH uses the private key to
// sign the the export payload and the signing public key certificate is provided
// to KRD to verify the signature. The KRD can import the root certificate into its
// Hardware Security Module (HSM), as required. The export token and the associated
// KDH signing certificate expires after 30 days.
//
// Next the KRD generates a key pair for the the purpose of encrypting the KDH key
// and provides the public key cerificate (also known as KRD wrapping certificate)
// back to KDH. The KRD will also import the root cerificate chain into Amazon Web
// Services Payment Cryptography by calling [ImportKey]for RootCertificatePublicKey . The KDH,
// Amazon Web Services Payment Cryptography, will use the KRD wrapping cerificate
// to encrypt (wrap) the key under export and signs it with signing private key to
// generate a TR-34 WrappedKeyBlock. For more information on TR-34 key export, see
// section [Exporting symmetric keys]in the Amazon Web Services Payment Cryptography User Guide.
//
// Set the following parameters:
//
// - ExportAttributes : Specify export attributes in case of IPEK export. This
// parameter is optional for KEK export.
//
// - ExportKeyIdentifier : The KeyARN of the KEK or BDK (in case of IPEK) under
// export.
//
// - KeyMaterial : Use Tr34KeyBlock parameters.
//
// - CertificateAuthorityPublicKeyIdentifier : The KeyARN of the certificate
// chain that signed the KRD wrapping key certificate.
//
// - ExportToken : Obtained from KDH by calling [GetParametersForImport].
//
// - WrappingKeyCertificate : The public key certificate in PEM format (base64
// encoded) of the KRD wrapping key Amazon Web Services Payment Cryptography uses
// for encryption of the TR-34 export payload. This certificate must be signed by
// the root certificate (CertificateAuthorityPublicKeyIdentifier) imported into
// Amazon Web Services Payment Cryptography.
//
// When this operation is successful, Amazon Web Services Payment Cryptography
// returns the KEK or IPEK as a TR-34 WrappedKeyBlock.
//
// # To export initial keys (KEK) or IPEK using RSA Wrap and Unwrap
//
// Using this operation, you can export initial key using asymmetric RSA wrap and
// unwrap key exchange method. To initiate export, generate an asymmetric key pair
// on the receiving HSM and obtain the public key certificate in PEM format (base64
// encoded) for the purpose of wrapping and the root certifiate chain. Import the
// root certificate into Amazon Web Services Payment Cryptography by calling [ImportKey]for
// RootCertificatePublicKey .
//
// Next call ExportKey and set the following parameters:
//
// - CertificateAuthorityPublicKeyIdentifier : The KeyARN of the certificate
// chain that signed wrapping key certificate.
//
// - KeyMaterial : Set to KeyCryptogram .
//
// - WrappingKeyCertificate : The public key certificate in PEM format (base64
// encoded) obtained by the receiving HSM and signed by the root certificate
// (CertificateAuthorityPublicKeyIdentifier) imported into Amazon Web Services
// Payment Cryptography. The receiving HSM uses its private key component to unwrap
// the WrappedKeyCryptogram.
//
// When this operation is successful, Amazon Web Services Payment Cryptography
// returns the WrappedKeyCryptogram.
//
// # To export working keys or IPEK using TR-31
//
// Using this operation, you can export working keys or IPEK using TR-31 symmetric
// key exchange. In TR-31, you must use an initial key such as KEK to encrypt or
// wrap the key under export. To establish a KEK, you can use [CreateKey]or [ImportKey].
//
// Set the following parameters:
//
// - ExportAttributes : Specify export attributes in case of IPEK export. This
// parameter is optional for KEK export.
//
// - ExportKeyIdentifier : The KeyARN of the KEK or BDK (in case of IPEK) under
// export.
//
// - KeyMaterial : Use Tr31KeyBlock parameters.
//
// # To export working keys using ECDH
//
// You can also use ECDH key agreement to export working keys in a TR-31 keyblock,
// where the wrapping key is an ECDH derived key.
//
// To initiate a TR-31 key export using ECDH, both sides must create an ECC key
// pair with key usage K3 and exchange public key certificates. In Amazon Web
// Services Payment Cryptography, you can do this by calling CreateKey . If you
// have not already done so, you must import the CA chain that issued the receiving
// public key certificate by calling ImportKey with input RootCertificatePublicKey
// for root CA or TrustedPublicKey for intermediate CA. You can then complete a
// TR-31 key export by deriving a shared wrapping key using the service ECC key
// pair, public certificate of your ECC key pair outside of Amazon Web Services
// Payment Cryptography, and the key derivation parameters including key derivation
// function, hash algorithm, derivation data, key algorithm.
//
// - KeyMaterial : Use DiffieHellmanTr31KeyBlock parameters.
//
// - PrivateKeyIdentifier : The KeyArn of the ECC key pair created within Amazon
// Web Services Payment Cryptography to derive a shared KEK.
//
// - PublicKeyCertificate : The public key certificate of the receiving ECC key
// pair in PEM format (base64 encoded) to derive a shared KEK.
//
// - CertificateAuthorityPublicKeyIdentifier : The keyARN of the CA that signed
// the public key certificate of the receiving ECC key pair.
//
// When this operation is successful, Amazon Web Services Payment Cryptography
// returns the working key as a TR-31 WrappedKeyBlock, where the wrapping key is
// the ECDH derived key.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// [GetParametersForExport]
//
// [ImportKey]
//
// [Exporting symmetric keys]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-export.html
// [GetParametersForExport]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_GetParametersForExport.html
// [ImportKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_ImportKey.html
// [ASC X9.143-2022]: https://webstore.ansi.org/standards/ascx9/ansix91432022
// [GetParametersForImport]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_GetParametersForImport.html
// [Importing and exporting keys]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-importexport.html
// [CreateKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_CreateKey.html
func paymentcryptography_ExportKey(cfg aws.Config, client *paymentcryptography.Client) {
	input := &paymentcryptography.ExportKeyInput{
		// ExportKeyIdentifier: *string, // Required
		// KeyMaterial: types.ExportKeyMaterial, // Required
	}

	if len(_paymentcryptographyExportKeyIdentifier) > 0 {
		input.ExportKeyIdentifier = aws.String(_paymentcryptographyExportKeyIdentifier)
	}
	if len(_paymentcryptographyKeyMaterial) > 0 {
		if err := assignInputField(input, "KeyMaterial", _paymentcryptographyKeyMaterial); err != nil {
			log.Errorf("invalid --key-material: %s", err.Error())
			return
		}
	}
	if len(_paymentcryptographyExportAttributes) > 0 {
		if err := assignInputField(input, "ExportAttributes", _paymentcryptographyExportAttributes); err != nil {
			log.Errorf("invalid --export-attributes: %s", err.Error())
			return
		}
	}

	if resp, err := client.ExportKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the Amazon Web Services Payment Cryptography key associated with the alias.
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// [CreateAlias]
//
// [DeleteAlias]
//
// [ListAliases]
//
// [UpdateAlias]
//
// [ListAliases]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_ListAliases.html
// [DeleteAlias]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_DeleteAlias.html
// [UpdateAlias]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_UpdateAlias.html
// [CreateAlias]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_CreateAlias.html
func paymentcryptography_GetAlias(cfg aws.Config, client *paymentcryptography.Client) {
	input := &paymentcryptography.GetAliasInput{
		// AliasName: *string, // Required
	}

	if len(_paymentcryptographyAliasName) > 0 {
		input.AliasName = aws.String(_paymentcryptographyAliasName)
	}

	if resp, err := client.GetAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a certificate signing request (CSR) from a key pair.
func paymentcryptography_GetCertificateSigningRequest(cfg aws.Config, client *paymentcryptography.Client) {
	input := &paymentcryptography.GetCertificateSigningRequestInput{
		// CertificateSubject: *types.CertificateSubjectType, // Required
		// KeyIdentifier: *string, // Required
		// SigningAlgorithm: types.SigningAlgorithmType, // Required
	}

	if len(_paymentcryptographyCertificateSubject) > 0 {
		if err := assignInputField(input, "CertificateSubject", _paymentcryptographyCertificateSubject); err != nil {
			log.Errorf("invalid --certificate-subject: %s", err.Error())
			return
		}
	}
	if len(_paymentcryptographyKeyIdentifier) > 0 {
		input.KeyIdentifier = aws.String(_paymentcryptographyKeyIdentifier)
	}
	if len(_paymentcryptographySigningAlgorithm) > 0 {
		if err := assignInputField(input, "SigningAlgorithm", _paymentcryptographySigningAlgorithm); err != nil {
			log.Errorf("invalid --signing-algorithm: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetCertificateSigningRequest(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the list of Amazon Web Services Regions where [Multi-Region key replication] is currently enabled
// for your Amazon Web Services account.
//
// This operation returns the current Multi-Region key replication configuration.
// New keys created in your account will be automatically replicated to these
// regions unless explicitly overridden during key creation.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// [EnableDefaultKeyReplicationRegions]
//
// [DisableDefaultKeyReplicationRegions]
//
// [DisableDefaultKeyReplicationRegions]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_DisableDefaultKeyReplicationRegions.html
// [EnableDefaultKeyReplicationRegions]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_EnableDefaultKeyReplicationRegions.html
// [Multi-Region key replication]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-multi-region-replication.html
func paymentcryptography_GetDefaultKeyReplicationRegions(cfg aws.Config, client *paymentcryptography.Client) {
	input := &paymentcryptography.GetDefaultKeyReplicationRegionsInput{}

	if resp, err := client.GetDefaultKeyReplicationRegions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the key metadata for an Amazon Web Services Payment Cryptography key,
// including the immutable and mutable attributes specified when the key was
// created. Returns key metadata including attributes, state, and timestamps, but
// does not return the actual cryptographic key material.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// [CreateKey]
//
// [DeleteKey]
//
// [ListKeys]
//
// [DeleteKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_DeleteKey.html
// [ListKeys]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_ListKeys.html
// [CreateKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_CreateKey.html
func paymentcryptography_GetKey(cfg aws.Config, client *paymentcryptography.Client) {
	input := &paymentcryptography.GetKeyInput{
		// KeyIdentifier: *string, // Required
	}

	if len(_paymentcryptographyKeyIdentifier) > 0 {
		input.KeyIdentifier = aws.String(_paymentcryptographyKeyIdentifier)
	}

	if resp, err := client.GetKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the export token and the signing key certificate to initiate a TR-34 key
// export from Amazon Web Services Payment Cryptography.
//
// The signing key certificate signs the wrapped key under export within the TR-34
// key payload. The export token and signing key certificate must be in place and
// operational before calling [ExportKey]. The export token expires in 30 days. You can use
// the same export token to export multiple keys from your service account.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// [ExportKey]
//
// [GetParametersForImport]
//
// [ExportKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_ExportKey.html
// [GetParametersForImport]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_GetParametersForImport.html
func paymentcryptography_GetParametersForExport(cfg aws.Config, client *paymentcryptography.Client) {
	input := &paymentcryptography.GetParametersForExportInput{
		// KeyMaterialType: types.KeyMaterialType, // Required
		// SigningKeyAlgorithm: types.KeyAlgorithm, // Required
	}

	if len(_paymentcryptographyKeyMaterialType) > 0 {
		if err := assignInputField(input, "KeyMaterialType", _paymentcryptographyKeyMaterialType); err != nil {
			log.Errorf("invalid --key-material-type: %s", err.Error())
			return
		}
	}
	if len(_paymentcryptographySigningKeyAlgorithm) > 0 {
		if err := assignInputField(input, "SigningKeyAlgorithm", _paymentcryptographySigningKeyAlgorithm); err != nil {
			log.Errorf("invalid --signing-key-algorithm: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetParametersForExport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the import token and the wrapping key certificate in PEM format (base64
// encoded) to initiate a TR-34 WrappedKeyBlock or a RSA WrappedKeyCryptogram
// import into Amazon Web Services Payment Cryptography.
//
// The wrapping key certificate wraps the key under import. The import token and
// wrapping key certificate must be in place and operational before calling [ImportKey]. The
// import token expires in 30 days. You can use the same import token to import
// multiple keys into your service account.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// [GetParametersForExport]
//
// [ImportKey]
//
// [GetParametersForExport]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_GetParametersForExport.html
// [ImportKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_ImportKey.html
func paymentcryptography_GetParametersForImport(cfg aws.Config, client *paymentcryptography.Client) {
	input := &paymentcryptography.GetParametersForImportInput{
		// KeyMaterialType: types.KeyMaterialType, // Required
		// WrappingKeyAlgorithm: types.KeyAlgorithm, // Required
	}

	if len(_paymentcryptographyKeyMaterialType) > 0 {
		if err := assignInputField(input, "KeyMaterialType", _paymentcryptographyKeyMaterialType); err != nil {
			log.Errorf("invalid --key-material-type: %s", err.Error())
			return
		}
	}
	if len(_paymentcryptographyWrappingKeyAlgorithm) > 0 {
		if err := assignInputField(input, "WrappingKeyAlgorithm", _paymentcryptographyWrappingKeyAlgorithm); err != nil {
			log.Errorf("invalid --wrapping-key-algorithm: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetParametersForImport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the public key certificate of the asymmetric key pair that exists within
// Amazon Web Services Payment Cryptography.
//
// Unlike the private key of an asymmetric key, which never leaves Amazon Web
// Services Payment Cryptography unencrypted, callers with GetPublicKeyCertificate
// permission can download the public key certificate of the asymmetric key. You
// can share the public key certificate to allow others to encrypt messages and
// verify signatures outside of Amazon Web Services Payment Cryptography
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
func paymentcryptography_GetPublicKeyCertificate(cfg aws.Config, client *paymentcryptography.Client) {
	input := &paymentcryptography.GetPublicKeyCertificateInput{
		// KeyIdentifier: *string, // Required
	}

	if len(_paymentcryptographyKeyIdentifier) > 0 {
		input.KeyIdentifier = aws.String(_paymentcryptographyKeyIdentifier)
	}

	if resp, err := client.GetPublicKeyCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Imports symmetric keys and public key certificates in PEM format (base64
// encoded) into Amazon Web Services Payment Cryptography.
//
// Amazon Web Services Payment Cryptography simplifies key exchange by replacing
// the existing paper-based approach with a modern electronic approach. With
// ImportKey you can import symmetric keys using either symmetric and asymmetric
// key exchange mechanisms.
//
// For symmetric key exchange, Amazon Web Services Payment Cryptography uses the
// ANSI X9 TR-31 norm in accordance with PCI PIN guidelines. And for asymmetric key
// exchange, Amazon Web Services Payment Cryptography supports ANSI X9 TR-34 norm,
// RSA unwrap, and ECDH (Elliptic Curve Diffie-Hellman) key exchange mechanisms.
// Asymmetric key exchange methods are typically used to establish bi-directional
// trust between the two parties exhanging keys and are used for initial key
// exchange such as Key Encryption Key (KEK) or Zone Master Key (ZMK). After which
// you can import working keys using symmetric method to perform various
// cryptographic operations within Amazon Web Services Payment Cryptography.
//
// PCI requires specific minimum key strength of wrapping keys used to protect the
// keys being exchanged electronically. These requirements can change when PCI
// standards are revised. The rules specify that wrapping keys used for transport
// must be at least as strong as the key being protected. For more information on
// recommended key strength of wrapping keys and key exchange mechanism, see [Importing and exporting keys]in
// the Amazon Web Services Payment Cryptography User Guide.
//
// You can also import a root public key certificate, used to sign other public
// key certificates, or a trusted public key certificate under an already
// established root public key certificate.
//
// # To import a public root key certificate
//
// Using this operation, you can import the public component (in PEM cerificate
// format) of your private root key. You can use the imported public root key
// certificate for digital signatures, for example signing wrapping key or signing
// key in TR-34, within your Amazon Web Services Payment Cryptography account.
//
// Set the following parameters:
//
// - KeyMaterial : RootCertificatePublicKey
//
// - KeyClass : PUBLIC_KEY
//
// - KeyModesOfUse : Verify
//
// - KeyUsage : TR31_S0_ASYMMETRIC_KEY_FOR_DIGITAL_SIGNATURE
//
// - PublicKeyCertificate : The public key certificate in PEM format (base64
// encoded) of the private root key under import.
//
// # To import a trusted public key certificate
//
// The root public key certificate must be in place and operational before you
// import a trusted public key certificate. Set the following parameters:
//
// - KeyMaterial : TrustedCertificatePublicKey
//
// - CertificateAuthorityPublicKeyIdentifier : KeyArn of the
// RootCertificatePublicKey .
//
// - KeyModesOfUse and KeyUsage : Corresponding to the cryptographic operations
// such as wrap, sign, or encrypt that you will allow the trusted public key
// certificate to perform.
//
// - PublicKeyCertificate : The trusted public key certificate in PEM format
// (base64 encoded) under import.
//
// # To import initial keys (KEK or ZMK or similar) using TR-34
//
// Using this operation, you can import initial key using TR-34 asymmetric key
// exchange. In TR-34 terminology, the sending party of the key is called Key
// Distribution Host (KDH) and the receiving party of the key is called Key
// Receiving Device (KRD). During the key import process, KDH is the user who
// initiates the key import and KRD is Amazon Web Services Payment Cryptography who
// receives the key.
//
// To initiate TR-34 key import, the KDH must obtain an import token by calling [GetParametersForImport].
// This operation generates an encryption keypair for the purpose of key import,
// signs the key and returns back the wrapping key certificate (also known as KRD
// wrapping certificate) and the root certificate chain. The KDH must trust and
// install the KRD wrapping certificate on its HSM and use it to encrypt (wrap) the
// KDH key during TR-34 WrappedKeyBlock generation. The import token and associated
// KRD wrapping certificate expires after 30 days.
//
// Next the KDH generates a key pair for the purpose of signing the encrypted KDH
// key and provides the public certificate of the signing key to Amazon Web
// Services Payment Cryptography. The KDH will also need to import the root
// certificate chain of the KDH signing certificate by calling ImportKey for
// RootCertificatePublicKey . For more information on TR-34 key import, see section [Importing symmetric keys]
// in the Amazon Web Services Payment Cryptography User Guide.
//
// Set the following parameters:
//
// - KeyMaterial : Use Tr34KeyBlock parameters.
//
// - CertificateAuthorityPublicKeyIdentifier : The KeyARN of the certificate
// chain that signed the KDH signing key certificate.
//
// - ImportToken : Obtained from KRD by calling [GetParametersForImport].
//
// - WrappedKeyBlock : The TR-34 wrapped key material from KDH. It contains the
// KDH key under import, wrapped with KRD wrapping certificate and signed by KDH
// signing private key. This TR-34 key block is typically generated by the KDH
// Hardware Security Module (HSM) outside of Amazon Web Services Payment
// Cryptography.
//
// - SigningKeyCertificate : The public key certificate in PEM format (base64
// encoded) of the KDH signing key generated under the root certificate
// (CertificateAuthorityPublicKeyIdentifier) imported in Amazon Web Services
// Payment Cryptography.
//
// # To import initial keys (KEK or ZMK or similar) using RSA Wrap and Unwrap
//
// Using this operation, you can import initial key using asymmetric RSA wrap and
// unwrap key exchange method. To initiate import, call [GetParametersForImport]with KeyMaterial set to
// KEY_CRYPTOGRAM to generate an import token. This operation also generates an
// encryption keypair for the purpose of key import, signs the key and returns back
// the wrapping key certificate in PEM format (base64 encoded) and its root
// certificate chain. The import token and associated KRD wrapping certificate
// expires after 30 days.
//
// You must trust and install the wrapping certificate and its certificate chain
// on the sending HSM and use it to wrap the key under export for
// WrappedKeyCryptogram generation. Next call ImportKey with KeyMaterial set to
// KEY_CRYPTOGRAM and provide the ImportToken and KeyAttributes for the key under
// import.
//
// # To import working keys using TR-31
//
// Amazon Web Services Payment Cryptography uses TR-31 symmetric key exchange norm
// to import working keys. A KEK must be established within Amazon Web Services
// Payment Cryptography by using TR-34 key import or by using [CreateKey]. To initiate a
// TR-31 key import, set the following parameters:
//
// - KeyMaterial : Use Tr31KeyBlock parameters.
//
// - WrappedKeyBlock : The TR-31 wrapped key material. It contains the key under
// import, encrypted using KEK. The TR-31 key block is typically generated by a HSM
// outside of Amazon Web Services Payment Cryptography.
//
// - WrappingKeyIdentifier : The KeyArn of the KEK that Amazon Web Services
// Payment Cryptography uses to decrypt or unwrap the key under import.
//
// # To import working keys using ECDH
//
// You can also use ECDH key agreement to import working keys as a TR-31 keyblock,
// where the wrapping key is an ECDH derived key.
//
// To initiate a TR-31 key import using ECDH, both sides must create an ECC key
// pair with key usage K3 and exchange public key certificates. In Amazon Web
// Services Payment Cryptography, you can do this by calling CreateKey and then
// GetPublicKeyCertificate to retrieve its public key certificate. Next, you can
// then generate a TR-31 WrappedKeyBlock using your own ECC key pair, the public
// certificate of the service's ECC key pair, and the key derivation parameters
// including key derivation function, hash algorithm, derivation data, and key
// algorithm. If you have not already done so, you must import the CA chain that
// issued the receiving public key certificate by calling ImportKey with input
// RootCertificatePublicKey for root CA or TrustedPublicKey for intermediate CA.
// To complete the TR-31 key import, you can use the following parameters. It is
// important that the ECDH key derivation parameters you use should match those
// used during import to derive the same shared wrapping key within Amazon Web
// Services Payment Cryptography.
//
// - KeyMaterial : Use DiffieHellmanTr31KeyBlock parameters.
//
// - PrivateKeyIdentifier : The KeyArn of the ECC key pair created within Amazon
// Web Services Payment Cryptography to derive a shared KEK.
//
// - PublicKeyCertificate : The public key certificate of the receiving ECC key
// pair in PEM format (base64 encoded) to derive a shared KEK.
//
// - CertificateAuthorityPublicKeyIdentifier : The keyARN of the CA that signed
// the public key certificate of the receiving ECC key pair.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// [ExportKey]
//
// [GetParametersForImport]
//
// [Importing symmetric keys]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-import.html
// [ExportKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_ExportKey.html
// [GetParametersForImport]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_GetParametersForImport.html
// [Importing and exporting keys]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-importexport.html
// [CreateKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_CreateKey.html
func paymentcryptography_ImportKey(cfg aws.Config, client *paymentcryptography.Client) {
	input := &paymentcryptography.ImportKeyInput{
		// KeyMaterial: types.ImportKeyMaterial, // Required
	}

	if len(_paymentcryptographyKeyMaterial) > 0 {
		if err := assignInputField(input, "KeyMaterial", _paymentcryptographyKeyMaterial); err != nil {
			log.Errorf("invalid --key-material: %s", err.Error())
			return
		}
	}
	if len(_paymentcryptographyEnabled) > 0 {
		if err := assignInputField(input, "Enabled", _paymentcryptographyEnabled); err != nil {
			log.Errorf("invalid --enabled: %s", err.Error())
			return
		}
	}
	if len(_paymentcryptographyKeyCheckValueAlgorithm) > 0 {
		if err := assignInputField(input, "KeyCheckValueAlgorithm", _paymentcryptographyKeyCheckValueAlgorithm); err != nil {
			log.Errorf("invalid --key-check-value-algorithm: %s", err.Error())
			return
		}
	}
	if len(_paymentcryptographyReplicationRegions) > 0 {
		input.ReplicationRegions = append([]string(nil), _paymentcryptographyReplicationRegions...)
	}
	if len(_paymentcryptographyTags) > 0 {
		if err := assignInputField(input, "Tags", _paymentcryptographyTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.ImportKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the aliases for all keys in the caller's Amazon Web Services account and
// Amazon Web Services Region. You can filter the aliases by keyARN . For more
// information, see [Using aliases]in the Amazon Web Services Payment Cryptography User Guide.
//
// This is a paginated operation, which means that each response might contain
// only a subset of all the aliases. When the response contains only a subset of
// aliases, it includes a NextToken value. Use this value in a subsequent
// ListAliases request to get more aliases. When you receive a response with no
// NextToken (or an empty or null value), that means there are no more aliases to
// get.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// [CreateAlias]
//
// [DeleteAlias]
//
// [GetAlias]
//
// [UpdateAlias]
//
// [DeleteAlias]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_DeleteAlias.html
// [UpdateAlias]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_UpdateAlias.html
// [Using aliases]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-managealias.html
// [CreateAlias]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_CreateAlias.html
// [GetAlias]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_GetAlias.html
func paymentcryptography_ListAliases(cfg aws.Config, client *paymentcryptography.Client) {
	input := &paymentcryptography.ListAliasesInput{}

	if len(_paymentcryptographyKeyArn) > 0 {
		input.KeyArn = aws.String(_paymentcryptographyKeyArn)
	}
	if len(_paymentcryptographyMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _paymentcryptographyMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_paymentcryptographyNextToken) > 0 {
		input.NextToken = aws.String(_paymentcryptographyNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAliases(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*paymentcryptography.ListAliasesOutput
	p := paymentcryptography.NewListAliasesPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Lists the keys in the caller's Amazon Web Services account and Amazon Web
// Services Region. You can filter the list of keys.
//
// This is a paginated operation, which means that each response might contain
// only a subset of all the keys. When the response contains only a subset of keys,
// it includes a NextToken value. Use this value in a subsequent ListKeys request
// to get more keys. When you receive a response with no NextToken (or an empty or
// null value), that means there are no more keys to get.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// [CreateKey]
//
// [DeleteKey]
//
// [GetKey]
//
// [DeleteKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_DeleteKey.html
// [GetKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_GetKey.html
// [CreateKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_CreateKey.html
func paymentcryptography_ListKeys(cfg aws.Config, client *paymentcryptography.Client) {
	input := &paymentcryptography.ListKeysInput{}

	if len(_paymentcryptographyKeyState) > 0 {
		if err := assignInputField(input, "KeyState", _paymentcryptographyKeyState); err != nil {
			log.Errorf("invalid --key-state: %s", err.Error())
			return
		}
	}
	if len(_paymentcryptographyMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _paymentcryptographyMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_paymentcryptographyNextToken) > 0 {
		input.NextToken = aws.String(_paymentcryptographyNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListKeys(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*paymentcryptography.ListKeysOutput
	p := paymentcryptography.NewListKeysPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Lists the tags for an Amazon Web Services resource.
// This is a paginated operation, which means that each response might contain
// only a subset of all the tags. When the response contains only a subset of tags,
// it includes a NextToken value. Use this value in a subsequent
// ListTagsForResource request to get more tags. When you receive a response with
// no NextToken (or an empty or null value), that means there are no more tags to
// get.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// [TagResource]
//
// [UntagResource]
//
// [TagResource]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_TagResource.html
// [UntagResource]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_UntagResource.html
func paymentcryptography_ListTagsForResource(cfg aws.Config, client *paymentcryptography.Client) {
	input := &paymentcryptography.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_paymentcryptographyResourceArn) > 0 {
		input.ResourceArn = aws.String(_paymentcryptographyResourceArn)
	}
	if len(_paymentcryptographyMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _paymentcryptographyMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_paymentcryptographyNextToken) > 0 {
		input.NextToken = aws.String(_paymentcryptographyNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*paymentcryptography.ListTagsForResourceOutput
	p := paymentcryptography.NewListTagsForResourcePaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Removes Replication Regions from an existing Amazon Web Services Payment
// Cryptography key, disabling the key's availability for cryptographic operations
// in the specified Amazon Web Services Regions.
//
// When you remove Replication Regions, the key material is securely deleted from
// those regions and can no longer be used for cryptographic operations there. This
// operation is irreversible for the specified Amazon Web Services Regions. For
// more information, see [Multi-Region key replication].
//
// Ensure that no active cryptographic operations or applications depend on the
// key in the regions you're removing before performing this operation.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// [AddKeyReplicationRegions]
//
// [DisableDefaultKeyReplicationRegions]
//
// [Multi-Region key replication]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-multi-region-replication.html
// [DisableDefaultKeyReplicationRegions]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_DisableDefaultKeyReplicationRegions.html
// [AddKeyReplicationRegions]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_AddKeyReplicationRegions.html
func paymentcryptography_RemoveKeyReplicationRegions(cfg aws.Config, client *paymentcryptography.Client) {
	input := &paymentcryptography.RemoveKeyReplicationRegionsInput{
		// KeyIdentifier: *string, // Required
		// ReplicationRegions: []string, // Required
	}

	if len(_paymentcryptographyKeyIdentifier) > 0 {
		input.KeyIdentifier = aws.String(_paymentcryptographyKeyIdentifier)
	}
	if len(_paymentcryptographyReplicationRegions) > 0 {
		input.ReplicationRegions = append([]string(nil), _paymentcryptographyReplicationRegions...)
	}

	if resp, err := client.RemoveKeyReplicationRegions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels a scheduled key deletion during the waiting period. Use this operation
// to restore a Key that is scheduled for deletion.
//
// During the waiting period, the KeyState is DELETE_PENDING and
// deletePendingTimestamp contains the date and time after which the Key will be
// deleted. After Key is restored, the KeyState is CREATE_COMPLETE , and the value
// for deletePendingTimestamp is removed.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// [DeleteKey]
//
// [StartKeyUsage]
//
// [StopKeyUsage]
//
// [DeleteKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_DeleteKey.html
// [StartKeyUsage]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_StartKeyUsage.html
// [StopKeyUsage]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_StopKeyUsage.html
func paymentcryptography_RestoreKey(cfg aws.Config, client *paymentcryptography.Client) {
	input := &paymentcryptography.RestoreKeyInput{
		// KeyIdentifier: *string, // Required
	}

	if len(_paymentcryptographyKeyIdentifier) > 0 {
		input.KeyIdentifier = aws.String(_paymentcryptographyKeyIdentifier)
	}

	if resp, err := client.RestoreKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables an Amazon Web Services Payment Cryptography key, which makes it active
// for cryptographic operations within Amazon Web Services Payment Cryptography
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// [StopKeyUsage]
//
// [StopKeyUsage]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_StopKeyUsage.html
func paymentcryptography_StartKeyUsage(cfg aws.Config, client *paymentcryptography.Client) {
	input := &paymentcryptography.StartKeyUsageInput{
		// KeyIdentifier: *string, // Required
	}

	if len(_paymentcryptographyKeyIdentifier) > 0 {
		input.KeyIdentifier = aws.String(_paymentcryptographyKeyIdentifier)
	}

	if resp, err := client.StartKeyUsage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables an Amazon Web Services Payment Cryptography key, which makes it
// inactive within Amazon Web Services Payment Cryptography.
//
// You can use this operation instead of [DeleteKey] to deactivate a key. You can enable the
// key in the future by calling [StartKeyUsage].
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// [DeleteKey]
//
// [StartKeyUsage]
//
// [DeleteKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_DeleteKey.html
// [StartKeyUsage]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_StartKeyUsage.html
func paymentcryptography_StopKeyUsage(cfg aws.Config, client *paymentcryptography.Client) {
	input := &paymentcryptography.StopKeyUsageInput{
		// KeyIdentifier: *string, // Required
	}

	if len(_paymentcryptographyKeyIdentifier) > 0 {
		input.KeyIdentifier = aws.String(_paymentcryptographyKeyIdentifier)
	}

	if resp, err := client.StopKeyUsage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or edits tags on an Amazon Web Services Payment Cryptography key.
// Tagging or untagging an Amazon Web Services Payment Cryptography key can allow
// or deny permission to the key.
//
// Each tag consists of a tag key and a tag value, both of which are
// case-sensitive strings. The tag value can be an empty (null) string. To add a
// tag, specify a new tag key and a tag value. To edit a tag, specify an existing
// tag key and a new tag value. You can also add tags to an Amazon Web Services
// Payment Cryptography key when you create it with [CreateKey].
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// [ListTagsForResource]
//
// [UntagResource]
//
// [UntagResource]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_UntagResource.html
// [CreateKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_CreateKey.html
// [ListTagsForResource]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_ListTagsForResource.html
func paymentcryptography_TagResource(cfg aws.Config, client *paymentcryptography.Client) {
	input := &paymentcryptography.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_paymentcryptographyResourceArn) > 0 {
		input.ResourceArn = aws.String(_paymentcryptographyResourceArn)
	}
	if len(_paymentcryptographyTags) > 0 {
		if err := assignInputField(input, "Tags", _paymentcryptographyTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.TagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a tag from an Amazon Web Services Payment Cryptography key.
// Tagging or untagging an Amazon Web Services Payment Cryptography key can allow
// or deny permission to the key.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// [ListTagsForResource]
//
// [TagResource]
//
// [TagResource]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_TagResource.html
// [ListTagsForResource]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_ListTagsForResource.html
func paymentcryptography_UntagResource(cfg aws.Config, client *paymentcryptography.Client) {
	input := &paymentcryptography.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_paymentcryptographyResourceArn) > 0 {
		input.ResourceArn = aws.String(_paymentcryptographyResourceArn)
	}
	if len(_paymentcryptographyTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _paymentcryptographyTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates an existing Amazon Web Services Payment Cryptography alias with a
// different key. Each alias is associated with only one Amazon Web Services
// Payment Cryptography key at a time, although a key can have multiple aliases.
// The alias and the Amazon Web Services Payment Cryptography key must be in the
// same Amazon Web Services account and Amazon Web Services Region
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// [CreateAlias]
//
// [DeleteAlias]
//
// [GetAlias]
//
// [ListAliases]
//
// [ListAliases]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_ListAliases.html
// [DeleteAlias]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_DeleteAlias.html
// [CreateAlias]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_CreateAlias.html
// [GetAlias]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_GetAlias.html
func paymentcryptography_UpdateAlias(cfg aws.Config, client *paymentcryptography.Client) {
	input := &paymentcryptography.UpdateAliasInput{
		// AliasName: *string, // Required
	}

	if len(_paymentcryptographyAliasName) > 0 {
		input.AliasName = aws.String(_paymentcryptographyAliasName)
	}
	if len(_paymentcryptographyKeyArn) > 0 {
		input.KeyArn = aws.String(_paymentcryptographyKeyArn)
	}

	if resp, err := client.UpdateAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_paymentcryptographyCmd)
	_paymentcryptographyCmd.Flags().SortFlags = false

	_paymentcryptographyCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_paymentcryptographyCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_paymentcryptographyCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_paymentcryptographyCmd.Flags().StringVarP(&_paymentcryptographyAliasName, "alias-name", "", "", "Alias Name")
	_paymentcryptographyCmd.Flags().StringVarP(&_paymentcryptographyCertificateSubject, "certificate-subject", "", "", "Certificate Subject")
	_paymentcryptographyCmd.Flags().StringVarP(&_paymentcryptographyDeleteKeyInDays, "delete-key-in-days", "", "", "Delete Key In Days")
	_paymentcryptographyCmd.Flags().StringVarP(&_paymentcryptographyDeriveKeyUsage, "derive-key-usage", "", "", "Derive Key Usage")
	_paymentcryptographyCmd.Flags().StringVarP(&_paymentcryptographyEnabled, "enabled", "", "", "Enabled")
	_paymentcryptographyCmd.Flags().StringVarP(&_paymentcryptographyExportAttributes, "export-attributes", "", "", "Export Attributes")
	_paymentcryptographyCmd.Flags().StringVarP(&_paymentcryptographyExportKeyIdentifier, "export-key-identifier", "", "", "Export Key Identifier")
	_paymentcryptographyCmd.Flags().StringVarP(&_paymentcryptographyExportable, "exportable", "", "", "Exportable")
	_paymentcryptographyCmd.Flags().StringVarP(&_paymentcryptographyKeyArn, "key-arn", "", "", "Key ARN")
	_paymentcryptographyCmd.Flags().StringVarP(&_paymentcryptographyKeyAttributes, "key-attributes", "", "", "Key Attributes")
	_paymentcryptographyCmd.Flags().StringVarP(&_paymentcryptographyKeyCheckValueAlgorithm, "key-check-value-algorithm", "", "", "Key Check Value Algorithm")
	_paymentcryptographyCmd.Flags().StringVarP(&_paymentcryptographyKeyIdentifier, "key-identifier", "", "", "Key Identifier")
	_paymentcryptographyCmd.Flags().StringVarP(&_paymentcryptographyKeyMaterial, "key-material", "", "", "Key Material")
	_paymentcryptographyCmd.Flags().StringVarP(&_paymentcryptographyKeyMaterialType, "key-material-type", "", "", "Key Material Type")
	_paymentcryptographyCmd.Flags().StringVarP(&_paymentcryptographyKeyState, "key-state", "", "", "Key State")
	_paymentcryptographyCmd.Flags().StringVarP(&_paymentcryptographyMaxResults, "max-results", "", "", "Max Results")
	_paymentcryptographyCmd.Flags().StringVarP(&_paymentcryptographyNextToken, "next-token", "", "", "Next Token")
	_paymentcryptographyCmd.Flags().StringSliceVarP(&_paymentcryptographyReplicationRegions, "replication-regions", "", nil, "Replication Regions")
	_paymentcryptographyCmd.Flags().StringVarP(&_paymentcryptographyResourceArn, "resource-arn", "", "", "Resource ARN")
	_paymentcryptographyCmd.Flags().StringVarP(&_paymentcryptographySigningAlgorithm, "signing-algorithm", "", "", "Signing Algorithm")
	_paymentcryptographyCmd.Flags().StringVarP(&_paymentcryptographySigningKeyAlgorithm, "signing-key-algorithm", "", "", "Signing Key Algorithm")
	_paymentcryptographyCmd.Flags().StringSliceVarP(&_paymentcryptographyTagKeys, "tag-keys", "", nil, "Tag Keys")
	_paymentcryptographyCmd.Flags().StringVarP(&_paymentcryptographyTags, "tags", "", "", "Tags")
	_paymentcryptographyCmd.Flags().StringVarP(&_paymentcryptographyWrappingKeyAlgorithm, "wrapping-key-algorithm", "", "", "Wrapping Key Algorithm")

	_paymentcryptographyCmd.Flags().BoolVarP(&_paymentcryptographyAddKeyReplicationRegions, "add-key-replication-regions", "", false, "Add Key Replication Regions")
	_paymentcryptographyCmd.Flags().BoolVarP(&_paymentcryptographyCreateAlias, "create-alias", "", false, "Create Alias")
	_paymentcryptographyCmd.Flags().BoolVarP(&_paymentcryptographyCreateKey, "create-key", "", false, "Create Key")
	_paymentcryptographyCmd.Flags().BoolVarP(&_paymentcryptographyDeleteAlias, "delete-alias", "", false, "Delete Alias")
	_paymentcryptographyCmd.Flags().BoolVarP(&_paymentcryptographyDeleteKey, "delete-key", "", false, "Delete Key")
	_paymentcryptographyCmd.Flags().BoolVarP(&_paymentcryptographyDisableDefaultKeyReplicationRegions, "disable-default-key-replication-regions", "", false, "Disable Default Key Replication Regions")
	_paymentcryptographyCmd.Flags().BoolVarP(&_paymentcryptographyEnableDefaultKeyReplicationRegions, "enable-default-key-replication-regions", "", false, "Enable Default Key Replication Regions")
	_paymentcryptographyCmd.Flags().BoolVarP(&_paymentcryptographyExportKey, "export-key", "", false, "Export Key")
	_paymentcryptographyCmd.Flags().BoolVarP(&_paymentcryptographyGetAlias, "get-alias", "", false, "Get Alias")
	_paymentcryptographyCmd.Flags().BoolVarP(&_paymentcryptographyGetCertificateSigningRequest, "get-certificate-signing-request", "", false, "Get Certificate Signing Request")
	_paymentcryptographyCmd.Flags().BoolVarP(&_paymentcryptographyGetDefaultKeyReplicationRegions, "get-default-key-replication-regions", "", false, "Get Default Key Replication Regions")
	_paymentcryptographyCmd.Flags().BoolVarP(&_paymentcryptographyGetKey, "get-key", "", false, "Get Key")
	_paymentcryptographyCmd.Flags().BoolVarP(&_paymentcryptographyGetParametersForExport, "get-parameters-for-export", "", false, "Get Parameters For Export")
	_paymentcryptographyCmd.Flags().BoolVarP(&_paymentcryptographyGetParametersForImport, "get-parameters-for-import", "", false, "Get Parameters For Import")
	_paymentcryptographyCmd.Flags().BoolVarP(&_paymentcryptographyGetPublicKeyCertificate, "get-public-key-certificate", "", false, "Get Public Key Certificate")
	_paymentcryptographyCmd.Flags().BoolVarP(&_paymentcryptographyImportKey, "import-key", "", false, "Import Key")
	_paymentcryptographyCmd.Flags().BoolVarP(&_paymentcryptographyListAliases, "list-aliases", "", false, "List Aliases")
	_paymentcryptographyCmd.Flags().BoolVarP(&_paymentcryptographyListKeys, "list-keys", "", false, "List Keys")
	_paymentcryptographyCmd.Flags().BoolVarP(&_paymentcryptographyListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_paymentcryptographyCmd.Flags().BoolVarP(&_paymentcryptographyRemoveKeyReplicationRegions, "remove-key-replication-regions", "", false, "Remove Key Replication Regions")
	_paymentcryptographyCmd.Flags().BoolVarP(&_paymentcryptographyRestoreKey, "restore-key", "", false, "Restore Key")
	_paymentcryptographyCmd.Flags().BoolVarP(&_paymentcryptographyStartKeyUsage, "start-key-usage", "", false, "Start Key Usage")
	_paymentcryptographyCmd.Flags().BoolVarP(&_paymentcryptographyStopKeyUsage, "stop-key-usage", "", false, "Stop Key Usage")
	_paymentcryptographyCmd.Flags().BoolVarP(&_paymentcryptographyTagResource, "tag-resource", "", false, "Tag Resource")
	_paymentcryptographyCmd.Flags().BoolVarP(&_paymentcryptographyUntagResource, "untag-resource", "", false, "Untag Resource")
	_paymentcryptographyCmd.Flags().BoolVarP(&_paymentcryptographyUpdateAlias, "update-alias", "", false, "Update Alias")

}
