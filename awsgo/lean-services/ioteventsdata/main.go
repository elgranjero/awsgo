package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/ioteventsdata"
)

var fields_batch_acknowledge_alarm = []leanruntime.Field{
	{Name: "AcknowledgeActionRequests", Flag: "acknowledge-action-requests", Type: "[]types.AcknowledgeAlarmActionRequest", Required: true},
}

var fields_batch_delete_detector = []leanruntime.Field{
	{Name: "Detectors", Flag: "detectors", Type: "[]types.DeleteDetectorRequest", Required: true},
}

var fields_batch_disable_alarm = []leanruntime.Field{
	{Name: "DisableActionRequests", Flag: "disable-action-requests", Type: "[]types.DisableAlarmActionRequest", Required: true},
}

var fields_batch_enable_alarm = []leanruntime.Field{
	{Name: "EnableActionRequests", Flag: "enable-action-requests", Type: "[]types.EnableAlarmActionRequest", Required: true},
}

var fields_batch_put_message = []leanruntime.Field{
	{Name: "Messages", Flag: "messages", Type: "[]types.Message", Required: true},
}

var fields_batch_reset_alarm = []leanruntime.Field{
	{Name: "ResetActionRequests", Flag: "reset-action-requests", Type: "[]types.ResetAlarmActionRequest", Required: true},
}

var fields_batch_snooze_alarm = []leanruntime.Field{
	{Name: "SnoozeActionRequests", Flag: "snooze-action-requests", Type: "[]types.SnoozeAlarmActionRequest", Required: true},
}

var fields_batch_update_detector = []leanruntime.Field{
	{Name: "Detectors", Flag: "detectors", Type: "[]types.UpdateDetectorRequest", Required: true},
}

var fields_describe_alarm = []leanruntime.Field{
	{Name: "AlarmModelName", Flag: "alarm-model-name", Type: "*string", Required: true},
	{Name: "KeyValue", Flag: "key-value", Type: "*string", Required: false},
}

var fields_describe_detector = []leanruntime.Field{
	{Name: "DetectorModelName", Flag: "detector-model-name", Type: "*string", Required: true},
	{Name: "KeyValue", Flag: "key-value", Type: "*string", Required: false},
}

var fields_list_alarms = []leanruntime.Field{
	{Name: "AlarmModelName", Flag: "alarm-model-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_detectors = []leanruntime.Field{
	{Name: "DetectorModelName", Flag: "detector-model-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StateName", Flag: "state-name", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"batch-acknowledge-alarm": {
			Name:   "batch-acknowledge-alarm",
			Fields: fields_batch_acknowledge_alarm,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchAcknowledgeAlarmInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_acknowledge_alarm, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchAcknowledgeAlarm(ctx, input)
			},
		},
		"batch-delete-detector": {
			Name:   "batch-delete-detector",
			Fields: fields_batch_delete_detector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDeleteDetectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_delete_detector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDeleteDetector(ctx, input)
			},
		},
		"batch-disable-alarm": {
			Name:   "batch-disable-alarm",
			Fields: fields_batch_disable_alarm,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDisableAlarmInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_disable_alarm, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDisableAlarm(ctx, input)
			},
		},
		"batch-enable-alarm": {
			Name:   "batch-enable-alarm",
			Fields: fields_batch_enable_alarm,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchEnableAlarmInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_enable_alarm, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchEnableAlarm(ctx, input)
			},
		},
		"batch-put-message": {
			Name:   "batch-put-message",
			Fields: fields_batch_put_message,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchPutMessageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_put_message, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchPutMessage(ctx, input)
			},
		},
		"batch-reset-alarm": {
			Name:   "batch-reset-alarm",
			Fields: fields_batch_reset_alarm,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchResetAlarmInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_reset_alarm, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchResetAlarm(ctx, input)
			},
		},
		"batch-snooze-alarm": {
			Name:   "batch-snooze-alarm",
			Fields: fields_batch_snooze_alarm,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchSnoozeAlarmInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_snooze_alarm, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchSnoozeAlarm(ctx, input)
			},
		},
		"batch-update-detector": {
			Name:   "batch-update-detector",
			Fields: fields_batch_update_detector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchUpdateDetectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_update_detector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchUpdateDetector(ctx, input)
			},
		},
		"describe-alarm": {
			Name:   "describe-alarm",
			Fields: fields_describe_alarm,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAlarmInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_alarm, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAlarm(ctx, input)
			},
		},
		"describe-detector": {
			Name:   "describe-detector",
			Fields: fields_describe_detector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDetectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_detector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDetector(ctx, input)
			},
		},
		"list-alarms": {
			Name:   "list-alarms",
			Fields: fields_list_alarms,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAlarmsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_alarms, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListAlarms(ctx, input)
			},
		},
		"list-detectors": {
			Name:   "list-detectors",
			Fields: fields_list_detectors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDetectorsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_detectors, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListDetectors(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("ioteventsdata", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
