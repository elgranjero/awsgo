package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/healthlake"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// healthlakeCmd represents the healthlake command
var _healthlakeCmd = &cobra.Command{
	Use:   "healthlake",
	Short: "AWS healthlake CLI",
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
		client := healthlake.NewFromConfig(cfg)
		if _healthlakeCreateFHIRDatastore {
			healthlake_CreateFHIRDatastore(cfg, client)
			return
		}
		if _healthlakeDeleteFHIRDatastore {
			healthlake_DeleteFHIRDatastore(cfg, client)
			return
		}
		if _healthlakeDescribeFHIRDatastore {
			healthlake_DescribeFHIRDatastore(cfg, client)
			return
		}
		if _healthlakeDescribeFHIRExportJob {
			healthlake_DescribeFHIRExportJob(cfg, client)
			return
		}
		if _healthlakeDescribeFHIRImportJob {
			healthlake_DescribeFHIRImportJob(cfg, client)
			return
		}
		if _healthlakeListFHIRDatastores {
			healthlake_ListFHIRDatastores(cfg, client)
			return
		}
		if _healthlakeListFHIRExportJobs {
			healthlake_ListFHIRExportJobs(cfg, client)
			return
		}
		if _healthlakeListFHIRImportJobs {
			healthlake_ListFHIRImportJobs(cfg, client)
			return
		}
		if _healthlakeListTagsForResource {
			healthlake_ListTagsForResource(cfg, client)
			return
		}
		if _healthlakeStartFHIRExportJob {
			healthlake_StartFHIRExportJob(cfg, client)
			return
		}
		if _healthlakeStartFHIRImportJob {
			healthlake_StartFHIRImportJob(cfg, client)
			return
		}
		if _healthlakeTagResource {
			healthlake_TagResource(cfg, client)
			return
		}
		if _healthlakeUntagResource {
			healthlake_UntagResource(cfg, client)
			return
		}

	},
}

var (
	_healthlakeCreateFHIRDatastore   bool
	_healthlakeDeleteFHIRDatastore   bool
	_healthlakeDescribeFHIRDatastore bool
	_healthlakeDescribeFHIRExportJob bool
	_healthlakeDescribeFHIRImportJob bool
	_healthlakeListFHIRDatastores    bool
	_healthlakeListFHIRExportJobs    bool
	_healthlakeListFHIRImportJobs    bool
	_healthlakeListTagsForResource   bool
	_healthlakeStartFHIRExportJob    bool
	_healthlakeStartFHIRImportJob    bool
	_healthlakeTagResource           bool
	_healthlakeUntagResource         bool

	_healthlakeClientToken                   string
	_healthlakeDataAccessRoleArn             string
	_healthlakeDatastoreId                   string
	_healthlakeDatastoreName                 string
	_healthlakeDatastoreTypeVersion          string
	_healthlakeFilter                        string
	_healthlakeIdentityProviderConfiguration string
	_healthlakeInputDataConfig               string
	_healthlakeJobId                         string
	_healthlakeJobName                       string
	_healthlakeJobOutputDataConfig           string
	_healthlakeJobStatus                     string
	_healthlakeMaxResults                    string
	_healthlakeNextToken                     string
	_healthlakeOutputDataConfig              string
	_healthlakePreloadDataConfig             string
	_healthlakeResourceARN                   string
	_healthlakeSseConfiguration              string
	_healthlakeSubmittedAfter                string
	_healthlakeSubmittedBefore               string
	_healthlakeTagKeys                       []string
	_healthlakeTags                          string
	_healthlakeValidationLevel               string
)

// Create a FHIR-enabled data store.
func healthlake_CreateFHIRDatastore(cfg aws.Config, client *healthlake.Client) {
	input := &healthlake.CreateFHIRDatastoreInput{
		// DatastoreTypeVersion: types.FHIRVersion, // Required
	}

	if len(_healthlakeDatastoreTypeVersion) > 0 {
		if err := assignInputField(input, "DatastoreTypeVersion", _healthlakeDatastoreTypeVersion); err != nil {
			log.Errorf("invalid --datastore-type-version: %s", err.Error())
			return
		}
	}
	if len(_healthlakeClientToken) > 0 {
		input.ClientToken = aws.String(_healthlakeClientToken)
	}
	if len(_healthlakeDatastoreName) > 0 {
		input.DatastoreName = aws.String(_healthlakeDatastoreName)
	}
	if len(_healthlakeIdentityProviderConfiguration) > 0 {
		if err := assignInputField(input, "IdentityProviderConfiguration", _healthlakeIdentityProviderConfiguration); err != nil {
			log.Errorf("invalid --identity-provider-configuration: %s", err.Error())
			return
		}
	}
	if len(_healthlakePreloadDataConfig) > 0 {
		if err := assignInputField(input, "PreloadDataConfig", _healthlakePreloadDataConfig); err != nil {
			log.Errorf("invalid --preload-data-config: %s", err.Error())
			return
		}
	}
	if len(_healthlakeSseConfiguration) > 0 {
		if err := assignInputField(input, "SseConfiguration", _healthlakeSseConfiguration); err != nil {
			log.Errorf("invalid --sse-configuration: %s", err.Error())
			return
		}
	}
	if len(_healthlakeTags) > 0 {
		if err := assignInputField(input, "Tags", _healthlakeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateFHIRDatastore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a FHIR-enabled data store.
func healthlake_DeleteFHIRDatastore(cfg aws.Config, client *healthlake.Client) {
	input := &healthlake.DeleteFHIRDatastoreInput{
		// DatastoreId: *string, // Required
	}

	if len(_healthlakeDatastoreId) > 0 {
		input.DatastoreId = aws.String(_healthlakeDatastoreId)
	}

	if resp, err := client.DeleteFHIRDatastore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get properties for a FHIR-enabled data store.
func healthlake_DescribeFHIRDatastore(cfg aws.Config, client *healthlake.Client) {
	input := &healthlake.DescribeFHIRDatastoreInput{
		// DatastoreId: *string, // Required
	}

	if len(_healthlakeDatastoreId) > 0 {
		input.DatastoreId = aws.String(_healthlakeDatastoreId)
	}

	if resp, err := client.DescribeFHIRDatastore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get FHIR export job properties.
func healthlake_DescribeFHIRExportJob(cfg aws.Config, client *healthlake.Client) {
	input := &healthlake.DescribeFHIRExportJobInput{
		// DatastoreId: *string, // Required
		// JobId: *string, // Required
	}

	if len(_healthlakeDatastoreId) > 0 {
		input.DatastoreId = aws.String(_healthlakeDatastoreId)
	}
	if len(_healthlakeJobId) > 0 {
		input.JobId = aws.String(_healthlakeJobId)
	}

	if resp, err := client.DescribeFHIRExportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the import job properties to learn more about the job or job progress.
func healthlake_DescribeFHIRImportJob(cfg aws.Config, client *healthlake.Client) {
	input := &healthlake.DescribeFHIRImportJobInput{
		// DatastoreId: *string, // Required
		// JobId: *string, // Required
	}

	if len(_healthlakeDatastoreId) > 0 {
		input.DatastoreId = aws.String(_healthlakeDatastoreId)
	}
	if len(_healthlakeJobId) > 0 {
		input.JobId = aws.String(_healthlakeJobId)
	}

	if resp, err := client.DescribeFHIRImportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List all FHIR-enabled data stores in a user’s account, regardless of data store
// status.
func healthlake_ListFHIRDatastores(cfg aws.Config, client *healthlake.Client) {
	input := &healthlake.ListFHIRDatastoresInput{}

	if len(_healthlakeFilter) > 0 {
		if err := assignInputField(input, "Filter", _healthlakeFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_healthlakeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _healthlakeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_healthlakeNextToken) > 0 {
		input.NextToken = aws.String(_healthlakeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListFHIRDatastores(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*healthlake.ListFHIRDatastoresOutput
	p := healthlake.NewListFHIRDatastoresPaginator(client, input)
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

// Lists all FHIR export jobs associated with an account and their statuses.
func healthlake_ListFHIRExportJobs(cfg aws.Config, client *healthlake.Client) {
	input := &healthlake.ListFHIRExportJobsInput{
		// DatastoreId: *string, // Required
	}

	if len(_healthlakeDatastoreId) > 0 {
		input.DatastoreId = aws.String(_healthlakeDatastoreId)
	}
	if len(_healthlakeJobName) > 0 {
		input.JobName = aws.String(_healthlakeJobName)
	}
	if len(_healthlakeJobStatus) > 0 {
		if err := assignInputField(input, "JobStatus", _healthlakeJobStatus); err != nil {
			log.Errorf("invalid --job-status: %s", err.Error())
			return
		}
	}
	if len(_healthlakeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _healthlakeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_healthlakeNextToken) > 0 {
		input.NextToken = aws.String(_healthlakeNextToken)
	}
	if len(_healthlakeSubmittedAfter) > 0 {
		if err := assignInputField(input, "SubmittedAfter", _healthlakeSubmittedAfter); err != nil {
			log.Errorf("invalid --submitted-after: %s", err.Error())
			return
		}
	}
	if len(_healthlakeSubmittedBefore) > 0 {
		if err := assignInputField(input, "SubmittedBefore", _healthlakeSubmittedBefore); err != nil {
			log.Errorf("invalid --submitted-before: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListFHIRExportJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*healthlake.ListFHIRExportJobsOutput
	p := healthlake.NewListFHIRExportJobsPaginator(client, input)
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

// List all FHIR import jobs associated with an account and their statuses.
func healthlake_ListFHIRImportJobs(cfg aws.Config, client *healthlake.Client) {
	input := &healthlake.ListFHIRImportJobsInput{
		// DatastoreId: *string, // Required
	}

	if len(_healthlakeDatastoreId) > 0 {
		input.DatastoreId = aws.String(_healthlakeDatastoreId)
	}
	if len(_healthlakeJobName) > 0 {
		input.JobName = aws.String(_healthlakeJobName)
	}
	if len(_healthlakeJobStatus) > 0 {
		if err := assignInputField(input, "JobStatus", _healthlakeJobStatus); err != nil {
			log.Errorf("invalid --job-status: %s", err.Error())
			return
		}
	}
	if len(_healthlakeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _healthlakeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_healthlakeNextToken) > 0 {
		input.NextToken = aws.String(_healthlakeNextToken)
	}
	if len(_healthlakeSubmittedAfter) > 0 {
		if err := assignInputField(input, "SubmittedAfter", _healthlakeSubmittedAfter); err != nil {
			log.Errorf("invalid --submitted-after: %s", err.Error())
			return
		}
	}
	if len(_healthlakeSubmittedBefore) > 0 {
		if err := assignInputField(input, "SubmittedBefore", _healthlakeSubmittedBefore); err != nil {
			log.Errorf("invalid --submitted-before: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListFHIRImportJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*healthlake.ListFHIRImportJobsOutput
	p := healthlake.NewListFHIRImportJobsPaginator(client, input)
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

// Returns a list of all existing tags associated with a data store.
func healthlake_ListTagsForResource(cfg aws.Config, client *healthlake.Client) {
	input := &healthlake.ListTagsForResourceInput{
		// ResourceARN: *string, // Required
	}

	if len(_healthlakeResourceARN) > 0 {
		input.ResourceARN = aws.String(_healthlakeResourceARN)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Start a FHIR export job.
func healthlake_StartFHIRExportJob(cfg aws.Config, client *healthlake.Client) {
	input := &healthlake.StartFHIRExportJobInput{
		// DataAccessRoleArn: *string, // Required
		// DatastoreId: *string, // Required
		// OutputDataConfig: types.OutputDataConfig, // Required
	}

	if len(_healthlakeDataAccessRoleArn) > 0 {
		input.DataAccessRoleArn = aws.String(_healthlakeDataAccessRoleArn)
	}
	if len(_healthlakeDatastoreId) > 0 {
		input.DatastoreId = aws.String(_healthlakeDatastoreId)
	}
	if len(_healthlakeOutputDataConfig) > 0 {
		if err := assignInputField(input, "OutputDataConfig", _healthlakeOutputDataConfig); err != nil {
			log.Errorf("invalid --output-data-config: %s", err.Error())
			return
		}
	}
	if len(_healthlakeClientToken) > 0 {
		input.ClientToken = aws.String(_healthlakeClientToken)
	}
	if len(_healthlakeJobName) > 0 {
		input.JobName = aws.String(_healthlakeJobName)
	}

	if resp, err := client.StartFHIRExportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Start importing bulk FHIR data into an ACTIVE data store. The import job
// imports FHIR data found in the InputDataConfig object and stores processing
// results in the JobOutputDataConfig object.
func healthlake_StartFHIRImportJob(cfg aws.Config, client *healthlake.Client) {
	input := &healthlake.StartFHIRImportJobInput{
		// DataAccessRoleArn: *string, // Required
		// DatastoreId: *string, // Required
		// InputDataConfig: types.InputDataConfig, // Required
		// JobOutputDataConfig: types.OutputDataConfig, // Required
	}

	if len(_healthlakeDataAccessRoleArn) > 0 {
		input.DataAccessRoleArn = aws.String(_healthlakeDataAccessRoleArn)
	}
	if len(_healthlakeDatastoreId) > 0 {
		input.DatastoreId = aws.String(_healthlakeDatastoreId)
	}
	if len(_healthlakeInputDataConfig) > 0 {
		if err := assignInputField(input, "InputDataConfig", _healthlakeInputDataConfig); err != nil {
			log.Errorf("invalid --input-data-config: %s", err.Error())
			return
		}
	}
	if len(_healthlakeJobOutputDataConfig) > 0 {
		if err := assignInputField(input, "JobOutputDataConfig", _healthlakeJobOutputDataConfig); err != nil {
			log.Errorf("invalid --job-output-data-config: %s", err.Error())
			return
		}
	}
	if len(_healthlakeClientToken) > 0 {
		input.ClientToken = aws.String(_healthlakeClientToken)
	}
	if len(_healthlakeJobName) > 0 {
		input.JobName = aws.String(_healthlakeJobName)
	}
	if len(_healthlakeValidationLevel) > 0 {
		if err := assignInputField(input, "ValidationLevel", _healthlakeValidationLevel); err != nil {
			log.Errorf("invalid --validation-level: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartFHIRImportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Add a user-specifed key and value tag to a data store.
func healthlake_TagResource(cfg aws.Config, client *healthlake.Client) {
	input := &healthlake.TagResourceInput{
		// ResourceARN: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_healthlakeResourceARN) > 0 {
		input.ResourceARN = aws.String(_healthlakeResourceARN)
	}
	if len(_healthlakeTags) > 0 {
		if err := assignInputField(input, "Tags", _healthlakeTags); err != nil {
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

// Remove a user-specifed key and value tag from a data store.
func healthlake_UntagResource(cfg aws.Config, client *healthlake.Client) {
	input := &healthlake.UntagResourceInput{
		// ResourceARN: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_healthlakeResourceARN) > 0 {
		input.ResourceARN = aws.String(_healthlakeResourceARN)
	}
	if len(_healthlakeTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _healthlakeTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_healthlakeCmd)
	_healthlakeCmd.Flags().SortFlags = false

	_healthlakeCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_healthlakeCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_healthlakeCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_healthlakeCmd.Flags().StringVarP(&_healthlakeClientToken, "client-token", "", "", "Client Token")
	_healthlakeCmd.Flags().StringVarP(&_healthlakeDataAccessRoleArn, "data-access-role-arn", "", "", "Data Access Role ARN")
	_healthlakeCmd.Flags().StringVarP(&_healthlakeDatastoreId, "datastore-id", "", "", "Datastore ID")
	_healthlakeCmd.Flags().StringVarP(&_healthlakeDatastoreName, "datastore-name", "", "", "Datastore Name")
	_healthlakeCmd.Flags().StringVarP(&_healthlakeDatastoreTypeVersion, "datastore-type-version", "", "", "Datastore Type Version")
	_healthlakeCmd.Flags().StringVarP(&_healthlakeFilter, "filter", "", "", "Filter")
	_healthlakeCmd.Flags().StringVarP(&_healthlakeIdentityProviderConfiguration, "identity-provider-configuration", "", "", "Identity Provider Configuration")
	_healthlakeCmd.Flags().StringVarP(&_healthlakeInputDataConfig, "input-data-config", "", "", "Input Data Config")
	_healthlakeCmd.Flags().StringVarP(&_healthlakeJobId, "job-id", "", "", "Job ID")
	_healthlakeCmd.Flags().StringVarP(&_healthlakeJobName, "job-name", "", "", "Job Name")
	_healthlakeCmd.Flags().StringVarP(&_healthlakeJobOutputDataConfig, "job-output-data-config", "", "", "Job Output Data Config")
	_healthlakeCmd.Flags().StringVarP(&_healthlakeJobStatus, "job-status", "", "", "Job Status")
	_healthlakeCmd.Flags().StringVarP(&_healthlakeMaxResults, "max-results", "", "", "Max Results")
	_healthlakeCmd.Flags().StringVarP(&_healthlakeNextToken, "next-token", "", "", "Next Token")
	_healthlakeCmd.Flags().StringVarP(&_healthlakeOutputDataConfig, "output-data-config", "", "", "Output Data Config")
	_healthlakeCmd.Flags().StringVarP(&_healthlakePreloadDataConfig, "preload-data-config", "", "", "Preload Data Config")
	_healthlakeCmd.Flags().StringVarP(&_healthlakeResourceARN, "resource-arn", "", "", "Resource ARN")
	_healthlakeCmd.Flags().StringVarP(&_healthlakeSseConfiguration, "sse-configuration", "", "", "SSE Configuration")
	_healthlakeCmd.Flags().StringVarP(&_healthlakeSubmittedAfter, "submitted-after", "", "", "Submitted After")
	_healthlakeCmd.Flags().StringVarP(&_healthlakeSubmittedBefore, "submitted-before", "", "", "Submitted Before")
	_healthlakeCmd.Flags().StringSliceVarP(&_healthlakeTagKeys, "tag-keys", "", nil, "Tag Keys")
	_healthlakeCmd.Flags().StringVarP(&_healthlakeTags, "tags", "", "", "Tags")
	_healthlakeCmd.Flags().StringVarP(&_healthlakeValidationLevel, "validation-level", "", "", "Validation Level")

	_healthlakeCmd.Flags().BoolVarP(&_healthlakeCreateFHIRDatastore, "create-fhir-datastore", "", false, "Create Fhir Datastore")
	_healthlakeCmd.Flags().BoolVarP(&_healthlakeDeleteFHIRDatastore, "delete-fhir-datastore", "", false, "Delete Fhir Datastore")
	_healthlakeCmd.Flags().BoolVarP(&_healthlakeDescribeFHIRDatastore, "describe-fhir-datastore", "", false, "Describe Fhir Datastore")
	_healthlakeCmd.Flags().BoolVarP(&_healthlakeDescribeFHIRExportJob, "describe-fhir-export-job", "", false, "Describe Fhir Export Job")
	_healthlakeCmd.Flags().BoolVarP(&_healthlakeDescribeFHIRImportJob, "describe-fhir-import-job", "", false, "Describe Fhir Import Job")
	_healthlakeCmd.Flags().BoolVarP(&_healthlakeListFHIRDatastores, "list-fhir-datastores", "", false, "List Fhir Datastores")
	_healthlakeCmd.Flags().BoolVarP(&_healthlakeListFHIRExportJobs, "list-fhir-export-jobs", "", false, "List Fhir Export Jobs")
	_healthlakeCmd.Flags().BoolVarP(&_healthlakeListFHIRImportJobs, "list-fhir-import-jobs", "", false, "List Fhir Import Jobs")
	_healthlakeCmd.Flags().BoolVarP(&_healthlakeListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_healthlakeCmd.Flags().BoolVarP(&_healthlakeStartFHIRExportJob, "start-fhir-export-job", "", false, "Start Fhir Export Job")
	_healthlakeCmd.Flags().BoolVarP(&_healthlakeStartFHIRImportJob, "start-fhir-import-job", "", false, "Start Fhir Import Job")
	_healthlakeCmd.Flags().BoolVarP(&_healthlakeTagResource, "tag-resource", "", false, "Tag Resource")
	_healthlakeCmd.Flags().BoolVarP(&_healthlakeUntagResource, "untag-resource", "", false, "Untag Resource")

}
