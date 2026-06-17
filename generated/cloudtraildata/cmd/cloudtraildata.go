package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudtraildata"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// cloudtraildataCmd represents the cloudtraildata command
var _cloudtraildataCmd = &cobra.Command{
	Use:   "cloudtraildata",
	Short: "AWS cloudtraildata CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := cloudtraildata.NewFromConfig(cfg)
		if _cloudtraildataPutAuditEvents {
			cloudtraildata_PutAuditEvents(cfg, client)
			return
		}

	},
}

var (
	_cloudtraildataPutAuditEvents bool

	_cloudtraildataAuditEvents string
	_cloudtraildataChannelArn  string
	_cloudtraildataExternalId  string
)

// Ingests your application events into CloudTrail Lake. A required parameter,
// auditEvents , accepts the JSON records (also called payload) of events that you
// want CloudTrail to ingest. You can add up to 100 of these events (or up to 1 MB)
// per PutAuditEvents request.
func cloudtraildata_PutAuditEvents(cfg aws.Config, client *cloudtraildata.Client) {
	input := &cloudtraildata.PutAuditEventsInput{
		// AuditEvents: []types.AuditEvent, // Required
		// ChannelArn: *string, // Required
	}

	if len(_cloudtraildataAuditEvents) > 0 {
		if err := assignInputField(input, "AuditEvents", _cloudtraildataAuditEvents); err != nil {
			log.Errorf("invalid --audit-events: %s", err.Error())
			return
		}
	}
	if len(_cloudtraildataChannelArn) > 0 {
		input.ChannelArn = aws.String(_cloudtraildataChannelArn)
	}
	if len(_cloudtraildataExternalId) > 0 {
		input.ExternalId = aws.String(_cloudtraildataExternalId)
	}

	if resp, err := client.PutAuditEvents(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_cloudtraildataCmd)
	_cloudtraildataCmd.Flags().SortFlags = false

	_cloudtraildataCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_cloudtraildataCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_cloudtraildataCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_cloudtraildataCmd.Flags().StringVarP(&_cloudtraildataAuditEvents, "audit-events", "", "", "Audit Events")
	_cloudtraildataCmd.Flags().StringVarP(&_cloudtraildataChannelArn, "channel-arn", "", "", "Channel ARN")
	_cloudtraildataCmd.Flags().StringVarP(&_cloudtraildataExternalId, "external-id", "", "", "External ID")

	_cloudtraildataCmd.Flags().BoolVarP(&_cloudtraildataPutAuditEvents, "put-audit-events", "", false, "Put Audit Events")

}
