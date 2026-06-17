package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/licensemanagerlinuxsubscriptions/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"deregister-subscription-provider", "get-registered-subscription-provider", "get-service-settings", "list-linux-subscription-instances", "list-linux-subscriptions", "list-registered-subscription-providers", "list-tags-for-resource", "register-subscription-provider", "tag-resource", "untag-resource", "update-service-settings"},
		OperationSet: map[string]bool{"deregister-subscription-provider": true, "get-registered-subscription-provider": true, "get-service-settings": true, "list-linux-subscription-instances": true, "list-linux-subscriptions": true, "list-registered-subscription-providers": true, "list-tags-for-resource": true, "register-subscription-provider": true, "tag-resource": true, "untag-resource": true, "update-service-settings": true},
		OperationInputs: map[string][]string{
			"deregister-subscription-provider":       {"SubscriptionProviderArn"},
			"get-registered-subscription-provider":   {"SubscriptionProviderArn"},
			"get-service-settings":                   {},
			"list-linux-subscription-instances":      {"Filters", "MaxResults", "NextToken"},
			"list-linux-subscriptions":               {"Filters", "MaxResults", "NextToken"},
			"list-registered-subscription-providers": {"MaxResults", "NextToken", "SubscriptionProviderSources"},
			"list-tags-for-resource":                 {"ResourceArn"},
			"register-subscription-provider":         {"SecretArn", "SubscriptionProviderSource", "Tags"},
			"tag-resource":                           {"ResourceArn", "Tags"},
			"untag-resource":                         {"ResourceArn", "TagKeys"},
			"update-service-settings":                {"AllowUpdate", "LinuxSubscriptionsDiscovery", "LinuxSubscriptionsDiscoverySettings"},
		},
		OperationInputTypes: map[string]map[string]string{
			"deregister-subscription-provider":       {"SubscriptionProviderArn": "*string"},
			"get-registered-subscription-provider":   {"SubscriptionProviderArn": "*string"},
			"get-service-settings":                   {},
			"list-linux-subscription-instances":      {"Filters": "[]types.Filter", "MaxResults": "*int32", "NextToken": "*string"},
			"list-linux-subscriptions":               {"Filters": "[]types.Filter", "MaxResults": "*int32", "NextToken": "*string"},
			"list-registered-subscription-providers": {"MaxResults": "*int32", "NextToken": "*string", "SubscriptionProviderSources": "[]types.SubscriptionProviderSource"},
			"list-tags-for-resource":                 {"ResourceArn": "*string"},
			"register-subscription-provider":         {"SecretArn": "*string", "SubscriptionProviderSource": "types.SubscriptionProviderSource", "Tags": "map[string]string"},
			"tag-resource":                           {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":                         {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-service-settings":                {"AllowUpdate": "*bool", "LinuxSubscriptionsDiscovery": "types.LinuxSubscriptionsDiscovery", "LinuxSubscriptionsDiscoverySettings": "*types.LinuxSubscriptionsDiscoverySettings"},
		},
		OperationInputRequired: map[string][]string{
			"deregister-subscription-provider":       {"SubscriptionProviderArn"},
			"get-registered-subscription-provider":   {"SubscriptionProviderArn"},
			"get-service-settings":                   {},
			"list-linux-subscription-instances":      {},
			"list-linux-subscriptions":               {},
			"list-registered-subscription-providers": {},
			"list-tags-for-resource":                 {"ResourceArn"},
			"register-subscription-provider":         {"SecretArn", "SubscriptionProviderSource"},
			"tag-resource":                           {"ResourceArn", "Tags"},
			"untag-resource":                         {"ResourceArn", "TagKeys"},
			"update-service-settings":                {"LinuxSubscriptionsDiscovery", "LinuxSubscriptionsDiscoverySettings"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("licensemanagerlinuxsubscriptions", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
