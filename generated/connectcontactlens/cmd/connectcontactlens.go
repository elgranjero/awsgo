package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/connectcontactlens"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// connectcontactlensCmd represents the connectcontactlens command
var _connectcontactlensCmd = &cobra.Command{
	Use:   "connectcontactlens",
	Short: "AWS connectcontactlens CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := connectcontactlens.NewFromConfig(cfg)
		if _connectcontactlensListRealtimeContactAnalysisSegments {
			connectcontactlens_ListRealtimeContactAnalysisSegments(cfg, client)
			return
		}

	},
}

var (
	_connectcontactlensListRealtimeContactAnalysisSegments bool

	_connectcontactlensContactId  string
	_connectcontactlensInstanceId string
	_connectcontactlensMaxResults string
	_connectcontactlensNextToken  string
)

// Provides a list of analysis segments for a real-time analysis session.
func connectcontactlens_ListRealtimeContactAnalysisSegments(cfg aws.Config, client *connectcontactlens.Client) {
	input := &connectcontactlens.ListRealtimeContactAnalysisSegmentsInput{
		// ContactId: *string, // Required
		// InstanceId: *string, // Required
	}

	if len(_connectcontactlensContactId) > 0 {
		input.ContactId = aws.String(_connectcontactlensContactId)
	}
	if len(_connectcontactlensInstanceId) > 0 {
		input.InstanceId = aws.String(_connectcontactlensInstanceId)
	}
	if len(_connectcontactlensMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectcontactlensMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectcontactlensNextToken) > 0 {
		input.NextToken = aws.String(_connectcontactlensNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRealtimeContactAnalysisSegments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connectcontactlens.ListRealtimeContactAnalysisSegmentsOutput
	p := connectcontactlens.NewListRealtimeContactAnalysisSegmentsPaginator(client, input)
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

func init() {
	_rootCmd.AddCommand(_connectcontactlensCmd)
	_connectcontactlensCmd.Flags().SortFlags = false

	_connectcontactlensCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_connectcontactlensCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_connectcontactlensCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_connectcontactlensCmd.Flags().StringVarP(&_connectcontactlensContactId, "contact-id", "", "", "Contact ID")
	_connectcontactlensCmd.Flags().StringVarP(&_connectcontactlensInstanceId, "instance-id", "", "", "Instance ID")
	_connectcontactlensCmd.Flags().StringVarP(&_connectcontactlensMaxResults, "max-results", "", "", "Max Results")
	_connectcontactlensCmd.Flags().StringVarP(&_connectcontactlensNextToken, "next-token", "", "", "Next Token")

	_connectcontactlensCmd.Flags().BoolVarP(&_connectcontactlensListRealtimeContactAnalysisSegments, "list-realtime-contact-analysis-segments", "", false, "List Realtime Contact Analysis Segments")

}
