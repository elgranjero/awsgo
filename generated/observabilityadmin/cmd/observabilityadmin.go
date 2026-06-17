package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/observabilityadmin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// observabilityadminCmd represents the observabilityadmin command
var _observabilityadminCmd = &cobra.Command{
	Use:   "observabilityadmin",
	Short: "AWS observabilityadmin CLI",
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
		client := observabilityadmin.NewFromConfig(cfg)
		if _observabilityadminCreateCentralizationRuleForOrganization {
			observabilityadmin_CreateCentralizationRuleForOrganization(cfg, client)
			return
		}
		if _observabilityadminCreateS3TableIntegration {
			observabilityadmin_CreateS3TableIntegration(cfg, client)
			return
		}
		if _observabilityadminCreateTelemetryPipeline {
			observabilityadmin_CreateTelemetryPipeline(cfg, client)
			return
		}
		if _observabilityadminCreateTelemetryRule {
			observabilityadmin_CreateTelemetryRule(cfg, client)
			return
		}
		if _observabilityadminCreateTelemetryRuleForOrganization {
			observabilityadmin_CreateTelemetryRuleForOrganization(cfg, client)
			return
		}
		if _observabilityadminDeleteCentralizationRuleForOrganization {
			observabilityadmin_DeleteCentralizationRuleForOrganization(cfg, client)
			return
		}
		if _observabilityadminDeleteS3TableIntegration {
			observabilityadmin_DeleteS3TableIntegration(cfg, client)
			return
		}
		if _observabilityadminDeleteTelemetryPipeline {
			observabilityadmin_DeleteTelemetryPipeline(cfg, client)
			return
		}
		if _observabilityadminDeleteTelemetryRule {
			observabilityadmin_DeleteTelemetryRule(cfg, client)
			return
		}
		if _observabilityadminDeleteTelemetryRuleForOrganization {
			observabilityadmin_DeleteTelemetryRuleForOrganization(cfg, client)
			return
		}
		if _observabilityadminGetCentralizationRuleForOrganization {
			observabilityadmin_GetCentralizationRuleForOrganization(cfg, client)
			return
		}
		if _observabilityadminGetS3TableIntegration {
			observabilityadmin_GetS3TableIntegration(cfg, client)
			return
		}
		if _observabilityadminGetTelemetryEnrichmentStatus {
			observabilityadmin_GetTelemetryEnrichmentStatus(cfg, client)
			return
		}
		if _observabilityadminGetTelemetryEvaluationStatus {
			observabilityadmin_GetTelemetryEvaluationStatus(cfg, client)
			return
		}
		if _observabilityadminGetTelemetryEvaluationStatusForOrganization {
			observabilityadmin_GetTelemetryEvaluationStatusForOrganization(cfg, client)
			return
		}
		if _observabilityadminGetTelemetryPipeline {
			observabilityadmin_GetTelemetryPipeline(cfg, client)
			return
		}
		if _observabilityadminGetTelemetryRule {
			observabilityadmin_GetTelemetryRule(cfg, client)
			return
		}
		if _observabilityadminGetTelemetryRuleForOrganization {
			observabilityadmin_GetTelemetryRuleForOrganization(cfg, client)
			return
		}
		if _observabilityadminListCentralizationRulesForOrganization {
			observabilityadmin_ListCentralizationRulesForOrganization(cfg, client)
			return
		}
		if _observabilityadminListResourceTelemetry {
			observabilityadmin_ListResourceTelemetry(cfg, client)
			return
		}
		if _observabilityadminListResourceTelemetryForOrganization {
			observabilityadmin_ListResourceTelemetryForOrganization(cfg, client)
			return
		}
		if _observabilityadminListS3TableIntegrations {
			observabilityadmin_ListS3TableIntegrations(cfg, client)
			return
		}
		if _observabilityadminListTagsForResource {
			observabilityadmin_ListTagsForResource(cfg, client)
			return
		}
		if _observabilityadminListTelemetryPipelines {
			observabilityadmin_ListTelemetryPipelines(cfg, client)
			return
		}
		if _observabilityadminListTelemetryRules {
			observabilityadmin_ListTelemetryRules(cfg, client)
			return
		}
		if _observabilityadminListTelemetryRulesForOrganization {
			observabilityadmin_ListTelemetryRulesForOrganization(cfg, client)
			return
		}
		if _observabilityadminStartTelemetryEnrichment {
			observabilityadmin_StartTelemetryEnrichment(cfg, client)
			return
		}
		if _observabilityadminStartTelemetryEvaluation {
			observabilityadmin_StartTelemetryEvaluation(cfg, client)
			return
		}
		if _observabilityadminStartTelemetryEvaluationForOrganization {
			observabilityadmin_StartTelemetryEvaluationForOrganization(cfg, client)
			return
		}
		if _observabilityadminStopTelemetryEnrichment {
			observabilityadmin_StopTelemetryEnrichment(cfg, client)
			return
		}
		if _observabilityadminStopTelemetryEvaluation {
			observabilityadmin_StopTelemetryEvaluation(cfg, client)
			return
		}
		if _observabilityadminStopTelemetryEvaluationForOrganization {
			observabilityadmin_StopTelemetryEvaluationForOrganization(cfg, client)
			return
		}
		if _observabilityadminTagResource {
			observabilityadmin_TagResource(cfg, client)
			return
		}
		if _observabilityadminTestTelemetryPipeline {
			observabilityadmin_TestTelemetryPipeline(cfg, client)
			return
		}
		if _observabilityadminUntagResource {
			observabilityadmin_UntagResource(cfg, client)
			return
		}
		if _observabilityadminUpdateCentralizationRuleForOrganization {
			observabilityadmin_UpdateCentralizationRuleForOrganization(cfg, client)
			return
		}
		if _observabilityadminUpdateTelemetryPipeline {
			observabilityadmin_UpdateTelemetryPipeline(cfg, client)
			return
		}
		if _observabilityadminUpdateTelemetryRule {
			observabilityadmin_UpdateTelemetryRule(cfg, client)
			return
		}
		if _observabilityadminUpdateTelemetryRuleForOrganization {
			observabilityadmin_UpdateTelemetryRuleForOrganization(cfg, client)
			return
		}
		if _observabilityadminValidateTelemetryPipelineConfiguration {
			observabilityadmin_ValidateTelemetryPipelineConfiguration(cfg, client)
			return
		}

	},
}

var (
	_observabilityadminCreateCentralizationRuleForOrganization     bool
	_observabilityadminCreateS3TableIntegration                    bool
	_observabilityadminCreateTelemetryPipeline                     bool
	_observabilityadminCreateTelemetryRule                         bool
	_observabilityadminCreateTelemetryRuleForOrganization          bool
	_observabilityadminDeleteCentralizationRuleForOrganization     bool
	_observabilityadminDeleteS3TableIntegration                    bool
	_observabilityadminDeleteTelemetryPipeline                     bool
	_observabilityadminDeleteTelemetryRule                         bool
	_observabilityadminDeleteTelemetryRuleForOrganization          bool
	_observabilityadminGetCentralizationRuleForOrganization        bool
	_observabilityadminGetS3TableIntegration                       bool
	_observabilityadminGetTelemetryEnrichmentStatus                bool
	_observabilityadminGetTelemetryEvaluationStatus                bool
	_observabilityadminGetTelemetryEvaluationStatusForOrganization bool
	_observabilityadminGetTelemetryPipeline                        bool
	_observabilityadminGetTelemetryRule                            bool
	_observabilityadminGetTelemetryRuleForOrganization             bool
	_observabilityadminListCentralizationRulesForOrganization      bool
	_observabilityadminListResourceTelemetry                       bool
	_observabilityadminListResourceTelemetryForOrganization        bool
	_observabilityadminListS3TableIntegrations                     bool
	_observabilityadminListTagsForResource                         bool
	_observabilityadminListTelemetryPipelines                      bool
	_observabilityadminListTelemetryRules                          bool
	_observabilityadminListTelemetryRulesForOrganization           bool
	_observabilityadminStartTelemetryEnrichment                    bool
	_observabilityadminStartTelemetryEvaluation                    bool
	_observabilityadminStartTelemetryEvaluationForOrganization     bool
	_observabilityadminStopTelemetryEnrichment                     bool
	_observabilityadminStopTelemetryEvaluation                     bool
	_observabilityadminStopTelemetryEvaluationForOrganization      bool
	_observabilityadminTagResource                                 bool
	_observabilityadminTestTelemetryPipeline                       bool
	_observabilityadminUntagResource                               bool
	_observabilityadminUpdateCentralizationRuleForOrganization     bool
	_observabilityadminUpdateTelemetryPipeline                     bool
	_observabilityadminUpdateTelemetryRule                         bool
	_observabilityadminUpdateTelemetryRuleForOrganization          bool
	_observabilityadminValidateTelemetryPipelineConfiguration      bool

	_observabilityadminAccountIdentifiers          []string
	_observabilityadminAllRegions                  string
	_observabilityadminArn                         string
	_observabilityadminConfiguration               string
	_observabilityadminEncryption                  string
	_observabilityadminMaxResults                  string
	_observabilityadminName                        string
	_observabilityadminNextToken                   string
	_observabilityadminPipelineIdentifier          string
	_observabilityadminRecords                     string
	_observabilityadminResourceARN                 string
	_observabilityadminResourceIdentifierPrefix    string
	_observabilityadminResourceTags                string
	_observabilityadminResourceTypes               string
	_observabilityadminRoleArn                     string
	_observabilityadminRule                        string
	_observabilityadminRuleIdentifier              string
	_observabilityadminRuleName                    string
	_observabilityadminRuleNamePrefix              string
	_observabilityadminSourceAccountIds            []string
	_observabilityadminSourceOrganizationUnitIds   []string
	_observabilityadminTagKeys                     []string
	_observabilityadminTags                        string
	_observabilityadminTelemetryConfigurationState string
)

// Creates a centralization rule that applies across an Amazon Web Services
// Organization. This operation can only be called by the organization's management
// account or a delegated administrator account.
func observabilityadmin_CreateCentralizationRuleForOrganization(cfg aws.Config, client *observabilityadmin.Client) {
	input := &observabilityadmin.CreateCentralizationRuleForOrganizationInput{
		// Rule: *types.CentralizationRule, // Required
		// RuleName: *string, // Required
	}

	if len(_observabilityadminRule) > 0 {
		if err := assignInputField(input, "Rule", _observabilityadminRule); err != nil {
			log.Errorf("invalid --rule: %s", err.Error())
			return
		}
	}
	if len(_observabilityadminRuleName) > 0 {
		input.RuleName = aws.String(_observabilityadminRuleName)
	}
	if len(_observabilityadminTags) > 0 {
		if err := assignInputField(input, "Tags", _observabilityadminTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCentralizationRuleForOrganization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an integration between CloudWatch and S3 Tables for analytics. This
// integration enables querying CloudWatch telemetry data using analytics engines
// like Amazon Athena, Amazon Redshift, and Apache Spark.
func observabilityadmin_CreateS3TableIntegration(cfg aws.Config, client *observabilityadmin.Client) {
	input := &observabilityadmin.CreateS3TableIntegrationInput{
		// Encryption: *types.Encryption, // Required
		// RoleArn: *string, // Required
	}

	if len(_observabilityadminEncryption) > 0 {
		if err := assignInputField(input, "Encryption", _observabilityadminEncryption); err != nil {
			log.Errorf("invalid --encryption: %s", err.Error())
			return
		}
	}
	if len(_observabilityadminRoleArn) > 0 {
		input.RoleArn = aws.String(_observabilityadminRoleArn)
	}
	if len(_observabilityadminTags) > 0 {
		if err := assignInputField(input, "Tags", _observabilityadminTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateS3TableIntegration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a telemetry pipeline for processing and transforming telemetry data.
// The pipeline defines how data flows from sources through processors to
// destinations, enabling data transformation and delivering capabilities.
func observabilityadmin_CreateTelemetryPipeline(cfg aws.Config, client *observabilityadmin.Client) {
	input := &observabilityadmin.CreateTelemetryPipelineInput{
		// Configuration: *types.TelemetryPipelineConfiguration, // Required
		// Name: *string, // Required
	}

	if len(_observabilityadminConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _observabilityadminConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_observabilityadminName) > 0 {
		input.Name = aws.String(_observabilityadminName)
	}
	if len(_observabilityadminTags) > 0 {
		if err := assignInputField(input, "Tags", _observabilityadminTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTelemetryPipeline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a telemetry rule that defines how telemetry should be configured for
// Amazon Web Services resources in your account. The rule specifies which
// resources should have telemetry enabled and how that telemetry data should be
// collected based on resource type, telemetry type, and selection criteria.
func observabilityadmin_CreateTelemetryRule(cfg aws.Config, client *observabilityadmin.Client) {
	input := &observabilityadmin.CreateTelemetryRuleInput{
		// Rule: *types.TelemetryRule, // Required
		// RuleName: *string, // Required
	}

	if len(_observabilityadminRule) > 0 {
		if err := assignInputField(input, "Rule", _observabilityadminRule); err != nil {
			log.Errorf("invalid --rule: %s", err.Error())
			return
		}
	}
	if len(_observabilityadminRuleName) > 0 {
		input.RuleName = aws.String(_observabilityadminRuleName)
	}
	if len(_observabilityadminTags) > 0 {
		if err := assignInputField(input, "Tags", _observabilityadminTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTelemetryRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a telemetry rule that applies across an Amazon Web Services
// Organization. This operation can only be called by the organization's management
// account or a delegated administrator account.
func observabilityadmin_CreateTelemetryRuleForOrganization(cfg aws.Config, client *observabilityadmin.Client) {
	input := &observabilityadmin.CreateTelemetryRuleForOrganizationInput{
		// Rule: *types.TelemetryRule, // Required
		// RuleName: *string, // Required
	}

	if len(_observabilityadminRule) > 0 {
		if err := assignInputField(input, "Rule", _observabilityadminRule); err != nil {
			log.Errorf("invalid --rule: %s", err.Error())
			return
		}
	}
	if len(_observabilityadminRuleName) > 0 {
		input.RuleName = aws.String(_observabilityadminRuleName)
	}
	if len(_observabilityadminTags) > 0 {
		if err := assignInputField(input, "Tags", _observabilityadminTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTelemetryRuleForOrganization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an organization-wide centralization rule. This operation can only be
// called by the organization's management account or a delegated administrator
// account.
func observabilityadmin_DeleteCentralizationRuleForOrganization(cfg aws.Config, client *observabilityadmin.Client) {
	input := &observabilityadmin.DeleteCentralizationRuleForOrganizationInput{
		// RuleIdentifier: *string, // Required
	}

	if len(_observabilityadminRuleIdentifier) > 0 {
		input.RuleIdentifier = aws.String(_observabilityadminRuleIdentifier)
	}

	if resp, err := client.DeleteCentralizationRuleForOrganization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an S3 Table integration and its associated data. This operation removes
// the connection between CloudWatch Observability Admin and S3 Tables.
func observabilityadmin_DeleteS3TableIntegration(cfg aws.Config, client *observabilityadmin.Client) {
	input := &observabilityadmin.DeleteS3TableIntegrationInput{
		// Arn: *string, // Required
	}

	if len(_observabilityadminArn) > 0 {
		input.Arn = aws.String(_observabilityadminArn)
	}

	if resp, err := client.DeleteS3TableIntegration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a telemetry pipeline and its associated resources. This operation stops
// data processing and removes the pipeline configuration.
func observabilityadmin_DeleteTelemetryPipeline(cfg aws.Config, client *observabilityadmin.Client) {
	input := &observabilityadmin.DeleteTelemetryPipelineInput{
		// PipelineIdentifier: *string, // Required
	}

	if len(_observabilityadminPipelineIdentifier) > 0 {
		input.PipelineIdentifier = aws.String(_observabilityadminPipelineIdentifier)
	}

	if resp, err := client.DeleteTelemetryPipeline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a telemetry rule from your account. Any telemetry configurations
// previously created by the rule will remain but no new resources will be
// configured by this rule.
func observabilityadmin_DeleteTelemetryRule(cfg aws.Config, client *observabilityadmin.Client) {
	input := &observabilityadmin.DeleteTelemetryRuleInput{
		// RuleIdentifier: *string, // Required
	}

	if len(_observabilityadminRuleIdentifier) > 0 {
		input.RuleIdentifier = aws.String(_observabilityadminRuleIdentifier)
	}

	if resp, err := client.DeleteTelemetryRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an organization-wide telemetry rule. This operation can only be called
// by the organization's management account or a delegated administrator account.
func observabilityadmin_DeleteTelemetryRuleForOrganization(cfg aws.Config, client *observabilityadmin.Client) {
	input := &observabilityadmin.DeleteTelemetryRuleForOrganizationInput{
		// RuleIdentifier: *string, // Required
	}

	if len(_observabilityadminRuleIdentifier) > 0 {
		input.RuleIdentifier = aws.String(_observabilityadminRuleIdentifier)
	}

	if resp, err := client.DeleteTelemetryRuleForOrganization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the details of a specific organization centralization rule. This
// operation can only be called by the organization's management account or a
// delegated administrator account.
func observabilityadmin_GetCentralizationRuleForOrganization(cfg aws.Config, client *observabilityadmin.Client) {
	input := &observabilityadmin.GetCentralizationRuleForOrganizationInput{
		// RuleIdentifier: *string, // Required
	}

	if len(_observabilityadminRuleIdentifier) > 0 {
		input.RuleIdentifier = aws.String(_observabilityadminRuleIdentifier)
	}

	if resp, err := client.GetCentralizationRuleForOrganization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a specific S3 Table integration, including its
// configuration, status, and metadata.
func observabilityadmin_GetS3TableIntegration(cfg aws.Config, client *observabilityadmin.Client) {
	input := &observabilityadmin.GetS3TableIntegrationInput{
		// Arn: *string, // Required
	}

	if len(_observabilityadminArn) > 0 {
		input.Arn = aws.String(_observabilityadminArn)
	}

	if resp, err := client.GetS3TableIntegration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the current status of the resource tags for telemetry feature, which
// enhances telemetry data with additional resource metadata from Resource
// Explorer.
func observabilityadmin_GetTelemetryEnrichmentStatus(cfg aws.Config, client *observabilityadmin.Client) {
	input := &observabilityadmin.GetTelemetryEnrichmentStatusInput{}

	if resp, err := client.GetTelemetryEnrichmentStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the current onboarding status of the telemetry config feature,
// including the status of the feature and reason the feature failed to start or
// stop.
func observabilityadmin_GetTelemetryEvaluationStatus(cfg aws.Config, client *observabilityadmin.Client) {
	input := &observabilityadmin.GetTelemetryEvaluationStatusInput{}

	if resp, err := client.GetTelemetryEvaluationStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This returns the onboarding status of the telemetry configuration feature for
// the organization. It can only be called by a Management Account of an Amazon Web
// Services Organization or an assigned Delegated Admin Account of Amazon
// CloudWatch telemetry config.
func observabilityadmin_GetTelemetryEvaluationStatusForOrganization(cfg aws.Config, client *observabilityadmin.Client) {
	input := &observabilityadmin.GetTelemetryEvaluationStatusForOrganizationInput{}

	if resp, err := client.GetTelemetryEvaluationStatusForOrganization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a specific telemetry pipeline, including its
// configuration, status, and metadata.
func observabilityadmin_GetTelemetryPipeline(cfg aws.Config, client *observabilityadmin.Client) {
	input := &observabilityadmin.GetTelemetryPipelineInput{
		// PipelineIdentifier: *string, // Required
	}

	if len(_observabilityadminPipelineIdentifier) > 0 {
		input.PipelineIdentifier = aws.String(_observabilityadminPipelineIdentifier)
	}

	if resp, err := client.GetTelemetryPipeline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the details of a specific telemetry rule in your account.
func observabilityadmin_GetTelemetryRule(cfg aws.Config, client *observabilityadmin.Client) {
	input := &observabilityadmin.GetTelemetryRuleInput{
		// RuleIdentifier: *string, // Required
	}

	if len(_observabilityadminRuleIdentifier) > 0 {
		input.RuleIdentifier = aws.String(_observabilityadminRuleIdentifier)
	}

	if resp, err := client.GetTelemetryRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the details of a specific organization telemetry rule. This
// operation can only be called by the organization's management account or a
// delegated administrator account.
func observabilityadmin_GetTelemetryRuleForOrganization(cfg aws.Config, client *observabilityadmin.Client) {
	input := &observabilityadmin.GetTelemetryRuleForOrganizationInput{
		// RuleIdentifier: *string, // Required
	}

	if len(_observabilityadminRuleIdentifier) > 0 {
		input.RuleIdentifier = aws.String(_observabilityadminRuleIdentifier)
	}

	if resp, err := client.GetTelemetryRuleForOrganization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all centralization rules in your organization. This operation can only be
// called by the organization's management account or a delegated administrator
// account.
func observabilityadmin_ListCentralizationRulesForOrganization(cfg aws.Config, client *observabilityadmin.Client) {
	input := &observabilityadmin.ListCentralizationRulesForOrganizationInput{}

	if len(_observabilityadminAllRegions) > 0 {
		if err := assignInputField(input, "AllRegions", _observabilityadminAllRegions); err != nil {
			log.Errorf("invalid --all-regions: %s", err.Error())
			return
		}
	}
	if len(_observabilityadminMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _observabilityadminMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_observabilityadminNextToken) > 0 {
		input.NextToken = aws.String(_observabilityadminNextToken)
	}
	if len(_observabilityadminRuleNamePrefix) > 0 {
		input.RuleNamePrefix = aws.String(_observabilityadminRuleNamePrefix)
	}

	if disablePaginator() {
		if resp, err := client.ListCentralizationRulesForOrganization(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*observabilityadmin.ListCentralizationRulesForOrganizationOutput
	p := observabilityadmin.NewListCentralizationRulesForOrganizationPaginator(client, input)
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

// Returns a list of telemetry configurations for Amazon Web Services resources
// supported by telemetry config. For more information, see [Auditing CloudWatch telemetry configurations].
//
// [Auditing CloudWatch telemetry configurations]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/telemetry-config-cloudwatch.html
func observabilityadmin_ListResourceTelemetry(cfg aws.Config, client *observabilityadmin.Client) {
	input := &observabilityadmin.ListResourceTelemetryInput{}

	if len(_observabilityadminMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _observabilityadminMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_observabilityadminNextToken) > 0 {
		input.NextToken = aws.String(_observabilityadminNextToken)
	}
	if len(_observabilityadminResourceIdentifierPrefix) > 0 {
		input.ResourceIdentifierPrefix = aws.String(_observabilityadminResourceIdentifierPrefix)
	}
	if len(_observabilityadminResourceTags) > 0 {
		if err := assignInputField(input, "ResourceTags", _observabilityadminResourceTags); err != nil {
			log.Errorf("invalid --resource-tags: %s", err.Error())
			return
		}
	}
	if len(_observabilityadminResourceTypes) > 0 {
		if err := assignInputField(input, "ResourceTypes", _observabilityadminResourceTypes); err != nil {
			log.Errorf("invalid --resource-types: %s", err.Error())
			return
		}
	}
	if len(_observabilityadminTelemetryConfigurationState) > 0 {
		if err := assignInputField(input, "TelemetryConfigurationState", _observabilityadminTelemetryConfigurationState); err != nil {
			log.Errorf("invalid --telemetry-configuration-state: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListResourceTelemetry(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*observabilityadmin.ListResourceTelemetryOutput
	p := observabilityadmin.NewListResourceTelemetryPaginator(client, input)
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

// Returns a list of telemetry configurations for Amazon Web Services resources
// supported by telemetry config in the organization.
func observabilityadmin_ListResourceTelemetryForOrganization(cfg aws.Config, client *observabilityadmin.Client) {
	input := &observabilityadmin.ListResourceTelemetryForOrganizationInput{}

	if len(_observabilityadminAccountIdentifiers) > 0 {
		input.AccountIdentifiers = append([]string(nil), _observabilityadminAccountIdentifiers...)
	}
	if len(_observabilityadminMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _observabilityadminMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_observabilityadminNextToken) > 0 {
		input.NextToken = aws.String(_observabilityadminNextToken)
	}
	if len(_observabilityadminResourceIdentifierPrefix) > 0 {
		input.ResourceIdentifierPrefix = aws.String(_observabilityadminResourceIdentifierPrefix)
	}
	if len(_observabilityadminResourceTags) > 0 {
		if err := assignInputField(input, "ResourceTags", _observabilityadminResourceTags); err != nil {
			log.Errorf("invalid --resource-tags: %s", err.Error())
			return
		}
	}
	if len(_observabilityadminResourceTypes) > 0 {
		if err := assignInputField(input, "ResourceTypes", _observabilityadminResourceTypes); err != nil {
			log.Errorf("invalid --resource-types: %s", err.Error())
			return
		}
	}
	if len(_observabilityadminTelemetryConfigurationState) > 0 {
		if err := assignInputField(input, "TelemetryConfigurationState", _observabilityadminTelemetryConfigurationState); err != nil {
			log.Errorf("invalid --telemetry-configuration-state: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListResourceTelemetryForOrganization(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*observabilityadmin.ListResourceTelemetryForOrganizationOutput
	p := observabilityadmin.NewListResourceTelemetryForOrganizationPaginator(client, input)
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

// Lists all S3 Table integrations in your account. We recommend using pagination
// to ensure that the operation returns quickly and successfully.
func observabilityadmin_ListS3TableIntegrations(cfg aws.Config, client *observabilityadmin.Client) {
	input := &observabilityadmin.ListS3TableIntegrationsInput{}

	if len(_observabilityadminMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _observabilityadminMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_observabilityadminNextToken) > 0 {
		input.NextToken = aws.String(_observabilityadminNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListS3TableIntegrations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*observabilityadmin.ListS3TableIntegrationsOutput
	p := observabilityadmin.NewListS3TableIntegrationsPaginator(client, input)
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

// Lists all tags attached to the specified resource. Supports telemetry rule
// resources and telemetry pipeline resources.
func observabilityadmin_ListTagsForResource(cfg aws.Config, client *observabilityadmin.Client) {
	input := &observabilityadmin.ListTagsForResourceInput{
		// ResourceARN: *string, // Required
	}

	if len(_observabilityadminResourceARN) > 0 {
		input.ResourceARN = aws.String(_observabilityadminResourceARN)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of telemetry pipelines in your account. Returns up to 100
// results. If more than 100 telemetry pipelines exist, include the NextToken
// value from the response to retrieve the next set of results.
func observabilityadmin_ListTelemetryPipelines(cfg aws.Config, client *observabilityadmin.Client) {
	input := &observabilityadmin.ListTelemetryPipelinesInput{}

	if len(_observabilityadminMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _observabilityadminMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_observabilityadminNextToken) > 0 {
		input.NextToken = aws.String(_observabilityadminNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTelemetryPipelines(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*observabilityadmin.ListTelemetryPipelinesOutput
	p := observabilityadmin.NewListTelemetryPipelinesPaginator(client, input)
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

// Lists all telemetry rules in your account. You can filter the results by
// specifying a rule name prefix.
func observabilityadmin_ListTelemetryRules(cfg aws.Config, client *observabilityadmin.Client) {
	input := &observabilityadmin.ListTelemetryRulesInput{}

	if len(_observabilityadminMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _observabilityadminMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_observabilityadminNextToken) > 0 {
		input.NextToken = aws.String(_observabilityadminNextToken)
	}
	if len(_observabilityadminRuleNamePrefix) > 0 {
		input.RuleNamePrefix = aws.String(_observabilityadminRuleNamePrefix)
	}

	if disablePaginator() {
		if resp, err := client.ListTelemetryRules(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*observabilityadmin.ListTelemetryRulesOutput
	p := observabilityadmin.NewListTelemetryRulesPaginator(client, input)
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

// Lists all telemetry rules in your organization. This operation can only be
// called by the organization's management account or a delegated administrator
// account.
func observabilityadmin_ListTelemetryRulesForOrganization(cfg aws.Config, client *observabilityadmin.Client) {
	input := &observabilityadmin.ListTelemetryRulesForOrganizationInput{}

	if len(_observabilityadminMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _observabilityadminMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_observabilityadminNextToken) > 0 {
		input.NextToken = aws.String(_observabilityadminNextToken)
	}
	if len(_observabilityadminRuleNamePrefix) > 0 {
		input.RuleNamePrefix = aws.String(_observabilityadminRuleNamePrefix)
	}
	if len(_observabilityadminSourceAccountIds) > 0 {
		input.SourceAccountIds = append([]string(nil), _observabilityadminSourceAccountIds...)
	}
	if len(_observabilityadminSourceOrganizationUnitIds) > 0 {
		input.SourceOrganizationUnitIds = append([]string(nil), _observabilityadminSourceOrganizationUnitIds...)
	}

	if disablePaginator() {
		if resp, err := client.ListTelemetryRulesForOrganization(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*observabilityadmin.ListTelemetryRulesForOrganizationOutput
	p := observabilityadmin.NewListTelemetryRulesForOrganizationPaginator(client, input)
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

// Enables the resource tags for telemetry feature for your account, which
// enhances telemetry data with additional resource metadata from Resource Explorer
// to provide richer context for monitoring and observability.
func observabilityadmin_StartTelemetryEnrichment(cfg aws.Config, client *observabilityadmin.Client) {
	input := &observabilityadmin.StartTelemetryEnrichmentInput{}

	if resp, err := client.StartTelemetryEnrichment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This action begins onboarding the caller Amazon Web Services account to the
// telemetry config feature.
func observabilityadmin_StartTelemetryEvaluation(cfg aws.Config, client *observabilityadmin.Client) {
	input := &observabilityadmin.StartTelemetryEvaluationInput{}

	if resp, err := client.StartTelemetryEvaluation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This actions begins onboarding the organization and all member accounts to the
// telemetry config feature.
func observabilityadmin_StartTelemetryEvaluationForOrganization(cfg aws.Config, client *observabilityadmin.Client) {
	input := &observabilityadmin.StartTelemetryEvaluationForOrganizationInput{}

	if resp, err := client.StartTelemetryEvaluationForOrganization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables the resource tags for telemetry feature for your account, stopping
// the enhancement of telemetry data with additional resource metadata.
func observabilityadmin_StopTelemetryEnrichment(cfg aws.Config, client *observabilityadmin.Client) {
	input := &observabilityadmin.StopTelemetryEnrichmentInput{}

	if resp, err := client.StopTelemetryEnrichment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This action begins offboarding the caller Amazon Web Services account from the
// telemetry config feature.
func observabilityadmin_StopTelemetryEvaluation(cfg aws.Config, client *observabilityadmin.Client) {
	input := &observabilityadmin.StopTelemetryEvaluationInput{}

	if resp, err := client.StopTelemetryEvaluation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This action offboards the Organization of the caller Amazon Web Services
// account from the telemetry config feature.
func observabilityadmin_StopTelemetryEvaluationForOrganization(cfg aws.Config, client *observabilityadmin.Client) {
	input := &observabilityadmin.StopTelemetryEvaluationForOrganizationInput{}

	if resp, err := client.StopTelemetryEvaluationForOrganization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or updates tags for a resource. Supports telemetry rule resources and
// telemetry pipeline resources.
func observabilityadmin_TagResource(cfg aws.Config, client *observabilityadmin.Client) {
	input := &observabilityadmin.TagResourceInput{
		// ResourceARN: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_observabilityadminResourceARN) > 0 {
		input.ResourceARN = aws.String(_observabilityadminResourceARN)
	}
	if len(_observabilityadminTags) > 0 {
		if err := assignInputField(input, "Tags", _observabilityadminTags); err != nil {
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

// Tests a pipeline configuration with sample records to validate data processing
// before deployment. This operation helps ensure your pipeline configuration works
// as expected.
func observabilityadmin_TestTelemetryPipeline(cfg aws.Config, client *observabilityadmin.Client) {
	input := &observabilityadmin.TestTelemetryPipelineInput{
		// Configuration: *types.TelemetryPipelineConfiguration, // Required
		// Records: []types.Record, // Required
	}

	if len(_observabilityadminConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _observabilityadminConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_observabilityadminRecords) > 0 {
		if err := assignInputField(input, "Records", _observabilityadminRecords); err != nil {
			log.Errorf("invalid --records: %s", err.Error())
			return
		}
	}

	if resp, err := client.TestTelemetryPipeline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes tags from a resource. Supports telemetry rule resources and telemetry
// pipeline resources.
func observabilityadmin_UntagResource(cfg aws.Config, client *observabilityadmin.Client) {
	input := &observabilityadmin.UntagResourceInput{
		// ResourceARN: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_observabilityadminResourceARN) > 0 {
		input.ResourceARN = aws.String(_observabilityadminResourceARN)
	}
	if len(_observabilityadminTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _observabilityadminTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing centralization rule that applies across an Amazon Web
// Services Organization. This operation can only be called by the organization's
// management account or a delegated administrator account.
func observabilityadmin_UpdateCentralizationRuleForOrganization(cfg aws.Config, client *observabilityadmin.Client) {
	input := &observabilityadmin.UpdateCentralizationRuleForOrganizationInput{
		// Rule: *types.CentralizationRule, // Required
		// RuleIdentifier: *string, // Required
	}

	if len(_observabilityadminRule) > 0 {
		if err := assignInputField(input, "Rule", _observabilityadminRule); err != nil {
			log.Errorf("invalid --rule: %s", err.Error())
			return
		}
	}
	if len(_observabilityadminRuleIdentifier) > 0 {
		input.RuleIdentifier = aws.String(_observabilityadminRuleIdentifier)
	}

	if resp, err := client.UpdateCentralizationRuleForOrganization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration of an existing telemetry pipeline.
// The following attributes cannot be updated after pipeline creation:
//
// - Pipeline name - The pipeline name is immutable
//
// - Pipeline ARN - The ARN is automatically generated and cannot be changed
//
// - Source type - Once a pipeline is created with a specific source type (such
// as S3, CloudWatch Logs, GitHub, or third-party sources), it cannot be changed to
// a different source type
//
// Processors can be added, removed, or modified. However, some processors are not
// supported for third-party pipelines and cannot be added through updates.
//
// # Source-Specific Update Rules
//
// CloudWatch Logs Sources (Vended and Custom)  Updatable: sts_role_arn
//
// Fixed: data_source_name , data_source_type , sink (must remain (at)original )
//
// S3 Sources (Crowdstrike, Zscaler, SentinelOne, Custom)  Updatable: All SQS
// configuration parameters, sts_role_arn , codec settings, compression type,
// bucket ownership settings, sink log group
//
// Fixed: notification_type , aws.region
//
// GitHub Audit Logs  Updatable: All Amazon Web Services Secrets Manager
// attributes, scope (can switch between ORGANIZATION/ENTERPRISE), organization or
// enterprise name, range , authentication credentials (PAT or GitHub App)
//
// Microsoft Sources (Entra ID, Office365, Windows)  Updatable: All Amazon Web
// Services Secrets Manager attributes, tenant_id , workspace_id (Windows only),
// OAuth2 credentials ( client_id , client_secret )
//
// Okta Sources (SSO, Auth0)  Updatable: All Amazon Web Services Secrets Manager
// attributes, domain , range , OAuth2 credentials ( client_id , client_secret )
//
// Palo Alto Networks  Updatable: All Amazon Web Services Secrets Manager
// attributes, hostname , basic authentication credentials ( username , password )
//
// ServiceNow CMDB  Updatable: All Amazon Web Services Secrets Manager attributes,
// instance_url , range , OAuth2 credentials ( client_id , client_secret )
//
// Wiz CNAPP  Updatable: All Amazon Web Services Secrets Manager attributes, region
// , range , OAuth2 credentials ( client_id , client_secret )
func observabilityadmin_UpdateTelemetryPipeline(cfg aws.Config, client *observabilityadmin.Client) {
	input := &observabilityadmin.UpdateTelemetryPipelineInput{
		// Configuration: *types.TelemetryPipelineConfiguration, // Required
		// PipelineIdentifier: *string, // Required
	}

	if len(_observabilityadminConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _observabilityadminConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_observabilityadminPipelineIdentifier) > 0 {
		input.PipelineIdentifier = aws.String(_observabilityadminPipelineIdentifier)
	}

	if resp, err := client.UpdateTelemetryPipeline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing telemetry rule in your account. If multiple users attempt
// to modify the same telemetry rule simultaneously, a ConflictException is
// returned to provide specific error information for concurrent modification
// scenarios.
func observabilityadmin_UpdateTelemetryRule(cfg aws.Config, client *observabilityadmin.Client) {
	input := &observabilityadmin.UpdateTelemetryRuleInput{
		// Rule: *types.TelemetryRule, // Required
		// RuleIdentifier: *string, // Required
	}

	if len(_observabilityadminRule) > 0 {
		if err := assignInputField(input, "Rule", _observabilityadminRule); err != nil {
			log.Errorf("invalid --rule: %s", err.Error())
			return
		}
	}
	if len(_observabilityadminRuleIdentifier) > 0 {
		input.RuleIdentifier = aws.String(_observabilityadminRuleIdentifier)
	}

	if resp, err := client.UpdateTelemetryRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing telemetry rule that applies across an Amazon Web Services
// Organization. This operation can only be called by the organization's management
// account or a delegated administrator account.
func observabilityadmin_UpdateTelemetryRuleForOrganization(cfg aws.Config, client *observabilityadmin.Client) {
	input := &observabilityadmin.UpdateTelemetryRuleForOrganizationInput{
		// Rule: *types.TelemetryRule, // Required
		// RuleIdentifier: *string, // Required
	}

	if len(_observabilityadminRule) > 0 {
		if err := assignInputField(input, "Rule", _observabilityadminRule); err != nil {
			log.Errorf("invalid --rule: %s", err.Error())
			return
		}
	}
	if len(_observabilityadminRuleIdentifier) > 0 {
		input.RuleIdentifier = aws.String(_observabilityadminRuleIdentifier)
	}

	if resp, err := client.UpdateTelemetryRuleForOrganization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Validates a pipeline configuration without creating the pipeline. This
// operation checks the configuration for syntax errors and compatibility issues.
func observabilityadmin_ValidateTelemetryPipelineConfiguration(cfg aws.Config, client *observabilityadmin.Client) {
	input := &observabilityadmin.ValidateTelemetryPipelineConfigurationInput{
		// Configuration: *types.TelemetryPipelineConfiguration, // Required
	}

	if len(_observabilityadminConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _observabilityadminConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.ValidateTelemetryPipelineConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_observabilityadminCmd)
	_observabilityadminCmd.Flags().SortFlags = false

	_observabilityadminCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_observabilityadminCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_observabilityadminCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_observabilityadminCmd.Flags().StringSliceVarP(&_observabilityadminAccountIdentifiers, "account-identifiers", "", nil, "Account Identifiers")
	_observabilityadminCmd.Flags().StringVarP(&_observabilityadminAllRegions, "all-regions", "", "", "All Regions")
	_observabilityadminCmd.Flags().StringVarP(&_observabilityadminArn, "arn", "", "", "ARN")
	_observabilityadminCmd.Flags().StringVarP(&_observabilityadminConfiguration, "configuration", "", "", "Configuration")
	_observabilityadminCmd.Flags().StringVarP(&_observabilityadminEncryption, "encryption", "", "", "Encryption")
	_observabilityadminCmd.Flags().StringVarP(&_observabilityadminMaxResults, "max-results", "", "", "Max Results")
	_observabilityadminCmd.Flags().StringVarP(&_observabilityadminName, "name", "", "", "Name")
	_observabilityadminCmd.Flags().StringVarP(&_observabilityadminNextToken, "next-token", "", "", "Next Token")
	_observabilityadminCmd.Flags().StringVarP(&_observabilityadminPipelineIdentifier, "pipeline-identifier", "", "", "Pipeline Identifier")
	_observabilityadminCmd.Flags().StringVarP(&_observabilityadminRecords, "records", "", "", "Records")
	_observabilityadminCmd.Flags().StringVarP(&_observabilityadminResourceARN, "resource-arn", "", "", "Resource ARN")
	_observabilityadminCmd.Flags().StringVarP(&_observabilityadminResourceIdentifierPrefix, "resource-identifier-prefix", "", "", "Resource Identifier Prefix")
	_observabilityadminCmd.Flags().StringVarP(&_observabilityadminResourceTags, "resource-tags", "", "", "Resource Tags")
	_observabilityadminCmd.Flags().StringVarP(&_observabilityadminResourceTypes, "resource-types", "", "", "Resource Types")
	_observabilityadminCmd.Flags().StringVarP(&_observabilityadminRoleArn, "role-arn", "", "", "Role ARN")
	_observabilityadminCmd.Flags().StringVarP(&_observabilityadminRule, "rule", "", "", "Rule")
	_observabilityadminCmd.Flags().StringVarP(&_observabilityadminRuleIdentifier, "rule-identifier", "", "", "Rule Identifier")
	_observabilityadminCmd.Flags().StringVarP(&_observabilityadminRuleName, "rule-name", "", "", "Rule Name")
	_observabilityadminCmd.Flags().StringVarP(&_observabilityadminRuleNamePrefix, "rule-name-prefix", "", "", "Rule Name Prefix")
	_observabilityadminCmd.Flags().StringSliceVarP(&_observabilityadminSourceAccountIds, "source-account-ids", "", nil, "Source Account Ids")
	_observabilityadminCmd.Flags().StringSliceVarP(&_observabilityadminSourceOrganizationUnitIds, "source-organization-unit-ids", "", nil, "Source Organization Unit Ids")
	_observabilityadminCmd.Flags().StringSliceVarP(&_observabilityadminTagKeys, "tag-keys", "", nil, "Tag Keys")
	_observabilityadminCmd.Flags().StringVarP(&_observabilityadminTags, "tags", "", "", "Tags")
	_observabilityadminCmd.Flags().StringVarP(&_observabilityadminTelemetryConfigurationState, "telemetry-configuration-state", "", "", "Telemetry Configuration State")

	_observabilityadminCmd.Flags().BoolVarP(&_observabilityadminCreateCentralizationRuleForOrganization, "create-centralization-rule-for-organization", "", false, "Create Centralization Rule For Organization")
	_observabilityadminCmd.Flags().BoolVarP(&_observabilityadminCreateS3TableIntegration, "create-s3-table-integration", "", false, "Create S3 Table Integration")
	_observabilityadminCmd.Flags().BoolVarP(&_observabilityadminCreateTelemetryPipeline, "create-telemetry-pipeline", "", false, "Create Telemetry Pipeline")
	_observabilityadminCmd.Flags().BoolVarP(&_observabilityadminCreateTelemetryRule, "create-telemetry-rule", "", false, "Create Telemetry Rule")
	_observabilityadminCmd.Flags().BoolVarP(&_observabilityadminCreateTelemetryRuleForOrganization, "create-telemetry-rule-for-organization", "", false, "Create Telemetry Rule For Organization")
	_observabilityadminCmd.Flags().BoolVarP(&_observabilityadminDeleteCentralizationRuleForOrganization, "delete-centralization-rule-for-organization", "", false, "Delete Centralization Rule For Organization")
	_observabilityadminCmd.Flags().BoolVarP(&_observabilityadminDeleteS3TableIntegration, "delete-s3-table-integration", "", false, "Delete S3 Table Integration")
	_observabilityadminCmd.Flags().BoolVarP(&_observabilityadminDeleteTelemetryPipeline, "delete-telemetry-pipeline", "", false, "Delete Telemetry Pipeline")
	_observabilityadminCmd.Flags().BoolVarP(&_observabilityadminDeleteTelemetryRule, "delete-telemetry-rule", "", false, "Delete Telemetry Rule")
	_observabilityadminCmd.Flags().BoolVarP(&_observabilityadminDeleteTelemetryRuleForOrganization, "delete-telemetry-rule-for-organization", "", false, "Delete Telemetry Rule For Organization")
	_observabilityadminCmd.Flags().BoolVarP(&_observabilityadminGetCentralizationRuleForOrganization, "get-centralization-rule-for-organization", "", false, "Get Centralization Rule For Organization")
	_observabilityadminCmd.Flags().BoolVarP(&_observabilityadminGetS3TableIntegration, "get-s3-table-integration", "", false, "Get S3 Table Integration")
	_observabilityadminCmd.Flags().BoolVarP(&_observabilityadminGetTelemetryEnrichmentStatus, "get-telemetry-enrichment-status", "", false, "Get Telemetry Enrichment Status")
	_observabilityadminCmd.Flags().BoolVarP(&_observabilityadminGetTelemetryEvaluationStatus, "get-telemetry-evaluation-status", "", false, "Get Telemetry Evaluation Status")
	_observabilityadminCmd.Flags().BoolVarP(&_observabilityadminGetTelemetryEvaluationStatusForOrganization, "get-telemetry-evaluation-status-for-organization", "", false, "Get Telemetry Evaluation Status For Organization")
	_observabilityadminCmd.Flags().BoolVarP(&_observabilityadminGetTelemetryPipeline, "get-telemetry-pipeline", "", false, "Get Telemetry Pipeline")
	_observabilityadminCmd.Flags().BoolVarP(&_observabilityadminGetTelemetryRule, "get-telemetry-rule", "", false, "Get Telemetry Rule")
	_observabilityadminCmd.Flags().BoolVarP(&_observabilityadminGetTelemetryRuleForOrganization, "get-telemetry-rule-for-organization", "", false, "Get Telemetry Rule For Organization")
	_observabilityadminCmd.Flags().BoolVarP(&_observabilityadminListCentralizationRulesForOrganization, "list-centralization-rules-for-organization", "", false, "List Centralization Rules For Organization")
	_observabilityadminCmd.Flags().BoolVarP(&_observabilityadminListResourceTelemetry, "list-resource-telemetry", "", false, "List Resource Telemetry")
	_observabilityadminCmd.Flags().BoolVarP(&_observabilityadminListResourceTelemetryForOrganization, "list-resource-telemetry-for-organization", "", false, "List Resource Telemetry For Organization")
	_observabilityadminCmd.Flags().BoolVarP(&_observabilityadminListS3TableIntegrations, "list-s3-table-integrations", "", false, "List S3 Table Integrations")
	_observabilityadminCmd.Flags().BoolVarP(&_observabilityadminListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_observabilityadminCmd.Flags().BoolVarP(&_observabilityadminListTelemetryPipelines, "list-telemetry-pipelines", "", false, "List Telemetry Pipelines")
	_observabilityadminCmd.Flags().BoolVarP(&_observabilityadminListTelemetryRules, "list-telemetry-rules", "", false, "List Telemetry Rules")
	_observabilityadminCmd.Flags().BoolVarP(&_observabilityadminListTelemetryRulesForOrganization, "list-telemetry-rules-for-organization", "", false, "List Telemetry Rules For Organization")
	_observabilityadminCmd.Flags().BoolVarP(&_observabilityadminStartTelemetryEnrichment, "start-telemetry-enrichment", "", false, "Start Telemetry Enrichment")
	_observabilityadminCmd.Flags().BoolVarP(&_observabilityadminStartTelemetryEvaluation, "start-telemetry-evaluation", "", false, "Start Telemetry Evaluation")
	_observabilityadminCmd.Flags().BoolVarP(&_observabilityadminStartTelemetryEvaluationForOrganization, "start-telemetry-evaluation-for-organization", "", false, "Start Telemetry Evaluation For Organization")
	_observabilityadminCmd.Flags().BoolVarP(&_observabilityadminStopTelemetryEnrichment, "stop-telemetry-enrichment", "", false, "Stop Telemetry Enrichment")
	_observabilityadminCmd.Flags().BoolVarP(&_observabilityadminStopTelemetryEvaluation, "stop-telemetry-evaluation", "", false, "Stop Telemetry Evaluation")
	_observabilityadminCmd.Flags().BoolVarP(&_observabilityadminStopTelemetryEvaluationForOrganization, "stop-telemetry-evaluation-for-organization", "", false, "Stop Telemetry Evaluation For Organization")
	_observabilityadminCmd.Flags().BoolVarP(&_observabilityadminTagResource, "tag-resource", "", false, "Tag Resource")
	_observabilityadminCmd.Flags().BoolVarP(&_observabilityadminTestTelemetryPipeline, "test-telemetry-pipeline", "", false, "Test Telemetry Pipeline")
	_observabilityadminCmd.Flags().BoolVarP(&_observabilityadminUntagResource, "untag-resource", "", false, "Untag Resource")
	_observabilityadminCmd.Flags().BoolVarP(&_observabilityadminUpdateCentralizationRuleForOrganization, "update-centralization-rule-for-organization", "", false, "Update Centralization Rule For Organization")
	_observabilityadminCmd.Flags().BoolVarP(&_observabilityadminUpdateTelemetryPipeline, "update-telemetry-pipeline", "", false, "Update Telemetry Pipeline")
	_observabilityadminCmd.Flags().BoolVarP(&_observabilityadminUpdateTelemetryRule, "update-telemetry-rule", "", false, "Update Telemetry Rule")
	_observabilityadminCmd.Flags().BoolVarP(&_observabilityadminUpdateTelemetryRuleForOrganization, "update-telemetry-rule-for-organization", "", false, "Update Telemetry Rule For Organization")
	_observabilityadminCmd.Flags().BoolVarP(&_observabilityadminValidateTelemetryPipelineConfiguration, "validate-telemetry-pipeline-configuration", "", false, "Validate Telemetry Pipeline Configuration")

}
