package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/globalaccelerator"
)

var fields_add_custom_routing_endpoints = []leanruntime.Field{
	{Name: "EndpointConfigurations", Flag: "endpoint-configurations", Type: "[]types.CustomRoutingEndpointConfiguration", Required: true},
	{Name: "EndpointGroupArn", Flag: "endpoint-group-arn", Type: "*string", Required: true},
}

var fields_add_endpoints = []leanruntime.Field{
	{Name: "EndpointConfigurations", Flag: "endpoint-configurations", Type: "[]types.EndpointConfiguration", Required: true},
	{Name: "EndpointGroupArn", Flag: "endpoint-group-arn", Type: "*string", Required: true},
}

var fields_advertise_byoip_cidr = []leanruntime.Field{
	{Name: "Cidr", Flag: "cidr", Type: "*string", Required: true},
}

var fields_allow_custom_routing_traffic = []leanruntime.Field{
	{Name: "AllowAllTrafficToEndpoint", Flag: "allow-all-traffic-to-endpoint", Type: "*bool", Required: false},
	{Name: "DestinationAddresses", Flag: "destination-addresses", Type: "[]string", Required: false},
	{Name: "DestinationPorts", Flag: "destination-ports", Type: "[]int32", Required: false},
	{Name: "EndpointGroupArn", Flag: "endpoint-group-arn", Type: "*string", Required: true},
	{Name: "EndpointId", Flag: "endpoint-id", Type: "*string", Required: true},
}

var fields_create_accelerator = []leanruntime.Field{
	{Name: "Enabled", Flag: "enabled", Type: "*bool", Required: false},
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: true},
	{Name: "IpAddressType", Flag: "ip-address-type", Type: "types.IpAddressType", Required: false},
	{Name: "IpAddresses", Flag: "ip-addresses", Type: "[]string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_cross_account_attachment = []leanruntime.Field{
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Principals", Flag: "principals", Type: "[]string", Required: false},
	{Name: "Resources", Flag: "resources", Type: "[]types.Resource", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_custom_routing_accelerator = []leanruntime.Field{
	{Name: "Enabled", Flag: "enabled", Type: "*bool", Required: false},
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: true},
	{Name: "IpAddressType", Flag: "ip-address-type", Type: "types.IpAddressType", Required: false},
	{Name: "IpAddresses", Flag: "ip-addresses", Type: "[]string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_custom_routing_endpoint_group = []leanruntime.Field{
	{Name: "DestinationConfigurations", Flag: "destination-configurations", Type: "[]types.CustomRoutingDestinationConfiguration", Required: true},
	{Name: "EndpointGroupRegion", Flag: "endpoint-group-region", Type: "*string", Required: true},
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: true},
	{Name: "ListenerArn", Flag: "listener-arn", Type: "*string", Required: true},
}

var fields_create_custom_routing_listener = []leanruntime.Field{
	{Name: "AcceleratorArn", Flag: "accelerator-arn", Type: "*string", Required: true},
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: true},
	{Name: "PortRanges", Flag: "port-ranges", Type: "[]types.PortRange", Required: true},
}

var fields_create_endpoint_group = []leanruntime.Field{
	{Name: "EndpointConfigurations", Flag: "endpoint-configurations", Type: "[]types.EndpointConfiguration", Required: false},
	{Name: "EndpointGroupRegion", Flag: "endpoint-group-region", Type: "*string", Required: true},
	{Name: "HealthCheckIntervalSeconds", Flag: "health-check-interval-seconds", Type: "*int32", Required: false},
	{Name: "HealthCheckPath", Flag: "health-check-path", Type: "*string", Required: false},
	{Name: "HealthCheckPort", Flag: "health-check-port", Type: "*int32", Required: false},
	{Name: "HealthCheckProtocol", Flag: "health-check-protocol", Type: "types.HealthCheckProtocol", Required: false},
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: true},
	{Name: "ListenerArn", Flag: "listener-arn", Type: "*string", Required: true},
	{Name: "PortOverrides", Flag: "port-overrides", Type: "[]types.PortOverride", Required: false},
	{Name: "ThresholdCount", Flag: "threshold-count", Type: "*int32", Required: false},
	{Name: "TrafficDialPercentage", Flag: "traffic-dial-percentage", Type: "*float32", Required: false},
}

var fields_create_listener = []leanruntime.Field{
	{Name: "AcceleratorArn", Flag: "accelerator-arn", Type: "*string", Required: true},
	{Name: "ClientAffinity", Flag: "client-affinity", Type: "types.ClientAffinity", Required: false},
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: true},
	{Name: "PortRanges", Flag: "port-ranges", Type: "[]types.PortRange", Required: true},
	{Name: "Protocol", Flag: "protocol", Type: "types.Protocol", Required: true},
}

var fields_delete_accelerator = []leanruntime.Field{
	{Name: "AcceleratorArn", Flag: "accelerator-arn", Type: "*string", Required: true},
}

var fields_delete_cross_account_attachment = []leanruntime.Field{
	{Name: "AttachmentArn", Flag: "attachment-arn", Type: "*string", Required: true},
}

var fields_delete_custom_routing_accelerator = []leanruntime.Field{
	{Name: "AcceleratorArn", Flag: "accelerator-arn", Type: "*string", Required: true},
}

var fields_delete_custom_routing_endpoint_group = []leanruntime.Field{
	{Name: "EndpointGroupArn", Flag: "endpoint-group-arn", Type: "*string", Required: true},
}

var fields_delete_custom_routing_listener = []leanruntime.Field{
	{Name: "ListenerArn", Flag: "listener-arn", Type: "*string", Required: true},
}

var fields_delete_endpoint_group = []leanruntime.Field{
	{Name: "EndpointGroupArn", Flag: "endpoint-group-arn", Type: "*string", Required: true},
}

var fields_delete_listener = []leanruntime.Field{
	{Name: "ListenerArn", Flag: "listener-arn", Type: "*string", Required: true},
}

var fields_deny_custom_routing_traffic = []leanruntime.Field{
	{Name: "DenyAllTrafficToEndpoint", Flag: "deny-all-traffic-to-endpoint", Type: "*bool", Required: false},
	{Name: "DestinationAddresses", Flag: "destination-addresses", Type: "[]string", Required: false},
	{Name: "DestinationPorts", Flag: "destination-ports", Type: "[]int32", Required: false},
	{Name: "EndpointGroupArn", Flag: "endpoint-group-arn", Type: "*string", Required: true},
	{Name: "EndpointId", Flag: "endpoint-id", Type: "*string", Required: true},
}

var fields_deprovision_byoip_cidr = []leanruntime.Field{
	{Name: "Cidr", Flag: "cidr", Type: "*string", Required: true},
}

var fields_describe_accelerator = []leanruntime.Field{
	{Name: "AcceleratorArn", Flag: "accelerator-arn", Type: "*string", Required: true},
}

var fields_describe_accelerator_attributes = []leanruntime.Field{
	{Name: "AcceleratorArn", Flag: "accelerator-arn", Type: "*string", Required: true},
}

var fields_describe_cross_account_attachment = []leanruntime.Field{
	{Name: "AttachmentArn", Flag: "attachment-arn", Type: "*string", Required: true},
}

var fields_describe_custom_routing_accelerator = []leanruntime.Field{
	{Name: "AcceleratorArn", Flag: "accelerator-arn", Type: "*string", Required: true},
}

var fields_describe_custom_routing_accelerator_attributes = []leanruntime.Field{
	{Name: "AcceleratorArn", Flag: "accelerator-arn", Type: "*string", Required: true},
}

var fields_describe_custom_routing_endpoint_group = []leanruntime.Field{
	{Name: "EndpointGroupArn", Flag: "endpoint-group-arn", Type: "*string", Required: true},
}

var fields_describe_custom_routing_listener = []leanruntime.Field{
	{Name: "ListenerArn", Flag: "listener-arn", Type: "*string", Required: true},
}

var fields_describe_endpoint_group = []leanruntime.Field{
	{Name: "EndpointGroupArn", Flag: "endpoint-group-arn", Type: "*string", Required: true},
}

var fields_describe_listener = []leanruntime.Field{
	{Name: "ListenerArn", Flag: "listener-arn", Type: "*string", Required: true},
}

var fields_list_accelerators = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_byoip_cidrs = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_cross_account_attachments = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_cross_account_resource_accounts = []leanruntime.Field{}

var fields_list_cross_account_resources = []leanruntime.Field{
	{Name: "AcceleratorArn", Flag: "accelerator-arn", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceOwnerAwsAccountId", Flag: "resource-owner-aws-account-id", Type: "*string", Required: true},
}

var fields_list_custom_routing_accelerators = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_custom_routing_endpoint_groups = []leanruntime.Field{
	{Name: "ListenerArn", Flag: "listener-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_custom_routing_listeners = []leanruntime.Field{
	{Name: "AcceleratorArn", Flag: "accelerator-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_custom_routing_port_mappings = []leanruntime.Field{
	{Name: "AcceleratorArn", Flag: "accelerator-arn", Type: "*string", Required: true},
	{Name: "EndpointGroupArn", Flag: "endpoint-group-arn", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_custom_routing_port_mappings_by_destination = []leanruntime.Field{
	{Name: "DestinationAddress", Flag: "destination-address", Type: "*string", Required: true},
	{Name: "EndpointId", Flag: "endpoint-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_endpoint_groups = []leanruntime.Field{
	{Name: "ListenerArn", Flag: "listener-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_listeners = []leanruntime.Field{
	{Name: "AcceleratorArn", Flag: "accelerator-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_provision_byoip_cidr = []leanruntime.Field{
	{Name: "Cidr", Flag: "cidr", Type: "*string", Required: true},
	{Name: "CidrAuthorizationContext", Flag: "cidr-authorization-context", Type: "*types.CidrAuthorizationContext", Required: true},
}

var fields_remove_custom_routing_endpoints = []leanruntime.Field{
	{Name: "EndpointGroupArn", Flag: "endpoint-group-arn", Type: "*string", Required: true},
	{Name: "EndpointIds", Flag: "endpoint-ids", Type: "[]string", Required: true},
}

var fields_remove_endpoints = []leanruntime.Field{
	{Name: "EndpointGroupArn", Flag: "endpoint-group-arn", Type: "*string", Required: true},
	{Name: "EndpointIdentifiers", Flag: "endpoint-identifiers", Type: "[]types.EndpointIdentifier", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_accelerator = []leanruntime.Field{
	{Name: "AcceleratorArn", Flag: "accelerator-arn", Type: "*string", Required: true},
	{Name: "Enabled", Flag: "enabled", Type: "*bool", Required: false},
	{Name: "IpAddressType", Flag: "ip-address-type", Type: "types.IpAddressType", Required: false},
	{Name: "IpAddresses", Flag: "ip-addresses", Type: "[]string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_accelerator_attributes = []leanruntime.Field{
	{Name: "AcceleratorArn", Flag: "accelerator-arn", Type: "*string", Required: true},
	{Name: "FlowLogsEnabled", Flag: "flow-logs-enabled", Type: "*bool", Required: false},
	{Name: "FlowLogsS3Bucket", Flag: "flow-logs-s3-bucket", Type: "*string", Required: false},
	{Name: "FlowLogsS3Prefix", Flag: "flow-logs-s3-prefix", Type: "*string", Required: false},
}

var fields_update_cross_account_attachment = []leanruntime.Field{
	{Name: "AddPrincipals", Flag: "add-principals", Type: "[]string", Required: false},
	{Name: "AddResources", Flag: "add-resources", Type: "[]types.Resource", Required: false},
	{Name: "AttachmentArn", Flag: "attachment-arn", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "RemovePrincipals", Flag: "remove-principals", Type: "[]string", Required: false},
	{Name: "RemoveResources", Flag: "remove-resources", Type: "[]types.Resource", Required: false},
}

var fields_update_custom_routing_accelerator = []leanruntime.Field{
	{Name: "AcceleratorArn", Flag: "accelerator-arn", Type: "*string", Required: true},
	{Name: "Enabled", Flag: "enabled", Type: "*bool", Required: false},
	{Name: "IpAddressType", Flag: "ip-address-type", Type: "types.IpAddressType", Required: false},
	{Name: "IpAddresses", Flag: "ip-addresses", Type: "[]string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_custom_routing_accelerator_attributes = []leanruntime.Field{
	{Name: "AcceleratorArn", Flag: "accelerator-arn", Type: "*string", Required: true},
	{Name: "FlowLogsEnabled", Flag: "flow-logs-enabled", Type: "*bool", Required: false},
	{Name: "FlowLogsS3Bucket", Flag: "flow-logs-s3-bucket", Type: "*string", Required: false},
	{Name: "FlowLogsS3Prefix", Flag: "flow-logs-s3-prefix", Type: "*string", Required: false},
}

var fields_update_custom_routing_listener = []leanruntime.Field{
	{Name: "ListenerArn", Flag: "listener-arn", Type: "*string", Required: true},
	{Name: "PortRanges", Flag: "port-ranges", Type: "[]types.PortRange", Required: true},
}

var fields_update_endpoint_group = []leanruntime.Field{
	{Name: "EndpointConfigurations", Flag: "endpoint-configurations", Type: "[]types.EndpointConfiguration", Required: false},
	{Name: "EndpointGroupArn", Flag: "endpoint-group-arn", Type: "*string", Required: true},
	{Name: "HealthCheckIntervalSeconds", Flag: "health-check-interval-seconds", Type: "*int32", Required: false},
	{Name: "HealthCheckPath", Flag: "health-check-path", Type: "*string", Required: false},
	{Name: "HealthCheckPort", Flag: "health-check-port", Type: "*int32", Required: false},
	{Name: "HealthCheckProtocol", Flag: "health-check-protocol", Type: "types.HealthCheckProtocol", Required: false},
	{Name: "PortOverrides", Flag: "port-overrides", Type: "[]types.PortOverride", Required: false},
	{Name: "ThresholdCount", Flag: "threshold-count", Type: "*int32", Required: false},
	{Name: "TrafficDialPercentage", Flag: "traffic-dial-percentage", Type: "*float32", Required: false},
}

var fields_update_listener = []leanruntime.Field{
	{Name: "ClientAffinity", Flag: "client-affinity", Type: "types.ClientAffinity", Required: false},
	{Name: "ListenerArn", Flag: "listener-arn", Type: "*string", Required: true},
	{Name: "PortRanges", Flag: "port-ranges", Type: "[]types.PortRange", Required: false},
	{Name: "Protocol", Flag: "protocol", Type: "types.Protocol", Required: false},
}

var fields_withdraw_byoip_cidr = []leanruntime.Field{
	{Name: "Cidr", Flag: "cidr", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"add-custom-routing-endpoints": {
			Name:   "add-custom-routing-endpoints",
			Fields: fields_add_custom_routing_endpoints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddCustomRoutingEndpointsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_custom_routing_endpoints, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddCustomRoutingEndpoints(ctx, input)
			},
		},
		"add-endpoints": {
			Name:   "add-endpoints",
			Fields: fields_add_endpoints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddEndpointsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_endpoints, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddEndpoints(ctx, input)
			},
		},
		"advertise-byoip-cidr": {
			Name:   "advertise-byoip-cidr",
			Fields: fields_advertise_byoip_cidr,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AdvertiseByoipCidrInput{}
				if _, err := leanruntime.ApplyInput(input, fields_advertise_byoip_cidr, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AdvertiseByoipCidr(ctx, input)
			},
		},
		"allow-custom-routing-traffic": {
			Name:   "allow-custom-routing-traffic",
			Fields: fields_allow_custom_routing_traffic,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AllowCustomRoutingTrafficInput{}
				if _, err := leanruntime.ApplyInput(input, fields_allow_custom_routing_traffic, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AllowCustomRoutingTraffic(ctx, input)
			},
		},
		"create-accelerator": {
			Name:   "create-accelerator",
			Fields: fields_create_accelerator,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAcceleratorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_accelerator, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAccelerator(ctx, input)
			},
		},
		"create-cross-account-attachment": {
			Name:   "create-cross-account-attachment",
			Fields: fields_create_cross_account_attachment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCrossAccountAttachmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_cross_account_attachment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCrossAccountAttachment(ctx, input)
			},
		},
		"create-custom-routing-accelerator": {
			Name:   "create-custom-routing-accelerator",
			Fields: fields_create_custom_routing_accelerator,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCustomRoutingAcceleratorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_custom_routing_accelerator, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCustomRoutingAccelerator(ctx, input)
			},
		},
		"create-custom-routing-endpoint-group": {
			Name:   "create-custom-routing-endpoint-group",
			Fields: fields_create_custom_routing_endpoint_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCustomRoutingEndpointGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_custom_routing_endpoint_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCustomRoutingEndpointGroup(ctx, input)
			},
		},
		"create-custom-routing-listener": {
			Name:   "create-custom-routing-listener",
			Fields: fields_create_custom_routing_listener,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCustomRoutingListenerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_custom_routing_listener, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCustomRoutingListener(ctx, input)
			},
		},
		"create-endpoint-group": {
			Name:   "create-endpoint-group",
			Fields: fields_create_endpoint_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEndpointGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_endpoint_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEndpointGroup(ctx, input)
			},
		},
		"create-listener": {
			Name:   "create-listener",
			Fields: fields_create_listener,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateListenerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_listener, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateListener(ctx, input)
			},
		},
		"delete-accelerator": {
			Name:   "delete-accelerator",
			Fields: fields_delete_accelerator,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAcceleratorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_accelerator, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAccelerator(ctx, input)
			},
		},
		"delete-cross-account-attachment": {
			Name:   "delete-cross-account-attachment",
			Fields: fields_delete_cross_account_attachment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCrossAccountAttachmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_cross_account_attachment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCrossAccountAttachment(ctx, input)
			},
		},
		"delete-custom-routing-accelerator": {
			Name:   "delete-custom-routing-accelerator",
			Fields: fields_delete_custom_routing_accelerator,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCustomRoutingAcceleratorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_custom_routing_accelerator, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCustomRoutingAccelerator(ctx, input)
			},
		},
		"delete-custom-routing-endpoint-group": {
			Name:   "delete-custom-routing-endpoint-group",
			Fields: fields_delete_custom_routing_endpoint_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCustomRoutingEndpointGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_custom_routing_endpoint_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCustomRoutingEndpointGroup(ctx, input)
			},
		},
		"delete-custom-routing-listener": {
			Name:   "delete-custom-routing-listener",
			Fields: fields_delete_custom_routing_listener,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCustomRoutingListenerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_custom_routing_listener, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCustomRoutingListener(ctx, input)
			},
		},
		"delete-endpoint-group": {
			Name:   "delete-endpoint-group",
			Fields: fields_delete_endpoint_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEndpointGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_endpoint_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEndpointGroup(ctx, input)
			},
		},
		"delete-listener": {
			Name:   "delete-listener",
			Fields: fields_delete_listener,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteListenerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_listener, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteListener(ctx, input)
			},
		},
		"deny-custom-routing-traffic": {
			Name:   "deny-custom-routing-traffic",
			Fields: fields_deny_custom_routing_traffic,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DenyCustomRoutingTrafficInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deny_custom_routing_traffic, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DenyCustomRoutingTraffic(ctx, input)
			},
		},
		"deprovision-byoip-cidr": {
			Name:   "deprovision-byoip-cidr",
			Fields: fields_deprovision_byoip_cidr,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeprovisionByoipCidrInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deprovision_byoip_cidr, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeprovisionByoipCidr(ctx, input)
			},
		},
		"describe-accelerator": {
			Name:   "describe-accelerator",
			Fields: fields_describe_accelerator,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAcceleratorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_accelerator, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAccelerator(ctx, input)
			},
		},
		"describe-accelerator-attributes": {
			Name:   "describe-accelerator-attributes",
			Fields: fields_describe_accelerator_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAcceleratorAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_accelerator_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAcceleratorAttributes(ctx, input)
			},
		},
		"describe-cross-account-attachment": {
			Name:   "describe-cross-account-attachment",
			Fields: fields_describe_cross_account_attachment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCrossAccountAttachmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_cross_account_attachment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeCrossAccountAttachment(ctx, input)
			},
		},
		"describe-custom-routing-accelerator": {
			Name:   "describe-custom-routing-accelerator",
			Fields: fields_describe_custom_routing_accelerator,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCustomRoutingAcceleratorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_custom_routing_accelerator, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeCustomRoutingAccelerator(ctx, input)
			},
		},
		"describe-custom-routing-accelerator-attributes": {
			Name:   "describe-custom-routing-accelerator-attributes",
			Fields: fields_describe_custom_routing_accelerator_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCustomRoutingAcceleratorAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_custom_routing_accelerator_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeCustomRoutingAcceleratorAttributes(ctx, input)
			},
		},
		"describe-custom-routing-endpoint-group": {
			Name:   "describe-custom-routing-endpoint-group",
			Fields: fields_describe_custom_routing_endpoint_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCustomRoutingEndpointGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_custom_routing_endpoint_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeCustomRoutingEndpointGroup(ctx, input)
			},
		},
		"describe-custom-routing-listener": {
			Name:   "describe-custom-routing-listener",
			Fields: fields_describe_custom_routing_listener,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCustomRoutingListenerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_custom_routing_listener, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeCustomRoutingListener(ctx, input)
			},
		},
		"describe-endpoint-group": {
			Name:   "describe-endpoint-group",
			Fields: fields_describe_endpoint_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEndpointGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_endpoint_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEndpointGroup(ctx, input)
			},
		},
		"describe-listener": {
			Name:   "describe-listener",
			Fields: fields_describe_listener,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeListenerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_listener, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeListener(ctx, input)
			},
		},
		"list-accelerators": {
			Name:   "list-accelerators",
			Fields: fields_list_accelerators,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAcceleratorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_accelerators, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAccelerators(ctx, input)
				}
				var results []*svc.ListAcceleratorsOutput
				p := svc.NewListAcceleratorsPaginator(client, input)
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
		"list-byoip-cidrs": {
			Name:   "list-byoip-cidrs",
			Fields: fields_list_byoip_cidrs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListByoipCidrsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_byoip_cidrs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListByoipCidrs(ctx, input)
				}
				var results []*svc.ListByoipCidrsOutput
				p := svc.NewListByoipCidrsPaginator(client, input)
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
		"list-cross-account-attachments": {
			Name:   "list-cross-account-attachments",
			Fields: fields_list_cross_account_attachments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCrossAccountAttachmentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_cross_account_attachments, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCrossAccountAttachments(ctx, input)
				}
				var results []*svc.ListCrossAccountAttachmentsOutput
				p := svc.NewListCrossAccountAttachmentsPaginator(client, input)
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
		"list-cross-account-resource-accounts": {
			Name:   "list-cross-account-resource-accounts",
			Fields: fields_list_cross_account_resource_accounts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCrossAccountResourceAccountsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_cross_account_resource_accounts, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListCrossAccountResourceAccounts(ctx, input)
			},
		},
		"list-cross-account-resources": {
			Name:   "list-cross-account-resources",
			Fields: fields_list_cross_account_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCrossAccountResourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_cross_account_resources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCrossAccountResources(ctx, input)
				}
				var results []*svc.ListCrossAccountResourcesOutput
				p := svc.NewListCrossAccountResourcesPaginator(client, input)
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
		"list-custom-routing-accelerators": {
			Name:   "list-custom-routing-accelerators",
			Fields: fields_list_custom_routing_accelerators,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCustomRoutingAcceleratorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_custom_routing_accelerators, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCustomRoutingAccelerators(ctx, input)
				}
				var results []*svc.ListCustomRoutingAcceleratorsOutput
				p := svc.NewListCustomRoutingAcceleratorsPaginator(client, input)
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
		"list-custom-routing-endpoint-groups": {
			Name:   "list-custom-routing-endpoint-groups",
			Fields: fields_list_custom_routing_endpoint_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCustomRoutingEndpointGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_custom_routing_endpoint_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCustomRoutingEndpointGroups(ctx, input)
				}
				var results []*svc.ListCustomRoutingEndpointGroupsOutput
				p := svc.NewListCustomRoutingEndpointGroupsPaginator(client, input)
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
		"list-custom-routing-listeners": {
			Name:   "list-custom-routing-listeners",
			Fields: fields_list_custom_routing_listeners,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCustomRoutingListenersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_custom_routing_listeners, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCustomRoutingListeners(ctx, input)
				}
				var results []*svc.ListCustomRoutingListenersOutput
				p := svc.NewListCustomRoutingListenersPaginator(client, input)
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
		"list-custom-routing-port-mappings": {
			Name:   "list-custom-routing-port-mappings",
			Fields: fields_list_custom_routing_port_mappings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCustomRoutingPortMappingsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_custom_routing_port_mappings, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCustomRoutingPortMappings(ctx, input)
				}
				var results []*svc.ListCustomRoutingPortMappingsOutput
				p := svc.NewListCustomRoutingPortMappingsPaginator(client, input)
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
		"list-custom-routing-port-mappings-by-destination": {
			Name:   "list-custom-routing-port-mappings-by-destination",
			Fields: fields_list_custom_routing_port_mappings_by_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCustomRoutingPortMappingsByDestinationInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_custom_routing_port_mappings_by_destination, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCustomRoutingPortMappingsByDestination(ctx, input)
				}
				var results []*svc.ListCustomRoutingPortMappingsByDestinationOutput
				p := svc.NewListCustomRoutingPortMappingsByDestinationPaginator(client, input)
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
		"list-endpoint-groups": {
			Name:   "list-endpoint-groups",
			Fields: fields_list_endpoint_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEndpointGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_endpoint_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEndpointGroups(ctx, input)
				}
				var results []*svc.ListEndpointGroupsOutput
				p := svc.NewListEndpointGroupsPaginator(client, input)
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
		"list-listeners": {
			Name:   "list-listeners",
			Fields: fields_list_listeners,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListListenersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_listeners, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListListeners(ctx, input)
				}
				var results []*svc.ListListenersOutput
				p := svc.NewListListenersPaginator(client, input)
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
		"provision-byoip-cidr": {
			Name:   "provision-byoip-cidr",
			Fields: fields_provision_byoip_cidr,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ProvisionByoipCidrInput{}
				if _, err := leanruntime.ApplyInput(input, fields_provision_byoip_cidr, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ProvisionByoipCidr(ctx, input)
			},
		},
		"remove-custom-routing-endpoints": {
			Name:   "remove-custom-routing-endpoints",
			Fields: fields_remove_custom_routing_endpoints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveCustomRoutingEndpointsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_custom_routing_endpoints, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveCustomRoutingEndpoints(ctx, input)
			},
		},
		"remove-endpoints": {
			Name:   "remove-endpoints",
			Fields: fields_remove_endpoints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveEndpointsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_endpoints, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveEndpoints(ctx, input)
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
		"update-accelerator": {
			Name:   "update-accelerator",
			Fields: fields_update_accelerator,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAcceleratorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_accelerator, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAccelerator(ctx, input)
			},
		},
		"update-accelerator-attributes": {
			Name:   "update-accelerator-attributes",
			Fields: fields_update_accelerator_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAcceleratorAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_accelerator_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAcceleratorAttributes(ctx, input)
			},
		},
		"update-cross-account-attachment": {
			Name:   "update-cross-account-attachment",
			Fields: fields_update_cross_account_attachment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCrossAccountAttachmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_cross_account_attachment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCrossAccountAttachment(ctx, input)
			},
		},
		"update-custom-routing-accelerator": {
			Name:   "update-custom-routing-accelerator",
			Fields: fields_update_custom_routing_accelerator,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCustomRoutingAcceleratorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_custom_routing_accelerator, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCustomRoutingAccelerator(ctx, input)
			},
		},
		"update-custom-routing-accelerator-attributes": {
			Name:   "update-custom-routing-accelerator-attributes",
			Fields: fields_update_custom_routing_accelerator_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCustomRoutingAcceleratorAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_custom_routing_accelerator_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCustomRoutingAcceleratorAttributes(ctx, input)
			},
		},
		"update-custom-routing-listener": {
			Name:   "update-custom-routing-listener",
			Fields: fields_update_custom_routing_listener,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCustomRoutingListenerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_custom_routing_listener, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCustomRoutingListener(ctx, input)
			},
		},
		"update-endpoint-group": {
			Name:   "update-endpoint-group",
			Fields: fields_update_endpoint_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEndpointGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_endpoint_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEndpointGroup(ctx, input)
			},
		},
		"update-listener": {
			Name:   "update-listener",
			Fields: fields_update_listener,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateListenerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_listener, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateListener(ctx, input)
			},
		},
		"withdraw-byoip-cidr": {
			Name:   "withdraw-byoip-cidr",
			Fields: fields_withdraw_byoip_cidr,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.WithdrawByoipCidrInput{}
				if _, err := leanruntime.ApplyInput(input, fields_withdraw_byoip_cidr, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.WithdrawByoipCidr(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("globalaccelerator", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
