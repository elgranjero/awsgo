package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/sesv2"
)

var fields_batch_get_metric_data = []leanruntime.Field{
	{Name: "Queries", Flag: "queries", Type: "[]types.BatchGetMetricDataQuery", Required: true},
}

var fields_cancel_export_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_create_configuration_set = []leanruntime.Field{
	{Name: "ArchivingOptions", Flag: "archiving-options", Type: "*types.ArchivingOptions", Required: false},
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
	{Name: "DeliveryOptions", Flag: "delivery-options", Type: "*types.DeliveryOptions", Required: false},
	{Name: "ReputationOptions", Flag: "reputation-options", Type: "*types.ReputationOptions", Required: false},
	{Name: "SendingOptions", Flag: "sending-options", Type: "*types.SendingOptions", Required: false},
	{Name: "SuppressionOptions", Flag: "suppression-options", Type: "*types.SuppressionOptions", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TrackingOptions", Flag: "tracking-options", Type: "*types.TrackingOptions", Required: false},
	{Name: "VdmOptions", Flag: "vdm-options", Type: "*types.VdmOptions", Required: false},
}

var fields_create_configuration_set_event_destination = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
	{Name: "EventDestination", Flag: "event-destination", Type: "*types.EventDestinationDefinition", Required: true},
	{Name: "EventDestinationName", Flag: "event-destination-name", Type: "*string", Required: true},
}

var fields_create_contact = []leanruntime.Field{
	{Name: "AttributesData", Flag: "attributes-data", Type: "*string", Required: false},
	{Name: "ContactListName", Flag: "contact-list-name", Type: "*string", Required: true},
	{Name: "EmailAddress", Flag: "email-address", Type: "*string", Required: true},
	{Name: "TopicPreferences", Flag: "topic-preferences", Type: "[]types.TopicPreference", Required: false},
	{Name: "UnsubscribeAll", Flag: "unsubscribe-all", Type: "bool", Required: false},
}

var fields_create_contact_list = []leanruntime.Field{
	{Name: "ContactListName", Flag: "contact-list-name", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Topics", Flag: "topics", Type: "[]types.Topic", Required: false},
}

var fields_create_custom_verification_email_template = []leanruntime.Field{
	{Name: "FailureRedirectionURL", Flag: "failure-redirection-url", Type: "*string", Required: true},
	{Name: "FromEmailAddress", Flag: "from-email-address", Type: "*string", Required: true},
	{Name: "SuccessRedirectionURL", Flag: "success-redirection-url", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TemplateContent", Flag: "template-content", Type: "*string", Required: true},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
	{Name: "TemplateSubject", Flag: "template-subject", Type: "*string", Required: true},
}

var fields_create_dedicated_ip_pool = []leanruntime.Field{
	{Name: "PoolName", Flag: "pool-name", Type: "*string", Required: true},
	{Name: "ScalingMode", Flag: "scaling-mode", Type: "types.ScalingMode", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_deliverability_test_report = []leanruntime.Field{
	{Name: "Content", Flag: "content", Type: "*types.EmailContent", Required: true},
	{Name: "FromEmailAddress", Flag: "from-email-address", Type: "*string", Required: true},
	{Name: "ReportName", Flag: "report-name", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_email_identity = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: false},
	{Name: "DkimSigningAttributes", Flag: "dkim-signing-attributes", Type: "*types.DkimSigningAttributes", Required: false},
	{Name: "EmailIdentity", Flag: "email-identity", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_email_identity_policy = []leanruntime.Field{
	{Name: "EmailIdentity", Flag: "email-identity", Type: "*string", Required: true},
	{Name: "Policy", Flag: "policy", Type: "*string", Required: true},
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
}

var fields_create_email_template = []leanruntime.Field{
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TemplateContent", Flag: "template-content", Type: "*types.EmailTemplateContent", Required: true},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
}

var fields_create_export_job = []leanruntime.Field{
	{Name: "ExportDataSource", Flag: "export-data-source", Type: "*types.ExportDataSource", Required: true},
	{Name: "ExportDestination", Flag: "export-destination", Type: "*types.ExportDestination", Required: true},
}

var fields_create_import_job = []leanruntime.Field{
	{Name: "ImportDataSource", Flag: "import-data-source", Type: "*types.ImportDataSource", Required: true},
	{Name: "ImportDestination", Flag: "import-destination", Type: "*types.ImportDestination", Required: true},
}

var fields_create_multi_region_endpoint = []leanruntime.Field{
	{Name: "Details", Flag: "details", Type: "*types.Details", Required: true},
	{Name: "EndpointName", Flag: "endpoint-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_tenant = []leanruntime.Field{
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TenantName", Flag: "tenant-name", Type: "*string", Required: true},
}

var fields_create_tenant_resource_association = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TenantName", Flag: "tenant-name", Type: "*string", Required: true},
}

var fields_delete_configuration_set = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
}

var fields_delete_configuration_set_event_destination = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
	{Name: "EventDestinationName", Flag: "event-destination-name", Type: "*string", Required: true},
}

var fields_delete_contact = []leanruntime.Field{
	{Name: "ContactListName", Flag: "contact-list-name", Type: "*string", Required: true},
	{Name: "EmailAddress", Flag: "email-address", Type: "*string", Required: true},
}

var fields_delete_contact_list = []leanruntime.Field{
	{Name: "ContactListName", Flag: "contact-list-name", Type: "*string", Required: true},
}

var fields_delete_custom_verification_email_template = []leanruntime.Field{
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
}

var fields_delete_dedicated_ip_pool = []leanruntime.Field{
	{Name: "PoolName", Flag: "pool-name", Type: "*string", Required: true},
}

var fields_delete_email_identity = []leanruntime.Field{
	{Name: "EmailIdentity", Flag: "email-identity", Type: "*string", Required: true},
}

var fields_delete_email_identity_policy = []leanruntime.Field{
	{Name: "EmailIdentity", Flag: "email-identity", Type: "*string", Required: true},
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
}

var fields_delete_email_template = []leanruntime.Field{
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
}

var fields_delete_multi_region_endpoint = []leanruntime.Field{
	{Name: "EndpointName", Flag: "endpoint-name", Type: "*string", Required: true},
}

var fields_delete_suppressed_destination = []leanruntime.Field{
	{Name: "EmailAddress", Flag: "email-address", Type: "*string", Required: true},
}

var fields_delete_tenant = []leanruntime.Field{
	{Name: "TenantName", Flag: "tenant-name", Type: "*string", Required: true},
}

var fields_delete_tenant_resource_association = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TenantName", Flag: "tenant-name", Type: "*string", Required: true},
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

var fields_get_contact = []leanruntime.Field{
	{Name: "ContactListName", Flag: "contact-list-name", Type: "*string", Required: true},
	{Name: "EmailAddress", Flag: "email-address", Type: "*string", Required: true},
}

var fields_get_contact_list = []leanruntime.Field{
	{Name: "ContactListName", Flag: "contact-list-name", Type: "*string", Required: true},
}

var fields_get_custom_verification_email_template = []leanruntime.Field{
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
}

var fields_get_dedicated_ip = []leanruntime.Field{
	{Name: "Ip", Flag: "ip", Type: "*string", Required: true},
}

var fields_get_dedicated_ip_pool = []leanruntime.Field{
	{Name: "PoolName", Flag: "pool-name", Type: "*string", Required: true},
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

var fields_get_email_address_insights = []leanruntime.Field{
	{Name: "EmailAddress", Flag: "email-address", Type: "*string", Required: true},
}

var fields_get_email_identity = []leanruntime.Field{
	{Name: "EmailIdentity", Flag: "email-identity", Type: "*string", Required: true},
}

var fields_get_email_identity_policies = []leanruntime.Field{
	{Name: "EmailIdentity", Flag: "email-identity", Type: "*string", Required: true},
}

var fields_get_email_template = []leanruntime.Field{
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
}

var fields_get_export_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_get_import_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_get_message_insights = []leanruntime.Field{
	{Name: "MessageId", Flag: "message-id", Type: "*string", Required: true},
}

var fields_get_multi_region_endpoint = []leanruntime.Field{
	{Name: "EndpointName", Flag: "endpoint-name", Type: "*string", Required: true},
}

var fields_get_reputation_entity = []leanruntime.Field{
	{Name: "ReputationEntityReference", Flag: "reputation-entity-reference", Type: "*string", Required: true},
	{Name: "ReputationEntityType", Flag: "reputation-entity-type", Type: "types.ReputationEntityType", Required: true},
}

var fields_get_suppressed_destination = []leanruntime.Field{
	{Name: "EmailAddress", Flag: "email-address", Type: "*string", Required: true},
}

var fields_get_tenant = []leanruntime.Field{
	{Name: "TenantName", Flag: "tenant-name", Type: "*string", Required: true},
}

var fields_list_configuration_sets = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
}

var fields_list_contact_lists = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
}

var fields_list_contacts = []leanruntime.Field{
	{Name: "ContactListName", Flag: "contact-list-name", Type: "*string", Required: true},
	{Name: "Filter", Flag: "filter", Type: "*types.ListContactsFilter", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
}

var fields_list_custom_verification_email_templates = []leanruntime.Field{
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

var fields_list_email_templates = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
}

var fields_list_export_jobs = []leanruntime.Field{
	{Name: "ExportSourceType", Flag: "export-source-type", Type: "types.ExportSourceType", Required: false},
	{Name: "JobStatus", Flag: "job-status", Type: "types.JobStatus", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
}

var fields_list_import_jobs = []leanruntime.Field{
	{Name: "ImportDestinationType", Flag: "import-destination-type", Type: "types.ImportDestinationType", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
}

var fields_list_multi_region_endpoints = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
}

var fields_list_recommendations = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "map[string]string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
}

var fields_list_reputation_entities = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "map[string]string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
}

var fields_list_resource_tenants = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_suppressed_destinations = []leanruntime.Field{
	{Name: "EndDate", Flag: "end-date", Type: "*time.Time", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
	{Name: "Reasons", Flag: "reasons", Type: "[]types.SuppressionListReason", Required: false},
	{Name: "StartDate", Flag: "start-date", Type: "*time.Time", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_tenant_resources = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "map[string]string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
	{Name: "TenantName", Flag: "tenant-name", Type: "*string", Required: true},
}

var fields_list_tenants = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
}

var fields_put_account_dedicated_ip_warmup_attributes = []leanruntime.Field{
	{Name: "AutoWarmupEnabled", Flag: "auto-warmup-enabled", Type: "bool", Required: false},
}

var fields_put_account_details = []leanruntime.Field{
	{Name: "AdditionalContactEmailAddresses", Flag: "additional-contact-email-addresses", Type: "[]string", Required: false},
	{Name: "ContactLanguage", Flag: "contact-language", Type: "types.ContactLanguage", Required: false},
	{Name: "MailType", Flag: "mail-type", Type: "types.MailType", Required: true},
	{Name: "ProductionAccessEnabled", Flag: "production-access-enabled", Type: "*bool", Required: false},
	{Name: "UseCaseDescription", Flag: "use-case-description", Type: "*string", Required: false},
	{Name: "WebsiteURL", Flag: "website-url", Type: "*string", Required: true},
}

var fields_put_account_sending_attributes = []leanruntime.Field{
	{Name: "SendingEnabled", Flag: "sending-enabled", Type: "bool", Required: false},
}

var fields_put_account_suppression_attributes = []leanruntime.Field{
	{Name: "SuppressedReasons", Flag: "suppressed-reasons", Type: "[]types.SuppressionListReason", Required: false},
	{Name: "ValidationAttributes", Flag: "validation-attributes", Type: "*types.SuppressionValidationAttributes", Required: false},
}

var fields_put_account_vdm_attributes = []leanruntime.Field{
	{Name: "VdmAttributes", Flag: "vdm-attributes", Type: "*types.VdmAttributes", Required: true},
}

var fields_put_configuration_set_archiving_options = []leanruntime.Field{
	{Name: "ArchiveArn", Flag: "archive-arn", Type: "*string", Required: false},
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
}

var fields_put_configuration_set_delivery_options = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
	{Name: "MaxDeliverySeconds", Flag: "max-delivery-seconds", Type: "*int64", Required: false},
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

var fields_put_configuration_set_suppression_options = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
	{Name: "SuppressedReasons", Flag: "suppressed-reasons", Type: "[]types.SuppressionListReason", Required: false},
	{Name: "ValidationOptions", Flag: "validation-options", Type: "*types.SuppressionValidationOptions", Required: false},
}

var fields_put_configuration_set_tracking_options = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
	{Name: "CustomRedirectDomain", Flag: "custom-redirect-domain", Type: "*string", Required: false},
	{Name: "HttpsPolicy", Flag: "https-policy", Type: "types.HttpsPolicy", Required: false},
}

var fields_put_configuration_set_vdm_options = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
	{Name: "VdmOptions", Flag: "vdm-options", Type: "*types.VdmOptions", Required: false},
}

var fields_put_dedicated_ip_in_pool = []leanruntime.Field{
	{Name: "DestinationPoolName", Flag: "destination-pool-name", Type: "*string", Required: true},
	{Name: "Ip", Flag: "ip", Type: "*string", Required: true},
}

var fields_put_dedicated_ip_pool_scaling_attributes = []leanruntime.Field{
	{Name: "PoolName", Flag: "pool-name", Type: "*string", Required: true},
	{Name: "ScalingMode", Flag: "scaling-mode", Type: "types.ScalingMode", Required: true},
}

var fields_put_dedicated_ip_warmup_attributes = []leanruntime.Field{
	{Name: "Ip", Flag: "ip", Type: "*string", Required: true},
	{Name: "WarmupPercentage", Flag: "warmup-percentage", Type: "*int32", Required: true},
}

var fields_put_deliverability_dashboard_option = []leanruntime.Field{
	{Name: "DashboardEnabled", Flag: "dashboard-enabled", Type: "bool", Required: true},
	{Name: "SubscribedDomains", Flag: "subscribed-domains", Type: "[]types.DomainDeliverabilityTrackingOption", Required: false},
}

var fields_put_email_identity_configuration_set_attributes = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: false},
	{Name: "EmailIdentity", Flag: "email-identity", Type: "*string", Required: true},
}

var fields_put_email_identity_dkim_attributes = []leanruntime.Field{
	{Name: "EmailIdentity", Flag: "email-identity", Type: "*string", Required: true},
	{Name: "SigningEnabled", Flag: "signing-enabled", Type: "bool", Required: false},
}

var fields_put_email_identity_dkim_signing_attributes = []leanruntime.Field{
	{Name: "EmailIdentity", Flag: "email-identity", Type: "*string", Required: true},
	{Name: "SigningAttributes", Flag: "signing-attributes", Type: "*types.DkimSigningAttributes", Required: false},
	{Name: "SigningAttributesOrigin", Flag: "signing-attributes-origin", Type: "types.DkimSigningAttributesOrigin", Required: true},
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

var fields_put_suppressed_destination = []leanruntime.Field{
	{Name: "EmailAddress", Flag: "email-address", Type: "*string", Required: true},
	{Name: "Reason", Flag: "reason", Type: "types.SuppressionListReason", Required: true},
}

var fields_send_bulk_email = []leanruntime.Field{
	{Name: "BulkEmailEntries", Flag: "bulk-email-entries", Type: "[]types.BulkEmailEntry", Required: true},
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: false},
	{Name: "DefaultContent", Flag: "default-content", Type: "*types.BulkEmailContent", Required: true},
	{Name: "DefaultEmailTags", Flag: "default-email-tags", Type: "[]types.MessageTag", Required: false},
	{Name: "EndpointId", Flag: "endpoint-id", Type: "*string", Required: false},
	{Name: "FeedbackForwardingEmailAddress", Flag: "feedback-forwarding-email-address", Type: "*string", Required: false},
	{Name: "FeedbackForwardingEmailAddressIdentityArn", Flag: "feedback-forwarding-email-address-identity-arn", Type: "*string", Required: false},
	{Name: "FromEmailAddress", Flag: "from-email-address", Type: "*string", Required: false},
	{Name: "FromEmailAddressIdentityArn", Flag: "from-email-address-identity-arn", Type: "*string", Required: false},
	{Name: "ReplyToAddresses", Flag: "reply-to-addresses", Type: "[]string", Required: false},
	{Name: "TenantName", Flag: "tenant-name", Type: "*string", Required: false},
}

var fields_send_custom_verification_email = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: false},
	{Name: "EmailAddress", Flag: "email-address", Type: "*string", Required: true},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
}

var fields_send_email = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: false},
	{Name: "Content", Flag: "content", Type: "*types.EmailContent", Required: true},
	{Name: "Destination", Flag: "destination", Type: "*types.Destination", Required: false},
	{Name: "EmailTags", Flag: "email-tags", Type: "[]types.MessageTag", Required: false},
	{Name: "EndpointId", Flag: "endpoint-id", Type: "*string", Required: false},
	{Name: "FeedbackForwardingEmailAddress", Flag: "feedback-forwarding-email-address", Type: "*string", Required: false},
	{Name: "FeedbackForwardingEmailAddressIdentityArn", Flag: "feedback-forwarding-email-address-identity-arn", Type: "*string", Required: false},
	{Name: "FromEmailAddress", Flag: "from-email-address", Type: "*string", Required: false},
	{Name: "FromEmailAddressIdentityArn", Flag: "from-email-address-identity-arn", Type: "*string", Required: false},
	{Name: "ListManagementOptions", Flag: "list-management-options", Type: "*types.ListManagementOptions", Required: false},
	{Name: "ReplyToAddresses", Flag: "reply-to-addresses", Type: "[]string", Required: false},
	{Name: "TenantName", Flag: "tenant-name", Type: "*string", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_test_render_email_template = []leanruntime.Field{
	{Name: "TemplateData", Flag: "template-data", Type: "*string", Required: true},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
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

var fields_update_contact = []leanruntime.Field{
	{Name: "AttributesData", Flag: "attributes-data", Type: "*string", Required: false},
	{Name: "ContactListName", Flag: "contact-list-name", Type: "*string", Required: true},
	{Name: "EmailAddress", Flag: "email-address", Type: "*string", Required: true},
	{Name: "TopicPreferences", Flag: "topic-preferences", Type: "[]types.TopicPreference", Required: false},
	{Name: "UnsubscribeAll", Flag: "unsubscribe-all", Type: "bool", Required: false},
}

var fields_update_contact_list = []leanruntime.Field{
	{Name: "ContactListName", Flag: "contact-list-name", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Topics", Flag: "topics", Type: "[]types.Topic", Required: false},
}

var fields_update_custom_verification_email_template = []leanruntime.Field{
	{Name: "FailureRedirectionURL", Flag: "failure-redirection-url", Type: "*string", Required: true},
	{Name: "FromEmailAddress", Flag: "from-email-address", Type: "*string", Required: true},
	{Name: "SuccessRedirectionURL", Flag: "success-redirection-url", Type: "*string", Required: true},
	{Name: "TemplateContent", Flag: "template-content", Type: "*string", Required: true},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
	{Name: "TemplateSubject", Flag: "template-subject", Type: "*string", Required: true},
}

var fields_update_email_identity_policy = []leanruntime.Field{
	{Name: "EmailIdentity", Flag: "email-identity", Type: "*string", Required: true},
	{Name: "Policy", Flag: "policy", Type: "*string", Required: true},
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
}

var fields_update_email_template = []leanruntime.Field{
	{Name: "TemplateContent", Flag: "template-content", Type: "*types.EmailTemplateContent", Required: true},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
}

var fields_update_reputation_entity_customer_managed_status = []leanruntime.Field{
	{Name: "ReputationEntityReference", Flag: "reputation-entity-reference", Type: "*string", Required: true},
	{Name: "ReputationEntityType", Flag: "reputation-entity-type", Type: "types.ReputationEntityType", Required: true},
	{Name: "SendingStatus", Flag: "sending-status", Type: "types.SendingStatus", Required: true},
}

var fields_update_reputation_entity_policy = []leanruntime.Field{
	{Name: "ReputationEntityPolicy", Flag: "reputation-entity-policy", Type: "*string", Required: true},
	{Name: "ReputationEntityReference", Flag: "reputation-entity-reference", Type: "*string", Required: true},
	{Name: "ReputationEntityType", Flag: "reputation-entity-type", Type: "types.ReputationEntityType", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"batch-get-metric-data": {
			Name:   "batch-get-metric-data",
			Fields: fields_batch_get_metric_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetMetricDataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_metric_data, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetMetricData(ctx, input)
			},
		},
		"cancel-export-job": {
			Name:   "cancel-export-job",
			Fields: fields_cancel_export_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelExportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_export_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelExportJob(ctx, input)
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
		"create-contact-list": {
			Name:   "create-contact-list",
			Fields: fields_create_contact_list,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateContactListInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_contact_list, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateContactList(ctx, input)
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
		"create-email-identity-policy": {
			Name:   "create-email-identity-policy",
			Fields: fields_create_email_identity_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEmailIdentityPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_email_identity_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEmailIdentityPolicy(ctx, input)
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
		"create-multi-region-endpoint": {
			Name:   "create-multi-region-endpoint",
			Fields: fields_create_multi_region_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMultiRegionEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_multi_region_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMultiRegionEndpoint(ctx, input)
			},
		},
		"create-tenant": {
			Name:   "create-tenant",
			Fields: fields_create_tenant,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTenantInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_tenant, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTenant(ctx, input)
			},
		},
		"create-tenant-resource-association": {
			Name:   "create-tenant-resource-association",
			Fields: fields_create_tenant_resource_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTenantResourceAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_tenant_resource_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTenantResourceAssociation(ctx, input)
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
		"delete-contact-list": {
			Name:   "delete-contact-list",
			Fields: fields_delete_contact_list,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteContactListInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_contact_list, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteContactList(ctx, input)
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
		"delete-email-identity-policy": {
			Name:   "delete-email-identity-policy",
			Fields: fields_delete_email_identity_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEmailIdentityPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_email_identity_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEmailIdentityPolicy(ctx, input)
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
		"delete-multi-region-endpoint": {
			Name:   "delete-multi-region-endpoint",
			Fields: fields_delete_multi_region_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMultiRegionEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_multi_region_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMultiRegionEndpoint(ctx, input)
			},
		},
		"delete-suppressed-destination": {
			Name:   "delete-suppressed-destination",
			Fields: fields_delete_suppressed_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSuppressedDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_suppressed_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSuppressedDestination(ctx, input)
			},
		},
		"delete-tenant": {
			Name:   "delete-tenant",
			Fields: fields_delete_tenant,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTenantInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_tenant, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTenant(ctx, input)
			},
		},
		"delete-tenant-resource-association": {
			Name:   "delete-tenant-resource-association",
			Fields: fields_delete_tenant_resource_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTenantResourceAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_tenant_resource_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTenantResourceAssociation(ctx, input)
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
		"get-contact-list": {
			Name:   "get-contact-list",
			Fields: fields_get_contact_list,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetContactListInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_contact_list, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetContactList(ctx, input)
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
		"get-dedicated-ip-pool": {
			Name:   "get-dedicated-ip-pool",
			Fields: fields_get_dedicated_ip_pool,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDedicatedIpPoolInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_dedicated_ip_pool, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDedicatedIpPool(ctx, input)
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
		"get-email-address-insights": {
			Name:   "get-email-address-insights",
			Fields: fields_get_email_address_insights,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEmailAddressInsightsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_email_address_insights, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEmailAddressInsights(ctx, input)
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
		"get-email-identity-policies": {
			Name:   "get-email-identity-policies",
			Fields: fields_get_email_identity_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEmailIdentityPoliciesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_email_identity_policies, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEmailIdentityPolicies(ctx, input)
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
		"get-message-insights": {
			Name:   "get-message-insights",
			Fields: fields_get_message_insights,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMessageInsightsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_message_insights, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMessageInsights(ctx, input)
			},
		},
		"get-multi-region-endpoint": {
			Name:   "get-multi-region-endpoint",
			Fields: fields_get_multi_region_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMultiRegionEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_multi_region_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMultiRegionEndpoint(ctx, input)
			},
		},
		"get-reputation-entity": {
			Name:   "get-reputation-entity",
			Fields: fields_get_reputation_entity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetReputationEntityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_reputation_entity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetReputationEntity(ctx, input)
			},
		},
		"get-suppressed-destination": {
			Name:   "get-suppressed-destination",
			Fields: fields_get_suppressed_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSuppressedDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_suppressed_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSuppressedDestination(ctx, input)
			},
		},
		"get-tenant": {
			Name:   "get-tenant",
			Fields: fields_get_tenant,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTenantInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_tenant, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTenant(ctx, input)
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
		"list-contact-lists": {
			Name:   "list-contact-lists",
			Fields: fields_list_contact_lists,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListContactListsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_contact_lists, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListContactLists(ctx, input)
				}
				var results []*svc.ListContactListsOutput
				p := svc.NewListContactListsPaginator(client, input)
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
		"list-email-templates": {
			Name:   "list-email-templates",
			Fields: fields_list_email_templates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEmailTemplatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_email_templates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEmailTemplates(ctx, input)
				}
				var results []*svc.ListEmailTemplatesOutput
				p := svc.NewListEmailTemplatesPaginator(client, input)
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
		"list-export-jobs": {
			Name:   "list-export-jobs",
			Fields: fields_list_export_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListExportJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_export_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListExportJobs(ctx, input)
				}
				var results []*svc.ListExportJobsOutput
				p := svc.NewListExportJobsPaginator(client, input)
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
		"list-import-jobs": {
			Name:   "list-import-jobs",
			Fields: fields_list_import_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListImportJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_import_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListImportJobs(ctx, input)
				}
				var results []*svc.ListImportJobsOutput
				p := svc.NewListImportJobsPaginator(client, input)
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
		"list-multi-region-endpoints": {
			Name:   "list-multi-region-endpoints",
			Fields: fields_list_multi_region_endpoints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMultiRegionEndpointsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_multi_region_endpoints, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMultiRegionEndpoints(ctx, input)
				}
				var results []*svc.ListMultiRegionEndpointsOutput
				p := svc.NewListMultiRegionEndpointsPaginator(client, input)
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
		"list-recommendations": {
			Name:   "list-recommendations",
			Fields: fields_list_recommendations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRecommendationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_recommendations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRecommendations(ctx, input)
				}
				var results []*svc.ListRecommendationsOutput
				p := svc.NewListRecommendationsPaginator(client, input)
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
		"list-reputation-entities": {
			Name:   "list-reputation-entities",
			Fields: fields_list_reputation_entities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListReputationEntitiesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_reputation_entities, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListReputationEntities(ctx, input)
				}
				var results []*svc.ListReputationEntitiesOutput
				p := svc.NewListReputationEntitiesPaginator(client, input)
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
		"list-resource-tenants": {
			Name:   "list-resource-tenants",
			Fields: fields_list_resource_tenants,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResourceTenantsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_resource_tenants, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResourceTenants(ctx, input)
				}
				var results []*svc.ListResourceTenantsOutput
				p := svc.NewListResourceTenantsPaginator(client, input)
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
		"list-suppressed-destinations": {
			Name:   "list-suppressed-destinations",
			Fields: fields_list_suppressed_destinations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSuppressedDestinationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_suppressed_destinations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSuppressedDestinations(ctx, input)
				}
				var results []*svc.ListSuppressedDestinationsOutput
				p := svc.NewListSuppressedDestinationsPaginator(client, input)
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
		"list-tenant-resources": {
			Name:   "list-tenant-resources",
			Fields: fields_list_tenant_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTenantResourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_tenant_resources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTenantResources(ctx, input)
				}
				var results []*svc.ListTenantResourcesOutput
				p := svc.NewListTenantResourcesPaginator(client, input)
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
		"list-tenants": {
			Name:   "list-tenants",
			Fields: fields_list_tenants,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTenantsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_tenants, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTenants(ctx, input)
				}
				var results []*svc.ListTenantsOutput
				p := svc.NewListTenantsPaginator(client, input)
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
		"put-account-details": {
			Name:   "put-account-details",
			Fields: fields_put_account_details,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutAccountDetailsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_account_details, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutAccountDetails(ctx, input)
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
		"put-account-suppression-attributes": {
			Name:   "put-account-suppression-attributes",
			Fields: fields_put_account_suppression_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutAccountSuppressionAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_account_suppression_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutAccountSuppressionAttributes(ctx, input)
			},
		},
		"put-account-vdm-attributes": {
			Name:   "put-account-vdm-attributes",
			Fields: fields_put_account_vdm_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutAccountVdmAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_account_vdm_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutAccountVdmAttributes(ctx, input)
			},
		},
		"put-configuration-set-archiving-options": {
			Name:   "put-configuration-set-archiving-options",
			Fields: fields_put_configuration_set_archiving_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutConfigurationSetArchivingOptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_configuration_set_archiving_options, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutConfigurationSetArchivingOptions(ctx, input)
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
		"put-configuration-set-suppression-options": {
			Name:   "put-configuration-set-suppression-options",
			Fields: fields_put_configuration_set_suppression_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutConfigurationSetSuppressionOptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_configuration_set_suppression_options, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutConfigurationSetSuppressionOptions(ctx, input)
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
		"put-configuration-set-vdm-options": {
			Name:   "put-configuration-set-vdm-options",
			Fields: fields_put_configuration_set_vdm_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutConfigurationSetVdmOptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_configuration_set_vdm_options, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutConfigurationSetVdmOptions(ctx, input)
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
		"put-dedicated-ip-pool-scaling-attributes": {
			Name:   "put-dedicated-ip-pool-scaling-attributes",
			Fields: fields_put_dedicated_ip_pool_scaling_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutDedicatedIpPoolScalingAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_dedicated_ip_pool_scaling_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutDedicatedIpPoolScalingAttributes(ctx, input)
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
		"put-email-identity-configuration-set-attributes": {
			Name:   "put-email-identity-configuration-set-attributes",
			Fields: fields_put_email_identity_configuration_set_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutEmailIdentityConfigurationSetAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_email_identity_configuration_set_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutEmailIdentityConfigurationSetAttributes(ctx, input)
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
		"put-email-identity-dkim-signing-attributes": {
			Name:   "put-email-identity-dkim-signing-attributes",
			Fields: fields_put_email_identity_dkim_signing_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutEmailIdentityDkimSigningAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_email_identity_dkim_signing_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutEmailIdentityDkimSigningAttributes(ctx, input)
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
		"put-suppressed-destination": {
			Name:   "put-suppressed-destination",
			Fields: fields_put_suppressed_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutSuppressedDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_suppressed_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutSuppressedDestination(ctx, input)
			},
		},
		"send-bulk-email": {
			Name:   "send-bulk-email",
			Fields: fields_send_bulk_email,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendBulkEmailInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_bulk_email, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendBulkEmail(ctx, input)
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
		"test-render-email-template": {
			Name:   "test-render-email-template",
			Fields: fields_test_render_email_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TestRenderEmailTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_test_render_email_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TestRenderEmailTemplate(ctx, input)
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
		"update-contact-list": {
			Name:   "update-contact-list",
			Fields: fields_update_contact_list,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateContactListInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_contact_list, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateContactList(ctx, input)
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
		"update-email-identity-policy": {
			Name:   "update-email-identity-policy",
			Fields: fields_update_email_identity_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEmailIdentityPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_email_identity_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEmailIdentityPolicy(ctx, input)
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
		"update-reputation-entity-customer-managed-status": {
			Name:   "update-reputation-entity-customer-managed-status",
			Fields: fields_update_reputation_entity_customer_managed_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateReputationEntityCustomerManagedStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_reputation_entity_customer_managed_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateReputationEntityCustomerManagedStatus(ctx, input)
			},
		},
		"update-reputation-entity-policy": {
			Name:   "update-reputation-entity-policy",
			Fields: fields_update_reputation_entity_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateReputationEntityPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_reputation_entity_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateReputationEntityPolicy(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("sesv2", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
