package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoice"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// pinpointsmsvoiceCmd represents the pinpointsmsvoice command
var _pinpointsmsvoiceCmd = &cobra.Command{
	Use:   "pinpointsmsvoice",
	Short: "AWS pinpointsmsvoice CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := pinpointsmsvoice.NewFromConfig(cfg)
		if _pinpointsmsvoiceCreateConfigurationSet {
			pinpointsmsvoice_CreateConfigurationSet(cfg, client)
			return
		}
		if _pinpointsmsvoiceCreateConfigurationSetEventDestination {
			pinpointsmsvoice_CreateConfigurationSetEventDestination(cfg, client)
			return
		}
		if _pinpointsmsvoiceDeleteConfigurationSet {
			pinpointsmsvoice_DeleteConfigurationSet(cfg, client)
			return
		}
		if _pinpointsmsvoiceDeleteConfigurationSetEventDestination {
			pinpointsmsvoice_DeleteConfigurationSetEventDestination(cfg, client)
			return
		}
		if _pinpointsmsvoiceGetConfigurationSetEventDestinations {
			pinpointsmsvoice_GetConfigurationSetEventDestinations(cfg, client)
			return
		}
		if _pinpointsmsvoiceListConfigurationSets {
			pinpointsmsvoice_ListConfigurationSets(cfg, client)
			return
		}
		if _pinpointsmsvoiceSendVoiceMessage {
			pinpointsmsvoice_SendVoiceMessage(cfg, client)
			return
		}
		if _pinpointsmsvoiceUpdateConfigurationSetEventDestination {
			pinpointsmsvoice_UpdateConfigurationSetEventDestination(cfg, client)
			return
		}

	},
}

var (
	_pinpointsmsvoiceCreateConfigurationSet                 bool
	_pinpointsmsvoiceCreateConfigurationSetEventDestination bool
	_pinpointsmsvoiceDeleteConfigurationSet                 bool
	_pinpointsmsvoiceDeleteConfigurationSetEventDestination bool
	_pinpointsmsvoiceGetConfigurationSetEventDestinations   bool
	_pinpointsmsvoiceListConfigurationSets                  bool
	_pinpointsmsvoiceSendVoiceMessage                       bool
	_pinpointsmsvoiceUpdateConfigurationSetEventDestination bool

	_pinpointsmsvoiceCallerId               string
	_pinpointsmsvoiceConfigurationSetName   string
	_pinpointsmsvoiceContent                string
	_pinpointsmsvoiceDestinationPhoneNumber string
	_pinpointsmsvoiceEventDestination       string
	_pinpointsmsvoiceEventDestinationName   string
	_pinpointsmsvoiceNextToken              string
	_pinpointsmsvoiceOriginationPhoneNumber string
	_pinpointsmsvoicePageSize               string
)

// Create a new configuration set. After you create the configuration set, you can
// add one or more event destinations to it.
func pinpointsmsvoice_CreateConfigurationSet(cfg aws.Config, client *pinpointsmsvoice.Client) {
	input := &pinpointsmsvoice.CreateConfigurationSetInput{}

	if len(_pinpointsmsvoiceConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_pinpointsmsvoiceConfigurationSetName)
	}

	if resp, err := client.CreateConfigurationSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a new event destination in a configuration set.
func pinpointsmsvoice_CreateConfigurationSetEventDestination(cfg aws.Config, client *pinpointsmsvoice.Client) {
	input := &pinpointsmsvoice.CreateConfigurationSetEventDestinationInput{
		// ConfigurationSetName: *string, // Required
	}

	if len(_pinpointsmsvoiceConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_pinpointsmsvoiceConfigurationSetName)
	}
	if len(_pinpointsmsvoiceEventDestination) > 0 {
		if err := assignInputField(input, "EventDestination", _pinpointsmsvoiceEventDestination); err != nil {
			log.Errorf("invalid --event-destination: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoiceEventDestinationName) > 0 {
		input.EventDestinationName = aws.String(_pinpointsmsvoiceEventDestinationName)
	}

	if resp, err := client.CreateConfigurationSetEventDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing configuration set.
func pinpointsmsvoice_DeleteConfigurationSet(cfg aws.Config, client *pinpointsmsvoice.Client) {
	input := &pinpointsmsvoice.DeleteConfigurationSetInput{
		// ConfigurationSetName: *string, // Required
	}

	if len(_pinpointsmsvoiceConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_pinpointsmsvoiceConfigurationSetName)
	}

	if resp, err := client.DeleteConfigurationSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an event destination in a configuration set.
func pinpointsmsvoice_DeleteConfigurationSetEventDestination(cfg aws.Config, client *pinpointsmsvoice.Client) {
	input := &pinpointsmsvoice.DeleteConfigurationSetEventDestinationInput{
		// ConfigurationSetName: *string, // Required
		// EventDestinationName: *string, // Required
	}

	if len(_pinpointsmsvoiceConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_pinpointsmsvoiceConfigurationSetName)
	}
	if len(_pinpointsmsvoiceEventDestinationName) > 0 {
		input.EventDestinationName = aws.String(_pinpointsmsvoiceEventDestinationName)
	}

	if resp, err := client.DeleteConfigurationSetEventDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Obtain information about an event destination, including the types of events it
// reports, the Amazon Resource Name (ARN) of the destination, and the name of the
// event destination.
func pinpointsmsvoice_GetConfigurationSetEventDestinations(cfg aws.Config, client *pinpointsmsvoice.Client) {
	input := &pinpointsmsvoice.GetConfigurationSetEventDestinationsInput{
		// ConfigurationSetName: *string, // Required
	}

	if len(_pinpointsmsvoiceConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_pinpointsmsvoiceConfigurationSetName)
	}

	if resp, err := client.GetConfigurationSetEventDestinations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List all of the configuration sets associated with your Amazon Pinpoint account
// in the current region.
func pinpointsmsvoice_ListConfigurationSets(cfg aws.Config, client *pinpointsmsvoice.Client) {
	input := &pinpointsmsvoice.ListConfigurationSetsInput{}

	if len(_pinpointsmsvoiceNextToken) > 0 {
		input.NextToken = aws.String(_pinpointsmsvoiceNextToken)
	}
	if len(_pinpointsmsvoicePageSize) > 0 {
		input.PageSize = aws.String(_pinpointsmsvoicePageSize)
	}

	if resp, err := client.ListConfigurationSets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a new voice message and send it to a recipient's phone number.
func pinpointsmsvoice_SendVoiceMessage(cfg aws.Config, client *pinpointsmsvoice.Client) {
	input := &pinpointsmsvoice.SendVoiceMessageInput{}

	if len(_pinpointsmsvoiceCallerId) > 0 {
		input.CallerId = aws.String(_pinpointsmsvoiceCallerId)
	}
	if len(_pinpointsmsvoiceConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_pinpointsmsvoiceConfigurationSetName)
	}
	if len(_pinpointsmsvoiceContent) > 0 {
		if err := assignInputField(input, "Content", _pinpointsmsvoiceContent); err != nil {
			log.Errorf("invalid --content: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoiceDestinationPhoneNumber) > 0 {
		input.DestinationPhoneNumber = aws.String(_pinpointsmsvoiceDestinationPhoneNumber)
	}
	if len(_pinpointsmsvoiceOriginationPhoneNumber) > 0 {
		input.OriginationPhoneNumber = aws.String(_pinpointsmsvoiceOriginationPhoneNumber)
	}

	if resp, err := client.SendVoiceMessage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update an event destination in a configuration set. An event destination is a
// location that you publish information about your voice calls to. For example,
// you can log an event to an Amazon CloudWatch destination when a call fails.
func pinpointsmsvoice_UpdateConfigurationSetEventDestination(cfg aws.Config, client *pinpointsmsvoice.Client) {
	input := &pinpointsmsvoice.UpdateConfigurationSetEventDestinationInput{
		// ConfigurationSetName: *string, // Required
		// EventDestinationName: *string, // Required
	}

	if len(_pinpointsmsvoiceConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_pinpointsmsvoiceConfigurationSetName)
	}
	if len(_pinpointsmsvoiceEventDestinationName) > 0 {
		input.EventDestinationName = aws.String(_pinpointsmsvoiceEventDestinationName)
	}
	if len(_pinpointsmsvoiceEventDestination) > 0 {
		if err := assignInputField(input, "EventDestination", _pinpointsmsvoiceEventDestination); err != nil {
			log.Errorf("invalid --event-destination: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateConfigurationSetEventDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_pinpointsmsvoiceCmd)
	_pinpointsmsvoiceCmd.Flags().SortFlags = false

	_pinpointsmsvoiceCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_pinpointsmsvoiceCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_pinpointsmsvoiceCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_pinpointsmsvoiceCmd.Flags().StringVarP(&_pinpointsmsvoiceCallerId, "caller-id", "", "", "Caller ID")
	_pinpointsmsvoiceCmd.Flags().StringVarP(&_pinpointsmsvoiceConfigurationSetName, "configuration-set-name", "", "", "Configuration Set Name")
	_pinpointsmsvoiceCmd.Flags().StringVarP(&_pinpointsmsvoiceContent, "content", "", "", "Content")
	_pinpointsmsvoiceCmd.Flags().StringVarP(&_pinpointsmsvoiceDestinationPhoneNumber, "destination-phone-number", "", "", "Destination Phone Number")
	_pinpointsmsvoiceCmd.Flags().StringVarP(&_pinpointsmsvoiceEventDestination, "event-destination", "", "", "Event Destination")
	_pinpointsmsvoiceCmd.Flags().StringVarP(&_pinpointsmsvoiceEventDestinationName, "event-destination-name", "", "", "Event Destination Name")
	_pinpointsmsvoiceCmd.Flags().StringVarP(&_pinpointsmsvoiceNextToken, "next-token", "", "", "Next Token")
	_pinpointsmsvoiceCmd.Flags().StringVarP(&_pinpointsmsvoiceOriginationPhoneNumber, "origination-phone-number", "", "", "Origination Phone Number")
	_pinpointsmsvoiceCmd.Flags().StringVarP(&_pinpointsmsvoicePageSize, "page-size", "", "", "Page Size")

	_pinpointsmsvoiceCmd.Flags().BoolVarP(&_pinpointsmsvoiceCreateConfigurationSet, "create-configuration-set", "", false, "Create Configuration Set")
	_pinpointsmsvoiceCmd.Flags().BoolVarP(&_pinpointsmsvoiceCreateConfigurationSetEventDestination, "create-configuration-set-event-destination", "", false, "Create Configuration Set Event Destination")
	_pinpointsmsvoiceCmd.Flags().BoolVarP(&_pinpointsmsvoiceDeleteConfigurationSet, "delete-configuration-set", "", false, "Delete Configuration Set")
	_pinpointsmsvoiceCmd.Flags().BoolVarP(&_pinpointsmsvoiceDeleteConfigurationSetEventDestination, "delete-configuration-set-event-destination", "", false, "Delete Configuration Set Event Destination")
	_pinpointsmsvoiceCmd.Flags().BoolVarP(&_pinpointsmsvoiceGetConfigurationSetEventDestinations, "get-configuration-set-event-destinations", "", false, "Get Configuration Set Event Destinations")
	_pinpointsmsvoiceCmd.Flags().BoolVarP(&_pinpointsmsvoiceListConfigurationSets, "list-configuration-sets", "", false, "List Configuration Sets")
	_pinpointsmsvoiceCmd.Flags().BoolVarP(&_pinpointsmsvoiceSendVoiceMessage, "send-voice-message", "", false, "Send Voice Message")
	_pinpointsmsvoiceCmd.Flags().BoolVarP(&_pinpointsmsvoiceUpdateConfigurationSetEventDestination, "update-configuration-set-event-destination", "", false, "Update Configuration Set Event Destination")

}
