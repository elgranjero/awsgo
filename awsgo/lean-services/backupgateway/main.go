package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/backupgateway"
)

var fields_associate_gateway_to_server = []leanruntime.Field{
	{Name: "GatewayArn", Flag: "gateway-arn", Type: "*string", Required: true},
	{Name: "ServerArn", Flag: "server-arn", Type: "*string", Required: true},
}

var fields_create_gateway = []leanruntime.Field{
	{Name: "ActivationKey", Flag: "activation-key", Type: "*string", Required: true},
	{Name: "GatewayDisplayName", Flag: "gateway-display-name", Type: "*string", Required: true},
	{Name: "GatewayType", Flag: "gateway-type", Type: "types.GatewayType", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_delete_gateway = []leanruntime.Field{
	{Name: "GatewayArn", Flag: "gateway-arn", Type: "*string", Required: true},
}

var fields_delete_hypervisor = []leanruntime.Field{
	{Name: "HypervisorArn", Flag: "hypervisor-arn", Type: "*string", Required: true},
}

var fields_disassociate_gateway_from_server = []leanruntime.Field{
	{Name: "GatewayArn", Flag: "gateway-arn", Type: "*string", Required: true},
}

var fields_get_bandwidth_rate_limit_schedule = []leanruntime.Field{
	{Name: "GatewayArn", Flag: "gateway-arn", Type: "*string", Required: true},
}

var fields_get_gateway = []leanruntime.Field{
	{Name: "GatewayArn", Flag: "gateway-arn", Type: "*string", Required: true},
}

var fields_get_hypervisor = []leanruntime.Field{
	{Name: "HypervisorArn", Flag: "hypervisor-arn", Type: "*string", Required: true},
}

var fields_get_hypervisor_property_mappings = []leanruntime.Field{
	{Name: "HypervisorArn", Flag: "hypervisor-arn", Type: "*string", Required: true},
}

var fields_get_virtual_machine = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_import_hypervisor_configuration = []leanruntime.Field{
	{Name: "Host", Flag: "host", Type: "*string", Required: true},
	{Name: "KmsKeyArn", Flag: "kms-key-arn", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Password", Flag: "password", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Username", Flag: "username", Type: "*string", Required: false},
}

var fields_list_gateways = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_hypervisors = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_virtual_machines = []leanruntime.Field{
	{Name: "HypervisorArn", Flag: "hypervisor-arn", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_put_bandwidth_rate_limit_schedule = []leanruntime.Field{
	{Name: "BandwidthRateLimitIntervals", Flag: "bandwidth-rate-limit-intervals", Type: "[]types.BandwidthRateLimitInterval", Required: true},
	{Name: "GatewayArn", Flag: "gateway-arn", Type: "*string", Required: true},
}

var fields_put_hypervisor_property_mappings = []leanruntime.Field{
	{Name: "HypervisorArn", Flag: "hypervisor-arn", Type: "*string", Required: true},
	{Name: "IamRoleArn", Flag: "iam-role-arn", Type: "*string", Required: true},
	{Name: "VmwareToAwsTagMappings", Flag: "vmware-to-aws-tag-mappings", Type: "[]types.VmwareToAwsTagMapping", Required: true},
}

var fields_put_maintenance_start_time = []leanruntime.Field{
	{Name: "DayOfMonth", Flag: "day-of-month", Type: "*int32", Required: false},
	{Name: "DayOfWeek", Flag: "day-of-week", Type: "*int32", Required: false},
	{Name: "GatewayArn", Flag: "gateway-arn", Type: "*string", Required: true},
	{Name: "HourOfDay", Flag: "hour-of-day", Type: "*int32", Required: true},
	{Name: "MinuteOfHour", Flag: "minute-of-hour", Type: "*int32", Required: true},
}

var fields_start_virtual_machines_metadata_sync = []leanruntime.Field{
	{Name: "HypervisorArn", Flag: "hypervisor-arn", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_test_hypervisor_configuration = []leanruntime.Field{
	{Name: "GatewayArn", Flag: "gateway-arn", Type: "*string", Required: true},
	{Name: "Host", Flag: "host", Type: "*string", Required: true},
	{Name: "Password", Flag: "password", Type: "*string", Required: false},
	{Name: "Username", Flag: "username", Type: "*string", Required: false},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_gateway_information = []leanruntime.Field{
	{Name: "GatewayArn", Flag: "gateway-arn", Type: "*string", Required: true},
	{Name: "GatewayDisplayName", Flag: "gateway-display-name", Type: "*string", Required: false},
}

var fields_update_gateway_software_now = []leanruntime.Field{
	{Name: "GatewayArn", Flag: "gateway-arn", Type: "*string", Required: true},
}

var fields_update_hypervisor = []leanruntime.Field{
	{Name: "Host", Flag: "host", Type: "*string", Required: false},
	{Name: "HypervisorArn", Flag: "hypervisor-arn", Type: "*string", Required: true},
	{Name: "LogGroupArn", Flag: "log-group-arn", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Password", Flag: "password", Type: "*string", Required: false},
	{Name: "Username", Flag: "username", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-gateway-to-server": {
			Name:   "associate-gateway-to-server",
			Fields: fields_associate_gateway_to_server,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateGatewayToServerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_gateway_to_server, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateGatewayToServer(ctx, input)
			},
		},
		"create-gateway": {
			Name:   "create-gateway",
			Fields: fields_create_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateGateway(ctx, input)
			},
		},
		"delete-gateway": {
			Name:   "delete-gateway",
			Fields: fields_delete_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteGateway(ctx, input)
			},
		},
		"delete-hypervisor": {
			Name:   "delete-hypervisor",
			Fields: fields_delete_hypervisor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteHypervisorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_hypervisor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteHypervisor(ctx, input)
			},
		},
		"disassociate-gateway-from-server": {
			Name:   "disassociate-gateway-from-server",
			Fields: fields_disassociate_gateway_from_server,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateGatewayFromServerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_gateway_from_server, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateGatewayFromServer(ctx, input)
			},
		},
		"get-bandwidth-rate-limit-schedule": {
			Name:   "get-bandwidth-rate-limit-schedule",
			Fields: fields_get_bandwidth_rate_limit_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBandwidthRateLimitScheduleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_bandwidth_rate_limit_schedule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBandwidthRateLimitSchedule(ctx, input)
			},
		},
		"get-gateway": {
			Name:   "get-gateway",
			Fields: fields_get_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetGateway(ctx, input)
			},
		},
		"get-hypervisor": {
			Name:   "get-hypervisor",
			Fields: fields_get_hypervisor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetHypervisorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_hypervisor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetHypervisor(ctx, input)
			},
		},
		"get-hypervisor-property-mappings": {
			Name:   "get-hypervisor-property-mappings",
			Fields: fields_get_hypervisor_property_mappings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetHypervisorPropertyMappingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_hypervisor_property_mappings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetHypervisorPropertyMappings(ctx, input)
			},
		},
		"get-virtual-machine": {
			Name:   "get-virtual-machine",
			Fields: fields_get_virtual_machine,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetVirtualMachineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_virtual_machine, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetVirtualMachine(ctx, input)
			},
		},
		"import-hypervisor-configuration": {
			Name:   "import-hypervisor-configuration",
			Fields: fields_import_hypervisor_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportHypervisorConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_hypervisor_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportHypervisorConfiguration(ctx, input)
			},
		},
		"list-gateways": {
			Name:   "list-gateways",
			Fields: fields_list_gateways,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGatewaysInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_gateways, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListGateways(ctx, input)
				}
				var results []*svc.ListGatewaysOutput
				p := svc.NewListGatewaysPaginator(client, input)
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
		"list-hypervisors": {
			Name:   "list-hypervisors",
			Fields: fields_list_hypervisors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListHypervisorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_hypervisors, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListHypervisors(ctx, input)
				}
				var results []*svc.ListHypervisorsOutput
				p := svc.NewListHypervisorsPaginator(client, input)
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
		"list-virtual-machines": {
			Name:   "list-virtual-machines",
			Fields: fields_list_virtual_machines,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVirtualMachinesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_virtual_machines, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListVirtualMachines(ctx, input)
				}
				var results []*svc.ListVirtualMachinesOutput
				p := svc.NewListVirtualMachinesPaginator(client, input)
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
		"put-bandwidth-rate-limit-schedule": {
			Name:   "put-bandwidth-rate-limit-schedule",
			Fields: fields_put_bandwidth_rate_limit_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutBandwidthRateLimitScheduleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_bandwidth_rate_limit_schedule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutBandwidthRateLimitSchedule(ctx, input)
			},
		},
		"put-hypervisor-property-mappings": {
			Name:   "put-hypervisor-property-mappings",
			Fields: fields_put_hypervisor_property_mappings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutHypervisorPropertyMappingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_hypervisor_property_mappings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutHypervisorPropertyMappings(ctx, input)
			},
		},
		"put-maintenance-start-time": {
			Name:   "put-maintenance-start-time",
			Fields: fields_put_maintenance_start_time,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutMaintenanceStartTimeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_maintenance_start_time, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutMaintenanceStartTime(ctx, input)
			},
		},
		"start-virtual-machines-metadata-sync": {
			Name:   "start-virtual-machines-metadata-sync",
			Fields: fields_start_virtual_machines_metadata_sync,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartVirtualMachinesMetadataSyncInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_virtual_machines_metadata_sync, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartVirtualMachinesMetadataSync(ctx, input)
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
		"test-hypervisor-configuration": {
			Name:   "test-hypervisor-configuration",
			Fields: fields_test_hypervisor_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TestHypervisorConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_test_hypervisor_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TestHypervisorConfiguration(ctx, input)
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
		"update-gateway-information": {
			Name:   "update-gateway-information",
			Fields: fields_update_gateway_information,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateGatewayInformationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_gateway_information, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateGatewayInformation(ctx, input)
			},
		},
		"update-gateway-software-now": {
			Name:   "update-gateway-software-now",
			Fields: fields_update_gateway_software_now,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateGatewaySoftwareNowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_gateway_software_now, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateGatewaySoftwareNow(ctx, input)
			},
		},
		"update-hypervisor": {
			Name:   "update-hypervisor",
			Fields: fields_update_hypervisor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateHypervisorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_hypervisor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateHypervisor(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("backupgateway", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
