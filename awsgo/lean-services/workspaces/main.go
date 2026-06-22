package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/workspaces"
)

var fields_accept_account_link_invitation = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "LinkId", Flag: "link-id", Type: "*string", Required: true},
}

var fields_associate_connection_alias = []leanruntime.Field{
	{Name: "AliasId", Flag: "alias-id", Type: "*string", Required: true},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
}

var fields_associate_ip_groups = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "GroupIds", Flag: "group-ids", Type: "[]string", Required: true},
}

var fields_associate_workspace_application = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_authorize_ip_rules = []leanruntime.Field{
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: true},
	{Name: "UserRules", Flag: "user-rules", Type: "[]types.IpRuleItem", Required: true},
}

var fields_copy_workspace_image = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "SourceImageId", Flag: "source-image-id", Type: "*string", Required: true},
	{Name: "SourceRegion", Flag: "source-region", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_account_link_invitation = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "TargetAccountId", Flag: "target-account-id", Type: "*string", Required: true},
}

var fields_create_connect_client_add_in = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "URL", Flag: "url", Type: "*string", Required: true},
}

var fields_create_connection_alias = []leanruntime.Field{
	{Name: "ConnectionString", Flag: "connection-string", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_ip_group = []leanruntime.Field{
	{Name: "GroupDesc", Flag: "group-desc", Type: "*string", Required: false},
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "UserRules", Flag: "user-rules", Type: "[]types.IpRuleItem", Required: false},
}

var fields_create_standby_workspaces = []leanruntime.Field{
	{Name: "PrimaryRegion", Flag: "primary-region", Type: "*string", Required: true},
	{Name: "StandbyWorkspaces", Flag: "standby-workspaces", Type: "[]types.StandbyWorkspace", Required: true},
}

var fields_create_tags = []leanruntime.Field{
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_create_updated_workspace_image = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "SourceImageId", Flag: "source-image-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_workspace_bundle = []leanruntime.Field{
	{Name: "BundleDescription", Flag: "bundle-description", Type: "*string", Required: true},
	{Name: "BundleName", Flag: "bundle-name", Type: "*string", Required: true},
	{Name: "ComputeType", Flag: "compute-type", Type: "*types.ComputeType", Required: true},
	{Name: "ImageId", Flag: "image-id", Type: "*string", Required: true},
	{Name: "RootStorage", Flag: "root-storage", Type: "*types.RootStorage", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "UserStorage", Flag: "user-storage", Type: "*types.UserStorage", Required: true},
}

var fields_create_workspace_image = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_create_workspaces = []leanruntime.Field{
	{Name: "Workspaces", Flag: "workspaces", Type: "[]types.WorkspaceRequest", Required: true},
}

var fields_create_workspaces_pool = []leanruntime.Field{
	{Name: "ApplicationSettings", Flag: "application-settings", Type: "*types.ApplicationSettingsRequest", Required: false},
	{Name: "BundleId", Flag: "bundle-id", Type: "*string", Required: true},
	{Name: "Capacity", Flag: "capacity", Type: "*types.Capacity", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "PoolName", Flag: "pool-name", Type: "*string", Required: true},
	{Name: "RunningMode", Flag: "running-mode", Type: "types.PoolsRunningMode", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TimeoutSettings", Flag: "timeout-settings", Type: "*types.TimeoutSettings", Required: false},
}

var fields_delete_account_link_invitation = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "LinkId", Flag: "link-id", Type: "*string", Required: true},
}

var fields_delete_client_branding = []leanruntime.Field{
	{Name: "Platforms", Flag: "platforms", Type: "[]types.ClientDeviceType", Required: true},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
}

var fields_delete_connect_client_add_in = []leanruntime.Field{
	{Name: "AddInId", Flag: "add-in-id", Type: "*string", Required: true},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
}

var fields_delete_connection_alias = []leanruntime.Field{
	{Name: "AliasId", Flag: "alias-id", Type: "*string", Required: true},
}

var fields_delete_ip_group = []leanruntime.Field{
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: true},
}

var fields_delete_tags = []leanruntime.Field{
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_delete_workspace_bundle = []leanruntime.Field{
	{Name: "BundleId", Flag: "bundle-id", Type: "*string", Required: false},
}

var fields_delete_workspace_image = []leanruntime.Field{
	{Name: "ImageId", Flag: "image-id", Type: "*string", Required: true},
}

var fields_deploy_workspace_applications = []leanruntime.Field{
	{Name: "Force", Flag: "force", Type: "*bool", Required: false},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_deregister_workspace_directory = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
}

var fields_describe_account = []leanruntime.Field{}

var fields_describe_account_modifications = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_application_associations = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "AssociatedResourceTypes", Flag: "associated-resource-types", Type: "[]types.ApplicationAssociatedResourceType", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_applications = []leanruntime.Field{
	{Name: "ApplicationIds", Flag: "application-ids", Type: "[]string", Required: false},
	{Name: "ComputeTypeNames", Flag: "compute-type-names", Type: "[]types.Compute", Required: false},
	{Name: "LicenseType", Flag: "license-type", Type: "types.WorkSpaceApplicationLicenseType", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OperatingSystemNames", Flag: "operating-system-names", Type: "[]types.OperatingSystemName", Required: false},
	{Name: "Owner", Flag: "owner", Type: "*string", Required: false},
}

var fields_describe_bundle_associations = []leanruntime.Field{
	{Name: "AssociatedResourceTypes", Flag: "associated-resource-types", Type: "[]types.BundleAssociatedResourceType", Required: true},
	{Name: "BundleId", Flag: "bundle-id", Type: "*string", Required: true},
}

var fields_describe_client_branding = []leanruntime.Field{
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
}

var fields_describe_client_properties = []leanruntime.Field{
	{Name: "ResourceIds", Flag: "resource-ids", Type: "[]string", Required: true},
}

var fields_describe_connect_client_add_ins = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
}

var fields_describe_connection_alias_permissions = []leanruntime.Field{
	{Name: "AliasId", Flag: "alias-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_connection_aliases = []leanruntime.Field{
	{Name: "AliasIds", Flag: "alias-ids", Type: "[]string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: false},
}

var fields_describe_custom_workspace_image_import = []leanruntime.Field{
	{Name: "ImageId", Flag: "image-id", Type: "*string", Required: true},
}

var fields_describe_image_associations = []leanruntime.Field{
	{Name: "AssociatedResourceTypes", Flag: "associated-resource-types", Type: "[]types.ImageAssociatedResourceType", Required: true},
	{Name: "ImageId", Flag: "image-id", Type: "*string", Required: true},
}

var fields_describe_ip_groups = []leanruntime.Field{
	{Name: "GroupIds", Flag: "group-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_tags = []leanruntime.Field{
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
}

var fields_describe_workspace_associations = []leanruntime.Field{
	{Name: "AssociatedResourceTypes", Flag: "associated-resource-types", Type: "[]types.WorkSpaceAssociatedResourceType", Required: true},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_describe_workspace_bundles = []leanruntime.Field{
	{Name: "BundleIds", Flag: "bundle-ids", Type: "[]string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Owner", Flag: "owner", Type: "*string", Required: false},
}

var fields_describe_workspace_directories = []leanruntime.Field{
	{Name: "DirectoryIds", Flag: "directory-ids", Type: "[]string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.DescribeWorkspaceDirectoriesFilter", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "WorkspaceDirectoryNames", Flag: "workspace-directory-names", Type: "[]string", Required: false},
}

var fields_describe_workspace_image_permissions = []leanruntime.Field{
	{Name: "ImageId", Flag: "image-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_workspace_images = []leanruntime.Field{
	{Name: "ImageIds", Flag: "image-ids", Type: "[]string", Required: false},
	{Name: "ImageType", Flag: "image-type", Type: "types.ImageType", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_workspace_snapshots = []leanruntime.Field{
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_describe_workspaces = []leanruntime.Field{
	{Name: "BundleId", Flag: "bundle-id", Type: "*string", Required: false},
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: false},
	{Name: "WorkspaceIds", Flag: "workspace-ids", Type: "[]string", Required: false},
	{Name: "WorkspaceName", Flag: "workspace-name", Type: "*string", Required: false},
}

var fields_describe_workspaces_connection_status = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "WorkspaceIds", Flag: "workspace-ids", Type: "[]string", Required: false},
}

var fields_describe_workspaces_pool_sessions = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PoolId", Flag: "pool-id", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: false},
}

var fields_describe_workspaces_pools = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.DescribeWorkspacesPoolsFilter", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PoolIds", Flag: "pool-ids", Type: "[]string", Required: false},
}

var fields_disassociate_connection_alias = []leanruntime.Field{
	{Name: "AliasId", Flag: "alias-id", Type: "*string", Required: true},
}

var fields_disassociate_ip_groups = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "GroupIds", Flag: "group-ids", Type: "[]string", Required: true},
}

var fields_disassociate_workspace_application = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_get_account_link = []leanruntime.Field{
	{Name: "LinkId", Flag: "link-id", Type: "*string", Required: false},
	{Name: "LinkedAccountId", Flag: "linked-account-id", Type: "*string", Required: false},
}

var fields_import_client_branding = []leanruntime.Field{
	{Name: "DeviceTypeAndroid", Flag: "device-type-android", Type: "*types.DefaultImportClientBrandingAttributes", Required: false},
	{Name: "DeviceTypeIos", Flag: "device-type-ios", Type: "*types.IosImportClientBrandingAttributes", Required: false},
	{Name: "DeviceTypeLinux", Flag: "device-type-linux", Type: "*types.DefaultImportClientBrandingAttributes", Required: false},
	{Name: "DeviceTypeOsx", Flag: "device-type-osx", Type: "*types.DefaultImportClientBrandingAttributes", Required: false},
	{Name: "DeviceTypeWeb", Flag: "device-type-web", Type: "*types.DefaultImportClientBrandingAttributes", Required: false},
	{Name: "DeviceTypeWindows", Flag: "device-type-windows", Type: "*types.DefaultImportClientBrandingAttributes", Required: false},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
}

var fields_import_custom_workspace_image = []leanruntime.Field{
	{Name: "ComputeType", Flag: "compute-type", Type: "types.ImageComputeType", Required: true},
	{Name: "ImageDescription", Flag: "image-description", Type: "*string", Required: true},
	{Name: "ImageName", Flag: "image-name", Type: "*string", Required: true},
	{Name: "ImageSource", Flag: "image-source", Type: "types.ImageSourceIdentifier", Required: true},
	{Name: "InfrastructureConfigurationArn", Flag: "infrastructure-configuration-arn", Type: "*string", Required: true},
	{Name: "OsVersion", Flag: "os-version", Type: "types.OSVersion", Required: true},
	{Name: "Platform", Flag: "platform", Type: "types.Platform", Required: true},
	{Name: "Protocol", Flag: "protocol", Type: "types.CustomImageProtocol", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_import_workspace_image = []leanruntime.Field{
	{Name: "Applications", Flag: "applications", Type: "[]types.Application", Required: false},
	{Name: "Ec2ImageId", Flag: "ec2-image-id", Type: "*string", Required: true},
	{Name: "ImageDescription", Flag: "image-description", Type: "*string", Required: true},
	{Name: "ImageName", Flag: "image-name", Type: "*string", Required: true},
	{Name: "IngestionProcess", Flag: "ingestion-process", Type: "types.WorkspaceImageIngestionProcess", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_list_account_links = []leanruntime.Field{
	{Name: "LinkStatusFilter", Flag: "link-status-filter", Type: "[]types.AccountLinkStatusEnum", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_available_management_cidr_ranges = []leanruntime.Field{
	{Name: "ManagementCidrRangeConstraint", Flag: "management-cidr-range-constraint", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_migrate_workspace = []leanruntime.Field{
	{Name: "BundleId", Flag: "bundle-id", Type: "*string", Required: true},
	{Name: "SourceWorkspaceId", Flag: "source-workspace-id", Type: "*string", Required: true},
}

var fields_modify_account = []leanruntime.Field{
	{Name: "DedicatedTenancyManagementCidrRange", Flag: "dedicated-tenancy-management-cidr-range", Type: "*string", Required: false},
	{Name: "DedicatedTenancySupport", Flag: "dedicated-tenancy-support", Type: "types.DedicatedTenancySupportEnum", Required: false},
}

var fields_modify_certificate_based_auth_properties = []leanruntime.Field{
	{Name: "CertificateBasedAuthProperties", Flag: "certificate-based-auth-properties", Type: "*types.CertificateBasedAuthProperties", Required: false},
	{Name: "PropertiesToDelete", Flag: "properties-to-delete", Type: "[]types.DeletableCertificateBasedAuthProperty", Required: false},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
}

var fields_modify_client_properties = []leanruntime.Field{
	{Name: "ClientProperties", Flag: "client-properties", Type: "*types.ClientProperties", Required: true},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
}

var fields_modify_endpoint_encryption_mode = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "EndpointEncryptionMode", Flag: "endpoint-encryption-mode", Type: "types.EndpointEncryptionMode", Required: true},
}

var fields_modify_saml_properties = []leanruntime.Field{
	{Name: "PropertiesToDelete", Flag: "properties-to-delete", Type: "[]types.DeletableSamlProperty", Required: false},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "SamlProperties", Flag: "saml-properties", Type: "*types.SamlProperties", Required: false},
}

var fields_modify_selfservice_permissions = []leanruntime.Field{
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "SelfservicePermissions", Flag: "selfservice-permissions", Type: "*types.SelfservicePermissions", Required: true},
}

var fields_modify_streaming_properties = []leanruntime.Field{
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "StreamingProperties", Flag: "streaming-properties", Type: "*types.StreamingProperties", Required: false},
}

var fields_modify_workspace_access_properties = []leanruntime.Field{
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "WorkspaceAccessProperties", Flag: "workspace-access-properties", Type: "*types.WorkspaceAccessProperties", Required: true},
}

var fields_modify_workspace_creation_properties = []leanruntime.Field{
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "WorkspaceCreationProperties", Flag: "workspace-creation-properties", Type: "*types.WorkspaceCreationProperties", Required: true},
}

var fields_modify_workspace_properties = []leanruntime.Field{
	{Name: "DataReplication", Flag: "data-replication", Type: "types.DataReplication", Required: false},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
	{Name: "WorkspaceProperties", Flag: "workspace-properties", Type: "*types.WorkspaceProperties", Required: false},
}

var fields_modify_workspace_state = []leanruntime.Field{
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
	{Name: "WorkspaceState", Flag: "workspace-state", Type: "types.TargetWorkspaceState", Required: true},
}

var fields_reboot_workspaces = []leanruntime.Field{
	{Name: "RebootWorkspaceRequests", Flag: "reboot-workspace-requests", Type: "[]types.RebootRequest", Required: true},
}

var fields_rebuild_workspaces = []leanruntime.Field{
	{Name: "RebuildWorkspaceRequests", Flag: "rebuild-workspace-requests", Type: "[]types.RebuildRequest", Required: true},
}

var fields_register_workspace_directory = []leanruntime.Field{
	{Name: "ActiveDirectoryConfig", Flag: "active-directory-config", Type: "*types.ActiveDirectoryConfig", Required: false},
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: false},
	{Name: "EnableSelfService", Flag: "enable-self-service", Type: "*bool", Required: false},
	{Name: "IdcInstanceArn", Flag: "idc-instance-arn", Type: "*string", Required: false},
	{Name: "MicrosoftEntraConfig", Flag: "microsoft-entra-config", Type: "*types.MicrosoftEntraConfig", Required: false},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Tenancy", Flag: "tenancy", Type: "types.Tenancy", Required: false},
	{Name: "UserIdentityType", Flag: "user-identity-type", Type: "types.UserIdentityType", Required: false},
	{Name: "WorkspaceDirectoryDescription", Flag: "workspace-directory-description", Type: "*string", Required: false},
	{Name: "WorkspaceDirectoryName", Flag: "workspace-directory-name", Type: "*string", Required: false},
	{Name: "WorkspaceType", Flag: "workspace-type", Type: "types.WorkspaceType", Required: false},
}

var fields_reject_account_link_invitation = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "LinkId", Flag: "link-id", Type: "*string", Required: true},
}

var fields_restore_workspace = []leanruntime.Field{
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_revoke_ip_rules = []leanruntime.Field{
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: true},
	{Name: "UserRules", Flag: "user-rules", Type: "[]string", Required: true},
}

var fields_start_workspaces = []leanruntime.Field{
	{Name: "StartWorkspaceRequests", Flag: "start-workspace-requests", Type: "[]types.StartRequest", Required: true},
}

var fields_start_workspaces_pool = []leanruntime.Field{
	{Name: "PoolId", Flag: "pool-id", Type: "*string", Required: true},
}

var fields_stop_workspaces = []leanruntime.Field{
	{Name: "StopWorkspaceRequests", Flag: "stop-workspace-requests", Type: "[]types.StopRequest", Required: true},
}

var fields_stop_workspaces_pool = []leanruntime.Field{
	{Name: "PoolId", Flag: "pool-id", Type: "*string", Required: true},
}

var fields_terminate_workspaces = []leanruntime.Field{
	{Name: "TerminateWorkspaceRequests", Flag: "terminate-workspace-requests", Type: "[]types.TerminateRequest", Required: true},
}

var fields_terminate_workspaces_pool = []leanruntime.Field{
	{Name: "PoolId", Flag: "pool-id", Type: "*string", Required: true},
}

var fields_terminate_workspaces_pool_session = []leanruntime.Field{
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
}

var fields_update_connect_client_add_in = []leanruntime.Field{
	{Name: "AddInId", Flag: "add-in-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "URL", Flag: "url", Type: "*string", Required: false},
}

var fields_update_connection_alias_permission = []leanruntime.Field{
	{Name: "AliasId", Flag: "alias-id", Type: "*string", Required: true},
	{Name: "ConnectionAliasPermission", Flag: "connection-alias-permission", Type: "*types.ConnectionAliasPermission", Required: true},
}

var fields_update_rules_of_ip_group = []leanruntime.Field{
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: true},
	{Name: "UserRules", Flag: "user-rules", Type: "[]types.IpRuleItem", Required: true},
}

var fields_update_workspace_bundle = []leanruntime.Field{
	{Name: "BundleId", Flag: "bundle-id", Type: "*string", Required: false},
	{Name: "ImageId", Flag: "image-id", Type: "*string", Required: false},
}

var fields_update_workspace_image_permission = []leanruntime.Field{
	{Name: "AllowCopyImage", Flag: "allow-copy-image", Type: "*bool", Required: true},
	{Name: "ImageId", Flag: "image-id", Type: "*string", Required: true},
	{Name: "SharedAccountId", Flag: "shared-account-id", Type: "*string", Required: true},
}

var fields_update_workspaces_pool = []leanruntime.Field{
	{Name: "ApplicationSettings", Flag: "application-settings", Type: "*types.ApplicationSettingsRequest", Required: false},
	{Name: "BundleId", Flag: "bundle-id", Type: "*string", Required: false},
	{Name: "Capacity", Flag: "capacity", Type: "*types.Capacity", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: false},
	{Name: "PoolId", Flag: "pool-id", Type: "*string", Required: true},
	{Name: "RunningMode", Flag: "running-mode", Type: "types.PoolsRunningMode", Required: false},
	{Name: "TimeoutSettings", Flag: "timeout-settings", Type: "*types.TimeoutSettings", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"accept-account-link-invitation": {
			Name:   "accept-account-link-invitation",
			Fields: fields_accept_account_link_invitation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AcceptAccountLinkInvitationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_accept_account_link_invitation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AcceptAccountLinkInvitation(ctx, input)
			},
		},
		"associate-connection-alias": {
			Name:   "associate-connection-alias",
			Fields: fields_associate_connection_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateConnectionAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_connection_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateConnectionAlias(ctx, input)
			},
		},
		"associate-ip-groups": {
			Name:   "associate-ip-groups",
			Fields: fields_associate_ip_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateIpGroupsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_ip_groups, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateIpGroups(ctx, input)
			},
		},
		"associate-workspace-application": {
			Name:   "associate-workspace-application",
			Fields: fields_associate_workspace_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateWorkspaceApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_workspace_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateWorkspaceApplication(ctx, input)
			},
		},
		"authorize-ip-rules": {
			Name:   "authorize-ip-rules",
			Fields: fields_authorize_ip_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AuthorizeIpRulesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_authorize_ip_rules, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AuthorizeIpRules(ctx, input)
			},
		},
		"copy-workspace-image": {
			Name:   "copy-workspace-image",
			Fields: fields_copy_workspace_image,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CopyWorkspaceImageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_copy_workspace_image, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CopyWorkspaceImage(ctx, input)
			},
		},
		"create-account-link-invitation": {
			Name:   "create-account-link-invitation",
			Fields: fields_create_account_link_invitation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAccountLinkInvitationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_account_link_invitation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAccountLinkInvitation(ctx, input)
			},
		},
		"create-connect-client-add-in": {
			Name:   "create-connect-client-add-in",
			Fields: fields_create_connect_client_add_in,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateConnectClientAddInInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_connect_client_add_in, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateConnectClientAddIn(ctx, input)
			},
		},
		"create-connection-alias": {
			Name:   "create-connection-alias",
			Fields: fields_create_connection_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateConnectionAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_connection_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateConnectionAlias(ctx, input)
			},
		},
		"create-ip-group": {
			Name:   "create-ip-group",
			Fields: fields_create_ip_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateIpGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_ip_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateIpGroup(ctx, input)
			},
		},
		"create-standby-workspaces": {
			Name:   "create-standby-workspaces",
			Fields: fields_create_standby_workspaces,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateStandbyWorkspacesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_standby_workspaces, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateStandbyWorkspaces(ctx, input)
			},
		},
		"create-tags": {
			Name:   "create-tags",
			Fields: fields_create_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTagsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_tags, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTags(ctx, input)
			},
		},
		"create-updated-workspace-image": {
			Name:   "create-updated-workspace-image",
			Fields: fields_create_updated_workspace_image,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateUpdatedWorkspaceImageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_updated_workspace_image, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateUpdatedWorkspaceImage(ctx, input)
			},
		},
		"create-workspace-bundle": {
			Name:   "create-workspace-bundle",
			Fields: fields_create_workspace_bundle,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWorkspaceBundleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_workspace_bundle, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWorkspaceBundle(ctx, input)
			},
		},
		"create-workspace-image": {
			Name:   "create-workspace-image",
			Fields: fields_create_workspace_image,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWorkspaceImageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_workspace_image, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWorkspaceImage(ctx, input)
			},
		},
		"create-workspaces": {
			Name:   "create-workspaces",
			Fields: fields_create_workspaces,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWorkspacesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_workspaces, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWorkspaces(ctx, input)
			},
		},
		"create-workspaces-pool": {
			Name:   "create-workspaces-pool",
			Fields: fields_create_workspaces_pool,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWorkspacesPoolInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_workspaces_pool, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWorkspacesPool(ctx, input)
			},
		},
		"delete-account-link-invitation": {
			Name:   "delete-account-link-invitation",
			Fields: fields_delete_account_link_invitation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAccountLinkInvitationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_account_link_invitation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAccountLinkInvitation(ctx, input)
			},
		},
		"delete-client-branding": {
			Name:   "delete-client-branding",
			Fields: fields_delete_client_branding,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteClientBrandingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_client_branding, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteClientBranding(ctx, input)
			},
		},
		"delete-connect-client-add-in": {
			Name:   "delete-connect-client-add-in",
			Fields: fields_delete_connect_client_add_in,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConnectClientAddInInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_connect_client_add_in, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConnectClientAddIn(ctx, input)
			},
		},
		"delete-connection-alias": {
			Name:   "delete-connection-alias",
			Fields: fields_delete_connection_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConnectionAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_connection_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConnectionAlias(ctx, input)
			},
		},
		"delete-ip-group": {
			Name:   "delete-ip-group",
			Fields: fields_delete_ip_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIpGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_ip_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIpGroup(ctx, input)
			},
		},
		"delete-tags": {
			Name:   "delete-tags",
			Fields: fields_delete_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTagsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_tags, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTags(ctx, input)
			},
		},
		"delete-workspace-bundle": {
			Name:   "delete-workspace-bundle",
			Fields: fields_delete_workspace_bundle,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWorkspaceBundleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_workspace_bundle, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWorkspaceBundle(ctx, input)
			},
		},
		"delete-workspace-image": {
			Name:   "delete-workspace-image",
			Fields: fields_delete_workspace_image,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWorkspaceImageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_workspace_image, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWorkspaceImage(ctx, input)
			},
		},
		"deploy-workspace-applications": {
			Name:   "deploy-workspace-applications",
			Fields: fields_deploy_workspace_applications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeployWorkspaceApplicationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deploy_workspace_applications, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeployWorkspaceApplications(ctx, input)
			},
		},
		"deregister-workspace-directory": {
			Name:   "deregister-workspace-directory",
			Fields: fields_deregister_workspace_directory,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterWorkspaceDirectoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_workspace_directory, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterWorkspaceDirectory(ctx, input)
			},
		},
		"describe-account": {
			Name:   "describe-account",
			Fields: fields_describe_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAccount(ctx, input)
			},
		},
		"describe-account-modifications": {
			Name:   "describe-account-modifications",
			Fields: fields_describe_account_modifications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAccountModificationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_account_modifications, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAccountModifications(ctx, input)
			},
		},
		"describe-application-associations": {
			Name:   "describe-application-associations",
			Fields: fields_describe_application_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeApplicationAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_application_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeApplicationAssociations(ctx, input)
				}
				var results []*svc.DescribeApplicationAssociationsOutput
				p := svc.NewDescribeApplicationAssociationsPaginator(client, input)
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
		"describe-applications": {
			Name:   "describe-applications",
			Fields: fields_describe_applications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeApplicationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_applications, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeApplications(ctx, input)
				}
				var results []*svc.DescribeApplicationsOutput
				p := svc.NewDescribeApplicationsPaginator(client, input)
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
		"describe-bundle-associations": {
			Name:   "describe-bundle-associations",
			Fields: fields_describe_bundle_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeBundleAssociationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_bundle_associations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeBundleAssociations(ctx, input)
			},
		},
		"describe-client-branding": {
			Name:   "describe-client-branding",
			Fields: fields_describe_client_branding,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeClientBrandingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_client_branding, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeClientBranding(ctx, input)
			},
		},
		"describe-client-properties": {
			Name:   "describe-client-properties",
			Fields: fields_describe_client_properties,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeClientPropertiesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_client_properties, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeClientProperties(ctx, input)
			},
		},
		"describe-connect-client-add-ins": {
			Name:   "describe-connect-client-add-ins",
			Fields: fields_describe_connect_client_add_ins,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeConnectClientAddInsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_connect_client_add_ins, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeConnectClientAddIns(ctx, input)
			},
		},
		"describe-connection-alias-permissions": {
			Name:   "describe-connection-alias-permissions",
			Fields: fields_describe_connection_alias_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeConnectionAliasPermissionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_connection_alias_permissions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeConnectionAliasPermissions(ctx, input)
			},
		},
		"describe-connection-aliases": {
			Name:   "describe-connection-aliases",
			Fields: fields_describe_connection_aliases,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeConnectionAliasesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_connection_aliases, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeConnectionAliases(ctx, input)
			},
		},
		"describe-custom-workspace-image-import": {
			Name:   "describe-custom-workspace-image-import",
			Fields: fields_describe_custom_workspace_image_import,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCustomWorkspaceImageImportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_custom_workspace_image_import, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeCustomWorkspaceImageImport(ctx, input)
			},
		},
		"describe-image-associations": {
			Name:   "describe-image-associations",
			Fields: fields_describe_image_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeImageAssociationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_image_associations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeImageAssociations(ctx, input)
			},
		},
		"describe-ip-groups": {
			Name:   "describe-ip-groups",
			Fields: fields_describe_ip_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeIpGroupsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_ip_groups, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeIpGroups(ctx, input)
			},
		},
		"describe-tags": {
			Name:   "describe-tags",
			Fields: fields_describe_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTagsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_tags, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTags(ctx, input)
			},
		},
		"describe-workspace-associations": {
			Name:   "describe-workspace-associations",
			Fields: fields_describe_workspace_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeWorkspaceAssociationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_workspace_associations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeWorkspaceAssociations(ctx, input)
			},
		},
		"describe-workspace-bundles": {
			Name:   "describe-workspace-bundles",
			Fields: fields_describe_workspace_bundles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeWorkspaceBundlesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_workspace_bundles, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeWorkspaceBundles(ctx, input)
				}
				var results []*svc.DescribeWorkspaceBundlesOutput
				p := svc.NewDescribeWorkspaceBundlesPaginator(client, input)
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
		"describe-workspace-directories": {
			Name:   "describe-workspace-directories",
			Fields: fields_describe_workspace_directories,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeWorkspaceDirectoriesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_workspace_directories, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeWorkspaceDirectories(ctx, input)
				}
				var results []*svc.DescribeWorkspaceDirectoriesOutput
				p := svc.NewDescribeWorkspaceDirectoriesPaginator(client, input)
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
		"describe-workspace-image-permissions": {
			Name:   "describe-workspace-image-permissions",
			Fields: fields_describe_workspace_image_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeWorkspaceImagePermissionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_workspace_image_permissions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeWorkspaceImagePermissions(ctx, input)
			},
		},
		"describe-workspace-images": {
			Name:   "describe-workspace-images",
			Fields: fields_describe_workspace_images,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeWorkspaceImagesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_workspace_images, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeWorkspaceImages(ctx, input)
			},
		},
		"describe-workspace-snapshots": {
			Name:   "describe-workspace-snapshots",
			Fields: fields_describe_workspace_snapshots,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeWorkspaceSnapshotsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_workspace_snapshots, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeWorkspaceSnapshots(ctx, input)
			},
		},
		"describe-workspaces": {
			Name:   "describe-workspaces",
			Fields: fields_describe_workspaces,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeWorkspacesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_workspaces, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeWorkspaces(ctx, input)
				}
				var results []*svc.DescribeWorkspacesOutput
				p := svc.NewDescribeWorkspacesPaginator(client, input)
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
		"describe-workspaces-connection-status": {
			Name:   "describe-workspaces-connection-status",
			Fields: fields_describe_workspaces_connection_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeWorkspacesConnectionStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_workspaces_connection_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeWorkspacesConnectionStatus(ctx, input)
			},
		},
		"describe-workspaces-pool-sessions": {
			Name:   "describe-workspaces-pool-sessions",
			Fields: fields_describe_workspaces_pool_sessions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeWorkspacesPoolSessionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_workspaces_pool_sessions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeWorkspacesPoolSessions(ctx, input)
			},
		},
		"describe-workspaces-pools": {
			Name:   "describe-workspaces-pools",
			Fields: fields_describe_workspaces_pools,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeWorkspacesPoolsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_workspaces_pools, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeWorkspacesPools(ctx, input)
			},
		},
		"disassociate-connection-alias": {
			Name:   "disassociate-connection-alias",
			Fields: fields_disassociate_connection_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateConnectionAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_connection_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateConnectionAlias(ctx, input)
			},
		},
		"disassociate-ip-groups": {
			Name:   "disassociate-ip-groups",
			Fields: fields_disassociate_ip_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateIpGroupsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_ip_groups, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateIpGroups(ctx, input)
			},
		},
		"disassociate-workspace-application": {
			Name:   "disassociate-workspace-application",
			Fields: fields_disassociate_workspace_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateWorkspaceApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_workspace_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateWorkspaceApplication(ctx, input)
			},
		},
		"get-account-link": {
			Name:   "get-account-link",
			Fields: fields_get_account_link,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccountLinkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_account_link, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccountLink(ctx, input)
			},
		},
		"import-client-branding": {
			Name:   "import-client-branding",
			Fields: fields_import_client_branding,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportClientBrandingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_client_branding, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportClientBranding(ctx, input)
			},
		},
		"import-custom-workspace-image": {
			Name:   "import-custom-workspace-image",
			Fields: fields_import_custom_workspace_image,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportCustomWorkspaceImageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_custom_workspace_image, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportCustomWorkspaceImage(ctx, input)
			},
		},
		"import-workspace-image": {
			Name:   "import-workspace-image",
			Fields: fields_import_workspace_image,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportWorkspaceImageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_workspace_image, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportWorkspaceImage(ctx, input)
			},
		},
		"list-account-links": {
			Name:   "list-account-links",
			Fields: fields_list_account_links,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAccountLinksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_account_links, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAccountLinks(ctx, input)
				}
				var results []*svc.ListAccountLinksOutput
				p := svc.NewListAccountLinksPaginator(client, input)
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
		"list-available-management-cidr-ranges": {
			Name:   "list-available-management-cidr-ranges",
			Fields: fields_list_available_management_cidr_ranges,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAvailableManagementCidrRangesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_available_management_cidr_ranges, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListAvailableManagementCidrRanges(ctx, input)
			},
		},
		"migrate-workspace": {
			Name:   "migrate-workspace",
			Fields: fields_migrate_workspace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.MigrateWorkspaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_migrate_workspace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.MigrateWorkspace(ctx, input)
			},
		},
		"modify-account": {
			Name:   "modify-account",
			Fields: fields_modify_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyAccount(ctx, input)
			},
		},
		"modify-certificate-based-auth-properties": {
			Name:   "modify-certificate-based-auth-properties",
			Fields: fields_modify_certificate_based_auth_properties,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyCertificateBasedAuthPropertiesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_certificate_based_auth_properties, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyCertificateBasedAuthProperties(ctx, input)
			},
		},
		"modify-client-properties": {
			Name:   "modify-client-properties",
			Fields: fields_modify_client_properties,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyClientPropertiesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_client_properties, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyClientProperties(ctx, input)
			},
		},
		"modify-endpoint-encryption-mode": {
			Name:   "modify-endpoint-encryption-mode",
			Fields: fields_modify_endpoint_encryption_mode,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyEndpointEncryptionModeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_endpoint_encryption_mode, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyEndpointEncryptionMode(ctx, input)
			},
		},
		"modify-saml-properties": {
			Name:   "modify-saml-properties",
			Fields: fields_modify_saml_properties,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifySamlPropertiesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_saml_properties, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifySamlProperties(ctx, input)
			},
		},
		"modify-selfservice-permissions": {
			Name:   "modify-selfservice-permissions",
			Fields: fields_modify_selfservice_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifySelfservicePermissionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_selfservice_permissions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifySelfservicePermissions(ctx, input)
			},
		},
		"modify-streaming-properties": {
			Name:   "modify-streaming-properties",
			Fields: fields_modify_streaming_properties,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyStreamingPropertiesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_streaming_properties, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyStreamingProperties(ctx, input)
			},
		},
		"modify-workspace-access-properties": {
			Name:   "modify-workspace-access-properties",
			Fields: fields_modify_workspace_access_properties,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyWorkspaceAccessPropertiesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_workspace_access_properties, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyWorkspaceAccessProperties(ctx, input)
			},
		},
		"modify-workspace-creation-properties": {
			Name:   "modify-workspace-creation-properties",
			Fields: fields_modify_workspace_creation_properties,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyWorkspaceCreationPropertiesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_workspace_creation_properties, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyWorkspaceCreationProperties(ctx, input)
			},
		},
		"modify-workspace-properties": {
			Name:   "modify-workspace-properties",
			Fields: fields_modify_workspace_properties,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyWorkspacePropertiesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_workspace_properties, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyWorkspaceProperties(ctx, input)
			},
		},
		"modify-workspace-state": {
			Name:   "modify-workspace-state",
			Fields: fields_modify_workspace_state,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyWorkspaceStateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_workspace_state, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyWorkspaceState(ctx, input)
			},
		},
		"reboot-workspaces": {
			Name:   "reboot-workspaces",
			Fields: fields_reboot_workspaces,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RebootWorkspacesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reboot_workspaces, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RebootWorkspaces(ctx, input)
			},
		},
		"rebuild-workspaces": {
			Name:   "rebuild-workspaces",
			Fields: fields_rebuild_workspaces,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RebuildWorkspacesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_rebuild_workspaces, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RebuildWorkspaces(ctx, input)
			},
		},
		"register-workspace-directory": {
			Name:   "register-workspace-directory",
			Fields: fields_register_workspace_directory,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterWorkspaceDirectoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_workspace_directory, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterWorkspaceDirectory(ctx, input)
			},
		},
		"reject-account-link-invitation": {
			Name:   "reject-account-link-invitation",
			Fields: fields_reject_account_link_invitation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RejectAccountLinkInvitationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reject_account_link_invitation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RejectAccountLinkInvitation(ctx, input)
			},
		},
		"restore-workspace": {
			Name:   "restore-workspace",
			Fields: fields_restore_workspace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RestoreWorkspaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_restore_workspace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RestoreWorkspace(ctx, input)
			},
		},
		"revoke-ip-rules": {
			Name:   "revoke-ip-rules",
			Fields: fields_revoke_ip_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RevokeIpRulesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_revoke_ip_rules, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RevokeIpRules(ctx, input)
			},
		},
		"start-workspaces": {
			Name:   "start-workspaces",
			Fields: fields_start_workspaces,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartWorkspacesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_workspaces, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartWorkspaces(ctx, input)
			},
		},
		"start-workspaces-pool": {
			Name:   "start-workspaces-pool",
			Fields: fields_start_workspaces_pool,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartWorkspacesPoolInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_workspaces_pool, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartWorkspacesPool(ctx, input)
			},
		},
		"stop-workspaces": {
			Name:   "stop-workspaces",
			Fields: fields_stop_workspaces,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopWorkspacesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_workspaces, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopWorkspaces(ctx, input)
			},
		},
		"stop-workspaces-pool": {
			Name:   "stop-workspaces-pool",
			Fields: fields_stop_workspaces_pool,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopWorkspacesPoolInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_workspaces_pool, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopWorkspacesPool(ctx, input)
			},
		},
		"terminate-workspaces": {
			Name:   "terminate-workspaces",
			Fields: fields_terminate_workspaces,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TerminateWorkspacesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_terminate_workspaces, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TerminateWorkspaces(ctx, input)
			},
		},
		"terminate-workspaces-pool": {
			Name:   "terminate-workspaces-pool",
			Fields: fields_terminate_workspaces_pool,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TerminateWorkspacesPoolInput{}
				if _, err := leanruntime.ApplyInput(input, fields_terminate_workspaces_pool, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TerminateWorkspacesPool(ctx, input)
			},
		},
		"terminate-workspaces-pool-session": {
			Name:   "terminate-workspaces-pool-session",
			Fields: fields_terminate_workspaces_pool_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TerminateWorkspacesPoolSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_terminate_workspaces_pool_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TerminateWorkspacesPoolSession(ctx, input)
			},
		},
		"update-connect-client-add-in": {
			Name:   "update-connect-client-add-in",
			Fields: fields_update_connect_client_add_in,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateConnectClientAddInInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_connect_client_add_in, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateConnectClientAddIn(ctx, input)
			},
		},
		"update-connection-alias-permission": {
			Name:   "update-connection-alias-permission",
			Fields: fields_update_connection_alias_permission,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateConnectionAliasPermissionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_connection_alias_permission, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateConnectionAliasPermission(ctx, input)
			},
		},
		"update-rules-of-ip-group": {
			Name:   "update-rules-of-ip-group",
			Fields: fields_update_rules_of_ip_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRulesOfIpGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_rules_of_ip_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRulesOfIpGroup(ctx, input)
			},
		},
		"update-workspace-bundle": {
			Name:   "update-workspace-bundle",
			Fields: fields_update_workspace_bundle,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateWorkspaceBundleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_workspace_bundle, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateWorkspaceBundle(ctx, input)
			},
		},
		"update-workspace-image-permission": {
			Name:   "update-workspace-image-permission",
			Fields: fields_update_workspace_image_permission,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateWorkspaceImagePermissionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_workspace_image_permission, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateWorkspaceImagePermission(ctx, input)
			},
		},
		"update-workspaces-pool": {
			Name:   "update-workspaces-pool",
			Fields: fields_update_workspaces_pool,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateWorkspacesPoolInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_workspaces_pool, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateWorkspacesPool(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("workspaces", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
