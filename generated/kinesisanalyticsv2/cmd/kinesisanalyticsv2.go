package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// kinesisanalyticsv2Cmd represents the kinesisanalyticsv2 command
var _kinesisanalyticsv2Cmd = &cobra.Command{
	Use:   "kinesisanalyticsv2",
	Short: "AWS kinesisanalyticsv2 CLI",
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
		client := kinesisanalyticsv2.NewFromConfig(cfg)
		if _kinesisanalyticsv2AddApplicationCloudWatchLoggingOption {
			kinesisanalyticsv2_AddApplicationCloudWatchLoggingOption(cfg, client)
			return
		}
		if _kinesisanalyticsv2AddApplicationInput {
			kinesisanalyticsv2_AddApplicationInput(cfg, client)
			return
		}
		if _kinesisanalyticsv2AddApplicationInputProcessingConfiguration {
			kinesisanalyticsv2_AddApplicationInputProcessingConfiguration(cfg, client)
			return
		}
		if _kinesisanalyticsv2AddApplicationOutput {
			kinesisanalyticsv2_AddApplicationOutput(cfg, client)
			return
		}
		if _kinesisanalyticsv2AddApplicationReferenceDataSource {
			kinesisanalyticsv2_AddApplicationReferenceDataSource(cfg, client)
			return
		}
		if _kinesisanalyticsv2AddApplicationVpcConfiguration {
			kinesisanalyticsv2_AddApplicationVpcConfiguration(cfg, client)
			return
		}
		if _kinesisanalyticsv2CreateApplication {
			kinesisanalyticsv2_CreateApplication(cfg, client)
			return
		}
		if _kinesisanalyticsv2CreateApplicationPresignedUrl {
			kinesisanalyticsv2_CreateApplicationPresignedUrl(cfg, client)
			return
		}
		if _kinesisanalyticsv2CreateApplicationSnapshot {
			kinesisanalyticsv2_CreateApplicationSnapshot(cfg, client)
			return
		}
		if _kinesisanalyticsv2DeleteApplication {
			kinesisanalyticsv2_DeleteApplication(cfg, client)
			return
		}
		if _kinesisanalyticsv2DeleteApplicationCloudWatchLoggingOption {
			kinesisanalyticsv2_DeleteApplicationCloudWatchLoggingOption(cfg, client)
			return
		}
		if _kinesisanalyticsv2DeleteApplicationInputProcessingConfiguration {
			kinesisanalyticsv2_DeleteApplicationInputProcessingConfiguration(cfg, client)
			return
		}
		if _kinesisanalyticsv2DeleteApplicationOutput {
			kinesisanalyticsv2_DeleteApplicationOutput(cfg, client)
			return
		}
		if _kinesisanalyticsv2DeleteApplicationReferenceDataSource {
			kinesisanalyticsv2_DeleteApplicationReferenceDataSource(cfg, client)
			return
		}
		if _kinesisanalyticsv2DeleteApplicationSnapshot {
			kinesisanalyticsv2_DeleteApplicationSnapshot(cfg, client)
			return
		}
		if _kinesisanalyticsv2DeleteApplicationVpcConfiguration {
			kinesisanalyticsv2_DeleteApplicationVpcConfiguration(cfg, client)
			return
		}
		if _kinesisanalyticsv2DescribeApplication {
			kinesisanalyticsv2_DescribeApplication(cfg, client)
			return
		}
		if _kinesisanalyticsv2DescribeApplicationOperation {
			kinesisanalyticsv2_DescribeApplicationOperation(cfg, client)
			return
		}
		if _kinesisanalyticsv2DescribeApplicationSnapshot {
			kinesisanalyticsv2_DescribeApplicationSnapshot(cfg, client)
			return
		}
		if _kinesisanalyticsv2DescribeApplicationVersion {
			kinesisanalyticsv2_DescribeApplicationVersion(cfg, client)
			return
		}
		if _kinesisanalyticsv2DiscoverInputSchema {
			kinesisanalyticsv2_DiscoverInputSchema(cfg, client)
			return
		}
		if _kinesisanalyticsv2ListApplicationOperations {
			kinesisanalyticsv2_ListApplicationOperations(cfg, client)
			return
		}
		if _kinesisanalyticsv2ListApplicationSnapshots {
			kinesisanalyticsv2_ListApplicationSnapshots(cfg, client)
			return
		}
		if _kinesisanalyticsv2ListApplicationVersions {
			kinesisanalyticsv2_ListApplicationVersions(cfg, client)
			return
		}
		if _kinesisanalyticsv2ListApplications {
			kinesisanalyticsv2_ListApplications(cfg, client)
			return
		}
		if _kinesisanalyticsv2ListTagsForResource {
			kinesisanalyticsv2_ListTagsForResource(cfg, client)
			return
		}
		if _kinesisanalyticsv2RollbackApplication {
			kinesisanalyticsv2_RollbackApplication(cfg, client)
			return
		}
		if _kinesisanalyticsv2StartApplication {
			kinesisanalyticsv2_StartApplication(cfg, client)
			return
		}
		if _kinesisanalyticsv2StopApplication {
			kinesisanalyticsv2_StopApplication(cfg, client)
			return
		}
		if _kinesisanalyticsv2TagResource {
			kinesisanalyticsv2_TagResource(cfg, client)
			return
		}
		if _kinesisanalyticsv2UntagResource {
			kinesisanalyticsv2_UntagResource(cfg, client)
			return
		}
		if _kinesisanalyticsv2UpdateApplication {
			kinesisanalyticsv2_UpdateApplication(cfg, client)
			return
		}
		if _kinesisanalyticsv2UpdateApplicationMaintenanceConfiguration {
			kinesisanalyticsv2_UpdateApplicationMaintenanceConfiguration(cfg, client)
			return
		}

	},
}

var (
	_kinesisanalyticsv2AddApplicationCloudWatchLoggingOption         bool
	_kinesisanalyticsv2AddApplicationInput                           bool
	_kinesisanalyticsv2AddApplicationInputProcessingConfiguration    bool
	_kinesisanalyticsv2AddApplicationOutput                          bool
	_kinesisanalyticsv2AddApplicationReferenceDataSource             bool
	_kinesisanalyticsv2AddApplicationVpcConfiguration                bool
	_kinesisanalyticsv2CreateApplication                             bool
	_kinesisanalyticsv2CreateApplicationPresignedUrl                 bool
	_kinesisanalyticsv2CreateApplicationSnapshot                     bool
	_kinesisanalyticsv2DeleteApplication                             bool
	_kinesisanalyticsv2DeleteApplicationCloudWatchLoggingOption      bool
	_kinesisanalyticsv2DeleteApplicationInputProcessingConfiguration bool
	_kinesisanalyticsv2DeleteApplicationOutput                       bool
	_kinesisanalyticsv2DeleteApplicationReferenceDataSource          bool
	_kinesisanalyticsv2DeleteApplicationSnapshot                     bool
	_kinesisanalyticsv2DeleteApplicationVpcConfiguration             bool
	_kinesisanalyticsv2DescribeApplication                           bool
	_kinesisanalyticsv2DescribeApplicationOperation                  bool
	_kinesisanalyticsv2DescribeApplicationSnapshot                   bool
	_kinesisanalyticsv2DescribeApplicationVersion                    bool
	_kinesisanalyticsv2DiscoverInputSchema                           bool
	_kinesisanalyticsv2ListApplicationOperations                     bool
	_kinesisanalyticsv2ListApplicationSnapshots                      bool
	_kinesisanalyticsv2ListApplicationVersions                       bool
	_kinesisanalyticsv2ListApplications                              bool
	_kinesisanalyticsv2ListTagsForResource                           bool
	_kinesisanalyticsv2RollbackApplication                           bool
	_kinesisanalyticsv2StartApplication                              bool
	_kinesisanalyticsv2StopApplication                               bool
	_kinesisanalyticsv2TagResource                                   bool
	_kinesisanalyticsv2UntagResource                                 bool
	_kinesisanalyticsv2UpdateApplication                             bool
	_kinesisanalyticsv2UpdateApplicationMaintenanceConfiguration     bool

	_kinesisanalyticsv2ApplicationConfiguration                  string
	_kinesisanalyticsv2ApplicationConfigurationUpdate            string
	_kinesisanalyticsv2ApplicationDescription                    string
	_kinesisanalyticsv2ApplicationMaintenanceConfigurationUpdate string
	_kinesisanalyticsv2ApplicationMode                           string
	_kinesisanalyticsv2ApplicationName                           string
	_kinesisanalyticsv2ApplicationVersionId                      string
	_kinesisanalyticsv2CloudWatchLoggingOption                   string
	_kinesisanalyticsv2CloudWatchLoggingOptionId                 string
	_kinesisanalyticsv2CloudWatchLoggingOptionUpdates            string
	_kinesisanalyticsv2CloudWatchLoggingOptions                  string
	_kinesisanalyticsv2ConditionalToken                          string
	_kinesisanalyticsv2CreateTimestamp                           string
	_kinesisanalyticsv2CurrentApplicationVersionId               string
	_kinesisanalyticsv2Force                                     string
	_kinesisanalyticsv2IncludeAdditionalDetails                  string
	_kinesisanalyticsv2Input                                     string
	_kinesisanalyticsv2InputId                                   string
	_kinesisanalyticsv2InputProcessingConfiguration              string
	_kinesisanalyticsv2InputStartingPositionConfiguration        string
	_kinesisanalyticsv2Limit                                     string
	_kinesisanalyticsv2NextToken                                 string
	_kinesisanalyticsv2Operation                                 string
	_kinesisanalyticsv2OperationId                               string
	_kinesisanalyticsv2OperationStatus                           string
	_kinesisanalyticsv2OutputId                                  string
	_kinesisanalyticsv2ReferenceDataSource                       string
	_kinesisanalyticsv2ReferenceId                               string
	_kinesisanalyticsv2ResourceARN                               string
	_kinesisanalyticsv2RunConfiguration                          string
	_kinesisanalyticsv2RunConfigurationUpdate                    string
	_kinesisanalyticsv2RuntimeEnvironment                        string
	_kinesisanalyticsv2RuntimeEnvironmentUpdate                  string
	_kinesisanalyticsv2S3Configuration                           string
	_kinesisanalyticsv2ServiceExecutionRole                      string
	_kinesisanalyticsv2ServiceExecutionRoleUpdate                string
	_kinesisanalyticsv2SessionExpirationDurationInSeconds        string
	_kinesisanalyticsv2SnapshotCreationTimestamp                 string
	_kinesisanalyticsv2SnapshotName                              string
	_kinesisanalyticsv2TagKeys                                   []string
	_kinesisanalyticsv2Tags                                      string
	_kinesisanalyticsv2UrlType                                   string
	_kinesisanalyticsv2VpcConfiguration                          string
	_kinesisanalyticsv2VpcConfigurationId                        string
)

// Adds an Amazon CloudWatch log stream to monitor application configuration
// errors.
func kinesisanalyticsv2_AddApplicationCloudWatchLoggingOption(cfg aws.Config, client *kinesisanalyticsv2.Client) {
	input := &kinesisanalyticsv2.AddApplicationCloudWatchLoggingOptionInput{
		// ApplicationName: *string, // Required
		// CloudWatchLoggingOption: *types.CloudWatchLoggingOption, // Required
	}

	if len(_kinesisanalyticsv2ApplicationName) > 0 {
		input.ApplicationName = aws.String(_kinesisanalyticsv2ApplicationName)
	}
	if len(_kinesisanalyticsv2CloudWatchLoggingOption) > 0 {
		if err := assignInputField(input, "CloudWatchLoggingOption", _kinesisanalyticsv2CloudWatchLoggingOption); err != nil {
			log.Errorf("invalid --cloud-watch-logging-option: %s", err.Error())
			return
		}
	}
	if len(_kinesisanalyticsv2ConditionalToken) > 0 {
		input.ConditionalToken = aws.String(_kinesisanalyticsv2ConditionalToken)
	}
	if len(_kinesisanalyticsv2CurrentApplicationVersionId) > 0 {
		if err := assignInputField(input, "CurrentApplicationVersionId", _kinesisanalyticsv2CurrentApplicationVersionId); err != nil {
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

// Adds a streaming source to your SQL-based Kinesis Data Analytics application.
// You can add a streaming source when you create an application, or you can use
// this operation to add a streaming source after you create an application. For
// more information, see CreateApplication.
//
// Any configuration update, including adding a streaming source using this
// operation, results in a new version of the application. You can use the DescribeApplication
// operation to find the current application version.
func kinesisanalyticsv2_AddApplicationInput(cfg aws.Config, client *kinesisanalyticsv2.Client) {
	input := &kinesisanalyticsv2.AddApplicationInputInput{
		// ApplicationName: *string, // Required
		// CurrentApplicationVersionId: *int64, // Required
		// Input: *types.Input, // Required
	}

	if len(_kinesisanalyticsv2ApplicationName) > 0 {
		input.ApplicationName = aws.String(_kinesisanalyticsv2ApplicationName)
	}
	if len(_kinesisanalyticsv2CurrentApplicationVersionId) > 0 {
		if err := assignInputField(input, "CurrentApplicationVersionId", _kinesisanalyticsv2CurrentApplicationVersionId); err != nil {
			log.Errorf("invalid --current-application-version-id: %s", err.Error())
			return
		}
	}
	if len(_kinesisanalyticsv2Input) > 0 {
		if err := assignInputField(input, "Input", _kinesisanalyticsv2Input); err != nil {
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

// Adds an InputProcessingConfiguration to a SQL-based Kinesis Data Analytics application. An input processor
// pre-processes records on the input stream before the application's SQL code
// executes. Currently, the only input processor available is [Amazon Lambda].
//
// [Amazon Lambda]: https://docs.aws.amazon.com/lambda/
func kinesisanalyticsv2_AddApplicationInputProcessingConfiguration(cfg aws.Config, client *kinesisanalyticsv2.Client) {
	input := &kinesisanalyticsv2.AddApplicationInputProcessingConfigurationInput{
		// ApplicationName: *string, // Required
		// CurrentApplicationVersionId: *int64, // Required
		// InputId: *string, // Required
		// InputProcessingConfiguration: *types.InputProcessingConfiguration, // Required
	}

	if len(_kinesisanalyticsv2ApplicationName) > 0 {
		input.ApplicationName = aws.String(_kinesisanalyticsv2ApplicationName)
	}
	if len(_kinesisanalyticsv2CurrentApplicationVersionId) > 0 {
		if err := assignInputField(input, "CurrentApplicationVersionId", _kinesisanalyticsv2CurrentApplicationVersionId); err != nil {
			log.Errorf("invalid --current-application-version-id: %s", err.Error())
			return
		}
	}
	if len(_kinesisanalyticsv2InputId) > 0 {
		input.InputId = aws.String(_kinesisanalyticsv2InputId)
	}
	if len(_kinesisanalyticsv2InputProcessingConfiguration) > 0 {
		if err := assignInputField(input, "InputProcessingConfiguration", _kinesisanalyticsv2InputProcessingConfiguration); err != nil {
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

// Adds an external destination to your SQL-based Kinesis Data Analytics
// application.
//
// If you want Kinesis Data Analytics to deliver data from an in-application
// stream within your application to an external destination (such as an Kinesis
// data stream, a Kinesis Data Firehose delivery stream, or an Amazon Lambda
// function), you add the relevant configuration to your application using this
// operation. You can configure one or more outputs for your application. Each
// output configuration maps an in-application stream and an external destination.
//
// You can use one of the output configurations to deliver data from your
// in-application error stream to an external destination so that you can analyze
// the errors.
//
// Any configuration update, including adding a streaming source using this
// operation, results in a new version of the application. You can use the DescribeApplication
// operation to find the current application version.
func kinesisanalyticsv2_AddApplicationOutput(cfg aws.Config, client *kinesisanalyticsv2.Client) {
	input := &kinesisanalyticsv2.AddApplicationOutputInput{
		// ApplicationName: *string, // Required
		// CurrentApplicationVersionId: *int64, // Required
		// Output: *types.Output, // Required
	}

	if len(_kinesisanalyticsv2ApplicationName) > 0 {
		input.ApplicationName = aws.String(_kinesisanalyticsv2ApplicationName)
	}
	if len(_kinesisanalyticsv2CurrentApplicationVersionId) > 0 {
		if err := assignInputField(input, "CurrentApplicationVersionId", _kinesisanalyticsv2CurrentApplicationVersionId); err != nil {
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

// Adds a reference data source to an existing SQL-based Kinesis Data Analytics
// application.
//
// Kinesis Data Analytics reads reference data (that is, an Amazon S3 object) and
// creates an in-application table within your application. In the request, you
// provide the source (S3 bucket name and object key name), name of the
// in-application table to create, and the necessary mapping information that
// describes how data in an Amazon S3 object maps to columns in the resulting
// in-application table.
func kinesisanalyticsv2_AddApplicationReferenceDataSource(cfg aws.Config, client *kinesisanalyticsv2.Client) {
	input := &kinesisanalyticsv2.AddApplicationReferenceDataSourceInput{
		// ApplicationName: *string, // Required
		// CurrentApplicationVersionId: *int64, // Required
		// ReferenceDataSource: *types.ReferenceDataSource, // Required
	}

	if len(_kinesisanalyticsv2ApplicationName) > 0 {
		input.ApplicationName = aws.String(_kinesisanalyticsv2ApplicationName)
	}
	if len(_kinesisanalyticsv2CurrentApplicationVersionId) > 0 {
		if err := assignInputField(input, "CurrentApplicationVersionId", _kinesisanalyticsv2CurrentApplicationVersionId); err != nil {
			log.Errorf("invalid --current-application-version-id: %s", err.Error())
			return
		}
	}
	if len(_kinesisanalyticsv2ReferenceDataSource) > 0 {
		if err := assignInputField(input, "ReferenceDataSource", _kinesisanalyticsv2ReferenceDataSource); err != nil {
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

// Adds a Virtual Private Cloud (VPC) configuration to the application.
// Applications can use VPCs to store and access resources securely.
//
// Note the following about VPC configurations for Managed Service for Apache
// Flink applications:
//
// - VPC configurations are not supported for SQL applications.
//
// - When a VPC is added to a Managed Service for Apache Flink application, the
// application can no longer be accessed from the Internet directly. To enable
// Internet access to the application, add an Internet gateway to your VPC.
func kinesisanalyticsv2_AddApplicationVpcConfiguration(cfg aws.Config, client *kinesisanalyticsv2.Client) {
	input := &kinesisanalyticsv2.AddApplicationVpcConfigurationInput{
		// ApplicationName: *string, // Required
		// VpcConfiguration: *types.VpcConfiguration, // Required
	}

	if len(_kinesisanalyticsv2ApplicationName) > 0 {
		input.ApplicationName = aws.String(_kinesisanalyticsv2ApplicationName)
	}
	if len(_kinesisanalyticsv2VpcConfiguration) > 0 {
		if err := assignInputField(input, "VpcConfiguration", _kinesisanalyticsv2VpcConfiguration); err != nil {
			log.Errorf("invalid --vpc-configuration: %s", err.Error())
			return
		}
	}
	if len(_kinesisanalyticsv2ConditionalToken) > 0 {
		input.ConditionalToken = aws.String(_kinesisanalyticsv2ConditionalToken)
	}
	if len(_kinesisanalyticsv2CurrentApplicationVersionId) > 0 {
		if err := assignInputField(input, "CurrentApplicationVersionId", _kinesisanalyticsv2CurrentApplicationVersionId); err != nil {
			log.Errorf("invalid --current-application-version-id: %s", err.Error())
			return
		}
	}

	if resp, err := client.AddApplicationVpcConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a Managed Service for Apache Flink application. For information about
// creating a Managed Service for Apache Flink application, see [Creating an Application].
//
// [Creating an Application]: https://docs.aws.amazon.com/kinesisanalytics/latest/java/getting-started.html
func kinesisanalyticsv2_CreateApplication(cfg aws.Config, client *kinesisanalyticsv2.Client) {
	input := &kinesisanalyticsv2.CreateApplicationInput{
		// ApplicationName: *string, // Required
		// RuntimeEnvironment: types.RuntimeEnvironment, // Required
		// ServiceExecutionRole: *string, // Required
	}

	if len(_kinesisanalyticsv2ApplicationName) > 0 {
		input.ApplicationName = aws.String(_kinesisanalyticsv2ApplicationName)
	}
	if len(_kinesisanalyticsv2RuntimeEnvironment) > 0 {
		if err := assignInputField(input, "RuntimeEnvironment", _kinesisanalyticsv2RuntimeEnvironment); err != nil {
			log.Errorf("invalid --runtime-environment: %s", err.Error())
			return
		}
	}
	if len(_kinesisanalyticsv2ServiceExecutionRole) > 0 {
		input.ServiceExecutionRole = aws.String(_kinesisanalyticsv2ServiceExecutionRole)
	}
	if len(_kinesisanalyticsv2ApplicationConfiguration) > 0 {
		if err := assignInputField(input, "ApplicationConfiguration", _kinesisanalyticsv2ApplicationConfiguration); err != nil {
			log.Errorf("invalid --application-configuration: %s", err.Error())
			return
		}
	}
	if len(_kinesisanalyticsv2ApplicationDescription) > 0 {
		input.ApplicationDescription = aws.String(_kinesisanalyticsv2ApplicationDescription)
	}
	if len(_kinesisanalyticsv2ApplicationMode) > 0 {
		if err := assignInputField(input, "ApplicationMode", _kinesisanalyticsv2ApplicationMode); err != nil {
			log.Errorf("invalid --application-mode: %s", err.Error())
			return
		}
	}
	if len(_kinesisanalyticsv2CloudWatchLoggingOptions) > 0 {
		if err := assignInputField(input, "CloudWatchLoggingOptions", _kinesisanalyticsv2CloudWatchLoggingOptions); err != nil {
			log.Errorf("invalid --cloud-watch-logging-options: %s", err.Error())
			return
		}
	}
	if len(_kinesisanalyticsv2Tags) > 0 {
		if err := assignInputField(input, "Tags", _kinesisanalyticsv2Tags); err != nil {
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

// Creates and returns a URL that you can use to connect to an application's
// extension.
//
// The IAM role or user used to call this API defines the permissions to access
// the extension. After the presigned URL is created, no additional permission is
// required to access this URL. IAM authorization policies for this API are also
// enforced for every HTTP request that attempts to connect to the extension.
//
// You control the amount of time that the URL will be valid using the
// SessionExpirationDurationInSeconds parameter. If you do not provide this
// parameter, the returned URL is valid for twelve hours.
//
// The URL that you get from a call to CreateApplicationPresignedUrl must be used
// within 3 minutes to be valid. If you first try to use the URL after the 3-minute
// limit expires, the service returns an HTTP 403 Forbidden error.
func kinesisanalyticsv2_CreateApplicationPresignedUrl(cfg aws.Config, client *kinesisanalyticsv2.Client) {
	input := &kinesisanalyticsv2.CreateApplicationPresignedUrlInput{
		// ApplicationName: *string, // Required
		// UrlType: types.UrlType, // Required
	}

	if len(_kinesisanalyticsv2ApplicationName) > 0 {
		input.ApplicationName = aws.String(_kinesisanalyticsv2ApplicationName)
	}
	if len(_kinesisanalyticsv2UrlType) > 0 {
		if err := assignInputField(input, "UrlType", _kinesisanalyticsv2UrlType); err != nil {
			log.Errorf("invalid --url-type: %s", err.Error())
			return
		}
	}
	if len(_kinesisanalyticsv2SessionExpirationDurationInSeconds) > 0 {
		if err := assignInputField(input, "SessionExpirationDurationInSeconds", _kinesisanalyticsv2SessionExpirationDurationInSeconds); err != nil {
			log.Errorf("invalid --session-expiration-duration-in-seconds: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateApplicationPresignedUrl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a snapshot of the application's state data.
func kinesisanalyticsv2_CreateApplicationSnapshot(cfg aws.Config, client *kinesisanalyticsv2.Client) {
	input := &kinesisanalyticsv2.CreateApplicationSnapshotInput{
		// ApplicationName: *string, // Required
		// SnapshotName: *string, // Required
	}

	if len(_kinesisanalyticsv2ApplicationName) > 0 {
		input.ApplicationName = aws.String(_kinesisanalyticsv2ApplicationName)
	}
	if len(_kinesisanalyticsv2SnapshotName) > 0 {
		input.SnapshotName = aws.String(_kinesisanalyticsv2SnapshotName)
	}

	if resp, err := client.CreateApplicationSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified application. Managed Service for Apache Flink halts
// application execution and deletes the application.
func kinesisanalyticsv2_DeleteApplication(cfg aws.Config, client *kinesisanalyticsv2.Client) {
	input := &kinesisanalyticsv2.DeleteApplicationInput{
		// ApplicationName: *string, // Required
		// CreateTimestamp: *time.Time, // Required
	}

	if len(_kinesisanalyticsv2ApplicationName) > 0 {
		input.ApplicationName = aws.String(_kinesisanalyticsv2ApplicationName)
	}
	if len(_kinesisanalyticsv2CreateTimestamp) > 0 {
		if err := assignInputField(input, "CreateTimestamp", _kinesisanalyticsv2CreateTimestamp); err != nil {
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

// Deletes an Amazon CloudWatch log stream from an SQL-based Kinesis Data
// Analytics application.
func kinesisanalyticsv2_DeleteApplicationCloudWatchLoggingOption(cfg aws.Config, client *kinesisanalyticsv2.Client) {
	input := &kinesisanalyticsv2.DeleteApplicationCloudWatchLoggingOptionInput{
		// ApplicationName: *string, // Required
		// CloudWatchLoggingOptionId: *string, // Required
	}

	if len(_kinesisanalyticsv2ApplicationName) > 0 {
		input.ApplicationName = aws.String(_kinesisanalyticsv2ApplicationName)
	}
	if len(_kinesisanalyticsv2CloudWatchLoggingOptionId) > 0 {
		input.CloudWatchLoggingOptionId = aws.String(_kinesisanalyticsv2CloudWatchLoggingOptionId)
	}
	if len(_kinesisanalyticsv2ConditionalToken) > 0 {
		input.ConditionalToken = aws.String(_kinesisanalyticsv2ConditionalToken)
	}
	if len(_kinesisanalyticsv2CurrentApplicationVersionId) > 0 {
		if err := assignInputField(input, "CurrentApplicationVersionId", _kinesisanalyticsv2CurrentApplicationVersionId); err != nil {
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

// Deletes an InputProcessingConfiguration from an input.
func kinesisanalyticsv2_DeleteApplicationInputProcessingConfiguration(cfg aws.Config, client *kinesisanalyticsv2.Client) {
	input := &kinesisanalyticsv2.DeleteApplicationInputProcessingConfigurationInput{
		// ApplicationName: *string, // Required
		// CurrentApplicationVersionId: *int64, // Required
		// InputId: *string, // Required
	}

	if len(_kinesisanalyticsv2ApplicationName) > 0 {
		input.ApplicationName = aws.String(_kinesisanalyticsv2ApplicationName)
	}
	if len(_kinesisanalyticsv2CurrentApplicationVersionId) > 0 {
		if err := assignInputField(input, "CurrentApplicationVersionId", _kinesisanalyticsv2CurrentApplicationVersionId); err != nil {
			log.Errorf("invalid --current-application-version-id: %s", err.Error())
			return
		}
	}
	if len(_kinesisanalyticsv2InputId) > 0 {
		input.InputId = aws.String(_kinesisanalyticsv2InputId)
	}

	if resp, err := client.DeleteApplicationInputProcessingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the output destination configuration from your SQL-based Kinesis Data
// Analytics application's configuration. Kinesis Data Analytics will no longer
// write data from the corresponding in-application stream to the external output
// destination.
func kinesisanalyticsv2_DeleteApplicationOutput(cfg aws.Config, client *kinesisanalyticsv2.Client) {
	input := &kinesisanalyticsv2.DeleteApplicationOutputInput{
		// ApplicationName: *string, // Required
		// CurrentApplicationVersionId: *int64, // Required
		// OutputId: *string, // Required
	}

	if len(_kinesisanalyticsv2ApplicationName) > 0 {
		input.ApplicationName = aws.String(_kinesisanalyticsv2ApplicationName)
	}
	if len(_kinesisanalyticsv2CurrentApplicationVersionId) > 0 {
		if err := assignInputField(input, "CurrentApplicationVersionId", _kinesisanalyticsv2CurrentApplicationVersionId); err != nil {
			log.Errorf("invalid --current-application-version-id: %s", err.Error())
			return
		}
	}
	if len(_kinesisanalyticsv2OutputId) > 0 {
		input.OutputId = aws.String(_kinesisanalyticsv2OutputId)
	}

	if resp, err := client.DeleteApplicationOutput(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a reference data source configuration from the specified SQL-based
// Kinesis Data Analytics application's configuration.
//
// If the application is running, Kinesis Data Analytics immediately removes the
// in-application table that you created using the AddApplicationReferenceDataSourceoperation.
func kinesisanalyticsv2_DeleteApplicationReferenceDataSource(cfg aws.Config, client *kinesisanalyticsv2.Client) {
	input := &kinesisanalyticsv2.DeleteApplicationReferenceDataSourceInput{
		// ApplicationName: *string, // Required
		// CurrentApplicationVersionId: *int64, // Required
		// ReferenceId: *string, // Required
	}

	if len(_kinesisanalyticsv2ApplicationName) > 0 {
		input.ApplicationName = aws.String(_kinesisanalyticsv2ApplicationName)
	}
	if len(_kinesisanalyticsv2CurrentApplicationVersionId) > 0 {
		if err := assignInputField(input, "CurrentApplicationVersionId", _kinesisanalyticsv2CurrentApplicationVersionId); err != nil {
			log.Errorf("invalid --current-application-version-id: %s", err.Error())
			return
		}
	}
	if len(_kinesisanalyticsv2ReferenceId) > 0 {
		input.ReferenceId = aws.String(_kinesisanalyticsv2ReferenceId)
	}

	if resp, err := client.DeleteApplicationReferenceDataSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a snapshot of application state.
func kinesisanalyticsv2_DeleteApplicationSnapshot(cfg aws.Config, client *kinesisanalyticsv2.Client) {
	input := &kinesisanalyticsv2.DeleteApplicationSnapshotInput{
		// ApplicationName: *string, // Required
		// SnapshotCreationTimestamp: *time.Time, // Required
		// SnapshotName: *string, // Required
	}

	if len(_kinesisanalyticsv2ApplicationName) > 0 {
		input.ApplicationName = aws.String(_kinesisanalyticsv2ApplicationName)
	}
	if len(_kinesisanalyticsv2SnapshotCreationTimestamp) > 0 {
		if err := assignInputField(input, "SnapshotCreationTimestamp", _kinesisanalyticsv2SnapshotCreationTimestamp); err != nil {
			log.Errorf("invalid --snapshot-creation-timestamp: %s", err.Error())
			return
		}
	}
	if len(_kinesisanalyticsv2SnapshotName) > 0 {
		input.SnapshotName = aws.String(_kinesisanalyticsv2SnapshotName)
	}

	if resp, err := client.DeleteApplicationSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a VPC configuration from a Managed Service for Apache Flink application.
func kinesisanalyticsv2_DeleteApplicationVpcConfiguration(cfg aws.Config, client *kinesisanalyticsv2.Client) {
	input := &kinesisanalyticsv2.DeleteApplicationVpcConfigurationInput{
		// ApplicationName: *string, // Required
		// VpcConfigurationId: *string, // Required
	}

	if len(_kinesisanalyticsv2ApplicationName) > 0 {
		input.ApplicationName = aws.String(_kinesisanalyticsv2ApplicationName)
	}
	if len(_kinesisanalyticsv2VpcConfigurationId) > 0 {
		input.VpcConfigurationId = aws.String(_kinesisanalyticsv2VpcConfigurationId)
	}
	if len(_kinesisanalyticsv2ConditionalToken) > 0 {
		input.ConditionalToken = aws.String(_kinesisanalyticsv2ConditionalToken)
	}
	if len(_kinesisanalyticsv2CurrentApplicationVersionId) > 0 {
		if err := assignInputField(input, "CurrentApplicationVersionId", _kinesisanalyticsv2CurrentApplicationVersionId); err != nil {
			log.Errorf("invalid --current-application-version-id: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteApplicationVpcConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a specific Managed Service for Apache Flink
// application.
//
// If you want to retrieve a list of all applications in your account, use the ListApplications
// operation.
func kinesisanalyticsv2_DescribeApplication(cfg aws.Config, client *kinesisanalyticsv2.Client) {
	input := &kinesisanalyticsv2.DescribeApplicationInput{
		// ApplicationName: *string, // Required
	}

	if len(_kinesisanalyticsv2ApplicationName) > 0 {
		input.ApplicationName = aws.String(_kinesisanalyticsv2ApplicationName)
	}
	if len(_kinesisanalyticsv2IncludeAdditionalDetails) > 0 {
		if err := assignInputField(input, "IncludeAdditionalDetails", _kinesisanalyticsv2IncludeAdditionalDetails); err != nil {
			log.Errorf("invalid --include-additional-details: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides a detailed description of a specified application operation. To see a
// list of all the operations of an application, invoke the ListApplicationOperationsoperation.
//
// This operation is supported only for Managed Service for Apache Flink.
func kinesisanalyticsv2_DescribeApplicationOperation(cfg aws.Config, client *kinesisanalyticsv2.Client) {
	input := &kinesisanalyticsv2.DescribeApplicationOperationInput{
		// ApplicationName: *string, // Required
		// OperationId: *string, // Required
	}

	if len(_kinesisanalyticsv2ApplicationName) > 0 {
		input.ApplicationName = aws.String(_kinesisanalyticsv2ApplicationName)
	}
	if len(_kinesisanalyticsv2OperationId) > 0 {
		input.OperationId = aws.String(_kinesisanalyticsv2OperationId)
	}

	if resp, err := client.DescribeApplicationOperation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a snapshot of application state data.
func kinesisanalyticsv2_DescribeApplicationSnapshot(cfg aws.Config, client *kinesisanalyticsv2.Client) {
	input := &kinesisanalyticsv2.DescribeApplicationSnapshotInput{
		// ApplicationName: *string, // Required
		// SnapshotName: *string, // Required
	}

	if len(_kinesisanalyticsv2ApplicationName) > 0 {
		input.ApplicationName = aws.String(_kinesisanalyticsv2ApplicationName)
	}
	if len(_kinesisanalyticsv2SnapshotName) > 0 {
		input.SnapshotName = aws.String(_kinesisanalyticsv2SnapshotName)
	}

	if resp, err := client.DescribeApplicationSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides a detailed description of a specified version of the application. To
// see a list of all the versions of an application, invoke the ListApplicationVersionsoperation.
//
// This operation is supported only for Managed Service for Apache Flink.
func kinesisanalyticsv2_DescribeApplicationVersion(cfg aws.Config, client *kinesisanalyticsv2.Client) {
	input := &kinesisanalyticsv2.DescribeApplicationVersionInput{
		// ApplicationName: *string, // Required
		// ApplicationVersionId: *int64, // Required
	}

	if len(_kinesisanalyticsv2ApplicationName) > 0 {
		input.ApplicationName = aws.String(_kinesisanalyticsv2ApplicationName)
	}
	if len(_kinesisanalyticsv2ApplicationVersionId) > 0 {
		if err := assignInputField(input, "ApplicationVersionId", _kinesisanalyticsv2ApplicationVersionId); err != nil {
			log.Errorf("invalid --application-version-id: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeApplicationVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Infers a schema for a SQL-based Kinesis Data Analytics application by
// evaluating sample records on the specified streaming source (Kinesis data stream
// or Kinesis Data Firehose delivery stream) or Amazon S3 object. In the response,
// the operation returns the inferred schema and also the sample records that the
// operation used to infer the schema.
//
// You can use the inferred schema when configuring a streaming source for your
// application. When you create an application using the Kinesis Data Analytics
// console, the console uses this operation to infer a schema and show it in the
// console user interface.
func kinesisanalyticsv2_DiscoverInputSchema(cfg aws.Config, client *kinesisanalyticsv2.Client) {
	input := &kinesisanalyticsv2.DiscoverInputSchemaInput{
		// ServiceExecutionRole: *string, // Required
	}

	if len(_kinesisanalyticsv2ServiceExecutionRole) > 0 {
		input.ServiceExecutionRole = aws.String(_kinesisanalyticsv2ServiceExecutionRole)
	}
	if len(_kinesisanalyticsv2InputProcessingConfiguration) > 0 {
		if err := assignInputField(input, "InputProcessingConfiguration", _kinesisanalyticsv2InputProcessingConfiguration); err != nil {
			log.Errorf("invalid --input-processing-configuration: %s", err.Error())
			return
		}
	}
	if len(_kinesisanalyticsv2InputStartingPositionConfiguration) > 0 {
		if err := assignInputField(input, "InputStartingPositionConfiguration", _kinesisanalyticsv2InputStartingPositionConfiguration); err != nil {
			log.Errorf("invalid --input-starting-position-configuration: %s", err.Error())
			return
		}
	}
	if len(_kinesisanalyticsv2ResourceARN) > 0 {
		input.ResourceARN = aws.String(_kinesisanalyticsv2ResourceARN)
	}
	if len(_kinesisanalyticsv2S3Configuration) > 0 {
		if err := assignInputField(input, "S3Configuration", _kinesisanalyticsv2S3Configuration); err != nil {
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

// Lists all the operations performed for the specified application such as
// UpdateApplication, StartApplication etc. The response also includes a summary of
// the operation.
//
// To get the complete description of a specific operation, invoke the DescribeApplicationOperation operation.
//
// This operation is supported only for Managed Service for Apache Flink.
func kinesisanalyticsv2_ListApplicationOperations(cfg aws.Config, client *kinesisanalyticsv2.Client) {
	input := &kinesisanalyticsv2.ListApplicationOperationsInput{
		// ApplicationName: *string, // Required
	}

	if len(_kinesisanalyticsv2ApplicationName) > 0 {
		input.ApplicationName = aws.String(_kinesisanalyticsv2ApplicationName)
	}
	if len(_kinesisanalyticsv2Limit) > 0 {
		if err := assignInputField(input, "Limit", _kinesisanalyticsv2Limit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_kinesisanalyticsv2NextToken) > 0 {
		input.NextToken = aws.String(_kinesisanalyticsv2NextToken)
	}
	if len(_kinesisanalyticsv2Operation) > 0 {
		input.Operation = aws.String(_kinesisanalyticsv2Operation)
	}
	if len(_kinesisanalyticsv2OperationStatus) > 0 {
		if err := assignInputField(input, "OperationStatus", _kinesisanalyticsv2OperationStatus); err != nil {
			log.Errorf("invalid --operation-status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListApplicationOperations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*kinesisanalyticsv2.ListApplicationOperationsOutput
	p := kinesisanalyticsv2.NewListApplicationOperationsPaginator(client, input)
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

// Lists information about the current application snapshots.
func kinesisanalyticsv2_ListApplicationSnapshots(cfg aws.Config, client *kinesisanalyticsv2.Client) {
	input := &kinesisanalyticsv2.ListApplicationSnapshotsInput{
		// ApplicationName: *string, // Required
	}

	if len(_kinesisanalyticsv2ApplicationName) > 0 {
		input.ApplicationName = aws.String(_kinesisanalyticsv2ApplicationName)
	}
	if len(_kinesisanalyticsv2Limit) > 0 {
		if err := assignInputField(input, "Limit", _kinesisanalyticsv2Limit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_kinesisanalyticsv2NextToken) > 0 {
		input.NextToken = aws.String(_kinesisanalyticsv2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListApplicationSnapshots(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*kinesisanalyticsv2.ListApplicationSnapshotsOutput
	p := kinesisanalyticsv2.NewListApplicationSnapshotsPaginator(client, input)
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

// Lists all the versions for the specified application, including versions that
// were rolled back. The response also includes a summary of the configuration
// associated with each version.
//
// To get the complete description of a specific application version, invoke the DescribeApplicationVersion
// operation.
//
// This operation is supported only for Managed Service for Apache Flink.
func kinesisanalyticsv2_ListApplicationVersions(cfg aws.Config, client *kinesisanalyticsv2.Client) {
	input := &kinesisanalyticsv2.ListApplicationVersionsInput{
		// ApplicationName: *string, // Required
	}

	if len(_kinesisanalyticsv2ApplicationName) > 0 {
		input.ApplicationName = aws.String(_kinesisanalyticsv2ApplicationName)
	}
	if len(_kinesisanalyticsv2Limit) > 0 {
		if err := assignInputField(input, "Limit", _kinesisanalyticsv2Limit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_kinesisanalyticsv2NextToken) > 0 {
		input.NextToken = aws.String(_kinesisanalyticsv2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListApplicationVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*kinesisanalyticsv2.ListApplicationVersionsOutput
	p := kinesisanalyticsv2.NewListApplicationVersionsPaginator(client, input)
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

// Returns a list of Managed Service for Apache Flink applications in your
// account. For each application, the response includes the application name,
// Amazon Resource Name (ARN), and status.
//
// If you want detailed information about a specific application, use DescribeApplication.
func kinesisanalyticsv2_ListApplications(cfg aws.Config, client *kinesisanalyticsv2.Client) {
	input := &kinesisanalyticsv2.ListApplicationsInput{}

	if len(_kinesisanalyticsv2Limit) > 0 {
		if err := assignInputField(input, "Limit", _kinesisanalyticsv2Limit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_kinesisanalyticsv2NextToken) > 0 {
		input.NextToken = aws.String(_kinesisanalyticsv2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListApplications(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*kinesisanalyticsv2.ListApplicationsOutput
	p := kinesisanalyticsv2.NewListApplicationsPaginator(client, input)
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

// Retrieves the list of key-value tags assigned to the application. For more
// information, see [Using Tagging].
//
// [Using Tagging]: https://docs.aws.amazon.com/kinesisanalytics/latest/java/how-tagging.html
func kinesisanalyticsv2_ListTagsForResource(cfg aws.Config, client *kinesisanalyticsv2.Client) {
	input := &kinesisanalyticsv2.ListTagsForResourceInput{
		// ResourceARN: *string, // Required
	}

	if len(_kinesisanalyticsv2ResourceARN) > 0 {
		input.ResourceARN = aws.String(_kinesisanalyticsv2ResourceARN)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Reverts the application to the previous running version. You can roll back an
// application if you suspect it is stuck in a transient status or in the running
// status.
//
// You can roll back an application only if it is in the UPDATING , AUTOSCALING ,
// or RUNNING statuses.
//
// When you rollback an application, it loads state data from the last successful
// snapshot. If the application has no snapshots, Managed Service for Apache Flink
// rejects the rollback request.
func kinesisanalyticsv2_RollbackApplication(cfg aws.Config, client *kinesisanalyticsv2.Client) {
	input := &kinesisanalyticsv2.RollbackApplicationInput{
		// ApplicationName: *string, // Required
		// CurrentApplicationVersionId: *int64, // Required
	}

	if len(_kinesisanalyticsv2ApplicationName) > 0 {
		input.ApplicationName = aws.String(_kinesisanalyticsv2ApplicationName)
	}
	if len(_kinesisanalyticsv2CurrentApplicationVersionId) > 0 {
		if err := assignInputField(input, "CurrentApplicationVersionId", _kinesisanalyticsv2CurrentApplicationVersionId); err != nil {
			log.Errorf("invalid --current-application-version-id: %s", err.Error())
			return
		}
	}

	if resp, err := client.RollbackApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts the specified Managed Service for Apache Flink application. After
// creating an application, you must exclusively call this operation to start your
// application.
func kinesisanalyticsv2_StartApplication(cfg aws.Config, client *kinesisanalyticsv2.Client) {
	input := &kinesisanalyticsv2.StartApplicationInput{
		// ApplicationName: *string, // Required
	}

	if len(_kinesisanalyticsv2ApplicationName) > 0 {
		input.ApplicationName = aws.String(_kinesisanalyticsv2ApplicationName)
	}
	if len(_kinesisanalyticsv2RunConfiguration) > 0 {
		if err := assignInputField(input, "RunConfiguration", _kinesisanalyticsv2RunConfiguration); err != nil {
			log.Errorf("invalid --run-configuration: %s", err.Error())
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

// Stops the application from processing data. You can stop an application only if
// it is in the running status, unless you set the Force parameter to true .
//
// You can use the DescribeApplication operation to find the application status.
//
// Managed Service for Apache Flink takes a snapshot when the application is
// stopped, unless Force is set to true .
func kinesisanalyticsv2_StopApplication(cfg aws.Config, client *kinesisanalyticsv2.Client) {
	input := &kinesisanalyticsv2.StopApplicationInput{
		// ApplicationName: *string, // Required
	}

	if len(_kinesisanalyticsv2ApplicationName) > 0 {
		input.ApplicationName = aws.String(_kinesisanalyticsv2ApplicationName)
	}
	if len(_kinesisanalyticsv2Force) > 0 {
		if err := assignInputField(input, "Force", _kinesisanalyticsv2Force); err != nil {
			log.Errorf("invalid --force: %s", err.Error())
			return
		}
	}

	if resp, err := client.StopApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds one or more key-value tags to a Managed Service for Apache Flink
// application. Note that the maximum number of application tags includes system
// tags. The maximum number of user-defined application tags is 50. For more
// information, see [Using Tagging].
//
// [Using Tagging]: https://docs.aws.amazon.com/kinesisanalytics/latest/java/how-tagging.html
func kinesisanalyticsv2_TagResource(cfg aws.Config, client *kinesisanalyticsv2.Client) {
	input := &kinesisanalyticsv2.TagResourceInput{
		// ResourceARN: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_kinesisanalyticsv2ResourceARN) > 0 {
		input.ResourceARN = aws.String(_kinesisanalyticsv2ResourceARN)
	}
	if len(_kinesisanalyticsv2Tags) > 0 {
		if err := assignInputField(input, "Tags", _kinesisanalyticsv2Tags); err != nil {
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

// Removes one or more tags from a Managed Service for Apache Flink application.
// For more information, see [Using Tagging].
//
// [Using Tagging]: https://docs.aws.amazon.com/kinesisanalytics/latest/java/how-tagging.html
func kinesisanalyticsv2_UntagResource(cfg aws.Config, client *kinesisanalyticsv2.Client) {
	input := &kinesisanalyticsv2.UntagResourceInput{
		// ResourceARN: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_kinesisanalyticsv2ResourceARN) > 0 {
		input.ResourceARN = aws.String(_kinesisanalyticsv2ResourceARN)
	}
	if len(_kinesisanalyticsv2TagKeys) > 0 {
		input.TagKeys = append([]string(nil), _kinesisanalyticsv2TagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing Managed Service for Apache Flink application. Using this
// operation, you can update application code, input configuration, and output
// configuration.
//
// Managed Service for Apache Flink updates the ApplicationVersionId each time you
// update your application.
func kinesisanalyticsv2_UpdateApplication(cfg aws.Config, client *kinesisanalyticsv2.Client) {
	input := &kinesisanalyticsv2.UpdateApplicationInput{
		// ApplicationName: *string, // Required
	}

	if len(_kinesisanalyticsv2ApplicationName) > 0 {
		input.ApplicationName = aws.String(_kinesisanalyticsv2ApplicationName)
	}
	if len(_kinesisanalyticsv2ApplicationConfigurationUpdate) > 0 {
		if err := assignInputField(input, "ApplicationConfigurationUpdate", _kinesisanalyticsv2ApplicationConfigurationUpdate); err != nil {
			log.Errorf("invalid --application-configuration-update: %s", err.Error())
			return
		}
	}
	if len(_kinesisanalyticsv2CloudWatchLoggingOptionUpdates) > 0 {
		if err := assignInputField(input, "CloudWatchLoggingOptionUpdates", _kinesisanalyticsv2CloudWatchLoggingOptionUpdates); err != nil {
			log.Errorf("invalid --cloud-watch-logging-option-updates: %s", err.Error())
			return
		}
	}
	if len(_kinesisanalyticsv2ConditionalToken) > 0 {
		input.ConditionalToken = aws.String(_kinesisanalyticsv2ConditionalToken)
	}
	if len(_kinesisanalyticsv2CurrentApplicationVersionId) > 0 {
		if err := assignInputField(input, "CurrentApplicationVersionId", _kinesisanalyticsv2CurrentApplicationVersionId); err != nil {
			log.Errorf("invalid --current-application-version-id: %s", err.Error())
			return
		}
	}
	if len(_kinesisanalyticsv2RunConfigurationUpdate) > 0 {
		if err := assignInputField(input, "RunConfigurationUpdate", _kinesisanalyticsv2RunConfigurationUpdate); err != nil {
			log.Errorf("invalid --run-configuration-update: %s", err.Error())
			return
		}
	}
	if len(_kinesisanalyticsv2RuntimeEnvironmentUpdate) > 0 {
		if err := assignInputField(input, "RuntimeEnvironmentUpdate", _kinesisanalyticsv2RuntimeEnvironmentUpdate); err != nil {
			log.Errorf("invalid --runtime-environment-update: %s", err.Error())
			return
		}
	}
	if len(_kinesisanalyticsv2ServiceExecutionRoleUpdate) > 0 {
		input.ServiceExecutionRoleUpdate = aws.String(_kinesisanalyticsv2ServiceExecutionRoleUpdate)
	}

	if resp, err := client.UpdateApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the maintenance configuration of the Managed Service for Apache Flink
// application.
//
// You can invoke this operation on an application that is in one of the two
// following states: READY or RUNNING . If you invoke it when the application is in
// a state other than these two states, it throws a ResourceInUseException . The
// service makes use of the updated configuration the next time it schedules
// maintenance for the application. If you invoke this operation after the service
// schedules maintenance, the service will apply the configuration update the next
// time it schedules maintenance for the application. This means that you might not
// see the maintenance configuration update applied to the maintenance process that
// follows a successful invocation of this operation, but to the following
// maintenance process instead.
//
// To see the current maintenance configuration of your application, invoke the DescribeApplication
// operation.
//
// For information about application maintenance, see [Managed Service for Apache Flink for Apache Flink Maintenance].
//
// This operation is supported only for Managed Service for Apache Flink.
//
// [Managed Service for Apache Flink for Apache Flink Maintenance]: https://docs.aws.amazon.com/kinesisanalytics/latest/java/maintenance.html
func kinesisanalyticsv2_UpdateApplicationMaintenanceConfiguration(cfg aws.Config, client *kinesisanalyticsv2.Client) {
	input := &kinesisanalyticsv2.UpdateApplicationMaintenanceConfigurationInput{
		// ApplicationMaintenanceConfigurationUpdate: *types.ApplicationMaintenanceConfigurationUpdate, // Required
		// ApplicationName: *string, // Required
	}

	if len(_kinesisanalyticsv2ApplicationMaintenanceConfigurationUpdate) > 0 {
		if err := assignInputField(input, "ApplicationMaintenanceConfigurationUpdate", _kinesisanalyticsv2ApplicationMaintenanceConfigurationUpdate); err != nil {
			log.Errorf("invalid --application-maintenance-configuration-update: %s", err.Error())
			return
		}
	}
	if len(_kinesisanalyticsv2ApplicationName) > 0 {
		input.ApplicationName = aws.String(_kinesisanalyticsv2ApplicationName)
	}

	if resp, err := client.UpdateApplicationMaintenanceConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_kinesisanalyticsv2Cmd)
	_kinesisanalyticsv2Cmd.Flags().SortFlags = false

	_kinesisanalyticsv2Cmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_kinesisanalyticsv2Cmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_kinesisanalyticsv2Cmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_kinesisanalyticsv2Cmd.Flags().StringVarP(&_kinesisanalyticsv2ApplicationConfiguration, "application-configuration", "", "", "Application Configuration")
	_kinesisanalyticsv2Cmd.Flags().StringVarP(&_kinesisanalyticsv2ApplicationConfigurationUpdate, "application-configuration-update", "", "", "Application Configuration Update")
	_kinesisanalyticsv2Cmd.Flags().StringVarP(&_kinesisanalyticsv2ApplicationDescription, "application-description", "", "", "Application Description")
	_kinesisanalyticsv2Cmd.Flags().StringVarP(&_kinesisanalyticsv2ApplicationMaintenanceConfigurationUpdate, "application-maintenance-configuration-update", "", "", "Application Maintenance Configuration Update")
	_kinesisanalyticsv2Cmd.Flags().StringVarP(&_kinesisanalyticsv2ApplicationMode, "application-mode", "", "", "Application Mode")
	_kinesisanalyticsv2Cmd.Flags().StringVarP(&_kinesisanalyticsv2ApplicationName, "application-name", "", "", "Application Name")
	_kinesisanalyticsv2Cmd.Flags().StringVarP(&_kinesisanalyticsv2ApplicationVersionId, "application-version-id", "", "", "Application Version ID")
	_kinesisanalyticsv2Cmd.Flags().StringVarP(&_kinesisanalyticsv2CloudWatchLoggingOption, "cloud-watch-logging-option", "", "", "Cloud Watch Logging Option")
	_kinesisanalyticsv2Cmd.Flags().StringVarP(&_kinesisanalyticsv2CloudWatchLoggingOptionId, "cloud-watch-logging-option-id", "", "", "Cloud Watch Logging Option ID")
	_kinesisanalyticsv2Cmd.Flags().StringVarP(&_kinesisanalyticsv2CloudWatchLoggingOptionUpdates, "cloud-watch-logging-option-updates", "", "", "Cloud Watch Logging Option Updates")
	_kinesisanalyticsv2Cmd.Flags().StringVarP(&_kinesisanalyticsv2CloudWatchLoggingOptions, "cloud-watch-logging-options", "", "", "Cloud Watch Logging Options")
	_kinesisanalyticsv2Cmd.Flags().StringVarP(&_kinesisanalyticsv2ConditionalToken, "conditional-token", "", "", "Conditional Token")
	_kinesisanalyticsv2Cmd.Flags().StringVarP(&_kinesisanalyticsv2CreateTimestamp, "create-timestamp", "", "", "Create Timestamp")
	_kinesisanalyticsv2Cmd.Flags().StringVarP(&_kinesisanalyticsv2CurrentApplicationVersionId, "current-application-version-id", "", "", "Current Application Version ID")
	_kinesisanalyticsv2Cmd.Flags().StringVarP(&_kinesisanalyticsv2Force, "force", "", "", "Force")
	_kinesisanalyticsv2Cmd.Flags().StringVarP(&_kinesisanalyticsv2IncludeAdditionalDetails, "include-additional-details", "", "", "Include Additional Details")
	_kinesisanalyticsv2Cmd.Flags().StringVarP(&_kinesisanalyticsv2Input, "input", "", "", "Input")
	_kinesisanalyticsv2Cmd.Flags().StringVarP(&_kinesisanalyticsv2InputId, "input-id", "", "", "Input ID")
	_kinesisanalyticsv2Cmd.Flags().StringVarP(&_kinesisanalyticsv2InputProcessingConfiguration, "input-processing-configuration", "", "", "Input Processing Configuration")
	_kinesisanalyticsv2Cmd.Flags().StringVarP(&_kinesisanalyticsv2InputStartingPositionConfiguration, "input-starting-position-configuration", "", "", "Input Starting Position Configuration")
	_kinesisanalyticsv2Cmd.Flags().StringVarP(&_kinesisanalyticsv2Limit, "limit", "", "", "Limit")
	_kinesisanalyticsv2Cmd.Flags().StringVarP(&_kinesisanalyticsv2NextToken, "next-token", "", "", "Next Token")
	_kinesisanalyticsv2Cmd.Flags().StringVarP(&_kinesisanalyticsv2Operation, "operation", "", "", "Operation")
	_kinesisanalyticsv2Cmd.Flags().StringVarP(&_kinesisanalyticsv2OperationId, "operation-id", "", "", "Operation ID")
	_kinesisanalyticsv2Cmd.Flags().StringVarP(&_kinesisanalyticsv2OperationStatus, "operation-status", "", "", "Operation Status")
	_kinesisanalyticsv2Cmd.Flags().StringVarP(&_kinesisanalyticsv2OutputId, "output-id", "", "", "Output ID")
	_kinesisanalyticsv2Cmd.Flags().StringVarP(&_kinesisanalyticsv2ReferenceDataSource, "reference-data-source", "", "", "Reference Data Source")
	_kinesisanalyticsv2Cmd.Flags().StringVarP(&_kinesisanalyticsv2ReferenceId, "reference-id", "", "", "Reference ID")
	_kinesisanalyticsv2Cmd.Flags().StringVarP(&_kinesisanalyticsv2ResourceARN, "resource-arn", "", "", "Resource ARN")
	_kinesisanalyticsv2Cmd.Flags().StringVarP(&_kinesisanalyticsv2RunConfiguration, "run-configuration", "", "", "Run Configuration")
	_kinesisanalyticsv2Cmd.Flags().StringVarP(&_kinesisanalyticsv2RunConfigurationUpdate, "run-configuration-update", "", "", "Run Configuration Update")
	_kinesisanalyticsv2Cmd.Flags().StringVarP(&_kinesisanalyticsv2RuntimeEnvironment, "runtime-environment", "", "", "Runtime Environment")
	_kinesisanalyticsv2Cmd.Flags().StringVarP(&_kinesisanalyticsv2RuntimeEnvironmentUpdate, "runtime-environment-update", "", "", "Runtime Environment Update")
	_kinesisanalyticsv2Cmd.Flags().StringVarP(&_kinesisanalyticsv2S3Configuration, "s3-configuration", "", "", "S3 Configuration")
	_kinesisanalyticsv2Cmd.Flags().StringVarP(&_kinesisanalyticsv2ServiceExecutionRole, "service-execution-role", "", "", "Service Execution Role")
	_kinesisanalyticsv2Cmd.Flags().StringVarP(&_kinesisanalyticsv2ServiceExecutionRoleUpdate, "service-execution-role-update", "", "", "Service Execution Role Update")
	_kinesisanalyticsv2Cmd.Flags().StringVarP(&_kinesisanalyticsv2SessionExpirationDurationInSeconds, "session-expiration-duration-in-seconds", "", "", "Session Expiration Duration In Seconds")
	_kinesisanalyticsv2Cmd.Flags().StringVarP(&_kinesisanalyticsv2SnapshotCreationTimestamp, "snapshot-creation-timestamp", "", "", "Snapshot Creation Timestamp")
	_kinesisanalyticsv2Cmd.Flags().StringVarP(&_kinesisanalyticsv2SnapshotName, "snapshot-name", "", "", "Snapshot Name")
	_kinesisanalyticsv2Cmd.Flags().StringSliceVarP(&_kinesisanalyticsv2TagKeys, "tag-keys", "", nil, "Tag Keys")
	_kinesisanalyticsv2Cmd.Flags().StringVarP(&_kinesisanalyticsv2Tags, "tags", "", "", "Tags")
	_kinesisanalyticsv2Cmd.Flags().StringVarP(&_kinesisanalyticsv2UrlType, "url-type", "", "", "URL Type")
	_kinesisanalyticsv2Cmd.Flags().StringVarP(&_kinesisanalyticsv2VpcConfiguration, "vpc-configuration", "", "", "VPC Configuration")
	_kinesisanalyticsv2Cmd.Flags().StringVarP(&_kinesisanalyticsv2VpcConfigurationId, "vpc-configuration-id", "", "", "VPC Configuration ID")

	_kinesisanalyticsv2Cmd.Flags().BoolVarP(&_kinesisanalyticsv2AddApplicationCloudWatchLoggingOption, "add-application-cloud-watch-logging-option", "", false, "Add Application Cloud Watch Logging Option")
	_kinesisanalyticsv2Cmd.Flags().BoolVarP(&_kinesisanalyticsv2AddApplicationInput, "add-application-input", "", false, "Add Application Input")
	_kinesisanalyticsv2Cmd.Flags().BoolVarP(&_kinesisanalyticsv2AddApplicationInputProcessingConfiguration, "add-application-input-processing-configuration", "", false, "Add Application Input Processing Configuration")
	_kinesisanalyticsv2Cmd.Flags().BoolVarP(&_kinesisanalyticsv2AddApplicationOutput, "add-application-output", "", false, "Add Application Output")
	_kinesisanalyticsv2Cmd.Flags().BoolVarP(&_kinesisanalyticsv2AddApplicationReferenceDataSource, "add-application-reference-data-source", "", false, "Add Application Reference Data Source")
	_kinesisanalyticsv2Cmd.Flags().BoolVarP(&_kinesisanalyticsv2AddApplicationVpcConfiguration, "add-application-vpc-configuration", "", false, "Add Application VPC Configuration")
	_kinesisanalyticsv2Cmd.Flags().BoolVarP(&_kinesisanalyticsv2CreateApplication, "create-application", "", false, "Create Application")
	_kinesisanalyticsv2Cmd.Flags().BoolVarP(&_kinesisanalyticsv2CreateApplicationPresignedUrl, "create-application-presigned-url", "", false, "Create Application Presigned URL")
	_kinesisanalyticsv2Cmd.Flags().BoolVarP(&_kinesisanalyticsv2CreateApplicationSnapshot, "create-application-snapshot", "", false, "Create Application Snapshot")
	_kinesisanalyticsv2Cmd.Flags().BoolVarP(&_kinesisanalyticsv2DeleteApplication, "delete-application", "", false, "Delete Application")
	_kinesisanalyticsv2Cmd.Flags().BoolVarP(&_kinesisanalyticsv2DeleteApplicationCloudWatchLoggingOption, "delete-application-cloud-watch-logging-option", "", false, "Delete Application Cloud Watch Logging Option")
	_kinesisanalyticsv2Cmd.Flags().BoolVarP(&_kinesisanalyticsv2DeleteApplicationInputProcessingConfiguration, "delete-application-input-processing-configuration", "", false, "Delete Application Input Processing Configuration")
	_kinesisanalyticsv2Cmd.Flags().BoolVarP(&_kinesisanalyticsv2DeleteApplicationOutput, "delete-application-output", "", false, "Delete Application Output")
	_kinesisanalyticsv2Cmd.Flags().BoolVarP(&_kinesisanalyticsv2DeleteApplicationReferenceDataSource, "delete-application-reference-data-source", "", false, "Delete Application Reference Data Source")
	_kinesisanalyticsv2Cmd.Flags().BoolVarP(&_kinesisanalyticsv2DeleteApplicationSnapshot, "delete-application-snapshot", "", false, "Delete Application Snapshot")
	_kinesisanalyticsv2Cmd.Flags().BoolVarP(&_kinesisanalyticsv2DeleteApplicationVpcConfiguration, "delete-application-vpc-configuration", "", false, "Delete Application VPC Configuration")
	_kinesisanalyticsv2Cmd.Flags().BoolVarP(&_kinesisanalyticsv2DescribeApplication, "describe-application", "", false, "Describe Application")
	_kinesisanalyticsv2Cmd.Flags().BoolVarP(&_kinesisanalyticsv2DescribeApplicationOperation, "describe-application-operation", "", false, "Describe Application Operation")
	_kinesisanalyticsv2Cmd.Flags().BoolVarP(&_kinesisanalyticsv2DescribeApplicationSnapshot, "describe-application-snapshot", "", false, "Describe Application Snapshot")
	_kinesisanalyticsv2Cmd.Flags().BoolVarP(&_kinesisanalyticsv2DescribeApplicationVersion, "describe-application-version", "", false, "Describe Application Version")
	_kinesisanalyticsv2Cmd.Flags().BoolVarP(&_kinesisanalyticsv2DiscoverInputSchema, "discover-input-schema", "", false, "Discover Input Schema")
	_kinesisanalyticsv2Cmd.Flags().BoolVarP(&_kinesisanalyticsv2ListApplicationOperations, "list-application-operations", "", false, "List Application Operations")
	_kinesisanalyticsv2Cmd.Flags().BoolVarP(&_kinesisanalyticsv2ListApplicationSnapshots, "list-application-snapshots", "", false, "List Application Snapshots")
	_kinesisanalyticsv2Cmd.Flags().BoolVarP(&_kinesisanalyticsv2ListApplicationVersions, "list-application-versions", "", false, "List Application Versions")
	_kinesisanalyticsv2Cmd.Flags().BoolVarP(&_kinesisanalyticsv2ListApplications, "list-applications", "", false, "List Applications")
	_kinesisanalyticsv2Cmd.Flags().BoolVarP(&_kinesisanalyticsv2ListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_kinesisanalyticsv2Cmd.Flags().BoolVarP(&_kinesisanalyticsv2RollbackApplication, "rollback-application", "", false, "Rollback Application")
	_kinesisanalyticsv2Cmd.Flags().BoolVarP(&_kinesisanalyticsv2StartApplication, "start-application", "", false, "Start Application")
	_kinesisanalyticsv2Cmd.Flags().BoolVarP(&_kinesisanalyticsv2StopApplication, "stop-application", "", false, "Stop Application")
	_kinesisanalyticsv2Cmd.Flags().BoolVarP(&_kinesisanalyticsv2TagResource, "tag-resource", "", false, "Tag Resource")
	_kinesisanalyticsv2Cmd.Flags().BoolVarP(&_kinesisanalyticsv2UntagResource, "untag-resource", "", false, "Untag Resource")
	_kinesisanalyticsv2Cmd.Flags().BoolVarP(&_kinesisanalyticsv2UpdateApplication, "update-application", "", false, "Update Application")
	_kinesisanalyticsv2Cmd.Flags().BoolVarP(&_kinesisanalyticsv2UpdateApplicationMaintenanceConfiguration, "update-application-maintenance-configuration", "", false, "Update Application Maintenance Configuration")

}
