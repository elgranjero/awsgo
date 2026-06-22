package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/personalizeevents"
)

var fields_put_action_interactions = []leanruntime.Field{
	{Name: "ActionInteractions", Flag: "action-interactions", Type: "[]types.ActionInteraction", Required: true},
	{Name: "TrackingId", Flag: "tracking-id", Type: "*string", Required: true},
}

var fields_put_actions = []leanruntime.Field{
	{Name: "Actions", Flag: "actions", Type: "[]types.Action", Required: true},
	{Name: "DatasetArn", Flag: "dataset-arn", Type: "*string", Required: true},
}

var fields_put_events = []leanruntime.Field{
	{Name: "EventList", Flag: "event-list", Type: "[]types.Event", Required: true},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
	{Name: "TrackingId", Flag: "tracking-id", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: false},
}

var fields_put_items = []leanruntime.Field{
	{Name: "DatasetArn", Flag: "dataset-arn", Type: "*string", Required: true},
	{Name: "Items", Flag: "items", Type: "[]types.Item", Required: true},
}

var fields_put_users = []leanruntime.Field{
	{Name: "DatasetArn", Flag: "dataset-arn", Type: "*string", Required: true},
	{Name: "Users", Flag: "users", Type: "[]types.User", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"put-action-interactions": {
			Name:   "put-action-interactions",
			Fields: fields_put_action_interactions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutActionInteractionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_action_interactions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutActionInteractions(ctx, input)
			},
		},
		"put-actions": {
			Name:   "put-actions",
			Fields: fields_put_actions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutActionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_actions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutActions(ctx, input)
			},
		},
		"put-events": {
			Name:   "put-events",
			Fields: fields_put_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutEventsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_events, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutEvents(ctx, input)
			},
		},
		"put-items": {
			Name:   "put-items",
			Fields: fields_put_items,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutItemsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_items, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutItems(ctx, input)
			},
		},
		"put-users": {
			Name:   "put-users",
			Fields: fields_put_users,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutUsersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_users, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutUsers(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("personalizeevents", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
