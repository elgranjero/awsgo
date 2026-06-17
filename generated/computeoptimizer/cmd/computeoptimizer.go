package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// computeoptimizerCmd represents the computeoptimizer command
var _computeoptimizerCmd = &cobra.Command{
	Use:   "computeoptimizer",
	Short: "AWS computeoptimizer CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := computeoptimizer.NewFromConfig(cfg)
		if _computeoptimizerDeleteRecommendationPreferences {
			computeoptimizer_DeleteRecommendationPreferences(cfg, client)
			return
		}
		if _computeoptimizerDescribeRecommendationExportJobs {
			computeoptimizer_DescribeRecommendationExportJobs(cfg, client)
			return
		}
		if _computeoptimizerExportAutoScalingGroupRecommendations {
			computeoptimizer_ExportAutoScalingGroupRecommendations(cfg, client)
			return
		}
		if _computeoptimizerExportEBSVolumeRecommendations {
			computeoptimizer_ExportEBSVolumeRecommendations(cfg, client)
			return
		}
		if _computeoptimizerExportEC2InstanceRecommendations {
			computeoptimizer_ExportEC2InstanceRecommendations(cfg, client)
			return
		}
		if _computeoptimizerExportECSServiceRecommendations {
			computeoptimizer_ExportECSServiceRecommendations(cfg, client)
			return
		}
		if _computeoptimizerExportIdleRecommendations {
			computeoptimizer_ExportIdleRecommendations(cfg, client)
			return
		}
		if _computeoptimizerExportLambdaFunctionRecommendations {
			computeoptimizer_ExportLambdaFunctionRecommendations(cfg, client)
			return
		}
		if _computeoptimizerExportLicenseRecommendations {
			computeoptimizer_ExportLicenseRecommendations(cfg, client)
			return
		}
		if _computeoptimizerExportRDSDatabaseRecommendations {
			computeoptimizer_ExportRDSDatabaseRecommendations(cfg, client)
			return
		}
		if _computeoptimizerGetAutoScalingGroupRecommendations {
			computeoptimizer_GetAutoScalingGroupRecommendations(cfg, client)
			return
		}
		if _computeoptimizerGetEBSVolumeRecommendations {
			computeoptimizer_GetEBSVolumeRecommendations(cfg, client)
			return
		}
		if _computeoptimizerGetEC2InstanceRecommendations {
			computeoptimizer_GetEC2InstanceRecommendations(cfg, client)
			return
		}
		if _computeoptimizerGetEC2RecommendationProjectedMetrics {
			computeoptimizer_GetEC2RecommendationProjectedMetrics(cfg, client)
			return
		}
		if _computeoptimizerGetECSServiceRecommendationProjectedMetrics {
			computeoptimizer_GetECSServiceRecommendationProjectedMetrics(cfg, client)
			return
		}
		if _computeoptimizerGetECSServiceRecommendations {
			computeoptimizer_GetECSServiceRecommendations(cfg, client)
			return
		}
		if _computeoptimizerGetEffectiveRecommendationPreferences {
			computeoptimizer_GetEffectiveRecommendationPreferences(cfg, client)
			return
		}
		if _computeoptimizerGetEnrollmentStatus {
			computeoptimizer_GetEnrollmentStatus(cfg, client)
			return
		}
		if _computeoptimizerGetEnrollmentStatusesForOrganization {
			computeoptimizer_GetEnrollmentStatusesForOrganization(cfg, client)
			return
		}
		if _computeoptimizerGetIdleRecommendations {
			computeoptimizer_GetIdleRecommendations(cfg, client)
			return
		}
		if _computeoptimizerGetLambdaFunctionRecommendations {
			computeoptimizer_GetLambdaFunctionRecommendations(cfg, client)
			return
		}
		if _computeoptimizerGetLicenseRecommendations {
			computeoptimizer_GetLicenseRecommendations(cfg, client)
			return
		}
		if _computeoptimizerGetRDSDatabaseRecommendationProjectedMetrics {
			computeoptimizer_GetRDSDatabaseRecommendationProjectedMetrics(cfg, client)
			return
		}
		if _computeoptimizerGetRDSDatabaseRecommendations {
			computeoptimizer_GetRDSDatabaseRecommendations(cfg, client)
			return
		}
		if _computeoptimizerGetRecommendationPreferences {
			computeoptimizer_GetRecommendationPreferences(cfg, client)
			return
		}
		if _computeoptimizerGetRecommendationSummaries {
			computeoptimizer_GetRecommendationSummaries(cfg, client)
			return
		}
		if _computeoptimizerPutRecommendationPreferences {
			computeoptimizer_PutRecommendationPreferences(cfg, client)
			return
		}
		if _computeoptimizerUpdateEnrollmentStatus {
			computeoptimizer_UpdateEnrollmentStatus(cfg, client)
			return
		}

	},
}

var (
	_computeoptimizerDeleteRecommendationPreferences              bool
	_computeoptimizerDescribeRecommendationExportJobs             bool
	_computeoptimizerExportAutoScalingGroupRecommendations        bool
	_computeoptimizerExportEBSVolumeRecommendations               bool
	_computeoptimizerExportEC2InstanceRecommendations             bool
	_computeoptimizerExportECSServiceRecommendations              bool
	_computeoptimizerExportIdleRecommendations                    bool
	_computeoptimizerExportLambdaFunctionRecommendations          bool
	_computeoptimizerExportLicenseRecommendations                 bool
	_computeoptimizerExportRDSDatabaseRecommendations             bool
	_computeoptimizerGetAutoScalingGroupRecommendations           bool
	_computeoptimizerGetEBSVolumeRecommendations                  bool
	_computeoptimizerGetEC2InstanceRecommendations                bool
	_computeoptimizerGetEC2RecommendationProjectedMetrics         bool
	_computeoptimizerGetECSServiceRecommendationProjectedMetrics  bool
	_computeoptimizerGetECSServiceRecommendations                 bool
	_computeoptimizerGetEffectiveRecommendationPreferences        bool
	_computeoptimizerGetEnrollmentStatus                          bool
	_computeoptimizerGetEnrollmentStatusesForOrganization         bool
	_computeoptimizerGetIdleRecommendations                       bool
	_computeoptimizerGetLambdaFunctionRecommendations             bool
	_computeoptimizerGetLicenseRecommendations                    bool
	_computeoptimizerGetRDSDatabaseRecommendationProjectedMetrics bool
	_computeoptimizerGetRDSDatabaseRecommendations                bool
	_computeoptimizerGetRecommendationPreferences                 bool
	_computeoptimizerGetRecommendationSummaries                   bool
	_computeoptimizerPutRecommendationPreferences                 bool
	_computeoptimizerUpdateEnrollmentStatus                       bool

	_computeoptimizerAccountIds                    []string
	_computeoptimizerAutoScalingGroupArns          []string
	_computeoptimizerEndTime                       string
	_computeoptimizerEnhancedInfrastructureMetrics string
	_computeoptimizerExternalMetricsPreference     string
	_computeoptimizerFieldsToExport                string
	_computeoptimizerFileFormat                    string
	_computeoptimizerFilters                       string
	_computeoptimizerFunctionArns                  []string
	_computeoptimizerIncludeMemberAccounts         string
	_computeoptimizerInferredWorkloadTypes         string
	_computeoptimizerInstanceArn                   string
	_computeoptimizerInstanceArns                  []string
	_computeoptimizerJobIds                        []string
	_computeoptimizerLookBackPeriod                string
	_computeoptimizerMaxResults                    string
	_computeoptimizerNextToken                     string
	_computeoptimizerOrderBy                       string
	_computeoptimizerPeriod                        string
	_computeoptimizerPreferredResources            string
	_computeoptimizerRecommendationPreferenceNames string
	_computeoptimizerRecommendationPreferences     string
	_computeoptimizerResourceArn                   string
	_computeoptimizerResourceArns                  []string
	_computeoptimizerResourceType                  string
	_computeoptimizerS3DestinationConfig           string
	_computeoptimizerSavingsEstimationMode         string
	_computeoptimizerScope                         string
	_computeoptimizerServiceArn                    string
	_computeoptimizerServiceArns                   []string
	_computeoptimizerStartTime                     string
	_computeoptimizerStat                          string
	_computeoptimizerStatus                        string
	_computeoptimizerUtilizationPreferences        string
	_computeoptimizerVolumeArns                    []string
)

// Deletes a recommendation preference, such as enhanced infrastructure metrics.
// For more information, see [Activating enhanced infrastructure metrics] in the Compute Optimizer User Guide.
//
// [Activating enhanced infrastructure metrics]: https://docs.aws.amazon.com/compute-optimizer/latest/ug/enhanced-infrastructure-metrics.html
func computeoptimizer_DeleteRecommendationPreferences(cfg aws.Config, client *computeoptimizer.Client) {
	input := &computeoptimizer.DeleteRecommendationPreferencesInput{
		// RecommendationPreferenceNames: []types.RecommendationPreferenceName, // Required
		// ResourceType: types.ResourceType, // Required
	}

	if len(_computeoptimizerRecommendationPreferenceNames) > 0 {
		if err := assignInputField(input, "RecommendationPreferenceNames", _computeoptimizerRecommendationPreferenceNames); err != nil {
			log.Errorf("invalid --recommendation-preference-names: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _computeoptimizerResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerScope) > 0 {
		if err := assignInputField(input, "Scope", _computeoptimizerScope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteRecommendationPreferences(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes recommendation export jobs created in the last seven days.
// Use the ExportAutoScalingGroupRecommendations or ExportEC2InstanceRecommendations actions to request an export of your recommendations. Then use the DescribeRecommendationExportJobs
// action to view your export jobs.
func computeoptimizer_DescribeRecommendationExportJobs(cfg aws.Config, client *computeoptimizer.Client) {
	input := &computeoptimizer.DescribeRecommendationExportJobsInput{}

	if len(_computeoptimizerFilters) > 0 {
		if err := assignInputField(input, "Filters", _computeoptimizerFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerJobIds) > 0 {
		input.JobIds = append([]string(nil), _computeoptimizerJobIds...)
	}
	if len(_computeoptimizerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _computeoptimizerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerNextToken) > 0 {
		input.NextToken = aws.String(_computeoptimizerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeRecommendationExportJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*computeoptimizer.DescribeRecommendationExportJobsOutput
	p := computeoptimizer.NewDescribeRecommendationExportJobsPaginator(client, input)
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

// Exports optimization recommendations for Amazon EC2 Auto Scaling groups.
// Recommendations are exported in a comma-separated values (.csv) file, and its
// metadata in a JavaScript Object Notation (JSON) (.json) file, to an existing
// Amazon Simple Storage Service (Amazon S3) bucket that you specify. For more
// information, see [Exporting Recommendations]in the Compute Optimizer User Guide.
//
// You can have only one Amazon EC2 Auto Scaling group export job in progress per
// Amazon Web Services Region.
//
// [Exporting Recommendations]: https://docs.aws.amazon.com/compute-optimizer/latest/ug/exporting-recommendations.html
func computeoptimizer_ExportAutoScalingGroupRecommendations(cfg aws.Config, client *computeoptimizer.Client) {
	input := &computeoptimizer.ExportAutoScalingGroupRecommendationsInput{
		// S3DestinationConfig: *types.S3DestinationConfig, // Required
	}

	if len(_computeoptimizerS3DestinationConfig) > 0 {
		if err := assignInputField(input, "S3DestinationConfig", _computeoptimizerS3DestinationConfig); err != nil {
			log.Errorf("invalid --s3-destination-config: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _computeoptimizerAccountIds...)
	}
	if len(_computeoptimizerFieldsToExport) > 0 {
		if err := assignInputField(input, "FieldsToExport", _computeoptimizerFieldsToExport); err != nil {
			log.Errorf("invalid --fields-to-export: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerFileFormat) > 0 {
		if err := assignInputField(input, "FileFormat", _computeoptimizerFileFormat); err != nil {
			log.Errorf("invalid --file-format: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerFilters) > 0 {
		if err := assignInputField(input, "Filters", _computeoptimizerFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerIncludeMemberAccounts) > 0 {
		if err := assignInputField(input, "IncludeMemberAccounts", _computeoptimizerIncludeMemberAccounts); err != nil {
			log.Errorf("invalid --include-member-accounts: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerRecommendationPreferences) > 0 {
		if err := assignInputField(input, "RecommendationPreferences", _computeoptimizerRecommendationPreferences); err != nil {
			log.Errorf("invalid --recommendation-preferences: %s", err.Error())
			return
		}
	}

	if resp, err := client.ExportAutoScalingGroupRecommendations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Exports optimization recommendations for Amazon EBS volumes.
// Recommendations are exported in a comma-separated values (.csv) file, and its
// metadata in a JavaScript Object Notation (JSON) (.json) file, to an existing
// Amazon Simple Storage Service (Amazon S3) bucket that you specify. For more
// information, see [Exporting Recommendations]in the Compute Optimizer User Guide.
//
// You can have only one Amazon EBS volume export job in progress per Amazon Web
// Services Region.
//
// [Exporting Recommendations]: https://docs.aws.amazon.com/compute-optimizer/latest/ug/exporting-recommendations.html
func computeoptimizer_ExportEBSVolumeRecommendations(cfg aws.Config, client *computeoptimizer.Client) {
	input := &computeoptimizer.ExportEBSVolumeRecommendationsInput{
		// S3DestinationConfig: *types.S3DestinationConfig, // Required
	}

	if len(_computeoptimizerS3DestinationConfig) > 0 {
		if err := assignInputField(input, "S3DestinationConfig", _computeoptimizerS3DestinationConfig); err != nil {
			log.Errorf("invalid --s3-destination-config: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _computeoptimizerAccountIds...)
	}
	if len(_computeoptimizerFieldsToExport) > 0 {
		if err := assignInputField(input, "FieldsToExport", _computeoptimizerFieldsToExport); err != nil {
			log.Errorf("invalid --fields-to-export: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerFileFormat) > 0 {
		if err := assignInputField(input, "FileFormat", _computeoptimizerFileFormat); err != nil {
			log.Errorf("invalid --file-format: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerFilters) > 0 {
		if err := assignInputField(input, "Filters", _computeoptimizerFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerIncludeMemberAccounts) > 0 {
		if err := assignInputField(input, "IncludeMemberAccounts", _computeoptimizerIncludeMemberAccounts); err != nil {
			log.Errorf("invalid --include-member-accounts: %s", err.Error())
			return
		}
	}

	if resp, err := client.ExportEBSVolumeRecommendations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Exports optimization recommendations for Amazon EC2 instances.
// Recommendations are exported in a comma-separated values (.csv) file, and its
// metadata in a JavaScript Object Notation (JSON) (.json) file, to an existing
// Amazon Simple Storage Service (Amazon S3) bucket that you specify. For more
// information, see [Exporting Recommendations]in the Compute Optimizer User Guide.
//
// You can have only one Amazon EC2 instance export job in progress per Amazon Web
// Services Region.
//
// [Exporting Recommendations]: https://docs.aws.amazon.com/compute-optimizer/latest/ug/exporting-recommendations.html
func computeoptimizer_ExportEC2InstanceRecommendations(cfg aws.Config, client *computeoptimizer.Client) {
	input := &computeoptimizer.ExportEC2InstanceRecommendationsInput{
		// S3DestinationConfig: *types.S3DestinationConfig, // Required
	}

	if len(_computeoptimizerS3DestinationConfig) > 0 {
		if err := assignInputField(input, "S3DestinationConfig", _computeoptimizerS3DestinationConfig); err != nil {
			log.Errorf("invalid --s3-destination-config: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _computeoptimizerAccountIds...)
	}
	if len(_computeoptimizerFieldsToExport) > 0 {
		if err := assignInputField(input, "FieldsToExport", _computeoptimizerFieldsToExport); err != nil {
			log.Errorf("invalid --fields-to-export: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerFileFormat) > 0 {
		if err := assignInputField(input, "FileFormat", _computeoptimizerFileFormat); err != nil {
			log.Errorf("invalid --file-format: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerFilters) > 0 {
		if err := assignInputField(input, "Filters", _computeoptimizerFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerIncludeMemberAccounts) > 0 {
		if err := assignInputField(input, "IncludeMemberAccounts", _computeoptimizerIncludeMemberAccounts); err != nil {
			log.Errorf("invalid --include-member-accounts: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerRecommendationPreferences) > 0 {
		if err := assignInputField(input, "RecommendationPreferences", _computeoptimizerRecommendationPreferences); err != nil {
			log.Errorf("invalid --recommendation-preferences: %s", err.Error())
			return
		}
	}

	if resp, err := client.ExportEC2InstanceRecommendations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Exports optimization recommendations for Amazon ECS services on Fargate.
// Recommendations are exported in a CSV file, and its metadata in a JSON file, to
// an existing Amazon Simple Storage Service (Amazon S3) bucket that you specify.
// For more information, see [Exporting Recommendations]in the Compute Optimizer User Guide.
//
// You can only have one Amazon ECS service export job in progress per Amazon Web
// Services Region.
//
// [Exporting Recommendations]: https://docs.aws.amazon.com/compute-optimizer/latest/ug/exporting-recommendations.html
func computeoptimizer_ExportECSServiceRecommendations(cfg aws.Config, client *computeoptimizer.Client) {
	input := &computeoptimizer.ExportECSServiceRecommendationsInput{
		// S3DestinationConfig: *types.S3DestinationConfig, // Required
	}

	if len(_computeoptimizerS3DestinationConfig) > 0 {
		if err := assignInputField(input, "S3DestinationConfig", _computeoptimizerS3DestinationConfig); err != nil {
			log.Errorf("invalid --s3-destination-config: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _computeoptimizerAccountIds...)
	}
	if len(_computeoptimizerFieldsToExport) > 0 {
		if err := assignInputField(input, "FieldsToExport", _computeoptimizerFieldsToExport); err != nil {
			log.Errorf("invalid --fields-to-export: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerFileFormat) > 0 {
		if err := assignInputField(input, "FileFormat", _computeoptimizerFileFormat); err != nil {
			log.Errorf("invalid --file-format: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerFilters) > 0 {
		if err := assignInputField(input, "Filters", _computeoptimizerFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerIncludeMemberAccounts) > 0 {
		if err := assignInputField(input, "IncludeMemberAccounts", _computeoptimizerIncludeMemberAccounts); err != nil {
			log.Errorf("invalid --include-member-accounts: %s", err.Error())
			return
		}
	}

	if resp, err := client.ExportECSServiceRecommendations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Export optimization recommendations for your idle resources.
// Recommendations are exported in a comma-separated values (CSV) file, and its
// metadata in a JavaScript Object Notation (JSON) file, to an existing Amazon
// Simple Storage Service (Amazon S3) bucket that you specify. For more
// information, see [Exporting Recommendations]in the Compute Optimizer User Guide.
//
// You can have only one idle resource export job in progress per Amazon Web
// Services Region.
//
// [Exporting Recommendations]: https://docs.aws.amazon.com/compute-optimizer/latest/ug/exporting-recommendations.html
func computeoptimizer_ExportIdleRecommendations(cfg aws.Config, client *computeoptimizer.Client) {
	input := &computeoptimizer.ExportIdleRecommendationsInput{
		// S3DestinationConfig: *types.S3DestinationConfig, // Required
	}

	if len(_computeoptimizerS3DestinationConfig) > 0 {
		if err := assignInputField(input, "S3DestinationConfig", _computeoptimizerS3DestinationConfig); err != nil {
			log.Errorf("invalid --s3-destination-config: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _computeoptimizerAccountIds...)
	}
	if len(_computeoptimizerFieldsToExport) > 0 {
		if err := assignInputField(input, "FieldsToExport", _computeoptimizerFieldsToExport); err != nil {
			log.Errorf("invalid --fields-to-export: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerFileFormat) > 0 {
		if err := assignInputField(input, "FileFormat", _computeoptimizerFileFormat); err != nil {
			log.Errorf("invalid --file-format: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerFilters) > 0 {
		if err := assignInputField(input, "Filters", _computeoptimizerFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerIncludeMemberAccounts) > 0 {
		if err := assignInputField(input, "IncludeMemberAccounts", _computeoptimizerIncludeMemberAccounts); err != nil {
			log.Errorf("invalid --include-member-accounts: %s", err.Error())
			return
		}
	}

	if resp, err := client.ExportIdleRecommendations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Exports optimization recommendations for Lambda functions.
// Recommendations are exported in a comma-separated values (.csv) file, and its
// metadata in a JavaScript Object Notation (JSON) (.json) file, to an existing
// Amazon Simple Storage Service (Amazon S3) bucket that you specify. For more
// information, see [Exporting Recommendations]in the Compute Optimizer User Guide.
//
// You can have only one Lambda function export job in progress per Amazon Web
// Services Region.
//
// [Exporting Recommendations]: https://docs.aws.amazon.com/compute-optimizer/latest/ug/exporting-recommendations.html
func computeoptimizer_ExportLambdaFunctionRecommendations(cfg aws.Config, client *computeoptimizer.Client) {
	input := &computeoptimizer.ExportLambdaFunctionRecommendationsInput{
		// S3DestinationConfig: *types.S3DestinationConfig, // Required
	}

	if len(_computeoptimizerS3DestinationConfig) > 0 {
		if err := assignInputField(input, "S3DestinationConfig", _computeoptimizerS3DestinationConfig); err != nil {
			log.Errorf("invalid --s3-destination-config: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _computeoptimizerAccountIds...)
	}
	if len(_computeoptimizerFieldsToExport) > 0 {
		if err := assignInputField(input, "FieldsToExport", _computeoptimizerFieldsToExport); err != nil {
			log.Errorf("invalid --fields-to-export: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerFileFormat) > 0 {
		if err := assignInputField(input, "FileFormat", _computeoptimizerFileFormat); err != nil {
			log.Errorf("invalid --file-format: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerFilters) > 0 {
		if err := assignInputField(input, "Filters", _computeoptimizerFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerIncludeMemberAccounts) > 0 {
		if err := assignInputField(input, "IncludeMemberAccounts", _computeoptimizerIncludeMemberAccounts); err != nil {
			log.Errorf("invalid --include-member-accounts: %s", err.Error())
			return
		}
	}

	if resp, err := client.ExportLambdaFunctionRecommendations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Export optimization recommendations for your licenses.
// Recommendations are exported in a comma-separated values (CSV) file, and its
// metadata in a JavaScript Object Notation (JSON) file, to an existing Amazon
// Simple Storage Service (Amazon S3) bucket that you specify. For more
// information, see [Exporting Recommendations]in the Compute Optimizer User Guide.
//
// You can have only one license export job in progress per Amazon Web Services
// Region.
//
// [Exporting Recommendations]: https://docs.aws.amazon.com/compute-optimizer/latest/ug/exporting-recommendations.html
func computeoptimizer_ExportLicenseRecommendations(cfg aws.Config, client *computeoptimizer.Client) {
	input := &computeoptimizer.ExportLicenseRecommendationsInput{
		// S3DestinationConfig: *types.S3DestinationConfig, // Required
	}

	if len(_computeoptimizerS3DestinationConfig) > 0 {
		if err := assignInputField(input, "S3DestinationConfig", _computeoptimizerS3DestinationConfig); err != nil {
			log.Errorf("invalid --s3-destination-config: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _computeoptimizerAccountIds...)
	}
	if len(_computeoptimizerFieldsToExport) > 0 {
		if err := assignInputField(input, "FieldsToExport", _computeoptimizerFieldsToExport); err != nil {
			log.Errorf("invalid --fields-to-export: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerFileFormat) > 0 {
		if err := assignInputField(input, "FileFormat", _computeoptimizerFileFormat); err != nil {
			log.Errorf("invalid --file-format: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerFilters) > 0 {
		if err := assignInputField(input, "Filters", _computeoptimizerFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerIncludeMemberAccounts) > 0 {
		if err := assignInputField(input, "IncludeMemberAccounts", _computeoptimizerIncludeMemberAccounts); err != nil {
			log.Errorf("invalid --include-member-accounts: %s", err.Error())
			return
		}
	}

	if resp, err := client.ExportLicenseRecommendations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Export optimization recommendations for your Amazon Aurora and Amazon
// Relational Database Service (Amazon RDS) databases.
//
// Recommendations are exported in a comma-separated values (CSV) file, and its
// metadata in a JavaScript Object Notation (JSON) file, to an existing Amazon
// Simple Storage Service (Amazon S3) bucket that you specify. For more
// information, see [Exporting Recommendations]in the Compute Optimizer User Guide.
//
// You can have only one Amazon Aurora or RDS export job in progress per Amazon
// Web Services Region.
//
// [Exporting Recommendations]: https://docs.aws.amazon.com/compute-optimizer/latest/ug/exporting-recommendations.html
func computeoptimizer_ExportRDSDatabaseRecommendations(cfg aws.Config, client *computeoptimizer.Client) {
	input := &computeoptimizer.ExportRDSDatabaseRecommendationsInput{
		// S3DestinationConfig: *types.S3DestinationConfig, // Required
	}

	if len(_computeoptimizerS3DestinationConfig) > 0 {
		if err := assignInputField(input, "S3DestinationConfig", _computeoptimizerS3DestinationConfig); err != nil {
			log.Errorf("invalid --s3-destination-config: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _computeoptimizerAccountIds...)
	}
	if len(_computeoptimizerFieldsToExport) > 0 {
		if err := assignInputField(input, "FieldsToExport", _computeoptimizerFieldsToExport); err != nil {
			log.Errorf("invalid --fields-to-export: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerFileFormat) > 0 {
		if err := assignInputField(input, "FileFormat", _computeoptimizerFileFormat); err != nil {
			log.Errorf("invalid --file-format: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerFilters) > 0 {
		if err := assignInputField(input, "Filters", _computeoptimizerFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerIncludeMemberAccounts) > 0 {
		if err := assignInputField(input, "IncludeMemberAccounts", _computeoptimizerIncludeMemberAccounts); err != nil {
			log.Errorf("invalid --include-member-accounts: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerRecommendationPreferences) > 0 {
		if err := assignInputField(input, "RecommendationPreferences", _computeoptimizerRecommendationPreferences); err != nil {
			log.Errorf("invalid --recommendation-preferences: %s", err.Error())
			return
		}
	}

	if resp, err := client.ExportRDSDatabaseRecommendations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns Amazon EC2 Auto Scaling group recommendations.
// Compute Optimizer generates recommendations for Amazon EC2 Auto Scaling groups
// that meet a specific set of requirements. For more information, see the [Supported resources and requirements]in the
// Compute Optimizer User Guide.
//
// [Supported resources and requirements]: https://docs.aws.amazon.com/compute-optimizer/latest/ug/requirements.html
func computeoptimizer_GetAutoScalingGroupRecommendations(cfg aws.Config, client *computeoptimizer.Client) {
	input := &computeoptimizer.GetAutoScalingGroupRecommendationsInput{}

	if len(_computeoptimizerAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _computeoptimizerAccountIds...)
	}
	if len(_computeoptimizerAutoScalingGroupArns) > 0 {
		input.AutoScalingGroupArns = append([]string(nil), _computeoptimizerAutoScalingGroupArns...)
	}
	if len(_computeoptimizerFilters) > 0 {
		if err := assignInputField(input, "Filters", _computeoptimizerFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _computeoptimizerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerNextToken) > 0 {
		input.NextToken = aws.String(_computeoptimizerNextToken)
	}
	if len(_computeoptimizerRecommendationPreferences) > 0 {
		if err := assignInputField(input, "RecommendationPreferences", _computeoptimizerRecommendationPreferences); err != nil {
			log.Errorf("invalid --recommendation-preferences: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetAutoScalingGroupRecommendations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns Amazon Elastic Block Store (Amazon EBS) volume recommendations.
// Compute Optimizer generates recommendations for Amazon EBS volumes that meet a
// specific set of requirements. For more information, see the [Supported resources and requirements]in the Compute
// Optimizer User Guide.
//
// [Supported resources and requirements]: https://docs.aws.amazon.com/compute-optimizer/latest/ug/requirements.html
func computeoptimizer_GetEBSVolumeRecommendations(cfg aws.Config, client *computeoptimizer.Client) {
	input := &computeoptimizer.GetEBSVolumeRecommendationsInput{}

	if len(_computeoptimizerAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _computeoptimizerAccountIds...)
	}
	if len(_computeoptimizerFilters) > 0 {
		if err := assignInputField(input, "Filters", _computeoptimizerFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _computeoptimizerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerNextToken) > 0 {
		input.NextToken = aws.String(_computeoptimizerNextToken)
	}
	if len(_computeoptimizerVolumeArns) > 0 {
		input.VolumeArns = append([]string(nil), _computeoptimizerVolumeArns...)
	}

	if resp, err := client.GetEBSVolumeRecommendations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns Amazon EC2 instance recommendations.
// Compute Optimizer generates recommendations for Amazon Elastic Compute Cloud
// (Amazon EC2) instances that meet a specific set of requirements. For more
// information, see the [Supported resources and requirements]in the Compute Optimizer User Guide.
//
// [Supported resources and requirements]: https://docs.aws.amazon.com/compute-optimizer/latest/ug/requirements.html
func computeoptimizer_GetEC2InstanceRecommendations(cfg aws.Config, client *computeoptimizer.Client) {
	input := &computeoptimizer.GetEC2InstanceRecommendationsInput{}

	if len(_computeoptimizerAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _computeoptimizerAccountIds...)
	}
	if len(_computeoptimizerFilters) > 0 {
		if err := assignInputField(input, "Filters", _computeoptimizerFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerInstanceArns) > 0 {
		input.InstanceArns = append([]string(nil), _computeoptimizerInstanceArns...)
	}
	if len(_computeoptimizerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _computeoptimizerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerNextToken) > 0 {
		input.NextToken = aws.String(_computeoptimizerNextToken)
	}
	if len(_computeoptimizerRecommendationPreferences) > 0 {
		if err := assignInputField(input, "RecommendationPreferences", _computeoptimizerRecommendationPreferences); err != nil {
			log.Errorf("invalid --recommendation-preferences: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetEC2InstanceRecommendations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the projected utilization metrics of Amazon EC2 instance
// recommendations.
//
// The Cpu and Memory metrics are the only projected utilization metrics returned
// when you run this action. Additionally, the Memory metric is returned only for
// resources that have the unified CloudWatch agent installed on them. For more
// information, see [Enabling Memory Utilization with the CloudWatch Agent].
//
// [Enabling Memory Utilization with the CloudWatch Agent]: https://docs.aws.amazon.com/compute-optimizer/latest/ug/metrics.html#cw-agent
func computeoptimizer_GetEC2RecommendationProjectedMetrics(cfg aws.Config, client *computeoptimizer.Client) {
	input := &computeoptimizer.GetEC2RecommendationProjectedMetricsInput{
		// EndTime: *time.Time, // Required
		// InstanceArn: *string, // Required
		// Period: int32, // Required
		// StartTime: *time.Time, // Required
		// Stat: types.MetricStatistic, // Required
	}

	if len(_computeoptimizerEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _computeoptimizerEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerInstanceArn) > 0 {
		input.InstanceArn = aws.String(_computeoptimizerInstanceArn)
	}
	if len(_computeoptimizerPeriod) > 0 {
		if err := assignInputField(input, "Period", _computeoptimizerPeriod); err != nil {
			log.Errorf("invalid --period: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _computeoptimizerStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerStat) > 0 {
		if err := assignInputField(input, "Stat", _computeoptimizerStat); err != nil {
			log.Errorf("invalid --stat: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerRecommendationPreferences) > 0 {
		if err := assignInputField(input, "RecommendationPreferences", _computeoptimizerRecommendationPreferences); err != nil {
			log.Errorf("invalid --recommendation-preferences: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetEC2RecommendationProjectedMetrics(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the projected metrics of Amazon ECS service recommendations.
func computeoptimizer_GetECSServiceRecommendationProjectedMetrics(cfg aws.Config, client *computeoptimizer.Client) {
	input := &computeoptimizer.GetECSServiceRecommendationProjectedMetricsInput{
		// EndTime: *time.Time, // Required
		// Period: int32, // Required
		// ServiceArn: *string, // Required
		// StartTime: *time.Time, // Required
		// Stat: types.MetricStatistic, // Required
	}

	if len(_computeoptimizerEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _computeoptimizerEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerPeriod) > 0 {
		if err := assignInputField(input, "Period", _computeoptimizerPeriod); err != nil {
			log.Errorf("invalid --period: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerServiceArn) > 0 {
		input.ServiceArn = aws.String(_computeoptimizerServiceArn)
	}
	if len(_computeoptimizerStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _computeoptimizerStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerStat) > 0 {
		if err := assignInputField(input, "Stat", _computeoptimizerStat); err != nil {
			log.Errorf("invalid --stat: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetECSServiceRecommendationProjectedMetrics(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns Amazon ECS service recommendations.
// Compute Optimizer generates recommendations for Amazon ECS services on Fargate
// that meet a specific set of requirements. For more information, see the [Supported resources and requirements]in the
// Compute Optimizer User Guide.
//
// [Supported resources and requirements]: https://docs.aws.amazon.com/compute-optimizer/latest/ug/requirements.html
func computeoptimizer_GetECSServiceRecommendations(cfg aws.Config, client *computeoptimizer.Client) {
	input := &computeoptimizer.GetECSServiceRecommendationsInput{}

	if len(_computeoptimizerAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _computeoptimizerAccountIds...)
	}
	if len(_computeoptimizerFilters) > 0 {
		if err := assignInputField(input, "Filters", _computeoptimizerFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _computeoptimizerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerNextToken) > 0 {
		input.NextToken = aws.String(_computeoptimizerNextToken)
	}
	if len(_computeoptimizerServiceArns) > 0 {
		input.ServiceArns = append([]string(nil), _computeoptimizerServiceArns...)
	}

	if resp, err := client.GetECSServiceRecommendations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the recommendation preferences that are in effect for a given resource,
// such as enhanced infrastructure metrics. Considers all applicable preferences
// that you might have set at the resource, account, and organization level.
//
// When you create a recommendation preference, you can set its status to Active
// or Inactive . Use this action to view the recommendation preferences that are in
// effect, or Active .
func computeoptimizer_GetEffectiveRecommendationPreferences(cfg aws.Config, client *computeoptimizer.Client) {
	input := &computeoptimizer.GetEffectiveRecommendationPreferencesInput{
		// ResourceArn: *string, // Required
	}

	if len(_computeoptimizerResourceArn) > 0 {
		input.ResourceArn = aws.String(_computeoptimizerResourceArn)
	}

	if resp, err := client.GetEffectiveRecommendationPreferences(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the enrollment (opt in) status of an account to the Compute Optimizer
// service.
//
// If the account is the management account of an organization, this action also
// confirms the enrollment status of member accounts of the organization. Use the GetEnrollmentStatusesForOrganization
// action to get detailed information about the enrollment status of member
// accounts of an organization.
func computeoptimizer_GetEnrollmentStatus(cfg aws.Config, client *computeoptimizer.Client) {
	input := &computeoptimizer.GetEnrollmentStatusInput{}

	if resp, err := client.GetEnrollmentStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the Compute Optimizer enrollment (opt-in) status of organization member
// accounts, if your account is an organization management account.
//
// To get the enrollment status of standalone accounts, use the GetEnrollmentStatus action.
func computeoptimizer_GetEnrollmentStatusesForOrganization(cfg aws.Config, client *computeoptimizer.Client) {
	input := &computeoptimizer.GetEnrollmentStatusesForOrganizationInput{}

	if len(_computeoptimizerFilters) > 0 {
		if err := assignInputField(input, "Filters", _computeoptimizerFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _computeoptimizerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerNextToken) > 0 {
		input.NextToken = aws.String(_computeoptimizerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetEnrollmentStatusesForOrganization(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*computeoptimizer.GetEnrollmentStatusesForOrganizationOutput
	p := computeoptimizer.NewGetEnrollmentStatusesForOrganizationPaginator(client, input)
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

// Returns idle resource recommendations. Compute Optimizer generates
// recommendations for idle resources that meet a specific set of requirements. For
// more information, see [Resource requirements]in the Compute Optimizer User Guide
//
// [Resource requirements]: https://docs.aws.amazon.com/compute-optimizer/latest/ug/requirements.html
func computeoptimizer_GetIdleRecommendations(cfg aws.Config, client *computeoptimizer.Client) {
	input := &computeoptimizer.GetIdleRecommendationsInput{}

	if len(_computeoptimizerAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _computeoptimizerAccountIds...)
	}
	if len(_computeoptimizerFilters) > 0 {
		if err := assignInputField(input, "Filters", _computeoptimizerFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _computeoptimizerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerNextToken) > 0 {
		input.NextToken = aws.String(_computeoptimizerNextToken)
	}
	if len(_computeoptimizerOrderBy) > 0 {
		if err := assignInputField(input, "OrderBy", _computeoptimizerOrderBy); err != nil {
			log.Errorf("invalid --order-by: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerResourceArns) > 0 {
		input.ResourceArns = append([]string(nil), _computeoptimizerResourceArns...)
	}

	if resp, err := client.GetIdleRecommendations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns Lambda function recommendations.
// Compute Optimizer generates recommendations for functions that meet a specific
// set of requirements. For more information, see the [Supported resources and requirements]in the Compute Optimizer
// User Guide.
//
// [Supported resources and requirements]: https://docs.aws.amazon.com/compute-optimizer/latest/ug/requirements.html
func computeoptimizer_GetLambdaFunctionRecommendations(cfg aws.Config, client *computeoptimizer.Client) {
	input := &computeoptimizer.GetLambdaFunctionRecommendationsInput{}

	if len(_computeoptimizerAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _computeoptimizerAccountIds...)
	}
	if len(_computeoptimizerFilters) > 0 {
		if err := assignInputField(input, "Filters", _computeoptimizerFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerFunctionArns) > 0 {
		input.FunctionArns = append([]string(nil), _computeoptimizerFunctionArns...)
	}
	if len(_computeoptimizerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _computeoptimizerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerNextToken) > 0 {
		input.NextToken = aws.String(_computeoptimizerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetLambdaFunctionRecommendations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*computeoptimizer.GetLambdaFunctionRecommendationsOutput
	p := computeoptimizer.NewGetLambdaFunctionRecommendationsPaginator(client, input)
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

// Returns license recommendations for Amazon EC2 instances that run on a specific
// license.
//
// Compute Optimizer generates recommendations for licenses that meet a specific
// set of requirements. For more information, see the [Supported resources and requirements]in the Compute Optimizer
// User Guide.
//
// [Supported resources and requirements]: https://docs.aws.amazon.com/compute-optimizer/latest/ug/requirements.html
func computeoptimizer_GetLicenseRecommendations(cfg aws.Config, client *computeoptimizer.Client) {
	input := &computeoptimizer.GetLicenseRecommendationsInput{}

	if len(_computeoptimizerAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _computeoptimizerAccountIds...)
	}
	if len(_computeoptimizerFilters) > 0 {
		if err := assignInputField(input, "Filters", _computeoptimizerFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _computeoptimizerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerNextToken) > 0 {
		input.NextToken = aws.String(_computeoptimizerNextToken)
	}
	if len(_computeoptimizerResourceArns) > 0 {
		input.ResourceArns = append([]string(nil), _computeoptimizerResourceArns...)
	}

	if resp, err := client.GetLicenseRecommendations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the projected metrics of Aurora and RDS database recommendations.
func computeoptimizer_GetRDSDatabaseRecommendationProjectedMetrics(cfg aws.Config, client *computeoptimizer.Client) {
	input := &computeoptimizer.GetRDSDatabaseRecommendationProjectedMetricsInput{
		// EndTime: *time.Time, // Required
		// Period: int32, // Required
		// ResourceArn: *string, // Required
		// StartTime: *time.Time, // Required
		// Stat: types.MetricStatistic, // Required
	}

	if len(_computeoptimizerEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _computeoptimizerEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerPeriod) > 0 {
		if err := assignInputField(input, "Period", _computeoptimizerPeriod); err != nil {
			log.Errorf("invalid --period: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerResourceArn) > 0 {
		input.ResourceArn = aws.String(_computeoptimizerResourceArn)
	}
	if len(_computeoptimizerStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _computeoptimizerStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerStat) > 0 {
		if err := assignInputField(input, "Stat", _computeoptimizerStat); err != nil {
			log.Errorf("invalid --stat: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerRecommendationPreferences) > 0 {
		if err := assignInputField(input, "RecommendationPreferences", _computeoptimizerRecommendationPreferences); err != nil {
			log.Errorf("invalid --recommendation-preferences: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetRDSDatabaseRecommendationProjectedMetrics(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns Amazon Aurora and RDS database recommendations.
// Compute Optimizer generates recommendations for Amazon Aurora and RDS databases
// that meet a specific set of requirements. For more information, see the [Supported resources and requirements]in the
// Compute Optimizer User Guide.
//
// [Supported resources and requirements]: https://docs.aws.amazon.com/compute-optimizer/latest/ug/requirements.html
func computeoptimizer_GetRDSDatabaseRecommendations(cfg aws.Config, client *computeoptimizer.Client) {
	input := &computeoptimizer.GetRDSDatabaseRecommendationsInput{}

	if len(_computeoptimizerAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _computeoptimizerAccountIds...)
	}
	if len(_computeoptimizerFilters) > 0 {
		if err := assignInputField(input, "Filters", _computeoptimizerFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _computeoptimizerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerNextToken) > 0 {
		input.NextToken = aws.String(_computeoptimizerNextToken)
	}
	if len(_computeoptimizerRecommendationPreferences) > 0 {
		if err := assignInputField(input, "RecommendationPreferences", _computeoptimizerRecommendationPreferences); err != nil {
			log.Errorf("invalid --recommendation-preferences: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerResourceArns) > 0 {
		input.ResourceArns = append([]string(nil), _computeoptimizerResourceArns...)
	}

	if resp, err := client.GetRDSDatabaseRecommendations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns existing recommendation preferences, such as enhanced infrastructure
// metrics.
//
// Use the scope parameter to specify which preferences to return. You can specify
// to return preferences for an organization, a specific account ID, or a specific
// EC2 instance or Amazon EC2 Auto Scaling group Amazon Resource Name (ARN).
//
// For more information, see [Activating enhanced infrastructure metrics] in the Compute Optimizer User Guide.
//
// [Activating enhanced infrastructure metrics]: https://docs.aws.amazon.com/compute-optimizer/latest/ug/enhanced-infrastructure-metrics.html
func computeoptimizer_GetRecommendationPreferences(cfg aws.Config, client *computeoptimizer.Client) {
	input := &computeoptimizer.GetRecommendationPreferencesInput{
		// ResourceType: types.ResourceType, // Required
	}

	if len(_computeoptimizerResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _computeoptimizerResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _computeoptimizerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerNextToken) > 0 {
		input.NextToken = aws.String(_computeoptimizerNextToken)
	}
	if len(_computeoptimizerScope) > 0 {
		if err := assignInputField(input, "Scope", _computeoptimizerScope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetRecommendationPreferences(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*computeoptimizer.GetRecommendationPreferencesOutput
	p := computeoptimizer.NewGetRecommendationPreferencesPaginator(client, input)
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

// Returns the optimization findings for an account.
// It returns the number of:
//
// - Amazon EC2 instances in an account that are Underprovisioned ,
// Overprovisioned , or Optimized .
//
// - EC2Amazon EC2 Auto Scaling groups in an account that are NotOptimized , or
// Optimized .
//
// - Amazon EBS volumes in an account that are NotOptimized , or Optimized .
//
// - Lambda functions in an account that are NotOptimized , or Optimized .
//
// - Amazon ECS services in an account that are Underprovisioned ,
// Overprovisioned , or Optimized .
//
// - Commercial software licenses in an account that are InsufficientMetrics ,
// NotOptimized or Optimized .
//
// - Amazon Aurora and Amazon RDS databases in an account that are
// Underprovisioned , Overprovisioned , Optimized , or NotOptimized .
func computeoptimizer_GetRecommendationSummaries(cfg aws.Config, client *computeoptimizer.Client) {
	input := &computeoptimizer.GetRecommendationSummariesInput{}

	if len(_computeoptimizerAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _computeoptimizerAccountIds...)
	}
	if len(_computeoptimizerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _computeoptimizerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerNextToken) > 0 {
		input.NextToken = aws.String(_computeoptimizerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetRecommendationSummaries(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*computeoptimizer.GetRecommendationSummariesOutput
	p := computeoptimizer.NewGetRecommendationSummariesPaginator(client, input)
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

// Creates a new recommendation preference or updates an existing recommendation
// preference, such as enhanced infrastructure metrics.
//
// For more information, see [Activating enhanced infrastructure metrics] in the Compute Optimizer User Guide.
//
// [Activating enhanced infrastructure metrics]: https://docs.aws.amazon.com/compute-optimizer/latest/ug/enhanced-infrastructure-metrics.html
func computeoptimizer_PutRecommendationPreferences(cfg aws.Config, client *computeoptimizer.Client) {
	input := &computeoptimizer.PutRecommendationPreferencesInput{
		// ResourceType: types.ResourceType, // Required
	}

	if len(_computeoptimizerResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _computeoptimizerResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerEnhancedInfrastructureMetrics) > 0 {
		if err := assignInputField(input, "EnhancedInfrastructureMetrics", _computeoptimizerEnhancedInfrastructureMetrics); err != nil {
			log.Errorf("invalid --enhanced-infrastructure-metrics: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerExternalMetricsPreference) > 0 {
		if err := assignInputField(input, "ExternalMetricsPreference", _computeoptimizerExternalMetricsPreference); err != nil {
			log.Errorf("invalid --external-metrics-preference: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerInferredWorkloadTypes) > 0 {
		if err := assignInputField(input, "InferredWorkloadTypes", _computeoptimizerInferredWorkloadTypes); err != nil {
			log.Errorf("invalid --inferred-workload-types: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerLookBackPeriod) > 0 {
		if err := assignInputField(input, "LookBackPeriod", _computeoptimizerLookBackPeriod); err != nil {
			log.Errorf("invalid --look-back-period: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerPreferredResources) > 0 {
		if err := assignInputField(input, "PreferredResources", _computeoptimizerPreferredResources); err != nil {
			log.Errorf("invalid --preferred-resources: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerSavingsEstimationMode) > 0 {
		if err := assignInputField(input, "SavingsEstimationMode", _computeoptimizerSavingsEstimationMode); err != nil {
			log.Errorf("invalid --savings-estimation-mode: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerScope) > 0 {
		if err := assignInputField(input, "Scope", _computeoptimizerScope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerUtilizationPreferences) > 0 {
		if err := assignInputField(input, "UtilizationPreferences", _computeoptimizerUtilizationPreferences); err != nil {
			log.Errorf("invalid --utilization-preferences: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutRecommendationPreferences(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the enrollment (opt in and opt out) status of an account to the Compute
// Optimizer service.
//
// If the account is a management account of an organization, this action can also
// be used to enroll member accounts of the organization.
//
// You must have the appropriate permissions to opt in to Compute Optimizer, to
// view its recommendations, and to opt out. For more information, see [Controlling access with Amazon Web Services Identity and Access Management]in the
// Compute Optimizer User Guide.
//
// When you opt in, Compute Optimizer automatically creates a service-linked role
// in your account to access its data. For more information, see [Using Service-Linked Roles for Compute Optimizer]in the Compute
// Optimizer User Guide.
//
// [Controlling access with Amazon Web Services Identity and Access Management]: https://docs.aws.amazon.com/compute-optimizer/latest/ug/security-iam.html
// [Using Service-Linked Roles for Compute Optimizer]: https://docs.aws.amazon.com/compute-optimizer/latest/ug/using-service-linked-roles.html
func computeoptimizer_UpdateEnrollmentStatus(cfg aws.Config, client *computeoptimizer.Client) {
	input := &computeoptimizer.UpdateEnrollmentStatusInput{
		// Status: types.Status, // Required
	}

	if len(_computeoptimizerStatus) > 0 {
		if err := assignInputField(input, "Status", _computeoptimizerStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerIncludeMemberAccounts) > 0 {
		if err := assignInputField(input, "IncludeMemberAccounts", _computeoptimizerIncludeMemberAccounts); err != nil {
			log.Errorf("invalid --include-member-accounts: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateEnrollmentStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_computeoptimizerCmd)
	_computeoptimizerCmd.Flags().SortFlags = false

	_computeoptimizerCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_computeoptimizerCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_computeoptimizerCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_computeoptimizerCmd.Flags().StringSliceVarP(&_computeoptimizerAccountIds, "account-ids", "", nil, "Account Ids")
	_computeoptimizerCmd.Flags().StringSliceVarP(&_computeoptimizerAutoScalingGroupArns, "auto-scaling-group-arns", "", nil, "Auto Scaling Group Arns")
	_computeoptimizerCmd.Flags().StringVarP(&_computeoptimizerEndTime, "end-time", "", "", "End Time")
	_computeoptimizerCmd.Flags().StringVarP(&_computeoptimizerEnhancedInfrastructureMetrics, "enhanced-infrastructure-metrics", "", "", "Enhanced Infrastructure Metrics")
	_computeoptimizerCmd.Flags().StringVarP(&_computeoptimizerExternalMetricsPreference, "external-metrics-preference", "", "", "External Metrics Preference")
	_computeoptimizerCmd.Flags().StringVarP(&_computeoptimizerFieldsToExport, "fields-to-export", "", "", "Fields To Export")
	_computeoptimizerCmd.Flags().StringVarP(&_computeoptimizerFileFormat, "file-format", "", "", "File Format")
	_computeoptimizerCmd.Flags().StringVarP(&_computeoptimizerFilters, "filters", "", "", "Filters")
	_computeoptimizerCmd.Flags().StringSliceVarP(&_computeoptimizerFunctionArns, "function-arns", "", nil, "Function Arns")
	_computeoptimizerCmd.Flags().StringVarP(&_computeoptimizerIncludeMemberAccounts, "include-member-accounts", "", "", "Include Member Accounts")
	_computeoptimizerCmd.Flags().StringVarP(&_computeoptimizerInferredWorkloadTypes, "inferred-workload-types", "", "", "Inferred Workload Types")
	_computeoptimizerCmd.Flags().StringVarP(&_computeoptimizerInstanceArn, "instance-arn", "", "", "Instance ARN")
	_computeoptimizerCmd.Flags().StringSliceVarP(&_computeoptimizerInstanceArns, "instance-arns", "", nil, "Instance Arns")
	_computeoptimizerCmd.Flags().StringSliceVarP(&_computeoptimizerJobIds, "job-ids", "", nil, "Job Ids")
	_computeoptimizerCmd.Flags().StringVarP(&_computeoptimizerLookBackPeriod, "look-back-period", "", "", "Look Back Period")
	_computeoptimizerCmd.Flags().StringVarP(&_computeoptimizerMaxResults, "max-results", "", "", "Max Results")
	_computeoptimizerCmd.Flags().StringVarP(&_computeoptimizerNextToken, "next-token", "", "", "Next Token")
	_computeoptimizerCmd.Flags().StringVarP(&_computeoptimizerOrderBy, "order-by", "", "", "Order By")
	_computeoptimizerCmd.Flags().StringVarP(&_computeoptimizerPeriod, "period", "", "", "Period")
	_computeoptimizerCmd.Flags().StringVarP(&_computeoptimizerPreferredResources, "preferred-resources", "", "", "Preferred Resources")
	_computeoptimizerCmd.Flags().StringVarP(&_computeoptimizerRecommendationPreferenceNames, "recommendation-preference-names", "", "", "Recommendation Preference Names")
	_computeoptimizerCmd.Flags().StringVarP(&_computeoptimizerRecommendationPreferences, "recommendation-preferences", "", "", "Recommendation Preferences")
	_computeoptimizerCmd.Flags().StringVarP(&_computeoptimizerResourceArn, "resource-arn", "", "", "Resource ARN")
	_computeoptimizerCmd.Flags().StringSliceVarP(&_computeoptimizerResourceArns, "resource-arns", "", nil, "Resource Arns")
	_computeoptimizerCmd.Flags().StringVarP(&_computeoptimizerResourceType, "resource-type", "", "", "Resource Type")
	_computeoptimizerCmd.Flags().StringVarP(&_computeoptimizerS3DestinationConfig, "s3-destination-config", "", "", "S3 Destination Config")
	_computeoptimizerCmd.Flags().StringVarP(&_computeoptimizerSavingsEstimationMode, "savings-estimation-mode", "", "", "Savings Estimation Mode")
	_computeoptimizerCmd.Flags().StringVarP(&_computeoptimizerScope, "scope", "", "", "Scope")
	_computeoptimizerCmd.Flags().StringVarP(&_computeoptimizerServiceArn, "service-arn", "", "", "Service ARN")
	_computeoptimizerCmd.Flags().StringSliceVarP(&_computeoptimizerServiceArns, "service-arns", "", nil, "Service Arns")
	_computeoptimizerCmd.Flags().StringVarP(&_computeoptimizerStartTime, "start-time", "", "", "Start Time")
	_computeoptimizerCmd.Flags().StringVarP(&_computeoptimizerStat, "stat", "", "", "Stat")
	_computeoptimizerCmd.Flags().StringVarP(&_computeoptimizerStatus, "status", "", "", "Status")
	_computeoptimizerCmd.Flags().StringVarP(&_computeoptimizerUtilizationPreferences, "utilization-preferences", "", "", "Utilization Preferences")
	_computeoptimizerCmd.Flags().StringSliceVarP(&_computeoptimizerVolumeArns, "volume-arns", "", nil, "Volume Arns")

	_computeoptimizerCmd.Flags().BoolVarP(&_computeoptimizerDeleteRecommendationPreferences, "delete-recommendation-preferences", "", false, "Delete Recommendation Preferences")
	_computeoptimizerCmd.Flags().BoolVarP(&_computeoptimizerDescribeRecommendationExportJobs, "describe-recommendation-export-jobs", "", false, "Describe Recommendation Export Jobs")
	_computeoptimizerCmd.Flags().BoolVarP(&_computeoptimizerExportAutoScalingGroupRecommendations, "export-auto-scaling-group-recommendations", "", false, "Export Auto Scaling Group Recommendations")
	_computeoptimizerCmd.Flags().BoolVarP(&_computeoptimizerExportEBSVolumeRecommendations, "export-ebs-volume-recommendations", "", false, "Export Ebs Volume Recommendations")
	_computeoptimizerCmd.Flags().BoolVarP(&_computeoptimizerExportEC2InstanceRecommendations, "export-ec2-instance-recommendations", "", false, "Export EC2 Instance Recommendations")
	_computeoptimizerCmd.Flags().BoolVarP(&_computeoptimizerExportECSServiceRecommendations, "export-ecs-service-recommendations", "", false, "Export Ecs Service Recommendations")
	_computeoptimizerCmd.Flags().BoolVarP(&_computeoptimizerExportIdleRecommendations, "export-idle-recommendations", "", false, "Export Idle Recommendations")
	_computeoptimizerCmd.Flags().BoolVarP(&_computeoptimizerExportLambdaFunctionRecommendations, "export-lambda-function-recommendations", "", false, "Export Lambda Function Recommendations")
	_computeoptimizerCmd.Flags().BoolVarP(&_computeoptimizerExportLicenseRecommendations, "export-license-recommendations", "", false, "Export License Recommendations")
	_computeoptimizerCmd.Flags().BoolVarP(&_computeoptimizerExportRDSDatabaseRecommendations, "export-rds-database-recommendations", "", false, "Export RDS Database Recommendations")
	_computeoptimizerCmd.Flags().BoolVarP(&_computeoptimizerGetAutoScalingGroupRecommendations, "get-auto-scaling-group-recommendations", "", false, "Get Auto Scaling Group Recommendations")
	_computeoptimizerCmd.Flags().BoolVarP(&_computeoptimizerGetEBSVolumeRecommendations, "get-ebs-volume-recommendations", "", false, "Get Ebs Volume Recommendations")
	_computeoptimizerCmd.Flags().BoolVarP(&_computeoptimizerGetEC2InstanceRecommendations, "get-ec2-instance-recommendations", "", false, "Get EC2 Instance Recommendations")
	_computeoptimizerCmd.Flags().BoolVarP(&_computeoptimizerGetEC2RecommendationProjectedMetrics, "get-ec2-recommendation-projected-metrics", "", false, "Get EC2 Recommendation Projected Metrics")
	_computeoptimizerCmd.Flags().BoolVarP(&_computeoptimizerGetECSServiceRecommendationProjectedMetrics, "get-ecs-service-recommendation-projected-metrics", "", false, "Get Ecs Service Recommendation Projected Metrics")
	_computeoptimizerCmd.Flags().BoolVarP(&_computeoptimizerGetECSServiceRecommendations, "get-ecs-service-recommendations", "", false, "Get Ecs Service Recommendations")
	_computeoptimizerCmd.Flags().BoolVarP(&_computeoptimizerGetEffectiveRecommendationPreferences, "get-effective-recommendation-preferences", "", false, "Get Effective Recommendation Preferences")
	_computeoptimizerCmd.Flags().BoolVarP(&_computeoptimizerGetEnrollmentStatus, "get-enrollment-status", "", false, "Get Enrollment Status")
	_computeoptimizerCmd.Flags().BoolVarP(&_computeoptimizerGetEnrollmentStatusesForOrganization, "get-enrollment-statuses-for-organization", "", false, "Get Enrollment Statuses For Organization")
	_computeoptimizerCmd.Flags().BoolVarP(&_computeoptimizerGetIdleRecommendations, "get-idle-recommendations", "", false, "Get Idle Recommendations")
	_computeoptimizerCmd.Flags().BoolVarP(&_computeoptimizerGetLambdaFunctionRecommendations, "get-lambda-function-recommendations", "", false, "Get Lambda Function Recommendations")
	_computeoptimizerCmd.Flags().BoolVarP(&_computeoptimizerGetLicenseRecommendations, "get-license-recommendations", "", false, "Get License Recommendations")
	_computeoptimizerCmd.Flags().BoolVarP(&_computeoptimizerGetRDSDatabaseRecommendationProjectedMetrics, "get-rds-database-recommendation-projected-metrics", "", false, "Get RDS Database Recommendation Projected Metrics")
	_computeoptimizerCmd.Flags().BoolVarP(&_computeoptimizerGetRDSDatabaseRecommendations, "get-rds-database-recommendations", "", false, "Get RDS Database Recommendations")
	_computeoptimizerCmd.Flags().BoolVarP(&_computeoptimizerGetRecommendationPreferences, "get-recommendation-preferences", "", false, "Get Recommendation Preferences")
	_computeoptimizerCmd.Flags().BoolVarP(&_computeoptimizerGetRecommendationSummaries, "get-recommendation-summaries", "", false, "Get Recommendation Summaries")
	_computeoptimizerCmd.Flags().BoolVarP(&_computeoptimizerPutRecommendationPreferences, "put-recommendation-preferences", "", false, "Put Recommendation Preferences")
	_computeoptimizerCmd.Flags().BoolVarP(&_computeoptimizerUpdateEnrollmentStatus, "update-enrollment-status", "", false, "Update Enrollment Status")

}
