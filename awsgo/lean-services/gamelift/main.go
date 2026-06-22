package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/gamelift"
)

var fields_accept_match = []leanruntime.Field{
	{Name: "AcceptanceType", Flag: "acceptance-type", Type: "types.AcceptanceType", Required: true},
	{Name: "PlayerIds", Flag: "player-ids", Type: "[]string", Required: true},
	{Name: "TicketId", Flag: "ticket-id", Type: "*string", Required: true},
}

var fields_claim_game_server = []leanruntime.Field{
	{Name: "FilterOption", Flag: "filter-option", Type: "*types.ClaimFilterOption", Required: false},
	{Name: "GameServerData", Flag: "game-server-data", Type: "*string", Required: false},
	{Name: "GameServerGroupName", Flag: "game-server-group-name", Type: "*string", Required: true},
	{Name: "GameServerId", Flag: "game-server-id", Type: "*string", Required: false},
}

var fields_create_alias = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RoutingStrategy", Flag: "routing-strategy", Type: "*types.RoutingStrategy", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_build = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "OperatingSystem", Flag: "operating-system", Type: "types.OperatingSystem", Required: false},
	{Name: "ServerSdkVersion", Flag: "server-sdk-version", Type: "*string", Required: false},
	{Name: "StorageLocation", Flag: "storage-location", Type: "*types.S3Location", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Version", Flag: "version", Type: "*string", Required: false},
}

var fields_create_container_fleet = []leanruntime.Field{
	{Name: "BillingType", Flag: "billing-type", Type: "types.ContainerFleetBillingType", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "FleetRoleArn", Flag: "fleet-role-arn", Type: "*string", Required: true},
	{Name: "GameServerContainerGroupDefinitionName", Flag: "game-server-container-group-definition-name", Type: "*string", Required: false},
	{Name: "GameServerContainerGroupsPerInstance", Flag: "game-server-container-groups-per-instance", Type: "*int32", Required: false},
	{Name: "GameSessionCreationLimitPolicy", Flag: "game-session-creation-limit-policy", Type: "*types.GameSessionCreationLimitPolicy", Required: false},
	{Name: "InstanceConnectionPortRange", Flag: "instance-connection-port-range", Type: "*types.ConnectionPortRange", Required: false},
	{Name: "InstanceInboundPermissions", Flag: "instance-inbound-permissions", Type: "[]types.IpPermission", Required: false},
	{Name: "InstanceType", Flag: "instance-type", Type: "*string", Required: false},
	{Name: "Locations", Flag: "locations", Type: "[]types.LocationConfiguration", Required: false},
	{Name: "LogConfiguration", Flag: "log-configuration", Type: "*types.LogConfiguration", Required: false},
	{Name: "MetricGroups", Flag: "metric-groups", Type: "[]string", Required: false},
	{Name: "NewGameSessionProtectionPolicy", Flag: "new-game-session-protection-policy", Type: "types.ProtectionPolicy", Required: false},
	{Name: "PerInstanceContainerGroupDefinitionName", Flag: "per-instance-container-group-definition-name", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_container_group_definition = []leanruntime.Field{
	{Name: "ContainerGroupType", Flag: "container-group-type", Type: "types.ContainerGroupType", Required: false},
	{Name: "GameServerContainerDefinition", Flag: "game-server-container-definition", Type: "*types.GameServerContainerDefinitionInput", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OperatingSystem", Flag: "operating-system", Type: "types.ContainerOperatingSystem", Required: true},
	{Name: "SupportContainerDefinitions", Flag: "support-container-definitions", Type: "[]types.SupportContainerDefinitionInput", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TotalMemoryLimitMebibytes", Flag: "total-memory-limit-mebibytes", Type: "*int32", Required: true},
	{Name: "TotalVcpuLimit", Flag: "total-vcpu-limit", Type: "*float64", Required: true},
	{Name: "VersionDescription", Flag: "version-description", Type: "*string", Required: false},
}

var fields_create_fleet = []leanruntime.Field{
	{Name: "AnywhereConfiguration", Flag: "anywhere-configuration", Type: "*types.AnywhereConfiguration", Required: false},
	{Name: "BuildId", Flag: "build-id", Type: "*string", Required: false},
	{Name: "CertificateConfiguration", Flag: "certificate-configuration", Type: "*types.CertificateConfiguration", Required: false},
	{Name: "ComputeType", Flag: "compute-type", Type: "types.ComputeType", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EC2InboundPermissions", Flag: "ec2-inbound-permissions", Type: "[]types.IpPermission", Required: false},
	{Name: "EC2InstanceType", Flag: "ec2-instance-type", Type: "types.EC2InstanceType", Required: false},
	{Name: "FleetType", Flag: "fleet-type", Type: "types.FleetType", Required: false},
	{Name: "InstanceRoleArn", Flag: "instance-role-arn", Type: "*string", Required: false},
	{Name: "InstanceRoleCredentialsProvider", Flag: "instance-role-credentials-provider", Type: "types.InstanceRoleCredentialsProvider", Required: false},
	{Name: "Locations", Flag: "locations", Type: "[]types.LocationConfiguration", Required: false},
	{Name: "LogPaths", Flag: "log-paths", Type: "[]string", Required: false},
	{Name: "MetricGroups", Flag: "metric-groups", Type: "[]string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NewGameSessionProtectionPolicy", Flag: "new-game-session-protection-policy", Type: "types.ProtectionPolicy", Required: false},
	{Name: "PeerVpcAwsAccountId", Flag: "peer-vpc-aws-account-id", Type: "*string", Required: false},
	{Name: "PeerVpcId", Flag: "peer-vpc-id", Type: "*string", Required: false},
	{Name: "ResourceCreationLimitPolicy", Flag: "resource-creation-limit-policy", Type: "*types.ResourceCreationLimitPolicy", Required: false},
	{Name: "RuntimeConfiguration", Flag: "runtime-configuration", Type: "*types.RuntimeConfiguration", Required: false},
	{Name: "ScriptId", Flag: "script-id", Type: "*string", Required: false},
	{Name: "ServerLaunchParameters", Flag: "server-launch-parameters", Type: "*string", Required: false},
	{Name: "ServerLaunchPath", Flag: "server-launch-path", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_fleet_locations = []leanruntime.Field{
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "Locations", Flag: "locations", Type: "[]types.LocationConfiguration", Required: true},
}

var fields_create_game_server_group = []leanruntime.Field{
	{Name: "AutoScalingPolicy", Flag: "auto-scaling-policy", Type: "*types.GameServerGroupAutoScalingPolicy", Required: false},
	{Name: "BalancingStrategy", Flag: "balancing-strategy", Type: "types.BalancingStrategy", Required: false},
	{Name: "GameServerGroupName", Flag: "game-server-group-name", Type: "*string", Required: true},
	{Name: "GameServerProtectionPolicy", Flag: "game-server-protection-policy", Type: "types.GameServerProtectionPolicy", Required: false},
	{Name: "InstanceDefinitions", Flag: "instance-definitions", Type: "[]types.InstanceDefinition", Required: true},
	{Name: "LaunchTemplate", Flag: "launch-template", Type: "*types.LaunchTemplateSpecification", Required: true},
	{Name: "MaxSize", Flag: "max-size", Type: "*int32", Required: true},
	{Name: "MinSize", Flag: "min-size", Type: "*int32", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VpcSubnets", Flag: "vpc-subnets", Type: "[]string", Required: false},
}

var fields_create_game_session = []leanruntime.Field{
	{Name: "AliasId", Flag: "alias-id", Type: "*string", Required: false},
	{Name: "CreatorId", Flag: "creator-id", Type: "*string", Required: false},
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: false},
	{Name: "GameProperties", Flag: "game-properties", Type: "[]types.GameProperty", Required: false},
	{Name: "GameSessionData", Flag: "game-session-data", Type: "*string", Required: false},
	{Name: "GameSessionId", Flag: "game-session-id", Type: "*string", Required: false},
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: false},
	{Name: "Location", Flag: "location", Type: "*string", Required: false},
	{Name: "MaximumPlayerSessionCount", Flag: "maximum-player-session-count", Type: "*int32", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_create_game_session_queue = []leanruntime.Field{
	{Name: "CustomEventData", Flag: "custom-event-data", Type: "*string", Required: false},
	{Name: "Destinations", Flag: "destinations", Type: "[]types.GameSessionQueueDestination", Required: false},
	{Name: "FilterConfiguration", Flag: "filter-configuration", Type: "*types.FilterConfiguration", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NotificationTarget", Flag: "notification-target", Type: "*string", Required: false},
	{Name: "PlayerLatencyPolicies", Flag: "player-latency-policies", Type: "[]types.PlayerLatencyPolicy", Required: false},
	{Name: "PriorityConfiguration", Flag: "priority-configuration", Type: "*types.PriorityConfiguration", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TimeoutInSeconds", Flag: "timeout-in-seconds", Type: "*int32", Required: false},
}

var fields_create_location = []leanruntime.Field{
	{Name: "LocationName", Flag: "location-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_matchmaking_configuration = []leanruntime.Field{
	{Name: "AcceptanceRequired", Flag: "acceptance-required", Type: "*bool", Required: true},
	{Name: "AcceptanceTimeoutSeconds", Flag: "acceptance-timeout-seconds", Type: "*int32", Required: false},
	{Name: "AdditionalPlayerCount", Flag: "additional-player-count", Type: "*int32", Required: false},
	{Name: "BackfillMode", Flag: "backfill-mode", Type: "types.BackfillMode", Required: false},
	{Name: "CustomEventData", Flag: "custom-event-data", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "FlexMatchMode", Flag: "flex-match-mode", Type: "types.FlexMatchMode", Required: false},
	{Name: "GameProperties", Flag: "game-properties", Type: "[]types.GameProperty", Required: false},
	{Name: "GameSessionData", Flag: "game-session-data", Type: "*string", Required: false},
	{Name: "GameSessionQueueArns", Flag: "game-session-queue-arns", Type: "[]string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NotificationTarget", Flag: "notification-target", Type: "*string", Required: false},
	{Name: "RequestTimeoutSeconds", Flag: "request-timeout-seconds", Type: "*int32", Required: true},
	{Name: "RuleSetName", Flag: "rule-set-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_matchmaking_rule_set = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RuleSetBody", Flag: "rule-set-body", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_player_session = []leanruntime.Field{
	{Name: "GameSessionId", Flag: "game-session-id", Type: "*string", Required: true},
	{Name: "PlayerData", Flag: "player-data", Type: "*string", Required: false},
	{Name: "PlayerId", Flag: "player-id", Type: "*string", Required: true},
}

var fields_create_player_sessions = []leanruntime.Field{
	{Name: "GameSessionId", Flag: "game-session-id", Type: "*string", Required: true},
	{Name: "PlayerDataMap", Flag: "player-data-map", Type: "map[string]string", Required: false},
	{Name: "PlayerIds", Flag: "player-ids", Type: "[]string", Required: true},
}

var fields_create_script = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "NodeJsVersion", Flag: "node-js-version", Type: "*string", Required: false},
	{Name: "StorageLocation", Flag: "storage-location", Type: "*types.S3Location", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Version", Flag: "version", Type: "*string", Required: false},
	{Name: "ZipFile", Flag: "zip-file", Type: "[]byte", Required: false},
}

var fields_create_vpc_peering_authorization = []leanruntime.Field{
	{Name: "GameLiftAwsAccountId", Flag: "game-lift-aws-account-id", Type: "*string", Required: true},
	{Name: "PeerVpcId", Flag: "peer-vpc-id", Type: "*string", Required: true},
}

var fields_create_vpc_peering_connection = []leanruntime.Field{
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "PeerVpcAwsAccountId", Flag: "peer-vpc-aws-account-id", Type: "*string", Required: true},
	{Name: "PeerVpcId", Flag: "peer-vpc-id", Type: "*string", Required: true},
}

var fields_delete_alias = []leanruntime.Field{
	{Name: "AliasId", Flag: "alias-id", Type: "*string", Required: true},
}

var fields_delete_build = []leanruntime.Field{
	{Name: "BuildId", Flag: "build-id", Type: "*string", Required: true},
}

var fields_delete_container_fleet = []leanruntime.Field{
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
}

var fields_delete_container_group_definition = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "VersionCountToRetain", Flag: "version-count-to-retain", Type: "*int32", Required: false},
	{Name: "VersionNumber", Flag: "version-number", Type: "*int32", Required: false},
}

var fields_delete_fleet = []leanruntime.Field{
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
}

var fields_delete_fleet_locations = []leanruntime.Field{
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "Locations", Flag: "locations", Type: "[]string", Required: true},
}

var fields_delete_game_server_group = []leanruntime.Field{
	{Name: "DeleteOption", Flag: "delete-option", Type: "types.GameServerGroupDeleteOption", Required: false},
	{Name: "GameServerGroupName", Flag: "game-server-group-name", Type: "*string", Required: true},
}

var fields_delete_game_session_queue = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_location = []leanruntime.Field{
	{Name: "LocationName", Flag: "location-name", Type: "*string", Required: true},
}

var fields_delete_matchmaking_configuration = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_matchmaking_rule_set = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_scaling_policy = []leanruntime.Field{
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_script = []leanruntime.Field{
	{Name: "ScriptId", Flag: "script-id", Type: "*string", Required: true},
}

var fields_delete_vpc_peering_authorization = []leanruntime.Field{
	{Name: "GameLiftAwsAccountId", Flag: "game-lift-aws-account-id", Type: "*string", Required: true},
	{Name: "PeerVpcId", Flag: "peer-vpc-id", Type: "*string", Required: true},
}

var fields_delete_vpc_peering_connection = []leanruntime.Field{
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "VpcPeeringConnectionId", Flag: "vpc-peering-connection-id", Type: "*string", Required: true},
}

var fields_deregister_compute = []leanruntime.Field{
	{Name: "ComputeName", Flag: "compute-name", Type: "*string", Required: true},
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
}

var fields_deregister_game_server = []leanruntime.Field{
	{Name: "GameServerGroupName", Flag: "game-server-group-name", Type: "*string", Required: true},
	{Name: "GameServerId", Flag: "game-server-id", Type: "*string", Required: true},
}

var fields_describe_alias = []leanruntime.Field{
	{Name: "AliasId", Flag: "alias-id", Type: "*string", Required: true},
}

var fields_describe_build = []leanruntime.Field{
	{Name: "BuildId", Flag: "build-id", Type: "*string", Required: true},
}

var fields_describe_compute = []leanruntime.Field{
	{Name: "ComputeName", Flag: "compute-name", Type: "*string", Required: true},
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
}

var fields_describe_container_fleet = []leanruntime.Field{
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
}

var fields_describe_container_group_definition = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "VersionNumber", Flag: "version-number", Type: "*int32", Required: false},
}

var fields_describe_ec2_instance_limits = []leanruntime.Field{
	{Name: "EC2InstanceType", Flag: "ec2-instance-type", Type: "types.EC2InstanceType", Required: false},
	{Name: "Location", Flag: "location", Type: "*string", Required: false},
}

var fields_describe_fleet_attributes = []leanruntime.Field{
	{Name: "FleetIds", Flag: "fleet-ids", Type: "[]string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_fleet_capacity = []leanruntime.Field{
	{Name: "FleetIds", Flag: "fleet-ids", Type: "[]string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_fleet_deployment = []leanruntime.Field{
	{Name: "DeploymentId", Flag: "deployment-id", Type: "*string", Required: false},
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
}

var fields_describe_fleet_events = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
}

var fields_describe_fleet_location_attributes = []leanruntime.Field{
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Locations", Flag: "locations", Type: "[]string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_fleet_location_capacity = []leanruntime.Field{
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "Location", Flag: "location", Type: "*string", Required: true},
}

var fields_describe_fleet_location_utilization = []leanruntime.Field{
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "Location", Flag: "location", Type: "*string", Required: true},
}

var fields_describe_fleet_port_settings = []leanruntime.Field{
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "Location", Flag: "location", Type: "*string", Required: false},
}

var fields_describe_fleet_utilization = []leanruntime.Field{
	{Name: "FleetIds", Flag: "fleet-ids", Type: "[]string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_game_server = []leanruntime.Field{
	{Name: "GameServerGroupName", Flag: "game-server-group-name", Type: "*string", Required: true},
	{Name: "GameServerId", Flag: "game-server-id", Type: "*string", Required: true},
}

var fields_describe_game_server_group = []leanruntime.Field{
	{Name: "GameServerGroupName", Flag: "game-server-group-name", Type: "*string", Required: true},
}

var fields_describe_game_server_instances = []leanruntime.Field{
	{Name: "GameServerGroupName", Flag: "game-server-group-name", Type: "*string", Required: true},
	{Name: "InstanceIds", Flag: "instance-ids", Type: "[]string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_game_session_details = []leanruntime.Field{
	{Name: "AliasId", Flag: "alias-id", Type: "*string", Required: false},
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: false},
	{Name: "GameSessionId", Flag: "game-session-id", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Location", Flag: "location", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StatusFilter", Flag: "status-filter", Type: "*string", Required: false},
}

var fields_describe_game_session_placement = []leanruntime.Field{
	{Name: "PlacementId", Flag: "placement-id", Type: "*string", Required: true},
}

var fields_describe_game_session_queues = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Names", Flag: "names", Type: "[]string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_game_sessions = []leanruntime.Field{
	{Name: "AliasId", Flag: "alias-id", Type: "*string", Required: false},
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: false},
	{Name: "GameSessionId", Flag: "game-session-id", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Location", Flag: "location", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StatusFilter", Flag: "status-filter", Type: "*string", Required: false},
}

var fields_describe_instances = []leanruntime.Field{
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Location", Flag: "location", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_matchmaking = []leanruntime.Field{
	{Name: "TicketIds", Flag: "ticket-ids", Type: "[]string", Required: true},
}

var fields_describe_matchmaking_configurations = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Names", Flag: "names", Type: "[]string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RuleSetName", Flag: "rule-set-name", Type: "*string", Required: false},
}

var fields_describe_matchmaking_rule_sets = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Names", Flag: "names", Type: "[]string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_player_sessions = []leanruntime.Field{
	{Name: "GameSessionId", Flag: "game-session-id", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PlayerId", Flag: "player-id", Type: "*string", Required: false},
	{Name: "PlayerSessionId", Flag: "player-session-id", Type: "*string", Required: false},
	{Name: "PlayerSessionStatusFilter", Flag: "player-session-status-filter", Type: "*string", Required: false},
}

var fields_describe_runtime_configuration = []leanruntime.Field{
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
}

var fields_describe_scaling_policies = []leanruntime.Field{
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Location", Flag: "location", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StatusFilter", Flag: "status-filter", Type: "types.ScalingStatusType", Required: false},
}

var fields_describe_script = []leanruntime.Field{
	{Name: "ScriptId", Flag: "script-id", Type: "*string", Required: true},
}

var fields_describe_vpc_peering_authorizations = []leanruntime.Field{}

var fields_describe_vpc_peering_connections = []leanruntime.Field{
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: false},
}

var fields_get_compute_access = []leanruntime.Field{
	{Name: "ComputeName", Flag: "compute-name", Type: "*string", Required: true},
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
}

var fields_get_compute_auth_token = []leanruntime.Field{
	{Name: "ComputeName", Flag: "compute-name", Type: "*string", Required: true},
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
}

var fields_get_game_session_log_url = []leanruntime.Field{
	{Name: "GameSessionId", Flag: "game-session-id", Type: "*string", Required: true},
}

var fields_get_instance_access = []leanruntime.Field{
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_list_aliases = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RoutingStrategyType", Flag: "routing-strategy-type", Type: "types.RoutingStrategyType", Required: false},
}

var fields_list_builds = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.BuildStatus", Required: false},
}

var fields_list_compute = []leanruntime.Field{
	{Name: "ComputeStatus", Flag: "compute-status", Type: "types.ListComputeInputStatus", Required: false},
	{Name: "ContainerGroupDefinitionName", Flag: "container-group-definition-name", Type: "*string", Required: false},
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Location", Flag: "location", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_container_fleets = []leanruntime.Field{
	{Name: "ContainerGroupDefinitionName", Flag: "container-group-definition-name", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_container_group_definition_versions = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_container_group_definitions = []leanruntime.Field{
	{Name: "ContainerGroupType", Flag: "container-group-type", Type: "types.ContainerGroupType", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_fleet_deployments = []leanruntime.Field{
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_fleets = []leanruntime.Field{
	{Name: "BuildId", Flag: "build-id", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ScriptId", Flag: "script-id", Type: "*string", Required: false},
}

var fields_list_game_server_groups = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_game_servers = []leanruntime.Field{
	{Name: "GameServerGroupName", Flag: "game-server-group-name", Type: "*string", Required: true},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_list_locations = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.LocationFilter", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_scripts = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_put_scaling_policy = []leanruntime.Field{
	{Name: "ComparisonOperator", Flag: "comparison-operator", Type: "types.ComparisonOperatorType", Required: false},
	{Name: "EvaluationPeriods", Flag: "evaluation-periods", Type: "*int32", Required: false},
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "MetricName", Flag: "metric-name", Type: "types.MetricName", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "PolicyType", Flag: "policy-type", Type: "types.PolicyType", Required: false},
	{Name: "ScalingAdjustment", Flag: "scaling-adjustment", Type: "*int32", Required: false},
	{Name: "ScalingAdjustmentType", Flag: "scaling-adjustment-type", Type: "types.ScalingAdjustmentType", Required: false},
	{Name: "TargetConfiguration", Flag: "target-configuration", Type: "*types.TargetConfiguration", Required: false},
	{Name: "Threshold", Flag: "threshold", Type: "*float64", Required: false},
}

var fields_register_compute = []leanruntime.Field{
	{Name: "CertificatePath", Flag: "certificate-path", Type: "*string", Required: false},
	{Name: "ComputeName", Flag: "compute-name", Type: "*string", Required: true},
	{Name: "DnsName", Flag: "dns-name", Type: "*string", Required: false},
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "IpAddress", Flag: "ip-address", Type: "*string", Required: false},
	{Name: "Location", Flag: "location", Type: "*string", Required: false},
}

var fields_register_game_server = []leanruntime.Field{
	{Name: "ConnectionInfo", Flag: "connection-info", Type: "*string", Required: false},
	{Name: "GameServerData", Flag: "game-server-data", Type: "*string", Required: false},
	{Name: "GameServerGroupName", Flag: "game-server-group-name", Type: "*string", Required: true},
	{Name: "GameServerId", Flag: "game-server-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_request_upload_credentials = []leanruntime.Field{
	{Name: "BuildId", Flag: "build-id", Type: "*string", Required: true},
}

var fields_resolve_alias = []leanruntime.Field{
	{Name: "AliasId", Flag: "alias-id", Type: "*string", Required: true},
}

var fields_resume_game_server_group = []leanruntime.Field{
	{Name: "GameServerGroupName", Flag: "game-server-group-name", Type: "*string", Required: true},
	{Name: "ResumeActions", Flag: "resume-actions", Type: "[]types.GameServerGroupAction", Required: true},
}

var fields_search_game_sessions = []leanruntime.Field{
	{Name: "AliasId", Flag: "alias-id", Type: "*string", Required: false},
	{Name: "FilterExpression", Flag: "filter-expression", Type: "*string", Required: false},
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Location", Flag: "location", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortExpression", Flag: "sort-expression", Type: "*string", Required: false},
}

var fields_start_fleet_actions = []leanruntime.Field{
	{Name: "Actions", Flag: "actions", Type: "[]types.FleetAction", Required: true},
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "Location", Flag: "location", Type: "*string", Required: false},
}

var fields_start_game_session_placement = []leanruntime.Field{
	{Name: "DesiredPlayerSessions", Flag: "desired-player-sessions", Type: "[]types.DesiredPlayerSession", Required: false},
	{Name: "GameProperties", Flag: "game-properties", Type: "[]types.GameProperty", Required: false},
	{Name: "GameSessionData", Flag: "game-session-data", Type: "*string", Required: false},
	{Name: "GameSessionName", Flag: "game-session-name", Type: "*string", Required: false},
	{Name: "GameSessionQueueName", Flag: "game-session-queue-name", Type: "*string", Required: true},
	{Name: "MaximumPlayerSessionCount", Flag: "maximum-player-session-count", Type: "*int32", Required: true},
	{Name: "PlacementId", Flag: "placement-id", Type: "*string", Required: true},
	{Name: "PlayerLatencies", Flag: "player-latencies", Type: "[]types.PlayerLatency", Required: false},
	{Name: "PriorityConfigurationOverride", Flag: "priority-configuration-override", Type: "*types.PriorityConfigurationOverride", Required: false},
}

var fields_start_match_backfill = []leanruntime.Field{
	{Name: "ConfigurationName", Flag: "configuration-name", Type: "*string", Required: true},
	{Name: "GameSessionArn", Flag: "game-session-arn", Type: "*string", Required: false},
	{Name: "Players", Flag: "players", Type: "[]types.Player", Required: true},
	{Name: "TicketId", Flag: "ticket-id", Type: "*string", Required: false},
}

var fields_start_matchmaking = []leanruntime.Field{
	{Name: "ConfigurationName", Flag: "configuration-name", Type: "*string", Required: true},
	{Name: "Players", Flag: "players", Type: "[]types.Player", Required: true},
	{Name: "TicketId", Flag: "ticket-id", Type: "*string", Required: false},
}

var fields_stop_fleet_actions = []leanruntime.Field{
	{Name: "Actions", Flag: "actions", Type: "[]types.FleetAction", Required: true},
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "Location", Flag: "location", Type: "*string", Required: false},
}

var fields_stop_game_session_placement = []leanruntime.Field{
	{Name: "PlacementId", Flag: "placement-id", Type: "*string", Required: true},
}

var fields_stop_matchmaking = []leanruntime.Field{
	{Name: "TicketId", Flag: "ticket-id", Type: "*string", Required: true},
}

var fields_suspend_game_server_group = []leanruntime.Field{
	{Name: "GameServerGroupName", Flag: "game-server-group-name", Type: "*string", Required: true},
	{Name: "SuspendActions", Flag: "suspend-actions", Type: "[]types.GameServerGroupAction", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_terminate_game_session = []leanruntime.Field{
	{Name: "GameSessionId", Flag: "game-session-id", Type: "*string", Required: true},
	{Name: "TerminationMode", Flag: "termination-mode", Type: "types.TerminationMode", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_alias = []leanruntime.Field{
	{Name: "AliasId", Flag: "alias-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "RoutingStrategy", Flag: "routing-strategy", Type: "*types.RoutingStrategy", Required: false},
}

var fields_update_build = []leanruntime.Field{
	{Name: "BuildId", Flag: "build-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Version", Flag: "version", Type: "*string", Required: false},
}

var fields_update_container_fleet = []leanruntime.Field{
	{Name: "DeploymentConfiguration", Flag: "deployment-configuration", Type: "*types.DeploymentConfiguration", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "GameServerContainerGroupDefinitionName", Flag: "game-server-container-group-definition-name", Type: "*string", Required: false},
	{Name: "GameServerContainerGroupsPerInstance", Flag: "game-server-container-groups-per-instance", Type: "*int32", Required: false},
	{Name: "GameSessionCreationLimitPolicy", Flag: "game-session-creation-limit-policy", Type: "*types.GameSessionCreationLimitPolicy", Required: false},
	{Name: "InstanceConnectionPortRange", Flag: "instance-connection-port-range", Type: "*types.ConnectionPortRange", Required: false},
	{Name: "InstanceInboundPermissionAuthorizations", Flag: "instance-inbound-permission-authorizations", Type: "[]types.IpPermission", Required: false},
	{Name: "InstanceInboundPermissionRevocations", Flag: "instance-inbound-permission-revocations", Type: "[]types.IpPermission", Required: false},
	{Name: "LogConfiguration", Flag: "log-configuration", Type: "*types.LogConfiguration", Required: false},
	{Name: "MetricGroups", Flag: "metric-groups", Type: "[]string", Required: false},
	{Name: "NewGameSessionProtectionPolicy", Flag: "new-game-session-protection-policy", Type: "types.ProtectionPolicy", Required: false},
	{Name: "PerInstanceContainerGroupDefinitionName", Flag: "per-instance-container-group-definition-name", Type: "*string", Required: false},
	{Name: "RemoveAttributes", Flag: "remove-attributes", Type: "[]types.ContainerFleetRemoveAttribute", Required: false},
}

var fields_update_container_group_definition = []leanruntime.Field{
	{Name: "GameServerContainerDefinition", Flag: "game-server-container-definition", Type: "*types.GameServerContainerDefinitionInput", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OperatingSystem", Flag: "operating-system", Type: "types.ContainerOperatingSystem", Required: false},
	{Name: "SourceVersionNumber", Flag: "source-version-number", Type: "*int32", Required: false},
	{Name: "SupportContainerDefinitions", Flag: "support-container-definitions", Type: "[]types.SupportContainerDefinitionInput", Required: false},
	{Name: "TotalMemoryLimitMebibytes", Flag: "total-memory-limit-mebibytes", Type: "*int32", Required: false},
	{Name: "TotalVcpuLimit", Flag: "total-vcpu-limit", Type: "*float64", Required: false},
	{Name: "VersionDescription", Flag: "version-description", Type: "*string", Required: false},
}

var fields_update_fleet_attributes = []leanruntime.Field{
	{Name: "AnywhereConfiguration", Flag: "anywhere-configuration", Type: "*types.AnywhereConfiguration", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "MetricGroups", Flag: "metric-groups", Type: "[]string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "NewGameSessionProtectionPolicy", Flag: "new-game-session-protection-policy", Type: "types.ProtectionPolicy", Required: false},
	{Name: "ResourceCreationLimitPolicy", Flag: "resource-creation-limit-policy", Type: "*types.ResourceCreationLimitPolicy", Required: false},
}

var fields_update_fleet_capacity = []leanruntime.Field{
	{Name: "DesiredInstances", Flag: "desired-instances", Type: "*int32", Required: false},
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "Location", Flag: "location", Type: "*string", Required: false},
	{Name: "ManagedCapacityConfiguration", Flag: "managed-capacity-configuration", Type: "*types.ManagedCapacityConfiguration", Required: false},
	{Name: "MaxSize", Flag: "max-size", Type: "*int32", Required: false},
	{Name: "MinSize", Flag: "min-size", Type: "*int32", Required: false},
}

var fields_update_fleet_port_settings = []leanruntime.Field{
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "InboundPermissionAuthorizations", Flag: "inbound-permission-authorizations", Type: "[]types.IpPermission", Required: false},
	{Name: "InboundPermissionRevocations", Flag: "inbound-permission-revocations", Type: "[]types.IpPermission", Required: false},
}

var fields_update_game_server = []leanruntime.Field{
	{Name: "GameServerData", Flag: "game-server-data", Type: "*string", Required: false},
	{Name: "GameServerGroupName", Flag: "game-server-group-name", Type: "*string", Required: true},
	{Name: "GameServerId", Flag: "game-server-id", Type: "*string", Required: true},
	{Name: "HealthCheck", Flag: "health-check", Type: "types.GameServerHealthCheck", Required: false},
	{Name: "UtilizationStatus", Flag: "utilization-status", Type: "types.GameServerUtilizationStatus", Required: false},
}

var fields_update_game_server_group = []leanruntime.Field{
	{Name: "BalancingStrategy", Flag: "balancing-strategy", Type: "types.BalancingStrategy", Required: false},
	{Name: "GameServerGroupName", Flag: "game-server-group-name", Type: "*string", Required: true},
	{Name: "GameServerProtectionPolicy", Flag: "game-server-protection-policy", Type: "types.GameServerProtectionPolicy", Required: false},
	{Name: "InstanceDefinitions", Flag: "instance-definitions", Type: "[]types.InstanceDefinition", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
}

var fields_update_game_session = []leanruntime.Field{
	{Name: "GameProperties", Flag: "game-properties", Type: "[]types.GameProperty", Required: false},
	{Name: "GameSessionId", Flag: "game-session-id", Type: "*string", Required: true},
	{Name: "MaximumPlayerSessionCount", Flag: "maximum-player-session-count", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "PlayerSessionCreationPolicy", Flag: "player-session-creation-policy", Type: "types.PlayerSessionCreationPolicy", Required: false},
	{Name: "ProtectionPolicy", Flag: "protection-policy", Type: "types.ProtectionPolicy", Required: false},
}

var fields_update_game_session_queue = []leanruntime.Field{
	{Name: "CustomEventData", Flag: "custom-event-data", Type: "*string", Required: false},
	{Name: "Destinations", Flag: "destinations", Type: "[]types.GameSessionQueueDestination", Required: false},
	{Name: "FilterConfiguration", Flag: "filter-configuration", Type: "*types.FilterConfiguration", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NotificationTarget", Flag: "notification-target", Type: "*string", Required: false},
	{Name: "PlayerLatencyPolicies", Flag: "player-latency-policies", Type: "[]types.PlayerLatencyPolicy", Required: false},
	{Name: "PriorityConfiguration", Flag: "priority-configuration", Type: "*types.PriorityConfiguration", Required: false},
	{Name: "TimeoutInSeconds", Flag: "timeout-in-seconds", Type: "*int32", Required: false},
}

var fields_update_matchmaking_configuration = []leanruntime.Field{
	{Name: "AcceptanceRequired", Flag: "acceptance-required", Type: "*bool", Required: false},
	{Name: "AcceptanceTimeoutSeconds", Flag: "acceptance-timeout-seconds", Type: "*int32", Required: false},
	{Name: "AdditionalPlayerCount", Flag: "additional-player-count", Type: "*int32", Required: false},
	{Name: "BackfillMode", Flag: "backfill-mode", Type: "types.BackfillMode", Required: false},
	{Name: "CustomEventData", Flag: "custom-event-data", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "FlexMatchMode", Flag: "flex-match-mode", Type: "types.FlexMatchMode", Required: false},
	{Name: "GameProperties", Flag: "game-properties", Type: "[]types.GameProperty", Required: false},
	{Name: "GameSessionData", Flag: "game-session-data", Type: "*string", Required: false},
	{Name: "GameSessionQueueArns", Flag: "game-session-queue-arns", Type: "[]string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NotificationTarget", Flag: "notification-target", Type: "*string", Required: false},
	{Name: "RequestTimeoutSeconds", Flag: "request-timeout-seconds", Type: "*int32", Required: false},
	{Name: "RuleSetName", Flag: "rule-set-name", Type: "*string", Required: false},
}

var fields_update_runtime_configuration = []leanruntime.Field{
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "RuntimeConfiguration", Flag: "runtime-configuration", Type: "*types.RuntimeConfiguration", Required: true},
}

var fields_update_script = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "ScriptId", Flag: "script-id", Type: "*string", Required: true},
	{Name: "StorageLocation", Flag: "storage-location", Type: "*types.S3Location", Required: false},
	{Name: "Version", Flag: "version", Type: "*string", Required: false},
	{Name: "ZipFile", Flag: "zip-file", Type: "[]byte", Required: false},
}

var fields_validate_matchmaking_rule_set = []leanruntime.Field{
	{Name: "RuleSetBody", Flag: "rule-set-body", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"accept-match": {
			Name:   "accept-match",
			Fields: fields_accept_match,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AcceptMatchInput{}
				if _, err := leanruntime.ApplyInput(input, fields_accept_match, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AcceptMatch(ctx, input)
			},
		},
		"claim-game-server": {
			Name:   "claim-game-server",
			Fields: fields_claim_game_server,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ClaimGameServerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_claim_game_server, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ClaimGameServer(ctx, input)
			},
		},
		"create-alias": {
			Name:   "create-alias",
			Fields: fields_create_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAlias(ctx, input)
			},
		},
		"create-build": {
			Name:   "create-build",
			Fields: fields_create_build,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBuildInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_build, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBuild(ctx, input)
			},
		},
		"create-container-fleet": {
			Name:   "create-container-fleet",
			Fields: fields_create_container_fleet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateContainerFleetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_container_fleet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateContainerFleet(ctx, input)
			},
		},
		"create-container-group-definition": {
			Name:   "create-container-group-definition",
			Fields: fields_create_container_group_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateContainerGroupDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_container_group_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateContainerGroupDefinition(ctx, input)
			},
		},
		"create-fleet": {
			Name:   "create-fleet",
			Fields: fields_create_fleet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFleetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_fleet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFleet(ctx, input)
			},
		},
		"create-fleet-locations": {
			Name:   "create-fleet-locations",
			Fields: fields_create_fleet_locations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFleetLocationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_fleet_locations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFleetLocations(ctx, input)
			},
		},
		"create-game-server-group": {
			Name:   "create-game-server-group",
			Fields: fields_create_game_server_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateGameServerGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_game_server_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateGameServerGroup(ctx, input)
			},
		},
		"create-game-session": {
			Name:   "create-game-session",
			Fields: fields_create_game_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateGameSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_game_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateGameSession(ctx, input)
			},
		},
		"create-game-session-queue": {
			Name:   "create-game-session-queue",
			Fields: fields_create_game_session_queue,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateGameSessionQueueInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_game_session_queue, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateGameSessionQueue(ctx, input)
			},
		},
		"create-location": {
			Name:   "create-location",
			Fields: fields_create_location,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLocationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_location, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLocation(ctx, input)
			},
		},
		"create-matchmaking-configuration": {
			Name:   "create-matchmaking-configuration",
			Fields: fields_create_matchmaking_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMatchmakingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_matchmaking_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMatchmakingConfiguration(ctx, input)
			},
		},
		"create-matchmaking-rule-set": {
			Name:   "create-matchmaking-rule-set",
			Fields: fields_create_matchmaking_rule_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMatchmakingRuleSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_matchmaking_rule_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMatchmakingRuleSet(ctx, input)
			},
		},
		"create-player-session": {
			Name:   "create-player-session",
			Fields: fields_create_player_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePlayerSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_player_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePlayerSession(ctx, input)
			},
		},
		"create-player-sessions": {
			Name:   "create-player-sessions",
			Fields: fields_create_player_sessions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePlayerSessionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_player_sessions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePlayerSessions(ctx, input)
			},
		},
		"create-script": {
			Name:   "create-script",
			Fields: fields_create_script,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateScriptInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_script, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateScript(ctx, input)
			},
		},
		"create-vpc-peering-authorization": {
			Name:   "create-vpc-peering-authorization",
			Fields: fields_create_vpc_peering_authorization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVpcPeeringAuthorizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_vpc_peering_authorization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVpcPeeringAuthorization(ctx, input)
			},
		},
		"create-vpc-peering-connection": {
			Name:   "create-vpc-peering-connection",
			Fields: fields_create_vpc_peering_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVpcPeeringConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_vpc_peering_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVpcPeeringConnection(ctx, input)
			},
		},
		"delete-alias": {
			Name:   "delete-alias",
			Fields: fields_delete_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAlias(ctx, input)
			},
		},
		"delete-build": {
			Name:   "delete-build",
			Fields: fields_delete_build,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBuildInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_build, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBuild(ctx, input)
			},
		},
		"delete-container-fleet": {
			Name:   "delete-container-fleet",
			Fields: fields_delete_container_fleet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteContainerFleetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_container_fleet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteContainerFleet(ctx, input)
			},
		},
		"delete-container-group-definition": {
			Name:   "delete-container-group-definition",
			Fields: fields_delete_container_group_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteContainerGroupDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_container_group_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteContainerGroupDefinition(ctx, input)
			},
		},
		"delete-fleet": {
			Name:   "delete-fleet",
			Fields: fields_delete_fleet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFleetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_fleet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFleet(ctx, input)
			},
		},
		"delete-fleet-locations": {
			Name:   "delete-fleet-locations",
			Fields: fields_delete_fleet_locations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFleetLocationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_fleet_locations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFleetLocations(ctx, input)
			},
		},
		"delete-game-server-group": {
			Name:   "delete-game-server-group",
			Fields: fields_delete_game_server_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteGameServerGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_game_server_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteGameServerGroup(ctx, input)
			},
		},
		"delete-game-session-queue": {
			Name:   "delete-game-session-queue",
			Fields: fields_delete_game_session_queue,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteGameSessionQueueInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_game_session_queue, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteGameSessionQueue(ctx, input)
			},
		},
		"delete-location": {
			Name:   "delete-location",
			Fields: fields_delete_location,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLocationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_location, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLocation(ctx, input)
			},
		},
		"delete-matchmaking-configuration": {
			Name:   "delete-matchmaking-configuration",
			Fields: fields_delete_matchmaking_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMatchmakingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_matchmaking_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMatchmakingConfiguration(ctx, input)
			},
		},
		"delete-matchmaking-rule-set": {
			Name:   "delete-matchmaking-rule-set",
			Fields: fields_delete_matchmaking_rule_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMatchmakingRuleSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_matchmaking_rule_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMatchmakingRuleSet(ctx, input)
			},
		},
		"delete-scaling-policy": {
			Name:   "delete-scaling-policy",
			Fields: fields_delete_scaling_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteScalingPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_scaling_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteScalingPolicy(ctx, input)
			},
		},
		"delete-script": {
			Name:   "delete-script",
			Fields: fields_delete_script,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteScriptInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_script, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteScript(ctx, input)
			},
		},
		"delete-vpc-peering-authorization": {
			Name:   "delete-vpc-peering-authorization",
			Fields: fields_delete_vpc_peering_authorization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVpcPeeringAuthorizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_vpc_peering_authorization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVpcPeeringAuthorization(ctx, input)
			},
		},
		"delete-vpc-peering-connection": {
			Name:   "delete-vpc-peering-connection",
			Fields: fields_delete_vpc_peering_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVpcPeeringConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_vpc_peering_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVpcPeeringConnection(ctx, input)
			},
		},
		"deregister-compute": {
			Name:   "deregister-compute",
			Fields: fields_deregister_compute,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterComputeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_compute, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterCompute(ctx, input)
			},
		},
		"deregister-game-server": {
			Name:   "deregister-game-server",
			Fields: fields_deregister_game_server,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterGameServerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_game_server, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterGameServer(ctx, input)
			},
		},
		"describe-alias": {
			Name:   "describe-alias",
			Fields: fields_describe_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAlias(ctx, input)
			},
		},
		"describe-build": {
			Name:   "describe-build",
			Fields: fields_describe_build,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeBuildInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_build, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeBuild(ctx, input)
			},
		},
		"describe-compute": {
			Name:   "describe-compute",
			Fields: fields_describe_compute,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeComputeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_compute, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeCompute(ctx, input)
			},
		},
		"describe-container-fleet": {
			Name:   "describe-container-fleet",
			Fields: fields_describe_container_fleet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeContainerFleetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_container_fleet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeContainerFleet(ctx, input)
			},
		},
		"describe-container-group-definition": {
			Name:   "describe-container-group-definition",
			Fields: fields_describe_container_group_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeContainerGroupDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_container_group_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeContainerGroupDefinition(ctx, input)
			},
		},
		"describe-ec2-instance-limits": {
			Name:   "describe-ec2-instance-limits",
			Fields: fields_describe_ec2_instance_limits,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEC2InstanceLimitsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_ec2_instance_limits, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEC2InstanceLimits(ctx, input)
			},
		},
		"describe-fleet-attributes": {
			Name:   "describe-fleet-attributes",
			Fields: fields_describe_fleet_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFleetAttributesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_fleet_attributes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeFleetAttributes(ctx, input)
				}
				var results []*svc.DescribeFleetAttributesOutput
				p := svc.NewDescribeFleetAttributesPaginator(client, input)
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
		"describe-fleet-capacity": {
			Name:   "describe-fleet-capacity",
			Fields: fields_describe_fleet_capacity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFleetCapacityInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_fleet_capacity, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeFleetCapacity(ctx, input)
				}
				var results []*svc.DescribeFleetCapacityOutput
				p := svc.NewDescribeFleetCapacityPaginator(client, input)
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
		"describe-fleet-deployment": {
			Name:   "describe-fleet-deployment",
			Fields: fields_describe_fleet_deployment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFleetDeploymentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_fleet_deployment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeFleetDeployment(ctx, input)
			},
		},
		"describe-fleet-events": {
			Name:   "describe-fleet-events",
			Fields: fields_describe_fleet_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFleetEventsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_fleet_events, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeFleetEvents(ctx, input)
				}
				var results []*svc.DescribeFleetEventsOutput
				p := svc.NewDescribeFleetEventsPaginator(client, input)
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
		"describe-fleet-location-attributes": {
			Name:   "describe-fleet-location-attributes",
			Fields: fields_describe_fleet_location_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFleetLocationAttributesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_fleet_location_attributes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeFleetLocationAttributes(ctx, input)
				}
				var results []*svc.DescribeFleetLocationAttributesOutput
				p := svc.NewDescribeFleetLocationAttributesPaginator(client, input)
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
		"describe-fleet-location-capacity": {
			Name:   "describe-fleet-location-capacity",
			Fields: fields_describe_fleet_location_capacity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFleetLocationCapacityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_fleet_location_capacity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeFleetLocationCapacity(ctx, input)
			},
		},
		"describe-fleet-location-utilization": {
			Name:   "describe-fleet-location-utilization",
			Fields: fields_describe_fleet_location_utilization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFleetLocationUtilizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_fleet_location_utilization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeFleetLocationUtilization(ctx, input)
			},
		},
		"describe-fleet-port-settings": {
			Name:   "describe-fleet-port-settings",
			Fields: fields_describe_fleet_port_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFleetPortSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_fleet_port_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeFleetPortSettings(ctx, input)
			},
		},
		"describe-fleet-utilization": {
			Name:   "describe-fleet-utilization",
			Fields: fields_describe_fleet_utilization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFleetUtilizationInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_fleet_utilization, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeFleetUtilization(ctx, input)
				}
				var results []*svc.DescribeFleetUtilizationOutput
				p := svc.NewDescribeFleetUtilizationPaginator(client, input)
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
		"describe-game-server": {
			Name:   "describe-game-server",
			Fields: fields_describe_game_server,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeGameServerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_game_server, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeGameServer(ctx, input)
			},
		},
		"describe-game-server-group": {
			Name:   "describe-game-server-group",
			Fields: fields_describe_game_server_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeGameServerGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_game_server_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeGameServerGroup(ctx, input)
			},
		},
		"describe-game-server-instances": {
			Name:   "describe-game-server-instances",
			Fields: fields_describe_game_server_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeGameServerInstancesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_game_server_instances, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeGameServerInstances(ctx, input)
				}
				var results []*svc.DescribeGameServerInstancesOutput
				p := svc.NewDescribeGameServerInstancesPaginator(client, input)
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
		"describe-game-session-details": {
			Name:   "describe-game-session-details",
			Fields: fields_describe_game_session_details,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeGameSessionDetailsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_game_session_details, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeGameSessionDetails(ctx, input)
				}
				var results []*svc.DescribeGameSessionDetailsOutput
				p := svc.NewDescribeGameSessionDetailsPaginator(client, input)
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
		"describe-game-session-placement": {
			Name:   "describe-game-session-placement",
			Fields: fields_describe_game_session_placement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeGameSessionPlacementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_game_session_placement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeGameSessionPlacement(ctx, input)
			},
		},
		"describe-game-session-queues": {
			Name:   "describe-game-session-queues",
			Fields: fields_describe_game_session_queues,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeGameSessionQueuesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_game_session_queues, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeGameSessionQueues(ctx, input)
				}
				var results []*svc.DescribeGameSessionQueuesOutput
				p := svc.NewDescribeGameSessionQueuesPaginator(client, input)
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
		"describe-game-sessions": {
			Name:   "describe-game-sessions",
			Fields: fields_describe_game_sessions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeGameSessionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_game_sessions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeGameSessions(ctx, input)
				}
				var results []*svc.DescribeGameSessionsOutput
				p := svc.NewDescribeGameSessionsPaginator(client, input)
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
		"describe-instances": {
			Name:   "describe-instances",
			Fields: fields_describe_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInstancesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_instances, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeInstances(ctx, input)
				}
				var results []*svc.DescribeInstancesOutput
				p := svc.NewDescribeInstancesPaginator(client, input)
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
		"describe-matchmaking": {
			Name:   "describe-matchmaking",
			Fields: fields_describe_matchmaking,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMatchmakingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_matchmaking, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeMatchmaking(ctx, input)
			},
		},
		"describe-matchmaking-configurations": {
			Name:   "describe-matchmaking-configurations",
			Fields: fields_describe_matchmaking_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMatchmakingConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_matchmaking_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeMatchmakingConfigurations(ctx, input)
				}
				var results []*svc.DescribeMatchmakingConfigurationsOutput
				p := svc.NewDescribeMatchmakingConfigurationsPaginator(client, input)
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
		"describe-matchmaking-rule-sets": {
			Name:   "describe-matchmaking-rule-sets",
			Fields: fields_describe_matchmaking_rule_sets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMatchmakingRuleSetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_matchmaking_rule_sets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeMatchmakingRuleSets(ctx, input)
				}
				var results []*svc.DescribeMatchmakingRuleSetsOutput
				p := svc.NewDescribeMatchmakingRuleSetsPaginator(client, input)
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
		"describe-player-sessions": {
			Name:   "describe-player-sessions",
			Fields: fields_describe_player_sessions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePlayerSessionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_player_sessions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribePlayerSessions(ctx, input)
				}
				var results []*svc.DescribePlayerSessionsOutput
				p := svc.NewDescribePlayerSessionsPaginator(client, input)
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
		"describe-runtime-configuration": {
			Name:   "describe-runtime-configuration",
			Fields: fields_describe_runtime_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRuntimeConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_runtime_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeRuntimeConfiguration(ctx, input)
			},
		},
		"describe-scaling-policies": {
			Name:   "describe-scaling-policies",
			Fields: fields_describe_scaling_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeScalingPoliciesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_scaling_policies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeScalingPolicies(ctx, input)
				}
				var results []*svc.DescribeScalingPoliciesOutput
				p := svc.NewDescribeScalingPoliciesPaginator(client, input)
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
		"describe-script": {
			Name:   "describe-script",
			Fields: fields_describe_script,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeScriptInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_script, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeScript(ctx, input)
			},
		},
		"describe-vpc-peering-authorizations": {
			Name:   "describe-vpc-peering-authorizations",
			Fields: fields_describe_vpc_peering_authorizations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVpcPeeringAuthorizationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_vpc_peering_authorizations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeVpcPeeringAuthorizations(ctx, input)
			},
		},
		"describe-vpc-peering-connections": {
			Name:   "describe-vpc-peering-connections",
			Fields: fields_describe_vpc_peering_connections,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVpcPeeringConnectionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_vpc_peering_connections, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeVpcPeeringConnections(ctx, input)
			},
		},
		"get-compute-access": {
			Name:   "get-compute-access",
			Fields: fields_get_compute_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetComputeAccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_compute_access, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetComputeAccess(ctx, input)
			},
		},
		"get-compute-auth-token": {
			Name:   "get-compute-auth-token",
			Fields: fields_get_compute_auth_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetComputeAuthTokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_compute_auth_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetComputeAuthToken(ctx, input)
			},
		},
		"get-game-session-log-url": {
			Name:   "get-game-session-log-url",
			Fields: fields_get_game_session_log_url,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetGameSessionLogUrlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_game_session_log_url, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetGameSessionLogUrl(ctx, input)
			},
		},
		"get-instance-access": {
			Name:   "get-instance-access",
			Fields: fields_get_instance_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetInstanceAccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_instance_access, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetInstanceAccess(ctx, input)
			},
		},
		"list-aliases": {
			Name:   "list-aliases",
			Fields: fields_list_aliases,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAliasesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_aliases, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAliases(ctx, input)
				}
				var results []*svc.ListAliasesOutput
				p := svc.NewListAliasesPaginator(client, input)
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
		"list-builds": {
			Name:   "list-builds",
			Fields: fields_list_builds,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBuildsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_builds, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBuilds(ctx, input)
				}
				var results []*svc.ListBuildsOutput
				p := svc.NewListBuildsPaginator(client, input)
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
		"list-compute": {
			Name:   "list-compute",
			Fields: fields_list_compute,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListComputeInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_compute, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCompute(ctx, input)
				}
				var results []*svc.ListComputeOutput
				p := svc.NewListComputePaginator(client, input)
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
		"list-container-fleets": {
			Name:   "list-container-fleets",
			Fields: fields_list_container_fleets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListContainerFleetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_container_fleets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListContainerFleets(ctx, input)
				}
				var results []*svc.ListContainerFleetsOutput
				p := svc.NewListContainerFleetsPaginator(client, input)
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
		"list-container-group-definition-versions": {
			Name:   "list-container-group-definition-versions",
			Fields: fields_list_container_group_definition_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListContainerGroupDefinitionVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_container_group_definition_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListContainerGroupDefinitionVersions(ctx, input)
				}
				var results []*svc.ListContainerGroupDefinitionVersionsOutput
				p := svc.NewListContainerGroupDefinitionVersionsPaginator(client, input)
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
		"list-container-group-definitions": {
			Name:   "list-container-group-definitions",
			Fields: fields_list_container_group_definitions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListContainerGroupDefinitionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_container_group_definitions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListContainerGroupDefinitions(ctx, input)
				}
				var results []*svc.ListContainerGroupDefinitionsOutput
				p := svc.NewListContainerGroupDefinitionsPaginator(client, input)
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
		"list-fleet-deployments": {
			Name:   "list-fleet-deployments",
			Fields: fields_list_fleet_deployments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFleetDeploymentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_fleet_deployments, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFleetDeployments(ctx, input)
				}
				var results []*svc.ListFleetDeploymentsOutput
				p := svc.NewListFleetDeploymentsPaginator(client, input)
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
		"list-fleets": {
			Name:   "list-fleets",
			Fields: fields_list_fleets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFleetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_fleets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFleets(ctx, input)
				}
				var results []*svc.ListFleetsOutput
				p := svc.NewListFleetsPaginator(client, input)
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
		"list-game-server-groups": {
			Name:   "list-game-server-groups",
			Fields: fields_list_game_server_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGameServerGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_game_server_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListGameServerGroups(ctx, input)
				}
				var results []*svc.ListGameServerGroupsOutput
				p := svc.NewListGameServerGroupsPaginator(client, input)
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
		"list-game-servers": {
			Name:   "list-game-servers",
			Fields: fields_list_game_servers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGameServersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_game_servers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListGameServers(ctx, input)
				}
				var results []*svc.ListGameServersOutput
				p := svc.NewListGameServersPaginator(client, input)
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
		"list-locations": {
			Name:   "list-locations",
			Fields: fields_list_locations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLocationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_locations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLocations(ctx, input)
				}
				var results []*svc.ListLocationsOutput
				p := svc.NewListLocationsPaginator(client, input)
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
		"list-scripts": {
			Name:   "list-scripts",
			Fields: fields_list_scripts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListScriptsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_scripts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListScripts(ctx, input)
				}
				var results []*svc.ListScriptsOutput
				p := svc.NewListScriptsPaginator(client, input)
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
		"put-scaling-policy": {
			Name:   "put-scaling-policy",
			Fields: fields_put_scaling_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutScalingPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_scaling_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutScalingPolicy(ctx, input)
			},
		},
		"register-compute": {
			Name:   "register-compute",
			Fields: fields_register_compute,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterComputeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_compute, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterCompute(ctx, input)
			},
		},
		"register-game-server": {
			Name:   "register-game-server",
			Fields: fields_register_game_server,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterGameServerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_game_server, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterGameServer(ctx, input)
			},
		},
		"request-upload-credentials": {
			Name:   "request-upload-credentials",
			Fields: fields_request_upload_credentials,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RequestUploadCredentialsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_request_upload_credentials, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RequestUploadCredentials(ctx, input)
			},
		},
		"resolve-alias": {
			Name:   "resolve-alias",
			Fields: fields_resolve_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResolveAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_resolve_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResolveAlias(ctx, input)
			},
		},
		"resume-game-server-group": {
			Name:   "resume-game-server-group",
			Fields: fields_resume_game_server_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResumeGameServerGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_resume_game_server_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResumeGameServerGroup(ctx, input)
			},
		},
		"search-game-sessions": {
			Name:   "search-game-sessions",
			Fields: fields_search_game_sessions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchGameSessionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_game_sessions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchGameSessions(ctx, input)
				}
				var results []*svc.SearchGameSessionsOutput
				p := svc.NewSearchGameSessionsPaginator(client, input)
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
		"start-fleet-actions": {
			Name:   "start-fleet-actions",
			Fields: fields_start_fleet_actions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartFleetActionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_fleet_actions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartFleetActions(ctx, input)
			},
		},
		"start-game-session-placement": {
			Name:   "start-game-session-placement",
			Fields: fields_start_game_session_placement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartGameSessionPlacementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_game_session_placement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartGameSessionPlacement(ctx, input)
			},
		},
		"start-match-backfill": {
			Name:   "start-match-backfill",
			Fields: fields_start_match_backfill,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartMatchBackfillInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_match_backfill, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartMatchBackfill(ctx, input)
			},
		},
		"start-matchmaking": {
			Name:   "start-matchmaking",
			Fields: fields_start_matchmaking,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartMatchmakingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_matchmaking, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartMatchmaking(ctx, input)
			},
		},
		"stop-fleet-actions": {
			Name:   "stop-fleet-actions",
			Fields: fields_stop_fleet_actions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopFleetActionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_fleet_actions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopFleetActions(ctx, input)
			},
		},
		"stop-game-session-placement": {
			Name:   "stop-game-session-placement",
			Fields: fields_stop_game_session_placement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopGameSessionPlacementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_game_session_placement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopGameSessionPlacement(ctx, input)
			},
		},
		"stop-matchmaking": {
			Name:   "stop-matchmaking",
			Fields: fields_stop_matchmaking,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopMatchmakingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_matchmaking, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopMatchmaking(ctx, input)
			},
		},
		"suspend-game-server-group": {
			Name:   "suspend-game-server-group",
			Fields: fields_suspend_game_server_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SuspendGameServerGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_suspend_game_server_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SuspendGameServerGroup(ctx, input)
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
		"terminate-game-session": {
			Name:   "terminate-game-session",
			Fields: fields_terminate_game_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TerminateGameSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_terminate_game_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TerminateGameSession(ctx, input)
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
		"update-alias": {
			Name:   "update-alias",
			Fields: fields_update_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAlias(ctx, input)
			},
		},
		"update-build": {
			Name:   "update-build",
			Fields: fields_update_build,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateBuildInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_build, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateBuild(ctx, input)
			},
		},
		"update-container-fleet": {
			Name:   "update-container-fleet",
			Fields: fields_update_container_fleet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateContainerFleetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_container_fleet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateContainerFleet(ctx, input)
			},
		},
		"update-container-group-definition": {
			Name:   "update-container-group-definition",
			Fields: fields_update_container_group_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateContainerGroupDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_container_group_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateContainerGroupDefinition(ctx, input)
			},
		},
		"update-fleet-attributes": {
			Name:   "update-fleet-attributes",
			Fields: fields_update_fleet_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFleetAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_fleet_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFleetAttributes(ctx, input)
			},
		},
		"update-fleet-capacity": {
			Name:   "update-fleet-capacity",
			Fields: fields_update_fleet_capacity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFleetCapacityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_fleet_capacity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFleetCapacity(ctx, input)
			},
		},
		"update-fleet-port-settings": {
			Name:   "update-fleet-port-settings",
			Fields: fields_update_fleet_port_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFleetPortSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_fleet_port_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFleetPortSettings(ctx, input)
			},
		},
		"update-game-server": {
			Name:   "update-game-server",
			Fields: fields_update_game_server,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateGameServerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_game_server, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateGameServer(ctx, input)
			},
		},
		"update-game-server-group": {
			Name:   "update-game-server-group",
			Fields: fields_update_game_server_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateGameServerGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_game_server_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateGameServerGroup(ctx, input)
			},
		},
		"update-game-session": {
			Name:   "update-game-session",
			Fields: fields_update_game_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateGameSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_game_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateGameSession(ctx, input)
			},
		},
		"update-game-session-queue": {
			Name:   "update-game-session-queue",
			Fields: fields_update_game_session_queue,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateGameSessionQueueInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_game_session_queue, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateGameSessionQueue(ctx, input)
			},
		},
		"update-matchmaking-configuration": {
			Name:   "update-matchmaking-configuration",
			Fields: fields_update_matchmaking_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMatchmakingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_matchmaking_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMatchmakingConfiguration(ctx, input)
			},
		},
		"update-runtime-configuration": {
			Name:   "update-runtime-configuration",
			Fields: fields_update_runtime_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRuntimeConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_runtime_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRuntimeConfiguration(ctx, input)
			},
		},
		"update-script": {
			Name:   "update-script",
			Fields: fields_update_script,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateScriptInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_script, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateScript(ctx, input)
			},
		},
		"validate-matchmaking-rule-set": {
			Name:   "validate-matchmaking-rule-set",
			Fields: fields_validate_matchmaking_rule_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ValidateMatchmakingRuleSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_validate_matchmaking_rule_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ValidateMatchmakingRuleSet(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("gamelift", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
