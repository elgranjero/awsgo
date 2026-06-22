package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/migrationhubstrategy"
)

var fields_get_application_component_details = []leanruntime.Field{
	{Name: "ApplicationComponentId", Flag: "application-component-id", Type: "*string", Required: true},
}

var fields_get_application_component_strategies = []leanruntime.Field{
	{Name: "ApplicationComponentId", Flag: "application-component-id", Type: "*string", Required: true},
}

var fields_get_assessment = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_import_file_task = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_latest_assessment_id = []leanruntime.Field{}

var fields_get_portfolio_preferences = []leanruntime.Field{}

var fields_get_portfolio_summary = []leanruntime.Field{}

var fields_get_recommendation_report_details = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_server_details = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ServerId", Flag: "server-id", Type: "*string", Required: true},
}

var fields_get_server_strategies = []leanruntime.Field{
	{Name: "ServerId", Flag: "server-id", Type: "*string", Required: true},
}

var fields_list_analyzable_servers = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Sort", Flag: "sort", Type: "types.SortOrder", Required: false},
}

var fields_list_application_components = []leanruntime.Field{
	{Name: "ApplicationComponentCriteria", Flag: "application-component-criteria", Type: "types.ApplicationComponentCriteria", Required: false},
	{Name: "FilterValue", Flag: "filter-value", Type: "*string", Required: false},
	{Name: "GroupIdFilter", Flag: "group-id-filter", Type: "[]types.Group", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Sort", Flag: "sort", Type: "types.SortOrder", Required: false},
}

var fields_list_collectors = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_import_file_task = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_servers = []leanruntime.Field{
	{Name: "FilterValue", Flag: "filter-value", Type: "*string", Required: false},
	{Name: "GroupIdFilter", Flag: "group-id-filter", Type: "[]types.Group", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ServerCriteria", Flag: "server-criteria", Type: "types.ServerCriteria", Required: false},
	{Name: "Sort", Flag: "sort", Type: "types.SortOrder", Required: false},
}

var fields_put_portfolio_preferences = []leanruntime.Field{
	{Name: "ApplicationMode", Flag: "application-mode", Type: "types.ApplicationMode", Required: false},
	{Name: "ApplicationPreferences", Flag: "application-preferences", Type: "*types.ApplicationPreferences", Required: false},
	{Name: "DatabasePreferences", Flag: "database-preferences", Type: "*types.DatabasePreferences", Required: false},
	{Name: "PrioritizeBusinessGoals", Flag: "prioritize-business-goals", Type: "*types.PrioritizeBusinessGoals", Required: false},
}

var fields_start_assessment = []leanruntime.Field{
	{Name: "AssessmentDataSourceType", Flag: "assessment-data-source-type", Type: "types.AssessmentDataSourceType", Required: false},
	{Name: "AssessmentTargets", Flag: "assessment-targets", Type: "[]types.AssessmentTarget", Required: false},
	{Name: "S3bucketForAnalysisData", Flag: "s3bucket-for-analysis-data", Type: "*string", Required: false},
	{Name: "S3bucketForReportData", Flag: "s3bucket-for-report-data", Type: "*string", Required: false},
}

var fields_start_import_file_task = []leanruntime.Field{
	{Name: "DataSourceType", Flag: "data-source-type", Type: "types.DataSourceType", Required: false},
	{Name: "GroupId", Flag: "group-id", Type: "[]types.Group", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "S3Bucket", Flag: "s3-bucket", Type: "*string", Required: true},
	{Name: "S3bucketForReportData", Flag: "s3bucket-for-report-data", Type: "*string", Required: false},
	{Name: "S3key", Flag: "s3key", Type: "*string", Required: true},
}

var fields_start_recommendation_report_generation = []leanruntime.Field{
	{Name: "GroupIdFilter", Flag: "group-id-filter", Type: "[]types.Group", Required: false},
	{Name: "OutputFormat", Flag: "output-format", Type: "types.OutputFormat", Required: false},
}

var fields_stop_assessment = []leanruntime.Field{
	{Name: "AssessmentId", Flag: "assessment-id", Type: "*string", Required: true},
}

var fields_update_application_component_config = []leanruntime.Field{
	{Name: "AppType", Flag: "app-type", Type: "types.AppType", Required: false},
	{Name: "ApplicationComponentId", Flag: "application-component-id", Type: "*string", Required: true},
	{Name: "ConfigureOnly", Flag: "configure-only", Type: "*bool", Required: false},
	{Name: "InclusionStatus", Flag: "inclusion-status", Type: "types.InclusionStatus", Required: false},
	{Name: "SecretsManagerKey", Flag: "secrets-manager-key", Type: "*string", Required: false},
	{Name: "SourceCodeList", Flag: "source-code-list", Type: "[]types.SourceCode", Required: false},
	{Name: "StrategyOption", Flag: "strategy-option", Type: "*types.StrategyOption", Required: false},
}

var fields_update_server_config = []leanruntime.Field{
	{Name: "ServerId", Flag: "server-id", Type: "*string", Required: true},
	{Name: "StrategyOption", Flag: "strategy-option", Type: "*types.StrategyOption", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"get-application-component-details": {
			Name:   "get-application-component-details",
			Fields: fields_get_application_component_details,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetApplicationComponentDetailsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_application_component_details, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetApplicationComponentDetails(ctx, input)
			},
		},
		"get-application-component-strategies": {
			Name:   "get-application-component-strategies",
			Fields: fields_get_application_component_strategies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetApplicationComponentStrategiesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_application_component_strategies, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetApplicationComponentStrategies(ctx, input)
			},
		},
		"get-assessment": {
			Name:   "get-assessment",
			Fields: fields_get_assessment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAssessmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_assessment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAssessment(ctx, input)
			},
		},
		"get-import-file-task": {
			Name:   "get-import-file-task",
			Fields: fields_get_import_file_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetImportFileTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_import_file_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetImportFileTask(ctx, input)
			},
		},
		"get-latest-assessment-id": {
			Name:   "get-latest-assessment-id",
			Fields: fields_get_latest_assessment_id,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLatestAssessmentIdInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_latest_assessment_id, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLatestAssessmentId(ctx, input)
			},
		},
		"get-portfolio-preferences": {
			Name:   "get-portfolio-preferences",
			Fields: fields_get_portfolio_preferences,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPortfolioPreferencesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_portfolio_preferences, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPortfolioPreferences(ctx, input)
			},
		},
		"get-portfolio-summary": {
			Name:   "get-portfolio-summary",
			Fields: fields_get_portfolio_summary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPortfolioSummaryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_portfolio_summary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPortfolioSummary(ctx, input)
			},
		},
		"get-recommendation-report-details": {
			Name:   "get-recommendation-report-details",
			Fields: fields_get_recommendation_report_details,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRecommendationReportDetailsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_recommendation_report_details, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRecommendationReportDetails(ctx, input)
			},
		},
		"get-server-details": {
			Name:   "get-server-details",
			Fields: fields_get_server_details,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetServerDetailsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_server_details, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetServerDetails(ctx, input)
				}
				var results []*svc.GetServerDetailsOutput
				p := svc.NewGetServerDetailsPaginator(client, input)
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
		"get-server-strategies": {
			Name:   "get-server-strategies",
			Fields: fields_get_server_strategies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetServerStrategiesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_server_strategies, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetServerStrategies(ctx, input)
			},
		},
		"list-analyzable-servers": {
			Name:   "list-analyzable-servers",
			Fields: fields_list_analyzable_servers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAnalyzableServersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_analyzable_servers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAnalyzableServers(ctx, input)
				}
				var results []*svc.ListAnalyzableServersOutput
				p := svc.NewListAnalyzableServersPaginator(client, input)
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
		"list-application-components": {
			Name:   "list-application-components",
			Fields: fields_list_application_components,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListApplicationComponentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_application_components, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListApplicationComponents(ctx, input)
				}
				var results []*svc.ListApplicationComponentsOutput
				p := svc.NewListApplicationComponentsPaginator(client, input)
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
		"list-collectors": {
			Name:   "list-collectors",
			Fields: fields_list_collectors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCollectorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_collectors, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCollectors(ctx, input)
				}
				var results []*svc.ListCollectorsOutput
				p := svc.NewListCollectorsPaginator(client, input)
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
		"list-import-file-task": {
			Name:   "list-import-file-task",
			Fields: fields_list_import_file_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListImportFileTaskInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_import_file_task, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListImportFileTask(ctx, input)
				}
				var results []*svc.ListImportFileTaskOutput
				p := svc.NewListImportFileTaskPaginator(client, input)
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
		"list-servers": {
			Name:   "list-servers",
			Fields: fields_list_servers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListServersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_servers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListServers(ctx, input)
				}
				var results []*svc.ListServersOutput
				p := svc.NewListServersPaginator(client, input)
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
		"put-portfolio-preferences": {
			Name:   "put-portfolio-preferences",
			Fields: fields_put_portfolio_preferences,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutPortfolioPreferencesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_portfolio_preferences, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutPortfolioPreferences(ctx, input)
			},
		},
		"start-assessment": {
			Name:   "start-assessment",
			Fields: fields_start_assessment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartAssessmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_assessment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartAssessment(ctx, input)
			},
		},
		"start-import-file-task": {
			Name:   "start-import-file-task",
			Fields: fields_start_import_file_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartImportFileTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_import_file_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartImportFileTask(ctx, input)
			},
		},
		"start-recommendation-report-generation": {
			Name:   "start-recommendation-report-generation",
			Fields: fields_start_recommendation_report_generation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartRecommendationReportGenerationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_recommendation_report_generation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartRecommendationReportGeneration(ctx, input)
			},
		},
		"stop-assessment": {
			Name:   "stop-assessment",
			Fields: fields_stop_assessment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopAssessmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_assessment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopAssessment(ctx, input)
			},
		},
		"update-application-component-config": {
			Name:   "update-application-component-config",
			Fields: fields_update_application_component_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateApplicationComponentConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_application_component_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateApplicationComponentConfig(ctx, input)
			},
		},
		"update-server-config": {
			Name:   "update-server-config",
			Fields: fields_update_server_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateServerConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_server_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateServerConfig(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("migrationhubstrategy", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
