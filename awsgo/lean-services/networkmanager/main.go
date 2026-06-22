package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/networkmanager"
)

var fields_accept_attachment = []leanruntime.Field{
	{Name: "AttachmentId", Flag: "attachment-id", Type: "*string", Required: true},
}

var fields_associate_connect_peer = []leanruntime.Field{
	{Name: "ConnectPeerId", Flag: "connect-peer-id", Type: "*string", Required: true},
	{Name: "DeviceId", Flag: "device-id", Type: "*string", Required: true},
	{Name: "GlobalNetworkId", Flag: "global-network-id", Type: "*string", Required: true},
	{Name: "LinkId", Flag: "link-id", Type: "*string", Required: false},
}

var fields_associate_customer_gateway = []leanruntime.Field{
	{Name: "CustomerGatewayArn", Flag: "customer-gateway-arn", Type: "*string", Required: true},
	{Name: "DeviceId", Flag: "device-id", Type: "*string", Required: true},
	{Name: "GlobalNetworkId", Flag: "global-network-id", Type: "*string", Required: true},
	{Name: "LinkId", Flag: "link-id", Type: "*string", Required: false},
}

var fields_associate_link = []leanruntime.Field{
	{Name: "DeviceId", Flag: "device-id", Type: "*string", Required: true},
	{Name: "GlobalNetworkId", Flag: "global-network-id", Type: "*string", Required: true},
	{Name: "LinkId", Flag: "link-id", Type: "*string", Required: true},
}

var fields_associate_transit_gateway_connect_peer = []leanruntime.Field{
	{Name: "DeviceId", Flag: "device-id", Type: "*string", Required: true},
	{Name: "GlobalNetworkId", Flag: "global-network-id", Type: "*string", Required: true},
	{Name: "LinkId", Flag: "link-id", Type: "*string", Required: false},
	{Name: "TransitGatewayConnectPeerArn", Flag: "transit-gateway-connect-peer-arn", Type: "*string", Required: true},
}

var fields_create_connect_attachment = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CoreNetworkId", Flag: "core-network-id", Type: "*string", Required: true},
	{Name: "EdgeLocation", Flag: "edge-location", Type: "*string", Required: true},
	{Name: "Options", Flag: "options", Type: "*types.ConnectAttachmentOptions", Required: true},
	{Name: "RoutingPolicyLabel", Flag: "routing-policy-label", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TransportAttachmentId", Flag: "transport-attachment-id", Type: "*string", Required: true},
}

var fields_create_connect_peer = []leanruntime.Field{
	{Name: "BgpOptions", Flag: "bgp-options", Type: "*types.BgpOptions", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ConnectAttachmentId", Flag: "connect-attachment-id", Type: "*string", Required: true},
	{Name: "CoreNetworkAddress", Flag: "core-network-address", Type: "*string", Required: false},
	{Name: "InsideCidrBlocks", Flag: "inside-cidr-blocks", Type: "[]string", Required: false},
	{Name: "PeerAddress", Flag: "peer-address", Type: "*string", Required: true},
	{Name: "SubnetArn", Flag: "subnet-arn", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_connection = []leanruntime.Field{
	{Name: "ConnectedDeviceId", Flag: "connected-device-id", Type: "*string", Required: true},
	{Name: "ConnectedLinkId", Flag: "connected-link-id", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DeviceId", Flag: "device-id", Type: "*string", Required: true},
	{Name: "GlobalNetworkId", Flag: "global-network-id", Type: "*string", Required: true},
	{Name: "LinkId", Flag: "link-id", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_core_network = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "GlobalNetworkId", Flag: "global-network-id", Type: "*string", Required: true},
	{Name: "PolicyDocument", Flag: "policy-document", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_core_network_prefix_list_association = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CoreNetworkId", Flag: "core-network-id", Type: "*string", Required: true},
	{Name: "PrefixListAlias", Flag: "prefix-list-alias", Type: "*string", Required: true},
	{Name: "PrefixListArn", Flag: "prefix-list-arn", Type: "*string", Required: true},
}

var fields_create_device = []leanruntime.Field{
	{Name: "AWSLocation", Flag: "aws-location", Type: "*types.AWSLocation", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "GlobalNetworkId", Flag: "global-network-id", Type: "*string", Required: true},
	{Name: "Location", Flag: "location", Type: "*types.Location", Required: false},
	{Name: "Model", Flag: "model", Type: "*string", Required: false},
	{Name: "SerialNumber", Flag: "serial-number", Type: "*string", Required: false},
	{Name: "SiteId", Flag: "site-id", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Type", Flag: "type", Type: "*string", Required: false},
	{Name: "Vendor", Flag: "vendor", Type: "*string", Required: false},
}

var fields_create_direct_connect_gateway_attachment = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CoreNetworkId", Flag: "core-network-id", Type: "*string", Required: true},
	{Name: "DirectConnectGatewayArn", Flag: "direct-connect-gateway-arn", Type: "*string", Required: true},
	{Name: "EdgeLocations", Flag: "edge-locations", Type: "[]string", Required: true},
	{Name: "RoutingPolicyLabel", Flag: "routing-policy-label", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_global_network = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_link = []leanruntime.Field{
	{Name: "Bandwidth", Flag: "bandwidth", Type: "*types.Bandwidth", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "GlobalNetworkId", Flag: "global-network-id", Type: "*string", Required: true},
	{Name: "Provider", Flag: "provider", Type: "*string", Required: false},
	{Name: "SiteId", Flag: "site-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Type", Flag: "type", Type: "*string", Required: false},
}

var fields_create_site = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "GlobalNetworkId", Flag: "global-network-id", Type: "*string", Required: true},
	{Name: "Location", Flag: "location", Type: "*types.Location", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_site_to_site_vpn_attachment = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CoreNetworkId", Flag: "core-network-id", Type: "*string", Required: true},
	{Name: "RoutingPolicyLabel", Flag: "routing-policy-label", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VpnConnectionArn", Flag: "vpn-connection-arn", Type: "*string", Required: true},
}

var fields_create_transit_gateway_peering = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CoreNetworkId", Flag: "core-network-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TransitGatewayArn", Flag: "transit-gateway-arn", Type: "*string", Required: true},
}

var fields_create_transit_gateway_route_table_attachment = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "PeeringId", Flag: "peering-id", Type: "*string", Required: true},
	{Name: "RoutingPolicyLabel", Flag: "routing-policy-label", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TransitGatewayRouteTableArn", Flag: "transit-gateway-route-table-arn", Type: "*string", Required: true},
}

var fields_create_vpc_attachment = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CoreNetworkId", Flag: "core-network-id", Type: "*string", Required: true},
	{Name: "Options", Flag: "options", Type: "*types.VpcOptions", Required: false},
	{Name: "RoutingPolicyLabel", Flag: "routing-policy-label", Type: "*string", Required: false},
	{Name: "SubnetArns", Flag: "subnet-arns", Type: "[]string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VpcArn", Flag: "vpc-arn", Type: "*string", Required: true},
}

var fields_delete_attachment = []leanruntime.Field{
	{Name: "AttachmentId", Flag: "attachment-id", Type: "*string", Required: true},
}

var fields_delete_connect_peer = []leanruntime.Field{
	{Name: "ConnectPeerId", Flag: "connect-peer-id", Type: "*string", Required: true},
}

var fields_delete_connection = []leanruntime.Field{
	{Name: "ConnectionId", Flag: "connection-id", Type: "*string", Required: true},
	{Name: "GlobalNetworkId", Flag: "global-network-id", Type: "*string", Required: true},
}

var fields_delete_core_network = []leanruntime.Field{
	{Name: "CoreNetworkId", Flag: "core-network-id", Type: "*string", Required: true},
}

var fields_delete_core_network_policy_version = []leanruntime.Field{
	{Name: "CoreNetworkId", Flag: "core-network-id", Type: "*string", Required: true},
	{Name: "PolicyVersionId", Flag: "policy-version-id", Type: "*int32", Required: true},
}

var fields_delete_core_network_prefix_list_association = []leanruntime.Field{
	{Name: "CoreNetworkId", Flag: "core-network-id", Type: "*string", Required: true},
	{Name: "PrefixListArn", Flag: "prefix-list-arn", Type: "*string", Required: true},
}

var fields_delete_device = []leanruntime.Field{
	{Name: "DeviceId", Flag: "device-id", Type: "*string", Required: true},
	{Name: "GlobalNetworkId", Flag: "global-network-id", Type: "*string", Required: true},
}

var fields_delete_global_network = []leanruntime.Field{
	{Name: "GlobalNetworkId", Flag: "global-network-id", Type: "*string", Required: true},
}

var fields_delete_link = []leanruntime.Field{
	{Name: "GlobalNetworkId", Flag: "global-network-id", Type: "*string", Required: true},
	{Name: "LinkId", Flag: "link-id", Type: "*string", Required: true},
}

var fields_delete_peering = []leanruntime.Field{
	{Name: "PeeringId", Flag: "peering-id", Type: "*string", Required: true},
}

var fields_delete_resource_policy = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_delete_site = []leanruntime.Field{
	{Name: "GlobalNetworkId", Flag: "global-network-id", Type: "*string", Required: true},
	{Name: "SiteId", Flag: "site-id", Type: "*string", Required: true},
}

var fields_deregister_transit_gateway = []leanruntime.Field{
	{Name: "GlobalNetworkId", Flag: "global-network-id", Type: "*string", Required: true},
	{Name: "TransitGatewayArn", Flag: "transit-gateway-arn", Type: "*string", Required: true},
}

var fields_describe_global_networks = []leanruntime.Field{
	{Name: "GlobalNetworkIds", Flag: "global-network-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_disassociate_connect_peer = []leanruntime.Field{
	{Name: "ConnectPeerId", Flag: "connect-peer-id", Type: "*string", Required: true},
	{Name: "GlobalNetworkId", Flag: "global-network-id", Type: "*string", Required: true},
}

var fields_disassociate_customer_gateway = []leanruntime.Field{
	{Name: "CustomerGatewayArn", Flag: "customer-gateway-arn", Type: "*string", Required: true},
	{Name: "GlobalNetworkId", Flag: "global-network-id", Type: "*string", Required: true},
}

var fields_disassociate_link = []leanruntime.Field{
	{Name: "DeviceId", Flag: "device-id", Type: "*string", Required: true},
	{Name: "GlobalNetworkId", Flag: "global-network-id", Type: "*string", Required: true},
	{Name: "LinkId", Flag: "link-id", Type: "*string", Required: true},
}

var fields_disassociate_transit_gateway_connect_peer = []leanruntime.Field{
	{Name: "GlobalNetworkId", Flag: "global-network-id", Type: "*string", Required: true},
	{Name: "TransitGatewayConnectPeerArn", Flag: "transit-gateway-connect-peer-arn", Type: "*string", Required: true},
}

var fields_execute_core_network_change_set = []leanruntime.Field{
	{Name: "CoreNetworkId", Flag: "core-network-id", Type: "*string", Required: true},
	{Name: "PolicyVersionId", Flag: "policy-version-id", Type: "*int32", Required: true},
}

var fields_get_connect_attachment = []leanruntime.Field{
	{Name: "AttachmentId", Flag: "attachment-id", Type: "*string", Required: true},
}

var fields_get_connect_peer = []leanruntime.Field{
	{Name: "ConnectPeerId", Flag: "connect-peer-id", Type: "*string", Required: true},
}

var fields_get_connect_peer_associations = []leanruntime.Field{
	{Name: "ConnectPeerIds", Flag: "connect-peer-ids", Type: "[]string", Required: false},
	{Name: "GlobalNetworkId", Flag: "global-network-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_connections = []leanruntime.Field{
	{Name: "ConnectionIds", Flag: "connection-ids", Type: "[]string", Required: false},
	{Name: "DeviceId", Flag: "device-id", Type: "*string", Required: false},
	{Name: "GlobalNetworkId", Flag: "global-network-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_core_network = []leanruntime.Field{
	{Name: "CoreNetworkId", Flag: "core-network-id", Type: "*string", Required: true},
}

var fields_get_core_network_change_events = []leanruntime.Field{
	{Name: "CoreNetworkId", Flag: "core-network-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PolicyVersionId", Flag: "policy-version-id", Type: "*int32", Required: true},
}

var fields_get_core_network_change_set = []leanruntime.Field{
	{Name: "CoreNetworkId", Flag: "core-network-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PolicyVersionId", Flag: "policy-version-id", Type: "*int32", Required: true},
}

var fields_get_core_network_policy = []leanruntime.Field{
	{Name: "Alias", Flag: "alias", Type: "types.CoreNetworkPolicyAlias", Required: false},
	{Name: "CoreNetworkId", Flag: "core-network-id", Type: "*string", Required: true},
	{Name: "PolicyVersionId", Flag: "policy-version-id", Type: "*int32", Required: false},
}

var fields_get_customer_gateway_associations = []leanruntime.Field{
	{Name: "CustomerGatewayArns", Flag: "customer-gateway-arns", Type: "[]string", Required: false},
	{Name: "GlobalNetworkId", Flag: "global-network-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_devices = []leanruntime.Field{
	{Name: "DeviceIds", Flag: "device-ids", Type: "[]string", Required: false},
	{Name: "GlobalNetworkId", Flag: "global-network-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SiteId", Flag: "site-id", Type: "*string", Required: false},
}

var fields_get_direct_connect_gateway_attachment = []leanruntime.Field{
	{Name: "AttachmentId", Flag: "attachment-id", Type: "*string", Required: true},
}

var fields_get_link_associations = []leanruntime.Field{
	{Name: "DeviceId", Flag: "device-id", Type: "*string", Required: false},
	{Name: "GlobalNetworkId", Flag: "global-network-id", Type: "*string", Required: true},
	{Name: "LinkId", Flag: "link-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_links = []leanruntime.Field{
	{Name: "GlobalNetworkId", Flag: "global-network-id", Type: "*string", Required: true},
	{Name: "LinkIds", Flag: "link-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Provider", Flag: "provider", Type: "*string", Required: false},
	{Name: "SiteId", Flag: "site-id", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "*string", Required: false},
}

var fields_get_network_resource_counts = []leanruntime.Field{
	{Name: "GlobalNetworkId", Flag: "global-network-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceType", Flag: "resource-type", Type: "*string", Required: false},
}

var fields_get_network_resource_relationships = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
	{Name: "AwsRegion", Flag: "aws-region", Type: "*string", Required: false},
	{Name: "CoreNetworkId", Flag: "core-network-id", Type: "*string", Required: false},
	{Name: "GlobalNetworkId", Flag: "global-network-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RegisteredGatewayArn", Flag: "registered-gateway-arn", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: false},
	{Name: "ResourceType", Flag: "resource-type", Type: "*string", Required: false},
}

var fields_get_network_resources = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
	{Name: "AwsRegion", Flag: "aws-region", Type: "*string", Required: false},
	{Name: "CoreNetworkId", Flag: "core-network-id", Type: "*string", Required: false},
	{Name: "GlobalNetworkId", Flag: "global-network-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RegisteredGatewayArn", Flag: "registered-gateway-arn", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: false},
	{Name: "ResourceType", Flag: "resource-type", Type: "*string", Required: false},
}

var fields_get_network_routes = []leanruntime.Field{
	{Name: "DestinationFilters", Flag: "destination-filters", Type: "map[string][]string", Required: false},
	{Name: "ExactCidrMatches", Flag: "exact-cidr-matches", Type: "[]string", Required: false},
	{Name: "GlobalNetworkId", Flag: "global-network-id", Type: "*string", Required: true},
	{Name: "LongestPrefixMatches", Flag: "longest-prefix-matches", Type: "[]string", Required: false},
	{Name: "PrefixListIds", Flag: "prefix-list-ids", Type: "[]string", Required: false},
	{Name: "RouteTableIdentifier", Flag: "route-table-identifier", Type: "*types.RouteTableIdentifier", Required: true},
	{Name: "States", Flag: "states", Type: "[]types.RouteState", Required: false},
	{Name: "SubnetOfMatches", Flag: "subnet-of-matches", Type: "[]string", Required: false},
	{Name: "SupernetOfMatches", Flag: "supernet-of-matches", Type: "[]string", Required: false},
	{Name: "Types", Flag: "types", Type: "[]types.RouteType", Required: false},
}

var fields_get_network_telemetry = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
	{Name: "AwsRegion", Flag: "aws-region", Type: "*string", Required: false},
	{Name: "CoreNetworkId", Flag: "core-network-id", Type: "*string", Required: false},
	{Name: "GlobalNetworkId", Flag: "global-network-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RegisteredGatewayArn", Flag: "registered-gateway-arn", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: false},
	{Name: "ResourceType", Flag: "resource-type", Type: "*string", Required: false},
}

var fields_get_resource_policy = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_get_route_analysis = []leanruntime.Field{
	{Name: "GlobalNetworkId", Flag: "global-network-id", Type: "*string", Required: true},
	{Name: "RouteAnalysisId", Flag: "route-analysis-id", Type: "*string", Required: true},
}

var fields_get_site_to_site_vpn_attachment = []leanruntime.Field{
	{Name: "AttachmentId", Flag: "attachment-id", Type: "*string", Required: true},
}

var fields_get_sites = []leanruntime.Field{
	{Name: "GlobalNetworkId", Flag: "global-network-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SiteIds", Flag: "site-ids", Type: "[]string", Required: false},
}

var fields_get_transit_gateway_connect_peer_associations = []leanruntime.Field{
	{Name: "GlobalNetworkId", Flag: "global-network-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TransitGatewayConnectPeerArns", Flag: "transit-gateway-connect-peer-arns", Type: "[]string", Required: false},
}

var fields_get_transit_gateway_peering = []leanruntime.Field{
	{Name: "PeeringId", Flag: "peering-id", Type: "*string", Required: true},
}

var fields_get_transit_gateway_registrations = []leanruntime.Field{
	{Name: "GlobalNetworkId", Flag: "global-network-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TransitGatewayArns", Flag: "transit-gateway-arns", Type: "[]string", Required: false},
}

var fields_get_transit_gateway_route_table_attachment = []leanruntime.Field{
	{Name: "AttachmentId", Flag: "attachment-id", Type: "*string", Required: true},
}

var fields_get_vpc_attachment = []leanruntime.Field{
	{Name: "AttachmentId", Flag: "attachment-id", Type: "*string", Required: true},
}

var fields_list_attachment_routing_policy_associations = []leanruntime.Field{
	{Name: "AttachmentId", Flag: "attachment-id", Type: "*string", Required: false},
	{Name: "CoreNetworkId", Flag: "core-network-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_attachments = []leanruntime.Field{
	{Name: "AttachmentType", Flag: "attachment-type", Type: "types.AttachmentType", Required: false},
	{Name: "CoreNetworkId", Flag: "core-network-id", Type: "*string", Required: false},
	{Name: "EdgeLocation", Flag: "edge-location", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "State", Flag: "state", Type: "types.AttachmentState", Required: false},
}

var fields_list_connect_peers = []leanruntime.Field{
	{Name: "ConnectAttachmentId", Flag: "connect-attachment-id", Type: "*string", Required: false},
	{Name: "CoreNetworkId", Flag: "core-network-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_core_network_policy_versions = []leanruntime.Field{
	{Name: "CoreNetworkId", Flag: "core-network-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_core_network_prefix_list_associations = []leanruntime.Field{
	{Name: "CoreNetworkId", Flag: "core-network-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PrefixListArn", Flag: "prefix-list-arn", Type: "*string", Required: false},
}

var fields_list_core_network_routing_information = []leanruntime.Field{
	{Name: "CommunityMatches", Flag: "community-matches", Type: "[]string", Required: false},
	{Name: "CoreNetworkId", Flag: "core-network-id", Type: "*string", Required: true},
	{Name: "EdgeLocation", Flag: "edge-location", Type: "*string", Required: true},
	{Name: "ExactAsPathMatches", Flag: "exact-as-path-matches", Type: "[]string", Required: false},
	{Name: "LocalPreferenceMatches", Flag: "local-preference-matches", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MedMatches", Flag: "med-matches", Type: "[]string", Required: false},
	{Name: "NextHopFilters", Flag: "next-hop-filters", Type: "map[string][]string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SegmentName", Flag: "segment-name", Type: "*string", Required: true},
}

var fields_list_core_networks = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_organization_service_access_status = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_peerings = []leanruntime.Field{
	{Name: "CoreNetworkId", Flag: "core-network-id", Type: "*string", Required: false},
	{Name: "EdgeLocation", Flag: "edge-location", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PeeringType", Flag: "peering-type", Type: "types.PeeringType", Required: false},
	{Name: "State", Flag: "state", Type: "types.PeeringState", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_put_attachment_routing_policy_label = []leanruntime.Field{
	{Name: "AttachmentId", Flag: "attachment-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CoreNetworkId", Flag: "core-network-id", Type: "*string", Required: true},
	{Name: "RoutingPolicyLabel", Flag: "routing-policy-label", Type: "*string", Required: true},
}

var fields_put_core_network_policy = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CoreNetworkId", Flag: "core-network-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "LatestVersionId", Flag: "latest-version-id", Type: "*int32", Required: false},
	{Name: "PolicyDocument", Flag: "policy-document", Type: "*string", Required: true},
}

var fields_put_resource_policy = []leanruntime.Field{
	{Name: "PolicyDocument", Flag: "policy-document", Type: "*string", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_register_transit_gateway = []leanruntime.Field{
	{Name: "GlobalNetworkId", Flag: "global-network-id", Type: "*string", Required: true},
	{Name: "TransitGatewayArn", Flag: "transit-gateway-arn", Type: "*string", Required: true},
}

var fields_reject_attachment = []leanruntime.Field{
	{Name: "AttachmentId", Flag: "attachment-id", Type: "*string", Required: true},
}

var fields_remove_attachment_routing_policy_label = []leanruntime.Field{
	{Name: "AttachmentId", Flag: "attachment-id", Type: "*string", Required: true},
	{Name: "CoreNetworkId", Flag: "core-network-id", Type: "*string", Required: true},
}

var fields_restore_core_network_policy_version = []leanruntime.Field{
	{Name: "CoreNetworkId", Flag: "core-network-id", Type: "*string", Required: true},
	{Name: "PolicyVersionId", Flag: "policy-version-id", Type: "*int32", Required: true},
}

var fields_start_organization_service_access_update = []leanruntime.Field{
	{Name: "Action", Flag: "action", Type: "*string", Required: true},
}

var fields_start_route_analysis = []leanruntime.Field{
	{Name: "Destination", Flag: "destination", Type: "*types.RouteAnalysisEndpointOptionsSpecification", Required: true},
	{Name: "GlobalNetworkId", Flag: "global-network-id", Type: "*string", Required: true},
	{Name: "IncludeReturnPath", Flag: "include-return-path", Type: "bool", Required: false},
	{Name: "Source", Flag: "source", Type: "*types.RouteAnalysisEndpointOptionsSpecification", Required: true},
	{Name: "UseMiddleboxes", Flag: "use-middleboxes", Type: "bool", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_connection = []leanruntime.Field{
	{Name: "ConnectedLinkId", Flag: "connected-link-id", Type: "*string", Required: false},
	{Name: "ConnectionId", Flag: "connection-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "GlobalNetworkId", Flag: "global-network-id", Type: "*string", Required: true},
	{Name: "LinkId", Flag: "link-id", Type: "*string", Required: false},
}

var fields_update_core_network = []leanruntime.Field{
	{Name: "CoreNetworkId", Flag: "core-network-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
}

var fields_update_device = []leanruntime.Field{
	{Name: "AWSLocation", Flag: "aws-location", Type: "*types.AWSLocation", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DeviceId", Flag: "device-id", Type: "*string", Required: true},
	{Name: "GlobalNetworkId", Flag: "global-network-id", Type: "*string", Required: true},
	{Name: "Location", Flag: "location", Type: "*types.Location", Required: false},
	{Name: "Model", Flag: "model", Type: "*string", Required: false},
	{Name: "SerialNumber", Flag: "serial-number", Type: "*string", Required: false},
	{Name: "SiteId", Flag: "site-id", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "*string", Required: false},
	{Name: "Vendor", Flag: "vendor", Type: "*string", Required: false},
}

var fields_update_direct_connect_gateway_attachment = []leanruntime.Field{
	{Name: "AttachmentId", Flag: "attachment-id", Type: "*string", Required: true},
	{Name: "EdgeLocations", Flag: "edge-locations", Type: "[]string", Required: false},
}

var fields_update_global_network = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "GlobalNetworkId", Flag: "global-network-id", Type: "*string", Required: true},
}

var fields_update_link = []leanruntime.Field{
	{Name: "Bandwidth", Flag: "bandwidth", Type: "*types.Bandwidth", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "GlobalNetworkId", Flag: "global-network-id", Type: "*string", Required: true},
	{Name: "LinkId", Flag: "link-id", Type: "*string", Required: true},
	{Name: "Provider", Flag: "provider", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "*string", Required: false},
}

var fields_update_network_resource_metadata = []leanruntime.Field{
	{Name: "GlobalNetworkId", Flag: "global-network-id", Type: "*string", Required: true},
	{Name: "Metadata", Flag: "metadata", Type: "map[string]string", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_update_site = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "GlobalNetworkId", Flag: "global-network-id", Type: "*string", Required: true},
	{Name: "Location", Flag: "location", Type: "*types.Location", Required: false},
	{Name: "SiteId", Flag: "site-id", Type: "*string", Required: true},
}

var fields_update_vpc_attachment = []leanruntime.Field{
	{Name: "AddSubnetArns", Flag: "add-subnet-arns", Type: "[]string", Required: false},
	{Name: "AttachmentId", Flag: "attachment-id", Type: "*string", Required: true},
	{Name: "Options", Flag: "options", Type: "*types.VpcOptions", Required: false},
	{Name: "RemoveSubnetArns", Flag: "remove-subnet-arns", Type: "[]string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"accept-attachment": {
			Name:   "accept-attachment",
			Fields: fields_accept_attachment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AcceptAttachmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_accept_attachment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AcceptAttachment(ctx, input)
			},
		},
		"associate-connect-peer": {
			Name:   "associate-connect-peer",
			Fields: fields_associate_connect_peer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateConnectPeerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_connect_peer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateConnectPeer(ctx, input)
			},
		},
		"associate-customer-gateway": {
			Name:   "associate-customer-gateway",
			Fields: fields_associate_customer_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateCustomerGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_customer_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateCustomerGateway(ctx, input)
			},
		},
		"associate-link": {
			Name:   "associate-link",
			Fields: fields_associate_link,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateLinkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_link, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateLink(ctx, input)
			},
		},
		"associate-transit-gateway-connect-peer": {
			Name:   "associate-transit-gateway-connect-peer",
			Fields: fields_associate_transit_gateway_connect_peer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateTransitGatewayConnectPeerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_transit_gateway_connect_peer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateTransitGatewayConnectPeer(ctx, input)
			},
		},
		"create-connect-attachment": {
			Name:   "create-connect-attachment",
			Fields: fields_create_connect_attachment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateConnectAttachmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_connect_attachment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateConnectAttachment(ctx, input)
			},
		},
		"create-connect-peer": {
			Name:   "create-connect-peer",
			Fields: fields_create_connect_peer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateConnectPeerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_connect_peer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateConnectPeer(ctx, input)
			},
		},
		"create-connection": {
			Name:   "create-connection",
			Fields: fields_create_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateConnection(ctx, input)
			},
		},
		"create-core-network": {
			Name:   "create-core-network",
			Fields: fields_create_core_network,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCoreNetworkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_core_network, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCoreNetwork(ctx, input)
			},
		},
		"create-core-network-prefix-list-association": {
			Name:   "create-core-network-prefix-list-association",
			Fields: fields_create_core_network_prefix_list_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCoreNetworkPrefixListAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_core_network_prefix_list_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCoreNetworkPrefixListAssociation(ctx, input)
			},
		},
		"create-device": {
			Name:   "create-device",
			Fields: fields_create_device,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDeviceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_device, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDevice(ctx, input)
			},
		},
		"create-direct-connect-gateway-attachment": {
			Name:   "create-direct-connect-gateway-attachment",
			Fields: fields_create_direct_connect_gateway_attachment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDirectConnectGatewayAttachmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_direct_connect_gateway_attachment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDirectConnectGatewayAttachment(ctx, input)
			},
		},
		"create-global-network": {
			Name:   "create-global-network",
			Fields: fields_create_global_network,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateGlobalNetworkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_global_network, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateGlobalNetwork(ctx, input)
			},
		},
		"create-link": {
			Name:   "create-link",
			Fields: fields_create_link,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLinkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_link, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLink(ctx, input)
			},
		},
		"create-site": {
			Name:   "create-site",
			Fields: fields_create_site,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSiteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_site, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSite(ctx, input)
			},
		},
		"create-site-to-site-vpn-attachment": {
			Name:   "create-site-to-site-vpn-attachment",
			Fields: fields_create_site_to_site_vpn_attachment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSiteToSiteVpnAttachmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_site_to_site_vpn_attachment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSiteToSiteVpnAttachment(ctx, input)
			},
		},
		"create-transit-gateway-peering": {
			Name:   "create-transit-gateway-peering",
			Fields: fields_create_transit_gateway_peering,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTransitGatewayPeeringInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_transit_gateway_peering, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTransitGatewayPeering(ctx, input)
			},
		},
		"create-transit-gateway-route-table-attachment": {
			Name:   "create-transit-gateway-route-table-attachment",
			Fields: fields_create_transit_gateway_route_table_attachment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTransitGatewayRouteTableAttachmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_transit_gateway_route_table_attachment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTransitGatewayRouteTableAttachment(ctx, input)
			},
		},
		"create-vpc-attachment": {
			Name:   "create-vpc-attachment",
			Fields: fields_create_vpc_attachment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVpcAttachmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_vpc_attachment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVpcAttachment(ctx, input)
			},
		},
		"delete-attachment": {
			Name:   "delete-attachment",
			Fields: fields_delete_attachment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAttachmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_attachment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAttachment(ctx, input)
			},
		},
		"delete-connect-peer": {
			Name:   "delete-connect-peer",
			Fields: fields_delete_connect_peer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConnectPeerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_connect_peer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConnectPeer(ctx, input)
			},
		},
		"delete-connection": {
			Name:   "delete-connection",
			Fields: fields_delete_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConnection(ctx, input)
			},
		},
		"delete-core-network": {
			Name:   "delete-core-network",
			Fields: fields_delete_core_network,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCoreNetworkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_core_network, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCoreNetwork(ctx, input)
			},
		},
		"delete-core-network-policy-version": {
			Name:   "delete-core-network-policy-version",
			Fields: fields_delete_core_network_policy_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCoreNetworkPolicyVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_core_network_policy_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCoreNetworkPolicyVersion(ctx, input)
			},
		},
		"delete-core-network-prefix-list-association": {
			Name:   "delete-core-network-prefix-list-association",
			Fields: fields_delete_core_network_prefix_list_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCoreNetworkPrefixListAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_core_network_prefix_list_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCoreNetworkPrefixListAssociation(ctx, input)
			},
		},
		"delete-device": {
			Name:   "delete-device",
			Fields: fields_delete_device,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDeviceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_device, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDevice(ctx, input)
			},
		},
		"delete-global-network": {
			Name:   "delete-global-network",
			Fields: fields_delete_global_network,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteGlobalNetworkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_global_network, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteGlobalNetwork(ctx, input)
			},
		},
		"delete-link": {
			Name:   "delete-link",
			Fields: fields_delete_link,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLinkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_link, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLink(ctx, input)
			},
		},
		"delete-peering": {
			Name:   "delete-peering",
			Fields: fields_delete_peering,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePeeringInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_peering, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePeering(ctx, input)
			},
		},
		"delete-resource-policy": {
			Name:   "delete-resource-policy",
			Fields: fields_delete_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteResourcePolicy(ctx, input)
			},
		},
		"delete-site": {
			Name:   "delete-site",
			Fields: fields_delete_site,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSiteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_site, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSite(ctx, input)
			},
		},
		"deregister-transit-gateway": {
			Name:   "deregister-transit-gateway",
			Fields: fields_deregister_transit_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterTransitGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_transit_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterTransitGateway(ctx, input)
			},
		},
		"describe-global-networks": {
			Name:   "describe-global-networks",
			Fields: fields_describe_global_networks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeGlobalNetworksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_global_networks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeGlobalNetworks(ctx, input)
				}
				var results []*svc.DescribeGlobalNetworksOutput
				p := svc.NewDescribeGlobalNetworksPaginator(client, input)
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
		"disassociate-connect-peer": {
			Name:   "disassociate-connect-peer",
			Fields: fields_disassociate_connect_peer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateConnectPeerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_connect_peer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateConnectPeer(ctx, input)
			},
		},
		"disassociate-customer-gateway": {
			Name:   "disassociate-customer-gateway",
			Fields: fields_disassociate_customer_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateCustomerGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_customer_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateCustomerGateway(ctx, input)
			},
		},
		"disassociate-link": {
			Name:   "disassociate-link",
			Fields: fields_disassociate_link,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateLinkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_link, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateLink(ctx, input)
			},
		},
		"disassociate-transit-gateway-connect-peer": {
			Name:   "disassociate-transit-gateway-connect-peer",
			Fields: fields_disassociate_transit_gateway_connect_peer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateTransitGatewayConnectPeerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_transit_gateway_connect_peer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateTransitGatewayConnectPeer(ctx, input)
			},
		},
		"execute-core-network-change-set": {
			Name:   "execute-core-network-change-set",
			Fields: fields_execute_core_network_change_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExecuteCoreNetworkChangeSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_execute_core_network_change_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExecuteCoreNetworkChangeSet(ctx, input)
			},
		},
		"get-connect-attachment": {
			Name:   "get-connect-attachment",
			Fields: fields_get_connect_attachment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConnectAttachmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_connect_attachment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConnectAttachment(ctx, input)
			},
		},
		"get-connect-peer": {
			Name:   "get-connect-peer",
			Fields: fields_get_connect_peer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConnectPeerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_connect_peer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConnectPeer(ctx, input)
			},
		},
		"get-connect-peer-associations": {
			Name:   "get-connect-peer-associations",
			Fields: fields_get_connect_peer_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConnectPeerAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_connect_peer_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetConnectPeerAssociations(ctx, input)
				}
				var results []*svc.GetConnectPeerAssociationsOutput
				p := svc.NewGetConnectPeerAssociationsPaginator(client, input)
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
		"get-connections": {
			Name:   "get-connections",
			Fields: fields_get_connections,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConnectionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_connections, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetConnections(ctx, input)
				}
				var results []*svc.GetConnectionsOutput
				p := svc.NewGetConnectionsPaginator(client, input)
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
		"get-core-network": {
			Name:   "get-core-network",
			Fields: fields_get_core_network,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCoreNetworkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_core_network, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCoreNetwork(ctx, input)
			},
		},
		"get-core-network-change-events": {
			Name:   "get-core-network-change-events",
			Fields: fields_get_core_network_change_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCoreNetworkChangeEventsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_core_network_change_events, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetCoreNetworkChangeEvents(ctx, input)
				}
				var results []*svc.GetCoreNetworkChangeEventsOutput
				p := svc.NewGetCoreNetworkChangeEventsPaginator(client, input)
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
		"get-core-network-change-set": {
			Name:   "get-core-network-change-set",
			Fields: fields_get_core_network_change_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCoreNetworkChangeSetInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_core_network_change_set, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetCoreNetworkChangeSet(ctx, input)
				}
				var results []*svc.GetCoreNetworkChangeSetOutput
				p := svc.NewGetCoreNetworkChangeSetPaginator(client, input)
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
		"get-core-network-policy": {
			Name:   "get-core-network-policy",
			Fields: fields_get_core_network_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCoreNetworkPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_core_network_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCoreNetworkPolicy(ctx, input)
			},
		},
		"get-customer-gateway-associations": {
			Name:   "get-customer-gateway-associations",
			Fields: fields_get_customer_gateway_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCustomerGatewayAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_customer_gateway_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetCustomerGatewayAssociations(ctx, input)
				}
				var results []*svc.GetCustomerGatewayAssociationsOutput
				p := svc.NewGetCustomerGatewayAssociationsPaginator(client, input)
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
		"get-devices": {
			Name:   "get-devices",
			Fields: fields_get_devices,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDevicesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_devices, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetDevices(ctx, input)
				}
				var results []*svc.GetDevicesOutput
				p := svc.NewGetDevicesPaginator(client, input)
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
		"get-direct-connect-gateway-attachment": {
			Name:   "get-direct-connect-gateway-attachment",
			Fields: fields_get_direct_connect_gateway_attachment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDirectConnectGatewayAttachmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_direct_connect_gateway_attachment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDirectConnectGatewayAttachment(ctx, input)
			},
		},
		"get-link-associations": {
			Name:   "get-link-associations",
			Fields: fields_get_link_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLinkAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_link_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetLinkAssociations(ctx, input)
				}
				var results []*svc.GetLinkAssociationsOutput
				p := svc.NewGetLinkAssociationsPaginator(client, input)
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
		"get-links": {
			Name:   "get-links",
			Fields: fields_get_links,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLinksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_links, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetLinks(ctx, input)
				}
				var results []*svc.GetLinksOutput
				p := svc.NewGetLinksPaginator(client, input)
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
		"get-network-resource-counts": {
			Name:   "get-network-resource-counts",
			Fields: fields_get_network_resource_counts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetNetworkResourceCountsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_network_resource_counts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetNetworkResourceCounts(ctx, input)
				}
				var results []*svc.GetNetworkResourceCountsOutput
				p := svc.NewGetNetworkResourceCountsPaginator(client, input)
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
		"get-network-resource-relationships": {
			Name:   "get-network-resource-relationships",
			Fields: fields_get_network_resource_relationships,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetNetworkResourceRelationshipsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_network_resource_relationships, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetNetworkResourceRelationships(ctx, input)
				}
				var results []*svc.GetNetworkResourceRelationshipsOutput
				p := svc.NewGetNetworkResourceRelationshipsPaginator(client, input)
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
		"get-network-resources": {
			Name:   "get-network-resources",
			Fields: fields_get_network_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetNetworkResourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_network_resources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetNetworkResources(ctx, input)
				}
				var results []*svc.GetNetworkResourcesOutput
				p := svc.NewGetNetworkResourcesPaginator(client, input)
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
		"get-network-routes": {
			Name:   "get-network-routes",
			Fields: fields_get_network_routes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetNetworkRoutesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_network_routes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetNetworkRoutes(ctx, input)
			},
		},
		"get-network-telemetry": {
			Name:   "get-network-telemetry",
			Fields: fields_get_network_telemetry,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetNetworkTelemetryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_network_telemetry, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetNetworkTelemetry(ctx, input)
				}
				var results []*svc.GetNetworkTelemetryOutput
				p := svc.NewGetNetworkTelemetryPaginator(client, input)
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
		"get-resource-policy": {
			Name:   "get-resource-policy",
			Fields: fields_get_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResourcePolicy(ctx, input)
			},
		},
		"get-route-analysis": {
			Name:   "get-route-analysis",
			Fields: fields_get_route_analysis,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRouteAnalysisInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_route_analysis, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRouteAnalysis(ctx, input)
			},
		},
		"get-site-to-site-vpn-attachment": {
			Name:   "get-site-to-site-vpn-attachment",
			Fields: fields_get_site_to_site_vpn_attachment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSiteToSiteVpnAttachmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_site_to_site_vpn_attachment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSiteToSiteVpnAttachment(ctx, input)
			},
		},
		"get-sites": {
			Name:   "get-sites",
			Fields: fields_get_sites,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSitesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_sites, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetSites(ctx, input)
				}
				var results []*svc.GetSitesOutput
				p := svc.NewGetSitesPaginator(client, input)
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
		"get-transit-gateway-connect-peer-associations": {
			Name:   "get-transit-gateway-connect-peer-associations",
			Fields: fields_get_transit_gateway_connect_peer_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTransitGatewayConnectPeerAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_transit_gateway_connect_peer_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetTransitGatewayConnectPeerAssociations(ctx, input)
				}
				var results []*svc.GetTransitGatewayConnectPeerAssociationsOutput
				p := svc.NewGetTransitGatewayConnectPeerAssociationsPaginator(client, input)
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
		"get-transit-gateway-peering": {
			Name:   "get-transit-gateway-peering",
			Fields: fields_get_transit_gateway_peering,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTransitGatewayPeeringInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_transit_gateway_peering, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTransitGatewayPeering(ctx, input)
			},
		},
		"get-transit-gateway-registrations": {
			Name:   "get-transit-gateway-registrations",
			Fields: fields_get_transit_gateway_registrations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTransitGatewayRegistrationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_transit_gateway_registrations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetTransitGatewayRegistrations(ctx, input)
				}
				var results []*svc.GetTransitGatewayRegistrationsOutput
				p := svc.NewGetTransitGatewayRegistrationsPaginator(client, input)
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
		"get-transit-gateway-route-table-attachment": {
			Name:   "get-transit-gateway-route-table-attachment",
			Fields: fields_get_transit_gateway_route_table_attachment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTransitGatewayRouteTableAttachmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_transit_gateway_route_table_attachment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTransitGatewayRouteTableAttachment(ctx, input)
			},
		},
		"get-vpc-attachment": {
			Name:   "get-vpc-attachment",
			Fields: fields_get_vpc_attachment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetVpcAttachmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_vpc_attachment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetVpcAttachment(ctx, input)
			},
		},
		"list-attachment-routing-policy-associations": {
			Name:   "list-attachment-routing-policy-associations",
			Fields: fields_list_attachment_routing_policy_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAttachmentRoutingPolicyAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_attachment_routing_policy_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAttachmentRoutingPolicyAssociations(ctx, input)
				}
				var results []*svc.ListAttachmentRoutingPolicyAssociationsOutput
				p := svc.NewListAttachmentRoutingPolicyAssociationsPaginator(client, input)
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
		"list-attachments": {
			Name:   "list-attachments",
			Fields: fields_list_attachments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAttachmentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_attachments, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAttachments(ctx, input)
				}
				var results []*svc.ListAttachmentsOutput
				p := svc.NewListAttachmentsPaginator(client, input)
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
		"list-connect-peers": {
			Name:   "list-connect-peers",
			Fields: fields_list_connect_peers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConnectPeersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_connect_peers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListConnectPeers(ctx, input)
				}
				var results []*svc.ListConnectPeersOutput
				p := svc.NewListConnectPeersPaginator(client, input)
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
		"list-core-network-policy-versions": {
			Name:   "list-core-network-policy-versions",
			Fields: fields_list_core_network_policy_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCoreNetworkPolicyVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_core_network_policy_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCoreNetworkPolicyVersions(ctx, input)
				}
				var results []*svc.ListCoreNetworkPolicyVersionsOutput
				p := svc.NewListCoreNetworkPolicyVersionsPaginator(client, input)
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
		"list-core-network-prefix-list-associations": {
			Name:   "list-core-network-prefix-list-associations",
			Fields: fields_list_core_network_prefix_list_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCoreNetworkPrefixListAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_core_network_prefix_list_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCoreNetworkPrefixListAssociations(ctx, input)
				}
				var results []*svc.ListCoreNetworkPrefixListAssociationsOutput
				p := svc.NewListCoreNetworkPrefixListAssociationsPaginator(client, input)
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
		"list-core-network-routing-information": {
			Name:   "list-core-network-routing-information",
			Fields: fields_list_core_network_routing_information,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCoreNetworkRoutingInformationInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_core_network_routing_information, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCoreNetworkRoutingInformation(ctx, input)
				}
				var results []*svc.ListCoreNetworkRoutingInformationOutput
				p := svc.NewListCoreNetworkRoutingInformationPaginator(client, input)
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
		"list-core-networks": {
			Name:   "list-core-networks",
			Fields: fields_list_core_networks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCoreNetworksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_core_networks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCoreNetworks(ctx, input)
				}
				var results []*svc.ListCoreNetworksOutput
				p := svc.NewListCoreNetworksPaginator(client, input)
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
		"list-organization-service-access-status": {
			Name:   "list-organization-service-access-status",
			Fields: fields_list_organization_service_access_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOrganizationServiceAccessStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_organization_service_access_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListOrganizationServiceAccessStatus(ctx, input)
			},
		},
		"list-peerings": {
			Name:   "list-peerings",
			Fields: fields_list_peerings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPeeringsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_peerings, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPeerings(ctx, input)
				}
				var results []*svc.ListPeeringsOutput
				p := svc.NewListPeeringsPaginator(client, input)
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
		"put-attachment-routing-policy-label": {
			Name:   "put-attachment-routing-policy-label",
			Fields: fields_put_attachment_routing_policy_label,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutAttachmentRoutingPolicyLabelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_attachment_routing_policy_label, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutAttachmentRoutingPolicyLabel(ctx, input)
			},
		},
		"put-core-network-policy": {
			Name:   "put-core-network-policy",
			Fields: fields_put_core_network_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutCoreNetworkPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_core_network_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutCoreNetworkPolicy(ctx, input)
			},
		},
		"put-resource-policy": {
			Name:   "put-resource-policy",
			Fields: fields_put_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutResourcePolicy(ctx, input)
			},
		},
		"register-transit-gateway": {
			Name:   "register-transit-gateway",
			Fields: fields_register_transit_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterTransitGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_transit_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterTransitGateway(ctx, input)
			},
		},
		"reject-attachment": {
			Name:   "reject-attachment",
			Fields: fields_reject_attachment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RejectAttachmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reject_attachment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RejectAttachment(ctx, input)
			},
		},
		"remove-attachment-routing-policy-label": {
			Name:   "remove-attachment-routing-policy-label",
			Fields: fields_remove_attachment_routing_policy_label,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveAttachmentRoutingPolicyLabelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_attachment_routing_policy_label, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveAttachmentRoutingPolicyLabel(ctx, input)
			},
		},
		"restore-core-network-policy-version": {
			Name:   "restore-core-network-policy-version",
			Fields: fields_restore_core_network_policy_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RestoreCoreNetworkPolicyVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_restore_core_network_policy_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RestoreCoreNetworkPolicyVersion(ctx, input)
			},
		},
		"start-organization-service-access-update": {
			Name:   "start-organization-service-access-update",
			Fields: fields_start_organization_service_access_update,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartOrganizationServiceAccessUpdateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_organization_service_access_update, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartOrganizationServiceAccessUpdate(ctx, input)
			},
		},
		"start-route-analysis": {
			Name:   "start-route-analysis",
			Fields: fields_start_route_analysis,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartRouteAnalysisInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_route_analysis, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartRouteAnalysis(ctx, input)
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
		"update-connection": {
			Name:   "update-connection",
			Fields: fields_update_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateConnection(ctx, input)
			},
		},
		"update-core-network": {
			Name:   "update-core-network",
			Fields: fields_update_core_network,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCoreNetworkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_core_network, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCoreNetwork(ctx, input)
			},
		},
		"update-device": {
			Name:   "update-device",
			Fields: fields_update_device,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDeviceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_device, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDevice(ctx, input)
			},
		},
		"update-direct-connect-gateway-attachment": {
			Name:   "update-direct-connect-gateway-attachment",
			Fields: fields_update_direct_connect_gateway_attachment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDirectConnectGatewayAttachmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_direct_connect_gateway_attachment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDirectConnectGatewayAttachment(ctx, input)
			},
		},
		"update-global-network": {
			Name:   "update-global-network",
			Fields: fields_update_global_network,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateGlobalNetworkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_global_network, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateGlobalNetwork(ctx, input)
			},
		},
		"update-link": {
			Name:   "update-link",
			Fields: fields_update_link,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateLinkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_link, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateLink(ctx, input)
			},
		},
		"update-network-resource-metadata": {
			Name:   "update-network-resource-metadata",
			Fields: fields_update_network_resource_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateNetworkResourceMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_network_resource_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateNetworkResourceMetadata(ctx, input)
			},
		},
		"update-site": {
			Name:   "update-site",
			Fields: fields_update_site,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSiteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_site, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSite(ctx, input)
			},
		},
		"update-vpc-attachment": {
			Name:   "update-vpc-attachment",
			Fields: fields_update_vpc_attachment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateVpcAttachmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_vpc_attachment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateVpcAttachment(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("networkmanager", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
