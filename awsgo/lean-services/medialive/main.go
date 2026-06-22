package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/medialive"
)

var fields_accept_input_device_transfer = []leanruntime.Field{
	{Name: "InputDeviceId", Flag: "input-device-id", Type: "*string", Required: true},
}

var fields_batch_delete = []leanruntime.Field{
	{Name: "ChannelIds", Flag: "channel-ids", Type: "[]string", Required: false},
	{Name: "InputIds", Flag: "input-ids", Type: "[]string", Required: false},
	{Name: "InputSecurityGroupIds", Flag: "input-security-group-ids", Type: "[]string", Required: false},
	{Name: "MultiplexIds", Flag: "multiplex-ids", Type: "[]string", Required: false},
}

var fields_batch_start = []leanruntime.Field{
	{Name: "ChannelIds", Flag: "channel-ids", Type: "[]string", Required: false},
	{Name: "MultiplexIds", Flag: "multiplex-ids", Type: "[]string", Required: false},
}

var fields_batch_stop = []leanruntime.Field{
	{Name: "ChannelIds", Flag: "channel-ids", Type: "[]string", Required: false},
	{Name: "MultiplexIds", Flag: "multiplex-ids", Type: "[]string", Required: false},
}

var fields_batch_update_schedule = []leanruntime.Field{
	{Name: "ChannelId", Flag: "channel-id", Type: "*string", Required: true},
	{Name: "Creates", Flag: "creates", Type: "*types.BatchScheduleActionCreateRequest", Required: false},
	{Name: "Deletes", Flag: "deletes", Type: "*types.BatchScheduleActionDeleteRequest", Required: false},
}

var fields_cancel_input_device_transfer = []leanruntime.Field{
	{Name: "InputDeviceId", Flag: "input-device-id", Type: "*string", Required: true},
}

var fields_claim_device = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: false},
}

var fields_create_channel = []leanruntime.Field{
	{Name: "AnywhereSettings", Flag: "anywhere-settings", Type: "*types.AnywhereSettings", Required: false},
	{Name: "CdiInputSpecification", Flag: "cdi-input-specification", Type: "*types.CdiInputSpecification", Required: false},
	{Name: "ChannelClass", Flag: "channel-class", Type: "types.ChannelClass", Required: false},
	{Name: "ChannelEngineVersion", Flag: "channel-engine-version", Type: "*types.ChannelEngineVersionRequest", Required: false},
	{Name: "ChannelSecurityGroups", Flag: "channel-security-groups", Type: "[]string", Required: false},
	{Name: "Destinations", Flag: "destinations", Type: "[]types.OutputDestination", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "EncoderSettings", Flag: "encoder-settings", Type: "*types.EncoderSettings", Required: false},
	{Name: "InferenceSettings", Flag: "inference-settings", Type: "*types.InferenceSettings", Required: false},
	{Name: "InputAttachments", Flag: "input-attachments", Type: "[]types.InputAttachment", Required: false},
	{Name: "InputSpecification", Flag: "input-specification", Type: "*types.InputSpecification", Required: false},
	{Name: "LinkedChannelSettings", Flag: "linked-channel-settings", Type: "*types.LinkedChannelSettings", Required: false},
	{Name: "LogLevel", Flag: "log-level", Type: "types.LogLevel", Required: false},
	{Name: "Maintenance", Flag: "maintenance", Type: "*types.MaintenanceCreateSettings", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "RequestId", Flag: "request-id", Type: "*string", Required: false},
	{Name: "Reserved", Flag: "reserved", Type: "*string", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Vpc", Flag: "vpc", Type: "*types.VpcOutputSettings", Required: false},
}

var fields_create_channel_placement_group = []leanruntime.Field{
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Nodes", Flag: "nodes", Type: "[]string", Required: false},
	{Name: "RequestId", Flag: "request-id", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_cloud_watch_alarm_template = []leanruntime.Field{
	{Name: "ComparisonOperator", Flag: "comparison-operator", Type: "types.CloudWatchAlarmTemplateComparisonOperator", Required: true},
	{Name: "DatapointsToAlarm", Flag: "datapoints-to-alarm", Type: "*int32", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EvaluationPeriods", Flag: "evaluation-periods", Type: "*int32", Required: true},
	{Name: "GroupIdentifier", Flag: "group-identifier", Type: "*string", Required: true},
	{Name: "MetricName", Flag: "metric-name", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Period", Flag: "period", Type: "*int32", Required: true},
	{Name: "RequestId", Flag: "request-id", Type: "*string", Required: false},
	{Name: "Statistic", Flag: "statistic", Type: "types.CloudWatchAlarmTemplateStatistic", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TargetResourceType", Flag: "target-resource-type", Type: "types.CloudWatchAlarmTemplateTargetResourceType", Required: true},
	{Name: "Threshold", Flag: "threshold", Type: "*float64", Required: true},
	{Name: "TreatMissingData", Flag: "treat-missing-data", Type: "types.CloudWatchAlarmTemplateTreatMissingData", Required: true},
}

var fields_create_cloud_watch_alarm_template_group = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RequestId", Flag: "request-id", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_cluster = []leanruntime.Field{
	{Name: "ClusterType", Flag: "cluster-type", Type: "types.ClusterType", Required: false},
	{Name: "InstanceRoleArn", Flag: "instance-role-arn", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "NetworkSettings", Flag: "network-settings", Type: "*types.ClusterNetworkSettingsCreateRequest", Required: false},
	{Name: "RequestId", Flag: "request-id", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_event_bridge_rule_template = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EventTargets", Flag: "event-targets", Type: "[]types.EventBridgeRuleTemplateTarget", Required: false},
	{Name: "EventType", Flag: "event-type", Type: "types.EventBridgeRuleTemplateEventType", Required: true},
	{Name: "GroupIdentifier", Flag: "group-identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RequestId", Flag: "request-id", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_event_bridge_rule_template_group = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RequestId", Flag: "request-id", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_input = []leanruntime.Field{
	{Name: "Destinations", Flag: "destinations", Type: "[]types.InputDestinationRequest", Required: false},
	{Name: "InputDevices", Flag: "input-devices", Type: "[]types.InputDeviceSettings", Required: false},
	{Name: "InputNetworkLocation", Flag: "input-network-location", Type: "types.InputNetworkLocation", Required: false},
	{Name: "InputSecurityGroups", Flag: "input-security-groups", Type: "[]string", Required: false},
	{Name: "MediaConnectFlows", Flag: "media-connect-flows", Type: "[]types.MediaConnectFlowRequest", Required: false},
	{Name: "MulticastSettings", Flag: "multicast-settings", Type: "*types.MulticastSettingsCreateRequest", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "RequestId", Flag: "request-id", Type: "*string", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "RouterSettings", Flag: "router-settings", Type: "*types.RouterSettings", Required: false},
	{Name: "SdiSources", Flag: "sdi-sources", Type: "[]string", Required: false},
	{Name: "Smpte2110ReceiverGroupSettings", Flag: "smpte2110-receiver-group-settings", Type: "*types.Smpte2110ReceiverGroupSettings", Required: false},
	{Name: "Sources", Flag: "sources", Type: "[]types.InputSourceRequest", Required: false},
	{Name: "SrtSettings", Flag: "srt-settings", Type: "*types.SrtSettingsRequest", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.InputType", Required: false},
	{Name: "Vpc", Flag: "vpc", Type: "*types.InputVpcRequest", Required: false},
}

var fields_create_input_security_group = []leanruntime.Field{
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "WhitelistRules", Flag: "whitelist-rules", Type: "[]types.InputWhitelistRuleCidr", Required: false},
}

var fields_create_multiplex = []leanruntime.Field{
	{Name: "AvailabilityZones", Flag: "availability-zones", Type: "[]string", Required: true},
	{Name: "MultiplexSettings", Flag: "multiplex-settings", Type: "*types.MultiplexSettings", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RequestId", Flag: "request-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_multiplex_program = []leanruntime.Field{
	{Name: "MultiplexId", Flag: "multiplex-id", Type: "*string", Required: true},
	{Name: "MultiplexProgramSettings", Flag: "multiplex-program-settings", Type: "*types.MultiplexProgramSettings", Required: true},
	{Name: "ProgramName", Flag: "program-name", Type: "*string", Required: true},
	{Name: "RequestId", Flag: "request-id", Type: "*string", Required: true},
}

var fields_create_network = []leanruntime.Field{
	{Name: "IpPools", Flag: "ip-pools", Type: "[]types.IpPoolCreateRequest", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "RequestId", Flag: "request-id", Type: "*string", Required: false},
	{Name: "Routes", Flag: "routes", Type: "[]types.RouteCreateRequest", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_node = []leanruntime.Field{
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "NodeInterfaceMappings", Flag: "node-interface-mappings", Type: "[]types.NodeInterfaceMappingCreateRequest", Required: false},
	{Name: "RequestId", Flag: "request-id", Type: "*string", Required: false},
	{Name: "Role", Flag: "role", Type: "types.NodeRole", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_node_registration_script = []leanruntime.Field{
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "NodeInterfaceMappings", Flag: "node-interface-mappings", Type: "[]types.NodeInterfaceMapping", Required: false},
	{Name: "RequestId", Flag: "request-id", Type: "*string", Required: false},
	{Name: "Role", Flag: "role", Type: "types.NodeRole", Required: false},
}

var fields_create_partner_input = []leanruntime.Field{
	{Name: "InputId", Flag: "input-id", Type: "*string", Required: true},
	{Name: "RequestId", Flag: "request-id", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_sdi_source = []leanruntime.Field{
	{Name: "Mode", Flag: "mode", Type: "types.SdiSourceMode", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "RequestId", Flag: "request-id", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.SdiSourceType", Required: false},
}

var fields_create_signal_map = []leanruntime.Field{
	{Name: "CloudWatchAlarmTemplateGroupIdentifiers", Flag: "cloud-watch-alarm-template-group-identifiers", Type: "[]string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DiscoveryEntryPointArn", Flag: "discovery-entry-point-arn", Type: "*string", Required: true},
	{Name: "EventBridgeRuleTemplateGroupIdentifiers", Flag: "event-bridge-rule-template-group-identifiers", Type: "[]string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RequestId", Flag: "request-id", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_tags = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_channel = []leanruntime.Field{
	{Name: "ChannelId", Flag: "channel-id", Type: "*string", Required: true},
}

var fields_delete_channel_placement_group = []leanruntime.Field{
	{Name: "ChannelPlacementGroupId", Flag: "channel-placement-group-id", Type: "*string", Required: true},
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: true},
}

var fields_delete_cloud_watch_alarm_template = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_delete_cloud_watch_alarm_template_group = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_delete_cluster = []leanruntime.Field{
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: true},
}

var fields_delete_event_bridge_rule_template = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_delete_event_bridge_rule_template_group = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_delete_input = []leanruntime.Field{
	{Name: "InputId", Flag: "input-id", Type: "*string", Required: true},
}

var fields_delete_input_security_group = []leanruntime.Field{
	{Name: "InputSecurityGroupId", Flag: "input-security-group-id", Type: "*string", Required: true},
}

var fields_delete_multiplex = []leanruntime.Field{
	{Name: "MultiplexId", Flag: "multiplex-id", Type: "*string", Required: true},
}

var fields_delete_multiplex_program = []leanruntime.Field{
	{Name: "MultiplexId", Flag: "multiplex-id", Type: "*string", Required: true},
	{Name: "ProgramName", Flag: "program-name", Type: "*string", Required: true},
}

var fields_delete_network = []leanruntime.Field{
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
}

var fields_delete_node = []leanruntime.Field{
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: true},
	{Name: "NodeId", Flag: "node-id", Type: "*string", Required: true},
}

var fields_delete_reservation = []leanruntime.Field{
	{Name: "ReservationId", Flag: "reservation-id", Type: "*string", Required: true},
}

var fields_delete_schedule = []leanruntime.Field{
	{Name: "ChannelId", Flag: "channel-id", Type: "*string", Required: true},
}

var fields_delete_sdi_source = []leanruntime.Field{
	{Name: "SdiSourceId", Flag: "sdi-source-id", Type: "*string", Required: true},
}

var fields_delete_signal_map = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_delete_tags = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_describe_account_configuration = []leanruntime.Field{}

var fields_describe_channel = []leanruntime.Field{
	{Name: "ChannelId", Flag: "channel-id", Type: "*string", Required: true},
}

var fields_describe_channel_placement_group = []leanruntime.Field{
	{Name: "ChannelPlacementGroupId", Flag: "channel-placement-group-id", Type: "*string", Required: true},
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: true},
}

var fields_describe_cluster = []leanruntime.Field{
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: true},
}

var fields_describe_input = []leanruntime.Field{
	{Name: "InputId", Flag: "input-id", Type: "*string", Required: true},
}

var fields_describe_input_device = []leanruntime.Field{
	{Name: "InputDeviceId", Flag: "input-device-id", Type: "*string", Required: true},
}

var fields_describe_input_device_thumbnail = []leanruntime.Field{
	{Name: "Accept", Flag: "accept", Type: "types.AcceptHeader", Required: true},
	{Name: "InputDeviceId", Flag: "input-device-id", Type: "*string", Required: true},
}

var fields_describe_input_security_group = []leanruntime.Field{
	{Name: "InputSecurityGroupId", Flag: "input-security-group-id", Type: "*string", Required: true},
}

var fields_describe_multiplex = []leanruntime.Field{
	{Name: "MultiplexId", Flag: "multiplex-id", Type: "*string", Required: true},
}

var fields_describe_multiplex_program = []leanruntime.Field{
	{Name: "MultiplexId", Flag: "multiplex-id", Type: "*string", Required: true},
	{Name: "ProgramName", Flag: "program-name", Type: "*string", Required: true},
}

var fields_describe_network = []leanruntime.Field{
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
}

var fields_describe_node = []leanruntime.Field{
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: true},
	{Name: "NodeId", Flag: "node-id", Type: "*string", Required: true},
}

var fields_describe_offering = []leanruntime.Field{
	{Name: "OfferingId", Flag: "offering-id", Type: "*string", Required: true},
}

var fields_describe_reservation = []leanruntime.Field{
	{Name: "ReservationId", Flag: "reservation-id", Type: "*string", Required: true},
}

var fields_describe_schedule = []leanruntime.Field{
	{Name: "ChannelId", Flag: "channel-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_sdi_source = []leanruntime.Field{
	{Name: "SdiSourceId", Flag: "sdi-source-id", Type: "*string", Required: true},
}

var fields_describe_thumbnails = []leanruntime.Field{
	{Name: "ChannelId", Flag: "channel-id", Type: "*string", Required: true},
	{Name: "PipelineId", Flag: "pipeline-id", Type: "*string", Required: true},
	{Name: "ThumbnailType", Flag: "thumbnail-type", Type: "*string", Required: true},
}

var fields_get_cloud_watch_alarm_template = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_cloud_watch_alarm_template_group = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_event_bridge_rule_template = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_event_bridge_rule_template_group = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_signal_map = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_list_alerts = []leanruntime.Field{
	{Name: "ChannelId", Flag: "channel-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StateFilter", Flag: "state-filter", Type: "*string", Required: false},
}

var fields_list_channel_placement_groups = []leanruntime.Field{
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_channels = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_cloud_watch_alarm_template_groups = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Scope", Flag: "scope", Type: "*string", Required: false},
	{Name: "SignalMapIdentifier", Flag: "signal-map-identifier", Type: "*string", Required: false},
}

var fields_list_cloud_watch_alarm_templates = []leanruntime.Field{
	{Name: "GroupIdentifier", Flag: "group-identifier", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Scope", Flag: "scope", Type: "*string", Required: false},
	{Name: "SignalMapIdentifier", Flag: "signal-map-identifier", Type: "*string", Required: false},
}

var fields_list_cluster_alerts = []leanruntime.Field{
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StateFilter", Flag: "state-filter", Type: "*string", Required: false},
}

var fields_list_clusters = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_event_bridge_rule_template_groups = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SignalMapIdentifier", Flag: "signal-map-identifier", Type: "*string", Required: false},
}

var fields_list_event_bridge_rule_templates = []leanruntime.Field{
	{Name: "GroupIdentifier", Flag: "group-identifier", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SignalMapIdentifier", Flag: "signal-map-identifier", Type: "*string", Required: false},
}

var fields_list_input_device_transfers = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TransferType", Flag: "transfer-type", Type: "*string", Required: true},
}

var fields_list_input_devices = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_input_security_groups = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_inputs = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_multiplex_alerts = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MultiplexId", Flag: "multiplex-id", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StateFilter", Flag: "state-filter", Type: "*string", Required: false},
}

var fields_list_multiplex_programs = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MultiplexId", Flag: "multiplex-id", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_multiplexes = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_networks = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_nodes = []leanruntime.Field{
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_offerings = []leanruntime.Field{
	{Name: "ChannelClass", Flag: "channel-class", Type: "*string", Required: false},
	{Name: "ChannelConfiguration", Flag: "channel-configuration", Type: "*string", Required: false},
	{Name: "Codec", Flag: "codec", Type: "*string", Required: false},
	{Name: "Duration", Flag: "duration", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MaximumBitrate", Flag: "maximum-bitrate", Type: "*string", Required: false},
	{Name: "MaximumFramerate", Flag: "maximum-framerate", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Resolution", Flag: "resolution", Type: "*string", Required: false},
	{Name: "ResourceType", Flag: "resource-type", Type: "*string", Required: false},
	{Name: "SpecialFeature", Flag: "special-feature", Type: "*string", Required: false},
	{Name: "VideoQuality", Flag: "video-quality", Type: "*string", Required: false},
}

var fields_list_reservations = []leanruntime.Field{
	{Name: "ChannelClass", Flag: "channel-class", Type: "*string", Required: false},
	{Name: "Codec", Flag: "codec", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MaximumBitrate", Flag: "maximum-bitrate", Type: "*string", Required: false},
	{Name: "MaximumFramerate", Flag: "maximum-framerate", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Resolution", Flag: "resolution", Type: "*string", Required: false},
	{Name: "ResourceType", Flag: "resource-type", Type: "*string", Required: false},
	{Name: "SpecialFeature", Flag: "special-feature", Type: "*string", Required: false},
	{Name: "VideoQuality", Flag: "video-quality", Type: "*string", Required: false},
}

var fields_list_sdi_sources = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_signal_maps = []leanruntime.Field{
	{Name: "CloudWatchAlarmTemplateGroupIdentifier", Flag: "cloud-watch-alarm-template-group-identifier", Type: "*string", Required: false},
	{Name: "EventBridgeRuleTemplateGroupIdentifier", Flag: "event-bridge-rule-template-group-identifier", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_versions = []leanruntime.Field{}

var fields_purchase_offering = []leanruntime.Field{
	{Name: "Count", Flag: "count", Type: "*int32", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "OfferingId", Flag: "offering-id", Type: "*string", Required: true},
	{Name: "RenewalSettings", Flag: "renewal-settings", Type: "*types.RenewalSettings", Required: false},
	{Name: "RequestId", Flag: "request-id", Type: "*string", Required: false},
	{Name: "Start", Flag: "start", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_reboot_input_device = []leanruntime.Field{
	{Name: "Force", Flag: "force", Type: "types.RebootInputDeviceForce", Required: false},
	{Name: "InputDeviceId", Flag: "input-device-id", Type: "*string", Required: true},
}

var fields_reject_input_device_transfer = []leanruntime.Field{
	{Name: "InputDeviceId", Flag: "input-device-id", Type: "*string", Required: true},
}

var fields_restart_channel_pipelines = []leanruntime.Field{
	{Name: "ChannelId", Flag: "channel-id", Type: "*string", Required: true},
	{Name: "PipelineIds", Flag: "pipeline-ids", Type: "[]types.ChannelPipelineIdToRestart", Required: false},
}

var fields_start_channel = []leanruntime.Field{
	{Name: "ChannelId", Flag: "channel-id", Type: "*string", Required: true},
}

var fields_start_delete_monitor_deployment = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_start_input_device = []leanruntime.Field{
	{Name: "InputDeviceId", Flag: "input-device-id", Type: "*string", Required: true},
}

var fields_start_input_device_maintenance_window = []leanruntime.Field{
	{Name: "InputDeviceId", Flag: "input-device-id", Type: "*string", Required: true},
}

var fields_start_monitor_deployment = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_start_multiplex = []leanruntime.Field{
	{Name: "MultiplexId", Flag: "multiplex-id", Type: "*string", Required: true},
}

var fields_start_update_signal_map = []leanruntime.Field{
	{Name: "CloudWatchAlarmTemplateGroupIdentifiers", Flag: "cloud-watch-alarm-template-group-identifiers", Type: "[]string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DiscoveryEntryPointArn", Flag: "discovery-entry-point-arn", Type: "*string", Required: false},
	{Name: "EventBridgeRuleTemplateGroupIdentifiers", Flag: "event-bridge-rule-template-group-identifiers", Type: "[]string", Required: false},
	{Name: "ForceRediscovery", Flag: "force-rediscovery", Type: "*bool", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_stop_channel = []leanruntime.Field{
	{Name: "ChannelId", Flag: "channel-id", Type: "*string", Required: true},
}

var fields_stop_input_device = []leanruntime.Field{
	{Name: "InputDeviceId", Flag: "input-device-id", Type: "*string", Required: true},
}

var fields_stop_multiplex = []leanruntime.Field{
	{Name: "MultiplexId", Flag: "multiplex-id", Type: "*string", Required: true},
}

var fields_transfer_input_device = []leanruntime.Field{
	{Name: "InputDeviceId", Flag: "input-device-id", Type: "*string", Required: true},
	{Name: "TargetCustomerId", Flag: "target-customer-id", Type: "*string", Required: false},
	{Name: "TargetRegion", Flag: "target-region", Type: "*string", Required: false},
	{Name: "TransferMessage", Flag: "transfer-message", Type: "*string", Required: false},
}

var fields_update_account_configuration = []leanruntime.Field{
	{Name: "AccountConfiguration", Flag: "account-configuration", Type: "*types.AccountConfiguration", Required: false},
}

var fields_update_channel = []leanruntime.Field{
	{Name: "AnywhereSettings", Flag: "anywhere-settings", Type: "*types.AnywhereSettings", Required: false},
	{Name: "CdiInputSpecification", Flag: "cdi-input-specification", Type: "*types.CdiInputSpecification", Required: false},
	{Name: "ChannelEngineVersion", Flag: "channel-engine-version", Type: "*types.ChannelEngineVersionRequest", Required: false},
	{Name: "ChannelId", Flag: "channel-id", Type: "*string", Required: true},
	{Name: "ChannelSecurityGroups", Flag: "channel-security-groups", Type: "[]string", Required: false},
	{Name: "Destinations", Flag: "destinations", Type: "[]types.OutputDestination", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "EncoderSettings", Flag: "encoder-settings", Type: "*types.EncoderSettings", Required: false},
	{Name: "InferenceSettings", Flag: "inference-settings", Type: "*types.InferenceSettings", Required: false},
	{Name: "InputAttachments", Flag: "input-attachments", Type: "[]types.InputAttachment", Required: false},
	{Name: "InputSpecification", Flag: "input-specification", Type: "*types.InputSpecification", Required: false},
	{Name: "LinkedChannelSettings", Flag: "linked-channel-settings", Type: "*types.LinkedChannelSettings", Required: false},
	{Name: "LogLevel", Flag: "log-level", Type: "types.LogLevel", Required: false},
	{Name: "Maintenance", Flag: "maintenance", Type: "*types.MaintenanceUpdateSettings", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
}

var fields_update_channel_class = []leanruntime.Field{
	{Name: "ChannelClass", Flag: "channel-class", Type: "types.ChannelClass", Required: true},
	{Name: "ChannelId", Flag: "channel-id", Type: "*string", Required: true},
	{Name: "Destinations", Flag: "destinations", Type: "[]types.OutputDestination", Required: false},
}

var fields_update_channel_placement_group = []leanruntime.Field{
	{Name: "ChannelPlacementGroupId", Flag: "channel-placement-group-id", Type: "*string", Required: true},
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Nodes", Flag: "nodes", Type: "[]string", Required: false},
}

var fields_update_cloud_watch_alarm_template = []leanruntime.Field{
	{Name: "ComparisonOperator", Flag: "comparison-operator", Type: "types.CloudWatchAlarmTemplateComparisonOperator", Required: false},
	{Name: "DatapointsToAlarm", Flag: "datapoints-to-alarm", Type: "*int32", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EvaluationPeriods", Flag: "evaluation-periods", Type: "*int32", Required: false},
	{Name: "GroupIdentifier", Flag: "group-identifier", Type: "*string", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "MetricName", Flag: "metric-name", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Period", Flag: "period", Type: "*int32", Required: false},
	{Name: "Statistic", Flag: "statistic", Type: "types.CloudWatchAlarmTemplateStatistic", Required: false},
	{Name: "TargetResourceType", Flag: "target-resource-type", Type: "types.CloudWatchAlarmTemplateTargetResourceType", Required: false},
	{Name: "Threshold", Flag: "threshold", Type: "*float64", Required: false},
	{Name: "TreatMissingData", Flag: "treat-missing-data", Type: "types.CloudWatchAlarmTemplateTreatMissingData", Required: false},
}

var fields_update_cloud_watch_alarm_template_group = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_update_cluster = []leanruntime.Field{
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "NetworkSettings", Flag: "network-settings", Type: "*types.ClusterNetworkSettingsUpdateRequest", Required: false},
}

var fields_update_event_bridge_rule_template = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EventTargets", Flag: "event-targets", Type: "[]types.EventBridgeRuleTemplateTarget", Required: false},
	{Name: "EventType", Flag: "event-type", Type: "types.EventBridgeRuleTemplateEventType", Required: false},
	{Name: "GroupIdentifier", Flag: "group-identifier", Type: "*string", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_event_bridge_rule_template_group = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_update_input = []leanruntime.Field{
	{Name: "Destinations", Flag: "destinations", Type: "[]types.InputDestinationRequest", Required: false},
	{Name: "InputDevices", Flag: "input-devices", Type: "[]types.InputDeviceRequest", Required: false},
	{Name: "InputId", Flag: "input-id", Type: "*string", Required: true},
	{Name: "InputSecurityGroups", Flag: "input-security-groups", Type: "[]string", Required: false},
	{Name: "MediaConnectFlows", Flag: "media-connect-flows", Type: "[]types.MediaConnectFlowRequest", Required: false},
	{Name: "MulticastSettings", Flag: "multicast-settings", Type: "*types.MulticastSettingsUpdateRequest", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "SdiSources", Flag: "sdi-sources", Type: "[]string", Required: false},
	{Name: "Smpte2110ReceiverGroupSettings", Flag: "smpte2110-receiver-group-settings", Type: "*types.Smpte2110ReceiverGroupSettings", Required: false},
	{Name: "Sources", Flag: "sources", Type: "[]types.InputSourceRequest", Required: false},
	{Name: "SpecialRouterSettings", Flag: "special-router-settings", Type: "*types.SpecialRouterSettings", Required: false},
	{Name: "SrtSettings", Flag: "srt-settings", Type: "*types.SrtSettingsRequest", Required: false},
}

var fields_update_input_device = []leanruntime.Field{
	{Name: "AvailabilityZone", Flag: "availability-zone", Type: "*string", Required: false},
	{Name: "HdDeviceSettings", Flag: "hd-device-settings", Type: "*types.InputDeviceConfigurableSettings", Required: false},
	{Name: "InputDeviceId", Flag: "input-device-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "UhdDeviceSettings", Flag: "uhd-device-settings", Type: "*types.InputDeviceConfigurableSettings", Required: false},
}

var fields_update_input_security_group = []leanruntime.Field{
	{Name: "InputSecurityGroupId", Flag: "input-security-group-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "WhitelistRules", Flag: "whitelist-rules", Type: "[]types.InputWhitelistRuleCidr", Required: false},
}

var fields_update_multiplex = []leanruntime.Field{
	{Name: "MultiplexId", Flag: "multiplex-id", Type: "*string", Required: true},
	{Name: "MultiplexSettings", Flag: "multiplex-settings", Type: "*types.MultiplexSettings", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "PacketIdentifiersMapping", Flag: "packet-identifiers-mapping", Type: "map[string]types.MultiplexProgramPacketIdentifiersMap", Required: false},
}

var fields_update_multiplex_program = []leanruntime.Field{
	{Name: "MultiplexId", Flag: "multiplex-id", Type: "*string", Required: true},
	{Name: "MultiplexProgramSettings", Flag: "multiplex-program-settings", Type: "*types.MultiplexProgramSettings", Required: false},
	{Name: "ProgramName", Flag: "program-name", Type: "*string", Required: true},
}

var fields_update_network = []leanruntime.Field{
	{Name: "IpPools", Flag: "ip-pools", Type: "[]types.IpPoolUpdateRequest", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
	{Name: "Routes", Flag: "routes", Type: "[]types.RouteUpdateRequest", Required: false},
}

var fields_update_node = []leanruntime.Field{
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "NodeId", Flag: "node-id", Type: "*string", Required: true},
	{Name: "Role", Flag: "role", Type: "types.NodeRole", Required: false},
	{Name: "SdiSourceMappings", Flag: "sdi-source-mappings", Type: "[]types.SdiSourceMappingUpdateRequest", Required: false},
}

var fields_update_node_state = []leanruntime.Field{
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: true},
	{Name: "NodeId", Flag: "node-id", Type: "*string", Required: true},
	{Name: "State", Flag: "state", Type: "types.UpdateNodeStateShape", Required: false},
}

var fields_update_reservation = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "RenewalSettings", Flag: "renewal-settings", Type: "*types.RenewalSettings", Required: false},
	{Name: "ReservationId", Flag: "reservation-id", Type: "*string", Required: true},
}

var fields_update_sdi_source = []leanruntime.Field{
	{Name: "Mode", Flag: "mode", Type: "types.SdiSourceMode", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "SdiSourceId", Flag: "sdi-source-id", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.SdiSourceType", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"accept-input-device-transfer": {
			Name:   "accept-input-device-transfer",
			Fields: fields_accept_input_device_transfer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AcceptInputDeviceTransferInput{}
				if _, err := leanruntime.ApplyInput(input, fields_accept_input_device_transfer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AcceptInputDeviceTransfer(ctx, input)
			},
		},
		"batch-delete": {
			Name:   "batch-delete",
			Fields: fields_batch_delete,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDeleteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_delete, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDelete(ctx, input)
			},
		},
		"batch-start": {
			Name:   "batch-start",
			Fields: fields_batch_start,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchStartInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_start, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchStart(ctx, input)
			},
		},
		"batch-stop": {
			Name:   "batch-stop",
			Fields: fields_batch_stop,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchStopInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_stop, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchStop(ctx, input)
			},
		},
		"batch-update-schedule": {
			Name:   "batch-update-schedule",
			Fields: fields_batch_update_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchUpdateScheduleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_update_schedule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchUpdateSchedule(ctx, input)
			},
		},
		"cancel-input-device-transfer": {
			Name:   "cancel-input-device-transfer",
			Fields: fields_cancel_input_device_transfer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelInputDeviceTransferInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_input_device_transfer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelInputDeviceTransfer(ctx, input)
			},
		},
		"claim-device": {
			Name:   "claim-device",
			Fields: fields_claim_device,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ClaimDeviceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_claim_device, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ClaimDevice(ctx, input)
			},
		},
		"create-channel": {
			Name:   "create-channel",
			Fields: fields_create_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateChannel(ctx, input)
			},
		},
		"create-channel-placement-group": {
			Name:   "create-channel-placement-group",
			Fields: fields_create_channel_placement_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateChannelPlacementGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_channel_placement_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateChannelPlacementGroup(ctx, input)
			},
		},
		"create-cloud-watch-alarm-template": {
			Name:   "create-cloud-watch-alarm-template",
			Fields: fields_create_cloud_watch_alarm_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCloudWatchAlarmTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_cloud_watch_alarm_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCloudWatchAlarmTemplate(ctx, input)
			},
		},
		"create-cloud-watch-alarm-template-group": {
			Name:   "create-cloud-watch-alarm-template-group",
			Fields: fields_create_cloud_watch_alarm_template_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCloudWatchAlarmTemplateGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_cloud_watch_alarm_template_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCloudWatchAlarmTemplateGroup(ctx, input)
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
		"create-event-bridge-rule-template": {
			Name:   "create-event-bridge-rule-template",
			Fields: fields_create_event_bridge_rule_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEventBridgeRuleTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_event_bridge_rule_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEventBridgeRuleTemplate(ctx, input)
			},
		},
		"create-event-bridge-rule-template-group": {
			Name:   "create-event-bridge-rule-template-group",
			Fields: fields_create_event_bridge_rule_template_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEventBridgeRuleTemplateGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_event_bridge_rule_template_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEventBridgeRuleTemplateGroup(ctx, input)
			},
		},
		"create-input": {
			Name:   "create-input",
			Fields: fields_create_input,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateInputInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_input, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateInput(ctx, input)
			},
		},
		"create-input-security-group": {
			Name:   "create-input-security-group",
			Fields: fields_create_input_security_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateInputSecurityGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_input_security_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateInputSecurityGroup(ctx, input)
			},
		},
		"create-multiplex": {
			Name:   "create-multiplex",
			Fields: fields_create_multiplex,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMultiplexInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_multiplex, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMultiplex(ctx, input)
			},
		},
		"create-multiplex-program": {
			Name:   "create-multiplex-program",
			Fields: fields_create_multiplex_program,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMultiplexProgramInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_multiplex_program, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMultiplexProgram(ctx, input)
			},
		},
		"create-network": {
			Name:   "create-network",
			Fields: fields_create_network,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateNetworkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_network, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateNetwork(ctx, input)
			},
		},
		"create-node": {
			Name:   "create-node",
			Fields: fields_create_node,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateNodeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_node, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateNode(ctx, input)
			},
		},
		"create-node-registration-script": {
			Name:   "create-node-registration-script",
			Fields: fields_create_node_registration_script,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateNodeRegistrationScriptInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_node_registration_script, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateNodeRegistrationScript(ctx, input)
			},
		},
		"create-partner-input": {
			Name:   "create-partner-input",
			Fields: fields_create_partner_input,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePartnerInputInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_partner_input, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePartnerInput(ctx, input)
			},
		},
		"create-sdi-source": {
			Name:   "create-sdi-source",
			Fields: fields_create_sdi_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSdiSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_sdi_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSdiSource(ctx, input)
			},
		},
		"create-signal-map": {
			Name:   "create-signal-map",
			Fields: fields_create_signal_map,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSignalMapInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_signal_map, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSignalMap(ctx, input)
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
		"delete-channel": {
			Name:   "delete-channel",
			Fields: fields_delete_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteChannel(ctx, input)
			},
		},
		"delete-channel-placement-group": {
			Name:   "delete-channel-placement-group",
			Fields: fields_delete_channel_placement_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteChannelPlacementGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_channel_placement_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteChannelPlacementGroup(ctx, input)
			},
		},
		"delete-cloud-watch-alarm-template": {
			Name:   "delete-cloud-watch-alarm-template",
			Fields: fields_delete_cloud_watch_alarm_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCloudWatchAlarmTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_cloud_watch_alarm_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCloudWatchAlarmTemplate(ctx, input)
			},
		},
		"delete-cloud-watch-alarm-template-group": {
			Name:   "delete-cloud-watch-alarm-template-group",
			Fields: fields_delete_cloud_watch_alarm_template_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCloudWatchAlarmTemplateGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_cloud_watch_alarm_template_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCloudWatchAlarmTemplateGroup(ctx, input)
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
		"delete-event-bridge-rule-template": {
			Name:   "delete-event-bridge-rule-template",
			Fields: fields_delete_event_bridge_rule_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEventBridgeRuleTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_event_bridge_rule_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEventBridgeRuleTemplate(ctx, input)
			},
		},
		"delete-event-bridge-rule-template-group": {
			Name:   "delete-event-bridge-rule-template-group",
			Fields: fields_delete_event_bridge_rule_template_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEventBridgeRuleTemplateGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_event_bridge_rule_template_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEventBridgeRuleTemplateGroup(ctx, input)
			},
		},
		"delete-input": {
			Name:   "delete-input",
			Fields: fields_delete_input,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteInputInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_input, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteInput(ctx, input)
			},
		},
		"delete-input-security-group": {
			Name:   "delete-input-security-group",
			Fields: fields_delete_input_security_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteInputSecurityGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_input_security_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteInputSecurityGroup(ctx, input)
			},
		},
		"delete-multiplex": {
			Name:   "delete-multiplex",
			Fields: fields_delete_multiplex,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMultiplexInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_multiplex, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMultiplex(ctx, input)
			},
		},
		"delete-multiplex-program": {
			Name:   "delete-multiplex-program",
			Fields: fields_delete_multiplex_program,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMultiplexProgramInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_multiplex_program, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMultiplexProgram(ctx, input)
			},
		},
		"delete-network": {
			Name:   "delete-network",
			Fields: fields_delete_network,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteNetworkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_network, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteNetwork(ctx, input)
			},
		},
		"delete-node": {
			Name:   "delete-node",
			Fields: fields_delete_node,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteNodeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_node, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteNode(ctx, input)
			},
		},
		"delete-reservation": {
			Name:   "delete-reservation",
			Fields: fields_delete_reservation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteReservationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_reservation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteReservation(ctx, input)
			},
		},
		"delete-schedule": {
			Name:   "delete-schedule",
			Fields: fields_delete_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteScheduleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_schedule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSchedule(ctx, input)
			},
		},
		"delete-sdi-source": {
			Name:   "delete-sdi-source",
			Fields: fields_delete_sdi_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSdiSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_sdi_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSdiSource(ctx, input)
			},
		},
		"delete-signal-map": {
			Name:   "delete-signal-map",
			Fields: fields_delete_signal_map,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSignalMapInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_signal_map, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSignalMap(ctx, input)
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
		"describe-account-configuration": {
			Name:   "describe-account-configuration",
			Fields: fields_describe_account_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAccountConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_account_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAccountConfiguration(ctx, input)
			},
		},
		"describe-channel": {
			Name:   "describe-channel",
			Fields: fields_describe_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeChannel(ctx, input)
			},
		},
		"describe-channel-placement-group": {
			Name:   "describe-channel-placement-group",
			Fields: fields_describe_channel_placement_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeChannelPlacementGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_channel_placement_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeChannelPlacementGroup(ctx, input)
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
		"describe-input": {
			Name:   "describe-input",
			Fields: fields_describe_input,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInputInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_input, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeInput(ctx, input)
			},
		},
		"describe-input-device": {
			Name:   "describe-input-device",
			Fields: fields_describe_input_device,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInputDeviceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_input_device, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeInputDevice(ctx, input)
			},
		},
		"describe-input-device-thumbnail": {
			Name:   "describe-input-device-thumbnail",
			Fields: fields_describe_input_device_thumbnail,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInputDeviceThumbnailInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_input_device_thumbnail, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeInputDeviceThumbnail(ctx, input)
			},
		},
		"describe-input-security-group": {
			Name:   "describe-input-security-group",
			Fields: fields_describe_input_security_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInputSecurityGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_input_security_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeInputSecurityGroup(ctx, input)
			},
		},
		"describe-multiplex": {
			Name:   "describe-multiplex",
			Fields: fields_describe_multiplex,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMultiplexInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_multiplex, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeMultiplex(ctx, input)
			},
		},
		"describe-multiplex-program": {
			Name:   "describe-multiplex-program",
			Fields: fields_describe_multiplex_program,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMultiplexProgramInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_multiplex_program, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeMultiplexProgram(ctx, input)
			},
		},
		"describe-network": {
			Name:   "describe-network",
			Fields: fields_describe_network,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeNetworkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_network, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeNetwork(ctx, input)
			},
		},
		"describe-node": {
			Name:   "describe-node",
			Fields: fields_describe_node,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeNodeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_node, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeNode(ctx, input)
			},
		},
		"describe-offering": {
			Name:   "describe-offering",
			Fields: fields_describe_offering,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeOfferingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_offering, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeOffering(ctx, input)
			},
		},
		"describe-reservation": {
			Name:   "describe-reservation",
			Fields: fields_describe_reservation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeReservationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_reservation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeReservation(ctx, input)
			},
		},
		"describe-schedule": {
			Name:   "describe-schedule",
			Fields: fields_describe_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeScheduleInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_schedule, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeSchedule(ctx, input)
				}
				var results []*svc.DescribeScheduleOutput
				p := svc.NewDescribeSchedulePaginator(client, input)
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
		"describe-sdi-source": {
			Name:   "describe-sdi-source",
			Fields: fields_describe_sdi_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSdiSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_sdi_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSdiSource(ctx, input)
			},
		},
		"describe-thumbnails": {
			Name:   "describe-thumbnails",
			Fields: fields_describe_thumbnails,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeThumbnailsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_thumbnails, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeThumbnails(ctx, input)
			},
		},
		"get-cloud-watch-alarm-template": {
			Name:   "get-cloud-watch-alarm-template",
			Fields: fields_get_cloud_watch_alarm_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCloudWatchAlarmTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_cloud_watch_alarm_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCloudWatchAlarmTemplate(ctx, input)
			},
		},
		"get-cloud-watch-alarm-template-group": {
			Name:   "get-cloud-watch-alarm-template-group",
			Fields: fields_get_cloud_watch_alarm_template_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCloudWatchAlarmTemplateGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_cloud_watch_alarm_template_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCloudWatchAlarmTemplateGroup(ctx, input)
			},
		},
		"get-event-bridge-rule-template": {
			Name:   "get-event-bridge-rule-template",
			Fields: fields_get_event_bridge_rule_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEventBridgeRuleTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_event_bridge_rule_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEventBridgeRuleTemplate(ctx, input)
			},
		},
		"get-event-bridge-rule-template-group": {
			Name:   "get-event-bridge-rule-template-group",
			Fields: fields_get_event_bridge_rule_template_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEventBridgeRuleTemplateGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_event_bridge_rule_template_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEventBridgeRuleTemplateGroup(ctx, input)
			},
		},
		"get-signal-map": {
			Name:   "get-signal-map",
			Fields: fields_get_signal_map,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSignalMapInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_signal_map, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSignalMap(ctx, input)
			},
		},
		"list-alerts": {
			Name:   "list-alerts",
			Fields: fields_list_alerts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAlertsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_alerts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAlerts(ctx, input)
				}
				var results []*svc.ListAlertsOutput
				p := svc.NewListAlertsPaginator(client, input)
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
		"list-channel-placement-groups": {
			Name:   "list-channel-placement-groups",
			Fields: fields_list_channel_placement_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListChannelPlacementGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_channel_placement_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListChannelPlacementGroups(ctx, input)
				}
				var results []*svc.ListChannelPlacementGroupsOutput
				p := svc.NewListChannelPlacementGroupsPaginator(client, input)
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
		"list-channels": {
			Name:   "list-channels",
			Fields: fields_list_channels,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListChannelsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_channels, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListChannels(ctx, input)
				}
				var results []*svc.ListChannelsOutput
				p := svc.NewListChannelsPaginator(client, input)
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
		"list-cloud-watch-alarm-template-groups": {
			Name:   "list-cloud-watch-alarm-template-groups",
			Fields: fields_list_cloud_watch_alarm_template_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCloudWatchAlarmTemplateGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_cloud_watch_alarm_template_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCloudWatchAlarmTemplateGroups(ctx, input)
				}
				var results []*svc.ListCloudWatchAlarmTemplateGroupsOutput
				p := svc.NewListCloudWatchAlarmTemplateGroupsPaginator(client, input)
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
		"list-cloud-watch-alarm-templates": {
			Name:   "list-cloud-watch-alarm-templates",
			Fields: fields_list_cloud_watch_alarm_templates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCloudWatchAlarmTemplatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_cloud_watch_alarm_templates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCloudWatchAlarmTemplates(ctx, input)
				}
				var results []*svc.ListCloudWatchAlarmTemplatesOutput
				p := svc.NewListCloudWatchAlarmTemplatesPaginator(client, input)
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
		"list-cluster-alerts": {
			Name:   "list-cluster-alerts",
			Fields: fields_list_cluster_alerts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListClusterAlertsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_cluster_alerts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListClusterAlerts(ctx, input)
				}
				var results []*svc.ListClusterAlertsOutput
				p := svc.NewListClusterAlertsPaginator(client, input)
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
		"list-event-bridge-rule-template-groups": {
			Name:   "list-event-bridge-rule-template-groups",
			Fields: fields_list_event_bridge_rule_template_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEventBridgeRuleTemplateGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_event_bridge_rule_template_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEventBridgeRuleTemplateGroups(ctx, input)
				}
				var results []*svc.ListEventBridgeRuleTemplateGroupsOutput
				p := svc.NewListEventBridgeRuleTemplateGroupsPaginator(client, input)
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
		"list-event-bridge-rule-templates": {
			Name:   "list-event-bridge-rule-templates",
			Fields: fields_list_event_bridge_rule_templates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEventBridgeRuleTemplatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_event_bridge_rule_templates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEventBridgeRuleTemplates(ctx, input)
				}
				var results []*svc.ListEventBridgeRuleTemplatesOutput
				p := svc.NewListEventBridgeRuleTemplatesPaginator(client, input)
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
		"list-input-device-transfers": {
			Name:   "list-input-device-transfers",
			Fields: fields_list_input_device_transfers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInputDeviceTransfersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_input_device_transfers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListInputDeviceTransfers(ctx, input)
				}
				var results []*svc.ListInputDeviceTransfersOutput
				p := svc.NewListInputDeviceTransfersPaginator(client, input)
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
		"list-input-devices": {
			Name:   "list-input-devices",
			Fields: fields_list_input_devices,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInputDevicesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_input_devices, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListInputDevices(ctx, input)
				}
				var results []*svc.ListInputDevicesOutput
				p := svc.NewListInputDevicesPaginator(client, input)
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
		"list-input-security-groups": {
			Name:   "list-input-security-groups",
			Fields: fields_list_input_security_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInputSecurityGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_input_security_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListInputSecurityGroups(ctx, input)
				}
				var results []*svc.ListInputSecurityGroupsOutput
				p := svc.NewListInputSecurityGroupsPaginator(client, input)
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
		"list-inputs": {
			Name:   "list-inputs",
			Fields: fields_list_inputs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInputsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_inputs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListInputs(ctx, input)
				}
				var results []*svc.ListInputsOutput
				p := svc.NewListInputsPaginator(client, input)
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
		"list-multiplex-alerts": {
			Name:   "list-multiplex-alerts",
			Fields: fields_list_multiplex_alerts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMultiplexAlertsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_multiplex_alerts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMultiplexAlerts(ctx, input)
				}
				var results []*svc.ListMultiplexAlertsOutput
				p := svc.NewListMultiplexAlertsPaginator(client, input)
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
		"list-multiplex-programs": {
			Name:   "list-multiplex-programs",
			Fields: fields_list_multiplex_programs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMultiplexProgramsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_multiplex_programs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMultiplexPrograms(ctx, input)
				}
				var results []*svc.ListMultiplexProgramsOutput
				p := svc.NewListMultiplexProgramsPaginator(client, input)
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
		"list-multiplexes": {
			Name:   "list-multiplexes",
			Fields: fields_list_multiplexes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMultiplexesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_multiplexes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMultiplexes(ctx, input)
				}
				var results []*svc.ListMultiplexesOutput
				p := svc.NewListMultiplexesPaginator(client, input)
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
		"list-networks": {
			Name:   "list-networks",
			Fields: fields_list_networks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListNetworksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_networks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListNetworks(ctx, input)
				}
				var results []*svc.ListNetworksOutput
				p := svc.NewListNetworksPaginator(client, input)
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
		"list-nodes": {
			Name:   "list-nodes",
			Fields: fields_list_nodes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListNodesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_nodes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListNodes(ctx, input)
				}
				var results []*svc.ListNodesOutput
				p := svc.NewListNodesPaginator(client, input)
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
		"list-reservations": {
			Name:   "list-reservations",
			Fields: fields_list_reservations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListReservationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_reservations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListReservations(ctx, input)
				}
				var results []*svc.ListReservationsOutput
				p := svc.NewListReservationsPaginator(client, input)
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
		"list-sdi-sources": {
			Name:   "list-sdi-sources",
			Fields: fields_list_sdi_sources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSdiSourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_sdi_sources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSdiSources(ctx, input)
				}
				var results []*svc.ListSdiSourcesOutput
				p := svc.NewListSdiSourcesPaginator(client, input)
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
		"list-signal-maps": {
			Name:   "list-signal-maps",
			Fields: fields_list_signal_maps,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSignalMapsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_signal_maps, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSignalMaps(ctx, input)
				}
				var results []*svc.ListSignalMapsOutput
				p := svc.NewListSignalMapsPaginator(client, input)
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
		"list-versions": {
			Name:   "list-versions",
			Fields: fields_list_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVersionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_versions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListVersions(ctx, input)
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
		"reboot-input-device": {
			Name:   "reboot-input-device",
			Fields: fields_reboot_input_device,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RebootInputDeviceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reboot_input_device, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RebootInputDevice(ctx, input)
			},
		},
		"reject-input-device-transfer": {
			Name:   "reject-input-device-transfer",
			Fields: fields_reject_input_device_transfer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RejectInputDeviceTransferInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reject_input_device_transfer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RejectInputDeviceTransfer(ctx, input)
			},
		},
		"restart-channel-pipelines": {
			Name:   "restart-channel-pipelines",
			Fields: fields_restart_channel_pipelines,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RestartChannelPipelinesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_restart_channel_pipelines, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RestartChannelPipelines(ctx, input)
			},
		},
		"start-channel": {
			Name:   "start-channel",
			Fields: fields_start_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartChannel(ctx, input)
			},
		},
		"start-delete-monitor-deployment": {
			Name:   "start-delete-monitor-deployment",
			Fields: fields_start_delete_monitor_deployment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartDeleteMonitorDeploymentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_delete_monitor_deployment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartDeleteMonitorDeployment(ctx, input)
			},
		},
		"start-input-device": {
			Name:   "start-input-device",
			Fields: fields_start_input_device,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartInputDeviceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_input_device, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartInputDevice(ctx, input)
			},
		},
		"start-input-device-maintenance-window": {
			Name:   "start-input-device-maintenance-window",
			Fields: fields_start_input_device_maintenance_window,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartInputDeviceMaintenanceWindowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_input_device_maintenance_window, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartInputDeviceMaintenanceWindow(ctx, input)
			},
		},
		"start-monitor-deployment": {
			Name:   "start-monitor-deployment",
			Fields: fields_start_monitor_deployment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartMonitorDeploymentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_monitor_deployment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartMonitorDeployment(ctx, input)
			},
		},
		"start-multiplex": {
			Name:   "start-multiplex",
			Fields: fields_start_multiplex,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartMultiplexInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_multiplex, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartMultiplex(ctx, input)
			},
		},
		"start-update-signal-map": {
			Name:   "start-update-signal-map",
			Fields: fields_start_update_signal_map,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartUpdateSignalMapInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_update_signal_map, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartUpdateSignalMap(ctx, input)
			},
		},
		"stop-channel": {
			Name:   "stop-channel",
			Fields: fields_stop_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopChannel(ctx, input)
			},
		},
		"stop-input-device": {
			Name:   "stop-input-device",
			Fields: fields_stop_input_device,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopInputDeviceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_input_device, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopInputDevice(ctx, input)
			},
		},
		"stop-multiplex": {
			Name:   "stop-multiplex",
			Fields: fields_stop_multiplex,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopMultiplexInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_multiplex, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopMultiplex(ctx, input)
			},
		},
		"transfer-input-device": {
			Name:   "transfer-input-device",
			Fields: fields_transfer_input_device,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TransferInputDeviceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_transfer_input_device, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TransferInputDevice(ctx, input)
			},
		},
		"update-account-configuration": {
			Name:   "update-account-configuration",
			Fields: fields_update_account_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAccountConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_account_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAccountConfiguration(ctx, input)
			},
		},
		"update-channel": {
			Name:   "update-channel",
			Fields: fields_update_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateChannel(ctx, input)
			},
		},
		"update-channel-class": {
			Name:   "update-channel-class",
			Fields: fields_update_channel_class,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateChannelClassInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_channel_class, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateChannelClass(ctx, input)
			},
		},
		"update-channel-placement-group": {
			Name:   "update-channel-placement-group",
			Fields: fields_update_channel_placement_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateChannelPlacementGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_channel_placement_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateChannelPlacementGroup(ctx, input)
			},
		},
		"update-cloud-watch-alarm-template": {
			Name:   "update-cloud-watch-alarm-template",
			Fields: fields_update_cloud_watch_alarm_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCloudWatchAlarmTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_cloud_watch_alarm_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCloudWatchAlarmTemplate(ctx, input)
			},
		},
		"update-cloud-watch-alarm-template-group": {
			Name:   "update-cloud-watch-alarm-template-group",
			Fields: fields_update_cloud_watch_alarm_template_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCloudWatchAlarmTemplateGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_cloud_watch_alarm_template_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCloudWatchAlarmTemplateGroup(ctx, input)
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
		"update-event-bridge-rule-template": {
			Name:   "update-event-bridge-rule-template",
			Fields: fields_update_event_bridge_rule_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEventBridgeRuleTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_event_bridge_rule_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEventBridgeRuleTemplate(ctx, input)
			},
		},
		"update-event-bridge-rule-template-group": {
			Name:   "update-event-bridge-rule-template-group",
			Fields: fields_update_event_bridge_rule_template_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEventBridgeRuleTemplateGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_event_bridge_rule_template_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEventBridgeRuleTemplateGroup(ctx, input)
			},
		},
		"update-input": {
			Name:   "update-input",
			Fields: fields_update_input,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateInputInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_input, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateInput(ctx, input)
			},
		},
		"update-input-device": {
			Name:   "update-input-device",
			Fields: fields_update_input_device,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateInputDeviceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_input_device, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateInputDevice(ctx, input)
			},
		},
		"update-input-security-group": {
			Name:   "update-input-security-group",
			Fields: fields_update_input_security_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateInputSecurityGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_input_security_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateInputSecurityGroup(ctx, input)
			},
		},
		"update-multiplex": {
			Name:   "update-multiplex",
			Fields: fields_update_multiplex,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMultiplexInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_multiplex, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMultiplex(ctx, input)
			},
		},
		"update-multiplex-program": {
			Name:   "update-multiplex-program",
			Fields: fields_update_multiplex_program,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMultiplexProgramInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_multiplex_program, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMultiplexProgram(ctx, input)
			},
		},
		"update-network": {
			Name:   "update-network",
			Fields: fields_update_network,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateNetworkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_network, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateNetwork(ctx, input)
			},
		},
		"update-node": {
			Name:   "update-node",
			Fields: fields_update_node,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateNodeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_node, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateNode(ctx, input)
			},
		},
		"update-node-state": {
			Name:   "update-node-state",
			Fields: fields_update_node_state,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateNodeStateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_node_state, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateNodeState(ctx, input)
			},
		},
		"update-reservation": {
			Name:   "update-reservation",
			Fields: fields_update_reservation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateReservationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_reservation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateReservation(ctx, input)
			},
		},
		"update-sdi-source": {
			Name:   "update-sdi-source",
			Fields: fields_update_sdi_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSdiSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_sdi_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSdiSource(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("medialive", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
