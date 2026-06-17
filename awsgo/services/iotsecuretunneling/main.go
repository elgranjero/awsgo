package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/iotsecuretunneling/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"close-tunnel", "describe-tunnel", "list-tags-for-resource", "list-tunnels", "open-tunnel", "rotate-tunnel-access-token", "tag-resource", "untag-resource"},
		OperationSet: map[string]bool{"close-tunnel": true, "describe-tunnel": true, "list-tags-for-resource": true, "list-tunnels": true, "open-tunnel": true, "rotate-tunnel-access-token": true, "tag-resource": true, "untag-resource": true},
		OperationInputs: map[string][]string{
			"close-tunnel":               {"Delete", "TunnelId"},
			"describe-tunnel":            {"TunnelId"},
			"list-tags-for-resource":     {"ResourceArn"},
			"list-tunnels":               {"MaxResults", "NextToken", "ThingName"},
			"open-tunnel":                {"Description", "DestinationConfig", "Tags", "TimeoutConfig"},
			"rotate-tunnel-access-token": {"ClientMode", "DestinationConfig", "TunnelId"},
			"tag-resource":               {"ResourceArn", "Tags"},
			"untag-resource":             {"ResourceArn", "TagKeys"},
		},
		OperationInputTypes: map[string]map[string]string{
			"close-tunnel":               {"Delete": "*bool", "TunnelId": "*string"},
			"describe-tunnel":            {"TunnelId": "*string"},
			"list-tags-for-resource":     {"ResourceArn": "*string"},
			"list-tunnels":               {"MaxResults": "*int32", "NextToken": "*string", "ThingName": "*string"},
			"open-tunnel":                {"Description": "*string", "DestinationConfig": "*types.DestinationConfig", "Tags": "[]types.Tag", "TimeoutConfig": "*types.TimeoutConfig"},
			"rotate-tunnel-access-token": {"ClientMode": "types.ClientMode", "DestinationConfig": "*types.DestinationConfig", "TunnelId": "*string"},
			"tag-resource":               {"ResourceArn": "*string", "Tags": "[]types.Tag"},
			"untag-resource":             {"ResourceArn": "*string", "TagKeys": "[]string"},
		},
		OperationInputRequired: map[string][]string{
			"close-tunnel":               {"TunnelId"},
			"describe-tunnel":            {"TunnelId"},
			"list-tags-for-resource":     {"ResourceArn"},
			"list-tunnels":               {},
			"open-tunnel":                {},
			"rotate-tunnel-access-token": {"ClientMode", "TunnelId"},
			"tag-resource":               {"ResourceArn", "Tags"},
			"untag-resource":             {"ResourceArn", "TagKeys"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("iotsecuretunneling", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
