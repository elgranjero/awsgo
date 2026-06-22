package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/computeoptimizer"
)

var fields_delete_recommendation_preferences = []leanruntime.Field{
	{Name: "RecommendationPreferenceNames", Flag: "recommendation-preference-names", Type: "[]types.RecommendationPreferenceName", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.ResourceType", Required: true},
	{Name: "Scope", Flag: "scope", Type: "*types.Scope", Required: false},
}

var fields_describe_recommendation_export_jobs = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.JobFilter", Required: false},
	{Name: "JobIds", Flag: "job-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_export_auto_scaling_group_recommendations = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: false},
	{Name: "FieldsToExport", Flag: "fields-to-export", Type: "[]types.ExportableAutoScalingGroupField", Required: false},
	{Name: "FileFormat", Flag: "file-format", Type: "types.FileFormat", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "IncludeMemberAccounts", Flag: "include-member-accounts", Type: "bool", Required: false},
	{Name: "RecommendationPreferences", Flag: "recommendation-preferences", Type: "*types.RecommendationPreferences", Required: false},
	{Name: "S3DestinationConfig", Flag: "s3-destination-config", Type: "*types.S3DestinationConfig", Required: true},
}

var fields_export_ebs_volume_recommendations = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: false},
	{Name: "FieldsToExport", Flag: "fields-to-export", Type: "[]types.ExportableVolumeField", Required: false},
	{Name: "FileFormat", Flag: "file-format", Type: "types.FileFormat", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.EBSFilter", Required: false},
	{Name: "IncludeMemberAccounts", Flag: "include-member-accounts", Type: "bool", Required: false},
	{Name: "S3DestinationConfig", Flag: "s3-destination-config", Type: "*types.S3DestinationConfig", Required: true},
}

var fields_export_ec2_instance_recommendations = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: false},
	{Name: "FieldsToExport", Flag: "fields-to-export", Type: "[]types.ExportableInstanceField", Required: false},
	{Name: "FileFormat", Flag: "file-format", Type: "types.FileFormat", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "IncludeMemberAccounts", Flag: "include-member-accounts", Type: "bool", Required: false},
	{Name: "RecommendationPreferences", Flag: "recommendation-preferences", Type: "*types.RecommendationPreferences", Required: false},
	{Name: "S3DestinationConfig", Flag: "s3-destination-config", Type: "*types.S3DestinationConfig", Required: true},
}

var fields_export_ecs_service_recommendations = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: false},
	{Name: "FieldsToExport", Flag: "fields-to-export", Type: "[]types.ExportableECSServiceField", Required: false},
	{Name: "FileFormat", Flag: "file-format", Type: "types.FileFormat", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.ECSServiceRecommendationFilter", Required: false},
	{Name: "IncludeMemberAccounts", Flag: "include-member-accounts", Type: "bool", Required: false},
	{Name: "S3DestinationConfig", Flag: "s3-destination-config", Type: "*types.S3DestinationConfig", Required: true},
}

var fields_export_idle_recommendations = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: false},
	{Name: "FieldsToExport", Flag: "fields-to-export", Type: "[]types.ExportableIdleField", Required: false},
	{Name: "FileFormat", Flag: "file-format", Type: "types.FileFormat", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.IdleRecommendationFilter", Required: false},
	{Name: "IncludeMemberAccounts", Flag: "include-member-accounts", Type: "bool", Required: false},
	{Name: "S3DestinationConfig", Flag: "s3-destination-config", Type: "*types.S3DestinationConfig", Required: true},
}

var fields_export_lambda_function_recommendations = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: false},
	{Name: "FieldsToExport", Flag: "fields-to-export", Type: "[]types.ExportableLambdaFunctionField", Required: false},
	{Name: "FileFormat", Flag: "file-format", Type: "types.FileFormat", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.LambdaFunctionRecommendationFilter", Required: false},
	{Name: "IncludeMemberAccounts", Flag: "include-member-accounts", Type: "bool", Required: false},
	{Name: "S3DestinationConfig", Flag: "s3-destination-config", Type: "*types.S3DestinationConfig", Required: true},
}

var fields_export_license_recommendations = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: false},
	{Name: "FieldsToExport", Flag: "fields-to-export", Type: "[]types.ExportableLicenseField", Required: false},
	{Name: "FileFormat", Flag: "file-format", Type: "types.FileFormat", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.LicenseRecommendationFilter", Required: false},
	{Name: "IncludeMemberAccounts", Flag: "include-member-accounts", Type: "bool", Required: false},
	{Name: "S3DestinationConfig", Flag: "s3-destination-config", Type: "*types.S3DestinationConfig", Required: true},
}

var fields_export_rds_database_recommendations = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: false},
	{Name: "FieldsToExport", Flag: "fields-to-export", Type: "[]types.ExportableRDSDBField", Required: false},
	{Name: "FileFormat", Flag: "file-format", Type: "types.FileFormat", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.RDSDBRecommendationFilter", Required: false},
	{Name: "IncludeMemberAccounts", Flag: "include-member-accounts", Type: "bool", Required: false},
	{Name: "RecommendationPreferences", Flag: "recommendation-preferences", Type: "*types.RecommendationPreferences", Required: false},
	{Name: "S3DestinationConfig", Flag: "s3-destination-config", Type: "*types.S3DestinationConfig", Required: true},
}

var fields_get_auto_scaling_group_recommendations = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: false},
	{Name: "AutoScalingGroupArns", Flag: "auto-scaling-group-arns", Type: "[]string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RecommendationPreferences", Flag: "recommendation-preferences", Type: "*types.RecommendationPreferences", Required: false},
}

var fields_get_ebs_volume_recommendations = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.EBSFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "VolumeArns", Flag: "volume-arns", Type: "[]string", Required: false},
}

var fields_get_ec2_instance_recommendations = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "InstanceArns", Flag: "instance-arns", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RecommendationPreferences", Flag: "recommendation-preferences", Type: "*types.RecommendationPreferences", Required: false},
}

var fields_get_ec2_recommendation_projected_metrics = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: true},
	{Name: "Period", Flag: "period", Type: "int32", Required: true},
	{Name: "RecommendationPreferences", Flag: "recommendation-preferences", Type: "*types.RecommendationPreferences", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
	{Name: "Stat", Flag: "stat", Type: "types.MetricStatistic", Required: true},
}

var fields_get_ecs_service_recommendation_projected_metrics = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "Period", Flag: "period", Type: "int32", Required: true},
	{Name: "ServiceArn", Flag: "service-arn", Type: "*string", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
	{Name: "Stat", Flag: "stat", Type: "types.MetricStatistic", Required: true},
}

var fields_get_ecs_service_recommendations = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.ECSServiceRecommendationFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ServiceArns", Flag: "service-arns", Type: "[]string", Required: false},
}

var fields_get_effective_recommendation_preferences = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_get_enrollment_status = []leanruntime.Field{}

var fields_get_enrollment_statuses_for_organization = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.EnrollmentFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_idle_recommendations = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.IdleRecommendationFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OrderBy", Flag: "order-by", Type: "*types.OrderBy", Required: false},
	{Name: "ResourceArns", Flag: "resource-arns", Type: "[]string", Required: false},
}

var fields_get_lambda_function_recommendations = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.LambdaFunctionRecommendationFilter", Required: false},
	{Name: "FunctionArns", Flag: "function-arns", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_license_recommendations = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.LicenseRecommendationFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceArns", Flag: "resource-arns", Type: "[]string", Required: false},
}

var fields_get_rds_database_recommendation_projected_metrics = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "Period", Flag: "period", Type: "int32", Required: true},
	{Name: "RecommendationPreferences", Flag: "recommendation-preferences", Type: "*types.RecommendationPreferences", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
	{Name: "Stat", Flag: "stat", Type: "types.MetricStatistic", Required: true},
}

var fields_get_rds_database_recommendations = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.RDSDBRecommendationFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RecommendationPreferences", Flag: "recommendation-preferences", Type: "*types.RecommendationPreferences", Required: false},
	{Name: "ResourceArns", Flag: "resource-arns", Type: "[]string", Required: false},
}

var fields_get_recommendation_preferences = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.ResourceType", Required: true},
	{Name: "Scope", Flag: "scope", Type: "*types.Scope", Required: false},
}

var fields_get_recommendation_summaries = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_put_recommendation_preferences = []leanruntime.Field{
	{Name: "EnhancedInfrastructureMetrics", Flag: "enhanced-infrastructure-metrics", Type: "types.EnhancedInfrastructureMetrics", Required: false},
	{Name: "ExternalMetricsPreference", Flag: "external-metrics-preference", Type: "*types.ExternalMetricsPreference", Required: false},
	{Name: "InferredWorkloadTypes", Flag: "inferred-workload-types", Type: "types.InferredWorkloadTypesPreference", Required: false},
	{Name: "LookBackPeriod", Flag: "look-back-period", Type: "types.LookBackPeriodPreference", Required: false},
	{Name: "PreferredResources", Flag: "preferred-resources", Type: "[]types.PreferredResource", Required: false},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.ResourceType", Required: true},
	{Name: "SavingsEstimationMode", Flag: "savings-estimation-mode", Type: "types.SavingsEstimationMode", Required: false},
	{Name: "Scope", Flag: "scope", Type: "*types.Scope", Required: false},
	{Name: "UtilizationPreferences", Flag: "utilization-preferences", Type: "[]types.UtilizationPreference", Required: false},
}

var fields_update_enrollment_status = []leanruntime.Field{
	{Name: "IncludeMemberAccounts", Flag: "include-member-accounts", Type: "bool", Required: false},
	{Name: "Status", Flag: "status", Type: "types.Status", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"delete-recommendation-preferences": {
			Name:   "delete-recommendation-preferences",
			Fields: fields_delete_recommendation_preferences,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRecommendationPreferencesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_recommendation_preferences, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRecommendationPreferences(ctx, input)
			},
		},
		"describe-recommendation-export-jobs": {
			Name:   "describe-recommendation-export-jobs",
			Fields: fields_describe_recommendation_export_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRecommendationExportJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_recommendation_export_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeRecommendationExportJobs(ctx, input)
				}
				var results []*svc.DescribeRecommendationExportJobsOutput
				p := svc.NewDescribeRecommendationExportJobsPaginator(client, input)
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
		"export-auto-scaling-group-recommendations": {
			Name:   "export-auto-scaling-group-recommendations",
			Fields: fields_export_auto_scaling_group_recommendations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExportAutoScalingGroupRecommendationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_export_auto_scaling_group_recommendations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExportAutoScalingGroupRecommendations(ctx, input)
			},
		},
		"export-ebs-volume-recommendations": {
			Name:   "export-ebs-volume-recommendations",
			Fields: fields_export_ebs_volume_recommendations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExportEBSVolumeRecommendationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_export_ebs_volume_recommendations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExportEBSVolumeRecommendations(ctx, input)
			},
		},
		"export-ec2-instance-recommendations": {
			Name:   "export-ec2-instance-recommendations",
			Fields: fields_export_ec2_instance_recommendations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExportEC2InstanceRecommendationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_export_ec2_instance_recommendations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExportEC2InstanceRecommendations(ctx, input)
			},
		},
		"export-ecs-service-recommendations": {
			Name:   "export-ecs-service-recommendations",
			Fields: fields_export_ecs_service_recommendations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExportECSServiceRecommendationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_export_ecs_service_recommendations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExportECSServiceRecommendations(ctx, input)
			},
		},
		"export-idle-recommendations": {
			Name:   "export-idle-recommendations",
			Fields: fields_export_idle_recommendations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExportIdleRecommendationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_export_idle_recommendations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExportIdleRecommendations(ctx, input)
			},
		},
		"export-lambda-function-recommendations": {
			Name:   "export-lambda-function-recommendations",
			Fields: fields_export_lambda_function_recommendations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExportLambdaFunctionRecommendationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_export_lambda_function_recommendations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExportLambdaFunctionRecommendations(ctx, input)
			},
		},
		"export-license-recommendations": {
			Name:   "export-license-recommendations",
			Fields: fields_export_license_recommendations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExportLicenseRecommendationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_export_license_recommendations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExportLicenseRecommendations(ctx, input)
			},
		},
		"export-rds-database-recommendations": {
			Name:   "export-rds-database-recommendations",
			Fields: fields_export_rds_database_recommendations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExportRDSDatabaseRecommendationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_export_rds_database_recommendations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExportRDSDatabaseRecommendations(ctx, input)
			},
		},
		"get-auto-scaling-group-recommendations": {
			Name:   "get-auto-scaling-group-recommendations",
			Fields: fields_get_auto_scaling_group_recommendations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAutoScalingGroupRecommendationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_auto_scaling_group_recommendations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAutoScalingGroupRecommendations(ctx, input)
			},
		},
		"get-ebs-volume-recommendations": {
			Name:   "get-ebs-volume-recommendations",
			Fields: fields_get_ebs_volume_recommendations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEBSVolumeRecommendationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_ebs_volume_recommendations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEBSVolumeRecommendations(ctx, input)
			},
		},
		"get-ec2-instance-recommendations": {
			Name:   "get-ec2-instance-recommendations",
			Fields: fields_get_ec2_instance_recommendations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEC2InstanceRecommendationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_ec2_instance_recommendations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEC2InstanceRecommendations(ctx, input)
			},
		},
		"get-ec2-recommendation-projected-metrics": {
			Name:   "get-ec2-recommendation-projected-metrics",
			Fields: fields_get_ec2_recommendation_projected_metrics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEC2RecommendationProjectedMetricsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_ec2_recommendation_projected_metrics, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEC2RecommendationProjectedMetrics(ctx, input)
			},
		},
		"get-ecs-service-recommendation-projected-metrics": {
			Name:   "get-ecs-service-recommendation-projected-metrics",
			Fields: fields_get_ecs_service_recommendation_projected_metrics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetECSServiceRecommendationProjectedMetricsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_ecs_service_recommendation_projected_metrics, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetECSServiceRecommendationProjectedMetrics(ctx, input)
			},
		},
		"get-ecs-service-recommendations": {
			Name:   "get-ecs-service-recommendations",
			Fields: fields_get_ecs_service_recommendations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetECSServiceRecommendationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_ecs_service_recommendations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetECSServiceRecommendations(ctx, input)
			},
		},
		"get-effective-recommendation-preferences": {
			Name:   "get-effective-recommendation-preferences",
			Fields: fields_get_effective_recommendation_preferences,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEffectiveRecommendationPreferencesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_effective_recommendation_preferences, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEffectiveRecommendationPreferences(ctx, input)
			},
		},
		"get-enrollment-status": {
			Name:   "get-enrollment-status",
			Fields: fields_get_enrollment_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEnrollmentStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_enrollment_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEnrollmentStatus(ctx, input)
			},
		},
		"get-enrollment-statuses-for-organization": {
			Name:   "get-enrollment-statuses-for-organization",
			Fields: fields_get_enrollment_statuses_for_organization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEnrollmentStatusesForOrganizationInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_enrollment_statuses_for_organization, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetEnrollmentStatusesForOrganization(ctx, input)
				}
				var results []*svc.GetEnrollmentStatusesForOrganizationOutput
				p := svc.NewGetEnrollmentStatusesForOrganizationPaginator(client, input)
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
		"get-idle-recommendations": {
			Name:   "get-idle-recommendations",
			Fields: fields_get_idle_recommendations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIdleRecommendationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_idle_recommendations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIdleRecommendations(ctx, input)
			},
		},
		"get-lambda-function-recommendations": {
			Name:   "get-lambda-function-recommendations",
			Fields: fields_get_lambda_function_recommendations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLambdaFunctionRecommendationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_lambda_function_recommendations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetLambdaFunctionRecommendations(ctx, input)
				}
				var results []*svc.GetLambdaFunctionRecommendationsOutput
				p := svc.NewGetLambdaFunctionRecommendationsPaginator(client, input)
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
		"get-license-recommendations": {
			Name:   "get-license-recommendations",
			Fields: fields_get_license_recommendations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLicenseRecommendationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_license_recommendations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLicenseRecommendations(ctx, input)
			},
		},
		"get-rds-database-recommendation-projected-metrics": {
			Name:   "get-rds-database-recommendation-projected-metrics",
			Fields: fields_get_rds_database_recommendation_projected_metrics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRDSDatabaseRecommendationProjectedMetricsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_rds_database_recommendation_projected_metrics, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRDSDatabaseRecommendationProjectedMetrics(ctx, input)
			},
		},
		"get-rds-database-recommendations": {
			Name:   "get-rds-database-recommendations",
			Fields: fields_get_rds_database_recommendations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRDSDatabaseRecommendationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_rds_database_recommendations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRDSDatabaseRecommendations(ctx, input)
			},
		},
		"get-recommendation-preferences": {
			Name:   "get-recommendation-preferences",
			Fields: fields_get_recommendation_preferences,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRecommendationPreferencesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_recommendation_preferences, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetRecommendationPreferences(ctx, input)
				}
				var results []*svc.GetRecommendationPreferencesOutput
				p := svc.NewGetRecommendationPreferencesPaginator(client, input)
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
		"get-recommendation-summaries": {
			Name:   "get-recommendation-summaries",
			Fields: fields_get_recommendation_summaries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRecommendationSummariesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_recommendation_summaries, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetRecommendationSummaries(ctx, input)
				}
				var results []*svc.GetRecommendationSummariesOutput
				p := svc.NewGetRecommendationSummariesPaginator(client, input)
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
		"put-recommendation-preferences": {
			Name:   "put-recommendation-preferences",
			Fields: fields_put_recommendation_preferences,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutRecommendationPreferencesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_recommendation_preferences, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutRecommendationPreferences(ctx, input)
			},
		},
		"update-enrollment-status": {
			Name:   "update-enrollment-status",
			Fields: fields_update_enrollment_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEnrollmentStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_enrollment_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEnrollmentStatus(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("computeoptimizer", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
