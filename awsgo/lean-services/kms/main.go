package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/kms"
)

var fields_cancel_key_deletion = []leanruntime.Field{
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: true},
}

var fields_connect_custom_key_store = []leanruntime.Field{
	{Name: "CustomKeyStoreId", Flag: "custom-key-store-id", Type: "*string", Required: true},
}

var fields_create_alias = []leanruntime.Field{
	{Name: "AliasName", Flag: "alias-name", Type: "*string", Required: true},
	{Name: "TargetKeyId", Flag: "target-key-id", Type: "*string", Required: true},
}

var fields_create_custom_key_store = []leanruntime.Field{
	{Name: "CloudHsmClusterId", Flag: "cloud-hsm-cluster-id", Type: "*string", Required: false},
	{Name: "CustomKeyStoreName", Flag: "custom-key-store-name", Type: "*string", Required: true},
	{Name: "CustomKeyStoreType", Flag: "custom-key-store-type", Type: "types.CustomKeyStoreType", Required: false},
	{Name: "KeyStorePassword", Flag: "key-store-password", Type: "*string", Required: false},
	{Name: "TrustAnchorCertificate", Flag: "trust-anchor-certificate", Type: "*string", Required: false},
	{Name: "XksProxyAuthenticationCredential", Flag: "xks-proxy-authentication-credential", Type: "*types.XksProxyAuthenticationCredentialType", Required: false},
	{Name: "XksProxyConnectivity", Flag: "xks-proxy-connectivity", Type: "types.XksProxyConnectivityType", Required: false},
	{Name: "XksProxyUriEndpoint", Flag: "xks-proxy-uri-endpoint", Type: "*string", Required: false},
	{Name: "XksProxyUriPath", Flag: "xks-proxy-uri-path", Type: "*string", Required: false},
	{Name: "XksProxyVpcEndpointServiceName", Flag: "xks-proxy-vpc-endpoint-service-name", Type: "*string", Required: false},
	{Name: "XksProxyVpcEndpointServiceOwner", Flag: "xks-proxy-vpc-endpoint-service-owner", Type: "*string", Required: false},
}

var fields_create_grant = []leanruntime.Field{
	{Name: "Constraints", Flag: "constraints", Type: "*types.GrantConstraints", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "GrantTokens", Flag: "grant-tokens", Type: "[]string", Required: false},
	{Name: "GranteePrincipal", Flag: "grantee-principal", Type: "*string", Required: true},
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Operations", Flag: "operations", Type: "[]types.GrantOperation", Required: true},
	{Name: "RetiringPrincipal", Flag: "retiring-principal", Type: "*string", Required: false},
}

var fields_create_key = []leanruntime.Field{
	{Name: "BypassPolicyLockoutSafetyCheck", Flag: "bypass-policy-lockout-safety-check", Type: "bool", Required: false},
	{Name: "CustomKeyStoreId", Flag: "custom-key-store-id", Type: "*string", Required: false},
	{Name: "CustomerMasterKeySpec", Flag: "customer-master-key-spec", Type: "types.CustomerMasterKeySpec", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "KeySpec", Flag: "key-spec", Type: "types.KeySpec", Required: false},
	{Name: "KeyUsage", Flag: "key-usage", Type: "types.KeyUsageType", Required: false},
	{Name: "MultiRegion", Flag: "multi-region", Type: "*bool", Required: false},
	{Name: "Origin", Flag: "origin", Type: "types.OriginType", Required: false},
	{Name: "Policy", Flag: "policy", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "XksKeyId", Flag: "xks-key-id", Type: "*string", Required: false},
}

var fields_decrypt = []leanruntime.Field{
	{Name: "CiphertextBlob", Flag: "ciphertext-blob", Type: "[]byte", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "DryRunModifiers", Flag: "dry-run-modifiers", Type: "[]types.DryRunModifierType", Required: false},
	{Name: "EncryptionAlgorithm", Flag: "encryption-algorithm", Type: "types.EncryptionAlgorithmSpec", Required: false},
	{Name: "EncryptionContext", Flag: "encryption-context", Type: "map[string]string", Required: false},
	{Name: "GrantTokens", Flag: "grant-tokens", Type: "[]string", Required: false},
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: false},
	{Name: "Recipient", Flag: "recipient", Type: "*types.RecipientInfo", Required: false},
}

var fields_delete_alias = []leanruntime.Field{
	{Name: "AliasName", Flag: "alias-name", Type: "*string", Required: true},
}

var fields_delete_custom_key_store = []leanruntime.Field{
	{Name: "CustomKeyStoreId", Flag: "custom-key-store-id", Type: "*string", Required: true},
}

var fields_delete_imported_key_material = []leanruntime.Field{
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: true},
	{Name: "KeyMaterialId", Flag: "key-material-id", Type: "*string", Required: false},
}

var fields_derive_shared_secret = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "GrantTokens", Flag: "grant-tokens", Type: "[]string", Required: false},
	{Name: "KeyAgreementAlgorithm", Flag: "key-agreement-algorithm", Type: "types.KeyAgreementAlgorithmSpec", Required: true},
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: true},
	{Name: "PublicKey", Flag: "public-key", Type: "[]byte", Required: true},
	{Name: "Recipient", Flag: "recipient", Type: "*types.RecipientInfo", Required: false},
}

var fields_describe_custom_key_stores = []leanruntime.Field{
	{Name: "CustomKeyStoreId", Flag: "custom-key-store-id", Type: "*string", Required: false},
	{Name: "CustomKeyStoreName", Flag: "custom-key-store-name", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
}

var fields_describe_key = []leanruntime.Field{
	{Name: "GrantTokens", Flag: "grant-tokens", Type: "[]string", Required: false},
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: true},
}

var fields_disable_key = []leanruntime.Field{
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: true},
}

var fields_disable_key_rotation = []leanruntime.Field{
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: true},
}

var fields_disconnect_custom_key_store = []leanruntime.Field{
	{Name: "CustomKeyStoreId", Flag: "custom-key-store-id", Type: "*string", Required: true},
}

var fields_enable_key = []leanruntime.Field{
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: true},
}

var fields_enable_key_rotation = []leanruntime.Field{
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: true},
	{Name: "RotationPeriodInDays", Flag: "rotation-period-in-days", Type: "*int32", Required: false},
}

var fields_encrypt = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "EncryptionAlgorithm", Flag: "encryption-algorithm", Type: "types.EncryptionAlgorithmSpec", Required: false},
	{Name: "EncryptionContext", Flag: "encryption-context", Type: "map[string]string", Required: false},
	{Name: "GrantTokens", Flag: "grant-tokens", Type: "[]string", Required: false},
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: true},
	{Name: "Plaintext", Flag: "plaintext", Type: "[]byte", Required: true},
}

var fields_generate_data_key = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "EncryptionContext", Flag: "encryption-context", Type: "map[string]string", Required: false},
	{Name: "GrantTokens", Flag: "grant-tokens", Type: "[]string", Required: false},
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: true},
	{Name: "KeySpec", Flag: "key-spec", Type: "types.DataKeySpec", Required: false},
	{Name: "NumberOfBytes", Flag: "number-of-bytes", Type: "*int32", Required: false},
	{Name: "Recipient", Flag: "recipient", Type: "*types.RecipientInfo", Required: false},
}

var fields_generate_data_key_pair = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "EncryptionContext", Flag: "encryption-context", Type: "map[string]string", Required: false},
	{Name: "GrantTokens", Flag: "grant-tokens", Type: "[]string", Required: false},
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: true},
	{Name: "KeyPairSpec", Flag: "key-pair-spec", Type: "types.DataKeyPairSpec", Required: true},
	{Name: "Recipient", Flag: "recipient", Type: "*types.RecipientInfo", Required: false},
}

var fields_generate_data_key_pair_without_plaintext = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "EncryptionContext", Flag: "encryption-context", Type: "map[string]string", Required: false},
	{Name: "GrantTokens", Flag: "grant-tokens", Type: "[]string", Required: false},
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: true},
	{Name: "KeyPairSpec", Flag: "key-pair-spec", Type: "types.DataKeyPairSpec", Required: true},
}

var fields_generate_data_key_without_plaintext = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "EncryptionContext", Flag: "encryption-context", Type: "map[string]string", Required: false},
	{Name: "GrantTokens", Flag: "grant-tokens", Type: "[]string", Required: false},
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: true},
	{Name: "KeySpec", Flag: "key-spec", Type: "types.DataKeySpec", Required: false},
	{Name: "NumberOfBytes", Flag: "number-of-bytes", Type: "*int32", Required: false},
}

var fields_generate_mac = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "GrantTokens", Flag: "grant-tokens", Type: "[]string", Required: false},
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: true},
	{Name: "MacAlgorithm", Flag: "mac-algorithm", Type: "types.MacAlgorithmSpec", Required: true},
	{Name: "Message", Flag: "message", Type: "[]byte", Required: true},
}

var fields_generate_random = []leanruntime.Field{
	{Name: "CustomKeyStoreId", Flag: "custom-key-store-id", Type: "*string", Required: false},
	{Name: "NumberOfBytes", Flag: "number-of-bytes", Type: "*int32", Required: false},
	{Name: "Recipient", Flag: "recipient", Type: "*types.RecipientInfo", Required: false},
}

var fields_get_key_policy = []leanruntime.Field{
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: true},
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: false},
}

var fields_get_key_rotation_status = []leanruntime.Field{
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: true},
}

var fields_get_parameters_for_import = []leanruntime.Field{
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: true},
	{Name: "WrappingAlgorithm", Flag: "wrapping-algorithm", Type: "types.AlgorithmSpec", Required: true},
	{Name: "WrappingKeySpec", Flag: "wrapping-key-spec", Type: "types.WrappingKeySpec", Required: true},
}

var fields_get_public_key = []leanruntime.Field{
	{Name: "GrantTokens", Flag: "grant-tokens", Type: "[]string", Required: false},
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: true},
}

var fields_import_key_material = []leanruntime.Field{
	{Name: "EncryptedKeyMaterial", Flag: "encrypted-key-material", Type: "[]byte", Required: true},
	{Name: "ExpirationModel", Flag: "expiration-model", Type: "types.ExpirationModelType", Required: false},
	{Name: "ImportToken", Flag: "import-token", Type: "[]byte", Required: true},
	{Name: "ImportType", Flag: "import-type", Type: "types.ImportType", Required: false},
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: true},
	{Name: "KeyMaterialDescription", Flag: "key-material-description", Type: "*string", Required: false},
	{Name: "KeyMaterialId", Flag: "key-material-id", Type: "*string", Required: false},
	{Name: "ValidTo", Flag: "valid-to", Type: "*time.Time", Required: false},
}

var fields_list_aliases = []leanruntime.Field{
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
}

var fields_list_grants = []leanruntime.Field{
	{Name: "GrantId", Flag: "grant-id", Type: "*string", Required: false},
	{Name: "GranteePrincipal", Flag: "grantee-principal", Type: "*string", Required: false},
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: true},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
}

var fields_list_key_policies = []leanruntime.Field{
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: true},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
}

var fields_list_key_rotations = []leanruntime.Field{
	{Name: "IncludeKeyMaterial", Flag: "include-key-material", Type: "types.IncludeKeyMaterial", Required: false},
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: true},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
}

var fields_list_keys = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
}

var fields_list_resource_tags = []leanruntime.Field{
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: true},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
}

var fields_list_retirable_grants = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "RetiringPrincipal", Flag: "retiring-principal", Type: "*string", Required: true},
}

var fields_put_key_policy = []leanruntime.Field{
	{Name: "BypassPolicyLockoutSafetyCheck", Flag: "bypass-policy-lockout-safety-check", Type: "bool", Required: false},
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: true},
	{Name: "Policy", Flag: "policy", Type: "*string", Required: true},
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: false},
}

var fields_re_encrypt = []leanruntime.Field{
	{Name: "CiphertextBlob", Flag: "ciphertext-blob", Type: "[]byte", Required: false},
	{Name: "DestinationEncryptionAlgorithm", Flag: "destination-encryption-algorithm", Type: "types.EncryptionAlgorithmSpec", Required: false},
	{Name: "DestinationEncryptionContext", Flag: "destination-encryption-context", Type: "map[string]string", Required: false},
	{Name: "DestinationKeyId", Flag: "destination-key-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "DryRunModifiers", Flag: "dry-run-modifiers", Type: "[]types.DryRunModifierType", Required: false},
	{Name: "GrantTokens", Flag: "grant-tokens", Type: "[]string", Required: false},
	{Name: "SourceEncryptionAlgorithm", Flag: "source-encryption-algorithm", Type: "types.EncryptionAlgorithmSpec", Required: false},
	{Name: "SourceEncryptionContext", Flag: "source-encryption-context", Type: "map[string]string", Required: false},
	{Name: "SourceKeyId", Flag: "source-key-id", Type: "*string", Required: false},
}

var fields_replicate_key = []leanruntime.Field{
	{Name: "BypassPolicyLockoutSafetyCheck", Flag: "bypass-policy-lockout-safety-check", Type: "bool", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: true},
	{Name: "Policy", Flag: "policy", Type: "*string", Required: false},
	{Name: "ReplicaRegion", Flag: "replica-region", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_retire_grant = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "GrantId", Flag: "grant-id", Type: "*string", Required: false},
	{Name: "GrantToken", Flag: "grant-token", Type: "*string", Required: false},
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: false},
}

var fields_revoke_grant = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "GrantId", Flag: "grant-id", Type: "*string", Required: true},
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: true},
}

var fields_rotate_key_on_demand = []leanruntime.Field{
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: true},
}

var fields_schedule_key_deletion = []leanruntime.Field{
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: true},
	{Name: "PendingWindowInDays", Flag: "pending-window-in-days", Type: "*int32", Required: false},
}

var fields_sign = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "GrantTokens", Flag: "grant-tokens", Type: "[]string", Required: false},
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: true},
	{Name: "Message", Flag: "message", Type: "[]byte", Required: true},
	{Name: "MessageType", Flag: "message-type", Type: "types.MessageType", Required: false},
	{Name: "SigningAlgorithm", Flag: "signing-algorithm", Type: "types.SigningAlgorithmSpec", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_alias = []leanruntime.Field{
	{Name: "AliasName", Flag: "alias-name", Type: "*string", Required: true},
	{Name: "TargetKeyId", Flag: "target-key-id", Type: "*string", Required: true},
}

var fields_update_custom_key_store = []leanruntime.Field{
	{Name: "CloudHsmClusterId", Flag: "cloud-hsm-cluster-id", Type: "*string", Required: false},
	{Name: "CustomKeyStoreId", Flag: "custom-key-store-id", Type: "*string", Required: true},
	{Name: "KeyStorePassword", Flag: "key-store-password", Type: "*string", Required: false},
	{Name: "NewCustomKeyStoreName", Flag: "new-custom-key-store-name", Type: "*string", Required: false},
	{Name: "XksProxyAuthenticationCredential", Flag: "xks-proxy-authentication-credential", Type: "*types.XksProxyAuthenticationCredentialType", Required: false},
	{Name: "XksProxyConnectivity", Flag: "xks-proxy-connectivity", Type: "types.XksProxyConnectivityType", Required: false},
	{Name: "XksProxyUriEndpoint", Flag: "xks-proxy-uri-endpoint", Type: "*string", Required: false},
	{Name: "XksProxyUriPath", Flag: "xks-proxy-uri-path", Type: "*string", Required: false},
	{Name: "XksProxyVpcEndpointServiceName", Flag: "xks-proxy-vpc-endpoint-service-name", Type: "*string", Required: false},
	{Name: "XksProxyVpcEndpointServiceOwner", Flag: "xks-proxy-vpc-endpoint-service-owner", Type: "*string", Required: false},
}

var fields_update_key_description = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: true},
}

var fields_update_primary_region = []leanruntime.Field{
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: true},
	{Name: "PrimaryRegion", Flag: "primary-region", Type: "*string", Required: true},
}

var fields_verify = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "GrantTokens", Flag: "grant-tokens", Type: "[]string", Required: false},
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: true},
	{Name: "Message", Flag: "message", Type: "[]byte", Required: true},
	{Name: "MessageType", Flag: "message-type", Type: "types.MessageType", Required: false},
	{Name: "Signature", Flag: "signature", Type: "[]byte", Required: true},
	{Name: "SigningAlgorithm", Flag: "signing-algorithm", Type: "types.SigningAlgorithmSpec", Required: true},
}

var fields_verify_mac = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "GrantTokens", Flag: "grant-tokens", Type: "[]string", Required: false},
	{Name: "KeyId", Flag: "key-id", Type: "*string", Required: true},
	{Name: "Mac", Flag: "mac", Type: "[]byte", Required: true},
	{Name: "MacAlgorithm", Flag: "mac-algorithm", Type: "types.MacAlgorithmSpec", Required: true},
	{Name: "Message", Flag: "message", Type: "[]byte", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"cancel-key-deletion": {
			Name:   "cancel-key-deletion",
			Fields: fields_cancel_key_deletion,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelKeyDeletionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_key_deletion, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelKeyDeletion(ctx, input)
			},
		},
		"connect-custom-key-store": {
			Name:   "connect-custom-key-store",
			Fields: fields_connect_custom_key_store,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ConnectCustomKeyStoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_connect_custom_key_store, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ConnectCustomKeyStore(ctx, input)
			},
		},
		"create-alias": {
			Name:   "create-alias",
			Fields: fields_create_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAlias(ctx, input)
			},
		},
		"create-custom-key-store": {
			Name:   "create-custom-key-store",
			Fields: fields_create_custom_key_store,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCustomKeyStoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_custom_key_store, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCustomKeyStore(ctx, input)
			},
		},
		"create-grant": {
			Name:   "create-grant",
			Fields: fields_create_grant,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateGrantInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_grant, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateGrant(ctx, input)
			},
		},
		"create-key": {
			Name:   "create-key",
			Fields: fields_create_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateKey(ctx, input)
			},
		},
		"decrypt": {
			Name:   "decrypt",
			Fields: fields_decrypt,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DecryptInput{}
				if _, err := leanruntime.ApplyInput(input, fields_decrypt, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.Decrypt(ctx, input)
			},
		},
		"delete-alias": {
			Name:   "delete-alias",
			Fields: fields_delete_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAlias(ctx, input)
			},
		},
		"delete-custom-key-store": {
			Name:   "delete-custom-key-store",
			Fields: fields_delete_custom_key_store,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCustomKeyStoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_custom_key_store, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCustomKeyStore(ctx, input)
			},
		},
		"delete-imported-key-material": {
			Name:   "delete-imported-key-material",
			Fields: fields_delete_imported_key_material,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteImportedKeyMaterialInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_imported_key_material, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteImportedKeyMaterial(ctx, input)
			},
		},
		"derive-shared-secret": {
			Name:   "derive-shared-secret",
			Fields: fields_derive_shared_secret,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeriveSharedSecretInput{}
				if _, err := leanruntime.ApplyInput(input, fields_derive_shared_secret, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeriveSharedSecret(ctx, input)
			},
		},
		"describe-custom-key-stores": {
			Name:   "describe-custom-key-stores",
			Fields: fields_describe_custom_key_stores,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCustomKeyStoresInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_custom_key_stores, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeCustomKeyStores(ctx, input)
				}
				var results []*svc.DescribeCustomKeyStoresOutput
				p := svc.NewDescribeCustomKeyStoresPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-key": {
			Name:   "describe-key",
			Fields: fields_describe_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeKey(ctx, input)
			},
		},
		"disable-key": {
			Name:   "disable-key",
			Fields: fields_disable_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableKey(ctx, input)
			},
		},
		"disable-key-rotation": {
			Name:   "disable-key-rotation",
			Fields: fields_disable_key_rotation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableKeyRotationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_key_rotation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableKeyRotation(ctx, input)
			},
		},
		"disconnect-custom-key-store": {
			Name:   "disconnect-custom-key-store",
			Fields: fields_disconnect_custom_key_store,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisconnectCustomKeyStoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disconnect_custom_key_store, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisconnectCustomKeyStore(ctx, input)
			},
		},
		"enable-key": {
			Name:   "enable-key",
			Fields: fields_enable_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableKey(ctx, input)
			},
		},
		"enable-key-rotation": {
			Name:   "enable-key-rotation",
			Fields: fields_enable_key_rotation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableKeyRotationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_key_rotation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableKeyRotation(ctx, input)
			},
		},
		"encrypt": {
			Name:   "encrypt",
			Fields: fields_encrypt,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EncryptInput{}
				if _, err := leanruntime.ApplyInput(input, fields_encrypt, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.Encrypt(ctx, input)
			},
		},
		"generate-data-key": {
			Name:   "generate-data-key",
			Fields: fields_generate_data_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GenerateDataKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_generate_data_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GenerateDataKey(ctx, input)
			},
		},
		"generate-data-key-pair": {
			Name:   "generate-data-key-pair",
			Fields: fields_generate_data_key_pair,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GenerateDataKeyPairInput{}
				if _, err := leanruntime.ApplyInput(input, fields_generate_data_key_pair, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GenerateDataKeyPair(ctx, input)
			},
		},
		"generate-data-key-pair-without-plaintext": {
			Name:   "generate-data-key-pair-without-plaintext",
			Fields: fields_generate_data_key_pair_without_plaintext,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GenerateDataKeyPairWithoutPlaintextInput{}
				if _, err := leanruntime.ApplyInput(input, fields_generate_data_key_pair_without_plaintext, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GenerateDataKeyPairWithoutPlaintext(ctx, input)
			},
		},
		"generate-data-key-without-plaintext": {
			Name:   "generate-data-key-without-plaintext",
			Fields: fields_generate_data_key_without_plaintext,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GenerateDataKeyWithoutPlaintextInput{}
				if _, err := leanruntime.ApplyInput(input, fields_generate_data_key_without_plaintext, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GenerateDataKeyWithoutPlaintext(ctx, input)
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
		"generate-random": {
			Name:   "generate-random",
			Fields: fields_generate_random,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GenerateRandomInput{}
				if _, err := leanruntime.ApplyInput(input, fields_generate_random, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GenerateRandom(ctx, input)
			},
		},
		"get-key-policy": {
			Name:   "get-key-policy",
			Fields: fields_get_key_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetKeyPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_key_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetKeyPolicy(ctx, input)
			},
		},
		"get-key-rotation-status": {
			Name:   "get-key-rotation-status",
			Fields: fields_get_key_rotation_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetKeyRotationStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_key_rotation_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetKeyRotationStatus(ctx, input)
			},
		},
		"get-parameters-for-import": {
			Name:   "get-parameters-for-import",
			Fields: fields_get_parameters_for_import,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetParametersForImportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_parameters_for_import, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetParametersForImport(ctx, input)
			},
		},
		"get-public-key": {
			Name:   "get-public-key",
			Fields: fields_get_public_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPublicKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_public_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPublicKey(ctx, input)
			},
		},
		"import-key-material": {
			Name:   "import-key-material",
			Fields: fields_import_key_material,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportKeyMaterialInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_key_material, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportKeyMaterial(ctx, input)
			},
		},
		"list-aliases": {
			Name:   "list-aliases",
			Fields: fields_list_aliases,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAliasesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_aliases, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAliases(ctx, input)
				}
				var results []*svc.ListAliasesOutput
				p := svc.NewListAliasesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-grants": {
			Name:   "list-grants",
			Fields: fields_list_grants,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGrantsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_grants, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListGrants(ctx, input)
				}
				var results []*svc.ListGrantsOutput
				p := svc.NewListGrantsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-key-policies": {
			Name:   "list-key-policies",
			Fields: fields_list_key_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListKeyPoliciesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_key_policies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListKeyPolicies(ctx, input)
				}
				var results []*svc.ListKeyPoliciesOutput
				p := svc.NewListKeyPoliciesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-key-rotations": {
			Name:   "list-key-rotations",
			Fields: fields_list_key_rotations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListKeyRotationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_key_rotations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListKeyRotations(ctx, input)
				}
				var results []*svc.ListKeyRotationsOutput
				p := svc.NewListKeyRotationsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-keys": {
			Name:   "list-keys",
			Fields: fields_list_keys,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListKeysInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_keys, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListKeys(ctx, input)
				}
				var results []*svc.ListKeysOutput
				p := svc.NewListKeysPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-resource-tags": {
			Name:   "list-resource-tags",
			Fields: fields_list_resource_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResourceTagsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_resource_tags, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResourceTags(ctx, input)
				}
				var results []*svc.ListResourceTagsOutput
				p := svc.NewListResourceTagsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-retirable-grants": {
			Name:   "list-retirable-grants",
			Fields: fields_list_retirable_grants,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRetirableGrantsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_retirable_grants, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRetirableGrants(ctx, input)
				}
				var results []*svc.ListRetirableGrantsOutput
				p := svc.NewListRetirableGrantsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"put-key-policy": {
			Name:   "put-key-policy",
			Fields: fields_put_key_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutKeyPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_key_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutKeyPolicy(ctx, input)
			},
		},
		"re-encrypt": {
			Name:   "re-encrypt",
			Fields: fields_re_encrypt,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ReEncryptInput{}
				if _, err := leanruntime.ApplyInput(input, fields_re_encrypt, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ReEncrypt(ctx, input)
			},
		},
		"replicate-key": {
			Name:   "replicate-key",
			Fields: fields_replicate_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ReplicateKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_replicate_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ReplicateKey(ctx, input)
			},
		},
		"retire-grant": {
			Name:   "retire-grant",
			Fields: fields_retire_grant,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RetireGrantInput{}
				if _, err := leanruntime.ApplyInput(input, fields_retire_grant, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RetireGrant(ctx, input)
			},
		},
		"revoke-grant": {
			Name:   "revoke-grant",
			Fields: fields_revoke_grant,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RevokeGrantInput{}
				if _, err := leanruntime.ApplyInput(input, fields_revoke_grant, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RevokeGrant(ctx, input)
			},
		},
		"rotate-key-on-demand": {
			Name:   "rotate-key-on-demand",
			Fields: fields_rotate_key_on_demand,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RotateKeyOnDemandInput{}
				if _, err := leanruntime.ApplyInput(input, fields_rotate_key_on_demand, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RotateKeyOnDemand(ctx, input)
			},
		},
		"schedule-key-deletion": {
			Name:   "schedule-key-deletion",
			Fields: fields_schedule_key_deletion,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ScheduleKeyDeletionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_schedule_key_deletion, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ScheduleKeyDeletion(ctx, input)
			},
		},
		"sign": {
			Name:   "sign",
			Fields: fields_sign,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SignInput{}
				if _, err := leanruntime.ApplyInput(input, fields_sign, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.Sign(ctx, input)
			},
		},
		"tag-resource": {
			Name:   "tag-resource",
			Fields: fields_tag_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TagResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_tag_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TagResource(ctx, input)
			},
		},
		"untag-resource": {
			Name:   "untag-resource",
			Fields: fields_untag_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UntagResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_untag_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UntagResource(ctx, input)
			},
		},
		"update-alias": {
			Name:   "update-alias",
			Fields: fields_update_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAlias(ctx, input)
			},
		},
		"update-custom-key-store": {
			Name:   "update-custom-key-store",
			Fields: fields_update_custom_key_store,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCustomKeyStoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_custom_key_store, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCustomKeyStore(ctx, input)
			},
		},
		"update-key-description": {
			Name:   "update-key-description",
			Fields: fields_update_key_description,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateKeyDescriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_key_description, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateKeyDescription(ctx, input)
			},
		},
		"update-primary-region": {
			Name:   "update-primary-region",
			Fields: fields_update_primary_region,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePrimaryRegionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_primary_region, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePrimaryRegion(ctx, input)
			},
		},
		"verify": {
			Name:   "verify",
			Fields: fields_verify,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.VerifyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_verify, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.Verify(ctx, input)
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
	}
	if err := leanruntime.Execute("kms", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
