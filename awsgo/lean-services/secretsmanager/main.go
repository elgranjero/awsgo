package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

var fields_batch_get_secret_value = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SecretIdList", Flag: "secret-id-list", Type: "[]string", Required: false},
}

var fields_cancel_rotate_secret = []leanruntime.Field{
	{Name: "SecretId", Flag: "secret-id", Type: "*string", Required: true},
}

var fields_create_secret = []leanruntime.Field{
	{Name: "AddReplicaRegions", Flag: "add-replica-regions", Type: "[]types.ReplicaRegionType", Required: false},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ForceOverwriteReplicaSecret", Flag: "force-overwrite-replica-secret", Type: "bool", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "SecretBinary", Flag: "secret-binary", Type: "[]byte", Required: false},
	{Name: "SecretString", Flag: "secret-string", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Type", Flag: "type", Type: "*string", Required: false},
}

var fields_delete_resource_policy = []leanruntime.Field{
	{Name: "SecretId", Flag: "secret-id", Type: "*string", Required: true},
}

var fields_delete_secret = []leanruntime.Field{
	{Name: "ForceDeleteWithoutRecovery", Flag: "force-delete-without-recovery", Type: "*bool", Required: false},
	{Name: "RecoveryWindowInDays", Flag: "recovery-window-in-days", Type: "*int64", Required: false},
	{Name: "SecretId", Flag: "secret-id", Type: "*string", Required: true},
}

var fields_describe_secret = []leanruntime.Field{
	{Name: "SecretId", Flag: "secret-id", Type: "*string", Required: true},
}

var fields_get_random_password = []leanruntime.Field{
	{Name: "ExcludeCharacters", Flag: "exclude-characters", Type: "*string", Required: false},
	{Name: "ExcludeLowercase", Flag: "exclude-lowercase", Type: "*bool", Required: false},
	{Name: "ExcludeNumbers", Flag: "exclude-numbers", Type: "*bool", Required: false},
	{Name: "ExcludePunctuation", Flag: "exclude-punctuation", Type: "*bool", Required: false},
	{Name: "ExcludeUppercase", Flag: "exclude-uppercase", Type: "*bool", Required: false},
	{Name: "IncludeSpace", Flag: "include-space", Type: "*bool", Required: false},
	{Name: "PasswordLength", Flag: "password-length", Type: "*int64", Required: false},
	{Name: "RequireEachIncludedType", Flag: "require-each-included-type", Type: "*bool", Required: false},
}

var fields_get_resource_policy = []leanruntime.Field{
	{Name: "SecretId", Flag: "secret-id", Type: "*string", Required: true},
}

var fields_get_secret_value = []leanruntime.Field{
	{Name: "SecretId", Flag: "secret-id", Type: "*string", Required: true},
	{Name: "VersionId", Flag: "version-id", Type: "*string", Required: false},
	{Name: "VersionStage", Flag: "version-stage", Type: "*string", Required: false},
}

var fields_list_secret_version_ids = []leanruntime.Field{
	{Name: "IncludeDeprecated", Flag: "include-deprecated", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SecretId", Flag: "secret-id", Type: "*string", Required: true},
}

var fields_list_secrets = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "IncludePlannedDeletion", Flag: "include-planned-deletion", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.SortByType", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrderType", Required: false},
}

var fields_put_resource_policy = []leanruntime.Field{
	{Name: "BlockPublicPolicy", Flag: "block-public-policy", Type: "*bool", Required: false},
	{Name: "ResourcePolicy", Flag: "resource-policy", Type: "*string", Required: true},
	{Name: "SecretId", Flag: "secret-id", Type: "*string", Required: true},
}

var fields_put_secret_value = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "RotationToken", Flag: "rotation-token", Type: "*string", Required: false},
	{Name: "SecretBinary", Flag: "secret-binary", Type: "[]byte", Required: false},
	{Name: "SecretId", Flag: "secret-id", Type: "*string", Required: true},
	{Name: "SecretString", Flag: "secret-string", Type: "*string", Required: false},
	{Name: "VersionStages", Flag: "version-stages", Type: "[]string", Required: false},
}

var fields_remove_regions_from_replication = []leanruntime.Field{
	{Name: "RemoveReplicaRegions", Flag: "remove-replica-regions", Type: "[]string", Required: true},
	{Name: "SecretId", Flag: "secret-id", Type: "*string", Required: true},
}

var fields_replicate_secret_to_regions = []leanruntime.Field{
	{Name: "AddReplicaRegions", Flag: "add-replica-regions", Type: "[]types.ReplicaRegionType", Required: true},
	{Name: "ForceOverwriteReplicaSecret", Flag: "force-overwrite-replica-secret", Type: "bool", Required: false},
	{Name: "SecretId", Flag: "secret-id", Type: "*string", Required: true},
}

var fields_restore_secret = []leanruntime.Field{
	{Name: "SecretId", Flag: "secret-id", Type: "*string", Required: true},
}

var fields_rotate_secret = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "ExternalSecretRotationMetadata", Flag: "external-secret-rotation-metadata", Type: "[]types.ExternalSecretRotationMetadataItem", Required: false},
	{Name: "ExternalSecretRotationRoleArn", Flag: "external-secret-rotation-role-arn", Type: "*string", Required: false},
	{Name: "RotateImmediately", Flag: "rotate-immediately", Type: "*bool", Required: false},
	{Name: "RotationLambdaARN", Flag: "rotation-lambda-arn", Type: "*string", Required: false},
	{Name: "RotationRules", Flag: "rotation-rules", Type: "*types.RotationRulesType", Required: false},
	{Name: "SecretId", Flag: "secret-id", Type: "*string", Required: true},
}

var fields_stop_replication_to_replica = []leanruntime.Field{
	{Name: "SecretId", Flag: "secret-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "SecretId", Flag: "secret-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "SecretId", Flag: "secret-id", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_secret = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "SecretBinary", Flag: "secret-binary", Type: "[]byte", Required: false},
	{Name: "SecretId", Flag: "secret-id", Type: "*string", Required: true},
	{Name: "SecretString", Flag: "secret-string", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "*string", Required: false},
}

var fields_update_secret_version_stage = []leanruntime.Field{
	{Name: "MoveToVersionId", Flag: "move-to-version-id", Type: "*string", Required: false},
	{Name: "RemoveFromVersionId", Flag: "remove-from-version-id", Type: "*string", Required: false},
	{Name: "SecretId", Flag: "secret-id", Type: "*string", Required: true},
	{Name: "VersionStage", Flag: "version-stage", Type: "*string", Required: true},
}

var fields_validate_resource_policy = []leanruntime.Field{
	{Name: "ResourcePolicy", Flag: "resource-policy", Type: "*string", Required: true},
	{Name: "SecretId", Flag: "secret-id", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"batch-get-secret-value": {
			Name:   "batch-get-secret-value",
			Fields: fields_batch_get_secret_value,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetSecretValueInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_batch_get_secret_value, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.BatchGetSecretValue(ctx, input)
				}
				var results []*svc.BatchGetSecretValueOutput
				p := svc.NewBatchGetSecretValuePaginator(client, input)
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
		"cancel-rotate-secret": {
			Name:   "cancel-rotate-secret",
			Fields: fields_cancel_rotate_secret,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelRotateSecretInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_rotate_secret, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelRotateSecret(ctx, input)
			},
		},
		"create-secret": {
			Name:   "create-secret",
			Fields: fields_create_secret,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSecretInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_secret, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSecret(ctx, input)
			},
		},
		"delete-resource-policy": {
			Name:   "delete-resource-policy",
			Fields: fields_delete_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteResourcePolicy(ctx, input)
			},
		},
		"delete-secret": {
			Name:   "delete-secret",
			Fields: fields_delete_secret,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSecretInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_secret, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSecret(ctx, input)
			},
		},
		"describe-secret": {
			Name:   "describe-secret",
			Fields: fields_describe_secret,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSecretInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_secret, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSecret(ctx, input)
			},
		},
		"get-random-password": {
			Name:   "get-random-password",
			Fields: fields_get_random_password,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRandomPasswordInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_random_password, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRandomPassword(ctx, input)
			},
		},
		"get-resource-policy": {
			Name:   "get-resource-policy",
			Fields: fields_get_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResourcePolicy(ctx, input)
			},
		},
		"get-secret-value": {
			Name:   "get-secret-value",
			Fields: fields_get_secret_value,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSecretValueInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_secret_value, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSecretValue(ctx, input)
			},
		},
		"list-secret-version-ids": {
			Name:   "list-secret-version-ids",
			Fields: fields_list_secret_version_ids,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSecretVersionIdsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_secret_version_ids, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSecretVersionIds(ctx, input)
				}
				var results []*svc.ListSecretVersionIdsOutput
				p := svc.NewListSecretVersionIdsPaginator(client, input)
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
		"list-secrets": {
			Name:   "list-secrets",
			Fields: fields_list_secrets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSecretsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_secrets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSecrets(ctx, input)
				}
				var results []*svc.ListSecretsOutput
				p := svc.NewListSecretsPaginator(client, input)
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
		"put-resource-policy": {
			Name:   "put-resource-policy",
			Fields: fields_put_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutResourcePolicy(ctx, input)
			},
		},
		"put-secret-value": {
			Name:   "put-secret-value",
			Fields: fields_put_secret_value,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutSecretValueInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_secret_value, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutSecretValue(ctx, input)
			},
		},
		"remove-regions-from-replication": {
			Name:   "remove-regions-from-replication",
			Fields: fields_remove_regions_from_replication,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveRegionsFromReplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_regions_from_replication, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveRegionsFromReplication(ctx, input)
			},
		},
		"replicate-secret-to-regions": {
			Name:   "replicate-secret-to-regions",
			Fields: fields_replicate_secret_to_regions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ReplicateSecretToRegionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_replicate_secret_to_regions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ReplicateSecretToRegions(ctx, input)
			},
		},
		"restore-secret": {
			Name:   "restore-secret",
			Fields: fields_restore_secret,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RestoreSecretInput{}
				if _, err := leanruntime.ApplyInput(input, fields_restore_secret, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RestoreSecret(ctx, input)
			},
		},
		"rotate-secret": {
			Name:   "rotate-secret",
			Fields: fields_rotate_secret,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RotateSecretInput{}
				if _, err := leanruntime.ApplyInput(input, fields_rotate_secret, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RotateSecret(ctx, input)
			},
		},
		"stop-replication-to-replica": {
			Name:   "stop-replication-to-replica",
			Fields: fields_stop_replication_to_replica,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopReplicationToReplicaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_replication_to_replica, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopReplicationToReplica(ctx, input)
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
		"update-secret": {
			Name:   "update-secret",
			Fields: fields_update_secret,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSecretInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_secret, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSecret(ctx, input)
			},
		},
		"update-secret-version-stage": {
			Name:   "update-secret-version-stage",
			Fields: fields_update_secret_version_stage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSecretVersionStageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_secret_version_stage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSecretVersionStage(ctx, input)
			},
		},
		"validate-resource-policy": {
			Name:   "validate-resource-policy",
			Fields: fields_validate_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ValidateResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_validate_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ValidateResourcePolicy(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("secretsmanager", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
