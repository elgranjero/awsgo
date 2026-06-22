package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
)

var fields_delete_connection = []leanruntime.Field{
	{Name: "ConnectionId", Flag: "connection-id", Type: "*string", Required: true},
}

var fields_get_connection = []leanruntime.Field{
	{Name: "ConnectionId", Flag: "connection-id", Type: "*string", Required: true},
}

var fields_post_to_connection = []leanruntime.Field{
	{Name: "ConnectionId", Flag: "connection-id", Type: "*string", Required: true},
	{Name: "Data", Flag: "data", Type: "[]byte", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
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
		"get-connection": {
			Name:   "get-connection",
			Fields: fields_get_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConnection(ctx, input)
			},
		},
		"post-to-connection": {
			Name:   "post-to-connection",
			Fields: fields_post_to_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PostToConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_post_to_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PostToConnection(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("apigatewaymanagementapi", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
