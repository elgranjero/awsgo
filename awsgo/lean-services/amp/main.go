package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/amp"
)

var fields_create_alert_manager_definition = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Data", Flag: "data", Type: "[]byte", Required: true},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_create_anomaly_detector = []leanruntime.Field{
	{Name: "Alias", Flag: "alias", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Configuration", Flag: "configuration", Type: "types.AnomalyDetectorConfiguration", Required: true},
	{Name: "EvaluationIntervalInSeconds", Flag: "evaluation-interval-in-seconds", Type: "*int32", Required: false},
	{Name: "Labels", Flag: "labels", Type: "map[string]string", Required: false},
	{Name: "MissingDataAction", Flag: "missing-data-action", Type: "types.AnomalyDetectorMissingDataAction", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_create_logging_configuration = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "LogGroupArn", Flag: "log-group-arn", Type: "*string", Required: true},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_create_query_logging_configuration = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Destinations", Flag: "destinations", Type: "[]types.LoggingDestination", Required: true},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_create_rule_groups_namespace = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Data", Flag: "data", Type: "[]byte", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_create_scraper = []leanruntime.Field{
	{Name: "Alias", Flag: "alias", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Destination", Flag: "destination", Type: "types.Destination", Required: true},
	{Name: "RoleConfiguration", Flag: "role-configuration", Type: "*types.RoleConfiguration", Required: false},
	{Name: "ScrapeConfiguration", Flag: "scrape-configuration", Type: "types.ScrapeConfiguration", Required: true},
	{Name: "Source", Flag: "source", Type: "types.Source", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_workspace = []leanruntime.Field{
	{Name: "Alias", Flag: "alias", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "KmsKeyArn", Flag: "kms-key-arn", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_alert_manager_definition = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_delete_anomaly_detector = []leanruntime.Field{
	{Name: "AnomalyDetectorId", Flag: "anomaly-detector-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_delete_logging_configuration = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_delete_query_logging_configuration = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_delete_resource_policy = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "RevisionId", Flag: "revision-id", Type: "*string", Required: false},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_delete_rule_groups_namespace = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_delete_scraper = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ScraperId", Flag: "scraper-id", Type: "*string", Required: true},
}

var fields_delete_scraper_logging_configuration = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ScraperId", Flag: "scraper-id", Type: "*string", Required: true},
}

var fields_delete_workspace = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_describe_alert_manager_definition = []leanruntime.Field{
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_describe_anomaly_detector = []leanruntime.Field{
	{Name: "AnomalyDetectorId", Flag: "anomaly-detector-id", Type: "*string", Required: true},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_describe_logging_configuration = []leanruntime.Field{
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_describe_query_logging_configuration = []leanruntime.Field{
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_describe_resource_policy = []leanruntime.Field{
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_describe_rule_groups_namespace = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_describe_scraper = []leanruntime.Field{
	{Name: "ScraperId", Flag: "scraper-id", Type: "*string", Required: true},
}

var fields_describe_scraper_logging_configuration = []leanruntime.Field{
	{Name: "ScraperId", Flag: "scraper-id", Type: "*string", Required: true},
}

var fields_describe_workspace = []leanruntime.Field{
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_describe_workspace_configuration = []leanruntime.Field{
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_get_default_scraper_configuration = []leanruntime.Field{}

var fields_list_anomaly_detectors = []leanruntime.Field{
	{Name: "Alias", Flag: "alias", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_list_rule_groups_namespaces = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_list_scrapers = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "map[string][]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_workspaces = []leanruntime.Field{
	{Name: "Alias", Flag: "alias", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_put_alert_manager_definition = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Data", Flag: "data", Type: "[]byte", Required: true},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_put_anomaly_detector = []leanruntime.Field{
	{Name: "AnomalyDetectorId", Flag: "anomaly-detector-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Configuration", Flag: "configuration", Type: "types.AnomalyDetectorConfiguration", Required: true},
	{Name: "EvaluationIntervalInSeconds", Flag: "evaluation-interval-in-seconds", Type: "*int32", Required: false},
	{Name: "Labels", Flag: "labels", Type: "map[string]string", Required: false},
	{Name: "MissingDataAction", Flag: "missing-data-action", Type: "types.AnomalyDetectorMissingDataAction", Required: false},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_put_resource_policy = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "PolicyDocument", Flag: "policy-document", Type: "*string", Required: true},
	{Name: "RevisionId", Flag: "revision-id", Type: "*string", Required: false},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_put_rule_groups_namespace = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Data", Flag: "data", Type: "[]byte", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_logging_configuration = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "LogGroupArn", Flag: "log-group-arn", Type: "*string", Required: true},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_update_query_logging_configuration = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Destinations", Flag: "destinations", Type: "[]types.LoggingDestination", Required: true},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_update_scraper = []leanruntime.Field{
	{Name: "Alias", Flag: "alias", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Destination", Flag: "destination", Type: "types.Destination", Required: false},
	{Name: "RoleConfiguration", Flag: "role-configuration", Type: "*types.RoleConfiguration", Required: false},
	{Name: "ScrapeConfiguration", Flag: "scrape-configuration", Type: "types.ScrapeConfiguration", Required: false},
	{Name: "ScraperId", Flag: "scraper-id", Type: "*string", Required: true},
}

var fields_update_scraper_logging_configuration = []leanruntime.Field{
	{Name: "LoggingDestination", Flag: "logging-destination", Type: "types.ScraperLoggingDestination", Required: true},
	{Name: "ScraperComponents", Flag: "scraper-components", Type: "[]types.ScraperComponent", Required: false},
	{Name: "ScraperId", Flag: "scraper-id", Type: "*string", Required: true},
}

var fields_update_workspace_alias = []leanruntime.Field{
	{Name: "Alias", Flag: "alias", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_update_workspace_configuration = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "LimitsPerLabelSet", Flag: "limits-per-label-set", Type: "[]types.LimitsPerLabelSet", Required: false},
	{Name: "RetentionPeriodInDays", Flag: "retention-period-in-days", Type: "*int32", Required: false},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-alert-manager-definition": {
			Name:   "create-alert-manager-definition",
			Fields: fields_create_alert_manager_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAlertManagerDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_alert_manager_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAlertManagerDefinition(ctx, input)
			},
		},
		"create-anomaly-detector": {
			Name:   "create-anomaly-detector",
			Fields: fields_create_anomaly_detector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAnomalyDetectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_anomaly_detector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAnomalyDetector(ctx, input)
			},
		},
		"create-logging-configuration": {
			Name:   "create-logging-configuration",
			Fields: fields_create_logging_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLoggingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_logging_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLoggingConfiguration(ctx, input)
			},
		},
		"create-query-logging-configuration": {
			Name:   "create-query-logging-configuration",
			Fields: fields_create_query_logging_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateQueryLoggingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_query_logging_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateQueryLoggingConfiguration(ctx, input)
			},
		},
		"create-rule-groups-namespace": {
			Name:   "create-rule-groups-namespace",
			Fields: fields_create_rule_groups_namespace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRuleGroupsNamespaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_rule_groups_namespace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRuleGroupsNamespace(ctx, input)
			},
		},
		"create-scraper": {
			Name:   "create-scraper",
			Fields: fields_create_scraper,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateScraperInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_scraper, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateScraper(ctx, input)
			},
		},
		"create-workspace": {
			Name:   "create-workspace",
			Fields: fields_create_workspace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWorkspaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_workspace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWorkspace(ctx, input)
			},
		},
		"delete-alert-manager-definition": {
			Name:   "delete-alert-manager-definition",
			Fields: fields_delete_alert_manager_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAlertManagerDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_alert_manager_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAlertManagerDefinition(ctx, input)
			},
		},
		"delete-anomaly-detector": {
			Name:   "delete-anomaly-detector",
			Fields: fields_delete_anomaly_detector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAnomalyDetectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_anomaly_detector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAnomalyDetector(ctx, input)
			},
		},
		"delete-logging-configuration": {
			Name:   "delete-logging-configuration",
			Fields: fields_delete_logging_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLoggingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_logging_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLoggingConfiguration(ctx, input)
			},
		},
		"delete-query-logging-configuration": {
			Name:   "delete-query-logging-configuration",
			Fields: fields_delete_query_logging_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteQueryLoggingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_query_logging_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteQueryLoggingConfiguration(ctx, input)
			},
		},
		"delete-resource-policy": {
			Name:   "delete-resource-policy",
			Fields: fields_delete_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteResourcePolicy(ctx, input)
			},
		},
		"delete-rule-groups-namespace": {
			Name:   "delete-rule-groups-namespace",
			Fields: fields_delete_rule_groups_namespace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRuleGroupsNamespaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_rule_groups_namespace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRuleGroupsNamespace(ctx, input)
			},
		},
		"delete-scraper": {
			Name:   "delete-scraper",
			Fields: fields_delete_scraper,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteScraperInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_scraper, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteScraper(ctx, input)
			},
		},
		"delete-scraper-logging-configuration": {
			Name:   "delete-scraper-logging-configuration",
			Fields: fields_delete_scraper_logging_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteScraperLoggingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_scraper_logging_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteScraperLoggingConfiguration(ctx, input)
			},
		},
		"delete-workspace": {
			Name:   "delete-workspace",
			Fields: fields_delete_workspace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWorkspaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_workspace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWorkspace(ctx, input)
			},
		},
		"describe-alert-manager-definition": {
			Name:   "describe-alert-manager-definition",
			Fields: fields_describe_alert_manager_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAlertManagerDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_alert_manager_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAlertManagerDefinition(ctx, input)
			},
		},
		"describe-anomaly-detector": {
			Name:   "describe-anomaly-detector",
			Fields: fields_describe_anomaly_detector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAnomalyDetectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_anomaly_detector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAnomalyDetector(ctx, input)
			},
		},
		"describe-logging-configuration": {
			Name:   "describe-logging-configuration",
			Fields: fields_describe_logging_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLoggingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_logging_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeLoggingConfiguration(ctx, input)
			},
		},
		"describe-query-logging-configuration": {
			Name:   "describe-query-logging-configuration",
			Fields: fields_describe_query_logging_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeQueryLoggingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_query_logging_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeQueryLoggingConfiguration(ctx, input)
			},
		},
		"describe-resource-policy": {
			Name:   "describe-resource-policy",
			Fields: fields_describe_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeResourcePolicy(ctx, input)
			},
		},
		"describe-rule-groups-namespace": {
			Name:   "describe-rule-groups-namespace",
			Fields: fields_describe_rule_groups_namespace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRuleGroupsNamespaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_rule_groups_namespace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeRuleGroupsNamespace(ctx, input)
			},
		},
		"describe-scraper": {
			Name:   "describe-scraper",
			Fields: fields_describe_scraper,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeScraperInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_scraper, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeScraper(ctx, input)
			},
		},
		"describe-scraper-logging-configuration": {
			Name:   "describe-scraper-logging-configuration",
			Fields: fields_describe_scraper_logging_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeScraperLoggingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_scraper_logging_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeScraperLoggingConfiguration(ctx, input)
			},
		},
		"describe-workspace": {
			Name:   "describe-workspace",
			Fields: fields_describe_workspace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeWorkspaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_workspace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeWorkspace(ctx, input)
			},
		},
		"describe-workspace-configuration": {
			Name:   "describe-workspace-configuration",
			Fields: fields_describe_workspace_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeWorkspaceConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_workspace_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeWorkspaceConfiguration(ctx, input)
			},
		},
		"get-default-scraper-configuration": {
			Name:   "get-default-scraper-configuration",
			Fields: fields_get_default_scraper_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDefaultScraperConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_default_scraper_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDefaultScraperConfiguration(ctx, input)
			},
		},
		"list-anomaly-detectors": {
			Name:   "list-anomaly-detectors",
			Fields: fields_list_anomaly_detectors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAnomalyDetectorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_anomaly_detectors, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAnomalyDetectors(ctx, input)
				}
				var results []*svc.ListAnomalyDetectorsOutput
				p := svc.NewListAnomalyDetectorsPaginator(client, input)
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
		"list-rule-groups-namespaces": {
			Name:   "list-rule-groups-namespaces",
			Fields: fields_list_rule_groups_namespaces,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRuleGroupsNamespacesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_rule_groups_namespaces, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRuleGroupsNamespaces(ctx, input)
				}
				var results []*svc.ListRuleGroupsNamespacesOutput
				p := svc.NewListRuleGroupsNamespacesPaginator(client, input)
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
		"list-scrapers": {
			Name:   "list-scrapers",
			Fields: fields_list_scrapers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListScrapersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_scrapers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListScrapers(ctx, input)
				}
				var results []*svc.ListScrapersOutput
				p := svc.NewListScrapersPaginator(client, input)
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
		"list-workspaces": {
			Name:   "list-workspaces",
			Fields: fields_list_workspaces,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWorkspacesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_workspaces, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWorkspaces(ctx, input)
				}
				var results []*svc.ListWorkspacesOutput
				p := svc.NewListWorkspacesPaginator(client, input)
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
		"put-alert-manager-definition": {
			Name:   "put-alert-manager-definition",
			Fields: fields_put_alert_manager_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutAlertManagerDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_alert_manager_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutAlertManagerDefinition(ctx, input)
			},
		},
		"put-anomaly-detector": {
			Name:   "put-anomaly-detector",
			Fields: fields_put_anomaly_detector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutAnomalyDetectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_anomaly_detector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutAnomalyDetector(ctx, input)
			},
		},
		"put-resource-policy": {
			Name:   "put-resource-policy",
			Fields: fields_put_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutResourcePolicy(ctx, input)
			},
		},
		"put-rule-groups-namespace": {
			Name:   "put-rule-groups-namespace",
			Fields: fields_put_rule_groups_namespace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutRuleGroupsNamespaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_rule_groups_namespace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutRuleGroupsNamespace(ctx, input)
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
		"update-logging-configuration": {
			Name:   "update-logging-configuration",
			Fields: fields_update_logging_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateLoggingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_logging_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateLoggingConfiguration(ctx, input)
			},
		},
		"update-query-logging-configuration": {
			Name:   "update-query-logging-configuration",
			Fields: fields_update_query_logging_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateQueryLoggingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_query_logging_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateQueryLoggingConfiguration(ctx, input)
			},
		},
		"update-scraper": {
			Name:   "update-scraper",
			Fields: fields_update_scraper,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateScraperInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_scraper, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateScraper(ctx, input)
			},
		},
		"update-scraper-logging-configuration": {
			Name:   "update-scraper-logging-configuration",
			Fields: fields_update_scraper_logging_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateScraperLoggingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_scraper_logging_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateScraperLoggingConfiguration(ctx, input)
			},
		},
		"update-workspace-alias": {
			Name:   "update-workspace-alias",
			Fields: fields_update_workspace_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateWorkspaceAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_workspace_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateWorkspaceAlias(ctx, input)
			},
		},
		"update-workspace-configuration": {
			Name:   "update-workspace-configuration",
			Fields: fields_update_workspace_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateWorkspaceConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_workspace_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateWorkspaceConfiguration(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("amp", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
