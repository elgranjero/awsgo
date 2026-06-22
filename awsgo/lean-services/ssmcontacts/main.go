package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/ssmcontacts"
)

var fields_accept_page = []leanruntime.Field{
	{Name: "AcceptCode", Flag: "accept-code", Type: "*string", Required: true},
	{Name: "AcceptCodeValidation", Flag: "accept-code-validation", Type: "types.AcceptCodeValidation", Required: false},
	{Name: "AcceptType", Flag: "accept-type", Type: "types.AcceptType", Required: true},
	{Name: "ContactChannelId", Flag: "contact-channel-id", Type: "*string", Required: false},
	{Name: "Note", Flag: "note", Type: "*string", Required: false},
	{Name: "PageId", Flag: "page-id", Type: "*string", Required: true},
}

var fields_activate_contact_channel = []leanruntime.Field{
	{Name: "ActivationCode", Flag: "activation-code", Type: "*string", Required: true},
	{Name: "ContactChannelId", Flag: "contact-channel-id", Type: "*string", Required: true},
}

var fields_create_contact = []leanruntime.Field{
	{Name: "Alias", Flag: "alias", Type: "*string", Required: true},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: false},
	{Name: "Plan", Flag: "plan", Type: "*types.Plan", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Type", Flag: "type", Type: "types.ContactType", Required: true},
}

var fields_create_contact_channel = []leanruntime.Field{
	{Name: "ContactId", Flag: "contact-id", Type: "*string", Required: true},
	{Name: "DeferActivation", Flag: "defer-activation", Type: "*bool", Required: false},
	{Name: "DeliveryAddress", Flag: "delivery-address", Type: "*types.ContactChannelAddress", Required: true},
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.ChannelType", Required: true},
}

var fields_create_rotation = []leanruntime.Field{
	{Name: "ContactIds", Flag: "contact-ids", Type: "[]string", Required: true},
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Recurrence", Flag: "recurrence", Type: "*types.RecurrenceSettings", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TimeZoneId", Flag: "time-zone-id", Type: "*string", Required: true},
}

var fields_create_rotation_override = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: false},
	{Name: "NewContactIds", Flag: "new-contact-ids", Type: "[]string", Required: true},
	{Name: "RotationId", Flag: "rotation-id", Type: "*string", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
}

var fields_deactivate_contact_channel = []leanruntime.Field{
	{Name: "ContactChannelId", Flag: "contact-channel-id", Type: "*string", Required: true},
}

var fields_delete_contact = []leanruntime.Field{
	{Name: "ContactId", Flag: "contact-id", Type: "*string", Required: true},
}

var fields_delete_contact_channel = []leanruntime.Field{
	{Name: "ContactChannelId", Flag: "contact-channel-id", Type: "*string", Required: true},
}

var fields_delete_rotation = []leanruntime.Field{
	{Name: "RotationId", Flag: "rotation-id", Type: "*string", Required: true},
}

var fields_delete_rotation_override = []leanruntime.Field{
	{Name: "RotationId", Flag: "rotation-id", Type: "*string", Required: true},
	{Name: "RotationOverrideId", Flag: "rotation-override-id", Type: "*string", Required: true},
}

var fields_describe_engagement = []leanruntime.Field{
	{Name: "EngagementId", Flag: "engagement-id", Type: "*string", Required: true},
}

var fields_describe_page = []leanruntime.Field{
	{Name: "PageId", Flag: "page-id", Type: "*string", Required: true},
}

var fields_get_contact = []leanruntime.Field{
	{Name: "ContactId", Flag: "contact-id", Type: "*string", Required: true},
}

var fields_get_contact_channel = []leanruntime.Field{
	{Name: "ContactChannelId", Flag: "contact-channel-id", Type: "*string", Required: true},
}

var fields_get_contact_policy = []leanruntime.Field{
	{Name: "ContactArn", Flag: "contact-arn", Type: "*string", Required: true},
}

var fields_get_rotation = []leanruntime.Field{
	{Name: "RotationId", Flag: "rotation-id", Type: "*string", Required: true},
}

var fields_get_rotation_override = []leanruntime.Field{
	{Name: "RotationId", Flag: "rotation-id", Type: "*string", Required: true},
	{Name: "RotationOverrideId", Flag: "rotation-override-id", Type: "*string", Required: true},
}

var fields_list_contact_channels = []leanruntime.Field{
	{Name: "ContactId", Flag: "contact-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_contacts = []leanruntime.Field{
	{Name: "AliasPrefix", Flag: "alias-prefix", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.ContactType", Required: false},
}

var fields_list_engagements = []leanruntime.Field{
	{Name: "IncidentId", Flag: "incident-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TimeRangeValue", Flag: "time-range-value", Type: "*types.TimeRange", Required: false},
}

var fields_list_page_receipts = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageId", Flag: "page-id", Type: "*string", Required: true},
}

var fields_list_page_resolutions = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageId", Flag: "page-id", Type: "*string", Required: true},
}

var fields_list_pages_by_contact = []leanruntime.Field{
	{Name: "ContactId", Flag: "contact-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_pages_by_engagement = []leanruntime.Field{
	{Name: "EngagementId", Flag: "engagement-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_preview_rotation_shifts = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Members", Flag: "members", Type: "[]string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Overrides", Flag: "overrides", Type: "[]types.PreviewOverride", Required: false},
	{Name: "Recurrence", Flag: "recurrence", Type: "*types.RecurrenceSettings", Required: true},
	{Name: "RotationStartTime", Flag: "rotation-start-time", Type: "*time.Time", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
	{Name: "TimeZoneId", Flag: "time-zone-id", Type: "*string", Required: true},
}

var fields_list_rotation_overrides = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RotationId", Flag: "rotation-id", Type: "*string", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
}

var fields_list_rotation_shifts = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RotationId", Flag: "rotation-id", Type: "*string", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
}

var fields_list_rotations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RotationNamePrefix", Flag: "rotation-name-prefix", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_put_contact_policy = []leanruntime.Field{
	{Name: "ContactArn", Flag: "contact-arn", Type: "*string", Required: true},
	{Name: "Policy", Flag: "policy", Type: "*string", Required: true},
}

var fields_send_activation_code = []leanruntime.Field{
	{Name: "ContactChannelId", Flag: "contact-channel-id", Type: "*string", Required: true},
}

var fields_start_engagement = []leanruntime.Field{
	{Name: "ContactId", Flag: "contact-id", Type: "*string", Required: true},
	{Name: "Content", Flag: "content", Type: "*string", Required: true},
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: false},
	{Name: "IncidentId", Flag: "incident-id", Type: "*string", Required: false},
	{Name: "PublicContent", Flag: "public-content", Type: "*string", Required: false},
	{Name: "PublicSubject", Flag: "public-subject", Type: "*string", Required: false},
	{Name: "Sender", Flag: "sender", Type: "*string", Required: true},
	{Name: "Subject", Flag: "subject", Type: "*string", Required: true},
}

var fields_stop_engagement = []leanruntime.Field{
	{Name: "EngagementId", Flag: "engagement-id", Type: "*string", Required: true},
	{Name: "Reason", Flag: "reason", Type: "*string", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_contact = []leanruntime.Field{
	{Name: "ContactId", Flag: "contact-id", Type: "*string", Required: true},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "Plan", Flag: "plan", Type: "*types.Plan", Required: false},
}

var fields_update_contact_channel = []leanruntime.Field{
	{Name: "ContactChannelId", Flag: "contact-channel-id", Type: "*string", Required: true},
	{Name: "DeliveryAddress", Flag: "delivery-address", Type: "*types.ContactChannelAddress", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_rotation = []leanruntime.Field{
	{Name: "ContactIds", Flag: "contact-ids", Type: "[]string", Required: false},
	{Name: "Recurrence", Flag: "recurrence", Type: "*types.RecurrenceSettings", Required: true},
	{Name: "RotationId", Flag: "rotation-id", Type: "*string", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
	{Name: "TimeZoneId", Flag: "time-zone-id", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"accept-page": {
			Name:   "accept-page",
			Fields: fields_accept_page,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AcceptPageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_accept_page, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AcceptPage(ctx, input)
			},
		},
		"activate-contact-channel": {
			Name:   "activate-contact-channel",
			Fields: fields_activate_contact_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ActivateContactChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_activate_contact_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ActivateContactChannel(ctx, input)
			},
		},
		"create-contact": {
			Name:   "create-contact",
			Fields: fields_create_contact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateContactInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_contact, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateContact(ctx, input)
			},
		},
		"create-contact-channel": {
			Name:   "create-contact-channel",
			Fields: fields_create_contact_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateContactChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_contact_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateContactChannel(ctx, input)
			},
		},
		"create-rotation": {
			Name:   "create-rotation",
			Fields: fields_create_rotation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRotationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_rotation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRotation(ctx, input)
			},
		},
		"create-rotation-override": {
			Name:   "create-rotation-override",
			Fields: fields_create_rotation_override,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRotationOverrideInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_rotation_override, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRotationOverride(ctx, input)
			},
		},
		"deactivate-contact-channel": {
			Name:   "deactivate-contact-channel",
			Fields: fields_deactivate_contact_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeactivateContactChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deactivate_contact_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeactivateContactChannel(ctx, input)
			},
		},
		"delete-contact": {
			Name:   "delete-contact",
			Fields: fields_delete_contact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteContactInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_contact, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteContact(ctx, input)
			},
		},
		"delete-contact-channel": {
			Name:   "delete-contact-channel",
			Fields: fields_delete_contact_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteContactChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_contact_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteContactChannel(ctx, input)
			},
		},
		"delete-rotation": {
			Name:   "delete-rotation",
			Fields: fields_delete_rotation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRotationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_rotation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRotation(ctx, input)
			},
		},
		"delete-rotation-override": {
			Name:   "delete-rotation-override",
			Fields: fields_delete_rotation_override,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRotationOverrideInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_rotation_override, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRotationOverride(ctx, input)
			},
		},
		"describe-engagement": {
			Name:   "describe-engagement",
			Fields: fields_describe_engagement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEngagementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_engagement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEngagement(ctx, input)
			},
		},
		"describe-page": {
			Name:   "describe-page",
			Fields: fields_describe_page,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_page, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribePage(ctx, input)
			},
		},
		"get-contact": {
			Name:   "get-contact",
			Fields: fields_get_contact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetContactInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_contact, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetContact(ctx, input)
			},
		},
		"get-contact-channel": {
			Name:   "get-contact-channel",
			Fields: fields_get_contact_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetContactChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_contact_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetContactChannel(ctx, input)
			},
		},
		"get-contact-policy": {
			Name:   "get-contact-policy",
			Fields: fields_get_contact_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetContactPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_contact_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetContactPolicy(ctx, input)
			},
		},
		"get-rotation": {
			Name:   "get-rotation",
			Fields: fields_get_rotation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRotationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_rotation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRotation(ctx, input)
			},
		},
		"get-rotation-override": {
			Name:   "get-rotation-override",
			Fields: fields_get_rotation_override,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRotationOverrideInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_rotation_override, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRotationOverride(ctx, input)
			},
		},
		"list-contact-channels": {
			Name:   "list-contact-channels",
			Fields: fields_list_contact_channels,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListContactChannelsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_contact_channels, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListContactChannels(ctx, input)
				}
				var results []*svc.ListContactChannelsOutput
				p := svc.NewListContactChannelsPaginator(client, input)
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
		"list-contacts": {
			Name:   "list-contacts",
			Fields: fields_list_contacts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListContactsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_contacts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListContacts(ctx, input)
				}
				var results []*svc.ListContactsOutput
				p := svc.NewListContactsPaginator(client, input)
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
		"list-engagements": {
			Name:   "list-engagements",
			Fields: fields_list_engagements,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEngagementsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_engagements, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEngagements(ctx, input)
				}
				var results []*svc.ListEngagementsOutput
				p := svc.NewListEngagementsPaginator(client, input)
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
		"list-page-receipts": {
			Name:   "list-page-receipts",
			Fields: fields_list_page_receipts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPageReceiptsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_page_receipts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPageReceipts(ctx, input)
				}
				var results []*svc.ListPageReceiptsOutput
				p := svc.NewListPageReceiptsPaginator(client, input)
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
		"list-page-resolutions": {
			Name:   "list-page-resolutions",
			Fields: fields_list_page_resolutions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPageResolutionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_page_resolutions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPageResolutions(ctx, input)
				}
				var results []*svc.ListPageResolutionsOutput
				p := svc.NewListPageResolutionsPaginator(client, input)
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
		"list-pages-by-contact": {
			Name:   "list-pages-by-contact",
			Fields: fields_list_pages_by_contact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPagesByContactInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_pages_by_contact, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPagesByContact(ctx, input)
				}
				var results []*svc.ListPagesByContactOutput
				p := svc.NewListPagesByContactPaginator(client, input)
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
		"list-pages-by-engagement": {
			Name:   "list-pages-by-engagement",
			Fields: fields_list_pages_by_engagement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPagesByEngagementInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_pages_by_engagement, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPagesByEngagement(ctx, input)
				}
				var results []*svc.ListPagesByEngagementOutput
				p := svc.NewListPagesByEngagementPaginator(client, input)
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
		"list-preview-rotation-shifts": {
			Name:   "list-preview-rotation-shifts",
			Fields: fields_list_preview_rotation_shifts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPreviewRotationShiftsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_preview_rotation_shifts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPreviewRotationShifts(ctx, input)
				}
				var results []*svc.ListPreviewRotationShiftsOutput
				p := svc.NewListPreviewRotationShiftsPaginator(client, input)
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
		"list-rotation-overrides": {
			Name:   "list-rotation-overrides",
			Fields: fields_list_rotation_overrides,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRotationOverridesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_rotation_overrides, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRotationOverrides(ctx, input)
				}
				var results []*svc.ListRotationOverridesOutput
				p := svc.NewListRotationOverridesPaginator(client, input)
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
		"list-rotation-shifts": {
			Name:   "list-rotation-shifts",
			Fields: fields_list_rotation_shifts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRotationShiftsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_rotation_shifts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRotationShifts(ctx, input)
				}
				var results []*svc.ListRotationShiftsOutput
				p := svc.NewListRotationShiftsPaginator(client, input)
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
		"list-rotations": {
			Name:   "list-rotations",
			Fields: fields_list_rotations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRotationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_rotations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRotations(ctx, input)
				}
				var results []*svc.ListRotationsOutput
				p := svc.NewListRotationsPaginator(client, input)
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
		"put-contact-policy": {
			Name:   "put-contact-policy",
			Fields: fields_put_contact_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutContactPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_contact_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutContactPolicy(ctx, input)
			},
		},
		"send-activation-code": {
			Name:   "send-activation-code",
			Fields: fields_send_activation_code,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendActivationCodeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_activation_code, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendActivationCode(ctx, input)
			},
		},
		"start-engagement": {
			Name:   "start-engagement",
			Fields: fields_start_engagement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartEngagementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_engagement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartEngagement(ctx, input)
			},
		},
		"stop-engagement": {
			Name:   "stop-engagement",
			Fields: fields_stop_engagement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopEngagementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_engagement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopEngagement(ctx, input)
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
		"update-contact": {
			Name:   "update-contact",
			Fields: fields_update_contact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateContactInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_contact, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateContact(ctx, input)
			},
		},
		"update-contact-channel": {
			Name:   "update-contact-channel",
			Fields: fields_update_contact_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateContactChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_contact_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateContactChannel(ctx, input)
			},
		},
		"update-rotation": {
			Name:   "update-rotation",
			Fields: fields_update_rotation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRotationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_rotation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRotation(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("ssmcontacts", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
