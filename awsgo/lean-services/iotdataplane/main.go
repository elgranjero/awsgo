package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/iotdataplane"
)

var fields_delete_connection = []leanruntime.Field{
	{Name: "CleanSession", Flag: "clean-session", Type: "bool", Required: false},
	{Name: "ClientId", Flag: "client-id", Type: "*string", Required: true},
	{Name: "PreventWillMessage", Flag: "prevent-will-message", Type: "bool", Required: false},
}

var fields_delete_thing_shadow = []leanruntime.Field{
	{Name: "ShadowName", Flag: "shadow-name", Type: "*string", Required: false},
	{Name: "ThingName", Flag: "thing-name", Type: "*string", Required: true},
}

var fields_get_retained_message = []leanruntime.Field{
	{Name: "Topic", Flag: "topic", Type: "*string", Required: true},
}

var fields_get_thing_shadow = []leanruntime.Field{
	{Name: "ShadowName", Flag: "shadow-name", Type: "*string", Required: false},
	{Name: "ThingName", Flag: "thing-name", Type: "*string", Required: true},
}

var fields_list_named_shadows_for_thing = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
	{Name: "ThingName", Flag: "thing-name", Type: "*string", Required: true},
}

var fields_list_retained_messages = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_publish = []leanruntime.Field{
	{Name: "ContentType", Flag: "content-type", Type: "*string", Required: false},
	{Name: "CorrelationData", Flag: "correlation-data", Type: "*string", Required: false},
	{Name: "MessageExpiry", Flag: "message-expiry", Type: "int64", Required: false},
	{Name: "Payload", Flag: "payload", Type: "[]byte", Required: false},
	{Name: "PayloadFormatIndicator", Flag: "payload-format-indicator", Type: "types.PayloadFormatIndicator", Required: false},
	{Name: "Qos", Flag: "qos", Type: "int32", Required: false},
	{Name: "ResponseTopic", Flag: "response-topic", Type: "*string", Required: false},
	{Name: "Retain", Flag: "retain", Type: "bool", Required: false},
	{Name: "Topic", Flag: "topic", Type: "*string", Required: true},
	{Name: "UserProperties", Flag: "user-properties", Type: "*string", Required: false},
}

var fields_update_thing_shadow = []leanruntime.Field{
	{Name: "Payload", Flag: "payload", Type: "[]byte", Required: true},
	{Name: "ShadowName", Flag: "shadow-name", Type: "*string", Required: false},
	{Name: "ThingName", Flag: "thing-name", Type: "*string", Required: true},
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
		"delete-thing-shadow": {
			Name:   "delete-thing-shadow",
			Fields: fields_delete_thing_shadow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteThingShadowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_thing_shadow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteThingShadow(ctx, input)
			},
		},
		"get-retained-message": {
			Name:   "get-retained-message",
			Fields: fields_get_retained_message,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRetainedMessageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_retained_message, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRetainedMessage(ctx, input)
			},
		},
		"get-thing-shadow": {
			Name:   "get-thing-shadow",
			Fields: fields_get_thing_shadow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetThingShadowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_thing_shadow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetThingShadow(ctx, input)
			},
		},
		"list-named-shadows-for-thing": {
			Name:   "list-named-shadows-for-thing",
			Fields: fields_list_named_shadows_for_thing,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListNamedShadowsForThingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_named_shadows_for_thing, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListNamedShadowsForThing(ctx, input)
			},
		},
		"list-retained-messages": {
			Name:   "list-retained-messages",
			Fields: fields_list_retained_messages,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRetainedMessagesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_retained_messages, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRetainedMessages(ctx, input)
				}
				var results []*svc.ListRetainedMessagesOutput
				p := svc.NewListRetainedMessagesPaginator(client, input)
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
		"publish": {
			Name:   "publish",
			Fields: fields_publish,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PublishInput{}
				if _, err := leanruntime.ApplyInput(input, fields_publish, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.Publish(ctx, input)
			},
		},
		"update-thing-shadow": {
			Name:   "update-thing-shadow",
			Fields: fields_update_thing_shadow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateThingShadowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_thing_shadow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateThingShadow(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("iotdataplane", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
