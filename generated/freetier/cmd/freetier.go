package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/freetier"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// freetierCmd represents the freetier command
var _freetierCmd = &cobra.Command{
	Use:   "freetier",
	Short: "AWS freetier CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := freetier.NewFromConfig(cfg)
		if _freetierGetAccountActivity {
			freetier_GetAccountActivity(cfg, client)
			return
		}
		if _freetierGetAccountPlanState {
			freetier_GetAccountPlanState(cfg, client)
			return
		}
		if _freetierGetFreeTierUsage {
			freetier_GetFreeTierUsage(cfg, client)
			return
		}
		if _freetierListAccountActivities {
			freetier_ListAccountActivities(cfg, client)
			return
		}
		if _freetierUpgradeAccountPlan {
			freetier_UpgradeAccountPlan(cfg, client)
			return
		}

	},
}

var (
	_freetierGetAccountActivity    bool
	_freetierGetAccountPlanState   bool
	_freetierGetFreeTierUsage      bool
	_freetierListAccountActivities bool
	_freetierUpgradeAccountPlan    bool

	_freetierAccountPlanType        string
	_freetierActivityId             string
	_freetierFilter                 string
	_freetierFilterActivityStatuses string
	_freetierLanguageCode           string
	_freetierMaxResults             string
	_freetierNextToken              string
)

// Returns a specific activity record that is available to the customer.
func freetier_GetAccountActivity(cfg aws.Config, client *freetier.Client) {
	input := &freetier.GetAccountActivityInput{
		// ActivityId: *string, // Required
	}

	if len(_freetierActivityId) > 0 {
		input.ActivityId = aws.String(_freetierActivityId)
	}
	if len(_freetierLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _freetierLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetAccountActivity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This returns all of the information related to the state of the account plan
// related to Free Tier.
func freetier_GetAccountPlanState(cfg aws.Config, client *freetier.Client) {
	input := &freetier.GetAccountPlanStateInput{}

	if resp, err := client.GetAccountPlanState(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of all Free Tier usage objects that match your filters.
func freetier_GetFreeTierUsage(cfg aws.Config, client *freetier.Client) {
	input := &freetier.GetFreeTierUsageInput{}

	if len(_freetierFilter) > 0 {
		if err := assignInputField(input, "Filter", _freetierFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_freetierMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _freetierMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_freetierNextToken) > 0 {
		input.NextToken = aws.String(_freetierNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetFreeTierUsage(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*freetier.GetFreeTierUsageOutput
	p := freetier.NewGetFreeTierUsagePaginator(client, input)
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

// Returns a list of activities that are available. This operation supports
// pagination and filtering by status.
func freetier_ListAccountActivities(cfg aws.Config, client *freetier.Client) {
	input := &freetier.ListAccountActivitiesInput{}

	if len(_freetierFilterActivityStatuses) > 0 {
		if err := assignInputField(input, "FilterActivityStatuses", _freetierFilterActivityStatuses); err != nil {
			log.Errorf("invalid --filter-activity-statuses: %s", err.Error())
			return
		}
	}
	if len(_freetierLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _freetierLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_freetierMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _freetierMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_freetierNextToken) > 0 {
		input.NextToken = aws.String(_freetierNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAccountActivities(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*freetier.ListAccountActivitiesOutput
	p := freetier.NewListAccountActivitiesPaginator(client, input)
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

// The account plan type for the Amazon Web Services account.
func freetier_UpgradeAccountPlan(cfg aws.Config, client *freetier.Client) {
	input := &freetier.UpgradeAccountPlanInput{
		// AccountPlanType: types.AccountPlanType, // Required
	}

	if len(_freetierAccountPlanType) > 0 {
		if err := assignInputField(input, "AccountPlanType", _freetierAccountPlanType); err != nil {
			log.Errorf("invalid --account-plan-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpgradeAccountPlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_freetierCmd)
	_freetierCmd.Flags().SortFlags = false

	_freetierCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_freetierCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_freetierCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_freetierCmd.Flags().StringVarP(&_freetierAccountPlanType, "account-plan-type", "", "", "Account Plan Type")
	_freetierCmd.Flags().StringVarP(&_freetierActivityId, "activity-id", "", "", "Activity ID")
	_freetierCmd.Flags().StringVarP(&_freetierFilter, "filter", "", "", "Filter")
	_freetierCmd.Flags().StringVarP(&_freetierFilterActivityStatuses, "filter-activity-statuses", "", "", "Filter Activity Statuses")
	_freetierCmd.Flags().StringVarP(&_freetierLanguageCode, "language-code", "", "", "Language Code")
	_freetierCmd.Flags().StringVarP(&_freetierMaxResults, "max-results", "", "", "Max Results")
	_freetierCmd.Flags().StringVarP(&_freetierNextToken, "next-token", "", "", "Next Token")

	_freetierCmd.Flags().BoolVarP(&_freetierGetAccountActivity, "get-account-activity", "", false, "Get Account Activity")
	_freetierCmd.Flags().BoolVarP(&_freetierGetAccountPlanState, "get-account-plan-state", "", false, "Get Account Plan State")
	_freetierCmd.Flags().BoolVarP(&_freetierGetFreeTierUsage, "get-free-tier-usage", "", false, "Get Free Tier Usage")
	_freetierCmd.Flags().BoolVarP(&_freetierListAccountActivities, "list-account-activities", "", false, "List Account Activities")
	_freetierCmd.Flags().BoolVarP(&_freetierUpgradeAccountPlan, "upgrade-account-plan", "", false, "Upgrade Account Plan")

}
