package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/arczonalshift"
)

var fields_cancel_practice_run = []leanruntime.Field{
	{Name: "ZonalShiftId", Flag: "zonal-shift-id", Type: "*string", Required: true},
}

var fields_cancel_zonal_shift = []leanruntime.Field{
	{Name: "ZonalShiftId", Flag: "zonal-shift-id", Type: "*string", Required: true},
}

var fields_create_practice_run_configuration = []leanruntime.Field{
	{Name: "AllowedWindows", Flag: "allowed-windows", Type: "[]string", Required: false},
	{Name: "BlockedDates", Flag: "blocked-dates", Type: "[]string", Required: false},
	{Name: "BlockedWindows", Flag: "blocked-windows", Type: "[]string", Required: false},
	{Name: "BlockingAlarms", Flag: "blocking-alarms", Type: "[]types.ControlCondition", Required: false},
	{Name: "OutcomeAlarms", Flag: "outcome-alarms", Type: "[]types.ControlCondition", Required: true},
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*string", Required: true},
}

var fields_delete_practice_run_configuration = []leanruntime.Field{
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*string", Required: true},
}

var fields_get_autoshift_observer_notification_status = []leanruntime.Field{}

var fields_get_managed_resource = []leanruntime.Field{
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*string", Required: true},
}

var fields_list_autoshifts = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.AutoshiftExecutionStatus", Required: false},
}

var fields_list_managed_resources = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_zonal_shifts = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.ZonalShiftStatus", Required: false},
}

var fields_start_practice_run = []leanruntime.Field{
	{Name: "AwayFrom", Flag: "away-from", Type: "*string", Required: true},
	{Name: "Comment", Flag: "comment", Type: "*string", Required: true},
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*string", Required: true},
}

var fields_start_zonal_shift = []leanruntime.Field{
	{Name: "AwayFrom", Flag: "away-from", Type: "*string", Required: true},
	{Name: "Comment", Flag: "comment", Type: "*string", Required: true},
	{Name: "ExpiresIn", Flag: "expires-in", Type: "*string", Required: true},
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*string", Required: true},
}

var fields_update_autoshift_observer_notification_status = []leanruntime.Field{
	{Name: "Status", Flag: "status", Type: "types.AutoshiftObserverNotificationStatus", Required: true},
}

var fields_update_practice_run_configuration = []leanruntime.Field{
	{Name: "AllowedWindows", Flag: "allowed-windows", Type: "[]string", Required: false},
	{Name: "BlockedDates", Flag: "blocked-dates", Type: "[]string", Required: false},
	{Name: "BlockedWindows", Flag: "blocked-windows", Type: "[]string", Required: false},
	{Name: "BlockingAlarms", Flag: "blocking-alarms", Type: "[]types.ControlCondition", Required: false},
	{Name: "OutcomeAlarms", Flag: "outcome-alarms", Type: "[]types.ControlCondition", Required: false},
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*string", Required: true},
}

var fields_update_zonal_autoshift_configuration = []leanruntime.Field{
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*string", Required: true},
	{Name: "ZonalAutoshiftStatus", Flag: "zonal-autoshift-status", Type: "types.ZonalAutoshiftStatus", Required: true},
}

var fields_update_zonal_shift = []leanruntime.Field{
	{Name: "Comment", Flag: "comment", Type: "*string", Required: false},
	{Name: "ExpiresIn", Flag: "expires-in", Type: "*string", Required: false},
	{Name: "ZonalShiftId", Flag: "zonal-shift-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"cancel-practice-run": {
			Name:   "cancel-practice-run",
			Fields: fields_cancel_practice_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelPracticeRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_practice_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelPracticeRun(ctx, input)
			},
		},
		"cancel-zonal-shift": {
			Name:   "cancel-zonal-shift",
			Fields: fields_cancel_zonal_shift,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelZonalShiftInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_zonal_shift, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelZonalShift(ctx, input)
			},
		},
		"create-practice-run-configuration": {
			Name:   "create-practice-run-configuration",
			Fields: fields_create_practice_run_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePracticeRunConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_practice_run_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePracticeRunConfiguration(ctx, input)
			},
		},
		"delete-practice-run-configuration": {
			Name:   "delete-practice-run-configuration",
			Fields: fields_delete_practice_run_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePracticeRunConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_practice_run_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePracticeRunConfiguration(ctx, input)
			},
		},
		"get-autoshift-observer-notification-status": {
			Name:   "get-autoshift-observer-notification-status",
			Fields: fields_get_autoshift_observer_notification_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAutoshiftObserverNotificationStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_autoshift_observer_notification_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAutoshiftObserverNotificationStatus(ctx, input)
			},
		},
		"get-managed-resource": {
			Name:   "get-managed-resource",
			Fields: fields_get_managed_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetManagedResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_managed_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetManagedResource(ctx, input)
			},
		},
		"list-autoshifts": {
			Name:   "list-autoshifts",
			Fields: fields_list_autoshifts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAutoshiftsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_autoshifts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAutoshifts(ctx, input)
				}
				var results []*svc.ListAutoshiftsOutput
				p := svc.NewListAutoshiftsPaginator(client, input)
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
		"list-managed-resources": {
			Name:   "list-managed-resources",
			Fields: fields_list_managed_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListManagedResourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_managed_resources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListManagedResources(ctx, input)
				}
				var results []*svc.ListManagedResourcesOutput
				p := svc.NewListManagedResourcesPaginator(client, input)
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
		"list-zonal-shifts": {
			Name:   "list-zonal-shifts",
			Fields: fields_list_zonal_shifts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListZonalShiftsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_zonal_shifts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListZonalShifts(ctx, input)
				}
				var results []*svc.ListZonalShiftsOutput
				p := svc.NewListZonalShiftsPaginator(client, input)
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
		"start-practice-run": {
			Name:   "start-practice-run",
			Fields: fields_start_practice_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartPracticeRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_practice_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartPracticeRun(ctx, input)
			},
		},
		"start-zonal-shift": {
			Name:   "start-zonal-shift",
			Fields: fields_start_zonal_shift,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartZonalShiftInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_zonal_shift, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartZonalShift(ctx, input)
			},
		},
		"update-autoshift-observer-notification-status": {
			Name:   "update-autoshift-observer-notification-status",
			Fields: fields_update_autoshift_observer_notification_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAutoshiftObserverNotificationStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_autoshift_observer_notification_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAutoshiftObserverNotificationStatus(ctx, input)
			},
		},
		"update-practice-run-configuration": {
			Name:   "update-practice-run-configuration",
			Fields: fields_update_practice_run_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePracticeRunConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_practice_run_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePracticeRunConfiguration(ctx, input)
			},
		},
		"update-zonal-autoshift-configuration": {
			Name:   "update-zonal-autoshift-configuration",
			Fields: fields_update_zonal_autoshift_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateZonalAutoshiftConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_zonal_autoshift_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateZonalAutoshiftConfiguration(ctx, input)
			},
		},
		"update-zonal-shift": {
			Name:   "update-zonal-shift",
			Fields: fields_update_zonal_shift,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateZonalShiftInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_zonal_shift, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateZonalShift(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("arczonalshift", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
