package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/iotsitewise"
)

var fields_associate_assets = []leanruntime.Field{
	{Name: "AssetId", Flag: "asset-id", Type: "*string", Required: true},
	{Name: "ChildAssetId", Flag: "child-asset-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "HierarchyId", Flag: "hierarchy-id", Type: "*string", Required: true},
}

var fields_associate_time_series_to_asset_property = []leanruntime.Field{
	{Name: "Alias", Flag: "alias", Type: "*string", Required: true},
	{Name: "AssetId", Flag: "asset-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "PropertyId", Flag: "property-id", Type: "*string", Required: true},
}

var fields_batch_associate_project_assets = []leanruntime.Field{
	{Name: "AssetIds", Flag: "asset-ids", Type: "[]string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ProjectId", Flag: "project-id", Type: "*string", Required: true},
}

var fields_batch_disassociate_project_assets = []leanruntime.Field{
	{Name: "AssetIds", Flag: "asset-ids", Type: "[]string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ProjectId", Flag: "project-id", Type: "*string", Required: true},
}

var fields_batch_get_asset_property_aggregates = []leanruntime.Field{
	{Name: "Entries", Flag: "entries", Type: "[]types.BatchGetAssetPropertyAggregatesEntry", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_batch_get_asset_property_value = []leanruntime.Field{
	{Name: "Entries", Flag: "entries", Type: "[]types.BatchGetAssetPropertyValueEntry", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_batch_get_asset_property_value_history = []leanruntime.Field{
	{Name: "Entries", Flag: "entries", Type: "[]types.BatchGetAssetPropertyValueHistoryEntry", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_batch_put_asset_property_value = []leanruntime.Field{
	{Name: "EnablePartialEntryProcessing", Flag: "enable-partial-entry-processing", Type: "*bool", Required: false},
	{Name: "Entries", Flag: "entries", Type: "[]types.PutAssetPropertyValueEntry", Required: true},
}

var fields_create_access_policy = []leanruntime.Field{
	{Name: "AccessPolicyIdentity", Flag: "access-policy-identity", Type: "*types.Identity", Required: true},
	{Name: "AccessPolicyPermission", Flag: "access-policy-permission", Type: "types.Permission", Required: true},
	{Name: "AccessPolicyResource", Flag: "access-policy-resource", Type: "*types.Resource", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_asset = []leanruntime.Field{
	{Name: "AssetDescription", Flag: "asset-description", Type: "*string", Required: false},
	{Name: "AssetExternalId", Flag: "asset-external-id", Type: "*string", Required: false},
	{Name: "AssetId", Flag: "asset-id", Type: "*string", Required: false},
	{Name: "AssetModelId", Flag: "asset-model-id", Type: "*string", Required: true},
	{Name: "AssetName", Flag: "asset-name", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_asset_model = []leanruntime.Field{
	{Name: "AssetModelCompositeModels", Flag: "asset-model-composite-models", Type: "[]types.AssetModelCompositeModelDefinition", Required: false},
	{Name: "AssetModelDescription", Flag: "asset-model-description", Type: "*string", Required: false},
	{Name: "AssetModelExternalId", Flag: "asset-model-external-id", Type: "*string", Required: false},
	{Name: "AssetModelHierarchies", Flag: "asset-model-hierarchies", Type: "[]types.AssetModelHierarchyDefinition", Required: false},
	{Name: "AssetModelId", Flag: "asset-model-id", Type: "*string", Required: false},
	{Name: "AssetModelName", Flag: "asset-model-name", Type: "*string", Required: true},
	{Name: "AssetModelProperties", Flag: "asset-model-properties", Type: "[]types.AssetModelPropertyDefinition", Required: false},
	{Name: "AssetModelType", Flag: "asset-model-type", Type: "types.AssetModelType", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_asset_model_composite_model = []leanruntime.Field{
	{Name: "AssetModelCompositeModelDescription", Flag: "asset-model-composite-model-description", Type: "*string", Required: false},
	{Name: "AssetModelCompositeModelExternalId", Flag: "asset-model-composite-model-external-id", Type: "*string", Required: false},
	{Name: "AssetModelCompositeModelId", Flag: "asset-model-composite-model-id", Type: "*string", Required: false},
	{Name: "AssetModelCompositeModelName", Flag: "asset-model-composite-model-name", Type: "*string", Required: true},
	{Name: "AssetModelCompositeModelProperties", Flag: "asset-model-composite-model-properties", Type: "[]types.AssetModelPropertyDefinition", Required: false},
	{Name: "AssetModelCompositeModelType", Flag: "asset-model-composite-model-type", Type: "*string", Required: true},
	{Name: "AssetModelId", Flag: "asset-model-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ComposedAssetModelId", Flag: "composed-asset-model-id", Type: "*string", Required: false},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: false},
	{Name: "IfNoneMatch", Flag: "if-none-match", Type: "*string", Required: false},
	{Name: "MatchForVersionType", Flag: "match-for-version-type", Type: "types.AssetModelVersionType", Required: false},
	{Name: "ParentAssetModelCompositeModelId", Flag: "parent-asset-model-composite-model-id", Type: "*string", Required: false},
}

var fields_create_bulk_import_job = []leanruntime.Field{
	{Name: "AdaptiveIngestion", Flag: "adaptive-ingestion", Type: "*bool", Required: false},
	{Name: "DeleteFilesAfterImport", Flag: "delete-files-after-import", Type: "*bool", Required: false},
	{Name: "ErrorReportLocation", Flag: "error-report-location", Type: "*types.ErrorReportLocation", Required: true},
	{Name: "Files", Flag: "files", Type: "[]types.File", Required: true},
	{Name: "JobConfiguration", Flag: "job-configuration", Type: "*types.JobConfiguration", Required: true},
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: true},
	{Name: "JobRoleArn", Flag: "job-role-arn", Type: "*string", Required: true},
}

var fields_create_computation_model = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ComputationModelConfiguration", Flag: "computation-model-configuration", Type: "*types.ComputationModelConfiguration", Required: true},
	{Name: "ComputationModelDataBinding", Flag: "computation-model-data-binding", Type: "map[string]types.ComputationModelDataBindingValue", Required: true},
	{Name: "ComputationModelDescription", Flag: "computation-model-description", Type: "*string", Required: false},
	{Name: "ComputationModelName", Flag: "computation-model-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_dashboard = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DashboardDefinition", Flag: "dashboard-definition", Type: "*string", Required: true},
	{Name: "DashboardDescription", Flag: "dashboard-description", Type: "*string", Required: false},
	{Name: "DashboardName", Flag: "dashboard-name", Type: "*string", Required: true},
	{Name: "ProjectId", Flag: "project-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_dataset = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DatasetDescription", Flag: "dataset-description", Type: "*string", Required: false},
	{Name: "DatasetId", Flag: "dataset-id", Type: "*string", Required: false},
	{Name: "DatasetName", Flag: "dataset-name", Type: "*string", Required: true},
	{Name: "DatasetSource", Flag: "dataset-source", Type: "*types.DatasetSource", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_gateway = []leanruntime.Field{
	{Name: "GatewayName", Flag: "gateway-name", Type: "*string", Required: true},
	{Name: "GatewayPlatform", Flag: "gateway-platform", Type: "*types.GatewayPlatform", Required: true},
	{Name: "GatewayVersion", Flag: "gateway-version", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_portal = []leanruntime.Field{
	{Name: "Alarms", Flag: "alarms", Type: "*types.Alarms", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "NotificationSenderEmail", Flag: "notification-sender-email", Type: "*string", Required: false},
	{Name: "PortalAuthMode", Flag: "portal-auth-mode", Type: "types.AuthMode", Required: false},
	{Name: "PortalContactEmail", Flag: "portal-contact-email", Type: "*string", Required: true},
	{Name: "PortalDescription", Flag: "portal-description", Type: "*string", Required: false},
	{Name: "PortalLogoImageFile", Flag: "portal-logo-image-file", Type: "*types.ImageFile", Required: false},
	{Name: "PortalName", Flag: "portal-name", Type: "*string", Required: true},
	{Name: "PortalType", Flag: "portal-type", Type: "types.PortalType", Required: false},
	{Name: "PortalTypeConfiguration", Flag: "portal-type-configuration", Type: "map[string]types.PortalTypeEntry", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_project = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "PortalId", Flag: "portal-id", Type: "*string", Required: true},
	{Name: "ProjectDescription", Flag: "project-description", Type: "*string", Required: false},
	{Name: "ProjectName", Flag: "project-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_access_policy = []leanruntime.Field{
	{Name: "AccessPolicyId", Flag: "access-policy-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
}

var fields_delete_asset = []leanruntime.Field{
	{Name: "AssetId", Flag: "asset-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
}

var fields_delete_asset_model = []leanruntime.Field{
	{Name: "AssetModelId", Flag: "asset-model-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: false},
	{Name: "IfNoneMatch", Flag: "if-none-match", Type: "*string", Required: false},
	{Name: "MatchForVersionType", Flag: "match-for-version-type", Type: "types.AssetModelVersionType", Required: false},
}

var fields_delete_asset_model_composite_model = []leanruntime.Field{
	{Name: "AssetModelCompositeModelId", Flag: "asset-model-composite-model-id", Type: "*string", Required: true},
	{Name: "AssetModelId", Flag: "asset-model-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: false},
	{Name: "IfNoneMatch", Flag: "if-none-match", Type: "*string", Required: false},
	{Name: "MatchForVersionType", Flag: "match-for-version-type", Type: "types.AssetModelVersionType", Required: false},
}

var fields_delete_asset_model_interface_relationship = []leanruntime.Field{
	{Name: "AssetModelId", Flag: "asset-model-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "InterfaceAssetModelId", Flag: "interface-asset-model-id", Type: "*string", Required: true},
}

var fields_delete_computation_model = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ComputationModelId", Flag: "computation-model-id", Type: "*string", Required: true},
}

var fields_delete_dashboard = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DashboardId", Flag: "dashboard-id", Type: "*string", Required: true},
}

var fields_delete_dataset = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DatasetId", Flag: "dataset-id", Type: "*string", Required: true},
}

var fields_delete_gateway = []leanruntime.Field{
	{Name: "GatewayId", Flag: "gateway-id", Type: "*string", Required: true},
}

var fields_delete_portal = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "PortalId", Flag: "portal-id", Type: "*string", Required: true},
}

var fields_delete_project = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ProjectId", Flag: "project-id", Type: "*string", Required: true},
}

var fields_delete_time_series = []leanruntime.Field{
	{Name: "Alias", Flag: "alias", Type: "*string", Required: false},
	{Name: "AssetId", Flag: "asset-id", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "PropertyId", Flag: "property-id", Type: "*string", Required: false},
}

var fields_describe_access_policy = []leanruntime.Field{
	{Name: "AccessPolicyId", Flag: "access-policy-id", Type: "*string", Required: true},
}

var fields_describe_action = []leanruntime.Field{
	{Name: "ActionId", Flag: "action-id", Type: "*string", Required: true},
}

var fields_describe_asset = []leanruntime.Field{
	{Name: "AssetId", Flag: "asset-id", Type: "*string", Required: true},
	{Name: "ExcludeProperties", Flag: "exclude-properties", Type: "bool", Required: false},
}

var fields_describe_asset_composite_model = []leanruntime.Field{
	{Name: "AssetCompositeModelId", Flag: "asset-composite-model-id", Type: "*string", Required: true},
	{Name: "AssetId", Flag: "asset-id", Type: "*string", Required: true},
}

var fields_describe_asset_model = []leanruntime.Field{
	{Name: "AssetModelId", Flag: "asset-model-id", Type: "*string", Required: true},
	{Name: "AssetModelVersion", Flag: "asset-model-version", Type: "*string", Required: false},
	{Name: "ExcludeProperties", Flag: "exclude-properties", Type: "bool", Required: false},
}

var fields_describe_asset_model_composite_model = []leanruntime.Field{
	{Name: "AssetModelCompositeModelId", Flag: "asset-model-composite-model-id", Type: "*string", Required: true},
	{Name: "AssetModelId", Flag: "asset-model-id", Type: "*string", Required: true},
	{Name: "AssetModelVersion", Flag: "asset-model-version", Type: "*string", Required: false},
}

var fields_describe_asset_model_interface_relationship = []leanruntime.Field{
	{Name: "AssetModelId", Flag: "asset-model-id", Type: "*string", Required: true},
	{Name: "InterfaceAssetModelId", Flag: "interface-asset-model-id", Type: "*string", Required: true},
}

var fields_describe_asset_property = []leanruntime.Field{
	{Name: "AssetId", Flag: "asset-id", Type: "*string", Required: true},
	{Name: "PropertyId", Flag: "property-id", Type: "*string", Required: true},
}

var fields_describe_bulk_import_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_describe_computation_model = []leanruntime.Field{
	{Name: "ComputationModelId", Flag: "computation-model-id", Type: "*string", Required: true},
	{Name: "ComputationModelVersion", Flag: "computation-model-version", Type: "*string", Required: false},
}

var fields_describe_computation_model_execution_summary = []leanruntime.Field{
	{Name: "ComputationModelId", Flag: "computation-model-id", Type: "*string", Required: true},
	{Name: "ResolveToResourceId", Flag: "resolve-to-resource-id", Type: "*string", Required: false},
	{Name: "ResolveToResourceType", Flag: "resolve-to-resource-type", Type: "types.ResolveToResourceType", Required: false},
}

var fields_describe_dashboard = []leanruntime.Field{
	{Name: "DashboardId", Flag: "dashboard-id", Type: "*string", Required: true},
}

var fields_describe_dataset = []leanruntime.Field{
	{Name: "DatasetId", Flag: "dataset-id", Type: "*string", Required: true},
}

var fields_describe_default_encryption_configuration = []leanruntime.Field{}

var fields_describe_execution = []leanruntime.Field{
	{Name: "ExecutionId", Flag: "execution-id", Type: "*string", Required: true},
}

var fields_describe_gateway = []leanruntime.Field{
	{Name: "GatewayId", Flag: "gateway-id", Type: "*string", Required: true},
}

var fields_describe_gateway_capability_configuration = []leanruntime.Field{
	{Name: "CapabilityNamespace", Flag: "capability-namespace", Type: "*string", Required: true},
	{Name: "GatewayId", Flag: "gateway-id", Type: "*string", Required: true},
}

var fields_describe_logging_options = []leanruntime.Field{}

var fields_describe_portal = []leanruntime.Field{
	{Name: "PortalId", Flag: "portal-id", Type: "*string", Required: true},
}

var fields_describe_project = []leanruntime.Field{
	{Name: "ProjectId", Flag: "project-id", Type: "*string", Required: true},
}

var fields_describe_storage_configuration = []leanruntime.Field{}

var fields_describe_time_series = []leanruntime.Field{
	{Name: "Alias", Flag: "alias", Type: "*string", Required: false},
	{Name: "AssetId", Flag: "asset-id", Type: "*string", Required: false},
	{Name: "PropertyId", Flag: "property-id", Type: "*string", Required: false},
}

var fields_disassociate_assets = []leanruntime.Field{
	{Name: "AssetId", Flag: "asset-id", Type: "*string", Required: true},
	{Name: "ChildAssetId", Flag: "child-asset-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "HierarchyId", Flag: "hierarchy-id", Type: "*string", Required: true},
}

var fields_disassociate_time_series_from_asset_property = []leanruntime.Field{
	{Name: "Alias", Flag: "alias", Type: "*string", Required: true},
	{Name: "AssetId", Flag: "asset-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "PropertyId", Flag: "property-id", Type: "*string", Required: true},
}

var fields_execute_action = []leanruntime.Field{
	{Name: "ActionDefinitionId", Flag: "action-definition-id", Type: "*string", Required: true},
	{Name: "ActionPayload", Flag: "action-payload", Type: "*types.ActionPayload", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ResolveTo", Flag: "resolve-to", Type: "*types.ResolveTo", Required: false},
	{Name: "TargetResource", Flag: "target-resource", Type: "*types.TargetResource", Required: true},
}

var fields_execute_query = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "QueryStatement", Flag: "query-statement", Type: "*string", Required: true},
}

var fields_get_asset_property_aggregates = []leanruntime.Field{
	{Name: "AggregateTypes", Flag: "aggregate-types", Type: "[]types.AggregateType", Required: true},
	{Name: "AssetId", Flag: "asset-id", Type: "*string", Required: false},
	{Name: "EndDate", Flag: "end-date", Type: "*time.Time", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PropertyAlias", Flag: "property-alias", Type: "*string", Required: false},
	{Name: "PropertyId", Flag: "property-id", Type: "*string", Required: false},
	{Name: "Qualities", Flag: "qualities", Type: "[]types.Quality", Required: false},
	{Name: "Resolution", Flag: "resolution", Type: "*string", Required: true},
	{Name: "StartDate", Flag: "start-date", Type: "*time.Time", Required: true},
	{Name: "TimeOrdering", Flag: "time-ordering", Type: "types.TimeOrdering", Required: false},
}

var fields_get_asset_property_value = []leanruntime.Field{
	{Name: "AssetId", Flag: "asset-id", Type: "*string", Required: false},
	{Name: "PropertyAlias", Flag: "property-alias", Type: "*string", Required: false},
	{Name: "PropertyId", Flag: "property-id", Type: "*string", Required: false},
}

var fields_get_asset_property_value_history = []leanruntime.Field{
	{Name: "AssetId", Flag: "asset-id", Type: "*string", Required: false},
	{Name: "EndDate", Flag: "end-date", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PropertyAlias", Flag: "property-alias", Type: "*string", Required: false},
	{Name: "PropertyId", Flag: "property-id", Type: "*string", Required: false},
	{Name: "Qualities", Flag: "qualities", Type: "[]types.Quality", Required: false},
	{Name: "StartDate", Flag: "start-date", Type: "*time.Time", Required: false},
	{Name: "TimeOrdering", Flag: "time-ordering", Type: "types.TimeOrdering", Required: false},
}

var fields_get_interpolated_asset_property_values = []leanruntime.Field{
	{Name: "AssetId", Flag: "asset-id", Type: "*string", Required: false},
	{Name: "EndTimeInSeconds", Flag: "end-time-in-seconds", Type: "*int64", Required: true},
	{Name: "EndTimeOffsetInNanos", Flag: "end-time-offset-in-nanos", Type: "*int32", Required: false},
	{Name: "IntervalInSeconds", Flag: "interval-in-seconds", Type: "*int64", Required: true},
	{Name: "IntervalWindowInSeconds", Flag: "interval-window-in-seconds", Type: "*int64", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PropertyAlias", Flag: "property-alias", Type: "*string", Required: false},
	{Name: "PropertyId", Flag: "property-id", Type: "*string", Required: false},
	{Name: "Quality", Flag: "quality", Type: "types.Quality", Required: true},
	{Name: "StartTimeInSeconds", Flag: "start-time-in-seconds", Type: "*int64", Required: true},
	{Name: "StartTimeOffsetInNanos", Flag: "start-time-offset-in-nanos", Type: "*int32", Required: false},
	{Name: "Type", Flag: "type", Type: "*string", Required: true},
}

var fields_invoke_assistant = []leanruntime.Field{
	{Name: "ConversationId", Flag: "conversation-id", Type: "*string", Required: false},
	{Name: "EnableTrace", Flag: "enable-trace", Type: "bool", Required: false},
	{Name: "Message", Flag: "message", Type: "*string", Required: true},
}

var fields_list_access_policies = []leanruntime.Field{
	{Name: "IamArn", Flag: "iam-arn", Type: "*string", Required: false},
	{Name: "IdentityId", Flag: "identity-id", Type: "*string", Required: false},
	{Name: "IdentityType", Flag: "identity-type", Type: "types.IdentityType", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: false},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.ResourceType", Required: false},
}

var fields_list_actions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResolveToResourceId", Flag: "resolve-to-resource-id", Type: "*string", Required: false},
	{Name: "ResolveToResourceType", Flag: "resolve-to-resource-type", Type: "types.ResolveToResourceType", Required: false},
	{Name: "TargetResourceId", Flag: "target-resource-id", Type: "*string", Required: true},
	{Name: "TargetResourceType", Flag: "target-resource-type", Type: "types.TargetResourceType", Required: true},
}

var fields_list_asset_model_composite_models = []leanruntime.Field{
	{Name: "AssetModelId", Flag: "asset-model-id", Type: "*string", Required: true},
	{Name: "AssetModelVersion", Flag: "asset-model-version", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_asset_model_properties = []leanruntime.Field{
	{Name: "AssetModelId", Flag: "asset-model-id", Type: "*string", Required: true},
	{Name: "AssetModelVersion", Flag: "asset-model-version", Type: "*string", Required: false},
	{Name: "Filter", Flag: "filter", Type: "types.ListAssetModelPropertiesFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_asset_models = []leanruntime.Field{
	{Name: "AssetModelTypes", Flag: "asset-model-types", Type: "[]types.AssetModelType", Required: false},
	{Name: "AssetModelVersion", Flag: "asset-model-version", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_asset_properties = []leanruntime.Field{
	{Name: "AssetId", Flag: "asset-id", Type: "*string", Required: true},
	{Name: "Filter", Flag: "filter", Type: "types.ListAssetPropertiesFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_asset_relationships = []leanruntime.Field{
	{Name: "AssetId", Flag: "asset-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TraversalType", Flag: "traversal-type", Type: "types.TraversalType", Required: true},
}

var fields_list_assets = []leanruntime.Field{
	{Name: "AssetModelId", Flag: "asset-model-id", Type: "*string", Required: false},
	{Name: "Filter", Flag: "filter", Type: "types.ListAssetsFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_associated_assets = []leanruntime.Field{
	{Name: "AssetId", Flag: "asset-id", Type: "*string", Required: true},
	{Name: "HierarchyId", Flag: "hierarchy-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TraversalDirection", Flag: "traversal-direction", Type: "types.TraversalDirection", Required: false},
}

var fields_list_bulk_import_jobs = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "types.ListBulkImportJobsFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_composition_relationships = []leanruntime.Field{
	{Name: "AssetModelId", Flag: "asset-model-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_computation_model_data_binding_usages = []leanruntime.Field{
	{Name: "DataBindingValueFilter", Flag: "data-binding-value-filter", Type: "*types.DataBindingValueFilter", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_computation_model_resolve_to_resources = []leanruntime.Field{
	{Name: "ComputationModelId", Flag: "computation-model-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_computation_models = []leanruntime.Field{
	{Name: "ComputationModelType", Flag: "computation-model-type", Type: "types.ComputationModelType", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_dashboards = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProjectId", Flag: "project-id", Type: "*string", Required: true},
}

var fields_list_datasets = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SourceType", Flag: "source-type", Type: "types.DatasetSourceType", Required: true},
}

var fields_list_executions = []leanruntime.Field{
	{Name: "ActionType", Flag: "action-type", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResolveToResourceId", Flag: "resolve-to-resource-id", Type: "*string", Required: false},
	{Name: "ResolveToResourceType", Flag: "resolve-to-resource-type", Type: "types.ResolveToResourceType", Required: false},
	{Name: "TargetResourceId", Flag: "target-resource-id", Type: "*string", Required: true},
	{Name: "TargetResourceType", Flag: "target-resource-type", Type: "types.TargetResourceType", Required: true},
}

var fields_list_gateways = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_interface_relationships = []leanruntime.Field{
	{Name: "InterfaceAssetModelId", Flag: "interface-asset-model-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_portals = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_project_assets = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProjectId", Flag: "project-id", Type: "*string", Required: true},
}

var fields_list_projects = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PortalId", Flag: "portal-id", Type: "*string", Required: true},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_time_series = []leanruntime.Field{
	{Name: "AliasPrefix", Flag: "alias-prefix", Type: "*string", Required: false},
	{Name: "AssetId", Flag: "asset-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TimeSeriesType", Flag: "time-series-type", Type: "types.ListTimeSeriesType", Required: false},
}

var fields_put_asset_model_interface_relationship = []leanruntime.Field{
	{Name: "AssetModelId", Flag: "asset-model-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "InterfaceAssetModelId", Flag: "interface-asset-model-id", Type: "*string", Required: true},
	{Name: "PropertyMappingConfiguration", Flag: "property-mapping-configuration", Type: "*types.PropertyMappingConfiguration", Required: true},
}

var fields_put_default_encryption_configuration = []leanruntime.Field{
	{Name: "EncryptionType", Flag: "encryption-type", Type: "types.EncryptionType", Required: true},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
}

var fields_put_logging_options = []leanruntime.Field{
	{Name: "LoggingOptions", Flag: "logging-options", Type: "*types.LoggingOptions", Required: true},
}

var fields_put_storage_configuration = []leanruntime.Field{
	{Name: "DisallowIngestNullNaN", Flag: "disallow-ingest-null-na-n", Type: "*bool", Required: false},
	{Name: "DisassociatedDataStorage", Flag: "disassociated-data-storage", Type: "types.DisassociatedDataStorageState", Required: false},
	{Name: "MultiLayerStorage", Flag: "multi-layer-storage", Type: "*types.MultiLayerStorage", Required: false},
	{Name: "RetentionPeriod", Flag: "retention-period", Type: "*types.RetentionPeriod", Required: false},
	{Name: "StorageType", Flag: "storage-type", Type: "types.StorageType", Required: true},
	{Name: "WarmTier", Flag: "warm-tier", Type: "types.WarmTierState", Required: false},
	{Name: "WarmTierRetentionPeriod", Flag: "warm-tier-retention-period", Type: "*types.WarmTierRetentionPeriod", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_access_policy = []leanruntime.Field{
	{Name: "AccessPolicyId", Flag: "access-policy-id", Type: "*string", Required: true},
	{Name: "AccessPolicyIdentity", Flag: "access-policy-identity", Type: "*types.Identity", Required: true},
	{Name: "AccessPolicyPermission", Flag: "access-policy-permission", Type: "types.Permission", Required: true},
	{Name: "AccessPolicyResource", Flag: "access-policy-resource", Type: "*types.Resource", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
}

var fields_update_asset = []leanruntime.Field{
	{Name: "AssetDescription", Flag: "asset-description", Type: "*string", Required: false},
	{Name: "AssetExternalId", Flag: "asset-external-id", Type: "*string", Required: false},
	{Name: "AssetId", Flag: "asset-id", Type: "*string", Required: true},
	{Name: "AssetName", Flag: "asset-name", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
}

var fields_update_asset_model = []leanruntime.Field{
	{Name: "AssetModelCompositeModels", Flag: "asset-model-composite-models", Type: "[]types.AssetModelCompositeModel", Required: false},
	{Name: "AssetModelDescription", Flag: "asset-model-description", Type: "*string", Required: false},
	{Name: "AssetModelExternalId", Flag: "asset-model-external-id", Type: "*string", Required: false},
	{Name: "AssetModelHierarchies", Flag: "asset-model-hierarchies", Type: "[]types.AssetModelHierarchy", Required: false},
	{Name: "AssetModelId", Flag: "asset-model-id", Type: "*string", Required: true},
	{Name: "AssetModelName", Flag: "asset-model-name", Type: "*string", Required: true},
	{Name: "AssetModelProperties", Flag: "asset-model-properties", Type: "[]types.AssetModelProperty", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: false},
	{Name: "IfNoneMatch", Flag: "if-none-match", Type: "*string", Required: false},
	{Name: "MatchForVersionType", Flag: "match-for-version-type", Type: "types.AssetModelVersionType", Required: false},
}

var fields_update_asset_model_composite_model = []leanruntime.Field{
	{Name: "AssetModelCompositeModelDescription", Flag: "asset-model-composite-model-description", Type: "*string", Required: false},
	{Name: "AssetModelCompositeModelExternalId", Flag: "asset-model-composite-model-external-id", Type: "*string", Required: false},
	{Name: "AssetModelCompositeModelId", Flag: "asset-model-composite-model-id", Type: "*string", Required: true},
	{Name: "AssetModelCompositeModelName", Flag: "asset-model-composite-model-name", Type: "*string", Required: true},
	{Name: "AssetModelCompositeModelProperties", Flag: "asset-model-composite-model-properties", Type: "[]types.AssetModelProperty", Required: false},
	{Name: "AssetModelId", Flag: "asset-model-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "IfMatch", Flag: "if-match", Type: "*string", Required: false},
	{Name: "IfNoneMatch", Flag: "if-none-match", Type: "*string", Required: false},
	{Name: "MatchForVersionType", Flag: "match-for-version-type", Type: "types.AssetModelVersionType", Required: false},
}

var fields_update_asset_property = []leanruntime.Field{
	{Name: "AssetId", Flag: "asset-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "PropertyAlias", Flag: "property-alias", Type: "*string", Required: false},
	{Name: "PropertyId", Flag: "property-id", Type: "*string", Required: true},
	{Name: "PropertyNotificationState", Flag: "property-notification-state", Type: "types.PropertyNotificationState", Required: false},
	{Name: "PropertyUnit", Flag: "property-unit", Type: "*string", Required: false},
}

var fields_update_computation_model = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ComputationModelConfiguration", Flag: "computation-model-configuration", Type: "*types.ComputationModelConfiguration", Required: true},
	{Name: "ComputationModelDataBinding", Flag: "computation-model-data-binding", Type: "map[string]types.ComputationModelDataBindingValue", Required: true},
	{Name: "ComputationModelDescription", Flag: "computation-model-description", Type: "*string", Required: false},
	{Name: "ComputationModelId", Flag: "computation-model-id", Type: "*string", Required: true},
	{Name: "ComputationModelName", Flag: "computation-model-name", Type: "*string", Required: true},
}

var fields_update_dashboard = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DashboardDefinition", Flag: "dashboard-definition", Type: "*string", Required: true},
	{Name: "DashboardDescription", Flag: "dashboard-description", Type: "*string", Required: false},
	{Name: "DashboardId", Flag: "dashboard-id", Type: "*string", Required: true},
	{Name: "DashboardName", Flag: "dashboard-name", Type: "*string", Required: true},
}

var fields_update_dataset = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DatasetDescription", Flag: "dataset-description", Type: "*string", Required: false},
	{Name: "DatasetId", Flag: "dataset-id", Type: "*string", Required: true},
	{Name: "DatasetName", Flag: "dataset-name", Type: "*string", Required: true},
	{Name: "DatasetSource", Flag: "dataset-source", Type: "*types.DatasetSource", Required: true},
}

var fields_update_gateway = []leanruntime.Field{
	{Name: "GatewayId", Flag: "gateway-id", Type: "*string", Required: true},
	{Name: "GatewayName", Flag: "gateway-name", Type: "*string", Required: true},
}

var fields_update_gateway_capability_configuration = []leanruntime.Field{
	{Name: "CapabilityConfiguration", Flag: "capability-configuration", Type: "*string", Required: true},
	{Name: "CapabilityNamespace", Flag: "capability-namespace", Type: "*string", Required: true},
	{Name: "GatewayId", Flag: "gateway-id", Type: "*string", Required: true},
}

var fields_update_portal = []leanruntime.Field{
	{Name: "Alarms", Flag: "alarms", Type: "*types.Alarms", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "NotificationSenderEmail", Flag: "notification-sender-email", Type: "*string", Required: false},
	{Name: "PortalContactEmail", Flag: "portal-contact-email", Type: "*string", Required: true},
	{Name: "PortalDescription", Flag: "portal-description", Type: "*string", Required: false},
	{Name: "PortalId", Flag: "portal-id", Type: "*string", Required: true},
	{Name: "PortalLogoImage", Flag: "portal-logo-image", Type: "*types.Image", Required: false},
	{Name: "PortalName", Flag: "portal-name", Type: "*string", Required: true},
	{Name: "PortalType", Flag: "portal-type", Type: "types.PortalType", Required: false},
	{Name: "PortalTypeConfiguration", Flag: "portal-type-configuration", Type: "map[string]types.PortalTypeEntry", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
}

var fields_update_project = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ProjectDescription", Flag: "project-description", Type: "*string", Required: false},
	{Name: "ProjectId", Flag: "project-id", Type: "*string", Required: true},
	{Name: "ProjectName", Flag: "project-name", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-assets": {
			Name:   "associate-assets",
			Fields: fields_associate_assets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateAssetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_assets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateAssets(ctx, input)
			},
		},
		"associate-time-series-to-asset-property": {
			Name:   "associate-time-series-to-asset-property",
			Fields: fields_associate_time_series_to_asset_property,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateTimeSeriesToAssetPropertyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_time_series_to_asset_property, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateTimeSeriesToAssetProperty(ctx, input)
			},
		},
		"batch-associate-project-assets": {
			Name:   "batch-associate-project-assets",
			Fields: fields_batch_associate_project_assets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchAssociateProjectAssetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_associate_project_assets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchAssociateProjectAssets(ctx, input)
			},
		},
		"batch-disassociate-project-assets": {
			Name:   "batch-disassociate-project-assets",
			Fields: fields_batch_disassociate_project_assets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDisassociateProjectAssetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_disassociate_project_assets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDisassociateProjectAssets(ctx, input)
			},
		},
		"batch-get-asset-property-aggregates": {
			Name:   "batch-get-asset-property-aggregates",
			Fields: fields_batch_get_asset_property_aggregates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetAssetPropertyAggregatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_batch_get_asset_property_aggregates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.BatchGetAssetPropertyAggregates(ctx, input)
				}
				var results []*svc.BatchGetAssetPropertyAggregatesOutput
				p := svc.NewBatchGetAssetPropertyAggregatesPaginator(client, input)
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
		"batch-get-asset-property-value": {
			Name:   "batch-get-asset-property-value",
			Fields: fields_batch_get_asset_property_value,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetAssetPropertyValueInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_batch_get_asset_property_value, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.BatchGetAssetPropertyValue(ctx, input)
				}
				var results []*svc.BatchGetAssetPropertyValueOutput
				p := svc.NewBatchGetAssetPropertyValuePaginator(client, input)
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
		"batch-get-asset-property-value-history": {
			Name:   "batch-get-asset-property-value-history",
			Fields: fields_batch_get_asset_property_value_history,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetAssetPropertyValueHistoryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_batch_get_asset_property_value_history, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.BatchGetAssetPropertyValueHistory(ctx, input)
				}
				var results []*svc.BatchGetAssetPropertyValueHistoryOutput
				p := svc.NewBatchGetAssetPropertyValueHistoryPaginator(client, input)
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
		"batch-put-asset-property-value": {
			Name:   "batch-put-asset-property-value",
			Fields: fields_batch_put_asset_property_value,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchPutAssetPropertyValueInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_put_asset_property_value, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchPutAssetPropertyValue(ctx, input)
			},
		},
		"create-access-policy": {
			Name:   "create-access-policy",
			Fields: fields_create_access_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAccessPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_access_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAccessPolicy(ctx, input)
			},
		},
		"create-asset": {
			Name:   "create-asset",
			Fields: fields_create_asset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAssetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_asset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAsset(ctx, input)
			},
		},
		"create-asset-model": {
			Name:   "create-asset-model",
			Fields: fields_create_asset_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAssetModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_asset_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAssetModel(ctx, input)
			},
		},
		"create-asset-model-composite-model": {
			Name:   "create-asset-model-composite-model",
			Fields: fields_create_asset_model_composite_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAssetModelCompositeModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_asset_model_composite_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAssetModelCompositeModel(ctx, input)
			},
		},
		"create-bulk-import-job": {
			Name:   "create-bulk-import-job",
			Fields: fields_create_bulk_import_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBulkImportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_bulk_import_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBulkImportJob(ctx, input)
			},
		},
		"create-computation-model": {
			Name:   "create-computation-model",
			Fields: fields_create_computation_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateComputationModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_computation_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateComputationModel(ctx, input)
			},
		},
		"create-dashboard": {
			Name:   "create-dashboard",
			Fields: fields_create_dashboard,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDashboardInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_dashboard, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDashboard(ctx, input)
			},
		},
		"create-dataset": {
			Name:   "create-dataset",
			Fields: fields_create_dataset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDatasetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_dataset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDataset(ctx, input)
			},
		},
		"create-gateway": {
			Name:   "create-gateway",
			Fields: fields_create_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateGateway(ctx, input)
			},
		},
		"create-portal": {
			Name:   "create-portal",
			Fields: fields_create_portal,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePortalInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_portal, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePortal(ctx, input)
			},
		},
		"create-project": {
			Name:   "create-project",
			Fields: fields_create_project,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateProjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_project, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateProject(ctx, input)
			},
		},
		"delete-access-policy": {
			Name:   "delete-access-policy",
			Fields: fields_delete_access_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAccessPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_access_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAccessPolicy(ctx, input)
			},
		},
		"delete-asset": {
			Name:   "delete-asset",
			Fields: fields_delete_asset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAssetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_asset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAsset(ctx, input)
			},
		},
		"delete-asset-model": {
			Name:   "delete-asset-model",
			Fields: fields_delete_asset_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAssetModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_asset_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAssetModel(ctx, input)
			},
		},
		"delete-asset-model-composite-model": {
			Name:   "delete-asset-model-composite-model",
			Fields: fields_delete_asset_model_composite_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAssetModelCompositeModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_asset_model_composite_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAssetModelCompositeModel(ctx, input)
			},
		},
		"delete-asset-model-interface-relationship": {
			Name:   "delete-asset-model-interface-relationship",
			Fields: fields_delete_asset_model_interface_relationship,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAssetModelInterfaceRelationshipInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_asset_model_interface_relationship, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAssetModelInterfaceRelationship(ctx, input)
			},
		},
		"delete-computation-model": {
			Name:   "delete-computation-model",
			Fields: fields_delete_computation_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteComputationModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_computation_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteComputationModel(ctx, input)
			},
		},
		"delete-dashboard": {
			Name:   "delete-dashboard",
			Fields: fields_delete_dashboard,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDashboardInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_dashboard, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDashboard(ctx, input)
			},
		},
		"delete-dataset": {
			Name:   "delete-dataset",
			Fields: fields_delete_dataset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDatasetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_dataset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDataset(ctx, input)
			},
		},
		"delete-gateway": {
			Name:   "delete-gateway",
			Fields: fields_delete_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteGateway(ctx, input)
			},
		},
		"delete-portal": {
			Name:   "delete-portal",
			Fields: fields_delete_portal,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePortalInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_portal, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePortal(ctx, input)
			},
		},
		"delete-project": {
			Name:   "delete-project",
			Fields: fields_delete_project,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteProjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_project, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteProject(ctx, input)
			},
		},
		"delete-time-series": {
			Name:   "delete-time-series",
			Fields: fields_delete_time_series,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTimeSeriesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_time_series, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTimeSeries(ctx, input)
			},
		},
		"describe-access-policy": {
			Name:   "describe-access-policy",
			Fields: fields_describe_access_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAccessPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_access_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAccessPolicy(ctx, input)
			},
		},
		"describe-action": {
			Name:   "describe-action",
			Fields: fields_describe_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAction(ctx, input)
			},
		},
		"describe-asset": {
			Name:   "describe-asset",
			Fields: fields_describe_asset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAssetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_asset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAsset(ctx, input)
			},
		},
		"describe-asset-composite-model": {
			Name:   "describe-asset-composite-model",
			Fields: fields_describe_asset_composite_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAssetCompositeModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_asset_composite_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAssetCompositeModel(ctx, input)
			},
		},
		"describe-asset-model": {
			Name:   "describe-asset-model",
			Fields: fields_describe_asset_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAssetModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_asset_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAssetModel(ctx, input)
			},
		},
		"describe-asset-model-composite-model": {
			Name:   "describe-asset-model-composite-model",
			Fields: fields_describe_asset_model_composite_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAssetModelCompositeModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_asset_model_composite_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAssetModelCompositeModel(ctx, input)
			},
		},
		"describe-asset-model-interface-relationship": {
			Name:   "describe-asset-model-interface-relationship",
			Fields: fields_describe_asset_model_interface_relationship,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAssetModelInterfaceRelationshipInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_asset_model_interface_relationship, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAssetModelInterfaceRelationship(ctx, input)
			},
		},
		"describe-asset-property": {
			Name:   "describe-asset-property",
			Fields: fields_describe_asset_property,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAssetPropertyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_asset_property, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAssetProperty(ctx, input)
			},
		},
		"describe-bulk-import-job": {
			Name:   "describe-bulk-import-job",
			Fields: fields_describe_bulk_import_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeBulkImportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_bulk_import_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeBulkImportJob(ctx, input)
			},
		},
		"describe-computation-model": {
			Name:   "describe-computation-model",
			Fields: fields_describe_computation_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeComputationModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_computation_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeComputationModel(ctx, input)
			},
		},
		"describe-computation-model-execution-summary": {
			Name:   "describe-computation-model-execution-summary",
			Fields: fields_describe_computation_model_execution_summary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeComputationModelExecutionSummaryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_computation_model_execution_summary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeComputationModelExecutionSummary(ctx, input)
			},
		},
		"describe-dashboard": {
			Name:   "describe-dashboard",
			Fields: fields_describe_dashboard,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDashboardInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_dashboard, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDashboard(ctx, input)
			},
		},
		"describe-dataset": {
			Name:   "describe-dataset",
			Fields: fields_describe_dataset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDatasetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_dataset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDataset(ctx, input)
			},
		},
		"describe-default-encryption-configuration": {
			Name:   "describe-default-encryption-configuration",
			Fields: fields_describe_default_encryption_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDefaultEncryptionConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_default_encryption_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDefaultEncryptionConfiguration(ctx, input)
			},
		},
		"describe-execution": {
			Name:   "describe-execution",
			Fields: fields_describe_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeExecution(ctx, input)
			},
		},
		"describe-gateway": {
			Name:   "describe-gateway",
			Fields: fields_describe_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeGateway(ctx, input)
			},
		},
		"describe-gateway-capability-configuration": {
			Name:   "describe-gateway-capability-configuration",
			Fields: fields_describe_gateway_capability_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeGatewayCapabilityConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_gateway_capability_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeGatewayCapabilityConfiguration(ctx, input)
			},
		},
		"describe-logging-options": {
			Name:   "describe-logging-options",
			Fields: fields_describe_logging_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLoggingOptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_logging_options, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeLoggingOptions(ctx, input)
			},
		},
		"describe-portal": {
			Name:   "describe-portal",
			Fields: fields_describe_portal,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePortalInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_portal, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribePortal(ctx, input)
			},
		},
		"describe-project": {
			Name:   "describe-project",
			Fields: fields_describe_project,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeProjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_project, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeProject(ctx, input)
			},
		},
		"describe-storage-configuration": {
			Name:   "describe-storage-configuration",
			Fields: fields_describe_storage_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeStorageConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_storage_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeStorageConfiguration(ctx, input)
			},
		},
		"describe-time-series": {
			Name:   "describe-time-series",
			Fields: fields_describe_time_series,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTimeSeriesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_time_series, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTimeSeries(ctx, input)
			},
		},
		"disassociate-assets": {
			Name:   "disassociate-assets",
			Fields: fields_disassociate_assets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateAssetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_assets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateAssets(ctx, input)
			},
		},
		"disassociate-time-series-from-asset-property": {
			Name:   "disassociate-time-series-from-asset-property",
			Fields: fields_disassociate_time_series_from_asset_property,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateTimeSeriesFromAssetPropertyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_time_series_from_asset_property, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateTimeSeriesFromAssetProperty(ctx, input)
			},
		},
		"execute-action": {
			Name:   "execute-action",
			Fields: fields_execute_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExecuteActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_execute_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExecuteAction(ctx, input)
			},
		},
		"execute-query": {
			Name:   "execute-query",
			Fields: fields_execute_query,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExecuteQueryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_execute_query, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ExecuteQuery(ctx, input)
				}
				var results []*svc.ExecuteQueryOutput
				p := svc.NewExecuteQueryPaginator(client, input)
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
		"get-asset-property-aggregates": {
			Name:   "get-asset-property-aggregates",
			Fields: fields_get_asset_property_aggregates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAssetPropertyAggregatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_asset_property_aggregates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetAssetPropertyAggregates(ctx, input)
				}
				var results []*svc.GetAssetPropertyAggregatesOutput
				p := svc.NewGetAssetPropertyAggregatesPaginator(client, input)
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
		"get-asset-property-value": {
			Name:   "get-asset-property-value",
			Fields: fields_get_asset_property_value,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAssetPropertyValueInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_asset_property_value, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAssetPropertyValue(ctx, input)
			},
		},
		"get-asset-property-value-history": {
			Name:   "get-asset-property-value-history",
			Fields: fields_get_asset_property_value_history,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAssetPropertyValueHistoryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_asset_property_value_history, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetAssetPropertyValueHistory(ctx, input)
				}
				var results []*svc.GetAssetPropertyValueHistoryOutput
				p := svc.NewGetAssetPropertyValueHistoryPaginator(client, input)
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
		"get-interpolated-asset-property-values": {
			Name:   "get-interpolated-asset-property-values",
			Fields: fields_get_interpolated_asset_property_values,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetInterpolatedAssetPropertyValuesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_interpolated_asset_property_values, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetInterpolatedAssetPropertyValues(ctx, input)
				}
				var results []*svc.GetInterpolatedAssetPropertyValuesOutput
				p := svc.NewGetInterpolatedAssetPropertyValuesPaginator(client, input)
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
		"invoke-assistant": {
			Name:   "invoke-assistant",
			Fields: fields_invoke_assistant,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.InvokeAssistantInput{}
				if _, err := leanruntime.ApplyInput(input, fields_invoke_assistant, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.InvokeAssistant(ctx, input)
			},
		},
		"list-access-policies": {
			Name:   "list-access-policies",
			Fields: fields_list_access_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAccessPoliciesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_access_policies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAccessPolicies(ctx, input)
				}
				var results []*svc.ListAccessPoliciesOutput
				p := svc.NewListAccessPoliciesPaginator(client, input)
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
		"list-actions": {
			Name:   "list-actions",
			Fields: fields_list_actions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListActionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_actions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListActions(ctx, input)
			},
		},
		"list-asset-model-composite-models": {
			Name:   "list-asset-model-composite-models",
			Fields: fields_list_asset_model_composite_models,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAssetModelCompositeModelsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_asset_model_composite_models, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAssetModelCompositeModels(ctx, input)
				}
				var results []*svc.ListAssetModelCompositeModelsOutput
				p := svc.NewListAssetModelCompositeModelsPaginator(client, input)
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
		"list-asset-model-properties": {
			Name:   "list-asset-model-properties",
			Fields: fields_list_asset_model_properties,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAssetModelPropertiesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_asset_model_properties, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAssetModelProperties(ctx, input)
				}
				var results []*svc.ListAssetModelPropertiesOutput
				p := svc.NewListAssetModelPropertiesPaginator(client, input)
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
		"list-asset-models": {
			Name:   "list-asset-models",
			Fields: fields_list_asset_models,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAssetModelsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_asset_models, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAssetModels(ctx, input)
				}
				var results []*svc.ListAssetModelsOutput
				p := svc.NewListAssetModelsPaginator(client, input)
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
		"list-asset-properties": {
			Name:   "list-asset-properties",
			Fields: fields_list_asset_properties,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAssetPropertiesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_asset_properties, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAssetProperties(ctx, input)
				}
				var results []*svc.ListAssetPropertiesOutput
				p := svc.NewListAssetPropertiesPaginator(client, input)
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
		"list-asset-relationships": {
			Name:   "list-asset-relationships",
			Fields: fields_list_asset_relationships,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAssetRelationshipsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_asset_relationships, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAssetRelationships(ctx, input)
				}
				var results []*svc.ListAssetRelationshipsOutput
				p := svc.NewListAssetRelationshipsPaginator(client, input)
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
		"list-assets": {
			Name:   "list-assets",
			Fields: fields_list_assets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAssetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_assets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAssets(ctx, input)
				}
				var results []*svc.ListAssetsOutput
				p := svc.NewListAssetsPaginator(client, input)
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
		"list-associated-assets": {
			Name:   "list-associated-assets",
			Fields: fields_list_associated_assets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAssociatedAssetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_associated_assets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAssociatedAssets(ctx, input)
				}
				var results []*svc.ListAssociatedAssetsOutput
				p := svc.NewListAssociatedAssetsPaginator(client, input)
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
		"list-bulk-import-jobs": {
			Name:   "list-bulk-import-jobs",
			Fields: fields_list_bulk_import_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBulkImportJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_bulk_import_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBulkImportJobs(ctx, input)
				}
				var results []*svc.ListBulkImportJobsOutput
				p := svc.NewListBulkImportJobsPaginator(client, input)
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
		"list-composition-relationships": {
			Name:   "list-composition-relationships",
			Fields: fields_list_composition_relationships,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCompositionRelationshipsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_composition_relationships, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCompositionRelationships(ctx, input)
				}
				var results []*svc.ListCompositionRelationshipsOutput
				p := svc.NewListCompositionRelationshipsPaginator(client, input)
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
		"list-computation-model-data-binding-usages": {
			Name:   "list-computation-model-data-binding-usages",
			Fields: fields_list_computation_model_data_binding_usages,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListComputationModelDataBindingUsagesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_computation_model_data_binding_usages, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListComputationModelDataBindingUsages(ctx, input)
				}
				var results []*svc.ListComputationModelDataBindingUsagesOutput
				p := svc.NewListComputationModelDataBindingUsagesPaginator(client, input)
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
		"list-computation-model-resolve-to-resources": {
			Name:   "list-computation-model-resolve-to-resources",
			Fields: fields_list_computation_model_resolve_to_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListComputationModelResolveToResourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_computation_model_resolve_to_resources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListComputationModelResolveToResources(ctx, input)
				}
				var results []*svc.ListComputationModelResolveToResourcesOutput
				p := svc.NewListComputationModelResolveToResourcesPaginator(client, input)
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
		"list-computation-models": {
			Name:   "list-computation-models",
			Fields: fields_list_computation_models,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListComputationModelsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_computation_models, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListComputationModels(ctx, input)
				}
				var results []*svc.ListComputationModelsOutput
				p := svc.NewListComputationModelsPaginator(client, input)
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
		"list-dashboards": {
			Name:   "list-dashboards",
			Fields: fields_list_dashboards,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDashboardsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_dashboards, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDashboards(ctx, input)
				}
				var results []*svc.ListDashboardsOutput
				p := svc.NewListDashboardsPaginator(client, input)
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
		"list-datasets": {
			Name:   "list-datasets",
			Fields: fields_list_datasets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDatasetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_datasets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDatasets(ctx, input)
				}
				var results []*svc.ListDatasetsOutput
				p := svc.NewListDatasetsPaginator(client, input)
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
		"list-executions": {
			Name:   "list-executions",
			Fields: fields_list_executions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListExecutionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_executions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListExecutions(ctx, input)
				}
				var results []*svc.ListExecutionsOutput
				p := svc.NewListExecutionsPaginator(client, input)
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
		"list-gateways": {
			Name:   "list-gateways",
			Fields: fields_list_gateways,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGatewaysInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_gateways, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListGateways(ctx, input)
				}
				var results []*svc.ListGatewaysOutput
				p := svc.NewListGatewaysPaginator(client, input)
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
		"list-interface-relationships": {
			Name:   "list-interface-relationships",
			Fields: fields_list_interface_relationships,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInterfaceRelationshipsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_interface_relationships, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListInterfaceRelationships(ctx, input)
				}
				var results []*svc.ListInterfaceRelationshipsOutput
				p := svc.NewListInterfaceRelationshipsPaginator(client, input)
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
		"list-portals": {
			Name:   "list-portals",
			Fields: fields_list_portals,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPortalsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_portals, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPortals(ctx, input)
				}
				var results []*svc.ListPortalsOutput
				p := svc.NewListPortalsPaginator(client, input)
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
		"list-project-assets": {
			Name:   "list-project-assets",
			Fields: fields_list_project_assets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProjectAssetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_project_assets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProjectAssets(ctx, input)
				}
				var results []*svc.ListProjectAssetsOutput
				p := svc.NewListProjectAssetsPaginator(client, input)
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
		"list-projects": {
			Name:   "list-projects",
			Fields: fields_list_projects,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProjectsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_projects, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProjects(ctx, input)
				}
				var results []*svc.ListProjectsOutput
				p := svc.NewListProjectsPaginator(client, input)
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
		"list-time-series": {
			Name:   "list-time-series",
			Fields: fields_list_time_series,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTimeSeriesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_time_series, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTimeSeries(ctx, input)
				}
				var results []*svc.ListTimeSeriesOutput
				p := svc.NewListTimeSeriesPaginator(client, input)
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
		"put-asset-model-interface-relationship": {
			Name:   "put-asset-model-interface-relationship",
			Fields: fields_put_asset_model_interface_relationship,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutAssetModelInterfaceRelationshipInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_asset_model_interface_relationship, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutAssetModelInterfaceRelationship(ctx, input)
			},
		},
		"put-default-encryption-configuration": {
			Name:   "put-default-encryption-configuration",
			Fields: fields_put_default_encryption_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutDefaultEncryptionConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_default_encryption_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutDefaultEncryptionConfiguration(ctx, input)
			},
		},
		"put-logging-options": {
			Name:   "put-logging-options",
			Fields: fields_put_logging_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutLoggingOptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_logging_options, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutLoggingOptions(ctx, input)
			},
		},
		"put-storage-configuration": {
			Name:   "put-storage-configuration",
			Fields: fields_put_storage_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutStorageConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_storage_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutStorageConfiguration(ctx, input)
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
		"update-access-policy": {
			Name:   "update-access-policy",
			Fields: fields_update_access_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAccessPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_access_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAccessPolicy(ctx, input)
			},
		},
		"update-asset": {
			Name:   "update-asset",
			Fields: fields_update_asset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAssetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_asset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAsset(ctx, input)
			},
		},
		"update-asset-model": {
			Name:   "update-asset-model",
			Fields: fields_update_asset_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAssetModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_asset_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAssetModel(ctx, input)
			},
		},
		"update-asset-model-composite-model": {
			Name:   "update-asset-model-composite-model",
			Fields: fields_update_asset_model_composite_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAssetModelCompositeModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_asset_model_composite_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAssetModelCompositeModel(ctx, input)
			},
		},
		"update-asset-property": {
			Name:   "update-asset-property",
			Fields: fields_update_asset_property,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAssetPropertyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_asset_property, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAssetProperty(ctx, input)
			},
		},
		"update-computation-model": {
			Name:   "update-computation-model",
			Fields: fields_update_computation_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateComputationModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_computation_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateComputationModel(ctx, input)
			},
		},
		"update-dashboard": {
			Name:   "update-dashboard",
			Fields: fields_update_dashboard,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDashboardInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_dashboard, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDashboard(ctx, input)
			},
		},
		"update-dataset": {
			Name:   "update-dataset",
			Fields: fields_update_dataset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDatasetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_dataset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDataset(ctx, input)
			},
		},
		"update-gateway": {
			Name:   "update-gateway",
			Fields: fields_update_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateGateway(ctx, input)
			},
		},
		"update-gateway-capability-configuration": {
			Name:   "update-gateway-capability-configuration",
			Fields: fields_update_gateway_capability_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateGatewayCapabilityConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_gateway_capability_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateGatewayCapabilityConfiguration(ctx, input)
			},
		},
		"update-portal": {
			Name:   "update-portal",
			Fields: fields_update_portal,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePortalInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_portal, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePortal(ctx, input)
			},
		},
		"update-project": {
			Name:   "update-project",
			Fields: fields_update_project,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateProjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_project, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateProject(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("iotsitewise", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
