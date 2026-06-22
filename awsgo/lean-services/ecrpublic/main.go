package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/ecrpublic"
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

var fields_complete_layer_upload = []leanruntime.Field{
	{Name: "LayerDigests", Flag: "layer-digests", Type: "[]string", Required: true},
	{Name: "RegistryId", Flag: "registry-id", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
	{Name: "UploadId", Flag: "upload-id", Type: "*string", Required: true},
}

var fields_create_repository = []leanruntime.Field{
	{Name: "CatalogData", Flag: "catalog-data", Type: "*types.RepositoryCatalogDataInput", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_delete_repository = []leanruntime.Field{
	{Name: "Force", Flag: "force", Type: "bool", Required: false},
	{Name: "RegistryId", Flag: "registry-id", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_delete_repository_policy = []leanruntime.Field{
	{Name: "RegistryId", Flag: "registry-id", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_describe_image_tags = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RegistryId", Flag: "registry-id", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_describe_images = []leanruntime.Field{
	{Name: "ImageIds", Flag: "image-ids", Type: "[]types.ImageIdentifier", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RegistryId", Flag: "registry-id", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_describe_registries = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_repositories = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RegistryId", Flag: "registry-id", Type: "*string", Required: false},
	{Name: "RepositoryNames", Flag: "repository-names", Type: "[]string", Required: false},
}

var fields_get_authorization_token = []leanruntime.Field{}

var fields_get_registry_catalog_data = []leanruntime.Field{}

var fields_get_repository_catalog_data = []leanruntime.Field{
	{Name: "RegistryId", Flag: "registry-id", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_get_repository_policy = []leanruntime.Field{
	{Name: "RegistryId", Flag: "registry-id", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_initiate_layer_upload = []leanruntime.Field{
	{Name: "RegistryId", Flag: "registry-id", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_put_image = []leanruntime.Field{
	{Name: "ImageDigest", Flag: "image-digest", Type: "*string", Required: false},
	{Name: "ImageManifest", Flag: "image-manifest", Type: "*string", Required: true},
	{Name: "ImageManifestMediaType", Flag: "image-manifest-media-type", Type: "*string", Required: false},
	{Name: "ImageTag", Flag: "image-tag", Type: "*string", Required: false},
	{Name: "RegistryId", Flag: "registry-id", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_put_registry_catalog_data = []leanruntime.Field{
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
}

var fields_put_repository_catalog_data = []leanruntime.Field{
	{Name: "CatalogData", Flag: "catalog-data", Type: "*types.RepositoryCatalogDataInput", Required: true},
	{Name: "RegistryId", Flag: "registry-id", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_set_repository_policy = []leanruntime.Field{
	{Name: "Force", Flag: "force", Type: "bool", Required: false},
	{Name: "PolicyText", Flag: "policy-text", Type: "*string", Required: true},
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

var fields_upload_layer_part = []leanruntime.Field{
	{Name: "LayerPartBlob", Flag: "layer-part-blob", Type: "[]byte", Required: true},
	{Name: "PartFirstByte", Flag: "part-first-byte", Type: "*int64", Required: true},
	{Name: "PartLastByte", Flag: "part-last-byte", Type: "*int64", Required: true},
	{Name: "RegistryId", Flag: "registry-id", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
	{Name: "UploadId", Flag: "upload-id", Type: "*string", Required: true},
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
		"describe-image-tags": {
			Name:   "describe-image-tags",
			Fields: fields_describe_image_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeImageTagsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_image_tags, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeImageTags(ctx, input)
				}
				var results []*svc.DescribeImageTagsOutput
				p := svc.NewDescribeImageTagsPaginator(client, input)
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
		"describe-registries": {
			Name:   "describe-registries",
			Fields: fields_describe_registries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRegistriesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_registries, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeRegistries(ctx, input)
				}
				var results []*svc.DescribeRegistriesOutput
				p := svc.NewDescribeRegistriesPaginator(client, input)
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
		"get-registry-catalog-data": {
			Name:   "get-registry-catalog-data",
			Fields: fields_get_registry_catalog_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRegistryCatalogDataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_registry_catalog_data, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRegistryCatalogData(ctx, input)
			},
		},
		"get-repository-catalog-data": {
			Name:   "get-repository-catalog-data",
			Fields: fields_get_repository_catalog_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRepositoryCatalogDataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_repository_catalog_data, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRepositoryCatalogData(ctx, input)
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
		"put-registry-catalog-data": {
			Name:   "put-registry-catalog-data",
			Fields: fields_put_registry_catalog_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutRegistryCatalogDataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_registry_catalog_data, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutRegistryCatalogData(ctx, input)
			},
		},
		"put-repository-catalog-data": {
			Name:   "put-repository-catalog-data",
			Fields: fields_put_repository_catalog_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutRepositoryCatalogDataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_repository_catalog_data, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutRepositoryCatalogData(ctx, input)
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
	}
	if err := leanruntime.Execute("ecrpublic", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
