package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/marketplacereporting"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// marketplacereportingCmd represents the marketplacereporting command
var _marketplacereportingCmd = &cobra.Command{
	Use:   "marketplacereporting",
	Short: "AWS marketplacereporting CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := marketplacereporting.NewFromConfig(cfg)
		if _marketplacereportingGetBuyerDashboard {
			marketplacereporting_GetBuyerDashboard(cfg, client)
			return
		}

	},
}

var (
	_marketplacereportingGetBuyerDashboard bool

	_marketplacereportingDashboardIdentifier string
	_marketplacereportingEmbeddingDomains    []string
)

// Generates an embedding URL for an Amazon QuickSight dashboard for an anonymous
// user.
//
// This API is available only to Amazon Web Services Organization management
// accounts or delegated administrators registered for the procurement insights (
// procurement-insights.marketplace.amazonaws.com ) feature.
//
// The following rules apply to a generated URL:
//
// - It contains a temporary bearer token, valid for 5 minutes after it is
// generated. Once redeemed within that period, it cannot be re-used again.
//
// - It has a session lifetime of one hour. The 5-minute validity period runs
// separately from the session lifetime.
func marketplacereporting_GetBuyerDashboard(cfg aws.Config, client *marketplacereporting.Client) {
	input := &marketplacereporting.GetBuyerDashboardInput{
		// DashboardIdentifier: *string, // Required
		// EmbeddingDomains: []string, // Required
	}

	if len(_marketplacereportingDashboardIdentifier) > 0 {
		input.DashboardIdentifier = aws.String(_marketplacereportingDashboardIdentifier)
	}
	if len(_marketplacereportingEmbeddingDomains) > 0 {
		input.EmbeddingDomains = append([]string(nil), _marketplacereportingEmbeddingDomains...)
	}

	if resp, err := client.GetBuyerDashboard(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_marketplacereportingCmd)
	_marketplacereportingCmd.Flags().SortFlags = false

	_marketplacereportingCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_marketplacereportingCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_marketplacereportingCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_marketplacereportingCmd.Flags().StringVarP(&_marketplacereportingDashboardIdentifier, "dashboard-identifier", "", "", "Dashboard Identifier")
	_marketplacereportingCmd.Flags().StringSliceVarP(&_marketplacereportingEmbeddingDomains, "embedding-domains", "", nil, "Embedding Domains")

	_marketplacereportingCmd.Flags().BoolVarP(&_marketplacereportingGetBuyerDashboard, "get-buyer-dashboard", "", false, "Get Buyer Dashboard")

}
