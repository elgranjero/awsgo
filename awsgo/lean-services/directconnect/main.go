package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/directconnect"
)

var fields_accept_direct_connect_gateway_association_proposal = []leanruntime.Field{
	{Name: "AssociatedGatewayOwnerAccount", Flag: "associated-gateway-owner-account", Type: "*string", Required: true},
	{Name: "DirectConnectGatewayId", Flag: "direct-connect-gateway-id", Type: "*string", Required: true},
	{Name: "OverrideAllowedPrefixesToDirectConnectGateway", Flag: "override-allowed-prefixes-to-direct-connect-gateway", Type: "[]types.RouteFilterPrefix", Required: false},
	{Name: "ProposalId", Flag: "proposal-id", Type: "*string", Required: true},
}

var fields_allocate_connection_on_interconnect = []leanruntime.Field{
	{Name: "Bandwidth", Flag: "bandwidth", Type: "*string", Required: true},
	{Name: "ConnectionName", Flag: "connection-name", Type: "*string", Required: true},
	{Name: "InterconnectId", Flag: "interconnect-id", Type: "*string", Required: true},
	{Name: "OwnerAccount", Flag: "owner-account", Type: "*string", Required: true},
	{Name: "Vlan", Flag: "vlan", Type: "int32", Required: true},
}

var fields_allocate_hosted_connection = []leanruntime.Field{
	{Name: "Bandwidth", Flag: "bandwidth", Type: "*string", Required: true},
	{Name: "ConnectionId", Flag: "connection-id", Type: "*string", Required: true},
	{Name: "ConnectionName", Flag: "connection-name", Type: "*string", Required: true},
	{Name: "OwnerAccount", Flag: "owner-account", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Vlan", Flag: "vlan", Type: "int32", Required: true},
}

var fields_allocate_private_virtual_interface = []leanruntime.Field{
	{Name: "ConnectionId", Flag: "connection-id", Type: "*string", Required: true},
	{Name: "NewPrivateVirtualInterfaceAllocation", Flag: "new-private-virtual-interface-allocation", Type: "*types.NewPrivateVirtualInterfaceAllocation", Required: true},
	{Name: "OwnerAccount", Flag: "owner-account", Type: "*string", Required: true},
}

var fields_allocate_public_virtual_interface = []leanruntime.Field{
	{Name: "ConnectionId", Flag: "connection-id", Type: "*string", Required: true},
	{Name: "NewPublicVirtualInterfaceAllocation", Flag: "new-public-virtual-interface-allocation", Type: "*types.NewPublicVirtualInterfaceAllocation", Required: true},
	{Name: "OwnerAccount", Flag: "owner-account", Type: "*string", Required: true},
}

var fields_allocate_transit_virtual_interface = []leanruntime.Field{
	{Name: "ConnectionId", Flag: "connection-id", Type: "*string", Required: true},
	{Name: "NewTransitVirtualInterfaceAllocation", Flag: "new-transit-virtual-interface-allocation", Type: "*types.NewTransitVirtualInterfaceAllocation", Required: true},
	{Name: "OwnerAccount", Flag: "owner-account", Type: "*string", Required: true},
}

var fields_associate_connection_with_lag = []leanruntime.Field{
	{Name: "ConnectionId", Flag: "connection-id", Type: "*string", Required: true},
	{Name: "LagId", Flag: "lag-id", Type: "*string", Required: true},
}

var fields_associate_hosted_connection = []leanruntime.Field{
	{Name: "ConnectionId", Flag: "connection-id", Type: "*string", Required: true},
	{Name: "ParentConnectionId", Flag: "parent-connection-id", Type: "*string", Required: true},
}

var fields_associate_mac_sec_key = []leanruntime.Field{
	{Name: "Cak", Flag: "cak", Type: "*string", Required: false},
	{Name: "Ckn", Flag: "ckn", Type: "*string", Required: false},
	{Name: "ConnectionId", Flag: "connection-id", Type: "*string", Required: true},
	{Name: "SecretARN", Flag: "secret-arn", Type: "*string", Required: false},
}

var fields_associate_virtual_interface = []leanruntime.Field{
	{Name: "ConnectionId", Flag: "connection-id", Type: "*string", Required: true},
	{Name: "VirtualInterfaceId", Flag: "virtual-interface-id", Type: "*string", Required: true},
}

var fields_confirm_connection = []leanruntime.Field{
	{Name: "ConnectionId", Flag: "connection-id", Type: "*string", Required: true},
}

var fields_confirm_customer_agreement = []leanruntime.Field{
	{Name: "AgreementName", Flag: "agreement-name", Type: "*string", Required: false},
}

var fields_confirm_private_virtual_interface = []leanruntime.Field{
	{Name: "DirectConnectGatewayId", Flag: "direct-connect-gateway-id", Type: "*string", Required: false},
	{Name: "VirtualGatewayId", Flag: "virtual-gateway-id", Type: "*string", Required: false},
	{Name: "VirtualInterfaceId", Flag: "virtual-interface-id", Type: "*string", Required: true},
}

var fields_confirm_public_virtual_interface = []leanruntime.Field{
	{Name: "VirtualInterfaceId", Flag: "virtual-interface-id", Type: "*string", Required: true},
}

var fields_confirm_transit_virtual_interface = []leanruntime.Field{
	{Name: "DirectConnectGatewayId", Flag: "direct-connect-gateway-id", Type: "*string", Required: true},
	{Name: "VirtualInterfaceId", Flag: "virtual-interface-id", Type: "*string", Required: true},
}

var fields_create_bgp_peer = []leanruntime.Field{
	{Name: "NewBGPPeer", Flag: "new-bgp-peer", Type: "*types.NewBGPPeer", Required: false},
	{Name: "VirtualInterfaceId", Flag: "virtual-interface-id", Type: "*string", Required: false},
}

var fields_create_connection = []leanruntime.Field{
	{Name: "Bandwidth", Flag: "bandwidth", Type: "*string", Required: true},
	{Name: "ConnectionName", Flag: "connection-name", Type: "*string", Required: true},
	{Name: "LagId", Flag: "lag-id", Type: "*string", Required: false},
	{Name: "Location", Flag: "location", Type: "*string", Required: true},
	{Name: "ProviderName", Flag: "provider-name", Type: "*string", Required: false},
	{Name: "RequestMACSec", Flag: "request-mac-sec", Type: "*bool", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_direct_connect_gateway = []leanruntime.Field{
	{Name: "AmazonSideAsn", Flag: "amazon-side-asn", Type: "*int64", Required: false},
	{Name: "DirectConnectGatewayName", Flag: "direct-connect-gateway-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_direct_connect_gateway_association = []leanruntime.Field{
	{Name: "AddAllowedPrefixesToDirectConnectGateway", Flag: "add-allowed-prefixes-to-direct-connect-gateway", Type: "[]types.RouteFilterPrefix", Required: false},
	{Name: "DirectConnectGatewayId", Flag: "direct-connect-gateway-id", Type: "*string", Required: true},
	{Name: "GatewayId", Flag: "gateway-id", Type: "*string", Required: false},
	{Name: "VirtualGatewayId", Flag: "virtual-gateway-id", Type: "*string", Required: false},
}

var fields_create_direct_connect_gateway_association_proposal = []leanruntime.Field{
	{Name: "AddAllowedPrefixesToDirectConnectGateway", Flag: "add-allowed-prefixes-to-direct-connect-gateway", Type: "[]types.RouteFilterPrefix", Required: false},
	{Name: "DirectConnectGatewayId", Flag: "direct-connect-gateway-id", Type: "*string", Required: true},
	{Name: "DirectConnectGatewayOwnerAccount", Flag: "direct-connect-gateway-owner-account", Type: "*string", Required: true},
	{Name: "GatewayId", Flag: "gateway-id", Type: "*string", Required: true},
	{Name: "RemoveAllowedPrefixesToDirectConnectGateway", Flag: "remove-allowed-prefixes-to-direct-connect-gateway", Type: "[]types.RouteFilterPrefix", Required: false},
}

var fields_create_interconnect = []leanruntime.Field{
	{Name: "Bandwidth", Flag: "bandwidth", Type: "*string", Required: true},
	{Name: "InterconnectName", Flag: "interconnect-name", Type: "*string", Required: true},
	{Name: "LagId", Flag: "lag-id", Type: "*string", Required: false},
	{Name: "Location", Flag: "location", Type: "*string", Required: true},
	{Name: "ProviderName", Flag: "provider-name", Type: "*string", Required: false},
	{Name: "RequestMACSec", Flag: "request-mac-sec", Type: "*bool", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_lag = []leanruntime.Field{
	{Name: "ChildConnectionTags", Flag: "child-connection-tags", Type: "[]types.Tag", Required: false},
	{Name: "ConnectionId", Flag: "connection-id", Type: "*string", Required: false},
	{Name: "ConnectionsBandwidth", Flag: "connections-bandwidth", Type: "*string", Required: true},
	{Name: "LagName", Flag: "lag-name", Type: "*string", Required: true},
	{Name: "Location", Flag: "location", Type: "*string", Required: true},
	{Name: "NumberOfConnections", Flag: "number-of-connections", Type: "int32", Required: true},
	{Name: "ProviderName", Flag: "provider-name", Type: "*string", Required: false},
	{Name: "RequestMACSec", Flag: "request-mac-sec", Type: "*bool", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_private_virtual_interface = []leanruntime.Field{
	{Name: "ConnectionId", Flag: "connection-id", Type: "*string", Required: true},
	{Name: "NewPrivateVirtualInterface", Flag: "new-private-virtual-interface", Type: "*types.NewPrivateVirtualInterface", Required: true},
}

var fields_create_public_virtual_interface = []leanruntime.Field{
	{Name: "ConnectionId", Flag: "connection-id", Type: "*string", Required: true},
	{Name: "NewPublicVirtualInterface", Flag: "new-public-virtual-interface", Type: "*types.NewPublicVirtualInterface", Required: true},
}

var fields_create_transit_virtual_interface = []leanruntime.Field{
	{Name: "ConnectionId", Flag: "connection-id", Type: "*string", Required: true},
	{Name: "NewTransitVirtualInterface", Flag: "new-transit-virtual-interface", Type: "*types.NewTransitVirtualInterface", Required: true},
}

var fields_delete_bgp_peer = []leanruntime.Field{
	{Name: "Asn", Flag: "asn", Type: "int32", Required: false},
	{Name: "AsnLong", Flag: "asn-long", Type: "*int64", Required: false},
	{Name: "BgpPeerId", Flag: "bgp-peer-id", Type: "*string", Required: false},
	{Name: "CustomerAddress", Flag: "customer-address", Type: "*string", Required: false},
	{Name: "VirtualInterfaceId", Flag: "virtual-interface-id", Type: "*string", Required: false},
}

var fields_delete_connection = []leanruntime.Field{
	{Name: "ConnectionId", Flag: "connection-id", Type: "*string", Required: true},
}

var fields_delete_direct_connect_gateway = []leanruntime.Field{
	{Name: "DirectConnectGatewayId", Flag: "direct-connect-gateway-id", Type: "*string", Required: true},
}

var fields_delete_direct_connect_gateway_association = []leanruntime.Field{
	{Name: "AssociationId", Flag: "association-id", Type: "*string", Required: false},
	{Name: "DirectConnectGatewayId", Flag: "direct-connect-gateway-id", Type: "*string", Required: false},
	{Name: "VirtualGatewayId", Flag: "virtual-gateway-id", Type: "*string", Required: false},
}

var fields_delete_direct_connect_gateway_association_proposal = []leanruntime.Field{
	{Name: "ProposalId", Flag: "proposal-id", Type: "*string", Required: true},
}

var fields_delete_interconnect = []leanruntime.Field{
	{Name: "InterconnectId", Flag: "interconnect-id", Type: "*string", Required: true},
}

var fields_delete_lag = []leanruntime.Field{
	{Name: "LagId", Flag: "lag-id", Type: "*string", Required: true},
}

var fields_delete_virtual_interface = []leanruntime.Field{
	{Name: "VirtualInterfaceId", Flag: "virtual-interface-id", Type: "*string", Required: true},
}

var fields_describe_connection_loa = []leanruntime.Field{
	{Name: "ConnectionId", Flag: "connection-id", Type: "*string", Required: true},
	{Name: "LoaContentType", Flag: "loa-content-type", Type: "types.LoaContentType", Required: false},
	{Name: "ProviderName", Flag: "provider-name", Type: "*string", Required: false},
}

var fields_describe_connections = []leanruntime.Field{
	{Name: "ConnectionId", Flag: "connection-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_connections_on_interconnect = []leanruntime.Field{
	{Name: "InterconnectId", Flag: "interconnect-id", Type: "*string", Required: true},
}

var fields_describe_customer_metadata = []leanruntime.Field{}

var fields_describe_direct_connect_gateway_association_proposals = []leanruntime.Field{
	{Name: "AssociatedGatewayId", Flag: "associated-gateway-id", Type: "*string", Required: false},
	{Name: "DirectConnectGatewayId", Flag: "direct-connect-gateway-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProposalId", Flag: "proposal-id", Type: "*string", Required: false},
}

var fields_describe_direct_connect_gateway_associations = []leanruntime.Field{
	{Name: "AssociatedGatewayId", Flag: "associated-gateway-id", Type: "*string", Required: false},
	{Name: "AssociationId", Flag: "association-id", Type: "*string", Required: false},
	{Name: "DirectConnectGatewayId", Flag: "direct-connect-gateway-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "VirtualGatewayId", Flag: "virtual-gateway-id", Type: "*string", Required: false},
}

var fields_describe_direct_connect_gateway_attachments = []leanruntime.Field{
	{Name: "DirectConnectGatewayId", Flag: "direct-connect-gateway-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "VirtualInterfaceId", Flag: "virtual-interface-id", Type: "*string", Required: false},
}

var fields_describe_direct_connect_gateways = []leanruntime.Field{
	{Name: "DirectConnectGatewayId", Flag: "direct-connect-gateway-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_hosted_connections = []leanruntime.Field{
	{Name: "ConnectionId", Flag: "connection-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_interconnect_loa = []leanruntime.Field{
	{Name: "InterconnectId", Flag: "interconnect-id", Type: "*string", Required: true},
	{Name: "LoaContentType", Flag: "loa-content-type", Type: "types.LoaContentType", Required: false},
	{Name: "ProviderName", Flag: "provider-name", Type: "*string", Required: false},
}

var fields_describe_interconnects = []leanruntime.Field{
	{Name: "InterconnectId", Flag: "interconnect-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_lags = []leanruntime.Field{
	{Name: "LagId", Flag: "lag-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_loa = []leanruntime.Field{
	{Name: "ConnectionId", Flag: "connection-id", Type: "*string", Required: true},
	{Name: "LoaContentType", Flag: "loa-content-type", Type: "types.LoaContentType", Required: false},
	{Name: "ProviderName", Flag: "provider-name", Type: "*string", Required: false},
}

var fields_describe_locations = []leanruntime.Field{}

var fields_describe_router_configuration = []leanruntime.Field{
	{Name: "RouterTypeIdentifier", Flag: "router-type-identifier", Type: "*string", Required: false},
	{Name: "VirtualInterfaceId", Flag: "virtual-interface-id", Type: "*string", Required: true},
}

var fields_describe_tags = []leanruntime.Field{
	{Name: "ResourceArns", Flag: "resource-arns", Type: "[]string", Required: true},
}

var fields_describe_virtual_gateways = []leanruntime.Field{}

var fields_describe_virtual_interfaces = []leanruntime.Field{
	{Name: "ConnectionId", Flag: "connection-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "VirtualInterfaceId", Flag: "virtual-interface-id", Type: "*string", Required: false},
}

var fields_disassociate_connection_from_lag = []leanruntime.Field{
	{Name: "ConnectionId", Flag: "connection-id", Type: "*string", Required: true},
	{Name: "LagId", Flag: "lag-id", Type: "*string", Required: true},
}

var fields_disassociate_mac_sec_key = []leanruntime.Field{
	{Name: "ConnectionId", Flag: "connection-id", Type: "*string", Required: true},
	{Name: "SecretARN", Flag: "secret-arn", Type: "*string", Required: true},
}

var fields_list_virtual_interface_test_history = []leanruntime.Field{
	{Name: "BgpPeers", Flag: "bgp-peers", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "*string", Required: false},
	{Name: "TestId", Flag: "test-id", Type: "*string", Required: false},
	{Name: "VirtualInterfaceId", Flag: "virtual-interface-id", Type: "*string", Required: false},
}

var fields_start_bgp_failover_test = []leanruntime.Field{
	{Name: "BgpPeers", Flag: "bgp-peers", Type: "[]string", Required: false},
	{Name: "TestDurationInMinutes", Flag: "test-duration-in-minutes", Type: "*int32", Required: false},
	{Name: "VirtualInterfaceId", Flag: "virtual-interface-id", Type: "*string", Required: true},
}

var fields_stop_bgp_failover_test = []leanruntime.Field{
	{Name: "VirtualInterfaceId", Flag: "virtual-interface-id", Type: "*string", Required: true},
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
	{Name: "ConnectionId", Flag: "connection-id", Type: "*string", Required: true},
	{Name: "ConnectionName", Flag: "connection-name", Type: "*string", Required: false},
	{Name: "EncryptionMode", Flag: "encryption-mode", Type: "*string", Required: false},
}

var fields_update_direct_connect_gateway = []leanruntime.Field{
	{Name: "DirectConnectGatewayId", Flag: "direct-connect-gateway-id", Type: "*string", Required: true},
	{Name: "NewDirectConnectGatewayName", Flag: "new-direct-connect-gateway-name", Type: "*string", Required: true},
}

var fields_update_direct_connect_gateway_association = []leanruntime.Field{
	{Name: "AddAllowedPrefixesToDirectConnectGateway", Flag: "add-allowed-prefixes-to-direct-connect-gateway", Type: "[]types.RouteFilterPrefix", Required: false},
	{Name: "AssociationId", Flag: "association-id", Type: "*string", Required: false},
	{Name: "RemoveAllowedPrefixesToDirectConnectGateway", Flag: "remove-allowed-prefixes-to-direct-connect-gateway", Type: "[]types.RouteFilterPrefix", Required: false},
}

var fields_update_lag = []leanruntime.Field{
	{Name: "EncryptionMode", Flag: "encryption-mode", Type: "*string", Required: false},
	{Name: "LagId", Flag: "lag-id", Type: "*string", Required: true},
	{Name: "LagName", Flag: "lag-name", Type: "*string", Required: false},
	{Name: "MinimumLinks", Flag: "minimum-links", Type: "int32", Required: false},
}

var fields_update_virtual_interface_attributes = []leanruntime.Field{
	{Name: "EnableSiteLink", Flag: "enable-site-link", Type: "*bool", Required: false},
	{Name: "Mtu", Flag: "mtu", Type: "*int32", Required: false},
	{Name: "VirtualInterfaceId", Flag: "virtual-interface-id", Type: "*string", Required: true},
	{Name: "VirtualInterfaceName", Flag: "virtual-interface-name", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"accept-direct-connect-gateway-association-proposal": {
			Name:   "accept-direct-connect-gateway-association-proposal",
			Fields: fields_accept_direct_connect_gateway_association_proposal,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AcceptDirectConnectGatewayAssociationProposalInput{}
				if _, err := leanruntime.ApplyInput(input, fields_accept_direct_connect_gateway_association_proposal, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AcceptDirectConnectGatewayAssociationProposal(ctx, input)
			},
		},
		"allocate-connection-on-interconnect": {
			Name:   "allocate-connection-on-interconnect",
			Fields: fields_allocate_connection_on_interconnect,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AllocateConnectionOnInterconnectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_allocate_connection_on_interconnect, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AllocateConnectionOnInterconnect(ctx, input)
			},
		},
		"allocate-hosted-connection": {
			Name:   "allocate-hosted-connection",
			Fields: fields_allocate_hosted_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AllocateHostedConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_allocate_hosted_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AllocateHostedConnection(ctx, input)
			},
		},
		"allocate-private-virtual-interface": {
			Name:   "allocate-private-virtual-interface",
			Fields: fields_allocate_private_virtual_interface,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AllocatePrivateVirtualInterfaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_allocate_private_virtual_interface, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AllocatePrivateVirtualInterface(ctx, input)
			},
		},
		"allocate-public-virtual-interface": {
			Name:   "allocate-public-virtual-interface",
			Fields: fields_allocate_public_virtual_interface,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AllocatePublicVirtualInterfaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_allocate_public_virtual_interface, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AllocatePublicVirtualInterface(ctx, input)
			},
		},
		"allocate-transit-virtual-interface": {
			Name:   "allocate-transit-virtual-interface",
			Fields: fields_allocate_transit_virtual_interface,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AllocateTransitVirtualInterfaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_allocate_transit_virtual_interface, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AllocateTransitVirtualInterface(ctx, input)
			},
		},
		"associate-connection-with-lag": {
			Name:   "associate-connection-with-lag",
			Fields: fields_associate_connection_with_lag,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateConnectionWithLagInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_connection_with_lag, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateConnectionWithLag(ctx, input)
			},
		},
		"associate-hosted-connection": {
			Name:   "associate-hosted-connection",
			Fields: fields_associate_hosted_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateHostedConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_hosted_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateHostedConnection(ctx, input)
			},
		},
		"associate-mac-sec-key": {
			Name:   "associate-mac-sec-key",
			Fields: fields_associate_mac_sec_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateMacSecKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_mac_sec_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateMacSecKey(ctx, input)
			},
		},
		"associate-virtual-interface": {
			Name:   "associate-virtual-interface",
			Fields: fields_associate_virtual_interface,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateVirtualInterfaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_virtual_interface, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateVirtualInterface(ctx, input)
			},
		},
		"confirm-connection": {
			Name:   "confirm-connection",
			Fields: fields_confirm_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ConfirmConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_confirm_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ConfirmConnection(ctx, input)
			},
		},
		"confirm-customer-agreement": {
			Name:   "confirm-customer-agreement",
			Fields: fields_confirm_customer_agreement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ConfirmCustomerAgreementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_confirm_customer_agreement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ConfirmCustomerAgreement(ctx, input)
			},
		},
		"confirm-private-virtual-interface": {
			Name:   "confirm-private-virtual-interface",
			Fields: fields_confirm_private_virtual_interface,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ConfirmPrivateVirtualInterfaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_confirm_private_virtual_interface, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ConfirmPrivateVirtualInterface(ctx, input)
			},
		},
		"confirm-public-virtual-interface": {
			Name:   "confirm-public-virtual-interface",
			Fields: fields_confirm_public_virtual_interface,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ConfirmPublicVirtualInterfaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_confirm_public_virtual_interface, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ConfirmPublicVirtualInterface(ctx, input)
			},
		},
		"confirm-transit-virtual-interface": {
			Name:   "confirm-transit-virtual-interface",
			Fields: fields_confirm_transit_virtual_interface,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ConfirmTransitVirtualInterfaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_confirm_transit_virtual_interface, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ConfirmTransitVirtualInterface(ctx, input)
			},
		},
		"create-bgp-peer": {
			Name:   "create-bgp-peer",
			Fields: fields_create_bgp_peer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBGPPeerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_bgp_peer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBGPPeer(ctx, input)
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
		"create-direct-connect-gateway": {
			Name:   "create-direct-connect-gateway",
			Fields: fields_create_direct_connect_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDirectConnectGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_direct_connect_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDirectConnectGateway(ctx, input)
			},
		},
		"create-direct-connect-gateway-association": {
			Name:   "create-direct-connect-gateway-association",
			Fields: fields_create_direct_connect_gateway_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDirectConnectGatewayAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_direct_connect_gateway_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDirectConnectGatewayAssociation(ctx, input)
			},
		},
		"create-direct-connect-gateway-association-proposal": {
			Name:   "create-direct-connect-gateway-association-proposal",
			Fields: fields_create_direct_connect_gateway_association_proposal,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDirectConnectGatewayAssociationProposalInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_direct_connect_gateway_association_proposal, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDirectConnectGatewayAssociationProposal(ctx, input)
			},
		},
		"create-interconnect": {
			Name:   "create-interconnect",
			Fields: fields_create_interconnect,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateInterconnectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_interconnect, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateInterconnect(ctx, input)
			},
		},
		"create-lag": {
			Name:   "create-lag",
			Fields: fields_create_lag,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLagInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_lag, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLag(ctx, input)
			},
		},
		"create-private-virtual-interface": {
			Name:   "create-private-virtual-interface",
			Fields: fields_create_private_virtual_interface,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePrivateVirtualInterfaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_private_virtual_interface, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePrivateVirtualInterface(ctx, input)
			},
		},
		"create-public-virtual-interface": {
			Name:   "create-public-virtual-interface",
			Fields: fields_create_public_virtual_interface,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePublicVirtualInterfaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_public_virtual_interface, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePublicVirtualInterface(ctx, input)
			},
		},
		"create-transit-virtual-interface": {
			Name:   "create-transit-virtual-interface",
			Fields: fields_create_transit_virtual_interface,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTransitVirtualInterfaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_transit_virtual_interface, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTransitVirtualInterface(ctx, input)
			},
		},
		"delete-bgp-peer": {
			Name:   "delete-bgp-peer",
			Fields: fields_delete_bgp_peer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBGPPeerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_bgp_peer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBGPPeer(ctx, input)
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
		"delete-direct-connect-gateway": {
			Name:   "delete-direct-connect-gateway",
			Fields: fields_delete_direct_connect_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDirectConnectGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_direct_connect_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDirectConnectGateway(ctx, input)
			},
		},
		"delete-direct-connect-gateway-association": {
			Name:   "delete-direct-connect-gateway-association",
			Fields: fields_delete_direct_connect_gateway_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDirectConnectGatewayAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_direct_connect_gateway_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDirectConnectGatewayAssociation(ctx, input)
			},
		},
		"delete-direct-connect-gateway-association-proposal": {
			Name:   "delete-direct-connect-gateway-association-proposal",
			Fields: fields_delete_direct_connect_gateway_association_proposal,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDirectConnectGatewayAssociationProposalInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_direct_connect_gateway_association_proposal, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDirectConnectGatewayAssociationProposal(ctx, input)
			},
		},
		"delete-interconnect": {
			Name:   "delete-interconnect",
			Fields: fields_delete_interconnect,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteInterconnectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_interconnect, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteInterconnect(ctx, input)
			},
		},
		"delete-lag": {
			Name:   "delete-lag",
			Fields: fields_delete_lag,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLagInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_lag, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLag(ctx, input)
			},
		},
		"delete-virtual-interface": {
			Name:   "delete-virtual-interface",
			Fields: fields_delete_virtual_interface,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVirtualInterfaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_virtual_interface, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVirtualInterface(ctx, input)
			},
		},
		"describe-connection-loa": {
			Name:   "describe-connection-loa",
			Fields: fields_describe_connection_loa,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeConnectionLoaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_connection_loa, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeConnectionLoa(ctx, input)
			},
		},
		"describe-connections": {
			Name:   "describe-connections",
			Fields: fields_describe_connections,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeConnectionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_connections, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeConnections(ctx, input)
			},
		},
		"describe-connections-on-interconnect": {
			Name:   "describe-connections-on-interconnect",
			Fields: fields_describe_connections_on_interconnect,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeConnectionsOnInterconnectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_connections_on_interconnect, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeConnectionsOnInterconnect(ctx, input)
			},
		},
		"describe-customer-metadata": {
			Name:   "describe-customer-metadata",
			Fields: fields_describe_customer_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCustomerMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_customer_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeCustomerMetadata(ctx, input)
			},
		},
		"describe-direct-connect-gateway-association-proposals": {
			Name:   "describe-direct-connect-gateway-association-proposals",
			Fields: fields_describe_direct_connect_gateway_association_proposals,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDirectConnectGatewayAssociationProposalsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_direct_connect_gateway_association_proposals, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDirectConnectGatewayAssociationProposals(ctx, input)
			},
		},
		"describe-direct-connect-gateway-associations": {
			Name:   "describe-direct-connect-gateway-associations",
			Fields: fields_describe_direct_connect_gateway_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDirectConnectGatewayAssociationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_direct_connect_gateway_associations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDirectConnectGatewayAssociations(ctx, input)
			},
		},
		"describe-direct-connect-gateway-attachments": {
			Name:   "describe-direct-connect-gateway-attachments",
			Fields: fields_describe_direct_connect_gateway_attachments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDirectConnectGatewayAttachmentsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_direct_connect_gateway_attachments, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDirectConnectGatewayAttachments(ctx, input)
			},
		},
		"describe-direct-connect-gateways": {
			Name:   "describe-direct-connect-gateways",
			Fields: fields_describe_direct_connect_gateways,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDirectConnectGatewaysInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_direct_connect_gateways, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDirectConnectGateways(ctx, input)
			},
		},
		"describe-hosted-connections": {
			Name:   "describe-hosted-connections",
			Fields: fields_describe_hosted_connections,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeHostedConnectionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_hosted_connections, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeHostedConnections(ctx, input)
			},
		},
		"describe-interconnect-loa": {
			Name:   "describe-interconnect-loa",
			Fields: fields_describe_interconnect_loa,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInterconnectLoaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_interconnect_loa, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeInterconnectLoa(ctx, input)
			},
		},
		"describe-interconnects": {
			Name:   "describe-interconnects",
			Fields: fields_describe_interconnects,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInterconnectsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_interconnects, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeInterconnects(ctx, input)
			},
		},
		"describe-lags": {
			Name:   "describe-lags",
			Fields: fields_describe_lags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLagsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_lags, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeLags(ctx, input)
			},
		},
		"describe-loa": {
			Name:   "describe-loa",
			Fields: fields_describe_loa,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLoaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_loa, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeLoa(ctx, input)
			},
		},
		"describe-locations": {
			Name:   "describe-locations",
			Fields: fields_describe_locations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLocationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_locations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeLocations(ctx, input)
			},
		},
		"describe-router-configuration": {
			Name:   "describe-router-configuration",
			Fields: fields_describe_router_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRouterConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_router_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeRouterConfiguration(ctx, input)
			},
		},
		"describe-tags": {
			Name:   "describe-tags",
			Fields: fields_describe_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTagsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_tags, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTags(ctx, input)
			},
		},
		"describe-virtual-gateways": {
			Name:   "describe-virtual-gateways",
			Fields: fields_describe_virtual_gateways,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVirtualGatewaysInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_virtual_gateways, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeVirtualGateways(ctx, input)
			},
		},
		"describe-virtual-interfaces": {
			Name:   "describe-virtual-interfaces",
			Fields: fields_describe_virtual_interfaces,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVirtualInterfacesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_virtual_interfaces, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeVirtualInterfaces(ctx, input)
			},
		},
		"disassociate-connection-from-lag": {
			Name:   "disassociate-connection-from-lag",
			Fields: fields_disassociate_connection_from_lag,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateConnectionFromLagInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_connection_from_lag, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateConnectionFromLag(ctx, input)
			},
		},
		"disassociate-mac-sec-key": {
			Name:   "disassociate-mac-sec-key",
			Fields: fields_disassociate_mac_sec_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateMacSecKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_mac_sec_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateMacSecKey(ctx, input)
			},
		},
		"list-virtual-interface-test-history": {
			Name:   "list-virtual-interface-test-history",
			Fields: fields_list_virtual_interface_test_history,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVirtualInterfaceTestHistoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_virtual_interface_test_history, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListVirtualInterfaceTestHistory(ctx, input)
			},
		},
		"start-bgp-failover-test": {
			Name:   "start-bgp-failover-test",
			Fields: fields_start_bgp_failover_test,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartBgpFailoverTestInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_bgp_failover_test, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartBgpFailoverTest(ctx, input)
			},
		},
		"stop-bgp-failover-test": {
			Name:   "stop-bgp-failover-test",
			Fields: fields_stop_bgp_failover_test,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopBgpFailoverTestInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_bgp_failover_test, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopBgpFailoverTest(ctx, input)
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
		"update-direct-connect-gateway": {
			Name:   "update-direct-connect-gateway",
			Fields: fields_update_direct_connect_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDirectConnectGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_direct_connect_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDirectConnectGateway(ctx, input)
			},
		},
		"update-direct-connect-gateway-association": {
			Name:   "update-direct-connect-gateway-association",
			Fields: fields_update_direct_connect_gateway_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDirectConnectGatewayAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_direct_connect_gateway_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDirectConnectGatewayAssociation(ctx, input)
			},
		},
		"update-lag": {
			Name:   "update-lag",
			Fields: fields_update_lag,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateLagInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_lag, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateLag(ctx, input)
			},
		},
		"update-virtual-interface-attributes": {
			Name:   "update-virtual-interface-attributes",
			Fields: fields_update_virtual_interface_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateVirtualInterfaceAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_virtual_interface_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateVirtualInterfaceAttributes(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("directconnect", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
