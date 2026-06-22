package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/workmailmessageflow"
)

var fields_get_raw_message_content = []leanruntime.Field{
	{Name: "MessageId", Flag: "message-id", Type: "*string", Required: true},
}

var fields_put_raw_message_content = []leanruntime.Field{
	{Name: "Content", Flag: "content", Type: "*types.RawMessageContent", Required: true},
	{Name: "MessageId", Flag: "message-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"get-raw-message-content": {
			Name:   "get-raw-message-content",
			Fields: fields_get_raw_message_content,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRawMessageContentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_raw_message_content, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRawMessageContent(ctx, input)
			},
		},
		"put-raw-message-content": {
			Name:   "put-raw-message-content",
			Fields: fields_put_raw_message_content,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutRawMessageContentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_raw_message_content, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutRawMessageContent(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("workmailmessageflow", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
