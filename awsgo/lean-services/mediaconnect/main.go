package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/mediaconnect"
)

var fields_add_bridge_outputs = []leanruntime.Field{
	{Name: "BridgeArn", Flag: "bridge-arn", Type: "*string", Required: true},
	{Name: "Outputs", Flag: "outputs", Type: "[]types.AddBridgeOutputRequest", Required: true},
}

var fields_add_bridge_sources = []leanruntime.Field{
	{Name: "BridgeArn", Flag: "bridge-arn", Type: "*string", Required: true},
	{Name: "Sources", Flag: "sources", Type: "[]types.AddBridgeSourceRequest", Required: true},
}

var fields_add_flow_media_streams = []leanruntime.Field{
	{Name: "FlowArn", Flag: "flow-arn", Type: "*string", Required: true},
	{Name: "MediaStreams", Flag: "media-streams", Type: "[]types.AddMediaStreamRequest", Required: true},
}

var fields_add_flow_outputs = []leanruntime.Field{
	{Name: "FlowArn", Flag: "flow-arn", Type: "*string", Required: true},
	{Name: "Outputs", Flag: "outputs", Type: "[]types.AddOutputRequest", Required: true},
}

var fields_add_flow_sources = []leanruntime.Field{
	{Name: "FlowArn", Flag: "flow-arn", Type: "*string", Required: true},
	{Name: "Sources", Flag: "sources", Type: "[]types.SetSourceRequest", Required: true},
}

var fields_add_flow_vpc_interfaces = []leanruntime.Field{
	{Name: "FlowArn", Flag: "flow-arn", Type: "*string", Required: true},
	{Name: "VpcInterfaces", Flag: "vpc-interfaces", Type: "[]types.VpcInterfaceRequest", Required: true},
}

var fields_batch_get_router_input = []leanruntime.Field{
	{Name: "Arns", Flag: "arns", Type: "[]string", Required: true},
}

var fields_batch_get_router_network_interface = []leanruntime.Field{
	{Name: "Arns", Flag: "arns", Type: "[]string", Required: true},
}

var fields_batch_get_router_output = []leanruntime.Field{
	{Name: "Arns", Flag: "arns", Type: "[]string", Required: true},
}

var fields_create_bridge = []leanruntime.Field{
	{Name: "EgressGatewayBridge", Flag: "egress-gateway-bridge", Type: "*types.AddEgressGatewayBridgeRequest", Required: false},
	{Name: "IngressGatewayBridge", Flag: "ingress-gateway-bridge", Type: "*types.AddIngressGatewayBridgeRequest", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Outputs", Flag: "outputs", Type: "[]types.AddBridgeOutputRequest", Required: false},
	{Name: "PlacementArn", Flag: "placement-arn", Type: "*string", Required: true},
	{Name: "SourceFailoverConfig", Flag: "source-failover-config", Type: "*types.FailoverConfig", Required: false},
	{Name: "Sources", Flag: "sources", Type: "[]types.AddBridgeSourceRequest", Required: true},
}

var fields_create_flow = []leanruntime.Field{
	{Name: "AvailabilityZone", Flag: "availability-zone", Type: "*string", Required: false},
	{Name: "EncodingConfig", Flag: "encoding-config", Type: "*types.EncodingConfig", Required: false},
	{Name: "Entitlements", Flag: "entitlements", Type: "[]types.GrantEntitlementRequest", Required: false},
	{Name: "FlowSize", Flag: "flow-size", Type: "types.FlowSize", Required: false},
	{Name: "FlowTags", Flag: "flow-tags", Type: "map[string]string", Required: false},
	{Name: "Maintenance", Flag: "maintenance", Type: "*types.AddMaintenance", Required: false},
	{Name: "MediaStreams", Flag: "media-streams", Type: "[]types.AddMediaStreamRequest", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NdiConfig", Flag: "ndi-config", Type: "*types.NdiConfig", Required: false},
	{Name: "Outputs", Flag: "outputs", Type: "[]types.AddOutputRequest", Required: false},
	{Name: "Source", Flag: "source", Type: "*types.SetSourceRequest", Required: false},
	{Name: "SourceFailoverConfig", Flag: "source-failover-config", Type: "*types.FailoverConfig", Required: false},
	{Name: "SourceMonitoringConfig", Flag: "source-monitoring-config", Type: "*types.MonitoringConfig", Required: false},
	{Name: "Sources", Flag: "sources", Type: "[]types.SetSourceRequest", Required: false},
	{Name: "VpcInterfaces", Flag: "vpc-interfaces", Type: "[]types.VpcInterfaceRequest", Required: false},
}

var fields_create_gateway = []leanruntime.Field{
	{Name: "EgressCidrBlocks", Flag: "egress-cidr-blocks", Type: "[]string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Networks", Flag: "networks", Type: "[]types.GatewayNetwork", Required: true},
}

var fields_create_router_input = []leanruntime.Field{
	{Name: "AvailabilityZone", Flag: "availability-zone", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Configuration", Flag: "configuration", Type: "types.RouterInputConfiguration", Required: true},
	{Name: "MaintenanceConfiguration", Flag: "maintenance-configuration", Type: "types.MaintenanceConfiguration", Required: false},
	{Name: "MaximumBitrate", Flag: "maximum-bitrate", Type: "*int64", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RegionName", Flag: "region-name", Type: "*string", Required: false},
	{Name: "RoutingScope", Flag: "routing-scope", Type: "types.RoutingScope", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Tier", Flag: "tier", Type: "types.RouterInputTier", Required: true},
	{Name: "TransitEncryption", Flag: "transit-encryption", Type: "*types.RouterInputTransitEncryption", Required: false},
}

var fields_create_router_network_interface = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Configuration", Flag: "configuration", Type: "types.RouterNetworkInterfaceConfiguration", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RegionName", Flag: "region-name", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_router_output = []leanruntime.Field{
	{Name: "AvailabilityZone", Flag: "availability-zone", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Configuration", Flag: "configuration", Type: "types.RouterOutputConfiguration", Required: true},
	{Name: "MaintenanceConfiguration", Flag: "maintenance-configuration", Type: "types.MaintenanceConfiguration", Required: false},
	{Name: "MaximumBitrate", Flag: "maximum-bitrate", Type: "*int64", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RegionName", Flag: "region-name", Type: "*string", Required: false},
	{Name: "RoutingScope", Flag: "routing-scope", Type: "types.RoutingScope", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Tier", Flag: "tier", Type: "types.RouterOutputTier", Required: true},
}

var fields_delete_bridge = []leanruntime.Field{
	{Name: "BridgeArn", Flag: "bridge-arn", Type: "*string", Required: true},
}

var fields_delete_flow = []leanruntime.Field{
	{Name: "FlowArn", Flag: "flow-arn", Type: "*string", Required: true},
}

var fields_delete_gateway = []leanruntime.Field{
	{Name: "GatewayArn", Flag: "gateway-arn", Type: "*string", Required: true},
}

var fields_delete_router_input = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_delete_router_network_interface = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_delete_router_output = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_deregister_gateway_instance = []leanruntime.Field{
	{Name: "Force", Flag: "force", Type: "*bool", Required: false},
	{Name: "GatewayInstanceArn", Flag: "gateway-instance-arn", Type: "*string", Required: true},
}

var fields_describe_bridge = []leanruntime.Field{
	{Name: "BridgeArn", Flag: "bridge-arn", Type: "*string", Required: true},
}

var fields_describe_flow = []leanruntime.Field{
	{Name: "FlowArn", Flag: "flow-arn", Type: "*string", Required: true},
}

var fields_describe_flow_source_metadata = []leanruntime.Field{
	{Name: "FlowArn", Flag: "flow-arn", Type: "*string", Required: true},
}

var fields_describe_flow_source_thumbnail = []leanruntime.Field{
	{Name: "FlowArn", Flag: "flow-arn", Type: "*string", Required: true},
}

var fields_describe_gateway = []leanruntime.Field{
	{Name: "GatewayArn", Flag: "gateway-arn", Type: "*string", Required: true},
}

var fields_describe_gateway_instance = []leanruntime.Field{
	{Name: "GatewayInstanceArn", Flag: "gateway-instance-arn", Type: "*string", Required: true},
}

var fields_describe_offering = []leanruntime.Field{
	{Name: "OfferingArn", Flag: "offering-arn", Type: "*string", Required: true},
}

var fields_describe_reservation = []leanruntime.Field{
	{Name: "ReservationArn", Flag: "reservation-arn", Type: "*string", Required: true},
}

var fields_get_router_input = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_router_input_source_metadata = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_router_input_thumbnail = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_router_network_interface = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_router_output = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_grant_flow_entitlements = []leanruntime.Field{
	{Name: "Entitlements", Flag: "entitlements", Type: "[]types.GrantEntitlementRequest", Required: true},
	{Name: "FlowArn", Flag: "flow-arn", Type: "*string", Required: true},
}

var fields_list_bridges = []leanruntime.Field{
	{Name: "FilterArn", Flag: "filter-arn", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_entitlements = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_flows = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_gateway_instances = []leanruntime.Field{
	{Name: "FilterArn", Flag: "filter-arn", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_gateways = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_offerings = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_reservations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_router_inputs = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.RouterInputFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_router_network_interfaces = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.RouterNetworkInterfaceFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_router_outputs = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.RouterOutputFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_global_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_purchase_offering = []leanruntime.Field{
	{Name: "OfferingArn", Flag: "offering-arn", Type: "*string", Required: true},
	{Name: "ReservationName", Flag: "reservation-name", Type: "*string", Required: true},
	{Name: "Start", Flag: "start", Type: "*string", Required: true},
}

var fields_remove_bridge_output = []leanruntime.Field{
	{Name: "BridgeArn", Flag: "bridge-arn", Type: "*string", Required: true},
	{Name: "OutputName", Flag: "output-name", Type: "*string", Required: true},
}

var fields_remove_bridge_source = []leanruntime.Field{
	{Name: "BridgeArn", Flag: "bridge-arn", Type: "*string", Required: true},
	{Name: "SourceName", Flag: "source-name", Type: "*string", Required: true},
}

var fields_remove_flow_media_stream = []leanruntime.Field{
	{Name: "FlowArn", Flag: "flow-arn", Type: "*string", Required: true},
	{Name: "MediaStreamName", Flag: "media-stream-name", Type: "*string", Required: true},
}

var fields_remove_flow_output = []leanruntime.Field{
	{Name: "FlowArn", Flag: "flow-arn", Type: "*string", Required: true},
	{Name: "OutputArn", Flag: "output-arn", Type: "*string", Required: true},
}

var fields_remove_flow_source = []leanruntime.Field{
	{Name: "FlowArn", Flag: "flow-arn", Type: "*string", Required: true},
	{Name: "SourceArn", Flag: "source-arn", Type: "*string", Required: true},
}

var fields_remove_flow_vpc_interface = []leanruntime.Field{
	{Name: "FlowArn", Flag: "flow-arn", Type: "*string", Required: true},
	{Name: "VpcInterfaceName", Flag: "vpc-interface-name", Type: "*string", Required: true},
}

var fields_restart_router_input = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_restart_router_output = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_revoke_flow_entitlement = []leanruntime.Field{
	{Name: "EntitlementArn", Flag: "entitlement-arn", Type: "*string", Required: true},
	{Name: "FlowArn", Flag: "flow-arn", Type: "*string", Required: true},
}

var fields_start_flow = []leanruntime.Field{
	{Name: "FlowArn", Flag: "flow-arn", Type: "*string", Required: true},
}

var fields_start_router_input = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_start_router_output = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_stop_flow = []leanruntime.Field{
	{Name: "FlowArn", Flag: "flow-arn", Type: "*string", Required: true},
}

var fields_stop_router_input = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_stop_router_output = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_tag_global_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_take_router_input = []leanruntime.Field{
	{Name: "RouterInputArn", Flag: "router-input-arn", Type: "*string", Required: false},
	{Name: "RouterOutputArn", Flag: "router-output-arn", Type: "*string", Required: true},
}

var fields_untag_global_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_bridge = []leanruntime.Field{
	{Name: "BridgeArn", Flag: "bridge-arn", Type: "*string", Required: true},
	{Name: "EgressGatewayBridge", Flag: "egress-gateway-bridge", Type: "*types.UpdateEgressGatewayBridgeRequest", Required: false},
	{Name: "IngressGatewayBridge", Flag: "ingress-gateway-bridge", Type: "*types.UpdateIngressGatewayBridgeRequest", Required: false},
	{Name: "SourceFailoverConfig", Flag: "source-failover-config", Type: "*types.UpdateFailoverConfig", Required: false},
}

var fields_update_bridge_output = []leanruntime.Field{
	{Name: "BridgeArn", Flag: "bridge-arn", Type: "*string", Required: true},
	{Name: "NetworkOutput", Flag: "network-output", Type: "*types.UpdateBridgeNetworkOutputRequest", Required: false},
	{Name: "OutputName", Flag: "output-name", Type: "*string", Required: true},
}

var fields_update_bridge_source = []leanruntime.Field{
	{Name: "BridgeArn", Flag: "bridge-arn", Type: "*string", Required: true},
	{Name: "FlowSource", Flag: "flow-source", Type: "*types.UpdateBridgeFlowSourceRequest", Required: false},
	{Name: "NetworkSource", Flag: "network-source", Type: "*types.UpdateBridgeNetworkSourceRequest", Required: false},
	{Name: "SourceName", Flag: "source-name", Type: "*string", Required: true},
}

var fields_update_bridge_state = []leanruntime.Field{
	{Name: "BridgeArn", Flag: "bridge-arn", Type: "*string", Required: true},
	{Name: "DesiredState", Flag: "desired-state", Type: "types.DesiredState", Required: true},
}

var fields_update_flow = []leanruntime.Field{
	{Name: "EncodingConfig", Flag: "encoding-config", Type: "*types.EncodingConfig", Required: false},
	{Name: "FlowArn", Flag: "flow-arn", Type: "*string", Required: true},
	{Name: "FlowSize", Flag: "flow-size", Type: "types.FlowSize", Required: false},
	{Name: "Maintenance", Flag: "maintenance", Type: "*types.UpdateMaintenance", Required: false},
	{Name: "NdiConfig", Flag: "ndi-config", Type: "*types.NdiConfig", Required: false},
	{Name: "SourceFailoverConfig", Flag: "source-failover-config", Type: "*types.UpdateFailoverConfig", Required: false},
	{Name: "SourceMonitoringConfig", Flag: "source-monitoring-config", Type: "*types.MonitoringConfig", Required: false},
}

var fields_update_flow_entitlement = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Encryption", Flag: "encryption", Type: "*types.UpdateEncryption", Required: false},
	{Name: "EntitlementArn", Flag: "entitlement-arn", Type: "*string", Required: true},
	{Name: "EntitlementStatus", Flag: "entitlement-status", Type: "types.EntitlementStatus", Required: false},
	{Name: "FlowArn", Flag: "flow-arn", Type: "*string", Required: true},
	{Name: "Subscribers", Flag: "subscribers", Type: "[]string", Required: false},
}

var fields_update_flow_media_stream = []leanruntime.Field{
	{Name: "Attributes", Flag: "attributes", Type: "*types.MediaStreamAttributesRequest", Required: false},
	{Name: "ClockRate", Flag: "clock-rate", Type: "*int32", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "FlowArn", Flag: "flow-arn", Type: "*string", Required: true},
	{Name: "MediaStreamName", Flag: "media-stream-name", Type: "*string", Required: true},
	{Name: "MediaStreamType", Flag: "media-stream-type", Type: "types.MediaStreamType", Required: false},
	{Name: "VideoFormat", Flag: "video-format", Type: "*string", Required: false},
}

var fields_update_flow_output = []leanruntime.Field{
	{Name: "CidrAllowList", Flag: "cidr-allow-list", Type: "[]string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Destination", Flag: "destination", Type: "*string", Required: false},
	{Name: "Encryption", Flag: "encryption", Type: "*types.UpdateEncryption", Required: false},
	{Name: "FlowArn", Flag: "flow-arn", Type: "*string", Required: true},
	{Name: "MaxLatency", Flag: "max-latency", Type: "*int32", Required: false},
	{Name: "MediaStreamOutputConfigurations", Flag: "media-stream-output-configurations", Type: "[]types.MediaStreamOutputConfigurationRequest", Required: false},
	{Name: "MinLatency", Flag: "min-latency", Type: "*int32", Required: false},
	{Name: "NdiProgramName", Flag: "ndi-program-name", Type: "*string", Required: false},
	{Name: "NdiSpeedHqQuality", Flag: "ndi-speed-hq-quality", Type: "*int32", Required: false},
	{Name: "OutputArn", Flag: "output-arn", Type: "*string", Required: true},
	{Name: "OutputStatus", Flag: "output-status", Type: "types.OutputStatus", Required: false},
	{Name: "Port", Flag: "port", Type: "*int32", Required: false},
	{Name: "Protocol", Flag: "protocol", Type: "types.Protocol", Required: false},
	{Name: "RemoteId", Flag: "remote-id", Type: "*string", Required: false},
	{Name: "RouterIntegrationState", Flag: "router-integration-state", Type: "types.State", Required: false},
	{Name: "RouterIntegrationTransitEncryption", Flag: "router-integration-transit-encryption", Type: "*types.FlowTransitEncryption", Required: false},
	{Name: "SenderControlPort", Flag: "sender-control-port", Type: "*int32", Required: false},
	{Name: "SenderIpAddress", Flag: "sender-ip-address", Type: "*string", Required: false},
	{Name: "SmoothingLatency", Flag: "smoothing-latency", Type: "*int32", Required: false},
	{Name: "StreamId", Flag: "stream-id", Type: "*string", Required: false},
	{Name: "VpcInterfaceAttachment", Flag: "vpc-interface-attachment", Type: "*types.VpcInterfaceAttachment", Required: false},
}

var fields_update_flow_source = []leanruntime.Field{
	{Name: "Decryption", Flag: "decryption", Type: "*types.UpdateEncryption", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EntitlementArn", Flag: "entitlement-arn", Type: "*string", Required: false},
	{Name: "FlowArn", Flag: "flow-arn", Type: "*string", Required: true},
	{Name: "GatewayBridgeSource", Flag: "gateway-bridge-source", Type: "*types.UpdateGatewayBridgeSourceRequest", Required: false},
	{Name: "IngestPort", Flag: "ingest-port", Type: "*int32", Required: false},
	{Name: "MaxBitrate", Flag: "max-bitrate", Type: "*int32", Required: false},
	{Name: "MaxLatency", Flag: "max-latency", Type: "*int32", Required: false},
	{Name: "MaxSyncBuffer", Flag: "max-sync-buffer", Type: "*int32", Required: false},
	{Name: "MediaStreamSourceConfigurations", Flag: "media-stream-source-configurations", Type: "[]types.MediaStreamSourceConfigurationRequest", Required: false},
	{Name: "MinLatency", Flag: "min-latency", Type: "*int32", Required: false},
	{Name: "NdiSourceSettings", Flag: "ndi-source-settings", Type: "*types.NdiSourceSettings", Required: false},
	{Name: "Protocol", Flag: "protocol", Type: "types.Protocol", Required: false},
	{Name: "RouterIntegrationState", Flag: "router-integration-state", Type: "types.State", Required: false},
	{Name: "RouterIntegrationTransitDecryption", Flag: "router-integration-transit-decryption", Type: "*types.FlowTransitEncryption", Required: false},
	{Name: "SenderControlPort", Flag: "sender-control-port", Type: "*int32", Required: false},
	{Name: "SenderIpAddress", Flag: "sender-ip-address", Type: "*string", Required: false},
	{Name: "SourceArn", Flag: "source-arn", Type: "*string", Required: true},
	{Name: "SourceListenerAddress", Flag: "source-listener-address", Type: "*string", Required: false},
	{Name: "SourceListenerPort", Flag: "source-listener-port", Type: "*int32", Required: false},
	{Name: "StreamId", Flag: "stream-id", Type: "*string", Required: false},
	{Name: "VpcInterfaceName", Flag: "vpc-interface-name", Type: "*string", Required: false},
	{Name: "WhitelistCidr", Flag: "whitelist-cidr", Type: "*string", Required: false},
}

var fields_update_gateway_instance = []leanruntime.Field{
	{Name: "BridgePlacement", Flag: "bridge-placement", Type: "types.BridgePlacement", Required: false},
	{Name: "GatewayInstanceArn", Flag: "gateway-instance-arn", Type: "*string", Required: true},
}

var fields_update_router_input = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "Configuration", Flag: "configuration", Type: "types.RouterInputConfiguration", Required: false},
	{Name: "MaintenanceConfiguration", Flag: "maintenance-configuration", Type: "types.MaintenanceConfiguration", Required: false},
	{Name: "MaximumBitrate", Flag: "maximum-bitrate", Type: "*int64", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "RoutingScope", Flag: "routing-scope", Type: "types.RoutingScope", Required: false},
	{Name: "Tier", Flag: "tier", Type: "types.RouterInputTier", Required: false},
	{Name: "TransitEncryption", Flag: "transit-encryption", Type: "*types.RouterInputTransitEncryption", Required: false},
}

var fields_update_router_network_interface = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "Configuration", Flag: "configuration", Type: "types.RouterNetworkInterfaceConfiguration", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_router_output = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "Configuration", Flag: "configuration", Type: "types.RouterOutputConfiguration", Required: false},
	{Name: "MaintenanceConfiguration", Flag: "maintenance-configuration", Type: "types.MaintenanceConfiguration", Required: false},
	{Name: "MaximumBitrate", Flag: "maximum-bitrate", Type: "*int64", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "RoutingScope", Flag: "routing-scope", Type: "types.RoutingScope", Required: false},
	{Name: "Tier", Flag: "tier", Type: "types.RouterOutputTier", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"add-bridge-outputs": {
			Name:   "add-bridge-outputs",
			Fields: fields_add_bridge_outputs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddBridgeOutputsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_bridge_outputs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddBridgeOutputs(ctx, input)
			},
		},
		"add-bridge-sources": {
			Name:   "add-bridge-sources",
			Fields: fields_add_bridge_sources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddBridgeSourcesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_bridge_sources, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddBridgeSources(ctx, input)
			},
		},
		"add-flow-media-streams": {
			Name:   "add-flow-media-streams",
			Fields: fields_add_flow_media_streams,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddFlowMediaStreamsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_flow_media_streams, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddFlowMediaStreams(ctx, input)
			},
		},
		"add-flow-outputs": {
			Name:   "add-flow-outputs",
			Fields: fields_add_flow_outputs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddFlowOutputsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_flow_outputs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddFlowOutputs(ctx, input)
			},
		},
		"add-flow-sources": {
			Name:   "add-flow-sources",
			Fields: fields_add_flow_sources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddFlowSourcesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_flow_sources, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddFlowSources(ctx, input)
			},
		},
		"add-flow-vpc-interfaces": {
			Name:   "add-flow-vpc-interfaces",
			Fields: fields_add_flow_vpc_interfaces,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddFlowVpcInterfacesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_flow_vpc_interfaces, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddFlowVpcInterfaces(ctx, input)
			},
		},
		"batch-get-router-input": {
			Name:   "batch-get-router-input",
			Fields: fields_batch_get_router_input,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetRouterInputInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_router_input, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetRouterInput(ctx, input)
			},
		},
		"batch-get-router-network-interface": {
			Name:   "batch-get-router-network-interface",
			Fields: fields_batch_get_router_network_interface,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetRouterNetworkInterfaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_router_network_interface, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetRouterNetworkInterface(ctx, input)
			},
		},
		"batch-get-router-output": {
			Name:   "batch-get-router-output",
			Fields: fields_batch_get_router_output,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetRouterOutputInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_router_output, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetRouterOutput(ctx, input)
			},
		},
		"create-bridge": {
			Name:   "create-bridge",
			Fields: fields_create_bridge,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBridgeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_bridge, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBridge(ctx, input)
			},
		},
		"create-flow": {
			Name:   "create-flow",
			Fields: fields_create_flow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFlowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_flow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFlow(ctx, input)
			},
		},
		"create-gateway": {
			Name:   "create-gateway",
			Fields: fields_create_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateGateway(ctx, input)
			},
		},
		"create-router-input": {
			Name:   "create-router-input",
			Fields: fields_create_router_input,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRouterInputInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_router_input, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRouterInput(ctx, input)
			},
		},
		"create-router-network-interface": {
			Name:   "create-router-network-interface",
			Fields: fields_create_router_network_interface,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRouterNetworkInterfaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_router_network_interface, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRouterNetworkInterface(ctx, input)
			},
		},
		"create-router-output": {
			Name:   "create-router-output",
			Fields: fields_create_router_output,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRouterOutputInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_router_output, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRouterOutput(ctx, input)
			},
		},
		"delete-bridge": {
			Name:   "delete-bridge",
			Fields: fields_delete_bridge,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBridgeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_bridge, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBridge(ctx, input)
			},
		},
		"delete-flow": {
			Name:   "delete-flow",
			Fields: fields_delete_flow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFlowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_flow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFlow(ctx, input)
			},
		},
		"delete-gateway": {
			Name:   "delete-gateway",
			Fields: fields_delete_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteGateway(ctx, input)
			},
		},
		"delete-router-input": {
			Name:   "delete-router-input",
			Fields: fields_delete_router_input,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRouterInputInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_router_input, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRouterInput(ctx, input)
			},
		},
		"delete-router-network-interface": {
			Name:   "delete-router-network-interface",
			Fields: fields_delete_router_network_interface,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRouterNetworkInterfaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_router_network_interface, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRouterNetworkInterface(ctx, input)
			},
		},
		"delete-router-output": {
			Name:   "delete-router-output",
			Fields: fields_delete_router_output,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRouterOutputInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_router_output, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRouterOutput(ctx, input)
			},
		},
		"deregister-gateway-instance": {
			Name:   "deregister-gateway-instance",
			Fields: fields_deregister_gateway_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterGatewayInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_gateway_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterGatewayInstance(ctx, input)
			},
		},
		"describe-bridge": {
			Name:   "describe-bridge",
			Fields: fields_describe_bridge,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeBridgeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_bridge, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeBridge(ctx, input)
			},
		},
		"describe-flow": {
			Name:   "describe-flow",
			Fields: fields_describe_flow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFlowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_flow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeFlow(ctx, input)
			},
		},
		"describe-flow-source-metadata": {
			Name:   "describe-flow-source-metadata",
			Fields: fields_describe_flow_source_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFlowSourceMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_flow_source_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeFlowSourceMetadata(ctx, input)
			},
		},
		"describe-flow-source-thumbnail": {
			Name:   "describe-flow-source-thumbnail",
			Fields: fields_describe_flow_source_thumbnail,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFlowSourceThumbnailInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_flow_source_thumbnail, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeFlowSourceThumbnail(ctx, input)
			},
		},
		"describe-gateway": {
			Name:   "describe-gateway",
			Fields: fields_describe_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeGateway(ctx, input)
			},
		},
		"describe-gateway-instance": {
			Name:   "describe-gateway-instance",
			Fields: fields_describe_gateway_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeGatewayInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_gateway_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeGatewayInstance(ctx, input)
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
		"get-router-input": {
			Name:   "get-router-input",
			Fields: fields_get_router_input,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRouterInputInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_router_input, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRouterInput(ctx, input)
			},
		},
		"get-router-input-source-metadata": {
			Name:   "get-router-input-source-metadata",
			Fields: fields_get_router_input_source_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRouterInputSourceMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_router_input_source_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRouterInputSourceMetadata(ctx, input)
			},
		},
		"get-router-input-thumbnail": {
			Name:   "get-router-input-thumbnail",
			Fields: fields_get_router_input_thumbnail,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRouterInputThumbnailInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_router_input_thumbnail, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRouterInputThumbnail(ctx, input)
			},
		},
		"get-router-network-interface": {
			Name:   "get-router-network-interface",
			Fields: fields_get_router_network_interface,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRouterNetworkInterfaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_router_network_interface, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRouterNetworkInterface(ctx, input)
			},
		},
		"get-router-output": {
			Name:   "get-router-output",
			Fields: fields_get_router_output,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRouterOutputInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_router_output, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRouterOutput(ctx, input)
			},
		},
		"grant-flow-entitlements": {
			Name:   "grant-flow-entitlements",
			Fields: fields_grant_flow_entitlements,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GrantFlowEntitlementsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_grant_flow_entitlements, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GrantFlowEntitlements(ctx, input)
			},
		},
		"list-bridges": {
			Name:   "list-bridges",
			Fields: fields_list_bridges,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBridgesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_bridges, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBridges(ctx, input)
				}
				var results []*svc.ListBridgesOutput
				p := svc.NewListBridgesPaginator(client, input)
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
		"list-entitlements": {
			Name:   "list-entitlements",
			Fields: fields_list_entitlements,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEntitlementsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_entitlements, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEntitlements(ctx, input)
				}
				var results []*svc.ListEntitlementsOutput
				p := svc.NewListEntitlementsPaginator(client, input)
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
		"list-flows": {
			Name:   "list-flows",
			Fields: fields_list_flows,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFlowsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_flows, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFlows(ctx, input)
				}
				var results []*svc.ListFlowsOutput
				p := svc.NewListFlowsPaginator(client, input)
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
		"list-gateway-instances": {
			Name:   "list-gateway-instances",
			Fields: fields_list_gateway_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGatewayInstancesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_gateway_instances, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListGatewayInstances(ctx, input)
				}
				var results []*svc.ListGatewayInstancesOutput
				p := svc.NewListGatewayInstancesPaginator(client, input)
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
		"list-gateways": {
			Name:   "list-gateways",
			Fields: fields_list_gateways,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGatewaysInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_gateways, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListGateways(ctx, input)
				}
				var results []*svc.ListGatewaysOutput
				p := svc.NewListGatewaysPaginator(client, input)
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
		"list-router-inputs": {
			Name:   "list-router-inputs",
			Fields: fields_list_router_inputs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRouterInputsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_router_inputs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRouterInputs(ctx, input)
				}
				var results []*svc.ListRouterInputsOutput
				p := svc.NewListRouterInputsPaginator(client, input)
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
		"list-router-network-interfaces": {
			Name:   "list-router-network-interfaces",
			Fields: fields_list_router_network_interfaces,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRouterNetworkInterfacesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_router_network_interfaces, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRouterNetworkInterfaces(ctx, input)
				}
				var results []*svc.ListRouterNetworkInterfacesOutput
				p := svc.NewListRouterNetworkInterfacesPaginator(client, input)
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
		"list-router-outputs": {
			Name:   "list-router-outputs",
			Fields: fields_list_router_outputs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRouterOutputsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_router_outputs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRouterOutputs(ctx, input)
				}
				var results []*svc.ListRouterOutputsOutput
				p := svc.NewListRouterOutputsPaginator(client, input)
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
		"list-tags-for-global-resource": {
			Name:   "list-tags-for-global-resource",
			Fields: fields_list_tags_for_global_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTagsForGlobalResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_tags_for_global_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTagsForGlobalResource(ctx, input)
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
		"remove-bridge-output": {
			Name:   "remove-bridge-output",
			Fields: fields_remove_bridge_output,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveBridgeOutputInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_bridge_output, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveBridgeOutput(ctx, input)
			},
		},
		"remove-bridge-source": {
			Name:   "remove-bridge-source",
			Fields: fields_remove_bridge_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveBridgeSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_bridge_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveBridgeSource(ctx, input)
			},
		},
		"remove-flow-media-stream": {
			Name:   "remove-flow-media-stream",
			Fields: fields_remove_flow_media_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveFlowMediaStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_flow_media_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveFlowMediaStream(ctx, input)
			},
		},
		"remove-flow-output": {
			Name:   "remove-flow-output",
			Fields: fields_remove_flow_output,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveFlowOutputInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_flow_output, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveFlowOutput(ctx, input)
			},
		},
		"remove-flow-source": {
			Name:   "remove-flow-source",
			Fields: fields_remove_flow_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveFlowSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_flow_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveFlowSource(ctx, input)
			},
		},
		"remove-flow-vpc-interface": {
			Name:   "remove-flow-vpc-interface",
			Fields: fields_remove_flow_vpc_interface,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveFlowVpcInterfaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_flow_vpc_interface, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveFlowVpcInterface(ctx, input)
			},
		},
		"restart-router-input": {
			Name:   "restart-router-input",
			Fields: fields_restart_router_input,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RestartRouterInputInput{}
				if _, err := leanruntime.ApplyInput(input, fields_restart_router_input, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RestartRouterInput(ctx, input)
			},
		},
		"restart-router-output": {
			Name:   "restart-router-output",
			Fields: fields_restart_router_output,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RestartRouterOutputInput{}
				if _, err := leanruntime.ApplyInput(input, fields_restart_router_output, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RestartRouterOutput(ctx, input)
			},
		},
		"revoke-flow-entitlement": {
			Name:   "revoke-flow-entitlement",
			Fields: fields_revoke_flow_entitlement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RevokeFlowEntitlementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_revoke_flow_entitlement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RevokeFlowEntitlement(ctx, input)
			},
		},
		"start-flow": {
			Name:   "start-flow",
			Fields: fields_start_flow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartFlowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_flow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartFlow(ctx, input)
			},
		},
		"start-router-input": {
			Name:   "start-router-input",
			Fields: fields_start_router_input,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartRouterInputInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_router_input, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartRouterInput(ctx, input)
			},
		},
		"start-router-output": {
			Name:   "start-router-output",
			Fields: fields_start_router_output,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartRouterOutputInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_router_output, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartRouterOutput(ctx, input)
			},
		},
		"stop-flow": {
			Name:   "stop-flow",
			Fields: fields_stop_flow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopFlowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_flow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopFlow(ctx, input)
			},
		},
		"stop-router-input": {
			Name:   "stop-router-input",
			Fields: fields_stop_router_input,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopRouterInputInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_router_input, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopRouterInput(ctx, input)
			},
		},
		"stop-router-output": {
			Name:   "stop-router-output",
			Fields: fields_stop_router_output,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopRouterOutputInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_router_output, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopRouterOutput(ctx, input)
			},
		},
		"tag-global-resource": {
			Name:   "tag-global-resource",
			Fields: fields_tag_global_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TagGlobalResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_tag_global_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TagGlobalResource(ctx, input)
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
		"take-router-input": {
			Name:   "take-router-input",
			Fields: fields_take_router_input,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TakeRouterInputInput{}
				if _, err := leanruntime.ApplyInput(input, fields_take_router_input, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TakeRouterInput(ctx, input)
			},
		},
		"untag-global-resource": {
			Name:   "untag-global-resource",
			Fields: fields_untag_global_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UntagGlobalResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_untag_global_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UntagGlobalResource(ctx, input)
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
		"update-bridge": {
			Name:   "update-bridge",
			Fields: fields_update_bridge,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateBridgeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_bridge, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateBridge(ctx, input)
			},
		},
		"update-bridge-output": {
			Name:   "update-bridge-output",
			Fields: fields_update_bridge_output,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateBridgeOutputInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_bridge_output, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateBridgeOutput(ctx, input)
			},
		},
		"update-bridge-source": {
			Name:   "update-bridge-source",
			Fields: fields_update_bridge_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateBridgeSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_bridge_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateBridgeSource(ctx, input)
			},
		},
		"update-bridge-state": {
			Name:   "update-bridge-state",
			Fields: fields_update_bridge_state,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateBridgeStateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_bridge_state, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateBridgeState(ctx, input)
			},
		},
		"update-flow": {
			Name:   "update-flow",
			Fields: fields_update_flow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFlowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_flow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFlow(ctx, input)
			},
		},
		"update-flow-entitlement": {
			Name:   "update-flow-entitlement",
			Fields: fields_update_flow_entitlement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFlowEntitlementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_flow_entitlement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFlowEntitlement(ctx, input)
			},
		},
		"update-flow-media-stream": {
			Name:   "update-flow-media-stream",
			Fields: fields_update_flow_media_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFlowMediaStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_flow_media_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFlowMediaStream(ctx, input)
			},
		},
		"update-flow-output": {
			Name:   "update-flow-output",
			Fields: fields_update_flow_output,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFlowOutputInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_flow_output, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFlowOutput(ctx, input)
			},
		},
		"update-flow-source": {
			Name:   "update-flow-source",
			Fields: fields_update_flow_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFlowSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_flow_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFlowSource(ctx, input)
			},
		},
		"update-gateway-instance": {
			Name:   "update-gateway-instance",
			Fields: fields_update_gateway_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateGatewayInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_gateway_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateGatewayInstance(ctx, input)
			},
		},
		"update-router-input": {
			Name:   "update-router-input",
			Fields: fields_update_router_input,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRouterInputInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_router_input, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRouterInput(ctx, input)
			},
		},
		"update-router-network-interface": {
			Name:   "update-router-network-interface",
			Fields: fields_update_router_network_interface,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRouterNetworkInterfaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_router_network_interface, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRouterNetworkInterface(ctx, input)
			},
		},
		"update-router-output": {
			Name:   "update-router-output",
			Fields: fields_update_router_output,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRouterOutputInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_router_output, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRouterOutput(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("mediaconnect", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
