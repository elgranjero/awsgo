package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoicev2"
)

var fields_associate_origination_identity = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "IsoCountryCode", Flag: "iso-country-code", Type: "*string", Required: true},
	{Name: "OriginationIdentity", Flag: "origination-identity", Type: "*string", Required: true},
	{Name: "PoolId", Flag: "pool-id", Type: "*string", Required: true},
}

var fields_associate_protect_configuration = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
	{Name: "ProtectConfigurationId", Flag: "protect-configuration-id", Type: "*string", Required: true},
}

var fields_carrier_lookup = []leanruntime.Field{
	{Name: "PhoneNumber", Flag: "phone-number", Type: "*string", Required: true},
}

var fields_create_configuration_set = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_event_destination = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CloudWatchLogsDestination", Flag: "cloud-watch-logs-destination", Type: "*types.CloudWatchLogsDestination", Required: false},
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
	{Name: "EventDestinationName", Flag: "event-destination-name", Type: "*string", Required: true},
	{Name: "KinesisFirehoseDestination", Flag: "kinesis-firehose-destination", Type: "*types.KinesisFirehoseDestination", Required: false},
	{Name: "MatchingEventTypes", Flag: "matching-event-types", Type: "[]types.EventType", Required: true},
	{Name: "SnsDestination", Flag: "sns-destination", Type: "*types.SnsDestination", Required: false},
}

var fields_create_opt_out_list = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "OptOutListName", Flag: "opt-out-list-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_pool = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DeletionProtectionEnabled", Flag: "deletion-protection-enabled", Type: "*bool", Required: false},
	{Name: "IsoCountryCode", Flag: "iso-country-code", Type: "*string", Required: true},
	{Name: "MessageType", Flag: "message-type", Type: "types.MessageType", Required: true},
	{Name: "OriginationIdentity", Flag: "origination-identity", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_protect_configuration = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DeletionProtectionEnabled", Flag: "deletion-protection-enabled", Type: "*bool", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_registration = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "RegistrationType", Flag: "registration-type", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_registration_association = []leanruntime.Field{
	{Name: "RegistrationId", Flag: "registration-id", Type: "*string", Required: true},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
}

var fields_create_registration_attachment = []leanruntime.Field{
	{Name: "AttachmentBody", Flag: "attachment-body", Type: "[]byte", Required: false},
	{Name: "AttachmentUrl", Flag: "attachment-url", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_registration_version = []leanruntime.Field{
	{Name: "RegistrationId", Flag: "registration-id", Type: "*string", Required: true},
}

var fields_create_verified_destination_number = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DestinationPhoneNumber", Flag: "destination-phone-number", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_delete_account_default_protect_configuration = []leanruntime.Field{}

var fields_delete_configuration_set = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
}

var fields_delete_default_message_type = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
}

var fields_delete_default_sender_id = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
}

var fields_delete_event_destination = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
	{Name: "EventDestinationName", Flag: "event-destination-name", Type: "*string", Required: true},
}

var fields_delete_keyword = []leanruntime.Field{
	{Name: "Keyword", Flag: "keyword", Type: "*string", Required: true},
	{Name: "OriginationIdentity", Flag: "origination-identity", Type: "*string", Required: true},
}

var fields_delete_media_message_spend_limit_override = []leanruntime.Field{}

var fields_delete_opt_out_list = []leanruntime.Field{
	{Name: "OptOutListName", Flag: "opt-out-list-name", Type: "*string", Required: true},
}

var fields_delete_opted_out_number = []leanruntime.Field{
	{Name: "OptOutListName", Flag: "opt-out-list-name", Type: "*string", Required: true},
	{Name: "OptedOutNumber", Flag: "opted-out-number", Type: "*string", Required: true},
}

var fields_delete_pool = []leanruntime.Field{
	{Name: "PoolId", Flag: "pool-id", Type: "*string", Required: true},
}

var fields_delete_protect_configuration = []leanruntime.Field{
	{Name: "ProtectConfigurationId", Flag: "protect-configuration-id", Type: "*string", Required: true},
}

var fields_delete_protect_configuration_rule_set_number_override = []leanruntime.Field{
	{Name: "DestinationPhoneNumber", Flag: "destination-phone-number", Type: "*string", Required: true},
	{Name: "ProtectConfigurationId", Flag: "protect-configuration-id", Type: "*string", Required: true},
}

var fields_delete_registration = []leanruntime.Field{
	{Name: "RegistrationId", Flag: "registration-id", Type: "*string", Required: true},
}

var fields_delete_registration_attachment = []leanruntime.Field{
	{Name: "RegistrationAttachmentId", Flag: "registration-attachment-id", Type: "*string", Required: true},
}

var fields_delete_registration_field_value = []leanruntime.Field{
	{Name: "FieldPath", Flag: "field-path", Type: "*string", Required: true},
	{Name: "RegistrationId", Flag: "registration-id", Type: "*string", Required: true},
}

var fields_delete_resource_policy = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_delete_text_message_spend_limit_override = []leanruntime.Field{}

var fields_delete_verified_destination_number = []leanruntime.Field{
	{Name: "VerifiedDestinationNumberId", Flag: "verified-destination-number-id", Type: "*string", Required: true},
}

var fields_delete_voice_message_spend_limit_override = []leanruntime.Field{}

var fields_describe_account_attributes = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_account_limits = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_configuration_sets = []leanruntime.Field{
	{Name: "ConfigurationSetNames", Flag: "configuration-set-names", Type: "[]string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.ConfigurationSetFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_keywords = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.KeywordFilter", Required: false},
	{Name: "Keywords", Flag: "keywords", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OriginationIdentity", Flag: "origination-identity", Type: "*string", Required: true},
}

var fields_describe_opt_out_lists = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OptOutListNames", Flag: "opt-out-list-names", Type: "[]string", Required: false},
	{Name: "Owner", Flag: "owner", Type: "types.Owner", Required: false},
}

var fields_describe_opted_out_numbers = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.OptedOutFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OptOutListName", Flag: "opt-out-list-name", Type: "*string", Required: true},
	{Name: "OptedOutNumbers", Flag: "opted-out-numbers", Type: "[]string", Required: false},
}

var fields_describe_phone_numbers = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.PhoneNumberFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Owner", Flag: "owner", Type: "types.Owner", Required: false},
	{Name: "PhoneNumberIds", Flag: "phone-number-ids", Type: "[]string", Required: false},
}

var fields_describe_pools = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.PoolFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Owner", Flag: "owner", Type: "types.Owner", Required: false},
	{Name: "PoolIds", Flag: "pool-ids", Type: "[]string", Required: false},
}

var fields_describe_protect_configurations = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.ProtectConfigurationFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProtectConfigurationIds", Flag: "protect-configuration-ids", Type: "[]string", Required: false},
}

var fields_describe_registration_attachments = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.RegistrationAttachmentFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RegistrationAttachmentIds", Flag: "registration-attachment-ids", Type: "[]string", Required: false},
}

var fields_describe_registration_field_definitions = []leanruntime.Field{
	{Name: "FieldPaths", Flag: "field-paths", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RegistrationType", Flag: "registration-type", Type: "*string", Required: true},
	{Name: "SectionPath", Flag: "section-path", Type: "*string", Required: false},
}

var fields_describe_registration_field_values = []leanruntime.Field{
	{Name: "FieldPaths", Flag: "field-paths", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RegistrationId", Flag: "registration-id", Type: "*string", Required: true},
	{Name: "SectionPath", Flag: "section-path", Type: "*string", Required: false},
	{Name: "VersionNumber", Flag: "version-number", Type: "*int64", Required: false},
}

var fields_describe_registration_section_definitions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RegistrationType", Flag: "registration-type", Type: "*string", Required: true},
	{Name: "SectionPaths", Flag: "section-paths", Type: "[]string", Required: false},
}

var fields_describe_registration_type_definitions = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.RegistrationTypeFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RegistrationTypes", Flag: "registration-types", Type: "[]string", Required: false},
}

var fields_describe_registration_versions = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.RegistrationVersionFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RegistrationId", Flag: "registration-id", Type: "*string", Required: true},
	{Name: "VersionNumbers", Flag: "version-numbers", Type: "[]int64", Required: false},
}

var fields_describe_registrations = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.RegistrationFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RegistrationIds", Flag: "registration-ids", Type: "[]string", Required: false},
}

var fields_describe_sender_ids = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.SenderIdFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Owner", Flag: "owner", Type: "types.Owner", Required: false},
	{Name: "SenderIds", Flag: "sender-ids", Type: "[]types.SenderIdAndCountry", Required: false},
}

var fields_describe_spend_limits = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_verified_destination_numbers = []leanruntime.Field{
	{Name: "DestinationPhoneNumbers", Flag: "destination-phone-numbers", Type: "[]string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.VerifiedDestinationNumberFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "VerifiedDestinationNumberIds", Flag: "verified-destination-number-ids", Type: "[]string", Required: false},
}

var fields_disassociate_origination_identity = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "IsoCountryCode", Flag: "iso-country-code", Type: "*string", Required: true},
	{Name: "OriginationIdentity", Flag: "origination-identity", Type: "*string", Required: true},
	{Name: "PoolId", Flag: "pool-id", Type: "*string", Required: true},
}

var fields_disassociate_protect_configuration = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
	{Name: "ProtectConfigurationId", Flag: "protect-configuration-id", Type: "*string", Required: true},
}

var fields_discard_registration_version = []leanruntime.Field{
	{Name: "RegistrationId", Flag: "registration-id", Type: "*string", Required: true},
}

var fields_get_protect_configuration_country_rule_set = []leanruntime.Field{
	{Name: "NumberCapability", Flag: "number-capability", Type: "types.NumberCapability", Required: true},
	{Name: "ProtectConfigurationId", Flag: "protect-configuration-id", Type: "*string", Required: true},
}

var fields_get_resource_policy = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_pool_origination_identities = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.PoolOriginationIdentitiesFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PoolId", Flag: "pool-id", Type: "*string", Required: true},
}

var fields_list_protect_configuration_rule_set_number_overrides = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.ProtectConfigurationRuleSetNumberOverrideFilterItem", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProtectConfigurationId", Flag: "protect-configuration-id", Type: "*string", Required: true},
}

var fields_list_registration_associations = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.RegistrationAssociationFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RegistrationId", Flag: "registration-id", Type: "*string", Required: true},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_put_keyword = []leanruntime.Field{
	{Name: "Keyword", Flag: "keyword", Type: "*string", Required: true},
	{Name: "KeywordAction", Flag: "keyword-action", Type: "types.KeywordAction", Required: false},
	{Name: "KeywordMessage", Flag: "keyword-message", Type: "*string", Required: true},
	{Name: "OriginationIdentity", Flag: "origination-identity", Type: "*string", Required: true},
}

var fields_put_message_feedback = []leanruntime.Field{
	{Name: "MessageFeedbackStatus", Flag: "message-feedback-status", Type: "types.MessageFeedbackStatus", Required: true},
	{Name: "MessageId", Flag: "message-id", Type: "*string", Required: true},
}

var fields_put_opted_out_number = []leanruntime.Field{
	{Name: "OptOutListName", Flag: "opt-out-list-name", Type: "*string", Required: true},
	{Name: "OptedOutNumber", Flag: "opted-out-number", Type: "*string", Required: true},
}

var fields_put_protect_configuration_rule_set_number_override = []leanruntime.Field{
	{Name: "Action", Flag: "action", Type: "types.ProtectConfigurationRuleOverrideAction", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DestinationPhoneNumber", Flag: "destination-phone-number", Type: "*string", Required: true},
	{Name: "ExpirationTimestamp", Flag: "expiration-timestamp", Type: "*time.Time", Required: false},
	{Name: "ProtectConfigurationId", Flag: "protect-configuration-id", Type: "*string", Required: true},
}

var fields_put_registration_field_value = []leanruntime.Field{
	{Name: "FieldPath", Flag: "field-path", Type: "*string", Required: true},
	{Name: "RegistrationAttachmentId", Flag: "registration-attachment-id", Type: "*string", Required: false},
	{Name: "RegistrationId", Flag: "registration-id", Type: "*string", Required: true},
	{Name: "SelectChoices", Flag: "select-choices", Type: "[]string", Required: false},
	{Name: "TextValue", Flag: "text-value", Type: "*string", Required: false},
}

var fields_put_resource_policy = []leanruntime.Field{
	{Name: "Policy", Flag: "policy", Type: "*string", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_release_phone_number = []leanruntime.Field{
	{Name: "PhoneNumberId", Flag: "phone-number-id", Type: "*string", Required: true},
}

var fields_release_sender_id = []leanruntime.Field{
	{Name: "IsoCountryCode", Flag: "iso-country-code", Type: "*string", Required: true},
	{Name: "SenderId", Flag: "sender-id", Type: "*string", Required: true},
}

var fields_request_phone_number = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DeletionProtectionEnabled", Flag: "deletion-protection-enabled", Type: "*bool", Required: false},
	{Name: "InternationalSendingEnabled", Flag: "international-sending-enabled", Type: "*bool", Required: false},
	{Name: "IsoCountryCode", Flag: "iso-country-code", Type: "*string", Required: true},
	{Name: "MessageType", Flag: "message-type", Type: "types.MessageType", Required: true},
	{Name: "NumberCapabilities", Flag: "number-capabilities", Type: "[]types.NumberCapability", Required: true},
	{Name: "NumberType", Flag: "number-type", Type: "types.RequestableNumberType", Required: true},
	{Name: "OptOutListName", Flag: "opt-out-list-name", Type: "*string", Required: false},
	{Name: "PoolId", Flag: "pool-id", Type: "*string", Required: false},
	{Name: "RegistrationId", Flag: "registration-id", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_request_sender_id = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DeletionProtectionEnabled", Flag: "deletion-protection-enabled", Type: "*bool", Required: false},
	{Name: "IsoCountryCode", Flag: "iso-country-code", Type: "*string", Required: true},
	{Name: "MessageTypes", Flag: "message-types", Type: "[]types.MessageType", Required: false},
	{Name: "SenderId", Flag: "sender-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_send_destination_number_verification_code = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: false},
	{Name: "Context", Flag: "context", Type: "map[string]string", Required: false},
	{Name: "DestinationCountryParameters", Flag: "destination-country-parameters", Type: "map[string]string", Required: false},
	{Name: "LanguageCode", Flag: "language-code", Type: "types.LanguageCode", Required: false},
	{Name: "OriginationIdentity", Flag: "origination-identity", Type: "*string", Required: false},
	{Name: "VerificationChannel", Flag: "verification-channel", Type: "types.VerificationChannel", Required: true},
	{Name: "VerifiedDestinationNumberId", Flag: "verified-destination-number-id", Type: "*string", Required: true},
}

var fields_send_media_message = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: false},
	{Name: "Context", Flag: "context", Type: "map[string]string", Required: false},
	{Name: "DestinationPhoneNumber", Flag: "destination-phone-number", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "bool", Required: false},
	{Name: "MaxPrice", Flag: "max-price", Type: "*string", Required: false},
	{Name: "MediaUrls", Flag: "media-urls", Type: "[]string", Required: false},
	{Name: "MessageBody", Flag: "message-body", Type: "*string", Required: false},
	{Name: "MessageFeedbackEnabled", Flag: "message-feedback-enabled", Type: "*bool", Required: false},
	{Name: "OriginationIdentity", Flag: "origination-identity", Type: "*string", Required: true},
	{Name: "ProtectConfigurationId", Flag: "protect-configuration-id", Type: "*string", Required: false},
	{Name: "TimeToLive", Flag: "time-to-live", Type: "*int32", Required: false},
}

var fields_send_text_message = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: false},
	{Name: "Context", Flag: "context", Type: "map[string]string", Required: false},
	{Name: "DestinationCountryParameters", Flag: "destination-country-parameters", Type: "map[string]string", Required: false},
	{Name: "DestinationPhoneNumber", Flag: "destination-phone-number", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "bool", Required: false},
	{Name: "Keyword", Flag: "keyword", Type: "*string", Required: false},
	{Name: "MaxPrice", Flag: "max-price", Type: "*string", Required: false},
	{Name: "MessageBody", Flag: "message-body", Type: "*string", Required: false},
	{Name: "MessageFeedbackEnabled", Flag: "message-feedback-enabled", Type: "*bool", Required: false},
	{Name: "MessageType", Flag: "message-type", Type: "types.MessageType", Required: false},
	{Name: "OriginationIdentity", Flag: "origination-identity", Type: "*string", Required: false},
	{Name: "ProtectConfigurationId", Flag: "protect-configuration-id", Type: "*string", Required: false},
	{Name: "TimeToLive", Flag: "time-to-live", Type: "*int32", Required: false},
}

var fields_send_voice_message = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: false},
	{Name: "Context", Flag: "context", Type: "map[string]string", Required: false},
	{Name: "DestinationPhoneNumber", Flag: "destination-phone-number", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "bool", Required: false},
	{Name: "MaxPricePerMinute", Flag: "max-price-per-minute", Type: "*string", Required: false},
	{Name: "MessageBody", Flag: "message-body", Type: "*string", Required: false},
	{Name: "MessageBodyTextType", Flag: "message-body-text-type", Type: "types.VoiceMessageBodyTextType", Required: false},
	{Name: "MessageFeedbackEnabled", Flag: "message-feedback-enabled", Type: "*bool", Required: false},
	{Name: "OriginationIdentity", Flag: "origination-identity", Type: "*string", Required: true},
	{Name: "ProtectConfigurationId", Flag: "protect-configuration-id", Type: "*string", Required: false},
	{Name: "TimeToLive", Flag: "time-to-live", Type: "*int32", Required: false},
	{Name: "VoiceId", Flag: "voice-id", Type: "types.VoiceId", Required: false},
}

var fields_set_account_default_protect_configuration = []leanruntime.Field{
	{Name: "ProtectConfigurationId", Flag: "protect-configuration-id", Type: "*string", Required: true},
}

var fields_set_default_message_feedback_enabled = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
	{Name: "MessageFeedbackEnabled", Flag: "message-feedback-enabled", Type: "*bool", Required: true},
}

var fields_set_default_message_type = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
	{Name: "MessageType", Flag: "message-type", Type: "types.MessageType", Required: true},
}

var fields_set_default_sender_id = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
	{Name: "SenderId", Flag: "sender-id", Type: "*string", Required: true},
}

var fields_set_media_message_spend_limit_override = []leanruntime.Field{
	{Name: "MonthlyLimit", Flag: "monthly-limit", Type: "*int64", Required: true},
}

var fields_set_text_message_spend_limit_override = []leanruntime.Field{
	{Name: "MonthlyLimit", Flag: "monthly-limit", Type: "*int64", Required: true},
}

var fields_set_voice_message_spend_limit_override = []leanruntime.Field{
	{Name: "MonthlyLimit", Flag: "monthly-limit", Type: "*int64", Required: true},
}

var fields_submit_registration_version = []leanruntime.Field{
	{Name: "AwsReview", Flag: "aws-review", Type: "bool", Required: false},
	{Name: "RegistrationId", Flag: "registration-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_event_destination = []leanruntime.Field{
	{Name: "CloudWatchLogsDestination", Flag: "cloud-watch-logs-destination", Type: "*types.CloudWatchLogsDestination", Required: false},
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
	{Name: "Enabled", Flag: "enabled", Type: "*bool", Required: false},
	{Name: "EventDestinationName", Flag: "event-destination-name", Type: "*string", Required: true},
	{Name: "KinesisFirehoseDestination", Flag: "kinesis-firehose-destination", Type: "*types.KinesisFirehoseDestination", Required: false},
	{Name: "MatchingEventTypes", Flag: "matching-event-types", Type: "[]types.EventType", Required: false},
	{Name: "SnsDestination", Flag: "sns-destination", Type: "*types.SnsDestination", Required: false},
}

var fields_update_phone_number = []leanruntime.Field{
	{Name: "DeletionProtectionEnabled", Flag: "deletion-protection-enabled", Type: "*bool", Required: false},
	{Name: "InternationalSendingEnabled", Flag: "international-sending-enabled", Type: "*bool", Required: false},
	{Name: "OptOutListName", Flag: "opt-out-list-name", Type: "*string", Required: false},
	{Name: "PhoneNumberId", Flag: "phone-number-id", Type: "*string", Required: true},
	{Name: "SelfManagedOptOutsEnabled", Flag: "self-managed-opt-outs-enabled", Type: "*bool", Required: false},
	{Name: "TwoWayChannelArn", Flag: "two-way-channel-arn", Type: "*string", Required: false},
	{Name: "TwoWayChannelRole", Flag: "two-way-channel-role", Type: "*string", Required: false},
	{Name: "TwoWayEnabled", Flag: "two-way-enabled", Type: "*bool", Required: false},
}

var fields_update_pool = []leanruntime.Field{
	{Name: "DeletionProtectionEnabled", Flag: "deletion-protection-enabled", Type: "*bool", Required: false},
	{Name: "OptOutListName", Flag: "opt-out-list-name", Type: "*string", Required: false},
	{Name: "PoolId", Flag: "pool-id", Type: "*string", Required: true},
	{Name: "SelfManagedOptOutsEnabled", Flag: "self-managed-opt-outs-enabled", Type: "*bool", Required: false},
	{Name: "SharedRoutesEnabled", Flag: "shared-routes-enabled", Type: "*bool", Required: false},
	{Name: "TwoWayChannelArn", Flag: "two-way-channel-arn", Type: "*string", Required: false},
	{Name: "TwoWayChannelRole", Flag: "two-way-channel-role", Type: "*string", Required: false},
	{Name: "TwoWayEnabled", Flag: "two-way-enabled", Type: "*bool", Required: false},
}

var fields_update_protect_configuration = []leanruntime.Field{
	{Name: "DeletionProtectionEnabled", Flag: "deletion-protection-enabled", Type: "*bool", Required: false},
	{Name: "ProtectConfigurationId", Flag: "protect-configuration-id", Type: "*string", Required: true},
}

var fields_update_protect_configuration_country_rule_set = []leanruntime.Field{
	{Name: "CountryRuleSetUpdates", Flag: "country-rule-set-updates", Type: "map[string]types.ProtectConfigurationCountryRuleSetInformation", Required: true},
	{Name: "NumberCapability", Flag: "number-capability", Type: "types.NumberCapability", Required: true},
	{Name: "ProtectConfigurationId", Flag: "protect-configuration-id", Type: "*string", Required: true},
}

var fields_update_sender_id = []leanruntime.Field{
	{Name: "DeletionProtectionEnabled", Flag: "deletion-protection-enabled", Type: "*bool", Required: false},
	{Name: "IsoCountryCode", Flag: "iso-country-code", Type: "*string", Required: true},
	{Name: "SenderId", Flag: "sender-id", Type: "*string", Required: true},
}

var fields_verify_destination_number = []leanruntime.Field{
	{Name: "VerificationCode", Flag: "verification-code", Type: "*string", Required: true},
	{Name: "VerifiedDestinationNumberId", Flag: "verified-destination-number-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-origination-identity": {
			Name:   "associate-origination-identity",
			Fields: fields_associate_origination_identity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateOriginationIdentityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_origination_identity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateOriginationIdentity(ctx, input)
			},
		},
		"associate-protect-configuration": {
			Name:   "associate-protect-configuration",
			Fields: fields_associate_protect_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateProtectConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_protect_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateProtectConfiguration(ctx, input)
			},
		},
		"carrier-lookup": {
			Name:   "carrier-lookup",
			Fields: fields_carrier_lookup,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CarrierLookupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_carrier_lookup, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CarrierLookup(ctx, input)
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
		"create-event-destination": {
			Name:   "create-event-destination",
			Fields: fields_create_event_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEventDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_event_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEventDestination(ctx, input)
			},
		},
		"create-opt-out-list": {
			Name:   "create-opt-out-list",
			Fields: fields_create_opt_out_list,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateOptOutListInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_opt_out_list, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateOptOutList(ctx, input)
			},
		},
		"create-pool": {
			Name:   "create-pool",
			Fields: fields_create_pool,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePoolInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_pool, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePool(ctx, input)
			},
		},
		"create-protect-configuration": {
			Name:   "create-protect-configuration",
			Fields: fields_create_protect_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateProtectConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_protect_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateProtectConfiguration(ctx, input)
			},
		},
		"create-registration": {
			Name:   "create-registration",
			Fields: fields_create_registration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRegistrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_registration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRegistration(ctx, input)
			},
		},
		"create-registration-association": {
			Name:   "create-registration-association",
			Fields: fields_create_registration_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRegistrationAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_registration_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRegistrationAssociation(ctx, input)
			},
		},
		"create-registration-attachment": {
			Name:   "create-registration-attachment",
			Fields: fields_create_registration_attachment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRegistrationAttachmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_registration_attachment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRegistrationAttachment(ctx, input)
			},
		},
		"create-registration-version": {
			Name:   "create-registration-version",
			Fields: fields_create_registration_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRegistrationVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_registration_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRegistrationVersion(ctx, input)
			},
		},
		"create-verified-destination-number": {
			Name:   "create-verified-destination-number",
			Fields: fields_create_verified_destination_number,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVerifiedDestinationNumberInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_verified_destination_number, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVerifiedDestinationNumber(ctx, input)
			},
		},
		"delete-account-default-protect-configuration": {
			Name:   "delete-account-default-protect-configuration",
			Fields: fields_delete_account_default_protect_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAccountDefaultProtectConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_account_default_protect_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAccountDefaultProtectConfiguration(ctx, input)
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
		"delete-default-message-type": {
			Name:   "delete-default-message-type",
			Fields: fields_delete_default_message_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDefaultMessageTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_default_message_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDefaultMessageType(ctx, input)
			},
		},
		"delete-default-sender-id": {
			Name:   "delete-default-sender-id",
			Fields: fields_delete_default_sender_id,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDefaultSenderIdInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_default_sender_id, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDefaultSenderId(ctx, input)
			},
		},
		"delete-event-destination": {
			Name:   "delete-event-destination",
			Fields: fields_delete_event_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEventDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_event_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEventDestination(ctx, input)
			},
		},
		"delete-keyword": {
			Name:   "delete-keyword",
			Fields: fields_delete_keyword,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteKeywordInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_keyword, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteKeyword(ctx, input)
			},
		},
		"delete-media-message-spend-limit-override": {
			Name:   "delete-media-message-spend-limit-override",
			Fields: fields_delete_media_message_spend_limit_override,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMediaMessageSpendLimitOverrideInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_media_message_spend_limit_override, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMediaMessageSpendLimitOverride(ctx, input)
			},
		},
		"delete-opt-out-list": {
			Name:   "delete-opt-out-list",
			Fields: fields_delete_opt_out_list,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteOptOutListInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_opt_out_list, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteOptOutList(ctx, input)
			},
		},
		"delete-opted-out-number": {
			Name:   "delete-opted-out-number",
			Fields: fields_delete_opted_out_number,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteOptedOutNumberInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_opted_out_number, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteOptedOutNumber(ctx, input)
			},
		},
		"delete-pool": {
			Name:   "delete-pool",
			Fields: fields_delete_pool,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePoolInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_pool, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePool(ctx, input)
			},
		},
		"delete-protect-configuration": {
			Name:   "delete-protect-configuration",
			Fields: fields_delete_protect_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteProtectConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_protect_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteProtectConfiguration(ctx, input)
			},
		},
		"delete-protect-configuration-rule-set-number-override": {
			Name:   "delete-protect-configuration-rule-set-number-override",
			Fields: fields_delete_protect_configuration_rule_set_number_override,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteProtectConfigurationRuleSetNumberOverrideInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_protect_configuration_rule_set_number_override, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteProtectConfigurationRuleSetNumberOverride(ctx, input)
			},
		},
		"delete-registration": {
			Name:   "delete-registration",
			Fields: fields_delete_registration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRegistrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_registration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRegistration(ctx, input)
			},
		},
		"delete-registration-attachment": {
			Name:   "delete-registration-attachment",
			Fields: fields_delete_registration_attachment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRegistrationAttachmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_registration_attachment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRegistrationAttachment(ctx, input)
			},
		},
		"delete-registration-field-value": {
			Name:   "delete-registration-field-value",
			Fields: fields_delete_registration_field_value,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRegistrationFieldValueInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_registration_field_value, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRegistrationFieldValue(ctx, input)
			},
		},
		"delete-resource-policy": {
			Name:   "delete-resource-policy",
			Fields: fields_delete_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteResourcePolicy(ctx, input)
			},
		},
		"delete-text-message-spend-limit-override": {
			Name:   "delete-text-message-spend-limit-override",
			Fields: fields_delete_text_message_spend_limit_override,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTextMessageSpendLimitOverrideInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_text_message_spend_limit_override, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTextMessageSpendLimitOverride(ctx, input)
			},
		},
		"delete-verified-destination-number": {
			Name:   "delete-verified-destination-number",
			Fields: fields_delete_verified_destination_number,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVerifiedDestinationNumberInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_verified_destination_number, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVerifiedDestinationNumber(ctx, input)
			},
		},
		"delete-voice-message-spend-limit-override": {
			Name:   "delete-voice-message-spend-limit-override",
			Fields: fields_delete_voice_message_spend_limit_override,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVoiceMessageSpendLimitOverrideInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_voice_message_spend_limit_override, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVoiceMessageSpendLimitOverride(ctx, input)
			},
		},
		"describe-account-attributes": {
			Name:   "describe-account-attributes",
			Fields: fields_describe_account_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAccountAttributesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_account_attributes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeAccountAttributes(ctx, input)
				}
				var results []*svc.DescribeAccountAttributesOutput
				p := svc.NewDescribeAccountAttributesPaginator(client, input)
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
		"describe-account-limits": {
			Name:   "describe-account-limits",
			Fields: fields_describe_account_limits,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAccountLimitsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_account_limits, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeAccountLimits(ctx, input)
				}
				var results []*svc.DescribeAccountLimitsOutput
				p := svc.NewDescribeAccountLimitsPaginator(client, input)
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
		"describe-configuration-sets": {
			Name:   "describe-configuration-sets",
			Fields: fields_describe_configuration_sets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeConfigurationSetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_configuration_sets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeConfigurationSets(ctx, input)
				}
				var results []*svc.DescribeConfigurationSetsOutput
				p := svc.NewDescribeConfigurationSetsPaginator(client, input)
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
		"describe-keywords": {
			Name:   "describe-keywords",
			Fields: fields_describe_keywords,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeKeywordsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_keywords, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeKeywords(ctx, input)
				}
				var results []*svc.DescribeKeywordsOutput
				p := svc.NewDescribeKeywordsPaginator(client, input)
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
		"describe-opt-out-lists": {
			Name:   "describe-opt-out-lists",
			Fields: fields_describe_opt_out_lists,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeOptOutListsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_opt_out_lists, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeOptOutLists(ctx, input)
				}
				var results []*svc.DescribeOptOutListsOutput
				p := svc.NewDescribeOptOutListsPaginator(client, input)
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
		"describe-opted-out-numbers": {
			Name:   "describe-opted-out-numbers",
			Fields: fields_describe_opted_out_numbers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeOptedOutNumbersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_opted_out_numbers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeOptedOutNumbers(ctx, input)
				}
				var results []*svc.DescribeOptedOutNumbersOutput
				p := svc.NewDescribeOptedOutNumbersPaginator(client, input)
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
		"describe-phone-numbers": {
			Name:   "describe-phone-numbers",
			Fields: fields_describe_phone_numbers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePhoneNumbersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_phone_numbers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribePhoneNumbers(ctx, input)
				}
				var results []*svc.DescribePhoneNumbersOutput
				p := svc.NewDescribePhoneNumbersPaginator(client, input)
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
		"describe-pools": {
			Name:   "describe-pools",
			Fields: fields_describe_pools,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePoolsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_pools, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribePools(ctx, input)
				}
				var results []*svc.DescribePoolsOutput
				p := svc.NewDescribePoolsPaginator(client, input)
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
		"describe-protect-configurations": {
			Name:   "describe-protect-configurations",
			Fields: fields_describe_protect_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeProtectConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_protect_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeProtectConfigurations(ctx, input)
				}
				var results []*svc.DescribeProtectConfigurationsOutput
				p := svc.NewDescribeProtectConfigurationsPaginator(client, input)
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
		"describe-registration-attachments": {
			Name:   "describe-registration-attachments",
			Fields: fields_describe_registration_attachments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRegistrationAttachmentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_registration_attachments, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeRegistrationAttachments(ctx, input)
				}
				var results []*svc.DescribeRegistrationAttachmentsOutput
				p := svc.NewDescribeRegistrationAttachmentsPaginator(client, input)
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
		"describe-registration-field-definitions": {
			Name:   "describe-registration-field-definitions",
			Fields: fields_describe_registration_field_definitions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRegistrationFieldDefinitionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_registration_field_definitions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeRegistrationFieldDefinitions(ctx, input)
				}
				var results []*svc.DescribeRegistrationFieldDefinitionsOutput
				p := svc.NewDescribeRegistrationFieldDefinitionsPaginator(client, input)
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
		"describe-registration-field-values": {
			Name:   "describe-registration-field-values",
			Fields: fields_describe_registration_field_values,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRegistrationFieldValuesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_registration_field_values, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeRegistrationFieldValues(ctx, input)
				}
				var results []*svc.DescribeRegistrationFieldValuesOutput
				p := svc.NewDescribeRegistrationFieldValuesPaginator(client, input)
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
		"describe-registration-section-definitions": {
			Name:   "describe-registration-section-definitions",
			Fields: fields_describe_registration_section_definitions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRegistrationSectionDefinitionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_registration_section_definitions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeRegistrationSectionDefinitions(ctx, input)
				}
				var results []*svc.DescribeRegistrationSectionDefinitionsOutput
				p := svc.NewDescribeRegistrationSectionDefinitionsPaginator(client, input)
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
		"describe-registration-type-definitions": {
			Name:   "describe-registration-type-definitions",
			Fields: fields_describe_registration_type_definitions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRegistrationTypeDefinitionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_registration_type_definitions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeRegistrationTypeDefinitions(ctx, input)
				}
				var results []*svc.DescribeRegistrationTypeDefinitionsOutput
				p := svc.NewDescribeRegistrationTypeDefinitionsPaginator(client, input)
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
		"describe-registration-versions": {
			Name:   "describe-registration-versions",
			Fields: fields_describe_registration_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRegistrationVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_registration_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeRegistrationVersions(ctx, input)
				}
				var results []*svc.DescribeRegistrationVersionsOutput
				p := svc.NewDescribeRegistrationVersionsPaginator(client, input)
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
		"describe-registrations": {
			Name:   "describe-registrations",
			Fields: fields_describe_registrations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRegistrationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_registrations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeRegistrations(ctx, input)
				}
				var results []*svc.DescribeRegistrationsOutput
				p := svc.NewDescribeRegistrationsPaginator(client, input)
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
		"describe-sender-ids": {
			Name:   "describe-sender-ids",
			Fields: fields_describe_sender_ids,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSenderIdsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_sender_ids, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeSenderIds(ctx, input)
				}
				var results []*svc.DescribeSenderIdsOutput
				p := svc.NewDescribeSenderIdsPaginator(client, input)
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
		"describe-spend-limits": {
			Name:   "describe-spend-limits",
			Fields: fields_describe_spend_limits,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSpendLimitsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_spend_limits, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeSpendLimits(ctx, input)
				}
				var results []*svc.DescribeSpendLimitsOutput
				p := svc.NewDescribeSpendLimitsPaginator(client, input)
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
		"describe-verified-destination-numbers": {
			Name:   "describe-verified-destination-numbers",
			Fields: fields_describe_verified_destination_numbers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVerifiedDestinationNumbersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_verified_destination_numbers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeVerifiedDestinationNumbers(ctx, input)
				}
				var results []*svc.DescribeVerifiedDestinationNumbersOutput
				p := svc.NewDescribeVerifiedDestinationNumbersPaginator(client, input)
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
		"disassociate-origination-identity": {
			Name:   "disassociate-origination-identity",
			Fields: fields_disassociate_origination_identity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateOriginationIdentityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_origination_identity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateOriginationIdentity(ctx, input)
			},
		},
		"disassociate-protect-configuration": {
			Name:   "disassociate-protect-configuration",
			Fields: fields_disassociate_protect_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateProtectConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_protect_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateProtectConfiguration(ctx, input)
			},
		},
		"discard-registration-version": {
			Name:   "discard-registration-version",
			Fields: fields_discard_registration_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DiscardRegistrationVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_discard_registration_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DiscardRegistrationVersion(ctx, input)
			},
		},
		"get-protect-configuration-country-rule-set": {
			Name:   "get-protect-configuration-country-rule-set",
			Fields: fields_get_protect_configuration_country_rule_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetProtectConfigurationCountryRuleSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_protect_configuration_country_rule_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetProtectConfigurationCountryRuleSet(ctx, input)
			},
		},
		"get-resource-policy": {
			Name:   "get-resource-policy",
			Fields: fields_get_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResourcePolicy(ctx, input)
			},
		},
		"list-pool-origination-identities": {
			Name:   "list-pool-origination-identities",
			Fields: fields_list_pool_origination_identities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPoolOriginationIdentitiesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_pool_origination_identities, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPoolOriginationIdentities(ctx, input)
				}
				var results []*svc.ListPoolOriginationIdentitiesOutput
				p := svc.NewListPoolOriginationIdentitiesPaginator(client, input)
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
		"list-protect-configuration-rule-set-number-overrides": {
			Name:   "list-protect-configuration-rule-set-number-overrides",
			Fields: fields_list_protect_configuration_rule_set_number_overrides,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProtectConfigurationRuleSetNumberOverridesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_protect_configuration_rule_set_number_overrides, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProtectConfigurationRuleSetNumberOverrides(ctx, input)
				}
				var results []*svc.ListProtectConfigurationRuleSetNumberOverridesOutput
				p := svc.NewListProtectConfigurationRuleSetNumberOverridesPaginator(client, input)
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
		"list-registration-associations": {
			Name:   "list-registration-associations",
			Fields: fields_list_registration_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRegistrationAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_registration_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRegistrationAssociations(ctx, input)
				}
				var results []*svc.ListRegistrationAssociationsOutput
				p := svc.NewListRegistrationAssociationsPaginator(client, input)
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
		"put-keyword": {
			Name:   "put-keyword",
			Fields: fields_put_keyword,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutKeywordInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_keyword, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutKeyword(ctx, input)
			},
		},
		"put-message-feedback": {
			Name:   "put-message-feedback",
			Fields: fields_put_message_feedback,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutMessageFeedbackInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_message_feedback, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutMessageFeedback(ctx, input)
			},
		},
		"put-opted-out-number": {
			Name:   "put-opted-out-number",
			Fields: fields_put_opted_out_number,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutOptedOutNumberInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_opted_out_number, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutOptedOutNumber(ctx, input)
			},
		},
		"put-protect-configuration-rule-set-number-override": {
			Name:   "put-protect-configuration-rule-set-number-override",
			Fields: fields_put_protect_configuration_rule_set_number_override,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutProtectConfigurationRuleSetNumberOverrideInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_protect_configuration_rule_set_number_override, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutProtectConfigurationRuleSetNumberOverride(ctx, input)
			},
		},
		"put-registration-field-value": {
			Name:   "put-registration-field-value",
			Fields: fields_put_registration_field_value,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutRegistrationFieldValueInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_registration_field_value, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutRegistrationFieldValue(ctx, input)
			},
		},
		"put-resource-policy": {
			Name:   "put-resource-policy",
			Fields: fields_put_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutResourcePolicy(ctx, input)
			},
		},
		"release-phone-number": {
			Name:   "release-phone-number",
			Fields: fields_release_phone_number,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ReleasePhoneNumberInput{}
				if _, err := leanruntime.ApplyInput(input, fields_release_phone_number, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ReleasePhoneNumber(ctx, input)
			},
		},
		"release-sender-id": {
			Name:   "release-sender-id",
			Fields: fields_release_sender_id,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ReleaseSenderIdInput{}
				if _, err := leanruntime.ApplyInput(input, fields_release_sender_id, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ReleaseSenderId(ctx, input)
			},
		},
		"request-phone-number": {
			Name:   "request-phone-number",
			Fields: fields_request_phone_number,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RequestPhoneNumberInput{}
				if _, err := leanruntime.ApplyInput(input, fields_request_phone_number, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RequestPhoneNumber(ctx, input)
			},
		},
		"request-sender-id": {
			Name:   "request-sender-id",
			Fields: fields_request_sender_id,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RequestSenderIdInput{}
				if _, err := leanruntime.ApplyInput(input, fields_request_sender_id, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RequestSenderId(ctx, input)
			},
		},
		"send-destination-number-verification-code": {
			Name:   "send-destination-number-verification-code",
			Fields: fields_send_destination_number_verification_code,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendDestinationNumberVerificationCodeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_destination_number_verification_code, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendDestinationNumberVerificationCode(ctx, input)
			},
		},
		"send-media-message": {
			Name:   "send-media-message",
			Fields: fields_send_media_message,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendMediaMessageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_media_message, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendMediaMessage(ctx, input)
			},
		},
		"send-text-message": {
			Name:   "send-text-message",
			Fields: fields_send_text_message,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendTextMessageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_text_message, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendTextMessage(ctx, input)
			},
		},
		"send-voice-message": {
			Name:   "send-voice-message",
			Fields: fields_send_voice_message,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendVoiceMessageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_voice_message, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendVoiceMessage(ctx, input)
			},
		},
		"set-account-default-protect-configuration": {
			Name:   "set-account-default-protect-configuration",
			Fields: fields_set_account_default_protect_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetAccountDefaultProtectConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_account_default_protect_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetAccountDefaultProtectConfiguration(ctx, input)
			},
		},
		"set-default-message-feedback-enabled": {
			Name:   "set-default-message-feedback-enabled",
			Fields: fields_set_default_message_feedback_enabled,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetDefaultMessageFeedbackEnabledInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_default_message_feedback_enabled, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetDefaultMessageFeedbackEnabled(ctx, input)
			},
		},
		"set-default-message-type": {
			Name:   "set-default-message-type",
			Fields: fields_set_default_message_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetDefaultMessageTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_default_message_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetDefaultMessageType(ctx, input)
			},
		},
		"set-default-sender-id": {
			Name:   "set-default-sender-id",
			Fields: fields_set_default_sender_id,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetDefaultSenderIdInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_default_sender_id, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetDefaultSenderId(ctx, input)
			},
		},
		"set-media-message-spend-limit-override": {
			Name:   "set-media-message-spend-limit-override",
			Fields: fields_set_media_message_spend_limit_override,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetMediaMessageSpendLimitOverrideInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_media_message_spend_limit_override, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetMediaMessageSpendLimitOverride(ctx, input)
			},
		},
		"set-text-message-spend-limit-override": {
			Name:   "set-text-message-spend-limit-override",
			Fields: fields_set_text_message_spend_limit_override,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetTextMessageSpendLimitOverrideInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_text_message_spend_limit_override, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetTextMessageSpendLimitOverride(ctx, input)
			},
		},
		"set-voice-message-spend-limit-override": {
			Name:   "set-voice-message-spend-limit-override",
			Fields: fields_set_voice_message_spend_limit_override,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetVoiceMessageSpendLimitOverrideInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_voice_message_spend_limit_override, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetVoiceMessageSpendLimitOverride(ctx, input)
			},
		},
		"submit-registration-version": {
			Name:   "submit-registration-version",
			Fields: fields_submit_registration_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SubmitRegistrationVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_submit_registration_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SubmitRegistrationVersion(ctx, input)
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
		"update-event-destination": {
			Name:   "update-event-destination",
			Fields: fields_update_event_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEventDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_event_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEventDestination(ctx, input)
			},
		},
		"update-phone-number": {
			Name:   "update-phone-number",
			Fields: fields_update_phone_number,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePhoneNumberInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_phone_number, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePhoneNumber(ctx, input)
			},
		},
		"update-pool": {
			Name:   "update-pool",
			Fields: fields_update_pool,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePoolInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_pool, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePool(ctx, input)
			},
		},
		"update-protect-configuration": {
			Name:   "update-protect-configuration",
			Fields: fields_update_protect_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateProtectConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_protect_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateProtectConfiguration(ctx, input)
			},
		},
		"update-protect-configuration-country-rule-set": {
			Name:   "update-protect-configuration-country-rule-set",
			Fields: fields_update_protect_configuration_country_rule_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateProtectConfigurationCountryRuleSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_protect_configuration_country_rule_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateProtectConfigurationCountryRuleSet(ctx, input)
			},
		},
		"update-sender-id": {
			Name:   "update-sender-id",
			Fields: fields_update_sender_id,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSenderIdInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_sender_id, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSenderId(ctx, input)
			},
		},
		"verify-destination-number": {
			Name:   "verify-destination-number",
			Fields: fields_verify_destination_number,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.VerifyDestinationNumberInput{}
				if _, err := leanruntime.ApplyInput(input, fields_verify_destination_number, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.VerifyDestinationNumber(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("pinpointsmsvoicev2", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
