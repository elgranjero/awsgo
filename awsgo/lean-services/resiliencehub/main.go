package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/resiliencehub"
)

var fields_accept_resource_grouping_recommendations = []leanruntime.Field{
	{Name: "AppArn", Flag: "app-arn", Type: "*string", Required: true},
	{Name: "Entries", Flag: "entries", Type: "[]types.AcceptGroupingRecommendationEntry", Required: true},
}

var fields_add_draft_app_version_resource_mappings = []leanruntime.Field{
	{Name: "AppArn", Flag: "app-arn", Type: "*string", Required: true},
	{Name: "ResourceMappings", Flag: "resource-mappings", Type: "[]types.ResourceMapping", Required: true},
}

var fields_batch_update_recommendation_status = []leanruntime.Field{
	{Name: "AppArn", Flag: "app-arn", Type: "*string", Required: true},
	{Name: "RequestEntries", Flag: "request-entries", Type: "[]types.UpdateRecommendationStatusRequestEntry", Required: true},
}

var fields_create_app = []leanruntime.Field{
	{Name: "AssessmentSchedule", Flag: "assessment-schedule", Type: "types.AppAssessmentScheduleType", Required: false},
	{Name: "AwsApplicationArn", Flag: "aws-application-arn", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EventSubscriptions", Flag: "event-subscriptions", Type: "[]types.EventSubscription", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "PermissionModel", Flag: "permission-model", Type: "*types.PermissionModel", Required: false},
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_app_version_app_component = []leanruntime.Field{
	{Name: "AdditionalInfo", Flag: "additional-info", Type: "map[string][]string", Required: false},
	{Name: "AppArn", Flag: "app-arn", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "*string", Required: true},
}

var fields_create_app_version_resource = []leanruntime.Field{
	{Name: "AdditionalInfo", Flag: "additional-info", Type: "map[string][]string", Required: false},
	{Name: "AppArn", Flag: "app-arn", Type: "*string", Required: true},
	{Name: "AppComponents", Flag: "app-components", Type: "[]string", Required: true},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: false},
	{Name: "AwsRegion", Flag: "aws-region", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "LogicalResourceId", Flag: "logical-resource-id", Type: "*types.LogicalResourceId", Required: true},
	{Name: "PhysicalResourceId", Flag: "physical-resource-id", Type: "*string", Required: true},
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: false},
	{Name: "ResourceType", Flag: "resource-type", Type: "*string", Required: true},
}

var fields_create_recommendation_template = []leanruntime.Field{
	{Name: "AssessmentArn", Flag: "assessment-arn", Type: "*string", Required: true},
	{Name: "BucketName", Flag: "bucket-name", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Format", Flag: "format", Type: "types.TemplateFormat", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RecommendationIds", Flag: "recommendation-ids", Type: "[]string", Required: false},
	{Name: "RecommendationTypes", Flag: "recommendation-types", Type: "[]types.RenderRecommendationType", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_resiliency_policy = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DataLocationConstraint", Flag: "data-location-constraint", Type: "types.DataLocationConstraint", Required: false},
	{Name: "Policy", Flag: "policy", Type: "map[string]types.FailurePolicy", Required: true},
	{Name: "PolicyDescription", Flag: "policy-description", Type: "*string", Required: false},
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Tier", Flag: "tier", Type: "types.ResiliencyPolicyTier", Required: true},
}

var fields_delete_app = []leanruntime.Field{
	{Name: "AppArn", Flag: "app-arn", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ForceDelete", Flag: "force-delete", Type: "*bool", Required: false},
}

var fields_delete_app_assessment = []leanruntime.Field{
	{Name: "AssessmentArn", Flag: "assessment-arn", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
}

var fields_delete_app_input_source = []leanruntime.Field{
	{Name: "AppArn", Flag: "app-arn", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "EksSourceClusterNamespace", Flag: "eks-source-cluster-namespace", Type: "*types.EksSourceClusterNamespace", Required: false},
	{Name: "SourceArn", Flag: "source-arn", Type: "*string", Required: false},
	{Name: "TerraformSource", Flag: "terraform-source", Type: "*types.TerraformSource", Required: false},
}

var fields_delete_app_version_app_component = []leanruntime.Field{
	{Name: "AppArn", Flag: "app-arn", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_app_version_resource = []leanruntime.Field{
	{Name: "AppArn", Flag: "app-arn", Type: "*string", Required: true},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: false},
	{Name: "AwsRegion", Flag: "aws-region", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "LogicalResourceId", Flag: "logical-resource-id", Type: "*types.LogicalResourceId", Required: false},
	{Name: "PhysicalResourceId", Flag: "physical-resource-id", Type: "*string", Required: false},
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: false},
}

var fields_delete_recommendation_template = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "RecommendationTemplateArn", Flag: "recommendation-template-arn", Type: "*string", Required: true},
}

var fields_delete_resiliency_policy = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: true},
}

var fields_describe_app = []leanruntime.Field{
	{Name: "AppArn", Flag: "app-arn", Type: "*string", Required: true},
}

var fields_describe_app_assessment = []leanruntime.Field{
	{Name: "AssessmentArn", Flag: "assessment-arn", Type: "*string", Required: true},
}

var fields_describe_app_version = []leanruntime.Field{
	{Name: "AppArn", Flag: "app-arn", Type: "*string", Required: true},
	{Name: "AppVersion", Flag: "app-version", Type: "*string", Required: true},
}

var fields_describe_app_version_app_component = []leanruntime.Field{
	{Name: "AppArn", Flag: "app-arn", Type: "*string", Required: true},
	{Name: "AppVersion", Flag: "app-version", Type: "*string", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_describe_app_version_resource = []leanruntime.Field{
	{Name: "AppArn", Flag: "app-arn", Type: "*string", Required: true},
	{Name: "AppVersion", Flag: "app-version", Type: "*string", Required: true},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: false},
	{Name: "AwsRegion", Flag: "aws-region", Type: "*string", Required: false},
	{Name: "LogicalResourceId", Flag: "logical-resource-id", Type: "*types.LogicalResourceId", Required: false},
	{Name: "PhysicalResourceId", Flag: "physical-resource-id", Type: "*string", Required: false},
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: false},
}

var fields_describe_app_version_resources_resolution_status = []leanruntime.Field{
	{Name: "AppArn", Flag: "app-arn", Type: "*string", Required: true},
	{Name: "AppVersion", Flag: "app-version", Type: "*string", Required: true},
	{Name: "ResolutionId", Flag: "resolution-id", Type: "*string", Required: false},
}

var fields_describe_app_version_template = []leanruntime.Field{
	{Name: "AppArn", Flag: "app-arn", Type: "*string", Required: true},
	{Name: "AppVersion", Flag: "app-version", Type: "*string", Required: true},
}

var fields_describe_draft_app_version_resources_import_status = []leanruntime.Field{
	{Name: "AppArn", Flag: "app-arn", Type: "*string", Required: true},
}

var fields_describe_metrics_export = []leanruntime.Field{
	{Name: "MetricsExportId", Flag: "metrics-export-id", Type: "*string", Required: true},
}

var fields_describe_resiliency_policy = []leanruntime.Field{
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: true},
}

var fields_describe_resource_grouping_recommendation_task = []leanruntime.Field{
	{Name: "AppArn", Flag: "app-arn", Type: "*string", Required: true},
	{Name: "GroupingId", Flag: "grouping-id", Type: "*string", Required: false},
}

var fields_import_resources_to_draft_app_version = []leanruntime.Field{
	{Name: "AppArn", Flag: "app-arn", Type: "*string", Required: true},
	{Name: "EksSources", Flag: "eks-sources", Type: "[]types.EksSource", Required: false},
	{Name: "ImportStrategy", Flag: "import-strategy", Type: "types.ResourceImportStrategyType", Required: false},
	{Name: "SourceArns", Flag: "source-arns", Type: "[]string", Required: false},
	{Name: "TerraformSources", Flag: "terraform-sources", Type: "[]types.TerraformSource", Required: false},
}

var fields_list_alarm_recommendations = []leanruntime.Field{
	{Name: "AssessmentArn", Flag: "assessment-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_app_assessment_compliance_drifts = []leanruntime.Field{
	{Name: "AssessmentArn", Flag: "assessment-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_app_assessment_resource_drifts = []leanruntime.Field{
	{Name: "AssessmentArn", Flag: "assessment-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_app_assessments = []leanruntime.Field{
	{Name: "AppArn", Flag: "app-arn", Type: "*string", Required: false},
	{Name: "AssessmentName", Flag: "assessment-name", Type: "*string", Required: false},
	{Name: "AssessmentStatus", Flag: "assessment-status", Type: "[]types.AssessmentStatus", Required: false},
	{Name: "ComplianceStatus", Flag: "compliance-status", Type: "types.ComplianceStatus", Required: false},
	{Name: "Invoker", Flag: "invoker", Type: "types.AssessmentInvoker", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ReverseOrder", Flag: "reverse-order", Type: "*bool", Required: false},
}

var fields_list_app_component_compliances = []leanruntime.Field{
	{Name: "AssessmentArn", Flag: "assessment-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_app_component_recommendations = []leanruntime.Field{
	{Name: "AssessmentArn", Flag: "assessment-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_app_input_sources = []leanruntime.Field{
	{Name: "AppArn", Flag: "app-arn", Type: "*string", Required: true},
	{Name: "AppVersion", Flag: "app-version", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_app_version_app_components = []leanruntime.Field{
	{Name: "AppArn", Flag: "app-arn", Type: "*string", Required: true},
	{Name: "AppVersion", Flag: "app-version", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_app_version_resource_mappings = []leanruntime.Field{
	{Name: "AppArn", Flag: "app-arn", Type: "*string", Required: true},
	{Name: "AppVersion", Flag: "app-version", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_app_version_resources = []leanruntime.Field{
	{Name: "AppArn", Flag: "app-arn", Type: "*string", Required: true},
	{Name: "AppVersion", Flag: "app-version", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResolutionId", Flag: "resolution-id", Type: "*string", Required: false},
}

var fields_list_app_versions = []leanruntime.Field{
	{Name: "AppArn", Flag: "app-arn", Type: "*string", Required: true},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
}

var fields_list_apps = []leanruntime.Field{
	{Name: "AppArn", Flag: "app-arn", Type: "*string", Required: false},
	{Name: "AwsApplicationArn", Flag: "aws-application-arn", Type: "*string", Required: false},
	{Name: "FromLastAssessmentTime", Flag: "from-last-assessment-time", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ReverseOrder", Flag: "reverse-order", Type: "*bool", Required: false},
	{Name: "ToLastAssessmentTime", Flag: "to-last-assessment-time", Type: "*time.Time", Required: false},
}

var fields_list_metrics = []leanruntime.Field{
	{Name: "Conditions", Flag: "conditions", Type: "[]types.Condition", Required: false},
	{Name: "DataSource", Flag: "data-source", Type: "*string", Required: false},
	{Name: "Fields", Flag: "fields", Type: "[]types.Field", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Sorts", Flag: "sorts", Type: "[]types.Sort", Required: false},
}

var fields_list_recommendation_templates = []leanruntime.Field{
	{Name: "AssessmentArn", Flag: "assessment-arn", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RecommendationTemplateArn", Flag: "recommendation-template-arn", Type: "*string", Required: false},
	{Name: "ReverseOrder", Flag: "reverse-order", Type: "*bool", Required: false},
	{Name: "Status", Flag: "status", Type: "[]types.RecommendationTemplateStatus", Required: false},
}

var fields_list_resiliency_policies = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: false},
}

var fields_list_resource_grouping_recommendations = []leanruntime.Field{
	{Name: "AppArn", Flag: "app-arn", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_sop_recommendations = []leanruntime.Field{
	{Name: "AssessmentArn", Flag: "assessment-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_suggested_resiliency_policies = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_test_recommendations = []leanruntime.Field{
	{Name: "AssessmentArn", Flag: "assessment-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_unsupported_app_version_resources = []leanruntime.Field{
	{Name: "AppArn", Flag: "app-arn", Type: "*string", Required: true},
	{Name: "AppVersion", Flag: "app-version", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResolutionId", Flag: "resolution-id", Type: "*string", Required: false},
}

var fields_publish_app_version = []leanruntime.Field{
	{Name: "AppArn", Flag: "app-arn", Type: "*string", Required: true},
	{Name: "VersionName", Flag: "version-name", Type: "*string", Required: false},
}

var fields_put_draft_app_version_template = []leanruntime.Field{
	{Name: "AppArn", Flag: "app-arn", Type: "*string", Required: true},
	{Name: "AppTemplateBody", Flag: "app-template-body", Type: "*string", Required: true},
}

var fields_reject_resource_grouping_recommendations = []leanruntime.Field{
	{Name: "AppArn", Flag: "app-arn", Type: "*string", Required: true},
	{Name: "Entries", Flag: "entries", Type: "[]types.RejectGroupingRecommendationEntry", Required: true},
}

var fields_remove_draft_app_version_resource_mappings = []leanruntime.Field{
	{Name: "AppArn", Flag: "app-arn", Type: "*string", Required: true},
	{Name: "AppRegistryAppNames", Flag: "app-registry-app-names", Type: "[]string", Required: false},
	{Name: "EksSourceNames", Flag: "eks-source-names", Type: "[]string", Required: false},
	{Name: "LogicalStackNames", Flag: "logical-stack-names", Type: "[]string", Required: false},
	{Name: "ResourceGroupNames", Flag: "resource-group-names", Type: "[]string", Required: false},
	{Name: "ResourceNames", Flag: "resource-names", Type: "[]string", Required: false},
	{Name: "TerraformSourceNames", Flag: "terraform-source-names", Type: "[]string", Required: false},
}

var fields_resolve_app_version_resources = []leanruntime.Field{
	{Name: "AppArn", Flag: "app-arn", Type: "*string", Required: true},
	{Name: "AppVersion", Flag: "app-version", Type: "*string", Required: true},
}

var fields_start_app_assessment = []leanruntime.Field{
	{Name: "AppArn", Flag: "app-arn", Type: "*string", Required: true},
	{Name: "AppVersion", Flag: "app-version", Type: "*string", Required: true},
	{Name: "AssessmentName", Flag: "assessment-name", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_start_metrics_export = []leanruntime.Field{
	{Name: "BucketName", Flag: "bucket-name", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
}

var fields_start_resource_grouping_recommendation_task = []leanruntime.Field{
	{Name: "AppArn", Flag: "app-arn", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_app = []leanruntime.Field{
	{Name: "AppArn", Flag: "app-arn", Type: "*string", Required: true},
	{Name: "AssessmentSchedule", Flag: "assessment-schedule", Type: "types.AppAssessmentScheduleType", Required: false},
	{Name: "ClearResiliencyPolicyArn", Flag: "clear-resiliency-policy-arn", Type: "*bool", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EventSubscriptions", Flag: "event-subscriptions", Type: "[]types.EventSubscription", Required: false},
	{Name: "PermissionModel", Flag: "permission-model", Type: "*types.PermissionModel", Required: false},
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: false},
}

var fields_update_app_version = []leanruntime.Field{
	{Name: "AdditionalInfo", Flag: "additional-info", Type: "map[string][]string", Required: false},
	{Name: "AppArn", Flag: "app-arn", Type: "*string", Required: true},
}

var fields_update_app_version_app_component = []leanruntime.Field{
	{Name: "AdditionalInfo", Flag: "additional-info", Type: "map[string][]string", Required: false},
	{Name: "AppArn", Flag: "app-arn", Type: "*string", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "*string", Required: false},
}

var fields_update_app_version_resource = []leanruntime.Field{
	{Name: "AdditionalInfo", Flag: "additional-info", Type: "map[string][]string", Required: false},
	{Name: "AppArn", Flag: "app-arn", Type: "*string", Required: true},
	{Name: "AppComponents", Flag: "app-components", Type: "[]string", Required: false},
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: false},
	{Name: "AwsRegion", Flag: "aws-region", Type: "*string", Required: false},
	{Name: "Excluded", Flag: "excluded", Type: "*bool", Required: false},
	{Name: "LogicalResourceId", Flag: "logical-resource-id", Type: "*types.LogicalResourceId", Required: false},
	{Name: "PhysicalResourceId", Flag: "physical-resource-id", Type: "*string", Required: false},
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: false},
	{Name: "ResourceType", Flag: "resource-type", Type: "*string", Required: false},
}

var fields_update_resiliency_policy = []leanruntime.Field{
	{Name: "DataLocationConstraint", Flag: "data-location-constraint", Type: "types.DataLocationConstraint", Required: false},
	{Name: "Policy", Flag: "policy", Type: "map[string]types.FailurePolicy", Required: false},
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: true},
	{Name: "PolicyDescription", Flag: "policy-description", Type: "*string", Required: false},
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: false},
	{Name: "Tier", Flag: "tier", Type: "types.ResiliencyPolicyTier", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"accept-resource-grouping-recommendations": {
			Name:   "accept-resource-grouping-recommendations",
			Fields: fields_accept_resource_grouping_recommendations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AcceptResourceGroupingRecommendationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_accept_resource_grouping_recommendations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AcceptResourceGroupingRecommendations(ctx, input)
			},
		},
		"add-draft-app-version-resource-mappings": {
			Name:   "add-draft-app-version-resource-mappings",
			Fields: fields_add_draft_app_version_resource_mappings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddDraftAppVersionResourceMappingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_draft_app_version_resource_mappings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddDraftAppVersionResourceMappings(ctx, input)
			},
		},
		"batch-update-recommendation-status": {
			Name:   "batch-update-recommendation-status",
			Fields: fields_batch_update_recommendation_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchUpdateRecommendationStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_update_recommendation_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchUpdateRecommendationStatus(ctx, input)
			},
		},
		"create-app": {
			Name:   "create-app",
			Fields: fields_create_app,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAppInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_app, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateApp(ctx, input)
			},
		},
		"create-app-version-app-component": {
			Name:   "create-app-version-app-component",
			Fields: fields_create_app_version_app_component,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAppVersionAppComponentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_app_version_app_component, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAppVersionAppComponent(ctx, input)
			},
		},
		"create-app-version-resource": {
			Name:   "create-app-version-resource",
			Fields: fields_create_app_version_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAppVersionResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_app_version_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAppVersionResource(ctx, input)
			},
		},
		"create-recommendation-template": {
			Name:   "create-recommendation-template",
			Fields: fields_create_recommendation_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRecommendationTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_recommendation_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRecommendationTemplate(ctx, input)
			},
		},
		"create-resiliency-policy": {
			Name:   "create-resiliency-policy",
			Fields: fields_create_resiliency_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateResiliencyPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_resiliency_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateResiliencyPolicy(ctx, input)
			},
		},
		"delete-app": {
			Name:   "delete-app",
			Fields: fields_delete_app,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAppInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_app, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteApp(ctx, input)
			},
		},
		"delete-app-assessment": {
			Name:   "delete-app-assessment",
			Fields: fields_delete_app_assessment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAppAssessmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_app_assessment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAppAssessment(ctx, input)
			},
		},
		"delete-app-input-source": {
			Name:   "delete-app-input-source",
			Fields: fields_delete_app_input_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAppInputSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_app_input_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAppInputSource(ctx, input)
			},
		},
		"delete-app-version-app-component": {
			Name:   "delete-app-version-app-component",
			Fields: fields_delete_app_version_app_component,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAppVersionAppComponentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_app_version_app_component, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAppVersionAppComponent(ctx, input)
			},
		},
		"delete-app-version-resource": {
			Name:   "delete-app-version-resource",
			Fields: fields_delete_app_version_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAppVersionResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_app_version_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAppVersionResource(ctx, input)
			},
		},
		"delete-recommendation-template": {
			Name:   "delete-recommendation-template",
			Fields: fields_delete_recommendation_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRecommendationTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_recommendation_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRecommendationTemplate(ctx, input)
			},
		},
		"delete-resiliency-policy": {
			Name:   "delete-resiliency-policy",
			Fields: fields_delete_resiliency_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteResiliencyPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_resiliency_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteResiliencyPolicy(ctx, input)
			},
		},
		"describe-app": {
			Name:   "describe-app",
			Fields: fields_describe_app,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAppInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_app, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeApp(ctx, input)
			},
		},
		"describe-app-assessment": {
			Name:   "describe-app-assessment",
			Fields: fields_describe_app_assessment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAppAssessmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_app_assessment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAppAssessment(ctx, input)
			},
		},
		"describe-app-version": {
			Name:   "describe-app-version",
			Fields: fields_describe_app_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAppVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_app_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAppVersion(ctx, input)
			},
		},
		"describe-app-version-app-component": {
			Name:   "describe-app-version-app-component",
			Fields: fields_describe_app_version_app_component,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAppVersionAppComponentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_app_version_app_component, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAppVersionAppComponent(ctx, input)
			},
		},
		"describe-app-version-resource": {
			Name:   "describe-app-version-resource",
			Fields: fields_describe_app_version_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAppVersionResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_app_version_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAppVersionResource(ctx, input)
			},
		},
		"describe-app-version-resources-resolution-status": {
			Name:   "describe-app-version-resources-resolution-status",
			Fields: fields_describe_app_version_resources_resolution_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAppVersionResourcesResolutionStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_app_version_resources_resolution_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAppVersionResourcesResolutionStatus(ctx, input)
			},
		},
		"describe-app-version-template": {
			Name:   "describe-app-version-template",
			Fields: fields_describe_app_version_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAppVersionTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_app_version_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAppVersionTemplate(ctx, input)
			},
		},
		"describe-draft-app-version-resources-import-status": {
			Name:   "describe-draft-app-version-resources-import-status",
			Fields: fields_describe_draft_app_version_resources_import_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDraftAppVersionResourcesImportStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_draft_app_version_resources_import_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDraftAppVersionResourcesImportStatus(ctx, input)
			},
		},
		"describe-metrics-export": {
			Name:   "describe-metrics-export",
			Fields: fields_describe_metrics_export,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMetricsExportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_metrics_export, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeMetricsExport(ctx, input)
			},
		},
		"describe-resiliency-policy": {
			Name:   "describe-resiliency-policy",
			Fields: fields_describe_resiliency_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeResiliencyPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_resiliency_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeResiliencyPolicy(ctx, input)
			},
		},
		"describe-resource-grouping-recommendation-task": {
			Name:   "describe-resource-grouping-recommendation-task",
			Fields: fields_describe_resource_grouping_recommendation_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeResourceGroupingRecommendationTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_resource_grouping_recommendation_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeResourceGroupingRecommendationTask(ctx, input)
			},
		},
		"import-resources-to-draft-app-version": {
			Name:   "import-resources-to-draft-app-version",
			Fields: fields_import_resources_to_draft_app_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportResourcesToDraftAppVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_resources_to_draft_app_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportResourcesToDraftAppVersion(ctx, input)
			},
		},
		"list-alarm-recommendations": {
			Name:   "list-alarm-recommendations",
			Fields: fields_list_alarm_recommendations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAlarmRecommendationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_alarm_recommendations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAlarmRecommendations(ctx, input)
				}
				var results []*svc.ListAlarmRecommendationsOutput
				p := svc.NewListAlarmRecommendationsPaginator(client, input)
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
		"list-app-assessment-compliance-drifts": {
			Name:   "list-app-assessment-compliance-drifts",
			Fields: fields_list_app_assessment_compliance_drifts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAppAssessmentComplianceDriftsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_app_assessment_compliance_drifts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAppAssessmentComplianceDrifts(ctx, input)
				}
				var results []*svc.ListAppAssessmentComplianceDriftsOutput
				p := svc.NewListAppAssessmentComplianceDriftsPaginator(client, input)
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
		"list-app-assessment-resource-drifts": {
			Name:   "list-app-assessment-resource-drifts",
			Fields: fields_list_app_assessment_resource_drifts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAppAssessmentResourceDriftsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_app_assessment_resource_drifts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAppAssessmentResourceDrifts(ctx, input)
				}
				var results []*svc.ListAppAssessmentResourceDriftsOutput
				p := svc.NewListAppAssessmentResourceDriftsPaginator(client, input)
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
		"list-app-assessments": {
			Name:   "list-app-assessments",
			Fields: fields_list_app_assessments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAppAssessmentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_app_assessments, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAppAssessments(ctx, input)
				}
				var results []*svc.ListAppAssessmentsOutput
				p := svc.NewListAppAssessmentsPaginator(client, input)
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
		"list-app-component-compliances": {
			Name:   "list-app-component-compliances",
			Fields: fields_list_app_component_compliances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAppComponentCompliancesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_app_component_compliances, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAppComponentCompliances(ctx, input)
				}
				var results []*svc.ListAppComponentCompliancesOutput
				p := svc.NewListAppComponentCompliancesPaginator(client, input)
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
		"list-app-component-recommendations": {
			Name:   "list-app-component-recommendations",
			Fields: fields_list_app_component_recommendations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAppComponentRecommendationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_app_component_recommendations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAppComponentRecommendations(ctx, input)
				}
				var results []*svc.ListAppComponentRecommendationsOutput
				p := svc.NewListAppComponentRecommendationsPaginator(client, input)
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
		"list-app-input-sources": {
			Name:   "list-app-input-sources",
			Fields: fields_list_app_input_sources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAppInputSourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_app_input_sources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAppInputSources(ctx, input)
				}
				var results []*svc.ListAppInputSourcesOutput
				p := svc.NewListAppInputSourcesPaginator(client, input)
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
		"list-app-version-app-components": {
			Name:   "list-app-version-app-components",
			Fields: fields_list_app_version_app_components,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAppVersionAppComponentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_app_version_app_components, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAppVersionAppComponents(ctx, input)
				}
				var results []*svc.ListAppVersionAppComponentsOutput
				p := svc.NewListAppVersionAppComponentsPaginator(client, input)
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
		"list-app-version-resource-mappings": {
			Name:   "list-app-version-resource-mappings",
			Fields: fields_list_app_version_resource_mappings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAppVersionResourceMappingsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_app_version_resource_mappings, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAppVersionResourceMappings(ctx, input)
				}
				var results []*svc.ListAppVersionResourceMappingsOutput
				p := svc.NewListAppVersionResourceMappingsPaginator(client, input)
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
		"list-app-version-resources": {
			Name:   "list-app-version-resources",
			Fields: fields_list_app_version_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAppVersionResourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_app_version_resources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAppVersionResources(ctx, input)
				}
				var results []*svc.ListAppVersionResourcesOutput
				p := svc.NewListAppVersionResourcesPaginator(client, input)
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
		"list-app-versions": {
			Name:   "list-app-versions",
			Fields: fields_list_app_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAppVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_app_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAppVersions(ctx, input)
				}
				var results []*svc.ListAppVersionsOutput
				p := svc.NewListAppVersionsPaginator(client, input)
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
		"list-apps": {
			Name:   "list-apps",
			Fields: fields_list_apps,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAppsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_apps, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListApps(ctx, input)
				}
				var results []*svc.ListAppsOutput
				p := svc.NewListAppsPaginator(client, input)
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
		"list-metrics": {
			Name:   "list-metrics",
			Fields: fields_list_metrics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMetricsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_metrics, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMetrics(ctx, input)
				}
				var results []*svc.ListMetricsOutput
				p := svc.NewListMetricsPaginator(client, input)
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
		"list-recommendation-templates": {
			Name:   "list-recommendation-templates",
			Fields: fields_list_recommendation_templates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRecommendationTemplatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_recommendation_templates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRecommendationTemplates(ctx, input)
				}
				var results []*svc.ListRecommendationTemplatesOutput
				p := svc.NewListRecommendationTemplatesPaginator(client, input)
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
		"list-resiliency-policies": {
			Name:   "list-resiliency-policies",
			Fields: fields_list_resiliency_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResiliencyPoliciesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_resiliency_policies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResiliencyPolicies(ctx, input)
				}
				var results []*svc.ListResiliencyPoliciesOutput
				p := svc.NewListResiliencyPoliciesPaginator(client, input)
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
		"list-resource-grouping-recommendations": {
			Name:   "list-resource-grouping-recommendations",
			Fields: fields_list_resource_grouping_recommendations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResourceGroupingRecommendationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_resource_grouping_recommendations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResourceGroupingRecommendations(ctx, input)
				}
				var results []*svc.ListResourceGroupingRecommendationsOutput
				p := svc.NewListResourceGroupingRecommendationsPaginator(client, input)
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
		"list-sop-recommendations": {
			Name:   "list-sop-recommendations",
			Fields: fields_list_sop_recommendations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSopRecommendationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_sop_recommendations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSopRecommendations(ctx, input)
				}
				var results []*svc.ListSopRecommendationsOutput
				p := svc.NewListSopRecommendationsPaginator(client, input)
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
		"list-suggested-resiliency-policies": {
			Name:   "list-suggested-resiliency-policies",
			Fields: fields_list_suggested_resiliency_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSuggestedResiliencyPoliciesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_suggested_resiliency_policies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSuggestedResiliencyPolicies(ctx, input)
				}
				var results []*svc.ListSuggestedResiliencyPoliciesOutput
				p := svc.NewListSuggestedResiliencyPoliciesPaginator(client, input)
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
		"list-test-recommendations": {
			Name:   "list-test-recommendations",
			Fields: fields_list_test_recommendations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTestRecommendationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_test_recommendations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTestRecommendations(ctx, input)
				}
				var results []*svc.ListTestRecommendationsOutput
				p := svc.NewListTestRecommendationsPaginator(client, input)
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
		"list-unsupported-app-version-resources": {
			Name:   "list-unsupported-app-version-resources",
			Fields: fields_list_unsupported_app_version_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListUnsupportedAppVersionResourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_unsupported_app_version_resources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListUnsupportedAppVersionResources(ctx, input)
				}
				var results []*svc.ListUnsupportedAppVersionResourcesOutput
				p := svc.NewListUnsupportedAppVersionResourcesPaginator(client, input)
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
		"publish-app-version": {
			Name:   "publish-app-version",
			Fields: fields_publish_app_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PublishAppVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_publish_app_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PublishAppVersion(ctx, input)
			},
		},
		"put-draft-app-version-template": {
			Name:   "put-draft-app-version-template",
			Fields: fields_put_draft_app_version_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutDraftAppVersionTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_draft_app_version_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutDraftAppVersionTemplate(ctx, input)
			},
		},
		"reject-resource-grouping-recommendations": {
			Name:   "reject-resource-grouping-recommendations",
			Fields: fields_reject_resource_grouping_recommendations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RejectResourceGroupingRecommendationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reject_resource_grouping_recommendations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RejectResourceGroupingRecommendations(ctx, input)
			},
		},
		"remove-draft-app-version-resource-mappings": {
			Name:   "remove-draft-app-version-resource-mappings",
			Fields: fields_remove_draft_app_version_resource_mappings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveDraftAppVersionResourceMappingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_draft_app_version_resource_mappings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveDraftAppVersionResourceMappings(ctx, input)
			},
		},
		"resolve-app-version-resources": {
			Name:   "resolve-app-version-resources",
			Fields: fields_resolve_app_version_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResolveAppVersionResourcesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_resolve_app_version_resources, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResolveAppVersionResources(ctx, input)
			},
		},
		"start-app-assessment": {
			Name:   "start-app-assessment",
			Fields: fields_start_app_assessment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartAppAssessmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_app_assessment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartAppAssessment(ctx, input)
			},
		},
		"start-metrics-export": {
			Name:   "start-metrics-export",
			Fields: fields_start_metrics_export,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartMetricsExportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_metrics_export, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartMetricsExport(ctx, input)
			},
		},
		"start-resource-grouping-recommendation-task": {
			Name:   "start-resource-grouping-recommendation-task",
			Fields: fields_start_resource_grouping_recommendation_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartResourceGroupingRecommendationTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_resource_grouping_recommendation_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartResourceGroupingRecommendationTask(ctx, input)
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
		"update-app": {
			Name:   "update-app",
			Fields: fields_update_app,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAppInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_app, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateApp(ctx, input)
			},
		},
		"update-app-version": {
			Name:   "update-app-version",
			Fields: fields_update_app_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAppVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_app_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAppVersion(ctx, input)
			},
		},
		"update-app-version-app-component": {
			Name:   "update-app-version-app-component",
			Fields: fields_update_app_version_app_component,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAppVersionAppComponentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_app_version_app_component, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAppVersionAppComponent(ctx, input)
			},
		},
		"update-app-version-resource": {
			Name:   "update-app-version-resource",
			Fields: fields_update_app_version_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAppVersionResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_app_version_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAppVersionResource(ctx, input)
			},
		},
		"update-resiliency-policy": {
			Name:   "update-resiliency-policy",
			Fields: fields_update_resiliency_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateResiliencyPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_resiliency_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateResiliencyPolicy(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("resiliencehub", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
