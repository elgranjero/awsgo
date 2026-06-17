package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/applicationsignals"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// applicationsignalsCmd represents the applicationsignals command
var _applicationsignalsCmd = &cobra.Command{
	Use:   "applicationsignals",
	Short: "AWS applicationsignals CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := applicationsignals.NewFromConfig(cfg)
		if _applicationsignalsBatchGetServiceLevelObjectiveBudgetReport {
			applicationsignals_BatchGetServiceLevelObjectiveBudgetReport(cfg, client)
			return
		}
		if _applicationsignalsBatchUpdateExclusionWindows {
			applicationsignals_BatchUpdateExclusionWindows(cfg, client)
			return
		}
		if _applicationsignalsCreateServiceLevelObjective {
			applicationsignals_CreateServiceLevelObjective(cfg, client)
			return
		}
		if _applicationsignalsDeleteGroupingConfiguration {
			applicationsignals_DeleteGroupingConfiguration(cfg, client)
			return
		}
		if _applicationsignalsDeleteServiceLevelObjective {
			applicationsignals_DeleteServiceLevelObjective(cfg, client)
			return
		}
		if _applicationsignalsGetService {
			applicationsignals_GetService(cfg, client)
			return
		}
		if _applicationsignalsGetServiceLevelObjective {
			applicationsignals_GetServiceLevelObjective(cfg, client)
			return
		}
		if _applicationsignalsListAuditFindings {
			applicationsignals_ListAuditFindings(cfg, client)
			return
		}
		if _applicationsignalsListEntityEvents {
			applicationsignals_ListEntityEvents(cfg, client)
			return
		}
		if _applicationsignalsListGroupingAttributeDefinitions {
			applicationsignals_ListGroupingAttributeDefinitions(cfg, client)
			return
		}
		if _applicationsignalsListServiceDependencies {
			applicationsignals_ListServiceDependencies(cfg, client)
			return
		}
		if _applicationsignalsListServiceDependents {
			applicationsignals_ListServiceDependents(cfg, client)
			return
		}
		if _applicationsignalsListServiceLevelObjectiveExclusionWindows {
			applicationsignals_ListServiceLevelObjectiveExclusionWindows(cfg, client)
			return
		}
		if _applicationsignalsListServiceLevelObjectives {
			applicationsignals_ListServiceLevelObjectives(cfg, client)
			return
		}
		if _applicationsignalsListServiceOperations {
			applicationsignals_ListServiceOperations(cfg, client)
			return
		}
		if _applicationsignalsListServiceStates {
			applicationsignals_ListServiceStates(cfg, client)
			return
		}
		if _applicationsignalsListServices {
			applicationsignals_ListServices(cfg, client)
			return
		}
		if _applicationsignalsListTagsForResource {
			applicationsignals_ListTagsForResource(cfg, client)
			return
		}
		if _applicationsignalsPutGroupingConfiguration {
			applicationsignals_PutGroupingConfiguration(cfg, client)
			return
		}
		if _applicationsignalsStartDiscovery {
			applicationsignals_StartDiscovery(cfg, client)
			return
		}
		if _applicationsignalsTagResource {
			applicationsignals_TagResource(cfg, client)
			return
		}
		if _applicationsignalsUntagResource {
			applicationsignals_UntagResource(cfg, client)
			return
		}
		if _applicationsignalsUpdateServiceLevelObjective {
			applicationsignals_UpdateServiceLevelObjective(cfg, client)
			return
		}

	},
}

var (
	_applicationsignalsBatchGetServiceLevelObjectiveBudgetReport bool
	_applicationsignalsBatchUpdateExclusionWindows               bool
	_applicationsignalsCreateServiceLevelObjective               bool
	_applicationsignalsDeleteGroupingConfiguration               bool
	_applicationsignalsDeleteServiceLevelObjective               bool
	_applicationsignalsGetService                                bool
	_applicationsignalsGetServiceLevelObjective                  bool
	_applicationsignalsListAuditFindings                         bool
	_applicationsignalsListEntityEvents                          bool
	_applicationsignalsListGroupingAttributeDefinitions          bool
	_applicationsignalsListServiceDependencies                   bool
	_applicationsignalsListServiceDependents                     bool
	_applicationsignalsListServiceLevelObjectiveExclusionWindows bool
	_applicationsignalsListServiceLevelObjectives                bool
	_applicationsignalsListServiceOperations                     bool
	_applicationsignalsListServiceStates                         bool
	_applicationsignalsListServices                              bool
	_applicationsignalsListTagsForResource                       bool
	_applicationsignalsPutGroupingConfiguration                  bool
	_applicationsignalsStartDiscovery                            bool
	_applicationsignalsTagResource                               bool
	_applicationsignalsUntagResource                             bool
	_applicationsignalsUpdateServiceLevelObjective               bool

	_applicationsignalsAddExclusionWindows          string
	_applicationsignalsAttributeFilters             string
	_applicationsignalsAuditTargets                 string
	_applicationsignalsAuditors                     []string
	_applicationsignalsAwsAccountId                 string
	_applicationsignalsBurnRateConfigurations       string
	_applicationsignalsDependencyConfig             string
	_applicationsignalsDescription                  string
	_applicationsignalsDetailLevel                  string
	_applicationsignalsEndTime                      string
	_applicationsignalsEntity                       string
	_applicationsignalsGoal                         string
	_applicationsignalsGroupingAttributeDefinitions string
	_applicationsignalsId                           string
	_applicationsignalsIncludeLinkedAccounts        string
	_applicationsignalsKeyAttributes                string
	_applicationsignalsMaxResults                   string
	_applicationsignalsMetricSourceTypes            string
	_applicationsignalsName                         string
	_applicationsignalsNextToken                    string
	_applicationsignalsOperationName                string
	_applicationsignalsRemoveExclusionWindows       string
	_applicationsignalsRequestBasedSliConfig        string
	_applicationsignalsResourceArn                  string
	_applicationsignalsSliConfig                    string
	_applicationsignalsSloIds                       []string
	_applicationsignalsSloOwnerAwsAccountId         string
	_applicationsignalsStartTime                    string
	_applicationsignalsTagKeys                      []string
	_applicationsignalsTags                         string
	_applicationsignalsTimestamp                    string
)

// Use this operation to retrieve one or more service level objective (SLO) budget
// reports.
//
// An error budget is the amount of time or requests in an unhealthy state that
// your service can accumulate during an interval before your overall SLO budget
// health is breached and the SLO is considered to be unmet. For example, an SLO
// with a threshold of 99.95% and a monthly interval translates to an error budget
// of 21.9 minutes of downtime in a 30-day month.
//
// Budget reports include a health indicator, the attainment value, and remaining
// budget.
//
// For more information about SLO error budgets, see [SLO concepts].
//
// [SLO concepts]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-ServiceLevelObjectives.html#CloudWatch-ServiceLevelObjectives-concepts
func applicationsignals_BatchGetServiceLevelObjectiveBudgetReport(cfg aws.Config, client *applicationsignals.Client) {
	input := &applicationsignals.BatchGetServiceLevelObjectiveBudgetReportInput{
		// SloIds: []string, // Required
		// Timestamp: *time.Time, // Required
	}

	if len(_applicationsignalsSloIds) > 0 {
		input.SloIds = append([]string(nil), _applicationsignalsSloIds...)
	}
	if len(_applicationsignalsTimestamp) > 0 {
		if err := assignInputField(input, "Timestamp", _applicationsignalsTimestamp); err != nil {
			log.Errorf("invalid --timestamp: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchGetServiceLevelObjectiveBudgetReport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Add or remove time window exclusions for one or more Service Level Objectives
// (SLOs).
func applicationsignals_BatchUpdateExclusionWindows(cfg aws.Config, client *applicationsignals.Client) {
	input := &applicationsignals.BatchUpdateExclusionWindowsInput{
		// SloIds: []string, // Required
	}

	if len(_applicationsignalsSloIds) > 0 {
		input.SloIds = append([]string(nil), _applicationsignalsSloIds...)
	}
	if len(_applicationsignalsAddExclusionWindows) > 0 {
		if err := assignInputField(input, "AddExclusionWindows", _applicationsignalsAddExclusionWindows); err != nil {
			log.Errorf("invalid --add-exclusion-windows: %s", err.Error())
			return
		}
	}
	if len(_applicationsignalsRemoveExclusionWindows) > 0 {
		if err := assignInputField(input, "RemoveExclusionWindows", _applicationsignalsRemoveExclusionWindows); err != nil {
			log.Errorf("invalid --remove-exclusion-windows: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchUpdateExclusionWindows(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a service level objective (SLO), which can help you ensure that your
// critical business operations are meeting customer expectations. Use SLOs to set
// and track specific target levels for the reliability and availability of your
// applications and services. SLOs use service level indicators (SLIs) to calculate
// whether the application is performing at the level that you want.
//
// Create an SLO to set a target for a service or operation’s availability or
// latency. CloudWatch measures this target frequently you can find whether it has
// been breached.
//
// The target performance quality that is defined for an SLO is the attainment
// goal.
//
// You can set SLO targets for your applications that are discovered by
// Application Signals, using critical metrics such as latency and availability.
// You can also set SLOs against any CloudWatch metric or math expression that
// produces a time series.
//
// You can't create an SLO for a service operation that was discovered by
// Application Signals until after that operation has reported standard metrics to
// Application Signals.
//
// When you create an SLO, you specify whether it is a period-based SLO or a
// request-based SLO. Each type of SLO has a different way of evaluating your
// application's performance against its attainment goal.
//
// - A period-based SLO uses defined periods of time within a specified total
// time interval. For each period of time, Application Signals determines whether
// the application met its goal. The attainment rate is calculated as the number
// of good periods/number of total periods .
//
// For example, for a period-based SLO, meeting an attainment goal of 99.9% means
//
// that within your interval, your application must meet its performance goal
// during at least 99.9% of the time periods.
//
// - A request-based SLO doesn't use pre-defined periods of time. Instead, the
// SLO measures number of good requests/number of total requests during the
// interval. At any time, you can find the ratio of good requests to total requests
// for the interval up to the time stamp that you specify, and measure that ratio
// against the goal set in your SLO.
//
// After you have created an SLO, you can retrieve error budget reports for it. An
// error budget is the amount of time or amount of requests that your application
// can be non-compliant with the SLO's goal, and still have your application meet
// the goal.
//
// - For a period-based SLO, the error budget starts at a number defined by the
// highest number of periods that can fail to meet the threshold, while still
// meeting the overall goal. The remaining error budget decreases with every failed
// period that is recorded. The error budget within one interval can never
// increase.
//
// For example, an SLO with a threshold that 99.95% of requests must be completed
//
// under 2000ms every month translates to an error budget of 21.9 minutes of
// downtime per month.
//
// - For a request-based SLO, the remaining error budget is dynamic and can
// increase or decrease, depending on the ratio of good requests to total requests.
//
// For more information about SLOs, see [Service level objectives (SLOs)].
//
// When you perform a CreateServiceLevelObjective operation, Application Signals
// creates the AWSServiceRoleForCloudWatchApplicationSignals service-linked role,
// if it doesn't already exist in your account. This service- linked role has the
// following permissions:
//
// - xray:GetServiceGraph
//
// - logs:StartQuery
//
// - logs:GetQueryResults
//
// - cloudwatch:GetMetricData
//
// - cloudwatch:ListMetrics
//
// - tag:GetResources
//
// - autoscaling:DescribeAutoScalingGroups
//
// [Service level objectives (SLOs)]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-ServiceLevelObjectives.html
func applicationsignals_CreateServiceLevelObjective(cfg aws.Config, client *applicationsignals.Client) {
	input := &applicationsignals.CreateServiceLevelObjectiveInput{
		// Name: *string, // Required
	}

	if len(_applicationsignalsName) > 0 {
		input.Name = aws.String(_applicationsignalsName)
	}
	if len(_applicationsignalsBurnRateConfigurations) > 0 {
		if err := assignInputField(input, "BurnRateConfigurations", _applicationsignalsBurnRateConfigurations); err != nil {
			log.Errorf("invalid --burn-rate-configurations: %s", err.Error())
			return
		}
	}
	if len(_applicationsignalsDescription) > 0 {
		input.Description = aws.String(_applicationsignalsDescription)
	}
	if len(_applicationsignalsGoal) > 0 {
		if err := assignInputField(input, "Goal", _applicationsignalsGoal); err != nil {
			log.Errorf("invalid --goal: %s", err.Error())
			return
		}
	}
	if len(_applicationsignalsRequestBasedSliConfig) > 0 {
		if err := assignInputField(input, "RequestBasedSliConfig", _applicationsignalsRequestBasedSliConfig); err != nil {
			log.Errorf("invalid --request-based-sli-config: %s", err.Error())
			return
		}
	}
	if len(_applicationsignalsSliConfig) > 0 {
		if err := assignInputField(input, "SliConfig", _applicationsignalsSliConfig); err != nil {
			log.Errorf("invalid --sli-config: %s", err.Error())
			return
		}
	}
	if len(_applicationsignalsTags) > 0 {
		if err := assignInputField(input, "Tags", _applicationsignalsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateServiceLevelObjective(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the grouping configuration for this account. This removes all custom
// grouping attribute definitions that were previously configured.
func applicationsignals_DeleteGroupingConfiguration(cfg aws.Config, client *applicationsignals.Client) {
	input := &applicationsignals.DeleteGroupingConfigurationInput{}

	if resp, err := client.DeleteGroupingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified service level objective.
func applicationsignals_DeleteServiceLevelObjective(cfg aws.Config, client *applicationsignals.Client) {
	input := &applicationsignals.DeleteServiceLevelObjectiveInput{
		// Id: *string, // Required
	}

	if len(_applicationsignalsId) > 0 {
		input.Id = aws.String(_applicationsignalsId)
	}

	if resp, err := client.DeleteServiceLevelObjective(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a service discovered by Application Signals.
func applicationsignals_GetService(cfg aws.Config, client *applicationsignals.Client) {
	input := &applicationsignals.GetServiceInput{
		// EndTime: *time.Time, // Required
		// KeyAttributes: map[string]string, // Required
		// StartTime: *time.Time, // Required
	}

	if len(_applicationsignalsEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _applicationsignalsEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_applicationsignalsKeyAttributes) > 0 {
		if err := assignInputField(input, "KeyAttributes", _applicationsignalsKeyAttributes); err != nil {
			log.Errorf("invalid --key-attributes: %s", err.Error())
			return
		}
	}
	if len(_applicationsignalsStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _applicationsignalsStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetService(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about one SLO created in the account.
func applicationsignals_GetServiceLevelObjective(cfg aws.Config, client *applicationsignals.Client) {
	input := &applicationsignals.GetServiceLevelObjectiveInput{
		// Id: *string, // Required
	}

	if len(_applicationsignalsId) > 0 {
		input.Id = aws.String(_applicationsignalsId)
	}

	if resp, err := client.GetServiceLevelObjective(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of audit findings that provide automated analysis of service
// behavior and root cause analysis. These findings help identify the most
// significant observations about your services, including performance issues,
// anomalies, and potential problems. The findings are generated using heuristic
// algorithms based on established troubleshooting patterns.
func applicationsignals_ListAuditFindings(cfg aws.Config, client *applicationsignals.Client) {
	input := &applicationsignals.ListAuditFindingsInput{
		// AuditTargets: []types.AuditTarget, // Required
		// EndTime: *time.Time, // Required
		// StartTime: *time.Time, // Required
	}

	if len(_applicationsignalsAuditTargets) > 0 {
		if err := assignInputField(input, "AuditTargets", _applicationsignalsAuditTargets); err != nil {
			log.Errorf("invalid --audit-targets: %s", err.Error())
			return
		}
	}
	if len(_applicationsignalsEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _applicationsignalsEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_applicationsignalsStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _applicationsignalsStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_applicationsignalsAuditors) > 0 {
		input.Auditors = append([]string(nil), _applicationsignalsAuditors...)
	}
	if len(_applicationsignalsDetailLevel) > 0 {
		if err := assignInputField(input, "DetailLevel", _applicationsignalsDetailLevel); err != nil {
			log.Errorf("invalid --detail-level: %s", err.Error())
			return
		}
	}
	if len(_applicationsignalsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _applicationsignalsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_applicationsignalsNextToken) > 0 {
		input.NextToken = aws.String(_applicationsignalsNextToken)
	}

	if resp, err := client.ListAuditFindings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of change events for a specific entity, such as deployments,
// configuration changes, or other state-changing activities. This operation helps
// track the history of changes that may have affected service performance.
func applicationsignals_ListEntityEvents(cfg aws.Config, client *applicationsignals.Client) {
	input := &applicationsignals.ListEntityEventsInput{
		// EndTime: *time.Time, // Required
		// Entity: map[string]string, // Required
		// StartTime: *time.Time, // Required
	}

	if len(_applicationsignalsEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _applicationsignalsEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_applicationsignalsEntity) > 0 {
		if err := assignInputField(input, "Entity", _applicationsignalsEntity); err != nil {
			log.Errorf("invalid --entity: %s", err.Error())
			return
		}
	}
	if len(_applicationsignalsStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _applicationsignalsStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_applicationsignalsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _applicationsignalsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_applicationsignalsNextToken) > 0 {
		input.NextToken = aws.String(_applicationsignalsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEntityEvents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*applicationsignals.ListEntityEventsOutput
	p := applicationsignals.NewListEntityEventsPaginator(client, input)
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

// Returns the current grouping configuration for this account, including all
// custom grouping attribute definitions that have been configured. These
// definitions determine how services are logically grouped based on telemetry
// attributes, Amazon Web Services tags, or predefined mappings.
func applicationsignals_ListGroupingAttributeDefinitions(cfg aws.Config, client *applicationsignals.Client) {
	input := &applicationsignals.ListGroupingAttributeDefinitionsInput{}

	if len(_applicationsignalsAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_applicationsignalsAwsAccountId)
	}
	if len(_applicationsignalsIncludeLinkedAccounts) > 0 {
		if err := assignInputField(input, "IncludeLinkedAccounts", _applicationsignalsIncludeLinkedAccounts); err != nil {
			log.Errorf("invalid --include-linked-accounts: %s", err.Error())
			return
		}
	}
	if len(_applicationsignalsNextToken) > 0 {
		input.NextToken = aws.String(_applicationsignalsNextToken)
	}

	if resp, err := client.ListGroupingAttributeDefinitions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of service dependencies of the service that you specify. A
// dependency is an infrastructure component that an operation of this service
// connects with. Dependencies can include Amazon Web Services services, Amazon Web
// Services resources, and third-party services.
func applicationsignals_ListServiceDependencies(cfg aws.Config, client *applicationsignals.Client) {
	input := &applicationsignals.ListServiceDependenciesInput{
		// EndTime: *time.Time, // Required
		// KeyAttributes: map[string]string, // Required
		// StartTime: *time.Time, // Required
	}

	if len(_applicationsignalsEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _applicationsignalsEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_applicationsignalsKeyAttributes) > 0 {
		if err := assignInputField(input, "KeyAttributes", _applicationsignalsKeyAttributes); err != nil {
			log.Errorf("invalid --key-attributes: %s", err.Error())
			return
		}
	}
	if len(_applicationsignalsStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _applicationsignalsStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_applicationsignalsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _applicationsignalsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_applicationsignalsNextToken) > 0 {
		input.NextToken = aws.String(_applicationsignalsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListServiceDependencies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*applicationsignals.ListServiceDependenciesOutput
	p := applicationsignals.NewListServiceDependenciesPaginator(client, input)
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

// Returns the list of dependents that invoked the specified service during the
// provided time range. Dependents include other services, CloudWatch Synthetics
// canaries, and clients that are instrumented with CloudWatch RUM app monitors.
func applicationsignals_ListServiceDependents(cfg aws.Config, client *applicationsignals.Client) {
	input := &applicationsignals.ListServiceDependentsInput{
		// EndTime: *time.Time, // Required
		// KeyAttributes: map[string]string, // Required
		// StartTime: *time.Time, // Required
	}

	if len(_applicationsignalsEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _applicationsignalsEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_applicationsignalsKeyAttributes) > 0 {
		if err := assignInputField(input, "KeyAttributes", _applicationsignalsKeyAttributes); err != nil {
			log.Errorf("invalid --key-attributes: %s", err.Error())
			return
		}
	}
	if len(_applicationsignalsStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _applicationsignalsStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_applicationsignalsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _applicationsignalsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_applicationsignalsNextToken) > 0 {
		input.NextToken = aws.String(_applicationsignalsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListServiceDependents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*applicationsignals.ListServiceDependentsOutput
	p := applicationsignals.NewListServiceDependentsPaginator(client, input)
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

// Retrieves all exclusion windows configured for a specific SLO.
func applicationsignals_ListServiceLevelObjectiveExclusionWindows(cfg aws.Config, client *applicationsignals.Client) {
	input := &applicationsignals.ListServiceLevelObjectiveExclusionWindowsInput{
		// Id: *string, // Required
	}

	if len(_applicationsignalsId) > 0 {
		input.Id = aws.String(_applicationsignalsId)
	}
	if len(_applicationsignalsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _applicationsignalsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_applicationsignalsNextToken) > 0 {
		input.NextToken = aws.String(_applicationsignalsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListServiceLevelObjectiveExclusionWindows(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*applicationsignals.ListServiceLevelObjectiveExclusionWindowsOutput
	p := applicationsignals.NewListServiceLevelObjectiveExclusionWindowsPaginator(client, input)
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

// Returns a list of SLOs created in this account.
func applicationsignals_ListServiceLevelObjectives(cfg aws.Config, client *applicationsignals.Client) {
	input := &applicationsignals.ListServiceLevelObjectivesInput{}

	if len(_applicationsignalsDependencyConfig) > 0 {
		if err := assignInputField(input, "DependencyConfig", _applicationsignalsDependencyConfig); err != nil {
			log.Errorf("invalid --dependency-config: %s", err.Error())
			return
		}
	}
	if len(_applicationsignalsIncludeLinkedAccounts) > 0 {
		if err := assignInputField(input, "IncludeLinkedAccounts", _applicationsignalsIncludeLinkedAccounts); err != nil {
			log.Errorf("invalid --include-linked-accounts: %s", err.Error())
			return
		}
	}
	if len(_applicationsignalsKeyAttributes) > 0 {
		if err := assignInputField(input, "KeyAttributes", _applicationsignalsKeyAttributes); err != nil {
			log.Errorf("invalid --key-attributes: %s", err.Error())
			return
		}
	}
	if len(_applicationsignalsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _applicationsignalsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_applicationsignalsMetricSourceTypes) > 0 {
		if err := assignInputField(input, "MetricSourceTypes", _applicationsignalsMetricSourceTypes); err != nil {
			log.Errorf("invalid --metric-source-types: %s", err.Error())
			return
		}
	}
	if len(_applicationsignalsNextToken) > 0 {
		input.NextToken = aws.String(_applicationsignalsNextToken)
	}
	if len(_applicationsignalsOperationName) > 0 {
		input.OperationName = aws.String(_applicationsignalsOperationName)
	}
	if len(_applicationsignalsSloOwnerAwsAccountId) > 0 {
		input.SloOwnerAwsAccountId = aws.String(_applicationsignalsSloOwnerAwsAccountId)
	}

	if disablePaginator() {
		if resp, err := client.ListServiceLevelObjectives(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*applicationsignals.ListServiceLevelObjectivesOutput
	p := applicationsignals.NewListServiceLevelObjectivesPaginator(client, input)
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

// Returns a list of the operations of this service that have been discovered by
// Application Signals. Only the operations that were invoked during the specified
// time range are returned.
func applicationsignals_ListServiceOperations(cfg aws.Config, client *applicationsignals.Client) {
	input := &applicationsignals.ListServiceOperationsInput{
		// EndTime: *time.Time, // Required
		// KeyAttributes: map[string]string, // Required
		// StartTime: *time.Time, // Required
	}

	if len(_applicationsignalsEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _applicationsignalsEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_applicationsignalsKeyAttributes) > 0 {
		if err := assignInputField(input, "KeyAttributes", _applicationsignalsKeyAttributes); err != nil {
			log.Errorf("invalid --key-attributes: %s", err.Error())
			return
		}
	}
	if len(_applicationsignalsStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _applicationsignalsStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_applicationsignalsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _applicationsignalsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_applicationsignalsNextToken) > 0 {
		input.NextToken = aws.String(_applicationsignalsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListServiceOperations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*applicationsignals.ListServiceOperationsOutput
	p := applicationsignals.NewListServiceOperationsPaginator(client, input)
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

// Returns information about the last deployment and other change states of
// services. This API provides visibility into recent changes that may have
// affected service performance, helping with troubleshooting and change
// correlation.
func applicationsignals_ListServiceStates(cfg aws.Config, client *applicationsignals.Client) {
	input := &applicationsignals.ListServiceStatesInput{
		// EndTime: *time.Time, // Required
		// StartTime: *time.Time, // Required
	}

	if len(_applicationsignalsEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _applicationsignalsEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_applicationsignalsStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _applicationsignalsStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_applicationsignalsAttributeFilters) > 0 {
		if err := assignInputField(input, "AttributeFilters", _applicationsignalsAttributeFilters); err != nil {
			log.Errorf("invalid --attribute-filters: %s", err.Error())
			return
		}
	}
	if len(_applicationsignalsAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_applicationsignalsAwsAccountId)
	}
	if len(_applicationsignalsIncludeLinkedAccounts) > 0 {
		if err := assignInputField(input, "IncludeLinkedAccounts", _applicationsignalsIncludeLinkedAccounts); err != nil {
			log.Errorf("invalid --include-linked-accounts: %s", err.Error())
			return
		}
	}
	if len(_applicationsignalsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _applicationsignalsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_applicationsignalsNextToken) > 0 {
		input.NextToken = aws.String(_applicationsignalsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListServiceStates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*applicationsignals.ListServiceStatesOutput
	p := applicationsignals.NewListServiceStatesPaginator(client, input)
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

// Returns a list of services that have been discovered by Application Signals. A
// service represents a minimum logical and transactional unit that completes a
// business function. Services are discovered through Application Signals
// instrumentation.
func applicationsignals_ListServices(cfg aws.Config, client *applicationsignals.Client) {
	input := &applicationsignals.ListServicesInput{
		// EndTime: *time.Time, // Required
		// StartTime: *time.Time, // Required
	}

	if len(_applicationsignalsEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _applicationsignalsEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_applicationsignalsStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _applicationsignalsStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_applicationsignalsAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_applicationsignalsAwsAccountId)
	}
	if len(_applicationsignalsIncludeLinkedAccounts) > 0 {
		if err := assignInputField(input, "IncludeLinkedAccounts", _applicationsignalsIncludeLinkedAccounts); err != nil {
			log.Errorf("invalid --include-linked-accounts: %s", err.Error())
			return
		}
	}
	if len(_applicationsignalsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _applicationsignalsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_applicationsignalsNextToken) > 0 {
		input.NextToken = aws.String(_applicationsignalsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListServices(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*applicationsignals.ListServicesOutput
	p := applicationsignals.NewListServicesPaginator(client, input)
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

// Displays the tags associated with a CloudWatch resource. Tags can be assigned
// to service level objectives.
func applicationsignals_ListTagsForResource(cfg aws.Config, client *applicationsignals.Client) {
	input := &applicationsignals.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_applicationsignalsResourceArn) > 0 {
		input.ResourceArn = aws.String(_applicationsignalsResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates the grouping configuration for this account. This operation
// allows you to define custom grouping attributes that determine how services are
// logically grouped based on telemetry attributes, Amazon Web Services tags, or
// predefined mappings. These grouping attributes can then be used to organize and
// filter services in the Application Signals console and APIs.
func applicationsignals_PutGroupingConfiguration(cfg aws.Config, client *applicationsignals.Client) {
	input := &applicationsignals.PutGroupingConfigurationInput{
		// GroupingAttributeDefinitions: []types.GroupingAttributeDefinition, // Required
	}

	if len(_applicationsignalsGroupingAttributeDefinitions) > 0 {
		if err := assignInputField(input, "GroupingAttributeDefinitions", _applicationsignalsGroupingAttributeDefinitions); err != nil {
			log.Errorf("invalid --grouping-attribute-definitions: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutGroupingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables this Amazon Web Services account to be able to use CloudWatch
// Application Signals by creating the
// AWSServiceRoleForCloudWatchApplicationSignals service-linked role. This service-
// linked role has the following permissions:
//
// - xray:GetServiceGraph
//
// - logs:StartQuery
//
// - logs:GetQueryResults
//
// - cloudwatch:GetMetricData
//
// - cloudwatch:ListMetrics
//
// - tag:GetResources
//
// - autoscaling:DescribeAutoScalingGroups
//
// A service-linked CloudTrail event channel is created to process CloudTrail
// events and return change event information. This includes last deployment time,
// userName, eventName, and other event metadata.
//
// After completing this step, you still need to instrument your Java and Python
// applications to send data to Application Signals. For more information, see [Enabling Application Signals].
//
// [Enabling Application Signals]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Application-Signals-Enable.html
func applicationsignals_StartDiscovery(cfg aws.Config, client *applicationsignals.Client) {
	input := &applicationsignals.StartDiscoveryInput{}

	if resp, err := client.StartDiscovery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns one or more tags (key-value pairs) to the specified CloudWatch
// resource, such as a service level objective.
//
// Tags can help you organize and categorize your resources. You can also use them
// to scope user permissions by granting a user permission to access or change only
// resources with certain tag values.
//
// Tags don't have any semantic meaning to Amazon Web Services and are interpreted
// strictly as strings of characters.
//
// You can use the TagResource action with an alarm that already has tags. If you
// specify a new tag key for the alarm, this tag is appended to the list of tags
// associated with the alarm. If you specify a tag key that is already associated
// with the alarm, the new tag value that you specify replaces the previous value
// for that tag.
//
// You can associate as many as 50 tags with a CloudWatch resource.
func applicationsignals_TagResource(cfg aws.Config, client *applicationsignals.Client) {
	input := &applicationsignals.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_applicationsignalsResourceArn) > 0 {
		input.ResourceArn = aws.String(_applicationsignalsResourceArn)
	}
	if len(_applicationsignalsTags) > 0 {
		if err := assignInputField(input, "Tags", _applicationsignalsTags); err != nil {
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
func applicationsignals_UntagResource(cfg aws.Config, client *applicationsignals.Client) {
	input := &applicationsignals.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_applicationsignalsResourceArn) > 0 {
		input.ResourceArn = aws.String(_applicationsignalsResourceArn)
	}
	if len(_applicationsignalsTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _applicationsignalsTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing service level objective (SLO). If you omit parameters, the
// previous values of those parameters are retained.
//
// You cannot change from a period-based SLO to a request-based SLO, or change
// from a request-based SLO to a period-based SLO.
func applicationsignals_UpdateServiceLevelObjective(cfg aws.Config, client *applicationsignals.Client) {
	input := &applicationsignals.UpdateServiceLevelObjectiveInput{
		// Id: *string, // Required
	}

	if len(_applicationsignalsId) > 0 {
		input.Id = aws.String(_applicationsignalsId)
	}
	if len(_applicationsignalsBurnRateConfigurations) > 0 {
		if err := assignInputField(input, "BurnRateConfigurations", _applicationsignalsBurnRateConfigurations); err != nil {
			log.Errorf("invalid --burn-rate-configurations: %s", err.Error())
			return
		}
	}
	if len(_applicationsignalsDescription) > 0 {
		input.Description = aws.String(_applicationsignalsDescription)
	}
	if len(_applicationsignalsGoal) > 0 {
		if err := assignInputField(input, "Goal", _applicationsignalsGoal); err != nil {
			log.Errorf("invalid --goal: %s", err.Error())
			return
		}
	}
	if len(_applicationsignalsRequestBasedSliConfig) > 0 {
		if err := assignInputField(input, "RequestBasedSliConfig", _applicationsignalsRequestBasedSliConfig); err != nil {
			log.Errorf("invalid --request-based-sli-config: %s", err.Error())
			return
		}
	}
	if len(_applicationsignalsSliConfig) > 0 {
		if err := assignInputField(input, "SliConfig", _applicationsignalsSliConfig); err != nil {
			log.Errorf("invalid --sli-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateServiceLevelObjective(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_applicationsignalsCmd)
	_applicationsignalsCmd.Flags().SortFlags = false

	_applicationsignalsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_applicationsignalsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_applicationsignalsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_applicationsignalsCmd.Flags().StringVarP(&_applicationsignalsAddExclusionWindows, "add-exclusion-windows", "", "", "Add Exclusion Windows")
	_applicationsignalsCmd.Flags().StringVarP(&_applicationsignalsAttributeFilters, "attribute-filters", "", "", "Attribute Filters")
	_applicationsignalsCmd.Flags().StringVarP(&_applicationsignalsAuditTargets, "audit-targets", "", "", "Audit Targets")
	_applicationsignalsCmd.Flags().StringSliceVarP(&_applicationsignalsAuditors, "auditors", "", nil, "Auditors")
	_applicationsignalsCmd.Flags().StringVarP(&_applicationsignalsAwsAccountId, "aws-account-id", "", "", "AWS Account ID")
	_applicationsignalsCmd.Flags().StringVarP(&_applicationsignalsBurnRateConfigurations, "burn-rate-configurations", "", "", "Burn Rate Configurations")
	_applicationsignalsCmd.Flags().StringVarP(&_applicationsignalsDependencyConfig, "dependency-config", "", "", "Dependency Config")
	_applicationsignalsCmd.Flags().StringVarP(&_applicationsignalsDescription, "description", "", "", "Description")
	_applicationsignalsCmd.Flags().StringVarP(&_applicationsignalsDetailLevel, "detail-level", "", "", "Detail Level")
	_applicationsignalsCmd.Flags().StringVarP(&_applicationsignalsEndTime, "end-time", "", "", "End Time")
	_applicationsignalsCmd.Flags().StringVarP(&_applicationsignalsEntity, "entity", "", "", "Entity")
	_applicationsignalsCmd.Flags().StringVarP(&_applicationsignalsGoal, "goal", "", "", "Goal")
	_applicationsignalsCmd.Flags().StringVarP(&_applicationsignalsGroupingAttributeDefinitions, "grouping-attribute-definitions", "", "", "Grouping Attribute Definitions")
	_applicationsignalsCmd.Flags().StringVarP(&_applicationsignalsId, "id", "", "", "ID")
	_applicationsignalsCmd.Flags().StringVarP(&_applicationsignalsIncludeLinkedAccounts, "include-linked-accounts", "", "", "Include Linked Accounts")
	_applicationsignalsCmd.Flags().StringVarP(&_applicationsignalsKeyAttributes, "key-attributes", "", "", "Key Attributes")
	_applicationsignalsCmd.Flags().StringVarP(&_applicationsignalsMaxResults, "max-results", "", "", "Max Results")
	_applicationsignalsCmd.Flags().StringVarP(&_applicationsignalsMetricSourceTypes, "metric-source-types", "", "", "Metric Source Types")
	_applicationsignalsCmd.Flags().StringVarP(&_applicationsignalsName, "name", "", "", "Name")
	_applicationsignalsCmd.Flags().StringVarP(&_applicationsignalsNextToken, "next-token", "", "", "Next Token")
	_applicationsignalsCmd.Flags().StringVarP(&_applicationsignalsOperationName, "operation-name", "", "", "Operation Name")
	_applicationsignalsCmd.Flags().StringVarP(&_applicationsignalsRemoveExclusionWindows, "remove-exclusion-windows", "", "", "Remove Exclusion Windows")
	_applicationsignalsCmd.Flags().StringVarP(&_applicationsignalsRequestBasedSliConfig, "request-based-sli-config", "", "", "Request Based Sli Config")
	_applicationsignalsCmd.Flags().StringVarP(&_applicationsignalsResourceArn, "resource-arn", "", "", "Resource ARN")
	_applicationsignalsCmd.Flags().StringVarP(&_applicationsignalsSliConfig, "sli-config", "", "", "Sli Config")
	_applicationsignalsCmd.Flags().StringSliceVarP(&_applicationsignalsSloIds, "slo-ids", "", nil, "Slo Ids")
	_applicationsignalsCmd.Flags().StringVarP(&_applicationsignalsSloOwnerAwsAccountId, "slo-owner-aws-account-id", "", "", "Slo Owner AWS Account ID")
	_applicationsignalsCmd.Flags().StringVarP(&_applicationsignalsStartTime, "start-time", "", "", "Start Time")
	_applicationsignalsCmd.Flags().StringSliceVarP(&_applicationsignalsTagKeys, "tag-keys", "", nil, "Tag Keys")
	_applicationsignalsCmd.Flags().StringVarP(&_applicationsignalsTags, "tags", "", "", "Tags")
	_applicationsignalsCmd.Flags().StringVarP(&_applicationsignalsTimestamp, "timestamp", "", "", "Timestamp")

	_applicationsignalsCmd.Flags().BoolVarP(&_applicationsignalsBatchGetServiceLevelObjectiveBudgetReport, "batch-get-service-level-objective-budget-report", "", false, "Batch Get Service Level Objective Budget Report")
	_applicationsignalsCmd.Flags().BoolVarP(&_applicationsignalsBatchUpdateExclusionWindows, "batch-update-exclusion-windows", "", false, "Batch Update Exclusion Windows")
	_applicationsignalsCmd.Flags().BoolVarP(&_applicationsignalsCreateServiceLevelObjective, "create-service-level-objective", "", false, "Create Service Level Objective")
	_applicationsignalsCmd.Flags().BoolVarP(&_applicationsignalsDeleteGroupingConfiguration, "delete-grouping-configuration", "", false, "Delete Grouping Configuration")
	_applicationsignalsCmd.Flags().BoolVarP(&_applicationsignalsDeleteServiceLevelObjective, "delete-service-level-objective", "", false, "Delete Service Level Objective")
	_applicationsignalsCmd.Flags().BoolVarP(&_applicationsignalsGetService, "get-service", "", false, "Get Service")
	_applicationsignalsCmd.Flags().BoolVarP(&_applicationsignalsGetServiceLevelObjective, "get-service-level-objective", "", false, "Get Service Level Objective")
	_applicationsignalsCmd.Flags().BoolVarP(&_applicationsignalsListAuditFindings, "list-audit-findings", "", false, "List Audit Findings")
	_applicationsignalsCmd.Flags().BoolVarP(&_applicationsignalsListEntityEvents, "list-entity-events", "", false, "List Entity Events")
	_applicationsignalsCmd.Flags().BoolVarP(&_applicationsignalsListGroupingAttributeDefinitions, "list-grouping-attribute-definitions", "", false, "List Grouping Attribute Definitions")
	_applicationsignalsCmd.Flags().BoolVarP(&_applicationsignalsListServiceDependencies, "list-service-dependencies", "", false, "List Service Dependencies")
	_applicationsignalsCmd.Flags().BoolVarP(&_applicationsignalsListServiceDependents, "list-service-dependents", "", false, "List Service Dependents")
	_applicationsignalsCmd.Flags().BoolVarP(&_applicationsignalsListServiceLevelObjectiveExclusionWindows, "list-service-level-objective-exclusion-windows", "", false, "List Service Level Objective Exclusion Windows")
	_applicationsignalsCmd.Flags().BoolVarP(&_applicationsignalsListServiceLevelObjectives, "list-service-level-objectives", "", false, "List Service Level Objectives")
	_applicationsignalsCmd.Flags().BoolVarP(&_applicationsignalsListServiceOperations, "list-service-operations", "", false, "List Service Operations")
	_applicationsignalsCmd.Flags().BoolVarP(&_applicationsignalsListServiceStates, "list-service-states", "", false, "List Service States")
	_applicationsignalsCmd.Flags().BoolVarP(&_applicationsignalsListServices, "list-services", "", false, "List Services")
	_applicationsignalsCmd.Flags().BoolVarP(&_applicationsignalsListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_applicationsignalsCmd.Flags().BoolVarP(&_applicationsignalsPutGroupingConfiguration, "put-grouping-configuration", "", false, "Put Grouping Configuration")
	_applicationsignalsCmd.Flags().BoolVarP(&_applicationsignalsStartDiscovery, "start-discovery", "", false, "Start Discovery")
	_applicationsignalsCmd.Flags().BoolVarP(&_applicationsignalsTagResource, "tag-resource", "", false, "Tag Resource")
	_applicationsignalsCmd.Flags().BoolVarP(&_applicationsignalsUntagResource, "untag-resource", "", false, "Untag Resource")
	_applicationsignalsCmd.Flags().BoolVarP(&_applicationsignalsUpdateServiceLevelObjective, "update-service-level-objective", "", false, "Update Service Level Objective")

}
