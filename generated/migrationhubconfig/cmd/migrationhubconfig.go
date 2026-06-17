package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/migrationhubconfig"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// migrationhubconfigCmd represents the migrationhubconfig command
var _migrationhubconfigCmd = &cobra.Command{
	Use:   "migrationhubconfig",
	Short: "AWS migrationhubconfig CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := migrationhubconfig.NewFromConfig(cfg)
		if _migrationhubconfigCreateHomeRegionControl {
			migrationhubconfig_CreateHomeRegionControl(cfg, client)
			return
		}
		if _migrationhubconfigDeleteHomeRegionControl {
			migrationhubconfig_DeleteHomeRegionControl(cfg, client)
			return
		}
		if _migrationhubconfigDescribeHomeRegionControls {
			migrationhubconfig_DescribeHomeRegionControls(cfg, client)
			return
		}
		if _migrationhubconfigGetHomeRegion {
			migrationhubconfig_GetHomeRegion(cfg, client)
			return
		}

	},
}

var (
	_migrationhubconfigCreateHomeRegionControl    bool
	_migrationhubconfigDeleteHomeRegionControl    bool
	_migrationhubconfigDescribeHomeRegionControls bool
	_migrationhubconfigGetHomeRegion              bool

	_migrationhubconfigControlId  string
	_migrationhubconfigDryRun     string
	_migrationhubconfigHomeRegion string
	_migrationhubconfigMaxResults string
	_migrationhubconfigNextToken  string
	_migrationhubconfigTarget     string
)

// This API sets up the home region for the calling account only.
func migrationhubconfig_CreateHomeRegionControl(cfg aws.Config, client *migrationhubconfig.Client) {
	input := &migrationhubconfig.CreateHomeRegionControlInput{
		// HomeRegion: *string, // Required
		// Target: *types.Target, // Required
	}

	if len(_migrationhubconfigHomeRegion) > 0 {
		input.HomeRegion = aws.String(_migrationhubconfigHomeRegion)
	}
	if len(_migrationhubconfigTarget) > 0 {
		if err := assignInputField(input, "Target", _migrationhubconfigTarget); err != nil {
			log.Errorf("invalid --target: %s", err.Error())
			return
		}
	}
	if len(_migrationhubconfigDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _migrationhubconfigDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateHomeRegionControl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation deletes the home region configuration for the calling account.
// The operation does not delete discovery or migration tracking data in the home
// region.
func migrationhubconfig_DeleteHomeRegionControl(cfg aws.Config, client *migrationhubconfig.Client) {
	input := &migrationhubconfig.DeleteHomeRegionControlInput{
		// ControlId: *string, // Required
	}

	if len(_migrationhubconfigControlId) > 0 {
		input.ControlId = aws.String(_migrationhubconfigControlId)
	}

	if resp, err := client.DeleteHomeRegionControl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API permits filtering on the ControlId and HomeRegion fields.
func migrationhubconfig_DescribeHomeRegionControls(cfg aws.Config, client *migrationhubconfig.Client) {
	input := &migrationhubconfig.DescribeHomeRegionControlsInput{}

	if len(_migrationhubconfigControlId) > 0 {
		input.ControlId = aws.String(_migrationhubconfigControlId)
	}
	if len(_migrationhubconfigHomeRegion) > 0 {
		input.HomeRegion = aws.String(_migrationhubconfigHomeRegion)
	}
	if len(_migrationhubconfigMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _migrationhubconfigMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_migrationhubconfigNextToken) > 0 {
		input.NextToken = aws.String(_migrationhubconfigNextToken)
	}
	if len(_migrationhubconfigTarget) > 0 {
		if err := assignInputField(input, "Target", _migrationhubconfigTarget); err != nil {
			log.Errorf("invalid --target: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeHomeRegionControls(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*migrationhubconfig.DescribeHomeRegionControlsOutput
	p := migrationhubconfig.NewDescribeHomeRegionControlsPaginator(client, input)
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

// Returns the calling account’s home region, if configured. This API is used by
// other AWS services to determine the regional endpoint for calling AWS
// Application Discovery Service and Migration Hub. You must call GetHomeRegion at
// least once before you call any other AWS Application Discovery Service and AWS
// Migration Hub APIs, to obtain the account's Migration Hub home region.
func migrationhubconfig_GetHomeRegion(cfg aws.Config, client *migrationhubconfig.Client) {
	input := &migrationhubconfig.GetHomeRegionInput{}

	if resp, err := client.GetHomeRegion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_migrationhubconfigCmd)
	_migrationhubconfigCmd.Flags().SortFlags = false

	_migrationhubconfigCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_migrationhubconfigCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_migrationhubconfigCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_migrationhubconfigCmd.Flags().StringVarP(&_migrationhubconfigControlId, "control-id", "", "", "Control ID")
	_migrationhubconfigCmd.Flags().StringVarP(&_migrationhubconfigDryRun, "dry-run", "", "", "Dry Run")
	_migrationhubconfigCmd.Flags().StringVarP(&_migrationhubconfigHomeRegion, "home-region", "", "", "Home Region")
	_migrationhubconfigCmd.Flags().StringVarP(&_migrationhubconfigMaxResults, "max-results", "", "", "Max Results")
	_migrationhubconfigCmd.Flags().StringVarP(&_migrationhubconfigNextToken, "next-token", "", "", "Next Token")
	_migrationhubconfigCmd.Flags().StringVarP(&_migrationhubconfigTarget, "target", "", "", "Target")

	_migrationhubconfigCmd.Flags().BoolVarP(&_migrationhubconfigCreateHomeRegionControl, "create-home-region-control", "", false, "Create Home Region Control")
	_migrationhubconfigCmd.Flags().BoolVarP(&_migrationhubconfigDeleteHomeRegionControl, "delete-home-region-control", "", false, "Delete Home Region Control")
	_migrationhubconfigCmd.Flags().BoolVarP(&_migrationhubconfigDescribeHomeRegionControls, "describe-home-region-controls", "", false, "Describe Home Region Controls")
	_migrationhubconfigCmd.Flags().BoolVarP(&_migrationhubconfigGetHomeRegion, "get-home-region", "", false, "Get Home Region")

}
