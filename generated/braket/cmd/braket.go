package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/braket"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// braketCmd represents the braket command
var _braketCmd = &cobra.Command{
	Use:   "braket",
	Short: "AWS braket CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := braket.NewFromConfig(cfg)
		if _braketCancelJob {
			braket_CancelJob(cfg, client)
			return
		}
		if _braketCancelQuantumTask {
			braket_CancelQuantumTask(cfg, client)
			return
		}
		if _braketCreateJob {
			braket_CreateJob(cfg, client)
			return
		}
		if _braketCreateQuantumTask {
			braket_CreateQuantumTask(cfg, client)
			return
		}
		if _braketCreateSpendingLimit {
			braket_CreateSpendingLimit(cfg, client)
			return
		}
		if _braketDeleteSpendingLimit {
			braket_DeleteSpendingLimit(cfg, client)
			return
		}
		if _braketGetDevice {
			braket_GetDevice(cfg, client)
			return
		}
		if _braketGetJob {
			braket_GetJob(cfg, client)
			return
		}
		if _braketGetQuantumTask {
			braket_GetQuantumTask(cfg, client)
			return
		}
		if _braketListTagsForResource {
			braket_ListTagsForResource(cfg, client)
			return
		}
		if _braketSearchDevices {
			braket_SearchDevices(cfg, client)
			return
		}
		if _braketSearchJobs {
			braket_SearchJobs(cfg, client)
			return
		}
		if _braketSearchQuantumTasks {
			braket_SearchQuantumTasks(cfg, client)
			return
		}
		if _braketSearchSpendingLimits {
			braket_SearchSpendingLimits(cfg, client)
			return
		}
		if _braketTagResource {
			braket_TagResource(cfg, client)
			return
		}
		if _braketUntagResource {
			braket_UntagResource(cfg, client)
			return
		}
		if _braketUpdateSpendingLimit {
			braket_UpdateSpendingLimit(cfg, client)
			return
		}

	},
}

var (
	_braketCancelJob            bool
	_braketCancelQuantumTask    bool
	_braketCreateJob            bool
	_braketCreateQuantumTask    bool
	_braketCreateSpendingLimit  bool
	_braketDeleteSpendingLimit  bool
	_braketGetDevice            bool
	_braketGetJob               bool
	_braketGetQuantumTask       bool
	_braketListTagsForResource  bool
	_braketSearchDevices        bool
	_braketSearchJobs           bool
	_braketSearchQuantumTasks   bool
	_braketSearchSpendingLimits bool
	_braketTagResource          bool
	_braketUntagResource        bool
	_braketUpdateSpendingLimit  bool

	_braketAction                   string
	_braketAdditionalAttributeNames string
	_braketAlgorithmSpecification   string
	_braketAssociations             string
	_braketCheckpointConfig         string
	_braketClientToken              string
	_braketDeviceArn                string
	_braketDeviceConfig             string
	_braketDeviceParameters         string
	_braketExperimentalCapabilities string
	_braketFilters                  string
	_braketHyperParameters          string
	_braketInputDataConfig          string
	_braketInstanceConfig           string
	_braketJobArn                   string
	_braketJobName                  string
	_braketJobToken                 string
	_braketMaxResults               string
	_braketNextToken                string
	_braketOutputDataConfig         string
	_braketOutputS3Bucket           string
	_braketOutputS3KeyPrefix        string
	_braketQuantumTaskArn           string
	_braketResourceArn              string
	_braketRoleArn                  string
	_braketShots                    string
	_braketSpendingLimit            string
	_braketSpendingLimitArn         string
	_braketStoppingCondition        string
	_braketTagKeys                  []string
	_braketTags                     string
	_braketTimePeriod               string
)

// Cancels an Amazon Braket hybrid job.
func braket_CancelJob(cfg aws.Config, client *braket.Client) {
	input := &braket.CancelJobInput{
		// JobArn: *string, // Required
	}

	if len(_braketJobArn) > 0 {
		input.JobArn = aws.String(_braketJobArn)
	}

	if resp, err := client.CancelJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels the specified task.
func braket_CancelQuantumTask(cfg aws.Config, client *braket.Client) {
	input := &braket.CancelQuantumTaskInput{
		// ClientToken: *string, // Required
		// QuantumTaskArn: *string, // Required
	}

	if len(_braketClientToken) > 0 {
		input.ClientToken = aws.String(_braketClientToken)
	}
	if len(_braketQuantumTaskArn) > 0 {
		input.QuantumTaskArn = aws.String(_braketQuantumTaskArn)
	}

	if resp, err := client.CancelQuantumTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Braket hybrid job.
func braket_CreateJob(cfg aws.Config, client *braket.Client) {
	input := &braket.CreateJobInput{
		// AlgorithmSpecification: *types.AlgorithmSpecification, // Required
		// ClientToken: *string, // Required
		// DeviceConfig: *types.DeviceConfig, // Required
		// InstanceConfig: *types.InstanceConfig, // Required
		// JobName: *string, // Required
		// OutputDataConfig: *types.JobOutputDataConfig, // Required
		// RoleArn: *string, // Required
	}

	if len(_braketAlgorithmSpecification) > 0 {
		if err := assignInputField(input, "AlgorithmSpecification", _braketAlgorithmSpecification); err != nil {
			log.Errorf("invalid --algorithm-specification: %s", err.Error())
			return
		}
	}
	if len(_braketClientToken) > 0 {
		input.ClientToken = aws.String(_braketClientToken)
	}
	if len(_braketDeviceConfig) > 0 {
		if err := assignInputField(input, "DeviceConfig", _braketDeviceConfig); err != nil {
			log.Errorf("invalid --device-config: %s", err.Error())
			return
		}
	}
	if len(_braketInstanceConfig) > 0 {
		if err := assignInputField(input, "InstanceConfig", _braketInstanceConfig); err != nil {
			log.Errorf("invalid --instance-config: %s", err.Error())
			return
		}
	}
	if len(_braketJobName) > 0 {
		input.JobName = aws.String(_braketJobName)
	}
	if len(_braketOutputDataConfig) > 0 {
		if err := assignInputField(input, "OutputDataConfig", _braketOutputDataConfig); err != nil {
			log.Errorf("invalid --output-data-config: %s", err.Error())
			return
		}
	}
	if len(_braketRoleArn) > 0 {
		input.RoleArn = aws.String(_braketRoleArn)
	}
	if len(_braketAssociations) > 0 {
		if err := assignInputField(input, "Associations", _braketAssociations); err != nil {
			log.Errorf("invalid --associations: %s", err.Error())
			return
		}
	}
	if len(_braketCheckpointConfig) > 0 {
		if err := assignInputField(input, "CheckpointConfig", _braketCheckpointConfig); err != nil {
			log.Errorf("invalid --checkpoint-config: %s", err.Error())
			return
		}
	}
	if len(_braketHyperParameters) > 0 {
		if err := assignInputField(input, "HyperParameters", _braketHyperParameters); err != nil {
			log.Errorf("invalid --hyper-parameters: %s", err.Error())
			return
		}
	}
	if len(_braketInputDataConfig) > 0 {
		if err := assignInputField(input, "InputDataConfig", _braketInputDataConfig); err != nil {
			log.Errorf("invalid --input-data-config: %s", err.Error())
			return
		}
	}
	if len(_braketStoppingCondition) > 0 {
		if err := assignInputField(input, "StoppingCondition", _braketStoppingCondition); err != nil {
			log.Errorf("invalid --stopping-condition: %s", err.Error())
			return
		}
	}
	if len(_braketTags) > 0 {
		if err := assignInputField(input, "Tags", _braketTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a quantum task.
func braket_CreateQuantumTask(cfg aws.Config, client *braket.Client) {
	input := &braket.CreateQuantumTaskInput{
		// Action: *string, // Required
		// ClientToken: *string, // Required
		// DeviceArn: *string, // Required
		// OutputS3Bucket: *string, // Required
		// OutputS3KeyPrefix: *string, // Required
		// Shots: *int64, // Required
	}

	if len(_braketAction) > 0 {
		input.Action = aws.String(_braketAction)
	}
	if len(_braketClientToken) > 0 {
		input.ClientToken = aws.String(_braketClientToken)
	}
	if len(_braketDeviceArn) > 0 {
		input.DeviceArn = aws.String(_braketDeviceArn)
	}
	if len(_braketOutputS3Bucket) > 0 {
		input.OutputS3Bucket = aws.String(_braketOutputS3Bucket)
	}
	if len(_braketOutputS3KeyPrefix) > 0 {
		input.OutputS3KeyPrefix = aws.String(_braketOutputS3KeyPrefix)
	}
	if len(_braketShots) > 0 {
		if err := assignInputField(input, "Shots", _braketShots); err != nil {
			log.Errorf("invalid --shots: %s", err.Error())
			return
		}
	}
	if len(_braketAssociations) > 0 {
		if err := assignInputField(input, "Associations", _braketAssociations); err != nil {
			log.Errorf("invalid --associations: %s", err.Error())
			return
		}
	}
	if len(_braketDeviceParameters) > 0 {
		input.DeviceParameters = aws.String(_braketDeviceParameters)
	}
	if len(_braketExperimentalCapabilities) > 0 {
		if err := assignInputField(input, "ExperimentalCapabilities", _braketExperimentalCapabilities); err != nil {
			log.Errorf("invalid --experimental-capabilities: %s", err.Error())
			return
		}
	}
	if len(_braketJobToken) > 0 {
		input.JobToken = aws.String(_braketJobToken)
	}
	if len(_braketTags) > 0 {
		if err := assignInputField(input, "Tags", _braketTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateQuantumTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a spending limit for a specified quantum device. Spending limits help
// you control costs by setting maximum amounts that can be spent on quantum
// computing tasks within a specified time period. Simulators do not support
// spending limits.
func braket_CreateSpendingLimit(cfg aws.Config, client *braket.Client) {
	input := &braket.CreateSpendingLimitInput{
		// ClientToken: *string, // Required
		// DeviceArn: *string, // Required
		// SpendingLimit: *string, // Required
	}

	if len(_braketClientToken) > 0 {
		input.ClientToken = aws.String(_braketClientToken)
	}
	if len(_braketDeviceArn) > 0 {
		input.DeviceArn = aws.String(_braketDeviceArn)
	}
	if len(_braketSpendingLimit) > 0 {
		input.SpendingLimit = aws.String(_braketSpendingLimit)
	}
	if len(_braketTags) > 0 {
		if err := assignInputField(input, "Tags", _braketTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_braketTimePeriod) > 0 {
		if err := assignInputField(input, "TimePeriod", _braketTimePeriod); err != nil {
			log.Errorf("invalid --time-period: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSpendingLimit(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing spending limit. This operation permanently removes the
// spending limit and cannot be undone. After deletion, the associated device
// becomes unrestricted for spending.
func braket_DeleteSpendingLimit(cfg aws.Config, client *braket.Client) {
	input := &braket.DeleteSpendingLimitInput{
		// SpendingLimitArn: *string, // Required
	}

	if len(_braketSpendingLimitArn) > 0 {
		input.SpendingLimitArn = aws.String(_braketSpendingLimitArn)
	}

	if resp, err := client.DeleteSpendingLimit(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the devices available in Amazon Braket.
// For backwards compatibility with older versions of BraketSchemas, OpenQASM
// information is omitted from GetDevice API calls. To get this information the
// user-agent needs to present a recent version of the BraketSchemas (1.8.0 or
// later). The Braket SDK automatically reports this for you. If you do not see
// OpenQASM results in the GetDevice response when using a Braket SDK, you may need
// to set AWS_EXECUTION_ENV environment variable to configure user-agent. See the
// code examples provided below for how to do this for the AWS CLI, Boto3, and the
// Go, Java, and JavaScript/TypeScript SDKs.
func braket_GetDevice(cfg aws.Config, client *braket.Client) {
	input := &braket.GetDeviceInput{
		// DeviceArn: *string, // Required
	}

	if len(_braketDeviceArn) > 0 {
		input.DeviceArn = aws.String(_braketDeviceArn)
	}

	if resp, err := client.GetDevice(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the specified Amazon Braket hybrid job.
func braket_GetJob(cfg aws.Config, client *braket.Client) {
	input := &braket.GetJobInput{
		// JobArn: *string, // Required
	}

	if len(_braketJobArn) > 0 {
		input.JobArn = aws.String(_braketJobArn)
	}
	if len(_braketAdditionalAttributeNames) > 0 {
		if err := assignInputField(input, "AdditionalAttributeNames", _braketAdditionalAttributeNames); err != nil {
			log.Errorf("invalid --additional-attribute-names: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the specified quantum task.
func braket_GetQuantumTask(cfg aws.Config, client *braket.Client) {
	input := &braket.GetQuantumTaskInput{
		// QuantumTaskArn: *string, // Required
	}

	if len(_braketQuantumTaskArn) > 0 {
		input.QuantumTaskArn = aws.String(_braketQuantumTaskArn)
	}
	if len(_braketAdditionalAttributeNames) > 0 {
		if err := assignInputField(input, "AdditionalAttributeNames", _braketAdditionalAttributeNames); err != nil {
			log.Errorf("invalid --additional-attribute-names: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetQuantumTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Shows the tags associated with this resource.
func braket_ListTagsForResource(cfg aws.Config, client *braket.Client) {
	input := &braket.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_braketResourceArn) > 0 {
		input.ResourceArn = aws.String(_braketResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Searches for devices using the specified filters.
func braket_SearchDevices(cfg aws.Config, client *braket.Client) {
	input := &braket.SearchDevicesInput{
		// Filters: []types.SearchDevicesFilter, // Required
	}

	if len(_braketFilters) > 0 {
		if err := assignInputField(input, "Filters", _braketFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_braketMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _braketMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_braketNextToken) > 0 {
		input.NextToken = aws.String(_braketNextToken)
	}

	if disablePaginator() {
		if resp, err := client.SearchDevices(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*braket.SearchDevicesOutput
	p := braket.NewSearchDevicesPaginator(client, input)
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

// Searches for Amazon Braket hybrid jobs that match the specified filter values.
func braket_SearchJobs(cfg aws.Config, client *braket.Client) {
	input := &braket.SearchJobsInput{
		// Filters: []types.SearchJobsFilter, // Required
	}

	if len(_braketFilters) > 0 {
		if err := assignInputField(input, "Filters", _braketFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_braketMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _braketMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_braketNextToken) > 0 {
		input.NextToken = aws.String(_braketNextToken)
	}

	if disablePaginator() {
		if resp, err := client.SearchJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*braket.SearchJobsOutput
	p := braket.NewSearchJobsPaginator(client, input)
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

// Searches for tasks that match the specified filter values.
func braket_SearchQuantumTasks(cfg aws.Config, client *braket.Client) {
	input := &braket.SearchQuantumTasksInput{
		// Filters: []types.SearchQuantumTasksFilter, // Required
	}

	if len(_braketFilters) > 0 {
		if err := assignInputField(input, "Filters", _braketFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_braketMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _braketMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_braketNextToken) > 0 {
		input.NextToken = aws.String(_braketNextToken)
	}

	if disablePaginator() {
		if resp, err := client.SearchQuantumTasks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*braket.SearchQuantumTasksOutput
	p := braket.NewSearchQuantumTasksPaginator(client, input)
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

// Searches and lists spending limits based on specified filters. This operation
// supports pagination and allows filtering by various criteria to find specific
// spending limits. We recommend using pagination to ensure that the operation
// returns quickly and successfully.
func braket_SearchSpendingLimits(cfg aws.Config, client *braket.Client) {
	input := &braket.SearchSpendingLimitsInput{}

	if len(_braketFilters) > 0 {
		if err := assignInputField(input, "Filters", _braketFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_braketMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _braketMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_braketNextToken) > 0 {
		input.NextToken = aws.String(_braketNextToken)
	}

	if disablePaginator() {
		if resp, err := client.SearchSpendingLimits(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*braket.SearchSpendingLimitsOutput
	p := braket.NewSearchSpendingLimitsPaginator(client, input)
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

// Add a tag to the specified resource.
func braket_TagResource(cfg aws.Config, client *braket.Client) {
	input := &braket.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_braketResourceArn) > 0 {
		input.ResourceArn = aws.String(_braketResourceArn)
	}
	if len(_braketTags) > 0 {
		if err := assignInputField(input, "Tags", _braketTags); err != nil {
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

// Remove tags from a resource.
func braket_UntagResource(cfg aws.Config, client *braket.Client) {
	input := &braket.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_braketResourceArn) > 0 {
		input.ResourceArn = aws.String(_braketResourceArn)
	}
	if len(_braketTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _braketTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing spending limit. You can modify the spending amount or time
// period. Changes take effect immediately.
func braket_UpdateSpendingLimit(cfg aws.Config, client *braket.Client) {
	input := &braket.UpdateSpendingLimitInput{
		// ClientToken: *string, // Required
		// SpendingLimitArn: *string, // Required
	}

	if len(_braketClientToken) > 0 {
		input.ClientToken = aws.String(_braketClientToken)
	}
	if len(_braketSpendingLimitArn) > 0 {
		input.SpendingLimitArn = aws.String(_braketSpendingLimitArn)
	}
	if len(_braketSpendingLimit) > 0 {
		input.SpendingLimit = aws.String(_braketSpendingLimit)
	}
	if len(_braketTimePeriod) > 0 {
		if err := assignInputField(input, "TimePeriod", _braketTimePeriod); err != nil {
			log.Errorf("invalid --time-period: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSpendingLimit(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_braketCmd)
	_braketCmd.Flags().SortFlags = false

	_braketCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_braketCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_braketCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_braketCmd.Flags().StringVarP(&_braketAction, "action", "", "", "Action")
	_braketCmd.Flags().StringVarP(&_braketAdditionalAttributeNames, "additional-attribute-names", "", "", "Additional Attribute Names")
	_braketCmd.Flags().StringVarP(&_braketAlgorithmSpecification, "algorithm-specification", "", "", "Algorithm Specification")
	_braketCmd.Flags().StringVarP(&_braketAssociations, "associations", "", "", "Associations")
	_braketCmd.Flags().StringVarP(&_braketCheckpointConfig, "checkpoint-config", "", "", "Checkpoint Config")
	_braketCmd.Flags().StringVarP(&_braketClientToken, "client-token", "", "", "Client Token")
	_braketCmd.Flags().StringVarP(&_braketDeviceArn, "device-arn", "", "", "Device ARN")
	_braketCmd.Flags().StringVarP(&_braketDeviceConfig, "device-config", "", "", "Device Config")
	_braketCmd.Flags().StringVarP(&_braketDeviceParameters, "device-parameters", "", "", "Device Parameters")
	_braketCmd.Flags().StringVarP(&_braketExperimentalCapabilities, "experimental-capabilities", "", "", "Experimental Capabilities")
	_braketCmd.Flags().StringVarP(&_braketFilters, "filters", "", "", "Filters")
	_braketCmd.Flags().StringVarP(&_braketHyperParameters, "hyper-parameters", "", "", "Hyper Parameters")
	_braketCmd.Flags().StringVarP(&_braketInputDataConfig, "input-data-config", "", "", "Input Data Config")
	_braketCmd.Flags().StringVarP(&_braketInstanceConfig, "instance-config", "", "", "Instance Config")
	_braketCmd.Flags().StringVarP(&_braketJobArn, "job-arn", "", "", "Job ARN")
	_braketCmd.Flags().StringVarP(&_braketJobName, "job-name", "", "", "Job Name")
	_braketCmd.Flags().StringVarP(&_braketJobToken, "job-token", "", "", "Job Token")
	_braketCmd.Flags().StringVarP(&_braketMaxResults, "max-results", "", "", "Max Results")
	_braketCmd.Flags().StringVarP(&_braketNextToken, "next-token", "", "", "Next Token")
	_braketCmd.Flags().StringVarP(&_braketOutputDataConfig, "output-data-config", "", "", "Output Data Config")
	_braketCmd.Flags().StringVarP(&_braketOutputS3Bucket, "output-s3-bucket", "", "", "Output S3 Bucket")
	_braketCmd.Flags().StringVarP(&_braketOutputS3KeyPrefix, "output-s3-key-prefix", "", "", "Output S3 Key Prefix")
	_braketCmd.Flags().StringVarP(&_braketQuantumTaskArn, "quantum-task-arn", "", "", "Quantum Task ARN")
	_braketCmd.Flags().StringVarP(&_braketResourceArn, "resource-arn", "", "", "Resource ARN")
	_braketCmd.Flags().StringVarP(&_braketRoleArn, "role-arn", "", "", "Role ARN")
	_braketCmd.Flags().StringVarP(&_braketShots, "shots", "", "", "Shots")
	_braketCmd.Flags().StringVarP(&_braketSpendingLimit, "spending-limit", "", "", "Spending Limit")
	_braketCmd.Flags().StringVarP(&_braketSpendingLimitArn, "spending-limit-arn", "", "", "Spending Limit ARN")
	_braketCmd.Flags().StringVarP(&_braketStoppingCondition, "stopping-condition", "", "", "Stopping Condition")
	_braketCmd.Flags().StringSliceVarP(&_braketTagKeys, "tag-keys", "", nil, "Tag Keys")
	_braketCmd.Flags().StringVarP(&_braketTags, "tags", "", "", "Tags")
	_braketCmd.Flags().StringVarP(&_braketTimePeriod, "time-period", "", "", "Time Period")

	_braketCmd.Flags().BoolVarP(&_braketCancelJob, "cancel-job", "", false, "Cancel Job")
	_braketCmd.Flags().BoolVarP(&_braketCancelQuantumTask, "cancel-quantum-task", "", false, "Cancel Quantum Task")
	_braketCmd.Flags().BoolVarP(&_braketCreateJob, "create-job", "", false, "Create Job")
	_braketCmd.Flags().BoolVarP(&_braketCreateQuantumTask, "create-quantum-task", "", false, "Create Quantum Task")
	_braketCmd.Flags().BoolVarP(&_braketCreateSpendingLimit, "create-spending-limit", "", false, "Create Spending Limit")
	_braketCmd.Flags().BoolVarP(&_braketDeleteSpendingLimit, "delete-spending-limit", "", false, "Delete Spending Limit")
	_braketCmd.Flags().BoolVarP(&_braketGetDevice, "get-device", "", false, "Get Device")
	_braketCmd.Flags().BoolVarP(&_braketGetJob, "get-job", "", false, "Get Job")
	_braketCmd.Flags().BoolVarP(&_braketGetQuantumTask, "get-quantum-task", "", false, "Get Quantum Task")
	_braketCmd.Flags().BoolVarP(&_braketListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_braketCmd.Flags().BoolVarP(&_braketSearchDevices, "search-devices", "", false, "Search Devices")
	_braketCmd.Flags().BoolVarP(&_braketSearchJobs, "search-jobs", "", false, "Search Jobs")
	_braketCmd.Flags().BoolVarP(&_braketSearchQuantumTasks, "search-quantum-tasks", "", false, "Search Quantum Tasks")
	_braketCmd.Flags().BoolVarP(&_braketSearchSpendingLimits, "search-spending-limits", "", false, "Search Spending Limits")
	_braketCmd.Flags().BoolVarP(&_braketTagResource, "tag-resource", "", false, "Tag Resource")
	_braketCmd.Flags().BoolVarP(&_braketUntagResource, "untag-resource", "", false, "Untag Resource")
	_braketCmd.Flags().BoolVarP(&_braketUpdateSpendingLimit, "update-spending-limit", "", false, "Update Spending Limit")

}
