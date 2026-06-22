package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/pinpointemail"
)

var fields_create_configuration_set = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
	{Name: "DeliveryOptions", Flag: "delivery-options", Type: "*types.DeliveryOptions", Required: false},
	{Name: "ReputationOptions", Flag: "reputation-options", Type: "*types.ReputationOptions", Required: false},
	{Name: "SendingOptions", Flag: "sending-options", Type: "*types.SendingOptions", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TrackingOptions", Flag: "tracking-options", Type: "*types.TrackingOptions", Required: false},
}

var fields_create_configuration_set_event_destination = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
	{Name: "EventDestination", Flag: "event-destination", Type: "*types.EventDestinationDefinition", Required: true},
	{Name: "EventDestinationName", Flag: "event-destination-name", Type: "*string", Required: true},
}

var fields_create_dedicated_ip_pool = []leanruntime.Field{
	{Name: "PoolName", Flag: "pool-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_deliverability_test_report = []leanruntime.Field{
	{Name: "Content", Flag: "content", Type: "*types.EmailContent", Required: true},
	{Name: "FromEmailAddress", Flag: "from-email-address", Type: "*string", Required: true},
	{Name: "ReportName", Flag: "report-name", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_email_identity = []leanruntime.Field{
	{Name: "EmailIdentity", Flag: "email-identity", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_delete_configuration_set = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
}

var fields_delete_configuration_set_event_destination = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
	{Name: "EventDestinationName", Flag: "event-destination-name", Type: "*string", Required: true},
}

var fields_delete_dedicated_ip_pool = []leanruntime.Field{
	{Name: "PoolName", Flag: "pool-name", Type: "*string", Required: true},
}

var fields_delete_email_identity = []leanruntime.Field{
	{Name: "EmailIdentity", Flag: "email-identity", Type: "*string", Required: true},
}

var fields_get_account = []leanruntime.Field{}

var fields_get_blacklist_reports = []leanruntime.Field{
	{Name: "BlacklistItemNames", Flag: "blacklist-item-names", Type: "[]string", Required: true},
}

var fields_get_configuration_set = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
}

var fields_get_configuration_set_event_destinations = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
}

var fields_get_dedicated_ip = []leanruntime.Field{
	{Name: "Ip", Flag: "ip", Type: "*string", Required: true},
}

var fields_get_dedicated_ips = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
	{Name: "PoolName", Flag: "pool-name", Type: "*string", Required: false},
}

var fields_get_deliverability_dashboard_options = []leanruntime.Field{}

var fields_get_deliverability_test_report = []leanruntime.Field{
	{Name: "ReportId", Flag: "report-id", Type: "*string", Required: true},
}

var fields_get_domain_deliverability_campaign = []leanruntime.Field{
	{Name: "CampaignId", Flag: "campaign-id", Type: "*string", Required: true},
}

var fields_get_domain_statistics_report = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "EndDate", Flag: "end-date", Type: "*time.Time", Required: true},
	{Name: "StartDate", Flag: "start-date", Type: "*time.Time", Required: true},
}

var fields_get_email_identity = []leanruntime.Field{
	{Name: "EmailIdentity", Flag: "email-identity", Type: "*string", Required: true},
}

var fields_list_configuration_sets = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
}

var fields_list_dedicated_ip_pools = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
}

var fields_list_deliverability_test_reports = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
}

var fields_list_domain_deliverability_campaigns = []leanruntime.Field{
	{Name: "EndDate", Flag: "end-date", Type: "*time.Time", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
	{Name: "StartDate", Flag: "start-date", Type: "*time.Time", Required: true},
	{Name: "SubscribedDomain", Flag: "subscribed-domain", Type: "*string", Required: true},
}

var fields_list_email_identities = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_put_account_dedicated_ip_warmup_attributes = []leanruntime.Field{
	{Name: "AutoWarmupEnabled", Flag: "auto-warmup-enabled", Type: "bool", Required: false},
}

var fields_put_account_sending_attributes = []leanruntime.Field{
	{Name: "SendingEnabled", Flag: "sending-enabled", Type: "bool", Required: false},
}

var fields_put_configuration_set_delivery_options = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
	{Name: "SendingPoolName", Flag: "sending-pool-name", Type: "*string", Required: false},
	{Name: "TlsPolicy", Flag: "tls-policy", Type: "types.TlsPolicy", Required: false},
}

var fields_put_configuration_set_reputation_options = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
	{Name: "ReputationMetricsEnabled", Flag: "reputation-metrics-enabled", Type: "bool", Required: false},
}

var fields_put_configuration_set_sending_options = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
	{Name: "SendingEnabled", Flag: "sending-enabled", Type: "bool", Required: false},
}

var fields_put_configuration_set_tracking_options = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
	{Name: "CustomRedirectDomain", Flag: "custom-redirect-domain", Type: "*string", Required: false},
}

var fields_put_dedicated_ip_in_pool = []leanruntime.Field{
	{Name: "DestinationPoolName", Flag: "destination-pool-name", Type: "*string", Required: true},
	{Name: "Ip", Flag: "ip", Type: "*string", Required: true},
}

var fields_put_dedicated_ip_warmup_attributes = []leanruntime.Field{
	{Name: "Ip", Flag: "ip", Type: "*string", Required: true},
	{Name: "WarmupPercentage", Flag: "warmup-percentage", Type: "*int32", Required: true},
}

var fields_put_deliverability_dashboard_option = []leanruntime.Field{
	{Name: "DashboardEnabled", Flag: "dashboard-enabled", Type: "bool", Required: true},
	{Name: "SubscribedDomains", Flag: "subscribed-domains", Type: "[]types.DomainDeliverabilityTrackingOption", Required: false},
}

var fields_put_email_identity_dkim_attributes = []leanruntime.Field{
	{Name: "EmailIdentity", Flag: "email-identity", Type: "*string", Required: true},
	{Name: "SigningEnabled", Flag: "signing-enabled", Type: "bool", Required: false},
}

var fields_put_email_identity_feedback_attributes = []leanruntime.Field{
	{Name: "EmailForwardingEnabled", Flag: "email-forwarding-enabled", Type: "bool", Required: false},
	{Name: "EmailIdentity", Flag: "email-identity", Type: "*string", Required: true},
}

var fields_put_email_identity_mail_from_attributes = []leanruntime.Field{
	{Name: "BehaviorOnMxFailure", Flag: "behavior-on-mx-failure", Type: "types.BehaviorOnMxFailure", Required: false},
	{Name: "EmailIdentity", Flag: "email-identity", Type: "*string", Required: true},
	{Name: "MailFromDomain", Flag: "mail-from-domain", Type: "*string", Required: false},
}

var fields_send_email = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: false},
	{Name: "Content", Flag: "content", Type: "*types.EmailContent", Required: true},
	{Name: "Destination", Flag: "destination", Type: "*types.Destination", Required: true},
	{Name: "EmailTags", Flag: "email-tags", Type: "[]types.MessageTag", Required: false},
	{Name: "FeedbackForwardingEmailAddress", Flag: "feedback-forwarding-email-address", Type: "*string", Required: false},
	{Name: "FromEmailAddress", Flag: "from-email-address", Type: "*string", Required: false},
	{Name: "ReplyToAddresses", Flag: "reply-to-addresses", Type: "[]string", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_configuration_set_event_destination = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
	{Name: "EventDestination", Flag: "event-destination", Type: "*types.EventDestinationDefinition", Required: true},
	{Name: "EventDestinationName", Flag: "event-destination-name", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-configuration-set": {
			Name:   "create-configuration-set",
			Fields: fields_create_configuration_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateConfigurationSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_configuration_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateConfigurationSet(ctx, input)
			},
		},
		"create-configuration-set-event-destination": {
			Name:   "create-configuration-set-event-destination",
			Fields: fields_create_configuration_set_event_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateConfigurationSetEventDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_configuration_set_event_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateConfigurationSetEventDestination(ctx, input)
			},
		},
		"create-dedicated-ip-pool": {
			Name:   "create-dedicated-ip-pool",
			Fields: fields_create_dedicated_ip_pool,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDedicatedIpPoolInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_dedicated_ip_pool, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDedicatedIpPool(ctx, input)
			},
		},
		"create-deliverability-test-report": {
			Name:   "create-deliverability-test-report",
			Fields: fields_create_deliverability_test_report,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDeliverabilityTestReportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_deliverability_test_report, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDeliverabilityTestReport(ctx, input)
			},
		},
		"create-email-identity": {
			Name:   "create-email-identity",
			Fields: fields_create_email_identity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEmailIdentityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_email_identity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEmailIdentity(ctx, input)
			},
		},
		"delete-configuration-set": {
			Name:   "delete-configuration-set",
			Fields: fields_delete_configuration_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConfigurationSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_configuration_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConfigurationSet(ctx, input)
			},
		},
		"delete-configuration-set-event-destination": {
			Name:   "delete-configuration-set-event-destination",
			Fields: fields_delete_configuration_set_event_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConfigurationSetEventDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_configuration_set_event_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConfigurationSetEventDestination(ctx, input)
			},
		},
		"delete-dedicated-ip-pool": {
			Name:   "delete-dedicated-ip-pool",
			Fields: fields_delete_dedicated_ip_pool,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDedicatedIpPoolInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_dedicated_ip_pool, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDedicatedIpPool(ctx, input)
			},
		},
		"delete-email-identity": {
			Name:   "delete-email-identity",
			Fields: fields_delete_email_identity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEmailIdentityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_email_identity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEmailIdentity(ctx, input)
			},
		},
		"get-account": {
			Name:   "get-account",
			Fields: fields_get_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccount(ctx, input)
			},
		},
		"get-blacklist-reports": {
			Name:   "get-blacklist-reports",
			Fields: fields_get_blacklist_reports,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBlacklistReportsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_blacklist_reports, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBlacklistReports(ctx, input)
			},
		},
		"get-configuration-set": {
			Name:   "get-configuration-set",
			Fields: fields_get_configuration_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConfigurationSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_configuration_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConfigurationSet(ctx, input)
			},
		},
		"get-configuration-set-event-destinations": {
			Name:   "get-configuration-set-event-destinations",
			Fields: fields_get_configuration_set_event_destinations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConfigurationSetEventDestinationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_configuration_set_event_destinations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConfigurationSetEventDestinations(ctx, input)
			},
		},
		"get-dedicated-ip": {
			Name:   "get-dedicated-ip",
			Fields: fields_get_dedicated_ip,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDedicatedIpInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_dedicated_ip, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDedicatedIp(ctx, input)
			},
		},
		"get-dedicated-ips": {
			Name:   "get-dedicated-ips",
			Fields: fields_get_dedicated_ips,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDedicatedIpsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_dedicated_ips, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetDedicatedIps(ctx, input)
				}
				var results []*svc.GetDedicatedIpsOutput
				p := svc.NewGetDedicatedIpsPaginator(client, input)
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
		"get-deliverability-dashboard-options": {
			Name:   "get-deliverability-dashboard-options",
			Fields: fields_get_deliverability_dashboard_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDeliverabilityDashboardOptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_deliverability_dashboard_options, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDeliverabilityDashboardOptions(ctx, input)
			},
		},
		"get-deliverability-test-report": {
			Name:   "get-deliverability-test-report",
			Fields: fields_get_deliverability_test_report,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDeliverabilityTestReportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_deliverability_test_report, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDeliverabilityTestReport(ctx, input)
			},
		},
		"get-domain-deliverability-campaign": {
			Name:   "get-domain-deliverability-campaign",
			Fields: fields_get_domain_deliverability_campaign,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDomainDeliverabilityCampaignInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_domain_deliverability_campaign, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDomainDeliverabilityCampaign(ctx, input)
			},
		},
		"get-domain-statistics-report": {
			Name:   "get-domain-statistics-report",
			Fields: fields_get_domain_statistics_report,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDomainStatisticsReportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_domain_statistics_report, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDomainStatisticsReport(ctx, input)
			},
		},
		"get-email-identity": {
			Name:   "get-email-identity",
			Fields: fields_get_email_identity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEmailIdentityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_email_identity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEmailIdentity(ctx, input)
			},
		},
		"list-configuration-sets": {
			Name:   "list-configuration-sets",
			Fields: fields_list_configuration_sets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConfigurationSetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_configuration_sets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListConfigurationSets(ctx, input)
				}
				var results []*svc.ListConfigurationSetsOutput
				p := svc.NewListConfigurationSetsPaginator(client, input)
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
		"list-dedicated-ip-pools": {
			Name:   "list-dedicated-ip-pools",
			Fields: fields_list_dedicated_ip_pools,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDedicatedIpPoolsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_dedicated_ip_pools, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDedicatedIpPools(ctx, input)
				}
				var results []*svc.ListDedicatedIpPoolsOutput
				p := svc.NewListDedicatedIpPoolsPaginator(client, input)
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
		"list-deliverability-test-reports": {
			Name:   "list-deliverability-test-reports",
			Fields: fields_list_deliverability_test_reports,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDeliverabilityTestReportsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_deliverability_test_reports, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDeliverabilityTestReports(ctx, input)
				}
				var results []*svc.ListDeliverabilityTestReportsOutput
				p := svc.NewListDeliverabilityTestReportsPaginator(client, input)
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
		"list-domain-deliverability-campaigns": {
			Name:   "list-domain-deliverability-campaigns",
			Fields: fields_list_domain_deliverability_campaigns,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDomainDeliverabilityCampaignsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_domain_deliverability_campaigns, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDomainDeliverabilityCampaigns(ctx, input)
				}
				var results []*svc.ListDomainDeliverabilityCampaignsOutput
				p := svc.NewListDomainDeliverabilityCampaignsPaginator(client, input)
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
		"list-email-identities": {
			Name:   "list-email-identities",
			Fields: fields_list_email_identities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEmailIdentitiesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_email_identities, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEmailIdentities(ctx, input)
				}
				var results []*svc.ListEmailIdentitiesOutput
				p := svc.NewListEmailIdentitiesPaginator(client, input)
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
		"put-account-dedicated-ip-warmup-attributes": {
			Name:   "put-account-dedicated-ip-warmup-attributes",
			Fields: fields_put_account_dedicated_ip_warmup_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutAccountDedicatedIpWarmupAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_account_dedicated_ip_warmup_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutAccountDedicatedIpWarmupAttributes(ctx, input)
			},
		},
		"put-account-sending-attributes": {
			Name:   "put-account-sending-attributes",
			Fields: fields_put_account_sending_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutAccountSendingAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_account_sending_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutAccountSendingAttributes(ctx, input)
			},
		},
		"put-configuration-set-delivery-options": {
			Name:   "put-configuration-set-delivery-options",
			Fields: fields_put_configuration_set_delivery_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutConfigurationSetDeliveryOptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_configuration_set_delivery_options, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutConfigurationSetDeliveryOptions(ctx, input)
			},
		},
		"put-configuration-set-reputation-options": {
			Name:   "put-configuration-set-reputation-options",
			Fields: fields_put_configuration_set_reputation_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutConfigurationSetReputationOptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_configuration_set_reputation_options, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutConfigurationSetReputationOptions(ctx, input)
			},
		},
		"put-configuration-set-sending-options": {
			Name:   "put-configuration-set-sending-options",
			Fields: fields_put_configuration_set_sending_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutConfigurationSetSendingOptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_configuration_set_sending_options, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutConfigurationSetSendingOptions(ctx, input)
			},
		},
		"put-configuration-set-tracking-options": {
			Name:   "put-configuration-set-tracking-options",
			Fields: fields_put_configuration_set_tracking_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutConfigurationSetTrackingOptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_configuration_set_tracking_options, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutConfigurationSetTrackingOptions(ctx, input)
			},
		},
		"put-dedicated-ip-in-pool": {
			Name:   "put-dedicated-ip-in-pool",
			Fields: fields_put_dedicated_ip_in_pool,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutDedicatedIpInPoolInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_dedicated_ip_in_pool, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutDedicatedIpInPool(ctx, input)
			},
		},
		"put-dedicated-ip-warmup-attributes": {
			Name:   "put-dedicated-ip-warmup-attributes",
			Fields: fields_put_dedicated_ip_warmup_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutDedicatedIpWarmupAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_dedicated_ip_warmup_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutDedicatedIpWarmupAttributes(ctx, input)
			},
		},
		"put-deliverability-dashboard-option": {
			Name:   "put-deliverability-dashboard-option",
			Fields: fields_put_deliverability_dashboard_option,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutDeliverabilityDashboardOptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_deliverability_dashboard_option, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutDeliverabilityDashboardOption(ctx, input)
			},
		},
		"put-email-identity-dkim-attributes": {
			Name:   "put-email-identity-dkim-attributes",
			Fields: fields_put_email_identity_dkim_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutEmailIdentityDkimAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_email_identity_dkim_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutEmailIdentityDkimAttributes(ctx, input)
			},
		},
		"put-email-identity-feedback-attributes": {
			Name:   "put-email-identity-feedback-attributes",
			Fields: fields_put_email_identity_feedback_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutEmailIdentityFeedbackAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_email_identity_feedback_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutEmailIdentityFeedbackAttributes(ctx, input)
			},
		},
		"put-email-identity-mail-from-attributes": {
			Name:   "put-email-identity-mail-from-attributes",
			Fields: fields_put_email_identity_mail_from_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutEmailIdentityMailFromAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_email_identity_mail_from_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutEmailIdentityMailFromAttributes(ctx, input)
			},
		},
		"send-email": {
			Name:   "send-email",
			Fields: fields_send_email,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendEmailInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_email, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendEmail(ctx, input)
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
		"update-configuration-set-event-destination": {
			Name:   "update-configuration-set-event-destination",
			Fields: fields_update_configuration_set_event_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateConfigurationSetEventDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_configuration_set_event_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateConfigurationSetEventDestination(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("pinpointemail", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
