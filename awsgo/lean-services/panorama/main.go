package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/panorama"
)

var fields_create_application_instance = []leanruntime.Field{
	{Name: "ApplicationInstanceIdToReplace", Flag: "application-instance-id-to-replace", Type: "*string", Required: false},
	{Name: "DefaultRuntimeContextDevice", Flag: "default-runtime-context-device", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ManifestOverridesPayload", Flag: "manifest-overrides-payload", Type: "types.ManifestOverridesPayload", Required: false},
	{Name: "ManifestPayload", Flag: "manifest-payload", Type: "types.ManifestPayload", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "RuntimeRoleArn", Flag: "runtime-role-arn", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_job_for_devices = []leanruntime.Field{
	{Name: "DeviceIds", Flag: "device-ids", Type: "[]string", Required: true},
	{Name: "DeviceJobConfig", Flag: "device-job-config", Type: "*types.DeviceJobConfig", Required: false},
	{Name: "JobType", Flag: "job-type", Type: "types.JobType", Required: true},
}

var fields_create_node_from_template_job = []leanruntime.Field{
	{Name: "JobTags", Flag: "job-tags", Type: "[]types.JobResourceTags", Required: false},
	{Name: "NodeDescription", Flag: "node-description", Type: "*string", Required: false},
	{Name: "NodeName", Flag: "node-name", Type: "*string", Required: true},
	{Name: "OutputPackageName", Flag: "output-package-name", Type: "*string", Required: true},
	{Name: "OutputPackageVersion", Flag: "output-package-version", Type: "*string", Required: true},
	{Name: "TemplateParameters", Flag: "template-parameters", Type: "map[string]string", Required: true},
	{Name: "TemplateType", Flag: "template-type", Type: "types.TemplateType", Required: true},
}

var fields_create_package = []leanruntime.Field{
	{Name: "PackageName", Flag: "package-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_package_import_job = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "InputConfig", Flag: "input-config", Type: "*types.PackageImportJobInputConfig", Required: true},
	{Name: "JobTags", Flag: "job-tags", Type: "[]types.JobResourceTags", Required: false},
	{Name: "JobType", Flag: "job-type", Type: "types.PackageImportJobType", Required: true},
	{Name: "OutputConfig", Flag: "output-config", Type: "*types.PackageImportJobOutputConfig", Required: true},
}

var fields_delete_device = []leanruntime.Field{
	{Name: "DeviceId", Flag: "device-id", Type: "*string", Required: true},
}

var fields_delete_package = []leanruntime.Field{
	{Name: "ForceDelete", Flag: "force-delete", Type: "bool", Required: false},
	{Name: "PackageId", Flag: "package-id", Type: "*string", Required: true},
}

var fields_deregister_package_version = []leanruntime.Field{
	{Name: "OwnerAccount", Flag: "owner-account", Type: "*string", Required: false},
	{Name: "PackageId", Flag: "package-id", Type: "*string", Required: true},
	{Name: "PackageVersion", Flag: "package-version", Type: "*string", Required: true},
	{Name: "PatchVersion", Flag: "patch-version", Type: "*string", Required: true},
	{Name: "UpdatedLatestPatchVersion", Flag: "updated-latest-patch-version", Type: "*string", Required: false},
}

var fields_describe_application_instance = []leanruntime.Field{
	{Name: "ApplicationInstanceId", Flag: "application-instance-id", Type: "*string", Required: true},
}

var fields_describe_application_instance_details = []leanruntime.Field{
	{Name: "ApplicationInstanceId", Flag: "application-instance-id", Type: "*string", Required: true},
}

var fields_describe_device = []leanruntime.Field{
	{Name: "DeviceId", Flag: "device-id", Type: "*string", Required: true},
}

var fields_describe_device_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_describe_node = []leanruntime.Field{
	{Name: "NodeId", Flag: "node-id", Type: "*string", Required: true},
	{Name: "OwnerAccount", Flag: "owner-account", Type: "*string", Required: false},
}

var fields_describe_node_from_template_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_describe_package = []leanruntime.Field{
	{Name: "PackageId", Flag: "package-id", Type: "*string", Required: true},
}

var fields_describe_package_import_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_describe_package_version = []leanruntime.Field{
	{Name: "OwnerAccount", Flag: "owner-account", Type: "*string", Required: false},
	{Name: "PackageId", Flag: "package-id", Type: "*string", Required: true},
	{Name: "PackageVersion", Flag: "package-version", Type: "*string", Required: true},
	{Name: "PatchVersion", Flag: "patch-version", Type: "*string", Required: false},
}

var fields_list_application_instance_dependencies = []leanruntime.Field{
	{Name: "ApplicationInstanceId", Flag: "application-instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_application_instance_node_instances = []leanruntime.Field{
	{Name: "ApplicationInstanceId", Flag: "application-instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_application_instances = []leanruntime.Field{
	{Name: "DeviceId", Flag: "device-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StatusFilter", Flag: "status-filter", Type: "types.StatusFilter", Required: false},
}

var fields_list_devices = []leanruntime.Field{
	{Name: "DeviceAggregatedStatusFilter", Flag: "device-aggregated-status-filter", Type: "types.DeviceAggregatedStatus", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NameFilter", Flag: "name-filter", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.ListDevicesSortBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_list_devices_jobs = []leanruntime.Field{
	{Name: "DeviceId", Flag: "device-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_node_from_template_jobs = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_nodes = []leanruntime.Field{
	{Name: "Category", Flag: "category", Type: "types.NodeCategory", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OwnerAccount", Flag: "owner-account", Type: "*string", Required: false},
	{Name: "PackageName", Flag: "package-name", Type: "*string", Required: false},
	{Name: "PackageVersion", Flag: "package-version", Type: "*string", Required: false},
	{Name: "PatchVersion", Flag: "patch-version", Type: "*string", Required: false},
}

var fields_list_package_import_jobs = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_packages = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_provision_device = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NetworkingConfiguration", Flag: "networking-configuration", Type: "*types.NetworkPayload", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_register_package_version = []leanruntime.Field{
	{Name: "MarkLatest", Flag: "mark-latest", Type: "bool", Required: false},
	{Name: "OwnerAccount", Flag: "owner-account", Type: "*string", Required: false},
	{Name: "PackageId", Flag: "package-id", Type: "*string", Required: true},
	{Name: "PackageVersion", Flag: "package-version", Type: "*string", Required: true},
	{Name: "PatchVersion", Flag: "patch-version", Type: "*string", Required: true},
}

var fields_remove_application_instance = []leanruntime.Field{
	{Name: "ApplicationInstanceId", Flag: "application-instance-id", Type: "*string", Required: true},
}

var fields_signal_application_instance_node_instances = []leanruntime.Field{
	{Name: "ApplicationInstanceId", Flag: "application-instance-id", Type: "*string", Required: true},
	{Name: "NodeSignals", Flag: "node-signals", Type: "[]types.NodeSignal", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_device_metadata = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DeviceId", Flag: "device-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-application-instance": {
			Name:   "create-application-instance",
			Fields: fields_create_application_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateApplicationInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_application_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateApplicationInstance(ctx, input)
			},
		},
		"create-job-for-devices": {
			Name:   "create-job-for-devices",
			Fields: fields_create_job_for_devices,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateJobForDevicesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_job_for_devices, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateJobForDevices(ctx, input)
			},
		},
		"create-node-from-template-job": {
			Name:   "create-node-from-template-job",
			Fields: fields_create_node_from_template_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateNodeFromTemplateJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_node_from_template_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateNodeFromTemplateJob(ctx, input)
			},
		},
		"create-package": {
			Name:   "create-package",
			Fields: fields_create_package,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePackageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_package, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePackage(ctx, input)
			},
		},
		"create-package-import-job": {
			Name:   "create-package-import-job",
			Fields: fields_create_package_import_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePackageImportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_package_import_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePackageImportJob(ctx, input)
			},
		},
		"delete-device": {
			Name:   "delete-device",
			Fields: fields_delete_device,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDeviceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_device, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDevice(ctx, input)
			},
		},
		"delete-package": {
			Name:   "delete-package",
			Fields: fields_delete_package,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePackageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_package, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePackage(ctx, input)
			},
		},
		"deregister-package-version": {
			Name:   "deregister-package-version",
			Fields: fields_deregister_package_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterPackageVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_package_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterPackageVersion(ctx, input)
			},
		},
		"describe-application-instance": {
			Name:   "describe-application-instance",
			Fields: fields_describe_application_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeApplicationInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_application_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeApplicationInstance(ctx, input)
			},
		},
		"describe-application-instance-details": {
			Name:   "describe-application-instance-details",
			Fields: fields_describe_application_instance_details,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeApplicationInstanceDetailsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_application_instance_details, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeApplicationInstanceDetails(ctx, input)
			},
		},
		"describe-device": {
			Name:   "describe-device",
			Fields: fields_describe_device,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDeviceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_device, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDevice(ctx, input)
			},
		},
		"describe-device-job": {
			Name:   "describe-device-job",
			Fields: fields_describe_device_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDeviceJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_device_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDeviceJob(ctx, input)
			},
		},
		"describe-node": {
			Name:   "describe-node",
			Fields: fields_describe_node,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeNodeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_node, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeNode(ctx, input)
			},
		},
		"describe-node-from-template-job": {
			Name:   "describe-node-from-template-job",
			Fields: fields_describe_node_from_template_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeNodeFromTemplateJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_node_from_template_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeNodeFromTemplateJob(ctx, input)
			},
		},
		"describe-package": {
			Name:   "describe-package",
			Fields: fields_describe_package,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePackageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_package, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribePackage(ctx, input)
			},
		},
		"describe-package-import-job": {
			Name:   "describe-package-import-job",
			Fields: fields_describe_package_import_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePackageImportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_package_import_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribePackageImportJob(ctx, input)
			},
		},
		"describe-package-version": {
			Name:   "describe-package-version",
			Fields: fields_describe_package_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePackageVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_package_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribePackageVersion(ctx, input)
			},
		},
		"list-application-instance-dependencies": {
			Name:   "list-application-instance-dependencies",
			Fields: fields_list_application_instance_dependencies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListApplicationInstanceDependenciesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_application_instance_dependencies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListApplicationInstanceDependencies(ctx, input)
				}
				var results []*svc.ListApplicationInstanceDependenciesOutput
				p := svc.NewListApplicationInstanceDependenciesPaginator(client, input)
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
		"list-application-instance-node-instances": {
			Name:   "list-application-instance-node-instances",
			Fields: fields_list_application_instance_node_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListApplicationInstanceNodeInstancesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_application_instance_node_instances, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListApplicationInstanceNodeInstances(ctx, input)
				}
				var results []*svc.ListApplicationInstanceNodeInstancesOutput
				p := svc.NewListApplicationInstanceNodeInstancesPaginator(client, input)
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
		"list-application-instances": {
			Name:   "list-application-instances",
			Fields: fields_list_application_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListApplicationInstancesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_application_instances, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListApplicationInstances(ctx, input)
				}
				var results []*svc.ListApplicationInstancesOutput
				p := svc.NewListApplicationInstancesPaginator(client, input)
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
		"list-devices": {
			Name:   "list-devices",
			Fields: fields_list_devices,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDevicesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_devices, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDevices(ctx, input)
				}
				var results []*svc.ListDevicesOutput
				p := svc.NewListDevicesPaginator(client, input)
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
		"list-devices-jobs": {
			Name:   "list-devices-jobs",
			Fields: fields_list_devices_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDevicesJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_devices_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDevicesJobs(ctx, input)
				}
				var results []*svc.ListDevicesJobsOutput
				p := svc.NewListDevicesJobsPaginator(client, input)
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
		"list-node-from-template-jobs": {
			Name:   "list-node-from-template-jobs",
			Fields: fields_list_node_from_template_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListNodeFromTemplateJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_node_from_template_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListNodeFromTemplateJobs(ctx, input)
				}
				var results []*svc.ListNodeFromTemplateJobsOutput
				p := svc.NewListNodeFromTemplateJobsPaginator(client, input)
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
		"list-nodes": {
			Name:   "list-nodes",
			Fields: fields_list_nodes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListNodesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_nodes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListNodes(ctx, input)
				}
				var results []*svc.ListNodesOutput
				p := svc.NewListNodesPaginator(client, input)
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
		"list-package-import-jobs": {
			Name:   "list-package-import-jobs",
			Fields: fields_list_package_import_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPackageImportJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_package_import_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPackageImportJobs(ctx, input)
				}
				var results []*svc.ListPackageImportJobsOutput
				p := svc.NewListPackageImportJobsPaginator(client, input)
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
		"list-packages": {
			Name:   "list-packages",
			Fields: fields_list_packages,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPackagesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_packages, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPackages(ctx, input)
				}
				var results []*svc.ListPackagesOutput
				p := svc.NewListPackagesPaginator(client, input)
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
		"provision-device": {
			Name:   "provision-device",
			Fields: fields_provision_device,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ProvisionDeviceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_provision_device, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ProvisionDevice(ctx, input)
			},
		},
		"register-package-version": {
			Name:   "register-package-version",
			Fields: fields_register_package_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterPackageVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_package_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterPackageVersion(ctx, input)
			},
		},
		"remove-application-instance": {
			Name:   "remove-application-instance",
			Fields: fields_remove_application_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveApplicationInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_application_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveApplicationInstance(ctx, input)
			},
		},
		"signal-application-instance-node-instances": {
			Name:   "signal-application-instance-node-instances",
			Fields: fields_signal_application_instance_node_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SignalApplicationInstanceNodeInstancesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_signal_application_instance_node_instances, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SignalApplicationInstanceNodeInstances(ctx, input)
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
		"update-device-metadata": {
			Name:   "update-device-metadata",
			Fields: fields_update_device_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDeviceMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_device_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDeviceMetadata(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("panorama", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
