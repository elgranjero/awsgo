package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/appconfigdata"
)

var fields_get_latest_configuration = []leanruntime.Field{
	{Name: "ConfigurationToken", Flag: "configuration-token", Type: "*string", Required: true},
}

var fields_start_configuration_session = []leanruntime.Field{
	{Name: "ApplicationIdentifier", Flag: "application-identifier", Type: "*string", Required: true},
	{Name: "ConfigurationProfileIdentifier", Flag: "configuration-profile-identifier", Type: "*string", Required: true},
	{Name: "EnvironmentIdentifier", Flag: "environment-identifier", Type: "*string", Required: true},
	{Name: "RequiredMinimumPollIntervalInSeconds", Flag: "required-minimum-poll-interval-in-seconds", Type: "*int32", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"get-latest-configuration": {
			Name:   "get-latest-configuration",
			Fields: fields_get_latest_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLatestConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_latest_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLatestConfiguration(ctx, input)
			},
		},
		"start-configuration-session": {
			Name:   "start-configuration-session",
			Fields: fields_start_configuration_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartConfigurationSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_configuration_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartConfigurationSession(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("appconfigdata", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
