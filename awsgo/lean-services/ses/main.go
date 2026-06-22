package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/ses"
)

var fields_clone_receipt_rule_set = []leanruntime.Field{
	{Name: "OriginalRuleSetName", Flag: "original-rule-set-name", Type: "*string", Required: true},
	{Name: "RuleSetName", Flag: "rule-set-name", Type: "*string", Required: true},
}

var fields_create_configuration_set = []leanruntime.Field{
	{Name: "ConfigurationSet", Flag: "configuration-set", Type: "*types.ConfigurationSet", Required: true},
}

var fields_create_configuration_set_event_destination = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
	{Name: "EventDestination", Flag: "event-destination", Type: "*types.EventDestination", Required: true},
}

var fields_create_configuration_set_tracking_options = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
	{Name: "TrackingOptions", Flag: "tracking-options", Type: "*types.TrackingOptions", Required: true},
}

var fields_create_custom_verification_email_template = []leanruntime.Field{
	{Name: "FailureRedirectionURL", Flag: "failure-redirection-url", Type: "*string", Required: true},
	{Name: "FromEmailAddress", Flag: "from-email-address", Type: "*string", Required: true},
	{Name: "SuccessRedirectionURL", Flag: "success-redirection-url", Type: "*string", Required: true},
	{Name: "TemplateContent", Flag: "template-content", Type: "*string", Required: true},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
	{Name: "TemplateSubject", Flag: "template-subject", Type: "*string", Required: true},
}

var fields_create_receipt_filter = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.ReceiptFilter", Required: true},
}

var fields_create_receipt_rule = []leanruntime.Field{
	{Name: "After", Flag: "after", Type: "*string", Required: false},
	{Name: "Rule", Flag: "rule", Type: "*types.ReceiptRule", Required: true},
	{Name: "RuleSetName", Flag: "rule-set-name", Type: "*string", Required: true},
}

var fields_create_receipt_rule_set = []leanruntime.Field{
	{Name: "RuleSetName", Flag: "rule-set-name", Type: "*string", Required: true},
}

var fields_create_template = []leanruntime.Field{
	{Name: "Template", Flag: "template", Type: "*types.Template", Required: true},
}

var fields_delete_configuration_set = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
}

var fields_delete_configuration_set_event_destination = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
	{Name: "EventDestinationName", Flag: "event-destination-name", Type: "*string", Required: true},
}

var fields_delete_configuration_set_tracking_options = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
}

var fields_delete_custom_verification_email_template = []leanruntime.Field{
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
}

var fields_delete_identity = []leanruntime.Field{
	{Name: "Identity", Flag: "identity", Type: "*string", Required: true},
}

var fields_delete_identity_policy = []leanruntime.Field{
	{Name: "Identity", Flag: "identity", Type: "*string", Required: true},
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
}

var fields_delete_receipt_filter = []leanruntime.Field{
	{Name: "FilterName", Flag: "filter-name", Type: "*string", Required: true},
}

var fields_delete_receipt_rule = []leanruntime.Field{
	{Name: "RuleName", Flag: "rule-name", Type: "*string", Required: true},
	{Name: "RuleSetName", Flag: "rule-set-name", Type: "*string", Required: true},
}

var fields_delete_receipt_rule_set = []leanruntime.Field{
	{Name: "RuleSetName", Flag: "rule-set-name", Type: "*string", Required: true},
}

var fields_delete_template = []leanruntime.Field{
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
}

var fields_delete_verified_email_address = []leanruntime.Field{
	{Name: "EmailAddress", Flag: "email-address", Type: "*string", Required: true},
}

var fields_describe_active_receipt_rule_set = []leanruntime.Field{}

var fields_describe_configuration_set = []leanruntime.Field{
	{Name: "ConfigurationSetAttributeNames", Flag: "configuration-set-attribute-names", Type: "[]types.ConfigurationSetAttribute", Required: false},
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
}

var fields_describe_receipt_rule = []leanruntime.Field{
	{Name: "RuleName", Flag: "rule-name", Type: "*string", Required: true},
	{Name: "RuleSetName", Flag: "rule-set-name", Type: "*string", Required: true},
}

var fields_describe_receipt_rule_set = []leanruntime.Field{
	{Name: "RuleSetName", Flag: "rule-set-name", Type: "*string", Required: true},
}

var fields_get_account_sending_enabled = []leanruntime.Field{}

var fields_get_custom_verification_email_template = []leanruntime.Field{
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
}

var fields_get_identity_dkim_attributes = []leanruntime.Field{
	{Name: "Identities", Flag: "identities", Type: "[]string", Required: true},
}

var fields_get_identity_mail_from_domain_attributes = []leanruntime.Field{
	{Name: "Identities", Flag: "identities", Type: "[]string", Required: true},
}

var fields_get_identity_notification_attributes = []leanruntime.Field{
	{Name: "Identities", Flag: "identities", Type: "[]string", Required: true},
}

var fields_get_identity_policies = []leanruntime.Field{
	{Name: "Identity", Flag: "identity", Type: "*string", Required: true},
	{Name: "PolicyNames", Flag: "policy-names", Type: "[]string", Required: true},
}

var fields_get_identity_verification_attributes = []leanruntime.Field{
	{Name: "Identities", Flag: "identities", Type: "[]string", Required: true},
}

var fields_get_send_quota = []leanruntime.Field{}

var fields_get_send_statistics = []leanruntime.Field{}

var fields_get_template = []leanruntime.Field{
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
}

var fields_list_configuration_sets = []leanruntime.Field{
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_custom_verification_email_templates = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_identities = []leanruntime.Field{
	{Name: "IdentityType", Flag: "identity-type", Type: "types.IdentityType", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_identity_policies = []leanruntime.Field{
	{Name: "Identity", Flag: "identity", Type: "*string", Required: true},
}

var fields_list_receipt_filters = []leanruntime.Field{}

var fields_list_receipt_rule_sets = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_templates = []leanruntime.Field{
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_verified_email_addresses = []leanruntime.Field{}

var fields_put_configuration_set_delivery_options = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
	{Name: "DeliveryOptions", Flag: "delivery-options", Type: "*types.DeliveryOptions", Required: false},
}

var fields_put_identity_policy = []leanruntime.Field{
	{Name: "Identity", Flag: "identity", Type: "*string", Required: true},
	{Name: "Policy", Flag: "policy", Type: "*string", Required: true},
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
}

var fields_reorder_receipt_rule_set = []leanruntime.Field{
	{Name: "RuleNames", Flag: "rule-names", Type: "[]string", Required: true},
	{Name: "RuleSetName", Flag: "rule-set-name", Type: "*string", Required: true},
}

var fields_send_bounce = []leanruntime.Field{
	{Name: "BounceSender", Flag: "bounce-sender", Type: "*string", Required: true},
	{Name: "BounceSenderArn", Flag: "bounce-sender-arn", Type: "*string", Required: false},
	{Name: "BouncedRecipientInfoList", Flag: "bounced-recipient-info-list", Type: "[]types.BouncedRecipientInfo", Required: true},
	{Name: "Explanation", Flag: "explanation", Type: "*string", Required: false},
	{Name: "MessageDsn", Flag: "message-dsn", Type: "*types.MessageDsn", Required: false},
	{Name: "OriginalMessageId", Flag: "original-message-id", Type: "*string", Required: true},
}

var fields_send_bulk_templated_email = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: false},
	{Name: "DefaultTags", Flag: "default-tags", Type: "[]types.MessageTag", Required: false},
	{Name: "DefaultTemplateData", Flag: "default-template-data", Type: "*string", Required: true},
	{Name: "Destinations", Flag: "destinations", Type: "[]types.BulkEmailDestination", Required: true},
	{Name: "ReplyToAddresses", Flag: "reply-to-addresses", Type: "[]string", Required: false},
	{Name: "ReturnPath", Flag: "return-path", Type: "*string", Required: false},
	{Name: "ReturnPathArn", Flag: "return-path-arn", Type: "*string", Required: false},
	{Name: "Source", Flag: "source", Type: "*string", Required: true},
	{Name: "SourceArn", Flag: "source-arn", Type: "*string", Required: false},
	{Name: "Template", Flag: "template", Type: "*string", Required: true},
	{Name: "TemplateArn", Flag: "template-arn", Type: "*string", Required: false},
}

var fields_send_custom_verification_email = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: false},
	{Name: "EmailAddress", Flag: "email-address", Type: "*string", Required: true},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
}

var fields_send_email = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: false},
	{Name: "Destination", Flag: "destination", Type: "*types.Destination", Required: true},
	{Name: "Message", Flag: "message", Type: "*types.Message", Required: true},
	{Name: "ReplyToAddresses", Flag: "reply-to-addresses", Type: "[]string", Required: false},
	{Name: "ReturnPath", Flag: "return-path", Type: "*string", Required: false},
	{Name: "ReturnPathArn", Flag: "return-path-arn", Type: "*string", Required: false},
	{Name: "Source", Flag: "source", Type: "*string", Required: true},
	{Name: "SourceArn", Flag: "source-arn", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.MessageTag", Required: false},
}

var fields_send_raw_email = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: false},
	{Name: "Destinations", Flag: "destinations", Type: "[]string", Required: false},
	{Name: "FromArn", Flag: "from-arn", Type: "*string", Required: false},
	{Name: "RawMessage", Flag: "raw-message", Type: "*types.RawMessage", Required: true},
	{Name: "ReturnPathArn", Flag: "return-path-arn", Type: "*string", Required: false},
	{Name: "Source", Flag: "source", Type: "*string", Required: false},
	{Name: "SourceArn", Flag: "source-arn", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.MessageTag", Required: false},
}

var fields_send_templated_email = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: false},
	{Name: "Destination", Flag: "destination", Type: "*types.Destination", Required: true},
	{Name: "ReplyToAddresses", Flag: "reply-to-addresses", Type: "[]string", Required: false},
	{Name: "ReturnPath", Flag: "return-path", Type: "*string", Required: false},
	{Name: "ReturnPathArn", Flag: "return-path-arn", Type: "*string", Required: false},
	{Name: "Source", Flag: "source", Type: "*string", Required: true},
	{Name: "SourceArn", Flag: "source-arn", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.MessageTag", Required: false},
	{Name: "Template", Flag: "template", Type: "*string", Required: true},
	{Name: "TemplateArn", Flag: "template-arn", Type: "*string", Required: false},
	{Name: "TemplateData", Flag: "template-data", Type: "*string", Required: true},
}

var fields_set_active_receipt_rule_set = []leanruntime.Field{
	{Name: "RuleSetName", Flag: "rule-set-name", Type: "*string", Required: false},
}

var fields_set_identity_dkim_enabled = []leanruntime.Field{
	{Name: "DkimEnabled", Flag: "dkim-enabled", Type: "bool", Required: true},
	{Name: "Identity", Flag: "identity", Type: "*string", Required: true},
}

var fields_set_identity_feedback_forwarding_enabled = []leanruntime.Field{
	{Name: "ForwardingEnabled", Flag: "forwarding-enabled", Type: "bool", Required: true},
	{Name: "Identity", Flag: "identity", Type: "*string", Required: true},
}

var fields_set_identity_headers_in_notifications_enabled = []leanruntime.Field{
	{Name: "Enabled", Flag: "enabled", Type: "bool", Required: true},
	{Name: "Identity", Flag: "identity", Type: "*string", Required: true},
	{Name: "NotificationType", Flag: "notification-type", Type: "types.NotificationType", Required: true},
}

var fields_set_identity_mail_from_domain = []leanruntime.Field{
	{Name: "BehaviorOnMXFailure", Flag: "behavior-on-mx-failure", Type: "types.BehaviorOnMXFailure", Required: false},
	{Name: "Identity", Flag: "identity", Type: "*string", Required: true},
	{Name: "MailFromDomain", Flag: "mail-from-domain", Type: "*string", Required: false},
}

var fields_set_identity_notification_topic = []leanruntime.Field{
	{Name: "Identity", Flag: "identity", Type: "*string", Required: true},
	{Name: "NotificationType", Flag: "notification-type", Type: "types.NotificationType", Required: true},
	{Name: "SnsTopic", Flag: "sns-topic", Type: "*string", Required: false},
}

var fields_set_receipt_rule_position = []leanruntime.Field{
	{Name: "After", Flag: "after", Type: "*string", Required: false},
	{Name: "RuleName", Flag: "rule-name", Type: "*string", Required: true},
	{Name: "RuleSetName", Flag: "rule-set-name", Type: "*string", Required: true},
}

var fields_test_render_template = []leanruntime.Field{
	{Name: "TemplateData", Flag: "template-data", Type: "*string", Required: true},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
}

var fields_update_account_sending_enabled = []leanruntime.Field{
	{Name: "Enabled", Flag: "enabled", Type: "bool", Required: false},
}

var fields_update_configuration_set_event_destination = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
	{Name: "EventDestination", Flag: "event-destination", Type: "*types.EventDestination", Required: true},
}

var fields_update_configuration_set_reputation_metrics_enabled = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
	{Name: "Enabled", Flag: "enabled", Type: "bool", Required: true},
}

var fields_update_configuration_set_sending_enabled = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
	{Name: "Enabled", Flag: "enabled", Type: "bool", Required: true},
}

var fields_update_configuration_set_tracking_options = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
	{Name: "TrackingOptions", Flag: "tracking-options", Type: "*types.TrackingOptions", Required: true},
}

var fields_update_custom_verification_email_template = []leanruntime.Field{
	{Name: "FailureRedirectionURL", Flag: "failure-redirection-url", Type: "*string", Required: false},
	{Name: "FromEmailAddress", Flag: "from-email-address", Type: "*string", Required: false},
	{Name: "SuccessRedirectionURL", Flag: "success-redirection-url", Type: "*string", Required: false},
	{Name: "TemplateContent", Flag: "template-content", Type: "*string", Required: false},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
	{Name: "TemplateSubject", Flag: "template-subject", Type: "*string", Required: false},
}

var fields_update_receipt_rule = []leanruntime.Field{
	{Name: "Rule", Flag: "rule", Type: "*types.ReceiptRule", Required: true},
	{Name: "RuleSetName", Flag: "rule-set-name", Type: "*string", Required: true},
}

var fields_update_template = []leanruntime.Field{
	{Name: "Template", Flag: "template", Type: "*types.Template", Required: true},
}

var fields_verify_domain_dkim = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
}

var fields_verify_domain_identity = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
}

var fields_verify_email_address = []leanruntime.Field{
	{Name: "EmailAddress", Flag: "email-address", Type: "*string", Required: true},
}

var fields_verify_email_identity = []leanruntime.Field{
	{Name: "EmailAddress", Flag: "email-address", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"clone-receipt-rule-set": {
			Name:   "clone-receipt-rule-set",
			Fields: fields_clone_receipt_rule_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CloneReceiptRuleSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_clone_receipt_rule_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CloneReceiptRuleSet(ctx, input)
			},
		},
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
		"create-configuration-set-tracking-options": {
			Name:   "create-configuration-set-tracking-options",
			Fields: fields_create_configuration_set_tracking_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateConfigurationSetTrackingOptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_configuration_set_tracking_options, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateConfigurationSetTrackingOptions(ctx, input)
			},
		},
		"create-custom-verification-email-template": {
			Name:   "create-custom-verification-email-template",
			Fields: fields_create_custom_verification_email_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCustomVerificationEmailTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_custom_verification_email_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCustomVerificationEmailTemplate(ctx, input)
			},
		},
		"create-receipt-filter": {
			Name:   "create-receipt-filter",
			Fields: fields_create_receipt_filter,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateReceiptFilterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_receipt_filter, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateReceiptFilter(ctx, input)
			},
		},
		"create-receipt-rule": {
			Name:   "create-receipt-rule",
			Fields: fields_create_receipt_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateReceiptRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_receipt_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateReceiptRule(ctx, input)
			},
		},
		"create-receipt-rule-set": {
			Name:   "create-receipt-rule-set",
			Fields: fields_create_receipt_rule_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateReceiptRuleSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_receipt_rule_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateReceiptRuleSet(ctx, input)
			},
		},
		"create-template": {
			Name:   "create-template",
			Fields: fields_create_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTemplate(ctx, input)
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
		"delete-configuration-set-tracking-options": {
			Name:   "delete-configuration-set-tracking-options",
			Fields: fields_delete_configuration_set_tracking_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConfigurationSetTrackingOptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_configuration_set_tracking_options, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConfigurationSetTrackingOptions(ctx, input)
			},
		},
		"delete-custom-verification-email-template": {
			Name:   "delete-custom-verification-email-template",
			Fields: fields_delete_custom_verification_email_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCustomVerificationEmailTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_custom_verification_email_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCustomVerificationEmailTemplate(ctx, input)
			},
		},
		"delete-identity": {
			Name:   "delete-identity",
			Fields: fields_delete_identity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIdentityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_identity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIdentity(ctx, input)
			},
		},
		"delete-identity-policy": {
			Name:   "delete-identity-policy",
			Fields: fields_delete_identity_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIdentityPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_identity_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIdentityPolicy(ctx, input)
			},
		},
		"delete-receipt-filter": {
			Name:   "delete-receipt-filter",
			Fields: fields_delete_receipt_filter,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteReceiptFilterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_receipt_filter, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteReceiptFilter(ctx, input)
			},
		},
		"delete-receipt-rule": {
			Name:   "delete-receipt-rule",
			Fields: fields_delete_receipt_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteReceiptRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_receipt_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteReceiptRule(ctx, input)
			},
		},
		"delete-receipt-rule-set": {
			Name:   "delete-receipt-rule-set",
			Fields: fields_delete_receipt_rule_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteReceiptRuleSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_receipt_rule_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteReceiptRuleSet(ctx, input)
			},
		},
		"delete-template": {
			Name:   "delete-template",
			Fields: fields_delete_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTemplate(ctx, input)
			},
		},
		"delete-verified-email-address": {
			Name:   "delete-verified-email-address",
			Fields: fields_delete_verified_email_address,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVerifiedEmailAddressInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_verified_email_address, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVerifiedEmailAddress(ctx, input)
			},
		},
		"describe-active-receipt-rule-set": {
			Name:   "describe-active-receipt-rule-set",
			Fields: fields_describe_active_receipt_rule_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeActiveReceiptRuleSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_active_receipt_rule_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeActiveReceiptRuleSet(ctx, input)
			},
		},
		"describe-configuration-set": {
			Name:   "describe-configuration-set",
			Fields: fields_describe_configuration_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeConfigurationSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_configuration_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeConfigurationSet(ctx, input)
			},
		},
		"describe-receipt-rule": {
			Name:   "describe-receipt-rule",
			Fields: fields_describe_receipt_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeReceiptRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_receipt_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeReceiptRule(ctx, input)
			},
		},
		"describe-receipt-rule-set": {
			Name:   "describe-receipt-rule-set",
			Fields: fields_describe_receipt_rule_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeReceiptRuleSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_receipt_rule_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeReceiptRuleSet(ctx, input)
			},
		},
		"get-account-sending-enabled": {
			Name:   "get-account-sending-enabled",
			Fields: fields_get_account_sending_enabled,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccountSendingEnabledInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_account_sending_enabled, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccountSendingEnabled(ctx, input)
			},
		},
		"get-custom-verification-email-template": {
			Name:   "get-custom-verification-email-template",
			Fields: fields_get_custom_verification_email_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCustomVerificationEmailTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_custom_verification_email_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCustomVerificationEmailTemplate(ctx, input)
			},
		},
		"get-identity-dkim-attributes": {
			Name:   "get-identity-dkim-attributes",
			Fields: fields_get_identity_dkim_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIdentityDkimAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_identity_dkim_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIdentityDkimAttributes(ctx, input)
			},
		},
		"get-identity-mail-from-domain-attributes": {
			Name:   "get-identity-mail-from-domain-attributes",
			Fields: fields_get_identity_mail_from_domain_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIdentityMailFromDomainAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_identity_mail_from_domain_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIdentityMailFromDomainAttributes(ctx, input)
			},
		},
		"get-identity-notification-attributes": {
			Name:   "get-identity-notification-attributes",
			Fields: fields_get_identity_notification_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIdentityNotificationAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_identity_notification_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIdentityNotificationAttributes(ctx, input)
			},
		},
		"get-identity-policies": {
			Name:   "get-identity-policies",
			Fields: fields_get_identity_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIdentityPoliciesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_identity_policies, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIdentityPolicies(ctx, input)
			},
		},
		"get-identity-verification-attributes": {
			Name:   "get-identity-verification-attributes",
			Fields: fields_get_identity_verification_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIdentityVerificationAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_identity_verification_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIdentityVerificationAttributes(ctx, input)
			},
		},
		"get-send-quota": {
			Name:   "get-send-quota",
			Fields: fields_get_send_quota,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSendQuotaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_send_quota, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSendQuota(ctx, input)
			},
		},
		"get-send-statistics": {
			Name:   "get-send-statistics",
			Fields: fields_get_send_statistics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSendStatisticsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_send_statistics, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSendStatistics(ctx, input)
			},
		},
		"get-template": {
			Name:   "get-template",
			Fields: fields_get_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTemplate(ctx, input)
			},
		},
		"list-configuration-sets": {
			Name:   "list-configuration-sets",
			Fields: fields_list_configuration_sets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConfigurationSetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_configuration_sets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListConfigurationSets(ctx, input)
			},
		},
		"list-custom-verification-email-templates": {
			Name:   "list-custom-verification-email-templates",
			Fields: fields_list_custom_verification_email_templates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCustomVerificationEmailTemplatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_custom_verification_email_templates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCustomVerificationEmailTemplates(ctx, input)
				}
				var results []*svc.ListCustomVerificationEmailTemplatesOutput
				p := svc.NewListCustomVerificationEmailTemplatesPaginator(client, input)
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
		"list-identities": {
			Name:   "list-identities",
			Fields: fields_list_identities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIdentitiesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_identities, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListIdentities(ctx, input)
				}
				var results []*svc.ListIdentitiesOutput
				p := svc.NewListIdentitiesPaginator(client, input)
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
		"list-identity-policies": {
			Name:   "list-identity-policies",
			Fields: fields_list_identity_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIdentityPoliciesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_identity_policies, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListIdentityPolicies(ctx, input)
			},
		},
		"list-receipt-filters": {
			Name:   "list-receipt-filters",
			Fields: fields_list_receipt_filters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListReceiptFiltersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_receipt_filters, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListReceiptFilters(ctx, input)
			},
		},
		"list-receipt-rule-sets": {
			Name:   "list-receipt-rule-sets",
			Fields: fields_list_receipt_rule_sets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListReceiptRuleSetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_receipt_rule_sets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListReceiptRuleSets(ctx, input)
			},
		},
		"list-templates": {
			Name:   "list-templates",
			Fields: fields_list_templates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTemplatesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_templates, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTemplates(ctx, input)
			},
		},
		"list-verified-email-addresses": {
			Name:   "list-verified-email-addresses",
			Fields: fields_list_verified_email_addresses,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVerifiedEmailAddressesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_verified_email_addresses, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListVerifiedEmailAddresses(ctx, input)
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
		"put-identity-policy": {
			Name:   "put-identity-policy",
			Fields: fields_put_identity_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutIdentityPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_identity_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutIdentityPolicy(ctx, input)
			},
		},
		"reorder-receipt-rule-set": {
			Name:   "reorder-receipt-rule-set",
			Fields: fields_reorder_receipt_rule_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ReorderReceiptRuleSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reorder_receipt_rule_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ReorderReceiptRuleSet(ctx, input)
			},
		},
		"send-bounce": {
			Name:   "send-bounce",
			Fields: fields_send_bounce,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendBounceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_bounce, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendBounce(ctx, input)
			},
		},
		"send-bulk-templated-email": {
			Name:   "send-bulk-templated-email",
			Fields: fields_send_bulk_templated_email,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendBulkTemplatedEmailInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_bulk_templated_email, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendBulkTemplatedEmail(ctx, input)
			},
		},
		"send-custom-verification-email": {
			Name:   "send-custom-verification-email",
			Fields: fields_send_custom_verification_email,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendCustomVerificationEmailInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_custom_verification_email, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendCustomVerificationEmail(ctx, input)
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
		"send-raw-email": {
			Name:   "send-raw-email",
			Fields: fields_send_raw_email,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendRawEmailInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_raw_email, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendRawEmail(ctx, input)
			},
		},
		"send-templated-email": {
			Name:   "send-templated-email",
			Fields: fields_send_templated_email,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendTemplatedEmailInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_templated_email, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendTemplatedEmail(ctx, input)
			},
		},
		"set-active-receipt-rule-set": {
			Name:   "set-active-receipt-rule-set",
			Fields: fields_set_active_receipt_rule_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetActiveReceiptRuleSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_active_receipt_rule_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetActiveReceiptRuleSet(ctx, input)
			},
		},
		"set-identity-dkim-enabled": {
			Name:   "set-identity-dkim-enabled",
			Fields: fields_set_identity_dkim_enabled,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetIdentityDkimEnabledInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_identity_dkim_enabled, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetIdentityDkimEnabled(ctx, input)
			},
		},
		"set-identity-feedback-forwarding-enabled": {
			Name:   "set-identity-feedback-forwarding-enabled",
			Fields: fields_set_identity_feedback_forwarding_enabled,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetIdentityFeedbackForwardingEnabledInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_identity_feedback_forwarding_enabled, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetIdentityFeedbackForwardingEnabled(ctx, input)
			},
		},
		"set-identity-headers-in-notifications-enabled": {
			Name:   "set-identity-headers-in-notifications-enabled",
			Fields: fields_set_identity_headers_in_notifications_enabled,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetIdentityHeadersInNotificationsEnabledInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_identity_headers_in_notifications_enabled, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetIdentityHeadersInNotificationsEnabled(ctx, input)
			},
		},
		"set-identity-mail-from-domain": {
			Name:   "set-identity-mail-from-domain",
			Fields: fields_set_identity_mail_from_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetIdentityMailFromDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_identity_mail_from_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetIdentityMailFromDomain(ctx, input)
			},
		},
		"set-identity-notification-topic": {
			Name:   "set-identity-notification-topic",
			Fields: fields_set_identity_notification_topic,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetIdentityNotificationTopicInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_identity_notification_topic, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetIdentityNotificationTopic(ctx, input)
			},
		},
		"set-receipt-rule-position": {
			Name:   "set-receipt-rule-position",
			Fields: fields_set_receipt_rule_position,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetReceiptRulePositionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_receipt_rule_position, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetReceiptRulePosition(ctx, input)
			},
		},
		"test-render-template": {
			Name:   "test-render-template",
			Fields: fields_test_render_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TestRenderTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_test_render_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TestRenderTemplate(ctx, input)
			},
		},
		"update-account-sending-enabled": {
			Name:   "update-account-sending-enabled",
			Fields: fields_update_account_sending_enabled,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAccountSendingEnabledInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_account_sending_enabled, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAccountSendingEnabled(ctx, input)
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
		"update-configuration-set-reputation-metrics-enabled": {
			Name:   "update-configuration-set-reputation-metrics-enabled",
			Fields: fields_update_configuration_set_reputation_metrics_enabled,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateConfigurationSetReputationMetricsEnabledInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_configuration_set_reputation_metrics_enabled, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateConfigurationSetReputationMetricsEnabled(ctx, input)
			},
		},
		"update-configuration-set-sending-enabled": {
			Name:   "update-configuration-set-sending-enabled",
			Fields: fields_update_configuration_set_sending_enabled,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateConfigurationSetSendingEnabledInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_configuration_set_sending_enabled, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateConfigurationSetSendingEnabled(ctx, input)
			},
		},
		"update-configuration-set-tracking-options": {
			Name:   "update-configuration-set-tracking-options",
			Fields: fields_update_configuration_set_tracking_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateConfigurationSetTrackingOptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_configuration_set_tracking_options, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateConfigurationSetTrackingOptions(ctx, input)
			},
		},
		"update-custom-verification-email-template": {
			Name:   "update-custom-verification-email-template",
			Fields: fields_update_custom_verification_email_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCustomVerificationEmailTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_custom_verification_email_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCustomVerificationEmailTemplate(ctx, input)
			},
		},
		"update-receipt-rule": {
			Name:   "update-receipt-rule",
			Fields: fields_update_receipt_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateReceiptRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_receipt_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateReceiptRule(ctx, input)
			},
		},
		"update-template": {
			Name:   "update-template",
			Fields: fields_update_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTemplate(ctx, input)
			},
		},
		"verify-domain-dkim": {
			Name:   "verify-domain-dkim",
			Fields: fields_verify_domain_dkim,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.VerifyDomainDkimInput{}
				if _, err := leanruntime.ApplyInput(input, fields_verify_domain_dkim, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.VerifyDomainDkim(ctx, input)
			},
		},
		"verify-domain-identity": {
			Name:   "verify-domain-identity",
			Fields: fields_verify_domain_identity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.VerifyDomainIdentityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_verify_domain_identity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.VerifyDomainIdentity(ctx, input)
			},
		},
		"verify-email-address": {
			Name:   "verify-email-address",
			Fields: fields_verify_email_address,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.VerifyEmailAddressInput{}
				if _, err := leanruntime.ApplyInput(input, fields_verify_email_address, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.VerifyEmailAddress(ctx, input)
			},
		},
		"verify-email-identity": {
			Name:   "verify-email-identity",
			Fields: fields_verify_email_identity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.VerifyEmailIdentityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_verify_email_identity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.VerifyEmailIdentity(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("ses", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
