package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/health/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"describe-affected-accounts-for-organization", "describe-affected-entities", "describe-affected-entities-for-organization", "describe-entity-aggregates", "describe-entity-aggregates-for-organization", "describe-event-aggregates", "describe-event-details", "describe-event-details-for-organization", "describe-event-types", "describe-events", "describe-events-for-organization", "describe-health-service-status-for-organization", "disable-health-service-access-for-organization", "enable-health-service-access-for-organization"},
		OperationSet: map[string]bool{"describe-affected-accounts-for-organization": true, "describe-affected-entities": true, "describe-affected-entities-for-organization": true, "describe-entity-aggregates": true, "describe-entity-aggregates-for-organization": true, "describe-event-aggregates": true, "describe-event-details": true, "describe-event-details-for-organization": true, "describe-event-types": true, "describe-events": true, "describe-events-for-organization": true, "describe-health-service-status-for-organization": true, "disable-health-service-access-for-organization": true, "enable-health-service-access-for-organization": true},
		OperationInputs: map[string][]string{
			"describe-affected-accounts-for-organization":     {"EventArn", "MaxResults", "NextToken"},
			"describe-affected-entities":                      {"Filter", "Locale", "MaxResults", "NextToken"},
			"describe-affected-entities-for-organization":     {"Locale", "MaxResults", "NextToken", "OrganizationEntityAccountFilters", "OrganizationEntityFilters"},
			"describe-entity-aggregates":                      {"EventArns"},
			"describe-entity-aggregates-for-organization":     {"AwsAccountIds", "EventArns"},
			"describe-event-aggregates":                       {"AggregateField", "Filter", "MaxResults", "NextToken"},
			"describe-event-details":                          {"EventArns", "Locale"},
			"describe-event-details-for-organization":         {"Locale", "OrganizationEventDetailFilters"},
			"describe-event-types":                            {"Filter", "Locale", "MaxResults", "NextToken"},
			"describe-events":                                 {"Filter", "Locale", "MaxResults", "NextToken"},
			"describe-events-for-organization":                {"Filter", "Locale", "MaxResults", "NextToken"},
			"describe-health-service-status-for-organization": {},
			"disable-health-service-access-for-organization":  {},
			"enable-health-service-access-for-organization":   {},
		},
		OperationInputTypes: map[string]map[string]string{
			"describe-affected-accounts-for-organization":     {"EventArn": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"describe-affected-entities":                      {"Filter": "*types.EntityFilter", "Locale": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"describe-affected-entities-for-organization":     {"Locale": "*string", "MaxResults": "*int32", "NextToken": "*string", "OrganizationEntityAccountFilters": "[]types.EntityAccountFilter", "OrganizationEntityFilters": "[]types.EventAccountFilter"},
			"describe-entity-aggregates":                      {"EventArns": "[]string"},
			"describe-entity-aggregates-for-organization":     {"AwsAccountIds": "[]string", "EventArns": "[]string"},
			"describe-event-aggregates":                       {"AggregateField": "types.EventAggregateField", "Filter": "*types.EventFilter", "MaxResults": "*int32", "NextToken": "*string"},
			"describe-event-details":                          {"EventArns": "[]string", "Locale": "*string"},
			"describe-event-details-for-organization":         {"Locale": "*string", "OrganizationEventDetailFilters": "[]types.EventAccountFilter"},
			"describe-event-types":                            {"Filter": "*types.EventTypeFilter", "Locale": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"describe-events":                                 {"Filter": "*types.EventFilter", "Locale": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"describe-events-for-organization":                {"Filter": "*types.OrganizationEventFilter", "Locale": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"describe-health-service-status-for-organization": {},
			"disable-health-service-access-for-organization":  {},
			"enable-health-service-access-for-organization":   {},
		},
		OperationInputRequired: map[string][]string{
			"describe-affected-accounts-for-organization":     {"EventArn"},
			"describe-affected-entities":                      {"Filter"},
			"describe-affected-entities-for-organization":     {},
			"describe-entity-aggregates":                      {},
			"describe-entity-aggregates-for-organization":     {"EventArns"},
			"describe-event-aggregates":                       {"AggregateField"},
			"describe-event-details":                          {"EventArns"},
			"describe-event-details-for-organization":         {"OrganizationEventDetailFilters"},
			"describe-event-types":                            {},
			"describe-events":                                 {},
			"describe-events-for-organization":                {},
			"describe-health-service-status-for-organization": {},
			"disable-health-service-access-for-organization":  {},
			"enable-health-service-access-for-organization":   {},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("health", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
