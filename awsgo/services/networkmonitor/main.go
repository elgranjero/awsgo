package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/networkmonitor/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-monitor", "create-probe", "delete-monitor", "delete-probe", "get-monitor", "get-probe", "list-monitors", "list-tags-for-resource", "tag-resource", "untag-resource", "update-monitor", "update-probe"},
		OperationSet: map[string]bool{"create-monitor": true, "create-probe": true, "delete-monitor": true, "delete-probe": true, "get-monitor": true, "get-probe": true, "list-monitors": true, "list-tags-for-resource": true, "tag-resource": true, "untag-resource": true, "update-monitor": true, "update-probe": true},
		OperationInputs: map[string][]string{
			"create-monitor":         {"AggregationPeriod", "ClientToken", "MonitorName", "Probes", "Tags"},
			"create-probe":           {"ClientToken", "MonitorName", "Probe", "Tags"},
			"delete-monitor":         {"MonitorName"},
			"delete-probe":           {"MonitorName", "ProbeId"},
			"get-monitor":            {"MonitorName"},
			"get-probe":              {"MonitorName", "ProbeId"},
			"list-monitors":          {"MaxResults", "NextToken", "State"},
			"list-tags-for-resource": {"ResourceArn"},
			"tag-resource":           {"ResourceArn", "Tags"},
			"untag-resource":         {"ResourceArn", "TagKeys"},
			"update-monitor":         {"AggregationPeriod", "MonitorName"},
			"update-probe":           {"Destination", "DestinationPort", "MonitorName", "PacketSize", "ProbeId", "Protocol", "State"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-monitor":         {"AggregationPeriod": "*int64", "ClientToken": "*string", "MonitorName": "*string", "Probes": "[]types.CreateMonitorProbeInput", "Tags": "map[string]string"},
			"create-probe":           {"ClientToken": "*string", "MonitorName": "*string", "Probe": "*types.ProbeInput", "Tags": "map[string]string"},
			"delete-monitor":         {"MonitorName": "*string"},
			"delete-probe":           {"MonitorName": "*string", "ProbeId": "*string"},
			"get-monitor":            {"MonitorName": "*string"},
			"get-probe":              {"MonitorName": "*string", "ProbeId": "*string"},
			"list-monitors":          {"MaxResults": "*int32", "NextToken": "*string", "State": "*string"},
			"list-tags-for-resource": {"ResourceArn": "*string"},
			"tag-resource":           {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":         {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-monitor":         {"AggregationPeriod": "*int64", "MonitorName": "*string"},
			"update-probe":           {"Destination": "*string", "DestinationPort": "*int32", "MonitorName": "*string", "PacketSize": "*int32", "ProbeId": "*string", "Protocol": "types.Protocol", "State": "types.ProbeState"},
		},
		OperationInputRequired: map[string][]string{
			"create-monitor":         {"MonitorName"},
			"create-probe":           {"MonitorName", "Probe"},
			"delete-monitor":         {"MonitorName"},
			"delete-probe":           {"MonitorName", "ProbeId"},
			"get-monitor":            {"MonitorName"},
			"get-probe":              {"MonitorName", "ProbeId"},
			"list-monitors":          {},
			"list-tags-for-resource": {"ResourceArn"},
			"tag-resource":           {"ResourceArn", "Tags"},
			"untag-resource":         {"ResourceArn", "TagKeys"},
			"update-monitor":         {"AggregationPeriod", "MonitorName"},
			"update-probe":           {"MonitorName", "ProbeId"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("networkmonitor", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
