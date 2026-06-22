package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/workspacesweb"
)

var fields_associate_browser_settings = []leanruntime.Field{
	{Name: "BrowserSettingsArn", Flag: "browser-settings-arn", Type: "*string", Required: true},
	{Name: "PortalArn", Flag: "portal-arn", Type: "*string", Required: true},
}

var fields_associate_data_protection_settings = []leanruntime.Field{
	{Name: "DataProtectionSettingsArn", Flag: "data-protection-settings-arn", Type: "*string", Required: true},
	{Name: "PortalArn", Flag: "portal-arn", Type: "*string", Required: true},
}

var fields_associate_ip_access_settings = []leanruntime.Field{
	{Name: "IpAccessSettingsArn", Flag: "ip-access-settings-arn", Type: "*string", Required: true},
	{Name: "PortalArn", Flag: "portal-arn", Type: "*string", Required: true},
}

var fields_associate_network_settings = []leanruntime.Field{
	{Name: "NetworkSettingsArn", Flag: "network-settings-arn", Type: "*string", Required: true},
	{Name: "PortalArn", Flag: "portal-arn", Type: "*string", Required: true},
}

var fields_associate_session_logger = []leanruntime.Field{
	{Name: "PortalArn", Flag: "portal-arn", Type: "*string", Required: true},
	{Name: "SessionLoggerArn", Flag: "session-logger-arn", Type: "*string", Required: true},
}

var fields_associate_trust_store = []leanruntime.Field{
	{Name: "PortalArn", Flag: "portal-arn", Type: "*string", Required: true},
	{Name: "TrustStoreArn", Flag: "trust-store-arn", Type: "*string", Required: true},
}

var fields_associate_user_access_logging_settings = []leanruntime.Field{
	{Name: "PortalArn", Flag: "portal-arn", Type: "*string", Required: true},
	{Name: "UserAccessLoggingSettingsArn", Flag: "user-access-logging-settings-arn", Type: "*string", Required: true},
}

var fields_associate_user_settings = []leanruntime.Field{
	{Name: "PortalArn", Flag: "portal-arn", Type: "*string", Required: true},
	{Name: "UserSettingsArn", Flag: "user-settings-arn", Type: "*string", Required: true},
}

var fields_create_browser_settings = []leanruntime.Field{
	{Name: "AdditionalEncryptionContext", Flag: "additional-encryption-context", Type: "map[string]string", Required: false},
	{Name: "BrowserPolicy", Flag: "browser-policy", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CustomerManagedKey", Flag: "customer-managed-key", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "WebContentFilteringPolicy", Flag: "web-content-filtering-policy", Type: "*types.WebContentFilteringPolicy", Required: false},
}

var fields_create_data_protection_settings = []leanruntime.Field{
	{Name: "AdditionalEncryptionContext", Flag: "additional-encryption-context", Type: "map[string]string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CustomerManagedKey", Flag: "customer-managed-key", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "InlineRedactionConfiguration", Flag: "inline-redaction-configuration", Type: "*types.InlineRedactionConfiguration", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_identity_provider = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "IdentityProviderDetails", Flag: "identity-provider-details", Type: "map[string]string", Required: true},
	{Name: "IdentityProviderName", Flag: "identity-provider-name", Type: "*string", Required: true},
	{Name: "IdentityProviderType", Flag: "identity-provider-type", Type: "types.IdentityProviderType", Required: true},
	{Name: "PortalArn", Flag: "portal-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_ip_access_settings = []leanruntime.Field{
	{Name: "AdditionalEncryptionContext", Flag: "additional-encryption-context", Type: "map[string]string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CustomerManagedKey", Flag: "customer-managed-key", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "IpRules", Flag: "ip-rules", Type: "[]types.IpRule", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_network_settings = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "SecurityGroupIds", Flag: "security-group-ids", Type: "[]string", Required: true},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: true},
}

var fields_create_portal = []leanruntime.Field{
	{Name: "AdditionalEncryptionContext", Flag: "additional-encryption-context", Type: "map[string]string", Required: false},
	{Name: "AuthenticationType", Flag: "authentication-type", Type: "types.AuthenticationType", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CustomerManagedKey", Flag: "customer-managed-key", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "InstanceType", Flag: "instance-type", Type: "types.InstanceType", Required: false},
	{Name: "MaxConcurrentSessions", Flag: "max-concurrent-sessions", Type: "*int32", Required: false},
	{Name: "PortalCustomDomain", Flag: "portal-custom-domain", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_session_logger = []leanruntime.Field{
	{Name: "AdditionalEncryptionContext", Flag: "additional-encryption-context", Type: "map[string]string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CustomerManagedKey", Flag: "customer-managed-key", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "EventFilter", Flag: "event-filter", Type: "types.EventFilter", Required: true},
	{Name: "LogConfiguration", Flag: "log-configuration", Type: "*types.LogConfiguration", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_trust_store = []leanruntime.Field{
	{Name: "CertificateList", Flag: "certificate-list", Type: "[][]byte", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_user_access_logging_settings = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "KinesisStreamArn", Flag: "kinesis-stream-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_user_settings = []leanruntime.Field{
	{Name: "AdditionalEncryptionContext", Flag: "additional-encryption-context", Type: "map[string]string", Required: false},
	{Name: "BrandingConfigurationInput", Flag: "branding-configuration-input", Type: "*types.BrandingConfigurationCreateInput", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CookieSynchronizationConfiguration", Flag: "cookie-synchronization-configuration", Type: "*types.CookieSynchronizationConfiguration", Required: false},
	{Name: "CopyAllowed", Flag: "copy-allowed", Type: "types.EnabledType", Required: true},
	{Name: "CustomerManagedKey", Flag: "customer-managed-key", Type: "*string", Required: false},
	{Name: "DeepLinkAllowed", Flag: "deep-link-allowed", Type: "types.EnabledType", Required: false},
	{Name: "DisconnectTimeoutInMinutes", Flag: "disconnect-timeout-in-minutes", Type: "*int32", Required: false},
	{Name: "DownloadAllowed", Flag: "download-allowed", Type: "types.EnabledType", Required: true},
	{Name: "IdleDisconnectTimeoutInMinutes", Flag: "idle-disconnect-timeout-in-minutes", Type: "*int32", Required: false},
	{Name: "PasteAllowed", Flag: "paste-allowed", Type: "types.EnabledType", Required: true},
	{Name: "PrintAllowed", Flag: "print-allowed", Type: "types.EnabledType", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "ToolbarConfiguration", Flag: "toolbar-configuration", Type: "*types.ToolbarConfiguration", Required: false},
	{Name: "UploadAllowed", Flag: "upload-allowed", Type: "types.EnabledType", Required: true},
	{Name: "WebAuthnAllowed", Flag: "web-authn-allowed", Type: "types.EnabledType", Required: false},
}

var fields_delete_browser_settings = []leanruntime.Field{
	{Name: "BrowserSettingsArn", Flag: "browser-settings-arn", Type: "*string", Required: true},
}

var fields_delete_data_protection_settings = []leanruntime.Field{
	{Name: "DataProtectionSettingsArn", Flag: "data-protection-settings-arn", Type: "*string", Required: true},
}

var fields_delete_identity_provider = []leanruntime.Field{
	{Name: "IdentityProviderArn", Flag: "identity-provider-arn", Type: "*string", Required: true},
}

var fields_delete_ip_access_settings = []leanruntime.Field{
	{Name: "IpAccessSettingsArn", Flag: "ip-access-settings-arn", Type: "*string", Required: true},
}

var fields_delete_network_settings = []leanruntime.Field{
	{Name: "NetworkSettingsArn", Flag: "network-settings-arn", Type: "*string", Required: true},
}

var fields_delete_portal = []leanruntime.Field{
	{Name: "PortalArn", Flag: "portal-arn", Type: "*string", Required: true},
}

var fields_delete_session_logger = []leanruntime.Field{
	{Name: "SessionLoggerArn", Flag: "session-logger-arn", Type: "*string", Required: true},
}

var fields_delete_trust_store = []leanruntime.Field{
	{Name: "TrustStoreArn", Flag: "trust-store-arn", Type: "*string", Required: true},
}

var fields_delete_user_access_logging_settings = []leanruntime.Field{
	{Name: "UserAccessLoggingSettingsArn", Flag: "user-access-logging-settings-arn", Type: "*string", Required: true},
}

var fields_delete_user_settings = []leanruntime.Field{
	{Name: "UserSettingsArn", Flag: "user-settings-arn", Type: "*string", Required: true},
}

var fields_disassociate_browser_settings = []leanruntime.Field{
	{Name: "PortalArn", Flag: "portal-arn", Type: "*string", Required: true},
}

var fields_disassociate_data_protection_settings = []leanruntime.Field{
	{Name: "PortalArn", Flag: "portal-arn", Type: "*string", Required: true},
}

var fields_disassociate_ip_access_settings = []leanruntime.Field{
	{Name: "PortalArn", Flag: "portal-arn", Type: "*string", Required: true},
}

var fields_disassociate_network_settings = []leanruntime.Field{
	{Name: "PortalArn", Flag: "portal-arn", Type: "*string", Required: true},
}

var fields_disassociate_session_logger = []leanruntime.Field{
	{Name: "PortalArn", Flag: "portal-arn", Type: "*string", Required: true},
}

var fields_disassociate_trust_store = []leanruntime.Field{
	{Name: "PortalArn", Flag: "portal-arn", Type: "*string", Required: true},
}

var fields_disassociate_user_access_logging_settings = []leanruntime.Field{
	{Name: "PortalArn", Flag: "portal-arn", Type: "*string", Required: true},
}

var fields_disassociate_user_settings = []leanruntime.Field{
	{Name: "PortalArn", Flag: "portal-arn", Type: "*string", Required: true},
}

var fields_expire_session = []leanruntime.Field{
	{Name: "PortalId", Flag: "portal-id", Type: "*string", Required: true},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
}

var fields_get_browser_settings = []leanruntime.Field{
	{Name: "BrowserSettingsArn", Flag: "browser-settings-arn", Type: "*string", Required: true},
}

var fields_get_data_protection_settings = []leanruntime.Field{
	{Name: "DataProtectionSettingsArn", Flag: "data-protection-settings-arn", Type: "*string", Required: true},
}

var fields_get_identity_provider = []leanruntime.Field{
	{Name: "IdentityProviderArn", Flag: "identity-provider-arn", Type: "*string", Required: true},
}

var fields_get_ip_access_settings = []leanruntime.Field{
	{Name: "IpAccessSettingsArn", Flag: "ip-access-settings-arn", Type: "*string", Required: true},
}

var fields_get_network_settings = []leanruntime.Field{
	{Name: "NetworkSettingsArn", Flag: "network-settings-arn", Type: "*string", Required: true},
}

var fields_get_portal = []leanruntime.Field{
	{Name: "PortalArn", Flag: "portal-arn", Type: "*string", Required: true},
}

var fields_get_portal_service_provider_metadata = []leanruntime.Field{
	{Name: "PortalArn", Flag: "portal-arn", Type: "*string", Required: true},
}

var fields_get_session = []leanruntime.Field{
	{Name: "PortalId", Flag: "portal-id", Type: "*string", Required: true},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
}

var fields_get_session_logger = []leanruntime.Field{
	{Name: "SessionLoggerArn", Flag: "session-logger-arn", Type: "*string", Required: true},
}

var fields_get_trust_store = []leanruntime.Field{
	{Name: "TrustStoreArn", Flag: "trust-store-arn", Type: "*string", Required: true},
}

var fields_get_trust_store_certificate = []leanruntime.Field{
	{Name: "Thumbprint", Flag: "thumbprint", Type: "*string", Required: true},
	{Name: "TrustStoreArn", Flag: "trust-store-arn", Type: "*string", Required: true},
}

var fields_get_user_access_logging_settings = []leanruntime.Field{
	{Name: "UserAccessLoggingSettingsArn", Flag: "user-access-logging-settings-arn", Type: "*string", Required: true},
}

var fields_get_user_settings = []leanruntime.Field{
	{Name: "UserSettingsArn", Flag: "user-settings-arn", Type: "*string", Required: true},
}

var fields_list_browser_settings = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_data_protection_settings = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_identity_providers = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PortalArn", Flag: "portal-arn", Type: "*string", Required: true},
}

var fields_list_ip_access_settings = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_network_settings = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_portals = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_session_loggers = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_sessions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PortalId", Flag: "portal-id", Type: "*string", Required: true},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.SessionSortBy", Required: false},
	{Name: "Status", Flag: "status", Type: "types.SessionStatus", Required: false},
	{Name: "Username", Flag: "username", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_trust_store_certificates = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TrustStoreArn", Flag: "trust-store-arn", Type: "*string", Required: true},
}

var fields_list_trust_stores = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_user_access_logging_settings = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_user_settings = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_browser_settings = []leanruntime.Field{
	{Name: "BrowserPolicy", Flag: "browser-policy", Type: "*string", Required: false},
	{Name: "BrowserSettingsArn", Flag: "browser-settings-arn", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "WebContentFilteringPolicy", Flag: "web-content-filtering-policy", Type: "*types.WebContentFilteringPolicy", Required: false},
}

var fields_update_data_protection_settings = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DataProtectionSettingsArn", Flag: "data-protection-settings-arn", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "InlineRedactionConfiguration", Flag: "inline-redaction-configuration", Type: "*types.InlineRedactionConfiguration", Required: false},
}

var fields_update_identity_provider = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "IdentityProviderArn", Flag: "identity-provider-arn", Type: "*string", Required: true},
	{Name: "IdentityProviderDetails", Flag: "identity-provider-details", Type: "map[string]string", Required: false},
	{Name: "IdentityProviderName", Flag: "identity-provider-name", Type: "*string", Required: false},
	{Name: "IdentityProviderType", Flag: "identity-provider-type", Type: "types.IdentityProviderType", Required: false},
}

var fields_update_ip_access_settings = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "IpAccessSettingsArn", Flag: "ip-access-settings-arn", Type: "*string", Required: true},
	{Name: "IpRules", Flag: "ip-rules", Type: "[]types.IpRule", Required: false},
}

var fields_update_network_settings = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "NetworkSettingsArn", Flag: "network-settings-arn", Type: "*string", Required: true},
	{Name: "SecurityGroupIds", Flag: "security-group-ids", Type: "[]string", Required: false},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: false},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: false},
}

var fields_update_portal = []leanruntime.Field{
	{Name: "AuthenticationType", Flag: "authentication-type", Type: "types.AuthenticationType", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "InstanceType", Flag: "instance-type", Type: "types.InstanceType", Required: false},
	{Name: "MaxConcurrentSessions", Flag: "max-concurrent-sessions", Type: "*int32", Required: false},
	{Name: "PortalArn", Flag: "portal-arn", Type: "*string", Required: true},
	{Name: "PortalCustomDomain", Flag: "portal-custom-domain", Type: "*string", Required: false},
}

var fields_update_session_logger = []leanruntime.Field{
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "EventFilter", Flag: "event-filter", Type: "types.EventFilter", Required: false},
	{Name: "LogConfiguration", Flag: "log-configuration", Type: "*types.LogConfiguration", Required: false},
	{Name: "SessionLoggerArn", Flag: "session-logger-arn", Type: "*string", Required: true},
}

var fields_update_trust_store = []leanruntime.Field{
	{Name: "CertificatesToAdd", Flag: "certificates-to-add", Type: "[][]byte", Required: false},
	{Name: "CertificatesToDelete", Flag: "certificates-to-delete", Type: "[]string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "TrustStoreArn", Flag: "trust-store-arn", Type: "*string", Required: true},
}

var fields_update_user_access_logging_settings = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "KinesisStreamArn", Flag: "kinesis-stream-arn", Type: "*string", Required: false},
	{Name: "UserAccessLoggingSettingsArn", Flag: "user-access-logging-settings-arn", Type: "*string", Required: true},
}

var fields_update_user_settings = []leanruntime.Field{
	{Name: "BrandingConfigurationInput", Flag: "branding-configuration-input", Type: "*types.BrandingConfigurationUpdateInput", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CookieSynchronizationConfiguration", Flag: "cookie-synchronization-configuration", Type: "*types.CookieSynchronizationConfiguration", Required: false},
	{Name: "CopyAllowed", Flag: "copy-allowed", Type: "types.EnabledType", Required: false},
	{Name: "DeepLinkAllowed", Flag: "deep-link-allowed", Type: "types.EnabledType", Required: false},
	{Name: "DisconnectTimeoutInMinutes", Flag: "disconnect-timeout-in-minutes", Type: "*int32", Required: false},
	{Name: "DownloadAllowed", Flag: "download-allowed", Type: "types.EnabledType", Required: false},
	{Name: "IdleDisconnectTimeoutInMinutes", Flag: "idle-disconnect-timeout-in-minutes", Type: "*int32", Required: false},
	{Name: "PasteAllowed", Flag: "paste-allowed", Type: "types.EnabledType", Required: false},
	{Name: "PrintAllowed", Flag: "print-allowed", Type: "types.EnabledType", Required: false},
	{Name: "ToolbarConfiguration", Flag: "toolbar-configuration", Type: "*types.ToolbarConfiguration", Required: false},
	{Name: "UploadAllowed", Flag: "upload-allowed", Type: "types.EnabledType", Required: false},
	{Name: "UserSettingsArn", Flag: "user-settings-arn", Type: "*string", Required: true},
	{Name: "WebAuthnAllowed", Flag: "web-authn-allowed", Type: "types.EnabledType", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-browser-settings": {
			Name:   "associate-browser-settings",
			Fields: fields_associate_browser_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateBrowserSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_browser_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateBrowserSettings(ctx, input)
			},
		},
		"associate-data-protection-settings": {
			Name:   "associate-data-protection-settings",
			Fields: fields_associate_data_protection_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateDataProtectionSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_data_protection_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateDataProtectionSettings(ctx, input)
			},
		},
		"associate-ip-access-settings": {
			Name:   "associate-ip-access-settings",
			Fields: fields_associate_ip_access_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateIpAccessSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_ip_access_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateIpAccessSettings(ctx, input)
			},
		},
		"associate-network-settings": {
			Name:   "associate-network-settings",
			Fields: fields_associate_network_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateNetworkSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_network_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateNetworkSettings(ctx, input)
			},
		},
		"associate-session-logger": {
			Name:   "associate-session-logger",
			Fields: fields_associate_session_logger,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateSessionLoggerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_session_logger, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateSessionLogger(ctx, input)
			},
		},
		"associate-trust-store": {
			Name:   "associate-trust-store",
			Fields: fields_associate_trust_store,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateTrustStoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_trust_store, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateTrustStore(ctx, input)
			},
		},
		"associate-user-access-logging-settings": {
			Name:   "associate-user-access-logging-settings",
			Fields: fields_associate_user_access_logging_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateUserAccessLoggingSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_user_access_logging_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateUserAccessLoggingSettings(ctx, input)
			},
		},
		"associate-user-settings": {
			Name:   "associate-user-settings",
			Fields: fields_associate_user_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateUserSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_user_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateUserSettings(ctx, input)
			},
		},
		"create-browser-settings": {
			Name:   "create-browser-settings",
			Fields: fields_create_browser_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBrowserSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_browser_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBrowserSettings(ctx, input)
			},
		},
		"create-data-protection-settings": {
			Name:   "create-data-protection-settings",
			Fields: fields_create_data_protection_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDataProtectionSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_data_protection_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDataProtectionSettings(ctx, input)
			},
		},
		"create-identity-provider": {
			Name:   "create-identity-provider",
			Fields: fields_create_identity_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateIdentityProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_identity_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateIdentityProvider(ctx, input)
			},
		},
		"create-ip-access-settings": {
			Name:   "create-ip-access-settings",
			Fields: fields_create_ip_access_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateIpAccessSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_ip_access_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateIpAccessSettings(ctx, input)
			},
		},
		"create-network-settings": {
			Name:   "create-network-settings",
			Fields: fields_create_network_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateNetworkSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_network_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateNetworkSettings(ctx, input)
			},
		},
		"create-portal": {
			Name:   "create-portal",
			Fields: fields_create_portal,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePortalInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_portal, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePortal(ctx, input)
			},
		},
		"create-session-logger": {
			Name:   "create-session-logger",
			Fields: fields_create_session_logger,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSessionLoggerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_session_logger, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSessionLogger(ctx, input)
			},
		},
		"create-trust-store": {
			Name:   "create-trust-store",
			Fields: fields_create_trust_store,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTrustStoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_trust_store, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTrustStore(ctx, input)
			},
		},
		"create-user-access-logging-settings": {
			Name:   "create-user-access-logging-settings",
			Fields: fields_create_user_access_logging_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateUserAccessLoggingSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_user_access_logging_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateUserAccessLoggingSettings(ctx, input)
			},
		},
		"create-user-settings": {
			Name:   "create-user-settings",
			Fields: fields_create_user_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateUserSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_user_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateUserSettings(ctx, input)
			},
		},
		"delete-browser-settings": {
			Name:   "delete-browser-settings",
			Fields: fields_delete_browser_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBrowserSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_browser_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBrowserSettings(ctx, input)
			},
		},
		"delete-data-protection-settings": {
			Name:   "delete-data-protection-settings",
			Fields: fields_delete_data_protection_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDataProtectionSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_data_protection_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDataProtectionSettings(ctx, input)
			},
		},
		"delete-identity-provider": {
			Name:   "delete-identity-provider",
			Fields: fields_delete_identity_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIdentityProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_identity_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIdentityProvider(ctx, input)
			},
		},
		"delete-ip-access-settings": {
			Name:   "delete-ip-access-settings",
			Fields: fields_delete_ip_access_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIpAccessSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_ip_access_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIpAccessSettings(ctx, input)
			},
		},
		"delete-network-settings": {
			Name:   "delete-network-settings",
			Fields: fields_delete_network_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteNetworkSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_network_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteNetworkSettings(ctx, input)
			},
		},
		"delete-portal": {
			Name:   "delete-portal",
			Fields: fields_delete_portal,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePortalInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_portal, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePortal(ctx, input)
			},
		},
		"delete-session-logger": {
			Name:   "delete-session-logger",
			Fields: fields_delete_session_logger,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSessionLoggerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_session_logger, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSessionLogger(ctx, input)
			},
		},
		"delete-trust-store": {
			Name:   "delete-trust-store",
			Fields: fields_delete_trust_store,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTrustStoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_trust_store, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTrustStore(ctx, input)
			},
		},
		"delete-user-access-logging-settings": {
			Name:   "delete-user-access-logging-settings",
			Fields: fields_delete_user_access_logging_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteUserAccessLoggingSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_user_access_logging_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteUserAccessLoggingSettings(ctx, input)
			},
		},
		"delete-user-settings": {
			Name:   "delete-user-settings",
			Fields: fields_delete_user_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteUserSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_user_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteUserSettings(ctx, input)
			},
		},
		"disassociate-browser-settings": {
			Name:   "disassociate-browser-settings",
			Fields: fields_disassociate_browser_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateBrowserSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_browser_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateBrowserSettings(ctx, input)
			},
		},
		"disassociate-data-protection-settings": {
			Name:   "disassociate-data-protection-settings",
			Fields: fields_disassociate_data_protection_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateDataProtectionSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_data_protection_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateDataProtectionSettings(ctx, input)
			},
		},
		"disassociate-ip-access-settings": {
			Name:   "disassociate-ip-access-settings",
			Fields: fields_disassociate_ip_access_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateIpAccessSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_ip_access_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateIpAccessSettings(ctx, input)
			},
		},
		"disassociate-network-settings": {
			Name:   "disassociate-network-settings",
			Fields: fields_disassociate_network_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateNetworkSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_network_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateNetworkSettings(ctx, input)
			},
		},
		"disassociate-session-logger": {
			Name:   "disassociate-session-logger",
			Fields: fields_disassociate_session_logger,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateSessionLoggerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_session_logger, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateSessionLogger(ctx, input)
			},
		},
		"disassociate-trust-store": {
			Name:   "disassociate-trust-store",
			Fields: fields_disassociate_trust_store,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateTrustStoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_trust_store, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateTrustStore(ctx, input)
			},
		},
		"disassociate-user-access-logging-settings": {
			Name:   "disassociate-user-access-logging-settings",
			Fields: fields_disassociate_user_access_logging_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateUserAccessLoggingSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_user_access_logging_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateUserAccessLoggingSettings(ctx, input)
			},
		},
		"disassociate-user-settings": {
			Name:   "disassociate-user-settings",
			Fields: fields_disassociate_user_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateUserSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_user_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateUserSettings(ctx, input)
			},
		},
		"expire-session": {
			Name:   "expire-session",
			Fields: fields_expire_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExpireSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_expire_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExpireSession(ctx, input)
			},
		},
		"get-browser-settings": {
			Name:   "get-browser-settings",
			Fields: fields_get_browser_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBrowserSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_browser_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBrowserSettings(ctx, input)
			},
		},
		"get-data-protection-settings": {
			Name:   "get-data-protection-settings",
			Fields: fields_get_data_protection_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDataProtectionSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_data_protection_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDataProtectionSettings(ctx, input)
			},
		},
		"get-identity-provider": {
			Name:   "get-identity-provider",
			Fields: fields_get_identity_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIdentityProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_identity_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIdentityProvider(ctx, input)
			},
		},
		"get-ip-access-settings": {
			Name:   "get-ip-access-settings",
			Fields: fields_get_ip_access_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIpAccessSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_ip_access_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIpAccessSettings(ctx, input)
			},
		},
		"get-network-settings": {
			Name:   "get-network-settings",
			Fields: fields_get_network_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetNetworkSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_network_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetNetworkSettings(ctx, input)
			},
		},
		"get-portal": {
			Name:   "get-portal",
			Fields: fields_get_portal,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPortalInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_portal, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPortal(ctx, input)
			},
		},
		"get-portal-service-provider-metadata": {
			Name:   "get-portal-service-provider-metadata",
			Fields: fields_get_portal_service_provider_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPortalServiceProviderMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_portal_service_provider_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPortalServiceProviderMetadata(ctx, input)
			},
		},
		"get-session": {
			Name:   "get-session",
			Fields: fields_get_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSession(ctx, input)
			},
		},
		"get-session-logger": {
			Name:   "get-session-logger",
			Fields: fields_get_session_logger,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSessionLoggerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_session_logger, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSessionLogger(ctx, input)
			},
		},
		"get-trust-store": {
			Name:   "get-trust-store",
			Fields: fields_get_trust_store,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTrustStoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_trust_store, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTrustStore(ctx, input)
			},
		},
		"get-trust-store-certificate": {
			Name:   "get-trust-store-certificate",
			Fields: fields_get_trust_store_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTrustStoreCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_trust_store_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTrustStoreCertificate(ctx, input)
			},
		},
		"get-user-access-logging-settings": {
			Name:   "get-user-access-logging-settings",
			Fields: fields_get_user_access_logging_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetUserAccessLoggingSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_user_access_logging_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetUserAccessLoggingSettings(ctx, input)
			},
		},
		"get-user-settings": {
			Name:   "get-user-settings",
			Fields: fields_get_user_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetUserSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_user_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetUserSettings(ctx, input)
			},
		},
		"list-browser-settings": {
			Name:   "list-browser-settings",
			Fields: fields_list_browser_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBrowserSettingsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_browser_settings, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBrowserSettings(ctx, input)
				}
				var results []*svc.ListBrowserSettingsOutput
				p := svc.NewListBrowserSettingsPaginator(client, input)
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
		"list-data-protection-settings": {
			Name:   "list-data-protection-settings",
			Fields: fields_list_data_protection_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDataProtectionSettingsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_data_protection_settings, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDataProtectionSettings(ctx, input)
				}
				var results []*svc.ListDataProtectionSettingsOutput
				p := svc.NewListDataProtectionSettingsPaginator(client, input)
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
		"list-identity-providers": {
			Name:   "list-identity-providers",
			Fields: fields_list_identity_providers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIdentityProvidersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_identity_providers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListIdentityProviders(ctx, input)
				}
				var results []*svc.ListIdentityProvidersOutput
				p := svc.NewListIdentityProvidersPaginator(client, input)
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
		"list-ip-access-settings": {
			Name:   "list-ip-access-settings",
			Fields: fields_list_ip_access_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIpAccessSettingsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_ip_access_settings, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListIpAccessSettings(ctx, input)
				}
				var results []*svc.ListIpAccessSettingsOutput
				p := svc.NewListIpAccessSettingsPaginator(client, input)
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
		"list-network-settings": {
			Name:   "list-network-settings",
			Fields: fields_list_network_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListNetworkSettingsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_network_settings, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListNetworkSettings(ctx, input)
				}
				var results []*svc.ListNetworkSettingsOutput
				p := svc.NewListNetworkSettingsPaginator(client, input)
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
		"list-portals": {
			Name:   "list-portals",
			Fields: fields_list_portals,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPortalsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_portals, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPortals(ctx, input)
				}
				var results []*svc.ListPortalsOutput
				p := svc.NewListPortalsPaginator(client, input)
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
		"list-session-loggers": {
			Name:   "list-session-loggers",
			Fields: fields_list_session_loggers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSessionLoggersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_session_loggers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSessionLoggers(ctx, input)
				}
				var results []*svc.ListSessionLoggersOutput
				p := svc.NewListSessionLoggersPaginator(client, input)
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
		"list-sessions": {
			Name:   "list-sessions",
			Fields: fields_list_sessions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSessionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_sessions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSessions(ctx, input)
				}
				var results []*svc.ListSessionsOutput
				p := svc.NewListSessionsPaginator(client, input)
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
		"list-tags-for-resource": {
			Name:   "list-tags-for-resource",
			Fields: fields_list_tags_for_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTagsForResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_tags_for_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTagsForResource(ctx, input)
			},
		},
		"list-trust-store-certificates": {
			Name:   "list-trust-store-certificates",
			Fields: fields_list_trust_store_certificates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTrustStoreCertificatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_trust_store_certificates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTrustStoreCertificates(ctx, input)
				}
				var results []*svc.ListTrustStoreCertificatesOutput
				p := svc.NewListTrustStoreCertificatesPaginator(client, input)
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
		"list-trust-stores": {
			Name:   "list-trust-stores",
			Fields: fields_list_trust_stores,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTrustStoresInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_trust_stores, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTrustStores(ctx, input)
				}
				var results []*svc.ListTrustStoresOutput
				p := svc.NewListTrustStoresPaginator(client, input)
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
		"list-user-access-logging-settings": {
			Name:   "list-user-access-logging-settings",
			Fields: fields_list_user_access_logging_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListUserAccessLoggingSettingsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_user_access_logging_settings, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListUserAccessLoggingSettings(ctx, input)
				}
				var results []*svc.ListUserAccessLoggingSettingsOutput
				p := svc.NewListUserAccessLoggingSettingsPaginator(client, input)
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
		"list-user-settings": {
			Name:   "list-user-settings",
			Fields: fields_list_user_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListUserSettingsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_user_settings, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListUserSettings(ctx, input)
				}
				var results []*svc.ListUserSettingsOutput
				p := svc.NewListUserSettingsPaginator(client, input)
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
		"update-browser-settings": {
			Name:   "update-browser-settings",
			Fields: fields_update_browser_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateBrowserSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_browser_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateBrowserSettings(ctx, input)
			},
		},
		"update-data-protection-settings": {
			Name:   "update-data-protection-settings",
			Fields: fields_update_data_protection_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDataProtectionSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_data_protection_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDataProtectionSettings(ctx, input)
			},
		},
		"update-identity-provider": {
			Name:   "update-identity-provider",
			Fields: fields_update_identity_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateIdentityProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_identity_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateIdentityProvider(ctx, input)
			},
		},
		"update-ip-access-settings": {
			Name:   "update-ip-access-settings",
			Fields: fields_update_ip_access_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateIpAccessSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_ip_access_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateIpAccessSettings(ctx, input)
			},
		},
		"update-network-settings": {
			Name:   "update-network-settings",
			Fields: fields_update_network_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateNetworkSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_network_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateNetworkSettings(ctx, input)
			},
		},
		"update-portal": {
			Name:   "update-portal",
			Fields: fields_update_portal,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePortalInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_portal, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePortal(ctx, input)
			},
		},
		"update-session-logger": {
			Name:   "update-session-logger",
			Fields: fields_update_session_logger,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSessionLoggerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_session_logger, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSessionLogger(ctx, input)
			},
		},
		"update-trust-store": {
			Name:   "update-trust-store",
			Fields: fields_update_trust_store,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTrustStoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_trust_store, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTrustStore(ctx, input)
			},
		},
		"update-user-access-logging-settings": {
			Name:   "update-user-access-logging-settings",
			Fields: fields_update_user_access_logging_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateUserAccessLoggingSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_user_access_logging_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateUserAccessLoggingSettings(ctx, input)
			},
		},
		"update-user-settings": {
			Name:   "update-user-settings",
			Fields: fields_update_user_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateUserSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_user_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateUserSettings(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("workspacesweb", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
