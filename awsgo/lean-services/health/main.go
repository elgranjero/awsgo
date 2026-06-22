package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/health"
)

var fields_describe_affected_accounts_for_organization = []leanruntime.Field{
	{Name: "EventArn", Flag: "event-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_affected_entities = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.EntityFilter", Required: true},
	{Name: "Locale", Flag: "locale", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_affected_entities_for_organization = []leanruntime.Field{
	{Name: "Locale", Flag: "locale", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OrganizationEntityAccountFilters", Flag: "organization-entity-account-filters", Type: "[]types.EntityAccountFilter", Required: false},
	{Name: "OrganizationEntityFilters", Flag: "organization-entity-filters", Type: "[]types.EventAccountFilter", Required: false},
}

var fields_describe_entity_aggregates = []leanruntime.Field{
	{Name: "EventArns", Flag: "event-arns", Type: "[]string", Required: false},
}

var fields_describe_entity_aggregates_for_organization = []leanruntime.Field{
	{Name: "AwsAccountIds", Flag: "aws-account-ids", Type: "[]string", Required: false},
	{Name: "EventArns", Flag: "event-arns", Type: "[]string", Required: true},
}

var fields_describe_event_aggregates = []leanruntime.Field{
	{Name: "AggregateField", Flag: "aggregate-field", Type: "types.EventAggregateField", Required: true},
	{Name: "Filter", Flag: "filter", Type: "*types.EventFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_event_details = []leanruntime.Field{
	{Name: "EventArns", Flag: "event-arns", Type: "[]string", Required: true},
	{Name: "Locale", Flag: "locale", Type: "*string", Required: false},
}

var fields_describe_event_details_for_organization = []leanruntime.Field{
	{Name: "Locale", Flag: "locale", Type: "*string", Required: false},
	{Name: "OrganizationEventDetailFilters", Flag: "organization-event-detail-filters", Type: "[]types.EventAccountFilter", Required: true},
}

var fields_describe_event_types = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.EventTypeFilter", Required: false},
	{Name: "Locale", Flag: "locale", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_events = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.EventFilter", Required: false},
	{Name: "Locale", Flag: "locale", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_events_for_organization = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.OrganizationEventFilter", Required: false},
	{Name: "Locale", Flag: "locale", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_health_service_status_for_organization = []leanruntime.Field{}

var fields_disable_health_service_access_for_organization = []leanruntime.Field{}

var fields_enable_health_service_access_for_organization = []leanruntime.Field{}

func main() {
	ops := map[string]leanruntime.Operation{
		"describe-affected-accounts-for-organization": {
			Name:   "describe-affected-accounts-for-organization",
			Fields: fields_describe_affected_accounts_for_organization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAffectedAccountsForOrganizationInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_affected_accounts_for_organization, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeAffectedAccountsForOrganization(ctx, input)
				}
				var results []*svc.DescribeAffectedAccountsForOrganizationOutput
				p := svc.NewDescribeAffectedAccountsForOrganizationPaginator(client, input)
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
		"describe-affected-entities": {
			Name:   "describe-affected-entities",
			Fields: fields_describe_affected_entities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAffectedEntitiesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_affected_entities, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeAffectedEntities(ctx, input)
				}
				var results []*svc.DescribeAffectedEntitiesOutput
				p := svc.NewDescribeAffectedEntitiesPaginator(client, input)
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
		"describe-affected-entities-for-organization": {
			Name:   "describe-affected-entities-for-organization",
			Fields: fields_describe_affected_entities_for_organization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAffectedEntitiesForOrganizationInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_affected_entities_for_organization, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeAffectedEntitiesForOrganization(ctx, input)
				}
				var results []*svc.DescribeAffectedEntitiesForOrganizationOutput
				p := svc.NewDescribeAffectedEntitiesForOrganizationPaginator(client, input)
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
		"describe-entity-aggregates": {
			Name:   "describe-entity-aggregates",
			Fields: fields_describe_entity_aggregates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEntityAggregatesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_entity_aggregates, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEntityAggregates(ctx, input)
			},
		},
		"describe-entity-aggregates-for-organization": {
			Name:   "describe-entity-aggregates-for-organization",
			Fields: fields_describe_entity_aggregates_for_organization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEntityAggregatesForOrganizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_entity_aggregates_for_organization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEntityAggregatesForOrganization(ctx, input)
			},
		},
		"describe-event-aggregates": {
			Name:   "describe-event-aggregates",
			Fields: fields_describe_event_aggregates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEventAggregatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_event_aggregates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeEventAggregates(ctx, input)
				}
				var results []*svc.DescribeEventAggregatesOutput
				p := svc.NewDescribeEventAggregatesPaginator(client, input)
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
		"describe-event-details": {
			Name:   "describe-event-details",
			Fields: fields_describe_event_details,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEventDetailsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_event_details, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEventDetails(ctx, input)
			},
		},
		"describe-event-details-for-organization": {
			Name:   "describe-event-details-for-organization",
			Fields: fields_describe_event_details_for_organization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEventDetailsForOrganizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_event_details_for_organization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEventDetailsForOrganization(ctx, input)
			},
		},
		"describe-event-types": {
			Name:   "describe-event-types",
			Fields: fields_describe_event_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEventTypesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_event_types, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeEventTypes(ctx, input)
				}
				var results []*svc.DescribeEventTypesOutput
				p := svc.NewDescribeEventTypesPaginator(client, input)
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
		"describe-events": {
			Name:   "describe-events",
			Fields: fields_describe_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEventsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_events, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeEvents(ctx, input)
				}
				var results []*svc.DescribeEventsOutput
				p := svc.NewDescribeEventsPaginator(client, input)
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
		"describe-events-for-organization": {
			Name:   "describe-events-for-organization",
			Fields: fields_describe_events_for_organization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEventsForOrganizationInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_events_for_organization, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeEventsForOrganization(ctx, input)
				}
				var results []*svc.DescribeEventsForOrganizationOutput
				p := svc.NewDescribeEventsForOrganizationPaginator(client, input)
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
		"describe-health-service-status-for-organization": {
			Name:   "describe-health-service-status-for-organization",
			Fields: fields_describe_health_service_status_for_organization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeHealthServiceStatusForOrganizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_health_service_status_for_organization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeHealthServiceStatusForOrganization(ctx, input)
			},
		},
		"disable-health-service-access-for-organization": {
			Name:   "disable-health-service-access-for-organization",
			Fields: fields_disable_health_service_access_for_organization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableHealthServiceAccessForOrganizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_health_service_access_for_organization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableHealthServiceAccessForOrganization(ctx, input)
			},
		},
		"enable-health-service-access-for-organization": {
			Name:   "enable-health-service-access-for-organization",
			Fields: fields_enable_health_service_access_for_organization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableHealthServiceAccessForOrganizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_health_service_access_for_organization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableHealthServiceAccessForOrganization(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("health", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
