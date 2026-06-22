package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
)

var fields_cancel_resource_request = []leanruntime.Field{
	{Name: "RequestToken", Flag: "request-token", Type: "*string", Required: true},
}

var fields_create_resource = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DesiredState", Flag: "desired-state", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "TypeName", Flag: "type-name", Type: "*string", Required: true},
	{Name: "TypeVersionId", Flag: "type-version-id", Type: "*string", Required: false},
}

var fields_delete_resource = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "TypeName", Flag: "type-name", Type: "*string", Required: true},
	{Name: "TypeVersionId", Flag: "type-version-id", Type: "*string", Required: false},
}

var fields_get_resource = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "TypeName", Flag: "type-name", Type: "*string", Required: true},
	{Name: "TypeVersionId", Flag: "type-version-id", Type: "*string", Required: false},
}

var fields_get_resource_request_status = []leanruntime.Field{
	{Name: "RequestToken", Flag: "request-token", Type: "*string", Required: true},
}

var fields_list_resource_requests = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceRequestStatusFilter", Flag: "resource-request-status-filter", Type: "*types.ResourceRequestStatusFilter", Required: false},
}

var fields_list_resources = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceModel", Flag: "resource-model", Type: "*string", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "TypeName", Flag: "type-name", Type: "*string", Required: true},
	{Name: "TypeVersionId", Flag: "type-version-id", Type: "*string", Required: false},
}

var fields_update_resource = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "PatchDocument", Flag: "patch-document", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "TypeName", Flag: "type-name", Type: "*string", Required: true},
	{Name: "TypeVersionId", Flag: "type-version-id", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"cancel-resource-request": {
			Name:   "cancel-resource-request",
			Fields: fields_cancel_resource_request,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelResourceRequestInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_resource_request, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelResourceRequest(ctx, input)
			},
		},
		"create-resource": {
			Name:   "create-resource",
			Fields: fields_create_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateResource(ctx, input)
			},
		},
		"delete-resource": {
			Name:   "delete-resource",
			Fields: fields_delete_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteResource(ctx, input)
			},
		},
		"get-resource": {
			Name:   "get-resource",
			Fields: fields_get_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResource(ctx, input)
			},
		},
		"get-resource-request-status": {
			Name:   "get-resource-request-status",
			Fields: fields_get_resource_request_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourceRequestStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_resource_request_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResourceRequestStatus(ctx, input)
			},
		},
		"list-resource-requests": {
			Name:   "list-resource-requests",
			Fields: fields_list_resource_requests,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResourceRequestsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_resource_requests, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResourceRequests(ctx, input)
				}
				var results []*svc.ListResourceRequestsOutput
				p := svc.NewListResourceRequestsPaginator(client, input)
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
		"list-resources": {
			Name:   "list-resources",
			Fields: fields_list_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_resources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResources(ctx, input)
				}
				var results []*svc.ListResourcesOutput
				p := svc.NewListResourcesPaginator(client, input)
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
		"update-resource": {
			Name:   "update-resource",
			Fields: fields_update_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateResource(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("cloudcontrol", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
