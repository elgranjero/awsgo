package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/iotevents"
)

var fields_create_alarm_model = []leanruntime.Field{
	{Name: "AlarmCapabilities", Flag: "alarm-capabilities", Type: "*types.AlarmCapabilities", Required: false},
	{Name: "AlarmEventActions", Flag: "alarm-event-actions", Type: "*types.AlarmEventActions", Required: false},
	{Name: "AlarmModelDescription", Flag: "alarm-model-description", Type: "*string", Required: false},
	{Name: "AlarmModelName", Flag: "alarm-model-name", Type: "*string", Required: true},
	{Name: "AlarmNotification", Flag: "alarm-notification", Type: "*types.AlarmNotification", Required: false},
	{Name: "AlarmRule", Flag: "alarm-rule", Type: "*types.AlarmRule", Required: true},
	{Name: "Key", Flag: "key", Type: "*string", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Severity", Flag: "severity", Type: "*int32", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_detector_model = []leanruntime.Field{
	{Name: "DetectorModelDefinition", Flag: "detector-model-definition", Type: "*types.DetectorModelDefinition", Required: true},
	{Name: "DetectorModelDescription", Flag: "detector-model-description", Type: "*string", Required: false},
	{Name: "DetectorModelName", Flag: "detector-model-name", Type: "*string", Required: true},
	{Name: "EvaluationMethod", Flag: "evaluation-method", Type: "types.EvaluationMethod", Required: false},
	{Name: "Key", Flag: "key", Type: "*string", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_input = []leanruntime.Field{
	{Name: "InputDefinition", Flag: "input-definition", Type: "*types.InputDefinition", Required: true},
	{Name: "InputDescription", Flag: "input-description", Type: "*string", Required: false},
	{Name: "InputName", Flag: "input-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_delete_alarm_model = []leanruntime.Field{
	{Name: "AlarmModelName", Flag: "alarm-model-name", Type: "*string", Required: true},
}

var fields_delete_detector_model = []leanruntime.Field{
	{Name: "DetectorModelName", Flag: "detector-model-name", Type: "*string", Required: true},
}

var fields_delete_input = []leanruntime.Field{
	{Name: "InputName", Flag: "input-name", Type: "*string", Required: true},
}

var fields_describe_alarm_model = []leanruntime.Field{
	{Name: "AlarmModelName", Flag: "alarm-model-name", Type: "*string", Required: true},
	{Name: "AlarmModelVersion", Flag: "alarm-model-version", Type: "*string", Required: false},
}

var fields_describe_detector_model = []leanruntime.Field{
	{Name: "DetectorModelName", Flag: "detector-model-name", Type: "*string", Required: true},
	{Name: "DetectorModelVersion", Flag: "detector-model-version", Type: "*string", Required: false},
}

var fields_describe_detector_model_analysis = []leanruntime.Field{
	{Name: "AnalysisId", Flag: "analysis-id", Type: "*string", Required: true},
}

var fields_describe_input = []leanruntime.Field{
	{Name: "InputName", Flag: "input-name", Type: "*string", Required: true},
}

var fields_describe_logging_options = []leanruntime.Field{}

var fields_get_detector_model_analysis_results = []leanruntime.Field{
	{Name: "AnalysisId", Flag: "analysis-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_alarm_model_versions = []leanruntime.Field{
	{Name: "AlarmModelName", Flag: "alarm-model-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_alarm_models = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_detector_model_versions = []leanruntime.Field{
	{Name: "DetectorModelName", Flag: "detector-model-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_detector_models = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_input_routings = []leanruntime.Field{
	{Name: "InputIdentifier", Flag: "input-identifier", Type: "*types.InputIdentifier", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_inputs = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_put_logging_options = []leanruntime.Field{
	{Name: "LoggingOptions", Flag: "logging-options", Type: "*types.LoggingOptions", Required: true},
}

var fields_start_detector_model_analysis = []leanruntime.Field{
	{Name: "DetectorModelDefinition", Flag: "detector-model-definition", Type: "*types.DetectorModelDefinition", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_alarm_model = []leanruntime.Field{
	{Name: "AlarmCapabilities", Flag: "alarm-capabilities", Type: "*types.AlarmCapabilities", Required: false},
	{Name: "AlarmEventActions", Flag: "alarm-event-actions", Type: "*types.AlarmEventActions", Required: false},
	{Name: "AlarmModelDescription", Flag: "alarm-model-description", Type: "*string", Required: false},
	{Name: "AlarmModelName", Flag: "alarm-model-name", Type: "*string", Required: true},
	{Name: "AlarmNotification", Flag: "alarm-notification", Type: "*types.AlarmNotification", Required: false},
	{Name: "AlarmRule", Flag: "alarm-rule", Type: "*types.AlarmRule", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Severity", Flag: "severity", Type: "*int32", Required: false},
}

var fields_update_detector_model = []leanruntime.Field{
	{Name: "DetectorModelDefinition", Flag: "detector-model-definition", Type: "*types.DetectorModelDefinition", Required: true},
	{Name: "DetectorModelDescription", Flag: "detector-model-description", Type: "*string", Required: false},
	{Name: "DetectorModelName", Flag: "detector-model-name", Type: "*string", Required: true},
	{Name: "EvaluationMethod", Flag: "evaluation-method", Type: "types.EvaluationMethod", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
}

var fields_update_input = []leanruntime.Field{
	{Name: "InputDefinition", Flag: "input-definition", Type: "*types.InputDefinition", Required: true},
	{Name: "InputDescription", Flag: "input-description", Type: "*string", Required: false},
	{Name: "InputName", Flag: "input-name", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-alarm-model": {
			Name:   "create-alarm-model",
			Fields: fields_create_alarm_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAlarmModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_alarm_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAlarmModel(ctx, input)
			},
		},
		"create-detector-model": {
			Name:   "create-detector-model",
			Fields: fields_create_detector_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDetectorModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_detector_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDetectorModel(ctx, input)
			},
		},
		"create-input": {
			Name:   "create-input",
			Fields: fields_create_input,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateInputInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_input, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateInput(ctx, input)
			},
		},
		"delete-alarm-model": {
			Name:   "delete-alarm-model",
			Fields: fields_delete_alarm_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAlarmModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_alarm_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAlarmModel(ctx, input)
			},
		},
		"delete-detector-model": {
			Name:   "delete-detector-model",
			Fields: fields_delete_detector_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDetectorModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_detector_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDetectorModel(ctx, input)
			},
		},
		"delete-input": {
			Name:   "delete-input",
			Fields: fields_delete_input,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteInputInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_input, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteInput(ctx, input)
			},
		},
		"describe-alarm-model": {
			Name:   "describe-alarm-model",
			Fields: fields_describe_alarm_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAlarmModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_alarm_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAlarmModel(ctx, input)
			},
		},
		"describe-detector-model": {
			Name:   "describe-detector-model",
			Fields: fields_describe_detector_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDetectorModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_detector_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDetectorModel(ctx, input)
			},
		},
		"describe-detector-model-analysis": {
			Name:   "describe-detector-model-analysis",
			Fields: fields_describe_detector_model_analysis,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDetectorModelAnalysisInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_detector_model_analysis, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDetectorModelAnalysis(ctx, input)
			},
		},
		"describe-input": {
			Name:   "describe-input",
			Fields: fields_describe_input,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInputInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_input, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeInput(ctx, input)
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
		"get-detector-model-analysis-results": {
			Name:   "get-detector-model-analysis-results",
			Fields: fields_get_detector_model_analysis_results,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDetectorModelAnalysisResultsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_detector_model_analysis_results, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDetectorModelAnalysisResults(ctx, input)
			},
		},
		"list-alarm-model-versions": {
			Name:   "list-alarm-model-versions",
			Fields: fields_list_alarm_model_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAlarmModelVersionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_alarm_model_versions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListAlarmModelVersions(ctx, input)
			},
		},
		"list-alarm-models": {
			Name:   "list-alarm-models",
			Fields: fields_list_alarm_models,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAlarmModelsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_alarm_models, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListAlarmModels(ctx, input)
			},
		},
		"list-detector-model-versions": {
			Name:   "list-detector-model-versions",
			Fields: fields_list_detector_model_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDetectorModelVersionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_detector_model_versions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListDetectorModelVersions(ctx, input)
			},
		},
		"list-detector-models": {
			Name:   "list-detector-models",
			Fields: fields_list_detector_models,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDetectorModelsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_detector_models, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListDetectorModels(ctx, input)
			},
		},
		"list-input-routings": {
			Name:   "list-input-routings",
			Fields: fields_list_input_routings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInputRoutingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_input_routings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListInputRoutings(ctx, input)
			},
		},
		"list-inputs": {
			Name:   "list-inputs",
			Fields: fields_list_inputs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInputsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_inputs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListInputs(ctx, input)
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
		"start-detector-model-analysis": {
			Name:   "start-detector-model-analysis",
			Fields: fields_start_detector_model_analysis,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartDetectorModelAnalysisInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_detector_model_analysis, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartDetectorModelAnalysis(ctx, input)
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
		"update-alarm-model": {
			Name:   "update-alarm-model",
			Fields: fields_update_alarm_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAlarmModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_alarm_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAlarmModel(ctx, input)
			},
		},
		"update-detector-model": {
			Name:   "update-detector-model",
			Fields: fields_update_detector_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDetectorModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_detector_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDetectorModel(ctx, input)
			},
		},
		"update-input": {
			Name:   "update-input",
			Fields: fields_update_input,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateInputInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_input, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateInput(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("iotevents", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
