package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/ec2instanceconnect"
)

var fields_send_serial_console_ssh_public_key = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "SSHPublicKey", Flag: "ssh-public-key", Type: "*string", Required: true},
	{Name: "SerialPort", Flag: "serial-port", Type: "int32", Required: false},
}

var fields_send_ssh_public_key = []leanruntime.Field{
	{Name: "AvailabilityZone", Flag: "availability-zone", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "InstanceOSUser", Flag: "instance-os-user", Type: "*string", Required: true},
	{Name: "SSHPublicKey", Flag: "ssh-public-key", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"send-serial-console-ssh-public-key": {
			Name:   "send-serial-console-ssh-public-key",
			Fields: fields_send_serial_console_ssh_public_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendSerialConsoleSSHPublicKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_serial_console_ssh_public_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendSerialConsoleSSHPublicKey(ctx, input)
			},
		},
		"send-ssh-public-key": {
			Name:   "send-ssh-public-key",
			Fields: fields_send_ssh_public_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendSSHPublicKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_ssh_public_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendSSHPublicKey(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("ec2instanceconnect", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
