package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/rtbfabric"
)

var fields_accept_link = []leanruntime.Field{
	{Name: "Attributes", Flag: "attributes", Type: "*types.LinkAttributes", Required: false},
	{Name: "GatewayId", Flag: "gateway-id", Type: "*string", Required: true},
	{Name: "LinkId", Flag: "link-id", Type: "*string", Required: true},
	{Name: "LogSettings", Flag: "log-settings", Type: "*types.LinkLogSettings", Required: true},
}

var fields_create_inbound_external_link = []leanruntime.Field{
	{Name: "Attributes", Flag: "attributes", Type: "*types.LinkAttributes", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "GatewayId", Flag: "gateway-id", Type: "*string", Required: true},
	{Name: "LogSettings", Flag: "log-settings", Type: "*types.LinkLogSettings", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_link = []leanruntime.Field{
	{Name: "Attributes", Flag: "attributes", Type: "*types.LinkAttributes", Required: false},
	{Name: "GatewayId", Flag: "gateway-id", Type: "*string", Required: true},
	{Name: "HttpResponderAllowed", Flag: "http-responder-allowed", Type: "*bool", Required: false},
	{Name: "LogSettings", Flag: "log-settings", Type: "*types.LinkLogSettings", Required: true},
	{Name: "PeerGatewayId", Flag: "peer-gateway-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_outbound_external_link = []leanruntime.Field{
	{Name: "Attributes", Flag: "attributes", Type: "*types.LinkAttributes", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "GatewayId", Flag: "gateway-id", Type: "*string", Required: true},
	{Name: "LogSettings", Flag: "log-settings", Type: "*types.LinkLogSettings", Required: true},
	{Name: "PublicEndpoint", Flag: "public-endpoint", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_requester_gateway = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "SecurityGroupIds", Flag: "security-group-ids", Type: "[]string", Required: true},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: true},
}

var fields_create_responder_gateway = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: false},
	{Name: "ManagedEndpointConfiguration", Flag: "managed-endpoint-configuration", Type: "types.ManagedEndpointConfiguration", Required: false},
	{Name: "Port", Flag: "port", Type: "*int32", Required: true},
	{Name: "Protocol", Flag: "protocol", Type: "types.Protocol", Required: true},
	{Name: "SecurityGroupIds", Flag: "security-group-ids", Type: "[]string", Required: true},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TrustStoreConfiguration", Flag: "trust-store-configuration", Type: "*types.TrustStoreConfiguration", Required: false},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: true},
}

var fields_delete_inbound_external_link = []leanruntime.Field{
	{Name: "GatewayId", Flag: "gateway-id", Type: "*string", Required: true},
	{Name: "LinkId", Flag: "link-id", Type: "*string", Required: true},
}

var fields_delete_link = []leanruntime.Field{
	{Name: "GatewayId", Flag: "gateway-id", Type: "*string", Required: true},
	{Name: "LinkId", Flag: "link-id", Type: "*string", Required: true},
}

var fields_delete_outbound_external_link = []leanruntime.Field{
	{Name: "GatewayId", Flag: "gateway-id", Type: "*string", Required: true},
	{Name: "LinkId", Flag: "link-id", Type: "*string", Required: true},
}

var fields_delete_requester_gateway = []leanruntime.Field{
	{Name: "GatewayId", Flag: "gateway-id", Type: "*string", Required: true},
}

var fields_delete_responder_gateway = []leanruntime.Field{
	{Name: "GatewayId", Flag: "gateway-id", Type: "*string", Required: true},
}

var fields_get_inbound_external_link = []leanruntime.Field{
	{Name: "GatewayId", Flag: "gateway-id", Type: "*string", Required: true},
	{Name: "LinkId", Flag: "link-id", Type: "*string", Required: true},
}

var fields_get_link = []leanruntime.Field{
	{Name: "GatewayId", Flag: "gateway-id", Type: "*string", Required: true},
	{Name: "LinkId", Flag: "link-id", Type: "*string", Required: true},
}

var fields_get_outbound_external_link = []leanruntime.Field{
	{Name: "GatewayId", Flag: "gateway-id", Type: "*string", Required: true},
	{Name: "LinkId", Flag: "link-id", Type: "*string", Required: true},
}

var fields_get_requester_gateway = []leanruntime.Field{
	{Name: "GatewayId", Flag: "gateway-id", Type: "*string", Required: true},
}

var fields_get_responder_gateway = []leanruntime.Field{
	{Name: "GatewayId", Flag: "gateway-id", Type: "*string", Required: true},
}

var fields_list_links = []leanruntime.Field{
	{Name: "GatewayId", Flag: "gateway-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_requester_gateways = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_responder_gateways = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_reject_link = []leanruntime.Field{
	{Name: "GatewayId", Flag: "gateway-id", Type: "*string", Required: true},
	{Name: "LinkId", Flag: "link-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_link = []leanruntime.Field{
	{Name: "GatewayId", Flag: "gateway-id", Type: "*string", Required: true},
	{Name: "LinkId", Flag: "link-id", Type: "*string", Required: true},
	{Name: "LogSettings", Flag: "log-settings", Type: "*types.LinkLogSettings", Required: false},
}

var fields_update_link_module_flow = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "GatewayId", Flag: "gateway-id", Type: "*string", Required: true},
	{Name: "LinkId", Flag: "link-id", Type: "*string", Required: true},
	{Name: "Modules", Flag: "modules", Type: "[]types.ModuleConfiguration", Required: true},
}

var fields_update_requester_gateway = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "GatewayId", Flag: "gateway-id", Type: "*string", Required: true},
}

var fields_update_responder_gateway = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: false},
	{Name: "GatewayId", Flag: "gateway-id", Type: "*string", Required: true},
	{Name: "ManagedEndpointConfiguration", Flag: "managed-endpoint-configuration", Type: "types.ManagedEndpointConfiguration", Required: false},
	{Name: "Port", Flag: "port", Type: "*int32", Required: true},
	{Name: "Protocol", Flag: "protocol", Type: "types.Protocol", Required: true},
	{Name: "TrustStoreConfiguration", Flag: "trust-store-configuration", Type: "*types.TrustStoreConfiguration", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"accept-link": {
			Name:   "accept-link",
			Fields: fields_accept_link,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AcceptLinkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_accept_link, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AcceptLink(ctx, input)
			},
		},
		"create-inbound-external-link": {
			Name:   "create-inbound-external-link",
			Fields: fields_create_inbound_external_link,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateInboundExternalLinkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_inbound_external_link, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateInboundExternalLink(ctx, input)
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
		"create-outbound-external-link": {
			Name:   "create-outbound-external-link",
			Fields: fields_create_outbound_external_link,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateOutboundExternalLinkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_outbound_external_link, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateOutboundExternalLink(ctx, input)
			},
		},
		"create-requester-gateway": {
			Name:   "create-requester-gateway",
			Fields: fields_create_requester_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRequesterGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_requester_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRequesterGateway(ctx, input)
			},
		},
		"create-responder-gateway": {
			Name:   "create-responder-gateway",
			Fields: fields_create_responder_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateResponderGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_responder_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateResponderGateway(ctx, input)
			},
		},
		"delete-inbound-external-link": {
			Name:   "delete-inbound-external-link",
			Fields: fields_delete_inbound_external_link,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteInboundExternalLinkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_inbound_external_link, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteInboundExternalLink(ctx, input)
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
		"delete-outbound-external-link": {
			Name:   "delete-outbound-external-link",
			Fields: fields_delete_outbound_external_link,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteOutboundExternalLinkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_outbound_external_link, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteOutboundExternalLink(ctx, input)
			},
		},
		"delete-requester-gateway": {
			Name:   "delete-requester-gateway",
			Fields: fields_delete_requester_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRequesterGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_requester_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRequesterGateway(ctx, input)
			},
		},
		"delete-responder-gateway": {
			Name:   "delete-responder-gateway",
			Fields: fields_delete_responder_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteResponderGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_responder_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteResponderGateway(ctx, input)
			},
		},
		"get-inbound-external-link": {
			Name:   "get-inbound-external-link",
			Fields: fields_get_inbound_external_link,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetInboundExternalLinkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_inbound_external_link, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetInboundExternalLink(ctx, input)
			},
		},
		"get-link": {
			Name:   "get-link",
			Fields: fields_get_link,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLinkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_link, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLink(ctx, input)
			},
		},
		"get-outbound-external-link": {
			Name:   "get-outbound-external-link",
			Fields: fields_get_outbound_external_link,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOutboundExternalLinkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_outbound_external_link, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetOutboundExternalLink(ctx, input)
			},
		},
		"get-requester-gateway": {
			Name:   "get-requester-gateway",
			Fields: fields_get_requester_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRequesterGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_requester_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRequesterGateway(ctx, input)
			},
		},
		"get-responder-gateway": {
			Name:   "get-responder-gateway",
			Fields: fields_get_responder_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResponderGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_responder_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResponderGateway(ctx, input)
			},
		},
		"list-links": {
			Name:   "list-links",
			Fields: fields_list_links,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLinksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_links, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLinks(ctx, input)
				}
				var results []*svc.ListLinksOutput
				p := svc.NewListLinksPaginator(client, input)
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
		"list-requester-gateways": {
			Name:   "list-requester-gateways",
			Fields: fields_list_requester_gateways,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRequesterGatewaysInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_requester_gateways, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRequesterGateways(ctx, input)
				}
				var results []*svc.ListRequesterGatewaysOutput
				p := svc.NewListRequesterGatewaysPaginator(client, input)
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
		"list-responder-gateways": {
			Name:   "list-responder-gateways",
			Fields: fields_list_responder_gateways,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResponderGatewaysInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_responder_gateways, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResponderGateways(ctx, input)
				}
				var results []*svc.ListResponderGatewaysOutput
				p := svc.NewListResponderGatewaysPaginator(client, input)
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
		"reject-link": {
			Name:   "reject-link",
			Fields: fields_reject_link,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RejectLinkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reject_link, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RejectLink(ctx, input)
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
		"update-link-module-flow": {
			Name:   "update-link-module-flow",
			Fields: fields_update_link_module_flow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateLinkModuleFlowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_link_module_flow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateLinkModuleFlow(ctx, input)
			},
		},
		"update-requester-gateway": {
			Name:   "update-requester-gateway",
			Fields: fields_update_requester_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRequesterGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_requester_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRequesterGateway(ctx, input)
			},
		},
		"update-responder-gateway": {
			Name:   "update-responder-gateway",
			Fields: fields_update_responder_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateResponderGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_responder_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateResponderGateway(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("rtbfabric", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
