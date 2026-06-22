package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/kinesisvideowebrtcstorage"
)

var fields_join_storage_session = []leanruntime.Field{
	{Name: "ChannelArn", Flag: "channel-arn", Type: "*string", Required: true},
}

var fields_join_storage_session_as_viewer = []leanruntime.Field{
	{Name: "ChannelArn", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "ClientId", Flag: "client-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"join-storage-session": {
			Name:   "join-storage-session",
			Fields: fields_join_storage_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.JoinStorageSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_join_storage_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.JoinStorageSession(ctx, input)
			},
		},
		"join-storage-session-as-viewer": {
			Name:   "join-storage-session-as-viewer",
			Fields: fields_join_storage_session_as_viewer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.JoinStorageSessionAsViewerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_join_storage_session_as_viewer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.JoinStorageSessionAsViewer(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("kinesisvideowebrtcstorage", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
