package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/greengrassv2"
)

var fields_associate_service_role_to_account = []leanruntime.Field{
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
}

var fields_batch_associate_client_device_with_core_device = []leanruntime.Field{
	{Name: "CoreDeviceThingName", Flag: "core-device-thing-name", Type: "*string", Required: true},
	{Name: "Entries", Flag: "entries", Type: "[]types.AssociateClientDeviceWithCoreDeviceEntry", Required: false},
}

var fields_batch_disassociate_client_device_from_core_device = []leanruntime.Field{
	{Name: "CoreDeviceThingName", Flag: "core-device-thing-name", Type: "*string", Required: true},
	{Name: "Entries", Flag: "entries", Type: "[]types.DisassociateClientDeviceFromCoreDeviceEntry", Required: false},
}

var fields_cancel_deployment = []leanruntime.Field{
	{Name: "DeploymentId", Flag: "deployment-id", Type: "*string", Required: true},
}

var fields_create_component_version = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "InlineRecipe", Flag: "inline-recipe", Type: "[]byte", Required: false},
	{Name: "LambdaFunction", Flag: "lambda-function", Type: "*types.LambdaFunctionRecipeSource", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_deployment = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Components", Flag: "components", Type: "map[string]types.ComponentDeploymentSpecification", Required: false},
	{Name: "DeploymentName", Flag: "deployment-name", Type: "*string", Required: false},
	{Name: "DeploymentPolicies", Flag: "deployment-policies", Type: "*types.DeploymentPolicies", Required: false},
	{Name: "IotJobConfiguration", Flag: "iot-job-configuration", Type: "*types.DeploymentIoTJobConfiguration", Required: false},
	{Name: "ParentTargetArn", Flag: "parent-target-arn", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TargetArn", Flag: "target-arn", Type: "*string", Required: true},
}

var fields_delete_component = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_delete_core_device = []leanruntime.Field{
	{Name: "CoreDeviceThingName", Flag: "core-device-thing-name", Type: "*string", Required: true},
}

var fields_delete_deployment = []leanruntime.Field{
	{Name: "DeploymentId", Flag: "deployment-id", Type: "*string", Required: true},
}

var fields_describe_component = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_disassociate_service_role_from_account = []leanruntime.Field{}

var fields_get_component = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "RecipeOutputFormat", Flag: "recipe-output-format", Type: "types.RecipeOutputFormat", Required: false},
}

var fields_get_component_version_artifact = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "ArtifactName", Flag: "artifact-name", Type: "*string", Required: true},
	{Name: "IotEndpointType", Flag: "iot-endpoint-type", Type: "types.IotEndpointType", Required: false},
	{Name: "S3EndpointType", Flag: "s3-endpoint-type", Type: "types.S3EndpointType", Required: false},
}

var fields_get_connectivity_info = []leanruntime.Field{
	{Name: "ThingName", Flag: "thing-name", Type: "*string", Required: true},
}

var fields_get_core_device = []leanruntime.Field{
	{Name: "CoreDeviceThingName", Flag: "core-device-thing-name", Type: "*string", Required: true},
}

var fields_get_deployment = []leanruntime.Field{
	{Name: "DeploymentId", Flag: "deployment-id", Type: "*string", Required: true},
}

var fields_get_service_role_for_account = []leanruntime.Field{}

var fields_list_client_devices_associated_with_core_device = []leanruntime.Field{
	{Name: "CoreDeviceThingName", Flag: "core-device-thing-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_component_versions = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_components = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Scope", Flag: "scope", Type: "types.ComponentVisibilityScope", Required: false},
}

var fields_list_core_devices = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Runtime", Flag: "runtime", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.CoreDeviceStatus", Required: false},
	{Name: "ThingGroupArn", Flag: "thing-group-arn", Type: "*string", Required: false},
}

var fields_list_deployments = []leanruntime.Field{
	{Name: "HistoryFilter", Flag: "history-filter", Type: "types.DeploymentHistoryFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ParentTargetArn", Flag: "parent-target-arn", Type: "*string", Required: false},
	{Name: "TargetArn", Flag: "target-arn", Type: "*string", Required: false},
}

var fields_list_effective_deployments = []leanruntime.Field{
	{Name: "CoreDeviceThingName", Flag: "core-device-thing-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_installed_components = []leanruntime.Field{
	{Name: "CoreDeviceThingName", Flag: "core-device-thing-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TopologyFilter", Flag: "topology-filter", Type: "types.InstalledComponentTopologyFilter", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_resolve_component_candidates = []leanruntime.Field{
	{Name: "ComponentCandidates", Flag: "component-candidates", Type: "[]types.ComponentCandidate", Required: false},
	{Name: "Platform", Flag: "platform", Type: "*types.ComponentPlatform", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_connectivity_info = []leanruntime.Field{
	{Name: "ConnectivityInfo", Flag: "connectivity-info", Type: "[]types.ConnectivityInfo", Required: true},
	{Name: "ThingName", Flag: "thing-name", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-service-role-to-account": {
			Name:   "associate-service-role-to-account",
			Fields: fields_associate_service_role_to_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateServiceRoleToAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_service_role_to_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateServiceRoleToAccount(ctx, input)
			},
		},
		"batch-associate-client-device-with-core-device": {
			Name:   "batch-associate-client-device-with-core-device",
			Fields: fields_batch_associate_client_device_with_core_device,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchAssociateClientDeviceWithCoreDeviceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_associate_client_device_with_core_device, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchAssociateClientDeviceWithCoreDevice(ctx, input)
			},
		},
		"batch-disassociate-client-device-from-core-device": {
			Name:   "batch-disassociate-client-device-from-core-device",
			Fields: fields_batch_disassociate_client_device_from_core_device,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDisassociateClientDeviceFromCoreDeviceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_disassociate_client_device_from_core_device, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDisassociateClientDeviceFromCoreDevice(ctx, input)
			},
		},
		"cancel-deployment": {
			Name:   "cancel-deployment",
			Fields: fields_cancel_deployment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelDeploymentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_deployment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelDeployment(ctx, input)
			},
		},
		"create-component-version": {
			Name:   "create-component-version",
			Fields: fields_create_component_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateComponentVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_component_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateComponentVersion(ctx, input)
			},
		},
		"create-deployment": {
			Name:   "create-deployment",
			Fields: fields_create_deployment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDeploymentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_deployment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDeployment(ctx, input)
			},
		},
		"delete-component": {
			Name:   "delete-component",
			Fields: fields_delete_component,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteComponentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_component, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteComponent(ctx, input)
			},
		},
		"delete-core-device": {
			Name:   "delete-core-device",
			Fields: fields_delete_core_device,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCoreDeviceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_core_device, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCoreDevice(ctx, input)
			},
		},
		"delete-deployment": {
			Name:   "delete-deployment",
			Fields: fields_delete_deployment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDeploymentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_deployment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDeployment(ctx, input)
			},
		},
		"describe-component": {
			Name:   "describe-component",
			Fields: fields_describe_component,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeComponentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_component, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeComponent(ctx, input)
			},
		},
		"disassociate-service-role-from-account": {
			Name:   "disassociate-service-role-from-account",
			Fields: fields_disassociate_service_role_from_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateServiceRoleFromAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_service_role_from_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateServiceRoleFromAccount(ctx, input)
			},
		},
		"get-component": {
			Name:   "get-component",
			Fields: fields_get_component,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetComponentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_component, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetComponent(ctx, input)
			},
		},
		"get-component-version-artifact": {
			Name:   "get-component-version-artifact",
			Fields: fields_get_component_version_artifact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetComponentVersionArtifactInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_component_version_artifact, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetComponentVersionArtifact(ctx, input)
			},
		},
		"get-connectivity-info": {
			Name:   "get-connectivity-info",
			Fields: fields_get_connectivity_info,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConnectivityInfoInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_connectivity_info, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConnectivityInfo(ctx, input)
			},
		},
		"get-core-device": {
			Name:   "get-core-device",
			Fields: fields_get_core_device,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCoreDeviceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_core_device, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCoreDevice(ctx, input)
			},
		},
		"get-deployment": {
			Name:   "get-deployment",
			Fields: fields_get_deployment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDeploymentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_deployment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDeployment(ctx, input)
			},
		},
		"get-service-role-for-account": {
			Name:   "get-service-role-for-account",
			Fields: fields_get_service_role_for_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetServiceRoleForAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_service_role_for_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetServiceRoleForAccount(ctx, input)
			},
		},
		"list-client-devices-associated-with-core-device": {
			Name:   "list-client-devices-associated-with-core-device",
			Fields: fields_list_client_devices_associated_with_core_device,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListClientDevicesAssociatedWithCoreDeviceInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_client_devices_associated_with_core_device, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListClientDevicesAssociatedWithCoreDevice(ctx, input)
				}
				var results []*svc.ListClientDevicesAssociatedWithCoreDeviceOutput
				p := svc.NewListClientDevicesAssociatedWithCoreDevicePaginator(client, input)
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
		"list-component-versions": {
			Name:   "list-component-versions",
			Fields: fields_list_component_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListComponentVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_component_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListComponentVersions(ctx, input)
				}
				var results []*svc.ListComponentVersionsOutput
				p := svc.NewListComponentVersionsPaginator(client, input)
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
		"list-components": {
			Name:   "list-components",
			Fields: fields_list_components,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListComponentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_components, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListComponents(ctx, input)
				}
				var results []*svc.ListComponentsOutput
				p := svc.NewListComponentsPaginator(client, input)
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
		"list-core-devices": {
			Name:   "list-core-devices",
			Fields: fields_list_core_devices,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCoreDevicesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_core_devices, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCoreDevices(ctx, input)
				}
				var results []*svc.ListCoreDevicesOutput
				p := svc.NewListCoreDevicesPaginator(client, input)
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
		"list-deployments": {
			Name:   "list-deployments",
			Fields: fields_list_deployments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDeploymentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_deployments, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDeployments(ctx, input)
				}
				var results []*svc.ListDeploymentsOutput
				p := svc.NewListDeploymentsPaginator(client, input)
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
		"list-effective-deployments": {
			Name:   "list-effective-deployments",
			Fields: fields_list_effective_deployments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEffectiveDeploymentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_effective_deployments, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEffectiveDeployments(ctx, input)
				}
				var results []*svc.ListEffectiveDeploymentsOutput
				p := svc.NewListEffectiveDeploymentsPaginator(client, input)
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
		"list-installed-components": {
			Name:   "list-installed-components",
			Fields: fields_list_installed_components,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInstalledComponentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_installed_components, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListInstalledComponents(ctx, input)
				}
				var results []*svc.ListInstalledComponentsOutput
				p := svc.NewListInstalledComponentsPaginator(client, input)
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
		"resolve-component-candidates": {
			Name:   "resolve-component-candidates",
			Fields: fields_resolve_component_candidates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResolveComponentCandidatesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_resolve_component_candidates, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResolveComponentCandidates(ctx, input)
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
		"update-connectivity-info": {
			Name:   "update-connectivity-info",
			Fields: fields_update_connectivity_info,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateConnectivityInfoInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_connectivity_info, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateConnectivityInfo(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("greengrassv2", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
