package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/notifications"
)

var fields_associate_channel = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "NotificationConfigurationArn", Flag: "notification-configuration-arn", Type: "*string", Required: true},
}

var fields_associate_managed_notification_account_contact = []leanruntime.Field{
	{Name: "ContactIdentifier", Flag: "contact-identifier", Type: "types.AccountContactType", Required: true},
	{Name: "ManagedNotificationConfigurationArn", Flag: "managed-notification-configuration-arn", Type: "*string", Required: true},
}

var fields_associate_managed_notification_additional_channel = []leanruntime.Field{
	{Name: "ChannelArn", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "ManagedNotificationConfigurationArn", Flag: "managed-notification-configuration-arn", Type: "*string", Required: true},
}

var fields_associate_organizational_unit = []leanruntime.Field{
	{Name: "NotificationConfigurationArn", Flag: "notification-configuration-arn", Type: "*string", Required: true},
	{Name: "OrganizationalUnitId", Flag: "organizational-unit-id", Type: "*string", Required: true},
}

var fields_create_event_rule = []leanruntime.Field{
	{Name: "EventPattern", Flag: "event-pattern", Type: "*string", Required: false},
	{Name: "EventType", Flag: "event-type", Type: "*string", Required: true},
	{Name: "NotificationConfigurationArn", Flag: "notification-configuration-arn", Type: "*string", Required: true},
	{Name: "Regions", Flag: "regions", Type: "[]string", Required: true},
	{Name: "Source", Flag: "source", Type: "*string", Required: true},
}

var fields_create_notification_configuration = []leanruntime.Field{
	{Name: "AggregationDuration", Flag: "aggregation-duration", Type: "types.AggregationDuration", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_event_rule = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_delete_notification_configuration = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_deregister_notification_hub = []leanruntime.Field{
	{Name: "NotificationHubRegion", Flag: "notification-hub-region", Type: "*string", Required: true},
}

var fields_disable_notifications_access_for_organization = []leanruntime.Field{}

var fields_disassociate_channel = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "NotificationConfigurationArn", Flag: "notification-configuration-arn", Type: "*string", Required: true},
}

var fields_disassociate_managed_notification_account_contact = []leanruntime.Field{
	{Name: "ContactIdentifier", Flag: "contact-identifier", Type: "types.AccountContactType", Required: true},
	{Name: "ManagedNotificationConfigurationArn", Flag: "managed-notification-configuration-arn", Type: "*string", Required: true},
}

var fields_disassociate_managed_notification_additional_channel = []leanruntime.Field{
	{Name: "ChannelArn", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "ManagedNotificationConfigurationArn", Flag: "managed-notification-configuration-arn", Type: "*string", Required: true},
}

var fields_disassociate_organizational_unit = []leanruntime.Field{
	{Name: "NotificationConfigurationArn", Flag: "notification-configuration-arn", Type: "*string", Required: true},
	{Name: "OrganizationalUnitId", Flag: "organizational-unit-id", Type: "*string", Required: true},
}

var fields_enable_notifications_access_for_organization = []leanruntime.Field{}

var fields_get_event_rule = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_managed_notification_child_event = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "Locale", Flag: "locale", Type: "types.LocaleCode", Required: false},
}

var fields_get_managed_notification_configuration = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_managed_notification_event = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "Locale", Flag: "locale", Type: "types.LocaleCode", Required: false},
}

var fields_get_notification_configuration = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_notification_event = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "Locale", Flag: "locale", Type: "types.LocaleCode", Required: false},
}

var fields_get_notifications_access_for_organization = []leanruntime.Field{}

var fields_list_channels = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "NotificationConfigurationArn", Flag: "notification-configuration-arn", Type: "*string", Required: true},
}

var fields_list_event_rules = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "NotificationConfigurationArn", Flag: "notification-configuration-arn", Type: "*string", Required: true},
}

var fields_list_managed_notification_channel_associations = []leanruntime.Field{
	{Name: "ManagedNotificationConfigurationArn", Flag: "managed-notification-configuration-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_managed_notification_child_events = []leanruntime.Field{
	{Name: "AggregateManagedNotificationEventArn", Flag: "aggregate-managed-notification-event-arn", Type: "*string", Required: true},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "Locale", Flag: "locale", Type: "types.LocaleCode", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OrganizationalUnitId", Flag: "organizational-unit-id", Type: "*string", Required: false},
	{Name: "RelatedAccount", Flag: "related-account", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
}

var fields_list_managed_notification_configurations = []leanruntime.Field{
	{Name: "ChannelIdentifier", Flag: "channel-identifier", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_managed_notification_events = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "Locale", Flag: "locale", Type: "types.LocaleCode", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OrganizationalUnitId", Flag: "organizational-unit-id", Type: "*string", Required: false},
	{Name: "RelatedAccount", Flag: "related-account", Type: "*string", Required: false},
	{Name: "Source", Flag: "source", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
}

var fields_list_member_accounts = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MemberAccount", Flag: "member-account", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "NotificationConfigurationArn", Flag: "notification-configuration-arn", Type: "*string", Required: true},
	{Name: "OrganizationalUnitId", Flag: "organizational-unit-id", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.MemberAccountNotificationConfigurationStatus", Required: false},
}

var fields_list_notification_configurations = []leanruntime.Field{
	{Name: "ChannelArn", Flag: "channel-arn", Type: "*string", Required: false},
	{Name: "EventRuleSource", Flag: "event-rule-source", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.NotificationConfigurationStatus", Required: false},
	{Name: "Subtype", Flag: "subtype", Type: "types.NotificationConfigurationSubtype", Required: false},
}

var fields_list_notification_events = []leanruntime.Field{
	{Name: "AggregateNotificationEventArn", Flag: "aggregate-notification-event-arn", Type: "*string", Required: false},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "IncludeChildEvents", Flag: "include-child-events", Type: "*bool", Required: false},
	{Name: "Locale", Flag: "locale", Type: "types.LocaleCode", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OrganizationalUnitId", Flag: "organizational-unit-id", Type: "*string", Required: false},
	{Name: "Source", Flag: "source", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
}

var fields_list_notification_hubs = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_organizational_units = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "NotificationConfigurationArn", Flag: "notification-configuration-arn", Type: "*string", Required: true},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_register_notification_hub = []leanruntime.Field{
	{Name: "NotificationHubRegion", Flag: "notification-hub-region", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_event_rule = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "EventPattern", Flag: "event-pattern", Type: "*string", Required: false},
	{Name: "Regions", Flag: "regions", Type: "[]string", Required: false},
}

var fields_update_notification_configuration = []leanruntime.Field{
	{Name: "AggregationDuration", Flag: "aggregation-duration", Type: "types.AggregationDuration", Required: false},
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-channel": {
			Name:   "associate-channel",
			Fields: fields_associate_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateChannel(ctx, input)
			},
		},
		"associate-managed-notification-account-contact": {
			Name:   "associate-managed-notification-account-contact",
			Fields: fields_associate_managed_notification_account_contact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateManagedNotificationAccountContactInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_managed_notification_account_contact, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateManagedNotificationAccountContact(ctx, input)
			},
		},
		"associate-managed-notification-additional-channel": {
			Name:   "associate-managed-notification-additional-channel",
			Fields: fields_associate_managed_notification_additional_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateManagedNotificationAdditionalChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_managed_notification_additional_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateManagedNotificationAdditionalChannel(ctx, input)
			},
		},
		"associate-organizational-unit": {
			Name:   "associate-organizational-unit",
			Fields: fields_associate_organizational_unit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateOrganizationalUnitInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_organizational_unit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateOrganizationalUnit(ctx, input)
			},
		},
		"create-event-rule": {
			Name:   "create-event-rule",
			Fields: fields_create_event_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEventRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_event_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEventRule(ctx, input)
			},
		},
		"create-notification-configuration": {
			Name:   "create-notification-configuration",
			Fields: fields_create_notification_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateNotificationConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_notification_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateNotificationConfiguration(ctx, input)
			},
		},
		"delete-event-rule": {
			Name:   "delete-event-rule",
			Fields: fields_delete_event_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEventRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_event_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEventRule(ctx, input)
			},
		},
		"delete-notification-configuration": {
			Name:   "delete-notification-configuration",
			Fields: fields_delete_notification_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteNotificationConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_notification_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteNotificationConfiguration(ctx, input)
			},
		},
		"deregister-notification-hub": {
			Name:   "deregister-notification-hub",
			Fields: fields_deregister_notification_hub,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterNotificationHubInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_notification_hub, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterNotificationHub(ctx, input)
			},
		},
		"disable-notifications-access-for-organization": {
			Name:   "disable-notifications-access-for-organization",
			Fields: fields_disable_notifications_access_for_organization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableNotificationsAccessForOrganizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_notifications_access_for_organization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableNotificationsAccessForOrganization(ctx, input)
			},
		},
		"disassociate-channel": {
			Name:   "disassociate-channel",
			Fields: fields_disassociate_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateChannel(ctx, input)
			},
		},
		"disassociate-managed-notification-account-contact": {
			Name:   "disassociate-managed-notification-account-contact",
			Fields: fields_disassociate_managed_notification_account_contact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateManagedNotificationAccountContactInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_managed_notification_account_contact, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateManagedNotificationAccountContact(ctx, input)
			},
		},
		"disassociate-managed-notification-additional-channel": {
			Name:   "disassociate-managed-notification-additional-channel",
			Fields: fields_disassociate_managed_notification_additional_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateManagedNotificationAdditionalChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_managed_notification_additional_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateManagedNotificationAdditionalChannel(ctx, input)
			},
		},
		"disassociate-organizational-unit": {
			Name:   "disassociate-organizational-unit",
			Fields: fields_disassociate_organizational_unit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateOrganizationalUnitInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_organizational_unit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateOrganizationalUnit(ctx, input)
			},
		},
		"enable-notifications-access-for-organization": {
			Name:   "enable-notifications-access-for-organization",
			Fields: fields_enable_notifications_access_for_organization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableNotificationsAccessForOrganizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_notifications_access_for_organization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableNotificationsAccessForOrganization(ctx, input)
			},
		},
		"get-event-rule": {
			Name:   "get-event-rule",
			Fields: fields_get_event_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEventRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_event_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEventRule(ctx, input)
			},
		},
		"get-managed-notification-child-event": {
			Name:   "get-managed-notification-child-event",
			Fields: fields_get_managed_notification_child_event,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetManagedNotificationChildEventInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_managed_notification_child_event, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetManagedNotificationChildEvent(ctx, input)
			},
		},
		"get-managed-notification-configuration": {
			Name:   "get-managed-notification-configuration",
			Fields: fields_get_managed_notification_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetManagedNotificationConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_managed_notification_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetManagedNotificationConfiguration(ctx, input)
			},
		},
		"get-managed-notification-event": {
			Name:   "get-managed-notification-event",
			Fields: fields_get_managed_notification_event,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetManagedNotificationEventInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_managed_notification_event, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetManagedNotificationEvent(ctx, input)
			},
		},
		"get-notification-configuration": {
			Name:   "get-notification-configuration",
			Fields: fields_get_notification_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetNotificationConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_notification_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetNotificationConfiguration(ctx, input)
			},
		},
		"get-notification-event": {
			Name:   "get-notification-event",
			Fields: fields_get_notification_event,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetNotificationEventInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_notification_event, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetNotificationEvent(ctx, input)
			},
		},
		"get-notifications-access-for-organization": {
			Name:   "get-notifications-access-for-organization",
			Fields: fields_get_notifications_access_for_organization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetNotificationsAccessForOrganizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_notifications_access_for_organization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetNotificationsAccessForOrganization(ctx, input)
			},
		},
		"list-channels": {
			Name:   "list-channels",
			Fields: fields_list_channels,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListChannelsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_channels, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListChannels(ctx, input)
				}
				var results []*svc.ListChannelsOutput
				p := svc.NewListChannelsPaginator(client, input)
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
		"list-event-rules": {
			Name:   "list-event-rules",
			Fields: fields_list_event_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEventRulesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_event_rules, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEventRules(ctx, input)
				}
				var results []*svc.ListEventRulesOutput
				p := svc.NewListEventRulesPaginator(client, input)
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
		"list-managed-notification-channel-associations": {
			Name:   "list-managed-notification-channel-associations",
			Fields: fields_list_managed_notification_channel_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListManagedNotificationChannelAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_managed_notification_channel_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListManagedNotificationChannelAssociations(ctx, input)
				}
				var results []*svc.ListManagedNotificationChannelAssociationsOutput
				p := svc.NewListManagedNotificationChannelAssociationsPaginator(client, input)
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
		"list-managed-notification-child-events": {
			Name:   "list-managed-notification-child-events",
			Fields: fields_list_managed_notification_child_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListManagedNotificationChildEventsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_managed_notification_child_events, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListManagedNotificationChildEvents(ctx, input)
				}
				var results []*svc.ListManagedNotificationChildEventsOutput
				p := svc.NewListManagedNotificationChildEventsPaginator(client, input)
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
		"list-managed-notification-configurations": {
			Name:   "list-managed-notification-configurations",
			Fields: fields_list_managed_notification_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListManagedNotificationConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_managed_notification_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListManagedNotificationConfigurations(ctx, input)
				}
				var results []*svc.ListManagedNotificationConfigurationsOutput
				p := svc.NewListManagedNotificationConfigurationsPaginator(client, input)
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
		"list-managed-notification-events": {
			Name:   "list-managed-notification-events",
			Fields: fields_list_managed_notification_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListManagedNotificationEventsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_managed_notification_events, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListManagedNotificationEvents(ctx, input)
				}
				var results []*svc.ListManagedNotificationEventsOutput
				p := svc.NewListManagedNotificationEventsPaginator(client, input)
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
		"list-member-accounts": {
			Name:   "list-member-accounts",
			Fields: fields_list_member_accounts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMemberAccountsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_member_accounts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMemberAccounts(ctx, input)
				}
				var results []*svc.ListMemberAccountsOutput
				p := svc.NewListMemberAccountsPaginator(client, input)
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
		"list-notification-configurations": {
			Name:   "list-notification-configurations",
			Fields: fields_list_notification_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListNotificationConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_notification_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListNotificationConfigurations(ctx, input)
				}
				var results []*svc.ListNotificationConfigurationsOutput
				p := svc.NewListNotificationConfigurationsPaginator(client, input)
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
		"list-notification-events": {
			Name:   "list-notification-events",
			Fields: fields_list_notification_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListNotificationEventsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_notification_events, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListNotificationEvents(ctx, input)
				}
				var results []*svc.ListNotificationEventsOutput
				p := svc.NewListNotificationEventsPaginator(client, input)
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
		"list-notification-hubs": {
			Name:   "list-notification-hubs",
			Fields: fields_list_notification_hubs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListNotificationHubsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_notification_hubs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListNotificationHubs(ctx, input)
				}
				var results []*svc.ListNotificationHubsOutput
				p := svc.NewListNotificationHubsPaginator(client, input)
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
		"list-organizational-units": {
			Name:   "list-organizational-units",
			Fields: fields_list_organizational_units,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOrganizationalUnitsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_organizational_units, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListOrganizationalUnits(ctx, input)
				}
				var results []*svc.ListOrganizationalUnitsOutput
				p := svc.NewListOrganizationalUnitsPaginator(client, input)
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
		"register-notification-hub": {
			Name:   "register-notification-hub",
			Fields: fields_register_notification_hub,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterNotificationHubInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_notification_hub, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterNotificationHub(ctx, input)
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
		"update-event-rule": {
			Name:   "update-event-rule",
			Fields: fields_update_event_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEventRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_event_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEventRule(ctx, input)
			},
		},
		"update-notification-configuration": {
			Name:   "update-notification-configuration",
			Fields: fields_update_notification_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateNotificationConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_notification_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateNotificationConfiguration(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("notifications", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
