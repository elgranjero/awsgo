package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/tnb"
)

var fields_cancel_sol_network_operation = []leanruntime.Field{
	{Name: "NsLcmOpOccId", Flag: "ns-lcm-op-occ-id", Type: "*string", Required: true},
}

var fields_create_sol_function_package = []leanruntime.Field{
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_sol_network_instance = []leanruntime.Field{
	{Name: "NsDescription", Flag: "ns-description", Type: "*string", Required: false},
	{Name: "NsName", Flag: "ns-name", Type: "*string", Required: true},
	{Name: "NsdInfoId", Flag: "nsd-info-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_sol_network_package = []leanruntime.Field{
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_sol_function_package = []leanruntime.Field{
	{Name: "VnfPkgId", Flag: "vnf-pkg-id", Type: "*string", Required: true},
}

var fields_delete_sol_network_instance = []leanruntime.Field{
	{Name: "NsInstanceId", Flag: "ns-instance-id", Type: "*string", Required: true},
}

var fields_delete_sol_network_package = []leanruntime.Field{
	{Name: "NsdInfoId", Flag: "nsd-info-id", Type: "*string", Required: true},
}

var fields_get_sol_function_instance = []leanruntime.Field{
	{Name: "VnfInstanceId", Flag: "vnf-instance-id", Type: "*string", Required: true},
}

var fields_get_sol_function_package = []leanruntime.Field{
	{Name: "VnfPkgId", Flag: "vnf-pkg-id", Type: "*string", Required: true},
}

var fields_get_sol_function_package_content = []leanruntime.Field{
	{Name: "Accept", Flag: "accept", Type: "types.PackageContentType", Required: true},
	{Name: "VnfPkgId", Flag: "vnf-pkg-id", Type: "*string", Required: true},
}

var fields_get_sol_function_package_descriptor = []leanruntime.Field{
	{Name: "Accept", Flag: "accept", Type: "types.DescriptorContentType", Required: true},
	{Name: "VnfPkgId", Flag: "vnf-pkg-id", Type: "*string", Required: true},
}

var fields_get_sol_network_instance = []leanruntime.Field{
	{Name: "NsInstanceId", Flag: "ns-instance-id", Type: "*string", Required: true},
}

var fields_get_sol_network_operation = []leanruntime.Field{
	{Name: "NsLcmOpOccId", Flag: "ns-lcm-op-occ-id", Type: "*string", Required: true},
}

var fields_get_sol_network_package = []leanruntime.Field{
	{Name: "NsdInfoId", Flag: "nsd-info-id", Type: "*string", Required: true},
}

var fields_get_sol_network_package_content = []leanruntime.Field{
	{Name: "Accept", Flag: "accept", Type: "types.PackageContentType", Required: true},
	{Name: "NsdInfoId", Flag: "nsd-info-id", Type: "*string", Required: true},
}

var fields_get_sol_network_package_descriptor = []leanruntime.Field{
	{Name: "NsdInfoId", Flag: "nsd-info-id", Type: "*string", Required: true},
}

var fields_instantiate_sol_network_instance = []leanruntime.Field{
	{Name: "AdditionalParamsForNs", Flag: "additional-params-for-ns", Type: "document.Interface", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "NsInstanceId", Flag: "ns-instance-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_list_sol_function_instances = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_sol_function_packages = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_sol_network_instances = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_sol_network_operations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "NsInstanceId", Flag: "ns-instance-id", Type: "*string", Required: false},
}

var fields_list_sol_network_packages = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_put_sol_function_package_content = []leanruntime.Field{
	{Name: "ContentType", Flag: "content-type", Type: "types.PackageContentType", Required: false},
	{Name: "File", Flag: "file", Type: "[]byte", Required: true},
	{Name: "VnfPkgId", Flag: "vnf-pkg-id", Type: "*string", Required: true},
}

var fields_put_sol_network_package_content = []leanruntime.Field{
	{Name: "ContentType", Flag: "content-type", Type: "types.PackageContentType", Required: false},
	{Name: "File", Flag: "file", Type: "[]byte", Required: true},
	{Name: "NsdInfoId", Flag: "nsd-info-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_terminate_sol_network_instance = []leanruntime.Field{
	{Name: "NsInstanceId", Flag: "ns-instance-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_sol_function_package = []leanruntime.Field{
	{Name: "OperationalState", Flag: "operational-state", Type: "types.OperationalState", Required: true},
	{Name: "VnfPkgId", Flag: "vnf-pkg-id", Type: "*string", Required: true},
}

var fields_update_sol_network_instance = []leanruntime.Field{
	{Name: "ModifyVnfInfoData", Flag: "modify-vnf-info-data", Type: "*types.UpdateSolNetworkModify", Required: false},
	{Name: "NsInstanceId", Flag: "ns-instance-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "UpdateNs", Flag: "update-ns", Type: "*types.UpdateSolNetworkServiceData", Required: false},
	{Name: "UpdateType", Flag: "update-type", Type: "types.UpdateSolNetworkType", Required: true},
}

var fields_update_sol_network_package = []leanruntime.Field{
	{Name: "NsdInfoId", Flag: "nsd-info-id", Type: "*string", Required: true},
	{Name: "NsdOperationalState", Flag: "nsd-operational-state", Type: "types.NsdOperationalState", Required: true},
}

var fields_validate_sol_function_package_content = []leanruntime.Field{
	{Name: "ContentType", Flag: "content-type", Type: "types.PackageContentType", Required: false},
	{Name: "File", Flag: "file", Type: "[]byte", Required: true},
	{Name: "VnfPkgId", Flag: "vnf-pkg-id", Type: "*string", Required: true},
}

var fields_validate_sol_network_package_content = []leanruntime.Field{
	{Name: "ContentType", Flag: "content-type", Type: "types.PackageContentType", Required: false},
	{Name: "File", Flag: "file", Type: "[]byte", Required: true},
	{Name: "NsdInfoId", Flag: "nsd-info-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"cancel-sol-network-operation": {
			Name:   "cancel-sol-network-operation",
			Fields: fields_cancel_sol_network_operation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelSolNetworkOperationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_sol_network_operation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelSolNetworkOperation(ctx, input)
			},
		},
		"create-sol-function-package": {
			Name:   "create-sol-function-package",
			Fields: fields_create_sol_function_package,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSolFunctionPackageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_sol_function_package, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSolFunctionPackage(ctx, input)
			},
		},
		"create-sol-network-instance": {
			Name:   "create-sol-network-instance",
			Fields: fields_create_sol_network_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSolNetworkInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_sol_network_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSolNetworkInstance(ctx, input)
			},
		},
		"create-sol-network-package": {
			Name:   "create-sol-network-package",
			Fields: fields_create_sol_network_package,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSolNetworkPackageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_sol_network_package, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSolNetworkPackage(ctx, input)
			},
		},
		"delete-sol-function-package": {
			Name:   "delete-sol-function-package",
			Fields: fields_delete_sol_function_package,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSolFunctionPackageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_sol_function_package, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSolFunctionPackage(ctx, input)
			},
		},
		"delete-sol-network-instance": {
			Name:   "delete-sol-network-instance",
			Fields: fields_delete_sol_network_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSolNetworkInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_sol_network_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSolNetworkInstance(ctx, input)
			},
		},
		"delete-sol-network-package": {
			Name:   "delete-sol-network-package",
			Fields: fields_delete_sol_network_package,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSolNetworkPackageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_sol_network_package, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSolNetworkPackage(ctx, input)
			},
		},
		"get-sol-function-instance": {
			Name:   "get-sol-function-instance",
			Fields: fields_get_sol_function_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSolFunctionInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_sol_function_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSolFunctionInstance(ctx, input)
			},
		},
		"get-sol-function-package": {
			Name:   "get-sol-function-package",
			Fields: fields_get_sol_function_package,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSolFunctionPackageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_sol_function_package, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSolFunctionPackage(ctx, input)
			},
		},
		"get-sol-function-package-content": {
			Name:   "get-sol-function-package-content",
			Fields: fields_get_sol_function_package_content,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSolFunctionPackageContentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_sol_function_package_content, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSolFunctionPackageContent(ctx, input)
			},
		},
		"get-sol-function-package-descriptor": {
			Name:   "get-sol-function-package-descriptor",
			Fields: fields_get_sol_function_package_descriptor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSolFunctionPackageDescriptorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_sol_function_package_descriptor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSolFunctionPackageDescriptor(ctx, input)
			},
		},
		"get-sol-network-instance": {
			Name:   "get-sol-network-instance",
			Fields: fields_get_sol_network_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSolNetworkInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_sol_network_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSolNetworkInstance(ctx, input)
			},
		},
		"get-sol-network-operation": {
			Name:   "get-sol-network-operation",
			Fields: fields_get_sol_network_operation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSolNetworkOperationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_sol_network_operation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSolNetworkOperation(ctx, input)
			},
		},
		"get-sol-network-package": {
			Name:   "get-sol-network-package",
			Fields: fields_get_sol_network_package,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSolNetworkPackageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_sol_network_package, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSolNetworkPackage(ctx, input)
			},
		},
		"get-sol-network-package-content": {
			Name:   "get-sol-network-package-content",
			Fields: fields_get_sol_network_package_content,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSolNetworkPackageContentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_sol_network_package_content, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSolNetworkPackageContent(ctx, input)
			},
		},
		"get-sol-network-package-descriptor": {
			Name:   "get-sol-network-package-descriptor",
			Fields: fields_get_sol_network_package_descriptor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSolNetworkPackageDescriptorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_sol_network_package_descriptor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSolNetworkPackageDescriptor(ctx, input)
			},
		},
		"instantiate-sol-network-instance": {
			Name:   "instantiate-sol-network-instance",
			Fields: fields_instantiate_sol_network_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.InstantiateSolNetworkInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_instantiate_sol_network_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.InstantiateSolNetworkInstance(ctx, input)
			},
		},
		"list-sol-function-instances": {
			Name:   "list-sol-function-instances",
			Fields: fields_list_sol_function_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSolFunctionInstancesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_sol_function_instances, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSolFunctionInstances(ctx, input)
				}
				var results []*svc.ListSolFunctionInstancesOutput
				p := svc.NewListSolFunctionInstancesPaginator(client, input)
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
		"list-sol-function-packages": {
			Name:   "list-sol-function-packages",
			Fields: fields_list_sol_function_packages,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSolFunctionPackagesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_sol_function_packages, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSolFunctionPackages(ctx, input)
				}
				var results []*svc.ListSolFunctionPackagesOutput
				p := svc.NewListSolFunctionPackagesPaginator(client, input)
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
		"list-sol-network-instances": {
			Name:   "list-sol-network-instances",
			Fields: fields_list_sol_network_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSolNetworkInstancesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_sol_network_instances, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSolNetworkInstances(ctx, input)
				}
				var results []*svc.ListSolNetworkInstancesOutput
				p := svc.NewListSolNetworkInstancesPaginator(client, input)
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
		"list-sol-network-operations": {
			Name:   "list-sol-network-operations",
			Fields: fields_list_sol_network_operations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSolNetworkOperationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_sol_network_operations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSolNetworkOperations(ctx, input)
				}
				var results []*svc.ListSolNetworkOperationsOutput
				p := svc.NewListSolNetworkOperationsPaginator(client, input)
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
		"list-sol-network-packages": {
			Name:   "list-sol-network-packages",
			Fields: fields_list_sol_network_packages,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSolNetworkPackagesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_sol_network_packages, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSolNetworkPackages(ctx, input)
				}
				var results []*svc.ListSolNetworkPackagesOutput
				p := svc.NewListSolNetworkPackagesPaginator(client, input)
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
		"put-sol-function-package-content": {
			Name:   "put-sol-function-package-content",
			Fields: fields_put_sol_function_package_content,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutSolFunctionPackageContentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_sol_function_package_content, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutSolFunctionPackageContent(ctx, input)
			},
		},
		"put-sol-network-package-content": {
			Name:   "put-sol-network-package-content",
			Fields: fields_put_sol_network_package_content,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutSolNetworkPackageContentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_sol_network_package_content, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutSolNetworkPackageContent(ctx, input)
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
		"terminate-sol-network-instance": {
			Name:   "terminate-sol-network-instance",
			Fields: fields_terminate_sol_network_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TerminateSolNetworkInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_terminate_sol_network_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TerminateSolNetworkInstance(ctx, input)
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
		"update-sol-function-package": {
			Name:   "update-sol-function-package",
			Fields: fields_update_sol_function_package,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSolFunctionPackageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_sol_function_package, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSolFunctionPackage(ctx, input)
			},
		},
		"update-sol-network-instance": {
			Name:   "update-sol-network-instance",
			Fields: fields_update_sol_network_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSolNetworkInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_sol_network_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSolNetworkInstance(ctx, input)
			},
		},
		"update-sol-network-package": {
			Name:   "update-sol-network-package",
			Fields: fields_update_sol_network_package,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSolNetworkPackageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_sol_network_package, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSolNetworkPackage(ctx, input)
			},
		},
		"validate-sol-function-package-content": {
			Name:   "validate-sol-function-package-content",
			Fields: fields_validate_sol_function_package_content,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ValidateSolFunctionPackageContentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_validate_sol_function_package_content, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ValidateSolFunctionPackageContent(ctx, input)
			},
		},
		"validate-sol-network-package-content": {
			Name:   "validate-sol-network-package-content",
			Fields: fields_validate_sol_network_package_content,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ValidateSolNetworkPackageContentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_validate_sol_network_package_content, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ValidateSolNetworkPackageContent(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("tnb", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
