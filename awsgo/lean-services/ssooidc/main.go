package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/ssooidc"
)

var fields_create_token = []leanruntime.Field{
	{Name: "ClientId", Flag: "client-id", Type: "*string", Required: true},
	{Name: "ClientSecret", Flag: "client-secret", Type: "*string", Required: true},
	{Name: "Code", Flag: "code", Type: "*string", Required: false},
	{Name: "CodeVerifier", Flag: "code-verifier", Type: "*string", Required: false},
	{Name: "DeviceCode", Flag: "device-code", Type: "*string", Required: false},
	{Name: "GrantType", Flag: "grant-type", Type: "*string", Required: true},
	{Name: "RedirectUri", Flag: "redirect-uri", Type: "*string", Required: false},
	{Name: "RefreshToken", Flag: "refresh-token", Type: "*string", Required: false},
	{Name: "Scope", Flag: "scope", Type: "[]string", Required: false},
}

var fields_create_token_with_iam = []leanruntime.Field{
	{Name: "Assertion", Flag: "assertion", Type: "*string", Required: false},
	{Name: "ClientId", Flag: "client-id", Type: "*string", Required: true},
	{Name: "Code", Flag: "code", Type: "*string", Required: false},
	{Name: "CodeVerifier", Flag: "code-verifier", Type: "*string", Required: false},
	{Name: "GrantType", Flag: "grant-type", Type: "*string", Required: true},
	{Name: "RedirectUri", Flag: "redirect-uri", Type: "*string", Required: false},
	{Name: "RefreshToken", Flag: "refresh-token", Type: "*string", Required: false},
	{Name: "RequestedTokenType", Flag: "requested-token-type", Type: "*string", Required: false},
	{Name: "Scope", Flag: "scope", Type: "[]string", Required: false},
	{Name: "SubjectToken", Flag: "subject-token", Type: "*string", Required: false},
	{Name: "SubjectTokenType", Flag: "subject-token-type", Type: "*string", Required: false},
}

var fields_register_client = []leanruntime.Field{
	{Name: "ClientName", Flag: "client-name", Type: "*string", Required: true},
	{Name: "ClientType", Flag: "client-type", Type: "*string", Required: true},
	{Name: "EntitledApplicationArn", Flag: "entitled-application-arn", Type: "*string", Required: false},
	{Name: "GrantTypes", Flag: "grant-types", Type: "[]string", Required: false},
	{Name: "IssuerUrl", Flag: "issuer-url", Type: "*string", Required: false},
	{Name: "RedirectUris", Flag: "redirect-uris", Type: "[]string", Required: false},
	{Name: "Scopes", Flag: "scopes", Type: "[]string", Required: false},
}

var fields_start_device_authorization = []leanruntime.Field{
	{Name: "ClientId", Flag: "client-id", Type: "*string", Required: true},
	{Name: "ClientSecret", Flag: "client-secret", Type: "*string", Required: true},
	{Name: "StartUrl", Flag: "start-url", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-token": {
			Name:   "create-token",
			Fields: fields_create_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateToken(ctx, input)
			},
		},
		"create-token-with-iam": {
			Name:   "create-token-with-iam",
			Fields: fields_create_token_with_iam,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTokenWithIAMInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_token_with_iam, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTokenWithIAM(ctx, input)
			},
		},
		"register-client": {
			Name:   "register-client",
			Fields: fields_register_client,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterClientInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_client, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterClient(ctx, input)
			},
		},
		"start-device-authorization": {
			Name:   "start-device-authorization",
			Fields: fields_start_device_authorization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartDeviceAuthorizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_device_authorization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartDeviceAuthorization(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("ssooidc", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
