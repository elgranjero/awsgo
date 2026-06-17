package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mwaa"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// mwaaCmd represents the mwaa command
var _mwaaCmd = &cobra.Command{
	Use:   "mwaa",
	Short: "AWS mwaa CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := mwaa.NewFromConfig(cfg)
		if _mwaaCreateCliToken {
			mwaa_CreateCliToken(cfg, client)
			return
		}
		if _mwaaCreateEnvironment {
			mwaa_CreateEnvironment(cfg, client)
			return
		}
		if _mwaaCreateWebLoginToken {
			mwaa_CreateWebLoginToken(cfg, client)
			return
		}
		if _mwaaDeleteEnvironment {
			mwaa_DeleteEnvironment(cfg, client)
			return
		}
		if _mwaaGetEnvironment {
			mwaa_GetEnvironment(cfg, client)
			return
		}
		if _mwaaInvokeRestApi {
			mwaa_InvokeRestApi(cfg, client)
			return
		}
		if _mwaaListEnvironments {
			mwaa_ListEnvironments(cfg, client)
			return
		}
		if _mwaaListTagsForResource {
			mwaa_ListTagsForResource(cfg, client)
			return
		}
		if _mwaaPublishMetrics {
			mwaa_PublishMetrics(cfg, client)
			return
		}
		if _mwaaTagResource {
			mwaa_TagResource(cfg, client)
			return
		}
		if _mwaaUntagResource {
			mwaa_UntagResource(cfg, client)
			return
		}
		if _mwaaUpdateEnvironment {
			mwaa_UpdateEnvironment(cfg, client)
			return
		}

	},
}

var (
	_mwaaCreateCliToken      bool
	_mwaaCreateEnvironment   bool
	_mwaaCreateWebLoginToken bool
	_mwaaDeleteEnvironment   bool
	_mwaaGetEnvironment      bool
	_mwaaInvokeRestApi       bool
	_mwaaListEnvironments    bool
	_mwaaListTagsForResource bool
	_mwaaPublishMetrics      bool
	_mwaaTagResource         bool
	_mwaaUntagResource       bool
	_mwaaUpdateEnvironment   bool

	_mwaaAirflowConfigurationOptions  string
	_mwaaAirflowVersion               string
	_mwaaBody                         string
	_mwaaDagS3Path                    string
	_mwaaEndpointManagement           string
	_mwaaEnvironmentClass             string
	_mwaaEnvironmentName              string
	_mwaaExecutionRoleArn             string
	_mwaaKmsKey                       string
	_mwaaLoggingConfiguration         string
	_mwaaMaxResults                   string
	_mwaaMaxWebservers                string
	_mwaaMaxWorkers                   string
	_mwaaMethod                       string
	_mwaaMetricData                   string
	_mwaaMinWebservers                string
	_mwaaMinWorkers                   string
	_mwaaName                         string
	_mwaaNetworkConfiguration         string
	_mwaaNextToken                    string
	_mwaaPath                         string
	_mwaaPluginsS3ObjectVersion       string
	_mwaaPluginsS3Path                string
	_mwaaQueryParameters              string
	_mwaaRequirementsS3ObjectVersion  string
	_mwaaRequirementsS3Path           string
	_mwaaResourceArn                  string
	_mwaaSchedulers                   string
	_mwaaSourceBucketArn              string
	_mwaaStartupScriptS3ObjectVersion string
	_mwaaStartupScriptS3Path          string
	_mwaaTagKeys                      []string
	_mwaaTags                         string
	_mwaaWebserverAccessMode          string
	_mwaaWeeklyMaintenanceWindowStart string
	_mwaaWorkerReplacementStrategy    string
)

// Creates a CLI token for the Airflow CLI. To learn more, see [Creating an Apache Airflow CLI token].
//
// [Creating an Apache Airflow CLI token]: https://docs.aws.amazon.com/mwaa/latest/userguide/call-mwaa-apis-cli.html
func mwaa_CreateCliToken(cfg aws.Config, client *mwaa.Client) {
	input := &mwaa.CreateCliTokenInput{
		// Name: *string, // Required
	}

	if len(_mwaaName) > 0 {
		input.Name = aws.String(_mwaaName)
	}

	if resp, err := client.CreateCliToken(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Managed Workflows for Apache Airflow (Amazon MWAA)
// environment.
func mwaa_CreateEnvironment(cfg aws.Config, client *mwaa.Client) {
	input := &mwaa.CreateEnvironmentInput{
		// DagS3Path: *string, // Required
		// ExecutionRoleArn: *string, // Required
		// Name: *string, // Required
		// NetworkConfiguration: *types.NetworkConfiguration, // Required
		// SourceBucketArn: *string, // Required
	}

	if len(_mwaaDagS3Path) > 0 {
		input.DagS3Path = aws.String(_mwaaDagS3Path)
	}
	if len(_mwaaExecutionRoleArn) > 0 {
		input.ExecutionRoleArn = aws.String(_mwaaExecutionRoleArn)
	}
	if len(_mwaaName) > 0 {
		input.Name = aws.String(_mwaaName)
	}
	if len(_mwaaNetworkConfiguration) > 0 {
		if err := assignInputField(input, "NetworkConfiguration", _mwaaNetworkConfiguration); err != nil {
			log.Errorf("invalid --network-configuration: %s", err.Error())
			return
		}
	}
	if len(_mwaaSourceBucketArn) > 0 {
		input.SourceBucketArn = aws.String(_mwaaSourceBucketArn)
	}
	if len(_mwaaAirflowConfigurationOptions) > 0 {
		if err := assignInputField(input, "AirflowConfigurationOptions", _mwaaAirflowConfigurationOptions); err != nil {
			log.Errorf("invalid --airflow-configuration-options: %s", err.Error())
			return
		}
	}
	if len(_mwaaAirflowVersion) > 0 {
		input.AirflowVersion = aws.String(_mwaaAirflowVersion)
	}
	if len(_mwaaEndpointManagement) > 0 {
		if err := assignInputField(input, "EndpointManagement", _mwaaEndpointManagement); err != nil {
			log.Errorf("invalid --endpoint-management: %s", err.Error())
			return
		}
	}
	if len(_mwaaEnvironmentClass) > 0 {
		input.EnvironmentClass = aws.String(_mwaaEnvironmentClass)
	}
	if len(_mwaaKmsKey) > 0 {
		input.KmsKey = aws.String(_mwaaKmsKey)
	}
	if len(_mwaaLoggingConfiguration) > 0 {
		if err := assignInputField(input, "LoggingConfiguration", _mwaaLoggingConfiguration); err != nil {
			log.Errorf("invalid --logging-configuration: %s", err.Error())
			return
		}
	}
	if len(_mwaaMaxWebservers) > 0 {
		if err := assignInputField(input, "MaxWebservers", _mwaaMaxWebservers); err != nil {
			log.Errorf("invalid --max-webservers: %s", err.Error())
			return
		}
	}
	if len(_mwaaMaxWorkers) > 0 {
		if err := assignInputField(input, "MaxWorkers", _mwaaMaxWorkers); err != nil {
			log.Errorf("invalid --max-workers: %s", err.Error())
			return
		}
	}
	if len(_mwaaMinWebservers) > 0 {
		if err := assignInputField(input, "MinWebservers", _mwaaMinWebservers); err != nil {
			log.Errorf("invalid --min-webservers: %s", err.Error())
			return
		}
	}
	if len(_mwaaMinWorkers) > 0 {
		if err := assignInputField(input, "MinWorkers", _mwaaMinWorkers); err != nil {
			log.Errorf("invalid --min-workers: %s", err.Error())
			return
		}
	}
	if len(_mwaaPluginsS3ObjectVersion) > 0 {
		input.PluginsS3ObjectVersion = aws.String(_mwaaPluginsS3ObjectVersion)
	}
	if len(_mwaaPluginsS3Path) > 0 {
		input.PluginsS3Path = aws.String(_mwaaPluginsS3Path)
	}
	if len(_mwaaRequirementsS3ObjectVersion) > 0 {
		input.RequirementsS3ObjectVersion = aws.String(_mwaaRequirementsS3ObjectVersion)
	}
	if len(_mwaaRequirementsS3Path) > 0 {
		input.RequirementsS3Path = aws.String(_mwaaRequirementsS3Path)
	}
	if len(_mwaaSchedulers) > 0 {
		if err := assignInputField(input, "Schedulers", _mwaaSchedulers); err != nil {
			log.Errorf("invalid --schedulers: %s", err.Error())
			return
		}
	}
	if len(_mwaaStartupScriptS3ObjectVersion) > 0 {
		input.StartupScriptS3ObjectVersion = aws.String(_mwaaStartupScriptS3ObjectVersion)
	}
	if len(_mwaaStartupScriptS3Path) > 0 {
		input.StartupScriptS3Path = aws.String(_mwaaStartupScriptS3Path)
	}
	if len(_mwaaTags) > 0 {
		if err := assignInputField(input, "Tags", _mwaaTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_mwaaWebserverAccessMode) > 0 {
		if err := assignInputField(input, "WebserverAccessMode", _mwaaWebserverAccessMode); err != nil {
			log.Errorf("invalid --webserver-access-mode: %s", err.Error())
			return
		}
	}
	if len(_mwaaWeeklyMaintenanceWindowStart) > 0 {
		input.WeeklyMaintenanceWindowStart = aws.String(_mwaaWeeklyMaintenanceWindowStart)
	}

	if resp, err := client.CreateEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a web login token for the Airflow Web UI. To learn more, see [Creating an Apache Airflow web login token].
//
// [Creating an Apache Airflow web login token]: https://docs.aws.amazon.com/mwaa/latest/userguide/call-mwaa-apis-web.html
func mwaa_CreateWebLoginToken(cfg aws.Config, client *mwaa.Client) {
	input := &mwaa.CreateWebLoginTokenInput{
		// Name: *string, // Required
	}

	if len(_mwaaName) > 0 {
		input.Name = aws.String(_mwaaName)
	}

	if resp, err := client.CreateWebLoginToken(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon Managed Workflows for Apache Airflow (Amazon MWAA)
// environment.
func mwaa_DeleteEnvironment(cfg aws.Config, client *mwaa.Client) {
	input := &mwaa.DeleteEnvironmentInput{
		// Name: *string, // Required
	}

	if len(_mwaaName) > 0 {
		input.Name = aws.String(_mwaaName)
	}

	if resp, err := client.DeleteEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes an Amazon Managed Workflows for Apache Airflow (MWAA) environment.
func mwaa_GetEnvironment(cfg aws.Config, client *mwaa.Client) {
	input := &mwaa.GetEnvironmentInput{
		// Name: *string, // Required
	}

	if len(_mwaaName) > 0 {
		input.Name = aws.String(_mwaaName)
	}

	if resp, err := client.GetEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Invokes the Apache Airflow REST API on the webserver with the specified inputs.
// To learn more, see [Using the Apache Airflow REST API]
//
// [Using the Apache Airflow REST API]: https://docs.aws.amazon.com/mwaa/latest/userguide/access-mwaa-apache-airflow-rest-api.html
func mwaa_InvokeRestApi(cfg aws.Config, client *mwaa.Client) {
	input := &mwaa.InvokeRestApiInput{
		// Method: types.RestApiMethod, // Required
		// Name: *string, // Required
		// Path: *string, // Required
	}

	if len(_mwaaMethod) > 0 {
		if err := assignInputField(input, "Method", _mwaaMethod); err != nil {
			log.Errorf("invalid --method: %s", err.Error())
			return
		}
	}
	if len(_mwaaName) > 0 {
		input.Name = aws.String(_mwaaName)
	}
	if len(_mwaaPath) > 0 {
		input.Path = aws.String(_mwaaPath)
	}
	if len(_mwaaBody) > 0 {
		if err := assignInputField(input, "Body", _mwaaBody); err != nil {
			log.Errorf("invalid --body: %s", err.Error())
			return
		}
	}
	if len(_mwaaQueryParameters) > 0 {
		if err := assignInputField(input, "QueryParameters", _mwaaQueryParameters); err != nil {
			log.Errorf("invalid --query-parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.InvokeRestApi(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the Amazon Managed Workflows for Apache Airflow (MWAA) environments.
func mwaa_ListEnvironments(cfg aws.Config, client *mwaa.Client) {
	input := &mwaa.ListEnvironmentsInput{}

	if len(_mwaaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mwaaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mwaaNextToken) > 0 {
		input.NextToken = aws.String(_mwaaNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEnvironments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mwaa.ListEnvironmentsOutput
	p := mwaa.NewListEnvironmentsPaginator(client, input)
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

// Lists the key-value tag pairs associated to the Amazon Managed Workflows for
// Apache Airflow (MWAA) environment. For example, "Environment": "Staging" .
func mwaa_ListTagsForResource(cfg aws.Config, client *mwaa.Client) {
	input := &mwaa.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_mwaaResourceArn) > 0 {
		input.ResourceArn = aws.String(_mwaaResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Internal only. Publishes environment health metrics to Amazon CloudWatch.
// Deprecated: This API is for internal use and not meant for public use, and is
// no longer available.
func mwaa_PublishMetrics(cfg aws.Config, client *mwaa.Client) {
	input := &mwaa.PublishMetricsInput{
		// EnvironmentName: *string, // Required
		// MetricData: []types.MetricDatum, // Required
	}

	if len(_mwaaEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_mwaaEnvironmentName)
	}
	if len(_mwaaMetricData) > 0 {
		if err := assignInputField(input, "MetricData", _mwaaMetricData); err != nil {
			log.Errorf("invalid --metric-data: %s", err.Error())
			return
		}
	}

	if resp, err := client.PublishMetrics(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates key-value tag pairs to your Amazon Managed Workflows for Apache
// Airflow (MWAA) environment.
func mwaa_TagResource(cfg aws.Config, client *mwaa.Client) {
	input := &mwaa.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_mwaaResourceArn) > 0 {
		input.ResourceArn = aws.String(_mwaaResourceArn)
	}
	if len(_mwaaTags) > 0 {
		if err := assignInputField(input, "Tags", _mwaaTags); err != nil {
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

// Removes key-value tag pairs associated to your Amazon Managed Workflows for
// Apache Airflow (MWAA) environment. For example, "Environment": "Staging" .
func mwaa_UntagResource(cfg aws.Config, client *mwaa.Client) {
	input := &mwaa.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_mwaaResourceArn) > 0 {
		input.ResourceArn = aws.String(_mwaaResourceArn)
	}
	if len(_mwaaTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _mwaaTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an Amazon Managed Workflows for Apache Airflow (MWAA) environment.
func mwaa_UpdateEnvironment(cfg aws.Config, client *mwaa.Client) {
	input := &mwaa.UpdateEnvironmentInput{
		// Name: *string, // Required
	}

	if len(_mwaaName) > 0 {
		input.Name = aws.String(_mwaaName)
	}
	if len(_mwaaAirflowConfigurationOptions) > 0 {
		if err := assignInputField(input, "AirflowConfigurationOptions", _mwaaAirflowConfigurationOptions); err != nil {
			log.Errorf("invalid --airflow-configuration-options: %s", err.Error())
			return
		}
	}
	if len(_mwaaAirflowVersion) > 0 {
		input.AirflowVersion = aws.String(_mwaaAirflowVersion)
	}
	if len(_mwaaDagS3Path) > 0 {
		input.DagS3Path = aws.String(_mwaaDagS3Path)
	}
	if len(_mwaaEnvironmentClass) > 0 {
		input.EnvironmentClass = aws.String(_mwaaEnvironmentClass)
	}
	if len(_mwaaExecutionRoleArn) > 0 {
		input.ExecutionRoleArn = aws.String(_mwaaExecutionRoleArn)
	}
	if len(_mwaaLoggingConfiguration) > 0 {
		if err := assignInputField(input, "LoggingConfiguration", _mwaaLoggingConfiguration); err != nil {
			log.Errorf("invalid --logging-configuration: %s", err.Error())
			return
		}
	}
	if len(_mwaaMaxWebservers) > 0 {
		if err := assignInputField(input, "MaxWebservers", _mwaaMaxWebservers); err != nil {
			log.Errorf("invalid --max-webservers: %s", err.Error())
			return
		}
	}
	if len(_mwaaMaxWorkers) > 0 {
		if err := assignInputField(input, "MaxWorkers", _mwaaMaxWorkers); err != nil {
			log.Errorf("invalid --max-workers: %s", err.Error())
			return
		}
	}
	if len(_mwaaMinWebservers) > 0 {
		if err := assignInputField(input, "MinWebservers", _mwaaMinWebservers); err != nil {
			log.Errorf("invalid --min-webservers: %s", err.Error())
			return
		}
	}
	if len(_mwaaMinWorkers) > 0 {
		if err := assignInputField(input, "MinWorkers", _mwaaMinWorkers); err != nil {
			log.Errorf("invalid --min-workers: %s", err.Error())
			return
		}
	}
	if len(_mwaaNetworkConfiguration) > 0 {
		if err := assignInputField(input, "NetworkConfiguration", _mwaaNetworkConfiguration); err != nil {
			log.Errorf("invalid --network-configuration: %s", err.Error())
			return
		}
	}
	if len(_mwaaPluginsS3ObjectVersion) > 0 {
		input.PluginsS3ObjectVersion = aws.String(_mwaaPluginsS3ObjectVersion)
	}
	if len(_mwaaPluginsS3Path) > 0 {
		input.PluginsS3Path = aws.String(_mwaaPluginsS3Path)
	}
	if len(_mwaaRequirementsS3ObjectVersion) > 0 {
		input.RequirementsS3ObjectVersion = aws.String(_mwaaRequirementsS3ObjectVersion)
	}
	if len(_mwaaRequirementsS3Path) > 0 {
		input.RequirementsS3Path = aws.String(_mwaaRequirementsS3Path)
	}
	if len(_mwaaSchedulers) > 0 {
		if err := assignInputField(input, "Schedulers", _mwaaSchedulers); err != nil {
			log.Errorf("invalid --schedulers: %s", err.Error())
			return
		}
	}
	if len(_mwaaSourceBucketArn) > 0 {
		input.SourceBucketArn = aws.String(_mwaaSourceBucketArn)
	}
	if len(_mwaaStartupScriptS3ObjectVersion) > 0 {
		input.StartupScriptS3ObjectVersion = aws.String(_mwaaStartupScriptS3ObjectVersion)
	}
	if len(_mwaaStartupScriptS3Path) > 0 {
		input.StartupScriptS3Path = aws.String(_mwaaStartupScriptS3Path)
	}
	if len(_mwaaWebserverAccessMode) > 0 {
		if err := assignInputField(input, "WebserverAccessMode", _mwaaWebserverAccessMode); err != nil {
			log.Errorf("invalid --webserver-access-mode: %s", err.Error())
			return
		}
	}
	if len(_mwaaWeeklyMaintenanceWindowStart) > 0 {
		input.WeeklyMaintenanceWindowStart = aws.String(_mwaaWeeklyMaintenanceWindowStart)
	}
	if len(_mwaaWorkerReplacementStrategy) > 0 {
		if err := assignInputField(input, "WorkerReplacementStrategy", _mwaaWorkerReplacementStrategy); err != nil {
			log.Errorf("invalid --worker-replacement-strategy: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_mwaaCmd)
	_mwaaCmd.Flags().SortFlags = false

	_mwaaCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_mwaaCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_mwaaCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_mwaaCmd.Flags().StringVarP(&_mwaaAirflowConfigurationOptions, "airflow-configuration-options", "", "", "Airflow Configuration Options")
	_mwaaCmd.Flags().StringVarP(&_mwaaAirflowVersion, "airflow-version", "", "", "Airflow Version")
	_mwaaCmd.Flags().StringVarP(&_mwaaBody, "body", "", "", "Body")
	_mwaaCmd.Flags().StringVarP(&_mwaaDagS3Path, "dag-s3-path", "", "", "Dag S3 Path")
	_mwaaCmd.Flags().StringVarP(&_mwaaEndpointManagement, "endpoint-management", "", "", "Endpoint Management")
	_mwaaCmd.Flags().StringVarP(&_mwaaEnvironmentClass, "environment-class", "", "", "Environment Class")
	_mwaaCmd.Flags().StringVarP(&_mwaaEnvironmentName, "environment-name", "", "", "Environment Name")
	_mwaaCmd.Flags().StringVarP(&_mwaaExecutionRoleArn, "execution-role-arn", "", "", "Execution Role ARN")
	_mwaaCmd.Flags().StringVarP(&_mwaaKmsKey, "kms-key", "", "", "KMS Key")
	_mwaaCmd.Flags().StringVarP(&_mwaaLoggingConfiguration, "logging-configuration", "", "", "Logging Configuration")
	_mwaaCmd.Flags().StringVarP(&_mwaaMaxResults, "max-results", "", "", "Max Results")
	_mwaaCmd.Flags().StringVarP(&_mwaaMaxWebservers, "max-webservers", "", "", "Max Webservers")
	_mwaaCmd.Flags().StringVarP(&_mwaaMaxWorkers, "max-workers", "", "", "Max Workers")
	_mwaaCmd.Flags().StringVarP(&_mwaaMethod, "method", "", "", "Method")
	_mwaaCmd.Flags().StringVarP(&_mwaaMetricData, "metric-data", "", "", "Metric Data")
	_mwaaCmd.Flags().StringVarP(&_mwaaMinWebservers, "min-webservers", "", "", "Min Webservers")
	_mwaaCmd.Flags().StringVarP(&_mwaaMinWorkers, "min-workers", "", "", "Min Workers")
	_mwaaCmd.Flags().StringVarP(&_mwaaName, "name", "", "", "Name")
	_mwaaCmd.Flags().StringVarP(&_mwaaNetworkConfiguration, "network-configuration", "", "", "Network Configuration")
	_mwaaCmd.Flags().StringVarP(&_mwaaNextToken, "next-token", "", "", "Next Token")
	_mwaaCmd.Flags().StringVarP(&_mwaaPath, "path", "", "", "Path")
	_mwaaCmd.Flags().StringVarP(&_mwaaPluginsS3ObjectVersion, "plugins-s3-object-version", "", "", "Plugins S3 Object Version")
	_mwaaCmd.Flags().StringVarP(&_mwaaPluginsS3Path, "plugins-s3-path", "", "", "Plugins S3 Path")
	_mwaaCmd.Flags().StringVarP(&_mwaaQueryParameters, "query-parameters", "", "", "Query Parameters")
	_mwaaCmd.Flags().StringVarP(&_mwaaRequirementsS3ObjectVersion, "requirements-s3-object-version", "", "", "Requirements S3 Object Version")
	_mwaaCmd.Flags().StringVarP(&_mwaaRequirementsS3Path, "requirements-s3-path", "", "", "Requirements S3 Path")
	_mwaaCmd.Flags().StringVarP(&_mwaaResourceArn, "resource-arn", "", "", "Resource ARN")
	_mwaaCmd.Flags().StringVarP(&_mwaaSchedulers, "schedulers", "", "", "Schedulers")
	_mwaaCmd.Flags().StringVarP(&_mwaaSourceBucketArn, "source-bucket-arn", "", "", "Source Bucket ARN")
	_mwaaCmd.Flags().StringVarP(&_mwaaStartupScriptS3ObjectVersion, "startup-script-s3-object-version", "", "", "Startup Script S3 Object Version")
	_mwaaCmd.Flags().StringVarP(&_mwaaStartupScriptS3Path, "startup-script-s3-path", "", "", "Startup Script S3 Path")
	_mwaaCmd.Flags().StringSliceVarP(&_mwaaTagKeys, "tag-keys", "", nil, "Tag Keys")
	_mwaaCmd.Flags().StringVarP(&_mwaaTags, "tags", "", "", "Tags")
	_mwaaCmd.Flags().StringVarP(&_mwaaWebserverAccessMode, "webserver-access-mode", "", "", "Webserver Access Mode")
	_mwaaCmd.Flags().StringVarP(&_mwaaWeeklyMaintenanceWindowStart, "weekly-maintenance-window-start", "", "", "Weekly Maintenance Window Start")
	_mwaaCmd.Flags().StringVarP(&_mwaaWorkerReplacementStrategy, "worker-replacement-strategy", "", "", "Worker Replacement Strategy")

	_mwaaCmd.Flags().BoolVarP(&_mwaaCreateCliToken, "create-cli-token", "", false, "Create Cli Token")
	_mwaaCmd.Flags().BoolVarP(&_mwaaCreateEnvironment, "create-environment", "", false, "Create Environment")
	_mwaaCmd.Flags().BoolVarP(&_mwaaCreateWebLoginToken, "create-web-login-token", "", false, "Create Web Login Token")
	_mwaaCmd.Flags().BoolVarP(&_mwaaDeleteEnvironment, "delete-environment", "", false, "Delete Environment")
	_mwaaCmd.Flags().BoolVarP(&_mwaaGetEnvironment, "get-environment", "", false, "Get Environment")
	_mwaaCmd.Flags().BoolVarP(&_mwaaInvokeRestApi, "invoke-rest-api", "", false, "Invoke Rest API")
	_mwaaCmd.Flags().BoolVarP(&_mwaaListEnvironments, "list-environments", "", false, "List Environments")
	_mwaaCmd.Flags().BoolVarP(&_mwaaListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_mwaaCmd.Flags().BoolVarP(&_mwaaPublishMetrics, "publish-metrics", "", false, "Publish Metrics")
	_mwaaCmd.Flags().BoolVarP(&_mwaaTagResource, "tag-resource", "", false, "Tag Resource")
	_mwaaCmd.Flags().BoolVarP(&_mwaaUntagResource, "untag-resource", "", false, "Untag Resource")
	_mwaaCmd.Flags().BoolVarP(&_mwaaUpdateEnvironment, "update-environment", "", false, "Update Environment")

}
