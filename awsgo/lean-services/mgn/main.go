package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/mgn"
)

var fields_archive_application = []leanruntime.Field{
	{Name: "AccountID", Flag: "account-id", Type: "*string", Required: false},
	{Name: "ApplicationID", Flag: "application-id", Type: "*string", Required: true},
}

var fields_archive_wave = []leanruntime.Field{
	{Name: "AccountID", Flag: "account-id", Type: "*string", Required: false},
	{Name: "WaveID", Flag: "wave-id", Type: "*string", Required: true},
}

var fields_associate_applications = []leanruntime.Field{
	{Name: "AccountID", Flag: "account-id", Type: "*string", Required: false},
	{Name: "ApplicationIDs", Flag: "application-ids", Type: "[]string", Required: true},
	{Name: "WaveID", Flag: "wave-id", Type: "*string", Required: true},
}

var fields_associate_source_servers = []leanruntime.Field{
	{Name: "AccountID", Flag: "account-id", Type: "*string", Required: false},
	{Name: "ApplicationID", Flag: "application-id", Type: "*string", Required: true},
	{Name: "SourceServerIDs", Flag: "source-server-ids", Type: "[]string", Required: true},
}

var fields_change_server_life_cycle_state = []leanruntime.Field{
	{Name: "AccountID", Flag: "account-id", Type: "*string", Required: false},
	{Name: "LifeCycle", Flag: "life-cycle", Type: "*types.ChangeServerLifeCycleStateSourceServerLifecycle", Required: true},
	{Name: "SourceServerID", Flag: "source-server-id", Type: "*string", Required: true},
}

var fields_create_application = []leanruntime.Field{
	{Name: "AccountID", Flag: "account-id", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_connector = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "SsmCommandConfig", Flag: "ssm-command-config", Type: "*types.ConnectorSsmCommandConfig", Required: false},
	{Name: "SsmInstanceID", Flag: "ssm-instance-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_launch_configuration_template = []leanruntime.Field{
	{Name: "AssociatePublicIpAddress", Flag: "associate-public-ip-address", Type: "*bool", Required: false},
	{Name: "BootMode", Flag: "boot-mode", Type: "types.BootMode", Required: false},
	{Name: "CopyPrivateIp", Flag: "copy-private-ip", Type: "*bool", Required: false},
	{Name: "CopyTags", Flag: "copy-tags", Type: "*bool", Required: false},
	{Name: "EnableMapAutoTagging", Flag: "enable-map-auto-tagging", Type: "*bool", Required: false},
	{Name: "EnableParametersEncryption", Flag: "enable-parameters-encryption", Type: "*bool", Required: false},
	{Name: "LargeVolumeConf", Flag: "large-volume-conf", Type: "*types.LaunchTemplateDiskConf", Required: false},
	{Name: "LaunchDisposition", Flag: "launch-disposition", Type: "types.LaunchDisposition", Required: false},
	{Name: "Licensing", Flag: "licensing", Type: "*types.Licensing", Required: false},
	{Name: "MapAutoTaggingMpeID", Flag: "map-auto-tagging-mpe-id", Type: "*string", Required: false},
	{Name: "ParametersEncryptionKey", Flag: "parameters-encryption-key", Type: "*string", Required: false},
	{Name: "PostLaunchActions", Flag: "post-launch-actions", Type: "*types.PostLaunchActions", Required: false},
	{Name: "SmallVolumeConf", Flag: "small-volume-conf", Type: "*types.LaunchTemplateDiskConf", Required: false},
	{Name: "SmallVolumeMaxSize", Flag: "small-volume-max-size", Type: "int64", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TargetInstanceTypeRightSizingMethod", Flag: "target-instance-type-right-sizing-method", Type: "types.TargetInstanceTypeRightSizingMethod", Required: false},
}

var fields_create_replication_configuration_template = []leanruntime.Field{
	{Name: "AssociateDefaultSecurityGroup", Flag: "associate-default-security-group", Type: "*bool", Required: true},
	{Name: "BandwidthThrottling", Flag: "bandwidth-throttling", Type: "int64", Required: true},
	{Name: "CreatePublicIP", Flag: "create-public-ip", Type: "*bool", Required: true},
	{Name: "DataPlaneRouting", Flag: "data-plane-routing", Type: "types.ReplicationConfigurationDataPlaneRouting", Required: true},
	{Name: "DefaultLargeStagingDiskType", Flag: "default-large-staging-disk-type", Type: "types.ReplicationConfigurationDefaultLargeStagingDiskType", Required: true},
	{Name: "EbsEncryption", Flag: "ebs-encryption", Type: "types.ReplicationConfigurationEbsEncryption", Required: true},
	{Name: "EbsEncryptionKeyArn", Flag: "ebs-encryption-key-arn", Type: "*string", Required: false},
	{Name: "InternetProtocol", Flag: "internet-protocol", Type: "types.InternetProtocol", Required: false},
	{Name: "ReplicationServerInstanceType", Flag: "replication-server-instance-type", Type: "*string", Required: true},
	{Name: "ReplicationServersSecurityGroupsIDs", Flag: "replication-servers-security-groups-ids", Type: "[]string", Required: true},
	{Name: "StagingAreaSubnetId", Flag: "staging-area-subnet-id", Type: "*string", Required: true},
	{Name: "StagingAreaTags", Flag: "staging-area-tags", Type: "map[string]string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "UseDedicatedReplicationServer", Flag: "use-dedicated-replication-server", Type: "*bool", Required: true},
	{Name: "UseFipsEndpoint", Flag: "use-fips-endpoint", Type: "*bool", Required: false},
}

var fields_create_wave = []leanruntime.Field{
	{Name: "AccountID", Flag: "account-id", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_application = []leanruntime.Field{
	{Name: "AccountID", Flag: "account-id", Type: "*string", Required: false},
	{Name: "ApplicationID", Flag: "application-id", Type: "*string", Required: true},
}

var fields_delete_connector = []leanruntime.Field{
	{Name: "ConnectorID", Flag: "connector-id", Type: "*string", Required: true},
}

var fields_delete_job = []leanruntime.Field{
	{Name: "AccountID", Flag: "account-id", Type: "*string", Required: false},
	{Name: "JobID", Flag: "job-id", Type: "*string", Required: true},
}

var fields_delete_launch_configuration_template = []leanruntime.Field{
	{Name: "LaunchConfigurationTemplateID", Flag: "launch-configuration-template-id", Type: "*string", Required: true},
}

var fields_delete_replication_configuration_template = []leanruntime.Field{
	{Name: "ReplicationConfigurationTemplateID", Flag: "replication-configuration-template-id", Type: "*string", Required: true},
}

var fields_delete_source_server = []leanruntime.Field{
	{Name: "AccountID", Flag: "account-id", Type: "*string", Required: false},
	{Name: "SourceServerID", Flag: "source-server-id", Type: "*string", Required: true},
}

var fields_delete_vcenter_client = []leanruntime.Field{
	{Name: "VcenterClientID", Flag: "vcenter-client-id", Type: "*string", Required: true},
}

var fields_delete_wave = []leanruntime.Field{
	{Name: "AccountID", Flag: "account-id", Type: "*string", Required: false},
	{Name: "WaveID", Flag: "wave-id", Type: "*string", Required: true},
}

var fields_describe_job_log_items = []leanruntime.Field{
	{Name: "AccountID", Flag: "account-id", Type: "*string", Required: false},
	{Name: "JobID", Flag: "job-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_jobs = []leanruntime.Field{
	{Name: "AccountID", Flag: "account-id", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "*types.DescribeJobsRequestFilters", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_launch_configuration_templates = []leanruntime.Field{
	{Name: "LaunchConfigurationTemplateIDs", Flag: "launch-configuration-template-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_replication_configuration_templates = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ReplicationConfigurationTemplateIDs", Flag: "replication-configuration-template-ids", Type: "[]string", Required: false},
}

var fields_describe_source_servers = []leanruntime.Field{
	{Name: "AccountID", Flag: "account-id", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "*types.DescribeSourceServersRequestFilters", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_vcenter_clients = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_disassociate_applications = []leanruntime.Field{
	{Name: "AccountID", Flag: "account-id", Type: "*string", Required: false},
	{Name: "ApplicationIDs", Flag: "application-ids", Type: "[]string", Required: true},
	{Name: "WaveID", Flag: "wave-id", Type: "*string", Required: true},
}

var fields_disassociate_source_servers = []leanruntime.Field{
	{Name: "AccountID", Flag: "account-id", Type: "*string", Required: false},
	{Name: "ApplicationID", Flag: "application-id", Type: "*string", Required: true},
	{Name: "SourceServerIDs", Flag: "source-server-ids", Type: "[]string", Required: true},
}

var fields_disconnect_from_service = []leanruntime.Field{
	{Name: "AccountID", Flag: "account-id", Type: "*string", Required: false},
	{Name: "SourceServerID", Flag: "source-server-id", Type: "*string", Required: true},
}

var fields_finalize_cutover = []leanruntime.Field{
	{Name: "AccountID", Flag: "account-id", Type: "*string", Required: false},
	{Name: "SourceServerID", Flag: "source-server-id", Type: "*string", Required: true},
}

var fields_get_launch_configuration = []leanruntime.Field{
	{Name: "AccountID", Flag: "account-id", Type: "*string", Required: false},
	{Name: "SourceServerID", Flag: "source-server-id", Type: "*string", Required: true},
}

var fields_get_replication_configuration = []leanruntime.Field{
	{Name: "AccountID", Flag: "account-id", Type: "*string", Required: false},
	{Name: "SourceServerID", Flag: "source-server-id", Type: "*string", Required: true},
}

var fields_initialize_service = []leanruntime.Field{}

var fields_list_applications = []leanruntime.Field{
	{Name: "AccountID", Flag: "account-id", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "*types.ListApplicationsRequestFilters", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_connectors = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "*types.ListConnectorsRequestFilters", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_export_errors = []leanruntime.Field{
	{Name: "ExportID", Flag: "export-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_exports = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "*types.ListExportsRequestFilters", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_import_errors = []leanruntime.Field{
	{Name: "ImportID", Flag: "import-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_imports = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "*types.ListImportsRequestFilters", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_managed_accounts = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_source_server_actions = []leanruntime.Field{
	{Name: "AccountID", Flag: "account-id", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "*types.SourceServerActionsRequestFilters", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SourceServerID", Flag: "source-server-id", Type: "*string", Required: true},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_template_actions = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "*types.TemplateActionsRequestFilters", Required: false},
	{Name: "LaunchConfigurationTemplateID", Flag: "launch-configuration-template-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_waves = []leanruntime.Field{
	{Name: "AccountID", Flag: "account-id", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "*types.ListWavesRequestFilters", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_mark_as_archived = []leanruntime.Field{
	{Name: "AccountID", Flag: "account-id", Type: "*string", Required: false},
	{Name: "SourceServerID", Flag: "source-server-id", Type: "*string", Required: true},
}

var fields_pause_replication = []leanruntime.Field{
	{Name: "AccountID", Flag: "account-id", Type: "*string", Required: false},
	{Name: "SourceServerID", Flag: "source-server-id", Type: "*string", Required: true},
}

var fields_put_source_server_action = []leanruntime.Field{
	{Name: "AccountID", Flag: "account-id", Type: "*string", Required: false},
	{Name: "ActionID", Flag: "action-id", Type: "*string", Required: true},
	{Name: "ActionName", Flag: "action-name", Type: "*string", Required: true},
	{Name: "Active", Flag: "active", Type: "*bool", Required: false},
	{Name: "Category", Flag: "category", Type: "types.ActionCategory", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DocumentIdentifier", Flag: "document-identifier", Type: "*string", Required: true},
	{Name: "DocumentVersion", Flag: "document-version", Type: "*string", Required: false},
	{Name: "ExternalParameters", Flag: "external-parameters", Type: "map[string]types.SsmExternalParameter", Required: false},
	{Name: "MustSucceedForCutover", Flag: "must-succeed-for-cutover", Type: "*bool", Required: false},
	{Name: "Order", Flag: "order", Type: "*int32", Required: true},
	{Name: "Parameters", Flag: "parameters", Type: "map[string][]types.SsmParameterStoreParameter", Required: false},
	{Name: "SourceServerID", Flag: "source-server-id", Type: "*string", Required: true},
	{Name: "TimeoutSeconds", Flag: "timeout-seconds", Type: "*int32", Required: false},
}

var fields_put_template_action = []leanruntime.Field{
	{Name: "ActionID", Flag: "action-id", Type: "*string", Required: true},
	{Name: "ActionName", Flag: "action-name", Type: "*string", Required: true},
	{Name: "Active", Flag: "active", Type: "*bool", Required: false},
	{Name: "Category", Flag: "category", Type: "types.ActionCategory", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DocumentIdentifier", Flag: "document-identifier", Type: "*string", Required: true},
	{Name: "DocumentVersion", Flag: "document-version", Type: "*string", Required: false},
	{Name: "ExternalParameters", Flag: "external-parameters", Type: "map[string]types.SsmExternalParameter", Required: false},
	{Name: "LaunchConfigurationTemplateID", Flag: "launch-configuration-template-id", Type: "*string", Required: true},
	{Name: "MustSucceedForCutover", Flag: "must-succeed-for-cutover", Type: "*bool", Required: false},
	{Name: "OperatingSystem", Flag: "operating-system", Type: "*string", Required: false},
	{Name: "Order", Flag: "order", Type: "*int32", Required: true},
	{Name: "Parameters", Flag: "parameters", Type: "map[string][]types.SsmParameterStoreParameter", Required: false},
	{Name: "TimeoutSeconds", Flag: "timeout-seconds", Type: "*int32", Required: false},
}

var fields_remove_source_server_action = []leanruntime.Field{
	{Name: "AccountID", Flag: "account-id", Type: "*string", Required: false},
	{Name: "ActionID", Flag: "action-id", Type: "*string", Required: true},
	{Name: "SourceServerID", Flag: "source-server-id", Type: "*string", Required: true},
}

var fields_remove_template_action = []leanruntime.Field{
	{Name: "ActionID", Flag: "action-id", Type: "*string", Required: true},
	{Name: "LaunchConfigurationTemplateID", Flag: "launch-configuration-template-id", Type: "*string", Required: true},
}

var fields_resume_replication = []leanruntime.Field{
	{Name: "AccountID", Flag: "account-id", Type: "*string", Required: false},
	{Name: "SourceServerID", Flag: "source-server-id", Type: "*string", Required: true},
}

var fields_retry_data_replication = []leanruntime.Field{
	{Name: "AccountID", Flag: "account-id", Type: "*string", Required: false},
	{Name: "SourceServerID", Flag: "source-server-id", Type: "*string", Required: true},
}

var fields_start_cutover = []leanruntime.Field{
	{Name: "AccountID", Flag: "account-id", Type: "*string", Required: false},
	{Name: "SourceServerIDs", Flag: "source-server-ids", Type: "[]string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_start_export = []leanruntime.Field{
	{Name: "S3Bucket", Flag: "s3-bucket", Type: "*string", Required: true},
	{Name: "S3BucketOwner", Flag: "s3-bucket-owner", Type: "*string", Required: false},
	{Name: "S3Key", Flag: "s3-key", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_start_import = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "S3BucketSource", Flag: "s3-bucket-source", Type: "*types.S3BucketSource", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_start_replication = []leanruntime.Field{
	{Name: "AccountID", Flag: "account-id", Type: "*string", Required: false},
	{Name: "SourceServerID", Flag: "source-server-id", Type: "*string", Required: true},
}

var fields_start_test = []leanruntime.Field{
	{Name: "AccountID", Flag: "account-id", Type: "*string", Required: false},
	{Name: "SourceServerIDs", Flag: "source-server-ids", Type: "[]string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_stop_replication = []leanruntime.Field{
	{Name: "AccountID", Flag: "account-id", Type: "*string", Required: false},
	{Name: "SourceServerID", Flag: "source-server-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_terminate_target_instances = []leanruntime.Field{
	{Name: "AccountID", Flag: "account-id", Type: "*string", Required: false},
	{Name: "SourceServerIDs", Flag: "source-server-ids", Type: "[]string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_unarchive_application = []leanruntime.Field{
	{Name: "AccountID", Flag: "account-id", Type: "*string", Required: false},
	{Name: "ApplicationID", Flag: "application-id", Type: "*string", Required: true},
}

var fields_unarchive_wave = []leanruntime.Field{
	{Name: "AccountID", Flag: "account-id", Type: "*string", Required: false},
	{Name: "WaveID", Flag: "wave-id", Type: "*string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_application = []leanruntime.Field{
	{Name: "AccountID", Flag: "account-id", Type: "*string", Required: false},
	{Name: "ApplicationID", Flag: "application-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_connector = []leanruntime.Field{
	{Name: "ConnectorID", Flag: "connector-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "SsmCommandConfig", Flag: "ssm-command-config", Type: "*types.ConnectorSsmCommandConfig", Required: false},
}

var fields_update_launch_configuration = []leanruntime.Field{
	{Name: "AccountID", Flag: "account-id", Type: "*string", Required: false},
	{Name: "BootMode", Flag: "boot-mode", Type: "types.BootMode", Required: false},
	{Name: "CopyPrivateIp", Flag: "copy-private-ip", Type: "*bool", Required: false},
	{Name: "CopyTags", Flag: "copy-tags", Type: "*bool", Required: false},
	{Name: "EnableMapAutoTagging", Flag: "enable-map-auto-tagging", Type: "*bool", Required: false},
	{Name: "LaunchDisposition", Flag: "launch-disposition", Type: "types.LaunchDisposition", Required: false},
	{Name: "Licensing", Flag: "licensing", Type: "*types.Licensing", Required: false},
	{Name: "MapAutoTaggingMpeID", Flag: "map-auto-tagging-mpe-id", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "PostLaunchActions", Flag: "post-launch-actions", Type: "*types.PostLaunchActions", Required: false},
	{Name: "SourceServerID", Flag: "source-server-id", Type: "*string", Required: true},
	{Name: "TargetInstanceTypeRightSizingMethod", Flag: "target-instance-type-right-sizing-method", Type: "types.TargetInstanceTypeRightSizingMethod", Required: false},
}

var fields_update_launch_configuration_template = []leanruntime.Field{
	{Name: "AssociatePublicIpAddress", Flag: "associate-public-ip-address", Type: "*bool", Required: false},
	{Name: "BootMode", Flag: "boot-mode", Type: "types.BootMode", Required: false},
	{Name: "CopyPrivateIp", Flag: "copy-private-ip", Type: "*bool", Required: false},
	{Name: "CopyTags", Flag: "copy-tags", Type: "*bool", Required: false},
	{Name: "EnableMapAutoTagging", Flag: "enable-map-auto-tagging", Type: "*bool", Required: false},
	{Name: "EnableParametersEncryption", Flag: "enable-parameters-encryption", Type: "*bool", Required: false},
	{Name: "LargeVolumeConf", Flag: "large-volume-conf", Type: "*types.LaunchTemplateDiskConf", Required: false},
	{Name: "LaunchConfigurationTemplateID", Flag: "launch-configuration-template-id", Type: "*string", Required: true},
	{Name: "LaunchDisposition", Flag: "launch-disposition", Type: "types.LaunchDisposition", Required: false},
	{Name: "Licensing", Flag: "licensing", Type: "*types.Licensing", Required: false},
	{Name: "MapAutoTaggingMpeID", Flag: "map-auto-tagging-mpe-id", Type: "*string", Required: false},
	{Name: "ParametersEncryptionKey", Flag: "parameters-encryption-key", Type: "*string", Required: false},
	{Name: "PostLaunchActions", Flag: "post-launch-actions", Type: "*types.PostLaunchActions", Required: false},
	{Name: "SmallVolumeConf", Flag: "small-volume-conf", Type: "*types.LaunchTemplateDiskConf", Required: false},
	{Name: "SmallVolumeMaxSize", Flag: "small-volume-max-size", Type: "int64", Required: false},
	{Name: "TargetInstanceTypeRightSizingMethod", Flag: "target-instance-type-right-sizing-method", Type: "types.TargetInstanceTypeRightSizingMethod", Required: false},
}

var fields_update_replication_configuration = []leanruntime.Field{
	{Name: "AccountID", Flag: "account-id", Type: "*string", Required: false},
	{Name: "AssociateDefaultSecurityGroup", Flag: "associate-default-security-group", Type: "*bool", Required: false},
	{Name: "BandwidthThrottling", Flag: "bandwidth-throttling", Type: "int64", Required: false},
	{Name: "CreatePublicIP", Flag: "create-public-ip", Type: "*bool", Required: false},
	{Name: "DataPlaneRouting", Flag: "data-plane-routing", Type: "types.ReplicationConfigurationDataPlaneRouting", Required: false},
	{Name: "DefaultLargeStagingDiskType", Flag: "default-large-staging-disk-type", Type: "types.ReplicationConfigurationDefaultLargeStagingDiskType", Required: false},
	{Name: "EbsEncryption", Flag: "ebs-encryption", Type: "types.ReplicationConfigurationEbsEncryption", Required: false},
	{Name: "EbsEncryptionKeyArn", Flag: "ebs-encryption-key-arn", Type: "*string", Required: false},
	{Name: "InternetProtocol", Flag: "internet-protocol", Type: "types.InternetProtocol", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "ReplicatedDisks", Flag: "replicated-disks", Type: "[]types.ReplicationConfigurationReplicatedDisk", Required: false},
	{Name: "ReplicationServerInstanceType", Flag: "replication-server-instance-type", Type: "*string", Required: false},
	{Name: "ReplicationServersSecurityGroupsIDs", Flag: "replication-servers-security-groups-ids", Type: "[]string", Required: false},
	{Name: "SourceServerID", Flag: "source-server-id", Type: "*string", Required: true},
	{Name: "StagingAreaSubnetId", Flag: "staging-area-subnet-id", Type: "*string", Required: false},
	{Name: "StagingAreaTags", Flag: "staging-area-tags", Type: "map[string]string", Required: false},
	{Name: "UseDedicatedReplicationServer", Flag: "use-dedicated-replication-server", Type: "*bool", Required: false},
	{Name: "UseFipsEndpoint", Flag: "use-fips-endpoint", Type: "*bool", Required: false},
}

var fields_update_replication_configuration_template = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: false},
	{Name: "AssociateDefaultSecurityGroup", Flag: "associate-default-security-group", Type: "*bool", Required: false},
	{Name: "BandwidthThrottling", Flag: "bandwidth-throttling", Type: "int64", Required: false},
	{Name: "CreatePublicIP", Flag: "create-public-ip", Type: "*bool", Required: false},
	{Name: "DataPlaneRouting", Flag: "data-plane-routing", Type: "types.ReplicationConfigurationDataPlaneRouting", Required: false},
	{Name: "DefaultLargeStagingDiskType", Flag: "default-large-staging-disk-type", Type: "types.ReplicationConfigurationDefaultLargeStagingDiskType", Required: false},
	{Name: "EbsEncryption", Flag: "ebs-encryption", Type: "types.ReplicationConfigurationEbsEncryption", Required: false},
	{Name: "EbsEncryptionKeyArn", Flag: "ebs-encryption-key-arn", Type: "*string", Required: false},
	{Name: "InternetProtocol", Flag: "internet-protocol", Type: "types.InternetProtocol", Required: false},
	{Name: "ReplicationConfigurationTemplateID", Flag: "replication-configuration-template-id", Type: "*string", Required: true},
	{Name: "ReplicationServerInstanceType", Flag: "replication-server-instance-type", Type: "*string", Required: false},
	{Name: "ReplicationServersSecurityGroupsIDs", Flag: "replication-servers-security-groups-ids", Type: "[]string", Required: false},
	{Name: "StagingAreaSubnetId", Flag: "staging-area-subnet-id", Type: "*string", Required: false},
	{Name: "StagingAreaTags", Flag: "staging-area-tags", Type: "map[string]string", Required: false},
	{Name: "UseDedicatedReplicationServer", Flag: "use-dedicated-replication-server", Type: "*bool", Required: false},
	{Name: "UseFipsEndpoint", Flag: "use-fips-endpoint", Type: "*bool", Required: false},
}

var fields_update_source_server = []leanruntime.Field{
	{Name: "AccountID", Flag: "account-id", Type: "*string", Required: false},
	{Name: "ConnectorAction", Flag: "connector-action", Type: "*types.SourceServerConnectorAction", Required: false},
	{Name: "SourceServerID", Flag: "source-server-id", Type: "*string", Required: true},
}

var fields_update_source_server_replication_type = []leanruntime.Field{
	{Name: "AccountID", Flag: "account-id", Type: "*string", Required: false},
	{Name: "ReplicationType", Flag: "replication-type", Type: "types.ReplicationType", Required: true},
	{Name: "SourceServerID", Flag: "source-server-id", Type: "*string", Required: true},
}

var fields_update_wave = []leanruntime.Field{
	{Name: "AccountID", Flag: "account-id", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "WaveID", Flag: "wave-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"archive-application": {
			Name:   "archive-application",
			Fields: fields_archive_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ArchiveApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_archive_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ArchiveApplication(ctx, input)
			},
		},
		"archive-wave": {
			Name:   "archive-wave",
			Fields: fields_archive_wave,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ArchiveWaveInput{}
				if _, err := leanruntime.ApplyInput(input, fields_archive_wave, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ArchiveWave(ctx, input)
			},
		},
		"associate-applications": {
			Name:   "associate-applications",
			Fields: fields_associate_applications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateApplicationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_applications, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateApplications(ctx, input)
			},
		},
		"associate-source-servers": {
			Name:   "associate-source-servers",
			Fields: fields_associate_source_servers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateSourceServersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_source_servers, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateSourceServers(ctx, input)
			},
		},
		"change-server-life-cycle-state": {
			Name:   "change-server-life-cycle-state",
			Fields: fields_change_server_life_cycle_state,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ChangeServerLifeCycleStateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_change_server_life_cycle_state, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ChangeServerLifeCycleState(ctx, input)
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
		"create-connector": {
			Name:   "create-connector",
			Fields: fields_create_connector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateConnectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_connector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateConnector(ctx, input)
			},
		},
		"create-launch-configuration-template": {
			Name:   "create-launch-configuration-template",
			Fields: fields_create_launch_configuration_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLaunchConfigurationTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_launch_configuration_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLaunchConfigurationTemplate(ctx, input)
			},
		},
		"create-replication-configuration-template": {
			Name:   "create-replication-configuration-template",
			Fields: fields_create_replication_configuration_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateReplicationConfigurationTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_replication_configuration_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateReplicationConfigurationTemplate(ctx, input)
			},
		},
		"create-wave": {
			Name:   "create-wave",
			Fields: fields_create_wave,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWaveInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_wave, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWave(ctx, input)
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
		"delete-connector": {
			Name:   "delete-connector",
			Fields: fields_delete_connector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConnectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_connector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConnector(ctx, input)
			},
		},
		"delete-job": {
			Name:   "delete-job",
			Fields: fields_delete_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteJob(ctx, input)
			},
		},
		"delete-launch-configuration-template": {
			Name:   "delete-launch-configuration-template",
			Fields: fields_delete_launch_configuration_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLaunchConfigurationTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_launch_configuration_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLaunchConfigurationTemplate(ctx, input)
			},
		},
		"delete-replication-configuration-template": {
			Name:   "delete-replication-configuration-template",
			Fields: fields_delete_replication_configuration_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteReplicationConfigurationTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_replication_configuration_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteReplicationConfigurationTemplate(ctx, input)
			},
		},
		"delete-source-server": {
			Name:   "delete-source-server",
			Fields: fields_delete_source_server,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSourceServerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_source_server, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSourceServer(ctx, input)
			},
		},
		"delete-vcenter-client": {
			Name:   "delete-vcenter-client",
			Fields: fields_delete_vcenter_client,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVcenterClientInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_vcenter_client, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVcenterClient(ctx, input)
			},
		},
		"delete-wave": {
			Name:   "delete-wave",
			Fields: fields_delete_wave,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWaveInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_wave, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWave(ctx, input)
			},
		},
		"describe-job-log-items": {
			Name:   "describe-job-log-items",
			Fields: fields_describe_job_log_items,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeJobLogItemsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_job_log_items, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeJobLogItems(ctx, input)
				}
				var results []*svc.DescribeJobLogItemsOutput
				p := svc.NewDescribeJobLogItemsPaginator(client, input)
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
		"describe-jobs": {
			Name:   "describe-jobs",
			Fields: fields_describe_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeJobs(ctx, input)
				}
				var results []*svc.DescribeJobsOutput
				p := svc.NewDescribeJobsPaginator(client, input)
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
		"describe-launch-configuration-templates": {
			Name:   "describe-launch-configuration-templates",
			Fields: fields_describe_launch_configuration_templates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLaunchConfigurationTemplatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_launch_configuration_templates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeLaunchConfigurationTemplates(ctx, input)
				}
				var results []*svc.DescribeLaunchConfigurationTemplatesOutput
				p := svc.NewDescribeLaunchConfigurationTemplatesPaginator(client, input)
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
		"describe-replication-configuration-templates": {
			Name:   "describe-replication-configuration-templates",
			Fields: fields_describe_replication_configuration_templates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeReplicationConfigurationTemplatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_replication_configuration_templates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeReplicationConfigurationTemplates(ctx, input)
				}
				var results []*svc.DescribeReplicationConfigurationTemplatesOutput
				p := svc.NewDescribeReplicationConfigurationTemplatesPaginator(client, input)
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
		"describe-source-servers": {
			Name:   "describe-source-servers",
			Fields: fields_describe_source_servers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSourceServersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_source_servers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeSourceServers(ctx, input)
				}
				var results []*svc.DescribeSourceServersOutput
				p := svc.NewDescribeSourceServersPaginator(client, input)
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
		"describe-vcenter-clients": {
			Name:   "describe-vcenter-clients",
			Fields: fields_describe_vcenter_clients,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVcenterClientsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_vcenter_clients, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeVcenterClients(ctx, input)
				}
				var results []*svc.DescribeVcenterClientsOutput
				p := svc.NewDescribeVcenterClientsPaginator(client, input)
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
		"disassociate-applications": {
			Name:   "disassociate-applications",
			Fields: fields_disassociate_applications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateApplicationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_applications, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateApplications(ctx, input)
			},
		},
		"disassociate-source-servers": {
			Name:   "disassociate-source-servers",
			Fields: fields_disassociate_source_servers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateSourceServersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_source_servers, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateSourceServers(ctx, input)
			},
		},
		"disconnect-from-service": {
			Name:   "disconnect-from-service",
			Fields: fields_disconnect_from_service,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisconnectFromServiceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disconnect_from_service, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisconnectFromService(ctx, input)
			},
		},
		"finalize-cutover": {
			Name:   "finalize-cutover",
			Fields: fields_finalize_cutover,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.FinalizeCutoverInput{}
				if _, err := leanruntime.ApplyInput(input, fields_finalize_cutover, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.FinalizeCutover(ctx, input)
			},
		},
		"get-launch-configuration": {
			Name:   "get-launch-configuration",
			Fields: fields_get_launch_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLaunchConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_launch_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLaunchConfiguration(ctx, input)
			},
		},
		"get-replication-configuration": {
			Name:   "get-replication-configuration",
			Fields: fields_get_replication_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetReplicationConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_replication_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetReplicationConfiguration(ctx, input)
			},
		},
		"initialize-service": {
			Name:   "initialize-service",
			Fields: fields_initialize_service,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.InitializeServiceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_initialize_service, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.InitializeService(ctx, input)
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
		"list-connectors": {
			Name:   "list-connectors",
			Fields: fields_list_connectors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConnectorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_connectors, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListConnectors(ctx, input)
				}
				var results []*svc.ListConnectorsOutput
				p := svc.NewListConnectorsPaginator(client, input)
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
		"list-export-errors": {
			Name:   "list-export-errors",
			Fields: fields_list_export_errors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListExportErrorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_export_errors, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListExportErrors(ctx, input)
				}
				var results []*svc.ListExportErrorsOutput
				p := svc.NewListExportErrorsPaginator(client, input)
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
		"list-exports": {
			Name:   "list-exports",
			Fields: fields_list_exports,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListExportsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_exports, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListExports(ctx, input)
				}
				var results []*svc.ListExportsOutput
				p := svc.NewListExportsPaginator(client, input)
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
		"list-import-errors": {
			Name:   "list-import-errors",
			Fields: fields_list_import_errors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListImportErrorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_import_errors, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListImportErrors(ctx, input)
				}
				var results []*svc.ListImportErrorsOutput
				p := svc.NewListImportErrorsPaginator(client, input)
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
		"list-imports": {
			Name:   "list-imports",
			Fields: fields_list_imports,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListImportsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_imports, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListImports(ctx, input)
				}
				var results []*svc.ListImportsOutput
				p := svc.NewListImportsPaginator(client, input)
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
		"list-managed-accounts": {
			Name:   "list-managed-accounts",
			Fields: fields_list_managed_accounts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListManagedAccountsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_managed_accounts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListManagedAccounts(ctx, input)
				}
				var results []*svc.ListManagedAccountsOutput
				p := svc.NewListManagedAccountsPaginator(client, input)
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
		"list-source-server-actions": {
			Name:   "list-source-server-actions",
			Fields: fields_list_source_server_actions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSourceServerActionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_source_server_actions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSourceServerActions(ctx, input)
				}
				var results []*svc.ListSourceServerActionsOutput
				p := svc.NewListSourceServerActionsPaginator(client, input)
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
		"list-template-actions": {
			Name:   "list-template-actions",
			Fields: fields_list_template_actions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTemplateActionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_template_actions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTemplateActions(ctx, input)
				}
				var results []*svc.ListTemplateActionsOutput
				p := svc.NewListTemplateActionsPaginator(client, input)
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
		"list-waves": {
			Name:   "list-waves",
			Fields: fields_list_waves,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWavesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_waves, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWaves(ctx, input)
				}
				var results []*svc.ListWavesOutput
				p := svc.NewListWavesPaginator(client, input)
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
		"mark-as-archived": {
			Name:   "mark-as-archived",
			Fields: fields_mark_as_archived,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.MarkAsArchivedInput{}
				if _, err := leanruntime.ApplyInput(input, fields_mark_as_archived, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.MarkAsArchived(ctx, input)
			},
		},
		"pause-replication": {
			Name:   "pause-replication",
			Fields: fields_pause_replication,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PauseReplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_pause_replication, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PauseReplication(ctx, input)
			},
		},
		"put-source-server-action": {
			Name:   "put-source-server-action",
			Fields: fields_put_source_server_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutSourceServerActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_source_server_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutSourceServerAction(ctx, input)
			},
		},
		"put-template-action": {
			Name:   "put-template-action",
			Fields: fields_put_template_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutTemplateActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_template_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutTemplateAction(ctx, input)
			},
		},
		"remove-source-server-action": {
			Name:   "remove-source-server-action",
			Fields: fields_remove_source_server_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveSourceServerActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_source_server_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveSourceServerAction(ctx, input)
			},
		},
		"remove-template-action": {
			Name:   "remove-template-action",
			Fields: fields_remove_template_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveTemplateActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_template_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveTemplateAction(ctx, input)
			},
		},
		"resume-replication": {
			Name:   "resume-replication",
			Fields: fields_resume_replication,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResumeReplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_resume_replication, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResumeReplication(ctx, input)
			},
		},
		"retry-data-replication": {
			Name:   "retry-data-replication",
			Fields: fields_retry_data_replication,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RetryDataReplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_retry_data_replication, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RetryDataReplication(ctx, input)
			},
		},
		"start-cutover": {
			Name:   "start-cutover",
			Fields: fields_start_cutover,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartCutoverInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_cutover, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartCutover(ctx, input)
			},
		},
		"start-export": {
			Name:   "start-export",
			Fields: fields_start_export,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartExportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_export, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartExport(ctx, input)
			},
		},
		"start-import": {
			Name:   "start-import",
			Fields: fields_start_import,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartImportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_import, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartImport(ctx, input)
			},
		},
		"start-replication": {
			Name:   "start-replication",
			Fields: fields_start_replication,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartReplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_replication, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartReplication(ctx, input)
			},
		},
		"start-test": {
			Name:   "start-test",
			Fields: fields_start_test,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartTestInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_test, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartTest(ctx, input)
			},
		},
		"stop-replication": {
			Name:   "stop-replication",
			Fields: fields_stop_replication,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopReplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_replication, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopReplication(ctx, input)
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
		"terminate-target-instances": {
			Name:   "terminate-target-instances",
			Fields: fields_terminate_target_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TerminateTargetInstancesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_terminate_target_instances, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TerminateTargetInstances(ctx, input)
			},
		},
		"unarchive-application": {
			Name:   "unarchive-application",
			Fields: fields_unarchive_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UnarchiveApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_unarchive_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UnarchiveApplication(ctx, input)
			},
		},
		"unarchive-wave": {
			Name:   "unarchive-wave",
			Fields: fields_unarchive_wave,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UnarchiveWaveInput{}
				if _, err := leanruntime.ApplyInput(input, fields_unarchive_wave, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UnarchiveWave(ctx, input)
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
		"update-connector": {
			Name:   "update-connector",
			Fields: fields_update_connector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateConnectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_connector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateConnector(ctx, input)
			},
		},
		"update-launch-configuration": {
			Name:   "update-launch-configuration",
			Fields: fields_update_launch_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateLaunchConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_launch_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateLaunchConfiguration(ctx, input)
			},
		},
		"update-launch-configuration-template": {
			Name:   "update-launch-configuration-template",
			Fields: fields_update_launch_configuration_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateLaunchConfigurationTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_launch_configuration_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateLaunchConfigurationTemplate(ctx, input)
			},
		},
		"update-replication-configuration": {
			Name:   "update-replication-configuration",
			Fields: fields_update_replication_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateReplicationConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_replication_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateReplicationConfiguration(ctx, input)
			},
		},
		"update-replication-configuration-template": {
			Name:   "update-replication-configuration-template",
			Fields: fields_update_replication_configuration_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateReplicationConfigurationTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_replication_configuration_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateReplicationConfigurationTemplate(ctx, input)
			},
		},
		"update-source-server": {
			Name:   "update-source-server",
			Fields: fields_update_source_server,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSourceServerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_source_server, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSourceServer(ctx, input)
			},
		},
		"update-source-server-replication-type": {
			Name:   "update-source-server-replication-type",
			Fields: fields_update_source_server_replication_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSourceServerReplicationTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_source_server_replication_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSourceServerReplicationType(ctx, input)
			},
		},
		"update-wave": {
			Name:   "update-wave",
			Fields: fields_update_wave,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateWaveInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_wave, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateWave(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("mgn", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
