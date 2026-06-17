package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/applicationsignals/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"batch-get-service-level-objective-budget-report", "batch-update-exclusion-windows", "create-service-level-objective", "delete-grouping-configuration", "delete-service-level-objective", "get-service", "get-service-level-objective", "list-audit-findings", "list-entity-events", "list-grouping-attribute-definitions", "list-service-dependencies", "list-service-dependents", "list-service-level-objective-exclusion-windows", "list-service-level-objectives", "list-service-operations", "list-service-states", "list-services", "list-tags-for-resource", "put-grouping-configuration", "start-discovery", "tag-resource", "untag-resource", "update-service-level-objective"},
		OperationSet: map[string]bool{"batch-get-service-level-objective-budget-report": true, "batch-update-exclusion-windows": true, "create-service-level-objective": true, "delete-grouping-configuration": true, "delete-service-level-objective": true, "get-service": true, "get-service-level-objective": true, "list-audit-findings": true, "list-entity-events": true, "list-grouping-attribute-definitions": true, "list-service-dependencies": true, "list-service-dependents": true, "list-service-level-objective-exclusion-windows": true, "list-service-level-objectives": true, "list-service-operations": true, "list-service-states": true, "list-services": true, "list-tags-for-resource": true, "put-grouping-configuration": true, "start-discovery": true, "tag-resource": true, "untag-resource": true, "update-service-level-objective": true},
		OperationInputs: map[string][]string{
			"batch-get-service-level-objective-budget-report": {"SloIds", "Timestamp"},
			"batch-update-exclusion-windows":                  {"AddExclusionWindows", "RemoveExclusionWindows", "SloIds"},
			"create-service-level-objective":                  {"BurnRateConfigurations", "Description", "Goal", "Name", "RequestBasedSliConfig", "SliConfig", "Tags"},
			"delete-grouping-configuration":                   {},
			"delete-service-level-objective":                  {"Id"},
			"get-service":                                     {"EndTime", "KeyAttributes", "StartTime"},
			"get-service-level-objective":                     {"Id"},
			"list-audit-findings":                             {"AuditTargets", "Auditors", "DetailLevel", "EndTime", "MaxResults", "NextToken", "StartTime"},
			"list-entity-events":                              {"EndTime", "Entity", "MaxResults", "NextToken", "StartTime"},
			"list-grouping-attribute-definitions":             {"AwsAccountId", "IncludeLinkedAccounts", "NextToken"},
			"list-service-dependencies":                       {"EndTime", "KeyAttributes", "MaxResults", "NextToken", "StartTime"},
			"list-service-dependents":                         {"EndTime", "KeyAttributes", "MaxResults", "NextToken", "StartTime"},
			"list-service-level-objective-exclusion-windows":  {"Id", "MaxResults", "NextToken"},
			"list-service-level-objectives":                   {"DependencyConfig", "IncludeLinkedAccounts", "KeyAttributes", "MaxResults", "MetricSourceTypes", "NextToken", "OperationName", "SloOwnerAwsAccountId"},
			"list-service-operations":                         {"EndTime", "KeyAttributes", "MaxResults", "NextToken", "StartTime"},
			"list-service-states":                             {"AttributeFilters", "AwsAccountId", "EndTime", "IncludeLinkedAccounts", "MaxResults", "NextToken", "StartTime"},
			"list-services":                                   {"AwsAccountId", "EndTime", "IncludeLinkedAccounts", "MaxResults", "NextToken", "StartTime"},
			"list-tags-for-resource":                          {"ResourceArn"},
			"put-grouping-configuration":                      {"GroupingAttributeDefinitions"},
			"start-discovery":                                 {},
			"tag-resource":                                    {"ResourceArn", "Tags"},
			"untag-resource":                                  {"ResourceArn", "TagKeys"},
			"update-service-level-objective":                  {"BurnRateConfigurations", "Description", "Goal", "Id", "RequestBasedSliConfig", "SliConfig"},
		},
		OperationInputTypes: map[string]map[string]string{
			"batch-get-service-level-objective-budget-report": {"SloIds": "[]string", "Timestamp": "*time.Time"},
			"batch-update-exclusion-windows":                  {"AddExclusionWindows": "[]types.ExclusionWindow", "RemoveExclusionWindows": "[]types.ExclusionWindow", "SloIds": "[]string"},
			"create-service-level-objective":                  {"BurnRateConfigurations": "[]types.BurnRateConfiguration", "Description": "*string", "Goal": "*types.Goal", "Name": "*string", "RequestBasedSliConfig": "*types.RequestBasedServiceLevelIndicatorConfig", "SliConfig": "*types.ServiceLevelIndicatorConfig", "Tags": "[]types.Tag"},
			"delete-grouping-configuration":                   {},
			"delete-service-level-objective":                  {"Id": "*string"},
			"get-service":                                     {"EndTime": "*time.Time", "KeyAttributes": "map[string]string", "StartTime": "*time.Time"},
			"get-service-level-objective":                     {"Id": "*string"},
			"list-audit-findings":                             {"AuditTargets": "[]types.AuditTarget", "Auditors": "[]string", "DetailLevel": "types.DetailLevel", "EndTime": "*time.Time", "MaxResults": "*int32", "NextToken": "*string", "StartTime": "*time.Time"},
			"list-entity-events":                              {"EndTime": "*time.Time", "Entity": "map[string]string", "MaxResults": "*int32", "NextToken": "*string", "StartTime": "*time.Time"},
			"list-grouping-attribute-definitions":             {"AwsAccountId": "*string", "IncludeLinkedAccounts": "bool", "NextToken": "*string"},
			"list-service-dependencies":                       {"EndTime": "*time.Time", "KeyAttributes": "map[string]string", "MaxResults": "*int32", "NextToken": "*string", "StartTime": "*time.Time"},
			"list-service-dependents":                         {"EndTime": "*time.Time", "KeyAttributes": "map[string]string", "MaxResults": "*int32", "NextToken": "*string", "StartTime": "*time.Time"},
			"list-service-level-objective-exclusion-windows":  {"Id": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-service-level-objectives":                   {"DependencyConfig": "*types.DependencyConfig", "IncludeLinkedAccounts": "bool", "KeyAttributes": "map[string]string", "MaxResults": "*int32", "MetricSourceTypes": "[]types.MetricSourceType", "NextToken": "*string", "OperationName": "*string", "SloOwnerAwsAccountId": "*string"},
			"list-service-operations":                         {"EndTime": "*time.Time", "KeyAttributes": "map[string]string", "MaxResults": "*int32", "NextToken": "*string", "StartTime": "*time.Time"},
			"list-service-states":                             {"AttributeFilters": "[]types.AttributeFilter", "AwsAccountId": "*string", "EndTime": "*time.Time", "IncludeLinkedAccounts": "bool", "MaxResults": "*int32", "NextToken": "*string", "StartTime": "*time.Time"},
			"list-services":                                   {"AwsAccountId": "*string", "EndTime": "*time.Time", "IncludeLinkedAccounts": "bool", "MaxResults": "*int32", "NextToken": "*string", "StartTime": "*time.Time"},
			"list-tags-for-resource":                          {"ResourceArn": "*string"},
			"put-grouping-configuration":                      {"GroupingAttributeDefinitions": "[]types.GroupingAttributeDefinition"},
			"start-discovery":                                 {},
			"tag-resource":                                    {"ResourceArn": "*string", "Tags": "[]types.Tag"},
			"untag-resource":                                  {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-service-level-objective":                  {"BurnRateConfigurations": "[]types.BurnRateConfiguration", "Description": "*string", "Goal": "*types.Goal", "Id": "*string", "RequestBasedSliConfig": "*types.RequestBasedServiceLevelIndicatorConfig", "SliConfig": "*types.ServiceLevelIndicatorConfig"},
		},
		OperationInputRequired: map[string][]string{
			"batch-get-service-level-objective-budget-report": {"SloIds", "Timestamp"},
			"batch-update-exclusion-windows":                  {"SloIds"},
			"create-service-level-objective":                  {"Name"},
			"delete-grouping-configuration":                   {},
			"delete-service-level-objective":                  {"Id"},
			"get-service":                                     {"EndTime", "KeyAttributes", "StartTime"},
			"get-service-level-objective":                     {"Id"},
			"list-audit-findings":                             {"AuditTargets", "EndTime", "StartTime"},
			"list-entity-events":                              {"EndTime", "Entity", "StartTime"},
			"list-grouping-attribute-definitions":             {},
			"list-service-dependencies":                       {"EndTime", "KeyAttributes", "StartTime"},
			"list-service-dependents":                         {"EndTime", "KeyAttributes", "StartTime"},
			"list-service-level-objective-exclusion-windows":  {"Id"},
			"list-service-level-objectives":                   {},
			"list-service-operations":                         {"EndTime", "KeyAttributes", "StartTime"},
			"list-service-states":                             {"EndTime", "StartTime"},
			"list-services":                                   {"EndTime", "StartTime"},
			"list-tags-for-resource":                          {"ResourceArn"},
			"put-grouping-configuration":                      {"GroupingAttributeDefinitions"},
			"start-discovery":                                 {},
			"tag-resource":                                    {"ResourceArn", "Tags"},
			"untag-resource":                                  {"ResourceArn", "TagKeys"},
			"update-service-level-objective":                  {"Id"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("applicationsignals", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
