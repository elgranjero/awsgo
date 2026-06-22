package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/connect"
)

var fields_activate_evaluation_form = []leanruntime.Field{
	{Name: "EvaluationFormId", Flag: "evaluation-form-id", Type: "*string", Required: true},
	{Name: "EvaluationFormVersion", Flag: "evaluation-form-version", Type: "int32", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_associate_analytics_data_set = []leanruntime.Field{
	{Name: "DataSetId", Flag: "data-set-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "TargetAccountId", Flag: "target-account-id", Type: "*string", Required: false},
}

var fields_associate_approved_origin = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Origin", Flag: "origin", Type: "*string", Required: true},
}

var fields_associate_bot = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "LexBot", Flag: "lex-bot", Type: "*types.LexBot", Required: false},
	{Name: "LexV2Bot", Flag: "lex-v2-bot", Type: "*types.LexV2Bot", Required: false},
}

var fields_associate_contact_with_user = []leanruntime.Field{
	{Name: "ContactId", Flag: "contact-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_associate_default_vocabulary = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "LanguageCode", Flag: "language-code", Type: "types.VocabularyLanguageCode", Required: true},
	{Name: "VocabularyId", Flag: "vocabulary-id", Type: "*string", Required: false},
}

var fields_associate_email_address_alias = []leanruntime.Field{
	{Name: "AliasConfiguration", Flag: "alias-configuration", Type: "*types.AliasConfiguration", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "EmailAddressId", Flag: "email-address-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_associate_flow = []leanruntime.Field{
	{Name: "FlowId", Flag: "flow-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.FlowAssociationResourceType", Required: true},
}

var fields_associate_hours_of_operations = []leanruntime.Field{
	{Name: "HoursOfOperationId", Flag: "hours-of-operation-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "ParentHoursOfOperationConfigs", Flag: "parent-hours-of-operation-configs", Type: "[]types.ParentHoursOfOperationConfig", Required: true},
}

var fields_associate_instance_storage_config = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.InstanceStorageResourceType", Required: true},
	{Name: "StorageConfig", Flag: "storage-config", Type: "*types.InstanceStorageConfig", Required: true},
}

var fields_associate_lambda_function = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "FunctionArn", Flag: "function-arn", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_associate_lex_bot = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "LexBot", Flag: "lex-bot", Type: "*types.LexBot", Required: true},
}

var fields_associate_phone_number_contact_flow = []leanruntime.Field{
	{Name: "ContactFlowId", Flag: "contact-flow-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "PhoneNumberId", Flag: "phone-number-id", Type: "*string", Required: true},
}

var fields_associate_queue_quick_connects = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
	{Name: "QuickConnectIds", Flag: "quick-connect-ids", Type: "[]string", Required: true},
}

var fields_associate_routing_profile_queues = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "ManualAssignmentQueueConfigs", Flag: "manual-assignment-queue-configs", Type: "[]types.RoutingProfileManualAssignmentQueueConfig", Required: false},
	{Name: "QueueConfigs", Flag: "queue-configs", Type: "[]types.RoutingProfileQueueConfig", Required: false},
	{Name: "RoutingProfileId", Flag: "routing-profile-id", Type: "*string", Required: true},
}

var fields_associate_security_key = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Key", Flag: "key", Type: "*string", Required: true},
}

var fields_associate_security_profiles = []leanruntime.Field{
	{Name: "EntityArn", Flag: "entity-arn", Type: "*string", Required: true},
	{Name: "EntityType", Flag: "entity-type", Type: "types.EntityType", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "SecurityProfiles", Flag: "security-profiles", Type: "[]types.SecurityProfileItem", Required: true},
}

var fields_associate_traffic_distribution_group_user = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "TrafficDistributionGroupId", Flag: "traffic-distribution-group-id", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_associate_user_proficiencies = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
	{Name: "UserProficiencies", Flag: "user-proficiencies", Type: "[]types.UserProficiency", Required: true},
}

var fields_associate_workspace = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "ResourceArns", Flag: "resource-arns", Type: "[]string", Required: true},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_batch_associate_analytics_data_set = []leanruntime.Field{
	{Name: "DataSetIds", Flag: "data-set-ids", Type: "[]string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "TargetAccountId", Flag: "target-account-id", Type: "*string", Required: false},
}

var fields_batch_create_data_table_value = []leanruntime.Field{
	{Name: "DataTableId", Flag: "data-table-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Values", Flag: "values", Type: "[]types.DataTableValue", Required: true},
}

var fields_batch_delete_data_table_value = []leanruntime.Field{
	{Name: "DataTableId", Flag: "data-table-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Values", Flag: "values", Type: "[]types.DataTableDeleteValueIdentifier", Required: true},
}

var fields_batch_describe_data_table_value = []leanruntime.Field{
	{Name: "DataTableId", Flag: "data-table-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Values", Flag: "values", Type: "[]types.DataTableValueIdentifier", Required: true},
}

var fields_batch_disassociate_analytics_data_set = []leanruntime.Field{
	{Name: "DataSetIds", Flag: "data-set-ids", Type: "[]string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "TargetAccountId", Flag: "target-account-id", Type: "*string", Required: false},
}

var fields_batch_get_attached_file_metadata = []leanruntime.Field{
	{Name: "AssociatedResourceArn", Flag: "associated-resource-arn", Type: "*string", Required: true},
	{Name: "FileIds", Flag: "file-ids", Type: "[]string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_batch_get_flow_association = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "ResourceIds", Flag: "resource-ids", Type: "[]string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.ListFlowAssociationResourceType", Required: false},
}

var fields_batch_put_contact = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ContactDataRequestList", Flag: "contact-data-request-list", Type: "[]types.ContactDataRequest", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_batch_update_data_table_value = []leanruntime.Field{
	{Name: "DataTableId", Flag: "data-table-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Values", Flag: "values", Type: "[]types.DataTableValue", Required: true},
}

var fields_claim_phone_number = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: false},
	{Name: "PhoneNumber", Flag: "phone-number", Type: "*string", Required: true},
	{Name: "PhoneNumberDescription", Flag: "phone-number-description", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TargetArn", Flag: "target-arn", Type: "*string", Required: false},
}

var fields_complete_attached_file_upload = []leanruntime.Field{
	{Name: "AssociatedResourceArn", Flag: "associated-resource-arn", Type: "*string", Required: true},
	{Name: "FileId", Flag: "file-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_create_agent_status = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayOrder", Flag: "display-order", Type: "*int32", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "State", Flag: "state", Type: "types.AgentStatusState", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_contact = []leanruntime.Field{
	{Name: "Attributes", Flag: "attributes", Type: "map[string]string", Required: false},
	{Name: "Channel", Flag: "channel", Type: "types.Channel", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ExpiryDurationInMinutes", Flag: "expiry-duration-in-minutes", Type: "*int32", Required: false},
	{Name: "InitiateAs", Flag: "initiate-as", Type: "types.InitiateAs", Required: false},
	{Name: "InitiationMethod", Flag: "initiation-method", Type: "types.ContactInitiationMethod", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "PreviousContactId", Flag: "previous-contact-id", Type: "*string", Required: false},
	{Name: "References", Flag: "references", Type: "map[string]types.Reference", Required: false},
	{Name: "RelatedContactId", Flag: "related-contact-id", Type: "*string", Required: false},
	{Name: "SegmentAttributes", Flag: "segment-attributes", Type: "map[string]types.SegmentAttributeValue", Required: false},
	{Name: "UserInfo", Flag: "user-info", Type: "*types.UserInfo", Required: false},
}

var fields_create_contact_flow = []leanruntime.Field{
	{Name: "Content", Flag: "content", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.ContactFlowStatus", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.ContactFlowType", Required: true},
}

var fields_create_contact_flow_module = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Content", Flag: "content", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ExternalInvocationConfiguration", Flag: "external-invocation-configuration", Type: "*types.ExternalInvocationConfiguration", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Settings", Flag: "settings", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_contact_flow_module_alias = []leanruntime.Field{
	{Name: "AliasName", Flag: "alias-name", Type: "*string", Required: true},
	{Name: "ContactFlowModuleId", Flag: "contact-flow-module-id", Type: "*string", Required: true},
	{Name: "ContactFlowModuleVersion", Flag: "contact-flow-module-version", Type: "*int64", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_create_contact_flow_module_version = []leanruntime.Field{
	{Name: "ContactFlowModuleId", Flag: "contact-flow-module-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "FlowModuleContentSha256", Flag: "flow-module-content-sha256", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_create_contact_flow_version = []leanruntime.Field{
	{Name: "ContactFlowId", Flag: "contact-flow-id", Type: "*string", Required: true},
	{Name: "ContactFlowVersion", Flag: "contact-flow-version", Type: "*int64", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "FlowContentSha256", Flag: "flow-content-sha256", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "LastModifiedRegion", Flag: "last-modified-region", Type: "*string", Required: false},
	{Name: "LastModifiedTime", Flag: "last-modified-time", Type: "*time.Time", Required: false},
}

var fields_create_data_table = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.DataTableStatus", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TimeZone", Flag: "time-zone", Type: "*string", Required: true},
	{Name: "ValueLockLevel", Flag: "value-lock-level", Type: "types.DataTableLockLevel", Required: true},
}

var fields_create_data_table_attribute = []leanruntime.Field{
	{Name: "DataTableId", Flag: "data-table-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Primary", Flag: "primary", Type: "bool", Required: false},
	{Name: "Validation", Flag: "validation", Type: "*types.Validation", Required: false},
	{Name: "ValueType", Flag: "value-type", Type: "types.DataTableAttributeValueType", Required: true},
}

var fields_create_email_address = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "EmailAddress", Flag: "email-address", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_evaluation_form = []leanruntime.Field{
	{Name: "AsDraft", Flag: "as-draft", Type: "bool", Required: false},
	{Name: "AutoEvaluationConfiguration", Flag: "auto-evaluation-configuration", Type: "*types.EvaluationFormAutoEvaluationConfiguration", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Items", Flag: "items", Type: "[]types.EvaluationFormItem", Required: true},
	{Name: "LanguageConfiguration", Flag: "language-configuration", Type: "*types.EvaluationFormLanguageConfiguration", Required: false},
	{Name: "ReviewConfiguration", Flag: "review-configuration", Type: "*types.EvaluationReviewConfiguration", Required: false},
	{Name: "ScoringStrategy", Flag: "scoring-strategy", Type: "*types.EvaluationFormScoringStrategy", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TargetConfiguration", Flag: "target-configuration", Type: "*types.EvaluationFormTargetConfiguration", Required: false},
	{Name: "Title", Flag: "title", Type: "*string", Required: true},
}

var fields_create_hours_of_operation = []leanruntime.Field{
	{Name: "Config", Flag: "config", Type: "[]types.HoursOfOperationConfig", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ParentHoursOfOperationConfigs", Flag: "parent-hours-of-operation-configs", Type: "[]types.ParentHoursOfOperationConfig", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TimeZone", Flag: "time-zone", Type: "*string", Required: true},
}

var fields_create_hours_of_operation_override = []leanruntime.Field{
	{Name: "Config", Flag: "config", Type: "[]types.HoursOfOperationOverrideConfig", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EffectiveFrom", Flag: "effective-from", Type: "*string", Required: true},
	{Name: "EffectiveTill", Flag: "effective-till", Type: "*string", Required: true},
	{Name: "HoursOfOperationId", Flag: "hours-of-operation-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OverrideType", Flag: "override-type", Type: "types.OverrideType", Required: false},
	{Name: "RecurrenceConfig", Flag: "recurrence-config", Type: "*types.RecurrenceConfig", Required: false},
}

var fields_create_instance = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: false},
	{Name: "IdentityManagementType", Flag: "identity-management-type", Type: "types.DirectoryType", Required: true},
	{Name: "InboundCallsEnabled", Flag: "inbound-calls-enabled", Type: "*bool", Required: true},
	{Name: "InstanceAlias", Flag: "instance-alias", Type: "*string", Required: false},
	{Name: "OutboundCallsEnabled", Flag: "outbound-calls-enabled", Type: "*bool", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_integration_association = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "IntegrationArn", Flag: "integration-arn", Type: "*string", Required: true},
	{Name: "IntegrationType", Flag: "integration-type", Type: "types.IntegrationType", Required: true},
	{Name: "SourceApplicationName", Flag: "source-application-name", Type: "*string", Required: false},
	{Name: "SourceApplicationUrl", Flag: "source-application-url", Type: "*string", Required: false},
	{Name: "SourceType", Flag: "source-type", Type: "types.SourceType", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_notification = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Content", Flag: "content", Type: "map[string]string", Required: true},
	{Name: "ExpiresAt", Flag: "expires-at", Type: "*time.Time", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "PredefinedNotificationId", Flag: "predefined-notification-id", Type: "*string", Required: false},
	{Name: "Priority", Flag: "priority", Type: "types.ConfigurableNotificationPriority", Required: false},
	{Name: "Recipients", Flag: "recipients", Type: "[]string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_participant = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ContactId", Flag: "contact-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "ParticipantDetails", Flag: "participant-details", Type: "*types.ParticipantDetailsToAdd", Required: true},
}

var fields_create_persistent_contact_association = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "InitialContactId", Flag: "initial-contact-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "RehydrationType", Flag: "rehydration-type", Type: "types.RehydrationType", Required: true},
	{Name: "SourceContactId", Flag: "source-contact-id", Type: "*string", Required: true},
}

var fields_create_predefined_attribute = []leanruntime.Field{
	{Name: "AttributeConfiguration", Flag: "attribute-configuration", Type: "*types.InputPredefinedAttributeConfiguration", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Purposes", Flag: "purposes", Type: "[]string", Required: false},
	{Name: "Values", Flag: "values", Type: "types.PredefinedAttributeValues", Required: false},
}

var fields_create_prompt = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "S3Uri", Flag: "s3-uri", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_push_notification_registration = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ContactConfiguration", Flag: "contact-configuration", Type: "*types.ContactConfiguration", Required: true},
	{Name: "DeviceToken", Flag: "device-token", Type: "*string", Required: true},
	{Name: "DeviceType", Flag: "device-type", Type: "types.DeviceType", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "PinpointAppArn", Flag: "pinpoint-app-arn", Type: "*string", Required: true},
}

var fields_create_queue = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "HoursOfOperationId", Flag: "hours-of-operation-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxContacts", Flag: "max-contacts", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OutboundCallerConfig", Flag: "outbound-caller-config", Type: "*types.OutboundCallerConfig", Required: false},
	{Name: "OutboundEmailConfig", Flag: "outbound-email-config", Type: "*types.OutboundEmailConfig", Required: false},
	{Name: "QuickConnectIds", Flag: "quick-connect-ids", Type: "[]string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_quick_connect = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "QuickConnectConfig", Flag: "quick-connect-config", Type: "*types.QuickConnectConfig", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_routing_profile = []leanruntime.Field{
	{Name: "AgentAvailabilityTimer", Flag: "agent-availability-timer", Type: "types.AgentAvailabilityTimer", Required: false},
	{Name: "DefaultOutboundQueueId", Flag: "default-outbound-queue-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "ManualAssignmentQueueConfigs", Flag: "manual-assignment-queue-configs", Type: "[]types.RoutingProfileManualAssignmentQueueConfig", Required: false},
	{Name: "MediaConcurrencies", Flag: "media-concurrencies", Type: "[]types.MediaConcurrency", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "QueueConfigs", Flag: "queue-configs", Type: "[]types.RoutingProfileQueueConfig", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_rule = []leanruntime.Field{
	{Name: "Actions", Flag: "actions", Type: "[]types.RuleAction", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Function", Flag: "function", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "PublishStatus", Flag: "publish-status", Type: "types.RulePublishStatus", Required: true},
	{Name: "TriggerEventSource", Flag: "trigger-event-source", Type: "*types.RuleTriggerEventSource", Required: true},
}

var fields_create_security_profile = []leanruntime.Field{
	{Name: "AllowedAccessControlHierarchyGroupId", Flag: "allowed-access-control-hierarchy-group-id", Type: "*string", Required: false},
	{Name: "AllowedAccessControlTags", Flag: "allowed-access-control-tags", Type: "map[string]string", Required: false},
	{Name: "AllowedFlowModules", Flag: "allowed-flow-modules", Type: "[]types.FlowModule", Required: false},
	{Name: "Applications", Flag: "applications", Type: "[]types.Application", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "GranularAccessControlConfiguration", Flag: "granular-access-control-configuration", Type: "*types.GranularAccessControlConfiguration", Required: false},
	{Name: "HierarchyRestrictedResources", Flag: "hierarchy-restricted-resources", Type: "[]string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Permissions", Flag: "permissions", Type: "[]string", Required: false},
	{Name: "SecurityProfileName", Flag: "security-profile-name", Type: "*string", Required: true},
	{Name: "TagRestrictedResources", Flag: "tag-restricted-resources", Type: "[]string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_task_template = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Constraints", Flag: "constraints", Type: "*types.TaskTemplateConstraints", Required: false},
	{Name: "ContactFlowId", Flag: "contact-flow-id", Type: "*string", Required: false},
	{Name: "Defaults", Flag: "defaults", Type: "*types.TaskTemplateDefaults", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Fields", Flag: "fields", Type: "[]types.TaskTemplateField", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "SelfAssignFlowId", Flag: "self-assign-flow-id", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.TaskTemplateStatus", Required: false},
}

var fields_create_test_case = []leanruntime.Field{
	{Name: "Content", Flag: "content", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EntryPoint", Flag: "entry-point", Type: "*types.TestCaseEntryPoint", Required: false},
	{Name: "InitializationData", Flag: "initialization-data", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "LastModifiedRegion", Flag: "last-modified-region", Type: "*string", Required: false},
	{Name: "LastModifiedTime", Flag: "last-modified-time", Type: "*time.Time", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.TestCaseStatus", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TestCaseId", Flag: "test-case-id", Type: "*string", Required: false},
}

var fields_create_traffic_distribution_group = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_use_case = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "IntegrationAssociationId", Flag: "integration-association-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "UseCaseType", Flag: "use-case-type", Type: "types.UseCaseType", Required: true},
}

var fields_create_user = []leanruntime.Field{
	{Name: "AfterContactWorkConfigs", Flag: "after-contact-work-configs", Type: "[]types.AfterContactWorkConfigPerChannel", Required: false},
	{Name: "AutoAcceptConfigs", Flag: "auto-accept-configs", Type: "[]types.AutoAcceptConfig", Required: false},
	{Name: "DirectoryUserId", Flag: "directory-user-id", Type: "*string", Required: false},
	{Name: "HierarchyGroupId", Flag: "hierarchy-group-id", Type: "*string", Required: false},
	{Name: "IdentityInfo", Flag: "identity-info", Type: "*types.UserIdentityInfo", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Password", Flag: "password", Type: "*string", Required: false},
	{Name: "PersistentConnectionConfigs", Flag: "persistent-connection-configs", Type: "[]types.PersistentConnectionConfig", Required: false},
	{Name: "PhoneConfig", Flag: "phone-config", Type: "*types.UserPhoneConfig", Required: false},
	{Name: "PhoneNumberConfigs", Flag: "phone-number-configs", Type: "[]types.PhoneNumberConfig", Required: false},
	{Name: "RoutingProfileId", Flag: "routing-profile-id", Type: "*string", Required: true},
	{Name: "SecurityProfileIds", Flag: "security-profile-ids", Type: "[]string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Username", Flag: "username", Type: "*string", Required: true},
	{Name: "VoiceEnhancementConfigs", Flag: "voice-enhancement-configs", Type: "[]types.VoiceEnhancementConfig", Required: false},
}

var fields_create_user_hierarchy_group = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ParentGroupId", Flag: "parent-group-id", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_view = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Content", Flag: "content", Type: "*types.ViewInputContent", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.ViewStatus", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_view_version = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "VersionDescription", Flag: "version-description", Type: "*string", Required: false},
	{Name: "ViewContentSha256", Flag: "view-content-sha256", Type: "*string", Required: false},
	{Name: "ViewId", Flag: "view-id", Type: "*string", Required: true},
}

var fields_create_vocabulary = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Content", Flag: "content", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "LanguageCode", Flag: "language-code", Type: "types.VocabularyLanguageCode", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "VocabularyName", Flag: "vocabulary-name", Type: "*string", Required: true},
}

var fields_create_workspace = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Theme", Flag: "theme", Type: "*types.WorkspaceTheme", Required: false},
	{Name: "Title", Flag: "title", Type: "*string", Required: false},
}

var fields_create_workspace_page = []leanruntime.Field{
	{Name: "InputData", Flag: "input-data", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Page", Flag: "page", Type: "*string", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Slug", Flag: "slug", Type: "*string", Required: false},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_deactivate_evaluation_form = []leanruntime.Field{
	{Name: "EvaluationFormId", Flag: "evaluation-form-id", Type: "*string", Required: true},
	{Name: "EvaluationFormVersion", Flag: "evaluation-form-version", Type: "int32", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_delete_attached_file = []leanruntime.Field{
	{Name: "AssociatedResourceArn", Flag: "associated-resource-arn", Type: "*string", Required: true},
	{Name: "FileId", Flag: "file-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_delete_contact_evaluation = []leanruntime.Field{
	{Name: "EvaluationId", Flag: "evaluation-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_delete_contact_flow = []leanruntime.Field{
	{Name: "ContactFlowId", Flag: "contact-flow-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_delete_contact_flow_module = []leanruntime.Field{
	{Name: "ContactFlowModuleId", Flag: "contact-flow-module-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_delete_contact_flow_module_alias = []leanruntime.Field{
	{Name: "AliasId", Flag: "alias-id", Type: "*string", Required: true},
	{Name: "ContactFlowModuleId", Flag: "contact-flow-module-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_delete_contact_flow_module_version = []leanruntime.Field{
	{Name: "ContactFlowModuleId", Flag: "contact-flow-module-id", Type: "*string", Required: true},
	{Name: "ContactFlowModuleVersion", Flag: "contact-flow-module-version", Type: "*int64", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_delete_contact_flow_version = []leanruntime.Field{
	{Name: "ContactFlowId", Flag: "contact-flow-id", Type: "*string", Required: true},
	{Name: "ContactFlowVersion", Flag: "contact-flow-version", Type: "*int64", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_delete_data_table = []leanruntime.Field{
	{Name: "DataTableId", Flag: "data-table-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_delete_data_table_attribute = []leanruntime.Field{
	{Name: "AttributeName", Flag: "attribute-name", Type: "*string", Required: true},
	{Name: "DataTableId", Flag: "data-table-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_delete_email_address = []leanruntime.Field{
	{Name: "EmailAddressId", Flag: "email-address-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_delete_evaluation_form = []leanruntime.Field{
	{Name: "EvaluationFormId", Flag: "evaluation-form-id", Type: "*string", Required: true},
	{Name: "EvaluationFormVersion", Flag: "evaluation-form-version", Type: "*int32", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_delete_hours_of_operation = []leanruntime.Field{
	{Name: "HoursOfOperationId", Flag: "hours-of-operation-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_delete_hours_of_operation_override = []leanruntime.Field{
	{Name: "HoursOfOperationId", Flag: "hours-of-operation-id", Type: "*string", Required: true},
	{Name: "HoursOfOperationOverrideId", Flag: "hours-of-operation-override-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_delete_instance = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_delete_integration_association = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "IntegrationAssociationId", Flag: "integration-association-id", Type: "*string", Required: true},
}

var fields_delete_notification = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "NotificationId", Flag: "notification-id", Type: "*string", Required: true},
}

var fields_delete_predefined_attribute = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_prompt = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "PromptId", Flag: "prompt-id", Type: "*string", Required: true},
}

var fields_delete_push_notification_registration = []leanruntime.Field{
	{Name: "ContactId", Flag: "contact-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "RegistrationId", Flag: "registration-id", Type: "*string", Required: true},
}

var fields_delete_queue = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
}

var fields_delete_quick_connect = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "QuickConnectId", Flag: "quick-connect-id", Type: "*string", Required: true},
}

var fields_delete_routing_profile = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "RoutingProfileId", Flag: "routing-profile-id", Type: "*string", Required: true},
}

var fields_delete_rule = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "RuleId", Flag: "rule-id", Type: "*string", Required: true},
}

var fields_delete_security_profile = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "SecurityProfileId", Flag: "security-profile-id", Type: "*string", Required: true},
}

var fields_delete_task_template = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "TaskTemplateId", Flag: "task-template-id", Type: "*string", Required: true},
}

var fields_delete_test_case = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "TestCaseId", Flag: "test-case-id", Type: "*string", Required: true},
}

var fields_delete_traffic_distribution_group = []leanruntime.Field{
	{Name: "TrafficDistributionGroupId", Flag: "traffic-distribution-group-id", Type: "*string", Required: true},
}

var fields_delete_use_case = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "IntegrationAssociationId", Flag: "integration-association-id", Type: "*string", Required: true},
	{Name: "UseCaseId", Flag: "use-case-id", Type: "*string", Required: true},
}

var fields_delete_user = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_delete_user_hierarchy_group = []leanruntime.Field{
	{Name: "HierarchyGroupId", Flag: "hierarchy-group-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_delete_view = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "ViewId", Flag: "view-id", Type: "*string", Required: true},
}

var fields_delete_view_version = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "ViewId", Flag: "view-id", Type: "*string", Required: true},
	{Name: "ViewVersion", Flag: "view-version", Type: "*int32", Required: true},
}

var fields_delete_vocabulary = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "VocabularyId", Flag: "vocabulary-id", Type: "*string", Required: true},
}

var fields_delete_workspace = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_delete_workspace_media = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MediaType", Flag: "media-type", Type: "types.MediaType", Required: true},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_delete_workspace_page = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Page", Flag: "page", Type: "*string", Required: true},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_describe_agent_status = []leanruntime.Field{
	{Name: "AgentStatusId", Flag: "agent-status-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_describe_authentication_profile = []leanruntime.Field{
	{Name: "AuthenticationProfileId", Flag: "authentication-profile-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_describe_contact = []leanruntime.Field{
	{Name: "ContactId", Flag: "contact-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_describe_contact_evaluation = []leanruntime.Field{
	{Name: "EvaluationId", Flag: "evaluation-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_describe_contact_flow = []leanruntime.Field{
	{Name: "ContactFlowId", Flag: "contact-flow-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_describe_contact_flow_module = []leanruntime.Field{
	{Name: "ContactFlowModuleId", Flag: "contact-flow-module-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_describe_contact_flow_module_alias = []leanruntime.Field{
	{Name: "AliasId", Flag: "alias-id", Type: "*string", Required: true},
	{Name: "ContactFlowModuleId", Flag: "contact-flow-module-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_describe_data_table = []leanruntime.Field{
	{Name: "DataTableId", Flag: "data-table-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_describe_data_table_attribute = []leanruntime.Field{
	{Name: "AttributeName", Flag: "attribute-name", Type: "*string", Required: true},
	{Name: "DataTableId", Flag: "data-table-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_describe_email_address = []leanruntime.Field{
	{Name: "EmailAddressId", Flag: "email-address-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_describe_evaluation_form = []leanruntime.Field{
	{Name: "EvaluationFormId", Flag: "evaluation-form-id", Type: "*string", Required: true},
	{Name: "EvaluationFormVersion", Flag: "evaluation-form-version", Type: "*int32", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_describe_hours_of_operation = []leanruntime.Field{
	{Name: "HoursOfOperationId", Flag: "hours-of-operation-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_describe_hours_of_operation_override = []leanruntime.Field{
	{Name: "HoursOfOperationId", Flag: "hours-of-operation-id", Type: "*string", Required: true},
	{Name: "HoursOfOperationOverrideId", Flag: "hours-of-operation-override-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_describe_instance = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_describe_instance_attribute = []leanruntime.Field{
	{Name: "AttributeType", Flag: "attribute-type", Type: "types.InstanceAttributeType", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_describe_instance_storage_config = []leanruntime.Field{
	{Name: "AssociationId", Flag: "association-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.InstanceStorageResourceType", Required: true},
}

var fields_describe_notification = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "NotificationId", Flag: "notification-id", Type: "*string", Required: true},
}

var fields_describe_phone_number = []leanruntime.Field{
	{Name: "PhoneNumberId", Flag: "phone-number-id", Type: "*string", Required: true},
}

var fields_describe_predefined_attribute = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_describe_prompt = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "PromptId", Flag: "prompt-id", Type: "*string", Required: true},
}

var fields_describe_queue = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
}

var fields_describe_quick_connect = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "QuickConnectId", Flag: "quick-connect-id", Type: "*string", Required: true},
}

var fields_describe_routing_profile = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "RoutingProfileId", Flag: "routing-profile-id", Type: "*string", Required: true},
}

var fields_describe_rule = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "RuleId", Flag: "rule-id", Type: "*string", Required: true},
}

var fields_describe_security_profile = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "SecurityProfileId", Flag: "security-profile-id", Type: "*string", Required: true},
}

var fields_describe_test_case = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.TestCaseStatus", Required: false},
	{Name: "TestCaseId", Flag: "test-case-id", Type: "*string", Required: true},
}

var fields_describe_traffic_distribution_group = []leanruntime.Field{
	{Name: "TrafficDistributionGroupId", Flag: "traffic-distribution-group-id", Type: "*string", Required: true},
}

var fields_describe_user = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_describe_user_hierarchy_group = []leanruntime.Field{
	{Name: "HierarchyGroupId", Flag: "hierarchy-group-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_describe_user_hierarchy_structure = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_describe_view = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "ViewId", Flag: "view-id", Type: "*string", Required: true},
}

var fields_describe_vocabulary = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "VocabularyId", Flag: "vocabulary-id", Type: "*string", Required: true},
}

var fields_describe_workspace = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_disassociate_analytics_data_set = []leanruntime.Field{
	{Name: "DataSetId", Flag: "data-set-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "TargetAccountId", Flag: "target-account-id", Type: "*string", Required: false},
}

var fields_disassociate_approved_origin = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Origin", Flag: "origin", Type: "*string", Required: true},
}

var fields_disassociate_bot = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "LexBot", Flag: "lex-bot", Type: "*types.LexBot", Required: false},
	{Name: "LexV2Bot", Flag: "lex-v2-bot", Type: "*types.LexV2Bot", Required: false},
}

var fields_disassociate_email_address_alias = []leanruntime.Field{
	{Name: "AliasConfiguration", Flag: "alias-configuration", Type: "*types.AliasConfiguration", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "EmailAddressId", Flag: "email-address-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_disassociate_flow = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.FlowAssociationResourceType", Required: true},
}

var fields_disassociate_hours_of_operations = []leanruntime.Field{
	{Name: "HoursOfOperationId", Flag: "hours-of-operation-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "ParentHoursOfOperationIds", Flag: "parent-hours-of-operation-ids", Type: "[]string", Required: true},
}

var fields_disassociate_instance_storage_config = []leanruntime.Field{
	{Name: "AssociationId", Flag: "association-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.InstanceStorageResourceType", Required: true},
}

var fields_disassociate_lambda_function = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "FunctionArn", Flag: "function-arn", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_disassociate_lex_bot = []leanruntime.Field{
	{Name: "BotName", Flag: "bot-name", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "LexRegion", Flag: "lex-region", Type: "*string", Required: true},
}

var fields_disassociate_phone_number_contact_flow = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "PhoneNumberId", Flag: "phone-number-id", Type: "*string", Required: true},
}

var fields_disassociate_queue_quick_connects = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
	{Name: "QuickConnectIds", Flag: "quick-connect-ids", Type: "[]string", Required: true},
}

var fields_disassociate_routing_profile_queues = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "ManualAssignmentQueueReferences", Flag: "manual-assignment-queue-references", Type: "[]types.RoutingProfileQueueReference", Required: false},
	{Name: "QueueReferences", Flag: "queue-references", Type: "[]types.RoutingProfileQueueReference", Required: false},
	{Name: "RoutingProfileId", Flag: "routing-profile-id", Type: "*string", Required: true},
}

var fields_disassociate_security_key = []leanruntime.Field{
	{Name: "AssociationId", Flag: "association-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_disassociate_security_profiles = []leanruntime.Field{
	{Name: "EntityArn", Flag: "entity-arn", Type: "*string", Required: true},
	{Name: "EntityType", Flag: "entity-type", Type: "types.EntityType", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "SecurityProfiles", Flag: "security-profiles", Type: "[]types.SecurityProfileItem", Required: true},
}

var fields_disassociate_traffic_distribution_group_user = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "TrafficDistributionGroupId", Flag: "traffic-distribution-group-id", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_disassociate_user_proficiencies = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
	{Name: "UserProficiencies", Flag: "user-proficiencies", Type: "[]types.UserProficiencyDisassociate", Required: true},
}

var fields_disassociate_workspace = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "ResourceArns", Flag: "resource-arns", Type: "[]string", Required: true},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_dismiss_user_contact = []leanruntime.Field{
	{Name: "ContactId", Flag: "contact-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_evaluate_data_table_values = []leanruntime.Field{
	{Name: "DataTableId", Flag: "data-table-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TimeZone", Flag: "time-zone", Type: "*string", Required: false},
	{Name: "Values", Flag: "values", Type: "[]types.DataTableValueEvaluationSet", Required: true},
}

var fields_get_attached_file = []leanruntime.Field{
	{Name: "AssociatedResourceArn", Flag: "associated-resource-arn", Type: "*string", Required: true},
	{Name: "FileId", Flag: "file-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "UrlExpiryInSeconds", Flag: "url-expiry-in-seconds", Type: "*int32", Required: false},
}

var fields_get_contact_attributes = []leanruntime.Field{
	{Name: "InitialContactId", Flag: "initial-contact-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_get_contact_metrics = []leanruntime.Field{
	{Name: "ContactId", Flag: "contact-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Metrics", Flag: "metrics", Type: "[]types.ContactMetricInfo", Required: true},
}

var fields_get_current_metric_data = []leanruntime.Field{
	{Name: "CurrentMetrics", Flag: "current-metrics", Type: "[]types.CurrentMetric", Required: true},
	{Name: "Filters", Flag: "filters", Type: "*types.Filters", Required: true},
	{Name: "Groupings", Flag: "groupings", Type: "[]types.Grouping", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortCriteria", Flag: "sort-criteria", Type: "[]types.CurrentMetricSortCriteria", Required: false},
}

var fields_get_current_user_data = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "*types.UserDataFilters", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_effective_hours_of_operations = []leanruntime.Field{
	{Name: "FromDate", Flag: "from-date", Type: "*string", Required: true},
	{Name: "HoursOfOperationId", Flag: "hours-of-operation-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "ToDate", Flag: "to-date", Type: "*string", Required: true},
}

var fields_get_federation_token = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_get_flow_association = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.FlowAssociationResourceType", Required: true},
}

var fields_get_metric_data = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "Filters", Flag: "filters", Type: "*types.Filters", Required: true},
	{Name: "Groupings", Flag: "groupings", Type: "[]types.Grouping", Required: false},
	{Name: "HistoricalMetrics", Flag: "historical-metrics", Type: "[]types.HistoricalMetric", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
}

var fields_get_metric_data_v2 = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.FilterV2", Required: true},
	{Name: "Groupings", Flag: "groupings", Type: "[]string", Required: false},
	{Name: "Interval", Flag: "interval", Type: "*types.IntervalDetails", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Metrics", Flag: "metrics", Type: "[]types.MetricV2", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
}

var fields_get_prompt_file = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "PromptId", Flag: "prompt-id", Type: "*string", Required: true},
}

var fields_get_task_template = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "SnapshotVersion", Flag: "snapshot-version", Type: "*string", Required: false},
	{Name: "TaskTemplateId", Flag: "task-template-id", Type: "*string", Required: true},
}

var fields_get_test_case_execution_summary = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "TestCaseExecutionId", Flag: "test-case-execution-id", Type: "*string", Required: true},
	{Name: "TestCaseId", Flag: "test-case-id", Type: "*string", Required: true},
}

var fields_get_traffic_distribution = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_import_phone_number = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "PhoneNumberDescription", Flag: "phone-number-description", Type: "*string", Required: false},
	{Name: "SourcePhoneNumberArn", Flag: "source-phone-number-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_import_workspace_media = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MediaSource", Flag: "media-source", Type: "*string", Required: true},
	{Name: "MediaType", Flag: "media-type", Type: "types.MediaType", Required: true},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_list_agent_statuses = []leanruntime.Field{
	{Name: "AgentStatusTypes", Flag: "agent-status-types", Type: "[]types.AgentStatusType", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_analytics_data_associations = []leanruntime.Field{
	{Name: "DataSetId", Flag: "data-set-id", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_analytics_data_lake_data_sets = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_approved_origins = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_associated_contacts = []leanruntime.Field{
	{Name: "ContactId", Flag: "contact-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_authentication_profiles = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_bots = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "LexVersion", Flag: "lex-version", Type: "types.LexVersion", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_child_hours_of_operations = []leanruntime.Field{
	{Name: "HoursOfOperationId", Flag: "hours-of-operation-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_contact_evaluations = []leanruntime.Field{
	{Name: "ContactId", Flag: "contact-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_contact_flow_module_aliases = []leanruntime.Field{
	{Name: "ContactFlowModuleId", Flag: "contact-flow-module-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_contact_flow_module_versions = []leanruntime.Field{
	{Name: "ContactFlowModuleId", Flag: "contact-flow-module-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_contact_flow_modules = []leanruntime.Field{
	{Name: "ContactFlowModuleState", Flag: "contact-flow-module-state", Type: "types.ContactFlowModuleState", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_contact_flow_versions = []leanruntime.Field{
	{Name: "ContactFlowId", Flag: "contact-flow-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_contact_flows = []leanruntime.Field{
	{Name: "ContactFlowTypes", Flag: "contact-flow-types", Type: "[]types.ContactFlowType", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_contact_references = []leanruntime.Field{
	{Name: "ContactId", Flag: "contact-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ReferenceTypes", Flag: "reference-types", Type: "[]types.ReferenceType", Required: true},
}

var fields_list_data_table_attributes = []leanruntime.Field{
	{Name: "AttributeIds", Flag: "attribute-ids", Type: "[]string", Required: false},
	{Name: "DataTableId", Flag: "data-table-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_data_table_primary_values = []leanruntime.Field{
	{Name: "DataTableId", Flag: "data-table-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PrimaryAttributeValues", Flag: "primary-attribute-values", Type: "[]types.PrimaryAttributeValueFilter", Required: false},
	{Name: "RecordIds", Flag: "record-ids", Type: "[]string", Required: false},
}

var fields_list_data_table_values = []leanruntime.Field{
	{Name: "DataTableId", Flag: "data-table-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PrimaryAttributeValues", Flag: "primary-attribute-values", Type: "[]types.PrimaryAttributeValueFilter", Required: false},
	{Name: "RecordIds", Flag: "record-ids", Type: "[]string", Required: false},
}

var fields_list_data_tables = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_default_vocabularies = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "LanguageCode", Flag: "language-code", Type: "types.VocabularyLanguageCode", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_entity_security_profiles = []leanruntime.Field{
	{Name: "EntityArn", Flag: "entity-arn", Type: "*string", Required: true},
	{Name: "EntityType", Flag: "entity-type", Type: "types.EntityType", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_evaluation_form_versions = []leanruntime.Field{
	{Name: "EvaluationFormId", Flag: "evaluation-form-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_evaluation_forms = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_flow_associations = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.ListFlowAssociationResourceType", Required: false},
}

var fields_list_hours_of_operation_overrides = []leanruntime.Field{
	{Name: "HoursOfOperationId", Flag: "hours-of-operation-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_hours_of_operations = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_instance_attributes = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_instance_storage_configs = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.InstanceStorageResourceType", Required: true},
}

var fields_list_instances = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_integration_associations = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "IntegrationArn", Flag: "integration-arn", Type: "*string", Required: false},
	{Name: "IntegrationType", Flag: "integration-type", Type: "types.IntegrationType", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_lambda_functions = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_lex_bots = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_notifications = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_phone_numbers = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PhoneNumberCountryCodes", Flag: "phone-number-country-codes", Type: "[]types.PhoneNumberCountryCode", Required: false},
	{Name: "PhoneNumberTypes", Flag: "phone-number-types", Type: "[]types.PhoneNumberType", Required: false},
}

var fields_list_phone_numbers_v2 = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PhoneNumberCountryCodes", Flag: "phone-number-country-codes", Type: "[]types.PhoneNumberCountryCode", Required: false},
	{Name: "PhoneNumberPrefix", Flag: "phone-number-prefix", Type: "*string", Required: false},
	{Name: "PhoneNumberTypes", Flag: "phone-number-types", Type: "[]types.PhoneNumberType", Required: false},
	{Name: "TargetArn", Flag: "target-arn", Type: "*string", Required: false},
}

var fields_list_predefined_attributes = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_prompts = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_queue_quick_connects = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
}

var fields_list_queues = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "QueueTypes", Flag: "queue-types", Type: "[]types.QueueType", Required: false},
}

var fields_list_quick_connects = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "QuickConnectTypes", Flag: "quick-connect-types", Type: "[]types.QuickConnectType", Required: false},
}

var fields_list_realtime_contact_analysis_segments_v2 = []leanruntime.Field{
	{Name: "ContactId", Flag: "contact-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OutputType", Flag: "output-type", Type: "types.RealTimeContactAnalysisOutputType", Required: true},
	{Name: "SegmentTypes", Flag: "segment-types", Type: "[]types.RealTimeContactAnalysisSegmentType", Required: true},
}

var fields_list_routing_profile_manual_assignment_queues = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RoutingProfileId", Flag: "routing-profile-id", Type: "*string", Required: true},
}

var fields_list_routing_profile_queues = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RoutingProfileId", Flag: "routing-profile-id", Type: "*string", Required: true},
}

var fields_list_routing_profiles = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_rules = []leanruntime.Field{
	{Name: "EventSourceName", Flag: "event-source-name", Type: "types.EventSourceName", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PublishStatus", Flag: "publish-status", Type: "types.RulePublishStatus", Required: false},
}

var fields_list_security_keys = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_security_profile_applications = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SecurityProfileId", Flag: "security-profile-id", Type: "*string", Required: true},
}

var fields_list_security_profile_flow_modules = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SecurityProfileId", Flag: "security-profile-id", Type: "*string", Required: true},
}

var fields_list_security_profile_permissions = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SecurityProfileId", Flag: "security-profile-id", Type: "*string", Required: true},
}

var fields_list_security_profiles = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_task_templates = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.TaskTemplateStatus", Required: false},
}

var fields_list_test_case_execution_records = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.TestCaseExecutionStatus", Required: false},
	{Name: "TestCaseExecutionId", Flag: "test-case-execution-id", Type: "*string", Required: true},
	{Name: "TestCaseId", Flag: "test-case-id", Type: "*string", Required: true},
}

var fields_list_test_case_executions = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
	{Name: "Status", Flag: "status", Type: "types.TestCaseExecutionStatus", Required: false},
	{Name: "TestCaseId", Flag: "test-case-id", Type: "*string", Required: false},
	{Name: "TestCaseName", Flag: "test-case-name", Type: "*string", Required: false},
}

var fields_list_test_cases = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_traffic_distribution_group_users = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TrafficDistributionGroupId", Flag: "traffic-distribution-group-id", Type: "*string", Required: true},
}

var fields_list_traffic_distribution_groups = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_use_cases = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "IntegrationAssociationId", Flag: "integration-association-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_user_hierarchy_groups = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_user_notifications = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_list_user_proficiencies = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_list_users = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_view_versions = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ViewId", Flag: "view-id", Type: "*string", Required: true},
}

var fields_list_views = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.ViewType", Required: false},
}

var fields_list_workspace_media = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_list_workspace_pages = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_list_workspaces = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_monitor_contact = []leanruntime.Field{
	{Name: "AllowedMonitorCapabilities", Flag: "allowed-monitor-capabilities", Type: "[]types.MonitorCapability", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ContactId", Flag: "contact-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_pause_contact = []leanruntime.Field{
	{Name: "ContactFlowId", Flag: "contact-flow-id", Type: "*string", Required: false},
	{Name: "ContactId", Flag: "contact-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_put_user_status = []leanruntime.Field{
	{Name: "AgentStatusId", Flag: "agent-status-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_release_phone_number = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "PhoneNumberId", Flag: "phone-number-id", Type: "*string", Required: true},
}

var fields_replicate_instance = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "ReplicaAlias", Flag: "replica-alias", Type: "*string", Required: true},
	{Name: "ReplicaRegion", Flag: "replica-region", Type: "*string", Required: true},
}

var fields_resume_contact = []leanruntime.Field{
	{Name: "ContactFlowId", Flag: "contact-flow-id", Type: "*string", Required: false},
	{Name: "ContactId", Flag: "contact-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_resume_contact_recording = []leanruntime.Field{
	{Name: "ContactId", Flag: "contact-id", Type: "*string", Required: true},
	{Name: "ContactRecordingType", Flag: "contact-recording-type", Type: "types.ContactRecordingType", Required: false},
	{Name: "InitialContactId", Flag: "initial-contact-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_search_agent_statuses = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SearchCriteria", Flag: "search-criteria", Type: "*types.AgentStatusSearchCriteria", Required: false},
	{Name: "SearchFilter", Flag: "search-filter", Type: "*types.AgentStatusSearchFilter", Required: false},
}

var fields_search_available_phone_numbers = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PhoneNumberCountryCode", Flag: "phone-number-country-code", Type: "types.PhoneNumberCountryCode", Required: true},
	{Name: "PhoneNumberPrefix", Flag: "phone-number-prefix", Type: "*string", Required: false},
	{Name: "PhoneNumberType", Flag: "phone-number-type", Type: "types.PhoneNumberType", Required: true},
	{Name: "TargetArn", Flag: "target-arn", Type: "*string", Required: false},
}

var fields_search_contact_evaluations = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SearchCriteria", Flag: "search-criteria", Type: "*types.EvaluationSearchCriteria", Required: false},
	{Name: "SearchFilter", Flag: "search-filter", Type: "*types.EvaluationSearchFilter", Required: false},
}

var fields_search_contact_flow_modules = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SearchCriteria", Flag: "search-criteria", Type: "*types.ContactFlowModuleSearchCriteria", Required: false},
	{Name: "SearchFilter", Flag: "search-filter", Type: "*types.ContactFlowModuleSearchFilter", Required: false},
}

var fields_search_contact_flows = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SearchCriteria", Flag: "search-criteria", Type: "*types.ContactFlowSearchCriteria", Required: false},
	{Name: "SearchFilter", Flag: "search-filter", Type: "*types.ContactFlowSearchFilter", Required: false},
}

var fields_search_contacts = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SearchCriteria", Flag: "search-criteria", Type: "*types.SearchCriteria", Required: false},
	{Name: "Sort", Flag: "sort", Type: "*types.Sort", Required: false},
	{Name: "TimeRange", Flag: "time-range", Type: "*types.SearchContactsTimeRange", Required: true},
}

var fields_search_data_tables = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SearchCriteria", Flag: "search-criteria", Type: "*types.DataTableSearchCriteria", Required: false},
	{Name: "SearchFilter", Flag: "search-filter", Type: "*types.DataTableSearchFilter", Required: false},
}

var fields_search_email_addresses = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SearchCriteria", Flag: "search-criteria", Type: "*types.EmailAddressSearchCriteria", Required: false},
	{Name: "SearchFilter", Flag: "search-filter", Type: "*types.EmailAddressSearchFilter", Required: false},
}

var fields_search_evaluation_forms = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SearchCriteria", Flag: "search-criteria", Type: "*types.EvaluationFormSearchCriteria", Required: false},
	{Name: "SearchFilter", Flag: "search-filter", Type: "*types.EvaluationFormSearchFilter", Required: false},
}

var fields_search_hours_of_operation_overrides = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SearchCriteria", Flag: "search-criteria", Type: "*types.HoursOfOperationOverrideSearchCriteria", Required: false},
	{Name: "SearchFilter", Flag: "search-filter", Type: "*types.HoursOfOperationSearchFilter", Required: false},
}

var fields_search_hours_of_operations = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SearchCriteria", Flag: "search-criteria", Type: "*types.HoursOfOperationSearchCriteria", Required: false},
	{Name: "SearchFilter", Flag: "search-filter", Type: "*types.HoursOfOperationSearchFilter", Required: false},
}

var fields_search_notifications = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SearchCriteria", Flag: "search-criteria", Type: "*types.NotificationSearchCriteria", Required: false},
	{Name: "SearchFilter", Flag: "search-filter", Type: "*types.NotificationSearchFilter", Required: false},
}

var fields_search_predefined_attributes = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SearchCriteria", Flag: "search-criteria", Type: "*types.PredefinedAttributeSearchCriteria", Required: false},
}

var fields_search_prompts = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SearchCriteria", Flag: "search-criteria", Type: "*types.PromptSearchCriteria", Required: false},
	{Name: "SearchFilter", Flag: "search-filter", Type: "*types.PromptSearchFilter", Required: false},
}

var fields_search_queues = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SearchCriteria", Flag: "search-criteria", Type: "*types.QueueSearchCriteria", Required: false},
	{Name: "SearchFilter", Flag: "search-filter", Type: "*types.QueueSearchFilter", Required: false},
}

var fields_search_quick_connects = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SearchCriteria", Flag: "search-criteria", Type: "*types.QuickConnectSearchCriteria", Required: false},
	{Name: "SearchFilter", Flag: "search-filter", Type: "*types.QuickConnectSearchFilter", Required: false},
}

var fields_search_resource_tags = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceTypes", Flag: "resource-types", Type: "[]string", Required: false},
	{Name: "SearchCriteria", Flag: "search-criteria", Type: "*types.ResourceTagsSearchCriteria", Required: false},
}

var fields_search_routing_profiles = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SearchCriteria", Flag: "search-criteria", Type: "*types.RoutingProfileSearchCriteria", Required: false},
	{Name: "SearchFilter", Flag: "search-filter", Type: "*types.RoutingProfileSearchFilter", Required: false},
}

var fields_search_security_profiles = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SearchCriteria", Flag: "search-criteria", Type: "*types.SecurityProfileSearchCriteria", Required: false},
	{Name: "SearchFilter", Flag: "search-filter", Type: "*types.SecurityProfilesSearchFilter", Required: false},
}

var fields_search_test_cases = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SearchCriteria", Flag: "search-criteria", Type: "*types.TestCaseSearchCriteria", Required: false},
	{Name: "SearchFilter", Flag: "search-filter", Type: "*types.TestCaseSearchFilter", Required: false},
}

var fields_search_user_hierarchy_groups = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SearchCriteria", Flag: "search-criteria", Type: "*types.UserHierarchyGroupSearchCriteria", Required: false},
	{Name: "SearchFilter", Flag: "search-filter", Type: "*types.UserHierarchyGroupSearchFilter", Required: false},
}

var fields_search_users = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SearchCriteria", Flag: "search-criteria", Type: "*types.UserSearchCriteria", Required: false},
	{Name: "SearchFilter", Flag: "search-filter", Type: "*types.UserSearchFilter", Required: false},
}

var fields_search_views = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SearchCriteria", Flag: "search-criteria", Type: "*types.ViewSearchCriteria", Required: false},
	{Name: "SearchFilter", Flag: "search-filter", Type: "*types.ViewSearchFilter", Required: false},
}

var fields_search_vocabularies = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "LanguageCode", Flag: "language-code", Type: "types.VocabularyLanguageCode", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameStartsWith", Flag: "name-starts-with", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "State", Flag: "state", Type: "types.VocabularyState", Required: false},
}

var fields_search_workspace_associations = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SearchCriteria", Flag: "search-criteria", Type: "*types.WorkspaceAssociationSearchCriteria", Required: false},
	{Name: "SearchFilter", Flag: "search-filter", Type: "*types.WorkspaceAssociationSearchFilter", Required: false},
}

var fields_search_workspaces = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SearchCriteria", Flag: "search-criteria", Type: "*types.WorkspaceSearchCriteria", Required: false},
	{Name: "SearchFilter", Flag: "search-filter", Type: "*types.WorkspaceSearchFilter", Required: false},
}

var fields_send_chat_integration_event = []leanruntime.Field{
	{Name: "DestinationId", Flag: "destination-id", Type: "*string", Required: true},
	{Name: "Event", Flag: "event", Type: "*types.ChatEvent", Required: true},
	{Name: "NewSessionDetails", Flag: "new-session-details", Type: "*types.NewSessionDetails", Required: false},
	{Name: "SourceId", Flag: "source-id", Type: "*string", Required: true},
	{Name: "Subtype", Flag: "subtype", Type: "*string", Required: false},
}

var fields_send_outbound_email = []leanruntime.Field{
	{Name: "AdditionalRecipients", Flag: "additional-recipients", Type: "*types.OutboundAdditionalRecipients", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DestinationEmailAddress", Flag: "destination-email-address", Type: "*types.EmailAddressInfo", Required: true},
	{Name: "EmailMessage", Flag: "email-message", Type: "*types.OutboundEmailContent", Required: true},
	{Name: "FromEmailAddress", Flag: "from-email-address", Type: "*types.EmailAddressInfo", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "SourceCampaign", Flag: "source-campaign", Type: "*types.SourceCampaign", Required: false},
	{Name: "TrafficType", Flag: "traffic-type", Type: "types.TrafficType", Required: true},
}

var fields_start_attached_file_upload = []leanruntime.Field{
	{Name: "AssociatedResourceArn", Flag: "associated-resource-arn", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CreatedBy", Flag: "created-by", Type: "types.CreatedByInfo", Required: false},
	{Name: "FileName", Flag: "file-name", Type: "*string", Required: true},
	{Name: "FileSizeInBytes", Flag: "file-size-in-bytes", Type: "*int64", Required: true},
	{Name: "FileUseCaseType", Flag: "file-use-case-type", Type: "types.FileUseCaseType", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "UrlExpiryInSeconds", Flag: "url-expiry-in-seconds", Type: "*int32", Required: false},
}

var fields_start_chat_contact = []leanruntime.Field{
	{Name: "Attributes", Flag: "attributes", Type: "map[string]string", Required: false},
	{Name: "ChatDurationInMinutes", Flag: "chat-duration-in-minutes", Type: "*int32", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ContactFlowId", Flag: "contact-flow-id", Type: "*string", Required: true},
	{Name: "CustomerId", Flag: "customer-id", Type: "*string", Required: false},
	{Name: "DisconnectOnCustomerExit", Flag: "disconnect-on-customer-exit", Type: "[]types.DisconnectOnCustomerExitParticipantType", Required: false},
	{Name: "InitialMessage", Flag: "initial-message", Type: "*types.ChatMessage", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "ParticipantConfiguration", Flag: "participant-configuration", Type: "*types.ParticipantConfiguration", Required: false},
	{Name: "ParticipantDetails", Flag: "participant-details", Type: "*types.ParticipantDetails", Required: true},
	{Name: "PersistentChat", Flag: "persistent-chat", Type: "*types.PersistentChat", Required: false},
	{Name: "RelatedContactId", Flag: "related-contact-id", Type: "*string", Required: false},
	{Name: "SegmentAttributes", Flag: "segment-attributes", Type: "map[string]types.SegmentAttributeValue", Required: false},
	{Name: "SupportedMessagingContentTypes", Flag: "supported-messaging-content-types", Type: "[]string", Required: false},
}

var fields_start_contact_evaluation = []leanruntime.Field{
	{Name: "AutoEvaluationConfiguration", Flag: "auto-evaluation-configuration", Type: "*types.AutoEvaluationConfiguration", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ContactId", Flag: "contact-id", Type: "*string", Required: true},
	{Name: "EvaluationFormId", Flag: "evaluation-form-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_start_contact_media_processing = []leanruntime.Field{
	{Name: "ContactId", Flag: "contact-id", Type: "*string", Required: false},
	{Name: "FailureMode", Flag: "failure-mode", Type: "types.ContactMediaProcessingFailureMode", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: false},
	{Name: "ProcessorArn", Flag: "processor-arn", Type: "*string", Required: false},
}

var fields_start_contact_recording = []leanruntime.Field{
	{Name: "ContactId", Flag: "contact-id", Type: "*string", Required: true},
	{Name: "InitialContactId", Flag: "initial-contact-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "VoiceRecordingConfiguration", Flag: "voice-recording-configuration", Type: "*types.VoiceRecordingConfiguration", Required: true},
}

var fields_start_contact_streaming = []leanruntime.Field{
	{Name: "ChatStreamingConfiguration", Flag: "chat-streaming-configuration", Type: "*types.ChatStreamingConfiguration", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "ContactId", Flag: "contact-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_start_email_contact = []leanruntime.Field{
	{Name: "AdditionalRecipients", Flag: "additional-recipients", Type: "*types.InboundAdditionalRecipients", Required: false},
	{Name: "Attachments", Flag: "attachments", Type: "[]types.EmailAttachment", Required: false},
	{Name: "Attributes", Flag: "attributes", Type: "map[string]string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ContactFlowId", Flag: "contact-flow-id", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DestinationEmailAddress", Flag: "destination-email-address", Type: "*string", Required: true},
	{Name: "EmailMessage", Flag: "email-message", Type: "*types.InboundEmailContent", Required: true},
	{Name: "FromEmailAddress", Flag: "from-email-address", Type: "*types.EmailAddressInfo", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "References", Flag: "references", Type: "map[string]types.Reference", Required: false},
	{Name: "RelatedContactId", Flag: "related-contact-id", Type: "*string", Required: false},
	{Name: "SegmentAttributes", Flag: "segment-attributes", Type: "map[string]types.SegmentAttributeValue", Required: false},
}

var fields_start_outbound_chat_contact = []leanruntime.Field{
	{Name: "Attributes", Flag: "attributes", Type: "map[string]string", Required: false},
	{Name: "ChatDurationInMinutes", Flag: "chat-duration-in-minutes", Type: "*int32", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ContactFlowId", Flag: "contact-flow-id", Type: "*string", Required: true},
	{Name: "DestinationEndpoint", Flag: "destination-endpoint", Type: "*types.Endpoint", Required: true},
	{Name: "InitialSystemMessage", Flag: "initial-system-message", Type: "*types.ChatMessage", Required: false},
	{Name: "InitialTemplatedSystemMessage", Flag: "initial-templated-system-message", Type: "*types.TemplatedMessageConfig", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "ParticipantDetails", Flag: "participant-details", Type: "*types.ParticipantDetails", Required: false},
	{Name: "RelatedContactId", Flag: "related-contact-id", Type: "*string", Required: false},
	{Name: "SegmentAttributes", Flag: "segment-attributes", Type: "map[string]types.SegmentAttributeValue", Required: true},
	{Name: "SourceEndpoint", Flag: "source-endpoint", Type: "*types.Endpoint", Required: true},
	{Name: "SupportedMessagingContentTypes", Flag: "supported-messaging-content-types", Type: "[]string", Required: false},
}

var fields_start_outbound_email_contact = []leanruntime.Field{
	{Name: "AdditionalRecipients", Flag: "additional-recipients", Type: "*types.OutboundAdditionalRecipients", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ContactId", Flag: "contact-id", Type: "*string", Required: true},
	{Name: "DestinationEmailAddress", Flag: "destination-email-address", Type: "*types.EmailAddressInfo", Required: true},
	{Name: "EmailMessage", Flag: "email-message", Type: "*types.OutboundEmailContent", Required: true},
	{Name: "FromEmailAddress", Flag: "from-email-address", Type: "*types.EmailAddressInfo", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_start_outbound_voice_contact = []leanruntime.Field{
	{Name: "AnswerMachineDetectionConfig", Flag: "answer-machine-detection-config", Type: "*types.AnswerMachineDetectionConfig", Required: false},
	{Name: "Attributes", Flag: "attributes", Type: "map[string]string", Required: false},
	{Name: "CampaignId", Flag: "campaign-id", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ContactFlowId", Flag: "contact-flow-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DestinationPhoneNumber", Flag: "destination-phone-number", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "OutboundStrategy", Flag: "outbound-strategy", Type: "*types.OutboundStrategy", Required: false},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: false},
	{Name: "References", Flag: "references", Type: "map[string]types.Reference", Required: false},
	{Name: "RelatedContactId", Flag: "related-contact-id", Type: "*string", Required: false},
	{Name: "RingTimeoutInSeconds", Flag: "ring-timeout-in-seconds", Type: "*int32", Required: false},
	{Name: "SourcePhoneNumber", Flag: "source-phone-number", Type: "*string", Required: false},
	{Name: "TrafficType", Flag: "traffic-type", Type: "types.TrafficType", Required: false},
}

var fields_start_screen_sharing = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ContactId", Flag: "contact-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_start_task_contact = []leanruntime.Field{
	{Name: "Attachments", Flag: "attachments", Type: "[]types.TaskAttachment", Required: false},
	{Name: "Attributes", Flag: "attributes", Type: "map[string]string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ContactFlowId", Flag: "contact-flow-id", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "PreviousContactId", Flag: "previous-contact-id", Type: "*string", Required: false},
	{Name: "QuickConnectId", Flag: "quick-connect-id", Type: "*string", Required: false},
	{Name: "References", Flag: "references", Type: "map[string]types.Reference", Required: false},
	{Name: "RelatedContactId", Flag: "related-contact-id", Type: "*string", Required: false},
	{Name: "ScheduledTime", Flag: "scheduled-time", Type: "*time.Time", Required: false},
	{Name: "SegmentAttributes", Flag: "segment-attributes", Type: "map[string]types.SegmentAttributeValue", Required: false},
	{Name: "TaskTemplateId", Flag: "task-template-id", Type: "*string", Required: false},
}

var fields_start_test_case_execution = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "TestCaseId", Flag: "test-case-id", Type: "*string", Required: true},
}

var fields_start_web_rtc_contact = []leanruntime.Field{
	{Name: "AllowedCapabilities", Flag: "allowed-capabilities", Type: "*types.AllowedCapabilities", Required: false},
	{Name: "Attributes", Flag: "attributes", Type: "map[string]string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ContactFlowId", Flag: "contact-flow-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "ParticipantDetails", Flag: "participant-details", Type: "*types.ParticipantDetails", Required: true},
	{Name: "References", Flag: "references", Type: "map[string]types.Reference", Required: false},
	{Name: "RelatedContactId", Flag: "related-contact-id", Type: "*string", Required: false},
}

var fields_stop_contact = []leanruntime.Field{
	{Name: "ContactId", Flag: "contact-id", Type: "*string", Required: true},
	{Name: "DisconnectReason", Flag: "disconnect-reason", Type: "*types.DisconnectReason", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_stop_contact_media_processing = []leanruntime.Field{
	{Name: "ContactId", Flag: "contact-id", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: false},
}

var fields_stop_contact_recording = []leanruntime.Field{
	{Name: "ContactId", Flag: "contact-id", Type: "*string", Required: true},
	{Name: "ContactRecordingType", Flag: "contact-recording-type", Type: "types.ContactRecordingType", Required: false},
	{Name: "InitialContactId", Flag: "initial-contact-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_stop_contact_streaming = []leanruntime.Field{
	{Name: "ContactId", Flag: "contact-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "StreamingId", Flag: "streaming-id", Type: "*string", Required: true},
}

var fields_stop_test_case_execution = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "TestCaseExecutionId", Flag: "test-case-execution-id", Type: "*string", Required: true},
	{Name: "TestCaseId", Flag: "test-case-id", Type: "*string", Required: true},
}

var fields_submit_contact_evaluation = []leanruntime.Field{
	{Name: "Answers", Flag: "answers", Type: "map[string]types.EvaluationAnswerInput", Required: false},
	{Name: "EvaluationId", Flag: "evaluation-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Notes", Flag: "notes", Type: "map[string]types.EvaluationNote", Required: false},
	{Name: "SubmittedBy", Flag: "submitted-by", Type: "types.EvaluatorUserUnion", Required: false},
}

var fields_suspend_contact_recording = []leanruntime.Field{
	{Name: "ContactId", Flag: "contact-id", Type: "*string", Required: true},
	{Name: "ContactRecordingType", Flag: "contact-recording-type", Type: "types.ContactRecordingType", Required: false},
	{Name: "InitialContactId", Flag: "initial-contact-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_tag_contact = []leanruntime.Field{
	{Name: "ContactId", Flag: "contact-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_transfer_contact = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ContactFlowId", Flag: "contact-flow-id", Type: "*string", Required: true},
	{Name: "ContactId", Flag: "contact-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: false},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: false},
}

var fields_untag_contact = []leanruntime.Field{
	{Name: "ContactId", Flag: "contact-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_agent_status = []leanruntime.Field{
	{Name: "AgentStatusId", Flag: "agent-status-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayOrder", Flag: "display-order", Type: "*int32", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "ResetOrderNumber", Flag: "reset-order-number", Type: "bool", Required: false},
	{Name: "State", Flag: "state", Type: "types.AgentStatusState", Required: false},
}

var fields_update_authentication_profile = []leanruntime.Field{
	{Name: "AllowedIps", Flag: "allowed-ips", Type: "[]string", Required: false},
	{Name: "AuthenticationProfileId", Flag: "authentication-profile-id", Type: "*string", Required: true},
	{Name: "BlockedIps", Flag: "blocked-ips", Type: "[]string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "PeriodicSessionDuration", Flag: "periodic-session-duration", Type: "*int32", Required: false},
	{Name: "SessionInactivityDuration", Flag: "session-inactivity-duration", Type: "*int32", Required: false},
	{Name: "SessionInactivityHandlingEnabled", Flag: "session-inactivity-handling-enabled", Type: "*bool", Required: false},
}

var fields_update_contact = []leanruntime.Field{
	{Name: "ContactId", Flag: "contact-id", Type: "*string", Required: true},
	{Name: "CustomerEndpoint", Flag: "customer-endpoint", Type: "*types.Endpoint", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "QueueInfo", Flag: "queue-info", Type: "*types.QueueInfoInput", Required: false},
	{Name: "References", Flag: "references", Type: "map[string]types.Reference", Required: false},
	{Name: "SegmentAttributes", Flag: "segment-attributes", Type: "map[string]types.SegmentAttributeValue", Required: false},
	{Name: "SystemEndpoint", Flag: "system-endpoint", Type: "*types.Endpoint", Required: false},
	{Name: "UserInfo", Flag: "user-info", Type: "*types.UserInfo", Required: false},
}

var fields_update_contact_attributes = []leanruntime.Field{
	{Name: "Attributes", Flag: "attributes", Type: "map[string]string", Required: true},
	{Name: "InitialContactId", Flag: "initial-contact-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_update_contact_evaluation = []leanruntime.Field{
	{Name: "Answers", Flag: "answers", Type: "map[string]types.EvaluationAnswerInput", Required: false},
	{Name: "EvaluationId", Flag: "evaluation-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Notes", Flag: "notes", Type: "map[string]types.EvaluationNote", Required: false},
	{Name: "UpdatedBy", Flag: "updated-by", Type: "types.EvaluatorUserUnion", Required: false},
}

var fields_update_contact_flow_content = []leanruntime.Field{
	{Name: "ContactFlowId", Flag: "contact-flow-id", Type: "*string", Required: true},
	{Name: "Content", Flag: "content", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_update_contact_flow_metadata = []leanruntime.Field{
	{Name: "ContactFlowId", Flag: "contact-flow-id", Type: "*string", Required: true},
	{Name: "ContactFlowState", Flag: "contact-flow-state", Type: "types.ContactFlowState", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_contact_flow_module_alias = []leanruntime.Field{
	{Name: "AliasId", Flag: "alias-id", Type: "*string", Required: true},
	{Name: "ContactFlowModuleId", Flag: "contact-flow-module-id", Type: "*string", Required: true},
	{Name: "ContactFlowModuleVersion", Flag: "contact-flow-module-version", Type: "*int64", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_contact_flow_module_content = []leanruntime.Field{
	{Name: "ContactFlowModuleId", Flag: "contact-flow-module-id", Type: "*string", Required: true},
	{Name: "Content", Flag: "content", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Settings", Flag: "settings", Type: "*string", Required: false},
}

var fields_update_contact_flow_module_metadata = []leanruntime.Field{
	{Name: "ContactFlowModuleId", Flag: "contact-flow-module-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "State", Flag: "state", Type: "types.ContactFlowModuleState", Required: false},
}

var fields_update_contact_flow_name = []leanruntime.Field{
	{Name: "ContactFlowId", Flag: "contact-flow-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_contact_routing_data = []leanruntime.Field{
	{Name: "ContactId", Flag: "contact-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "QueuePriority", Flag: "queue-priority", Type: "*int64", Required: false},
	{Name: "QueueTimeAdjustmentSeconds", Flag: "queue-time-adjustment-seconds", Type: "*int32", Required: false},
	{Name: "RoutingCriteria", Flag: "routing-criteria", Type: "*types.RoutingCriteriaInput", Required: false},
}

var fields_update_contact_schedule = []leanruntime.Field{
	{Name: "ContactId", Flag: "contact-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "ScheduledTime", Flag: "scheduled-time", Type: "*time.Time", Required: true},
}

var fields_update_data_table_attribute = []leanruntime.Field{
	{Name: "AttributeName", Flag: "attribute-name", Type: "*string", Required: true},
	{Name: "DataTableId", Flag: "data-table-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Primary", Flag: "primary", Type: "bool", Required: false},
	{Name: "Validation", Flag: "validation", Type: "*types.Validation", Required: false},
	{Name: "ValueType", Flag: "value-type", Type: "types.DataTableAttributeValueType", Required: true},
}

var fields_update_data_table_metadata = []leanruntime.Field{
	{Name: "DataTableId", Flag: "data-table-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "TimeZone", Flag: "time-zone", Type: "*string", Required: true},
	{Name: "ValueLockLevel", Flag: "value-lock-level", Type: "types.DataTableLockLevel", Required: true},
}

var fields_update_data_table_primary_values = []leanruntime.Field{
	{Name: "DataTableId", Flag: "data-table-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "LockVersion", Flag: "lock-version", Type: "*types.DataTableLockVersion", Required: true},
	{Name: "NewPrimaryValues", Flag: "new-primary-values", Type: "[]types.PrimaryValue", Required: true},
	{Name: "PrimaryValues", Flag: "primary-values", Type: "[]types.PrimaryValue", Required: true},
}

var fields_update_email_address_metadata = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "EmailAddressId", Flag: "email-address-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_update_evaluation_form = []leanruntime.Field{
	{Name: "AsDraft", Flag: "as-draft", Type: "bool", Required: false},
	{Name: "AutoEvaluationConfiguration", Flag: "auto-evaluation-configuration", Type: "*types.EvaluationFormAutoEvaluationConfiguration", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CreateNewVersion", Flag: "create-new-version", Type: "*bool", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EvaluationFormId", Flag: "evaluation-form-id", Type: "*string", Required: true},
	{Name: "EvaluationFormVersion", Flag: "evaluation-form-version", Type: "int32", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Items", Flag: "items", Type: "[]types.EvaluationFormItem", Required: true},
	{Name: "LanguageConfiguration", Flag: "language-configuration", Type: "*types.EvaluationFormLanguageConfiguration", Required: false},
	{Name: "ReviewConfiguration", Flag: "review-configuration", Type: "*types.EvaluationReviewConfiguration", Required: false},
	{Name: "ScoringStrategy", Flag: "scoring-strategy", Type: "*types.EvaluationFormScoringStrategy", Required: false},
	{Name: "TargetConfiguration", Flag: "target-configuration", Type: "*types.EvaluationFormTargetConfiguration", Required: false},
	{Name: "Title", Flag: "title", Type: "*string", Required: true},
}

var fields_update_hours_of_operation = []leanruntime.Field{
	{Name: "Config", Flag: "config", Type: "[]types.HoursOfOperationConfig", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "HoursOfOperationId", Flag: "hours-of-operation-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "TimeZone", Flag: "time-zone", Type: "*string", Required: false},
}

var fields_update_hours_of_operation_override = []leanruntime.Field{
	{Name: "Config", Flag: "config", Type: "[]types.HoursOfOperationOverrideConfig", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EffectiveFrom", Flag: "effective-from", Type: "*string", Required: false},
	{Name: "EffectiveTill", Flag: "effective-till", Type: "*string", Required: false},
	{Name: "HoursOfOperationId", Flag: "hours-of-operation-id", Type: "*string", Required: true},
	{Name: "HoursOfOperationOverrideId", Flag: "hours-of-operation-override-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "OverrideType", Flag: "override-type", Type: "types.OverrideType", Required: false},
	{Name: "RecurrenceConfig", Flag: "recurrence-config", Type: "*types.RecurrenceConfig", Required: false},
}

var fields_update_instance_attribute = []leanruntime.Field{
	{Name: "AttributeType", Flag: "attribute-type", Type: "types.InstanceAttributeType", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Value", Flag: "value", Type: "*string", Required: true},
}

var fields_update_instance_storage_config = []leanruntime.Field{
	{Name: "AssociationId", Flag: "association-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.InstanceStorageResourceType", Required: true},
	{Name: "StorageConfig", Flag: "storage-config", Type: "*types.InstanceStorageConfig", Required: true},
}

var fields_update_notification_content = []leanruntime.Field{
	{Name: "Content", Flag: "content", Type: "map[string]string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "NotificationId", Flag: "notification-id", Type: "*string", Required: true},
}

var fields_update_participant_authentication = []leanruntime.Field{
	{Name: "Code", Flag: "code", Type: "*string", Required: false},
	{Name: "Error", Flag: "error", Type: "*string", Required: false},
	{Name: "ErrorDescription", Flag: "error-description", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "State", Flag: "state", Type: "*string", Required: true},
}

var fields_update_participant_role_config = []leanruntime.Field{
	{Name: "ChannelConfiguration", Flag: "channel-configuration", Type: "types.UpdateParticipantRoleConfigChannelInfo", Required: true},
	{Name: "ContactId", Flag: "contact-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_update_phone_number = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: false},
	{Name: "PhoneNumberId", Flag: "phone-number-id", Type: "*string", Required: true},
	{Name: "TargetArn", Flag: "target-arn", Type: "*string", Required: false},
}

var fields_update_phone_number_metadata = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "PhoneNumberDescription", Flag: "phone-number-description", Type: "*string", Required: false},
	{Name: "PhoneNumberId", Flag: "phone-number-id", Type: "*string", Required: true},
}

var fields_update_predefined_attribute = []leanruntime.Field{
	{Name: "AttributeConfiguration", Flag: "attribute-configuration", Type: "*types.InputPredefinedAttributeConfiguration", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Purposes", Flag: "purposes", Type: "[]string", Required: false},
	{Name: "Values", Flag: "values", Type: "types.PredefinedAttributeValues", Required: false},
}

var fields_update_prompt = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "PromptId", Flag: "prompt-id", Type: "*string", Required: true},
	{Name: "S3Uri", Flag: "s3-uri", Type: "*string", Required: false},
}

var fields_update_queue_hours_of_operation = []leanruntime.Field{
	{Name: "HoursOfOperationId", Flag: "hours-of-operation-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
}

var fields_update_queue_max_contacts = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxContacts", Flag: "max-contacts", Type: "*int32", Required: false},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
}

var fields_update_queue_name = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
}

var fields_update_queue_outbound_caller_config = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "OutboundCallerConfig", Flag: "outbound-caller-config", Type: "*types.OutboundCallerConfig", Required: true},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
}

var fields_update_queue_outbound_email_config = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "OutboundEmailConfig", Flag: "outbound-email-config", Type: "*types.OutboundEmailConfig", Required: true},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
}

var fields_update_queue_status = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.QueueStatus", Required: true},
}

var fields_update_quick_connect_config = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "QuickConnectConfig", Flag: "quick-connect-config", Type: "*types.QuickConnectConfig", Required: true},
	{Name: "QuickConnectId", Flag: "quick-connect-id", Type: "*string", Required: true},
}

var fields_update_quick_connect_name = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "QuickConnectId", Flag: "quick-connect-id", Type: "*string", Required: true},
}

var fields_update_routing_profile_agent_availability_timer = []leanruntime.Field{
	{Name: "AgentAvailabilityTimer", Flag: "agent-availability-timer", Type: "types.AgentAvailabilityTimer", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "RoutingProfileId", Flag: "routing-profile-id", Type: "*string", Required: true},
}

var fields_update_routing_profile_concurrency = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MediaConcurrencies", Flag: "media-concurrencies", Type: "[]types.MediaConcurrency", Required: true},
	{Name: "RoutingProfileId", Flag: "routing-profile-id", Type: "*string", Required: true},
}

var fields_update_routing_profile_default_outbound_queue = []leanruntime.Field{
	{Name: "DefaultOutboundQueueId", Flag: "default-outbound-queue-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "RoutingProfileId", Flag: "routing-profile-id", Type: "*string", Required: true},
}

var fields_update_routing_profile_name = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "RoutingProfileId", Flag: "routing-profile-id", Type: "*string", Required: true},
}

var fields_update_routing_profile_queues = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "QueueConfigs", Flag: "queue-configs", Type: "[]types.RoutingProfileQueueConfig", Required: true},
	{Name: "RoutingProfileId", Flag: "routing-profile-id", Type: "*string", Required: true},
}

var fields_update_rule = []leanruntime.Field{
	{Name: "Actions", Flag: "actions", Type: "[]types.RuleAction", Required: true},
	{Name: "Function", Flag: "function", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "PublishStatus", Flag: "publish-status", Type: "types.RulePublishStatus", Required: true},
	{Name: "RuleId", Flag: "rule-id", Type: "*string", Required: true},
}

var fields_update_security_profile = []leanruntime.Field{
	{Name: "AllowedAccessControlHierarchyGroupId", Flag: "allowed-access-control-hierarchy-group-id", Type: "*string", Required: false},
	{Name: "AllowedAccessControlTags", Flag: "allowed-access-control-tags", Type: "map[string]string", Required: false},
	{Name: "AllowedFlowModules", Flag: "allowed-flow-modules", Type: "[]types.FlowModule", Required: false},
	{Name: "Applications", Flag: "applications", Type: "[]types.Application", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "GranularAccessControlConfiguration", Flag: "granular-access-control-configuration", Type: "*types.GranularAccessControlConfiguration", Required: false},
	{Name: "HierarchyRestrictedResources", Flag: "hierarchy-restricted-resources", Type: "[]string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Permissions", Flag: "permissions", Type: "[]string", Required: false},
	{Name: "SecurityProfileId", Flag: "security-profile-id", Type: "*string", Required: true},
	{Name: "TagRestrictedResources", Flag: "tag-restricted-resources", Type: "[]string", Required: false},
}

var fields_update_task_template = []leanruntime.Field{
	{Name: "Constraints", Flag: "constraints", Type: "*types.TaskTemplateConstraints", Required: false},
	{Name: "ContactFlowId", Flag: "contact-flow-id", Type: "*string", Required: false},
	{Name: "Defaults", Flag: "defaults", Type: "*types.TaskTemplateDefaults", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Fields", Flag: "fields", Type: "[]types.TaskTemplateField", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "SelfAssignFlowId", Flag: "self-assign-flow-id", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.TaskTemplateStatus", Required: false},
	{Name: "TaskTemplateId", Flag: "task-template-id", Type: "*string", Required: true},
}

var fields_update_test_case = []leanruntime.Field{
	{Name: "Content", Flag: "content", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EntryPoint", Flag: "entry-point", Type: "*types.TestCaseEntryPoint", Required: false},
	{Name: "InitializationData", Flag: "initialization-data", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "LastModifiedRegion", Flag: "last-modified-region", Type: "*string", Required: false},
	{Name: "LastModifiedTime", Flag: "last-modified-time", Type: "*time.Time", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.TestCaseStatus", Required: false},
	{Name: "TestCaseId", Flag: "test-case-id", Type: "*string", Required: true},
}

var fields_update_traffic_distribution = []leanruntime.Field{
	{Name: "AgentConfig", Flag: "agent-config", Type: "*types.AgentConfig", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "SignInConfig", Flag: "sign-in-config", Type: "*types.SignInConfig", Required: false},
	{Name: "TelephonyConfig", Flag: "telephony-config", Type: "*types.TelephonyConfig", Required: false},
}

var fields_update_user_config = []leanruntime.Field{
	{Name: "AfterContactWorkConfigs", Flag: "after-contact-work-configs", Type: "[]types.AfterContactWorkConfigPerChannel", Required: false},
	{Name: "AutoAcceptConfigs", Flag: "auto-accept-configs", Type: "[]types.AutoAcceptConfig", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "PersistentConnectionConfigs", Flag: "persistent-connection-configs", Type: "[]types.PersistentConnectionConfig", Required: false},
	{Name: "PhoneNumberConfigs", Flag: "phone-number-configs", Type: "[]types.PhoneNumberConfig", Required: false},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
	{Name: "VoiceEnhancementConfigs", Flag: "voice-enhancement-configs", Type: "[]types.VoiceEnhancementConfig", Required: false},
}

var fields_update_user_hierarchy = []leanruntime.Field{
	{Name: "HierarchyGroupId", Flag: "hierarchy-group-id", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_update_user_hierarchy_group_name = []leanruntime.Field{
	{Name: "HierarchyGroupId", Flag: "hierarchy-group-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_update_user_hierarchy_structure = []leanruntime.Field{
	{Name: "HierarchyStructure", Flag: "hierarchy-structure", Type: "*types.HierarchyStructureUpdate", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_update_user_identity_info = []leanruntime.Field{
	{Name: "IdentityInfo", Flag: "identity-info", Type: "*types.UserIdentityInfo", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_update_user_notification_status = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "LastModifiedRegion", Flag: "last-modified-region", Type: "*string", Required: false},
	{Name: "LastModifiedTime", Flag: "last-modified-time", Type: "*time.Time", Required: false},
	{Name: "NotificationId", Flag: "notification-id", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.NotificationStatus", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_update_user_phone_config = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "PhoneConfig", Flag: "phone-config", Type: "*types.UserPhoneConfig", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_update_user_proficiencies = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
	{Name: "UserProficiencies", Flag: "user-proficiencies", Type: "[]types.UserProficiency", Required: true},
}

var fields_update_user_routing_profile = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "RoutingProfileId", Flag: "routing-profile-id", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_update_user_security_profiles = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "SecurityProfileIds", Flag: "security-profile-ids", Type: "[]string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_update_view_content = []leanruntime.Field{
	{Name: "Content", Flag: "content", Type: "*types.ViewInputContent", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.ViewStatus", Required: true},
	{Name: "ViewId", Flag: "view-id", Type: "*string", Required: true},
}

var fields_update_view_metadata = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "ViewId", Flag: "view-id", Type: "*string", Required: true},
}

var fields_update_workspace_metadata = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Title", Flag: "title", Type: "*string", Required: false},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_update_workspace_page = []leanruntime.Field{
	{Name: "InputData", Flag: "input-data", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "NewPage", Flag: "new-page", Type: "*string", Required: false},
	{Name: "Page", Flag: "page", Type: "*string", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: false},
	{Name: "Slug", Flag: "slug", Type: "*string", Required: false},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_update_workspace_theme = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Theme", Flag: "theme", Type: "*types.WorkspaceTheme", Required: false},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

var fields_update_workspace_visibility = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Visibility", Flag: "visibility", Type: "types.Visibility", Required: true},
	{Name: "WorkspaceId", Flag: "workspace-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"activate-evaluation-form": {
			Name:   "activate-evaluation-form",
			Fields: fields_activate_evaluation_form,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ActivateEvaluationFormInput{}
				if _, err := leanruntime.ApplyInput(input, fields_activate_evaluation_form, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ActivateEvaluationForm(ctx, input)
			},
		},
		"associate-analytics-data-set": {
			Name:   "associate-analytics-data-set",
			Fields: fields_associate_analytics_data_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateAnalyticsDataSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_analytics_data_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateAnalyticsDataSet(ctx, input)
			},
		},
		"associate-approved-origin": {
			Name:   "associate-approved-origin",
			Fields: fields_associate_approved_origin,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateApprovedOriginInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_approved_origin, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateApprovedOrigin(ctx, input)
			},
		},
		"associate-bot": {
			Name:   "associate-bot",
			Fields: fields_associate_bot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateBotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_bot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateBot(ctx, input)
			},
		},
		"associate-contact-with-user": {
			Name:   "associate-contact-with-user",
			Fields: fields_associate_contact_with_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateContactWithUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_contact_with_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateContactWithUser(ctx, input)
			},
		},
		"associate-default-vocabulary": {
			Name:   "associate-default-vocabulary",
			Fields: fields_associate_default_vocabulary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateDefaultVocabularyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_default_vocabulary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateDefaultVocabulary(ctx, input)
			},
		},
		"associate-email-address-alias": {
			Name:   "associate-email-address-alias",
			Fields: fields_associate_email_address_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateEmailAddressAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_email_address_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateEmailAddressAlias(ctx, input)
			},
		},
		"associate-flow": {
			Name:   "associate-flow",
			Fields: fields_associate_flow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateFlowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_flow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateFlow(ctx, input)
			},
		},
		"associate-hours-of-operations": {
			Name:   "associate-hours-of-operations",
			Fields: fields_associate_hours_of_operations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateHoursOfOperationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_hours_of_operations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateHoursOfOperations(ctx, input)
			},
		},
		"associate-instance-storage-config": {
			Name:   "associate-instance-storage-config",
			Fields: fields_associate_instance_storage_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateInstanceStorageConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_instance_storage_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateInstanceStorageConfig(ctx, input)
			},
		},
		"associate-lambda-function": {
			Name:   "associate-lambda-function",
			Fields: fields_associate_lambda_function,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateLambdaFunctionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_lambda_function, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateLambdaFunction(ctx, input)
			},
		},
		"associate-lex-bot": {
			Name:   "associate-lex-bot",
			Fields: fields_associate_lex_bot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateLexBotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_lex_bot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateLexBot(ctx, input)
			},
		},
		"associate-phone-number-contact-flow": {
			Name:   "associate-phone-number-contact-flow",
			Fields: fields_associate_phone_number_contact_flow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociatePhoneNumberContactFlowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_phone_number_contact_flow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociatePhoneNumberContactFlow(ctx, input)
			},
		},
		"associate-queue-quick-connects": {
			Name:   "associate-queue-quick-connects",
			Fields: fields_associate_queue_quick_connects,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateQueueQuickConnectsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_queue_quick_connects, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateQueueQuickConnects(ctx, input)
			},
		},
		"associate-routing-profile-queues": {
			Name:   "associate-routing-profile-queues",
			Fields: fields_associate_routing_profile_queues,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateRoutingProfileQueuesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_routing_profile_queues, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateRoutingProfileQueues(ctx, input)
			},
		},
		"associate-security-key": {
			Name:   "associate-security-key",
			Fields: fields_associate_security_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateSecurityKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_security_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateSecurityKey(ctx, input)
			},
		},
		"associate-security-profiles": {
			Name:   "associate-security-profiles",
			Fields: fields_associate_security_profiles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateSecurityProfilesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_security_profiles, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateSecurityProfiles(ctx, input)
			},
		},
		"associate-traffic-distribution-group-user": {
			Name:   "associate-traffic-distribution-group-user",
			Fields: fields_associate_traffic_distribution_group_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateTrafficDistributionGroupUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_traffic_distribution_group_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateTrafficDistributionGroupUser(ctx, input)
			},
		},
		"associate-user-proficiencies": {
			Name:   "associate-user-proficiencies",
			Fields: fields_associate_user_proficiencies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateUserProficienciesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_user_proficiencies, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateUserProficiencies(ctx, input)
			},
		},
		"associate-workspace": {
			Name:   "associate-workspace",
			Fields: fields_associate_workspace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateWorkspaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_workspace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateWorkspace(ctx, input)
			},
		},
		"batch-associate-analytics-data-set": {
			Name:   "batch-associate-analytics-data-set",
			Fields: fields_batch_associate_analytics_data_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchAssociateAnalyticsDataSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_associate_analytics_data_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchAssociateAnalyticsDataSet(ctx, input)
			},
		},
		"batch-create-data-table-value": {
			Name:   "batch-create-data-table-value",
			Fields: fields_batch_create_data_table_value,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchCreateDataTableValueInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_create_data_table_value, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchCreateDataTableValue(ctx, input)
			},
		},
		"batch-delete-data-table-value": {
			Name:   "batch-delete-data-table-value",
			Fields: fields_batch_delete_data_table_value,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDeleteDataTableValueInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_delete_data_table_value, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDeleteDataTableValue(ctx, input)
			},
		},
		"batch-describe-data-table-value": {
			Name:   "batch-describe-data-table-value",
			Fields: fields_batch_describe_data_table_value,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDescribeDataTableValueInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_describe_data_table_value, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDescribeDataTableValue(ctx, input)
			},
		},
		"batch-disassociate-analytics-data-set": {
			Name:   "batch-disassociate-analytics-data-set",
			Fields: fields_batch_disassociate_analytics_data_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDisassociateAnalyticsDataSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_disassociate_analytics_data_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDisassociateAnalyticsDataSet(ctx, input)
			},
		},
		"batch-get-attached-file-metadata": {
			Name:   "batch-get-attached-file-metadata",
			Fields: fields_batch_get_attached_file_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetAttachedFileMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_attached_file_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetAttachedFileMetadata(ctx, input)
			},
		},
		"batch-get-flow-association": {
			Name:   "batch-get-flow-association",
			Fields: fields_batch_get_flow_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetFlowAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_flow_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetFlowAssociation(ctx, input)
			},
		},
		"batch-put-contact": {
			Name:   "batch-put-contact",
			Fields: fields_batch_put_contact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchPutContactInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_put_contact, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchPutContact(ctx, input)
			},
		},
		"batch-update-data-table-value": {
			Name:   "batch-update-data-table-value",
			Fields: fields_batch_update_data_table_value,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchUpdateDataTableValueInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_update_data_table_value, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchUpdateDataTableValue(ctx, input)
			},
		},
		"claim-phone-number": {
			Name:   "claim-phone-number",
			Fields: fields_claim_phone_number,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ClaimPhoneNumberInput{}
				if _, err := leanruntime.ApplyInput(input, fields_claim_phone_number, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ClaimPhoneNumber(ctx, input)
			},
		},
		"complete-attached-file-upload": {
			Name:   "complete-attached-file-upload",
			Fields: fields_complete_attached_file_upload,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CompleteAttachedFileUploadInput{}
				if _, err := leanruntime.ApplyInput(input, fields_complete_attached_file_upload, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CompleteAttachedFileUpload(ctx, input)
			},
		},
		"create-agent-status": {
			Name:   "create-agent-status",
			Fields: fields_create_agent_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAgentStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_agent_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAgentStatus(ctx, input)
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
		"create-contact-flow": {
			Name:   "create-contact-flow",
			Fields: fields_create_contact_flow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateContactFlowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_contact_flow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateContactFlow(ctx, input)
			},
		},
		"create-contact-flow-module": {
			Name:   "create-contact-flow-module",
			Fields: fields_create_contact_flow_module,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateContactFlowModuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_contact_flow_module, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateContactFlowModule(ctx, input)
			},
		},
		"create-contact-flow-module-alias": {
			Name:   "create-contact-flow-module-alias",
			Fields: fields_create_contact_flow_module_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateContactFlowModuleAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_contact_flow_module_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateContactFlowModuleAlias(ctx, input)
			},
		},
		"create-contact-flow-module-version": {
			Name:   "create-contact-flow-module-version",
			Fields: fields_create_contact_flow_module_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateContactFlowModuleVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_contact_flow_module_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateContactFlowModuleVersion(ctx, input)
			},
		},
		"create-contact-flow-version": {
			Name:   "create-contact-flow-version",
			Fields: fields_create_contact_flow_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateContactFlowVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_contact_flow_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateContactFlowVersion(ctx, input)
			},
		},
		"create-data-table": {
			Name:   "create-data-table",
			Fields: fields_create_data_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDataTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_data_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDataTable(ctx, input)
			},
		},
		"create-data-table-attribute": {
			Name:   "create-data-table-attribute",
			Fields: fields_create_data_table_attribute,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDataTableAttributeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_data_table_attribute, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDataTableAttribute(ctx, input)
			},
		},
		"create-email-address": {
			Name:   "create-email-address",
			Fields: fields_create_email_address,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEmailAddressInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_email_address, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEmailAddress(ctx, input)
			},
		},
		"create-evaluation-form": {
			Name:   "create-evaluation-form",
			Fields: fields_create_evaluation_form,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEvaluationFormInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_evaluation_form, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEvaluationForm(ctx, input)
			},
		},
		"create-hours-of-operation": {
			Name:   "create-hours-of-operation",
			Fields: fields_create_hours_of_operation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateHoursOfOperationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_hours_of_operation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateHoursOfOperation(ctx, input)
			},
		},
		"create-hours-of-operation-override": {
			Name:   "create-hours-of-operation-override",
			Fields: fields_create_hours_of_operation_override,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateHoursOfOperationOverrideInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_hours_of_operation_override, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateHoursOfOperationOverride(ctx, input)
			},
		},
		"create-instance": {
			Name:   "create-instance",
			Fields: fields_create_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateInstance(ctx, input)
			},
		},
		"create-integration-association": {
			Name:   "create-integration-association",
			Fields: fields_create_integration_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateIntegrationAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_integration_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateIntegrationAssociation(ctx, input)
			},
		},
		"create-notification": {
			Name:   "create-notification",
			Fields: fields_create_notification,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateNotificationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_notification, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateNotification(ctx, input)
			},
		},
		"create-participant": {
			Name:   "create-participant",
			Fields: fields_create_participant,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateParticipantInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_participant, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateParticipant(ctx, input)
			},
		},
		"create-persistent-contact-association": {
			Name:   "create-persistent-contact-association",
			Fields: fields_create_persistent_contact_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePersistentContactAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_persistent_contact_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePersistentContactAssociation(ctx, input)
			},
		},
		"create-predefined-attribute": {
			Name:   "create-predefined-attribute",
			Fields: fields_create_predefined_attribute,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePredefinedAttributeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_predefined_attribute, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePredefinedAttribute(ctx, input)
			},
		},
		"create-prompt": {
			Name:   "create-prompt",
			Fields: fields_create_prompt,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePromptInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_prompt, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePrompt(ctx, input)
			},
		},
		"create-push-notification-registration": {
			Name:   "create-push-notification-registration",
			Fields: fields_create_push_notification_registration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePushNotificationRegistrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_push_notification_registration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePushNotificationRegistration(ctx, input)
			},
		},
		"create-queue": {
			Name:   "create-queue",
			Fields: fields_create_queue,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateQueueInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_queue, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateQueue(ctx, input)
			},
		},
		"create-quick-connect": {
			Name:   "create-quick-connect",
			Fields: fields_create_quick_connect,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateQuickConnectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_quick_connect, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateQuickConnect(ctx, input)
			},
		},
		"create-routing-profile": {
			Name:   "create-routing-profile",
			Fields: fields_create_routing_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRoutingProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_routing_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRoutingProfile(ctx, input)
			},
		},
		"create-rule": {
			Name:   "create-rule",
			Fields: fields_create_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRule(ctx, input)
			},
		},
		"create-security-profile": {
			Name:   "create-security-profile",
			Fields: fields_create_security_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSecurityProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_security_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSecurityProfile(ctx, input)
			},
		},
		"create-task-template": {
			Name:   "create-task-template",
			Fields: fields_create_task_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTaskTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_task_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTaskTemplate(ctx, input)
			},
		},
		"create-test-case": {
			Name:   "create-test-case",
			Fields: fields_create_test_case,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTestCaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_test_case, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTestCase(ctx, input)
			},
		},
		"create-traffic-distribution-group": {
			Name:   "create-traffic-distribution-group",
			Fields: fields_create_traffic_distribution_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTrafficDistributionGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_traffic_distribution_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTrafficDistributionGroup(ctx, input)
			},
		},
		"create-use-case": {
			Name:   "create-use-case",
			Fields: fields_create_use_case,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateUseCaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_use_case, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateUseCase(ctx, input)
			},
		},
		"create-user": {
			Name:   "create-user",
			Fields: fields_create_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateUser(ctx, input)
			},
		},
		"create-user-hierarchy-group": {
			Name:   "create-user-hierarchy-group",
			Fields: fields_create_user_hierarchy_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateUserHierarchyGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_user_hierarchy_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateUserHierarchyGroup(ctx, input)
			},
		},
		"create-view": {
			Name:   "create-view",
			Fields: fields_create_view,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateViewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_view, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateView(ctx, input)
			},
		},
		"create-view-version": {
			Name:   "create-view-version",
			Fields: fields_create_view_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateViewVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_view_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateViewVersion(ctx, input)
			},
		},
		"create-vocabulary": {
			Name:   "create-vocabulary",
			Fields: fields_create_vocabulary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVocabularyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_vocabulary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVocabulary(ctx, input)
			},
		},
		"create-workspace": {
			Name:   "create-workspace",
			Fields: fields_create_workspace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWorkspaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_workspace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWorkspace(ctx, input)
			},
		},
		"create-workspace-page": {
			Name:   "create-workspace-page",
			Fields: fields_create_workspace_page,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWorkspacePageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_workspace_page, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWorkspacePage(ctx, input)
			},
		},
		"deactivate-evaluation-form": {
			Name:   "deactivate-evaluation-form",
			Fields: fields_deactivate_evaluation_form,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeactivateEvaluationFormInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deactivate_evaluation_form, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeactivateEvaluationForm(ctx, input)
			},
		},
		"delete-attached-file": {
			Name:   "delete-attached-file",
			Fields: fields_delete_attached_file,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAttachedFileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_attached_file, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAttachedFile(ctx, input)
			},
		},
		"delete-contact-evaluation": {
			Name:   "delete-contact-evaluation",
			Fields: fields_delete_contact_evaluation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteContactEvaluationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_contact_evaluation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteContactEvaluation(ctx, input)
			},
		},
		"delete-contact-flow": {
			Name:   "delete-contact-flow",
			Fields: fields_delete_contact_flow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteContactFlowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_contact_flow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteContactFlow(ctx, input)
			},
		},
		"delete-contact-flow-module": {
			Name:   "delete-contact-flow-module",
			Fields: fields_delete_contact_flow_module,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteContactFlowModuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_contact_flow_module, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteContactFlowModule(ctx, input)
			},
		},
		"delete-contact-flow-module-alias": {
			Name:   "delete-contact-flow-module-alias",
			Fields: fields_delete_contact_flow_module_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteContactFlowModuleAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_contact_flow_module_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteContactFlowModuleAlias(ctx, input)
			},
		},
		"delete-contact-flow-module-version": {
			Name:   "delete-contact-flow-module-version",
			Fields: fields_delete_contact_flow_module_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteContactFlowModuleVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_contact_flow_module_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteContactFlowModuleVersion(ctx, input)
			},
		},
		"delete-contact-flow-version": {
			Name:   "delete-contact-flow-version",
			Fields: fields_delete_contact_flow_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteContactFlowVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_contact_flow_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteContactFlowVersion(ctx, input)
			},
		},
		"delete-data-table": {
			Name:   "delete-data-table",
			Fields: fields_delete_data_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDataTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_data_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDataTable(ctx, input)
			},
		},
		"delete-data-table-attribute": {
			Name:   "delete-data-table-attribute",
			Fields: fields_delete_data_table_attribute,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDataTableAttributeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_data_table_attribute, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDataTableAttribute(ctx, input)
			},
		},
		"delete-email-address": {
			Name:   "delete-email-address",
			Fields: fields_delete_email_address,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEmailAddressInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_email_address, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEmailAddress(ctx, input)
			},
		},
		"delete-evaluation-form": {
			Name:   "delete-evaluation-form",
			Fields: fields_delete_evaluation_form,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEvaluationFormInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_evaluation_form, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEvaluationForm(ctx, input)
			},
		},
		"delete-hours-of-operation": {
			Name:   "delete-hours-of-operation",
			Fields: fields_delete_hours_of_operation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteHoursOfOperationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_hours_of_operation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteHoursOfOperation(ctx, input)
			},
		},
		"delete-hours-of-operation-override": {
			Name:   "delete-hours-of-operation-override",
			Fields: fields_delete_hours_of_operation_override,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteHoursOfOperationOverrideInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_hours_of_operation_override, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteHoursOfOperationOverride(ctx, input)
			},
		},
		"delete-instance": {
			Name:   "delete-instance",
			Fields: fields_delete_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteInstance(ctx, input)
			},
		},
		"delete-integration-association": {
			Name:   "delete-integration-association",
			Fields: fields_delete_integration_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIntegrationAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_integration_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIntegrationAssociation(ctx, input)
			},
		},
		"delete-notification": {
			Name:   "delete-notification",
			Fields: fields_delete_notification,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteNotificationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_notification, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteNotification(ctx, input)
			},
		},
		"delete-predefined-attribute": {
			Name:   "delete-predefined-attribute",
			Fields: fields_delete_predefined_attribute,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePredefinedAttributeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_predefined_attribute, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePredefinedAttribute(ctx, input)
			},
		},
		"delete-prompt": {
			Name:   "delete-prompt",
			Fields: fields_delete_prompt,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePromptInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_prompt, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePrompt(ctx, input)
			},
		},
		"delete-push-notification-registration": {
			Name:   "delete-push-notification-registration",
			Fields: fields_delete_push_notification_registration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePushNotificationRegistrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_push_notification_registration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePushNotificationRegistration(ctx, input)
			},
		},
		"delete-queue": {
			Name:   "delete-queue",
			Fields: fields_delete_queue,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteQueueInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_queue, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteQueue(ctx, input)
			},
		},
		"delete-quick-connect": {
			Name:   "delete-quick-connect",
			Fields: fields_delete_quick_connect,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteQuickConnectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_quick_connect, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteQuickConnect(ctx, input)
			},
		},
		"delete-routing-profile": {
			Name:   "delete-routing-profile",
			Fields: fields_delete_routing_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRoutingProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_routing_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRoutingProfile(ctx, input)
			},
		},
		"delete-rule": {
			Name:   "delete-rule",
			Fields: fields_delete_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRule(ctx, input)
			},
		},
		"delete-security-profile": {
			Name:   "delete-security-profile",
			Fields: fields_delete_security_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSecurityProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_security_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSecurityProfile(ctx, input)
			},
		},
		"delete-task-template": {
			Name:   "delete-task-template",
			Fields: fields_delete_task_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTaskTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_task_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTaskTemplate(ctx, input)
			},
		},
		"delete-test-case": {
			Name:   "delete-test-case",
			Fields: fields_delete_test_case,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTestCaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_test_case, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTestCase(ctx, input)
			},
		},
		"delete-traffic-distribution-group": {
			Name:   "delete-traffic-distribution-group",
			Fields: fields_delete_traffic_distribution_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTrafficDistributionGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_traffic_distribution_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTrafficDistributionGroup(ctx, input)
			},
		},
		"delete-use-case": {
			Name:   "delete-use-case",
			Fields: fields_delete_use_case,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteUseCaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_use_case, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteUseCase(ctx, input)
			},
		},
		"delete-user": {
			Name:   "delete-user",
			Fields: fields_delete_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteUser(ctx, input)
			},
		},
		"delete-user-hierarchy-group": {
			Name:   "delete-user-hierarchy-group",
			Fields: fields_delete_user_hierarchy_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteUserHierarchyGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_user_hierarchy_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteUserHierarchyGroup(ctx, input)
			},
		},
		"delete-view": {
			Name:   "delete-view",
			Fields: fields_delete_view,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteViewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_view, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteView(ctx, input)
			},
		},
		"delete-view-version": {
			Name:   "delete-view-version",
			Fields: fields_delete_view_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteViewVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_view_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteViewVersion(ctx, input)
			},
		},
		"delete-vocabulary": {
			Name:   "delete-vocabulary",
			Fields: fields_delete_vocabulary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVocabularyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_vocabulary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVocabulary(ctx, input)
			},
		},
		"delete-workspace": {
			Name:   "delete-workspace",
			Fields: fields_delete_workspace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWorkspaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_workspace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWorkspace(ctx, input)
			},
		},
		"delete-workspace-media": {
			Name:   "delete-workspace-media",
			Fields: fields_delete_workspace_media,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWorkspaceMediaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_workspace_media, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWorkspaceMedia(ctx, input)
			},
		},
		"delete-workspace-page": {
			Name:   "delete-workspace-page",
			Fields: fields_delete_workspace_page,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWorkspacePageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_workspace_page, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWorkspacePage(ctx, input)
			},
		},
		"describe-agent-status": {
			Name:   "describe-agent-status",
			Fields: fields_describe_agent_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAgentStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_agent_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAgentStatus(ctx, input)
			},
		},
		"describe-authentication-profile": {
			Name:   "describe-authentication-profile",
			Fields: fields_describe_authentication_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAuthenticationProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_authentication_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAuthenticationProfile(ctx, input)
			},
		},
		"describe-contact": {
			Name:   "describe-contact",
			Fields: fields_describe_contact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeContactInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_contact, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeContact(ctx, input)
			},
		},
		"describe-contact-evaluation": {
			Name:   "describe-contact-evaluation",
			Fields: fields_describe_contact_evaluation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeContactEvaluationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_contact_evaluation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeContactEvaluation(ctx, input)
			},
		},
		"describe-contact-flow": {
			Name:   "describe-contact-flow",
			Fields: fields_describe_contact_flow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeContactFlowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_contact_flow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeContactFlow(ctx, input)
			},
		},
		"describe-contact-flow-module": {
			Name:   "describe-contact-flow-module",
			Fields: fields_describe_contact_flow_module,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeContactFlowModuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_contact_flow_module, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeContactFlowModule(ctx, input)
			},
		},
		"describe-contact-flow-module-alias": {
			Name:   "describe-contact-flow-module-alias",
			Fields: fields_describe_contact_flow_module_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeContactFlowModuleAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_contact_flow_module_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeContactFlowModuleAlias(ctx, input)
			},
		},
		"describe-data-table": {
			Name:   "describe-data-table",
			Fields: fields_describe_data_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDataTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_data_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDataTable(ctx, input)
			},
		},
		"describe-data-table-attribute": {
			Name:   "describe-data-table-attribute",
			Fields: fields_describe_data_table_attribute,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDataTableAttributeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_data_table_attribute, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDataTableAttribute(ctx, input)
			},
		},
		"describe-email-address": {
			Name:   "describe-email-address",
			Fields: fields_describe_email_address,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEmailAddressInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_email_address, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEmailAddress(ctx, input)
			},
		},
		"describe-evaluation-form": {
			Name:   "describe-evaluation-form",
			Fields: fields_describe_evaluation_form,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEvaluationFormInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_evaluation_form, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEvaluationForm(ctx, input)
			},
		},
		"describe-hours-of-operation": {
			Name:   "describe-hours-of-operation",
			Fields: fields_describe_hours_of_operation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeHoursOfOperationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_hours_of_operation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeHoursOfOperation(ctx, input)
			},
		},
		"describe-hours-of-operation-override": {
			Name:   "describe-hours-of-operation-override",
			Fields: fields_describe_hours_of_operation_override,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeHoursOfOperationOverrideInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_hours_of_operation_override, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeHoursOfOperationOverride(ctx, input)
			},
		},
		"describe-instance": {
			Name:   "describe-instance",
			Fields: fields_describe_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeInstance(ctx, input)
			},
		},
		"describe-instance-attribute": {
			Name:   "describe-instance-attribute",
			Fields: fields_describe_instance_attribute,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInstanceAttributeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_instance_attribute, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeInstanceAttribute(ctx, input)
			},
		},
		"describe-instance-storage-config": {
			Name:   "describe-instance-storage-config",
			Fields: fields_describe_instance_storage_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInstanceStorageConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_instance_storage_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeInstanceStorageConfig(ctx, input)
			},
		},
		"describe-notification": {
			Name:   "describe-notification",
			Fields: fields_describe_notification,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeNotificationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_notification, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeNotification(ctx, input)
			},
		},
		"describe-phone-number": {
			Name:   "describe-phone-number",
			Fields: fields_describe_phone_number,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePhoneNumberInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_phone_number, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribePhoneNumber(ctx, input)
			},
		},
		"describe-predefined-attribute": {
			Name:   "describe-predefined-attribute",
			Fields: fields_describe_predefined_attribute,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePredefinedAttributeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_predefined_attribute, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribePredefinedAttribute(ctx, input)
			},
		},
		"describe-prompt": {
			Name:   "describe-prompt",
			Fields: fields_describe_prompt,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePromptInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_prompt, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribePrompt(ctx, input)
			},
		},
		"describe-queue": {
			Name:   "describe-queue",
			Fields: fields_describe_queue,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeQueueInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_queue, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeQueue(ctx, input)
			},
		},
		"describe-quick-connect": {
			Name:   "describe-quick-connect",
			Fields: fields_describe_quick_connect,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeQuickConnectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_quick_connect, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeQuickConnect(ctx, input)
			},
		},
		"describe-routing-profile": {
			Name:   "describe-routing-profile",
			Fields: fields_describe_routing_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRoutingProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_routing_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeRoutingProfile(ctx, input)
			},
		},
		"describe-rule": {
			Name:   "describe-rule",
			Fields: fields_describe_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeRule(ctx, input)
			},
		},
		"describe-security-profile": {
			Name:   "describe-security-profile",
			Fields: fields_describe_security_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSecurityProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_security_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSecurityProfile(ctx, input)
			},
		},
		"describe-test-case": {
			Name:   "describe-test-case",
			Fields: fields_describe_test_case,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTestCaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_test_case, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTestCase(ctx, input)
			},
		},
		"describe-traffic-distribution-group": {
			Name:   "describe-traffic-distribution-group",
			Fields: fields_describe_traffic_distribution_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTrafficDistributionGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_traffic_distribution_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTrafficDistributionGroup(ctx, input)
			},
		},
		"describe-user": {
			Name:   "describe-user",
			Fields: fields_describe_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeUser(ctx, input)
			},
		},
		"describe-user-hierarchy-group": {
			Name:   "describe-user-hierarchy-group",
			Fields: fields_describe_user_hierarchy_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeUserHierarchyGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_user_hierarchy_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeUserHierarchyGroup(ctx, input)
			},
		},
		"describe-user-hierarchy-structure": {
			Name:   "describe-user-hierarchy-structure",
			Fields: fields_describe_user_hierarchy_structure,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeUserHierarchyStructureInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_user_hierarchy_structure, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeUserHierarchyStructure(ctx, input)
			},
		},
		"describe-view": {
			Name:   "describe-view",
			Fields: fields_describe_view,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeViewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_view, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeView(ctx, input)
			},
		},
		"describe-vocabulary": {
			Name:   "describe-vocabulary",
			Fields: fields_describe_vocabulary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVocabularyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_vocabulary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeVocabulary(ctx, input)
			},
		},
		"describe-workspace": {
			Name:   "describe-workspace",
			Fields: fields_describe_workspace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeWorkspaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_workspace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeWorkspace(ctx, input)
			},
		},
		"disassociate-analytics-data-set": {
			Name:   "disassociate-analytics-data-set",
			Fields: fields_disassociate_analytics_data_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateAnalyticsDataSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_analytics_data_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateAnalyticsDataSet(ctx, input)
			},
		},
		"disassociate-approved-origin": {
			Name:   "disassociate-approved-origin",
			Fields: fields_disassociate_approved_origin,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateApprovedOriginInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_approved_origin, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateApprovedOrigin(ctx, input)
			},
		},
		"disassociate-bot": {
			Name:   "disassociate-bot",
			Fields: fields_disassociate_bot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateBotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_bot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateBot(ctx, input)
			},
		},
		"disassociate-email-address-alias": {
			Name:   "disassociate-email-address-alias",
			Fields: fields_disassociate_email_address_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateEmailAddressAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_email_address_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateEmailAddressAlias(ctx, input)
			},
		},
		"disassociate-flow": {
			Name:   "disassociate-flow",
			Fields: fields_disassociate_flow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateFlowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_flow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateFlow(ctx, input)
			},
		},
		"disassociate-hours-of-operations": {
			Name:   "disassociate-hours-of-operations",
			Fields: fields_disassociate_hours_of_operations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateHoursOfOperationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_hours_of_operations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateHoursOfOperations(ctx, input)
			},
		},
		"disassociate-instance-storage-config": {
			Name:   "disassociate-instance-storage-config",
			Fields: fields_disassociate_instance_storage_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateInstanceStorageConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_instance_storage_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateInstanceStorageConfig(ctx, input)
			},
		},
		"disassociate-lambda-function": {
			Name:   "disassociate-lambda-function",
			Fields: fields_disassociate_lambda_function,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateLambdaFunctionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_lambda_function, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateLambdaFunction(ctx, input)
			},
		},
		"disassociate-lex-bot": {
			Name:   "disassociate-lex-bot",
			Fields: fields_disassociate_lex_bot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateLexBotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_lex_bot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateLexBot(ctx, input)
			},
		},
		"disassociate-phone-number-contact-flow": {
			Name:   "disassociate-phone-number-contact-flow",
			Fields: fields_disassociate_phone_number_contact_flow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociatePhoneNumberContactFlowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_phone_number_contact_flow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociatePhoneNumberContactFlow(ctx, input)
			},
		},
		"disassociate-queue-quick-connects": {
			Name:   "disassociate-queue-quick-connects",
			Fields: fields_disassociate_queue_quick_connects,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateQueueQuickConnectsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_queue_quick_connects, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateQueueQuickConnects(ctx, input)
			},
		},
		"disassociate-routing-profile-queues": {
			Name:   "disassociate-routing-profile-queues",
			Fields: fields_disassociate_routing_profile_queues,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateRoutingProfileQueuesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_routing_profile_queues, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateRoutingProfileQueues(ctx, input)
			},
		},
		"disassociate-security-key": {
			Name:   "disassociate-security-key",
			Fields: fields_disassociate_security_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateSecurityKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_security_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateSecurityKey(ctx, input)
			},
		},
		"disassociate-security-profiles": {
			Name:   "disassociate-security-profiles",
			Fields: fields_disassociate_security_profiles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateSecurityProfilesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_security_profiles, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateSecurityProfiles(ctx, input)
			},
		},
		"disassociate-traffic-distribution-group-user": {
			Name:   "disassociate-traffic-distribution-group-user",
			Fields: fields_disassociate_traffic_distribution_group_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateTrafficDistributionGroupUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_traffic_distribution_group_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateTrafficDistributionGroupUser(ctx, input)
			},
		},
		"disassociate-user-proficiencies": {
			Name:   "disassociate-user-proficiencies",
			Fields: fields_disassociate_user_proficiencies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateUserProficienciesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_user_proficiencies, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateUserProficiencies(ctx, input)
			},
		},
		"disassociate-workspace": {
			Name:   "disassociate-workspace",
			Fields: fields_disassociate_workspace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateWorkspaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_workspace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateWorkspace(ctx, input)
			},
		},
		"dismiss-user-contact": {
			Name:   "dismiss-user-contact",
			Fields: fields_dismiss_user_contact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DismissUserContactInput{}
				if _, err := leanruntime.ApplyInput(input, fields_dismiss_user_contact, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DismissUserContact(ctx, input)
			},
		},
		"evaluate-data-table-values": {
			Name:   "evaluate-data-table-values",
			Fields: fields_evaluate_data_table_values,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EvaluateDataTableValuesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_evaluate_data_table_values, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.EvaluateDataTableValues(ctx, input)
				}
				var results []*svc.EvaluateDataTableValuesOutput
				p := svc.NewEvaluateDataTableValuesPaginator(client, input)
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
		"get-attached-file": {
			Name:   "get-attached-file",
			Fields: fields_get_attached_file,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAttachedFileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_attached_file, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAttachedFile(ctx, input)
			},
		},
		"get-contact-attributes": {
			Name:   "get-contact-attributes",
			Fields: fields_get_contact_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetContactAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_contact_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetContactAttributes(ctx, input)
			},
		},
		"get-contact-metrics": {
			Name:   "get-contact-metrics",
			Fields: fields_get_contact_metrics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetContactMetricsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_contact_metrics, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetContactMetrics(ctx, input)
			},
		},
		"get-current-metric-data": {
			Name:   "get-current-metric-data",
			Fields: fields_get_current_metric_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCurrentMetricDataInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_current_metric_data, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetCurrentMetricData(ctx, input)
				}
				var results []*svc.GetCurrentMetricDataOutput
				p := svc.NewGetCurrentMetricDataPaginator(client, input)
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
		"get-current-user-data": {
			Name:   "get-current-user-data",
			Fields: fields_get_current_user_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCurrentUserDataInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_current_user_data, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetCurrentUserData(ctx, input)
				}
				var results []*svc.GetCurrentUserDataOutput
				p := svc.NewGetCurrentUserDataPaginator(client, input)
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
		"get-effective-hours-of-operations": {
			Name:   "get-effective-hours-of-operations",
			Fields: fields_get_effective_hours_of_operations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEffectiveHoursOfOperationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_effective_hours_of_operations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEffectiveHoursOfOperations(ctx, input)
			},
		},
		"get-federation-token": {
			Name:   "get-federation-token",
			Fields: fields_get_federation_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFederationTokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_federation_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFederationToken(ctx, input)
			},
		},
		"get-flow-association": {
			Name:   "get-flow-association",
			Fields: fields_get_flow_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFlowAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_flow_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFlowAssociation(ctx, input)
			},
		},
		"get-metric-data": {
			Name:   "get-metric-data",
			Fields: fields_get_metric_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMetricDataInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_metric_data, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetMetricData(ctx, input)
				}
				var results []*svc.GetMetricDataOutput
				p := svc.NewGetMetricDataPaginator(client, input)
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
		"get-metric-data-v2": {
			Name:   "get-metric-data-v2",
			Fields: fields_get_metric_data_v2,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMetricDataV2Input{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_metric_data_v2, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetMetricDataV2(ctx, input)
				}
				var results []*svc.GetMetricDataV2Output
				p := svc.NewGetMetricDataV2Paginator(client, input)
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
		"get-prompt-file": {
			Name:   "get-prompt-file",
			Fields: fields_get_prompt_file,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPromptFileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_prompt_file, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPromptFile(ctx, input)
			},
		},
		"get-task-template": {
			Name:   "get-task-template",
			Fields: fields_get_task_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTaskTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_task_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTaskTemplate(ctx, input)
			},
		},
		"get-test-case-execution-summary": {
			Name:   "get-test-case-execution-summary",
			Fields: fields_get_test_case_execution_summary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTestCaseExecutionSummaryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_test_case_execution_summary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTestCaseExecutionSummary(ctx, input)
			},
		},
		"get-traffic-distribution": {
			Name:   "get-traffic-distribution",
			Fields: fields_get_traffic_distribution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTrafficDistributionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_traffic_distribution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTrafficDistribution(ctx, input)
			},
		},
		"import-phone-number": {
			Name:   "import-phone-number",
			Fields: fields_import_phone_number,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportPhoneNumberInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_phone_number, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportPhoneNumber(ctx, input)
			},
		},
		"import-workspace-media": {
			Name:   "import-workspace-media",
			Fields: fields_import_workspace_media,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportWorkspaceMediaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_workspace_media, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportWorkspaceMedia(ctx, input)
			},
		},
		"list-agent-statuses": {
			Name:   "list-agent-statuses",
			Fields: fields_list_agent_statuses,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAgentStatusesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_agent_statuses, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAgentStatuses(ctx, input)
				}
				var results []*svc.ListAgentStatusesOutput
				p := svc.NewListAgentStatusesPaginator(client, input)
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
		"list-analytics-data-associations": {
			Name:   "list-analytics-data-associations",
			Fields: fields_list_analytics_data_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAnalyticsDataAssociationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_analytics_data_associations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListAnalyticsDataAssociations(ctx, input)
			},
		},
		"list-analytics-data-lake-data-sets": {
			Name:   "list-analytics-data-lake-data-sets",
			Fields: fields_list_analytics_data_lake_data_sets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAnalyticsDataLakeDataSetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_analytics_data_lake_data_sets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListAnalyticsDataLakeDataSets(ctx, input)
			},
		},
		"list-approved-origins": {
			Name:   "list-approved-origins",
			Fields: fields_list_approved_origins,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListApprovedOriginsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_approved_origins, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListApprovedOrigins(ctx, input)
				}
				var results []*svc.ListApprovedOriginsOutput
				p := svc.NewListApprovedOriginsPaginator(client, input)
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
		"list-associated-contacts": {
			Name:   "list-associated-contacts",
			Fields: fields_list_associated_contacts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAssociatedContactsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_associated_contacts, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListAssociatedContacts(ctx, input)
			},
		},
		"list-authentication-profiles": {
			Name:   "list-authentication-profiles",
			Fields: fields_list_authentication_profiles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAuthenticationProfilesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_authentication_profiles, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAuthenticationProfiles(ctx, input)
				}
				var results []*svc.ListAuthenticationProfilesOutput
				p := svc.NewListAuthenticationProfilesPaginator(client, input)
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
		"list-bots": {
			Name:   "list-bots",
			Fields: fields_list_bots,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBotsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_bots, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBots(ctx, input)
				}
				var results []*svc.ListBotsOutput
				p := svc.NewListBotsPaginator(client, input)
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
		"list-child-hours-of-operations": {
			Name:   "list-child-hours-of-operations",
			Fields: fields_list_child_hours_of_operations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListChildHoursOfOperationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_child_hours_of_operations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListChildHoursOfOperations(ctx, input)
				}
				var results []*svc.ListChildHoursOfOperationsOutput
				p := svc.NewListChildHoursOfOperationsPaginator(client, input)
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
		"list-contact-evaluations": {
			Name:   "list-contact-evaluations",
			Fields: fields_list_contact_evaluations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListContactEvaluationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_contact_evaluations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListContactEvaluations(ctx, input)
				}
				var results []*svc.ListContactEvaluationsOutput
				p := svc.NewListContactEvaluationsPaginator(client, input)
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
		"list-contact-flow-module-aliases": {
			Name:   "list-contact-flow-module-aliases",
			Fields: fields_list_contact_flow_module_aliases,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListContactFlowModuleAliasesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_contact_flow_module_aliases, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListContactFlowModuleAliases(ctx, input)
				}
				var results []*svc.ListContactFlowModuleAliasesOutput
				p := svc.NewListContactFlowModuleAliasesPaginator(client, input)
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
		"list-contact-flow-module-versions": {
			Name:   "list-contact-flow-module-versions",
			Fields: fields_list_contact_flow_module_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListContactFlowModuleVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_contact_flow_module_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListContactFlowModuleVersions(ctx, input)
				}
				var results []*svc.ListContactFlowModuleVersionsOutput
				p := svc.NewListContactFlowModuleVersionsPaginator(client, input)
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
		"list-contact-flow-modules": {
			Name:   "list-contact-flow-modules",
			Fields: fields_list_contact_flow_modules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListContactFlowModulesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_contact_flow_modules, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListContactFlowModules(ctx, input)
				}
				var results []*svc.ListContactFlowModulesOutput
				p := svc.NewListContactFlowModulesPaginator(client, input)
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
		"list-contact-flow-versions": {
			Name:   "list-contact-flow-versions",
			Fields: fields_list_contact_flow_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListContactFlowVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_contact_flow_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListContactFlowVersions(ctx, input)
				}
				var results []*svc.ListContactFlowVersionsOutput
				p := svc.NewListContactFlowVersionsPaginator(client, input)
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
		"list-contact-flows": {
			Name:   "list-contact-flows",
			Fields: fields_list_contact_flows,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListContactFlowsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_contact_flows, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListContactFlows(ctx, input)
				}
				var results []*svc.ListContactFlowsOutput
				p := svc.NewListContactFlowsPaginator(client, input)
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
		"list-contact-references": {
			Name:   "list-contact-references",
			Fields: fields_list_contact_references,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListContactReferencesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_contact_references, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListContactReferences(ctx, input)
				}
				var results []*svc.ListContactReferencesOutput
				p := svc.NewListContactReferencesPaginator(client, input)
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
		"list-data-table-attributes": {
			Name:   "list-data-table-attributes",
			Fields: fields_list_data_table_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDataTableAttributesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_data_table_attributes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDataTableAttributes(ctx, input)
				}
				var results []*svc.ListDataTableAttributesOutput
				p := svc.NewListDataTableAttributesPaginator(client, input)
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
		"list-data-table-primary-values": {
			Name:   "list-data-table-primary-values",
			Fields: fields_list_data_table_primary_values,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDataTablePrimaryValuesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_data_table_primary_values, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDataTablePrimaryValues(ctx, input)
				}
				var results []*svc.ListDataTablePrimaryValuesOutput
				p := svc.NewListDataTablePrimaryValuesPaginator(client, input)
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
		"list-data-table-values": {
			Name:   "list-data-table-values",
			Fields: fields_list_data_table_values,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDataTableValuesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_data_table_values, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDataTableValues(ctx, input)
				}
				var results []*svc.ListDataTableValuesOutput
				p := svc.NewListDataTableValuesPaginator(client, input)
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
		"list-data-tables": {
			Name:   "list-data-tables",
			Fields: fields_list_data_tables,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDataTablesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_data_tables, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDataTables(ctx, input)
				}
				var results []*svc.ListDataTablesOutput
				p := svc.NewListDataTablesPaginator(client, input)
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
		"list-default-vocabularies": {
			Name:   "list-default-vocabularies",
			Fields: fields_list_default_vocabularies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDefaultVocabulariesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_default_vocabularies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDefaultVocabularies(ctx, input)
				}
				var results []*svc.ListDefaultVocabulariesOutput
				p := svc.NewListDefaultVocabulariesPaginator(client, input)
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
		"list-entity-security-profiles": {
			Name:   "list-entity-security-profiles",
			Fields: fields_list_entity_security_profiles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEntitySecurityProfilesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_entity_security_profiles, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEntitySecurityProfiles(ctx, input)
				}
				var results []*svc.ListEntitySecurityProfilesOutput
				p := svc.NewListEntitySecurityProfilesPaginator(client, input)
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
		"list-evaluation-form-versions": {
			Name:   "list-evaluation-form-versions",
			Fields: fields_list_evaluation_form_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEvaluationFormVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_evaluation_form_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEvaluationFormVersions(ctx, input)
				}
				var results []*svc.ListEvaluationFormVersionsOutput
				p := svc.NewListEvaluationFormVersionsPaginator(client, input)
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
		"list-evaluation-forms": {
			Name:   "list-evaluation-forms",
			Fields: fields_list_evaluation_forms,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEvaluationFormsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_evaluation_forms, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEvaluationForms(ctx, input)
				}
				var results []*svc.ListEvaluationFormsOutput
				p := svc.NewListEvaluationFormsPaginator(client, input)
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
		"list-flow-associations": {
			Name:   "list-flow-associations",
			Fields: fields_list_flow_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFlowAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_flow_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFlowAssociations(ctx, input)
				}
				var results []*svc.ListFlowAssociationsOutput
				p := svc.NewListFlowAssociationsPaginator(client, input)
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
		"list-hours-of-operation-overrides": {
			Name:   "list-hours-of-operation-overrides",
			Fields: fields_list_hours_of_operation_overrides,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListHoursOfOperationOverridesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_hours_of_operation_overrides, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListHoursOfOperationOverrides(ctx, input)
				}
				var results []*svc.ListHoursOfOperationOverridesOutput
				p := svc.NewListHoursOfOperationOverridesPaginator(client, input)
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
		"list-hours-of-operations": {
			Name:   "list-hours-of-operations",
			Fields: fields_list_hours_of_operations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListHoursOfOperationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_hours_of_operations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListHoursOfOperations(ctx, input)
				}
				var results []*svc.ListHoursOfOperationsOutput
				p := svc.NewListHoursOfOperationsPaginator(client, input)
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
		"list-instance-attributes": {
			Name:   "list-instance-attributes",
			Fields: fields_list_instance_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInstanceAttributesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_instance_attributes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListInstanceAttributes(ctx, input)
				}
				var results []*svc.ListInstanceAttributesOutput
				p := svc.NewListInstanceAttributesPaginator(client, input)
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
		"list-instance-storage-configs": {
			Name:   "list-instance-storage-configs",
			Fields: fields_list_instance_storage_configs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInstanceStorageConfigsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_instance_storage_configs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListInstanceStorageConfigs(ctx, input)
				}
				var results []*svc.ListInstanceStorageConfigsOutput
				p := svc.NewListInstanceStorageConfigsPaginator(client, input)
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
		"list-instances": {
			Name:   "list-instances",
			Fields: fields_list_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInstancesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_instances, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListInstances(ctx, input)
				}
				var results []*svc.ListInstancesOutput
				p := svc.NewListInstancesPaginator(client, input)
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
		"list-integration-associations": {
			Name:   "list-integration-associations",
			Fields: fields_list_integration_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIntegrationAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_integration_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListIntegrationAssociations(ctx, input)
				}
				var results []*svc.ListIntegrationAssociationsOutput
				p := svc.NewListIntegrationAssociationsPaginator(client, input)
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
		"list-lambda-functions": {
			Name:   "list-lambda-functions",
			Fields: fields_list_lambda_functions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLambdaFunctionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_lambda_functions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLambdaFunctions(ctx, input)
				}
				var results []*svc.ListLambdaFunctionsOutput
				p := svc.NewListLambdaFunctionsPaginator(client, input)
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
		"list-lex-bots": {
			Name:   "list-lex-bots",
			Fields: fields_list_lex_bots,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLexBotsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_lex_bots, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLexBots(ctx, input)
				}
				var results []*svc.ListLexBotsOutput
				p := svc.NewListLexBotsPaginator(client, input)
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
		"list-notifications": {
			Name:   "list-notifications",
			Fields: fields_list_notifications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListNotificationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_notifications, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListNotifications(ctx, input)
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
		"list-phone-numbers-v2": {
			Name:   "list-phone-numbers-v2",
			Fields: fields_list_phone_numbers_v2,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPhoneNumbersV2Input{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_phone_numbers_v2, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPhoneNumbersV2(ctx, input)
				}
				var results []*svc.ListPhoneNumbersV2Output
				p := svc.NewListPhoneNumbersV2Paginator(client, input)
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
		"list-predefined-attributes": {
			Name:   "list-predefined-attributes",
			Fields: fields_list_predefined_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPredefinedAttributesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_predefined_attributes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPredefinedAttributes(ctx, input)
				}
				var results []*svc.ListPredefinedAttributesOutput
				p := svc.NewListPredefinedAttributesPaginator(client, input)
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
		"list-prompts": {
			Name:   "list-prompts",
			Fields: fields_list_prompts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPromptsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_prompts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPrompts(ctx, input)
				}
				var results []*svc.ListPromptsOutput
				p := svc.NewListPromptsPaginator(client, input)
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
		"list-queue-quick-connects": {
			Name:   "list-queue-quick-connects",
			Fields: fields_list_queue_quick_connects,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListQueueQuickConnectsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_queue_quick_connects, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListQueueQuickConnects(ctx, input)
				}
				var results []*svc.ListQueueQuickConnectsOutput
				p := svc.NewListQueueQuickConnectsPaginator(client, input)
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
		"list-queues": {
			Name:   "list-queues",
			Fields: fields_list_queues,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListQueuesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_queues, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListQueues(ctx, input)
				}
				var results []*svc.ListQueuesOutput
				p := svc.NewListQueuesPaginator(client, input)
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
		"list-quick-connects": {
			Name:   "list-quick-connects",
			Fields: fields_list_quick_connects,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListQuickConnectsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_quick_connects, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListQuickConnects(ctx, input)
				}
				var results []*svc.ListQuickConnectsOutput
				p := svc.NewListQuickConnectsPaginator(client, input)
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
		"list-realtime-contact-analysis-segments-v2": {
			Name:   "list-realtime-contact-analysis-segments-v2",
			Fields: fields_list_realtime_contact_analysis_segments_v2,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRealtimeContactAnalysisSegmentsV2Input{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_realtime_contact_analysis_segments_v2, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRealtimeContactAnalysisSegmentsV2(ctx, input)
				}
				var results []*svc.ListRealtimeContactAnalysisSegmentsV2Output
				p := svc.NewListRealtimeContactAnalysisSegmentsV2Paginator(client, input)
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
		"list-routing-profile-manual-assignment-queues": {
			Name:   "list-routing-profile-manual-assignment-queues",
			Fields: fields_list_routing_profile_manual_assignment_queues,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRoutingProfileManualAssignmentQueuesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_routing_profile_manual_assignment_queues, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRoutingProfileManualAssignmentQueues(ctx, input)
				}
				var results []*svc.ListRoutingProfileManualAssignmentQueuesOutput
				p := svc.NewListRoutingProfileManualAssignmentQueuesPaginator(client, input)
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
		"list-routing-profile-queues": {
			Name:   "list-routing-profile-queues",
			Fields: fields_list_routing_profile_queues,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRoutingProfileQueuesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_routing_profile_queues, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRoutingProfileQueues(ctx, input)
				}
				var results []*svc.ListRoutingProfileQueuesOutput
				p := svc.NewListRoutingProfileQueuesPaginator(client, input)
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
		"list-routing-profiles": {
			Name:   "list-routing-profiles",
			Fields: fields_list_routing_profiles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRoutingProfilesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_routing_profiles, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRoutingProfiles(ctx, input)
				}
				var results []*svc.ListRoutingProfilesOutput
				p := svc.NewListRoutingProfilesPaginator(client, input)
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
		"list-rules": {
			Name:   "list-rules",
			Fields: fields_list_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRulesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_rules, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRules(ctx, input)
				}
				var results []*svc.ListRulesOutput
				p := svc.NewListRulesPaginator(client, input)
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
		"list-security-keys": {
			Name:   "list-security-keys",
			Fields: fields_list_security_keys,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSecurityKeysInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_security_keys, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSecurityKeys(ctx, input)
				}
				var results []*svc.ListSecurityKeysOutput
				p := svc.NewListSecurityKeysPaginator(client, input)
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
		"list-security-profile-applications": {
			Name:   "list-security-profile-applications",
			Fields: fields_list_security_profile_applications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSecurityProfileApplicationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_security_profile_applications, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSecurityProfileApplications(ctx, input)
				}
				var results []*svc.ListSecurityProfileApplicationsOutput
				p := svc.NewListSecurityProfileApplicationsPaginator(client, input)
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
		"list-security-profile-flow-modules": {
			Name:   "list-security-profile-flow-modules",
			Fields: fields_list_security_profile_flow_modules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSecurityProfileFlowModulesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_security_profile_flow_modules, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSecurityProfileFlowModules(ctx, input)
				}
				var results []*svc.ListSecurityProfileFlowModulesOutput
				p := svc.NewListSecurityProfileFlowModulesPaginator(client, input)
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
		"list-security-profile-permissions": {
			Name:   "list-security-profile-permissions",
			Fields: fields_list_security_profile_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSecurityProfilePermissionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_security_profile_permissions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSecurityProfilePermissions(ctx, input)
				}
				var results []*svc.ListSecurityProfilePermissionsOutput
				p := svc.NewListSecurityProfilePermissionsPaginator(client, input)
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
		"list-security-profiles": {
			Name:   "list-security-profiles",
			Fields: fields_list_security_profiles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSecurityProfilesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_security_profiles, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSecurityProfiles(ctx, input)
				}
				var results []*svc.ListSecurityProfilesOutput
				p := svc.NewListSecurityProfilesPaginator(client, input)
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
		"list-task-templates": {
			Name:   "list-task-templates",
			Fields: fields_list_task_templates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTaskTemplatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_task_templates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTaskTemplates(ctx, input)
				}
				var results []*svc.ListTaskTemplatesOutput
				p := svc.NewListTaskTemplatesPaginator(client, input)
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
		"list-test-case-execution-records": {
			Name:   "list-test-case-execution-records",
			Fields: fields_list_test_case_execution_records,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTestCaseExecutionRecordsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_test_case_execution_records, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTestCaseExecutionRecords(ctx, input)
			},
		},
		"list-test-case-executions": {
			Name:   "list-test-case-executions",
			Fields: fields_list_test_case_executions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTestCaseExecutionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_test_case_executions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTestCaseExecutions(ctx, input)
			},
		},
		"list-test-cases": {
			Name:   "list-test-cases",
			Fields: fields_list_test_cases,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTestCasesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_test_cases, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTestCases(ctx, input)
				}
				var results []*svc.ListTestCasesOutput
				p := svc.NewListTestCasesPaginator(client, input)
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
		"list-traffic-distribution-group-users": {
			Name:   "list-traffic-distribution-group-users",
			Fields: fields_list_traffic_distribution_group_users,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTrafficDistributionGroupUsersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_traffic_distribution_group_users, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTrafficDistributionGroupUsers(ctx, input)
				}
				var results []*svc.ListTrafficDistributionGroupUsersOutput
				p := svc.NewListTrafficDistributionGroupUsersPaginator(client, input)
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
		"list-traffic-distribution-groups": {
			Name:   "list-traffic-distribution-groups",
			Fields: fields_list_traffic_distribution_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTrafficDistributionGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_traffic_distribution_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTrafficDistributionGroups(ctx, input)
				}
				var results []*svc.ListTrafficDistributionGroupsOutput
				p := svc.NewListTrafficDistributionGroupsPaginator(client, input)
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
		"list-use-cases": {
			Name:   "list-use-cases",
			Fields: fields_list_use_cases,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListUseCasesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_use_cases, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListUseCases(ctx, input)
				}
				var results []*svc.ListUseCasesOutput
				p := svc.NewListUseCasesPaginator(client, input)
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
		"list-user-hierarchy-groups": {
			Name:   "list-user-hierarchy-groups",
			Fields: fields_list_user_hierarchy_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListUserHierarchyGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_user_hierarchy_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListUserHierarchyGroups(ctx, input)
				}
				var results []*svc.ListUserHierarchyGroupsOutput
				p := svc.NewListUserHierarchyGroupsPaginator(client, input)
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
		"list-user-notifications": {
			Name:   "list-user-notifications",
			Fields: fields_list_user_notifications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListUserNotificationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_user_notifications, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListUserNotifications(ctx, input)
			},
		},
		"list-user-proficiencies": {
			Name:   "list-user-proficiencies",
			Fields: fields_list_user_proficiencies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListUserProficienciesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_user_proficiencies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListUserProficiencies(ctx, input)
				}
				var results []*svc.ListUserProficienciesOutput
				p := svc.NewListUserProficienciesPaginator(client, input)
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
		"list-users": {
			Name:   "list-users",
			Fields: fields_list_users,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListUsersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_users, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListUsers(ctx, input)
				}
				var results []*svc.ListUsersOutput
				p := svc.NewListUsersPaginator(client, input)
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
		"list-view-versions": {
			Name:   "list-view-versions",
			Fields: fields_list_view_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListViewVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_view_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListViewVersions(ctx, input)
				}
				var results []*svc.ListViewVersionsOutput
				p := svc.NewListViewVersionsPaginator(client, input)
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
		"list-views": {
			Name:   "list-views",
			Fields: fields_list_views,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListViewsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_views, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListViews(ctx, input)
				}
				var results []*svc.ListViewsOutput
				p := svc.NewListViewsPaginator(client, input)
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
		"list-workspace-media": {
			Name:   "list-workspace-media",
			Fields: fields_list_workspace_media,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWorkspaceMediaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_workspace_media, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListWorkspaceMedia(ctx, input)
			},
		},
		"list-workspace-pages": {
			Name:   "list-workspace-pages",
			Fields: fields_list_workspace_pages,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWorkspacePagesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_workspace_pages, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWorkspacePages(ctx, input)
				}
				var results []*svc.ListWorkspacePagesOutput
				p := svc.NewListWorkspacePagesPaginator(client, input)
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
		"list-workspaces": {
			Name:   "list-workspaces",
			Fields: fields_list_workspaces,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWorkspacesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_workspaces, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWorkspaces(ctx, input)
				}
				var results []*svc.ListWorkspacesOutput
				p := svc.NewListWorkspacesPaginator(client, input)
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
		"monitor-contact": {
			Name:   "monitor-contact",
			Fields: fields_monitor_contact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.MonitorContactInput{}
				if _, err := leanruntime.ApplyInput(input, fields_monitor_contact, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.MonitorContact(ctx, input)
			},
		},
		"pause-contact": {
			Name:   "pause-contact",
			Fields: fields_pause_contact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PauseContactInput{}
				if _, err := leanruntime.ApplyInput(input, fields_pause_contact, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PauseContact(ctx, input)
			},
		},
		"put-user-status": {
			Name:   "put-user-status",
			Fields: fields_put_user_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutUserStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_user_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutUserStatus(ctx, input)
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
		"replicate-instance": {
			Name:   "replicate-instance",
			Fields: fields_replicate_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ReplicateInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_replicate_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ReplicateInstance(ctx, input)
			},
		},
		"resume-contact": {
			Name:   "resume-contact",
			Fields: fields_resume_contact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResumeContactInput{}
				if _, err := leanruntime.ApplyInput(input, fields_resume_contact, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResumeContact(ctx, input)
			},
		},
		"resume-contact-recording": {
			Name:   "resume-contact-recording",
			Fields: fields_resume_contact_recording,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResumeContactRecordingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_resume_contact_recording, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResumeContactRecording(ctx, input)
			},
		},
		"search-agent-statuses": {
			Name:   "search-agent-statuses",
			Fields: fields_search_agent_statuses,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchAgentStatusesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_agent_statuses, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchAgentStatuses(ctx, input)
				}
				var results []*svc.SearchAgentStatusesOutput
				p := svc.NewSearchAgentStatusesPaginator(client, input)
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
		"search-contact-evaluations": {
			Name:   "search-contact-evaluations",
			Fields: fields_search_contact_evaluations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchContactEvaluationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_search_contact_evaluations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SearchContactEvaluations(ctx, input)
			},
		},
		"search-contact-flow-modules": {
			Name:   "search-contact-flow-modules",
			Fields: fields_search_contact_flow_modules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchContactFlowModulesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_contact_flow_modules, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchContactFlowModules(ctx, input)
				}
				var results []*svc.SearchContactFlowModulesOutput
				p := svc.NewSearchContactFlowModulesPaginator(client, input)
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
		"search-contact-flows": {
			Name:   "search-contact-flows",
			Fields: fields_search_contact_flows,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchContactFlowsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_contact_flows, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchContactFlows(ctx, input)
				}
				var results []*svc.SearchContactFlowsOutput
				p := svc.NewSearchContactFlowsPaginator(client, input)
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
		"search-contacts": {
			Name:   "search-contacts",
			Fields: fields_search_contacts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchContactsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_contacts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchContacts(ctx, input)
				}
				var results []*svc.SearchContactsOutput
				p := svc.NewSearchContactsPaginator(client, input)
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
		"search-data-tables": {
			Name:   "search-data-tables",
			Fields: fields_search_data_tables,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchDataTablesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_data_tables, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchDataTables(ctx, input)
				}
				var results []*svc.SearchDataTablesOutput
				p := svc.NewSearchDataTablesPaginator(client, input)
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
		"search-email-addresses": {
			Name:   "search-email-addresses",
			Fields: fields_search_email_addresses,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchEmailAddressesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_search_email_addresses, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SearchEmailAddresses(ctx, input)
			},
		},
		"search-evaluation-forms": {
			Name:   "search-evaluation-forms",
			Fields: fields_search_evaluation_forms,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchEvaluationFormsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_search_evaluation_forms, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SearchEvaluationForms(ctx, input)
			},
		},
		"search-hours-of-operation-overrides": {
			Name:   "search-hours-of-operation-overrides",
			Fields: fields_search_hours_of_operation_overrides,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchHoursOfOperationOverridesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_hours_of_operation_overrides, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchHoursOfOperationOverrides(ctx, input)
				}
				var results []*svc.SearchHoursOfOperationOverridesOutput
				p := svc.NewSearchHoursOfOperationOverridesPaginator(client, input)
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
		"search-hours-of-operations": {
			Name:   "search-hours-of-operations",
			Fields: fields_search_hours_of_operations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchHoursOfOperationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_hours_of_operations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchHoursOfOperations(ctx, input)
				}
				var results []*svc.SearchHoursOfOperationsOutput
				p := svc.NewSearchHoursOfOperationsPaginator(client, input)
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
		"search-notifications": {
			Name:   "search-notifications",
			Fields: fields_search_notifications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchNotificationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_search_notifications, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SearchNotifications(ctx, input)
			},
		},
		"search-predefined-attributes": {
			Name:   "search-predefined-attributes",
			Fields: fields_search_predefined_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchPredefinedAttributesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_predefined_attributes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchPredefinedAttributes(ctx, input)
				}
				var results []*svc.SearchPredefinedAttributesOutput
				p := svc.NewSearchPredefinedAttributesPaginator(client, input)
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
		"search-prompts": {
			Name:   "search-prompts",
			Fields: fields_search_prompts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchPromptsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_prompts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchPrompts(ctx, input)
				}
				var results []*svc.SearchPromptsOutput
				p := svc.NewSearchPromptsPaginator(client, input)
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
		"search-queues": {
			Name:   "search-queues",
			Fields: fields_search_queues,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchQueuesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_queues, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchQueues(ctx, input)
				}
				var results []*svc.SearchQueuesOutput
				p := svc.NewSearchQueuesPaginator(client, input)
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
		"search-quick-connects": {
			Name:   "search-quick-connects",
			Fields: fields_search_quick_connects,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchQuickConnectsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_quick_connects, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchQuickConnects(ctx, input)
				}
				var results []*svc.SearchQuickConnectsOutput
				p := svc.NewSearchQuickConnectsPaginator(client, input)
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
		"search-resource-tags": {
			Name:   "search-resource-tags",
			Fields: fields_search_resource_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchResourceTagsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_resource_tags, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchResourceTags(ctx, input)
				}
				var results []*svc.SearchResourceTagsOutput
				p := svc.NewSearchResourceTagsPaginator(client, input)
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
		"search-routing-profiles": {
			Name:   "search-routing-profiles",
			Fields: fields_search_routing_profiles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchRoutingProfilesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_routing_profiles, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchRoutingProfiles(ctx, input)
				}
				var results []*svc.SearchRoutingProfilesOutput
				p := svc.NewSearchRoutingProfilesPaginator(client, input)
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
		"search-security-profiles": {
			Name:   "search-security-profiles",
			Fields: fields_search_security_profiles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchSecurityProfilesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_security_profiles, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchSecurityProfiles(ctx, input)
				}
				var results []*svc.SearchSecurityProfilesOutput
				p := svc.NewSearchSecurityProfilesPaginator(client, input)
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
		"search-test-cases": {
			Name:   "search-test-cases",
			Fields: fields_search_test_cases,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchTestCasesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_test_cases, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchTestCases(ctx, input)
				}
				var results []*svc.SearchTestCasesOutput
				p := svc.NewSearchTestCasesPaginator(client, input)
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
		"search-user-hierarchy-groups": {
			Name:   "search-user-hierarchy-groups",
			Fields: fields_search_user_hierarchy_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchUserHierarchyGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_user_hierarchy_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchUserHierarchyGroups(ctx, input)
				}
				var results []*svc.SearchUserHierarchyGroupsOutput
				p := svc.NewSearchUserHierarchyGroupsPaginator(client, input)
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
		"search-users": {
			Name:   "search-users",
			Fields: fields_search_users,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchUsersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_users, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchUsers(ctx, input)
				}
				var results []*svc.SearchUsersOutput
				p := svc.NewSearchUsersPaginator(client, input)
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
		"search-views": {
			Name:   "search-views",
			Fields: fields_search_views,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchViewsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_views, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchViews(ctx, input)
				}
				var results []*svc.SearchViewsOutput
				p := svc.NewSearchViewsPaginator(client, input)
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
		"search-vocabularies": {
			Name:   "search-vocabularies",
			Fields: fields_search_vocabularies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchVocabulariesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_vocabularies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchVocabularies(ctx, input)
				}
				var results []*svc.SearchVocabulariesOutput
				p := svc.NewSearchVocabulariesPaginator(client, input)
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
		"search-workspace-associations": {
			Name:   "search-workspace-associations",
			Fields: fields_search_workspace_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchWorkspaceAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_workspace_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchWorkspaceAssociations(ctx, input)
				}
				var results []*svc.SearchWorkspaceAssociationsOutput
				p := svc.NewSearchWorkspaceAssociationsPaginator(client, input)
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
		"search-workspaces": {
			Name:   "search-workspaces",
			Fields: fields_search_workspaces,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchWorkspacesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_workspaces, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchWorkspaces(ctx, input)
				}
				var results []*svc.SearchWorkspacesOutput
				p := svc.NewSearchWorkspacesPaginator(client, input)
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
		"send-chat-integration-event": {
			Name:   "send-chat-integration-event",
			Fields: fields_send_chat_integration_event,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendChatIntegrationEventInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_chat_integration_event, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendChatIntegrationEvent(ctx, input)
			},
		},
		"send-outbound-email": {
			Name:   "send-outbound-email",
			Fields: fields_send_outbound_email,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendOutboundEmailInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_outbound_email, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendOutboundEmail(ctx, input)
			},
		},
		"start-attached-file-upload": {
			Name:   "start-attached-file-upload",
			Fields: fields_start_attached_file_upload,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartAttachedFileUploadInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_attached_file_upload, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartAttachedFileUpload(ctx, input)
			},
		},
		"start-chat-contact": {
			Name:   "start-chat-contact",
			Fields: fields_start_chat_contact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartChatContactInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_chat_contact, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartChatContact(ctx, input)
			},
		},
		"start-contact-evaluation": {
			Name:   "start-contact-evaluation",
			Fields: fields_start_contact_evaluation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartContactEvaluationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_contact_evaluation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartContactEvaluation(ctx, input)
			},
		},
		"start-contact-media-processing": {
			Name:   "start-contact-media-processing",
			Fields: fields_start_contact_media_processing,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartContactMediaProcessingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_contact_media_processing, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartContactMediaProcessing(ctx, input)
			},
		},
		"start-contact-recording": {
			Name:   "start-contact-recording",
			Fields: fields_start_contact_recording,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartContactRecordingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_contact_recording, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartContactRecording(ctx, input)
			},
		},
		"start-contact-streaming": {
			Name:   "start-contact-streaming",
			Fields: fields_start_contact_streaming,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartContactStreamingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_contact_streaming, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartContactStreaming(ctx, input)
			},
		},
		"start-email-contact": {
			Name:   "start-email-contact",
			Fields: fields_start_email_contact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartEmailContactInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_email_contact, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartEmailContact(ctx, input)
			},
		},
		"start-outbound-chat-contact": {
			Name:   "start-outbound-chat-contact",
			Fields: fields_start_outbound_chat_contact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartOutboundChatContactInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_outbound_chat_contact, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartOutboundChatContact(ctx, input)
			},
		},
		"start-outbound-email-contact": {
			Name:   "start-outbound-email-contact",
			Fields: fields_start_outbound_email_contact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartOutboundEmailContactInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_outbound_email_contact, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartOutboundEmailContact(ctx, input)
			},
		},
		"start-outbound-voice-contact": {
			Name:   "start-outbound-voice-contact",
			Fields: fields_start_outbound_voice_contact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartOutboundVoiceContactInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_outbound_voice_contact, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartOutboundVoiceContact(ctx, input)
			},
		},
		"start-screen-sharing": {
			Name:   "start-screen-sharing",
			Fields: fields_start_screen_sharing,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartScreenSharingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_screen_sharing, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartScreenSharing(ctx, input)
			},
		},
		"start-task-contact": {
			Name:   "start-task-contact",
			Fields: fields_start_task_contact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartTaskContactInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_task_contact, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartTaskContact(ctx, input)
			},
		},
		"start-test-case-execution": {
			Name:   "start-test-case-execution",
			Fields: fields_start_test_case_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartTestCaseExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_test_case_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartTestCaseExecution(ctx, input)
			},
		},
		"start-web-rtc-contact": {
			Name:   "start-web-rtc-contact",
			Fields: fields_start_web_rtc_contact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartWebRTCContactInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_web_rtc_contact, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartWebRTCContact(ctx, input)
			},
		},
		"stop-contact": {
			Name:   "stop-contact",
			Fields: fields_stop_contact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopContactInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_contact, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopContact(ctx, input)
			},
		},
		"stop-contact-media-processing": {
			Name:   "stop-contact-media-processing",
			Fields: fields_stop_contact_media_processing,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopContactMediaProcessingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_contact_media_processing, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopContactMediaProcessing(ctx, input)
			},
		},
		"stop-contact-recording": {
			Name:   "stop-contact-recording",
			Fields: fields_stop_contact_recording,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopContactRecordingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_contact_recording, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopContactRecording(ctx, input)
			},
		},
		"stop-contact-streaming": {
			Name:   "stop-contact-streaming",
			Fields: fields_stop_contact_streaming,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopContactStreamingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_contact_streaming, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopContactStreaming(ctx, input)
			},
		},
		"stop-test-case-execution": {
			Name:   "stop-test-case-execution",
			Fields: fields_stop_test_case_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopTestCaseExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_test_case_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopTestCaseExecution(ctx, input)
			},
		},
		"submit-contact-evaluation": {
			Name:   "submit-contact-evaluation",
			Fields: fields_submit_contact_evaluation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SubmitContactEvaluationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_submit_contact_evaluation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SubmitContactEvaluation(ctx, input)
			},
		},
		"suspend-contact-recording": {
			Name:   "suspend-contact-recording",
			Fields: fields_suspend_contact_recording,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SuspendContactRecordingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_suspend_contact_recording, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SuspendContactRecording(ctx, input)
			},
		},
		"tag-contact": {
			Name:   "tag-contact",
			Fields: fields_tag_contact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TagContactInput{}
				if _, err := leanruntime.ApplyInput(input, fields_tag_contact, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TagContact(ctx, input)
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
		"transfer-contact": {
			Name:   "transfer-contact",
			Fields: fields_transfer_contact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TransferContactInput{}
				if _, err := leanruntime.ApplyInput(input, fields_transfer_contact, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TransferContact(ctx, input)
			},
		},
		"untag-contact": {
			Name:   "untag-contact",
			Fields: fields_untag_contact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UntagContactInput{}
				if _, err := leanruntime.ApplyInput(input, fields_untag_contact, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UntagContact(ctx, input)
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
		"update-agent-status": {
			Name:   "update-agent-status",
			Fields: fields_update_agent_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAgentStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_agent_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAgentStatus(ctx, input)
			},
		},
		"update-authentication-profile": {
			Name:   "update-authentication-profile",
			Fields: fields_update_authentication_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAuthenticationProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_authentication_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAuthenticationProfile(ctx, input)
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
		"update-contact-attributes": {
			Name:   "update-contact-attributes",
			Fields: fields_update_contact_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateContactAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_contact_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateContactAttributes(ctx, input)
			},
		},
		"update-contact-evaluation": {
			Name:   "update-contact-evaluation",
			Fields: fields_update_contact_evaluation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateContactEvaluationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_contact_evaluation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateContactEvaluation(ctx, input)
			},
		},
		"update-contact-flow-content": {
			Name:   "update-contact-flow-content",
			Fields: fields_update_contact_flow_content,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateContactFlowContentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_contact_flow_content, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateContactFlowContent(ctx, input)
			},
		},
		"update-contact-flow-metadata": {
			Name:   "update-contact-flow-metadata",
			Fields: fields_update_contact_flow_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateContactFlowMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_contact_flow_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateContactFlowMetadata(ctx, input)
			},
		},
		"update-contact-flow-module-alias": {
			Name:   "update-contact-flow-module-alias",
			Fields: fields_update_contact_flow_module_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateContactFlowModuleAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_contact_flow_module_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateContactFlowModuleAlias(ctx, input)
			},
		},
		"update-contact-flow-module-content": {
			Name:   "update-contact-flow-module-content",
			Fields: fields_update_contact_flow_module_content,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateContactFlowModuleContentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_contact_flow_module_content, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateContactFlowModuleContent(ctx, input)
			},
		},
		"update-contact-flow-module-metadata": {
			Name:   "update-contact-flow-module-metadata",
			Fields: fields_update_contact_flow_module_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateContactFlowModuleMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_contact_flow_module_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateContactFlowModuleMetadata(ctx, input)
			},
		},
		"update-contact-flow-name": {
			Name:   "update-contact-flow-name",
			Fields: fields_update_contact_flow_name,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateContactFlowNameInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_contact_flow_name, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateContactFlowName(ctx, input)
			},
		},
		"update-contact-routing-data": {
			Name:   "update-contact-routing-data",
			Fields: fields_update_contact_routing_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateContactRoutingDataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_contact_routing_data, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateContactRoutingData(ctx, input)
			},
		},
		"update-contact-schedule": {
			Name:   "update-contact-schedule",
			Fields: fields_update_contact_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateContactScheduleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_contact_schedule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateContactSchedule(ctx, input)
			},
		},
		"update-data-table-attribute": {
			Name:   "update-data-table-attribute",
			Fields: fields_update_data_table_attribute,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDataTableAttributeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_data_table_attribute, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDataTableAttribute(ctx, input)
			},
		},
		"update-data-table-metadata": {
			Name:   "update-data-table-metadata",
			Fields: fields_update_data_table_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDataTableMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_data_table_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDataTableMetadata(ctx, input)
			},
		},
		"update-data-table-primary-values": {
			Name:   "update-data-table-primary-values",
			Fields: fields_update_data_table_primary_values,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDataTablePrimaryValuesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_data_table_primary_values, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDataTablePrimaryValues(ctx, input)
			},
		},
		"update-email-address-metadata": {
			Name:   "update-email-address-metadata",
			Fields: fields_update_email_address_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEmailAddressMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_email_address_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEmailAddressMetadata(ctx, input)
			},
		},
		"update-evaluation-form": {
			Name:   "update-evaluation-form",
			Fields: fields_update_evaluation_form,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEvaluationFormInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_evaluation_form, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEvaluationForm(ctx, input)
			},
		},
		"update-hours-of-operation": {
			Name:   "update-hours-of-operation",
			Fields: fields_update_hours_of_operation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateHoursOfOperationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_hours_of_operation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateHoursOfOperation(ctx, input)
			},
		},
		"update-hours-of-operation-override": {
			Name:   "update-hours-of-operation-override",
			Fields: fields_update_hours_of_operation_override,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateHoursOfOperationOverrideInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_hours_of_operation_override, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateHoursOfOperationOverride(ctx, input)
			},
		},
		"update-instance-attribute": {
			Name:   "update-instance-attribute",
			Fields: fields_update_instance_attribute,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateInstanceAttributeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_instance_attribute, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateInstanceAttribute(ctx, input)
			},
		},
		"update-instance-storage-config": {
			Name:   "update-instance-storage-config",
			Fields: fields_update_instance_storage_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateInstanceStorageConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_instance_storage_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateInstanceStorageConfig(ctx, input)
			},
		},
		"update-notification-content": {
			Name:   "update-notification-content",
			Fields: fields_update_notification_content,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateNotificationContentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_notification_content, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateNotificationContent(ctx, input)
			},
		},
		"update-participant-authentication": {
			Name:   "update-participant-authentication",
			Fields: fields_update_participant_authentication,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateParticipantAuthenticationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_participant_authentication, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateParticipantAuthentication(ctx, input)
			},
		},
		"update-participant-role-config": {
			Name:   "update-participant-role-config",
			Fields: fields_update_participant_role_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateParticipantRoleConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_participant_role_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateParticipantRoleConfig(ctx, input)
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
		"update-phone-number-metadata": {
			Name:   "update-phone-number-metadata",
			Fields: fields_update_phone_number_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePhoneNumberMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_phone_number_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePhoneNumberMetadata(ctx, input)
			},
		},
		"update-predefined-attribute": {
			Name:   "update-predefined-attribute",
			Fields: fields_update_predefined_attribute,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePredefinedAttributeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_predefined_attribute, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePredefinedAttribute(ctx, input)
			},
		},
		"update-prompt": {
			Name:   "update-prompt",
			Fields: fields_update_prompt,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePromptInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_prompt, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePrompt(ctx, input)
			},
		},
		"update-queue-hours-of-operation": {
			Name:   "update-queue-hours-of-operation",
			Fields: fields_update_queue_hours_of_operation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateQueueHoursOfOperationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_queue_hours_of_operation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateQueueHoursOfOperation(ctx, input)
			},
		},
		"update-queue-max-contacts": {
			Name:   "update-queue-max-contacts",
			Fields: fields_update_queue_max_contacts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateQueueMaxContactsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_queue_max_contacts, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateQueueMaxContacts(ctx, input)
			},
		},
		"update-queue-name": {
			Name:   "update-queue-name",
			Fields: fields_update_queue_name,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateQueueNameInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_queue_name, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateQueueName(ctx, input)
			},
		},
		"update-queue-outbound-caller-config": {
			Name:   "update-queue-outbound-caller-config",
			Fields: fields_update_queue_outbound_caller_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateQueueOutboundCallerConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_queue_outbound_caller_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateQueueOutboundCallerConfig(ctx, input)
			},
		},
		"update-queue-outbound-email-config": {
			Name:   "update-queue-outbound-email-config",
			Fields: fields_update_queue_outbound_email_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateQueueOutboundEmailConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_queue_outbound_email_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateQueueOutboundEmailConfig(ctx, input)
			},
		},
		"update-queue-status": {
			Name:   "update-queue-status",
			Fields: fields_update_queue_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateQueueStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_queue_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateQueueStatus(ctx, input)
			},
		},
		"update-quick-connect-config": {
			Name:   "update-quick-connect-config",
			Fields: fields_update_quick_connect_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateQuickConnectConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_quick_connect_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateQuickConnectConfig(ctx, input)
			},
		},
		"update-quick-connect-name": {
			Name:   "update-quick-connect-name",
			Fields: fields_update_quick_connect_name,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateQuickConnectNameInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_quick_connect_name, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateQuickConnectName(ctx, input)
			},
		},
		"update-routing-profile-agent-availability-timer": {
			Name:   "update-routing-profile-agent-availability-timer",
			Fields: fields_update_routing_profile_agent_availability_timer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRoutingProfileAgentAvailabilityTimerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_routing_profile_agent_availability_timer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRoutingProfileAgentAvailabilityTimer(ctx, input)
			},
		},
		"update-routing-profile-concurrency": {
			Name:   "update-routing-profile-concurrency",
			Fields: fields_update_routing_profile_concurrency,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRoutingProfileConcurrencyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_routing_profile_concurrency, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRoutingProfileConcurrency(ctx, input)
			},
		},
		"update-routing-profile-default-outbound-queue": {
			Name:   "update-routing-profile-default-outbound-queue",
			Fields: fields_update_routing_profile_default_outbound_queue,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRoutingProfileDefaultOutboundQueueInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_routing_profile_default_outbound_queue, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRoutingProfileDefaultOutboundQueue(ctx, input)
			},
		},
		"update-routing-profile-name": {
			Name:   "update-routing-profile-name",
			Fields: fields_update_routing_profile_name,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRoutingProfileNameInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_routing_profile_name, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRoutingProfileName(ctx, input)
			},
		},
		"update-routing-profile-queues": {
			Name:   "update-routing-profile-queues",
			Fields: fields_update_routing_profile_queues,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRoutingProfileQueuesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_routing_profile_queues, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRoutingProfileQueues(ctx, input)
			},
		},
		"update-rule": {
			Name:   "update-rule",
			Fields: fields_update_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRule(ctx, input)
			},
		},
		"update-security-profile": {
			Name:   "update-security-profile",
			Fields: fields_update_security_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSecurityProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_security_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSecurityProfile(ctx, input)
			},
		},
		"update-task-template": {
			Name:   "update-task-template",
			Fields: fields_update_task_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTaskTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_task_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTaskTemplate(ctx, input)
			},
		},
		"update-test-case": {
			Name:   "update-test-case",
			Fields: fields_update_test_case,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTestCaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_test_case, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTestCase(ctx, input)
			},
		},
		"update-traffic-distribution": {
			Name:   "update-traffic-distribution",
			Fields: fields_update_traffic_distribution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTrafficDistributionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_traffic_distribution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTrafficDistribution(ctx, input)
			},
		},
		"update-user-config": {
			Name:   "update-user-config",
			Fields: fields_update_user_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateUserConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_user_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateUserConfig(ctx, input)
			},
		},
		"update-user-hierarchy": {
			Name:   "update-user-hierarchy",
			Fields: fields_update_user_hierarchy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateUserHierarchyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_user_hierarchy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateUserHierarchy(ctx, input)
			},
		},
		"update-user-hierarchy-group-name": {
			Name:   "update-user-hierarchy-group-name",
			Fields: fields_update_user_hierarchy_group_name,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateUserHierarchyGroupNameInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_user_hierarchy_group_name, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateUserHierarchyGroupName(ctx, input)
			},
		},
		"update-user-hierarchy-structure": {
			Name:   "update-user-hierarchy-structure",
			Fields: fields_update_user_hierarchy_structure,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateUserHierarchyStructureInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_user_hierarchy_structure, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateUserHierarchyStructure(ctx, input)
			},
		},
		"update-user-identity-info": {
			Name:   "update-user-identity-info",
			Fields: fields_update_user_identity_info,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateUserIdentityInfoInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_user_identity_info, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateUserIdentityInfo(ctx, input)
			},
		},
		"update-user-notification-status": {
			Name:   "update-user-notification-status",
			Fields: fields_update_user_notification_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateUserNotificationStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_user_notification_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateUserNotificationStatus(ctx, input)
			},
		},
		"update-user-phone-config": {
			Name:   "update-user-phone-config",
			Fields: fields_update_user_phone_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateUserPhoneConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_user_phone_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateUserPhoneConfig(ctx, input)
			},
		},
		"update-user-proficiencies": {
			Name:   "update-user-proficiencies",
			Fields: fields_update_user_proficiencies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateUserProficienciesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_user_proficiencies, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateUserProficiencies(ctx, input)
			},
		},
		"update-user-routing-profile": {
			Name:   "update-user-routing-profile",
			Fields: fields_update_user_routing_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateUserRoutingProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_user_routing_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateUserRoutingProfile(ctx, input)
			},
		},
		"update-user-security-profiles": {
			Name:   "update-user-security-profiles",
			Fields: fields_update_user_security_profiles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateUserSecurityProfilesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_user_security_profiles, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateUserSecurityProfiles(ctx, input)
			},
		},
		"update-view-content": {
			Name:   "update-view-content",
			Fields: fields_update_view_content,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateViewContentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_view_content, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateViewContent(ctx, input)
			},
		},
		"update-view-metadata": {
			Name:   "update-view-metadata",
			Fields: fields_update_view_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateViewMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_view_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateViewMetadata(ctx, input)
			},
		},
		"update-workspace-metadata": {
			Name:   "update-workspace-metadata",
			Fields: fields_update_workspace_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateWorkspaceMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_workspace_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateWorkspaceMetadata(ctx, input)
			},
		},
		"update-workspace-page": {
			Name:   "update-workspace-page",
			Fields: fields_update_workspace_page,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateWorkspacePageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_workspace_page, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateWorkspacePage(ctx, input)
			},
		},
		"update-workspace-theme": {
			Name:   "update-workspace-theme",
			Fields: fields_update_workspace_theme,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateWorkspaceThemeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_workspace_theme, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateWorkspaceTheme(ctx, input)
			},
		},
		"update-workspace-visibility": {
			Name:   "update-workspace-visibility",
			Fields: fields_update_workspace_visibility,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateWorkspaceVisibilityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_workspace_visibility, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateWorkspaceVisibility(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("connect", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
