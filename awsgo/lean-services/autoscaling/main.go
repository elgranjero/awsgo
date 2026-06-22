package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/autoscaling"
)

var fields_attach_instances = []leanruntime.Field{
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: true},
	{Name: "InstanceIds", Flag: "instance-ids", Type: "[]string", Required: false},
}

var fields_attach_load_balancer_target_groups = []leanruntime.Field{
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: true},
	{Name: "TargetGroupARNs", Flag: "target-group-arns", Type: "[]string", Required: true},
}

var fields_attach_load_balancers = []leanruntime.Field{
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: true},
	{Name: "LoadBalancerNames", Flag: "load-balancer-names", Type: "[]string", Required: true},
}

var fields_attach_traffic_sources = []leanruntime.Field{
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: true},
	{Name: "SkipZonalShiftValidation", Flag: "skip-zonal-shift-validation", Type: "*bool", Required: false},
	{Name: "TrafficSources", Flag: "traffic-sources", Type: "[]types.TrafficSourceIdentifier", Required: true},
}

var fields_batch_delete_scheduled_action = []leanruntime.Field{
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: true},
	{Name: "ScheduledActionNames", Flag: "scheduled-action-names", Type: "[]string", Required: true},
}

var fields_batch_put_scheduled_update_group_action = []leanruntime.Field{
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: true},
	{Name: "ScheduledUpdateGroupActions", Flag: "scheduled-update-group-actions", Type: "[]types.ScheduledUpdateGroupActionRequest", Required: true},
}

var fields_cancel_instance_refresh = []leanruntime.Field{
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: true},
	{Name: "WaitForTransitioningInstances", Flag: "wait-for-transitioning-instances", Type: "*bool", Required: false},
}

var fields_complete_lifecycle_action = []leanruntime.Field{
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: false},
	{Name: "LifecycleActionResult", Flag: "lifecycle-action-result", Type: "*string", Required: true},
	{Name: "LifecycleActionToken", Flag: "lifecycle-action-token", Type: "*string", Required: false},
	{Name: "LifecycleHookName", Flag: "lifecycle-hook-name", Type: "*string", Required: true},
}

var fields_create_auto_scaling_group = []leanruntime.Field{
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: true},
	{Name: "AvailabilityZoneDistribution", Flag: "availability-zone-distribution", Type: "*types.AvailabilityZoneDistribution", Required: false},
	{Name: "AvailabilityZoneImpairmentPolicy", Flag: "availability-zone-impairment-policy", Type: "*types.AvailabilityZoneImpairmentPolicy", Required: false},
	{Name: "AvailabilityZones", Flag: "availability-zones", Type: "[]string", Required: false},
	{Name: "CapacityRebalance", Flag: "capacity-rebalance", Type: "*bool", Required: false},
	{Name: "CapacityReservationSpecification", Flag: "capacity-reservation-specification", Type: "*types.CapacityReservationSpecification", Required: false},
	{Name: "Context", Flag: "context", Type: "*string", Required: false},
	{Name: "DefaultCooldown", Flag: "default-cooldown", Type: "*int32", Required: false},
	{Name: "DefaultInstanceWarmup", Flag: "default-instance-warmup", Type: "*int32", Required: false},
	{Name: "DeletionProtection", Flag: "deletion-protection", Type: "types.DeletionProtection", Required: false},
	{Name: "DesiredCapacity", Flag: "desired-capacity", Type: "*int32", Required: false},
	{Name: "DesiredCapacityType", Flag: "desired-capacity-type", Type: "*string", Required: false},
	{Name: "HealthCheckGracePeriod", Flag: "health-check-grace-period", Type: "*int32", Required: false},
	{Name: "HealthCheckType", Flag: "health-check-type", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: false},
	{Name: "InstanceLifecyclePolicy", Flag: "instance-lifecycle-policy", Type: "*types.InstanceLifecyclePolicy", Required: false},
	{Name: "InstanceMaintenancePolicy", Flag: "instance-maintenance-policy", Type: "*types.InstanceMaintenancePolicy", Required: false},
	{Name: "LaunchConfigurationName", Flag: "launch-configuration-name", Type: "*string", Required: false},
	{Name: "LaunchTemplate", Flag: "launch-template", Type: "*types.LaunchTemplateSpecification", Required: false},
	{Name: "LifecycleHookSpecificationList", Flag: "lifecycle-hook-specification-list", Type: "[]types.LifecycleHookSpecification", Required: false},
	{Name: "LoadBalancerNames", Flag: "load-balancer-names", Type: "[]string", Required: false},
	{Name: "MaxInstanceLifetime", Flag: "max-instance-lifetime", Type: "*int32", Required: false},
	{Name: "MaxSize", Flag: "max-size", Type: "*int32", Required: true},
	{Name: "MinSize", Flag: "min-size", Type: "*int32", Required: true},
	{Name: "MixedInstancesPolicy", Flag: "mixed-instances-policy", Type: "*types.MixedInstancesPolicy", Required: false},
	{Name: "NewInstancesProtectedFromScaleIn", Flag: "new-instances-protected-from-scale-in", Type: "*bool", Required: false},
	{Name: "PlacementGroup", Flag: "placement-group", Type: "*string", Required: false},
	{Name: "ServiceLinkedRoleARN", Flag: "service-linked-role-arn", Type: "*string", Required: false},
	{Name: "SkipZonalShiftValidation", Flag: "skip-zonal-shift-validation", Type: "*bool", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TargetGroupARNs", Flag: "target-group-arns", Type: "[]string", Required: false},
	{Name: "TerminationPolicies", Flag: "termination-policies", Type: "[]string", Required: false},
	{Name: "TrafficSources", Flag: "traffic-sources", Type: "[]types.TrafficSourceIdentifier", Required: false},
	{Name: "VPCZoneIdentifier", Flag: "vpc-zone-identifier", Type: "*string", Required: false},
}

var fields_create_launch_configuration = []leanruntime.Field{
	{Name: "AssociatePublicIpAddress", Flag: "associate-public-ip-address", Type: "*bool", Required: false},
	{Name: "BlockDeviceMappings", Flag: "block-device-mappings", Type: "[]types.BlockDeviceMapping", Required: false},
	{Name: "ClassicLinkVPCId", Flag: "classic-link-vpcid", Type: "*string", Required: false},
	{Name: "ClassicLinkVPCSecurityGroups", Flag: "classic-link-vpc-security-groups", Type: "[]string", Required: false},
	{Name: "EbsOptimized", Flag: "ebs-optimized", Type: "*bool", Required: false},
	{Name: "IamInstanceProfile", Flag: "iam-instance-profile", Type: "*string", Required: false},
	{Name: "ImageId", Flag: "image-id", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: false},
	{Name: "InstanceMonitoring", Flag: "instance-monitoring", Type: "*types.InstanceMonitoring", Required: false},
	{Name: "InstanceType", Flag: "instance-type", Type: "*string", Required: false},
	{Name: "KernelId", Flag: "kernel-id", Type: "*string", Required: false},
	{Name: "KeyName", Flag: "key-name", Type: "*string", Required: false},
	{Name: "LaunchConfigurationName", Flag: "launch-configuration-name", Type: "*string", Required: true},
	{Name: "MetadataOptions", Flag: "metadata-options", Type: "*types.InstanceMetadataOptions", Required: false},
	{Name: "PlacementTenancy", Flag: "placement-tenancy", Type: "*string", Required: false},
	{Name: "RamdiskId", Flag: "ramdisk-id", Type: "*string", Required: false},
	{Name: "SecurityGroups", Flag: "security-groups", Type: "[]string", Required: false},
	{Name: "SpotPrice", Flag: "spot-price", Type: "*string", Required: false},
	{Name: "UserData", Flag: "user-data", Type: "*string", Required: false},
}

var fields_create_or_update_tags = []leanruntime.Field{
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_delete_auto_scaling_group = []leanruntime.Field{
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: true},
	{Name: "ForceDelete", Flag: "force-delete", Type: "*bool", Required: false},
}

var fields_delete_launch_configuration = []leanruntime.Field{
	{Name: "LaunchConfigurationName", Flag: "launch-configuration-name", Type: "*string", Required: true},
}

var fields_delete_lifecycle_hook = []leanruntime.Field{
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: true},
	{Name: "LifecycleHookName", Flag: "lifecycle-hook-name", Type: "*string", Required: true},
}

var fields_delete_notification_configuration = []leanruntime.Field{
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: true},
	{Name: "TopicARN", Flag: "topic-arn", Type: "*string", Required: true},
}

var fields_delete_policy = []leanruntime.Field{
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: false},
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
}

var fields_delete_scheduled_action = []leanruntime.Field{
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: true},
	{Name: "ScheduledActionName", Flag: "scheduled-action-name", Type: "*string", Required: true},
}

var fields_delete_tags = []leanruntime.Field{
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_delete_warm_pool = []leanruntime.Field{
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: true},
	{Name: "ForceDelete", Flag: "force-delete", Type: "*bool", Required: false},
}

var fields_describe_account_limits = []leanruntime.Field{}

var fields_describe_adjustment_types = []leanruntime.Field{}

var fields_describe_auto_scaling_groups = []leanruntime.Field{
	{Name: "AutoScalingGroupNames", Flag: "auto-scaling-group-names", Type: "[]string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "IncludeInstances", Flag: "include-instances", Type: "*bool", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_auto_scaling_instances = []leanruntime.Field{
	{Name: "InstanceIds", Flag: "instance-ids", Type: "[]string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_auto_scaling_notification_types = []leanruntime.Field{}

var fields_describe_instance_refreshes = []leanruntime.Field{
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: true},
	{Name: "InstanceRefreshIds", Flag: "instance-refresh-ids", Type: "[]string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_launch_configurations = []leanruntime.Field{
	{Name: "LaunchConfigurationNames", Flag: "launch-configuration-names", Type: "[]string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_lifecycle_hook_types = []leanruntime.Field{}

var fields_describe_lifecycle_hooks = []leanruntime.Field{
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: true},
	{Name: "LifecycleHookNames", Flag: "lifecycle-hook-names", Type: "[]string", Required: false},
}

var fields_describe_load_balancer_target_groups = []leanruntime.Field{
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: true},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_load_balancers = []leanruntime.Field{
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: true},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_metric_collection_types = []leanruntime.Field{}

var fields_describe_notification_configurations = []leanruntime.Field{
	{Name: "AutoScalingGroupNames", Flag: "auto-scaling-group-names", Type: "[]string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_policies = []leanruntime.Field{
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PolicyNames", Flag: "policy-names", Type: "[]string", Required: false},
	{Name: "PolicyTypes", Flag: "policy-types", Type: "[]string", Required: false},
}

var fields_describe_scaling_activities = []leanruntime.Field{
	{Name: "ActivityIds", Flag: "activity-ids", Type: "[]string", Required: false},
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "IncludeDeletedGroups", Flag: "include-deleted-groups", Type: "*bool", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_scaling_process_types = []leanruntime.Field{}

var fields_describe_scheduled_actions = []leanruntime.Field{
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: false},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ScheduledActionNames", Flag: "scheduled-action-names", Type: "[]string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
}

var fields_describe_tags = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_termination_policy_types = []leanruntime.Field{}

var fields_describe_traffic_sources = []leanruntime.Field{
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: true},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TrafficSourceType", Flag: "traffic-source-type", Type: "*string", Required: false},
}

var fields_describe_warm_pool = []leanruntime.Field{
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: true},
	{Name: "MaxRecords", Flag: "max-records", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_detach_instances = []leanruntime.Field{
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: true},
	{Name: "InstanceIds", Flag: "instance-ids", Type: "[]string", Required: false},
	{Name: "ShouldDecrementDesiredCapacity", Flag: "should-decrement-desired-capacity", Type: "*bool", Required: true},
}

var fields_detach_load_balancer_target_groups = []leanruntime.Field{
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: true},
	{Name: "TargetGroupARNs", Flag: "target-group-arns", Type: "[]string", Required: true},
}

var fields_detach_load_balancers = []leanruntime.Field{
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: true},
	{Name: "LoadBalancerNames", Flag: "load-balancer-names", Type: "[]string", Required: true},
}

var fields_detach_traffic_sources = []leanruntime.Field{
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: true},
	{Name: "TrafficSources", Flag: "traffic-sources", Type: "[]types.TrafficSourceIdentifier", Required: true},
}

var fields_disable_metrics_collection = []leanruntime.Field{
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: true},
	{Name: "Metrics", Flag: "metrics", Type: "[]string", Required: false},
}

var fields_enable_metrics_collection = []leanruntime.Field{
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: true},
	{Name: "Granularity", Flag: "granularity", Type: "*string", Required: true},
	{Name: "Metrics", Flag: "metrics", Type: "[]string", Required: false},
}

var fields_enter_standby = []leanruntime.Field{
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: true},
	{Name: "InstanceIds", Flag: "instance-ids", Type: "[]string", Required: false},
	{Name: "ShouldDecrementDesiredCapacity", Flag: "should-decrement-desired-capacity", Type: "*bool", Required: true},
}

var fields_execute_policy = []leanruntime.Field{
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: false},
	{Name: "BreachThreshold", Flag: "breach-threshold", Type: "*float64", Required: false},
	{Name: "HonorCooldown", Flag: "honor-cooldown", Type: "*bool", Required: false},
	{Name: "MetricValue", Flag: "metric-value", Type: "*float64", Required: false},
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
}

var fields_exit_standby = []leanruntime.Field{
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: true},
	{Name: "InstanceIds", Flag: "instance-ids", Type: "[]string", Required: false},
}

var fields_get_predictive_scaling_forecast = []leanruntime.Field{
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: true},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
}

var fields_launch_instances = []leanruntime.Field{
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: true},
	{Name: "AvailabilityZoneIds", Flag: "availability-zone-ids", Type: "[]string", Required: false},
	{Name: "AvailabilityZones", Flag: "availability-zones", Type: "[]string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "RequestedCapacity", Flag: "requested-capacity", Type: "*int32", Required: true},
	{Name: "RetryStrategy", Flag: "retry-strategy", Type: "types.RetryStrategy", Required: false},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: false},
}

var fields_put_lifecycle_hook = []leanruntime.Field{
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: true},
	{Name: "DefaultResult", Flag: "default-result", Type: "*string", Required: false},
	{Name: "HeartbeatTimeout", Flag: "heartbeat-timeout", Type: "*int32", Required: false},
	{Name: "LifecycleHookName", Flag: "lifecycle-hook-name", Type: "*string", Required: true},
	{Name: "LifecycleTransition", Flag: "lifecycle-transition", Type: "*string", Required: false},
	{Name: "NotificationMetadata", Flag: "notification-metadata", Type: "*string", Required: false},
	{Name: "NotificationTargetARN", Flag: "notification-target-arn", Type: "*string", Required: false},
	{Name: "RoleARN", Flag: "role-arn", Type: "*string", Required: false},
}

var fields_put_notification_configuration = []leanruntime.Field{
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: true},
	{Name: "NotificationTypes", Flag: "notification-types", Type: "[]string", Required: true},
	{Name: "TopicARN", Flag: "topic-arn", Type: "*string", Required: true},
}

var fields_put_scaling_policy = []leanruntime.Field{
	{Name: "AdjustmentType", Flag: "adjustment-type", Type: "*string", Required: false},
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: true},
	{Name: "Cooldown", Flag: "cooldown", Type: "*int32", Required: false},
	{Name: "Enabled", Flag: "enabled", Type: "*bool", Required: false},
	{Name: "EstimatedInstanceWarmup", Flag: "estimated-instance-warmup", Type: "*int32", Required: false},
	{Name: "MetricAggregationType", Flag: "metric-aggregation-type", Type: "*string", Required: false},
	{Name: "MinAdjustmentMagnitude", Flag: "min-adjustment-magnitude", Type: "*int32", Required: false},
	{Name: "MinAdjustmentStep", Flag: "min-adjustment-step", Type: "*int32", Required: false},
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
	{Name: "PolicyType", Flag: "policy-type", Type: "*string", Required: false},
	{Name: "PredictiveScalingConfiguration", Flag: "predictive-scaling-configuration", Type: "*types.PredictiveScalingConfiguration", Required: false},
	{Name: "ScalingAdjustment", Flag: "scaling-adjustment", Type: "*int32", Required: false},
	{Name: "StepAdjustments", Flag: "step-adjustments", Type: "[]types.StepAdjustment", Required: false},
	{Name: "TargetTrackingConfiguration", Flag: "target-tracking-configuration", Type: "*types.TargetTrackingConfiguration", Required: false},
}

var fields_put_scheduled_update_group_action = []leanruntime.Field{
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: true},
	{Name: "DesiredCapacity", Flag: "desired-capacity", Type: "*int32", Required: false},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "MaxSize", Flag: "max-size", Type: "*int32", Required: false},
	{Name: "MinSize", Flag: "min-size", Type: "*int32", Required: false},
	{Name: "Recurrence", Flag: "recurrence", Type: "*string", Required: false},
	{Name: "ScheduledActionName", Flag: "scheduled-action-name", Type: "*string", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
	{Name: "Time", Flag: "time", Type: "*time.Time", Required: false},
	{Name: "TimeZone", Flag: "time-zone", Type: "*string", Required: false},
}

var fields_put_warm_pool = []leanruntime.Field{
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: true},
	{Name: "InstanceReusePolicy", Flag: "instance-reuse-policy", Type: "*types.InstanceReusePolicy", Required: false},
	{Name: "MaxGroupPreparedCapacity", Flag: "max-group-prepared-capacity", Type: "*int32", Required: false},
	{Name: "MinSize", Flag: "min-size", Type: "*int32", Required: false},
	{Name: "PoolState", Flag: "pool-state", Type: "types.WarmPoolState", Required: false},
}

var fields_record_lifecycle_action_heartbeat = []leanruntime.Field{
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: false},
	{Name: "LifecycleActionToken", Flag: "lifecycle-action-token", Type: "*string", Required: false},
	{Name: "LifecycleHookName", Flag: "lifecycle-hook-name", Type: "*string", Required: true},
}

var fields_resume_processes = []leanruntime.Field{
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: true},
	{Name: "ScalingProcesses", Flag: "scaling-processes", Type: "[]string", Required: false},
}

var fields_rollback_instance_refresh = []leanruntime.Field{
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: true},
}

var fields_set_desired_capacity = []leanruntime.Field{
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: true},
	{Name: "DesiredCapacity", Flag: "desired-capacity", Type: "*int32", Required: true},
	{Name: "HonorCooldown", Flag: "honor-cooldown", Type: "*bool", Required: false},
}

var fields_set_instance_health = []leanruntime.Field{
	{Name: "HealthStatus", Flag: "health-status", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "ShouldRespectGracePeriod", Flag: "should-respect-grace-period", Type: "*bool", Required: false},
}

var fields_set_instance_protection = []leanruntime.Field{
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: true},
	{Name: "InstanceIds", Flag: "instance-ids", Type: "[]string", Required: true},
	{Name: "ProtectedFromScaleIn", Flag: "protected-from-scale-in", Type: "*bool", Required: true},
}

var fields_start_instance_refresh = []leanruntime.Field{
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: true},
	{Name: "DesiredConfiguration", Flag: "desired-configuration", Type: "*types.DesiredConfiguration", Required: false},
	{Name: "Preferences", Flag: "preferences", Type: "*types.RefreshPreferences", Required: false},
	{Name: "Strategy", Flag: "strategy", Type: "types.RefreshStrategy", Required: false},
}

var fields_suspend_processes = []leanruntime.Field{
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: true},
	{Name: "ScalingProcesses", Flag: "scaling-processes", Type: "[]string", Required: false},
}

var fields_terminate_instance_in_auto_scaling_group = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "ShouldDecrementDesiredCapacity", Flag: "should-decrement-desired-capacity", Type: "*bool", Required: true},
}

var fields_update_auto_scaling_group = []leanruntime.Field{
	{Name: "AutoScalingGroupName", Flag: "auto-scaling-group-name", Type: "*string", Required: true},
	{Name: "AvailabilityZoneDistribution", Flag: "availability-zone-distribution", Type: "*types.AvailabilityZoneDistribution", Required: false},
	{Name: "AvailabilityZoneImpairmentPolicy", Flag: "availability-zone-impairment-policy", Type: "*types.AvailabilityZoneImpairmentPolicy", Required: false},
	{Name: "AvailabilityZones", Flag: "availability-zones", Type: "[]string", Required: false},
	{Name: "CapacityRebalance", Flag: "capacity-rebalance", Type: "*bool", Required: false},
	{Name: "CapacityReservationSpecification", Flag: "capacity-reservation-specification", Type: "*types.CapacityReservationSpecification", Required: false},
	{Name: "Context", Flag: "context", Type: "*string", Required: false},
	{Name: "DefaultCooldown", Flag: "default-cooldown", Type: "*int32", Required: false},
	{Name: "DefaultInstanceWarmup", Flag: "default-instance-warmup", Type: "*int32", Required: false},
	{Name: "DeletionProtection", Flag: "deletion-protection", Type: "types.DeletionProtection", Required: false},
	{Name: "DesiredCapacity", Flag: "desired-capacity", Type: "*int32", Required: false},
	{Name: "DesiredCapacityType", Flag: "desired-capacity-type", Type: "*string", Required: false},
	{Name: "HealthCheckGracePeriod", Flag: "health-check-grace-period", Type: "*int32", Required: false},
	{Name: "HealthCheckType", Flag: "health-check-type", Type: "*string", Required: false},
	{Name: "InstanceLifecyclePolicy", Flag: "instance-lifecycle-policy", Type: "*types.InstanceLifecyclePolicy", Required: false},
	{Name: "InstanceMaintenancePolicy", Flag: "instance-maintenance-policy", Type: "*types.InstanceMaintenancePolicy", Required: false},
	{Name: "LaunchConfigurationName", Flag: "launch-configuration-name", Type: "*string", Required: false},
	{Name: "LaunchTemplate", Flag: "launch-template", Type: "*types.LaunchTemplateSpecification", Required: false},
	{Name: "MaxInstanceLifetime", Flag: "max-instance-lifetime", Type: "*int32", Required: false},
	{Name: "MaxSize", Flag: "max-size", Type: "*int32", Required: false},
	{Name: "MinSize", Flag: "min-size", Type: "*int32", Required: false},
	{Name: "MixedInstancesPolicy", Flag: "mixed-instances-policy", Type: "*types.MixedInstancesPolicy", Required: false},
	{Name: "NewInstancesProtectedFromScaleIn", Flag: "new-instances-protected-from-scale-in", Type: "*bool", Required: false},
	{Name: "PlacementGroup", Flag: "placement-group", Type: "*string", Required: false},
	{Name: "ServiceLinkedRoleARN", Flag: "service-linked-role-arn", Type: "*string", Required: false},
	{Name: "SkipZonalShiftValidation", Flag: "skip-zonal-shift-validation", Type: "*bool", Required: false},
	{Name: "TerminationPolicies", Flag: "termination-policies", Type: "[]string", Required: false},
	{Name: "VPCZoneIdentifier", Flag: "vpc-zone-identifier", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"attach-instances": {
			Name:   "attach-instances",
			Fields: fields_attach_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AttachInstancesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_attach_instances, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AttachInstances(ctx, input)
			},
		},
		"attach-load-balancer-target-groups": {
			Name:   "attach-load-balancer-target-groups",
			Fields: fields_attach_load_balancer_target_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AttachLoadBalancerTargetGroupsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_attach_load_balancer_target_groups, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AttachLoadBalancerTargetGroups(ctx, input)
			},
		},
		"attach-load-balancers": {
			Name:   "attach-load-balancers",
			Fields: fields_attach_load_balancers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AttachLoadBalancersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_attach_load_balancers, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AttachLoadBalancers(ctx, input)
			},
		},
		"attach-traffic-sources": {
			Name:   "attach-traffic-sources",
			Fields: fields_attach_traffic_sources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AttachTrafficSourcesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_attach_traffic_sources, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AttachTrafficSources(ctx, input)
			},
		},
		"batch-delete-scheduled-action": {
			Name:   "batch-delete-scheduled-action",
			Fields: fields_batch_delete_scheduled_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDeleteScheduledActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_delete_scheduled_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDeleteScheduledAction(ctx, input)
			},
		},
		"batch-put-scheduled-update-group-action": {
			Name:   "batch-put-scheduled-update-group-action",
			Fields: fields_batch_put_scheduled_update_group_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchPutScheduledUpdateGroupActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_put_scheduled_update_group_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchPutScheduledUpdateGroupAction(ctx, input)
			},
		},
		"cancel-instance-refresh": {
			Name:   "cancel-instance-refresh",
			Fields: fields_cancel_instance_refresh,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelInstanceRefreshInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_instance_refresh, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelInstanceRefresh(ctx, input)
			},
		},
		"complete-lifecycle-action": {
			Name:   "complete-lifecycle-action",
			Fields: fields_complete_lifecycle_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CompleteLifecycleActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_complete_lifecycle_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CompleteLifecycleAction(ctx, input)
			},
		},
		"create-auto-scaling-group": {
			Name:   "create-auto-scaling-group",
			Fields: fields_create_auto_scaling_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAutoScalingGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_auto_scaling_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAutoScalingGroup(ctx, input)
			},
		},
		"create-launch-configuration": {
			Name:   "create-launch-configuration",
			Fields: fields_create_launch_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLaunchConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_launch_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLaunchConfiguration(ctx, input)
			},
		},
		"create-or-update-tags": {
			Name:   "create-or-update-tags",
			Fields: fields_create_or_update_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateOrUpdateTagsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_or_update_tags, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateOrUpdateTags(ctx, input)
			},
		},
		"delete-auto-scaling-group": {
			Name:   "delete-auto-scaling-group",
			Fields: fields_delete_auto_scaling_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAutoScalingGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_auto_scaling_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAutoScalingGroup(ctx, input)
			},
		},
		"delete-launch-configuration": {
			Name:   "delete-launch-configuration",
			Fields: fields_delete_launch_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLaunchConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_launch_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLaunchConfiguration(ctx, input)
			},
		},
		"delete-lifecycle-hook": {
			Name:   "delete-lifecycle-hook",
			Fields: fields_delete_lifecycle_hook,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLifecycleHookInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_lifecycle_hook, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLifecycleHook(ctx, input)
			},
		},
		"delete-notification-configuration": {
			Name:   "delete-notification-configuration",
			Fields: fields_delete_notification_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteNotificationConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_notification_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteNotificationConfiguration(ctx, input)
			},
		},
		"delete-policy": {
			Name:   "delete-policy",
			Fields: fields_delete_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePolicy(ctx, input)
			},
		},
		"delete-scheduled-action": {
			Name:   "delete-scheduled-action",
			Fields: fields_delete_scheduled_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteScheduledActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_scheduled_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteScheduledAction(ctx, input)
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
		"delete-warm-pool": {
			Name:   "delete-warm-pool",
			Fields: fields_delete_warm_pool,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWarmPoolInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_warm_pool, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWarmPool(ctx, input)
			},
		},
		"describe-account-limits": {
			Name:   "describe-account-limits",
			Fields: fields_describe_account_limits,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAccountLimitsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_account_limits, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAccountLimits(ctx, input)
			},
		},
		"describe-adjustment-types": {
			Name:   "describe-adjustment-types",
			Fields: fields_describe_adjustment_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAdjustmentTypesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_adjustment_types, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAdjustmentTypes(ctx, input)
			},
		},
		"describe-auto-scaling-groups": {
			Name:   "describe-auto-scaling-groups",
			Fields: fields_describe_auto_scaling_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAutoScalingGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_auto_scaling_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeAutoScalingGroups(ctx, input)
				}
				var results []*svc.DescribeAutoScalingGroupsOutput
				p := svc.NewDescribeAutoScalingGroupsPaginator(client, input)
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
		"describe-auto-scaling-instances": {
			Name:   "describe-auto-scaling-instances",
			Fields: fields_describe_auto_scaling_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAutoScalingInstancesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_auto_scaling_instances, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeAutoScalingInstances(ctx, input)
				}
				var results []*svc.DescribeAutoScalingInstancesOutput
				p := svc.NewDescribeAutoScalingInstancesPaginator(client, input)
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
		"describe-auto-scaling-notification-types": {
			Name:   "describe-auto-scaling-notification-types",
			Fields: fields_describe_auto_scaling_notification_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAutoScalingNotificationTypesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_auto_scaling_notification_types, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAutoScalingNotificationTypes(ctx, input)
			},
		},
		"describe-instance-refreshes": {
			Name:   "describe-instance-refreshes",
			Fields: fields_describe_instance_refreshes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInstanceRefreshesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_instance_refreshes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeInstanceRefreshes(ctx, input)
				}
				var results []*svc.DescribeInstanceRefreshesOutput
				p := svc.NewDescribeInstanceRefreshesPaginator(client, input)
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
		"describe-launch-configurations": {
			Name:   "describe-launch-configurations",
			Fields: fields_describe_launch_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLaunchConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_launch_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeLaunchConfigurations(ctx, input)
				}
				var results []*svc.DescribeLaunchConfigurationsOutput
				p := svc.NewDescribeLaunchConfigurationsPaginator(client, input)
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
		"describe-lifecycle-hook-types": {
			Name:   "describe-lifecycle-hook-types",
			Fields: fields_describe_lifecycle_hook_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLifecycleHookTypesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_lifecycle_hook_types, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeLifecycleHookTypes(ctx, input)
			},
		},
		"describe-lifecycle-hooks": {
			Name:   "describe-lifecycle-hooks",
			Fields: fields_describe_lifecycle_hooks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLifecycleHooksInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_lifecycle_hooks, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeLifecycleHooks(ctx, input)
			},
		},
		"describe-load-balancer-target-groups": {
			Name:   "describe-load-balancer-target-groups",
			Fields: fields_describe_load_balancer_target_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLoadBalancerTargetGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_load_balancer_target_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeLoadBalancerTargetGroups(ctx, input)
				}
				var results []*svc.DescribeLoadBalancerTargetGroupsOutput
				p := svc.NewDescribeLoadBalancerTargetGroupsPaginator(client, input)
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
		"describe-load-balancers": {
			Name:   "describe-load-balancers",
			Fields: fields_describe_load_balancers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLoadBalancersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_load_balancers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeLoadBalancers(ctx, input)
				}
				var results []*svc.DescribeLoadBalancersOutput
				p := svc.NewDescribeLoadBalancersPaginator(client, input)
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
		"describe-metric-collection-types": {
			Name:   "describe-metric-collection-types",
			Fields: fields_describe_metric_collection_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMetricCollectionTypesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_metric_collection_types, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeMetricCollectionTypes(ctx, input)
			},
		},
		"describe-notification-configurations": {
			Name:   "describe-notification-configurations",
			Fields: fields_describe_notification_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeNotificationConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_notification_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeNotificationConfigurations(ctx, input)
				}
				var results []*svc.DescribeNotificationConfigurationsOutput
				p := svc.NewDescribeNotificationConfigurationsPaginator(client, input)
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
		"describe-policies": {
			Name:   "describe-policies",
			Fields: fields_describe_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePoliciesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_policies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribePolicies(ctx, input)
				}
				var results []*svc.DescribePoliciesOutput
				p := svc.NewDescribePoliciesPaginator(client, input)
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
		"describe-scaling-activities": {
			Name:   "describe-scaling-activities",
			Fields: fields_describe_scaling_activities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeScalingActivitiesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_scaling_activities, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeScalingActivities(ctx, input)
				}
				var results []*svc.DescribeScalingActivitiesOutput
				p := svc.NewDescribeScalingActivitiesPaginator(client, input)
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
		"describe-scaling-process-types": {
			Name:   "describe-scaling-process-types",
			Fields: fields_describe_scaling_process_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeScalingProcessTypesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_scaling_process_types, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeScalingProcessTypes(ctx, input)
			},
		},
		"describe-scheduled-actions": {
			Name:   "describe-scheduled-actions",
			Fields: fields_describe_scheduled_actions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeScheduledActionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_scheduled_actions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeScheduledActions(ctx, input)
				}
				var results []*svc.DescribeScheduledActionsOutput
				p := svc.NewDescribeScheduledActionsPaginator(client, input)
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
		"describe-tags": {
			Name:   "describe-tags",
			Fields: fields_describe_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTagsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_tags, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeTags(ctx, input)
				}
				var results []*svc.DescribeTagsOutput
				p := svc.NewDescribeTagsPaginator(client, input)
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
		"describe-termination-policy-types": {
			Name:   "describe-termination-policy-types",
			Fields: fields_describe_termination_policy_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTerminationPolicyTypesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_termination_policy_types, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTerminationPolicyTypes(ctx, input)
			},
		},
		"describe-traffic-sources": {
			Name:   "describe-traffic-sources",
			Fields: fields_describe_traffic_sources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTrafficSourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_traffic_sources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeTrafficSources(ctx, input)
				}
				var results []*svc.DescribeTrafficSourcesOutput
				p := svc.NewDescribeTrafficSourcesPaginator(client, input)
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
		"describe-warm-pool": {
			Name:   "describe-warm-pool",
			Fields: fields_describe_warm_pool,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeWarmPoolInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_warm_pool, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeWarmPool(ctx, input)
				}
				var results []*svc.DescribeWarmPoolOutput
				p := svc.NewDescribeWarmPoolPaginator(client, input)
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
		"detach-instances": {
			Name:   "detach-instances",
			Fields: fields_detach_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetachInstancesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detach_instances, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetachInstances(ctx, input)
			},
		},
		"detach-load-balancer-target-groups": {
			Name:   "detach-load-balancer-target-groups",
			Fields: fields_detach_load_balancer_target_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetachLoadBalancerTargetGroupsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detach_load_balancer_target_groups, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetachLoadBalancerTargetGroups(ctx, input)
			},
		},
		"detach-load-balancers": {
			Name:   "detach-load-balancers",
			Fields: fields_detach_load_balancers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetachLoadBalancersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detach_load_balancers, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetachLoadBalancers(ctx, input)
			},
		},
		"detach-traffic-sources": {
			Name:   "detach-traffic-sources",
			Fields: fields_detach_traffic_sources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetachTrafficSourcesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detach_traffic_sources, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetachTrafficSources(ctx, input)
			},
		},
		"disable-metrics-collection": {
			Name:   "disable-metrics-collection",
			Fields: fields_disable_metrics_collection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableMetricsCollectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_metrics_collection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableMetricsCollection(ctx, input)
			},
		},
		"enable-metrics-collection": {
			Name:   "enable-metrics-collection",
			Fields: fields_enable_metrics_collection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableMetricsCollectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_metrics_collection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableMetricsCollection(ctx, input)
			},
		},
		"enter-standby": {
			Name:   "enter-standby",
			Fields: fields_enter_standby,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnterStandbyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enter_standby, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnterStandby(ctx, input)
			},
		},
		"execute-policy": {
			Name:   "execute-policy",
			Fields: fields_execute_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExecutePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_execute_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExecutePolicy(ctx, input)
			},
		},
		"exit-standby": {
			Name:   "exit-standby",
			Fields: fields_exit_standby,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExitStandbyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_exit_standby, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExitStandby(ctx, input)
			},
		},
		"get-predictive-scaling-forecast": {
			Name:   "get-predictive-scaling-forecast",
			Fields: fields_get_predictive_scaling_forecast,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPredictiveScalingForecastInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_predictive_scaling_forecast, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPredictiveScalingForecast(ctx, input)
			},
		},
		"launch-instances": {
			Name:   "launch-instances",
			Fields: fields_launch_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.LaunchInstancesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_launch_instances, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.LaunchInstances(ctx, input)
			},
		},
		"put-lifecycle-hook": {
			Name:   "put-lifecycle-hook",
			Fields: fields_put_lifecycle_hook,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutLifecycleHookInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_lifecycle_hook, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutLifecycleHook(ctx, input)
			},
		},
		"put-notification-configuration": {
			Name:   "put-notification-configuration",
			Fields: fields_put_notification_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutNotificationConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_notification_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutNotificationConfiguration(ctx, input)
			},
		},
		"put-scaling-policy": {
			Name:   "put-scaling-policy",
			Fields: fields_put_scaling_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutScalingPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_scaling_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutScalingPolicy(ctx, input)
			},
		},
		"put-scheduled-update-group-action": {
			Name:   "put-scheduled-update-group-action",
			Fields: fields_put_scheduled_update_group_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutScheduledUpdateGroupActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_scheduled_update_group_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutScheduledUpdateGroupAction(ctx, input)
			},
		},
		"put-warm-pool": {
			Name:   "put-warm-pool",
			Fields: fields_put_warm_pool,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutWarmPoolInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_warm_pool, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutWarmPool(ctx, input)
			},
		},
		"record-lifecycle-action-heartbeat": {
			Name:   "record-lifecycle-action-heartbeat",
			Fields: fields_record_lifecycle_action_heartbeat,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RecordLifecycleActionHeartbeatInput{}
				if _, err := leanruntime.ApplyInput(input, fields_record_lifecycle_action_heartbeat, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RecordLifecycleActionHeartbeat(ctx, input)
			},
		},
		"resume-processes": {
			Name:   "resume-processes",
			Fields: fields_resume_processes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResumeProcessesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_resume_processes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResumeProcesses(ctx, input)
			},
		},
		"rollback-instance-refresh": {
			Name:   "rollback-instance-refresh",
			Fields: fields_rollback_instance_refresh,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RollbackInstanceRefreshInput{}
				if _, err := leanruntime.ApplyInput(input, fields_rollback_instance_refresh, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RollbackInstanceRefresh(ctx, input)
			},
		},
		"set-desired-capacity": {
			Name:   "set-desired-capacity",
			Fields: fields_set_desired_capacity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetDesiredCapacityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_desired_capacity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetDesiredCapacity(ctx, input)
			},
		},
		"set-instance-health": {
			Name:   "set-instance-health",
			Fields: fields_set_instance_health,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetInstanceHealthInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_instance_health, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetInstanceHealth(ctx, input)
			},
		},
		"set-instance-protection": {
			Name:   "set-instance-protection",
			Fields: fields_set_instance_protection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetInstanceProtectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_instance_protection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetInstanceProtection(ctx, input)
			},
		},
		"start-instance-refresh": {
			Name:   "start-instance-refresh",
			Fields: fields_start_instance_refresh,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartInstanceRefreshInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_instance_refresh, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartInstanceRefresh(ctx, input)
			},
		},
		"suspend-processes": {
			Name:   "suspend-processes",
			Fields: fields_suspend_processes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SuspendProcessesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_suspend_processes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SuspendProcesses(ctx, input)
			},
		},
		"terminate-instance-in-auto-scaling-group": {
			Name:   "terminate-instance-in-auto-scaling-group",
			Fields: fields_terminate_instance_in_auto_scaling_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TerminateInstanceInAutoScalingGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_terminate_instance_in_auto_scaling_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TerminateInstanceInAutoScalingGroup(ctx, input)
			},
		},
		"update-auto-scaling-group": {
			Name:   "update-auto-scaling-group",
			Fields: fields_update_auto_scaling_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAutoScalingGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_auto_scaling_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAutoScalingGroup(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("autoscaling", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
