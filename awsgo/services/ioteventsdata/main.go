package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/ioteventsdata/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"batch-acknowledge-alarm", "batch-delete-detector", "batch-disable-alarm", "batch-enable-alarm", "batch-put-message", "batch-reset-alarm", "batch-snooze-alarm", "batch-update-detector", "describe-alarm", "describe-detector", "list-alarms", "list-detectors"},
		OperationSet: map[string]bool{"batch-acknowledge-alarm": true, "batch-delete-detector": true, "batch-disable-alarm": true, "batch-enable-alarm": true, "batch-put-message": true, "batch-reset-alarm": true, "batch-snooze-alarm": true, "batch-update-detector": true, "describe-alarm": true, "describe-detector": true, "list-alarms": true, "list-detectors": true},
		OperationInputs: map[string][]string{
			"batch-acknowledge-alarm": {"AcknowledgeActionRequests"},
			"batch-delete-detector":   {"Detectors"},
			"batch-disable-alarm":     {"DisableActionRequests"},
			"batch-enable-alarm":      {"EnableActionRequests"},
			"batch-put-message":       {"Messages"},
			"batch-reset-alarm":       {"ResetActionRequests"},
			"batch-snooze-alarm":      {"SnoozeActionRequests"},
			"batch-update-detector":   {"Detectors"},
			"describe-alarm":          {"AlarmModelName", "KeyValue"},
			"describe-detector":       {"DetectorModelName", "KeyValue"},
			"list-alarms":             {"AlarmModelName", "MaxResults", "NextToken"},
			"list-detectors":          {"DetectorModelName", "MaxResults", "NextToken", "StateName"},
		},
		OperationInputTypes: map[string]map[string]string{
			"batch-acknowledge-alarm": {"AcknowledgeActionRequests": "[]types.AcknowledgeAlarmActionRequest"},
			"batch-delete-detector":   {"Detectors": "[]types.DeleteDetectorRequest"},
			"batch-disable-alarm":     {"DisableActionRequests": "[]types.DisableAlarmActionRequest"},
			"batch-enable-alarm":      {"EnableActionRequests": "[]types.EnableAlarmActionRequest"},
			"batch-put-message":       {"Messages": "[]types.Message"},
			"batch-reset-alarm":       {"ResetActionRequests": "[]types.ResetAlarmActionRequest"},
			"batch-snooze-alarm":      {"SnoozeActionRequests": "[]types.SnoozeAlarmActionRequest"},
			"batch-update-detector":   {"Detectors": "[]types.UpdateDetectorRequest"},
			"describe-alarm":          {"AlarmModelName": "*string", "KeyValue": "*string"},
			"describe-detector":       {"DetectorModelName": "*string", "KeyValue": "*string"},
			"list-alarms":             {"AlarmModelName": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-detectors":          {"DetectorModelName": "*string", "MaxResults": "*int32", "NextToken": "*string", "StateName": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"batch-acknowledge-alarm": {"AcknowledgeActionRequests"},
			"batch-delete-detector":   {"Detectors"},
			"batch-disable-alarm":     {"DisableActionRequests"},
			"batch-enable-alarm":      {"EnableActionRequests"},
			"batch-put-message":       {"Messages"},
			"batch-reset-alarm":       {"ResetActionRequests"},
			"batch-snooze-alarm":      {"SnoozeActionRequests"},
			"batch-update-detector":   {"Detectors"},
			"describe-alarm":          {"AlarmModelName"},
			"describe-detector":       {"DetectorModelName"},
			"list-alarms":             {"AlarmModelName"},
			"list-detectors":          {"DetectorModelName"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("ioteventsdata", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
