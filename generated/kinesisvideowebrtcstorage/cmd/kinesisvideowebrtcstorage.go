package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideowebrtcstorage"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// kinesisvideowebrtcstorageCmd represents the kinesisvideowebrtcstorage command
var _kinesisvideowebrtcstorageCmd = &cobra.Command{
	Use:   "kinesisvideowebrtcstorage",
	Short: "AWS kinesisvideowebrtcstorage CLI",
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
		client := kinesisvideowebrtcstorage.NewFromConfig(cfg)
		if _kinesisvideowebrtcstorageJoinStorageSession {
			kinesisvideowebrtcstorage_JoinStorageSession(cfg, client)
			return
		}
		if _kinesisvideowebrtcstorageJoinStorageSessionAsViewer {
			kinesisvideowebrtcstorage_JoinStorageSessionAsViewer(cfg, client)
			return
		}

	},
}

var (
	_kinesisvideowebrtcstorageJoinStorageSession         bool
	_kinesisvideowebrtcstorageJoinStorageSessionAsViewer bool

	_kinesisvideowebrtcstorageChannelArn string
	_kinesisvideowebrtcstorageClientId   string
)

// Before using this API, you must call the GetSignalingChannelEndpoint API to
// request the WEBRTC endpoint. You then specify the endpoint and region in your
// JoinStorageSession API request.
//
// Join the ongoing one way-video and/or multi-way audio WebRTC session as a video
// producing device for an input channel. If there’s no existing session for the
// channel, a new streaming session needs to be created, and the Amazon Resource
// Name (ARN) of the signaling channel must be provided.
//
// Currently for the SINGLE_MASTER type, a video producing device is able to
// ingest both audio and video media into a stream. Only video producing devices
// can join the session and record media.
//
// Both audio and video tracks are currently required for WebRTC ingestion.
//
// Current requirements:
//
// - Video track: H.264
//
// - Audio track: Opus
//
// The resulting ingested video in the Kinesis video stream will have the
// following parameters: H.264 video and AAC audio.
//
// Once a master participant has negotiated a connection through WebRTC, the
// ingested media session will be stored in the Kinesis video stream. Multiple
// viewers are then able to play back real-time media through our Playback APIs.
//
// You can also use existing Kinesis Video Streams features like HLS or DASH
// playback, image generation via [GetImages], and more with ingested WebRTC media.
//
// S3 image delivery and notifications are not currently supported.
//
// Assume that only one video producing device client can be associated with a
// session for the channel. If more than one client joins the session of a specific
// channel as a video producing device, the most recent client request takes
// precedence.
//
// # Additional information
//
// - Idempotent - This API is not idempotent.
//
// - Retry behavior - This is counted as a new API call.
//
// - Concurrent calls - Concurrent calls are allowed. An offer is sent once per
// each call.
//
// [GetImages]: https://docs.aws.amazon.com/kinesisvideostreams/latest/dg/gs-getImages.html
func kinesisvideowebrtcstorage_JoinStorageSession(cfg aws.Config, client *kinesisvideowebrtcstorage.Client) {
	input := &kinesisvideowebrtcstorage.JoinStorageSessionInput{
		// ChannelArn: *string, // Required
	}

	if len(_kinesisvideowebrtcstorageChannelArn) > 0 {
		input.ChannelArn = aws.String(_kinesisvideowebrtcstorageChannelArn)
	}

	if resp, err := client.JoinStorageSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Join the ongoing one way-video and/or multi-way audio WebRTC session as a
// viewer for an input channel. If there’s no existing session for the channel,
// create a new streaming session and provide the Amazon Resource Name (ARN) of the
// signaling channel ( channelArn ) and client id ( clientId ).
//
// Currently for SINGLE_MASTER type, a video producing device is able to ingest
// both audio and video media into a stream, while viewers can only ingest audio.
// Both a video producing device and viewers can join a session first and wait for
// other participants. While participants are having peer to peer conversations
// through WebRTC, the ingested media session will be stored into the Kinesis Video
// Stream. Multiple viewers are able to playback real-time media.
//
// Customers can also use existing Kinesis Video Streams features like HLS or DASH
// playback, Image generation, and more with ingested WebRTC media. If there’s an
// existing session with the same clientId that's found in the join session
// request, the new request takes precedence.
func kinesisvideowebrtcstorage_JoinStorageSessionAsViewer(cfg aws.Config, client *kinesisvideowebrtcstorage.Client) {
	input := &kinesisvideowebrtcstorage.JoinStorageSessionAsViewerInput{
		// ChannelArn: *string, // Required
		// ClientId: *string, // Required
	}

	if len(_kinesisvideowebrtcstorageChannelArn) > 0 {
		input.ChannelArn = aws.String(_kinesisvideowebrtcstorageChannelArn)
	}
	if len(_kinesisvideowebrtcstorageClientId) > 0 {
		input.ClientId = aws.String(_kinesisvideowebrtcstorageClientId)
	}

	if resp, err := client.JoinStorageSessionAsViewer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_kinesisvideowebrtcstorageCmd)
	_kinesisvideowebrtcstorageCmd.Flags().SortFlags = false

	_kinesisvideowebrtcstorageCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_kinesisvideowebrtcstorageCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_kinesisvideowebrtcstorageCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_kinesisvideowebrtcstorageCmd.Flags().StringVarP(&_kinesisvideowebrtcstorageChannelArn, "channel-arn", "", "", "Channel ARN")
	_kinesisvideowebrtcstorageCmd.Flags().StringVarP(&_kinesisvideowebrtcstorageClientId, "client-id", "", "", "Client ID")

	_kinesisvideowebrtcstorageCmd.Flags().BoolVarP(&_kinesisvideowebrtcstorageJoinStorageSession, "join-storage-session", "", false, "Join Storage Session")
	_kinesisvideowebrtcstorageCmd.Flags().BoolVarP(&_kinesisvideowebrtcstorageJoinStorageSessionAsViewer, "join-storage-session-as-viewer", "", false, "Join Storage Session As Viewer")

}
