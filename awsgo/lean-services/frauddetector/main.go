package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/frauddetector"
)

var fields_batch_create_variable = []leanruntime.Field{
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VariableEntries", Flag: "variable-entries", Type: "[]types.VariableEntry", Required: true},
}

var fields_batch_get_variable = []leanruntime.Field{
	{Name: "Names", Flag: "names", Type: "[]string", Required: true},
}

var fields_cancel_batch_import_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_cancel_batch_prediction_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_create_batch_import_job = []leanruntime.Field{
	{Name: "EventTypeName", Flag: "event-type-name", Type: "*string", Required: true},
	{Name: "IamRoleArn", Flag: "iam-role-arn", Type: "*string", Required: true},
	{Name: "InputPath", Flag: "input-path", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "OutputPath", Flag: "output-path", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_batch_prediction_job = []leanruntime.Field{
	{Name: "DetectorName", Flag: "detector-name", Type: "*string", Required: true},
	{Name: "DetectorVersion", Flag: "detector-version", Type: "*string", Required: false},
	{Name: "EventTypeName", Flag: "event-type-name", Type: "*string", Required: true},
	{Name: "IamRoleArn", Flag: "iam-role-arn", Type: "*string", Required: true},
	{Name: "InputPath", Flag: "input-path", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "OutputPath", Flag: "output-path", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_detector_version = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "ExternalModelEndpoints", Flag: "external-model-endpoints", Type: "[]string", Required: false},
	{Name: "ModelVersions", Flag: "model-versions", Type: "[]types.ModelVersion", Required: false},
	{Name: "RuleExecutionMode", Flag: "rule-execution-mode", Type: "types.RuleExecutionMode", Required: false},
	{Name: "Rules", Flag: "rules", Type: "[]types.Rule", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_list = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Elements", Flag: "elements", Type: "[]string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VariableType", Flag: "variable-type", Type: "*string", Required: false},
}

var fields_create_model = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EventTypeName", Flag: "event-type-name", Type: "*string", Required: true},
	{Name: "ModelId", Flag: "model-id", Type: "*string", Required: true},
	{Name: "ModelType", Flag: "model-type", Type: "types.ModelTypeEnum", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_model_version = []leanruntime.Field{
	{Name: "ExternalEventsDetail", Flag: "external-events-detail", Type: "*types.ExternalEventsDetail", Required: false},
	{Name: "IngestedEventsDetail", Flag: "ingested-events-detail", Type: "*types.IngestedEventsDetail", Required: false},
	{Name: "ModelId", Flag: "model-id", Type: "*string", Required: true},
	{Name: "ModelType", Flag: "model-type", Type: "types.ModelTypeEnum", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TrainingDataSchema", Flag: "training-data-schema", Type: "*types.TrainingDataSchema", Required: true},
	{Name: "TrainingDataSource", Flag: "training-data-source", Type: "types.TrainingDataSourceEnum", Required: true},
}

var fields_create_rule = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "Expression", Flag: "expression", Type: "*string", Required: true},
	{Name: "Language", Flag: "language", Type: "types.Language", Required: true},
	{Name: "Outcomes", Flag: "outcomes", Type: "[]string", Required: true},
	{Name: "RuleId", Flag: "rule-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_variable = []leanruntime.Field{
	{Name: "DataSource", Flag: "data-source", Type: "types.DataSource", Required: true},
	{Name: "DataType", Flag: "data-type", Type: "types.DataType", Required: true},
	{Name: "DefaultValue", Flag: "default-value", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VariableType", Flag: "variable-type", Type: "*string", Required: false},
}

var fields_delete_batch_import_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_delete_batch_prediction_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_delete_detector = []leanruntime.Field{
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
}

var fields_delete_detector_version = []leanruntime.Field{
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "DetectorVersionId", Flag: "detector-version-id", Type: "*string", Required: true},
}

var fields_delete_entity_type = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_event = []leanruntime.Field{
	{Name: "DeleteAuditHistory", Flag: "delete-audit-history", Type: "*bool", Required: false},
	{Name: "EventId", Flag: "event-id", Type: "*string", Required: true},
	{Name: "EventTypeName", Flag: "event-type-name", Type: "*string", Required: true},
}

var fields_delete_event_type = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_events_by_event_type = []leanruntime.Field{
	{Name: "EventTypeName", Flag: "event-type-name", Type: "*string", Required: true},
}

var fields_delete_external_model = []leanruntime.Field{
	{Name: "ModelEndpoint", Flag: "model-endpoint", Type: "*string", Required: true},
}

var fields_delete_label = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_list = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_model = []leanruntime.Field{
	{Name: "ModelId", Flag: "model-id", Type: "*string", Required: true},
	{Name: "ModelType", Flag: "model-type", Type: "types.ModelTypeEnum", Required: true},
}

var fields_delete_model_version = []leanruntime.Field{
	{Name: "ModelId", Flag: "model-id", Type: "*string", Required: true},
	{Name: "ModelType", Flag: "model-type", Type: "types.ModelTypeEnum", Required: true},
	{Name: "ModelVersionNumber", Flag: "model-version-number", Type: "*string", Required: true},
}

var fields_delete_outcome = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_rule = []leanruntime.Field{
	{Name: "Rule", Flag: "rule", Type: "*types.Rule", Required: true},
}

var fields_delete_variable = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_describe_detector = []leanruntime.Field{
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_model_versions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "ModelId", Flag: "model-id", Type: "*string", Required: false},
	{Name: "ModelType", Flag: "model-type", Type: "types.ModelTypeEnum", Required: false},
	{Name: "ModelVersionNumber", Flag: "model-version-number", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_batch_import_jobs = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_batch_prediction_jobs = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_delete_events_by_event_type_status = []leanruntime.Field{
	{Name: "EventTypeName", Flag: "event-type-name", Type: "*string", Required: true},
}

var fields_get_detector_version = []leanruntime.Field{
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "DetectorVersionId", Flag: "detector-version-id", Type: "*string", Required: true},
}

var fields_get_detectors = []leanruntime.Field{
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_entity_types = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_event = []leanruntime.Field{
	{Name: "EventId", Flag: "event-id", Type: "*string", Required: true},
	{Name: "EventTypeName", Flag: "event-type-name", Type: "*string", Required: true},
}

var fields_get_event_prediction = []leanruntime.Field{
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "DetectorVersionId", Flag: "detector-version-id", Type: "*string", Required: false},
	{Name: "Entities", Flag: "entities", Type: "[]types.Entity", Required: true},
	{Name: "EventId", Flag: "event-id", Type: "*string", Required: true},
	{Name: "EventTimestamp", Flag: "event-timestamp", Type: "*string", Required: true},
	{Name: "EventTypeName", Flag: "event-type-name", Type: "*string", Required: true},
	{Name: "EventVariables", Flag: "event-variables", Type: "map[string]string", Required: true},
	{Name: "ExternalModelEndpointDataBlobs", Flag: "external-model-endpoint-data-blobs", Type: "map[string]types.ModelEndpointDataBlob", Required: false},
}

var fields_get_event_prediction_metadata = []leanruntime.Field{
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "DetectorVersionId", Flag: "detector-version-id", Type: "*string", Required: true},
	{Name: "EventId", Flag: "event-id", Type: "*string", Required: true},
	{Name: "EventTypeName", Flag: "event-type-name", Type: "*string", Required: true},
	{Name: "PredictionTimestamp", Flag: "prediction-timestamp", Type: "*string", Required: true},
}

var fields_get_event_types = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_external_models = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "ModelEndpoint", Flag: "model-endpoint", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_kms_encryption_key = []leanruntime.Field{}

var fields_get_labels = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_list_elements = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_lists_metadata = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_model_version = []leanruntime.Field{
	{Name: "ModelId", Flag: "model-id", Type: "*string", Required: true},
	{Name: "ModelType", Flag: "model-type", Type: "types.ModelTypeEnum", Required: true},
	{Name: "ModelVersionNumber", Flag: "model-version-number", Type: "*string", Required: true},
}

var fields_get_models = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "ModelId", Flag: "model-id", Type: "*string", Required: false},
	{Name: "ModelType", Flag: "model-type", Type: "types.ModelTypeEnum", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_outcomes = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_rules = []leanruntime.Field{
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RuleId", Flag: "rule-id", Type: "*string", Required: false},
	{Name: "RuleVersion", Flag: "rule-version", Type: "*string", Required: false},
}

var fields_get_variables = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_event_predictions = []leanruntime.Field{
	{Name: "DetectorId", Flag: "detector-id", Type: "*types.FilterCondition", Required: false},
	{Name: "DetectorVersionId", Flag: "detector-version-id", Type: "*types.FilterCondition", Required: false},
	{Name: "EventId", Flag: "event-id", Type: "*types.FilterCondition", Required: false},
	{Name: "EventType", Flag: "event-type", Type: "*types.FilterCondition", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PredictionTimeRange", Flag: "prediction-time-range", Type: "*types.PredictionTimeRange", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_put_detector = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "EventTypeName", Flag: "event-type-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_put_entity_type = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_put_event_type = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EntityTypes", Flag: "entity-types", Type: "[]string", Required: true},
	{Name: "EventIngestion", Flag: "event-ingestion", Type: "types.EventIngestion", Required: false},
	{Name: "EventOrchestration", Flag: "event-orchestration", Type: "*types.EventOrchestration", Required: false},
	{Name: "EventVariables", Flag: "event-variables", Type: "[]string", Required: true},
	{Name: "Labels", Flag: "labels", Type: "[]string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_put_external_model = []leanruntime.Field{
	{Name: "InputConfiguration", Flag: "input-configuration", Type: "*types.ModelInputConfiguration", Required: true},
	{Name: "InvokeModelEndpointRoleArn", Flag: "invoke-model-endpoint-role-arn", Type: "*string", Required: true},
	{Name: "ModelEndpoint", Flag: "model-endpoint", Type: "*string", Required: true},
	{Name: "ModelEndpointStatus", Flag: "model-endpoint-status", Type: "types.ModelEndpointStatus", Required: true},
	{Name: "ModelSource", Flag: "model-source", Type: "types.ModelSource", Required: true},
	{Name: "OutputConfiguration", Flag: "output-configuration", Type: "*types.ModelOutputConfiguration", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_put_kms_encryption_key = []leanruntime.Field{
	{Name: "KmsEncryptionKeyArn", Flag: "kms-encryption-key-arn", Type: "*string", Required: true},
}

var fields_put_label = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_put_outcome = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_send_event = []leanruntime.Field{
	{Name: "AssignedLabel", Flag: "assigned-label", Type: "*string", Required: false},
	{Name: "Entities", Flag: "entities", Type: "[]types.Entity", Required: true},
	{Name: "EventId", Flag: "event-id", Type: "*string", Required: true},
	{Name: "EventTimestamp", Flag: "event-timestamp", Type: "*string", Required: true},
	{Name: "EventTypeName", Flag: "event-type-name", Type: "*string", Required: true},
	{Name: "EventVariables", Flag: "event-variables", Type: "map[string]string", Required: true},
	{Name: "LabelTimestamp", Flag: "label-timestamp", Type: "*string", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_detector_version = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "DetectorVersionId", Flag: "detector-version-id", Type: "*string", Required: true},
	{Name: "ExternalModelEndpoints", Flag: "external-model-endpoints", Type: "[]string", Required: true},
	{Name: "ModelVersions", Flag: "model-versions", Type: "[]types.ModelVersion", Required: false},
	{Name: "RuleExecutionMode", Flag: "rule-execution-mode", Type: "types.RuleExecutionMode", Required: false},
	{Name: "Rules", Flag: "rules", Type: "[]types.Rule", Required: true},
}

var fields_update_detector_version_metadata = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "DetectorVersionId", Flag: "detector-version-id", Type: "*string", Required: true},
}

var fields_update_detector_version_status = []leanruntime.Field{
	{Name: "DetectorId", Flag: "detector-id", Type: "*string", Required: true},
	{Name: "DetectorVersionId", Flag: "detector-version-id", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.DetectorVersionStatus", Required: true},
}

var fields_update_event_label = []leanruntime.Field{
	{Name: "AssignedLabel", Flag: "assigned-label", Type: "*string", Required: true},
	{Name: "EventId", Flag: "event-id", Type: "*string", Required: true},
	{Name: "EventTypeName", Flag: "event-type-name", Type: "*string", Required: true},
	{Name: "LabelTimestamp", Flag: "label-timestamp", Type: "*string", Required: true},
}

var fields_update_list = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Elements", Flag: "elements", Type: "[]string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "UpdateMode", Flag: "update-mode", Type: "types.ListUpdateMode", Required: false},
	{Name: "VariableType", Flag: "variable-type", Type: "*string", Required: false},
}

var fields_update_model = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ModelId", Flag: "model-id", Type: "*string", Required: true},
	{Name: "ModelType", Flag: "model-type", Type: "types.ModelTypeEnum", Required: true},
}

var fields_update_model_version = []leanruntime.Field{
	{Name: "ExternalEventsDetail", Flag: "external-events-detail", Type: "*types.ExternalEventsDetail", Required: false},
	{Name: "IngestedEventsDetail", Flag: "ingested-events-detail", Type: "*types.IngestedEventsDetail", Required: false},
	{Name: "MajorVersionNumber", Flag: "major-version-number", Type: "*string", Required: true},
	{Name: "ModelId", Flag: "model-id", Type: "*string", Required: true},
	{Name: "ModelType", Flag: "model-type", Type: "types.ModelTypeEnum", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_update_model_version_status = []leanruntime.Field{
	{Name: "ModelId", Flag: "model-id", Type: "*string", Required: true},
	{Name: "ModelType", Flag: "model-type", Type: "types.ModelTypeEnum", Required: true},
	{Name: "ModelVersionNumber", Flag: "model-version-number", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.ModelVersionStatus", Required: true},
}

var fields_update_rule_metadata = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "Rule", Flag: "rule", Type: "*types.Rule", Required: true},
}

var fields_update_rule_version = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Expression", Flag: "expression", Type: "*string", Required: true},
	{Name: "Language", Flag: "language", Type: "types.Language", Required: true},
	{Name: "Outcomes", Flag: "outcomes", Type: "[]string", Required: true},
	{Name: "Rule", Flag: "rule", Type: "*types.Rule", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_update_variable = []leanruntime.Field{
	{Name: "DefaultValue", Flag: "default-value", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "VariableType", Flag: "variable-type", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"batch-create-variable": {
			Name:   "batch-create-variable",
			Fields: fields_batch_create_variable,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchCreateVariableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_create_variable, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchCreateVariable(ctx, input)
			},
		},
		"batch-get-variable": {
			Name:   "batch-get-variable",
			Fields: fields_batch_get_variable,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetVariableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_variable, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetVariable(ctx, input)
			},
		},
		"cancel-batch-import-job": {
			Name:   "cancel-batch-import-job",
			Fields: fields_cancel_batch_import_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelBatchImportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_batch_import_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelBatchImportJob(ctx, input)
			},
		},
		"cancel-batch-prediction-job": {
			Name:   "cancel-batch-prediction-job",
			Fields: fields_cancel_batch_prediction_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelBatchPredictionJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_batch_prediction_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelBatchPredictionJob(ctx, input)
			},
		},
		"create-batch-import-job": {
			Name:   "create-batch-import-job",
			Fields: fields_create_batch_import_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBatchImportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_batch_import_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBatchImportJob(ctx, input)
			},
		},
		"create-batch-prediction-job": {
			Name:   "create-batch-prediction-job",
			Fields: fields_create_batch_prediction_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBatchPredictionJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_batch_prediction_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBatchPredictionJob(ctx, input)
			},
		},
		"create-detector-version": {
			Name:   "create-detector-version",
			Fields: fields_create_detector_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDetectorVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_detector_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDetectorVersion(ctx, input)
			},
		},
		"create-list": {
			Name:   "create-list",
			Fields: fields_create_list,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateListInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_list, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateList(ctx, input)
			},
		},
		"create-model": {
			Name:   "create-model",
			Fields: fields_create_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateModel(ctx, input)
			},
		},
		"create-model-version": {
			Name:   "create-model-version",
			Fields: fields_create_model_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateModelVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_model_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateModelVersion(ctx, input)
			},
		},
		"create-rule": {
			Name:   "create-rule",
			Fields: fields_create_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRule(ctx, input)
			},
		},
		"create-variable": {
			Name:   "create-variable",
			Fields: fields_create_variable,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVariableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_variable, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVariable(ctx, input)
			},
		},
		"delete-batch-import-job": {
			Name:   "delete-batch-import-job",
			Fields: fields_delete_batch_import_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBatchImportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_batch_import_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBatchImportJob(ctx, input)
			},
		},
		"delete-batch-prediction-job": {
			Name:   "delete-batch-prediction-job",
			Fields: fields_delete_batch_prediction_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBatchPredictionJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_batch_prediction_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBatchPredictionJob(ctx, input)
			},
		},
		"delete-detector": {
			Name:   "delete-detector",
			Fields: fields_delete_detector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDetectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_detector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDetector(ctx, input)
			},
		},
		"delete-detector-version": {
			Name:   "delete-detector-version",
			Fields: fields_delete_detector_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDetectorVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_detector_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDetectorVersion(ctx, input)
			},
		},
		"delete-entity-type": {
			Name:   "delete-entity-type",
			Fields: fields_delete_entity_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEntityTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_entity_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEntityType(ctx, input)
			},
		},
		"delete-event": {
			Name:   "delete-event",
			Fields: fields_delete_event,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEventInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_event, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEvent(ctx, input)
			},
		},
		"delete-event-type": {
			Name:   "delete-event-type",
			Fields: fields_delete_event_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEventTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_event_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEventType(ctx, input)
			},
		},
		"delete-events-by-event-type": {
			Name:   "delete-events-by-event-type",
			Fields: fields_delete_events_by_event_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEventsByEventTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_events_by_event_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEventsByEventType(ctx, input)
			},
		},
		"delete-external-model": {
			Name:   "delete-external-model",
			Fields: fields_delete_external_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteExternalModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_external_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteExternalModel(ctx, input)
			},
		},
		"delete-label": {
			Name:   "delete-label",
			Fields: fields_delete_label,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLabelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_label, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLabel(ctx, input)
			},
		},
		"delete-list": {
			Name:   "delete-list",
			Fields: fields_delete_list,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteListInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_list, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteList(ctx, input)
			},
		},
		"delete-model": {
			Name:   "delete-model",
			Fields: fields_delete_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteModel(ctx, input)
			},
		},
		"delete-model-version": {
			Name:   "delete-model-version",
			Fields: fields_delete_model_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteModelVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_model_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteModelVersion(ctx, input)
			},
		},
		"delete-outcome": {
			Name:   "delete-outcome",
			Fields: fields_delete_outcome,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteOutcomeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_outcome, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteOutcome(ctx, input)
			},
		},
		"delete-rule": {
			Name:   "delete-rule",
			Fields: fields_delete_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRule(ctx, input)
			},
		},
		"delete-variable": {
			Name:   "delete-variable",
			Fields: fields_delete_variable,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVariableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_variable, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVariable(ctx, input)
			},
		},
		"describe-detector": {
			Name:   "describe-detector",
			Fields: fields_describe_detector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDetectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_detector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDetector(ctx, input)
			},
		},
		"describe-model-versions": {
			Name:   "describe-model-versions",
			Fields: fields_describe_model_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeModelVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_model_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeModelVersions(ctx, input)
				}
				var results []*svc.DescribeModelVersionsOutput
				p := svc.NewDescribeModelVersionsPaginator(client, input)
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
		"get-batch-import-jobs": {
			Name:   "get-batch-import-jobs",
			Fields: fields_get_batch_import_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBatchImportJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_batch_import_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetBatchImportJobs(ctx, input)
				}
				var results []*svc.GetBatchImportJobsOutput
				p := svc.NewGetBatchImportJobsPaginator(client, input)
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
		"get-batch-prediction-jobs": {
			Name:   "get-batch-prediction-jobs",
			Fields: fields_get_batch_prediction_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBatchPredictionJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_batch_prediction_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetBatchPredictionJobs(ctx, input)
				}
				var results []*svc.GetBatchPredictionJobsOutput
				p := svc.NewGetBatchPredictionJobsPaginator(client, input)
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
		"get-delete-events-by-event-type-status": {
			Name:   "get-delete-events-by-event-type-status",
			Fields: fields_get_delete_events_by_event_type_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDeleteEventsByEventTypeStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_delete_events_by_event_type_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDeleteEventsByEventTypeStatus(ctx, input)
			},
		},
		"get-detector-version": {
			Name:   "get-detector-version",
			Fields: fields_get_detector_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDetectorVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_detector_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDetectorVersion(ctx, input)
			},
		},
		"get-detectors": {
			Name:   "get-detectors",
			Fields: fields_get_detectors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDetectorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_detectors, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetDetectors(ctx, input)
				}
				var results []*svc.GetDetectorsOutput
				p := svc.NewGetDetectorsPaginator(client, input)
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
		"get-entity-types": {
			Name:   "get-entity-types",
			Fields: fields_get_entity_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEntityTypesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_entity_types, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetEntityTypes(ctx, input)
				}
				var results []*svc.GetEntityTypesOutput
				p := svc.NewGetEntityTypesPaginator(client, input)
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
		"get-event": {
			Name:   "get-event",
			Fields: fields_get_event,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEventInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_event, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEvent(ctx, input)
			},
		},
		"get-event-prediction": {
			Name:   "get-event-prediction",
			Fields: fields_get_event_prediction,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEventPredictionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_event_prediction, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEventPrediction(ctx, input)
			},
		},
		"get-event-prediction-metadata": {
			Name:   "get-event-prediction-metadata",
			Fields: fields_get_event_prediction_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEventPredictionMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_event_prediction_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEventPredictionMetadata(ctx, input)
			},
		},
		"get-event-types": {
			Name:   "get-event-types",
			Fields: fields_get_event_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEventTypesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_event_types, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetEventTypes(ctx, input)
				}
				var results []*svc.GetEventTypesOutput
				p := svc.NewGetEventTypesPaginator(client, input)
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
		"get-external-models": {
			Name:   "get-external-models",
			Fields: fields_get_external_models,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetExternalModelsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_external_models, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetExternalModels(ctx, input)
				}
				var results []*svc.GetExternalModelsOutput
				p := svc.NewGetExternalModelsPaginator(client, input)
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
		"get-kms-encryption-key": {
			Name:   "get-kms-encryption-key",
			Fields: fields_get_kms_encryption_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetKMSEncryptionKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_kms_encryption_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetKMSEncryptionKey(ctx, input)
			},
		},
		"get-labels": {
			Name:   "get-labels",
			Fields: fields_get_labels,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLabelsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_labels, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetLabels(ctx, input)
				}
				var results []*svc.GetLabelsOutput
				p := svc.NewGetLabelsPaginator(client, input)
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
		"get-list-elements": {
			Name:   "get-list-elements",
			Fields: fields_get_list_elements,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetListElementsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_list_elements, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetListElements(ctx, input)
				}
				var results []*svc.GetListElementsOutput
				p := svc.NewGetListElementsPaginator(client, input)
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
		"get-lists-metadata": {
			Name:   "get-lists-metadata",
			Fields: fields_get_lists_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetListsMetadataInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_lists_metadata, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetListsMetadata(ctx, input)
				}
				var results []*svc.GetListsMetadataOutput
				p := svc.NewGetListsMetadataPaginator(client, input)
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
		"get-model-version": {
			Name:   "get-model-version",
			Fields: fields_get_model_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetModelVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_model_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetModelVersion(ctx, input)
			},
		},
		"get-models": {
			Name:   "get-models",
			Fields: fields_get_models,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetModelsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_models, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetModels(ctx, input)
				}
				var results []*svc.GetModelsOutput
				p := svc.NewGetModelsPaginator(client, input)
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
		"get-outcomes": {
			Name:   "get-outcomes",
			Fields: fields_get_outcomes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOutcomesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_outcomes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetOutcomes(ctx, input)
				}
				var results []*svc.GetOutcomesOutput
				p := svc.NewGetOutcomesPaginator(client, input)
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
		"get-rules": {
			Name:   "get-rules",
			Fields: fields_get_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRulesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_rules, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetRules(ctx, input)
				}
				var results []*svc.GetRulesOutput
				p := svc.NewGetRulesPaginator(client, input)
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
		"get-variables": {
			Name:   "get-variables",
			Fields: fields_get_variables,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetVariablesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_variables, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetVariables(ctx, input)
				}
				var results []*svc.GetVariablesOutput
				p := svc.NewGetVariablesPaginator(client, input)
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
		"list-event-predictions": {
			Name:   "list-event-predictions",
			Fields: fields_list_event_predictions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEventPredictionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_event_predictions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEventPredictions(ctx, input)
				}
				var results []*svc.ListEventPredictionsOutput
				p := svc.NewListEventPredictionsPaginator(client, input)
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
		"put-detector": {
			Name:   "put-detector",
			Fields: fields_put_detector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutDetectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_detector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutDetector(ctx, input)
			},
		},
		"put-entity-type": {
			Name:   "put-entity-type",
			Fields: fields_put_entity_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutEntityTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_entity_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutEntityType(ctx, input)
			},
		},
		"put-event-type": {
			Name:   "put-event-type",
			Fields: fields_put_event_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutEventTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_event_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutEventType(ctx, input)
			},
		},
		"put-external-model": {
			Name:   "put-external-model",
			Fields: fields_put_external_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutExternalModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_external_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutExternalModel(ctx, input)
			},
		},
		"put-kms-encryption-key": {
			Name:   "put-kms-encryption-key",
			Fields: fields_put_kms_encryption_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutKMSEncryptionKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_kms_encryption_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutKMSEncryptionKey(ctx, input)
			},
		},
		"put-label": {
			Name:   "put-label",
			Fields: fields_put_label,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutLabelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_label, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutLabel(ctx, input)
			},
		},
		"put-outcome": {
			Name:   "put-outcome",
			Fields: fields_put_outcome,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutOutcomeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_outcome, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutOutcome(ctx, input)
			},
		},
		"send-event": {
			Name:   "send-event",
			Fields: fields_send_event,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendEventInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_event, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendEvent(ctx, input)
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
		"update-detector-version": {
			Name:   "update-detector-version",
			Fields: fields_update_detector_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDetectorVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_detector_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDetectorVersion(ctx, input)
			},
		},
		"update-detector-version-metadata": {
			Name:   "update-detector-version-metadata",
			Fields: fields_update_detector_version_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDetectorVersionMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_detector_version_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDetectorVersionMetadata(ctx, input)
			},
		},
		"update-detector-version-status": {
			Name:   "update-detector-version-status",
			Fields: fields_update_detector_version_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDetectorVersionStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_detector_version_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDetectorVersionStatus(ctx, input)
			},
		},
		"update-event-label": {
			Name:   "update-event-label",
			Fields: fields_update_event_label,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEventLabelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_event_label, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEventLabel(ctx, input)
			},
		},
		"update-list": {
			Name:   "update-list",
			Fields: fields_update_list,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateListInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_list, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateList(ctx, input)
			},
		},
		"update-model": {
			Name:   "update-model",
			Fields: fields_update_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateModel(ctx, input)
			},
		},
		"update-model-version": {
			Name:   "update-model-version",
			Fields: fields_update_model_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateModelVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_model_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateModelVersion(ctx, input)
			},
		},
		"update-model-version-status": {
			Name:   "update-model-version-status",
			Fields: fields_update_model_version_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateModelVersionStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_model_version_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateModelVersionStatus(ctx, input)
			},
		},
		"update-rule-metadata": {
			Name:   "update-rule-metadata",
			Fields: fields_update_rule_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRuleMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_rule_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRuleMetadata(ctx, input)
			},
		},
		"update-rule-version": {
			Name:   "update-rule-version",
			Fields: fields_update_rule_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRuleVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_rule_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRuleVersion(ctx, input)
			},
		},
		"update-variable": {
			Name:   "update-variable",
			Fields: fields_update_variable,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateVariableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_variable, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateVariable(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("frauddetector", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
