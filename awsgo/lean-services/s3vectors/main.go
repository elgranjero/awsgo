package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/s3vectors"
)

var fields_create_index = []leanruntime.Field{
	{Name: "DataType", Flag: "data-type", Type: "types.DataType", Required: true},
	{Name: "Dimension", Flag: "dimension", Type: "*int32", Required: true},
	{Name: "DistanceMetric", Flag: "distance-metric", Type: "types.DistanceMetric", Required: true},
	{Name: "EncryptionConfiguration", Flag: "encryption-configuration", Type: "*types.EncryptionConfiguration", Required: false},
	{Name: "IndexName", Flag: "index-name", Type: "*string", Required: true},
	{Name: "MetadataConfiguration", Flag: "metadata-configuration", Type: "*types.MetadataConfiguration", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "VectorBucketArn", Flag: "vector-bucket-arn", Type: "*string", Required: false},
	{Name: "VectorBucketName", Flag: "vector-bucket-name", Type: "*string", Required: false},
}

var fields_create_vector_bucket = []leanruntime.Field{
	{Name: "EncryptionConfiguration", Flag: "encryption-configuration", Type: "*types.EncryptionConfiguration", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "VectorBucketName", Flag: "vector-bucket-name", Type: "*string", Required: true},
}

var fields_delete_index = []leanruntime.Field{
	{Name: "IndexArn", Flag: "index-arn", Type: "*string", Required: false},
	{Name: "IndexName", Flag: "index-name", Type: "*string", Required: false},
	{Name: "VectorBucketName", Flag: "vector-bucket-name", Type: "*string", Required: false},
}

var fields_delete_vector_bucket = []leanruntime.Field{
	{Name: "VectorBucketArn", Flag: "vector-bucket-arn", Type: "*string", Required: false},
	{Name: "VectorBucketName", Flag: "vector-bucket-name", Type: "*string", Required: false},
}

var fields_delete_vector_bucket_policy = []leanruntime.Field{
	{Name: "VectorBucketArn", Flag: "vector-bucket-arn", Type: "*string", Required: false},
	{Name: "VectorBucketName", Flag: "vector-bucket-name", Type: "*string", Required: false},
}

var fields_delete_vectors = []leanruntime.Field{
	{Name: "IndexArn", Flag: "index-arn", Type: "*string", Required: false},
	{Name: "IndexName", Flag: "index-name", Type: "*string", Required: false},
	{Name: "Keys", Flag: "keys", Type: "[]string", Required: true},
	{Name: "VectorBucketName", Flag: "vector-bucket-name", Type: "*string", Required: false},
}

var fields_get_index = []leanruntime.Field{
	{Name: "IndexArn", Flag: "index-arn", Type: "*string", Required: false},
	{Name: "IndexName", Flag: "index-name", Type: "*string", Required: false},
	{Name: "VectorBucketName", Flag: "vector-bucket-name", Type: "*string", Required: false},
}

var fields_get_vector_bucket = []leanruntime.Field{
	{Name: "VectorBucketArn", Flag: "vector-bucket-arn", Type: "*string", Required: false},
	{Name: "VectorBucketName", Flag: "vector-bucket-name", Type: "*string", Required: false},
}

var fields_get_vector_bucket_policy = []leanruntime.Field{
	{Name: "VectorBucketArn", Flag: "vector-bucket-arn", Type: "*string", Required: false},
	{Name: "VectorBucketName", Flag: "vector-bucket-name", Type: "*string", Required: false},
}

var fields_get_vectors = []leanruntime.Field{
	{Name: "IndexArn", Flag: "index-arn", Type: "*string", Required: false},
	{Name: "IndexName", Flag: "index-name", Type: "*string", Required: false},
	{Name: "Keys", Flag: "keys", Type: "[]string", Required: true},
	{Name: "ReturnData", Flag: "return-data", Type: "bool", Required: false},
	{Name: "ReturnMetadata", Flag: "return-metadata", Type: "bool", Required: false},
	{Name: "VectorBucketName", Flag: "vector-bucket-name", Type: "*string", Required: false},
}

var fields_list_indexes = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Prefix", Flag: "prefix", Type: "*string", Required: false},
	{Name: "VectorBucketArn", Flag: "vector-bucket-arn", Type: "*string", Required: false},
	{Name: "VectorBucketName", Flag: "vector-bucket-name", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_vector_buckets = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Prefix", Flag: "prefix", Type: "*string", Required: false},
}

var fields_list_vectors = []leanruntime.Field{
	{Name: "IndexArn", Flag: "index-arn", Type: "*string", Required: false},
	{Name: "IndexName", Flag: "index-name", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ReturnData", Flag: "return-data", Type: "bool", Required: false},
	{Name: "ReturnMetadata", Flag: "return-metadata", Type: "bool", Required: false},
	{Name: "SegmentCount", Flag: "segment-count", Type: "*int32", Required: false},
	{Name: "SegmentIndex", Flag: "segment-index", Type: "int32", Required: false},
	{Name: "VectorBucketName", Flag: "vector-bucket-name", Type: "*string", Required: false},
}

var fields_put_vector_bucket_policy = []leanruntime.Field{
	{Name: "Policy", Flag: "policy", Type: "*string", Required: true},
	{Name: "VectorBucketArn", Flag: "vector-bucket-arn", Type: "*string", Required: false},
	{Name: "VectorBucketName", Flag: "vector-bucket-name", Type: "*string", Required: false},
}

var fields_put_vectors = []leanruntime.Field{
	{Name: "IndexArn", Flag: "index-arn", Type: "*string", Required: false},
	{Name: "IndexName", Flag: "index-name", Type: "*string", Required: false},
	{Name: "VectorBucketName", Flag: "vector-bucket-name", Type: "*string", Required: false},
	{Name: "Vectors", Flag: "vectors", Type: "[]types.PutInputVector", Required: true},
}

var fields_query_vectors = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "document.Interface", Required: false},
	{Name: "IndexArn", Flag: "index-arn", Type: "*string", Required: false},
	{Name: "IndexName", Flag: "index-name", Type: "*string", Required: false},
	{Name: "QueryVector", Flag: "query-vector", Type: "types.VectorData", Required: true},
	{Name: "ReturnDistance", Flag: "return-distance", Type: "bool", Required: false},
	{Name: "ReturnMetadata", Flag: "return-metadata", Type: "bool", Required: false},
	{Name: "TopK", Flag: "top-k", Type: "*int32", Required: true},
	{Name: "VectorBucketName", Flag: "vector-bucket-name", Type: "*string", Required: false},
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
		"create-index": {
			Name:   "create-index",
			Fields: fields_create_index,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateIndexInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_index, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateIndex(ctx, input)
			},
		},
		"create-vector-bucket": {
			Name:   "create-vector-bucket",
			Fields: fields_create_vector_bucket,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVectorBucketInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_vector_bucket, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVectorBucket(ctx, input)
			},
		},
		"delete-index": {
			Name:   "delete-index",
			Fields: fields_delete_index,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIndexInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_index, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIndex(ctx, input)
			},
		},
		"delete-vector-bucket": {
			Name:   "delete-vector-bucket",
			Fields: fields_delete_vector_bucket,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVectorBucketInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_vector_bucket, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVectorBucket(ctx, input)
			},
		},
		"delete-vector-bucket-policy": {
			Name:   "delete-vector-bucket-policy",
			Fields: fields_delete_vector_bucket_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVectorBucketPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_vector_bucket_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVectorBucketPolicy(ctx, input)
			},
		},
		"delete-vectors": {
			Name:   "delete-vectors",
			Fields: fields_delete_vectors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVectorsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_vectors, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVectors(ctx, input)
			},
		},
		"get-index": {
			Name:   "get-index",
			Fields: fields_get_index,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIndexInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_index, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIndex(ctx, input)
			},
		},
		"get-vector-bucket": {
			Name:   "get-vector-bucket",
			Fields: fields_get_vector_bucket,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetVectorBucketInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_vector_bucket, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetVectorBucket(ctx, input)
			},
		},
		"get-vector-bucket-policy": {
			Name:   "get-vector-bucket-policy",
			Fields: fields_get_vector_bucket_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetVectorBucketPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_vector_bucket_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetVectorBucketPolicy(ctx, input)
			},
		},
		"get-vectors": {
			Name:   "get-vectors",
			Fields: fields_get_vectors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetVectorsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_vectors, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetVectors(ctx, input)
			},
		},
		"list-indexes": {
			Name:   "list-indexes",
			Fields: fields_list_indexes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIndexesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_indexes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListIndexes(ctx, input)
				}
				var results []*svc.ListIndexesOutput
				p := svc.NewListIndexesPaginator(client, input)
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
		"list-vector-buckets": {
			Name:   "list-vector-buckets",
			Fields: fields_list_vector_buckets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVectorBucketsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_vector_buckets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListVectorBuckets(ctx, input)
				}
				var results []*svc.ListVectorBucketsOutput
				p := svc.NewListVectorBucketsPaginator(client, input)
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
		"list-vectors": {
			Name:   "list-vectors",
			Fields: fields_list_vectors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVectorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_vectors, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListVectors(ctx, input)
				}
				var results []*svc.ListVectorsOutput
				p := svc.NewListVectorsPaginator(client, input)
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
		"put-vector-bucket-policy": {
			Name:   "put-vector-bucket-policy",
			Fields: fields_put_vector_bucket_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutVectorBucketPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_vector_bucket_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutVectorBucketPolicy(ctx, input)
			},
		},
		"put-vectors": {
			Name:   "put-vectors",
			Fields: fields_put_vectors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutVectorsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_vectors, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutVectors(ctx, input)
			},
		},
		"query-vectors": {
			Name:   "query-vectors",
			Fields: fields_query_vectors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.QueryVectorsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_query_vectors, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.QueryVectors(ctx, input)
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
	if err := leanruntime.Execute("s3vectors", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
