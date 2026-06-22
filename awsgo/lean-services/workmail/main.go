package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/workmail"
)

var fields_associate_delegate_to_resource = []leanruntime.Field{
	{Name: "EntityId", Flag: "entity-id", Type: "*string", Required: true},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
}

var fields_associate_member_to_group = []leanruntime.Field{
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: true},
	{Name: "MemberId", Flag: "member-id", Type: "*string", Required: true},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_assume_impersonation_role = []leanruntime.Field{
	{Name: "ImpersonationRoleId", Flag: "impersonation-role-id", Type: "*string", Required: true},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_cancel_mailbox_export_job = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_create_alias = []leanruntime.Field{
	{Name: "Alias", Flag: "alias", Type: "*string", Required: true},
	{Name: "EntityId", Flag: "entity-id", Type: "*string", Required: true},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_create_availability_configuration = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "EwsProvider", Flag: "ews-provider", Type: "*types.EwsAvailabilityProvider", Required: false},
	{Name: "LambdaProvider", Flag: "lambda-provider", Type: "*types.LambdaAvailabilityProvider", Required: false},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_create_group = []leanruntime.Field{
	{Name: "HiddenFromGlobalAddressList", Flag: "hidden-from-global-address-list", Type: "bool", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_create_identity_center_application = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_create_impersonation_role = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
	{Name: "Rules", Flag: "rules", Type: "[]types.ImpersonationRule", Required: true},
	{Name: "Type", Flag: "type", Type: "types.ImpersonationRoleType", Required: true},
}

var fields_create_mobile_device_access_rule = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DeviceModels", Flag: "device-models", Type: "[]string", Required: false},
	{Name: "DeviceOperatingSystems", Flag: "device-operating-systems", Type: "[]string", Required: false},
	{Name: "DeviceTypes", Flag: "device-types", Type: "[]string", Required: false},
	{Name: "DeviceUserAgents", Flag: "device-user-agents", Type: "[]string", Required: false},
	{Name: "Effect", Flag: "effect", Type: "types.MobileDeviceAccessRuleEffect", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NotDeviceModels", Flag: "not-device-models", Type: "[]string", Required: false},
	{Name: "NotDeviceOperatingSystems", Flag: "not-device-operating-systems", Type: "[]string", Required: false},
	{Name: "NotDeviceTypes", Flag: "not-device-types", Type: "[]string", Required: false},
	{Name: "NotDeviceUserAgents", Flag: "not-device-user-agents", Type: "[]string", Required: false},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_create_organization = []leanruntime.Field{
	{Name: "Alias", Flag: "alias", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: false},
	{Name: "Domains", Flag: "domains", Type: "[]types.Domain", Required: false},
	{Name: "EnableInteroperability", Flag: "enable-interoperability", Type: "bool", Required: false},
	{Name: "KmsKeyArn", Flag: "kms-key-arn", Type: "*string", Required: false},
}

var fields_create_resource = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "HiddenFromGlobalAddressList", Flag: "hidden-from-global-address-list", Type: "bool", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.ResourceType", Required: true},
}

var fields_create_user = []leanruntime.Field{
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: true},
	{Name: "FirstName", Flag: "first-name", Type: "*string", Required: false},
	{Name: "HiddenFromGlobalAddressList", Flag: "hidden-from-global-address-list", Type: "bool", Required: false},
	{Name: "IdentityProviderUserId", Flag: "identity-provider-user-id", Type: "*string", Required: false},
	{Name: "LastName", Flag: "last-name", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
	{Name: "Password", Flag: "password", Type: "*string", Required: false},
	{Name: "Role", Flag: "role", Type: "types.UserRole", Required: false},
}

var fields_delete_access_control_rule = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_delete_alias = []leanruntime.Field{
	{Name: "Alias", Flag: "alias", Type: "*string", Required: true},
	{Name: "EntityId", Flag: "entity-id", Type: "*string", Required: true},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_delete_availability_configuration = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_delete_email_monitoring_configuration = []leanruntime.Field{
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_delete_group = []leanruntime.Field{
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: true},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_delete_identity_center_application = []leanruntime.Field{
	{Name: "ApplicationArn", Flag: "application-arn", Type: "*string", Required: true},
}

var fields_delete_identity_provider_configuration = []leanruntime.Field{
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_delete_impersonation_role = []leanruntime.Field{
	{Name: "ImpersonationRoleId", Flag: "impersonation-role-id", Type: "*string", Required: true},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_delete_mailbox_permissions = []leanruntime.Field{
	{Name: "EntityId", Flag: "entity-id", Type: "*string", Required: true},
	{Name: "GranteeId", Flag: "grantee-id", Type: "*string", Required: true},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_delete_mobile_device_access_override = []leanruntime.Field{
	{Name: "DeviceId", Flag: "device-id", Type: "*string", Required: true},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_delete_mobile_device_access_rule = []leanruntime.Field{
	{Name: "MobileDeviceAccessRuleId", Flag: "mobile-device-access-rule-id", Type: "*string", Required: true},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_delete_organization = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DeleteDirectory", Flag: "delete-directory", Type: "bool", Required: true},
	{Name: "DeleteIdentityCenterApplication", Flag: "delete-identity-center-application", Type: "bool", Required: false},
	{Name: "ForceDelete", Flag: "force-delete", Type: "bool", Required: false},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_delete_personal_access_token = []leanruntime.Field{
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
	{Name: "PersonalAccessTokenId", Flag: "personal-access-token-id", Type: "*string", Required: true},
}

var fields_delete_resource = []leanruntime.Field{
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
}

var fields_delete_retention_policy = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_delete_user = []leanruntime.Field{
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_deregister_from_work_mail = []leanruntime.Field{
	{Name: "EntityId", Flag: "entity-id", Type: "*string", Required: true},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_deregister_mail_domain = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_describe_email_monitoring_configuration = []leanruntime.Field{
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_describe_entity = []leanruntime.Field{
	{Name: "Email", Flag: "email", Type: "*string", Required: true},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_describe_group = []leanruntime.Field{
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: true},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_describe_identity_provider_configuration = []leanruntime.Field{
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_describe_inbound_dmarc_settings = []leanruntime.Field{
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_describe_mailbox_export_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_describe_organization = []leanruntime.Field{
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_describe_resource = []leanruntime.Field{
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
}

var fields_describe_user = []leanruntime.Field{
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_disassociate_delegate_from_resource = []leanruntime.Field{
	{Name: "EntityId", Flag: "entity-id", Type: "*string", Required: true},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
}

var fields_disassociate_member_from_group = []leanruntime.Field{
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: true},
	{Name: "MemberId", Flag: "member-id", Type: "*string", Required: true},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_get_access_control_effect = []leanruntime.Field{
	{Name: "Action", Flag: "action", Type: "*string", Required: true},
	{Name: "ImpersonationRoleId", Flag: "impersonation-role-id", Type: "*string", Required: false},
	{Name: "IpAddress", Flag: "ip-address", Type: "*string", Required: true},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: false},
}

var fields_get_default_retention_policy = []leanruntime.Field{
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_get_impersonation_role = []leanruntime.Field{
	{Name: "ImpersonationRoleId", Flag: "impersonation-role-id", Type: "*string", Required: true},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_get_impersonation_role_effect = []leanruntime.Field{
	{Name: "ImpersonationRoleId", Flag: "impersonation-role-id", Type: "*string", Required: true},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
	{Name: "TargetUser", Flag: "target-user", Type: "*string", Required: true},
}

var fields_get_mail_domain = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_get_mailbox_details = []leanruntime.Field{
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_get_mobile_device_access_effect = []leanruntime.Field{
	{Name: "DeviceModel", Flag: "device-model", Type: "*string", Required: false},
	{Name: "DeviceOperatingSystem", Flag: "device-operating-system", Type: "*string", Required: false},
	{Name: "DeviceType", Flag: "device-type", Type: "*string", Required: false},
	{Name: "DeviceUserAgent", Flag: "device-user-agent", Type: "*string", Required: false},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_get_mobile_device_access_override = []leanruntime.Field{
	{Name: "DeviceId", Flag: "device-id", Type: "*string", Required: true},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_get_personal_access_token_metadata = []leanruntime.Field{
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
	{Name: "PersonalAccessTokenId", Flag: "personal-access-token-id", Type: "*string", Required: true},
}

var fields_list_access_control_rules = []leanruntime.Field{
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_list_aliases = []leanruntime.Field{
	{Name: "EntityId", Flag: "entity-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_list_availability_configurations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_list_group_members = []leanruntime.Field{
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_list_groups = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "*types.ListGroupsFilters", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_list_groups_for_entity = []leanruntime.Field{
	{Name: "EntityId", Flag: "entity-id", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "*types.ListGroupsForEntityFilters", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_list_impersonation_roles = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_list_mail_domains = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_list_mailbox_export_jobs = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_list_mailbox_permissions = []leanruntime.Field{
	{Name: "EntityId", Flag: "entity-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_list_mobile_device_access_overrides = []leanruntime.Field{
	{Name: "DeviceId", Flag: "device-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: false},
}

var fields_list_mobile_device_access_rules = []leanruntime.Field{
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_list_organizations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_personal_access_tokens = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: false},
}

var fields_list_resource_delegates = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
}

var fields_list_resources = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "*types.ListResourcesFilters", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_users = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "*types.ListUsersFilters", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_put_access_control_rule = []leanruntime.Field{
	{Name: "Actions", Flag: "actions", Type: "[]string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "Effect", Flag: "effect", Type: "types.AccessControlRuleEffect", Required: true},
	{Name: "ImpersonationRoleIds", Flag: "impersonation-role-ids", Type: "[]string", Required: false},
	{Name: "IpRanges", Flag: "ip-ranges", Type: "[]string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NotActions", Flag: "not-actions", Type: "[]string", Required: false},
	{Name: "NotImpersonationRoleIds", Flag: "not-impersonation-role-ids", Type: "[]string", Required: false},
	{Name: "NotIpRanges", Flag: "not-ip-ranges", Type: "[]string", Required: false},
	{Name: "NotUserIds", Flag: "not-user-ids", Type: "[]string", Required: false},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
	{Name: "UserIds", Flag: "user-ids", Type: "[]string", Required: false},
}

var fields_put_email_monitoring_configuration = []leanruntime.Field{
	{Name: "LogGroupArn", Flag: "log-group-arn", Type: "*string", Required: true},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
}

var fields_put_identity_provider_configuration = []leanruntime.Field{
	{Name: "AuthenticationMode", Flag: "authentication-mode", Type: "types.IdentityProviderAuthenticationMode", Required: true},
	{Name: "IdentityCenterConfiguration", Flag: "identity-center-configuration", Type: "*types.IdentityCenterConfiguration", Required: true},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
	{Name: "PersonalAccessTokenConfiguration", Flag: "personal-access-token-configuration", Type: "*types.PersonalAccessTokenConfiguration", Required: true},
}

var fields_put_inbound_dmarc_settings = []leanruntime.Field{
	{Name: "Enforced", Flag: "enforced", Type: "*bool", Required: true},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_put_mailbox_permissions = []leanruntime.Field{
	{Name: "EntityId", Flag: "entity-id", Type: "*string", Required: true},
	{Name: "GranteeId", Flag: "grantee-id", Type: "*string", Required: true},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
	{Name: "PermissionValues", Flag: "permission-values", Type: "[]types.PermissionType", Required: true},
}

var fields_put_mobile_device_access_override = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DeviceId", Flag: "device-id", Type: "*string", Required: true},
	{Name: "Effect", Flag: "effect", Type: "types.MobileDeviceAccessRuleEffect", Required: true},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_put_retention_policy = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "FolderConfigurations", Flag: "folder-configurations", Type: "[]types.FolderConfiguration", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_register_mail_domain = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_register_to_work_mail = []leanruntime.Field{
	{Name: "Email", Flag: "email", Type: "*string", Required: true},
	{Name: "EntityId", Flag: "entity-id", Type: "*string", Required: true},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_reset_password = []leanruntime.Field{
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
	{Name: "Password", Flag: "password", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_start_mailbox_export_job = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EntityId", Flag: "entity-id", Type: "*string", Required: true},
	{Name: "KmsKeyArn", Flag: "kms-key-arn", Type: "*string", Required: true},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "S3BucketName", Flag: "s3-bucket-name", Type: "*string", Required: true},
	{Name: "S3Prefix", Flag: "s3-prefix", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_test_availability_configuration = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: false},
	{Name: "EwsProvider", Flag: "ews-provider", Type: "*types.EwsAvailabilityProvider", Required: false},
	{Name: "LambdaProvider", Flag: "lambda-provider", Type: "*types.LambdaAvailabilityProvider", Required: false},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_availability_configuration = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "EwsProvider", Flag: "ews-provider", Type: "*types.EwsAvailabilityProvider", Required: false},
	{Name: "LambdaProvider", Flag: "lambda-provider", Type: "*types.LambdaAvailabilityProvider", Required: false},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_update_default_mail_domain = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_update_group = []leanruntime.Field{
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: true},
	{Name: "HiddenFromGlobalAddressList", Flag: "hidden-from-global-address-list", Type: "*bool", Required: false},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_update_impersonation_role = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ImpersonationRoleId", Flag: "impersonation-role-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
	{Name: "Rules", Flag: "rules", Type: "[]types.ImpersonationRule", Required: true},
	{Name: "Type", Flag: "type", Type: "types.ImpersonationRoleType", Required: true},
}

var fields_update_mailbox_quota = []leanruntime.Field{
	{Name: "MailboxQuota", Flag: "mailbox-quota", Type: "*int32", Required: true},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_update_mobile_device_access_rule = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DeviceModels", Flag: "device-models", Type: "[]string", Required: false},
	{Name: "DeviceOperatingSystems", Flag: "device-operating-systems", Type: "[]string", Required: false},
	{Name: "DeviceTypes", Flag: "device-types", Type: "[]string", Required: false},
	{Name: "DeviceUserAgents", Flag: "device-user-agents", Type: "[]string", Required: false},
	{Name: "Effect", Flag: "effect", Type: "types.MobileDeviceAccessRuleEffect", Required: true},
	{Name: "MobileDeviceAccessRuleId", Flag: "mobile-device-access-rule-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NotDeviceModels", Flag: "not-device-models", Type: "[]string", Required: false},
	{Name: "NotDeviceOperatingSystems", Flag: "not-device-operating-systems", Type: "[]string", Required: false},
	{Name: "NotDeviceTypes", Flag: "not-device-types", Type: "[]string", Required: false},
	{Name: "NotDeviceUserAgents", Flag: "not-device-user-agents", Type: "[]string", Required: false},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_update_primary_email_address = []leanruntime.Field{
	{Name: "Email", Flag: "email", Type: "*string", Required: true},
	{Name: "EntityId", Flag: "entity-id", Type: "*string", Required: true},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_update_resource = []leanruntime.Field{
	{Name: "BookingOptions", Flag: "booking-options", Type: "*types.BookingOptions", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "HiddenFromGlobalAddressList", Flag: "hidden-from-global-address-list", Type: "*bool", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.ResourceType", Required: false},
}

var fields_update_user = []leanruntime.Field{
	{Name: "City", Flag: "city", Type: "*string", Required: false},
	{Name: "Company", Flag: "company", Type: "*string", Required: false},
	{Name: "Country", Flag: "country", Type: "*string", Required: false},
	{Name: "Department", Flag: "department", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "FirstName", Flag: "first-name", Type: "*string", Required: false},
	{Name: "HiddenFromGlobalAddressList", Flag: "hidden-from-global-address-list", Type: "*bool", Required: false},
	{Name: "IdentityProviderUserId", Flag: "identity-provider-user-id", Type: "*string", Required: false},
	{Name: "Initials", Flag: "initials", Type: "*string", Required: false},
	{Name: "JobTitle", Flag: "job-title", Type: "*string", Required: false},
	{Name: "LastName", Flag: "last-name", Type: "*string", Required: false},
	{Name: "Office", Flag: "office", Type: "*string", Required: false},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
	{Name: "Role", Flag: "role", Type: "types.UserRole", Required: false},
	{Name: "Street", Flag: "street", Type: "*string", Required: false},
	{Name: "Telephone", Flag: "telephone", Type: "*string", Required: false},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
	{Name: "ZipCode", Flag: "zip-code", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-delegate-to-resource": {
			Name:   "associate-delegate-to-resource",
			Fields: fields_associate_delegate_to_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateDelegateToResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_delegate_to_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateDelegateToResource(ctx, input)
			},
		},
		"associate-member-to-group": {
			Name:   "associate-member-to-group",
			Fields: fields_associate_member_to_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateMemberToGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_member_to_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateMemberToGroup(ctx, input)
			},
		},
		"assume-impersonation-role": {
			Name:   "assume-impersonation-role",
			Fields: fields_assume_impersonation_role,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssumeImpersonationRoleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_assume_impersonation_role, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssumeImpersonationRole(ctx, input)
			},
		},
		"cancel-mailbox-export-job": {
			Name:   "cancel-mailbox-export-job",
			Fields: fields_cancel_mailbox_export_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelMailboxExportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_mailbox_export_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelMailboxExportJob(ctx, input)
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
		"create-availability-configuration": {
			Name:   "create-availability-configuration",
			Fields: fields_create_availability_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAvailabilityConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_availability_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAvailabilityConfiguration(ctx, input)
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
		"create-identity-center-application": {
			Name:   "create-identity-center-application",
			Fields: fields_create_identity_center_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateIdentityCenterApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_identity_center_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateIdentityCenterApplication(ctx, input)
			},
		},
		"create-impersonation-role": {
			Name:   "create-impersonation-role",
			Fields: fields_create_impersonation_role,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateImpersonationRoleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_impersonation_role, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateImpersonationRole(ctx, input)
			},
		},
		"create-mobile-device-access-rule": {
			Name:   "create-mobile-device-access-rule",
			Fields: fields_create_mobile_device_access_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMobileDeviceAccessRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_mobile_device_access_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMobileDeviceAccessRule(ctx, input)
			},
		},
		"create-organization": {
			Name:   "create-organization",
			Fields: fields_create_organization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateOrganizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_organization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateOrganization(ctx, input)
			},
		},
		"create-resource": {
			Name:   "create-resource",
			Fields: fields_create_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateResource(ctx, input)
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
		"delete-access-control-rule": {
			Name:   "delete-access-control-rule",
			Fields: fields_delete_access_control_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAccessControlRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_access_control_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAccessControlRule(ctx, input)
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
		"delete-availability-configuration": {
			Name:   "delete-availability-configuration",
			Fields: fields_delete_availability_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAvailabilityConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_availability_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAvailabilityConfiguration(ctx, input)
			},
		},
		"delete-email-monitoring-configuration": {
			Name:   "delete-email-monitoring-configuration",
			Fields: fields_delete_email_monitoring_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEmailMonitoringConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_email_monitoring_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEmailMonitoringConfiguration(ctx, input)
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
		"delete-identity-center-application": {
			Name:   "delete-identity-center-application",
			Fields: fields_delete_identity_center_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIdentityCenterApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_identity_center_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIdentityCenterApplication(ctx, input)
			},
		},
		"delete-identity-provider-configuration": {
			Name:   "delete-identity-provider-configuration",
			Fields: fields_delete_identity_provider_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIdentityProviderConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_identity_provider_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIdentityProviderConfiguration(ctx, input)
			},
		},
		"delete-impersonation-role": {
			Name:   "delete-impersonation-role",
			Fields: fields_delete_impersonation_role,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteImpersonationRoleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_impersonation_role, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteImpersonationRole(ctx, input)
			},
		},
		"delete-mailbox-permissions": {
			Name:   "delete-mailbox-permissions",
			Fields: fields_delete_mailbox_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMailboxPermissionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_mailbox_permissions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMailboxPermissions(ctx, input)
			},
		},
		"delete-mobile-device-access-override": {
			Name:   "delete-mobile-device-access-override",
			Fields: fields_delete_mobile_device_access_override,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMobileDeviceAccessOverrideInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_mobile_device_access_override, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMobileDeviceAccessOverride(ctx, input)
			},
		},
		"delete-mobile-device-access-rule": {
			Name:   "delete-mobile-device-access-rule",
			Fields: fields_delete_mobile_device_access_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMobileDeviceAccessRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_mobile_device_access_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMobileDeviceAccessRule(ctx, input)
			},
		},
		"delete-organization": {
			Name:   "delete-organization",
			Fields: fields_delete_organization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteOrganizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_organization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteOrganization(ctx, input)
			},
		},
		"delete-personal-access-token": {
			Name:   "delete-personal-access-token",
			Fields: fields_delete_personal_access_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePersonalAccessTokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_personal_access_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePersonalAccessToken(ctx, input)
			},
		},
		"delete-resource": {
			Name:   "delete-resource",
			Fields: fields_delete_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteResource(ctx, input)
			},
		},
		"delete-retention-policy": {
			Name:   "delete-retention-policy",
			Fields: fields_delete_retention_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRetentionPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_retention_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRetentionPolicy(ctx, input)
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
		"deregister-from-work-mail": {
			Name:   "deregister-from-work-mail",
			Fields: fields_deregister_from_work_mail,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterFromWorkMailInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_from_work_mail, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterFromWorkMail(ctx, input)
			},
		},
		"deregister-mail-domain": {
			Name:   "deregister-mail-domain",
			Fields: fields_deregister_mail_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterMailDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_mail_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterMailDomain(ctx, input)
			},
		},
		"describe-email-monitoring-configuration": {
			Name:   "describe-email-monitoring-configuration",
			Fields: fields_describe_email_monitoring_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEmailMonitoringConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_email_monitoring_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEmailMonitoringConfiguration(ctx, input)
			},
		},
		"describe-entity": {
			Name:   "describe-entity",
			Fields: fields_describe_entity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEntityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_entity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEntity(ctx, input)
			},
		},
		"describe-group": {
			Name:   "describe-group",
			Fields: fields_describe_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeGroup(ctx, input)
			},
		},
		"describe-identity-provider-configuration": {
			Name:   "describe-identity-provider-configuration",
			Fields: fields_describe_identity_provider_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeIdentityProviderConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_identity_provider_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeIdentityProviderConfiguration(ctx, input)
			},
		},
		"describe-inbound-dmarc-settings": {
			Name:   "describe-inbound-dmarc-settings",
			Fields: fields_describe_inbound_dmarc_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInboundDmarcSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_inbound_dmarc_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeInboundDmarcSettings(ctx, input)
			},
		},
		"describe-mailbox-export-job": {
			Name:   "describe-mailbox-export-job",
			Fields: fields_describe_mailbox_export_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMailboxExportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_mailbox_export_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeMailboxExportJob(ctx, input)
			},
		},
		"describe-organization": {
			Name:   "describe-organization",
			Fields: fields_describe_organization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeOrganizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_organization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeOrganization(ctx, input)
			},
		},
		"describe-resource": {
			Name:   "describe-resource",
			Fields: fields_describe_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeResource(ctx, input)
			},
		},
		"describe-user": {
			Name:   "describe-user",
			Fields: fields_describe_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeUser(ctx, input)
			},
		},
		"disassociate-delegate-from-resource": {
			Name:   "disassociate-delegate-from-resource",
			Fields: fields_disassociate_delegate_from_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateDelegateFromResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_delegate_from_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateDelegateFromResource(ctx, input)
			},
		},
		"disassociate-member-from-group": {
			Name:   "disassociate-member-from-group",
			Fields: fields_disassociate_member_from_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateMemberFromGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_member_from_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateMemberFromGroup(ctx, input)
			},
		},
		"get-access-control-effect": {
			Name:   "get-access-control-effect",
			Fields: fields_get_access_control_effect,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccessControlEffectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_access_control_effect, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccessControlEffect(ctx, input)
			},
		},
		"get-default-retention-policy": {
			Name:   "get-default-retention-policy",
			Fields: fields_get_default_retention_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDefaultRetentionPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_default_retention_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDefaultRetentionPolicy(ctx, input)
			},
		},
		"get-impersonation-role": {
			Name:   "get-impersonation-role",
			Fields: fields_get_impersonation_role,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetImpersonationRoleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_impersonation_role, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetImpersonationRole(ctx, input)
			},
		},
		"get-impersonation-role-effect": {
			Name:   "get-impersonation-role-effect",
			Fields: fields_get_impersonation_role_effect,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetImpersonationRoleEffectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_impersonation_role_effect, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetImpersonationRoleEffect(ctx, input)
			},
		},
		"get-mail-domain": {
			Name:   "get-mail-domain",
			Fields: fields_get_mail_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMailDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_mail_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMailDomain(ctx, input)
			},
		},
		"get-mailbox-details": {
			Name:   "get-mailbox-details",
			Fields: fields_get_mailbox_details,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMailboxDetailsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_mailbox_details, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMailboxDetails(ctx, input)
			},
		},
		"get-mobile-device-access-effect": {
			Name:   "get-mobile-device-access-effect",
			Fields: fields_get_mobile_device_access_effect,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMobileDeviceAccessEffectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_mobile_device_access_effect, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMobileDeviceAccessEffect(ctx, input)
			},
		},
		"get-mobile-device-access-override": {
			Name:   "get-mobile-device-access-override",
			Fields: fields_get_mobile_device_access_override,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMobileDeviceAccessOverrideInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_mobile_device_access_override, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMobileDeviceAccessOverride(ctx, input)
			},
		},
		"get-personal-access-token-metadata": {
			Name:   "get-personal-access-token-metadata",
			Fields: fields_get_personal_access_token_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPersonalAccessTokenMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_personal_access_token_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPersonalAccessTokenMetadata(ctx, input)
			},
		},
		"list-access-control-rules": {
			Name:   "list-access-control-rules",
			Fields: fields_list_access_control_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAccessControlRulesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_access_control_rules, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListAccessControlRules(ctx, input)
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
		"list-availability-configurations": {
			Name:   "list-availability-configurations",
			Fields: fields_list_availability_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAvailabilityConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_availability_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAvailabilityConfigurations(ctx, input)
				}
				var results []*svc.ListAvailabilityConfigurationsOutput
				p := svc.NewListAvailabilityConfigurationsPaginator(client, input)
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
		"list-group-members": {
			Name:   "list-group-members",
			Fields: fields_list_group_members,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGroupMembersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_group_members, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListGroupMembers(ctx, input)
				}
				var results []*svc.ListGroupMembersOutput
				p := svc.NewListGroupMembersPaginator(client, input)
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
		"list-groups-for-entity": {
			Name:   "list-groups-for-entity",
			Fields: fields_list_groups_for_entity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGroupsForEntityInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_groups_for_entity, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListGroupsForEntity(ctx, input)
				}
				var results []*svc.ListGroupsForEntityOutput
				p := svc.NewListGroupsForEntityPaginator(client, input)
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
		"list-impersonation-roles": {
			Name:   "list-impersonation-roles",
			Fields: fields_list_impersonation_roles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListImpersonationRolesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_impersonation_roles, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListImpersonationRoles(ctx, input)
				}
				var results []*svc.ListImpersonationRolesOutput
				p := svc.NewListImpersonationRolesPaginator(client, input)
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
		"list-mail-domains": {
			Name:   "list-mail-domains",
			Fields: fields_list_mail_domains,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMailDomainsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_mail_domains, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMailDomains(ctx, input)
				}
				var results []*svc.ListMailDomainsOutput
				p := svc.NewListMailDomainsPaginator(client, input)
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
		"list-mailbox-export-jobs": {
			Name:   "list-mailbox-export-jobs",
			Fields: fields_list_mailbox_export_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMailboxExportJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_mailbox_export_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMailboxExportJobs(ctx, input)
				}
				var results []*svc.ListMailboxExportJobsOutput
				p := svc.NewListMailboxExportJobsPaginator(client, input)
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
		"list-mailbox-permissions": {
			Name:   "list-mailbox-permissions",
			Fields: fields_list_mailbox_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMailboxPermissionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_mailbox_permissions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMailboxPermissions(ctx, input)
				}
				var results []*svc.ListMailboxPermissionsOutput
				p := svc.NewListMailboxPermissionsPaginator(client, input)
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
		"list-mobile-device-access-overrides": {
			Name:   "list-mobile-device-access-overrides",
			Fields: fields_list_mobile_device_access_overrides,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMobileDeviceAccessOverridesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_mobile_device_access_overrides, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMobileDeviceAccessOverrides(ctx, input)
				}
				var results []*svc.ListMobileDeviceAccessOverridesOutput
				p := svc.NewListMobileDeviceAccessOverridesPaginator(client, input)
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
		"list-mobile-device-access-rules": {
			Name:   "list-mobile-device-access-rules",
			Fields: fields_list_mobile_device_access_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMobileDeviceAccessRulesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_mobile_device_access_rules, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListMobileDeviceAccessRules(ctx, input)
			},
		},
		"list-organizations": {
			Name:   "list-organizations",
			Fields: fields_list_organizations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOrganizationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_organizations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListOrganizations(ctx, input)
				}
				var results []*svc.ListOrganizationsOutput
				p := svc.NewListOrganizationsPaginator(client, input)
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
		"list-personal-access-tokens": {
			Name:   "list-personal-access-tokens",
			Fields: fields_list_personal_access_tokens,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPersonalAccessTokensInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_personal_access_tokens, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPersonalAccessTokens(ctx, input)
				}
				var results []*svc.ListPersonalAccessTokensOutput
				p := svc.NewListPersonalAccessTokensPaginator(client, input)
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
		"list-resource-delegates": {
			Name:   "list-resource-delegates",
			Fields: fields_list_resource_delegates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResourceDelegatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_resource_delegates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResourceDelegates(ctx, input)
				}
				var results []*svc.ListResourceDelegatesOutput
				p := svc.NewListResourceDelegatesPaginator(client, input)
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
		"list-resources": {
			Name:   "list-resources",
			Fields: fields_list_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_resources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResources(ctx, input)
				}
				var results []*svc.ListResourcesOutput
				p := svc.NewListResourcesPaginator(client, input)
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
		"put-access-control-rule": {
			Name:   "put-access-control-rule",
			Fields: fields_put_access_control_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutAccessControlRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_access_control_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutAccessControlRule(ctx, input)
			},
		},
		"put-email-monitoring-configuration": {
			Name:   "put-email-monitoring-configuration",
			Fields: fields_put_email_monitoring_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutEmailMonitoringConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_email_monitoring_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutEmailMonitoringConfiguration(ctx, input)
			},
		},
		"put-identity-provider-configuration": {
			Name:   "put-identity-provider-configuration",
			Fields: fields_put_identity_provider_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutIdentityProviderConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_identity_provider_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutIdentityProviderConfiguration(ctx, input)
			},
		},
		"put-inbound-dmarc-settings": {
			Name:   "put-inbound-dmarc-settings",
			Fields: fields_put_inbound_dmarc_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutInboundDmarcSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_inbound_dmarc_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutInboundDmarcSettings(ctx, input)
			},
		},
		"put-mailbox-permissions": {
			Name:   "put-mailbox-permissions",
			Fields: fields_put_mailbox_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutMailboxPermissionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_mailbox_permissions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutMailboxPermissions(ctx, input)
			},
		},
		"put-mobile-device-access-override": {
			Name:   "put-mobile-device-access-override",
			Fields: fields_put_mobile_device_access_override,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutMobileDeviceAccessOverrideInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_mobile_device_access_override, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutMobileDeviceAccessOverride(ctx, input)
			},
		},
		"put-retention-policy": {
			Name:   "put-retention-policy",
			Fields: fields_put_retention_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutRetentionPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_retention_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutRetentionPolicy(ctx, input)
			},
		},
		"register-mail-domain": {
			Name:   "register-mail-domain",
			Fields: fields_register_mail_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterMailDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_mail_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterMailDomain(ctx, input)
			},
		},
		"register-to-work-mail": {
			Name:   "register-to-work-mail",
			Fields: fields_register_to_work_mail,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterToWorkMailInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_to_work_mail, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterToWorkMail(ctx, input)
			},
		},
		"reset-password": {
			Name:   "reset-password",
			Fields: fields_reset_password,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResetPasswordInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reset_password, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResetPassword(ctx, input)
			},
		},
		"start-mailbox-export-job": {
			Name:   "start-mailbox-export-job",
			Fields: fields_start_mailbox_export_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartMailboxExportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_mailbox_export_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartMailboxExportJob(ctx, input)
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
		"test-availability-configuration": {
			Name:   "test-availability-configuration",
			Fields: fields_test_availability_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TestAvailabilityConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_test_availability_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TestAvailabilityConfiguration(ctx, input)
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
		"update-availability-configuration": {
			Name:   "update-availability-configuration",
			Fields: fields_update_availability_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAvailabilityConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_availability_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAvailabilityConfiguration(ctx, input)
			},
		},
		"update-default-mail-domain": {
			Name:   "update-default-mail-domain",
			Fields: fields_update_default_mail_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDefaultMailDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_default_mail_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDefaultMailDomain(ctx, input)
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
		"update-impersonation-role": {
			Name:   "update-impersonation-role",
			Fields: fields_update_impersonation_role,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateImpersonationRoleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_impersonation_role, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateImpersonationRole(ctx, input)
			},
		},
		"update-mailbox-quota": {
			Name:   "update-mailbox-quota",
			Fields: fields_update_mailbox_quota,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMailboxQuotaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_mailbox_quota, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMailboxQuota(ctx, input)
			},
		},
		"update-mobile-device-access-rule": {
			Name:   "update-mobile-device-access-rule",
			Fields: fields_update_mobile_device_access_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMobileDeviceAccessRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_mobile_device_access_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMobileDeviceAccessRule(ctx, input)
			},
		},
		"update-primary-email-address": {
			Name:   "update-primary-email-address",
			Fields: fields_update_primary_email_address,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePrimaryEmailAddressInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_primary_email_address, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePrimaryEmailAddress(ctx, input)
			},
		},
		"update-resource": {
			Name:   "update-resource",
			Fields: fields_update_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateResource(ctx, input)
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
	}
	if err := leanruntime.Execute("workmail", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
