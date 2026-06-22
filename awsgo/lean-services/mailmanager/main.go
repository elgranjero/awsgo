package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/mailmanager"
)

var fields_create_addon_instance = []leanruntime.Field{
	{Name: "AddonSubscriptionId", Flag: "addon-subscription-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_addon_subscription = []leanruntime.Field{
	{Name: "AddonName", Flag: "addon-name", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_address_list = []leanruntime.Field{
	{Name: "AddressListName", Flag: "address-list-name", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_address_list_import_job = []leanruntime.Field{
	{Name: "AddressListId", Flag: "address-list-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ImportDataFormat", Flag: "import-data-format", Type: "*types.ImportDataFormat", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_create_archive = []leanruntime.Field{
	{Name: "ArchiveName", Flag: "archive-name", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "KmsKeyArn", Flag: "kms-key-arn", Type: "*string", Required: false},
	{Name: "Retention", Flag: "retention", Type: "types.ArchiveRetention", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_ingress_point = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "IngressPointConfiguration", Flag: "ingress-point-configuration", Type: "types.IngressPointConfiguration", Required: false},
	{Name: "IngressPointName", Flag: "ingress-point-name", Type: "*string", Required: true},
	{Name: "NetworkConfiguration", Flag: "network-configuration", Type: "types.NetworkConfiguration", Required: false},
	{Name: "RuleSetId", Flag: "rule-set-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TrafficPolicyId", Flag: "traffic-policy-id", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.IngressPointType", Required: true},
}

var fields_create_relay = []leanruntime.Field{
	{Name: "Authentication", Flag: "authentication", Type: "types.RelayAuthentication", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "RelayName", Flag: "relay-name", Type: "*string", Required: true},
	{Name: "ServerName", Flag: "server-name", Type: "*string", Required: true},
	{Name: "ServerPort", Flag: "server-port", Type: "*int32", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_rule_set = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "RuleSetName", Flag: "rule-set-name", Type: "*string", Required: true},
	{Name: "Rules", Flag: "rules", Type: "[]types.Rule", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_traffic_policy = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DefaultAction", Flag: "default-action", Type: "types.AcceptAction", Required: true},
	{Name: "MaxMessageSizeBytes", Flag: "max-message-size-bytes", Type: "*int32", Required: false},
	{Name: "PolicyStatements", Flag: "policy-statements", Type: "[]types.PolicyStatement", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TrafficPolicyName", Flag: "traffic-policy-name", Type: "*string", Required: true},
}

var fields_delete_addon_instance = []leanruntime.Field{
	{Name: "AddonInstanceId", Flag: "addon-instance-id", Type: "*string", Required: true},
}

var fields_delete_addon_subscription = []leanruntime.Field{
	{Name: "AddonSubscriptionId", Flag: "addon-subscription-id", Type: "*string", Required: true},
}

var fields_delete_address_list = []leanruntime.Field{
	{Name: "AddressListId", Flag: "address-list-id", Type: "*string", Required: true},
}

var fields_delete_archive = []leanruntime.Field{
	{Name: "ArchiveId", Flag: "archive-id", Type: "*string", Required: true},
}

var fields_delete_ingress_point = []leanruntime.Field{
	{Name: "IngressPointId", Flag: "ingress-point-id", Type: "*string", Required: true},
}

var fields_delete_relay = []leanruntime.Field{
	{Name: "RelayId", Flag: "relay-id", Type: "*string", Required: true},
}

var fields_delete_rule_set = []leanruntime.Field{
	{Name: "RuleSetId", Flag: "rule-set-id", Type: "*string", Required: true},
}

var fields_delete_traffic_policy = []leanruntime.Field{
	{Name: "TrafficPolicyId", Flag: "traffic-policy-id", Type: "*string", Required: true},
}

var fields_deregister_member_from_address_list = []leanruntime.Field{
	{Name: "Address", Flag: "address", Type: "*string", Required: true},
	{Name: "AddressListId", Flag: "address-list-id", Type: "*string", Required: true},
}

var fields_get_addon_instance = []leanruntime.Field{
	{Name: "AddonInstanceId", Flag: "addon-instance-id", Type: "*string", Required: true},
}

var fields_get_addon_subscription = []leanruntime.Field{
	{Name: "AddonSubscriptionId", Flag: "addon-subscription-id", Type: "*string", Required: true},
}

var fields_get_address_list = []leanruntime.Field{
	{Name: "AddressListId", Flag: "address-list-id", Type: "*string", Required: true},
}

var fields_get_address_list_import_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_get_archive = []leanruntime.Field{
	{Name: "ArchiveId", Flag: "archive-id", Type: "*string", Required: true},
}

var fields_get_archive_export = []leanruntime.Field{
	{Name: "ExportId", Flag: "export-id", Type: "*string", Required: true},
}

var fields_get_archive_message = []leanruntime.Field{
	{Name: "ArchivedMessageId", Flag: "archived-message-id", Type: "*string", Required: true},
}

var fields_get_archive_message_content = []leanruntime.Field{
	{Name: "ArchivedMessageId", Flag: "archived-message-id", Type: "*string", Required: true},
}

var fields_get_archive_search = []leanruntime.Field{
	{Name: "SearchId", Flag: "search-id", Type: "*string", Required: true},
}

var fields_get_archive_search_results = []leanruntime.Field{
	{Name: "SearchId", Flag: "search-id", Type: "*string", Required: true},
}

var fields_get_ingress_point = []leanruntime.Field{
	{Name: "IngressPointId", Flag: "ingress-point-id", Type: "*string", Required: true},
}

var fields_get_member_of_address_list = []leanruntime.Field{
	{Name: "Address", Flag: "address", Type: "*string", Required: true},
	{Name: "AddressListId", Flag: "address-list-id", Type: "*string", Required: true},
}

var fields_get_relay = []leanruntime.Field{
	{Name: "RelayId", Flag: "relay-id", Type: "*string", Required: true},
}

var fields_get_rule_set = []leanruntime.Field{
	{Name: "RuleSetId", Flag: "rule-set-id", Type: "*string", Required: true},
}

var fields_get_traffic_policy = []leanruntime.Field{
	{Name: "TrafficPolicyId", Flag: "traffic-policy-id", Type: "*string", Required: true},
}

var fields_list_addon_instances = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
}

var fields_list_addon_subscriptions = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
}

var fields_list_address_list_import_jobs = []leanruntime.Field{
	{Name: "AddressListId", Flag: "address-list-id", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
}

var fields_list_address_lists = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
}

var fields_list_archive_exports = []leanruntime.Field{
	{Name: "ArchiveId", Flag: "archive-id", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
}

var fields_list_archive_searches = []leanruntime.Field{
	{Name: "ArchiveId", Flag: "archive-id", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
}

var fields_list_archives = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
}

var fields_list_ingress_points = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
}

var fields_list_members_of_address_list = []leanruntime.Field{
	{Name: "AddressListId", Flag: "address-list-id", Type: "*string", Required: true},
	{Name: "Filter", Flag: "filter", Type: "*types.AddressFilter", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
}

var fields_list_relays = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
}

var fields_list_rule_sets = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_traffic_policies = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
}

var fields_register_member_to_address_list = []leanruntime.Field{
	{Name: "Address", Flag: "address", Type: "*string", Required: true},
	{Name: "AddressListId", Flag: "address-list-id", Type: "*string", Required: true},
}

var fields_start_address_list_import_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_start_archive_export = []leanruntime.Field{
	{Name: "ArchiveId", Flag: "archive-id", Type: "*string", Required: true},
	{Name: "ExportDestinationConfiguration", Flag: "export-destination-configuration", Type: "types.ExportDestinationConfiguration", Required: true},
	{Name: "Filters", Flag: "filters", Type: "*types.ArchiveFilters", Required: false},
	{Name: "FromTimestamp", Flag: "from-timestamp", Type: "*time.Time", Required: true},
	{Name: "IncludeMetadata", Flag: "include-metadata", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "ToTimestamp", Flag: "to-timestamp", Type: "*time.Time", Required: true},
}

var fields_start_archive_search = []leanruntime.Field{
	{Name: "ArchiveId", Flag: "archive-id", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "*types.ArchiveFilters", Required: false},
	{Name: "FromTimestamp", Flag: "from-timestamp", Type: "*time.Time", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: true},
	{Name: "ToTimestamp", Flag: "to-timestamp", Type: "*time.Time", Required: true},
}

var fields_stop_address_list_import_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_stop_archive_export = []leanruntime.Field{
	{Name: "ExportId", Flag: "export-id", Type: "*string", Required: true},
}

var fields_stop_archive_search = []leanruntime.Field{
	{Name: "SearchId", Flag: "search-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_archive = []leanruntime.Field{
	{Name: "ArchiveId", Flag: "archive-id", Type: "*string", Required: true},
	{Name: "ArchiveName", Flag: "archive-name", Type: "*string", Required: false},
	{Name: "Retention", Flag: "retention", Type: "types.ArchiveRetention", Required: false},
}

var fields_update_ingress_point = []leanruntime.Field{
	{Name: "IngressPointConfiguration", Flag: "ingress-point-configuration", Type: "types.IngressPointConfiguration", Required: false},
	{Name: "IngressPointId", Flag: "ingress-point-id", Type: "*string", Required: true},
	{Name: "IngressPointName", Flag: "ingress-point-name", Type: "*string", Required: false},
	{Name: "RuleSetId", Flag: "rule-set-id", Type: "*string", Required: false},
	{Name: "StatusToUpdate", Flag: "status-to-update", Type: "types.IngressPointStatusToUpdate", Required: false},
	{Name: "TrafficPolicyId", Flag: "traffic-policy-id", Type: "*string", Required: false},
}

var fields_update_relay = []leanruntime.Field{
	{Name: "Authentication", Flag: "authentication", Type: "types.RelayAuthentication", Required: false},
	{Name: "RelayId", Flag: "relay-id", Type: "*string", Required: true},
	{Name: "RelayName", Flag: "relay-name", Type: "*string", Required: false},
	{Name: "ServerName", Flag: "server-name", Type: "*string", Required: false},
	{Name: "ServerPort", Flag: "server-port", Type: "*int32", Required: false},
}

var fields_update_rule_set = []leanruntime.Field{
	{Name: "RuleSetId", Flag: "rule-set-id", Type: "*string", Required: true},
	{Name: "RuleSetName", Flag: "rule-set-name", Type: "*string", Required: false},
	{Name: "Rules", Flag: "rules", Type: "[]types.Rule", Required: false},
}

var fields_update_traffic_policy = []leanruntime.Field{
	{Name: "DefaultAction", Flag: "default-action", Type: "types.AcceptAction", Required: false},
	{Name: "MaxMessageSizeBytes", Flag: "max-message-size-bytes", Type: "*int32", Required: false},
	{Name: "PolicyStatements", Flag: "policy-statements", Type: "[]types.PolicyStatement", Required: false},
	{Name: "TrafficPolicyId", Flag: "traffic-policy-id", Type: "*string", Required: true},
	{Name: "TrafficPolicyName", Flag: "traffic-policy-name", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-addon-instance": {
			Name:   "create-addon-instance",
			Fields: fields_create_addon_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAddonInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_addon_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAddonInstance(ctx, input)
			},
		},
		"create-addon-subscription": {
			Name:   "create-addon-subscription",
			Fields: fields_create_addon_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAddonSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_addon_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAddonSubscription(ctx, input)
			},
		},
		"create-address-list": {
			Name:   "create-address-list",
			Fields: fields_create_address_list,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAddressListInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_address_list, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAddressList(ctx, input)
			},
		},
		"create-address-list-import-job": {
			Name:   "create-address-list-import-job",
			Fields: fields_create_address_list_import_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAddressListImportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_address_list_import_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAddressListImportJob(ctx, input)
			},
		},
		"create-archive": {
			Name:   "create-archive",
			Fields: fields_create_archive,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateArchiveInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_archive, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateArchive(ctx, input)
			},
		},
		"create-ingress-point": {
			Name:   "create-ingress-point",
			Fields: fields_create_ingress_point,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateIngressPointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_ingress_point, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateIngressPoint(ctx, input)
			},
		},
		"create-relay": {
			Name:   "create-relay",
			Fields: fields_create_relay,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRelayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_relay, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRelay(ctx, input)
			},
		},
		"create-rule-set": {
			Name:   "create-rule-set",
			Fields: fields_create_rule_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRuleSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_rule_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRuleSet(ctx, input)
			},
		},
		"create-traffic-policy": {
			Name:   "create-traffic-policy",
			Fields: fields_create_traffic_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTrafficPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_traffic_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTrafficPolicy(ctx, input)
			},
		},
		"delete-addon-instance": {
			Name:   "delete-addon-instance",
			Fields: fields_delete_addon_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAddonInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_addon_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAddonInstance(ctx, input)
			},
		},
		"delete-addon-subscription": {
			Name:   "delete-addon-subscription",
			Fields: fields_delete_addon_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAddonSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_addon_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAddonSubscription(ctx, input)
			},
		},
		"delete-address-list": {
			Name:   "delete-address-list",
			Fields: fields_delete_address_list,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAddressListInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_address_list, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAddressList(ctx, input)
			},
		},
		"delete-archive": {
			Name:   "delete-archive",
			Fields: fields_delete_archive,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteArchiveInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_archive, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteArchive(ctx, input)
			},
		},
		"delete-ingress-point": {
			Name:   "delete-ingress-point",
			Fields: fields_delete_ingress_point,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIngressPointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_ingress_point, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIngressPoint(ctx, input)
			},
		},
		"delete-relay": {
			Name:   "delete-relay",
			Fields: fields_delete_relay,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRelayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_relay, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRelay(ctx, input)
			},
		},
		"delete-rule-set": {
			Name:   "delete-rule-set",
			Fields: fields_delete_rule_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRuleSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_rule_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRuleSet(ctx, input)
			},
		},
		"delete-traffic-policy": {
			Name:   "delete-traffic-policy",
			Fields: fields_delete_traffic_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTrafficPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_traffic_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTrafficPolicy(ctx, input)
			},
		},
		"deregister-member-from-address-list": {
			Name:   "deregister-member-from-address-list",
			Fields: fields_deregister_member_from_address_list,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterMemberFromAddressListInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_member_from_address_list, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterMemberFromAddressList(ctx, input)
			},
		},
		"get-addon-instance": {
			Name:   "get-addon-instance",
			Fields: fields_get_addon_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAddonInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_addon_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAddonInstance(ctx, input)
			},
		},
		"get-addon-subscription": {
			Name:   "get-addon-subscription",
			Fields: fields_get_addon_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAddonSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_addon_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAddonSubscription(ctx, input)
			},
		},
		"get-address-list": {
			Name:   "get-address-list",
			Fields: fields_get_address_list,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAddressListInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_address_list, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAddressList(ctx, input)
			},
		},
		"get-address-list-import-job": {
			Name:   "get-address-list-import-job",
			Fields: fields_get_address_list_import_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAddressListImportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_address_list_import_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAddressListImportJob(ctx, input)
			},
		},
		"get-archive": {
			Name:   "get-archive",
			Fields: fields_get_archive,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetArchiveInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_archive, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetArchive(ctx, input)
			},
		},
		"get-archive-export": {
			Name:   "get-archive-export",
			Fields: fields_get_archive_export,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetArchiveExportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_archive_export, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetArchiveExport(ctx, input)
			},
		},
		"get-archive-message": {
			Name:   "get-archive-message",
			Fields: fields_get_archive_message,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetArchiveMessageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_archive_message, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetArchiveMessage(ctx, input)
			},
		},
		"get-archive-message-content": {
			Name:   "get-archive-message-content",
			Fields: fields_get_archive_message_content,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetArchiveMessageContentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_archive_message_content, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetArchiveMessageContent(ctx, input)
			},
		},
		"get-archive-search": {
			Name:   "get-archive-search",
			Fields: fields_get_archive_search,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetArchiveSearchInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_archive_search, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetArchiveSearch(ctx, input)
			},
		},
		"get-archive-search-results": {
			Name:   "get-archive-search-results",
			Fields: fields_get_archive_search_results,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetArchiveSearchResultsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_archive_search_results, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetArchiveSearchResults(ctx, input)
			},
		},
		"get-ingress-point": {
			Name:   "get-ingress-point",
			Fields: fields_get_ingress_point,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIngressPointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_ingress_point, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIngressPoint(ctx, input)
			},
		},
		"get-member-of-address-list": {
			Name:   "get-member-of-address-list",
			Fields: fields_get_member_of_address_list,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMemberOfAddressListInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_member_of_address_list, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMemberOfAddressList(ctx, input)
			},
		},
		"get-relay": {
			Name:   "get-relay",
			Fields: fields_get_relay,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRelayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_relay, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRelay(ctx, input)
			},
		},
		"get-rule-set": {
			Name:   "get-rule-set",
			Fields: fields_get_rule_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRuleSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_rule_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRuleSet(ctx, input)
			},
		},
		"get-traffic-policy": {
			Name:   "get-traffic-policy",
			Fields: fields_get_traffic_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTrafficPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_traffic_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTrafficPolicy(ctx, input)
			},
		},
		"list-addon-instances": {
			Name:   "list-addon-instances",
			Fields: fields_list_addon_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAddonInstancesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_addon_instances, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAddonInstances(ctx, input)
				}
				var results []*svc.ListAddonInstancesOutput
				p := svc.NewListAddonInstancesPaginator(client, input)
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
		"list-addon-subscriptions": {
			Name:   "list-addon-subscriptions",
			Fields: fields_list_addon_subscriptions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAddonSubscriptionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_addon_subscriptions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAddonSubscriptions(ctx, input)
				}
				var results []*svc.ListAddonSubscriptionsOutput
				p := svc.NewListAddonSubscriptionsPaginator(client, input)
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
		"list-address-list-import-jobs": {
			Name:   "list-address-list-import-jobs",
			Fields: fields_list_address_list_import_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAddressListImportJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_address_list_import_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAddressListImportJobs(ctx, input)
				}
				var results []*svc.ListAddressListImportJobsOutput
				p := svc.NewListAddressListImportJobsPaginator(client, input)
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
		"list-address-lists": {
			Name:   "list-address-lists",
			Fields: fields_list_address_lists,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAddressListsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_address_lists, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAddressLists(ctx, input)
				}
				var results []*svc.ListAddressListsOutput
				p := svc.NewListAddressListsPaginator(client, input)
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
		"list-archive-exports": {
			Name:   "list-archive-exports",
			Fields: fields_list_archive_exports,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListArchiveExportsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_archive_exports, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListArchiveExports(ctx, input)
				}
				var results []*svc.ListArchiveExportsOutput
				p := svc.NewListArchiveExportsPaginator(client, input)
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
		"list-archive-searches": {
			Name:   "list-archive-searches",
			Fields: fields_list_archive_searches,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListArchiveSearchesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_archive_searches, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListArchiveSearches(ctx, input)
				}
				var results []*svc.ListArchiveSearchesOutput
				p := svc.NewListArchiveSearchesPaginator(client, input)
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
		"list-archives": {
			Name:   "list-archives",
			Fields: fields_list_archives,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListArchivesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_archives, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListArchives(ctx, input)
				}
				var results []*svc.ListArchivesOutput
				p := svc.NewListArchivesPaginator(client, input)
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
		"list-ingress-points": {
			Name:   "list-ingress-points",
			Fields: fields_list_ingress_points,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIngressPointsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_ingress_points, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListIngressPoints(ctx, input)
				}
				var results []*svc.ListIngressPointsOutput
				p := svc.NewListIngressPointsPaginator(client, input)
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
		"list-members-of-address-list": {
			Name:   "list-members-of-address-list",
			Fields: fields_list_members_of_address_list,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMembersOfAddressListInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_members_of_address_list, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMembersOfAddressList(ctx, input)
				}
				var results []*svc.ListMembersOfAddressListOutput
				p := svc.NewListMembersOfAddressListPaginator(client, input)
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
		"list-relays": {
			Name:   "list-relays",
			Fields: fields_list_relays,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRelaysInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_relays, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRelays(ctx, input)
				}
				var results []*svc.ListRelaysOutput
				p := svc.NewListRelaysPaginator(client, input)
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
		"list-rule-sets": {
			Name:   "list-rule-sets",
			Fields: fields_list_rule_sets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRuleSetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_rule_sets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRuleSets(ctx, input)
				}
				var results []*svc.ListRuleSetsOutput
				p := svc.NewListRuleSetsPaginator(client, input)
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
		"list-traffic-policies": {
			Name:   "list-traffic-policies",
			Fields: fields_list_traffic_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTrafficPoliciesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_traffic_policies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTrafficPolicies(ctx, input)
				}
				var results []*svc.ListTrafficPoliciesOutput
				p := svc.NewListTrafficPoliciesPaginator(client, input)
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
		"register-member-to-address-list": {
			Name:   "register-member-to-address-list",
			Fields: fields_register_member_to_address_list,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterMemberToAddressListInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_member_to_address_list, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterMemberToAddressList(ctx, input)
			},
		},
		"start-address-list-import-job": {
			Name:   "start-address-list-import-job",
			Fields: fields_start_address_list_import_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartAddressListImportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_address_list_import_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartAddressListImportJob(ctx, input)
			},
		},
		"start-archive-export": {
			Name:   "start-archive-export",
			Fields: fields_start_archive_export,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartArchiveExportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_archive_export, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartArchiveExport(ctx, input)
			},
		},
		"start-archive-search": {
			Name:   "start-archive-search",
			Fields: fields_start_archive_search,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartArchiveSearchInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_archive_search, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartArchiveSearch(ctx, input)
			},
		},
		"stop-address-list-import-job": {
			Name:   "stop-address-list-import-job",
			Fields: fields_stop_address_list_import_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopAddressListImportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_address_list_import_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopAddressListImportJob(ctx, input)
			},
		},
		"stop-archive-export": {
			Name:   "stop-archive-export",
			Fields: fields_stop_archive_export,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopArchiveExportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_archive_export, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopArchiveExport(ctx, input)
			},
		},
		"stop-archive-search": {
			Name:   "stop-archive-search",
			Fields: fields_stop_archive_search,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopArchiveSearchInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_archive_search, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopArchiveSearch(ctx, input)
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
		"update-archive": {
			Name:   "update-archive",
			Fields: fields_update_archive,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateArchiveInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_archive, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateArchive(ctx, input)
			},
		},
		"update-ingress-point": {
			Name:   "update-ingress-point",
			Fields: fields_update_ingress_point,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateIngressPointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_ingress_point, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateIngressPoint(ctx, input)
			},
		},
		"update-relay": {
			Name:   "update-relay",
			Fields: fields_update_relay,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRelayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_relay, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRelay(ctx, input)
			},
		},
		"update-rule-set": {
			Name:   "update-rule-set",
			Fields: fields_update_rule_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRuleSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_rule_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRuleSet(ctx, input)
			},
		},
		"update-traffic-policy": {
			Name:   "update-traffic-policy",
			Fields: fields_update_traffic_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTrafficPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_traffic_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTrafficPolicy(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("mailmanager", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
