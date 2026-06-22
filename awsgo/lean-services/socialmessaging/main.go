package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/socialmessaging"
)

var fields_associate_whats_app_business_account = []leanruntime.Field{
	{Name: "SetupFinalization", Flag: "setup-finalization", Type: "*types.WhatsAppSetupFinalization", Required: false},
	{Name: "SignupCallback", Flag: "signup-callback", Type: "*types.WhatsAppSignupCallback", Required: false},
}

var fields_create_whats_app_message_template = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "TemplateDefinition", Flag: "template-definition", Type: "[]byte", Required: true},
}

var fields_create_whats_app_message_template_from_library = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "MetaLibraryTemplate", Flag: "meta-library-template", Type: "*types.MetaLibraryTemplate", Required: true},
}

var fields_create_whats_app_message_template_media = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "SourceS3File", Flag: "source-s3-file", Type: "*types.S3File", Required: false},
}

var fields_delete_whats_app_message_media = []leanruntime.Field{
	{Name: "MediaId", Flag: "media-id", Type: "*string", Required: true},
	{Name: "OriginationPhoneNumberId", Flag: "origination-phone-number-id", Type: "*string", Required: true},
}

var fields_delete_whats_app_message_template = []leanruntime.Field{
	{Name: "DeleteAllLanguages", Flag: "delete-all-languages", Type: "*bool", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "MetaTemplateId", Flag: "meta-template-id", Type: "*string", Required: false},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
}

var fields_disassociate_whats_app_business_account = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_linked_whats_app_business_account = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_linked_whats_app_business_account_phone_number = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_whats_app_message_media = []leanruntime.Field{
	{Name: "DestinationS3File", Flag: "destination-s3-file", Type: "*types.S3File", Required: false},
	{Name: "DestinationS3PresignedUrl", Flag: "destination-s3-presigned-url", Type: "*types.S3PresignedUrl", Required: false},
	{Name: "MediaId", Flag: "media-id", Type: "*string", Required: true},
	{Name: "MetadataOnly", Flag: "metadata-only", Type: "*bool", Required: false},
	{Name: "OriginationPhoneNumberId", Flag: "origination-phone-number-id", Type: "*string", Required: true},
}

var fields_get_whats_app_message_template = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "MetaTemplateId", Flag: "meta-template-id", Type: "*string", Required: true},
}

var fields_list_linked_whats_app_business_accounts = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_whats_app_message_templates = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_whats_app_template_library = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "map[string]string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_post_whats_app_message_media = []leanruntime.Field{
	{Name: "OriginationPhoneNumberId", Flag: "origination-phone-number-id", Type: "*string", Required: true},
	{Name: "SourceS3File", Flag: "source-s3-file", Type: "*types.S3File", Required: false},
	{Name: "SourceS3PresignedUrl", Flag: "source-s3-presigned-url", Type: "*types.S3PresignedUrl", Required: false},
}

var fields_put_whats_app_business_account_event_destinations = []leanruntime.Field{
	{Name: "EventDestinations", Flag: "event-destinations", Type: "[]types.WhatsAppBusinessAccountEventDestination", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_send_whats_app_message = []leanruntime.Field{
	{Name: "Message", Flag: "message", Type: "[]byte", Required: true},
	{Name: "MetaApiVersion", Flag: "meta-api-version", Type: "*string", Required: true},
	{Name: "OriginationPhoneNumberId", Flag: "origination-phone-number-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_whats_app_message_template = []leanruntime.Field{
	{Name: "CtaUrlLinkTrackingOptedOut", Flag: "cta-url-link-tracking-opted-out", Type: "*bool", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "MetaTemplateId", Flag: "meta-template-id", Type: "*string", Required: true},
	{Name: "ParameterFormat", Flag: "parameter-format", Type: "*string", Required: false},
	{Name: "TemplateCategory", Flag: "template-category", Type: "*string", Required: false},
	{Name: "TemplateComponents", Flag: "template-components", Type: "[]byte", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-whats-app-business-account": {
			Name:   "associate-whats-app-business-account",
			Fields: fields_associate_whats_app_business_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateWhatsAppBusinessAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_whats_app_business_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateWhatsAppBusinessAccount(ctx, input)
			},
		},
		"create-whats-app-message-template": {
			Name:   "create-whats-app-message-template",
			Fields: fields_create_whats_app_message_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWhatsAppMessageTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_whats_app_message_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWhatsAppMessageTemplate(ctx, input)
			},
		},
		"create-whats-app-message-template-from-library": {
			Name:   "create-whats-app-message-template-from-library",
			Fields: fields_create_whats_app_message_template_from_library,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWhatsAppMessageTemplateFromLibraryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_whats_app_message_template_from_library, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWhatsAppMessageTemplateFromLibrary(ctx, input)
			},
		},
		"create-whats-app-message-template-media": {
			Name:   "create-whats-app-message-template-media",
			Fields: fields_create_whats_app_message_template_media,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWhatsAppMessageTemplateMediaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_whats_app_message_template_media, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWhatsAppMessageTemplateMedia(ctx, input)
			},
		},
		"delete-whats-app-message-media": {
			Name:   "delete-whats-app-message-media",
			Fields: fields_delete_whats_app_message_media,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWhatsAppMessageMediaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_whats_app_message_media, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWhatsAppMessageMedia(ctx, input)
			},
		},
		"delete-whats-app-message-template": {
			Name:   "delete-whats-app-message-template",
			Fields: fields_delete_whats_app_message_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWhatsAppMessageTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_whats_app_message_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWhatsAppMessageTemplate(ctx, input)
			},
		},
		"disassociate-whats-app-business-account": {
			Name:   "disassociate-whats-app-business-account",
			Fields: fields_disassociate_whats_app_business_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateWhatsAppBusinessAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_whats_app_business_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateWhatsAppBusinessAccount(ctx, input)
			},
		},
		"get-linked-whats-app-business-account": {
			Name:   "get-linked-whats-app-business-account",
			Fields: fields_get_linked_whats_app_business_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLinkedWhatsAppBusinessAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_linked_whats_app_business_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLinkedWhatsAppBusinessAccount(ctx, input)
			},
		},
		"get-linked-whats-app-business-account-phone-number": {
			Name:   "get-linked-whats-app-business-account-phone-number",
			Fields: fields_get_linked_whats_app_business_account_phone_number,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLinkedWhatsAppBusinessAccountPhoneNumberInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_linked_whats_app_business_account_phone_number, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLinkedWhatsAppBusinessAccountPhoneNumber(ctx, input)
			},
		},
		"get-whats-app-message-media": {
			Name:   "get-whats-app-message-media",
			Fields: fields_get_whats_app_message_media,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWhatsAppMessageMediaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_whats_app_message_media, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetWhatsAppMessageMedia(ctx, input)
			},
		},
		"get-whats-app-message-template": {
			Name:   "get-whats-app-message-template",
			Fields: fields_get_whats_app_message_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWhatsAppMessageTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_whats_app_message_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetWhatsAppMessageTemplate(ctx, input)
			},
		},
		"list-linked-whats-app-business-accounts": {
			Name:   "list-linked-whats-app-business-accounts",
			Fields: fields_list_linked_whats_app_business_accounts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLinkedWhatsAppBusinessAccountsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_linked_whats_app_business_accounts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLinkedWhatsAppBusinessAccounts(ctx, input)
				}
				var results []*svc.ListLinkedWhatsAppBusinessAccountsOutput
				p := svc.NewListLinkedWhatsAppBusinessAccountsPaginator(client, input)
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
		"list-whats-app-message-templates": {
			Name:   "list-whats-app-message-templates",
			Fields: fields_list_whats_app_message_templates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWhatsAppMessageTemplatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_whats_app_message_templates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWhatsAppMessageTemplates(ctx, input)
				}
				var results []*svc.ListWhatsAppMessageTemplatesOutput
				p := svc.NewListWhatsAppMessageTemplatesPaginator(client, input)
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
		"list-whats-app-template-library": {
			Name:   "list-whats-app-template-library",
			Fields: fields_list_whats_app_template_library,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWhatsAppTemplateLibraryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_whats_app_template_library, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWhatsAppTemplateLibrary(ctx, input)
				}
				var results []*svc.ListWhatsAppTemplateLibraryOutput
				p := svc.NewListWhatsAppTemplateLibraryPaginator(client, input)
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
		"post-whats-app-message-media": {
			Name:   "post-whats-app-message-media",
			Fields: fields_post_whats_app_message_media,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PostWhatsAppMessageMediaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_post_whats_app_message_media, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PostWhatsAppMessageMedia(ctx, input)
			},
		},
		"put-whats-app-business-account-event-destinations": {
			Name:   "put-whats-app-business-account-event-destinations",
			Fields: fields_put_whats_app_business_account_event_destinations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutWhatsAppBusinessAccountEventDestinationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_whats_app_business_account_event_destinations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutWhatsAppBusinessAccountEventDestinations(ctx, input)
			},
		},
		"send-whats-app-message": {
			Name:   "send-whats-app-message",
			Fields: fields_send_whats_app_message,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendWhatsAppMessageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_whats_app_message, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendWhatsAppMessage(ctx, input)
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
		"update-whats-app-message-template": {
			Name:   "update-whats-app-message-template",
			Fields: fields_update_whats_app_message_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateWhatsAppMessageTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_whats_app_message_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateWhatsAppMessageTemplate(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("socialmessaging", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
