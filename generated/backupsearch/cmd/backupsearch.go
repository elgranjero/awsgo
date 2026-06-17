package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/backupsearch"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// backupsearchCmd represents the backupsearch command
var _backupsearchCmd = &cobra.Command{
	Use:   "backupsearch",
	Short: "AWS backupsearch CLI",
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
		client := backupsearch.NewFromConfig(cfg)
		if _backupsearchGetSearchJob {
			backupsearch_GetSearchJob(cfg, client)
			return
		}
		if _backupsearchGetSearchResultExportJob {
			backupsearch_GetSearchResultExportJob(cfg, client)
			return
		}
		if _backupsearchListSearchJobBackups {
			backupsearch_ListSearchJobBackups(cfg, client)
			return
		}
		if _backupsearchListSearchJobResults {
			backupsearch_ListSearchJobResults(cfg, client)
			return
		}
		if _backupsearchListSearchJobs {
			backupsearch_ListSearchJobs(cfg, client)
			return
		}
		if _backupsearchListSearchResultExportJobs {
			backupsearch_ListSearchResultExportJobs(cfg, client)
			return
		}
		if _backupsearchListTagsForResource {
			backupsearch_ListTagsForResource(cfg, client)
			return
		}
		if _backupsearchStartSearchJob {
			backupsearch_StartSearchJob(cfg, client)
			return
		}
		if _backupsearchStartSearchResultExportJob {
			backupsearch_StartSearchResultExportJob(cfg, client)
			return
		}
		if _backupsearchStopSearchJob {
			backupsearch_StopSearchJob(cfg, client)
			return
		}
		if _backupsearchTagResource {
			backupsearch_TagResource(cfg, client)
			return
		}
		if _backupsearchUntagResource {
			backupsearch_UntagResource(cfg, client)
			return
		}

	},
}

var (
	_backupsearchGetSearchJob               bool
	_backupsearchGetSearchResultExportJob   bool
	_backupsearchListSearchJobBackups       bool
	_backupsearchListSearchJobResults       bool
	_backupsearchListSearchJobs             bool
	_backupsearchListSearchResultExportJobs bool
	_backupsearchListTagsForResource        bool
	_backupsearchStartSearchJob             bool
	_backupsearchStartSearchResultExportJob bool
	_backupsearchStopSearchJob              bool
	_backupsearchTagResource                bool
	_backupsearchUntagResource              bool

	_backupsearchByStatus            string
	_backupsearchClientToken         string
	_backupsearchEncryptionKeyArn    string
	_backupsearchExportJobIdentifier string
	_backupsearchExportSpecification string
	_backupsearchItemFilters         string
	_backupsearchMaxResults          string
	_backupsearchName                string
	_backupsearchNextToken           string
	_backupsearchResourceArn         string
	_backupsearchRoleArn             string
	_backupsearchSearchJobIdentifier string
	_backupsearchSearchScope         string
	_backupsearchStatus              string
	_backupsearchTagKeys             []string
	_backupsearchTags                string
)

// This operation retrieves metadata of a search job, including its progress.
func backupsearch_GetSearchJob(cfg aws.Config, client *backupsearch.Client) {
	input := &backupsearch.GetSearchJobInput{
		// SearchJobIdentifier: *string, // Required
	}

	if len(_backupsearchSearchJobIdentifier) > 0 {
		input.SearchJobIdentifier = aws.String(_backupsearchSearchJobIdentifier)
	}

	if resp, err := client.GetSearchJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation retrieves the metadata of an export job.
// An export job is an operation that transmits the results of a search job to a
// specified S3 bucket in a .csv file.
//
// An export job allows you to retain results of a search beyond the search job's
// scheduled retention of 7 days.
func backupsearch_GetSearchResultExportJob(cfg aws.Config, client *backupsearch.Client) {
	input := &backupsearch.GetSearchResultExportJobInput{
		// ExportJobIdentifier: *string, // Required
	}

	if len(_backupsearchExportJobIdentifier) > 0 {
		input.ExportJobIdentifier = aws.String(_backupsearchExportJobIdentifier)
	}

	if resp, err := client.GetSearchResultExportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation returns a list of all backups (recovery points) in a paginated
// format that were included in the search job.
//
// If a search does not display an expected backup in the results, you can call
// this operation to display each backup included in the search. Any backups that
// were not included because they have a FAILED status from a permissions issue
// will be displayed, along with a status message.
//
// Only recovery points with a backup index that has a status of ACTIVE will be
// included in search results. If the index has any other status, its status will
// be displayed along with a status message.
func backupsearch_ListSearchJobBackups(cfg aws.Config, client *backupsearch.Client) {
	input := &backupsearch.ListSearchJobBackupsInput{
		// SearchJobIdentifier: *string, // Required
	}

	if len(_backupsearchSearchJobIdentifier) > 0 {
		input.SearchJobIdentifier = aws.String(_backupsearchSearchJobIdentifier)
	}
	if len(_backupsearchMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _backupsearchMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_backupsearchNextToken) > 0 {
		input.NextToken = aws.String(_backupsearchNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSearchJobBackups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*backupsearch.ListSearchJobBackupsOutput
	p := backupsearch.NewListSearchJobBackupsPaginator(client, input)
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

// This operation returns a list of a specified search job.
func backupsearch_ListSearchJobResults(cfg aws.Config, client *backupsearch.Client) {
	input := &backupsearch.ListSearchJobResultsInput{
		// SearchJobIdentifier: *string, // Required
	}

	if len(_backupsearchSearchJobIdentifier) > 0 {
		input.SearchJobIdentifier = aws.String(_backupsearchSearchJobIdentifier)
	}
	if len(_backupsearchMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _backupsearchMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_backupsearchNextToken) > 0 {
		input.NextToken = aws.String(_backupsearchNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSearchJobResults(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*backupsearch.ListSearchJobResultsOutput
	p := backupsearch.NewListSearchJobResultsPaginator(client, input)
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

// This operation returns a list of search jobs belonging to an account.
func backupsearch_ListSearchJobs(cfg aws.Config, client *backupsearch.Client) {
	input := &backupsearch.ListSearchJobsInput{}

	if len(_backupsearchByStatus) > 0 {
		if err := assignInputField(input, "ByStatus", _backupsearchByStatus); err != nil {
			log.Errorf("invalid --by-status: %s", err.Error())
			return
		}
	}
	if len(_backupsearchMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _backupsearchMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_backupsearchNextToken) > 0 {
		input.NextToken = aws.String(_backupsearchNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSearchJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*backupsearch.ListSearchJobsOutput
	p := backupsearch.NewListSearchJobsPaginator(client, input)
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

// This operation exports search results of a search job to a specified
// destination S3 bucket.
func backupsearch_ListSearchResultExportJobs(cfg aws.Config, client *backupsearch.Client) {
	input := &backupsearch.ListSearchResultExportJobsInput{}

	if len(_backupsearchMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _backupsearchMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_backupsearchNextToken) > 0 {
		input.NextToken = aws.String(_backupsearchNextToken)
	}
	if len(_backupsearchSearchJobIdentifier) > 0 {
		input.SearchJobIdentifier = aws.String(_backupsearchSearchJobIdentifier)
	}
	if len(_backupsearchStatus) > 0 {
		if err := assignInputField(input, "Status", _backupsearchStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListSearchResultExportJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*backupsearch.ListSearchResultExportJobsOutput
	p := backupsearch.NewListSearchResultExportJobsPaginator(client, input)
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

// This operation returns the tags for a resource type.
func backupsearch_ListTagsForResource(cfg aws.Config, client *backupsearch.Client) {
	input := &backupsearch.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_backupsearchResourceArn) > 0 {
		input.ResourceArn = aws.String(_backupsearchResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation creates a search job which returns recovery points filtered by
// SearchScope and items filtered by ItemFilters.
//
// You can optionally include ClientToken, EncryptionKeyArn, Name, and/or Tags.
func backupsearch_StartSearchJob(cfg aws.Config, client *backupsearch.Client) {
	input := &backupsearch.StartSearchJobInput{
		// SearchScope: *types.SearchScope, // Required
	}

	if len(_backupsearchSearchScope) > 0 {
		if err := assignInputField(input, "SearchScope", _backupsearchSearchScope); err != nil {
			log.Errorf("invalid --search-scope: %s", err.Error())
			return
		}
	}
	if len(_backupsearchClientToken) > 0 {
		input.ClientToken = aws.String(_backupsearchClientToken)
	}
	if len(_backupsearchEncryptionKeyArn) > 0 {
		input.EncryptionKeyArn = aws.String(_backupsearchEncryptionKeyArn)
	}
	if len(_backupsearchItemFilters) > 0 {
		if err := assignInputField(input, "ItemFilters", _backupsearchItemFilters); err != nil {
			log.Errorf("invalid --item-filters: %s", err.Error())
			return
		}
	}
	if len(_backupsearchName) > 0 {
		input.Name = aws.String(_backupsearchName)
	}
	if len(_backupsearchTags) > 0 {
		if err := assignInputField(input, "Tags", _backupsearchTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartSearchJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operations starts a job to export the results of search job to a
// designated S3 bucket.
func backupsearch_StartSearchResultExportJob(cfg aws.Config, client *backupsearch.Client) {
	input := &backupsearch.StartSearchResultExportJobInput{
		// ExportSpecification: types.ExportSpecification, // Required
		// SearchJobIdentifier: *string, // Required
	}

	if len(_backupsearchExportSpecification) > 0 {
		if err := assignInputField(input, "ExportSpecification", _backupsearchExportSpecification); err != nil {
			log.Errorf("invalid --export-specification: %s", err.Error())
			return
		}
	}
	if len(_backupsearchSearchJobIdentifier) > 0 {
		input.SearchJobIdentifier = aws.String(_backupsearchSearchJobIdentifier)
	}
	if len(_backupsearchClientToken) > 0 {
		input.ClientToken = aws.String(_backupsearchClientToken)
	}
	if len(_backupsearchRoleArn) > 0 {
		input.RoleArn = aws.String(_backupsearchRoleArn)
	}
	if len(_backupsearchTags) > 0 {
		if err := assignInputField(input, "Tags", _backupsearchTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartSearchResultExportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operations ends a search job.
// Only a search job with a status of RUNNING can be stopped.
func backupsearch_StopSearchJob(cfg aws.Config, client *backupsearch.Client) {
	input := &backupsearch.StopSearchJobInput{
		// SearchJobIdentifier: *string, // Required
	}

	if len(_backupsearchSearchJobIdentifier) > 0 {
		input.SearchJobIdentifier = aws.String(_backupsearchSearchJobIdentifier)
	}

	if resp, err := client.StopSearchJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation puts tags on the resource you indicate.
func backupsearch_TagResource(cfg aws.Config, client *backupsearch.Client) {
	input := &backupsearch.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]*string, // Required
	}

	if len(_backupsearchResourceArn) > 0 {
		input.ResourceArn = aws.String(_backupsearchResourceArn)
	}
	if len(_backupsearchTags) > 0 {
		if err := assignInputField(input, "Tags", _backupsearchTags); err != nil {
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

// This operation removes tags from the specified resource.
func backupsearch_UntagResource(cfg aws.Config, client *backupsearch.Client) {
	input := &backupsearch.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_backupsearchResourceArn) > 0 {
		input.ResourceArn = aws.String(_backupsearchResourceArn)
	}
	if len(_backupsearchTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _backupsearchTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_backupsearchCmd)
	_backupsearchCmd.Flags().SortFlags = false

	_backupsearchCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_backupsearchCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_backupsearchCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_backupsearchCmd.Flags().StringVarP(&_backupsearchByStatus, "by-status", "", "", "By Status")
	_backupsearchCmd.Flags().StringVarP(&_backupsearchClientToken, "client-token", "", "", "Client Token")
	_backupsearchCmd.Flags().StringVarP(&_backupsearchEncryptionKeyArn, "encryption-key-arn", "", "", "Encryption Key ARN")
	_backupsearchCmd.Flags().StringVarP(&_backupsearchExportJobIdentifier, "export-job-identifier", "", "", "Export Job Identifier")
	_backupsearchCmd.Flags().StringVarP(&_backupsearchExportSpecification, "export-specification", "", "", "Export Specification")
	_backupsearchCmd.Flags().StringVarP(&_backupsearchItemFilters, "item-filters", "", "", "Item Filters")
	_backupsearchCmd.Flags().StringVarP(&_backupsearchMaxResults, "max-results", "", "", "Max Results")
	_backupsearchCmd.Flags().StringVarP(&_backupsearchName, "name", "", "", "Name")
	_backupsearchCmd.Flags().StringVarP(&_backupsearchNextToken, "next-token", "", "", "Next Token")
	_backupsearchCmd.Flags().StringVarP(&_backupsearchResourceArn, "resource-arn", "", "", "Resource ARN")
	_backupsearchCmd.Flags().StringVarP(&_backupsearchRoleArn, "role-arn", "", "", "Role ARN")
	_backupsearchCmd.Flags().StringVarP(&_backupsearchSearchJobIdentifier, "search-job-identifier", "", "", "Search Job Identifier")
	_backupsearchCmd.Flags().StringVarP(&_backupsearchSearchScope, "search-scope", "", "", "Search Scope")
	_backupsearchCmd.Flags().StringVarP(&_backupsearchStatus, "status", "", "", "Status")
	_backupsearchCmd.Flags().StringSliceVarP(&_backupsearchTagKeys, "tag-keys", "", nil, "Tag Keys")
	_backupsearchCmd.Flags().StringVarP(&_backupsearchTags, "tags", "", "", "Tags")

	_backupsearchCmd.Flags().BoolVarP(&_backupsearchGetSearchJob, "get-search-job", "", false, "Get Search Job")
	_backupsearchCmd.Flags().BoolVarP(&_backupsearchGetSearchResultExportJob, "get-search-result-export-job", "", false, "Get Search Result Export Job")
	_backupsearchCmd.Flags().BoolVarP(&_backupsearchListSearchJobBackups, "list-search-job-backups", "", false, "List Search Job Backups")
	_backupsearchCmd.Flags().BoolVarP(&_backupsearchListSearchJobResults, "list-search-job-results", "", false, "List Search Job Results")
	_backupsearchCmd.Flags().BoolVarP(&_backupsearchListSearchJobs, "list-search-jobs", "", false, "List Search Jobs")
	_backupsearchCmd.Flags().BoolVarP(&_backupsearchListSearchResultExportJobs, "list-search-result-export-jobs", "", false, "List Search Result Export Jobs")
	_backupsearchCmd.Flags().BoolVarP(&_backupsearchListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_backupsearchCmd.Flags().BoolVarP(&_backupsearchStartSearchJob, "start-search-job", "", false, "Start Search Job")
	_backupsearchCmd.Flags().BoolVarP(&_backupsearchStartSearchResultExportJob, "start-search-result-export-job", "", false, "Start Search Result Export Job")
	_backupsearchCmd.Flags().BoolVarP(&_backupsearchStopSearchJob, "stop-search-job", "", false, "Stop Search Job")
	_backupsearchCmd.Flags().BoolVarP(&_backupsearchTagResource, "tag-resource", "", false, "Tag Resource")
	_backupsearchCmd.Flags().BoolVarP(&_backupsearchUntagResource, "untag-resource", "", false, "Untag Resource")

}
