package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/controltower"
)

var fields_create_landing_zone = []leanruntime.Field{
	{Name: "Manifest", Flag: "manifest", Type: "document.Interface", Required: false},
	{Name: "RemediationTypes", Flag: "remediation-types", Type: "[]types.RemediationType", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Version", Flag: "version", Type: "*string", Required: true},
}

var fields_delete_landing_zone = []leanruntime.Field{
	{Name: "LandingZoneIdentifier", Flag: "landing-zone-identifier", Type: "*string", Required: true},
}

var fields_disable_baseline = []leanruntime.Field{
	{Name: "EnabledBaselineIdentifier", Flag: "enabled-baseline-identifier", Type: "*string", Required: true},
}

var fields_disable_control = []leanruntime.Field{
	{Name: "ControlIdentifier", Flag: "control-identifier", Type: "*string", Required: false},
	{Name: "EnabledControlIdentifier", Flag: "enabled-control-identifier", Type: "*string", Required: false},
	{Name: "TargetIdentifier", Flag: "target-identifier", Type: "*string", Required: false},
}

var fields_enable_baseline = []leanruntime.Field{
	{Name: "BaselineIdentifier", Flag: "baseline-identifier", Type: "*string", Required: true},
	{Name: "BaselineVersion", Flag: "baseline-version", Type: "*string", Required: true},
	{Name: "Parameters", Flag: "parameters", Type: "[]types.EnabledBaselineParameter", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TargetIdentifier", Flag: "target-identifier", Type: "*string", Required: true},
}

var fields_enable_control = []leanruntime.Field{
	{Name: "ControlIdentifier", Flag: "control-identifier", Type: "*string", Required: true},
	{Name: "Parameters", Flag: "parameters", Type: "[]types.EnabledControlParameter", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TargetIdentifier", Flag: "target-identifier", Type: "*string", Required: true},
}

var fields_get_baseline = []leanruntime.Field{
	{Name: "BaselineIdentifier", Flag: "baseline-identifier", Type: "*string", Required: true},
}

var fields_get_baseline_operation = []leanruntime.Field{
	{Name: "OperationIdentifier", Flag: "operation-identifier", Type: "*string", Required: true},
}

var fields_get_control_operation = []leanruntime.Field{
	{Name: "OperationIdentifier", Flag: "operation-identifier", Type: "*string", Required: true},
}

var fields_get_enabled_baseline = []leanruntime.Field{
	{Name: "EnabledBaselineIdentifier", Flag: "enabled-baseline-identifier", Type: "*string", Required: true},
}

var fields_get_enabled_control = []leanruntime.Field{
	{Name: "EnabledControlIdentifier", Flag: "enabled-control-identifier", Type: "*string", Required: true},
}

var fields_get_landing_zone = []leanruntime.Field{
	{Name: "LandingZoneIdentifier", Flag: "landing-zone-identifier", Type: "*string", Required: true},
}

var fields_get_landing_zone_operation = []leanruntime.Field{
	{Name: "OperationIdentifier", Flag: "operation-identifier", Type: "*string", Required: true},
}

var fields_list_baselines = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_control_operations = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.ControlOperationFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_enabled_baselines = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.EnabledBaselineFilter", Required: false},
	{Name: "IncludeChildren", Flag: "include-children", Type: "bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_enabled_controls = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.EnabledControlFilter", Required: false},
	{Name: "IncludeChildren", Flag: "include-children", Type: "bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TargetIdentifier", Flag: "target-identifier", Type: "*string", Required: false},
}

var fields_list_landing_zone_operations = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.LandingZoneOperationFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_landing_zones = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_reset_enabled_baseline = []leanruntime.Field{
	{Name: "EnabledBaselineIdentifier", Flag: "enabled-baseline-identifier", Type: "*string", Required: true},
}

var fields_reset_enabled_control = []leanruntime.Field{
	{Name: "EnabledControlIdentifier", Flag: "enabled-control-identifier", Type: "*string", Required: true},
}

var fields_reset_landing_zone = []leanruntime.Field{
	{Name: "LandingZoneIdentifier", Flag: "landing-zone-identifier", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_enabled_baseline = []leanruntime.Field{
	{Name: "BaselineVersion", Flag: "baseline-version", Type: "*string", Required: true},
	{Name: "EnabledBaselineIdentifier", Flag: "enabled-baseline-identifier", Type: "*string", Required: true},
	{Name: "Parameters", Flag: "parameters", Type: "[]types.EnabledBaselineParameter", Required: false},
}

var fields_update_enabled_control = []leanruntime.Field{
	{Name: "EnabledControlIdentifier", Flag: "enabled-control-identifier", Type: "*string", Required: true},
	{Name: "Parameters", Flag: "parameters", Type: "[]types.EnabledControlParameter", Required: true},
}

var fields_update_landing_zone = []leanruntime.Field{
	{Name: "LandingZoneIdentifier", Flag: "landing-zone-identifier", Type: "*string", Required: true},
	{Name: "Manifest", Flag: "manifest", Type: "document.Interface", Required: false},
	{Name: "RemediationTypes", Flag: "remediation-types", Type: "[]types.RemediationType", Required: false},
	{Name: "Version", Flag: "version", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-landing-zone": {
			Name:   "create-landing-zone",
			Fields: fields_create_landing_zone,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLandingZoneInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_landing_zone, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLandingZone(ctx, input)
			},
		},
		"delete-landing-zone": {
			Name:   "delete-landing-zone",
			Fields: fields_delete_landing_zone,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLandingZoneInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_landing_zone, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLandingZone(ctx, input)
			},
		},
		"disable-baseline": {
			Name:   "disable-baseline",
			Fields: fields_disable_baseline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableBaselineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_baseline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableBaseline(ctx, input)
			},
		},
		"disable-control": {
			Name:   "disable-control",
			Fields: fields_disable_control,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableControlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_control, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableControl(ctx, input)
			},
		},
		"enable-baseline": {
			Name:   "enable-baseline",
			Fields: fields_enable_baseline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableBaselineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_baseline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableBaseline(ctx, input)
			},
		},
		"enable-control": {
			Name:   "enable-control",
			Fields: fields_enable_control,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableControlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_control, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableControl(ctx, input)
			},
		},
		"get-baseline": {
			Name:   "get-baseline",
			Fields: fields_get_baseline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBaselineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_baseline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBaseline(ctx, input)
			},
		},
		"get-baseline-operation": {
			Name:   "get-baseline-operation",
			Fields: fields_get_baseline_operation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBaselineOperationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_baseline_operation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBaselineOperation(ctx, input)
			},
		},
		"get-control-operation": {
			Name:   "get-control-operation",
			Fields: fields_get_control_operation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetControlOperationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_control_operation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetControlOperation(ctx, input)
			},
		},
		"get-enabled-baseline": {
			Name:   "get-enabled-baseline",
			Fields: fields_get_enabled_baseline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEnabledBaselineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_enabled_baseline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEnabledBaseline(ctx, input)
			},
		},
		"get-enabled-control": {
			Name:   "get-enabled-control",
			Fields: fields_get_enabled_control,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEnabledControlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_enabled_control, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEnabledControl(ctx, input)
			},
		},
		"get-landing-zone": {
			Name:   "get-landing-zone",
			Fields: fields_get_landing_zone,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLandingZoneInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_landing_zone, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLandingZone(ctx, input)
			},
		},
		"get-landing-zone-operation": {
			Name:   "get-landing-zone-operation",
			Fields: fields_get_landing_zone_operation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLandingZoneOperationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_landing_zone_operation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLandingZoneOperation(ctx, input)
			},
		},
		"list-baselines": {
			Name:   "list-baselines",
			Fields: fields_list_baselines,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBaselinesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_baselines, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBaselines(ctx, input)
				}
				var results []*svc.ListBaselinesOutput
				p := svc.NewListBaselinesPaginator(client, input)
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
		"list-control-operations": {
			Name:   "list-control-operations",
			Fields: fields_list_control_operations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListControlOperationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_control_operations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListControlOperations(ctx, input)
				}
				var results []*svc.ListControlOperationsOutput
				p := svc.NewListControlOperationsPaginator(client, input)
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
		"list-enabled-baselines": {
			Name:   "list-enabled-baselines",
			Fields: fields_list_enabled_baselines,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEnabledBaselinesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_enabled_baselines, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEnabledBaselines(ctx, input)
				}
				var results []*svc.ListEnabledBaselinesOutput
				p := svc.NewListEnabledBaselinesPaginator(client, input)
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
		"list-enabled-controls": {
			Name:   "list-enabled-controls",
			Fields: fields_list_enabled_controls,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEnabledControlsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_enabled_controls, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEnabledControls(ctx, input)
				}
				var results []*svc.ListEnabledControlsOutput
				p := svc.NewListEnabledControlsPaginator(client, input)
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
		"list-landing-zone-operations": {
			Name:   "list-landing-zone-operations",
			Fields: fields_list_landing_zone_operations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLandingZoneOperationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_landing_zone_operations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLandingZoneOperations(ctx, input)
				}
				var results []*svc.ListLandingZoneOperationsOutput
				p := svc.NewListLandingZoneOperationsPaginator(client, input)
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
		"list-landing-zones": {
			Name:   "list-landing-zones",
			Fields: fields_list_landing_zones,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLandingZonesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_landing_zones, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLandingZones(ctx, input)
				}
				var results []*svc.ListLandingZonesOutput
				p := svc.NewListLandingZonesPaginator(client, input)
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
		"reset-enabled-baseline": {
			Name:   "reset-enabled-baseline",
			Fields: fields_reset_enabled_baseline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResetEnabledBaselineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reset_enabled_baseline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResetEnabledBaseline(ctx, input)
			},
		},
		"reset-enabled-control": {
			Name:   "reset-enabled-control",
			Fields: fields_reset_enabled_control,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResetEnabledControlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reset_enabled_control, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResetEnabledControl(ctx, input)
			},
		},
		"reset-landing-zone": {
			Name:   "reset-landing-zone",
			Fields: fields_reset_landing_zone,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResetLandingZoneInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reset_landing_zone, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResetLandingZone(ctx, input)
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
		"update-enabled-baseline": {
			Name:   "update-enabled-baseline",
			Fields: fields_update_enabled_baseline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEnabledBaselineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_enabled_baseline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEnabledBaseline(ctx, input)
			},
		},
		"update-enabled-control": {
			Name:   "update-enabled-control",
			Fields: fields_update_enabled_control,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEnabledControlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_enabled_control, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEnabledControl(ctx, input)
			},
		},
		"update-landing-zone": {
			Name:   "update-landing-zone",
			Fields: fields_update_landing_zone,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateLandingZoneInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_landing_zone, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateLandingZone(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("controltower", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
