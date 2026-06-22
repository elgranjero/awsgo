package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/customerprofiles"
)

var fields_add_profile_key = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "KeyName", Flag: "key-name", Type: "*string", Required: true},
	{Name: "ProfileId", Flag: "profile-id", Type: "*string", Required: true},
	{Name: "Values", Flag: "values", Type: "[]string", Required: true},
}

var fields_batch_get_calculated_attribute_for_profile = []leanruntime.Field{
	{Name: "CalculatedAttributeName", Flag: "calculated-attribute-name", Type: "*string", Required: true},
	{Name: "ConditionOverrides", Flag: "condition-overrides", Type: "*types.ConditionOverrides", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "ProfileIds", Flag: "profile-ids", Type: "[]string", Required: true},
}

var fields_batch_get_profile = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "ProfileIds", Flag: "profile-ids", Type: "[]string", Required: true},
}

var fields_create_calculated_attribute_definition = []leanruntime.Field{
	{Name: "AttributeDetails", Flag: "attribute-details", Type: "*types.AttributeDetails", Required: true},
	{Name: "CalculatedAttributeName", Flag: "calculated-attribute-name", Type: "*string", Required: true},
	{Name: "Conditions", Flag: "conditions", Type: "*types.Conditions", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "Filter", Flag: "filter", Type: "*types.Filter", Required: false},
	{Name: "Statistic", Flag: "statistic", Type: "types.Statistic", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "UseHistoricalData", Flag: "use-historical-data", Type: "*bool", Required: false},
}

var fields_create_domain = []leanruntime.Field{
	{Name: "DataStore", Flag: "data-store", Type: "*types.DataStoreRequest", Required: false},
	{Name: "DeadLetterQueueUrl", Flag: "dead-letter-queue-url", Type: "*string", Required: false},
	{Name: "DefaultEncryptionKey", Flag: "default-encryption-key", Type: "*string", Required: false},
	{Name: "DefaultExpirationDays", Flag: "default-expiration-days", Type: "*int32", Required: true},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "Matching", Flag: "matching", Type: "*types.MatchingRequest", Required: false},
	{Name: "RuleBasedMatching", Flag: "rule-based-matching", Type: "*types.RuleBasedMatchingRequest", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_domain_layout = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: true},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "IsDefault", Flag: "is-default", Type: "bool", Required: false},
	{Name: "Layout", Flag: "layout", Type: "*string", Required: true},
	{Name: "LayoutDefinitionName", Flag: "layout-definition-name", Type: "*string", Required: true},
	{Name: "LayoutType", Flag: "layout-type", Type: "types.LayoutType", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_event_stream = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "EventStreamName", Flag: "event-stream-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Uri", Flag: "uri", Type: "*string", Required: true},
}

var fields_create_event_trigger = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "EventTriggerConditions", Flag: "event-trigger-conditions", Type: "[]types.EventTriggerCondition", Required: true},
	{Name: "EventTriggerLimits", Flag: "event-trigger-limits", Type: "*types.EventTriggerLimits", Required: false},
	{Name: "EventTriggerName", Flag: "event-trigger-name", Type: "*string", Required: true},
	{Name: "ObjectTypeName", Flag: "object-type-name", Type: "*string", Required: true},
	{Name: "SegmentFilter", Flag: "segment-filter", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_integration_workflow = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "IntegrationConfig", Flag: "integration-config", Type: "*types.IntegrationConfig", Required: true},
	{Name: "ObjectTypeName", Flag: "object-type-name", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "WorkflowType", Flag: "workflow-type", Type: "types.WorkflowType", Required: true},
}

var fields_create_profile = []leanruntime.Field{
	{Name: "AccountNumber", Flag: "account-number", Type: "*string", Required: false},
	{Name: "AdditionalInformation", Flag: "additional-information", Type: "*string", Required: false},
	{Name: "Address", Flag: "address", Type: "*types.Address", Required: false},
	{Name: "Attributes", Flag: "attributes", Type: "map[string]string", Required: false},
	{Name: "BillingAddress", Flag: "billing-address", Type: "*types.Address", Required: false},
	{Name: "BirthDate", Flag: "birth-date", Type: "*string", Required: false},
	{Name: "BusinessEmailAddress", Flag: "business-email-address", Type: "*string", Required: false},
	{Name: "BusinessName", Flag: "business-name", Type: "*string", Required: false},
	{Name: "BusinessPhoneNumber", Flag: "business-phone-number", Type: "*string", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "EmailAddress", Flag: "email-address", Type: "*string", Required: false},
	{Name: "EngagementPreferences", Flag: "engagement-preferences", Type: "*types.EngagementPreferences", Required: false},
	{Name: "FirstName", Flag: "first-name", Type: "*string", Required: false},
	{Name: "Gender", Flag: "gender", Type: "types.Gender", Required: false},
	{Name: "GenderString", Flag: "gender-string", Type: "*string", Required: false},
	{Name: "HomePhoneNumber", Flag: "home-phone-number", Type: "*string", Required: false},
	{Name: "LastName", Flag: "last-name", Type: "*string", Required: false},
	{Name: "MailingAddress", Flag: "mailing-address", Type: "*types.Address", Required: false},
	{Name: "MiddleName", Flag: "middle-name", Type: "*string", Required: false},
	{Name: "MobilePhoneNumber", Flag: "mobile-phone-number", Type: "*string", Required: false},
	{Name: "PartyType", Flag: "party-type", Type: "types.PartyType", Required: false},
	{Name: "PartyTypeString", Flag: "party-type-string", Type: "*string", Required: false},
	{Name: "PersonalEmailAddress", Flag: "personal-email-address", Type: "*string", Required: false},
	{Name: "PhoneNumber", Flag: "phone-number", Type: "*string", Required: false},
	{Name: "ProfileType", Flag: "profile-type", Type: "types.ProfileType", Required: false},
	{Name: "ShippingAddress", Flag: "shipping-address", Type: "*types.Address", Required: false},
}

var fields_create_recommender = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "RecommenderConfig", Flag: "recommender-config", Type: "*types.RecommenderConfig", Required: false},
	{Name: "RecommenderName", Flag: "recommender-name", Type: "*string", Required: true},
	{Name: "RecommenderRecipeName", Flag: "recommender-recipe-name", Type: "types.RecommenderRecipeName", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_segment_definition = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: true},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "SegmentDefinitionName", Flag: "segment-definition-name", Type: "*string", Required: true},
	{Name: "SegmentGroups", Flag: "segment-groups", Type: "*types.SegmentGroup", Required: false},
	{Name: "SegmentSqlQuery", Flag: "segment-sql-query", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_segment_estimate = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "SegmentQuery", Flag: "segment-query", Type: "*types.SegmentGroupStructure", Required: false},
	{Name: "SegmentSqlQuery", Flag: "segment-sql-query", Type: "*string", Required: false},
}

var fields_create_segment_snapshot = []leanruntime.Field{
	{Name: "DataFormat", Flag: "data-format", Type: "types.DataFormat", Required: true},
	{Name: "DestinationUri", Flag: "destination-uri", Type: "*string", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "EncryptionKey", Flag: "encryption-key", Type: "*string", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "SegmentDefinitionName", Flag: "segment-definition-name", Type: "*string", Required: true},
}

var fields_create_upload_job = []leanruntime.Field{
	{Name: "DataExpiry", Flag: "data-expiry", Type: "*int32", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: true},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "Fields", Flag: "fields", Type: "map[string]types.ObjectTypeField", Required: true},
	{Name: "UniqueKey", Flag: "unique-key", Type: "*string", Required: true},
}

var fields_delete_calculated_attribute_definition = []leanruntime.Field{
	{Name: "CalculatedAttributeName", Flag: "calculated-attribute-name", Type: "*string", Required: true},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_delete_domain = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_delete_domain_layout = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "LayoutDefinitionName", Flag: "layout-definition-name", Type: "*string", Required: true},
}

var fields_delete_domain_object_type = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "ObjectTypeName", Flag: "object-type-name", Type: "*string", Required: true},
}

var fields_delete_event_stream = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "EventStreamName", Flag: "event-stream-name", Type: "*string", Required: true},
}

var fields_delete_event_trigger = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "EventTriggerName", Flag: "event-trigger-name", Type: "*string", Required: true},
}

var fields_delete_integration = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "Uri", Flag: "uri", Type: "*string", Required: true},
}

var fields_delete_profile = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "ProfileId", Flag: "profile-id", Type: "*string", Required: true},
}

var fields_delete_profile_key = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "KeyName", Flag: "key-name", Type: "*string", Required: true},
	{Name: "ProfileId", Flag: "profile-id", Type: "*string", Required: true},
	{Name: "Values", Flag: "values", Type: "[]string", Required: true},
}

var fields_delete_profile_object = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "ObjectTypeName", Flag: "object-type-name", Type: "*string", Required: true},
	{Name: "ProfileId", Flag: "profile-id", Type: "*string", Required: true},
	{Name: "ProfileObjectUniqueKey", Flag: "profile-object-unique-key", Type: "*string", Required: true},
}

var fields_delete_profile_object_type = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "ObjectTypeName", Flag: "object-type-name", Type: "*string", Required: true},
}

var fields_delete_recommender = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "RecommenderName", Flag: "recommender-name", Type: "*string", Required: true},
}

var fields_delete_segment_definition = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "SegmentDefinitionName", Flag: "segment-definition-name", Type: "*string", Required: true},
}

var fields_delete_workflow = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "WorkflowId", Flag: "workflow-id", Type: "*string", Required: true},
}

var fields_detect_profile_object_type = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "Objects", Flag: "objects", Type: "[]string", Required: true},
}

var fields_get_auto_merging_preview = []leanruntime.Field{
	{Name: "ConflictResolution", Flag: "conflict-resolution", Type: "*types.ConflictResolution", Required: true},
	{Name: "Consolidation", Flag: "consolidation", Type: "*types.Consolidation", Required: true},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "MinAllowedConfidenceScoreForMerging", Flag: "min-allowed-confidence-score-for-merging", Type: "*float64", Required: false},
}

var fields_get_calculated_attribute_definition = []leanruntime.Field{
	{Name: "CalculatedAttributeName", Flag: "calculated-attribute-name", Type: "*string", Required: true},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_get_calculated_attribute_for_profile = []leanruntime.Field{
	{Name: "CalculatedAttributeName", Flag: "calculated-attribute-name", Type: "*string", Required: true},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "ProfileId", Flag: "profile-id", Type: "*string", Required: true},
}

var fields_get_domain = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_get_domain_layout = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "LayoutDefinitionName", Flag: "layout-definition-name", Type: "*string", Required: true},
}

var fields_get_domain_object_type = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "ObjectTypeName", Flag: "object-type-name", Type: "*string", Required: true},
}

var fields_get_event_stream = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "EventStreamName", Flag: "event-stream-name", Type: "*string", Required: true},
}

var fields_get_event_trigger = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "EventTriggerName", Flag: "event-trigger-name", Type: "*string", Required: true},
}

var fields_get_identity_resolution_job = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_get_integration = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "Uri", Flag: "uri", Type: "*string", Required: true},
}

var fields_get_matches = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_object_type_attribute_statistics = []leanruntime.Field{
	{Name: "AttributeName", Flag: "attribute-name", Type: "*string", Required: true},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "ObjectTypeName", Flag: "object-type-name", Type: "*string", Required: true},
}

var fields_get_profile_history_record = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "ProfileId", Flag: "profile-id", Type: "*string", Required: true},
}

var fields_get_profile_object_type = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "ObjectTypeName", Flag: "object-type-name", Type: "*string", Required: true},
}

var fields_get_profile_object_type_template = []leanruntime.Field{
	{Name: "TemplateId", Flag: "template-id", Type: "*string", Required: true},
}

var fields_get_profile_recommendations = []leanruntime.Field{
	{Name: "Context", Flag: "context", Type: "map[string]string", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "ProfileId", Flag: "profile-id", Type: "*string", Required: true},
	{Name: "RecommenderName", Flag: "recommender-name", Type: "*string", Required: true},
}

var fields_get_recommender = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "RecommenderName", Flag: "recommender-name", Type: "*string", Required: true},
	{Name: "TrainingMetricsCount", Flag: "training-metrics-count", Type: "*int32", Required: false},
}

var fields_get_segment_definition = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "SegmentDefinitionName", Flag: "segment-definition-name", Type: "*string", Required: true},
}

var fields_get_segment_estimate = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "EstimateId", Flag: "estimate-id", Type: "*string", Required: true},
}

var fields_get_segment_membership = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "ProfileIds", Flag: "profile-ids", Type: "[]string", Required: true},
	{Name: "SegmentDefinitionName", Flag: "segment-definition-name", Type: "*string", Required: true},
}

var fields_get_segment_snapshot = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "SegmentDefinitionName", Flag: "segment-definition-name", Type: "*string", Required: true},
	{Name: "SnapshotId", Flag: "snapshot-id", Type: "*string", Required: true},
}

var fields_get_similar_profiles = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "MatchType", Flag: "match-type", Type: "types.MatchType", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SearchKey", Flag: "search-key", Type: "*string", Required: true},
	{Name: "SearchValue", Flag: "search-value", Type: "*string", Required: true},
}

var fields_get_upload_job = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_get_upload_job_path = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_get_workflow = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "WorkflowId", Flag: "workflow-id", Type: "*string", Required: true},
}

var fields_get_workflow_steps = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "WorkflowId", Flag: "workflow-id", Type: "*string", Required: true},
}

var fields_list_account_integrations = []leanruntime.Field{
	{Name: "IncludeHidden", Flag: "include-hidden", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Uri", Flag: "uri", Type: "*string", Required: true},
}

var fields_list_calculated_attribute_definitions = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_calculated_attributes_for_profile = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProfileId", Flag: "profile-id", Type: "*string", Required: true},
}

var fields_list_domain_layouts = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_domain_object_types = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_domains = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_event_streams = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_event_triggers = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_identity_resolution_jobs = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_integrations = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "IncludeHidden", Flag: "include-hidden", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_object_type_attribute_values = []leanruntime.Field{
	{Name: "AttributeName", Flag: "attribute-name", Type: "*string", Required: true},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ObjectTypeName", Flag: "object-type-name", Type: "*string", Required: true},
}

var fields_list_object_type_attributes = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ObjectTypeName", Flag: "object-type-name", Type: "*string", Required: true},
}

var fields_list_profile_attribute_values = []leanruntime.Field{
	{Name: "AttributeName", Flag: "attribute-name", Type: "*string", Required: true},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_list_profile_history_records = []leanruntime.Field{
	{Name: "ActionType", Flag: "action-type", Type: "types.ActionType", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ObjectTypeName", Flag: "object-type-name", Type: "*string", Required: false},
	{Name: "PerformedBy", Flag: "performed-by", Type: "*string", Required: false},
	{Name: "ProfileId", Flag: "profile-id", Type: "*string", Required: true},
}

var fields_list_profile_object_type_templates = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_profile_object_types = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_profile_objects = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ObjectFilter", Flag: "object-filter", Type: "*types.ObjectFilter", Required: false},
	{Name: "ObjectTypeName", Flag: "object-type-name", Type: "*string", Required: true},
	{Name: "ProfileId", Flag: "profile-id", Type: "*string", Required: true},
}

var fields_list_recommender_recipes = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_recommenders = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_rule_based_matches = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_segment_definitions = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_upload_jobs = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_workflows = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "QueryEndDate", Flag: "query-end-date", Type: "*time.Time", Required: false},
	{Name: "QueryStartDate", Flag: "query-start-date", Type: "*time.Time", Required: false},
	{Name: "Status", Flag: "status", Type: "types.Status", Required: false},
	{Name: "WorkflowType", Flag: "workflow-type", Type: "types.WorkflowType", Required: false},
}

var fields_merge_profiles = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "FieldSourceProfileIds", Flag: "field-source-profile-ids", Type: "*types.FieldSourceProfileIds", Required: false},
	{Name: "MainProfileId", Flag: "main-profile-id", Type: "*string", Required: true},
	{Name: "ProfileIdsToBeMerged", Flag: "profile-ids-to-be-merged", Type: "[]string", Required: true},
}

var fields_put_domain_object_type = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "EncryptionKey", Flag: "encryption-key", Type: "*string", Required: false},
	{Name: "Fields", Flag: "fields", Type: "map[string]types.DomainObjectTypeField", Required: true},
	{Name: "ObjectTypeName", Flag: "object-type-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_put_integration = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "EventTriggerNames", Flag: "event-trigger-names", Type: "[]string", Required: false},
	{Name: "FlowDefinition", Flag: "flow-definition", Type: "*types.FlowDefinition", Required: false},
	{Name: "ObjectTypeName", Flag: "object-type-name", Type: "*string", Required: false},
	{Name: "ObjectTypeNames", Flag: "object-type-names", Type: "map[string]string", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "Scope", Flag: "scope", Type: "types.Scope", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Uri", Flag: "uri", Type: "*string", Required: false},
}

var fields_put_profile_object = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "Object", Flag: "object", Type: "*string", Required: true},
	{Name: "ObjectTypeName", Flag: "object-type-name", Type: "*string", Required: true},
}

var fields_put_profile_object_type = []leanruntime.Field{
	{Name: "AllowProfileCreation", Flag: "allow-profile-creation", Type: "bool", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "EncryptionKey", Flag: "encryption-key", Type: "*string", Required: false},
	{Name: "ExpirationDays", Flag: "expiration-days", Type: "*int32", Required: false},
	{Name: "Fields", Flag: "fields", Type: "map[string]types.ObjectTypeField", Required: false},
	{Name: "Keys", Flag: "keys", Type: "map[string][]types.ObjectTypeKey", Required: false},
	{Name: "MaxProfileObjectCount", Flag: "max-profile-object-count", Type: "*int32", Required: false},
	{Name: "ObjectTypeName", Flag: "object-type-name", Type: "*string", Required: true},
	{Name: "SourceLastUpdatedTimestampFormat", Flag: "source-last-updated-timestamp-format", Type: "*string", Required: false},
	{Name: "SourcePriority", Flag: "source-priority", Type: "*int32", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TemplateId", Flag: "template-id", Type: "*string", Required: false},
}

var fields_search_profiles = []leanruntime.Field{
	{Name: "AdditionalSearchKeys", Flag: "additional-search-keys", Type: "[]types.AdditionalSearchKey", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "KeyName", Flag: "key-name", Type: "*string", Required: true},
	{Name: "LogicalOperator", Flag: "logical-operator", Type: "types.LogicalOperator", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Values", Flag: "values", Type: "[]string", Required: true},
}

var fields_start_recommender = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "RecommenderName", Flag: "recommender-name", Type: "*string", Required: true},
}

var fields_start_upload_job = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_stop_recommender = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "RecommenderName", Flag: "recommender-name", Type: "*string", Required: true},
}

var fields_stop_upload_job = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_calculated_attribute_definition = []leanruntime.Field{
	{Name: "CalculatedAttributeName", Flag: "calculated-attribute-name", Type: "*string", Required: true},
	{Name: "Conditions", Flag: "conditions", Type: "*types.Conditions", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_update_domain = []leanruntime.Field{
	{Name: "DataStore", Flag: "data-store", Type: "*types.DataStoreRequest", Required: false},
	{Name: "DeadLetterQueueUrl", Flag: "dead-letter-queue-url", Type: "*string", Required: false},
	{Name: "DefaultEncryptionKey", Flag: "default-encryption-key", Type: "*string", Required: false},
	{Name: "DefaultExpirationDays", Flag: "default-expiration-days", Type: "*int32", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "Matching", Flag: "matching", Type: "*types.MatchingRequest", Required: false},
	{Name: "RuleBasedMatching", Flag: "rule-based-matching", Type: "*types.RuleBasedMatchingRequest", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_update_domain_layout = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "IsDefault", Flag: "is-default", Type: "bool", Required: false},
	{Name: "Layout", Flag: "layout", Type: "*string", Required: false},
	{Name: "LayoutDefinitionName", Flag: "layout-definition-name", Type: "*string", Required: true},
	{Name: "LayoutType", Flag: "layout-type", Type: "types.LayoutType", Required: false},
}

var fields_update_event_trigger = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "EventTriggerConditions", Flag: "event-trigger-conditions", Type: "[]types.EventTriggerCondition", Required: false},
	{Name: "EventTriggerLimits", Flag: "event-trigger-limits", Type: "*types.EventTriggerLimits", Required: false},
	{Name: "EventTriggerName", Flag: "event-trigger-name", Type: "*string", Required: true},
	{Name: "ObjectTypeName", Flag: "object-type-name", Type: "*string", Required: false},
	{Name: "SegmentFilter", Flag: "segment-filter", Type: "*string", Required: false},
}

var fields_update_profile = []leanruntime.Field{
	{Name: "AccountNumber", Flag: "account-number", Type: "*string", Required: false},
	{Name: "AdditionalInformation", Flag: "additional-information", Type: "*string", Required: false},
	{Name: "Address", Flag: "address", Type: "*types.UpdateAddress", Required: false},
	{Name: "Attributes", Flag: "attributes", Type: "map[string]string", Required: false},
	{Name: "BillingAddress", Flag: "billing-address", Type: "*types.UpdateAddress", Required: false},
	{Name: "BirthDate", Flag: "birth-date", Type: "*string", Required: false},
	{Name: "BusinessEmailAddress", Flag: "business-email-address", Type: "*string", Required: false},
	{Name: "BusinessName", Flag: "business-name", Type: "*string", Required: false},
	{Name: "BusinessPhoneNumber", Flag: "business-phone-number", Type: "*string", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "EmailAddress", Flag: "email-address", Type: "*string", Required: false},
	{Name: "EngagementPreferences", Flag: "engagement-preferences", Type: "*types.EngagementPreferences", Required: false},
	{Name: "FirstName", Flag: "first-name", Type: "*string", Required: false},
	{Name: "Gender", Flag: "gender", Type: "types.Gender", Required: false},
	{Name: "GenderString", Flag: "gender-string", Type: "*string", Required: false},
	{Name: "HomePhoneNumber", Flag: "home-phone-number", Type: "*string", Required: false},
	{Name: "LastName", Flag: "last-name", Type: "*string", Required: false},
	{Name: "MailingAddress", Flag: "mailing-address", Type: "*types.UpdateAddress", Required: false},
	{Name: "MiddleName", Flag: "middle-name", Type: "*string", Required: false},
	{Name: "MobilePhoneNumber", Flag: "mobile-phone-number", Type: "*string", Required: false},
	{Name: "PartyType", Flag: "party-type", Type: "types.PartyType", Required: false},
	{Name: "PartyTypeString", Flag: "party-type-string", Type: "*string", Required: false},
	{Name: "PersonalEmailAddress", Flag: "personal-email-address", Type: "*string", Required: false},
	{Name: "PhoneNumber", Flag: "phone-number", Type: "*string", Required: false},
	{Name: "ProfileId", Flag: "profile-id", Type: "*string", Required: true},
	{Name: "ProfileType", Flag: "profile-type", Type: "types.ProfileType", Required: false},
	{Name: "ShippingAddress", Flag: "shipping-address", Type: "*types.UpdateAddress", Required: false},
}

var fields_update_recommender = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "RecommenderConfig", Flag: "recommender-config", Type: "*types.RecommenderConfig", Required: false},
	{Name: "RecommenderName", Flag: "recommender-name", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"add-profile-key": {
			Name:   "add-profile-key",
			Fields: fields_add_profile_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddProfileKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_profile_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddProfileKey(ctx, input)
			},
		},
		"batch-get-calculated-attribute-for-profile": {
			Name:   "batch-get-calculated-attribute-for-profile",
			Fields: fields_batch_get_calculated_attribute_for_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetCalculatedAttributeForProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_calculated_attribute_for_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetCalculatedAttributeForProfile(ctx, input)
			},
		},
		"batch-get-profile": {
			Name:   "batch-get-profile",
			Fields: fields_batch_get_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetProfile(ctx, input)
			},
		},
		"create-calculated-attribute-definition": {
			Name:   "create-calculated-attribute-definition",
			Fields: fields_create_calculated_attribute_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCalculatedAttributeDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_calculated_attribute_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCalculatedAttributeDefinition(ctx, input)
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
		"create-domain-layout": {
			Name:   "create-domain-layout",
			Fields: fields_create_domain_layout,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDomainLayoutInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_domain_layout, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDomainLayout(ctx, input)
			},
		},
		"create-event-stream": {
			Name:   "create-event-stream",
			Fields: fields_create_event_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEventStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_event_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEventStream(ctx, input)
			},
		},
		"create-event-trigger": {
			Name:   "create-event-trigger",
			Fields: fields_create_event_trigger,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEventTriggerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_event_trigger, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEventTrigger(ctx, input)
			},
		},
		"create-integration-workflow": {
			Name:   "create-integration-workflow",
			Fields: fields_create_integration_workflow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateIntegrationWorkflowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_integration_workflow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateIntegrationWorkflow(ctx, input)
			},
		},
		"create-profile": {
			Name:   "create-profile",
			Fields: fields_create_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateProfile(ctx, input)
			},
		},
		"create-recommender": {
			Name:   "create-recommender",
			Fields: fields_create_recommender,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRecommenderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_recommender, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRecommender(ctx, input)
			},
		},
		"create-segment-definition": {
			Name:   "create-segment-definition",
			Fields: fields_create_segment_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSegmentDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_segment_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSegmentDefinition(ctx, input)
			},
		},
		"create-segment-estimate": {
			Name:   "create-segment-estimate",
			Fields: fields_create_segment_estimate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSegmentEstimateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_segment_estimate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSegmentEstimate(ctx, input)
			},
		},
		"create-segment-snapshot": {
			Name:   "create-segment-snapshot",
			Fields: fields_create_segment_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSegmentSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_segment_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSegmentSnapshot(ctx, input)
			},
		},
		"create-upload-job": {
			Name:   "create-upload-job",
			Fields: fields_create_upload_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateUploadJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_upload_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateUploadJob(ctx, input)
			},
		},
		"delete-calculated-attribute-definition": {
			Name:   "delete-calculated-attribute-definition",
			Fields: fields_delete_calculated_attribute_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCalculatedAttributeDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_calculated_attribute_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCalculatedAttributeDefinition(ctx, input)
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
		"delete-domain-layout": {
			Name:   "delete-domain-layout",
			Fields: fields_delete_domain_layout,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDomainLayoutInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_domain_layout, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDomainLayout(ctx, input)
			},
		},
		"delete-domain-object-type": {
			Name:   "delete-domain-object-type",
			Fields: fields_delete_domain_object_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDomainObjectTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_domain_object_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDomainObjectType(ctx, input)
			},
		},
		"delete-event-stream": {
			Name:   "delete-event-stream",
			Fields: fields_delete_event_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEventStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_event_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEventStream(ctx, input)
			},
		},
		"delete-event-trigger": {
			Name:   "delete-event-trigger",
			Fields: fields_delete_event_trigger,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEventTriggerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_event_trigger, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEventTrigger(ctx, input)
			},
		},
		"delete-integration": {
			Name:   "delete-integration",
			Fields: fields_delete_integration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIntegrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_integration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIntegration(ctx, input)
			},
		},
		"delete-profile": {
			Name:   "delete-profile",
			Fields: fields_delete_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteProfile(ctx, input)
			},
		},
		"delete-profile-key": {
			Name:   "delete-profile-key",
			Fields: fields_delete_profile_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteProfileKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_profile_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteProfileKey(ctx, input)
			},
		},
		"delete-profile-object": {
			Name:   "delete-profile-object",
			Fields: fields_delete_profile_object,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteProfileObjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_profile_object, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteProfileObject(ctx, input)
			},
		},
		"delete-profile-object-type": {
			Name:   "delete-profile-object-type",
			Fields: fields_delete_profile_object_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteProfileObjectTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_profile_object_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteProfileObjectType(ctx, input)
			},
		},
		"delete-recommender": {
			Name:   "delete-recommender",
			Fields: fields_delete_recommender,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRecommenderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_recommender, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRecommender(ctx, input)
			},
		},
		"delete-segment-definition": {
			Name:   "delete-segment-definition",
			Fields: fields_delete_segment_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSegmentDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_segment_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSegmentDefinition(ctx, input)
			},
		},
		"delete-workflow": {
			Name:   "delete-workflow",
			Fields: fields_delete_workflow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWorkflowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_workflow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWorkflow(ctx, input)
			},
		},
		"detect-profile-object-type": {
			Name:   "detect-profile-object-type",
			Fields: fields_detect_profile_object_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetectProfileObjectTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detect_profile_object_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetectProfileObjectType(ctx, input)
			},
		},
		"get-auto-merging-preview": {
			Name:   "get-auto-merging-preview",
			Fields: fields_get_auto_merging_preview,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAutoMergingPreviewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_auto_merging_preview, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAutoMergingPreview(ctx, input)
			},
		},
		"get-calculated-attribute-definition": {
			Name:   "get-calculated-attribute-definition",
			Fields: fields_get_calculated_attribute_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCalculatedAttributeDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_calculated_attribute_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCalculatedAttributeDefinition(ctx, input)
			},
		},
		"get-calculated-attribute-for-profile": {
			Name:   "get-calculated-attribute-for-profile",
			Fields: fields_get_calculated_attribute_for_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCalculatedAttributeForProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_calculated_attribute_for_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCalculatedAttributeForProfile(ctx, input)
			},
		},
		"get-domain": {
			Name:   "get-domain",
			Fields: fields_get_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDomain(ctx, input)
			},
		},
		"get-domain-layout": {
			Name:   "get-domain-layout",
			Fields: fields_get_domain_layout,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDomainLayoutInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_domain_layout, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDomainLayout(ctx, input)
			},
		},
		"get-domain-object-type": {
			Name:   "get-domain-object-type",
			Fields: fields_get_domain_object_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDomainObjectTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_domain_object_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDomainObjectType(ctx, input)
			},
		},
		"get-event-stream": {
			Name:   "get-event-stream",
			Fields: fields_get_event_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEventStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_event_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEventStream(ctx, input)
			},
		},
		"get-event-trigger": {
			Name:   "get-event-trigger",
			Fields: fields_get_event_trigger,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEventTriggerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_event_trigger, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEventTrigger(ctx, input)
			},
		},
		"get-identity-resolution-job": {
			Name:   "get-identity-resolution-job",
			Fields: fields_get_identity_resolution_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIdentityResolutionJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_identity_resolution_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIdentityResolutionJob(ctx, input)
			},
		},
		"get-integration": {
			Name:   "get-integration",
			Fields: fields_get_integration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIntegrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_integration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIntegration(ctx, input)
			},
		},
		"get-matches": {
			Name:   "get-matches",
			Fields: fields_get_matches,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMatchesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_matches, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMatches(ctx, input)
			},
		},
		"get-object-type-attribute-statistics": {
			Name:   "get-object-type-attribute-statistics",
			Fields: fields_get_object_type_attribute_statistics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetObjectTypeAttributeStatisticsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_object_type_attribute_statistics, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetObjectTypeAttributeStatistics(ctx, input)
			},
		},
		"get-profile-history-record": {
			Name:   "get-profile-history-record",
			Fields: fields_get_profile_history_record,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetProfileHistoryRecordInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_profile_history_record, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetProfileHistoryRecord(ctx, input)
			},
		},
		"get-profile-object-type": {
			Name:   "get-profile-object-type",
			Fields: fields_get_profile_object_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetProfileObjectTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_profile_object_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetProfileObjectType(ctx, input)
			},
		},
		"get-profile-object-type-template": {
			Name:   "get-profile-object-type-template",
			Fields: fields_get_profile_object_type_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetProfileObjectTypeTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_profile_object_type_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetProfileObjectTypeTemplate(ctx, input)
			},
		},
		"get-profile-recommendations": {
			Name:   "get-profile-recommendations",
			Fields: fields_get_profile_recommendations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetProfileRecommendationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_profile_recommendations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetProfileRecommendations(ctx, input)
			},
		},
		"get-recommender": {
			Name:   "get-recommender",
			Fields: fields_get_recommender,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRecommenderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_recommender, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRecommender(ctx, input)
			},
		},
		"get-segment-definition": {
			Name:   "get-segment-definition",
			Fields: fields_get_segment_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSegmentDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_segment_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSegmentDefinition(ctx, input)
			},
		},
		"get-segment-estimate": {
			Name:   "get-segment-estimate",
			Fields: fields_get_segment_estimate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSegmentEstimateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_segment_estimate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSegmentEstimate(ctx, input)
			},
		},
		"get-segment-membership": {
			Name:   "get-segment-membership",
			Fields: fields_get_segment_membership,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSegmentMembershipInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_segment_membership, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSegmentMembership(ctx, input)
			},
		},
		"get-segment-snapshot": {
			Name:   "get-segment-snapshot",
			Fields: fields_get_segment_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSegmentSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_segment_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSegmentSnapshot(ctx, input)
			},
		},
		"get-similar-profiles": {
			Name:   "get-similar-profiles",
			Fields: fields_get_similar_profiles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSimilarProfilesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_similar_profiles, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetSimilarProfiles(ctx, input)
				}
				var results []*svc.GetSimilarProfilesOutput
				p := svc.NewGetSimilarProfilesPaginator(client, input)
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
		"get-upload-job": {
			Name:   "get-upload-job",
			Fields: fields_get_upload_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetUploadJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_upload_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetUploadJob(ctx, input)
			},
		},
		"get-upload-job-path": {
			Name:   "get-upload-job-path",
			Fields: fields_get_upload_job_path,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetUploadJobPathInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_upload_job_path, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetUploadJobPath(ctx, input)
			},
		},
		"get-workflow": {
			Name:   "get-workflow",
			Fields: fields_get_workflow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWorkflowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_workflow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetWorkflow(ctx, input)
			},
		},
		"get-workflow-steps": {
			Name:   "get-workflow-steps",
			Fields: fields_get_workflow_steps,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWorkflowStepsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_workflow_steps, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetWorkflowSteps(ctx, input)
			},
		},
		"list-account-integrations": {
			Name:   "list-account-integrations",
			Fields: fields_list_account_integrations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAccountIntegrationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_account_integrations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListAccountIntegrations(ctx, input)
			},
		},
		"list-calculated-attribute-definitions": {
			Name:   "list-calculated-attribute-definitions",
			Fields: fields_list_calculated_attribute_definitions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCalculatedAttributeDefinitionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_calculated_attribute_definitions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListCalculatedAttributeDefinitions(ctx, input)
			},
		},
		"list-calculated-attributes-for-profile": {
			Name:   "list-calculated-attributes-for-profile",
			Fields: fields_list_calculated_attributes_for_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCalculatedAttributesForProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_calculated_attributes_for_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListCalculatedAttributesForProfile(ctx, input)
			},
		},
		"list-domain-layouts": {
			Name:   "list-domain-layouts",
			Fields: fields_list_domain_layouts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDomainLayoutsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_domain_layouts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDomainLayouts(ctx, input)
				}
				var results []*svc.ListDomainLayoutsOutput
				p := svc.NewListDomainLayoutsPaginator(client, input)
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
		"list-domain-object-types": {
			Name:   "list-domain-object-types",
			Fields: fields_list_domain_object_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDomainObjectTypesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_domain_object_types, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDomainObjectTypes(ctx, input)
				}
				var results []*svc.ListDomainObjectTypesOutput
				p := svc.NewListDomainObjectTypesPaginator(client, input)
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
				if _, err := leanruntime.ApplyInput(input, fields_list_domains, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListDomains(ctx, input)
			},
		},
		"list-event-streams": {
			Name:   "list-event-streams",
			Fields: fields_list_event_streams,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEventStreamsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_event_streams, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEventStreams(ctx, input)
				}
				var results []*svc.ListEventStreamsOutput
				p := svc.NewListEventStreamsPaginator(client, input)
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
		"list-event-triggers": {
			Name:   "list-event-triggers",
			Fields: fields_list_event_triggers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEventTriggersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_event_triggers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEventTriggers(ctx, input)
				}
				var results []*svc.ListEventTriggersOutput
				p := svc.NewListEventTriggersPaginator(client, input)
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
		"list-identity-resolution-jobs": {
			Name:   "list-identity-resolution-jobs",
			Fields: fields_list_identity_resolution_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIdentityResolutionJobsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_identity_resolution_jobs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListIdentityResolutionJobs(ctx, input)
			},
		},
		"list-integrations": {
			Name:   "list-integrations",
			Fields: fields_list_integrations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIntegrationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_integrations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListIntegrations(ctx, input)
			},
		},
		"list-object-type-attribute-values": {
			Name:   "list-object-type-attribute-values",
			Fields: fields_list_object_type_attribute_values,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListObjectTypeAttributeValuesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_object_type_attribute_values, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListObjectTypeAttributeValues(ctx, input)
			},
		},
		"list-object-type-attributes": {
			Name:   "list-object-type-attributes",
			Fields: fields_list_object_type_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListObjectTypeAttributesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_object_type_attributes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListObjectTypeAttributes(ctx, input)
				}
				var results []*svc.ListObjectTypeAttributesOutput
				p := svc.NewListObjectTypeAttributesPaginator(client, input)
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
		"list-profile-attribute-values": {
			Name:   "list-profile-attribute-values",
			Fields: fields_list_profile_attribute_values,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProfileAttributeValuesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_profile_attribute_values, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListProfileAttributeValues(ctx, input)
			},
		},
		"list-profile-history-records": {
			Name:   "list-profile-history-records",
			Fields: fields_list_profile_history_records,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProfileHistoryRecordsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_profile_history_records, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListProfileHistoryRecords(ctx, input)
			},
		},
		"list-profile-object-type-templates": {
			Name:   "list-profile-object-type-templates",
			Fields: fields_list_profile_object_type_templates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProfileObjectTypeTemplatesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_profile_object_type_templates, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListProfileObjectTypeTemplates(ctx, input)
			},
		},
		"list-profile-object-types": {
			Name:   "list-profile-object-types",
			Fields: fields_list_profile_object_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProfileObjectTypesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_profile_object_types, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListProfileObjectTypes(ctx, input)
			},
		},
		"list-profile-objects": {
			Name:   "list-profile-objects",
			Fields: fields_list_profile_objects,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProfileObjectsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_profile_objects, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListProfileObjects(ctx, input)
			},
		},
		"list-recommender-recipes": {
			Name:   "list-recommender-recipes",
			Fields: fields_list_recommender_recipes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRecommenderRecipesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_recommender_recipes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRecommenderRecipes(ctx, input)
				}
				var results []*svc.ListRecommenderRecipesOutput
				p := svc.NewListRecommenderRecipesPaginator(client, input)
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
		"list-recommenders": {
			Name:   "list-recommenders",
			Fields: fields_list_recommenders,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRecommendersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_recommenders, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRecommenders(ctx, input)
				}
				var results []*svc.ListRecommendersOutput
				p := svc.NewListRecommendersPaginator(client, input)
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
		"list-rule-based-matches": {
			Name:   "list-rule-based-matches",
			Fields: fields_list_rule_based_matches,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRuleBasedMatchesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_rule_based_matches, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRuleBasedMatches(ctx, input)
				}
				var results []*svc.ListRuleBasedMatchesOutput
				p := svc.NewListRuleBasedMatchesPaginator(client, input)
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
		"list-segment-definitions": {
			Name:   "list-segment-definitions",
			Fields: fields_list_segment_definitions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSegmentDefinitionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_segment_definitions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSegmentDefinitions(ctx, input)
				}
				var results []*svc.ListSegmentDefinitionsOutput
				p := svc.NewListSegmentDefinitionsPaginator(client, input)
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
		"list-upload-jobs": {
			Name:   "list-upload-jobs",
			Fields: fields_list_upload_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListUploadJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_upload_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListUploadJobs(ctx, input)
				}
				var results []*svc.ListUploadJobsOutput
				p := svc.NewListUploadJobsPaginator(client, input)
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
		"list-workflows": {
			Name:   "list-workflows",
			Fields: fields_list_workflows,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWorkflowsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_workflows, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListWorkflows(ctx, input)
			},
		},
		"merge-profiles": {
			Name:   "merge-profiles",
			Fields: fields_merge_profiles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.MergeProfilesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_merge_profiles, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.MergeProfiles(ctx, input)
			},
		},
		"put-domain-object-type": {
			Name:   "put-domain-object-type",
			Fields: fields_put_domain_object_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutDomainObjectTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_domain_object_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutDomainObjectType(ctx, input)
			},
		},
		"put-integration": {
			Name:   "put-integration",
			Fields: fields_put_integration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutIntegrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_integration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutIntegration(ctx, input)
			},
		},
		"put-profile-object": {
			Name:   "put-profile-object",
			Fields: fields_put_profile_object,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutProfileObjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_profile_object, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutProfileObject(ctx, input)
			},
		},
		"put-profile-object-type": {
			Name:   "put-profile-object-type",
			Fields: fields_put_profile_object_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutProfileObjectTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_profile_object_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutProfileObjectType(ctx, input)
			},
		},
		"search-profiles": {
			Name:   "search-profiles",
			Fields: fields_search_profiles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchProfilesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_search_profiles, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SearchProfiles(ctx, input)
			},
		},
		"start-recommender": {
			Name:   "start-recommender",
			Fields: fields_start_recommender,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartRecommenderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_recommender, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartRecommender(ctx, input)
			},
		},
		"start-upload-job": {
			Name:   "start-upload-job",
			Fields: fields_start_upload_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartUploadJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_upload_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartUploadJob(ctx, input)
			},
		},
		"stop-recommender": {
			Name:   "stop-recommender",
			Fields: fields_stop_recommender,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopRecommenderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_recommender, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopRecommender(ctx, input)
			},
		},
		"stop-upload-job": {
			Name:   "stop-upload-job",
			Fields: fields_stop_upload_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopUploadJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_upload_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopUploadJob(ctx, input)
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
		"update-calculated-attribute-definition": {
			Name:   "update-calculated-attribute-definition",
			Fields: fields_update_calculated_attribute_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCalculatedAttributeDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_calculated_attribute_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCalculatedAttributeDefinition(ctx, input)
			},
		},
		"update-domain": {
			Name:   "update-domain",
			Fields: fields_update_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDomain(ctx, input)
			},
		},
		"update-domain-layout": {
			Name:   "update-domain-layout",
			Fields: fields_update_domain_layout,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDomainLayoutInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_domain_layout, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDomainLayout(ctx, input)
			},
		},
		"update-event-trigger": {
			Name:   "update-event-trigger",
			Fields: fields_update_event_trigger,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEventTriggerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_event_trigger, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEventTrigger(ctx, input)
			},
		},
		"update-profile": {
			Name:   "update-profile",
			Fields: fields_update_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateProfile(ctx, input)
			},
		},
		"update-recommender": {
			Name:   "update-recommender",
			Fields: fields_update_recommender,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRecommenderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_recommender, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRecommender(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("customerprofiles", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
