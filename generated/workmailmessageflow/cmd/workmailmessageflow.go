package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/workmailmessageflow"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// workmailmessageflowCmd represents the workmailmessageflow command
var _workmailmessageflowCmd = &cobra.Command{
	Use:   "workmailmessageflow",
	Short: "AWS workmailmessageflow CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := workmailmessageflow.NewFromConfig(cfg)
		if _workmailmessageflowGetRawMessageContent {
			workmailmessageflow_GetRawMessageContent(cfg, client)
			return
		}
		if _workmailmessageflowPutRawMessageContent {
			workmailmessageflow_PutRawMessageContent(cfg, client)
			return
		}

	},
}

var (
	_workmailmessageflowGetRawMessageContent bool
	_workmailmessageflowPutRawMessageContent bool

	_workmailmessageflowContent   string
	_workmailmessageflowMessageId string
)

// Retrieves the raw content of an in-transit email message, in MIME format.
func workmailmessageflow_GetRawMessageContent(cfg aws.Config, client *workmailmessageflow.Client) {
	input := &workmailmessageflow.GetRawMessageContentInput{
		// MessageId: *string, // Required
	}

	if len(_workmailmessageflowMessageId) > 0 {
		input.MessageId = aws.String(_workmailmessageflowMessageId)
	}

	if resp, err := client.GetRawMessageContent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the raw content of an in-transit email message, in MIME format.
// This example describes how to update in-transit email message. For more
// information and examples for using this API, see [Updating message content with AWS Lambda].
//
// Updates to an in-transit message only appear when you call PutRawMessageContent
// from an AWS Lambda function configured with a synchronous [Run Lambda]rule. If you call
// PutRawMessageContent on a delivered or sent message, the message remains
// unchanged, even though [GetRawMessageContent]returns an updated message.
//
// [GetRawMessageContent]: https://docs.aws.amazon.com/workmail/latest/APIReference/API_messageflow_GetRawMessageContent.html
// [Run Lambda]: https://docs.aws.amazon.com/workmail/latest/adminguide/lambda.html#synchronous-rules
// [Updating message content with AWS Lambda]: https://docs.aws.amazon.com/workmail/latest/adminguide/update-with-lambda.html
func workmailmessageflow_PutRawMessageContent(cfg aws.Config, client *workmailmessageflow.Client) {
	input := &workmailmessageflow.PutRawMessageContentInput{
		// Content: *types.RawMessageContent, // Required
		// MessageId: *string, // Required
	}

	if len(_workmailmessageflowContent) > 0 {
		if err := assignInputField(input, "Content", _workmailmessageflowContent); err != nil {
			log.Errorf("invalid --content: %s", err.Error())
			return
		}
	}
	if len(_workmailmessageflowMessageId) > 0 {
		input.MessageId = aws.String(_workmailmessageflowMessageId)
	}

	if resp, err := client.PutRawMessageContent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_workmailmessageflowCmd)
	_workmailmessageflowCmd.Flags().SortFlags = false

	_workmailmessageflowCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_workmailmessageflowCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_workmailmessageflowCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_workmailmessageflowCmd.Flags().StringVarP(&_workmailmessageflowContent, "content", "", "", "Content")
	_workmailmessageflowCmd.Flags().StringVarP(&_workmailmessageflowMessageId, "message-id", "", "", "Message ID")

	_workmailmessageflowCmd.Flags().BoolVarP(&_workmailmessageflowGetRawMessageContent, "get-raw-message-content", "", false, "Get Raw Message Content")
	_workmailmessageflowCmd.Flags().BoolVarP(&_workmailmessageflowPutRawMessageContent, "put-raw-message-content", "", false, "Put Raw Message Content")

}
