package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/sns"
)

var fields_add_permission = []leanruntime.Field{
	{Name: "AWSAccountId", Flag: "aws-account-id", Type: "[]string", Required: true},
	{Name: "ActionName", Flag: "action-name", Type: "[]string", Required: true},
	{Name: "Label", Flag: "label", Type: "*string", Required: true},
	{Name: "TopicArn", Flag: "topic-arn", Type: "*string", Required: true},
}

var fields_check_if_phone_number_is_opted_out = []leanruntime.Field{
	{Name: "PhoneNumber", Flag: "phone-number", Type: "*string", Required: true},
}

var fields_confirm_subscription = []leanruntime.Field{
	{Name: "AuthenticateOnUnsubscribe", Flag: "authenticate-on-unsubscribe", Type: "*string", Required: false},
	{Name: "Token", Flag: "token", Type: "*string", Required: true},
	{Name: "TopicArn", Flag: "topic-arn", Type: "*string", Required: true},
}

var fields_create_platform_application = []leanruntime.Field{
	{Name: "Attributes", Flag: "attributes", Type: "map[string]string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Platform", Flag: "platform", Type: "*string", Required: true},
}

var fields_create_platform_endpoint = []leanruntime.Field{
	{Name: "Attributes", Flag: "attributes", Type: "map[string]string", Required: false},
	{Name: "CustomUserData", Flag: "custom-user-data", Type: "*string", Required: false},
	{Name: "PlatformApplicationArn", Flag: "platform-application-arn", Type: "*string", Required: true},
	{Name: "Token", Flag: "token", Type: "*string", Required: true},
}

var fields_create_sms_sandbox_phone_number = []leanruntime.Field{
	{Name: "LanguageCode", Flag: "language-code", Type: "types.LanguageCodeString", Required: false},
	{Name: "PhoneNumber", Flag: "phone-number", Type: "*string", Required: true},
}

var fields_create_topic = []leanruntime.Field{
	{Name: "Attributes", Flag: "attributes", Type: "map[string]string", Required: false},
	{Name: "DataProtectionPolicy", Flag: "data-protection-policy", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_delete_endpoint = []leanruntime.Field{
	{Name: "EndpointArn", Flag: "endpoint-arn", Type: "*string", Required: true},
}

var fields_delete_platform_application = []leanruntime.Field{
	{Name: "PlatformApplicationArn", Flag: "platform-application-arn", Type: "*string", Required: true},
}

var fields_delete_sms_sandbox_phone_number = []leanruntime.Field{
	{Name: "PhoneNumber", Flag: "phone-number", Type: "*string", Required: true},
}

var fields_delete_topic = []leanruntime.Field{
	{Name: "TopicArn", Flag: "topic-arn", Type: "*string", Required: true},
}

var fields_get_data_protection_policy = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_get_endpoint_attributes = []leanruntime.Field{
	{Name: "EndpointArn", Flag: "endpoint-arn", Type: "*string", Required: true},
}

var fields_get_platform_application_attributes = []leanruntime.Field{
	{Name: "PlatformApplicationArn", Flag: "platform-application-arn", Type: "*string", Required: true},
}

var fields_get_sms_attributes = []leanruntime.Field{
	{Name: "Attributes", Flag: "attributes", Type: "[]string", Required: false},
}

var fields_get_sms_sandbox_account_status = []leanruntime.Field{}

var fields_get_subscription_attributes = []leanruntime.Field{
	{Name: "SubscriptionArn", Flag: "subscription-arn", Type: "*string", Required: true},
}

var fields_get_topic_attributes = []leanruntime.Field{
	{Name: "TopicArn", Flag: "topic-arn", Type: "*string", Required: true},
}

var fields_list_endpoints_by_platform_application = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PlatformApplicationArn", Flag: "platform-application-arn", Type: "*string", Required: true},
}

var fields_list_origination_numbers = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_phone_numbers_opted_out = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_platform_applications = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_sms_sandbox_phone_numbers = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_subscriptions = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_subscriptions_by_topic = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TopicArn", Flag: "topic-arn", Type: "*string", Required: true},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_topics = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_opt_in_phone_number = []leanruntime.Field{
	{Name: "PhoneNumber", Flag: "phone-number", Type: "*string", Required: true},
}

var fields_publish = []leanruntime.Field{
	{Name: "Message", Flag: "message", Type: "*string", Required: true},
	{Name: "MessageAttributes", Flag: "message-attributes", Type: "map[string]types.MessageAttributeValue", Required: false},
	{Name: "MessageDeduplicationId", Flag: "message-deduplication-id", Type: "*string", Required: false},
	{Name: "MessageGroupId", Flag: "message-group-id", Type: "*string", Required: false},
	{Name: "MessageStructure", Flag: "message-structure", Type: "*string", Required: false},
	{Name: "PhoneNumber", Flag: "phone-number", Type: "*string", Required: false},
	{Name: "Subject", Flag: "subject", Type: "*string", Required: false},
	{Name: "TargetArn", Flag: "target-arn", Type: "*string", Required: false},
	{Name: "TopicArn", Flag: "topic-arn", Type: "*string", Required: false},
}

var fields_publish_batch = []leanruntime.Field{
	{Name: "PublishBatchRequestEntries", Flag: "publish-batch-request-entries", Type: "[]types.PublishBatchRequestEntry", Required: true},
	{Name: "TopicArn", Flag: "topic-arn", Type: "*string", Required: true},
}

var fields_put_data_protection_policy = []leanruntime.Field{
	{Name: "DataProtectionPolicy", Flag: "data-protection-policy", Type: "*string", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_remove_permission = []leanruntime.Field{
	{Name: "Label", Flag: "label", Type: "*string", Required: true},
	{Name: "TopicArn", Flag: "topic-arn", Type: "*string", Required: true},
}

var fields_set_endpoint_attributes = []leanruntime.Field{
	{Name: "Attributes", Flag: "attributes", Type: "map[string]string", Required: true},
	{Name: "EndpointArn", Flag: "endpoint-arn", Type: "*string", Required: true},
}

var fields_set_platform_application_attributes = []leanruntime.Field{
	{Name: "Attributes", Flag: "attributes", Type: "map[string]string", Required: true},
	{Name: "PlatformApplicationArn", Flag: "platform-application-arn", Type: "*string", Required: true},
}

var fields_set_sms_attributes = []leanruntime.Field{
	{Name: "Attributes", Flag: "attributes", Type: "map[string]string", Required: true},
}

var fields_set_subscription_attributes = []leanruntime.Field{
	{Name: "AttributeName", Flag: "attribute-name", Type: "*string", Required: true},
	{Name: "AttributeValue", Flag: "attribute-value", Type: "*string", Required: false},
	{Name: "SubscriptionArn", Flag: "subscription-arn", Type: "*string", Required: true},
}

var fields_set_topic_attributes = []leanruntime.Field{
	{Name: "AttributeName", Flag: "attribute-name", Type: "*string", Required: true},
	{Name: "AttributeValue", Flag: "attribute-value", Type: "*string", Required: false},
	{Name: "TopicArn", Flag: "topic-arn", Type: "*string", Required: true},
}

var fields_subscribe = []leanruntime.Field{
	{Name: "Attributes", Flag: "attributes", Type: "map[string]string", Required: false},
	{Name: "Endpoint", Flag: "endpoint", Type: "*string", Required: false},
	{Name: "Protocol", Flag: "protocol", Type: "*string", Required: true},
	{Name: "ReturnSubscriptionArn", Flag: "return-subscription-arn", Type: "bool", Required: false},
	{Name: "TopicArn", Flag: "topic-arn", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_unsubscribe = []leanruntime.Field{
	{Name: "SubscriptionArn", Flag: "subscription-arn", Type: "*string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_verify_sms_sandbox_phone_number = []leanruntime.Field{
	{Name: "OneTimePassword", Flag: "one-time-password", Type: "*string", Required: true},
	{Name: "PhoneNumber", Flag: "phone-number", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"add-permission": {
			Name:   "add-permission",
			Fields: fields_add_permission,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddPermissionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_permission, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddPermission(ctx, input)
			},
		},
		"check-if-phone-number-is-opted-out": {
			Name:   "check-if-phone-number-is-opted-out",
			Fields: fields_check_if_phone_number_is_opted_out,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CheckIfPhoneNumberIsOptedOutInput{}
				if _, err := leanruntime.ApplyInput(input, fields_check_if_phone_number_is_opted_out, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CheckIfPhoneNumberIsOptedOut(ctx, input)
			},
		},
		"confirm-subscription": {
			Name:   "confirm-subscription",
			Fields: fields_confirm_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ConfirmSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_confirm_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ConfirmSubscription(ctx, input)
			},
		},
		"create-platform-application": {
			Name:   "create-platform-application",
			Fields: fields_create_platform_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePlatformApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_platform_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePlatformApplication(ctx, input)
			},
		},
		"create-platform-endpoint": {
			Name:   "create-platform-endpoint",
			Fields: fields_create_platform_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePlatformEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_platform_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePlatformEndpoint(ctx, input)
			},
		},
		"create-sms-sandbox-phone-number": {
			Name:   "create-sms-sandbox-phone-number",
			Fields: fields_create_sms_sandbox_phone_number,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSMSSandboxPhoneNumberInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_sms_sandbox_phone_number, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSMSSandboxPhoneNumber(ctx, input)
			},
		},
		"create-topic": {
			Name:   "create-topic",
			Fields: fields_create_topic,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTopicInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_topic, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTopic(ctx, input)
			},
		},
		"delete-endpoint": {
			Name:   "delete-endpoint",
			Fields: fields_delete_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEndpoint(ctx, input)
			},
		},
		"delete-platform-application": {
			Name:   "delete-platform-application",
			Fields: fields_delete_platform_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePlatformApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_platform_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePlatformApplication(ctx, input)
			},
		},
		"delete-sms-sandbox-phone-number": {
			Name:   "delete-sms-sandbox-phone-number",
			Fields: fields_delete_sms_sandbox_phone_number,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSMSSandboxPhoneNumberInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_sms_sandbox_phone_number, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSMSSandboxPhoneNumber(ctx, input)
			},
		},
		"delete-topic": {
			Name:   "delete-topic",
			Fields: fields_delete_topic,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTopicInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_topic, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTopic(ctx, input)
			},
		},
		"get-data-protection-policy": {
			Name:   "get-data-protection-policy",
			Fields: fields_get_data_protection_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDataProtectionPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_data_protection_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDataProtectionPolicy(ctx, input)
			},
		},
		"get-endpoint-attributes": {
			Name:   "get-endpoint-attributes",
			Fields: fields_get_endpoint_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEndpointAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_endpoint_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEndpointAttributes(ctx, input)
			},
		},
		"get-platform-application-attributes": {
			Name:   "get-platform-application-attributes",
			Fields: fields_get_platform_application_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPlatformApplicationAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_platform_application_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPlatformApplicationAttributes(ctx, input)
			},
		},
		"get-sms-attributes": {
			Name:   "get-sms-attributes",
			Fields: fields_get_sms_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSMSAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_sms_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSMSAttributes(ctx, input)
			},
		},
		"get-sms-sandbox-account-status": {
			Name:   "get-sms-sandbox-account-status",
			Fields: fields_get_sms_sandbox_account_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSMSSandboxAccountStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_sms_sandbox_account_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSMSSandboxAccountStatus(ctx, input)
			},
		},
		"get-subscription-attributes": {
			Name:   "get-subscription-attributes",
			Fields: fields_get_subscription_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSubscriptionAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_subscription_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSubscriptionAttributes(ctx, input)
			},
		},
		"get-topic-attributes": {
			Name:   "get-topic-attributes",
			Fields: fields_get_topic_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTopicAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_topic_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTopicAttributes(ctx, input)
			},
		},
		"list-endpoints-by-platform-application": {
			Name:   "list-endpoints-by-platform-application",
			Fields: fields_list_endpoints_by_platform_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEndpointsByPlatformApplicationInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_endpoints_by_platform_application, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEndpointsByPlatformApplication(ctx, input)
				}
				var results []*svc.ListEndpointsByPlatformApplicationOutput
				p := svc.NewListEndpointsByPlatformApplicationPaginator(client, input)
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
		"list-origination-numbers": {
			Name:   "list-origination-numbers",
			Fields: fields_list_origination_numbers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOriginationNumbersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_origination_numbers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListOriginationNumbers(ctx, input)
				}
				var results []*svc.ListOriginationNumbersOutput
				p := svc.NewListOriginationNumbersPaginator(client, input)
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
		"list-phone-numbers-opted-out": {
			Name:   "list-phone-numbers-opted-out",
			Fields: fields_list_phone_numbers_opted_out,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPhoneNumbersOptedOutInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_phone_numbers_opted_out, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPhoneNumbersOptedOut(ctx, input)
				}
				var results []*svc.ListPhoneNumbersOptedOutOutput
				p := svc.NewListPhoneNumbersOptedOutPaginator(client, input)
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
		"list-platform-applications": {
			Name:   "list-platform-applications",
			Fields: fields_list_platform_applications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPlatformApplicationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_platform_applications, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPlatformApplications(ctx, input)
				}
				var results []*svc.ListPlatformApplicationsOutput
				p := svc.NewListPlatformApplicationsPaginator(client, input)
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
		"list-sms-sandbox-phone-numbers": {
			Name:   "list-sms-sandbox-phone-numbers",
			Fields: fields_list_sms_sandbox_phone_numbers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSMSSandboxPhoneNumbersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_sms_sandbox_phone_numbers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSMSSandboxPhoneNumbers(ctx, input)
				}
				var results []*svc.ListSMSSandboxPhoneNumbersOutput
				p := svc.NewListSMSSandboxPhoneNumbersPaginator(client, input)
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
		"list-subscriptions": {
			Name:   "list-subscriptions",
			Fields: fields_list_subscriptions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSubscriptionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_subscriptions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSubscriptions(ctx, input)
				}
				var results []*svc.ListSubscriptionsOutput
				p := svc.NewListSubscriptionsPaginator(client, input)
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
		"list-subscriptions-by-topic": {
			Name:   "list-subscriptions-by-topic",
			Fields: fields_list_subscriptions_by_topic,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSubscriptionsByTopicInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_subscriptions_by_topic, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSubscriptionsByTopic(ctx, input)
				}
				var results []*svc.ListSubscriptionsByTopicOutput
				p := svc.NewListSubscriptionsByTopicPaginator(client, input)
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
		"list-topics": {
			Name:   "list-topics",
			Fields: fields_list_topics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTopicsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_topics, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTopics(ctx, input)
				}
				var results []*svc.ListTopicsOutput
				p := svc.NewListTopicsPaginator(client, input)
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
		"opt-in-phone-number": {
			Name:   "opt-in-phone-number",
			Fields: fields_opt_in_phone_number,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.OptInPhoneNumberInput{}
				if _, err := leanruntime.ApplyInput(input, fields_opt_in_phone_number, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.OptInPhoneNumber(ctx, input)
			},
		},
		"publish": {
			Name:   "publish",
			Fields: fields_publish,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PublishInput{}
				if _, err := leanruntime.ApplyInput(input, fields_publish, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.Publish(ctx, input)
			},
		},
		"publish-batch": {
			Name:   "publish-batch",
			Fields: fields_publish_batch,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PublishBatchInput{}
				if _, err := leanruntime.ApplyInput(input, fields_publish_batch, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PublishBatch(ctx, input)
			},
		},
		"put-data-protection-policy": {
			Name:   "put-data-protection-policy",
			Fields: fields_put_data_protection_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutDataProtectionPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_data_protection_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutDataProtectionPolicy(ctx, input)
			},
		},
		"remove-permission": {
			Name:   "remove-permission",
			Fields: fields_remove_permission,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemovePermissionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_permission, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemovePermission(ctx, input)
			},
		},
		"set-endpoint-attributes": {
			Name:   "set-endpoint-attributes",
			Fields: fields_set_endpoint_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetEndpointAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_endpoint_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetEndpointAttributes(ctx, input)
			},
		},
		"set-platform-application-attributes": {
			Name:   "set-platform-application-attributes",
			Fields: fields_set_platform_application_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetPlatformApplicationAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_platform_application_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetPlatformApplicationAttributes(ctx, input)
			},
		},
		"set-sms-attributes": {
			Name:   "set-sms-attributes",
			Fields: fields_set_sms_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetSMSAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_sms_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetSMSAttributes(ctx, input)
			},
		},
		"set-subscription-attributes": {
			Name:   "set-subscription-attributes",
			Fields: fields_set_subscription_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetSubscriptionAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_subscription_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetSubscriptionAttributes(ctx, input)
			},
		},
		"set-topic-attributes": {
			Name:   "set-topic-attributes",
			Fields: fields_set_topic_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetTopicAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_topic_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetTopicAttributes(ctx, input)
			},
		},
		"subscribe": {
			Name:   "subscribe",
			Fields: fields_subscribe,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SubscribeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_subscribe, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.Subscribe(ctx, input)
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
		"unsubscribe": {
			Name:   "unsubscribe",
			Fields: fields_unsubscribe,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UnsubscribeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_unsubscribe, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.Unsubscribe(ctx, input)
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
		"verify-sms-sandbox-phone-number": {
			Name:   "verify-sms-sandbox-phone-number",
			Fields: fields_verify_sms_sandbox_phone_number,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.VerifySMSSandboxPhoneNumberInput{}
				if _, err := leanruntime.ApplyInput(input, fields_verify_sms_sandbox_phone_number, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.VerifySMSSandboxPhoneNumber(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("sns", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
