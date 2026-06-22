package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/outposts"
)

var fields_cancel_capacity_task = []leanruntime.Field{
	{Name: "CapacityTaskId", Flag: "capacity-task-id", Type: "*string", Required: true},
	{Name: "OutpostIdentifier", Flag: "outpost-identifier", Type: "*string", Required: true},
}

var fields_cancel_order = []leanruntime.Field{
	{Name: "OrderId", Flag: "order-id", Type: "*string", Required: true},
}

var fields_create_order = []leanruntime.Field{
	{Name: "LineItems", Flag: "line-items", Type: "[]types.LineItemRequest", Required: false},
	{Name: "OutpostIdentifier", Flag: "outpost-identifier", Type: "*string", Required: true},
	{Name: "PaymentOption", Flag: "payment-option", Type: "types.PaymentOption", Required: true},
	{Name: "PaymentTerm", Flag: "payment-term", Type: "types.PaymentTerm", Required: false},
}

var fields_create_outpost = []leanruntime.Field{
	{Name: "AvailabilityZone", Flag: "availability-zone", Type: "*string", Required: false},
	{Name: "AvailabilityZoneId", Flag: "availability-zone-id", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "SiteId", Flag: "site-id", Type: "*string", Required: true},
	{Name: "SupportedHardwareType", Flag: "supported-hardware-type", Type: "types.SupportedHardwareType", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_site = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Notes", Flag: "notes", Type: "*string", Required: false},
	{Name: "OperatingAddress", Flag: "operating-address", Type: "*types.Address", Required: false},
	{Name: "RackPhysicalProperties", Flag: "rack-physical-properties", Type: "*types.RackPhysicalProperties", Required: false},
	{Name: "ShippingAddress", Flag: "shipping-address", Type: "*types.Address", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_outpost = []leanruntime.Field{
	{Name: "OutpostId", Flag: "outpost-id", Type: "*string", Required: true},
}

var fields_delete_site = []leanruntime.Field{
	{Name: "SiteId", Flag: "site-id", Type: "*string", Required: true},
}

var fields_get_capacity_task = []leanruntime.Field{
	{Name: "CapacityTaskId", Flag: "capacity-task-id", Type: "*string", Required: true},
	{Name: "OutpostIdentifier", Flag: "outpost-identifier", Type: "*string", Required: true},
}

var fields_get_catalog_item = []leanruntime.Field{
	{Name: "CatalogItemId", Flag: "catalog-item-id", Type: "*string", Required: true},
}

var fields_get_connection = []leanruntime.Field{
	{Name: "ConnectionId", Flag: "connection-id", Type: "*string", Required: true},
}

var fields_get_order = []leanruntime.Field{
	{Name: "OrderId", Flag: "order-id", Type: "*string", Required: true},
}

var fields_get_outpost = []leanruntime.Field{
	{Name: "OutpostId", Flag: "outpost-id", Type: "*string", Required: true},
}

var fields_get_outpost_billing_information = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OutpostIdentifier", Flag: "outpost-identifier", Type: "*string", Required: true},
}

var fields_get_outpost_instance_types = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OutpostId", Flag: "outpost-id", Type: "*string", Required: true},
}

var fields_get_outpost_supported_instance_types = []leanruntime.Field{
	{Name: "AssetId", Flag: "asset-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OrderId", Flag: "order-id", Type: "*string", Required: false},
	{Name: "OutpostIdentifier", Flag: "outpost-identifier", Type: "*string", Required: true},
}

var fields_get_site = []leanruntime.Field{
	{Name: "SiteId", Flag: "site-id", Type: "*string", Required: true},
}

var fields_get_site_address = []leanruntime.Field{
	{Name: "AddressType", Flag: "address-type", Type: "types.AddressType", Required: true},
	{Name: "SiteId", Flag: "site-id", Type: "*string", Required: true},
}

var fields_list_asset_instances = []leanruntime.Field{
	{Name: "AccountIdFilter", Flag: "account-id-filter", Type: "[]string", Required: false},
	{Name: "AssetIdFilter", Flag: "asset-id-filter", Type: "[]string", Required: false},
	{Name: "AwsServiceFilter", Flag: "aws-service-filter", Type: "[]types.AWSServiceName", Required: false},
	{Name: "InstanceTypeFilter", Flag: "instance-type-filter", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OutpostIdentifier", Flag: "outpost-identifier", Type: "*string", Required: true},
}

var fields_list_assets = []leanruntime.Field{
	{Name: "HostIdFilter", Flag: "host-id-filter", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OutpostIdentifier", Flag: "outpost-identifier", Type: "*string", Required: true},
	{Name: "StatusFilter", Flag: "status-filter", Type: "[]types.AssetState", Required: false},
}

var fields_list_blocking_instances_for_capacity_task = []leanruntime.Field{
	{Name: "CapacityTaskId", Flag: "capacity-task-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OutpostIdentifier", Flag: "outpost-identifier", Type: "*string", Required: true},
}

var fields_list_capacity_tasks = []leanruntime.Field{
	{Name: "CapacityTaskStatusFilter", Flag: "capacity-task-status-filter", Type: "[]types.CapacityTaskStatus", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OutpostIdentifierFilter", Flag: "outpost-identifier-filter", Type: "*string", Required: false},
}

var fields_list_catalog_items = []leanruntime.Field{
	{Name: "EC2FamilyFilter", Flag: "ec2-family-filter", Type: "[]string", Required: false},
	{Name: "ItemClassFilter", Flag: "item-class-filter", Type: "[]types.CatalogItemClass", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SupportedStorageFilter", Flag: "supported-storage-filter", Type: "[]types.SupportedStorageEnum", Required: false},
}

var fields_list_orders = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OutpostIdentifierFilter", Flag: "outpost-identifier-filter", Type: "*string", Required: false},
}

var fields_list_outposts = []leanruntime.Field{
	{Name: "AvailabilityZoneFilter", Flag: "availability-zone-filter", Type: "[]string", Required: false},
	{Name: "AvailabilityZoneIdFilter", Flag: "availability-zone-id-filter", Type: "[]string", Required: false},
	{Name: "LifeCycleStatusFilter", Flag: "life-cycle-status-filter", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_sites = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OperatingAddressCityFilter", Flag: "operating-address-city-filter", Type: "[]string", Required: false},
	{Name: "OperatingAddressCountryCodeFilter", Flag: "operating-address-country-code-filter", Type: "[]string", Required: false},
	{Name: "OperatingAddressStateOrRegionFilter", Flag: "operating-address-state-or-region-filter", Type: "[]string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_start_capacity_task = []leanruntime.Field{
	{Name: "AssetId", Flag: "asset-id", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "bool", Required: false},
	{Name: "InstancePools", Flag: "instance-pools", Type: "[]types.InstanceTypeCapacity", Required: true},
	{Name: "InstancesToExclude", Flag: "instances-to-exclude", Type: "*types.InstancesToExclude", Required: false},
	{Name: "OrderId", Flag: "order-id", Type: "*string", Required: false},
	{Name: "OutpostIdentifier", Flag: "outpost-identifier", Type: "*string", Required: true},
	{Name: "TaskActionOnBlockingInstances", Flag: "task-action-on-blocking-instances", Type: "types.TaskActionOnBlockingInstances", Required: false},
}

var fields_start_connection = []leanruntime.Field{
	{Name: "AssetId", Flag: "asset-id", Type: "*string", Required: true},
	{Name: "ClientPublicKey", Flag: "client-public-key", Type: "*string", Required: true},
	{Name: "DeviceSerialNumber", Flag: "device-serial-number", Type: "*string", Required: false},
	{Name: "NetworkInterfaceDeviceIndex", Flag: "network-interface-device-index", Type: "int32", Required: true},
}

var fields_start_outpost_decommission = []leanruntime.Field{
	{Name: "OutpostIdentifier", Flag: "outpost-identifier", Type: "*string", Required: true},
	{Name: "ValidateOnly", Flag: "validate-only", Type: "bool", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_outpost = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "OutpostId", Flag: "outpost-id", Type: "*string", Required: true},
	{Name: "SupportedHardwareType", Flag: "supported-hardware-type", Type: "types.SupportedHardwareType", Required: false},
}

var fields_update_site = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Notes", Flag: "notes", Type: "*string", Required: false},
	{Name: "SiteId", Flag: "site-id", Type: "*string", Required: true},
}

var fields_update_site_address = []leanruntime.Field{
	{Name: "Address", Flag: "address", Type: "*types.Address", Required: true},
	{Name: "AddressType", Flag: "address-type", Type: "types.AddressType", Required: true},
	{Name: "SiteId", Flag: "site-id", Type: "*string", Required: true},
}

var fields_update_site_rack_physical_properties = []leanruntime.Field{
	{Name: "FiberOpticCableType", Flag: "fiber-optic-cable-type", Type: "types.FiberOpticCableType", Required: false},
	{Name: "MaximumSupportedWeightLbs", Flag: "maximum-supported-weight-lbs", Type: "types.MaximumSupportedWeightLbs", Required: false},
	{Name: "OpticalStandard", Flag: "optical-standard", Type: "types.OpticalStandard", Required: false},
	{Name: "PowerConnector", Flag: "power-connector", Type: "types.PowerConnector", Required: false},
	{Name: "PowerDrawKva", Flag: "power-draw-kva", Type: "types.PowerDrawKva", Required: false},
	{Name: "PowerFeedDrop", Flag: "power-feed-drop", Type: "types.PowerFeedDrop", Required: false},
	{Name: "PowerPhase", Flag: "power-phase", Type: "types.PowerPhase", Required: false},
	{Name: "SiteId", Flag: "site-id", Type: "*string", Required: true},
	{Name: "UplinkCount", Flag: "uplink-count", Type: "types.UplinkCount", Required: false},
	{Name: "UplinkGbps", Flag: "uplink-gbps", Type: "types.UplinkGbps", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"cancel-capacity-task": {
			Name:   "cancel-capacity-task",
			Fields: fields_cancel_capacity_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelCapacityTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_capacity_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelCapacityTask(ctx, input)
			},
		},
		"cancel-order": {
			Name:   "cancel-order",
			Fields: fields_cancel_order,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelOrderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_order, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelOrder(ctx, input)
			},
		},
		"create-order": {
			Name:   "create-order",
			Fields: fields_create_order,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateOrderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_order, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateOrder(ctx, input)
			},
		},
		"create-outpost": {
			Name:   "create-outpost",
			Fields: fields_create_outpost,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateOutpostInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_outpost, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateOutpost(ctx, input)
			},
		},
		"create-site": {
			Name:   "create-site",
			Fields: fields_create_site,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSiteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_site, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSite(ctx, input)
			},
		},
		"delete-outpost": {
			Name:   "delete-outpost",
			Fields: fields_delete_outpost,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteOutpostInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_outpost, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteOutpost(ctx, input)
			},
		},
		"delete-site": {
			Name:   "delete-site",
			Fields: fields_delete_site,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSiteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_site, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSite(ctx, input)
			},
		},
		"get-capacity-task": {
			Name:   "get-capacity-task",
			Fields: fields_get_capacity_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCapacityTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_capacity_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCapacityTask(ctx, input)
			},
		},
		"get-catalog-item": {
			Name:   "get-catalog-item",
			Fields: fields_get_catalog_item,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCatalogItemInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_catalog_item, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCatalogItem(ctx, input)
			},
		},
		"get-connection": {
			Name:   "get-connection",
			Fields: fields_get_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConnection(ctx, input)
			},
		},
		"get-order": {
			Name:   "get-order",
			Fields: fields_get_order,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOrderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_order, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetOrder(ctx, input)
			},
		},
		"get-outpost": {
			Name:   "get-outpost",
			Fields: fields_get_outpost,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOutpostInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_outpost, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetOutpost(ctx, input)
			},
		},
		"get-outpost-billing-information": {
			Name:   "get-outpost-billing-information",
			Fields: fields_get_outpost_billing_information,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOutpostBillingInformationInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_outpost_billing_information, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetOutpostBillingInformation(ctx, input)
				}
				var results []*svc.GetOutpostBillingInformationOutput
				p := svc.NewGetOutpostBillingInformationPaginator(client, input)
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
		"get-outpost-instance-types": {
			Name:   "get-outpost-instance-types",
			Fields: fields_get_outpost_instance_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOutpostInstanceTypesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_outpost_instance_types, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetOutpostInstanceTypes(ctx, input)
				}
				var results []*svc.GetOutpostInstanceTypesOutput
				p := svc.NewGetOutpostInstanceTypesPaginator(client, input)
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
		"get-outpost-supported-instance-types": {
			Name:   "get-outpost-supported-instance-types",
			Fields: fields_get_outpost_supported_instance_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOutpostSupportedInstanceTypesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_outpost_supported_instance_types, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetOutpostSupportedInstanceTypes(ctx, input)
				}
				var results []*svc.GetOutpostSupportedInstanceTypesOutput
				p := svc.NewGetOutpostSupportedInstanceTypesPaginator(client, input)
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
		"get-site": {
			Name:   "get-site",
			Fields: fields_get_site,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSiteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_site, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSite(ctx, input)
			},
		},
		"get-site-address": {
			Name:   "get-site-address",
			Fields: fields_get_site_address,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSiteAddressInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_site_address, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSiteAddress(ctx, input)
			},
		},
		"list-asset-instances": {
			Name:   "list-asset-instances",
			Fields: fields_list_asset_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAssetInstancesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_asset_instances, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAssetInstances(ctx, input)
				}
				var results []*svc.ListAssetInstancesOutput
				p := svc.NewListAssetInstancesPaginator(client, input)
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
		"list-assets": {
			Name:   "list-assets",
			Fields: fields_list_assets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAssetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_assets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAssets(ctx, input)
				}
				var results []*svc.ListAssetsOutput
				p := svc.NewListAssetsPaginator(client, input)
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
		"list-blocking-instances-for-capacity-task": {
			Name:   "list-blocking-instances-for-capacity-task",
			Fields: fields_list_blocking_instances_for_capacity_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBlockingInstancesForCapacityTaskInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_blocking_instances_for_capacity_task, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBlockingInstancesForCapacityTask(ctx, input)
				}
				var results []*svc.ListBlockingInstancesForCapacityTaskOutput
				p := svc.NewListBlockingInstancesForCapacityTaskPaginator(client, input)
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
		"list-capacity-tasks": {
			Name:   "list-capacity-tasks",
			Fields: fields_list_capacity_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCapacityTasksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_capacity_tasks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCapacityTasks(ctx, input)
				}
				var results []*svc.ListCapacityTasksOutput
				p := svc.NewListCapacityTasksPaginator(client, input)
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
		"list-catalog-items": {
			Name:   "list-catalog-items",
			Fields: fields_list_catalog_items,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCatalogItemsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_catalog_items, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCatalogItems(ctx, input)
				}
				var results []*svc.ListCatalogItemsOutput
				p := svc.NewListCatalogItemsPaginator(client, input)
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
		"list-orders": {
			Name:   "list-orders",
			Fields: fields_list_orders,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOrdersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_orders, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListOrders(ctx, input)
				}
				var results []*svc.ListOrdersOutput
				p := svc.NewListOrdersPaginator(client, input)
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
		"list-outposts": {
			Name:   "list-outposts",
			Fields: fields_list_outposts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOutpostsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_outposts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListOutposts(ctx, input)
				}
				var results []*svc.ListOutpostsOutput
				p := svc.NewListOutpostsPaginator(client, input)
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
		"list-sites": {
			Name:   "list-sites",
			Fields: fields_list_sites,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSitesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_sites, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSites(ctx, input)
				}
				var results []*svc.ListSitesOutput
				p := svc.NewListSitesPaginator(client, input)
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
		"start-capacity-task": {
			Name:   "start-capacity-task",
			Fields: fields_start_capacity_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartCapacityTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_capacity_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartCapacityTask(ctx, input)
			},
		},
		"start-connection": {
			Name:   "start-connection",
			Fields: fields_start_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartConnection(ctx, input)
			},
		},
		"start-outpost-decommission": {
			Name:   "start-outpost-decommission",
			Fields: fields_start_outpost_decommission,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartOutpostDecommissionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_outpost_decommission, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartOutpostDecommission(ctx, input)
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
		"update-outpost": {
			Name:   "update-outpost",
			Fields: fields_update_outpost,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateOutpostInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_outpost, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateOutpost(ctx, input)
			},
		},
		"update-site": {
			Name:   "update-site",
			Fields: fields_update_site,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSiteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_site, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSite(ctx, input)
			},
		},
		"update-site-address": {
			Name:   "update-site-address",
			Fields: fields_update_site_address,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSiteAddressInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_site_address, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSiteAddress(ctx, input)
			},
		},
		"update-site-rack-physical-properties": {
			Name:   "update-site-rack-physical-properties",
			Fields: fields_update_site_rack_physical_properties,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSiteRackPhysicalPropertiesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_site_rack_physical_properties, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSiteRackPhysicalProperties(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("outposts", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
