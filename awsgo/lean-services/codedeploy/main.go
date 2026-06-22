package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/codedeploy"
)

var fields_add_tags_to_on_premises_instances = []leanruntime.Field{
	{Name: "InstanceNames", Flag: "instance-names", Type: "[]string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_batch_get_application_revisions = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "Revisions", Flag: "revisions", Type: "[]types.RevisionLocation", Required: true},
}

var fields_batch_get_applications = []leanruntime.Field{
	{Name: "ApplicationNames", Flag: "application-names", Type: "[]string", Required: true},
}

var fields_batch_get_deployment_groups = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "DeploymentGroupNames", Flag: "deployment-group-names", Type: "[]string", Required: true},
}

var fields_batch_get_deployment_instances = []leanruntime.Field{
	{Name: "DeploymentId", Flag: "deployment-id", Type: "*string", Required: true},
	{Name: "InstanceIds", Flag: "instance-ids", Type: "[]string", Required: true},
}

var fields_batch_get_deployment_targets = []leanruntime.Field{
	{Name: "DeploymentId", Flag: "deployment-id", Type: "*string", Required: true},
	{Name: "TargetIds", Flag: "target-ids", Type: "[]string", Required: true},
}

var fields_batch_get_deployments = []leanruntime.Field{
	{Name: "DeploymentIds", Flag: "deployment-ids", Type: "[]string", Required: true},
}

var fields_batch_get_on_premises_instances = []leanruntime.Field{
	{Name: "InstanceNames", Flag: "instance-names", Type: "[]string", Required: true},
}

var fields_continue_deployment = []leanruntime.Field{
	{Name: "DeploymentId", Flag: "deployment-id", Type: "*string", Required: false},
	{Name: "DeploymentWaitType", Flag: "deployment-wait-type", Type: "types.DeploymentWaitType", Required: false},
}

var fields_create_application = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "ComputePlatform", Flag: "compute-platform", Type: "types.ComputePlatform", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_deployment = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "AutoRollbackConfiguration", Flag: "auto-rollback-configuration", Type: "*types.AutoRollbackConfiguration", Required: false},
	{Name: "DeploymentConfigName", Flag: "deployment-config-name", Type: "*string", Required: false},
	{Name: "DeploymentGroupName", Flag: "deployment-group-name", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "FileExistsBehavior", Flag: "file-exists-behavior", Type: "types.FileExistsBehavior", Required: false},
	{Name: "IgnoreApplicationStopFailures", Flag: "ignore-application-stop-failures", Type: "bool", Required: false},
	{Name: "OverrideAlarmConfiguration", Flag: "override-alarm-configuration", Type: "*types.AlarmConfiguration", Required: false},
	{Name: "Revision", Flag: "revision", Type: "*types.RevisionLocation", Required: false},
	{Name: "TargetInstances", Flag: "target-instances", Type: "*types.TargetInstances", Required: false},
	{Name: "UpdateOutdatedInstancesOnly", Flag: "update-outdated-instances-only", Type: "bool", Required: false},
}

var fields_create_deployment_config = []leanruntime.Field{
	{Name: "ComputePlatform", Flag: "compute-platform", Type: "types.ComputePlatform", Required: false},
	{Name: "DeploymentConfigName", Flag: "deployment-config-name", Type: "*string", Required: true},
	{Name: "MinimumHealthyHosts", Flag: "minimum-healthy-hosts", Type: "*types.MinimumHealthyHosts", Required: false},
	{Name: "TrafficRoutingConfig", Flag: "traffic-routing-config", Type: "*types.TrafficRoutingConfig", Required: false},
	{Name: "ZonalConfig", Flag: "zonal-config", Type: "*types.ZonalConfig", Required: false},
}

var fields_create_deployment_group = []leanruntime.Field{
	{Name: "AlarmConfiguration", Flag: "alarm-configuration", Type: "*types.AlarmConfiguration", Required: false},
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "AutoRollbackConfiguration", Flag: "auto-rollback-configuration", Type: "*types.AutoRollbackConfiguration", Required: false},
	{Name: "AutoScalingGroups", Flag: "auto-scaling-groups", Type: "[]string", Required: false},
	{Name: "BlueGreenDeploymentConfiguration", Flag: "blue-green-deployment-configuration", Type: "*types.BlueGreenDeploymentConfiguration", Required: false},
	{Name: "DeploymentConfigName", Flag: "deployment-config-name", Type: "*string", Required: false},
	{Name: "DeploymentGroupName", Flag: "deployment-group-name", Type: "*string", Required: true},
	{Name: "DeploymentStyle", Flag: "deployment-style", Type: "*types.DeploymentStyle", Required: false},
	{Name: "Ec2TagFilters", Flag: "ec2-tag-filters", Type: "[]types.EC2TagFilter", Required: false},
	{Name: "Ec2TagSet", Flag: "ec2-tag-set", Type: "*types.EC2TagSet", Required: false},
	{Name: "EcsServices", Flag: "ecs-services", Type: "[]types.ECSService", Required: false},
	{Name: "LoadBalancerInfo", Flag: "load-balancer-info", Type: "*types.LoadBalancerInfo", Required: false},
	{Name: "OnPremisesInstanceTagFilters", Flag: "on-premises-instance-tag-filters", Type: "[]types.TagFilter", Required: false},
	{Name: "OnPremisesTagSet", Flag: "on-premises-tag-set", Type: "*types.OnPremisesTagSet", Required: false},
	{Name: "OutdatedInstancesStrategy", Flag: "outdated-instances-strategy", Type: "types.OutdatedInstancesStrategy", Required: false},
	{Name: "ServiceRoleArn", Flag: "service-role-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TerminationHookEnabled", Flag: "termination-hook-enabled", Type: "*bool", Required: false},
	{Name: "TriggerConfigurations", Flag: "trigger-configurations", Type: "[]types.TriggerConfig", Required: false},
}

var fields_delete_application = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
}

var fields_delete_deployment_config = []leanruntime.Field{
	{Name: "DeploymentConfigName", Flag: "deployment-config-name", Type: "*string", Required: true},
}

var fields_delete_deployment_group = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "DeploymentGroupName", Flag: "deployment-group-name", Type: "*string", Required: true},
}

var fields_delete_git_hub_account_token = []leanruntime.Field{
	{Name: "TokenName", Flag: "token-name", Type: "*string", Required: false},
}

var fields_delete_resources_by_external_id = []leanruntime.Field{
	{Name: "ExternalId", Flag: "external-id", Type: "*string", Required: false},
}

var fields_deregister_on_premises_instance = []leanruntime.Field{
	{Name: "InstanceName", Flag: "instance-name", Type: "*string", Required: true},
}

var fields_get_application = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
}

var fields_get_application_revision = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "Revision", Flag: "revision", Type: "*types.RevisionLocation", Required: true},
}

var fields_get_deployment = []leanruntime.Field{
	{Name: "DeploymentId", Flag: "deployment-id", Type: "*string", Required: true},
}

var fields_get_deployment_config = []leanruntime.Field{
	{Name: "DeploymentConfigName", Flag: "deployment-config-name", Type: "*string", Required: true},
}

var fields_get_deployment_group = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "DeploymentGroupName", Flag: "deployment-group-name", Type: "*string", Required: true},
}

var fields_get_deployment_instance = []leanruntime.Field{
	{Name: "DeploymentId", Flag: "deployment-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_get_deployment_target = []leanruntime.Field{
	{Name: "DeploymentId", Flag: "deployment-id", Type: "*string", Required: true},
	{Name: "TargetId", Flag: "target-id", Type: "*string", Required: true},
}

var fields_get_on_premises_instance = []leanruntime.Field{
	{Name: "InstanceName", Flag: "instance-name", Type: "*string", Required: true},
}

var fields_list_application_revisions = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "Deployed", Flag: "deployed", Type: "types.ListStateFilterAction", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "S3Bucket", Flag: "s3-bucket", Type: "*string", Required: false},
	{Name: "S3KeyPrefix", Flag: "s3-key-prefix", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.ApplicationRevisionSortBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_list_applications = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_deployment_configs = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_deployment_groups = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_deployment_instances = []leanruntime.Field{
	{Name: "DeploymentId", Flag: "deployment-id", Type: "*string", Required: true},
	{Name: "InstanceStatusFilter", Flag: "instance-status-filter", Type: "[]types.InstanceStatus", Required: false},
	{Name: "InstanceTypeFilter", Flag: "instance-type-filter", Type: "[]types.InstanceType", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_deployment_targets = []leanruntime.Field{
	{Name: "DeploymentId", Flag: "deployment-id", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TargetFilters", Flag: "target-filters", Type: "map[string][]string", Required: false},
}

var fields_list_deployments = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: false},
	{Name: "CreateTimeRange", Flag: "create-time-range", Type: "*types.TimeRange", Required: false},
	{Name: "DeploymentGroupName", Flag: "deployment-group-name", Type: "*string", Required: false},
	{Name: "ExternalId", Flag: "external-id", Type: "*string", Required: false},
	{Name: "IncludeOnlyStatuses", Flag: "include-only-statuses", Type: "[]types.DeploymentStatus", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_git_hub_account_token_names = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_on_premises_instances = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RegistrationStatus", Flag: "registration-status", Type: "types.RegistrationStatus", Required: false},
	{Name: "TagFilters", Flag: "tag-filters", Type: "[]types.TagFilter", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_put_lifecycle_event_hook_execution_status = []leanruntime.Field{
	{Name: "DeploymentId", Flag: "deployment-id", Type: "*string", Required: false},
	{Name: "LifecycleEventHookExecutionId", Flag: "lifecycle-event-hook-execution-id", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.LifecycleEventStatus", Required: false},
}

var fields_register_application_revision = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Revision", Flag: "revision", Type: "*types.RevisionLocation", Required: true},
}

var fields_register_on_premises_instance = []leanruntime.Field{
	{Name: "IamSessionArn", Flag: "iam-session-arn", Type: "*string", Required: false},
	{Name: "IamUserArn", Flag: "iam-user-arn", Type: "*string", Required: false},
	{Name: "InstanceName", Flag: "instance-name", Type: "*string", Required: true},
}

var fields_remove_tags_from_on_premises_instances = []leanruntime.Field{
	{Name: "InstanceNames", Flag: "instance-names", Type: "[]string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_skip_wait_time_for_instance_termination = []leanruntime.Field{
	{Name: "DeploymentId", Flag: "deployment-id", Type: "*string", Required: false},
}

var fields_stop_deployment = []leanruntime.Field{
	{Name: "AutoRollbackEnabled", Flag: "auto-rollback-enabled", Type: "*bool", Required: false},
	{Name: "DeploymentId", Flag: "deployment-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_application = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: false},
	{Name: "NewApplicationName", Flag: "new-application-name", Type: "*string", Required: false},
}

var fields_update_deployment_group = []leanruntime.Field{
	{Name: "AlarmConfiguration", Flag: "alarm-configuration", Type: "*types.AlarmConfiguration", Required: false},
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "AutoRollbackConfiguration", Flag: "auto-rollback-configuration", Type: "*types.AutoRollbackConfiguration", Required: false},
	{Name: "AutoScalingGroups", Flag: "auto-scaling-groups", Type: "[]string", Required: false},
	{Name: "BlueGreenDeploymentConfiguration", Flag: "blue-green-deployment-configuration", Type: "*types.BlueGreenDeploymentConfiguration", Required: false},
	{Name: "CurrentDeploymentGroupName", Flag: "current-deployment-group-name", Type: "*string", Required: true},
	{Name: "DeploymentConfigName", Flag: "deployment-config-name", Type: "*string", Required: false},
	{Name: "DeploymentStyle", Flag: "deployment-style", Type: "*types.DeploymentStyle", Required: false},
	{Name: "Ec2TagFilters", Flag: "ec2-tag-filters", Type: "[]types.EC2TagFilter", Required: false},
	{Name: "Ec2TagSet", Flag: "ec2-tag-set", Type: "*types.EC2TagSet", Required: false},
	{Name: "EcsServices", Flag: "ecs-services", Type: "[]types.ECSService", Required: false},
	{Name: "LoadBalancerInfo", Flag: "load-balancer-info", Type: "*types.LoadBalancerInfo", Required: false},
	{Name: "NewDeploymentGroupName", Flag: "new-deployment-group-name", Type: "*string", Required: false},
	{Name: "OnPremisesInstanceTagFilters", Flag: "on-premises-instance-tag-filters", Type: "[]types.TagFilter", Required: false},
	{Name: "OnPremisesTagSet", Flag: "on-premises-tag-set", Type: "*types.OnPremisesTagSet", Required: false},
	{Name: "OutdatedInstancesStrategy", Flag: "outdated-instances-strategy", Type: "types.OutdatedInstancesStrategy", Required: false},
	{Name: "ServiceRoleArn", Flag: "service-role-arn", Type: "*string", Required: false},
	{Name: "TerminationHookEnabled", Flag: "termination-hook-enabled", Type: "*bool", Required: false},
	{Name: "TriggerConfigurations", Flag: "trigger-configurations", Type: "[]types.TriggerConfig", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"add-tags-to-on-premises-instances": {
			Name:   "add-tags-to-on-premises-instances",
			Fields: fields_add_tags_to_on_premises_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddTagsToOnPremisesInstancesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_tags_to_on_premises_instances, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddTagsToOnPremisesInstances(ctx, input)
			},
		},
		"batch-get-application-revisions": {
			Name:   "batch-get-application-revisions",
			Fields: fields_batch_get_application_revisions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetApplicationRevisionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_application_revisions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetApplicationRevisions(ctx, input)
			},
		},
		"batch-get-applications": {
			Name:   "batch-get-applications",
			Fields: fields_batch_get_applications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetApplicationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_applications, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetApplications(ctx, input)
			},
		},
		"batch-get-deployment-groups": {
			Name:   "batch-get-deployment-groups",
			Fields: fields_batch_get_deployment_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetDeploymentGroupsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_deployment_groups, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetDeploymentGroups(ctx, input)
			},
		},
		"batch-get-deployment-instances": {
			Name:   "batch-get-deployment-instances",
			Fields: fields_batch_get_deployment_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetDeploymentInstancesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_deployment_instances, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetDeploymentInstances(ctx, input)
			},
		},
		"batch-get-deployment-targets": {
			Name:   "batch-get-deployment-targets",
			Fields: fields_batch_get_deployment_targets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetDeploymentTargetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_deployment_targets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetDeploymentTargets(ctx, input)
			},
		},
		"batch-get-deployments": {
			Name:   "batch-get-deployments",
			Fields: fields_batch_get_deployments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetDeploymentsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_deployments, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetDeployments(ctx, input)
			},
		},
		"batch-get-on-premises-instances": {
			Name:   "batch-get-on-premises-instances",
			Fields: fields_batch_get_on_premises_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetOnPremisesInstancesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_on_premises_instances, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetOnPremisesInstances(ctx, input)
			},
		},
		"continue-deployment": {
			Name:   "continue-deployment",
			Fields: fields_continue_deployment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ContinueDeploymentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_continue_deployment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ContinueDeployment(ctx, input)
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
		"create-deployment": {
			Name:   "create-deployment",
			Fields: fields_create_deployment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDeploymentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_deployment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDeployment(ctx, input)
			},
		},
		"create-deployment-config": {
			Name:   "create-deployment-config",
			Fields: fields_create_deployment_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDeploymentConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_deployment_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDeploymentConfig(ctx, input)
			},
		},
		"create-deployment-group": {
			Name:   "create-deployment-group",
			Fields: fields_create_deployment_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDeploymentGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_deployment_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDeploymentGroup(ctx, input)
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
		"delete-deployment-config": {
			Name:   "delete-deployment-config",
			Fields: fields_delete_deployment_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDeploymentConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_deployment_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDeploymentConfig(ctx, input)
			},
		},
		"delete-deployment-group": {
			Name:   "delete-deployment-group",
			Fields: fields_delete_deployment_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDeploymentGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_deployment_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDeploymentGroup(ctx, input)
			},
		},
		"delete-git-hub-account-token": {
			Name:   "delete-git-hub-account-token",
			Fields: fields_delete_git_hub_account_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteGitHubAccountTokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_git_hub_account_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteGitHubAccountToken(ctx, input)
			},
		},
		"delete-resources-by-external-id": {
			Name:   "delete-resources-by-external-id",
			Fields: fields_delete_resources_by_external_id,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteResourcesByExternalIdInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_resources_by_external_id, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteResourcesByExternalId(ctx, input)
			},
		},
		"deregister-on-premises-instance": {
			Name:   "deregister-on-premises-instance",
			Fields: fields_deregister_on_premises_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterOnPremisesInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_on_premises_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterOnPremisesInstance(ctx, input)
			},
		},
		"get-application": {
			Name:   "get-application",
			Fields: fields_get_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetApplication(ctx, input)
			},
		},
		"get-application-revision": {
			Name:   "get-application-revision",
			Fields: fields_get_application_revision,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetApplicationRevisionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_application_revision, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetApplicationRevision(ctx, input)
			},
		},
		"get-deployment": {
			Name:   "get-deployment",
			Fields: fields_get_deployment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDeploymentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_deployment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDeployment(ctx, input)
			},
		},
		"get-deployment-config": {
			Name:   "get-deployment-config",
			Fields: fields_get_deployment_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDeploymentConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_deployment_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDeploymentConfig(ctx, input)
			},
		},
		"get-deployment-group": {
			Name:   "get-deployment-group",
			Fields: fields_get_deployment_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDeploymentGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_deployment_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDeploymentGroup(ctx, input)
			},
		},
		"get-deployment-instance": {
			Name:   "get-deployment-instance",
			Fields: fields_get_deployment_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDeploymentInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_deployment_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDeploymentInstance(ctx, input)
			},
		},
		"get-deployment-target": {
			Name:   "get-deployment-target",
			Fields: fields_get_deployment_target,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDeploymentTargetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_deployment_target, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDeploymentTarget(ctx, input)
			},
		},
		"get-on-premises-instance": {
			Name:   "get-on-premises-instance",
			Fields: fields_get_on_premises_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOnPremisesInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_on_premises_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetOnPremisesInstance(ctx, input)
			},
		},
		"list-application-revisions": {
			Name:   "list-application-revisions",
			Fields: fields_list_application_revisions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListApplicationRevisionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_application_revisions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListApplicationRevisions(ctx, input)
				}
				var results []*svc.ListApplicationRevisionsOutput
				p := svc.NewListApplicationRevisionsPaginator(client, input)
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
		"list-deployment-configs": {
			Name:   "list-deployment-configs",
			Fields: fields_list_deployment_configs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDeploymentConfigsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_deployment_configs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDeploymentConfigs(ctx, input)
				}
				var results []*svc.ListDeploymentConfigsOutput
				p := svc.NewListDeploymentConfigsPaginator(client, input)
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
		"list-deployment-groups": {
			Name:   "list-deployment-groups",
			Fields: fields_list_deployment_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDeploymentGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_deployment_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDeploymentGroups(ctx, input)
				}
				var results []*svc.ListDeploymentGroupsOutput
				p := svc.NewListDeploymentGroupsPaginator(client, input)
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
		"list-deployment-instances": {
			Name:   "list-deployment-instances",
			Fields: fields_list_deployment_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDeploymentInstancesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_deployment_instances, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDeploymentInstances(ctx, input)
				}
				var results []*svc.ListDeploymentInstancesOutput
				p := svc.NewListDeploymentInstancesPaginator(client, input)
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
		"list-deployment-targets": {
			Name:   "list-deployment-targets",
			Fields: fields_list_deployment_targets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDeploymentTargetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_deployment_targets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListDeploymentTargets(ctx, input)
			},
		},
		"list-deployments": {
			Name:   "list-deployments",
			Fields: fields_list_deployments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDeploymentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_deployments, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDeployments(ctx, input)
				}
				var results []*svc.ListDeploymentsOutput
				p := svc.NewListDeploymentsPaginator(client, input)
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
		"list-git-hub-account-token-names": {
			Name:   "list-git-hub-account-token-names",
			Fields: fields_list_git_hub_account_token_names,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGitHubAccountTokenNamesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_git_hub_account_token_names, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListGitHubAccountTokenNames(ctx, input)
			},
		},
		"list-on-premises-instances": {
			Name:   "list-on-premises-instances",
			Fields: fields_list_on_premises_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOnPremisesInstancesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_on_premises_instances, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListOnPremisesInstances(ctx, input)
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
		"put-lifecycle-event-hook-execution-status": {
			Name:   "put-lifecycle-event-hook-execution-status",
			Fields: fields_put_lifecycle_event_hook_execution_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutLifecycleEventHookExecutionStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_lifecycle_event_hook_execution_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutLifecycleEventHookExecutionStatus(ctx, input)
			},
		},
		"register-application-revision": {
			Name:   "register-application-revision",
			Fields: fields_register_application_revision,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterApplicationRevisionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_application_revision, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterApplicationRevision(ctx, input)
			},
		},
		"register-on-premises-instance": {
			Name:   "register-on-premises-instance",
			Fields: fields_register_on_premises_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterOnPremisesInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_on_premises_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterOnPremisesInstance(ctx, input)
			},
		},
		"remove-tags-from-on-premises-instances": {
			Name:   "remove-tags-from-on-premises-instances",
			Fields: fields_remove_tags_from_on_premises_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveTagsFromOnPremisesInstancesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_tags_from_on_premises_instances, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveTagsFromOnPremisesInstances(ctx, input)
			},
		},
		"skip-wait-time-for-instance-termination": {
			Name:   "skip-wait-time-for-instance-termination",
			Fields: fields_skip_wait_time_for_instance_termination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SkipWaitTimeForInstanceTerminationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_skip_wait_time_for_instance_termination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SkipWaitTimeForInstanceTermination(ctx, input)
			},
		},
		"stop-deployment": {
			Name:   "stop-deployment",
			Fields: fields_stop_deployment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopDeploymentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_deployment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopDeployment(ctx, input)
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
		"update-deployment-group": {
			Name:   "update-deployment-group",
			Fields: fields_update_deployment_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDeploymentGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_deployment_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDeploymentGroup(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("codedeploy", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
