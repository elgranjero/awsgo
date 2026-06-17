package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/migrationhubstrategy"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// migrationhubstrategyCmd represents the migrationhubstrategy command
var _migrationhubstrategyCmd = &cobra.Command{
	Use:   "migrationhubstrategy",
	Short: "AWS migrationhubstrategy CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := migrationhubstrategy.NewFromConfig(cfg)
		if _migrationhubstrategyGetApplicationComponentDetails {
			migrationhubstrategy_GetApplicationComponentDetails(cfg, client)
			return
		}
		if _migrationhubstrategyGetApplicationComponentStrategies {
			migrationhubstrategy_GetApplicationComponentStrategies(cfg, client)
			return
		}
		if _migrationhubstrategyGetAssessment {
			migrationhubstrategy_GetAssessment(cfg, client)
			return
		}
		if _migrationhubstrategyGetImportFileTask {
			migrationhubstrategy_GetImportFileTask(cfg, client)
			return
		}
		if _migrationhubstrategyGetLatestAssessmentId {
			migrationhubstrategy_GetLatestAssessmentId(cfg, client)
			return
		}
		if _migrationhubstrategyGetPortfolioPreferences {
			migrationhubstrategy_GetPortfolioPreferences(cfg, client)
			return
		}
		if _migrationhubstrategyGetPortfolioSummary {
			migrationhubstrategy_GetPortfolioSummary(cfg, client)
			return
		}
		if _migrationhubstrategyGetRecommendationReportDetails {
			migrationhubstrategy_GetRecommendationReportDetails(cfg, client)
			return
		}
		if _migrationhubstrategyGetServerDetails {
			migrationhubstrategy_GetServerDetails(cfg, client)
			return
		}
		if _migrationhubstrategyGetServerStrategies {
			migrationhubstrategy_GetServerStrategies(cfg, client)
			return
		}
		if _migrationhubstrategyListAnalyzableServers {
			migrationhubstrategy_ListAnalyzableServers(cfg, client)
			return
		}
		if _migrationhubstrategyListApplicationComponents {
			migrationhubstrategy_ListApplicationComponents(cfg, client)
			return
		}
		if _migrationhubstrategyListCollectors {
			migrationhubstrategy_ListCollectors(cfg, client)
			return
		}
		if _migrationhubstrategyListImportFileTask {
			migrationhubstrategy_ListImportFileTask(cfg, client)
			return
		}
		if _migrationhubstrategyListServers {
			migrationhubstrategy_ListServers(cfg, client)
			return
		}
		if _migrationhubstrategyPutPortfolioPreferences {
			migrationhubstrategy_PutPortfolioPreferences(cfg, client)
			return
		}
		if _migrationhubstrategyStartAssessment {
			migrationhubstrategy_StartAssessment(cfg, client)
			return
		}
		if _migrationhubstrategyStartImportFileTask {
			migrationhubstrategy_StartImportFileTask(cfg, client)
			return
		}
		if _migrationhubstrategyStartRecommendationReportGeneration {
			migrationhubstrategy_StartRecommendationReportGeneration(cfg, client)
			return
		}
		if _migrationhubstrategyStopAssessment {
			migrationhubstrategy_StopAssessment(cfg, client)
			return
		}
		if _migrationhubstrategyUpdateApplicationComponentConfig {
			migrationhubstrategy_UpdateApplicationComponentConfig(cfg, client)
			return
		}
		if _migrationhubstrategyUpdateServerConfig {
			migrationhubstrategy_UpdateServerConfig(cfg, client)
			return
		}

	},
}

var (
	_migrationhubstrategyGetApplicationComponentDetails      bool
	_migrationhubstrategyGetApplicationComponentStrategies   bool
	_migrationhubstrategyGetAssessment                       bool
	_migrationhubstrategyGetImportFileTask                   bool
	_migrationhubstrategyGetLatestAssessmentId               bool
	_migrationhubstrategyGetPortfolioPreferences             bool
	_migrationhubstrategyGetPortfolioSummary                 bool
	_migrationhubstrategyGetRecommendationReportDetails      bool
	_migrationhubstrategyGetServerDetails                    bool
	_migrationhubstrategyGetServerStrategies                 bool
	_migrationhubstrategyListAnalyzableServers               bool
	_migrationhubstrategyListApplicationComponents           bool
	_migrationhubstrategyListCollectors                      bool
	_migrationhubstrategyListImportFileTask                  bool
	_migrationhubstrategyListServers                         bool
	_migrationhubstrategyPutPortfolioPreferences             bool
	_migrationhubstrategyStartAssessment                     bool
	_migrationhubstrategyStartImportFileTask                 bool
	_migrationhubstrategyStartRecommendationReportGeneration bool
	_migrationhubstrategyStopAssessment                      bool
	_migrationhubstrategyUpdateApplicationComponentConfig    bool
	_migrationhubstrategyUpdateServerConfig                  bool

	_migrationhubstrategyAppType                      string
	_migrationhubstrategyApplicationComponentCriteria string
	_migrationhubstrategyApplicationComponentId       string
	_migrationhubstrategyApplicationMode              string
	_migrationhubstrategyApplicationPreferences       string
	_migrationhubstrategyAssessmentDataSourceType     string
	_migrationhubstrategyAssessmentId                 string
	_migrationhubstrategyAssessmentTargets            string
	_migrationhubstrategyConfigureOnly                string
	_migrationhubstrategyDataSourceType               string
	_migrationhubstrategyDatabasePreferences          string
	_migrationhubstrategyFilterValue                  string
	_migrationhubstrategyGroupId                      string
	_migrationhubstrategyGroupIdFilter                string
	_migrationhubstrategyId                           string
	_migrationhubstrategyInclusionStatus              string
	_migrationhubstrategyMaxResults                   string
	_migrationhubstrategyName                         string
	_migrationhubstrategyNextToken                    string
	_migrationhubstrategyOutputFormat                 string
	_migrationhubstrategyPrioritizeBusinessGoals      string
	_migrationhubstrategyS3Bucket                     string
	_migrationhubstrategyS3bucketForAnalysisData      string
	_migrationhubstrategyS3bucketForReportData        string
	_migrationhubstrategyS3key                        string
	_migrationhubstrategySecretsManagerKey            string
	_migrationhubstrategyServerCriteria               string
	_migrationhubstrategyServerId                     string
	_migrationhubstrategySort                         string
	_migrationhubstrategySourceCodeList               string
	_migrationhubstrategyStrategyOption               string
)

// Retrieves details about an application component.
func migrationhubstrategy_GetApplicationComponentDetails(cfg aws.Config, client *migrationhubstrategy.Client) {
	input := &migrationhubstrategy.GetApplicationComponentDetailsInput{
		// ApplicationComponentId: *string, // Required
	}

	if len(_migrationhubstrategyApplicationComponentId) > 0 {
		input.ApplicationComponentId = aws.String(_migrationhubstrategyApplicationComponentId)
	}

	if resp, err := client.GetApplicationComponentDetails(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of all the recommended strategies and tools for an
// application component running on a server.
func migrationhubstrategy_GetApplicationComponentStrategies(cfg aws.Config, client *migrationhubstrategy.Client) {
	input := &migrationhubstrategy.GetApplicationComponentStrategiesInput{
		// ApplicationComponentId: *string, // Required
	}

	if len(_migrationhubstrategyApplicationComponentId) > 0 {
		input.ApplicationComponentId = aws.String(_migrationhubstrategyApplicationComponentId)
	}

	if resp, err := client.GetApplicationComponentStrategies(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the status of an on-going assessment.
func migrationhubstrategy_GetAssessment(cfg aws.Config, client *migrationhubstrategy.Client) {
	input := &migrationhubstrategy.GetAssessmentInput{
		// Id: *string, // Required
	}

	if len(_migrationhubstrategyId) > 0 {
		input.Id = aws.String(_migrationhubstrategyId)
	}

	if resp, err := client.GetAssessment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the details about a specific import task.
func migrationhubstrategy_GetImportFileTask(cfg aws.Config, client *migrationhubstrategy.Client) {
	input := &migrationhubstrategy.GetImportFileTaskInput{
		// Id: *string, // Required
	}

	if len(_migrationhubstrategyId) > 0 {
		input.Id = aws.String(_migrationhubstrategyId)
	}

	if resp, err := client.GetImportFileTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieve the latest ID of a specific assessment task.
func migrationhubstrategy_GetLatestAssessmentId(cfg aws.Config, client *migrationhubstrategy.Client) {
	input := &migrationhubstrategy.GetLatestAssessmentIdInput{}

	if resp, err := client.GetLatestAssessmentId(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves your migration and modernization preferences.
func migrationhubstrategy_GetPortfolioPreferences(cfg aws.Config, client *migrationhubstrategy.Client) {
	input := &migrationhubstrategy.GetPortfolioPreferencesInput{}

	if resp, err := client.GetPortfolioPreferences(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves overall summary including the number of servers to rehost and the
// overall number of anti-patterns.
func migrationhubstrategy_GetPortfolioSummary(cfg aws.Config, client *migrationhubstrategy.Client) {
	input := &migrationhubstrategy.GetPortfolioSummaryInput{}

	if resp, err := client.GetPortfolioSummary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves detailed information about the specified recommendation report.
func migrationhubstrategy_GetRecommendationReportDetails(cfg aws.Config, client *migrationhubstrategy.Client) {
	input := &migrationhubstrategy.GetRecommendationReportDetailsInput{
		// Id: *string, // Required
	}

	if len(_migrationhubstrategyId) > 0 {
		input.Id = aws.String(_migrationhubstrategyId)
	}

	if resp, err := client.GetRecommendationReportDetails(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves detailed information about a specified server.
func migrationhubstrategy_GetServerDetails(cfg aws.Config, client *migrationhubstrategy.Client) {
	input := &migrationhubstrategy.GetServerDetailsInput{
		// ServerId: *string, // Required
	}

	if len(_migrationhubstrategyServerId) > 0 {
		input.ServerId = aws.String(_migrationhubstrategyServerId)
	}
	if len(_migrationhubstrategyMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _migrationhubstrategyMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_migrationhubstrategyNextToken) > 0 {
		input.NextToken = aws.String(_migrationhubstrategyNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetServerDetails(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*migrationhubstrategy.GetServerDetailsOutput
	p := migrationhubstrategy.NewGetServerDetailsPaginator(client, input)
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

// Retrieves recommended strategies and tools for the specified server.
func migrationhubstrategy_GetServerStrategies(cfg aws.Config, client *migrationhubstrategy.Client) {
	input := &migrationhubstrategy.GetServerStrategiesInput{
		// ServerId: *string, // Required
	}

	if len(_migrationhubstrategyServerId) > 0 {
		input.ServerId = aws.String(_migrationhubstrategyServerId)
	}

	if resp, err := client.GetServerStrategies(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of all the servers fetched from customer vCenter using
// Strategy Recommendation Collector.
func migrationhubstrategy_ListAnalyzableServers(cfg aws.Config, client *migrationhubstrategy.Client) {
	input := &migrationhubstrategy.ListAnalyzableServersInput{}

	if len(_migrationhubstrategyMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _migrationhubstrategyMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_migrationhubstrategyNextToken) > 0 {
		input.NextToken = aws.String(_migrationhubstrategyNextToken)
	}
	if len(_migrationhubstrategySort) > 0 {
		if err := assignInputField(input, "Sort", _migrationhubstrategySort); err != nil {
			log.Errorf("invalid --sort: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListAnalyzableServers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*migrationhubstrategy.ListAnalyzableServersOutput
	p := migrationhubstrategy.NewListAnalyzableServersPaginator(client, input)
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

// Retrieves a list of all the application components (processes).
func migrationhubstrategy_ListApplicationComponents(cfg aws.Config, client *migrationhubstrategy.Client) {
	input := &migrationhubstrategy.ListApplicationComponentsInput{}

	if len(_migrationhubstrategyApplicationComponentCriteria) > 0 {
		if err := assignInputField(input, "ApplicationComponentCriteria", _migrationhubstrategyApplicationComponentCriteria); err != nil {
			log.Errorf("invalid --application-component-criteria: %s", err.Error())
			return
		}
	}
	if len(_migrationhubstrategyFilterValue) > 0 {
		input.FilterValue = aws.String(_migrationhubstrategyFilterValue)
	}
	if len(_migrationhubstrategyGroupIdFilter) > 0 {
		if err := assignInputField(input, "GroupIdFilter", _migrationhubstrategyGroupIdFilter); err != nil {
			log.Errorf("invalid --group-id-filter: %s", err.Error())
			return
		}
	}
	if len(_migrationhubstrategyMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _migrationhubstrategyMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_migrationhubstrategyNextToken) > 0 {
		input.NextToken = aws.String(_migrationhubstrategyNextToken)
	}
	if len(_migrationhubstrategySort) > 0 {
		if err := assignInputField(input, "Sort", _migrationhubstrategySort); err != nil {
			log.Errorf("invalid --sort: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListApplicationComponents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*migrationhubstrategy.ListApplicationComponentsOutput
	p := migrationhubstrategy.NewListApplicationComponentsPaginator(client, input)
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

// Retrieves a list of all the installed collectors.
func migrationhubstrategy_ListCollectors(cfg aws.Config, client *migrationhubstrategy.Client) {
	input := &migrationhubstrategy.ListCollectorsInput{}

	if len(_migrationhubstrategyMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _migrationhubstrategyMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_migrationhubstrategyNextToken) > 0 {
		input.NextToken = aws.String(_migrationhubstrategyNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCollectors(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*migrationhubstrategy.ListCollectorsOutput
	p := migrationhubstrategy.NewListCollectorsPaginator(client, input)
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

// Retrieves a list of all the imports performed.
func migrationhubstrategy_ListImportFileTask(cfg aws.Config, client *migrationhubstrategy.Client) {
	input := &migrationhubstrategy.ListImportFileTaskInput{}

	if len(_migrationhubstrategyMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _migrationhubstrategyMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_migrationhubstrategyNextToken) > 0 {
		input.NextToken = aws.String(_migrationhubstrategyNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListImportFileTask(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*migrationhubstrategy.ListImportFileTaskOutput
	p := migrationhubstrategy.NewListImportFileTaskPaginator(client, input)
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

// Returns a list of all the servers.
func migrationhubstrategy_ListServers(cfg aws.Config, client *migrationhubstrategy.Client) {
	input := &migrationhubstrategy.ListServersInput{}

	if len(_migrationhubstrategyFilterValue) > 0 {
		input.FilterValue = aws.String(_migrationhubstrategyFilterValue)
	}
	if len(_migrationhubstrategyGroupIdFilter) > 0 {
		if err := assignInputField(input, "GroupIdFilter", _migrationhubstrategyGroupIdFilter); err != nil {
			log.Errorf("invalid --group-id-filter: %s", err.Error())
			return
		}
	}
	if len(_migrationhubstrategyMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _migrationhubstrategyMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_migrationhubstrategyNextToken) > 0 {
		input.NextToken = aws.String(_migrationhubstrategyNextToken)
	}
	if len(_migrationhubstrategyServerCriteria) > 0 {
		if err := assignInputField(input, "ServerCriteria", _migrationhubstrategyServerCriteria); err != nil {
			log.Errorf("invalid --server-criteria: %s", err.Error())
			return
		}
	}
	if len(_migrationhubstrategySort) > 0 {
		if err := assignInputField(input, "Sort", _migrationhubstrategySort); err != nil {
			log.Errorf("invalid --sort: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListServers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*migrationhubstrategy.ListServersOutput
	p := migrationhubstrategy.NewListServersPaginator(client, input)
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

// Saves the specified migration and modernization preferences.
func migrationhubstrategy_PutPortfolioPreferences(cfg aws.Config, client *migrationhubstrategy.Client) {
	input := &migrationhubstrategy.PutPortfolioPreferencesInput{}

	if len(_migrationhubstrategyApplicationMode) > 0 {
		if err := assignInputField(input, "ApplicationMode", _migrationhubstrategyApplicationMode); err != nil {
			log.Errorf("invalid --application-mode: %s", err.Error())
			return
		}
	}
	if len(_migrationhubstrategyApplicationPreferences) > 0 {
		if err := assignInputField(input, "ApplicationPreferences", _migrationhubstrategyApplicationPreferences); err != nil {
			log.Errorf("invalid --application-preferences: %s", err.Error())
			return
		}
	}
	if len(_migrationhubstrategyDatabasePreferences) > 0 {
		if err := assignInputField(input, "DatabasePreferences", _migrationhubstrategyDatabasePreferences); err != nil {
			log.Errorf("invalid --database-preferences: %s", err.Error())
			return
		}
	}
	if len(_migrationhubstrategyPrioritizeBusinessGoals) > 0 {
		if err := assignInputField(input, "PrioritizeBusinessGoals", _migrationhubstrategyPrioritizeBusinessGoals); err != nil {
			log.Errorf("invalid --prioritize-business-goals: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutPortfolioPreferences(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts the assessment of an on-premises environment.
func migrationhubstrategy_StartAssessment(cfg aws.Config, client *migrationhubstrategy.Client) {
	input := &migrationhubstrategy.StartAssessmentInput{}

	if len(_migrationhubstrategyAssessmentDataSourceType) > 0 {
		if err := assignInputField(input, "AssessmentDataSourceType", _migrationhubstrategyAssessmentDataSourceType); err != nil {
			log.Errorf("invalid --assessment-data-source-type: %s", err.Error())
			return
		}
	}
	if len(_migrationhubstrategyAssessmentTargets) > 0 {
		if err := assignInputField(input, "AssessmentTargets", _migrationhubstrategyAssessmentTargets); err != nil {
			log.Errorf("invalid --assessment-targets: %s", err.Error())
			return
		}
	}
	if len(_migrationhubstrategyS3bucketForAnalysisData) > 0 {
		input.S3bucketForAnalysisData = aws.String(_migrationhubstrategyS3bucketForAnalysisData)
	}
	if len(_migrationhubstrategyS3bucketForReportData) > 0 {
		input.S3bucketForReportData = aws.String(_migrationhubstrategyS3bucketForReportData)
	}

	if resp, err := client.StartAssessment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a file import.
func migrationhubstrategy_StartImportFileTask(cfg aws.Config, client *migrationhubstrategy.Client) {
	input := &migrationhubstrategy.StartImportFileTaskInput{
		// Name: *string, // Required
		// S3Bucket: *string, // Required
		// S3key: *string, // Required
	}

	if len(_migrationhubstrategyName) > 0 {
		input.Name = aws.String(_migrationhubstrategyName)
	}
	if len(_migrationhubstrategyS3Bucket) > 0 {
		input.S3Bucket = aws.String(_migrationhubstrategyS3Bucket)
	}
	if len(_migrationhubstrategyS3key) > 0 {
		input.S3key = aws.String(_migrationhubstrategyS3key)
	}
	if len(_migrationhubstrategyDataSourceType) > 0 {
		if err := assignInputField(input, "DataSourceType", _migrationhubstrategyDataSourceType); err != nil {
			log.Errorf("invalid --data-source-type: %s", err.Error())
			return
		}
	}
	if len(_migrationhubstrategyGroupId) > 0 {
		if err := assignInputField(input, "GroupId", _migrationhubstrategyGroupId); err != nil {
			log.Errorf("invalid --group-id: %s", err.Error())
			return
		}
	}
	if len(_migrationhubstrategyS3bucketForReportData) > 0 {
		input.S3bucketForReportData = aws.String(_migrationhubstrategyS3bucketForReportData)
	}

	if resp, err := client.StartImportFileTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts generating a recommendation report.
func migrationhubstrategy_StartRecommendationReportGeneration(cfg aws.Config, client *migrationhubstrategy.Client) {
	input := &migrationhubstrategy.StartRecommendationReportGenerationInput{}

	if len(_migrationhubstrategyGroupIdFilter) > 0 {
		if err := assignInputField(input, "GroupIdFilter", _migrationhubstrategyGroupIdFilter); err != nil {
			log.Errorf("invalid --group-id-filter: %s", err.Error())
			return
		}
	}
	if len(_migrationhubstrategyOutputFormat) > 0 {
		if err := assignInputField(input, "OutputFormat", _migrationhubstrategyOutputFormat); err != nil {
			log.Errorf("invalid --output-format: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartRecommendationReportGeneration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops the assessment of an on-premises environment.
func migrationhubstrategy_StopAssessment(cfg aws.Config, client *migrationhubstrategy.Client) {
	input := &migrationhubstrategy.StopAssessmentInput{
		// AssessmentId: *string, // Required
	}

	if len(_migrationhubstrategyAssessmentId) > 0 {
		input.AssessmentId = aws.String(_migrationhubstrategyAssessmentId)
	}

	if resp, err := client.StopAssessment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration of an application component.
func migrationhubstrategy_UpdateApplicationComponentConfig(cfg aws.Config, client *migrationhubstrategy.Client) {
	input := &migrationhubstrategy.UpdateApplicationComponentConfigInput{
		// ApplicationComponentId: *string, // Required
	}

	if len(_migrationhubstrategyApplicationComponentId) > 0 {
		input.ApplicationComponentId = aws.String(_migrationhubstrategyApplicationComponentId)
	}
	if len(_migrationhubstrategyAppType) > 0 {
		if err := assignInputField(input, "AppType", _migrationhubstrategyAppType); err != nil {
			log.Errorf("invalid --app-type: %s", err.Error())
			return
		}
	}
	if len(_migrationhubstrategyConfigureOnly) > 0 {
		if err := assignInputField(input, "ConfigureOnly", _migrationhubstrategyConfigureOnly); err != nil {
			log.Errorf("invalid --configure-only: %s", err.Error())
			return
		}
	}
	if len(_migrationhubstrategyInclusionStatus) > 0 {
		if err := assignInputField(input, "InclusionStatus", _migrationhubstrategyInclusionStatus); err != nil {
			log.Errorf("invalid --inclusion-status: %s", err.Error())
			return
		}
	}
	if len(_migrationhubstrategySecretsManagerKey) > 0 {
		input.SecretsManagerKey = aws.String(_migrationhubstrategySecretsManagerKey)
	}
	if len(_migrationhubstrategySourceCodeList) > 0 {
		if err := assignInputField(input, "SourceCodeList", _migrationhubstrategySourceCodeList); err != nil {
			log.Errorf("invalid --source-code-list: %s", err.Error())
			return
		}
	}
	if len(_migrationhubstrategyStrategyOption) > 0 {
		if err := assignInputField(input, "StrategyOption", _migrationhubstrategyStrategyOption); err != nil {
			log.Errorf("invalid --strategy-option: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateApplicationComponentConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration of the specified server.
func migrationhubstrategy_UpdateServerConfig(cfg aws.Config, client *migrationhubstrategy.Client) {
	input := &migrationhubstrategy.UpdateServerConfigInput{
		// ServerId: *string, // Required
	}

	if len(_migrationhubstrategyServerId) > 0 {
		input.ServerId = aws.String(_migrationhubstrategyServerId)
	}
	if len(_migrationhubstrategyStrategyOption) > 0 {
		if err := assignInputField(input, "StrategyOption", _migrationhubstrategyStrategyOption); err != nil {
			log.Errorf("invalid --strategy-option: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateServerConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_migrationhubstrategyCmd)
	_migrationhubstrategyCmd.Flags().SortFlags = false

	_migrationhubstrategyCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_migrationhubstrategyCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_migrationhubstrategyCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_migrationhubstrategyCmd.Flags().StringVarP(&_migrationhubstrategyAppType, "app-type", "", "", "App Type")
	_migrationhubstrategyCmd.Flags().StringVarP(&_migrationhubstrategyApplicationComponentCriteria, "application-component-criteria", "", "", "Application Component Criteria")
	_migrationhubstrategyCmd.Flags().StringVarP(&_migrationhubstrategyApplicationComponentId, "application-component-id", "", "", "Application Component ID")
	_migrationhubstrategyCmd.Flags().StringVarP(&_migrationhubstrategyApplicationMode, "application-mode", "", "", "Application Mode")
	_migrationhubstrategyCmd.Flags().StringVarP(&_migrationhubstrategyApplicationPreferences, "application-preferences", "", "", "Application Preferences")
	_migrationhubstrategyCmd.Flags().StringVarP(&_migrationhubstrategyAssessmentDataSourceType, "assessment-data-source-type", "", "", "Assessment Data Source Type")
	_migrationhubstrategyCmd.Flags().StringVarP(&_migrationhubstrategyAssessmentId, "assessment-id", "", "", "Assessment ID")
	_migrationhubstrategyCmd.Flags().StringVarP(&_migrationhubstrategyAssessmentTargets, "assessment-targets", "", "", "Assessment Targets")
	_migrationhubstrategyCmd.Flags().StringVarP(&_migrationhubstrategyConfigureOnly, "configure-only", "", "", "Configure Only")
	_migrationhubstrategyCmd.Flags().StringVarP(&_migrationhubstrategyDataSourceType, "data-source-type", "", "", "Data Source Type")
	_migrationhubstrategyCmd.Flags().StringVarP(&_migrationhubstrategyDatabasePreferences, "database-preferences", "", "", "Database Preferences")
	_migrationhubstrategyCmd.Flags().StringVarP(&_migrationhubstrategyFilterValue, "filter-value", "", "", "Filter Value")
	_migrationhubstrategyCmd.Flags().StringVarP(&_migrationhubstrategyGroupId, "group-id", "", "", "Group ID")
	_migrationhubstrategyCmd.Flags().StringVarP(&_migrationhubstrategyGroupIdFilter, "group-id-filter", "", "", "Group ID Filter")
	_migrationhubstrategyCmd.Flags().StringVarP(&_migrationhubstrategyId, "id", "", "", "ID")
	_migrationhubstrategyCmd.Flags().StringVarP(&_migrationhubstrategyInclusionStatus, "inclusion-status", "", "", "Inclusion Status")
	_migrationhubstrategyCmd.Flags().StringVarP(&_migrationhubstrategyMaxResults, "max-results", "", "", "Max Results")
	_migrationhubstrategyCmd.Flags().StringVarP(&_migrationhubstrategyName, "name", "", "", "Name")
	_migrationhubstrategyCmd.Flags().StringVarP(&_migrationhubstrategyNextToken, "next-token", "", "", "Next Token")
	_migrationhubstrategyCmd.Flags().StringVarP(&_migrationhubstrategyOutputFormat, "output-format", "", "", "Output Format")
	_migrationhubstrategyCmd.Flags().StringVarP(&_migrationhubstrategyPrioritizeBusinessGoals, "prioritize-business-goals", "", "", "Prioritize Business Goals")
	_migrationhubstrategyCmd.Flags().StringVarP(&_migrationhubstrategyS3Bucket, "s3-bucket", "", "", "S3 Bucket")
	_migrationhubstrategyCmd.Flags().StringVarP(&_migrationhubstrategyS3bucketForAnalysisData, "s3bucket-for-analysis-data", "", "", "S3bucket For Analysis Data")
	_migrationhubstrategyCmd.Flags().StringVarP(&_migrationhubstrategyS3bucketForReportData, "s3bucket-for-report-data", "", "", "S3bucket For Report Data")
	_migrationhubstrategyCmd.Flags().StringVarP(&_migrationhubstrategyS3key, "s3key", "", "", "S3key")
	_migrationhubstrategyCmd.Flags().StringVarP(&_migrationhubstrategySecretsManagerKey, "secrets-manager-key", "", "", "Secrets Manager Key")
	_migrationhubstrategyCmd.Flags().StringVarP(&_migrationhubstrategyServerCriteria, "server-criteria", "", "", "Server Criteria")
	_migrationhubstrategyCmd.Flags().StringVarP(&_migrationhubstrategyServerId, "server-id", "", "", "Server ID")
	_migrationhubstrategyCmd.Flags().StringVarP(&_migrationhubstrategySort, "sort", "", "", "Sort")
	_migrationhubstrategyCmd.Flags().StringVarP(&_migrationhubstrategySourceCodeList, "source-code-list", "", "", "Source Code List")
	_migrationhubstrategyCmd.Flags().StringVarP(&_migrationhubstrategyStrategyOption, "strategy-option", "", "", "Strategy Option")

	_migrationhubstrategyCmd.Flags().BoolVarP(&_migrationhubstrategyGetApplicationComponentDetails, "get-application-component-details", "", false, "Get Application Component Details")
	_migrationhubstrategyCmd.Flags().BoolVarP(&_migrationhubstrategyGetApplicationComponentStrategies, "get-application-component-strategies", "", false, "Get Application Component Strategies")
	_migrationhubstrategyCmd.Flags().BoolVarP(&_migrationhubstrategyGetAssessment, "get-assessment", "", false, "Get Assessment")
	_migrationhubstrategyCmd.Flags().BoolVarP(&_migrationhubstrategyGetImportFileTask, "get-import-file-task", "", false, "Get Import File Task")
	_migrationhubstrategyCmd.Flags().BoolVarP(&_migrationhubstrategyGetLatestAssessmentId, "get-latest-assessment-id", "", false, "Get Latest Assessment ID")
	_migrationhubstrategyCmd.Flags().BoolVarP(&_migrationhubstrategyGetPortfolioPreferences, "get-portfolio-preferences", "", false, "Get Portfolio Preferences")
	_migrationhubstrategyCmd.Flags().BoolVarP(&_migrationhubstrategyGetPortfolioSummary, "get-portfolio-summary", "", false, "Get Portfolio Summary")
	_migrationhubstrategyCmd.Flags().BoolVarP(&_migrationhubstrategyGetRecommendationReportDetails, "get-recommendation-report-details", "", false, "Get Recommendation Report Details")
	_migrationhubstrategyCmd.Flags().BoolVarP(&_migrationhubstrategyGetServerDetails, "get-server-details", "", false, "Get Server Details")
	_migrationhubstrategyCmd.Flags().BoolVarP(&_migrationhubstrategyGetServerStrategies, "get-server-strategies", "", false, "Get Server Strategies")
	_migrationhubstrategyCmd.Flags().BoolVarP(&_migrationhubstrategyListAnalyzableServers, "list-analyzable-servers", "", false, "List Analyzable Servers")
	_migrationhubstrategyCmd.Flags().BoolVarP(&_migrationhubstrategyListApplicationComponents, "list-application-components", "", false, "List Application Components")
	_migrationhubstrategyCmd.Flags().BoolVarP(&_migrationhubstrategyListCollectors, "list-collectors", "", false, "List Collectors")
	_migrationhubstrategyCmd.Flags().BoolVarP(&_migrationhubstrategyListImportFileTask, "list-import-file-task", "", false, "List Import File Task")
	_migrationhubstrategyCmd.Flags().BoolVarP(&_migrationhubstrategyListServers, "list-servers", "", false, "List Servers")
	_migrationhubstrategyCmd.Flags().BoolVarP(&_migrationhubstrategyPutPortfolioPreferences, "put-portfolio-preferences", "", false, "Put Portfolio Preferences")
	_migrationhubstrategyCmd.Flags().BoolVarP(&_migrationhubstrategyStartAssessment, "start-assessment", "", false, "Start Assessment")
	_migrationhubstrategyCmd.Flags().BoolVarP(&_migrationhubstrategyStartImportFileTask, "start-import-file-task", "", false, "Start Import File Task")
	_migrationhubstrategyCmd.Flags().BoolVarP(&_migrationhubstrategyStartRecommendationReportGeneration, "start-recommendation-report-generation", "", false, "Start Recommendation Report Generation")
	_migrationhubstrategyCmd.Flags().BoolVarP(&_migrationhubstrategyStopAssessment, "stop-assessment", "", false, "Stop Assessment")
	_migrationhubstrategyCmd.Flags().BoolVarP(&_migrationhubstrategyUpdateApplicationComponentConfig, "update-application-component-config", "", false, "Update Application Component Config")
	_migrationhubstrategyCmd.Flags().BoolVarP(&_migrationhubstrategyUpdateServerConfig, "update-server-config", "", false, "Update Server Config")

}
