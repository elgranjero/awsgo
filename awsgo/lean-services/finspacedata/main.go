package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/finspacedata"
)

var fields_associate_user_to_permission_group = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "PermissionGroupId", Flag: "permission-group-id", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_create_changeset = []leanruntime.Field{
	{Name: "ChangeType", Flag: "change-type", Type: "types.ChangeType", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DatasetId", Flag: "dataset-id", Type: "*string", Required: true},
	{Name: "FormatParams", Flag: "format-params", Type: "map[string]string", Required: true},
	{Name: "SourceParams", Flag: "source-params", Type: "map[string]string", Required: true},
}

var fields_create_data_view = []leanruntime.Field{
	{Name: "AsOfTimestamp", Flag: "as-of-timestamp", Type: "*int64", Required: false},
	{Name: "AutoUpdate", Flag: "auto-update", Type: "bool", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DatasetId", Flag: "dataset-id", Type: "*string", Required: true},
	{Name: "DestinationTypeParams", Flag: "destination-type-params", Type: "*types.DataViewDestinationTypeParams", Required: true},
	{Name: "PartitionColumns", Flag: "partition-columns", Type: "[]string", Required: false},
	{Name: "SortColumns", Flag: "sort-columns", Type: "[]string", Required: false},
}

var fields_create_dataset = []leanruntime.Field{
	{Name: "Alias", Flag: "alias", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DatasetDescription", Flag: "dataset-description", Type: "*string", Required: false},
	{Name: "DatasetTitle", Flag: "dataset-title", Type: "*string", Required: true},
	{Name: "Kind", Flag: "kind", Type: "types.DatasetKind", Required: true},
	{Name: "OwnerInfo", Flag: "owner-info", Type: "*types.DatasetOwnerInfo", Required: false},
	{Name: "PermissionGroupParams", Flag: "permission-group-params", Type: "*types.PermissionGroupParams", Required: true},
	{Name: "SchemaDefinition", Flag: "schema-definition", Type: "*types.SchemaUnion", Required: false},
}

var fields_create_permission_group = []leanruntime.Field{
	{Name: "ApplicationPermissions", Flag: "application-permissions", Type: "[]types.ApplicationPermission", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_create_user = []leanruntime.Field{
	{Name: "ApiAccess", Flag: "api-access", Type: "types.ApiAccess", Required: false},
	{Name: "ApiAccessPrincipalArn", Flag: "api-access-principal-arn", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "EmailAddress", Flag: "email-address", Type: "*string", Required: true},
	{Name: "FirstName", Flag: "first-name", Type: "*string", Required: false},
	{Name: "LastName", Flag: "last-name", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.UserType", Required: true},
}

var fields_delete_dataset = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DatasetId", Flag: "dataset-id", Type: "*string", Required: true},
}

var fields_delete_permission_group = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "PermissionGroupId", Flag: "permission-group-id", Type: "*string", Required: true},
}

var fields_disable_user = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_disassociate_user_from_permission_group = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "PermissionGroupId", Flag: "permission-group-id", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_enable_user = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_get_changeset = []leanruntime.Field{
	{Name: "ChangesetId", Flag: "changeset-id", Type: "*string", Required: true},
	{Name: "DatasetId", Flag: "dataset-id", Type: "*string", Required: true},
}

var fields_get_data_view = []leanruntime.Field{
	{Name: "DataViewId", Flag: "data-view-id", Type: "*string", Required: true},
	{Name: "DatasetId", Flag: "dataset-id", Type: "*string", Required: true},
}

var fields_get_dataset = []leanruntime.Field{
	{Name: "DatasetId", Flag: "dataset-id", Type: "*string", Required: true},
}

var fields_get_external_data_view_access_details = []leanruntime.Field{
	{Name: "DataViewId", Flag: "data-view-id", Type: "*string", Required: true},
	{Name: "DatasetId", Flag: "dataset-id", Type: "*string", Required: true},
}

var fields_get_permission_group = []leanruntime.Field{
	{Name: "PermissionGroupId", Flag: "permission-group-id", Type: "*string", Required: true},
}

var fields_get_programmatic_access_credentials = []leanruntime.Field{
	{Name: "DurationInMinutes", Flag: "duration-in-minutes", Type: "*int64", Required: false},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
}

var fields_get_user = []leanruntime.Field{
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_get_working_location = []leanruntime.Field{
	{Name: "LocationType", Flag: "location-type", Type: "types.LocationType", Required: false},
}

var fields_list_changesets = []leanruntime.Field{
	{Name: "DatasetId", Flag: "dataset-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_data_views = []leanruntime.Field{
	{Name: "DatasetId", Flag: "dataset-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_datasets = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_permission_groups = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_permission_groups_by_user = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_list_users = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_users_by_permission_group = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PermissionGroupId", Flag: "permission-group-id", Type: "*string", Required: true},
}

var fields_reset_user_password = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_update_changeset = []leanruntime.Field{
	{Name: "ChangesetId", Flag: "changeset-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DatasetId", Flag: "dataset-id", Type: "*string", Required: true},
	{Name: "FormatParams", Flag: "format-params", Type: "map[string]string", Required: true},
	{Name: "SourceParams", Flag: "source-params", Type: "map[string]string", Required: true},
}

var fields_update_dataset = []leanruntime.Field{
	{Name: "Alias", Flag: "alias", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DatasetDescription", Flag: "dataset-description", Type: "*string", Required: false},
	{Name: "DatasetId", Flag: "dataset-id", Type: "*string", Required: true},
	{Name: "DatasetTitle", Flag: "dataset-title", Type: "*string", Required: true},
	{Name: "Kind", Flag: "kind", Type: "types.DatasetKind", Required: true},
	{Name: "SchemaDefinition", Flag: "schema-definition", Type: "*types.SchemaUnion", Required: false},
}

var fields_update_permission_group = []leanruntime.Field{
	{Name: "ApplicationPermissions", Flag: "application-permissions", Type: "[]types.ApplicationPermission", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "PermissionGroupId", Flag: "permission-group-id", Type: "*string", Required: true},
}

var fields_update_user = []leanruntime.Field{
	{Name: "ApiAccess", Flag: "api-access", Type: "types.ApiAccess", Required: false},
	{Name: "ApiAccessPrincipalArn", Flag: "api-access-principal-arn", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "FirstName", Flag: "first-name", Type: "*string", Required: false},
	{Name: "LastName", Flag: "last-name", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.UserType", Required: false},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-user-to-permission-group": {
			Name:   "associate-user-to-permission-group",
			Fields: fields_associate_user_to_permission_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateUserToPermissionGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_user_to_permission_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateUserToPermissionGroup(ctx, input)
			},
		},
		"create-changeset": {
			Name:   "create-changeset",
			Fields: fields_create_changeset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateChangesetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_changeset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateChangeset(ctx, input)
			},
		},
		"create-data-view": {
			Name:   "create-data-view",
			Fields: fields_create_data_view,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDataViewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_data_view, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDataView(ctx, input)
			},
		},
		"create-dataset": {
			Name:   "create-dataset",
			Fields: fields_create_dataset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDatasetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_dataset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDataset(ctx, input)
			},
		},
		"create-permission-group": {
			Name:   "create-permission-group",
			Fields: fields_create_permission_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePermissionGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_permission_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePermissionGroup(ctx, input)
			},
		},
		"create-user": {
			Name:   "create-user",
			Fields: fields_create_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateUser(ctx, input)
			},
		},
		"delete-dataset": {
			Name:   "delete-dataset",
			Fields: fields_delete_dataset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDatasetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_dataset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDataset(ctx, input)
			},
		},
		"delete-permission-group": {
			Name:   "delete-permission-group",
			Fields: fields_delete_permission_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePermissionGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_permission_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePermissionGroup(ctx, input)
			},
		},
		"disable-user": {
			Name:   "disable-user",
			Fields: fields_disable_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableUser(ctx, input)
			},
		},
		"disassociate-user-from-permission-group": {
			Name:   "disassociate-user-from-permission-group",
			Fields: fields_disassociate_user_from_permission_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateUserFromPermissionGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_user_from_permission_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateUserFromPermissionGroup(ctx, input)
			},
		},
		"enable-user": {
			Name:   "enable-user",
			Fields: fields_enable_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableUser(ctx, input)
			},
		},
		"get-changeset": {
			Name:   "get-changeset",
			Fields: fields_get_changeset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetChangesetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_changeset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetChangeset(ctx, input)
			},
		},
		"get-data-view": {
			Name:   "get-data-view",
			Fields: fields_get_data_view,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDataViewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_data_view, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDataView(ctx, input)
			},
		},
		"get-dataset": {
			Name:   "get-dataset",
			Fields: fields_get_dataset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDatasetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_dataset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDataset(ctx, input)
			},
		},
		"get-external-data-view-access-details": {
			Name:   "get-external-data-view-access-details",
			Fields: fields_get_external_data_view_access_details,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetExternalDataViewAccessDetailsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_external_data_view_access_details, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetExternalDataViewAccessDetails(ctx, input)
			},
		},
		"get-permission-group": {
			Name:   "get-permission-group",
			Fields: fields_get_permission_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPermissionGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_permission_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPermissionGroup(ctx, input)
			},
		},
		"get-programmatic-access-credentials": {
			Name:   "get-programmatic-access-credentials",
			Fields: fields_get_programmatic_access_credentials,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetProgrammaticAccessCredentialsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_programmatic_access_credentials, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetProgrammaticAccessCredentials(ctx, input)
			},
		},
		"get-user": {
			Name:   "get-user",
			Fields: fields_get_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetUser(ctx, input)
			},
		},
		"get-working-location": {
			Name:   "get-working-location",
			Fields: fields_get_working_location,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWorkingLocationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_working_location, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetWorkingLocation(ctx, input)
			},
		},
		"list-changesets": {
			Name:   "list-changesets",
			Fields: fields_list_changesets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListChangesetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_changesets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListChangesets(ctx, input)
				}
				var results []*svc.ListChangesetsOutput
				p := svc.NewListChangesetsPaginator(client, input)
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
		"list-data-views": {
			Name:   "list-data-views",
			Fields: fields_list_data_views,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDataViewsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_data_views, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDataViews(ctx, input)
				}
				var results []*svc.ListDataViewsOutput
				p := svc.NewListDataViewsPaginator(client, input)
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
		"list-datasets": {
			Name:   "list-datasets",
			Fields: fields_list_datasets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDatasetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_datasets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDatasets(ctx, input)
				}
				var results []*svc.ListDatasetsOutput
				p := svc.NewListDatasetsPaginator(client, input)
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
		"list-permission-groups": {
			Name:   "list-permission-groups",
			Fields: fields_list_permission_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPermissionGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_permission_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPermissionGroups(ctx, input)
				}
				var results []*svc.ListPermissionGroupsOutput
				p := svc.NewListPermissionGroupsPaginator(client, input)
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
		"list-permission-groups-by-user": {
			Name:   "list-permission-groups-by-user",
			Fields: fields_list_permission_groups_by_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPermissionGroupsByUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_permission_groups_by_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListPermissionGroupsByUser(ctx, input)
			},
		},
		"list-users": {
			Name:   "list-users",
			Fields: fields_list_users,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListUsersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_users, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListUsers(ctx, input)
				}
				var results []*svc.ListUsersOutput
				p := svc.NewListUsersPaginator(client, input)
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
		"list-users-by-permission-group": {
			Name:   "list-users-by-permission-group",
			Fields: fields_list_users_by_permission_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListUsersByPermissionGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_users_by_permission_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListUsersByPermissionGroup(ctx, input)
			},
		},
		"reset-user-password": {
			Name:   "reset-user-password",
			Fields: fields_reset_user_password,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResetUserPasswordInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reset_user_password, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResetUserPassword(ctx, input)
			},
		},
		"update-changeset": {
			Name:   "update-changeset",
			Fields: fields_update_changeset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateChangesetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_changeset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateChangeset(ctx, input)
			},
		},
		"update-dataset": {
			Name:   "update-dataset",
			Fields: fields_update_dataset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDatasetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_dataset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDataset(ctx, input)
			},
		},
		"update-permission-group": {
			Name:   "update-permission-group",
			Fields: fields_update_permission_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePermissionGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_permission_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePermissionGroup(ctx, input)
			},
		},
		"update-user": {
			Name:   "update-user",
			Fields: fields_update_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateUser(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("finspacedata", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
