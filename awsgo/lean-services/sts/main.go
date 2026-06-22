package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/sts"
)

var fields_assume_role = []leanruntime.Field{
	{Name: "DurationSeconds", Flag: "duration-seconds", Type: "*int32", Required: false},
	{Name: "ExternalId", Flag: "external-id", Type: "*string", Required: false},
	{Name: "Policy", Flag: "policy", Type: "*string", Required: false},
	{Name: "PolicyArns", Flag: "policy-arns", Type: "[]types.PolicyDescriptorType", Required: false},
	{Name: "ProvidedContexts", Flag: "provided-contexts", Type: "[]types.ProvidedContext", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "RoleSessionName", Flag: "role-session-name", Type: "*string", Required: true},
	{Name: "SerialNumber", Flag: "serial-number", Type: "*string", Required: false},
	{Name: "SourceIdentity", Flag: "source-identity", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TokenCode", Flag: "token-code", Type: "*string", Required: false},
	{Name: "TransitiveTagKeys", Flag: "transitive-tag-keys", Type: "[]string", Required: false},
}

var fields_assume_role_with_saml = []leanruntime.Field{
	{Name: "DurationSeconds", Flag: "duration-seconds", Type: "*int32", Required: false},
	{Name: "Policy", Flag: "policy", Type: "*string", Required: false},
	{Name: "PolicyArns", Flag: "policy-arns", Type: "[]types.PolicyDescriptorType", Required: false},
	{Name: "PrincipalArn", Flag: "principal-arn", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "SAMLAssertion", Flag: "saml-assertion", Type: "*string", Required: true},
}

var fields_assume_role_with_web_identity = []leanruntime.Field{
	{Name: "DurationSeconds", Flag: "duration-seconds", Type: "*int32", Required: false},
	{Name: "Policy", Flag: "policy", Type: "*string", Required: false},
	{Name: "PolicyArns", Flag: "policy-arns", Type: "[]types.PolicyDescriptorType", Required: false},
	{Name: "ProviderId", Flag: "provider-id", Type: "*string", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "RoleSessionName", Flag: "role-session-name", Type: "*string", Required: true},
	{Name: "WebIdentityToken", Flag: "web-identity-token", Type: "*string", Required: true},
}

var fields_assume_root = []leanruntime.Field{
	{Name: "DurationSeconds", Flag: "duration-seconds", Type: "*int32", Required: false},
	{Name: "TargetPrincipal", Flag: "target-principal", Type: "*string", Required: true},
	{Name: "TaskPolicyArn", Flag: "task-policy-arn", Type: "*types.PolicyDescriptorType", Required: true},
}

var fields_decode_authorization_message = []leanruntime.Field{
	{Name: "EncodedMessage", Flag: "encoded-message", Type: "*string", Required: true},
}

var fields_get_access_key_info = []leanruntime.Field{
	{Name: "AccessKeyId", Flag: "access-key-id", Type: "*string", Required: true},
}

var fields_get_caller_identity = []leanruntime.Field{}

var fields_get_delegated_access_token = []leanruntime.Field{
	{Name: "TradeInToken", Flag: "trade-in-token", Type: "*string", Required: true},
}

var fields_get_federation_token = []leanruntime.Field{
	{Name: "DurationSeconds", Flag: "duration-seconds", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Policy", Flag: "policy", Type: "*string", Required: false},
	{Name: "PolicyArns", Flag: "policy-arns", Type: "[]types.PolicyDescriptorType", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_get_session_token = []leanruntime.Field{
	{Name: "DurationSeconds", Flag: "duration-seconds", Type: "*int32", Required: false},
	{Name: "SerialNumber", Flag: "serial-number", Type: "*string", Required: false},
	{Name: "TokenCode", Flag: "token-code", Type: "*string", Required: false},
}

var fields_get_web_identity_token = []leanruntime.Field{
	{Name: "Audience", Flag: "audience", Type: "[]string", Required: true},
	{Name: "DurationSeconds", Flag: "duration-seconds", Type: "*int32", Required: false},
	{Name: "SigningAlgorithm", Flag: "signing-algorithm", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"assume-role": {
			Name:   "assume-role",
			Fields: fields_assume_role,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssumeRoleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_assume_role, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssumeRole(ctx, input)
			},
		},
		"assume-role-with-saml": {
			Name:   "assume-role-with-saml",
			Fields: fields_assume_role_with_saml,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssumeRoleWithSAMLInput{}
				if _, err := leanruntime.ApplyInput(input, fields_assume_role_with_saml, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssumeRoleWithSAML(ctx, input)
			},
		},
		"assume-role-with-web-identity": {
			Name:   "assume-role-with-web-identity",
			Fields: fields_assume_role_with_web_identity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssumeRoleWithWebIdentityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_assume_role_with_web_identity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssumeRoleWithWebIdentity(ctx, input)
			},
		},
		"assume-root": {
			Name:   "assume-root",
			Fields: fields_assume_root,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssumeRootInput{}
				if _, err := leanruntime.ApplyInput(input, fields_assume_root, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssumeRoot(ctx, input)
			},
		},
		"decode-authorization-message": {
			Name:   "decode-authorization-message",
			Fields: fields_decode_authorization_message,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DecodeAuthorizationMessageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_decode_authorization_message, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DecodeAuthorizationMessage(ctx, input)
			},
		},
		"get-access-key-info": {
			Name:   "get-access-key-info",
			Fields: fields_get_access_key_info,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccessKeyInfoInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_access_key_info, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccessKeyInfo(ctx, input)
			},
		},
		"get-caller-identity": {
			Name:   "get-caller-identity",
			Fields: fields_get_caller_identity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCallerIdentityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_caller_identity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCallerIdentity(ctx, input)
			},
		},
		"get-delegated-access-token": {
			Name:   "get-delegated-access-token",
			Fields: fields_get_delegated_access_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDelegatedAccessTokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_delegated_access_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDelegatedAccessToken(ctx, input)
			},
		},
		"get-federation-token": {
			Name:   "get-federation-token",
			Fields: fields_get_federation_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFederationTokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_federation_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFederationToken(ctx, input)
			},
		},
		"get-session-token": {
			Name:   "get-session-token",
			Fields: fields_get_session_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSessionTokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_session_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSessionToken(ctx, input)
			},
		},
		"get-web-identity-token": {
			Name:   "get-web-identity-token",
			Fields: fields_get_web_identity_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWebIdentityTokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_web_identity_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetWebIdentityToken(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("sts", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
