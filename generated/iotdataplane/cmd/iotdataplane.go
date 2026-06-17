package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iotdataplane"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// iotdataplaneCmd represents the iotdataplane command
var _iotdataplaneCmd = &cobra.Command{
	Use:   "iotdataplane",
	Short: "AWS iotdataplane CLI",
	Run: func(cmd *cobra.Command, args []string) {
		_awsOutput = resolveAWSOutput(_awsProfile, cmd.Flags().Changed("output"))
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := iotdataplane.NewFromConfig(cfg)
		if _iotdataplaneDeleteConnection {
			iotdataplane_DeleteConnection(cfg, client)
			return
		}
		if _iotdataplaneDeleteThingShadow {
			iotdataplane_DeleteThingShadow(cfg, client)
			return
		}
		if _iotdataplaneGetRetainedMessage {
			iotdataplane_GetRetainedMessage(cfg, client)
			return
		}
		if _iotdataplaneGetThingShadow {
			iotdataplane_GetThingShadow(cfg, client)
			return
		}
		if _iotdataplaneListNamedShadowsForThing {
			iotdataplane_ListNamedShadowsForThing(cfg, client)
			return
		}
		if _iotdataplaneListRetainedMessages {
			iotdataplane_ListRetainedMessages(cfg, client)
			return
		}
		if _iotdataplanePublish {
			iotdataplane_Publish(cfg, client)
			return
		}
		if _iotdataplaneUpdateThingShadow {
			iotdataplane_UpdateThingShadow(cfg, client)
			return
		}

	},
}

var (
	_iotdataplaneDeleteConnection         bool
	_iotdataplaneDeleteThingShadow        bool
	_iotdataplaneGetRetainedMessage       bool
	_iotdataplaneGetThingShadow           bool
	_iotdataplaneListNamedShadowsForThing bool
	_iotdataplaneListRetainedMessages     bool
	_iotdataplanePublish                  bool
	_iotdataplaneUpdateThingShadow        bool

	_iotdataplaneCleanSession           string
	_iotdataplaneClientId               string
	_iotdataplaneContentType            string
	_iotdataplaneCorrelationData        string
	_iotdataplaneMaxResults             string
	_iotdataplaneMessageExpiry          string
	_iotdataplaneNextToken              string
	_iotdataplanePageSize               string
	_iotdataplanePayload                string
	_iotdataplanePayloadFormatIndicator string
	_iotdataplanePreventWillMessage     string
	_iotdataplaneQos                    string
	_iotdataplaneResponseTopic          string
	_iotdataplaneRetain                 string
	_iotdataplaneShadowName             string
	_iotdataplaneThingName              string
	_iotdataplaneTopic                  string
	_iotdataplaneUserProperties         string
)

// Disconnects a connected MQTT client from Amazon Web Services IoT Core. When you
// disconnect a client, Amazon Web Services IoT Core closes the client's network
// connection and optionally cleans the session state.
func iotdataplane_DeleteConnection(cfg aws.Config, client *iotdataplane.Client) {
	input := &iotdataplane.DeleteConnectionInput{
		// ClientId: *string, // Required
	}

	if len(_iotdataplaneClientId) > 0 {
		input.ClientId = aws.String(_iotdataplaneClientId)
	}
	if len(_iotdataplaneCleanSession) > 0 {
		if err := assignInputField(input, "CleanSession", _iotdataplaneCleanSession); err != nil {
			log.Errorf("invalid --clean-session: %s", err.Error())
			return
		}
	}
	if len(_iotdataplanePreventWillMessage) > 0 {
		if err := assignInputField(input, "PreventWillMessage", _iotdataplanePreventWillMessage); err != nil {
			log.Errorf("invalid --prevent-will-message: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the shadow for the specified thing.
// Requires permission to access the [DeleteThingShadow] action.
//
// For more information, see [DeleteThingShadow] in the IoT Developer Guide.
//
// [DeleteThingShadow]: http://docs.aws.amazon.com/iot/latest/developerguide/API_DeleteThingShadow.html
func iotdataplane_DeleteThingShadow(cfg aws.Config, client *iotdataplane.Client) {
	input := &iotdataplane.DeleteThingShadowInput{
		// ThingName: *string, // Required
	}

	if len(_iotdataplaneThingName) > 0 {
		input.ThingName = aws.String(_iotdataplaneThingName)
	}
	if len(_iotdataplaneShadowName) > 0 {
		input.ShadowName = aws.String(_iotdataplaneShadowName)
	}

	if resp, err := client.DeleteThingShadow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the details of a single retained message for the specified topic.
// This action returns the message payload of the retained message, which can
// incur messaging costs. To list only the topic names of the retained messages,
// call [ListRetainedMessages].
//
// Requires permission to access the [GetRetainedMessage] action.
//
// For more information about messaging costs, see [Amazon Web Services IoT Core pricing - Messaging].
//
// [GetRetainedMessage]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html
// [Amazon Web Services IoT Core pricing - Messaging]: http://aws.amazon.com/iot-core/pricing/#Messaging
// [ListRetainedMessages]: https://docs.aws.amazon.com/iot/latest/apireference/API_iotdata_ListRetainedMessages.html
func iotdataplane_GetRetainedMessage(cfg aws.Config, client *iotdataplane.Client) {
	input := &iotdataplane.GetRetainedMessageInput{
		// Topic: *string, // Required
	}

	if len(_iotdataplaneTopic) > 0 {
		input.Topic = aws.String(_iotdataplaneTopic)
	}

	if resp, err := client.GetRetainedMessage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the shadow for the specified thing.
// Requires permission to access the [GetThingShadow] action.
//
// For more information, see [GetThingShadow] in the IoT Developer Guide.
//
// [GetThingShadow]: http://docs.aws.amazon.com/iot/latest/developerguide/API_GetThingShadow.html
func iotdataplane_GetThingShadow(cfg aws.Config, client *iotdataplane.Client) {
	input := &iotdataplane.GetThingShadowInput{
		// ThingName: *string, // Required
	}

	if len(_iotdataplaneThingName) > 0 {
		input.ThingName = aws.String(_iotdataplaneThingName)
	}
	if len(_iotdataplaneShadowName) > 0 {
		input.ShadowName = aws.String(_iotdataplaneShadowName)
	}

	if resp, err := client.GetThingShadow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the shadows for the specified thing.
// Requires permission to access the [ListNamedShadowsForThing] action.
//
// [ListNamedShadowsForThing]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iotdataplane_ListNamedShadowsForThing(cfg aws.Config, client *iotdataplane.Client) {
	input := &iotdataplane.ListNamedShadowsForThingInput{
		// ThingName: *string, // Required
	}

	if len(_iotdataplaneThingName) > 0 {
		input.ThingName = aws.String(_iotdataplaneThingName)
	}
	if len(_iotdataplaneNextToken) > 0 {
		input.NextToken = aws.String(_iotdataplaneNextToken)
	}
	if len(_iotdataplanePageSize) > 0 {
		if err := assignInputField(input, "PageSize", _iotdataplanePageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListNamedShadowsForThing(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists summary information about the retained messages stored for the account.
// This action returns only the topic names of the retained messages. It doesn't
// return any message payloads. Although this action doesn't return a message
// payload, it can still incur messaging costs.
//
// To get the message payload of a retained message, call [GetRetainedMessage] with the topic name of
// the retained message.
//
// Requires permission to access the [ListRetainedMessages] action.
//
// For more information about messaging costs, see [Amazon Web Services IoT Core pricing - Messaging].
//
// [GetRetainedMessage]: https://docs.aws.amazon.com/iot/latest/apireference/API_iotdata_GetRetainedMessage.html
// [Amazon Web Services IoT Core pricing - Messaging]: http://aws.amazon.com/iot-core/pricing/#Messaging
// [ListRetainedMessages]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html
func iotdataplane_ListRetainedMessages(cfg aws.Config, client *iotdataplane.Client) {
	input := &iotdataplane.ListRetainedMessagesInput{}

	if len(_iotdataplaneMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotdataplaneMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotdataplaneNextToken) > 0 {
		input.NextToken = aws.String(_iotdataplaneNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRetainedMessages(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotdataplane.ListRetainedMessagesOutput
	p := iotdataplane.NewListRetainedMessagesPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Publishes an MQTT message.
// Requires permission to access the [Publish] action.
//
// For more information about MQTT messages, see [MQTT Protocol] in the IoT Developer Guide.
//
// For more information about messaging costs, see [Amazon Web Services IoT Core pricing - Messaging].
//
// [MQTT Protocol]: http://docs.aws.amazon.com/iot/latest/developerguide/mqtt.html
// [Amazon Web Services IoT Core pricing - Messaging]: http://aws.amazon.com/iot-core/pricing/#Messaging
// [Publish]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iotdataplane_Publish(cfg aws.Config, client *iotdataplane.Client) {
	input := &iotdataplane.PublishInput{
		// Topic: *string, // Required
	}

	if len(_iotdataplaneTopic) > 0 {
		input.Topic = aws.String(_iotdataplaneTopic)
	}
	if len(_iotdataplaneContentType) > 0 {
		input.ContentType = aws.String(_iotdataplaneContentType)
	}
	if len(_iotdataplaneCorrelationData) > 0 {
		input.CorrelationData = aws.String(_iotdataplaneCorrelationData)
	}
	if len(_iotdataplaneMessageExpiry) > 0 {
		if err := assignInputField(input, "MessageExpiry", _iotdataplaneMessageExpiry); err != nil {
			log.Errorf("invalid --message-expiry: %s", err.Error())
			return
		}
	}
	if len(_iotdataplanePayload) > 0 {
		if err := assignInputField(input, "Payload", _iotdataplanePayload); err != nil {
			log.Errorf("invalid --payload: %s", err.Error())
			return
		}
	}
	if len(_iotdataplanePayloadFormatIndicator) > 0 {
		if err := assignInputField(input, "PayloadFormatIndicator", _iotdataplanePayloadFormatIndicator); err != nil {
			log.Errorf("invalid --payload-format-indicator: %s", err.Error())
			return
		}
	}
	if len(_iotdataplaneQos) > 0 {
		if err := assignInputField(input, "Qos", _iotdataplaneQos); err != nil {
			log.Errorf("invalid --qos: %s", err.Error())
			return
		}
	}
	if len(_iotdataplaneResponseTopic) > 0 {
		input.ResponseTopic = aws.String(_iotdataplaneResponseTopic)
	}
	if len(_iotdataplaneRetain) > 0 {
		if err := assignInputField(input, "Retain", _iotdataplaneRetain); err != nil {
			log.Errorf("invalid --retain: %s", err.Error())
			return
		}
	}
	if len(_iotdataplaneUserProperties) > 0 {
		input.UserProperties = aws.String(_iotdataplaneUserProperties)
	}

	if resp, err := client.Publish(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the shadow for the specified thing.
// Requires permission to access the [UpdateThingShadow] action.
//
// For more information, see [UpdateThingShadow] in the IoT Developer Guide.
//
// [UpdateThingShadow]: http://docs.aws.amazon.com/iot/latest/developerguide/API_UpdateThingShadow.html
func iotdataplane_UpdateThingShadow(cfg aws.Config, client *iotdataplane.Client) {
	input := &iotdataplane.UpdateThingShadowInput{
		// Payload: []byte, // Required
		// ThingName: *string, // Required
	}

	if len(_iotdataplanePayload) > 0 {
		if err := assignInputField(input, "Payload", _iotdataplanePayload); err != nil {
			log.Errorf("invalid --payload: %s", err.Error())
			return
		}
	}
	if len(_iotdataplaneThingName) > 0 {
		input.ThingName = aws.String(_iotdataplaneThingName)
	}
	if len(_iotdataplaneShadowName) > 0 {
		input.ShadowName = aws.String(_iotdataplaneShadowName)
	}

	if resp, err := client.UpdateThingShadow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_iotdataplaneCmd)
	_iotdataplaneCmd.Flags().SortFlags = false

	_iotdataplaneCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_iotdataplaneCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_iotdataplaneCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_iotdataplaneCmd.Flags().StringVarP(&_iotdataplaneCleanSession, "clean-session", "", "", "Clean Session")
	_iotdataplaneCmd.Flags().StringVarP(&_iotdataplaneClientId, "client-id", "", "", "Client ID")
	_iotdataplaneCmd.Flags().StringVarP(&_iotdataplaneContentType, "content-type", "", "", "Content Type")
	_iotdataplaneCmd.Flags().StringVarP(&_iotdataplaneCorrelationData, "correlation-data", "", "", "Correlation Data")
	_iotdataplaneCmd.Flags().StringVarP(&_iotdataplaneMaxResults, "max-results", "", "", "Max Results")
	_iotdataplaneCmd.Flags().StringVarP(&_iotdataplaneMessageExpiry, "message-expiry", "", "", "Message Expiry")
	_iotdataplaneCmd.Flags().StringVarP(&_iotdataplaneNextToken, "next-token", "", "", "Next Token")
	_iotdataplaneCmd.Flags().StringVarP(&_iotdataplanePageSize, "page-size", "", "", "Page Size")
	_iotdataplaneCmd.Flags().StringVarP(&_iotdataplanePayload, "payload", "", "", "Payload")
	_iotdataplaneCmd.Flags().StringVarP(&_iotdataplanePayloadFormatIndicator, "payload-format-indicator", "", "", "Payload Format Indicator")
	_iotdataplaneCmd.Flags().StringVarP(&_iotdataplanePreventWillMessage, "prevent-will-message", "", "", "Prevent Will Message")
	_iotdataplaneCmd.Flags().StringVarP(&_iotdataplaneQos, "qos", "", "", "Qos")
	_iotdataplaneCmd.Flags().StringVarP(&_iotdataplaneResponseTopic, "response-topic", "", "", "Response Topic")
	_iotdataplaneCmd.Flags().StringVarP(&_iotdataplaneRetain, "retain", "", "", "Retain")
	_iotdataplaneCmd.Flags().StringVarP(&_iotdataplaneShadowName, "shadow-name", "", "", "Shadow Name")
	_iotdataplaneCmd.Flags().StringVarP(&_iotdataplaneThingName, "thing-name", "", "", "Thing Name")
	_iotdataplaneCmd.Flags().StringVarP(&_iotdataplaneTopic, "topic", "", "", "Topic")
	_iotdataplaneCmd.Flags().StringVarP(&_iotdataplaneUserProperties, "user-properties", "", "", "User Properties")

	_iotdataplaneCmd.Flags().BoolVarP(&_iotdataplaneDeleteConnection, "delete-connection", "", false, "Delete Connection")
	_iotdataplaneCmd.Flags().BoolVarP(&_iotdataplaneDeleteThingShadow, "delete-thing-shadow", "", false, "Delete Thing Shadow")
	_iotdataplaneCmd.Flags().BoolVarP(&_iotdataplaneGetRetainedMessage, "get-retained-message", "", false, "Get Retained Message")
	_iotdataplaneCmd.Flags().BoolVarP(&_iotdataplaneGetThingShadow, "get-thing-shadow", "", false, "Get Thing Shadow")
	_iotdataplaneCmd.Flags().BoolVarP(&_iotdataplaneListNamedShadowsForThing, "list-named-shadows-for-thing", "", false, "List Named Shadows For Thing")
	_iotdataplaneCmd.Flags().BoolVarP(&_iotdataplaneListRetainedMessages, "list-retained-messages", "", false, "List Retained Messages")
	_iotdataplaneCmd.Flags().BoolVarP(&_iotdataplanePublish, "publish", "", false, "Publish")
	_iotdataplaneCmd.Flags().BoolVarP(&_iotdataplaneUpdateThingShadow, "update-thing-shadow", "", false, "Update Thing Shadow")

}
