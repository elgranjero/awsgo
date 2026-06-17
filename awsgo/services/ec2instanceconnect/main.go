package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/ec2instanceconnect/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"send-serial-console-ssh-public-key", "send-ssh-public-key"},
		OperationSet: map[string]bool{"send-serial-console-ssh-public-key": true, "send-ssh-public-key": true},
		OperationInputs: map[string][]string{
			"send-serial-console-ssh-public-key": {"InstanceId", "SSHPublicKey", "SerialPort"},
			"send-ssh-public-key":                {"AvailabilityZone", "InstanceId", "InstanceOSUser", "SSHPublicKey"},
		},
		OperationInputTypes: map[string]map[string]string{
			"send-serial-console-ssh-public-key": {"InstanceId": "*string", "SSHPublicKey": "*string", "SerialPort": "int32"},
			"send-ssh-public-key":                {"AvailabilityZone": "*string", "InstanceId": "*string", "InstanceOSUser": "*string", "SSHPublicKey": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"send-serial-console-ssh-public-key": {"InstanceId", "SSHPublicKey"},
			"send-ssh-public-key":                {"InstanceId", "InstanceOSUser", "SSHPublicKey"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("ec2instanceconnect", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
