package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/cloudhsm"
)

var fields_add_tags_to_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagList", Flag: "tag-list", Type: "[]types.Tag", Required: true},
}

var fields_create_hapg = []leanruntime.Field{
	{Name: "Label", Flag: "label", Type: "*string", Required: true},
}

var fields_create_hsm = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "EniIp", Flag: "eni-ip", Type: "*string", Required: false},
	{Name: "ExternalId", Flag: "external-id", Type: "*string", Required: false},
	{Name: "IamRoleArn", Flag: "iam-role-arn", Type: "*string", Required: true},
	{Name: "SshKey", Flag: "ssh-key", Type: "*string", Required: true},
	{Name: "SubnetId", Flag: "subnet-id", Type: "*string", Required: true},
	{Name: "SubscriptionType", Flag: "subscription-type", Type: "types.SubscriptionType", Required: true},
	{Name: "SyslogIp", Flag: "syslog-ip", Type: "*string", Required: false},
}

var fields_create_luna_client = []leanruntime.Field{
	{Name: "Certificate", Flag: "certificate", Type: "*string", Required: true},
	{Name: "Label", Flag: "label", Type: "*string", Required: false},
}

var fields_delete_hapg = []leanruntime.Field{
	{Name: "HapgArn", Flag: "hapg-arn", Type: "*string", Required: true},
}

var fields_delete_hsm = []leanruntime.Field{
	{Name: "HsmArn", Flag: "hsm-arn", Type: "*string", Required: true},
}

var fields_delete_luna_client = []leanruntime.Field{
	{Name: "ClientArn", Flag: "client-arn", Type: "*string", Required: true},
}

var fields_describe_hapg = []leanruntime.Field{
	{Name: "HapgArn", Flag: "hapg-arn", Type: "*string", Required: true},
}

var fields_describe_hsm = []leanruntime.Field{
	{Name: "HsmArn", Flag: "hsm-arn", Type: "*string", Required: false},
	{Name: "HsmSerialNumber", Flag: "hsm-serial-number", Type: "*string", Required: false},
}

var fields_describe_luna_client = []leanruntime.Field{
	{Name: "CertificateFingerprint", Flag: "certificate-fingerprint", Type: "*string", Required: false},
	{Name: "ClientArn", Flag: "client-arn", Type: "*string", Required: false},
}

var fields_get_config = []leanruntime.Field{
	{Name: "ClientArn", Flag: "client-arn", Type: "*string", Required: true},
	{Name: "ClientVersion", Flag: "client-version", Type: "types.ClientVersion", Required: true},
	{Name: "HapgList", Flag: "hapg-list", Type: "[]string", Required: true},
}

var fields_list_available_zones = []leanruntime.Field{}

var fields_list_hapgs = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_hsms = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_luna_clients = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_modify_hapg = []leanruntime.Field{
	{Name: "HapgArn", Flag: "hapg-arn", Type: "*string", Required: true},
	{Name: "Label", Flag: "label", Type: "*string", Required: false},
	{Name: "PartitionSerialList", Flag: "partition-serial-list", Type: "[]string", Required: false},
}

var fields_modify_hsm = []leanruntime.Field{
	{Name: "EniIp", Flag: "eni-ip", Type: "*string", Required: false},
	{Name: "ExternalId", Flag: "external-id", Type: "*string", Required: false},
	{Name: "HsmArn", Flag: "hsm-arn", Type: "*string", Required: true},
	{Name: "IamRoleArn", Flag: "iam-role-arn", Type: "*string", Required: false},
	{Name: "SubnetId", Flag: "subnet-id", Type: "*string", Required: false},
	{Name: "SyslogIp", Flag: "syslog-ip", Type: "*string", Required: false},
}

var fields_modify_luna_client = []leanruntime.Field{
	{Name: "Certificate", Flag: "certificate", Type: "*string", Required: true},
	{Name: "ClientArn", Flag: "client-arn", Type: "*string", Required: true},
}

var fields_remove_tags_from_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeyList", Flag: "tag-key-list", Type: "[]string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"add-tags-to-resource": {
			Name:   "add-tags-to-resource",
			Fields: fields_add_tags_to_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddTagsToResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_tags_to_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddTagsToResource(ctx, input)
			},
		},
		"create-hapg": {
			Name:   "create-hapg",
			Fields: fields_create_hapg,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateHapgInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_hapg, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateHapg(ctx, input)
			},
		},
		"create-hsm": {
			Name:   "create-hsm",
			Fields: fields_create_hsm,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateHsmInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_hsm, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateHsm(ctx, input)
			},
		},
		"create-luna-client": {
			Name:   "create-luna-client",
			Fields: fields_create_luna_client,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLunaClientInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_luna_client, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLunaClient(ctx, input)
			},
		},
		"delete-hapg": {
			Name:   "delete-hapg",
			Fields: fields_delete_hapg,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteHapgInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_hapg, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteHapg(ctx, input)
			},
		},
		"delete-hsm": {
			Name:   "delete-hsm",
			Fields: fields_delete_hsm,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteHsmInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_hsm, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteHsm(ctx, input)
			},
		},
		"delete-luna-client": {
			Name:   "delete-luna-client",
			Fields: fields_delete_luna_client,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLunaClientInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_luna_client, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLunaClient(ctx, input)
			},
		},
		"describe-hapg": {
			Name:   "describe-hapg",
			Fields: fields_describe_hapg,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeHapgInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_hapg, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeHapg(ctx, input)
			},
		},
		"describe-hsm": {
			Name:   "describe-hsm",
			Fields: fields_describe_hsm,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeHsmInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_hsm, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeHsm(ctx, input)
			},
		},
		"describe-luna-client": {
			Name:   "describe-luna-client",
			Fields: fields_describe_luna_client,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLunaClientInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_luna_client, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeLunaClient(ctx, input)
			},
		},
		"get-config": {
			Name:   "get-config",
			Fields: fields_get_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConfig(ctx, input)
			},
		},
		"list-available-zones": {
			Name:   "list-available-zones",
			Fields: fields_list_available_zones,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAvailableZonesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_available_zones, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListAvailableZones(ctx, input)
			},
		},
		"list-hapgs": {
			Name:   "list-hapgs",
			Fields: fields_list_hapgs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListHapgsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_hapgs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListHapgs(ctx, input)
			},
		},
		"list-hsms": {
			Name:   "list-hsms",
			Fields: fields_list_hsms,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListHsmsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_hsms, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListHsms(ctx, input)
			},
		},
		"list-luna-clients": {
			Name:   "list-luna-clients",
			Fields: fields_list_luna_clients,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLunaClientsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_luna_clients, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListLunaClients(ctx, input)
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
		"modify-hapg": {
			Name:   "modify-hapg",
			Fields: fields_modify_hapg,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyHapgInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_hapg, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyHapg(ctx, input)
			},
		},
		"modify-hsm": {
			Name:   "modify-hsm",
			Fields: fields_modify_hsm,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyHsmInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_hsm, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyHsm(ctx, input)
			},
		},
		"modify-luna-client": {
			Name:   "modify-luna-client",
			Fields: fields_modify_luna_client,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyLunaClientInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_luna_client, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyLunaClient(ctx, input)
			},
		},
		"remove-tags-from-resource": {
			Name:   "remove-tags-from-resource",
			Fields: fields_remove_tags_from_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveTagsFromResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_tags_from_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveTagsFromResource(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("cloudhsm", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
