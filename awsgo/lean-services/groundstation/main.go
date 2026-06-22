package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/groundstation"
)

var fields_cancel_contact = []leanruntime.Field{
	{Name: "ContactId", Flag: "contact-id", Type: "*string", Required: true},
}

var fields_create_config = []leanruntime.Field{
	{Name: "ConfigData", Flag: "config-data", Type: "types.ConfigTypeData", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_dataflow_endpoint_group = []leanruntime.Field{
	{Name: "ContactPostPassDurationSeconds", Flag: "contact-post-pass-duration-seconds", Type: "*int32", Required: false},
	{Name: "ContactPrePassDurationSeconds", Flag: "contact-pre-pass-duration-seconds", Type: "*int32", Required: false},
	{Name: "EndpointDetails", Flag: "endpoint-details", Type: "[]types.EndpointDetails", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_dataflow_endpoint_group_v2 = []leanruntime.Field{
	{Name: "ContactPostPassDurationSeconds", Flag: "contact-post-pass-duration-seconds", Type: "*int32", Required: false},
	{Name: "ContactPrePassDurationSeconds", Flag: "contact-pre-pass-duration-seconds", Type: "*int32", Required: false},
	{Name: "Endpoints", Flag: "endpoints", Type: "[]types.CreateEndpointDetails", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_ephemeris = []leanruntime.Field{
	{Name: "Enabled", Flag: "enabled", Type: "*bool", Required: false},
	{Name: "Ephemeris", Flag: "ephemeris", Type: "types.EphemerisData", Required: false},
	{Name: "ExpirationTime", Flag: "expiration-time", Type: "*time.Time", Required: false},
	{Name: "KmsKeyArn", Flag: "kms-key-arn", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Priority", Flag: "priority", Type: "*int32", Required: false},
	{Name: "SatelliteId", Flag: "satellite-id", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_mission_profile = []leanruntime.Field{
	{Name: "ContactPostPassDurationSeconds", Flag: "contact-post-pass-duration-seconds", Type: "*int32", Required: false},
	{Name: "ContactPrePassDurationSeconds", Flag: "contact-pre-pass-duration-seconds", Type: "*int32", Required: false},
	{Name: "DataflowEdges", Flag: "dataflow-edges", Type: "[][]string", Required: true},
	{Name: "MinimumViableContactDurationSeconds", Flag: "minimum-viable-contact-duration-seconds", Type: "*int32", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "StreamsKmsKey", Flag: "streams-kms-key", Type: "types.KmsKey", Required: false},
	{Name: "StreamsKmsRole", Flag: "streams-kms-role", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TelemetrySinkConfigArn", Flag: "telemetry-sink-config-arn", Type: "*string", Required: false},
	{Name: "TrackingConfigArn", Flag: "tracking-config-arn", Type: "*string", Required: true},
}

var fields_delete_config = []leanruntime.Field{
	{Name: "ConfigId", Flag: "config-id", Type: "*string", Required: true},
	{Name: "ConfigType", Flag: "config-type", Type: "types.ConfigCapabilityType", Required: true},
}

var fields_delete_dataflow_endpoint_group = []leanruntime.Field{
	{Name: "DataflowEndpointGroupId", Flag: "dataflow-endpoint-group-id", Type: "*string", Required: true},
}

var fields_delete_ephemeris = []leanruntime.Field{
	{Name: "EphemerisId", Flag: "ephemeris-id", Type: "*string", Required: true},
}

var fields_delete_mission_profile = []leanruntime.Field{
	{Name: "MissionProfileId", Flag: "mission-profile-id", Type: "*string", Required: true},
}

var fields_describe_contact = []leanruntime.Field{
	{Name: "ContactId", Flag: "contact-id", Type: "*string", Required: true},
}

var fields_describe_ephemeris = []leanruntime.Field{
	{Name: "EphemerisId", Flag: "ephemeris-id", Type: "*string", Required: true},
}

var fields_get_agent_configuration = []leanruntime.Field{
	{Name: "AgentId", Flag: "agent-id", Type: "*string", Required: true},
}

var fields_get_agent_task_response_url = []leanruntime.Field{
	{Name: "AgentId", Flag: "agent-id", Type: "*string", Required: true},
	{Name: "TaskId", Flag: "task-id", Type: "*string", Required: true},
}

var fields_get_config = []leanruntime.Field{
	{Name: "ConfigId", Flag: "config-id", Type: "*string", Required: true},
	{Name: "ConfigType", Flag: "config-type", Type: "types.ConfigCapabilityType", Required: true},
}

var fields_get_dataflow_endpoint_group = []leanruntime.Field{
	{Name: "DataflowEndpointGroupId", Flag: "dataflow-endpoint-group-id", Type: "*string", Required: true},
}

var fields_get_minute_usage = []leanruntime.Field{
	{Name: "Month", Flag: "month", Type: "*int32", Required: true},
	{Name: "Year", Flag: "year", Type: "*int32", Required: true},
}

var fields_get_mission_profile = []leanruntime.Field{
	{Name: "MissionProfileId", Flag: "mission-profile-id", Type: "*string", Required: true},
}

var fields_get_satellite = []leanruntime.Field{
	{Name: "SatelliteId", Flag: "satellite-id", Type: "*string", Required: true},
}

var fields_list_configs = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_contacts = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "Ephemeris", Flag: "ephemeris", Type: "types.EphemerisFilter", Required: false},
	{Name: "GroundStation", Flag: "ground-station", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MissionProfileArn", Flag: "mission-profile-arn", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SatelliteArn", Flag: "satellite-arn", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
	{Name: "StatusList", Flag: "status-list", Type: "[]types.ContactStatus", Required: true},
}

var fields_list_dataflow_endpoint_groups = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_ephemerides = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "EphemerisType", Flag: "ephemeris-type", Type: "types.EphemerisType", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SatelliteId", Flag: "satellite-id", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
	{Name: "StatusList", Flag: "status-list", Type: "[]types.EphemerisStatus", Required: false},
}

var fields_list_ground_stations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SatelliteId", Flag: "satellite-id", Type: "*string", Required: false},
}

var fields_list_mission_profiles = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_satellites = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_register_agent = []leanruntime.Field{
	{Name: "AgentDetails", Flag: "agent-details", Type: "*types.AgentDetails", Required: true},
	{Name: "DiscoveryData", Flag: "discovery-data", Type: "*types.DiscoveryData", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_reserve_contact = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "GroundStation", Flag: "ground-station", Type: "*string", Required: true},
	{Name: "MissionProfileArn", Flag: "mission-profile-arn", Type: "*string", Required: true},
	{Name: "SatelliteArn", Flag: "satellite-arn", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TrackingOverrides", Flag: "tracking-overrides", Type: "*types.TrackingOverrides", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_agent_status = []leanruntime.Field{
	{Name: "AgentId", Flag: "agent-id", Type: "*string", Required: true},
	{Name: "AggregateStatus", Flag: "aggregate-status", Type: "*types.AggregateStatus", Required: true},
	{Name: "ComponentStatuses", Flag: "component-statuses", Type: "[]types.ComponentStatusData", Required: true},
	{Name: "TaskId", Flag: "task-id", Type: "*string", Required: true},
}

var fields_update_config = []leanruntime.Field{
	{Name: "ConfigData", Flag: "config-data", Type: "types.ConfigTypeData", Required: true},
	{Name: "ConfigId", Flag: "config-id", Type: "*string", Required: true},
	{Name: "ConfigType", Flag: "config-type", Type: "types.ConfigCapabilityType", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_update_ephemeris = []leanruntime.Field{
	{Name: "Enabled", Flag: "enabled", Type: "*bool", Required: true},
	{Name: "EphemerisId", Flag: "ephemeris-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Priority", Flag: "priority", Type: "*int32", Required: false},
}

var fields_update_mission_profile = []leanruntime.Field{
	{Name: "ContactPostPassDurationSeconds", Flag: "contact-post-pass-duration-seconds", Type: "*int32", Required: false},
	{Name: "ContactPrePassDurationSeconds", Flag: "contact-pre-pass-duration-seconds", Type: "*int32", Required: false},
	{Name: "DataflowEdges", Flag: "dataflow-edges", Type: "[][]string", Required: false},
	{Name: "MinimumViableContactDurationSeconds", Flag: "minimum-viable-contact-duration-seconds", Type: "*int32", Required: false},
	{Name: "MissionProfileId", Flag: "mission-profile-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "StreamsKmsKey", Flag: "streams-kms-key", Type: "types.KmsKey", Required: false},
	{Name: "StreamsKmsRole", Flag: "streams-kms-role", Type: "*string", Required: false},
	{Name: "TelemetrySinkConfigArn", Flag: "telemetry-sink-config-arn", Type: "*string", Required: false},
	{Name: "TrackingConfigArn", Flag: "tracking-config-arn", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"cancel-contact": {
			Name:   "cancel-contact",
			Fields: fields_cancel_contact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelContactInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_contact, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelContact(ctx, input)
			},
		},
		"create-config": {
			Name:   "create-config",
			Fields: fields_create_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateConfig(ctx, input)
			},
		},
		"create-dataflow-endpoint-group": {
			Name:   "create-dataflow-endpoint-group",
			Fields: fields_create_dataflow_endpoint_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDataflowEndpointGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_dataflow_endpoint_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDataflowEndpointGroup(ctx, input)
			},
		},
		"create-dataflow-endpoint-group-v2": {
			Name:   "create-dataflow-endpoint-group-v2",
			Fields: fields_create_dataflow_endpoint_group_v2,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDataflowEndpointGroupV2Input{}
				if _, err := leanruntime.ApplyInput(input, fields_create_dataflow_endpoint_group_v2, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDataflowEndpointGroupV2(ctx, input)
			},
		},
		"create-ephemeris": {
			Name:   "create-ephemeris",
			Fields: fields_create_ephemeris,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEphemerisInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_ephemeris, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEphemeris(ctx, input)
			},
		},
		"create-mission-profile": {
			Name:   "create-mission-profile",
			Fields: fields_create_mission_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMissionProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_mission_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMissionProfile(ctx, input)
			},
		},
		"delete-config": {
			Name:   "delete-config",
			Fields: fields_delete_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConfig(ctx, input)
			},
		},
		"delete-dataflow-endpoint-group": {
			Name:   "delete-dataflow-endpoint-group",
			Fields: fields_delete_dataflow_endpoint_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDataflowEndpointGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_dataflow_endpoint_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDataflowEndpointGroup(ctx, input)
			},
		},
		"delete-ephemeris": {
			Name:   "delete-ephemeris",
			Fields: fields_delete_ephemeris,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEphemerisInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_ephemeris, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEphemeris(ctx, input)
			},
		},
		"delete-mission-profile": {
			Name:   "delete-mission-profile",
			Fields: fields_delete_mission_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMissionProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_mission_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMissionProfile(ctx, input)
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
		"describe-ephemeris": {
			Name:   "describe-ephemeris",
			Fields: fields_describe_ephemeris,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEphemerisInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_ephemeris, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEphemeris(ctx, input)
			},
		},
		"get-agent-configuration": {
			Name:   "get-agent-configuration",
			Fields: fields_get_agent_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAgentConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_agent_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAgentConfiguration(ctx, input)
			},
		},
		"get-agent-task-response-url": {
			Name:   "get-agent-task-response-url",
			Fields: fields_get_agent_task_response_url,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAgentTaskResponseUrlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_agent_task_response_url, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAgentTaskResponseUrl(ctx, input)
			},
		},
		"get-config": {
			Name:   "get-config",
			Fields: fields_get_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConfig(ctx, input)
			},
		},
		"get-dataflow-endpoint-group": {
			Name:   "get-dataflow-endpoint-group",
			Fields: fields_get_dataflow_endpoint_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDataflowEndpointGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_dataflow_endpoint_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDataflowEndpointGroup(ctx, input)
			},
		},
		"get-minute-usage": {
			Name:   "get-minute-usage",
			Fields: fields_get_minute_usage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMinuteUsageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_minute_usage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMinuteUsage(ctx, input)
			},
		},
		"get-mission-profile": {
			Name:   "get-mission-profile",
			Fields: fields_get_mission_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMissionProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_mission_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMissionProfile(ctx, input)
			},
		},
		"get-satellite": {
			Name:   "get-satellite",
			Fields: fields_get_satellite,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSatelliteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_satellite, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSatellite(ctx, input)
			},
		},
		"list-configs": {
			Name:   "list-configs",
			Fields: fields_list_configs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConfigsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_configs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListConfigs(ctx, input)
				}
				var results []*svc.ListConfigsOutput
				p := svc.NewListConfigsPaginator(client, input)
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
		"list-dataflow-endpoint-groups": {
			Name:   "list-dataflow-endpoint-groups",
			Fields: fields_list_dataflow_endpoint_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDataflowEndpointGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_dataflow_endpoint_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDataflowEndpointGroups(ctx, input)
				}
				var results []*svc.ListDataflowEndpointGroupsOutput
				p := svc.NewListDataflowEndpointGroupsPaginator(client, input)
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
		"list-ephemerides": {
			Name:   "list-ephemerides",
			Fields: fields_list_ephemerides,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEphemeridesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_ephemerides, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEphemerides(ctx, input)
				}
				var results []*svc.ListEphemeridesOutput
				p := svc.NewListEphemeridesPaginator(client, input)
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
		"list-ground-stations": {
			Name:   "list-ground-stations",
			Fields: fields_list_ground_stations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGroundStationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_ground_stations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListGroundStations(ctx, input)
				}
				var results []*svc.ListGroundStationsOutput
				p := svc.NewListGroundStationsPaginator(client, input)
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
		"list-mission-profiles": {
			Name:   "list-mission-profiles",
			Fields: fields_list_mission_profiles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMissionProfilesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_mission_profiles, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMissionProfiles(ctx, input)
				}
				var results []*svc.ListMissionProfilesOutput
				p := svc.NewListMissionProfilesPaginator(client, input)
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
		"list-satellites": {
			Name:   "list-satellites",
			Fields: fields_list_satellites,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSatellitesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_satellites, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSatellites(ctx, input)
				}
				var results []*svc.ListSatellitesOutput
				p := svc.NewListSatellitesPaginator(client, input)
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
		"register-agent": {
			Name:   "register-agent",
			Fields: fields_register_agent,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterAgentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_agent, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterAgent(ctx, input)
			},
		},
		"reserve-contact": {
			Name:   "reserve-contact",
			Fields: fields_reserve_contact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ReserveContactInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reserve_contact, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ReserveContact(ctx, input)
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
		"update-config": {
			Name:   "update-config",
			Fields: fields_update_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateConfig(ctx, input)
			},
		},
		"update-ephemeris": {
			Name:   "update-ephemeris",
			Fields: fields_update_ephemeris,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEphemerisInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_ephemeris, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEphemeris(ctx, input)
			},
		},
		"update-mission-profile": {
			Name:   "update-mission-profile",
			Fields: fields_update_mission_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMissionProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_mission_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMissionProfile(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("groundstation", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
