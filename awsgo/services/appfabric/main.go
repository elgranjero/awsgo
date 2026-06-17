package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/appfabric/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"batch-get-user-access-tasks", "connect-app-authorization", "create-app-authorization", "create-app-bundle", "create-ingestion", "create-ingestion-destination", "delete-app-authorization", "delete-app-bundle", "delete-ingestion", "delete-ingestion-destination", "get-app-authorization", "get-app-bundle", "get-ingestion", "get-ingestion-destination", "list-app-authorizations", "list-app-bundles", "list-ingestion-destinations", "list-ingestions", "list-tags-for-resource", "start-ingestion", "start-user-access-tasks", "stop-ingestion", "tag-resource", "untag-resource", "update-app-authorization", "update-ingestion-destination"},
		OperationSet: map[string]bool{"batch-get-user-access-tasks": true, "connect-app-authorization": true, "create-app-authorization": true, "create-app-bundle": true, "create-ingestion": true, "create-ingestion-destination": true, "delete-app-authorization": true, "delete-app-bundle": true, "delete-ingestion": true, "delete-ingestion-destination": true, "get-app-authorization": true, "get-app-bundle": true, "get-ingestion": true, "get-ingestion-destination": true, "list-app-authorizations": true, "list-app-bundles": true, "list-ingestion-destinations": true, "list-ingestions": true, "list-tags-for-resource": true, "start-ingestion": true, "start-user-access-tasks": true, "stop-ingestion": true, "tag-resource": true, "untag-resource": true, "update-app-authorization": true, "update-ingestion-destination": true},
		OperationInputs: map[string][]string{
			"batch-get-user-access-tasks":  {"AppBundleIdentifier", "TaskIdList"},
			"connect-app-authorization":    {"AppAuthorizationIdentifier", "AppBundleIdentifier", "AuthRequest"},
			"create-app-authorization":     {"App", "AppBundleIdentifier", "AuthType", "ClientToken", "Credential", "Tags", "Tenant"},
			"create-app-bundle":            {"ClientToken", "CustomerManagedKeyIdentifier", "Tags"},
			"create-ingestion":             {"App", "AppBundleIdentifier", "ClientToken", "IngestionType", "Tags", "TenantId"},
			"create-ingestion-destination": {"AppBundleIdentifier", "ClientToken", "DestinationConfiguration", "IngestionIdentifier", "ProcessingConfiguration", "Tags"},
			"delete-app-authorization":     {"AppAuthorizationIdentifier", "AppBundleIdentifier"},
			"delete-app-bundle":            {"AppBundleIdentifier"},
			"delete-ingestion":             {"AppBundleIdentifier", "IngestionIdentifier"},
			"delete-ingestion-destination": {"AppBundleIdentifier", "IngestionDestinationIdentifier", "IngestionIdentifier"},
			"get-app-authorization":        {"AppAuthorizationIdentifier", "AppBundleIdentifier"},
			"get-app-bundle":               {"AppBundleIdentifier"},
			"get-ingestion":                {"AppBundleIdentifier", "IngestionIdentifier"},
			"get-ingestion-destination":    {"AppBundleIdentifier", "IngestionDestinationIdentifier", "IngestionIdentifier"},
			"list-app-authorizations":      {"AppBundleIdentifier", "MaxResults", "NextToken"},
			"list-app-bundles":             {"MaxResults", "NextToken"},
			"list-ingestion-destinations":  {"AppBundleIdentifier", "IngestionIdentifier", "MaxResults", "NextToken"},
			"list-ingestions":              {"AppBundleIdentifier", "MaxResults", "NextToken"},
			"list-tags-for-resource":       {"ResourceArn"},
			"start-ingestion":              {"AppBundleIdentifier", "IngestionIdentifier"},
			"start-user-access-tasks":      {"AppBundleIdentifier", "Email"},
			"stop-ingestion":               {"AppBundleIdentifier", "IngestionIdentifier"},
			"tag-resource":                 {"ResourceArn", "Tags"},
			"untag-resource":               {"ResourceArn", "TagKeys"},
			"update-app-authorization":     {"AppAuthorizationIdentifier", "AppBundleIdentifier", "Credential", "Tenant"},
			"update-ingestion-destination": {"AppBundleIdentifier", "DestinationConfiguration", "IngestionDestinationIdentifier", "IngestionIdentifier"},
		},
		OperationInputTypes: map[string]map[string]string{
			"batch-get-user-access-tasks":  {"AppBundleIdentifier": "*string", "TaskIdList": "[]string"},
			"connect-app-authorization":    {"AppAuthorizationIdentifier": "*string", "AppBundleIdentifier": "*string", "AuthRequest": "*types.AuthRequest"},
			"create-app-authorization":     {"App": "*string", "AppBundleIdentifier": "*string", "AuthType": "types.AuthType", "ClientToken": "*string", "Credential": "types.Credential", "Tags": "[]types.Tag", "Tenant": "*types.Tenant"},
			"create-app-bundle":            {"ClientToken": "*string", "CustomerManagedKeyIdentifier": "*string", "Tags": "[]types.Tag"},
			"create-ingestion":             {"App": "*string", "AppBundleIdentifier": "*string", "ClientToken": "*string", "IngestionType": "types.IngestionType", "Tags": "[]types.Tag", "TenantId": "*string"},
			"create-ingestion-destination": {"AppBundleIdentifier": "*string", "ClientToken": "*string", "DestinationConfiguration": "types.DestinationConfiguration", "IngestionIdentifier": "*string", "ProcessingConfiguration": "types.ProcessingConfiguration", "Tags": "[]types.Tag"},
			"delete-app-authorization":     {"AppAuthorizationIdentifier": "*string", "AppBundleIdentifier": "*string"},
			"delete-app-bundle":            {"AppBundleIdentifier": "*string"},
			"delete-ingestion":             {"AppBundleIdentifier": "*string", "IngestionIdentifier": "*string"},
			"delete-ingestion-destination": {"AppBundleIdentifier": "*string", "IngestionDestinationIdentifier": "*string", "IngestionIdentifier": "*string"},
			"get-app-authorization":        {"AppAuthorizationIdentifier": "*string", "AppBundleIdentifier": "*string"},
			"get-app-bundle":               {"AppBundleIdentifier": "*string"},
			"get-ingestion":                {"AppBundleIdentifier": "*string", "IngestionIdentifier": "*string"},
			"get-ingestion-destination":    {"AppBundleIdentifier": "*string", "IngestionDestinationIdentifier": "*string", "IngestionIdentifier": "*string"},
			"list-app-authorizations":      {"AppBundleIdentifier": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-app-bundles":             {"MaxResults": "*int32", "NextToken": "*string"},
			"list-ingestion-destinations":  {"AppBundleIdentifier": "*string", "IngestionIdentifier": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-ingestions":              {"AppBundleIdentifier": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":       {"ResourceArn": "*string"},
			"start-ingestion":              {"AppBundleIdentifier": "*string", "IngestionIdentifier": "*string"},
			"start-user-access-tasks":      {"AppBundleIdentifier": "*string", "Email": "*string"},
			"stop-ingestion":               {"AppBundleIdentifier": "*string", "IngestionIdentifier": "*string"},
			"tag-resource":                 {"ResourceArn": "*string", "Tags": "[]types.Tag"},
			"untag-resource":               {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-app-authorization":     {"AppAuthorizationIdentifier": "*string", "AppBundleIdentifier": "*string", "Credential": "types.Credential", "Tenant": "*types.Tenant"},
			"update-ingestion-destination": {"AppBundleIdentifier": "*string", "DestinationConfiguration": "types.DestinationConfiguration", "IngestionDestinationIdentifier": "*string", "IngestionIdentifier": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"batch-get-user-access-tasks":  {"AppBundleIdentifier", "TaskIdList"},
			"connect-app-authorization":    {"AppAuthorizationIdentifier", "AppBundleIdentifier"},
			"create-app-authorization":     {"App", "AppBundleIdentifier", "AuthType", "Credential", "Tenant"},
			"create-app-bundle":            {},
			"create-ingestion":             {"App", "AppBundleIdentifier", "IngestionType", "TenantId"},
			"create-ingestion-destination": {"AppBundleIdentifier", "DestinationConfiguration", "IngestionIdentifier", "ProcessingConfiguration"},
			"delete-app-authorization":     {"AppAuthorizationIdentifier", "AppBundleIdentifier"},
			"delete-app-bundle":            {"AppBundleIdentifier"},
			"delete-ingestion":             {"AppBundleIdentifier", "IngestionIdentifier"},
			"delete-ingestion-destination": {"AppBundleIdentifier", "IngestionDestinationIdentifier", "IngestionIdentifier"},
			"get-app-authorization":        {"AppAuthorizationIdentifier", "AppBundleIdentifier"},
			"get-app-bundle":               {"AppBundleIdentifier"},
			"get-ingestion":                {"AppBundleIdentifier", "IngestionIdentifier"},
			"get-ingestion-destination":    {"AppBundleIdentifier", "IngestionDestinationIdentifier", "IngestionIdentifier"},
			"list-app-authorizations":      {"AppBundleIdentifier"},
			"list-app-bundles":             {},
			"list-ingestion-destinations":  {"AppBundleIdentifier", "IngestionIdentifier"},
			"list-ingestions":              {"AppBundleIdentifier"},
			"list-tags-for-resource":       {"ResourceArn"},
			"start-ingestion":              {"AppBundleIdentifier", "IngestionIdentifier"},
			"start-user-access-tasks":      {"AppBundleIdentifier", "Email"},
			"stop-ingestion":               {"AppBundleIdentifier", "IngestionIdentifier"},
			"tag-resource":                 {"ResourceArn", "Tags"},
			"untag-resource":               {"ResourceArn", "TagKeys"},
			"update-app-authorization":     {"AppAuthorizationIdentifier", "AppBundleIdentifier"},
			"update-ingestion-destination": {"AppBundleIdentifier", "DestinationConfiguration", "IngestionDestinationIdentifier", "IngestionIdentifier"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("appfabric", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
