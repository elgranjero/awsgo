package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/ssmguiconnect"
)

var fields_delete_connection_recording_preferences = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
}

var fields_get_connection_recording_preferences = []leanruntime.Field{}

var fields_update_connection_recording_preferences = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ConnectionRecordingPreferences", Flag: "connection-recording-preferences", Type: "*types.ConnectionRecordingPreferences", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"delete-connection-recording-preferences": {
			Name:   "delete-connection-recording-preferences",
			Fields: fields_delete_connection_recording_preferences,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConnectionRecordingPreferencesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_connection_recording_preferences, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConnectionRecordingPreferences(ctx, input)
			},
		},
		"get-connection-recording-preferences": {
			Name:   "get-connection-recording-preferences",
			Fields: fields_get_connection_recording_preferences,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConnectionRecordingPreferencesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_connection_recording_preferences, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConnectionRecordingPreferences(ctx, input)
			},
		},
		"update-connection-recording-preferences": {
			Name:   "update-connection-recording-preferences",
			Fields: fields_update_connection_recording_preferences,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateConnectionRecordingPreferencesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_connection_recording_preferences, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateConnectionRecordingPreferences(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("ssmguiconnect", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
