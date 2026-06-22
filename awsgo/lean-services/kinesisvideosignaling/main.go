package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/kinesisvideosignaling"
)

var fields_get_ice_server_config = []leanruntime.Field{
	{Name: "ChannelARN", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "ClientId", Flag: "client-id", Type: "*string", Required: false},
	{Name: "Service", Flag: "service", Type: "types.Service", Required: false},
	{Name: "Username", Flag: "username", Type: "*string", Required: false},
}

var fields_send_alexa_offer_to_master = []leanruntime.Field{
	{Name: "ChannelARN", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "MessagePayload", Flag: "message-payload", Type: "*string", Required: true},
	{Name: "SenderClientId", Flag: "sender-client-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"get-ice-server-config": {
			Name:   "get-ice-server-config",
			Fields: fields_get_ice_server_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIceServerConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_ice_server_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIceServerConfig(ctx, input)
			},
		},
		"send-alexa-offer-to-master": {
			Name:   "send-alexa-offer-to-master",
			Fields: fields_send_alexa_offer_to_master,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendAlexaOfferToMasterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_alexa_offer_to_master, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendAlexaOfferToMaster(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("kinesisvideosignaling", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
