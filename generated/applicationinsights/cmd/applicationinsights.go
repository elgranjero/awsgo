package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/applicationinsights"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// applicationinsightsCmd represents the applicationinsights command
var _applicationinsightsCmd = &cobra.Command{
	Use:   "applicationinsights",
	Short: "AWS applicationinsights CLI",
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
		client := applicationinsights.NewFromConfig(cfg)
		if _applicationinsightsAddWorkload {
			applicationinsights_AddWorkload(cfg, client)
			return
		}
		if _applicationinsightsCreateApplication {
			applicationinsights_CreateApplication(cfg, client)
			return
		}
		if _applicationinsightsCreateComponent {
			applicationinsights_CreateComponent(cfg, client)
			return
		}
		if _applicationinsightsCreateLogPattern {
			applicationinsights_CreateLogPattern(cfg, client)
			return
		}
		if _applicationinsightsDeleteApplication {
			applicationinsights_DeleteApplication(cfg, client)
			return
		}
		if _applicationinsightsDeleteComponent {
			applicationinsights_DeleteComponent(cfg, client)
			return
		}
		if _applicationinsightsDeleteLogPattern {
			applicationinsights_DeleteLogPattern(cfg, client)
			return
		}
		if _applicationinsightsDescribeApplication {
			applicationinsights_DescribeApplication(cfg, client)
			return
		}
		if _applicationinsightsDescribeComponent {
			applicationinsights_DescribeComponent(cfg, client)
			return
		}
		if _applicationinsightsDescribeComponentConfiguration {
			applicationinsights_DescribeComponentConfiguration(cfg, client)
			return
		}
		if _applicationinsightsDescribeComponentConfigurationRecommendation {
			applicationinsights_DescribeComponentConfigurationRecommendation(cfg, client)
			return
		}
		if _applicationinsightsDescribeLogPattern {
			applicationinsights_DescribeLogPattern(cfg, client)
			return
		}
		if _applicationinsightsDescribeObservation {
			applicationinsights_DescribeObservation(cfg, client)
			return
		}
		if _applicationinsightsDescribeProblem {
			applicationinsights_DescribeProblem(cfg, client)
			return
		}
		if _applicationinsightsDescribeProblemObservations {
			applicationinsights_DescribeProblemObservations(cfg, client)
			return
		}
		if _applicationinsightsDescribeWorkload {
			applicationinsights_DescribeWorkload(cfg, client)
			return
		}
		if _applicationinsightsListApplications {
			applicationinsights_ListApplications(cfg, client)
			return
		}
		if _applicationinsightsListComponents {
			applicationinsights_ListComponents(cfg, client)
			return
		}
		if _applicationinsightsListConfigurationHistory {
			applicationinsights_ListConfigurationHistory(cfg, client)
			return
		}
		if _applicationinsightsListLogPatternSets {
			applicationinsights_ListLogPatternSets(cfg, client)
			return
		}
		if _applicationinsightsListLogPatterns {
			applicationinsights_ListLogPatterns(cfg, client)
			return
		}
		if _applicationinsightsListProblems {
			applicationinsights_ListProblems(cfg, client)
			return
		}
		if _applicationinsightsListTagsForResource {
			applicationinsights_ListTagsForResource(cfg, client)
			return
		}
		if _applicationinsightsListWorkloads {
			applicationinsights_ListWorkloads(cfg, client)
			return
		}
		if _applicationinsightsRemoveWorkload {
			applicationinsights_RemoveWorkload(cfg, client)
			return
		}
		if _applicationinsightsTagResource {
			applicationinsights_TagResource(cfg, client)
			return
		}
		if _applicationinsightsUntagResource {
			applicationinsights_UntagResource(cfg, client)
			return
		}
		if _applicationinsightsUpdateApplication {
			applicationinsights_UpdateApplication(cfg, client)
			return
		}
		if _applicationinsightsUpdateComponent {
			applicationinsights_UpdateComponent(cfg, client)
			return
		}
		if _applicationinsightsUpdateComponentConfiguration {
			applicationinsights_UpdateComponentConfiguration(cfg, client)
			return
		}
		if _applicationinsightsUpdateLogPattern {
			applicationinsights_UpdateLogPattern(cfg, client)
			return
		}
		if _applicationinsightsUpdateProblem {
			applicationinsights_UpdateProblem(cfg, client)
			return
		}
		if _applicationinsightsUpdateWorkload {
			applicationinsights_UpdateWorkload(cfg, client)
			return
		}

	},
}

var (
	_applicationinsightsAddWorkload                                  bool
	_applicationinsightsCreateApplication                            bool
	_applicationinsightsCreateComponent                              bool
	_applicationinsightsCreateLogPattern                             bool
	_applicationinsightsDeleteApplication                            bool
	_applicationinsightsDeleteComponent                              bool
	_applicationinsightsDeleteLogPattern                             bool
	_applicationinsightsDescribeApplication                          bool
	_applicationinsightsDescribeComponent                            bool
	_applicationinsightsDescribeComponentConfiguration               bool
	_applicationinsightsDescribeComponentConfigurationRecommendation bool
	_applicationinsightsDescribeLogPattern                           bool
	_applicationinsightsDescribeObservation                          bool
	_applicationinsightsDescribeProblem                              bool
	_applicationinsightsDescribeProblemObservations                  bool
	_applicationinsightsDescribeWorkload                             bool
	_applicationinsightsListApplications                             bool
	_applicationinsightsListComponents                               bool
	_applicationinsightsListConfigurationHistory                     bool
	_applicationinsightsListLogPatternSets                           bool
	_applicationinsightsListLogPatterns                              bool
	_applicationinsightsListProblems                                 bool
	_applicationinsightsListTagsForResource                          bool
	_applicationinsightsListWorkloads                                bool
	_applicationinsightsRemoveWorkload                               bool
	_applicationinsightsTagResource                                  bool
	_applicationinsightsUntagResource                                bool
	_applicationinsightsUpdateApplication                            bool
	_applicationinsightsUpdateComponent                              bool
	_applicationinsightsUpdateComponentConfiguration                 bool
	_applicationinsightsUpdateLogPattern                             bool
	_applicationinsightsUpdateProblem                                bool
	_applicationinsightsUpdateWorkload                               bool

	_applicationinsightsAccountId               string
	_applicationinsightsAttachMissingPermission string
	_applicationinsightsAutoConfigEnabled       string
	_applicationinsightsAutoCreate              string
	_applicationinsightsComponentConfiguration  string
	_applicationinsightsComponentName           string
	_applicationinsightsCWEMonitorEnabled       string
	_applicationinsightsEndTime                 string
	_applicationinsightsEventStatus             string
	_applicationinsightsGroupingType            string
	_applicationinsightsMaxResults              string
	_applicationinsightsMonitor                 string
	_applicationinsightsNewComponentName        string
	_applicationinsightsNextToken               string
	_applicationinsightsObservationId           string
	_applicationinsightsOpsCenterEnabled        string
	_applicationinsightsOpsItemSNSTopicArn      string
	_applicationinsightsPattern                 string
	_applicationinsightsPatternName             string
	_applicationinsightsPatternSetName          string
	_applicationinsightsProblemId               string
	_applicationinsightsRank                    string
	_applicationinsightsRecommendationType      string
	_applicationinsightsRemoveSNSTopic          string
	_applicationinsightsResourceARN             string
	_applicationinsightsResourceGroupName       string
	_applicationinsightsResourceList            []string
	_applicationinsightsSNSNotificationArn      string
	_applicationinsightsStartTime               string
	_applicationinsightsTagKeys                 []string
	_applicationinsightsTags                    string
	_applicationinsightsTier                    string
	_applicationinsightsUpdateStatus            string
	_applicationinsightsVisibility              string
	_applicationinsightsWorkloadConfiguration   string
	_applicationinsightsWorkloadId              string
	_applicationinsightsWorkloadName            string
)

// Adds a workload to a component. Each component can have at most five workloads.
func applicationinsights_AddWorkload(cfg aws.Config, client *applicationinsights.Client) {
	input := &applicationinsights.AddWorkloadInput{
		// ComponentName: *string, // Required
		// ResourceGroupName: *string, // Required
		// WorkloadConfiguration: *types.WorkloadConfiguration, // Required
	}

	if len(_applicationinsightsComponentName) > 0 {
		input.ComponentName = aws.String(_applicationinsightsComponentName)
	}
	if len(_applicationinsightsResourceGroupName) > 0 {
		input.ResourceGroupName = aws.String(_applicationinsightsResourceGroupName)
	}
	if len(_applicationinsightsWorkloadConfiguration) > 0 {
		if err := assignInputField(input, "WorkloadConfiguration", _applicationinsightsWorkloadConfiguration); err != nil {
			log.Errorf("invalid --workload-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.AddWorkload(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds an application that is created from a resource group.
func applicationinsights_CreateApplication(cfg aws.Config, client *applicationinsights.Client) {
	input := &applicationinsights.CreateApplicationInput{}

	if len(_applicationinsightsAttachMissingPermission) > 0 {
		if err := assignInputField(input, "AttachMissingPermission", _applicationinsightsAttachMissingPermission); err != nil {
			log.Errorf("invalid --attach-missing-permission: %s", err.Error())
			return
		}
	}
	if len(_applicationinsightsAutoConfigEnabled) > 0 {
		if err := assignInputField(input, "AutoConfigEnabled", _applicationinsightsAutoConfigEnabled); err != nil {
			log.Errorf("invalid --auto-config-enabled: %s", err.Error())
			return
		}
	}
	if len(_applicationinsightsAutoCreate) > 0 {
		if err := assignInputField(input, "AutoCreate", _applicationinsightsAutoCreate); err != nil {
			log.Errorf("invalid --auto-create: %s", err.Error())
			return
		}
	}
	if len(_applicationinsightsCWEMonitorEnabled) > 0 {
		if err := assignInputField(input, "CWEMonitorEnabled", _applicationinsightsCWEMonitorEnabled); err != nil {
			log.Errorf("invalid --cwe-monitor-enabled: %s", err.Error())
			return
		}
	}
	if len(_applicationinsightsGroupingType) > 0 {
		if err := assignInputField(input, "GroupingType", _applicationinsightsGroupingType); err != nil {
			log.Errorf("invalid --grouping-type: %s", err.Error())
			return
		}
	}
	if len(_applicationinsightsOpsCenterEnabled) > 0 {
		if err := assignInputField(input, "OpsCenterEnabled", _applicationinsightsOpsCenterEnabled); err != nil {
			log.Errorf("invalid --ops-center-enabled: %s", err.Error())
			return
		}
	}
	if len(_applicationinsightsOpsItemSNSTopicArn) > 0 {
		input.OpsItemSNSTopicArn = aws.String(_applicationinsightsOpsItemSNSTopicArn)
	}
	if len(_applicationinsightsResourceGroupName) > 0 {
		input.ResourceGroupName = aws.String(_applicationinsightsResourceGroupName)
	}
	if len(_applicationinsightsSNSNotificationArn) > 0 {
		input.SNSNotificationArn = aws.String(_applicationinsightsSNSNotificationArn)
	}
	if len(_applicationinsightsTags) > 0 {
		if err := assignInputField(input, "Tags", _applicationinsightsTags); err != nil {
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

// Creates a custom component by grouping similar standalone instances to monitor.
func applicationinsights_CreateComponent(cfg aws.Config, client *applicationinsights.Client) {
	input := &applicationinsights.CreateComponentInput{
		// ComponentName: *string, // Required
		// ResourceGroupName: *string, // Required
		// ResourceList: []string, // Required
	}

	if len(_applicationinsightsComponentName) > 0 {
		input.ComponentName = aws.String(_applicationinsightsComponentName)
	}
	if len(_applicationinsightsResourceGroupName) > 0 {
		input.ResourceGroupName = aws.String(_applicationinsightsResourceGroupName)
	}
	if len(_applicationinsightsResourceList) > 0 {
		input.ResourceList = append([]string(nil), _applicationinsightsResourceList...)
	}

	if resp, err := client.CreateComponent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds an log pattern to a LogPatternSet .
func applicationinsights_CreateLogPattern(cfg aws.Config, client *applicationinsights.Client) {
	input := &applicationinsights.CreateLogPatternInput{
		// Pattern: *string, // Required
		// PatternName: *string, // Required
		// PatternSetName: *string, // Required
		// Rank: int32, // Required
		// ResourceGroupName: *string, // Required
	}

	if len(_applicationinsightsPattern) > 0 {
		input.Pattern = aws.String(_applicationinsightsPattern)
	}
	if len(_applicationinsightsPatternName) > 0 {
		input.PatternName = aws.String(_applicationinsightsPatternName)
	}
	if len(_applicationinsightsPatternSetName) > 0 {
		input.PatternSetName = aws.String(_applicationinsightsPatternSetName)
	}
	if len(_applicationinsightsRank) > 0 {
		if err := assignInputField(input, "Rank", _applicationinsightsRank); err != nil {
			log.Errorf("invalid --rank: %s", err.Error())
			return
		}
	}
	if len(_applicationinsightsResourceGroupName) > 0 {
		input.ResourceGroupName = aws.String(_applicationinsightsResourceGroupName)
	}

	if resp, err := client.CreateLogPattern(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified application from monitoring. Does not delete the
// application.
func applicationinsights_DeleteApplication(cfg aws.Config, client *applicationinsights.Client) {
	input := &applicationinsights.DeleteApplicationInput{
		// ResourceGroupName: *string, // Required
	}

	if len(_applicationinsightsResourceGroupName) > 0 {
		input.ResourceGroupName = aws.String(_applicationinsightsResourceGroupName)
	}

	if resp, err := client.DeleteApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Ungroups a custom component. When you ungroup custom components, all applicable
// monitors that are set up for the component are removed and the instances revert
// to their standalone status.
func applicationinsights_DeleteComponent(cfg aws.Config, client *applicationinsights.Client) {
	input := &applicationinsights.DeleteComponentInput{
		// ComponentName: *string, // Required
		// ResourceGroupName: *string, // Required
	}

	if len(_applicationinsightsComponentName) > 0 {
		input.ComponentName = aws.String(_applicationinsightsComponentName)
	}
	if len(_applicationinsightsResourceGroupName) > 0 {
		input.ResourceGroupName = aws.String(_applicationinsightsResourceGroupName)
	}

	if resp, err := client.DeleteComponent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified log pattern from a LogPatternSet .
func applicationinsights_DeleteLogPattern(cfg aws.Config, client *applicationinsights.Client) {
	input := &applicationinsights.DeleteLogPatternInput{
		// PatternName: *string, // Required
		// PatternSetName: *string, // Required
		// ResourceGroupName: *string, // Required
	}

	if len(_applicationinsightsPatternName) > 0 {
		input.PatternName = aws.String(_applicationinsightsPatternName)
	}
	if len(_applicationinsightsPatternSetName) > 0 {
		input.PatternSetName = aws.String(_applicationinsightsPatternSetName)
	}
	if len(_applicationinsightsResourceGroupName) > 0 {
		input.ResourceGroupName = aws.String(_applicationinsightsResourceGroupName)
	}

	if resp, err := client.DeleteLogPattern(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the application.
func applicationinsights_DescribeApplication(cfg aws.Config, client *applicationinsights.Client) {
	input := &applicationinsights.DescribeApplicationInput{
		// ResourceGroupName: *string, // Required
	}

	if len(_applicationinsightsResourceGroupName) > 0 {
		input.ResourceGroupName = aws.String(_applicationinsightsResourceGroupName)
	}
	if len(_applicationinsightsAccountId) > 0 {
		input.AccountId = aws.String(_applicationinsightsAccountId)
	}

	if resp, err := client.DescribeApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a component and lists the resources that are grouped together in a
// component.
func applicationinsights_DescribeComponent(cfg aws.Config, client *applicationinsights.Client) {
	input := &applicationinsights.DescribeComponentInput{
		// ComponentName: *string, // Required
		// ResourceGroupName: *string, // Required
	}

	if len(_applicationinsightsComponentName) > 0 {
		input.ComponentName = aws.String(_applicationinsightsComponentName)
	}
	if len(_applicationinsightsResourceGroupName) > 0 {
		input.ResourceGroupName = aws.String(_applicationinsightsResourceGroupName)
	}
	if len(_applicationinsightsAccountId) > 0 {
		input.AccountId = aws.String(_applicationinsightsAccountId)
	}

	if resp, err := client.DescribeComponent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the monitoring configuration of the component.
func applicationinsights_DescribeComponentConfiguration(cfg aws.Config, client *applicationinsights.Client) {
	input := &applicationinsights.DescribeComponentConfigurationInput{
		// ComponentName: *string, // Required
		// ResourceGroupName: *string, // Required
	}

	if len(_applicationinsightsComponentName) > 0 {
		input.ComponentName = aws.String(_applicationinsightsComponentName)
	}
	if len(_applicationinsightsResourceGroupName) > 0 {
		input.ResourceGroupName = aws.String(_applicationinsightsResourceGroupName)
	}
	if len(_applicationinsightsAccountId) > 0 {
		input.AccountId = aws.String(_applicationinsightsAccountId)
	}

	if resp, err := client.DescribeComponentConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the recommended monitoring configuration of the component.
func applicationinsights_DescribeComponentConfigurationRecommendation(cfg aws.Config, client *applicationinsights.Client) {
	input := &applicationinsights.DescribeComponentConfigurationRecommendationInput{
		// ComponentName: *string, // Required
		// ResourceGroupName: *string, // Required
		// Tier: types.Tier, // Required
	}

	if len(_applicationinsightsComponentName) > 0 {
		input.ComponentName = aws.String(_applicationinsightsComponentName)
	}
	if len(_applicationinsightsResourceGroupName) > 0 {
		input.ResourceGroupName = aws.String(_applicationinsightsResourceGroupName)
	}
	if len(_applicationinsightsTier) > 0 {
		if err := assignInputField(input, "Tier", _applicationinsightsTier); err != nil {
			log.Errorf("invalid --tier: %s", err.Error())
			return
		}
	}
	if len(_applicationinsightsRecommendationType) > 0 {
		if err := assignInputField(input, "RecommendationType", _applicationinsightsRecommendationType); err != nil {
			log.Errorf("invalid --recommendation-type: %s", err.Error())
			return
		}
	}
	if len(_applicationinsightsWorkloadName) > 0 {
		input.WorkloadName = aws.String(_applicationinsightsWorkloadName)
	}

	if resp, err := client.DescribeComponentConfigurationRecommendation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describe a specific log pattern from a LogPatternSet .
func applicationinsights_DescribeLogPattern(cfg aws.Config, client *applicationinsights.Client) {
	input := &applicationinsights.DescribeLogPatternInput{
		// PatternName: *string, // Required
		// PatternSetName: *string, // Required
		// ResourceGroupName: *string, // Required
	}

	if len(_applicationinsightsPatternName) > 0 {
		input.PatternName = aws.String(_applicationinsightsPatternName)
	}
	if len(_applicationinsightsPatternSetName) > 0 {
		input.PatternSetName = aws.String(_applicationinsightsPatternSetName)
	}
	if len(_applicationinsightsResourceGroupName) > 0 {
		input.ResourceGroupName = aws.String(_applicationinsightsResourceGroupName)
	}
	if len(_applicationinsightsAccountId) > 0 {
		input.AccountId = aws.String(_applicationinsightsAccountId)
	}

	if resp, err := client.DescribeLogPattern(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes an anomaly or error with the application.
func applicationinsights_DescribeObservation(cfg aws.Config, client *applicationinsights.Client) {
	input := &applicationinsights.DescribeObservationInput{
		// ObservationId: *string, // Required
	}

	if len(_applicationinsightsObservationId) > 0 {
		input.ObservationId = aws.String(_applicationinsightsObservationId)
	}
	if len(_applicationinsightsAccountId) > 0 {
		input.AccountId = aws.String(_applicationinsightsAccountId)
	}

	if resp, err := client.DescribeObservation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes an application problem.
func applicationinsights_DescribeProblem(cfg aws.Config, client *applicationinsights.Client) {
	input := &applicationinsights.DescribeProblemInput{
		// ProblemId: *string, // Required
	}

	if len(_applicationinsightsProblemId) > 0 {
		input.ProblemId = aws.String(_applicationinsightsProblemId)
	}
	if len(_applicationinsightsAccountId) > 0 {
		input.AccountId = aws.String(_applicationinsightsAccountId)
	}

	if resp, err := client.DescribeProblem(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the anomalies or errors associated with the problem.
func applicationinsights_DescribeProblemObservations(cfg aws.Config, client *applicationinsights.Client) {
	input := &applicationinsights.DescribeProblemObservationsInput{
		// ProblemId: *string, // Required
	}

	if len(_applicationinsightsProblemId) > 0 {
		input.ProblemId = aws.String(_applicationinsightsProblemId)
	}
	if len(_applicationinsightsAccountId) > 0 {
		input.AccountId = aws.String(_applicationinsightsAccountId)
	}

	if resp, err := client.DescribeProblemObservations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a workload and its configuration.
func applicationinsights_DescribeWorkload(cfg aws.Config, client *applicationinsights.Client) {
	input := &applicationinsights.DescribeWorkloadInput{
		// ComponentName: *string, // Required
		// ResourceGroupName: *string, // Required
		// WorkloadId: *string, // Required
	}

	if len(_applicationinsightsComponentName) > 0 {
		input.ComponentName = aws.String(_applicationinsightsComponentName)
	}
	if len(_applicationinsightsResourceGroupName) > 0 {
		input.ResourceGroupName = aws.String(_applicationinsightsResourceGroupName)
	}
	if len(_applicationinsightsWorkloadId) > 0 {
		input.WorkloadId = aws.String(_applicationinsightsWorkloadId)
	}
	if len(_applicationinsightsAccountId) > 0 {
		input.AccountId = aws.String(_applicationinsightsAccountId)
	}

	if resp, err := client.DescribeWorkload(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the IDs of the applications that you are monitoring.
func applicationinsights_ListApplications(cfg aws.Config, client *applicationinsights.Client) {
	input := &applicationinsights.ListApplicationsInput{}

	if len(_applicationinsightsAccountId) > 0 {
		input.AccountId = aws.String(_applicationinsightsAccountId)
	}
	if len(_applicationinsightsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _applicationinsightsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_applicationinsightsNextToken) > 0 {
		input.NextToken = aws.String(_applicationinsightsNextToken)
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

	var results []*applicationinsights.ListApplicationsOutput
	p := applicationinsights.NewListApplicationsPaginator(client, input)
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

// Lists the auto-grouped, standalone, and custom components of the application.
func applicationinsights_ListComponents(cfg aws.Config, client *applicationinsights.Client) {
	input := &applicationinsights.ListComponentsInput{
		// ResourceGroupName: *string, // Required
	}

	if len(_applicationinsightsResourceGroupName) > 0 {
		input.ResourceGroupName = aws.String(_applicationinsightsResourceGroupName)
	}
	if len(_applicationinsightsAccountId) > 0 {
		input.AccountId = aws.String(_applicationinsightsAccountId)
	}
	if len(_applicationinsightsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _applicationinsightsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_applicationinsightsNextToken) > 0 {
		input.NextToken = aws.String(_applicationinsightsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListComponents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*applicationinsights.ListComponentsOutput
	p := applicationinsights.NewListComponentsPaginator(client, input)
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

// Lists the INFO, WARN, and ERROR events for periodic configuration updates
// performed by Application Insights. Examples of events represented are:
//
// - INFO: creating a new alarm or updating an alarm threshold.
//
// - WARN: alarm not created due to insufficient data points used to predict
// thresholds.
//
// - ERROR: alarm not created due to permission errors or exceeding quotas.
func applicationinsights_ListConfigurationHistory(cfg aws.Config, client *applicationinsights.Client) {
	input := &applicationinsights.ListConfigurationHistoryInput{}

	if len(_applicationinsightsAccountId) > 0 {
		input.AccountId = aws.String(_applicationinsightsAccountId)
	}
	if len(_applicationinsightsEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _applicationinsightsEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_applicationinsightsEventStatus) > 0 {
		if err := assignInputField(input, "EventStatus", _applicationinsightsEventStatus); err != nil {
			log.Errorf("invalid --event-status: %s", err.Error())
			return
		}
	}
	if len(_applicationinsightsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _applicationinsightsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_applicationinsightsNextToken) > 0 {
		input.NextToken = aws.String(_applicationinsightsNextToken)
	}
	if len(_applicationinsightsResourceGroupName) > 0 {
		input.ResourceGroupName = aws.String(_applicationinsightsResourceGroupName)
	}
	if len(_applicationinsightsStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _applicationinsightsStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListConfigurationHistory(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*applicationinsights.ListConfigurationHistoryOutput
	p := applicationinsights.NewListConfigurationHistoryPaginator(client, input)
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

// Lists the log pattern sets in the specific application.
func applicationinsights_ListLogPatternSets(cfg aws.Config, client *applicationinsights.Client) {
	input := &applicationinsights.ListLogPatternSetsInput{
		// ResourceGroupName: *string, // Required
	}

	if len(_applicationinsightsResourceGroupName) > 0 {
		input.ResourceGroupName = aws.String(_applicationinsightsResourceGroupName)
	}
	if len(_applicationinsightsAccountId) > 0 {
		input.AccountId = aws.String(_applicationinsightsAccountId)
	}
	if len(_applicationinsightsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _applicationinsightsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_applicationinsightsNextToken) > 0 {
		input.NextToken = aws.String(_applicationinsightsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListLogPatternSets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*applicationinsights.ListLogPatternSetsOutput
	p := applicationinsights.NewListLogPatternSetsPaginator(client, input)
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

// Lists the log patterns in the specific log LogPatternSet .
func applicationinsights_ListLogPatterns(cfg aws.Config, client *applicationinsights.Client) {
	input := &applicationinsights.ListLogPatternsInput{
		// ResourceGroupName: *string, // Required
	}

	if len(_applicationinsightsResourceGroupName) > 0 {
		input.ResourceGroupName = aws.String(_applicationinsightsResourceGroupName)
	}
	if len(_applicationinsightsAccountId) > 0 {
		input.AccountId = aws.String(_applicationinsightsAccountId)
	}
	if len(_applicationinsightsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _applicationinsightsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_applicationinsightsNextToken) > 0 {
		input.NextToken = aws.String(_applicationinsightsNextToken)
	}
	if len(_applicationinsightsPatternSetName) > 0 {
		input.PatternSetName = aws.String(_applicationinsightsPatternSetName)
	}

	if disablePaginator() {
		if resp, err := client.ListLogPatterns(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*applicationinsights.ListLogPatternsOutput
	p := applicationinsights.NewListLogPatternsPaginator(client, input)
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

// Lists the problems with your application.
func applicationinsights_ListProblems(cfg aws.Config, client *applicationinsights.Client) {
	input := &applicationinsights.ListProblemsInput{}

	if len(_applicationinsightsAccountId) > 0 {
		input.AccountId = aws.String(_applicationinsightsAccountId)
	}
	if len(_applicationinsightsComponentName) > 0 {
		input.ComponentName = aws.String(_applicationinsightsComponentName)
	}
	if len(_applicationinsightsEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _applicationinsightsEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_applicationinsightsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _applicationinsightsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_applicationinsightsNextToken) > 0 {
		input.NextToken = aws.String(_applicationinsightsNextToken)
	}
	if len(_applicationinsightsResourceGroupName) > 0 {
		input.ResourceGroupName = aws.String(_applicationinsightsResourceGroupName)
	}
	if len(_applicationinsightsStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _applicationinsightsStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_applicationinsightsVisibility) > 0 {
		if err := assignInputField(input, "Visibility", _applicationinsightsVisibility); err != nil {
			log.Errorf("invalid --visibility: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListProblems(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*applicationinsights.ListProblemsOutput
	p := applicationinsights.NewListProblemsPaginator(client, input)
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

// Retrieve a list of the tags (keys and values) that are associated with a
// specified application. A tag is a label that you optionally define and associate
// with an application. Each tag consists of a required tag key and an optional
// associated tag value. A tag key is a general label that acts as a category for
// more specific tag values. A tag value acts as a descriptor within a tag key.
func applicationinsights_ListTagsForResource(cfg aws.Config, client *applicationinsights.Client) {
	input := &applicationinsights.ListTagsForResourceInput{
		// ResourceARN: *string, // Required
	}

	if len(_applicationinsightsResourceARN) > 0 {
		input.ResourceARN = aws.String(_applicationinsightsResourceARN)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the workloads that are configured on a given component.
func applicationinsights_ListWorkloads(cfg aws.Config, client *applicationinsights.Client) {
	input := &applicationinsights.ListWorkloadsInput{
		// ComponentName: *string, // Required
		// ResourceGroupName: *string, // Required
	}

	if len(_applicationinsightsComponentName) > 0 {
		input.ComponentName = aws.String(_applicationinsightsComponentName)
	}
	if len(_applicationinsightsResourceGroupName) > 0 {
		input.ResourceGroupName = aws.String(_applicationinsightsResourceGroupName)
	}
	if len(_applicationinsightsAccountId) > 0 {
		input.AccountId = aws.String(_applicationinsightsAccountId)
	}
	if len(_applicationinsightsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _applicationinsightsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_applicationinsightsNextToken) > 0 {
		input.NextToken = aws.String(_applicationinsightsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListWorkloads(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*applicationinsights.ListWorkloadsOutput
	p := applicationinsights.NewListWorkloadsPaginator(client, input)
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

// Remove workload from a component.
func applicationinsights_RemoveWorkload(cfg aws.Config, client *applicationinsights.Client) {
	input := &applicationinsights.RemoveWorkloadInput{
		// ComponentName: *string, // Required
		// ResourceGroupName: *string, // Required
		// WorkloadId: *string, // Required
	}

	if len(_applicationinsightsComponentName) > 0 {
		input.ComponentName = aws.String(_applicationinsightsComponentName)
	}
	if len(_applicationinsightsResourceGroupName) > 0 {
		input.ResourceGroupName = aws.String(_applicationinsightsResourceGroupName)
	}
	if len(_applicationinsightsWorkloadId) > 0 {
		input.WorkloadId = aws.String(_applicationinsightsWorkloadId)
	}

	if resp, err := client.RemoveWorkload(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Add one or more tags (keys and values) to a specified application. A tag is a
// label that you optionally define and associate with an application. Tags can
// help you categorize and manage application in different ways, such as by
// purpose, owner, environment, or other criteria.
//
// Each tag consists of a required tag key and an associated tag value, both of
// which you define. A tag key is a general label that acts as a category for more
// specific tag values. A tag value acts as a descriptor within a tag key.
func applicationinsights_TagResource(cfg aws.Config, client *applicationinsights.Client) {
	input := &applicationinsights.TagResourceInput{
		// ResourceARN: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_applicationinsightsResourceARN) > 0 {
		input.ResourceARN = aws.String(_applicationinsightsResourceARN)
	}
	if len(_applicationinsightsTags) > 0 {
		if err := assignInputField(input, "Tags", _applicationinsightsTags); err != nil {
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

// Remove one or more tags (keys and values) from a specified application.
func applicationinsights_UntagResource(cfg aws.Config, client *applicationinsights.Client) {
	input := &applicationinsights.UntagResourceInput{
		// ResourceARN: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_applicationinsightsResourceARN) > 0 {
		input.ResourceARN = aws.String(_applicationinsightsResourceARN)
	}
	if len(_applicationinsightsTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _applicationinsightsTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the application.
func applicationinsights_UpdateApplication(cfg aws.Config, client *applicationinsights.Client) {
	input := &applicationinsights.UpdateApplicationInput{
		// ResourceGroupName: *string, // Required
	}

	if len(_applicationinsightsResourceGroupName) > 0 {
		input.ResourceGroupName = aws.String(_applicationinsightsResourceGroupName)
	}
	if len(_applicationinsightsAttachMissingPermission) > 0 {
		if err := assignInputField(input, "AttachMissingPermission", _applicationinsightsAttachMissingPermission); err != nil {
			log.Errorf("invalid --attach-missing-permission: %s", err.Error())
			return
		}
	}
	if len(_applicationinsightsAutoConfigEnabled) > 0 {
		if err := assignInputField(input, "AutoConfigEnabled", _applicationinsightsAutoConfigEnabled); err != nil {
			log.Errorf("invalid --auto-config-enabled: %s", err.Error())
			return
		}
	}
	if len(_applicationinsightsCWEMonitorEnabled) > 0 {
		if err := assignInputField(input, "CWEMonitorEnabled", _applicationinsightsCWEMonitorEnabled); err != nil {
			log.Errorf("invalid --cwe-monitor-enabled: %s", err.Error())
			return
		}
	}
	if len(_applicationinsightsOpsCenterEnabled) > 0 {
		if err := assignInputField(input, "OpsCenterEnabled", _applicationinsightsOpsCenterEnabled); err != nil {
			log.Errorf("invalid --ops-center-enabled: %s", err.Error())
			return
		}
	}
	if len(_applicationinsightsOpsItemSNSTopicArn) > 0 {
		input.OpsItemSNSTopicArn = aws.String(_applicationinsightsOpsItemSNSTopicArn)
	}
	if len(_applicationinsightsRemoveSNSTopic) > 0 {
		if err := assignInputField(input, "RemoveSNSTopic", _applicationinsightsRemoveSNSTopic); err != nil {
			log.Errorf("invalid --remove-sns-topic: %s", err.Error())
			return
		}
	}
	if len(_applicationinsightsSNSNotificationArn) > 0 {
		input.SNSNotificationArn = aws.String(_applicationinsightsSNSNotificationArn)
	}

	if resp, err := client.UpdateApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the custom component name and/or the list of resources that make up the
// component.
func applicationinsights_UpdateComponent(cfg aws.Config, client *applicationinsights.Client) {
	input := &applicationinsights.UpdateComponentInput{
		// ComponentName: *string, // Required
		// ResourceGroupName: *string, // Required
	}

	if len(_applicationinsightsComponentName) > 0 {
		input.ComponentName = aws.String(_applicationinsightsComponentName)
	}
	if len(_applicationinsightsResourceGroupName) > 0 {
		input.ResourceGroupName = aws.String(_applicationinsightsResourceGroupName)
	}
	if len(_applicationinsightsNewComponentName) > 0 {
		input.NewComponentName = aws.String(_applicationinsightsNewComponentName)
	}
	if len(_applicationinsightsResourceList) > 0 {
		input.ResourceList = append([]string(nil), _applicationinsightsResourceList...)
	}

	if resp, err := client.UpdateComponent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the monitoring configurations for the component. The configuration
// input parameter is an escaped JSON of the configuration and should match the
// schema of what is returned by DescribeComponentConfigurationRecommendation .
func applicationinsights_UpdateComponentConfiguration(cfg aws.Config, client *applicationinsights.Client) {
	input := &applicationinsights.UpdateComponentConfigurationInput{
		// ComponentName: *string, // Required
		// ResourceGroupName: *string, // Required
	}

	if len(_applicationinsightsComponentName) > 0 {
		input.ComponentName = aws.String(_applicationinsightsComponentName)
	}
	if len(_applicationinsightsResourceGroupName) > 0 {
		input.ResourceGroupName = aws.String(_applicationinsightsResourceGroupName)
	}
	if len(_applicationinsightsAutoConfigEnabled) > 0 {
		if err := assignInputField(input, "AutoConfigEnabled", _applicationinsightsAutoConfigEnabled); err != nil {
			log.Errorf("invalid --auto-config-enabled: %s", err.Error())
			return
		}
	}
	if len(_applicationinsightsComponentConfiguration) > 0 {
		input.ComponentConfiguration = aws.String(_applicationinsightsComponentConfiguration)
	}
	if len(_applicationinsightsMonitor) > 0 {
		if err := assignInputField(input, "Monitor", _applicationinsightsMonitor); err != nil {
			log.Errorf("invalid --monitor: %s", err.Error())
			return
		}
	}
	if len(_applicationinsightsTier) > 0 {
		if err := assignInputField(input, "Tier", _applicationinsightsTier); err != nil {
			log.Errorf("invalid --tier: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateComponentConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a log pattern to a LogPatternSet .
func applicationinsights_UpdateLogPattern(cfg aws.Config, client *applicationinsights.Client) {
	input := &applicationinsights.UpdateLogPatternInput{
		// PatternName: *string, // Required
		// PatternSetName: *string, // Required
		// ResourceGroupName: *string, // Required
	}

	if len(_applicationinsightsPatternName) > 0 {
		input.PatternName = aws.String(_applicationinsightsPatternName)
	}
	if len(_applicationinsightsPatternSetName) > 0 {
		input.PatternSetName = aws.String(_applicationinsightsPatternSetName)
	}
	if len(_applicationinsightsResourceGroupName) > 0 {
		input.ResourceGroupName = aws.String(_applicationinsightsResourceGroupName)
	}
	if len(_applicationinsightsPattern) > 0 {
		input.Pattern = aws.String(_applicationinsightsPattern)
	}
	if len(_applicationinsightsRank) > 0 {
		if err := assignInputField(input, "Rank", _applicationinsightsRank); err != nil {
			log.Errorf("invalid --rank: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateLogPattern(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the visibility of the problem or specifies the problem as RESOLVED .
func applicationinsights_UpdateProblem(cfg aws.Config, client *applicationinsights.Client) {
	input := &applicationinsights.UpdateProblemInput{
		// ProblemId: *string, // Required
	}

	if len(_applicationinsightsProblemId) > 0 {
		input.ProblemId = aws.String(_applicationinsightsProblemId)
	}
	if len(_applicationinsightsUpdateStatus) > 0 {
		if err := assignInputField(input, "UpdateStatus", _applicationinsightsUpdateStatus); err != nil {
			log.Errorf("invalid --update-status: %s", err.Error())
			return
		}
	}
	if len(_applicationinsightsVisibility) > 0 {
		if err := assignInputField(input, "Visibility", _applicationinsightsVisibility); err != nil {
			log.Errorf("invalid --visibility: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateProblem(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a workload to a component. Each component can have at most five workloads.
func applicationinsights_UpdateWorkload(cfg aws.Config, client *applicationinsights.Client) {
	input := &applicationinsights.UpdateWorkloadInput{
		// ComponentName: *string, // Required
		// ResourceGroupName: *string, // Required
		// WorkloadConfiguration: *types.WorkloadConfiguration, // Required
	}

	if len(_applicationinsightsComponentName) > 0 {
		input.ComponentName = aws.String(_applicationinsightsComponentName)
	}
	if len(_applicationinsightsResourceGroupName) > 0 {
		input.ResourceGroupName = aws.String(_applicationinsightsResourceGroupName)
	}
	if len(_applicationinsightsWorkloadConfiguration) > 0 {
		if err := assignInputField(input, "WorkloadConfiguration", _applicationinsightsWorkloadConfiguration); err != nil {
			log.Errorf("invalid --workload-configuration: %s", err.Error())
			return
		}
	}
	if len(_applicationinsightsWorkloadId) > 0 {
		input.WorkloadId = aws.String(_applicationinsightsWorkloadId)
	}

	if resp, err := client.UpdateWorkload(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_applicationinsightsCmd)
	_applicationinsightsCmd.Flags().SortFlags = false

	_applicationinsightsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_applicationinsightsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_applicationinsightsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_applicationinsightsCmd.Flags().StringVarP(&_applicationinsightsAccountId, "account-id", "", "", "Account ID")
	_applicationinsightsCmd.Flags().StringVarP(&_applicationinsightsAttachMissingPermission, "attach-missing-permission", "", "", "Attach Missing Permission")
	_applicationinsightsCmd.Flags().StringVarP(&_applicationinsightsAutoConfigEnabled, "auto-config-enabled", "", "", "Auto Config Enabled")
	_applicationinsightsCmd.Flags().StringVarP(&_applicationinsightsAutoCreate, "auto-create", "", "", "Auto Create")
	_applicationinsightsCmd.Flags().StringVarP(&_applicationinsightsComponentConfiguration, "component-configuration", "", "", "Component Configuration")
	_applicationinsightsCmd.Flags().StringVarP(&_applicationinsightsComponentName, "component-name", "", "", "Component Name")
	_applicationinsightsCmd.Flags().StringVarP(&_applicationinsightsCWEMonitorEnabled, "cwe-monitor-enabled", "", "", "Cwe Monitor Enabled")
	_applicationinsightsCmd.Flags().StringVarP(&_applicationinsightsEndTime, "end-time", "", "", "End Time")
	_applicationinsightsCmd.Flags().StringVarP(&_applicationinsightsEventStatus, "event-status", "", "", "Event Status")
	_applicationinsightsCmd.Flags().StringVarP(&_applicationinsightsGroupingType, "grouping-type", "", "", "Grouping Type")
	_applicationinsightsCmd.Flags().StringVarP(&_applicationinsightsMaxResults, "max-results", "", "", "Max Results")
	_applicationinsightsCmd.Flags().StringVarP(&_applicationinsightsMonitor, "monitor", "", "", "Monitor")
	_applicationinsightsCmd.Flags().StringVarP(&_applicationinsightsNewComponentName, "new-component-name", "", "", "New Component Name")
	_applicationinsightsCmd.Flags().StringVarP(&_applicationinsightsNextToken, "next-token", "", "", "Next Token")
	_applicationinsightsCmd.Flags().StringVarP(&_applicationinsightsObservationId, "observation-id", "", "", "Observation ID")
	_applicationinsightsCmd.Flags().StringVarP(&_applicationinsightsOpsCenterEnabled, "ops-center-enabled", "", "", "Ops Center Enabled")
	_applicationinsightsCmd.Flags().StringVarP(&_applicationinsightsOpsItemSNSTopicArn, "ops-item-sns-topic-arn", "", "", "Ops Item SNS Topic ARN")
	_applicationinsightsCmd.Flags().StringVarP(&_applicationinsightsPattern, "pattern", "", "", "Pattern")
	_applicationinsightsCmd.Flags().StringVarP(&_applicationinsightsPatternName, "pattern-name", "", "", "Pattern Name")
	_applicationinsightsCmd.Flags().StringVarP(&_applicationinsightsPatternSetName, "pattern-set-name", "", "", "Pattern Set Name")
	_applicationinsightsCmd.Flags().StringVarP(&_applicationinsightsProblemId, "problem-id", "", "", "Problem ID")
	_applicationinsightsCmd.Flags().StringVarP(&_applicationinsightsRank, "rank", "", "", "Rank")
	_applicationinsightsCmd.Flags().StringVarP(&_applicationinsightsRecommendationType, "recommendation-type", "", "", "Recommendation Type")
	_applicationinsightsCmd.Flags().StringVarP(&_applicationinsightsRemoveSNSTopic, "remove-sns-topic", "", "", "Remove SNS Topic")
	_applicationinsightsCmd.Flags().StringVarP(&_applicationinsightsResourceARN, "resource-arn", "", "", "Resource ARN")
	_applicationinsightsCmd.Flags().StringVarP(&_applicationinsightsResourceGroupName, "resource-group-name", "", "", "Resource Group Name")
	_applicationinsightsCmd.Flags().StringSliceVarP(&_applicationinsightsResourceList, "resource-list", "", nil, "Resource List")
	_applicationinsightsCmd.Flags().StringVarP(&_applicationinsightsSNSNotificationArn, "sns-notification-arn", "", "", "SNS Notification ARN")
	_applicationinsightsCmd.Flags().StringVarP(&_applicationinsightsStartTime, "start-time", "", "", "Start Time")
	_applicationinsightsCmd.Flags().StringSliceVarP(&_applicationinsightsTagKeys, "tag-keys", "", nil, "Tag Keys")
	_applicationinsightsCmd.Flags().StringVarP(&_applicationinsightsTags, "tags", "", "", "Tags")
	_applicationinsightsCmd.Flags().StringVarP(&_applicationinsightsTier, "tier", "", "", "Tier")
	_applicationinsightsCmd.Flags().StringVarP(&_applicationinsightsUpdateStatus, "update-status", "", "", "Update Status")
	_applicationinsightsCmd.Flags().StringVarP(&_applicationinsightsVisibility, "visibility", "", "", "Visibility")
	_applicationinsightsCmd.Flags().StringVarP(&_applicationinsightsWorkloadConfiguration, "workload-configuration", "", "", "Workload Configuration")
	_applicationinsightsCmd.Flags().StringVarP(&_applicationinsightsWorkloadId, "workload-id", "", "", "Workload ID")
	_applicationinsightsCmd.Flags().StringVarP(&_applicationinsightsWorkloadName, "workload-name", "", "", "Workload Name")

	_applicationinsightsCmd.Flags().BoolVarP(&_applicationinsightsAddWorkload, "add-workload", "", false, "Add Workload")
	_applicationinsightsCmd.Flags().BoolVarP(&_applicationinsightsCreateApplication, "create-application", "", false, "Create Application")
	_applicationinsightsCmd.Flags().BoolVarP(&_applicationinsightsCreateComponent, "create-component", "", false, "Create Component")
	_applicationinsightsCmd.Flags().BoolVarP(&_applicationinsightsCreateLogPattern, "create-log-pattern", "", false, "Create Log Pattern")
	_applicationinsightsCmd.Flags().BoolVarP(&_applicationinsightsDeleteApplication, "delete-application", "", false, "Delete Application")
	_applicationinsightsCmd.Flags().BoolVarP(&_applicationinsightsDeleteComponent, "delete-component", "", false, "Delete Component")
	_applicationinsightsCmd.Flags().BoolVarP(&_applicationinsightsDeleteLogPattern, "delete-log-pattern", "", false, "Delete Log Pattern")
	_applicationinsightsCmd.Flags().BoolVarP(&_applicationinsightsDescribeApplication, "describe-application", "", false, "Describe Application")
	_applicationinsightsCmd.Flags().BoolVarP(&_applicationinsightsDescribeComponent, "describe-component", "", false, "Describe Component")
	_applicationinsightsCmd.Flags().BoolVarP(&_applicationinsightsDescribeComponentConfiguration, "describe-component-configuration", "", false, "Describe Component Configuration")
	_applicationinsightsCmd.Flags().BoolVarP(&_applicationinsightsDescribeComponentConfigurationRecommendation, "describe-component-configuration-recommendation", "", false, "Describe Component Configuration Recommendation")
	_applicationinsightsCmd.Flags().BoolVarP(&_applicationinsightsDescribeLogPattern, "describe-log-pattern", "", false, "Describe Log Pattern")
	_applicationinsightsCmd.Flags().BoolVarP(&_applicationinsightsDescribeObservation, "describe-observation", "", false, "Describe Observation")
	_applicationinsightsCmd.Flags().BoolVarP(&_applicationinsightsDescribeProblem, "describe-problem", "", false, "Describe Problem")
	_applicationinsightsCmd.Flags().BoolVarP(&_applicationinsightsDescribeProblemObservations, "describe-problem-observations", "", false, "Describe Problem Observations")
	_applicationinsightsCmd.Flags().BoolVarP(&_applicationinsightsDescribeWorkload, "describe-workload", "", false, "Describe Workload")
	_applicationinsightsCmd.Flags().BoolVarP(&_applicationinsightsListApplications, "list-applications", "", false, "List Applications")
	_applicationinsightsCmd.Flags().BoolVarP(&_applicationinsightsListComponents, "list-components", "", false, "List Components")
	_applicationinsightsCmd.Flags().BoolVarP(&_applicationinsightsListConfigurationHistory, "list-configuration-history", "", false, "List Configuration History")
	_applicationinsightsCmd.Flags().BoolVarP(&_applicationinsightsListLogPatternSets, "list-log-pattern-sets", "", false, "List Log Pattern Sets")
	_applicationinsightsCmd.Flags().BoolVarP(&_applicationinsightsListLogPatterns, "list-log-patterns", "", false, "List Log Patterns")
	_applicationinsightsCmd.Flags().BoolVarP(&_applicationinsightsListProblems, "list-problems", "", false, "List Problems")
	_applicationinsightsCmd.Flags().BoolVarP(&_applicationinsightsListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_applicationinsightsCmd.Flags().BoolVarP(&_applicationinsightsListWorkloads, "list-workloads", "", false, "List Workloads")
	_applicationinsightsCmd.Flags().BoolVarP(&_applicationinsightsRemoveWorkload, "remove-workload", "", false, "Remove Workload")
	_applicationinsightsCmd.Flags().BoolVarP(&_applicationinsightsTagResource, "tag-resource", "", false, "Tag Resource")
	_applicationinsightsCmd.Flags().BoolVarP(&_applicationinsightsUntagResource, "untag-resource", "", false, "Untag Resource")
	_applicationinsightsCmd.Flags().BoolVarP(&_applicationinsightsUpdateApplication, "update-application", "", false, "Update Application")
	_applicationinsightsCmd.Flags().BoolVarP(&_applicationinsightsUpdateComponent, "update-component", "", false, "Update Component")
	_applicationinsightsCmd.Flags().BoolVarP(&_applicationinsightsUpdateComponentConfiguration, "update-component-configuration", "", false, "Update Component Configuration")
	_applicationinsightsCmd.Flags().BoolVarP(&_applicationinsightsUpdateLogPattern, "update-log-pattern", "", false, "Update Log Pattern")
	_applicationinsightsCmd.Flags().BoolVarP(&_applicationinsightsUpdateProblem, "update-problem", "", false, "Update Problem")
	_applicationinsightsCmd.Flags().BoolVarP(&_applicationinsightsUpdateWorkload, "update-workload", "", false, "Update Workload")

}
