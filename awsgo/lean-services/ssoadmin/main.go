package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/ssoadmin"
)

var fields_add_region = []leanruntime.Field{
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
	{Name: "RegionName", Flag: "region-name", Type: "*string", Required: true},
}

var fields_attach_customer_managed_policy_reference_to_permission_set = []leanruntime.Field{
	{Name: "CustomerManagedPolicyReference", Flag: "customer-managed-policy-reference", Type: "*types.CustomerManagedPolicyReference", Required: true},
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
	{Name: "PermissionSetArn", Flag: "permission-set-arn", Type: "*string", Required: true},
}

var fields_attach_managed_policy_to_permission_set = []leanruntime.Field{
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
	{Name: "ManagedPolicyArn", Flag: "managed-policy-arn", Type: "*string", Required: true},
	{Name: "PermissionSetArn", Flag: "permission-set-arn", Type: "*string", Required: true},
}

var fields_create_account_assignment = []leanruntime.Field{
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
	{Name: "PermissionSetArn", Flag: "permission-set-arn", Type: "*string", Required: true},
	{Name: "PrincipalId", Flag: "principal-id", Type: "*string", Required: true},
	{Name: "PrincipalType", Flag: "principal-type", Type: "types.PrincipalType", Required: true},
	{Name: "TargetId", Flag: "target-id", Type: "*string", Required: true},
	{Name: "TargetType", Flag: "target-type", Type: "types.TargetType", Required: true},
}

var fields_create_application = []leanruntime.Field{
	{Name: "ApplicationProviderArn", Flag: "application-provider-arn", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "PortalOptions", Flag: "portal-options", Type: "*types.PortalOptions", Required: false},
	{Name: "Status", Flag: "status", Type: "types.ApplicationStatus", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_application_assignment = []leanruntime.Field{
	{Name: "ApplicationArn", Flag: "application-arn", Type: "*string", Required: true},
	{Name: "PrincipalId", Flag: "principal-id", Type: "*string", Required: true},
	{Name: "PrincipalType", Flag: "principal-type", Type: "types.PrincipalType", Required: true},
}

var fields_create_instance = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_instance_access_control_attribute_configuration = []leanruntime.Field{
	{Name: "InstanceAccessControlAttributeConfiguration", Flag: "instance-access-control-attribute-configuration", Type: "*types.InstanceAccessControlAttributeConfiguration", Required: true},
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
}

var fields_create_permission_set = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RelayState", Flag: "relay-state", Type: "*string", Required: false},
	{Name: "SessionDuration", Flag: "session-duration", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_trusted_token_issuer = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TrustedTokenIssuerConfiguration", Flag: "trusted-token-issuer-configuration", Type: "types.TrustedTokenIssuerConfiguration", Required: true},
	{Name: "TrustedTokenIssuerType", Flag: "trusted-token-issuer-type", Type: "types.TrustedTokenIssuerType", Required: true},
}

var fields_delete_account_assignment = []leanruntime.Field{
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
	{Name: "PermissionSetArn", Flag: "permission-set-arn", Type: "*string", Required: true},
	{Name: "PrincipalId", Flag: "principal-id", Type: "*string", Required: true},
	{Name: "PrincipalType", Flag: "principal-type", Type: "types.PrincipalType", Required: true},
	{Name: "TargetId", Flag: "target-id", Type: "*string", Required: true},
	{Name: "TargetType", Flag: "target-type", Type: "types.TargetType", Required: true},
}

var fields_delete_application = []leanruntime.Field{
	{Name: "ApplicationArn", Flag: "application-arn", Type: "*string", Required: true},
}

var fields_delete_application_access_scope = []leanruntime.Field{
	{Name: "ApplicationArn", Flag: "application-arn", Type: "*string", Required: true},
	{Name: "Scope", Flag: "scope", Type: "*string", Required: true},
}

var fields_delete_application_assignment = []leanruntime.Field{
	{Name: "ApplicationArn", Flag: "application-arn", Type: "*string", Required: true},
	{Name: "PrincipalId", Flag: "principal-id", Type: "*string", Required: true},
	{Name: "PrincipalType", Flag: "principal-type", Type: "types.PrincipalType", Required: true},
}

var fields_delete_application_authentication_method = []leanruntime.Field{
	{Name: "ApplicationArn", Flag: "application-arn", Type: "*string", Required: true},
	{Name: "AuthenticationMethodType", Flag: "authentication-method-type", Type: "types.AuthenticationMethodType", Required: true},
}

var fields_delete_application_grant = []leanruntime.Field{
	{Name: "ApplicationArn", Flag: "application-arn", Type: "*string", Required: true},
	{Name: "GrantType", Flag: "grant-type", Type: "types.GrantType", Required: true},
}

var fields_delete_inline_policy_from_permission_set = []leanruntime.Field{
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
	{Name: "PermissionSetArn", Flag: "permission-set-arn", Type: "*string", Required: true},
}

var fields_delete_instance = []leanruntime.Field{
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
}

var fields_delete_instance_access_control_attribute_configuration = []leanruntime.Field{
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
}

var fields_delete_permission_set = []leanruntime.Field{
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
	{Name: "PermissionSetArn", Flag: "permission-set-arn", Type: "*string", Required: true},
}

var fields_delete_permissions_boundary_from_permission_set = []leanruntime.Field{
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
	{Name: "PermissionSetArn", Flag: "permission-set-arn", Type: "*string", Required: true},
}

var fields_delete_trusted_token_issuer = []leanruntime.Field{
	{Name: "TrustedTokenIssuerArn", Flag: "trusted-token-issuer-arn", Type: "*string", Required: true},
}

var fields_describe_account_assignment_creation_status = []leanruntime.Field{
	{Name: "AccountAssignmentCreationRequestId", Flag: "account-assignment-creation-request-id", Type: "*string", Required: true},
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
}

var fields_describe_account_assignment_deletion_status = []leanruntime.Field{
	{Name: "AccountAssignmentDeletionRequestId", Flag: "account-assignment-deletion-request-id", Type: "*string", Required: true},
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
}

var fields_describe_application = []leanruntime.Field{
	{Name: "ApplicationArn", Flag: "application-arn", Type: "*string", Required: true},
}

var fields_describe_application_assignment = []leanruntime.Field{
	{Name: "ApplicationArn", Flag: "application-arn", Type: "*string", Required: true},
	{Name: "PrincipalId", Flag: "principal-id", Type: "*string", Required: true},
	{Name: "PrincipalType", Flag: "principal-type", Type: "types.PrincipalType", Required: true},
}

var fields_describe_application_provider = []leanruntime.Field{
	{Name: "ApplicationProviderArn", Flag: "application-provider-arn", Type: "*string", Required: true},
}

var fields_describe_instance = []leanruntime.Field{
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
}

var fields_describe_instance_access_control_attribute_configuration = []leanruntime.Field{
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
}

var fields_describe_permission_set = []leanruntime.Field{
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
	{Name: "PermissionSetArn", Flag: "permission-set-arn", Type: "*string", Required: true},
}

var fields_describe_permission_set_provisioning_status = []leanruntime.Field{
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
	{Name: "ProvisionPermissionSetRequestId", Flag: "provision-permission-set-request-id", Type: "*string", Required: true},
}

var fields_describe_region = []leanruntime.Field{
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
	{Name: "RegionName", Flag: "region-name", Type: "*string", Required: true},
}

var fields_describe_trusted_token_issuer = []leanruntime.Field{
	{Name: "TrustedTokenIssuerArn", Flag: "trusted-token-issuer-arn", Type: "*string", Required: true},
}

var fields_detach_customer_managed_policy_reference_from_permission_set = []leanruntime.Field{
	{Name: "CustomerManagedPolicyReference", Flag: "customer-managed-policy-reference", Type: "*types.CustomerManagedPolicyReference", Required: true},
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
	{Name: "PermissionSetArn", Flag: "permission-set-arn", Type: "*string", Required: true},
}

var fields_detach_managed_policy_from_permission_set = []leanruntime.Field{
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
	{Name: "ManagedPolicyArn", Flag: "managed-policy-arn", Type: "*string", Required: true},
	{Name: "PermissionSetArn", Flag: "permission-set-arn", Type: "*string", Required: true},
}

var fields_get_application_access_scope = []leanruntime.Field{
	{Name: "ApplicationArn", Flag: "application-arn", Type: "*string", Required: true},
	{Name: "Scope", Flag: "scope", Type: "*string", Required: true},
}

var fields_get_application_assignment_configuration = []leanruntime.Field{
	{Name: "ApplicationArn", Flag: "application-arn", Type: "*string", Required: true},
}

var fields_get_application_authentication_method = []leanruntime.Field{
	{Name: "ApplicationArn", Flag: "application-arn", Type: "*string", Required: true},
	{Name: "AuthenticationMethodType", Flag: "authentication-method-type", Type: "types.AuthenticationMethodType", Required: true},
}

var fields_get_application_grant = []leanruntime.Field{
	{Name: "ApplicationArn", Flag: "application-arn", Type: "*string", Required: true},
	{Name: "GrantType", Flag: "grant-type", Type: "types.GrantType", Required: true},
}

var fields_get_application_session_configuration = []leanruntime.Field{
	{Name: "ApplicationArn", Flag: "application-arn", Type: "*string", Required: true},
}

var fields_get_inline_policy_for_permission_set = []leanruntime.Field{
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
	{Name: "PermissionSetArn", Flag: "permission-set-arn", Type: "*string", Required: true},
}

var fields_get_permissions_boundary_for_permission_set = []leanruntime.Field{
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
	{Name: "PermissionSetArn", Flag: "permission-set-arn", Type: "*string", Required: true},
}

var fields_list_account_assignment_creation_status = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.OperationStatusFilter", Required: false},
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_account_assignment_deletion_status = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.OperationStatusFilter", Required: false},
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_account_assignments = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PermissionSetArn", Flag: "permission-set-arn", Type: "*string", Required: true},
}

var fields_list_account_assignments_for_principal = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.ListAccountAssignmentsFilter", Required: false},
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PrincipalId", Flag: "principal-id", Type: "*string", Required: true},
	{Name: "PrincipalType", Flag: "principal-type", Type: "types.PrincipalType", Required: true},
}

var fields_list_accounts_for_provisioned_permission_set = []leanruntime.Field{
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PermissionSetArn", Flag: "permission-set-arn", Type: "*string", Required: true},
	{Name: "ProvisioningStatus", Flag: "provisioning-status", Type: "types.ProvisioningStatus", Required: false},
}

var fields_list_application_access_scopes = []leanruntime.Field{
	{Name: "ApplicationArn", Flag: "application-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_application_assignments = []leanruntime.Field{
	{Name: "ApplicationArn", Flag: "application-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_application_assignments_for_principal = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.ListApplicationAssignmentsFilter", Required: false},
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PrincipalId", Flag: "principal-id", Type: "*string", Required: true},
	{Name: "PrincipalType", Flag: "principal-type", Type: "types.PrincipalType", Required: true},
}

var fields_list_application_authentication_methods = []leanruntime.Field{
	{Name: "ApplicationArn", Flag: "application-arn", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_application_grants = []leanruntime.Field{
	{Name: "ApplicationArn", Flag: "application-arn", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_application_providers = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_applications = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.ListApplicationsFilter", Required: false},
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_customer_managed_policy_references_in_permission_set = []leanruntime.Field{
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PermissionSetArn", Flag: "permission-set-arn", Type: "*string", Required: true},
}

var fields_list_instances = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_managed_policies_in_permission_set = []leanruntime.Field{
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PermissionSetArn", Flag: "permission-set-arn", Type: "*string", Required: true},
}

var fields_list_permission_set_provisioning_status = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.OperationStatusFilter", Required: false},
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_permission_sets = []leanruntime.Field{
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_permission_sets_provisioned_to_account = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProvisioningStatus", Flag: "provisioning-status", Type: "types.ProvisioningStatus", Required: false},
}

var fields_list_regions = []leanruntime.Field{
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_trusted_token_issuers = []leanruntime.Field{
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_provision_permission_set = []leanruntime.Field{
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
	{Name: "PermissionSetArn", Flag: "permission-set-arn", Type: "*string", Required: true},
	{Name: "TargetId", Flag: "target-id", Type: "*string", Required: false},
	{Name: "TargetType", Flag: "target-type", Type: "types.ProvisionTargetType", Required: true},
}

var fields_put_application_access_scope = []leanruntime.Field{
	{Name: "ApplicationArn", Flag: "application-arn", Type: "*string", Required: true},
	{Name: "AuthorizedTargets", Flag: "authorized-targets", Type: "[]string", Required: false},
	{Name: "Scope", Flag: "scope", Type: "*string", Required: true},
}

var fields_put_application_assignment_configuration = []leanruntime.Field{
	{Name: "ApplicationArn", Flag: "application-arn", Type: "*string", Required: true},
	{Name: "AssignmentRequired", Flag: "assignment-required", Type: "*bool", Required: true},
}

var fields_put_application_authentication_method = []leanruntime.Field{
	{Name: "ApplicationArn", Flag: "application-arn", Type: "*string", Required: true},
	{Name: "AuthenticationMethod", Flag: "authentication-method", Type: "types.AuthenticationMethod", Required: true},
	{Name: "AuthenticationMethodType", Flag: "authentication-method-type", Type: "types.AuthenticationMethodType", Required: true},
}

var fields_put_application_grant = []leanruntime.Field{
	{Name: "ApplicationArn", Flag: "application-arn", Type: "*string", Required: true},
	{Name: "Grant", Flag: "grant", Type: "types.Grant", Required: true},
	{Name: "GrantType", Flag: "grant-type", Type: "types.GrantType", Required: true},
}

var fields_put_application_session_configuration = []leanruntime.Field{
	{Name: "ApplicationArn", Flag: "application-arn", Type: "*string", Required: true},
	{Name: "UserBackgroundSessionApplicationStatus", Flag: "user-background-session-application-status", Type: "types.UserBackgroundSessionApplicationStatus", Required: false},
}

var fields_put_inline_policy_to_permission_set = []leanruntime.Field{
	{Name: "InlinePolicy", Flag: "inline-policy", Type: "*string", Required: true},
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
	{Name: "PermissionSetArn", Flag: "permission-set-arn", Type: "*string", Required: true},
}

var fields_put_permissions_boundary_to_permission_set = []leanruntime.Field{
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
	{Name: "PermissionSetArn", Flag: "permission-set-arn", Type: "*string", Required: true},
	{Name: "PermissionsBoundary", Flag: "permissions-boundary", Type: "*types.PermissionsBoundary", Required: true},
}

var fields_remove_region = []leanruntime.Field{
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
	{Name: "RegionName", Flag: "region-name", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_application = []leanruntime.Field{
	{Name: "ApplicationArn", Flag: "application-arn", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "PortalOptions", Flag: "portal-options", Type: "*types.UpdateApplicationPortalOptions", Required: false},
	{Name: "Status", Flag: "status", Type: "types.ApplicationStatus", Required: false},
}

var fields_update_instance = []leanruntime.Field{
	{Name: "EncryptionConfiguration", Flag: "encryption-configuration", Type: "*types.EncryptionConfiguration", Required: false},
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_instance_access_control_attribute_configuration = []leanruntime.Field{
	{Name: "InstanceAccessControlAttributeConfiguration", Flag: "instance-access-control-attribute-configuration", Type: "*types.InstanceAccessControlAttributeConfiguration", Required: true},
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
}

var fields_update_permission_set = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
	{Name: "PermissionSetArn", Flag: "permission-set-arn", Type: "*string", Required: true},
	{Name: "RelayState", Flag: "relay-state", Type: "*string", Required: false},
	{Name: "SessionDuration", Flag: "session-duration", Type: "*string", Required: false},
}

var fields_update_trusted_token_issuer = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "TrustedTokenIssuerArn", Flag: "trusted-token-issuer-arn", Type: "*string", Required: true},
	{Name: "TrustedTokenIssuerConfiguration", Flag: "trusted-token-issuer-configuration", Type: "types.TrustedTokenIssuerUpdateConfiguration", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"add-region": {
			Name:   "add-region",
			Fields: fields_add_region,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddRegionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_region, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddRegion(ctx, input)
			},
		},
		"attach-customer-managed-policy-reference-to-permission-set": {
			Name:   "attach-customer-managed-policy-reference-to-permission-set",
			Fields: fields_attach_customer_managed_policy_reference_to_permission_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AttachCustomerManagedPolicyReferenceToPermissionSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_attach_customer_managed_policy_reference_to_permission_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AttachCustomerManagedPolicyReferenceToPermissionSet(ctx, input)
			},
		},
		"attach-managed-policy-to-permission-set": {
			Name:   "attach-managed-policy-to-permission-set",
			Fields: fields_attach_managed_policy_to_permission_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AttachManagedPolicyToPermissionSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_attach_managed_policy_to_permission_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AttachManagedPolicyToPermissionSet(ctx, input)
			},
		},
		"create-account-assignment": {
			Name:   "create-account-assignment",
			Fields: fields_create_account_assignment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAccountAssignmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_account_assignment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAccountAssignment(ctx, input)
			},
		},
		"create-application": {
			Name:   "create-application",
			Fields: fields_create_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateApplication(ctx, input)
			},
		},
		"create-application-assignment": {
			Name:   "create-application-assignment",
			Fields: fields_create_application_assignment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateApplicationAssignmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_application_assignment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateApplicationAssignment(ctx, input)
			},
		},
		"create-instance": {
			Name:   "create-instance",
			Fields: fields_create_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateInstance(ctx, input)
			},
		},
		"create-instance-access-control-attribute-configuration": {
			Name:   "create-instance-access-control-attribute-configuration",
			Fields: fields_create_instance_access_control_attribute_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateInstanceAccessControlAttributeConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_instance_access_control_attribute_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateInstanceAccessControlAttributeConfiguration(ctx, input)
			},
		},
		"create-permission-set": {
			Name:   "create-permission-set",
			Fields: fields_create_permission_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePermissionSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_permission_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePermissionSet(ctx, input)
			},
		},
		"create-trusted-token-issuer": {
			Name:   "create-trusted-token-issuer",
			Fields: fields_create_trusted_token_issuer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTrustedTokenIssuerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_trusted_token_issuer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTrustedTokenIssuer(ctx, input)
			},
		},
		"delete-account-assignment": {
			Name:   "delete-account-assignment",
			Fields: fields_delete_account_assignment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAccountAssignmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_account_assignment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAccountAssignment(ctx, input)
			},
		},
		"delete-application": {
			Name:   "delete-application",
			Fields: fields_delete_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteApplication(ctx, input)
			},
		},
		"delete-application-access-scope": {
			Name:   "delete-application-access-scope",
			Fields: fields_delete_application_access_scope,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteApplicationAccessScopeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_application_access_scope, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteApplicationAccessScope(ctx, input)
			},
		},
		"delete-application-assignment": {
			Name:   "delete-application-assignment",
			Fields: fields_delete_application_assignment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteApplicationAssignmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_application_assignment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteApplicationAssignment(ctx, input)
			},
		},
		"delete-application-authentication-method": {
			Name:   "delete-application-authentication-method",
			Fields: fields_delete_application_authentication_method,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteApplicationAuthenticationMethodInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_application_authentication_method, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteApplicationAuthenticationMethod(ctx, input)
			},
		},
		"delete-application-grant": {
			Name:   "delete-application-grant",
			Fields: fields_delete_application_grant,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteApplicationGrantInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_application_grant, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteApplicationGrant(ctx, input)
			},
		},
		"delete-inline-policy-from-permission-set": {
			Name:   "delete-inline-policy-from-permission-set",
			Fields: fields_delete_inline_policy_from_permission_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteInlinePolicyFromPermissionSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_inline_policy_from_permission_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteInlinePolicyFromPermissionSet(ctx, input)
			},
		},
		"delete-instance": {
			Name:   "delete-instance",
			Fields: fields_delete_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteInstance(ctx, input)
			},
		},
		"delete-instance-access-control-attribute-configuration": {
			Name:   "delete-instance-access-control-attribute-configuration",
			Fields: fields_delete_instance_access_control_attribute_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteInstanceAccessControlAttributeConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_instance_access_control_attribute_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteInstanceAccessControlAttributeConfiguration(ctx, input)
			},
		},
		"delete-permission-set": {
			Name:   "delete-permission-set",
			Fields: fields_delete_permission_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePermissionSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_permission_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePermissionSet(ctx, input)
			},
		},
		"delete-permissions-boundary-from-permission-set": {
			Name:   "delete-permissions-boundary-from-permission-set",
			Fields: fields_delete_permissions_boundary_from_permission_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePermissionsBoundaryFromPermissionSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_permissions_boundary_from_permission_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePermissionsBoundaryFromPermissionSet(ctx, input)
			},
		},
		"delete-trusted-token-issuer": {
			Name:   "delete-trusted-token-issuer",
			Fields: fields_delete_trusted_token_issuer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTrustedTokenIssuerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_trusted_token_issuer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTrustedTokenIssuer(ctx, input)
			},
		},
		"describe-account-assignment-creation-status": {
			Name:   "describe-account-assignment-creation-status",
			Fields: fields_describe_account_assignment_creation_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAccountAssignmentCreationStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_account_assignment_creation_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAccountAssignmentCreationStatus(ctx, input)
			},
		},
		"describe-account-assignment-deletion-status": {
			Name:   "describe-account-assignment-deletion-status",
			Fields: fields_describe_account_assignment_deletion_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAccountAssignmentDeletionStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_account_assignment_deletion_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAccountAssignmentDeletionStatus(ctx, input)
			},
		},
		"describe-application": {
			Name:   "describe-application",
			Fields: fields_describe_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeApplication(ctx, input)
			},
		},
		"describe-application-assignment": {
			Name:   "describe-application-assignment",
			Fields: fields_describe_application_assignment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeApplicationAssignmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_application_assignment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeApplicationAssignment(ctx, input)
			},
		},
		"describe-application-provider": {
			Name:   "describe-application-provider",
			Fields: fields_describe_application_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeApplicationProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_application_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeApplicationProvider(ctx, input)
			},
		},
		"describe-instance": {
			Name:   "describe-instance",
			Fields: fields_describe_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeInstance(ctx, input)
			},
		},
		"describe-instance-access-control-attribute-configuration": {
			Name:   "describe-instance-access-control-attribute-configuration",
			Fields: fields_describe_instance_access_control_attribute_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInstanceAccessControlAttributeConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_instance_access_control_attribute_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeInstanceAccessControlAttributeConfiguration(ctx, input)
			},
		},
		"describe-permission-set": {
			Name:   "describe-permission-set",
			Fields: fields_describe_permission_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePermissionSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_permission_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribePermissionSet(ctx, input)
			},
		},
		"describe-permission-set-provisioning-status": {
			Name:   "describe-permission-set-provisioning-status",
			Fields: fields_describe_permission_set_provisioning_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePermissionSetProvisioningStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_permission_set_provisioning_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribePermissionSetProvisioningStatus(ctx, input)
			},
		},
		"describe-region": {
			Name:   "describe-region",
			Fields: fields_describe_region,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRegionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_region, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeRegion(ctx, input)
			},
		},
		"describe-trusted-token-issuer": {
			Name:   "describe-trusted-token-issuer",
			Fields: fields_describe_trusted_token_issuer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTrustedTokenIssuerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_trusted_token_issuer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTrustedTokenIssuer(ctx, input)
			},
		},
		"detach-customer-managed-policy-reference-from-permission-set": {
			Name:   "detach-customer-managed-policy-reference-from-permission-set",
			Fields: fields_detach_customer_managed_policy_reference_from_permission_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetachCustomerManagedPolicyReferenceFromPermissionSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detach_customer_managed_policy_reference_from_permission_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetachCustomerManagedPolicyReferenceFromPermissionSet(ctx, input)
			},
		},
		"detach-managed-policy-from-permission-set": {
			Name:   "detach-managed-policy-from-permission-set",
			Fields: fields_detach_managed_policy_from_permission_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetachManagedPolicyFromPermissionSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detach_managed_policy_from_permission_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetachManagedPolicyFromPermissionSet(ctx, input)
			},
		},
		"get-application-access-scope": {
			Name:   "get-application-access-scope",
			Fields: fields_get_application_access_scope,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetApplicationAccessScopeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_application_access_scope, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetApplicationAccessScope(ctx, input)
			},
		},
		"get-application-assignment-configuration": {
			Name:   "get-application-assignment-configuration",
			Fields: fields_get_application_assignment_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetApplicationAssignmentConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_application_assignment_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetApplicationAssignmentConfiguration(ctx, input)
			},
		},
		"get-application-authentication-method": {
			Name:   "get-application-authentication-method",
			Fields: fields_get_application_authentication_method,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetApplicationAuthenticationMethodInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_application_authentication_method, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetApplicationAuthenticationMethod(ctx, input)
			},
		},
		"get-application-grant": {
			Name:   "get-application-grant",
			Fields: fields_get_application_grant,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetApplicationGrantInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_application_grant, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetApplicationGrant(ctx, input)
			},
		},
		"get-application-session-configuration": {
			Name:   "get-application-session-configuration",
			Fields: fields_get_application_session_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetApplicationSessionConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_application_session_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetApplicationSessionConfiguration(ctx, input)
			},
		},
		"get-inline-policy-for-permission-set": {
			Name:   "get-inline-policy-for-permission-set",
			Fields: fields_get_inline_policy_for_permission_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetInlinePolicyForPermissionSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_inline_policy_for_permission_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetInlinePolicyForPermissionSet(ctx, input)
			},
		},
		"get-permissions-boundary-for-permission-set": {
			Name:   "get-permissions-boundary-for-permission-set",
			Fields: fields_get_permissions_boundary_for_permission_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPermissionsBoundaryForPermissionSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_permissions_boundary_for_permission_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPermissionsBoundaryForPermissionSet(ctx, input)
			},
		},
		"list-account-assignment-creation-status": {
			Name:   "list-account-assignment-creation-status",
			Fields: fields_list_account_assignment_creation_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAccountAssignmentCreationStatusInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_account_assignment_creation_status, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAccountAssignmentCreationStatus(ctx, input)
				}
				var results []*svc.ListAccountAssignmentCreationStatusOutput
				p := svc.NewListAccountAssignmentCreationStatusPaginator(client, input)
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
		"list-account-assignment-deletion-status": {
			Name:   "list-account-assignment-deletion-status",
			Fields: fields_list_account_assignment_deletion_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAccountAssignmentDeletionStatusInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_account_assignment_deletion_status, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAccountAssignmentDeletionStatus(ctx, input)
				}
				var results []*svc.ListAccountAssignmentDeletionStatusOutput
				p := svc.NewListAccountAssignmentDeletionStatusPaginator(client, input)
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
		"list-account-assignments": {
			Name:   "list-account-assignments",
			Fields: fields_list_account_assignments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAccountAssignmentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_account_assignments, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAccountAssignments(ctx, input)
				}
				var results []*svc.ListAccountAssignmentsOutput
				p := svc.NewListAccountAssignmentsPaginator(client, input)
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
		"list-account-assignments-for-principal": {
			Name:   "list-account-assignments-for-principal",
			Fields: fields_list_account_assignments_for_principal,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAccountAssignmentsForPrincipalInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_account_assignments_for_principal, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAccountAssignmentsForPrincipal(ctx, input)
				}
				var results []*svc.ListAccountAssignmentsForPrincipalOutput
				p := svc.NewListAccountAssignmentsForPrincipalPaginator(client, input)
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
		"list-accounts-for-provisioned-permission-set": {
			Name:   "list-accounts-for-provisioned-permission-set",
			Fields: fields_list_accounts_for_provisioned_permission_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAccountsForProvisionedPermissionSetInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_accounts_for_provisioned_permission_set, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAccountsForProvisionedPermissionSet(ctx, input)
				}
				var results []*svc.ListAccountsForProvisionedPermissionSetOutput
				p := svc.NewListAccountsForProvisionedPermissionSetPaginator(client, input)
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
		"list-application-access-scopes": {
			Name:   "list-application-access-scopes",
			Fields: fields_list_application_access_scopes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListApplicationAccessScopesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_application_access_scopes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListApplicationAccessScopes(ctx, input)
				}
				var results []*svc.ListApplicationAccessScopesOutput
				p := svc.NewListApplicationAccessScopesPaginator(client, input)
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
		"list-application-assignments": {
			Name:   "list-application-assignments",
			Fields: fields_list_application_assignments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListApplicationAssignmentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_application_assignments, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListApplicationAssignments(ctx, input)
				}
				var results []*svc.ListApplicationAssignmentsOutput
				p := svc.NewListApplicationAssignmentsPaginator(client, input)
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
		"list-application-assignments-for-principal": {
			Name:   "list-application-assignments-for-principal",
			Fields: fields_list_application_assignments_for_principal,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListApplicationAssignmentsForPrincipalInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_application_assignments_for_principal, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListApplicationAssignmentsForPrincipal(ctx, input)
				}
				var results []*svc.ListApplicationAssignmentsForPrincipalOutput
				p := svc.NewListApplicationAssignmentsForPrincipalPaginator(client, input)
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
		"list-application-authentication-methods": {
			Name:   "list-application-authentication-methods",
			Fields: fields_list_application_authentication_methods,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListApplicationAuthenticationMethodsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_application_authentication_methods, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListApplicationAuthenticationMethods(ctx, input)
				}
				var results []*svc.ListApplicationAuthenticationMethodsOutput
				p := svc.NewListApplicationAuthenticationMethodsPaginator(client, input)
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
		"list-application-grants": {
			Name:   "list-application-grants",
			Fields: fields_list_application_grants,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListApplicationGrantsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_application_grants, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListApplicationGrants(ctx, input)
				}
				var results []*svc.ListApplicationGrantsOutput
				p := svc.NewListApplicationGrantsPaginator(client, input)
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
		"list-application-providers": {
			Name:   "list-application-providers",
			Fields: fields_list_application_providers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListApplicationProvidersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_application_providers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListApplicationProviders(ctx, input)
				}
				var results []*svc.ListApplicationProvidersOutput
				p := svc.NewListApplicationProvidersPaginator(client, input)
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
		"list-applications": {
			Name:   "list-applications",
			Fields: fields_list_applications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListApplicationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_applications, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListApplications(ctx, input)
				}
				var results []*svc.ListApplicationsOutput
				p := svc.NewListApplicationsPaginator(client, input)
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
		"list-customer-managed-policy-references-in-permission-set": {
			Name:   "list-customer-managed-policy-references-in-permission-set",
			Fields: fields_list_customer_managed_policy_references_in_permission_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCustomerManagedPolicyReferencesInPermissionSetInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_customer_managed_policy_references_in_permission_set, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCustomerManagedPolicyReferencesInPermissionSet(ctx, input)
				}
				var results []*svc.ListCustomerManagedPolicyReferencesInPermissionSetOutput
				p := svc.NewListCustomerManagedPolicyReferencesInPermissionSetPaginator(client, input)
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
		"list-instances": {
			Name:   "list-instances",
			Fields: fields_list_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInstancesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_instances, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListInstances(ctx, input)
				}
				var results []*svc.ListInstancesOutput
				p := svc.NewListInstancesPaginator(client, input)
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
		"list-managed-policies-in-permission-set": {
			Name:   "list-managed-policies-in-permission-set",
			Fields: fields_list_managed_policies_in_permission_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListManagedPoliciesInPermissionSetInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_managed_policies_in_permission_set, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListManagedPoliciesInPermissionSet(ctx, input)
				}
				var results []*svc.ListManagedPoliciesInPermissionSetOutput
				p := svc.NewListManagedPoliciesInPermissionSetPaginator(client, input)
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
		"list-permission-set-provisioning-status": {
			Name:   "list-permission-set-provisioning-status",
			Fields: fields_list_permission_set_provisioning_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPermissionSetProvisioningStatusInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_permission_set_provisioning_status, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPermissionSetProvisioningStatus(ctx, input)
				}
				var results []*svc.ListPermissionSetProvisioningStatusOutput
				p := svc.NewListPermissionSetProvisioningStatusPaginator(client, input)
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
		"list-permission-sets": {
			Name:   "list-permission-sets",
			Fields: fields_list_permission_sets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPermissionSetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_permission_sets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPermissionSets(ctx, input)
				}
				var results []*svc.ListPermissionSetsOutput
				p := svc.NewListPermissionSetsPaginator(client, input)
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
		"list-permission-sets-provisioned-to-account": {
			Name:   "list-permission-sets-provisioned-to-account",
			Fields: fields_list_permission_sets_provisioned_to_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPermissionSetsProvisionedToAccountInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_permission_sets_provisioned_to_account, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPermissionSetsProvisionedToAccount(ctx, input)
				}
				var results []*svc.ListPermissionSetsProvisionedToAccountOutput
				p := svc.NewListPermissionSetsProvisionedToAccountPaginator(client, input)
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
		"list-regions": {
			Name:   "list-regions",
			Fields: fields_list_regions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRegionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_regions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRegions(ctx, input)
				}
				var results []*svc.ListRegionsOutput
				p := svc.NewListRegionsPaginator(client, input)
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
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_tags_for_resource, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTagsForResource(ctx, input)
				}
				var results []*svc.ListTagsForResourceOutput
				p := svc.NewListTagsForResourcePaginator(client, input)
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
		"list-trusted-token-issuers": {
			Name:   "list-trusted-token-issuers",
			Fields: fields_list_trusted_token_issuers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTrustedTokenIssuersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_trusted_token_issuers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTrustedTokenIssuers(ctx, input)
				}
				var results []*svc.ListTrustedTokenIssuersOutput
				p := svc.NewListTrustedTokenIssuersPaginator(client, input)
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
		"provision-permission-set": {
			Name:   "provision-permission-set",
			Fields: fields_provision_permission_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ProvisionPermissionSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_provision_permission_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ProvisionPermissionSet(ctx, input)
			},
		},
		"put-application-access-scope": {
			Name:   "put-application-access-scope",
			Fields: fields_put_application_access_scope,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutApplicationAccessScopeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_application_access_scope, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutApplicationAccessScope(ctx, input)
			},
		},
		"put-application-assignment-configuration": {
			Name:   "put-application-assignment-configuration",
			Fields: fields_put_application_assignment_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutApplicationAssignmentConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_application_assignment_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutApplicationAssignmentConfiguration(ctx, input)
			},
		},
		"put-application-authentication-method": {
			Name:   "put-application-authentication-method",
			Fields: fields_put_application_authentication_method,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutApplicationAuthenticationMethodInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_application_authentication_method, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutApplicationAuthenticationMethod(ctx, input)
			},
		},
		"put-application-grant": {
			Name:   "put-application-grant",
			Fields: fields_put_application_grant,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutApplicationGrantInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_application_grant, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutApplicationGrant(ctx, input)
			},
		},
		"put-application-session-configuration": {
			Name:   "put-application-session-configuration",
			Fields: fields_put_application_session_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutApplicationSessionConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_application_session_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutApplicationSessionConfiguration(ctx, input)
			},
		},
		"put-inline-policy-to-permission-set": {
			Name:   "put-inline-policy-to-permission-set",
			Fields: fields_put_inline_policy_to_permission_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutInlinePolicyToPermissionSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_inline_policy_to_permission_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutInlinePolicyToPermissionSet(ctx, input)
			},
		},
		"put-permissions-boundary-to-permission-set": {
			Name:   "put-permissions-boundary-to-permission-set",
			Fields: fields_put_permissions_boundary_to_permission_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutPermissionsBoundaryToPermissionSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_permissions_boundary_to_permission_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutPermissionsBoundaryToPermissionSet(ctx, input)
			},
		},
		"remove-region": {
			Name:   "remove-region",
			Fields: fields_remove_region,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveRegionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_region, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveRegion(ctx, input)
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
		"update-application": {
			Name:   "update-application",
			Fields: fields_update_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateApplication(ctx, input)
			},
		},
		"update-instance": {
			Name:   "update-instance",
			Fields: fields_update_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateInstance(ctx, input)
			},
		},
		"update-instance-access-control-attribute-configuration": {
			Name:   "update-instance-access-control-attribute-configuration",
			Fields: fields_update_instance_access_control_attribute_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateInstanceAccessControlAttributeConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_instance_access_control_attribute_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateInstanceAccessControlAttributeConfiguration(ctx, input)
			},
		},
		"update-permission-set": {
			Name:   "update-permission-set",
			Fields: fields_update_permission_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePermissionSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_permission_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePermissionSet(ctx, input)
			},
		},
		"update-trusted-token-issuer": {
			Name:   "update-trusted-token-issuer",
			Fields: fields_update_trusted_token_issuer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTrustedTokenIssuerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_trusted_token_issuer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTrustedTokenIssuer(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("ssoadmin", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
