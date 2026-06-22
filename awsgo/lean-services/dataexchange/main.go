package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/dataexchange"
)

var fields_accept_data_grant = []leanruntime.Field{
	{Name: "DataGrantArn", Flag: "data-grant-arn", Type: "*string", Required: true},
}

var fields_cancel_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_create_data_grant = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EndsAt", Flag: "ends-at", Type: "*time.Time", Required: false},
	{Name: "GrantDistributionScope", Flag: "grant-distribution-scope", Type: "types.GrantDistributionScope", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ReceiverPrincipal", Flag: "receiver-principal", Type: "*string", Required: true},
	{Name: "SourceDataSetId", Flag: "source-data-set-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_data_set = []leanruntime.Field{
	{Name: "AssetType", Flag: "asset-type", Type: "types.AssetType", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_event_action = []leanruntime.Field{
	{Name: "Action", Flag: "action", Type: "*types.Action", Required: true},
	{Name: "Event", Flag: "event", Type: "*types.Event", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_job = []leanruntime.Field{
	{Name: "Details", Flag: "details", Type: "*types.RequestDetails", Required: true},
	{Name: "Type", Flag: "type", Type: "types.Type", Required: true},
}

var fields_create_revision = []leanruntime.Field{
	{Name: "Comment", Flag: "comment", Type: "*string", Required: false},
	{Name: "DataSetId", Flag: "data-set-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_asset = []leanruntime.Field{
	{Name: "AssetId", Flag: "asset-id", Type: "*string", Required: true},
	{Name: "DataSetId", Flag: "data-set-id", Type: "*string", Required: true},
	{Name: "RevisionId", Flag: "revision-id", Type: "*string", Required: true},
}

var fields_delete_data_grant = []leanruntime.Field{
	{Name: "DataGrantId", Flag: "data-grant-id", Type: "*string", Required: true},
}

var fields_delete_data_set = []leanruntime.Field{
	{Name: "DataSetId", Flag: "data-set-id", Type: "*string", Required: true},
}

var fields_delete_event_action = []leanruntime.Field{
	{Name: "EventActionId", Flag: "event-action-id", Type: "*string", Required: true},
}

var fields_delete_revision = []leanruntime.Field{
	{Name: "DataSetId", Flag: "data-set-id", Type: "*string", Required: true},
	{Name: "RevisionId", Flag: "revision-id", Type: "*string", Required: true},
}

var fields_get_asset = []leanruntime.Field{
	{Name: "AssetId", Flag: "asset-id", Type: "*string", Required: true},
	{Name: "DataSetId", Flag: "data-set-id", Type: "*string", Required: true},
	{Name: "RevisionId", Flag: "revision-id", Type: "*string", Required: true},
}

var fields_get_data_grant = []leanruntime.Field{
	{Name: "DataGrantId", Flag: "data-grant-id", Type: "*string", Required: true},
}

var fields_get_data_set = []leanruntime.Field{
	{Name: "DataSetId", Flag: "data-set-id", Type: "*string", Required: true},
}

var fields_get_event_action = []leanruntime.Field{
	{Name: "EventActionId", Flag: "event-action-id", Type: "*string", Required: true},
}

var fields_get_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_get_received_data_grant = []leanruntime.Field{
	{Name: "DataGrantArn", Flag: "data-grant-arn", Type: "*string", Required: true},
}

var fields_get_revision = []leanruntime.Field{
	{Name: "DataSetId", Flag: "data-set-id", Type: "*string", Required: true},
	{Name: "RevisionId", Flag: "revision-id", Type: "*string", Required: true},
}

var fields_list_data_grants = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_data_set_revisions = []leanruntime.Field{
	{Name: "DataSetId", Flag: "data-set-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_data_sets = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Origin", Flag: "origin", Type: "*string", Required: false},
}

var fields_list_event_actions = []leanruntime.Field{
	{Name: "EventSourceId", Flag: "event-source-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_jobs = []leanruntime.Field{
	{Name: "DataSetId", Flag: "data-set-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RevisionId", Flag: "revision-id", Type: "*string", Required: false},
}

var fields_list_received_data_grants = []leanruntime.Field{
	{Name: "AcceptanceState", Flag: "acceptance-state", Type: "[]types.AcceptanceStateFilterValue", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_revision_assets = []leanruntime.Field{
	{Name: "DataSetId", Flag: "data-set-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RevisionId", Flag: "revision-id", Type: "*string", Required: true},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_revoke_revision = []leanruntime.Field{
	{Name: "DataSetId", Flag: "data-set-id", Type: "*string", Required: true},
	{Name: "RevisionId", Flag: "revision-id", Type: "*string", Required: true},
	{Name: "RevocationComment", Flag: "revocation-comment", Type: "*string", Required: true},
}

var fields_send_api_asset = []leanruntime.Field{
	{Name: "AssetId", Flag: "asset-id", Type: "*string", Required: true},
	{Name: "Body", Flag: "body", Type: "*string", Required: false},
	{Name: "DataSetId", Flag: "data-set-id", Type: "*string", Required: true},
	{Name: "Method", Flag: "method", Type: "*string", Required: false},
	{Name: "Path", Flag: "path", Type: "*string", Required: false},
	{Name: "QueryStringParameters", Flag: "query-string-parameters", Type: "map[string]string", Required: false},
	{Name: "RequestHeaders", Flag: "request-headers", Type: "map[string]string", Required: false},
	{Name: "RevisionId", Flag: "revision-id", Type: "*string", Required: true},
}

var fields_send_data_set_notification = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Comment", Flag: "comment", Type: "*string", Required: false},
	{Name: "DataSetId", Flag: "data-set-id", Type: "*string", Required: true},
	{Name: "Details", Flag: "details", Type: "*types.NotificationDetails", Required: false},
	{Name: "Scope", Flag: "scope", Type: "*types.ScopeDetails", Required: false},
	{Name: "Type", Flag: "type", Type: "types.NotificationType", Required: true},
}

var fields_start_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_asset = []leanruntime.Field{
	{Name: "AssetId", Flag: "asset-id", Type: "*string", Required: true},
	{Name: "DataSetId", Flag: "data-set-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RevisionId", Flag: "revision-id", Type: "*string", Required: true},
}

var fields_update_data_set = []leanruntime.Field{
	{Name: "DataSetId", Flag: "data-set-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_event_action = []leanruntime.Field{
	{Name: "Action", Flag: "action", Type: "*types.Action", Required: false},
	{Name: "EventActionId", Flag: "event-action-id", Type: "*string", Required: true},
}

var fields_update_revision = []leanruntime.Field{
	{Name: "Comment", Flag: "comment", Type: "*string", Required: false},
	{Name: "DataSetId", Flag: "data-set-id", Type: "*string", Required: true},
	{Name: "Finalized", Flag: "finalized", Type: "*bool", Required: false},
	{Name: "RevisionId", Flag: "revision-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"accept-data-grant": {
			Name:   "accept-data-grant",
			Fields: fields_accept_data_grant,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AcceptDataGrantInput{}
				if _, err := leanruntime.ApplyInput(input, fields_accept_data_grant, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AcceptDataGrant(ctx, input)
			},
		},
		"cancel-job": {
			Name:   "cancel-job",
			Fields: fields_cancel_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelJob(ctx, input)
			},
		},
		"create-data-grant": {
			Name:   "create-data-grant",
			Fields: fields_create_data_grant,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDataGrantInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_data_grant, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDataGrant(ctx, input)
			},
		},
		"create-data-set": {
			Name:   "create-data-set",
			Fields: fields_create_data_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDataSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_data_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDataSet(ctx, input)
			},
		},
		"create-event-action": {
			Name:   "create-event-action",
			Fields: fields_create_event_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEventActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_event_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEventAction(ctx, input)
			},
		},
		"create-job": {
			Name:   "create-job",
			Fields: fields_create_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateJob(ctx, input)
			},
		},
		"create-revision": {
			Name:   "create-revision",
			Fields: fields_create_revision,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRevisionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_revision, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRevision(ctx, input)
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
		"delete-data-grant": {
			Name:   "delete-data-grant",
			Fields: fields_delete_data_grant,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDataGrantInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_data_grant, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDataGrant(ctx, input)
			},
		},
		"delete-data-set": {
			Name:   "delete-data-set",
			Fields: fields_delete_data_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDataSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_data_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDataSet(ctx, input)
			},
		},
		"delete-event-action": {
			Name:   "delete-event-action",
			Fields: fields_delete_event_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEventActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_event_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEventAction(ctx, input)
			},
		},
		"delete-revision": {
			Name:   "delete-revision",
			Fields: fields_delete_revision,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRevisionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_revision, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRevision(ctx, input)
			},
		},
		"get-asset": {
			Name:   "get-asset",
			Fields: fields_get_asset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAssetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_asset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAsset(ctx, input)
			},
		},
		"get-data-grant": {
			Name:   "get-data-grant",
			Fields: fields_get_data_grant,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDataGrantInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_data_grant, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDataGrant(ctx, input)
			},
		},
		"get-data-set": {
			Name:   "get-data-set",
			Fields: fields_get_data_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDataSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_data_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDataSet(ctx, input)
			},
		},
		"get-event-action": {
			Name:   "get-event-action",
			Fields: fields_get_event_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEventActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_event_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEventAction(ctx, input)
			},
		},
		"get-job": {
			Name:   "get-job",
			Fields: fields_get_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetJob(ctx, input)
			},
		},
		"get-received-data-grant": {
			Name:   "get-received-data-grant",
			Fields: fields_get_received_data_grant,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetReceivedDataGrantInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_received_data_grant, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetReceivedDataGrant(ctx, input)
			},
		},
		"get-revision": {
			Name:   "get-revision",
			Fields: fields_get_revision,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRevisionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_revision, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRevision(ctx, input)
			},
		},
		"list-data-grants": {
			Name:   "list-data-grants",
			Fields: fields_list_data_grants,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDataGrantsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_data_grants, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDataGrants(ctx, input)
				}
				var results []*svc.ListDataGrantsOutput
				p := svc.NewListDataGrantsPaginator(client, input)
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
		"list-data-set-revisions": {
			Name:   "list-data-set-revisions",
			Fields: fields_list_data_set_revisions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDataSetRevisionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_data_set_revisions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDataSetRevisions(ctx, input)
				}
				var results []*svc.ListDataSetRevisionsOutput
				p := svc.NewListDataSetRevisionsPaginator(client, input)
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
		"list-data-sets": {
			Name:   "list-data-sets",
			Fields: fields_list_data_sets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDataSetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_data_sets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDataSets(ctx, input)
				}
				var results []*svc.ListDataSetsOutput
				p := svc.NewListDataSetsPaginator(client, input)
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
		"list-event-actions": {
			Name:   "list-event-actions",
			Fields: fields_list_event_actions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEventActionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_event_actions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEventActions(ctx, input)
				}
				var results []*svc.ListEventActionsOutput
				p := svc.NewListEventActionsPaginator(client, input)
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
		"list-jobs": {
			Name:   "list-jobs",
			Fields: fields_list_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListJobs(ctx, input)
				}
				var results []*svc.ListJobsOutput
				p := svc.NewListJobsPaginator(client, input)
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
		"list-received-data-grants": {
			Name:   "list-received-data-grants",
			Fields: fields_list_received_data_grants,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListReceivedDataGrantsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_received_data_grants, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListReceivedDataGrants(ctx, input)
				}
				var results []*svc.ListReceivedDataGrantsOutput
				p := svc.NewListReceivedDataGrantsPaginator(client, input)
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
		"list-revision-assets": {
			Name:   "list-revision-assets",
			Fields: fields_list_revision_assets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRevisionAssetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_revision_assets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRevisionAssets(ctx, input)
				}
				var results []*svc.ListRevisionAssetsOutput
				p := svc.NewListRevisionAssetsPaginator(client, input)
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
		"revoke-revision": {
			Name:   "revoke-revision",
			Fields: fields_revoke_revision,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RevokeRevisionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_revoke_revision, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RevokeRevision(ctx, input)
			},
		},
		"send-api-asset": {
			Name:   "send-api-asset",
			Fields: fields_send_api_asset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendApiAssetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_api_asset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendApiAsset(ctx, input)
			},
		},
		"send-data-set-notification": {
			Name:   "send-data-set-notification",
			Fields: fields_send_data_set_notification,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendDataSetNotificationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_data_set_notification, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendDataSetNotification(ctx, input)
			},
		},
		"start-job": {
			Name:   "start-job",
			Fields: fields_start_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartJob(ctx, input)
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
		"update-data-set": {
			Name:   "update-data-set",
			Fields: fields_update_data_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDataSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_data_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDataSet(ctx, input)
			},
		},
		"update-event-action": {
			Name:   "update-event-action",
			Fields: fields_update_event_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEventActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_event_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEventAction(ctx, input)
			},
		},
		"update-revision": {
			Name:   "update-revision",
			Fields: fields_update_revision,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRevisionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_revision, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRevision(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("dataexchange", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
