package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/pinpoint"
)

var fields_create_app = []leanruntime.Field{
	{Name: "CreateApplicationRequest", Flag: "create-application-request", Type: "*types.CreateApplicationRequest", Required: true},
}

var fields_create_campaign = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "WriteCampaignRequest", Flag: "write-campaign-request", Type: "*types.WriteCampaignRequest", Required: true},
}

var fields_create_email_template = []leanruntime.Field{
	{Name: "EmailTemplateRequest", Flag: "email-template-request", Type: "*types.EmailTemplateRequest", Required: true},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
}

var fields_create_export_job = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "ExportJobRequest", Flag: "export-job-request", Type: "*types.ExportJobRequest", Required: true},
}

var fields_create_import_job = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "ImportJobRequest", Flag: "import-job-request", Type: "*types.ImportJobRequest", Required: true},
}

var fields_create_in_app_template = []leanruntime.Field{
	{Name: "InAppTemplateRequest", Flag: "in-app-template-request", Type: "*types.InAppTemplateRequest", Required: true},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
}

var fields_create_journey = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "WriteJourneyRequest", Flag: "write-journey-request", Type: "*types.WriteJourneyRequest", Required: true},
}

var fields_create_push_template = []leanruntime.Field{
	{Name: "PushNotificationTemplateRequest", Flag: "push-notification-template-request", Type: "*types.PushNotificationTemplateRequest", Required: true},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
}

var fields_create_recommender_configuration = []leanruntime.Field{
	{Name: "CreateRecommenderConfiguration", Flag: "create-recommender-configuration", Type: "*types.CreateRecommenderConfigurationShape", Required: true},
}

var fields_create_segment = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "WriteSegmentRequest", Flag: "write-segment-request", Type: "*types.WriteSegmentRequest", Required: true},
}

var fields_create_sms_template = []leanruntime.Field{
	{Name: "SMSTemplateRequest", Flag: "sms-template-request", Type: "*types.SMSTemplateRequest", Required: true},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
}

var fields_create_voice_template = []leanruntime.Field{
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
	{Name: "VoiceTemplateRequest", Flag: "voice-template-request", Type: "*types.VoiceTemplateRequest", Required: true},
}

var fields_delete_adm_channel = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_delete_apns_channel = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_delete_apns_sandbox_channel = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_delete_apns_voip_channel = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_delete_apns_voip_sandbox_channel = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_delete_app = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_delete_baidu_channel = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_delete_campaign = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "CampaignId", Flag: "campaign-id", Type: "*string", Required: true},
}

var fields_delete_email_channel = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_delete_email_template = []leanruntime.Field{
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
	{Name: "Version", Flag: "version", Type: "*string", Required: false},
}

var fields_delete_endpoint = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "EndpointId", Flag: "endpoint-id", Type: "*string", Required: true},
}

var fields_delete_event_stream = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_delete_gcm_channel = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_delete_in_app_template = []leanruntime.Field{
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
	{Name: "Version", Flag: "version", Type: "*string", Required: false},
}

var fields_delete_journey = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "JourneyId", Flag: "journey-id", Type: "*string", Required: true},
}

var fields_delete_push_template = []leanruntime.Field{
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
	{Name: "Version", Flag: "version", Type: "*string", Required: false},
}

var fields_delete_recommender_configuration = []leanruntime.Field{
	{Name: "RecommenderId", Flag: "recommender-id", Type: "*string", Required: true},
}

var fields_delete_segment = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "SegmentId", Flag: "segment-id", Type: "*string", Required: true},
}

var fields_delete_sms_channel = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_delete_sms_template = []leanruntime.Field{
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
	{Name: "Version", Flag: "version", Type: "*string", Required: false},
}

var fields_delete_user_endpoints = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_delete_voice_channel = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_delete_voice_template = []leanruntime.Field{
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
	{Name: "Version", Flag: "version", Type: "*string", Required: false},
}

var fields_get_adm_channel = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_get_apns_channel = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_get_apns_sandbox_channel = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_get_apns_voip_channel = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_get_apns_voip_sandbox_channel = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_get_app = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_get_application_date_range_kpi = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "KpiName", Flag: "kpi-name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
}

var fields_get_application_settings = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_get_apps = []leanruntime.Field{
	{Name: "PageSize", Flag: "page-size", Type: "*string", Required: false},
	{Name: "Token", Flag: "token", Type: "*string", Required: false},
}

var fields_get_baidu_channel = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_get_campaign = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "CampaignId", Flag: "campaign-id", Type: "*string", Required: true},
}

var fields_get_campaign_activities = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "CampaignId", Flag: "campaign-id", Type: "*string", Required: true},
	{Name: "PageSize", Flag: "page-size", Type: "*string", Required: false},
	{Name: "Token", Flag: "token", Type: "*string", Required: false},
}

var fields_get_campaign_date_range_kpi = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "CampaignId", Flag: "campaign-id", Type: "*string", Required: true},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "KpiName", Flag: "kpi-name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
}

var fields_get_campaign_version = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "CampaignId", Flag: "campaign-id", Type: "*string", Required: true},
	{Name: "Version", Flag: "version", Type: "*string", Required: true},
}

var fields_get_campaign_versions = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "CampaignId", Flag: "campaign-id", Type: "*string", Required: true},
	{Name: "PageSize", Flag: "page-size", Type: "*string", Required: false},
	{Name: "Token", Flag: "token", Type: "*string", Required: false},
}

var fields_get_campaigns = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "PageSize", Flag: "page-size", Type: "*string", Required: false},
	{Name: "Token", Flag: "token", Type: "*string", Required: false},
}

var fields_get_channels = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_get_email_channel = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_get_email_template = []leanruntime.Field{
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
	{Name: "Version", Flag: "version", Type: "*string", Required: false},
}

var fields_get_endpoint = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "EndpointId", Flag: "endpoint-id", Type: "*string", Required: true},
}

var fields_get_event_stream = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_get_export_job = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_get_export_jobs = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "PageSize", Flag: "page-size", Type: "*string", Required: false},
	{Name: "Token", Flag: "token", Type: "*string", Required: false},
}

var fields_get_gcm_channel = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_get_import_job = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_get_import_jobs = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "PageSize", Flag: "page-size", Type: "*string", Required: false},
	{Name: "Token", Flag: "token", Type: "*string", Required: false},
}

var fields_get_in_app_messages = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "EndpointId", Flag: "endpoint-id", Type: "*string", Required: true},
}

var fields_get_in_app_template = []leanruntime.Field{
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
	{Name: "Version", Flag: "version", Type: "*string", Required: false},
}

var fields_get_journey = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "JourneyId", Flag: "journey-id", Type: "*string", Required: true},
}

var fields_get_journey_date_range_kpi = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "JourneyId", Flag: "journey-id", Type: "*string", Required: true},
	{Name: "KpiName", Flag: "kpi-name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
}

var fields_get_journey_execution_activity_metrics = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "JourneyActivityId", Flag: "journey-activity-id", Type: "*string", Required: true},
	{Name: "JourneyId", Flag: "journey-id", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*string", Required: false},
}

var fields_get_journey_execution_metrics = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "JourneyId", Flag: "journey-id", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*string", Required: false},
}

var fields_get_journey_run_execution_activity_metrics = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "JourneyActivityId", Flag: "journey-activity-id", Type: "*string", Required: true},
	{Name: "JourneyId", Flag: "journey-id", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*string", Required: false},
	{Name: "RunId", Flag: "run-id", Type: "*string", Required: true},
}

var fields_get_journey_run_execution_metrics = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "JourneyId", Flag: "journey-id", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*string", Required: false},
	{Name: "RunId", Flag: "run-id", Type: "*string", Required: true},
}

var fields_get_journey_runs = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "JourneyId", Flag: "journey-id", Type: "*string", Required: true},
	{Name: "PageSize", Flag: "page-size", Type: "*string", Required: false},
	{Name: "Token", Flag: "token", Type: "*string", Required: false},
}

var fields_get_push_template = []leanruntime.Field{
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
	{Name: "Version", Flag: "version", Type: "*string", Required: false},
}

var fields_get_recommender_configuration = []leanruntime.Field{
	{Name: "RecommenderId", Flag: "recommender-id", Type: "*string", Required: true},
}

var fields_get_recommender_configurations = []leanruntime.Field{
	{Name: "PageSize", Flag: "page-size", Type: "*string", Required: false},
	{Name: "Token", Flag: "token", Type: "*string", Required: false},
}

var fields_get_segment = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "SegmentId", Flag: "segment-id", Type: "*string", Required: true},
}

var fields_get_segment_export_jobs = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "PageSize", Flag: "page-size", Type: "*string", Required: false},
	{Name: "SegmentId", Flag: "segment-id", Type: "*string", Required: true},
	{Name: "Token", Flag: "token", Type: "*string", Required: false},
}

var fields_get_segment_import_jobs = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "PageSize", Flag: "page-size", Type: "*string", Required: false},
	{Name: "SegmentId", Flag: "segment-id", Type: "*string", Required: true},
	{Name: "Token", Flag: "token", Type: "*string", Required: false},
}

var fields_get_segment_version = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "SegmentId", Flag: "segment-id", Type: "*string", Required: true},
	{Name: "Version", Flag: "version", Type: "*string", Required: true},
}

var fields_get_segment_versions = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "PageSize", Flag: "page-size", Type: "*string", Required: false},
	{Name: "SegmentId", Flag: "segment-id", Type: "*string", Required: true},
	{Name: "Token", Flag: "token", Type: "*string", Required: false},
}

var fields_get_segments = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "PageSize", Flag: "page-size", Type: "*string", Required: false},
	{Name: "Token", Flag: "token", Type: "*string", Required: false},
}

var fields_get_sms_channel = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_get_sms_template = []leanruntime.Field{
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
	{Name: "Version", Flag: "version", Type: "*string", Required: false},
}

var fields_get_user_endpoints = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_get_voice_channel = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_get_voice_template = []leanruntime.Field{
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
	{Name: "Version", Flag: "version", Type: "*string", Required: false},
}

var fields_list_journeys = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "PageSize", Flag: "page-size", Type: "*string", Required: false},
	{Name: "Token", Flag: "token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_template_versions = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*string", Required: false},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
	{Name: "TemplateType", Flag: "template-type", Type: "*string", Required: true},
}

var fields_list_templates = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*string", Required: false},
	{Name: "Prefix", Flag: "prefix", Type: "*string", Required: false},
	{Name: "TemplateType", Flag: "template-type", Type: "*string", Required: false},
}

var fields_phone_number_validate = []leanruntime.Field{
	{Name: "NumberValidateRequest", Flag: "number-validate-request", Type: "*types.NumberValidateRequest", Required: true},
}

var fields_put_event_stream = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "WriteEventStream", Flag: "write-event-stream", Type: "*types.WriteEventStream", Required: true},
}

var fields_put_events = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "EventsRequest", Flag: "events-request", Type: "*types.EventsRequest", Required: true},
}

var fields_remove_attributes = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "AttributeType", Flag: "attribute-type", Type: "*string", Required: true},
	{Name: "UpdateAttributesRequest", Flag: "update-attributes-request", Type: "*types.UpdateAttributesRequest", Required: true},
}

var fields_send_messages = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "MessageRequest", Flag: "message-request", Type: "*types.MessageRequest", Required: true},
}

var fields_send_otp_message = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "SendOTPMessageRequestParameters", Flag: "send-otp-message-request-parameters", Type: "*types.SendOTPMessageRequestParameters", Required: true},
}

var fields_send_users_messages = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "SendUsersMessageRequest", Flag: "send-users-message-request", Type: "*types.SendUsersMessageRequest", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagsModel", Flag: "tags-model", Type: "*types.TagsModel", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_adm_channel = []leanruntime.Field{
	{Name: "ADMChannelRequest", Flag: "adm-channel-request", Type: "*types.ADMChannelRequest", Required: true},
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_update_apns_channel = []leanruntime.Field{
	{Name: "APNSChannelRequest", Flag: "apns-channel-request", Type: "*types.APNSChannelRequest", Required: true},
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_update_apns_sandbox_channel = []leanruntime.Field{
	{Name: "APNSSandboxChannelRequest", Flag: "apns-sandbox-channel-request", Type: "*types.APNSSandboxChannelRequest", Required: true},
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_update_apns_voip_channel = []leanruntime.Field{
	{Name: "APNSVoipChannelRequest", Flag: "apns-voip-channel-request", Type: "*types.APNSVoipChannelRequest", Required: true},
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_update_apns_voip_sandbox_channel = []leanruntime.Field{
	{Name: "APNSVoipSandboxChannelRequest", Flag: "apns-voip-sandbox-channel-request", Type: "*types.APNSVoipSandboxChannelRequest", Required: true},
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_update_application_settings = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "WriteApplicationSettingsRequest", Flag: "write-application-settings-request", Type: "*types.WriteApplicationSettingsRequest", Required: true},
}

var fields_update_baidu_channel = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "BaiduChannelRequest", Flag: "baidu-channel-request", Type: "*types.BaiduChannelRequest", Required: true},
}

var fields_update_campaign = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "CampaignId", Flag: "campaign-id", Type: "*string", Required: true},
	{Name: "WriteCampaignRequest", Flag: "write-campaign-request", Type: "*types.WriteCampaignRequest", Required: true},
}

var fields_update_email_channel = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "EmailChannelRequest", Flag: "email-channel-request", Type: "*types.EmailChannelRequest", Required: true},
}

var fields_update_email_template = []leanruntime.Field{
	{Name: "CreateNewVersion", Flag: "create-new-version", Type: "*bool", Required: false},
	{Name: "EmailTemplateRequest", Flag: "email-template-request", Type: "*types.EmailTemplateRequest", Required: true},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
	{Name: "Version", Flag: "version", Type: "*string", Required: false},
}

var fields_update_endpoint = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "EndpointId", Flag: "endpoint-id", Type: "*string", Required: true},
	{Name: "EndpointRequest", Flag: "endpoint-request", Type: "*types.EndpointRequest", Required: true},
}

var fields_update_endpoints_batch = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "EndpointBatchRequest", Flag: "endpoint-batch-request", Type: "*types.EndpointBatchRequest", Required: true},
}

var fields_update_gcm_channel = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "GCMChannelRequest", Flag: "gcm-channel-request", Type: "*types.GCMChannelRequest", Required: true},
}

var fields_update_in_app_template = []leanruntime.Field{
	{Name: "CreateNewVersion", Flag: "create-new-version", Type: "*bool", Required: false},
	{Name: "InAppTemplateRequest", Flag: "in-app-template-request", Type: "*types.InAppTemplateRequest", Required: true},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
	{Name: "Version", Flag: "version", Type: "*string", Required: false},
}

var fields_update_journey = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "JourneyId", Flag: "journey-id", Type: "*string", Required: true},
	{Name: "WriteJourneyRequest", Flag: "write-journey-request", Type: "*types.WriteJourneyRequest", Required: true},
}

var fields_update_journey_state = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "JourneyId", Flag: "journey-id", Type: "*string", Required: true},
	{Name: "JourneyStateRequest", Flag: "journey-state-request", Type: "*types.JourneyStateRequest", Required: true},
}

var fields_update_push_template = []leanruntime.Field{
	{Name: "CreateNewVersion", Flag: "create-new-version", Type: "*bool", Required: false},
	{Name: "PushNotificationTemplateRequest", Flag: "push-notification-template-request", Type: "*types.PushNotificationTemplateRequest", Required: true},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
	{Name: "Version", Flag: "version", Type: "*string", Required: false},
}

var fields_update_recommender_configuration = []leanruntime.Field{
	{Name: "RecommenderId", Flag: "recommender-id", Type: "*string", Required: true},
	{Name: "UpdateRecommenderConfiguration", Flag: "update-recommender-configuration", Type: "*types.UpdateRecommenderConfigurationShape", Required: true},
}

var fields_update_segment = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "SegmentId", Flag: "segment-id", Type: "*string", Required: true},
	{Name: "WriteSegmentRequest", Flag: "write-segment-request", Type: "*types.WriteSegmentRequest", Required: true},
}

var fields_update_sms_channel = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "SMSChannelRequest", Flag: "sms-channel-request", Type: "*types.SMSChannelRequest", Required: true},
}

var fields_update_sms_template = []leanruntime.Field{
	{Name: "CreateNewVersion", Flag: "create-new-version", Type: "*bool", Required: false},
	{Name: "SMSTemplateRequest", Flag: "sms-template-request", Type: "*types.SMSTemplateRequest", Required: true},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
	{Name: "Version", Flag: "version", Type: "*string", Required: false},
}

var fields_update_template_active_version = []leanruntime.Field{
	{Name: "TemplateActiveVersionRequest", Flag: "template-active-version-request", Type: "*types.TemplateActiveVersionRequest", Required: true},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
	{Name: "TemplateType", Flag: "template-type", Type: "*string", Required: true},
}

var fields_update_voice_channel = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "VoiceChannelRequest", Flag: "voice-channel-request", Type: "*types.VoiceChannelRequest", Required: true},
}

var fields_update_voice_template = []leanruntime.Field{
	{Name: "CreateNewVersion", Flag: "create-new-version", Type: "*bool", Required: false},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
	{Name: "Version", Flag: "version", Type: "*string", Required: false},
	{Name: "VoiceTemplateRequest", Flag: "voice-template-request", Type: "*types.VoiceTemplateRequest", Required: true},
}

var fields_verify_otp_message = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "VerifyOTPMessageRequestParameters", Flag: "verify-otp-message-request-parameters", Type: "*types.VerifyOTPMessageRequestParameters", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-app": {
			Name:   "create-app",
			Fields: fields_create_app,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAppInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_app, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateApp(ctx, input)
			},
		},
		"create-campaign": {
			Name:   "create-campaign",
			Fields: fields_create_campaign,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCampaignInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_campaign, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCampaign(ctx, input)
			},
		},
		"create-email-template": {
			Name:   "create-email-template",
			Fields: fields_create_email_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEmailTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_email_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEmailTemplate(ctx, input)
			},
		},
		"create-export-job": {
			Name:   "create-export-job",
			Fields: fields_create_export_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateExportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_export_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateExportJob(ctx, input)
			},
		},
		"create-import-job": {
			Name:   "create-import-job",
			Fields: fields_create_import_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateImportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_import_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateImportJob(ctx, input)
			},
		},
		"create-in-app-template": {
			Name:   "create-in-app-template",
			Fields: fields_create_in_app_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateInAppTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_in_app_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateInAppTemplate(ctx, input)
			},
		},
		"create-journey": {
			Name:   "create-journey",
			Fields: fields_create_journey,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateJourneyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_journey, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateJourney(ctx, input)
			},
		},
		"create-push-template": {
			Name:   "create-push-template",
			Fields: fields_create_push_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePushTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_push_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePushTemplate(ctx, input)
			},
		},
		"create-recommender-configuration": {
			Name:   "create-recommender-configuration",
			Fields: fields_create_recommender_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRecommenderConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_recommender_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRecommenderConfiguration(ctx, input)
			},
		},
		"create-segment": {
			Name:   "create-segment",
			Fields: fields_create_segment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSegmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_segment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSegment(ctx, input)
			},
		},
		"create-sms-template": {
			Name:   "create-sms-template",
			Fields: fields_create_sms_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSmsTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_sms_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSmsTemplate(ctx, input)
			},
		},
		"create-voice-template": {
			Name:   "create-voice-template",
			Fields: fields_create_voice_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVoiceTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_voice_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVoiceTemplate(ctx, input)
			},
		},
		"delete-adm-channel": {
			Name:   "delete-adm-channel",
			Fields: fields_delete_adm_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAdmChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_adm_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAdmChannel(ctx, input)
			},
		},
		"delete-apns-channel": {
			Name:   "delete-apns-channel",
			Fields: fields_delete_apns_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteApnsChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_apns_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteApnsChannel(ctx, input)
			},
		},
		"delete-apns-sandbox-channel": {
			Name:   "delete-apns-sandbox-channel",
			Fields: fields_delete_apns_sandbox_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteApnsSandboxChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_apns_sandbox_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteApnsSandboxChannel(ctx, input)
			},
		},
		"delete-apns-voip-channel": {
			Name:   "delete-apns-voip-channel",
			Fields: fields_delete_apns_voip_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteApnsVoipChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_apns_voip_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteApnsVoipChannel(ctx, input)
			},
		},
		"delete-apns-voip-sandbox-channel": {
			Name:   "delete-apns-voip-sandbox-channel",
			Fields: fields_delete_apns_voip_sandbox_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteApnsVoipSandboxChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_apns_voip_sandbox_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteApnsVoipSandboxChannel(ctx, input)
			},
		},
		"delete-app": {
			Name:   "delete-app",
			Fields: fields_delete_app,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAppInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_app, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteApp(ctx, input)
			},
		},
		"delete-baidu-channel": {
			Name:   "delete-baidu-channel",
			Fields: fields_delete_baidu_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBaiduChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_baidu_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBaiduChannel(ctx, input)
			},
		},
		"delete-campaign": {
			Name:   "delete-campaign",
			Fields: fields_delete_campaign,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCampaignInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_campaign, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCampaign(ctx, input)
			},
		},
		"delete-email-channel": {
			Name:   "delete-email-channel",
			Fields: fields_delete_email_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEmailChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_email_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEmailChannel(ctx, input)
			},
		},
		"delete-email-template": {
			Name:   "delete-email-template",
			Fields: fields_delete_email_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEmailTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_email_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEmailTemplate(ctx, input)
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
		"delete-event-stream": {
			Name:   "delete-event-stream",
			Fields: fields_delete_event_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEventStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_event_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEventStream(ctx, input)
			},
		},
		"delete-gcm-channel": {
			Name:   "delete-gcm-channel",
			Fields: fields_delete_gcm_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteGcmChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_gcm_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteGcmChannel(ctx, input)
			},
		},
		"delete-in-app-template": {
			Name:   "delete-in-app-template",
			Fields: fields_delete_in_app_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteInAppTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_in_app_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteInAppTemplate(ctx, input)
			},
		},
		"delete-journey": {
			Name:   "delete-journey",
			Fields: fields_delete_journey,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteJourneyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_journey, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteJourney(ctx, input)
			},
		},
		"delete-push-template": {
			Name:   "delete-push-template",
			Fields: fields_delete_push_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePushTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_push_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePushTemplate(ctx, input)
			},
		},
		"delete-recommender-configuration": {
			Name:   "delete-recommender-configuration",
			Fields: fields_delete_recommender_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRecommenderConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_recommender_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRecommenderConfiguration(ctx, input)
			},
		},
		"delete-segment": {
			Name:   "delete-segment",
			Fields: fields_delete_segment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSegmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_segment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSegment(ctx, input)
			},
		},
		"delete-sms-channel": {
			Name:   "delete-sms-channel",
			Fields: fields_delete_sms_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSmsChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_sms_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSmsChannel(ctx, input)
			},
		},
		"delete-sms-template": {
			Name:   "delete-sms-template",
			Fields: fields_delete_sms_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSmsTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_sms_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSmsTemplate(ctx, input)
			},
		},
		"delete-user-endpoints": {
			Name:   "delete-user-endpoints",
			Fields: fields_delete_user_endpoints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteUserEndpointsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_user_endpoints, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteUserEndpoints(ctx, input)
			},
		},
		"delete-voice-channel": {
			Name:   "delete-voice-channel",
			Fields: fields_delete_voice_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVoiceChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_voice_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVoiceChannel(ctx, input)
			},
		},
		"delete-voice-template": {
			Name:   "delete-voice-template",
			Fields: fields_delete_voice_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVoiceTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_voice_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVoiceTemplate(ctx, input)
			},
		},
		"get-adm-channel": {
			Name:   "get-adm-channel",
			Fields: fields_get_adm_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAdmChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_adm_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAdmChannel(ctx, input)
			},
		},
		"get-apns-channel": {
			Name:   "get-apns-channel",
			Fields: fields_get_apns_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetApnsChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_apns_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetApnsChannel(ctx, input)
			},
		},
		"get-apns-sandbox-channel": {
			Name:   "get-apns-sandbox-channel",
			Fields: fields_get_apns_sandbox_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetApnsSandboxChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_apns_sandbox_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetApnsSandboxChannel(ctx, input)
			},
		},
		"get-apns-voip-channel": {
			Name:   "get-apns-voip-channel",
			Fields: fields_get_apns_voip_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetApnsVoipChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_apns_voip_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetApnsVoipChannel(ctx, input)
			},
		},
		"get-apns-voip-sandbox-channel": {
			Name:   "get-apns-voip-sandbox-channel",
			Fields: fields_get_apns_voip_sandbox_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetApnsVoipSandboxChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_apns_voip_sandbox_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetApnsVoipSandboxChannel(ctx, input)
			},
		},
		"get-app": {
			Name:   "get-app",
			Fields: fields_get_app,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAppInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_app, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetApp(ctx, input)
			},
		},
		"get-application-date-range-kpi": {
			Name:   "get-application-date-range-kpi",
			Fields: fields_get_application_date_range_kpi,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetApplicationDateRangeKpiInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_application_date_range_kpi, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetApplicationDateRangeKpi(ctx, input)
			},
		},
		"get-application-settings": {
			Name:   "get-application-settings",
			Fields: fields_get_application_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetApplicationSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_application_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetApplicationSettings(ctx, input)
			},
		},
		"get-apps": {
			Name:   "get-apps",
			Fields: fields_get_apps,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAppsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_apps, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetApps(ctx, input)
			},
		},
		"get-baidu-channel": {
			Name:   "get-baidu-channel",
			Fields: fields_get_baidu_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBaiduChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_baidu_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBaiduChannel(ctx, input)
			},
		},
		"get-campaign": {
			Name:   "get-campaign",
			Fields: fields_get_campaign,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCampaignInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_campaign, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCampaign(ctx, input)
			},
		},
		"get-campaign-activities": {
			Name:   "get-campaign-activities",
			Fields: fields_get_campaign_activities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCampaignActivitiesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_campaign_activities, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCampaignActivities(ctx, input)
			},
		},
		"get-campaign-date-range-kpi": {
			Name:   "get-campaign-date-range-kpi",
			Fields: fields_get_campaign_date_range_kpi,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCampaignDateRangeKpiInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_campaign_date_range_kpi, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCampaignDateRangeKpi(ctx, input)
			},
		},
		"get-campaign-version": {
			Name:   "get-campaign-version",
			Fields: fields_get_campaign_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCampaignVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_campaign_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCampaignVersion(ctx, input)
			},
		},
		"get-campaign-versions": {
			Name:   "get-campaign-versions",
			Fields: fields_get_campaign_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCampaignVersionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_campaign_versions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCampaignVersions(ctx, input)
			},
		},
		"get-campaigns": {
			Name:   "get-campaigns",
			Fields: fields_get_campaigns,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCampaignsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_campaigns, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCampaigns(ctx, input)
			},
		},
		"get-channels": {
			Name:   "get-channels",
			Fields: fields_get_channels,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetChannelsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_channels, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetChannels(ctx, input)
			},
		},
		"get-email-channel": {
			Name:   "get-email-channel",
			Fields: fields_get_email_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEmailChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_email_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEmailChannel(ctx, input)
			},
		},
		"get-email-template": {
			Name:   "get-email-template",
			Fields: fields_get_email_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEmailTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_email_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEmailTemplate(ctx, input)
			},
		},
		"get-endpoint": {
			Name:   "get-endpoint",
			Fields: fields_get_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEndpoint(ctx, input)
			},
		},
		"get-event-stream": {
			Name:   "get-event-stream",
			Fields: fields_get_event_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEventStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_event_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEventStream(ctx, input)
			},
		},
		"get-export-job": {
			Name:   "get-export-job",
			Fields: fields_get_export_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetExportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_export_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetExportJob(ctx, input)
			},
		},
		"get-export-jobs": {
			Name:   "get-export-jobs",
			Fields: fields_get_export_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetExportJobsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_export_jobs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetExportJobs(ctx, input)
			},
		},
		"get-gcm-channel": {
			Name:   "get-gcm-channel",
			Fields: fields_get_gcm_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetGcmChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_gcm_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetGcmChannel(ctx, input)
			},
		},
		"get-import-job": {
			Name:   "get-import-job",
			Fields: fields_get_import_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetImportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_import_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetImportJob(ctx, input)
			},
		},
		"get-import-jobs": {
			Name:   "get-import-jobs",
			Fields: fields_get_import_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetImportJobsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_import_jobs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetImportJobs(ctx, input)
			},
		},
		"get-in-app-messages": {
			Name:   "get-in-app-messages",
			Fields: fields_get_in_app_messages,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetInAppMessagesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_in_app_messages, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetInAppMessages(ctx, input)
			},
		},
		"get-in-app-template": {
			Name:   "get-in-app-template",
			Fields: fields_get_in_app_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetInAppTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_in_app_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetInAppTemplate(ctx, input)
			},
		},
		"get-journey": {
			Name:   "get-journey",
			Fields: fields_get_journey,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetJourneyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_journey, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetJourney(ctx, input)
			},
		},
		"get-journey-date-range-kpi": {
			Name:   "get-journey-date-range-kpi",
			Fields: fields_get_journey_date_range_kpi,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetJourneyDateRangeKpiInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_journey_date_range_kpi, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetJourneyDateRangeKpi(ctx, input)
			},
		},
		"get-journey-execution-activity-metrics": {
			Name:   "get-journey-execution-activity-metrics",
			Fields: fields_get_journey_execution_activity_metrics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetJourneyExecutionActivityMetricsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_journey_execution_activity_metrics, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetJourneyExecutionActivityMetrics(ctx, input)
			},
		},
		"get-journey-execution-metrics": {
			Name:   "get-journey-execution-metrics",
			Fields: fields_get_journey_execution_metrics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetJourneyExecutionMetricsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_journey_execution_metrics, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetJourneyExecutionMetrics(ctx, input)
			},
		},
		"get-journey-run-execution-activity-metrics": {
			Name:   "get-journey-run-execution-activity-metrics",
			Fields: fields_get_journey_run_execution_activity_metrics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetJourneyRunExecutionActivityMetricsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_journey_run_execution_activity_metrics, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetJourneyRunExecutionActivityMetrics(ctx, input)
			},
		},
		"get-journey-run-execution-metrics": {
			Name:   "get-journey-run-execution-metrics",
			Fields: fields_get_journey_run_execution_metrics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetJourneyRunExecutionMetricsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_journey_run_execution_metrics, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetJourneyRunExecutionMetrics(ctx, input)
			},
		},
		"get-journey-runs": {
			Name:   "get-journey-runs",
			Fields: fields_get_journey_runs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetJourneyRunsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_journey_runs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetJourneyRuns(ctx, input)
			},
		},
		"get-push-template": {
			Name:   "get-push-template",
			Fields: fields_get_push_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPushTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_push_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPushTemplate(ctx, input)
			},
		},
		"get-recommender-configuration": {
			Name:   "get-recommender-configuration",
			Fields: fields_get_recommender_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRecommenderConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_recommender_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRecommenderConfiguration(ctx, input)
			},
		},
		"get-recommender-configurations": {
			Name:   "get-recommender-configurations",
			Fields: fields_get_recommender_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRecommenderConfigurationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_recommender_configurations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRecommenderConfigurations(ctx, input)
			},
		},
		"get-segment": {
			Name:   "get-segment",
			Fields: fields_get_segment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSegmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_segment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSegment(ctx, input)
			},
		},
		"get-segment-export-jobs": {
			Name:   "get-segment-export-jobs",
			Fields: fields_get_segment_export_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSegmentExportJobsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_segment_export_jobs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSegmentExportJobs(ctx, input)
			},
		},
		"get-segment-import-jobs": {
			Name:   "get-segment-import-jobs",
			Fields: fields_get_segment_import_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSegmentImportJobsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_segment_import_jobs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSegmentImportJobs(ctx, input)
			},
		},
		"get-segment-version": {
			Name:   "get-segment-version",
			Fields: fields_get_segment_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSegmentVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_segment_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSegmentVersion(ctx, input)
			},
		},
		"get-segment-versions": {
			Name:   "get-segment-versions",
			Fields: fields_get_segment_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSegmentVersionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_segment_versions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSegmentVersions(ctx, input)
			},
		},
		"get-segments": {
			Name:   "get-segments",
			Fields: fields_get_segments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSegmentsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_segments, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSegments(ctx, input)
			},
		},
		"get-sms-channel": {
			Name:   "get-sms-channel",
			Fields: fields_get_sms_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSmsChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_sms_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSmsChannel(ctx, input)
			},
		},
		"get-sms-template": {
			Name:   "get-sms-template",
			Fields: fields_get_sms_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSmsTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_sms_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSmsTemplate(ctx, input)
			},
		},
		"get-user-endpoints": {
			Name:   "get-user-endpoints",
			Fields: fields_get_user_endpoints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetUserEndpointsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_user_endpoints, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetUserEndpoints(ctx, input)
			},
		},
		"get-voice-channel": {
			Name:   "get-voice-channel",
			Fields: fields_get_voice_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetVoiceChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_voice_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetVoiceChannel(ctx, input)
			},
		},
		"get-voice-template": {
			Name:   "get-voice-template",
			Fields: fields_get_voice_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetVoiceTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_voice_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetVoiceTemplate(ctx, input)
			},
		},
		"list-journeys": {
			Name:   "list-journeys",
			Fields: fields_list_journeys,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListJourneysInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_journeys, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListJourneys(ctx, input)
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
		"list-template-versions": {
			Name:   "list-template-versions",
			Fields: fields_list_template_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTemplateVersionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_template_versions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTemplateVersions(ctx, input)
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
		"phone-number-validate": {
			Name:   "phone-number-validate",
			Fields: fields_phone_number_validate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PhoneNumberValidateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_phone_number_validate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PhoneNumberValidate(ctx, input)
			},
		},
		"put-event-stream": {
			Name:   "put-event-stream",
			Fields: fields_put_event_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutEventStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_event_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutEventStream(ctx, input)
			},
		},
		"put-events": {
			Name:   "put-events",
			Fields: fields_put_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutEventsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_events, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutEvents(ctx, input)
			},
		},
		"remove-attributes": {
			Name:   "remove-attributes",
			Fields: fields_remove_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveAttributes(ctx, input)
			},
		},
		"send-messages": {
			Name:   "send-messages",
			Fields: fields_send_messages,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendMessagesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_messages, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendMessages(ctx, input)
			},
		},
		"send-otp-message": {
			Name:   "send-otp-message",
			Fields: fields_send_otp_message,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendOTPMessageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_otp_message, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendOTPMessage(ctx, input)
			},
		},
		"send-users-messages": {
			Name:   "send-users-messages",
			Fields: fields_send_users_messages,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendUsersMessagesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_users_messages, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendUsersMessages(ctx, input)
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
		"update-adm-channel": {
			Name:   "update-adm-channel",
			Fields: fields_update_adm_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAdmChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_adm_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAdmChannel(ctx, input)
			},
		},
		"update-apns-channel": {
			Name:   "update-apns-channel",
			Fields: fields_update_apns_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateApnsChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_apns_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateApnsChannel(ctx, input)
			},
		},
		"update-apns-sandbox-channel": {
			Name:   "update-apns-sandbox-channel",
			Fields: fields_update_apns_sandbox_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateApnsSandboxChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_apns_sandbox_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateApnsSandboxChannel(ctx, input)
			},
		},
		"update-apns-voip-channel": {
			Name:   "update-apns-voip-channel",
			Fields: fields_update_apns_voip_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateApnsVoipChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_apns_voip_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateApnsVoipChannel(ctx, input)
			},
		},
		"update-apns-voip-sandbox-channel": {
			Name:   "update-apns-voip-sandbox-channel",
			Fields: fields_update_apns_voip_sandbox_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateApnsVoipSandboxChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_apns_voip_sandbox_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateApnsVoipSandboxChannel(ctx, input)
			},
		},
		"update-application-settings": {
			Name:   "update-application-settings",
			Fields: fields_update_application_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateApplicationSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_application_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateApplicationSettings(ctx, input)
			},
		},
		"update-baidu-channel": {
			Name:   "update-baidu-channel",
			Fields: fields_update_baidu_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateBaiduChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_baidu_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateBaiduChannel(ctx, input)
			},
		},
		"update-campaign": {
			Name:   "update-campaign",
			Fields: fields_update_campaign,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCampaignInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_campaign, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCampaign(ctx, input)
			},
		},
		"update-email-channel": {
			Name:   "update-email-channel",
			Fields: fields_update_email_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEmailChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_email_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEmailChannel(ctx, input)
			},
		},
		"update-email-template": {
			Name:   "update-email-template",
			Fields: fields_update_email_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEmailTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_email_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEmailTemplate(ctx, input)
			},
		},
		"update-endpoint": {
			Name:   "update-endpoint",
			Fields: fields_update_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEndpoint(ctx, input)
			},
		},
		"update-endpoints-batch": {
			Name:   "update-endpoints-batch",
			Fields: fields_update_endpoints_batch,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEndpointsBatchInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_endpoints_batch, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEndpointsBatch(ctx, input)
			},
		},
		"update-gcm-channel": {
			Name:   "update-gcm-channel",
			Fields: fields_update_gcm_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateGcmChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_gcm_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateGcmChannel(ctx, input)
			},
		},
		"update-in-app-template": {
			Name:   "update-in-app-template",
			Fields: fields_update_in_app_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateInAppTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_in_app_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateInAppTemplate(ctx, input)
			},
		},
		"update-journey": {
			Name:   "update-journey",
			Fields: fields_update_journey,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateJourneyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_journey, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateJourney(ctx, input)
			},
		},
		"update-journey-state": {
			Name:   "update-journey-state",
			Fields: fields_update_journey_state,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateJourneyStateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_journey_state, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateJourneyState(ctx, input)
			},
		},
		"update-push-template": {
			Name:   "update-push-template",
			Fields: fields_update_push_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePushTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_push_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePushTemplate(ctx, input)
			},
		},
		"update-recommender-configuration": {
			Name:   "update-recommender-configuration",
			Fields: fields_update_recommender_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRecommenderConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_recommender_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRecommenderConfiguration(ctx, input)
			},
		},
		"update-segment": {
			Name:   "update-segment",
			Fields: fields_update_segment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSegmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_segment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSegment(ctx, input)
			},
		},
		"update-sms-channel": {
			Name:   "update-sms-channel",
			Fields: fields_update_sms_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSmsChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_sms_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSmsChannel(ctx, input)
			},
		},
		"update-sms-template": {
			Name:   "update-sms-template",
			Fields: fields_update_sms_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSmsTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_sms_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSmsTemplate(ctx, input)
			},
		},
		"update-template-active-version": {
			Name:   "update-template-active-version",
			Fields: fields_update_template_active_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTemplateActiveVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_template_active_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTemplateActiveVersion(ctx, input)
			},
		},
		"update-voice-channel": {
			Name:   "update-voice-channel",
			Fields: fields_update_voice_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateVoiceChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_voice_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateVoiceChannel(ctx, input)
			},
		},
		"update-voice-template": {
			Name:   "update-voice-template",
			Fields: fields_update_voice_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateVoiceTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_voice_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateVoiceTemplate(ctx, input)
			},
		},
		"verify-otp-message": {
			Name:   "verify-otp-message",
			Fields: fields_verify_otp_message,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.VerifyOTPMessageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_verify_otp_message, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.VerifyOTPMessage(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("pinpoint", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
