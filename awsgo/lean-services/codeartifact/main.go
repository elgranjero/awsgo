package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/codeartifact"
)

var fields_associate_external_connection = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "DomainOwner", Flag: "domain-owner", Type: "*string", Required: false},
	{Name: "ExternalConnection", Flag: "external-connection", Type: "*string", Required: true},
	{Name: "Repository", Flag: "repository", Type: "*string", Required: true},
}

var fields_copy_package_versions = []leanruntime.Field{
	{Name: "AllowOverwrite", Flag: "allow-overwrite", Type: "*bool", Required: false},
	{Name: "DestinationRepository", Flag: "destination-repository", Type: "*string", Required: true},
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "DomainOwner", Flag: "domain-owner", Type: "*string", Required: false},
	{Name: "Format", Flag: "format", Type: "types.PackageFormat", Required: true},
	{Name: "IncludeFromUpstream", Flag: "include-from-upstream", Type: "*bool", Required: false},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: false},
	{Name: "Package", Flag: "package", Type: "*string", Required: true},
	{Name: "SourceRepository", Flag: "source-repository", Type: "*string", Required: true},
	{Name: "VersionRevisions", Flag: "version-revisions", Type: "map[string]string", Required: false},
	{Name: "Versions", Flag: "versions", Type: "[]string", Required: false},
}

var fields_create_domain = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "EncryptionKey", Flag: "encryption-key", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_package_group = []leanruntime.Field{
	{Name: "ContactInfo", Flag: "contact-info", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "DomainOwner", Flag: "domain-owner", Type: "*string", Required: false},
	{Name: "PackageGroup", Flag: "package-group", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_repository = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "DomainOwner", Flag: "domain-owner", Type: "*string", Required: false},
	{Name: "Repository", Flag: "repository", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Upstreams", Flag: "upstreams", Type: "[]types.UpstreamRepository", Required: false},
}

var fields_delete_domain = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "DomainOwner", Flag: "domain-owner", Type: "*string", Required: false},
}

var fields_delete_domain_permissions_policy = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "DomainOwner", Flag: "domain-owner", Type: "*string", Required: false},
	{Name: "PolicyRevision", Flag: "policy-revision", Type: "*string", Required: false},
}

var fields_delete_package = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "DomainOwner", Flag: "domain-owner", Type: "*string", Required: false},
	{Name: "Format", Flag: "format", Type: "types.PackageFormat", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: false},
	{Name: "Package", Flag: "package", Type: "*string", Required: true},
	{Name: "Repository", Flag: "repository", Type: "*string", Required: true},
}

var fields_delete_package_group = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "DomainOwner", Flag: "domain-owner", Type: "*string", Required: false},
	{Name: "PackageGroup", Flag: "package-group", Type: "*string", Required: true},
}

var fields_delete_package_versions = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "DomainOwner", Flag: "domain-owner", Type: "*string", Required: false},
	{Name: "ExpectedStatus", Flag: "expected-status", Type: "types.PackageVersionStatus", Required: false},
	{Name: "Format", Flag: "format", Type: "types.PackageFormat", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: false},
	{Name: "Package", Flag: "package", Type: "*string", Required: true},
	{Name: "Repository", Flag: "repository", Type: "*string", Required: true},
	{Name: "Versions", Flag: "versions", Type: "[]string", Required: true},
}

var fields_delete_repository = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "DomainOwner", Flag: "domain-owner", Type: "*string", Required: false},
	{Name: "Repository", Flag: "repository", Type: "*string", Required: true},
}

var fields_delete_repository_permissions_policy = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "DomainOwner", Flag: "domain-owner", Type: "*string", Required: false},
	{Name: "PolicyRevision", Flag: "policy-revision", Type: "*string", Required: false},
	{Name: "Repository", Flag: "repository", Type: "*string", Required: true},
}

var fields_describe_domain = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "DomainOwner", Flag: "domain-owner", Type: "*string", Required: false},
}

var fields_describe_package = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "DomainOwner", Flag: "domain-owner", Type: "*string", Required: false},
	{Name: "Format", Flag: "format", Type: "types.PackageFormat", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: false},
	{Name: "Package", Flag: "package", Type: "*string", Required: true},
	{Name: "Repository", Flag: "repository", Type: "*string", Required: true},
}

var fields_describe_package_group = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "DomainOwner", Flag: "domain-owner", Type: "*string", Required: false},
	{Name: "PackageGroup", Flag: "package-group", Type: "*string", Required: true},
}

var fields_describe_package_version = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "DomainOwner", Flag: "domain-owner", Type: "*string", Required: false},
	{Name: "Format", Flag: "format", Type: "types.PackageFormat", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: false},
	{Name: "Package", Flag: "package", Type: "*string", Required: true},
	{Name: "PackageVersion", Flag: "package-version", Type: "*string", Required: true},
	{Name: "Repository", Flag: "repository", Type: "*string", Required: true},
}

var fields_describe_repository = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "DomainOwner", Flag: "domain-owner", Type: "*string", Required: false},
	{Name: "Repository", Flag: "repository", Type: "*string", Required: true},
}

var fields_disassociate_external_connection = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "DomainOwner", Flag: "domain-owner", Type: "*string", Required: false},
	{Name: "ExternalConnection", Flag: "external-connection", Type: "*string", Required: true},
	{Name: "Repository", Flag: "repository", Type: "*string", Required: true},
}

var fields_dispose_package_versions = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "DomainOwner", Flag: "domain-owner", Type: "*string", Required: false},
	{Name: "ExpectedStatus", Flag: "expected-status", Type: "types.PackageVersionStatus", Required: false},
	{Name: "Format", Flag: "format", Type: "types.PackageFormat", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: false},
	{Name: "Package", Flag: "package", Type: "*string", Required: true},
	{Name: "Repository", Flag: "repository", Type: "*string", Required: true},
	{Name: "VersionRevisions", Flag: "version-revisions", Type: "map[string]string", Required: false},
	{Name: "Versions", Flag: "versions", Type: "[]string", Required: true},
}

var fields_get_associated_package_group = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "DomainOwner", Flag: "domain-owner", Type: "*string", Required: false},
	{Name: "Format", Flag: "format", Type: "types.PackageFormat", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: false},
	{Name: "Package", Flag: "package", Type: "*string", Required: true},
}

var fields_get_authorization_token = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "DomainOwner", Flag: "domain-owner", Type: "*string", Required: false},
	{Name: "DurationSeconds", Flag: "duration-seconds", Type: "*int64", Required: false},
}

var fields_get_domain_permissions_policy = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "DomainOwner", Flag: "domain-owner", Type: "*string", Required: false},
}

var fields_get_package_version_asset = []leanruntime.Field{
	{Name: "Asset", Flag: "asset", Type: "*string", Required: true},
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "DomainOwner", Flag: "domain-owner", Type: "*string", Required: false},
	{Name: "Format", Flag: "format", Type: "types.PackageFormat", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: false},
	{Name: "Package", Flag: "package", Type: "*string", Required: true},
	{Name: "PackageVersion", Flag: "package-version", Type: "*string", Required: true},
	{Name: "PackageVersionRevision", Flag: "package-version-revision", Type: "*string", Required: false},
	{Name: "Repository", Flag: "repository", Type: "*string", Required: true},
}

var fields_get_package_version_readme = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "DomainOwner", Flag: "domain-owner", Type: "*string", Required: false},
	{Name: "Format", Flag: "format", Type: "types.PackageFormat", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: false},
	{Name: "Package", Flag: "package", Type: "*string", Required: true},
	{Name: "PackageVersion", Flag: "package-version", Type: "*string", Required: true},
	{Name: "Repository", Flag: "repository", Type: "*string", Required: true},
}

var fields_get_repository_endpoint = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "DomainOwner", Flag: "domain-owner", Type: "*string", Required: false},
	{Name: "EndpointType", Flag: "endpoint-type", Type: "types.EndpointType", Required: false},
	{Name: "Format", Flag: "format", Type: "types.PackageFormat", Required: true},
	{Name: "Repository", Flag: "repository", Type: "*string", Required: true},
}

var fields_get_repository_permissions_policy = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "DomainOwner", Flag: "domain-owner", Type: "*string", Required: false},
	{Name: "Repository", Flag: "repository", Type: "*string", Required: true},
}

var fields_list_allowed_repositories_for_group = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "DomainOwner", Flag: "domain-owner", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OriginRestrictionType", Flag: "origin-restriction-type", Type: "types.PackageGroupOriginRestrictionType", Required: true},
	{Name: "PackageGroup", Flag: "package-group", Type: "*string", Required: true},
}

var fields_list_associated_packages = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "DomainOwner", Flag: "domain-owner", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PackageGroup", Flag: "package-group", Type: "*string", Required: true},
	{Name: "Preview", Flag: "preview", Type: "*bool", Required: false},
}

var fields_list_domains = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_package_groups = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "DomainOwner", Flag: "domain-owner", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Prefix", Flag: "prefix", Type: "*string", Required: false},
}

var fields_list_package_version_assets = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "DomainOwner", Flag: "domain-owner", Type: "*string", Required: false},
	{Name: "Format", Flag: "format", Type: "types.PackageFormat", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Package", Flag: "package", Type: "*string", Required: true},
	{Name: "PackageVersion", Flag: "package-version", Type: "*string", Required: true},
	{Name: "Repository", Flag: "repository", Type: "*string", Required: true},
}

var fields_list_package_version_dependencies = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "DomainOwner", Flag: "domain-owner", Type: "*string", Required: false},
	{Name: "Format", Flag: "format", Type: "types.PackageFormat", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Package", Flag: "package", Type: "*string", Required: true},
	{Name: "PackageVersion", Flag: "package-version", Type: "*string", Required: true},
	{Name: "Repository", Flag: "repository", Type: "*string", Required: true},
}

var fields_list_package_versions = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "DomainOwner", Flag: "domain-owner", Type: "*string", Required: false},
	{Name: "Format", Flag: "format", Type: "types.PackageFormat", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OriginType", Flag: "origin-type", Type: "types.PackageVersionOriginType", Required: false},
	{Name: "Package", Flag: "package", Type: "*string", Required: true},
	{Name: "Repository", Flag: "repository", Type: "*string", Required: true},
	{Name: "SortBy", Flag: "sort-by", Type: "types.PackageVersionSortType", Required: false},
	{Name: "Status", Flag: "status", Type: "types.PackageVersionStatus", Required: false},
}

var fields_list_packages = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "DomainOwner", Flag: "domain-owner", Type: "*string", Required: false},
	{Name: "Format", Flag: "format", Type: "types.PackageFormat", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PackagePrefix", Flag: "package-prefix", Type: "*string", Required: false},
	{Name: "Publish", Flag: "publish", Type: "types.AllowPublish", Required: false},
	{Name: "Repository", Flag: "repository", Type: "*string", Required: true},
	{Name: "Upstream", Flag: "upstream", Type: "types.AllowUpstream", Required: false},
}

var fields_list_repositories = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RepositoryPrefix", Flag: "repository-prefix", Type: "*string", Required: false},
}

var fields_list_repositories_in_domain = []leanruntime.Field{
	{Name: "AdministratorAccount", Flag: "administrator-account", Type: "*string", Required: false},
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "DomainOwner", Flag: "domain-owner", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RepositoryPrefix", Flag: "repository-prefix", Type: "*string", Required: false},
}

var fields_list_sub_package_groups = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "DomainOwner", Flag: "domain-owner", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PackageGroup", Flag: "package-group", Type: "*string", Required: true},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_publish_package_version = []leanruntime.Field{
	{Name: "AssetContent", Flag: "asset-content", Type: "io.Reader", Required: true},
	{Name: "AssetName", Flag: "asset-name", Type: "*string", Required: true},
	{Name: "AssetSHA256", Flag: "asset-sha256", Type: "*string", Required: true},
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "DomainOwner", Flag: "domain-owner", Type: "*string", Required: false},
	{Name: "Format", Flag: "format", Type: "types.PackageFormat", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: false},
	{Name: "Package", Flag: "package", Type: "*string", Required: true},
	{Name: "PackageVersion", Flag: "package-version", Type: "*string", Required: true},
	{Name: "Repository", Flag: "repository", Type: "*string", Required: true},
	{Name: "Unfinished", Flag: "unfinished", Type: "*bool", Required: false},
}

var fields_put_domain_permissions_policy = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "DomainOwner", Flag: "domain-owner", Type: "*string", Required: false},
	{Name: "PolicyDocument", Flag: "policy-document", Type: "*string", Required: true},
	{Name: "PolicyRevision", Flag: "policy-revision", Type: "*string", Required: false},
}

var fields_put_package_origin_configuration = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "DomainOwner", Flag: "domain-owner", Type: "*string", Required: false},
	{Name: "Format", Flag: "format", Type: "types.PackageFormat", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: false},
	{Name: "Package", Flag: "package", Type: "*string", Required: true},
	{Name: "Repository", Flag: "repository", Type: "*string", Required: true},
	{Name: "Restrictions", Flag: "restrictions", Type: "*types.PackageOriginRestrictions", Required: true},
}

var fields_put_repository_permissions_policy = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "DomainOwner", Flag: "domain-owner", Type: "*string", Required: false},
	{Name: "PolicyDocument", Flag: "policy-document", Type: "*string", Required: true},
	{Name: "PolicyRevision", Flag: "policy-revision", Type: "*string", Required: false},
	{Name: "Repository", Flag: "repository", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_package_group = []leanruntime.Field{
	{Name: "ContactInfo", Flag: "contact-info", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "DomainOwner", Flag: "domain-owner", Type: "*string", Required: false},
	{Name: "PackageGroup", Flag: "package-group", Type: "*string", Required: true},
}

var fields_update_package_group_origin_configuration = []leanruntime.Field{
	{Name: "AddAllowedRepositories", Flag: "add-allowed-repositories", Type: "[]types.PackageGroupAllowedRepository", Required: false},
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "DomainOwner", Flag: "domain-owner", Type: "*string", Required: false},
	{Name: "PackageGroup", Flag: "package-group", Type: "*string", Required: true},
	{Name: "RemoveAllowedRepositories", Flag: "remove-allowed-repositories", Type: "[]types.PackageGroupAllowedRepository", Required: false},
	{Name: "Restrictions", Flag: "restrictions", Type: "map[string]types.PackageGroupOriginRestrictionMode", Required: false},
}

var fields_update_package_versions_status = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "DomainOwner", Flag: "domain-owner", Type: "*string", Required: false},
	{Name: "ExpectedStatus", Flag: "expected-status", Type: "types.PackageVersionStatus", Required: false},
	{Name: "Format", Flag: "format", Type: "types.PackageFormat", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: false},
	{Name: "Package", Flag: "package", Type: "*string", Required: true},
	{Name: "Repository", Flag: "repository", Type: "*string", Required: true},
	{Name: "TargetStatus", Flag: "target-status", Type: "types.PackageVersionStatus", Required: true},
	{Name: "VersionRevisions", Flag: "version-revisions", Type: "map[string]string", Required: false},
	{Name: "Versions", Flag: "versions", Type: "[]string", Required: true},
}

var fields_update_repository = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "DomainOwner", Flag: "domain-owner", Type: "*string", Required: false},
	{Name: "Repository", Flag: "repository", Type: "*string", Required: true},
	{Name: "Upstreams", Flag: "upstreams", Type: "[]types.UpstreamRepository", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-external-connection": {
			Name:   "associate-external-connection",
			Fields: fields_associate_external_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateExternalConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_external_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateExternalConnection(ctx, input)
			},
		},
		"copy-package-versions": {
			Name:   "copy-package-versions",
			Fields: fields_copy_package_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CopyPackageVersionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_copy_package_versions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CopyPackageVersions(ctx, input)
			},
		},
		"create-domain": {
			Name:   "create-domain",
			Fields: fields_create_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDomain(ctx, input)
			},
		},
		"create-package-group": {
			Name:   "create-package-group",
			Fields: fields_create_package_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePackageGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_package_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePackageGroup(ctx, input)
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
		"delete-domain": {
			Name:   "delete-domain",
			Fields: fields_delete_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDomain(ctx, input)
			},
		},
		"delete-domain-permissions-policy": {
			Name:   "delete-domain-permissions-policy",
			Fields: fields_delete_domain_permissions_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDomainPermissionsPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_domain_permissions_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDomainPermissionsPolicy(ctx, input)
			},
		},
		"delete-package": {
			Name:   "delete-package",
			Fields: fields_delete_package,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePackageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_package, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePackage(ctx, input)
			},
		},
		"delete-package-group": {
			Name:   "delete-package-group",
			Fields: fields_delete_package_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePackageGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_package_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePackageGroup(ctx, input)
			},
		},
		"delete-package-versions": {
			Name:   "delete-package-versions",
			Fields: fields_delete_package_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePackageVersionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_package_versions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePackageVersions(ctx, input)
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
		"delete-repository-permissions-policy": {
			Name:   "delete-repository-permissions-policy",
			Fields: fields_delete_repository_permissions_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRepositoryPermissionsPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_repository_permissions_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRepositoryPermissionsPolicy(ctx, input)
			},
		},
		"describe-domain": {
			Name:   "describe-domain",
			Fields: fields_describe_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDomain(ctx, input)
			},
		},
		"describe-package": {
			Name:   "describe-package",
			Fields: fields_describe_package,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePackageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_package, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribePackage(ctx, input)
			},
		},
		"describe-package-group": {
			Name:   "describe-package-group",
			Fields: fields_describe_package_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePackageGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_package_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribePackageGroup(ctx, input)
			},
		},
		"describe-package-version": {
			Name:   "describe-package-version",
			Fields: fields_describe_package_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePackageVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_package_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribePackageVersion(ctx, input)
			},
		},
		"describe-repository": {
			Name:   "describe-repository",
			Fields: fields_describe_repository,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRepositoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_repository, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeRepository(ctx, input)
			},
		},
		"disassociate-external-connection": {
			Name:   "disassociate-external-connection",
			Fields: fields_disassociate_external_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateExternalConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_external_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateExternalConnection(ctx, input)
			},
		},
		"dispose-package-versions": {
			Name:   "dispose-package-versions",
			Fields: fields_dispose_package_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisposePackageVersionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_dispose_package_versions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisposePackageVersions(ctx, input)
			},
		},
		"get-associated-package-group": {
			Name:   "get-associated-package-group",
			Fields: fields_get_associated_package_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAssociatedPackageGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_associated_package_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAssociatedPackageGroup(ctx, input)
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
		"get-domain-permissions-policy": {
			Name:   "get-domain-permissions-policy",
			Fields: fields_get_domain_permissions_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDomainPermissionsPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_domain_permissions_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDomainPermissionsPolicy(ctx, input)
			},
		},
		"get-package-version-asset": {
			Name:   "get-package-version-asset",
			Fields: fields_get_package_version_asset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPackageVersionAssetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_package_version_asset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPackageVersionAsset(ctx, input)
			},
		},
		"get-package-version-readme": {
			Name:   "get-package-version-readme",
			Fields: fields_get_package_version_readme,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPackageVersionReadmeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_package_version_readme, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPackageVersionReadme(ctx, input)
			},
		},
		"get-repository-endpoint": {
			Name:   "get-repository-endpoint",
			Fields: fields_get_repository_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRepositoryEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_repository_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRepositoryEndpoint(ctx, input)
			},
		},
		"get-repository-permissions-policy": {
			Name:   "get-repository-permissions-policy",
			Fields: fields_get_repository_permissions_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRepositoryPermissionsPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_repository_permissions_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRepositoryPermissionsPolicy(ctx, input)
			},
		},
		"list-allowed-repositories-for-group": {
			Name:   "list-allowed-repositories-for-group",
			Fields: fields_list_allowed_repositories_for_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAllowedRepositoriesForGroupInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_allowed_repositories_for_group, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAllowedRepositoriesForGroup(ctx, input)
				}
				var results []*svc.ListAllowedRepositoriesForGroupOutput
				p := svc.NewListAllowedRepositoriesForGroupPaginator(client, input)
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
		"list-associated-packages": {
			Name:   "list-associated-packages",
			Fields: fields_list_associated_packages,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAssociatedPackagesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_associated_packages, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAssociatedPackages(ctx, input)
				}
				var results []*svc.ListAssociatedPackagesOutput
				p := svc.NewListAssociatedPackagesPaginator(client, input)
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
		"list-domains": {
			Name:   "list-domains",
			Fields: fields_list_domains,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDomainsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_domains, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDomains(ctx, input)
				}
				var results []*svc.ListDomainsOutput
				p := svc.NewListDomainsPaginator(client, input)
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
		"list-package-groups": {
			Name:   "list-package-groups",
			Fields: fields_list_package_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPackageGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_package_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPackageGroups(ctx, input)
				}
				var results []*svc.ListPackageGroupsOutput
				p := svc.NewListPackageGroupsPaginator(client, input)
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
		"list-package-version-assets": {
			Name:   "list-package-version-assets",
			Fields: fields_list_package_version_assets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPackageVersionAssetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_package_version_assets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPackageVersionAssets(ctx, input)
				}
				var results []*svc.ListPackageVersionAssetsOutput
				p := svc.NewListPackageVersionAssetsPaginator(client, input)
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
		"list-package-version-dependencies": {
			Name:   "list-package-version-dependencies",
			Fields: fields_list_package_version_dependencies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPackageVersionDependenciesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_package_version_dependencies, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListPackageVersionDependencies(ctx, input)
			},
		},
		"list-package-versions": {
			Name:   "list-package-versions",
			Fields: fields_list_package_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPackageVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_package_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPackageVersions(ctx, input)
				}
				var results []*svc.ListPackageVersionsOutput
				p := svc.NewListPackageVersionsPaginator(client, input)
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
		"list-packages": {
			Name:   "list-packages",
			Fields: fields_list_packages,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPackagesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_packages, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPackages(ctx, input)
				}
				var results []*svc.ListPackagesOutput
				p := svc.NewListPackagesPaginator(client, input)
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
		"list-repositories": {
			Name:   "list-repositories",
			Fields: fields_list_repositories,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRepositoriesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_repositories, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRepositories(ctx, input)
				}
				var results []*svc.ListRepositoriesOutput
				p := svc.NewListRepositoriesPaginator(client, input)
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
		"list-repositories-in-domain": {
			Name:   "list-repositories-in-domain",
			Fields: fields_list_repositories_in_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRepositoriesInDomainInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_repositories_in_domain, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRepositoriesInDomain(ctx, input)
				}
				var results []*svc.ListRepositoriesInDomainOutput
				p := svc.NewListRepositoriesInDomainPaginator(client, input)
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
		"list-sub-package-groups": {
			Name:   "list-sub-package-groups",
			Fields: fields_list_sub_package_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSubPackageGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_sub_package_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSubPackageGroups(ctx, input)
				}
				var results []*svc.ListSubPackageGroupsOutput
				p := svc.NewListSubPackageGroupsPaginator(client, input)
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
		"publish-package-version": {
			Name:   "publish-package-version",
			Fields: fields_publish_package_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PublishPackageVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_publish_package_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PublishPackageVersion(ctx, input)
			},
		},
		"put-domain-permissions-policy": {
			Name:   "put-domain-permissions-policy",
			Fields: fields_put_domain_permissions_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutDomainPermissionsPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_domain_permissions_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutDomainPermissionsPolicy(ctx, input)
			},
		},
		"put-package-origin-configuration": {
			Name:   "put-package-origin-configuration",
			Fields: fields_put_package_origin_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutPackageOriginConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_package_origin_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutPackageOriginConfiguration(ctx, input)
			},
		},
		"put-repository-permissions-policy": {
			Name:   "put-repository-permissions-policy",
			Fields: fields_put_repository_permissions_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutRepositoryPermissionsPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_repository_permissions_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutRepositoryPermissionsPolicy(ctx, input)
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
		"update-package-group": {
			Name:   "update-package-group",
			Fields: fields_update_package_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePackageGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_package_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePackageGroup(ctx, input)
			},
		},
		"update-package-group-origin-configuration": {
			Name:   "update-package-group-origin-configuration",
			Fields: fields_update_package_group_origin_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePackageGroupOriginConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_package_group_origin_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePackageGroupOriginConfiguration(ctx, input)
			},
		},
		"update-package-versions-status": {
			Name:   "update-package-versions-status",
			Fields: fields_update_package_versions_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePackageVersionsStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_package_versions_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePackageVersionsStatus(ctx, input)
			},
		},
		"update-repository": {
			Name:   "update-repository",
			Fields: fields_update_repository,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRepositoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_repository, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRepository(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("codeartifact", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
