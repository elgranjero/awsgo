package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/bcmrecommendedactions"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// bcmrecommendedactionsCmd represents the bcmrecommendedactions command
var _bcmrecommendedactionsCmd = &cobra.Command{
	Use:   "bcmrecommendedactions",
	Short: "AWS bcmrecommendedactions CLI",
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
		client := bcmrecommendedactions.NewFromConfig(cfg)
		if _bcmrecommendedactionsListRecommendedActions {
			bcmrecommendedactions_ListRecommendedActions(cfg, client)
			return
		}

	},
}

var (
	_bcmrecommendedactionsListRecommendedActions bool

	_bcmrecommendedactionsFilter     string
	_bcmrecommendedactionsMaxResults string
	_bcmrecommendedactionsNextToken  string
)

// Returns a list of recommended actions that match the filter criteria.
func bcmrecommendedactions_ListRecommendedActions(cfg aws.Config, client *bcmrecommendedactions.Client) {
	input := &bcmrecommendedactions.ListRecommendedActionsInput{}

	if len(_bcmrecommendedactionsFilter) > 0 {
		if err := assignInputField(input, "Filter", _bcmrecommendedactionsFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_bcmrecommendedactionsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bcmrecommendedactionsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bcmrecommendedactionsNextToken) > 0 {
		input.NextToken = aws.String(_bcmrecommendedactionsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRecommendedActions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bcmrecommendedactions.ListRecommendedActionsOutput
	p := bcmrecommendedactions.NewListRecommendedActionsPaginator(client, input)
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
	_rootCmd.AddCommand(_bcmrecommendedactionsCmd)
	_bcmrecommendedactionsCmd.Flags().SortFlags = false

	_bcmrecommendedactionsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_bcmrecommendedactionsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_bcmrecommendedactionsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_bcmrecommendedactionsCmd.Flags().StringVarP(&_bcmrecommendedactionsFilter, "filter", "", "", "Filter")
	_bcmrecommendedactionsCmd.Flags().StringVarP(&_bcmrecommendedactionsMaxResults, "max-results", "", "", "Max Results")
	_bcmrecommendedactionsCmd.Flags().StringVarP(&_bcmrecommendedactionsNextToken, "next-token", "", "", "Next Token")

	_bcmrecommendedactionsCmd.Flags().BoolVarP(&_bcmrecommendedactionsListRecommendedActions, "list-recommended-actions", "", false, "List Recommended Actions")

}
