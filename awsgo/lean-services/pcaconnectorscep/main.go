package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/pcaconnectorscep"
)

var fields_create_challenge = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ConnectorArn", Flag: "connector-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_connector = []leanruntime.Field{
	{Name: "CertificateAuthorityArn", Flag: "certificate-authority-arn", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "MobileDeviceManagement", Flag: "mobile-device-management", Type: "types.MobileDeviceManagement", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "VpcEndpointId", Flag: "vpc-endpoint-id", Type: "*string", Required: false},
}

var fields_delete_challenge = []leanruntime.Field{
	{Name: "ChallengeArn", Flag: "challenge-arn", Type: "*string", Required: true},
}

var fields_delete_connector = []leanruntime.Field{
	{Name: "ConnectorArn", Flag: "connector-arn", Type: "*string", Required: true},
}

var fields_get_challenge_metadata = []leanruntime.Field{
	{Name: "ChallengeArn", Flag: "challenge-arn", Type: "*string", Required: true},
}

var fields_get_challenge_password = []leanruntime.Field{
	{Name: "ChallengeArn", Flag: "challenge-arn", Type: "*string", Required: true},
}

var fields_get_connector = []leanruntime.Field{
	{Name: "ConnectorArn", Flag: "connector-arn", Type: "*string", Required: true},
}

var fields_list_challenge_metadata = []leanruntime.Field{
	{Name: "ConnectorArn", Flag: "connector-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_connectors = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-challenge": {
			Name:   "create-challenge",
			Fields: fields_create_challenge,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateChallengeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_challenge, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateChallenge(ctx, input)
			},
		},
		"create-connector": {
			Name:   "create-connector",
			Fields: fields_create_connector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateConnectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_connector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateConnector(ctx, input)
			},
		},
		"delete-challenge": {
			Name:   "delete-challenge",
			Fields: fields_delete_challenge,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteChallengeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_challenge, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteChallenge(ctx, input)
			},
		},
		"delete-connector": {
			Name:   "delete-connector",
			Fields: fields_delete_connector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConnectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_connector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConnector(ctx, input)
			},
		},
		"get-challenge-metadata": {
			Name:   "get-challenge-metadata",
			Fields: fields_get_challenge_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetChallengeMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_challenge_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetChallengeMetadata(ctx, input)
			},
		},
		"get-challenge-password": {
			Name:   "get-challenge-password",
			Fields: fields_get_challenge_password,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetChallengePasswordInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_challenge_password, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetChallengePassword(ctx, input)
			},
		},
		"get-connector": {
			Name:   "get-connector",
			Fields: fields_get_connector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConnectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_connector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConnector(ctx, input)
			},
		},
		"list-challenge-metadata": {
			Name:   "list-challenge-metadata",
			Fields: fields_list_challenge_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListChallengeMetadataInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_challenge_metadata, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListChallengeMetadata(ctx, input)
				}
				var results []*svc.ListChallengeMetadataOutput
				p := svc.NewListChallengeMetadataPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-connectors": {
			Name:   "list-connectors",
			Fields: fields_list_connectors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConnectorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_connectors, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListConnectors(ctx, input)
				}
				var results []*svc.ListConnectorsOutput
				p := svc.NewListConnectorsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-tags-for-resource": {
			Name:   "list-tags-for-resource",
			Fields: fields_list_tags_for_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTagsForResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_tags_for_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTagsForResource(ctx, input)
			},
		},
		"tag-resource": {
			Name:   "tag-resource",
			Fields: fields_tag_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TagResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_tag_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TagResource(ctx, input)
			},
		},
		"untag-resource": {
			Name:   "untag-resource",
			Fields: fields_untag_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UntagResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_untag_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UntagResource(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("pcaconnectorscep", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
