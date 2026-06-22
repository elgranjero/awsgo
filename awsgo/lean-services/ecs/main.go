package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/ecs"
)

var fields_create_capacity_provider = []leanruntime.Field{
	{Name: "AutoScalingGroupProvider", Flag: "auto-scaling-group-provider", Type: "*types.AutoScalingGroupProvider", Required: false},
	{Name: "Cluster", Flag: "cluster", Type: "*string", Required: false},
	{Name: "ManagedInstancesProvider", Flag: "managed-instances-provider", Type: "*types.CreateManagedInstancesProviderConfiguration", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_cluster = []leanruntime.Field{
	{Name: "CapacityProviders", Flag: "capacity-providers", Type: "[]string", Required: false},
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: false},
	{Name: "Configuration", Flag: "configuration", Type: "*types.ClusterConfiguration", Required: false},
	{Name: "DefaultCapacityProviderStrategy", Flag: "default-capacity-provider-strategy", Type: "[]types.CapacityProviderStrategyItem", Required: false},
	{Name: "ServiceConnectDefaults", Flag: "service-connect-defaults", Type: "*types.ClusterServiceConnectDefaultsRequest", Required: false},
	{Name: "Settings", Flag: "settings", Type: "[]types.ClusterSetting", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_express_gateway_service = []leanruntime.Field{
	{Name: "Cluster", Flag: "cluster", Type: "*string", Required: false},
	{Name: "Cpu", Flag: "cpu", Type: "*string", Required: false},
	{Name: "ExecutionRoleArn", Flag: "execution-role-arn", Type: "*string", Required: true},
	{Name: "HealthCheckPath", Flag: "health-check-path", Type: "*string", Required: false},
	{Name: "InfrastructureRoleArn", Flag: "infrastructure-role-arn", Type: "*string", Required: true},
	{Name: "Memory", Flag: "memory", Type: "*string", Required: false},
	{Name: "NetworkConfiguration", Flag: "network-configuration", Type: "*types.ExpressGatewayServiceNetworkConfiguration", Required: false},
	{Name: "PrimaryContainer", Flag: "primary-container", Type: "*types.ExpressGatewayContainer", Required: true},
	{Name: "ScalingTarget", Flag: "scaling-target", Type: "*types.ExpressGatewayScalingTarget", Required: false},
	{Name: "ServiceName", Flag: "service-name", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TaskRoleArn", Flag: "task-role-arn", Type: "*string", Required: false},
}

var fields_create_service = []leanruntime.Field{
	{Name: "AvailabilityZoneRebalancing", Flag: "availability-zone-rebalancing", Type: "types.AvailabilityZoneRebalancing", Required: false},
	{Name: "CapacityProviderStrategy", Flag: "capacity-provider-strategy", Type: "[]types.CapacityProviderStrategyItem", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Cluster", Flag: "cluster", Type: "*string", Required: false},
	{Name: "DeploymentConfiguration", Flag: "deployment-configuration", Type: "*types.DeploymentConfiguration", Required: false},
	{Name: "DeploymentController", Flag: "deployment-controller", Type: "*types.DeploymentController", Required: false},
	{Name: "DesiredCount", Flag: "desired-count", Type: "*int32", Required: false},
	{Name: "EnableECSManagedTags", Flag: "enable-ecs-managed-tags", Type: "bool", Required: false},
	{Name: "EnableExecuteCommand", Flag: "enable-execute-command", Type: "bool", Required: false},
	{Name: "HealthCheckGracePeriodSeconds", Flag: "health-check-grace-period-seconds", Type: "*int32", Required: false},
	{Name: "LaunchType", Flag: "launch-type", Type: "types.LaunchType", Required: false},
	{Name: "LoadBalancers", Flag: "load-balancers", Type: "[]types.LoadBalancer", Required: false},
	{Name: "NetworkConfiguration", Flag: "network-configuration", Type: "*types.NetworkConfiguration", Required: false},
	{Name: "PlacementConstraints", Flag: "placement-constraints", Type: "[]types.PlacementConstraint", Required: false},
	{Name: "PlacementStrategy", Flag: "placement-strategy", Type: "[]types.PlacementStrategy", Required: false},
	{Name: "PlatformVersion", Flag: "platform-version", Type: "*string", Required: false},
	{Name: "PropagateTags", Flag: "propagate-tags", Type: "types.PropagateTags", Required: false},
	{Name: "Role", Flag: "role", Type: "*string", Required: false},
	{Name: "SchedulingStrategy", Flag: "scheduling-strategy", Type: "types.SchedulingStrategy", Required: false},
	{Name: "ServiceConnectConfiguration", Flag: "service-connect-configuration", Type: "*types.ServiceConnectConfiguration", Required: false},
	{Name: "ServiceName", Flag: "service-name", Type: "*string", Required: true},
	{Name: "ServiceRegistries", Flag: "service-registries", Type: "[]types.ServiceRegistry", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TaskDefinition", Flag: "task-definition", Type: "*string", Required: false},
	{Name: "VolumeConfigurations", Flag: "volume-configurations", Type: "[]types.ServiceVolumeConfiguration", Required: false},
	{Name: "VpcLatticeConfigurations", Flag: "vpc-lattice-configurations", Type: "[]types.VpcLatticeConfiguration", Required: false},
}

var fields_create_task_set = []leanruntime.Field{
	{Name: "CapacityProviderStrategy", Flag: "capacity-provider-strategy", Type: "[]types.CapacityProviderStrategyItem", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Cluster", Flag: "cluster", Type: "*string", Required: true},
	{Name: "ExternalId", Flag: "external-id", Type: "*string", Required: false},
	{Name: "LaunchType", Flag: "launch-type", Type: "types.LaunchType", Required: false},
	{Name: "LoadBalancers", Flag: "load-balancers", Type: "[]types.LoadBalancer", Required: false},
	{Name: "NetworkConfiguration", Flag: "network-configuration", Type: "*types.NetworkConfiguration", Required: false},
	{Name: "PlatformVersion", Flag: "platform-version", Type: "*string", Required: false},
	{Name: "Scale", Flag: "scale", Type: "*types.Scale", Required: false},
	{Name: "Service", Flag: "service", Type: "*string", Required: true},
	{Name: "ServiceRegistries", Flag: "service-registries", Type: "[]types.ServiceRegistry", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TaskDefinition", Flag: "task-definition", Type: "*string", Required: true},
}

var fields_delete_account_setting = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "types.SettingName", Required: true},
	{Name: "PrincipalArn", Flag: "principal-arn", Type: "*string", Required: false},
}

var fields_delete_attributes = []leanruntime.Field{
	{Name: "Attributes", Flag: "attributes", Type: "[]types.Attribute", Required: true},
	{Name: "Cluster", Flag: "cluster", Type: "*string", Required: false},
}

var fields_delete_capacity_provider = []leanruntime.Field{
	{Name: "CapacityProvider", Flag: "capacity-provider", Type: "*string", Required: true},
	{Name: "Cluster", Flag: "cluster", Type: "*string", Required: false},
}

var fields_delete_cluster = []leanruntime.Field{
	{Name: "Cluster", Flag: "cluster", Type: "*string", Required: true},
}

var fields_delete_express_gateway_service = []leanruntime.Field{
	{Name: "ServiceArn", Flag: "service-arn", Type: "*string", Required: true},
}

var fields_delete_service = []leanruntime.Field{
	{Name: "Cluster", Flag: "cluster", Type: "*string", Required: false},
	{Name: "Force", Flag: "force", Type: "*bool", Required: false},
	{Name: "Service", Flag: "service", Type: "*string", Required: true},
}

var fields_delete_task_definitions = []leanruntime.Field{
	{Name: "TaskDefinitions", Flag: "task-definitions", Type: "[]string", Required: true},
}

var fields_delete_task_set = []leanruntime.Field{
	{Name: "Cluster", Flag: "cluster", Type: "*string", Required: true},
	{Name: "Force", Flag: "force", Type: "*bool", Required: false},
	{Name: "Service", Flag: "service", Type: "*string", Required: true},
	{Name: "TaskSet", Flag: "task-set", Type: "*string", Required: true},
}

var fields_deregister_container_instance = []leanruntime.Field{
	{Name: "Cluster", Flag: "cluster", Type: "*string", Required: false},
	{Name: "ContainerInstance", Flag: "container-instance", Type: "*string", Required: true},
	{Name: "Force", Flag: "force", Type: "*bool", Required: false},
}

var fields_deregister_task_definition = []leanruntime.Field{
	{Name: "TaskDefinition", Flag: "task-definition", Type: "*string", Required: true},
}

var fields_describe_capacity_providers = []leanruntime.Field{
	{Name: "CapacityProviders", Flag: "capacity-providers", Type: "[]string", Required: false},
	{Name: "Cluster", Flag: "cluster", Type: "*string", Required: false},
	{Name: "Include", Flag: "include", Type: "[]types.CapacityProviderField", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_clusters = []leanruntime.Field{
	{Name: "Clusters", Flag: "clusters", Type: "[]string", Required: false},
	{Name: "Include", Flag: "include", Type: "[]types.ClusterField", Required: false},
}

var fields_describe_container_instances = []leanruntime.Field{
	{Name: "Cluster", Flag: "cluster", Type: "*string", Required: false},
	{Name: "ContainerInstances", Flag: "container-instances", Type: "[]string", Required: true},
	{Name: "Include", Flag: "include", Type: "[]types.ContainerInstanceField", Required: false},
}

var fields_describe_express_gateway_service = []leanruntime.Field{
	{Name: "Include", Flag: "include", Type: "[]types.ExpressGatewayServiceInclude", Required: false},
	{Name: "ServiceArn", Flag: "service-arn", Type: "*string", Required: true},
}

var fields_describe_service_deployments = []leanruntime.Field{
	{Name: "ServiceDeploymentArns", Flag: "service-deployment-arns", Type: "[]string", Required: true},
}

var fields_describe_service_revisions = []leanruntime.Field{
	{Name: "ServiceRevisionArns", Flag: "service-revision-arns", Type: "[]string", Required: true},
}

var fields_describe_services = []leanruntime.Field{
	{Name: "Cluster", Flag: "cluster", Type: "*string", Required: false},
	{Name: "Include", Flag: "include", Type: "[]types.ServiceField", Required: false},
	{Name: "Services", Flag: "services", Type: "[]string", Required: true},
}

var fields_describe_task_definition = []leanruntime.Field{
	{Name: "Include", Flag: "include", Type: "[]types.TaskDefinitionField", Required: false},
	{Name: "TaskDefinition", Flag: "task-definition", Type: "*string", Required: true},
}

var fields_describe_task_sets = []leanruntime.Field{
	{Name: "Cluster", Flag: "cluster", Type: "*string", Required: true},
	{Name: "Include", Flag: "include", Type: "[]types.TaskSetField", Required: false},
	{Name: "Service", Flag: "service", Type: "*string", Required: true},
	{Name: "TaskSets", Flag: "task-sets", Type: "[]string", Required: false},
}

var fields_describe_tasks = []leanruntime.Field{
	{Name: "Cluster", Flag: "cluster", Type: "*string", Required: false},
	{Name: "Include", Flag: "include", Type: "[]types.TaskField", Required: false},
	{Name: "Tasks", Flag: "tasks", Type: "[]string", Required: true},
}

var fields_discover_poll_endpoint = []leanruntime.Field{
	{Name: "Cluster", Flag: "cluster", Type: "*string", Required: false},
	{Name: "ContainerInstance", Flag: "container-instance", Type: "*string", Required: false},
}

var fields_execute_command = []leanruntime.Field{
	{Name: "Cluster", Flag: "cluster", Type: "*string", Required: false},
	{Name: "Command", Flag: "command", Type: "*string", Required: true},
	{Name: "Container", Flag: "container", Type: "*string", Required: false},
	{Name: "Interactive", Flag: "interactive", Type: "bool", Required: true},
	{Name: "Task", Flag: "task", Type: "*string", Required: true},
}

var fields_get_task_protection = []leanruntime.Field{
	{Name: "Cluster", Flag: "cluster", Type: "*string", Required: true},
	{Name: "Tasks", Flag: "tasks", Type: "[]string", Required: false},
}

var fields_list_account_settings = []leanruntime.Field{
	{Name: "EffectiveSettings", Flag: "effective-settings", Type: "bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "Name", Flag: "name", Type: "types.SettingName", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PrincipalArn", Flag: "principal-arn", Type: "*string", Required: false},
	{Name: "Value", Flag: "value", Type: "*string", Required: false},
}

var fields_list_attributes = []leanruntime.Field{
	{Name: "AttributeName", Flag: "attribute-name", Type: "*string", Required: false},
	{Name: "AttributeValue", Flag: "attribute-value", Type: "*string", Required: false},
	{Name: "Cluster", Flag: "cluster", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TargetType", Flag: "target-type", Type: "types.TargetType", Required: true},
}

var fields_list_clusters = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_container_instances = []leanruntime.Field{
	{Name: "Cluster", Flag: "cluster", Type: "*string", Required: false},
	{Name: "Filter", Flag: "filter", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.ContainerInstanceStatus", Required: false},
}

var fields_list_service_deployments = []leanruntime.Field{
	{Name: "Cluster", Flag: "cluster", Type: "*string", Required: false},
	{Name: "CreatedAt", Flag: "created-at", Type: "*types.CreatedAt", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Service", Flag: "service", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "[]types.ServiceDeploymentStatus", Required: false},
}

var fields_list_services = []leanruntime.Field{
	{Name: "Cluster", Flag: "cluster", Type: "*string", Required: false},
	{Name: "LaunchType", Flag: "launch-type", Type: "types.LaunchType", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceManagementType", Flag: "resource-management-type", Type: "types.ResourceManagementType", Required: false},
	{Name: "SchedulingStrategy", Flag: "scheduling-strategy", Type: "types.SchedulingStrategy", Required: false},
}

var fields_list_services_by_namespace = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_task_definition_families = []leanruntime.Field{
	{Name: "FamilyPrefix", Flag: "family-prefix", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.TaskDefinitionFamilyStatus", Required: false},
}

var fields_list_task_definitions = []leanruntime.Field{
	{Name: "FamilyPrefix", Flag: "family-prefix", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Sort", Flag: "sort", Type: "types.SortOrder", Required: false},
	{Name: "Status", Flag: "status", Type: "types.TaskDefinitionStatus", Required: false},
}

var fields_list_tasks = []leanruntime.Field{
	{Name: "Cluster", Flag: "cluster", Type: "*string", Required: false},
	{Name: "ContainerInstance", Flag: "container-instance", Type: "*string", Required: false},
	{Name: "DesiredStatus", Flag: "desired-status", Type: "types.DesiredStatus", Required: false},
	{Name: "Family", Flag: "family", Type: "*string", Required: false},
	{Name: "LaunchType", Flag: "launch-type", Type: "types.LaunchType", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ServiceName", Flag: "service-name", Type: "*string", Required: false},
	{Name: "StartedBy", Flag: "started-by", Type: "*string", Required: false},
}

var fields_put_account_setting = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "types.SettingName", Required: true},
	{Name: "PrincipalArn", Flag: "principal-arn", Type: "*string", Required: false},
	{Name: "Value", Flag: "value", Type: "*string", Required: true},
}

var fields_put_account_setting_default = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "types.SettingName", Required: true},
	{Name: "Value", Flag: "value", Type: "*string", Required: true},
}

var fields_put_attributes = []leanruntime.Field{
	{Name: "Attributes", Flag: "attributes", Type: "[]types.Attribute", Required: true},
	{Name: "Cluster", Flag: "cluster", Type: "*string", Required: false},
}

var fields_put_cluster_capacity_providers = []leanruntime.Field{
	{Name: "CapacityProviders", Flag: "capacity-providers", Type: "[]string", Required: true},
	{Name: "Cluster", Flag: "cluster", Type: "*string", Required: true},
	{Name: "DefaultCapacityProviderStrategy", Flag: "default-capacity-provider-strategy", Type: "[]types.CapacityProviderStrategyItem", Required: true},
}

var fields_register_container_instance = []leanruntime.Field{
	{Name: "Attributes", Flag: "attributes", Type: "[]types.Attribute", Required: false},
	{Name: "Cluster", Flag: "cluster", Type: "*string", Required: false},
	{Name: "ContainerInstanceArn", Flag: "container-instance-arn", Type: "*string", Required: false},
	{Name: "InstanceIdentityDocument", Flag: "instance-identity-document", Type: "*string", Required: false},
	{Name: "InstanceIdentityDocumentSignature", Flag: "instance-identity-document-signature", Type: "*string", Required: false},
	{Name: "PlatformDevices", Flag: "platform-devices", Type: "[]types.PlatformDevice", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TotalResources", Flag: "total-resources", Type: "[]types.Resource", Required: false},
	{Name: "VersionInfo", Flag: "version-info", Type: "*types.VersionInfo", Required: false},
}

var fields_register_task_definition = []leanruntime.Field{
	{Name: "ContainerDefinitions", Flag: "container-definitions", Type: "[]types.ContainerDefinition", Required: true},
	{Name: "Cpu", Flag: "cpu", Type: "*string", Required: false},
	{Name: "EnableFaultInjection", Flag: "enable-fault-injection", Type: "*bool", Required: false},
	{Name: "EphemeralStorage", Flag: "ephemeral-storage", Type: "*types.EphemeralStorage", Required: false},
	{Name: "ExecutionRoleArn", Flag: "execution-role-arn", Type: "*string", Required: false},
	{Name: "Family", Flag: "family", Type: "*string", Required: true},
	{Name: "InferenceAccelerators", Flag: "inference-accelerators", Type: "[]types.InferenceAccelerator", Required: false},
	{Name: "IpcMode", Flag: "ipc-mode", Type: "types.IpcMode", Required: false},
	{Name: "Memory", Flag: "memory", Type: "*string", Required: false},
	{Name: "NetworkMode", Flag: "network-mode", Type: "types.NetworkMode", Required: false},
	{Name: "PidMode", Flag: "pid-mode", Type: "types.PidMode", Required: false},
	{Name: "PlacementConstraints", Flag: "placement-constraints", Type: "[]types.TaskDefinitionPlacementConstraint", Required: false},
	{Name: "ProxyConfiguration", Flag: "proxy-configuration", Type: "*types.ProxyConfiguration", Required: false},
	{Name: "RequiresCompatibilities", Flag: "requires-compatibilities", Type: "[]types.Compatibility", Required: false},
	{Name: "RuntimePlatform", Flag: "runtime-platform", Type: "*types.RuntimePlatform", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TaskRoleArn", Flag: "task-role-arn", Type: "*string", Required: false},
	{Name: "Volumes", Flag: "volumes", Type: "[]types.Volume", Required: false},
}

var fields_run_task = []leanruntime.Field{
	{Name: "CapacityProviderStrategy", Flag: "capacity-provider-strategy", Type: "[]types.CapacityProviderStrategyItem", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Cluster", Flag: "cluster", Type: "*string", Required: false},
	{Name: "Count", Flag: "count", Type: "*int32", Required: false},
	{Name: "EnableECSManagedTags", Flag: "enable-ecs-managed-tags", Type: "bool", Required: false},
	{Name: "EnableExecuteCommand", Flag: "enable-execute-command", Type: "bool", Required: false},
	{Name: "Group", Flag: "group", Type: "*string", Required: false},
	{Name: "LaunchType", Flag: "launch-type", Type: "types.LaunchType", Required: false},
	{Name: "NetworkConfiguration", Flag: "network-configuration", Type: "*types.NetworkConfiguration", Required: false},
	{Name: "Overrides", Flag: "overrides", Type: "*types.TaskOverride", Required: false},
	{Name: "PlacementConstraints", Flag: "placement-constraints", Type: "[]types.PlacementConstraint", Required: false},
	{Name: "PlacementStrategy", Flag: "placement-strategy", Type: "[]types.PlacementStrategy", Required: false},
	{Name: "PlatformVersion", Flag: "platform-version", Type: "*string", Required: false},
	{Name: "PropagateTags", Flag: "propagate-tags", Type: "types.PropagateTags", Required: false},
	{Name: "ReferenceId", Flag: "reference-id", Type: "*string", Required: false},
	{Name: "StartedBy", Flag: "started-by", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TaskDefinition", Flag: "task-definition", Type: "*string", Required: true},
	{Name: "VolumeConfigurations", Flag: "volume-configurations", Type: "[]types.TaskVolumeConfiguration", Required: false},
}

var fields_start_task = []leanruntime.Field{
	{Name: "Cluster", Flag: "cluster", Type: "*string", Required: false},
	{Name: "ContainerInstances", Flag: "container-instances", Type: "[]string", Required: true},
	{Name: "EnableECSManagedTags", Flag: "enable-ecs-managed-tags", Type: "bool", Required: false},
	{Name: "EnableExecuteCommand", Flag: "enable-execute-command", Type: "bool", Required: false},
	{Name: "Group", Flag: "group", Type: "*string", Required: false},
	{Name: "NetworkConfiguration", Flag: "network-configuration", Type: "*types.NetworkConfiguration", Required: false},
	{Name: "Overrides", Flag: "overrides", Type: "*types.TaskOverride", Required: false},
	{Name: "PropagateTags", Flag: "propagate-tags", Type: "types.PropagateTags", Required: false},
	{Name: "ReferenceId", Flag: "reference-id", Type: "*string", Required: false},
	{Name: "StartedBy", Flag: "started-by", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TaskDefinition", Flag: "task-definition", Type: "*string", Required: true},
	{Name: "VolumeConfigurations", Flag: "volume-configurations", Type: "[]types.TaskVolumeConfiguration", Required: false},
}

var fields_stop_service_deployment = []leanruntime.Field{
	{Name: "ServiceDeploymentArn", Flag: "service-deployment-arn", Type: "*string", Required: true},
	{Name: "StopType", Flag: "stop-type", Type: "types.StopServiceDeploymentStopType", Required: false},
}

var fields_stop_task = []leanruntime.Field{
	{Name: "Cluster", Flag: "cluster", Type: "*string", Required: false},
	{Name: "Reason", Flag: "reason", Type: "*string", Required: false},
	{Name: "Task", Flag: "task", Type: "*string", Required: true},
}

var fields_submit_attachment_state_changes = []leanruntime.Field{
	{Name: "Attachments", Flag: "attachments", Type: "[]types.AttachmentStateChange", Required: true},
	{Name: "Cluster", Flag: "cluster", Type: "*string", Required: false},
}

var fields_submit_container_state_change = []leanruntime.Field{
	{Name: "Cluster", Flag: "cluster", Type: "*string", Required: false},
	{Name: "ContainerName", Flag: "container-name", Type: "*string", Required: false},
	{Name: "ExitCode", Flag: "exit-code", Type: "*int32", Required: false},
	{Name: "NetworkBindings", Flag: "network-bindings", Type: "[]types.NetworkBinding", Required: false},
	{Name: "Reason", Flag: "reason", Type: "*string", Required: false},
	{Name: "RuntimeId", Flag: "runtime-id", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "*string", Required: false},
	{Name: "Task", Flag: "task", Type: "*string", Required: false},
}

var fields_submit_task_state_change = []leanruntime.Field{
	{Name: "Attachments", Flag: "attachments", Type: "[]types.AttachmentStateChange", Required: false},
	{Name: "Cluster", Flag: "cluster", Type: "*string", Required: false},
	{Name: "Containers", Flag: "containers", Type: "[]types.ContainerStateChange", Required: false},
	{Name: "ExecutionStoppedAt", Flag: "execution-stopped-at", Type: "*time.Time", Required: false},
	{Name: "ManagedAgents", Flag: "managed-agents", Type: "[]types.ManagedAgentStateChange", Required: false},
	{Name: "PullStartedAt", Flag: "pull-started-at", Type: "*time.Time", Required: false},
	{Name: "PullStoppedAt", Flag: "pull-stopped-at", Type: "*time.Time", Required: false},
	{Name: "Reason", Flag: "reason", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "*string", Required: false},
	{Name: "Task", Flag: "task", Type: "*string", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_capacity_provider = []leanruntime.Field{
	{Name: "AutoScalingGroupProvider", Flag: "auto-scaling-group-provider", Type: "*types.AutoScalingGroupProviderUpdate", Required: false},
	{Name: "Cluster", Flag: "cluster", Type: "*string", Required: false},
	{Name: "ManagedInstancesProvider", Flag: "managed-instances-provider", Type: "*types.UpdateManagedInstancesProviderConfiguration", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_update_cluster = []leanruntime.Field{
	{Name: "Cluster", Flag: "cluster", Type: "*string", Required: true},
	{Name: "Configuration", Flag: "configuration", Type: "*types.ClusterConfiguration", Required: false},
	{Name: "ServiceConnectDefaults", Flag: "service-connect-defaults", Type: "*types.ClusterServiceConnectDefaultsRequest", Required: false},
	{Name: "Settings", Flag: "settings", Type: "[]types.ClusterSetting", Required: false},
}

var fields_update_cluster_settings = []leanruntime.Field{
	{Name: "Cluster", Flag: "cluster", Type: "*string", Required: true},
	{Name: "Settings", Flag: "settings", Type: "[]types.ClusterSetting", Required: true},
}

var fields_update_container_agent = []leanruntime.Field{
	{Name: "Cluster", Flag: "cluster", Type: "*string", Required: false},
	{Name: "ContainerInstance", Flag: "container-instance", Type: "*string", Required: true},
}

var fields_update_container_instances_state = []leanruntime.Field{
	{Name: "Cluster", Flag: "cluster", Type: "*string", Required: false},
	{Name: "ContainerInstances", Flag: "container-instances", Type: "[]string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.ContainerInstanceStatus", Required: true},
}

var fields_update_express_gateway_service = []leanruntime.Field{
	{Name: "Cpu", Flag: "cpu", Type: "*string", Required: false},
	{Name: "ExecutionRoleArn", Flag: "execution-role-arn", Type: "*string", Required: false},
	{Name: "HealthCheckPath", Flag: "health-check-path", Type: "*string", Required: false},
	{Name: "Memory", Flag: "memory", Type: "*string", Required: false},
	{Name: "NetworkConfiguration", Flag: "network-configuration", Type: "*types.ExpressGatewayServiceNetworkConfiguration", Required: false},
	{Name: "PrimaryContainer", Flag: "primary-container", Type: "*types.ExpressGatewayContainer", Required: false},
	{Name: "ScalingTarget", Flag: "scaling-target", Type: "*types.ExpressGatewayScalingTarget", Required: false},
	{Name: "ServiceArn", Flag: "service-arn", Type: "*string", Required: true},
	{Name: "TaskRoleArn", Flag: "task-role-arn", Type: "*string", Required: false},
}

var fields_update_service = []leanruntime.Field{
	{Name: "AvailabilityZoneRebalancing", Flag: "availability-zone-rebalancing", Type: "types.AvailabilityZoneRebalancing", Required: false},
	{Name: "CapacityProviderStrategy", Flag: "capacity-provider-strategy", Type: "[]types.CapacityProviderStrategyItem", Required: false},
	{Name: "Cluster", Flag: "cluster", Type: "*string", Required: false},
	{Name: "DeploymentConfiguration", Flag: "deployment-configuration", Type: "*types.DeploymentConfiguration", Required: false},
	{Name: "DeploymentController", Flag: "deployment-controller", Type: "*types.DeploymentController", Required: false},
	{Name: "DesiredCount", Flag: "desired-count", Type: "*int32", Required: false},
	{Name: "EnableECSManagedTags", Flag: "enable-ecs-managed-tags", Type: "*bool", Required: false},
	{Name: "EnableExecuteCommand", Flag: "enable-execute-command", Type: "*bool", Required: false},
	{Name: "ForceNewDeployment", Flag: "force-new-deployment", Type: "bool", Required: false},
	{Name: "HealthCheckGracePeriodSeconds", Flag: "health-check-grace-period-seconds", Type: "*int32", Required: false},
	{Name: "LoadBalancers", Flag: "load-balancers", Type: "[]types.LoadBalancer", Required: false},
	{Name: "NetworkConfiguration", Flag: "network-configuration", Type: "*types.NetworkConfiguration", Required: false},
	{Name: "PlacementConstraints", Flag: "placement-constraints", Type: "[]types.PlacementConstraint", Required: false},
	{Name: "PlacementStrategy", Flag: "placement-strategy", Type: "[]types.PlacementStrategy", Required: false},
	{Name: "PlatformVersion", Flag: "platform-version", Type: "*string", Required: false},
	{Name: "PropagateTags", Flag: "propagate-tags", Type: "types.PropagateTags", Required: false},
	{Name: "Service", Flag: "service", Type: "*string", Required: true},
	{Name: "ServiceConnectConfiguration", Flag: "service-connect-configuration", Type: "*types.ServiceConnectConfiguration", Required: false},
	{Name: "ServiceRegistries", Flag: "service-registries", Type: "[]types.ServiceRegistry", Required: false},
	{Name: "TaskDefinition", Flag: "task-definition", Type: "*string", Required: false},
	{Name: "VolumeConfigurations", Flag: "volume-configurations", Type: "[]types.ServiceVolumeConfiguration", Required: false},
	{Name: "VpcLatticeConfigurations", Flag: "vpc-lattice-configurations", Type: "[]types.VpcLatticeConfiguration", Required: false},
}

var fields_update_service_primary_task_set = []leanruntime.Field{
	{Name: "Cluster", Flag: "cluster", Type: "*string", Required: true},
	{Name: "PrimaryTaskSet", Flag: "primary-task-set", Type: "*string", Required: true},
	{Name: "Service", Flag: "service", Type: "*string", Required: true},
}

var fields_update_task_protection = []leanruntime.Field{
	{Name: "Cluster", Flag: "cluster", Type: "*string", Required: true},
	{Name: "ExpiresInMinutes", Flag: "expires-in-minutes", Type: "*int32", Required: false},
	{Name: "ProtectionEnabled", Flag: "protection-enabled", Type: "bool", Required: true},
	{Name: "Tasks", Flag: "tasks", Type: "[]string", Required: true},
}

var fields_update_task_set = []leanruntime.Field{
	{Name: "Cluster", Flag: "cluster", Type: "*string", Required: true},
	{Name: "Scale", Flag: "scale", Type: "*types.Scale", Required: true},
	{Name: "Service", Flag: "service", Type: "*string", Required: true},
	{Name: "TaskSet", Flag: "task-set", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-capacity-provider": {
			Name:   "create-capacity-provider",
			Fields: fields_create_capacity_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCapacityProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_capacity_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCapacityProvider(ctx, input)
			},
		},
		"create-cluster": {
			Name:   "create-cluster",
			Fields: fields_create_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCluster(ctx, input)
			},
		},
		"create-express-gateway-service": {
			Name:   "create-express-gateway-service",
			Fields: fields_create_express_gateway_service,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateExpressGatewayServiceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_express_gateway_service, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateExpressGatewayService(ctx, input)
			},
		},
		"create-service": {
			Name:   "create-service",
			Fields: fields_create_service,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateServiceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_service, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateService(ctx, input)
			},
		},
		"create-task-set": {
			Name:   "create-task-set",
			Fields: fields_create_task_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTaskSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_task_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTaskSet(ctx, input)
			},
		},
		"delete-account-setting": {
			Name:   "delete-account-setting",
			Fields: fields_delete_account_setting,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAccountSettingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_account_setting, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAccountSetting(ctx, input)
			},
		},
		"delete-attributes": {
			Name:   "delete-attributes",
			Fields: fields_delete_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAttributes(ctx, input)
			},
		},
		"delete-capacity-provider": {
			Name:   "delete-capacity-provider",
			Fields: fields_delete_capacity_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCapacityProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_capacity_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCapacityProvider(ctx, input)
			},
		},
		"delete-cluster": {
			Name:   "delete-cluster",
			Fields: fields_delete_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCluster(ctx, input)
			},
		},
		"delete-express-gateway-service": {
			Name:   "delete-express-gateway-service",
			Fields: fields_delete_express_gateway_service,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteExpressGatewayServiceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_express_gateway_service, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteExpressGatewayService(ctx, input)
			},
		},
		"delete-service": {
			Name:   "delete-service",
			Fields: fields_delete_service,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteServiceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_service, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteService(ctx, input)
			},
		},
		"delete-task-definitions": {
			Name:   "delete-task-definitions",
			Fields: fields_delete_task_definitions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTaskDefinitionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_task_definitions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTaskDefinitions(ctx, input)
			},
		},
		"delete-task-set": {
			Name:   "delete-task-set",
			Fields: fields_delete_task_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTaskSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_task_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTaskSet(ctx, input)
			},
		},
		"deregister-container-instance": {
			Name:   "deregister-container-instance",
			Fields: fields_deregister_container_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterContainerInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_container_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterContainerInstance(ctx, input)
			},
		},
		"deregister-task-definition": {
			Name:   "deregister-task-definition",
			Fields: fields_deregister_task_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterTaskDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_task_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterTaskDefinition(ctx, input)
			},
		},
		"describe-capacity-providers": {
			Name:   "describe-capacity-providers",
			Fields: fields_describe_capacity_providers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCapacityProvidersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_capacity_providers, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeCapacityProviders(ctx, input)
			},
		},
		"describe-clusters": {
			Name:   "describe-clusters",
			Fields: fields_describe_clusters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeClustersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_clusters, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeClusters(ctx, input)
			},
		},
		"describe-container-instances": {
			Name:   "describe-container-instances",
			Fields: fields_describe_container_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeContainerInstancesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_container_instances, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeContainerInstances(ctx, input)
			},
		},
		"describe-express-gateway-service": {
			Name:   "describe-express-gateway-service",
			Fields: fields_describe_express_gateway_service,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeExpressGatewayServiceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_express_gateway_service, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeExpressGatewayService(ctx, input)
			},
		},
		"describe-service-deployments": {
			Name:   "describe-service-deployments",
			Fields: fields_describe_service_deployments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeServiceDeploymentsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_service_deployments, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeServiceDeployments(ctx, input)
			},
		},
		"describe-service-revisions": {
			Name:   "describe-service-revisions",
			Fields: fields_describe_service_revisions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeServiceRevisionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_service_revisions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeServiceRevisions(ctx, input)
			},
		},
		"describe-services": {
			Name:   "describe-services",
			Fields: fields_describe_services,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeServicesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_services, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeServices(ctx, input)
			},
		},
		"describe-task-definition": {
			Name:   "describe-task-definition",
			Fields: fields_describe_task_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTaskDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_task_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTaskDefinition(ctx, input)
			},
		},
		"describe-task-sets": {
			Name:   "describe-task-sets",
			Fields: fields_describe_task_sets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTaskSetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_task_sets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTaskSets(ctx, input)
			},
		},
		"describe-tasks": {
			Name:   "describe-tasks",
			Fields: fields_describe_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTasksInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_tasks, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTasks(ctx, input)
			},
		},
		"discover-poll-endpoint": {
			Name:   "discover-poll-endpoint",
			Fields: fields_discover_poll_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DiscoverPollEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_discover_poll_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DiscoverPollEndpoint(ctx, input)
			},
		},
		"execute-command": {
			Name:   "execute-command",
			Fields: fields_execute_command,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExecuteCommandInput{}
				if _, err := leanruntime.ApplyInput(input, fields_execute_command, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExecuteCommand(ctx, input)
			},
		},
		"get-task-protection": {
			Name:   "get-task-protection",
			Fields: fields_get_task_protection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTaskProtectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_task_protection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTaskProtection(ctx, input)
			},
		},
		"list-account-settings": {
			Name:   "list-account-settings",
			Fields: fields_list_account_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAccountSettingsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_account_settings, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAccountSettings(ctx, input)
				}
				var results []*svc.ListAccountSettingsOutput
				p := svc.NewListAccountSettingsPaginator(client, input)
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
		"list-attributes": {
			Name:   "list-attributes",
			Fields: fields_list_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAttributesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_attributes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAttributes(ctx, input)
				}
				var results []*svc.ListAttributesOutput
				p := svc.NewListAttributesPaginator(client, input)
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
		"list-container-instances": {
			Name:   "list-container-instances",
			Fields: fields_list_container_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListContainerInstancesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_container_instances, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListContainerInstances(ctx, input)
				}
				var results []*svc.ListContainerInstancesOutput
				p := svc.NewListContainerInstancesPaginator(client, input)
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
		"list-service-deployments": {
			Name:   "list-service-deployments",
			Fields: fields_list_service_deployments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListServiceDeploymentsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_service_deployments, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListServiceDeployments(ctx, input)
			},
		},
		"list-services": {
			Name:   "list-services",
			Fields: fields_list_services,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListServicesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_services, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListServices(ctx, input)
				}
				var results []*svc.ListServicesOutput
				p := svc.NewListServicesPaginator(client, input)
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
		"list-services-by-namespace": {
			Name:   "list-services-by-namespace",
			Fields: fields_list_services_by_namespace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListServicesByNamespaceInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_services_by_namespace, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListServicesByNamespace(ctx, input)
				}
				var results []*svc.ListServicesByNamespaceOutput
				p := svc.NewListServicesByNamespacePaginator(client, input)
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
		"list-task-definition-families": {
			Name:   "list-task-definition-families",
			Fields: fields_list_task_definition_families,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTaskDefinitionFamiliesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_task_definition_families, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTaskDefinitionFamilies(ctx, input)
				}
				var results []*svc.ListTaskDefinitionFamiliesOutput
				p := svc.NewListTaskDefinitionFamiliesPaginator(client, input)
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
		"list-task-definitions": {
			Name:   "list-task-definitions",
			Fields: fields_list_task_definitions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTaskDefinitionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_task_definitions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTaskDefinitions(ctx, input)
				}
				var results []*svc.ListTaskDefinitionsOutput
				p := svc.NewListTaskDefinitionsPaginator(client, input)
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
		"list-tasks": {
			Name:   "list-tasks",
			Fields: fields_list_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTasksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_tasks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTasks(ctx, input)
				}
				var results []*svc.ListTasksOutput
				p := svc.NewListTasksPaginator(client, input)
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
		"put-account-setting": {
			Name:   "put-account-setting",
			Fields: fields_put_account_setting,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutAccountSettingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_account_setting, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutAccountSetting(ctx, input)
			},
		},
		"put-account-setting-default": {
			Name:   "put-account-setting-default",
			Fields: fields_put_account_setting_default,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutAccountSettingDefaultInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_account_setting_default, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutAccountSettingDefault(ctx, input)
			},
		},
		"put-attributes": {
			Name:   "put-attributes",
			Fields: fields_put_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutAttributes(ctx, input)
			},
		},
		"put-cluster-capacity-providers": {
			Name:   "put-cluster-capacity-providers",
			Fields: fields_put_cluster_capacity_providers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutClusterCapacityProvidersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_cluster_capacity_providers, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutClusterCapacityProviders(ctx, input)
			},
		},
		"register-container-instance": {
			Name:   "register-container-instance",
			Fields: fields_register_container_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterContainerInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_container_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterContainerInstance(ctx, input)
			},
		},
		"register-task-definition": {
			Name:   "register-task-definition",
			Fields: fields_register_task_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterTaskDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_task_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterTaskDefinition(ctx, input)
			},
		},
		"run-task": {
			Name:   "run-task",
			Fields: fields_run_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RunTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_run_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RunTask(ctx, input)
			},
		},
		"start-task": {
			Name:   "start-task",
			Fields: fields_start_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartTask(ctx, input)
			},
		},
		"stop-service-deployment": {
			Name:   "stop-service-deployment",
			Fields: fields_stop_service_deployment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopServiceDeploymentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_service_deployment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopServiceDeployment(ctx, input)
			},
		},
		"stop-task": {
			Name:   "stop-task",
			Fields: fields_stop_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopTask(ctx, input)
			},
		},
		"submit-attachment-state-changes": {
			Name:   "submit-attachment-state-changes",
			Fields: fields_submit_attachment_state_changes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SubmitAttachmentStateChangesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_submit_attachment_state_changes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SubmitAttachmentStateChanges(ctx, input)
			},
		},
		"submit-container-state-change": {
			Name:   "submit-container-state-change",
			Fields: fields_submit_container_state_change,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SubmitContainerStateChangeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_submit_container_state_change, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SubmitContainerStateChange(ctx, input)
			},
		},
		"submit-task-state-change": {
			Name:   "submit-task-state-change",
			Fields: fields_submit_task_state_change,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SubmitTaskStateChangeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_submit_task_state_change, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SubmitTaskStateChange(ctx, input)
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
		"update-capacity-provider": {
			Name:   "update-capacity-provider",
			Fields: fields_update_capacity_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCapacityProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_capacity_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCapacityProvider(ctx, input)
			},
		},
		"update-cluster": {
			Name:   "update-cluster",
			Fields: fields_update_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCluster(ctx, input)
			},
		},
		"update-cluster-settings": {
			Name:   "update-cluster-settings",
			Fields: fields_update_cluster_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateClusterSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_cluster_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateClusterSettings(ctx, input)
			},
		},
		"update-container-agent": {
			Name:   "update-container-agent",
			Fields: fields_update_container_agent,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateContainerAgentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_container_agent, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateContainerAgent(ctx, input)
			},
		},
		"update-container-instances-state": {
			Name:   "update-container-instances-state",
			Fields: fields_update_container_instances_state,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateContainerInstancesStateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_container_instances_state, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateContainerInstancesState(ctx, input)
			},
		},
		"update-express-gateway-service": {
			Name:   "update-express-gateway-service",
			Fields: fields_update_express_gateway_service,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateExpressGatewayServiceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_express_gateway_service, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateExpressGatewayService(ctx, input)
			},
		},
		"update-service": {
			Name:   "update-service",
			Fields: fields_update_service,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateServiceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_service, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateService(ctx, input)
			},
		},
		"update-service-primary-task-set": {
			Name:   "update-service-primary-task-set",
			Fields: fields_update_service_primary_task_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateServicePrimaryTaskSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_service_primary_task_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateServicePrimaryTaskSet(ctx, input)
			},
		},
		"update-task-protection": {
			Name:   "update-task-protection",
			Fields: fields_update_task_protection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTaskProtectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_task_protection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTaskProtection(ctx, input)
			},
		},
		"update-task-set": {
			Name:   "update-task-set",
			Fields: fields_update_task_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTaskSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_task_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTaskSet(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("ecs", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
