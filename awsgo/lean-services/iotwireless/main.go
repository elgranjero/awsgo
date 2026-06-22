package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/iotwireless"
)

var fields_associate_aws_account_with_partner_account = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Sidewalk", Flag: "sidewalk", Type: "*types.SidewalkAccountInfo", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_associate_multicast_group_with_fuota_task = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "MulticastGroupId", Flag: "multicast-group-id", Type: "*string", Required: true},
}

var fields_associate_wireless_device_with_fuota_task = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "WirelessDeviceId", Flag: "wireless-device-id", Type: "*string", Required: true},
}

var fields_associate_wireless_device_with_multicast_group = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "WirelessDeviceId", Flag: "wireless-device-id", Type: "*string", Required: true},
}

var fields_associate_wireless_device_with_thing = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "ThingArn", Flag: "thing-arn", Type: "*string", Required: true},
}

var fields_associate_wireless_gateway_with_certificate = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IotCertificateId", Flag: "iot-certificate-id", Type: "*string", Required: true},
}

var fields_associate_wireless_gateway_with_thing = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "ThingArn", Flag: "thing-arn", Type: "*string", Required: true},
}

var fields_cancel_multicast_group_session = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_create_destination = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Expression", Flag: "expression", Type: "*string", Required: true},
	{Name: "ExpressionType", Flag: "expression-type", Type: "types.ExpressionType", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_device_profile = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "LoRaWAN", Flag: "lo-ra-wan", Type: "*types.LoRaWANDeviceProfile", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Sidewalk", Flag: "sidewalk", Type: "*types.SidewalkCreateDeviceProfile", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_fuota_task = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Descriptor", Flag: "descriptor", Type: "*string", Required: false},
	{Name: "FirmwareUpdateImage", Flag: "firmware-update-image", Type: "*string", Required: true},
	{Name: "FirmwareUpdateRole", Flag: "firmware-update-role", Type: "*string", Required: true},
	{Name: "FragmentIntervalMS", Flag: "fragment-interval-ms", Type: "*int32", Required: false},
	{Name: "FragmentSizeBytes", Flag: "fragment-size-bytes", Type: "*int32", Required: false},
	{Name: "LoRaWAN", Flag: "lo-ra-wan", Type: "*types.LoRaWANFuotaTask", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "RedundancyPercent", Flag: "redundancy-percent", Type: "*int32", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_multicast_group = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "LoRaWAN", Flag: "lo-ra-wan", Type: "*types.LoRaWANMulticast", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_network_analyzer_configuration = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "MulticastGroups", Flag: "multicast-groups", Type: "[]string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TraceContent", Flag: "trace-content", Type: "*types.TraceContent", Required: false},
	{Name: "WirelessDevices", Flag: "wireless-devices", Type: "[]string", Required: false},
	{Name: "WirelessGateways", Flag: "wireless-gateways", Type: "[]string", Required: false},
}

var fields_create_service_profile = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "LoRaWAN", Flag: "lo-ra-wan", Type: "*types.LoRaWANServiceProfile", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_wireless_device = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DestinationName", Flag: "destination-name", Type: "*string", Required: true},
	{Name: "LoRaWAN", Flag: "lo-ra-wan", Type: "*types.LoRaWANDevice", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Positioning", Flag: "positioning", Type: "types.PositioningConfigStatus", Required: false},
	{Name: "Sidewalk", Flag: "sidewalk", Type: "*types.SidewalkCreateWirelessDevice", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Type", Flag: "type", Type: "types.WirelessDeviceType", Required: true},
}

var fields_create_wireless_gateway = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "LoRaWAN", Flag: "lo-ra-wan", Type: "*types.LoRaWANGateway", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_wireless_gateway_task = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "WirelessGatewayTaskDefinitionId", Flag: "wireless-gateway-task-definition-id", Type: "*string", Required: true},
}

var fields_create_wireless_gateway_task_definition = []leanruntime.Field{
	{Name: "AutoCreateTasks", Flag: "auto-create-tasks", Type: "bool", Required: true},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Update", Flag: "update", Type: "*types.UpdateWirelessGatewayTaskCreate", Required: false},
}

var fields_delete_destination = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_device_profile = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_fuota_task = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_multicast_group = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_network_analyzer_configuration = []leanruntime.Field{
	{Name: "ConfigurationName", Flag: "configuration-name", Type: "*string", Required: true},
}

var fields_delete_queued_messages = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "MessageId", Flag: "message-id", Type: "*string", Required: true},
	{Name: "WirelessDeviceType", Flag: "wireless-device-type", Type: "types.WirelessDeviceType", Required: false},
}

var fields_delete_service_profile = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_wireless_device = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_wireless_device_import_task = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_wireless_gateway = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_wireless_gateway_task = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_wireless_gateway_task_definition = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_deregister_wireless_device = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "WirelessDeviceType", Flag: "wireless-device-type", Type: "types.WirelessDeviceType", Required: false},
}

var fields_disassociate_aws_account_from_partner_account = []leanruntime.Field{
	{Name: "PartnerAccountId", Flag: "partner-account-id", Type: "*string", Required: true},
	{Name: "PartnerType", Flag: "partner-type", Type: "types.PartnerType", Required: true},
}

var fields_disassociate_multicast_group_from_fuota_task = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "MulticastGroupId", Flag: "multicast-group-id", Type: "*string", Required: true},
}

var fields_disassociate_wireless_device_from_fuota_task = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "WirelessDeviceId", Flag: "wireless-device-id", Type: "*string", Required: true},
}

var fields_disassociate_wireless_device_from_multicast_group = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "WirelessDeviceId", Flag: "wireless-device-id", Type: "*string", Required: true},
}

var fields_disassociate_wireless_device_from_thing = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_disassociate_wireless_gateway_from_certificate = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_disassociate_wireless_gateway_from_thing = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_destination = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_device_profile = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_event_configuration_by_resource_types = []leanruntime.Field{}

var fields_get_fuota_task = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_log_levels_by_resource_types = []leanruntime.Field{}

var fields_get_metric_configuration = []leanruntime.Field{}

var fields_get_metrics = []leanruntime.Field{
	{Name: "SummaryMetricQueries", Flag: "summary-metric-queries", Type: "[]types.SummaryMetricQuery", Required: false},
}

var fields_get_multicast_group = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_multicast_group_session = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_network_analyzer_configuration = []leanruntime.Field{
	{Name: "ConfigurationName", Flag: "configuration-name", Type: "*string", Required: true},
}

var fields_get_partner_account = []leanruntime.Field{
	{Name: "PartnerAccountId", Flag: "partner-account-id", Type: "*string", Required: true},
	{Name: "PartnerType", Flag: "partner-type", Type: "types.PartnerType", Required: true},
}

var fields_get_position = []leanruntime.Field{
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.PositionResourceType", Required: true},
}

var fields_get_position_configuration = []leanruntime.Field{
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.PositionResourceType", Required: true},
}

var fields_get_position_estimate = []leanruntime.Field{
	{Name: "CellTowers", Flag: "cell-towers", Type: "*types.CellTowers", Required: false},
	{Name: "Gnss", Flag: "gnss", Type: "*types.Gnss", Required: false},
	{Name: "Ip", Flag: "ip", Type: "*types.Ip", Required: false},
	{Name: "Timestamp", Flag: "timestamp", Type: "*time.Time", Required: false},
	{Name: "WiFiAccessPoints", Flag: "wi-fi-access-points", Type: "[]types.WiFiAccessPoint", Required: false},
}

var fields_get_resource_event_configuration = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "IdentifierType", Flag: "identifier-type", Type: "types.IdentifierType", Required: true},
	{Name: "PartnerType", Flag: "partner-type", Type: "types.EventNotificationPartnerType", Required: false},
}

var fields_get_resource_log_level = []leanruntime.Field{
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "*string", Required: true},
}

var fields_get_resource_position = []leanruntime.Field{
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.PositionResourceType", Required: true},
}

var fields_get_service_endpoint = []leanruntime.Field{
	{Name: "ServiceType", Flag: "service-type", Type: "types.WirelessGatewayServiceType", Required: false},
}

var fields_get_service_profile = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_wireless_device = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "IdentifierType", Flag: "identifier-type", Type: "types.WirelessDeviceIdType", Required: true},
}

var fields_get_wireless_device_import_task = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_wireless_device_statistics = []leanruntime.Field{
	{Name: "WirelessDeviceId", Flag: "wireless-device-id", Type: "*string", Required: true},
}

var fields_get_wireless_gateway = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "IdentifierType", Flag: "identifier-type", Type: "types.WirelessGatewayIdType", Required: true},
}

var fields_get_wireless_gateway_certificate = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_wireless_gateway_firmware_information = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_wireless_gateway_statistics = []leanruntime.Field{
	{Name: "WirelessGatewayId", Flag: "wireless-gateway-id", Type: "*string", Required: true},
}

var fields_get_wireless_gateway_task = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_wireless_gateway_task_definition = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_list_destinations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_device_profiles = []leanruntime.Field{
	{Name: "DeviceProfileType", Flag: "device-profile-type", Type: "types.DeviceProfileType", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_devices_for_wireless_device_import_task = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.OnboardStatus", Required: false},
}

var fields_list_event_configurations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.EventNotificationResourceType", Required: true},
}

var fields_list_fuota_tasks = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_multicast_groups = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_multicast_groups_by_fuota_task = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_network_analyzer_configurations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_partner_accounts = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_position_configurations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.PositionResourceType", Required: false},
}

var fields_list_queued_messages = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "WirelessDeviceType", Flag: "wireless-device-type", Type: "types.WirelessDeviceType", Required: false},
}

var fields_list_service_profiles = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_wireless_device_import_tasks = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_wireless_devices = []leanruntime.Field{
	{Name: "DestinationName", Flag: "destination-name", Type: "*string", Required: false},
	{Name: "DeviceProfileId", Flag: "device-profile-id", Type: "*string", Required: false},
	{Name: "FuotaTaskId", Flag: "fuota-task-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "MulticastGroupId", Flag: "multicast-group-id", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ServiceProfileId", Flag: "service-profile-id", Type: "*string", Required: false},
	{Name: "WirelessDeviceType", Flag: "wireless-device-type", Type: "types.WirelessDeviceType", Required: false},
}

var fields_list_wireless_gateway_task_definitions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TaskDefinitionType", Flag: "task-definition-type", Type: "types.WirelessGatewayTaskDefinitionType", Required: false},
}

var fields_list_wireless_gateways = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_put_position_configuration = []leanruntime.Field{
	{Name: "Destination", Flag: "destination", Type: "*string", Required: false},
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.PositionResourceType", Required: true},
	{Name: "Solvers", Flag: "solvers", Type: "*types.PositionSolverConfigurations", Required: false},
}

var fields_put_resource_log_level = []leanruntime.Field{
	{Name: "LogLevel", Flag: "log-level", Type: "types.LogLevel", Required: true},
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "*string", Required: true},
}

var fields_reset_all_resource_log_levels = []leanruntime.Field{}

var fields_reset_resource_log_level = []leanruntime.Field{
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "*string", Required: true},
}

var fields_send_data_to_multicast_group = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "PayloadData", Flag: "payload-data", Type: "*string", Required: true},
	{Name: "WirelessMetadata", Flag: "wireless-metadata", Type: "*types.MulticastWirelessMetadata", Required: true},
}

var fields_send_data_to_wireless_device = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "PayloadData", Flag: "payload-data", Type: "*string", Required: true},
	{Name: "TransmitMode", Flag: "transmit-mode", Type: "*int32", Required: true},
	{Name: "WirelessMetadata", Flag: "wireless-metadata", Type: "*types.WirelessMetadata", Required: false},
}

var fields_start_bulk_associate_wireless_device_with_multicast_group = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "QueryString", Flag: "query-string", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_start_bulk_disassociate_wireless_device_from_multicast_group = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "QueryString", Flag: "query-string", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_start_fuota_task = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "LoRaWAN", Flag: "lo-ra-wan", Type: "*types.LoRaWANStartFuotaTask", Required: false},
}

var fields_start_multicast_group_session = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "LoRaWAN", Flag: "lo-ra-wan", Type: "*types.LoRaWANMulticastSession", Required: true},
}

var fields_start_single_wireless_device_import_task = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "DestinationName", Flag: "destination-name", Type: "*string", Required: true},
	{Name: "DeviceName", Flag: "device-name", Type: "*string", Required: false},
	{Name: "Positioning", Flag: "positioning", Type: "types.PositioningConfigStatus", Required: false},
	{Name: "Sidewalk", Flag: "sidewalk", Type: "*types.SidewalkSingleStartImportInfo", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_start_wireless_device_import_task = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "DestinationName", Flag: "destination-name", Type: "*string", Required: true},
	{Name: "Positioning", Flag: "positioning", Type: "types.PositioningConfigStatus", Required: false},
	{Name: "Sidewalk", Flag: "sidewalk", Type: "*types.SidewalkStartImportInfo", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_test_wireless_device = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_destination = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Expression", Flag: "expression", Type: "*string", Required: false},
	{Name: "ExpressionType", Flag: "expression-type", Type: "types.ExpressionType", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
}

var fields_update_event_configuration_by_resource_types = []leanruntime.Field{
	{Name: "ConnectionStatus", Flag: "connection-status", Type: "*types.ConnectionStatusResourceTypeEventConfiguration", Required: false},
	{Name: "DeviceRegistrationState", Flag: "device-registration-state", Type: "*types.DeviceRegistrationStateResourceTypeEventConfiguration", Required: false},
	{Name: "Join", Flag: "join", Type: "*types.JoinResourceTypeEventConfiguration", Required: false},
	{Name: "MessageDeliveryStatus", Flag: "message-delivery-status", Type: "*types.MessageDeliveryStatusResourceTypeEventConfiguration", Required: false},
	{Name: "Proximity", Flag: "proximity", Type: "*types.ProximityResourceTypeEventConfiguration", Required: false},
}

var fields_update_fuota_task = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Descriptor", Flag: "descriptor", Type: "*string", Required: false},
	{Name: "FirmwareUpdateImage", Flag: "firmware-update-image", Type: "*string", Required: false},
	{Name: "FirmwareUpdateRole", Flag: "firmware-update-role", Type: "*string", Required: false},
	{Name: "FragmentIntervalMS", Flag: "fragment-interval-ms", Type: "*int32", Required: false},
	{Name: "FragmentSizeBytes", Flag: "fragment-size-bytes", Type: "*int32", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "LoRaWAN", Flag: "lo-ra-wan", Type: "*types.LoRaWANFuotaTask", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "RedundancyPercent", Flag: "redundancy-percent", Type: "*int32", Required: false},
}

var fields_update_log_levels_by_resource_types = []leanruntime.Field{
	{Name: "DefaultLogLevel", Flag: "default-log-level", Type: "types.LogLevel", Required: false},
	{Name: "FuotaTaskLogOptions", Flag: "fuota-task-log-options", Type: "[]types.FuotaTaskLogOption", Required: false},
	{Name: "WirelessDeviceLogOptions", Flag: "wireless-device-log-options", Type: "[]types.WirelessDeviceLogOption", Required: false},
	{Name: "WirelessGatewayLogOptions", Flag: "wireless-gateway-log-options", Type: "[]types.WirelessGatewayLogOption", Required: false},
}

var fields_update_metric_configuration = []leanruntime.Field{
	{Name: "SummaryMetric", Flag: "summary-metric", Type: "*types.SummaryMetricConfiguration", Required: false},
}

var fields_update_multicast_group = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "LoRaWAN", Flag: "lo-ra-wan", Type: "*types.LoRaWANMulticast", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_network_analyzer_configuration = []leanruntime.Field{
	{Name: "ConfigurationName", Flag: "configuration-name", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "MulticastGroupsToAdd", Flag: "multicast-groups-to-add", Type: "[]string", Required: false},
	{Name: "MulticastGroupsToRemove", Flag: "multicast-groups-to-remove", Type: "[]string", Required: false},
	{Name: "TraceContent", Flag: "trace-content", Type: "*types.TraceContent", Required: false},
	{Name: "WirelessDevicesToAdd", Flag: "wireless-devices-to-add", Type: "[]string", Required: false},
	{Name: "WirelessDevicesToRemove", Flag: "wireless-devices-to-remove", Type: "[]string", Required: false},
	{Name: "WirelessGatewaysToAdd", Flag: "wireless-gateways-to-add", Type: "[]string", Required: false},
	{Name: "WirelessGatewaysToRemove", Flag: "wireless-gateways-to-remove", Type: "[]string", Required: false},
}

var fields_update_partner_account = []leanruntime.Field{
	{Name: "PartnerAccountId", Flag: "partner-account-id", Type: "*string", Required: true},
	{Name: "PartnerType", Flag: "partner-type", Type: "types.PartnerType", Required: true},
	{Name: "Sidewalk", Flag: "sidewalk", Type: "*types.SidewalkUpdateAccount", Required: true},
}

var fields_update_position = []leanruntime.Field{
	{Name: "Position", Flag: "position", Type: "[]float32", Required: true},
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.PositionResourceType", Required: true},
}

var fields_update_resource_event_configuration = []leanruntime.Field{
	{Name: "ConnectionStatus", Flag: "connection-status", Type: "*types.ConnectionStatusEventConfiguration", Required: false},
	{Name: "DeviceRegistrationState", Flag: "device-registration-state", Type: "*types.DeviceRegistrationStateEventConfiguration", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "IdentifierType", Flag: "identifier-type", Type: "types.IdentifierType", Required: true},
	{Name: "Join", Flag: "join", Type: "*types.JoinEventConfiguration", Required: false},
	{Name: "MessageDeliveryStatus", Flag: "message-delivery-status", Type: "*types.MessageDeliveryStatusEventConfiguration", Required: false},
	{Name: "PartnerType", Flag: "partner-type", Type: "types.EventNotificationPartnerType", Required: false},
	{Name: "Proximity", Flag: "proximity", Type: "*types.ProximityEventConfiguration", Required: false},
}

var fields_update_resource_position = []leanruntime.Field{
	{Name: "GeoJsonPayload", Flag: "geo-json-payload", Type: "[]byte", Required: false},
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.PositionResourceType", Required: true},
}

var fields_update_wireless_device = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DestinationName", Flag: "destination-name", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "LoRaWAN", Flag: "lo-ra-wan", Type: "*types.LoRaWANUpdateDevice", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Positioning", Flag: "positioning", Type: "types.PositioningConfigStatus", Required: false},
	{Name: "Sidewalk", Flag: "sidewalk", Type: "*types.SidewalkUpdateWirelessDevice", Required: false},
}

var fields_update_wireless_device_import_task = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "Sidewalk", Flag: "sidewalk", Type: "*types.SidewalkUpdateImportInfo", Required: true},
}

var fields_update_wireless_gateway = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "JoinEuiFilters", Flag: "join-eui-filters", Type: "[][]string", Required: false},
	{Name: "MaxEirp", Flag: "max-eirp", Type: "*float32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "NetIdFilters", Flag: "net-id-filters", Type: "[]string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-aws-account-with-partner-account": {
			Name:   "associate-aws-account-with-partner-account",
			Fields: fields_associate_aws_account_with_partner_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateAwsAccountWithPartnerAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_aws_account_with_partner_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateAwsAccountWithPartnerAccount(ctx, input)
			},
		},
		"associate-multicast-group-with-fuota-task": {
			Name:   "associate-multicast-group-with-fuota-task",
			Fields: fields_associate_multicast_group_with_fuota_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateMulticastGroupWithFuotaTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_multicast_group_with_fuota_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateMulticastGroupWithFuotaTask(ctx, input)
			},
		},
		"associate-wireless-device-with-fuota-task": {
			Name:   "associate-wireless-device-with-fuota-task",
			Fields: fields_associate_wireless_device_with_fuota_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateWirelessDeviceWithFuotaTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_wireless_device_with_fuota_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateWirelessDeviceWithFuotaTask(ctx, input)
			},
		},
		"associate-wireless-device-with-multicast-group": {
			Name:   "associate-wireless-device-with-multicast-group",
			Fields: fields_associate_wireless_device_with_multicast_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateWirelessDeviceWithMulticastGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_wireless_device_with_multicast_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateWirelessDeviceWithMulticastGroup(ctx, input)
			},
		},
		"associate-wireless-device-with-thing": {
			Name:   "associate-wireless-device-with-thing",
			Fields: fields_associate_wireless_device_with_thing,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateWirelessDeviceWithThingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_wireless_device_with_thing, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateWirelessDeviceWithThing(ctx, input)
			},
		},
		"associate-wireless-gateway-with-certificate": {
			Name:   "associate-wireless-gateway-with-certificate",
			Fields: fields_associate_wireless_gateway_with_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateWirelessGatewayWithCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_wireless_gateway_with_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateWirelessGatewayWithCertificate(ctx, input)
			},
		},
		"associate-wireless-gateway-with-thing": {
			Name:   "associate-wireless-gateway-with-thing",
			Fields: fields_associate_wireless_gateway_with_thing,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateWirelessGatewayWithThingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_wireless_gateway_with_thing, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateWirelessGatewayWithThing(ctx, input)
			},
		},
		"cancel-multicast-group-session": {
			Name:   "cancel-multicast-group-session",
			Fields: fields_cancel_multicast_group_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelMulticastGroupSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_multicast_group_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelMulticastGroupSession(ctx, input)
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
		"create-device-profile": {
			Name:   "create-device-profile",
			Fields: fields_create_device_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDeviceProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_device_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDeviceProfile(ctx, input)
			},
		},
		"create-fuota-task": {
			Name:   "create-fuota-task",
			Fields: fields_create_fuota_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFuotaTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_fuota_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFuotaTask(ctx, input)
			},
		},
		"create-multicast-group": {
			Name:   "create-multicast-group",
			Fields: fields_create_multicast_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMulticastGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_multicast_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMulticastGroup(ctx, input)
			},
		},
		"create-network-analyzer-configuration": {
			Name:   "create-network-analyzer-configuration",
			Fields: fields_create_network_analyzer_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateNetworkAnalyzerConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_network_analyzer_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateNetworkAnalyzerConfiguration(ctx, input)
			},
		},
		"create-service-profile": {
			Name:   "create-service-profile",
			Fields: fields_create_service_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateServiceProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_service_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateServiceProfile(ctx, input)
			},
		},
		"create-wireless-device": {
			Name:   "create-wireless-device",
			Fields: fields_create_wireless_device,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWirelessDeviceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_wireless_device, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWirelessDevice(ctx, input)
			},
		},
		"create-wireless-gateway": {
			Name:   "create-wireless-gateway",
			Fields: fields_create_wireless_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWirelessGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_wireless_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWirelessGateway(ctx, input)
			},
		},
		"create-wireless-gateway-task": {
			Name:   "create-wireless-gateway-task",
			Fields: fields_create_wireless_gateway_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWirelessGatewayTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_wireless_gateway_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWirelessGatewayTask(ctx, input)
			},
		},
		"create-wireless-gateway-task-definition": {
			Name:   "create-wireless-gateway-task-definition",
			Fields: fields_create_wireless_gateway_task_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWirelessGatewayTaskDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_wireless_gateway_task_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWirelessGatewayTaskDefinition(ctx, input)
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
		"delete-device-profile": {
			Name:   "delete-device-profile",
			Fields: fields_delete_device_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDeviceProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_device_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDeviceProfile(ctx, input)
			},
		},
		"delete-fuota-task": {
			Name:   "delete-fuota-task",
			Fields: fields_delete_fuota_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFuotaTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_fuota_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFuotaTask(ctx, input)
			},
		},
		"delete-multicast-group": {
			Name:   "delete-multicast-group",
			Fields: fields_delete_multicast_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMulticastGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_multicast_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMulticastGroup(ctx, input)
			},
		},
		"delete-network-analyzer-configuration": {
			Name:   "delete-network-analyzer-configuration",
			Fields: fields_delete_network_analyzer_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteNetworkAnalyzerConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_network_analyzer_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteNetworkAnalyzerConfiguration(ctx, input)
			},
		},
		"delete-queued-messages": {
			Name:   "delete-queued-messages",
			Fields: fields_delete_queued_messages,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteQueuedMessagesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_queued_messages, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteQueuedMessages(ctx, input)
			},
		},
		"delete-service-profile": {
			Name:   "delete-service-profile",
			Fields: fields_delete_service_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteServiceProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_service_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteServiceProfile(ctx, input)
			},
		},
		"delete-wireless-device": {
			Name:   "delete-wireless-device",
			Fields: fields_delete_wireless_device,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWirelessDeviceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_wireless_device, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWirelessDevice(ctx, input)
			},
		},
		"delete-wireless-device-import-task": {
			Name:   "delete-wireless-device-import-task",
			Fields: fields_delete_wireless_device_import_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWirelessDeviceImportTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_wireless_device_import_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWirelessDeviceImportTask(ctx, input)
			},
		},
		"delete-wireless-gateway": {
			Name:   "delete-wireless-gateway",
			Fields: fields_delete_wireless_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWirelessGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_wireless_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWirelessGateway(ctx, input)
			},
		},
		"delete-wireless-gateway-task": {
			Name:   "delete-wireless-gateway-task",
			Fields: fields_delete_wireless_gateway_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWirelessGatewayTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_wireless_gateway_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWirelessGatewayTask(ctx, input)
			},
		},
		"delete-wireless-gateway-task-definition": {
			Name:   "delete-wireless-gateway-task-definition",
			Fields: fields_delete_wireless_gateway_task_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWirelessGatewayTaskDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_wireless_gateway_task_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWirelessGatewayTaskDefinition(ctx, input)
			},
		},
		"deregister-wireless-device": {
			Name:   "deregister-wireless-device",
			Fields: fields_deregister_wireless_device,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterWirelessDeviceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_wireless_device, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterWirelessDevice(ctx, input)
			},
		},
		"disassociate-aws-account-from-partner-account": {
			Name:   "disassociate-aws-account-from-partner-account",
			Fields: fields_disassociate_aws_account_from_partner_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateAwsAccountFromPartnerAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_aws_account_from_partner_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateAwsAccountFromPartnerAccount(ctx, input)
			},
		},
		"disassociate-multicast-group-from-fuota-task": {
			Name:   "disassociate-multicast-group-from-fuota-task",
			Fields: fields_disassociate_multicast_group_from_fuota_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateMulticastGroupFromFuotaTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_multicast_group_from_fuota_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateMulticastGroupFromFuotaTask(ctx, input)
			},
		},
		"disassociate-wireless-device-from-fuota-task": {
			Name:   "disassociate-wireless-device-from-fuota-task",
			Fields: fields_disassociate_wireless_device_from_fuota_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateWirelessDeviceFromFuotaTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_wireless_device_from_fuota_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateWirelessDeviceFromFuotaTask(ctx, input)
			},
		},
		"disassociate-wireless-device-from-multicast-group": {
			Name:   "disassociate-wireless-device-from-multicast-group",
			Fields: fields_disassociate_wireless_device_from_multicast_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateWirelessDeviceFromMulticastGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_wireless_device_from_multicast_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateWirelessDeviceFromMulticastGroup(ctx, input)
			},
		},
		"disassociate-wireless-device-from-thing": {
			Name:   "disassociate-wireless-device-from-thing",
			Fields: fields_disassociate_wireless_device_from_thing,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateWirelessDeviceFromThingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_wireless_device_from_thing, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateWirelessDeviceFromThing(ctx, input)
			},
		},
		"disassociate-wireless-gateway-from-certificate": {
			Name:   "disassociate-wireless-gateway-from-certificate",
			Fields: fields_disassociate_wireless_gateway_from_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateWirelessGatewayFromCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_wireless_gateway_from_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateWirelessGatewayFromCertificate(ctx, input)
			},
		},
		"disassociate-wireless-gateway-from-thing": {
			Name:   "disassociate-wireless-gateway-from-thing",
			Fields: fields_disassociate_wireless_gateway_from_thing,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateWirelessGatewayFromThingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_wireless_gateway_from_thing, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateWirelessGatewayFromThing(ctx, input)
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
		"get-device-profile": {
			Name:   "get-device-profile",
			Fields: fields_get_device_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDeviceProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_device_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDeviceProfile(ctx, input)
			},
		},
		"get-event-configuration-by-resource-types": {
			Name:   "get-event-configuration-by-resource-types",
			Fields: fields_get_event_configuration_by_resource_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEventConfigurationByResourceTypesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_event_configuration_by_resource_types, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEventConfigurationByResourceTypes(ctx, input)
			},
		},
		"get-fuota-task": {
			Name:   "get-fuota-task",
			Fields: fields_get_fuota_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFuotaTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_fuota_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFuotaTask(ctx, input)
			},
		},
		"get-log-levels-by-resource-types": {
			Name:   "get-log-levels-by-resource-types",
			Fields: fields_get_log_levels_by_resource_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLogLevelsByResourceTypesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_log_levels_by_resource_types, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLogLevelsByResourceTypes(ctx, input)
			},
		},
		"get-metric-configuration": {
			Name:   "get-metric-configuration",
			Fields: fields_get_metric_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMetricConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_metric_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMetricConfiguration(ctx, input)
			},
		},
		"get-metrics": {
			Name:   "get-metrics",
			Fields: fields_get_metrics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMetricsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_metrics, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMetrics(ctx, input)
			},
		},
		"get-multicast-group": {
			Name:   "get-multicast-group",
			Fields: fields_get_multicast_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMulticastGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_multicast_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMulticastGroup(ctx, input)
			},
		},
		"get-multicast-group-session": {
			Name:   "get-multicast-group-session",
			Fields: fields_get_multicast_group_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMulticastGroupSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_multicast_group_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMulticastGroupSession(ctx, input)
			},
		},
		"get-network-analyzer-configuration": {
			Name:   "get-network-analyzer-configuration",
			Fields: fields_get_network_analyzer_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetNetworkAnalyzerConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_network_analyzer_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetNetworkAnalyzerConfiguration(ctx, input)
			},
		},
		"get-partner-account": {
			Name:   "get-partner-account",
			Fields: fields_get_partner_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPartnerAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_partner_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPartnerAccount(ctx, input)
			},
		},
		"get-position": {
			Name:   "get-position",
			Fields: fields_get_position,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPositionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_position, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPosition(ctx, input)
			},
		},
		"get-position-configuration": {
			Name:   "get-position-configuration",
			Fields: fields_get_position_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPositionConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_position_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPositionConfiguration(ctx, input)
			},
		},
		"get-position-estimate": {
			Name:   "get-position-estimate",
			Fields: fields_get_position_estimate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPositionEstimateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_position_estimate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPositionEstimate(ctx, input)
			},
		},
		"get-resource-event-configuration": {
			Name:   "get-resource-event-configuration",
			Fields: fields_get_resource_event_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourceEventConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_resource_event_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResourceEventConfiguration(ctx, input)
			},
		},
		"get-resource-log-level": {
			Name:   "get-resource-log-level",
			Fields: fields_get_resource_log_level,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourceLogLevelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_resource_log_level, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResourceLogLevel(ctx, input)
			},
		},
		"get-resource-position": {
			Name:   "get-resource-position",
			Fields: fields_get_resource_position,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourcePositionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_resource_position, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResourcePosition(ctx, input)
			},
		},
		"get-service-endpoint": {
			Name:   "get-service-endpoint",
			Fields: fields_get_service_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetServiceEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_service_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetServiceEndpoint(ctx, input)
			},
		},
		"get-service-profile": {
			Name:   "get-service-profile",
			Fields: fields_get_service_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetServiceProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_service_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetServiceProfile(ctx, input)
			},
		},
		"get-wireless-device": {
			Name:   "get-wireless-device",
			Fields: fields_get_wireless_device,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWirelessDeviceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_wireless_device, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetWirelessDevice(ctx, input)
			},
		},
		"get-wireless-device-import-task": {
			Name:   "get-wireless-device-import-task",
			Fields: fields_get_wireless_device_import_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWirelessDeviceImportTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_wireless_device_import_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetWirelessDeviceImportTask(ctx, input)
			},
		},
		"get-wireless-device-statistics": {
			Name:   "get-wireless-device-statistics",
			Fields: fields_get_wireless_device_statistics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWirelessDeviceStatisticsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_wireless_device_statistics, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetWirelessDeviceStatistics(ctx, input)
			},
		},
		"get-wireless-gateway": {
			Name:   "get-wireless-gateway",
			Fields: fields_get_wireless_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWirelessGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_wireless_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetWirelessGateway(ctx, input)
			},
		},
		"get-wireless-gateway-certificate": {
			Name:   "get-wireless-gateway-certificate",
			Fields: fields_get_wireless_gateway_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWirelessGatewayCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_wireless_gateway_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetWirelessGatewayCertificate(ctx, input)
			},
		},
		"get-wireless-gateway-firmware-information": {
			Name:   "get-wireless-gateway-firmware-information",
			Fields: fields_get_wireless_gateway_firmware_information,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWirelessGatewayFirmwareInformationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_wireless_gateway_firmware_information, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetWirelessGatewayFirmwareInformation(ctx, input)
			},
		},
		"get-wireless-gateway-statistics": {
			Name:   "get-wireless-gateway-statistics",
			Fields: fields_get_wireless_gateway_statistics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWirelessGatewayStatisticsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_wireless_gateway_statistics, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetWirelessGatewayStatistics(ctx, input)
			},
		},
		"get-wireless-gateway-task": {
			Name:   "get-wireless-gateway-task",
			Fields: fields_get_wireless_gateway_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWirelessGatewayTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_wireless_gateway_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetWirelessGatewayTask(ctx, input)
			},
		},
		"get-wireless-gateway-task-definition": {
			Name:   "get-wireless-gateway-task-definition",
			Fields: fields_get_wireless_gateway_task_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWirelessGatewayTaskDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_wireless_gateway_task_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetWirelessGatewayTaskDefinition(ctx, input)
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
		"list-device-profiles": {
			Name:   "list-device-profiles",
			Fields: fields_list_device_profiles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDeviceProfilesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_device_profiles, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDeviceProfiles(ctx, input)
				}
				var results []*svc.ListDeviceProfilesOutput
				p := svc.NewListDeviceProfilesPaginator(client, input)
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
		"list-devices-for-wireless-device-import-task": {
			Name:   "list-devices-for-wireless-device-import-task",
			Fields: fields_list_devices_for_wireless_device_import_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDevicesForWirelessDeviceImportTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_devices_for_wireless_device_import_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListDevicesForWirelessDeviceImportTask(ctx, input)
			},
		},
		"list-event-configurations": {
			Name:   "list-event-configurations",
			Fields: fields_list_event_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEventConfigurationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_event_configurations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListEventConfigurations(ctx, input)
			},
		},
		"list-fuota-tasks": {
			Name:   "list-fuota-tasks",
			Fields: fields_list_fuota_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFuotaTasksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_fuota_tasks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFuotaTasks(ctx, input)
				}
				var results []*svc.ListFuotaTasksOutput
				p := svc.NewListFuotaTasksPaginator(client, input)
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
		"list-multicast-groups": {
			Name:   "list-multicast-groups",
			Fields: fields_list_multicast_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMulticastGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_multicast_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMulticastGroups(ctx, input)
				}
				var results []*svc.ListMulticastGroupsOutput
				p := svc.NewListMulticastGroupsPaginator(client, input)
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
		"list-multicast-groups-by-fuota-task": {
			Name:   "list-multicast-groups-by-fuota-task",
			Fields: fields_list_multicast_groups_by_fuota_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMulticastGroupsByFuotaTaskInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_multicast_groups_by_fuota_task, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMulticastGroupsByFuotaTask(ctx, input)
				}
				var results []*svc.ListMulticastGroupsByFuotaTaskOutput
				p := svc.NewListMulticastGroupsByFuotaTaskPaginator(client, input)
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
		"list-network-analyzer-configurations": {
			Name:   "list-network-analyzer-configurations",
			Fields: fields_list_network_analyzer_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListNetworkAnalyzerConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_network_analyzer_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListNetworkAnalyzerConfigurations(ctx, input)
				}
				var results []*svc.ListNetworkAnalyzerConfigurationsOutput
				p := svc.NewListNetworkAnalyzerConfigurationsPaginator(client, input)
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
		"list-partner-accounts": {
			Name:   "list-partner-accounts",
			Fields: fields_list_partner_accounts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPartnerAccountsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_partner_accounts, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListPartnerAccounts(ctx, input)
			},
		},
		"list-position-configurations": {
			Name:   "list-position-configurations",
			Fields: fields_list_position_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPositionConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_position_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPositionConfigurations(ctx, input)
				}
				var results []*svc.ListPositionConfigurationsOutput
				p := svc.NewListPositionConfigurationsPaginator(client, input)
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
		"list-queued-messages": {
			Name:   "list-queued-messages",
			Fields: fields_list_queued_messages,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListQueuedMessagesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_queued_messages, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListQueuedMessages(ctx, input)
				}
				var results []*svc.ListQueuedMessagesOutput
				p := svc.NewListQueuedMessagesPaginator(client, input)
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
		"list-service-profiles": {
			Name:   "list-service-profiles",
			Fields: fields_list_service_profiles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListServiceProfilesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_service_profiles, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListServiceProfiles(ctx, input)
				}
				var results []*svc.ListServiceProfilesOutput
				p := svc.NewListServiceProfilesPaginator(client, input)
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
		"list-wireless-device-import-tasks": {
			Name:   "list-wireless-device-import-tasks",
			Fields: fields_list_wireless_device_import_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWirelessDeviceImportTasksInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_wireless_device_import_tasks, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListWirelessDeviceImportTasks(ctx, input)
			},
		},
		"list-wireless-devices": {
			Name:   "list-wireless-devices",
			Fields: fields_list_wireless_devices,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWirelessDevicesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_wireless_devices, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWirelessDevices(ctx, input)
				}
				var results []*svc.ListWirelessDevicesOutput
				p := svc.NewListWirelessDevicesPaginator(client, input)
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
		"list-wireless-gateway-task-definitions": {
			Name:   "list-wireless-gateway-task-definitions",
			Fields: fields_list_wireless_gateway_task_definitions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWirelessGatewayTaskDefinitionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_wireless_gateway_task_definitions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListWirelessGatewayTaskDefinitions(ctx, input)
			},
		},
		"list-wireless-gateways": {
			Name:   "list-wireless-gateways",
			Fields: fields_list_wireless_gateways,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWirelessGatewaysInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_wireless_gateways, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWirelessGateways(ctx, input)
				}
				var results []*svc.ListWirelessGatewaysOutput
				p := svc.NewListWirelessGatewaysPaginator(client, input)
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
		"put-position-configuration": {
			Name:   "put-position-configuration",
			Fields: fields_put_position_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutPositionConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_position_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutPositionConfiguration(ctx, input)
			},
		},
		"put-resource-log-level": {
			Name:   "put-resource-log-level",
			Fields: fields_put_resource_log_level,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutResourceLogLevelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_resource_log_level, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutResourceLogLevel(ctx, input)
			},
		},
		"reset-all-resource-log-levels": {
			Name:   "reset-all-resource-log-levels",
			Fields: fields_reset_all_resource_log_levels,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResetAllResourceLogLevelsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reset_all_resource_log_levels, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResetAllResourceLogLevels(ctx, input)
			},
		},
		"reset-resource-log-level": {
			Name:   "reset-resource-log-level",
			Fields: fields_reset_resource_log_level,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResetResourceLogLevelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reset_resource_log_level, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResetResourceLogLevel(ctx, input)
			},
		},
		"send-data-to-multicast-group": {
			Name:   "send-data-to-multicast-group",
			Fields: fields_send_data_to_multicast_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendDataToMulticastGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_data_to_multicast_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendDataToMulticastGroup(ctx, input)
			},
		},
		"send-data-to-wireless-device": {
			Name:   "send-data-to-wireless-device",
			Fields: fields_send_data_to_wireless_device,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendDataToWirelessDeviceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_data_to_wireless_device, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendDataToWirelessDevice(ctx, input)
			},
		},
		"start-bulk-associate-wireless-device-with-multicast-group": {
			Name:   "start-bulk-associate-wireless-device-with-multicast-group",
			Fields: fields_start_bulk_associate_wireless_device_with_multicast_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartBulkAssociateWirelessDeviceWithMulticastGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_bulk_associate_wireless_device_with_multicast_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartBulkAssociateWirelessDeviceWithMulticastGroup(ctx, input)
			},
		},
		"start-bulk-disassociate-wireless-device-from-multicast-group": {
			Name:   "start-bulk-disassociate-wireless-device-from-multicast-group",
			Fields: fields_start_bulk_disassociate_wireless_device_from_multicast_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartBulkDisassociateWirelessDeviceFromMulticastGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_bulk_disassociate_wireless_device_from_multicast_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartBulkDisassociateWirelessDeviceFromMulticastGroup(ctx, input)
			},
		},
		"start-fuota-task": {
			Name:   "start-fuota-task",
			Fields: fields_start_fuota_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartFuotaTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_fuota_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartFuotaTask(ctx, input)
			},
		},
		"start-multicast-group-session": {
			Name:   "start-multicast-group-session",
			Fields: fields_start_multicast_group_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartMulticastGroupSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_multicast_group_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartMulticastGroupSession(ctx, input)
			},
		},
		"start-single-wireless-device-import-task": {
			Name:   "start-single-wireless-device-import-task",
			Fields: fields_start_single_wireless_device_import_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartSingleWirelessDeviceImportTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_single_wireless_device_import_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartSingleWirelessDeviceImportTask(ctx, input)
			},
		},
		"start-wireless-device-import-task": {
			Name:   "start-wireless-device-import-task",
			Fields: fields_start_wireless_device_import_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartWirelessDeviceImportTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_wireless_device_import_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartWirelessDeviceImportTask(ctx, input)
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
		"test-wireless-device": {
			Name:   "test-wireless-device",
			Fields: fields_test_wireless_device,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TestWirelessDeviceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_test_wireless_device, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TestWirelessDevice(ctx, input)
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
		"update-event-configuration-by-resource-types": {
			Name:   "update-event-configuration-by-resource-types",
			Fields: fields_update_event_configuration_by_resource_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEventConfigurationByResourceTypesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_event_configuration_by_resource_types, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEventConfigurationByResourceTypes(ctx, input)
			},
		},
		"update-fuota-task": {
			Name:   "update-fuota-task",
			Fields: fields_update_fuota_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFuotaTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_fuota_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFuotaTask(ctx, input)
			},
		},
		"update-log-levels-by-resource-types": {
			Name:   "update-log-levels-by-resource-types",
			Fields: fields_update_log_levels_by_resource_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateLogLevelsByResourceTypesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_log_levels_by_resource_types, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateLogLevelsByResourceTypes(ctx, input)
			},
		},
		"update-metric-configuration": {
			Name:   "update-metric-configuration",
			Fields: fields_update_metric_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMetricConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_metric_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMetricConfiguration(ctx, input)
			},
		},
		"update-multicast-group": {
			Name:   "update-multicast-group",
			Fields: fields_update_multicast_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMulticastGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_multicast_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMulticastGroup(ctx, input)
			},
		},
		"update-network-analyzer-configuration": {
			Name:   "update-network-analyzer-configuration",
			Fields: fields_update_network_analyzer_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateNetworkAnalyzerConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_network_analyzer_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateNetworkAnalyzerConfiguration(ctx, input)
			},
		},
		"update-partner-account": {
			Name:   "update-partner-account",
			Fields: fields_update_partner_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePartnerAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_partner_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePartnerAccount(ctx, input)
			},
		},
		"update-position": {
			Name:   "update-position",
			Fields: fields_update_position,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePositionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_position, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePosition(ctx, input)
			},
		},
		"update-resource-event-configuration": {
			Name:   "update-resource-event-configuration",
			Fields: fields_update_resource_event_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateResourceEventConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_resource_event_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateResourceEventConfiguration(ctx, input)
			},
		},
		"update-resource-position": {
			Name:   "update-resource-position",
			Fields: fields_update_resource_position,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateResourcePositionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_resource_position, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateResourcePosition(ctx, input)
			},
		},
		"update-wireless-device": {
			Name:   "update-wireless-device",
			Fields: fields_update_wireless_device,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateWirelessDeviceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_wireless_device, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateWirelessDevice(ctx, input)
			},
		},
		"update-wireless-device-import-task": {
			Name:   "update-wireless-device-import-task",
			Fields: fields_update_wireless_device_import_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateWirelessDeviceImportTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_wireless_device_import_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateWirelessDeviceImportTask(ctx, input)
			},
		},
		"update-wireless-gateway": {
			Name:   "update-wireless-gateway",
			Fields: fields_update_wireless_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateWirelessGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_wireless_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateWirelessGateway(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("iotwireless", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
