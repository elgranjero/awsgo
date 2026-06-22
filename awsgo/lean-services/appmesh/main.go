package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/appmesh"
)

var fields_create_gateway_route = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "GatewayRouteName", Flag: "gateway-route-name", Type: "*string", Required: true},
	{Name: "MeshName", Flag: "mesh-name", Type: "*string", Required: true},
	{Name: "MeshOwner", Flag: "mesh-owner", Type: "*string", Required: false},
	{Name: "Spec", Flag: "spec", Type: "*types.GatewayRouteSpec", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.TagRef", Required: false},
	{Name: "VirtualGatewayName", Flag: "virtual-gateway-name", Type: "*string", Required: true},
}

var fields_create_mesh = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "MeshName", Flag: "mesh-name", Type: "*string", Required: true},
	{Name: "Spec", Flag: "spec", Type: "*types.MeshSpec", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.TagRef", Required: false},
}

var fields_create_route = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "MeshName", Flag: "mesh-name", Type: "*string", Required: true},
	{Name: "MeshOwner", Flag: "mesh-owner", Type: "*string", Required: false},
	{Name: "RouteName", Flag: "route-name", Type: "*string", Required: true},
	{Name: "Spec", Flag: "spec", Type: "*types.RouteSpec", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.TagRef", Required: false},
	{Name: "VirtualRouterName", Flag: "virtual-router-name", Type: "*string", Required: true},
}

var fields_create_virtual_gateway = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "MeshName", Flag: "mesh-name", Type: "*string", Required: true},
	{Name: "MeshOwner", Flag: "mesh-owner", Type: "*string", Required: false},
	{Name: "Spec", Flag: "spec", Type: "*types.VirtualGatewaySpec", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.TagRef", Required: false},
	{Name: "VirtualGatewayName", Flag: "virtual-gateway-name", Type: "*string", Required: true},
}

var fields_create_virtual_node = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "MeshName", Flag: "mesh-name", Type: "*string", Required: true},
	{Name: "MeshOwner", Flag: "mesh-owner", Type: "*string", Required: false},
	{Name: "Spec", Flag: "spec", Type: "*types.VirtualNodeSpec", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.TagRef", Required: false},
	{Name: "VirtualNodeName", Flag: "virtual-node-name", Type: "*string", Required: true},
}

var fields_create_virtual_router = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "MeshName", Flag: "mesh-name", Type: "*string", Required: true},
	{Name: "MeshOwner", Flag: "mesh-owner", Type: "*string", Required: false},
	{Name: "Spec", Flag: "spec", Type: "*types.VirtualRouterSpec", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.TagRef", Required: false},
	{Name: "VirtualRouterName", Flag: "virtual-router-name", Type: "*string", Required: true},
}

var fields_create_virtual_service = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "MeshName", Flag: "mesh-name", Type: "*string", Required: true},
	{Name: "MeshOwner", Flag: "mesh-owner", Type: "*string", Required: false},
	{Name: "Spec", Flag: "spec", Type: "*types.VirtualServiceSpec", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.TagRef", Required: false},
	{Name: "VirtualServiceName", Flag: "virtual-service-name", Type: "*string", Required: true},
}

var fields_delete_gateway_route = []leanruntime.Field{
	{Name: "GatewayRouteName", Flag: "gateway-route-name", Type: "*string", Required: true},
	{Name: "MeshName", Flag: "mesh-name", Type: "*string", Required: true},
	{Name: "MeshOwner", Flag: "mesh-owner", Type: "*string", Required: false},
	{Name: "VirtualGatewayName", Flag: "virtual-gateway-name", Type: "*string", Required: true},
}

var fields_delete_mesh = []leanruntime.Field{
	{Name: "MeshName", Flag: "mesh-name", Type: "*string", Required: true},
}

var fields_delete_route = []leanruntime.Field{
	{Name: "MeshName", Flag: "mesh-name", Type: "*string", Required: true},
	{Name: "MeshOwner", Flag: "mesh-owner", Type: "*string", Required: false},
	{Name: "RouteName", Flag: "route-name", Type: "*string", Required: true},
	{Name: "VirtualRouterName", Flag: "virtual-router-name", Type: "*string", Required: true},
}

var fields_delete_virtual_gateway = []leanruntime.Field{
	{Name: "MeshName", Flag: "mesh-name", Type: "*string", Required: true},
	{Name: "MeshOwner", Flag: "mesh-owner", Type: "*string", Required: false},
	{Name: "VirtualGatewayName", Flag: "virtual-gateway-name", Type: "*string", Required: true},
}

var fields_delete_virtual_node = []leanruntime.Field{
	{Name: "MeshName", Flag: "mesh-name", Type: "*string", Required: true},
	{Name: "MeshOwner", Flag: "mesh-owner", Type: "*string", Required: false},
	{Name: "VirtualNodeName", Flag: "virtual-node-name", Type: "*string", Required: true},
}

var fields_delete_virtual_router = []leanruntime.Field{
	{Name: "MeshName", Flag: "mesh-name", Type: "*string", Required: true},
	{Name: "MeshOwner", Flag: "mesh-owner", Type: "*string", Required: false},
	{Name: "VirtualRouterName", Flag: "virtual-router-name", Type: "*string", Required: true},
}

var fields_delete_virtual_service = []leanruntime.Field{
	{Name: "MeshName", Flag: "mesh-name", Type: "*string", Required: true},
	{Name: "MeshOwner", Flag: "mesh-owner", Type: "*string", Required: false},
	{Name: "VirtualServiceName", Flag: "virtual-service-name", Type: "*string", Required: true},
}

var fields_describe_gateway_route = []leanruntime.Field{
	{Name: "GatewayRouteName", Flag: "gateway-route-name", Type: "*string", Required: true},
	{Name: "MeshName", Flag: "mesh-name", Type: "*string", Required: true},
	{Name: "MeshOwner", Flag: "mesh-owner", Type: "*string", Required: false},
	{Name: "VirtualGatewayName", Flag: "virtual-gateway-name", Type: "*string", Required: true},
}

var fields_describe_mesh = []leanruntime.Field{
	{Name: "MeshName", Flag: "mesh-name", Type: "*string", Required: true},
	{Name: "MeshOwner", Flag: "mesh-owner", Type: "*string", Required: false},
}

var fields_describe_route = []leanruntime.Field{
	{Name: "MeshName", Flag: "mesh-name", Type: "*string", Required: true},
	{Name: "MeshOwner", Flag: "mesh-owner", Type: "*string", Required: false},
	{Name: "RouteName", Flag: "route-name", Type: "*string", Required: true},
	{Name: "VirtualRouterName", Flag: "virtual-router-name", Type: "*string", Required: true},
}

var fields_describe_virtual_gateway = []leanruntime.Field{
	{Name: "MeshName", Flag: "mesh-name", Type: "*string", Required: true},
	{Name: "MeshOwner", Flag: "mesh-owner", Type: "*string", Required: false},
	{Name: "VirtualGatewayName", Flag: "virtual-gateway-name", Type: "*string", Required: true},
}

var fields_describe_virtual_node = []leanruntime.Field{
	{Name: "MeshName", Flag: "mesh-name", Type: "*string", Required: true},
	{Name: "MeshOwner", Flag: "mesh-owner", Type: "*string", Required: false},
	{Name: "VirtualNodeName", Flag: "virtual-node-name", Type: "*string", Required: true},
}

var fields_describe_virtual_router = []leanruntime.Field{
	{Name: "MeshName", Flag: "mesh-name", Type: "*string", Required: true},
	{Name: "MeshOwner", Flag: "mesh-owner", Type: "*string", Required: false},
	{Name: "VirtualRouterName", Flag: "virtual-router-name", Type: "*string", Required: true},
}

var fields_describe_virtual_service = []leanruntime.Field{
	{Name: "MeshName", Flag: "mesh-name", Type: "*string", Required: true},
	{Name: "MeshOwner", Flag: "mesh-owner", Type: "*string", Required: false},
	{Name: "VirtualServiceName", Flag: "virtual-service-name", Type: "*string", Required: true},
}

var fields_list_gateway_routes = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "MeshName", Flag: "mesh-name", Type: "*string", Required: true},
	{Name: "MeshOwner", Flag: "mesh-owner", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "VirtualGatewayName", Flag: "virtual-gateway-name", Type: "*string", Required: true},
}

var fields_list_meshes = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_routes = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "MeshName", Flag: "mesh-name", Type: "*string", Required: true},
	{Name: "MeshOwner", Flag: "mesh-owner", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "VirtualRouterName", Flag: "virtual-router-name", Type: "*string", Required: true},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_virtual_gateways = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "MeshName", Flag: "mesh-name", Type: "*string", Required: true},
	{Name: "MeshOwner", Flag: "mesh-owner", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_virtual_nodes = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "MeshName", Flag: "mesh-name", Type: "*string", Required: true},
	{Name: "MeshOwner", Flag: "mesh-owner", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_virtual_routers = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "MeshName", Flag: "mesh-name", Type: "*string", Required: true},
	{Name: "MeshOwner", Flag: "mesh-owner", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_virtual_services = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "MeshName", Flag: "mesh-name", Type: "*string", Required: true},
	{Name: "MeshOwner", Flag: "mesh-owner", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.TagRef", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_gateway_route = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "GatewayRouteName", Flag: "gateway-route-name", Type: "*string", Required: true},
	{Name: "MeshName", Flag: "mesh-name", Type: "*string", Required: true},
	{Name: "MeshOwner", Flag: "mesh-owner", Type: "*string", Required: false},
	{Name: "Spec", Flag: "spec", Type: "*types.GatewayRouteSpec", Required: true},
	{Name: "VirtualGatewayName", Flag: "virtual-gateway-name", Type: "*string", Required: true},
}

var fields_update_mesh = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "MeshName", Flag: "mesh-name", Type: "*string", Required: true},
	{Name: "Spec", Flag: "spec", Type: "*types.MeshSpec", Required: false},
}

var fields_update_route = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "MeshName", Flag: "mesh-name", Type: "*string", Required: true},
	{Name: "MeshOwner", Flag: "mesh-owner", Type: "*string", Required: false},
	{Name: "RouteName", Flag: "route-name", Type: "*string", Required: true},
	{Name: "Spec", Flag: "spec", Type: "*types.RouteSpec", Required: true},
	{Name: "VirtualRouterName", Flag: "virtual-router-name", Type: "*string", Required: true},
}

var fields_update_virtual_gateway = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "MeshName", Flag: "mesh-name", Type: "*string", Required: true},
	{Name: "MeshOwner", Flag: "mesh-owner", Type: "*string", Required: false},
	{Name: "Spec", Flag: "spec", Type: "*types.VirtualGatewaySpec", Required: true},
	{Name: "VirtualGatewayName", Flag: "virtual-gateway-name", Type: "*string", Required: true},
}

var fields_update_virtual_node = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "MeshName", Flag: "mesh-name", Type: "*string", Required: true},
	{Name: "MeshOwner", Flag: "mesh-owner", Type: "*string", Required: false},
	{Name: "Spec", Flag: "spec", Type: "*types.VirtualNodeSpec", Required: true},
	{Name: "VirtualNodeName", Flag: "virtual-node-name", Type: "*string", Required: true},
}

var fields_update_virtual_router = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "MeshName", Flag: "mesh-name", Type: "*string", Required: true},
	{Name: "MeshOwner", Flag: "mesh-owner", Type: "*string", Required: false},
	{Name: "Spec", Flag: "spec", Type: "*types.VirtualRouterSpec", Required: true},
	{Name: "VirtualRouterName", Flag: "virtual-router-name", Type: "*string", Required: true},
}

var fields_update_virtual_service = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "MeshName", Flag: "mesh-name", Type: "*string", Required: true},
	{Name: "MeshOwner", Flag: "mesh-owner", Type: "*string", Required: false},
	{Name: "Spec", Flag: "spec", Type: "*types.VirtualServiceSpec", Required: true},
	{Name: "VirtualServiceName", Flag: "virtual-service-name", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-gateway-route": {
			Name:   "create-gateway-route",
			Fields: fields_create_gateway_route,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateGatewayRouteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_gateway_route, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateGatewayRoute(ctx, input)
			},
		},
		"create-mesh": {
			Name:   "create-mesh",
			Fields: fields_create_mesh,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMeshInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_mesh, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMesh(ctx, input)
			},
		},
		"create-route": {
			Name:   "create-route",
			Fields: fields_create_route,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRouteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_route, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRoute(ctx, input)
			},
		},
		"create-virtual-gateway": {
			Name:   "create-virtual-gateway",
			Fields: fields_create_virtual_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVirtualGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_virtual_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVirtualGateway(ctx, input)
			},
		},
		"create-virtual-node": {
			Name:   "create-virtual-node",
			Fields: fields_create_virtual_node,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVirtualNodeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_virtual_node, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVirtualNode(ctx, input)
			},
		},
		"create-virtual-router": {
			Name:   "create-virtual-router",
			Fields: fields_create_virtual_router,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVirtualRouterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_virtual_router, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVirtualRouter(ctx, input)
			},
		},
		"create-virtual-service": {
			Name:   "create-virtual-service",
			Fields: fields_create_virtual_service,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVirtualServiceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_virtual_service, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVirtualService(ctx, input)
			},
		},
		"delete-gateway-route": {
			Name:   "delete-gateway-route",
			Fields: fields_delete_gateway_route,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteGatewayRouteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_gateway_route, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteGatewayRoute(ctx, input)
			},
		},
		"delete-mesh": {
			Name:   "delete-mesh",
			Fields: fields_delete_mesh,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMeshInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_mesh, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMesh(ctx, input)
			},
		},
		"delete-route": {
			Name:   "delete-route",
			Fields: fields_delete_route,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRouteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_route, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRoute(ctx, input)
			},
		},
		"delete-virtual-gateway": {
			Name:   "delete-virtual-gateway",
			Fields: fields_delete_virtual_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVirtualGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_virtual_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVirtualGateway(ctx, input)
			},
		},
		"delete-virtual-node": {
			Name:   "delete-virtual-node",
			Fields: fields_delete_virtual_node,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVirtualNodeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_virtual_node, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVirtualNode(ctx, input)
			},
		},
		"delete-virtual-router": {
			Name:   "delete-virtual-router",
			Fields: fields_delete_virtual_router,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVirtualRouterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_virtual_router, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVirtualRouter(ctx, input)
			},
		},
		"delete-virtual-service": {
			Name:   "delete-virtual-service",
			Fields: fields_delete_virtual_service,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVirtualServiceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_virtual_service, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVirtualService(ctx, input)
			},
		},
		"describe-gateway-route": {
			Name:   "describe-gateway-route",
			Fields: fields_describe_gateway_route,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeGatewayRouteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_gateway_route, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeGatewayRoute(ctx, input)
			},
		},
		"describe-mesh": {
			Name:   "describe-mesh",
			Fields: fields_describe_mesh,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMeshInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_mesh, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeMesh(ctx, input)
			},
		},
		"describe-route": {
			Name:   "describe-route",
			Fields: fields_describe_route,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRouteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_route, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeRoute(ctx, input)
			},
		},
		"describe-virtual-gateway": {
			Name:   "describe-virtual-gateway",
			Fields: fields_describe_virtual_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVirtualGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_virtual_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeVirtualGateway(ctx, input)
			},
		},
		"describe-virtual-node": {
			Name:   "describe-virtual-node",
			Fields: fields_describe_virtual_node,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVirtualNodeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_virtual_node, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeVirtualNode(ctx, input)
			},
		},
		"describe-virtual-router": {
			Name:   "describe-virtual-router",
			Fields: fields_describe_virtual_router,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVirtualRouterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_virtual_router, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeVirtualRouter(ctx, input)
			},
		},
		"describe-virtual-service": {
			Name:   "describe-virtual-service",
			Fields: fields_describe_virtual_service,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVirtualServiceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_virtual_service, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeVirtualService(ctx, input)
			},
		},
		"list-gateway-routes": {
			Name:   "list-gateway-routes",
			Fields: fields_list_gateway_routes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGatewayRoutesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_gateway_routes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListGatewayRoutes(ctx, input)
				}
				var results []*svc.ListGatewayRoutesOutput
				p := svc.NewListGatewayRoutesPaginator(client, input)
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
		"list-meshes": {
			Name:   "list-meshes",
			Fields: fields_list_meshes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMeshesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_meshes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMeshes(ctx, input)
				}
				var results []*svc.ListMeshesOutput
				p := svc.NewListMeshesPaginator(client, input)
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
		"list-routes": {
			Name:   "list-routes",
			Fields: fields_list_routes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRoutesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_routes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRoutes(ctx, input)
				}
				var results []*svc.ListRoutesOutput
				p := svc.NewListRoutesPaginator(client, input)
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
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_tags_for_resource, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTagsForResource(ctx, input)
				}
				var results []*svc.ListTagsForResourceOutput
				p := svc.NewListTagsForResourcePaginator(client, input)
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
		"list-virtual-gateways": {
			Name:   "list-virtual-gateways",
			Fields: fields_list_virtual_gateways,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVirtualGatewaysInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_virtual_gateways, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListVirtualGateways(ctx, input)
				}
				var results []*svc.ListVirtualGatewaysOutput
				p := svc.NewListVirtualGatewaysPaginator(client, input)
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
		"list-virtual-nodes": {
			Name:   "list-virtual-nodes",
			Fields: fields_list_virtual_nodes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVirtualNodesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_virtual_nodes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListVirtualNodes(ctx, input)
				}
				var results []*svc.ListVirtualNodesOutput
				p := svc.NewListVirtualNodesPaginator(client, input)
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
		"list-virtual-routers": {
			Name:   "list-virtual-routers",
			Fields: fields_list_virtual_routers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVirtualRoutersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_virtual_routers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListVirtualRouters(ctx, input)
				}
				var results []*svc.ListVirtualRoutersOutput
				p := svc.NewListVirtualRoutersPaginator(client, input)
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
		"list-virtual-services": {
			Name:   "list-virtual-services",
			Fields: fields_list_virtual_services,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVirtualServicesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_virtual_services, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListVirtualServices(ctx, input)
				}
				var results []*svc.ListVirtualServicesOutput
				p := svc.NewListVirtualServicesPaginator(client, input)
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
		"update-gateway-route": {
			Name:   "update-gateway-route",
			Fields: fields_update_gateway_route,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateGatewayRouteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_gateway_route, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateGatewayRoute(ctx, input)
			},
		},
		"update-mesh": {
			Name:   "update-mesh",
			Fields: fields_update_mesh,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMeshInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_mesh, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMesh(ctx, input)
			},
		},
		"update-route": {
			Name:   "update-route",
			Fields: fields_update_route,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRouteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_route, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRoute(ctx, input)
			},
		},
		"update-virtual-gateway": {
			Name:   "update-virtual-gateway",
			Fields: fields_update_virtual_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateVirtualGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_virtual_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateVirtualGateway(ctx, input)
			},
		},
		"update-virtual-node": {
			Name:   "update-virtual-node",
			Fields: fields_update_virtual_node,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateVirtualNodeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_virtual_node, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateVirtualNode(ctx, input)
			},
		},
		"update-virtual-router": {
			Name:   "update-virtual-router",
			Fields: fields_update_virtual_router,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateVirtualRouterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_virtual_router, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateVirtualRouter(ctx, input)
			},
		},
		"update-virtual-service": {
			Name:   "update-virtual-service",
			Fields: fields_update_virtual_service,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateVirtualServiceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_virtual_service, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateVirtualService(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("appmesh", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
