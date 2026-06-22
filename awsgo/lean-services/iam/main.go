package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/iam"
)

var fields_accept_delegation_request = []leanruntime.Field{
	{Name: "DelegationRequestId", Flag: "delegation-request-id", Type: "*string", Required: true},
}

var fields_add_client_idto_open_id_connect_provider = []leanruntime.Field{
	{Name: "ClientID", Flag: "client-id", Type: "*string", Required: true},
	{Name: "OpenIDConnectProviderArn", Flag: "open-id-connect-provider-arn", Type: "*string", Required: true},
}

var fields_add_role_to_instance_profile = []leanruntime.Field{
	{Name: "InstanceProfileName", Flag: "instance-profile-name", Type: "*string", Required: true},
	{Name: "RoleName", Flag: "role-name", Type: "*string", Required: true},
}

var fields_add_user_to_group = []leanruntime.Field{
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: true},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_associate_delegation_request = []leanruntime.Field{
	{Name: "DelegationRequestId", Flag: "delegation-request-id", Type: "*string", Required: true},
}

var fields_attach_group_policy = []leanruntime.Field{
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: true},
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: true},
}

var fields_attach_role_policy = []leanruntime.Field{
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: true},
	{Name: "RoleName", Flag: "role-name", Type: "*string", Required: true},
}

var fields_attach_user_policy = []leanruntime.Field{
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: true},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_change_password = []leanruntime.Field{
	{Name: "NewPassword", Flag: "new-password", Type: "*string", Required: true},
	{Name: "OldPassword", Flag: "old-password", Type: "*string", Required: true},
}

var fields_create_access_key = []leanruntime.Field{
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: false},
}

var fields_create_account_alias = []leanruntime.Field{
	{Name: "AccountAlias", Flag: "account-alias", Type: "*string", Required: true},
}

var fields_create_delegation_request = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "NotificationChannel", Flag: "notification-channel", Type: "*string", Required: true},
	{Name: "OnlySendByOwner", Flag: "only-send-by-owner", Type: "bool", Required: false},
	{Name: "OwnerAccountId", Flag: "owner-account-id", Type: "*string", Required: false},
	{Name: "Permissions", Flag: "permissions", Type: "*types.DelegationPermission", Required: true},
	{Name: "RedirectUrl", Flag: "redirect-url", Type: "*string", Required: false},
	{Name: "RequestMessage", Flag: "request-message", Type: "*string", Required: false},
	{Name: "RequestorWorkflowId", Flag: "requestor-workflow-id", Type: "*string", Required: true},
	{Name: "SessionDuration", Flag: "session-duration", Type: "*int32", Required: true},
}

var fields_create_group = []leanruntime.Field{
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: true},
	{Name: "Path", Flag: "path", Type: "*string", Required: false},
}

var fields_create_instance_profile = []leanruntime.Field{
	{Name: "InstanceProfileName", Flag: "instance-profile-name", Type: "*string", Required: true},
	{Name: "Path", Flag: "path", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_login_profile = []leanruntime.Field{
	{Name: "Password", Flag: "password", Type: "*string", Required: false},
	{Name: "PasswordResetRequired", Flag: "password-reset-required", Type: "bool", Required: false},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: false},
}

var fields_create_open_id_connect_provider = []leanruntime.Field{
	{Name: "ClientIDList", Flag: "client-id-list", Type: "[]string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "ThumbprintList", Flag: "thumbprint-list", Type: "[]string", Required: false},
	{Name: "Url", Flag: "url", Type: "*string", Required: true},
}

var fields_create_policy = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Path", Flag: "path", Type: "*string", Required: false},
	{Name: "PolicyDocument", Flag: "policy-document", Type: "*string", Required: true},
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_policy_version = []leanruntime.Field{
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: true},
	{Name: "PolicyDocument", Flag: "policy-document", Type: "*string", Required: true},
	{Name: "SetAsDefault", Flag: "set-as-default", Type: "bool", Required: false},
}

var fields_create_role = []leanruntime.Field{
	{Name: "AssumeRolePolicyDocument", Flag: "assume-role-policy-document", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "MaxSessionDuration", Flag: "max-session-duration", Type: "*int32", Required: false},
	{Name: "Path", Flag: "path", Type: "*string", Required: false},
	{Name: "PermissionsBoundary", Flag: "permissions-boundary", Type: "*string", Required: false},
	{Name: "RoleName", Flag: "role-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_saml_provider = []leanruntime.Field{
	{Name: "AddPrivateKey", Flag: "add-private-key", Type: "*string", Required: false},
	{Name: "AssertionEncryptionMode", Flag: "assertion-encryption-mode", Type: "types.AssertionEncryptionModeType", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "SAMLMetadataDocument", Flag: "saml-metadata-document", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_service_linked_role = []leanruntime.Field{
	{Name: "AWSServiceName", Flag: "aws-service-name", Type: "*string", Required: true},
	{Name: "CustomSuffix", Flag: "custom-suffix", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
}

var fields_create_service_specific_credential = []leanruntime.Field{
	{Name: "CredentialAgeDays", Flag: "credential-age-days", Type: "*int32", Required: false},
	{Name: "ServiceName", Flag: "service-name", Type: "*string", Required: true},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_create_user = []leanruntime.Field{
	{Name: "Path", Flag: "path", Type: "*string", Required: false},
	{Name: "PermissionsBoundary", Flag: "permissions-boundary", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_create_virtual_mfa_device = []leanruntime.Field{
	{Name: "Path", Flag: "path", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VirtualMFADeviceName", Flag: "virtual-mfa-device-name", Type: "*string", Required: true},
}

var fields_deactivate_mfa_device = []leanruntime.Field{
	{Name: "SerialNumber", Flag: "serial-number", Type: "*string", Required: true},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: false},
}

var fields_delete_access_key = []leanruntime.Field{
	{Name: "AccessKeyId", Flag: "access-key-id", Type: "*string", Required: true},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: false},
}

var fields_delete_account_alias = []leanruntime.Field{
	{Name: "AccountAlias", Flag: "account-alias", Type: "*string", Required: true},
}

var fields_delete_account_password_policy = []leanruntime.Field{}

var fields_delete_group = []leanruntime.Field{
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: true},
}

var fields_delete_group_policy = []leanruntime.Field{
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: true},
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
}

var fields_delete_instance_profile = []leanruntime.Field{
	{Name: "InstanceProfileName", Flag: "instance-profile-name", Type: "*string", Required: true},
}

var fields_delete_login_profile = []leanruntime.Field{
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: false},
}

var fields_delete_open_id_connect_provider = []leanruntime.Field{
	{Name: "OpenIDConnectProviderArn", Flag: "open-id-connect-provider-arn", Type: "*string", Required: true},
}

var fields_delete_policy = []leanruntime.Field{
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: true},
}

var fields_delete_policy_version = []leanruntime.Field{
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: true},
	{Name: "VersionId", Flag: "version-id", Type: "*string", Required: true},
}

var fields_delete_role = []leanruntime.Field{
	{Name: "RoleName", Flag: "role-name", Type: "*string", Required: true},
}

var fields_delete_role_permissions_boundary = []leanruntime.Field{
	{Name: "RoleName", Flag: "role-name", Type: "*string", Required: true},
}

var fields_delete_role_policy = []leanruntime.Field{
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
	{Name: "RoleName", Flag: "role-name", Type: "*string", Required: true},
}

var fields_delete_saml_provider = []leanruntime.Field{
	{Name: "SAMLProviderArn", Flag: "saml-provider-arn", Type: "*string", Required: true},
}

var fields_delete_server_certificate = []leanruntime.Field{
	{Name: "ServerCertificateName", Flag: "server-certificate-name", Type: "*string", Required: true},
}

var fields_delete_service_linked_role = []leanruntime.Field{
	{Name: "RoleName", Flag: "role-name", Type: "*string", Required: true},
}

var fields_delete_service_specific_credential = []leanruntime.Field{
	{Name: "ServiceSpecificCredentialId", Flag: "service-specific-credential-id", Type: "*string", Required: true},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: false},
}

var fields_delete_signing_certificate = []leanruntime.Field{
	{Name: "CertificateId", Flag: "certificate-id", Type: "*string", Required: true},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: false},
}

var fields_delete_ssh_public_key = []leanruntime.Field{
	{Name: "SSHPublicKeyId", Flag: "ssh-public-key-id", Type: "*string", Required: true},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_delete_user = []leanruntime.Field{
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_delete_user_permissions_boundary = []leanruntime.Field{
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_delete_user_policy = []leanruntime.Field{
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_delete_virtual_mfa_device = []leanruntime.Field{
	{Name: "SerialNumber", Flag: "serial-number", Type: "*string", Required: true},
}

var fields_detach_group_policy = []leanruntime.Field{
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: true},
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: true},
}

var fields_detach_role_policy = []leanruntime.Field{
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: true},
	{Name: "RoleName", Flag: "role-name", Type: "*string", Required: true},
}

var fields_detach_user_policy = []leanruntime.Field{
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: true},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_disable_organizations_root_credentials_management = []leanruntime.Field{}

var fields_disable_organizations_root_sessions = []leanruntime.Field{}

var fields_disable_outbound_web_identity_federation = []leanruntime.Field{}

var fields_enable_mfa_device = []leanruntime.Field{
	{Name: "AuthenticationCode1", Flag: "authentication-code1", Type: "*string", Required: true},
	{Name: "AuthenticationCode2", Flag: "authentication-code2", Type: "*string", Required: true},
	{Name: "SerialNumber", Flag: "serial-number", Type: "*string", Required: true},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_enable_organizations_root_credentials_management = []leanruntime.Field{}

var fields_enable_organizations_root_sessions = []leanruntime.Field{}

var fields_enable_outbound_web_identity_federation = []leanruntime.Field{}

var fields_generate_credential_report = []leanruntime.Field{}

var fields_generate_organizations_access_report = []leanruntime.Field{
	{Name: "EntityPath", Flag: "entity-path", Type: "*string", Required: true},
	{Name: "OrganizationsPolicyId", Flag: "organizations-policy-id", Type: "*string", Required: false},
}

var fields_generate_service_last_accessed_details = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "Granularity", Flag: "granularity", Type: "types.AccessAdvisorUsageGranularityType", Required: false},
}

var fields_get_access_key_last_used = []leanruntime.Field{
	{Name: "AccessKeyId", Flag: "access-key-id", Type: "*string", Required: true},
}

var fields_get_account_authorization_details = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "[]types.EntityType", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_get_account_password_policy = []leanruntime.Field{}

var fields_get_account_summary = []leanruntime.Field{}

var fields_get_context_keys_for_custom_policy = []leanruntime.Field{
	{Name: "PolicyInputList", Flag: "policy-input-list", Type: "[]string", Required: true},
}

var fields_get_context_keys_for_principal_policy = []leanruntime.Field{
	{Name: "PolicyInputList", Flag: "policy-input-list", Type: "[]string", Required: false},
	{Name: "PolicySourceArn", Flag: "policy-source-arn", Type: "*string", Required: true},
}

var fields_get_credential_report = []leanruntime.Field{}

var fields_get_delegation_request = []leanruntime.Field{
	{Name: "DelegationPermissionCheck", Flag: "delegation-permission-check", Type: "bool", Required: false},
	{Name: "DelegationRequestId", Flag: "delegation-request-id", Type: "*string", Required: true},
}

var fields_get_group = []leanruntime.Field{
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: true},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_get_group_policy = []leanruntime.Field{
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: true},
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
}

var fields_get_human_readable_summary = []leanruntime.Field{
	{Name: "EntityArn", Flag: "entity-arn", Type: "*string", Required: true},
	{Name: "Locale", Flag: "locale", Type: "*string", Required: false},
}

var fields_get_instance_profile = []leanruntime.Field{
	{Name: "InstanceProfileName", Flag: "instance-profile-name", Type: "*string", Required: true},
}

var fields_get_login_profile = []leanruntime.Field{
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: false},
}

var fields_get_mfa_device = []leanruntime.Field{
	{Name: "SerialNumber", Flag: "serial-number", Type: "*string", Required: true},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: false},
}

var fields_get_open_id_connect_provider = []leanruntime.Field{
	{Name: "OpenIDConnectProviderArn", Flag: "open-id-connect-provider-arn", Type: "*string", Required: true},
}

var fields_get_organizations_access_report = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "SortKey", Flag: "sort-key", Type: "types.SortKeyType", Required: false},
}

var fields_get_outbound_web_identity_federation_info = []leanruntime.Field{}

var fields_get_policy = []leanruntime.Field{
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: true},
}

var fields_get_policy_version = []leanruntime.Field{
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: true},
	{Name: "VersionId", Flag: "version-id", Type: "*string", Required: true},
}

var fields_get_role = []leanruntime.Field{
	{Name: "RoleName", Flag: "role-name", Type: "*string", Required: true},
}

var fields_get_role_policy = []leanruntime.Field{
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
	{Name: "RoleName", Flag: "role-name", Type: "*string", Required: true},
}

var fields_get_saml_provider = []leanruntime.Field{
	{Name: "SAMLProviderArn", Flag: "saml-provider-arn", Type: "*string", Required: true},
}

var fields_get_server_certificate = []leanruntime.Field{
	{Name: "ServerCertificateName", Flag: "server-certificate-name", Type: "*string", Required: true},
}

var fields_get_service_last_accessed_details = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_get_service_last_accessed_details_with_entities = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "ServiceNamespace", Flag: "service-namespace", Type: "*string", Required: true},
}

var fields_get_service_linked_role_deletion_status = []leanruntime.Field{
	{Name: "DeletionTaskId", Flag: "deletion-task-id", Type: "*string", Required: true},
}

var fields_get_ssh_public_key = []leanruntime.Field{
	{Name: "Encoding", Flag: "encoding", Type: "types.EncodingType", Required: true},
	{Name: "SSHPublicKeyId", Flag: "ssh-public-key-id", Type: "*string", Required: true},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_get_user = []leanruntime.Field{
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: false},
}

var fields_get_user_policy = []leanruntime.Field{
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_list_access_keys = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: false},
}

var fields_list_account_aliases = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_list_attached_group_policies = []leanruntime.Field{
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: true},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "PathPrefix", Flag: "path-prefix", Type: "*string", Required: false},
}

var fields_list_attached_role_policies = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "PathPrefix", Flag: "path-prefix", Type: "*string", Required: false},
	{Name: "RoleName", Flag: "role-name", Type: "*string", Required: true},
}

var fields_list_attached_user_policies = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "PathPrefix", Flag: "path-prefix", Type: "*string", Required: false},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_list_delegation_requests = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "OwnerId", Flag: "owner-id", Type: "*string", Required: false},
}

var fields_list_entities_for_policy = []leanruntime.Field{
	{Name: "EntityFilter", Flag: "entity-filter", Type: "types.EntityType", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "PathPrefix", Flag: "path-prefix", Type: "*string", Required: false},
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: true},
	{Name: "PolicyUsageFilter", Flag: "policy-usage-filter", Type: "types.PolicyUsageType", Required: false},
}

var fields_list_group_policies = []leanruntime.Field{
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: true},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_list_groups = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "PathPrefix", Flag: "path-prefix", Type: "*string", Required: false},
}

var fields_list_groups_for_user = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_list_instance_profile_tags = []leanruntime.Field{
	{Name: "InstanceProfileName", Flag: "instance-profile-name", Type: "*string", Required: true},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_list_instance_profiles = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "PathPrefix", Flag: "path-prefix", Type: "*string", Required: false},
}

var fields_list_instance_profiles_for_role = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "RoleName", Flag: "role-name", Type: "*string", Required: true},
}

var fields_list_mfa_device_tags = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "SerialNumber", Flag: "serial-number", Type: "*string", Required: true},
}

var fields_list_mfa_devices = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: false},
}

var fields_list_open_id_connect_provider_tags = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "OpenIDConnectProviderArn", Flag: "open-id-connect-provider-arn", Type: "*string", Required: true},
}

var fields_list_open_id_connect_providers = []leanruntime.Field{}

var fields_list_organizations_features = []leanruntime.Field{}

var fields_list_policies = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "OnlyAttached", Flag: "only-attached", Type: "bool", Required: false},
	{Name: "PathPrefix", Flag: "path-prefix", Type: "*string", Required: false},
	{Name: "PolicyUsageFilter", Flag: "policy-usage-filter", Type: "types.PolicyUsageType", Required: false},
	{Name: "Scope", Flag: "scope", Type: "types.PolicyScopeType", Required: false},
}

var fields_list_policies_granting_service_access = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "ServiceNamespaces", Flag: "service-namespaces", Type: "[]string", Required: true},
}

var fields_list_policy_tags = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: true},
}

var fields_list_policy_versions = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: true},
}

var fields_list_role_policies = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "RoleName", Flag: "role-name", Type: "*string", Required: true},
}

var fields_list_role_tags = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "RoleName", Flag: "role-name", Type: "*string", Required: true},
}

var fields_list_roles = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "PathPrefix", Flag: "path-prefix", Type: "*string", Required: false},
}

var fields_list_saml_provider_tags = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "SAMLProviderArn", Flag: "saml-provider-arn", Type: "*string", Required: true},
}

var fields_list_saml_providers = []leanruntime.Field{}

var fields_list_server_certificate_tags = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "ServerCertificateName", Flag: "server-certificate-name", Type: "*string", Required: true},
}

var fields_list_server_certificates = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "PathPrefix", Flag: "path-prefix", Type: "*string", Required: false},
}

var fields_list_service_specific_credentials = []leanruntime.Field{
	{Name: "AllUsers", Flag: "all-users", Type: "*bool", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "ServiceName", Flag: "service-name", Type: "*string", Required: false},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: false},
}

var fields_list_signing_certificates = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: false},
}

var fields_list_ssh_public_keys = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: false},
}

var fields_list_user_policies = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_list_user_tags = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_list_users = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "PathPrefix", Flag: "path-prefix", Type: "*string", Required: false},
}

var fields_list_virtual_mfa_devices = []leanruntime.Field{
	{Name: "AssignmentStatus", Flag: "assignment-status", Type: "types.AssignmentStatusType", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_put_group_policy = []leanruntime.Field{
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: true},
	{Name: "PolicyDocument", Flag: "policy-document", Type: "*string", Required: true},
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
}

var fields_put_role_permissions_boundary = []leanruntime.Field{
	{Name: "PermissionsBoundary", Flag: "permissions-boundary", Type: "*string", Required: true},
	{Name: "RoleName", Flag: "role-name", Type: "*string", Required: true},
}

var fields_put_role_policy = []leanruntime.Field{
	{Name: "PolicyDocument", Flag: "policy-document", Type: "*string", Required: true},
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
	{Name: "RoleName", Flag: "role-name", Type: "*string", Required: true},
}

var fields_put_user_permissions_boundary = []leanruntime.Field{
	{Name: "PermissionsBoundary", Flag: "permissions-boundary", Type: "*string", Required: true},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_put_user_policy = []leanruntime.Field{
	{Name: "PolicyDocument", Flag: "policy-document", Type: "*string", Required: true},
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_reject_delegation_request = []leanruntime.Field{
	{Name: "DelegationRequestId", Flag: "delegation-request-id", Type: "*string", Required: true},
	{Name: "Notes", Flag: "notes", Type: "*string", Required: false},
}

var fields_remove_client_id_from_open_id_connect_provider = []leanruntime.Field{
	{Name: "ClientID", Flag: "client-id", Type: "*string", Required: true},
	{Name: "OpenIDConnectProviderArn", Flag: "open-id-connect-provider-arn", Type: "*string", Required: true},
}

var fields_remove_role_from_instance_profile = []leanruntime.Field{
	{Name: "InstanceProfileName", Flag: "instance-profile-name", Type: "*string", Required: true},
	{Name: "RoleName", Flag: "role-name", Type: "*string", Required: true},
}

var fields_remove_user_from_group = []leanruntime.Field{
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: true},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_reset_service_specific_credential = []leanruntime.Field{
	{Name: "ServiceSpecificCredentialId", Flag: "service-specific-credential-id", Type: "*string", Required: true},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: false},
}

var fields_resync_mfa_device = []leanruntime.Field{
	{Name: "AuthenticationCode1", Flag: "authentication-code1", Type: "*string", Required: true},
	{Name: "AuthenticationCode2", Flag: "authentication-code2", Type: "*string", Required: true},
	{Name: "SerialNumber", Flag: "serial-number", Type: "*string", Required: true},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_send_delegation_token = []leanruntime.Field{
	{Name: "DelegationRequestId", Flag: "delegation-request-id", Type: "*string", Required: true},
}

var fields_set_default_policy_version = []leanruntime.Field{
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: true},
	{Name: "VersionId", Flag: "version-id", Type: "*string", Required: true},
}

var fields_set_security_token_service_preferences = []leanruntime.Field{
	{Name: "GlobalEndpointTokenVersion", Flag: "global-endpoint-token-version", Type: "types.GlobalEndpointTokenVersion", Required: true},
}

var fields_simulate_custom_policy = []leanruntime.Field{
	{Name: "ActionNames", Flag: "action-names", Type: "[]string", Required: true},
	{Name: "CallerArn", Flag: "caller-arn", Type: "*string", Required: false},
	{Name: "ContextEntries", Flag: "context-entries", Type: "[]types.ContextEntry", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "PermissionsBoundaryPolicyInputList", Flag: "permissions-boundary-policy-input-list", Type: "[]string", Required: false},
	{Name: "PolicyInputList", Flag: "policy-input-list", Type: "[]string", Required: true},
	{Name: "ResourceArns", Flag: "resource-arns", Type: "[]string", Required: false},
	{Name: "ResourceHandlingOption", Flag: "resource-handling-option", Type: "*string", Required: false},
	{Name: "ResourceOwner", Flag: "resource-owner", Type: "*string", Required: false},
	{Name: "ResourcePolicy", Flag: "resource-policy", Type: "*string", Required: false},
}

var fields_simulate_principal_policy = []leanruntime.Field{
	{Name: "ActionNames", Flag: "action-names", Type: "[]string", Required: true},
	{Name: "CallerArn", Flag: "caller-arn", Type: "*string", Required: false},
	{Name: "ContextEntries", Flag: "context-entries", Type: "[]types.ContextEntry", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "PermissionsBoundaryPolicyInputList", Flag: "permissions-boundary-policy-input-list", Type: "[]string", Required: false},
	{Name: "PolicyInputList", Flag: "policy-input-list", Type: "[]string", Required: false},
	{Name: "PolicySourceArn", Flag: "policy-source-arn", Type: "*string", Required: true},
	{Name: "ResourceArns", Flag: "resource-arns", Type: "[]string", Required: false},
	{Name: "ResourceHandlingOption", Flag: "resource-handling-option", Type: "*string", Required: false},
	{Name: "ResourceOwner", Flag: "resource-owner", Type: "*string", Required: false},
	{Name: "ResourcePolicy", Flag: "resource-policy", Type: "*string", Required: false},
}

var fields_tag_instance_profile = []leanruntime.Field{
	{Name: "InstanceProfileName", Flag: "instance-profile-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_tag_mfa_device = []leanruntime.Field{
	{Name: "SerialNumber", Flag: "serial-number", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_tag_open_id_connect_provider = []leanruntime.Field{
	{Name: "OpenIDConnectProviderArn", Flag: "open-id-connect-provider-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_tag_policy = []leanruntime.Field{
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_tag_role = []leanruntime.Field{
	{Name: "RoleName", Flag: "role-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_tag_saml_provider = []leanruntime.Field{
	{Name: "SAMLProviderArn", Flag: "saml-provider-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_tag_server_certificate = []leanruntime.Field{
	{Name: "ServerCertificateName", Flag: "server-certificate-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_tag_user = []leanruntime.Field{
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_untag_instance_profile = []leanruntime.Field{
	{Name: "InstanceProfileName", Flag: "instance-profile-name", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_untag_mfa_device = []leanruntime.Field{
	{Name: "SerialNumber", Flag: "serial-number", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_untag_open_id_connect_provider = []leanruntime.Field{
	{Name: "OpenIDConnectProviderArn", Flag: "open-id-connect-provider-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_untag_policy = []leanruntime.Field{
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_untag_role = []leanruntime.Field{
	{Name: "RoleName", Flag: "role-name", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_untag_saml_provider = []leanruntime.Field{
	{Name: "SAMLProviderArn", Flag: "saml-provider-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_untag_server_certificate = []leanruntime.Field{
	{Name: "ServerCertificateName", Flag: "server-certificate-name", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_untag_user = []leanruntime.Field{
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_update_access_key = []leanruntime.Field{
	{Name: "AccessKeyId", Flag: "access-key-id", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.StatusType", Required: true},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: false},
}

var fields_update_account_password_policy = []leanruntime.Field{
	{Name: "AllowUsersToChangePassword", Flag: "allow-users-to-change-password", Type: "bool", Required: false},
	{Name: "HardExpiry", Flag: "hard-expiry", Type: "*bool", Required: false},
	{Name: "MaxPasswordAge", Flag: "max-password-age", Type: "*int32", Required: false},
	{Name: "MinimumPasswordLength", Flag: "minimum-password-length", Type: "*int32", Required: false},
	{Name: "PasswordReusePrevention", Flag: "password-reuse-prevention", Type: "*int32", Required: false},
	{Name: "RequireLowercaseCharacters", Flag: "require-lowercase-characters", Type: "bool", Required: false},
	{Name: "RequireNumbers", Flag: "require-numbers", Type: "bool", Required: false},
	{Name: "RequireSymbols", Flag: "require-symbols", Type: "bool", Required: false},
	{Name: "RequireUppercaseCharacters", Flag: "require-uppercase-characters", Type: "bool", Required: false},
}

var fields_update_assume_role_policy = []leanruntime.Field{
	{Name: "PolicyDocument", Flag: "policy-document", Type: "*string", Required: true},
	{Name: "RoleName", Flag: "role-name", Type: "*string", Required: true},
}

var fields_update_delegation_request = []leanruntime.Field{
	{Name: "DelegationRequestId", Flag: "delegation-request-id", Type: "*string", Required: true},
	{Name: "Notes", Flag: "notes", Type: "*string", Required: false},
}

var fields_update_group = []leanruntime.Field{
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: true},
	{Name: "NewGroupName", Flag: "new-group-name", Type: "*string", Required: false},
	{Name: "NewPath", Flag: "new-path", Type: "*string", Required: false},
}

var fields_update_login_profile = []leanruntime.Field{
	{Name: "Password", Flag: "password", Type: "*string", Required: false},
	{Name: "PasswordResetRequired", Flag: "password-reset-required", Type: "*bool", Required: false},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_update_open_id_connect_provider_thumbprint = []leanruntime.Field{
	{Name: "OpenIDConnectProviderArn", Flag: "open-id-connect-provider-arn", Type: "*string", Required: true},
	{Name: "ThumbprintList", Flag: "thumbprint-list", Type: "[]string", Required: true},
}

var fields_update_role = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "MaxSessionDuration", Flag: "max-session-duration", Type: "*int32", Required: false},
	{Name: "RoleName", Flag: "role-name", Type: "*string", Required: true},
}

var fields_update_role_description = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "RoleName", Flag: "role-name", Type: "*string", Required: true},
}

var fields_update_saml_provider = []leanruntime.Field{
	{Name: "AddPrivateKey", Flag: "add-private-key", Type: "*string", Required: false},
	{Name: "AssertionEncryptionMode", Flag: "assertion-encryption-mode", Type: "types.AssertionEncryptionModeType", Required: false},
	{Name: "RemovePrivateKey", Flag: "remove-private-key", Type: "*string", Required: false},
	{Name: "SAMLMetadataDocument", Flag: "saml-metadata-document", Type: "*string", Required: false},
	{Name: "SAMLProviderArn", Flag: "saml-provider-arn", Type: "*string", Required: true},
}

var fields_update_server_certificate = []leanruntime.Field{
	{Name: "NewPath", Flag: "new-path", Type: "*string", Required: false},
	{Name: "NewServerCertificateName", Flag: "new-server-certificate-name", Type: "*string", Required: false},
	{Name: "ServerCertificateName", Flag: "server-certificate-name", Type: "*string", Required: true},
}

var fields_update_service_specific_credential = []leanruntime.Field{
	{Name: "ServiceSpecificCredentialId", Flag: "service-specific-credential-id", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.StatusType", Required: true},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: false},
}

var fields_update_signing_certificate = []leanruntime.Field{
	{Name: "CertificateId", Flag: "certificate-id", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.StatusType", Required: true},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: false},
}

var fields_update_ssh_public_key = []leanruntime.Field{
	{Name: "SSHPublicKeyId", Flag: "ssh-public-key-id", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.StatusType", Required: true},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_update_user = []leanruntime.Field{
	{Name: "NewPath", Flag: "new-path", Type: "*string", Required: false},
	{Name: "NewUserName", Flag: "new-user-name", Type: "*string", Required: false},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_upload_server_certificate = []leanruntime.Field{
	{Name: "CertificateBody", Flag: "certificate-body", Type: "*string", Required: true},
	{Name: "CertificateChain", Flag: "certificate-chain", Type: "*string", Required: false},
	{Name: "Path", Flag: "path", Type: "*string", Required: false},
	{Name: "PrivateKey", Flag: "private-key", Type: "*string", Required: true},
	{Name: "ServerCertificateName", Flag: "server-certificate-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_upload_signing_certificate = []leanruntime.Field{
	{Name: "CertificateBody", Flag: "certificate-body", Type: "*string", Required: true},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: false},
}

var fields_upload_ssh_public_key = []leanruntime.Field{
	{Name: "SSHPublicKeyBody", Flag: "ssh-public-key-body", Type: "*string", Required: true},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"accept-delegation-request": {
			Name:   "accept-delegation-request",
			Fields: fields_accept_delegation_request,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AcceptDelegationRequestInput{}
				if _, err := leanruntime.ApplyInput(input, fields_accept_delegation_request, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AcceptDelegationRequest(ctx, input)
			},
		},
		"add-client-idto-open-id-connect-provider": {
			Name:   "add-client-idto-open-id-connect-provider",
			Fields: fields_add_client_idto_open_id_connect_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddClientIDToOpenIDConnectProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_client_idto_open_id_connect_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddClientIDToOpenIDConnectProvider(ctx, input)
			},
		},
		"add-role-to-instance-profile": {
			Name:   "add-role-to-instance-profile",
			Fields: fields_add_role_to_instance_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddRoleToInstanceProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_role_to_instance_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddRoleToInstanceProfile(ctx, input)
			},
		},
		"add-user-to-group": {
			Name:   "add-user-to-group",
			Fields: fields_add_user_to_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddUserToGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_user_to_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddUserToGroup(ctx, input)
			},
		},
		"associate-delegation-request": {
			Name:   "associate-delegation-request",
			Fields: fields_associate_delegation_request,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateDelegationRequestInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_delegation_request, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateDelegationRequest(ctx, input)
			},
		},
		"attach-group-policy": {
			Name:   "attach-group-policy",
			Fields: fields_attach_group_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AttachGroupPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_attach_group_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AttachGroupPolicy(ctx, input)
			},
		},
		"attach-role-policy": {
			Name:   "attach-role-policy",
			Fields: fields_attach_role_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AttachRolePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_attach_role_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AttachRolePolicy(ctx, input)
			},
		},
		"attach-user-policy": {
			Name:   "attach-user-policy",
			Fields: fields_attach_user_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AttachUserPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_attach_user_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AttachUserPolicy(ctx, input)
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
		"create-access-key": {
			Name:   "create-access-key",
			Fields: fields_create_access_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAccessKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_access_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAccessKey(ctx, input)
			},
		},
		"create-account-alias": {
			Name:   "create-account-alias",
			Fields: fields_create_account_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAccountAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_account_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAccountAlias(ctx, input)
			},
		},
		"create-delegation-request": {
			Name:   "create-delegation-request",
			Fields: fields_create_delegation_request,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDelegationRequestInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_delegation_request, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDelegationRequest(ctx, input)
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
		"create-instance-profile": {
			Name:   "create-instance-profile",
			Fields: fields_create_instance_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateInstanceProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_instance_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateInstanceProfile(ctx, input)
			},
		},
		"create-login-profile": {
			Name:   "create-login-profile",
			Fields: fields_create_login_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLoginProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_login_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLoginProfile(ctx, input)
			},
		},
		"create-open-id-connect-provider": {
			Name:   "create-open-id-connect-provider",
			Fields: fields_create_open_id_connect_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateOpenIDConnectProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_open_id_connect_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateOpenIDConnectProvider(ctx, input)
			},
		},
		"create-policy": {
			Name:   "create-policy",
			Fields: fields_create_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePolicy(ctx, input)
			},
		},
		"create-policy-version": {
			Name:   "create-policy-version",
			Fields: fields_create_policy_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePolicyVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_policy_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePolicyVersion(ctx, input)
			},
		},
		"create-role": {
			Name:   "create-role",
			Fields: fields_create_role,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRoleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_role, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRole(ctx, input)
			},
		},
		"create-saml-provider": {
			Name:   "create-saml-provider",
			Fields: fields_create_saml_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSAMLProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_saml_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSAMLProvider(ctx, input)
			},
		},
		"create-service-linked-role": {
			Name:   "create-service-linked-role",
			Fields: fields_create_service_linked_role,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateServiceLinkedRoleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_service_linked_role, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateServiceLinkedRole(ctx, input)
			},
		},
		"create-service-specific-credential": {
			Name:   "create-service-specific-credential",
			Fields: fields_create_service_specific_credential,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateServiceSpecificCredentialInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_service_specific_credential, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateServiceSpecificCredential(ctx, input)
			},
		},
		"create-user": {
			Name:   "create-user",
			Fields: fields_create_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateUser(ctx, input)
			},
		},
		"create-virtual-mfa-device": {
			Name:   "create-virtual-mfa-device",
			Fields: fields_create_virtual_mfa_device,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVirtualMFADeviceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_virtual_mfa_device, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVirtualMFADevice(ctx, input)
			},
		},
		"deactivate-mfa-device": {
			Name:   "deactivate-mfa-device",
			Fields: fields_deactivate_mfa_device,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeactivateMFADeviceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deactivate_mfa_device, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeactivateMFADevice(ctx, input)
			},
		},
		"delete-access-key": {
			Name:   "delete-access-key",
			Fields: fields_delete_access_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAccessKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_access_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAccessKey(ctx, input)
			},
		},
		"delete-account-alias": {
			Name:   "delete-account-alias",
			Fields: fields_delete_account_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAccountAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_account_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAccountAlias(ctx, input)
			},
		},
		"delete-account-password-policy": {
			Name:   "delete-account-password-policy",
			Fields: fields_delete_account_password_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAccountPasswordPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_account_password_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAccountPasswordPolicy(ctx, input)
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
		"delete-group-policy": {
			Name:   "delete-group-policy",
			Fields: fields_delete_group_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteGroupPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_group_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteGroupPolicy(ctx, input)
			},
		},
		"delete-instance-profile": {
			Name:   "delete-instance-profile",
			Fields: fields_delete_instance_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteInstanceProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_instance_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteInstanceProfile(ctx, input)
			},
		},
		"delete-login-profile": {
			Name:   "delete-login-profile",
			Fields: fields_delete_login_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLoginProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_login_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLoginProfile(ctx, input)
			},
		},
		"delete-open-id-connect-provider": {
			Name:   "delete-open-id-connect-provider",
			Fields: fields_delete_open_id_connect_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteOpenIDConnectProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_open_id_connect_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteOpenIDConnectProvider(ctx, input)
			},
		},
		"delete-policy": {
			Name:   "delete-policy",
			Fields: fields_delete_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePolicy(ctx, input)
			},
		},
		"delete-policy-version": {
			Name:   "delete-policy-version",
			Fields: fields_delete_policy_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePolicyVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_policy_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePolicyVersion(ctx, input)
			},
		},
		"delete-role": {
			Name:   "delete-role",
			Fields: fields_delete_role,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRoleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_role, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRole(ctx, input)
			},
		},
		"delete-role-permissions-boundary": {
			Name:   "delete-role-permissions-boundary",
			Fields: fields_delete_role_permissions_boundary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRolePermissionsBoundaryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_role_permissions_boundary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRolePermissionsBoundary(ctx, input)
			},
		},
		"delete-role-policy": {
			Name:   "delete-role-policy",
			Fields: fields_delete_role_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRolePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_role_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRolePolicy(ctx, input)
			},
		},
		"delete-saml-provider": {
			Name:   "delete-saml-provider",
			Fields: fields_delete_saml_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSAMLProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_saml_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSAMLProvider(ctx, input)
			},
		},
		"delete-server-certificate": {
			Name:   "delete-server-certificate",
			Fields: fields_delete_server_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteServerCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_server_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteServerCertificate(ctx, input)
			},
		},
		"delete-service-linked-role": {
			Name:   "delete-service-linked-role",
			Fields: fields_delete_service_linked_role,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteServiceLinkedRoleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_service_linked_role, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteServiceLinkedRole(ctx, input)
			},
		},
		"delete-service-specific-credential": {
			Name:   "delete-service-specific-credential",
			Fields: fields_delete_service_specific_credential,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteServiceSpecificCredentialInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_service_specific_credential, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteServiceSpecificCredential(ctx, input)
			},
		},
		"delete-signing-certificate": {
			Name:   "delete-signing-certificate",
			Fields: fields_delete_signing_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSigningCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_signing_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSigningCertificate(ctx, input)
			},
		},
		"delete-ssh-public-key": {
			Name:   "delete-ssh-public-key",
			Fields: fields_delete_ssh_public_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSSHPublicKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_ssh_public_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSSHPublicKey(ctx, input)
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
		"delete-user-permissions-boundary": {
			Name:   "delete-user-permissions-boundary",
			Fields: fields_delete_user_permissions_boundary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteUserPermissionsBoundaryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_user_permissions_boundary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteUserPermissionsBoundary(ctx, input)
			},
		},
		"delete-user-policy": {
			Name:   "delete-user-policy",
			Fields: fields_delete_user_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteUserPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_user_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteUserPolicy(ctx, input)
			},
		},
		"delete-virtual-mfa-device": {
			Name:   "delete-virtual-mfa-device",
			Fields: fields_delete_virtual_mfa_device,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVirtualMFADeviceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_virtual_mfa_device, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVirtualMFADevice(ctx, input)
			},
		},
		"detach-group-policy": {
			Name:   "detach-group-policy",
			Fields: fields_detach_group_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetachGroupPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detach_group_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetachGroupPolicy(ctx, input)
			},
		},
		"detach-role-policy": {
			Name:   "detach-role-policy",
			Fields: fields_detach_role_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetachRolePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detach_role_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetachRolePolicy(ctx, input)
			},
		},
		"detach-user-policy": {
			Name:   "detach-user-policy",
			Fields: fields_detach_user_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetachUserPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detach_user_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetachUserPolicy(ctx, input)
			},
		},
		"disable-organizations-root-credentials-management": {
			Name:   "disable-organizations-root-credentials-management",
			Fields: fields_disable_organizations_root_credentials_management,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableOrganizationsRootCredentialsManagementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_organizations_root_credentials_management, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableOrganizationsRootCredentialsManagement(ctx, input)
			},
		},
		"disable-organizations-root-sessions": {
			Name:   "disable-organizations-root-sessions",
			Fields: fields_disable_organizations_root_sessions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableOrganizationsRootSessionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_organizations_root_sessions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableOrganizationsRootSessions(ctx, input)
			},
		},
		"disable-outbound-web-identity-federation": {
			Name:   "disable-outbound-web-identity-federation",
			Fields: fields_disable_outbound_web_identity_federation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableOutboundWebIdentityFederationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_outbound_web_identity_federation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableOutboundWebIdentityFederation(ctx, input)
			},
		},
		"enable-mfa-device": {
			Name:   "enable-mfa-device",
			Fields: fields_enable_mfa_device,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableMFADeviceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_mfa_device, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableMFADevice(ctx, input)
			},
		},
		"enable-organizations-root-credentials-management": {
			Name:   "enable-organizations-root-credentials-management",
			Fields: fields_enable_organizations_root_credentials_management,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableOrganizationsRootCredentialsManagementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_organizations_root_credentials_management, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableOrganizationsRootCredentialsManagement(ctx, input)
			},
		},
		"enable-organizations-root-sessions": {
			Name:   "enable-organizations-root-sessions",
			Fields: fields_enable_organizations_root_sessions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableOrganizationsRootSessionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_organizations_root_sessions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableOrganizationsRootSessions(ctx, input)
			},
		},
		"enable-outbound-web-identity-federation": {
			Name:   "enable-outbound-web-identity-federation",
			Fields: fields_enable_outbound_web_identity_federation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableOutboundWebIdentityFederationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_outbound_web_identity_federation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableOutboundWebIdentityFederation(ctx, input)
			},
		},
		"generate-credential-report": {
			Name:   "generate-credential-report",
			Fields: fields_generate_credential_report,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GenerateCredentialReportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_generate_credential_report, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GenerateCredentialReport(ctx, input)
			},
		},
		"generate-organizations-access-report": {
			Name:   "generate-organizations-access-report",
			Fields: fields_generate_organizations_access_report,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GenerateOrganizationsAccessReportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_generate_organizations_access_report, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GenerateOrganizationsAccessReport(ctx, input)
			},
		},
		"generate-service-last-accessed-details": {
			Name:   "generate-service-last-accessed-details",
			Fields: fields_generate_service_last_accessed_details,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GenerateServiceLastAccessedDetailsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_generate_service_last_accessed_details, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GenerateServiceLastAccessedDetails(ctx, input)
			},
		},
		"get-access-key-last-used": {
			Name:   "get-access-key-last-used",
			Fields: fields_get_access_key_last_used,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccessKeyLastUsedInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_access_key_last_used, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccessKeyLastUsed(ctx, input)
			},
		},
		"get-account-authorization-details": {
			Name:   "get-account-authorization-details",
			Fields: fields_get_account_authorization_details,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccountAuthorizationDetailsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_account_authorization_details, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetAccountAuthorizationDetails(ctx, input)
				}
				var results []*svc.GetAccountAuthorizationDetailsOutput
				p := svc.NewGetAccountAuthorizationDetailsPaginator(client, input)
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
		"get-account-password-policy": {
			Name:   "get-account-password-policy",
			Fields: fields_get_account_password_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccountPasswordPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_account_password_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccountPasswordPolicy(ctx, input)
			},
		},
		"get-account-summary": {
			Name:   "get-account-summary",
			Fields: fields_get_account_summary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccountSummaryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_account_summary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccountSummary(ctx, input)
			},
		},
		"get-context-keys-for-custom-policy": {
			Name:   "get-context-keys-for-custom-policy",
			Fields: fields_get_context_keys_for_custom_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetContextKeysForCustomPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_context_keys_for_custom_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetContextKeysForCustomPolicy(ctx, input)
			},
		},
		"get-context-keys-for-principal-policy": {
			Name:   "get-context-keys-for-principal-policy",
			Fields: fields_get_context_keys_for_principal_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetContextKeysForPrincipalPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_context_keys_for_principal_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetContextKeysForPrincipalPolicy(ctx, input)
			},
		},
		"get-credential-report": {
			Name:   "get-credential-report",
			Fields: fields_get_credential_report,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCredentialReportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_credential_report, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCredentialReport(ctx, input)
			},
		},
		"get-delegation-request": {
			Name:   "get-delegation-request",
			Fields: fields_get_delegation_request,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDelegationRequestInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_delegation_request, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDelegationRequest(ctx, input)
			},
		},
		"get-group": {
			Name:   "get-group",
			Fields: fields_get_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetGroupInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_group, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetGroup(ctx, input)
				}
				var results []*svc.GetGroupOutput
				p := svc.NewGetGroupPaginator(client, input)
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
		"get-group-policy": {
			Name:   "get-group-policy",
			Fields: fields_get_group_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetGroupPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_group_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetGroupPolicy(ctx, input)
			},
		},
		"get-human-readable-summary": {
			Name:   "get-human-readable-summary",
			Fields: fields_get_human_readable_summary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetHumanReadableSummaryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_human_readable_summary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetHumanReadableSummary(ctx, input)
			},
		},
		"get-instance-profile": {
			Name:   "get-instance-profile",
			Fields: fields_get_instance_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetInstanceProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_instance_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetInstanceProfile(ctx, input)
			},
		},
		"get-login-profile": {
			Name:   "get-login-profile",
			Fields: fields_get_login_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLoginProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_login_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLoginProfile(ctx, input)
			},
		},
		"get-mfa-device": {
			Name:   "get-mfa-device",
			Fields: fields_get_mfa_device,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMFADeviceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_mfa_device, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMFADevice(ctx, input)
			},
		},
		"get-open-id-connect-provider": {
			Name:   "get-open-id-connect-provider",
			Fields: fields_get_open_id_connect_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOpenIDConnectProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_open_id_connect_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetOpenIDConnectProvider(ctx, input)
			},
		},
		"get-organizations-access-report": {
			Name:   "get-organizations-access-report",
			Fields: fields_get_organizations_access_report,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOrganizationsAccessReportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_organizations_access_report, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetOrganizationsAccessReport(ctx, input)
			},
		},
		"get-outbound-web-identity-federation-info": {
			Name:   "get-outbound-web-identity-federation-info",
			Fields: fields_get_outbound_web_identity_federation_info,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOutboundWebIdentityFederationInfoInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_outbound_web_identity_federation_info, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetOutboundWebIdentityFederationInfo(ctx, input)
			},
		},
		"get-policy": {
			Name:   "get-policy",
			Fields: fields_get_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPolicy(ctx, input)
			},
		},
		"get-policy-version": {
			Name:   "get-policy-version",
			Fields: fields_get_policy_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPolicyVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_policy_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPolicyVersion(ctx, input)
			},
		},
		"get-role": {
			Name:   "get-role",
			Fields: fields_get_role,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRoleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_role, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRole(ctx, input)
			},
		},
		"get-role-policy": {
			Name:   "get-role-policy",
			Fields: fields_get_role_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRolePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_role_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRolePolicy(ctx, input)
			},
		},
		"get-saml-provider": {
			Name:   "get-saml-provider",
			Fields: fields_get_saml_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSAMLProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_saml_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSAMLProvider(ctx, input)
			},
		},
		"get-server-certificate": {
			Name:   "get-server-certificate",
			Fields: fields_get_server_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetServerCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_server_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetServerCertificate(ctx, input)
			},
		},
		"get-service-last-accessed-details": {
			Name:   "get-service-last-accessed-details",
			Fields: fields_get_service_last_accessed_details,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetServiceLastAccessedDetailsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_service_last_accessed_details, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetServiceLastAccessedDetails(ctx, input)
			},
		},
		"get-service-last-accessed-details-with-entities": {
			Name:   "get-service-last-accessed-details-with-entities",
			Fields: fields_get_service_last_accessed_details_with_entities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetServiceLastAccessedDetailsWithEntitiesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_service_last_accessed_details_with_entities, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetServiceLastAccessedDetailsWithEntities(ctx, input)
			},
		},
		"get-service-linked-role-deletion-status": {
			Name:   "get-service-linked-role-deletion-status",
			Fields: fields_get_service_linked_role_deletion_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetServiceLinkedRoleDeletionStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_service_linked_role_deletion_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetServiceLinkedRoleDeletionStatus(ctx, input)
			},
		},
		"get-ssh-public-key": {
			Name:   "get-ssh-public-key",
			Fields: fields_get_ssh_public_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSSHPublicKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_ssh_public_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSSHPublicKey(ctx, input)
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
		"get-user-policy": {
			Name:   "get-user-policy",
			Fields: fields_get_user_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetUserPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_user_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetUserPolicy(ctx, input)
			},
		},
		"list-access-keys": {
			Name:   "list-access-keys",
			Fields: fields_list_access_keys,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAccessKeysInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_access_keys, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAccessKeys(ctx, input)
				}
				var results []*svc.ListAccessKeysOutput
				p := svc.NewListAccessKeysPaginator(client, input)
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
		"list-account-aliases": {
			Name:   "list-account-aliases",
			Fields: fields_list_account_aliases,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAccountAliasesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_account_aliases, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAccountAliases(ctx, input)
				}
				var results []*svc.ListAccountAliasesOutput
				p := svc.NewListAccountAliasesPaginator(client, input)
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
		"list-attached-group-policies": {
			Name:   "list-attached-group-policies",
			Fields: fields_list_attached_group_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAttachedGroupPoliciesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_attached_group_policies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAttachedGroupPolicies(ctx, input)
				}
				var results []*svc.ListAttachedGroupPoliciesOutput
				p := svc.NewListAttachedGroupPoliciesPaginator(client, input)
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
		"list-attached-role-policies": {
			Name:   "list-attached-role-policies",
			Fields: fields_list_attached_role_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAttachedRolePoliciesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_attached_role_policies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAttachedRolePolicies(ctx, input)
				}
				var results []*svc.ListAttachedRolePoliciesOutput
				p := svc.NewListAttachedRolePoliciesPaginator(client, input)
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
		"list-attached-user-policies": {
			Name:   "list-attached-user-policies",
			Fields: fields_list_attached_user_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAttachedUserPoliciesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_attached_user_policies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAttachedUserPolicies(ctx, input)
				}
				var results []*svc.ListAttachedUserPoliciesOutput
				p := svc.NewListAttachedUserPoliciesPaginator(client, input)
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
		"list-delegation-requests": {
			Name:   "list-delegation-requests",
			Fields: fields_list_delegation_requests,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDelegationRequestsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_delegation_requests, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListDelegationRequests(ctx, input)
			},
		},
		"list-entities-for-policy": {
			Name:   "list-entities-for-policy",
			Fields: fields_list_entities_for_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEntitiesForPolicyInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_entities_for_policy, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEntitiesForPolicy(ctx, input)
				}
				var results []*svc.ListEntitiesForPolicyOutput
				p := svc.NewListEntitiesForPolicyPaginator(client, input)
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
		"list-group-policies": {
			Name:   "list-group-policies",
			Fields: fields_list_group_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGroupPoliciesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_group_policies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListGroupPolicies(ctx, input)
				}
				var results []*svc.ListGroupPoliciesOutput
				p := svc.NewListGroupPoliciesPaginator(client, input)
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
		"list-groups-for-user": {
			Name:   "list-groups-for-user",
			Fields: fields_list_groups_for_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGroupsForUserInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_groups_for_user, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListGroupsForUser(ctx, input)
				}
				var results []*svc.ListGroupsForUserOutput
				p := svc.NewListGroupsForUserPaginator(client, input)
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
		"list-instance-profile-tags": {
			Name:   "list-instance-profile-tags",
			Fields: fields_list_instance_profile_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInstanceProfileTagsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_instance_profile_tags, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListInstanceProfileTags(ctx, input)
				}
				var results []*svc.ListInstanceProfileTagsOutput
				p := svc.NewListInstanceProfileTagsPaginator(client, input)
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
		"list-instance-profiles": {
			Name:   "list-instance-profiles",
			Fields: fields_list_instance_profiles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInstanceProfilesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_instance_profiles, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListInstanceProfiles(ctx, input)
				}
				var results []*svc.ListInstanceProfilesOutput
				p := svc.NewListInstanceProfilesPaginator(client, input)
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
		"list-instance-profiles-for-role": {
			Name:   "list-instance-profiles-for-role",
			Fields: fields_list_instance_profiles_for_role,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInstanceProfilesForRoleInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_instance_profiles_for_role, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListInstanceProfilesForRole(ctx, input)
				}
				var results []*svc.ListInstanceProfilesForRoleOutput
				p := svc.NewListInstanceProfilesForRolePaginator(client, input)
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
		"list-mfa-device-tags": {
			Name:   "list-mfa-device-tags",
			Fields: fields_list_mfa_device_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMFADeviceTagsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_mfa_device_tags, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMFADeviceTags(ctx, input)
				}
				var results []*svc.ListMFADeviceTagsOutput
				p := svc.NewListMFADeviceTagsPaginator(client, input)
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
		"list-mfa-devices": {
			Name:   "list-mfa-devices",
			Fields: fields_list_mfa_devices,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMFADevicesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_mfa_devices, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMFADevices(ctx, input)
				}
				var results []*svc.ListMFADevicesOutput
				p := svc.NewListMFADevicesPaginator(client, input)
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
		"list-open-id-connect-provider-tags": {
			Name:   "list-open-id-connect-provider-tags",
			Fields: fields_list_open_id_connect_provider_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOpenIDConnectProviderTagsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_open_id_connect_provider_tags, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListOpenIDConnectProviderTags(ctx, input)
				}
				var results []*svc.ListOpenIDConnectProviderTagsOutput
				p := svc.NewListOpenIDConnectProviderTagsPaginator(client, input)
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
		"list-open-id-connect-providers": {
			Name:   "list-open-id-connect-providers",
			Fields: fields_list_open_id_connect_providers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOpenIDConnectProvidersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_open_id_connect_providers, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListOpenIDConnectProviders(ctx, input)
			},
		},
		"list-organizations-features": {
			Name:   "list-organizations-features",
			Fields: fields_list_organizations_features,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOrganizationsFeaturesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_organizations_features, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListOrganizationsFeatures(ctx, input)
			},
		},
		"list-policies": {
			Name:   "list-policies",
			Fields: fields_list_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPoliciesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_policies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPolicies(ctx, input)
				}
				var results []*svc.ListPoliciesOutput
				p := svc.NewListPoliciesPaginator(client, input)
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
		"list-policies-granting-service-access": {
			Name:   "list-policies-granting-service-access",
			Fields: fields_list_policies_granting_service_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPoliciesGrantingServiceAccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_policies_granting_service_access, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListPoliciesGrantingServiceAccess(ctx, input)
			},
		},
		"list-policy-tags": {
			Name:   "list-policy-tags",
			Fields: fields_list_policy_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPolicyTagsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_policy_tags, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPolicyTags(ctx, input)
				}
				var results []*svc.ListPolicyTagsOutput
				p := svc.NewListPolicyTagsPaginator(client, input)
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
		"list-policy-versions": {
			Name:   "list-policy-versions",
			Fields: fields_list_policy_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPolicyVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_policy_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPolicyVersions(ctx, input)
				}
				var results []*svc.ListPolicyVersionsOutput
				p := svc.NewListPolicyVersionsPaginator(client, input)
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
		"list-role-policies": {
			Name:   "list-role-policies",
			Fields: fields_list_role_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRolePoliciesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_role_policies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRolePolicies(ctx, input)
				}
				var results []*svc.ListRolePoliciesOutput
				p := svc.NewListRolePoliciesPaginator(client, input)
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
		"list-role-tags": {
			Name:   "list-role-tags",
			Fields: fields_list_role_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRoleTagsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_role_tags, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRoleTags(ctx, input)
				}
				var results []*svc.ListRoleTagsOutput
				p := svc.NewListRoleTagsPaginator(client, input)
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
		"list-roles": {
			Name:   "list-roles",
			Fields: fields_list_roles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRolesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_roles, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRoles(ctx, input)
				}
				var results []*svc.ListRolesOutput
				p := svc.NewListRolesPaginator(client, input)
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
		"list-saml-provider-tags": {
			Name:   "list-saml-provider-tags",
			Fields: fields_list_saml_provider_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSAMLProviderTagsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_saml_provider_tags, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSAMLProviderTags(ctx, input)
				}
				var results []*svc.ListSAMLProviderTagsOutput
				p := svc.NewListSAMLProviderTagsPaginator(client, input)
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
		"list-saml-providers": {
			Name:   "list-saml-providers",
			Fields: fields_list_saml_providers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSAMLProvidersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_saml_providers, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListSAMLProviders(ctx, input)
			},
		},
		"list-server-certificate-tags": {
			Name:   "list-server-certificate-tags",
			Fields: fields_list_server_certificate_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListServerCertificateTagsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_server_certificate_tags, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListServerCertificateTags(ctx, input)
				}
				var results []*svc.ListServerCertificateTagsOutput
				p := svc.NewListServerCertificateTagsPaginator(client, input)
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
		"list-server-certificates": {
			Name:   "list-server-certificates",
			Fields: fields_list_server_certificates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListServerCertificatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_server_certificates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListServerCertificates(ctx, input)
				}
				var results []*svc.ListServerCertificatesOutput
				p := svc.NewListServerCertificatesPaginator(client, input)
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
		"list-service-specific-credentials": {
			Name:   "list-service-specific-credentials",
			Fields: fields_list_service_specific_credentials,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListServiceSpecificCredentialsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_service_specific_credentials, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListServiceSpecificCredentials(ctx, input)
			},
		},
		"list-signing-certificates": {
			Name:   "list-signing-certificates",
			Fields: fields_list_signing_certificates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSigningCertificatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_signing_certificates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSigningCertificates(ctx, input)
				}
				var results []*svc.ListSigningCertificatesOutput
				p := svc.NewListSigningCertificatesPaginator(client, input)
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
		"list-ssh-public-keys": {
			Name:   "list-ssh-public-keys",
			Fields: fields_list_ssh_public_keys,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSSHPublicKeysInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_ssh_public_keys, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSSHPublicKeys(ctx, input)
				}
				var results []*svc.ListSSHPublicKeysOutput
				p := svc.NewListSSHPublicKeysPaginator(client, input)
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
		"list-user-policies": {
			Name:   "list-user-policies",
			Fields: fields_list_user_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListUserPoliciesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_user_policies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListUserPolicies(ctx, input)
				}
				var results []*svc.ListUserPoliciesOutput
				p := svc.NewListUserPoliciesPaginator(client, input)
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
		"list-user-tags": {
			Name:   "list-user-tags",
			Fields: fields_list_user_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListUserTagsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_user_tags, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListUserTags(ctx, input)
				}
				var results []*svc.ListUserTagsOutput
				p := svc.NewListUserTagsPaginator(client, input)
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
		"list-virtual-mfa-devices": {
			Name:   "list-virtual-mfa-devices",
			Fields: fields_list_virtual_mfa_devices,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVirtualMFADevicesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_virtual_mfa_devices, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListVirtualMFADevices(ctx, input)
				}
				var results []*svc.ListVirtualMFADevicesOutput
				p := svc.NewListVirtualMFADevicesPaginator(client, input)
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
		"put-group-policy": {
			Name:   "put-group-policy",
			Fields: fields_put_group_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutGroupPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_group_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutGroupPolicy(ctx, input)
			},
		},
		"put-role-permissions-boundary": {
			Name:   "put-role-permissions-boundary",
			Fields: fields_put_role_permissions_boundary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutRolePermissionsBoundaryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_role_permissions_boundary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutRolePermissionsBoundary(ctx, input)
			},
		},
		"put-role-policy": {
			Name:   "put-role-policy",
			Fields: fields_put_role_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutRolePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_role_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutRolePolicy(ctx, input)
			},
		},
		"put-user-permissions-boundary": {
			Name:   "put-user-permissions-boundary",
			Fields: fields_put_user_permissions_boundary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutUserPermissionsBoundaryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_user_permissions_boundary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutUserPermissionsBoundary(ctx, input)
			},
		},
		"put-user-policy": {
			Name:   "put-user-policy",
			Fields: fields_put_user_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutUserPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_user_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutUserPolicy(ctx, input)
			},
		},
		"reject-delegation-request": {
			Name:   "reject-delegation-request",
			Fields: fields_reject_delegation_request,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RejectDelegationRequestInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reject_delegation_request, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RejectDelegationRequest(ctx, input)
			},
		},
		"remove-client-id-from-open-id-connect-provider": {
			Name:   "remove-client-id-from-open-id-connect-provider",
			Fields: fields_remove_client_id_from_open_id_connect_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveClientIDFromOpenIDConnectProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_client_id_from_open_id_connect_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveClientIDFromOpenIDConnectProvider(ctx, input)
			},
		},
		"remove-role-from-instance-profile": {
			Name:   "remove-role-from-instance-profile",
			Fields: fields_remove_role_from_instance_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveRoleFromInstanceProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_role_from_instance_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveRoleFromInstanceProfile(ctx, input)
			},
		},
		"remove-user-from-group": {
			Name:   "remove-user-from-group",
			Fields: fields_remove_user_from_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveUserFromGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_user_from_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveUserFromGroup(ctx, input)
			},
		},
		"reset-service-specific-credential": {
			Name:   "reset-service-specific-credential",
			Fields: fields_reset_service_specific_credential,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResetServiceSpecificCredentialInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reset_service_specific_credential, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResetServiceSpecificCredential(ctx, input)
			},
		},
		"resync-mfa-device": {
			Name:   "resync-mfa-device",
			Fields: fields_resync_mfa_device,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResyncMFADeviceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_resync_mfa_device, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResyncMFADevice(ctx, input)
			},
		},
		"send-delegation-token": {
			Name:   "send-delegation-token",
			Fields: fields_send_delegation_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendDelegationTokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_delegation_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendDelegationToken(ctx, input)
			},
		},
		"set-default-policy-version": {
			Name:   "set-default-policy-version",
			Fields: fields_set_default_policy_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetDefaultPolicyVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_default_policy_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetDefaultPolicyVersion(ctx, input)
			},
		},
		"set-security-token-service-preferences": {
			Name:   "set-security-token-service-preferences",
			Fields: fields_set_security_token_service_preferences,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetSecurityTokenServicePreferencesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_security_token_service_preferences, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetSecurityTokenServicePreferences(ctx, input)
			},
		},
		"simulate-custom-policy": {
			Name:   "simulate-custom-policy",
			Fields: fields_simulate_custom_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SimulateCustomPolicyInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_simulate_custom_policy, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SimulateCustomPolicy(ctx, input)
				}
				var results []*svc.SimulateCustomPolicyOutput
				p := svc.NewSimulateCustomPolicyPaginator(client, input)
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
		"simulate-principal-policy": {
			Name:   "simulate-principal-policy",
			Fields: fields_simulate_principal_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SimulatePrincipalPolicyInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_simulate_principal_policy, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SimulatePrincipalPolicy(ctx, input)
				}
				var results []*svc.SimulatePrincipalPolicyOutput
				p := svc.NewSimulatePrincipalPolicyPaginator(client, input)
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
		"tag-instance-profile": {
			Name:   "tag-instance-profile",
			Fields: fields_tag_instance_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TagInstanceProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_tag_instance_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TagInstanceProfile(ctx, input)
			},
		},
		"tag-mfa-device": {
			Name:   "tag-mfa-device",
			Fields: fields_tag_mfa_device,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TagMFADeviceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_tag_mfa_device, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TagMFADevice(ctx, input)
			},
		},
		"tag-open-id-connect-provider": {
			Name:   "tag-open-id-connect-provider",
			Fields: fields_tag_open_id_connect_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TagOpenIDConnectProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_tag_open_id_connect_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TagOpenIDConnectProvider(ctx, input)
			},
		},
		"tag-policy": {
			Name:   "tag-policy",
			Fields: fields_tag_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TagPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_tag_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TagPolicy(ctx, input)
			},
		},
		"tag-role": {
			Name:   "tag-role",
			Fields: fields_tag_role,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TagRoleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_tag_role, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TagRole(ctx, input)
			},
		},
		"tag-saml-provider": {
			Name:   "tag-saml-provider",
			Fields: fields_tag_saml_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TagSAMLProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_tag_saml_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TagSAMLProvider(ctx, input)
			},
		},
		"tag-server-certificate": {
			Name:   "tag-server-certificate",
			Fields: fields_tag_server_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TagServerCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_tag_server_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TagServerCertificate(ctx, input)
			},
		},
		"tag-user": {
			Name:   "tag-user",
			Fields: fields_tag_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TagUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_tag_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TagUser(ctx, input)
			},
		},
		"untag-instance-profile": {
			Name:   "untag-instance-profile",
			Fields: fields_untag_instance_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UntagInstanceProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_untag_instance_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UntagInstanceProfile(ctx, input)
			},
		},
		"untag-mfa-device": {
			Name:   "untag-mfa-device",
			Fields: fields_untag_mfa_device,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UntagMFADeviceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_untag_mfa_device, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UntagMFADevice(ctx, input)
			},
		},
		"untag-open-id-connect-provider": {
			Name:   "untag-open-id-connect-provider",
			Fields: fields_untag_open_id_connect_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UntagOpenIDConnectProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_untag_open_id_connect_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UntagOpenIDConnectProvider(ctx, input)
			},
		},
		"untag-policy": {
			Name:   "untag-policy",
			Fields: fields_untag_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UntagPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_untag_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UntagPolicy(ctx, input)
			},
		},
		"untag-role": {
			Name:   "untag-role",
			Fields: fields_untag_role,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UntagRoleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_untag_role, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UntagRole(ctx, input)
			},
		},
		"untag-saml-provider": {
			Name:   "untag-saml-provider",
			Fields: fields_untag_saml_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UntagSAMLProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_untag_saml_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UntagSAMLProvider(ctx, input)
			},
		},
		"untag-server-certificate": {
			Name:   "untag-server-certificate",
			Fields: fields_untag_server_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UntagServerCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_untag_server_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UntagServerCertificate(ctx, input)
			},
		},
		"untag-user": {
			Name:   "untag-user",
			Fields: fields_untag_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UntagUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_untag_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UntagUser(ctx, input)
			},
		},
		"update-access-key": {
			Name:   "update-access-key",
			Fields: fields_update_access_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAccessKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_access_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAccessKey(ctx, input)
			},
		},
		"update-account-password-policy": {
			Name:   "update-account-password-policy",
			Fields: fields_update_account_password_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAccountPasswordPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_account_password_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAccountPasswordPolicy(ctx, input)
			},
		},
		"update-assume-role-policy": {
			Name:   "update-assume-role-policy",
			Fields: fields_update_assume_role_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAssumeRolePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_assume_role_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAssumeRolePolicy(ctx, input)
			},
		},
		"update-delegation-request": {
			Name:   "update-delegation-request",
			Fields: fields_update_delegation_request,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDelegationRequestInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_delegation_request, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDelegationRequest(ctx, input)
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
		"update-login-profile": {
			Name:   "update-login-profile",
			Fields: fields_update_login_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateLoginProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_login_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateLoginProfile(ctx, input)
			},
		},
		"update-open-id-connect-provider-thumbprint": {
			Name:   "update-open-id-connect-provider-thumbprint",
			Fields: fields_update_open_id_connect_provider_thumbprint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateOpenIDConnectProviderThumbprintInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_open_id_connect_provider_thumbprint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateOpenIDConnectProviderThumbprint(ctx, input)
			},
		},
		"update-role": {
			Name:   "update-role",
			Fields: fields_update_role,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRoleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_role, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRole(ctx, input)
			},
		},
		"update-role-description": {
			Name:   "update-role-description",
			Fields: fields_update_role_description,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRoleDescriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_role_description, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRoleDescription(ctx, input)
			},
		},
		"update-saml-provider": {
			Name:   "update-saml-provider",
			Fields: fields_update_saml_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSAMLProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_saml_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSAMLProvider(ctx, input)
			},
		},
		"update-server-certificate": {
			Name:   "update-server-certificate",
			Fields: fields_update_server_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateServerCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_server_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateServerCertificate(ctx, input)
			},
		},
		"update-service-specific-credential": {
			Name:   "update-service-specific-credential",
			Fields: fields_update_service_specific_credential,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateServiceSpecificCredentialInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_service_specific_credential, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateServiceSpecificCredential(ctx, input)
			},
		},
		"update-signing-certificate": {
			Name:   "update-signing-certificate",
			Fields: fields_update_signing_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSigningCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_signing_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSigningCertificate(ctx, input)
			},
		},
		"update-ssh-public-key": {
			Name:   "update-ssh-public-key",
			Fields: fields_update_ssh_public_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSSHPublicKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_ssh_public_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSSHPublicKey(ctx, input)
			},
		},
		"update-user": {
			Name:   "update-user",
			Fields: fields_update_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateUser(ctx, input)
			},
		},
		"upload-server-certificate": {
			Name:   "upload-server-certificate",
			Fields: fields_upload_server_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UploadServerCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_upload_server_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UploadServerCertificate(ctx, input)
			},
		},
		"upload-signing-certificate": {
			Name:   "upload-signing-certificate",
			Fields: fields_upload_signing_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UploadSigningCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_upload_signing_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UploadSigningCertificate(ctx, input)
			},
		},
		"upload-ssh-public-key": {
			Name:   "upload-ssh-public-key",
			Fields: fields_upload_ssh_public_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UploadSSHPublicKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_upload_ssh_public_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UploadSSHPublicKey(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("iam", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
