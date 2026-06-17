package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/sagemakeredge/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"get-deployments", "get-device-registration", "send-heartbeat"},
		OperationSet: map[string]bool{"get-deployments": true, "get-device-registration": true, "send-heartbeat": true},
		OperationInputs: map[string][]string{
			"get-deployments":         {"DeviceFleetName", "DeviceName"},
			"get-device-registration": {"DeviceFleetName", "DeviceName"},
			"send-heartbeat":          {"AgentMetrics", "AgentVersion", "DeploymentResult", "DeviceFleetName", "DeviceName", "Models"},
		},
		OperationInputTypes: map[string]map[string]string{
			"get-deployments":         {"DeviceFleetName": "*string", "DeviceName": "*string"},
			"get-device-registration": {"DeviceFleetName": "*string", "DeviceName": "*string"},
			"send-heartbeat":          {"AgentMetrics": "[]types.EdgeMetric", "AgentVersion": "*string", "DeploymentResult": "*types.DeploymentResult", "DeviceFleetName": "*string", "DeviceName": "*string", "Models": "[]types.Model"},
		},
		OperationInputRequired: map[string][]string{
			"get-deployments":         {"DeviceFleetName", "DeviceName"},
			"get-device-registration": {"DeviceFleetName", "DeviceName"},
			"send-heartbeat":          {"AgentVersion", "DeviceFleetName", "DeviceName"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("sagemakeredge", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
