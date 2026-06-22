package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/appintegrations"
)

var fields_create_application = []leanruntime.Field{
	{Name: "ApplicationConfig", Flag: "application-config", Type: "*types.ApplicationConfig", Required: false},
	{Name: "ApplicationSourceConfig", Flag: "application-source-config", Type: "*types.ApplicationSourceConfig", Required: true},
	{Name: "ApplicationType", Flag: "application-type", Type: "types.ApplicationType", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "IframeConfig", Flag: "iframe-config", Type: "*types.IframeConfig", Required: false},
	{Name: "InitializationTimeout", Flag: "initialization-timeout", Type: "*int32", Required: false},
	{Name: "IsService", Flag: "is-service", Type: "bool", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "Permissions", Flag: "permissions", Type: "[]string", Required: false},
	{Name: "Publications", Flag: "publications", Type: "[]types.Publication", Required: false},
	{Name: "Subscriptions", Flag: "subscriptions", Type: "[]types.Subscription", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_data_integration = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "FileConfiguration", Flag: "file-configuration", Type: "*types.FileConfiguration", Required: false},
	{Name: "KmsKey", Flag: "kms-key", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ObjectConfiguration", Flag: "object-configuration", Type: "map[string]map[string][]string", Required: false},
	{Name: "ScheduleConfig", Flag: "schedule-config", Type: "*types.ScheduleConfiguration", Required: false},
	{Name: "SourceURI", Flag: "source-uri", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_data_integration_association = []leanruntime.Field{
	{Name: "ClientAssociationMetadata", Flag: "client-association-metadata", Type: "map[string]string", Required: false},
	{Name: "ClientId", Flag: "client-id", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DataIntegrationIdentifier", Flag: "data-integration-identifier", Type: "*string", Required: true},
	{Name: "DestinationURI", Flag: "destination-uri", Type: "*string", Required: false},
	{Name: "ExecutionConfiguration", Flag: "execution-configuration", Type: "*types.ExecutionConfiguration", Required: false},
	{Name: "ObjectConfiguration", Flag: "object-configuration", Type: "map[string]map[string][]string", Required: false},
}

var fields_create_event_integration = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EventBridgeBus", Flag: "event-bridge-bus", Type: "*string", Required: true},
	{Name: "EventFilter", Flag: "event-filter", Type: "*types.EventFilter", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_application = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_delete_data_integration = []leanruntime.Field{
	{Name: "DataIntegrationIdentifier", Flag: "data-integration-identifier", Type: "*string", Required: true},
}

var fields_delete_event_integration = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_application = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_data_integration = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_event_integration = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_list_application_associations = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_applications = []leanruntime.Field{
	{Name: "ApplicationType", Flag: "application-type", Type: "types.ApplicationType", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_data_integration_associations = []leanruntime.Field{
	{Name: "DataIntegrationIdentifier", Flag: "data-integration-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_data_integrations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_event_integration_associations = []leanruntime.Field{
	{Name: "EventIntegrationName", Flag: "event-integration-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_event_integrations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_application = []leanruntime.Field{
	{Name: "ApplicationConfig", Flag: "application-config", Type: "*types.ApplicationConfig", Required: false},
	{Name: "ApplicationSourceConfig", Flag: "application-source-config", Type: "*types.ApplicationSourceConfig", Required: false},
	{Name: "ApplicationType", Flag: "application-type", Type: "types.ApplicationType", Required: false},
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "IframeConfig", Flag: "iframe-config", Type: "*types.IframeConfig", Required: false},
	{Name: "InitializationTimeout", Flag: "initialization-timeout", Type: "*int32", Required: false},
	{Name: "IsService", Flag: "is-service", Type: "*bool", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Permissions", Flag: "permissions", Type: "[]string", Required: false},
	{Name: "Publications", Flag: "publications", Type: "[]types.Publication", Required: false},
	{Name: "Subscriptions", Flag: "subscriptions", Type: "[]types.Subscription", Required: false},
}

var fields_update_data_integration = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_data_integration_association = []leanruntime.Field{
	{Name: "DataIntegrationAssociationIdentifier", Flag: "data-integration-association-identifier", Type: "*string", Required: true},
	{Name: "DataIntegrationIdentifier", Flag: "data-integration-identifier", Type: "*string", Required: true},
	{Name: "ExecutionConfiguration", Flag: "execution-configuration", Type: "*types.ExecutionConfiguration", Required: true},
}

var fields_update_event_integration = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-application": {
			Name:   "create-application",
			Fields: fields_create_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateApplication(ctx, input)
			},
		},
		"create-data-integration": {
			Name:   "create-data-integration",
			Fields: fields_create_data_integration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDataIntegrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_data_integration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDataIntegration(ctx, input)
			},
		},
		"create-data-integration-association": {
			Name:   "create-data-integration-association",
			Fields: fields_create_data_integration_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDataIntegrationAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_data_integration_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDataIntegrationAssociation(ctx, input)
			},
		},
		"create-event-integration": {
			Name:   "create-event-integration",
			Fields: fields_create_event_integration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEventIntegrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_event_integration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEventIntegration(ctx, input)
			},
		},
		"delete-application": {
			Name:   "delete-application",
			Fields: fields_delete_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteApplication(ctx, input)
			},
		},
		"delete-data-integration": {
			Name:   "delete-data-integration",
			Fields: fields_delete_data_integration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDataIntegrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_data_integration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDataIntegration(ctx, input)
			},
		},
		"delete-event-integration": {
			Name:   "delete-event-integration",
			Fields: fields_delete_event_integration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEventIntegrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_event_integration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEventIntegration(ctx, input)
			},
		},
		"get-application": {
			Name:   "get-application",
			Fields: fields_get_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetApplication(ctx, input)
			},
		},
		"get-data-integration": {
			Name:   "get-data-integration",
			Fields: fields_get_data_integration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDataIntegrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_data_integration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDataIntegration(ctx, input)
			},
		},
		"get-event-integration": {
			Name:   "get-event-integration",
			Fields: fields_get_event_integration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEventIntegrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_event_integration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEventIntegration(ctx, input)
			},
		},
		"list-application-associations": {
			Name:   "list-application-associations",
			Fields: fields_list_application_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListApplicationAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_application_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListApplicationAssociations(ctx, input)
				}
				var results []*svc.ListApplicationAssociationsOutput
				p := svc.NewListApplicationAssociationsPaginator(client, input)
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
		"list-applications": {
			Name:   "list-applications",
			Fields: fields_list_applications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListApplicationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_applications, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListApplications(ctx, input)
				}
				var results []*svc.ListApplicationsOutput
				p := svc.NewListApplicationsPaginator(client, input)
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
		"list-data-integration-associations": {
			Name:   "list-data-integration-associations",
			Fields: fields_list_data_integration_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDataIntegrationAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_data_integration_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDataIntegrationAssociations(ctx, input)
				}
				var results []*svc.ListDataIntegrationAssociationsOutput
				p := svc.NewListDataIntegrationAssociationsPaginator(client, input)
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
		"list-data-integrations": {
			Name:   "list-data-integrations",
			Fields: fields_list_data_integrations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDataIntegrationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_data_integrations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDataIntegrations(ctx, input)
				}
				var results []*svc.ListDataIntegrationsOutput
				p := svc.NewListDataIntegrationsPaginator(client, input)
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
		"list-event-integration-associations": {
			Name:   "list-event-integration-associations",
			Fields: fields_list_event_integration_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEventIntegrationAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_event_integration_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEventIntegrationAssociations(ctx, input)
				}
				var results []*svc.ListEventIntegrationAssociationsOutput
				p := svc.NewListEventIntegrationAssociationsPaginator(client, input)
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
		"list-event-integrations": {
			Name:   "list-event-integrations",
			Fields: fields_list_event_integrations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEventIntegrationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_event_integrations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEventIntegrations(ctx, input)
				}
				var results []*svc.ListEventIntegrationsOutput
				p := svc.NewListEventIntegrationsPaginator(client, input)
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
		"update-application": {
			Name:   "update-application",
			Fields: fields_update_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateApplication(ctx, input)
			},
		},
		"update-data-integration": {
			Name:   "update-data-integration",
			Fields: fields_update_data_integration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDataIntegrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_data_integration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDataIntegration(ctx, input)
			},
		},
		"update-data-integration-association": {
			Name:   "update-data-integration-association",
			Fields: fields_update_data_integration_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDataIntegrationAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_data_integration_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDataIntegrationAssociation(ctx, input)
			},
		},
		"update-event-integration": {
			Name:   "update-event-integration",
			Fields: fields_update_event_integration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEventIntegrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_event_integration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEventIntegration(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("appintegrations", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
