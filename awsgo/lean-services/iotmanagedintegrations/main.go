package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/iotmanagedintegrations"
)

var fields_create_account_association = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ConnectorDestinationId", Flag: "connector-destination-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "GeneralAuthorization", Flag: "general-authorization", Type: "*types.GeneralAuthorizationName", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_cloud_connector = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EndpointConfig", Flag: "endpoint-config", Type: "*types.EndpointConfig", Required: true},
	{Name: "EndpointType", Flag: "endpoint-type", Type: "types.EndpointType", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_create_connector_destination = []leanruntime.Field{
	{Name: "AuthConfig", Flag: "auth-config", Type: "*types.AuthConfig", Required: true},
	{Name: "AuthType", Flag: "auth-type", Type: "types.AuthType", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CloudConnectorId", Flag: "cloud-connector-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "SecretsManager", Flag: "secrets-manager", Type: "*types.SecretsManager", Required: false},
}

var fields_create_credential_locker = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_destination = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DeliveryDestinationArn", Flag: "delivery-destination-arn", Type: "*string", Required: true},
	{Name: "DeliveryDestinationType", Flag: "delivery-destination-type", Type: "types.DeliveryDestinationType", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_event_log_configuration = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "EventLogLevel", Flag: "event-log-level", Type: "types.LogLevel", Required: true},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: false},
	{Name: "ResourceType", Flag: "resource-type", Type: "*string", Required: true},
}

var fields_create_managed_thing = []leanruntime.Field{
	{Name: "AuthenticationMaterial", Flag: "authentication-material", Type: "*string", Required: true},
	{Name: "AuthenticationMaterialType", Flag: "authentication-material-type", Type: "types.AuthMaterialType", Required: true},
	{Name: "Brand", Flag: "brand", Type: "*string", Required: false},
	{Name: "Capabilities", Flag: "capabilities", Type: "*string", Required: false},
	{Name: "CapabilityReport", Flag: "capability-report", Type: "*types.CapabilityReport", Required: false},
	{Name: "CapabilitySchemas", Flag: "capability-schemas", Type: "[]types.CapabilitySchemaItem", Required: false},
	{Name: "Classification", Flag: "classification", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CredentialLockerId", Flag: "credential-locker-id", Type: "*string", Required: false},
	{Name: "MetaData", Flag: "meta-data", Type: "map[string]string", Required: false},
	{Name: "Model", Flag: "model", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Owner", Flag: "owner", Type: "*string", Required: false},
	{Name: "Role", Flag: "role", Type: "types.Role", Required: true},
	{Name: "SerialNumber", Flag: "serial-number", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "WiFiSimpleSetupConfiguration", Flag: "wi-fi-simple-setup-configuration", Type: "*types.WiFiSimpleSetupConfiguration", Required: false},
}

var fields_create_notification_configuration = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DestinationName", Flag: "destination-name", Type: "*string", Required: true},
	{Name: "EventType", Flag: "event-type", Type: "types.EventType", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_ota_task = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "OtaMechanism", Flag: "ota-mechanism", Type: "types.OtaMechanism", Required: false},
	{Name: "OtaSchedulingConfig", Flag: "ota-scheduling-config", Type: "*types.OtaTaskSchedulingConfig", Required: false},
	{Name: "OtaTargetQueryString", Flag: "ota-target-query-string", Type: "*string", Required: false},
	{Name: "OtaTaskExecutionRetryConfig", Flag: "ota-task-execution-retry-config", Type: "*types.OtaTaskExecutionRetryConfig", Required: false},
	{Name: "OtaType", Flag: "ota-type", Type: "types.OtaType", Required: true},
	{Name: "Protocol", Flag: "protocol", Type: "types.OtaProtocol", Required: false},
	{Name: "S3Url", Flag: "s3-url", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Target", Flag: "target", Type: "[]string", Required: false},
	{Name: "TaskConfigurationId", Flag: "task-configuration-id", Type: "*string", Required: false},
}

var fields_create_ota_task_configuration = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "PushConfig", Flag: "push-config", Type: "*types.PushConfig", Required: false},
}

var fields_create_provisioning_profile = []leanruntime.Field{
	{Name: "CaCertificate", Flag: "ca-certificate", Type: "*string", Required: false},
	{Name: "ClaimCertificate", Flag: "claim-certificate", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "ProvisioningType", Flag: "provisioning-type", Type: "types.ProvisioningType", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_account_association = []leanruntime.Field{
	{Name: "AccountAssociationId", Flag: "account-association-id", Type: "*string", Required: true},
}

var fields_delete_cloud_connector = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_delete_connector_destination = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_delete_credential_locker = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_delete_destination = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_event_log_configuration = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_managed_thing = []leanruntime.Field{
	{Name: "Force", Flag: "force", Type: "*bool", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_delete_notification_configuration = []leanruntime.Field{
	{Name: "EventType", Flag: "event-type", Type: "types.EventType", Required: true},
}

var fields_delete_ota_task = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_delete_ota_task_configuration = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_delete_provisioning_profile = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_deregister_account_association = []leanruntime.Field{
	{Name: "AccountAssociationId", Flag: "account-association-id", Type: "*string", Required: true},
	{Name: "ManagedThingId", Flag: "managed-thing-id", Type: "*string", Required: true},
}

var fields_get_account_association = []leanruntime.Field{
	{Name: "AccountAssociationId", Flag: "account-association-id", Type: "*string", Required: true},
}

var fields_get_cloud_connector = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_connector_destination = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_credential_locker = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_custom_endpoint = []leanruntime.Field{}

var fields_get_default_encryption_configuration = []leanruntime.Field{}

var fields_get_destination = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_device_discovery = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_event_log_configuration = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_hub_configuration = []leanruntime.Field{}

var fields_get_managed_thing = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_managed_thing_capabilities = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_managed_thing_certificate = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_managed_thing_connectivity_data = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_managed_thing_meta_data = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_managed_thing_state = []leanruntime.Field{
	{Name: "ManagedThingId", Flag: "managed-thing-id", Type: "*string", Required: true},
}

var fields_get_notification_configuration = []leanruntime.Field{
	{Name: "EventType", Flag: "event-type", Type: "types.EventType", Required: true},
}

var fields_get_ota_task = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_ota_task_configuration = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_provisioning_profile = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_runtime_log_configuration = []leanruntime.Field{
	{Name: "ManagedThingId", Flag: "managed-thing-id", Type: "*string", Required: true},
}

var fields_get_schema_version = []leanruntime.Field{
	{Name: "Format", Flag: "format", Type: "types.SchemaVersionFormat", Required: false},
	{Name: "SchemaVersionedId", Flag: "schema-versioned-id", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.SchemaVersionType", Required: true},
}

var fields_list_account_associations = []leanruntime.Field{
	{Name: "ConnectorDestinationId", Flag: "connector-destination-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_cloud_connectors = []leanruntime.Field{
	{Name: "LambdaArn", Flag: "lambda-arn", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.CloudConnectorType", Required: false},
}

var fields_list_connector_destinations = []leanruntime.Field{
	{Name: "CloudConnectorId", Flag: "cloud-connector-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_credential_lockers = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_destinations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_device_discoveries = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StatusFilter", Flag: "status-filter", Type: "types.DeviceDiscoveryStatus", Required: false},
	{Name: "TypeFilter", Flag: "type-filter", Type: "types.DiscoveryType", Required: false},
}

var fields_list_discovered_devices = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_event_log_configurations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_managed_thing_account_associations = []leanruntime.Field{
	{Name: "AccountAssociationId", Flag: "account-association-id", Type: "*string", Required: false},
	{Name: "ManagedThingId", Flag: "managed-thing-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_managed_thing_schemas = []leanruntime.Field{
	{Name: "CapabilityIdFilter", Flag: "capability-id-filter", Type: "*string", Required: false},
	{Name: "EndpointIdFilter", Flag: "endpoint-id-filter", Type: "*string", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_managed_things = []leanruntime.Field{
	{Name: "ConnectorDestinationIdFilter", Flag: "connector-destination-id-filter", Type: "*string", Required: false},
	{Name: "ConnectorDeviceIdFilter", Flag: "connector-device-id-filter", Type: "*string", Required: false},
	{Name: "ConnectorPolicyIdFilter", Flag: "connector-policy-id-filter", Type: "*string", Required: false},
	{Name: "CredentialLockerFilter", Flag: "credential-locker-filter", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OwnerFilter", Flag: "owner-filter", Type: "*string", Required: false},
	{Name: "ParentControllerIdentifierFilter", Flag: "parent-controller-identifier-filter", Type: "*string", Required: false},
	{Name: "ProvisioningStatusFilter", Flag: "provisioning-status-filter", Type: "types.ProvisioningStatus", Required: false},
	{Name: "RoleFilter", Flag: "role-filter", Type: "types.Role", Required: false},
	{Name: "SerialNumberFilter", Flag: "serial-number-filter", Type: "*string", Required: false},
}

var fields_list_notification_configurations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_ota_task_configurations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_ota_task_executions = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_ota_tasks = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_provisioning_profiles = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_schema_versions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SchemaId", Flag: "schema-id", Type: "*string", Required: false},
	{Name: "SemanticVersion", Flag: "semantic-version", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.SchemaVersionType", Required: true},
	{Name: "Visibility", Flag: "visibility", Type: "types.SchemaVersionVisibility", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_put_default_encryption_configuration = []leanruntime.Field{
	{Name: "EncryptionType", Flag: "encryption-type", Type: "types.EncryptionType", Required: true},
	{Name: "KmsKeyArn", Flag: "kms-key-arn", Type: "*string", Required: false},
}

var fields_put_hub_configuration = []leanruntime.Field{
	{Name: "HubTokenTimerExpirySettingInSeconds", Flag: "hub-token-timer-expiry-setting-in-seconds", Type: "*int64", Required: true},
}

var fields_put_runtime_log_configuration = []leanruntime.Field{
	{Name: "ManagedThingId", Flag: "managed-thing-id", Type: "*string", Required: true},
	{Name: "RuntimeLogConfigurations", Flag: "runtime-log-configurations", Type: "*types.RuntimeLogConfigurations", Required: true},
}

var fields_register_account_association = []leanruntime.Field{
	{Name: "AccountAssociationId", Flag: "account-association-id", Type: "*string", Required: true},
	{Name: "DeviceDiscoveryId", Flag: "device-discovery-id", Type: "*string", Required: true},
	{Name: "ManagedThingId", Flag: "managed-thing-id", Type: "*string", Required: true},
}

var fields_register_custom_endpoint = []leanruntime.Field{}

var fields_reset_runtime_log_configuration = []leanruntime.Field{
	{Name: "ManagedThingId", Flag: "managed-thing-id", Type: "*string", Required: true},
}

var fields_send_connector_event = []leanruntime.Field{
	{Name: "ConnectorDeviceId", Flag: "connector-device-id", Type: "*string", Required: false},
	{Name: "ConnectorId", Flag: "connector-id", Type: "*string", Required: true},
	{Name: "DeviceDiscoveryId", Flag: "device-discovery-id", Type: "*string", Required: false},
	{Name: "Devices", Flag: "devices", Type: "[]types.Device", Required: false},
	{Name: "MatterEndpoint", Flag: "matter-endpoint", Type: "*types.MatterEndpoint", Required: false},
	{Name: "Message", Flag: "message", Type: "*string", Required: false},
	{Name: "Operation", Flag: "operation", Type: "types.ConnectorEventOperation", Required: true},
	{Name: "OperationVersion", Flag: "operation-version", Type: "*string", Required: false},
	{Name: "StatusCode", Flag: "status-code", Type: "*int32", Required: false},
	{Name: "TraceId", Flag: "trace-id", Type: "*string", Required: false},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: false},
}

var fields_send_managed_thing_command = []leanruntime.Field{
	{Name: "AccountAssociationId", Flag: "account-association-id", Type: "*string", Required: false},
	{Name: "ConnectorAssociationId", Flag: "connector-association-id", Type: "*string", Required: false},
	{Name: "Endpoints", Flag: "endpoints", Type: "[]types.CommandEndpoint", Required: true},
	{Name: "ManagedThingId", Flag: "managed-thing-id", Type: "*string", Required: true},
}

var fields_start_account_association_refresh = []leanruntime.Field{
	{Name: "AccountAssociationId", Flag: "account-association-id", Type: "*string", Required: true},
}

var fields_start_device_discovery = []leanruntime.Field{
	{Name: "AccountAssociationId", Flag: "account-association-id", Type: "*string", Required: false},
	{Name: "AuthenticationMaterial", Flag: "authentication-material", Type: "*string", Required: false},
	{Name: "AuthenticationMaterialType", Flag: "authentication-material-type", Type: "types.DiscoveryAuthMaterialType", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ConnectorAssociationIdentifier", Flag: "connector-association-identifier", Type: "*string", Required: false},
	{Name: "ConnectorDeviceIdList", Flag: "connector-device-id-list", Type: "[]string", Required: false},
	{Name: "ControllerIdentifier", Flag: "controller-identifier", Type: "*string", Required: false},
	{Name: "CustomProtocolDetail", Flag: "custom-protocol-detail", Type: "map[string]string", Required: false},
	{Name: "DiscoveryType", Flag: "discovery-type", Type: "types.DiscoveryType", Required: true},
	{Name: "EndDeviceIdentifier", Flag: "end-device-identifier", Type: "*string", Required: false},
	{Name: "Protocol", Flag: "protocol", Type: "types.ProtocolType", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_account_association = []leanruntime.Field{
	{Name: "AccountAssociationId", Flag: "account-association-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_cloud_connector = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_connector_destination = []leanruntime.Field{
	{Name: "AuthConfig", Flag: "auth-config", Type: "*types.AuthConfigUpdate", Required: false},
	{Name: "AuthType", Flag: "auth-type", Type: "types.AuthType", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "SecretsManager", Flag: "secrets-manager", Type: "*types.SecretsManager", Required: false},
}

var fields_update_destination = []leanruntime.Field{
	{Name: "DeliveryDestinationArn", Flag: "delivery-destination-arn", Type: "*string", Required: false},
	{Name: "DeliveryDestinationType", Flag: "delivery-destination-type", Type: "types.DeliveryDestinationType", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
}

var fields_update_event_log_configuration = []leanruntime.Field{
	{Name: "EventLogLevel", Flag: "event-log-level", Type: "types.LogLevel", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_update_managed_thing = []leanruntime.Field{
	{Name: "Brand", Flag: "brand", Type: "*string", Required: false},
	{Name: "Capabilities", Flag: "capabilities", Type: "*string", Required: false},
	{Name: "CapabilityReport", Flag: "capability-report", Type: "*types.CapabilityReport", Required: false},
	{Name: "CapabilitySchemas", Flag: "capability-schemas", Type: "[]types.CapabilitySchemaItem", Required: false},
	{Name: "Classification", Flag: "classification", Type: "*string", Required: false},
	{Name: "CredentialLockerId", Flag: "credential-locker-id", Type: "*string", Required: false},
	{Name: "HubNetworkMode", Flag: "hub-network-mode", Type: "types.HubNetworkMode", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "MetaData", Flag: "meta-data", Type: "map[string]string", Required: false},
	{Name: "Model", Flag: "model", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Owner", Flag: "owner", Type: "*string", Required: false},
	{Name: "SerialNumber", Flag: "serial-number", Type: "*string", Required: false},
	{Name: "WiFiSimpleSetupConfiguration", Flag: "wi-fi-simple-setup-configuration", Type: "*types.WiFiSimpleSetupConfiguration", Required: false},
}

var fields_update_notification_configuration = []leanruntime.Field{
	{Name: "DestinationName", Flag: "destination-name", Type: "*string", Required: true},
	{Name: "EventType", Flag: "event-type", Type: "types.EventType", Required: true},
}

var fields_update_ota_task = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "TaskConfigurationId", Flag: "task-configuration-id", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-account-association": {
			Name:   "create-account-association",
			Fields: fields_create_account_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAccountAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_account_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAccountAssociation(ctx, input)
			},
		},
		"create-cloud-connector": {
			Name:   "create-cloud-connector",
			Fields: fields_create_cloud_connector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCloudConnectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_cloud_connector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCloudConnector(ctx, input)
			},
		},
		"create-connector-destination": {
			Name:   "create-connector-destination",
			Fields: fields_create_connector_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateConnectorDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_connector_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateConnectorDestination(ctx, input)
			},
		},
		"create-credential-locker": {
			Name:   "create-credential-locker",
			Fields: fields_create_credential_locker,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCredentialLockerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_credential_locker, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCredentialLocker(ctx, input)
			},
		},
		"create-destination": {
			Name:   "create-destination",
			Fields: fields_create_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDestination(ctx, input)
			},
		},
		"create-event-log-configuration": {
			Name:   "create-event-log-configuration",
			Fields: fields_create_event_log_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEventLogConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_event_log_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEventLogConfiguration(ctx, input)
			},
		},
		"create-managed-thing": {
			Name:   "create-managed-thing",
			Fields: fields_create_managed_thing,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateManagedThingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_managed_thing, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateManagedThing(ctx, input)
			},
		},
		"create-notification-configuration": {
			Name:   "create-notification-configuration",
			Fields: fields_create_notification_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateNotificationConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_notification_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateNotificationConfiguration(ctx, input)
			},
		},
		"create-ota-task": {
			Name:   "create-ota-task",
			Fields: fields_create_ota_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateOtaTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_ota_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateOtaTask(ctx, input)
			},
		},
		"create-ota-task-configuration": {
			Name:   "create-ota-task-configuration",
			Fields: fields_create_ota_task_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateOtaTaskConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_ota_task_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateOtaTaskConfiguration(ctx, input)
			},
		},
		"create-provisioning-profile": {
			Name:   "create-provisioning-profile",
			Fields: fields_create_provisioning_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateProvisioningProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_provisioning_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateProvisioningProfile(ctx, input)
			},
		},
		"delete-account-association": {
			Name:   "delete-account-association",
			Fields: fields_delete_account_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAccountAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_account_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAccountAssociation(ctx, input)
			},
		},
		"delete-cloud-connector": {
			Name:   "delete-cloud-connector",
			Fields: fields_delete_cloud_connector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCloudConnectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_cloud_connector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCloudConnector(ctx, input)
			},
		},
		"delete-connector-destination": {
			Name:   "delete-connector-destination",
			Fields: fields_delete_connector_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConnectorDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_connector_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConnectorDestination(ctx, input)
			},
		},
		"delete-credential-locker": {
			Name:   "delete-credential-locker",
			Fields: fields_delete_credential_locker,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCredentialLockerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_credential_locker, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCredentialLocker(ctx, input)
			},
		},
		"delete-destination": {
			Name:   "delete-destination",
			Fields: fields_delete_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDestination(ctx, input)
			},
		},
		"delete-event-log-configuration": {
			Name:   "delete-event-log-configuration",
			Fields: fields_delete_event_log_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEventLogConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_event_log_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEventLogConfiguration(ctx, input)
			},
		},
		"delete-managed-thing": {
			Name:   "delete-managed-thing",
			Fields: fields_delete_managed_thing,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteManagedThingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_managed_thing, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteManagedThing(ctx, input)
			},
		},
		"delete-notification-configuration": {
			Name:   "delete-notification-configuration",
			Fields: fields_delete_notification_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteNotificationConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_notification_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteNotificationConfiguration(ctx, input)
			},
		},
		"delete-ota-task": {
			Name:   "delete-ota-task",
			Fields: fields_delete_ota_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteOtaTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_ota_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteOtaTask(ctx, input)
			},
		},
		"delete-ota-task-configuration": {
			Name:   "delete-ota-task-configuration",
			Fields: fields_delete_ota_task_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteOtaTaskConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_ota_task_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteOtaTaskConfiguration(ctx, input)
			},
		},
		"delete-provisioning-profile": {
			Name:   "delete-provisioning-profile",
			Fields: fields_delete_provisioning_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteProvisioningProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_provisioning_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteProvisioningProfile(ctx, input)
			},
		},
		"deregister-account-association": {
			Name:   "deregister-account-association",
			Fields: fields_deregister_account_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterAccountAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_account_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterAccountAssociation(ctx, input)
			},
		},
		"get-account-association": {
			Name:   "get-account-association",
			Fields: fields_get_account_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccountAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_account_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccountAssociation(ctx, input)
			},
		},
		"get-cloud-connector": {
			Name:   "get-cloud-connector",
			Fields: fields_get_cloud_connector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCloudConnectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_cloud_connector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCloudConnector(ctx, input)
			},
		},
		"get-connector-destination": {
			Name:   "get-connector-destination",
			Fields: fields_get_connector_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConnectorDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_connector_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConnectorDestination(ctx, input)
			},
		},
		"get-credential-locker": {
			Name:   "get-credential-locker",
			Fields: fields_get_credential_locker,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCredentialLockerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_credential_locker, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCredentialLocker(ctx, input)
			},
		},
		"get-custom-endpoint": {
			Name:   "get-custom-endpoint",
			Fields: fields_get_custom_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCustomEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_custom_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCustomEndpoint(ctx, input)
			},
		},
		"get-default-encryption-configuration": {
			Name:   "get-default-encryption-configuration",
			Fields: fields_get_default_encryption_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDefaultEncryptionConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_default_encryption_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDefaultEncryptionConfiguration(ctx, input)
			},
		},
		"get-destination": {
			Name:   "get-destination",
			Fields: fields_get_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDestination(ctx, input)
			},
		},
		"get-device-discovery": {
			Name:   "get-device-discovery",
			Fields: fields_get_device_discovery,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDeviceDiscoveryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_device_discovery, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDeviceDiscovery(ctx, input)
			},
		},
		"get-event-log-configuration": {
			Name:   "get-event-log-configuration",
			Fields: fields_get_event_log_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEventLogConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_event_log_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEventLogConfiguration(ctx, input)
			},
		},
		"get-hub-configuration": {
			Name:   "get-hub-configuration",
			Fields: fields_get_hub_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetHubConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_hub_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetHubConfiguration(ctx, input)
			},
		},
		"get-managed-thing": {
			Name:   "get-managed-thing",
			Fields: fields_get_managed_thing,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetManagedThingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_managed_thing, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetManagedThing(ctx, input)
			},
		},
		"get-managed-thing-capabilities": {
			Name:   "get-managed-thing-capabilities",
			Fields: fields_get_managed_thing_capabilities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetManagedThingCapabilitiesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_managed_thing_capabilities, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetManagedThingCapabilities(ctx, input)
			},
		},
		"get-managed-thing-certificate": {
			Name:   "get-managed-thing-certificate",
			Fields: fields_get_managed_thing_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetManagedThingCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_managed_thing_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetManagedThingCertificate(ctx, input)
			},
		},
		"get-managed-thing-connectivity-data": {
			Name:   "get-managed-thing-connectivity-data",
			Fields: fields_get_managed_thing_connectivity_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetManagedThingConnectivityDataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_managed_thing_connectivity_data, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetManagedThingConnectivityData(ctx, input)
			},
		},
		"get-managed-thing-meta-data": {
			Name:   "get-managed-thing-meta-data",
			Fields: fields_get_managed_thing_meta_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetManagedThingMetaDataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_managed_thing_meta_data, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetManagedThingMetaData(ctx, input)
			},
		},
		"get-managed-thing-state": {
			Name:   "get-managed-thing-state",
			Fields: fields_get_managed_thing_state,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetManagedThingStateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_managed_thing_state, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetManagedThingState(ctx, input)
			},
		},
		"get-notification-configuration": {
			Name:   "get-notification-configuration",
			Fields: fields_get_notification_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetNotificationConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_notification_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetNotificationConfiguration(ctx, input)
			},
		},
		"get-ota-task": {
			Name:   "get-ota-task",
			Fields: fields_get_ota_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOtaTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_ota_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetOtaTask(ctx, input)
			},
		},
		"get-ota-task-configuration": {
			Name:   "get-ota-task-configuration",
			Fields: fields_get_ota_task_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOtaTaskConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_ota_task_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetOtaTaskConfiguration(ctx, input)
			},
		},
		"get-provisioning-profile": {
			Name:   "get-provisioning-profile",
			Fields: fields_get_provisioning_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetProvisioningProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_provisioning_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetProvisioningProfile(ctx, input)
			},
		},
		"get-runtime-log-configuration": {
			Name:   "get-runtime-log-configuration",
			Fields: fields_get_runtime_log_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRuntimeLogConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_runtime_log_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRuntimeLogConfiguration(ctx, input)
			},
		},
		"get-schema-version": {
			Name:   "get-schema-version",
			Fields: fields_get_schema_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSchemaVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_schema_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSchemaVersion(ctx, input)
			},
		},
		"list-account-associations": {
			Name:   "list-account-associations",
			Fields: fields_list_account_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAccountAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_account_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAccountAssociations(ctx, input)
				}
				var results []*svc.ListAccountAssociationsOutput
				p := svc.NewListAccountAssociationsPaginator(client, input)
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
		"list-cloud-connectors": {
			Name:   "list-cloud-connectors",
			Fields: fields_list_cloud_connectors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCloudConnectorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_cloud_connectors, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCloudConnectors(ctx, input)
				}
				var results []*svc.ListCloudConnectorsOutput
				p := svc.NewListCloudConnectorsPaginator(client, input)
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
		"list-connector-destinations": {
			Name:   "list-connector-destinations",
			Fields: fields_list_connector_destinations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConnectorDestinationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_connector_destinations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListConnectorDestinations(ctx, input)
				}
				var results []*svc.ListConnectorDestinationsOutput
				p := svc.NewListConnectorDestinationsPaginator(client, input)
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
		"list-credential-lockers": {
			Name:   "list-credential-lockers",
			Fields: fields_list_credential_lockers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCredentialLockersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_credential_lockers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCredentialLockers(ctx, input)
				}
				var results []*svc.ListCredentialLockersOutput
				p := svc.NewListCredentialLockersPaginator(client, input)
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
		"list-destinations": {
			Name:   "list-destinations",
			Fields: fields_list_destinations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDestinationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_destinations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDestinations(ctx, input)
				}
				var results []*svc.ListDestinationsOutput
				p := svc.NewListDestinationsPaginator(client, input)
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
		"list-device-discoveries": {
			Name:   "list-device-discoveries",
			Fields: fields_list_device_discoveries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDeviceDiscoveriesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_device_discoveries, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDeviceDiscoveries(ctx, input)
				}
				var results []*svc.ListDeviceDiscoveriesOutput
				p := svc.NewListDeviceDiscoveriesPaginator(client, input)
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
		"list-discovered-devices": {
			Name:   "list-discovered-devices",
			Fields: fields_list_discovered_devices,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDiscoveredDevicesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_discovered_devices, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDiscoveredDevices(ctx, input)
				}
				var results []*svc.ListDiscoveredDevicesOutput
				p := svc.NewListDiscoveredDevicesPaginator(client, input)
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
		"list-event-log-configurations": {
			Name:   "list-event-log-configurations",
			Fields: fields_list_event_log_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEventLogConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_event_log_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEventLogConfigurations(ctx, input)
				}
				var results []*svc.ListEventLogConfigurationsOutput
				p := svc.NewListEventLogConfigurationsPaginator(client, input)
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
		"list-managed-thing-account-associations": {
			Name:   "list-managed-thing-account-associations",
			Fields: fields_list_managed_thing_account_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListManagedThingAccountAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_managed_thing_account_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListManagedThingAccountAssociations(ctx, input)
				}
				var results []*svc.ListManagedThingAccountAssociationsOutput
				p := svc.NewListManagedThingAccountAssociationsPaginator(client, input)
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
		"list-managed-thing-schemas": {
			Name:   "list-managed-thing-schemas",
			Fields: fields_list_managed_thing_schemas,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListManagedThingSchemasInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_managed_thing_schemas, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListManagedThingSchemas(ctx, input)
				}
				var results []*svc.ListManagedThingSchemasOutput
				p := svc.NewListManagedThingSchemasPaginator(client, input)
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
		"list-managed-things": {
			Name:   "list-managed-things",
			Fields: fields_list_managed_things,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListManagedThingsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_managed_things, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListManagedThings(ctx, input)
				}
				var results []*svc.ListManagedThingsOutput
				p := svc.NewListManagedThingsPaginator(client, input)
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
		"list-notification-configurations": {
			Name:   "list-notification-configurations",
			Fields: fields_list_notification_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListNotificationConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_notification_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListNotificationConfigurations(ctx, input)
				}
				var results []*svc.ListNotificationConfigurationsOutput
				p := svc.NewListNotificationConfigurationsPaginator(client, input)
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
		"list-ota-task-configurations": {
			Name:   "list-ota-task-configurations",
			Fields: fields_list_ota_task_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOtaTaskConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_ota_task_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListOtaTaskConfigurations(ctx, input)
				}
				var results []*svc.ListOtaTaskConfigurationsOutput
				p := svc.NewListOtaTaskConfigurationsPaginator(client, input)
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
		"list-ota-task-executions": {
			Name:   "list-ota-task-executions",
			Fields: fields_list_ota_task_executions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOtaTaskExecutionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_ota_task_executions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListOtaTaskExecutions(ctx, input)
				}
				var results []*svc.ListOtaTaskExecutionsOutput
				p := svc.NewListOtaTaskExecutionsPaginator(client, input)
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
		"list-ota-tasks": {
			Name:   "list-ota-tasks",
			Fields: fields_list_ota_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOtaTasksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_ota_tasks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListOtaTasks(ctx, input)
				}
				var results []*svc.ListOtaTasksOutput
				p := svc.NewListOtaTasksPaginator(client, input)
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
		"list-provisioning-profiles": {
			Name:   "list-provisioning-profiles",
			Fields: fields_list_provisioning_profiles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProvisioningProfilesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_provisioning_profiles, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProvisioningProfiles(ctx, input)
				}
				var results []*svc.ListProvisioningProfilesOutput
				p := svc.NewListProvisioningProfilesPaginator(client, input)
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
		"list-schema-versions": {
			Name:   "list-schema-versions",
			Fields: fields_list_schema_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSchemaVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_schema_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSchemaVersions(ctx, input)
				}
				var results []*svc.ListSchemaVersionsOutput
				p := svc.NewListSchemaVersionsPaginator(client, input)
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
		"put-default-encryption-configuration": {
			Name:   "put-default-encryption-configuration",
			Fields: fields_put_default_encryption_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutDefaultEncryptionConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_default_encryption_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutDefaultEncryptionConfiguration(ctx, input)
			},
		},
		"put-hub-configuration": {
			Name:   "put-hub-configuration",
			Fields: fields_put_hub_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutHubConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_hub_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutHubConfiguration(ctx, input)
			},
		},
		"put-runtime-log-configuration": {
			Name:   "put-runtime-log-configuration",
			Fields: fields_put_runtime_log_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutRuntimeLogConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_runtime_log_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutRuntimeLogConfiguration(ctx, input)
			},
		},
		"register-account-association": {
			Name:   "register-account-association",
			Fields: fields_register_account_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterAccountAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_account_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterAccountAssociation(ctx, input)
			},
		},
		"register-custom-endpoint": {
			Name:   "register-custom-endpoint",
			Fields: fields_register_custom_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterCustomEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_custom_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterCustomEndpoint(ctx, input)
			},
		},
		"reset-runtime-log-configuration": {
			Name:   "reset-runtime-log-configuration",
			Fields: fields_reset_runtime_log_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResetRuntimeLogConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reset_runtime_log_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResetRuntimeLogConfiguration(ctx, input)
			},
		},
		"send-connector-event": {
			Name:   "send-connector-event",
			Fields: fields_send_connector_event,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendConnectorEventInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_connector_event, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendConnectorEvent(ctx, input)
			},
		},
		"send-managed-thing-command": {
			Name:   "send-managed-thing-command",
			Fields: fields_send_managed_thing_command,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendManagedThingCommandInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_managed_thing_command, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendManagedThingCommand(ctx, input)
			},
		},
		"start-account-association-refresh": {
			Name:   "start-account-association-refresh",
			Fields: fields_start_account_association_refresh,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartAccountAssociationRefreshInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_account_association_refresh, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartAccountAssociationRefresh(ctx, input)
			},
		},
		"start-device-discovery": {
			Name:   "start-device-discovery",
			Fields: fields_start_device_discovery,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartDeviceDiscoveryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_device_discovery, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartDeviceDiscovery(ctx, input)
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
		"update-account-association": {
			Name:   "update-account-association",
			Fields: fields_update_account_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAccountAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_account_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAccountAssociation(ctx, input)
			},
		},
		"update-cloud-connector": {
			Name:   "update-cloud-connector",
			Fields: fields_update_cloud_connector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCloudConnectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_cloud_connector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCloudConnector(ctx, input)
			},
		},
		"update-connector-destination": {
			Name:   "update-connector-destination",
			Fields: fields_update_connector_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateConnectorDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_connector_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateConnectorDestination(ctx, input)
			},
		},
		"update-destination": {
			Name:   "update-destination",
			Fields: fields_update_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDestination(ctx, input)
			},
		},
		"update-event-log-configuration": {
			Name:   "update-event-log-configuration",
			Fields: fields_update_event_log_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEventLogConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_event_log_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEventLogConfiguration(ctx, input)
			},
		},
		"update-managed-thing": {
			Name:   "update-managed-thing",
			Fields: fields_update_managed_thing,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateManagedThingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_managed_thing, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateManagedThing(ctx, input)
			},
		},
		"update-notification-configuration": {
			Name:   "update-notification-configuration",
			Fields: fields_update_notification_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateNotificationConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_notification_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateNotificationConfiguration(ctx, input)
			},
		},
		"update-ota-task": {
			Name:   "update-ota-task",
			Fields: fields_update_ota_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateOtaTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_ota_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateOtaTask(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("iotmanagedintegrations", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
