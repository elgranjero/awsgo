package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/s3vectors/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-index", "create-vector-bucket", "delete-index", "delete-vector-bucket", "delete-vector-bucket-policy", "delete-vectors", "get-index", "get-vector-bucket", "get-vector-bucket-policy", "get-vectors", "list-indexes", "list-tags-for-resource", "list-vector-buckets", "list-vectors", "put-vector-bucket-policy", "put-vectors", "query-vectors", "tag-resource", "untag-resource"},
		OperationSet: map[string]bool{"create-index": true, "create-vector-bucket": true, "delete-index": true, "delete-vector-bucket": true, "delete-vector-bucket-policy": true, "delete-vectors": true, "get-index": true, "get-vector-bucket": true, "get-vector-bucket-policy": true, "get-vectors": true, "list-indexes": true, "list-tags-for-resource": true, "list-vector-buckets": true, "list-vectors": true, "put-vector-bucket-policy": true, "put-vectors": true, "query-vectors": true, "tag-resource": true, "untag-resource": true},
		OperationInputs: map[string][]string{
			"create-index":                {"DataType", "Dimension", "DistanceMetric", "EncryptionConfiguration", "IndexName", "MetadataConfiguration", "Tags", "VectorBucketArn", "VectorBucketName"},
			"create-vector-bucket":        {"EncryptionConfiguration", "Tags", "VectorBucketName"},
			"delete-index":                {"IndexArn", "IndexName", "VectorBucketName"},
			"delete-vector-bucket":        {"VectorBucketArn", "VectorBucketName"},
			"delete-vector-bucket-policy": {"VectorBucketArn", "VectorBucketName"},
			"delete-vectors":              {"IndexArn", "IndexName", "Keys", "VectorBucketName"},
			"get-index":                   {"IndexArn", "IndexName", "VectorBucketName"},
			"get-vector-bucket":           {"VectorBucketArn", "VectorBucketName"},
			"get-vector-bucket-policy":    {"VectorBucketArn", "VectorBucketName"},
			"get-vectors":                 {"IndexArn", "IndexName", "Keys", "ReturnData", "ReturnMetadata", "VectorBucketName"},
			"list-indexes":                {"MaxResults", "NextToken", "Prefix", "VectorBucketArn", "VectorBucketName"},
			"list-tags-for-resource":      {"ResourceArn"},
			"list-vector-buckets":         {"MaxResults", "NextToken", "Prefix"},
			"list-vectors":                {"IndexArn", "IndexName", "MaxResults", "NextToken", "ReturnData", "ReturnMetadata", "SegmentCount", "SegmentIndex", "VectorBucketName"},
			"put-vector-bucket-policy":    {"Policy", "VectorBucketArn", "VectorBucketName"},
			"put-vectors":                 {"IndexArn", "IndexName", "VectorBucketName", "Vectors"},
			"query-vectors":               {"Filter", "IndexArn", "IndexName", "QueryVector", "ReturnDistance", "ReturnMetadata", "TopK", "VectorBucketName"},
			"tag-resource":                {"ResourceArn", "Tags"},
			"untag-resource":              {"ResourceArn", "TagKeys"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-index":                {"DataType": "types.DataType", "Dimension": "*int32", "DistanceMetric": "types.DistanceMetric", "EncryptionConfiguration": "*types.EncryptionConfiguration", "IndexName": "*string", "MetadataConfiguration": "*types.MetadataConfiguration", "Tags": "map[string]string", "VectorBucketArn": "*string", "VectorBucketName": "*string"},
			"create-vector-bucket":        {"EncryptionConfiguration": "*types.EncryptionConfiguration", "Tags": "map[string]string", "VectorBucketName": "*string"},
			"delete-index":                {"IndexArn": "*string", "IndexName": "*string", "VectorBucketName": "*string"},
			"delete-vector-bucket":        {"VectorBucketArn": "*string", "VectorBucketName": "*string"},
			"delete-vector-bucket-policy": {"VectorBucketArn": "*string", "VectorBucketName": "*string"},
			"delete-vectors":              {"IndexArn": "*string", "IndexName": "*string", "Keys": "[]string", "VectorBucketName": "*string"},
			"get-index":                   {"IndexArn": "*string", "IndexName": "*string", "VectorBucketName": "*string"},
			"get-vector-bucket":           {"VectorBucketArn": "*string", "VectorBucketName": "*string"},
			"get-vector-bucket-policy":    {"VectorBucketArn": "*string", "VectorBucketName": "*string"},
			"get-vectors":                 {"IndexArn": "*string", "IndexName": "*string", "Keys": "[]string", "ReturnData": "bool", "ReturnMetadata": "bool", "VectorBucketName": "*string"},
			"list-indexes":                {"MaxResults": "*int32", "NextToken": "*string", "Prefix": "*string", "VectorBucketArn": "*string", "VectorBucketName": "*string"},
			"list-tags-for-resource":      {"ResourceArn": "*string"},
			"list-vector-buckets":         {"MaxResults": "*int32", "NextToken": "*string", "Prefix": "*string"},
			"list-vectors":                {"IndexArn": "*string", "IndexName": "*string", "MaxResults": "*int32", "NextToken": "*string", "ReturnData": "bool", "ReturnMetadata": "bool", "SegmentCount": "*int32", "SegmentIndex": "int32", "VectorBucketName": "*string"},
			"put-vector-bucket-policy":    {"Policy": "*string", "VectorBucketArn": "*string", "VectorBucketName": "*string"},
			"put-vectors":                 {"IndexArn": "*string", "IndexName": "*string", "VectorBucketName": "*string", "Vectors": "[]types.PutInputVector"},
			"query-vectors":               {"Filter": "document.Interface", "IndexArn": "*string", "IndexName": "*string", "QueryVector": "types.VectorData", "ReturnDistance": "bool", "ReturnMetadata": "bool", "TopK": "*int32", "VectorBucketName": "*string"},
			"tag-resource":                {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":              {"ResourceArn": "*string", "TagKeys": "[]string"},
		},
		OperationInputRequired: map[string][]string{
			"create-index":                {"DataType", "Dimension", "DistanceMetric", "IndexName"},
			"create-vector-bucket":        {"VectorBucketName"},
			"delete-index":                {},
			"delete-vector-bucket":        {},
			"delete-vector-bucket-policy": {},
			"delete-vectors":              {"Keys"},
			"get-index":                   {},
			"get-vector-bucket":           {},
			"get-vector-bucket-policy":    {},
			"get-vectors":                 {"Keys"},
			"list-indexes":                {},
			"list-tags-for-resource":      {"ResourceArn"},
			"list-vector-buckets":         {},
			"list-vectors":                {},
			"put-vector-bucket-policy":    {"Policy"},
			"put-vectors":                 {"Vectors"},
			"query-vectors":               {"QueryVector", "TopK"},
			"tag-resource":                {"ResourceArn", "Tags"},
			"untag-resource":              {"ResourceArn", "TagKeys"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("s3vectors", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
