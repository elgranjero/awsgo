package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/drs"
)

var fields_associate_source_network_stack = []leanruntime.Field{
	{Name: "CfnStackName", Flag: "cfn-stack-name", Type: "*string", Required: true},
	{Name: "SourceNetworkID", Flag: "source-network-id", Type: "*string", Required: true},
}

var fields_create_extended_source_server = []leanruntime.Field{
	{Name: "SourceServerArn", Flag: "source-server-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_launch_configuration_template = []leanruntime.Field{
	{Name: "CopyPrivateIp", Flag: "copy-private-ip", Type: "*bool", Required: false},
	{Name: "CopyTags", Flag: "copy-tags", Type: "*bool", Required: false},
	{Name: "ExportBucketArn", Flag: "export-bucket-arn", Type: "*string", Required: false},
	{Name: "LaunchDisposition", Flag: "launch-disposition", Type: "types.LaunchDisposition", Required: false},
	{Name: "LaunchIntoSourceInstance", Flag: "launch-into-source-instance", Type: "*bool", Required: false},
	{Name: "Licensing", Flag: "licensing", Type: "*types.Licensing", Required: false},
	{Name: "PostLaunchEnabled", Flag: "post-launch-enabled", Type: "*bool", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TargetInstanceTypeRightSizingMethod", Flag: "target-instance-type-right-sizing-method", Type: "types.TargetInstanceTypeRightSizingMethod", Required: false},
}

var fields_create_replication_configuration_template = []leanruntime.Field{
	{Name: "AssociateDefaultSecurityGroup", Flag: "associate-default-security-group", Type: "*bool", Required: true},
	{Name: "AutoReplicateNewDisks", Flag: "auto-replicate-new-disks", Type: "*bool", Required: false},
	{Name: "BandwidthThrottling", Flag: "bandwidth-throttling", Type: "int64", Required: true},
	{Name: "CreatePublicIP", Flag: "create-public-ip", Type: "*bool", Required: true},
	{Name: "DataPlaneRouting", Flag: "data-plane-routing", Type: "types.ReplicationConfigurationDataPlaneRouting", Required: true},
	{Name: "DefaultLargeStagingDiskType", Flag: "default-large-staging-disk-type", Type: "types.ReplicationConfigurationDefaultLargeStagingDiskType", Required: true},
	{Name: "EbsEncryption", Flag: "ebs-encryption", Type: "types.ReplicationConfigurationEbsEncryption", Required: true},
	{Name: "EbsEncryptionKeyArn", Flag: "ebs-encryption-key-arn", Type: "*string", Required: false},
	{Name: "PitPolicy", Flag: "pit-policy", Type: "[]types.PITPolicyRule", Required: true},
	{Name: "ReplicationServerInstanceType", Flag: "replication-server-instance-type", Type: "*string", Required: true},
	{Name: "ReplicationServersSecurityGroupsIDs", Flag: "replication-servers-security-groups-ids", Type: "[]string", Required: true},
	{Name: "StagingAreaSubnetId", Flag: "staging-area-subnet-id", Type: "*string", Required: true},
	{Name: "StagingAreaTags", Flag: "staging-area-tags", Type: "map[string]string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "UseDedicatedReplicationServer", Flag: "use-dedicated-replication-server", Type: "*bool", Required: true},
}

var fields_create_source_network = []leanruntime.Field{
	{Name: "OriginAccountID", Flag: "origin-account-id", Type: "*string", Required: true},
	{Name: "OriginRegion", Flag: "origin-region", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "VpcID", Flag: "vpc-id", Type: "*string", Required: true},
}

var fields_delete_job = []leanruntime.Field{
	{Name: "JobID", Flag: "job-id", Type: "*string", Required: true},
}

var fields_delete_launch_action = []leanruntime.Field{
	{Name: "ActionId", Flag: "action-id", Type: "*string", Required: true},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
}

var fields_delete_launch_configuration_template = []leanruntime.Field{
	{Name: "LaunchConfigurationTemplateID", Flag: "launch-configuration-template-id", Type: "*string", Required: true},
}

var fields_delete_recovery_instance = []leanruntime.Field{
	{Name: "RecoveryInstanceID", Flag: "recovery-instance-id", Type: "*string", Required: true},
}

var fields_delete_replication_configuration_template = []leanruntime.Field{
	{Name: "ReplicationConfigurationTemplateID", Flag: "replication-configuration-template-id", Type: "*string", Required: true},
}

var fields_delete_source_network = []leanruntime.Field{
	{Name: "SourceNetworkID", Flag: "source-network-id", Type: "*string", Required: true},
}

var fields_delete_source_server = []leanruntime.Field{
	{Name: "SourceServerID", Flag: "source-server-id", Type: "*string", Required: true},
}

var fields_describe_job_log_items = []leanruntime.Field{
	{Name: "JobID", Flag: "job-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_jobs = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "*types.DescribeJobsRequestFilters", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_launch_configuration_templates = []leanruntime.Field{
	{Name: "LaunchConfigurationTemplateIDs", Flag: "launch-configuration-template-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_recovery_instances = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "*types.DescribeRecoveryInstancesRequestFilters", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_recovery_snapshots = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "*types.DescribeRecoverySnapshotsRequestFilters", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Order", Flag: "order", Type: "types.RecoverySnapshotsOrder", Required: false},
	{Name: "SourceServerID", Flag: "source-server-id", Type: "*string", Required: true},
}

var fields_describe_replication_configuration_templates = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ReplicationConfigurationTemplateIDs", Flag: "replication-configuration-template-ids", Type: "[]string", Required: false},
}

var fields_describe_source_networks = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "*types.DescribeSourceNetworksRequestFilters", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_source_servers = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "*types.DescribeSourceServersRequestFilters", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_disconnect_recovery_instance = []leanruntime.Field{
	{Name: "RecoveryInstanceID", Flag: "recovery-instance-id", Type: "*string", Required: true},
}

var fields_disconnect_source_server = []leanruntime.Field{
	{Name: "SourceServerID", Flag: "source-server-id", Type: "*string", Required: true},
}

var fields_export_source_network_cfn_template = []leanruntime.Field{
	{Name: "SourceNetworkID", Flag: "source-network-id", Type: "*string", Required: true},
}

var fields_get_failback_replication_configuration = []leanruntime.Field{
	{Name: "RecoveryInstanceID", Flag: "recovery-instance-id", Type: "*string", Required: true},
}

var fields_get_launch_configuration = []leanruntime.Field{
	{Name: "SourceServerID", Flag: "source-server-id", Type: "*string", Required: true},
}

var fields_get_replication_configuration = []leanruntime.Field{
	{Name: "SourceServerID", Flag: "source-server-id", Type: "*string", Required: true},
}

var fields_initialize_service = []leanruntime.Field{}

var fields_list_extensible_source_servers = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StagingAccountID", Flag: "staging-account-id", Type: "*string", Required: true},
}

var fields_list_launch_actions = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "*types.LaunchActionsRequestFilters", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
}

var fields_list_staging_accounts = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_put_launch_action = []leanruntime.Field{
	{Name: "ActionCode", Flag: "action-code", Type: "*string", Required: true},
	{Name: "ActionId", Flag: "action-id", Type: "*string", Required: true},
	{Name: "ActionVersion", Flag: "action-version", Type: "*string", Required: true},
	{Name: "Active", Flag: "active", Type: "*bool", Required: true},
	{Name: "Category", Flag: "category", Type: "types.LaunchActionCategory", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Optional", Flag: "optional", Type: "*bool", Required: true},
	{Name: "Order", Flag: "order", Type: "*int32", Required: true},
	{Name: "Parameters", Flag: "parameters", Type: "map[string]types.LaunchActionParameter", Required: false},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
}

var fields_retry_data_replication = []leanruntime.Field{
	{Name: "SourceServerID", Flag: "source-server-id", Type: "*string", Required: true},
}

var fields_reverse_replication = []leanruntime.Field{
	{Name: "RecoveryInstanceID", Flag: "recovery-instance-id", Type: "*string", Required: true},
}

var fields_start_failback_launch = []leanruntime.Field{
	{Name: "RecoveryInstanceIDs", Flag: "recovery-instance-ids", Type: "[]string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_start_recovery = []leanruntime.Field{
	{Name: "IsDrill", Flag: "is-drill", Type: "*bool", Required: false},
	{Name: "SourceServers", Flag: "source-servers", Type: "[]types.StartRecoveryRequestSourceServer", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_start_replication = []leanruntime.Field{
	{Name: "SourceServerID", Flag: "source-server-id", Type: "*string", Required: true},
}

var fields_start_source_network_recovery = []leanruntime.Field{
	{Name: "DeployAsNew", Flag: "deploy-as-new", Type: "*bool", Required: false},
	{Name: "SourceNetworks", Flag: "source-networks", Type: "[]types.StartSourceNetworkRecoveryRequestNetworkEntry", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_start_source_network_replication = []leanruntime.Field{
	{Name: "SourceNetworkID", Flag: "source-network-id", Type: "*string", Required: true},
}

var fields_stop_failback = []leanruntime.Field{
	{Name: "RecoveryInstanceID", Flag: "recovery-instance-id", Type: "*string", Required: true},
}

var fields_stop_replication = []leanruntime.Field{
	{Name: "SourceServerID", Flag: "source-server-id", Type: "*string", Required: true},
}

var fields_stop_source_network_replication = []leanruntime.Field{
	{Name: "SourceNetworkID", Flag: "source-network-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_terminate_recovery_instances = []leanruntime.Field{
	{Name: "RecoveryInstanceIDs", Flag: "recovery-instance-ids", Type: "[]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_failback_replication_configuration = []leanruntime.Field{
	{Name: "BandwidthThrottling", Flag: "bandwidth-throttling", Type: "int64", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "RecoveryInstanceID", Flag: "recovery-instance-id", Type: "*string", Required: true},
	{Name: "UsePrivateIP", Flag: "use-private-ip", Type: "*bool", Required: false},
}

var fields_update_launch_configuration = []leanruntime.Field{
	{Name: "CopyPrivateIp", Flag: "copy-private-ip", Type: "*bool", Required: false},
	{Name: "CopyTags", Flag: "copy-tags", Type: "*bool", Required: false},
	{Name: "LaunchDisposition", Flag: "launch-disposition", Type: "types.LaunchDisposition", Required: false},
	{Name: "LaunchIntoInstanceProperties", Flag: "launch-into-instance-properties", Type: "*types.LaunchIntoInstanceProperties", Required: false},
	{Name: "Licensing", Flag: "licensing", Type: "*types.Licensing", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "PostLaunchEnabled", Flag: "post-launch-enabled", Type: "*bool", Required: false},
	{Name: "SourceServerID", Flag: "source-server-id", Type: "*string", Required: true},
	{Name: "TargetInstanceTypeRightSizingMethod", Flag: "target-instance-type-right-sizing-method", Type: "types.TargetInstanceTypeRightSizingMethod", Required: false},
}

var fields_update_launch_configuration_template = []leanruntime.Field{
	{Name: "CopyPrivateIp", Flag: "copy-private-ip", Type: "*bool", Required: false},
	{Name: "CopyTags", Flag: "copy-tags", Type: "*bool", Required: false},
	{Name: "ExportBucketArn", Flag: "export-bucket-arn", Type: "*string", Required: false},
	{Name: "LaunchConfigurationTemplateID", Flag: "launch-configuration-template-id", Type: "*string", Required: true},
	{Name: "LaunchDisposition", Flag: "launch-disposition", Type: "types.LaunchDisposition", Required: false},
	{Name: "LaunchIntoSourceInstance", Flag: "launch-into-source-instance", Type: "*bool", Required: false},
	{Name: "Licensing", Flag: "licensing", Type: "*types.Licensing", Required: false},
	{Name: "PostLaunchEnabled", Flag: "post-launch-enabled", Type: "*bool", Required: false},
	{Name: "TargetInstanceTypeRightSizingMethod", Flag: "target-instance-type-right-sizing-method", Type: "types.TargetInstanceTypeRightSizingMethod", Required: false},
}

var fields_update_replication_configuration = []leanruntime.Field{
	{Name: "AssociateDefaultSecurityGroup", Flag: "associate-default-security-group", Type: "*bool", Required: false},
	{Name: "AutoReplicateNewDisks", Flag: "auto-replicate-new-disks", Type: "*bool", Required: false},
	{Name: "BandwidthThrottling", Flag: "bandwidth-throttling", Type: "int64", Required: false},
	{Name: "CreatePublicIP", Flag: "create-public-ip", Type: "*bool", Required: false},
	{Name: "DataPlaneRouting", Flag: "data-plane-routing", Type: "types.ReplicationConfigurationDataPlaneRouting", Required: false},
	{Name: "DefaultLargeStagingDiskType", Flag: "default-large-staging-disk-type", Type: "types.ReplicationConfigurationDefaultLargeStagingDiskType", Required: false},
	{Name: "EbsEncryption", Flag: "ebs-encryption", Type: "types.ReplicationConfigurationEbsEncryption", Required: false},
	{Name: "EbsEncryptionKeyArn", Flag: "ebs-encryption-key-arn", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "PitPolicy", Flag: "pit-policy", Type: "[]types.PITPolicyRule", Required: false},
	{Name: "ReplicatedDisks", Flag: "replicated-disks", Type: "[]types.ReplicationConfigurationReplicatedDisk", Required: false},
	{Name: "ReplicationServerInstanceType", Flag: "replication-server-instance-type", Type: "*string", Required: false},
	{Name: "ReplicationServersSecurityGroupsIDs", Flag: "replication-servers-security-groups-ids", Type: "[]string", Required: false},
	{Name: "SourceServerID", Flag: "source-server-id", Type: "*string", Required: true},
	{Name: "StagingAreaSubnetId", Flag: "staging-area-subnet-id", Type: "*string", Required: false},
	{Name: "StagingAreaTags", Flag: "staging-area-tags", Type: "map[string]string", Required: false},
	{Name: "UseDedicatedReplicationServer", Flag: "use-dedicated-replication-server", Type: "*bool", Required: false},
}

var fields_update_replication_configuration_template = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: false},
	{Name: "AssociateDefaultSecurityGroup", Flag: "associate-default-security-group", Type: "*bool", Required: false},
	{Name: "AutoReplicateNewDisks", Flag: "auto-replicate-new-disks", Type: "*bool", Required: false},
	{Name: "BandwidthThrottling", Flag: "bandwidth-throttling", Type: "int64", Required: false},
	{Name: "CreatePublicIP", Flag: "create-public-ip", Type: "*bool", Required: false},
	{Name: "DataPlaneRouting", Flag: "data-plane-routing", Type: "types.ReplicationConfigurationDataPlaneRouting", Required: false},
	{Name: "DefaultLargeStagingDiskType", Flag: "default-large-staging-disk-type", Type: "types.ReplicationConfigurationDefaultLargeStagingDiskType", Required: false},
	{Name: "EbsEncryption", Flag: "ebs-encryption", Type: "types.ReplicationConfigurationEbsEncryption", Required: false},
	{Name: "EbsEncryptionKeyArn", Flag: "ebs-encryption-key-arn", Type: "*string", Required: false},
	{Name: "PitPolicy", Flag: "pit-policy", Type: "[]types.PITPolicyRule", Required: false},
	{Name: "ReplicationConfigurationTemplateID", Flag: "replication-configuration-template-id", Type: "*string", Required: true},
	{Name: "ReplicationServerInstanceType", Flag: "replication-server-instance-type", Type: "*string", Required: false},
	{Name: "ReplicationServersSecurityGroupsIDs", Flag: "replication-servers-security-groups-ids", Type: "[]string", Required: false},
	{Name: "StagingAreaSubnetId", Flag: "staging-area-subnet-id", Type: "*string", Required: false},
	{Name: "StagingAreaTags", Flag: "staging-area-tags", Type: "map[string]string", Required: false},
	{Name: "UseDedicatedReplicationServer", Flag: "use-dedicated-replication-server", Type: "*bool", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-source-network-stack": {
			Name:   "associate-source-network-stack",
			Fields: fields_associate_source_network_stack,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateSourceNetworkStackInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_source_network_stack, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateSourceNetworkStack(ctx, input)
			},
		},
		"create-extended-source-server": {
			Name:   "create-extended-source-server",
			Fields: fields_create_extended_source_server,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateExtendedSourceServerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_extended_source_server, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateExtendedSourceServer(ctx, input)
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
		"create-source-network": {
			Name:   "create-source-network",
			Fields: fields_create_source_network,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSourceNetworkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_source_network, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSourceNetwork(ctx, input)
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
		"delete-launch-action": {
			Name:   "delete-launch-action",
			Fields: fields_delete_launch_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLaunchActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_launch_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLaunchAction(ctx, input)
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
		"delete-recovery-instance": {
			Name:   "delete-recovery-instance",
			Fields: fields_delete_recovery_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRecoveryInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_recovery_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRecoveryInstance(ctx, input)
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
		"delete-source-network": {
			Name:   "delete-source-network",
			Fields: fields_delete_source_network,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSourceNetworkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_source_network, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSourceNetwork(ctx, input)
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
		"describe-recovery-instances": {
			Name:   "describe-recovery-instances",
			Fields: fields_describe_recovery_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRecoveryInstancesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_recovery_instances, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeRecoveryInstances(ctx, input)
				}
				var results []*svc.DescribeRecoveryInstancesOutput
				p := svc.NewDescribeRecoveryInstancesPaginator(client, input)
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
		"describe-recovery-snapshots": {
			Name:   "describe-recovery-snapshots",
			Fields: fields_describe_recovery_snapshots,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRecoverySnapshotsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_recovery_snapshots, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeRecoverySnapshots(ctx, input)
				}
				var results []*svc.DescribeRecoverySnapshotsOutput
				p := svc.NewDescribeRecoverySnapshotsPaginator(client, input)
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
		"describe-source-networks": {
			Name:   "describe-source-networks",
			Fields: fields_describe_source_networks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSourceNetworksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_source_networks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeSourceNetworks(ctx, input)
				}
				var results []*svc.DescribeSourceNetworksOutput
				p := svc.NewDescribeSourceNetworksPaginator(client, input)
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
		"disconnect-recovery-instance": {
			Name:   "disconnect-recovery-instance",
			Fields: fields_disconnect_recovery_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisconnectRecoveryInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disconnect_recovery_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisconnectRecoveryInstance(ctx, input)
			},
		},
		"disconnect-source-server": {
			Name:   "disconnect-source-server",
			Fields: fields_disconnect_source_server,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisconnectSourceServerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disconnect_source_server, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisconnectSourceServer(ctx, input)
			},
		},
		"export-source-network-cfn-template": {
			Name:   "export-source-network-cfn-template",
			Fields: fields_export_source_network_cfn_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExportSourceNetworkCfnTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_export_source_network_cfn_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExportSourceNetworkCfnTemplate(ctx, input)
			},
		},
		"get-failback-replication-configuration": {
			Name:   "get-failback-replication-configuration",
			Fields: fields_get_failback_replication_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFailbackReplicationConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_failback_replication_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFailbackReplicationConfiguration(ctx, input)
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
		"list-extensible-source-servers": {
			Name:   "list-extensible-source-servers",
			Fields: fields_list_extensible_source_servers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListExtensibleSourceServersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_extensible_source_servers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListExtensibleSourceServers(ctx, input)
				}
				var results []*svc.ListExtensibleSourceServersOutput
				p := svc.NewListExtensibleSourceServersPaginator(client, input)
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
		"list-launch-actions": {
			Name:   "list-launch-actions",
			Fields: fields_list_launch_actions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLaunchActionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_launch_actions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLaunchActions(ctx, input)
				}
				var results []*svc.ListLaunchActionsOutput
				p := svc.NewListLaunchActionsPaginator(client, input)
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
		"list-staging-accounts": {
			Name:   "list-staging-accounts",
			Fields: fields_list_staging_accounts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListStagingAccountsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_staging_accounts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListStagingAccounts(ctx, input)
				}
				var results []*svc.ListStagingAccountsOutput
				p := svc.NewListStagingAccountsPaginator(client, input)
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
		"put-launch-action": {
			Name:   "put-launch-action",
			Fields: fields_put_launch_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutLaunchActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_launch_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutLaunchAction(ctx, input)
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
		"reverse-replication": {
			Name:   "reverse-replication",
			Fields: fields_reverse_replication,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ReverseReplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reverse_replication, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ReverseReplication(ctx, input)
			},
		},
		"start-failback-launch": {
			Name:   "start-failback-launch",
			Fields: fields_start_failback_launch,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartFailbackLaunchInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_failback_launch, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartFailbackLaunch(ctx, input)
			},
		},
		"start-recovery": {
			Name:   "start-recovery",
			Fields: fields_start_recovery,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartRecoveryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_recovery, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartRecovery(ctx, input)
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
		"start-source-network-recovery": {
			Name:   "start-source-network-recovery",
			Fields: fields_start_source_network_recovery,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartSourceNetworkRecoveryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_source_network_recovery, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartSourceNetworkRecovery(ctx, input)
			},
		},
		"start-source-network-replication": {
			Name:   "start-source-network-replication",
			Fields: fields_start_source_network_replication,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartSourceNetworkReplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_source_network_replication, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartSourceNetworkReplication(ctx, input)
			},
		},
		"stop-failback": {
			Name:   "stop-failback",
			Fields: fields_stop_failback,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopFailbackInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_failback, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopFailback(ctx, input)
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
		"stop-source-network-replication": {
			Name:   "stop-source-network-replication",
			Fields: fields_stop_source_network_replication,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopSourceNetworkReplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_source_network_replication, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopSourceNetworkReplication(ctx, input)
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
		"terminate-recovery-instances": {
			Name:   "terminate-recovery-instances",
			Fields: fields_terminate_recovery_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TerminateRecoveryInstancesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_terminate_recovery_instances, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TerminateRecoveryInstances(ctx, input)
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
		"update-failback-replication-configuration": {
			Name:   "update-failback-replication-configuration",
			Fields: fields_update_failback_replication_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFailbackReplicationConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_failback_replication_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFailbackReplicationConfiguration(ctx, input)
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
	}
	if err := leanruntime.Execute("drs", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
