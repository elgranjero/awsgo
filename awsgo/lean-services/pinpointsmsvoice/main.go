package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoice"
)

var fields_create_configuration_set = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: false},
}

var fields_create_configuration_set_event_destination = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
	{Name: "EventDestination", Flag: "event-destination", Type: "*types.EventDestinationDefinition", Required: false},
	{Name: "EventDestinationName", Flag: "event-destination-name", Type: "*string", Required: false},
}

var fields_delete_configuration_set = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
}

var fields_delete_configuration_set_event_destination = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
	{Name: "EventDestinationName", Flag: "event-destination-name", Type: "*string", Required: true},
}

var fields_get_configuration_set_event_destinations = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
}

var fields_list_configuration_sets = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*string", Required: false},
}

var fields_send_voice_message = []leanruntime.Field{
	{Name: "CallerId", Flag: "caller-id", Type: "*string", Required: false},
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: false},
	{Name: "Content", Flag: "content", Type: "*types.VoiceMessageContent", Required: false},
	{Name: "DestinationPhoneNumber", Flag: "destination-phone-number", Type: "*string", Required: false},
	{Name: "OriginationPhoneNumber", Flag: "origination-phone-number", Type: "*string", Required: false},
}

var fields_update_configuration_set_event_destination = []leanruntime.Field{
	{Name: "ConfigurationSetName", Flag: "configuration-set-name", Type: "*string", Required: true},
	{Name: "EventDestination", Flag: "event-destination", Type: "*types.EventDestinationDefinition", Required: false},
	{Name: "EventDestinationName", Flag: "event-destination-name", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-configuration-set": {
			Name:   "create-configuration-set",
			Fields: fields_create_configuration_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateConfigurationSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_configuration_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateConfigurationSet(ctx, input)
			},
		},
		"create-configuration-set-event-destination": {
			Name:   "create-configuration-set-event-destination",
			Fields: fields_create_configuration_set_event_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateConfigurationSetEventDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_configuration_set_event_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateConfigurationSetEventDestination(ctx, input)
			},
		},
		"delete-configuration-set": {
			Name:   "delete-configuration-set",
			Fields: fields_delete_configuration_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConfigurationSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_configuration_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConfigurationSet(ctx, input)
			},
		},
		"delete-configuration-set-event-destination": {
			Name:   "delete-configuration-set-event-destination",
			Fields: fields_delete_configuration_set_event_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConfigurationSetEventDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_configuration_set_event_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConfigurationSetEventDestination(ctx, input)
			},
		},
		"get-configuration-set-event-destinations": {
			Name:   "get-configuration-set-event-destinations",
			Fields: fields_get_configuration_set_event_destinations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConfigurationSetEventDestinationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_configuration_set_event_destinations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConfigurationSetEventDestinations(ctx, input)
			},
		},
		"list-configuration-sets": {
			Name:   "list-configuration-sets",
			Fields: fields_list_configuration_sets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConfigurationSetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_configuration_sets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListConfigurationSets(ctx, input)
			},
		},
		"send-voice-message": {
			Name:   "send-voice-message",
			Fields: fields_send_voice_message,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendVoiceMessageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_voice_message, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendVoiceMessage(ctx, input)
			},
		},
		"update-configuration-set-event-destination": {
			Name:   "update-configuration-set-event-destination",
			Fields: fields_update_configuration_set_event_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateConfigurationSetEventDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_configuration_set_event_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateConfigurationSetEventDestination(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("pinpointsmsvoice", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
