package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/workdocs"
)

var fields_abort_document_version_upload = []leanruntime.Field{
	{Name: "AuthenticationToken", Flag: "authentication-token", Type: "*string", Required: false},
	{Name: "DocumentId", Flag: "document-id", Type: "*string", Required: true},
	{Name: "VersionId", Flag: "version-id", Type: "*string", Required: true},
}

var fields_activate_user = []leanruntime.Field{
	{Name: "AuthenticationToken", Flag: "authentication-token", Type: "*string", Required: false},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_add_resource_permissions = []leanruntime.Field{
	{Name: "AuthenticationToken", Flag: "authentication-token", Type: "*string", Required: false},
	{Name: "NotificationOptions", Flag: "notification-options", Type: "*types.NotificationOptions", Required: false},
	{Name: "Principals", Flag: "principals", Type: "[]types.SharePrincipal", Required: true},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
}

var fields_create_comment = []leanruntime.Field{
	{Name: "AuthenticationToken", Flag: "authentication-token", Type: "*string", Required: false},
	{Name: "DocumentId", Flag: "document-id", Type: "*string", Required: true},
	{Name: "NotifyCollaborators", Flag: "notify-collaborators", Type: "bool", Required: false},
	{Name: "ParentId", Flag: "parent-id", Type: "*string", Required: false},
	{Name: "Text", Flag: "text", Type: "*string", Required: true},
	{Name: "ThreadId", Flag: "thread-id", Type: "*string", Required: false},
	{Name: "VersionId", Flag: "version-id", Type: "*string", Required: true},
	{Name: "Visibility", Flag: "visibility", Type: "types.CommentVisibilityType", Required: false},
}

var fields_create_custom_metadata = []leanruntime.Field{
	{Name: "AuthenticationToken", Flag: "authentication-token", Type: "*string", Required: false},
	{Name: "CustomMetadata", Flag: "custom-metadata", Type: "map[string]string", Required: true},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "VersionId", Flag: "version-id", Type: "*string", Required: false},
}

var fields_create_folder = []leanruntime.Field{
	{Name: "AuthenticationToken", Flag: "authentication-token", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "ParentFolderId", Flag: "parent-folder-id", Type: "*string", Required: true},
}

var fields_create_labels = []leanruntime.Field{
	{Name: "AuthenticationToken", Flag: "authentication-token", Type: "*string", Required: false},
	{Name: "Labels", Flag: "labels", Type: "[]string", Required: true},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
}

var fields_create_notification_subscription = []leanruntime.Field{
	{Name: "Endpoint", Flag: "endpoint", Type: "*string", Required: true},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
	{Name: "Protocol", Flag: "protocol", Type: "types.SubscriptionProtocolType", Required: true},
	{Name: "SubscriptionType", Flag: "subscription-type", Type: "types.SubscriptionType", Required: true},
}

var fields_create_user = []leanruntime.Field{
	{Name: "AuthenticationToken", Flag: "authentication-token", Type: "*string", Required: false},
	{Name: "EmailAddress", Flag: "email-address", Type: "*string", Required: false},
	{Name: "GivenName", Flag: "given-name", Type: "*string", Required: true},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: false},
	{Name: "Password", Flag: "password", Type: "*string", Required: true},
	{Name: "StorageRule", Flag: "storage-rule", Type: "*types.StorageRuleType", Required: false},
	{Name: "Surname", Flag: "surname", Type: "*string", Required: true},
	{Name: "TimeZoneId", Flag: "time-zone-id", Type: "*string", Required: false},
	{Name: "Username", Flag: "username", Type: "*string", Required: true},
}

var fields_deactivate_user = []leanruntime.Field{
	{Name: "AuthenticationToken", Flag: "authentication-token", Type: "*string", Required: false},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_delete_comment = []leanruntime.Field{
	{Name: "AuthenticationToken", Flag: "authentication-token", Type: "*string", Required: false},
	{Name: "CommentId", Flag: "comment-id", Type: "*string", Required: true},
	{Name: "DocumentId", Flag: "document-id", Type: "*string", Required: true},
	{Name: "VersionId", Flag: "version-id", Type: "*string", Required: true},
}

var fields_delete_custom_metadata = []leanruntime.Field{
	{Name: "AuthenticationToken", Flag: "authentication-token", Type: "*string", Required: false},
	{Name: "DeleteAll", Flag: "delete-all", Type: "bool", Required: false},
	{Name: "Keys", Flag: "keys", Type: "[]string", Required: false},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "VersionId", Flag: "version-id", Type: "*string", Required: false},
}

var fields_delete_document = []leanruntime.Field{
	{Name: "AuthenticationToken", Flag: "authentication-token", Type: "*string", Required: false},
	{Name: "DocumentId", Flag: "document-id", Type: "*string", Required: true},
}

var fields_delete_document_version = []leanruntime.Field{
	{Name: "AuthenticationToken", Flag: "authentication-token", Type: "*string", Required: false},
	{Name: "DeletePriorVersions", Flag: "delete-prior-versions", Type: "bool", Required: true},
	{Name: "DocumentId", Flag: "document-id", Type: "*string", Required: true},
	{Name: "VersionId", Flag: "version-id", Type: "*string", Required: true},
}

var fields_delete_folder = []leanruntime.Field{
	{Name: "AuthenticationToken", Flag: "authentication-token", Type: "*string", Required: false},
	{Name: "FolderId", Flag: "folder-id", Type: "*string", Required: true},
}

var fields_delete_folder_contents = []leanruntime.Field{
	{Name: "AuthenticationToken", Flag: "authentication-token", Type: "*string", Required: false},
	{Name: "FolderId", Flag: "folder-id", Type: "*string", Required: true},
}

var fields_delete_labels = []leanruntime.Field{
	{Name: "AuthenticationToken", Flag: "authentication-token", Type: "*string", Required: false},
	{Name: "DeleteAll", Flag: "delete-all", Type: "bool", Required: false},
	{Name: "Labels", Flag: "labels", Type: "[]string", Required: false},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
}

var fields_delete_notification_subscription = []leanruntime.Field{
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
	{Name: "SubscriptionId", Flag: "subscription-id", Type: "*string", Required: true},
}

var fields_delete_user = []leanruntime.Field{
	{Name: "AuthenticationToken", Flag: "authentication-token", Type: "*string", Required: false},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_describe_activities = []leanruntime.Field{
	{Name: "ActivityTypes", Flag: "activity-types", Type: "*string", Required: false},
	{Name: "AuthenticationToken", Flag: "authentication-token", Type: "*string", Required: false},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "IncludeIndirectActivities", Flag: "include-indirect-activities", Type: "bool", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: false},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: false},
}

var fields_describe_comments = []leanruntime.Field{
	{Name: "AuthenticationToken", Flag: "authentication-token", Type: "*string", Required: false},
	{Name: "DocumentId", Flag: "document-id", Type: "*string", Required: true},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "VersionId", Flag: "version-id", Type: "*string", Required: true},
}

var fields_describe_document_versions = []leanruntime.Field{
	{Name: "AuthenticationToken", Flag: "authentication-token", Type: "*string", Required: false},
	{Name: "DocumentId", Flag: "document-id", Type: "*string", Required: true},
	{Name: "Fields", Flag: "fields", Type: "*string", Required: false},
	{Name: "Include", Flag: "include", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
}

var fields_describe_folder_contents = []leanruntime.Field{
	{Name: "AuthenticationToken", Flag: "authentication-token", Type: "*string", Required: false},
	{Name: "FolderId", Flag: "folder-id", Type: "*string", Required: true},
	{Name: "Include", Flag: "include", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "Order", Flag: "order", Type: "types.OrderType", Required: false},
	{Name: "Sort", Flag: "sort", Type: "types.ResourceSortType", Required: false},
	{Name: "Type", Flag: "type", Type: "types.FolderContentType", Required: false},
}

var fields_describe_groups = []leanruntime.Field{
	{Name: "AuthenticationToken", Flag: "authentication-token", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: false},
	{Name: "SearchQuery", Flag: "search-query", Type: "*string", Required: true},
}

var fields_describe_notification_subscriptions = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_describe_resource_permissions = []leanruntime.Field{
	{Name: "AuthenticationToken", Flag: "authentication-token", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "PrincipalId", Flag: "principal-id", Type: "*string", Required: false},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
}

var fields_describe_root_folders = []leanruntime.Field{
	{Name: "AuthenticationToken", Flag: "authentication-token", Type: "*string", Required: true},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
}

var fields_describe_users = []leanruntime.Field{
	{Name: "AuthenticationToken", Flag: "authentication-token", Type: "*string", Required: false},
	{Name: "Fields", Flag: "fields", Type: "*string", Required: false},
	{Name: "Include", Flag: "include", Type: "types.UserFilterType", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "Order", Flag: "order", Type: "types.OrderType", Required: false},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: false},
	{Name: "Query", Flag: "query", Type: "*string", Required: false},
	{Name: "Sort", Flag: "sort", Type: "types.UserSortType", Required: false},
	{Name: "UserIds", Flag: "user-ids", Type: "*string", Required: false},
}

var fields_get_current_user = []leanruntime.Field{
	{Name: "AuthenticationToken", Flag: "authentication-token", Type: "*string", Required: true},
}

var fields_get_document = []leanruntime.Field{
	{Name: "AuthenticationToken", Flag: "authentication-token", Type: "*string", Required: false},
	{Name: "DocumentId", Flag: "document-id", Type: "*string", Required: true},
	{Name: "IncludeCustomMetadata", Flag: "include-custom-metadata", Type: "bool", Required: false},
}

var fields_get_document_path = []leanruntime.Field{
	{Name: "AuthenticationToken", Flag: "authentication-token", Type: "*string", Required: false},
	{Name: "DocumentId", Flag: "document-id", Type: "*string", Required: true},
	{Name: "Fields", Flag: "fields", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
}

var fields_get_document_version = []leanruntime.Field{
	{Name: "AuthenticationToken", Flag: "authentication-token", Type: "*string", Required: false},
	{Name: "DocumentId", Flag: "document-id", Type: "*string", Required: true},
	{Name: "Fields", Flag: "fields", Type: "*string", Required: false},
	{Name: "IncludeCustomMetadata", Flag: "include-custom-metadata", Type: "bool", Required: false},
	{Name: "VersionId", Flag: "version-id", Type: "*string", Required: true},
}

var fields_get_folder = []leanruntime.Field{
	{Name: "AuthenticationToken", Flag: "authentication-token", Type: "*string", Required: false},
	{Name: "FolderId", Flag: "folder-id", Type: "*string", Required: true},
	{Name: "IncludeCustomMetadata", Flag: "include-custom-metadata", Type: "bool", Required: false},
}

var fields_get_folder_path = []leanruntime.Field{
	{Name: "AuthenticationToken", Flag: "authentication-token", Type: "*string", Required: false},
	{Name: "Fields", Flag: "fields", Type: "*string", Required: false},
	{Name: "FolderId", Flag: "folder-id", Type: "*string", Required: true},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
}

var fields_get_resources = []leanruntime.Field{
	{Name: "AuthenticationToken", Flag: "authentication-token", Type: "*string", Required: false},
	{Name: "CollectionType", Flag: "collection-type", Type: "types.ResourceCollectionType", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: false},
}

var fields_initiate_document_version_upload = []leanruntime.Field{
	{Name: "AuthenticationToken", Flag: "authentication-token", Type: "*string", Required: false},
	{Name: "ContentCreatedTimestamp", Flag: "content-created-timestamp", Type: "*time.Time", Required: false},
	{Name: "ContentModifiedTimestamp", Flag: "content-modified-timestamp", Type: "*time.Time", Required: false},
	{Name: "ContentType", Flag: "content-type", Type: "*string", Required: false},
	{Name: "DocumentSizeInBytes", Flag: "document-size-in-bytes", Type: "*int64", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "ParentFolderId", Flag: "parent-folder-id", Type: "*string", Required: false},
}

var fields_remove_all_resource_permissions = []leanruntime.Field{
	{Name: "AuthenticationToken", Flag: "authentication-token", Type: "*string", Required: false},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
}

var fields_remove_resource_permission = []leanruntime.Field{
	{Name: "AuthenticationToken", Flag: "authentication-token", Type: "*string", Required: false},
	{Name: "PrincipalId", Flag: "principal-id", Type: "*string", Required: true},
	{Name: "PrincipalType", Flag: "principal-type", Type: "types.PrincipalType", Required: false},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
}

var fields_restore_document_versions = []leanruntime.Field{
	{Name: "AuthenticationToken", Flag: "authentication-token", Type: "*string", Required: false},
	{Name: "DocumentId", Flag: "document-id", Type: "*string", Required: true},
}

var fields_search_resources = []leanruntime.Field{
	{Name: "AdditionalResponseFields", Flag: "additional-response-fields", Type: "[]types.AdditionalResponseFieldType", Required: false},
	{Name: "AuthenticationToken", Flag: "authentication-token", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "*types.Filters", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "OrderBy", Flag: "order-by", Type: "[]types.SearchSortResult", Required: false},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: false},
	{Name: "QueryScopes", Flag: "query-scopes", Type: "[]types.SearchQueryScopeType", Required: false},
	{Name: "QueryText", Flag: "query-text", Type: "*string", Required: false},
}

var fields_update_document = []leanruntime.Field{
	{Name: "AuthenticationToken", Flag: "authentication-token", Type: "*string", Required: false},
	{Name: "DocumentId", Flag: "document-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "ParentFolderId", Flag: "parent-folder-id", Type: "*string", Required: false},
	{Name: "ResourceState", Flag: "resource-state", Type: "types.ResourceStateType", Required: false},
}

var fields_update_document_version = []leanruntime.Field{
	{Name: "AuthenticationToken", Flag: "authentication-token", Type: "*string", Required: false},
	{Name: "DocumentId", Flag: "document-id", Type: "*string", Required: true},
	{Name: "VersionId", Flag: "version-id", Type: "*string", Required: true},
	{Name: "VersionStatus", Flag: "version-status", Type: "types.DocumentVersionStatus", Required: false},
}

var fields_update_folder = []leanruntime.Field{
	{Name: "AuthenticationToken", Flag: "authentication-token", Type: "*string", Required: false},
	{Name: "FolderId", Flag: "folder-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "ParentFolderId", Flag: "parent-folder-id", Type: "*string", Required: false},
	{Name: "ResourceState", Flag: "resource-state", Type: "types.ResourceStateType", Required: false},
}

var fields_update_user = []leanruntime.Field{
	{Name: "AuthenticationToken", Flag: "authentication-token", Type: "*string", Required: false},
	{Name: "GivenName", Flag: "given-name", Type: "*string", Required: false},
	{Name: "GrantPoweruserPrivileges", Flag: "grant-poweruser-privileges", Type: "types.BooleanEnumType", Required: false},
	{Name: "Locale", Flag: "locale", Type: "types.LocaleType", Required: false},
	{Name: "StorageRule", Flag: "storage-rule", Type: "*types.StorageRuleType", Required: false},
	{Name: "Surname", Flag: "surname", Type: "*string", Required: false},
	{Name: "TimeZoneId", Flag: "time-zone-id", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.UserType", Required: false},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"abort-document-version-upload": {
			Name:   "abort-document-version-upload",
			Fields: fields_abort_document_version_upload,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AbortDocumentVersionUploadInput{}
				if _, err := leanruntime.ApplyInput(input, fields_abort_document_version_upload, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AbortDocumentVersionUpload(ctx, input)
			},
		},
		"activate-user": {
			Name:   "activate-user",
			Fields: fields_activate_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ActivateUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_activate_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ActivateUser(ctx, input)
			},
		},
		"add-resource-permissions": {
			Name:   "add-resource-permissions",
			Fields: fields_add_resource_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddResourcePermissionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_resource_permissions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddResourcePermissions(ctx, input)
			},
		},
		"create-comment": {
			Name:   "create-comment",
			Fields: fields_create_comment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCommentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_comment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateComment(ctx, input)
			},
		},
		"create-custom-metadata": {
			Name:   "create-custom-metadata",
			Fields: fields_create_custom_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCustomMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_custom_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCustomMetadata(ctx, input)
			},
		},
		"create-folder": {
			Name:   "create-folder",
			Fields: fields_create_folder,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFolderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_folder, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFolder(ctx, input)
			},
		},
		"create-labels": {
			Name:   "create-labels",
			Fields: fields_create_labels,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLabelsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_labels, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLabels(ctx, input)
			},
		},
		"create-notification-subscription": {
			Name:   "create-notification-subscription",
			Fields: fields_create_notification_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateNotificationSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_notification_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateNotificationSubscription(ctx, input)
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
		"deactivate-user": {
			Name:   "deactivate-user",
			Fields: fields_deactivate_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeactivateUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deactivate_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeactivateUser(ctx, input)
			},
		},
		"delete-comment": {
			Name:   "delete-comment",
			Fields: fields_delete_comment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCommentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_comment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteComment(ctx, input)
			},
		},
		"delete-custom-metadata": {
			Name:   "delete-custom-metadata",
			Fields: fields_delete_custom_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCustomMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_custom_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCustomMetadata(ctx, input)
			},
		},
		"delete-document": {
			Name:   "delete-document",
			Fields: fields_delete_document,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDocumentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_document, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDocument(ctx, input)
			},
		},
		"delete-document-version": {
			Name:   "delete-document-version",
			Fields: fields_delete_document_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDocumentVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_document_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDocumentVersion(ctx, input)
			},
		},
		"delete-folder": {
			Name:   "delete-folder",
			Fields: fields_delete_folder,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFolderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_folder, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFolder(ctx, input)
			},
		},
		"delete-folder-contents": {
			Name:   "delete-folder-contents",
			Fields: fields_delete_folder_contents,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFolderContentsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_folder_contents, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFolderContents(ctx, input)
			},
		},
		"delete-labels": {
			Name:   "delete-labels",
			Fields: fields_delete_labels,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLabelsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_labels, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLabels(ctx, input)
			},
		},
		"delete-notification-subscription": {
			Name:   "delete-notification-subscription",
			Fields: fields_delete_notification_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteNotificationSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_notification_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteNotificationSubscription(ctx, input)
			},
		},
		"delete-user": {
			Name:   "delete-user",
			Fields: fields_delete_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteUser(ctx, input)
			},
		},
		"describe-activities": {
			Name:   "describe-activities",
			Fields: fields_describe_activities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeActivitiesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_activities, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeActivities(ctx, input)
				}
				var results []*svc.DescribeActivitiesOutput
				p := svc.NewDescribeActivitiesPaginator(client, input)
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
		"describe-comments": {
			Name:   "describe-comments",
			Fields: fields_describe_comments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCommentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_comments, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeComments(ctx, input)
				}
				var results []*svc.DescribeCommentsOutput
				p := svc.NewDescribeCommentsPaginator(client, input)
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
		"describe-document-versions": {
			Name:   "describe-document-versions",
			Fields: fields_describe_document_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDocumentVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_document_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeDocumentVersions(ctx, input)
				}
				var results []*svc.DescribeDocumentVersionsOutput
				p := svc.NewDescribeDocumentVersionsPaginator(client, input)
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
		"describe-folder-contents": {
			Name:   "describe-folder-contents",
			Fields: fields_describe_folder_contents,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFolderContentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_folder_contents, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeFolderContents(ctx, input)
				}
				var results []*svc.DescribeFolderContentsOutput
				p := svc.NewDescribeFolderContentsPaginator(client, input)
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
		"describe-groups": {
			Name:   "describe-groups",
			Fields: fields_describe_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeGroups(ctx, input)
				}
				var results []*svc.DescribeGroupsOutput
				p := svc.NewDescribeGroupsPaginator(client, input)
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
		"describe-notification-subscriptions": {
			Name:   "describe-notification-subscriptions",
			Fields: fields_describe_notification_subscriptions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeNotificationSubscriptionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_notification_subscriptions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeNotificationSubscriptions(ctx, input)
				}
				var results []*svc.DescribeNotificationSubscriptionsOutput
				p := svc.NewDescribeNotificationSubscriptionsPaginator(client, input)
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
		"describe-resource-permissions": {
			Name:   "describe-resource-permissions",
			Fields: fields_describe_resource_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeResourcePermissionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_resource_permissions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeResourcePermissions(ctx, input)
				}
				var results []*svc.DescribeResourcePermissionsOutput
				p := svc.NewDescribeResourcePermissionsPaginator(client, input)
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
		"describe-root-folders": {
			Name:   "describe-root-folders",
			Fields: fields_describe_root_folders,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRootFoldersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_root_folders, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeRootFolders(ctx, input)
				}
				var results []*svc.DescribeRootFoldersOutput
				p := svc.NewDescribeRootFoldersPaginator(client, input)
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
		"describe-users": {
			Name:   "describe-users",
			Fields: fields_describe_users,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeUsersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_users, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeUsers(ctx, input)
				}
				var results []*svc.DescribeUsersOutput
				p := svc.NewDescribeUsersPaginator(client, input)
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
		"get-current-user": {
			Name:   "get-current-user",
			Fields: fields_get_current_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCurrentUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_current_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCurrentUser(ctx, input)
			},
		},
		"get-document": {
			Name:   "get-document",
			Fields: fields_get_document,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDocumentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_document, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDocument(ctx, input)
			},
		},
		"get-document-path": {
			Name:   "get-document-path",
			Fields: fields_get_document_path,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDocumentPathInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_document_path, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDocumentPath(ctx, input)
			},
		},
		"get-document-version": {
			Name:   "get-document-version",
			Fields: fields_get_document_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDocumentVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_document_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDocumentVersion(ctx, input)
			},
		},
		"get-folder": {
			Name:   "get-folder",
			Fields: fields_get_folder,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFolderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_folder, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFolder(ctx, input)
			},
		},
		"get-folder-path": {
			Name:   "get-folder-path",
			Fields: fields_get_folder_path,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFolderPathInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_folder_path, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFolderPath(ctx, input)
			},
		},
		"get-resources": {
			Name:   "get-resources",
			Fields: fields_get_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourcesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_resources, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResources(ctx, input)
			},
		},
		"initiate-document-version-upload": {
			Name:   "initiate-document-version-upload",
			Fields: fields_initiate_document_version_upload,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.InitiateDocumentVersionUploadInput{}
				if _, err := leanruntime.ApplyInput(input, fields_initiate_document_version_upload, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.InitiateDocumentVersionUpload(ctx, input)
			},
		},
		"remove-all-resource-permissions": {
			Name:   "remove-all-resource-permissions",
			Fields: fields_remove_all_resource_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveAllResourcePermissionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_all_resource_permissions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveAllResourcePermissions(ctx, input)
			},
		},
		"remove-resource-permission": {
			Name:   "remove-resource-permission",
			Fields: fields_remove_resource_permission,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveResourcePermissionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_resource_permission, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveResourcePermission(ctx, input)
			},
		},
		"restore-document-versions": {
			Name:   "restore-document-versions",
			Fields: fields_restore_document_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RestoreDocumentVersionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_restore_document_versions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RestoreDocumentVersions(ctx, input)
			},
		},
		"search-resources": {
			Name:   "search-resources",
			Fields: fields_search_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchResourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_resources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchResources(ctx, input)
				}
				var results []*svc.SearchResourcesOutput
				p := svc.NewSearchResourcesPaginator(client, input)
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
		"update-document": {
			Name:   "update-document",
			Fields: fields_update_document,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDocumentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_document, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDocument(ctx, input)
			},
		},
		"update-document-version": {
			Name:   "update-document-version",
			Fields: fields_update_document_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDocumentVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_document_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDocumentVersion(ctx, input)
			},
		},
		"update-folder": {
			Name:   "update-folder",
			Fields: fields_update_folder,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFolderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_folder, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFolder(ctx, input)
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
	if err := leanruntime.Execute("workdocs", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
