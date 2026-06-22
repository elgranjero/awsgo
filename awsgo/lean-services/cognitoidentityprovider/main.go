package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

var fields_add_custom_attributes = []leanruntime.Field{
	{Name: "CustomAttributes", Flag: "custom-attributes", Type: "[]types.SchemaAttributeType", Required: true},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_add_user_pool_client_secret = []leanruntime.Field{
	{Name: "ClientId", Flag: "client-id", Type: "*string", Required: true},
	{Name: "ClientSecret", Flag: "client-secret", Type: "*string", Required: false},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_admin_add_user_to_group = []leanruntime.Field{
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: true},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
	{Name: "Username", Flag: "username", Type: "*string", Required: true},
}

var fields_admin_confirm_sign_up = []leanruntime.Field{
	{Name: "ClientMetadata", Flag: "client-metadata", Type: "map[string]string", Required: false},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
	{Name: "Username", Flag: "username", Type: "*string", Required: true},
}

var fields_admin_create_user = []leanruntime.Field{
	{Name: "ClientMetadata", Flag: "client-metadata", Type: "map[string]string", Required: false},
	{Name: "DesiredDeliveryMediums", Flag: "desired-delivery-mediums", Type: "[]types.DeliveryMediumType", Required: false},
	{Name: "ForceAliasCreation", Flag: "force-alias-creation", Type: "bool", Required: false},
	{Name: "MessageAction", Flag: "message-action", Type: "types.MessageActionType", Required: false},
	{Name: "TemporaryPassword", Flag: "temporary-password", Type: "*string", Required: false},
	{Name: "UserAttributes", Flag: "user-attributes", Type: "[]types.AttributeType", Required: false},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
	{Name: "Username", Flag: "username", Type: "*string", Required: true},
	{Name: "ValidationData", Flag: "validation-data", Type: "[]types.AttributeType", Required: false},
}

var fields_admin_delete_user = []leanruntime.Field{
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
	{Name: "Username", Flag: "username", Type: "*string", Required: true},
}

var fields_admin_delete_user_attributes = []leanruntime.Field{
	{Name: "UserAttributeNames", Flag: "user-attribute-names", Type: "[]string", Required: true},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
	{Name: "Username", Flag: "username", Type: "*string", Required: true},
}

var fields_admin_disable_provider_for_user = []leanruntime.Field{
	{Name: "User", Flag: "user", Type: "*types.ProviderUserIdentifierType", Required: true},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_admin_disable_user = []leanruntime.Field{
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
	{Name: "Username", Flag: "username", Type: "*string", Required: true},
}

var fields_admin_enable_user = []leanruntime.Field{
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
	{Name: "Username", Flag: "username", Type: "*string", Required: true},
}

var fields_admin_forget_device = []leanruntime.Field{
	{Name: "DeviceKey", Flag: "device-key", Type: "*string", Required: true},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
	{Name: "Username", Flag: "username", Type: "*string", Required: true},
}

var fields_admin_get_device = []leanruntime.Field{
	{Name: "DeviceKey", Flag: "device-key", Type: "*string", Required: true},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
	{Name: "Username", Flag: "username", Type: "*string", Required: true},
}

var fields_admin_get_user = []leanruntime.Field{
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
	{Name: "Username", Flag: "username", Type: "*string", Required: true},
}

var fields_admin_initiate_auth = []leanruntime.Field{
	{Name: "AnalyticsMetadata", Flag: "analytics-metadata", Type: "*types.AnalyticsMetadataType", Required: false},
	{Name: "AuthFlow", Flag: "auth-flow", Type: "types.AuthFlowType", Required: true},
	{Name: "AuthParameters", Flag: "auth-parameters", Type: "map[string]string", Required: false},
	{Name: "ClientId", Flag: "client-id", Type: "*string", Required: true},
	{Name: "ClientMetadata", Flag: "client-metadata", Type: "map[string]string", Required: false},
	{Name: "ContextData", Flag: "context-data", Type: "*types.ContextDataType", Required: false},
	{Name: "Session", Flag: "session", Type: "*string", Required: false},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_admin_link_provider_for_user = []leanruntime.Field{
	{Name: "DestinationUser", Flag: "destination-user", Type: "*types.ProviderUserIdentifierType", Required: true},
	{Name: "SourceUser", Flag: "source-user", Type: "*types.ProviderUserIdentifierType", Required: true},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_admin_list_devices = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "PaginationToken", Flag: "pagination-token", Type: "*string", Required: false},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
	{Name: "Username", Flag: "username", Type: "*string", Required: true},
}

var fields_admin_list_groups_for_user = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
	{Name: "Username", Flag: "username", Type: "*string", Required: true},
}

var fields_admin_list_user_auth_events = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
	{Name: "Username", Flag: "username", Type: "*string", Required: true},
}

var fields_admin_remove_user_from_group = []leanruntime.Field{
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: true},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
	{Name: "Username", Flag: "username", Type: "*string", Required: true},
}

var fields_admin_reset_user_password = []leanruntime.Field{
	{Name: "ClientMetadata", Flag: "client-metadata", Type: "map[string]string", Required: false},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
	{Name: "Username", Flag: "username", Type: "*string", Required: true},
}

var fields_admin_respond_to_auth_challenge = []leanruntime.Field{
	{Name: "AnalyticsMetadata", Flag: "analytics-metadata", Type: "*types.AnalyticsMetadataType", Required: false},
	{Name: "ChallengeName", Flag: "challenge-name", Type: "types.ChallengeNameType", Required: true},
	{Name: "ChallengeResponses", Flag: "challenge-responses", Type: "map[string]string", Required: false},
	{Name: "ClientId", Flag: "client-id", Type: "*string", Required: true},
	{Name: "ClientMetadata", Flag: "client-metadata", Type: "map[string]string", Required: false},
	{Name: "ContextData", Flag: "context-data", Type: "*types.ContextDataType", Required: false},
	{Name: "Session", Flag: "session", Type: "*string", Required: false},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_admin_set_user_mfa_preference = []leanruntime.Field{
	{Name: "EmailMfaSettings", Flag: "email-mfa-settings", Type: "*types.EmailMfaSettingsType", Required: false},
	{Name: "SMSMfaSettings", Flag: "sms-mfa-settings", Type: "*types.SMSMfaSettingsType", Required: false},
	{Name: "SoftwareTokenMfaSettings", Flag: "software-token-mfa-settings", Type: "*types.SoftwareTokenMfaSettingsType", Required: false},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
	{Name: "Username", Flag: "username", Type: "*string", Required: true},
}

var fields_admin_set_user_password = []leanruntime.Field{
	{Name: "Password", Flag: "password", Type: "*string", Required: true},
	{Name: "Permanent", Flag: "permanent", Type: "bool", Required: false},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
	{Name: "Username", Flag: "username", Type: "*string", Required: true},
}

var fields_admin_set_user_settings = []leanruntime.Field{
	{Name: "MFAOptions", Flag: "mfa-options", Type: "[]types.MFAOptionType", Required: true},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
	{Name: "Username", Flag: "username", Type: "*string", Required: true},
}

var fields_admin_update_auth_event_feedback = []leanruntime.Field{
	{Name: "EventId", Flag: "event-id", Type: "*string", Required: true},
	{Name: "FeedbackValue", Flag: "feedback-value", Type: "types.FeedbackValueType", Required: true},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
	{Name: "Username", Flag: "username", Type: "*string", Required: true},
}

var fields_admin_update_device_status = []leanruntime.Field{
	{Name: "DeviceKey", Flag: "device-key", Type: "*string", Required: true},
	{Name: "DeviceRememberedStatus", Flag: "device-remembered-status", Type: "types.DeviceRememberedStatusType", Required: false},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
	{Name: "Username", Flag: "username", Type: "*string", Required: true},
}

var fields_admin_update_user_attributes = []leanruntime.Field{
	{Name: "ClientMetadata", Flag: "client-metadata", Type: "map[string]string", Required: false},
	{Name: "UserAttributes", Flag: "user-attributes", Type: "[]types.AttributeType", Required: true},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
	{Name: "Username", Flag: "username", Type: "*string", Required: true},
}

var fields_admin_user_global_sign_out = []leanruntime.Field{
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
	{Name: "Username", Flag: "username", Type: "*string", Required: true},
}

var fields_associate_software_token = []leanruntime.Field{
	{Name: "AccessToken", Flag: "access-token", Type: "*string", Required: false},
	{Name: "Session", Flag: "session", Type: "*string", Required: false},
}

var fields_change_password = []leanruntime.Field{
	{Name: "AccessToken", Flag: "access-token", Type: "*string", Required: true},
	{Name: "PreviousPassword", Flag: "previous-password", Type: "*string", Required: false},
	{Name: "ProposedPassword", Flag: "proposed-password", Type: "*string", Required: true},
}

var fields_complete_web_authn_registration = []leanruntime.Field{
	{Name: "AccessToken", Flag: "access-token", Type: "*string", Required: true},
	{Name: "Credential", Flag: "credential", Type: "document.Interface", Required: true},
}

var fields_confirm_device = []leanruntime.Field{
	{Name: "AccessToken", Flag: "access-token", Type: "*string", Required: true},
	{Name: "DeviceKey", Flag: "device-key", Type: "*string", Required: true},
	{Name: "DeviceName", Flag: "device-name", Type: "*string", Required: false},
	{Name: "DeviceSecretVerifierConfig", Flag: "device-secret-verifier-config", Type: "*types.DeviceSecretVerifierConfigType", Required: false},
}

var fields_confirm_forgot_password = []leanruntime.Field{
	{Name: "AnalyticsMetadata", Flag: "analytics-metadata", Type: "*types.AnalyticsMetadataType", Required: false},
	{Name: "ClientId", Flag: "client-id", Type: "*string", Required: true},
	{Name: "ClientMetadata", Flag: "client-metadata", Type: "map[string]string", Required: false},
	{Name: "ConfirmationCode", Flag: "confirmation-code", Type: "*string", Required: true},
	{Name: "Password", Flag: "password", Type: "*string", Required: true},
	{Name: "SecretHash", Flag: "secret-hash", Type: "*string", Required: false},
	{Name: "UserContextData", Flag: "user-context-data", Type: "*types.UserContextDataType", Required: false},
	{Name: "Username", Flag: "username", Type: "*string", Required: true},
}

var fields_confirm_sign_up = []leanruntime.Field{
	{Name: "AnalyticsMetadata", Flag: "analytics-metadata", Type: "*types.AnalyticsMetadataType", Required: false},
	{Name: "ClientId", Flag: "client-id", Type: "*string", Required: true},
	{Name: "ClientMetadata", Flag: "client-metadata", Type: "map[string]string", Required: false},
	{Name: "ConfirmationCode", Flag: "confirmation-code", Type: "*string", Required: true},
	{Name: "ForceAliasCreation", Flag: "force-alias-creation", Type: "bool", Required: false},
	{Name: "SecretHash", Flag: "secret-hash", Type: "*string", Required: false},
	{Name: "Session", Flag: "session", Type: "*string", Required: false},
	{Name: "UserContextData", Flag: "user-context-data", Type: "*types.UserContextDataType", Required: false},
	{Name: "Username", Flag: "username", Type: "*string", Required: true},
}

var fields_create_group = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: true},
	{Name: "Precedence", Flag: "precedence", Type: "*int32", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_create_identity_provider = []leanruntime.Field{
	{Name: "AttributeMapping", Flag: "attribute-mapping", Type: "map[string]string", Required: false},
	{Name: "IdpIdentifiers", Flag: "idp-identifiers", Type: "[]string", Required: false},
	{Name: "ProviderDetails", Flag: "provider-details", Type: "map[string]string", Required: true},
	{Name: "ProviderName", Flag: "provider-name", Type: "*string", Required: true},
	{Name: "ProviderType", Flag: "provider-type", Type: "types.IdentityProviderTypeType", Required: true},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_create_managed_login_branding = []leanruntime.Field{
	{Name: "Assets", Flag: "assets", Type: "[]types.AssetType", Required: false},
	{Name: "ClientId", Flag: "client-id", Type: "*string", Required: true},
	{Name: "Settings", Flag: "settings", Type: "document.Interface", Required: false},
	{Name: "UseCognitoProvidedValues", Flag: "use-cognito-provided-values", Type: "bool", Required: false},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_create_resource_server = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Scopes", Flag: "scopes", Type: "[]types.ResourceServerScopeType", Required: false},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_create_terms = []leanruntime.Field{
	{Name: "ClientId", Flag: "client-id", Type: "*string", Required: true},
	{Name: "Enforcement", Flag: "enforcement", Type: "types.TermsEnforcementType", Required: true},
	{Name: "Links", Flag: "links", Type: "map[string]string", Required: false},
	{Name: "TermsName", Flag: "terms-name", Type: "*string", Required: true},
	{Name: "TermsSource", Flag: "terms-source", Type: "types.TermsSourceType", Required: true},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_create_user_import_job = []leanruntime.Field{
	{Name: "CloudWatchLogsRoleArn", Flag: "cloud-watch-logs-role-arn", Type: "*string", Required: true},
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: true},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_create_user_pool = []leanruntime.Field{
	{Name: "AccountRecoverySetting", Flag: "account-recovery-setting", Type: "*types.AccountRecoverySettingType", Required: false},
	{Name: "AdminCreateUserConfig", Flag: "admin-create-user-config", Type: "*types.AdminCreateUserConfigType", Required: false},
	{Name: "AliasAttributes", Flag: "alias-attributes", Type: "[]types.AliasAttributeType", Required: false},
	{Name: "AutoVerifiedAttributes", Flag: "auto-verified-attributes", Type: "[]types.VerifiedAttributeType", Required: false},
	{Name: "DeletionProtection", Flag: "deletion-protection", Type: "types.DeletionProtectionType", Required: false},
	{Name: "DeviceConfiguration", Flag: "device-configuration", Type: "*types.DeviceConfigurationType", Required: false},
	{Name: "EmailConfiguration", Flag: "email-configuration", Type: "*types.EmailConfigurationType", Required: false},
	{Name: "EmailVerificationMessage", Flag: "email-verification-message", Type: "*string", Required: false},
	{Name: "EmailVerificationSubject", Flag: "email-verification-subject", Type: "*string", Required: false},
	{Name: "LambdaConfig", Flag: "lambda-config", Type: "*types.LambdaConfigType", Required: false},
	{Name: "MfaConfiguration", Flag: "mfa-configuration", Type: "types.UserPoolMfaType", Required: false},
	{Name: "Policies", Flag: "policies", Type: "*types.UserPoolPolicyType", Required: false},
	{Name: "PoolName", Flag: "pool-name", Type: "*string", Required: true},
	{Name: "Schema", Flag: "schema", Type: "[]types.SchemaAttributeType", Required: false},
	{Name: "SmsAuthenticationMessage", Flag: "sms-authentication-message", Type: "*string", Required: false},
	{Name: "SmsConfiguration", Flag: "sms-configuration", Type: "*types.SmsConfigurationType", Required: false},
	{Name: "SmsVerificationMessage", Flag: "sms-verification-message", Type: "*string", Required: false},
	{Name: "UserAttributeUpdateSettings", Flag: "user-attribute-update-settings", Type: "*types.UserAttributeUpdateSettingsType", Required: false},
	{Name: "UserPoolAddOns", Flag: "user-pool-add-ons", Type: "*types.UserPoolAddOnsType", Required: false},
	{Name: "UserPoolTags", Flag: "user-pool-tags", Type: "map[string]string", Required: false},
	{Name: "UserPoolTier", Flag: "user-pool-tier", Type: "types.UserPoolTierType", Required: false},
	{Name: "UsernameAttributes", Flag: "username-attributes", Type: "[]types.UsernameAttributeType", Required: false},
	{Name: "UsernameConfiguration", Flag: "username-configuration", Type: "*types.UsernameConfigurationType", Required: false},
	{Name: "VerificationMessageTemplate", Flag: "verification-message-template", Type: "*types.VerificationMessageTemplateType", Required: false},
}

var fields_create_user_pool_client = []leanruntime.Field{
	{Name: "AccessTokenValidity", Flag: "access-token-validity", Type: "*int32", Required: false},
	{Name: "AllowedOAuthFlows", Flag: "allowed-oauth-flows", Type: "[]types.OAuthFlowType", Required: false},
	{Name: "AllowedOAuthFlowsUserPoolClient", Flag: "allowed-oauth-flows-user-pool-client", Type: "bool", Required: false},
	{Name: "AllowedOAuthScopes", Flag: "allowed-oauth-scopes", Type: "[]string", Required: false},
	{Name: "AnalyticsConfiguration", Flag: "analytics-configuration", Type: "*types.AnalyticsConfigurationType", Required: false},
	{Name: "AuthSessionValidity", Flag: "auth-session-validity", Type: "*int32", Required: false},
	{Name: "CallbackURLs", Flag: "callback-urls", Type: "[]string", Required: false},
	{Name: "ClientName", Flag: "client-name", Type: "*string", Required: true},
	{Name: "ClientSecret", Flag: "client-secret", Type: "*string", Required: false},
	{Name: "DefaultRedirectURI", Flag: "default-redirect-uri", Type: "*string", Required: false},
	{Name: "EnablePropagateAdditionalUserContextData", Flag: "enable-propagate-additional-user-context-data", Type: "*bool", Required: false},
	{Name: "EnableTokenRevocation", Flag: "enable-token-revocation", Type: "*bool", Required: false},
	{Name: "ExplicitAuthFlows", Flag: "explicit-auth-flows", Type: "[]types.ExplicitAuthFlowsType", Required: false},
	{Name: "GenerateSecret", Flag: "generate-secret", Type: "bool", Required: false},
	{Name: "IdTokenValidity", Flag: "id-token-validity", Type: "*int32", Required: false},
	{Name: "LogoutURLs", Flag: "logout-urls", Type: "[]string", Required: false},
	{Name: "PreventUserExistenceErrors", Flag: "prevent-user-existence-errors", Type: "types.PreventUserExistenceErrorTypes", Required: false},
	{Name: "ReadAttributes", Flag: "read-attributes", Type: "[]string", Required: false},
	{Name: "RefreshTokenRotation", Flag: "refresh-token-rotation", Type: "*types.RefreshTokenRotationType", Required: false},
	{Name: "RefreshTokenValidity", Flag: "refresh-token-validity", Type: "int32", Required: false},
	{Name: "SupportedIdentityProviders", Flag: "supported-identity-providers", Type: "[]string", Required: false},
	{Name: "TokenValidityUnits", Flag: "token-validity-units", Type: "*types.TokenValidityUnitsType", Required: false},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
	{Name: "WriteAttributes", Flag: "write-attributes", Type: "[]string", Required: false},
}

var fields_create_user_pool_domain = []leanruntime.Field{
	{Name: "CustomDomainConfig", Flag: "custom-domain-config", Type: "*types.CustomDomainConfigType", Required: false},
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "ManagedLoginVersion", Flag: "managed-login-version", Type: "*int32", Required: false},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_delete_group = []leanruntime.Field{
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: true},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_delete_identity_provider = []leanruntime.Field{
	{Name: "ProviderName", Flag: "provider-name", Type: "*string", Required: true},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_delete_managed_login_branding = []leanruntime.Field{
	{Name: "ManagedLoginBrandingId", Flag: "managed-login-branding-id", Type: "*string", Required: true},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_delete_resource_server = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_delete_terms = []leanruntime.Field{
	{Name: "TermsId", Flag: "terms-id", Type: "*string", Required: true},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_delete_user = []leanruntime.Field{
	{Name: "AccessToken", Flag: "access-token", Type: "*string", Required: true},
}

var fields_delete_user_attributes = []leanruntime.Field{
	{Name: "AccessToken", Flag: "access-token", Type: "*string", Required: true},
	{Name: "UserAttributeNames", Flag: "user-attribute-names", Type: "[]string", Required: true},
}

var fields_delete_user_pool = []leanruntime.Field{
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_delete_user_pool_client = []leanruntime.Field{
	{Name: "ClientId", Flag: "client-id", Type: "*string", Required: true},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_delete_user_pool_client_secret = []leanruntime.Field{
	{Name: "ClientId", Flag: "client-id", Type: "*string", Required: true},
	{Name: "ClientSecretId", Flag: "client-secret-id", Type: "*string", Required: true},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_delete_user_pool_domain = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_delete_web_authn_credential = []leanruntime.Field{
	{Name: "AccessToken", Flag: "access-token", Type: "*string", Required: true},
	{Name: "CredentialId", Flag: "credential-id", Type: "*string", Required: true},
}

var fields_describe_identity_provider = []leanruntime.Field{
	{Name: "ProviderName", Flag: "provider-name", Type: "*string", Required: true},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_describe_managed_login_branding = []leanruntime.Field{
	{Name: "ManagedLoginBrandingId", Flag: "managed-login-branding-id", Type: "*string", Required: true},
	{Name: "ReturnMergedResources", Flag: "return-merged-resources", Type: "bool", Required: false},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_describe_managed_login_branding_by_client = []leanruntime.Field{
	{Name: "ClientId", Flag: "client-id", Type: "*string", Required: true},
	{Name: "ReturnMergedResources", Flag: "return-merged-resources", Type: "bool", Required: false},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_describe_resource_server = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_describe_risk_configuration = []leanruntime.Field{
	{Name: "ClientId", Flag: "client-id", Type: "*string", Required: false},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_describe_terms = []leanruntime.Field{
	{Name: "TermsId", Flag: "terms-id", Type: "*string", Required: true},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_describe_user_import_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_describe_user_pool = []leanruntime.Field{
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_describe_user_pool_client = []leanruntime.Field{
	{Name: "ClientId", Flag: "client-id", Type: "*string", Required: true},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_describe_user_pool_domain = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
}

var fields_forget_device = []leanruntime.Field{
	{Name: "AccessToken", Flag: "access-token", Type: "*string", Required: false},
	{Name: "DeviceKey", Flag: "device-key", Type: "*string", Required: true},
}

var fields_forgot_password = []leanruntime.Field{
	{Name: "AnalyticsMetadata", Flag: "analytics-metadata", Type: "*types.AnalyticsMetadataType", Required: false},
	{Name: "ClientId", Flag: "client-id", Type: "*string", Required: true},
	{Name: "ClientMetadata", Flag: "client-metadata", Type: "map[string]string", Required: false},
	{Name: "SecretHash", Flag: "secret-hash", Type: "*string", Required: false},
	{Name: "UserContextData", Flag: "user-context-data", Type: "*types.UserContextDataType", Required: false},
	{Name: "Username", Flag: "username", Type: "*string", Required: true},
}

var fields_get_csv_header = []leanruntime.Field{
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_get_device = []leanruntime.Field{
	{Name: "AccessToken", Flag: "access-token", Type: "*string", Required: false},
	{Name: "DeviceKey", Flag: "device-key", Type: "*string", Required: true},
}

var fields_get_group = []leanruntime.Field{
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: true},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_get_identity_provider_by_identifier = []leanruntime.Field{
	{Name: "IdpIdentifier", Flag: "idp-identifier", Type: "*string", Required: true},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_get_log_delivery_configuration = []leanruntime.Field{
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_get_signing_certificate = []leanruntime.Field{
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_get_tokens_from_refresh_token = []leanruntime.Field{
	{Name: "ClientId", Flag: "client-id", Type: "*string", Required: true},
	{Name: "ClientMetadata", Flag: "client-metadata", Type: "map[string]string", Required: false},
	{Name: "ClientSecret", Flag: "client-secret", Type: "*string", Required: false},
	{Name: "DeviceKey", Flag: "device-key", Type: "*string", Required: false},
	{Name: "RefreshToken", Flag: "refresh-token", Type: "*string", Required: true},
}

var fields_get_ui_customization = []leanruntime.Field{
	{Name: "ClientId", Flag: "client-id", Type: "*string", Required: false},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_get_user = []leanruntime.Field{
	{Name: "AccessToken", Flag: "access-token", Type: "*string", Required: true},
}

var fields_get_user_attribute_verification_code = []leanruntime.Field{
	{Name: "AccessToken", Flag: "access-token", Type: "*string", Required: true},
	{Name: "AttributeName", Flag: "attribute-name", Type: "*string", Required: true},
	{Name: "ClientMetadata", Flag: "client-metadata", Type: "map[string]string", Required: false},
}

var fields_get_user_auth_factors = []leanruntime.Field{
	{Name: "AccessToken", Flag: "access-token", Type: "*string", Required: true},
}

var fields_get_user_pool_mfa_config = []leanruntime.Field{
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_global_sign_out = []leanruntime.Field{
	{Name: "AccessToken", Flag: "access-token", Type: "*string", Required: true},
}

var fields_initiate_auth = []leanruntime.Field{
	{Name: "AnalyticsMetadata", Flag: "analytics-metadata", Type: "*types.AnalyticsMetadataType", Required: false},
	{Name: "AuthFlow", Flag: "auth-flow", Type: "types.AuthFlowType", Required: true},
	{Name: "AuthParameters", Flag: "auth-parameters", Type: "map[string]string", Required: false},
	{Name: "ClientId", Flag: "client-id", Type: "*string", Required: true},
	{Name: "ClientMetadata", Flag: "client-metadata", Type: "map[string]string", Required: false},
	{Name: "Session", Flag: "session", Type: "*string", Required: false},
	{Name: "UserContextData", Flag: "user-context-data", Type: "*types.UserContextDataType", Required: false},
}

var fields_list_devices = []leanruntime.Field{
	{Name: "AccessToken", Flag: "access-token", Type: "*string", Required: true},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "PaginationToken", Flag: "pagination-token", Type: "*string", Required: false},
}

var fields_list_groups = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_list_identity_providers = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_list_resource_servers = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_terms = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_list_user_import_jobs = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: true},
	{Name: "PaginationToken", Flag: "pagination-token", Type: "*string", Required: false},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_list_user_pool_client_secrets = []leanruntime.Field{
	{Name: "ClientId", Flag: "client-id", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_list_user_pool_clients = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_list_user_pools = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_users = []leanruntime.Field{
	{Name: "AttributesToGet", Flag: "attributes-to-get", Type: "[]string", Required: false},
	{Name: "Filter", Flag: "filter", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "PaginationToken", Flag: "pagination-token", Type: "*string", Required: false},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_list_users_in_group = []leanruntime.Field{
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: true},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_list_web_authn_credentials = []leanruntime.Field{
	{Name: "AccessToken", Flag: "access-token", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_resend_confirmation_code = []leanruntime.Field{
	{Name: "AnalyticsMetadata", Flag: "analytics-metadata", Type: "*types.AnalyticsMetadataType", Required: false},
	{Name: "ClientId", Flag: "client-id", Type: "*string", Required: true},
	{Name: "ClientMetadata", Flag: "client-metadata", Type: "map[string]string", Required: false},
	{Name: "SecretHash", Flag: "secret-hash", Type: "*string", Required: false},
	{Name: "UserContextData", Flag: "user-context-data", Type: "*types.UserContextDataType", Required: false},
	{Name: "Username", Flag: "username", Type: "*string", Required: true},
}

var fields_respond_to_auth_challenge = []leanruntime.Field{
	{Name: "AnalyticsMetadata", Flag: "analytics-metadata", Type: "*types.AnalyticsMetadataType", Required: false},
	{Name: "ChallengeName", Flag: "challenge-name", Type: "types.ChallengeNameType", Required: true},
	{Name: "ChallengeResponses", Flag: "challenge-responses", Type: "map[string]string", Required: false},
	{Name: "ClientId", Flag: "client-id", Type: "*string", Required: true},
	{Name: "ClientMetadata", Flag: "client-metadata", Type: "map[string]string", Required: false},
	{Name: "Session", Flag: "session", Type: "*string", Required: false},
	{Name: "UserContextData", Flag: "user-context-data", Type: "*types.UserContextDataType", Required: false},
}

var fields_revoke_token = []leanruntime.Field{
	{Name: "ClientId", Flag: "client-id", Type: "*string", Required: true},
	{Name: "ClientSecret", Flag: "client-secret", Type: "*string", Required: false},
	{Name: "Token", Flag: "token", Type: "*string", Required: true},
}

var fields_set_log_delivery_configuration = []leanruntime.Field{
	{Name: "LogConfigurations", Flag: "log-configurations", Type: "[]types.LogConfigurationType", Required: true},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_set_risk_configuration = []leanruntime.Field{
	{Name: "AccountTakeoverRiskConfiguration", Flag: "account-takeover-risk-configuration", Type: "*types.AccountTakeoverRiskConfigurationType", Required: false},
	{Name: "ClientId", Flag: "client-id", Type: "*string", Required: false},
	{Name: "CompromisedCredentialsRiskConfiguration", Flag: "compromised-credentials-risk-configuration", Type: "*types.CompromisedCredentialsRiskConfigurationType", Required: false},
	{Name: "RiskExceptionConfiguration", Flag: "risk-exception-configuration", Type: "*types.RiskExceptionConfigurationType", Required: false},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_set_ui_customization = []leanruntime.Field{
	{Name: "CSS", Flag: "css", Type: "*string", Required: false},
	{Name: "ClientId", Flag: "client-id", Type: "*string", Required: false},
	{Name: "ImageFile", Flag: "image-file", Type: "[]byte", Required: false},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_set_user_mfa_preference = []leanruntime.Field{
	{Name: "AccessToken", Flag: "access-token", Type: "*string", Required: true},
	{Name: "EmailMfaSettings", Flag: "email-mfa-settings", Type: "*types.EmailMfaSettingsType", Required: false},
	{Name: "SMSMfaSettings", Flag: "sms-mfa-settings", Type: "*types.SMSMfaSettingsType", Required: false},
	{Name: "SoftwareTokenMfaSettings", Flag: "software-token-mfa-settings", Type: "*types.SoftwareTokenMfaSettingsType", Required: false},
}

var fields_set_user_pool_mfa_config = []leanruntime.Field{
	{Name: "EmailMfaConfiguration", Flag: "email-mfa-configuration", Type: "*types.EmailMfaConfigType", Required: false},
	{Name: "MfaConfiguration", Flag: "mfa-configuration", Type: "types.UserPoolMfaType", Required: false},
	{Name: "SmsMfaConfiguration", Flag: "sms-mfa-configuration", Type: "*types.SmsMfaConfigType", Required: false},
	{Name: "SoftwareTokenMfaConfiguration", Flag: "software-token-mfa-configuration", Type: "*types.SoftwareTokenMfaConfigType", Required: false},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
	{Name: "WebAuthnConfiguration", Flag: "web-authn-configuration", Type: "*types.WebAuthnConfigurationType", Required: false},
}

var fields_set_user_settings = []leanruntime.Field{
	{Name: "AccessToken", Flag: "access-token", Type: "*string", Required: true},
	{Name: "MFAOptions", Flag: "mfa-options", Type: "[]types.MFAOptionType", Required: true},
}

var fields_sign_up = []leanruntime.Field{
	{Name: "AnalyticsMetadata", Flag: "analytics-metadata", Type: "*types.AnalyticsMetadataType", Required: false},
	{Name: "ClientId", Flag: "client-id", Type: "*string", Required: true},
	{Name: "ClientMetadata", Flag: "client-metadata", Type: "map[string]string", Required: false},
	{Name: "Password", Flag: "password", Type: "*string", Required: false},
	{Name: "SecretHash", Flag: "secret-hash", Type: "*string", Required: false},
	{Name: "UserAttributes", Flag: "user-attributes", Type: "[]types.AttributeType", Required: false},
	{Name: "UserContextData", Flag: "user-context-data", Type: "*types.UserContextDataType", Required: false},
	{Name: "Username", Flag: "username", Type: "*string", Required: true},
	{Name: "ValidationData", Flag: "validation-data", Type: "[]types.AttributeType", Required: false},
}

var fields_start_user_import_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_start_web_authn_registration = []leanruntime.Field{
	{Name: "AccessToken", Flag: "access-token", Type: "*string", Required: true},
}

var fields_stop_user_import_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_auth_event_feedback = []leanruntime.Field{
	{Name: "EventId", Flag: "event-id", Type: "*string", Required: true},
	{Name: "FeedbackToken", Flag: "feedback-token", Type: "*string", Required: true},
	{Name: "FeedbackValue", Flag: "feedback-value", Type: "types.FeedbackValueType", Required: true},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
	{Name: "Username", Flag: "username", Type: "*string", Required: true},
}

var fields_update_device_status = []leanruntime.Field{
	{Name: "AccessToken", Flag: "access-token", Type: "*string", Required: true},
	{Name: "DeviceKey", Flag: "device-key", Type: "*string", Required: true},
	{Name: "DeviceRememberedStatus", Flag: "device-remembered-status", Type: "types.DeviceRememberedStatusType", Required: false},
}

var fields_update_group = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: true},
	{Name: "Precedence", Flag: "precedence", Type: "*int32", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_update_identity_provider = []leanruntime.Field{
	{Name: "AttributeMapping", Flag: "attribute-mapping", Type: "map[string]string", Required: false},
	{Name: "IdpIdentifiers", Flag: "idp-identifiers", Type: "[]string", Required: false},
	{Name: "ProviderDetails", Flag: "provider-details", Type: "map[string]string", Required: false},
	{Name: "ProviderName", Flag: "provider-name", Type: "*string", Required: true},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_update_managed_login_branding = []leanruntime.Field{
	{Name: "Assets", Flag: "assets", Type: "[]types.AssetType", Required: false},
	{Name: "ManagedLoginBrandingId", Flag: "managed-login-branding-id", Type: "*string", Required: false},
	{Name: "Settings", Flag: "settings", Type: "document.Interface", Required: false},
	{Name: "UseCognitoProvidedValues", Flag: "use-cognito-provided-values", Type: "bool", Required: false},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: false},
}

var fields_update_resource_server = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Scopes", Flag: "scopes", Type: "[]types.ResourceServerScopeType", Required: false},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_update_terms = []leanruntime.Field{
	{Name: "Enforcement", Flag: "enforcement", Type: "types.TermsEnforcementType", Required: false},
	{Name: "Links", Flag: "links", Type: "map[string]string", Required: false},
	{Name: "TermsId", Flag: "terms-id", Type: "*string", Required: true},
	{Name: "TermsName", Flag: "terms-name", Type: "*string", Required: false},
	{Name: "TermsSource", Flag: "terms-source", Type: "types.TermsSourceType", Required: false},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_update_user_attributes = []leanruntime.Field{
	{Name: "AccessToken", Flag: "access-token", Type: "*string", Required: true},
	{Name: "ClientMetadata", Flag: "client-metadata", Type: "map[string]string", Required: false},
	{Name: "UserAttributes", Flag: "user-attributes", Type: "[]types.AttributeType", Required: true},
}

var fields_update_user_pool = []leanruntime.Field{
	{Name: "AccountRecoverySetting", Flag: "account-recovery-setting", Type: "*types.AccountRecoverySettingType", Required: false},
	{Name: "AdminCreateUserConfig", Flag: "admin-create-user-config", Type: "*types.AdminCreateUserConfigType", Required: false},
	{Name: "AutoVerifiedAttributes", Flag: "auto-verified-attributes", Type: "[]types.VerifiedAttributeType", Required: false},
	{Name: "DeletionProtection", Flag: "deletion-protection", Type: "types.DeletionProtectionType", Required: false},
	{Name: "DeviceConfiguration", Flag: "device-configuration", Type: "*types.DeviceConfigurationType", Required: false},
	{Name: "EmailConfiguration", Flag: "email-configuration", Type: "*types.EmailConfigurationType", Required: false},
	{Name: "EmailVerificationMessage", Flag: "email-verification-message", Type: "*string", Required: false},
	{Name: "EmailVerificationSubject", Flag: "email-verification-subject", Type: "*string", Required: false},
	{Name: "LambdaConfig", Flag: "lambda-config", Type: "*types.LambdaConfigType", Required: false},
	{Name: "MfaConfiguration", Flag: "mfa-configuration", Type: "types.UserPoolMfaType", Required: false},
	{Name: "Policies", Flag: "policies", Type: "*types.UserPoolPolicyType", Required: false},
	{Name: "PoolName", Flag: "pool-name", Type: "*string", Required: false},
	{Name: "SmsAuthenticationMessage", Flag: "sms-authentication-message", Type: "*string", Required: false},
	{Name: "SmsConfiguration", Flag: "sms-configuration", Type: "*types.SmsConfigurationType", Required: false},
	{Name: "SmsVerificationMessage", Flag: "sms-verification-message", Type: "*string", Required: false},
	{Name: "UserAttributeUpdateSettings", Flag: "user-attribute-update-settings", Type: "*types.UserAttributeUpdateSettingsType", Required: false},
	{Name: "UserPoolAddOns", Flag: "user-pool-add-ons", Type: "*types.UserPoolAddOnsType", Required: false},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
	{Name: "UserPoolTags", Flag: "user-pool-tags", Type: "map[string]string", Required: false},
	{Name: "UserPoolTier", Flag: "user-pool-tier", Type: "types.UserPoolTierType", Required: false},
	{Name: "VerificationMessageTemplate", Flag: "verification-message-template", Type: "*types.VerificationMessageTemplateType", Required: false},
}

var fields_update_user_pool_client = []leanruntime.Field{
	{Name: "AccessTokenValidity", Flag: "access-token-validity", Type: "*int32", Required: false},
	{Name: "AllowedOAuthFlows", Flag: "allowed-oauth-flows", Type: "[]types.OAuthFlowType", Required: false},
	{Name: "AllowedOAuthFlowsUserPoolClient", Flag: "allowed-oauth-flows-user-pool-client", Type: "bool", Required: false},
	{Name: "AllowedOAuthScopes", Flag: "allowed-oauth-scopes", Type: "[]string", Required: false},
	{Name: "AnalyticsConfiguration", Flag: "analytics-configuration", Type: "*types.AnalyticsConfigurationType", Required: false},
	{Name: "AuthSessionValidity", Flag: "auth-session-validity", Type: "*int32", Required: false},
	{Name: "CallbackURLs", Flag: "callback-urls", Type: "[]string", Required: false},
	{Name: "ClientId", Flag: "client-id", Type: "*string", Required: true},
	{Name: "ClientName", Flag: "client-name", Type: "*string", Required: false},
	{Name: "DefaultRedirectURI", Flag: "default-redirect-uri", Type: "*string", Required: false},
	{Name: "EnablePropagateAdditionalUserContextData", Flag: "enable-propagate-additional-user-context-data", Type: "*bool", Required: false},
	{Name: "EnableTokenRevocation", Flag: "enable-token-revocation", Type: "*bool", Required: false},
	{Name: "ExplicitAuthFlows", Flag: "explicit-auth-flows", Type: "[]types.ExplicitAuthFlowsType", Required: false},
	{Name: "IdTokenValidity", Flag: "id-token-validity", Type: "*int32", Required: false},
	{Name: "LogoutURLs", Flag: "logout-urls", Type: "[]string", Required: false},
	{Name: "PreventUserExistenceErrors", Flag: "prevent-user-existence-errors", Type: "types.PreventUserExistenceErrorTypes", Required: false},
	{Name: "ReadAttributes", Flag: "read-attributes", Type: "[]string", Required: false},
	{Name: "RefreshTokenRotation", Flag: "refresh-token-rotation", Type: "*types.RefreshTokenRotationType", Required: false},
	{Name: "RefreshTokenValidity", Flag: "refresh-token-validity", Type: "int32", Required: false},
	{Name: "SupportedIdentityProviders", Flag: "supported-identity-providers", Type: "[]string", Required: false},
	{Name: "TokenValidityUnits", Flag: "token-validity-units", Type: "*types.TokenValidityUnitsType", Required: false},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
	{Name: "WriteAttributes", Flag: "write-attributes", Type: "[]string", Required: false},
}

var fields_update_user_pool_domain = []leanruntime.Field{
	{Name: "CustomDomainConfig", Flag: "custom-domain-config", Type: "*types.CustomDomainConfigType", Required: false},
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "ManagedLoginVersion", Flag: "managed-login-version", Type: "*int32", Required: false},
	{Name: "UserPoolId", Flag: "user-pool-id", Type: "*string", Required: true},
}

var fields_verify_software_token = []leanruntime.Field{
	{Name: "AccessToken", Flag: "access-token", Type: "*string", Required: false},
	{Name: "FriendlyDeviceName", Flag: "friendly-device-name", Type: "*string", Required: false},
	{Name: "Session", Flag: "session", Type: "*string", Required: false},
	{Name: "UserCode", Flag: "user-code", Type: "*string", Required: true},
}

var fields_verify_user_attribute = []leanruntime.Field{
	{Name: "AccessToken", Flag: "access-token", Type: "*string", Required: true},
	{Name: "AttributeName", Flag: "attribute-name", Type: "*string", Required: true},
	{Name: "Code", Flag: "code", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"add-custom-attributes": {
			Name:   "add-custom-attributes",
			Fields: fields_add_custom_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddCustomAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_custom_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddCustomAttributes(ctx, input)
			},
		},
		"add-user-pool-client-secret": {
			Name:   "add-user-pool-client-secret",
			Fields: fields_add_user_pool_client_secret,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddUserPoolClientSecretInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_user_pool_client_secret, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddUserPoolClientSecret(ctx, input)
			},
		},
		"admin-add-user-to-group": {
			Name:   "admin-add-user-to-group",
			Fields: fields_admin_add_user_to_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AdminAddUserToGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_admin_add_user_to_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AdminAddUserToGroup(ctx, input)
			},
		},
		"admin-confirm-sign-up": {
			Name:   "admin-confirm-sign-up",
			Fields: fields_admin_confirm_sign_up,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AdminConfirmSignUpInput{}
				if _, err := leanruntime.ApplyInput(input, fields_admin_confirm_sign_up, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AdminConfirmSignUp(ctx, input)
			},
		},
		"admin-create-user": {
			Name:   "admin-create-user",
			Fields: fields_admin_create_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AdminCreateUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_admin_create_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AdminCreateUser(ctx, input)
			},
		},
		"admin-delete-user": {
			Name:   "admin-delete-user",
			Fields: fields_admin_delete_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AdminDeleteUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_admin_delete_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AdminDeleteUser(ctx, input)
			},
		},
		"admin-delete-user-attributes": {
			Name:   "admin-delete-user-attributes",
			Fields: fields_admin_delete_user_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AdminDeleteUserAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_admin_delete_user_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AdminDeleteUserAttributes(ctx, input)
			},
		},
		"admin-disable-provider-for-user": {
			Name:   "admin-disable-provider-for-user",
			Fields: fields_admin_disable_provider_for_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AdminDisableProviderForUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_admin_disable_provider_for_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AdminDisableProviderForUser(ctx, input)
			},
		},
		"admin-disable-user": {
			Name:   "admin-disable-user",
			Fields: fields_admin_disable_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AdminDisableUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_admin_disable_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AdminDisableUser(ctx, input)
			},
		},
		"admin-enable-user": {
			Name:   "admin-enable-user",
			Fields: fields_admin_enable_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AdminEnableUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_admin_enable_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AdminEnableUser(ctx, input)
			},
		},
		"admin-forget-device": {
			Name:   "admin-forget-device",
			Fields: fields_admin_forget_device,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AdminForgetDeviceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_admin_forget_device, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AdminForgetDevice(ctx, input)
			},
		},
		"admin-get-device": {
			Name:   "admin-get-device",
			Fields: fields_admin_get_device,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AdminGetDeviceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_admin_get_device, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AdminGetDevice(ctx, input)
			},
		},
		"admin-get-user": {
			Name:   "admin-get-user",
			Fields: fields_admin_get_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AdminGetUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_admin_get_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AdminGetUser(ctx, input)
			},
		},
		"admin-initiate-auth": {
			Name:   "admin-initiate-auth",
			Fields: fields_admin_initiate_auth,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AdminInitiateAuthInput{}
				if _, err := leanruntime.ApplyInput(input, fields_admin_initiate_auth, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AdminInitiateAuth(ctx, input)
			},
		},
		"admin-link-provider-for-user": {
			Name:   "admin-link-provider-for-user",
			Fields: fields_admin_link_provider_for_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AdminLinkProviderForUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_admin_link_provider_for_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AdminLinkProviderForUser(ctx, input)
			},
		},
		"admin-list-devices": {
			Name:   "admin-list-devices",
			Fields: fields_admin_list_devices,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AdminListDevicesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_admin_list_devices, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AdminListDevices(ctx, input)
			},
		},
		"admin-list-groups-for-user": {
			Name:   "admin-list-groups-for-user",
			Fields: fields_admin_list_groups_for_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AdminListGroupsForUserInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_admin_list_groups_for_user, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.AdminListGroupsForUser(ctx, input)
				}
				var results []*svc.AdminListGroupsForUserOutput
				p := svc.NewAdminListGroupsForUserPaginator(client, input)
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
		"admin-list-user-auth-events": {
			Name:   "admin-list-user-auth-events",
			Fields: fields_admin_list_user_auth_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AdminListUserAuthEventsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_admin_list_user_auth_events, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.AdminListUserAuthEvents(ctx, input)
				}
				var results []*svc.AdminListUserAuthEventsOutput
				p := svc.NewAdminListUserAuthEventsPaginator(client, input)
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
		"admin-remove-user-from-group": {
			Name:   "admin-remove-user-from-group",
			Fields: fields_admin_remove_user_from_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AdminRemoveUserFromGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_admin_remove_user_from_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AdminRemoveUserFromGroup(ctx, input)
			},
		},
		"admin-reset-user-password": {
			Name:   "admin-reset-user-password",
			Fields: fields_admin_reset_user_password,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AdminResetUserPasswordInput{}
				if _, err := leanruntime.ApplyInput(input, fields_admin_reset_user_password, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AdminResetUserPassword(ctx, input)
			},
		},
		"admin-respond-to-auth-challenge": {
			Name:   "admin-respond-to-auth-challenge",
			Fields: fields_admin_respond_to_auth_challenge,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AdminRespondToAuthChallengeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_admin_respond_to_auth_challenge, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AdminRespondToAuthChallenge(ctx, input)
			},
		},
		"admin-set-user-mfa-preference": {
			Name:   "admin-set-user-mfa-preference",
			Fields: fields_admin_set_user_mfa_preference,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AdminSetUserMFAPreferenceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_admin_set_user_mfa_preference, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AdminSetUserMFAPreference(ctx, input)
			},
		},
		"admin-set-user-password": {
			Name:   "admin-set-user-password",
			Fields: fields_admin_set_user_password,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AdminSetUserPasswordInput{}
				if _, err := leanruntime.ApplyInput(input, fields_admin_set_user_password, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AdminSetUserPassword(ctx, input)
			},
		},
		"admin-set-user-settings": {
			Name:   "admin-set-user-settings",
			Fields: fields_admin_set_user_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AdminSetUserSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_admin_set_user_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AdminSetUserSettings(ctx, input)
			},
		},
		"admin-update-auth-event-feedback": {
			Name:   "admin-update-auth-event-feedback",
			Fields: fields_admin_update_auth_event_feedback,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AdminUpdateAuthEventFeedbackInput{}
				if _, err := leanruntime.ApplyInput(input, fields_admin_update_auth_event_feedback, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AdminUpdateAuthEventFeedback(ctx, input)
			},
		},
		"admin-update-device-status": {
			Name:   "admin-update-device-status",
			Fields: fields_admin_update_device_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AdminUpdateDeviceStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_admin_update_device_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AdminUpdateDeviceStatus(ctx, input)
			},
		},
		"admin-update-user-attributes": {
			Name:   "admin-update-user-attributes",
			Fields: fields_admin_update_user_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AdminUpdateUserAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_admin_update_user_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AdminUpdateUserAttributes(ctx, input)
			},
		},
		"admin-user-global-sign-out": {
			Name:   "admin-user-global-sign-out",
			Fields: fields_admin_user_global_sign_out,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AdminUserGlobalSignOutInput{}
				if _, err := leanruntime.ApplyInput(input, fields_admin_user_global_sign_out, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AdminUserGlobalSignOut(ctx, input)
			},
		},
		"associate-software-token": {
			Name:   "associate-software-token",
			Fields: fields_associate_software_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateSoftwareTokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_software_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateSoftwareToken(ctx, input)
			},
		},
		"change-password": {
			Name:   "change-password",
			Fields: fields_change_password,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ChangePasswordInput{}
				if _, err := leanruntime.ApplyInput(input, fields_change_password, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ChangePassword(ctx, input)
			},
		},
		"complete-web-authn-registration": {
			Name:   "complete-web-authn-registration",
			Fields: fields_complete_web_authn_registration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CompleteWebAuthnRegistrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_complete_web_authn_registration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CompleteWebAuthnRegistration(ctx, input)
			},
		},
		"confirm-device": {
			Name:   "confirm-device",
			Fields: fields_confirm_device,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ConfirmDeviceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_confirm_device, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ConfirmDevice(ctx, input)
			},
		},
		"confirm-forgot-password": {
			Name:   "confirm-forgot-password",
			Fields: fields_confirm_forgot_password,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ConfirmForgotPasswordInput{}
				if _, err := leanruntime.ApplyInput(input, fields_confirm_forgot_password, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ConfirmForgotPassword(ctx, input)
			},
		},
		"confirm-sign-up": {
			Name:   "confirm-sign-up",
			Fields: fields_confirm_sign_up,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ConfirmSignUpInput{}
				if _, err := leanruntime.ApplyInput(input, fields_confirm_sign_up, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ConfirmSignUp(ctx, input)
			},
		},
		"create-group": {
			Name:   "create-group",
			Fields: fields_create_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateGroup(ctx, input)
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
		"create-managed-login-branding": {
			Name:   "create-managed-login-branding",
			Fields: fields_create_managed_login_branding,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateManagedLoginBrandingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_managed_login_branding, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateManagedLoginBranding(ctx, input)
			},
		},
		"create-resource-server": {
			Name:   "create-resource-server",
			Fields: fields_create_resource_server,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateResourceServerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_resource_server, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateResourceServer(ctx, input)
			},
		},
		"create-terms": {
			Name:   "create-terms",
			Fields: fields_create_terms,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTermsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_terms, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTerms(ctx, input)
			},
		},
		"create-user-import-job": {
			Name:   "create-user-import-job",
			Fields: fields_create_user_import_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateUserImportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_user_import_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateUserImportJob(ctx, input)
			},
		},
		"create-user-pool": {
			Name:   "create-user-pool",
			Fields: fields_create_user_pool,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateUserPoolInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_user_pool, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateUserPool(ctx, input)
			},
		},
		"create-user-pool-client": {
			Name:   "create-user-pool-client",
			Fields: fields_create_user_pool_client,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateUserPoolClientInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_user_pool_client, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateUserPoolClient(ctx, input)
			},
		},
		"create-user-pool-domain": {
			Name:   "create-user-pool-domain",
			Fields: fields_create_user_pool_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateUserPoolDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_user_pool_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateUserPoolDomain(ctx, input)
			},
		},
		"delete-group": {
			Name:   "delete-group",
			Fields: fields_delete_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteGroup(ctx, input)
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
		"delete-managed-login-branding": {
			Name:   "delete-managed-login-branding",
			Fields: fields_delete_managed_login_branding,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteManagedLoginBrandingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_managed_login_branding, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteManagedLoginBranding(ctx, input)
			},
		},
		"delete-resource-server": {
			Name:   "delete-resource-server",
			Fields: fields_delete_resource_server,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteResourceServerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_resource_server, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteResourceServer(ctx, input)
			},
		},
		"delete-terms": {
			Name:   "delete-terms",
			Fields: fields_delete_terms,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTermsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_terms, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTerms(ctx, input)
			},
		},
		"delete-user": {
			Name:   "delete-user",
			Fields: fields_delete_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteUser(ctx, input)
			},
		},
		"delete-user-attributes": {
			Name:   "delete-user-attributes",
			Fields: fields_delete_user_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteUserAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_user_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteUserAttributes(ctx, input)
			},
		},
		"delete-user-pool": {
			Name:   "delete-user-pool",
			Fields: fields_delete_user_pool,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteUserPoolInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_user_pool, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteUserPool(ctx, input)
			},
		},
		"delete-user-pool-client": {
			Name:   "delete-user-pool-client",
			Fields: fields_delete_user_pool_client,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteUserPoolClientInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_user_pool_client, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteUserPoolClient(ctx, input)
			},
		},
		"delete-user-pool-client-secret": {
			Name:   "delete-user-pool-client-secret",
			Fields: fields_delete_user_pool_client_secret,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteUserPoolClientSecretInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_user_pool_client_secret, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteUserPoolClientSecret(ctx, input)
			},
		},
		"delete-user-pool-domain": {
			Name:   "delete-user-pool-domain",
			Fields: fields_delete_user_pool_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteUserPoolDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_user_pool_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteUserPoolDomain(ctx, input)
			},
		},
		"delete-web-authn-credential": {
			Name:   "delete-web-authn-credential",
			Fields: fields_delete_web_authn_credential,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWebAuthnCredentialInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_web_authn_credential, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWebAuthnCredential(ctx, input)
			},
		},
		"describe-identity-provider": {
			Name:   "describe-identity-provider",
			Fields: fields_describe_identity_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeIdentityProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_identity_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeIdentityProvider(ctx, input)
			},
		},
		"describe-managed-login-branding": {
			Name:   "describe-managed-login-branding",
			Fields: fields_describe_managed_login_branding,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeManagedLoginBrandingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_managed_login_branding, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeManagedLoginBranding(ctx, input)
			},
		},
		"describe-managed-login-branding-by-client": {
			Name:   "describe-managed-login-branding-by-client",
			Fields: fields_describe_managed_login_branding_by_client,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeManagedLoginBrandingByClientInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_managed_login_branding_by_client, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeManagedLoginBrandingByClient(ctx, input)
			},
		},
		"describe-resource-server": {
			Name:   "describe-resource-server",
			Fields: fields_describe_resource_server,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeResourceServerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_resource_server, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeResourceServer(ctx, input)
			},
		},
		"describe-risk-configuration": {
			Name:   "describe-risk-configuration",
			Fields: fields_describe_risk_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRiskConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_risk_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeRiskConfiguration(ctx, input)
			},
		},
		"describe-terms": {
			Name:   "describe-terms",
			Fields: fields_describe_terms,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTermsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_terms, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTerms(ctx, input)
			},
		},
		"describe-user-import-job": {
			Name:   "describe-user-import-job",
			Fields: fields_describe_user_import_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeUserImportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_user_import_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeUserImportJob(ctx, input)
			},
		},
		"describe-user-pool": {
			Name:   "describe-user-pool",
			Fields: fields_describe_user_pool,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeUserPoolInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_user_pool, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeUserPool(ctx, input)
			},
		},
		"describe-user-pool-client": {
			Name:   "describe-user-pool-client",
			Fields: fields_describe_user_pool_client,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeUserPoolClientInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_user_pool_client, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeUserPoolClient(ctx, input)
			},
		},
		"describe-user-pool-domain": {
			Name:   "describe-user-pool-domain",
			Fields: fields_describe_user_pool_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeUserPoolDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_user_pool_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeUserPoolDomain(ctx, input)
			},
		},
		"forget-device": {
			Name:   "forget-device",
			Fields: fields_forget_device,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ForgetDeviceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_forget_device, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ForgetDevice(ctx, input)
			},
		},
		"forgot-password": {
			Name:   "forgot-password",
			Fields: fields_forgot_password,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ForgotPasswordInput{}
				if _, err := leanruntime.ApplyInput(input, fields_forgot_password, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ForgotPassword(ctx, input)
			},
		},
		"get-csv-header": {
			Name:   "get-csv-header",
			Fields: fields_get_csv_header,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCSVHeaderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_csv_header, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCSVHeader(ctx, input)
			},
		},
		"get-device": {
			Name:   "get-device",
			Fields: fields_get_device,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDeviceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_device, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDevice(ctx, input)
			},
		},
		"get-group": {
			Name:   "get-group",
			Fields: fields_get_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetGroup(ctx, input)
			},
		},
		"get-identity-provider-by-identifier": {
			Name:   "get-identity-provider-by-identifier",
			Fields: fields_get_identity_provider_by_identifier,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIdentityProviderByIdentifierInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_identity_provider_by_identifier, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIdentityProviderByIdentifier(ctx, input)
			},
		},
		"get-log-delivery-configuration": {
			Name:   "get-log-delivery-configuration",
			Fields: fields_get_log_delivery_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLogDeliveryConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_log_delivery_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLogDeliveryConfiguration(ctx, input)
			},
		},
		"get-signing-certificate": {
			Name:   "get-signing-certificate",
			Fields: fields_get_signing_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSigningCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_signing_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSigningCertificate(ctx, input)
			},
		},
		"get-tokens-from-refresh-token": {
			Name:   "get-tokens-from-refresh-token",
			Fields: fields_get_tokens_from_refresh_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTokensFromRefreshTokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_tokens_from_refresh_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTokensFromRefreshToken(ctx, input)
			},
		},
		"get-ui-customization": {
			Name:   "get-ui-customization",
			Fields: fields_get_ui_customization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetUICustomizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_ui_customization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetUICustomization(ctx, input)
			},
		},
		"get-user": {
			Name:   "get-user",
			Fields: fields_get_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetUser(ctx, input)
			},
		},
		"get-user-attribute-verification-code": {
			Name:   "get-user-attribute-verification-code",
			Fields: fields_get_user_attribute_verification_code,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetUserAttributeVerificationCodeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_user_attribute_verification_code, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetUserAttributeVerificationCode(ctx, input)
			},
		},
		"get-user-auth-factors": {
			Name:   "get-user-auth-factors",
			Fields: fields_get_user_auth_factors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetUserAuthFactorsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_user_auth_factors, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetUserAuthFactors(ctx, input)
			},
		},
		"get-user-pool-mfa-config": {
			Name:   "get-user-pool-mfa-config",
			Fields: fields_get_user_pool_mfa_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetUserPoolMfaConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_user_pool_mfa_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetUserPoolMfaConfig(ctx, input)
			},
		},
		"global-sign-out": {
			Name:   "global-sign-out",
			Fields: fields_global_sign_out,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GlobalSignOutInput{}
				if _, err := leanruntime.ApplyInput(input, fields_global_sign_out, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GlobalSignOut(ctx, input)
			},
		},
		"initiate-auth": {
			Name:   "initiate-auth",
			Fields: fields_initiate_auth,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.InitiateAuthInput{}
				if _, err := leanruntime.ApplyInput(input, fields_initiate_auth, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.InitiateAuth(ctx, input)
			},
		},
		"list-devices": {
			Name:   "list-devices",
			Fields: fields_list_devices,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDevicesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_devices, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListDevices(ctx, input)
			},
		},
		"list-groups": {
			Name:   "list-groups",
			Fields: fields_list_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListGroups(ctx, input)
				}
				var results []*svc.ListGroupsOutput
				p := svc.NewListGroupsPaginator(client, input)
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
		"list-resource-servers": {
			Name:   "list-resource-servers",
			Fields: fields_list_resource_servers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResourceServersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_resource_servers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResourceServers(ctx, input)
				}
				var results []*svc.ListResourceServersOutput
				p := svc.NewListResourceServersPaginator(client, input)
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
		"list-terms": {
			Name:   "list-terms",
			Fields: fields_list_terms,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTermsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_terms, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTerms(ctx, input)
			},
		},
		"list-user-import-jobs": {
			Name:   "list-user-import-jobs",
			Fields: fields_list_user_import_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListUserImportJobsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_user_import_jobs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListUserImportJobs(ctx, input)
			},
		},
		"list-user-pool-client-secrets": {
			Name:   "list-user-pool-client-secrets",
			Fields: fields_list_user_pool_client_secrets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListUserPoolClientSecretsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_user_pool_client_secrets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListUserPoolClientSecrets(ctx, input)
			},
		},
		"list-user-pool-clients": {
			Name:   "list-user-pool-clients",
			Fields: fields_list_user_pool_clients,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListUserPoolClientsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_user_pool_clients, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListUserPoolClients(ctx, input)
				}
				var results []*svc.ListUserPoolClientsOutput
				p := svc.NewListUserPoolClientsPaginator(client, input)
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
		"list-user-pools": {
			Name:   "list-user-pools",
			Fields: fields_list_user_pools,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListUserPoolsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_user_pools, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListUserPools(ctx, input)
				}
				var results []*svc.ListUserPoolsOutput
				p := svc.NewListUserPoolsPaginator(client, input)
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
		"list-users": {
			Name:   "list-users",
			Fields: fields_list_users,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListUsersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_users, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListUsers(ctx, input)
				}
				var results []*svc.ListUsersOutput
				p := svc.NewListUsersPaginator(client, input)
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
		"list-users-in-group": {
			Name:   "list-users-in-group",
			Fields: fields_list_users_in_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListUsersInGroupInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_users_in_group, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListUsersInGroup(ctx, input)
				}
				var results []*svc.ListUsersInGroupOutput
				p := svc.NewListUsersInGroupPaginator(client, input)
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
		"list-web-authn-credentials": {
			Name:   "list-web-authn-credentials",
			Fields: fields_list_web_authn_credentials,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWebAuthnCredentialsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_web_authn_credentials, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListWebAuthnCredentials(ctx, input)
			},
		},
		"resend-confirmation-code": {
			Name:   "resend-confirmation-code",
			Fields: fields_resend_confirmation_code,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResendConfirmationCodeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_resend_confirmation_code, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResendConfirmationCode(ctx, input)
			},
		},
		"respond-to-auth-challenge": {
			Name:   "respond-to-auth-challenge",
			Fields: fields_respond_to_auth_challenge,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RespondToAuthChallengeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_respond_to_auth_challenge, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RespondToAuthChallenge(ctx, input)
			},
		},
		"revoke-token": {
			Name:   "revoke-token",
			Fields: fields_revoke_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RevokeTokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_revoke_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RevokeToken(ctx, input)
			},
		},
		"set-log-delivery-configuration": {
			Name:   "set-log-delivery-configuration",
			Fields: fields_set_log_delivery_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetLogDeliveryConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_log_delivery_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetLogDeliveryConfiguration(ctx, input)
			},
		},
		"set-risk-configuration": {
			Name:   "set-risk-configuration",
			Fields: fields_set_risk_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetRiskConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_risk_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetRiskConfiguration(ctx, input)
			},
		},
		"set-ui-customization": {
			Name:   "set-ui-customization",
			Fields: fields_set_ui_customization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetUICustomizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_ui_customization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetUICustomization(ctx, input)
			},
		},
		"set-user-mfa-preference": {
			Name:   "set-user-mfa-preference",
			Fields: fields_set_user_mfa_preference,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetUserMFAPreferenceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_user_mfa_preference, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetUserMFAPreference(ctx, input)
			},
		},
		"set-user-pool-mfa-config": {
			Name:   "set-user-pool-mfa-config",
			Fields: fields_set_user_pool_mfa_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetUserPoolMfaConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_user_pool_mfa_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetUserPoolMfaConfig(ctx, input)
			},
		},
		"set-user-settings": {
			Name:   "set-user-settings",
			Fields: fields_set_user_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetUserSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_user_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetUserSettings(ctx, input)
			},
		},
		"sign-up": {
			Name:   "sign-up",
			Fields: fields_sign_up,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SignUpInput{}
				if _, err := leanruntime.ApplyInput(input, fields_sign_up, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SignUp(ctx, input)
			},
		},
		"start-user-import-job": {
			Name:   "start-user-import-job",
			Fields: fields_start_user_import_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartUserImportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_user_import_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartUserImportJob(ctx, input)
			},
		},
		"start-web-authn-registration": {
			Name:   "start-web-authn-registration",
			Fields: fields_start_web_authn_registration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartWebAuthnRegistrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_web_authn_registration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartWebAuthnRegistration(ctx, input)
			},
		},
		"stop-user-import-job": {
			Name:   "stop-user-import-job",
			Fields: fields_stop_user_import_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopUserImportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_user_import_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopUserImportJob(ctx, input)
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
		"update-auth-event-feedback": {
			Name:   "update-auth-event-feedback",
			Fields: fields_update_auth_event_feedback,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAuthEventFeedbackInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_auth_event_feedback, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAuthEventFeedback(ctx, input)
			},
		},
		"update-device-status": {
			Name:   "update-device-status",
			Fields: fields_update_device_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDeviceStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_device_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDeviceStatus(ctx, input)
			},
		},
		"update-group": {
			Name:   "update-group",
			Fields: fields_update_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateGroup(ctx, input)
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
		"update-managed-login-branding": {
			Name:   "update-managed-login-branding",
			Fields: fields_update_managed_login_branding,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateManagedLoginBrandingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_managed_login_branding, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateManagedLoginBranding(ctx, input)
			},
		},
		"update-resource-server": {
			Name:   "update-resource-server",
			Fields: fields_update_resource_server,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateResourceServerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_resource_server, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateResourceServer(ctx, input)
			},
		},
		"update-terms": {
			Name:   "update-terms",
			Fields: fields_update_terms,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTermsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_terms, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTerms(ctx, input)
			},
		},
		"update-user-attributes": {
			Name:   "update-user-attributes",
			Fields: fields_update_user_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateUserAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_user_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateUserAttributes(ctx, input)
			},
		},
		"update-user-pool": {
			Name:   "update-user-pool",
			Fields: fields_update_user_pool,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateUserPoolInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_user_pool, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateUserPool(ctx, input)
			},
		},
		"update-user-pool-client": {
			Name:   "update-user-pool-client",
			Fields: fields_update_user_pool_client,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateUserPoolClientInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_user_pool_client, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateUserPoolClient(ctx, input)
			},
		},
		"update-user-pool-domain": {
			Name:   "update-user-pool-domain",
			Fields: fields_update_user_pool_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateUserPoolDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_user_pool_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateUserPoolDomain(ctx, input)
			},
		},
		"verify-software-token": {
			Name:   "verify-software-token",
			Fields: fields_verify_software_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.VerifySoftwareTokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_verify_software_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.VerifySoftwareToken(ctx, input)
			},
		},
		"verify-user-attribute": {
			Name:   "verify-user-attribute",
			Fields: fields_verify_user_attribute,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.VerifyUserAttributeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_verify_user_attribute, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.VerifyUserAttribute(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("cognitoidentityprovider", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
