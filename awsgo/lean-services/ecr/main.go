package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/ecr"
)

var fields_batch_check_layer_availability = []leanruntime.Field{
	{Name: "LayerDigests", Flag: "layer-digests", Type: "[]string", Required: true},
	{Name: "RegistryId", Flag: "registry-id", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_batch_delete_image = []leanruntime.Field{
	{Name: "ImageIds", Flag: "image-ids", Type: "[]types.ImageIdentifier", Required: true},
	{Name: "RegistryId", Flag: "registry-id", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_batch_get_image = []leanruntime.Field{
	{Name: "AcceptedMediaTypes", Flag: "accepted-media-types", Type: "[]string", Required: false},
	{Name: "ImageIds", Flag: "image-ids", Type: "[]types.ImageIdentifier", Required: true},
	{Name: "RegistryId", Flag: "registry-id", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_batch_get_repository_scanning_configuration = []leanruntime.Field{
	{Name: "RepositoryNames", Flag: "repository-names", Type: "[]string", Required: true},
}

var fields_complete_layer_upload = []leanruntime.Field{
	{Name: "LayerDigests", Flag: "layer-digests", Type: "[]string", Required: true},
	{Name: "RegistryId", Flag: "registry-id", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
	{Name: "UploadId", Flag: "upload-id", Type: "*string", Required: true},
}

var fields_create_pull_through_cache_rule = []leanruntime.Field{
	{Name: "CredentialArn", Flag: "credential-arn", Type: "*string", Required: false},
	{Name: "CustomRoleArn", Flag: "custom-role-arn", Type: "*string", Required: false},
	{Name: "EcrRepositoryPrefix", Flag: "ecr-repository-prefix", Type: "*string", Required: true},
	{Name: "RegistryId", Flag: "registry-id", Type: "*string", Required: false},
	{Name: "UpstreamRegistry", Flag: "upstream-registry", Type: "types.UpstreamRegistry", Required: false},
	{Name: "UpstreamRegistryUrl", Flag: "upstream-registry-url", Type: "*string", Required: true},
	{Name: "UpstreamRepositoryPrefix", Flag: "upstream-repository-prefix", Type: "*string", Required: false},
}

var fields_create_repository = []leanruntime.Field{
	{Name: "EncryptionConfiguration", Flag: "encryption-configuration", Type: "*types.EncryptionConfiguration", Required: false},
	{Name: "ImageScanningConfiguration", Flag: "image-scanning-configuration", Type: "*types.ImageScanningConfiguration", Required: false},
	{Name: "ImageTagMutability", Flag: "image-tag-mutability", Type: "types.ImageTagMutability", Required: false},
	{Name: "ImageTagMutabilityExclusionFilters", Flag: "image-tag-mutability-exclusion-filters", Type: "[]types.ImageTagMutabilityExclusionFilter", Required: false},
	{Name: "RegistryId", Flag: "registry-id", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_repository_creation_template = []leanruntime.Field{
	{Name: "AppliedFor", Flag: "applied-for", Type: "[]types.RCTAppliedFor", Required: true},
	{Name: "CustomRoleArn", Flag: "custom-role-arn", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EncryptionConfiguration", Flag: "encryption-configuration", Type: "*types.EncryptionConfigurationForRepositoryCreationTemplate", Required: false},
	{Name: "ImageTagMutability", Flag: "image-tag-mutability", Type: "types.ImageTagMutability", Required: false},
	{Name: "ImageTagMutabilityExclusionFilters", Flag: "image-tag-mutability-exclusion-filters", Type: "[]types.ImageTagMutabilityExclusionFilter", Required: false},
	{Name: "LifecyclePolicy", Flag: "lifecycle-policy", Type: "*string", Required: false},
	{Name: "Prefix", Flag: "prefix", Type: "*string", Required: true},
	{Name: "RepositoryPolicy", Flag: "repository-policy", Type: "*string", Required: false},
	{Name: "ResourceTags", Flag: "resource-tags", Type: "[]types.Tag", Required: false},
}

var fields_delete_lifecycle_policy = []leanruntime.Field{
	{Name: "RegistryId", Flag: "registry-id", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_delete_pull_through_cache_rule = []leanruntime.Field{
	{Name: "EcrRepositoryPrefix", Flag: "ecr-repository-prefix", Type: "*string", Required: true},
	{Name: "RegistryId", Flag: "registry-id", Type: "*string", Required: false},
}

var fields_delete_registry_policy = []leanruntime.Field{}

var fields_delete_repository = []leanruntime.Field{
	{Name: "Force", Flag: "force", Type: "bool", Required: false},
	{Name: "RegistryId", Flag: "registry-id", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_delete_repository_creation_template = []leanruntime.Field{
	{Name: "Prefix", Flag: "prefix", Type: "*string", Required: true},
}

var fields_delete_repository_policy = []leanruntime.Field{
	{Name: "RegistryId", Flag: "registry-id", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_delete_signing_configuration = []leanruntime.Field{}

var fields_deregister_pull_time_update_exclusion = []leanruntime.Field{
	{Name: "PrincipalArn", Flag: "principal-arn", Type: "*string", Required: true},
}

var fields_describe_image_replication_status = []leanruntime.Field{
	{Name: "ImageId", Flag: "image-id", Type: "*types.ImageIdentifier", Required: true},
	{Name: "RegistryId", Flag: "registry-id", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_describe_image_scan_findings = []leanruntime.Field{
	{Name: "ImageId", Flag: "image-id", Type: "*types.ImageIdentifier", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RegistryId", Flag: "registry-id", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_describe_image_signing_status = []leanruntime.Field{
	{Name: "ImageId", Flag: "image-id", Type: "*types.ImageIdentifier", Required: true},
	{Name: "RegistryId", Flag: "registry-id", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_describe_images = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.DescribeImagesFilter", Required: false},
	{Name: "ImageIds", Flag: "image-ids", Type: "[]types.ImageIdentifier", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RegistryId", Flag: "registry-id", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_describe_pull_through_cache_rules = []leanruntime.Field{
	{Name: "EcrRepositoryPrefixes", Flag: "ecr-repository-prefixes", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RegistryId", Flag: "registry-id", Type: "*string", Required: false},
}

var fields_describe_registry = []leanruntime.Field{}

var fields_describe_repositories = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RegistryId", Flag: "registry-id", Type: "*string", Required: false},
	{Name: "RepositoryNames", Flag: "repository-names", Type: "[]string", Required: false},
}

var fields_describe_repository_creation_templates = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Prefixes", Flag: "prefixes", Type: "[]string", Required: false},
}

var fields_get_account_setting = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_authorization_token = []leanruntime.Field{
	{Name: "RegistryIds", Flag: "registry-ids", Type: "[]string", Required: false},
}

var fields_get_download_url_for_layer = []leanruntime.Field{
	{Name: "LayerDigest", Flag: "layer-digest", Type: "*string", Required: true},
	{Name: "RegistryId", Flag: "registry-id", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_get_lifecycle_policy = []leanruntime.Field{
	{Name: "RegistryId", Flag: "registry-id", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_get_lifecycle_policy_preview = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.LifecyclePolicyPreviewFilter", Required: false},
	{Name: "ImageIds", Flag: "image-ids", Type: "[]types.ImageIdentifier", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RegistryId", Flag: "registry-id", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_get_registry_policy = []leanruntime.Field{}

var fields_get_registry_scanning_configuration = []leanruntime.Field{}

var fields_get_repository_policy = []leanruntime.Field{
	{Name: "RegistryId", Flag: "registry-id", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_get_signing_configuration = []leanruntime.Field{}

var fields_initiate_layer_upload = []leanruntime.Field{
	{Name: "RegistryId", Flag: "registry-id", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_list_image_referrers = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.ListImageReferrersFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RegistryId", Flag: "registry-id", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
	{Name: "SubjectId", Flag: "subject-id", Type: "*types.SubjectIdentifier", Required: true},
}

var fields_list_images = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.ListImagesFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RegistryId", Flag: "registry-id", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_list_pull_time_update_exclusions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_put_account_setting = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Value", Flag: "value", Type: "*string", Required: true},
}

var fields_put_image = []leanruntime.Field{
	{Name: "ImageDigest", Flag: "image-digest", Type: "*string", Required: false},
	{Name: "ImageManifest", Flag: "image-manifest", Type: "*string", Required: true},
	{Name: "ImageManifestMediaType", Flag: "image-manifest-media-type", Type: "*string", Required: false},
	{Name: "ImageTag", Flag: "image-tag", Type: "*string", Required: false},
	{Name: "RegistryId", Flag: "registry-id", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_put_image_scanning_configuration = []leanruntime.Field{
	{Name: "ImageScanningConfiguration", Flag: "image-scanning-configuration", Type: "*types.ImageScanningConfiguration", Required: true},
	{Name: "RegistryId", Flag: "registry-id", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_put_image_tag_mutability = []leanruntime.Field{
	{Name: "ImageTagMutability", Flag: "image-tag-mutability", Type: "types.ImageTagMutability", Required: true},
	{Name: "ImageTagMutabilityExclusionFilters", Flag: "image-tag-mutability-exclusion-filters", Type: "[]types.ImageTagMutabilityExclusionFilter", Required: false},
	{Name: "RegistryId", Flag: "registry-id", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_put_lifecycle_policy = []leanruntime.Field{
	{Name: "LifecyclePolicyText", Flag: "lifecycle-policy-text", Type: "*string", Required: true},
	{Name: "RegistryId", Flag: "registry-id", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_put_registry_policy = []leanruntime.Field{
	{Name: "PolicyText", Flag: "policy-text", Type: "*string", Required: true},
}

var fields_put_registry_scanning_configuration = []leanruntime.Field{
	{Name: "Rules", Flag: "rules", Type: "[]types.RegistryScanningRule", Required: false},
	{Name: "ScanType", Flag: "scan-type", Type: "types.ScanType", Required: false},
}

var fields_put_replication_configuration = []leanruntime.Field{
	{Name: "ReplicationConfiguration", Flag: "replication-configuration", Type: "*types.ReplicationConfiguration", Required: true},
}

var fields_put_signing_configuration = []leanruntime.Field{
	{Name: "SigningConfiguration", Flag: "signing-configuration", Type: "*types.SigningConfiguration", Required: true},
}

var fields_register_pull_time_update_exclusion = []leanruntime.Field{
	{Name: "PrincipalArn", Flag: "principal-arn", Type: "*string", Required: true},
}

var fields_set_repository_policy = []leanruntime.Field{
	{Name: "Force", Flag: "force", Type: "bool", Required: false},
	{Name: "PolicyText", Flag: "policy-text", Type: "*string", Required: true},
	{Name: "RegistryId", Flag: "registry-id", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_start_image_scan = []leanruntime.Field{
	{Name: "ImageId", Flag: "image-id", Type: "*types.ImageIdentifier", Required: true},
	{Name: "RegistryId", Flag: "registry-id", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_start_lifecycle_policy_preview = []leanruntime.Field{
	{Name: "LifecyclePolicyText", Flag: "lifecycle-policy-text", Type: "*string", Required: false},
	{Name: "RegistryId", Flag: "registry-id", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_image_storage_class = []leanruntime.Field{
	{Name: "ImageId", Flag: "image-id", Type: "*types.ImageIdentifier", Required: true},
	{Name: "RegistryId", Flag: "registry-id", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
	{Name: "TargetStorageClass", Flag: "target-storage-class", Type: "types.TargetStorageClass", Required: true},
}

var fields_update_pull_through_cache_rule = []leanruntime.Field{
	{Name: "CredentialArn", Flag: "credential-arn", Type: "*string", Required: false},
	{Name: "CustomRoleArn", Flag: "custom-role-arn", Type: "*string", Required: false},
	{Name: "EcrRepositoryPrefix", Flag: "ecr-repository-prefix", Type: "*string", Required: true},
	{Name: "RegistryId", Flag: "registry-id", Type: "*string", Required: false},
}

var fields_update_repository_creation_template = []leanruntime.Field{
	{Name: "AppliedFor", Flag: "applied-for", Type: "[]types.RCTAppliedFor", Required: false},
	{Name: "CustomRoleArn", Flag: "custom-role-arn", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EncryptionConfiguration", Flag: "encryption-configuration", Type: "*types.EncryptionConfigurationForRepositoryCreationTemplate", Required: false},
	{Name: "ImageTagMutability", Flag: "image-tag-mutability", Type: "types.ImageTagMutability", Required: false},
	{Name: "ImageTagMutabilityExclusionFilters", Flag: "image-tag-mutability-exclusion-filters", Type: "[]types.ImageTagMutabilityExclusionFilter", Required: false},
	{Name: "LifecyclePolicy", Flag: "lifecycle-policy", Type: "*string", Required: false},
	{Name: "Prefix", Flag: "prefix", Type: "*string", Required: true},
	{Name: "RepositoryPolicy", Flag: "repository-policy", Type: "*string", Required: false},
	{Name: "ResourceTags", Flag: "resource-tags", Type: "[]types.Tag", Required: false},
}

var fields_upload_layer_part = []leanruntime.Field{
	{Name: "LayerPartBlob", Flag: "layer-part-blob", Type: "[]byte", Required: true},
	{Name: "PartFirstByte", Flag: "part-first-byte", Type: "*int64", Required: true},
	{Name: "PartLastByte", Flag: "part-last-byte", Type: "*int64", Required: true},
	{Name: "RegistryId", Flag: "registry-id", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
	{Name: "UploadId", Flag: "upload-id", Type: "*string", Required: true},
}

var fields_validate_pull_through_cache_rule = []leanruntime.Field{
	{Name: "EcrRepositoryPrefix", Flag: "ecr-repository-prefix", Type: "*string", Required: true},
	{Name: "RegistryId", Flag: "registry-id", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"batch-check-layer-availability": {
			Name:   "batch-check-layer-availability",
			Fields: fields_batch_check_layer_availability,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchCheckLayerAvailabilityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_check_layer_availability, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchCheckLayerAvailability(ctx, input)
			},
		},
		"batch-delete-image": {
			Name:   "batch-delete-image",
			Fields: fields_batch_delete_image,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDeleteImageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_delete_image, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDeleteImage(ctx, input)
			},
		},
		"batch-get-image": {
			Name:   "batch-get-image",
			Fields: fields_batch_get_image,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetImageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_image, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetImage(ctx, input)
			},
		},
		"batch-get-repository-scanning-configuration": {
			Name:   "batch-get-repository-scanning-configuration",
			Fields: fields_batch_get_repository_scanning_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetRepositoryScanningConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_repository_scanning_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetRepositoryScanningConfiguration(ctx, input)
			},
		},
		"complete-layer-upload": {
			Name:   "complete-layer-upload",
			Fields: fields_complete_layer_upload,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CompleteLayerUploadInput{}
				if _, err := leanruntime.ApplyInput(input, fields_complete_layer_upload, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CompleteLayerUpload(ctx, input)
			},
		},
		"create-pull-through-cache-rule": {
			Name:   "create-pull-through-cache-rule",
			Fields: fields_create_pull_through_cache_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePullThroughCacheRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_pull_through_cache_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePullThroughCacheRule(ctx, input)
			},
		},
		"create-repository": {
			Name:   "create-repository",
			Fields: fields_create_repository,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRepositoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_repository, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRepository(ctx, input)
			},
		},
		"create-repository-creation-template": {
			Name:   "create-repository-creation-template",
			Fields: fields_create_repository_creation_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRepositoryCreationTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_repository_creation_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRepositoryCreationTemplate(ctx, input)
			},
		},
		"delete-lifecycle-policy": {
			Name:   "delete-lifecycle-policy",
			Fields: fields_delete_lifecycle_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLifecyclePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_lifecycle_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLifecyclePolicy(ctx, input)
			},
		},
		"delete-pull-through-cache-rule": {
			Name:   "delete-pull-through-cache-rule",
			Fields: fields_delete_pull_through_cache_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePullThroughCacheRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_pull_through_cache_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePullThroughCacheRule(ctx, input)
			},
		},
		"delete-registry-policy": {
			Name:   "delete-registry-policy",
			Fields: fields_delete_registry_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRegistryPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_registry_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRegistryPolicy(ctx, input)
			},
		},
		"delete-repository": {
			Name:   "delete-repository",
			Fields: fields_delete_repository,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRepositoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_repository, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRepository(ctx, input)
			},
		},
		"delete-repository-creation-template": {
			Name:   "delete-repository-creation-template",
			Fields: fields_delete_repository_creation_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRepositoryCreationTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_repository_creation_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRepositoryCreationTemplate(ctx, input)
			},
		},
		"delete-repository-policy": {
			Name:   "delete-repository-policy",
			Fields: fields_delete_repository_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRepositoryPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_repository_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRepositoryPolicy(ctx, input)
			},
		},
		"delete-signing-configuration": {
			Name:   "delete-signing-configuration",
			Fields: fields_delete_signing_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSigningConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_signing_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSigningConfiguration(ctx, input)
			},
		},
		"deregister-pull-time-update-exclusion": {
			Name:   "deregister-pull-time-update-exclusion",
			Fields: fields_deregister_pull_time_update_exclusion,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterPullTimeUpdateExclusionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_pull_time_update_exclusion, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterPullTimeUpdateExclusion(ctx, input)
			},
		},
		"describe-image-replication-status": {
			Name:   "describe-image-replication-status",
			Fields: fields_describe_image_replication_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeImageReplicationStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_image_replication_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeImageReplicationStatus(ctx, input)
			},
		},
		"describe-image-scan-findings": {
			Name:   "describe-image-scan-findings",
			Fields: fields_describe_image_scan_findings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeImageScanFindingsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_image_scan_findings, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeImageScanFindings(ctx, input)
				}
				var results []*svc.DescribeImageScanFindingsOutput
				p := svc.NewDescribeImageScanFindingsPaginator(client, input)
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
		"describe-image-signing-status": {
			Name:   "describe-image-signing-status",
			Fields: fields_describe_image_signing_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeImageSigningStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_image_signing_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeImageSigningStatus(ctx, input)
			},
		},
		"describe-images": {
			Name:   "describe-images",
			Fields: fields_describe_images,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeImagesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_images, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeImages(ctx, input)
				}
				var results []*svc.DescribeImagesOutput
				p := svc.NewDescribeImagesPaginator(client, input)
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
		"describe-pull-through-cache-rules": {
			Name:   "describe-pull-through-cache-rules",
			Fields: fields_describe_pull_through_cache_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePullThroughCacheRulesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_pull_through_cache_rules, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribePullThroughCacheRules(ctx, input)
				}
				var results []*svc.DescribePullThroughCacheRulesOutput
				p := svc.NewDescribePullThroughCacheRulesPaginator(client, input)
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
		"describe-registry": {
			Name:   "describe-registry",
			Fields: fields_describe_registry,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRegistryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_registry, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeRegistry(ctx, input)
			},
		},
		"describe-repositories": {
			Name:   "describe-repositories",
			Fields: fields_describe_repositories,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRepositoriesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_repositories, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeRepositories(ctx, input)
				}
				var results []*svc.DescribeRepositoriesOutput
				p := svc.NewDescribeRepositoriesPaginator(client, input)
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
		"describe-repository-creation-templates": {
			Name:   "describe-repository-creation-templates",
			Fields: fields_describe_repository_creation_templates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRepositoryCreationTemplatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_repository_creation_templates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeRepositoryCreationTemplates(ctx, input)
				}
				var results []*svc.DescribeRepositoryCreationTemplatesOutput
				p := svc.NewDescribeRepositoryCreationTemplatesPaginator(client, input)
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
		"get-account-setting": {
			Name:   "get-account-setting",
			Fields: fields_get_account_setting,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccountSettingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_account_setting, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccountSetting(ctx, input)
			},
		},
		"get-authorization-token": {
			Name:   "get-authorization-token",
			Fields: fields_get_authorization_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAuthorizationTokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_authorization_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAuthorizationToken(ctx, input)
			},
		},
		"get-download-url-for-layer": {
			Name:   "get-download-url-for-layer",
			Fields: fields_get_download_url_for_layer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDownloadUrlForLayerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_download_url_for_layer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDownloadUrlForLayer(ctx, input)
			},
		},
		"get-lifecycle-policy": {
			Name:   "get-lifecycle-policy",
			Fields: fields_get_lifecycle_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLifecyclePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_lifecycle_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLifecyclePolicy(ctx, input)
			},
		},
		"get-lifecycle-policy-preview": {
			Name:   "get-lifecycle-policy-preview",
			Fields: fields_get_lifecycle_policy_preview,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLifecyclePolicyPreviewInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_lifecycle_policy_preview, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetLifecyclePolicyPreview(ctx, input)
				}
				var results []*svc.GetLifecyclePolicyPreviewOutput
				p := svc.NewGetLifecyclePolicyPreviewPaginator(client, input)
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
		"get-registry-policy": {
			Name:   "get-registry-policy",
			Fields: fields_get_registry_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRegistryPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_registry_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRegistryPolicy(ctx, input)
			},
		},
		"get-registry-scanning-configuration": {
			Name:   "get-registry-scanning-configuration",
			Fields: fields_get_registry_scanning_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRegistryScanningConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_registry_scanning_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRegistryScanningConfiguration(ctx, input)
			},
		},
		"get-repository-policy": {
			Name:   "get-repository-policy",
			Fields: fields_get_repository_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRepositoryPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_repository_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRepositoryPolicy(ctx, input)
			},
		},
		"get-signing-configuration": {
			Name:   "get-signing-configuration",
			Fields: fields_get_signing_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSigningConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_signing_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSigningConfiguration(ctx, input)
			},
		},
		"initiate-layer-upload": {
			Name:   "initiate-layer-upload",
			Fields: fields_initiate_layer_upload,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.InitiateLayerUploadInput{}
				if _, err := leanruntime.ApplyInput(input, fields_initiate_layer_upload, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.InitiateLayerUpload(ctx, input)
			},
		},
		"list-image-referrers": {
			Name:   "list-image-referrers",
			Fields: fields_list_image_referrers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListImageReferrersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_image_referrers, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListImageReferrers(ctx, input)
			},
		},
		"list-images": {
			Name:   "list-images",
			Fields: fields_list_images,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListImagesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_images, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListImages(ctx, input)
				}
				var results []*svc.ListImagesOutput
				p := svc.NewListImagesPaginator(client, input)
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
		"list-pull-time-update-exclusions": {
			Name:   "list-pull-time-update-exclusions",
			Fields: fields_list_pull_time_update_exclusions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPullTimeUpdateExclusionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_pull_time_update_exclusions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListPullTimeUpdateExclusions(ctx, input)
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
		"put-account-setting": {
			Name:   "put-account-setting",
			Fields: fields_put_account_setting,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutAccountSettingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_account_setting, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutAccountSetting(ctx, input)
			},
		},
		"put-image": {
			Name:   "put-image",
			Fields: fields_put_image,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutImageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_image, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutImage(ctx, input)
			},
		},
		"put-image-scanning-configuration": {
			Name:   "put-image-scanning-configuration",
			Fields: fields_put_image_scanning_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutImageScanningConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_image_scanning_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutImageScanningConfiguration(ctx, input)
			},
		},
		"put-image-tag-mutability": {
			Name:   "put-image-tag-mutability",
			Fields: fields_put_image_tag_mutability,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutImageTagMutabilityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_image_tag_mutability, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutImageTagMutability(ctx, input)
			},
		},
		"put-lifecycle-policy": {
			Name:   "put-lifecycle-policy",
			Fields: fields_put_lifecycle_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutLifecyclePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_lifecycle_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutLifecyclePolicy(ctx, input)
			},
		},
		"put-registry-policy": {
			Name:   "put-registry-policy",
			Fields: fields_put_registry_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutRegistryPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_registry_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutRegistryPolicy(ctx, input)
			},
		},
		"put-registry-scanning-configuration": {
			Name:   "put-registry-scanning-configuration",
			Fields: fields_put_registry_scanning_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutRegistryScanningConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_registry_scanning_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutRegistryScanningConfiguration(ctx, input)
			},
		},
		"put-replication-configuration": {
			Name:   "put-replication-configuration",
			Fields: fields_put_replication_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutReplicationConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_replication_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutReplicationConfiguration(ctx, input)
			},
		},
		"put-signing-configuration": {
			Name:   "put-signing-configuration",
			Fields: fields_put_signing_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutSigningConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_signing_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutSigningConfiguration(ctx, input)
			},
		},
		"register-pull-time-update-exclusion": {
			Name:   "register-pull-time-update-exclusion",
			Fields: fields_register_pull_time_update_exclusion,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterPullTimeUpdateExclusionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_pull_time_update_exclusion, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterPullTimeUpdateExclusion(ctx, input)
			},
		},
		"set-repository-policy": {
			Name:   "set-repository-policy",
			Fields: fields_set_repository_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetRepositoryPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_repository_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetRepositoryPolicy(ctx, input)
			},
		},
		"start-image-scan": {
			Name:   "start-image-scan",
			Fields: fields_start_image_scan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartImageScanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_image_scan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartImageScan(ctx, input)
			},
		},
		"start-lifecycle-policy-preview": {
			Name:   "start-lifecycle-policy-preview",
			Fields: fields_start_lifecycle_policy_preview,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartLifecyclePolicyPreviewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_lifecycle_policy_preview, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartLifecyclePolicyPreview(ctx, input)
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
		"update-image-storage-class": {
			Name:   "update-image-storage-class",
			Fields: fields_update_image_storage_class,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateImageStorageClassInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_image_storage_class, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateImageStorageClass(ctx, input)
			},
		},
		"update-pull-through-cache-rule": {
			Name:   "update-pull-through-cache-rule",
			Fields: fields_update_pull_through_cache_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePullThroughCacheRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_pull_through_cache_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePullThroughCacheRule(ctx, input)
			},
		},
		"update-repository-creation-template": {
			Name:   "update-repository-creation-template",
			Fields: fields_update_repository_creation_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRepositoryCreationTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_repository_creation_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRepositoryCreationTemplate(ctx, input)
			},
		},
		"upload-layer-part": {
			Name:   "upload-layer-part",
			Fields: fields_upload_layer_part,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UploadLayerPartInput{}
				if _, err := leanruntime.ApplyInput(input, fields_upload_layer_part, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UploadLayerPart(ctx, input)
			},
		},
		"validate-pull-through-cache-rule": {
			Name:   "validate-pull-through-cache-rule",
			Fields: fields_validate_pull_through_cache_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ValidatePullThroughCacheRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_validate_pull_through_cache_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ValidatePullThroughCacheRule(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("ecr", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
