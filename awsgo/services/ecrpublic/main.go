package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/ecrpublic/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"batch-check-layer-availability", "batch-delete-image", "complete-layer-upload", "create-repository", "delete-repository", "delete-repository-policy", "describe-image-tags", "describe-images", "describe-registries", "describe-repositories", "get-authorization-token", "get-registry-catalog-data", "get-repository-catalog-data", "get-repository-policy", "initiate-layer-upload", "list-tags-for-resource", "put-image", "put-registry-catalog-data", "put-repository-catalog-data", "set-repository-policy", "tag-resource", "untag-resource", "upload-layer-part"},
		OperationSet: map[string]bool{"batch-check-layer-availability": true, "batch-delete-image": true, "complete-layer-upload": true, "create-repository": true, "delete-repository": true, "delete-repository-policy": true, "describe-image-tags": true, "describe-images": true, "describe-registries": true, "describe-repositories": true, "get-authorization-token": true, "get-registry-catalog-data": true, "get-repository-catalog-data": true, "get-repository-policy": true, "initiate-layer-upload": true, "list-tags-for-resource": true, "put-image": true, "put-registry-catalog-data": true, "put-repository-catalog-data": true, "set-repository-policy": true, "tag-resource": true, "untag-resource": true, "upload-layer-part": true},
		OperationInputs: map[string][]string{
			"batch-check-layer-availability": {"LayerDigests", "RegistryId", "RepositoryName"},
			"batch-delete-image":             {"ImageIds", "RegistryId", "RepositoryName"},
			"complete-layer-upload":          {"LayerDigests", "RegistryId", "RepositoryName", "UploadId"},
			"create-repository":              {"CatalogData", "RepositoryName", "Tags"},
			"delete-repository":              {"Force", "RegistryId", "RepositoryName"},
			"delete-repository-policy":       {"RegistryId", "RepositoryName"},
			"describe-image-tags":            {"MaxResults", "NextToken", "RegistryId", "RepositoryName"},
			"describe-images":                {"ImageIds", "MaxResults", "NextToken", "RegistryId", "RepositoryName"},
			"describe-registries":            {"MaxResults", "NextToken"},
			"describe-repositories":          {"MaxResults", "NextToken", "RegistryId", "RepositoryNames"},
			"get-authorization-token":        {},
			"get-registry-catalog-data":      {},
			"get-repository-catalog-data":    {"RegistryId", "RepositoryName"},
			"get-repository-policy":          {"RegistryId", "RepositoryName"},
			"initiate-layer-upload":          {"RegistryId", "RepositoryName"},
			"list-tags-for-resource":         {"ResourceArn"},
			"put-image":                      {"ImageDigest", "ImageManifest", "ImageManifestMediaType", "ImageTag", "RegistryId", "RepositoryName"},
			"put-registry-catalog-data":      {"DisplayName"},
			"put-repository-catalog-data":    {"CatalogData", "RegistryId", "RepositoryName"},
			"set-repository-policy":          {"Force", "PolicyText", "RegistryId", "RepositoryName"},
			"tag-resource":                   {"ResourceArn", "Tags"},
			"untag-resource":                 {"ResourceArn", "TagKeys"},
			"upload-layer-part":              {"LayerPartBlob", "PartFirstByte", "PartLastByte", "RegistryId", "RepositoryName", "UploadId"},
		},
		OperationInputTypes: map[string]map[string]string{
			"batch-check-layer-availability": {"LayerDigests": "[]string", "RegistryId": "*string", "RepositoryName": "*string"},
			"batch-delete-image":             {"ImageIds": "[]types.ImageIdentifier", "RegistryId": "*string", "RepositoryName": "*string"},
			"complete-layer-upload":          {"LayerDigests": "[]string", "RegistryId": "*string", "RepositoryName": "*string", "UploadId": "*string"},
			"create-repository":              {"CatalogData": "*types.RepositoryCatalogDataInput", "RepositoryName": "*string", "Tags": "[]types.Tag"},
			"delete-repository":              {"Force": "bool", "RegistryId": "*string", "RepositoryName": "*string"},
			"delete-repository-policy":       {"RegistryId": "*string", "RepositoryName": "*string"},
			"describe-image-tags":            {"MaxResults": "*int32", "NextToken": "*string", "RegistryId": "*string", "RepositoryName": "*string"},
			"describe-images":                {"ImageIds": "[]types.ImageIdentifier", "MaxResults": "*int32", "NextToken": "*string", "RegistryId": "*string", "RepositoryName": "*string"},
			"describe-registries":            {"MaxResults": "*int32", "NextToken": "*string"},
			"describe-repositories":          {"MaxResults": "*int32", "NextToken": "*string", "RegistryId": "*string", "RepositoryNames": "[]string"},
			"get-authorization-token":        {},
			"get-registry-catalog-data":      {},
			"get-repository-catalog-data":    {"RegistryId": "*string", "RepositoryName": "*string"},
			"get-repository-policy":          {"RegistryId": "*string", "RepositoryName": "*string"},
			"initiate-layer-upload":          {"RegistryId": "*string", "RepositoryName": "*string"},
			"list-tags-for-resource":         {"ResourceArn": "*string"},
			"put-image":                      {"ImageDigest": "*string", "ImageManifest": "*string", "ImageManifestMediaType": "*string", "ImageTag": "*string", "RegistryId": "*string", "RepositoryName": "*string"},
			"put-registry-catalog-data":      {"DisplayName": "*string"},
			"put-repository-catalog-data":    {"CatalogData": "*types.RepositoryCatalogDataInput", "RegistryId": "*string", "RepositoryName": "*string"},
			"set-repository-policy":          {"Force": "bool", "PolicyText": "*string", "RegistryId": "*string", "RepositoryName": "*string"},
			"tag-resource":                   {"ResourceArn": "*string", "Tags": "[]types.Tag"},
			"untag-resource":                 {"ResourceArn": "*string", "TagKeys": "[]string"},
			"upload-layer-part":              {"LayerPartBlob": "[]byte", "PartFirstByte": "*int64", "PartLastByte": "*int64", "RegistryId": "*string", "RepositoryName": "*string", "UploadId": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"batch-check-layer-availability": {"LayerDigests", "RepositoryName"},
			"batch-delete-image":             {"ImageIds", "RepositoryName"},
			"complete-layer-upload":          {"LayerDigests", "RepositoryName", "UploadId"},
			"create-repository":              {"RepositoryName"},
			"delete-repository":              {"RepositoryName"},
			"delete-repository-policy":       {"RepositoryName"},
			"describe-image-tags":            {"RepositoryName"},
			"describe-images":                {"RepositoryName"},
			"describe-registries":            {},
			"describe-repositories":          {},
			"get-authorization-token":        {},
			"get-registry-catalog-data":      {},
			"get-repository-catalog-data":    {"RepositoryName"},
			"get-repository-policy":          {"RepositoryName"},
			"initiate-layer-upload":          {"RepositoryName"},
			"list-tags-for-resource":         {"ResourceArn"},
			"put-image":                      {"ImageManifest", "RepositoryName"},
			"put-registry-catalog-data":      {},
			"put-repository-catalog-data":    {"CatalogData", "RepositoryName"},
			"set-repository-policy":          {"PolicyText", "RepositoryName"},
			"tag-resource":                   {"ResourceArn", "Tags"},
			"untag-resource":                 {"ResourceArn", "TagKeys"},
			"upload-layer-part":              {"LayerPartBlob", "PartFirstByte", "PartLastByte", "RepositoryName", "UploadId"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("ecrpublic", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
