package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/amp"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ampCmd represents the amp command
var _ampCmd = &cobra.Command{
	Use:   "amp",
	Short: "AWS amp CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := amp.NewFromConfig(cfg)
		if _ampCreateAlertManagerDefinition {
			amp_CreateAlertManagerDefinition(cfg, client)
			return
		}
		if _ampCreateAnomalyDetector {
			amp_CreateAnomalyDetector(cfg, client)
			return
		}
		if _ampCreateLoggingConfiguration {
			amp_CreateLoggingConfiguration(cfg, client)
			return
		}
		if _ampCreateQueryLoggingConfiguration {
			amp_CreateQueryLoggingConfiguration(cfg, client)
			return
		}
		if _ampCreateRuleGroupsNamespace {
			amp_CreateRuleGroupsNamespace(cfg, client)
			return
		}
		if _ampCreateScraper {
			amp_CreateScraper(cfg, client)
			return
		}
		if _ampCreateWorkspace {
			amp_CreateWorkspace(cfg, client)
			return
		}
		if _ampDeleteAlertManagerDefinition {
			amp_DeleteAlertManagerDefinition(cfg, client)
			return
		}
		if _ampDeleteAnomalyDetector {
			amp_DeleteAnomalyDetector(cfg, client)
			return
		}
		if _ampDeleteLoggingConfiguration {
			amp_DeleteLoggingConfiguration(cfg, client)
			return
		}
		if _ampDeleteQueryLoggingConfiguration {
			amp_DeleteQueryLoggingConfiguration(cfg, client)
			return
		}
		if _ampDeleteResourcePolicy {
			amp_DeleteResourcePolicy(cfg, client)
			return
		}
		if _ampDeleteRuleGroupsNamespace {
			amp_DeleteRuleGroupsNamespace(cfg, client)
			return
		}
		if _ampDeleteScraper {
			amp_DeleteScraper(cfg, client)
			return
		}
		if _ampDeleteScraperLoggingConfiguration {
			amp_DeleteScraperLoggingConfiguration(cfg, client)
			return
		}
		if _ampDeleteWorkspace {
			amp_DeleteWorkspace(cfg, client)
			return
		}
		if _ampDescribeAlertManagerDefinition {
			amp_DescribeAlertManagerDefinition(cfg, client)
			return
		}
		if _ampDescribeAnomalyDetector {
			amp_DescribeAnomalyDetector(cfg, client)
			return
		}
		if _ampDescribeLoggingConfiguration {
			amp_DescribeLoggingConfiguration(cfg, client)
			return
		}
		if _ampDescribeQueryLoggingConfiguration {
			amp_DescribeQueryLoggingConfiguration(cfg, client)
			return
		}
		if _ampDescribeResourcePolicy {
			amp_DescribeResourcePolicy(cfg, client)
			return
		}
		if _ampDescribeRuleGroupsNamespace {
			amp_DescribeRuleGroupsNamespace(cfg, client)
			return
		}
		if _ampDescribeScraper {
			amp_DescribeScraper(cfg, client)
			return
		}
		if _ampDescribeScraperLoggingConfiguration {
			amp_DescribeScraperLoggingConfiguration(cfg, client)
			return
		}
		if _ampDescribeWorkspace {
			amp_DescribeWorkspace(cfg, client)
			return
		}
		if _ampDescribeWorkspaceConfiguration {
			amp_DescribeWorkspaceConfiguration(cfg, client)
			return
		}
		if _ampGetDefaultScraperConfiguration {
			amp_GetDefaultScraperConfiguration(cfg, client)
			return
		}
		if _ampListAnomalyDetectors {
			amp_ListAnomalyDetectors(cfg, client)
			return
		}
		if _ampListRuleGroupsNamespaces {
			amp_ListRuleGroupsNamespaces(cfg, client)
			return
		}
		if _ampListScrapers {
			amp_ListScrapers(cfg, client)
			return
		}
		if _ampListTagsForResource {
			amp_ListTagsForResource(cfg, client)
			return
		}
		if _ampListWorkspaces {
			amp_ListWorkspaces(cfg, client)
			return
		}
		if _ampPutAlertManagerDefinition {
			amp_PutAlertManagerDefinition(cfg, client)
			return
		}
		if _ampPutAnomalyDetector {
			amp_PutAnomalyDetector(cfg, client)
			return
		}
		if _ampPutResourcePolicy {
			amp_PutResourcePolicy(cfg, client)
			return
		}
		if _ampPutRuleGroupsNamespace {
			amp_PutRuleGroupsNamespace(cfg, client)
			return
		}
		if _ampTagResource {
			amp_TagResource(cfg, client)
			return
		}
		if _ampUntagResource {
			amp_UntagResource(cfg, client)
			return
		}
		if _ampUpdateLoggingConfiguration {
			amp_UpdateLoggingConfiguration(cfg, client)
			return
		}
		if _ampUpdateQueryLoggingConfiguration {
			amp_UpdateQueryLoggingConfiguration(cfg, client)
			return
		}
		if _ampUpdateScraper {
			amp_UpdateScraper(cfg, client)
			return
		}
		if _ampUpdateScraperLoggingConfiguration {
			amp_UpdateScraperLoggingConfiguration(cfg, client)
			return
		}
		if _ampUpdateWorkspaceAlias {
			amp_UpdateWorkspaceAlias(cfg, client)
			return
		}
		if _ampUpdateWorkspaceConfiguration {
			amp_UpdateWorkspaceConfiguration(cfg, client)
			return
		}

	},
}

var (
	_ampCreateAlertManagerDefinition        bool
	_ampCreateAnomalyDetector               bool
	_ampCreateLoggingConfiguration          bool
	_ampCreateQueryLoggingConfiguration     bool
	_ampCreateRuleGroupsNamespace           bool
	_ampCreateScraper                       bool
	_ampCreateWorkspace                     bool
	_ampDeleteAlertManagerDefinition        bool
	_ampDeleteAnomalyDetector               bool
	_ampDeleteLoggingConfiguration          bool
	_ampDeleteQueryLoggingConfiguration     bool
	_ampDeleteResourcePolicy                bool
	_ampDeleteRuleGroupsNamespace           bool
	_ampDeleteScraper                       bool
	_ampDeleteScraperLoggingConfiguration   bool
	_ampDeleteWorkspace                     bool
	_ampDescribeAlertManagerDefinition      bool
	_ampDescribeAnomalyDetector             bool
	_ampDescribeLoggingConfiguration        bool
	_ampDescribeQueryLoggingConfiguration   bool
	_ampDescribeResourcePolicy              bool
	_ampDescribeRuleGroupsNamespace         bool
	_ampDescribeScraper                     bool
	_ampDescribeScraperLoggingConfiguration bool
	_ampDescribeWorkspace                   bool
	_ampDescribeWorkspaceConfiguration      bool
	_ampGetDefaultScraperConfiguration      bool
	_ampListAnomalyDetectors                bool
	_ampListRuleGroupsNamespaces            bool
	_ampListScrapers                        bool
	_ampListTagsForResource                 bool
	_ampListWorkspaces                      bool
	_ampPutAlertManagerDefinition           bool
	_ampPutAnomalyDetector                  bool
	_ampPutResourcePolicy                   bool
	_ampPutRuleGroupsNamespace              bool
	_ampTagResource                         bool
	_ampUntagResource                       bool
	_ampUpdateLoggingConfiguration          bool
	_ampUpdateQueryLoggingConfiguration     bool
	_ampUpdateScraper                       bool
	_ampUpdateScraperLoggingConfiguration   bool
	_ampUpdateWorkspaceAlias                bool
	_ampUpdateWorkspaceConfiguration        bool

	_ampAlias                       string
	_ampAnomalyDetectorId           string
	_ampClientToken                 string
	_ampConfiguration               string
	_ampData                        string
	_ampDestination                 string
	_ampDestinations                string
	_ampEvaluationIntervalInSeconds string
	_ampFilters                     string
	_ampKmsKeyArn                   string
	_ampLabels                      string
	_ampLimitsPerLabelSet           string
	_ampLogGroupArn                 string
	_ampLoggingDestination          string
	_ampMaxResults                  string
	_ampMissingDataAction           string
	_ampName                        string
	_ampNextToken                   string
	_ampPolicyDocument              string
	_ampResourceArn                 string
	_ampRetentionPeriodInDays       string
	_ampRevisionId                  string
	_ampRoleConfiguration           string
	_ampScrapeConfiguration         string
	_ampScraperComponents           string
	_ampScraperId                   string
	_ampSource                      string
	_ampTagKeys                     []string
	_ampTags                        string
	_ampWorkspaceId                 string
)

// The CreateAlertManagerDefinition operation creates the alert manager definition
// in a workspace. If a workspace already has an alert manager definition, don't
// use this operation to update it. Instead, use PutAlertManagerDefinition .
func amp_CreateAlertManagerDefinition(cfg aws.Config, client *amp.Client) {
	input := &amp.CreateAlertManagerDefinitionInput{
		// Data: []byte, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_ampData) > 0 {
		if err := assignInputField(input, "Data", _ampData); err != nil {
			log.Errorf("invalid --data: %s", err.Error())
			return
		}
	}
	if len(_ampWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_ampWorkspaceId)
	}
	if len(_ampClientToken) > 0 {
		input.ClientToken = aws.String(_ampClientToken)
	}

	if resp, err := client.CreateAlertManagerDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an anomaly detector within a workspace using the Random Cut Forest
// algorithm for time-series analysis. The anomaly detector analyzes Amazon Managed
// Service for Prometheus metrics to identify unusual patterns and behaviors.
func amp_CreateAnomalyDetector(cfg aws.Config, client *amp.Client) {
	input := &amp.CreateAnomalyDetectorInput{
		// Alias: *string, // Required
		// Configuration: types.AnomalyDetectorConfiguration, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_ampAlias) > 0 {
		input.Alias = aws.String(_ampAlias)
	}
	if len(_ampConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _ampConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_ampWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_ampWorkspaceId)
	}
	if len(_ampClientToken) > 0 {
		input.ClientToken = aws.String(_ampClientToken)
	}
	if len(_ampEvaluationIntervalInSeconds) > 0 {
		if err := assignInputField(input, "EvaluationIntervalInSeconds", _ampEvaluationIntervalInSeconds); err != nil {
			log.Errorf("invalid --evaluation-interval-in-seconds: %s", err.Error())
			return
		}
	}
	if len(_ampLabels) > 0 {
		if err := assignInputField(input, "Labels", _ampLabels); err != nil {
			log.Errorf("invalid --labels: %s", err.Error())
			return
		}
	}
	if len(_ampMissingDataAction) > 0 {
		if err := assignInputField(input, "MissingDataAction", _ampMissingDataAction); err != nil {
			log.Errorf("invalid --missing-data-action: %s", err.Error())
			return
		}
	}
	if len(_ampTags) > 0 {
		if err := assignInputField(input, "Tags", _ampTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAnomalyDetector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The CreateLoggingConfiguration operation creates rules and alerting logging
// configuration for the workspace. Use this operation to set the CloudWatch log
// group to which the logs will be published to.
//
// These logging configurations are only for rules and alerting logs.
func amp_CreateLoggingConfiguration(cfg aws.Config, client *amp.Client) {
	input := &amp.CreateLoggingConfigurationInput{
		// LogGroupArn: *string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_ampLogGroupArn) > 0 {
		input.LogGroupArn = aws.String(_ampLogGroupArn)
	}
	if len(_ampWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_ampWorkspaceId)
	}
	if len(_ampClientToken) > 0 {
		input.ClientToken = aws.String(_ampClientToken)
	}

	if resp, err := client.CreateLoggingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a query logging configuration for the specified workspace. This
// operation enables logging of queries that exceed the specified QSP threshold.
func amp_CreateQueryLoggingConfiguration(cfg aws.Config, client *amp.Client) {
	input := &amp.CreateQueryLoggingConfigurationInput{
		// Destinations: []types.LoggingDestination, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_ampDestinations) > 0 {
		if err := assignInputField(input, "Destinations", _ampDestinations); err != nil {
			log.Errorf("invalid --destinations: %s", err.Error())
			return
		}
	}
	if len(_ampWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_ampWorkspaceId)
	}
	if len(_ampClientToken) > 0 {
		input.ClientToken = aws.String(_ampClientToken)
	}

	if resp, err := client.CreateQueryLoggingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The CreateRuleGroupsNamespace operation creates a rule groups namespace within
// a workspace. A rule groups namespace is associated with exactly one rules file.
// A workspace can have multiple rule groups namespaces.
//
// The combined length of a rule group namespace and a rule group name cannot
// exceed 721 UTF-8 bytes.
//
// Use this operation only to create new rule groups namespaces. To update an
// existing rule groups namespace, use PutRuleGroupsNamespace .
func amp_CreateRuleGroupsNamespace(cfg aws.Config, client *amp.Client) {
	input := &amp.CreateRuleGroupsNamespaceInput{
		// Data: []byte, // Required
		// Name: *string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_ampData) > 0 {
		if err := assignInputField(input, "Data", _ampData); err != nil {
			log.Errorf("invalid --data: %s", err.Error())
			return
		}
	}
	if len(_ampName) > 0 {
		input.Name = aws.String(_ampName)
	}
	if len(_ampWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_ampWorkspaceId)
	}
	if len(_ampClientToken) > 0 {
		input.ClientToken = aws.String(_ampClientToken)
	}
	if len(_ampTags) > 0 {
		if err := assignInputField(input, "Tags", _ampTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRuleGroupsNamespace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The CreateScraper operation creates a scraper to collect metrics. A scraper
// pulls metrics from Prometheus-compatible sources and sends them to your Amazon
// Managed Service for Prometheus workspace. You can configure scrapers to collect
// metrics from Amazon EKS clusters, Amazon MSK clusters, or from VPC-based sources
// that support DNS-based service discovery. Scrapers are flexible, and can be
// configured to control what metrics are collected, the frequency of collection,
// what transformations are applied to the metrics, and more.
//
// An IAM role will be created for you that Amazon Managed Service for Prometheus
// uses to access the metrics in your source. You must configure this role with a
// policy that allows it to scrape metrics from your source. For Amazon EKS
// sources, see [Configuring your Amazon EKS cluster]in the Amazon Managed Service for Prometheus User Guide.
//
// The scrapeConfiguration parameter contains the base-64 encoded YAML
// configuration for the scraper.
//
// When creating a scraper, the service creates a Network Interface in each
// Availability Zone that are passed into CreateScraper through subnets. These
// network interfaces are used to connect to your source within the VPC for
// scraping metrics.
//
// For more information about collectors, including what metrics are collected,
// and how to configure the scraper, see [Using an Amazon Web Services managed collector]in the Amazon Managed Service for
// Prometheus User Guide.
//
// [Using an Amazon Web Services managed collector]: https://docs.aws.amazon.com/prometheus/latest/userguide/AMP-collector-how-to.html
// [Configuring your Amazon EKS cluster]: https://docs.aws.amazon.com/prometheus/latest/userguide/AMP-collector-how-to.html#AMP-collector-eks-setup
func amp_CreateScraper(cfg aws.Config, client *amp.Client) {
	input := &amp.CreateScraperInput{
		// Destination: types.Destination, // Required
		// ScrapeConfiguration: types.ScrapeConfiguration, // Required
		// Source: types.Source, // Required
	}

	if len(_ampDestination) > 0 {
		if err := assignInputField(input, "Destination", _ampDestination); err != nil {
			log.Errorf("invalid --destination: %s", err.Error())
			return
		}
	}
	if len(_ampScrapeConfiguration) > 0 {
		if err := assignInputField(input, "ScrapeConfiguration", _ampScrapeConfiguration); err != nil {
			log.Errorf("invalid --scrape-configuration: %s", err.Error())
			return
		}
	}
	if len(_ampSource) > 0 {
		if err := assignInputField(input, "Source", _ampSource); err != nil {
			log.Errorf("invalid --source: %s", err.Error())
			return
		}
	}
	if len(_ampAlias) > 0 {
		input.Alias = aws.String(_ampAlias)
	}
	if len(_ampClientToken) > 0 {
		input.ClientToken = aws.String(_ampClientToken)
	}
	if len(_ampRoleConfiguration) > 0 {
		if err := assignInputField(input, "RoleConfiguration", _ampRoleConfiguration); err != nil {
			log.Errorf("invalid --role-configuration: %s", err.Error())
			return
		}
	}
	if len(_ampTags) > 0 {
		if err := assignInputField(input, "Tags", _ampTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateScraper(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a Prometheus workspace. A workspace is a logical space dedicated to the
// storage and querying of Prometheus metrics. You can have one or more workspaces
// in each Region in your account.
func amp_CreateWorkspace(cfg aws.Config, client *amp.Client) {
	input := &amp.CreateWorkspaceInput{}

	if len(_ampAlias) > 0 {
		input.Alias = aws.String(_ampAlias)
	}
	if len(_ampClientToken) > 0 {
		input.ClientToken = aws.String(_ampClientToken)
	}
	if len(_ampKmsKeyArn) > 0 {
		input.KmsKeyArn = aws.String(_ampKmsKeyArn)
	}
	if len(_ampTags) > 0 {
		if err := assignInputField(input, "Tags", _ampTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateWorkspace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the alert manager definition from a workspace.
func amp_DeleteAlertManagerDefinition(cfg aws.Config, client *amp.Client) {
	input := &amp.DeleteAlertManagerDefinitionInput{
		// WorkspaceId: *string, // Required
	}

	if len(_ampWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_ampWorkspaceId)
	}
	if len(_ampClientToken) > 0 {
		input.ClientToken = aws.String(_ampClientToken)
	}

	if resp, err := client.DeleteAlertManagerDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes an anomaly detector from a workspace. This operation is idempotent.
func amp_DeleteAnomalyDetector(cfg aws.Config, client *amp.Client) {
	input := &amp.DeleteAnomalyDetectorInput{
		// AnomalyDetectorId: *string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_ampAnomalyDetectorId) > 0 {
		input.AnomalyDetectorId = aws.String(_ampAnomalyDetectorId)
	}
	if len(_ampWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_ampWorkspaceId)
	}
	if len(_ampClientToken) > 0 {
		input.ClientToken = aws.String(_ampClientToken)
	}

	if resp, err := client.DeleteAnomalyDetector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the rules and alerting logging configuration for a workspace.
// These logging configurations are only for rules and alerting logs.
func amp_DeleteLoggingConfiguration(cfg aws.Config, client *amp.Client) {
	input := &amp.DeleteLoggingConfigurationInput{
		// WorkspaceId: *string, // Required
	}

	if len(_ampWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_ampWorkspaceId)
	}
	if len(_ampClientToken) > 0 {
		input.ClientToken = aws.String(_ampClientToken)
	}

	if resp, err := client.DeleteLoggingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the query logging configuration for the specified workspace.
func amp_DeleteQueryLoggingConfiguration(cfg aws.Config, client *amp.Client) {
	input := &amp.DeleteQueryLoggingConfigurationInput{
		// WorkspaceId: *string, // Required
	}

	if len(_ampWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_ampWorkspaceId)
	}
	if len(_ampClientToken) > 0 {
		input.ClientToken = aws.String(_ampClientToken)
	}

	if resp, err := client.DeleteQueryLoggingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the resource-based policy attached to an Amazon Managed Service for
// Prometheus workspace.
func amp_DeleteResourcePolicy(cfg aws.Config, client *amp.Client) {
	input := &amp.DeleteResourcePolicyInput{
		// WorkspaceId: *string, // Required
	}

	if len(_ampWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_ampWorkspaceId)
	}
	if len(_ampClientToken) > 0 {
		input.ClientToken = aws.String(_ampClientToken)
	}
	if len(_ampRevisionId) > 0 {
		input.RevisionId = aws.String(_ampRevisionId)
	}

	if resp, err := client.DeleteResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes one rule groups namespace and its associated rule groups definition.
func amp_DeleteRuleGroupsNamespace(cfg aws.Config, client *amp.Client) {
	input := &amp.DeleteRuleGroupsNamespaceInput{
		// Name: *string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_ampName) > 0 {
		input.Name = aws.String(_ampName)
	}
	if len(_ampWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_ampWorkspaceId)
	}
	if len(_ampClientToken) > 0 {
		input.ClientToken = aws.String(_ampClientToken)
	}

	if resp, err := client.DeleteRuleGroupsNamespace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The DeleteScraper operation deletes one scraper, and stops any metrics
// collection that the scraper performs.
func amp_DeleteScraper(cfg aws.Config, client *amp.Client) {
	input := &amp.DeleteScraperInput{
		// ScraperId: *string, // Required
	}

	if len(_ampScraperId) > 0 {
		input.ScraperId = aws.String(_ampScraperId)
	}
	if len(_ampClientToken) > 0 {
		input.ClientToken = aws.String(_ampClientToken)
	}

	if resp, err := client.DeleteScraper(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the logging configuration for a Amazon Managed Service for Prometheus
// scraper.
func amp_DeleteScraperLoggingConfiguration(cfg aws.Config, client *amp.Client) {
	input := &amp.DeleteScraperLoggingConfigurationInput{
		// ScraperId: *string, // Required
	}

	if len(_ampScraperId) > 0 {
		input.ScraperId = aws.String(_ampScraperId)
	}
	if len(_ampClientToken) > 0 {
		input.ClientToken = aws.String(_ampClientToken)
	}

	if resp, err := client.DeleteScraperLoggingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing workspace.
// When you delete a workspace, the data that has been ingested into it is not
// immediately deleted. It will be permanently deleted within one month.
func amp_DeleteWorkspace(cfg aws.Config, client *amp.Client) {
	input := &amp.DeleteWorkspaceInput{
		// WorkspaceId: *string, // Required
	}

	if len(_ampWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_ampWorkspaceId)
	}
	if len(_ampClientToken) > 0 {
		input.ClientToken = aws.String(_ampClientToken)
	}

	if resp, err := client.DeleteWorkspace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the full information about the alert manager definition for a
// workspace.
func amp_DescribeAlertManagerDefinition(cfg aws.Config, client *amp.Client) {
	input := &amp.DescribeAlertManagerDefinitionInput{
		// WorkspaceId: *string, // Required
	}

	if len(_ampWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_ampWorkspaceId)
	}

	if resp, err := client.DescribeAlertManagerDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves detailed information about a specific anomaly detector, including its
// status and configuration.
func amp_DescribeAnomalyDetector(cfg aws.Config, client *amp.Client) {
	input := &amp.DescribeAnomalyDetectorInput{
		// AnomalyDetectorId: *string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_ampAnomalyDetectorId) > 0 {
		input.AnomalyDetectorId = aws.String(_ampAnomalyDetectorId)
	}
	if len(_ampWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_ampWorkspaceId)
	}

	if resp, err := client.DescribeAnomalyDetector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns complete information about the current rules and alerting logging
// configuration of the workspace.
//
// These logging configurations are only for rules and alerting logs.
func amp_DescribeLoggingConfiguration(cfg aws.Config, client *amp.Client) {
	input := &amp.DescribeLoggingConfigurationInput{
		// WorkspaceId: *string, // Required
	}

	if len(_ampWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_ampWorkspaceId)
	}

	if resp, err := client.DescribeLoggingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the details of the query logging configuration for the specified
// workspace.
func amp_DescribeQueryLoggingConfiguration(cfg aws.Config, client *amp.Client) {
	input := &amp.DescribeQueryLoggingConfigurationInput{
		// WorkspaceId: *string, // Required
	}

	if len(_ampWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_ampWorkspaceId)
	}

	if resp, err := client.DescribeQueryLoggingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the resource-based policy attached to an Amazon
// Managed Service for Prometheus workspace.
func amp_DescribeResourcePolicy(cfg aws.Config, client *amp.Client) {
	input := &amp.DescribeResourcePolicyInput{
		// WorkspaceId: *string, // Required
	}

	if len(_ampWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_ampWorkspaceId)
	}

	if resp, err := client.DescribeResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns complete information about one rule groups namespace. To retrieve a
// list of rule groups namespaces, use ListRuleGroupsNamespaces .
func amp_DescribeRuleGroupsNamespace(cfg aws.Config, client *amp.Client) {
	input := &amp.DescribeRuleGroupsNamespaceInput{
		// Name: *string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_ampName) > 0 {
		input.Name = aws.String(_ampName)
	}
	if len(_ampWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_ampWorkspaceId)
	}

	if resp, err := client.DescribeRuleGroupsNamespace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The DescribeScraper operation displays information about an existing scraper.
func amp_DescribeScraper(cfg aws.Config, client *amp.Client) {
	input := &amp.DescribeScraperInput{
		// ScraperId: *string, // Required
	}

	if len(_ampScraperId) > 0 {
		input.ScraperId = aws.String(_ampScraperId)
	}

	if resp, err := client.DescribeScraper(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the logging configuration for a Amazon Managed Service for Prometheus
// scraper.
func amp_DescribeScraperLoggingConfiguration(cfg aws.Config, client *amp.Client) {
	input := &amp.DescribeScraperLoggingConfigurationInput{
		// ScraperId: *string, // Required
	}

	if len(_ampScraperId) > 0 {
		input.ScraperId = aws.String(_ampScraperId)
	}

	if resp, err := client.DescribeScraperLoggingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about an existing workspace.
func amp_DescribeWorkspace(cfg aws.Config, client *amp.Client) {
	input := &amp.DescribeWorkspaceInput{
		// WorkspaceId: *string, // Required
	}

	if len(_ampWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_ampWorkspaceId)
	}

	if resp, err := client.DescribeWorkspace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use this operation to return information about the configuration of a
// workspace. The configuration details returned include workspace configuration
// status, label set limits, and retention period.
func amp_DescribeWorkspaceConfiguration(cfg aws.Config, client *amp.Client) {
	input := &amp.DescribeWorkspaceConfigurationInput{
		// WorkspaceId: *string, // Required
	}

	if len(_ampWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_ampWorkspaceId)
	}

	if resp, err := client.DescribeWorkspaceConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The GetDefaultScraperConfiguration operation returns the default scraper
// configuration used when Amazon EKS creates a scraper for you.
func amp_GetDefaultScraperConfiguration(cfg aws.Config, client *amp.Client) {
	input := &amp.GetDefaultScraperConfigurationInput{}

	if resp, err := client.GetDefaultScraperConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a paginated list of anomaly detectors for a workspace with optional
// filtering by alias.
func amp_ListAnomalyDetectors(cfg aws.Config, client *amp.Client) {
	input := &amp.ListAnomalyDetectorsInput{
		// WorkspaceId: *string, // Required
	}

	if len(_ampWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_ampWorkspaceId)
	}
	if len(_ampAlias) > 0 {
		input.Alias = aws.String(_ampAlias)
	}
	if len(_ampMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ampMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ampNextToken) > 0 {
		input.NextToken = aws.String(_ampNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAnomalyDetectors(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*amp.ListAnomalyDetectorsOutput
	p := amp.NewListAnomalyDetectorsPaginator(client, input)
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

// Returns a list of rule groups namespaces in a workspace.
func amp_ListRuleGroupsNamespaces(cfg aws.Config, client *amp.Client) {
	input := &amp.ListRuleGroupsNamespacesInput{
		// WorkspaceId: *string, // Required
	}

	if len(_ampWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_ampWorkspaceId)
	}
	if len(_ampMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ampMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ampName) > 0 {
		input.Name = aws.String(_ampName)
	}
	if len(_ampNextToken) > 0 {
		input.NextToken = aws.String(_ampNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRuleGroupsNamespaces(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*amp.ListRuleGroupsNamespacesOutput
	p := amp.NewListRuleGroupsNamespacesPaginator(client, input)
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

// The ListScrapers operation lists all of the scrapers in your account. This
// includes scrapers being created or deleted. You can optionally filter the
// returned list.
func amp_ListScrapers(cfg aws.Config, client *amp.Client) {
	input := &amp.ListScrapersInput{}

	if len(_ampFilters) > 0 {
		if err := assignInputField(input, "Filters", _ampFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_ampMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ampMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ampNextToken) > 0 {
		input.NextToken = aws.String(_ampNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListScrapers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*amp.ListScrapersOutput
	p := amp.NewListScrapersPaginator(client, input)
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

// The ListTagsForResource operation returns the tags that are associated with an
// Amazon Managed Service for Prometheus resource. Currently, the only resources
// that can be tagged are scrapers, workspaces, and rule groups namespaces.
func amp_ListTagsForResource(cfg aws.Config, client *amp.Client) {
	input := &amp.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_ampResourceArn) > 0 {
		input.ResourceArn = aws.String(_ampResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all of the Amazon Managed Service for Prometheus workspaces in your
// account. This includes workspaces being created or deleted.
func amp_ListWorkspaces(cfg aws.Config, client *amp.Client) {
	input := &amp.ListWorkspacesInput{}

	if len(_ampAlias) > 0 {
		input.Alias = aws.String(_ampAlias)
	}
	if len(_ampMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ampMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ampNextToken) > 0 {
		input.NextToken = aws.String(_ampNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListWorkspaces(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*amp.ListWorkspacesOutput
	p := amp.NewListWorkspacesPaginator(client, input)
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

// Updates an existing alert manager definition in a workspace. If the workspace
// does not already have an alert manager definition, don't use this operation to
// create it. Instead, use CreateAlertManagerDefinition .
func amp_PutAlertManagerDefinition(cfg aws.Config, client *amp.Client) {
	input := &amp.PutAlertManagerDefinitionInput{
		// Data: []byte, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_ampData) > 0 {
		if err := assignInputField(input, "Data", _ampData); err != nil {
			log.Errorf("invalid --data: %s", err.Error())
			return
		}
	}
	if len(_ampWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_ampWorkspaceId)
	}
	if len(_ampClientToken) > 0 {
		input.ClientToken = aws.String(_ampClientToken)
	}

	if resp, err := client.PutAlertManagerDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// When you call PutAnomalyDetector , the operation creates a new anomaly detector
// if one doesn't exist, or updates an existing one. Each call to this operation
// triggers a complete retraining of the detector, which includes querying the
// minimum required samples and backfilling the detector with historical data. This
// process occurs regardless of whether you're making a minor change like updating
// the evaluation interval or making more substantial modifications. The operation
// serves as the single method for creating, updating, and retraining anomaly
// detectors.
func amp_PutAnomalyDetector(cfg aws.Config, client *amp.Client) {
	input := &amp.PutAnomalyDetectorInput{
		// AnomalyDetectorId: *string, // Required
		// Configuration: types.AnomalyDetectorConfiguration, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_ampAnomalyDetectorId) > 0 {
		input.AnomalyDetectorId = aws.String(_ampAnomalyDetectorId)
	}
	if len(_ampConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _ampConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_ampWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_ampWorkspaceId)
	}
	if len(_ampClientToken) > 0 {
		input.ClientToken = aws.String(_ampClientToken)
	}
	if len(_ampEvaluationIntervalInSeconds) > 0 {
		if err := assignInputField(input, "EvaluationIntervalInSeconds", _ampEvaluationIntervalInSeconds); err != nil {
			log.Errorf("invalid --evaluation-interval-in-seconds: %s", err.Error())
			return
		}
	}
	if len(_ampLabels) > 0 {
		if err := assignInputField(input, "Labels", _ampLabels); err != nil {
			log.Errorf("invalid --labels: %s", err.Error())
			return
		}
	}
	if len(_ampMissingDataAction) > 0 {
		if err := assignInputField(input, "MissingDataAction", _ampMissingDataAction); err != nil {
			log.Errorf("invalid --missing-data-action: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutAnomalyDetector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates a resource-based policy for an Amazon Managed Service for
// Prometheus workspace. Use resource-based policies to grant permissions to other
// AWS accounts or services to access your workspace.
//
// Only Prometheus-compatible APIs can be used for workspace sharing. You can add
// non-Prometheus-compatible APIs to the policy, but they will be ignored. For more
// information, see [Prometheus-compatible APIs]in the Amazon Managed Service for Prometheus User Guide.
//
// If your workspace uses customer-managed KMS keys for encryption, you must grant
// the principals in your resource-based policy access to those KMS keys. You can
// do this by creating KMS grants. For more information, see [CreateGrant]in the AWS Key
// Management Service API Reference and [Encryption at rest]in the Amazon Managed Service for
// Prometheus User Guide.
//
// For more information about working with IAM, see [Using Amazon Managed Service for Prometheus with IAM] in the Amazon Managed Service
// for Prometheus User Guide.
//
// [Prometheus-compatible APIs]: https://docs.aws.amazon.com/prometheus/latest/userguide/AMP-APIReference-Prometheus-Compatible-Apis.html
// [Using Amazon Managed Service for Prometheus with IAM]: https://docs.aws.amazon.com/prometheus/latest/userguide/security_iam_service-with-iam.html
// [Encryption at rest]: https://docs.aws.amazon.com/prometheus/latest/userguide/encryption-at-rest-Amazon-Service-Prometheus.html
// [CreateGrant]: https://docs.aws.amazon.com/kms/latest/APIReference/API_CreateGrant.html
func amp_PutResourcePolicy(cfg aws.Config, client *amp.Client) {
	input := &amp.PutResourcePolicyInput{
		// PolicyDocument: *string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_ampPolicyDocument) > 0 {
		input.PolicyDocument = aws.String(_ampPolicyDocument)
	}
	if len(_ampWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_ampWorkspaceId)
	}
	if len(_ampClientToken) > 0 {
		input.ClientToken = aws.String(_ampClientToken)
	}
	if len(_ampRevisionId) > 0 {
		input.RevisionId = aws.String(_ampRevisionId)
	}

	if resp, err := client.PutResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing rule groups namespace within a workspace. A rule groups
// namespace is associated with exactly one rules file. A workspace can have
// multiple rule groups namespaces.
//
// The combined length of a rule group namespace and a rule group name cannot
// exceed 721 UTF-8 bytes.
//
// Use this operation only to update existing rule groups namespaces. To create a
// new rule groups namespace, use CreateRuleGroupsNamespace .
//
// You can't use this operation to add tags to an existing rule groups namespace.
// Instead, use TagResource .
func amp_PutRuleGroupsNamespace(cfg aws.Config, client *amp.Client) {
	input := &amp.PutRuleGroupsNamespaceInput{
		// Data: []byte, // Required
		// Name: *string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_ampData) > 0 {
		if err := assignInputField(input, "Data", _ampData); err != nil {
			log.Errorf("invalid --data: %s", err.Error())
			return
		}
	}
	if len(_ampName) > 0 {
		input.Name = aws.String(_ampName)
	}
	if len(_ampWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_ampWorkspaceId)
	}
	if len(_ampClientToken) > 0 {
		input.ClientToken = aws.String(_ampClientToken)
	}

	if resp, err := client.PutRuleGroupsNamespace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The TagResource operation associates tags with an Amazon Managed Service for
// Prometheus resource. The only resources that can be tagged are rule groups
// namespaces, scrapers, and workspaces.
//
// If you specify a new tag key for the resource, this tag is appended to the list
// of tags associated with the resource. If you specify a tag key that is already
// associated with the resource, the new tag value that you specify replaces the
// previous value for that tag. To remove a tag, use UntagResource .
func amp_TagResource(cfg aws.Config, client *amp.Client) {
	input := &amp.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_ampResourceArn) > 0 {
		input.ResourceArn = aws.String(_ampResourceArn)
	}
	if len(_ampTags) > 0 {
		if err := assignInputField(input, "Tags", _ampTags); err != nil {
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

// Removes the specified tags from an Amazon Managed Service for Prometheus
// resource. The only resources that can be tagged are rule groups namespaces,
// scrapers, and workspaces.
func amp_UntagResource(cfg aws.Config, client *amp.Client) {
	input := &amp.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_ampResourceArn) > 0 {
		input.ResourceArn = aws.String(_ampResourceArn)
	}
	if len(_ampTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _ampTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the log group ARN or the workspace ID of the current rules and alerting
// logging configuration.
//
// These logging configurations are only for rules and alerting logs.
func amp_UpdateLoggingConfiguration(cfg aws.Config, client *amp.Client) {
	input := &amp.UpdateLoggingConfigurationInput{
		// LogGroupArn: *string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_ampLogGroupArn) > 0 {
		input.LogGroupArn = aws.String(_ampLogGroupArn)
	}
	if len(_ampWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_ampWorkspaceId)
	}
	if len(_ampClientToken) > 0 {
		input.ClientToken = aws.String(_ampClientToken)
	}

	if resp, err := client.UpdateLoggingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the query logging configuration for the specified workspace.
func amp_UpdateQueryLoggingConfiguration(cfg aws.Config, client *amp.Client) {
	input := &amp.UpdateQueryLoggingConfigurationInput{
		// Destinations: []types.LoggingDestination, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_ampDestinations) > 0 {
		if err := assignInputField(input, "Destinations", _ampDestinations); err != nil {
			log.Errorf("invalid --destinations: %s", err.Error())
			return
		}
	}
	if len(_ampWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_ampWorkspaceId)
	}
	if len(_ampClientToken) > 0 {
		input.ClientToken = aws.String(_ampClientToken)
	}

	if resp, err := client.UpdateQueryLoggingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing scraper.
// You can't use this function to update the source from which the scraper is
// collecting metrics. To change the source, delete the scraper and create a new
// one.
func amp_UpdateScraper(cfg aws.Config, client *amp.Client) {
	input := &amp.UpdateScraperInput{
		// ScraperId: *string, // Required
	}

	if len(_ampScraperId) > 0 {
		input.ScraperId = aws.String(_ampScraperId)
	}
	if len(_ampAlias) > 0 {
		input.Alias = aws.String(_ampAlias)
	}
	if len(_ampClientToken) > 0 {
		input.ClientToken = aws.String(_ampClientToken)
	}
	if len(_ampDestination) > 0 {
		if err := assignInputField(input, "Destination", _ampDestination); err != nil {
			log.Errorf("invalid --destination: %s", err.Error())
			return
		}
	}
	if len(_ampRoleConfiguration) > 0 {
		if err := assignInputField(input, "RoleConfiguration", _ampRoleConfiguration); err != nil {
			log.Errorf("invalid --role-configuration: %s", err.Error())
			return
		}
	}
	if len(_ampScrapeConfiguration) > 0 {
		if err := assignInputField(input, "ScrapeConfiguration", _ampScrapeConfiguration); err != nil {
			log.Errorf("invalid --scrape-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateScraper(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the logging configuration for a Amazon Managed Service for Prometheus
// scraper.
func amp_UpdateScraperLoggingConfiguration(cfg aws.Config, client *amp.Client) {
	input := &amp.UpdateScraperLoggingConfigurationInput{
		// LoggingDestination: types.ScraperLoggingDestination, // Required
		// ScraperId: *string, // Required
	}

	if len(_ampLoggingDestination) > 0 {
		if err := assignInputField(input, "LoggingDestination", _ampLoggingDestination); err != nil {
			log.Errorf("invalid --logging-destination: %s", err.Error())
			return
		}
	}
	if len(_ampScraperId) > 0 {
		input.ScraperId = aws.String(_ampScraperId)
	}
	if len(_ampScraperComponents) > 0 {
		if err := assignInputField(input, "ScraperComponents", _ampScraperComponents); err != nil {
			log.Errorf("invalid --scraper-components: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateScraperLoggingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the alias of an existing workspace.
func amp_UpdateWorkspaceAlias(cfg aws.Config, client *amp.Client) {
	input := &amp.UpdateWorkspaceAliasInput{
		// WorkspaceId: *string, // Required
	}

	if len(_ampWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_ampWorkspaceId)
	}
	if len(_ampAlias) > 0 {
		input.Alias = aws.String(_ampAlias)
	}
	if len(_ampClientToken) > 0 {
		input.ClientToken = aws.String(_ampClientToken)
	}

	if resp, err := client.UpdateWorkspaceAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use this operation to create or update the label sets, label set limits, and
// retention period of a workspace.
//
// You must specify at least one of limitsPerLabelSet or retentionPeriodInDays for
// the request to be valid.
func amp_UpdateWorkspaceConfiguration(cfg aws.Config, client *amp.Client) {
	input := &amp.UpdateWorkspaceConfigurationInput{
		// WorkspaceId: *string, // Required
	}

	if len(_ampWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_ampWorkspaceId)
	}
	if len(_ampClientToken) > 0 {
		input.ClientToken = aws.String(_ampClientToken)
	}
	if len(_ampLimitsPerLabelSet) > 0 {
		if err := assignInputField(input, "LimitsPerLabelSet", _ampLimitsPerLabelSet); err != nil {
			log.Errorf("invalid --limits-per-label-set: %s", err.Error())
			return
		}
	}
	if len(_ampRetentionPeriodInDays) > 0 {
		if err := assignInputField(input, "RetentionPeriodInDays", _ampRetentionPeriodInDays); err != nil {
			log.Errorf("invalid --retention-period-in-days: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateWorkspaceConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_ampCmd)
	_ampCmd.Flags().SortFlags = false

	_ampCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_ampCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_ampCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_ampCmd.Flags().StringVarP(&_ampAlias, "alias", "", "", "Alias")
	_ampCmd.Flags().StringVarP(&_ampAnomalyDetectorId, "anomaly-detector-id", "", "", "Anomaly Detector ID")
	_ampCmd.Flags().StringVarP(&_ampClientToken, "client-token", "", "", "Client Token")
	_ampCmd.Flags().StringVarP(&_ampConfiguration, "configuration", "", "", "Configuration")
	_ampCmd.Flags().StringVarP(&_ampData, "data", "", "", "Data")
	_ampCmd.Flags().StringVarP(&_ampDestination, "destination", "", "", "Destination")
	_ampCmd.Flags().StringVarP(&_ampDestinations, "destinations", "", "", "Destinations")
	_ampCmd.Flags().StringVarP(&_ampEvaluationIntervalInSeconds, "evaluation-interval-in-seconds", "", "", "Evaluation Interval In Seconds")
	_ampCmd.Flags().StringVarP(&_ampFilters, "filters", "", "", "Filters")
	_ampCmd.Flags().StringVarP(&_ampKmsKeyArn, "kms-key-arn", "", "", "KMS Key ARN")
	_ampCmd.Flags().StringVarP(&_ampLabels, "labels", "", "", "Labels")
	_ampCmd.Flags().StringVarP(&_ampLimitsPerLabelSet, "limits-per-label-set", "", "", "Limits Per Label Set")
	_ampCmd.Flags().StringVarP(&_ampLogGroupArn, "log-group-arn", "", "", "Log Group ARN")
	_ampCmd.Flags().StringVarP(&_ampLoggingDestination, "logging-destination", "", "", "Logging Destination")
	_ampCmd.Flags().StringVarP(&_ampMaxResults, "max-results", "", "", "Max Results")
	_ampCmd.Flags().StringVarP(&_ampMissingDataAction, "missing-data-action", "", "", "Missing Data Action")
	_ampCmd.Flags().StringVarP(&_ampName, "name", "", "", "Name")
	_ampCmd.Flags().StringVarP(&_ampNextToken, "next-token", "", "", "Next Token")
	_ampCmd.Flags().StringVarP(&_ampPolicyDocument, "policy-document", "", "", "Policy Document")
	_ampCmd.Flags().StringVarP(&_ampResourceArn, "resource-arn", "", "", "Resource ARN")
	_ampCmd.Flags().StringVarP(&_ampRetentionPeriodInDays, "retention-period-in-days", "", "", "Retention Period In Days")
	_ampCmd.Flags().StringVarP(&_ampRevisionId, "revision-id", "", "", "Revision ID")
	_ampCmd.Flags().StringVarP(&_ampRoleConfiguration, "role-configuration", "", "", "Role Configuration")
	_ampCmd.Flags().StringVarP(&_ampScrapeConfiguration, "scrape-configuration", "", "", "Scrape Configuration")
	_ampCmd.Flags().StringVarP(&_ampScraperComponents, "scraper-components", "", "", "Scraper Components")
	_ampCmd.Flags().StringVarP(&_ampScraperId, "scraper-id", "", "", "Scraper ID")
	_ampCmd.Flags().StringVarP(&_ampSource, "source", "", "", "Source")
	_ampCmd.Flags().StringSliceVarP(&_ampTagKeys, "tag-keys", "", nil, "Tag Keys")
	_ampCmd.Flags().StringVarP(&_ampTags, "tags", "", "", "Tags")
	_ampCmd.Flags().StringVarP(&_ampWorkspaceId, "workspace-id", "", "", "Workspace ID")

	_ampCmd.Flags().BoolVarP(&_ampCreateAlertManagerDefinition, "create-alert-manager-definition", "", false, "Create Alert Manager Definition")
	_ampCmd.Flags().BoolVarP(&_ampCreateAnomalyDetector, "create-anomaly-detector", "", false, "Create Anomaly Detector")
	_ampCmd.Flags().BoolVarP(&_ampCreateLoggingConfiguration, "create-logging-configuration", "", false, "Create Logging Configuration")
	_ampCmd.Flags().BoolVarP(&_ampCreateQueryLoggingConfiguration, "create-query-logging-configuration", "", false, "Create Query Logging Configuration")
	_ampCmd.Flags().BoolVarP(&_ampCreateRuleGroupsNamespace, "create-rule-groups-namespace", "", false, "Create Rule Groups Namespace")
	_ampCmd.Flags().BoolVarP(&_ampCreateScraper, "create-scraper", "", false, "Create Scraper")
	_ampCmd.Flags().BoolVarP(&_ampCreateWorkspace, "create-workspace", "", false, "Create Workspace")
	_ampCmd.Flags().BoolVarP(&_ampDeleteAlertManagerDefinition, "delete-alert-manager-definition", "", false, "Delete Alert Manager Definition")
	_ampCmd.Flags().BoolVarP(&_ampDeleteAnomalyDetector, "delete-anomaly-detector", "", false, "Delete Anomaly Detector")
	_ampCmd.Flags().BoolVarP(&_ampDeleteLoggingConfiguration, "delete-logging-configuration", "", false, "Delete Logging Configuration")
	_ampCmd.Flags().BoolVarP(&_ampDeleteQueryLoggingConfiguration, "delete-query-logging-configuration", "", false, "Delete Query Logging Configuration")
	_ampCmd.Flags().BoolVarP(&_ampDeleteResourcePolicy, "delete-resource-policy", "", false, "Delete Resource Policy")
	_ampCmd.Flags().BoolVarP(&_ampDeleteRuleGroupsNamespace, "delete-rule-groups-namespace", "", false, "Delete Rule Groups Namespace")
	_ampCmd.Flags().BoolVarP(&_ampDeleteScraper, "delete-scraper", "", false, "Delete Scraper")
	_ampCmd.Flags().BoolVarP(&_ampDeleteScraperLoggingConfiguration, "delete-scraper-logging-configuration", "", false, "Delete Scraper Logging Configuration")
	_ampCmd.Flags().BoolVarP(&_ampDeleteWorkspace, "delete-workspace", "", false, "Delete Workspace")
	_ampCmd.Flags().BoolVarP(&_ampDescribeAlertManagerDefinition, "describe-alert-manager-definition", "", false, "Describe Alert Manager Definition")
	_ampCmd.Flags().BoolVarP(&_ampDescribeAnomalyDetector, "describe-anomaly-detector", "", false, "Describe Anomaly Detector")
	_ampCmd.Flags().BoolVarP(&_ampDescribeLoggingConfiguration, "describe-logging-configuration", "", false, "Describe Logging Configuration")
	_ampCmd.Flags().BoolVarP(&_ampDescribeQueryLoggingConfiguration, "describe-query-logging-configuration", "", false, "Describe Query Logging Configuration")
	_ampCmd.Flags().BoolVarP(&_ampDescribeResourcePolicy, "describe-resource-policy", "", false, "Describe Resource Policy")
	_ampCmd.Flags().BoolVarP(&_ampDescribeRuleGroupsNamespace, "describe-rule-groups-namespace", "", false, "Describe Rule Groups Namespace")
	_ampCmd.Flags().BoolVarP(&_ampDescribeScraper, "describe-scraper", "", false, "Describe Scraper")
	_ampCmd.Flags().BoolVarP(&_ampDescribeScraperLoggingConfiguration, "describe-scraper-logging-configuration", "", false, "Describe Scraper Logging Configuration")
	_ampCmd.Flags().BoolVarP(&_ampDescribeWorkspace, "describe-workspace", "", false, "Describe Workspace")
	_ampCmd.Flags().BoolVarP(&_ampDescribeWorkspaceConfiguration, "describe-workspace-configuration", "", false, "Describe Workspace Configuration")
	_ampCmd.Flags().BoolVarP(&_ampGetDefaultScraperConfiguration, "get-default-scraper-configuration", "", false, "Get Default Scraper Configuration")
	_ampCmd.Flags().BoolVarP(&_ampListAnomalyDetectors, "list-anomaly-detectors", "", false, "List Anomaly Detectors")
	_ampCmd.Flags().BoolVarP(&_ampListRuleGroupsNamespaces, "list-rule-groups-namespaces", "", false, "List Rule Groups Namespaces")
	_ampCmd.Flags().BoolVarP(&_ampListScrapers, "list-scrapers", "", false, "List Scrapers")
	_ampCmd.Flags().BoolVarP(&_ampListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_ampCmd.Flags().BoolVarP(&_ampListWorkspaces, "list-workspaces", "", false, "List Workspaces")
	_ampCmd.Flags().BoolVarP(&_ampPutAlertManagerDefinition, "put-alert-manager-definition", "", false, "Put Alert Manager Definition")
	_ampCmd.Flags().BoolVarP(&_ampPutAnomalyDetector, "put-anomaly-detector", "", false, "Put Anomaly Detector")
	_ampCmd.Flags().BoolVarP(&_ampPutResourcePolicy, "put-resource-policy", "", false, "Put Resource Policy")
	_ampCmd.Flags().BoolVarP(&_ampPutRuleGroupsNamespace, "put-rule-groups-namespace", "", false, "Put Rule Groups Namespace")
	_ampCmd.Flags().BoolVarP(&_ampTagResource, "tag-resource", "", false, "Tag Resource")
	_ampCmd.Flags().BoolVarP(&_ampUntagResource, "untag-resource", "", false, "Untag Resource")
	_ampCmd.Flags().BoolVarP(&_ampUpdateLoggingConfiguration, "update-logging-configuration", "", false, "Update Logging Configuration")
	_ampCmd.Flags().BoolVarP(&_ampUpdateQueryLoggingConfiguration, "update-query-logging-configuration", "", false, "Update Query Logging Configuration")
	_ampCmd.Flags().BoolVarP(&_ampUpdateScraper, "update-scraper", "", false, "Update Scraper")
	_ampCmd.Flags().BoolVarP(&_ampUpdateScraperLoggingConfiguration, "update-scraper-logging-configuration", "", false, "Update Scraper Logging Configuration")
	_ampCmd.Flags().BoolVarP(&_ampUpdateWorkspaceAlias, "update-workspace-alias", "", false, "Update Workspace Alias")
	_ampCmd.Flags().BoolVarP(&_ampUpdateWorkspaceConfiguration, "update-workspace-configuration", "", false, "Update Workspace Configuration")

}
