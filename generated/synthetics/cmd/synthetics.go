package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/synthetics"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// syntheticsCmd represents the synthetics command
var _syntheticsCmd = &cobra.Command{
	Use:   "synthetics",
	Short: "AWS synthetics CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := synthetics.NewFromConfig(cfg)
		if _syntheticsAssociateResource {
			synthetics_AssociateResource(cfg, client)
			return
		}
		if _syntheticsCreateCanary {
			synthetics_CreateCanary(cfg, client)
			return
		}
		if _syntheticsCreateGroup {
			synthetics_CreateGroup(cfg, client)
			return
		}
		if _syntheticsDeleteCanary {
			synthetics_DeleteCanary(cfg, client)
			return
		}
		if _syntheticsDeleteGroup {
			synthetics_DeleteGroup(cfg, client)
			return
		}
		if _syntheticsDescribeCanaries {
			synthetics_DescribeCanaries(cfg, client)
			return
		}
		if _syntheticsDescribeCanariesLastRun {
			synthetics_DescribeCanariesLastRun(cfg, client)
			return
		}
		if _syntheticsDescribeRuntimeVersions {
			synthetics_DescribeRuntimeVersions(cfg, client)
			return
		}
		if _syntheticsDisassociateResource {
			synthetics_DisassociateResource(cfg, client)
			return
		}
		if _syntheticsGetCanary {
			synthetics_GetCanary(cfg, client)
			return
		}
		if _syntheticsGetCanaryRuns {
			synthetics_GetCanaryRuns(cfg, client)
			return
		}
		if _syntheticsGetGroup {
			synthetics_GetGroup(cfg, client)
			return
		}
		if _syntheticsListAssociatedGroups {
			synthetics_ListAssociatedGroups(cfg, client)
			return
		}
		if _syntheticsListGroupResources {
			synthetics_ListGroupResources(cfg, client)
			return
		}
		if _syntheticsListGroups {
			synthetics_ListGroups(cfg, client)
			return
		}
		if _syntheticsListTagsForResource {
			synthetics_ListTagsForResource(cfg, client)
			return
		}
		if _syntheticsStartCanary {
			synthetics_StartCanary(cfg, client)
			return
		}
		if _syntheticsStartCanaryDryRun {
			synthetics_StartCanaryDryRun(cfg, client)
			return
		}
		if _syntheticsStopCanary {
			synthetics_StopCanary(cfg, client)
			return
		}
		if _syntheticsTagResource {
			synthetics_TagResource(cfg, client)
			return
		}
		if _syntheticsUntagResource {
			synthetics_UntagResource(cfg, client)
			return
		}
		if _syntheticsUpdateCanary {
			synthetics_UpdateCanary(cfg, client)
			return
		}

	},
}

var (
	_syntheticsAssociateResource       bool
	_syntheticsCreateCanary            bool
	_syntheticsCreateGroup             bool
	_syntheticsDeleteCanary            bool
	_syntheticsDeleteGroup             bool
	_syntheticsDescribeCanaries        bool
	_syntheticsDescribeCanariesLastRun bool
	_syntheticsDescribeRuntimeVersions bool
	_syntheticsDisassociateResource    bool
	_syntheticsGetCanary               bool
	_syntheticsGetCanaryRuns           bool
	_syntheticsGetGroup                bool
	_syntheticsListAssociatedGroups    bool
	_syntheticsListGroupResources      bool
	_syntheticsListGroups              bool
	_syntheticsListTagsForResource     bool
	_syntheticsStartCanary             bool
	_syntheticsStartCanaryDryRun       bool
	_syntheticsStopCanary              bool
	_syntheticsTagResource             bool
	_syntheticsUntagResource           bool
	_syntheticsUpdateCanary            bool

	_syntheticsArtifactConfig               string
	_syntheticsArtifactS3Location           string
	_syntheticsBrowserConfigs               string
	_syntheticsBrowserType                  string
	_syntheticsCode                         string
	_syntheticsDeleteLambda                 string
	_syntheticsDryRunId                     string
	_syntheticsExecutionRoleArn             string
	_syntheticsFailureRetentionPeriodInDays string
	_syntheticsGroupIdentifier              string
	_syntheticsMaxResults                   string
	_syntheticsName                         string
	_syntheticsNames                        []string
	_syntheticsNextToken                    string
	_syntheticsProvisionedResourceCleanup   string
	_syntheticsResourceArn                  string
	_syntheticsResourcesToReplicateTags     string
	_syntheticsRunConfig                    string
	_syntheticsRunType                      string
	_syntheticsRuntimeVersion               string
	_syntheticsSchedule                     string
	_syntheticsSuccessRetentionPeriodInDays string
	_syntheticsTagKeys                      []string
	_syntheticsTags                         string
	_syntheticsVisualReference              string
	_syntheticsVisualReferences             string
	_syntheticsVpcConfig                    string
)

// Associates a canary with a group. Using groups can help you with managing and
// automating your canaries, and you can also view aggregated run results and
// statistics for all canaries in a group.
//
// You must run this operation in the Region where the canary exists.
func synthetics_AssociateResource(cfg aws.Config, client *synthetics.Client) {
	input := &synthetics.AssociateResourceInput{
		// GroupIdentifier: *string, // Required
		// ResourceArn: *string, // Required
	}

	if len(_syntheticsGroupIdentifier) > 0 {
		input.GroupIdentifier = aws.String(_syntheticsGroupIdentifier)
	}
	if len(_syntheticsResourceArn) > 0 {
		input.ResourceArn = aws.String(_syntheticsResourceArn)
	}

	if resp, err := client.AssociateResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a canary. Canaries are scripts that monitor your endpoints and APIs
// from the outside-in. Canaries help you check the availability and latency of
// your web services and troubleshoot anomalies by investigating load time data,
// screenshots of the UI, logs, and metrics. You can set up a canary to run
// continuously or just once.
//
// Do not use CreateCanary to modify an existing canary. Use [UpdateCanary] instead.
//
// To create canaries, you must have the CloudWatchSyntheticsFullAccess policy. If
// you are creating a new IAM role for the canary, you also need the iam:CreateRole
// , iam:CreatePolicy and iam:AttachRolePolicy permissions. For more information,
// see [Necessary Roles and Permissions].
//
// Do not include secrets or proprietary information in your canary names. The
// canary name makes up part of the Amazon Resource Name (ARN) for the canary, and
// the ARN is included in outbound calls over the internet. For more information,
// see [Security Considerations for Synthetics Canaries].
//
// [UpdateCanary]: https://docs.aws.amazon.com/AmazonSynthetics/latest/APIReference/API_UpdateCanary.html
// [Necessary Roles and Permissions]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Roles
// [Security Considerations for Synthetics Canaries]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/servicelens_canaries_security.html
func synthetics_CreateCanary(cfg aws.Config, client *synthetics.Client) {
	input := &synthetics.CreateCanaryInput{
		// ArtifactS3Location: *string, // Required
		// Code: *types.CanaryCodeInput, // Required
		// ExecutionRoleArn: *string, // Required
		// Name: *string, // Required
		// RuntimeVersion: *string, // Required
		// Schedule: *types.CanaryScheduleInput, // Required
	}

	if len(_syntheticsArtifactS3Location) > 0 {
		input.ArtifactS3Location = aws.String(_syntheticsArtifactS3Location)
	}
	if len(_syntheticsCode) > 0 {
		if err := assignInputField(input, "Code", _syntheticsCode); err != nil {
			log.Errorf("invalid --code: %s", err.Error())
			return
		}
	}
	if len(_syntheticsExecutionRoleArn) > 0 {
		input.ExecutionRoleArn = aws.String(_syntheticsExecutionRoleArn)
	}
	if len(_syntheticsName) > 0 {
		input.Name = aws.String(_syntheticsName)
	}
	if len(_syntheticsRuntimeVersion) > 0 {
		input.RuntimeVersion = aws.String(_syntheticsRuntimeVersion)
	}
	if len(_syntheticsSchedule) > 0 {
		if err := assignInputField(input, "Schedule", _syntheticsSchedule); err != nil {
			log.Errorf("invalid --schedule: %s", err.Error())
			return
		}
	}
	if len(_syntheticsArtifactConfig) > 0 {
		if err := assignInputField(input, "ArtifactConfig", _syntheticsArtifactConfig); err != nil {
			log.Errorf("invalid --artifact-config: %s", err.Error())
			return
		}
	}
	if len(_syntheticsBrowserConfigs) > 0 {
		if err := assignInputField(input, "BrowserConfigs", _syntheticsBrowserConfigs); err != nil {
			log.Errorf("invalid --browser-configs: %s", err.Error())
			return
		}
	}
	if len(_syntheticsFailureRetentionPeriodInDays) > 0 {
		if err := assignInputField(input, "FailureRetentionPeriodInDays", _syntheticsFailureRetentionPeriodInDays); err != nil {
			log.Errorf("invalid --failure-retention-period-in-days: %s", err.Error())
			return
		}
	}
	if len(_syntheticsProvisionedResourceCleanup) > 0 {
		if err := assignInputField(input, "ProvisionedResourceCleanup", _syntheticsProvisionedResourceCleanup); err != nil {
			log.Errorf("invalid --provisioned-resource-cleanup: %s", err.Error())
			return
		}
	}
	if len(_syntheticsResourcesToReplicateTags) > 0 {
		if err := assignInputField(input, "ResourcesToReplicateTags", _syntheticsResourcesToReplicateTags); err != nil {
			log.Errorf("invalid --resources-to-replicate-tags: %s", err.Error())
			return
		}
	}
	if len(_syntheticsRunConfig) > 0 {
		if err := assignInputField(input, "RunConfig", _syntheticsRunConfig); err != nil {
			log.Errorf("invalid --run-config: %s", err.Error())
			return
		}
	}
	if len(_syntheticsSuccessRetentionPeriodInDays) > 0 {
		if err := assignInputField(input, "SuccessRetentionPeriodInDays", _syntheticsSuccessRetentionPeriodInDays); err != nil {
			log.Errorf("invalid --success-retention-period-in-days: %s", err.Error())
			return
		}
	}
	if len(_syntheticsTags) > 0 {
		if err := assignInputField(input, "Tags", _syntheticsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_syntheticsVpcConfig) > 0 {
		if err := assignInputField(input, "VpcConfig", _syntheticsVpcConfig); err != nil {
			log.Errorf("invalid --vpc-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCanary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a group which you can use to associate canaries with each other,
// including cross-Region canaries. Using groups can help you with managing and
// automating your canaries, and you can also view aggregated run results and
// statistics for all canaries in a group.
//
// Groups are global resources. When you create a group, it is replicated across
// Amazon Web Services Regions, and you can view it and add canaries to it from any
// Region. Although the group ARN format reflects the Region name where it was
// created, a group is not constrained to any Region. This means that you can put
// canaries from multiple Regions into the same group, and then use that group to
// view and manage all of those canaries in a single view.
//
// Groups are supported in all Regions except the Regions that are disabled by
// default. For more information about these Regions, see [Enabling a Region].
//
// Each group can contain as many as 10 canaries. You can have as many as 20
// groups in your account. Any single canary can be a member of up to 10 groups.
//
// [Enabling a Region]: https://docs.aws.amazon.com/general/latest/gr/rande-manage.html#rande-manage-enable
func synthetics_CreateGroup(cfg aws.Config, client *synthetics.Client) {
	input := &synthetics.CreateGroupInput{
		// Name: *string, // Required
	}

	if len(_syntheticsName) > 0 {
		input.Name = aws.String(_syntheticsName)
	}
	if len(_syntheticsTags) > 0 {
		if err := assignInputField(input, "Tags", _syntheticsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Permanently deletes the specified canary.
// If the canary's ProvisionedResourceCleanup field is set to AUTOMATIC or you
// specify DeleteLambda in this operation as true , CloudWatch Synthetics also
// deletes the Lambda functions and layers that are used by the canary.
//
// Other resources used and created by the canary are not automatically deleted.
// After you delete a canary, you should also delete the following:
//
// - The CloudWatch alarms created for this canary. These alarms have a name of
// Synthetics-Alarm-first-198-characters-of-canary-name-canaryId-alarm number
//
// - Amazon S3 objects and buckets, such as the canary's artifact location.
//
// - IAM roles created for the canary. If they were created in the console,
// these roles have the name
// role/service-role/CloudWatchSyntheticsRole-First-21-Characters-of-CanaryName
//
// - CloudWatch Logs log groups created for the canary. These logs groups have
// the name /aws/lambda/cwsyn-First-21-Characters-of-CanaryName
//
// Before you delete a canary, you might want to use GetCanary to display the
// information about this canary. Make note of the information returned by this
// operation so that you can delete these resources after you delete the canary.
func synthetics_DeleteCanary(cfg aws.Config, client *synthetics.Client) {
	input := &synthetics.DeleteCanaryInput{
		// Name: *string, // Required
	}

	if len(_syntheticsName) > 0 {
		input.Name = aws.String(_syntheticsName)
	}
	if len(_syntheticsDeleteLambda) > 0 {
		if err := assignInputField(input, "DeleteLambda", _syntheticsDeleteLambda); err != nil {
			log.Errorf("invalid --delete-lambda: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteCanary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a group. The group doesn't need to be empty to be deleted. If there are
// canaries in the group, they are not deleted when you delete the group.
//
// Groups are a global resource that appear in all Regions, but the request to
// delete a group must be made from its home Region. You can find the home Region
// of a group within its ARN.
func synthetics_DeleteGroup(cfg aws.Config, client *synthetics.Client) {
	input := &synthetics.DeleteGroupInput{
		// GroupIdentifier: *string, // Required
	}

	if len(_syntheticsGroupIdentifier) > 0 {
		input.GroupIdentifier = aws.String(_syntheticsGroupIdentifier)
	}

	if resp, err := client.DeleteGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation returns a list of the canaries in your account, along with full
// details about each canary.
//
// This operation supports resource-level authorization using an IAM policy and
// the Names parameter. If you specify the Names parameter, the operation is
// successful only if you have authorization to view all the canaries that you
// specify in your request. If you do not have permission to view any of the
// canaries, the request fails with a 403 response.
//
// You are required to use the Names parameter if you are logged on to a user or
// role that has an IAM policy that restricts which canaries that you are allowed
// to view. For more information, see [Limiting a user to viewing specific canaries].
//
// [Limiting a user to viewing specific canaries]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Restricted.html
func synthetics_DescribeCanaries(cfg aws.Config, client *synthetics.Client) {
	input := &synthetics.DescribeCanariesInput{}

	if len(_syntheticsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _syntheticsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_syntheticsNames) > 0 {
		input.Names = append([]string(nil), _syntheticsNames...)
	}
	if len(_syntheticsNextToken) > 0 {
		input.NextToken = aws.String(_syntheticsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeCanaries(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*synthetics.DescribeCanariesOutput
	p := synthetics.NewDescribeCanariesPaginator(client, input)
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

// Use this operation to see information from the most recent run of each canary
// that you have created.
//
// This operation supports resource-level authorization using an IAM policy and
// the Names parameter. If you specify the Names parameter, the operation is
// successful only if you have authorization to view all the canaries that you
// specify in your request. If you do not have permission to view any of the
// canaries, the request fails with a 403 response.
//
// You are required to use the Names parameter if you are logged on to a user or
// role that has an IAM policy that restricts which canaries that you are allowed
// to view. For more information, see [Limiting a user to viewing specific canaries].
//
// [Limiting a user to viewing specific canaries]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Restricted.html
func synthetics_DescribeCanariesLastRun(cfg aws.Config, client *synthetics.Client) {
	input := &synthetics.DescribeCanariesLastRunInput{}

	if len(_syntheticsBrowserType) > 0 {
		if err := assignInputField(input, "BrowserType", _syntheticsBrowserType); err != nil {
			log.Errorf("invalid --browser-type: %s", err.Error())
			return
		}
	}
	if len(_syntheticsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _syntheticsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_syntheticsNames) > 0 {
		input.Names = append([]string(nil), _syntheticsNames...)
	}
	if len(_syntheticsNextToken) > 0 {
		input.NextToken = aws.String(_syntheticsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeCanariesLastRun(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*synthetics.DescribeCanariesLastRunOutput
	p := synthetics.NewDescribeCanariesLastRunPaginator(client, input)
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

// Returns a list of Synthetics canary runtime versions. For more information, see [Canary Runtime Versions]
// .
//
// [Canary Runtime Versions]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Library.html
func synthetics_DescribeRuntimeVersions(cfg aws.Config, client *synthetics.Client) {
	input := &synthetics.DescribeRuntimeVersionsInput{}

	if len(_syntheticsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _syntheticsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_syntheticsNextToken) > 0 {
		input.NextToken = aws.String(_syntheticsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeRuntimeVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*synthetics.DescribeRuntimeVersionsOutput
	p := synthetics.NewDescribeRuntimeVersionsPaginator(client, input)
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

// Removes a canary from a group. You must run this operation in the Region where
// the canary exists.
func synthetics_DisassociateResource(cfg aws.Config, client *synthetics.Client) {
	input := &synthetics.DisassociateResourceInput{
		// GroupIdentifier: *string, // Required
		// ResourceArn: *string, // Required
	}

	if len(_syntheticsGroupIdentifier) > 0 {
		input.GroupIdentifier = aws.String(_syntheticsGroupIdentifier)
	}
	if len(_syntheticsResourceArn) > 0 {
		input.ResourceArn = aws.String(_syntheticsResourceArn)
	}

	if resp, err := client.DisassociateResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves complete information about one canary. You must specify the name of
// the canary that you want. To get a list of canaries and their names, use [DescribeCanaries].
//
// [DescribeCanaries]: https://docs.aws.amazon.com/AmazonSynthetics/latest/APIReference/API_DescribeCanaries.html
func synthetics_GetCanary(cfg aws.Config, client *synthetics.Client) {
	input := &synthetics.GetCanaryInput{
		// Name: *string, // Required
	}

	if len(_syntheticsName) > 0 {
		input.Name = aws.String(_syntheticsName)
	}
	if len(_syntheticsDryRunId) > 0 {
		input.DryRunId = aws.String(_syntheticsDryRunId)
	}

	if resp, err := client.GetCanary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of runs for a specified canary.
func synthetics_GetCanaryRuns(cfg aws.Config, client *synthetics.Client) {
	input := &synthetics.GetCanaryRunsInput{
		// Name: *string, // Required
	}

	if len(_syntheticsName) > 0 {
		input.Name = aws.String(_syntheticsName)
	}
	if len(_syntheticsDryRunId) > 0 {
		input.DryRunId = aws.String(_syntheticsDryRunId)
	}
	if len(_syntheticsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _syntheticsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_syntheticsNextToken) > 0 {
		input.NextToken = aws.String(_syntheticsNextToken)
	}
	if len(_syntheticsRunType) > 0 {
		if err := assignInputField(input, "RunType", _syntheticsRunType); err != nil {
			log.Errorf("invalid --run-type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetCanaryRuns(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*synthetics.GetCanaryRunsOutput
	p := synthetics.NewGetCanaryRunsPaginator(client, input)
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

// Returns information about one group. Groups are a global resource, so you can
// use this operation from any Region.
func synthetics_GetGroup(cfg aws.Config, client *synthetics.Client) {
	input := &synthetics.GetGroupInput{
		// GroupIdentifier: *string, // Required
	}

	if len(_syntheticsGroupIdentifier) > 0 {
		input.GroupIdentifier = aws.String(_syntheticsGroupIdentifier)
	}

	if resp, err := client.GetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of the groups that the specified canary is associated with. The
// canary that you specify must be in the current Region.
func synthetics_ListAssociatedGroups(cfg aws.Config, client *synthetics.Client) {
	input := &synthetics.ListAssociatedGroupsInput{
		// ResourceArn: *string, // Required
	}

	if len(_syntheticsResourceArn) > 0 {
		input.ResourceArn = aws.String(_syntheticsResourceArn)
	}
	if len(_syntheticsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _syntheticsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_syntheticsNextToken) > 0 {
		input.NextToken = aws.String(_syntheticsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAssociatedGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*synthetics.ListAssociatedGroupsOutput
	p := synthetics.NewListAssociatedGroupsPaginator(client, input)
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

// This operation returns a list of the ARNs of the canaries that are associated
// with the specified group.
func synthetics_ListGroupResources(cfg aws.Config, client *synthetics.Client) {
	input := &synthetics.ListGroupResourcesInput{
		// GroupIdentifier: *string, // Required
	}

	if len(_syntheticsGroupIdentifier) > 0 {
		input.GroupIdentifier = aws.String(_syntheticsGroupIdentifier)
	}
	if len(_syntheticsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _syntheticsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_syntheticsNextToken) > 0 {
		input.NextToken = aws.String(_syntheticsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListGroupResources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*synthetics.ListGroupResourcesOutput
	p := synthetics.NewListGroupResourcesPaginator(client, input)
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

// Returns a list of all groups in the account, displaying their names, unique
// IDs, and ARNs. The groups from all Regions are returned.
func synthetics_ListGroups(cfg aws.Config, client *synthetics.Client) {
	input := &synthetics.ListGroupsInput{}

	if len(_syntheticsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _syntheticsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_syntheticsNextToken) > 0 {
		input.NextToken = aws.String(_syntheticsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*synthetics.ListGroupsOutput
	p := synthetics.NewListGroupsPaginator(client, input)
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

// Displays the tags associated with a canary or group.
func synthetics_ListTagsForResource(cfg aws.Config, client *synthetics.Client) {
	input := &synthetics.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_syntheticsResourceArn) > 0 {
		input.ResourceArn = aws.String(_syntheticsResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use this operation to run a canary that has already been created. The frequency
// of the canary runs is determined by the value of the canary's Schedule . To see
// a canary's schedule, use [GetCanary].
//
// [GetCanary]: https://docs.aws.amazon.com/AmazonSynthetics/latest/APIReference/API_GetCanary.html
func synthetics_StartCanary(cfg aws.Config, client *synthetics.Client) {
	input := &synthetics.StartCanaryInput{
		// Name: *string, // Required
	}

	if len(_syntheticsName) > 0 {
		input.Name = aws.String(_syntheticsName)
	}

	if resp, err := client.StartCanary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use this operation to start a dry run for a canary that has already been created
func synthetics_StartCanaryDryRun(cfg aws.Config, client *synthetics.Client) {
	input := &synthetics.StartCanaryDryRunInput{
		// Name: *string, // Required
	}

	if len(_syntheticsName) > 0 {
		input.Name = aws.String(_syntheticsName)
	}
	if len(_syntheticsArtifactConfig) > 0 {
		if err := assignInputField(input, "ArtifactConfig", _syntheticsArtifactConfig); err != nil {
			log.Errorf("invalid --artifact-config: %s", err.Error())
			return
		}
	}
	if len(_syntheticsArtifactS3Location) > 0 {
		input.ArtifactS3Location = aws.String(_syntheticsArtifactS3Location)
	}
	if len(_syntheticsBrowserConfigs) > 0 {
		if err := assignInputField(input, "BrowserConfigs", _syntheticsBrowserConfigs); err != nil {
			log.Errorf("invalid --browser-configs: %s", err.Error())
			return
		}
	}
	if len(_syntheticsCode) > 0 {
		if err := assignInputField(input, "Code", _syntheticsCode); err != nil {
			log.Errorf("invalid --code: %s", err.Error())
			return
		}
	}
	if len(_syntheticsExecutionRoleArn) > 0 {
		input.ExecutionRoleArn = aws.String(_syntheticsExecutionRoleArn)
	}
	if len(_syntheticsFailureRetentionPeriodInDays) > 0 {
		if err := assignInputField(input, "FailureRetentionPeriodInDays", _syntheticsFailureRetentionPeriodInDays); err != nil {
			log.Errorf("invalid --failure-retention-period-in-days: %s", err.Error())
			return
		}
	}
	if len(_syntheticsProvisionedResourceCleanup) > 0 {
		if err := assignInputField(input, "ProvisionedResourceCleanup", _syntheticsProvisionedResourceCleanup); err != nil {
			log.Errorf("invalid --provisioned-resource-cleanup: %s", err.Error())
			return
		}
	}
	if len(_syntheticsRunConfig) > 0 {
		if err := assignInputField(input, "RunConfig", _syntheticsRunConfig); err != nil {
			log.Errorf("invalid --run-config: %s", err.Error())
			return
		}
	}
	if len(_syntheticsRuntimeVersion) > 0 {
		input.RuntimeVersion = aws.String(_syntheticsRuntimeVersion)
	}
	if len(_syntheticsSuccessRetentionPeriodInDays) > 0 {
		if err := assignInputField(input, "SuccessRetentionPeriodInDays", _syntheticsSuccessRetentionPeriodInDays); err != nil {
			log.Errorf("invalid --success-retention-period-in-days: %s", err.Error())
			return
		}
	}
	if len(_syntheticsVisualReference) > 0 {
		if err := assignInputField(input, "VisualReference", _syntheticsVisualReference); err != nil {
			log.Errorf("invalid --visual-reference: %s", err.Error())
			return
		}
	}
	if len(_syntheticsVisualReferences) > 0 {
		if err := assignInputField(input, "VisualReferences", _syntheticsVisualReferences); err != nil {
			log.Errorf("invalid --visual-references: %s", err.Error())
			return
		}
	}
	if len(_syntheticsVpcConfig) > 0 {
		if err := assignInputField(input, "VpcConfig", _syntheticsVpcConfig); err != nil {
			log.Errorf("invalid --vpc-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartCanaryDryRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops the canary to prevent all future runs. If the canary is currently
// running,the run that is in progress completes on its own, publishes metrics, and
// uploads artifacts, but it is not recorded in Synthetics as a completed run.
//
// You can use StartCanary to start it running again with the canary’s current
// schedule at any point in the future.
func synthetics_StopCanary(cfg aws.Config, client *synthetics.Client) {
	input := &synthetics.StopCanaryInput{
		// Name: *string, // Required
	}

	if len(_syntheticsName) > 0 {
		input.Name = aws.String(_syntheticsName)
	}

	if resp, err := client.StopCanary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns one or more tags (key-value pairs) to the specified canary or group.
// Tags can help you organize and categorize your resources. You can also use them
// to scope user permissions, by granting a user permission to access or change
// only resources with certain tag values.
//
// Tags don't have any semantic meaning to Amazon Web Services and are interpreted
// strictly as strings of characters.
//
// You can use the TagResource action with a resource that already has tags. If
// you specify a new tag key for the resource, this tag is appended to the list of
// tags associated with the resource. If you specify a tag key that is already
// associated with the resource, the new tag value that you specify replaces the
// previous value for that tag.
//
// You can associate as many as 50 tags with a canary or group.
func synthetics_TagResource(cfg aws.Config, client *synthetics.Client) {
	input := &synthetics.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_syntheticsResourceArn) > 0 {
		input.ResourceArn = aws.String(_syntheticsResourceArn)
	}
	if len(_syntheticsTags) > 0 {
		if err := assignInputField(input, "Tags", _syntheticsTags); err != nil {
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

// Removes one or more tags from the specified resource.
func synthetics_UntagResource(cfg aws.Config, client *synthetics.Client) {
	input := &synthetics.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_syntheticsResourceArn) > 0 {
		input.ResourceArn = aws.String(_syntheticsResourceArn)
	}
	if len(_syntheticsTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _syntheticsTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration of a canary that has already been created.
// For multibrowser canaries, you can add or remove browsers by updating the
// browserConfig list in the update call. For example:
//
// - To add Firefox to a canary that currently uses Chrome, specify
// browserConfigs as [CHROME, FIREFOX]
//
// - To remove Firefox and keep only Chrome, specify browserConfigs as [CHROME]
//
// You can't use this operation to update the tags of an existing canary. To
// change the tags of an existing canary, use [TagResource].
//
// When you use the dryRunId field when updating a canary, the only other field
// you can provide is the Schedule . Adding any other field will thrown an
// exception.
//
// [TagResource]: https://docs.aws.amazon.com/AmazonSynthetics/latest/APIReference/API_TagResource.html
func synthetics_UpdateCanary(cfg aws.Config, client *synthetics.Client) {
	input := &synthetics.UpdateCanaryInput{
		// Name: *string, // Required
	}

	if len(_syntheticsName) > 0 {
		input.Name = aws.String(_syntheticsName)
	}
	if len(_syntheticsArtifactConfig) > 0 {
		if err := assignInputField(input, "ArtifactConfig", _syntheticsArtifactConfig); err != nil {
			log.Errorf("invalid --artifact-config: %s", err.Error())
			return
		}
	}
	if len(_syntheticsArtifactS3Location) > 0 {
		input.ArtifactS3Location = aws.String(_syntheticsArtifactS3Location)
	}
	if len(_syntheticsBrowserConfigs) > 0 {
		if err := assignInputField(input, "BrowserConfigs", _syntheticsBrowserConfigs); err != nil {
			log.Errorf("invalid --browser-configs: %s", err.Error())
			return
		}
	}
	if len(_syntheticsCode) > 0 {
		if err := assignInputField(input, "Code", _syntheticsCode); err != nil {
			log.Errorf("invalid --code: %s", err.Error())
			return
		}
	}
	if len(_syntheticsDryRunId) > 0 {
		input.DryRunId = aws.String(_syntheticsDryRunId)
	}
	if len(_syntheticsExecutionRoleArn) > 0 {
		input.ExecutionRoleArn = aws.String(_syntheticsExecutionRoleArn)
	}
	if len(_syntheticsFailureRetentionPeriodInDays) > 0 {
		if err := assignInputField(input, "FailureRetentionPeriodInDays", _syntheticsFailureRetentionPeriodInDays); err != nil {
			log.Errorf("invalid --failure-retention-period-in-days: %s", err.Error())
			return
		}
	}
	if len(_syntheticsProvisionedResourceCleanup) > 0 {
		if err := assignInputField(input, "ProvisionedResourceCleanup", _syntheticsProvisionedResourceCleanup); err != nil {
			log.Errorf("invalid --provisioned-resource-cleanup: %s", err.Error())
			return
		}
	}
	if len(_syntheticsRunConfig) > 0 {
		if err := assignInputField(input, "RunConfig", _syntheticsRunConfig); err != nil {
			log.Errorf("invalid --run-config: %s", err.Error())
			return
		}
	}
	if len(_syntheticsRuntimeVersion) > 0 {
		input.RuntimeVersion = aws.String(_syntheticsRuntimeVersion)
	}
	if len(_syntheticsSchedule) > 0 {
		if err := assignInputField(input, "Schedule", _syntheticsSchedule); err != nil {
			log.Errorf("invalid --schedule: %s", err.Error())
			return
		}
	}
	if len(_syntheticsSuccessRetentionPeriodInDays) > 0 {
		if err := assignInputField(input, "SuccessRetentionPeriodInDays", _syntheticsSuccessRetentionPeriodInDays); err != nil {
			log.Errorf("invalid --success-retention-period-in-days: %s", err.Error())
			return
		}
	}
	if len(_syntheticsVisualReference) > 0 {
		if err := assignInputField(input, "VisualReference", _syntheticsVisualReference); err != nil {
			log.Errorf("invalid --visual-reference: %s", err.Error())
			return
		}
	}
	if len(_syntheticsVisualReferences) > 0 {
		if err := assignInputField(input, "VisualReferences", _syntheticsVisualReferences); err != nil {
			log.Errorf("invalid --visual-references: %s", err.Error())
			return
		}
	}
	if len(_syntheticsVpcConfig) > 0 {
		if err := assignInputField(input, "VpcConfig", _syntheticsVpcConfig); err != nil {
			log.Errorf("invalid --vpc-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateCanary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_syntheticsCmd)
	_syntheticsCmd.Flags().SortFlags = false

	_syntheticsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_syntheticsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_syntheticsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_syntheticsCmd.Flags().StringVarP(&_syntheticsArtifactConfig, "artifact-config", "", "", "Artifact Config")
	_syntheticsCmd.Flags().StringVarP(&_syntheticsArtifactS3Location, "artifact-s3-location", "", "", "Artifact S3 Location")
	_syntheticsCmd.Flags().StringVarP(&_syntheticsBrowserConfigs, "browser-configs", "", "", "Browser Configs")
	_syntheticsCmd.Flags().StringVarP(&_syntheticsBrowserType, "browser-type", "", "", "Browser Type")
	_syntheticsCmd.Flags().StringVarP(&_syntheticsCode, "code", "", "", "Code")
	_syntheticsCmd.Flags().StringVarP(&_syntheticsDeleteLambda, "delete-lambda", "", "", "Delete Lambda")
	_syntheticsCmd.Flags().StringVarP(&_syntheticsDryRunId, "dry-run-id", "", "", "Dry Run ID")
	_syntheticsCmd.Flags().StringVarP(&_syntheticsExecutionRoleArn, "execution-role-arn", "", "", "Execution Role ARN")
	_syntheticsCmd.Flags().StringVarP(&_syntheticsFailureRetentionPeriodInDays, "failure-retention-period-in-days", "", "", "Failure Retention Period In Days")
	_syntheticsCmd.Flags().StringVarP(&_syntheticsGroupIdentifier, "group-identifier", "", "", "Group Identifier")
	_syntheticsCmd.Flags().StringVarP(&_syntheticsMaxResults, "max-results", "", "", "Max Results")
	_syntheticsCmd.Flags().StringVarP(&_syntheticsName, "name", "", "", "Name")
	_syntheticsCmd.Flags().StringSliceVarP(&_syntheticsNames, "names", "", nil, "Names")
	_syntheticsCmd.Flags().StringVarP(&_syntheticsNextToken, "next-token", "", "", "Next Token")
	_syntheticsCmd.Flags().StringVarP(&_syntheticsProvisionedResourceCleanup, "provisioned-resource-cleanup", "", "", "Provisioned Resource Cleanup")
	_syntheticsCmd.Flags().StringVarP(&_syntheticsResourceArn, "resource-arn", "", "", "Resource ARN")
	_syntheticsCmd.Flags().StringVarP(&_syntheticsResourcesToReplicateTags, "resources-to-replicate-tags", "", "", "Resources To Replicate Tags")
	_syntheticsCmd.Flags().StringVarP(&_syntheticsRunConfig, "run-config", "", "", "Run Config")
	_syntheticsCmd.Flags().StringVarP(&_syntheticsRunType, "run-type", "", "", "Run Type")
	_syntheticsCmd.Flags().StringVarP(&_syntheticsRuntimeVersion, "runtime-version", "", "", "Runtime Version")
	_syntheticsCmd.Flags().StringVarP(&_syntheticsSchedule, "schedule", "", "", "Schedule")
	_syntheticsCmd.Flags().StringVarP(&_syntheticsSuccessRetentionPeriodInDays, "success-retention-period-in-days", "", "", "Success Retention Period In Days")
	_syntheticsCmd.Flags().StringSliceVarP(&_syntheticsTagKeys, "tag-keys", "", nil, "Tag Keys")
	_syntheticsCmd.Flags().StringVarP(&_syntheticsTags, "tags", "", "", "Tags")
	_syntheticsCmd.Flags().StringVarP(&_syntheticsVisualReference, "visual-reference", "", "", "Visual Reference")
	_syntheticsCmd.Flags().StringVarP(&_syntheticsVisualReferences, "visual-references", "", "", "Visual References")
	_syntheticsCmd.Flags().StringVarP(&_syntheticsVpcConfig, "vpc-config", "", "", "VPC Config")

	_syntheticsCmd.Flags().BoolVarP(&_syntheticsAssociateResource, "associate-resource", "", false, "Associate Resource")
	_syntheticsCmd.Flags().BoolVarP(&_syntheticsCreateCanary, "create-canary", "", false, "Create Canary")
	_syntheticsCmd.Flags().BoolVarP(&_syntheticsCreateGroup, "create-group", "", false, "Create Group")
	_syntheticsCmd.Flags().BoolVarP(&_syntheticsDeleteCanary, "delete-canary", "", false, "Delete Canary")
	_syntheticsCmd.Flags().BoolVarP(&_syntheticsDeleteGroup, "delete-group", "", false, "Delete Group")
	_syntheticsCmd.Flags().BoolVarP(&_syntheticsDescribeCanaries, "describe-canaries", "", false, "Describe Canaries")
	_syntheticsCmd.Flags().BoolVarP(&_syntheticsDescribeCanariesLastRun, "describe-canaries-last-run", "", false, "Describe Canaries Last Run")
	_syntheticsCmd.Flags().BoolVarP(&_syntheticsDescribeRuntimeVersions, "describe-runtime-versions", "", false, "Describe Runtime Versions")
	_syntheticsCmd.Flags().BoolVarP(&_syntheticsDisassociateResource, "disassociate-resource", "", false, "Disassociate Resource")
	_syntheticsCmd.Flags().BoolVarP(&_syntheticsGetCanary, "get-canary", "", false, "Get Canary")
	_syntheticsCmd.Flags().BoolVarP(&_syntheticsGetCanaryRuns, "get-canary-runs", "", false, "Get Canary Runs")
	_syntheticsCmd.Flags().BoolVarP(&_syntheticsGetGroup, "get-group", "", false, "Get Group")
	_syntheticsCmd.Flags().BoolVarP(&_syntheticsListAssociatedGroups, "list-associated-groups", "", false, "List Associated Groups")
	_syntheticsCmd.Flags().BoolVarP(&_syntheticsListGroupResources, "list-group-resources", "", false, "List Group Resources")
	_syntheticsCmd.Flags().BoolVarP(&_syntheticsListGroups, "list-groups", "", false, "List Groups")
	_syntheticsCmd.Flags().BoolVarP(&_syntheticsListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_syntheticsCmd.Flags().BoolVarP(&_syntheticsStartCanary, "start-canary", "", false, "Start Canary")
	_syntheticsCmd.Flags().BoolVarP(&_syntheticsStartCanaryDryRun, "start-canary-dry-run", "", false, "Start Canary Dry Run")
	_syntheticsCmd.Flags().BoolVarP(&_syntheticsStopCanary, "stop-canary", "", false, "Stop Canary")
	_syntheticsCmd.Flags().BoolVarP(&_syntheticsTagResource, "tag-resource", "", false, "Tag Resource")
	_syntheticsCmd.Flags().BoolVarP(&_syntheticsUntagResource, "untag-resource", "", false, "Untag Resource")
	_syntheticsCmd.Flags().BoolVarP(&_syntheticsUpdateCanary, "update-canary", "", false, "Update Canary")

}
