package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codegurusecurity"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// codegurusecurityCmd represents the codegurusecurity command
var _codegurusecurityCmd = &cobra.Command{
	Use:   "codegurusecurity",
	Short: "AWS codegurusecurity CLI",
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
		client := codegurusecurity.NewFromConfig(cfg)
		if _codegurusecurityBatchGetFindings {
			codegurusecurity_BatchGetFindings(cfg, client)
			return
		}
		if _codegurusecurityCreateScan {
			codegurusecurity_CreateScan(cfg, client)
			return
		}
		if _codegurusecurityCreateUploadUrl {
			codegurusecurity_CreateUploadUrl(cfg, client)
			return
		}
		if _codegurusecurityGetAccountConfiguration {
			codegurusecurity_GetAccountConfiguration(cfg, client)
			return
		}
		if _codegurusecurityGetFindings {
			codegurusecurity_GetFindings(cfg, client)
			return
		}
		if _codegurusecurityGetMetricsSummary {
			codegurusecurity_GetMetricsSummary(cfg, client)
			return
		}
		if _codegurusecurityGetScan {
			codegurusecurity_GetScan(cfg, client)
			return
		}
		if _codegurusecurityListFindingsMetrics {
			codegurusecurity_ListFindingsMetrics(cfg, client)
			return
		}
		if _codegurusecurityListScans {
			codegurusecurity_ListScans(cfg, client)
			return
		}
		if _codegurusecurityListTagsForResource {
			codegurusecurity_ListTagsForResource(cfg, client)
			return
		}
		if _codegurusecurityTagResource {
			codegurusecurity_TagResource(cfg, client)
			return
		}
		if _codegurusecurityUntagResource {
			codegurusecurity_UntagResource(cfg, client)
			return
		}
		if _codegurusecurityUpdateAccountConfiguration {
			codegurusecurity_UpdateAccountConfiguration(cfg, client)
			return
		}

	},
}

var (
	_codegurusecurityBatchGetFindings           bool
	_codegurusecurityCreateScan                 bool
	_codegurusecurityCreateUploadUrl            bool
	_codegurusecurityGetAccountConfiguration    bool
	_codegurusecurityGetFindings                bool
	_codegurusecurityGetMetricsSummary          bool
	_codegurusecurityGetScan                    bool
	_codegurusecurityListFindingsMetrics        bool
	_codegurusecurityListScans                  bool
	_codegurusecurityListTagsForResource        bool
	_codegurusecurityTagResource                bool
	_codegurusecurityUntagResource              bool
	_codegurusecurityUpdateAccountConfiguration bool

	_codegurusecurityAnalysisType       string
	_codegurusecurityClientToken        string
	_codegurusecurityDate               string
	_codegurusecurityEncryptionConfig   string
	_codegurusecurityEndDate            string
	_codegurusecurityFindingIdentifiers string
	_codegurusecurityMaxResults         string
	_codegurusecurityNextToken          string
	_codegurusecurityResourceArn        string
	_codegurusecurityResourceId         string
	_codegurusecurityRunId              string
	_codegurusecurityScanName           string
	_codegurusecurityScanType           string
	_codegurusecurityStartDate          string
	_codegurusecurityStatus             string
	_codegurusecurityTagKeys            []string
	_codegurusecurityTags               string
)

// Returns a list of requested findings from standard scans.
func codegurusecurity_BatchGetFindings(cfg aws.Config, client *codegurusecurity.Client) {
	input := &codegurusecurity.BatchGetFindingsInput{
		// FindingIdentifiers: []types.FindingIdentifier, // Required
	}

	if len(_codegurusecurityFindingIdentifiers) > 0 {
		if err := assignInputField(input, "FindingIdentifiers", _codegurusecurityFindingIdentifiers); err != nil {
			log.Errorf("invalid --finding-identifiers: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchGetFindings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use to create a scan using code uploaded to an Amazon S3 bucket.
func codegurusecurity_CreateScan(cfg aws.Config, client *codegurusecurity.Client) {
	input := &codegurusecurity.CreateScanInput{
		// ResourceId: types.ResourceId, // Required
		// ScanName: *string, // Required
	}

	if len(_codegurusecurityResourceId) > 0 {
		if err := assignInputField(input, "ResourceId", _codegurusecurityResourceId); err != nil {
			log.Errorf("invalid --resource-id: %s", err.Error())
			return
		}
	}
	if len(_codegurusecurityScanName) > 0 {
		input.ScanName = aws.String(_codegurusecurityScanName)
	}
	if len(_codegurusecurityAnalysisType) > 0 {
		if err := assignInputField(input, "AnalysisType", _codegurusecurityAnalysisType); err != nil {
			log.Errorf("invalid --analysis-type: %s", err.Error())
			return
		}
	}
	if len(_codegurusecurityClientToken) > 0 {
		input.ClientToken = aws.String(_codegurusecurityClientToken)
	}
	if len(_codegurusecurityScanType) > 0 {
		if err := assignInputField(input, "ScanType", _codegurusecurityScanType); err != nil {
			log.Errorf("invalid --scan-type: %s", err.Error())
			return
		}
	}
	if len(_codegurusecurityTags) > 0 {
		if err := assignInputField(input, "Tags", _codegurusecurityTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateScan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Generates a pre-signed URL, request headers used to upload a code resource, and
// code artifact identifier for the uploaded resource.
//
// You can upload your code resource to the URL with the request headers using any
// HTTP client.
func codegurusecurity_CreateUploadUrl(cfg aws.Config, client *codegurusecurity.Client) {
	input := &codegurusecurity.CreateUploadUrlInput{
		// ScanName: *string, // Required
	}

	if len(_codegurusecurityScanName) > 0 {
		input.ScanName = aws.String(_codegurusecurityScanName)
	}

	if resp, err := client.CreateUploadUrl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use to get the encryption configuration for an account.
func codegurusecurity_GetAccountConfiguration(cfg aws.Config, client *codegurusecurity.Client) {
	input := &codegurusecurity.GetAccountConfigurationInput{}

	if resp, err := client.GetAccountConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of all findings generated by a particular scan.
func codegurusecurity_GetFindings(cfg aws.Config, client *codegurusecurity.Client) {
	input := &codegurusecurity.GetFindingsInput{
		// ScanName: *string, // Required
	}

	if len(_codegurusecurityScanName) > 0 {
		input.ScanName = aws.String(_codegurusecurityScanName)
	}
	if len(_codegurusecurityMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codegurusecurityMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codegurusecurityNextToken) > 0 {
		input.NextToken = aws.String(_codegurusecurityNextToken)
	}
	if len(_codegurusecurityStatus) > 0 {
		if err := assignInputField(input, "Status", _codegurusecurityStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetFindings(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codegurusecurity.GetFindingsOutput
	p := codegurusecurity.NewGetFindingsPaginator(client, input)
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

// Returns a summary of metrics for an account from a specified date, including
// number of open findings, the categories with most findings, the scans with most
// open findings, and scans with most open critical findings.
func codegurusecurity_GetMetricsSummary(cfg aws.Config, client *codegurusecurity.Client) {
	input := &codegurusecurity.GetMetricsSummaryInput{
		// Date: *time.Time, // Required
	}

	if len(_codegurusecurityDate) > 0 {
		if err := assignInputField(input, "Date", _codegurusecurityDate); err != nil {
			log.Errorf("invalid --date: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetMetricsSummary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns details about a scan, including whether or not a scan has completed.
func codegurusecurity_GetScan(cfg aws.Config, client *codegurusecurity.Client) {
	input := &codegurusecurity.GetScanInput{
		// ScanName: *string, // Required
	}

	if len(_codegurusecurityScanName) > 0 {
		input.ScanName = aws.String(_codegurusecurityScanName)
	}
	if len(_codegurusecurityRunId) > 0 {
		input.RunId = aws.String(_codegurusecurityRunId)
	}

	if resp, err := client.GetScan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns metrics about all findings in an account within a specified time range.
func codegurusecurity_ListFindingsMetrics(cfg aws.Config, client *codegurusecurity.Client) {
	input := &codegurusecurity.ListFindingsMetricsInput{
		// EndDate: *time.Time, // Required
		// StartDate: *time.Time, // Required
	}

	if len(_codegurusecurityEndDate) > 0 {
		if err := assignInputField(input, "EndDate", _codegurusecurityEndDate); err != nil {
			log.Errorf("invalid --end-date: %s", err.Error())
			return
		}
	}
	if len(_codegurusecurityStartDate) > 0 {
		if err := assignInputField(input, "StartDate", _codegurusecurityStartDate); err != nil {
			log.Errorf("invalid --start-date: %s", err.Error())
			return
		}
	}
	if len(_codegurusecurityMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codegurusecurityMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codegurusecurityNextToken) > 0 {
		input.NextToken = aws.String(_codegurusecurityNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListFindingsMetrics(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codegurusecurity.ListFindingsMetricsOutput
	p := codegurusecurity.NewListFindingsMetricsPaginator(client, input)
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

// Returns a list of all scans in an account. Does not return EXPRESS scans.
func codegurusecurity_ListScans(cfg aws.Config, client *codegurusecurity.Client) {
	input := &codegurusecurity.ListScansInput{}

	if len(_codegurusecurityMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codegurusecurityMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codegurusecurityNextToken) > 0 {
		input.NextToken = aws.String(_codegurusecurityNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListScans(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codegurusecurity.ListScansOutput
	p := codegurusecurity.NewListScansPaginator(client, input)
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

// Returns a list of all tags associated with a scan.
func codegurusecurity_ListTagsForResource(cfg aws.Config, client *codegurusecurity.Client) {
	input := &codegurusecurity.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_codegurusecurityResourceArn) > 0 {
		input.ResourceArn = aws.String(_codegurusecurityResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use to add one or more tags to an existing scan.
func codegurusecurity_TagResource(cfg aws.Config, client *codegurusecurity.Client) {
	input := &codegurusecurity.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_codegurusecurityResourceArn) > 0 {
		input.ResourceArn = aws.String(_codegurusecurityResourceArn)
	}
	if len(_codegurusecurityTags) > 0 {
		if err := assignInputField(input, "Tags", _codegurusecurityTags); err != nil {
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

// Use to remove one or more tags from an existing scan.
func codegurusecurity_UntagResource(cfg aws.Config, client *codegurusecurity.Client) {
	input := &codegurusecurity.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_codegurusecurityResourceArn) > 0 {
		input.ResourceArn = aws.String(_codegurusecurityResourceArn)
	}
	if len(_codegurusecurityTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _codegurusecurityTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use to update the encryption configuration for an account.
func codegurusecurity_UpdateAccountConfiguration(cfg aws.Config, client *codegurusecurity.Client) {
	input := &codegurusecurity.UpdateAccountConfigurationInput{
		// EncryptionConfig: *types.EncryptionConfig, // Required
	}

	if len(_codegurusecurityEncryptionConfig) > 0 {
		if err := assignInputField(input, "EncryptionConfig", _codegurusecurityEncryptionConfig); err != nil {
			log.Errorf("invalid --encryption-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAccountConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_codegurusecurityCmd)
	_codegurusecurityCmd.Flags().SortFlags = false

	_codegurusecurityCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_codegurusecurityCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_codegurusecurityCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_codegurusecurityCmd.Flags().StringVarP(&_codegurusecurityAnalysisType, "analysis-type", "", "", "Analysis Type")
	_codegurusecurityCmd.Flags().StringVarP(&_codegurusecurityClientToken, "client-token", "", "", "Client Token")
	_codegurusecurityCmd.Flags().StringVarP(&_codegurusecurityDate, "date", "", "", "Date")
	_codegurusecurityCmd.Flags().StringVarP(&_codegurusecurityEncryptionConfig, "encryption-config", "", "", "Encryption Config")
	_codegurusecurityCmd.Flags().StringVarP(&_codegurusecurityEndDate, "end-date", "", "", "End Date")
	_codegurusecurityCmd.Flags().StringVarP(&_codegurusecurityFindingIdentifiers, "finding-identifiers", "", "", "Finding Identifiers")
	_codegurusecurityCmd.Flags().StringVarP(&_codegurusecurityMaxResults, "max-results", "", "", "Max Results")
	_codegurusecurityCmd.Flags().StringVarP(&_codegurusecurityNextToken, "next-token", "", "", "Next Token")
	_codegurusecurityCmd.Flags().StringVarP(&_codegurusecurityResourceArn, "resource-arn", "", "", "Resource ARN")
	_codegurusecurityCmd.Flags().StringVarP(&_codegurusecurityResourceId, "resource-id", "", "", "Resource ID")
	_codegurusecurityCmd.Flags().StringVarP(&_codegurusecurityRunId, "run-id", "", "", "Run ID")
	_codegurusecurityCmd.Flags().StringVarP(&_codegurusecurityScanName, "scan-name", "", "", "Scan Name")
	_codegurusecurityCmd.Flags().StringVarP(&_codegurusecurityScanType, "scan-type", "", "", "Scan Type")
	_codegurusecurityCmd.Flags().StringVarP(&_codegurusecurityStartDate, "start-date", "", "", "Start Date")
	_codegurusecurityCmd.Flags().StringVarP(&_codegurusecurityStatus, "status", "", "", "Status")
	_codegurusecurityCmd.Flags().StringSliceVarP(&_codegurusecurityTagKeys, "tag-keys", "", nil, "Tag Keys")
	_codegurusecurityCmd.Flags().StringVarP(&_codegurusecurityTags, "tags", "", "", "Tags")

	_codegurusecurityCmd.Flags().BoolVarP(&_codegurusecurityBatchGetFindings, "batch-get-findings", "", false, "Batch Get Findings")
	_codegurusecurityCmd.Flags().BoolVarP(&_codegurusecurityCreateScan, "create-scan", "", false, "Create Scan")
	_codegurusecurityCmd.Flags().BoolVarP(&_codegurusecurityCreateUploadUrl, "create-upload-url", "", false, "Create Upload URL")
	_codegurusecurityCmd.Flags().BoolVarP(&_codegurusecurityGetAccountConfiguration, "get-account-configuration", "", false, "Get Account Configuration")
	_codegurusecurityCmd.Flags().BoolVarP(&_codegurusecurityGetFindings, "get-findings", "", false, "Get Findings")
	_codegurusecurityCmd.Flags().BoolVarP(&_codegurusecurityGetMetricsSummary, "get-metrics-summary", "", false, "Get Metrics Summary")
	_codegurusecurityCmd.Flags().BoolVarP(&_codegurusecurityGetScan, "get-scan", "", false, "Get Scan")
	_codegurusecurityCmd.Flags().BoolVarP(&_codegurusecurityListFindingsMetrics, "list-findings-metrics", "", false, "List Findings Metrics")
	_codegurusecurityCmd.Flags().BoolVarP(&_codegurusecurityListScans, "list-scans", "", false, "List Scans")
	_codegurusecurityCmd.Flags().BoolVarP(&_codegurusecurityListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_codegurusecurityCmd.Flags().BoolVarP(&_codegurusecurityTagResource, "tag-resource", "", false, "Tag Resource")
	_codegurusecurityCmd.Flags().BoolVarP(&_codegurusecurityUntagResource, "untag-resource", "", false, "Untag Resource")
	_codegurusecurityCmd.Flags().BoolVarP(&_codegurusecurityUpdateAccountConfiguration, "update-account-configuration", "", false, "Update Account Configuration")

}
