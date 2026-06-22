package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/appstream"
)

var fields_associate_app_block_builder_app_block = []leanruntime.Field{
	{Name: "AppBlockArn", Flag: "app-block-arn", Type: "*string", Required: true},
	{Name: "AppBlockBuilderName", Flag: "app-block-builder-name", Type: "*string", Required: true},
}

var fields_associate_application_fleet = []leanruntime.Field{
	{Name: "ApplicationArn", Flag: "application-arn", Type: "*string", Required: true},
	{Name: "FleetName", Flag: "fleet-name", Type: "*string", Required: true},
}

var fields_associate_application_to_entitlement = []leanruntime.Field{
	{Name: "ApplicationIdentifier", Flag: "application-identifier", Type: "*string", Required: true},
	{Name: "EntitlementName", Flag: "entitlement-name", Type: "*string", Required: true},
	{Name: "StackName", Flag: "stack-name", Type: "*string", Required: true},
}

var fields_associate_fleet = []leanruntime.Field{
	{Name: "FleetName", Flag: "fleet-name", Type: "*string", Required: true},
	{Name: "StackName", Flag: "stack-name", Type: "*string", Required: true},
}

var fields_associate_software_to_image_builder = []leanruntime.Field{
	{Name: "ImageBuilderName", Flag: "image-builder-name", Type: "*string", Required: true},
	{Name: "SoftwareNames", Flag: "software-names", Type: "[]string", Required: true},
}

var fields_batch_associate_user_stack = []leanruntime.Field{
	{Name: "UserStackAssociations", Flag: "user-stack-associations", Type: "[]types.UserStackAssociation", Required: true},
}

var fields_batch_disassociate_user_stack = []leanruntime.Field{
	{Name: "UserStackAssociations", Flag: "user-stack-associations", Type: "[]types.UserStackAssociation", Required: true},
}

var fields_copy_image = []leanruntime.Field{
	{Name: "DestinationImageDescription", Flag: "destination-image-description", Type: "*string", Required: false},
	{Name: "DestinationImageName", Flag: "destination-image-name", Type: "*string", Required: true},
	{Name: "DestinationRegion", Flag: "destination-region", Type: "*string", Required: true},
	{Name: "SourceImageName", Flag: "source-image-name", Type: "*string", Required: true},
}

var fields_create_app_block = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "PackagingType", Flag: "packaging-type", Type: "types.PackagingType", Required: false},
	{Name: "PostSetupScriptDetails", Flag: "post-setup-script-details", Type: "*types.ScriptDetails", Required: false},
	{Name: "SetupScriptDetails", Flag: "setup-script-details", Type: "*types.ScriptDetails", Required: false},
	{Name: "SourceS3Location", Flag: "source-s3-location", Type: "*types.S3Location", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_app_block_builder = []leanruntime.Field{
	{Name: "AccessEndpoints", Flag: "access-endpoints", Type: "[]types.AccessEndpoint", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisableIMDSV1", Flag: "disable-imdsv1", Type: "*bool", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "EnableDefaultInternetAccess", Flag: "enable-default-internet-access", Type: "*bool", Required: false},
	{Name: "IamRoleArn", Flag: "iam-role-arn", Type: "*string", Required: false},
	{Name: "InstanceType", Flag: "instance-type", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Platform", Flag: "platform", Type: "types.AppBlockBuilderPlatformType", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "VpcConfig", Flag: "vpc-config", Type: "*types.VpcConfig", Required: true},
}

var fields_create_app_block_builder_streaming_url = []leanruntime.Field{
	{Name: "AppBlockBuilderName", Flag: "app-block-builder-name", Type: "*string", Required: true},
	{Name: "Validity", Flag: "validity", Type: "*int64", Required: false},
}

var fields_create_application = []leanruntime.Field{
	{Name: "AppBlockArn", Flag: "app-block-arn", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "IconS3Location", Flag: "icon-s3-location", Type: "*types.S3Location", Required: true},
	{Name: "InstanceFamilies", Flag: "instance-families", Type: "[]string", Required: true},
	{Name: "LaunchParameters", Flag: "launch-parameters", Type: "*string", Required: false},
	{Name: "LaunchPath", Flag: "launch-path", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Platforms", Flag: "platforms", Type: "[]types.PlatformType", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "WorkingDirectory", Flag: "working-directory", Type: "*string", Required: false},
}

var fields_create_directory_config = []leanruntime.Field{
	{Name: "CertificateBasedAuthProperties", Flag: "certificate-based-auth-properties", Type: "*types.CertificateBasedAuthProperties", Required: false},
	{Name: "DirectoryName", Flag: "directory-name", Type: "*string", Required: true},
	{Name: "OrganizationalUnitDistinguishedNames", Flag: "organizational-unit-distinguished-names", Type: "[]string", Required: true},
	{Name: "ServiceAccountCredentials", Flag: "service-account-credentials", Type: "*types.ServiceAccountCredentials", Required: false},
}

var fields_create_entitlement = []leanruntime.Field{
	{Name: "AppVisibility", Flag: "app-visibility", Type: "types.AppVisibility", Required: true},
	{Name: "Attributes", Flag: "attributes", Type: "[]types.EntitlementAttribute", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "StackName", Flag: "stack-name", Type: "*string", Required: true},
}

var fields_create_export_image_task = []leanruntime.Field{
	{Name: "AmiDescription", Flag: "ami-description", Type: "*string", Required: false},
	{Name: "AmiName", Flag: "ami-name", Type: "*string", Required: true},
	{Name: "IamRoleArn", Flag: "iam-role-arn", Type: "*string", Required: true},
	{Name: "ImageName", Flag: "image-name", Type: "*string", Required: true},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "map[string]string", Required: false},
}

var fields_create_fleet = []leanruntime.Field{
	{Name: "ComputeCapacity", Flag: "compute-capacity", Type: "*types.ComputeCapacity", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisableIMDSV1", Flag: "disable-imdsv1", Type: "*bool", Required: false},
	{Name: "DisconnectTimeoutInSeconds", Flag: "disconnect-timeout-in-seconds", Type: "*int32", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "DomainJoinInfo", Flag: "domain-join-info", Type: "*types.DomainJoinInfo", Required: false},
	{Name: "EnableDefaultInternetAccess", Flag: "enable-default-internet-access", Type: "*bool", Required: false},
	{Name: "FleetType", Flag: "fleet-type", Type: "types.FleetType", Required: false},
	{Name: "IamRoleArn", Flag: "iam-role-arn", Type: "*string", Required: false},
	{Name: "IdleDisconnectTimeoutInSeconds", Flag: "idle-disconnect-timeout-in-seconds", Type: "*int32", Required: false},
	{Name: "ImageArn", Flag: "image-arn", Type: "*string", Required: false},
	{Name: "ImageName", Flag: "image-name", Type: "*string", Required: false},
	{Name: "InstanceType", Flag: "instance-type", Type: "*string", Required: true},
	{Name: "MaxConcurrentSessions", Flag: "max-concurrent-sessions", Type: "*int32", Required: false},
	{Name: "MaxSessionsPerInstance", Flag: "max-sessions-per-instance", Type: "*int32", Required: false},
	{Name: "MaxUserDurationInSeconds", Flag: "max-user-duration-in-seconds", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Platform", Flag: "platform", Type: "types.PlatformType", Required: false},
	{Name: "RootVolumeConfig", Flag: "root-volume-config", Type: "*types.VolumeConfig", Required: false},
	{Name: "SessionScriptS3Location", Flag: "session-script-s3-location", Type: "*types.S3Location", Required: false},
	{Name: "StreamView", Flag: "stream-view", Type: "types.StreamView", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "UsbDeviceFilterStrings", Flag: "usb-device-filter-strings", Type: "[]string", Required: false},
	{Name: "VpcConfig", Flag: "vpc-config", Type: "*types.VpcConfig", Required: false},
}

var fields_create_image_builder = []leanruntime.Field{
	{Name: "AccessEndpoints", Flag: "access-endpoints", Type: "[]types.AccessEndpoint", Required: false},
	{Name: "AppstreamAgentVersion", Flag: "appstream-agent-version", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisableIMDSV1", Flag: "disable-imdsv1", Type: "*bool", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "DomainJoinInfo", Flag: "domain-join-info", Type: "*types.DomainJoinInfo", Required: false},
	{Name: "EnableDefaultInternetAccess", Flag: "enable-default-internet-access", Type: "*bool", Required: false},
	{Name: "IamRoleArn", Flag: "iam-role-arn", Type: "*string", Required: false},
	{Name: "ImageArn", Flag: "image-arn", Type: "*string", Required: false},
	{Name: "ImageName", Flag: "image-name", Type: "*string", Required: false},
	{Name: "InstanceType", Flag: "instance-type", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RootVolumeConfig", Flag: "root-volume-config", Type: "*types.VolumeConfig", Required: false},
	{Name: "SoftwaresToInstall", Flag: "softwares-to-install", Type: "[]string", Required: false},
	{Name: "SoftwaresToUninstall", Flag: "softwares-to-uninstall", Type: "[]string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "VpcConfig", Flag: "vpc-config", Type: "*types.VpcConfig", Required: false},
}

var fields_create_image_builder_streaming_url = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Validity", Flag: "validity", Type: "*int64", Required: false},
}

var fields_create_imported_image = []leanruntime.Field{
	{Name: "AgentSoftwareVersion", Flag: "agent-software-version", Type: "types.AgentSoftwareVersion", Required: false},
	{Name: "AppCatalogConfig", Flag: "app-catalog-config", Type: "[]types.ApplicationConfig", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "IamRoleArn", Flag: "iam-role-arn", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RuntimeValidationConfig", Flag: "runtime-validation-config", Type: "*types.RuntimeValidationConfig", Required: false},
	{Name: "SourceAmiId", Flag: "source-ami-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_stack = []leanruntime.Field{
	{Name: "AccessEndpoints", Flag: "access-endpoints", Type: "[]types.AccessEndpoint", Required: false},
	{Name: "ApplicationSettings", Flag: "application-settings", Type: "*types.ApplicationSettings", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "EmbedHostDomains", Flag: "embed-host-domains", Type: "[]string", Required: false},
	{Name: "FeedbackURL", Flag: "feedback-url", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RedirectURL", Flag: "redirect-url", Type: "*string", Required: false},
	{Name: "StorageConnectors", Flag: "storage-connectors", Type: "[]types.StorageConnector", Required: false},
	{Name: "StreamingExperienceSettings", Flag: "streaming-experience-settings", Type: "*types.StreamingExperienceSettings", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "UserSettings", Flag: "user-settings", Type: "[]types.UserSetting", Required: false},
}

var fields_create_streaming_url = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: false},
	{Name: "FleetName", Flag: "fleet-name", Type: "*string", Required: true},
	{Name: "SessionContext", Flag: "session-context", Type: "*string", Required: false},
	{Name: "StackName", Flag: "stack-name", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
	{Name: "Validity", Flag: "validity", Type: "*int64", Required: false},
}

var fields_create_theme_for_stack = []leanruntime.Field{
	{Name: "FaviconS3Location", Flag: "favicon-s3-location", Type: "*types.S3Location", Required: true},
	{Name: "FooterLinks", Flag: "footer-links", Type: "[]types.ThemeFooterLink", Required: false},
	{Name: "OrganizationLogoS3Location", Flag: "organization-logo-s3-location", Type: "*types.S3Location", Required: true},
	{Name: "StackName", Flag: "stack-name", Type: "*string", Required: true},
	{Name: "ThemeStyling", Flag: "theme-styling", Type: "types.ThemeStyling", Required: true},
	{Name: "TitleText", Flag: "title-text", Type: "*string", Required: true},
}

var fields_create_updated_image = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ExistingImageName", Flag: "existing-image-name", Type: "*string", Required: true},
	{Name: "NewImageDescription", Flag: "new-image-description", Type: "*string", Required: false},
	{Name: "NewImageDisplayName", Flag: "new-image-display-name", Type: "*string", Required: false},
	{Name: "NewImageName", Flag: "new-image-name", Type: "*string", Required: true},
	{Name: "NewImageTags", Flag: "new-image-tags", Type: "map[string]string", Required: false},
}

var fields_create_usage_report_subscription = []leanruntime.Field{}

var fields_create_user = []leanruntime.Field{
	{Name: "AuthenticationType", Flag: "authentication-type", Type: "types.AuthenticationType", Required: true},
	{Name: "FirstName", Flag: "first-name", Type: "*string", Required: false},
	{Name: "LastName", Flag: "last-name", Type: "*string", Required: false},
	{Name: "MessageAction", Flag: "message-action", Type: "types.MessageAction", Required: false},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_delete_app_block = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_app_block_builder = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_application = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_directory_config = []leanruntime.Field{
	{Name: "DirectoryName", Flag: "directory-name", Type: "*string", Required: true},
}

var fields_delete_entitlement = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "StackName", Flag: "stack-name", Type: "*string", Required: true},
}

var fields_delete_fleet = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_image = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_image_builder = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_image_permissions = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "SharedAccountId", Flag: "shared-account-id", Type: "*string", Required: true},
}

var fields_delete_stack = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_theme_for_stack = []leanruntime.Field{
	{Name: "StackName", Flag: "stack-name", Type: "*string", Required: true},
}

var fields_delete_usage_report_subscription = []leanruntime.Field{}

var fields_delete_user = []leanruntime.Field{
	{Name: "AuthenticationType", Flag: "authentication-type", Type: "types.AuthenticationType", Required: true},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_describe_app_block_builder_app_block_associations = []leanruntime.Field{
	{Name: "AppBlockArn", Flag: "app-block-arn", Type: "*string", Required: false},
	{Name: "AppBlockBuilderName", Flag: "app-block-builder-name", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_app_block_builders = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Names", Flag: "names", Type: "[]string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_app_blocks = []leanruntime.Field{
	{Name: "Arns", Flag: "arns", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_app_license_usage = []leanruntime.Field{
	{Name: "BillingPeriod", Flag: "billing-period", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_application_fleet_associations = []leanruntime.Field{
	{Name: "ApplicationArn", Flag: "application-arn", Type: "*string", Required: false},
	{Name: "FleetName", Flag: "fleet-name", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_applications = []leanruntime.Field{
	{Name: "Arns", Flag: "arns", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_directory_configs = []leanruntime.Field{
	{Name: "DirectoryNames", Flag: "directory-names", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_entitlements = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StackName", Flag: "stack-name", Type: "*string", Required: true},
}

var fields_describe_fleets = []leanruntime.Field{
	{Name: "Names", Flag: "names", Type: "[]string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_image_builders = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Names", Flag: "names", Type: "[]string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_image_permissions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SharedAwsAccountIds", Flag: "shared-aws-account-ids", Type: "[]string", Required: false},
}

var fields_describe_images = []leanruntime.Field{
	{Name: "Arns", Flag: "arns", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Names", Flag: "names", Type: "[]string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.VisibilityType", Required: false},
}

var fields_describe_sessions = []leanruntime.Field{
	{Name: "AuthenticationType", Flag: "authentication-type", Type: "types.AuthenticationType", Required: false},
	{Name: "FleetName", Flag: "fleet-name", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StackName", Flag: "stack-name", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: false},
}

var fields_describe_software_associations = []leanruntime.Field{
	{Name: "AssociatedResource", Flag: "associated-resource", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_stacks = []leanruntime.Field{
	{Name: "Names", Flag: "names", Type: "[]string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_theme_for_stack = []leanruntime.Field{
	{Name: "StackName", Flag: "stack-name", Type: "*string", Required: true},
}

var fields_describe_usage_report_subscriptions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_user_stack_associations = []leanruntime.Field{
	{Name: "AuthenticationType", Flag: "authentication-type", Type: "types.AuthenticationType", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StackName", Flag: "stack-name", Type: "*string", Required: false},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: false},
}

var fields_describe_users = []leanruntime.Field{
	{Name: "AuthenticationType", Flag: "authentication-type", Type: "types.AuthenticationType", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_disable_user = []leanruntime.Field{
	{Name: "AuthenticationType", Flag: "authentication-type", Type: "types.AuthenticationType", Required: true},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_disassociate_app_block_builder_app_block = []leanruntime.Field{
	{Name: "AppBlockArn", Flag: "app-block-arn", Type: "*string", Required: true},
	{Name: "AppBlockBuilderName", Flag: "app-block-builder-name", Type: "*string", Required: true},
}

var fields_disassociate_application_fleet = []leanruntime.Field{
	{Name: "ApplicationArn", Flag: "application-arn", Type: "*string", Required: true},
	{Name: "FleetName", Flag: "fleet-name", Type: "*string", Required: true},
}

var fields_disassociate_application_from_entitlement = []leanruntime.Field{
	{Name: "ApplicationIdentifier", Flag: "application-identifier", Type: "*string", Required: true},
	{Name: "EntitlementName", Flag: "entitlement-name", Type: "*string", Required: true},
	{Name: "StackName", Flag: "stack-name", Type: "*string", Required: true},
}

var fields_disassociate_fleet = []leanruntime.Field{
	{Name: "FleetName", Flag: "fleet-name", Type: "*string", Required: true},
	{Name: "StackName", Flag: "stack-name", Type: "*string", Required: true},
}

var fields_disassociate_software_from_image_builder = []leanruntime.Field{
	{Name: "ImageBuilderName", Flag: "image-builder-name", Type: "*string", Required: true},
	{Name: "SoftwareNames", Flag: "software-names", Type: "[]string", Required: true},
}

var fields_enable_user = []leanruntime.Field{
	{Name: "AuthenticationType", Flag: "authentication-type", Type: "types.AuthenticationType", Required: true},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_expire_session = []leanruntime.Field{
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
}

var fields_get_export_image_task = []leanruntime.Field{
	{Name: "TaskId", Flag: "task-id", Type: "*string", Required: false},
}

var fields_list_associated_fleets = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StackName", Flag: "stack-name", Type: "*string", Required: true},
}

var fields_list_associated_stacks = []leanruntime.Field{
	{Name: "FleetName", Flag: "fleet-name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_entitled_applications = []leanruntime.Field{
	{Name: "EntitlementName", Flag: "entitlement-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StackName", Flag: "stack-name", Type: "*string", Required: true},
}

var fields_list_export_image_tasks = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_start_app_block_builder = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_start_fleet = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_start_image_builder = []leanruntime.Field{
	{Name: "AppstreamAgentVersion", Flag: "appstream-agent-version", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_start_software_deployment_to_image_builder = []leanruntime.Field{
	{Name: "ImageBuilderName", Flag: "image-builder-name", Type: "*string", Required: true},
	{Name: "RetryFailedDeployments", Flag: "retry-failed-deployments", Type: "*bool", Required: false},
}

var fields_stop_app_block_builder = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_stop_fleet = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_stop_image_builder = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_app_block_builder = []leanruntime.Field{
	{Name: "AccessEndpoints", Flag: "access-endpoints", Type: "[]types.AccessEndpoint", Required: false},
	{Name: "AttributesToDelete", Flag: "attributes-to-delete", Type: "[]types.AppBlockBuilderAttribute", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisableIMDSV1", Flag: "disable-imdsv1", Type: "*bool", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "EnableDefaultInternetAccess", Flag: "enable-default-internet-access", Type: "*bool", Required: false},
	{Name: "IamRoleArn", Flag: "iam-role-arn", Type: "*string", Required: false},
	{Name: "InstanceType", Flag: "instance-type", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Platform", Flag: "platform", Type: "types.PlatformType", Required: false},
	{Name: "VpcConfig", Flag: "vpc-config", Type: "*types.VpcConfig", Required: false},
}

var fields_update_application = []leanruntime.Field{
	{Name: "AppBlockArn", Flag: "app-block-arn", Type: "*string", Required: false},
	{Name: "AttributesToDelete", Flag: "attributes-to-delete", Type: "[]types.ApplicationAttribute", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "IconS3Location", Flag: "icon-s3-location", Type: "*types.S3Location", Required: false},
	{Name: "LaunchParameters", Flag: "launch-parameters", Type: "*string", Required: false},
	{Name: "LaunchPath", Flag: "launch-path", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "WorkingDirectory", Flag: "working-directory", Type: "*string", Required: false},
}

var fields_update_directory_config = []leanruntime.Field{
	{Name: "CertificateBasedAuthProperties", Flag: "certificate-based-auth-properties", Type: "*types.CertificateBasedAuthProperties", Required: false},
	{Name: "DirectoryName", Flag: "directory-name", Type: "*string", Required: true},
	{Name: "OrganizationalUnitDistinguishedNames", Flag: "organizational-unit-distinguished-names", Type: "[]string", Required: false},
	{Name: "ServiceAccountCredentials", Flag: "service-account-credentials", Type: "*types.ServiceAccountCredentials", Required: false},
}

var fields_update_entitlement = []leanruntime.Field{
	{Name: "AppVisibility", Flag: "app-visibility", Type: "types.AppVisibility", Required: false},
	{Name: "Attributes", Flag: "attributes", Type: "[]types.EntitlementAttribute", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "StackName", Flag: "stack-name", Type: "*string", Required: true},
}

var fields_update_fleet = []leanruntime.Field{
	{Name: "AttributesToDelete", Flag: "attributes-to-delete", Type: "[]types.FleetAttribute", Required: false},
	{Name: "ComputeCapacity", Flag: "compute-capacity", Type: "*types.ComputeCapacity", Required: false},
	{Name: "DeleteVpcConfig", Flag: "delete-vpc-config", Type: "*bool", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisableIMDSV1", Flag: "disable-imdsv1", Type: "*bool", Required: false},
	{Name: "DisconnectTimeoutInSeconds", Flag: "disconnect-timeout-in-seconds", Type: "*int32", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "DomainJoinInfo", Flag: "domain-join-info", Type: "*types.DomainJoinInfo", Required: false},
	{Name: "EnableDefaultInternetAccess", Flag: "enable-default-internet-access", Type: "*bool", Required: false},
	{Name: "IamRoleArn", Flag: "iam-role-arn", Type: "*string", Required: false},
	{Name: "IdleDisconnectTimeoutInSeconds", Flag: "idle-disconnect-timeout-in-seconds", Type: "*int32", Required: false},
	{Name: "ImageArn", Flag: "image-arn", Type: "*string", Required: false},
	{Name: "ImageName", Flag: "image-name", Type: "*string", Required: false},
	{Name: "InstanceType", Flag: "instance-type", Type: "*string", Required: false},
	{Name: "MaxConcurrentSessions", Flag: "max-concurrent-sessions", Type: "*int32", Required: false},
	{Name: "MaxSessionsPerInstance", Flag: "max-sessions-per-instance", Type: "*int32", Required: false},
	{Name: "MaxUserDurationInSeconds", Flag: "max-user-duration-in-seconds", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Platform", Flag: "platform", Type: "types.PlatformType", Required: false},
	{Name: "RootVolumeConfig", Flag: "root-volume-config", Type: "*types.VolumeConfig", Required: false},
	{Name: "SessionScriptS3Location", Flag: "session-script-s3-location", Type: "*types.S3Location", Required: false},
	{Name: "StreamView", Flag: "stream-view", Type: "types.StreamView", Required: false},
	{Name: "UsbDeviceFilterStrings", Flag: "usb-device-filter-strings", Type: "[]string", Required: false},
	{Name: "VpcConfig", Flag: "vpc-config", Type: "*types.VpcConfig", Required: false},
}

var fields_update_image_permissions = []leanruntime.Field{
	{Name: "ImagePermissions", Flag: "image-permissions", Type: "*types.ImagePermissions", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "SharedAccountId", Flag: "shared-account-id", Type: "*string", Required: true},
}

var fields_update_stack = []leanruntime.Field{
	{Name: "AccessEndpoints", Flag: "access-endpoints", Type: "[]types.AccessEndpoint", Required: false},
	{Name: "ApplicationSettings", Flag: "application-settings", Type: "*types.ApplicationSettings", Required: false},
	{Name: "AttributesToDelete", Flag: "attributes-to-delete", Type: "[]types.StackAttribute", Required: false},
	{Name: "DeleteStorageConnectors", Flag: "delete-storage-connectors", Type: "*bool", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "EmbedHostDomains", Flag: "embed-host-domains", Type: "[]string", Required: false},
	{Name: "FeedbackURL", Flag: "feedback-url", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RedirectURL", Flag: "redirect-url", Type: "*string", Required: false},
	{Name: "StorageConnectors", Flag: "storage-connectors", Type: "[]types.StorageConnector", Required: false},
	{Name: "StreamingExperienceSettings", Flag: "streaming-experience-settings", Type: "*types.StreamingExperienceSettings", Required: false},
	{Name: "UserSettings", Flag: "user-settings", Type: "[]types.UserSetting", Required: false},
}

var fields_update_theme_for_stack = []leanruntime.Field{
	{Name: "AttributesToDelete", Flag: "attributes-to-delete", Type: "[]types.ThemeAttribute", Required: false},
	{Name: "FaviconS3Location", Flag: "favicon-s3-location", Type: "*types.S3Location", Required: false},
	{Name: "FooterLinks", Flag: "footer-links", Type: "[]types.ThemeFooterLink", Required: false},
	{Name: "OrganizationLogoS3Location", Flag: "organization-logo-s3-location", Type: "*types.S3Location", Required: false},
	{Name: "StackName", Flag: "stack-name", Type: "*string", Required: true},
	{Name: "State", Flag: "state", Type: "types.ThemeState", Required: false},
	{Name: "ThemeStyling", Flag: "theme-styling", Type: "types.ThemeStyling", Required: false},
	{Name: "TitleText", Flag: "title-text", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-app-block-builder-app-block": {
			Name:   "associate-app-block-builder-app-block",
			Fields: fields_associate_app_block_builder_app_block,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateAppBlockBuilderAppBlockInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_app_block_builder_app_block, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateAppBlockBuilderAppBlock(ctx, input)
			},
		},
		"associate-application-fleet": {
			Name:   "associate-application-fleet",
			Fields: fields_associate_application_fleet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateApplicationFleetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_application_fleet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateApplicationFleet(ctx, input)
			},
		},
		"associate-application-to-entitlement": {
			Name:   "associate-application-to-entitlement",
			Fields: fields_associate_application_to_entitlement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateApplicationToEntitlementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_application_to_entitlement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateApplicationToEntitlement(ctx, input)
			},
		},
		"associate-fleet": {
			Name:   "associate-fleet",
			Fields: fields_associate_fleet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateFleetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_fleet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateFleet(ctx, input)
			},
		},
		"associate-software-to-image-builder": {
			Name:   "associate-software-to-image-builder",
			Fields: fields_associate_software_to_image_builder,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateSoftwareToImageBuilderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_software_to_image_builder, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateSoftwareToImageBuilder(ctx, input)
			},
		},
		"batch-associate-user-stack": {
			Name:   "batch-associate-user-stack",
			Fields: fields_batch_associate_user_stack,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchAssociateUserStackInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_associate_user_stack, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchAssociateUserStack(ctx, input)
			},
		},
		"batch-disassociate-user-stack": {
			Name:   "batch-disassociate-user-stack",
			Fields: fields_batch_disassociate_user_stack,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDisassociateUserStackInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_disassociate_user_stack, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDisassociateUserStack(ctx, input)
			},
		},
		"copy-image": {
			Name:   "copy-image",
			Fields: fields_copy_image,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CopyImageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_copy_image, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CopyImage(ctx, input)
			},
		},
		"create-app-block": {
			Name:   "create-app-block",
			Fields: fields_create_app_block,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAppBlockInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_app_block, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAppBlock(ctx, input)
			},
		},
		"create-app-block-builder": {
			Name:   "create-app-block-builder",
			Fields: fields_create_app_block_builder,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAppBlockBuilderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_app_block_builder, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAppBlockBuilder(ctx, input)
			},
		},
		"create-app-block-builder-streaming-url": {
			Name:   "create-app-block-builder-streaming-url",
			Fields: fields_create_app_block_builder_streaming_url,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAppBlockBuilderStreamingURLInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_app_block_builder_streaming_url, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAppBlockBuilderStreamingURL(ctx, input)
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
		"create-directory-config": {
			Name:   "create-directory-config",
			Fields: fields_create_directory_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDirectoryConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_directory_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDirectoryConfig(ctx, input)
			},
		},
		"create-entitlement": {
			Name:   "create-entitlement",
			Fields: fields_create_entitlement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEntitlementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_entitlement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEntitlement(ctx, input)
			},
		},
		"create-export-image-task": {
			Name:   "create-export-image-task",
			Fields: fields_create_export_image_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateExportImageTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_export_image_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateExportImageTask(ctx, input)
			},
		},
		"create-fleet": {
			Name:   "create-fleet",
			Fields: fields_create_fleet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFleetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_fleet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFleet(ctx, input)
			},
		},
		"create-image-builder": {
			Name:   "create-image-builder",
			Fields: fields_create_image_builder,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateImageBuilderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_image_builder, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateImageBuilder(ctx, input)
			},
		},
		"create-image-builder-streaming-url": {
			Name:   "create-image-builder-streaming-url",
			Fields: fields_create_image_builder_streaming_url,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateImageBuilderStreamingURLInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_image_builder_streaming_url, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateImageBuilderStreamingURL(ctx, input)
			},
		},
		"create-imported-image": {
			Name:   "create-imported-image",
			Fields: fields_create_imported_image,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateImportedImageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_imported_image, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateImportedImage(ctx, input)
			},
		},
		"create-stack": {
			Name:   "create-stack",
			Fields: fields_create_stack,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateStackInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_stack, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateStack(ctx, input)
			},
		},
		"create-streaming-url": {
			Name:   "create-streaming-url",
			Fields: fields_create_streaming_url,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateStreamingURLInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_streaming_url, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateStreamingURL(ctx, input)
			},
		},
		"create-theme-for-stack": {
			Name:   "create-theme-for-stack",
			Fields: fields_create_theme_for_stack,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateThemeForStackInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_theme_for_stack, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateThemeForStack(ctx, input)
			},
		},
		"create-updated-image": {
			Name:   "create-updated-image",
			Fields: fields_create_updated_image,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateUpdatedImageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_updated_image, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateUpdatedImage(ctx, input)
			},
		},
		"create-usage-report-subscription": {
			Name:   "create-usage-report-subscription",
			Fields: fields_create_usage_report_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateUsageReportSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_usage_report_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateUsageReportSubscription(ctx, input)
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
		"delete-app-block": {
			Name:   "delete-app-block",
			Fields: fields_delete_app_block,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAppBlockInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_app_block, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAppBlock(ctx, input)
			},
		},
		"delete-app-block-builder": {
			Name:   "delete-app-block-builder",
			Fields: fields_delete_app_block_builder,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAppBlockBuilderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_app_block_builder, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAppBlockBuilder(ctx, input)
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
		"delete-directory-config": {
			Name:   "delete-directory-config",
			Fields: fields_delete_directory_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDirectoryConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_directory_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDirectoryConfig(ctx, input)
			},
		},
		"delete-entitlement": {
			Name:   "delete-entitlement",
			Fields: fields_delete_entitlement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEntitlementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_entitlement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEntitlement(ctx, input)
			},
		},
		"delete-fleet": {
			Name:   "delete-fleet",
			Fields: fields_delete_fleet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFleetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_fleet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFleet(ctx, input)
			},
		},
		"delete-image": {
			Name:   "delete-image",
			Fields: fields_delete_image,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteImageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_image, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteImage(ctx, input)
			},
		},
		"delete-image-builder": {
			Name:   "delete-image-builder",
			Fields: fields_delete_image_builder,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteImageBuilderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_image_builder, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteImageBuilder(ctx, input)
			},
		},
		"delete-image-permissions": {
			Name:   "delete-image-permissions",
			Fields: fields_delete_image_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteImagePermissionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_image_permissions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteImagePermissions(ctx, input)
			},
		},
		"delete-stack": {
			Name:   "delete-stack",
			Fields: fields_delete_stack,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteStackInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_stack, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteStack(ctx, input)
			},
		},
		"delete-theme-for-stack": {
			Name:   "delete-theme-for-stack",
			Fields: fields_delete_theme_for_stack,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteThemeForStackInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_theme_for_stack, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteThemeForStack(ctx, input)
			},
		},
		"delete-usage-report-subscription": {
			Name:   "delete-usage-report-subscription",
			Fields: fields_delete_usage_report_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteUsageReportSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_usage_report_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteUsageReportSubscription(ctx, input)
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
		"describe-app-block-builder-app-block-associations": {
			Name:   "describe-app-block-builder-app-block-associations",
			Fields: fields_describe_app_block_builder_app_block_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAppBlockBuilderAppBlockAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_app_block_builder_app_block_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeAppBlockBuilderAppBlockAssociations(ctx, input)
				}
				var results []*svc.DescribeAppBlockBuilderAppBlockAssociationsOutput
				p := svc.NewDescribeAppBlockBuilderAppBlockAssociationsPaginator(client, input)
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
		"describe-app-block-builders": {
			Name:   "describe-app-block-builders",
			Fields: fields_describe_app_block_builders,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAppBlockBuildersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_app_block_builders, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeAppBlockBuilders(ctx, input)
				}
				var results []*svc.DescribeAppBlockBuildersOutput
				p := svc.NewDescribeAppBlockBuildersPaginator(client, input)
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
		"describe-app-blocks": {
			Name:   "describe-app-blocks",
			Fields: fields_describe_app_blocks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAppBlocksInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_app_blocks, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAppBlocks(ctx, input)
			},
		},
		"describe-app-license-usage": {
			Name:   "describe-app-license-usage",
			Fields: fields_describe_app_license_usage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAppLicenseUsageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_app_license_usage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAppLicenseUsage(ctx, input)
			},
		},
		"describe-application-fleet-associations": {
			Name:   "describe-application-fleet-associations",
			Fields: fields_describe_application_fleet_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeApplicationFleetAssociationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_application_fleet_associations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeApplicationFleetAssociations(ctx, input)
			},
		},
		"describe-applications": {
			Name:   "describe-applications",
			Fields: fields_describe_applications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeApplicationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_applications, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeApplications(ctx, input)
			},
		},
		"describe-directory-configs": {
			Name:   "describe-directory-configs",
			Fields: fields_describe_directory_configs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDirectoryConfigsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_directory_configs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDirectoryConfigs(ctx, input)
			},
		},
		"describe-entitlements": {
			Name:   "describe-entitlements",
			Fields: fields_describe_entitlements,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEntitlementsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_entitlements, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEntitlements(ctx, input)
			},
		},
		"describe-fleets": {
			Name:   "describe-fleets",
			Fields: fields_describe_fleets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFleetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_fleets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeFleets(ctx, input)
			},
		},
		"describe-image-builders": {
			Name:   "describe-image-builders",
			Fields: fields_describe_image_builders,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeImageBuildersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_image_builders, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeImageBuilders(ctx, input)
			},
		},
		"describe-image-permissions": {
			Name:   "describe-image-permissions",
			Fields: fields_describe_image_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeImagePermissionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_image_permissions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeImagePermissions(ctx, input)
				}
				var results []*svc.DescribeImagePermissionsOutput
				p := svc.NewDescribeImagePermissionsPaginator(client, input)
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
		"describe-images": {
			Name:   "describe-images",
			Fields: fields_describe_images,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeImagesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_images, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeImages(ctx, input)
				}
				var results []*svc.DescribeImagesOutput
				p := svc.NewDescribeImagesPaginator(client, input)
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
		"describe-sessions": {
			Name:   "describe-sessions",
			Fields: fields_describe_sessions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSessionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_sessions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSessions(ctx, input)
			},
		},
		"describe-software-associations": {
			Name:   "describe-software-associations",
			Fields: fields_describe_software_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSoftwareAssociationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_software_associations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSoftwareAssociations(ctx, input)
			},
		},
		"describe-stacks": {
			Name:   "describe-stacks",
			Fields: fields_describe_stacks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeStacksInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_stacks, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeStacks(ctx, input)
			},
		},
		"describe-theme-for-stack": {
			Name:   "describe-theme-for-stack",
			Fields: fields_describe_theme_for_stack,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeThemeForStackInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_theme_for_stack, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeThemeForStack(ctx, input)
			},
		},
		"describe-usage-report-subscriptions": {
			Name:   "describe-usage-report-subscriptions",
			Fields: fields_describe_usage_report_subscriptions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeUsageReportSubscriptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_usage_report_subscriptions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeUsageReportSubscriptions(ctx, input)
			},
		},
		"describe-user-stack-associations": {
			Name:   "describe-user-stack-associations",
			Fields: fields_describe_user_stack_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeUserStackAssociationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_user_stack_associations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeUserStackAssociations(ctx, input)
			},
		},
		"describe-users": {
			Name:   "describe-users",
			Fields: fields_describe_users,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeUsersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_users, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeUsers(ctx, input)
			},
		},
		"disable-user": {
			Name:   "disable-user",
			Fields: fields_disable_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableUser(ctx, input)
			},
		},
		"disassociate-app-block-builder-app-block": {
			Name:   "disassociate-app-block-builder-app-block",
			Fields: fields_disassociate_app_block_builder_app_block,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateAppBlockBuilderAppBlockInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_app_block_builder_app_block, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateAppBlockBuilderAppBlock(ctx, input)
			},
		},
		"disassociate-application-fleet": {
			Name:   "disassociate-application-fleet",
			Fields: fields_disassociate_application_fleet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateApplicationFleetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_application_fleet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateApplicationFleet(ctx, input)
			},
		},
		"disassociate-application-from-entitlement": {
			Name:   "disassociate-application-from-entitlement",
			Fields: fields_disassociate_application_from_entitlement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateApplicationFromEntitlementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_application_from_entitlement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateApplicationFromEntitlement(ctx, input)
			},
		},
		"disassociate-fleet": {
			Name:   "disassociate-fleet",
			Fields: fields_disassociate_fleet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateFleetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_fleet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateFleet(ctx, input)
			},
		},
		"disassociate-software-from-image-builder": {
			Name:   "disassociate-software-from-image-builder",
			Fields: fields_disassociate_software_from_image_builder,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateSoftwareFromImageBuilderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_software_from_image_builder, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateSoftwareFromImageBuilder(ctx, input)
			},
		},
		"enable-user": {
			Name:   "enable-user",
			Fields: fields_enable_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableUser(ctx, input)
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
		"get-export-image-task": {
			Name:   "get-export-image-task",
			Fields: fields_get_export_image_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetExportImageTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_export_image_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetExportImageTask(ctx, input)
			},
		},
		"list-associated-fleets": {
			Name:   "list-associated-fleets",
			Fields: fields_list_associated_fleets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAssociatedFleetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_associated_fleets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListAssociatedFleets(ctx, input)
			},
		},
		"list-associated-stacks": {
			Name:   "list-associated-stacks",
			Fields: fields_list_associated_stacks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAssociatedStacksInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_associated_stacks, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListAssociatedStacks(ctx, input)
			},
		},
		"list-entitled-applications": {
			Name:   "list-entitled-applications",
			Fields: fields_list_entitled_applications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEntitledApplicationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_entitled_applications, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListEntitledApplications(ctx, input)
			},
		},
		"list-export-image-tasks": {
			Name:   "list-export-image-tasks",
			Fields: fields_list_export_image_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListExportImageTasksInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_export_image_tasks, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListExportImageTasks(ctx, input)
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
		"start-app-block-builder": {
			Name:   "start-app-block-builder",
			Fields: fields_start_app_block_builder,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartAppBlockBuilderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_app_block_builder, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartAppBlockBuilder(ctx, input)
			},
		},
		"start-fleet": {
			Name:   "start-fleet",
			Fields: fields_start_fleet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartFleetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_fleet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartFleet(ctx, input)
			},
		},
		"start-image-builder": {
			Name:   "start-image-builder",
			Fields: fields_start_image_builder,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartImageBuilderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_image_builder, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartImageBuilder(ctx, input)
			},
		},
		"start-software-deployment-to-image-builder": {
			Name:   "start-software-deployment-to-image-builder",
			Fields: fields_start_software_deployment_to_image_builder,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartSoftwareDeploymentToImageBuilderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_software_deployment_to_image_builder, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartSoftwareDeploymentToImageBuilder(ctx, input)
			},
		},
		"stop-app-block-builder": {
			Name:   "stop-app-block-builder",
			Fields: fields_stop_app_block_builder,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopAppBlockBuilderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_app_block_builder, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopAppBlockBuilder(ctx, input)
			},
		},
		"stop-fleet": {
			Name:   "stop-fleet",
			Fields: fields_stop_fleet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopFleetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_fleet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopFleet(ctx, input)
			},
		},
		"stop-image-builder": {
			Name:   "stop-image-builder",
			Fields: fields_stop_image_builder,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopImageBuilderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_image_builder, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopImageBuilder(ctx, input)
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
		"update-app-block-builder": {
			Name:   "update-app-block-builder",
			Fields: fields_update_app_block_builder,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAppBlockBuilderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_app_block_builder, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAppBlockBuilder(ctx, input)
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
		"update-directory-config": {
			Name:   "update-directory-config",
			Fields: fields_update_directory_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDirectoryConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_directory_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDirectoryConfig(ctx, input)
			},
		},
		"update-entitlement": {
			Name:   "update-entitlement",
			Fields: fields_update_entitlement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEntitlementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_entitlement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEntitlement(ctx, input)
			},
		},
		"update-fleet": {
			Name:   "update-fleet",
			Fields: fields_update_fleet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFleetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_fleet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFleet(ctx, input)
			},
		},
		"update-image-permissions": {
			Name:   "update-image-permissions",
			Fields: fields_update_image_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateImagePermissionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_image_permissions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateImagePermissions(ctx, input)
			},
		},
		"update-stack": {
			Name:   "update-stack",
			Fields: fields_update_stack,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateStackInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_stack, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateStack(ctx, input)
			},
		},
		"update-theme-for-stack": {
			Name:   "update-theme-for-stack",
			Fields: fields_update_theme_for_stack,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateThemeForStackInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_theme_for_stack, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateThemeForStack(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("appstream", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
