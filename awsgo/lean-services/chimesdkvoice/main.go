package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/chimesdkvoice"
)

var fields_associate_phone_numbers_with_voice_connector = []leanruntime.Field{
	{Name: "E164PhoneNumbers", Flag: "e164-phone-numbers", Type: "[]string", Required: true},
	{Name: "ForceAssociate", Flag: "force-associate", Type: "*bool", Required: false},
	{Name: "VoiceConnectorId", Flag: "voice-connector-id", Type: "*string", Required: true},
}

var fields_associate_phone_numbers_with_voice_connector_group = []leanruntime.Field{
	{Name: "E164PhoneNumbers", Flag: "e164-phone-numbers", Type: "[]string", Required: true},
	{Name: "ForceAssociate", Flag: "force-associate", Type: "*bool", Required: false},
	{Name: "VoiceConnectorGroupId", Flag: "voice-connector-group-id", Type: "*string", Required: true},
}

var fields_batch_delete_phone_number = []leanruntime.Field{
	{Name: "PhoneNumberIds", Flag: "phone-number-ids", Type: "[]string", Required: true},
}

var fields_batch_update_phone_number = []leanruntime.Field{
	{Name: "UpdatePhoneNumberRequestItems", Flag: "update-phone-number-request-items", Type: "[]types.UpdatePhoneNumberRequestItem", Required: true},
}

var fields_create_phone_number_order = []leanruntime.Field{
	{Name: "E164PhoneNumbers", Flag: "e164-phone-numbers", Type: "[]string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "ProductType", Flag: "product-type", Type: "types.PhoneNumberProductType", Required: true},
}

var fields_create_proxy_session = []leanruntime.Field{
	{Name: "Capabilities", Flag: "capabilities", Type: "[]types.Capability", Required: true},
	{Name: "ExpiryMinutes", Flag: "expiry-minutes", Type: "*int32", Required: false},
	{Name: "GeoMatchLevel", Flag: "geo-match-level", Type: "types.GeoMatchLevel", Required: false},
	{Name: "GeoMatchParams", Flag: "geo-match-params", Type: "*types.GeoMatchParams", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "NumberSelectionBehavior", Flag: "number-selection-behavior", Type: "types.NumberSelectionBehavior", Required: false},
	{Name: "ParticipantPhoneNumbers", Flag: "participant-phone-numbers", Type: "[]string", Required: true},
	{Name: "VoiceConnectorId", Flag: "voice-connector-id", Type: "*string", Required: true},
}

var fields_create_sip_media_application = []leanruntime.Field{
	{Name: "AwsRegion", Flag: "aws-region", Type: "*string", Required: true},
	{Name: "Endpoints", Flag: "endpoints", Type: "[]types.SipMediaApplicationEndpoint", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_sip_media_application_call = []leanruntime.Field{
	{Name: "ArgumentsMap", Flag: "arguments-map", Type: "map[string]string", Required: false},
	{Name: "FromPhoneNumber", Flag: "from-phone-number", Type: "*string", Required: true},
	{Name: "SipHeaders", Flag: "sip-headers", Type: "map[string]string", Required: false},
	{Name: "SipMediaApplicationId", Flag: "sip-media-application-id", Type: "*string", Required: true},
	{Name: "ToPhoneNumber", Flag: "to-phone-number", Type: "*string", Required: true},
}

var fields_create_sip_rule = []leanruntime.Field{
	{Name: "Disabled", Flag: "disabled", Type: "*bool", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "TargetApplications", Flag: "target-applications", Type: "[]types.SipRuleTargetApplication", Required: false},
	{Name: "TriggerType", Flag: "trigger-type", Type: "types.SipRuleTriggerType", Required: true},
	{Name: "TriggerValue", Flag: "trigger-value", Type: "*string", Required: true},
}

var fields_create_voice_connector = []leanruntime.Field{
	{Name: "AwsRegion", Flag: "aws-region", Type: "types.VoiceConnectorAwsRegion", Required: false},
	{Name: "IntegrationType", Flag: "integration-type", Type: "types.VoiceConnectorIntegrationType", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NetworkType", Flag: "network-type", Type: "types.NetworkType", Required: false},
	{Name: "RequireEncryption", Flag: "require-encryption", Type: "*bool", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_voice_connector_group = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "VoiceConnectorItems", Flag: "voice-connector-items", Type: "[]types.VoiceConnectorItem", Required: false},
}

var fields_create_voice_profile = []leanruntime.Field{
	{Name: "SpeakerSearchTaskId", Flag: "speaker-search-task-id", Type: "*string", Required: true},
}

var fields_create_voice_profile_domain = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ServerSideEncryptionConfiguration", Flag: "server-side-encryption-configuration", Type: "*types.ServerSideEncryptionConfiguration", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_delete_phone_number = []leanruntime.Field{
	{Name: "PhoneNumberId", Flag: "phone-number-id", Type: "*string", Required: true},
}

var fields_delete_proxy_session = []leanruntime.Field{
	{Name: "ProxySessionId", Flag: "proxy-session-id", Type: "*string", Required: true},
	{Name: "VoiceConnectorId", Flag: "voice-connector-id", Type: "*string", Required: true},
}

var fields_delete_sip_media_application = []leanruntime.Field{
	{Name: "SipMediaApplicationId", Flag: "sip-media-application-id", Type: "*string", Required: true},
}

var fields_delete_sip_rule = []leanruntime.Field{
	{Name: "SipRuleId", Flag: "sip-rule-id", Type: "*string", Required: true},
}

var fields_delete_voice_connector = []leanruntime.Field{
	{Name: "VoiceConnectorId", Flag: "voice-connector-id", Type: "*string", Required: true},
}

var fields_delete_voice_connector_emergency_calling_configuration = []leanruntime.Field{
	{Name: "VoiceConnectorId", Flag: "voice-connector-id", Type: "*string", Required: true},
}

var fields_delete_voice_connector_external_systems_configuration = []leanruntime.Field{
	{Name: "VoiceConnectorId", Flag: "voice-connector-id", Type: "*string", Required: true},
}

var fields_delete_voice_connector_group = []leanruntime.Field{
	{Name: "VoiceConnectorGroupId", Flag: "voice-connector-group-id", Type: "*string", Required: true},
}

var fields_delete_voice_connector_origination = []leanruntime.Field{
	{Name: "VoiceConnectorId", Flag: "voice-connector-id", Type: "*string", Required: true},
}

var fields_delete_voice_connector_proxy = []leanruntime.Field{
	{Name: "VoiceConnectorId", Flag: "voice-connector-id", Type: "*string", Required: true},
}

var fields_delete_voice_connector_streaming_configuration = []leanruntime.Field{
	{Name: "VoiceConnectorId", Flag: "voice-connector-id", Type: "*string", Required: true},
}

var fields_delete_voice_connector_termination = []leanruntime.Field{
	{Name: "VoiceConnectorId", Flag: "voice-connector-id", Type: "*string", Required: true},
}

var fields_delete_voice_connector_termination_credentials = []leanruntime.Field{
	{Name: "Usernames", Flag: "usernames", Type: "[]string", Required: true},
	{Name: "VoiceConnectorId", Flag: "voice-connector-id", Type: "*string", Required: true},
}

var fields_delete_voice_profile = []leanruntime.Field{
	{Name: "VoiceProfileId", Flag: "voice-profile-id", Type: "*string", Required: true},
}

var fields_delete_voice_profile_domain = []leanruntime.Field{
	{Name: "VoiceProfileDomainId", Flag: "voice-profile-domain-id", Type: "*string", Required: true},
}

var fields_disassociate_phone_numbers_from_voice_connector = []leanruntime.Field{
	{Name: "E164PhoneNumbers", Flag: "e164-phone-numbers", Type: "[]string", Required: true},
	{Name: "VoiceConnectorId", Flag: "voice-connector-id", Type: "*string", Required: true},
}

var fields_disassociate_phone_numbers_from_voice_connector_group = []leanruntime.Field{
	{Name: "E164PhoneNumbers", Flag: "e164-phone-numbers", Type: "[]string", Required: true},
	{Name: "VoiceConnectorGroupId", Flag: "voice-connector-group-id", Type: "*string", Required: true},
}

var fields_get_global_settings = []leanruntime.Field{}

var fields_get_phone_number = []leanruntime.Field{
	{Name: "PhoneNumberId", Flag: "phone-number-id", Type: "*string", Required: true},
}

var fields_get_phone_number_order = []leanruntime.Field{
	{Name: "PhoneNumberOrderId", Flag: "phone-number-order-id", Type: "*string", Required: true},
}

var fields_get_phone_number_settings = []leanruntime.Field{}

var fields_get_proxy_session = []leanruntime.Field{
	{Name: "ProxySessionId", Flag: "proxy-session-id", Type: "*string", Required: true},
	{Name: "VoiceConnectorId", Flag: "voice-connector-id", Type: "*string", Required: true},
}

var fields_get_sip_media_application = []leanruntime.Field{
	{Name: "SipMediaApplicationId", Flag: "sip-media-application-id", Type: "*string", Required: true},
}

var fields_get_sip_media_application_alexa_skill_configuration = []leanruntime.Field{
	{Name: "SipMediaApplicationId", Flag: "sip-media-application-id", Type: "*string", Required: true},
}

var fields_get_sip_media_application_logging_configuration = []leanruntime.Field{
	{Name: "SipMediaApplicationId", Flag: "sip-media-application-id", Type: "*string", Required: true},
}

var fields_get_sip_rule = []leanruntime.Field{
	{Name: "SipRuleId", Flag: "sip-rule-id", Type: "*string", Required: true},
}

var fields_get_speaker_search_task = []leanruntime.Field{
	{Name: "SpeakerSearchTaskId", Flag: "speaker-search-task-id", Type: "*string", Required: true},
	{Name: "VoiceConnectorId", Flag: "voice-connector-id", Type: "*string", Required: true},
}

var fields_get_voice_connector = []leanruntime.Field{
	{Name: "VoiceConnectorId", Flag: "voice-connector-id", Type: "*string", Required: true},
}

var fields_get_voice_connector_emergency_calling_configuration = []leanruntime.Field{
	{Name: "VoiceConnectorId", Flag: "voice-connector-id", Type: "*string", Required: true},
}

var fields_get_voice_connector_external_systems_configuration = []leanruntime.Field{
	{Name: "VoiceConnectorId", Flag: "voice-connector-id", Type: "*string", Required: true},
}

var fields_get_voice_connector_group = []leanruntime.Field{
	{Name: "VoiceConnectorGroupId", Flag: "voice-connector-group-id", Type: "*string", Required: true},
}

var fields_get_voice_connector_logging_configuration = []leanruntime.Field{
	{Name: "VoiceConnectorId", Flag: "voice-connector-id", Type: "*string", Required: true},
}

var fields_get_voice_connector_origination = []leanruntime.Field{
	{Name: "VoiceConnectorId", Flag: "voice-connector-id", Type: "*string", Required: true},
}

var fields_get_voice_connector_proxy = []leanruntime.Field{
	{Name: "VoiceConnectorId", Flag: "voice-connector-id", Type: "*string", Required: true},
}

var fields_get_voice_connector_streaming_configuration = []leanruntime.Field{
	{Name: "VoiceConnectorId", Flag: "voice-connector-id", Type: "*string", Required: true},
}

var fields_get_voice_connector_termination = []leanruntime.Field{
	{Name: "VoiceConnectorId", Flag: "voice-connector-id", Type: "*string", Required: true},
}

var fields_get_voice_connector_termination_health = []leanruntime.Field{
	{Name: "VoiceConnectorId", Flag: "voice-connector-id", Type: "*string", Required: true},
}

var fields_get_voice_profile = []leanruntime.Field{
	{Name: "VoiceProfileId", Flag: "voice-profile-id", Type: "*string", Required: true},
}

var fields_get_voice_profile_domain = []leanruntime.Field{
	{Name: "VoiceProfileDomainId", Flag: "voice-profile-domain-id", Type: "*string", Required: true},
}

var fields_get_voice_tone_analysis_task = []leanruntime.Field{
	{Name: "IsCaller", Flag: "is-caller", Type: "*bool", Required: true},
	{Name: "VoiceConnectorId", Flag: "voice-connector-id", Type: "*string", Required: true},
	{Name: "VoiceToneAnalysisTaskId", Flag: "voice-tone-analysis-task-id", Type: "*string", Required: true},
}

var fields_list_available_voice_connector_regions = []leanruntime.Field{}

var fields_list_phone_number_orders = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_phone_numbers = []leanruntime.Field{
	{Name: "FilterName", Flag: "filter-name", Type: "types.PhoneNumberAssociationName", Required: false},
	{Name: "FilterValue", Flag: "filter-value", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProductType", Flag: "product-type", Type: "types.PhoneNumberProductType", Required: false},
	{Name: "Status", Flag: "status", Type: "*string", Required: false},
}

var fields_list_proxy_sessions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.ProxySessionStatus", Required: false},
	{Name: "VoiceConnectorId", Flag: "voice-connector-id", Type: "*string", Required: true},
}

var fields_list_sip_media_applications = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_sip_rules = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SipMediaApplicationId", Flag: "sip-media-application-id", Type: "*string", Required: false},
}

var fields_list_supported_phone_number_countries = []leanruntime.Field{
	{Name: "ProductType", Flag: "product-type", Type: "types.PhoneNumberProductType", Required: true},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_voice_connector_groups = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_voice_connector_termination_credentials = []leanruntime.Field{
	{Name: "VoiceConnectorId", Flag: "voice-connector-id", Type: "*string", Required: true},
}

var fields_list_voice_connectors = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_voice_profile_domains = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_voice_profiles = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "VoiceProfileDomainId", Flag: "voice-profile-domain-id", Type: "*string", Required: true},
}

var fields_put_sip_media_application_alexa_skill_configuration = []leanruntime.Field{
	{Name: "SipMediaApplicationAlexaSkillConfiguration", Flag: "sip-media-application-alexa-skill-configuration", Type: "*types.SipMediaApplicationAlexaSkillConfiguration", Required: false},
	{Name: "SipMediaApplicationId", Flag: "sip-media-application-id", Type: "*string", Required: true},
}

var fields_put_sip_media_application_logging_configuration = []leanruntime.Field{
	{Name: "SipMediaApplicationId", Flag: "sip-media-application-id", Type: "*string", Required: true},
	{Name: "SipMediaApplicationLoggingConfiguration", Flag: "sip-media-application-logging-configuration", Type: "*types.SipMediaApplicationLoggingConfiguration", Required: false},
}

var fields_put_voice_connector_emergency_calling_configuration = []leanruntime.Field{
	{Name: "EmergencyCallingConfiguration", Flag: "emergency-calling-configuration", Type: "*types.EmergencyCallingConfiguration", Required: true},
	{Name: "VoiceConnectorId", Flag: "voice-connector-id", Type: "*string", Required: true},
}

var fields_put_voice_connector_external_systems_configuration = []leanruntime.Field{
	{Name: "ContactCenterSystemTypes", Flag: "contact-center-system-types", Type: "[]types.ContactCenterSystemType", Required: false},
	{Name: "SessionBorderControllerTypes", Flag: "session-border-controller-types", Type: "[]types.SessionBorderControllerType", Required: false},
	{Name: "VoiceConnectorId", Flag: "voice-connector-id", Type: "*string", Required: true},
}

var fields_put_voice_connector_logging_configuration = []leanruntime.Field{
	{Name: "LoggingConfiguration", Flag: "logging-configuration", Type: "*types.LoggingConfiguration", Required: true},
	{Name: "VoiceConnectorId", Flag: "voice-connector-id", Type: "*string", Required: true},
}

var fields_put_voice_connector_origination = []leanruntime.Field{
	{Name: "Origination", Flag: "origination", Type: "*types.Origination", Required: true},
	{Name: "VoiceConnectorId", Flag: "voice-connector-id", Type: "*string", Required: true},
}

var fields_put_voice_connector_proxy = []leanruntime.Field{
	{Name: "DefaultSessionExpiryMinutes", Flag: "default-session-expiry-minutes", Type: "*int32", Required: true},
	{Name: "Disabled", Flag: "disabled", Type: "*bool", Required: false},
	{Name: "FallBackPhoneNumber", Flag: "fall-back-phone-number", Type: "*string", Required: false},
	{Name: "PhoneNumberPoolCountries", Flag: "phone-number-pool-countries", Type: "[]string", Required: true},
	{Name: "VoiceConnectorId", Flag: "voice-connector-id", Type: "*string", Required: true},
}

var fields_put_voice_connector_streaming_configuration = []leanruntime.Field{
	{Name: "StreamingConfiguration", Flag: "streaming-configuration", Type: "*types.StreamingConfiguration", Required: true},
	{Name: "VoiceConnectorId", Flag: "voice-connector-id", Type: "*string", Required: true},
}

var fields_put_voice_connector_termination = []leanruntime.Field{
	{Name: "Termination", Flag: "termination", Type: "*types.Termination", Required: true},
	{Name: "VoiceConnectorId", Flag: "voice-connector-id", Type: "*string", Required: true},
}

var fields_put_voice_connector_termination_credentials = []leanruntime.Field{
	{Name: "Credentials", Flag: "credentials", Type: "[]types.Credential", Required: false},
	{Name: "VoiceConnectorId", Flag: "voice-connector-id", Type: "*string", Required: true},
}

var fields_restore_phone_number = []leanruntime.Field{
	{Name: "PhoneNumberId", Flag: "phone-number-id", Type: "*string", Required: true},
}

var fields_search_available_phone_numbers = []leanruntime.Field{
	{Name: "AreaCode", Flag: "area-code", Type: "*string", Required: false},
	{Name: "City", Flag: "city", Type: "*string", Required: false},
	{Name: "Country", Flag: "country", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PhoneNumberType", Flag: "phone-number-type", Type: "types.PhoneNumberType", Required: false},
	{Name: "State", Flag: "state", Type: "*string", Required: false},
	{Name: "TollFreePrefix", Flag: "toll-free-prefix", Type: "*string", Required: false},
}

var fields_start_speaker_search_task = []leanruntime.Field{
	{Name: "CallLeg", Flag: "call-leg", Type: "types.CallLegType", Required: false},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "TransactionId", Flag: "transaction-id", Type: "*string", Required: true},
	{Name: "VoiceConnectorId", Flag: "voice-connector-id", Type: "*string", Required: true},
	{Name: "VoiceProfileDomainId", Flag: "voice-profile-domain-id", Type: "*string", Required: true},
}

var fields_start_voice_tone_analysis_task = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "LanguageCode", Flag: "language-code", Type: "types.LanguageCode", Required: true},
	{Name: "TransactionId", Flag: "transaction-id", Type: "*string", Required: true},
	{Name: "VoiceConnectorId", Flag: "voice-connector-id", Type: "*string", Required: true},
}

var fields_stop_speaker_search_task = []leanruntime.Field{
	{Name: "SpeakerSearchTaskId", Flag: "speaker-search-task-id", Type: "*string", Required: true},
	{Name: "VoiceConnectorId", Flag: "voice-connector-id", Type: "*string", Required: true},
}

var fields_stop_voice_tone_analysis_task = []leanruntime.Field{
	{Name: "VoiceConnectorId", Flag: "voice-connector-id", Type: "*string", Required: true},
	{Name: "VoiceToneAnalysisTaskId", Flag: "voice-tone-analysis-task-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_global_settings = []leanruntime.Field{
	{Name: "VoiceConnector", Flag: "voice-connector", Type: "*types.VoiceConnectorSettings", Required: false},
}

var fields_update_phone_number = []leanruntime.Field{
	{Name: "CallingName", Flag: "calling-name", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "PhoneNumberId", Flag: "phone-number-id", Type: "*string", Required: true},
	{Name: "ProductType", Flag: "product-type", Type: "types.PhoneNumberProductType", Required: false},
}

var fields_update_phone_number_settings = []leanruntime.Field{
	{Name: "CallingName", Flag: "calling-name", Type: "*string", Required: true},
}

var fields_update_proxy_session = []leanruntime.Field{
	{Name: "Capabilities", Flag: "capabilities", Type: "[]types.Capability", Required: true},
	{Name: "ExpiryMinutes", Flag: "expiry-minutes", Type: "*int32", Required: false},
	{Name: "ProxySessionId", Flag: "proxy-session-id", Type: "*string", Required: true},
	{Name: "VoiceConnectorId", Flag: "voice-connector-id", Type: "*string", Required: true},
}

var fields_update_sip_media_application = []leanruntime.Field{
	{Name: "Endpoints", Flag: "endpoints", Type: "[]types.SipMediaApplicationEndpoint", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "SipMediaApplicationId", Flag: "sip-media-application-id", Type: "*string", Required: true},
}

var fields_update_sip_media_application_call = []leanruntime.Field{
	{Name: "Arguments", Flag: "arguments", Type: "map[string]string", Required: true},
	{Name: "SipMediaApplicationId", Flag: "sip-media-application-id", Type: "*string", Required: true},
	{Name: "TransactionId", Flag: "transaction-id", Type: "*string", Required: true},
}

var fields_update_sip_rule = []leanruntime.Field{
	{Name: "Disabled", Flag: "disabled", Type: "*bool", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "SipRuleId", Flag: "sip-rule-id", Type: "*string", Required: true},
	{Name: "TargetApplications", Flag: "target-applications", Type: "[]types.SipRuleTargetApplication", Required: false},
}

var fields_update_voice_connector = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RequireEncryption", Flag: "require-encryption", Type: "*bool", Required: true},
	{Name: "VoiceConnectorId", Flag: "voice-connector-id", Type: "*string", Required: true},
}

var fields_update_voice_connector_group = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "VoiceConnectorGroupId", Flag: "voice-connector-group-id", Type: "*string", Required: true},
	{Name: "VoiceConnectorItems", Flag: "voice-connector-items", Type: "[]types.VoiceConnectorItem", Required: true},
}

var fields_update_voice_profile = []leanruntime.Field{
	{Name: "SpeakerSearchTaskId", Flag: "speaker-search-task-id", Type: "*string", Required: true},
	{Name: "VoiceProfileId", Flag: "voice-profile-id", Type: "*string", Required: true},
}

var fields_update_voice_profile_domain = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "VoiceProfileDomainId", Flag: "voice-profile-domain-id", Type: "*string", Required: true},
}

var fields_validate_e911_address = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: true},
	{Name: "City", Flag: "city", Type: "*string", Required: true},
	{Name: "Country", Flag: "country", Type: "*string", Required: true},
	{Name: "PostalCode", Flag: "postal-code", Type: "*string", Required: true},
	{Name: "State", Flag: "state", Type: "*string", Required: true},
	{Name: "StreetInfo", Flag: "street-info", Type: "*string", Required: true},
	{Name: "StreetNumber", Flag: "street-number", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-phone-numbers-with-voice-connector": {
			Name:   "associate-phone-numbers-with-voice-connector",
			Fields: fields_associate_phone_numbers_with_voice_connector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociatePhoneNumbersWithVoiceConnectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_phone_numbers_with_voice_connector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociatePhoneNumbersWithVoiceConnector(ctx, input)
			},
		},
		"associate-phone-numbers-with-voice-connector-group": {
			Name:   "associate-phone-numbers-with-voice-connector-group",
			Fields: fields_associate_phone_numbers_with_voice_connector_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociatePhoneNumbersWithVoiceConnectorGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_phone_numbers_with_voice_connector_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociatePhoneNumbersWithVoiceConnectorGroup(ctx, input)
			},
		},
		"batch-delete-phone-number": {
			Name:   "batch-delete-phone-number",
			Fields: fields_batch_delete_phone_number,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDeletePhoneNumberInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_delete_phone_number, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDeletePhoneNumber(ctx, input)
			},
		},
		"batch-update-phone-number": {
			Name:   "batch-update-phone-number",
			Fields: fields_batch_update_phone_number,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchUpdatePhoneNumberInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_update_phone_number, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchUpdatePhoneNumber(ctx, input)
			},
		},
		"create-phone-number-order": {
			Name:   "create-phone-number-order",
			Fields: fields_create_phone_number_order,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePhoneNumberOrderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_phone_number_order, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePhoneNumberOrder(ctx, input)
			},
		},
		"create-proxy-session": {
			Name:   "create-proxy-session",
			Fields: fields_create_proxy_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateProxySessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_proxy_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateProxySession(ctx, input)
			},
		},
		"create-sip-media-application": {
			Name:   "create-sip-media-application",
			Fields: fields_create_sip_media_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSipMediaApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_sip_media_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSipMediaApplication(ctx, input)
			},
		},
		"create-sip-media-application-call": {
			Name:   "create-sip-media-application-call",
			Fields: fields_create_sip_media_application_call,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSipMediaApplicationCallInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_sip_media_application_call, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSipMediaApplicationCall(ctx, input)
			},
		},
		"create-sip-rule": {
			Name:   "create-sip-rule",
			Fields: fields_create_sip_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSipRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_sip_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSipRule(ctx, input)
			},
		},
		"create-voice-connector": {
			Name:   "create-voice-connector",
			Fields: fields_create_voice_connector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVoiceConnectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_voice_connector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVoiceConnector(ctx, input)
			},
		},
		"create-voice-connector-group": {
			Name:   "create-voice-connector-group",
			Fields: fields_create_voice_connector_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVoiceConnectorGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_voice_connector_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVoiceConnectorGroup(ctx, input)
			},
		},
		"create-voice-profile": {
			Name:   "create-voice-profile",
			Fields: fields_create_voice_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVoiceProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_voice_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVoiceProfile(ctx, input)
			},
		},
		"create-voice-profile-domain": {
			Name:   "create-voice-profile-domain",
			Fields: fields_create_voice_profile_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVoiceProfileDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_voice_profile_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVoiceProfileDomain(ctx, input)
			},
		},
		"delete-phone-number": {
			Name:   "delete-phone-number",
			Fields: fields_delete_phone_number,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePhoneNumberInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_phone_number, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePhoneNumber(ctx, input)
			},
		},
		"delete-proxy-session": {
			Name:   "delete-proxy-session",
			Fields: fields_delete_proxy_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteProxySessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_proxy_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteProxySession(ctx, input)
			},
		},
		"delete-sip-media-application": {
			Name:   "delete-sip-media-application",
			Fields: fields_delete_sip_media_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSipMediaApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_sip_media_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSipMediaApplication(ctx, input)
			},
		},
		"delete-sip-rule": {
			Name:   "delete-sip-rule",
			Fields: fields_delete_sip_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSipRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_sip_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSipRule(ctx, input)
			},
		},
		"delete-voice-connector": {
			Name:   "delete-voice-connector",
			Fields: fields_delete_voice_connector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVoiceConnectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_voice_connector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVoiceConnector(ctx, input)
			},
		},
		"delete-voice-connector-emergency-calling-configuration": {
			Name:   "delete-voice-connector-emergency-calling-configuration",
			Fields: fields_delete_voice_connector_emergency_calling_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVoiceConnectorEmergencyCallingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_voice_connector_emergency_calling_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVoiceConnectorEmergencyCallingConfiguration(ctx, input)
			},
		},
		"delete-voice-connector-external-systems-configuration": {
			Name:   "delete-voice-connector-external-systems-configuration",
			Fields: fields_delete_voice_connector_external_systems_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVoiceConnectorExternalSystemsConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_voice_connector_external_systems_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVoiceConnectorExternalSystemsConfiguration(ctx, input)
			},
		},
		"delete-voice-connector-group": {
			Name:   "delete-voice-connector-group",
			Fields: fields_delete_voice_connector_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVoiceConnectorGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_voice_connector_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVoiceConnectorGroup(ctx, input)
			},
		},
		"delete-voice-connector-origination": {
			Name:   "delete-voice-connector-origination",
			Fields: fields_delete_voice_connector_origination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVoiceConnectorOriginationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_voice_connector_origination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVoiceConnectorOrigination(ctx, input)
			},
		},
		"delete-voice-connector-proxy": {
			Name:   "delete-voice-connector-proxy",
			Fields: fields_delete_voice_connector_proxy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVoiceConnectorProxyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_voice_connector_proxy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVoiceConnectorProxy(ctx, input)
			},
		},
		"delete-voice-connector-streaming-configuration": {
			Name:   "delete-voice-connector-streaming-configuration",
			Fields: fields_delete_voice_connector_streaming_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVoiceConnectorStreamingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_voice_connector_streaming_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVoiceConnectorStreamingConfiguration(ctx, input)
			},
		},
		"delete-voice-connector-termination": {
			Name:   "delete-voice-connector-termination",
			Fields: fields_delete_voice_connector_termination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVoiceConnectorTerminationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_voice_connector_termination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVoiceConnectorTermination(ctx, input)
			},
		},
		"delete-voice-connector-termination-credentials": {
			Name:   "delete-voice-connector-termination-credentials",
			Fields: fields_delete_voice_connector_termination_credentials,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVoiceConnectorTerminationCredentialsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_voice_connector_termination_credentials, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVoiceConnectorTerminationCredentials(ctx, input)
			},
		},
		"delete-voice-profile": {
			Name:   "delete-voice-profile",
			Fields: fields_delete_voice_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVoiceProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_voice_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVoiceProfile(ctx, input)
			},
		},
		"delete-voice-profile-domain": {
			Name:   "delete-voice-profile-domain",
			Fields: fields_delete_voice_profile_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVoiceProfileDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_voice_profile_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVoiceProfileDomain(ctx, input)
			},
		},
		"disassociate-phone-numbers-from-voice-connector": {
			Name:   "disassociate-phone-numbers-from-voice-connector",
			Fields: fields_disassociate_phone_numbers_from_voice_connector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociatePhoneNumbersFromVoiceConnectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_phone_numbers_from_voice_connector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociatePhoneNumbersFromVoiceConnector(ctx, input)
			},
		},
		"disassociate-phone-numbers-from-voice-connector-group": {
			Name:   "disassociate-phone-numbers-from-voice-connector-group",
			Fields: fields_disassociate_phone_numbers_from_voice_connector_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociatePhoneNumbersFromVoiceConnectorGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_phone_numbers_from_voice_connector_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociatePhoneNumbersFromVoiceConnectorGroup(ctx, input)
			},
		},
		"get-global-settings": {
			Name:   "get-global-settings",
			Fields: fields_get_global_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetGlobalSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_global_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetGlobalSettings(ctx, input)
			},
		},
		"get-phone-number": {
			Name:   "get-phone-number",
			Fields: fields_get_phone_number,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPhoneNumberInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_phone_number, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPhoneNumber(ctx, input)
			},
		},
		"get-phone-number-order": {
			Name:   "get-phone-number-order",
			Fields: fields_get_phone_number_order,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPhoneNumberOrderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_phone_number_order, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPhoneNumberOrder(ctx, input)
			},
		},
		"get-phone-number-settings": {
			Name:   "get-phone-number-settings",
			Fields: fields_get_phone_number_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPhoneNumberSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_phone_number_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPhoneNumberSettings(ctx, input)
			},
		},
		"get-proxy-session": {
			Name:   "get-proxy-session",
			Fields: fields_get_proxy_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetProxySessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_proxy_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetProxySession(ctx, input)
			},
		},
		"get-sip-media-application": {
			Name:   "get-sip-media-application",
			Fields: fields_get_sip_media_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSipMediaApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_sip_media_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSipMediaApplication(ctx, input)
			},
		},
		"get-sip-media-application-alexa-skill-configuration": {
			Name:   "get-sip-media-application-alexa-skill-configuration",
			Fields: fields_get_sip_media_application_alexa_skill_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSipMediaApplicationAlexaSkillConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_sip_media_application_alexa_skill_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSipMediaApplicationAlexaSkillConfiguration(ctx, input)
			},
		},
		"get-sip-media-application-logging-configuration": {
			Name:   "get-sip-media-application-logging-configuration",
			Fields: fields_get_sip_media_application_logging_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSipMediaApplicationLoggingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_sip_media_application_logging_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSipMediaApplicationLoggingConfiguration(ctx, input)
			},
		},
		"get-sip-rule": {
			Name:   "get-sip-rule",
			Fields: fields_get_sip_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSipRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_sip_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSipRule(ctx, input)
			},
		},
		"get-speaker-search-task": {
			Name:   "get-speaker-search-task",
			Fields: fields_get_speaker_search_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSpeakerSearchTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_speaker_search_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSpeakerSearchTask(ctx, input)
			},
		},
		"get-voice-connector": {
			Name:   "get-voice-connector",
			Fields: fields_get_voice_connector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetVoiceConnectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_voice_connector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetVoiceConnector(ctx, input)
			},
		},
		"get-voice-connector-emergency-calling-configuration": {
			Name:   "get-voice-connector-emergency-calling-configuration",
			Fields: fields_get_voice_connector_emergency_calling_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetVoiceConnectorEmergencyCallingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_voice_connector_emergency_calling_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetVoiceConnectorEmergencyCallingConfiguration(ctx, input)
			},
		},
		"get-voice-connector-external-systems-configuration": {
			Name:   "get-voice-connector-external-systems-configuration",
			Fields: fields_get_voice_connector_external_systems_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetVoiceConnectorExternalSystemsConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_voice_connector_external_systems_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetVoiceConnectorExternalSystemsConfiguration(ctx, input)
			},
		},
		"get-voice-connector-group": {
			Name:   "get-voice-connector-group",
			Fields: fields_get_voice_connector_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetVoiceConnectorGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_voice_connector_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetVoiceConnectorGroup(ctx, input)
			},
		},
		"get-voice-connector-logging-configuration": {
			Name:   "get-voice-connector-logging-configuration",
			Fields: fields_get_voice_connector_logging_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetVoiceConnectorLoggingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_voice_connector_logging_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetVoiceConnectorLoggingConfiguration(ctx, input)
			},
		},
		"get-voice-connector-origination": {
			Name:   "get-voice-connector-origination",
			Fields: fields_get_voice_connector_origination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetVoiceConnectorOriginationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_voice_connector_origination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetVoiceConnectorOrigination(ctx, input)
			},
		},
		"get-voice-connector-proxy": {
			Name:   "get-voice-connector-proxy",
			Fields: fields_get_voice_connector_proxy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetVoiceConnectorProxyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_voice_connector_proxy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetVoiceConnectorProxy(ctx, input)
			},
		},
		"get-voice-connector-streaming-configuration": {
			Name:   "get-voice-connector-streaming-configuration",
			Fields: fields_get_voice_connector_streaming_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetVoiceConnectorStreamingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_voice_connector_streaming_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetVoiceConnectorStreamingConfiguration(ctx, input)
			},
		},
		"get-voice-connector-termination": {
			Name:   "get-voice-connector-termination",
			Fields: fields_get_voice_connector_termination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetVoiceConnectorTerminationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_voice_connector_termination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetVoiceConnectorTermination(ctx, input)
			},
		},
		"get-voice-connector-termination-health": {
			Name:   "get-voice-connector-termination-health",
			Fields: fields_get_voice_connector_termination_health,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetVoiceConnectorTerminationHealthInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_voice_connector_termination_health, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetVoiceConnectorTerminationHealth(ctx, input)
			},
		},
		"get-voice-profile": {
			Name:   "get-voice-profile",
			Fields: fields_get_voice_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetVoiceProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_voice_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetVoiceProfile(ctx, input)
			},
		},
		"get-voice-profile-domain": {
			Name:   "get-voice-profile-domain",
			Fields: fields_get_voice_profile_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetVoiceProfileDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_voice_profile_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetVoiceProfileDomain(ctx, input)
			},
		},
		"get-voice-tone-analysis-task": {
			Name:   "get-voice-tone-analysis-task",
			Fields: fields_get_voice_tone_analysis_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetVoiceToneAnalysisTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_voice_tone_analysis_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetVoiceToneAnalysisTask(ctx, input)
			},
		},
		"list-available-voice-connector-regions": {
			Name:   "list-available-voice-connector-regions",
			Fields: fields_list_available_voice_connector_regions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAvailableVoiceConnectorRegionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_available_voice_connector_regions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListAvailableVoiceConnectorRegions(ctx, input)
			},
		},
		"list-phone-number-orders": {
			Name:   "list-phone-number-orders",
			Fields: fields_list_phone_number_orders,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPhoneNumberOrdersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_phone_number_orders, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPhoneNumberOrders(ctx, input)
				}
				var results []*svc.ListPhoneNumberOrdersOutput
				p := svc.NewListPhoneNumberOrdersPaginator(client, input)
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
		"list-phone-numbers": {
			Name:   "list-phone-numbers",
			Fields: fields_list_phone_numbers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPhoneNumbersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_phone_numbers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPhoneNumbers(ctx, input)
				}
				var results []*svc.ListPhoneNumbersOutput
				p := svc.NewListPhoneNumbersPaginator(client, input)
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
		"list-proxy-sessions": {
			Name:   "list-proxy-sessions",
			Fields: fields_list_proxy_sessions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProxySessionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_proxy_sessions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProxySessions(ctx, input)
				}
				var results []*svc.ListProxySessionsOutput
				p := svc.NewListProxySessionsPaginator(client, input)
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
		"list-sip-media-applications": {
			Name:   "list-sip-media-applications",
			Fields: fields_list_sip_media_applications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSipMediaApplicationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_sip_media_applications, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSipMediaApplications(ctx, input)
				}
				var results []*svc.ListSipMediaApplicationsOutput
				p := svc.NewListSipMediaApplicationsPaginator(client, input)
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
		"list-sip-rules": {
			Name:   "list-sip-rules",
			Fields: fields_list_sip_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSipRulesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_sip_rules, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSipRules(ctx, input)
				}
				var results []*svc.ListSipRulesOutput
				p := svc.NewListSipRulesPaginator(client, input)
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
		"list-supported-phone-number-countries": {
			Name:   "list-supported-phone-number-countries",
			Fields: fields_list_supported_phone_number_countries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSupportedPhoneNumberCountriesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_supported_phone_number_countries, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListSupportedPhoneNumberCountries(ctx, input)
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
		"list-voice-connector-groups": {
			Name:   "list-voice-connector-groups",
			Fields: fields_list_voice_connector_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVoiceConnectorGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_voice_connector_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListVoiceConnectorGroups(ctx, input)
				}
				var results []*svc.ListVoiceConnectorGroupsOutput
				p := svc.NewListVoiceConnectorGroupsPaginator(client, input)
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
		"list-voice-connector-termination-credentials": {
			Name:   "list-voice-connector-termination-credentials",
			Fields: fields_list_voice_connector_termination_credentials,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVoiceConnectorTerminationCredentialsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_voice_connector_termination_credentials, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListVoiceConnectorTerminationCredentials(ctx, input)
			},
		},
		"list-voice-connectors": {
			Name:   "list-voice-connectors",
			Fields: fields_list_voice_connectors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVoiceConnectorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_voice_connectors, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListVoiceConnectors(ctx, input)
				}
				var results []*svc.ListVoiceConnectorsOutput
				p := svc.NewListVoiceConnectorsPaginator(client, input)
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
		"list-voice-profile-domains": {
			Name:   "list-voice-profile-domains",
			Fields: fields_list_voice_profile_domains,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVoiceProfileDomainsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_voice_profile_domains, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListVoiceProfileDomains(ctx, input)
				}
				var results []*svc.ListVoiceProfileDomainsOutput
				p := svc.NewListVoiceProfileDomainsPaginator(client, input)
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
		"list-voice-profiles": {
			Name:   "list-voice-profiles",
			Fields: fields_list_voice_profiles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVoiceProfilesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_voice_profiles, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListVoiceProfiles(ctx, input)
				}
				var results []*svc.ListVoiceProfilesOutput
				p := svc.NewListVoiceProfilesPaginator(client, input)
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
		"put-sip-media-application-alexa-skill-configuration": {
			Name:   "put-sip-media-application-alexa-skill-configuration",
			Fields: fields_put_sip_media_application_alexa_skill_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutSipMediaApplicationAlexaSkillConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_sip_media_application_alexa_skill_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutSipMediaApplicationAlexaSkillConfiguration(ctx, input)
			},
		},
		"put-sip-media-application-logging-configuration": {
			Name:   "put-sip-media-application-logging-configuration",
			Fields: fields_put_sip_media_application_logging_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutSipMediaApplicationLoggingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_sip_media_application_logging_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutSipMediaApplicationLoggingConfiguration(ctx, input)
			},
		},
		"put-voice-connector-emergency-calling-configuration": {
			Name:   "put-voice-connector-emergency-calling-configuration",
			Fields: fields_put_voice_connector_emergency_calling_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutVoiceConnectorEmergencyCallingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_voice_connector_emergency_calling_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutVoiceConnectorEmergencyCallingConfiguration(ctx, input)
			},
		},
		"put-voice-connector-external-systems-configuration": {
			Name:   "put-voice-connector-external-systems-configuration",
			Fields: fields_put_voice_connector_external_systems_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutVoiceConnectorExternalSystemsConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_voice_connector_external_systems_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutVoiceConnectorExternalSystemsConfiguration(ctx, input)
			},
		},
		"put-voice-connector-logging-configuration": {
			Name:   "put-voice-connector-logging-configuration",
			Fields: fields_put_voice_connector_logging_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutVoiceConnectorLoggingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_voice_connector_logging_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutVoiceConnectorLoggingConfiguration(ctx, input)
			},
		},
		"put-voice-connector-origination": {
			Name:   "put-voice-connector-origination",
			Fields: fields_put_voice_connector_origination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutVoiceConnectorOriginationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_voice_connector_origination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutVoiceConnectorOrigination(ctx, input)
			},
		},
		"put-voice-connector-proxy": {
			Name:   "put-voice-connector-proxy",
			Fields: fields_put_voice_connector_proxy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutVoiceConnectorProxyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_voice_connector_proxy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutVoiceConnectorProxy(ctx, input)
			},
		},
		"put-voice-connector-streaming-configuration": {
			Name:   "put-voice-connector-streaming-configuration",
			Fields: fields_put_voice_connector_streaming_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutVoiceConnectorStreamingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_voice_connector_streaming_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutVoiceConnectorStreamingConfiguration(ctx, input)
			},
		},
		"put-voice-connector-termination": {
			Name:   "put-voice-connector-termination",
			Fields: fields_put_voice_connector_termination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutVoiceConnectorTerminationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_voice_connector_termination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutVoiceConnectorTermination(ctx, input)
			},
		},
		"put-voice-connector-termination-credentials": {
			Name:   "put-voice-connector-termination-credentials",
			Fields: fields_put_voice_connector_termination_credentials,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutVoiceConnectorTerminationCredentialsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_voice_connector_termination_credentials, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutVoiceConnectorTerminationCredentials(ctx, input)
			},
		},
		"restore-phone-number": {
			Name:   "restore-phone-number",
			Fields: fields_restore_phone_number,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RestorePhoneNumberInput{}
				if _, err := leanruntime.ApplyInput(input, fields_restore_phone_number, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RestorePhoneNumber(ctx, input)
			},
		},
		"search-available-phone-numbers": {
			Name:   "search-available-phone-numbers",
			Fields: fields_search_available_phone_numbers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchAvailablePhoneNumbersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_available_phone_numbers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchAvailablePhoneNumbers(ctx, input)
				}
				var results []*svc.SearchAvailablePhoneNumbersOutput
				p := svc.NewSearchAvailablePhoneNumbersPaginator(client, input)
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
		"start-speaker-search-task": {
			Name:   "start-speaker-search-task",
			Fields: fields_start_speaker_search_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartSpeakerSearchTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_speaker_search_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartSpeakerSearchTask(ctx, input)
			},
		},
		"start-voice-tone-analysis-task": {
			Name:   "start-voice-tone-analysis-task",
			Fields: fields_start_voice_tone_analysis_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartVoiceToneAnalysisTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_voice_tone_analysis_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartVoiceToneAnalysisTask(ctx, input)
			},
		},
		"stop-speaker-search-task": {
			Name:   "stop-speaker-search-task",
			Fields: fields_stop_speaker_search_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopSpeakerSearchTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_speaker_search_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopSpeakerSearchTask(ctx, input)
			},
		},
		"stop-voice-tone-analysis-task": {
			Name:   "stop-voice-tone-analysis-task",
			Fields: fields_stop_voice_tone_analysis_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopVoiceToneAnalysisTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_voice_tone_analysis_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopVoiceToneAnalysisTask(ctx, input)
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
		"update-global-settings": {
			Name:   "update-global-settings",
			Fields: fields_update_global_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateGlobalSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_global_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateGlobalSettings(ctx, input)
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
		"update-phone-number-settings": {
			Name:   "update-phone-number-settings",
			Fields: fields_update_phone_number_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePhoneNumberSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_phone_number_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePhoneNumberSettings(ctx, input)
			},
		},
		"update-proxy-session": {
			Name:   "update-proxy-session",
			Fields: fields_update_proxy_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateProxySessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_proxy_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateProxySession(ctx, input)
			},
		},
		"update-sip-media-application": {
			Name:   "update-sip-media-application",
			Fields: fields_update_sip_media_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSipMediaApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_sip_media_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSipMediaApplication(ctx, input)
			},
		},
		"update-sip-media-application-call": {
			Name:   "update-sip-media-application-call",
			Fields: fields_update_sip_media_application_call,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSipMediaApplicationCallInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_sip_media_application_call, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSipMediaApplicationCall(ctx, input)
			},
		},
		"update-sip-rule": {
			Name:   "update-sip-rule",
			Fields: fields_update_sip_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSipRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_sip_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSipRule(ctx, input)
			},
		},
		"update-voice-connector": {
			Name:   "update-voice-connector",
			Fields: fields_update_voice_connector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateVoiceConnectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_voice_connector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateVoiceConnector(ctx, input)
			},
		},
		"update-voice-connector-group": {
			Name:   "update-voice-connector-group",
			Fields: fields_update_voice_connector_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateVoiceConnectorGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_voice_connector_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateVoiceConnectorGroup(ctx, input)
			},
		},
		"update-voice-profile": {
			Name:   "update-voice-profile",
			Fields: fields_update_voice_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateVoiceProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_voice_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateVoiceProfile(ctx, input)
			},
		},
		"update-voice-profile-domain": {
			Name:   "update-voice-profile-domain",
			Fields: fields_update_voice_profile_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateVoiceProfileDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_voice_profile_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateVoiceProfileDomain(ctx, input)
			},
		},
		"validate-e911-address": {
			Name:   "validate-e911-address",
			Fields: fields_validate_e911_address,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ValidateE911AddressInput{}
				if _, err := leanruntime.ApplyInput(input, fields_validate_e911_address, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ValidateE911Address(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("chimesdkvoice", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
