package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalytics"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// kinesisanalyticsCmd represents the kinesisanalytics command
var _kinesisanalyticsCmd = &cobra.Command{
	Use:   "kinesisanalytics",
	Short: "AWS kinesisanalytics CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := kinesisanalytics.NewFromConfig(cfg)
		if _kinesisanalyticsAddApplicationCloudWatchLoggingOption {
			kinesisanalytics_AddApplicationCloudWatchLoggingOption(cfg, client)
			return
		}
		if _kinesisanalyticsAddApplicationInput {
			kinesisanalytics_AddApplicationInput(cfg, client)
			return
		}
		if _kinesisanalyticsAddApplicationInputProcessingConfiguration {
			kinesisanalytics_AddApplicationInputProcessingConfiguration(cfg, client)
			return
		}
		if _kinesisanalyticsAddApplicationOutput {
			kinesisanalytics_AddApplicationOutput(cfg, client)
			return
		}
		if _kinesisanalyticsAddApplicationReferenceDataSource {
			kinesisanalytics_AddApplicationReferenceDataSource(cfg, client)
			return
		}
		if _kinesisanalyticsCreateApplication {
			kinesisanalytics_CreateApplication(cfg, client)
			return
		}
		if _kinesisanalyticsDeleteApplication {
			kinesisanalytics_DeleteApplication(cfg, client)
			return
		}
		if _kinesisanalyticsDeleteApplicationCloudWatchLoggingOption {
			kinesisanalytics_DeleteApplicationCloudWatchLoggingOption(cfg, client)
			return
		}
		if _kinesisanalyticsDeleteApplicationInputProcessingConfiguration {
			kinesisanalytics_DeleteApplicationInputProcessingConfiguration(cfg, client)
			return
		}
		if _kinesisanalyticsDeleteApplicationOutput {
			kinesisanalytics_DeleteApplicationOutput(cfg, client)
			return
		}
		if _kinesisanalyticsDeleteApplicationReferenceDataSource {
			kinesisanalytics_DeleteApplicationReferenceDataSource(cfg, client)
			return
		}
		if _kinesisanalyticsDescribeApplication {
			kinesisanalytics_DescribeApplication(cfg, client)
			return
		}
		if _kinesisanalyticsDiscoverInputSchema {
			kinesisanalytics_DiscoverInputSchema(cfg, client)
			return
		}
		if _kinesisanalyticsListApplications {
			kinesisanalytics_ListApplications(cfg, client)
			return
		}
		if _kinesisanalyticsListTagsForResource {
			kinesisanalytics_ListTagsForResource(cfg, client)
			return
		}
		if _kinesisanalyticsStartApplication {
			kinesisanalytics_StartApplication(cfg, client)
			return
		}
		if _kinesisanalyticsStopApplication {
			kinesisanalytics_StopApplication(cfg, client)
			return
		}
		if _kinesisanalyticsTagResource {
			kinesisanalytics_TagResource(cfg, client)
			return
		}
		if _kinesisanalyticsUntagResource {
			kinesisanalytics_UntagResource(cfg, client)
			return
		}
		if _kinesisanalyticsUpdateApplication {
			kinesisanalytics_UpdateApplication(cfg, client)
			return
		}

	},
}

var (
	_kinesisanalyticsAddApplicationCloudWatchLoggingOption         bool
	_kinesisanalyticsAddApplicationInput                           bool
	_kinesisanalyticsAddApplicationInputProcessingConfiguration    bool
	_kinesisanalyticsAddApplicationOutput                          bool
	_kinesisanalyticsAddApplicationReferenceDataSource             bool
	_kinesisanalyticsCreateApplication                             bool
	_kinesisanalyticsDeleteApplication                             bool
	_kinesisanalyticsDeleteApplicationCloudWatchLoggingOption      bool
	_kinesisanalyticsDeleteApplicationInputProcessingConfiguration bool
	_kinesisanalyticsDeleteApplicationOutput                       bool
	_kinesisanalyticsDeleteApplicationReferenceDataSource          bool
	_kinesisanalyticsDescribeApplication                           bool
	_kinesisanalyticsDiscoverInputSchema                           bool
	_kinesisanalyticsListApplications                              bool
	_kinesisanalyticsListTagsForResource                           bool
	_kinesisanalyticsStartApplication                              bool
	_kinesisanalyticsStopApplication                               bool
	_kinesisanalyticsTagResource                                   bool
	_kinesisanalyticsUntagResource                                 bool
	_kinesisanalyticsUpdateApplication                             bool

	_kinesisanalyticsApplicationCode                    string
	_kinesisanalyticsApplicationDescription             string
	_kinesisanalyticsApplicationName                    string
	_kinesisanalyticsApplicationUpdate                  string
	_kinesisanalyticsCloudWatchLoggingOption            string
	_kinesisanalyticsCloudWatchLoggingOptionId          string
	_kinesisanalyticsCloudWatchLoggingOptions           string
	_kinesisanalyticsCreateTimestamp                    string
	_kinesisanalyticsCurrentApplicationVersionId        string
	_kinesisanalyticsExclusiveStartApplicationName      string
	_kinesisanalyticsInput                              string
	_kinesisanalyticsInputConfigurations                string
	_kinesisanalyticsInputId                            string
	_kinesisanalyticsInputProcessingConfiguration       string
	_kinesisanalyticsInputStartingPositionConfiguration string
	_kinesisanalyticsInputs                             string
	_kinesisanalyticsLimit                              string
	_kinesisanalyticsOutputId                           string
	_kinesisanalyticsOutputs                            string
	_kinesisanalyticsReferenceDataSource                string
	_kinesisanalyticsReferenceId                        string
	_kinesisanalyticsResourceARN                        string
	_kinesisanalyticsRoleARN                            string
	_kinesisanalyticsS3Configuration                    string
	_kinesisanalyticsTagKeys                            []string
	_kinesisanalyticsTags                               string
)

// This documentation is for version 1 of the Amazon Kinesis Data Analytics API,
// which only supports SQL applications. Version 2 of the API supports SQL and Java
// applications. For more information about version 2, see Amazon Kinesis Data Analytics API V2 Documentation.
//
// Adds a CloudWatch log stream to monitor application configuration errors. For
// more information about using CloudWatch log streams with Amazon Kinesis
// Analytics applications, see [Working with Amazon CloudWatch Logs].
//
// [Working with Amazon CloudWatch Logs]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/cloudwatch-logs.html
func kinesisanalytics_AddApplicationCloudWatchLoggingOption(cfg aws.Config, client *kinesisanalytics.Client) {
	input := &kinesisanalytics.AddApplicationCloudWatchLoggingOptionInput{
		// ApplicationName: *string, // Required
		// CloudWatchLoggingOption: *types.CloudWatchLoggingOption, // Required
		// CurrentApplicationVersionId: *int64, // Required
	}

	if len(_kinesisanalyticsApplicationName) > 0 {
		input.ApplicationName = aws.String(_kinesisanalyticsApplicationName)
	}
	if len(_kinesisanalyticsCloudWatchLoggingOption) > 0 {
		if err := assignInputField(input, "CloudWatchLoggingOption", _kinesisanalyticsCloudWatchLoggingOption); err != nil {
			log.Errorf("invalid --cloud-watch-logging-option: %s", err.Error())
			return
		}
	}
	if len(_kinesisanalyticsCurrentApplicationVersionId) > 0 {
		if err := assignInputField(input, "CurrentApplicationVersionId", _kinesisanalyticsCurrentApplicationVersionId); err != nil {
			log.Errorf("invalid --current-application-version-id: %s", err.Error())
			return
		}
	}

	if resp, err := client.AddApplicationCloudWatchLoggingOption(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This documentation is for version 1 of the Amazon Kinesis Data Analytics API,
// which only supports SQL applications. Version 2 of the API supports SQL and Java
// applications. For more information about version 2, see Amazon Kinesis Data Analytics API V2 Documentation.
//
// Adds a streaming source to your Amazon Kinesis application. For conceptual
// information, see [Configuring Application Input].
//
// You can add a streaming source either when you create an application or you can
// use this operation to add a streaming source after you create an application.
// For more information, see [CreateApplication].
//
// Any configuration update, including adding a streaming source using this
// operation, results in a new version of the application. You can use the [DescribeApplication]
// operation to find the current application version.
//
// This operation requires permissions to perform the
// kinesisanalytics:AddApplicationInput action.
//
// [CreateApplication]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_CreateApplication.html
// [DescribeApplication]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_DescribeApplication.html
// [Configuring Application Input]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-input.html
func kinesisanalytics_AddApplicationInput(cfg aws.Config, client *kinesisanalytics.Client) {
	input := &kinesisanalytics.AddApplicationInputInput{
		// ApplicationName: *string, // Required
		// CurrentApplicationVersionId: *int64, // Required
		// Input: *types.Input, // Required
	}

	if len(_kinesisanalyticsApplicationName) > 0 {
		input.ApplicationName = aws.String(_kinesisanalyticsApplicationName)
	}
	if len(_kinesisanalyticsCurrentApplicationVersionId) > 0 {
		if err := assignInputField(input, "CurrentApplicationVersionId", _kinesisanalyticsCurrentApplicationVersionId); err != nil {
			log.Errorf("invalid --current-application-version-id: %s", err.Error())
			return
		}
	}
	if len(_kinesisanalyticsInput) > 0 {
		if err := assignInputField(input, "Input", _kinesisanalyticsInput); err != nil {
			log.Errorf("invalid --input: %s", err.Error())
			return
		}
	}

	if resp, err := client.AddApplicationInput(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This documentation is for version 1 of the Amazon Kinesis Data Analytics API,
// which only supports SQL applications. Version 2 of the API supports SQL and Java
// applications. For more information about version 2, see Amazon Kinesis Data Analytics API V2 Documentation.
//
// Adds an [InputProcessingConfiguration] to an application. An input processor preprocesses records on the
// input stream before the application's SQL code executes. Currently, the only
// input processor available is [AWS Lambda].
//
// [AWS Lambda]: https://docs.aws.amazon.com/lambda/
// [InputProcessingConfiguration]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_InputProcessingConfiguration.html
func kinesisanalytics_AddApplicationInputProcessingConfiguration(cfg aws.Config, client *kinesisanalytics.Client) {
	input := &kinesisanalytics.AddApplicationInputProcessingConfigurationInput{
		// ApplicationName: *string, // Required
		// CurrentApplicationVersionId: *int64, // Required
		// InputId: *string, // Required
		// InputProcessingConfiguration: *types.InputProcessingConfiguration, // Required
	}

	if len(_kinesisanalyticsApplicationName) > 0 {
		input.ApplicationName = aws.String(_kinesisanalyticsApplicationName)
	}
	if len(_kinesisanalyticsCurrentApplicationVersionId) > 0 {
		if err := assignInputField(input, "CurrentApplicationVersionId", _kinesisanalyticsCurrentApplicationVersionId); err != nil {
			log.Errorf("invalid --current-application-version-id: %s", err.Error())
			return
		}
	}
	if len(_kinesisanalyticsInputId) > 0 {
		input.InputId = aws.String(_kinesisanalyticsInputId)
	}
	if len(_kinesisanalyticsInputProcessingConfiguration) > 0 {
		if err := assignInputField(input, "InputProcessingConfiguration", _kinesisanalyticsInputProcessingConfiguration); err != nil {
			log.Errorf("invalid --input-processing-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.AddApplicationInputProcessingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This documentation is for version 1 of the Amazon Kinesis Data Analytics API,
// which only supports SQL applications. Version 2 of the API supports SQL and Java
// applications. For more information about version 2, see Amazon Kinesis Data Analytics API V2 Documentation.
//
// Adds an external destination to your Amazon Kinesis Analytics application.
//
// If you want Amazon Kinesis Analytics to deliver data from an in-application
// stream within your application to an external destination (such as an Amazon
// Kinesis stream, an Amazon Kinesis Firehose delivery stream, or an AWS Lambda
// function), you add the relevant configuration to your application using this
// operation. You can configure one or more outputs for your application. Each
// output configuration maps an in-application stream and an external destination.
//
// You can use one of the output configurations to deliver data from your
// in-application error stream to an external destination so that you can analyze
// the errors. For more information, see [Understanding Application Output (Destination)].
//
// Any configuration update, including adding a streaming source using this
// operation, results in a new version of the application. You can use the [DescribeApplication]
// operation to find the current application version.
//
// For the limits on the number of application inputs and outputs you can
// configure, see [Limits].
//
// This operation requires permissions to perform the
// kinesisanalytics:AddApplicationOutput action.
//
// [Limits]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/limits.html
// [Understanding Application Output (Destination)]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-output.html
// [DescribeApplication]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_DescribeApplication.html
func kinesisanalytics_AddApplicationOutput(cfg aws.Config, client *kinesisanalytics.Client) {
	input := &kinesisanalytics.AddApplicationOutputInput{
		// ApplicationName: *string, // Required
		// CurrentApplicationVersionId: *int64, // Required
		// Output: *types.Output, // Required
	}

	if len(_kinesisanalyticsApplicationName) > 0 {
		input.ApplicationName = aws.String(_kinesisanalyticsApplicationName)
	}
	if len(_kinesisanalyticsCurrentApplicationVersionId) > 0 {
		if err := assignInputField(input, "CurrentApplicationVersionId", _kinesisanalyticsCurrentApplicationVersionId); err != nil {
			log.Errorf("invalid --current-application-version-id: %s", err.Error())
			return
		}
	}

	if resp, err := client.AddApplicationOutput(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This documentation is for version 1 of the Amazon Kinesis Data Analytics API,
// which only supports SQL applications. Version 2 of the API supports SQL and Java
// applications. For more information about version 2, see Amazon Kinesis Data Analytics API V2 Documentation.
//
// Adds a reference data source to an existing application.
//
// Amazon Kinesis Analytics reads reference data (that is, an Amazon S3 object)
// and creates an in-application table within your application. In the request, you
// provide the source (S3 bucket name and object key name), name of the
// in-application table to create, and the necessary mapping information that
// describes how data in Amazon S3 object maps to columns in the resulting
// in-application table.
//
// For conceptual information, see [Configuring Application Input]. For the limits on data sources you can add to
// your application, see [Limits].
//
// This operation requires permissions to perform the
// kinesisanalytics:AddApplicationOutput action.
//
// [Limits]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/limits.html
// [Configuring Application Input]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-input.html
func kinesisanalytics_AddApplicationReferenceDataSource(cfg aws.Config, client *kinesisanalytics.Client) {
	input := &kinesisanalytics.AddApplicationReferenceDataSourceInput{
		// ApplicationName: *string, // Required
		// CurrentApplicationVersionId: *int64, // Required
		// ReferenceDataSource: *types.ReferenceDataSource, // Required
	}

	if len(_kinesisanalyticsApplicationName) > 0 {
		input.ApplicationName = aws.String(_kinesisanalyticsApplicationName)
	}
	if len(_kinesisanalyticsCurrentApplicationVersionId) > 0 {
		if err := assignInputField(input, "CurrentApplicationVersionId", _kinesisanalyticsCurrentApplicationVersionId); err != nil {
			log.Errorf("invalid --current-application-version-id: %s", err.Error())
			return
		}
	}
	if len(_kinesisanalyticsReferenceDataSource) > 0 {
		if err := assignInputField(input, "ReferenceDataSource", _kinesisanalyticsReferenceDataSource); err != nil {
			log.Errorf("invalid --reference-data-source: %s", err.Error())
			return
		}
	}

	if resp, err := client.AddApplicationReferenceDataSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This documentation is for version 1 of the Amazon Kinesis Data Analytics API,
// which only supports SQL applications. Version 2 of the API supports SQL and Java
// applications. For more information about version 2, see Amazon Kinesis Data Analytics API V2 Documentation.
//
// Creates an Amazon Kinesis Analytics application. You can configure each
// application with one streaming source as input, application code to process the
// input, and up to three destinations where you want Amazon Kinesis Analytics to
// write the output data from your application. For an overview, see [How it Works].
//
// In the input configuration, you map the streaming source to an in-application
// stream, which you can think of as a constantly updating table. In the mapping,
// you must provide a schema for the in-application stream and map each data column
// in the in-application stream to a data element in the streaming source.
//
// Your application code is one or more SQL statements that read input data,
// transform it, and generate output. Your application code can create one or more
// SQL artifacts like SQL streams or pumps.
//
// In the output configuration, you can configure the application to write data
// from in-application streams created in your applications to up to three
// destinations.
//
// To read data from your source stream or write data to destination streams,
// Amazon Kinesis Analytics needs your permissions. You grant these permissions by
// creating IAM roles. This operation requires permissions to perform the
// kinesisanalytics:CreateApplication action.
//
// For introductory exercises to create an Amazon Kinesis Analytics application,
// see [Getting Started].
//
// [Getting Started]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/getting-started.html
// [How it Works]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works.html
func kinesisanalytics_CreateApplication(cfg aws.Config, client *kinesisanalytics.Client) {
	input := &kinesisanalytics.CreateApplicationInput{
		// ApplicationName: *string, // Required
	}

	if len(_kinesisanalyticsApplicationName) > 0 {
		input.ApplicationName = aws.String(_kinesisanalyticsApplicationName)
	}
	if len(_kinesisanalyticsApplicationCode) > 0 {
		input.ApplicationCode = aws.String(_kinesisanalyticsApplicationCode)
	}
	if len(_kinesisanalyticsApplicationDescription) > 0 {
		input.ApplicationDescription = aws.String(_kinesisanalyticsApplicationDescription)
	}
	if len(_kinesisanalyticsCloudWatchLoggingOptions) > 0 {
		if err := assignInputField(input, "CloudWatchLoggingOptions", _kinesisanalyticsCloudWatchLoggingOptions); err != nil {
			log.Errorf("invalid --cloud-watch-logging-options: %s", err.Error())
			return
		}
	}
	if len(_kinesisanalyticsInputs) > 0 {
		if err := assignInputField(input, "Inputs", _kinesisanalyticsInputs); err != nil {
			log.Errorf("invalid --inputs: %s", err.Error())
			return
		}
	}
	if len(_kinesisanalyticsOutputs) > 0 {
		if err := assignInputField(input, "Outputs", _kinesisanalyticsOutputs); err != nil {
			log.Errorf("invalid --outputs: %s", err.Error())
			return
		}
	}
	if len(_kinesisanalyticsTags) > 0 {
		if err := assignInputField(input, "Tags", _kinesisanalyticsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This documentation is for version 1 of the Amazon Kinesis Data Analytics API,
// which only supports SQL applications. Version 2 of the API supports SQL and Java
// applications. For more information about version 2, see Amazon Kinesis Data Analytics API V2 Documentation.
//
// Deletes the specified application. Amazon Kinesis Analytics halts application
// execution and deletes the application, including any application artifacts (such
// as in-application streams, reference table, and application code).
//
// This operation requires permissions to perform the
// kinesisanalytics:DeleteApplication action.
func kinesisanalytics_DeleteApplication(cfg aws.Config, client *kinesisanalytics.Client) {
	input := &kinesisanalytics.DeleteApplicationInput{
		// ApplicationName: *string, // Required
		// CreateTimestamp: *time.Time, // Required
	}

	if len(_kinesisanalyticsApplicationName) > 0 {
		input.ApplicationName = aws.String(_kinesisanalyticsApplicationName)
	}
	if len(_kinesisanalyticsCreateTimestamp) > 0 {
		if err := assignInputField(input, "CreateTimestamp", _kinesisanalyticsCreateTimestamp); err != nil {
			log.Errorf("invalid --create-timestamp: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This documentation is for version 1 of the Amazon Kinesis Data Analytics API,
// which only supports SQL applications. Version 2 of the API supports SQL and Java
// applications. For more information about version 2, see Amazon Kinesis Data Analytics API V2 Documentation.
//
// Deletes a CloudWatch log stream from an application. For more information about
// using CloudWatch log streams with Amazon Kinesis Analytics applications, see [Working with Amazon CloudWatch Logs].
//
// [Working with Amazon CloudWatch Logs]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/cloudwatch-logs.html
func kinesisanalytics_DeleteApplicationCloudWatchLoggingOption(cfg aws.Config, client *kinesisanalytics.Client) {
	input := &kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionInput{
		// ApplicationName: *string, // Required
		// CloudWatchLoggingOptionId: *string, // Required
		// CurrentApplicationVersionId: *int64, // Required
	}

	if len(_kinesisanalyticsApplicationName) > 0 {
		input.ApplicationName = aws.String(_kinesisanalyticsApplicationName)
	}
	if len(_kinesisanalyticsCloudWatchLoggingOptionId) > 0 {
		input.CloudWatchLoggingOptionId = aws.String(_kinesisanalyticsCloudWatchLoggingOptionId)
	}
	if len(_kinesisanalyticsCurrentApplicationVersionId) > 0 {
		if err := assignInputField(input, "CurrentApplicationVersionId", _kinesisanalyticsCurrentApplicationVersionId); err != nil {
			log.Errorf("invalid --current-application-version-id: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteApplicationCloudWatchLoggingOption(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This documentation is for version 1 of the Amazon Kinesis Data Analytics API,
// which only supports SQL applications. Version 2 of the API supports SQL and Java
// applications. For more information about version 2, see Amazon Kinesis Data Analytics API V2 Documentation.
//
// Deletes an [InputProcessingConfiguration] from an input.
//
// [InputProcessingConfiguration]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_InputProcessingConfiguration.html
func kinesisanalytics_DeleteApplicationInputProcessingConfiguration(cfg aws.Config, client *kinesisanalytics.Client) {
	input := &kinesisanalytics.DeleteApplicationInputProcessingConfigurationInput{
		// ApplicationName: *string, // Required
		// CurrentApplicationVersionId: *int64, // Required
		// InputId: *string, // Required
	}

	if len(_kinesisanalyticsApplicationName) > 0 {
		input.ApplicationName = aws.String(_kinesisanalyticsApplicationName)
	}
	if len(_kinesisanalyticsCurrentApplicationVersionId) > 0 {
		if err := assignInputField(input, "CurrentApplicationVersionId", _kinesisanalyticsCurrentApplicationVersionId); err != nil {
			log.Errorf("invalid --current-application-version-id: %s", err.Error())
			return
		}
	}
	if len(_kinesisanalyticsInputId) > 0 {
		input.InputId = aws.String(_kinesisanalyticsInputId)
	}

	if resp, err := client.DeleteApplicationInputProcessingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This documentation is for version 1 of the Amazon Kinesis Data Analytics API,
// which only supports SQL applications. Version 2 of the API supports SQL and Java
// applications. For more information about version 2, see Amazon Kinesis Data Analytics API V2 Documentation.
//
// Deletes output destination configuration from your application configuration.
// Amazon Kinesis Analytics will no longer write data from the corresponding
// in-application stream to the external output destination.
//
// This operation requires permissions to perform the
// kinesisanalytics:DeleteApplicationOutput action.
func kinesisanalytics_DeleteApplicationOutput(cfg aws.Config, client *kinesisanalytics.Client) {
	input := &kinesisanalytics.DeleteApplicationOutputInput{
		// ApplicationName: *string, // Required
		// CurrentApplicationVersionId: *int64, // Required
		// OutputId: *string, // Required
	}

	if len(_kinesisanalyticsApplicationName) > 0 {
		input.ApplicationName = aws.String(_kinesisanalyticsApplicationName)
	}
	if len(_kinesisanalyticsCurrentApplicationVersionId) > 0 {
		if err := assignInputField(input, "CurrentApplicationVersionId", _kinesisanalyticsCurrentApplicationVersionId); err != nil {
			log.Errorf("invalid --current-application-version-id: %s", err.Error())
			return
		}
	}
	if len(_kinesisanalyticsOutputId) > 0 {
		input.OutputId = aws.String(_kinesisanalyticsOutputId)
	}

	if resp, err := client.DeleteApplicationOutput(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This documentation is for version 1 of the Amazon Kinesis Data Analytics API,
// which only supports SQL applications. Version 2 of the API supports SQL and Java
// applications. For more information about version 2, see Amazon Kinesis Data Analytics API V2 Documentation.
//
// Deletes a reference data source configuration from the specified application
// configuration.
//
// If the application is running, Amazon Kinesis Analytics immediately removes the
// in-application table that you created using the [AddApplicationReferenceDataSource]operation.
//
// This operation requires permissions to perform the
// kinesisanalytics.DeleteApplicationReferenceDataSource action.
//
// [AddApplicationReferenceDataSource]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_AddApplicationReferenceDataSource.html
func kinesisanalytics_DeleteApplicationReferenceDataSource(cfg aws.Config, client *kinesisanalytics.Client) {
	input := &kinesisanalytics.DeleteApplicationReferenceDataSourceInput{
		// ApplicationName: *string, // Required
		// CurrentApplicationVersionId: *int64, // Required
		// ReferenceId: *string, // Required
	}

	if len(_kinesisanalyticsApplicationName) > 0 {
		input.ApplicationName = aws.String(_kinesisanalyticsApplicationName)
	}
	if len(_kinesisanalyticsCurrentApplicationVersionId) > 0 {
		if err := assignInputField(input, "CurrentApplicationVersionId", _kinesisanalyticsCurrentApplicationVersionId); err != nil {
			log.Errorf("invalid --current-application-version-id: %s", err.Error())
			return
		}
	}
	if len(_kinesisanalyticsReferenceId) > 0 {
		input.ReferenceId = aws.String(_kinesisanalyticsReferenceId)
	}

	if resp, err := client.DeleteApplicationReferenceDataSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This documentation is for version 1 of the Amazon Kinesis Data Analytics API,
// which only supports SQL applications. Version 2 of the API supports SQL and Java
// applications. For more information about version 2, see Amazon Kinesis Data Analytics API V2 Documentation.
//
// Returns information about a specific Amazon Kinesis Analytics application.
//
// If you want to retrieve a list of all applications in your account, use the [ListApplications]
// operation.
//
// This operation requires permissions to perform the
// kinesisanalytics:DescribeApplication action. You can use DescribeApplication to
// get the current application versionId, which you need to call other operations
// such as Update .
//
// [ListApplications]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_ListApplications.html
func kinesisanalytics_DescribeApplication(cfg aws.Config, client *kinesisanalytics.Client) {
	input := &kinesisanalytics.DescribeApplicationInput{
		// ApplicationName: *string, // Required
	}

	if len(_kinesisanalyticsApplicationName) > 0 {
		input.ApplicationName = aws.String(_kinesisanalyticsApplicationName)
	}

	if resp, err := client.DescribeApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This documentation is for version 1 of the Amazon Kinesis Data Analytics API,
// which only supports SQL applications. Version 2 of the API supports SQL and Java
// applications. For more information about version 2, see Amazon Kinesis Data Analytics API V2 Documentation.
//
// Infers a schema by evaluating sample records on the specified streaming source
// (Amazon Kinesis stream or Amazon Kinesis Firehose delivery stream) or S3 object.
// In the response, the operation returns the inferred schema and also the sample
// records that the operation used to infer the schema.
//
// You can use the inferred schema when configuring a streaming source for your
// application. For conceptual information, see [Configuring Application Input]. Note that when you create an
// application using the Amazon Kinesis Analytics console, the console uses this
// operation to infer a schema and show it in the console user interface.
//
// This operation requires permissions to perform the
// kinesisanalytics:DiscoverInputSchema action.
//
// [Configuring Application Input]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-input.html
func kinesisanalytics_DiscoverInputSchema(cfg aws.Config, client *kinesisanalytics.Client) {
	input := &kinesisanalytics.DiscoverInputSchemaInput{}

	if len(_kinesisanalyticsInputProcessingConfiguration) > 0 {
		if err := assignInputField(input, "InputProcessingConfiguration", _kinesisanalyticsInputProcessingConfiguration); err != nil {
			log.Errorf("invalid --input-processing-configuration: %s", err.Error())
			return
		}
	}
	if len(_kinesisanalyticsInputStartingPositionConfiguration) > 0 {
		if err := assignInputField(input, "InputStartingPositionConfiguration", _kinesisanalyticsInputStartingPositionConfiguration); err != nil {
			log.Errorf("invalid --input-starting-position-configuration: %s", err.Error())
			return
		}
	}
	if len(_kinesisanalyticsResourceARN) > 0 {
		input.ResourceARN = aws.String(_kinesisanalyticsResourceARN)
	}
	if len(_kinesisanalyticsRoleARN) > 0 {
		input.RoleARN = aws.String(_kinesisanalyticsRoleARN)
	}
	if len(_kinesisanalyticsS3Configuration) > 0 {
		if err := assignInputField(input, "S3Configuration", _kinesisanalyticsS3Configuration); err != nil {
			log.Errorf("invalid --s3-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.DiscoverInputSchema(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This documentation is for version 1 of the Amazon Kinesis Data Analytics API,
// which only supports SQL applications. Version 2 of the API supports SQL and Java
// applications. For more information about version 2, see Amazon Kinesis Data Analytics API V2 Documentation.
//
// Returns a list of Amazon Kinesis Analytics applications in your account. For
// each application, the response includes the application name, Amazon Resource
// Name (ARN), and status.
//
// If the response returns the HasMoreApplications value as true,
//
// you can send another request by adding the ExclusiveStartApplicationName in the
// request body, and set the value of this to the last application name from the
// previous response.
//
// If you want detailed information about a specific application, use [DescribeApplication].
//
// This operation requires permissions to perform the
// kinesisanalytics:ListApplications action.
//
// [DescribeApplication]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_DescribeApplication.html
func kinesisanalytics_ListApplications(cfg aws.Config, client *kinesisanalytics.Client) {
	input := &kinesisanalytics.ListApplicationsInput{}

	if len(_kinesisanalyticsExclusiveStartApplicationName) > 0 {
		input.ExclusiveStartApplicationName = aws.String(_kinesisanalyticsExclusiveStartApplicationName)
	}
	if len(_kinesisanalyticsLimit) > 0 {
		if err := assignInputField(input, "Limit", _kinesisanalyticsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListApplications(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the list of key-value tags assigned to the application. For more
// information, see [Using Tagging].
//
// [Using Tagging]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-tagging.html
func kinesisanalytics_ListTagsForResource(cfg aws.Config, client *kinesisanalytics.Client) {
	input := &kinesisanalytics.ListTagsForResourceInput{
		// ResourceARN: *string, // Required
	}

	if len(_kinesisanalyticsResourceARN) > 0 {
		input.ResourceARN = aws.String(_kinesisanalyticsResourceARN)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This documentation is for version 1 of the Amazon Kinesis Data Analytics API,
// which only supports SQL applications. Version 2 of the API supports SQL and Java
// applications. For more information about version 2, see Amazon Kinesis Data Analytics API V2 Documentation.
//
// Starts the specified Amazon Kinesis Analytics application. After creating an
// application, you must exclusively call this operation to start your application.
//
// After the application starts, it begins consuming the input data, processes it,
// and writes the output to the configured destination.
//
// The application status must be READY for you to start an application. You can
// get the application status in the console or using the [DescribeApplication]operation.
//
// After you start the application, you can stop the application from processing
// the input by calling the [StopApplication]operation.
//
// This operation requires permissions to perform the
// kinesisanalytics:StartApplication action.
//
// [DescribeApplication]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_DescribeApplication.html
// [StopApplication]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_StopApplication.html
func kinesisanalytics_StartApplication(cfg aws.Config, client *kinesisanalytics.Client) {
	input := &kinesisanalytics.StartApplicationInput{
		// ApplicationName: *string, // Required
		// InputConfigurations: []types.InputConfiguration, // Required
	}

	if len(_kinesisanalyticsApplicationName) > 0 {
		input.ApplicationName = aws.String(_kinesisanalyticsApplicationName)
	}
	if len(_kinesisanalyticsInputConfigurations) > 0 {
		if err := assignInputField(input, "InputConfigurations", _kinesisanalyticsInputConfigurations); err != nil {
			log.Errorf("invalid --input-configurations: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This documentation is for version 1 of the Amazon Kinesis Data Analytics API,
// which only supports SQL applications. Version 2 of the API supports SQL and Java
// applications. For more information about version 2, see Amazon Kinesis Data Analytics API V2 Documentation.
//
// Stops the application from processing input data. You can stop an application
// only if it is in the running state. You can use the [DescribeApplication]operation to find the
// application state. After the application is stopped, Amazon Kinesis Analytics
// stops reading data from the input, the application stops processing data, and
// there is no output written to the destination.
//
// This operation requires permissions to perform the
// kinesisanalytics:StopApplication action.
//
// [DescribeApplication]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_DescribeApplication.html
func kinesisanalytics_StopApplication(cfg aws.Config, client *kinesisanalytics.Client) {
	input := &kinesisanalytics.StopApplicationInput{
		// ApplicationName: *string, // Required
	}

	if len(_kinesisanalyticsApplicationName) > 0 {
		input.ApplicationName = aws.String(_kinesisanalyticsApplicationName)
	}

	if resp, err := client.StopApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds one or more key-value tags to a Kinesis Analytics application. Note that
// the maximum number of application tags includes system tags. The maximum number
// of user-defined application tags is 50. For more information, see [Using Tagging].
//
// [Using Tagging]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-tagging.html
func kinesisanalytics_TagResource(cfg aws.Config, client *kinesisanalytics.Client) {
	input := &kinesisanalytics.TagResourceInput{
		// ResourceARN: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_kinesisanalyticsResourceARN) > 0 {
		input.ResourceARN = aws.String(_kinesisanalyticsResourceARN)
	}
	if len(_kinesisanalyticsTags) > 0 {
		if err := assignInputField(input, "Tags", _kinesisanalyticsTags); err != nil {
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

// Removes one or more tags from a Kinesis Analytics application. For more
// information, see [Using Tagging].
//
// [Using Tagging]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-tagging.html
func kinesisanalytics_UntagResource(cfg aws.Config, client *kinesisanalytics.Client) {
	input := &kinesisanalytics.UntagResourceInput{
		// ResourceARN: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_kinesisanalyticsResourceARN) > 0 {
		input.ResourceARN = aws.String(_kinesisanalyticsResourceARN)
	}
	if len(_kinesisanalyticsTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _kinesisanalyticsTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This documentation is for version 1 of the Amazon Kinesis Data Analytics API,
// which only supports SQL applications. Version 2 of the API supports SQL and Java
// applications. For more information about version 2, see Amazon Kinesis Data Analytics API V2 Documentation.
//
// Updates an existing Amazon Kinesis Analytics application. Using this API, you
// can update application code, input configuration, and output configuration.
//
// Note that Amazon Kinesis Analytics updates the CurrentApplicationVersionId each
// time you update your application.
//
// This operation requires permission for the kinesisanalytics:UpdateApplication
// action.
func kinesisanalytics_UpdateApplication(cfg aws.Config, client *kinesisanalytics.Client) {
	input := &kinesisanalytics.UpdateApplicationInput{
		// ApplicationName: *string, // Required
		// ApplicationUpdate: *types.ApplicationUpdate, // Required
		// CurrentApplicationVersionId: *int64, // Required
	}

	if len(_kinesisanalyticsApplicationName) > 0 {
		input.ApplicationName = aws.String(_kinesisanalyticsApplicationName)
	}
	if len(_kinesisanalyticsApplicationUpdate) > 0 {
		if err := assignInputField(input, "ApplicationUpdate", _kinesisanalyticsApplicationUpdate); err != nil {
			log.Errorf("invalid --application-update: %s", err.Error())
			return
		}
	}
	if len(_kinesisanalyticsCurrentApplicationVersionId) > 0 {
		if err := assignInputField(input, "CurrentApplicationVersionId", _kinesisanalyticsCurrentApplicationVersionId); err != nil {
			log.Errorf("invalid --current-application-version-id: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_kinesisanalyticsCmd)
	_kinesisanalyticsCmd.Flags().SortFlags = false

	_kinesisanalyticsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_kinesisanalyticsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_kinesisanalyticsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_kinesisanalyticsCmd.Flags().StringVarP(&_kinesisanalyticsApplicationCode, "application-code", "", "", "Application Code")
	_kinesisanalyticsCmd.Flags().StringVarP(&_kinesisanalyticsApplicationDescription, "application-description", "", "", "Application Description")
	_kinesisanalyticsCmd.Flags().StringVarP(&_kinesisanalyticsApplicationName, "application-name", "", "", "Application Name")
	_kinesisanalyticsCmd.Flags().StringVarP(&_kinesisanalyticsApplicationUpdate, "application-update", "", "", "Application Update")
	_kinesisanalyticsCmd.Flags().StringVarP(&_kinesisanalyticsCloudWatchLoggingOption, "cloud-watch-logging-option", "", "", "Cloud Watch Logging Option")
	_kinesisanalyticsCmd.Flags().StringVarP(&_kinesisanalyticsCloudWatchLoggingOptionId, "cloud-watch-logging-option-id", "", "", "Cloud Watch Logging Option ID")
	_kinesisanalyticsCmd.Flags().StringVarP(&_kinesisanalyticsCloudWatchLoggingOptions, "cloud-watch-logging-options", "", "", "Cloud Watch Logging Options")
	_kinesisanalyticsCmd.Flags().StringVarP(&_kinesisanalyticsCreateTimestamp, "create-timestamp", "", "", "Create Timestamp")
	_kinesisanalyticsCmd.Flags().StringVarP(&_kinesisanalyticsCurrentApplicationVersionId, "current-application-version-id", "", "", "Current Application Version ID")
	_kinesisanalyticsCmd.Flags().StringVarP(&_kinesisanalyticsExclusiveStartApplicationName, "exclusive-start-application-name", "", "", "Exclusive Start Application Name")
	_kinesisanalyticsCmd.Flags().StringVarP(&_kinesisanalyticsInput, "input", "", "", "Input")
	_kinesisanalyticsCmd.Flags().StringVarP(&_kinesisanalyticsInputConfigurations, "input-configurations", "", "", "Input Configurations")
	_kinesisanalyticsCmd.Flags().StringVarP(&_kinesisanalyticsInputId, "input-id", "", "", "Input ID")
	_kinesisanalyticsCmd.Flags().StringVarP(&_kinesisanalyticsInputProcessingConfiguration, "input-processing-configuration", "", "", "Input Processing Configuration")
	_kinesisanalyticsCmd.Flags().StringVarP(&_kinesisanalyticsInputStartingPositionConfiguration, "input-starting-position-configuration", "", "", "Input Starting Position Configuration")
	_kinesisanalyticsCmd.Flags().StringVarP(&_kinesisanalyticsInputs, "inputs", "", "", "Inputs")
	_kinesisanalyticsCmd.Flags().StringVarP(&_kinesisanalyticsLimit, "limit", "", "", "Limit")
	_kinesisanalyticsCmd.Flags().StringVarP(&_kinesisanalyticsOutputId, "output-id", "", "", "Output ID")
	_kinesisanalyticsCmd.Flags().StringVarP(&_kinesisanalyticsOutputs, "outputs", "", "", "Outputs")
	_kinesisanalyticsCmd.Flags().StringVarP(&_kinesisanalyticsReferenceDataSource, "reference-data-source", "", "", "Reference Data Source")
	_kinesisanalyticsCmd.Flags().StringVarP(&_kinesisanalyticsReferenceId, "reference-id", "", "", "Reference ID")
	_kinesisanalyticsCmd.Flags().StringVarP(&_kinesisanalyticsResourceARN, "resource-arn", "", "", "Resource ARN")
	_kinesisanalyticsCmd.Flags().StringVarP(&_kinesisanalyticsRoleARN, "role-arn", "", "", "Role ARN")
	_kinesisanalyticsCmd.Flags().StringVarP(&_kinesisanalyticsS3Configuration, "s3-configuration", "", "", "S3 Configuration")
	_kinesisanalyticsCmd.Flags().StringSliceVarP(&_kinesisanalyticsTagKeys, "tag-keys", "", nil, "Tag Keys")
	_kinesisanalyticsCmd.Flags().StringVarP(&_kinesisanalyticsTags, "tags", "", "", "Tags")

	_kinesisanalyticsCmd.Flags().BoolVarP(&_kinesisanalyticsAddApplicationCloudWatchLoggingOption, "add-application-cloud-watch-logging-option", "", false, "Add Application Cloud Watch Logging Option")
	_kinesisanalyticsCmd.Flags().BoolVarP(&_kinesisanalyticsAddApplicationInput, "add-application-input", "", false, "Add Application Input")
	_kinesisanalyticsCmd.Flags().BoolVarP(&_kinesisanalyticsAddApplicationInputProcessingConfiguration, "add-application-input-processing-configuration", "", false, "Add Application Input Processing Configuration")
	_kinesisanalyticsCmd.Flags().BoolVarP(&_kinesisanalyticsAddApplicationOutput, "add-application-output", "", false, "Add Application Output")
	_kinesisanalyticsCmd.Flags().BoolVarP(&_kinesisanalyticsAddApplicationReferenceDataSource, "add-application-reference-data-source", "", false, "Add Application Reference Data Source")
	_kinesisanalyticsCmd.Flags().BoolVarP(&_kinesisanalyticsCreateApplication, "create-application", "", false, "Create Application")
	_kinesisanalyticsCmd.Flags().BoolVarP(&_kinesisanalyticsDeleteApplication, "delete-application", "", false, "Delete Application")
	_kinesisanalyticsCmd.Flags().BoolVarP(&_kinesisanalyticsDeleteApplicationCloudWatchLoggingOption, "delete-application-cloud-watch-logging-option", "", false, "Delete Application Cloud Watch Logging Option")
	_kinesisanalyticsCmd.Flags().BoolVarP(&_kinesisanalyticsDeleteApplicationInputProcessingConfiguration, "delete-application-input-processing-configuration", "", false, "Delete Application Input Processing Configuration")
	_kinesisanalyticsCmd.Flags().BoolVarP(&_kinesisanalyticsDeleteApplicationOutput, "delete-application-output", "", false, "Delete Application Output")
	_kinesisanalyticsCmd.Flags().BoolVarP(&_kinesisanalyticsDeleteApplicationReferenceDataSource, "delete-application-reference-data-source", "", false, "Delete Application Reference Data Source")
	_kinesisanalyticsCmd.Flags().BoolVarP(&_kinesisanalyticsDescribeApplication, "describe-application", "", false, "Describe Application")
	_kinesisanalyticsCmd.Flags().BoolVarP(&_kinesisanalyticsDiscoverInputSchema, "discover-input-schema", "", false, "Discover Input Schema")
	_kinesisanalyticsCmd.Flags().BoolVarP(&_kinesisanalyticsListApplications, "list-applications", "", false, "List Applications")
	_kinesisanalyticsCmd.Flags().BoolVarP(&_kinesisanalyticsListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_kinesisanalyticsCmd.Flags().BoolVarP(&_kinesisanalyticsStartApplication, "start-application", "", false, "Start Application")
	_kinesisanalyticsCmd.Flags().BoolVarP(&_kinesisanalyticsStopApplication, "stop-application", "", false, "Stop Application")
	_kinesisanalyticsCmd.Flags().BoolVarP(&_kinesisanalyticsTagResource, "tag-resource", "", false, "Tag Resource")
	_kinesisanalyticsCmd.Flags().BoolVarP(&_kinesisanalyticsUntagResource, "untag-resource", "", false, "Untag Resource")
	_kinesisanalyticsCmd.Flags().BoolVarP(&_kinesisanalyticsUpdateApplication, "update-application", "", false, "Update Application")

}
