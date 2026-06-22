package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/clouddirectory"
)

var fields_add_facet_to_object = []leanruntime.Field{
	{Name: "DirectoryArn", Flag: "directory-arn", Type: "*string", Required: true},
	{Name: "ObjectAttributeList", Flag: "object-attribute-list", Type: "[]types.AttributeKeyAndValue", Required: false},
	{Name: "ObjectReference", Flag: "object-reference", Type: "*types.ObjectReference", Required: true},
	{Name: "SchemaFacet", Flag: "schema-facet", Type: "*types.SchemaFacet", Required: true},
}

var fields_apply_schema = []leanruntime.Field{
	{Name: "DirectoryArn", Flag: "directory-arn", Type: "*string", Required: true},
	{Name: "PublishedSchemaArn", Flag: "published-schema-arn", Type: "*string", Required: true},
}

var fields_attach_object = []leanruntime.Field{
	{Name: "ChildReference", Flag: "child-reference", Type: "*types.ObjectReference", Required: true},
	{Name: "DirectoryArn", Flag: "directory-arn", Type: "*string", Required: true},
	{Name: "LinkName", Flag: "link-name", Type: "*string", Required: true},
	{Name: "ParentReference", Flag: "parent-reference", Type: "*types.ObjectReference", Required: true},
}

var fields_attach_policy = []leanruntime.Field{
	{Name: "DirectoryArn", Flag: "directory-arn", Type: "*string", Required: true},
	{Name: "ObjectReference", Flag: "object-reference", Type: "*types.ObjectReference", Required: true},
	{Name: "PolicyReference", Flag: "policy-reference", Type: "*types.ObjectReference", Required: true},
}

var fields_attach_to_index = []leanruntime.Field{
	{Name: "DirectoryArn", Flag: "directory-arn", Type: "*string", Required: true},
	{Name: "IndexReference", Flag: "index-reference", Type: "*types.ObjectReference", Required: true},
	{Name: "TargetReference", Flag: "target-reference", Type: "*types.ObjectReference", Required: true},
}

var fields_attach_typed_link = []leanruntime.Field{
	{Name: "Attributes", Flag: "attributes", Type: "[]types.AttributeNameAndValue", Required: true},
	{Name: "DirectoryArn", Flag: "directory-arn", Type: "*string", Required: true},
	{Name: "SourceObjectReference", Flag: "source-object-reference", Type: "*types.ObjectReference", Required: true},
	{Name: "TargetObjectReference", Flag: "target-object-reference", Type: "*types.ObjectReference", Required: true},
	{Name: "TypedLinkFacet", Flag: "typed-link-facet", Type: "*types.TypedLinkSchemaAndFacetName", Required: true},
}

var fields_batch_read = []leanruntime.Field{
	{Name: "ConsistencyLevel", Flag: "consistency-level", Type: "types.ConsistencyLevel", Required: false},
	{Name: "DirectoryArn", Flag: "directory-arn", Type: "*string", Required: true},
	{Name: "Operations", Flag: "operations", Type: "[]types.BatchReadOperation", Required: true},
}

var fields_batch_write = []leanruntime.Field{
	{Name: "DirectoryArn", Flag: "directory-arn", Type: "*string", Required: true},
	{Name: "Operations", Flag: "operations", Type: "[]types.BatchWriteOperation", Required: true},
}

var fields_create_directory = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "SchemaArn", Flag: "schema-arn", Type: "*string", Required: true},
}

var fields_create_facet = []leanruntime.Field{
	{Name: "Attributes", Flag: "attributes", Type: "[]types.FacetAttribute", Required: false},
	{Name: "FacetStyle", Flag: "facet-style", Type: "types.FacetStyle", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ObjectType", Flag: "object-type", Type: "types.ObjectType", Required: false},
	{Name: "SchemaArn", Flag: "schema-arn", Type: "*string", Required: true},
}

var fields_create_index = []leanruntime.Field{
	{Name: "DirectoryArn", Flag: "directory-arn", Type: "*string", Required: true},
	{Name: "IsUnique", Flag: "is-unique", Type: "bool", Required: true},
	{Name: "LinkName", Flag: "link-name", Type: "*string", Required: false},
	{Name: "OrderedIndexedAttributeList", Flag: "ordered-indexed-attribute-list", Type: "[]types.AttributeKey", Required: true},
	{Name: "ParentReference", Flag: "parent-reference", Type: "*types.ObjectReference", Required: false},
}

var fields_create_object = []leanruntime.Field{
	{Name: "DirectoryArn", Flag: "directory-arn", Type: "*string", Required: true},
	{Name: "LinkName", Flag: "link-name", Type: "*string", Required: false},
	{Name: "ObjectAttributeList", Flag: "object-attribute-list", Type: "[]types.AttributeKeyAndValue", Required: false},
	{Name: "ParentReference", Flag: "parent-reference", Type: "*types.ObjectReference", Required: false},
	{Name: "SchemaFacets", Flag: "schema-facets", Type: "[]types.SchemaFacet", Required: true},
}

var fields_create_schema = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_create_typed_link_facet = []leanruntime.Field{
	{Name: "Facet", Flag: "facet", Type: "*types.TypedLinkFacet", Required: true},
	{Name: "SchemaArn", Flag: "schema-arn", Type: "*string", Required: true},
}

var fields_delete_directory = []leanruntime.Field{
	{Name: "DirectoryArn", Flag: "directory-arn", Type: "*string", Required: true},
}

var fields_delete_facet = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "SchemaArn", Flag: "schema-arn", Type: "*string", Required: true},
}

var fields_delete_object = []leanruntime.Field{
	{Name: "DirectoryArn", Flag: "directory-arn", Type: "*string", Required: true},
	{Name: "ObjectReference", Flag: "object-reference", Type: "*types.ObjectReference", Required: true},
}

var fields_delete_schema = []leanruntime.Field{
	{Name: "SchemaArn", Flag: "schema-arn", Type: "*string", Required: true},
}

var fields_delete_typed_link_facet = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "SchemaArn", Flag: "schema-arn", Type: "*string", Required: true},
}

var fields_detach_from_index = []leanruntime.Field{
	{Name: "DirectoryArn", Flag: "directory-arn", Type: "*string", Required: true},
	{Name: "IndexReference", Flag: "index-reference", Type: "*types.ObjectReference", Required: true},
	{Name: "TargetReference", Flag: "target-reference", Type: "*types.ObjectReference", Required: true},
}

var fields_detach_object = []leanruntime.Field{
	{Name: "DirectoryArn", Flag: "directory-arn", Type: "*string", Required: true},
	{Name: "LinkName", Flag: "link-name", Type: "*string", Required: true},
	{Name: "ParentReference", Flag: "parent-reference", Type: "*types.ObjectReference", Required: true},
}

var fields_detach_policy = []leanruntime.Field{
	{Name: "DirectoryArn", Flag: "directory-arn", Type: "*string", Required: true},
	{Name: "ObjectReference", Flag: "object-reference", Type: "*types.ObjectReference", Required: true},
	{Name: "PolicyReference", Flag: "policy-reference", Type: "*types.ObjectReference", Required: true},
}

var fields_detach_typed_link = []leanruntime.Field{
	{Name: "DirectoryArn", Flag: "directory-arn", Type: "*string", Required: true},
	{Name: "TypedLinkSpecifier", Flag: "typed-link-specifier", Type: "*types.TypedLinkSpecifier", Required: true},
}

var fields_disable_directory = []leanruntime.Field{
	{Name: "DirectoryArn", Flag: "directory-arn", Type: "*string", Required: true},
}

var fields_enable_directory = []leanruntime.Field{
	{Name: "DirectoryArn", Flag: "directory-arn", Type: "*string", Required: true},
}

var fields_get_applied_schema_version = []leanruntime.Field{
	{Name: "SchemaArn", Flag: "schema-arn", Type: "*string", Required: true},
}

var fields_get_directory = []leanruntime.Field{
	{Name: "DirectoryArn", Flag: "directory-arn", Type: "*string", Required: true},
}

var fields_get_facet = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "SchemaArn", Flag: "schema-arn", Type: "*string", Required: true},
}

var fields_get_link_attributes = []leanruntime.Field{
	{Name: "AttributeNames", Flag: "attribute-names", Type: "[]string", Required: true},
	{Name: "ConsistencyLevel", Flag: "consistency-level", Type: "types.ConsistencyLevel", Required: false},
	{Name: "DirectoryArn", Flag: "directory-arn", Type: "*string", Required: true},
	{Name: "TypedLinkSpecifier", Flag: "typed-link-specifier", Type: "*types.TypedLinkSpecifier", Required: true},
}

var fields_get_object_attributes = []leanruntime.Field{
	{Name: "AttributeNames", Flag: "attribute-names", Type: "[]string", Required: true},
	{Name: "ConsistencyLevel", Flag: "consistency-level", Type: "types.ConsistencyLevel", Required: false},
	{Name: "DirectoryArn", Flag: "directory-arn", Type: "*string", Required: true},
	{Name: "ObjectReference", Flag: "object-reference", Type: "*types.ObjectReference", Required: true},
	{Name: "SchemaFacet", Flag: "schema-facet", Type: "*types.SchemaFacet", Required: true},
}

var fields_get_object_information = []leanruntime.Field{
	{Name: "ConsistencyLevel", Flag: "consistency-level", Type: "types.ConsistencyLevel", Required: false},
	{Name: "DirectoryArn", Flag: "directory-arn", Type: "*string", Required: true},
	{Name: "ObjectReference", Flag: "object-reference", Type: "*types.ObjectReference", Required: true},
}

var fields_get_schema_as_json = []leanruntime.Field{
	{Name: "SchemaArn", Flag: "schema-arn", Type: "*string", Required: true},
}

var fields_get_typed_link_facet_information = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "SchemaArn", Flag: "schema-arn", Type: "*string", Required: true},
}

var fields_list_applied_schema_arns = []leanruntime.Field{
	{Name: "DirectoryArn", Flag: "directory-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SchemaArn", Flag: "schema-arn", Type: "*string", Required: false},
}

var fields_list_attached_indices = []leanruntime.Field{
	{Name: "ConsistencyLevel", Flag: "consistency-level", Type: "types.ConsistencyLevel", Required: false},
	{Name: "DirectoryArn", Flag: "directory-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TargetReference", Flag: "target-reference", Type: "*types.ObjectReference", Required: true},
}

var fields_list_development_schema_arns = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_directories = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "State", Flag: "state", Type: "types.DirectoryState", Required: false},
}

var fields_list_facet_attributes = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SchemaArn", Flag: "schema-arn", Type: "*string", Required: true},
}

var fields_list_facet_names = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SchemaArn", Flag: "schema-arn", Type: "*string", Required: true},
}

var fields_list_incoming_typed_links = []leanruntime.Field{
	{Name: "ConsistencyLevel", Flag: "consistency-level", Type: "types.ConsistencyLevel", Required: false},
	{Name: "DirectoryArn", Flag: "directory-arn", Type: "*string", Required: true},
	{Name: "FilterAttributeRanges", Flag: "filter-attribute-ranges", Type: "[]types.TypedLinkAttributeRange", Required: false},
	{Name: "FilterTypedLink", Flag: "filter-typed-link", Type: "*types.TypedLinkSchemaAndFacetName", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ObjectReference", Flag: "object-reference", Type: "*types.ObjectReference", Required: true},
}

var fields_list_index = []leanruntime.Field{
	{Name: "ConsistencyLevel", Flag: "consistency-level", Type: "types.ConsistencyLevel", Required: false},
	{Name: "DirectoryArn", Flag: "directory-arn", Type: "*string", Required: true},
	{Name: "IndexReference", Flag: "index-reference", Type: "*types.ObjectReference", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RangesOnIndexedValues", Flag: "ranges-on-indexed-values", Type: "[]types.ObjectAttributeRange", Required: false},
}

var fields_list_managed_schema_arns = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SchemaArn", Flag: "schema-arn", Type: "*string", Required: false},
}

var fields_list_object_attributes = []leanruntime.Field{
	{Name: "ConsistencyLevel", Flag: "consistency-level", Type: "types.ConsistencyLevel", Required: false},
	{Name: "DirectoryArn", Flag: "directory-arn", Type: "*string", Required: true},
	{Name: "FacetFilter", Flag: "facet-filter", Type: "*types.SchemaFacet", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ObjectReference", Flag: "object-reference", Type: "*types.ObjectReference", Required: true},
}

var fields_list_object_children = []leanruntime.Field{
	{Name: "ConsistencyLevel", Flag: "consistency-level", Type: "types.ConsistencyLevel", Required: false},
	{Name: "DirectoryArn", Flag: "directory-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ObjectReference", Flag: "object-reference", Type: "*types.ObjectReference", Required: true},
}

var fields_list_object_parent_paths = []leanruntime.Field{
	{Name: "DirectoryArn", Flag: "directory-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ObjectReference", Flag: "object-reference", Type: "*types.ObjectReference", Required: true},
}

var fields_list_object_parents = []leanruntime.Field{
	{Name: "ConsistencyLevel", Flag: "consistency-level", Type: "types.ConsistencyLevel", Required: false},
	{Name: "DirectoryArn", Flag: "directory-arn", Type: "*string", Required: true},
	{Name: "IncludeAllLinksToEachParent", Flag: "include-all-links-to-each-parent", Type: "bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ObjectReference", Flag: "object-reference", Type: "*types.ObjectReference", Required: true},
}

var fields_list_object_policies = []leanruntime.Field{
	{Name: "ConsistencyLevel", Flag: "consistency-level", Type: "types.ConsistencyLevel", Required: false},
	{Name: "DirectoryArn", Flag: "directory-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ObjectReference", Flag: "object-reference", Type: "*types.ObjectReference", Required: true},
}

var fields_list_outgoing_typed_links = []leanruntime.Field{
	{Name: "ConsistencyLevel", Flag: "consistency-level", Type: "types.ConsistencyLevel", Required: false},
	{Name: "DirectoryArn", Flag: "directory-arn", Type: "*string", Required: true},
	{Name: "FilterAttributeRanges", Flag: "filter-attribute-ranges", Type: "[]types.TypedLinkAttributeRange", Required: false},
	{Name: "FilterTypedLink", Flag: "filter-typed-link", Type: "*types.TypedLinkSchemaAndFacetName", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ObjectReference", Flag: "object-reference", Type: "*types.ObjectReference", Required: true},
}

var fields_list_policy_attachments = []leanruntime.Field{
	{Name: "ConsistencyLevel", Flag: "consistency-level", Type: "types.ConsistencyLevel", Required: false},
	{Name: "DirectoryArn", Flag: "directory-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PolicyReference", Flag: "policy-reference", Type: "*types.ObjectReference", Required: true},
}

var fields_list_published_schema_arns = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SchemaArn", Flag: "schema-arn", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_typed_link_facet_attributes = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SchemaArn", Flag: "schema-arn", Type: "*string", Required: true},
}

var fields_list_typed_link_facet_names = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SchemaArn", Flag: "schema-arn", Type: "*string", Required: true},
}

var fields_lookup_policy = []leanruntime.Field{
	{Name: "DirectoryArn", Flag: "directory-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ObjectReference", Flag: "object-reference", Type: "*types.ObjectReference", Required: true},
}

var fields_publish_schema = []leanruntime.Field{
	{Name: "DevelopmentSchemaArn", Flag: "development-schema-arn", Type: "*string", Required: true},
	{Name: "MinorVersion", Flag: "minor-version", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Version", Flag: "version", Type: "*string", Required: true},
}

var fields_put_schema_from_json = []leanruntime.Field{
	{Name: "Document", Flag: "document", Type: "*string", Required: true},
	{Name: "SchemaArn", Flag: "schema-arn", Type: "*string", Required: true},
}

var fields_remove_facet_from_object = []leanruntime.Field{
	{Name: "DirectoryArn", Flag: "directory-arn", Type: "*string", Required: true},
	{Name: "ObjectReference", Flag: "object-reference", Type: "*types.ObjectReference", Required: true},
	{Name: "SchemaFacet", Flag: "schema-facet", Type: "*types.SchemaFacet", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_facet = []leanruntime.Field{
	{Name: "AttributeUpdates", Flag: "attribute-updates", Type: "[]types.FacetAttributeUpdate", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ObjectType", Flag: "object-type", Type: "types.ObjectType", Required: false},
	{Name: "SchemaArn", Flag: "schema-arn", Type: "*string", Required: true},
}

var fields_update_link_attributes = []leanruntime.Field{
	{Name: "AttributeUpdates", Flag: "attribute-updates", Type: "[]types.LinkAttributeUpdate", Required: true},
	{Name: "DirectoryArn", Flag: "directory-arn", Type: "*string", Required: true},
	{Name: "TypedLinkSpecifier", Flag: "typed-link-specifier", Type: "*types.TypedLinkSpecifier", Required: true},
}

var fields_update_object_attributes = []leanruntime.Field{
	{Name: "AttributeUpdates", Flag: "attribute-updates", Type: "[]types.ObjectAttributeUpdate", Required: true},
	{Name: "DirectoryArn", Flag: "directory-arn", Type: "*string", Required: true},
	{Name: "ObjectReference", Flag: "object-reference", Type: "*types.ObjectReference", Required: true},
}

var fields_update_schema = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "SchemaArn", Flag: "schema-arn", Type: "*string", Required: true},
}

var fields_update_typed_link_facet = []leanruntime.Field{
	{Name: "AttributeUpdates", Flag: "attribute-updates", Type: "[]types.TypedLinkFacetAttributeUpdate", Required: true},
	{Name: "IdentityAttributeOrder", Flag: "identity-attribute-order", Type: "[]string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "SchemaArn", Flag: "schema-arn", Type: "*string", Required: true},
}

var fields_upgrade_applied_schema = []leanruntime.Field{
	{Name: "DirectoryArn", Flag: "directory-arn", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "bool", Required: false},
	{Name: "PublishedSchemaArn", Flag: "published-schema-arn", Type: "*string", Required: true},
}

var fields_upgrade_published_schema = []leanruntime.Field{
	{Name: "DevelopmentSchemaArn", Flag: "development-schema-arn", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "bool", Required: false},
	{Name: "MinorVersion", Flag: "minor-version", Type: "*string", Required: true},
	{Name: "PublishedSchemaArn", Flag: "published-schema-arn", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"add-facet-to-object": {
			Name:   "add-facet-to-object",
			Fields: fields_add_facet_to_object,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddFacetToObjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_facet_to_object, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddFacetToObject(ctx, input)
			},
		},
		"apply-schema": {
			Name:   "apply-schema",
			Fields: fields_apply_schema,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ApplySchemaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_apply_schema, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ApplySchema(ctx, input)
			},
		},
		"attach-object": {
			Name:   "attach-object",
			Fields: fields_attach_object,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AttachObjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_attach_object, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AttachObject(ctx, input)
			},
		},
		"attach-policy": {
			Name:   "attach-policy",
			Fields: fields_attach_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AttachPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_attach_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AttachPolicy(ctx, input)
			},
		},
		"attach-to-index": {
			Name:   "attach-to-index",
			Fields: fields_attach_to_index,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AttachToIndexInput{}
				if _, err := leanruntime.ApplyInput(input, fields_attach_to_index, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AttachToIndex(ctx, input)
			},
		},
		"attach-typed-link": {
			Name:   "attach-typed-link",
			Fields: fields_attach_typed_link,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AttachTypedLinkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_attach_typed_link, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AttachTypedLink(ctx, input)
			},
		},
		"batch-read": {
			Name:   "batch-read",
			Fields: fields_batch_read,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchReadInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_read, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchRead(ctx, input)
			},
		},
		"batch-write": {
			Name:   "batch-write",
			Fields: fields_batch_write,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchWriteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_write, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchWrite(ctx, input)
			},
		},
		"create-directory": {
			Name:   "create-directory",
			Fields: fields_create_directory,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDirectoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_directory, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDirectory(ctx, input)
			},
		},
		"create-facet": {
			Name:   "create-facet",
			Fields: fields_create_facet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFacetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_facet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFacet(ctx, input)
			},
		},
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
		"create-object": {
			Name:   "create-object",
			Fields: fields_create_object,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateObjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_object, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateObject(ctx, input)
			},
		},
		"create-schema": {
			Name:   "create-schema",
			Fields: fields_create_schema,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSchemaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_schema, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSchema(ctx, input)
			},
		},
		"create-typed-link-facet": {
			Name:   "create-typed-link-facet",
			Fields: fields_create_typed_link_facet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTypedLinkFacetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_typed_link_facet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTypedLinkFacet(ctx, input)
			},
		},
		"delete-directory": {
			Name:   "delete-directory",
			Fields: fields_delete_directory,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDirectoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_directory, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDirectory(ctx, input)
			},
		},
		"delete-facet": {
			Name:   "delete-facet",
			Fields: fields_delete_facet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFacetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_facet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFacet(ctx, input)
			},
		},
		"delete-object": {
			Name:   "delete-object",
			Fields: fields_delete_object,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteObjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_object, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteObject(ctx, input)
			},
		},
		"delete-schema": {
			Name:   "delete-schema",
			Fields: fields_delete_schema,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSchemaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_schema, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSchema(ctx, input)
			},
		},
		"delete-typed-link-facet": {
			Name:   "delete-typed-link-facet",
			Fields: fields_delete_typed_link_facet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTypedLinkFacetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_typed_link_facet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTypedLinkFacet(ctx, input)
			},
		},
		"detach-from-index": {
			Name:   "detach-from-index",
			Fields: fields_detach_from_index,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetachFromIndexInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detach_from_index, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetachFromIndex(ctx, input)
			},
		},
		"detach-object": {
			Name:   "detach-object",
			Fields: fields_detach_object,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetachObjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detach_object, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetachObject(ctx, input)
			},
		},
		"detach-policy": {
			Name:   "detach-policy",
			Fields: fields_detach_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetachPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detach_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetachPolicy(ctx, input)
			},
		},
		"detach-typed-link": {
			Name:   "detach-typed-link",
			Fields: fields_detach_typed_link,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetachTypedLinkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detach_typed_link, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetachTypedLink(ctx, input)
			},
		},
		"disable-directory": {
			Name:   "disable-directory",
			Fields: fields_disable_directory,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableDirectoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_directory, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableDirectory(ctx, input)
			},
		},
		"enable-directory": {
			Name:   "enable-directory",
			Fields: fields_enable_directory,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableDirectoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_directory, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableDirectory(ctx, input)
			},
		},
		"get-applied-schema-version": {
			Name:   "get-applied-schema-version",
			Fields: fields_get_applied_schema_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAppliedSchemaVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_applied_schema_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAppliedSchemaVersion(ctx, input)
			},
		},
		"get-directory": {
			Name:   "get-directory",
			Fields: fields_get_directory,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDirectoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_directory, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDirectory(ctx, input)
			},
		},
		"get-facet": {
			Name:   "get-facet",
			Fields: fields_get_facet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFacetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_facet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFacet(ctx, input)
			},
		},
		"get-link-attributes": {
			Name:   "get-link-attributes",
			Fields: fields_get_link_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLinkAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_link_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLinkAttributes(ctx, input)
			},
		},
		"get-object-attributes": {
			Name:   "get-object-attributes",
			Fields: fields_get_object_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetObjectAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_object_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetObjectAttributes(ctx, input)
			},
		},
		"get-object-information": {
			Name:   "get-object-information",
			Fields: fields_get_object_information,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetObjectInformationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_object_information, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetObjectInformation(ctx, input)
			},
		},
		"get-schema-as-json": {
			Name:   "get-schema-as-json",
			Fields: fields_get_schema_as_json,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSchemaAsJsonInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_schema_as_json, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSchemaAsJson(ctx, input)
			},
		},
		"get-typed-link-facet-information": {
			Name:   "get-typed-link-facet-information",
			Fields: fields_get_typed_link_facet_information,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTypedLinkFacetInformationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_typed_link_facet_information, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTypedLinkFacetInformation(ctx, input)
			},
		},
		"list-applied-schema-arns": {
			Name:   "list-applied-schema-arns",
			Fields: fields_list_applied_schema_arns,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAppliedSchemaArnsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_applied_schema_arns, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAppliedSchemaArns(ctx, input)
				}
				var results []*svc.ListAppliedSchemaArnsOutput
				p := svc.NewListAppliedSchemaArnsPaginator(client, input)
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
		"list-attached-indices": {
			Name:   "list-attached-indices",
			Fields: fields_list_attached_indices,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAttachedIndicesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_attached_indices, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAttachedIndices(ctx, input)
				}
				var results []*svc.ListAttachedIndicesOutput
				p := svc.NewListAttachedIndicesPaginator(client, input)
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
		"list-development-schema-arns": {
			Name:   "list-development-schema-arns",
			Fields: fields_list_development_schema_arns,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDevelopmentSchemaArnsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_development_schema_arns, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDevelopmentSchemaArns(ctx, input)
				}
				var results []*svc.ListDevelopmentSchemaArnsOutput
				p := svc.NewListDevelopmentSchemaArnsPaginator(client, input)
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
		"list-directories": {
			Name:   "list-directories",
			Fields: fields_list_directories,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDirectoriesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_directories, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDirectories(ctx, input)
				}
				var results []*svc.ListDirectoriesOutput
				p := svc.NewListDirectoriesPaginator(client, input)
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
		"list-facet-attributes": {
			Name:   "list-facet-attributes",
			Fields: fields_list_facet_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFacetAttributesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_facet_attributes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFacetAttributes(ctx, input)
				}
				var results []*svc.ListFacetAttributesOutput
				p := svc.NewListFacetAttributesPaginator(client, input)
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
		"list-facet-names": {
			Name:   "list-facet-names",
			Fields: fields_list_facet_names,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFacetNamesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_facet_names, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFacetNames(ctx, input)
				}
				var results []*svc.ListFacetNamesOutput
				p := svc.NewListFacetNamesPaginator(client, input)
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
		"list-incoming-typed-links": {
			Name:   "list-incoming-typed-links",
			Fields: fields_list_incoming_typed_links,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIncomingTypedLinksInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_incoming_typed_links, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListIncomingTypedLinks(ctx, input)
			},
		},
		"list-index": {
			Name:   "list-index",
			Fields: fields_list_index,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIndexInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_index, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListIndex(ctx, input)
				}
				var results []*svc.ListIndexOutput
				p := svc.NewListIndexPaginator(client, input)
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
		"list-managed-schema-arns": {
			Name:   "list-managed-schema-arns",
			Fields: fields_list_managed_schema_arns,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListManagedSchemaArnsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_managed_schema_arns, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListManagedSchemaArns(ctx, input)
				}
				var results []*svc.ListManagedSchemaArnsOutput
				p := svc.NewListManagedSchemaArnsPaginator(client, input)
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
		"list-object-attributes": {
			Name:   "list-object-attributes",
			Fields: fields_list_object_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListObjectAttributesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_object_attributes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListObjectAttributes(ctx, input)
				}
				var results []*svc.ListObjectAttributesOutput
				p := svc.NewListObjectAttributesPaginator(client, input)
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
		"list-object-children": {
			Name:   "list-object-children",
			Fields: fields_list_object_children,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListObjectChildrenInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_object_children, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListObjectChildren(ctx, input)
				}
				var results []*svc.ListObjectChildrenOutput
				p := svc.NewListObjectChildrenPaginator(client, input)
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
		"list-object-parent-paths": {
			Name:   "list-object-parent-paths",
			Fields: fields_list_object_parent_paths,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListObjectParentPathsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_object_parent_paths, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListObjectParentPaths(ctx, input)
				}
				var results []*svc.ListObjectParentPathsOutput
				p := svc.NewListObjectParentPathsPaginator(client, input)
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
		"list-object-parents": {
			Name:   "list-object-parents",
			Fields: fields_list_object_parents,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListObjectParentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_object_parents, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListObjectParents(ctx, input)
				}
				var results []*svc.ListObjectParentsOutput
				p := svc.NewListObjectParentsPaginator(client, input)
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
		"list-object-policies": {
			Name:   "list-object-policies",
			Fields: fields_list_object_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListObjectPoliciesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_object_policies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListObjectPolicies(ctx, input)
				}
				var results []*svc.ListObjectPoliciesOutput
				p := svc.NewListObjectPoliciesPaginator(client, input)
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
		"list-outgoing-typed-links": {
			Name:   "list-outgoing-typed-links",
			Fields: fields_list_outgoing_typed_links,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOutgoingTypedLinksInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_outgoing_typed_links, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListOutgoingTypedLinks(ctx, input)
			},
		},
		"list-policy-attachments": {
			Name:   "list-policy-attachments",
			Fields: fields_list_policy_attachments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPolicyAttachmentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_policy_attachments, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPolicyAttachments(ctx, input)
				}
				var results []*svc.ListPolicyAttachmentsOutput
				p := svc.NewListPolicyAttachmentsPaginator(client, input)
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
		"list-published-schema-arns": {
			Name:   "list-published-schema-arns",
			Fields: fields_list_published_schema_arns,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPublishedSchemaArnsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_published_schema_arns, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPublishedSchemaArns(ctx, input)
				}
				var results []*svc.ListPublishedSchemaArnsOutput
				p := svc.NewListPublishedSchemaArnsPaginator(client, input)
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
		"list-typed-link-facet-attributes": {
			Name:   "list-typed-link-facet-attributes",
			Fields: fields_list_typed_link_facet_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTypedLinkFacetAttributesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_typed_link_facet_attributes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTypedLinkFacetAttributes(ctx, input)
				}
				var results []*svc.ListTypedLinkFacetAttributesOutput
				p := svc.NewListTypedLinkFacetAttributesPaginator(client, input)
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
		"list-typed-link-facet-names": {
			Name:   "list-typed-link-facet-names",
			Fields: fields_list_typed_link_facet_names,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTypedLinkFacetNamesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_typed_link_facet_names, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTypedLinkFacetNames(ctx, input)
				}
				var results []*svc.ListTypedLinkFacetNamesOutput
				p := svc.NewListTypedLinkFacetNamesPaginator(client, input)
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
		"lookup-policy": {
			Name:   "lookup-policy",
			Fields: fields_lookup_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.LookupPolicyInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_lookup_policy, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.LookupPolicy(ctx, input)
				}
				var results []*svc.LookupPolicyOutput
				p := svc.NewLookupPolicyPaginator(client, input)
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
		"publish-schema": {
			Name:   "publish-schema",
			Fields: fields_publish_schema,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PublishSchemaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_publish_schema, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PublishSchema(ctx, input)
			},
		},
		"put-schema-from-json": {
			Name:   "put-schema-from-json",
			Fields: fields_put_schema_from_json,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutSchemaFromJsonInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_schema_from_json, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutSchemaFromJson(ctx, input)
			},
		},
		"remove-facet-from-object": {
			Name:   "remove-facet-from-object",
			Fields: fields_remove_facet_from_object,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveFacetFromObjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_facet_from_object, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveFacetFromObject(ctx, input)
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
		"update-facet": {
			Name:   "update-facet",
			Fields: fields_update_facet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFacetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_facet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFacet(ctx, input)
			},
		},
		"update-link-attributes": {
			Name:   "update-link-attributes",
			Fields: fields_update_link_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateLinkAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_link_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateLinkAttributes(ctx, input)
			},
		},
		"update-object-attributes": {
			Name:   "update-object-attributes",
			Fields: fields_update_object_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateObjectAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_object_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateObjectAttributes(ctx, input)
			},
		},
		"update-schema": {
			Name:   "update-schema",
			Fields: fields_update_schema,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSchemaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_schema, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSchema(ctx, input)
			},
		},
		"update-typed-link-facet": {
			Name:   "update-typed-link-facet",
			Fields: fields_update_typed_link_facet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTypedLinkFacetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_typed_link_facet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTypedLinkFacet(ctx, input)
			},
		},
		"upgrade-applied-schema": {
			Name:   "upgrade-applied-schema",
			Fields: fields_upgrade_applied_schema,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpgradeAppliedSchemaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_upgrade_applied_schema, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpgradeAppliedSchema(ctx, input)
			},
		},
		"upgrade-published-schema": {
			Name:   "upgrade-published-schema",
			Fields: fields_upgrade_published_schema,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpgradePublishedSchemaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_upgrade_published_schema, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpgradePublishedSchema(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("clouddirectory", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
