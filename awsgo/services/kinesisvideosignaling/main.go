package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/kinesisvideosignaling/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"get-ice-server-config", "send-alexa-offer-to-master"},
		OperationSet: map[string]bool{"get-ice-server-config": true, "send-alexa-offer-to-master": true},
		OperationInputs: map[string][]string{
			"get-ice-server-config":      {"ChannelARN", "ClientId", "Service", "Username"},
			"send-alexa-offer-to-master": {"ChannelARN", "MessagePayload", "SenderClientId"},
		},
		OperationInputTypes: map[string]map[string]string{
			"get-ice-server-config":      {"ChannelARN": "*string", "ClientId": "*string", "Service": "types.Service", "Username": "*string"},
			"send-alexa-offer-to-master": {"ChannelARN": "*string", "MessagePayload": "*string", "SenderClientId": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"get-ice-server-config":      {"ChannelARN"},
			"send-alexa-offer-to-master": {"ChannelARN", "MessagePayload", "SenderClientId"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("kinesisvideosignaling", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
