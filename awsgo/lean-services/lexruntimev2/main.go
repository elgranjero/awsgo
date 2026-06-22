package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/lexruntimev2"
)

var fields_delete_session = []leanruntime.Field{
	{Name: "BotAliasId", Flag: "bot-alias-id", Type: "*string", Required: true},
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: true},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
}

var fields_get_session = []leanruntime.Field{
	{Name: "BotAliasId", Flag: "bot-alias-id", Type: "*string", Required: true},
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: true},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
}

var fields_put_session = []leanruntime.Field{
	{Name: "BotAliasId", Flag: "bot-alias-id", Type: "*string", Required: true},
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: true},
	{Name: "Messages", Flag: "messages", Type: "[]types.Message", Required: false},
	{Name: "RequestAttributes", Flag: "request-attributes", Type: "map[string]string", Required: false},
	{Name: "ResponseContentType", Flag: "response-content-type", Type: "*string", Required: false},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
	{Name: "SessionState", Flag: "session-state", Type: "*types.SessionState", Required: true},
}

var fields_recognize_text = []leanruntime.Field{
	{Name: "BotAliasId", Flag: "bot-alias-id", Type: "*string", Required: true},
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: true},
	{Name: "RequestAttributes", Flag: "request-attributes", Type: "map[string]string", Required: false},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
	{Name: "SessionState", Flag: "session-state", Type: "*types.SessionState", Required: false},
	{Name: "Text", Flag: "text", Type: "*string", Required: true},
}

var fields_recognize_utterance = []leanruntime.Field{
	{Name: "BotAliasId", Flag: "bot-alias-id", Type: "*string", Required: true},
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "InputStream", Flag: "input-stream", Type: "io.Reader", Required: false},
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: true},
	{Name: "RequestAttributes", Flag: "request-attributes", Type: "*string", Required: false},
	{Name: "RequestContentType", Flag: "request-content-type", Type: "*string", Required: true},
	{Name: "ResponseContentType", Flag: "response-content-type", Type: "*string", Required: false},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
	{Name: "SessionState", Flag: "session-state", Type: "*string", Required: false},
}

var fields_start_conversation = []leanruntime.Field{
	{Name: "BotAliasId", Flag: "bot-alias-id", Type: "*string", Required: true},
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "ConversationMode", Flag: "conversation-mode", Type: "types.ConversationMode", Required: false},
	{Name: "LocaleId", Flag: "locale-id", Type: "*string", Required: true},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
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
		"recognize-text": {
			Name:   "recognize-text",
			Fields: fields_recognize_text,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RecognizeTextInput{}
				if _, err := leanruntime.ApplyInput(input, fields_recognize_text, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RecognizeText(ctx, input)
			},
		},
		"recognize-utterance": {
			Name:   "recognize-utterance",
			Fields: fields_recognize_utterance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RecognizeUtteranceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_recognize_utterance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RecognizeUtterance(ctx, input)
			},
		},
		"start-conversation": {
			Name:   "start-conversation",
			Fields: fields_start_conversation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartConversationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_conversation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartConversation(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("lexruntimev2", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
