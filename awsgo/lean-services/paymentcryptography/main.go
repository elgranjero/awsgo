package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/paymentcryptography"
)

var fields_add_key_replication_regions = []leanruntime.Field{
	{Name: "KeyIdentifier", Flag: "key-identifier", Type: "*string", Required: true},
	{Name: "ReplicationRegions", Flag: "replication-regions", Type: "[]string", Required: true},
}

var fields_create_alias = []leanruntime.Field{
	{Name: "AliasName", Flag: "alias-name", Type: "*string", Required: true},
	{Name: "KeyArn", Flag: "key-arn", Type: "*string", Required: false},
}

var fields_create_key = []leanruntime.Field{
	{Name: "DeriveKeyUsage", Flag: "derive-key-usage", Type: "types.DeriveKeyUsage", Required: false},
	{Name: "Enabled", Flag: "enabled", Type: "*bool", Required: false},
	{Name: "Exportable", Flag: "exportable", Type: "*bool", Required: true},
	{Name: "KeyAttributes", Flag: "key-attributes", Type: "*types.KeyAttributes", Required: true},
	{Name: "KeyCheckValueAlgorithm", Flag: "key-check-value-algorithm", Type: "types.KeyCheckValueAlgorithm", Required: false},
	{Name: "ReplicationRegions", Flag: "replication-regions", Type: "[]string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_delete_alias = []leanruntime.Field{
	{Name: "AliasName", Flag: "alias-name", Type: "*string", Required: true},
}

var fields_delete_key = []leanruntime.Field{
	{Name: "DeleteKeyInDays", Flag: "delete-key-in-days", Type: "*int32", Required: false},
	{Name: "KeyIdentifier", Flag: "key-identifier", Type: "*string", Required: true},
}

var fields_disable_default_key_replication_regions = []leanruntime.Field{
	{Name: "ReplicationRegions", Flag: "replication-regions", Type: "[]string", Required: true},
}

var fields_enable_default_key_replication_regions = []leanruntime.Field{
	{Name: "ReplicationRegions", Flag: "replication-regions", Type: "[]string", Required: true},
}

var fields_export_key = []leanruntime.Field{
	{Name: "ExportAttributes", Flag: "export-attributes", Type: "*types.ExportAttributes", Required: false},
	{Name: "ExportKeyIdentifier", Flag: "export-key-identifier", Type: "*string", Required: true},
	{Name: "KeyMaterial", Flag: "key-material", Type: "types.ExportKeyMaterial", Required: true},
}

var fields_get_alias = []leanruntime.Field{
	{Name: "AliasName", Flag: "alias-name", Type: "*string", Required: true},
}

var fields_get_certificate_signing_request = []leanruntime.Field{
	{Name: "CertificateSubject", Flag: "certificate-subject", Type: "*types.CertificateSubjectType", Required: true},
	{Name: "KeyIdentifier", Flag: "key-identifier", Type: "*string", Required: true},
	{Name: "SigningAlgorithm", Flag: "signing-algorithm", Type: "types.SigningAlgorithmType", Required: true},
}

var fields_get_default_key_replication_regions = []leanruntime.Field{}

var fields_get_key = []leanruntime.Field{
	{Name: "KeyIdentifier", Flag: "key-identifier", Type: "*string", Required: true},
}

var fields_get_parameters_for_export = []leanruntime.Field{
	{Name: "KeyMaterialType", Flag: "key-material-type", Type: "types.KeyMaterialType", Required: true},
	{Name: "SigningKeyAlgorithm", Flag: "signing-key-algorithm", Type: "types.KeyAlgorithm", Required: true},
}

var fields_get_parameters_for_import = []leanruntime.Field{
	{Name: "KeyMaterialType", Flag: "key-material-type", Type: "types.KeyMaterialType", Required: true},
	{Name: "WrappingKeyAlgorithm", Flag: "wrapping-key-algorithm", Type: "types.KeyAlgorithm", Required: true},
}

var fields_get_public_key_certificate = []leanruntime.Field{
	{Name: "KeyIdentifier", Flag: "key-identifier", Type: "*string", Required: true},
}

var fields_import_key = []leanruntime.Field{
	{Name: "Enabled", Flag: "enabled", Type: "*bool", Required: false},
	{Name: "KeyCheckValueAlgorithm", Flag: "key-check-value-algorithm", Type: "types.KeyCheckValueAlgorithm", Required: false},
	{Name: "KeyMaterial", Flag: "key-material", Type: "types.ImportKeyMaterial", Required: true},
	{Name: "ReplicationRegions", Flag: "replication-regions", Type: "[]string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_list_aliases = []leanruntime.Field{
	{Name: "KeyArn", Flag: "key-arn", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_keys = []leanruntime.Field{
	{Name: "KeyState", Flag: "key-state", Type: "types.KeyState", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_remove_key_replication_regions = []leanruntime.Field{
	{Name: "KeyIdentifier", Flag: "key-identifier", Type: "*string", Required: true},
	{Name: "ReplicationRegions", Flag: "replication-regions", Type: "[]string", Required: true},
}

var fields_restore_key = []leanruntime.Field{
	{Name: "KeyIdentifier", Flag: "key-identifier", Type: "*string", Required: true},
}

var fields_start_key_usage = []leanruntime.Field{
	{Name: "KeyIdentifier", Flag: "key-identifier", Type: "*string", Required: true},
}

var fields_stop_key_usage = []leanruntime.Field{
	{Name: "KeyIdentifier", Flag: "key-identifier", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_alias = []leanruntime.Field{
	{Name: "AliasName", Flag: "alias-name", Type: "*string", Required: true},
	{Name: "KeyArn", Flag: "key-arn", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"add-key-replication-regions": {
			Name:   "add-key-replication-regions",
			Fields: fields_add_key_replication_regions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddKeyReplicationRegionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_key_replication_regions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddKeyReplicationRegions(ctx, input)
			},
		},
		"create-alias": {
			Name:   "create-alias",
			Fields: fields_create_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAlias(ctx, input)
			},
		},
		"create-key": {
			Name:   "create-key",
			Fields: fields_create_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateKey(ctx, input)
			},
		},
		"delete-alias": {
			Name:   "delete-alias",
			Fields: fields_delete_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAlias(ctx, input)
			},
		},
		"delete-key": {
			Name:   "delete-key",
			Fields: fields_delete_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteKey(ctx, input)
			},
		},
		"disable-default-key-replication-regions": {
			Name:   "disable-default-key-replication-regions",
			Fields: fields_disable_default_key_replication_regions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableDefaultKeyReplicationRegionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_default_key_replication_regions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableDefaultKeyReplicationRegions(ctx, input)
			},
		},
		"enable-default-key-replication-regions": {
			Name:   "enable-default-key-replication-regions",
			Fields: fields_enable_default_key_replication_regions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableDefaultKeyReplicationRegionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_default_key_replication_regions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableDefaultKeyReplicationRegions(ctx, input)
			},
		},
		"export-key": {
			Name:   "export-key",
			Fields: fields_export_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExportKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_export_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExportKey(ctx, input)
			},
		},
		"get-alias": {
			Name:   "get-alias",
			Fields: fields_get_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAlias(ctx, input)
			},
		},
		"get-certificate-signing-request": {
			Name:   "get-certificate-signing-request",
			Fields: fields_get_certificate_signing_request,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCertificateSigningRequestInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_certificate_signing_request, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCertificateSigningRequest(ctx, input)
			},
		},
		"get-default-key-replication-regions": {
			Name:   "get-default-key-replication-regions",
			Fields: fields_get_default_key_replication_regions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDefaultKeyReplicationRegionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_default_key_replication_regions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDefaultKeyReplicationRegions(ctx, input)
			},
		},
		"get-key": {
			Name:   "get-key",
			Fields: fields_get_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetKey(ctx, input)
			},
		},
		"get-parameters-for-export": {
			Name:   "get-parameters-for-export",
			Fields: fields_get_parameters_for_export,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetParametersForExportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_parameters_for_export, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetParametersForExport(ctx, input)
			},
		},
		"get-parameters-for-import": {
			Name:   "get-parameters-for-import",
			Fields: fields_get_parameters_for_import,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetParametersForImportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_parameters_for_import, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetParametersForImport(ctx, input)
			},
		},
		"get-public-key-certificate": {
			Name:   "get-public-key-certificate",
			Fields: fields_get_public_key_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPublicKeyCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_public_key_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPublicKeyCertificate(ctx, input)
			},
		},
		"import-key": {
			Name:   "import-key",
			Fields: fields_import_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportKey(ctx, input)
			},
		},
		"list-aliases": {
			Name:   "list-aliases",
			Fields: fields_list_aliases,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAliasesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_aliases, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAliases(ctx, input)
				}
				var results []*svc.ListAliasesOutput
				p := svc.NewListAliasesPaginator(client, input)
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
		"list-keys": {
			Name:   "list-keys",
			Fields: fields_list_keys,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListKeysInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_keys, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListKeys(ctx, input)
				}
				var results []*svc.ListKeysOutput
				p := svc.NewListKeysPaginator(client, input)
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
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_tags_for_resource, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTagsForResource(ctx, input)
				}
				var results []*svc.ListTagsForResourceOutput
				p := svc.NewListTagsForResourcePaginator(client, input)
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
		"remove-key-replication-regions": {
			Name:   "remove-key-replication-regions",
			Fields: fields_remove_key_replication_regions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveKeyReplicationRegionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_key_replication_regions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveKeyReplicationRegions(ctx, input)
			},
		},
		"restore-key": {
			Name:   "restore-key",
			Fields: fields_restore_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RestoreKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_restore_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RestoreKey(ctx, input)
			},
		},
		"start-key-usage": {
			Name:   "start-key-usage",
			Fields: fields_start_key_usage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartKeyUsageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_key_usage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartKeyUsage(ctx, input)
			},
		},
		"stop-key-usage": {
			Name:   "stop-key-usage",
			Fields: fields_stop_key_usage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopKeyUsageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_key_usage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopKeyUsage(ctx, input)
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
		"update-alias": {
			Name:   "update-alias",
			Fields: fields_update_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAlias(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("paymentcryptography", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
