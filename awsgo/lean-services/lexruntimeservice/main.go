package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/lexruntimeservice"
)

var fields_delete_session = []leanruntime.Field{
	{Name: "BotAlias", Flag: "bot-alias", Type: "*string", Required: true},
	{Name: "BotName", Flag: "bot-name", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_get_session = []leanruntime.Field{
	{Name: "BotAlias", Flag: "bot-alias", Type: "*string", Required: true},
	{Name: "BotName", Flag: "bot-name", Type: "*string", Required: true},
	{Name: "CheckpointLabelFilter", Flag: "checkpoint-label-filter", Type: "*string", Required: false},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_post_content = []leanruntime.Field{
	{Name: "Accept", Flag: "accept", Type: "*string", Required: false},
	{Name: "ActiveContexts", Flag: "active-contexts", Type: "*string", Required: false},
	{Name: "BotAlias", Flag: "bot-alias", Type: "*string", Required: true},
	{Name: "BotName", Flag: "bot-name", Type: "*string", Required: true},
	{Name: "ContentType", Flag: "content-type", Type: "*string", Required: true},
	{Name: "InputStream", Flag: "input-stream", Type: "io.Reader", Required: true},
	{Name: "RequestAttributes", Flag: "request-attributes", Type: "*string", Required: false},
	{Name: "SessionAttributes", Flag: "session-attributes", Type: "*string", Required: false},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_post_text = []leanruntime.Field{
	{Name: "ActiveContexts", Flag: "active-contexts", Type: "[]types.ActiveContext", Required: false},
	{Name: "BotAlias", Flag: "bot-alias", Type: "*string", Required: true},
	{Name: "BotName", Flag: "bot-name", Type: "*string", Required: true},
	{Name: "InputText", Flag: "input-text", Type: "*string", Required: true},
	{Name: "RequestAttributes", Flag: "request-attributes", Type: "map[string]string", Required: false},
	{Name: "SessionAttributes", Flag: "session-attributes", Type: "map[string]string", Required: false},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_put_session = []leanruntime.Field{
	{Name: "Accept", Flag: "accept", Type: "*string", Required: false},
	{Name: "ActiveContexts", Flag: "active-contexts", Type: "[]types.ActiveContext", Required: false},
	{Name: "BotAlias", Flag: "bot-alias", Type: "*string", Required: true},
	{Name: "BotName", Flag: "bot-name", Type: "*string", Required: true},
	{Name: "DialogAction", Flag: "dialog-action", Type: "*types.DialogAction", Required: false},
	{Name: "RecentIntentSummaryView", Flag: "recent-intent-summary-view", Type: "[]types.IntentSummary", Required: false},
	{Name: "SessionAttributes", Flag: "session-attributes", Type: "map[string]string", Required: false},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"delete-session": {
			Name:   "delete-session",
			Fields: fields_delete_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSession(ctx, input)
			},
		},
		"get-session": {
			Name:   "get-session",
			Fields: fields_get_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSession(ctx, input)
			},
		},
		"post-content": {
			Name:   "post-content",
			Fields: fields_post_content,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PostContentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_post_content, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PostContent(ctx, input)
			},
		},
		"post-text": {
			Name:   "post-text",
			Fields: fields_post_text,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PostTextInput{}
				if _, err := leanruntime.ApplyInput(input, fields_post_text, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PostText(ctx, input)
			},
		},
		"put-session": {
			Name:   "put-session",
			Fields: fields_put_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutSession(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("lexruntimeservice", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
