package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/kendra"
)

var fields_associate_entities_to_experience = []leanruntime.Field{
	{Name: "EntityList", Flag: "entity-list", Type: "[]types.EntityConfiguration", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
}

var fields_associate_personas_to_entities = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
	{Name: "Personas", Flag: "personas", Type: "[]types.EntityPersonaConfiguration", Required: true},
}

var fields_batch_delete_document = []leanruntime.Field{
	{Name: "DataSourceSyncJobMetricTarget", Flag: "data-source-sync-job-metric-target", Type: "*types.DataSourceSyncJobMetricTarget", Required: false},
	{Name: "DocumentIdList", Flag: "document-id-list", Type: "[]string", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
}

var fields_batch_delete_featured_results_set = []leanruntime.Field{
	{Name: "FeaturedResultsSetIds", Flag: "featured-results-set-ids", Type: "[]string", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
}

var fields_batch_get_document_status = []leanruntime.Field{
	{Name: "DocumentInfoList", Flag: "document-info-list", Type: "[]types.DocumentInfo", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
}

var fields_batch_put_document = []leanruntime.Field{
	{Name: "CustomDocumentEnrichmentConfiguration", Flag: "custom-document-enrichment-configuration", Type: "*types.CustomDocumentEnrichmentConfiguration", Required: false},
	{Name: "Documents", Flag: "documents", Type: "[]types.Document", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
}

var fields_clear_query_suggestions = []leanruntime.Field{
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
}

var fields_create_access_control_configuration = []leanruntime.Field{
	{Name: "AccessControlList", Flag: "access-control-list", Type: "[]types.Principal", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "HierarchicalAccessControlList", Flag: "hierarchical-access-control-list", Type: "[]types.HierarchicalPrincipal", Required: false},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_create_data_source = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Configuration", Flag: "configuration", Type: "*types.DataSourceConfiguration", Required: false},
	{Name: "CustomDocumentEnrichmentConfiguration", Flag: "custom-document-enrichment-configuration", Type: "*types.CustomDocumentEnrichmentConfiguration", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
	{Name: "LanguageCode", Flag: "language-code", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "Schedule", Flag: "schedule", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Type", Flag: "type", Type: "types.DataSourceType", Required: true},
	{Name: "VpcConfiguration", Flag: "vpc-configuration", Type: "*types.DataSourceVpcConfiguration", Required: false},
}

var fields_create_experience = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Configuration", Flag: "configuration", Type: "*types.ExperienceConfiguration", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
}

var fields_create_faq = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "FileFormat", Flag: "file-format", Type: "types.FaqFileFormat", Required: false},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
	{Name: "LanguageCode", Flag: "language-code", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "S3Path", Flag: "s3-path", Type: "*types.S3Path", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_featured_results_set = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "FeaturedDocuments", Flag: "featured-documents", Type: "[]types.FeaturedDocument", Required: false},
	{Name: "FeaturedResultsSetName", Flag: "featured-results-set-name", Type: "*string", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
	{Name: "QueryTexts", Flag: "query-texts", Type: "[]string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.FeaturedResultsSetStatus", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_index = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Edition", Flag: "edition", Type: "types.IndexEdition", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "ServerSideEncryptionConfiguration", Flag: "server-side-encryption-configuration", Type: "*types.ServerSideEncryptionConfiguration", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "UserContextPolicy", Flag: "user-context-policy", Type: "types.UserContextPolicy", Required: false},
	{Name: "UserGroupResolutionConfiguration", Flag: "user-group-resolution-configuration", Type: "*types.UserGroupResolutionConfiguration", Required: false},
	{Name: "UserTokenConfigurations", Flag: "user-token-configurations", Type: "[]types.UserTokenConfiguration", Required: false},
}

var fields_create_query_suggestions_block_list = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "SourceS3Path", Flag: "source-s3-path", Type: "*types.S3Path", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_thesaurus = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "SourceS3Path", Flag: "source-s3-path", Type: "*types.S3Path", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_delete_access_control_configuration = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
}

var fields_delete_data_source = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
}

var fields_delete_experience = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
}

var fields_delete_faq = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
}

var fields_delete_index = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_principal_mapping = []leanruntime.Field{
	{Name: "DataSourceId", Flag: "data-source-id", Type: "*string", Required: false},
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
	{Name: "OrderingId", Flag: "ordering-id", Type: "*int64", Required: false},
}

var fields_delete_query_suggestions_block_list = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
}

var fields_delete_thesaurus = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
}

var fields_describe_access_control_configuration = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
}

var fields_describe_data_source = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
}

var fields_describe_experience = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
}

var fields_describe_faq = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
}

var fields_describe_featured_results_set = []leanruntime.Field{
	{Name: "FeaturedResultsSetId", Flag: "featured-results-set-id", Type: "*string", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
}

var fields_describe_index = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_describe_principal_mapping = []leanruntime.Field{
	{Name: "DataSourceId", Flag: "data-source-id", Type: "*string", Required: false},
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
}

var fields_describe_query_suggestions_block_list = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
}

var fields_describe_query_suggestions_config = []leanruntime.Field{
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
}

var fields_describe_thesaurus = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
}

var fields_disassociate_entities_from_experience = []leanruntime.Field{
	{Name: "EntityList", Flag: "entity-list", Type: "[]types.EntityConfiguration", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
}

var fields_disassociate_personas_from_entities = []leanruntime.Field{
	{Name: "EntityIds", Flag: "entity-ids", Type: "[]string", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
}

var fields_get_query_suggestions = []leanruntime.Field{
	{Name: "AttributeSuggestionsConfig", Flag: "attribute-suggestions-config", Type: "*types.AttributeSuggestionsGetConfig", Required: false},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
	{Name: "MaxSuggestionsCount", Flag: "max-suggestions-count", Type: "*int32", Required: false},
	{Name: "QueryText", Flag: "query-text", Type: "*string", Required: true},
	{Name: "SuggestionTypes", Flag: "suggestion-types", Type: "[]types.SuggestionType", Required: false},
}

var fields_get_snapshots = []leanruntime.Field{
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
	{Name: "Interval", Flag: "interval", Type: "types.Interval", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MetricType", Flag: "metric-type", Type: "types.MetricType", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_access_control_configurations = []leanruntime.Field{
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_data_source_sync_jobs = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartTimeFilter", Flag: "start-time-filter", Type: "*types.TimeRange", Required: false},
	{Name: "StatusFilter", Flag: "status-filter", Type: "types.DataSourceSyncJobStatus", Required: false},
}

var fields_list_data_sources = []leanruntime.Field{
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_entity_personas = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_experience_entities = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_experiences = []leanruntime.Field{
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_faqs = []leanruntime.Field{
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_featured_results_sets = []leanruntime.Field{
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_groups_older_than_ordering_id = []leanruntime.Field{
	{Name: "DataSourceId", Flag: "data-source-id", Type: "*string", Required: false},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OrderingId", Flag: "ordering-id", Type: "*int64", Required: true},
}

var fields_list_indices = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_query_suggestions_block_lists = []leanruntime.Field{
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_thesauri = []leanruntime.Field{
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_put_principal_mapping = []leanruntime.Field{
	{Name: "DataSourceId", Flag: "data-source-id", Type: "*string", Required: false},
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: true},
	{Name: "GroupMembers", Flag: "group-members", Type: "*types.GroupMembers", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
	{Name: "OrderingId", Flag: "ordering-id", Type: "*int64", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
}

var fields_query = []leanruntime.Field{
	{Name: "AttributeFilter", Flag: "attribute-filter", Type: "*types.AttributeFilter", Required: false},
	{Name: "CollapseConfiguration", Flag: "collapse-configuration", Type: "*types.CollapseConfiguration", Required: false},
	{Name: "DocumentRelevanceOverrideConfigurations", Flag: "document-relevance-override-configurations", Type: "[]types.DocumentRelevanceConfiguration", Required: false},
	{Name: "Facets", Flag: "facets", Type: "[]types.Facet", Required: false},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
	{Name: "PageNumber", Flag: "page-number", Type: "*int32", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
	{Name: "QueryResultTypeFilter", Flag: "query-result-type-filter", Type: "types.QueryResultType", Required: false},
	{Name: "QueryText", Flag: "query-text", Type: "*string", Required: false},
	{Name: "RequestedDocumentAttributes", Flag: "requested-document-attributes", Type: "[]string", Required: false},
	{Name: "SortingConfiguration", Flag: "sorting-configuration", Type: "*types.SortingConfiguration", Required: false},
	{Name: "SortingConfigurations", Flag: "sorting-configurations", Type: "[]types.SortingConfiguration", Required: false},
	{Name: "SpellCorrectionConfiguration", Flag: "spell-correction-configuration", Type: "*types.SpellCorrectionConfiguration", Required: false},
	{Name: "UserContext", Flag: "user-context", Type: "*types.UserContext", Required: false},
	{Name: "VisitorId", Flag: "visitor-id", Type: "*string", Required: false},
}

var fields_retrieve = []leanruntime.Field{
	{Name: "AttributeFilter", Flag: "attribute-filter", Type: "*types.AttributeFilter", Required: false},
	{Name: "DocumentRelevanceOverrideConfigurations", Flag: "document-relevance-override-configurations", Type: "[]types.DocumentRelevanceConfiguration", Required: false},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
	{Name: "PageNumber", Flag: "page-number", Type: "*int32", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
	{Name: "QueryText", Flag: "query-text", Type: "*string", Required: true},
	{Name: "RequestedDocumentAttributes", Flag: "requested-document-attributes", Type: "[]string", Required: false},
	{Name: "UserContext", Flag: "user-context", Type: "*types.UserContext", Required: false},
}

var fields_start_data_source_sync_job = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
}

var fields_stop_data_source_sync_job = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
}

var fields_submit_feedback = []leanruntime.Field{
	{Name: "ClickFeedbackItems", Flag: "click-feedback-items", Type: "[]types.ClickFeedback", Required: false},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
	{Name: "QueryId", Flag: "query-id", Type: "*string", Required: true},
	{Name: "RelevanceFeedbackItems", Flag: "relevance-feedback-items", Type: "[]types.RelevanceFeedback", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_access_control_configuration = []leanruntime.Field{
	{Name: "AccessControlList", Flag: "access-control-list", Type: "[]types.Principal", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "HierarchicalAccessControlList", Flag: "hierarchical-access-control-list", Type: "[]types.HierarchicalPrincipal", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_data_source = []leanruntime.Field{
	{Name: "Configuration", Flag: "configuration", Type: "*types.DataSourceConfiguration", Required: false},
	{Name: "CustomDocumentEnrichmentConfiguration", Flag: "custom-document-enrichment-configuration", Type: "*types.CustomDocumentEnrichmentConfiguration", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
	{Name: "LanguageCode", Flag: "language-code", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "Schedule", Flag: "schedule", Type: "*string", Required: false},
	{Name: "VpcConfiguration", Flag: "vpc-configuration", Type: "*types.DataSourceVpcConfiguration", Required: false},
}

var fields_update_experience = []leanruntime.Field{
	{Name: "Configuration", Flag: "configuration", Type: "*types.ExperienceConfiguration", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
}

var fields_update_featured_results_set = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "FeaturedDocuments", Flag: "featured-documents", Type: "[]types.FeaturedDocument", Required: false},
	{Name: "FeaturedResultsSetId", Flag: "featured-results-set-id", Type: "*string", Required: true},
	{Name: "FeaturedResultsSetName", Flag: "featured-results-set-name", Type: "*string", Required: false},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
	{Name: "QueryTexts", Flag: "query-texts", Type: "[]string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.FeaturedResultsSetStatus", Required: false},
}

var fields_update_index = []leanruntime.Field{
	{Name: "CapacityUnits", Flag: "capacity-units", Type: "*types.CapacityUnitsConfiguration", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DocumentMetadataConfigurationUpdates", Flag: "document-metadata-configuration-updates", Type: "[]types.DocumentMetadataConfiguration", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "UserContextPolicy", Flag: "user-context-policy", Type: "types.UserContextPolicy", Required: false},
	{Name: "UserGroupResolutionConfiguration", Flag: "user-group-resolution-configuration", Type: "*types.UserGroupResolutionConfiguration", Required: false},
	{Name: "UserTokenConfigurations", Flag: "user-token-configurations", Type: "[]types.UserTokenConfiguration", Required: false},
}

var fields_update_query_suggestions_block_list = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "SourceS3Path", Flag: "source-s3-path", Type: "*types.S3Path", Required: false},
}

var fields_update_query_suggestions_config = []leanruntime.Field{
	{Name: "AttributeSuggestionsConfig", Flag: "attribute-suggestions-config", Type: "*types.AttributeSuggestionsUpdateConfig", Required: false},
	{Name: "IncludeQueriesWithoutUserInformation", Flag: "include-queries-without-user-information", Type: "*bool", Required: false},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
	{Name: "MinimumNumberOfQueryingUsers", Flag: "minimum-number-of-querying-users", Type: "*int32", Required: false},
	{Name: "MinimumQueryCount", Flag: "minimum-query-count", Type: "*int32", Required: false},
	{Name: "Mode", Flag: "mode", Type: "types.Mode", Required: false},
	{Name: "QueryLogLookBackWindowInDays", Flag: "query-log-look-back-window-in-days", Type: "*int32", Required: false},
}

var fields_update_thesaurus = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "SourceS3Path", Flag: "source-s3-path", Type: "*types.S3Path", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-entities-to-experience": {
			Name:   "associate-entities-to-experience",
			Fields: fields_associate_entities_to_experience,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateEntitiesToExperienceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_entities_to_experience, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateEntitiesToExperience(ctx, input)
			},
		},
		"associate-personas-to-entities": {
			Name:   "associate-personas-to-entities",
			Fields: fields_associate_personas_to_entities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociatePersonasToEntitiesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_personas_to_entities, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociatePersonasToEntities(ctx, input)
			},
		},
		"batch-delete-document": {
			Name:   "batch-delete-document",
			Fields: fields_batch_delete_document,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDeleteDocumentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_delete_document, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDeleteDocument(ctx, input)
			},
		},
		"batch-delete-featured-results-set": {
			Name:   "batch-delete-featured-results-set",
			Fields: fields_batch_delete_featured_results_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDeleteFeaturedResultsSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_delete_featured_results_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDeleteFeaturedResultsSet(ctx, input)
			},
		},
		"batch-get-document-status": {
			Name:   "batch-get-document-status",
			Fields: fields_batch_get_document_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetDocumentStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_document_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetDocumentStatus(ctx, input)
			},
		},
		"batch-put-document": {
			Name:   "batch-put-document",
			Fields: fields_batch_put_document,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchPutDocumentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_put_document, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchPutDocument(ctx, input)
			},
		},
		"clear-query-suggestions": {
			Name:   "clear-query-suggestions",
			Fields: fields_clear_query_suggestions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ClearQuerySuggestionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_clear_query_suggestions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ClearQuerySuggestions(ctx, input)
			},
		},
		"create-access-control-configuration": {
			Name:   "create-access-control-configuration",
			Fields: fields_create_access_control_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAccessControlConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_access_control_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAccessControlConfiguration(ctx, input)
			},
		},
		"create-data-source": {
			Name:   "create-data-source",
			Fields: fields_create_data_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDataSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_data_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDataSource(ctx, input)
			},
		},
		"create-experience": {
			Name:   "create-experience",
			Fields: fields_create_experience,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateExperienceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_experience, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateExperience(ctx, input)
			},
		},
		"create-faq": {
			Name:   "create-faq",
			Fields: fields_create_faq,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFaqInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_faq, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFaq(ctx, input)
			},
		},
		"create-featured-results-set": {
			Name:   "create-featured-results-set",
			Fields: fields_create_featured_results_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFeaturedResultsSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_featured_results_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFeaturedResultsSet(ctx, input)
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
		"create-query-suggestions-block-list": {
			Name:   "create-query-suggestions-block-list",
			Fields: fields_create_query_suggestions_block_list,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateQuerySuggestionsBlockListInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_query_suggestions_block_list, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateQuerySuggestionsBlockList(ctx, input)
			},
		},
		"create-thesaurus": {
			Name:   "create-thesaurus",
			Fields: fields_create_thesaurus,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateThesaurusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_thesaurus, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateThesaurus(ctx, input)
			},
		},
		"delete-access-control-configuration": {
			Name:   "delete-access-control-configuration",
			Fields: fields_delete_access_control_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAccessControlConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_access_control_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAccessControlConfiguration(ctx, input)
			},
		},
		"delete-data-source": {
			Name:   "delete-data-source",
			Fields: fields_delete_data_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDataSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_data_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDataSource(ctx, input)
			},
		},
		"delete-experience": {
			Name:   "delete-experience",
			Fields: fields_delete_experience,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteExperienceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_experience, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteExperience(ctx, input)
			},
		},
		"delete-faq": {
			Name:   "delete-faq",
			Fields: fields_delete_faq,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFaqInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_faq, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFaq(ctx, input)
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
		"delete-principal-mapping": {
			Name:   "delete-principal-mapping",
			Fields: fields_delete_principal_mapping,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePrincipalMappingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_principal_mapping, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePrincipalMapping(ctx, input)
			},
		},
		"delete-query-suggestions-block-list": {
			Name:   "delete-query-suggestions-block-list",
			Fields: fields_delete_query_suggestions_block_list,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteQuerySuggestionsBlockListInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_query_suggestions_block_list, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteQuerySuggestionsBlockList(ctx, input)
			},
		},
		"delete-thesaurus": {
			Name:   "delete-thesaurus",
			Fields: fields_delete_thesaurus,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteThesaurusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_thesaurus, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteThesaurus(ctx, input)
			},
		},
		"describe-access-control-configuration": {
			Name:   "describe-access-control-configuration",
			Fields: fields_describe_access_control_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAccessControlConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_access_control_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAccessControlConfiguration(ctx, input)
			},
		},
		"describe-data-source": {
			Name:   "describe-data-source",
			Fields: fields_describe_data_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDataSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_data_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDataSource(ctx, input)
			},
		},
		"describe-experience": {
			Name:   "describe-experience",
			Fields: fields_describe_experience,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeExperienceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_experience, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeExperience(ctx, input)
			},
		},
		"describe-faq": {
			Name:   "describe-faq",
			Fields: fields_describe_faq,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFaqInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_faq, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeFaq(ctx, input)
			},
		},
		"describe-featured-results-set": {
			Name:   "describe-featured-results-set",
			Fields: fields_describe_featured_results_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFeaturedResultsSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_featured_results_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeFeaturedResultsSet(ctx, input)
			},
		},
		"describe-index": {
			Name:   "describe-index",
			Fields: fields_describe_index,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeIndexInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_index, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeIndex(ctx, input)
			},
		},
		"describe-principal-mapping": {
			Name:   "describe-principal-mapping",
			Fields: fields_describe_principal_mapping,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePrincipalMappingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_principal_mapping, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribePrincipalMapping(ctx, input)
			},
		},
		"describe-query-suggestions-block-list": {
			Name:   "describe-query-suggestions-block-list",
			Fields: fields_describe_query_suggestions_block_list,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeQuerySuggestionsBlockListInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_query_suggestions_block_list, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeQuerySuggestionsBlockList(ctx, input)
			},
		},
		"describe-query-suggestions-config": {
			Name:   "describe-query-suggestions-config",
			Fields: fields_describe_query_suggestions_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeQuerySuggestionsConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_query_suggestions_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeQuerySuggestionsConfig(ctx, input)
			},
		},
		"describe-thesaurus": {
			Name:   "describe-thesaurus",
			Fields: fields_describe_thesaurus,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeThesaurusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_thesaurus, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeThesaurus(ctx, input)
			},
		},
		"disassociate-entities-from-experience": {
			Name:   "disassociate-entities-from-experience",
			Fields: fields_disassociate_entities_from_experience,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateEntitiesFromExperienceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_entities_from_experience, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateEntitiesFromExperience(ctx, input)
			},
		},
		"disassociate-personas-from-entities": {
			Name:   "disassociate-personas-from-entities",
			Fields: fields_disassociate_personas_from_entities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociatePersonasFromEntitiesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_personas_from_entities, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociatePersonasFromEntities(ctx, input)
			},
		},
		"get-query-suggestions": {
			Name:   "get-query-suggestions",
			Fields: fields_get_query_suggestions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetQuerySuggestionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_query_suggestions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetQuerySuggestions(ctx, input)
			},
		},
		"get-snapshots": {
			Name:   "get-snapshots",
			Fields: fields_get_snapshots,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSnapshotsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_snapshots, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetSnapshots(ctx, input)
				}
				var results []*svc.GetSnapshotsOutput
				p := svc.NewGetSnapshotsPaginator(client, input)
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
		"list-access-control-configurations": {
			Name:   "list-access-control-configurations",
			Fields: fields_list_access_control_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAccessControlConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_access_control_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAccessControlConfigurations(ctx, input)
				}
				var results []*svc.ListAccessControlConfigurationsOutput
				p := svc.NewListAccessControlConfigurationsPaginator(client, input)
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
		"list-data-source-sync-jobs": {
			Name:   "list-data-source-sync-jobs",
			Fields: fields_list_data_source_sync_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDataSourceSyncJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_data_source_sync_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDataSourceSyncJobs(ctx, input)
				}
				var results []*svc.ListDataSourceSyncJobsOutput
				p := svc.NewListDataSourceSyncJobsPaginator(client, input)
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
		"list-data-sources": {
			Name:   "list-data-sources",
			Fields: fields_list_data_sources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDataSourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_data_sources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDataSources(ctx, input)
				}
				var results []*svc.ListDataSourcesOutput
				p := svc.NewListDataSourcesPaginator(client, input)
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
		"list-entity-personas": {
			Name:   "list-entity-personas",
			Fields: fields_list_entity_personas,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEntityPersonasInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_entity_personas, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEntityPersonas(ctx, input)
				}
				var results []*svc.ListEntityPersonasOutput
				p := svc.NewListEntityPersonasPaginator(client, input)
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
		"list-experience-entities": {
			Name:   "list-experience-entities",
			Fields: fields_list_experience_entities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListExperienceEntitiesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_experience_entities, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListExperienceEntities(ctx, input)
				}
				var results []*svc.ListExperienceEntitiesOutput
				p := svc.NewListExperienceEntitiesPaginator(client, input)
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
		"list-experiences": {
			Name:   "list-experiences",
			Fields: fields_list_experiences,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListExperiencesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_experiences, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListExperiences(ctx, input)
				}
				var results []*svc.ListExperiencesOutput
				p := svc.NewListExperiencesPaginator(client, input)
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
		"list-faqs": {
			Name:   "list-faqs",
			Fields: fields_list_faqs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFaqsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_faqs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFaqs(ctx, input)
				}
				var results []*svc.ListFaqsOutput
				p := svc.NewListFaqsPaginator(client, input)
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
		"list-featured-results-sets": {
			Name:   "list-featured-results-sets",
			Fields: fields_list_featured_results_sets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFeaturedResultsSetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_featured_results_sets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListFeaturedResultsSets(ctx, input)
			},
		},
		"list-groups-older-than-ordering-id": {
			Name:   "list-groups-older-than-ordering-id",
			Fields: fields_list_groups_older_than_ordering_id,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGroupsOlderThanOrderingIdInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_groups_older_than_ordering_id, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListGroupsOlderThanOrderingId(ctx, input)
				}
				var results []*svc.ListGroupsOlderThanOrderingIdOutput
				p := svc.NewListGroupsOlderThanOrderingIdPaginator(client, input)
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
		"list-indices": {
			Name:   "list-indices",
			Fields: fields_list_indices,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIndicesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_indices, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListIndices(ctx, input)
				}
				var results []*svc.ListIndicesOutput
				p := svc.NewListIndicesPaginator(client, input)
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
		"list-query-suggestions-block-lists": {
			Name:   "list-query-suggestions-block-lists",
			Fields: fields_list_query_suggestions_block_lists,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListQuerySuggestionsBlockListsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_query_suggestions_block_lists, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListQuerySuggestionsBlockLists(ctx, input)
				}
				var results []*svc.ListQuerySuggestionsBlockListsOutput
				p := svc.NewListQuerySuggestionsBlockListsPaginator(client, input)
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
		"list-thesauri": {
			Name:   "list-thesauri",
			Fields: fields_list_thesauri,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListThesauriInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_thesauri, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListThesauri(ctx, input)
				}
				var results []*svc.ListThesauriOutput
				p := svc.NewListThesauriPaginator(client, input)
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
		"put-principal-mapping": {
			Name:   "put-principal-mapping",
			Fields: fields_put_principal_mapping,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutPrincipalMappingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_principal_mapping, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutPrincipalMapping(ctx, input)
			},
		},
		"query": {
			Name:   "query",
			Fields: fields_query,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.QueryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_query, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.Query(ctx, input)
			},
		},
		"retrieve": {
			Name:   "retrieve",
			Fields: fields_retrieve,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RetrieveInput{}
				if _, err := leanruntime.ApplyInput(input, fields_retrieve, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.Retrieve(ctx, input)
			},
		},
		"start-data-source-sync-job": {
			Name:   "start-data-source-sync-job",
			Fields: fields_start_data_source_sync_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartDataSourceSyncJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_data_source_sync_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartDataSourceSyncJob(ctx, input)
			},
		},
		"stop-data-source-sync-job": {
			Name:   "stop-data-source-sync-job",
			Fields: fields_stop_data_source_sync_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopDataSourceSyncJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_data_source_sync_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopDataSourceSyncJob(ctx, input)
			},
		},
		"submit-feedback": {
			Name:   "submit-feedback",
			Fields: fields_submit_feedback,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SubmitFeedbackInput{}
				if _, err := leanruntime.ApplyInput(input, fields_submit_feedback, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SubmitFeedback(ctx, input)
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
		"update-access-control-configuration": {
			Name:   "update-access-control-configuration",
			Fields: fields_update_access_control_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAccessControlConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_access_control_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAccessControlConfiguration(ctx, input)
			},
		},
		"update-data-source": {
			Name:   "update-data-source",
			Fields: fields_update_data_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDataSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_data_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDataSource(ctx, input)
			},
		},
		"update-experience": {
			Name:   "update-experience",
			Fields: fields_update_experience,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateExperienceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_experience, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateExperience(ctx, input)
			},
		},
		"update-featured-results-set": {
			Name:   "update-featured-results-set",
			Fields: fields_update_featured_results_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFeaturedResultsSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_featured_results_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFeaturedResultsSet(ctx, input)
			},
		},
		"update-index": {
			Name:   "update-index",
			Fields: fields_update_index,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateIndexInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_index, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateIndex(ctx, input)
			},
		},
		"update-query-suggestions-block-list": {
			Name:   "update-query-suggestions-block-list",
			Fields: fields_update_query_suggestions_block_list,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateQuerySuggestionsBlockListInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_query_suggestions_block_list, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateQuerySuggestionsBlockList(ctx, input)
			},
		},
		"update-query-suggestions-config": {
			Name:   "update-query-suggestions-config",
			Fields: fields_update_query_suggestions_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateQuerySuggestionsConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_query_suggestions_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateQuerySuggestionsConfig(ctx, input)
			},
		},
		"update-thesaurus": {
			Name:   "update-thesaurus",
			Fields: fields_update_thesaurus,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateThesaurusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_thesaurus, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateThesaurus(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("kendra", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
