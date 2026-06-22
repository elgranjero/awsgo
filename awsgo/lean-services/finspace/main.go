package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/finspace"
)

var fields_create_environment = []leanruntime.Field{
	{Name: "DataBundles", Flag: "data-bundles", Type: "[]string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "FederationMode", Flag: "federation-mode", Type: "types.FederationMode", Required: false},
	{Name: "FederationParameters", Flag: "federation-parameters", Type: "*types.FederationParameters", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "SuperuserParameters", Flag: "superuser-parameters", Type: "*types.SuperuserParameters", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_kx_changeset = []leanruntime.Field{
	{Name: "ChangeRequests", Flag: "change-requests", Type: "[]types.ChangeRequest", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
}

var fields_create_kx_cluster = []leanruntime.Field{
	{Name: "AutoScalingConfiguration", Flag: "auto-scaling-configuration", Type: "*types.AutoScalingConfiguration", Required: false},
	{Name: "AvailabilityZoneId", Flag: "availability-zone-id", Type: "*string", Required: false},
	{Name: "AzMode", Flag: "az-mode", Type: "types.KxAzMode", Required: true},
	{Name: "CacheStorageConfigurations", Flag: "cache-storage-configurations", Type: "[]types.KxCacheStorageConfiguration", Required: false},
	{Name: "CapacityConfiguration", Flag: "capacity-configuration", Type: "*types.CapacityConfiguration", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ClusterDescription", Flag: "cluster-description", Type: "*string", Required: false},
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "ClusterType", Flag: "cluster-type", Type: "types.KxClusterType", Required: true},
	{Name: "Code", Flag: "code", Type: "*types.CodeConfiguration", Required: false},
	{Name: "CommandLineArguments", Flag: "command-line-arguments", Type: "[]types.KxCommandLineArgument", Required: false},
	{Name: "Databases", Flag: "databases", Type: "[]types.KxDatabaseConfiguration", Required: false},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
	{Name: "ExecutionRole", Flag: "execution-role", Type: "*string", Required: false},
	{Name: "InitializationScript", Flag: "initialization-script", Type: "*string", Required: false},
	{Name: "ReleaseLabel", Flag: "release-label", Type: "*string", Required: true},
	{Name: "SavedownStorageConfiguration", Flag: "savedown-storage-configuration", Type: "*types.KxSavedownStorageConfiguration", Required: false},
	{Name: "ScalingGroupConfiguration", Flag: "scaling-group-configuration", Type: "*types.KxScalingGroupConfiguration", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TickerplantLogConfiguration", Flag: "tickerplant-log-configuration", Type: "*types.TickerplantLogConfiguration", Required: false},
	{Name: "VpcConfiguration", Flag: "vpc-configuration", Type: "*types.VpcConfiguration", Required: true},
}

var fields_create_kx_database = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_kx_dataview = []leanruntime.Field{
	{Name: "AutoUpdate", Flag: "auto-update", Type: "bool", Required: false},
	{Name: "AvailabilityZoneId", Flag: "availability-zone-id", Type: "*string", Required: false},
	{Name: "AzMode", Flag: "az-mode", Type: "types.KxAzMode", Required: true},
	{Name: "ChangesetId", Flag: "changeset-id", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "DataviewName", Flag: "dataview-name", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
	{Name: "ReadWrite", Flag: "read-write", Type: "bool", Required: false},
	{Name: "SegmentConfigurations", Flag: "segment-configurations", Type: "[]types.KxDataviewSegmentConfiguration", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_kx_environment = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_kx_scaling_group = []leanruntime.Field{
	{Name: "AvailabilityZoneId", Flag: "availability-zone-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
	{Name: "HostType", Flag: "host-type", Type: "*string", Required: true},
	{Name: "ScalingGroupName", Flag: "scaling-group-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_kx_user = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
	{Name: "IamRole", Flag: "iam-role", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_create_kx_volume = []leanruntime.Field{
	{Name: "AvailabilityZoneIds", Flag: "availability-zone-ids", Type: "[]string", Required: true},
	{Name: "AzMode", Flag: "az-mode", Type: "types.KxAzMode", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
	{Name: "Nas1Configuration", Flag: "nas1-configuration", Type: "*types.KxNAS1Configuration", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "VolumeName", Flag: "volume-name", Type: "*string", Required: true},
	{Name: "VolumeType", Flag: "volume-type", Type: "types.KxVolumeType", Required: true},
}

var fields_delete_environment = []leanruntime.Field{
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
}

var fields_delete_kx_cluster = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
}

var fields_delete_kx_cluster_node = []leanruntime.Field{
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
	{Name: "NodeId", Flag: "node-id", Type: "*string", Required: true},
}

var fields_delete_kx_database = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
}

var fields_delete_kx_dataview = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "DataviewName", Flag: "dataview-name", Type: "*string", Required: true},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
}

var fields_delete_kx_environment = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
}

var fields_delete_kx_scaling_group = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
	{Name: "ScalingGroupName", Flag: "scaling-group-name", Type: "*string", Required: true},
}

var fields_delete_kx_user = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_delete_kx_volume = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
	{Name: "VolumeName", Flag: "volume-name", Type: "*string", Required: true},
}

var fields_get_environment = []leanruntime.Field{
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
}

var fields_get_kx_changeset = []leanruntime.Field{
	{Name: "ChangesetId", Flag: "changeset-id", Type: "*string", Required: true},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
}

var fields_get_kx_cluster = []leanruntime.Field{
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
}

var fields_get_kx_connection_string = []leanruntime.Field{
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
	{Name: "UserArn", Flag: "user-arn", Type: "*string", Required: true},
}

var fields_get_kx_database = []leanruntime.Field{
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
}

var fields_get_kx_dataview = []leanruntime.Field{
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "DataviewName", Flag: "dataview-name", Type: "*string", Required: true},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
}

var fields_get_kx_environment = []leanruntime.Field{
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
}

var fields_get_kx_scaling_group = []leanruntime.Field{
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
	{Name: "ScalingGroupName", Flag: "scaling-group-name", Type: "*string", Required: true},
}

var fields_get_kx_user = []leanruntime.Field{
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_get_kx_volume = []leanruntime.Field{
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
	{Name: "VolumeName", Flag: "volume-name", Type: "*string", Required: true},
}

var fields_list_environments = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_kx_changesets = []leanruntime.Field{
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_kx_cluster_nodes = []leanruntime.Field{
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_kx_clusters = []leanruntime.Field{
	{Name: "ClusterType", Flag: "cluster-type", Type: "types.KxClusterType", Required: false},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_kx_databases = []leanruntime.Field{
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_kx_dataviews = []leanruntime.Field{
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_kx_environments = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_kx_scaling_groups = []leanruntime.Field{
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_kx_users = []leanruntime.Field{
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_kx_volumes = []leanruntime.Field{
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "VolumeType", Flag: "volume-type", Type: "types.KxVolumeType", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_environment = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
	{Name: "FederationMode", Flag: "federation-mode", Type: "types.FederationMode", Required: false},
	{Name: "FederationParameters", Flag: "federation-parameters", Type: "*types.FederationParameters", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_kx_cluster_code_configuration = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "Code", Flag: "code", Type: "*types.CodeConfiguration", Required: true},
	{Name: "CommandLineArguments", Flag: "command-line-arguments", Type: "[]types.KxCommandLineArgument", Required: false},
	{Name: "DeploymentConfiguration", Flag: "deployment-configuration", Type: "*types.KxClusterCodeDeploymentConfiguration", Required: false},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
	{Name: "InitializationScript", Flag: "initialization-script", Type: "*string", Required: false},
}

var fields_update_kx_cluster_databases = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "Databases", Flag: "databases", Type: "[]types.KxDatabaseConfiguration", Required: true},
	{Name: "DeploymentConfiguration", Flag: "deployment-configuration", Type: "*types.KxDeploymentConfiguration", Required: false},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
}

var fields_update_kx_database = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
}

var fields_update_kx_dataview = []leanruntime.Field{
	{Name: "ChangesetId", Flag: "changeset-id", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "DataviewName", Flag: "dataview-name", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
	{Name: "SegmentConfigurations", Flag: "segment-configurations", Type: "[]types.KxDataviewSegmentConfiguration", Required: false},
}

var fields_update_kx_environment = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_kx_environment_network = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CustomDNSConfiguration", Flag: "custom-dns-configuration", Type: "[]types.CustomDNSServer", Required: false},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
	{Name: "TransitGatewayConfiguration", Flag: "transit-gateway-configuration", Type: "*types.TransitGatewayConfiguration", Required: false},
}

var fields_update_kx_user = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
	{Name: "IamRole", Flag: "iam-role", Type: "*string", Required: true},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_update_kx_volume = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
	{Name: "Nas1Configuration", Flag: "nas1-configuration", Type: "*types.KxNAS1Configuration", Required: false},
	{Name: "VolumeName", Flag: "volume-name", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-environment": {
			Name:   "create-environment",
			Fields: fields_create_environment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEnvironmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_environment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEnvironment(ctx, input)
			},
		},
		"create-kx-changeset": {
			Name:   "create-kx-changeset",
			Fields: fields_create_kx_changeset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateKxChangesetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_kx_changeset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateKxChangeset(ctx, input)
			},
		},
		"create-kx-cluster": {
			Name:   "create-kx-cluster",
			Fields: fields_create_kx_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateKxClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_kx_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateKxCluster(ctx, input)
			},
		},
		"create-kx-database": {
			Name:   "create-kx-database",
			Fields: fields_create_kx_database,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateKxDatabaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_kx_database, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateKxDatabase(ctx, input)
			},
		},
		"create-kx-dataview": {
			Name:   "create-kx-dataview",
			Fields: fields_create_kx_dataview,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateKxDataviewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_kx_dataview, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateKxDataview(ctx, input)
			},
		},
		"create-kx-environment": {
			Name:   "create-kx-environment",
			Fields: fields_create_kx_environment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateKxEnvironmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_kx_environment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateKxEnvironment(ctx, input)
			},
		},
		"create-kx-scaling-group": {
			Name:   "create-kx-scaling-group",
			Fields: fields_create_kx_scaling_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateKxScalingGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_kx_scaling_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateKxScalingGroup(ctx, input)
			},
		},
		"create-kx-user": {
			Name:   "create-kx-user",
			Fields: fields_create_kx_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateKxUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_kx_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateKxUser(ctx, input)
			},
		},
		"create-kx-volume": {
			Name:   "create-kx-volume",
			Fields: fields_create_kx_volume,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateKxVolumeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_kx_volume, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateKxVolume(ctx, input)
			},
		},
		"delete-environment": {
			Name:   "delete-environment",
			Fields: fields_delete_environment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEnvironmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_environment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEnvironment(ctx, input)
			},
		},
		"delete-kx-cluster": {
			Name:   "delete-kx-cluster",
			Fields: fields_delete_kx_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteKxClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_kx_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteKxCluster(ctx, input)
			},
		},
		"delete-kx-cluster-node": {
			Name:   "delete-kx-cluster-node",
			Fields: fields_delete_kx_cluster_node,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteKxClusterNodeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_kx_cluster_node, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteKxClusterNode(ctx, input)
			},
		},
		"delete-kx-database": {
			Name:   "delete-kx-database",
			Fields: fields_delete_kx_database,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteKxDatabaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_kx_database, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteKxDatabase(ctx, input)
			},
		},
		"delete-kx-dataview": {
			Name:   "delete-kx-dataview",
			Fields: fields_delete_kx_dataview,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteKxDataviewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_kx_dataview, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteKxDataview(ctx, input)
			},
		},
		"delete-kx-environment": {
			Name:   "delete-kx-environment",
			Fields: fields_delete_kx_environment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteKxEnvironmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_kx_environment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteKxEnvironment(ctx, input)
			},
		},
		"delete-kx-scaling-group": {
			Name:   "delete-kx-scaling-group",
			Fields: fields_delete_kx_scaling_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteKxScalingGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_kx_scaling_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteKxScalingGroup(ctx, input)
			},
		},
		"delete-kx-user": {
			Name:   "delete-kx-user",
			Fields: fields_delete_kx_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteKxUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_kx_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteKxUser(ctx, input)
			},
		},
		"delete-kx-volume": {
			Name:   "delete-kx-volume",
			Fields: fields_delete_kx_volume,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteKxVolumeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_kx_volume, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteKxVolume(ctx, input)
			},
		},
		"get-environment": {
			Name:   "get-environment",
			Fields: fields_get_environment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEnvironmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_environment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEnvironment(ctx, input)
			},
		},
		"get-kx-changeset": {
			Name:   "get-kx-changeset",
			Fields: fields_get_kx_changeset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetKxChangesetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_kx_changeset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetKxChangeset(ctx, input)
			},
		},
		"get-kx-cluster": {
			Name:   "get-kx-cluster",
			Fields: fields_get_kx_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetKxClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_kx_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetKxCluster(ctx, input)
			},
		},
		"get-kx-connection-string": {
			Name:   "get-kx-connection-string",
			Fields: fields_get_kx_connection_string,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetKxConnectionStringInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_kx_connection_string, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetKxConnectionString(ctx, input)
			},
		},
		"get-kx-database": {
			Name:   "get-kx-database",
			Fields: fields_get_kx_database,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetKxDatabaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_kx_database, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetKxDatabase(ctx, input)
			},
		},
		"get-kx-dataview": {
			Name:   "get-kx-dataview",
			Fields: fields_get_kx_dataview,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetKxDataviewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_kx_dataview, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetKxDataview(ctx, input)
			},
		},
		"get-kx-environment": {
			Name:   "get-kx-environment",
			Fields: fields_get_kx_environment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetKxEnvironmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_kx_environment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetKxEnvironment(ctx, input)
			},
		},
		"get-kx-scaling-group": {
			Name:   "get-kx-scaling-group",
			Fields: fields_get_kx_scaling_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetKxScalingGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_kx_scaling_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetKxScalingGroup(ctx, input)
			},
		},
		"get-kx-user": {
			Name:   "get-kx-user",
			Fields: fields_get_kx_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetKxUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_kx_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetKxUser(ctx, input)
			},
		},
		"get-kx-volume": {
			Name:   "get-kx-volume",
			Fields: fields_get_kx_volume,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetKxVolumeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_kx_volume, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetKxVolume(ctx, input)
			},
		},
		"list-environments": {
			Name:   "list-environments",
			Fields: fields_list_environments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEnvironmentsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_environments, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListEnvironments(ctx, input)
			},
		},
		"list-kx-changesets": {
			Name:   "list-kx-changesets",
			Fields: fields_list_kx_changesets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListKxChangesetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_kx_changesets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListKxChangesets(ctx, input)
				}
				var results []*svc.ListKxChangesetsOutput
				p := svc.NewListKxChangesetsPaginator(client, input)
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
		"list-kx-cluster-nodes": {
			Name:   "list-kx-cluster-nodes",
			Fields: fields_list_kx_cluster_nodes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListKxClusterNodesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_kx_cluster_nodes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListKxClusterNodes(ctx, input)
				}
				var results []*svc.ListKxClusterNodesOutput
				p := svc.NewListKxClusterNodesPaginator(client, input)
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
		"list-kx-clusters": {
			Name:   "list-kx-clusters",
			Fields: fields_list_kx_clusters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListKxClustersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_kx_clusters, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListKxClusters(ctx, input)
			},
		},
		"list-kx-databases": {
			Name:   "list-kx-databases",
			Fields: fields_list_kx_databases,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListKxDatabasesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_kx_databases, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListKxDatabases(ctx, input)
				}
				var results []*svc.ListKxDatabasesOutput
				p := svc.NewListKxDatabasesPaginator(client, input)
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
		"list-kx-dataviews": {
			Name:   "list-kx-dataviews",
			Fields: fields_list_kx_dataviews,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListKxDataviewsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_kx_dataviews, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListKxDataviews(ctx, input)
				}
				var results []*svc.ListKxDataviewsOutput
				p := svc.NewListKxDataviewsPaginator(client, input)
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
		"list-kx-environments": {
			Name:   "list-kx-environments",
			Fields: fields_list_kx_environments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListKxEnvironmentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_kx_environments, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListKxEnvironments(ctx, input)
				}
				var results []*svc.ListKxEnvironmentsOutput
				p := svc.NewListKxEnvironmentsPaginator(client, input)
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
		"list-kx-scaling-groups": {
			Name:   "list-kx-scaling-groups",
			Fields: fields_list_kx_scaling_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListKxScalingGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_kx_scaling_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListKxScalingGroups(ctx, input)
				}
				var results []*svc.ListKxScalingGroupsOutput
				p := svc.NewListKxScalingGroupsPaginator(client, input)
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
		"list-kx-users": {
			Name:   "list-kx-users",
			Fields: fields_list_kx_users,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListKxUsersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_kx_users, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListKxUsers(ctx, input)
			},
		},
		"list-kx-volumes": {
			Name:   "list-kx-volumes",
			Fields: fields_list_kx_volumes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListKxVolumesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_kx_volumes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListKxVolumes(ctx, input)
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
		"update-environment": {
			Name:   "update-environment",
			Fields: fields_update_environment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEnvironmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_environment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEnvironment(ctx, input)
			},
		},
		"update-kx-cluster-code-configuration": {
			Name:   "update-kx-cluster-code-configuration",
			Fields: fields_update_kx_cluster_code_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateKxClusterCodeConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_kx_cluster_code_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateKxClusterCodeConfiguration(ctx, input)
			},
		},
		"update-kx-cluster-databases": {
			Name:   "update-kx-cluster-databases",
			Fields: fields_update_kx_cluster_databases,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateKxClusterDatabasesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_kx_cluster_databases, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateKxClusterDatabases(ctx, input)
			},
		},
		"update-kx-database": {
			Name:   "update-kx-database",
			Fields: fields_update_kx_database,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateKxDatabaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_kx_database, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateKxDatabase(ctx, input)
			},
		},
		"update-kx-dataview": {
			Name:   "update-kx-dataview",
			Fields: fields_update_kx_dataview,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateKxDataviewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_kx_dataview, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateKxDataview(ctx, input)
			},
		},
		"update-kx-environment": {
			Name:   "update-kx-environment",
			Fields: fields_update_kx_environment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateKxEnvironmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_kx_environment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateKxEnvironment(ctx, input)
			},
		},
		"update-kx-environment-network": {
			Name:   "update-kx-environment-network",
			Fields: fields_update_kx_environment_network,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateKxEnvironmentNetworkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_kx_environment_network, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateKxEnvironmentNetwork(ctx, input)
			},
		},
		"update-kx-user": {
			Name:   "update-kx-user",
			Fields: fields_update_kx_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateKxUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_kx_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateKxUser(ctx, input)
			},
		},
		"update-kx-volume": {
			Name:   "update-kx-volume",
			Fields: fields_update_kx_volume,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateKxVolumeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_kx_volume, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateKxVolume(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("finspace", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
