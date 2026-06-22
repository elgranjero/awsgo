package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/devicefarm"
)

var fields_create_device_pool = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "MaxDevices", Flag: "max-devices", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ProjectArn", Flag: "project-arn", Type: "*string", Required: true},
	{Name: "Rules", Flag: "rules", Type: "[]types.Rule", Required: true},
}

var fields_create_instance_profile = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ExcludeAppPackagesFromCleanup", Flag: "exclude-app-packages-from-cleanup", Type: "[]string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "PackageCleanup", Flag: "package-cleanup", Type: "*bool", Required: false},
	{Name: "RebootAfterUse", Flag: "reboot-after-use", Type: "*bool", Required: false},
}

var fields_create_network_profile = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DownlinkBandwidthBits", Flag: "downlink-bandwidth-bits", Type: "*int64", Required: false},
	{Name: "DownlinkDelayMs", Flag: "downlink-delay-ms", Type: "*int64", Required: false},
	{Name: "DownlinkJitterMs", Flag: "downlink-jitter-ms", Type: "*int64", Required: false},
	{Name: "DownlinkLossPercent", Flag: "downlink-loss-percent", Type: "int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ProjectArn", Flag: "project-arn", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.NetworkProfileType", Required: false},
	{Name: "UplinkBandwidthBits", Flag: "uplink-bandwidth-bits", Type: "*int64", Required: false},
	{Name: "UplinkDelayMs", Flag: "uplink-delay-ms", Type: "*int64", Required: false},
	{Name: "UplinkJitterMs", Flag: "uplink-jitter-ms", Type: "*int64", Required: false},
	{Name: "UplinkLossPercent", Flag: "uplink-loss-percent", Type: "int32", Required: false},
}

var fields_create_project = []leanruntime.Field{
	{Name: "DefaultJobTimeoutMinutes", Flag: "default-job-timeout-minutes", Type: "*int32", Required: false},
	{Name: "EnvironmentVariables", Flag: "environment-variables", Type: "[]types.EnvironmentVariable", Required: false},
	{Name: "ExecutionRoleArn", Flag: "execution-role-arn", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "VpcConfig", Flag: "vpc-config", Type: "*types.VpcConfig", Required: false},
}

var fields_create_remote_access_session = []leanruntime.Field{
	{Name: "AppArn", Flag: "app-arn", Type: "*string", Required: false},
	{Name: "Configuration", Flag: "configuration", Type: "*types.CreateRemoteAccessSessionConfiguration", Required: false},
	{Name: "DeviceArn", Flag: "device-arn", Type: "*string", Required: true},
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: false},
	{Name: "InteractionMode", Flag: "interaction-mode", Type: "types.InteractionMode", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "ProjectArn", Flag: "project-arn", Type: "*string", Required: true},
	{Name: "SkipAppResign", Flag: "skip-app-resign", Type: "*bool", Required: false},
}

var fields_create_test_grid_project = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "VpcConfig", Flag: "vpc-config", Type: "*types.TestGridVpcConfig", Required: false},
}

var fields_create_test_grid_url = []leanruntime.Field{
	{Name: "ExpiresInSeconds", Flag: "expires-in-seconds", Type: "*int32", Required: true},
	{Name: "ProjectArn", Flag: "project-arn", Type: "*string", Required: true},
}

var fields_create_upload = []leanruntime.Field{
	{Name: "ContentType", Flag: "content-type", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ProjectArn", Flag: "project-arn", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.UploadType", Required: true},
}

var fields_create_vpce_configuration = []leanruntime.Field{
	{Name: "ServiceDnsName", Flag: "service-dns-name", Type: "*string", Required: true},
	{Name: "VpceConfigurationDescription", Flag: "vpce-configuration-description", Type: "*string", Required: false},
	{Name: "VpceConfigurationName", Flag: "vpce-configuration-name", Type: "*string", Required: true},
	{Name: "VpceServiceName", Flag: "vpce-service-name", Type: "*string", Required: true},
}

var fields_delete_device_pool = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_delete_instance_profile = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_delete_network_profile = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_delete_project = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_delete_remote_access_session = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_delete_run = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_delete_test_grid_project = []leanruntime.Field{
	{Name: "ProjectArn", Flag: "project-arn", Type: "*string", Required: true},
}

var fields_delete_upload = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_delete_vpce_configuration = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_account_settings = []leanruntime.Field{}

var fields_get_device = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_device_instance = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_device_pool = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_device_pool_compatibility = []leanruntime.Field{
	{Name: "AppArn", Flag: "app-arn", Type: "*string", Required: false},
	{Name: "Configuration", Flag: "configuration", Type: "*types.ScheduleRunConfiguration", Required: false},
	{Name: "DevicePoolArn", Flag: "device-pool-arn", Type: "*string", Required: true},
	{Name: "ProjectArn", Flag: "project-arn", Type: "*string", Required: false},
	{Name: "Test", Flag: "test", Type: "*types.ScheduleRunTest", Required: false},
	{Name: "TestType", Flag: "test-type", Type: "types.TestType", Required: false},
}

var fields_get_instance_profile = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_job = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_network_profile = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_offering_status = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_project = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_remote_access_session = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_run = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_suite = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_test = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_test_grid_project = []leanruntime.Field{
	{Name: "ProjectArn", Flag: "project-arn", Type: "*string", Required: true},
}

var fields_get_test_grid_session = []leanruntime.Field{
	{Name: "ProjectArn", Flag: "project-arn", Type: "*string", Required: false},
	{Name: "SessionArn", Flag: "session-arn", Type: "*string", Required: false},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: false},
}

var fields_get_upload = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_vpce_configuration = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_install_to_remote_access_session = []leanruntime.Field{
	{Name: "AppArn", Flag: "app-arn", Type: "*string", Required: true},
	{Name: "RemoteAccessSessionArn", Flag: "remote-access-session-arn", Type: "*string", Required: true},
}

var fields_list_artifacts = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.ArtifactCategory", Required: true},
}

var fields_list_device_instances = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_device_pools = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.DevicePoolType", Required: false},
}

var fields_list_devices = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.DeviceFilter", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_instance_profiles = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_jobs = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_network_profiles = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.NetworkProfileType", Required: false},
}

var fields_list_offering_promotions = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_offering_transactions = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_offerings = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_projects = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_remote_access_sessions = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_runs = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_samples = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_suites = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_test_grid_projects = []leanruntime.Field{
	{Name: "MaxResult", Flag: "max-result", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_test_grid_session_actions = []leanruntime.Field{
	{Name: "MaxResult", Flag: "max-result", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SessionArn", Flag: "session-arn", Type: "*string", Required: true},
}

var fields_list_test_grid_session_artifacts = []leanruntime.Field{
	{Name: "MaxResult", Flag: "max-result", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SessionArn", Flag: "session-arn", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.TestGridSessionArtifactCategory", Required: false},
}

var fields_list_test_grid_sessions = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "EndTimeAfter", Flag: "end-time-after", Type: "*time.Time", Required: false},
	{Name: "EndTimeBefore", Flag: "end-time-before", Type: "*time.Time", Required: false},
	{Name: "MaxResult", Flag: "max-result", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProjectArn", Flag: "project-arn", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.TestGridSessionStatus", Required: false},
}

var fields_list_tests = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_unique_problems = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_uploads = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.UploadType", Required: false},
}

var fields_list_vpce_configurations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_purchase_offering = []leanruntime.Field{
	{Name: "OfferingId", Flag: "offering-id", Type: "*string", Required: true},
	{Name: "OfferingPromotionId", Flag: "offering-promotion-id", Type: "*string", Required: false},
	{Name: "Quantity", Flag: "quantity", Type: "*int32", Required: true},
}

var fields_renew_offering = []leanruntime.Field{
	{Name: "OfferingId", Flag: "offering-id", Type: "*string", Required: true},
	{Name: "Quantity", Flag: "quantity", Type: "*int32", Required: true},
}

var fields_schedule_run = []leanruntime.Field{
	{Name: "AppArn", Flag: "app-arn", Type: "*string", Required: false},
	{Name: "Configuration", Flag: "configuration", Type: "*types.ScheduleRunConfiguration", Required: false},
	{Name: "DevicePoolArn", Flag: "device-pool-arn", Type: "*string", Required: false},
	{Name: "DeviceSelectionConfiguration", Flag: "device-selection-configuration", Type: "*types.DeviceSelectionConfiguration", Required: false},
	{Name: "ExecutionConfiguration", Flag: "execution-configuration", Type: "*types.ExecutionConfiguration", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "ProjectArn", Flag: "project-arn", Type: "*string", Required: true},
	{Name: "Test", Flag: "test", Type: "*types.ScheduleRunTest", Required: true},
}

var fields_stop_job = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_stop_remote_access_session = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_stop_run = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_device_instance = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "Labels", Flag: "labels", Type: "[]string", Required: false},
	{Name: "ProfileArn", Flag: "profile-arn", Type: "*string", Required: false},
}

var fields_update_device_pool = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "ClearMaxDevices", Flag: "clear-max-devices", Type: "*bool", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "MaxDevices", Flag: "max-devices", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Rules", Flag: "rules", Type: "[]types.Rule", Required: false},
}

var fields_update_instance_profile = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ExcludeAppPackagesFromCleanup", Flag: "exclude-app-packages-from-cleanup", Type: "[]string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "PackageCleanup", Flag: "package-cleanup", Type: "*bool", Required: false},
	{Name: "RebootAfterUse", Flag: "reboot-after-use", Type: "*bool", Required: false},
}

var fields_update_network_profile = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DownlinkBandwidthBits", Flag: "downlink-bandwidth-bits", Type: "*int64", Required: false},
	{Name: "DownlinkDelayMs", Flag: "downlink-delay-ms", Type: "*int64", Required: false},
	{Name: "DownlinkJitterMs", Flag: "downlink-jitter-ms", Type: "*int64", Required: false},
	{Name: "DownlinkLossPercent", Flag: "downlink-loss-percent", Type: "int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.NetworkProfileType", Required: false},
	{Name: "UplinkBandwidthBits", Flag: "uplink-bandwidth-bits", Type: "*int64", Required: false},
	{Name: "UplinkDelayMs", Flag: "uplink-delay-ms", Type: "*int64", Required: false},
	{Name: "UplinkJitterMs", Flag: "uplink-jitter-ms", Type: "*int64", Required: false},
	{Name: "UplinkLossPercent", Flag: "uplink-loss-percent", Type: "int32", Required: false},
}

var fields_update_project = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "DefaultJobTimeoutMinutes", Flag: "default-job-timeout-minutes", Type: "*int32", Required: false},
	{Name: "EnvironmentVariables", Flag: "environment-variables", Type: "[]types.EnvironmentVariable", Required: false},
	{Name: "ExecutionRoleArn", Flag: "execution-role-arn", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "VpcConfig", Flag: "vpc-config", Type: "*types.VpcConfig", Required: false},
}

var fields_update_test_grid_project = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "ProjectArn", Flag: "project-arn", Type: "*string", Required: true},
	{Name: "VpcConfig", Flag: "vpc-config", Type: "*types.TestGridVpcConfig", Required: false},
}

var fields_update_upload = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "ContentType", Flag: "content-type", Type: "*string", Required: false},
	{Name: "EditContent", Flag: "edit-content", Type: "*bool", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_vpce_configuration = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "ServiceDnsName", Flag: "service-dns-name", Type: "*string", Required: false},
	{Name: "VpceConfigurationDescription", Flag: "vpce-configuration-description", Type: "*string", Required: false},
	{Name: "VpceConfigurationName", Flag: "vpce-configuration-name", Type: "*string", Required: false},
	{Name: "VpceServiceName", Flag: "vpce-service-name", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-device-pool": {
			Name:   "create-device-pool",
			Fields: fields_create_device_pool,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDevicePoolInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_device_pool, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDevicePool(ctx, input)
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
		"create-network-profile": {
			Name:   "create-network-profile",
			Fields: fields_create_network_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateNetworkProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_network_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateNetworkProfile(ctx, input)
			},
		},
		"create-project": {
			Name:   "create-project",
			Fields: fields_create_project,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateProjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_project, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateProject(ctx, input)
			},
		},
		"create-remote-access-session": {
			Name:   "create-remote-access-session",
			Fields: fields_create_remote_access_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRemoteAccessSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_remote_access_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRemoteAccessSession(ctx, input)
			},
		},
		"create-test-grid-project": {
			Name:   "create-test-grid-project",
			Fields: fields_create_test_grid_project,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTestGridProjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_test_grid_project, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTestGridProject(ctx, input)
			},
		},
		"create-test-grid-url": {
			Name:   "create-test-grid-url",
			Fields: fields_create_test_grid_url,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTestGridUrlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_test_grid_url, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTestGridUrl(ctx, input)
			},
		},
		"create-upload": {
			Name:   "create-upload",
			Fields: fields_create_upload,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateUploadInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_upload, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateUpload(ctx, input)
			},
		},
		"create-vpce-configuration": {
			Name:   "create-vpce-configuration",
			Fields: fields_create_vpce_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVPCEConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_vpce_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVPCEConfiguration(ctx, input)
			},
		},
		"delete-device-pool": {
			Name:   "delete-device-pool",
			Fields: fields_delete_device_pool,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDevicePoolInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_device_pool, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDevicePool(ctx, input)
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
		"delete-network-profile": {
			Name:   "delete-network-profile",
			Fields: fields_delete_network_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteNetworkProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_network_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteNetworkProfile(ctx, input)
			},
		},
		"delete-project": {
			Name:   "delete-project",
			Fields: fields_delete_project,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteProjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_project, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteProject(ctx, input)
			},
		},
		"delete-remote-access-session": {
			Name:   "delete-remote-access-session",
			Fields: fields_delete_remote_access_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRemoteAccessSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_remote_access_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRemoteAccessSession(ctx, input)
			},
		},
		"delete-run": {
			Name:   "delete-run",
			Fields: fields_delete_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRun(ctx, input)
			},
		},
		"delete-test-grid-project": {
			Name:   "delete-test-grid-project",
			Fields: fields_delete_test_grid_project,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTestGridProjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_test_grid_project, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTestGridProject(ctx, input)
			},
		},
		"delete-upload": {
			Name:   "delete-upload",
			Fields: fields_delete_upload,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteUploadInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_upload, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteUpload(ctx, input)
			},
		},
		"delete-vpce-configuration": {
			Name:   "delete-vpce-configuration",
			Fields: fields_delete_vpce_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVPCEConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_vpce_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVPCEConfiguration(ctx, input)
			},
		},
		"get-account-settings": {
			Name:   "get-account-settings",
			Fields: fields_get_account_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccountSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_account_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccountSettings(ctx, input)
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
		"get-device-instance": {
			Name:   "get-device-instance",
			Fields: fields_get_device_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDeviceInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_device_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDeviceInstance(ctx, input)
			},
		},
		"get-device-pool": {
			Name:   "get-device-pool",
			Fields: fields_get_device_pool,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDevicePoolInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_device_pool, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDevicePool(ctx, input)
			},
		},
		"get-device-pool-compatibility": {
			Name:   "get-device-pool-compatibility",
			Fields: fields_get_device_pool_compatibility,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDevicePoolCompatibilityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_device_pool_compatibility, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDevicePoolCompatibility(ctx, input)
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
		"get-job": {
			Name:   "get-job",
			Fields: fields_get_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetJob(ctx, input)
			},
		},
		"get-network-profile": {
			Name:   "get-network-profile",
			Fields: fields_get_network_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetNetworkProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_network_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetNetworkProfile(ctx, input)
			},
		},
		"get-offering-status": {
			Name:   "get-offering-status",
			Fields: fields_get_offering_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOfferingStatusInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_offering_status, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetOfferingStatus(ctx, input)
				}
				var results []*svc.GetOfferingStatusOutput
				p := svc.NewGetOfferingStatusPaginator(client, input)
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
		"get-project": {
			Name:   "get-project",
			Fields: fields_get_project,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetProjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_project, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetProject(ctx, input)
			},
		},
		"get-remote-access-session": {
			Name:   "get-remote-access-session",
			Fields: fields_get_remote_access_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRemoteAccessSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_remote_access_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRemoteAccessSession(ctx, input)
			},
		},
		"get-run": {
			Name:   "get-run",
			Fields: fields_get_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRun(ctx, input)
			},
		},
		"get-suite": {
			Name:   "get-suite",
			Fields: fields_get_suite,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSuiteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_suite, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSuite(ctx, input)
			},
		},
		"get-test": {
			Name:   "get-test",
			Fields: fields_get_test,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTestInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_test, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTest(ctx, input)
			},
		},
		"get-test-grid-project": {
			Name:   "get-test-grid-project",
			Fields: fields_get_test_grid_project,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTestGridProjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_test_grid_project, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTestGridProject(ctx, input)
			},
		},
		"get-test-grid-session": {
			Name:   "get-test-grid-session",
			Fields: fields_get_test_grid_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTestGridSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_test_grid_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTestGridSession(ctx, input)
			},
		},
		"get-upload": {
			Name:   "get-upload",
			Fields: fields_get_upload,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetUploadInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_upload, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetUpload(ctx, input)
			},
		},
		"get-vpce-configuration": {
			Name:   "get-vpce-configuration",
			Fields: fields_get_vpce_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetVPCEConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_vpce_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetVPCEConfiguration(ctx, input)
			},
		},
		"install-to-remote-access-session": {
			Name:   "install-to-remote-access-session",
			Fields: fields_install_to_remote_access_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.InstallToRemoteAccessSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_install_to_remote_access_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.InstallToRemoteAccessSession(ctx, input)
			},
		},
		"list-artifacts": {
			Name:   "list-artifacts",
			Fields: fields_list_artifacts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListArtifactsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_artifacts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListArtifacts(ctx, input)
				}
				var results []*svc.ListArtifactsOutput
				p := svc.NewListArtifactsPaginator(client, input)
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
		"list-device-instances": {
			Name:   "list-device-instances",
			Fields: fields_list_device_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDeviceInstancesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_device_instances, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListDeviceInstances(ctx, input)
			},
		},
		"list-device-pools": {
			Name:   "list-device-pools",
			Fields: fields_list_device_pools,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDevicePoolsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_device_pools, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDevicePools(ctx, input)
				}
				var results []*svc.ListDevicePoolsOutput
				p := svc.NewListDevicePoolsPaginator(client, input)
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
		"list-devices": {
			Name:   "list-devices",
			Fields: fields_list_devices,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDevicesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_devices, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDevices(ctx, input)
				}
				var results []*svc.ListDevicesOutput
				p := svc.NewListDevicesPaginator(client, input)
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
				if _, err := leanruntime.ApplyInput(input, fields_list_instance_profiles, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListInstanceProfiles(ctx, input)
			},
		},
		"list-jobs": {
			Name:   "list-jobs",
			Fields: fields_list_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListJobs(ctx, input)
				}
				var results []*svc.ListJobsOutput
				p := svc.NewListJobsPaginator(client, input)
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
		"list-network-profiles": {
			Name:   "list-network-profiles",
			Fields: fields_list_network_profiles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListNetworkProfilesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_network_profiles, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListNetworkProfiles(ctx, input)
			},
		},
		"list-offering-promotions": {
			Name:   "list-offering-promotions",
			Fields: fields_list_offering_promotions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOfferingPromotionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_offering_promotions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListOfferingPromotions(ctx, input)
			},
		},
		"list-offering-transactions": {
			Name:   "list-offering-transactions",
			Fields: fields_list_offering_transactions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOfferingTransactionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_offering_transactions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListOfferingTransactions(ctx, input)
				}
				var results []*svc.ListOfferingTransactionsOutput
				p := svc.NewListOfferingTransactionsPaginator(client, input)
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
		"list-offerings": {
			Name:   "list-offerings",
			Fields: fields_list_offerings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOfferingsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_offerings, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListOfferings(ctx, input)
				}
				var results []*svc.ListOfferingsOutput
				p := svc.NewListOfferingsPaginator(client, input)
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
		"list-projects": {
			Name:   "list-projects",
			Fields: fields_list_projects,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProjectsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_projects, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProjects(ctx, input)
				}
				var results []*svc.ListProjectsOutput
				p := svc.NewListProjectsPaginator(client, input)
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
		"list-remote-access-sessions": {
			Name:   "list-remote-access-sessions",
			Fields: fields_list_remote_access_sessions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRemoteAccessSessionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_remote_access_sessions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListRemoteAccessSessions(ctx, input)
			},
		},
		"list-runs": {
			Name:   "list-runs",
			Fields: fields_list_runs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRunsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_runs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRuns(ctx, input)
				}
				var results []*svc.ListRunsOutput
				p := svc.NewListRunsPaginator(client, input)
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
		"list-samples": {
			Name:   "list-samples",
			Fields: fields_list_samples,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSamplesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_samples, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSamples(ctx, input)
				}
				var results []*svc.ListSamplesOutput
				p := svc.NewListSamplesPaginator(client, input)
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
		"list-suites": {
			Name:   "list-suites",
			Fields: fields_list_suites,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSuitesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_suites, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSuites(ctx, input)
				}
				var results []*svc.ListSuitesOutput
				p := svc.NewListSuitesPaginator(client, input)
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
		"list-test-grid-projects": {
			Name:   "list-test-grid-projects",
			Fields: fields_list_test_grid_projects,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTestGridProjectsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_test_grid_projects, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTestGridProjects(ctx, input)
				}
				var results []*svc.ListTestGridProjectsOutput
				p := svc.NewListTestGridProjectsPaginator(client, input)
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
		"list-test-grid-session-actions": {
			Name:   "list-test-grid-session-actions",
			Fields: fields_list_test_grid_session_actions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTestGridSessionActionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_test_grid_session_actions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTestGridSessionActions(ctx, input)
				}
				var results []*svc.ListTestGridSessionActionsOutput
				p := svc.NewListTestGridSessionActionsPaginator(client, input)
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
		"list-test-grid-session-artifacts": {
			Name:   "list-test-grid-session-artifacts",
			Fields: fields_list_test_grid_session_artifacts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTestGridSessionArtifactsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_test_grid_session_artifacts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTestGridSessionArtifacts(ctx, input)
				}
				var results []*svc.ListTestGridSessionArtifactsOutput
				p := svc.NewListTestGridSessionArtifactsPaginator(client, input)
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
		"list-test-grid-sessions": {
			Name:   "list-test-grid-sessions",
			Fields: fields_list_test_grid_sessions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTestGridSessionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_test_grid_sessions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTestGridSessions(ctx, input)
				}
				var results []*svc.ListTestGridSessionsOutput
				p := svc.NewListTestGridSessionsPaginator(client, input)
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
		"list-tests": {
			Name:   "list-tests",
			Fields: fields_list_tests,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTestsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_tests, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTests(ctx, input)
				}
				var results []*svc.ListTestsOutput
				p := svc.NewListTestsPaginator(client, input)
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
		"list-unique-problems": {
			Name:   "list-unique-problems",
			Fields: fields_list_unique_problems,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListUniqueProblemsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_unique_problems, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListUniqueProblems(ctx, input)
				}
				var results []*svc.ListUniqueProblemsOutput
				p := svc.NewListUniqueProblemsPaginator(client, input)
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
		"list-uploads": {
			Name:   "list-uploads",
			Fields: fields_list_uploads,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListUploadsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_uploads, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListUploads(ctx, input)
				}
				var results []*svc.ListUploadsOutput
				p := svc.NewListUploadsPaginator(client, input)
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
		"list-vpce-configurations": {
			Name:   "list-vpce-configurations",
			Fields: fields_list_vpce_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVPCEConfigurationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_vpce_configurations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListVPCEConfigurations(ctx, input)
			},
		},
		"purchase-offering": {
			Name:   "purchase-offering",
			Fields: fields_purchase_offering,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PurchaseOfferingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_purchase_offering, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PurchaseOffering(ctx, input)
			},
		},
		"renew-offering": {
			Name:   "renew-offering",
			Fields: fields_renew_offering,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RenewOfferingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_renew_offering, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RenewOffering(ctx, input)
			},
		},
		"schedule-run": {
			Name:   "schedule-run",
			Fields: fields_schedule_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ScheduleRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_schedule_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ScheduleRun(ctx, input)
			},
		},
		"stop-job": {
			Name:   "stop-job",
			Fields: fields_stop_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopJob(ctx, input)
			},
		},
		"stop-remote-access-session": {
			Name:   "stop-remote-access-session",
			Fields: fields_stop_remote_access_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopRemoteAccessSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_remote_access_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopRemoteAccessSession(ctx, input)
			},
		},
		"stop-run": {
			Name:   "stop-run",
			Fields: fields_stop_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopRun(ctx, input)
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
		"update-device-instance": {
			Name:   "update-device-instance",
			Fields: fields_update_device_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDeviceInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_device_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDeviceInstance(ctx, input)
			},
		},
		"update-device-pool": {
			Name:   "update-device-pool",
			Fields: fields_update_device_pool,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDevicePoolInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_device_pool, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDevicePool(ctx, input)
			},
		},
		"update-instance-profile": {
			Name:   "update-instance-profile",
			Fields: fields_update_instance_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateInstanceProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_instance_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateInstanceProfile(ctx, input)
			},
		},
		"update-network-profile": {
			Name:   "update-network-profile",
			Fields: fields_update_network_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateNetworkProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_network_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateNetworkProfile(ctx, input)
			},
		},
		"update-project": {
			Name:   "update-project",
			Fields: fields_update_project,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateProjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_project, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateProject(ctx, input)
			},
		},
		"update-test-grid-project": {
			Name:   "update-test-grid-project",
			Fields: fields_update_test_grid_project,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTestGridProjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_test_grid_project, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTestGridProject(ctx, input)
			},
		},
		"update-upload": {
			Name:   "update-upload",
			Fields: fields_update_upload,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateUploadInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_upload, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateUpload(ctx, input)
			},
		},
		"update-vpce-configuration": {
			Name:   "update-vpce-configuration",
			Fields: fields_update_vpce_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateVPCEConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_vpce_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateVPCEConfiguration(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("devicefarm", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
