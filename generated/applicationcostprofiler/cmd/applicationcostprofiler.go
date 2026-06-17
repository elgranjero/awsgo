package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/applicationcostprofiler"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// applicationcostprofilerCmd represents the applicationcostprofiler command
var _applicationcostprofilerCmd = &cobra.Command{
	Use:   "applicationcostprofiler",
	Short: "AWS applicationcostprofiler CLI",
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
		client := applicationcostprofiler.NewFromConfig(cfg)
		if _applicationcostprofilerDeleteReportDefinition {
			applicationcostprofiler_DeleteReportDefinition(cfg, client)
			return
		}
		if _applicationcostprofilerGetReportDefinition {
			applicationcostprofiler_GetReportDefinition(cfg, client)
			return
		}
		if _applicationcostprofilerImportApplicationUsage {
			applicationcostprofiler_ImportApplicationUsage(cfg, client)
			return
		}
		if _applicationcostprofilerListReportDefinitions {
			applicationcostprofiler_ListReportDefinitions(cfg, client)
			return
		}
		if _applicationcostprofilerPutReportDefinition {
			applicationcostprofiler_PutReportDefinition(cfg, client)
			return
		}
		if _applicationcostprofilerUpdateReportDefinition {
			applicationcostprofiler_UpdateReportDefinition(cfg, client)
			return
		}

	},
}

var (
	_applicationcostprofilerDeleteReportDefinition bool
	_applicationcostprofilerGetReportDefinition    bool
	_applicationcostprofilerImportApplicationUsage bool
	_applicationcostprofilerListReportDefinitions  bool
	_applicationcostprofilerPutReportDefinition    bool
	_applicationcostprofilerUpdateReportDefinition bool

	_applicationcostprofilerDestinationS3Location string
	_applicationcostprofilerFormat                string
	_applicationcostprofilerMaxResults            string
	_applicationcostprofilerNextToken             string
	_applicationcostprofilerReportDescription     string
	_applicationcostprofilerReportFrequency       string
	_applicationcostprofilerReportId              string
	_applicationcostprofilerSourceS3Location      string
)

// Deletes the specified report definition in AWS Application Cost Profiler. This
// stops the report from being generated.
func applicationcostprofiler_DeleteReportDefinition(cfg aws.Config, client *applicationcostprofiler.Client) {
	input := &applicationcostprofiler.DeleteReportDefinitionInput{
		// ReportId: *string, // Required
	}

	if len(_applicationcostprofilerReportId) > 0 {
		input.ReportId = aws.String(_applicationcostprofilerReportId)
	}

	if resp, err := client.DeleteReportDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the definition of a report already configured in AWS Application Cost
// Profiler.
func applicationcostprofiler_GetReportDefinition(cfg aws.Config, client *applicationcostprofiler.Client) {
	input := &applicationcostprofiler.GetReportDefinitionInput{
		// ReportId: *string, // Required
	}

	if len(_applicationcostprofilerReportId) > 0 {
		input.ReportId = aws.String(_applicationcostprofilerReportId)
	}

	if resp, err := client.GetReportDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Ingests application usage data from Amazon Simple Storage Service (Amazon S3).
// The data must already exist in the S3 location. As part of the action, AWS
// Application Cost Profiler copies the object from your S3 bucket to an S3 bucket
// owned by Amazon for processing asynchronously.
func applicationcostprofiler_ImportApplicationUsage(cfg aws.Config, client *applicationcostprofiler.Client) {
	input := &applicationcostprofiler.ImportApplicationUsageInput{
		// SourceS3Location: *types.SourceS3Location, // Required
	}

	if len(_applicationcostprofilerSourceS3Location) > 0 {
		if err := assignInputField(input, "SourceS3Location", _applicationcostprofilerSourceS3Location); err != nil {
			log.Errorf("invalid --source-s3-location: %s", err.Error())
			return
		}
	}

	if resp, err := client.ImportApplicationUsage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of all reports and their configurations for your AWS account.
// The maximum number of reports is one.
func applicationcostprofiler_ListReportDefinitions(cfg aws.Config, client *applicationcostprofiler.Client) {
	input := &applicationcostprofiler.ListReportDefinitionsInput{}

	if len(_applicationcostprofilerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _applicationcostprofilerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_applicationcostprofilerNextToken) > 0 {
		input.NextToken = aws.String(_applicationcostprofilerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListReportDefinitions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*applicationcostprofiler.ListReportDefinitionsOutput
	p := applicationcostprofiler.NewListReportDefinitionsPaginator(client, input)
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

// Creates the report definition for a report in Application Cost Profiler.
func applicationcostprofiler_PutReportDefinition(cfg aws.Config, client *applicationcostprofiler.Client) {
	input := &applicationcostprofiler.PutReportDefinitionInput{
		// DestinationS3Location: *types.S3Location, // Required
		// Format: types.Format, // Required
		// ReportDescription: *string, // Required
		// ReportFrequency: types.ReportFrequency, // Required
		// ReportId: *string, // Required
	}

	if len(_applicationcostprofilerDestinationS3Location) > 0 {
		if err := assignInputField(input, "DestinationS3Location", _applicationcostprofilerDestinationS3Location); err != nil {
			log.Errorf("invalid --destination-s3-location: %s", err.Error())
			return
		}
	}
	if len(_applicationcostprofilerFormat) > 0 {
		if err := assignInputField(input, "Format", _applicationcostprofilerFormat); err != nil {
			log.Errorf("invalid --format: %s", err.Error())
			return
		}
	}
	if len(_applicationcostprofilerReportDescription) > 0 {
		input.ReportDescription = aws.String(_applicationcostprofilerReportDescription)
	}
	if len(_applicationcostprofilerReportFrequency) > 0 {
		if err := assignInputField(input, "ReportFrequency", _applicationcostprofilerReportFrequency); err != nil {
			log.Errorf("invalid --report-frequency: %s", err.Error())
			return
		}
	}
	if len(_applicationcostprofilerReportId) > 0 {
		input.ReportId = aws.String(_applicationcostprofilerReportId)
	}

	if resp, err := client.PutReportDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates existing report in AWS Application Cost Profiler.
func applicationcostprofiler_UpdateReportDefinition(cfg aws.Config, client *applicationcostprofiler.Client) {
	input := &applicationcostprofiler.UpdateReportDefinitionInput{
		// DestinationS3Location: *types.S3Location, // Required
		// Format: types.Format, // Required
		// ReportDescription: *string, // Required
		// ReportFrequency: types.ReportFrequency, // Required
		// ReportId: *string, // Required
	}

	if len(_applicationcostprofilerDestinationS3Location) > 0 {
		if err := assignInputField(input, "DestinationS3Location", _applicationcostprofilerDestinationS3Location); err != nil {
			log.Errorf("invalid --destination-s3-location: %s", err.Error())
			return
		}
	}
	if len(_applicationcostprofilerFormat) > 0 {
		if err := assignInputField(input, "Format", _applicationcostprofilerFormat); err != nil {
			log.Errorf("invalid --format: %s", err.Error())
			return
		}
	}
	if len(_applicationcostprofilerReportDescription) > 0 {
		input.ReportDescription = aws.String(_applicationcostprofilerReportDescription)
	}
	if len(_applicationcostprofilerReportFrequency) > 0 {
		if err := assignInputField(input, "ReportFrequency", _applicationcostprofilerReportFrequency); err != nil {
			log.Errorf("invalid --report-frequency: %s", err.Error())
			return
		}
	}
	if len(_applicationcostprofilerReportId) > 0 {
		input.ReportId = aws.String(_applicationcostprofilerReportId)
	}

	if resp, err := client.UpdateReportDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_applicationcostprofilerCmd)
	_applicationcostprofilerCmd.Flags().SortFlags = false

	_applicationcostprofilerCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_applicationcostprofilerCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_applicationcostprofilerCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_applicationcostprofilerCmd.Flags().StringVarP(&_applicationcostprofilerDestinationS3Location, "destination-s3-location", "", "", "Destination S3 Location")
	_applicationcostprofilerCmd.Flags().StringVarP(&_applicationcostprofilerFormat, "format", "", "", "Format")
	_applicationcostprofilerCmd.Flags().StringVarP(&_applicationcostprofilerMaxResults, "max-results", "", "", "Max Results")
	_applicationcostprofilerCmd.Flags().StringVarP(&_applicationcostprofilerNextToken, "next-token", "", "", "Next Token")
	_applicationcostprofilerCmd.Flags().StringVarP(&_applicationcostprofilerReportDescription, "report-description", "", "", "Report Description")
	_applicationcostprofilerCmd.Flags().StringVarP(&_applicationcostprofilerReportFrequency, "report-frequency", "", "", "Report Frequency")
	_applicationcostprofilerCmd.Flags().StringVarP(&_applicationcostprofilerReportId, "report-id", "", "", "Report ID")
	_applicationcostprofilerCmd.Flags().StringVarP(&_applicationcostprofilerSourceS3Location, "source-s3-location", "", "", "Source S3 Location")

	_applicationcostprofilerCmd.Flags().BoolVarP(&_applicationcostprofilerDeleteReportDefinition, "delete-report-definition", "", false, "Delete Report Definition")
	_applicationcostprofilerCmd.Flags().BoolVarP(&_applicationcostprofilerGetReportDefinition, "get-report-definition", "", false, "Get Report Definition")
	_applicationcostprofilerCmd.Flags().BoolVarP(&_applicationcostprofilerImportApplicationUsage, "import-application-usage", "", false, "Import Application Usage")
	_applicationcostprofilerCmd.Flags().BoolVarP(&_applicationcostprofilerListReportDefinitions, "list-report-definitions", "", false, "List Report Definitions")
	_applicationcostprofilerCmd.Flags().BoolVarP(&_applicationcostprofilerPutReportDefinition, "put-report-definition", "", false, "Put Report Definition")
	_applicationcostprofilerCmd.Flags().BoolVarP(&_applicationcostprofilerUpdateReportDefinition, "update-report-definition", "", false, "Update Report Definition")

}
