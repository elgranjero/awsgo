package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/controlcatalog"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// controlcatalogCmd represents the controlcatalog command
var _controlcatalogCmd = &cobra.Command{
	Use:   "controlcatalog",
	Short: "AWS controlcatalog CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := controlcatalog.NewFromConfig(cfg)
		if _controlcatalogGetControl {
			controlcatalog_GetControl(cfg, client)
			return
		}
		if _controlcatalogListCommonControls {
			controlcatalog_ListCommonControls(cfg, client)
			return
		}
		if _controlcatalogListControlMappings {
			controlcatalog_ListControlMappings(cfg, client)
			return
		}
		if _controlcatalogListControls {
			controlcatalog_ListControls(cfg, client)
			return
		}
		if _controlcatalogListDomains {
			controlcatalog_ListDomains(cfg, client)
			return
		}
		if _controlcatalogListObjectives {
			controlcatalog_ListObjectives(cfg, client)
			return
		}

	},
}

var (
	_controlcatalogGetControl          bool
	_controlcatalogListCommonControls  bool
	_controlcatalogListControlMappings bool
	_controlcatalogListControls        bool
	_controlcatalogListDomains         bool
	_controlcatalogListObjectives      bool

	_controlcatalogCommonControlFilter string
	_controlcatalogControlArn          string
	_controlcatalogFilter              string
	_controlcatalogMaxResults          string
	_controlcatalogNextToken           string
	_controlcatalogObjectiveFilter     string
)

// Returns details about a specific control, most notably a list of Amazon Web
// Services Regions where this control is supported. Input a value for the
// ControlArn parameter, in ARN form. GetControl accepts controltower or
// controlcatalog control ARNs as input. Returns a controlcatalog ARN format.
//
// In the API response, controls that have the value GLOBAL in the Scope field do
// not show the DeployableRegions field, because it does not apply. Controls that
// have the value REGIONAL in the Scope field return a value for the
// DeployableRegions field, as shown in the example.
func controlcatalog_GetControl(cfg aws.Config, client *controlcatalog.Client) {
	input := &controlcatalog.GetControlInput{
		// ControlArn: *string, // Required
	}

	if len(_controlcatalogControlArn) > 0 {
		input.ControlArn = aws.String(_controlcatalogControlArn)
	}

	if resp, err := client.GetControl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a paginated list of common controls from the Amazon Web Services
// Control Catalog.
//
// You can apply an optional filter to see common controls that have a specific
// objective. If you don’t provide a filter, the operation returns all common
// controls.
func controlcatalog_ListCommonControls(cfg aws.Config, client *controlcatalog.Client) {
	input := &controlcatalog.ListCommonControlsInput{}

	if len(_controlcatalogCommonControlFilter) > 0 {
		if err := assignInputField(input, "CommonControlFilter", _controlcatalogCommonControlFilter); err != nil {
			log.Errorf("invalid --common-control-filter: %s", err.Error())
			return
		}
	}
	if len(_controlcatalogMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _controlcatalogMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_controlcatalogNextToken) > 0 {
		input.NextToken = aws.String(_controlcatalogNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCommonControls(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*controlcatalog.ListCommonControlsOutput
	p := controlcatalog.NewListCommonControlsPaginator(client, input)
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

// Returns a paginated list of control mappings from the Control Catalog. Control
// mappings show relationships between controls and other entities, such as common
// controls or compliance frameworks.
func controlcatalog_ListControlMappings(cfg aws.Config, client *controlcatalog.Client) {
	input := &controlcatalog.ListControlMappingsInput{}

	if len(_controlcatalogFilter) > 0 {
		if err := assignInputField(input, "Filter", _controlcatalogFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_controlcatalogMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _controlcatalogMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_controlcatalogNextToken) > 0 {
		input.NextToken = aws.String(_controlcatalogNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListControlMappings(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*controlcatalog.ListControlMappingsOutput
	p := controlcatalog.NewListControlMappingsPaginator(client, input)
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

// Returns a paginated list of all available controls in the Control Catalog
// library. Allows you to discover available controls. The list of controls is
// given as structures of type controlSummary. The ARN is returned in the global
// controlcatalog format, as shown in the examples.
func controlcatalog_ListControls(cfg aws.Config, client *controlcatalog.Client) {
	input := &controlcatalog.ListControlsInput{}

	if len(_controlcatalogFilter) > 0 {
		if err := assignInputField(input, "Filter", _controlcatalogFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_controlcatalogMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _controlcatalogMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_controlcatalogNextToken) > 0 {
		input.NextToken = aws.String(_controlcatalogNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListControls(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*controlcatalog.ListControlsOutput
	p := controlcatalog.NewListControlsPaginator(client, input)
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

// Returns a paginated list of domains from the Control Catalog.
func controlcatalog_ListDomains(cfg aws.Config, client *controlcatalog.Client) {
	input := &controlcatalog.ListDomainsInput{}

	if len(_controlcatalogMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _controlcatalogMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_controlcatalogNextToken) > 0 {
		input.NextToken = aws.String(_controlcatalogNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDomains(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*controlcatalog.ListDomainsOutput
	p := controlcatalog.NewListDomainsPaginator(client, input)
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

// Returns a paginated list of objectives from the Control Catalog.
// You can apply an optional filter to see the objectives that belong to a
// specific domain. If you don’t provide a filter, the operation returns all
// objectives.
func controlcatalog_ListObjectives(cfg aws.Config, client *controlcatalog.Client) {
	input := &controlcatalog.ListObjectivesInput{}

	if len(_controlcatalogMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _controlcatalogMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_controlcatalogNextToken) > 0 {
		input.NextToken = aws.String(_controlcatalogNextToken)
	}
	if len(_controlcatalogObjectiveFilter) > 0 {
		if err := assignInputField(input, "ObjectiveFilter", _controlcatalogObjectiveFilter); err != nil {
			log.Errorf("invalid --objective-filter: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListObjectives(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*controlcatalog.ListObjectivesOutput
	p := controlcatalog.NewListObjectivesPaginator(client, input)
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
	_rootCmd.AddCommand(_controlcatalogCmd)
	_controlcatalogCmd.Flags().SortFlags = false

	_controlcatalogCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_controlcatalogCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_controlcatalogCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_controlcatalogCmd.Flags().StringVarP(&_controlcatalogCommonControlFilter, "common-control-filter", "", "", "Common Control Filter")
	_controlcatalogCmd.Flags().StringVarP(&_controlcatalogControlArn, "control-arn", "", "", "Control ARN")
	_controlcatalogCmd.Flags().StringVarP(&_controlcatalogFilter, "filter", "", "", "Filter")
	_controlcatalogCmd.Flags().StringVarP(&_controlcatalogMaxResults, "max-results", "", "", "Max Results")
	_controlcatalogCmd.Flags().StringVarP(&_controlcatalogNextToken, "next-token", "", "", "Next Token")
	_controlcatalogCmd.Flags().StringVarP(&_controlcatalogObjectiveFilter, "objective-filter", "", "", "Objective Filter")

	_controlcatalogCmd.Flags().BoolVarP(&_controlcatalogGetControl, "get-control", "", false, "Get Control")
	_controlcatalogCmd.Flags().BoolVarP(&_controlcatalogListCommonControls, "list-common-controls", "", false, "List Common Controls")
	_controlcatalogCmd.Flags().BoolVarP(&_controlcatalogListControlMappings, "list-control-mappings", "", false, "List Control Mappings")
	_controlcatalogCmd.Flags().BoolVarP(&_controlcatalogListControls, "list-controls", "", false, "List Controls")
	_controlcatalogCmd.Flags().BoolVarP(&_controlcatalogListDomains, "list-domains", "", false, "List Domains")
	_controlcatalogCmd.Flags().BoolVarP(&_controlcatalogListObjectives, "list-objectives", "", false, "List Objectives")

}
