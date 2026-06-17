package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/marketplaceentitlementservice"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// marketplaceentitlementserviceCmd represents the marketplaceentitlementservice command
var _marketplaceentitlementserviceCmd = &cobra.Command{
	Use:   "marketplaceentitlementservice",
	Short: "AWS marketplaceentitlementservice CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := marketplaceentitlementservice.NewFromConfig(cfg)
		if _marketplaceentitlementserviceGetEntitlements {
			marketplaceentitlementservice_GetEntitlements(cfg, client)
			return
		}

	},
}

var (
	_marketplaceentitlementserviceGetEntitlements bool

	_marketplaceentitlementserviceFilter      string
	_marketplaceentitlementserviceMaxResults  string
	_marketplaceentitlementserviceNextToken   string
	_marketplaceentitlementserviceProductCode string
)

// GetEntitlements retrieves entitlement values for a given product. The results
// can be filtered based on customer identifier, AWS account ID, license ARN, or
// product dimensions.
func marketplaceentitlementservice_GetEntitlements(cfg aws.Config, client *marketplaceentitlementservice.Client) {
	input := &marketplaceentitlementservice.GetEntitlementsInput{
		// ProductCode: *string, // Required
	}

	if len(_marketplaceentitlementserviceProductCode) > 0 {
		input.ProductCode = aws.String(_marketplaceentitlementserviceProductCode)
	}
	if len(_marketplaceentitlementserviceFilter) > 0 {
		if err := assignInputField(input, "Filter", _marketplaceentitlementserviceFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_marketplaceentitlementserviceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _marketplaceentitlementserviceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_marketplaceentitlementserviceNextToken) > 0 {
		input.NextToken = aws.String(_marketplaceentitlementserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetEntitlements(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*marketplaceentitlementservice.GetEntitlementsOutput
	p := marketplaceentitlementservice.NewGetEntitlementsPaginator(client, input)
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
	_rootCmd.AddCommand(_marketplaceentitlementserviceCmd)
	_marketplaceentitlementserviceCmd.Flags().SortFlags = false

	_marketplaceentitlementserviceCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_marketplaceentitlementserviceCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_marketplaceentitlementserviceCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_marketplaceentitlementserviceCmd.Flags().StringVarP(&_marketplaceentitlementserviceFilter, "filter", "", "", "Filter")
	_marketplaceentitlementserviceCmd.Flags().StringVarP(&_marketplaceentitlementserviceMaxResults, "max-results", "", "", "Max Results")
	_marketplaceentitlementserviceCmd.Flags().StringVarP(&_marketplaceentitlementserviceNextToken, "next-token", "", "", "Next Token")
	_marketplaceentitlementserviceCmd.Flags().StringVarP(&_marketplaceentitlementserviceProductCode, "product-code", "", "", "Product Code")

	_marketplaceentitlementserviceCmd.Flags().BoolVarP(&_marketplaceentitlementserviceGetEntitlements, "get-entitlements", "", false, "Get Entitlements")

}
