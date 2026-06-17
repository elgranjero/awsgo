package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// kmsCmd represents the kms command
var _kmsCmd = &cobra.Command{
	Use:   "kms",
	Short: "AWS kms CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := kms.NewFromConfig(cfg)
		if _kmsCancelKeyDeletion {
			kms_CancelKeyDeletion(cfg, client)
			return
		}
		if _kmsConnectCustomKeyStore {
			kms_ConnectCustomKeyStore(cfg, client)
			return
		}
		if _kmsCreateAlias {
			kms_CreateAlias(cfg, client)
			return
		}
		if _kmsCreateCustomKeyStore {
			kms_CreateCustomKeyStore(cfg, client)
			return
		}
		if _kmsCreateGrant {
			kms_CreateGrant(cfg, client)
			return
		}
		if _kmsCreateKey {
			kms_CreateKey(cfg, client)
			return
		}
		if _kmsDecrypt {
			kms_Decrypt(cfg, client)
			return
		}
		if _kmsDeleteAlias {
			kms_DeleteAlias(cfg, client)
			return
		}
		if _kmsDeleteCustomKeyStore {
			kms_DeleteCustomKeyStore(cfg, client)
			return
		}
		if _kmsDeleteImportedKeyMaterial {
			kms_DeleteImportedKeyMaterial(cfg, client)
			return
		}
		if _kmsDeriveSharedSecret {
			kms_DeriveSharedSecret(cfg, client)
			return
		}
		if _kmsDescribeCustomKeyStores {
			kms_DescribeCustomKeyStores(cfg, client)
			return
		}
		if _kmsDescribeKey {
			kms_DescribeKey(cfg, client)
			return
		}
		if _kmsDisableKey {
			kms_DisableKey(cfg, client)
			return
		}
		if _kmsDisableKeyRotation {
			kms_DisableKeyRotation(cfg, client)
			return
		}
		if _kmsDisconnectCustomKeyStore {
			kms_DisconnectCustomKeyStore(cfg, client)
			return
		}
		if _kmsEnableKey {
			kms_EnableKey(cfg, client)
			return
		}
		if _kmsEnableKeyRotation {
			kms_EnableKeyRotation(cfg, client)
			return
		}
		if _kmsEncrypt {
			kms_Encrypt(cfg, client)
			return
		}
		if _kmsGenerateDataKey {
			kms_GenerateDataKey(cfg, client)
			return
		}
		if _kmsGenerateDataKeyPair {
			kms_GenerateDataKeyPair(cfg, client)
			return
		}
		if _kmsGenerateDataKeyPairWithoutPlaintext {
			kms_GenerateDataKeyPairWithoutPlaintext(cfg, client)
			return
		}
		if _kmsGenerateDataKeyWithoutPlaintext {
			kms_GenerateDataKeyWithoutPlaintext(cfg, client)
			return
		}
		if _kmsGenerateMac {
			kms_GenerateMac(cfg, client)
			return
		}
		if _kmsGenerateRandom {
			kms_GenerateRandom(cfg, client)
			return
		}
		if _kmsGetKeyPolicy {
			kms_GetKeyPolicy(cfg, client)
			return
		}
		if _kmsGetKeyRotationStatus {
			kms_GetKeyRotationStatus(cfg, client)
			return
		}
		if _kmsGetParametersForImport {
			kms_GetParametersForImport(cfg, client)
			return
		}
		if _kmsGetPublicKey {
			kms_GetPublicKey(cfg, client)
			return
		}
		if _kmsImportKeyMaterial {
			kms_ImportKeyMaterial(cfg, client)
			return
		}
		if _kmsListAliases {
			kms_ListAliases(cfg, client)
			return
		}
		if _kmsListGrants {
			kms_ListGrants(cfg, client)
			return
		}
		if _kmsListKeyPolicies {
			kms_ListKeyPolicies(cfg, client)
			return
		}
		if _kmsListKeyRotations {
			kms_ListKeyRotations(cfg, client)
			return
		}
		if _kmsListKeys {
			kms_ListKeys(cfg, client)
			return
		}
		if _kmsListResourceTags {
			kms_ListResourceTags(cfg, client)
			return
		}
		if _kmsListRetirableGrants {
			kms_ListRetirableGrants(cfg, client)
			return
		}
		if _kmsPutKeyPolicy {
			kms_PutKeyPolicy(cfg, client)
			return
		}
		if _kmsReEncrypt {
			kms_ReEncrypt(cfg, client)
			return
		}
		if _kmsReplicateKey {
			kms_ReplicateKey(cfg, client)
			return
		}
		if _kmsRetireGrant {
			kms_RetireGrant(cfg, client)
			return
		}
		if _kmsRevokeGrant {
			kms_RevokeGrant(cfg, client)
			return
		}
		if _kmsRotateKeyOnDemand {
			kms_RotateKeyOnDemand(cfg, client)
			return
		}
		if _kmsScheduleKeyDeletion {
			kms_ScheduleKeyDeletion(cfg, client)
			return
		}
		if _kmsSign {
			kms_Sign(cfg, client)
			return
		}
		if _kmsTagResource {
			kms_TagResource(cfg, client)
			return
		}
		if _kmsUntagResource {
			kms_UntagResource(cfg, client)
			return
		}
		if _kmsUpdateAlias {
			kms_UpdateAlias(cfg, client)
			return
		}
		if _kmsUpdateCustomKeyStore {
			kms_UpdateCustomKeyStore(cfg, client)
			return
		}
		if _kmsUpdateKeyDescription {
			kms_UpdateKeyDescription(cfg, client)
			return
		}
		if _kmsUpdatePrimaryRegion {
			kms_UpdatePrimaryRegion(cfg, client)
			return
		}
		if _kmsVerify {
			kms_Verify(cfg, client)
			return
		}
		if _kmsVerifyMac {
			kms_VerifyMac(cfg, client)
			return
		}

	},
}

var (
	_kmsCancelKeyDeletion                   bool
	_kmsConnectCustomKeyStore               bool
	_kmsCreateAlias                         bool
	_kmsCreateCustomKeyStore                bool
	_kmsCreateGrant                         bool
	_kmsCreateKey                           bool
	_kmsDecrypt                             bool
	_kmsDeleteAlias                         bool
	_kmsDeleteCustomKeyStore                bool
	_kmsDeleteImportedKeyMaterial           bool
	_kmsDeriveSharedSecret                  bool
	_kmsDescribeCustomKeyStores             bool
	_kmsDescribeKey                         bool
	_kmsDisableKey                          bool
	_kmsDisableKeyRotation                  bool
	_kmsDisconnectCustomKeyStore            bool
	_kmsEnableKey                           bool
	_kmsEnableKeyRotation                   bool
	_kmsEncrypt                             bool
	_kmsGenerateDataKey                     bool
	_kmsGenerateDataKeyPair                 bool
	_kmsGenerateDataKeyPairWithoutPlaintext bool
	_kmsGenerateDataKeyWithoutPlaintext     bool
	_kmsGenerateMac                         bool
	_kmsGenerateRandom                      bool
	_kmsGetKeyPolicy                        bool
	_kmsGetKeyRotationStatus                bool
	_kmsGetParametersForImport              bool
	_kmsGetPublicKey                        bool
	_kmsImportKeyMaterial                   bool
	_kmsListAliases                         bool
	_kmsListGrants                          bool
	_kmsListKeyPolicies                     bool
	_kmsListKeyRotations                    bool
	_kmsListKeys                            bool
	_kmsListResourceTags                    bool
	_kmsListRetirableGrants                 bool
	_kmsPutKeyPolicy                        bool
	_kmsReEncrypt                           bool
	_kmsReplicateKey                        bool
	_kmsRetireGrant                         bool
	_kmsRevokeGrant                         bool
	_kmsRotateKeyOnDemand                   bool
	_kmsScheduleKeyDeletion                 bool
	_kmsSign                                bool
	_kmsTagResource                         bool
	_kmsUntagResource                       bool
	_kmsUpdateAlias                         bool
	_kmsUpdateCustomKeyStore                bool
	_kmsUpdateKeyDescription                bool
	_kmsUpdatePrimaryRegion                 bool
	_kmsVerify                              bool
	_kmsVerifyMac                           bool

	_kmsAliasName                        string
	_kmsBypassPolicyLockoutSafetyCheck   string
	_kmsCiphertextBlob                   string
	_kmsCloudHsmClusterId                string
	_kmsConstraints                      string
	_kmsCustomKeyStoreId                 string
	_kmsCustomKeyStoreName               string
	_kmsCustomKeyStoreType               string
	_kmsCustomerMasterKeySpec            string
	_kmsDescription                      string
	_kmsDestinationEncryptionAlgorithm   string
	_kmsDestinationEncryptionContext     string
	_kmsDestinationKeyId                 string
	_kmsDryRun                           string
	_kmsDryRunModifiers                  string
	_kmsEncryptedKeyMaterial             string
	_kmsEncryptionAlgorithm              string
	_kmsEncryptionContext                string
	_kmsExpirationModel                  string
	_kmsGrantId                          string
	_kmsGrantToken                       string
	_kmsGrantTokens                      []string
	_kmsGranteePrincipal                 string
	_kmsImportToken                      string
	_kmsImportType                       string
	_kmsIncludeKeyMaterial               string
	_kmsKeyAgreementAlgorithm            string
	_kmsKeyId                            string
	_kmsKeyMaterialDescription           string
	_kmsKeyMaterialId                    string
	_kmsKeyPairSpec                      string
	_kmsKeySpec                          string
	_kmsKeyStorePassword                 string
	_kmsKeyUsage                         string
	_kmsLimit                            string
	_kmsMac                              string
	_kmsMacAlgorithm                     string
	_kmsMarker                           string
	_kmsMessage                          string
	_kmsMessageType                      string
	_kmsMultiRegion                      string
	_kmsName                             string
	_kmsNewCustomKeyStoreName            string
	_kmsNumberOfBytes                    string
	_kmsOperations                       string
	_kmsOrigin                           string
	_kmsPendingWindowInDays              string
	_kmsPlaintext                        string
	_kmsPolicy                           string
	_kmsPolicyName                       string
	_kmsPrimaryRegion                    string
	_kmsPublicKey                        string
	_kmsRecipient                        string
	_kmsReplicaRegion                    string
	_kmsRetiringPrincipal                string
	_kmsRotationPeriodInDays             string
	_kmsSignature                        string
	_kmsSigningAlgorithm                 string
	_kmsSourceEncryptionAlgorithm        string
	_kmsSourceEncryptionContext          string
	_kmsSourceKeyId                      string
	_kmsTagKeys                          []string
	_kmsTags                             string
	_kmsTargetKeyId                      string
	_kmsTrustAnchorCertificate           string
	_kmsValidTo                          string
	_kmsWrappingAlgorithm                string
	_kmsWrappingKeySpec                  string
	_kmsXksKeyId                         string
	_kmsXksProxyAuthenticationCredential string
	_kmsXksProxyConnectivity             string
	_kmsXksProxyUriEndpoint              string
	_kmsXksProxyUriPath                  string
	_kmsXksProxyVpcEndpointServiceName   string
	_kmsXksProxyVpcEndpointServiceOwner  string
)

// Cancels the deletion of a KMS key. When this operation succeeds, the key state
// of the KMS key is Disabled . To enable the KMS key, use EnableKey.
//
// For more information about scheduling and canceling deletion of a KMS key, see [Deleting KMS keys]
// in the Key Management Service Developer Guide.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: No. You cannot perform this operation on a KMS key in a
// different Amazon Web Services account.
//
// Required permissions: [kms:CancelKeyDeletion] (key policy)
//
// Related operations: ScheduleKeyDeletion
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [kms:CancelKeyDeletion]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [Deleting KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/deleting-keys.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
func kms_CancelKeyDeletion(cfg aws.Config, client *kms.Client) {
	input := &kms.CancelKeyDeletionInput{
		// KeyId: *string, // Required
	}

	if len(_kmsKeyId) > 0 {
		input.KeyId = aws.String(_kmsKeyId)
	}

	if resp, err := client.CancelKeyDeletion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Connects or reconnects a [custom key store] to its backing key store. For an CloudHSM key store,
// ConnectCustomKeyStore connects the key store to its associated CloudHSM cluster.
// For an external key store, ConnectCustomKeyStore connects the key store to the
// external key store proxy that communicates with your external key manager.
//
// The custom key store must be connected before you can create KMS keys in the
// key store or use the KMS keys it contains. You can disconnect and reconnect a
// custom key store at any time.
//
// The connection process for a custom key store can take an extended amount of
// time to complete. This operation starts the connection process, but it does not
// wait for it to complete. When it succeeds, this operation quickly returns an
// HTTP 200 response and a JSON object with no properties. However, this response
// does not indicate that the custom key store is connected. To get the connection
// state of the custom key store, use the DescribeCustomKeyStoresoperation.
//
// This operation is part of the custom key stores feature in KMS, which combines
// the convenience and extensive integration of KMS with the isolation and control
// of a key store that you own and manage.
//
// The ConnectCustomKeyStore operation might fail for various reasons. To find the
// reason, use the DescribeCustomKeyStoresoperation and see the ConnectionErrorCode in the response. For
// help interpreting the ConnectionErrorCode , see CustomKeyStoresListEntry.
//
// To fix the failure, use the DisconnectCustomKeyStore operation to disconnect the custom key store,
// correct the error, use the UpdateCustomKeyStoreoperation if necessary, and then use
// ConnectCustomKeyStore again.
//
// # CloudHSM key store
//
// During the connection process for an CloudHSM key store, KMS finds the CloudHSM
// cluster that is associated with the custom key store, creates the connection
// infrastructure, connects to the cluster, logs into the CloudHSM client as the
// kmsuser CU, and rotates its password.
//
// To connect an CloudHSM key store, its associated CloudHSM cluster must have at
// least one active HSM. To get the number of active HSMs in a cluster, use the [DescribeClusters]
// operation. To add HSMs to the cluster, use the [CreateHsm]operation. Also, the [kmsuser crypto user]kmsuser
// (CU) must not be logged into the cluster. This prevents KMS from using this
// account to log in.
//
// If you are having trouble connecting or disconnecting a CloudHSM key store, see [Troubleshooting an CloudHSM key store]
// in the Key Management Service Developer Guide.
//
// # External key store
//
// When you connect an external key store that uses public endpoint connectivity,
// KMS tests its ability to communicate with your external key manager by sending a
// request via the external key store proxy.
//
// When you connect to an external key store that uses VPC endpoint service
// connectivity, KMS establishes the networking elements that it needs to
// communicate with your external key manager via the external key store proxy.
// This includes creating an interface endpoint to the VPC endpoint service and a
// private hosted zone for traffic between KMS and the VPC endpoint service.
//
// To connect an external key store, KMS must be able to connect to the external
// key store proxy, the external key store proxy must be able to communicate with
// your external key manager, and the external key manager must be available for
// cryptographic operations.
//
// If you are having trouble connecting or disconnecting an external key store,
// see [Troubleshooting an external key store]in the Key Management Service Developer Guide.
//
// Cross-account use: No. You cannot perform this operation on a custom key store
// in a different Amazon Web Services account.
//
// Required permissions: [kms:ConnectCustomKeyStore] (IAM policy)
//
// # Related operations
//
// # CreateCustomKeyStore
//
// # DeleteCustomKeyStore
//
// # DescribeCustomKeyStores
//
// # DisconnectCustomKeyStore
//
// # UpdateCustomKeyStore
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [DescribeClusters]: https://docs.aws.amazon.com/cloudhsm/latest/APIReference/API_DescribeClusters.html
// [kmsuser crypto user]: https://docs.aws.amazon.com/kms/latest/developerguide/keystore-cloudhsm.html#concept-kmsuser
// [Troubleshooting an CloudHSM key store]: https://docs.aws.amazon.com/kms/latest/developerguide/fix-keystore.html
// [CreateHsm]: https://docs.aws.amazon.com/cloudhsm/latest/APIReference/API_CreateHsm.html
// [kms:ConnectCustomKeyStore]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [Troubleshooting an external key store]: https://docs.aws.amazon.com/kms/latest/developerguide/xks-troubleshooting.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [custom key store]: https://docs.aws.amazon.com/kms/latest/developerguide/key-store-overview.html
func kms_ConnectCustomKeyStore(cfg aws.Config, client *kms.Client) {
	input := &kms.ConnectCustomKeyStoreInput{
		// CustomKeyStoreId: *string, // Required
	}

	if len(_kmsCustomKeyStoreId) > 0 {
		input.CustomKeyStoreId = aws.String(_kmsCustomKeyStoreId)
	}

	if resp, err := client.ConnectCustomKeyStore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a friendly name for a KMS key.
// Adding, deleting, or updating an alias can allow or deny permission to the KMS
// key. For details, see [ABAC for KMS]in the Key Management Service Developer Guide.
//
// You can use an alias to identify a KMS key in the KMS console, in the DescribeKey
// operation and in [cryptographic operations], such as Encrypt and GenerateDataKey. You can also change the KMS key that's
// associated with the alias (UpdateAlias ) or delete the alias (DeleteAlias ) at any time. These
// operations don't affect the underlying KMS key.
//
// You can associate the alias with any customer managed key in the same Amazon
// Web Services Region. Each alias is associated with only one KMS key at a time,
// but a KMS key can have multiple aliases. A valid KMS key is required. You can't
// create an alias without a KMS key.
//
// The alias must be unique in the account and Region, but you can have aliases
// with the same name in different Regions. For detailed information about aliases,
// see [Aliases in KMS]in the Key Management Service Developer Guide.
//
// This operation does not return a response. To get the alias that you created,
// use the ListAliasesoperation.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: No. You cannot perform this operation on an alias in a
// different Amazon Web Services account.
//
// # Required permissions
//
// [kms:CreateAlias]
// - on the alias (IAM policy).
//
// [kms:CreateAlias]
// - on the KMS key (key policy).
//
// For details, see [Controlling access to aliases] in the Key Management Service Developer Guide.
//
// Related operations:
//
// # DeleteAlias
//
// # ListAliases
//
// # UpdateAlias
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [cryptographic operations]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-cryptography.html#cryptographic-operations
// [kms:CreateAlias]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [Aliases in KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-alias.html
// [ABAC for KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/abac.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [Controlling access to aliases]: https://docs.aws.amazon.com/kms/latest/developerguide/alias-access.html
func kms_CreateAlias(cfg aws.Config, client *kms.Client) {
	input := &kms.CreateAliasInput{
		// AliasName: *string, // Required
		// TargetKeyId: *string, // Required
	}

	if len(_kmsAliasName) > 0 {
		input.AliasName = aws.String(_kmsAliasName)
	}
	if len(_kmsTargetKeyId) > 0 {
		input.TargetKeyId = aws.String(_kmsTargetKeyId)
	}

	if resp, err := client.CreateAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a [custom key store] backed by a key store that you own and manage. When you use a KMS
// key in a custom key store for a cryptographic operation, the cryptographic
// operation is actually performed in your key store using your keys. KMS supports [CloudHSM key stores]
// backed by an [CloudHSM cluster]and [external key stores] backed by an external key store proxy and external key
// manager outside of Amazon Web Services.
//
// This operation is part of the custom key stores feature in KMS, which combines
// the convenience and extensive integration of KMS with the isolation and control
// of a key store that you own and manage.
//
// Before you create the custom key store, the required elements must be in place
// and operational. We recommend that you use the test tools that KMS provides to
// verify the configuration your external key store proxy. For details about the
// required elements and verification tests, see [Assemble the prerequisites (for CloudHSM key stores)]or [Assemble the prerequisites (for external key stores)] in the Key Management Service
// Developer Guide.
//
// To create a custom key store, use the following parameters.
//
// - To create an CloudHSM key store, specify the CustomKeyStoreName ,
// CloudHsmClusterId , KeyStorePassword , and TrustAnchorCertificate . The
// CustomKeyStoreType parameter is optional for CloudHSM key stores. If you
// include it, set it to the default value, AWS_CLOUDHSM . For help with
// failures, see [Troubleshooting an CloudHSM key store]in the Key Management Service Developer Guide.
//
// - To create an external key store, specify the CustomKeyStoreName and a
// CustomKeyStoreType of EXTERNAL_KEY_STORE . Also, specify values for
// XksProxyConnectivity , XksProxyAuthenticationCredential , XksProxyUriEndpoint
// , and XksProxyUriPath . If your XksProxyConnectivity value is
// VPC_ENDPOINT_SERVICE , specify the XksProxyVpcEndpointServiceName parameter.
// For help with failures, see [Troubleshooting an external key store]in the Key Management Service Developer Guide.
//
// For external key stores:
//
// Some external key managers provide a simpler method for creating an external
// key store. For details, see your external key manager documentation.
//
// When creating an external key store in the KMS console, you can upload a
// JSON-based proxy configuration file with the desired values. You cannot use a
// proxy configuration with the CreateCustomKeyStore operation. However, you can
// use the values in the file to help you determine the correct values for the
// CreateCustomKeyStore parameters.
//
// When the operation completes successfully, it returns the ID of the new custom
// key store. Before you can use your new custom key store, you need to use the ConnectCustomKeyStore
// operation to connect a new CloudHSM key store to its CloudHSM cluster, or to
// connect a new external key store to the external key store proxy for your
// external key manager. Even if you are not going to use your custom key store
// immediately, you might want to connect it to verify that all settings are
// correct and then disconnect it until you are ready to use it.
//
// Cross-account use: No. You cannot perform this operation on a custom key store
// in a different Amazon Web Services account.
//
// Required permissions: [kms:CreateCustomKeyStore] (IAM policy).
//
// Related operations:
//
// # ConnectCustomKeyStore
//
// # DeleteCustomKeyStore
//
// # DescribeCustomKeyStores
//
// # DisconnectCustomKeyStore
//
// # UpdateCustomKeyStore
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [CloudHSM key stores]: https://docs.aws.amazon.com/kms/latest/developerguide/keystore-cloudhsm.html
// [CloudHSM cluster]: https://docs.aws.amazon.com/cloudhsm/latest/userguide/clusters.html
// [external key stores]: https://docs.aws.amazon.com/kms/latest/developerguide/keystore-external.html
// [Troubleshooting an CloudHSM key store]: https://docs.aws.amazon.com/kms/latest/developerguide/fix-keystore.html
// [Assemble the prerequisites (for CloudHSM key stores)]: https://docs.aws.amazon.com/kms/latest/developerguide/create-keystore.html#before-keystore
// [Assemble the prerequisites (for external key stores)]: https://docs.aws.amazon.com/kms/latest/developerguide/create-xks-keystore.html#xks-requirements
// [Troubleshooting an external key store]: https://docs.aws.amazon.com/kms/latest/developerguide/xks-troubleshooting.html
// [kms:CreateCustomKeyStore]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [custom key store]: https://docs.aws.amazon.com/kms/latest/developerguide/key-store-overview.html
func kms_CreateCustomKeyStore(cfg aws.Config, client *kms.Client) {
	input := &kms.CreateCustomKeyStoreInput{
		// CustomKeyStoreName: *string, // Required
	}

	if len(_kmsCustomKeyStoreName) > 0 {
		input.CustomKeyStoreName = aws.String(_kmsCustomKeyStoreName)
	}
	if len(_kmsCloudHsmClusterId) > 0 {
		input.CloudHsmClusterId = aws.String(_kmsCloudHsmClusterId)
	}
	if len(_kmsCustomKeyStoreType) > 0 {
		if err := assignInputField(input, "CustomKeyStoreType", _kmsCustomKeyStoreType); err != nil {
			log.Errorf("invalid --custom-key-store-type: %s", err.Error())
			return
		}
	}
	if len(_kmsKeyStorePassword) > 0 {
		input.KeyStorePassword = aws.String(_kmsKeyStorePassword)
	}
	if len(_kmsTrustAnchorCertificate) > 0 {
		input.TrustAnchorCertificate = aws.String(_kmsTrustAnchorCertificate)
	}
	if len(_kmsXksProxyAuthenticationCredential) > 0 {
		if err := assignInputField(input, "XksProxyAuthenticationCredential", _kmsXksProxyAuthenticationCredential); err != nil {
			log.Errorf("invalid --xks-proxy-authentication-credential: %s", err.Error())
			return
		}
	}
	if len(_kmsXksProxyConnectivity) > 0 {
		if err := assignInputField(input, "XksProxyConnectivity", _kmsXksProxyConnectivity); err != nil {
			log.Errorf("invalid --xks-proxy-connectivity: %s", err.Error())
			return
		}
	}
	if len(_kmsXksProxyUriEndpoint) > 0 {
		input.XksProxyUriEndpoint = aws.String(_kmsXksProxyUriEndpoint)
	}
	if len(_kmsXksProxyUriPath) > 0 {
		input.XksProxyUriPath = aws.String(_kmsXksProxyUriPath)
	}
	if len(_kmsXksProxyVpcEndpointServiceName) > 0 {
		input.XksProxyVpcEndpointServiceName = aws.String(_kmsXksProxyVpcEndpointServiceName)
	}
	if len(_kmsXksProxyVpcEndpointServiceOwner) > 0 {
		input.XksProxyVpcEndpointServiceOwner = aws.String(_kmsXksProxyVpcEndpointServiceOwner)
	}

	if resp, err := client.CreateCustomKeyStore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a grant to a KMS key.
// A grant is a policy instrument that allows Amazon Web Services principals to
// use KMS keys in cryptographic operations. It also can allow them to view a KMS
// key (DescribeKey ) and create and manage grants. When authorizing access to a KMS key,
// grants are considered along with key policies and IAM policies. Grants are often
// used for temporary permissions because you can create one, use its permissions,
// and delete it without changing your key policies or IAM policies.
//
// For detailed information about grants, including grant terminology, see [Grants in KMS] in the
// Key Management Service Developer Guide . For examples of creating grants in
// several programming languages, see [Use CreateGrant with an Amazon Web Services SDK or CLI].
//
// The CreateGrant operation returns a GrantToken and a GrantId .
//
// - When you create, retire, or revoke a grant, there might be a brief delay,
// usually less than five minutes, until the grant is available throughout KMS.
// This state is known as eventual consistency. Once the grant has achieved
// eventual consistency, the grantee principal can use the permissions in the grant
// without identifying the grant.
//
// # However, to use the permissions in the grant immediately, use the GrantToken
//
// that CreateGrant returns. For details, see [Using a grant token]in the Key Management Service
// Developer Guide .
//
// - The CreateGrant operation also returns a GrantId . You can use the GrantId
// and a key identifier to identify the grant in the RetireGrantand RevokeGrantoperations. To find the
// grant ID, use the ListGrantsor ListRetirableGrantsoperations.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: Yes. To perform this operation on a KMS key in a different
// Amazon Web Services account, specify the key ARN in the value of the KeyId
// parameter.
//
// Required permissions: [kms:CreateGrant] (key policy)
//
// Related operations:
//
// # ListGrants
//
// # ListRetirableGrants
//
// # RetireGrant
//
// # RevokeGrant
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [Grants in KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/grants.html
// [kms:CreateGrant]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [Use CreateGrant with an Amazon Web Services SDK or CLI]: https://docs.aws.amazon.com/kms/latest/developerguide/example_kms_CreateGrant_section.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [Using a grant token]: https://docs.aws.amazon.com/kms/latest/developerguide/using-grant-token.html
func kms_CreateGrant(cfg aws.Config, client *kms.Client) {
	input := &kms.CreateGrantInput{
		// GranteePrincipal: *string, // Required
		// KeyId: *string, // Required
		// Operations: []types.GrantOperation, // Required
	}

	if len(_kmsGranteePrincipal) > 0 {
		input.GranteePrincipal = aws.String(_kmsGranteePrincipal)
	}
	if len(_kmsKeyId) > 0 {
		input.KeyId = aws.String(_kmsKeyId)
	}
	if len(_kmsOperations) > 0 {
		if err := assignInputField(input, "Operations", _kmsOperations); err != nil {
			log.Errorf("invalid --operations: %s", err.Error())
			return
		}
	}
	if len(_kmsConstraints) > 0 {
		if err := assignInputField(input, "Constraints", _kmsConstraints); err != nil {
			log.Errorf("invalid --constraints: %s", err.Error())
			return
		}
	}
	if len(_kmsDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _kmsDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}
	if len(_kmsGrantTokens) > 0 {
		input.GrantTokens = append([]string(nil), _kmsGrantTokens...)
	}
	if len(_kmsName) > 0 {
		input.Name = aws.String(_kmsName)
	}
	if len(_kmsRetiringPrincipal) > 0 {
		input.RetiringPrincipal = aws.String(_kmsRetiringPrincipal)
	}

	if resp, err := client.CreateGrant(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a unique customer managed [KMS key] in your Amazon Web Services account and
// Region. You can use a KMS key in cryptographic operations, such as encryption
// and signing. Some Amazon Web Services services let you use KMS keys that you
// create and manage to protect your service resources.
//
// A KMS key is a logical representation of a cryptographic key. In addition to
// the key material used in cryptographic operations, a KMS key includes metadata,
// such as the key ID, key policy, creation date, description, and key state.
//
// Use the parameters of CreateKey to specify the type of KMS key, the source of
// its key material, its key policy, description, tags, and other properties.
//
// KMS has replaced the term customer master key (CMK) with Key Management Service
// key and KMS key. The concept has not changed. To prevent breaking changes, KMS
// is keeping some variations of this term.
//
// To create different types of KMS keys, use the following guidance:
//
// Symmetric encryption KMS key By default, CreateKey creates a symmetric
// encryption KMS key with key material that KMS generates. This is the basic and
// most widely used type of KMS key, and provides the best performance.
//
// To create a symmetric encryption KMS key, you don't need to specify any
// parameters. The default value for KeySpec , SYMMETRIC_DEFAULT , the default
// value for KeyUsage , ENCRYPT_DECRYPT , and the default value for Origin ,
// AWS_KMS , create a symmetric encryption KMS key with KMS key material.
//
// If you need a key for basic encryption and decryption or you are creating a KMS
// key to protect your resources in an Amazon Web Services service, create a
// symmetric encryption KMS key. The key material in a symmetric encryption key
// never leaves KMS unencrypted. You can use a symmetric encryption KMS key to
// encrypt and decrypt data up to 4,096 bytes, but they are typically used to
// generate data keys and data keys pairs. For details, see GenerateDataKeyand GenerateDataKeyPair.
//
// Asymmetric KMS keys To create an asymmetric KMS key, use the KeySpec parameter
// to specify the type of key material in the KMS key. Then, use the KeyUsage
// parameter to determine whether the KMS key will be used to encrypt and decrypt
// or sign and verify. You can't change these properties after the KMS key is
// created.
//
// Asymmetric KMS keys contain an RSA key pair, Elliptic Curve (ECC) key pair,
// ML-DSA key pair or an SM2 key pair (China Regions only). The private key in an
// asymmetric KMS key never leaves KMS unencrypted. However, you can use the GetPublicKey
// operation to download the public key so it can be used outside of KMS. Each KMS
// key can have only one key usage. KMS keys with RSA key pairs can be used to
// encrypt and decrypt data or sign and verify messages (but not both). KMS keys
// with NIST-standard ECC key pairs can be used to sign and verify messages or
// derive shared secrets (but not both). KMS keys with ECC_SECG_P256K1 can be used
// only to sign and verify messages. KMS keys with ML-DSA key pairs can be used to
// sign and verify messages. KMS keys with SM2 key pairs (China Regions only) can
// be used to either encrypt and decrypt data, sign and verify messages, or derive
// shared secrets (you must choose one key usage type). For information about
// asymmetric KMS keys, see [Asymmetric KMS keys]in the Key Management Service Developer Guide.
//
// HMAC KMS key To create an HMAC KMS key, set the KeySpec parameter to a key spec
// value for HMAC KMS keys. Then set the KeyUsage parameter to GENERATE_VERIFY_MAC
// . You must set the key usage even though GENERATE_VERIFY_MAC is the only valid
// key usage value for HMAC KMS keys. You can't change these properties after the
// KMS key is created.
//
// HMAC KMS keys are symmetric keys that never leave KMS unencrypted. You can use
// HMAC keys to generate (GenerateMac ) and verify (VerifyMac ) HMAC codes for messages up to 4096
// bytes.
//
// Multi-Region primary keys To create a multi-Region primary key in the local
// Amazon Web Services Region, use the MultiRegion parameter with a value of True .
// To create a multi-Region replica key, that is, a KMS key with the same key ID
// and key material as a primary key, but in a different Amazon Web Services
// Region, use the ReplicateKeyoperation. To change a replica key to a primary key, and its
// primary key to a replica key, use the UpdatePrimaryRegionoperation.
//
// You can create multi-Region KMS keys for all supported KMS key types: symmetric
// encryption KMS keys, HMAC KMS keys, asymmetric encryption KMS keys, and
// asymmetric signing KMS keys. You can also create multi-Region keys with imported
// key material. However, you can't create multi-Region keys in a custom key store.
//
// This operation supports multi-Region keys, an KMS feature that lets you create
// multiple interoperable KMS keys in different Amazon Web Services Regions.
// Because these KMS keys have the same key ID, key material, and other metadata,
// you can use them interchangeably to encrypt data in one Amazon Web Services
// Region and decrypt it in a different Amazon Web Services Region without
// re-encrypting the data or making a cross-Region call. For more information about
// multi-Region keys, see [Multi-Region keys in KMS]in the Key Management Service Developer Guide.
//
// Imported key material To import your own key material into a KMS key, begin by
// creating a KMS key with no key material. To do this, use the Origin parameter
// of CreateKey with a value of EXTERNAL . Next, use GetParametersForImport operation to get a public
// key and import token. Use the wrapping public key to encrypt your key material.
// Then, use ImportKeyMaterialwith your import token to import the key material. For step-by-step
// instructions, see [Importing Key Material]in the Key Management Service Developer Guide .
//
// You can import key material into KMS keys of all supported KMS key types:
// symmetric encryption KMS keys, HMAC KMS keys, asymmetric encryption KMS keys,
// and asymmetric signing KMS keys. You can also create multi-Region keys with
// imported key material. However, you can't import key material into a KMS key in
// a custom key store.
//
// To create a multi-Region primary key with imported key material, use the Origin
// parameter of CreateKey with a value of EXTERNAL and the MultiRegion parameter
// with a value of True . To create replicas of the multi-Region primary key, use
// the ReplicateKeyoperation. For instructions, see [Importing key material step 1]. For more information about multi-Region
// keys, see [Multi-Region keys in KMS]in the Key Management Service Developer Guide.
//
// Custom key store A [custom key store] lets you protect your Amazon Web Services resources using
// keys in a backing key store that you own and manage. When you request a
// cryptographic operation with a KMS key in a custom key store, the operation is
// performed in the backing key store using its cryptographic keys.
//
// KMS supports [CloudHSM key stores] backed by an CloudHSM cluster and [external key stores] backed by an external key
// manager outside of Amazon Web Services. When you create a KMS key in an CloudHSM
// key store, KMS generates an encryption key in the CloudHSM cluster and
// associates it with the KMS key. When you create a KMS key in an external key
// store, you specify an existing encryption key in the external key manager.
//
// Some external key managers provide a simpler method for creating a KMS key in
// an external key store. For details, see your external key manager documentation.
//
// Before you create a KMS key in a custom key store, the ConnectionState of the
// key store must be CONNECTED . To connect the custom key store, use the ConnectCustomKeyStore
// operation. To find the ConnectionState , use the DescribeCustomKeyStores operation.
//
// To create a KMS key in a custom key store, use the CustomKeyStoreId . Use the
// default KeySpec value, SYMMETRIC_DEFAULT , and the default KeyUsage value,
// ENCRYPT_DECRYPT to create a symmetric encryption key. No other key type is
// supported in a custom key store.
//
// To create a KMS key in an [CloudHSM key store], use the Origin parameter with a value of
// AWS_CLOUDHSM . The CloudHSM cluster that is associated with the custom key store
// must have at least two active HSMs in different Availability Zones in the Amazon
// Web Services Region.
//
// To create a KMS key in an [external key store], use the Origin parameter with a value of
// EXTERNAL_KEY_STORE and an XksKeyId parameter that identifies an existing
// external key.
//
// Some external key managers provide a simpler method for creating a KMS key in
// an external key store. For details, see your external key manager documentation.
//
// Cross-account use: No. You cannot use this operation to create a KMS key in a
// different Amazon Web Services account.
//
// Required permissions: [kms:CreateKey] (IAM policy). To use the Tags parameter, [kms:TagResource] (IAM policy).
// For examples and information about related permissions, see [Allow a user to create KMS keys]in the Key
// Management Service Developer Guide.
//
// Related operations:
//
// # DescribeKey
//
// # ListKeys
//
// # ScheduleKeyDeletion
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [CloudHSM key stores]: https://docs.aws.amazon.com/kms/latest/developerguide/keystore-cloudhsm.html
// [external key store]: https://docs.aws.amazon.com/kms/latest/developerguide/create-xks-keys.html
// [external key stores]: https://docs.aws.amazon.com/kms/latest/developerguide/keystore-external.html
// [Asymmetric KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html
// [Multi-Region keys in KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html
// [Importing key material step 1]: https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys-create-cmk.html
// [Allow a user to create KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/customer-managed-policies.html#iam-policy-example-create-key
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [kms:TagResource]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [CloudHSM key store]: https://docs.aws.amazon.com/kms/latest/developerguide/create-cmk-keystore.html
// [kms:CreateKey]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [Importing Key Material]: https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html
// [custom key store]: https://docs.aws.amazon.com/kms/latest/developerguide/key-store-overview.html
// [KMS key]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#kms-keys
func kms_CreateKey(cfg aws.Config, client *kms.Client) {
	input := &kms.CreateKeyInput{}

	if len(_kmsBypassPolicyLockoutSafetyCheck) > 0 {
		if err := assignInputField(input, "BypassPolicyLockoutSafetyCheck", _kmsBypassPolicyLockoutSafetyCheck); err != nil {
			log.Errorf("invalid --bypass-policy-lockout-safety-check: %s", err.Error())
			return
		}
	}
	if len(_kmsCustomKeyStoreId) > 0 {
		input.CustomKeyStoreId = aws.String(_kmsCustomKeyStoreId)
	}
	if len(_kmsCustomerMasterKeySpec) > 0 {
		if err := assignInputField(input, "CustomerMasterKeySpec", _kmsCustomerMasterKeySpec); err != nil {
			log.Errorf("invalid --customer-master-key-spec: %s", err.Error())
			return
		}
	}
	if len(_kmsDescription) > 0 {
		input.Description = aws.String(_kmsDescription)
	}
	if len(_kmsKeySpec) > 0 {
		if err := assignInputField(input, "KeySpec", _kmsKeySpec); err != nil {
			log.Errorf("invalid --key-spec: %s", err.Error())
			return
		}
	}
	if len(_kmsKeyUsage) > 0 {
		if err := assignInputField(input, "KeyUsage", _kmsKeyUsage); err != nil {
			log.Errorf("invalid --key-usage: %s", err.Error())
			return
		}
	}
	if len(_kmsMultiRegion) > 0 {
		if err := assignInputField(input, "MultiRegion", _kmsMultiRegion); err != nil {
			log.Errorf("invalid --multi-region: %s", err.Error())
			return
		}
	}
	if len(_kmsOrigin) > 0 {
		if err := assignInputField(input, "Origin", _kmsOrigin); err != nil {
			log.Errorf("invalid --origin: %s", err.Error())
			return
		}
	}
	if len(_kmsPolicy) > 0 {
		input.Policy = aws.String(_kmsPolicy)
	}
	if len(_kmsTags) > 0 {
		if err := assignInputField(input, "Tags", _kmsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_kmsXksKeyId) > 0 {
		input.XksKeyId = aws.String(_kmsXksKeyId)
	}

	if resp, err := client.CreateKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Decrypts ciphertext that was encrypted by a KMS key using any of the following
// operations:
//
// # Encrypt
//
// # GenerateDataKey
//
// # GenerateDataKeyPair
//
// # GenerateDataKeyWithoutPlaintext
//
// # GenerateDataKeyPairWithoutPlaintext
//
// You can use this operation to decrypt ciphertext that was encrypted under a
// symmetric encryption KMS key or an asymmetric encryption KMS key. When the KMS
// key is asymmetric, you must specify the KMS key and the encryption algorithm
// that was used to encrypt the ciphertext. For information about asymmetric KMS
// keys, see [Asymmetric KMS keys]in the Key Management Service Developer Guide.
//
// The Decrypt operation also decrypts ciphertext that was encrypted outside of
// KMS by the public key in an KMS asymmetric KMS key. However, it cannot decrypt
// symmetric ciphertext produced by other libraries, such as the [Amazon Web Services Encryption SDK]or [Amazon S3 client-side encryption]. These
// libraries return a ciphertext format that is incompatible with KMS.
//
// If the ciphertext was encrypted under a symmetric encryption KMS key, the KeyId
// parameter is optional. KMS can get this information from metadata that it adds
// to the symmetric ciphertext blob. This feature adds durability to your
// implementation by ensuring that authorized users can decrypt ciphertext decades
// after it was encrypted, even if they've lost track of the key ID. However,
// specifying the KMS key is always recommended as a best practice. When you use
// the KeyId parameter to specify a KMS key, KMS only uses the KMS key you
// specify. If the ciphertext was encrypted under a different KMS key, the Decrypt
// operation fails. This practice ensures that you use the KMS key that you intend.
//
// Whenever possible, use key policies to give users permission to call the Decrypt
// operation on a particular KMS key, instead of using IAM policies. Otherwise, you
// might create an IAM policy that gives the user Decrypt permission on all KMS
// keys. This user could decrypt ciphertext that was encrypted by KMS keys in other
// accounts if the key policy for the cross-account KMS key permits it. If you must
// use an IAM policy for Decrypt permissions, limit the user to particular KMS
// keys or particular trusted accounts. For details, see [Best practices for IAM policies]in the Key Management
// Service Developer Guide.
//
// Decrypt also supports [Amazon Web Services Nitro Enclaves] and NitroTPM, which provide attested environments in
// Amazon EC2. To call Decrypt for a Nitro enclave or NitroTPM, use the [Amazon Web Services Nitro Enclaves SDK] or any
// Amazon Web Services SDK. Use the Recipient parameter to provide the attestation
// document for the attested environment. Instead of the plaintext data, the
// response includes the plaintext data encrypted with the public key from the
// attestation document ( CiphertextForRecipient ). For information about the
// interaction between KMS and Amazon Web Services Nitro Enclaves or Amazon Web
// Services NitroTPM, see [Cryptographic attestation support in KMS]in the Key Management Service Developer Guide.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: Yes. If you use the KeyId parameter to identify a KMS key in
// a different Amazon Web Services account, specify the key ARN or the alias ARN of
// the KMS key.
//
// Required permissions: [kms:Decrypt] (key policy)
//
// Related operations:
//
// # Encrypt
//
// # GenerateDataKey
//
// # GenerateDataKeyPair
//
// # ReEncrypt
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Amazon Web Services Encryption SDK]: https://docs.aws.amazon.com/encryption-sdk/latest/developer-guide/
// [Cryptographic attestation support in KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/cryptographic-attestation.html
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [kms:Decrypt]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [Asymmetric KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html
// [Amazon Web Services Nitro Enclaves]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/nitro-enclave.html
// [Amazon S3 client-side encryption]: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingClientSideEncryption.html
// [Best practices for IAM policies]: https://docs.aws.amazon.com/kms/latest/developerguide/iam-policies.html#iam-policies-best-practices
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [Amazon Web Services Nitro Enclaves SDK]: https://docs.aws.amazon.com/enclaves/latest/user/developing-applications.html#sdk
func kms_Decrypt(cfg aws.Config, client *kms.Client) {
	input := &kms.DecryptInput{}

	if len(_kmsCiphertextBlob) > 0 {
		if err := assignInputField(input, "CiphertextBlob", _kmsCiphertextBlob); err != nil {
			log.Errorf("invalid --ciphertext-blob: %s", err.Error())
			return
		}
	}
	if len(_kmsDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _kmsDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}
	if len(_kmsDryRunModifiers) > 0 {
		if err := assignInputField(input, "DryRunModifiers", _kmsDryRunModifiers); err != nil {
			log.Errorf("invalid --dry-run-modifiers: %s", err.Error())
			return
		}
	}
	if len(_kmsEncryptionAlgorithm) > 0 {
		if err := assignInputField(input, "EncryptionAlgorithm", _kmsEncryptionAlgorithm); err != nil {
			log.Errorf("invalid --encryption-algorithm: %s", err.Error())
			return
		}
	}
	if len(_kmsEncryptionContext) > 0 {
		if err := assignInputField(input, "EncryptionContext", _kmsEncryptionContext); err != nil {
			log.Errorf("invalid --encryption-context: %s", err.Error())
			return
		}
	}
	if len(_kmsGrantTokens) > 0 {
		input.GrantTokens = append([]string(nil), _kmsGrantTokens...)
	}
	if len(_kmsKeyId) > 0 {
		input.KeyId = aws.String(_kmsKeyId)
	}
	if len(_kmsRecipient) > 0 {
		if err := assignInputField(input, "Recipient", _kmsRecipient); err != nil {
			log.Errorf("invalid --recipient: %s", err.Error())
			return
		}
	}

	if resp, err := client.Decrypt(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified alias.
// Adding, deleting, or updating an alias can allow or deny permission to the KMS
// key. For details, see [ABAC for KMS]in the Key Management Service Developer Guide.
//
// Because an alias is not a property of a KMS key, you can delete and change the
// aliases of a KMS key without affecting the KMS key. Also, aliases do not appear
// in the response from the DescribeKeyoperation. To get the aliases of all KMS keys, use the ListAliases
// operation.
//
// Each KMS key can have multiple aliases. To change the alias of a KMS key, use DeleteAlias
// to delete the current alias and CreateAliasto create a new alias. To associate an existing
// alias with a different KMS key, call UpdateAlias.
//
// Cross-account use: No. You cannot perform this operation on an alias in a
// different Amazon Web Services account.
//
// # Required permissions
//
// [kms:DeleteAlias]
// - on the alias (IAM policy).
//
// [kms:DeleteAlias]
// - on the KMS key (key policy).
//
// For details, see [Controlling access to aliases] in the Key Management Service Developer Guide.
//
// Related operations:
//
// # CreateAlias
//
// # ListAliases
//
// # UpdateAlias
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [ABAC for KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/abac.html
// [kms:DeleteAlias]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [Controlling access to aliases]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-alias.html#alias-access
func kms_DeleteAlias(cfg aws.Config, client *kms.Client) {
	input := &kms.DeleteAliasInput{
		// AliasName: *string, // Required
	}

	if len(_kmsAliasName) > 0 {
		input.AliasName = aws.String(_kmsAliasName)
	}

	if resp, err := client.DeleteAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a [custom key store]. This operation does not affect any backing elements of the custom
// key store. It does not delete the CloudHSM cluster that is associated with an
// CloudHSM key store, or affect any users or keys in the cluster. For an external
// key store, it does not affect the external key store proxy, external key
// manager, or any external keys.
//
// This operation is part of the custom key stores feature in KMS, which combines
// the convenience and extensive integration of KMS with the isolation and control
// of a key store that you own and manage.
//
// The custom key store that you delete cannot contain any [KMS keys]. Before deleting the
// key store, verify that you will never need to use any of the KMS keys in the key
// store for any [cryptographic operations]. Then, use ScheduleKeyDeletion to delete the KMS keys from the key store. After the
// required waiting period expires and all KMS keys are deleted from the custom key
// store, use DisconnectCustomKeyStoreto disconnect the key store from KMS. Then, you can delete the
// custom key store.
//
// For keys in an CloudHSM key store, the ScheduleKeyDeletion operation makes a
// best effort to delete the key material from the associated cluster. However, you
// might need to manually [delete the orphaned key material]from the cluster and its backups. KMS never creates,
// manages, or deletes cryptographic keys in the external key manager associated
// with an external key store. You must manage them using your external key manager
// tools.
//
// Instead of deleting the custom key store, consider using the DisconnectCustomKeyStore operation to
// disconnect the custom key store from its backing key store. While the key store
// is disconnected, you cannot create or use the KMS keys in the key store. But,
// you do not need to delete KMS keys and you can reconnect a disconnected custom
// key store at any time.
//
// If the operation succeeds, it returns a JSON object with no properties.
//
// Cross-account use: No. You cannot perform this operation on a custom key store
// in a different Amazon Web Services account.
//
// Required permissions: [kms:DeleteCustomKeyStore] (IAM policy)
//
// Related operations:
//
// # ConnectCustomKeyStore
//
// # CreateCustomKeyStore
//
// # DescribeCustomKeyStores
//
// # DisconnectCustomKeyStore
//
// # UpdateCustomKeyStore
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [delete the orphaned key material]: https://docs.aws.amazon.com/kms/latest/developerguide/fix-keystore.html#fix-keystore-orphaned-key
// [kms:DeleteCustomKeyStore]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [cryptographic operations]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-cryptography.html#cryptographic-operations
// [KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#kms_keys
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [custom key store]: https://docs.aws.amazon.com/kms/latest/developerguide/key-store-overview.html
func kms_DeleteCustomKeyStore(cfg aws.Config, client *kms.Client) {
	input := &kms.DeleteCustomKeyStoreInput{
		// CustomKeyStoreId: *string, // Required
	}

	if len(_kmsCustomKeyStoreId) > 0 {
		input.CustomKeyStoreId = aws.String(_kmsCustomKeyStoreId)
	}

	if resp, err := client.DeleteCustomKeyStore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes key material that was previously imported. This operation makes the
// specified KMS key temporarily unusable. To restore the usability of the KMS key,
// reimport the same key material. For more information about importing key
// material into KMS, see [Importing Key Material]in the Key Management Service Developer Guide.
//
// When the specified KMS key is in the PendingDeletion state, this operation does
// not change the KMS key's state. Otherwise, it changes the KMS key's state to
// PendingImport .
//
// # Considerations for multi-Region symmetric encryption keys
//
// - When you delete the key material of a primary Region key that is in
// PENDING_ROTATION or PENDING_MULTI_REGION_IMPORT_AND_ROTATION state, you'll
// also be deleting the key materials for the replica Region keys.
//
// - If you delete any key material of a replica Region key, the primary Region
// key and other replica Region keys remain unchanged.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: No. You cannot perform this operation on a KMS key in a
// different Amazon Web Services account.
//
// Required permissions: [kms:DeleteImportedKeyMaterial] (key policy)
//
// Related operations:
//
// # GetParametersForImport
//
// # ListKeyRotations
//
// # ImportKeyMaterial
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [kms:DeleteImportedKeyMaterial]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [Importing Key Material]: https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
func kms_DeleteImportedKeyMaterial(cfg aws.Config, client *kms.Client) {
	input := &kms.DeleteImportedKeyMaterialInput{
		// KeyId: *string, // Required
	}

	if len(_kmsKeyId) > 0 {
		input.KeyId = aws.String(_kmsKeyId)
	}
	if len(_kmsKeyMaterialId) > 0 {
		input.KeyMaterialId = aws.String(_kmsKeyMaterialId)
	}

	if resp, err := client.DeleteImportedKeyMaterial(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Derives a shared secret using a key agreement algorithm.
// You must use an asymmetric NIST-standard elliptic curve (ECC) or SM2 (China
// Regions only) KMS key pair with a KeyUsage value of KEY_AGREEMENT to call
// DeriveSharedSecret.
//
// DeriveSharedSecret uses the [Elliptic Curve Cryptography Cofactor Diffie-Hellman Primitive] (ECDH) to establish a key agreement between two
// peers by deriving a shared secret from their elliptic curve public-private key
// pairs. You can use the raw shared secret that DeriveSharedSecret returns to
// derive a symmetric key that can encrypt and decrypt data that is sent between
// the two peers, or that can generate and verify HMACs. KMS recommends that you
// follow [NIST recommendations for key derivation]when using the raw shared secret to derive a symmetric key.
//
// The following workflow demonstrates how to establish key agreement over an
// insecure communication channel using DeriveSharedSecret.
//
// - Alice calls CreateKeyto create an asymmetric KMS key pair with a KeyUsage value of
// KEY_AGREEMENT .
//
// # The asymmetric KMS key must use a NIST-standard elliptic curve (ECC) or SM2
//
// (China Regions only) key spec.
//
// - Bob creates an elliptic curve key pair.
//
// # Bob can call CreateKeyto create an asymmetric KMS key pair or generate a key pair
//
// outside of KMS. Bob's key pair must use the same NIST-standard elliptic curve
// (ECC) or SM2 (China Regions ony) curve as Alice.
//
// - Alice and Bob exchange their public keys through an insecure communication
// channel (like the internet).
//
// Use GetPublicKeyto download the public key of your asymmetric KMS key pair.
//
// # KMS strongly recommends verifying that the public key you receive came from the
//
// expected party before using it to derive a shared secret.
//
// - Alice calls DeriveSharedSecret.
//
// # KMS uses the private key from the KMS key pair generated in Step 1, Bob's
//
// public key, and the Elliptic Curve Cryptography Cofactor Diffie-Hellman
// Primitive to derive the shared secret. The private key in your KMS key pair
// never leaves KMS unencrypted. DeriveSharedSecret returns the raw shared secret.
//
// - Bob uses the Elliptic Curve Cryptography Cofactor Diffie-Hellman Primitive
// to calculate the same raw secret using his private key and Alice's public key.
//
// To derive a shared secret you must provide a key agreement algorithm, the
// private key of the caller's asymmetric NIST-standard elliptic curve or SM2
// (China Regions only) KMS key pair, and the public key from your peer's
// NIST-standard elliptic curve or SM2 (China Regions only) key pair. The public
// key can be from another asymmetric KMS key pair or from a key pair generated
// outside of KMS, but both key pairs must be on the same elliptic curve.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: Yes. To perform this operation with a KMS key in a different
// Amazon Web Services account, specify the key ARN or alias ARN in the value of
// the KeyId parameter.
//
// Required permissions: [kms:DeriveSharedSecret] (key policy)
//
// Related operations:
//
// # CreateKey
//
// # GetPublicKey
//
// # DescribeKey
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [kms:DeriveSharedSecret]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [Elliptic Curve Cryptography Cofactor Diffie-Hellman Primitive]: https://nvlpubs.nist.gov/nistpubs/SpecialPublications/NIST.SP.800-56Ar3.pdf#page=60
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [NIST recommendations for key derivation]: https://nvlpubs.nist.gov/nistpubs/SpecialPublications/NIST.SP.800-56Cr2.pdf
func kms_DeriveSharedSecret(cfg aws.Config, client *kms.Client) {
	input := &kms.DeriveSharedSecretInput{
		// KeyAgreementAlgorithm: types.KeyAgreementAlgorithmSpec, // Required
		// KeyId: *string, // Required
		// PublicKey: []byte, // Required
	}

	if len(_kmsKeyAgreementAlgorithm) > 0 {
		if err := assignInputField(input, "KeyAgreementAlgorithm", _kmsKeyAgreementAlgorithm); err != nil {
			log.Errorf("invalid --key-agreement-algorithm: %s", err.Error())
			return
		}
	}
	if len(_kmsKeyId) > 0 {
		input.KeyId = aws.String(_kmsKeyId)
	}
	if len(_kmsPublicKey) > 0 {
		if err := assignInputField(input, "PublicKey", _kmsPublicKey); err != nil {
			log.Errorf("invalid --public-key: %s", err.Error())
			return
		}
	}
	if len(_kmsDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _kmsDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}
	if len(_kmsGrantTokens) > 0 {
		input.GrantTokens = append([]string(nil), _kmsGrantTokens...)
	}
	if len(_kmsRecipient) > 0 {
		if err := assignInputField(input, "Recipient", _kmsRecipient); err != nil {
			log.Errorf("invalid --recipient: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeriveSharedSecret(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about [custom key stores] in the account and Region.
// This operation is part of the custom key stores feature in KMS, which combines
// the convenience and extensive integration of KMS with the isolation and control
// of a key store that you own and manage.
//
// By default, this operation returns information about all custom key stores in
// the account and Region. To get only information about a particular custom key
// store, use either the CustomKeyStoreName or CustomKeyStoreId parameter (but not
// both).
//
// To determine whether the custom key store is connected to its CloudHSM cluster
// or external key store proxy, use the ConnectionState element in the response.
// If an attempt to connect the custom key store failed, the ConnectionState value
// is FAILED and the ConnectionErrorCode element in the response indicates the
// cause of the failure. For help interpreting the ConnectionErrorCode , see CustomKeyStoresListEntry.
//
// Custom key stores have a DISCONNECTED connection state if the key store has
// never been connected or you used the DisconnectCustomKeyStoreoperation to disconnect it. Otherwise, the
// connection state is CONNECTED. If your custom key store connection state is
// CONNECTED but you are having trouble using it, verify that the backing store is
// active and available. For an CloudHSM key store, verify that the associated
// CloudHSM cluster is active and contains the minimum number of HSMs required for
// the operation, if any. For an external key store, verify that the external key
// store proxy and its associated external key manager are reachable and enabled.
//
// For help repairing your CloudHSM key store, see the [Troubleshooting CloudHSM key stores]. For help repairing your
// external key store, see the [Troubleshooting external key stores]. Both topics are in the Key Management Service
// Developer Guide.
//
// Cross-account use: No. You cannot perform this operation on a custom key store
// in a different Amazon Web Services account.
//
// Required permissions: [kms:DescribeCustomKeyStores] (IAM policy)
//
// Related operations:
//
// # ConnectCustomKeyStore
//
// # CreateCustomKeyStore
//
// # DeleteCustomKeyStore
//
// # DisconnectCustomKeyStore
//
// # UpdateCustomKeyStore
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [kms:DescribeCustomKeyStores]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [Troubleshooting CloudHSM key stores]: https://docs.aws.amazon.com/kms/latest/developerguide/fix-keystore.html
// [Troubleshooting external key stores]: https://docs.aws.amazon.com/kms/latest/developerguide/xks-troubleshooting.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [custom key stores]: https://docs.aws.amazon.com/kms/latest/developerguide/key-store-overview.html
func kms_DescribeCustomKeyStores(cfg aws.Config, client *kms.Client) {
	input := &kms.DescribeCustomKeyStoresInput{}

	if len(_kmsCustomKeyStoreId) > 0 {
		input.CustomKeyStoreId = aws.String(_kmsCustomKeyStoreId)
	}
	if len(_kmsCustomKeyStoreName) > 0 {
		input.CustomKeyStoreName = aws.String(_kmsCustomKeyStoreName)
	}
	if len(_kmsLimit) > 0 {
		if err := assignInputField(input, "Limit", _kmsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_kmsMarker) > 0 {
		input.Marker = aws.String(_kmsMarker)
	}

	if disablePaginator() {
		if resp, err := client.DescribeCustomKeyStores(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*kms.DescribeCustomKeyStoresOutput
	p := kms.NewDescribeCustomKeyStoresPaginator(client, input)
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

// Provides detailed information about a KMS key. You can run DescribeKey on a [customer managed key] or
// an [Amazon Web Services managed key].
//
// This detailed information includes the key ARN, creation date (and deletion
// date, if applicable), the key state, and the origin and expiration date (if any)
// of the key material. It includes fields, like KeySpec , that help you
// distinguish different types of KMS keys. It also displays the key usage
// (encryption, signing, or generating and verifying MACs) and the algorithms that
// the KMS key supports.
//
// For [multi-Region keys], DescribeKey displays the primary key and all related replica keys. For
// KMS keys in [CloudHSM key stores], it includes information about the key store, such as the key
// store ID and the CloudHSM cluster ID. For KMS keys in [external key stores], it includes the custom
// key store ID and the ID of the external key.
//
// DescribeKey does not return the following information:
//
// - Aliases associated with the KMS key. To get this information, use ListAliases.
//
// - Whether automatic key rotation is enabled on the KMS key. To get this
// information, use GetKeyRotationStatus. Also, some key states prevent a KMS key from being
// automatically rotated. For details, see [How key rotation works]in the Key Management Service
// Developer Guide.
//
// - Tags on the KMS key. To get this information, use ListResourceTags.
//
// - Key policies and grants on the KMS key. To get this information, use GetKeyPolicyand ListGrants.
//
// In general, DescribeKey is a non-mutating operation. It returns data about KMS
// keys, but doesn't change them. However, Amazon Web Services services use
// DescribeKey to create [Amazon Web Services managed keys] from a predefined Amazon Web Services alias with no key
// ID.
//
// Cross-account use: Yes. To perform this operation with a KMS key in a different
// Amazon Web Services account, specify the key ARN or alias ARN in the value of
// the KeyId parameter.
//
// Required permissions: [kms:DescribeKey] (key policy)
//
// Related operations:
//
// # GetKeyPolicy
//
// # GetKeyRotationStatus
//
// # ListAliases
//
// # ListGrants
//
// # ListKeys
//
// # ListResourceTags
//
// # ListRetirableGrants
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [CloudHSM key stores]: https://docs.aws.amazon.com/kms/latest/developerguide/keystore-cloudhsm.html
// [external key stores]: https://docs.aws.amazon.com/kms/latest/developerguide/keystore-external.html
// [How key rotation works]: https://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html#rotate-keys-how-it-works
// [kms:DescribeKey]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [multi-Region keys]: https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html
// [Amazon Web Services managed keys]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-key
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [Amazon Web Services managed key]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-key
// [customer managed key]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#customer-mgn-key
func kms_DescribeKey(cfg aws.Config, client *kms.Client) {
	input := &kms.DescribeKeyInput{
		// KeyId: *string, // Required
	}

	if len(_kmsKeyId) > 0 {
		input.KeyId = aws.String(_kmsKeyId)
	}
	if len(_kmsGrantTokens) > 0 {
		input.GrantTokens = append([]string(nil), _kmsGrantTokens...)
	}

	if resp, err := client.DescribeKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the state of a KMS key to disabled. This change temporarily prevents use
// of the KMS key for [cryptographic operations].
//
// The KMS key that you use for this operation must be in a compatible key state.
// For more information about how key state affects the use of a KMS key, see [Key states of KMS keys]in
// the Key Management Service Developer Guide .
//
// Cross-account use: No. You cannot perform this operation on a KMS key in a
// different Amazon Web Services account.
//
// Required permissions: [kms:DisableKey] (key policy)
//
// Related operations: EnableKey
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [cryptographic operations]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-cryptography.html#cryptographic-operations
// [kms:DisableKey]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
func kms_DisableKey(cfg aws.Config, client *kms.Client) {
	input := &kms.DisableKeyInput{
		// KeyId: *string, // Required
	}

	if len(_kmsKeyId) > 0 {
		input.KeyId = aws.String(_kmsKeyId)
	}

	if resp, err := client.DisableKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables [automatic rotation of the key material] of the specified symmetric encryption KMS key.
// Automatic key rotation is supported only on symmetric encryption KMS keys. You
// cannot enable automatic rotation of [asymmetric KMS keys], [HMAC KMS keys], KMS keys with [imported key material], or KMS keys in a [custom key store]. To
// enable or disable automatic rotation of a set of related [multi-Region keys], set the property on
// the primary key.
//
// You can enable (EnableKeyRotation ) and disable automatic rotation of the key material in [customer managed KMS keys]. Key
// material rotation of [Amazon Web Services managed KMS keys]is not configurable. KMS always rotates the key material
// for every year. Rotation of [Amazon Web Services owned KMS keys]varies.
//
// In May 2022, KMS changed the rotation schedule for Amazon Web Services managed
// keys from every three years to every year. For details, see EnableKeyRotation.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: No. You cannot perform this operation on a KMS key in a
// different Amazon Web Services account.
//
// Required permissions: [kms:DisableKeyRotation] (key policy)
//
// Related operations:
//
// # EnableKeyRotation
//
// # GetKeyRotationStatus
//
// # ListKeyRotations
//
// # RotateKeyOnDemand
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [imported key material]: https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [HMAC KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/hmac.html
// [Amazon Web Services managed KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-key
// [asymmetric KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html
// [customer managed KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#customer-mgn-key
// [Amazon Web Services owned KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-owned-key
// [kms:DisableKeyRotation]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [multi-Region keys]: https://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html#multi-region-rotate
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [custom key store]: https://docs.aws.amazon.com/kms/latest/developerguide/key-store-overview.html
// [automatic rotation of the key material]: https://docs.aws.amazon.com/kms/latest/developerguide/rotating-keys-enable-disable.html
func kms_DisableKeyRotation(cfg aws.Config, client *kms.Client) {
	input := &kms.DisableKeyRotationInput{
		// KeyId: *string, // Required
	}

	if len(_kmsKeyId) > 0 {
		input.KeyId = aws.String(_kmsKeyId)
	}

	if resp, err := client.DisableKeyRotation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disconnects the [custom key store] from its backing key store. This operation disconnects an
// CloudHSM key store from its associated CloudHSM cluster or disconnects an
// external key store from the external key store proxy that communicates with your
// external key manager.
//
// This operation is part of the custom key stores feature in KMS, which combines
// the convenience and extensive integration of KMS with the isolation and control
// of a key store that you own and manage.
//
// While a custom key store is disconnected, you can manage the custom key store
// and its KMS keys, but you cannot create or use its KMS keys. You can reconnect
// the custom key store at any time.
//
// While a custom key store is disconnected, all attempts to create KMS keys in
// the custom key store or to use existing KMS keys in [cryptographic operations]will fail. This action can
// prevent users from storing and accessing sensitive data.
//
// When you disconnect a custom key store, its ConnectionState changes to
// Disconnected . To find the connection state of a custom key store, use the DescribeCustomKeyStores
// operation. To reconnect a custom key store, use the ConnectCustomKeyStoreoperation.
//
// If the operation succeeds, it returns a JSON object with no properties.
//
// Cross-account use: No. You cannot perform this operation on a custom key store
// in a different Amazon Web Services account.
//
// Required permissions: [kms:DisconnectCustomKeyStore] (IAM policy)
//
// Related operations:
//
// # ConnectCustomKeyStore
//
// # CreateCustomKeyStore
//
// # DeleteCustomKeyStore
//
// # DescribeCustomKeyStores
//
// # UpdateCustomKeyStore
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [cryptographic operations]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-cryptography.html#cryptographic-operations
// [kms:DisconnectCustomKeyStore]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [custom key store]: https://docs.aws.amazon.com/kms/latest/developerguide/key-store-overview.html
func kms_DisconnectCustomKeyStore(cfg aws.Config, client *kms.Client) {
	input := &kms.DisconnectCustomKeyStoreInput{
		// CustomKeyStoreId: *string, // Required
	}

	if len(_kmsCustomKeyStoreId) > 0 {
		input.CustomKeyStoreId = aws.String(_kmsCustomKeyStoreId)
	}

	if resp, err := client.DisconnectCustomKeyStore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the key state of a KMS key to enabled. This allows you to use the KMS key
// for [cryptographic operations].
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: No. You cannot perform this operation on a KMS key in a
// different Amazon Web Services account.
//
// Required permissions: [kms:EnableKey] (key policy)
//
// Related operations: DisableKey
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [kms:EnableKey]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [cryptographic operations]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-cryptography.html#cryptographic-operations
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
func kms_EnableKey(cfg aws.Config, client *kms.Client) {
	input := &kms.EnableKeyInput{
		// KeyId: *string, // Required
	}

	if len(_kmsKeyId) > 0 {
		input.KeyId = aws.String(_kmsKeyId)
	}

	if resp, err := client.EnableKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables [automatic rotation of the key material] of the specified symmetric encryption KMS key.
// By default, when you enable automatic rotation of a [customer managed KMS key], KMS rotates the key
// material of the KMS key one year (approximately 365 days) from the enable date
// and every year thereafter. You can use the optional RotationPeriodInDays
// parameter to specify a custom rotation period when you enable key rotation, or
// you can use RotationPeriodInDays to modify the rotation period of a key that
// you previously enabled automatic key rotation on.
//
// You can monitor rotation of the key material for your KMS keys in CloudTrail
// and Amazon CloudWatch. To disable rotation of the key material in a customer
// managed KMS key, use the DisableKeyRotationoperation. You can use the GetKeyRotationStatus operation to identify any
// in progress rotations. You can use the ListKeyRotationsoperation to view the details of
// completed rotations.
//
// Automatic key rotation is supported only on symmetric encryption KMS keys. You
// cannot enable automatic rotation of [asymmetric KMS keys], [HMAC KMS keys], KMS keys with [imported key material], or KMS keys in a [custom key store]. To
// enable or disable automatic rotation of a set of related [multi-Region keys], set the property on
// the primary key.
//
// You cannot enable or disable automatic rotation of [Amazon Web Services managed KMS keys]. KMS always rotates the key
// material of Amazon Web Services managed keys every year. Rotation of [Amazon Web Services owned KMS keys]is managed
// by the Amazon Web Services service that owns the key.
//
// In May 2022, KMS changed the rotation schedule for Amazon Web Services managed
// keys from every three years (approximately 1,095 days) to every year
// (approximately 365 days).
//
// New Amazon Web Services managed keys are automatically rotated one year after
// they are created, and approximately every year thereafter.
//
// Existing Amazon Web Services managed keys are automatically rotated one year
// after their most recent rotation, and every year thereafter.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: No. You cannot perform this operation on a KMS key in a
// different Amazon Web Services account.
//
// Required permissions: [kms:EnableKeyRotation] (key policy)
//
// Related operations:
//
// # DisableKeyRotation
//
// # GetKeyRotationStatus
//
// # ListKeyRotations
//
// # RotateKeyOnDemand
//
// - You can perform on-demand (RotateKeyOnDemand ) rotation of the key material in customer
// managed KMS keys, regardless of whether or not automatic key rotation is
// enabled.
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [kms:EnableKeyRotation]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [imported key material]: https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [HMAC KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/hmac.html
// [Amazon Web Services managed KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-key
// [customer managed KMS key]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#customer-mgn-key
// [asymmetric KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html
// [Amazon Web Services owned KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-owned-key
// [multi-Region keys]: https://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html#multi-region-rotate
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [custom key store]: https://docs.aws.amazon.com/kms/latest/developerguide/key-store-overview.html
// [automatic rotation of the key material]: https://docs.aws.amazon.com/kms/latest/developerguide/rotating-keys-enable-disable.html
func kms_EnableKeyRotation(cfg aws.Config, client *kms.Client) {
	input := &kms.EnableKeyRotationInput{
		// KeyId: *string, // Required
	}

	if len(_kmsKeyId) > 0 {
		input.KeyId = aws.String(_kmsKeyId)
	}
	if len(_kmsRotationPeriodInDays) > 0 {
		if err := assignInputField(input, "RotationPeriodInDays", _kmsRotationPeriodInDays); err != nil {
			log.Errorf("invalid --rotation-period-in-days: %s", err.Error())
			return
		}
	}

	if resp, err := client.EnableKeyRotation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Encrypts plaintext of up to 4,096 bytes using a KMS key. You can use a
// symmetric or asymmetric KMS key with a KeyUsage of ENCRYPT_DECRYPT .
//
// You can use this operation to encrypt small amounts of arbitrary data, such as
// a personal identifier or database password, or other sensitive information. You
// don't need to use the Encrypt operation to encrypt a data key. The GenerateDataKey and GenerateDataKeyPair
// operations return a plaintext data key and an encrypted copy of that data key.
//
// If you use a symmetric encryption KMS key, you can use an encryption context to
// add additional security to your encryption operation. If you specify an
// EncryptionContext when encrypting data, you must specify the same encryption
// context (a case-sensitive exact match) when decrypting the data. Otherwise, the
// request to decrypt fails with an InvalidCiphertextException . For more
// information, see [Encryption Context]in the Key Management Service Developer Guide.
//
// If you specify an asymmetric KMS key, you must also specify the encryption
// algorithm. The algorithm must be compatible with the KMS key spec.
//
// When you use an asymmetric KMS key to encrypt or reencrypt data, be sure to
// record the KMS key and encryption algorithm that you choose. You will be
// required to provide the same KMS key and encryption algorithm when you decrypt
// the data. If the KMS key and algorithm do not match the values used to encrypt
// the data, the decrypt operation fails.
//
// You are not required to supply the key ID and encryption algorithm when you
// decrypt with symmetric encryption KMS keys because KMS stores this information
// in the ciphertext blob. KMS cannot store metadata in ciphertext generated with
// asymmetric keys. The standard format for asymmetric key ciphertext does not
// include configurable fields.
//
// The maximum size of the data that you can encrypt varies with the type of KMS
// key and the encryption algorithm that you choose.
//
// - Symmetric encryption KMS keys
//
// - SYMMETRIC_DEFAULT : 4096 bytes
//
// - RSA_2048
//
// - RSAES_OAEP_SHA_1 : 214 bytes
//
// - RSAES_OAEP_SHA_256 : 190 bytes
//
// - RSA_3072
//
// - RSAES_OAEP_SHA_1 : 342 bytes
//
// - RSAES_OAEP_SHA_256 : 318 bytes
//
// - RSA_4096
//
// - RSAES_OAEP_SHA_1 : 470 bytes
//
// - RSAES_OAEP_SHA_256 : 446 bytes
//
// - SM2PKE : 1024 bytes (China Regions only)
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: Yes. To perform this operation with a KMS key in a different
// Amazon Web Services account, specify the key ARN or alias ARN in the value of
// the KeyId parameter.
//
// Required permissions: [kms:Encrypt] (key policy)
//
// Related operations:
//
// # Decrypt
//
// # GenerateDataKey
//
// # GenerateDataKeyPair
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [Encryption Context]: https://docs.aws.amazon.com/kms/latest/developerguide/encrypt_context.html
// [kms:Encrypt]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
func kms_Encrypt(cfg aws.Config, client *kms.Client) {
	input := &kms.EncryptInput{
		// KeyId: *string, // Required
		// Plaintext: []byte, // Required
	}

	if len(_kmsKeyId) > 0 {
		input.KeyId = aws.String(_kmsKeyId)
	}
	if len(_kmsPlaintext) > 0 {
		if err := assignInputField(input, "Plaintext", _kmsPlaintext); err != nil {
			log.Errorf("invalid --plaintext: %s", err.Error())
			return
		}
	}
	if len(_kmsDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _kmsDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}
	if len(_kmsEncryptionAlgorithm) > 0 {
		if err := assignInputField(input, "EncryptionAlgorithm", _kmsEncryptionAlgorithm); err != nil {
			log.Errorf("invalid --encryption-algorithm: %s", err.Error())
			return
		}
	}
	if len(_kmsEncryptionContext) > 0 {
		if err := assignInputField(input, "EncryptionContext", _kmsEncryptionContext); err != nil {
			log.Errorf("invalid --encryption-context: %s", err.Error())
			return
		}
	}
	if len(_kmsGrantTokens) > 0 {
		input.GrantTokens = append([]string(nil), _kmsGrantTokens...)
	}

	if resp, err := client.Encrypt(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a unique symmetric data key for use outside of KMS. This operation
// returns a plaintext copy of the data key and a copy that is encrypted under a
// symmetric encryption KMS key that you specify. The bytes in the plaintext key
// are random; they are not related to the caller or the KMS key. You can use the
// plaintext key to encrypt your data outside of KMS and store the encrypted data
// key with the encrypted data.
//
// To generate a data key, specify the symmetric encryption KMS key that will be
// used to encrypt the data key. You cannot use an asymmetric KMS key to encrypt
// data keys. To get the type of your KMS key, use the DescribeKeyoperation.
//
// You must also specify the length of the data key. Use either the KeySpec or
// NumberOfBytes parameters (but not both). For 128-bit and 256-bit data keys, use
// the KeySpec parameter.
//
// To generate a 128-bit SM4 data key (China Regions only), specify a KeySpec
// value of AES_128 or a NumberOfBytes value of 16 . The symmetric encryption key
// used in China Regions to encrypt your data key is an SM4 encryption key.
//
// To get only an encrypted copy of the data key, use GenerateDataKeyWithoutPlaintext. To generate an asymmetric
// data key pair, use the GenerateDataKeyPairor GenerateDataKeyPairWithoutPlaintext operation. To get a cryptographically secure random
// byte string, use GenerateRandom.
//
// You can use an optional encryption context to add additional security to the
// encryption operation. If you specify an EncryptionContext , you must specify the
// same encryption context (a case-sensitive exact match) when decrypting the
// encrypted data key. Otherwise, the request to decrypt fails with an
// InvalidCiphertextException . For more information, see [Encryption Context] in the Key Management
// Service Developer Guide.
//
// GenerateDataKey also supports [Amazon Web Services Nitro Enclaves], which provide an isolated compute environment
// in Amazon EC2. To call GenerateDataKey for an Amazon Web Services Nitro enclave
// or NitroTPM, use the [Amazon Web Services Nitro Enclaves SDK]or any Amazon Web Services SDK. Use the Recipient
// parameter to provide the attestation document for the attested environment.
// GenerateDataKey returns a copy of the data key encrypted under the specified KMS
// key, as usual. But instead of a plaintext copy of the data key, the response
// includes a copy of the data key encrypted under the public key from the
// attestation document ( CiphertextForRecipient ). For information about the
// interaction between KMS and Amazon Web Services Nitro Enclaves or Amazon Web
// Services NitroTPM, see [Cryptographic attestation support in KMS]in the Key Management Service Developer Guide.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// # How to use your data key
//
// We recommend that you use the following pattern to encrypt data locally in your
// application. You can write your own code or use a client-side encryption
// library, such as the [Amazon Web Services Encryption SDK], the [Amazon DynamoDB Encryption Client], or [Amazon S3 client-side encryption] to do these tasks for you.
//
// To encrypt data outside of KMS:
//
// - Use the GenerateDataKey operation to get a data key.
//
// - Use the plaintext data key (in the Plaintext field of the response) to
// encrypt your data outside of KMS. Then erase the plaintext data key from memory.
//
// - Store the encrypted data key (in the CiphertextBlob field of the response)
// with the encrypted data.
//
// To decrypt data outside of KMS:
//
// - Use the Decryptoperation to decrypt the encrypted data key. The operation returns
// a plaintext copy of the data key.
//
// - Use the plaintext data key to decrypt data outside of KMS, then erase the
// plaintext data key from memory.
//
// Cross-account use: Yes. To perform this operation with a KMS key in a different
// Amazon Web Services account, specify the key ARN or alias ARN in the value of
// the KeyId parameter.
//
// Required permissions: [kms:GenerateDataKey] (key policy)
//
// Related operations:
//
// # Decrypt
//
// # Encrypt
//
// # GenerateDataKeyPair
//
// # GenerateDataKeyPairWithoutPlaintext
//
// # GenerateDataKeyWithoutPlaintext
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Cryptographic attestation support in KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/cryptographic-attestation.html
// [Amazon Web Services Encryption SDK]: https://docs.aws.amazon.com/encryption-sdk/latest/developer-guide/
// [Amazon DynamoDB Encryption Client]: https://docs.aws.amazon.com/dynamodb-encryption-client/latest/devguide/
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [Encryption Context]: https://docs.aws.amazon.com/kms/latest/developerguide/encrypt_context.html
// [Amazon Web Services Nitro Enclaves]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/nitro-enclave.html
// [Amazon S3 client-side encryption]: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingClientSideEncryption.html
// [kms:GenerateDataKey]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [Amazon Web Services Nitro Enclaves SDK]: https://docs.aws.amazon.com/enclaves/latest/user/developing-applications.html#sdk
func kms_GenerateDataKey(cfg aws.Config, client *kms.Client) {
	input := &kms.GenerateDataKeyInput{
		// KeyId: *string, // Required
	}

	if len(_kmsKeyId) > 0 {
		input.KeyId = aws.String(_kmsKeyId)
	}
	if len(_kmsDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _kmsDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}
	if len(_kmsEncryptionContext) > 0 {
		if err := assignInputField(input, "EncryptionContext", _kmsEncryptionContext); err != nil {
			log.Errorf("invalid --encryption-context: %s", err.Error())
			return
		}
	}
	if len(_kmsGrantTokens) > 0 {
		input.GrantTokens = append([]string(nil), _kmsGrantTokens...)
	}
	if len(_kmsKeySpec) > 0 {
		if err := assignInputField(input, "KeySpec", _kmsKeySpec); err != nil {
			log.Errorf("invalid --key-spec: %s", err.Error())
			return
		}
	}
	if len(_kmsNumberOfBytes) > 0 {
		if err := assignInputField(input, "NumberOfBytes", _kmsNumberOfBytes); err != nil {
			log.Errorf("invalid --number-of-bytes: %s", err.Error())
			return
		}
	}
	if len(_kmsRecipient) > 0 {
		if err := assignInputField(input, "Recipient", _kmsRecipient); err != nil {
			log.Errorf("invalid --recipient: %s", err.Error())
			return
		}
	}

	if resp, err := client.GenerateDataKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a unique asymmetric data key pair for use outside of KMS. This
// operation returns a plaintext public key, a plaintext private key, and a copy of
// the private key that is encrypted under the symmetric encryption KMS key you
// specify. You can use the data key pair to perform asymmetric cryptography and
// implement digital signatures outside of KMS. The bytes in the keys are random;
// they are not related to the caller or to the KMS key that is used to encrypt the
// private key.
//
// You can use the public key that GenerateDataKeyPair returns to encrypt data or
// verify a signature outside of KMS. Then, store the encrypted private key with
// the data. When you are ready to decrypt data or sign a message, you can use the Decrypt
// operation to decrypt the encrypted private key.
//
// To generate a data key pair, you must specify a symmetric encryption KMS key to
// encrypt the private key in a data key pair. You cannot use an asymmetric KMS key
// or a KMS key in a custom key store. To get the type and origin of your KMS key,
// use the DescribeKeyoperation.
//
// Use the KeyPairSpec parameter to choose an RSA or Elliptic Curve (ECC) data key
// pair. In China Regions, you can also choose an SM2 data key pair. KMS recommends
// that you use ECC key pairs for signing, and use RSA and SM2 key pairs for either
// encryption or signing, but not both. However, KMS cannot enforce any
// restrictions on the use of data key pairs outside of KMS.
//
// If you are using the data key pair to encrypt data, or for any operation where
// you don't immediately need a private key, consider using the GenerateDataKeyPairWithoutPlaintextoperation.
// GenerateDataKeyPairWithoutPlaintext returns a plaintext public key and an
// encrypted private key, but omits the plaintext private key that you need only to
// decrypt ciphertext or sign a message. Later, when you need to decrypt the data
// or sign a message, use the Decryptoperation to decrypt the encrypted private key in
// the data key pair.
//
// GenerateDataKeyPair returns a unique data key pair for each request. The bytes
// in the keys are random; they are not related to the caller or the KMS key that
// is used to encrypt the private key. The public key is a DER-encoded X.509
// SubjectPublicKeyInfo, as specified in [RFC 5280]. The private key is a DER-encoded PKCS8
// PrivateKeyInfo, as specified in [RFC 5958].
//
// GenerateDataKeyPair also supports [Amazon Web Services Nitro Enclaves], which provide an isolated compute
// environment in Amazon EC2. To call GenerateDataKeyPair for an Amazon Web
// Services Nitro enclave or NitroTPM, use the [Amazon Web Services Nitro Enclaves SDK]or any Amazon Web Services SDK. Use
// the Recipient parameter to provide the attestation document for the attested
// environment. GenerateDataKeyPair returns the public data key and a copy of the
// private data key encrypted under the specified KMS key, as usual. But instead of
// a plaintext copy of the private data key ( PrivateKeyPlaintext ), the response
// includes a copy of the private data key encrypted under the public key from the
// attestation document ( CiphertextForRecipient ). For information about the
// interaction between KMS and Amazon Web Services Nitro Enclaves or Amazon Web
// Services NitroTPM, see [Cryptographic attestation support in KMS]in the Key Management Service Developer Guide.
//
// You can use an optional encryption context to add additional security to the
// encryption operation. If you specify an EncryptionContext , you must specify the
// same encryption context (a case-sensitive exact match) when decrypting the
// encrypted data key. Otherwise, the request to decrypt fails with an
// InvalidCiphertextException . For more information, see [Encryption Context] in the Key Management
// Service Developer Guide.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: Yes. To perform this operation with a KMS key in a different
// Amazon Web Services account, specify the key ARN or alias ARN in the value of
// the KeyId parameter.
//
// Required permissions: [kms:GenerateDataKeyPair] (key policy)
//
// Related operations:
//
// # Decrypt
//
// # Encrypt
//
// # GenerateDataKey
//
// # GenerateDataKeyPairWithoutPlaintext
//
// # GenerateDataKeyWithoutPlaintext
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Cryptographic attestation support in KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/cryptographic-attestation.html
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [RFC 5280]: https://tools.ietf.org/html/rfc5280
// [Encryption Context]: https://docs.aws.amazon.com/kms/latest/developerguide/encrypt_context.html
// [Amazon Web Services Nitro Enclaves]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/nitro-enclave.html
// [RFC 5958]: https://tools.ietf.org/html/rfc5958
// [kms:GenerateDataKeyPair]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [Amazon Web Services Nitro Enclaves SDK]: https://docs.aws.amazon.com/enclaves/latest/user/developing-applications.html#sdk
func kms_GenerateDataKeyPair(cfg aws.Config, client *kms.Client) {
	input := &kms.GenerateDataKeyPairInput{
		// KeyId: *string, // Required
		// KeyPairSpec: types.DataKeyPairSpec, // Required
	}

	if len(_kmsKeyId) > 0 {
		input.KeyId = aws.String(_kmsKeyId)
	}
	if len(_kmsKeyPairSpec) > 0 {
		if err := assignInputField(input, "KeyPairSpec", _kmsKeyPairSpec); err != nil {
			log.Errorf("invalid --key-pair-spec: %s", err.Error())
			return
		}
	}
	if len(_kmsDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _kmsDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}
	if len(_kmsEncryptionContext) > 0 {
		if err := assignInputField(input, "EncryptionContext", _kmsEncryptionContext); err != nil {
			log.Errorf("invalid --encryption-context: %s", err.Error())
			return
		}
	}
	if len(_kmsGrantTokens) > 0 {
		input.GrantTokens = append([]string(nil), _kmsGrantTokens...)
	}
	if len(_kmsRecipient) > 0 {
		if err := assignInputField(input, "Recipient", _kmsRecipient); err != nil {
			log.Errorf("invalid --recipient: %s", err.Error())
			return
		}
	}

	if resp, err := client.GenerateDataKeyPair(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a unique asymmetric data key pair for use outside of KMS. This
// operation returns a plaintext public key and a copy of the private key that is
// encrypted under the symmetric encryption KMS key you specify. Unlike GenerateDataKeyPair, this
// operation does not return a plaintext private key. The bytes in the keys are
// random; they are not related to the caller or to the KMS key that is used to
// encrypt the private key.
//
// You can use the public key that GenerateDataKeyPairWithoutPlaintext returns to
// encrypt data or verify a signature outside of KMS. Then, store the encrypted
// private key with the data. When you are ready to decrypt data or sign a message,
// you can use the Decryptoperation to decrypt the encrypted private key.
//
// To generate a data key pair, you must specify a symmetric encryption KMS key to
// encrypt the private key in a data key pair. You cannot use an asymmetric KMS key
// or a KMS key in a custom key store. To get the type and origin of your KMS key,
// use the DescribeKeyoperation.
//
// Use the KeyPairSpec parameter to choose an RSA or Elliptic Curve (ECC) data key
// pair. In China Regions, you can also choose an SM2 data key pair. KMS recommends
// that you use ECC key pairs for signing, and use RSA and SM2 key pairs for either
// encryption or signing, but not both. However, KMS cannot enforce any
// restrictions on the use of data key pairs outside of KMS.
//
// GenerateDataKeyPairWithoutPlaintext returns a unique data key pair for each
// request. The bytes in the key are not related to the caller or KMS key that is
// used to encrypt the private key. The public key is a DER-encoded X.509
// SubjectPublicKeyInfo, as specified in [RFC 5280].
//
// You can use an optional encryption context to add additional security to the
// encryption operation. If you specify an EncryptionContext , you must specify the
// same encryption context (a case-sensitive exact match) when decrypting the
// encrypted data key. Otherwise, the request to decrypt fails with an
// InvalidCiphertextException . For more information, see [Encryption Context] in the Key Management
// Service Developer Guide.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: Yes. To perform this operation with a KMS key in a different
// Amazon Web Services account, specify the key ARN or alias ARN in the value of
// the KeyId parameter.
//
// Required permissions: [kms:GenerateDataKeyPairWithoutPlaintext] (key policy)
//
// Related operations:
//
// # Decrypt
//
// # Encrypt
//
// # GenerateDataKey
//
// # GenerateDataKeyPair
//
// # GenerateDataKeyWithoutPlaintext
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [RFC 5280]: https://tools.ietf.org/html/rfc5280
// [Encryption Context]: https://docs.aws.amazon.com/kms/latest/developerguide/encrypt_context.html
// [kms:GenerateDataKeyPairWithoutPlaintext]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
func kms_GenerateDataKeyPairWithoutPlaintext(cfg aws.Config, client *kms.Client) {
	input := &kms.GenerateDataKeyPairWithoutPlaintextInput{
		// KeyId: *string, // Required
		// KeyPairSpec: types.DataKeyPairSpec, // Required
	}

	if len(_kmsKeyId) > 0 {
		input.KeyId = aws.String(_kmsKeyId)
	}
	if len(_kmsKeyPairSpec) > 0 {
		if err := assignInputField(input, "KeyPairSpec", _kmsKeyPairSpec); err != nil {
			log.Errorf("invalid --key-pair-spec: %s", err.Error())
			return
		}
	}
	if len(_kmsDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _kmsDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}
	if len(_kmsEncryptionContext) > 0 {
		if err := assignInputField(input, "EncryptionContext", _kmsEncryptionContext); err != nil {
			log.Errorf("invalid --encryption-context: %s", err.Error())
			return
		}
	}
	if len(_kmsGrantTokens) > 0 {
		input.GrantTokens = append([]string(nil), _kmsGrantTokens...)
	}

	if resp, err := client.GenerateDataKeyPairWithoutPlaintext(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a unique symmetric data key for use outside of KMS. This operation
// returns a data key that is encrypted under a symmetric encryption KMS key that
// you specify. The bytes in the key are random; they are not related to the caller
// or to the KMS key.
//
// GenerateDataKeyWithoutPlaintext is identical to the GenerateDataKey operation except that it
// does not return a plaintext copy of the data key.
//
// This operation is useful for systems that need to encrypt data at some point,
// but not immediately. When you need to encrypt the data, you call the Decryptoperation
// on the encrypted copy of the key.
//
// It's also useful in distributed systems with different levels of trust. For
// example, you might store encrypted data in containers. One component of your
// system creates new containers and stores an encrypted data key with each
// container. Then, a different component puts the data into the containers. That
// component first decrypts the data key, uses the plaintext data key to encrypt
// data, puts the encrypted data into the container, and then destroys the
// plaintext data key. In this system, the component that creates the containers
// never sees the plaintext data key.
//
// To request an asymmetric data key pair, use the GenerateDataKeyPair or GenerateDataKeyPairWithoutPlaintext operations.
//
// To generate a data key, you must specify the symmetric encryption KMS key that
// is used to encrypt the data key. You cannot use an asymmetric KMS key or a key
// in a custom key store to generate a data key. To get the type of your KMS key,
// use the DescribeKeyoperation.
//
// You must also specify the length of the data key. Use either the KeySpec or
// NumberOfBytes parameters (but not both). For 128-bit and 256-bit data keys, use
// the KeySpec parameter.
//
// To generate an SM4 data key (China Regions only), specify a KeySpec value of
// AES_128 or NumberOfBytes value of 16 . The symmetric encryption key used in
// China Regions to encrypt your data key is an SM4 encryption key.
//
// If the operation succeeds, you will find the encrypted copy of the data key in
// the CiphertextBlob field.
//
// You can use an optional encryption context to add additional security to the
// encryption operation. If you specify an EncryptionContext , you must specify the
// same encryption context (a case-sensitive exact match) when decrypting the
// encrypted data key. Otherwise, the request to decrypt fails with an
// InvalidCiphertextException . For more information, see [Encryption Context] in the Key Management
// Service Developer Guide.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: Yes. To perform this operation with a KMS key in a different
// Amazon Web Services account, specify the key ARN or alias ARN in the value of
// the KeyId parameter.
//
// Required permissions: [kms:GenerateDataKeyWithoutPlaintext] (key policy)
//
// Related operations:
//
// # Decrypt
//
// # Encrypt
//
// # GenerateDataKey
//
// # GenerateDataKeyPair
//
// # GenerateDataKeyPairWithoutPlaintext
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [Encryption Context]: https://docs.aws.amazon.com/kms/latest/developerguide/encrypt_context.html
// [kms:GenerateDataKeyWithoutPlaintext]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
func kms_GenerateDataKeyWithoutPlaintext(cfg aws.Config, client *kms.Client) {
	input := &kms.GenerateDataKeyWithoutPlaintextInput{
		// KeyId: *string, // Required
	}

	if len(_kmsKeyId) > 0 {
		input.KeyId = aws.String(_kmsKeyId)
	}
	if len(_kmsDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _kmsDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}
	if len(_kmsEncryptionContext) > 0 {
		if err := assignInputField(input, "EncryptionContext", _kmsEncryptionContext); err != nil {
			log.Errorf("invalid --encryption-context: %s", err.Error())
			return
		}
	}
	if len(_kmsGrantTokens) > 0 {
		input.GrantTokens = append([]string(nil), _kmsGrantTokens...)
	}
	if len(_kmsKeySpec) > 0 {
		if err := assignInputField(input, "KeySpec", _kmsKeySpec); err != nil {
			log.Errorf("invalid --key-spec: %s", err.Error())
			return
		}
	}
	if len(_kmsNumberOfBytes) > 0 {
		if err := assignInputField(input, "NumberOfBytes", _kmsNumberOfBytes); err != nil {
			log.Errorf("invalid --number-of-bytes: %s", err.Error())
			return
		}
	}

	if resp, err := client.GenerateDataKeyWithoutPlaintext(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Generates a hash-based message authentication code (HMAC) for a message using
// an HMAC KMS key and a MAC algorithm that the key supports. HMAC KMS keys and the
// HMAC algorithms that KMS uses conform to industry standards defined in [RFC 2104].
//
// You can use value that GenerateMac returns in the VerifyMac operation to demonstrate
// that the original message has not changed. Also, because a secret key is used to
// create the hash, you can verify that the party that generated the hash has the
// required secret key. You can also use the raw result to implement HMAC-based
// algorithms such as key derivation functions. This operation is part of KMS
// support for HMAC KMS keys. For details, see [HMAC keys in KMS]in the Key Management Service
// Developer Guide .
//
// Best practices recommend that you limit the time during which any signing
// mechanism, including an HMAC, is effective. This deters an attack where the
// actor uses a signed message to establish validity repeatedly or long after the
// message is superseded. HMAC tags do not include a timestamp, but you can include
// a timestamp in the token or message to help you detect when its time to refresh
// the HMAC.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: Yes. To perform this operation with a KMS key in a different
// Amazon Web Services account, specify the key ARN or alias ARN in the value of
// the KeyId parameter.
//
// Required permissions: [kms:GenerateMac] (key policy)
//
// Related operations: VerifyMac
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [kms:GenerateMac]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [RFC 2104]: https://datatracker.ietf.org/doc/html/rfc2104
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [HMAC keys in KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/hmac.html
func kms_GenerateMac(cfg aws.Config, client *kms.Client) {
	input := &kms.GenerateMacInput{
		// KeyId: *string, // Required
		// MacAlgorithm: types.MacAlgorithmSpec, // Required
		// Message: []byte, // Required
	}

	if len(_kmsKeyId) > 0 {
		input.KeyId = aws.String(_kmsKeyId)
	}
	if len(_kmsMacAlgorithm) > 0 {
		if err := assignInputField(input, "MacAlgorithm", _kmsMacAlgorithm); err != nil {
			log.Errorf("invalid --mac-algorithm: %s", err.Error())
			return
		}
	}
	if len(_kmsMessage) > 0 {
		if err := assignInputField(input, "Message", _kmsMessage); err != nil {
			log.Errorf("invalid --message: %s", err.Error())
			return
		}
	}
	if len(_kmsDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _kmsDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}
	if len(_kmsGrantTokens) > 0 {
		input.GrantTokens = append([]string(nil), _kmsGrantTokens...)
	}

	if resp, err := client.GenerateMac(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a random byte string that is cryptographically secure.
// You must use the NumberOfBytes parameter to specify the length of the random
// byte string. There is no default value for string length.
//
// By default, the random byte string is generated in KMS. To generate the byte
// string in the CloudHSM cluster associated with an CloudHSM key store, use the
// CustomKeyStoreId parameter.
//
// GenerateRandom also supports [Amazon Web Services Nitro Enclaves], which provide an isolated compute environment in
// Amazon EC2. To call GenerateRandom for a Nitro enclave or NitroTPM, use the [Amazon Web Services Nitro Enclaves SDK] or
// any Amazon Web Services SDK. Use the Recipient parameter to provide the
// attestation document for the attested environment. Instead of plaintext bytes,
// the response includes the plaintext bytes encrypted under the public key from
// the attestation document ( CiphertextForRecipient ). For information about the
// interaction between KMS and Amazon Web Services Nitro Enclaves or Amazon Web
// Services NitroTPM, see [Cryptographic attestation support in KMS]in the Key Management Service Developer Guide.
//
// For more information about entropy and random number generation, see [Entropy and random number generation] in the
// Key Management Service Developer Guide.
//
// Cross-account use: Not applicable. GenerateRandom does not use any
// account-specific resources, such as KMS keys.
//
// Required permissions: [kms:GenerateRandom] (IAM policy)
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Cryptographic attestation support in KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/cryptographic-attestation.html
// [kms:GenerateRandom]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [Amazon Web Services Nitro Enclaves]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/nitro-enclave.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [Amazon Web Services Nitro Enclaves SDK]: https://docs.aws.amazon.com/enclaves/latest/user/developing-applications.html#sdk
// [Entropy and random number generation]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-cryptography.html#entropy-and-random-numbers
func kms_GenerateRandom(cfg aws.Config, client *kms.Client) {
	input := &kms.GenerateRandomInput{}

	if len(_kmsCustomKeyStoreId) > 0 {
		input.CustomKeyStoreId = aws.String(_kmsCustomKeyStoreId)
	}
	if len(_kmsNumberOfBytes) > 0 {
		if err := assignInputField(input, "NumberOfBytes", _kmsNumberOfBytes); err != nil {
			log.Errorf("invalid --number-of-bytes: %s", err.Error())
			return
		}
	}
	if len(_kmsRecipient) > 0 {
		if err := assignInputField(input, "Recipient", _kmsRecipient); err != nil {
			log.Errorf("invalid --recipient: %s", err.Error())
			return
		}
	}

	if resp, err := client.GenerateRandom(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a key policy attached to the specified KMS key.
// Cross-account use: No. You cannot perform this operation on a KMS key in a
// different Amazon Web Services account.
//
// Required permissions: [kms:GetKeyPolicy] (key policy)
//
// Related operations: [PutKeyPolicy]
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [kms:GetKeyPolicy]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [PutKeyPolicy]: https://docs.aws.amazon.com/kms/latest/APIReference/API_PutKeyPolicy.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
func kms_GetKeyPolicy(cfg aws.Config, client *kms.Client) {
	input := &kms.GetKeyPolicyInput{
		// KeyId: *string, // Required
	}

	if len(_kmsKeyId) > 0 {
		input.KeyId = aws.String(_kmsKeyId)
	}
	if len(_kmsPolicyName) > 0 {
		input.PolicyName = aws.String(_kmsPolicyName)
	}

	if resp, err := client.GetKeyPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides detailed information about the rotation status for a KMS key,
// including whether [automatic rotation of the key material]is enabled for the specified KMS key, the [rotation period], and the next
// scheduled rotation date.
//
// Automatic key rotation is supported only on symmetric encryption KMS keys. You
// cannot enable automatic rotation of [asymmetric KMS keys], [HMAC KMS keys], KMS keys with [imported key material], or KMS keys in a [custom key store]. To
// enable or disable automatic rotation of a set of related [multi-Region keys], set the property on
// the primary key.
//
// You can enable (EnableKeyRotation ) and disable automatic rotation (DisableKeyRotation ) of the key material in
// customer managed KMS keys. Key material rotation of [Amazon Web Services managed KMS keys]is not configurable. KMS
// always rotates the key material in Amazon Web Services managed KMS keys every
// year. The key rotation status for Amazon Web Services managed KMS keys is always
// true .
//
// You can perform on-demand (RotateKeyOnDemand ) rotation of the key material in customer managed
// KMS keys, regardless of whether or not automatic key rotation is enabled. You
// can use GetKeyRotationStatus to identify the date and time that an in progress
// on-demand rotation was initiated. You can use ListKeyRotationsto view the details of completed
// rotations.
//
// In May 2022, KMS changed the rotation schedule for Amazon Web Services managed
// keys from every three years to every year. For details, see EnableKeyRotation.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// - Disabled: The key rotation status does not change when you disable a KMS
// key. However, while the KMS key is disabled, KMS does not rotate the key
// material. When you re-enable the KMS key, rotation resumes. If the key material
// in the re-enabled KMS key hasn't been rotated in one year, KMS rotates it
// immediately, and every year thereafter. If it's been less than a year since the
// key material in the re-enabled KMS key was rotated, the KMS key resumes its
// prior rotation schedule.
//
// - Pending deletion: While a KMS key is pending deletion, its key rotation
// status is false and KMS does not rotate the key material. If you cancel the
// deletion, the original key rotation status returns to true .
//
// Cross-account use: Yes. To perform this operation on a KMS key in a different
// Amazon Web Services account, specify the key ARN in the value of the KeyId
// parameter.
//
// Required permissions: [kms:GetKeyRotationStatus] (key policy)
//
// Related operations:
//
// # DisableKeyRotation
//
// # EnableKeyRotation
//
// # ListKeyRotations
//
// # RotateKeyOnDemand
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [imported key material]: https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [HMAC KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/hmac.html
// [rotation period]: https://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html#rotation-period
// [Amazon Web Services managed KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-key
// [kms:GetKeyRotationStatus]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [automatic rotation of the key material]: https://docs.aws.amazon.com/kms/latest/developerguide/rotating-keys-enable-disable.html
// [asymmetric KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html
// [multi-Region keys]: https://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html#multi-region-rotate
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [custom key store]: https://docs.aws.amazon.com/kms/latest/developerguide/key-store-overview.html
func kms_GetKeyRotationStatus(cfg aws.Config, client *kms.Client) {
	input := &kms.GetKeyRotationStatusInput{
		// KeyId: *string, // Required
	}

	if len(_kmsKeyId) > 0 {
		input.KeyId = aws.String(_kmsKeyId)
	}

	if resp, err := client.GetKeyRotationStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the public key and an import token you need to import or reimport key
// material for a KMS key.
//
// By default, KMS keys are created with key material that KMS generates. This
// operation supports [Importing key material], an advanced feature that lets you generate and import the
// cryptographic key material for a KMS key.
//
// Before calling GetParametersForImport , use the CreateKey operation with an Origin value
// of EXTERNAL to create a KMS key with no key material. You can import key
// material for a symmetric encryption KMS key, HMAC KMS key, asymmetric encryption
// KMS key, or asymmetric signing KMS key. You can also import key material into a [multi-Region key]
// of any supported type. However, you can't import key material into a KMS key in
// a [custom key store]. You can also use GetParametersForImport to get a public key and import
// token to [reimport the original key material]into a KMS key whose key material expired or was deleted.
//
// GetParametersForImport returns the items that you need to import your key
// material.
//
// - The public key (or "wrapping key") of an RSA key pair that KMS generates.
//
// You will use this public key to encrypt ("wrap") your key material while it's
//
// in transit to KMS.
//
// - A import token that ensures that KMS can decrypt your key material and
// associate it with the correct KMS key.
//
// The public key and its import token are permanently linked and must be used
// together. Each public key and import token set is valid for 24 hours. The
// expiration date and time appear in the ParametersValidTo field in the
// GetParametersForImport response. You cannot use an expired public key or import
// token in an ImportKeyMaterialrequest. If your key and token expire, send another
// GetParametersForImport request.
//
// GetParametersForImport requires the following information:
//
// - The key ID of the KMS key for which you are importing the key material.
//
// - The key spec of the public key ("wrapping key") that you will use to
// encrypt your key material during import.
//
// - The wrapping algorithm that you will use with the public key to encrypt
// your key material.
//
// You can use the same or a different public key spec and wrapping algorithm each
// time you import or reimport the same key material.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: No. You cannot perform this operation on a KMS key in a
// different Amazon Web Services account.
//
// Required permissions: [kms:GetParametersForImport] (key policy)
//
// Related operations:
//
// # ImportKeyMaterial
//
// # DeleteImportedKeyMaterial
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Importing key material]: https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html
// [kms:GetParametersForImport]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [reimport the original key material]: https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys-import-key-material.html#reimport-key-material
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [multi-Region key]: https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html
// [custom key store]: https://docs.aws.amazon.com/kms/latest/developerguide/key-store-overview.html
func kms_GetParametersForImport(cfg aws.Config, client *kms.Client) {
	input := &kms.GetParametersForImportInput{
		// KeyId: *string, // Required
		// WrappingAlgorithm: types.AlgorithmSpec, // Required
		// WrappingKeySpec: types.WrappingKeySpec, // Required
	}

	if len(_kmsKeyId) > 0 {
		input.KeyId = aws.String(_kmsKeyId)
	}
	if len(_kmsWrappingAlgorithm) > 0 {
		if err := assignInputField(input, "WrappingAlgorithm", _kmsWrappingAlgorithm); err != nil {
			log.Errorf("invalid --wrapping-algorithm: %s", err.Error())
			return
		}
	}
	if len(_kmsWrappingKeySpec) > 0 {
		if err := assignInputField(input, "WrappingKeySpec", _kmsWrappingKeySpec); err != nil {
			log.Errorf("invalid --wrapping-key-spec: %s", err.Error())
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

// Returns the public key of an asymmetric KMS key. Unlike the private key of a
// asymmetric KMS key, which never leaves KMS unencrypted, callers with
// kms:GetPublicKey permission can download the public key of an asymmetric KMS
// key. You can share the public key to allow others to encrypt messages and verify
// signatures outside of KMS. For information about asymmetric KMS keys, see [Asymmetric KMS keys]in
// the Key Management Service Developer Guide.
//
// You do not need to download the public key. Instead, you can use the public key
// within KMS by calling the Encrypt, ReEncrypt, or Verify operations with the identifier of an
// asymmetric KMS key. When you use the public key within KMS, you benefit from the
// authentication, authorization, and logging that are part of every KMS operation.
// You also reduce of risk of encrypting data that cannot be decrypted. These
// features are not effective outside of KMS.
//
// To help you use the public key safely outside of KMS, GetPublicKey returns
// important information about the public key in the response, including:
//
// [KeySpec]
// - : The type of key material in the public key, such as RSA_4096 or
// ECC_NIST_P521 .
//
// [KeyUsage]
// - : Whether the key is used for encryption, signing, or deriving a shared
// secret.
//
// [EncryptionAlgorithms]
// - , [KeyAgreementAlgorithms], or [SigningAlgorithms]: A list of the encryption algorithms, key agreement algorithms, or
// signing algorithms for the key.
//
// Although KMS cannot enforce these restrictions on external operations, it is
// crucial that you use this information to prevent the public key from being used
// improperly. For example, you can prevent a public signing key from being used
// encrypt data, or prevent a public key from being used with an encryption
// algorithm that is not supported by KMS. You can also avoid errors, such as using
// the wrong signing algorithm in a verification operation.
//
// To verify a signature outside of KMS with an SM2 public key (China Regions
// only), you must specify the distinguishing ID. By default, KMS uses
// 1234567812345678 as the distinguishing ID. For more information, see [Offline verification with SM2 key pairs].
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: Yes. To perform this operation with a KMS key in a different
// Amazon Web Services account, specify the key ARN or alias ARN in the value of
// the KeyId parameter.
//
// Required permissions: [kms:GetPublicKey] (key policy)
//
// Related operations: CreateKey
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [SigningAlgorithms]: https://docs.aws.amazon.com/kms/latest/APIReference/API_GetPublicKey.html#KMS-GetPublicKey-response-SigningAlgorithms
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [kms:GetPublicKey]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [EncryptionAlgorithms]: https://docs.aws.amazon.com/kms/latest/APIReference/API_GetPublicKey.html#KMS-GetPublicKey-response-EncryptionAlgorithms
// [Asymmetric KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html
// [KeyAgreementAlgorithms]: https://docs.aws.amazon.com/kms/latest/APIReference/API_GetPublicKey.html#KMS-GetPublicKey-response-KeyAgreementAlgorithms
// [KeySpec]: https://docs.aws.amazon.com/kms/latest/APIReference/API_GetPublicKey.html#KMS-GetPublicKey-response-KeySpec
// [Offline verification with SM2 key pairs]: https://docs.aws.amazon.com/kms/latest/developerguide/offline-operations.html#key-spec-sm-offline-verification
// [KeyUsage]: https://docs.aws.amazon.com/kms/latest/APIReference/API_GetPublicKey.html#KMS-GetPublicKey-response-KeyUsage
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
func kms_GetPublicKey(cfg aws.Config, client *kms.Client) {
	input := &kms.GetPublicKeyInput{
		// KeyId: *string, // Required
	}

	if len(_kmsKeyId) > 0 {
		input.KeyId = aws.String(_kmsKeyId)
	}
	if len(_kmsGrantTokens) > 0 {
		input.GrantTokens = append([]string(nil), _kmsGrantTokens...)
	}

	if resp, err := client.GetPublicKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Imports or reimports key material into an existing KMS key that was created
// without key material. You can also use this operation to set or update the
// expiration model and expiration date of the imported key material.
//
// By default, KMS creates KMS keys with key material that it generates. You can
// also generate and import your own key material. For more information about
// importing key material, see [Importing key material].
//
// For asymmetric and HMAC keys, you cannot change the key material after the
// initial import. You can import multiple key materials into symmetric encryption
// keys and rotate the key material on demand using RotateKeyOnDemand .
//
// You can import new key materials into multi-Region symmetric encryption keys.
// To do so, you must import the new key material into the primary Region key. Then
// you can import the same key materials into the replica Region keys. You cannot
// directly import new key material into the replica Region keys.
//
// To import new key material for a multi-Region symmetric key, you’ll need to
// complete the following:
//
// - Call ImportKeyMaterial on the primary Region key with the ImportType set to
// NEW_KEY_MATERIAL .
//
// - Call ImportKeyMaterial on the replica Region key with the ImportType set to
// EXISTING_KEY_MATERIAL using the same key material imported to the primary
// Region key. You must do this for every replica Region key before you can perform
// the RotateKeyOnDemandoperation on the primary Region key.
//
// After you import key material, you can [reimport the same key material] into that KMS key or, if the key
// supports on-demand rotation, import new key material. You can use the ImportType
// parameter to indicate whether you are importing new key material or re-importing
// previously imported key material. You might reimport key material to replace key
// material that expired or key material that you deleted. You might also reimport
// key material to change the expiration model or expiration date of the key
// material.
//
// Each time you import key material into KMS, you can determine whether (
// ExpirationModel ) and when ( ValidTo ) the key material expires. To change the
// expiration of your key material, you must import it again, either by calling
// ImportKeyMaterial or using the [import features] of the KMS console.
//
// Before you call ImportKeyMaterial , complete these steps:
//
// - Create or identify a KMS key with EXTERNAL origin, which indicates that the
// KMS key is designed for imported key material.
//
// # To create a new KMS key for imported key material, call the CreateKeyoperation with an
//
// Origin value of EXTERNAL . You can create a symmetric encryption KMS key, HMAC
// KMS key, asymmetric encryption KMS key, asymmetric key agreement key, or
// asymmetric signing KMS key. You can also import key material into a [multi-Region key]of any
// supported type. However, you can't import key material into a KMS key in a [custom key store].
//
// - Call the GetParametersForImportoperation to get a public key and import token set for importing
// key material.
//
// - Use the public key in the GetParametersForImportresponse to encrypt your key material.
//
// Then, in an ImportKeyMaterial request, you submit your encrypted key material
// and import token. When calling this operation, you must specify the following
// values:
//
// - The key ID or key ARN of the KMS key to associate with the imported key
// material. Its Origin must be EXTERNAL and its KeyState must be PendingImport
// or Enabled . You cannot perform this operation on a KMS key in a [custom key store], or on a
// KMS key in a different Amazon Web Services account. To get the Origin and
// KeyState of a KMS key, call DescribeKey.
//
// - The encrypted key material.
//
// - The import token that GetParametersForImportreturned. You must use a public key and token from
// the same GetParametersForImport response.
//
// - Whether the key material expires ( ExpirationModel ) and, if so, when (
// ValidTo ). For help with this choice, see [Setting an expiration time]in the Key Management Service
// Developer Guide.
//
// # If you set an expiration date, KMS deletes the key material from the KMS key on
//
// the specified date, making the KMS key unusable. To use the KMS key in
// cryptographic operations again, you must reimport the same key material.
// However, you can delete and reimport the key material at any time, including
// before the key material expires. Each time you reimport, you can eliminate or
// reset the expiration time.
//
// When this operation is successful, the state of the KMS key changes to Enabled ,
// and you can use the KMS key in cryptographic operations. For symmetric
// encryption keys, you will need to import all of the key materials associated
// with the KMS key to change its state to Enabled . Use the ListKeyRotations
// operation to list the ID and import state of each key material associated with a
// KMS key.
//
// If this operation fails, use the exception to help determine the problem. If
// the error is related to the key material, the import token, or wrapping key, use
// GetParametersForImportto get a new public key and import token for the KMS key and repeat the import
// procedure. For help, see [Create a KMS key with imported key material]in the Key Management Service Developer Guide.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: No. You cannot perform this operation on a KMS key in a
// different Amazon Web Services account.
//
// Required permissions: [kms:ImportKeyMaterial] (key policy)
//
// Related operations:
//
// # DeleteImportedKeyMaterial
//
// # GetParametersForImport
//
// # ListKeyRotations
//
// # RotateKeyOnDemand
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Importing key material]: https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [kms:ImportKeyMaterial]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [reimport the same key material]: https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys-import-key-material.html#reimport-key-material
// [import features]: https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys-import-key-material.html#importing-keys-import-key-material-console
// [Create a KMS key with imported key material]: https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys-conceptual.html
// [Setting an expiration time]: https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys-import-key-material.html#importing-keys-expiration
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [custom key store]: https://docs.aws.amazon.com/kms/latest/developerguide/key-store-overview.html
// [multi-Region key]: https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html
func kms_ImportKeyMaterial(cfg aws.Config, client *kms.Client) {
	input := &kms.ImportKeyMaterialInput{
		// EncryptedKeyMaterial: []byte, // Required
		// ImportToken: []byte, // Required
		// KeyId: *string, // Required
	}

	if len(_kmsEncryptedKeyMaterial) > 0 {
		if err := assignInputField(input, "EncryptedKeyMaterial", _kmsEncryptedKeyMaterial); err != nil {
			log.Errorf("invalid --encrypted-key-material: %s", err.Error())
			return
		}
	}
	if len(_kmsImportToken) > 0 {
		if err := assignInputField(input, "ImportToken", _kmsImportToken); err != nil {
			log.Errorf("invalid --import-token: %s", err.Error())
			return
		}
	}
	if len(_kmsKeyId) > 0 {
		input.KeyId = aws.String(_kmsKeyId)
	}
	if len(_kmsExpirationModel) > 0 {
		if err := assignInputField(input, "ExpirationModel", _kmsExpirationModel); err != nil {
			log.Errorf("invalid --expiration-model: %s", err.Error())
			return
		}
	}
	if len(_kmsImportType) > 0 {
		if err := assignInputField(input, "ImportType", _kmsImportType); err != nil {
			log.Errorf("invalid --import-type: %s", err.Error())
			return
		}
	}
	if len(_kmsKeyMaterialDescription) > 0 {
		input.KeyMaterialDescription = aws.String(_kmsKeyMaterialDescription)
	}
	if len(_kmsKeyMaterialId) > 0 {
		input.KeyMaterialId = aws.String(_kmsKeyMaterialId)
	}
	if len(_kmsValidTo) > 0 {
		if err := assignInputField(input, "ValidTo", _kmsValidTo); err != nil {
			log.Errorf("invalid --valid-to: %s", err.Error())
			return
		}
	}

	if resp, err := client.ImportKeyMaterial(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a list of aliases in the caller's Amazon Web Services account and region.
// For more information about aliases, see CreateAlias.
//
// By default, the ListAliases operation returns all aliases in the account and
// region. To get only the aliases associated with a particular KMS key, use the
// KeyId parameter.
//
// The ListAliases response can include aliases that you created and associated
// with your customer managed keys, and aliases that Amazon Web Services created
// and associated with Amazon Web Services managed keys in your account. You can
// recognize Amazon Web Services aliases because their names have the format aws/ ,
// such as aws/dynamodb .
//
// The response might also include aliases that have no TargetKeyId field. These
// are predefined aliases that Amazon Web Services has created but has not yet
// associated with a KMS key. Aliases that Amazon Web Services creates in your
// account, including predefined aliases, do not count against your [KMS aliases quota].
//
// Cross-account use: No. ListAliases does not return aliases in other Amazon Web
// Services accounts.
//
// Required permissions: [kms:ListAliases] (IAM policy)
//
// For details, see [Controlling access to aliases] in the Key Management Service Developer Guide.
//
// Related operations:
//
// # CreateAlias
//
// # DeleteAlias
//
// # UpdateAlias
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [KMS aliases quota]: https://docs.aws.amazon.com/kms/latest/developerguide/resource-limits.html#aliases-per-key
// [kms:ListAliases]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [Controlling access to aliases]: https://docs.aws.amazon.com/kms/latest/developerguide/alias-access.html
func kms_ListAliases(cfg aws.Config, client *kms.Client) {
	input := &kms.ListAliasesInput{}

	if len(_kmsKeyId) > 0 {
		input.KeyId = aws.String(_kmsKeyId)
	}
	if len(_kmsLimit) > 0 {
		if err := assignInputField(input, "Limit", _kmsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_kmsMarker) > 0 {
		input.Marker = aws.String(_kmsMarker)
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

	var results []*kms.ListAliasesOutput
	p := kms.NewListAliasesPaginator(client, input)
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

// Gets a list of all grants for the specified KMS key.
// You must specify the KMS key in all requests. You can filter the grant list by
// grant ID or grantee principal.
//
// For detailed information about grants, including grant terminology, see [Grants in KMS] in the
// Key Management Service Developer Guide . For examples of creating grants in
// several programming languages, see [Use CreateGrant with an Amazon Web Services SDK or CLI].
//
// The GranteePrincipal field in the ListGrants response usually contains the user
// or role designated as the grantee principal in the grant. However, when the
// grantee principal in the grant is an Amazon Web Services service, the
// GranteePrincipal field contains the [service principal], which might represent several different
// grantee principals.
//
// Cross-account use: Yes. To perform this operation on a KMS key in a different
// Amazon Web Services account, specify the key ARN in the value of the KeyId
// parameter.
//
// Required permissions: [kms:ListGrants] (key policy)
//
// Related operations:
//
// # CreateGrant
//
// # ListRetirableGrants
//
// # RetireGrant
//
// # RevokeGrant
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [service principal]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_principal.html#principal-services
// [Grants in KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/grants.html
// [Use CreateGrant with an Amazon Web Services SDK or CLI]: https://docs.aws.amazon.com/kms/latest/developerguide/example_kms_CreateGrant_section.html
// [kms:ListGrants]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
func kms_ListGrants(cfg aws.Config, client *kms.Client) {
	input := &kms.ListGrantsInput{
		// KeyId: *string, // Required
	}

	if len(_kmsKeyId) > 0 {
		input.KeyId = aws.String(_kmsKeyId)
	}
	if len(_kmsGrantId) > 0 {
		input.GrantId = aws.String(_kmsGrantId)
	}
	if len(_kmsGranteePrincipal) > 0 {
		input.GranteePrincipal = aws.String(_kmsGranteePrincipal)
	}
	if len(_kmsLimit) > 0 {
		if err := assignInputField(input, "Limit", _kmsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_kmsMarker) > 0 {
		input.Marker = aws.String(_kmsMarker)
	}

	if disablePaginator() {
		if resp, err := client.ListGrants(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*kms.ListGrantsOutput
	p := kms.NewListGrantsPaginator(client, input)
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

// Gets the names of the key policies that are attached to a KMS key. This
// operation is designed to get policy names that you can use in a GetKeyPolicyoperation.
// However, the only valid policy name is default .
//
// Cross-account use: No. You cannot perform this operation on a KMS key in a
// different Amazon Web Services account.
//
// Required permissions: [kms:ListKeyPolicies] (key policy)
//
// Related operations:
//
// # GetKeyPolicy
//
// [PutKeyPolicy]
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [kms:ListKeyPolicies]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [PutKeyPolicy]: https://docs.aws.amazon.com/kms/latest/APIReference/API_PutKeyPolicy.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
func kms_ListKeyPolicies(cfg aws.Config, client *kms.Client) {
	input := &kms.ListKeyPoliciesInput{
		// KeyId: *string, // Required
	}

	if len(_kmsKeyId) > 0 {
		input.KeyId = aws.String(_kmsKeyId)
	}
	if len(_kmsLimit) > 0 {
		if err := assignInputField(input, "Limit", _kmsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_kmsMarker) > 0 {
		input.Marker = aws.String(_kmsMarker)
	}

	if disablePaginator() {
		if resp, err := client.ListKeyPolicies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*kms.ListKeyPoliciesOutput
	p := kms.NewListKeyPoliciesPaginator(client, input)
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

// Returns information about the key materials associated with the specified KMS
// key. You can use the optional IncludeKeyMaterial parameter to control which key
// materials are included in the response.
//
// You must specify the KMS key in all requests. You can refine the key rotations
// list by limiting the number of rotations returned.
//
// For detailed information about automatic and on-demand key rotations, see [Rotate KMS keys] in
// the Key Management Service Developer Guide.
//
// Cross-account use: No. You cannot perform this operation on a KMS key in a
// different Amazon Web Services account.
//
// Required permissions: [kms:ListKeyRotations] (key policy)
//
// Related operations:
//
// # EnableKeyRotation
//
// # DeleteImportedKeyMaterial
//
// # DisableKeyRotation
//
// # GetKeyRotationStatus
//
// # ImportKeyMaterial
//
// # RotateKeyOnDemand
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Rotate KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html
// [kms:ListKeyRotations]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
func kms_ListKeyRotations(cfg aws.Config, client *kms.Client) {
	input := &kms.ListKeyRotationsInput{
		// KeyId: *string, // Required
	}

	if len(_kmsKeyId) > 0 {
		input.KeyId = aws.String(_kmsKeyId)
	}
	if len(_kmsIncludeKeyMaterial) > 0 {
		if err := assignInputField(input, "IncludeKeyMaterial", _kmsIncludeKeyMaterial); err != nil {
			log.Errorf("invalid --include-key-material: %s", err.Error())
			return
		}
	}
	if len(_kmsLimit) > 0 {
		if err := assignInputField(input, "Limit", _kmsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_kmsMarker) > 0 {
		input.Marker = aws.String(_kmsMarker)
	}

	if disablePaginator() {
		if resp, err := client.ListKeyRotations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*kms.ListKeyRotationsOutput
	p := kms.NewListKeyRotationsPaginator(client, input)
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

// Gets a list of all KMS keys in the caller's Amazon Web Services account and
// Region.
//
// Cross-account use: No. You cannot perform this operation on a KMS key in a
// different Amazon Web Services account.
//
// Required permissions: [kms:ListKeys] (IAM policy)
//
// Related operations:
//
// # CreateKey
//
// # DescribeKey
//
// # ListAliases
//
// # ListResourceTags
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [kms:ListKeys]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
func kms_ListKeys(cfg aws.Config, client *kms.Client) {
	input := &kms.ListKeysInput{}

	if len(_kmsLimit) > 0 {
		if err := assignInputField(input, "Limit", _kmsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_kmsMarker) > 0 {
		input.Marker = aws.String(_kmsMarker)
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

	var results []*kms.ListKeysOutput
	p := kms.NewListKeysPaginator(client, input)
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

// Returns all tags on the specified KMS key.
// For general information about tags, including the format and syntax, see [Tagging Amazon Web Services resources] in
// the Amazon Web Services General Reference. For information about using tags in
// KMS, see [Tags in KMS].
//
// Cross-account use: No. You cannot perform this operation on a KMS key in a
// different Amazon Web Services account.
//
// Required permissions: [kms:ListResourceTags] (key policy)
//
// Related operations:
//
// # CreateKey
//
// # ReplicateKey
//
// # TagResource
//
// # UntagResource
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [kms:ListResourceTags]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [Tags in KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/tagging-keys.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [Tagging Amazon Web Services resources]: https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html
func kms_ListResourceTags(cfg aws.Config, client *kms.Client) {
	input := &kms.ListResourceTagsInput{
		// KeyId: *string, // Required
	}

	if len(_kmsKeyId) > 0 {
		input.KeyId = aws.String(_kmsKeyId)
	}
	if len(_kmsLimit) > 0 {
		if err := assignInputField(input, "Limit", _kmsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_kmsMarker) > 0 {
		input.Marker = aws.String(_kmsMarker)
	}

	if disablePaginator() {
		if resp, err := client.ListResourceTags(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*kms.ListResourceTagsOutput
	p := kms.NewListResourceTagsPaginator(client, input)
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

// Returns information about all grants in the Amazon Web Services account and
// Region that have the specified retiring principal.
//
// You can specify any principal in your Amazon Web Services account. The grants
// that are returned include grants for KMS keys in your Amazon Web Services
// account and other Amazon Web Services accounts. You might use this operation to
// determine which grants you may retire. To retire a grant, use the RetireGrantoperation.
//
// For detailed information about grants, including grant terminology, see [Grants in KMS] in the
// Key Management Service Developer Guide . For examples of creating grants in
// several programming languages, see [Use CreateGrant with an Amazon Web Services SDK or CLI].
//
// Cross-account use: You must specify a principal in your Amazon Web Services
// account. This operation returns a list of grants where the retiring principal
// specified in the ListRetirableGrants request is the same retiring principal on
// the grant. This can include grants on KMS keys owned by other Amazon Web
// Services accounts, but you do not need kms:ListRetirableGrants permission (or
// any other additional permission) in any Amazon Web Services account other than
// your own.
//
// Required permissions: [kms:ListRetirableGrants] (IAM policy) in your Amazon Web Services account.
//
// KMS authorizes ListRetirableGrants requests by evaluating the caller account's
// kms:ListRetirableGrants permissions. The authorized resource in
// ListRetirableGrants calls is the retiring principal specified in the request.
// KMS does not evaluate the caller's permissions to verify their access to any KMS
// keys or grants that might be returned by the ListRetirableGrants call.
//
// Related operations:
//
// # CreateGrant
//
// # ListGrants
//
// # RetireGrant
//
// # RevokeGrant
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [kms:ListRetirableGrants]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [Grants in KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/grants.html
// [Use CreateGrant with an Amazon Web Services SDK or CLI]: https://docs.aws.amazon.com/kms/latest/developerguide/example_kms_CreateGrant_section.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
func kms_ListRetirableGrants(cfg aws.Config, client *kms.Client) {
	input := &kms.ListRetirableGrantsInput{
		// RetiringPrincipal: *string, // Required
	}

	if len(_kmsRetiringPrincipal) > 0 {
		input.RetiringPrincipal = aws.String(_kmsRetiringPrincipal)
	}
	if len(_kmsLimit) > 0 {
		if err := assignInputField(input, "Limit", _kmsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_kmsMarker) > 0 {
		input.Marker = aws.String(_kmsMarker)
	}

	if disablePaginator() {
		if resp, err := client.ListRetirableGrants(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*kms.ListRetirableGrantsOutput
	p := kms.NewListRetirableGrantsPaginator(client, input)
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

// Attaches a key policy to the specified KMS key.
// For more information about key policies, see [Key Policies] in the Key Management Service
// Developer Guide. For help writing and formatting a JSON policy document, see the
// [IAM JSON Policy Reference]in the Identity and Access Management User Guide . For examples of adding a key
// policy in multiple programming languages, see [Use PutKeyPolicy with an Amazon Web Services SDK or CLI]in the Key Management Service
// Developer Guide.
//
// Cross-account use: No. You cannot perform this operation on a KMS key in a
// different Amazon Web Services account.
//
// Required permissions: [kms:PutKeyPolicy] (key policy)
//
// Related operations: GetKeyPolicy
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [IAM JSON Policy Reference]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies.html
// [kms:PutKeyPolicy]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [Use PutKeyPolicy with an Amazon Web Services SDK or CLI]: https://docs.aws.amazon.com/kms/latest/developerguide/example_kms_PutKeyPolicy_section.html
// [Key Policies]: https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
func kms_PutKeyPolicy(cfg aws.Config, client *kms.Client) {
	input := &kms.PutKeyPolicyInput{
		// KeyId: *string, // Required
		// Policy: *string, // Required
	}

	if len(_kmsKeyId) > 0 {
		input.KeyId = aws.String(_kmsKeyId)
	}
	if len(_kmsPolicy) > 0 {
		input.Policy = aws.String(_kmsPolicy)
	}
	if len(_kmsBypassPolicyLockoutSafetyCheck) > 0 {
		if err := assignInputField(input, "BypassPolicyLockoutSafetyCheck", _kmsBypassPolicyLockoutSafetyCheck); err != nil {
			log.Errorf("invalid --bypass-policy-lockout-safety-check: %s", err.Error())
			return
		}
	}
	if len(_kmsPolicyName) > 0 {
		input.PolicyName = aws.String(_kmsPolicyName)
	}

	if resp, err := client.PutKeyPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Decrypts ciphertext and then reencrypts it entirely within KMS. You can use
// this operation to change the KMS key under which data is encrypted, such as when
// you [manually rotate]a KMS key or change the KMS key that protects a ciphertext. You can also
// use it to reencrypt ciphertext under the same KMS key, such as to change the [encryption context]of
// a ciphertext.
//
// The ReEncrypt operation can decrypt ciphertext that was encrypted by using a
// KMS key in an KMS operation, such as Encryptor GenerateDataKey. It can also decrypt ciphertext that
// was encrypted by using the public key of an [asymmetric KMS key]outside of KMS. However, it cannot
// decrypt ciphertext produced by other libraries, such as the [Amazon Web Services Encryption SDK]or [Amazon S3 client-side encryption]. These
// libraries return a ciphertext format that is incompatible with KMS.
//
// When you use the ReEncrypt operation, you need to provide information for the
// decrypt operation and the subsequent encrypt operation.
//
// - If your ciphertext was encrypted under an asymmetric KMS key, you must use
// the SourceKeyId parameter to identify the KMS key that encrypted the
// ciphertext. You must also supply the encryption algorithm that was used. This
// information is required to decrypt the data.
//
// - If your ciphertext was encrypted under a symmetric encryption KMS key, the
// SourceKeyId parameter is optional. KMS can get this information from metadata
// that it adds to the symmetric ciphertext blob. This feature adds durability to
// your implementation by ensuring that authorized users can decrypt ciphertext
// decades after it was encrypted, even if they've lost track of the key ID.
// However, specifying the source KMS key is always recommended as a best practice.
// When you use the SourceKeyId parameter to specify a KMS key, KMS uses only the
// KMS key you specify. If the ciphertext was encrypted under a different KMS key,
// the ReEncrypt operation fails. This practice ensures that you use the KMS key
// that you intend.
//
// - To reencrypt the data, you must use the DestinationKeyId parameter to
// specify the KMS key that re-encrypts the data after it is decrypted. If the
// destination KMS key is an asymmetric KMS key, you must also provide the
// encryption algorithm. The algorithm that you choose must be compatible with the
// KMS key.
//
// # When you use an asymmetric KMS key to encrypt or reencrypt data, be sure to
//
// record the KMS key and encryption algorithm that you choose. You will be
// required to provide the same KMS key and encryption algorithm when you decrypt
// the data. If the KMS key and algorithm do not match the values used to encrypt
// the data, the decrypt operation fails.
//
// # You are not required to supply the key ID and encryption algorithm when you
//
// decrypt with symmetric encryption KMS keys because KMS stores this information
// in the ciphertext blob. KMS cannot store metadata in ciphertext generated with
// asymmetric keys. The standard format for asymmetric key ciphertext does not
// include configurable fields.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: Yes. The source KMS key and destination KMS key can be in
// different Amazon Web Services accounts. Either or both KMS keys can be in a
// different account than the caller. To specify a KMS key in a different account,
// you must use its key ARN or alias ARN.
//
// Required permissions:
//
// [kms:ReEncryptFrom]
// - permission on the source KMS key (key policy)
//
// [kms:ReEncryptTo]
// - permission on the destination KMS key (key policy)
//
// To permit reencryption from or to a KMS key, include the "kms:ReEncrypt*"
// permission in your [key policy]. This permission is automatically included in the key
// policy when you use the console to create a KMS key. But you must include it
// manually when you create a KMS key programmatically or when you use the PutKeyPolicy
// operation to set a key policy.
//
// Related operations:
//
// # Decrypt
//
// # Encrypt
//
// # GenerateDataKey
//
// # GenerateDataKeyPair
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Amazon Web Services Encryption SDK]: https://docs.aws.amazon.com/encryption-sdk/latest/developer-guide/
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [asymmetric KMS key]: https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html
// [key policy]: https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html
// [Amazon S3 client-side encryption]: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingClientSideEncryption.html
// [kms:ReEncryptTo]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [encryption context]: https://docs.aws.amazon.com/kms/latest/developerguide/encrypt_context.html
// [manually rotate]: https://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys-manually.html
// [kms:ReEncryptFrom]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
func kms_ReEncrypt(cfg aws.Config, client *kms.Client) {
	input := &kms.ReEncryptInput{
		// DestinationKeyId: *string, // Required
	}

	if len(_kmsDestinationKeyId) > 0 {
		input.DestinationKeyId = aws.String(_kmsDestinationKeyId)
	}
	if len(_kmsCiphertextBlob) > 0 {
		if err := assignInputField(input, "CiphertextBlob", _kmsCiphertextBlob); err != nil {
			log.Errorf("invalid --ciphertext-blob: %s", err.Error())
			return
		}
	}
	if len(_kmsDestinationEncryptionAlgorithm) > 0 {
		if err := assignInputField(input, "DestinationEncryptionAlgorithm", _kmsDestinationEncryptionAlgorithm); err != nil {
			log.Errorf("invalid --destination-encryption-algorithm: %s", err.Error())
			return
		}
	}
	if len(_kmsDestinationEncryptionContext) > 0 {
		if err := assignInputField(input, "DestinationEncryptionContext", _kmsDestinationEncryptionContext); err != nil {
			log.Errorf("invalid --destination-encryption-context: %s", err.Error())
			return
		}
	}
	if len(_kmsDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _kmsDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}
	if len(_kmsDryRunModifiers) > 0 {
		if err := assignInputField(input, "DryRunModifiers", _kmsDryRunModifiers); err != nil {
			log.Errorf("invalid --dry-run-modifiers: %s", err.Error())
			return
		}
	}
	if len(_kmsGrantTokens) > 0 {
		input.GrantTokens = append([]string(nil), _kmsGrantTokens...)
	}
	if len(_kmsSourceEncryptionAlgorithm) > 0 {
		if err := assignInputField(input, "SourceEncryptionAlgorithm", _kmsSourceEncryptionAlgorithm); err != nil {
			log.Errorf("invalid --source-encryption-algorithm: %s", err.Error())
			return
		}
	}
	if len(_kmsSourceEncryptionContext) > 0 {
		if err := assignInputField(input, "SourceEncryptionContext", _kmsSourceEncryptionContext); err != nil {
			log.Errorf("invalid --source-encryption-context: %s", err.Error())
			return
		}
	}
	if len(_kmsSourceKeyId) > 0 {
		input.SourceKeyId = aws.String(_kmsSourceKeyId)
	}

	if resp, err := client.ReEncrypt(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Replicates a multi-Region key into the specified Region. This operation creates
// a multi-Region replica key based on a multi-Region primary key in a different
// Region of the same Amazon Web Services partition. You can create multiple
// replicas of a primary key, but each must be in a different Region. To create a
// multi-Region primary key, use the CreateKeyoperation.
//
// This operation supports multi-Region keys, an KMS feature that lets you create
// multiple interoperable KMS keys in different Amazon Web Services Regions.
// Because these KMS keys have the same key ID, key material, and other metadata,
// you can use them interchangeably to encrypt data in one Amazon Web Services
// Region and decrypt it in a different Amazon Web Services Region without
// re-encrypting the data or making a cross-Region call. For more information about
// multi-Region keys, see [Multi-Region keys in KMS]in the Key Management Service Developer Guide.
//
// A replica key is a fully-functional KMS key that can be used independently of
// its primary and peer replica keys. A primary key and its replica keys share
// properties that make them interoperable. They have the same [key ID]and key material.
// They also have the same key spec, key usage, key material origin, and automatic
// key rotation status. KMS automatically synchronizes these shared properties
// among related multi-Region keys. All other properties of a replica key can
// differ, including its [key policy], [tags], [aliases], and [key state]. KMS pricing and quotas for KMS keys apply to
// each primary key and replica key.
//
// When this operation completes, the new replica key has a transient key state of
// Creating . This key state changes to Enabled (or PendingImport ) after a few
// seconds when the process of creating the new replica key is complete. While the
// key state is Creating , you can manage key, but you cannot yet use it in
// cryptographic operations. If you are creating and using the replica key
// programmatically, retry on KMSInvalidStateException or call DescribeKey to
// check its KeyState value before using it. For details about the Creating key
// state, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// You cannot create more than one replica of a primary key in any Region. If the
// Region already includes a replica of the key you're trying to replicate,
// ReplicateKey returns an AlreadyExistsException error. If the key state of the
// existing replica is PendingDeletion , you can cancel the scheduled key deletion (CancelKeyDeletion
// ) or wait for the key to be deleted. The new replica key you create will have
// the same [shared properties]as the original replica key.
//
// The CloudTrail log of a ReplicateKey operation records a ReplicateKey operation
// in the primary key's Region and a CreateKeyoperation in the replica key's Region.
//
// If you replicate a multi-Region primary key with imported key material, the
// replica key is created with no key material. You must import the same key
// material that you imported into the primary key.
//
// To convert a replica key to a primary key, use the UpdatePrimaryRegion operation.
//
// ReplicateKey uses different default values for the KeyPolicy and Tags
// parameters than those used in the KMS console. For details, see the parameter
// descriptions.
//
// Cross-account use: No. You cannot use this operation to create a replica key in
// a different Amazon Web Services account.
//
// Required permissions:
//
// - kms:ReplicateKey on the primary key (in the primary key's Region). Include
// this permission in the primary key's key policy.
//
// - kms:CreateKey in an IAM policy in the replica Region.
//
// - To use the Tags parameter, kms:TagResource in an IAM policy in the replica
// Region.
//
// # Related operations
//
// # CreateKey
//
// # UpdatePrimaryRegion
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [key ID]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id-key-id
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [aliases]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-alias.html
// [Multi-Region keys in KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html
// [key policy]: https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html
// [key state]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [tags]: https://docs.aws.amazon.com/kms/latest/developerguide/tagging-keys.html
// [shared properties]: https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html#mrk-sync-properties
func kms_ReplicateKey(cfg aws.Config, client *kms.Client) {
	input := &kms.ReplicateKeyInput{
		// KeyId: *string, // Required
		// ReplicaRegion: *string, // Required
	}

	if len(_kmsKeyId) > 0 {
		input.KeyId = aws.String(_kmsKeyId)
	}
	if len(_kmsReplicaRegion) > 0 {
		input.ReplicaRegion = aws.String(_kmsReplicaRegion)
	}
	if len(_kmsBypassPolicyLockoutSafetyCheck) > 0 {
		if err := assignInputField(input, "BypassPolicyLockoutSafetyCheck", _kmsBypassPolicyLockoutSafetyCheck); err != nil {
			log.Errorf("invalid --bypass-policy-lockout-safety-check: %s", err.Error())
			return
		}
	}
	if len(_kmsDescription) > 0 {
		input.Description = aws.String(_kmsDescription)
	}
	if len(_kmsPolicy) > 0 {
		input.Policy = aws.String(_kmsPolicy)
	}
	if len(_kmsTags) > 0 {
		if err := assignInputField(input, "Tags", _kmsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.ReplicateKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a grant. Typically, you retire a grant when you no longer need its
// permissions. To identify the grant to retire, use a [grant token], or both the grant ID and
// a key identifier (key ID or key ARN) of the KMS key. The CreateGrantoperation returns both
// values.
//
// This operation can be called by the retiring principal for a grant, by the
// grantee principal if the grant allows the RetireGrant operation, and by the
// Amazon Web Services account in which the grant is created. It can also be called
// by principals to whom permission for retiring a grant is delegated.
//
// For detailed information about grants, including grant terminology, see [Grants in KMS] in the
// Key Management Service Developer Guide . For examples of creating grants in
// several programming languages, see [Use CreateGrant with an Amazon Web Services SDK or CLI].
//
// Cross-account use: Yes. You can retire a grant on a KMS key in a different
// Amazon Web Services account.
//
// Required permissions: Permission to retire a grant is determined primarily by
// the grant. For details, see [Retiring and revoking grants]in the Key Management Service Developer Guide.
//
// Related operations:
//
// # CreateGrant
//
// # ListGrants
//
// # ListRetirableGrants
//
// # RevokeGrant
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [grant token]: https://docs.aws.amazon.com/kms/latest/developerguide/grants.html#grant_token
// [Retiring and revoking grants]: https://docs.aws.amazon.com/kms/latest/developerguide/grant-delete.html
// [Grants in KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/grants.html
// [Use CreateGrant with an Amazon Web Services SDK or CLI]: https://docs.aws.amazon.com/kms/latest/developerguide/example_kms_CreateGrant_section.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
func kms_RetireGrant(cfg aws.Config, client *kms.Client) {
	input := &kms.RetireGrantInput{}

	if len(_kmsDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _kmsDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}
	if len(_kmsGrantId) > 0 {
		input.GrantId = aws.String(_kmsGrantId)
	}
	if len(_kmsGrantToken) > 0 {
		input.GrantToken = aws.String(_kmsGrantToken)
	}
	if len(_kmsKeyId) > 0 {
		input.KeyId = aws.String(_kmsKeyId)
	}

	if resp, err := client.RetireGrant(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified grant. You revoke a grant to terminate the permissions
// that the grant allows. For more information, see [Retiring and revoking grants]in the Key Management Service
// Developer Guide .
//
// When you create, retire, or revoke a grant, there might be a brief delay,
// usually less than five minutes, until the grant is available throughout KMS.
// This state is known as eventual consistency. For details, see [Eventual consistency]in the Key
// Management Service Developer Guide .
//
// For detailed information about grants, including grant terminology, see [Grants in KMS] in the
// Key Management Service Developer Guide . For examples of creating grants in
// several programming languages, see [Use CreateGrant with an Amazon Web Services SDK or CLI].
//
// Cross-account use: Yes. To perform this operation on a KMS key in a different
// Amazon Web Services account, specify the key ARN in the value of the KeyId
// parameter.
//
// Required permissions: [kms:RevokeGrant] (key policy).
//
// Related operations:
//
// # CreateGrant
//
// # ListGrants
//
// # ListRetirableGrants
//
// # RetireGrant
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/grants.html#terms-eventual-consistency
// [kms:RevokeGrant]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [Retiring and revoking grants]: https://docs.aws.amazon.com/kms/latest/developerguide/grant-delete.html
// [Grants in KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/grants.html
// [Use CreateGrant with an Amazon Web Services SDK or CLI]: https://docs.aws.amazon.com/kms/latest/developerguide/example_kms_CreateGrant_section.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
func kms_RevokeGrant(cfg aws.Config, client *kms.Client) {
	input := &kms.RevokeGrantInput{
		// GrantId: *string, // Required
		// KeyId: *string, // Required
	}

	if len(_kmsGrantId) > 0 {
		input.GrantId = aws.String(_kmsGrantId)
	}
	if len(_kmsKeyId) > 0 {
		input.KeyId = aws.String(_kmsKeyId)
	}
	if len(_kmsDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _kmsDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}

	if resp, err := client.RevokeGrant(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Immediately initiates rotation of the key material of the specified symmetric
// encryption KMS key.
//
// You can perform [on-demand rotation] of the key material in customer managed KMS keys, regardless
// of whether or not [automatic key rotation]is enabled. On-demand rotations do not change existing
// automatic rotation schedules. For example, consider a KMS key that has automatic
// key rotation enabled with a rotation period of 730 days. If the key is scheduled
// to automatically rotate on April 14, 2024, and you perform an on-demand rotation
// on April 10, 2024, the key will automatically rotate, as scheduled, on April 14,
// 2024 and every 730 days thereafter.
//
// You can perform on-demand key rotation a maximum of 25 times per KMS key. You
// can use the KMS console to view the number of remaining on-demand rotations
// available for a KMS key.
//
// You can use GetKeyRotationStatus to identify any in progress on-demand rotations. You can use ListKeyRotations to
// identify the date that completed on-demand rotations were performed. You can
// monitor rotation of the key material for your KMS keys in CloudTrail and Amazon
// CloudWatch.
//
// On-demand key rotation is supported only on symmetric encryption KMS keys. You
// cannot perform on-demand rotation of [asymmetric KMS keys], [HMAC KMS keys], or KMS keys in a [custom key store]. When you initiate
// on-demand key rotation on a symmetric encryption KMS key with imported key
// material, you must have already imported [new key material]and that key material's state should
// be PENDING_ROTATION . Use the ListKeyRotations operation to check the state of
// all key materials associated with a KMS key. To perform on-demand rotation of a
// set of related [multi-Region keys], import new key material in the primary Region key, import the
// same key material in each replica Region key, and invoke the on-demand rotation
// on the primary Region key.
//
// You cannot initiate on-demand rotation of [Amazon Web Services managed KMS keys]. KMS always rotates the key material
// of Amazon Web Services managed keys every year. Rotation of [Amazon Web Services owned KMS keys]is managed by the
// Amazon Web Services service that owns the key.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: No. You cannot perform this operation on a KMS key in a
// different Amazon Web Services account.
//
// Required permissions: [kms:RotateKeyOnDemand] (key policy)
//
// Related operations:
//
// # EnableKeyRotation
//
// # DisableKeyRotation
//
// # GetKeyRotationStatus
//
// # ImportKeyMaterial
//
// # ListKeyRotations
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [new key material]: https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys-import-key-material.html
// [HMAC KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/hmac.html
// [Amazon Web Services managed KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-key
// [on-demand rotation]: https://docs.aws.amazon.com/kms/latest/developerguide/rotating-keys-on-demand.html
// [asymmetric KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html
// [Amazon Web Services owned KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-owned-key
// [automatic key rotation]: https://docs.aws.amazon.com/kms/latest/developerguide/rotating-keys-enable-disable.html
// [kms:RotateKeyOnDemand]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [multi-Region keys]: https://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html#multi-region-rotate
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [custom key store]: https://docs.aws.amazon.com/kms/latest/developerguide/key-store-overview.html
func kms_RotateKeyOnDemand(cfg aws.Config, client *kms.Client) {
	input := &kms.RotateKeyOnDemandInput{
		// KeyId: *string, // Required
	}

	if len(_kmsKeyId) > 0 {
		input.KeyId = aws.String(_kmsKeyId)
	}

	if resp, err := client.RotateKeyOnDemand(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Schedules the deletion of a KMS key. By default, KMS applies a waiting period
// of 30 days, but you can specify a waiting period of 7-30 days. When this
// operation is successful, the key state of the KMS key changes to PendingDeletion
// and the key can't be used in any cryptographic operations. It remains in this
// state for the duration of the waiting period. Before the waiting period ends,
// you can use CancelKeyDeletionto cancel the deletion of the KMS key. After the waiting period
// ends, KMS deletes the KMS key, its key material, and all KMS data associated
// with it, including all aliases that refer to it.
//
// Deleting a KMS key is a destructive and potentially dangerous operation. When a
// KMS key is deleted, all data that was encrypted under the KMS key is
// unrecoverable. (The only exception is a [multi-Region replica key], or an [asymmetric or HMAC KMS key with imported key material].) To prevent the use of a KMS
// key without deleting it, use DisableKey.
//
// You can schedule the deletion of a multi-Region primary key and its replica
// keys at any time. However, KMS will not delete a multi-Region primary key with
// existing replica keys. If you schedule the deletion of a primary key with
// replicas, its key state changes to PendingReplicaDeletion and it cannot be
// replicated or used in cryptographic operations. This status can continue
// indefinitely. When the last of its replicas keys is deleted (not just
// scheduled), the key state of the primary key changes to PendingDeletion and its
// waiting period ( PendingWindowInDays ) begins. For details, see [Deleting multi-Region keys] in the Key
// Management Service Developer Guide.
//
// When KMS [deletes a KMS key from an CloudHSM key store], it makes a best effort to delete the associated key material from
// the associated CloudHSM cluster. However, you might need to manually [delete the orphaned key material]from the
// cluster and its backups. [Deleting a KMS key from an external key store]has no effect on the associated external key. However,
// for both types of custom key stores, deleting a KMS key is destructive and
// irreversible. You cannot decrypt ciphertext encrypted under the KMS key by using
// only its associated external key or CloudHSM key. Also, you cannot recreate a
// KMS key in an external key store by creating a new KMS key with the same key
// material.
//
// For more information about scheduling a KMS key for deletion, see [Deleting KMS keys] in the Key
// Management Service Developer Guide.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: No. You cannot perform this operation on a KMS key in a
// different Amazon Web Services account.
//
// Required permissions: kms:ScheduleKeyDeletion (key policy)
//
// # Related operations
//
// # CancelKeyDeletion
//
// # DisableKey
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [delete the orphaned key material]: https://docs.aws.amazon.com/kms/latest/developerguide/fix-keystore.html#fix-keystore-orphaned-key
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [Deleting a KMS key from an external key store]: https://docs.aws.amazon.com/kms/latest/developerguide/deleting-keys.html#delete-xks-key
// [Deleting multi-Region keys]: https://docs.aws.amazon.com/kms/latest/developerguide/deleting-keys.html#deleting-mrks
// [Deleting KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/deleting-keys.html
// [multi-Region replica key]: https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-delete.html
// [asymmetric or HMAC KMS key with imported key material]: https://docs.aws.amazon.com/kms/latest/developerguide/deleting-keys.html#import-delete-key
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [deletes a KMS key from an CloudHSM key store]: https://docs.aws.amazon.com/kms/latest/developerguide/deleting-keys.html#delete-cmk-keystore
func kms_ScheduleKeyDeletion(cfg aws.Config, client *kms.Client) {
	input := &kms.ScheduleKeyDeletionInput{
		// KeyId: *string, // Required
	}

	if len(_kmsKeyId) > 0 {
		input.KeyId = aws.String(_kmsKeyId)
	}
	if len(_kmsPendingWindowInDays) > 0 {
		if err := assignInputField(input, "PendingWindowInDays", _kmsPendingWindowInDays); err != nil {
			log.Errorf("invalid --pending-window-in-days: %s", err.Error())
			return
		}
	}

	if resp, err := client.ScheduleKeyDeletion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a [digital signature] for a message or message digest by using the private key in an
// asymmetric signing KMS key. To verify the signature, use the Verifyoperation, or use
// the public key in the same asymmetric KMS key outside of KMS. For information
// about asymmetric KMS keys, see [Asymmetric KMS keys]in the Key Management Service Developer Guide.
//
// Digital signatures are generated and verified by using asymmetric key pair,
// such as an RSA, ECC, or ML-DSA pair that is represented by an asymmetric KMS
// key. The key owner (or an authorized user) uses their private key to sign a
// message. Anyone with the public key can verify that the message was signed with
// that particular private key and that the message hasn't changed since it was
// signed.
//
// To use the Sign operation, provide the following information:
//
// - Use the KeyId parameter to identify an asymmetric KMS key with a KeyUsage
// value of SIGN_VERIFY . To get the KeyUsage value of a KMS key, use the DescribeKey
// operation. The caller must have kms:Sign permission on the KMS key.
//
// - Use the Message parameter to specify the message or message digest to sign.
// You can submit messages of up to 4096 bytes. To sign a larger message, generate
// a hash digest of the message, and then provide the hash digest in the Message
// parameter. To indicate whether the message is a full message, a digest, or an
// ML-DSA EXTERNAL_MU, use the MessageType parameter.
//
// - Choose a signing algorithm that is compatible with the KMS key.
//
// When signing a message, be sure to record the KMS key and the signing
// algorithm. This information is required to verify the signature.
//
// Best practices recommend that you limit the time during which any signature is
// effective. This deters an attack where the actor uses a signed message to
// establish validity repeatedly or long after the message is superseded.
// Signatures do not include a timestamp, but you can include a timestamp in the
// signed message to help you detect when its time to refresh the signature.
//
// To verify the signature that this operation generates, use the Verify operation. Or
// use the GetPublicKeyoperation to download the public key and then use the public key to
// verify the signature outside of KMS.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: Yes. To perform this operation with a KMS key in a different
// Amazon Web Services account, specify the key ARN or alias ARN in the value of
// the KeyId parameter.
//
// Required permissions: [kms:Sign] (key policy)
//
// Related operations: Verify
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [Asymmetric KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html
// [kms:Sign]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [digital signature]: https://en.wikipedia.org/wiki/Digital_signature
func kms_Sign(cfg aws.Config, client *kms.Client) {
	input := &kms.SignInput{
		// KeyId: *string, // Required
		// Message: []byte, // Required
		// SigningAlgorithm: types.SigningAlgorithmSpec, // Required
	}

	if len(_kmsKeyId) > 0 {
		input.KeyId = aws.String(_kmsKeyId)
	}
	if len(_kmsMessage) > 0 {
		if err := assignInputField(input, "Message", _kmsMessage); err != nil {
			log.Errorf("invalid --message: %s", err.Error())
			return
		}
	}
	if len(_kmsSigningAlgorithm) > 0 {
		if err := assignInputField(input, "SigningAlgorithm", _kmsSigningAlgorithm); err != nil {
			log.Errorf("invalid --signing-algorithm: %s", err.Error())
			return
		}
	}
	if len(_kmsDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _kmsDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}
	if len(_kmsGrantTokens) > 0 {
		input.GrantTokens = append([]string(nil), _kmsGrantTokens...)
	}
	if len(_kmsMessageType) > 0 {
		if err := assignInputField(input, "MessageType", _kmsMessageType); err != nil {
			log.Errorf("invalid --message-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.Sign(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or edits tags on a [customer managed key].
// Tagging or untagging a KMS key can allow or deny permission to the KMS key. For
// details, see [ABAC for KMS]in the Key Management Service Developer Guide.
//
// Each tag consists of a tag key and a tag value, both of which are
// case-sensitive strings. The tag value can be an empty (null) string. To add a
// tag, specify a new tag key and a tag value. To edit a tag, specify an existing
// tag key and a new tag value.
//
// You can use this operation to tag a [customer managed key], but you cannot tag an [Amazon Web Services managed key], an [Amazon Web Services owned key], a [custom key store], or an [alias].
//
// You can also add tags to a KMS key while creating it (CreateKey ) or replicating it (ReplicateKey ).
//
// For information about using tags in KMS, see [Tagging keys]. For general information about
// tags, including the format and syntax, see [Tagging Amazon Web Services resources]in the Amazon Web Services General
// Reference.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: No. You cannot perform this operation on a KMS key in a
// different Amazon Web Services account.
//
// Required permissions: [kms:TagResource] (key policy)
//
// # Related operations
//
// # CreateKey
//
// # ListResourceTags
//
// # ReplicateKey
//
// # UntagResource
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Amazon Web Services owned key]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-owned-key
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [kms:TagResource]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [customer managed key]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#customer-mgn-key
// [Tagging keys]: https://docs.aws.amazon.com/kms/latest/developerguide/tagging-keys.html
// [alias]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-alias.html
// [ABAC for KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/abac.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [Amazon Web Services managed key]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-key
// [custom key store]: https://docs.aws.amazon.com/kms/latest/developerguide/key-store-overview.html
// [Tagging Amazon Web Services resources]: https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html
func kms_TagResource(cfg aws.Config, client *kms.Client) {
	input := &kms.TagResourceInput{
		// KeyId: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_kmsKeyId) > 0 {
		input.KeyId = aws.String(_kmsKeyId)
	}
	if len(_kmsTags) > 0 {
		if err := assignInputField(input, "Tags", _kmsTags); err != nil {
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

// Deletes tags from a [customer managed key]. To delete a tag, specify the tag key and the KMS key.
// Tagging or untagging a KMS key can allow or deny permission to the KMS key. For
// details, see [ABAC for KMS]in the Key Management Service Developer Guide.
//
// When it succeeds, the UntagResource operation doesn't return any output. Also,
// if the specified tag key isn't found on the KMS key, it doesn't throw an
// exception or return a response. To confirm that the operation worked, use the ListResourceTags
// operation.
//
// For information about using tags in KMS, see [Tagging keys]. For general information about
// tags, including the format and syntax, see [Tagging Amazon Web Services resources]in the Amazon Web Services General
// Reference.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: No. You cannot perform this operation on a KMS key in a
// different Amazon Web Services account.
//
// Required permissions: [kms:UntagResource] (key policy)
//
// # Related operations
//
// # CreateKey
//
// # ListResourceTags
//
// # ReplicateKey
//
// # TagResource
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [kms:UntagResource]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [Tagging keys]: https://docs.aws.amazon.com/kms/latest/developerguide/tagging-keys.html
// [ABAC for KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/abac.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [Tagging Amazon Web Services resources]: https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html
// [customer managed key]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#customer-mgn-key
func kms_UntagResource(cfg aws.Config, client *kms.Client) {
	input := &kms.UntagResourceInput{
		// KeyId: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_kmsKeyId) > 0 {
		input.KeyId = aws.String(_kmsKeyId)
	}
	if len(_kmsTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _kmsTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates an existing KMS alias with a different KMS key. Each alias is
// associated with only one KMS key at a time, although a KMS key can have multiple
// aliases. The alias and the KMS key must be in the same Amazon Web Services
// account and Region.
//
// Adding, deleting, or updating an alias can allow or deny permission to the KMS
// key. For details, see [ABAC for KMS]in the Key Management Service Developer Guide.
//
// The current and new KMS key must be the same type (both symmetric or both
// asymmetric or both HMAC), and they must have the same key usage. This
// restriction prevents errors in code that uses aliases. If you must assign an
// alias to a different type of KMS key, use DeleteAliasto delete the old alias and CreateAlias to
// create a new alias.
//
// You cannot use UpdateAlias to change an alias name. To change an alias name,
// use DeleteAliasto delete the old alias and CreateAlias to create a new alias.
//
// Because an alias is not a property of a KMS key, you can create, update, and
// delete the aliases of a KMS key without affecting the KMS key. Also, aliases do
// not appear in the response from the DescribeKeyoperation. To get the aliases of all KMS
// keys in the account, use the ListAliasesoperation.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: No. You cannot perform this operation on a KMS key in a
// different Amazon Web Services account.
//
// # Required permissions
//
// [kms:UpdateAlias]
// - on the alias (IAM policy).
//
// [kms:UpdateAlias]
// - on the current KMS key (key policy).
//
// [kms:UpdateAlias]
// - on the new KMS key (key policy).
//
// For details, see [Controlling access to aliases] in the Key Management Service Developer Guide.
//
// Related operations:
//
// # CreateAlias
//
// # DeleteAlias
//
// # ListAliases
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [ABAC for KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/abac.html
// [kms:UpdateAlias]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [Controlling access to aliases]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-alias.html#alias-access
func kms_UpdateAlias(cfg aws.Config, client *kms.Client) {
	input := &kms.UpdateAliasInput{
		// AliasName: *string, // Required
		// TargetKeyId: *string, // Required
	}

	if len(_kmsAliasName) > 0 {
		input.AliasName = aws.String(_kmsAliasName)
	}
	if len(_kmsTargetKeyId) > 0 {
		input.TargetKeyId = aws.String(_kmsTargetKeyId)
	}

	if resp, err := client.UpdateAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes the properties of a custom key store. You can use this operation to
// change the properties of an CloudHSM key store or an external key store.
//
// Use the required CustomKeyStoreId parameter to identify the custom key store.
// Use the remaining optional parameters to change its properties. This operation
// does not return any property values. To verify the updated property values, use
// the DescribeCustomKeyStoresoperation.
//
// This operation is part of the custom key stores feature in KMS, which combines
// the convenience and extensive integration of KMS with the isolation and control
// of a key store that you own and manage.
//
// When updating the properties of an external key store, verify that the updated
// settings connect your key store, via the external key store proxy, to the same
// external key manager as the previous settings, or to a backup or snapshot of the
// external key manager with the same cryptographic keys. If the updated connection
// settings fail, you can fix them and retry, although an extended delay might
// disrupt Amazon Web Services services. However, if KMS permanently loses its
// access to cryptographic keys, ciphertext encrypted under those keys is
// unrecoverable.
//
// For external key stores:
//
// Some external key managers provide a simpler method for updating an external
// key store. For details, see your external key manager documentation.
//
// When updating an external key store in the KMS console, you can upload a
// JSON-based proxy configuration file with the desired values. You cannot upload
// the proxy configuration file to the UpdateCustomKeyStore operation. However,
// you can use the file to help you determine the correct values for the
// UpdateCustomKeyStore parameters.
//
// For an CloudHSM key store, you can use this operation to change the custom key
// store friendly name ( NewCustomKeyStoreName ), to tell KMS about a change to the
// kmsuser crypto user password ( KeyStorePassword ), or to associate the custom
// key store with a different, but related, CloudHSM cluster ( CloudHsmClusterId ).
// To update any property of an CloudHSM key store, the ConnectionState of the
// CloudHSM key store must be DISCONNECTED .
//
// For an external key store, you can use this operation to change the custom key
// store friendly name ( NewCustomKeyStoreName ), or to tell KMS about a change to
// the external key store proxy authentication credentials (
// XksProxyAuthenticationCredential ), connection method ( XksProxyConnectivity ),
// external proxy endpoint ( XksProxyUriEndpoint ) and path ( XksProxyUriPath ).
// For external key stores with an XksProxyConnectivity of VPC_ENDPOINT_SERVICE ,
// you can also update the Amazon VPC endpoint service name (
// XksProxyVpcEndpointServiceName ). To update most properties of an external key
// store, the ConnectionState of the external key store must be DISCONNECTED .
// However, you can update the CustomKeyStoreName ,
// XksProxyAuthenticationCredential , and XksProxyUriPath of an external key store
// when it is in the CONNECTED or DISCONNECTED state.
//
// If your update requires a DISCONNECTED state, before using UpdateCustomKeyStore
// , use the DisconnectCustomKeyStoreoperation to disconnect the custom key store. After the
// UpdateCustomKeyStore operation completes, use the ConnectCustomKeyStore to reconnect the custom key
// store. To find the ConnectionState of the custom key store, use the DescribeCustomKeyStores operation.
//
// Before updating the custom key store, verify that the new values allow KMS to
// connect the custom key store to its backing key store. For example, before you
// change the XksProxyUriPath value, verify that the external key store proxy is
// reachable at the new path.
//
// If the operation succeeds, it returns a JSON object with no properties.
//
// Cross-account use: No. You cannot perform this operation on a custom key store
// in a different Amazon Web Services account.
//
// Required permissions: [kms:UpdateCustomKeyStore] (IAM policy)
//
// Related operations:
//
// # ConnectCustomKeyStore
//
// # CreateCustomKeyStore
//
// # DeleteCustomKeyStore
//
// # DescribeCustomKeyStores
//
// # DisconnectCustomKeyStore
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [kms:UpdateCustomKeyStore]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
func kms_UpdateCustomKeyStore(cfg aws.Config, client *kms.Client) {
	input := &kms.UpdateCustomKeyStoreInput{
		// CustomKeyStoreId: *string, // Required
	}

	if len(_kmsCustomKeyStoreId) > 0 {
		input.CustomKeyStoreId = aws.String(_kmsCustomKeyStoreId)
	}
	if len(_kmsCloudHsmClusterId) > 0 {
		input.CloudHsmClusterId = aws.String(_kmsCloudHsmClusterId)
	}
	if len(_kmsKeyStorePassword) > 0 {
		input.KeyStorePassword = aws.String(_kmsKeyStorePassword)
	}
	if len(_kmsNewCustomKeyStoreName) > 0 {
		input.NewCustomKeyStoreName = aws.String(_kmsNewCustomKeyStoreName)
	}
	if len(_kmsXksProxyAuthenticationCredential) > 0 {
		if err := assignInputField(input, "XksProxyAuthenticationCredential", _kmsXksProxyAuthenticationCredential); err != nil {
			log.Errorf("invalid --xks-proxy-authentication-credential: %s", err.Error())
			return
		}
	}
	if len(_kmsXksProxyConnectivity) > 0 {
		if err := assignInputField(input, "XksProxyConnectivity", _kmsXksProxyConnectivity); err != nil {
			log.Errorf("invalid --xks-proxy-connectivity: %s", err.Error())
			return
		}
	}
	if len(_kmsXksProxyUriEndpoint) > 0 {
		input.XksProxyUriEndpoint = aws.String(_kmsXksProxyUriEndpoint)
	}
	if len(_kmsXksProxyUriPath) > 0 {
		input.XksProxyUriPath = aws.String(_kmsXksProxyUriPath)
	}
	if len(_kmsXksProxyVpcEndpointServiceName) > 0 {
		input.XksProxyVpcEndpointServiceName = aws.String(_kmsXksProxyVpcEndpointServiceName)
	}
	if len(_kmsXksProxyVpcEndpointServiceOwner) > 0 {
		input.XksProxyVpcEndpointServiceOwner = aws.String(_kmsXksProxyVpcEndpointServiceOwner)
	}

	if resp, err := client.UpdateCustomKeyStore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the description of a KMS key. To see the description of a KMS key, use DescribeKey
// .
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: No. You cannot perform this operation on a KMS key in a
// different Amazon Web Services account.
//
// Required permissions: [kms:UpdateKeyDescription] (key policy)
//
// # Related operations
//
// # CreateKey
//
// # DescribeKey
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [kms:UpdateKeyDescription]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
func kms_UpdateKeyDescription(cfg aws.Config, client *kms.Client) {
	input := &kms.UpdateKeyDescriptionInput{
		// Description: *string, // Required
		// KeyId: *string, // Required
	}

	if len(_kmsDescription) > 0 {
		input.Description = aws.String(_kmsDescription)
	}
	if len(_kmsKeyId) > 0 {
		input.KeyId = aws.String(_kmsKeyId)
	}

	if resp, err := client.UpdateKeyDescription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes the primary key of a multi-Region key.
// This operation changes the replica key in the specified Region to a primary key
// and changes the former primary key to a replica key. For example, suppose you
// have a primary key in us-east-1 and a replica key in eu-west-2 . If you run
// UpdatePrimaryRegion with a PrimaryRegion value of eu-west-2 , the primary key is
// now the key in eu-west-2 , and the key in us-east-1 becomes a replica key. For
// details, see [Change the primary key in a set of multi-Region keys]in the Key Management Service Developer Guide.
//
// This operation supports multi-Region keys, an KMS feature that lets you create
// multiple interoperable KMS keys in different Amazon Web Services Regions.
// Because these KMS keys have the same key ID, key material, and other metadata,
// you can use them interchangeably to encrypt data in one Amazon Web Services
// Region and decrypt it in a different Amazon Web Services Region without
// re-encrypting the data or making a cross-Region call. For more information about
// multi-Region keys, see [Multi-Region keys in KMS]in the Key Management Service Developer Guide.
//
// The primary key of a multi-Region key is the source for properties that are
// always shared by primary and replica keys, including the key material, [key ID], [key spec], [key usage], [key material origin],
// and [automatic key rotation]. It's the only key that can be replicated. You cannot [delete the primary key] until all replica
// keys are deleted.
//
// The key ID and primary Region that you specify uniquely identify the replica
// key that will become the primary key. The primary Region must already have a
// replica key. This operation does not create a KMS key in the specified Region.
// To find the replica keys, use the DescribeKeyoperation on the primary key or any replica
// key. To create a replica key, use the ReplicateKeyoperation.
//
// You can run this operation while using the affected multi-Region keys in
// cryptographic operations. This operation should not delay, interrupt, or cause
// failures in cryptographic operations.
//
// Even after this operation completes, the process of updating the primary Region
// might still be in progress for a few more seconds. Operations such as
// DescribeKey might display both the old and new primary keys as replicas. The old
// and new primary keys have a transient key state of Updating . The original key
// state is restored when the update is complete. While the key state is Updating ,
// you can use the keys in cryptographic operations, but you cannot replicate the
// new primary key or perform certain management operations, such as enabling or
// disabling these keys. For details about the Updating key state, see [Key states of KMS keys] in the Key
// Management Service Developer Guide.
//
// This operation does not return any output. To verify that primary key is
// changed, use the DescribeKeyoperation.
//
// Cross-account use: No. You cannot use this operation in a different Amazon Web
// Services account.
//
// Required permissions:
//
// - kms:UpdatePrimaryRegion on the current primary key (in the primary key's
// Region). Include this permission primary key's key policy.
//
// - kms:UpdatePrimaryRegion on the current replica key (in the replica key's
// Region). Include this permission in the replica key's key policy.
//
// # Related operations
//
// # CreateKey
//
// # ReplicateKey
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [key ID]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id-key-id
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [delete the primary key]: https://docs.aws.amazon.com/kms/latest/APIReference/API_ScheduleKeyDeletion.html
// [Change the primary key in a set of multi-Region keys]: https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-update.html
// [key usage]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-usage
// [Multi-Region keys in KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html
// [key spec]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-spec
// [key material origin]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-origin
// [automatic key rotation]: https://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
func kms_UpdatePrimaryRegion(cfg aws.Config, client *kms.Client) {
	input := &kms.UpdatePrimaryRegionInput{
		// KeyId: *string, // Required
		// PrimaryRegion: *string, // Required
	}

	if len(_kmsKeyId) > 0 {
		input.KeyId = aws.String(_kmsKeyId)
	}
	if len(_kmsPrimaryRegion) > 0 {
		input.PrimaryRegion = aws.String(_kmsPrimaryRegion)
	}

	if resp, err := client.UpdatePrimaryRegion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Verifies a digital signature that was generated by the Sign operation.
// Verification confirms that an authorized user signed the message with the
// specified KMS key and signing algorithm, and the message hasn't changed since it
// was signed. If the signature is verified, the value of the SignatureValid field
// in the response is True . If the signature verification fails, the Verify
// operation fails with an KMSInvalidSignatureException exception.
//
// A digital signature is generated by using the private key in an asymmetric KMS
// key. The signature is verified by using the public key in the same asymmetric
// KMS key. For information about asymmetric KMS keys, see [Asymmetric KMS keys]in the Key Management
// Service Developer Guide.
//
// To use the Verify operation, specify the same asymmetric KMS key, message, and
// signing algorithm that were used to produce the signature. The message type does
// not need to be the same as the one used for signing, but it must indicate
// whether the value of the Message parameter should be hashed as part of the
// verification process.
//
// You can also verify the digital signature by using the public key of the KMS
// key outside of KMS. Use the GetPublicKeyoperation to download the public key in the
// asymmetric KMS key and then use the public key to verify the signature outside
// of KMS. The advantage of using the Verify operation is that it is performed
// within KMS. As a result, it's easy to call, the operation is performed within
// the FIPS boundary, it is logged in CloudTrail, and you can use key policy and
// IAM policy to determine who is authorized to use the KMS key to verify
// signatures.
//
// To verify a signature outside of KMS with an SM2 public key (China Regions
// only), you must specify the distinguishing ID. By default, KMS uses
// 1234567812345678 as the distinguishing ID. For more information, see [Offline verification with SM2 key pairs].
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: Yes. To perform this operation with a KMS key in a different
// Amazon Web Services account, specify the key ARN or alias ARN in the value of
// the KeyId parameter.
//
// Required permissions: [kms:Verify] (key policy)
//
// Related operations: Sign
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [Asymmetric KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html
// [Offline verification with SM2 key pairs]: https://docs.aws.amazon.com/kms/latest/developerguide/offline-operations.html#key-spec-sm-offline-verification
// [kms:Verify]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
func kms_Verify(cfg aws.Config, client *kms.Client) {
	input := &kms.VerifyInput{
		// KeyId: *string, // Required
		// Message: []byte, // Required
		// Signature: []byte, // Required
		// SigningAlgorithm: types.SigningAlgorithmSpec, // Required
	}

	if len(_kmsKeyId) > 0 {
		input.KeyId = aws.String(_kmsKeyId)
	}
	if len(_kmsMessage) > 0 {
		if err := assignInputField(input, "Message", _kmsMessage); err != nil {
			log.Errorf("invalid --message: %s", err.Error())
			return
		}
	}
	if len(_kmsSignature) > 0 {
		if err := assignInputField(input, "Signature", _kmsSignature); err != nil {
			log.Errorf("invalid --signature: %s", err.Error())
			return
		}
	}
	if len(_kmsSigningAlgorithm) > 0 {
		if err := assignInputField(input, "SigningAlgorithm", _kmsSigningAlgorithm); err != nil {
			log.Errorf("invalid --signing-algorithm: %s", err.Error())
			return
		}
	}
	if len(_kmsDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _kmsDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}
	if len(_kmsGrantTokens) > 0 {
		input.GrantTokens = append([]string(nil), _kmsGrantTokens...)
	}
	if len(_kmsMessageType) > 0 {
		if err := assignInputField(input, "MessageType", _kmsMessageType); err != nil {
			log.Errorf("invalid --message-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.Verify(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Verifies the hash-based message authentication code (HMAC) for a specified
// message, HMAC KMS key, and MAC algorithm. To verify the HMAC, VerifyMac
// computes an HMAC using the message, HMAC KMS key, and MAC algorithm that you
// specify, and compares the computed HMAC to the HMAC that you specify. If the
// HMACs are identical, the verification succeeds; otherwise, it fails.
// Verification indicates that the message hasn't changed since the HMAC was
// calculated, and the specified key was used to generate and verify the HMAC.
//
// HMAC KMS keys and the HMAC algorithms that KMS uses conform to industry
// standards defined in [RFC 2104].
//
// This operation is part of KMS support for HMAC KMS keys. For details, see [HMAC keys in KMS] in
// the Key Management Service Developer Guide.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: Yes. To perform this operation with a KMS key in a different
// Amazon Web Services account, specify the key ARN or alias ARN in the value of
// the KeyId parameter.
//
// Required permissions: [kms:VerifyMac] (key policy)
//
// Related operations: GenerateMac
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [RFC 2104]: https://datatracker.ietf.org/doc/html/rfc2104
// [kms:VerifyMac]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [HMAC keys in KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/hmac.html
func kms_VerifyMac(cfg aws.Config, client *kms.Client) {
	input := &kms.VerifyMacInput{
		// KeyId: *string, // Required
		// Mac: []byte, // Required
		// MacAlgorithm: types.MacAlgorithmSpec, // Required
		// Message: []byte, // Required
	}

	if len(_kmsKeyId) > 0 {
		input.KeyId = aws.String(_kmsKeyId)
	}
	if len(_kmsMac) > 0 {
		if err := assignInputField(input, "Mac", _kmsMac); err != nil {
			log.Errorf("invalid --mac: %s", err.Error())
			return
		}
	}
	if len(_kmsMacAlgorithm) > 0 {
		if err := assignInputField(input, "MacAlgorithm", _kmsMacAlgorithm); err != nil {
			log.Errorf("invalid --mac-algorithm: %s", err.Error())
			return
		}
	}
	if len(_kmsMessage) > 0 {
		if err := assignInputField(input, "Message", _kmsMessage); err != nil {
			log.Errorf("invalid --message: %s", err.Error())
			return
		}
	}
	if len(_kmsDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _kmsDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}
	if len(_kmsGrantTokens) > 0 {
		input.GrantTokens = append([]string(nil), _kmsGrantTokens...)
	}

	if resp, err := client.VerifyMac(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_kmsCmd)
	_kmsCmd.Flags().SortFlags = false

	_kmsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_kmsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_kmsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_kmsCmd.Flags().StringVarP(&_kmsAliasName, "alias-name", "", "", "Alias Name")
	_kmsCmd.Flags().StringVarP(&_kmsBypassPolicyLockoutSafetyCheck, "bypass-policy-lockout-safety-check", "", "", "Bypass Policy Lockout Safety Check")
	_kmsCmd.Flags().StringVarP(&_kmsCiphertextBlob, "ciphertext-blob", "", "", "Ciphertext Blob")
	_kmsCmd.Flags().StringVarP(&_kmsCloudHsmClusterId, "cloud-hsm-cluster-id", "", "", "Cloud Hsm Cluster ID")
	_kmsCmd.Flags().StringVarP(&_kmsConstraints, "constraints", "", "", "Constraints")
	_kmsCmd.Flags().StringVarP(&_kmsCustomKeyStoreId, "custom-key-store-id", "", "", "Custom Key Store ID")
	_kmsCmd.Flags().StringVarP(&_kmsCustomKeyStoreName, "custom-key-store-name", "", "", "Custom Key Store Name")
	_kmsCmd.Flags().StringVarP(&_kmsCustomKeyStoreType, "custom-key-store-type", "", "", "Custom Key Store Type")
	_kmsCmd.Flags().StringVarP(&_kmsCustomerMasterKeySpec, "customer-master-key-spec", "", "", "Customer Master Key Spec")
	_kmsCmd.Flags().StringVarP(&_kmsDescription, "description", "", "", "Description")
	_kmsCmd.Flags().StringVarP(&_kmsDestinationEncryptionAlgorithm, "destination-encryption-algorithm", "", "", "Destination Encryption Algorithm")
	_kmsCmd.Flags().StringVarP(&_kmsDestinationEncryptionContext, "destination-encryption-context", "", "", "Destination Encryption Context")
	_kmsCmd.Flags().StringVarP(&_kmsDestinationKeyId, "destination-key-id", "", "", "Destination Key ID")
	_kmsCmd.Flags().StringVarP(&_kmsDryRun, "dry-run", "", "", "Dry Run")
	_kmsCmd.Flags().StringVarP(&_kmsDryRunModifiers, "dry-run-modifiers", "", "", "Dry Run Modifiers")
	_kmsCmd.Flags().StringVarP(&_kmsEncryptedKeyMaterial, "encrypted-key-material", "", "", "Encrypted Key Material")
	_kmsCmd.Flags().StringVarP(&_kmsEncryptionAlgorithm, "encryption-algorithm", "", "", "Encryption Algorithm")
	_kmsCmd.Flags().StringVarP(&_kmsEncryptionContext, "encryption-context", "", "", "Encryption Context")
	_kmsCmd.Flags().StringVarP(&_kmsExpirationModel, "expiration-model", "", "", "Expiration Model")
	_kmsCmd.Flags().StringVarP(&_kmsGrantId, "grant-id", "", "", "Grant ID")
	_kmsCmd.Flags().StringVarP(&_kmsGrantToken, "grant-token", "", "", "Grant Token")
	_kmsCmd.Flags().StringSliceVarP(&_kmsGrantTokens, "grant-tokens", "", nil, "Grant Tokens")
	_kmsCmd.Flags().StringVarP(&_kmsGranteePrincipal, "grantee-principal", "", "", "Grantee Principal")
	_kmsCmd.Flags().StringVarP(&_kmsImportToken, "import-token", "", "", "Import Token")
	_kmsCmd.Flags().StringVarP(&_kmsImportType, "import-type", "", "", "Import Type")
	_kmsCmd.Flags().StringVarP(&_kmsIncludeKeyMaterial, "include-key-material", "", "", "Include Key Material")
	_kmsCmd.Flags().StringVarP(&_kmsKeyAgreementAlgorithm, "key-agreement-algorithm", "", "", "Key Agreement Algorithm")
	_kmsCmd.Flags().StringVarP(&_kmsKeyId, "key-id", "", "", "Key ID")
	_kmsCmd.Flags().StringVarP(&_kmsKeyMaterialDescription, "key-material-description", "", "", "Key Material Description")
	_kmsCmd.Flags().StringVarP(&_kmsKeyMaterialId, "key-material-id", "", "", "Key Material ID")
	_kmsCmd.Flags().StringVarP(&_kmsKeyPairSpec, "key-pair-spec", "", "", "Key Pair Spec")
	_kmsCmd.Flags().StringVarP(&_kmsKeySpec, "key-spec", "", "", "Key Spec")
	_kmsCmd.Flags().StringVarP(&_kmsKeyStorePassword, "key-store-password", "", "", "Key Store Password")
	_kmsCmd.Flags().StringVarP(&_kmsKeyUsage, "key-usage", "", "", "Key Usage")
	_kmsCmd.Flags().StringVarP(&_kmsLimit, "limit", "", "", "Limit")
	_kmsCmd.Flags().StringVarP(&_kmsMac, "mac", "", "", "Mac")
	_kmsCmd.Flags().StringVarP(&_kmsMacAlgorithm, "mac-algorithm", "", "", "Mac Algorithm")
	_kmsCmd.Flags().StringVarP(&_kmsMarker, "marker", "", "", "Marker")
	_kmsCmd.Flags().StringVarP(&_kmsMessage, "message", "", "", "Message")
	_kmsCmd.Flags().StringVarP(&_kmsMessageType, "message-type", "", "", "Message Type")
	_kmsCmd.Flags().StringVarP(&_kmsMultiRegion, "multi-region", "", "", "Multi Region")
	_kmsCmd.Flags().StringVarP(&_kmsName, "name", "", "", "Name")
	_kmsCmd.Flags().StringVarP(&_kmsNewCustomKeyStoreName, "new-custom-key-store-name", "", "", "New Custom Key Store Name")
	_kmsCmd.Flags().StringVarP(&_kmsNumberOfBytes, "number-of-bytes", "", "", "Number Of Bytes")
	_kmsCmd.Flags().StringVarP(&_kmsOperations, "operations", "", "", "Operations")
	_kmsCmd.Flags().StringVarP(&_kmsOrigin, "origin", "", "", "Origin")
	_kmsCmd.Flags().StringVarP(&_kmsPendingWindowInDays, "pending-window-in-days", "", "", "Pending Window In Days")
	_kmsCmd.Flags().StringVarP(&_kmsPlaintext, "plaintext", "", "", "Plaintext")
	_kmsCmd.Flags().StringVarP(&_kmsPolicy, "policy", "", "", "Policy")
	_kmsCmd.Flags().StringVarP(&_kmsPolicyName, "policy-name", "", "", "Policy Name")
	_kmsCmd.Flags().StringVarP(&_kmsPrimaryRegion, "primary-region", "", "", "Primary Region")
	_kmsCmd.Flags().StringVarP(&_kmsPublicKey, "public-key", "", "", "Public Key")
	_kmsCmd.Flags().StringVarP(&_kmsRecipient, "recipient", "", "", "Recipient")
	_kmsCmd.Flags().StringVarP(&_kmsReplicaRegion, "replica-region", "", "", "Replica Region")
	_kmsCmd.Flags().StringVarP(&_kmsRetiringPrincipal, "retiring-principal", "", "", "Retiring Principal")
	_kmsCmd.Flags().StringVarP(&_kmsRotationPeriodInDays, "rotation-period-in-days", "", "", "Rotation Period In Days")
	_kmsCmd.Flags().StringVarP(&_kmsSignature, "signature", "", "", "Signature")
	_kmsCmd.Flags().StringVarP(&_kmsSigningAlgorithm, "signing-algorithm", "", "", "Signing Algorithm")
	_kmsCmd.Flags().StringVarP(&_kmsSourceEncryptionAlgorithm, "source-encryption-algorithm", "", "", "Source Encryption Algorithm")
	_kmsCmd.Flags().StringVarP(&_kmsSourceEncryptionContext, "source-encryption-context", "", "", "Source Encryption Context")
	_kmsCmd.Flags().StringVarP(&_kmsSourceKeyId, "source-key-id", "", "", "Source Key ID")
	_kmsCmd.Flags().StringSliceVarP(&_kmsTagKeys, "tag-keys", "", nil, "Tag Keys")
	_kmsCmd.Flags().StringVarP(&_kmsTags, "tags", "", "", "Tags")
	_kmsCmd.Flags().StringVarP(&_kmsTargetKeyId, "target-key-id", "", "", "Target Key ID")
	_kmsCmd.Flags().StringVarP(&_kmsTrustAnchorCertificate, "trust-anchor-certificate", "", "", "Trust Anchor Certificate")
	_kmsCmd.Flags().StringVarP(&_kmsValidTo, "valid-to", "", "", "Valid To")
	_kmsCmd.Flags().StringVarP(&_kmsWrappingAlgorithm, "wrapping-algorithm", "", "", "Wrapping Algorithm")
	_kmsCmd.Flags().StringVarP(&_kmsWrappingKeySpec, "wrapping-key-spec", "", "", "Wrapping Key Spec")
	_kmsCmd.Flags().StringVarP(&_kmsXksKeyId, "xks-key-id", "", "", "Xks Key ID")
	_kmsCmd.Flags().StringVarP(&_kmsXksProxyAuthenticationCredential, "xks-proxy-authentication-credential", "", "", "Xks Proxy Authentication Credential")
	_kmsCmd.Flags().StringVarP(&_kmsXksProxyConnectivity, "xks-proxy-connectivity", "", "", "Xks Proxy Connectivity")
	_kmsCmd.Flags().StringVarP(&_kmsXksProxyUriEndpoint, "xks-proxy-uri-endpoint", "", "", "Xks Proxy URI Endpoint")
	_kmsCmd.Flags().StringVarP(&_kmsXksProxyUriPath, "xks-proxy-uri-path", "", "", "Xks Proxy URI Path")
	_kmsCmd.Flags().StringVarP(&_kmsXksProxyVpcEndpointServiceName, "xks-proxy-vpc-endpoint-service-name", "", "", "Xks Proxy VPC Endpoint Service Name")
	_kmsCmd.Flags().StringVarP(&_kmsXksProxyVpcEndpointServiceOwner, "xks-proxy-vpc-endpoint-service-owner", "", "", "Xks Proxy VPC Endpoint Service Owner")

	_kmsCmd.Flags().BoolVarP(&_kmsCancelKeyDeletion, "cancel-key-deletion", "", false, "Cancel Key Deletion")
	_kmsCmd.Flags().BoolVarP(&_kmsConnectCustomKeyStore, "connect-custom-key-store", "", false, "Connect Custom Key Store")
	_kmsCmd.Flags().BoolVarP(&_kmsCreateAlias, "create-alias", "", false, "Create Alias")
	_kmsCmd.Flags().BoolVarP(&_kmsCreateCustomKeyStore, "create-custom-key-store", "", false, "Create Custom Key Store")
	_kmsCmd.Flags().BoolVarP(&_kmsCreateGrant, "create-grant", "", false, "Create Grant")
	_kmsCmd.Flags().BoolVarP(&_kmsCreateKey, "create-key", "", false, "Create Key")
	_kmsCmd.Flags().BoolVarP(&_kmsDecrypt, "decrypt", "", false, "Decrypt")
	_kmsCmd.Flags().BoolVarP(&_kmsDeleteAlias, "delete-alias", "", false, "Delete Alias")
	_kmsCmd.Flags().BoolVarP(&_kmsDeleteCustomKeyStore, "delete-custom-key-store", "", false, "Delete Custom Key Store")
	_kmsCmd.Flags().BoolVarP(&_kmsDeleteImportedKeyMaterial, "delete-imported-key-material", "", false, "Delete Imported Key Material")
	_kmsCmd.Flags().BoolVarP(&_kmsDeriveSharedSecret, "derive-shared-secret", "", false, "Derive Shared Secret")
	_kmsCmd.Flags().BoolVarP(&_kmsDescribeCustomKeyStores, "describe-custom-key-stores", "", false, "Describe Custom Key Stores")
	_kmsCmd.Flags().BoolVarP(&_kmsDescribeKey, "describe-key", "", false, "Describe Key")
	_kmsCmd.Flags().BoolVarP(&_kmsDisableKey, "disable-key", "", false, "Disable Key")
	_kmsCmd.Flags().BoolVarP(&_kmsDisableKeyRotation, "disable-key-rotation", "", false, "Disable Key Rotation")
	_kmsCmd.Flags().BoolVarP(&_kmsDisconnectCustomKeyStore, "disconnect-custom-key-store", "", false, "Disconnect Custom Key Store")
	_kmsCmd.Flags().BoolVarP(&_kmsEnableKey, "enable-key", "", false, "Enable Key")
	_kmsCmd.Flags().BoolVarP(&_kmsEnableKeyRotation, "enable-key-rotation", "", false, "Enable Key Rotation")
	_kmsCmd.Flags().BoolVarP(&_kmsEncrypt, "encrypt", "", false, "Encrypt")
	_kmsCmd.Flags().BoolVarP(&_kmsGenerateDataKey, "generate-data-key", "", false, "Generate Data Key")
	_kmsCmd.Flags().BoolVarP(&_kmsGenerateDataKeyPair, "generate-data-key-pair", "", false, "Generate Data Key Pair")
	_kmsCmd.Flags().BoolVarP(&_kmsGenerateDataKeyPairWithoutPlaintext, "generate-data-key-pair-without-plaintext", "", false, "Generate Data Key Pair Without Plaintext")
	_kmsCmd.Flags().BoolVarP(&_kmsGenerateDataKeyWithoutPlaintext, "generate-data-key-without-plaintext", "", false, "Generate Data Key Without Plaintext")
	_kmsCmd.Flags().BoolVarP(&_kmsGenerateMac, "generate-mac", "", false, "Generate Mac")
	_kmsCmd.Flags().BoolVarP(&_kmsGenerateRandom, "generate-random", "", false, "Generate Random")
	_kmsCmd.Flags().BoolVarP(&_kmsGetKeyPolicy, "get-key-policy", "", false, "Get Key Policy")
	_kmsCmd.Flags().BoolVarP(&_kmsGetKeyRotationStatus, "get-key-rotation-status", "", false, "Get Key Rotation Status")
	_kmsCmd.Flags().BoolVarP(&_kmsGetParametersForImport, "get-parameters-for-import", "", false, "Get Parameters For Import")
	_kmsCmd.Flags().BoolVarP(&_kmsGetPublicKey, "get-public-key", "", false, "Get Public Key")
	_kmsCmd.Flags().BoolVarP(&_kmsImportKeyMaterial, "import-key-material", "", false, "Import Key Material")
	_kmsCmd.Flags().BoolVarP(&_kmsListAliases, "list-aliases", "", false, "List Aliases")
	_kmsCmd.Flags().BoolVarP(&_kmsListGrants, "list-grants", "", false, "List Grants")
	_kmsCmd.Flags().BoolVarP(&_kmsListKeyPolicies, "list-key-policies", "", false, "List Key Policies")
	_kmsCmd.Flags().BoolVarP(&_kmsListKeyRotations, "list-key-rotations", "", false, "List Key Rotations")
	_kmsCmd.Flags().BoolVarP(&_kmsListKeys, "list-keys", "", false, "List Keys")
	_kmsCmd.Flags().BoolVarP(&_kmsListResourceTags, "list-resource-tags", "", false, "List Resource Tags")
	_kmsCmd.Flags().BoolVarP(&_kmsListRetirableGrants, "list-retirable-grants", "", false, "List Retirable Grants")
	_kmsCmd.Flags().BoolVarP(&_kmsPutKeyPolicy, "put-key-policy", "", false, "Put Key Policy")
	_kmsCmd.Flags().BoolVarP(&_kmsReEncrypt, "re-encrypt", "", false, "Re Encrypt")
	_kmsCmd.Flags().BoolVarP(&_kmsReplicateKey, "replicate-key", "", false, "Replicate Key")
	_kmsCmd.Flags().BoolVarP(&_kmsRetireGrant, "retire-grant", "", false, "Retire Grant")
	_kmsCmd.Flags().BoolVarP(&_kmsRevokeGrant, "revoke-grant", "", false, "Revoke Grant")
	_kmsCmd.Flags().BoolVarP(&_kmsRotateKeyOnDemand, "rotate-key-on-demand", "", false, "Rotate Key On Demand")
	_kmsCmd.Flags().BoolVarP(&_kmsScheduleKeyDeletion, "schedule-key-deletion", "", false, "Schedule Key Deletion")
	_kmsCmd.Flags().BoolVarP(&_kmsSign, "sign", "", false, "Sign")
	_kmsCmd.Flags().BoolVarP(&_kmsTagResource, "tag-resource", "", false, "Tag Resource")
	_kmsCmd.Flags().BoolVarP(&_kmsUntagResource, "untag-resource", "", false, "Untag Resource")
	_kmsCmd.Flags().BoolVarP(&_kmsUpdateAlias, "update-alias", "", false, "Update Alias")
	_kmsCmd.Flags().BoolVarP(&_kmsUpdateCustomKeyStore, "update-custom-key-store", "", false, "Update Custom Key Store")
	_kmsCmd.Flags().BoolVarP(&_kmsUpdateKeyDescription, "update-key-description", "", false, "Update Key Description")
	_kmsCmd.Flags().BoolVarP(&_kmsUpdatePrimaryRegion, "update-primary-region", "", false, "Update Primary Region")
	_kmsCmd.Flags().BoolVarP(&_kmsVerify, "verify", "", false, "Verify")
	_kmsCmd.Flags().BoolVarP(&_kmsVerifyMac, "verify-mac", "", false, "Verify Mac")

}
