package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/kinesisvideomedia"
)

var fields_get_media = []leanruntime.Field{
	{Name: "StartSelector", Flag: "start-selector", Type: "*types.StartSelector", Required: true},
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamName", Flag: "stream-name", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"get-media": {
			Name:   "get-media",
			Fields: fields_get_media,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMediaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_media, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMedia(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("kinesisvideomedia", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
