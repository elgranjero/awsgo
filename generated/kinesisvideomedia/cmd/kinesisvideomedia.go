package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideomedia"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// kinesisvideomediaCmd represents the kinesisvideomedia command
var _kinesisvideomediaCmd = &cobra.Command{
	Use:   "kinesisvideomedia",
	Short: "AWS kinesisvideomedia CLI",
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
		client := kinesisvideomedia.NewFromConfig(cfg)
		if _kinesisvideomediaGetMedia {
			kinesisvideomedia_GetMedia(cfg, client)
			return
		}

	},
}

var (
	_kinesisvideomediaGetMedia bool

	_kinesisvideomediaStartSelector string
	_kinesisvideomediaStreamARN     string
	_kinesisvideomediaStreamName    string
)

// Use this API to retrieve media content from a Kinesis video stream. In the
// request, you identify the stream name or stream Amazon Resource Name (ARN), and
// the starting chunk. Kinesis Video Streams then returns a stream of chunks in
// order by fragment number.
//
// You must first call the GetDataEndpoint API to get an endpoint. Then send the
// GetMedia requests to this endpoint using the [--endpoint-url parameter].
//
// When you put media data (fragments) on a stream, Kinesis Video Streams stores
// each incoming fragment and related metadata in what is called a "chunk." For
// more information, see [PutMedia]. The GetMedia API returns a stream of these chunks
// starting from the chunk that you specify in the request.
//
// The following limits apply when using the GetMedia API:
//
// - A client can call GetMedia up to five times per second per stream.
//
// - Kinesis Video Streams sends media data at a rate of up to 25 megabytes per
// second (or 200 megabits per second) during a GetMedia session.
//
// If an error is thrown after invoking a Kinesis Video Streams media API, in
// addition to the HTTP status code and the response body, it includes the
// following pieces of information:
//
// - x-amz-ErrorType HTTP header – contains a more specific error type in
// addition to what the HTTP status code provides.
//
// - x-amz-RequestId HTTP header – if you want to report an issue to AWS, the
// support team can better diagnose the problem if given the Request Id.
//
// Both the HTTP status code and the ErrorType header can be utilized to make
// programmatic decisions about whether errors are retry-able and under what
// conditions, as well as provide information on what actions the client programmer
// might need to take in order to successfully try again.
//
// For more information, see the Errors section at the bottom of this topic, as
// well as [Common Errors].
//
// [PutMedia]: https://docs.aws.amazon.com/kinesisvideostreams/latest/dg/API_dataplane_PutMedia.html
// [--endpoint-url parameter]: https://docs.aws.amazon.com/cli/latest/reference/
// [Common Errors]: https://docs.aws.amazon.com/kinesisvideostreams/latest/dg/CommonErrors.html
func kinesisvideomedia_GetMedia(cfg aws.Config, client *kinesisvideomedia.Client) {
	input := &kinesisvideomedia.GetMediaInput{
		// StartSelector: *types.StartSelector, // Required
	}

	if len(_kinesisvideomediaStartSelector) > 0 {
		if err := assignInputField(input, "StartSelector", _kinesisvideomediaStartSelector); err != nil {
			log.Errorf("invalid --start-selector: %s", err.Error())
			return
		}
	}
	if len(_kinesisvideomediaStreamARN) > 0 {
		input.StreamARN = aws.String(_kinesisvideomediaStreamARN)
	}
	if len(_kinesisvideomediaStreamName) > 0 {
		input.StreamName = aws.String(_kinesisvideomediaStreamName)
	}

	if resp, err := client.GetMedia(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_kinesisvideomediaCmd)
	_kinesisvideomediaCmd.Flags().SortFlags = false

	_kinesisvideomediaCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_kinesisvideomediaCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_kinesisvideomediaCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_kinesisvideomediaCmd.Flags().StringVarP(&_kinesisvideomediaStartSelector, "start-selector", "", "", "Start Selector")
	_kinesisvideomediaCmd.Flags().StringVarP(&_kinesisvideomediaStreamARN, "stream-arn", "", "", "Stream ARN")
	_kinesisvideomediaCmd.Flags().StringVarP(&_kinesisvideomediaStreamName, "stream-name", "", "", "Stream Name")

	_kinesisvideomediaCmd.Flags().BoolVarP(&_kinesisvideomediaGetMedia, "get-media", "", false, "Get Media")

}
