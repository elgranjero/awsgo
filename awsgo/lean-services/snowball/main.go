package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/snowball"
)

var fields_cancel_cluster = []leanruntime.Field{
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: true},
}

var fields_cancel_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_create_address = []leanruntime.Field{
	{Name: "Address", Flag: "address", Type: "*types.Address", Required: true},
}

var fields_create_cluster = []leanruntime.Field{
	{Name: "AddressId", Flag: "address-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ForceCreateJobs", Flag: "force-create-jobs", Type: "bool", Required: false},
	{Name: "ForwardingAddressId", Flag: "forwarding-address-id", Type: "*string", Required: false},
	{Name: "InitialClusterSize", Flag: "initial-cluster-size", Type: "*int32", Required: false},
	{Name: "JobType", Flag: "job-type", Type: "types.JobType", Required: true},
	{Name: "KmsKeyARN", Flag: "kms-key-arn", Type: "*string", Required: false},
	{Name: "LongTermPricingIds", Flag: "long-term-pricing-ids", Type: "[]string", Required: false},
	{Name: "Notification", Flag: "notification", Type: "*types.Notification", Required: false},
	{Name: "OnDeviceServiceConfiguration", Flag: "on-device-service-configuration", Type: "*types.OnDeviceServiceConfiguration", Required: false},
	{Name: "RemoteManagement", Flag: "remote-management", Type: "types.RemoteManagement", Required: false},
	{Name: "Resources", Flag: "resources", Type: "*types.JobResource", Required: false},
	{Name: "RoleARN", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "ShippingOption", Flag: "shipping-option", Type: "types.ShippingOption", Required: true},
	{Name: "SnowballCapacityPreference", Flag: "snowball-capacity-preference", Type: "types.SnowballCapacity", Required: false},
	{Name: "SnowballType", Flag: "snowball-type", Type: "types.SnowballType", Required: true},
	{Name: "TaxDocuments", Flag: "tax-documents", Type: "*types.TaxDocuments", Required: false},
}

var fields_create_job = []leanruntime.Field{
	{Name: "AddressId", Flag: "address-id", Type: "*string", Required: false},
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DeviceConfiguration", Flag: "device-configuration", Type: "*types.DeviceConfiguration", Required: false},
	{Name: "ForwardingAddressId", Flag: "forwarding-address-id", Type: "*string", Required: false},
	{Name: "ImpactLevel", Flag: "impact-level", Type: "types.ImpactLevel", Required: false},
	{Name: "JobType", Flag: "job-type", Type: "types.JobType", Required: false},
	{Name: "KmsKeyARN", Flag: "kms-key-arn", Type: "*string", Required: false},
	{Name: "LongTermPricingId", Flag: "long-term-pricing-id", Type: "*string", Required: false},
	{Name: "Notification", Flag: "notification", Type: "*types.Notification", Required: false},
	{Name: "OnDeviceServiceConfiguration", Flag: "on-device-service-configuration", Type: "*types.OnDeviceServiceConfiguration", Required: false},
	{Name: "PickupDetails", Flag: "pickup-details", Type: "*types.PickupDetails", Required: false},
	{Name: "RemoteManagement", Flag: "remote-management", Type: "types.RemoteManagement", Required: false},
	{Name: "Resources", Flag: "resources", Type: "*types.JobResource", Required: false},
	{Name: "RoleARN", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "ShippingOption", Flag: "shipping-option", Type: "types.ShippingOption", Required: false},
	{Name: "SnowballCapacityPreference", Flag: "snowball-capacity-preference", Type: "types.SnowballCapacity", Required: false},
	{Name: "SnowballType", Flag: "snowball-type", Type: "types.SnowballType", Required: false},
	{Name: "TaxDocuments", Flag: "tax-documents", Type: "*types.TaxDocuments", Required: false},
}

var fields_create_long_term_pricing = []leanruntime.Field{
	{Name: "IsLongTermPricingAutoRenew", Flag: "is-long-term-pricing-auto-renew", Type: "*bool", Required: false},
	{Name: "LongTermPricingType", Flag: "long-term-pricing-type", Type: "types.LongTermPricingType", Required: true},
	{Name: "SnowballType", Flag: "snowball-type", Type: "types.SnowballType", Required: true},
}

var fields_create_return_shipping_label = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "ShippingOption", Flag: "shipping-option", Type: "types.ShippingOption", Required: false},
}

var fields_describe_address = []leanruntime.Field{
	{Name: "AddressId", Flag: "address-id", Type: "*string", Required: true},
}

var fields_describe_addresses = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_cluster = []leanruntime.Field{
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: true},
}

var fields_describe_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_describe_return_shipping_label = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_get_job_manifest = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_get_job_unlock_code = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_get_snowball_usage = []leanruntime.Field{}

var fields_get_software_updates = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_list_cluster_jobs = []leanruntime.Field{
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_clusters = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_compatible_images = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_jobs = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_long_term_pricing = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_pickup_locations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_service_versions = []leanruntime.Field{
	{Name: "DependentServices", Flag: "dependent-services", Type: "[]types.DependentService", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ServiceName", Flag: "service-name", Type: "types.ServiceName", Required: true},
}

var fields_update_cluster = []leanruntime.Field{
	{Name: "AddressId", Flag: "address-id", Type: "*string", Required: false},
	{Name: "ClusterId", Flag: "cluster-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ForwardingAddressId", Flag: "forwarding-address-id", Type: "*string", Required: false},
	{Name: "Notification", Flag: "notification", Type: "*types.Notification", Required: false},
	{Name: "OnDeviceServiceConfiguration", Flag: "on-device-service-configuration", Type: "*types.OnDeviceServiceConfiguration", Required: false},
	{Name: "Resources", Flag: "resources", Type: "*types.JobResource", Required: false},
	{Name: "RoleARN", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "ShippingOption", Flag: "shipping-option", Type: "types.ShippingOption", Required: false},
}

var fields_update_job = []leanruntime.Field{
	{Name: "AddressId", Flag: "address-id", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ForwardingAddressId", Flag: "forwarding-address-id", Type: "*string", Required: false},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "Notification", Flag: "notification", Type: "*types.Notification", Required: false},
	{Name: "OnDeviceServiceConfiguration", Flag: "on-device-service-configuration", Type: "*types.OnDeviceServiceConfiguration", Required: false},
	{Name: "PickupDetails", Flag: "pickup-details", Type: "*types.PickupDetails", Required: false},
	{Name: "Resources", Flag: "resources", Type: "*types.JobResource", Required: false},
	{Name: "RoleARN", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "ShippingOption", Flag: "shipping-option", Type: "types.ShippingOption", Required: false},
	{Name: "SnowballCapacityPreference", Flag: "snowball-capacity-preference", Type: "types.SnowballCapacity", Required: false},
}

var fields_update_job_shipment_state = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "ShipmentState", Flag: "shipment-state", Type: "types.ShipmentState", Required: true},
}

var fields_update_long_term_pricing = []leanruntime.Field{
	{Name: "IsLongTermPricingAutoRenew", Flag: "is-long-term-pricing-auto-renew", Type: "*bool", Required: false},
	{Name: "LongTermPricingId", Flag: "long-term-pricing-id", Type: "*string", Required: true},
	{Name: "ReplacementJob", Flag: "replacement-job", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"cancel-cluster": {
			Name:   "cancel-cluster",
			Fields: fields_cancel_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelCluster(ctx, input)
			},
		},
		"cancel-job": {
			Name:   "cancel-job",
			Fields: fields_cancel_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelJob(ctx, input)
			},
		},
		"create-address": {
			Name:   "create-address",
			Fields: fields_create_address,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAddressInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_address, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAddress(ctx, input)
			},
		},
		"create-cluster": {
			Name:   "create-cluster",
			Fields: fields_create_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCluster(ctx, input)
			},
		},
		"create-job": {
			Name:   "create-job",
			Fields: fields_create_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateJob(ctx, input)
			},
		},
		"create-long-term-pricing": {
			Name:   "create-long-term-pricing",
			Fields: fields_create_long_term_pricing,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLongTermPricingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_long_term_pricing, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLongTermPricing(ctx, input)
			},
		},
		"create-return-shipping-label": {
			Name:   "create-return-shipping-label",
			Fields: fields_create_return_shipping_label,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateReturnShippingLabelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_return_shipping_label, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateReturnShippingLabel(ctx, input)
			},
		},
		"describe-address": {
			Name:   "describe-address",
			Fields: fields_describe_address,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAddressInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_address, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAddress(ctx, input)
			},
		},
		"describe-addresses": {
			Name:   "describe-addresses",
			Fields: fields_describe_addresses,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAddressesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_addresses, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeAddresses(ctx, input)
				}
				var results []*svc.DescribeAddressesOutput
				p := svc.NewDescribeAddressesPaginator(client, input)
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
		"describe-cluster": {
			Name:   "describe-cluster",
			Fields: fields_describe_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeCluster(ctx, input)
			},
		},
		"describe-job": {
			Name:   "describe-job",
			Fields: fields_describe_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeJob(ctx, input)
			},
		},
		"describe-return-shipping-label": {
			Name:   "describe-return-shipping-label",
			Fields: fields_describe_return_shipping_label,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeReturnShippingLabelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_return_shipping_label, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeReturnShippingLabel(ctx, input)
			},
		},
		"get-job-manifest": {
			Name:   "get-job-manifest",
			Fields: fields_get_job_manifest,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetJobManifestInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_job_manifest, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetJobManifest(ctx, input)
			},
		},
		"get-job-unlock-code": {
			Name:   "get-job-unlock-code",
			Fields: fields_get_job_unlock_code,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetJobUnlockCodeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_job_unlock_code, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetJobUnlockCode(ctx, input)
			},
		},
		"get-snowball-usage": {
			Name:   "get-snowball-usage",
			Fields: fields_get_snowball_usage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSnowballUsageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_snowball_usage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSnowballUsage(ctx, input)
			},
		},
		"get-software-updates": {
			Name:   "get-software-updates",
			Fields: fields_get_software_updates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSoftwareUpdatesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_software_updates, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSoftwareUpdates(ctx, input)
			},
		},
		"list-cluster-jobs": {
			Name:   "list-cluster-jobs",
			Fields: fields_list_cluster_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListClusterJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_cluster_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListClusterJobs(ctx, input)
				}
				var results []*svc.ListClusterJobsOutput
				p := svc.NewListClusterJobsPaginator(client, input)
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
		"list-clusters": {
			Name:   "list-clusters",
			Fields: fields_list_clusters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListClustersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_clusters, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListClusters(ctx, input)
				}
				var results []*svc.ListClustersOutput
				p := svc.NewListClustersPaginator(client, input)
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
		"list-compatible-images": {
			Name:   "list-compatible-images",
			Fields: fields_list_compatible_images,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCompatibleImagesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_compatible_images, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCompatibleImages(ctx, input)
				}
				var results []*svc.ListCompatibleImagesOutput
				p := svc.NewListCompatibleImagesPaginator(client, input)
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
		"list-jobs": {
			Name:   "list-jobs",
			Fields: fields_list_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListJobs(ctx, input)
				}
				var results []*svc.ListJobsOutput
				p := svc.NewListJobsPaginator(client, input)
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
		"list-long-term-pricing": {
			Name:   "list-long-term-pricing",
			Fields: fields_list_long_term_pricing,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLongTermPricingInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_long_term_pricing, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLongTermPricing(ctx, input)
				}
				var results []*svc.ListLongTermPricingOutput
				p := svc.NewListLongTermPricingPaginator(client, input)
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
		"list-pickup-locations": {
			Name:   "list-pickup-locations",
			Fields: fields_list_pickup_locations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPickupLocationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_pickup_locations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPickupLocations(ctx, input)
				}
				var results []*svc.ListPickupLocationsOutput
				p := svc.NewListPickupLocationsPaginator(client, input)
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
		"list-service-versions": {
			Name:   "list-service-versions",
			Fields: fields_list_service_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListServiceVersionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_service_versions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListServiceVersions(ctx, input)
			},
		},
		"update-cluster": {
			Name:   "update-cluster",
			Fields: fields_update_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCluster(ctx, input)
			},
		},
		"update-job": {
			Name:   "update-job",
			Fields: fields_update_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateJob(ctx, input)
			},
		},
		"update-job-shipment-state": {
			Name:   "update-job-shipment-state",
			Fields: fields_update_job_shipment_state,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateJobShipmentStateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_job_shipment_state, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateJobShipmentState(ctx, input)
			},
		},
		"update-long-term-pricing": {
			Name:   "update-long-term-pricing",
			Fields: fields_update_long_term_pricing,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateLongTermPricingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_long_term_pricing, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateLongTermPricing(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("snowball", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
