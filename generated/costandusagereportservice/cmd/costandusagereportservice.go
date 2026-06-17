package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/costandusagereportservice"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// costandusagereportserviceCmd represents the costandusagereportservice command
var _costandusagereportserviceCmd = &cobra.Command{
	Use:   "costandusagereportservice",
	Short: "AWS costandusagereportservice CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := costandusagereportservice.NewFromConfig(cfg)
		if _costandusagereportserviceDeleteReportDefinition {
			costandusagereportservice_DeleteReportDefinition(cfg, client)
			return
		}
		if _costandusagereportserviceDescribeReportDefinitions {
			costandusagereportservice_DescribeReportDefinitions(cfg, client)
			return
		}
		if _costandusagereportserviceListTagsForResource {
			costandusagereportservice_ListTagsForResource(cfg, client)
			return
		}
		if _costandusagereportserviceModifyReportDefinition {
			costandusagereportservice_ModifyReportDefinition(cfg, client)
			return
		}
		if _costandusagereportservicePutReportDefinition {
			costandusagereportservice_PutReportDefinition(cfg, client)
			return
		}
		if _costandusagereportserviceTagResource {
			costandusagereportservice_TagResource(cfg, client)
			return
		}
		if _costandusagereportserviceUntagResource {
			costandusagereportservice_UntagResource(cfg, client)
			return
		}

	},
}

var (
	_costandusagereportserviceDeleteReportDefinition    bool
	_costandusagereportserviceDescribeReportDefinitions bool
	_costandusagereportserviceListTagsForResource       bool
	_costandusagereportserviceModifyReportDefinition    bool
	_costandusagereportservicePutReportDefinition       bool
	_costandusagereportserviceTagResource               bool
	_costandusagereportserviceUntagResource             bool

	_costandusagereportserviceMaxResults       string
	_costandusagereportserviceNextToken        string
	_costandusagereportserviceReportDefinition string
	_costandusagereportserviceReportName       string
	_costandusagereportserviceTagKeys          []string
	_costandusagereportserviceTags             string
)

// Deletes the specified report. Any tags associated with the report are also
// deleted.
func costandusagereportservice_DeleteReportDefinition(cfg aws.Config, client *costandusagereportservice.Client) {
	input := &costandusagereportservice.DeleteReportDefinitionInput{
		// ReportName: *string, // Required
	}

	if len(_costandusagereportserviceReportName) > 0 {
		input.ReportName = aws.String(_costandusagereportserviceReportName)
	}

	if resp, err := client.DeleteReportDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the Amazon Web Services Cost and Usage Report available to this account.
func costandusagereportservice_DescribeReportDefinitions(cfg aws.Config, client *costandusagereportservice.Client) {
	input := &costandusagereportservice.DescribeReportDefinitionsInput{}

	if len(_costandusagereportserviceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _costandusagereportserviceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_costandusagereportserviceNextToken) > 0 {
		input.NextToken = aws.String(_costandusagereportserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeReportDefinitions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*costandusagereportservice.DescribeReportDefinitionsOutput
	p := costandusagereportservice.NewDescribeReportDefinitionsPaginator(client, input)
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

// Lists the tags associated with the specified report definition.
func costandusagereportservice_ListTagsForResource(cfg aws.Config, client *costandusagereportservice.Client) {
	input := &costandusagereportservice.ListTagsForResourceInput{
		// ReportName: *string, // Required
	}

	if len(_costandusagereportserviceReportName) > 0 {
		input.ReportName = aws.String(_costandusagereportserviceReportName)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Allows you to programmatically update your report preferences.
func costandusagereportservice_ModifyReportDefinition(cfg aws.Config, client *costandusagereportservice.Client) {
	input := &costandusagereportservice.ModifyReportDefinitionInput{
		// ReportDefinition: *types.ReportDefinition, // Required
		// ReportName: *string, // Required
	}

	if len(_costandusagereportserviceReportDefinition) > 0 {
		if err := assignInputField(input, "ReportDefinition", _costandusagereportserviceReportDefinition); err != nil {
			log.Errorf("invalid --report-definition: %s", err.Error())
			return
		}
	}
	if len(_costandusagereportserviceReportName) > 0 {
		input.ReportName = aws.String(_costandusagereportserviceReportName)
	}

	if resp, err := client.ModifyReportDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new report using the description that you provide.
func costandusagereportservice_PutReportDefinition(cfg aws.Config, client *costandusagereportservice.Client) {
	input := &costandusagereportservice.PutReportDefinitionInput{
		// ReportDefinition: *types.ReportDefinition, // Required
	}

	if len(_costandusagereportserviceReportDefinition) > 0 {
		if err := assignInputField(input, "ReportDefinition", _costandusagereportserviceReportDefinition); err != nil {
			log.Errorf("invalid --report-definition: %s", err.Error())
			return
		}
	}
	if len(_costandusagereportserviceTags) > 0 {
		if err := assignInputField(input, "Tags", _costandusagereportserviceTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutReportDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a set of tags with a report definition.
func costandusagereportservice_TagResource(cfg aws.Config, client *costandusagereportservice.Client) {
	input := &costandusagereportservice.TagResourceInput{
		// ReportName: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_costandusagereportserviceReportName) > 0 {
		input.ReportName = aws.String(_costandusagereportserviceReportName)
	}
	if len(_costandusagereportserviceTags) > 0 {
		if err := assignInputField(input, "Tags", _costandusagereportserviceTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.TagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a set of tags from a report definition.
func costandusagereportservice_UntagResource(cfg aws.Config, client *costandusagereportservice.Client) {
	input := &costandusagereportservice.UntagResourceInput{
		// ReportName: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_costandusagereportserviceReportName) > 0 {
		input.ReportName = aws.String(_costandusagereportserviceReportName)
	}
	if len(_costandusagereportserviceTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _costandusagereportserviceTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_costandusagereportserviceCmd)
	_costandusagereportserviceCmd.Flags().SortFlags = false

	_costandusagereportserviceCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_costandusagereportserviceCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_costandusagereportserviceCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_costandusagereportserviceCmd.Flags().StringVarP(&_costandusagereportserviceMaxResults, "max-results", "", "", "Max Results")
	_costandusagereportserviceCmd.Flags().StringVarP(&_costandusagereportserviceNextToken, "next-token", "", "", "Next Token")
	_costandusagereportserviceCmd.Flags().StringVarP(&_costandusagereportserviceReportDefinition, "report-definition", "", "", "Report Definition")
	_costandusagereportserviceCmd.Flags().StringVarP(&_costandusagereportserviceReportName, "report-name", "", "", "Report Name")
	_costandusagereportserviceCmd.Flags().StringSliceVarP(&_costandusagereportserviceTagKeys, "tag-keys", "", nil, "Tag Keys")
	_costandusagereportserviceCmd.Flags().StringVarP(&_costandusagereportserviceTags, "tags", "", "", "Tags")

	_costandusagereportserviceCmd.Flags().BoolVarP(&_costandusagereportserviceDeleteReportDefinition, "delete-report-definition", "", false, "Delete Report Definition")
	_costandusagereportserviceCmd.Flags().BoolVarP(&_costandusagereportserviceDescribeReportDefinitions, "describe-report-definitions", "", false, "Describe Report Definitions")
	_costandusagereportserviceCmd.Flags().BoolVarP(&_costandusagereportserviceListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_costandusagereportserviceCmd.Flags().BoolVarP(&_costandusagereportserviceModifyReportDefinition, "modify-report-definition", "", false, "Modify Report Definition")
	_costandusagereportserviceCmd.Flags().BoolVarP(&_costandusagereportservicePutReportDefinition, "put-report-definition", "", false, "Put Report Definition")
	_costandusagereportserviceCmd.Flags().BoolVarP(&_costandusagereportserviceTagResource, "tag-resource", "", false, "Tag Resource")
	_costandusagereportserviceCmd.Flags().BoolVarP(&_costandusagereportserviceUntagResource, "untag-resource", "", false, "Untag Resource")

}
