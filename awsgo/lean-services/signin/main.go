package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/signin"
)

var fields_create_oauth2_token = []leanruntime.Field{
	{Name: "TokenInput", Flag: "token-input", Type: "*types.CreateOAuth2TokenRequestBody", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-oauth2-token": {
			Name:   "create-oauth2-token",
			Fields: fields_create_oauth2_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateOAuth2TokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_oauth2_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateOAuth2Token(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("signin", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
