package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/kinesisvideowebrtcstorage/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"join-storage-session", "join-storage-session-as-viewer"},
		OperationSet: map[string]bool{"join-storage-session": true, "join-storage-session-as-viewer": true},
		OperationInputs: map[string][]string{
			"join-storage-session":           {"ChannelArn"},
			"join-storage-session-as-viewer": {"ChannelArn", "ClientId"},
		},
		OperationInputTypes: map[string]map[string]string{
			"join-storage-session":           {"ChannelArn": "*string"},
			"join-storage-session-as-viewer": {"ChannelArn": "*string", "ClientId": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"join-storage-session":           {"ChannelArn"},
			"join-storage-session-as-viewer": {"ChannelArn", "ClientId"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("kinesisvideowebrtcstorage", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
