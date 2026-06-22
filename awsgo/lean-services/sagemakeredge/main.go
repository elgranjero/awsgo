package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/sagemakeredge"
)

var fields_get_deployments = []leanruntime.Field{
	{Name: "DeviceFleetName", Flag: "device-fleet-name", Type: "*string", Required: true},
	{Name: "DeviceName", Flag: "device-name", Type: "*string", Required: true},
}

var fields_get_device_registration = []leanruntime.Field{
	{Name: "DeviceFleetName", Flag: "device-fleet-name", Type: "*string", Required: true},
	{Name: "DeviceName", Flag: "device-name", Type: "*string", Required: true},
}

var fields_send_heartbeat = []leanruntime.Field{
	{Name: "AgentMetrics", Flag: "agent-metrics", Type: "[]types.EdgeMetric", Required: false},
	{Name: "AgentVersion", Flag: "agent-version", Type: "*string", Required: true},
	{Name: "DeploymentResult", Flag: "deployment-result", Type: "*types.DeploymentResult", Required: false},
	{Name: "DeviceFleetName", Flag: "device-fleet-name", Type: "*string", Required: true},
	{Name: "DeviceName", Flag: "device-name", Type: "*string", Required: true},
	{Name: "Models", Flag: "models", Type: "[]types.Model", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"get-deployments": {
			Name:   "get-deployments",
			Fields: fields_get_deployments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDeploymentsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_deployments, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDeployments(ctx, input)
			},
		},
		"get-device-registration": {
			Name:   "get-device-registration",
			Fields: fields_get_device_registration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDeviceRegistrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_device_registration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDeviceRegistration(ctx, input)
			},
		},
		"send-heartbeat": {
			Name:   "send-heartbeat",
			Fields: fields_send_heartbeat,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendHeartbeatInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_heartbeat, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendHeartbeat(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("sagemakeredge", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
