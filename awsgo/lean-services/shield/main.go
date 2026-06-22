package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/shield"
)

var fields_associate_drt_log_bucket = []leanruntime.Field{
	{Name: "LogBucket", Flag: "log-bucket", Type: "*string", Required: true},
}

var fields_associate_drt_role = []leanruntime.Field{
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
}

var fields_associate_health_check = []leanruntime.Field{
	{Name: "HealthCheckArn", Flag: "health-check-arn", Type: "*string", Required: true},
	{Name: "ProtectionId", Flag: "protection-id", Type: "*string", Required: true},
}

var fields_associate_proactive_engagement_details = []leanruntime.Field{
	{Name: "EmergencyContactList", Flag: "emergency-contact-list", Type: "[]types.EmergencyContact", Required: true},
}

var fields_create_protection = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_protection_group = []leanruntime.Field{
	{Name: "Aggregation", Flag: "aggregation", Type: "types.ProtectionGroupAggregation", Required: true},
	{Name: "Members", Flag: "members", Type: "[]string", Required: false},
	{Name: "Pattern", Flag: "pattern", Type: "types.ProtectionGroupPattern", Required: true},
	{Name: "ProtectionGroupId", Flag: "protection-group-id", Type: "*string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.ProtectedResourceType", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_subscription = []leanruntime.Field{}

var fields_delete_protection = []leanruntime.Field{
	{Name: "ProtectionId", Flag: "protection-id", Type: "*string", Required: true},
}

var fields_delete_protection_group = []leanruntime.Field{
	{Name: "ProtectionGroupId", Flag: "protection-group-id", Type: "*string", Required: true},
}

var fields_delete_subscription = []leanruntime.Field{}

var fields_describe_attack = []leanruntime.Field{
	{Name: "AttackId", Flag: "attack-id", Type: "*string", Required: true},
}

var fields_describe_attack_statistics = []leanruntime.Field{}

var fields_describe_drt_access = []leanruntime.Field{}

var fields_describe_emergency_contact_settings = []leanruntime.Field{}

var fields_describe_protection = []leanruntime.Field{
	{Name: "ProtectionId", Flag: "protection-id", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: false},
}

var fields_describe_protection_group = []leanruntime.Field{
	{Name: "ProtectionGroupId", Flag: "protection-group-id", Type: "*string", Required: true},
}

var fields_describe_subscription = []leanruntime.Field{}

var fields_disable_application_layer_automatic_response = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_disable_proactive_engagement = []leanruntime.Field{}

var fields_disassociate_drt_log_bucket = []leanruntime.Field{
	{Name: "LogBucket", Flag: "log-bucket", Type: "*string", Required: true},
}

var fields_disassociate_drt_role = []leanruntime.Field{}

var fields_disassociate_health_check = []leanruntime.Field{
	{Name: "HealthCheckArn", Flag: "health-check-arn", Type: "*string", Required: true},
	{Name: "ProtectionId", Flag: "protection-id", Type: "*string", Required: true},
}

var fields_enable_application_layer_automatic_response = []leanruntime.Field{
	{Name: "Action", Flag: "action", Type: "*types.ResponseAction", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_enable_proactive_engagement = []leanruntime.Field{}

var fields_get_subscription_state = []leanruntime.Field{}

var fields_list_attacks = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*types.TimeRange", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceArns", Flag: "resource-arns", Type: "[]string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*types.TimeRange", Required: false},
}

var fields_list_protection_groups = []leanruntime.Field{
	{Name: "InclusionFilters", Flag: "inclusion-filters", Type: "*types.InclusionProtectionGroupFilters", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_protections = []leanruntime.Field{
	{Name: "InclusionFilters", Flag: "inclusion-filters", Type: "*types.InclusionProtectionFilters", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_resources_in_protection_group = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProtectionGroupId", Flag: "protection-group-id", Type: "*string", Required: true},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_application_layer_automatic_response = []leanruntime.Field{
	{Name: "Action", Flag: "action", Type: "*types.ResponseAction", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_update_emergency_contact_settings = []leanruntime.Field{
	{Name: "EmergencyContactList", Flag: "emergency-contact-list", Type: "[]types.EmergencyContact", Required: false},
}

var fields_update_protection_group = []leanruntime.Field{
	{Name: "Aggregation", Flag: "aggregation", Type: "types.ProtectionGroupAggregation", Required: true},
	{Name: "Members", Flag: "members", Type: "[]string", Required: false},
	{Name: "Pattern", Flag: "pattern", Type: "types.ProtectionGroupPattern", Required: true},
	{Name: "ProtectionGroupId", Flag: "protection-group-id", Type: "*string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.ProtectedResourceType", Required: false},
}

var fields_update_subscription = []leanruntime.Field{
	{Name: "AutoRenew", Flag: "auto-renew", Type: "types.AutoRenew", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-drt-log-bucket": {
			Name:   "associate-drt-log-bucket",
			Fields: fields_associate_drt_log_bucket,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateDRTLogBucketInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_drt_log_bucket, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateDRTLogBucket(ctx, input)
			},
		},
		"associate-drt-role": {
			Name:   "associate-drt-role",
			Fields: fields_associate_drt_role,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateDRTRoleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_drt_role, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateDRTRole(ctx, input)
			},
		},
		"associate-health-check": {
			Name:   "associate-health-check",
			Fields: fields_associate_health_check,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateHealthCheckInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_health_check, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateHealthCheck(ctx, input)
			},
		},
		"associate-proactive-engagement-details": {
			Name:   "associate-proactive-engagement-details",
			Fields: fields_associate_proactive_engagement_details,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateProactiveEngagementDetailsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_proactive_engagement_details, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateProactiveEngagementDetails(ctx, input)
			},
		},
		"create-protection": {
			Name:   "create-protection",
			Fields: fields_create_protection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateProtectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_protection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateProtection(ctx, input)
			},
		},
		"create-protection-group": {
			Name:   "create-protection-group",
			Fields: fields_create_protection_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateProtectionGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_protection_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateProtectionGroup(ctx, input)
			},
		},
		"create-subscription": {
			Name:   "create-subscription",
			Fields: fields_create_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSubscription(ctx, input)
			},
		},
		"delete-protection": {
			Name:   "delete-protection",
			Fields: fields_delete_protection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteProtectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_protection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteProtection(ctx, input)
			},
		},
		"delete-protection-group": {
			Name:   "delete-protection-group",
			Fields: fields_delete_protection_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteProtectionGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_protection_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteProtectionGroup(ctx, input)
			},
		},
		"delete-subscription": {
			Name:   "delete-subscription",
			Fields: fields_delete_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSubscription(ctx, input)
			},
		},
		"describe-attack": {
			Name:   "describe-attack",
			Fields: fields_describe_attack,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAttackInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_attack, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAttack(ctx, input)
			},
		},
		"describe-attack-statistics": {
			Name:   "describe-attack-statistics",
			Fields: fields_describe_attack_statistics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAttackStatisticsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_attack_statistics, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAttackStatistics(ctx, input)
			},
		},
		"describe-drt-access": {
			Name:   "describe-drt-access",
			Fields: fields_describe_drt_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDRTAccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_drt_access, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDRTAccess(ctx, input)
			},
		},
		"describe-emergency-contact-settings": {
			Name:   "describe-emergency-contact-settings",
			Fields: fields_describe_emergency_contact_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEmergencyContactSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_emergency_contact_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEmergencyContactSettings(ctx, input)
			},
		},
		"describe-protection": {
			Name:   "describe-protection",
			Fields: fields_describe_protection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeProtectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_protection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeProtection(ctx, input)
			},
		},
		"describe-protection-group": {
			Name:   "describe-protection-group",
			Fields: fields_describe_protection_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeProtectionGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_protection_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeProtectionGroup(ctx, input)
			},
		},
		"describe-subscription": {
			Name:   "describe-subscription",
			Fields: fields_describe_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSubscription(ctx, input)
			},
		},
		"disable-application-layer-automatic-response": {
			Name:   "disable-application-layer-automatic-response",
			Fields: fields_disable_application_layer_automatic_response,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableApplicationLayerAutomaticResponseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_application_layer_automatic_response, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableApplicationLayerAutomaticResponse(ctx, input)
			},
		},
		"disable-proactive-engagement": {
			Name:   "disable-proactive-engagement",
			Fields: fields_disable_proactive_engagement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableProactiveEngagementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_proactive_engagement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableProactiveEngagement(ctx, input)
			},
		},
		"disassociate-drt-log-bucket": {
			Name:   "disassociate-drt-log-bucket",
			Fields: fields_disassociate_drt_log_bucket,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateDRTLogBucketInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_drt_log_bucket, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateDRTLogBucket(ctx, input)
			},
		},
		"disassociate-drt-role": {
			Name:   "disassociate-drt-role",
			Fields: fields_disassociate_drt_role,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateDRTRoleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_drt_role, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateDRTRole(ctx, input)
			},
		},
		"disassociate-health-check": {
			Name:   "disassociate-health-check",
			Fields: fields_disassociate_health_check,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateHealthCheckInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_health_check, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateHealthCheck(ctx, input)
			},
		},
		"enable-application-layer-automatic-response": {
			Name:   "enable-application-layer-automatic-response",
			Fields: fields_enable_application_layer_automatic_response,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableApplicationLayerAutomaticResponseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_application_layer_automatic_response, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableApplicationLayerAutomaticResponse(ctx, input)
			},
		},
		"enable-proactive-engagement": {
			Name:   "enable-proactive-engagement",
			Fields: fields_enable_proactive_engagement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableProactiveEngagementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_proactive_engagement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableProactiveEngagement(ctx, input)
			},
		},
		"get-subscription-state": {
			Name:   "get-subscription-state",
			Fields: fields_get_subscription_state,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSubscriptionStateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_subscription_state, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSubscriptionState(ctx, input)
			},
		},
		"list-attacks": {
			Name:   "list-attacks",
			Fields: fields_list_attacks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAttacksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_attacks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAttacks(ctx, input)
				}
				var results []*svc.ListAttacksOutput
				p := svc.NewListAttacksPaginator(client, input)
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
		"list-protection-groups": {
			Name:   "list-protection-groups",
			Fields: fields_list_protection_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProtectionGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_protection_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProtectionGroups(ctx, input)
				}
				var results []*svc.ListProtectionGroupsOutput
				p := svc.NewListProtectionGroupsPaginator(client, input)
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
		"list-protections": {
			Name:   "list-protections",
			Fields: fields_list_protections,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProtectionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_protections, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProtections(ctx, input)
				}
				var results []*svc.ListProtectionsOutput
				p := svc.NewListProtectionsPaginator(client, input)
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
		"list-resources-in-protection-group": {
			Name:   "list-resources-in-protection-group",
			Fields: fields_list_resources_in_protection_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResourcesInProtectionGroupInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_resources_in_protection_group, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResourcesInProtectionGroup(ctx, input)
				}
				var results []*svc.ListResourcesInProtectionGroupOutput
				p := svc.NewListResourcesInProtectionGroupPaginator(client, input)
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
		"update-application-layer-automatic-response": {
			Name:   "update-application-layer-automatic-response",
			Fields: fields_update_application_layer_automatic_response,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateApplicationLayerAutomaticResponseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_application_layer_automatic_response, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateApplicationLayerAutomaticResponse(ctx, input)
			},
		},
		"update-emergency-contact-settings": {
			Name:   "update-emergency-contact-settings",
			Fields: fields_update_emergency_contact_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEmergencyContactSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_emergency_contact_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEmergencyContactSettings(ctx, input)
			},
		},
		"update-protection-group": {
			Name:   "update-protection-group",
			Fields: fields_update_protection_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateProtectionGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_protection_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateProtectionGroup(ctx, input)
			},
		},
		"update-subscription": {
			Name:   "update-subscription",
			Fields: fields_update_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSubscription(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("shield", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
