package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/emr"
)

var fields_add_instance_fleet = []leanruntime.Field{
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: true},
	{Name: "InstanceFleet", Flag: "instance-fleet", Type: "*types.InstanceFleetConfig", Required: true},
}

var fields_add_instance_groups = []leanruntime.Field{
	{Name: "InstanceGroups", Flag: "instance-groups", Type: "[]types.InstanceGroupConfig", Required: true},
	{Name: "JobFlowId", Flag: "job-flow-id", Type: "*string", Required: true},
}

var fields_add_job_flow_steps = []leanruntime.Field{
	{Name: "ExecutionRoleArn", Flag: "execution-role-arn", Type: "*string", Required: false},
	{Name: "JobFlowId", Flag: "job-flow-id", Type: "*string", Required: true},
	{Name: "Steps", Flag: "steps", Type: "[]types.StepConfig", Required: true},
}

var fields_add_tags = []leanruntime.Field{
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_cancel_steps = []leanruntime.Field{
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: true},
	{Name: "StepCancellationOption", Flag: "step-cancellation-option", Type: "types.StepCancellationOption", Required: false},
	{Name: "StepIds", Flag: "step-ids", Type: "[]string", Required: true},
}

var fields_create_persistent_app_ui = []leanruntime.Field{
	{Name: "EMRContainersConfig", Flag: "emr-containers-config", Type: "*types.EMRContainersConfig", Required: false},
	{Name: "ProfilerType", Flag: "profiler-type", Type: "types.ProfilerType", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TargetResourceArn", Flag: "target-resource-arn", Type: "*string", Required: true},
	{Name: "XReferer", Flag: "xreferer", Type: "*string", Required: false},
}

var fields_create_security_configuration = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "SecurityConfiguration", Flag: "security-configuration", Type: "*string", Required: true},
}

var fields_create_studio = []leanruntime.Field{
	{Name: "AuthMode", Flag: "auth-mode", Type: "types.AuthMode", Required: true},
	{Name: "DefaultS3Location", Flag: "default-s3-location", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EncryptionKeyArn", Flag: "encryption-key-arn", Type: "*string", Required: false},
	{Name: "EngineSecurityGroupId", Flag: "engine-security-group-id", Type: "*string", Required: true},
	{Name: "IdcInstanceArn", Flag: "idc-instance-arn", Type: "*string", Required: false},
	{Name: "IdcUserAssignment", Flag: "idc-user-assignment", Type: "types.IdcUserAssignment", Required: false},
	{Name: "IdpAuthUrl", Flag: "idp-auth-url", Type: "*string", Required: false},
	{Name: "IdpRelayStateParameterName", Flag: "idp-relay-state-parameter-name", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ServiceRole", Flag: "service-role", Type: "*string", Required: true},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TrustedIdentityPropagationEnabled", Flag: "trusted-identity-propagation-enabled", Type: "*bool", Required: false},
	{Name: "UserRole", Flag: "user-role", Type: "*string", Required: false},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: true},
	{Name: "WorkspaceSecurityGroupId", Flag: "workspace-security-group-id", Type: "*string", Required: true},
}

var fields_create_studio_session_mapping = []leanruntime.Field{
	{Name: "IdentityId", Flag: "identity-id", Type: "*string", Required: false},
	{Name: "IdentityName", Flag: "identity-name", Type: "*string", Required: false},
	{Name: "IdentityType", Flag: "identity-type", Type: "types.IdentityType", Required: true},
	{Name: "SessionPolicyArn", Flag: "session-policy-arn", Type: "*string", Required: true},
	{Name: "StudioId", Flag: "studio-id", Type: "*string", Required: true},
}

var fields_delete_security_configuration = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_studio = []leanruntime.Field{
	{Name: "StudioId", Flag: "studio-id", Type: "*string", Required: true},
}

var fields_delete_studio_session_mapping = []leanruntime.Field{
	{Name: "IdentityId", Flag: "identity-id", Type: "*string", Required: false},
	{Name: "IdentityName", Flag: "identity-name", Type: "*string", Required: false},
	{Name: "IdentityType", Flag: "identity-type", Type: "types.IdentityType", Required: true},
	{Name: "StudioId", Flag: "studio-id", Type: "*string", Required: true},
}

var fields_describe_cluster = []leanruntime.Field{
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: true},
}

var fields_describe_job_flows = []leanruntime.Field{
	{Name: "CreatedAfter", Flag: "created-after", Type: "*time.Time", Required: false},
	{Name: "CreatedBefore", Flag: "created-before", Type: "*time.Time", Required: false},
	{Name: "JobFlowIds", Flag: "job-flow-ids", Type: "[]string", Required: false},
	{Name: "JobFlowStates", Flag: "job-flow-states", Type: "[]types.JobFlowExecutionState", Required: false},
}

var fields_describe_notebook_execution = []leanruntime.Field{
	{Name: "NotebookExecutionId", Flag: "notebook-execution-id", Type: "*string", Required: true},
}

var fields_describe_persistent_app_ui = []leanruntime.Field{
	{Name: "PersistentAppUIId", Flag: "persistent-app-uiid", Type: "*string", Required: true},
}

var fields_describe_release_label = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ReleaseLabel", Flag: "release-label", Type: "*string", Required: false},
}

var fields_describe_security_configuration = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_describe_step = []leanruntime.Field{
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: true},
	{Name: "StepId", Flag: "step-id", Type: "*string", Required: true},
}

var fields_describe_studio = []leanruntime.Field{
	{Name: "StudioId", Flag: "studio-id", Type: "*string", Required: true},
}

var fields_get_auto_termination_policy = []leanruntime.Field{
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: true},
}

var fields_get_block_public_access_configuration = []leanruntime.Field{}

var fields_get_cluster_session_credentials = []leanruntime.Field{
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: true},
	{Name: "ExecutionRoleArn", Flag: "execution-role-arn", Type: "*string", Required: false},
}

var fields_get_managed_scaling_policy = []leanruntime.Field{
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: true},
}

var fields_get_on_cluster_app_ui_presigned_url = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: false},
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ExecutionRoleArn", Flag: "execution-role-arn", Type: "*string", Required: false},
	{Name: "OnClusterAppUIType", Flag: "on-cluster-app-ui-type", Type: "types.OnClusterAppUIType", Required: false},
}

var fields_get_persistent_app_ui_presigned_url = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: false},
	{Name: "AuthProxyCall", Flag: "auth-proxy-call", Type: "*bool", Required: false},
	{Name: "ExecutionRoleArn", Flag: "execution-role-arn", Type: "*string", Required: false},
	{Name: "PersistentAppUIId", Flag: "persistent-app-uiid", Type: "*string", Required: true},
	{Name: "PersistentAppUIType", Flag: "persistent-app-ui-type", Type: "types.PersistentAppUIType", Required: false},
}

var fields_get_studio_session_mapping = []leanruntime.Field{
	{Name: "IdentityId", Flag: "identity-id", Type: "*string", Required: false},
	{Name: "IdentityName", Flag: "identity-name", Type: "*string", Required: false},
	{Name: "IdentityType", Flag: "identity-type", Type: "types.IdentityType", Required: true},
	{Name: "StudioId", Flag: "studio-id", Type: "*string", Required: true},
}

var fields_list_bootstrap_actions = []leanruntime.Field{
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: true},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
}

var fields_list_clusters = []leanruntime.Field{
	{Name: "ClusterStates", Flag: "cluster-states", Type: "[]types.ClusterState", Required: false},
	{Name: "CreatedAfter", Flag: "created-after", Type: "*time.Time", Required: false},
	{Name: "CreatedBefore", Flag: "created-before", Type: "*time.Time", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
}

var fields_list_instance_fleets = []leanruntime.Field{
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: true},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
}

var fields_list_instance_groups = []leanruntime.Field{
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: true},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
}

var fields_list_instances = []leanruntime.Field{
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: true},
	{Name: "InstanceFleetId", Flag: "instance-fleet-id", Type: "*string", Required: false},
	{Name: "InstanceFleetType", Flag: "instance-fleet-type", Type: "types.InstanceFleetType", Required: false},
	{Name: "InstanceGroupId", Flag: "instance-group-id", Type: "*string", Required: false},
	{Name: "InstanceGroupTypes", Flag: "instance-group-types", Type: "[]types.InstanceGroupType", Required: false},
	{Name: "InstanceStates", Flag: "instance-states", Type: "[]types.InstanceState", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
}

var fields_list_notebook_executions = []leanruntime.Field{
	{Name: "EditorId", Flag: "editor-id", Type: "*string", Required: false},
	{Name: "ExecutionEngineId", Flag: "execution-engine-id", Type: "*string", Required: false},
	{Name: "From", Flag: "from", Type: "*time.Time", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.NotebookExecutionStatus", Required: false},
	{Name: "To", Flag: "to", Type: "*time.Time", Required: false},
}

var fields_list_release_labels = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "*types.ReleaseLabelFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_security_configurations = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
}

var fields_list_steps = []leanruntime.Field{
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: true},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "StepIds", Flag: "step-ids", Type: "[]string", Required: false},
	{Name: "StepStates", Flag: "step-states", Type: "[]types.StepState", Required: false},
}

var fields_list_studio_session_mappings = []leanruntime.Field{
	{Name: "IdentityType", Flag: "identity-type", Type: "types.IdentityType", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "StudioId", Flag: "studio-id", Type: "*string", Required: false},
}

var fields_list_studios = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
}

var fields_list_supported_instance_types = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "ReleaseLabel", Flag: "release-label", Type: "*string", Required: true},
}

var fields_modify_cluster = []leanruntime.Field{
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: true},
	{Name: "ExtendedSupport", Flag: "extended-support", Type: "*bool", Required: false},
	{Name: "StepConcurrencyLevel", Flag: "step-concurrency-level", Type: "*int32", Required: false},
}

var fields_modify_instance_fleet = []leanruntime.Field{
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: true},
	{Name: "InstanceFleet", Flag: "instance-fleet", Type: "*types.InstanceFleetModifyConfig", Required: true},
}

var fields_modify_instance_groups = []leanruntime.Field{
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: false},
	{Name: "InstanceGroups", Flag: "instance-groups", Type: "[]types.InstanceGroupModifyConfig", Required: false},
}

var fields_put_auto_scaling_policy = []leanruntime.Field{
	{Name: "AutoScalingPolicy", Flag: "auto-scaling-policy", Type: "*types.AutoScalingPolicy", Required: true},
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: true},
	{Name: "InstanceGroupId", Flag: "instance-group-id", Type: "*string", Required: true},
}

var fields_put_auto_termination_policy = []leanruntime.Field{
	{Name: "AutoTerminationPolicy", Flag: "auto-termination-policy", Type: "*types.AutoTerminationPolicy", Required: false},
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: true},
}

var fields_put_block_public_access_configuration = []leanruntime.Field{
	{Name: "BlockPublicAccessConfiguration", Flag: "block-public-access-configuration", Type: "*types.BlockPublicAccessConfiguration", Required: true},
}

var fields_put_managed_scaling_policy = []leanruntime.Field{
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: true},
	{Name: "ManagedScalingPolicy", Flag: "managed-scaling-policy", Type: "*types.ManagedScalingPolicy", Required: true},
}

var fields_remove_auto_scaling_policy = []leanruntime.Field{
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: true},
	{Name: "InstanceGroupId", Flag: "instance-group-id", Type: "*string", Required: true},
}

var fields_remove_auto_termination_policy = []leanruntime.Field{
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: true},
}

var fields_remove_managed_scaling_policy = []leanruntime.Field{
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: true},
}

var fields_remove_tags = []leanruntime.Field{
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_run_job_flow = []leanruntime.Field{
	{Name: "AdditionalInfo", Flag: "additional-info", Type: "*string", Required: false},
	{Name: "AmiVersion", Flag: "ami-version", Type: "*string", Required: false},
	{Name: "Applications", Flag: "applications", Type: "[]types.Application", Required: false},
	{Name: "AutoScalingRole", Flag: "auto-scaling-role", Type: "*string", Required: false},
	{Name: "AutoTerminationPolicy", Flag: "auto-termination-policy", Type: "*types.AutoTerminationPolicy", Required: false},
	{Name: "BootstrapActions", Flag: "bootstrap-actions", Type: "[]types.BootstrapActionConfig", Required: false},
	{Name: "Configurations", Flag: "configurations", Type: "[]types.Configuration", Required: false},
	{Name: "CustomAmiId", Flag: "custom-ami-id", Type: "*string", Required: false},
	{Name: "EbsRootVolumeIops", Flag: "ebs-root-volume-iops", Type: "*int32", Required: false},
	{Name: "EbsRootVolumeSize", Flag: "ebs-root-volume-size", Type: "*int32", Required: false},
	{Name: "EbsRootVolumeThroughput", Flag: "ebs-root-volume-throughput", Type: "*int32", Required: false},
	{Name: "ExtendedSupport", Flag: "extended-support", Type: "*bool", Required: false},
	{Name: "Instances", Flag: "instances", Type: "*types.JobFlowInstancesConfig", Required: true},
	{Name: "JobFlowRole", Flag: "job-flow-role", Type: "*string", Required: false},
	{Name: "KerberosAttributes", Flag: "kerberos-attributes", Type: "*types.KerberosAttributes", Required: false},
	{Name: "LogEncryptionKmsKeyId", Flag: "log-encryption-kms-key-id", Type: "*string", Required: false},
	{Name: "LogUri", Flag: "log-uri", Type: "*string", Required: false},
	{Name: "ManagedScalingPolicy", Flag: "managed-scaling-policy", Type: "*types.ManagedScalingPolicy", Required: false},
	{Name: "MonitoringConfiguration", Flag: "monitoring-configuration", Type: "*types.MonitoringConfiguration", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NewSupportedProducts", Flag: "new-supported-products", Type: "[]types.SupportedProductConfig", Required: false},
	{Name: "OSReleaseLabel", Flag: "os-release-label", Type: "*string", Required: false},
	{Name: "PlacementGroupConfigs", Flag: "placement-group-configs", Type: "[]types.PlacementGroupConfig", Required: false},
	{Name: "ReleaseLabel", Flag: "release-label", Type: "*string", Required: false},
	{Name: "RepoUpgradeOnBoot", Flag: "repo-upgrade-on-boot", Type: "types.RepoUpgradeOnBoot", Required: false},
	{Name: "ScaleDownBehavior", Flag: "scale-down-behavior", Type: "types.ScaleDownBehavior", Required: false},
	{Name: "SecurityConfiguration", Flag: "security-configuration", Type: "*string", Required: false},
	{Name: "ServiceRole", Flag: "service-role", Type: "*string", Required: false},
	{Name: "StepConcurrencyLevel", Flag: "step-concurrency-level", Type: "*int32", Required: false},
	{Name: "Steps", Flag: "steps", Type: "[]types.StepConfig", Required: false},
	{Name: "SupportedProducts", Flag: "supported-products", Type: "[]string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VisibleToAllUsers", Flag: "visible-to-all-users", Type: "*bool", Required: false},
}

var fields_set_keep_job_flow_alive_when_no_steps = []leanruntime.Field{
	{Name: "JobFlowIds", Flag: "job-flow-ids", Type: "[]string", Required: true},
	{Name: "KeepJobFlowAliveWhenNoSteps", Flag: "keep-job-flow-alive-when-no-steps", Type: "*bool", Required: true},
}

var fields_set_termination_protection = []leanruntime.Field{
	{Name: "JobFlowIds", Flag: "job-flow-ids", Type: "[]string", Required: true},
	{Name: "TerminationProtected", Flag: "termination-protected", Type: "*bool", Required: true},
}

var fields_set_unhealthy_node_replacement = []leanruntime.Field{
	{Name: "JobFlowIds", Flag: "job-flow-ids", Type: "[]string", Required: true},
	{Name: "UnhealthyNodeReplacement", Flag: "unhealthy-node-replacement", Type: "*bool", Required: true},
}

var fields_set_visible_to_all_users = []leanruntime.Field{
	{Name: "JobFlowIds", Flag: "job-flow-ids", Type: "[]string", Required: true},
	{Name: "VisibleToAllUsers", Flag: "visible-to-all-users", Type: "*bool", Required: true},
}

var fields_start_notebook_execution = []leanruntime.Field{
	{Name: "EditorId", Flag: "editor-id", Type: "*string", Required: false},
	{Name: "EnvironmentVariables", Flag: "environment-variables", Type: "map[string]string", Required: false},
	{Name: "ExecutionEngine", Flag: "execution-engine", Type: "*types.ExecutionEngineConfig", Required: true},
	{Name: "NotebookExecutionName", Flag: "notebook-execution-name", Type: "*string", Required: false},
	{Name: "NotebookInstanceSecurityGroupId", Flag: "notebook-instance-security-group-id", Type: "*string", Required: false},
	{Name: "NotebookParams", Flag: "notebook-params", Type: "*string", Required: false},
	{Name: "NotebookS3Location", Flag: "notebook-s3-location", Type: "*types.NotebookS3LocationFromInput", Required: false},
	{Name: "OutputNotebookFormat", Flag: "output-notebook-format", Type: "types.OutputNotebookFormat", Required: false},
	{Name: "OutputNotebookS3Location", Flag: "output-notebook-s3-location", Type: "*types.OutputNotebookS3LocationFromInput", Required: false},
	{Name: "RelativePath", Flag: "relative-path", Type: "*string", Required: false},
	{Name: "ServiceRole", Flag: "service-role", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_stop_notebook_execution = []leanruntime.Field{
	{Name: "NotebookExecutionId", Flag: "notebook-execution-id", Type: "*string", Required: true},
}

var fields_terminate_job_flows = []leanruntime.Field{
	{Name: "JobFlowIds", Flag: "job-flow-ids", Type: "[]string", Required: true},
}

var fields_update_studio = []leanruntime.Field{
	{Name: "DefaultS3Location", Flag: "default-s3-location", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EncryptionKeyArn", Flag: "encryption-key-arn", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "StudioId", Flag: "studio-id", Type: "*string", Required: true},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: false},
}

var fields_update_studio_session_mapping = []leanruntime.Field{
	{Name: "IdentityId", Flag: "identity-id", Type: "*string", Required: false},
	{Name: "IdentityName", Flag: "identity-name", Type: "*string", Required: false},
	{Name: "IdentityType", Flag: "identity-type", Type: "types.IdentityType", Required: true},
	{Name: "SessionPolicyArn", Flag: "session-policy-arn", Type: "*string", Required: true},
	{Name: "StudioId", Flag: "studio-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"add-instance-fleet": {
			Name:   "add-instance-fleet",
			Fields: fields_add_instance_fleet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddInstanceFleetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_instance_fleet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddInstanceFleet(ctx, input)
			},
		},
		"add-instance-groups": {
			Name:   "add-instance-groups",
			Fields: fields_add_instance_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddInstanceGroupsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_instance_groups, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddInstanceGroups(ctx, input)
			},
		},
		"add-job-flow-steps": {
			Name:   "add-job-flow-steps",
			Fields: fields_add_job_flow_steps,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddJobFlowStepsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_job_flow_steps, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddJobFlowSteps(ctx, input)
			},
		},
		"add-tags": {
			Name:   "add-tags",
			Fields: fields_add_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddTagsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_tags, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddTags(ctx, input)
			},
		},
		"cancel-steps": {
			Name:   "cancel-steps",
			Fields: fields_cancel_steps,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelStepsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_steps, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelSteps(ctx, input)
			},
		},
		"create-persistent-app-ui": {
			Name:   "create-persistent-app-ui",
			Fields: fields_create_persistent_app_ui,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePersistentAppUIInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_persistent_app_ui, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePersistentAppUI(ctx, input)
			},
		},
		"create-security-configuration": {
			Name:   "create-security-configuration",
			Fields: fields_create_security_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSecurityConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_security_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSecurityConfiguration(ctx, input)
			},
		},
		"create-studio": {
			Name:   "create-studio",
			Fields: fields_create_studio,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateStudioInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_studio, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateStudio(ctx, input)
			},
		},
		"create-studio-session-mapping": {
			Name:   "create-studio-session-mapping",
			Fields: fields_create_studio_session_mapping,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateStudioSessionMappingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_studio_session_mapping, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateStudioSessionMapping(ctx, input)
			},
		},
		"delete-security-configuration": {
			Name:   "delete-security-configuration",
			Fields: fields_delete_security_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSecurityConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_security_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSecurityConfiguration(ctx, input)
			},
		},
		"delete-studio": {
			Name:   "delete-studio",
			Fields: fields_delete_studio,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteStudioInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_studio, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteStudio(ctx, input)
			},
		},
		"delete-studio-session-mapping": {
			Name:   "delete-studio-session-mapping",
			Fields: fields_delete_studio_session_mapping,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteStudioSessionMappingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_studio_session_mapping, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteStudioSessionMapping(ctx, input)
			},
		},
		"describe-cluster": {
			Name:   "describe-cluster",
			Fields: fields_describe_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeCluster(ctx, input)
			},
		},
		"describe-job-flows": {
			Name:   "describe-job-flows",
			Fields: fields_describe_job_flows,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeJobFlowsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_job_flows, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeJobFlows(ctx, input)
			},
		},
		"describe-notebook-execution": {
			Name:   "describe-notebook-execution",
			Fields: fields_describe_notebook_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeNotebookExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_notebook_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeNotebookExecution(ctx, input)
			},
		},
		"describe-persistent-app-ui": {
			Name:   "describe-persistent-app-ui",
			Fields: fields_describe_persistent_app_ui,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePersistentAppUIInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_persistent_app_ui, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribePersistentAppUI(ctx, input)
			},
		},
		"describe-release-label": {
			Name:   "describe-release-label",
			Fields: fields_describe_release_label,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeReleaseLabelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_release_label, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeReleaseLabel(ctx, input)
			},
		},
		"describe-security-configuration": {
			Name:   "describe-security-configuration",
			Fields: fields_describe_security_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSecurityConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_security_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSecurityConfiguration(ctx, input)
			},
		},
		"describe-step": {
			Name:   "describe-step",
			Fields: fields_describe_step,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeStepInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_step, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeStep(ctx, input)
			},
		},
		"describe-studio": {
			Name:   "describe-studio",
			Fields: fields_describe_studio,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeStudioInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_studio, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeStudio(ctx, input)
			},
		},
		"get-auto-termination-policy": {
			Name:   "get-auto-termination-policy",
			Fields: fields_get_auto_termination_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAutoTerminationPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_auto_termination_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAutoTerminationPolicy(ctx, input)
			},
		},
		"get-block-public-access-configuration": {
			Name:   "get-block-public-access-configuration",
			Fields: fields_get_block_public_access_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBlockPublicAccessConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_block_public_access_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBlockPublicAccessConfiguration(ctx, input)
			},
		},
		"get-cluster-session-credentials": {
			Name:   "get-cluster-session-credentials",
			Fields: fields_get_cluster_session_credentials,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetClusterSessionCredentialsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_cluster_session_credentials, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetClusterSessionCredentials(ctx, input)
			},
		},
		"get-managed-scaling-policy": {
			Name:   "get-managed-scaling-policy",
			Fields: fields_get_managed_scaling_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetManagedScalingPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_managed_scaling_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetManagedScalingPolicy(ctx, input)
			},
		},
		"get-on-cluster-app-ui-presigned-url": {
			Name:   "get-on-cluster-app-ui-presigned-url",
			Fields: fields_get_on_cluster_app_ui_presigned_url,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOnClusterAppUIPresignedURLInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_on_cluster_app_ui_presigned_url, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetOnClusterAppUIPresignedURL(ctx, input)
			},
		},
		"get-persistent-app-ui-presigned-url": {
			Name:   "get-persistent-app-ui-presigned-url",
			Fields: fields_get_persistent_app_ui_presigned_url,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPersistentAppUIPresignedURLInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_persistent_app_ui_presigned_url, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPersistentAppUIPresignedURL(ctx, input)
			},
		},
		"get-studio-session-mapping": {
			Name:   "get-studio-session-mapping",
			Fields: fields_get_studio_session_mapping,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetStudioSessionMappingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_studio_session_mapping, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetStudioSessionMapping(ctx, input)
			},
		},
		"list-bootstrap-actions": {
			Name:   "list-bootstrap-actions",
			Fields: fields_list_bootstrap_actions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBootstrapActionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_bootstrap_actions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBootstrapActions(ctx, input)
				}
				var results []*svc.ListBootstrapActionsOutput
				p := svc.NewListBootstrapActionsPaginator(client, input)
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
		"list-clusters": {
			Name:   "list-clusters",
			Fields: fields_list_clusters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListClustersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_clusters, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListClusters(ctx, input)
				}
				var results []*svc.ListClustersOutput
				p := svc.NewListClustersPaginator(client, input)
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
		"list-instance-fleets": {
			Name:   "list-instance-fleets",
			Fields: fields_list_instance_fleets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInstanceFleetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_instance_fleets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListInstanceFleets(ctx, input)
				}
				var results []*svc.ListInstanceFleetsOutput
				p := svc.NewListInstanceFleetsPaginator(client, input)
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
		"list-instance-groups": {
			Name:   "list-instance-groups",
			Fields: fields_list_instance_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInstanceGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_instance_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListInstanceGroups(ctx, input)
				}
				var results []*svc.ListInstanceGroupsOutput
				p := svc.NewListInstanceGroupsPaginator(client, input)
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
		"list-notebook-executions": {
			Name:   "list-notebook-executions",
			Fields: fields_list_notebook_executions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListNotebookExecutionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_notebook_executions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListNotebookExecutions(ctx, input)
				}
				var results []*svc.ListNotebookExecutionsOutput
				p := svc.NewListNotebookExecutionsPaginator(client, input)
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
		"list-release-labels": {
			Name:   "list-release-labels",
			Fields: fields_list_release_labels,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListReleaseLabelsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_release_labels, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListReleaseLabels(ctx, input)
				}
				var results []*svc.ListReleaseLabelsOutput
				p := svc.NewListReleaseLabelsPaginator(client, input)
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
		"list-security-configurations": {
			Name:   "list-security-configurations",
			Fields: fields_list_security_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSecurityConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_security_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSecurityConfigurations(ctx, input)
				}
				var results []*svc.ListSecurityConfigurationsOutput
				p := svc.NewListSecurityConfigurationsPaginator(client, input)
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
		"list-steps": {
			Name:   "list-steps",
			Fields: fields_list_steps,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListStepsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_steps, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSteps(ctx, input)
				}
				var results []*svc.ListStepsOutput
				p := svc.NewListStepsPaginator(client, input)
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
		"list-studio-session-mappings": {
			Name:   "list-studio-session-mappings",
			Fields: fields_list_studio_session_mappings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListStudioSessionMappingsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_studio_session_mappings, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListStudioSessionMappings(ctx, input)
				}
				var results []*svc.ListStudioSessionMappingsOutput
				p := svc.NewListStudioSessionMappingsPaginator(client, input)
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
		"list-studios": {
			Name:   "list-studios",
			Fields: fields_list_studios,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListStudiosInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_studios, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListStudios(ctx, input)
				}
				var results []*svc.ListStudiosOutput
				p := svc.NewListStudiosPaginator(client, input)
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
		"list-supported-instance-types": {
			Name:   "list-supported-instance-types",
			Fields: fields_list_supported_instance_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSupportedInstanceTypesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_supported_instance_types, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSupportedInstanceTypes(ctx, input)
				}
				var results []*svc.ListSupportedInstanceTypesOutput
				p := svc.NewListSupportedInstanceTypesPaginator(client, input)
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
		"modify-cluster": {
			Name:   "modify-cluster",
			Fields: fields_modify_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyCluster(ctx, input)
			},
		},
		"modify-instance-fleet": {
			Name:   "modify-instance-fleet",
			Fields: fields_modify_instance_fleet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyInstanceFleetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_instance_fleet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyInstanceFleet(ctx, input)
			},
		},
		"modify-instance-groups": {
			Name:   "modify-instance-groups",
			Fields: fields_modify_instance_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyInstanceGroupsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_instance_groups, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyInstanceGroups(ctx, input)
			},
		},
		"put-auto-scaling-policy": {
			Name:   "put-auto-scaling-policy",
			Fields: fields_put_auto_scaling_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutAutoScalingPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_auto_scaling_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutAutoScalingPolicy(ctx, input)
			},
		},
		"put-auto-termination-policy": {
			Name:   "put-auto-termination-policy",
			Fields: fields_put_auto_termination_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutAutoTerminationPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_auto_termination_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutAutoTerminationPolicy(ctx, input)
			},
		},
		"put-block-public-access-configuration": {
			Name:   "put-block-public-access-configuration",
			Fields: fields_put_block_public_access_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutBlockPublicAccessConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_block_public_access_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutBlockPublicAccessConfiguration(ctx, input)
			},
		},
		"put-managed-scaling-policy": {
			Name:   "put-managed-scaling-policy",
			Fields: fields_put_managed_scaling_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutManagedScalingPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_managed_scaling_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutManagedScalingPolicy(ctx, input)
			},
		},
		"remove-auto-scaling-policy": {
			Name:   "remove-auto-scaling-policy",
			Fields: fields_remove_auto_scaling_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveAutoScalingPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_auto_scaling_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveAutoScalingPolicy(ctx, input)
			},
		},
		"remove-auto-termination-policy": {
			Name:   "remove-auto-termination-policy",
			Fields: fields_remove_auto_termination_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveAutoTerminationPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_auto_termination_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveAutoTerminationPolicy(ctx, input)
			},
		},
		"remove-managed-scaling-policy": {
			Name:   "remove-managed-scaling-policy",
			Fields: fields_remove_managed_scaling_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveManagedScalingPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_managed_scaling_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveManagedScalingPolicy(ctx, input)
			},
		},
		"remove-tags": {
			Name:   "remove-tags",
			Fields: fields_remove_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveTagsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_tags, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveTags(ctx, input)
			},
		},
		"run-job-flow": {
			Name:   "run-job-flow",
			Fields: fields_run_job_flow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RunJobFlowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_run_job_flow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RunJobFlow(ctx, input)
			},
		},
		"set-keep-job-flow-alive-when-no-steps": {
			Name:   "set-keep-job-flow-alive-when-no-steps",
			Fields: fields_set_keep_job_flow_alive_when_no_steps,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetKeepJobFlowAliveWhenNoStepsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_keep_job_flow_alive_when_no_steps, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetKeepJobFlowAliveWhenNoSteps(ctx, input)
			},
		},
		"set-termination-protection": {
			Name:   "set-termination-protection",
			Fields: fields_set_termination_protection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetTerminationProtectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_termination_protection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetTerminationProtection(ctx, input)
			},
		},
		"set-unhealthy-node-replacement": {
			Name:   "set-unhealthy-node-replacement",
			Fields: fields_set_unhealthy_node_replacement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetUnhealthyNodeReplacementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_unhealthy_node_replacement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetUnhealthyNodeReplacement(ctx, input)
			},
		},
		"set-visible-to-all-users": {
			Name:   "set-visible-to-all-users",
			Fields: fields_set_visible_to_all_users,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetVisibleToAllUsersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_visible_to_all_users, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetVisibleToAllUsers(ctx, input)
			},
		},
		"start-notebook-execution": {
			Name:   "start-notebook-execution",
			Fields: fields_start_notebook_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartNotebookExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_notebook_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartNotebookExecution(ctx, input)
			},
		},
		"stop-notebook-execution": {
			Name:   "stop-notebook-execution",
			Fields: fields_stop_notebook_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopNotebookExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_notebook_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopNotebookExecution(ctx, input)
			},
		},
		"terminate-job-flows": {
			Name:   "terminate-job-flows",
			Fields: fields_terminate_job_flows,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TerminateJobFlowsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_terminate_job_flows, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TerminateJobFlows(ctx, input)
			},
		},
		"update-studio": {
			Name:   "update-studio",
			Fields: fields_update_studio,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateStudioInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_studio, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateStudio(ctx, input)
			},
		},
		"update-studio-session-mapping": {
			Name:   "update-studio-session-mapping",
			Fields: fields_update_studio_session_mapping,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateStudioSessionMappingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_studio_session_mapping, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateStudioSessionMapping(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("emr", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
