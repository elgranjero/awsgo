package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideosignaling"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// kinesisvideosignalingCmd represents the kinesisvideosignaling command
var _kinesisvideosignalingCmd = &cobra.Command{
	Use:   "kinesisvideosignaling",
	Short: "AWS kinesisvideosignaling CLI",
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
		client := kinesisvideosignaling.NewFromConfig(cfg)
		if _kinesisvideosignalingGetIceServerConfig {
			kinesisvideosignaling_GetIceServerConfig(cfg, client)
			return
		}
		if _kinesisvideosignalingSendAlexaOfferToMaster {
			kinesisvideosignaling_SendAlexaOfferToMaster(cfg, client)
			return
		}

	},
}

var (
	_kinesisvideosignalingGetIceServerConfig     bool
	_kinesisvideosignalingSendAlexaOfferToMaster bool

	_kinesisvideosignalingChannelARN     string
	_kinesisvideosignalingClientId       string
	_kinesisvideosignalingMessagePayload string
	_kinesisvideosignalingSenderClientId string
	_kinesisvideosignalingService        string
	_kinesisvideosignalingUsername       string
)

// Gets the Interactive Connectivity Establishment (ICE) server configuration
// information, including URIs, username, and password which can be used to
// configure the WebRTC connection. The ICE component uses this configuration
// information to setup the WebRTC connection, including authenticating with the
// Traversal Using Relays around NAT (TURN) relay server.
//
// TURN is a protocol that is used to improve the connectivity of peer-to-peer
// applications. By providing a cloud-based relay service, TURN ensures that a
// connection can be established even when one or more peers are incapable of a
// direct peer-to-peer connection. For more information, see [A REST API For Access To TURN Services].
//
// You can invoke this API to establish a fallback mechanism in case either of the
// peers is unable to establish a direct peer-to-peer connection over a signaling
// channel. You must specify either a signaling channel ARN or the client ID in
// order to invoke this API.
//
// [A REST API For Access To TURN Services]: https://tools.ietf.org/html/draft-uberti-rtcweb-turn-rest-00
func kinesisvideosignaling_GetIceServerConfig(cfg aws.Config, client *kinesisvideosignaling.Client) {
	input := &kinesisvideosignaling.GetIceServerConfigInput{
		// ChannelARN: *string, // Required
	}

	if len(_kinesisvideosignalingChannelARN) > 0 {
		input.ChannelARN = aws.String(_kinesisvideosignalingChannelARN)
	}
	if len(_kinesisvideosignalingClientId) > 0 {
		input.ClientId = aws.String(_kinesisvideosignalingClientId)
	}
	if len(_kinesisvideosignalingService) > 0 {
		if err := assignInputField(input, "Service", _kinesisvideosignalingService); err != nil {
			log.Errorf("invalid --service: %s", err.Error())
			return
		}
	}
	if len(_kinesisvideosignalingUsername) > 0 {
		input.Username = aws.String(_kinesisvideosignalingUsername)
	}

	if resp, err := client.GetIceServerConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API allows you to connect WebRTC-enabled devices with Alexa display
// devices. When invoked, it sends the Alexa Session Description Protocol (SDP)
// offer to the master peer. The offer is delivered as soon as the master is
// connected to the specified signaling channel. This API returns the SDP answer
// from the connected master. If the master is not connected to the signaling
// channel, redelivery requests are made until the message expires.
func kinesisvideosignaling_SendAlexaOfferToMaster(cfg aws.Config, client *kinesisvideosignaling.Client) {
	input := &kinesisvideosignaling.SendAlexaOfferToMasterInput{
		// ChannelARN: *string, // Required
		// MessagePayload: *string, // Required
		// SenderClientId: *string, // Required
	}

	if len(_kinesisvideosignalingChannelARN) > 0 {
		input.ChannelARN = aws.String(_kinesisvideosignalingChannelARN)
	}
	if len(_kinesisvideosignalingMessagePayload) > 0 {
		input.MessagePayload = aws.String(_kinesisvideosignalingMessagePayload)
	}
	if len(_kinesisvideosignalingSenderClientId) > 0 {
		input.SenderClientId = aws.String(_kinesisvideosignalingSenderClientId)
	}

	if resp, err := client.SendAlexaOfferToMaster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_kinesisvideosignalingCmd)
	_kinesisvideosignalingCmd.Flags().SortFlags = false

	_kinesisvideosignalingCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_kinesisvideosignalingCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_kinesisvideosignalingCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_kinesisvideosignalingCmd.Flags().StringVarP(&_kinesisvideosignalingChannelARN, "channel-arn", "", "", "Channel ARN")
	_kinesisvideosignalingCmd.Flags().StringVarP(&_kinesisvideosignalingClientId, "client-id", "", "", "Client ID")
	_kinesisvideosignalingCmd.Flags().StringVarP(&_kinesisvideosignalingMessagePayload, "message-payload", "", "", "Message Payload")
	_kinesisvideosignalingCmd.Flags().StringVarP(&_kinesisvideosignalingSenderClientId, "sender-client-id", "", "", "Sender Client ID")
	_kinesisvideosignalingCmd.Flags().StringVarP(&_kinesisvideosignalingService, "service", "", "", "Service")
	_kinesisvideosignalingCmd.Flags().StringVarP(&_kinesisvideosignalingUsername, "username", "", "", "Username")

	_kinesisvideosignalingCmd.Flags().BoolVarP(&_kinesisvideosignalingGetIceServerConfig, "get-ice-server-config", "", false, "Get Ice Server Config")
	_kinesisvideosignalingCmd.Flags().BoolVarP(&_kinesisvideosignalingSendAlexaOfferToMaster, "send-alexa-offer-to-master", "", false, "Send Alexa Offer To Master")

}
